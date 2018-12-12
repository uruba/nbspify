package nbspers

import (
	"fmt"
	"nbspify/config"
	"regexp"
	"strings"
)

type Nbsper interface {
	GetCode() string
	Apply(input string, matchSegments []string, optionsNbsper *config.OptionsNbsper) string
}

func GetNbsperRegexp(matchSegments []string, patternMatch string, optionsNbsper *config.OptionsNbsper) *regexp.Regexp {
	patterns := strings.Join(matchSegments, "|")

	patternComplete := fmt.Sprintf(patternMatch, patterns)

	if !optionsNbsper.CaseSensitive {
		patternComplete = "(?i)" + patternComplete
	}

	return regexp.MustCompile(patternComplete)
}

func GetNbsperProcessFunc(
	matchSegments []string,
	optionsNbsper *config.OptionsNbsper) func(input string, patternMatch string, patternReplace string) string {
	return func(input string, patternMatch string, patternReplace string) string {
		expression := GetNbsperRegexp(matchSegments, patternMatch, optionsNbsper)

		return expression.ReplaceAllString(input, patternReplace)
	}
}
