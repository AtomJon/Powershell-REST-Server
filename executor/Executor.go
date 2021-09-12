package executor

import (
	"log"

	"github.com/AtomJon/Subscript-REST-Server/resource"
)

func ExecuteResource(resource resource.Resource) (string, error) {

	log.Printf("Executing resource: %v", resource)

	switch resource.ResourceExtension {

	case ".ps1":
		return ExecutePowershell(resource)

	default:
		return "Cannot specify resource type", nil
		
	}

}