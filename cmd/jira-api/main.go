package main

import (
	"flag"
	"fmt"
	"github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/configuration"
	"github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/http/rest"
	"os"
)

func main() {
	conf := configuration.JiraAPIResourceConfiguration{}
	ok := conf.SetupFlags()

	if *conf.Flags.ContextFlags.ShowHelp.Value {
		flag.Usage()
		os.Exit(0)
	} else if !ok {
		for i, err := range conf.Errors {
			fmt.Printf("Error [%d] --> %v\n", i, err)
		}
		flag.Usage()
		os.Exit(-1)
	} else {
		ctx := configuration.GetExecutionContext(conf)
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
