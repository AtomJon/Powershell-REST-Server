package resource

type ScriptExecutionRequest struct {
	Path       string
	Parameters map[string]string
}