

```
name: My Config
version: 0.0.1
schema: v1

models:
  - name: <MODEL_NAME>
    model: <MODEL_ID>
    provider: azure
    apiBase: https://just-an-example.openai.azure.com
    apiKey: <YOUR_AZURE_API_KEY>
    env:
      apiVersion: <API_VERSION>
      deployment: <MODEL_DEPLOYMENT>
      apiType: azure-openai
```
