package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"strings"
)

const DefinitionDir = "./definitions/"
const DefinitionFileExt = "yaml"

func LoadDefinitions() map[string][]string {
	fileList, err := ioutil.ReadDir(DefinitionDir)
	if err != nil {
		log.Fatalf("Could not list the definition folder: %v", err)
	}

	definitionMap := make(map[string][]string)

	for _, file := range fileList {
		fileName := file.Name()

		extension := fileName[strings.IndexByte(fileName, '.')+1:]
		if extension == DefinitionFileExt {
			fileContent, err := ioutil.ReadFile(DefinitionDir + fileName)
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
