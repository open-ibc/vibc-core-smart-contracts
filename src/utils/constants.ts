import path from "path";

export const PACKAGE_VERSION = require("../../package.json").version;

// Defaults
export const DEFAULT_DEPLOYER = "default";
export const DEFAULT_RPC_URL = "http://127.0.0.1:8545";
export const DEFAULT_CHAIN_ID = "31337";
export const DEFAULT_CHAIN_NAME = "local";
export const DEFAULT_VERIFIER = "blockscout";
export const MODULE_ROOT_PATH =
  process.env.MODULE_ROOT_PATH || path.resolve(__dirname, "..", ".."); // Note: this is impacted by the tsup config since that can potential change where this file ends up being built to
export const PARENT_MODULE_ROOT_PATH = path.dirname(path.resolve("package.json")); // Path of module using this package; path.resolve uses cwd so this should typically be the root of the module that's running the bin
const DEFAULT_ARTIFACTS_PATH = path.resolve(MODULE_ROOT_PATH, "out");
const DEFAULT_DEPLOYMENTS_PATH = path.resolve(
  PARENT_MODULE_ROOT_PATH,
  "deployments"
);
const DEFAULT_SPECS_PATH = path.resolve(MODULE_ROOT_PATH, "./specs");
const DEFAULT_DEPLOYMENT_ENVIRONMENT = "local"; // Deployment Environment disambiguates between deployment environments on the same chains (e.g. staging testnet vs prod testnet) This should be one of the following: local, staging, prod, mainnet

export const CHAIN_NAME = process.env.CHAIN_NAME || DEFAULT_CHAIN_NAME;
export const CHAIN_ID = parseInt(process.env.CHAIN_ID || DEFAULT_CHAIN_ID);
export const RPC_URL = process.env.RPC_URL || DEFAULT_RPC_URL;
export const ANVIL_PORT = process.env.ANVIL_PORT || "8545";
export const DEPLOYMENT_ENVIRONMENT =
  process.env.DEPLOYMENT_ENVIRONMENT || DEFAULT_DEPLOYMENT_ENVIRONMENT;

// The path where we access artifacts for already deployed contracts
export const ARTIFACTS_PATH = process.env.ARTIFACTS_PATH
  ? process.env.ARTIFACTS_PATH
  : DEFAULT_ARTIFACTS_PATH; // Used for importing both

// The path where we save deployments
export const DEPLOYMENTS_PATH = process.env.DEPLOYMENTS_PATH
  ? process.env.DEPLOYMENTS_PATH
  : DEFAULT_DEPLOYMENTS_PATH;

const SPECS_BASE_PATH = process.env.SPECS_BASE_PATH
  ? process.env.SPECS_BASE_PATH
  : DEFAULT_SPECS_PATH;

export const UPDATE_SPECS_PATH = process.env.UPDATE_SPECS_PATH
  ? process.env.UPDATE_SPECS_PATH
  : path.resolve(SPECS_BASE_PATH, "update.spec.yaml");

export const ACCOUNTS_SPECS_PATH = process.env.ACCOUNTS_SPECS_PATH
  ? process.env.ACCOUNTS_SPECS_PATH
  : path.resolve(SPECS_BASE_PATH, "evm.accounts.yaml");

export const EXTRA_BINDINGS_PATH = process.env.EXTRA_BINDINGS_PATH;

// The path where we access external artifacts for already deployed contracts
export const EXTRA_ARTIFACTS_PATH = process.env.EXTRA_ARTIFACTS_PATH;

