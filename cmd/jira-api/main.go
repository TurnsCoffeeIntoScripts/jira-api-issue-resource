package main

import (
	"fmt"
	"github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/configuration"
)

func main() {
	flags := configuration.JiraApiResourceFlags{}
	flags.SetupFlags(true)

	fmt.Println(*flags.Username)
}
