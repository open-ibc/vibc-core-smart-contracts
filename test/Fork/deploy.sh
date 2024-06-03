
export MODULE_ROOT_PATH="./"
anvil --fork-url ${OP_SEPOLIA_RPC_URL}
&
npx deploy-vibc-core-smart-contracts --ACCOUNT_SPECS_PATH="test/Fork/deployment-accounts-spec.yaml" --DEPLOY_SPECS_PATH="test/Fork/contract-spec.yaml"
