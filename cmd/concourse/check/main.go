package main

import (
	"fmt"
	"github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/action"
	"os"
)

func main() {
	if err := action.Check(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
