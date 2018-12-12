package nbspers

import "nbspify/config"

type NbsperReplace struct {
}

func (nbsperReplace *NbsperReplace) GetCode() string {
	return "replace"
}

func (nbsperReplace *NbsperReplace) Apply(input string, matchSegments []string, optionsNbsper *config.OptionsNbsper) string {
	return GetNbsperProcessFunc(
		matchSegments,
		optionsNbsper)(input, "%s", "&nbsp;")
}
