package nbspers

import (
	"fmt"
	"regexp"
	"strings"
)

type NbsperAfter struct {
}

func (nbsperAfter *NbsperAfter) GetCode() string {
	return "after"
}

func (nbsperAfter *NbsperAfter) Apply(input string, matchSegments []string) string {
	words := strings.Join(matchSegments, "|")
	expression := regexp.MustCompile(fmt.Sprintf("( |^)(%s) ", words))

	return expression.ReplaceAllString(input, "$1$2&nbsp;")
}
