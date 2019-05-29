package main

import (
	"flag"
	"github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/configuration"
	"os"
)

func main() {
	flags := configuration.JiraApiResourceFlags{}
	flags.SetupFlags(true)

	if *flags.ShowHelp {
		flag.Usage()
		os.Exit(0)
	} else {
		configuration.GetExecutionContext().Execute()
	}
}
