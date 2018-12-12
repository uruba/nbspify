package nbspers

import "nbspify/config"

type NbsperAround struct {
}

func (nbsperAround *NbsperAround) GetCode() string {
	return "around"
}

func (nbsperAround *NbsperAround) Apply(input string, matchSegments []string, optionsNbsper *config.OptionsNbsper) string {
	processFunc := GetNbsperProcessFunc(
		matchSegments,
		optionsNbsper)

	output := ""
	output = processFunc(input, " (%s) ", "&nbsp;$1&nbsp;")
	output = processFunc(output, ` (%s)\b`, "&nbsp;$1")
	output = processFunc(output, `\b(%s) `, "$1&nbsp;")

	return output
}
