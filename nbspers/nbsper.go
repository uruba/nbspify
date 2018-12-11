package nbspers

import (
	"fmt"
	"regexp"
	"strings"
)

type Nbsper interface {
	GetCode() string
	Apply(input string, matchSegments []string) string
}

func GetNbsperRegexp(matchSegments []string, patternMatch string) *regexp.Regexp {
	patterns := strings.Join(matchSegments, "|")

	return regexp.MustCompile(fmt.Sprintf(patternMatch, patterns))
}

func GetNbsperProcessFunc(
	input string,
	matchSegments []string,
	patternMatch string,
	patternReplace string) func() string {
	return func() string {
		expression := GetNbsperRegexp(matchSegments, patternMatch)

		return expression.ReplaceAllString(input, patternReplace)
	}
}
