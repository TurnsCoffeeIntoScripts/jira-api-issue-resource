package configuration

type Context int

const (
	ReadIssue Context = iota
	EditCustomField
	Unknown
)

var names = [...]string{"ReadIssue", "EditCustomField", "Unknown"}

func (c Context) String() string {

	if c < ReadIssue || c >= Unknown {
		return "Unknown"
	}

	return names[c]

}

func GetContext(contextString string) Context {
	for index, name := range names {
		if name == contextString {
			return Context(index)
		}
	}

	return Unknown
}
