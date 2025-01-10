pragma solidity ^0.4.25;
import "./BlsVerifyPrecompiled.sol";

contract BlsVerifyPrecompiledHelper {
    BlsVerifyPrecompiled test;
    function BlsVerifyPrecompiledHelper() {

        test = BlsVerifyPrecompiled(0x5023);
    }
    function blsverify(string pk_str, string agg_sig_str, uint256 num_user) public returns(string){
        return test.blsverify(pk_str,agg_sig_str,num_user);
    }

}
