import { defineConfig } from "tinacms";
import { new_docFields } from "./templates";

// Your hosting provider likely exposes this as an environment variable
const branch = process.env.HEAD || process.env.VERCEL_GIT_COMMIT_REF || "main";

export default defineConfig({
  branch,
  clientId: "a7ed6e7a-14da-4fc4-9bf4-6bafff7d5887", // Get this from tina.io
  token: "4e7e558580baf15d1e76127cfbc885f758ff593d", // Get this from tina.io
  client: { skip: true },
  build: {
    outputFolder: "admin",
    publicFolder: "static",
  },
  media: {
    tina: {
      mediaRoot: "",
      publicFolder: "static",
    },
  },
  schema: {
    collections: [
      {
        format: "md",
        label: "Docs",
        name: "docs",
        path: "content",
        frontmatterFormat: "yaml",
        match: {
          include: "**/*",
          exclude: "_index",
        },
        fields: [
          {
            type: "rich-text",
            name: "body",
            label: "Body of Document",
            description: "This is the markdown body",
            isBody: true,
            templates: [
              {
                name: 'vimeo',
                label: 'vimeo',
                match: {
                  start: '{{<',
                  end: '>}}',
                },
                fields: [
                  {
                    name: '_value',
                    label: 'value',
                    type: 'string',
                    required: true,
                  },
                ],
              },
            ],
          },
          ...new_docFields(),
        ],
      },
    ],
  },
});
