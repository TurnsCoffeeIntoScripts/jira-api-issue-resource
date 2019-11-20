package application

import (
	"errors"
	"flag"
	"github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/configuration"
)

type JiraApiResourceApp struct {
	params configuration.JiraAPIResourceParameters
}

func Run() error {
	app := &JiraApiResourceApp{}
	if err := initFlagsAndParameters(app); err != nil {
		return err
	}

	if err := configurationReady(app); err != nil {
		return err
	}

	return app.executeFromContext()
}

func initFlagsAndParameters(app *JiraApiResourceApp) error {
	app.params = configuration.JiraAPIResourceParameters{}
	app.params.Parse()
	if !app.params.Meta.AllMandatoryValuesPresent() {
		flag.Usage()
		return errors.New("missing mandatory flags/parameters")
	}

	return nil
}

func configurationReady(app *JiraApiResourceApp) error {
	if !app.params.Meta.Ready() {
		return errors.New("flags and parameters did not form a valid set")
	}

	return nil
}
