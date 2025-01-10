pragma solidity ^0.4.25;

contract EcdsaVerifySol {
     // 使用 ecrecover 恢复签名者的地址
    function verifySignature(
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
