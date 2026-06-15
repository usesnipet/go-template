import { defineConfig } from "@kubb/core";
import { pluginClient } from "@kubb/plugin-client";
import { pluginOas } from "@kubb/plugin-oas";
import { pluginReactQuery } from "@kubb/plugin-react-query";
import { pluginTs } from "@kubb/plugin-ts";
import { pluginZod } from "@kubb/plugin-zod";

import type { ResolveNameParams } from "@kubb/core";
function removeControllerSuffix(name: string) {
  return name.replace(/([-_]?controller[-_]?)/gi, "");
}

const nameTransformer = {
  name: (name: ResolveNameParams["name"]) => removeControllerSuffix(name),
};

export default defineConfig({
  root: ".",
  input: {
    path: "../docs/swagger.json",
  },
  output: {
    path: "./src/__generated__/api",
    clean: true,
  },
  plugins: [
    pluginOas(),
    pluginTs({
      output: {
        path: "types",
      },
      transformers: nameTransformer,
    }),
    pluginClient({
      output: {
        path: "client",
      },
      importPath: "@/lib/api-client",
      transformers: nameTransformer,
    }),
    pluginReactQuery({
      output: {
        path: "hooks",
      },
      client: {
        importPath: "@/lib/api-client",
      },
      transformers: nameTransformer,
    }),
    pluginZod({
      output: {
        path: "zod",
      },
      importPath: "zod",
      version: "4",
      // typed: true,
      dateType: "string",
      unknownType: "unknown",
      transformers: nameTransformer,
    }),
  ],
});
