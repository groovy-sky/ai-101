# Azure AI Search MCP Tool

This project implements a **Model Context Protocol (MCP) server** that allows AI assistants
(such as those in **Zed IDE**) to query an **Azure AI Search index**.

The server exposes a tool called `azure_ai_search` that executes search queries against
Azure AI Search and returns formatted results.

---

## Features

- Connects Zed AI assistants to **Azure AI Search**
- Implements an **MCP server using stdio transport**
- Allows querying any Azure Search index
- Returns formatted search results with document fields
- Configurable result size and selected fields

---

## Prerequisites

Before running this tool, make sure you have:

- Go installed (1.20 or newer recommended)
- An Azure account
- An **Azure AI Search service**
- At least one **search index**
- An **API key** for the search service
- Zed IDE (or another MCP-compatible client)

---

## Install Dependencies

Inside the project directory run:

go mod tidy

This installs the required dependency:

- github.com/metoro-io/mcp-golang

---

## Environment Variables

The server reads Azure credentials from environment variables.

Required variables:

AZURE_SEARCH_ENDPOINT  
Example:

https://your-search-service.search.windows.net

AZURE_SEARCH_API_KEY  
Your Azure AI Search **admin or query key**

Optional variable:

AZURE_SEARCH_API_VERSION  
Default used by the tool:

2024-07-01

Example setup (Linux/macOS):

export AZURE_SEARCH_ENDPOINT="https://mysearch.search.windows.net"
export AZURE_SEARCH_API_KEY="your_api_key_here"
export AZURE_SEARCH_API_VERSION="2024-07-01"

Example setup (Windows PowerShell):

$env:AZURE_SEARCH_ENDPOINT="https://mysearch.search.windows.net"
$env:AZURE_SEARCH_API_KEY="your_api_key_here"
$env:AZURE_SEARCH_API_VERSION="2024-07-01"

---

## Run the MCP Server

Start the server with:

go run main.go

The server communicates through **stdio**, which is required for MCP integration
with tools such as Zed IDE.

---

## Configure Zed IDE

To allow Zed to use this tool, register the MCP server in the Zed MCP configuration.

Example configuration:

{
  "mcpServers": {
    "azure-ai-search": {
      "command": "go",
      "args": ["run", "main.go"],
      "cwd": "/path/to/tools/mcp-az-ai-search"
    }
  }
}

After restarting Zed, the tool will become available to the AI assistant.

---

## Tool: azure_ai_search

The MCP server registers a tool named:

azure_ai_search

Description:

Execute a search query against an Azure AI Search index and return matching documents.

---

## Parameters

query  
Search query string.

indexName  
Name of the Azure AI Search index.

top (optional)  
Maximum number of results to return. Default: 10

select (optional)  
Comma-separated list of fields to return.

Example:

title,content,url

---

## Example Tool Call

Example request:

{
  "query": "machine learning",
  "indexName": "documents",
  "top": 5,
  "select": "title,content,url"
}

Example response:

Search completed. Found 5 results.

Result 1:
{
  "title": "Introduction to Machine Learning",
  "content": "...",
  "url": "https://example.com/ml"
}

---

## How It Works

1. The MCP server starts using stdio transport.
2. Zed connects to the server as an MCP tool provider.
3. The AI assistant calls the `azure_ai_search` tool.
4. The tool sends a REST request to Azure AI Search.
5. Results are formatted and returned to the assistant.

---

## Error Handling

The server will exit if required environment variables are missing.

Example:

Error: AZURE_SEARCH_ENDPOINT and AZURE_SEARCH_API_KEY environment variables must be set

The tool also returns errors if Azure Search returns a non‑200 HTTP status.
