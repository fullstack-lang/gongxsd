package models

import (
	"fmt"
	"log"
	"os"
)

func Generate(stage *StageStruct, outputFilePath string) {

	// generate the typescript file
	codeGO := ModelsFileTemplate

	file, err := os.Create(outputFilePath)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()
	fmt.Fprint(file, codeGO)
}
