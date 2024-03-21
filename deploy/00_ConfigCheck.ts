import { HardhatRuntimeEnvironment } from "hardhat/types";
import { DeployFunction } from "hardhat-deploy/types";
import "dotenv/config";

const func: DeployFunction = async function (hre: HardhatRuntimeEnvironment) {
  if (hre.network.name !== "hardhat") {
    if (!process.env.PRIVATE_KEY) {
      throw Error("private key not found! - set PRIVATE_KEY var in .env file");
    }

    // Check that:
    // All constructor args are set

    // Proof Verifier:
    // 1.) Light client type must be either DUMMY or OP-CLIENT
    // 2.) L2OO set

    // LightClient:
    // 1.) Light client type must be either DUMMY or OP-CLIENT
    // 2.) Fraud proof window &l1blockprovider

    // Dispatcher:
    //1.) init port prefix
    // eh I don't think we need to verify light client either should not be dependencies or

    // const accounts = await hre.getUnnamedAccounts();
    // const accounts0 = await hre.getNamedAccounts();
    // const accounts1 = await hre.ethers.getSigners();
    // console.log("acc", accounts, accounts1, accounts0);
    // console.log(`non-local network with from address: ${accounts[0]}`);
  } else {
    console.log("hardhat");
  }
};

func.tags = ["ConfigCheck"];
export default func;
