#pragma once
#include "libprecompiled/Common.h"
#include "libprecompiled/Precompiled.h"
#include <pbc/pbc.h>
#include <vector>

namespace dev
{
    namespace precompiled
    {
        class BlsVerifyPrecompiled: public dev::precompiled::Precompiled
        {
            // 定义BLS签名结构
            struct BLSSignature {
                element_t sig; // 签名
            };
        public:
            typedef std::shared_ptr<BlsVerifyPrecompiled> Ptr;
            BlsVerifyPrecompiled();
            virtual ~BlsVerifyPrecompiled(){};
            std::string toString() override;


//    PrecompiledExecResult::Ptr call(std::shared_ptr<dev::blockverifier::ExecutiveContext> _context,
//        bytesConstRef _param, Address const& _origin = Address(),
//        Address const& _sender = Address()) override;
            PrecompiledExecResult::Ptr call(std::shared_ptr<dev::blockverifier::ExecutiveContext> context,
                                            bytesConstRef param, Address const& origin = Address(),
                                            Address const& sender = Address()) override;

        private:
            bool verify_aggregated_signature(
                    pairing_t pairing,
                    element_t g,
                    const std::string& pk_string_combined,
                    const std::vector<const char*>& messages,
                    const std::string& aggregated_signature_string
            );
            std::vector<std::string> split_string(const std::string& str, char delimiter);
            void string_to_element(element_t& elem, const std::string& str);
            void string_to_signature(BLSSignature& sig, const std::string& str, pairing_t pairing);
        };
    }
}



