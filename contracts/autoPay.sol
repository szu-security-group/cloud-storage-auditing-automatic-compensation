pragma solidity >= 0.6.0;

import "./SolRsaVerify.sol";

contract AutoPayAudit{
    // the server blockchain ethernum account address
    // the client blockchain ehhernum account address
    address public server ; 
    address payable public client ;  
    
    // the pre-store assets 
    uint public asset;

    // public key
    bytes public E;
    bytes public N;

    // random number for challenge
    bytes public gs;

    // challenge tags and index
    uint[][] public indexs;
    bytes[] public chalTags;
    bool chalFlag;

    // proof of integrity from Server
    bytes[] public proofs;
    bool proveFlag;

    // result of verify
    uint result;

    constructor(address payable _client) payable{
        server  = msg.sender;
        client = _client;
        asset = msg.value;

        keyFlag = false;
        randsFlag = false;
        chalFlag = false;
        proveFlag = false;
        
    }

    function setPk(bytes memory _N, bytes memory _E) public{
        require(msg.sender == client, "Pk should set by client");
        N = _N;
        E = _E;
    }


    function setIndexAndGs(uint[][] memory _index, bytes memory _gs) public {
        require(msg.sender == client, "Index and RandomGs should set by client");
        indexs = _index;
        gs = _gs;
        randsFlag = true;
    }


    function setProve(bytes[] memory _proof) public {
        require(msg.sender == server, "Proof should set by server");
        proofs = _proof;
        proveFlag = true;

    }

    function setChals(bytes[] memory _chaltags) public {
        require(msg.sender == client, "Challenge should set by client");
        require(proveFlag, "Proof should set before challenge set");

        chalTags = _chaltags;
        chalFlag = true;

    }


    function verify() public payable{
        require(proveFlag==true && chalFlag==true, "Proof and challengeTag is not empty");
        require(chalTags.length == proofs.length, "The size of challenge and proof should be equal");

        for (uint256 index = 0; index < chalTags.length; index++) {
            result = SolRsaVerify.pkcs1Sha256VerifyRaw(proofs[index], chalTags[index], E, N);
            if (result == 0) {
                sendViaCall(client);
                // do transfer
            }
        }

    }

    function sendViaCall(address payable _to) public payable {
        (bool sent, bytes memory data) = _to.call{value: 0.01 ether}("");
        require(sent, "Failed to send Ether");
    }

    function getBalance() public view returns (uint) {
        return address(this).balance;
    }

    function getGs () public view returns (bytes memory) {
        return gs;
    }

    function getN() public view returns (bytes memory) {
        return N;
    }

    function getChIndex() public view returns (uint[][] memory) {
        return indexs;
    }
}
