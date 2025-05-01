package main

import (
	// "encoding/json"
	"fmt"
	// "io"
	"log"
	// "net/http"
	// "time"

	expr "github.com/expr-lang/expr"
	mcp_golang "github.com/metoro-io/mcp-golang"
	"github.com/metoro-io/mcp-golang/transport/stdio"
)

type Eval struct {
	expression string `json:"expression" jsonschema:"required,description=The expression to be evaluated"`
	env map[string]interface{} `json:"description" jsonschema:"description=Optional parameters map"`
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
    
    // Compile the expression received from arguments
    program, err := expr.Compile(arguments.expression, expr.Env(Env{}) )
    if err != nil {
        log.Fatalf("Error compiling expression: %v", err)
    }

    // Execute the compiled program
    env := Env{}
    output, err := expr.Run(program, env)
    if err != nil {
        return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Error executing expression: %v", err))), err
    }

    // Return the result as a formatted string
    resultText := fmt.Sprintf("%s = %v", arguments.expression, output)
    return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(resultText)), nil
  })
	if err != nil {
		log.Fatalf("Error registering hello tool: %v", err)
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