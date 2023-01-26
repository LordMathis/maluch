package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"

	"github.com/LordMathis/maluch/agents"
)

type PipelineConfig struct {
	version int
	stages  []struct {
		name   string
		agent  agents.BaseAgent
		script []string
	}
}

func Parse(filename string) PipelineConfig {
	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	var config PipelineConfig

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}

	return config

}
