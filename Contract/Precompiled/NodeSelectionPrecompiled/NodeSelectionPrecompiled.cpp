/*
 * @CopyRight:
 * FISCO-BCOS is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * FISCO-BCOS is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with FISCO-BCOS.  If not, see <http://www.gnu.org/licenses/>
 * (c) 2016-2019 fisco-dev contributors.
 */

#include <iostream>
#include <openssl/evp.h>
#include <openssl/sha.h>
#include <openssl/bn.h>
#include <openssl/obj_mac.h>
#include "libethcore/ABI.h"
#include "NodeSelectionPrecompiled.h"
#include <libprecompiled/Common.h>

#include <libblockverifier/ExecutiveContext.h>
#include <libethcore/ABI.h>

#include <openssl/ecdsa.h>
#include <openssl/obj_mac.h>
#include <openssl/sha.h>
#include <string>
#include <vector>
#include <iomanip>

using namespace dev::storage;

using namespace dev;
using namespace dev::eth;
using namespace dev::precompiled;
using namespace dev::blockverifier;

// add interface
const char* const METHOD_NAME = "selectNode(string[],uint256[],uint256)";

NodeSelectionPrecompiled::NodeSelectionPrecompiled()
{
    name2Selector[METHOD_NAME] = getFuncSelector(METHOD_NAME);  //构造函数内，把函数名转换为哈希值，相当于做一个函数签名，取前几个字节作为函数选择器
}
std::string NodeSelectionPrecompiled::toString()
{
    return "NodeSelection";
}
// 定义一个结构体来存储节点的ID和信任分数
struct Node {
    std::string id;
    u256 trustScore;
};

PrecompiledExecResult::Ptr NodeSelectionPrecompiled::call(
        ExecutiveContext::Ptr, bytesConstRef param, Address const&, Address const&)
{
    PRECOMPILED_LOG(TRACE) << LOG_BADGE("NodeSelectionPrecompiled") << LOG_DESC("call")
                           << LOG_KV("param", toHex(param));

    // parse function name
    uint32_t func = getParamFunc(param); // sol调用函数，前四个字节是函数签名，用来做函数选择
    bytesConstRef data = getParamData(param);
    auto callResult = m_precompiledExecResultFactory->createPrecompiledResult();
    callResult->gasPricer()->setMemUsed(param.size());
    dev::eth::ContractABI abi;

    if (func == name2Selector[METHOD_NAME])
    {  //  function call

        // 解析传递过来的参数
        std::vector<std::string> nodeIDs;
        std::vector<u256> trustScores;
        u256 randomValue;
        abi.abiOut(data, nodeIDs, trustScores,randomValue);
        // 将节点ID和信任分数打包为Node结构体
        std::vector<Node> nodes;
        for (size_t i = 0; i < nodeIDs.size(); ++i) {
            nodes.push_back({nodeIDs[i], trustScores[i]});
        }
        std::string ret_result;
        try {
            // 调用选择节点函数
            ret_result = selectNodeByTrust(nodes, randomValue);
        } catch (const std::exception& e) {
            ret_result = e.what();
        }
        callResult->setExecResult(abi.abiIn("", ret_result));
    }
    else
    {  // unknown function call
        PRECOMPILED_LOG(ERROR) << LOG_BADGE("NodeselectionPrecompiled") << LOG_DESC(" unknown func ")
                               << LOG_KV("func", func);
        callResult->setExecResult(abi.abiIn("", u256(int32_t(CODE_UNKNOW_FUNCTION_CALL))));
    }

    return callResult;
}

// 根据信任分数和随机数进行节点选择
std::string NodeSelectionPrecompiled::selectNodeByTrust(const std::vector<Node>& nodes, const u256& randomValue) {
    // 计算所有信任分数的总和
    u256 totalTrust = std::accumulate(nodes.begin(), nodes.end(), u256(0),
                                          [](u256 sum, const Node& node) { return sum + node.trustScore; });

    // 检查随机数是否在合法范围内
    if (randomValue < 0 || randomValue >= totalTrust) {
        throw std::invalid_argument("Random value is out of range.");
    }

    // 遍历节点并找到落入随机数区间的节点
    u256 cumulativeTrust = 0;
    for (const auto& node : nodes) {
        cumulativeTrust += node.trustScore;
        if (randomValue < cumulativeTrust) {
            return node.id;  // 选中的节点
        }
    }

    // 理论上不应该执行到这一步，除非输入不符合预期
    return "error selection";
}