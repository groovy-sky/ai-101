# AI 101 – Getting Started

## Introduction
This guide provides a simple introduction of Artificial Intelligence and a hands‑on setup of deploying, configuring and using LLM AI.

## Overview

Artificial Intelligence (AI) refers to computer systems designed to perform tasks that normally require human intelligence. These tasks include understanding language, recognizing images, generating content, making predictions, and assisting with decision making.

Modern AI systems are typically powered by **machine learning models**, especially **large language models (LLMs)** trained on large datasets.

### Common capabilities

- Natural language understanding
- Text generation
- Image recognition
- Code generation
- Data analysis
- Conversational assistants

### Most Popular models

Some widely used AI platforms and tools include:

- **ChatGPT (OpenAI)**
- **Claude (Anthropic)**
- **Google Gemini**
- **GitHub Copilot**
- **Ollama**
- **Azure AI Foundry**

## Prerequirements

To be able to follow this document, you will need **Active Azure subscription** and Admin rights on your PC. 


## Setup Deployment

Environment setup order:
1. Install Ollama
2. Install Visual Studio Code
3. Create Azure AI Foundry and Deploy a Model in it
4. Install Continue Extension in VS Code
5. Configure Ollama agent
6. Configure GPT agent
7. Result

### Install Ollama

Ollama allows running AI models locally on your machine.

1. Visit:
   
   https://ollama.com

2. Download the installer for your operating system.

3. Install Ollama.

4. Verify installation by running:

```
ollama --version
```

5. Pull a model to test:

```
ollama run llama3
```

If the model responds in the terminal, the installation works.

---

### Install Visual Studio Code

1. Download VS Code from:

https://code.visualstudio.com

2. Install it using the default options.

3. Launch VS Code.

Recommended extensions:

- GitHub Copilot (optional)
- Continue (installed later)

---

### 4. Create AI Foundry and Deploy a Model

Now we will deploy an AI model in Azure.

## Create AI Foundry

1. Go to the **Azure Portal**
2. Search for **Azure AI Foundry**
3. Create a new resource
4. Select your:
   - Subscription
   - Resource Group
   - Region

After deployment finishes, open the AI Foundry project.

---

#### Deploy a Model

1. Open **Model Catalog**
2. Choose a model (example: GPT or similar)
3. Click **Deploy**
4. Wait until deployment completes.

---

#### Test the Model in Playground

1. Open **Playground**
2. Select your deployed model
3. Enter a test prompt such as:

```
Explain what artificial intelligence is.
```

4. Verify the model generates a response.

---


You will need the following information for later configuration:

- **Endpoint URL**
- **Model name / deployment name**
- **API key**

---

### 5. Install Continue Extension in VS Code

Continue is an AI coding assistant extension.

## Install the extension

1. Open **VS Code**
2. Go to **Extensions**
3. Search for:

```
Continue
```

4. Install the extension.

---

### Configure Ollama in Continue

Edit the Continue configuration file.

Example configuration:

```json
{
  "models": [
    {
      "title": "Local Llama",
      "provider": "ollama",
      "model": "llama3"
    }
  ]
}
```

This connects Continue to the locally running Ollama model.

---

### 6. Configure OpenAI Agent

Now configure Continue to use your Azure OpenAI deployment.

Example configuration:

```json
{
  "models": [
    {
      "title": "Azure OpenAI",
      "provider": "openai",
      "apiBase": "https://your-resource.openai.azure.com/",
      "apiKey": "YOUR_API_KEY",
      "model": "YOUR_MODEL_NAME"
    }
  ]
}
```

Replace:

- `apiBase` with your endpoint
- `apiKey` with your Azure key
- `model` with your deployed model name

Save the configuration.

---

### 7. 

Now verify everything works.

1. Open a project folder in VS Code.
2. Open the **Continue panel**.
3. Ask a question such as:

