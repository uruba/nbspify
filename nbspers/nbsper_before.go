package nbspers

type NbsperBefore struct {
}

func (nbsperBefore *NbsperBefore) GetCode() string {
	return "before"
}

func (nbsperBefore *NbsperBefore) Apply(input string, matchSegments []string) string {
	return GetNbsperProcessFunc(
		input,
		matchSegments,
		` (%s)\b`,
		"&nbsp;$1")()
}
