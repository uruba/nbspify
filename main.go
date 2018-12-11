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
