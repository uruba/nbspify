package main

import (
	"fmt"
	"nbspify/config"
	"nbspify/input"
	nbsperDefinitions "nbspify/nbspers"
)

func main() {
	processedText := input.ReadInput()

	nbspers := getNbspers()
	definitions := config.LoadDefinitions()

	if nbspers != nil {
		for _, nbsper := range nbspers {
			definition := definitions[nbsper.GetCode()]

			if len(definition) > 0 {
				processedText = nbsper.Apply(processedText, definition)
			}
		}
	}

	// STDOUT
	fmt.Print(processedText)
}

func getNbspers() []nbsperDefinitions.Nbsper {
	nbspers := make([]nbsperDefinitions.Nbsper, 0)

	// The order matters as the output is modified in-place...
	nbspers = append(nbspers, &nbsperDefinitions.NbsperReplace{})
	nbspers = append(nbspers, &nbsperDefinitions.NbsperInside{})
	nbspers = append(nbspers, &nbsperDefinitions.NbsperAround{})
	nbspers = append(nbspers, &nbsperDefinitions.NbsperBefore{})
	nbspers = append(nbspers, &nbsperDefinitions.NbsperAfter{})

	return nbspers
}
