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
		`\b(%s) `,
		"$1&nbsp;")()
}
