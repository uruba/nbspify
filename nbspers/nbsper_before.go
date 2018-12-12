package nbspers

import "nbspify/config"

type NbsperBefore struct {
}

func (nbsperBefore *NbsperBefore) GetCode() string {
	return "before"
}

func (nbsperBefore *NbsperBefore) Apply(input string, matchSegments []string, optionsNbsper *config.OptionsNbsper) string {
	return GetNbsperProcessFunc(
		matchSegments,
		optionsNbsper)(input, ` (%s)\b`, "&nbsp;$1")
}
