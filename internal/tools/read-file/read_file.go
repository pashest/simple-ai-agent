package read_file

import (
	"encoding/json"
	"os"
	"simple_agent/internal/tools"
)

type ReadFileInput struct {
	Path string `json:"path" jsonschema_description:"The relative path of a file in the working directory."`
}

func NewReadFileTool() tools.ToolDefinition {
	return tools.ToolDefinition{
		Name:        "read_file",
		Description: "Read the contents of a given relative file path. Use this when you want to see what's inside a file. Do not use this with directory names.",
		InputSchema: tools.GenerateSchema[ReadFileInput](),
		Function:    ReadFile,
	}
}

func ReadFile(input json.RawMessage) (string, error) {
	readFileInput := ReadFileInput{}
	err := json.Unmarshal(input, &readFileInput)
	if err != nil {
		panic(err)
	}

	content, err := os.ReadFile(readFileInput.Path)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
