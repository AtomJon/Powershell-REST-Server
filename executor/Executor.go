package executor

import (
	"github.com/AtomJon/subscriptrestserver/resource"
)

func ExecuteResource(resource resource.Resource, request resource.ScriptExecutionRequest) (string, error) {

	switch resource.ResourceExtension {

	case ".ps1":
		return ExecutePowershell(resource, request)

	default:
		return "Cannot specify resource type", nil
		
	}

}