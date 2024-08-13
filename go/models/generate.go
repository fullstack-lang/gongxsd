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

	for _, ct := range GetGongstrucsSorted[*ComplexType](stage) {

		// not the inline complex type
		if ct.IsAnonymous {
			continue
		}

		fields := ct.GetFields(stage)
		templInsertionLevel0[ModelsFileTmplLevel0AllGongstructsCode] += Replace3(
			ModelsFileTmplLevel1Code[ModelsFileTmplLevel1OneGongstructCode],

			"{{"+string(rune(ModelsFileTmplLevel2Structname))+"}}", ct.GoIdentifier,

			"{{"+string(rune(ModelsFileTmplLevel2Source))+"}}",
			`named complex type "`+ct.Name+`"`,

			"{{"+string(rune(ModelsFileTmplLevel2Fields))+"}}",
			fields,
		)
	}

	for _, group := range GetGongstrucsSorted[*Group](stage) {

		// not the inline complex type
		if group.Ref != "" {
			continue
		}

		fields := group.GetFields(stage)
		templInsertionLevel0[ModelsFileTmplLevel0AllGongstructsCode] += Replace3(
			ModelsFileTmplLevel1Code[ModelsFileTmplLevel1OneGongstructCode],

			"{{"+string(rune(ModelsFileTmplLevel2Structname))+"}}", group.GoIdentifier,

			"{{"+string(rune(ModelsFileTmplLevel2Source))+"}}",
			`named group "`+group.Name+`"`,

			"{{"+string(rune(ModelsFileTmplLevel2Fields))+"}}",
			fields,
		)
	}

	// elements with inline complex type
	for _, element := range GetGongstrucsSorted[*Element](stage) {

		if element.ComplexType == nil {
			continue
		}

		// fail loud and fast
		if !element.ComplexType.IsAnonymous {
			log.Fatal("element", element.Name, "has inlined complex type")
		}

		source := `inlined complex type within element "` + element.Name + `"`

		if element.HasNameConflict {
			source += `
// Identifier is post fixed because more than one xsd element has the name "` + element.Name + `"`
		}

		fields := element.ComplexType.GetFields(stage)

		templInsertionLevel0[ModelsFileTmplLevel0AllGongstructsCode] += Replace3(
			ModelsFileTmplLevel1Code[ModelsFileTmplLevel1OneGongstructCode],

			"{{"+string(rune(ModelsFileTmplLevel2Structname))+"}}",
			element.GoIdentifier,

			"{{"+string(rune(ModelsFileTmplLevel2Source))+"}}",
			source,

			"{{"+string(rune(ModelsFileTmplLevel2Fields))+"}}",
			fields,
		)

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
