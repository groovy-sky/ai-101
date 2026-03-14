package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type PowerShellArgs struct {
	Command     string            `json:"command"`
	Script      string            `json:"script"`
	ScriptBlock string            `json:"scriptBlock"`
	Arguments   []string          `json:"arguments"`
	WorkingDir  string            `json:"workingDir"`
	Timeout     int               `json:"timeout"`
	Env         map[string]string `json:"env"`
}

type GetCommandArgs struct {
	Name    string `json:"name"`
	Module  string `json:"module"`
	Pattern string `json:"pattern"`
}

type ModuleArgs struct {
	Name string `json:"name"`
}

var (
	powershellPath string
	defaultTimeout = 30
)

func main() {
	log.SetOutput(os.Stderr)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	log.Println("Starting PowerShell MCP Server...")

	if err := detectPowerShell(); err != nil {
		log.Fatalf("PowerShell detection failed: %v", err)
	}

	log.Printf("Using PowerShell: %s", powershellPath)

	server := mcp.NewServer(&mcp.Implementation{
		Name:    "powershell-mcp",
		Version: "1.0.0",
	}, nil)

	// Tool 1: Execute PowerShell Command (FIXED)
	executeTool := mcp.Tool{
		Name:        "powershell_execute",
		Description: "Execute PowerShell commands or scripts. Provide ONE of: command (simple command), script (path to .ps1 file), or scriptBlock (multi-line script). Returns output, errors, and exit code.",
		InputSchema: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"command": map[string]interface{}{
					"type":        "string",
					"description": "PowerShell command to execute (e.g., 'Get-Process | Select-Object -First 5')",
				},
				"script": map[string]interface{}{
					"type":        "string",
					"description": "Path to PowerShell script file (.ps1)",
				},
				"scriptBlock": map[string]interface{}{
					"type":        "string",
					"description": "Multi-line PowerShell script block to execute",
				},
				"arguments": map[string]interface{}{
					"type":        "array",
					"description": "Arguments to pass to the script or command",
					"items": map[string]interface{}{
						"type": "string",
					},
				},
				"workingDir": map[string]interface{}{
					"type":        "string",
					"description": "Working directory for command execution",
				},
				"timeout": map[string]interface{}{
					"type":        "integer",
					"description": "Timeout in seconds (default: 30, max: 300)",
					"default":     30,
					"minimum":     1,
					"maximum":     300,
				},
				"env": map[string]interface{}{
					"type":        "object",
					"description": "Environment variables to set (key-value pairs)",
					"additionalProperties": map[string]interface{}{
						"type": "string",
					},
				},
			},
		},
	}

	// Tool 2: Get PowerShell Command Info
	getCommandTool := mcp.Tool{
		Name:        "powershell_get_command",
		Description: "Get information about PowerShell commands, including syntax, parameters, and examples. Similar to Get-Command and Get-Help.",
		InputSchema: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"name": map[string]interface{}{
					"type":        "string",
					"description": "Command name to get information about (e.g., 'Get-Process')",
				},
				"module": map[string]interface{}{
					"type":        "string",
					"description": "Module name to filter commands",
				},
				"pattern": map[string]interface{}{
					"type":        "string",
					"description": "Pattern to search for commands (wildcard supported, e.g., 'Get-*')",
				},
			},
		},
	}

	// Tool 3: List PowerShell Modules
	listModulesTool := mcp.Tool{
		Name:        "powershell_list_modules",
		Description: "List all available PowerShell modules installed on the system.",
		InputSchema: map[string]interface{}{
			"type":       "object",
			"properties": map[string]interface{}{},
		},
	}

	// Tool 4: Get Module Info
	getModuleTool := mcp.Tool{
		Name:        "powershell_get_module",
		Description: "Get detailed information about a specific PowerShell module, including version, description, and exported commands.",
		InputSchema: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"name": map[string]interface{}{
					"type":        "string",
					"description": "Module name (e.g., 'Microsoft.PowerShell.Management')",
				},
			},
			"required": []string{"name"},
		},
	}

	mcp.AddTool(server, &executeTool, executePowerShell)
	mcp.AddTool(server, &getCommandTool, getCommandInfo)
	mcp.AddTool(server, &listModulesTool, listModules)
	mcp.AddTool(server, &getModuleTool, getModuleInfo)

	log.Println("Tools registered successfully")
	log.Println("Server ready - waiting for MCP requests on stdin")

	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}

