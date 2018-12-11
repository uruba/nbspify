package nbspers

type NbsperAround struct {
}

func (nbsperAround *NbsperAround) GetCode() string {
	return "around"
}

func (nbsperAround *NbsperAround) Apply(input string, matchSegments []string) string {
	return GetNbsperProcessFunc(
		input,
		matchSegments,
		" (%s) ",
		"&nbsp;$1&nbsp;")()
}
