pragma solidity ^0.4.25;
pragma experimental ABIEncoderV2;
import "./VerifyEcdsaPrecompiled.sol";
import "./ARTrustPrecompiled.sol";
import "./TableDRoT.sol";
contract DRoTPrecompiledHelper {
    VerifyEcdsaPrecompiled verifyEcdsa;
    ARTrustPrecompiled arTrust;
    TableDRoT tableDRoT;

     event UpdateStatusResult(int256 count);
     // 构造函数，传入 TableRegistration 合约地址
    function DRoTPrecompiledHelper(address tableRegistrationAddress) {
        // 调用预编译合约
        verifyEcdsa = VerifyEcdsaPrecompiled(0x5009);
        arTrust = ARTrustPrecompiled(0x5058);
          tableDRoT = TableDRoT(tableRegistrationAddress);
    }

    function VerifyEcdsa(string message, string rHex, string sHex, string pubKeyHex) public constant returns(bool){
    return verifyEcdsa.VerifyEcdsa(message, rHex, sHex, pubKeyHex);
    }

     function verifyAR(string nodeID, uint256 attestationTime, string measurement, string referenceMeasurement, uint256 reliabilityThreshold) public constant returns(bool,uint256,string){
        return arTrust.verifyAR(nodeID, attestationTime, measurement, referenceMeasurement, reliabilityThreshold);
     }
  function verifyARandStore(
    string memory nodeID,
    uint256 attestationTime,
    string memory measurement,
    string memory referenceMeasurement,
    uint256 reliabilityThreshold
) public returns (bool, uint256, string memory) {
    // 检查 NodeID 是否已注册
    (int256 ret, string memory temp) = tableDRoT.Select(nodeID, "PublicKey");
    require(ret == 0, "NodeID not registered in the table");

    // 调用 ARTrust 合约的验证逻辑
    (bool result, uint256 trustScore, string memory info) = arTrust.verifyAR(
        nodeID,
        attestationTime,
        measurement,
        referenceMeasurement,
        reliabilityThreshold
    );

    // 更新信任分到数据库
    int256 updateCount = tableDRoT.updateTrustScore(nodeID, trustScore);

    // 检查更新是否成功
    require(updateCount > 0, "Failed to update trust score in the table");

    // 触发事件
    emit UpdateStatusResult(updateCount);

    return (result, trustScore, info);
}



function EcdsaVerifySol(
string messageHash, // 将类型从 bytes32 改为 string
uint8 v, // 签名的 v 值
string r, // 将类型从 bytes32 改为 string
string s, // 将类型从 bytes32 改为 string
address expectedSigner // 预期的签名者地址
) public pure returns (bool) {
// // 转换 bytes 为 bytes32
bytes32 mHash = hexStringToBytes32(messageHash);
bytes32 rHash = hexStringToBytes32(r);
bytes32 sHash = hexStringToBytes32(s);

// 使用 ecrecover 从签名中恢复地址
address signer = ecrecover(mHash, v, rHash, sHash);
// 检查恢复的地址是否等于预期的地址
return signer == expectedSigner;
}

function hexStringToBytes32(string hexString) public pure returns (bytes32) {
bytes memory temp = bytes(hexString);
uint256 length = temp.length;

// 确保字符串长度为 66（0x 开头 + 64 个十六进制字符）
require(length == 66, "Invalid hex string length");

bytes32 result;

// 从索引 2 开始，因为前两位是 '0x'
for (uint256 i = 2; i < length; i += 2) {
result = (result << 8) | bytes32(fromHexChar(uint8(temp[i])) * 16 + fromHexChar(uint8(temp[i + 1])));
}

return result;
}

// 辅助函数：将单个十六进制字符转换为其数值
function fromHexChar(uint8 c) internal pure returns (uint8) {
if (c >= 48 && c <= 57) { // '0' 到 '9'
return c - 48;
} else if (c >= 97 && c <= 102) { // 'a' 到 'f'
return 10 + c - 97;
} else if (c >= 65 && c <= 70) { // 'A' 到 'F'
return 10 + c - 65;
} else {
revert("Invalid hex character");
}
}

}
