package main

import (
	"fmt"
	"github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/action"
	"os"
)

func main() {
	if !action.Check() {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", "Jira Issue doesn't exist. An error might have occured during the API call")
		os.Exit(1)
	}
}
