package main

import (
	"fmt"
	"nbspify/config"
	"nbspify/input"
	nbsperKinds "nbspify/nbspers"
)

func main() {
	processedText := input.ReadInput()

	nbspers := getNbspers()
	definitions := config.LoadDefinitions()

	optionsNbsper := initOptions()

	if nbspers != nil {
		for _, nbsper := range nbspers {
			matchSegments := definitions[nbsper.GetCode()]

			if len(matchSegments) > 0 {
				processedText = nbsper.Apply(processedText, matchSegments, optionsNbsper)
			}
		}
	}

	// STDOUT
	fmt.Print(processedText)
}

func initOptions() *config.OptionsNbsper {
	// TODO - configurable by flags
	return &config.OptionsNbsper{
		CaseSensitive: false,
	}
}

func getNbspers() []nbsperKinds.Nbsper {
	nbspers := make([]nbsperKinds.Nbsper, 0)

	// The order matters as the output is modified in-place...
	nbspers = append(nbspers, &nbsperKinds.NbsperReplace{})
	nbspers = append(nbspers, &nbsperKinds.NbsperInside{})
	nbspers = append(nbspers, &nbsperKinds.NbsperAround{})
	nbspers = append(nbspers, &nbsperKinds.NbsperBefore{})
	nbspers = append(nbspers, &nbsperKinds.NbsperAfter{})

	return nbspers
}
