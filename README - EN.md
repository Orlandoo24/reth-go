### Go Script README (Updated Version)
#### Project Overview
This script provides a script that automatically mines reth, and the golang implementation has better performance than the js implementation.

#### Functionality Summary
The script is responsible for the following tasks:
1. **Ethereum Node Connection**: Establishes a connection to an Ethereum node using the `ethclient.Dial` function.
2. **Private Key Parsing**: Converts a hexadecimal private key into an ECDSA private key object.
3. **Address Derivation**: Derives the Ethereum address from the ECDSA public key.
4. **Challenge Hash Calculation**: Computes the hash of a given challenge, which is "rETH" in this case.
5. **Solution Finding**: Attempts to find a solution that meets a specific condition by generating random values and hashing them with the challenge hash.
6. **Transaction Construction**: Builds a transaction with specified parameters like nonce, value, gas limit, and gas price.
7. **Gas Price Management**: Waits and retries if the suggested gas price is higher than a predefined maximum.
8. **Transaction Signing and Sending**: Signs the transaction with the private key and sends it to the Ethereum network.

#### Usage Instructions
To use the Go script, follow these steps:
1. **Environment Setup**: Ensure that you have the Go programming language installed on your system.
2. **Private Key and Address Setup**: Replace the placeholder `"地址"` with your Ethereum account address and `"私钥"` with your actual private key in hexadecimal format.
3. **Connection String**: Update the Ethereum node RPC URL if necessary. The current URL is set to "https://rpc.flashbots.net".
4. **Run the Script**: Execute the script using the Go command: `go run <filename>.go`.

#### Contribution Guidelines
- Fork the repository and create a new branch for your feature (`Feat_xxx`).
- Make your changes and commit them with a descriptive commit message.
- Open a Pull Request to merge your feature branch back into the main project.

#### To-Do List
- [x] Fix the URL string format in `ethclient.Dial` to ensure a correct connection to the Ethereum node.
- [x] Implement an exit condition in the `findSolution` function to prevent potential infinite loops.
- [x] Verify that the fields in the `jsonData` structure comply with the ABI requirements of the smart contract.
- [x] Test the generation logic of `dataHex` to ensure the correctness of the constructed transaction data.
- [x] Implement error handling and retry mechanisms when transaction sending fails.
- [x] Add logging functionality to monitor the script's operation status and facilitate debugging.
- [x] Adjust the gas price and gas limit settings in the script according to real-time Ethereum network gas prices.
- [x] Develop unit tests to verify the correctness of the `findSolution` function and other critical functions.
- [x] Implement a confirmation mechanism after transaction sending to ensure the transaction has been successfully included in the blockchain.
- [x] Implement the distribution of airdrops or whitelist spots.
- [x] Implement the signing mechanism for distribution fees.
