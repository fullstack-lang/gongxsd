// generated code - do not edit
package models

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

const marshallRes = `package {{PackageName}}

import (
	"time"

	"{{ModelsPackageName}}"

	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ time.Time

// Injection point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.StageStruct) {

	// Declaration of instances to stage{{Identifiers}}

	// Setup of values{{ValueInitializers}}

	// Setup of pointers{{PointersInitializers}}
}`

const IdentifiersDecls = `
	{{Identifier}} := (&models.{{GeneratedStructName}}{Name: ` + "`" + `{{GeneratedFieldNameValue}}` + "`" + `}).Stage(stage)`

const StringInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = ` + "`" + `{{GeneratedFieldNameValue}}` + "`"

const StringEnumInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = {{GeneratedFieldNameValue}}`

const NumberInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = {{GeneratedFieldNameValue}}`

const PointerFieldInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = {{GeneratedFieldNameValue}}`

const SliceOfPointersFieldInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = append({{Identifier}}.{{GeneratedFieldName}}, {{GeneratedFieldNameValue}})`

const TimeInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}}, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "{{GeneratedFieldNameValue}}")`

// Marshall marshall the stage content into the file as an instanciation into a stage
func (stage *StageStruct) Marshall(file *os.File, modelsPackageName, packageName string) {

	name := file.Name()

	if !strings.HasSuffix(name, ".go") {
		log.Fatalln(name + " is not a go filename")
	}

	log.Println("filename of marshall output is " + name)
	newBase := filepath.Base(file.Name())

	res := marshallRes
	res = strings.ReplaceAll(res, "{{databaseName}}", strings.ReplaceAll(newBase, ".go", ""))
	res = strings.ReplaceAll(res, "{{PackageName}}", packageName)
	res = strings.ReplaceAll(res, "{{ModelsPackageName}}", modelsPackageName)

	// map of identifiers
	// var StageMapDstructIds map[*Dstruct]string
	identifiersDecl := ""
	initializerStatements := ""
	pointersInitializesStatements := ""

	id := ""
	_ = id
	decl := ""
	_ = decl
	setValueField := ""
	_ = setValueField

	// insertion initialization of objects to stage
	map_All_Identifiers := make(map[*All]string)
	_ = map_All_Identifiers

	allOrdered := []*All{}
	for all := range stage.Alls {
		allOrdered = append(allOrdered, all)
	}
	sort.Slice(allOrdered[:], func(i, j int) bool {
		return allOrdered[i].Name < allOrdered[j].Name
	})
	if len(allOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, all := range allOrdered {

		id = generatesIdentifier("All", idx, all.Name)
		map_All_Identifiers[all] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "All")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", all.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(all.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MinOccurs")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(all.MinOccurs))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MaxOccurs")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(all.MaxOccurs))
		initializerStatements += setValueField

	}

	map_Annotation_Identifiers := make(map[*Annotation]string)
	_ = map_Annotation_Identifiers

	annotationOrdered := []*Annotation{}
	for annotation := range stage.Annotations {
		annotationOrdered = append(annotationOrdered, annotation)
	}
	sort.Slice(annotationOrdered[:], func(i, j int) bool {
		return annotationOrdered[i].Name < annotationOrdered[j].Name
	})
	if len(annotationOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, annotation := range annotationOrdered {

		id = generatesIdentifier("Annotation", idx, annotation.Name)
		map_Annotation_Identifiers[annotation] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Annotation")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", annotation.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(annotation.Name))
		initializerStatements += setValueField

	}

	map_Attribute_Identifiers := make(map[*Attribute]string)
	_ = map_Attribute_Identifiers

	attributeOrdered := []*Attribute{}
	for attribute := range stage.Attributes {
		attributeOrdered = append(attributeOrdered, attribute)
	}
	sort.Slice(attributeOrdered[:], func(i, j int) bool {
		return attributeOrdered[i].Name < attributeOrdered[j].Name
	})
	if len(attributeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, attribute := range attributeOrdered {

		id = generatesIdentifier("Attribute", idx, attribute.Name)
		map_Attribute_Identifiers[attribute] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Attribute")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", attribute.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "NameXSD")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute.NameXSD))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Type")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute.Type))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Default")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute.Default))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Use")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute.Use))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Form")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute.Form))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Fixed")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute.Fixed))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Ref")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute.Ref))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "TargetNamespace")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute.TargetNamespace))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "SimpleType")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute.SimpleType))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IDXSD")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute.IDXSD))
		initializerStatements += setValueField

	}

	map_AttributeGroup_Identifiers := make(map[*AttributeGroup]string)
	_ = map_AttributeGroup_Identifiers

	attributegroupOrdered := []*AttributeGroup{}
	for attributegroup := range stage.AttributeGroups {
		attributegroupOrdered = append(attributegroupOrdered, attributegroup)
	}
	sort.Slice(attributegroupOrdered[:], func(i, j int) bool {
		return attributegroupOrdered[i].Name < attributegroupOrdered[j].Name
	})
	if len(attributegroupOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, attributegroup := range attributegroupOrdered {

		id = generatesIdentifier("AttributeGroup", idx, attributegroup.Name)
		map_AttributeGroup_Identifiers[attributegroup] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "AttributeGroup")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", attributegroup.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attributegroup.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "NameXSD")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attributegroup.NameXSD))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Ref")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attributegroup.Ref))
		initializerStatements += setValueField

	}

	map_Choice_Identifiers := make(map[*Choice]string)
	_ = map_Choice_Identifiers

	choiceOrdered := []*Choice{}
	for choice := range stage.Choices {
		choiceOrdered = append(choiceOrdered, choice)
	}
	sort.Slice(choiceOrdered[:], func(i, j int) bool {
		return choiceOrdered[i].Name < choiceOrdered[j].Name
	})
	if len(choiceOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, choice := range choiceOrdered {

		id = generatesIdentifier("Choice", idx, choice.Name)
		map_Choice_Identifiers[choice] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Choice")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", choice.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(choice.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MinOccurs")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(choice.MinOccurs))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MaxOccurs")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(choice.MaxOccurs))
		initializerStatements += setValueField

	}

	map_ComplexType_Identifiers := make(map[*ComplexType]string)
	_ = map_ComplexType_Identifiers

	complextypeOrdered := []*ComplexType{}
	for complextype := range stage.ComplexTypes {
		complextypeOrdered = append(complextypeOrdered, complextype)
	}
	sort.Slice(complextypeOrdered[:], func(i, j int) bool {
		return complextypeOrdered[i].Name < complextypeOrdered[j].Name
	})
	if len(complextypeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, complextype := range complextypeOrdered {

		id = generatesIdentifier("ComplexType", idx, complextype.Name)
		map_ComplexType_Identifiers[complextype] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ComplexType")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", complextype.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(complextype.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsInlined")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", complextype.IsInlined))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "NameXSD")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(complextype.NameXSD))
		initializerStatements += setValueField

	}

	map_Documentation_Identifiers := make(map[*Documentation]string)
	_ = map_Documentation_Identifiers

	documentationOrdered := []*Documentation{}
	for documentation := range stage.Documentations {
		documentationOrdered = append(documentationOrdered, documentation)
	}
	sort.Slice(documentationOrdered[:], func(i, j int) bool {
		return documentationOrdered[i].Name < documentationOrdered[j].Name
	})
	if len(documentationOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, documentation := range documentationOrdered {

		id = generatesIdentifier("Documentation", idx, documentation.Name)
		map_Documentation_Identifiers[documentation] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Documentation")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", documentation.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(documentation.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Text")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(documentation.Text))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Source")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(documentation.Source))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Lang")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(documentation.Lang))
		initializerStatements += setValueField

	}

	map_Element_Identifiers := make(map[*Element]string)
	_ = map_Element_Identifiers

	elementOrdered := []*Element{}
	for element := range stage.Elements {
		elementOrdered = append(elementOrdered, element)
	}
	sort.Slice(elementOrdered[:], func(i, j int) bool {
		return elementOrdered[i].Name < elementOrdered[j].Name
	})
	if len(elementOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, element := range elementOrdered {

		id = generatesIdentifier("Element", idx, element.Name)
		map_Element_Identifiers[element] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Element")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", element.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(element.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "HasNameConflict")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", element.HasNameConflict))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "GoIdentifier")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(element.GoIdentifier))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "NameXSD")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(element.NameXSD))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Type")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(element.Type))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MinOccurs")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(element.MinOccurs))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MaxOccurs")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(element.MaxOccurs))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Default")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(element.Default))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Fixed")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(element.Fixed))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Nillable")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(element.Nillable))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Ref")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(element.Ref))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Abstract")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(element.Abstract))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Form")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(element.Form))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Block")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(element.Block))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Final")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(element.Final))
		initializerStatements += setValueField

	}

	map_Enumeration_Identifiers := make(map[*Enumeration]string)
	_ = map_Enumeration_Identifiers

	enumerationOrdered := []*Enumeration{}
	for enumeration := range stage.Enumerations {
		enumerationOrdered = append(enumerationOrdered, enumeration)
	}
	sort.Slice(enumerationOrdered[:], func(i, j int) bool {
		return enumerationOrdered[i].Name < enumerationOrdered[j].Name
	})
	if len(enumerationOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, enumeration := range enumerationOrdered {

		id = generatesIdentifier("Enumeration", idx, enumeration.Name)
		map_Enumeration_Identifiers[enumeration] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Enumeration")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", enumeration.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(enumeration.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Value")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(enumeration.Value))
		initializerStatements += setValueField

	}

	map_Group_Identifiers := make(map[*Group]string)
	_ = map_Group_Identifiers

	groupOrdered := []*Group{}
	for group := range stage.Groups {
		groupOrdered = append(groupOrdered, group)
	}
	sort.Slice(groupOrdered[:], func(i, j int) bool {
		return groupOrdered[i].Name < groupOrdered[j].Name
	})
	if len(groupOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, group := range groupOrdered {

		id = generatesIdentifier("Group", idx, group.Name)
		map_Group_Identifiers[group] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Group")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", group.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(group.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "NameXSD")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(group.NameXSD))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Ref")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(group.Ref))
		initializerStatements += setValueField

	}

	map_Length_Identifiers := make(map[*Length]string)
	_ = map_Length_Identifiers

	lengthOrdered := []*Length{}
	for length := range stage.Lengths {
		lengthOrdered = append(lengthOrdered, length)
	}
	sort.Slice(lengthOrdered[:], func(i, j int) bool {
		return lengthOrdered[i].Name < lengthOrdered[j].Name
	})
	if len(lengthOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, length := range lengthOrdered {

		id = generatesIdentifier("Length", idx, length.Name)
		map_Length_Identifiers[length] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Length")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", length.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(length.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Value")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(length.Value))
		initializerStatements += setValueField

	}

	map_MaxInclusive_Identifiers := make(map[*MaxInclusive]string)
	_ = map_MaxInclusive_Identifiers

	maxinclusiveOrdered := []*MaxInclusive{}
	for maxinclusive := range stage.MaxInclusives {
		maxinclusiveOrdered = append(maxinclusiveOrdered, maxinclusive)
	}
	sort.Slice(maxinclusiveOrdered[:], func(i, j int) bool {
		return maxinclusiveOrdered[i].Name < maxinclusiveOrdered[j].Name
	})
	if len(maxinclusiveOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, maxinclusive := range maxinclusiveOrdered {

		id = generatesIdentifier("MaxInclusive", idx, maxinclusive.Name)
		map_MaxInclusive_Identifiers[maxinclusive] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "MaxInclusive")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", maxinclusive.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(maxinclusive.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Value")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(maxinclusive.Value))
		initializerStatements += setValueField

	}

	map_MaxLength_Identifiers := make(map[*MaxLength]string)
	_ = map_MaxLength_Identifiers

	maxlengthOrdered := []*MaxLength{}
	for maxlength := range stage.MaxLengths {
		maxlengthOrdered = append(maxlengthOrdered, maxlength)
	}
	sort.Slice(maxlengthOrdered[:], func(i, j int) bool {
		return maxlengthOrdered[i].Name < maxlengthOrdered[j].Name
	})
	if len(maxlengthOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, maxlength := range maxlengthOrdered {

		id = generatesIdentifier("MaxLength", idx, maxlength.Name)
		map_MaxLength_Identifiers[maxlength] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "MaxLength")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", maxlength.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(maxlength.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Value")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(maxlength.Value))
		initializerStatements += setValueField

	}

	map_MinInclusive_Identifiers := make(map[*MinInclusive]string)
	_ = map_MinInclusive_Identifiers

	mininclusiveOrdered := []*MinInclusive{}
	for mininclusive := range stage.MinInclusives {
		mininclusiveOrdered = append(mininclusiveOrdered, mininclusive)
	}
	sort.Slice(mininclusiveOrdered[:], func(i, j int) bool {
		return mininclusiveOrdered[i].Name < mininclusiveOrdered[j].Name
	})
	if len(mininclusiveOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, mininclusive := range mininclusiveOrdered {

		id = generatesIdentifier("MinInclusive", idx, mininclusive.Name)
		map_MinInclusive_Identifiers[mininclusive] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "MinInclusive")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", mininclusive.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(mininclusive.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Value")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(mininclusive.Value))
		initializerStatements += setValueField

	}

	map_MinLength_Identifiers := make(map[*MinLength]string)
	_ = map_MinLength_Identifiers

	minlengthOrdered := []*MinLength{}
	for minlength := range stage.MinLengths {
		minlengthOrdered = append(minlengthOrdered, minlength)
	}
	sort.Slice(minlengthOrdered[:], func(i, j int) bool {
		return minlengthOrdered[i].Name < minlengthOrdered[j].Name
	})
	if len(minlengthOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, minlength := range minlengthOrdered {

		id = generatesIdentifier("MinLength", idx, minlength.Name)
		map_MinLength_Identifiers[minlength] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "MinLength")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", minlength.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(minlength.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Value")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(minlength.Value))
		initializerStatements += setValueField

	}

	map_Pattern_Identifiers := make(map[*Pattern]string)
	_ = map_Pattern_Identifiers

	patternOrdered := []*Pattern{}
	for pattern := range stage.Patterns {
		patternOrdered = append(patternOrdered, pattern)
	}
	sort.Slice(patternOrdered[:], func(i, j int) bool {
		return patternOrdered[i].Name < patternOrdered[j].Name
	})
	if len(patternOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, pattern := range patternOrdered {

		id = generatesIdentifier("Pattern", idx, pattern.Name)
		map_Pattern_Identifiers[pattern] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Pattern")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", pattern.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(pattern.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Value")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(pattern.Value))
		initializerStatements += setValueField

	}

	map_Restriction_Identifiers := make(map[*Restriction]string)
	_ = map_Restriction_Identifiers

	restrictionOrdered := []*Restriction{}
	for restriction := range stage.Restrictions {
		restrictionOrdered = append(restrictionOrdered, restriction)
	}
	sort.Slice(restrictionOrdered[:], func(i, j int) bool {
		return restrictionOrdered[i].Name < restrictionOrdered[j].Name
	})
	if len(restrictionOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, restriction := range restrictionOrdered {

		id = generatesIdentifier("Restriction", idx, restriction.Name)
		map_Restriction_Identifiers[restriction] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Restriction")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", restriction.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(restriction.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Base")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(restriction.Base))
		initializerStatements += setValueField

	}

	map_Schema_Identifiers := make(map[*Schema]string)
	_ = map_Schema_Identifiers

	schemaOrdered := []*Schema{}
	for schema := range stage.Schemas {
		schemaOrdered = append(schemaOrdered, schema)
	}
	sort.Slice(schemaOrdered[:], func(i, j int) bool {
		return schemaOrdered[i].Name < schemaOrdered[j].Name
	})
	if len(schemaOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, schema := range schemaOrdered {

		id = generatesIdentifier("Schema", idx, schema.Name)
		map_Schema_Identifiers[schema] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Schema")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", schema.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(schema.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Xs")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(schema.Xs))
		initializerStatements += setValueField

	}

	map_Sequence_Identifiers := make(map[*Sequence]string)
	_ = map_Sequence_Identifiers

	sequenceOrdered := []*Sequence{}
	for sequence := range stage.Sequences {
		sequenceOrdered = append(sequenceOrdered, sequence)
	}
	sort.Slice(sequenceOrdered[:], func(i, j int) bool {
		return sequenceOrdered[i].Name < sequenceOrdered[j].Name
	})
	if len(sequenceOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, sequence := range sequenceOrdered {

		id = generatesIdentifier("Sequence", idx, sequence.Name)
		map_Sequence_Identifiers[sequence] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Sequence")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", sequence.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(sequence.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MinOccurs")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(sequence.MinOccurs))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MaxOccurs")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(sequence.MaxOccurs))
		initializerStatements += setValueField

	}

	map_SimpleType_Identifiers := make(map[*SimpleType]string)
	_ = map_SimpleType_Identifiers

	simpletypeOrdered := []*SimpleType{}
	for simpletype := range stage.SimpleTypes {
		simpletypeOrdered = append(simpletypeOrdered, simpletype)
	}
	sort.Slice(simpletypeOrdered[:], func(i, j int) bool {
		return simpletypeOrdered[i].Name < simpletypeOrdered[j].Name
	})
	if len(simpletypeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, simpletype := range simpletypeOrdered {

		id = generatesIdentifier("SimpleType", idx, simpletype.Name)
		map_SimpleType_Identifiers[simpletype] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SimpleType")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", simpletype.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(simpletype.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "NameXSD")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(simpletype.NameXSD))
		initializerStatements += setValueField

	}

	map_TotalDigit_Identifiers := make(map[*TotalDigit]string)
	_ = map_TotalDigit_Identifiers

	totaldigitOrdered := []*TotalDigit{}
	for totaldigit := range stage.TotalDigits {
		totaldigitOrdered = append(totaldigitOrdered, totaldigit)
	}
	sort.Slice(totaldigitOrdered[:], func(i, j int) bool {
		return totaldigitOrdered[i].Name < totaldigitOrdered[j].Name
	})
	if len(totaldigitOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, totaldigit := range totaldigitOrdered {

		id = generatesIdentifier("TotalDigit", idx, totaldigit.Name)
		map_TotalDigit_Identifiers[totaldigit] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "TotalDigit")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", totaldigit.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(totaldigit.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Value")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(totaldigit.Value))
		initializerStatements += setValueField

	}

	map_WhiteSpace_Identifiers := make(map[*WhiteSpace]string)
	_ = map_WhiteSpace_Identifiers

	whitespaceOrdered := []*WhiteSpace{}
	for whitespace := range stage.WhiteSpaces {
		whitespaceOrdered = append(whitespaceOrdered, whitespace)
	}
	sort.Slice(whitespaceOrdered[:], func(i, j int) bool {
		return whitespaceOrdered[i].Name < whitespaceOrdered[j].Name
	})
	if len(whitespaceOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, whitespace := range whitespaceOrdered {

		id = generatesIdentifier("WhiteSpace", idx, whitespace.Name)
		map_WhiteSpace_Identifiers[whitespace] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "WhiteSpace")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", whitespace.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(whitespace.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Value")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(whitespace.Value))
		initializerStatements += setValueField

	}

	// insertion initialization of objects to stage
	for idx, all := range allOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("All", idx, all.Name)
		map_All_Identifiers[all] = id

		// Initialisation of values
		if all.Annotation != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Annotation")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Annotation_Identifiers[all.Annotation])
			pointersInitializesStatements += setPointerField
		}

		for _, _element := range all.Elements {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Elements")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Element_Identifiers[_element])
			pointersInitializesStatements += setPointerField
		}

		for _, _group := range all.Groups {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Groups")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Group_Identifiers[_group])
			pointersInitializesStatements += setPointerField
		}

		if all.Sequence != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Sequence")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Sequence_Identifiers[all.Sequence])
			pointersInitializesStatements += setPointerField
		}

		if all.All != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "All")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_All_Identifiers[all.All])
			pointersInitializesStatements += setPointerField
		}

		if all.Choice != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Choice")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Choice_Identifiers[all.Choice])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, annotation := range annotationOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Annotation", idx, annotation.Name)
		map_Annotation_Identifiers[annotation] = id

		// Initialisation of values
		for _, _documentation := range annotation.Documentations {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Documentations")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Documentation_Identifiers[_documentation])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, attribute := range attributeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Attribute", idx, attribute.Name)
		map_Attribute_Identifiers[attribute] = id

		// Initialisation of values
		if attribute.Annotation != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Annotation")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Annotation_Identifiers[attribute.Annotation])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, attributegroup := range attributegroupOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("AttributeGroup", idx, attributegroup.Name)
		map_AttributeGroup_Identifiers[attributegroup] = id

		// Initialisation of values
		if attributegroup.Annotation != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Annotation")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Annotation_Identifiers[attributegroup.Annotation])
			pointersInitializesStatements += setPointerField
		}

		if attributegroup.AttributeGroup != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "AttributeGroup")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_AttributeGroup_Identifiers[attributegroup.AttributeGroup])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, choice := range choiceOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Choice", idx, choice.Name)
		map_Choice_Identifiers[choice] = id

		// Initialisation of values
		if choice.Annotation != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Annotation")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Annotation_Identifiers[choice.Annotation])
			pointersInitializesStatements += setPointerField
		}

		for _, _element := range choice.Elements {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Elements")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Element_Identifiers[_element])
			pointersInitializesStatements += setPointerField
		}

		for _, _group := range choice.Groups {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Groups")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Group_Identifiers[_group])
			pointersInitializesStatements += setPointerField
		}

		if choice.Sequence != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Sequence")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Sequence_Identifiers[choice.Sequence])
			pointersInitializesStatements += setPointerField
		}

		if choice.All != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "All")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_All_Identifiers[choice.All])
			pointersInitializesStatements += setPointerField
		}

		if choice.Choice != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Choice")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Choice_Identifiers[choice.Choice])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, complextype := range complextypeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ComplexType", idx, complextype.Name)
		map_ComplexType_Identifiers[complextype] = id

		// Initialisation of values
		if complextype.EnclosingElement != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "EnclosingElement")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Element_Identifiers[complextype.EnclosingElement])
			pointersInitializesStatements += setPointerField
		}

		if complextype.Annotation != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Annotation")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Annotation_Identifiers[complextype.Annotation])
			pointersInitializesStatements += setPointerField
		}

		if complextype.Sequence != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Sequence")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Sequence_Identifiers[complextype.Sequence])
			pointersInitializesStatements += setPointerField
		}

		if complextype.All != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "All")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_All_Identifiers[complextype.All])
			pointersInitializesStatements += setPointerField
		}

		if complextype.Choice != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Choice")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Choice_Identifiers[complextype.Choice])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute := range complextype.Attributes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Attributes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Attribute_Identifiers[_attribute])
			pointersInitializesStatements += setPointerField
		}

		for _, _attributegroup := range complextype.AttributeGroups {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "AttributeGroups")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_AttributeGroup_Identifiers[_attributegroup])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, documentation := range documentationOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Documentation", idx, documentation.Name)
		map_Documentation_Identifiers[documentation] = id

		// Initialisation of values
	}

	for idx, element := range elementOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Element", idx, element.Name)
		map_Element_Identifiers[element] = id

		// Initialisation of values
		if element.Annotation != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Annotation")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Annotation_Identifiers[element.Annotation])
			pointersInitializesStatements += setPointerField
		}

		if element.SimpleType != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SimpleType")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_SimpleType_Identifiers[element.SimpleType])
			pointersInitializesStatements += setPointerField
		}

		if element.ComplexType != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ComplexType")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ComplexType_Identifiers[element.ComplexType])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, enumeration := range enumerationOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Enumeration", idx, enumeration.Name)
		map_Enumeration_Identifiers[enumeration] = id

		// Initialisation of values
		if enumeration.Annotation != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Annotation")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Annotation_Identifiers[enumeration.Annotation])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, group := range groupOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Group", idx, group.Name)
		map_Group_Identifiers[group] = id

		// Initialisation of values
		if group.Annotation != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Annotation")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Annotation_Identifiers[group.Annotation])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, length := range lengthOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Length", idx, length.Name)
		map_Length_Identifiers[length] = id

		// Initialisation of values
		if length.Annotation != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Annotation")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Annotation_Identifiers[length.Annotation])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, maxinclusive := range maxinclusiveOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("MaxInclusive", idx, maxinclusive.Name)
		map_MaxInclusive_Identifiers[maxinclusive] = id

		// Initialisation of values
		if maxinclusive.Annotation != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Annotation")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Annotation_Identifiers[maxinclusive.Annotation])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, maxlength := range maxlengthOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("MaxLength", idx, maxlength.Name)
		map_MaxLength_Identifiers[maxlength] = id

		// Initialisation of values
		if maxlength.Annotation != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Annotation")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Annotation_Identifiers[maxlength.Annotation])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, mininclusive := range mininclusiveOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("MinInclusive", idx, mininclusive.Name)
		map_MinInclusive_Identifiers[mininclusive] = id

		// Initialisation of values
		if mininclusive.Annotation != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Annotation")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Annotation_Identifiers[mininclusive.Annotation])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, minlength := range minlengthOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("MinLength", idx, minlength.Name)
		map_MinLength_Identifiers[minlength] = id

		// Initialisation of values
		if minlength.Annotation != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Annotation")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Annotation_Identifiers[minlength.Annotation])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, pattern := range patternOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Pattern", idx, pattern.Name)
		map_Pattern_Identifiers[pattern] = id

		// Initialisation of values
		if pattern.Annotation != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Annotation")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Annotation_Identifiers[pattern.Annotation])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, restriction := range restrictionOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Restriction", idx, restriction.Name)
		map_Restriction_Identifiers[restriction] = id

		// Initialisation of values
		if restriction.Annotation != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Annotation")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Annotation_Identifiers[restriction.Annotation])
			pointersInitializesStatements += setPointerField
		}

		for _, _enumeration := range restriction.Enumerations {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Enumerations")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Enumeration_Identifiers[_enumeration])
			pointersInitializesStatements += setPointerField
		}

		if restriction.MinInclusive != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "MinInclusive")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_MinInclusive_Identifiers[restriction.MinInclusive])
			pointersInitializesStatements += setPointerField
		}

		if restriction.MaxInclusive != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "MaxInclusive")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_MaxInclusive_Identifiers[restriction.MaxInclusive])
			pointersInitializesStatements += setPointerField
		}

		if restriction.Pattern != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Pattern")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Pattern_Identifiers[restriction.Pattern])
			pointersInitializesStatements += setPointerField
		}

		if restriction.WhiteSpace != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "WhiteSpace")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_WhiteSpace_Identifiers[restriction.WhiteSpace])
			pointersInitializesStatements += setPointerField
		}

		if restriction.MinLength != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "MinLength")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_MinLength_Identifiers[restriction.MinLength])
			pointersInitializesStatements += setPointerField
		}

		if restriction.MaxLength != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "MaxLength")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_MaxLength_Identifiers[restriction.MaxLength])
			pointersInitializesStatements += setPointerField
		}

		if restriction.Length != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Length")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Length_Identifiers[restriction.Length])
			pointersInitializesStatements += setPointerField
		}

		if restriction.TotalDigit != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "TotalDigit")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_TotalDigit_Identifiers[restriction.TotalDigit])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, schema := range schemaOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Schema", idx, schema.Name)
		map_Schema_Identifiers[schema] = id

		// Initialisation of values
		if schema.Annotation != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Annotation")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Annotation_Identifiers[schema.Annotation])
			pointersInitializesStatements += setPointerField
		}

		for _, _element := range schema.Elements {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Elements")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Element_Identifiers[_element])
			pointersInitializesStatements += setPointerField
		}

		for _, _simpletype := range schema.SimpleTypes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SimpleTypes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_SimpleType_Identifiers[_simpletype])
			pointersInitializesStatements += setPointerField
		}

		for _, _complextype := range schema.ComplexTypes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ComplexTypes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ComplexType_Identifiers[_complextype])
			pointersInitializesStatements += setPointerField
		}

		for _, _attributegroup := range schema.AttributeGroups {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "AttributeGroups")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_AttributeGroup_Identifiers[_attributegroup])
			pointersInitializesStatements += setPointerField
		}

		for _, _group := range schema.Groups {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Groups")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Group_Identifiers[_group])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, sequence := range sequenceOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Sequence", idx, sequence.Name)
		map_Sequence_Identifiers[sequence] = id

		// Initialisation of values
		if sequence.Annotation != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Annotation")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Annotation_Identifiers[sequence.Annotation])
			pointersInitializesStatements += setPointerField
		}

		for _, _element := range sequence.Elements {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Elements")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Element_Identifiers[_element])
			pointersInitializesStatements += setPointerField
		}

		for _, _group := range sequence.Groups {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Groups")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Group_Identifiers[_group])
			pointersInitializesStatements += setPointerField
		}

		if sequence.Sequence != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Sequence")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Sequence_Identifiers[sequence.Sequence])
			pointersInitializesStatements += setPointerField
		}

		if sequence.All != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "All")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_All_Identifiers[sequence.All])
			pointersInitializesStatements += setPointerField
		}

		if sequence.Choice != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Choice")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Choice_Identifiers[sequence.Choice])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, simpletype := range simpletypeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("SimpleType", idx, simpletype.Name)
		map_SimpleType_Identifiers[simpletype] = id

		// Initialisation of values
		if simpletype.Annotation != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Annotation")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Annotation_Identifiers[simpletype.Annotation])
			pointersInitializesStatements += setPointerField
		}

		if simpletype.Restriction != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Restriction")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Restriction_Identifiers[simpletype.Restriction])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, totaldigit := range totaldigitOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("TotalDigit", idx, totaldigit.Name)
		map_TotalDigit_Identifiers[totaldigit] = id

		// Initialisation of values
		if totaldigit.Annotation != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Annotation")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Annotation_Identifiers[totaldigit.Annotation])
			pointersInitializesStatements += setPointerField
		}

	}

	for idx, whitespace := range whitespaceOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("WhiteSpace", idx, whitespace.Name)
		map_WhiteSpace_Identifiers[whitespace] = id

		// Initialisation of values
		if whitespace.Annotation != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Annotation")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Annotation_Identifiers[whitespace.Annotation])
			pointersInitializesStatements += setPointerField
		}

	}

	res = strings.ReplaceAll(res, "{{Identifiers}}", identifiersDecl)
	res = strings.ReplaceAll(res, "{{ValueInitializers}}", initializerStatements)
	res = strings.ReplaceAll(res, "{{PointersInitializers}}", pointersInitializesStatements)

	if stage.MetaPackageImportAlias != "" {
		res = strings.ReplaceAll(res, "{{ImportPackageDeclaration}}",
			fmt.Sprintf("\n\t%s %s", stage.MetaPackageImportAlias, stage.MetaPackageImportPath))

		res = strings.ReplaceAll(res, "{{ImportPackageDummyDeclaration}}",
			fmt.Sprintf("\nvar _ %s.StageStruct",
				stage.MetaPackageImportAlias))

		var entries string

		// regenerate the map of doc link renaming
		// the key and value are set to the value because
		// if it has been renamed, this is the new value that matters
		valuesOrdered := make([]GONG__Identifier, 0)
		for _, value := range stage.Map_DocLink_Renaming {
			valuesOrdered = append(valuesOrdered, value)
		}
		sort.Slice(valuesOrdered[:], func(i, j int) bool {
			return valuesOrdered[i].Ident < valuesOrdered[j].Ident
		})
		for _, value := range valuesOrdered {

			// get the number of points in the value to find if it is a field
			// or a struct

			switch value.Type {
			case GONG__ENUM_CAST_INT:
				entries += fmt.Sprintf("\n\n\t\"%s\": %s(0),", value.Ident, value.Ident)
			case GONG__ENUM_CAST_STRING:
				entries += fmt.Sprintf("\n\n\t\"%s\": %s(\"\"),", value.Ident, value.Ident)
			case GONG__FIELD_VALUE:
				// substitute the second point with "{})."
				joker := "__substitute_for_first_point__"
				valueIdentifier := strings.Replace(value.Ident, ".", joker, 1)
				valueIdentifier = strings.Replace(valueIdentifier, ".", "{}).", 1)
				valueIdentifier = strings.Replace(valueIdentifier, joker, ".", 1)
				entries += fmt.Sprintf("\n\n\t\"%s\": (%s,", value.Ident, valueIdentifier)
			case GONG__IDENTIFIER_CONST:
				entries += fmt.Sprintf("\n\n\t\"%s\": %s,", value.Ident, value.Ident)
			case GONG__STRUCT_INSTANCE:
				entries += fmt.Sprintf("\n\n\t\"%s\": &(%s{}),", value.Ident, value.Ident)
			}
		}

		// res = strings.ReplaceAll(res, "{{EntriesDocLinkStringDocLinkIdentifier}}", entries)
	}

	fmt.Fprintln(file, res)
}

// unique identifier per struct
func generatesIdentifier(gongStructName string, idx int, instanceName string) (identifier string) {

	identifier = instanceName
	// Make a Regex to say we only want letters and numbers
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	processedString := reg.ReplaceAllString(instanceName, "_")

	identifier = fmt.Sprintf("__%s__%06d_%s", gongStructName, idx, processedString)

	return
}
