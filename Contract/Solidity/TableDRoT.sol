pragma solidity ^0.4.25;
import "./Table.sol";

contract TableDRoT {
    // event
    event RegisterEvent(int256 ret, string NodeID, string PublicKey, string AttestationScoreFunc, string ReferenceMeasurement);
  //  event Select(string NodeID);
    event RemoveResult(int256 count);
    event UpdateStatusResult(int256 count);
    //event TransferEvent(int256 ret, string from_account, string to_account, uint256 amount);

    string constant TABLE_NAME = "DRoT_Table_Test"; //数据库名

    constructor() public {
        // 构造函数中创建OracleNode表
        createTable();
    }

    function createTable() private {
        NodeTable nt = NodeTable(0x1010); //The fixed address is 0x1001 for TableFactory  1010是KVTable合约，table合约的表项数据类型只能是string kvtable支持 string int256和 bytes 

        // OracleNode表, key : NodeID, field : PublicKey,PrivateKey,AttestationScoreFunc,ReferenceMeasurement
        // |  NodeID (主键)     |    PublicKey     | AttestationScoreFunc | ReferenceMeasurement |TrustScore |
        // |--------------------|------------------|----------------------|--------------------- | --========|
        // |        NodeID1     | PublicKeyAddress |ExampleAttestationfunc|  Sampemeasurement    | 970045    |
        // |--------------------|------------------|----------------------|----------------------|  ----- ---|
        //
        // 创建表 // the parameters of createTable are tableName,keyField(主键),"vlaueFiled1,vlaueFiled2,vlaueFiled3,..."
        nt.createTable(TABLE_NAME, "NodeID", "PublicKey,AttestationScoreFunc,ReferenceMeasurement,TrustScore");
    }


    function openTable() private returns(Table) {
        NodeTable nt = NodeTable(0x1001);
        Table table = nt.openTable(TABLE_NAME);
        return table;
    }

    /*
    描述 : 根据Oracle 节点ID 查询公钥
    参数 ：
            NodeID：Oracle 节点ID

    返回值：
            参数一： 成功返回0, 账户不存在返回-1
            参数二： 第一个参数为0时有效，资产金额
    */
    function Select(string memory NodeID, string memory field) public constant returns(int256, string memory) {
        // 打开表
        Table table = openTable();
        // 查询
        Entries entries = table.select(NodeID, table.newCondition());
        //        uint256 asset_value = 0;
        if (0 == uint256(entries.size())) {
        return (-1, "Node is not registered");
        } else {
        Entry entry = entries.get(0);
        return (0, entry.getString(field));
        }

    }

   

    //update records
    function updateTrustScore(string memory nodeID, uint256 trustscore)
    public
    returns (int256)
    {
    Table table = openTable();

    Entry entry = table.newEntry();
    entry.set("TrustScore", trustscore);
    Condition condition = table.newCondition();
    condition.EQ("NodeID", nodeID);
    

    int256 count = table.update(nodeID,entry, condition);
    emit UpdateStatusResult(count);

    return count;
    }
/*
描述 : Oracle节点注册
参数 ：
        NodeID : 节点ID
        PublicKeyAddress : 公钥地址
        AttestationScoreFunc ：
        ReferenceMeasurement ：

返回值：
        0  注册成功
        -1 节点已存在
        -2 其他错误
*/
function register(string NodeID, string PublicKeyAddress, string AttestationScoreFunc, string ReferenceMeasurement) public returns(int256){
    int256 ret_code = 0;
    int256 ret= 0;
    string memory temp;
    // 查询账户是否存在
    (ret, temp) = Select(NodeID,"PublicKey");
    //用户不存在
    if(ret != 0) {
        Table table = openTable();

        Entry entry = table.newEntry();
        entry.set("NodeID", NodeID);
        entry.set("PublicKey", PublicKeyAddress);
        entry.set("AttestationScoreFunc", AttestationScoreFunc);
        entry.set("ReferenceMeasurement", ReferenceMeasurement);
        entry.set("TrustScore", uint256(0));
        // 插入
        int count = table.insert(NodeID, entry);
        if (count == 1) {
        // 成功
        ret_code = 0;
        } else {
        // 失败? 无权限或者其他错误
        ret_code = -2;
        }
    } else {
    // 账户已存在
    ret_code = -1;
    }

    emit RegisterEvent(ret_code, NodeID, PublicKeyAddress, AttestationScoreFunc, ReferenceMeasurement);

    return ret_code;
    }

//remove records
function remove(string NodeID, string PublicKeyAddress) public returns (int256) {
    Table table = openTable();

    Condition condition = table.newCondition();
    condition.EQ("NodeID", NodeID);
    condition.EQ("PublicKey", PublicKeyAddress);

    int256 count = table.remove(NodeID, condition);
    emit RemoveResult(count);

    return count;
}

}