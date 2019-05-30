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
		ctx := configuration.GetExecutionContext(flags)
		if ctx != nil {
			ok, err := rest.ApiCall(*ctx)

			if !ok {
				fmt.Println(err)
				os.Exit(1)
			}
		} else {
			fmt.Println("Unable to get execution context")
			os.Exit(1)
		}
	}
}