func detectPowerShell() error {
	// Try PowerShell Core (pwsh) first
	if path, err := exec.LookPath("pwsh"); err == nil {
		powershellPath = path
		return nil
	}

	// Try PowerShell Core with .exe extension (Windows)
	if path, err := exec.LookPath("pwsh.exe"); err == nil {
		powershellPath = path
		return nil
	}

	// Fall back to Windows PowerShell
	if runtime.GOOS == "windows" {
		if path, err := exec.LookPath("powershell.exe"); err == nil {
			powershellPath = path
			return nil
		}
	}

	return fmt.Errorf("PowerShell not found. Install PowerShell Core (pwsh) or Windows PowerShell")
}

func executePowerShell(ctx context.Context, req *mcp.CallToolRequest, args PowerShellArgs) (*mcp.CallToolResult, any, error) {
	startTime := time.Now()
	log.Printf("PowerShell execution request")

	// Validate
	if err := validateExecuteArgs(args); err != nil {
		log.Printf("Validation error: %v", err)
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: fmt.Sprintf("Validation error: %v", err)},
			},
			IsError: true,
		}, nil, nil
	}

	// Set timeout
	timeout := time.Duration(args.Timeout) * time.Second
	if args.Timeout == 0 {
		timeout = time.Duration(defaultTimeout) * time.Second
	}

	execCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	// Build PowerShell command
	psArgs := []string{
		"-NoProfile",
		"-NonInteractive",
		"-ExecutionPolicy", "Bypass",
		"-OutputFormat", "Text",
	}

	var commandToRun string

	if args.Command != "" {
		psArgs = append(psArgs, "-Command")
		commandToRun = args.Command
		if len(args.Arguments) > 0 {
			commandToRun += " " + strings.Join(args.Arguments, " ")
		}
		psArgs = append(psArgs, commandToRun)
		log.Printf("Executing command: %s", args.Command)
	} else if args.Script != "" {
		psArgs = append(psArgs, "-File", args.Script)
		psArgs = append(psArgs, args.Arguments...)
		log.Printf("Executing script: %s", args.Script)
	} else if args.ScriptBlock != "" {
		psArgs = append(psArgs, "-Command")
		psArgs = append(psArgs, args.ScriptBlock)
		log.Printf("Executing script block (%d chars)", len(args.ScriptBlock))
	}

	// Create command
	cmd := exec.CommandContext(execCtx, powershellPath, psArgs...)

	// Set working directory
	if args.WorkingDir != "" {
		cmd.Dir = args.WorkingDir
	}

	// Set environment variables
	if len(args.Env) > 0 {
		cmd.Env = os.Environ()
		for key, value := range args.Env {
			cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", key, value))
		}
	}

	// Execute and capture output
	output, err := cmd.CombinedOutput()
	duration := time.Since(startTime)

	var result strings.Builder
	result.WriteString("# PowerShell Execution Result\n\n")
	result.WriteString(fmt.Sprintf("**Duration:** %v\n", duration))

	exitCode := 0
	if cmd.ProcessState != nil {
		exitCode = cmd.ProcessState.ExitCode()
	}
	result.WriteString(fmt.Sprintf("**Exit Code:** %d\n\n", exitCode))

	if err != nil {
		if execCtx.Err() == context.DeadlineExceeded {
			log.Printf("Command timed out after %v", timeout)
			result.WriteString(fmt.Sprintf("**Status:** ⚠️ Timeout (exceeded %v)\n\n", timeout))
		} else {
			log.Printf("Command failed: %v", err)
			result.WriteString(fmt.Sprintf("**Status:** ❌ Failed\n"))
			result.WriteString(fmt.Sprintf("**Error:** %v\n\n", err))
		}
	} else {
		result.WriteString("**Status:** ✅ Success\n\n")
	}

	if len(output) > 0 {
		result.WriteString("## Output\n\n")
		result.WriteString("```\n")
		result.WriteString(string(output))
		result.WriteString("\n```\n")
	} else {
		result.WriteString("(No output)\n")
	}

	log.Printf("Execution completed in %v, exit code: %d", duration, exitCode)

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: result.String()},
		},
	}, nil, nil
}

