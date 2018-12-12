package nbspers

import "nbspify/config"

type NbsperAfter struct {
}

func (nbsperAfter *NbsperAfter) GetCode() string {
	return "after"
}

func (nbsperAfter *NbsperAfter) Apply(input string, matchSegments []string, optionsNbsper *config.OptionsNbsper) string {
	return GetNbsperProcessFunc(
		matchSegments,
		optionsNbsper)(input, `\b(%s) `, "$1&nbsp;")
}
