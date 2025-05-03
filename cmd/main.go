package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	agentpkg "simple_agent/internal/agent"
	toolspkg "simple_agent/internal/tools"
	edit_file "simple_agent/internal/tools/edit-file"
	list_files "simple_agent/internal/tools/list-files"
	read_file "simple_agent/internal/tools/read-file"

	"github.com/anthropics/anthropic-sdk-go"
)

func main() {
	client := anthropic.NewClient()

	scanner := bufio.NewScanner(os.Stdin)
	getUserMessage := func() (string, bool) {
		if !scanner.Scan() {
			return "", false
		}
		return scanner.Text(), true
	}

	tools := []toolspkg.ToolDefinition{
		read_file.NewReadFileTool(),
		list_files.NewListFilesTool(),
		edit_file.NewEditFileTool(),
	}
	agent := agentpkg.NewAgent(&client, getUserMessage, tools)
	err := agent.Run(context.TODO())
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}
}
