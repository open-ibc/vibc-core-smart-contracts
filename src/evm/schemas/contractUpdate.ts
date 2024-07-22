import { z } from "zod";
import { TxItemSchema } from "./tx";
import { ContractItemSchema } from "./contract";
import { Registry } from "../../utils/registry";

// Represents the updates we can make to a contract, either deploying a new one or sending a tx to an existing one.
export const UpdateContractConfigSchema = z.union([
  TxItemSchema,
  ContractItemSchema,
]);
type UpdateContractConfig = z.infer<typeof UpdateContractConfigSchema>;
export class UpdateContractRegistry extends Registry<UpdateContractConfig> {}

// Load a parsed config file into a contract update registry that can be passed into the updateContract method
export function loadContractUpdateRegistry(
  config: any
): UpdateContractRegistry {
  const parsed = config.map((c: any) => {
    const parsed = UpdateContractConfigSchema.safeParse(c);
    if (!parsed.success) {
      throw new Error(`Invalid contract update config: ${parsed.error.errors}`);
    }
    return parsed.data;
  });
  return new UpdateContractRegistry(parsed, {
    nameFunc: (c) => c.name,
    nameInParent: "updateContracts",
  });
}
