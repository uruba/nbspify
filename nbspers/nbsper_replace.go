package nbspers

type NbsperReplace struct {
}

func (nbsperReplace *NbsperReplace) GetCode() string {
	return "replace"
}

func (nbsperReplace *NbsperReplace) Apply(input string, matchSegments []string) string {
	return GetNbsperProcessFunc(
		input,
		matchSegments,
		"%s",
		"&nbsp;")()
}
