package main

import (
	"flag"
	"github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/configuration"
	"os"
)

func main() {
	flags := configuration.JiraAPIResourceConfiguration{}
	ok := flags.SetupFlags()

	//if !ok || *flags.ShowHelp {
	if !ok {
		flag.Usage()
		os.Exit(0)
	} else {
		//ctx := configuration.GetExecutionContext(flags)
		/*if ctx != nil {
			ok, err := rest.ApiCall(*ctx)

			if !ok {
				fmt.Println(err)
				os.Exit(1)
			}
		} else {
			fmt.Println("Unable to get execution context")
			os.Exit(1)
		}*/
	}
}
