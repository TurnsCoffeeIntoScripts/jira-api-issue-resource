module github.com/TurnsCoffeeIntoScripts/jira-api-resource

go 1.12

replace (
	github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/action v0.0.0 => ./pkg/action
	github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/configuration v0.0.0 => ./pkg/configuration
	github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/domain v0.0.0 => ./pkg/domain
	github.com/TurnsCoffeeIntoScripts/jira-api-resource/pkg/http/rest v0.0.0 => ./pkg/http/rest
)
