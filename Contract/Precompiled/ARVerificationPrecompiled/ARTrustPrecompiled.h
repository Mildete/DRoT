#pragma once
#include "libprecompiled/Common.h"
#include "libprecompiled/Precompiled.h"


namespace dev
{
namespace precompiled
{
class ARTrustPrecompiled: public dev::precompiled::Precompiled
{
public:
    typedef std::shared_ptr<ARTrustPrecompiled> Ptr;
    ARTrustPrecompiled();
    virtual ~ARTrustPrecompiled(){};

    

//    PrecompiledExecResult::Ptr call(std::shared_ptr<dev::blockverifier::ExecutiveContext> _context,
//        bytesConstRef _param, Address const& _origin = Address(),
//        Address const& _sender = Address()) override;
     PrecompiledExecResult::Ptr call(std::shared_ptr<dev::blockverifier::ExecutiveContext> context,
               bytesConstRef param, Address const& origin = Address(),
               Address const& sender = Address()) override;

private:
    // 添加verifyAR函数声明
    void verifyAR(bytesConstRef _paramData, PrecompiledExecResult::Ptr _callResult);
};
}
}



