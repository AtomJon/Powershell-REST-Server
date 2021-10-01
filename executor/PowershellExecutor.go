package executor

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/AtomJon/subscriptrestserver/resource"
)

func ExecutePowershell(resource resource.Resource, request resource.ScriptExecutionRequest) (string, error) {
	parsedRequest := parseRequestAsString(request)
	
	containedContent := fmt.Sprintf("& {%s}", resource.Content)

	cmd := exec.Command("powershell.exe", "-Command", containedContent, parsedRequest);

	result, err := cmd.CombinedOutput();
	if (err != nil) {
		return "", fmt.Errorf("%s %v", string(result), err);
	}

	return string(result), nil;
}

func parseRequestAsString(request resource.ScriptExecutionRequest) (string) {
	builder := strings.Builder{}

	for identifier, value := range(request.Parameters) {
		builder.WriteRune('-')
		builder.WriteString(identifier)
		builder.WriteString(" \"")
		builder.WriteString(value)
		builder.WriteString("\" ")
	}

	return builder.String()
}