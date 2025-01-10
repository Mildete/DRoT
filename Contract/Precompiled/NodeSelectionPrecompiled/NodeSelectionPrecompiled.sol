pragma solidity ^0.4.25;
pragma experimental ABIEncoderV2;
contract NodeSelectionPrecompiled{
  function selectNode(string[] nodeIDs, uint256[] trustScores, uint256 randomValue) public constant returns(string);
}
