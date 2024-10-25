import { z } from "zod";
import { Registry } from "../../utils/registry";
import { parseZodSchema } from "../../utils/io";

// A contract may or may not be deployed (null address).
export const ContractItemSchema = z
  .object({
    name: z.string().min(1),
    description: z.optional(z.string()),
    factoryName: z.optional(z.string()),
    deployArgs: z.optional(z.array(z.any())),
    libraries: z.optional(
      z.array(
        z.object({
          name: z.string().min(1),
          address: z.string().min(1),
        })
      )
    ),
    // either a account name from account registry, or a private key or a mnemonic signer
    deployer: z.string().nullish(),
    address: z.string().nullish(),
    init: z.optional(
      z.object({
        signature: z.string().min(1),
        args: z.array(z.string().min(1)),
      })
    ),
    abi: z.optional(z.any()),
    solcVersion: z.optional(z.string()),
  })
  .strict();

const ContractItemList = z.array(ContractItemSchema);
const registryName = "contracts";

const MultiChainContractRegistrySchema = z.array(
  z.object({ chainName: z.string().min(1), [registryName]: ContractItemList })
);

export type ContractItem = z.infer<typeof ContractItemSchema>;

// export type ContractRegistry = Registry<ContractItem>
export type MultiChainContractRegistry = Registry<{
  chainName: string;
  [registryName]: ContractRegistry;
}>;

export class ContractRegistry extends Registry<ContractItem> {}

export class ContractRegistryLoader {
  static loadSingle(config: any): ContractRegistry {
    return loadContractRegistry(config);
  }

  static loadMultiple(config: any): MultiChainContractRegistry {
    return loadMultiChainContractRegistry(config);
  }

  static emptyMultiple(): MultiChainContractRegistry {
    //  @ts-ignore
    return new Registry<{ chainName: string; contracts: ContractRegistry }>(
      [],
      {
        //  @ts-ignore
        toObj: (c) => {
          return {
            chainName: c.chainName,
            [registryName]: c.contracts.serialize(),
          };
        },
      }
    );
  }

  static newMultiple(
    items: { chainName: string; contracts: ContractRegistry }[]
  ): MultiChainContractRegistry {
    // @ts-ignore
    return new Registry(items, {
      nameFunc: (c) => c.chainName,
      toObj: (c) => {
        return {
          chainName: c.chainName,
          [registryName]: c.contracts.serialize(),
        };
      },
    });
  }

  static emptySingle(): ContractRegistry {
    return new ContractRegistry([], { nameInParent: registryName });
  }
}

function loadContractRegistry(config: any): ContractRegistry {
  const parsed = parseZodSchema(
    "ContractRegistry",
    config,
    ContractItemList.parse
  );
  return new ContractRegistry(parsed, {
    nameFunc: (c) => c.name,
    nameInParent: registryName,
  });
}

function loadMultiChainContractRegistry(
  config: any
): MultiChainContractRegistry {
  const parsed = parseZodSchema(
    "MultiChainContractRegistry",
    config,
    MultiChainContractRegistrySchema.parse
  );
  const contractRegistries = parsed.map((item) => ({
    chainName: item.chainName,
    contracts: loadContractRegistry(item.contracts),
  }));
  return ContractRegistryLoader.newMultiple(contractRegistries);
}
