pragma solidity ^0.4.25;

contract NodeTable {
    /**
     *@brief 打开表，返回Table合约地址
     *@param tableName 表的名称
     *@return 返回Table的地址，当表不存在时，将会返回空地址即address(0x0)
     */
    function openTable(string memory) public view returns (Table) {} //open table


/**
 *@brief 创建表，返回是否成功
     *@param tableName 表的名称
     *@param key 表的主键名
     *@param valueFields 表的字段名，多个字段名以英文逗号分隔
     *@return 返回错误码，成功为0，错误则为负数
     */
    function createTable(string memory, string memory, string memory) public returns (int256) {} //create table
}

// 查询条件
contract Condition {
//等于
    function EQ(string memory, int256) public {}
    function EQ(string memory, string memory) public {}

//不等于
    function NE(string memory, int256) public {}
    function NE(string memory, string memory) public {}

//大于
    function GT(string, int) public;
    //大于或等于
    function GE(string memory, int256) public {}



//小于
    function LT(string memory, int256) public {}
//小于或等于
    function LE(string memory, int256) public {}

//限制返回记录条数
    function limit(int256) public {}
    function limit(int256, int256) public {}
}

// 单条数据记录
contract Entry {
    function getInt(string memory) public view returns (int256) {}
    function getUInt(string memory) public view returns (uint256) {}
    function getAddress(string memory) public view returns (address) {}
    function getBytes64(string memory) public view returns (bytes1[64] memory) {}
    function getBytes32(string memory) public view returns (bytes32) {}
    function getString(string memory) public view returns (string memory) {}

    function set(string memory, int256) public {}
    function set(string memory, uint256) public {}
    function set(string memory, string memory) public {}
    function set(string memory, address) public {}
}

// 数据记录集
contract Entries {
    function get(int256) public view returns (Entry) {}
    function size() public view returns (int256) {}
}

// Table主类
contract Table {
/**
 *@brief 查询接口
     *@param key 查询主键值
     *@param cond 查询条件
     *@return Entries合约地址，合约地址一定存在
     */
    function select(string memory, Condition) public view returns (Entries) {}
/**
 *@brief 插入接口
     *@param key 插入主键值
     *@param entry 插入字段值
     *@return 插入影响的行数
     */
    function insert(string memory, Entry) public returns (int256) {}
/**
 *@brief 更新接口
     *@param key 更新主键值
     *@param entry 更新字段值
     *@param cond 更新条件
     *@return 更新影响的行数
     */
    function update(string memory, Entry, Condition) public returns (int256) {}
/**
 *@brief 删除接口
     *@param key 删除的主键值
     *@param cond 删除条件
     *@return 删除影响的行数
     */
    function remove(string memory, Condition) public returns (int256) {}

    function newEntry() public view returns (Entry) {}
    function newCondition() public view returns (Condition) {}
}