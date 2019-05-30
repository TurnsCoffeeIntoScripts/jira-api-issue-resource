package main

import (
	"flag"
	"fmt"
	"github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/configuration"
	"github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/http/rest"
	"os"
)

func main() {
	flags := configuration.JiraApiResourceFlags{}
	ok := flags.SetupFlags(true)

	if !ok || *flags.ShowHelp {
		flag.Usage()
		os.Exit(0)
	} else {
		ok, err := rest.ApiCall(configuration.GetExecutionContext(flags))

		if !ok {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
