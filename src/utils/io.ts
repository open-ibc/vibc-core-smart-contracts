import { fileURLToPath } from "url";
import fs from "fs";
import fsAsync from "fs/promises";
import path from "path";
import yaml from "yaml";
import { z } from "zod";
import nunjucks from "nunjucks";
import assert from "assert";
import { ProcessPromise } from "zx";
import { Chain, ChainConfigSchema } from "../evm/chain";
import {
  DEPLOYMENTS_PATH,
  ARTIFACTS_PATH,
  ACCOUNTS_SPECS_PATH,
  CHAIN_NAME,
  CHAIN_ID,
  RPC_URL,
  DEPLOY_SPECS_PATH,
  PACKAGE_VERSION,
  DEPLOYMENT_ENVIRONMENT,
  ANVIL_PORT,
  UPGRADE_SPECS_PATH,
} from "./constants";
import yargs from "yargs/yargs";
import { hideBin } from "yargs/helpers";
import { AccountRegistry } from "../evm/account";
import { ethers } from "ethers";

export interface StringToStringMap {
  [key: string]: string | null | undefined;
}

export type DeployedContractObject = {
  factory: string;
  address: string;
  abi: any;
  bytecode: string;
  args: any[];
  metadata?: string;
  name: string;
};

// readYamlFile reads a yaml file and returns the parsed object.
export function readYamlFile(file: string): any {
  return yaml.parse(fs.readFileSync(file, "utf-8"));
}

export function contractNameToDeployFile(
  contractName: string,
  chainId: number
) {
  return `${contractName}-${chainId}.json`;
}

/**
 * Load a json or yaml file and return the parsed object.
 * @param file file path must have an extension of .json or .yaml|.yml
 * @returns parsed object
 * @throws error if file not found, file extension is not supported, or parsing error
 */
export function parseObjFromFile(
  file: string,
  options: BufferEncoding = "utf8"
): any {
  if (!fs.existsSync(file)) {
    throw new Error(`file not found: ${file}`);
  }
  const ext = path.parse(file).ext;
  switch (ext) {
    case ".json":
      return JSON.parse(fs.readFileSync(file, options));
    case ".yaml":
    case ".yml":
      return yaml.parse(fs.readFileSync(file, options));
    default:
      throw new Error(
        `unsupported file extension: ${ext}, only {.json, .yaml|.yml} are supported`
      );
  }
}

// configure the renderer to throw an error if a template variable is not found
export const renderEnv = nunjucks.configure({ throwOnUndefined: true });

/**
 * Renders a template string using the provided environment variables.
 * @param str - The string to render. `{{var1}}`/`{{obj1.name}}` will be replaced with the value of `env.var1`/`obj.name`.
 * @param env - The environment variables (key/value pairs) to use for rendering.
 * @returns The rendered string.
 */
export const renderString = (str: string, env: any) => {
  try {
    return renderEnv.renderString(str, env);
  } catch (err: any) {
    throw new Error(
      `failed to render string: ${JSON.stringify(str)}, error: ${err}`
    );
  }
};

/**
 * Provides a way to parse a zod schema and throw a meaningful error message if parsing fails.
 * @param className A meaningful name for the class being parsed. Used for error messages.
 * @param config Config object to parse
 * @param parseFunc Normally the `parse` method of a zod schema.
 * @returns Parsed object that matches the zod schema.
 */
export function parseZodSchema<T>(
  className: string,
  config: any,
  parseFunc: (config: any) => T
) {
  try {
    return parseFunc(config);
  } catch (e) {
    const zErr: z.ZodError = e as any;
    if (e instanceof z.ZodError) {
      throw new Error(
        `parsing ${className} failed. ${zErr.issues
          .map((i) => i.path)
          .join(", ")}: ${zErr.message}\nconfig obj:\n${JSON.stringify(
          config,
          null,
          2
        )}`
      );
    } else {
      throw e;
    }
  }
}

/**
 * create directories recursively if not exists, similar to `mkdir -p dirPath`
 */

/**
 * Remove a dir and its contents recursively, similar to `rm -rf dirPath`
 * Then recreate the dir, similar to `mkdir -p dirPath`
 * This effectively clears all existing content the dir if they exist; otherwise just create the dir.
 */
export function resetDir(dirPath: string) {
  fs.rmSync(dirPath, { recursive: true, force: true });
  fs.mkdirSync(dirPath, { recursive: true });
}

export function setStdoutStderr(
  proc: ProcessPromise,
  wd: string,
  stdoutName: string = "stdout",
  stderrName: string = "stderr"
): ProcessPromise {
  proc.pipe(fs.createWriteStream(path.resolve(wd, stdoutName)));
  proc.stderr.pipe(fs.createWriteStream(path.resolve(wd, stderrName)));
  proc.quiet();
  return proc;
}

export function toEnvVarName(e: string) {
  return e.replaceAll("-", "_").toUpperCase();
}

/** Reads a deployment metadata rom a foundry build file. used to generate bytecode for deployment files*/
async function readMetadata(factoryName: string) {
  const filePath = path.join(
    ARTIFACTS_PATH,
    `${factoryName}.sol`,
    `${factoryName}.json`
  );

  try {
    const data = await fsAsync.readFile(filePath, "utf8");
    return JSON.stringify(JSON.parse(data).metadata);
  } catch (e) {
    console.error(`error reading from file ${filePath}: \n`, e);
  }
}

