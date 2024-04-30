import { defineConfig } from "tsup";

export default defineConfig({
  entry: [
    "src/index.ts",
    "src/utils/index.ts",
    "src/evm/index.ts",
    "src/utils/cli.ts",
    "src/utils/io.ts",
    "src/utils/constants.ts",
    "src/evm/schemas/contract.ts",
    "src/evm/schemas/tx.ts",
    "src/evm/chain.ts",
    "src/evm/contracts/*.ts",
    "src/evm/account.ts",
    "src/scripts/deploy-script.ts",
    "src/scripts/upgrade-script.ts",
    "src/scripts/setup-dispatcher-script.ts",
  ],
  format: ["cjs", "esm"], // Build for commonJS and ESmodules
  dts: true, // Generate declaration file (.d.ts)
  splitting: false,
  sourcemap: true,
  clean: true,
  outDir: "./dist",
});
