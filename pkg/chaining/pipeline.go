package chaining

import (
	"github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/configuration"
	"github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/log"
	"github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/service"
)

type Pipeline struct {
	steps    []Step
	length   int
	csValues CrossStepsValues
}

func (p *Pipeline) BuildPipelineFromChain(chain []service.Service, params *configuration.JiraAPIResourceParameters) error {
	for s := range chain {
		p.addStep(chain[s], params)
	}

	p.length = len(p.steps)

	return nil
}

func (p *Pipeline) addStep(s service.Service, params *configuration.JiraAPIResourceParameters) {
	newStep := Step{}
	newStep.Service = s
	newStep.Name = s.Name()
	newStep.params = params

	p.steps = append(p.steps, newStep)
}

func (p *Pipeline) Execute() error {
	for index := range p.steps {
		log.Logger.Debug("Executing step (", index+1, "/", p.length, ")", p.steps[index].Name)
		err := p.steps[index].Execute(p.csValues)

		if err != nil {
			return err
		}

		if index < p.length-1 {
			ns := &p.steps[index+1]
			p.csValues = p.steps[index].PrepareNextStep(ns, p.csValues)
		}
	}

	return nil
}
