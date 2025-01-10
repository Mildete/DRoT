pragma solidity ^0.4.25;
contract VerifyEcdsaPrecompiled{
    function VerifyEcdsa(string message, string rHex, string sHex, string pubKeyHex) public constant returns(bool);
}
