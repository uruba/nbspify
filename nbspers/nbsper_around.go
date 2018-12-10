package nbspers

import (
	"fmt"
	"regexp"
	"strings"
)

type NbsperAround struct {
}

func (nbsperAround *NbsperAround) GetCode() string {
	return "around"
}

func (nbsperAround *NbsperAround) Apply(input string, matchSegments []string) string {
	words := strings.Join(matchSegments, "|")
	expression := regexp.MustCompile(fmt.Sprintf(" (%s) ", words))

	return expression.ReplaceAllString(input, "&nbsp;$1&nbsp;")
}
