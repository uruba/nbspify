package config

type OptionsNbsper struct {
	CaseSensitive bool
}

func CreateOptionsNbsper(caseSensitive bool) *OptionsNbsper {
	return &OptionsNbsper{CaseSensitive: caseSensitive}
}
