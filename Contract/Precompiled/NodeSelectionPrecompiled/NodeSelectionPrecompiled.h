#pragma once
#include "libprecompiled/Common.h"
#include "libprecompiled/Precompiled.h"


namespace dev
{
    namespace precompiled
    {
        class NodeSelectionPrecompiled: public dev::precompiled::Precompiled
        {
            // 定义一个结构体来存储节点的ID和信任分数
            struct Node {
                std::string id;
                u256 trustScore;
            };
        public:
            typedef std::shared_ptr<NodeSelectionPrecompiled> Ptr;
            NodeSelectionPrecompiled();
            virtual ~NodeSelectionPrecompiled(){};
            std::string toString() override;


//    PrecompiledExecResult::Ptr call(std::shared_ptr<dev::blockverifier::ExecutiveContext> _context,
//        bytesConstRef _param, Address const& _origin = Address(),
//        Address const& _sender = Address()) override;
            PrecompiledExecResult::Ptr call(std::shared_ptr<dev::blockverifier::ExecutiveContext> context,
                                            bytesConstRef param, Address const& origin = Address(),
                                            Address const& sender = Address()) override;

        private:
            // 添加成员函数声明
            std::string selectNodeByTrust(const std::vector<Node>& nodes, const u256& randomValue);

        };
    }
}



