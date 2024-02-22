## Deployment
1. Set the env vars in .env file
```bash
export IMPL_SALT=$(openssl rand -hex 32)
export PRIVATE_KEY=
export ETHERSCAN_API_KEY=
export L2_OUTPUT_ORACLE_ADDRESS=
export OP_RPC_URL=
export BASE_RPC_URL=
```
2. Run source command to set the env vars
```bash
source .env
```
3. Deploy the contracts with    
```bash
 CHAIN=<chain> forge script script/Deploy.s.sol --rpc-url $OP_RPC_URL  --broadcast --private-key $PRIVATE_KEY   --verifier-url https://< chain >-sepolia.blockscout.com/api --verifier blockscout --verify
```
<chain> can be `optimism` or `base`.

### Deploying a single contract
All the functions responsible for deploying a single contract are public, allowing the utilization of the --sig argument in the forge script to specifically direct the deployment of a single contract.
For example, to deploy a dummy proof verifier contract, the following command can be used:
```bash
forge script script/Deploy.s.sol --rpc-url $OP_RPC_URL  --broadcast --private-key $PRIVATE_KEY --sig 'deployDummyProofVerifier()'
```

### Troubleshooting
In the event of encountering an ambiguous error containing `EvmError: Revert` it is probable that adjusting the IMPL_SALT environment variable is necessary. 
This variable plays a crucial role in determining the addresses of smart contracts deployed through CREATE2. 
Deploying the same contracts with the identical IMPL_SALT value twice will result in a failure during the second deployment.
To generate a new IMPL_SALT value you can run `source .env`.