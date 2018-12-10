package config

import (
	"github.com/gobuffalo/packr"
	"gopkg.in/yaml.v2"
	"log"
	"strings"
)

const DefinitionDir = "../definitions"
const DefinitionFileExt = "yaml"

func LoadDefinitions() map[string][]string {
	box := packr.NewBox(DefinitionDir)
	fileNames := box.List()
	if len(fileNames) == 0 {
		log.Fatalf("No definition files have been bundled")
	}

	definitionMap := make(map[string][]string)

	for _, fileName := range fileNames {
		extension := fileName[strings.IndexByte(fileName, '.')+1:]
		if extension == DefinitionFileExt {
			fileContent, err := box.Find(fileName)
			if err != nil {
				log.Fatalf("Could not read a definition file: %v", err)
			}

			fileDefinitionChunk := make(map[string][]string)
			err = yaml.Unmarshal(fileContent, &fileDefinitionChunk)
			if err != nil {
				log.Fatalf("Could not parse a definition file: %v", err)
			}

			for definitionKey, definitionValues := range fileDefinitionChunk {
				// TODO - deduplication
				definitionMap[definitionKey] = append(definitionMap[definitionKey], definitionValues...)
			}
		}
	}

	return definitionMap
}
