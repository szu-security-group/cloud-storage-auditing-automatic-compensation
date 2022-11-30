# Blockchain enabled cloud storage auditing for automatic user compensation

With the large-scale application of cloud storage services, the security of stored data is valued by individuals and companies.
To enhance outsourced data security, cloud storage auditing protocols to verify the integrity of the data without downloading all the data.
To alleviate the need of a trusted third party in public auditing, recent work proposed using blockchain to support public cloud storage auditing.

## What we do
![model.png](model.png)

To sole the usability issue, we propose an efficient smart-contract-based cloud storage auditing protocol.
The protocol only requires one signature verification that is natively supported by Ethereum, which makes the proposed protocol very efficient.
The proposed protocol also supports automatic user compensation.

## Usage

Our project code is written in golang, make sure you have version 1.19 of golang on your computer. You can download the code by git clone. Run the command line as shown below.

### struct
Let's briefly introduce the structure of the project. 
contract Contains the code files of the smart contract. core includes each algorithm flow of the protocol. data is the method implementation of protocol processing data.
main.go is a complete demo of the protocol.

### deploy contract
In order to quickly deploy our contract, you can use the remix online compiler.
autopay.sol includes the main algorithm code of the protocol, and the other is the code to call the precompiled contract.
You can create a wallet through MetaMask, and then deploy smart contracts through remix to test the operation of the protocol and its performance.

