# AI 101: AI Agents (part 3)

![AI 101 course logo](../logo.png)

Continue is an AI coding assistant available in the VS Code extension. It supports several modes for different workflows and can be customized with rules and prompts.

## Modes in VS Code

### Agent mode
Use Agent mode for multi-step development tasks that involve tools, file changes, or external integrations.

### Chat mode
Use Chat mode to ask questions, understand code, and get conversational help.

### Plan mode
Use Plan mode to review intended actions before execution, especially for prompts that may change code or data.

## Rules in VS Code

Rules provide instructions to the model for Agent, Chat, and Edit mode requests.

### What rules do
Rules shape model behavior with project-specific or organization-specific guidance.

Examples:
- Use a specific library
- Follow project conventions
- Generate migrations after schema changes

### What rules do not affect
Rules are **not included** in:
- Autocomplete
- Apply

### Local rules
Local rules live in:

```text
.continue/rules

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