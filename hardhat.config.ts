import { HardhatUserConfig } from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox";
import "hardhat-contract-sizer";
import "@nomicfoundation/hardhat-foundry";

const config: HardhatUserConfig = {
  solidity: "0.8.15",
  // paths: {
  // sources: "./contracts",
  //   tests: "./test",
  //   cache: "./cache",
  // artifacts: "./artifacts",
  // },
};

export default config;
