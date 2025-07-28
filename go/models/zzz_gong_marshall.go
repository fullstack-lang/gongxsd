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

	"github.com/sergi/go-diff/diffmatchpatch"
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

// _ point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.Stage) {

	const __write__local_time = "{{LocalTimeStamp}}"
	const __write__utc_time__ = "{{UTCTimeStamp}}"

	const __commitId__ = "{{CommitId}}"

	// Declaration of instances to stage{{Identifiers}}

	// Setup of values{{ValueInitializers}}

	// Setup of pointers{{PointersInitializers}}
}
`

const IdentifiersDecls = `
	{{Identifier}} := (&models.{{GeneratedStructName}}{}).Stage(stage)`

const StringInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = ` + "`" + `{{GeneratedFieldNameValue}}` + "`"

const MetaFieldStructInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = ` + `{{GeneratedFieldNameValue}}`

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
func (stage *Stage) Marshall(file *os.File, modelsPackageName, packageName string) {

	name := file.Name()

	if !strings.HasSuffix(name, ".go") {
		log.Fatalln(name + " is not a go filename")
	}

	log.Printf("Marshalling %s", name)
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
		alli := allOrdered[i]
		allj := allOrdered[j]
		alli_order, oki := stage.AllMap_Staged_Order[alli]
		allj_order, okj := stage.AllMap_Staged_Order[allj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return alli_order < allj_order
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
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "OuterElementName")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(all.OuterElementName))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Order")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", all.Order))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Depth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", all.Depth))
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
		annotationi := annotationOrdered[i]
		annotationj := annotationOrdered[j]
		annotationi_order, oki := stage.AnnotationMap_Staged_Order[annotationi]
		annotationj_order, okj := stage.AnnotationMap_Staged_Order[annotationj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return annotationi_order < annotationj_order
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
		attributei := attributeOrdered[i]
		attributej := attributeOrdered[j]
		attributei_order, oki := stage.AttributeMap_Staged_Order[attributei]
		attributej_order, okj := stage.AttributeMap_Staged_Order[attributej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return attributei_order < attributej_order
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

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "HasNameConflict")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", attribute.HasNameConflict))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "GoIdentifier")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attribute.GoIdentifier))
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
		attributegroupi := attributegroupOrdered[i]
		attributegroupj := attributegroupOrdered[j]
		attributegroupi_order, oki := stage.AttributeGroupMap_Staged_Order[attributegroupi]
		attributegroupj_order, okj := stage.AttributeGroupMap_Staged_Order[attributegroupj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return attributegroupi_order < attributegroupj_order
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

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "HasNameConflict")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", attributegroup.HasNameConflict))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "GoIdentifier")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attributegroup.GoIdentifier))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Ref")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(attributegroup.Ref))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Order")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", attributegroup.Order))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Depth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", attributegroup.Depth))
		initializerStatements += setValueField

	}

	map_Choice_Identifiers := make(map[*Choice]string)
	_ = map_Choice_Identifiers

	choiceOrdered := []*Choice{}
	for choice := range stage.Choices {
		choiceOrdered = append(choiceOrdered, choice)
	}
	sort.Slice(choiceOrdered[:], func(i, j int) bool {
		choicei := choiceOrdered[i]
		choicej := choiceOrdered[j]
		choicei_order, oki := stage.ChoiceMap_Staged_Order[choicei]
		choicej_order, okj := stage.ChoiceMap_Staged_Order[choicej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return choicei_order < choicej_order
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
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "OuterElementName")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(choice.OuterElementName))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Order")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", choice.Order))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Depth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", choice.Depth))
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

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsDuplicatedInXSD")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", choice.IsDuplicatedInXSD))
		initializerStatements += setValueField

	}

	map_ComplexContent_Identifiers := make(map[*ComplexContent]string)
	_ = map_ComplexContent_Identifiers

	complexcontentOrdered := []*ComplexContent{}
	for complexcontent := range stage.ComplexContents {
		complexcontentOrdered = append(complexcontentOrdered, complexcontent)
	}
	sort.Slice(complexcontentOrdered[:], func(i, j int) bool {
		complexcontenti := complexcontentOrdered[i]
		complexcontentj := complexcontentOrdered[j]
		complexcontenti_order, oki := stage.ComplexContentMap_Staged_Order[complexcontenti]
		complexcontentj_order, okj := stage.ComplexContentMap_Staged_Order[complexcontentj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return complexcontenti_order < complexcontentj_order
	})
	if len(complexcontentOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, complexcontent := range complexcontentOrdered {

		id = generatesIdentifier("ComplexContent", idx, complexcontent.Name)
		map_ComplexContent_Identifiers[complexcontent] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ComplexContent")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", complexcontent.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(complexcontent.Name))
		initializerStatements += setValueField

	}

	map_ComplexType_Identifiers := make(map[*ComplexType]string)
	_ = map_ComplexType_Identifiers

	complextypeOrdered := []*ComplexType{}
	for complextype := range stage.ComplexTypes {
		complextypeOrdered = append(complextypeOrdered, complextype)
	}
	sort.Slice(complextypeOrdered[:], func(i, j int) bool {
		complextypei := complextypeOrdered[i]
		complextypej := complextypeOrdered[j]
		complextypei_order, oki := stage.ComplexTypeMap_Staged_Order[complextypei]
		complextypej_order, okj := stage.ComplexTypeMap_Staged_Order[complextypej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return complextypei_order < complextypej_order
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
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "HasNameConflict")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", complextype.HasNameConflict))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "GoIdentifier")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(complextype.GoIdentifier))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsAnonymous")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", complextype.IsAnonymous))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "NameXSD")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(complextype.NameXSD))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "OuterElementName")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(complextype.OuterElementName))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Order")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", complextype.Order))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Depth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", complextype.Depth))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MinOccurs")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(complextype.MinOccurs))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MaxOccurs")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(complextype.MaxOccurs))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsDuplicatedInXSD")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", complextype.IsDuplicatedInXSD))
		initializerStatements += setValueField

	}

	map_Documentation_Identifiers := make(map[*Documentation]string)
	_ = map_Documentation_Identifiers

	documentationOrdered := []*Documentation{}
	for documentation := range stage.Documentations {
		documentationOrdered = append(documentationOrdered, documentation)
	}
	sort.Slice(documentationOrdered[:], func(i, j int) bool {
		documentationi := documentationOrdered[i]
		documentationj := documentationOrdered[j]
		documentationi_order, oki := stage.DocumentationMap_Staged_Order[documentationi]
		documentationj_order, okj := stage.DocumentationMap_Staged_Order[documentationj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return documentationi_order < documentationj_order
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
		elementi := elementOrdered[i]
		elementj := elementOrdered[j]
		elementi_order, oki := stage.ElementMap_Staged_Order[elementi]
		elementj_order, okj := stage.ElementMap_Staged_Order[elementj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return elementi_order < elementj_order
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
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Order")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", element.Order))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Depth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", element.Depth))
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

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsDuplicatedInXSD")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", element.IsDuplicatedInXSD))
		initializerStatements += setValueField

	}

	map_Enumeration_Identifiers := make(map[*Enumeration]string)
	_ = map_Enumeration_Identifiers

	enumerationOrdered := []*Enumeration{}
	for enumeration := range stage.Enumerations {
		enumerationOrdered = append(enumerationOrdered, enumeration)
	}
	sort.Slice(enumerationOrdered[:], func(i, j int) bool {
		enumerationi := enumerationOrdered[i]
		enumerationj := enumerationOrdered[j]
		enumerationi_order, oki := stage.EnumerationMap_Staged_Order[enumerationi]
		enumerationj_order, okj := stage.EnumerationMap_Staged_Order[enumerationj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return enumerationi_order < enumerationj_order
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

	map_Extension_Identifiers := make(map[*Extension]string)
	_ = map_Extension_Identifiers

	extensionOrdered := []*Extension{}
	for extension := range stage.Extensions {
		extensionOrdered = append(extensionOrdered, extension)
	}
	sort.Slice(extensionOrdered[:], func(i, j int) bool {
		extensioni := extensionOrdered[i]
		extensionj := extensionOrdered[j]
		extensioni_order, oki := stage.ExtensionMap_Staged_Order[extensioni]
		extensionj_order, okj := stage.ExtensionMap_Staged_Order[extensionj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return extensioni_order < extensionj_order
	})
	if len(extensionOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, extension := range extensionOrdered {

		id = generatesIdentifier("Extension", idx, extension.Name)
		map_Extension_Identifiers[extension] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Extension")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", extension.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(extension.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "OuterElementName")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(extension.OuterElementName))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Order")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", extension.Order))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Depth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", extension.Depth))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MinOccurs")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(extension.MinOccurs))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MaxOccurs")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(extension.MaxOccurs))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Base")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(extension.Base))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Ref")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(extension.Ref))
		initializerStatements += setValueField

	}

	map_Group_Identifiers := make(map[*Group]string)
	_ = map_Group_Identifiers

	groupOrdered := []*Group{}
	for group := range stage.Groups {
		groupOrdered = append(groupOrdered, group)
	}
	sort.Slice(groupOrdered[:], func(i, j int) bool {
		groupi := groupOrdered[i]
		groupj := groupOrdered[j]
		groupi_order, oki := stage.GroupMap_Staged_Order[groupi]
		groupj_order, okj := stage.GroupMap_Staged_Order[groupj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return groupi_order < groupj_order
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

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsAnonymous")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", group.IsAnonymous))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "HasNameConflict")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", group.HasNameConflict))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "GoIdentifier")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(group.GoIdentifier))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "OuterElementName")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(group.OuterElementName))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Order")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", group.Order))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Depth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", group.Depth))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MinOccurs")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(group.MinOccurs))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MaxOccurs")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(group.MaxOccurs))
		initializerStatements += setValueField

	}

	map_Length_Identifiers := make(map[*Length]string)
	_ = map_Length_Identifiers

	lengthOrdered := []*Length{}
	for length := range stage.Lengths {
		lengthOrdered = append(lengthOrdered, length)
	}
	sort.Slice(lengthOrdered[:], func(i, j int) bool {
		lengthi := lengthOrdered[i]
		lengthj := lengthOrdered[j]
		lengthi_order, oki := stage.LengthMap_Staged_Order[lengthi]
		lengthj_order, okj := stage.LengthMap_Staged_Order[lengthj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return lengthi_order < lengthj_order
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
		maxinclusivei := maxinclusiveOrdered[i]
		maxinclusivej := maxinclusiveOrdered[j]
		maxinclusivei_order, oki := stage.MaxInclusiveMap_Staged_Order[maxinclusivei]
		maxinclusivej_order, okj := stage.MaxInclusiveMap_Staged_Order[maxinclusivej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return maxinclusivei_order < maxinclusivej_order
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
		maxlengthi := maxlengthOrdered[i]
		maxlengthj := maxlengthOrdered[j]
		maxlengthi_order, oki := stage.MaxLengthMap_Staged_Order[maxlengthi]
		maxlengthj_order, okj := stage.MaxLengthMap_Staged_Order[maxlengthj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return maxlengthi_order < maxlengthj_order
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
		mininclusivei := mininclusiveOrdered[i]
		mininclusivej := mininclusiveOrdered[j]
		mininclusivei_order, oki := stage.MinInclusiveMap_Staged_Order[mininclusivei]
		mininclusivej_order, okj := stage.MinInclusiveMap_Staged_Order[mininclusivej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return mininclusivei_order < mininclusivej_order
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
		minlengthi := minlengthOrdered[i]
		minlengthj := minlengthOrdered[j]
		minlengthi_order, oki := stage.MinLengthMap_Staged_Order[minlengthi]
		minlengthj_order, okj := stage.MinLengthMap_Staged_Order[minlengthj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return minlengthi_order < minlengthj_order
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
		patterni := patternOrdered[i]
		patternj := patternOrdered[j]
		patterni_order, oki := stage.PatternMap_Staged_Order[patterni]
		patternj_order, okj := stage.PatternMap_Staged_Order[patternj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return patterni_order < patternj_order
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
		restrictioni := restrictionOrdered[i]
		restrictionj := restrictionOrdered[j]
		restrictioni_order, oki := stage.RestrictionMap_Staged_Order[restrictioni]
		restrictionj_order, okj := stage.RestrictionMap_Staged_Order[restrictionj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return restrictioni_order < restrictionj_order
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
		schemai := schemaOrdered[i]
		schemaj := schemaOrdered[j]
		schemai_order, oki := stage.SchemaMap_Staged_Order[schemai]
		schemaj_order, okj := stage.SchemaMap_Staged_Order[schemaj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return schemai_order < schemaj_order
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

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Order")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", schema.Order))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Depth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", schema.Depth))
		initializerStatements += setValueField

	}

	map_Sequence_Identifiers := make(map[*Sequence]string)
	_ = map_Sequence_Identifiers

	sequenceOrdered := []*Sequence{}
	for sequence := range stage.Sequences {
		sequenceOrdered = append(sequenceOrdered, sequence)
	}
	sort.Slice(sequenceOrdered[:], func(i, j int) bool {
		sequencei := sequenceOrdered[i]
		sequencej := sequenceOrdered[j]
		sequencei_order, oki := stage.SequenceMap_Staged_Order[sequencei]
		sequencej_order, okj := stage.SequenceMap_Staged_Order[sequencej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return sequencei_order < sequencej_order
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
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "OuterElementName")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(sequence.OuterElementName))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Order")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", sequence.Order))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Depth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", sequence.Depth))
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

	map_SimpleContent_Identifiers := make(map[*SimpleContent]string)
	_ = map_SimpleContent_Identifiers

	simplecontentOrdered := []*SimpleContent{}
	for simplecontent := range stage.SimpleContents {
		simplecontentOrdered = append(simplecontentOrdered, simplecontent)
	}
	sort.Slice(simplecontentOrdered[:], func(i, j int) bool {
		simplecontenti := simplecontentOrdered[i]
		simplecontentj := simplecontentOrdered[j]
		simplecontenti_order, oki := stage.SimpleContentMap_Staged_Order[simplecontenti]
		simplecontentj_order, okj := stage.SimpleContentMap_Staged_Order[simplecontentj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return simplecontenti_order < simplecontentj_order
	})
	if len(simplecontentOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, simplecontent := range simplecontentOrdered {

		id = generatesIdentifier("SimpleContent", idx, simplecontent.Name)
		map_SimpleContent_Identifiers[simplecontent] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SimpleContent")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", simplecontent.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(simplecontent.Name))
		initializerStatements += setValueField

	}

	map_SimpleType_Identifiers := make(map[*SimpleType]string)
	_ = map_SimpleType_Identifiers

	simpletypeOrdered := []*SimpleType{}
	for simpletype := range stage.SimpleTypes {
		simpletypeOrdered = append(simpletypeOrdered, simpletype)
	}
	sort.Slice(simpletypeOrdered[:], func(i, j int) bool {
		simpletypei := simpletypeOrdered[i]
		simpletypej := simpletypeOrdered[j]
		simpletypei_order, oki := stage.SimpleTypeMap_Staged_Order[simpletypei]
		simpletypej_order, okj := stage.SimpleTypeMap_Staged_Order[simpletypej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return simpletypei_order < simpletypej_order
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

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Order")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", simpletype.Order))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Depth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", simpletype.Depth))
		initializerStatements += setValueField

	}

	map_TotalDigit_Identifiers := make(map[*TotalDigit]string)
	_ = map_TotalDigit_Identifiers

	totaldigitOrdered := []*TotalDigit{}
	for totaldigit := range stage.TotalDigits {
		totaldigitOrdered = append(totaldigitOrdered, totaldigit)
	}
	sort.Slice(totaldigitOrdered[:], func(i, j int) bool {
		totaldigiti := totaldigitOrdered[i]
		totaldigitj := totaldigitOrdered[j]
		totaldigiti_order, oki := stage.TotalDigitMap_Staged_Order[totaldigiti]
		totaldigitj_order, okj := stage.TotalDigitMap_Staged_Order[totaldigitj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return totaldigiti_order < totaldigitj_order
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

	map_Union_Identifiers := make(map[*Union]string)
	_ = map_Union_Identifiers

	unionOrdered := []*Union{}
	for union := range stage.Unions {
		unionOrdered = append(unionOrdered, union)
	}
	sort.Slice(unionOrdered[:], func(i, j int) bool {
		unioni := unionOrdered[i]
		unionj := unionOrdered[j]
		unioni_order, oki := stage.UnionMap_Staged_Order[unioni]
		unionj_order, okj := stage.UnionMap_Staged_Order[unionj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return unioni_order < unionj_order
	})
	if len(unionOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, union := range unionOrdered {

		id = generatesIdentifier("Union", idx, union.Name)
		map_Union_Identifiers[union] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Union")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", union.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(union.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MemberTypes")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(union.MemberTypes))
		initializerStatements += setValueField

	}

	map_WhiteSpace_Identifiers := make(map[*WhiteSpace]string)
	_ = map_WhiteSpace_Identifiers

	whitespaceOrdered := []*WhiteSpace{}
	for whitespace := range stage.WhiteSpaces {
		whitespaceOrdered = append(whitespaceOrdered, whitespace)
	}
	sort.Slice(whitespaceOrdered[:], func(i, j int) bool {
		whitespacei := whitespaceOrdered[i]
		whitespacej := whitespaceOrdered[j]
		whitespacei_order, oki := stage.WhiteSpaceMap_Staged_Order[whitespacei]
		whitespacej_order, okj := stage.WhiteSpaceMap_Staged_Order[whitespacej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return whitespacei_order < whitespacej_order
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
	if len(allOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of All instances pointers"
	}
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

		for _, _sequence := range all.Sequences {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Sequences")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Sequence_Identifiers[_sequence])
			pointersInitializesStatements += setPointerField
		}

		for _, _all := range all.Alls {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Alls")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_All_Identifiers[_all])
			pointersInitializesStatements += setPointerField
		}

		for _, _choice := range all.Choices {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Choices")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Choice_Identifiers[_choice])
			pointersInitializesStatements += setPointerField
		}

		for _, _group := range all.Groups {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Groups")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Group_Identifiers[_group])
			pointersInitializesStatements += setPointerField
		}

		for _, _element := range all.Elements {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Elements")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Element_Identifiers[_element])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(annotationOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Annotation instances pointers"
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

	if len(attributeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Attribute instances pointers"
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

	if len(attributegroupOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of AttributeGroup instances pointers"
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

		for _, _attributegroup := range attributegroup.AttributeGroups {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "AttributeGroups")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_AttributeGroup_Identifiers[_attributegroup])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute := range attributegroup.Attributes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Attributes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Attribute_Identifiers[_attribute])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(choiceOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Choice instances pointers"
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

		for _, _sequence := range choice.Sequences {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Sequences")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Sequence_Identifiers[_sequence])
			pointersInitializesStatements += setPointerField
		}

		for _, _all := range choice.Alls {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Alls")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_All_Identifiers[_all])
			pointersInitializesStatements += setPointerField
		}

		for _, _choice := range choice.Choices {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Choices")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Choice_Identifiers[_choice])
			pointersInitializesStatements += setPointerField
		}

		for _, _group := range choice.Groups {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Groups")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Group_Identifiers[_group])
			pointersInitializesStatements += setPointerField
		}

		for _, _element := range choice.Elements {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Elements")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Element_Identifiers[_element])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(complexcontentOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of ComplexContent instances pointers"
	}
	for idx, complexcontent := range complexcontentOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ComplexContent", idx, complexcontent.Name)
		map_ComplexContent_Identifiers[complexcontent] = id

		// Initialisation of values
	}

	if len(complextypeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of ComplexType instances pointers"
	}
	for idx, complextype := range complextypeOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("ComplexType", idx, complextype.Name)
		map_ComplexType_Identifiers[complextype] = id

		// Initialisation of values
		if complextype.OuterElement != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "OuterElement")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Element_Identifiers[complextype.OuterElement])
			pointersInitializesStatements += setPointerField
		}

		if complextype.Annotation != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Annotation")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Annotation_Identifiers[complextype.Annotation])
			pointersInitializesStatements += setPointerField
		}

		for _, _sequence := range complextype.Sequences {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Sequences")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Sequence_Identifiers[_sequence])
			pointersInitializesStatements += setPointerField
		}

		for _, _all := range complextype.Alls {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Alls")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_All_Identifiers[_all])
			pointersInitializesStatements += setPointerField
		}

		for _, _choice := range complextype.Choices {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Choices")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Choice_Identifiers[_choice])
			pointersInitializesStatements += setPointerField
		}

		for _, _group := range complextype.Groups {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Groups")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Group_Identifiers[_group])
			pointersInitializesStatements += setPointerField
		}

		for _, _element := range complextype.Elements {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Elements")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Element_Identifiers[_element])
			pointersInitializesStatements += setPointerField
		}

		if complextype.Extension != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Extension")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Extension_Identifiers[complextype.Extension])
			pointersInitializesStatements += setPointerField
		}

		if complextype.SimpleContent != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SimpleContent")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_SimpleContent_Identifiers[complextype.SimpleContent])
			pointersInitializesStatements += setPointerField
		}

		if complextype.ComplexContent != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "ComplexContent")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_ComplexContent_Identifiers[complextype.ComplexContent])
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

	if len(documentationOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Documentation instances pointers"
	}
	for idx, documentation := range documentationOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Documentation", idx, documentation.Name)
		map_Documentation_Identifiers[documentation] = id

		// Initialisation of values
	}

	if len(elementOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Element instances pointers"
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

		for _, _group := range element.Groups {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Groups")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Group_Identifiers[_group])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(enumerationOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Enumeration instances pointers"
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

	if len(extensionOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Extension instances pointers"
	}
	for idx, extension := range extensionOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Extension", idx, extension.Name)
		map_Extension_Identifiers[extension] = id

		// Initialisation of values
		for _, _sequence := range extension.Sequences {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Sequences")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Sequence_Identifiers[_sequence])
			pointersInitializesStatements += setPointerField
		}

		for _, _all := range extension.Alls {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Alls")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_All_Identifiers[_all])
			pointersInitializesStatements += setPointerField
		}

		for _, _choice := range extension.Choices {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Choices")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Choice_Identifiers[_choice])
			pointersInitializesStatements += setPointerField
		}

		for _, _group := range extension.Groups {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Groups")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Group_Identifiers[_group])
			pointersInitializesStatements += setPointerField
		}

		for _, _element := range extension.Elements {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Elements")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Element_Identifiers[_element])
			pointersInitializesStatements += setPointerField
		}

		for _, _attribute := range extension.Attributes {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Attributes")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Attribute_Identifiers[_attribute])
			pointersInitializesStatements += setPointerField
		}

		for _, _attributegroup := range extension.AttributeGroups {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "AttributeGroups")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_AttributeGroup_Identifiers[_attributegroup])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(groupOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Group instances pointers"
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

		if group.OuterElement != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "OuterElement")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Element_Identifiers[group.OuterElement])
			pointersInitializesStatements += setPointerField
		}

		for _, _sequence := range group.Sequences {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Sequences")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Sequence_Identifiers[_sequence])
			pointersInitializesStatements += setPointerField
		}

		for _, _all := range group.Alls {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Alls")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_All_Identifiers[_all])
			pointersInitializesStatements += setPointerField
		}

		for _, _choice := range group.Choices {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Choices")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Choice_Identifiers[_choice])
			pointersInitializesStatements += setPointerField
		}

		for _, _group := range group.Groups {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Groups")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Group_Identifiers[_group])
			pointersInitializesStatements += setPointerField
		}

		for _, _element := range group.Elements {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Elements")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Element_Identifiers[_element])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(lengthOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Length instances pointers"
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

	if len(maxinclusiveOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of MaxInclusive instances pointers"
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

	if len(maxlengthOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of MaxLength instances pointers"
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

	if len(mininclusiveOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of MinInclusive instances pointers"
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

	if len(minlengthOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of MinLength instances pointers"
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

	if len(patternOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Pattern instances pointers"
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

	if len(restrictionOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Restriction instances pointers"
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

	if len(schemaOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Schema instances pointers"
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

	if len(sequenceOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Sequence instances pointers"
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

		for _, _sequence := range sequence.Sequences {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Sequences")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Sequence_Identifiers[_sequence])
			pointersInitializesStatements += setPointerField
		}

		for _, _all := range sequence.Alls {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Alls")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_All_Identifiers[_all])
			pointersInitializesStatements += setPointerField
		}

		for _, _choice := range sequence.Choices {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Choices")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Choice_Identifiers[_choice])
			pointersInitializesStatements += setPointerField
		}

		for _, _group := range sequence.Groups {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Groups")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Group_Identifiers[_group])
			pointersInitializesStatements += setPointerField
		}

		for _, _element := range sequence.Elements {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Elements")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Element_Identifiers[_element])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(simplecontentOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of SimpleContent instances pointers"
	}
	for idx, simplecontent := range simplecontentOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("SimpleContent", idx, simplecontent.Name)
		map_SimpleContent_Identifiers[simplecontent] = id

		// Initialisation of values
		if simplecontent.Extension != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Extension")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Extension_Identifiers[simplecontent.Extension])
			pointersInitializesStatements += setPointerField
		}

		if simplecontent.Restriction != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Restriction")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Restriction_Identifiers[simplecontent.Restriction])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(simpletypeOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of SimpleType instances pointers"
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

		if simpletype.Union != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Union")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Union_Identifiers[simpletype.Union])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(totaldigitOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of TotalDigit instances pointers"
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

	if len(unionOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Union instances pointers"
	}
	for idx, union := range unionOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Union", idx, union.Name)
		map_Union_Identifiers[union] = id

		// Initialisation of values
		if union.Annotation != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Annotation")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Annotation_Identifiers[union.Annotation])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(whitespaceOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of WhiteSpace instances pointers"
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

	// Local time with timezone
	localTimestamp := stage.commitTimeStamp.Format("2006-01-02 15:04:05.000000 MST")

	// UTC time
	utcTimestamp := stage.commitTimeStamp.UTC().Format("2006-01-02 15:04:05.000000 UTC")
	res = strings.ReplaceAll(res, "{{LocalTimeStamp}}", localTimestamp)
	res = strings.ReplaceAll(res, "{{UTCTimeStamp}}", utcTimestamp)
	res = strings.ReplaceAll(res, "{{CommitId}}", fmt.Sprintf("%.10d", stage.commitId))

	if stage.MetaPackageImportAlias != "" {
		res = strings.ReplaceAll(res, "{{ImportPackageDeclaration}}",
			fmt.Sprintf("\n\t%s %s", stage.MetaPackageImportAlias, stage.MetaPackageImportPath))

		res = strings.ReplaceAll(res, "{{ImportPackageDummyDeclaration}}",
			fmt.Sprintf("\nvar _ %s.Stage",
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

	if stage.generatesDiff {
		diff := computeDiff(stage.contentWhenParsed, res)
		os.WriteFile(fmt.Sprintf("%s-%.10d-%.10d.delta", name, stage.commitIdWhenParsed, stage.commitId), []byte(diff), os.FileMode(0666))
		diff = ComputeDiff(stage.contentWhenParsed, res)
		os.WriteFile(fmt.Sprintf("%s-%.10d-%.10d.diff", name, stage.commitIdWhenParsed, stage.commitId), []byte(diff), os.FileMode(0666))
	}
	stage.contentWhenParsed = res
	stage.commitIdWhenParsed = stage.commitId

	fmt.Fprintln(file, res)
}

// computeDiff calculates the git-style unified diff between two strings.
func computeDiff(a, b string) string {
	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(a, b, false)
	return dmp.DiffToDelta(diffs)
}

// computePrettyDiff calculates the git-style unified diff between two strings.
func computePrettyDiff(a, b string) string {
	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(a, b, false)
	return dmp.DiffPrettyHtml(diffs)
}

// applyDiff reconstructs the original string 'a' from the new string 'b' and the diff string 'c'.
func applyDiff(b, c string) (string, error) {
	dmp := diffmatchpatch.New()
	diffs, err := dmp.DiffFromDelta(b, c)
	if err != nil {
		return "", err
	}
	patches := dmp.PatchMake(b, diffs)
	// We are applying the patch in reverse to get from 'b' to 'a'.
	// The library's PatchApply function returns the new string and a slice of booleans indicating the success of each patch application.
	result, _ := dmp.PatchApply(patches, b)
	return result, nil
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
