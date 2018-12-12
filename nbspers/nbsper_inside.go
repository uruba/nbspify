package nbspers

import (
	"nbspify/config"
	"regexp"
	"strings"
)

type NbsperInside struct {
}

func (nbsperInside *NbsperInside) GetCode() string {
	return "inside"
}

func (nbsperInside *NbsperInside) Apply(input string, matchSegments []string, optionsNbsper *config.OptionsNbsper) string {
	expression := GetNbsperRegexp(matchSegments, "(.|^)%s(.|$)", optionsNbsper)
	matches := expression.FindAllString(input, -1)

	if len(matches) > 0 {
		for _, match := range matches {
			replaceRegexp := regexp.MustCompile("([^^]) ([^$])")
			replacedMatch := replaceRegexp.ReplaceAllString(match, "$1&nbsp;$2")

			input = strings.Replace(input, match, replacedMatch, 1)
		}
	}

	return input
}
