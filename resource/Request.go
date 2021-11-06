package resource

type ScriptExecutionRequest struct {
	ScriptName string
	Parameters map[string]string
}