// Given a chain object, return the folder in which deployments for this chain will be
export const getDeployFolderForChain = (chain: Chain) => {
  return path.join(
    DEPLOYMENTS_PATH,
    chain.chainId.toString(),
    chain.deploymentEnvironment
  );
};

export async function writeDeployedContractToFile(
  chain: Chain,
  deployedContract: DeployedContractObject
) {
  const deployFileName = contractNameToDeployFile(
    deployedContract.name,
    chain.chainId
  );

  const deploymentFolder = getDeployFolderForChain(chain);
  const fullPath = path.join(deploymentFolder, deployFileName);
  ensureDir(deploymentFolder);

  // get metadata from contract from forge build output
  const metadata = await readMetadata(deployedContract.factory);
  const outData = JSON.stringify({
    ...deployedContract,
    metadata,
    version: PACKAGE_VERSION,
  });

  fs.writeFile(fullPath, outData, (err) => {
    if (err) {
      console.error(err);
      return;
    }
  });
}

// Read existing deployment files into env, so that we can use them in the deployment scripts
export async function readDeploymentFilesIntoEnv(env: any, chain: Chain) {
  const deploymentFolder = getDeployFolderForChain(chain);
  await ensureDir(deploymentFolder); // In case there are no existing files written yet to directory
  let files: any[] = [];
  try {
    files = await fsAsync.readdir(deploymentFolder);
  } catch (e) {
    console.log(`no files to read from`);
    return env;
  }
  for (const file of files) {
    if (file.endsWith(".json")) {
      try {
        const data = JSON.parse(
          fs.readFileSync(path.join(deploymentFolder, file), "utf8")
        );
        env[data.name] = data.address;
      } catch (e) {
        console.error(`error reading file ${file}`, e);
      }
    }
  }
  return env;
}

export async function readFromDeploymentFile(
  deploymentName: string,
  chain: Chain
) {
  const fileFolder = getDeployFolderForChain(chain);

  const filePath = path.join(
    fileFolder,
    contractNameToDeployFile(deploymentName, chain.chainId)
  );
  try {
    const data = JSON.parse(fs.readFileSync(filePath, "utf8"));
    return data;
  } catch (e) {
    console.error(`error reading file ${filePath}`, e);
  }
}

const compileInitArgs = (
  init: { args: any[]; signature: string } | undefined,
  env: StringToStringMap
) => {
  if (!init) {
    return "";
  }
  const initArgs = init.args.map((arg: any) => {
    return typeof arg === "string" ? renderString(arg, env) : arg;
  });
  const iFace = new ethers.Interface([`function ${init.signature}`]);
  return iFace.encodeFunctionData(init.signature, initArgs);
};

/**
 *
 * @param args Render the args for the contract deployment through looking them up in environment
 * @param init replace initArgs with the init string
 * @param env to look up the args in
 * @returns
 */
export const renderArgs = (
  args: any[] | undefined,
  init: { args: any[]; signature: string } | undefined,
  env: any
) => {
  const initData = compileInitArgs(init, env);

  return args
    ? args.map((arg: any) => {
        if (typeof arg !== "string") return arg;
        if (arg === "$INITARGS") {
          if (initData === "")
            throw new Error(`Found $INITARGS but no args to replace it with.`);
          return initData;
        }
        return renderString(arg, env);
      })
    : [];
};

/**
 * create directories recursively if not exists, similar to `mkdir -p dirPath`
 */
export function ensureDir(dirPath: string) {
  if (fs.existsSync(dirPath)) {
    assert(fs.statSync(dirPath).isDirectory(), `${dirPath} is not a directory`);
  } else {
    fs.mkdirSync(dirPath, { recursive: true });
  }
}

export async function parseArgsFromCLI() {
  // Load args from command line. CLI args take priority. Then env vars. Then default values if none are specified
  const argv1 = await yargs(hideBin(process.argv)).argv;

  // write a method that uses argv1 if specified, otherwsie use the imported files from utils/constants
  const chainName = argv1.CHAIN_NAME || CHAIN_NAME;
  const chainId = argv1.CHAIN_ID || CHAIN_ID;
  const rpcUrl = argv1.RPC_URL || RPC_URL;
  const deploymentEnvironment =
    argv1.DEPLOYMENT_ENVIRONMENT || DEPLOYMENT_ENVIRONMENT;
  const accountSpecs =
    (argv1.ACCOUNT_SPECS_PATH as string) || ACCOUNTS_SPECS_PATH;

  const deploySpecs = (argv1.DEPLOY_SPECS_PATH as string) || DEPLOY_SPECS_PATH;
  const upgradeSpecs =
    (argv1.UPGRADE_SPECS_PATH as string) || UPGRADE_SPECS_PATH;
  const anvilPort = (argv1.ANVIL_PORT as string) || ANVIL_PORT;

  const chainParse = ChainConfigSchema.safeParse({
    rpc: rpcUrl,
    chainId,
    chainName,
    vmType: "evm",
    description: "local chain",
    deploymentEnvironment,
  });
  if (!chainParse.success) {
    throw new Error(`failed to parse chain config: ${chainParse.error.errors}`);
  }

  const accountConfigFromYaml = {
    name: "local",
    registry: parseObjFromFile(accountSpecs),
  };

  const accounts = AccountRegistry.loadMultiple([accountConfigFromYaml]);

  return {
    chain: chainParse.data,
    accounts,
    accountSpecs,
    upgradeSpecs,
    deploySpecs,
    args: argv1,
    anvilPort,
  };
}
