import { parseZodSchema } from "../utils/io";
import { Registry } from "../utils/registry";
import { z } from "zod";

const ChainConfigSchema = z.object({
  chainName: z.string().min(1),
  chainId: z.number().int().min(1),
  vmType: z.enum(["evm", "polymer", "cosmos"]).optional().default("evm"),
  description: z.optional(z.string()),
  rpc: z.string().min(1),
});

export const chainRegistrySchema = z.array(ChainConfigSchema).min(1);
export type ChainRegistry = Registry<z.infer<typeof ChainConfigSchema>>;
export type Chain = ChainRegistry["Element"];

// load chain registry from a config object
export function loadChainRegistry(config: z.input<typeof chainRegistrySchema>) {
  const chainRegistry = parseZodSchema(
    "ChainRegistry",
    config,
    chainRegistrySchema.parse
  );
  return new Registry(chainRegistry, { nameFunc: (c) => c.chainName });
}
