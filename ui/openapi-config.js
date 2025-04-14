//import{ ConfigFile } from '@rtk-query/codegen-openapi'

/**@type {import('@rtk-query/codegen-openapi').ConfigFile} */
const config = {
  schemaFile: '../docs/swagger.json',
  apiFile: './src/store/api/baseApi.ts',
  apiImport: 'baseApi',
  outputFile: './src/store/api/index.ts',
  exportName: 'api',
  hooks: true,
}

export default config