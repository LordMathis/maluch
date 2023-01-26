package agents

type Action string

type BaseAgent struct {
	name   string
	action map[Action]interface{}
}

type AgentRunner interface {
	RunAction(string) error
}
