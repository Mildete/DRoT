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
#include "ARTrustPrecompiled.h"
#include <libprecompiled/Common.h>

#include <libblockverifier/ExecutiveContext.h>
#include <libethcore/ABI.h>

using namespace dev::storage;

using namespace dev;
using namespace dev::eth;
using namespace dev::precompiled;
using namespace dev::blockverifier;

// verify interface
const char* const AR_VERIFY_METHOD_STR = "verifyAR(string,uint256,string,string,uint256)";

// Constructor
ARTrustPrecompiled::ARTrustPrecompiled()
{
    name2Selector[AR_VERIFY_METHOD_STR] = getFuncSelector(AR_VERIFY_METHOD_STR);
}

// Precompiled `call` function
 // std::shared_ptr<ExecutiveContext> _context, bytesConstRef _param, Address const& _origin, Address const&)
PrecompiledExecResult::Ptr ARTrustPrecompiled::call(
     ExecutiveContext::Ptr, bytesConstRef param, Address const&, Address const&)
{
    PRECOMPILED_LOG(TRACE) << LOG_BADGE("ARTrustPrecompiled") << LOG_DESC("call")
                           << LOG_KV("param", toHex(param));

    // Parse the function selector
    uint32_t func = getParamFunc(param);
    auto callResult = m_precompiledExecResultFactory->createPrecompiledResult();
    callResult->gasPricer()->setMemUsed(param.size());
    dev::eth::ContractABI abi;

    if (func == name2Selector[AR_VERIFY_METHOD_STR])
    {
        verifyAR(getParamData(param), callResult);
    }
    else
    {
        PRECOMPILED_LOG(ERROR) << LOG_BADGE("ARTrustPrecompiled") << LOG_DESC("Unknown function call")
                               << LOG_KV("func", func);
        callResult->setExecResult(abi.abiIn("", u256(int32_t(CODE_UNKNOW_FUNCTION_CALL))));
    }

    return callResult;
}

// Function to verify the attestation report
void ARTrustPrecompiled::verifyAR(bytesConstRef _paramData, PrecompiledExecResult::Ptr _callResult)
{
    ContractABI abi;
    try
    {
        std::string nodeID, measurement, referenceMeasurement;
        u256 attestationTime, reliabilityThreshold;

        // Parse the parameters
        abi.abiOut(_paramData, nodeID, attestationTime, measurement, referenceMeasurement, reliabilityThreshold);

        // Step 1: Verify measurement and referenceMeasurement equality
        if (measurement != referenceMeasurement)
        {
            _callResult->setExecResult(abi.abiIn("", false, 0, std::string("Measurement verification failed")));
            return;
        }

        // Step 2: Calculate the reliability score
        time_t currentTime = std::time(nullptr);
        double timeDiff = std::abs(difftime(currentTime, static_cast<time_t>(attestationTime)));
        float reliabilityScore = (timeDiff < 60) ? 1.0 : std::exp(-0.01 * (timeDiff - 60));

        // Step 3: Scale reliabilityScore to a large integer (e.g., 1e6) for comparison
        uint64_t scaledReliabilityScore = static_cast<uint64_t>(reliabilityScore * 1e6);

        // Convert reliabilityThreshold to the same scale (multiplied by 1e6)
        u256 scaledReliabilityThreshold = reliabilityThreshold * u256(1e4);

        // Step 4: Compare the scaled reliability score with the threshold
        if (u256(scaledReliabilityScore) < scaledReliabilityThreshold)
        {
            _callResult->setExecResult(abi.abiIn("", false, 0, std::string("Reliability score below threshold")));
            return;
        }

        // Verification successful
        _callResult->setExecResult(abi.abiIn("", true, u256(scaledReliabilityScore), std::string("AR Verification successful")));
    }
    catch (std::exception const& e)
    {
        _callResult->setExecResult(abi.abiIn("", false, 0, std::string("Exception occurred")));
    }
} //一个扩大1e4 一个扩大1e6 最后参数取1-100