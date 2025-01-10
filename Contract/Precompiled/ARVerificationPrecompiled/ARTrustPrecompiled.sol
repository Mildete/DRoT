pragma solidity ^0.4.25;
contract ARTrustPrecompiled{
    function verifyAR(string nodeID, uint256 attestationTime, string measurement, string referenceMeasurement, uint256 reliabilityThreshold) public constant returns(bool,uint256,string);
}
