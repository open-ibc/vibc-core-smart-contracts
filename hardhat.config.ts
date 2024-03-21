import { HardhatUserConfig } from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox";
import "hardhat-contract-sizer";
import "@nomicfoundation/hardhat-foundry";
import "hardhat-deploy";

const config: HardhatUserConfig = {
  solidity: {
    compilers: [
      {
        version: "0.8.15",
        settings: {
          optimizer: {
            enabled: true,
            runs: 200,
          },
        },
      },
    ],
  },
  networks: {
    hardhat: {
      allowUnlimitedContractSize: true, // We need to set this to true since hardhat doesn't correctly calculate contract
    },
    optimismSepolia: {
      url: "localhost",
      chainId: 11155420,
      allowUnlimitedContractSize: true,
      url: "" process.env.OP_SEPOLIA_RPC_URL
        ? process.env.OP_SEPOLIA_RPC_URL
        : "https://sepolia.optimism.io",
      accounts: process.env.PRIVATE_KEY ? [process.env.PRIVATE_KEY] : "remote",
    },
  },
  namedAccounts: {
    deployer: {
      default: process.env.private_KEY ? process.env.private_KEY : 0,
    },
  },
};

export default config;
