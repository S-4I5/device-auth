package scope

import (
	"fmt"
	"slices"
)

type Scope int

const (
	Undefined Scope = iota - 1
	AuthV1Validation
	AuthV1Authenticate
	UserV1GetUserById
	UserV1GetUserByPhoneNumber
)

var (
	strings = []string{
		"AuthV1Validation",
		"AuthV1Authenticate",
		"UserV1GetUserById",
		"UserV1GetUserByPhoneNumber",
	}
	errCannotParseStringToScope = fmt.Errorf("cannot pasre string to scope")
)

func ParseScope(input string) (Scope, error) {
	intScope := slices.Index(strings, input)
	if intScope == -1 {
		return Undefined, errCannotParseStringToScope
	}
	return Scope(intScope), nil
}

func (s Scope) toValue() int {
	return int(s)
}

func (s Scope) String() string {
	if 0 < int(s) || int(s) >= len(strings) {
		return "unknown"
	}
	return strings[s]
}
