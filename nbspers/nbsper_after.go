package nbspers

type NbsperAfter struct {
}

func (nbsperAfter *NbsperAfter) GetCode() string {
	return "after"
}

func (nbsperAfter *NbsperAfter) Apply(input string, matchSegments []string) string {
	return GetNbsperProcessFunc(
		input,
		matchSegments,
		"( |^)(%s) ",
		"$1$2&nbsp;")()
}
