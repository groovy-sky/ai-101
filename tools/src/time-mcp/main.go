package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type TimeArgs struct {
	Format string `json:"format"`
	UTC    bool   `json:"utc"`
}

func main() {
	log.SetOutput(os.Stderr)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	log.Println("Starting Time MCP Server...")

	server := mcp.NewServer(&mcp.Implementation{
		Name:    "time-mcp",
		Version: "1.0.0",
	}, nil)

	timeTool := mcp.Tool{
		Name:        "time_now",
		Description: "Returns the current system time. Optional: provide a Go time format string and set utc=true for UTC time.",
		InputSchema: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"format": map[string]interface{}{
					"type":        "string",
					"description": "Optional Go time format (default: RFC3339)",
				},
				"utc": map[string]interface{}{
					"type":        "boolean",
					"description": "Return time in UTC instead of local time",
				},
			},
		},
	}

	unixTool := mcp.Tool{
		Name:        "time_unix",
		Description: "Returns the current Unix timestamp (seconds since 1970-01-01 UTC).",
		InputSchema: map[string]interface{}{
			"type":       "object",
			"properties": map[string]interface{}{},
		},
	}

	mcp.AddTool(server, &timeTool, getSystemTime)
	mcp.AddTool(server, &unixTool, getUnixTime)

	log.Println("Tool registered successfully")
	log.Println("Server ready - waiting for MCP requests on stdin")

	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}

func getSystemTime(ctx context.Context, req *mcp.CallToolRequest, args TimeArgs) (*mcp.CallToolResult, any, error) {
	now := time.Now()

	if args.UTC {
		now = now.UTC()
	}

	format := time.RFC3339
	if args.Format != "" {
		format = args.Format
	}

	result := now.Format(format)

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: result},
		},
	}, nil, nil
}

func getUnixTime(ctx context.Context, req *mcp.CallToolRequest, args struct{}) (*mcp.CallToolResult, any, error) {
	now := time.Now().Unix()

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: fmt.Sprintf("%d", now)},
		},
	}, nil, nil
}
