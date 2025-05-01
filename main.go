package main

import (
	"fmt"
	"log"

	expr "github.com/expr-lang/expr"
	mcp_golang "github.com/metoro-io/mcp-golang"
	"github.com/metoro-io/mcp-golang/transport/stdio"
)

type Eval struct {
	Expression string                 `json:"expression" jsonschema:"required,description=The expression to be evaluated"`
	Env        map[string]interface{} `json:"env" jsonschema:"description=Optional parameters map"`
}

type Content struct {
	Title       string  `json:"title" jsonschema:"required,description=The title to submit"`
	Description *string `json:"description" jsonschema:"description=The description to submit"`
}

type Env struct {
}

// main initializes and starts the MCP server, registers tools, prompts, and resources, and handles incoming requests.
func main() {
	log.Println("Starting MCP Server...")

	server := mcp_golang.NewServer(stdio.NewStdioServerTransport())

	// Register tools, prompts, and resources here...
	err := server.RegisterTool("eval", "Evaluate an expression", func(arguments Eval) (*mcp_golang.ToolResponse, error) {
		log.Println("Received request for eval tool")
	
		// Use the expression from arguments, or default to a simple expression if empty
		code := arguments.Expression
		if code == "" {
			code = "10 + 72 / 2"
		}

		program, err := expr.Compile(code, expr.Env(Env{}))
		if err != nil {
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Error compiling expression: %v", err))), err
		}

		// Create environment with any provided variables
		env := Env{}
		output, err := expr.Run(program, env)
		if err != nil {
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Error executing expression: %v", err))), err
		}

		return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("%s = %v", code, output))), nil
	})
	if err != nil {
		log.Fatalf("Error registering eval tool: %v", err)
	}

	// Register "prompt_test" prompt
	err = server.RegisterPrompt("prompt_test", "This is a test prompt", func(arguments Content) (*mcp_golang.PromptResponse, error) {
		log.Println("Received request for prompt_test")
		return mcp_golang.NewPromptResponse("description", mcp_golang.NewPromptMessage(mcp_golang.NewTextContent(fmt.Sprintf("Hello, %s!", arguments.Title)), mcp_golang.RoleUser)), nil
	})
	if err != nil {
		log.Fatalf("Error registering prompt_test: %v", err)
	}

	// Register test resource
	err = server.RegisterResource("test://resource", "resource_test", "This is a test resource", "application/json",
		func() (*mcp_golang.ResourceResponse, error) {
			log.Println("Received request for resource: test://resource")
			return mcp_golang.NewResourceResponse(mcp_golang.NewTextEmbeddedResource(
				"test://resource", "This is a test resource", "application/json",
			)), nil
		})
	if err != nil {
		log.Fatalf("Error registering resource: %v", err)
	} else {
		log.Println("Successfully registered resource: test://resource") // Debug log
	}
	// Start the server
	log.Println("MCP Server is now running and waiting for requests...")
	err = server.Serve()
	if err != nil {
		log.Fatalf("Server error: %v", err)
	}

	select {} // Keeps the server running
}