func getCommandInfo(ctx context.Context, req *mcp.CallToolRequest, args GetCommandArgs) (*mcp.CallToolResult, any, error) {
	log.Printf("Get command info: %s", args.Name)

	var psCommand string
	if args.Name != "" {
		psCommand = fmt.Sprintf("Get-Command -Name '%s' -ErrorAction SilentlyContinue | Select-Object Name, CommandType, Module, Version, Source | ConvertTo-Json; Get-Help '%s' -ErrorAction SilentlyContinue",
			escapeSingleQuote(args.Name), escapeSingleQuote(args.Name))
	} else if args.Module != "" {
		psCommand = fmt.Sprintf("Get-Command -Module '%s' | Select-Object Name, CommandType | ConvertTo-Json",
			escapeSingleQuote(args.Module))
	} else if args.Pattern != "" {
		psCommand = fmt.Sprintf("Get-Command -Name '%s' | Select-Object Name, CommandType, Module | ConvertTo-Json",
			escapeSingleQuote(args.Pattern))
	} else {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: "Please provide name, module, or pattern"},
			},
			IsError: true,
		}, nil, nil
	}

	cmd := exec.CommandContext(ctx, powershellPath, "-NoProfile", "-NonInteractive", "-Command", psCommand)
	output, err := cmd.CombinedOutput()

	if err != nil {
		log.Printf("Command failed: %v", err)
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: fmt.Sprintf("Failed to get command info: %v\n\nOutput:\n%s", err, string(output))},
			},
			IsError: true,
		}, nil, nil
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(output)},
		},
	}, nil, nil
}

func listModules(ctx context.Context, req *mcp.CallToolRequest, args struct{}) (*mcp.CallToolResult, any, error) {
	log.Printf("Listing PowerShell modules")

	psCommand := "Get-Module -ListAvailable | Select-Object Name, Version, Description | ConvertTo-Json"
	cmd := exec.CommandContext(ctx, powershellPath, "-NoProfile", "-NonInteractive", "-Command", psCommand)
	output, err := cmd.CombinedOutput()

	if err != nil {
		log.Printf("Command failed: %v", err)
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: fmt.Sprintf("Failed to list modules: %v", err)},
			},
			IsError: true,
		}, nil, nil
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(output)},
		},
	}, nil, nil
}

func getModuleInfo(ctx context.Context, req *mcp.CallToolRequest, args ModuleArgs) (*mcp.CallToolResult, any, error) {
	log.Printf("Get module info: %s", args.Name)

	psCommand := fmt.Sprintf(`
$module = Get-Module -Name '%s' -ListAvailable | Select-Object -First 1
if ($module) {
    $commands = Get-Command -Module '%s' | Select-Object Name, CommandType
    @{
        Name = $module.Name
        Version = $module.Version.ToString()
        Description = $module.Description
        Path = $module.Path
        ExportedCommands = $commands
    } | ConvertTo-Json -Depth 3
} else {
    Write-Output "Module not found: %s"
}
`, escapeSingleQuote(args.Name), escapeSingleQuote(args.Name), args.Name)

	cmd := exec.CommandContext(ctx, powershellPath, "-NoProfile", "-NonInteractive", "-Command", psCommand)
	output, err := cmd.CombinedOutput()

	if err != nil {
		log.Printf("Command failed: %v", err)
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: fmt.Sprintf("Failed to get module info: %v", err)},
			},
			IsError: true,
		}, nil, nil
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(output)},
		},
	}, nil, nil
}

func validateExecuteArgs(args PowerShellArgs) error {
	// Must have exactly one of: command, script, or scriptBlock
	count := 0
	if args.Command != "" {
		count++
	}
	if args.Script != "" {
		count++
	}
	if args.ScriptBlock != "" {
		count++
	}

	if count == 0 {
		return fmt.Errorf("must provide one of: command, script, or scriptBlock")
	}
	if count > 1 {
		return fmt.Errorf("provide only one of: command, script, or scriptBlock")
	}

	// Validate timeout
	if args.Timeout < 0 {
		return fmt.Errorf("timeout cannot be negative")
	}
	if args.Timeout > 300 {
		return fmt.Errorf("timeout cannot exceed 300 seconds")
	}

	// Validate script file exists
	if args.Script != "" {
		if _, err := os.Stat(args.Script); os.IsNotExist(err) {
			return fmt.Errorf("script file not found: %s", args.Script)
		}
	}

	return nil
}

func escapeSingleQuote(s string) string {
	return strings.ReplaceAll(s, "'", "''")
}
