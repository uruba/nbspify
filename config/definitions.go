package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"path/filepath"
)

const DefinitionFile = "definitions.yaml"

func LoadDefinitions() map[string][]string {
	definitionMap := make(map[string][]string)

	fileName := filepath.Join(GetConfigDir(), DefinitionFile)
	fileContent, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalf("Could not read the definition file: %v", err)
	}

	fileDefinitionChunk := make(map[string]map[string][]string)
	err = yaml.Unmarshal(fileContent, &fileDefinitionChunk)
	if err != nil {
		log.Fatalf("Could not parse the definition file: %v", err)
	}

	// TODO - localization
	for _, definitionKeys := range fileDefinitionChunk {
		for definitionKey, definitionValues := range definitionKeys {
			// TODO - deduplication
			definitionMap[definitionKey] = append(definitionMap[definitionKey], definitionValues...)
		}
	}

	return definitionMap
}
