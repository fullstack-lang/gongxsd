package models

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// Generate process the xsd
//
// - all Named Complex Type => Gongstruct
// - all Elements with inlined ComplexType => Gongstruct
func Generate(stage *StageStruct, outputFilePath string) {

	// generate the typescript file
	codeGO := ModelsFileTemplate

	templInsertionLevel0 := make(map[ModelsFileTmplLevel0]string)
	for subStructTemplate := range ModelsFileTmplLevel0Code {
		templInsertionLevel0[subStructTemplate] = ""
	}

	for ct := range *GetGongstructInstancesSet[ComplexType](stage) {

		if !ct.IsInlined {
			templInsertionLevel0[ModelsFileTmplLevel0AllGongstructsCode] += Replace1(
				ModelsFileTmplLevel1Code[ModelsFileTmplLevel1OneGongstructCode],
				"{{"+string(rune(ModelsFileTmplLevel2Structname))+"}}", capitalizeFirstLetter(ct.Name),
			)
		}
	}

	for insertionPerStructId := ModelsFileTmplLevel0(0); insertionPerStructId < ModelsFileTmplLevel0Nb; insertionPerStructId++ {
		toReplace := "{{" + string(rune(insertionPerStructId)) + "}}"
		codeGO = strings.ReplaceAll(codeGO, toReplace, templInsertionLevel0[insertionPerStructId])
	}

	file, err := os.Create(outputFilePath)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()
	fmt.Fprint(file, codeGO)
}
