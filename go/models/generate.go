package models

import (
	"cmp"
	"fmt"
	"log"
	"os"
	"slices"
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

	type ParticleCode struct {
		particle Particle
		code     string
	}
	var particleCodes []*ParticleCode

	for _, ct := range GetGongstrucsSorted[*ComplexType](stage) {

		fields := ct.GetFields(stage)

		var comment string
		if annotation := ct.Annotation; annotation != nil {
			if len(annotation.Documentations) > 0 {
				comment = PrefixLines(annotation.Documentations[0].Text)
			}
		}

		if !ct.IsAnonymous {
			tmp := Replace4(
				ModelsFileTmplLevel1Code[ModelsFileTmplLevel1NamedStructCode],

				"{{"+string(rune(ModelsFileTmplLevel2Comment))+"}}",
				comment,

				"{{"+string(rune(ModelsFileTmplLevel2Structname))+"}}", ct.GoIdentifier,

				"{{"+string(rune(ModelsFileTmplLevel2Source))+"}}",
				`named complex type "`+ct.Name+`"`,

				"{{"+string(rune(ModelsFileTmplLevel2Fields))+"}}",
				fields,
			)
			particleCodes = append(particleCodes, &ParticleCode{
				particle: ct,
				code:     tmp,
			})
		} else {
			var outerElementName string
			if ct.OuterElement != nil {
				outerElementName = ct.OuterElement.Name
			}
			tmp := Replace4(
				ModelsFileTmplLevel1Code[ModelsFileTmplLevel1NamedStructCode],

				"{{"+string(rune(ModelsFileTmplLevel2Comment))+"}}",
				comment,

				"{{"+string(rune(ModelsFileTmplLevel2Structname))+"}}", ct.GoIdentifier,

				"{{"+string(rune(ModelsFileTmplLevel2Source))+"}}",
				`within outer element "`+outerElementName+`"`,

				"{{"+string(rune(ModelsFileTmplLevel2Fields))+"}}",
				fields,
			)
			particleCodes = append(particleCodes, &ParticleCode{
				particle: ct,
				code:     tmp,
			})
		}

	}

	// groups are generated into unamed struct that can be composed
	// into named struct
	for _, group := range GetGongstrucsSorted[*Group](stage) {

		// not the inline complex type
		if group.Ref != "" {
			continue
		}

		fields := group.GetFields(stage)
		tmp := Replace3(
			ModelsFileTmplLevel1Code[ModelsFileTmplLevel1UnNamedStructCode],

			"{{"+string(rune(ModelsFileTmplLevel2Structname))+"}}", group.GoIdentifier,

			"{{"+string(rune(ModelsFileTmplLevel2Source))+"}}",
			`named group "`+group.Name+`"`,

			"{{"+string(rune(ModelsFileTmplLevel2Fields))+"}}",
			fields,
		)
		particleCodes = append(particleCodes, &ParticleCode{
			particle: group,
			code:     tmp,
		})
	}

	for _, ag := range GetGongstrucsSorted[*AttributeGroup](stage) {

		// not the inline complex type
		if ag.Ref != "" {
			continue
		}

		stMap := make(map[string]*SimpleType)
		for st := range *GetGongstructInstancesSet[SimpleType](stage) {
			stMap[st.Name] = st
		}
		agMap := make(map[string]*AttributeGroup)
		for ag := range *GetGongstructInstancesSet[AttributeGroup](stage) {
			agMap[ag.Name] = ag
		}

		var fields string
		setOfGoIdentifiers := make(map[string]any)
		generateAttributes(ag.Attributes, stMap, setOfGoIdentifiers, &fields)

		for _, referencedAg := range ag.AttributeGroups {
			if referencedAg.Ref != "" {
				if namedAg, ok := agMap[referencedAg.Ref]; ok {
					fields += "\n\n\t// generated from attribute group \"" + referencedAg.Ref +
						"\n\t" + namedAg.GoIdentifier
				} else {
					log.Fatalln("Unkown attribute group", referencedAg.Ref)
				}
			}
		}

		tmp := Replace3(
			ModelsFileTmplLevel1Code[ModelsFileTmplLevel1UnNamedStructCode],

			"{{"+string(rune(ModelsFileTmplLevel2Structname))+"}}", ag.GoIdentifier,

			"{{"+string(rune(ModelsFileTmplLevel2Source))+"}}",
			`named attribute group "`+ag.Name+`"`,

			"{{"+string(rune(ModelsFileTmplLevel2Fields))+"}}",
			fields,
		)
		particleCodes = append(particleCodes, &ParticleCode{
			particle: ag,
			code:     tmp,
		})
	}

	// elements do not need to be translated into gong struct
	for _, element := range GetGongstrucsSorted[*Element](stage) {

		_ = element
	}

	// parse THE schema
	if len(GetGongstrucsSorted[*Schema](stage)) != 1 {
		log.Fatalln("an XSD (XML Schema Definition) cannot contain more than one " +
			"<xs:schema> element directly within a single XSD file. The <xs:schema> element " +
			"is the root element of the schema, and there can only be one root element in an XML document.")
	}

	schema := GetGongstrucsSorted[*Schema](stage)[0]
	for _, element := range schema.Elements {
		// add the XMLName because it is a root element
		fields := "\n\n\t// necessary since it is a root element" +
			"\n\tXMLName xml.Name `xml:\"" + element.NameXSD + "\"`"

		var comment string
		if annotation := element.Annotation; annotation != nil {
			if len(annotation.Documentations) > 0 {
				comment = PrefixLines(annotation.Documentations[0].Text)
			}
		}

		if element.ComplexType != nil {
			fields += "\n\n\t// generated from inline complex type" +
				"\n\t" + element.ComplexType.GoIdentifier

			tmp := Replace4(
				ModelsFileTmplLevel1Code[ModelsFileTmplLevel1NamedStructCode],

				"{{"+string(rune(ModelsFileTmplLevel2Comment))+"}}",
				comment,

				"{{"+string(rune(ModelsFileTmplLevel2Structname))+"}}",
				element.GoIdentifier,

				"{{"+string(rune(ModelsFileTmplLevel2Source))+"}}",
				"element "+element.NameXSD+" within root schema",

				"{{"+string(rune(ModelsFileTmplLevel2Fields))+"}}",
				fields,
			)
			particleCodes = append(particleCodes, &ParticleCode{
				particle: element,
				code:     tmp,
			})
		}
	}

	for _, st := range GetGongstrucsSorted[*SimpleType](stage) {
		if st.IsStringEnumerate() {
			// log.Println("String Enumerate", st.Name)
		}
	}
	// generate all particle codes
	slices.SortFunc(particleCodes,
		func(a, b *ParticleCode) int {
			return cmp.Compare(a.particle.GetOrder(), b.particle.GetOrder())
		})

	for _, pc := range particleCodes {
		if pc.particle.GetOrder() == 0 {
			log.Println("order is zero")
		}
		templInsertionLevel0[ModelsFileTmplLevel0AllGongstructsCode] += pc.code
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
