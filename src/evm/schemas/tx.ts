import { z } from "zod";
import { Registry } from "../../utils/registry";
import { parseZodSchema } from "../../utils/io";

const TxItemSchema = z.object({
  name: z.string().min(1),
  description: z.optional(z.string()),
  // either a account name from account registry, or a private key or a mnemonic signer
  deployer: z.string().nullish(),
  signature: z.string().min(1),
  address: z.string().nullish(),
  factoryName: z.optional(z.string()),
  args: z.optional(z.array(z.any())),
  init: z.optional(
    z.object({
      signature: z.string().min(1),
      args: z.array(z.string().min(1)),
    })
  ),
});

type TxItem = z.infer<typeof TxItemSchema>;
const TxItemList = z.array(TxItemSchema);

export class TxRegistry extends Registry<TxItem> {}

export function loadTxRegistry(config: any): TxRegistry {
  const parsed = parseZodSchema("TxRegistry", config, TxItemList.parse);
  return new TxRegistry(parsed, {
    nameFunc: (c) => c.name,
    nameInParent: "transactions",
  });
}
