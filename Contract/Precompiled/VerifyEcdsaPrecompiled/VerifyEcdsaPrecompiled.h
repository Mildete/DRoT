#pragma once
#include "libprecompiled/Common.h"
#include "libprecompiled/Precompiled.h"


namespace dev
{
namespace precompiled
{
class VerifyEcdsaPrecompiled: public dev::precompiled::Precompiled
{
public:
    typedef std::shared_ptr<VerifyEcdsaPrecompiled> Ptr;
    VerifyEcdsaPrecompiled();
    virtual ~VerifyEcdsaPrecompiled(){};
    std::string toString() override;
    

//    PrecompiledExecResult::Ptr call(std::shared_ptr<dev::blockverifier::ExecutiveContext> _context,
//        bytesConstRef _param, Address const& _origin = Address(),
//        Address const& _sender = Address()) override;
     PrecompiledExecResult::Ptr call(std::shared_ptr<dev::blockverifier::ExecutiveContext> context,
               bytesConstRef param, Address const& origin = Address(),
               Address const& sender = Address()) override;

private:
    // 添加成员函数声明
    bool verifyECDSASignature(const std::string& message, const std::string& rHex, const std::string& sHex, const std::string& pubKeyHex);
    std::vector<uint8_t> hexStringToBytes(const std::string& hex);
};
}
}



