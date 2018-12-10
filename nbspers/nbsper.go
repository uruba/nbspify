package nbspers

type Nbsper interface {
	GetCode() string
	Apply(input string, matchSegments []string) string
}
