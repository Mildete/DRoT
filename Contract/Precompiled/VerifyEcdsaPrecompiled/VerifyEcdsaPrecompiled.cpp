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
#include "VerifyEcdsaPrecompiled.h"
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
const char* const METHOD_NAME = "VerifyEcdsa(string,string,string,string)";

VerifyEcdsaPrecompiled::VerifyEcdsaPrecompiled()
{
    name2Selector[METHOD_NAME] = getFuncSelector(METHOD_NAME);  //构造函数内，把函数名转换为哈希值，相当于做一个函数签名，取前几个字节作为函数选择器
}
std::string VerifyEcdsaPrecompiled::toString()
{
    return "VerifyEcdsaSignature";
}

PrecompiledExecResult::Ptr VerifyEcdsaPrecompiled::call(
        ExecutiveContext::Ptr, bytesConstRef param, Address const&, Address const&)
{
    PRECOMPILED_LOG(TRACE) << LOG_BADGE("VerifyEcdsaPrecompiled") << LOG_DESC("call")
                           << LOG_KV("param", toHex(param));

    // parse function name
    uint32_t func = getParamFunc(param); // sol调用函数，前四个字节是函数签名，用来做函数选择
    bytesConstRef data = getParamData(param);
    auto callResult = m_precompiledExecResultFactory->createPrecompiledResult();
    callResult->gasPricer()->setMemUsed(param.size());
    dev::eth::ContractABI abi;

    if (func == name2Selector[METHOD_NAME])
    {  //  function call

        std::string message, rHex, sHex,pubKeyHex;
        bool ret_result;
        abi.abiOut(data, message, rHex, sHex,pubKeyHex);
        ret_result = verifyECDSASignature(message, rHex, sHex,pubKeyHex);
        callResult->setExecResult(abi.abiIn("", ret_result));

    }
    else
    {  // unknown function call
        PRECOMPILED_LOG(ERROR) << LOG_BADGE("VerifyEcdsaPrecompiled") << LOG_DESC(" unknown func ")
                               << LOG_KV("func", func);
        callResult->setExecResult(abi.abiIn("", u256(int32_t(CODE_UNKNOW_FUNCTION_CALL))));
    }

    return callResult;
}


// ECDSA 签名验证函数
bool VerifyEcdsaPrecompiled::verifyECDSASignature(const std::string& message, const std::string& rHex, const std::string& sHex, const std::string& pubKeyHex) {
    // 对消息进行哈希 (与 Go 保持一致)
    uint8_t hash[SHA256_DIGEST_LENGTH];
    SHA256((const uint8_t*)message.c_str(), message.length(), hash);

    // 将 r, s, pubKey 从 hex 转为字节数组
    std::vector<uint8_t> rBytes = hexStringToBytes(rHex);
    std::vector<uint8_t> sBytes = hexStringToBytes(sHex);
    std::vector<uint8_t> pubKeyBytes = hexStringToBytes(pubKeyHex);

    // 检查 r 和 s 的长度是否为 32 字节
    if (rBytes.size() != 32 || sBytes.size() != 32) {
        std::cerr << "Invalid r or s length" << std::endl;
        return false;
    }

    // 将公钥添加 0x04 前缀，标识为未压缩公钥
    std::vector<uint8_t> fullPubKey = {0x04};
    fullPubKey.insert(fullPubKey.end(), pubKeyBytes.begin(), pubKeyBytes.end());

    // 检查公钥的长度是否为 65 字节（0x04 + 64 字节公钥）
    if (fullPubKey.size() != 65) {
        std::cerr << "Invalid public key length" << std::endl;
        return false;
    }

    // 将 r, s 解析为 BIGNUM
    BIGNUM* r = BN_bin2bn(rBytes.data(), rBytes.size(), NULL);
    BIGNUM* s = BN_bin2bn(sBytes.data(), sBytes.size(), NULL);

    // 生成 EC_KEY 对象
    EC_GROUP* group = EC_GROUP_new_by_curve_name(NID_X9_62_prime256v1);
    EC_KEY* eckey = EC_KEY_new();
    EC_KEY_set_group(eckey, group);

    // 设置公钥
    EC_POINT* pubKeyPoint = EC_POINT_new(group);
    EC_POINT_oct2point(group, pubKeyPoint, fullPubKey.data(), fullPubKey.size(), NULL);
    EC_KEY_set_public_key(eckey, pubKeyPoint);

    // 生成 ECDSA_SIG 结构并手动设置 r, s
    ECDSA_SIG* sig = ECDSA_SIG_new();
#if OPENSSL_VERSION_NUMBER < 0x10100000L  // 适用于 OpenSSL 1.0.x
    sig->r = r;
    sig->s = s;
#else  // OpenSSL 1.1.x 及更新版本
    ECDSA_SIG_set0(sig, r, s);
#endif

    // 验证签名
    int verifyStatus = ECDSA_do_verify(hash, SHA256_DIGEST_LENGTH, sig, eckey);

    // 清理资源
    EC_GROUP_free(group);
    EC_KEY_free(eckey);
    EC_POINT_free(pubKeyPoint);
    ECDSA_SIG_free(sig);

    // 返回验证结果
    return verifyStatus == 1;
}

// 将 hex 编码的字符串转换为字节数组
std::vector<uint8_t> VerifyEcdsaPrecompiled::hexStringToBytes(const std::string& hex) {
    std::vector<uint8_t> bytes(hex.length() / 2);
    for (size_t i = 0; i < hex.length(); i += 2) {
        std::string byteString = hex.substr(i, 2);
        bytes[i / 2] = (uint8_t)strtol(byteString.c_str(), nullptr, 16);
    }
    return bytes;
}