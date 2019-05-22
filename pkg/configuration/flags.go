package configuration

import (
	"flag"
	"strings"
)

type JiraApiResourceFlags struct {
	JiraApiUrl *string
	Protocol   *string
	Username   *string
	Password   *string
	IssueId    *string
	RawData    *string
	Data       map[string]string
}

func (f *JiraApiResourceFlags) SetupFlags(parse bool) {
	f.JiraApiUrl = flag.String("url", "", "The base URL of the Jira Rest API to be used (without the http|https)")
	f.Protocol = flag.String("protocol", "https", "The http protocol to be used (http|https)")
	f.Username = flag.String("username", "", "Username used to establish a secure connection with the Jira Rest API")
	f.Password = flag.String("password", "", "Password used by the username in the connection to the Jira Rest API")
	f.IssueId = flag.String("id", "", "The Jira ticket ID (Format: <PROJECT_KEY>-<NUMBER>")
	f.RawData = flag.String("data", "", "Map form: \\'key1-val1_key2_val2\\'")

	if parse {
		f.Parse()
	}
}

func (f *JiraApiResourceFlags) Parse() {
	flag.Parse()
}

func (f *JiraApiResourceFlags) Validate() bool {
	if *f.IssueId == "" || *f.JiraApiUrl == "" || *f.Username == "" || *f.Password == "" {
		flag.Usage()
		return false
	}

	return true
}

func (f *JiraApiResourceFlags) PopulateMap() {
	f.Data = make(map[string]string)
	for _, val := range strings.Split(*f.RawData, "_") {
		innerData := strings.Split(val, "-")
		f.Data[innerData[0]] = innerData[1]
	}
}
