pragma solidity ^0.4.25;
contract BlsVerifyPrecompiled{
  function blsverify(string pk_str, string agg_sig_str, uint256 num_user) public returns(string);
}
