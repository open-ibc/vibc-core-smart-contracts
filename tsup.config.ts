import { defineConfig } from "tsup";

export default defineConfig({
  entry: [
    "src/index.ts",
    "src/utils/index.ts",
    "src/evm/index.ts",
    "src/utils/cli.ts",
    "src/utils/io.ts",
    "src/utils/constants.ts",
    "src/evm/chain.ts",
    "src/evm/contracts/*.ts",
    "src/evm/schemas/*.ts",
    "src/scripts/update-contracts-script.ts",
    "src/scripts/verify-contract-script.ts",
    "src/scripts/fork-deployment-test.ts",
    "src/scripts/deploy-multisig.ts",
    "src/scripts/execute-multisig-tx.ts",
  ],
  format: ["cjs", "esm"], // Build for commonJS and ESmodules
  dts: true, // Generate declaration file (.d.ts)
  splitting: false,
  sourcemap: true,
  clean: true,
  outDir: "./dist",
});
