package executor

import (
	"github.com/AtomJon/subscriptrestserver/resource"
)

func ExecuteResource(resource resource.Resource) (string, error) {

	switch resource.ResourceExtension {

	case ".ps1":
		return ExecutePowershell(resource)

	default:
		return "Cannot specify resource type", nil
		
	}

}