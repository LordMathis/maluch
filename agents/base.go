package agents

import (
	"context"
	"fmt"
	"reflect"
)

type Action string

type BaseAgent struct {
	name   string
	action map[Action]interface{}
}

type AgentRunner interface {
	RunAction(string, context.Context) error
}

func (b BaseAgent) RunAction(a Action, ctx context.Context) error {
	f := reflect.ValueOf(b.action[a])

	in := []reflect.Value{reflect.ValueOf(ctx)}

	res := f.Call(in)
	result := res[0].Interface()

	fmt.Println(result)
	// TODO:
	return nil
}
