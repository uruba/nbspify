package nbspers

import (
	"fmt"
	"regexp"
	"strings"
)

type NbsperBefore struct {
}

func (nbsperBefore *NbsperBefore) GetCode() string {
	return "before"
}

func (nbsperBefore *NbsperBefore) Apply(input string, matchSegments []string) string {
	words := strings.Join(matchSegments, "|")
	expression := regexp.MustCompile(fmt.Sprintf(" (%s)( |$)", words))

	return expression.ReplaceAllString(input, "&nbsp;$1$2")
}
