
https://www.adithyan.io/blog/agent-engineering-101

* Prompts: for reusable slash commands and task instructions
* Roles: for choosing which model handles each task
* Config / agent files: for tying prompts and roles together

Typical places:

* config.yaml — main Continue configuration
* .continue/mcpServers/ — MCP server configs
* prompt markdown files — reusable custom prompts
* agent/config files in Mission Control — reusable custom workflows


```
your-workspace/
├─ .continue/
│  └─ mcpServers/
│     └─ continue-docs-mcp.yaml
├─ prompts/
│  └─ explain-invokable.md
└─ config.yaml
```