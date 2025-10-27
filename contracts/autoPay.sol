/*
The smart contract was deployed on the latest version of Ethereum.
We used solidity (version 0.60) to write the smart contract.
The protocol is implemented using Golang 1.17 (blockchain offline) and solidity (blockchain online).
The modulus $N$ is 1024 bits and the size of data block is 4KB. The experimental data ranges from 100MB to 1GB.
*/


pragma solidity >= 0.6.0;
pragma experimental ABIEncoderV2;


contract AutoPayAudit{
    // the server blockchain ethernum account address
    // the client blockchain ehhernum account address
    address payable public server ;   
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
    bytes public chalTag;  
    bool chalFlag;  

    // proof of integrity from Server
    bytes public proof;   
    bool proveFlag;

    bool keyFlag;  
    bool randsFlag;  

    bytes32 public hr;
    bytes32 public clientHr;
    bytes32 public serverHr;
    bool public hrMatched;
    bool hrFlag;

    bool public disputeResult;


    constructor(address payable _client) public payable{
        server  = msg.sender;  
        client = _client;
        asset = msg.value;  
        
        keyFlag = false;
        randsFlag = false;
        chalFlag = false;
        proveFlag = false;
        hrFlag = false;
        hrMatched = false;
    }

    function setHr(bytes32 _hr) public {
        if (msg.sender == client) {
            clientHr = _hr;
        } else if (msg.sender == server) {
            serverHr = _hr;
        } else {
            revert("Only client or server can set hr");
        }
        
        if (clientHr != bytes32(0) && serverHr != bytes32(0)) {
            if (clientHr == serverHr) {
                hr = clientHr;
                hrMatched = true;
                hrFlag = true;
            } else {
                revert("Merkle roots do not match");
            }
        }
    }

    function setPk(bytes memory _N, bytes memory _E) public{
        require(msg.sender == client, "Pk should set by client");  
        require(hrFlag, "Merkle root must be set and matched first");
        N = _N;
        E = _E; 
    }


    function setIndexAndGs(uint[][] memory _index, bytes memory _gs) public {
        require(msg.sender == client, "Index and RandomGs should set by client");  
        indexs = _index;
        gs = _gs;
        randsFlag = true;
    }

    function setProve(bytes memory _proof) public {
        require(msg.sender == server, "Proof should set by server");  
        proof = _proof;
        proveFlag = true;
    }

    function setChals(bytes memory _chaltag) public {
        require(msg.sender == client, "Challenge should set by client");  
        require(proveFlag, "Proof should set before challenge set");   
        chalTag = _chaltag;   
        chalFlag = true;
    }

    function directRsaVerify(
        bytes memory _data,
        bytes memory _e,
        bytes memory _m
    ) internal view returns (bytes memory) {
        uint256 dataLen = _data.length;
        uint256 eLen = _e.length;
        uint256 mLen = _m.length;

        bytes32 bDataLen = bytes32(dataLen);
        bytes32 bELen = bytes32(eLen);
        bytes32 bMLen = bytes32(mLen);

        bytes memory input = abi.encodePacked(bDataLen, bELen, bMLen, _data, _e, _m);

        bytes memory output = new bytes(mLen);

        bool success;
        assembly {
            success := staticcall(
                gas(),
                5,
                add(input,32),
                mload(input),
                add(output,32),
                mLen
            )
        }
        require(success, "modexp precompile failed"); 

        return output;
    }

    function verify() public payable {
        require(proveFlag && chalFlag, "Invalid verification state"); 
        bytes memory rsaResult = directRsaVerify(chalTag, E, N);
        if (keccak256(rsaResult) == keccak256(proof)) {
            sendViaCall(client);   
        }
    } 


    function disputeArbitration(
        uint idx,          
        bytes calldata blockData,  
        bytes32[] calldata proof   
    ) external {
        require(hrFlag, "Merkle root must be set and matched first");   

        // Calculate the value of the leaf node
        bytes memory idxBytes = abi.encodePacked(uint32(idx));
        bytes32 hash = keccak256(abi.encodePacked(bytes1(0x00), idxBytes, blockData));

        // Recalculate the root value 
        uint pos = idx;
        for (uint i = 0; i < proof.length; i++) {
            if (pos % 2 == 0) {
                hash = keccak256(abi.encodePacked(bytes1(0x01), hash, proof[i]));
            } else {
                hash = keccak256(abi.encodePacked(bytes1(0x01), proof[i], hash));
            }
            pos /= 2;
        }
        if (hash == hr) {
            disputeResult = true;
            sendViaCall(server);
        } else {
            disputeResult = false;
            sendViaCall(client);    
        }
    }

    function sendViaCall(address payable _to) internal {
        (bool sent,) = _to.call{value: 0.01 ether}("");
        require(sent, "Failed to send Ether");
    }

    function getBalance() public view returns (uint) {
        return address(this).balance;
    }
    function getGs() public view returns (bytes memory) {
        return gs;
    }
    function getN() public view returns (bytes memory) {
        return N;
    }
    function getChIndex() public view returns (uint[][] memory) {
        return indexs;
    }

    function getHr() public view returns (bytes32) {
        return hr;
    }
    
    function getDisputeResult() public view returns (bool) {
        return disputeResult;
    }
    
}
