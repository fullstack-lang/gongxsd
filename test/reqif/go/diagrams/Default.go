package diagrams

import (
	"time"

	"github.com/fullstack-lang/gongdoc/go/models"

	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gongxsd/test/reqif/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ time.Time

// _ point for meta package dummy declaration
var _ ref_models.StageStruct

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.StageStruct) {

	// Declaration of instances to stage

	__Classdiagram__000000_Default := (&models.Classdiagram{Name: `Default`}).Stage(stage)

	__Field__000000_ATTRIBUTE_DEFINITION_XHTML_REF := (&models.Field{Name: `ATTRIBUTE_DEFINITION_XHTML_REF`}).Stage(stage)
	__Field__000001_DESC := (&models.Field{Name: `DESC`}).Stage(stage)
	__Field__000002_IDENTIFIER := (&models.Field{Name: `IDENTIFIER`}).Stage(stage)
	__Field__000003_IS_SIMPLIFIED := (&models.Field{Name: `IS_SIMPLIFIED`}).Stage(stage)
	__Field__000004_LAST_CHANGE := (&models.Field{Name: `LAST_CHANGE`}).Stage(stage)
	__Field__000005_LONG_NAME := (&models.Field{Name: `LONG_NAME`}).Stage(stage)
	__Field__000006_THE_VALUE := (&models.Field{Name: `THE_VALUE`}).Stage(stage)

	__GongStructShape__000000_Default_ATTRIBUTE_DEFINITION_BOOLEAN := (&models.GongStructShape{Name: `Default-ATTRIBUTE_DEFINITION_BOOLEAN`}).Stage(stage)
	__GongStructShape__000001_Default_ATTRIBUTE_DEFINITION_DATE := (&models.GongStructShape{Name: `Default-ATTRIBUTE_DEFINITION_DATE`}).Stage(stage)
	__GongStructShape__000002_Default_ATTRIBUTE_DEFINITION_ENUMERATION := (&models.GongStructShape{Name: `Default-ATTRIBUTE_DEFINITION_ENUMERATION`}).Stage(stage)
	__GongStructShape__000003_Default_ATTRIBUTE_DEFINITION_INTEGER := (&models.GongStructShape{Name: `Default-ATTRIBUTE_DEFINITION_INTEGER`}).Stage(stage)
	__GongStructShape__000004_Default_ATTRIBUTE_DEFINITION_REAL := (&models.GongStructShape{Name: `Default-ATTRIBUTE_DEFINITION_REAL`}).Stage(stage)
	__GongStructShape__000005_Default_ATTRIBUTE_DEFINITION_STRING := (&models.GongStructShape{Name: `Default-ATTRIBUTE_DEFINITION_STRING`}).Stage(stage)
	__GongStructShape__000006_Default_ATTRIBUTE_DEFINITION_XHTML := (&models.GongStructShape{Name: `Default-ATTRIBUTE_DEFINITION_XHTML`}).Stage(stage)
	__GongStructShape__000007_Default_ATTRIBUTE_VALUE_BOOLEAN := (&models.GongStructShape{Name: `Default-ATTRIBUTE_VALUE_BOOLEAN`}).Stage(stage)
	__GongStructShape__000008_Default_ATTRIBUTE_VALUE_DATE := (&models.GongStructShape{Name: `Default-ATTRIBUTE_VALUE_DATE`}).Stage(stage)
	__GongStructShape__000009_Default_ATTRIBUTE_VALUE_ENUMERATION := (&models.GongStructShape{Name: `Default-ATTRIBUTE_VALUE_ENUMERATION`}).Stage(stage)
	__GongStructShape__000010_Default_ATTRIBUTE_VALUE_INTEGER := (&models.GongStructShape{Name: `Default-ATTRIBUTE_VALUE_INTEGER`}).Stage(stage)
	__GongStructShape__000011_Default_ATTRIBUTE_VALUE_REAL := (&models.GongStructShape{Name: `Default-ATTRIBUTE_VALUE_REAL`}).Stage(stage)
	__GongStructShape__000012_Default_ATTRIBUTE_VALUE_STRING := (&models.GongStructShape{Name: `Default-ATTRIBUTE_VALUE_STRING`}).Stage(stage)
	__GongStructShape__000013_Default_ATTRIBUTE_VALUE_XHTML := (&models.GongStructShape{Name: `Default-ATTRIBUTE_VALUE_XHTML`}).Stage(stage)
	__GongStructShape__000014_Default_A_ATTRIBUTE_DEFINITION_XHTML_REF := (&models.GongStructShape{Name: `Default-A_ATTRIBUTE_DEFINITION_XHTML_REF`}).Stage(stage)
	__GongStructShape__000015_Default_A_ATTRIBUTE_VALUE_XHTML := (&models.GongStructShape{Name: `Default-A_ATTRIBUTE_VALUE_XHTML`}).Stage(stage)
	__GongStructShape__000016_Default_A_ATTRIBUTE_VALUE_XHTML_1 := (&models.GongStructShape{Name: `Default-A_ATTRIBUTE_VALUE_XHTML_1`}).Stage(stage)
	__GongStructShape__000017_Default_A_CORE_CONTENT := (&models.GongStructShape{Name: `Default-A_CORE_CONTENT`}).Stage(stage)
	__GongStructShape__000018_Default_A_SPECIFICATIONS := (&models.GongStructShape{Name: `Default-A_SPECIFICATIONS`}).Stage(stage)
	__GongStructShape__000019_Default_A_SPEC_OBJECTS := (&models.GongStructShape{Name: `Default-A_SPEC_OBJECTS`}).Stage(stage)
	__GongStructShape__000020_Default_A_SPEC_OBJECT_TYPE_REF := (&models.GongStructShape{Name: `Default-A_SPEC_OBJECT_TYPE_REF`}).Stage(stage)
	__GongStructShape__000021_Default_A_SPEC_RELATIONS := (&models.GongStructShape{Name: `Default-A_SPEC_RELATIONS`}).Stage(stage)
	__GongStructShape__000022_Default_A_SPEC_RELATION_TYPE_REF := (&models.GongStructShape{Name: `Default-A_SPEC_RELATION_TYPE_REF`}).Stage(stage)
	__GongStructShape__000023_Default_A_THE_HEADER := (&models.GongStructShape{Name: `Default-A_THE_HEADER`}).Stage(stage)
	__GongStructShape__000024_Default_REQ_IF := (&models.GongStructShape{Name: `Default-REQ_IF`}).Stage(stage)
	__GongStructShape__000025_Default_REQ_IF_CONTENT := (&models.GongStructShape{Name: `Default-REQ_IF_CONTENT`}).Stage(stage)
	__GongStructShape__000026_Default_REQ_IF_HEADER := (&models.GongStructShape{Name: `Default-REQ_IF_HEADER`}).Stage(stage)
	__GongStructShape__000027_Default_SPEC_OBJECT := (&models.GongStructShape{Name: `Default-SPEC_OBJECT`}).Stage(stage)
	__GongStructShape__000028_Default_SPEC_RELATION := (&models.GongStructShape{Name: `Default-SPEC_RELATION`}).Stage(stage)
	__GongStructShape__000029_Default_XHTML_CONTENT := (&models.GongStructShape{Name: `Default-XHTML_CONTENT`}).Stage(stage)

	__Link__000000_ATTRIBUTE_VALUE_BOOLEAN := (&models.Link{Name: `ATTRIBUTE_VALUE_BOOLEAN`}).Stage(stage)
	__Link__000001_ATTRIBUTE_VALUE_DATE := (&models.Link{Name: `ATTRIBUTE_VALUE_DATE`}).Stage(stage)
	__Link__000002_ATTRIBUTE_VALUE_ENUMERATION := (&models.Link{Name: `ATTRIBUTE_VALUE_ENUMERATION`}).Stage(stage)
	__Link__000003_ATTRIBUTE_VALUE_INTEGER := (&models.Link{Name: `ATTRIBUTE_VALUE_INTEGER`}).Stage(stage)
	__Link__000004_ATTRIBUTE_VALUE_REAL := (&models.Link{Name: `ATTRIBUTE_VALUE_REAL`}).Stage(stage)
	__Link__000005_ATTRIBUTE_VALUE_STRING := (&models.Link{Name: `ATTRIBUTE_VALUE_STRING`}).Stage(stage)
	__Link__000006_ATTRIBUTE_VALUE_XHTML := (&models.Link{Name: `ATTRIBUTE_VALUE_XHTML`}).Stage(stage)
	__Link__000007_CORE_CONTENT := (&models.Link{Name: `CORE_CONTENT`}).Stage(stage)
	__Link__000008_DEFINITION := (&models.Link{Name: `DEFINITION`}).Stage(stage)
	__Link__000009_REQ_IF_CONTENT := (&models.Link{Name: `REQ_IF_CONTENT`}).Stage(stage)
	__Link__000010_REQ_IF_HEADER := (&models.Link{Name: `REQ_IF_HEADER`}).Stage(stage)
	__Link__000011_SPECIFICATIONS := (&models.Link{Name: `SPECIFICATIONS`}).Stage(stage)
	__Link__000012_SPEC_OBJECT := (&models.Link{Name: `SPEC_OBJECT`}).Stage(stage)
	__Link__000013_SPEC_OBJECTS := (&models.Link{Name: `SPEC_OBJECTS`}).Stage(stage)
	__Link__000014_SPEC_RELATIONS := (&models.Link{Name: `SPEC_RELATIONS`}).Stage(stage)
	__Link__000015_THE_HEADER := (&models.Link{Name: `THE_HEADER`}).Stage(stage)
	__Link__000016_THE_ORIGINAL_VALUE := (&models.Link{Name: `THE_ORIGINAL_VALUE`}).Stage(stage)
	__Link__000017_THE_VALUE := (&models.Link{Name: `THE_VALUE`}).Stage(stage)
	__Link__000018_TYPE := (&models.Link{Name: `TYPE`}).Stage(stage)
	__Link__000019_VALUES := (&models.Link{Name: `VALUES`}).Stage(stage)

	__Position__000000_Pos_Default_ATTRIBUTE_DEFINITION_BOOLEAN := (&models.Position{Name: `Pos-Default-ATTRIBUTE_DEFINITION_BOOLEAN`}).Stage(stage)
	__Position__000001_Pos_Default_ATTRIBUTE_DEFINITION_DATE := (&models.Position{Name: `Pos-Default-ATTRIBUTE_DEFINITION_DATE`}).Stage(stage)
	__Position__000002_Pos_Default_ATTRIBUTE_DEFINITION_ENUMERATION := (&models.Position{Name: `Pos-Default-ATTRIBUTE_DEFINITION_ENUMERATION`}).Stage(stage)
	__Position__000003_Pos_Default_ATTRIBUTE_DEFINITION_INTEGER := (&models.Position{Name: `Pos-Default-ATTRIBUTE_DEFINITION_INTEGER`}).Stage(stage)
	__Position__000004_Pos_Default_ATTRIBUTE_DEFINITION_REAL := (&models.Position{Name: `Pos-Default-ATTRIBUTE_DEFINITION_REAL`}).Stage(stage)
	__Position__000005_Pos_Default_ATTRIBUTE_DEFINITION_STRING := (&models.Position{Name: `Pos-Default-ATTRIBUTE_DEFINITION_STRING`}).Stage(stage)
	__Position__000006_Pos_Default_ATTRIBUTE_DEFINITION_XHTML := (&models.Position{Name: `Pos-Default-ATTRIBUTE_DEFINITION_XHTML`}).Stage(stage)
	__Position__000007_Pos_Default_ATTRIBUTE_VALUE_BOOLEAN := (&models.Position{Name: `Pos-Default-ATTRIBUTE_VALUE_BOOLEAN`}).Stage(stage)
	__Position__000008_Pos_Default_ATTRIBUTE_VALUE_DATE := (&models.Position{Name: `Pos-Default-ATTRIBUTE_VALUE_DATE`}).Stage(stage)
	__Position__000009_Pos_Default_ATTRIBUTE_VALUE_ENUMERATION := (&models.Position{Name: `Pos-Default-ATTRIBUTE_VALUE_ENUMERATION`}).Stage(stage)
	__Position__000010_Pos_Default_ATTRIBUTE_VALUE_INTEGER := (&models.Position{Name: `Pos-Default-ATTRIBUTE_VALUE_INTEGER`}).Stage(stage)
	__Position__000011_Pos_Default_ATTRIBUTE_VALUE_REAL := (&models.Position{Name: `Pos-Default-ATTRIBUTE_VALUE_REAL`}).Stage(stage)
	__Position__000012_Pos_Default_ATTRIBUTE_VALUE_STRING := (&models.Position{Name: `Pos-Default-ATTRIBUTE_VALUE_STRING`}).Stage(stage)
	__Position__000013_Pos_Default_ATTRIBUTE_VALUE_XHTML := (&models.Position{Name: `Pos-Default-ATTRIBUTE_VALUE_XHTML`}).Stage(stage)
	__Position__000014_Pos_Default_A_ATTRIBUTE_DEFINITION_XHTML_REF := (&models.Position{Name: `Pos-Default-A_ATTRIBUTE_DEFINITION_XHTML_REF`}).Stage(stage)
	__Position__000015_Pos_Default_A_ATTRIBUTE_VALUE_XHTML := (&models.Position{Name: `Pos-Default-A_ATTRIBUTE_VALUE_XHTML`}).Stage(stage)
	__Position__000016_Pos_Default_A_ATTRIBUTE_VALUE_XHTML_1 := (&models.Position{Name: `Pos-Default-A_ATTRIBUTE_VALUE_XHTML_1`}).Stage(stage)
	__Position__000017_Pos_Default_A_CORE_CONTENT := (&models.Position{Name: `Pos-Default-A_CORE_CONTENT`}).Stage(stage)
	__Position__000018_Pos_Default_A_SPECIFICATIONS := (&models.Position{Name: `Pos-Default-A_SPECIFICATIONS`}).Stage(stage)
	__Position__000019_Pos_Default_A_SPEC_OBJECTS := (&models.Position{Name: `Pos-Default-A_SPEC_OBJECTS`}).Stage(stage)
	__Position__000020_Pos_Default_A_SPEC_OBJECT_TYPE_REF := (&models.Position{Name: `Pos-Default-A_SPEC_OBJECT_TYPE_REF`}).Stage(stage)
	__Position__000021_Pos_Default_A_SPEC_RELATIONS := (&models.Position{Name: `Pos-Default-A_SPEC_RELATIONS`}).Stage(stage)
	__Position__000022_Pos_Default_A_SPEC_RELATION_TYPE_REF := (&models.Position{Name: `Pos-Default-A_SPEC_RELATION_TYPE_REF`}).Stage(stage)
	__Position__000023_Pos_Default_A_THE_HEADER := (&models.Position{Name: `Pos-Default-A_THE_HEADER`}).Stage(stage)
	__Position__000024_Pos_Default_REQ_IF := (&models.Position{Name: `Pos-Default-REQ_IF`}).Stage(stage)
	__Position__000025_Pos_Default_REQ_IF_CONTENT := (&models.Position{Name: `Pos-Default-REQ_IF_CONTENT`}).Stage(stage)
	__Position__000026_Pos_Default_REQ_IF_HEADER := (&models.Position{Name: `Pos-Default-REQ_IF_HEADER`}).Stage(stage)
	__Position__000027_Pos_Default_SPEC_OBJECT := (&models.Position{Name: `Pos-Default-SPEC_OBJECT`}).Stage(stage)
	__Position__000028_Pos_Default_SPEC_RELATION := (&models.Position{Name: `Pos-Default-SPEC_RELATION`}).Stage(stage)
	__Position__000029_Pos_Default_XHTML_CONTENT := (&models.Position{Name: `Pos-Default-XHTML_CONTENT`}).Stage(stage)

	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_ATTRIBUTE_VALUE_XHTML_and_Default_A_ATTRIBUTE_DEFINITION_XHTML_REF := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-ATTRIBUTE_VALUE_XHTML and Default-A_ATTRIBUTE_DEFINITION_XHTML_REF`}).Stage(stage)
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_ATTRIBUTE_VALUE_XHTML_and_Default_XHTML_CONTENT := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-ATTRIBUTE_VALUE_XHTML and Default-XHTML_CONTENT`}).Stage(stage)
	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_ATTRIBUTE_VALUE_XHTML_and_Default_XHTML_CONTENT := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-ATTRIBUTE_VALUE_XHTML and Default-XHTML_CONTENT`}).Stage(stage)
	__Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_A_ATTRIBUTE_VALUE_XHTML_1_and_Default_ATTRIBUTE_VALUE_BOOLEAN := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-A_ATTRIBUTE_VALUE_XHTML_1 and Default-ATTRIBUTE_VALUE_BOOLEAN`}).Stage(stage)
	__Vertice__000004_Verticle_in_class_diagram_Default_in_middle_between_Default_A_ATTRIBUTE_VALUE_XHTML_1_and_Default_ATTRIBUTE_VALUE_DATE := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-A_ATTRIBUTE_VALUE_XHTML_1 and Default-ATTRIBUTE_VALUE_DATE`}).Stage(stage)
	__Vertice__000005_Verticle_in_class_diagram_Default_in_middle_between_Default_A_ATTRIBUTE_VALUE_XHTML_1_and_Default_ATTRIBUTE_VALUE_ENUMERATION := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-A_ATTRIBUTE_VALUE_XHTML_1 and Default-ATTRIBUTE_VALUE_ENUMERATION`}).Stage(stage)
	__Vertice__000006_Verticle_in_class_diagram_Default_in_middle_between_Default_A_ATTRIBUTE_VALUE_XHTML_1_and_Default_ATTRIBUTE_VALUE_INTEGER := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-A_ATTRIBUTE_VALUE_XHTML_1 and Default-ATTRIBUTE_VALUE_INTEGER`}).Stage(stage)
	__Vertice__000007_Verticle_in_class_diagram_Default_in_middle_between_Default_A_ATTRIBUTE_VALUE_XHTML_1_and_Default_ATTRIBUTE_VALUE_REAL := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-A_ATTRIBUTE_VALUE_XHTML_1 and Default-ATTRIBUTE_VALUE_REAL`}).Stage(stage)
	__Vertice__000008_Verticle_in_class_diagram_Default_in_middle_between_Default_A_ATTRIBUTE_VALUE_XHTML_1_and_Default_ATTRIBUTE_VALUE_STRING := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-A_ATTRIBUTE_VALUE_XHTML_1 and Default-ATTRIBUTE_VALUE_STRING`}).Stage(stage)
	__Vertice__000009_Verticle_in_class_diagram_Default_in_middle_between_Default_A_ATTRIBUTE_VALUE_XHTML_1_and_Default_ATTRIBUTE_VALUE_XHTML := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-A_ATTRIBUTE_VALUE_XHTML_1 and Default-ATTRIBUTE_VALUE_XHTML`}).Stage(stage)
	__Vertice__000010_Verticle_in_class_diagram_Default_in_middle_between_Default_A_CORE_CONTENT_and_Default_REQ_IF_CONTENT := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-A_CORE_CONTENT and Default-REQ_IF_CONTENT`}).Stage(stage)
	__Vertice__000011_Verticle_in_class_diagram_Default_in_middle_between_Default_A_SPEC_OBJECTS_and_Default_SPEC_OBJECT := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-A_SPEC_OBJECTS and Default-SPEC_OBJECT`}).Stage(stage)
	__Vertice__000012_Verticle_in_class_diagram_Default_in_middle_between_Default_A_THE_HEADER_and_Default_REQ_IF_HEADER := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-A_THE_HEADER and Default-REQ_IF_HEADER`}).Stage(stage)
	__Vertice__000013_Verticle_in_class_diagram_Default_in_middle_between_Default_REQ_IF_and_Default_A_CORE_CONTENT := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-REQ_IF and Default-A_CORE_CONTENT`}).Stage(stage)
	__Vertice__000014_Verticle_in_class_diagram_Default_in_middle_between_Default_REQ_IF_and_Default_A_THE_HEADER := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-REQ_IF and Default-A_THE_HEADER`}).Stage(stage)
	__Vertice__000015_Verticle_in_class_diagram_Default_in_middle_between_Default_REQ_IF_CONTENT_and_Default_A_SPECIFICATIONS := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-REQ_IF_CONTENT and Default-A_SPECIFICATIONS`}).Stage(stage)
	__Vertice__000016_Verticle_in_class_diagram_Default_in_middle_between_Default_REQ_IF_CONTENT_and_Default_A_SPEC_OBJECTS := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-REQ_IF_CONTENT and Default-A_SPEC_OBJECTS`}).Stage(stage)
	__Vertice__000017_Verticle_in_class_diagram_Default_in_middle_between_Default_REQ_IF_CONTENT_and_Default_A_SPEC_RELATIONS := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-REQ_IF_CONTENT and Default-A_SPEC_RELATIONS`}).Stage(stage)
	__Vertice__000018_Verticle_in_class_diagram_Default_in_middle_between_Default_SPEC_OBJECT_and_Default_A_ATTRIBUTE_VALUE_XHTML_1 := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-SPEC_OBJECT and Default-A_ATTRIBUTE_VALUE_XHTML_1`}).Stage(stage)
	__Vertice__000019_Verticle_in_class_diagram_Default_in_middle_between_Default_SPEC_OBJECT_and_Default_A_SPEC_OBJECT_TYPE_REF := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-SPEC_OBJECT and Default-A_SPEC_OBJECT_TYPE_REF`}).Stage(stage)

	// Setup of values

	__Classdiagram__000000_Default.Name = `Default`
	__Classdiagram__000000_Default.IsInDrawMode = false

	__Field__000000_ATTRIBUTE_DEFINITION_XHTML_REF.Name = `ATTRIBUTE_DEFINITION_XHTML_REF`

	//gong:ident [ref_models.A_ATTRIBUTE_DEFINITION_XHTML_REF.ATTRIBUTE_DEFINITION_XHTML_REF] comment added to overcome the problem with the comment map association
	__Field__000000_ATTRIBUTE_DEFINITION_XHTML_REF.Identifier = `ref_models.A_ATTRIBUTE_DEFINITION_XHTML_REF.ATTRIBUTE_DEFINITION_XHTML_REF`
	__Field__000000_ATTRIBUTE_DEFINITION_XHTML_REF.FieldTypeAsString = ``
	__Field__000000_ATTRIBUTE_DEFINITION_XHTML_REF.Structname = `A_ATTRIBUTE_DEFINITION_XHTML_REF`
	__Field__000000_ATTRIBUTE_DEFINITION_XHTML_REF.Fieldtypename = `string`

	__Field__000001_DESC.Name = `DESC`

	//gong:ident [ref_models.SPEC_OBJECT.DESC] comment added to overcome the problem with the comment map association
	__Field__000001_DESC.Identifier = `ref_models.SPEC_OBJECT.DESC`
	__Field__000001_DESC.FieldTypeAsString = ``
	__Field__000001_DESC.Structname = `SPEC_OBJECT`
	__Field__000001_DESC.Fieldtypename = `string`

	__Field__000002_IDENTIFIER.Name = `IDENTIFIER`

	//gong:ident [ref_models.SPEC_OBJECT.IDENTIFIER] comment added to overcome the problem with the comment map association
	__Field__000002_IDENTIFIER.Identifier = `ref_models.SPEC_OBJECT.IDENTIFIER`
	__Field__000002_IDENTIFIER.FieldTypeAsString = ``
	__Field__000002_IDENTIFIER.Structname = `SPEC_OBJECT`
	__Field__000002_IDENTIFIER.Fieldtypename = `string`

	__Field__000003_IS_SIMPLIFIED.Name = `IS_SIMPLIFIED`

	//gong:ident [ref_models.ATTRIBUTE_VALUE_XHTML.IS_SIMPLIFIED] comment added to overcome the problem with the comment map association
	__Field__000003_IS_SIMPLIFIED.Identifier = `ref_models.ATTRIBUTE_VALUE_XHTML.IS_SIMPLIFIED`
	__Field__000003_IS_SIMPLIFIED.FieldTypeAsString = ``
	__Field__000003_IS_SIMPLIFIED.Structname = `ATTRIBUTE_VALUE_XHTML`
	__Field__000003_IS_SIMPLIFIED.Fieldtypename = `bool`

	__Field__000004_LAST_CHANGE.Name = `LAST_CHANGE`

	//gong:ident [ref_models.SPEC_OBJECT.LAST_CHANGE] comment added to overcome the problem with the comment map association
	__Field__000004_LAST_CHANGE.Identifier = `ref_models.SPEC_OBJECT.LAST_CHANGE`
	__Field__000004_LAST_CHANGE.FieldTypeAsString = ``
	__Field__000004_LAST_CHANGE.Structname = `SPEC_OBJECT`
	__Field__000004_LAST_CHANGE.Fieldtypename = `string`

	__Field__000005_LONG_NAME.Name = `LONG_NAME`

	//gong:ident [ref_models.SPEC_OBJECT.LONG_NAME] comment added to overcome the problem with the comment map association
	__Field__000005_LONG_NAME.Identifier = `ref_models.SPEC_OBJECT.LONG_NAME`
	__Field__000005_LONG_NAME.FieldTypeAsString = ``
	__Field__000005_LONG_NAME.Structname = `SPEC_OBJECT`
	__Field__000005_LONG_NAME.Fieldtypename = `string`

	__Field__000006_THE_VALUE.Name = `THE_VALUE`

	//gong:ident [ref_models.ATTRIBUTE_VALUE_STRING.THE_VALUE] comment added to overcome the problem with the comment map association
	__Field__000006_THE_VALUE.Identifier = `ref_models.ATTRIBUTE_VALUE_STRING.THE_VALUE`
	__Field__000006_THE_VALUE.FieldTypeAsString = ``
	__Field__000006_THE_VALUE.Structname = `ATTRIBUTE_VALUE_STRING`
	__Field__000006_THE_VALUE.Fieldtypename = `string`

	__GongStructShape__000000_Default_ATTRIBUTE_DEFINITION_BOOLEAN.Name = `Default-ATTRIBUTE_DEFINITION_BOOLEAN`

	//gong:ident [ref_models.ATTRIBUTE_DEFINITION_BOOLEAN] comment added to overcome the problem with the comment map association
	__GongStructShape__000000_Default_ATTRIBUTE_DEFINITION_BOOLEAN.Identifier = `ref_models.ATTRIBUTE_DEFINITION_BOOLEAN`
	__GongStructShape__000000_Default_ATTRIBUTE_DEFINITION_BOOLEAN.ShowNbInstances = false
	__GongStructShape__000000_Default_ATTRIBUTE_DEFINITION_BOOLEAN.NbInstances = 0
	__GongStructShape__000000_Default_ATTRIBUTE_DEFINITION_BOOLEAN.Width = 240.000000
	__GongStructShape__000000_Default_ATTRIBUTE_DEFINITION_BOOLEAN.Height = 63.000000
	__GongStructShape__000000_Default_ATTRIBUTE_DEFINITION_BOOLEAN.IsSelected = false

	__GongStructShape__000001_Default_ATTRIBUTE_DEFINITION_DATE.Name = `Default-ATTRIBUTE_DEFINITION_DATE`

	//gong:ident [ref_models.ATTRIBUTE_DEFINITION_DATE] comment added to overcome the problem with the comment map association
	__GongStructShape__000001_Default_ATTRIBUTE_DEFINITION_DATE.Identifier = `ref_models.ATTRIBUTE_DEFINITION_DATE`
	__GongStructShape__000001_Default_ATTRIBUTE_DEFINITION_DATE.ShowNbInstances = false
	__GongStructShape__000001_Default_ATTRIBUTE_DEFINITION_DATE.NbInstances = 0
	__GongStructShape__000001_Default_ATTRIBUTE_DEFINITION_DATE.Width = 240.000000
	__GongStructShape__000001_Default_ATTRIBUTE_DEFINITION_DATE.Height = 63.000000
	__GongStructShape__000001_Default_ATTRIBUTE_DEFINITION_DATE.IsSelected = false

	__GongStructShape__000002_Default_ATTRIBUTE_DEFINITION_ENUMERATION.Name = `Default-ATTRIBUTE_DEFINITION_ENUMERATION`

	//gong:ident [ref_models.ATTRIBUTE_DEFINITION_ENUMERATION] comment added to overcome the problem with the comment map association
	__GongStructShape__000002_Default_ATTRIBUTE_DEFINITION_ENUMERATION.Identifier = `ref_models.ATTRIBUTE_DEFINITION_ENUMERATION`
	__GongStructShape__000002_Default_ATTRIBUTE_DEFINITION_ENUMERATION.ShowNbInstances = false
	__GongStructShape__000002_Default_ATTRIBUTE_DEFINITION_ENUMERATION.NbInstances = 0
	__GongStructShape__000002_Default_ATTRIBUTE_DEFINITION_ENUMERATION.Width = 240.000000
	__GongStructShape__000002_Default_ATTRIBUTE_DEFINITION_ENUMERATION.Height = 63.000000
	__GongStructShape__000002_Default_ATTRIBUTE_DEFINITION_ENUMERATION.IsSelected = false

	__GongStructShape__000003_Default_ATTRIBUTE_DEFINITION_INTEGER.Name = `Default-ATTRIBUTE_DEFINITION_INTEGER`

	//gong:ident [ref_models.ATTRIBUTE_DEFINITION_INTEGER] comment added to overcome the problem with the comment map association
	__GongStructShape__000003_Default_ATTRIBUTE_DEFINITION_INTEGER.Identifier = `ref_models.ATTRIBUTE_DEFINITION_INTEGER`
	__GongStructShape__000003_Default_ATTRIBUTE_DEFINITION_INTEGER.ShowNbInstances = false
	__GongStructShape__000003_Default_ATTRIBUTE_DEFINITION_INTEGER.NbInstances = 0
	__GongStructShape__000003_Default_ATTRIBUTE_DEFINITION_INTEGER.Width = 240.000000
	__GongStructShape__000003_Default_ATTRIBUTE_DEFINITION_INTEGER.Height = 63.000000
	__GongStructShape__000003_Default_ATTRIBUTE_DEFINITION_INTEGER.IsSelected = false

	__GongStructShape__000004_Default_ATTRIBUTE_DEFINITION_REAL.Name = `Default-ATTRIBUTE_DEFINITION_REAL`

	//gong:ident [ref_models.ATTRIBUTE_DEFINITION_REAL] comment added to overcome the problem with the comment map association
	__GongStructShape__000004_Default_ATTRIBUTE_DEFINITION_REAL.Identifier = `ref_models.ATTRIBUTE_DEFINITION_REAL`
	__GongStructShape__000004_Default_ATTRIBUTE_DEFINITION_REAL.ShowNbInstances = false
	__GongStructShape__000004_Default_ATTRIBUTE_DEFINITION_REAL.NbInstances = 0
	__GongStructShape__000004_Default_ATTRIBUTE_DEFINITION_REAL.Width = 240.000000
	__GongStructShape__000004_Default_ATTRIBUTE_DEFINITION_REAL.Height = 63.000000
	__GongStructShape__000004_Default_ATTRIBUTE_DEFINITION_REAL.IsSelected = false

	__GongStructShape__000005_Default_ATTRIBUTE_DEFINITION_STRING.Name = `Default-ATTRIBUTE_DEFINITION_STRING`

	//gong:ident [ref_models.ATTRIBUTE_DEFINITION_STRING] comment added to overcome the problem with the comment map association
	__GongStructShape__000005_Default_ATTRIBUTE_DEFINITION_STRING.Identifier = `ref_models.ATTRIBUTE_DEFINITION_STRING`
	__GongStructShape__000005_Default_ATTRIBUTE_DEFINITION_STRING.ShowNbInstances = false
	__GongStructShape__000005_Default_ATTRIBUTE_DEFINITION_STRING.NbInstances = 0
	__GongStructShape__000005_Default_ATTRIBUTE_DEFINITION_STRING.Width = 240.000000
	__GongStructShape__000005_Default_ATTRIBUTE_DEFINITION_STRING.Height = 63.000000
	__GongStructShape__000005_Default_ATTRIBUTE_DEFINITION_STRING.IsSelected = false

	__GongStructShape__000006_Default_ATTRIBUTE_DEFINITION_XHTML.Name = `Default-ATTRIBUTE_DEFINITION_XHTML`

	//gong:ident [ref_models.ATTRIBUTE_DEFINITION_XHTML] comment added to overcome the problem with the comment map association
	__GongStructShape__000006_Default_ATTRIBUTE_DEFINITION_XHTML.Identifier = `ref_models.ATTRIBUTE_DEFINITION_XHTML`
	__GongStructShape__000006_Default_ATTRIBUTE_DEFINITION_XHTML.ShowNbInstances = false
	__GongStructShape__000006_Default_ATTRIBUTE_DEFINITION_XHTML.NbInstances = 0
	__GongStructShape__000006_Default_ATTRIBUTE_DEFINITION_XHTML.Width = 240.000000
	__GongStructShape__000006_Default_ATTRIBUTE_DEFINITION_XHTML.Height = 63.000000
	__GongStructShape__000006_Default_ATTRIBUTE_DEFINITION_XHTML.IsSelected = false

	__GongStructShape__000007_Default_ATTRIBUTE_VALUE_BOOLEAN.Name = `Default-ATTRIBUTE_VALUE_BOOLEAN`

	//gong:ident [ref_models.ATTRIBUTE_VALUE_BOOLEAN] comment added to overcome the problem with the comment map association
	__GongStructShape__000007_Default_ATTRIBUTE_VALUE_BOOLEAN.Identifier = `ref_models.ATTRIBUTE_VALUE_BOOLEAN`
	__GongStructShape__000007_Default_ATTRIBUTE_VALUE_BOOLEAN.ShowNbInstances = false
	__GongStructShape__000007_Default_ATTRIBUTE_VALUE_BOOLEAN.NbInstances = 0
	__GongStructShape__000007_Default_ATTRIBUTE_VALUE_BOOLEAN.Width = 240.000000
	__GongStructShape__000007_Default_ATTRIBUTE_VALUE_BOOLEAN.Height = 63.000000
	__GongStructShape__000007_Default_ATTRIBUTE_VALUE_BOOLEAN.IsSelected = false

	__GongStructShape__000008_Default_ATTRIBUTE_VALUE_DATE.Name = `Default-ATTRIBUTE_VALUE_DATE`

	//gong:ident [ref_models.ATTRIBUTE_VALUE_DATE] comment added to overcome the problem with the comment map association
	__GongStructShape__000008_Default_ATTRIBUTE_VALUE_DATE.Identifier = `ref_models.ATTRIBUTE_VALUE_DATE`
	__GongStructShape__000008_Default_ATTRIBUTE_VALUE_DATE.ShowNbInstances = false
	__GongStructShape__000008_Default_ATTRIBUTE_VALUE_DATE.NbInstances = 0
	__GongStructShape__000008_Default_ATTRIBUTE_VALUE_DATE.Width = 240.000000
	__GongStructShape__000008_Default_ATTRIBUTE_VALUE_DATE.Height = 63.000000
	__GongStructShape__000008_Default_ATTRIBUTE_VALUE_DATE.IsSelected = false

	__GongStructShape__000009_Default_ATTRIBUTE_VALUE_ENUMERATION.Name = `Default-ATTRIBUTE_VALUE_ENUMERATION`

	//gong:ident [ref_models.ATTRIBUTE_VALUE_ENUMERATION] comment added to overcome the problem with the comment map association
	__GongStructShape__000009_Default_ATTRIBUTE_VALUE_ENUMERATION.Identifier = `ref_models.ATTRIBUTE_VALUE_ENUMERATION`
	__GongStructShape__000009_Default_ATTRIBUTE_VALUE_ENUMERATION.ShowNbInstances = false
	__GongStructShape__000009_Default_ATTRIBUTE_VALUE_ENUMERATION.NbInstances = 0
	__GongStructShape__000009_Default_ATTRIBUTE_VALUE_ENUMERATION.Width = 240.000000
	__GongStructShape__000009_Default_ATTRIBUTE_VALUE_ENUMERATION.Height = 63.000000
	__GongStructShape__000009_Default_ATTRIBUTE_VALUE_ENUMERATION.IsSelected = false

	__GongStructShape__000010_Default_ATTRIBUTE_VALUE_INTEGER.Name = `Default-ATTRIBUTE_VALUE_INTEGER`

	//gong:ident [ref_models.ATTRIBUTE_VALUE_INTEGER] comment added to overcome the problem with the comment map association
	__GongStructShape__000010_Default_ATTRIBUTE_VALUE_INTEGER.Identifier = `ref_models.ATTRIBUTE_VALUE_INTEGER`
	__GongStructShape__000010_Default_ATTRIBUTE_VALUE_INTEGER.ShowNbInstances = false
	__GongStructShape__000010_Default_ATTRIBUTE_VALUE_INTEGER.NbInstances = 0
	__GongStructShape__000010_Default_ATTRIBUTE_VALUE_INTEGER.Width = 240.000000
	__GongStructShape__000010_Default_ATTRIBUTE_VALUE_INTEGER.Height = 63.000000
	__GongStructShape__000010_Default_ATTRIBUTE_VALUE_INTEGER.IsSelected = false

	__GongStructShape__000011_Default_ATTRIBUTE_VALUE_REAL.Name = `Default-ATTRIBUTE_VALUE_REAL`

	//gong:ident [ref_models.ATTRIBUTE_VALUE_REAL] comment added to overcome the problem with the comment map association
	__GongStructShape__000011_Default_ATTRIBUTE_VALUE_REAL.Identifier = `ref_models.ATTRIBUTE_VALUE_REAL`
	__GongStructShape__000011_Default_ATTRIBUTE_VALUE_REAL.ShowNbInstances = false
	__GongStructShape__000011_Default_ATTRIBUTE_VALUE_REAL.NbInstances = 0
	__GongStructShape__000011_Default_ATTRIBUTE_VALUE_REAL.Width = 240.000000
	__GongStructShape__000011_Default_ATTRIBUTE_VALUE_REAL.Height = 63.000000
	__GongStructShape__000011_Default_ATTRIBUTE_VALUE_REAL.IsSelected = false

	__GongStructShape__000012_Default_ATTRIBUTE_VALUE_STRING.Name = `Default-ATTRIBUTE_VALUE_STRING`

	//gong:ident [ref_models.ATTRIBUTE_VALUE_STRING] comment added to overcome the problem with the comment map association
	__GongStructShape__000012_Default_ATTRIBUTE_VALUE_STRING.Identifier = `ref_models.ATTRIBUTE_VALUE_STRING`
	__GongStructShape__000012_Default_ATTRIBUTE_VALUE_STRING.ShowNbInstances = false
	__GongStructShape__000012_Default_ATTRIBUTE_VALUE_STRING.NbInstances = 0
	__GongStructShape__000012_Default_ATTRIBUTE_VALUE_STRING.Width = 240.000000
	__GongStructShape__000012_Default_ATTRIBUTE_VALUE_STRING.Height = 78.000000
	__GongStructShape__000012_Default_ATTRIBUTE_VALUE_STRING.IsSelected = false

	__GongStructShape__000013_Default_ATTRIBUTE_VALUE_XHTML.Name = `Default-ATTRIBUTE_VALUE_XHTML`

	//gong:ident [ref_models.ATTRIBUTE_VALUE_XHTML] comment added to overcome the problem with the comment map association
	__GongStructShape__000013_Default_ATTRIBUTE_VALUE_XHTML.Identifier = `ref_models.ATTRIBUTE_VALUE_XHTML`
	__GongStructShape__000013_Default_ATTRIBUTE_VALUE_XHTML.ShowNbInstances = false
	__GongStructShape__000013_Default_ATTRIBUTE_VALUE_XHTML.NbInstances = 0
	__GongStructShape__000013_Default_ATTRIBUTE_VALUE_XHTML.Width = 240.000000
	__GongStructShape__000013_Default_ATTRIBUTE_VALUE_XHTML.Height = 78.000000
	__GongStructShape__000013_Default_ATTRIBUTE_VALUE_XHTML.IsSelected = false

	__GongStructShape__000014_Default_A_ATTRIBUTE_DEFINITION_XHTML_REF.Name = `Default-A_ATTRIBUTE_DEFINITION_XHTML_REF`

	//gong:ident [ref_models.A_ATTRIBUTE_DEFINITION_XHTML_REF] comment added to overcome the problem with the comment map association
	__GongStructShape__000014_Default_A_ATTRIBUTE_DEFINITION_XHTML_REF.Identifier = `ref_models.A_ATTRIBUTE_DEFINITION_XHTML_REF`
	__GongStructShape__000014_Default_A_ATTRIBUTE_DEFINITION_XHTML_REF.ShowNbInstances = false
	__GongStructShape__000014_Default_A_ATTRIBUTE_DEFINITION_XHTML_REF.NbInstances = 0
	__GongStructShape__000014_Default_A_ATTRIBUTE_DEFINITION_XHTML_REF.Width = 460.000000
	__GongStructShape__000014_Default_A_ATTRIBUTE_DEFINITION_XHTML_REF.Height = 78.000000
	__GongStructShape__000014_Default_A_ATTRIBUTE_DEFINITION_XHTML_REF.IsSelected = false

	__GongStructShape__000015_Default_A_ATTRIBUTE_VALUE_XHTML.Name = `Default-A_ATTRIBUTE_VALUE_XHTML`

	//gong:ident [ref_models.A_ATTRIBUTE_VALUE_XHTML] comment added to overcome the problem with the comment map association
	__GongStructShape__000015_Default_A_ATTRIBUTE_VALUE_XHTML.Identifier = `ref_models.A_ATTRIBUTE_VALUE_XHTML`
	__GongStructShape__000015_Default_A_ATTRIBUTE_VALUE_XHTML.ShowNbInstances = false
	__GongStructShape__000015_Default_A_ATTRIBUTE_VALUE_XHTML.NbInstances = 0
	__GongStructShape__000015_Default_A_ATTRIBUTE_VALUE_XHTML.Width = 240.000000
	__GongStructShape__000015_Default_A_ATTRIBUTE_VALUE_XHTML.Height = 63.000000
	__GongStructShape__000015_Default_A_ATTRIBUTE_VALUE_XHTML.IsSelected = false

	__GongStructShape__000016_Default_A_ATTRIBUTE_VALUE_XHTML_1.Name = `Default-A_ATTRIBUTE_VALUE_XHTML_1`

	//gong:ident [ref_models.A_ATTRIBUTE_VALUE_XHTML_1] comment added to overcome the problem with the comment map association
	__GongStructShape__000016_Default_A_ATTRIBUTE_VALUE_XHTML_1.Identifier = `ref_models.A_ATTRIBUTE_VALUE_XHTML_1`
	__GongStructShape__000016_Default_A_ATTRIBUTE_VALUE_XHTML_1.ShowNbInstances = false
	__GongStructShape__000016_Default_A_ATTRIBUTE_VALUE_XHTML_1.NbInstances = 0
	__GongStructShape__000016_Default_A_ATTRIBUTE_VALUE_XHTML_1.Width = 275.000000
	__GongStructShape__000016_Default_A_ATTRIBUTE_VALUE_XHTML_1.Height = 63.000000
	__GongStructShape__000016_Default_A_ATTRIBUTE_VALUE_XHTML_1.IsSelected = false

	__GongStructShape__000017_Default_A_CORE_CONTENT.Name = `Default-A_CORE_CONTENT`

	//gong:ident [ref_models.A_CORE_CONTENT] comment added to overcome the problem with the comment map association
	__GongStructShape__000017_Default_A_CORE_CONTENT.Identifier = `ref_models.A_CORE_CONTENT`
	__GongStructShape__000017_Default_A_CORE_CONTENT.ShowNbInstances = false
	__GongStructShape__000017_Default_A_CORE_CONTENT.NbInstances = 0
	__GongStructShape__000017_Default_A_CORE_CONTENT.Width = 240.000000
	__GongStructShape__000017_Default_A_CORE_CONTENT.Height = 63.000000
	__GongStructShape__000017_Default_A_CORE_CONTENT.IsSelected = false

	__GongStructShape__000018_Default_A_SPECIFICATIONS.Name = `Default-A_SPECIFICATIONS`

	//gong:ident [ref_models.A_SPECIFICATIONS] comment added to overcome the problem with the comment map association
	__GongStructShape__000018_Default_A_SPECIFICATIONS.Identifier = `ref_models.A_SPECIFICATIONS`
	__GongStructShape__000018_Default_A_SPECIFICATIONS.ShowNbInstances = false
	__GongStructShape__000018_Default_A_SPECIFICATIONS.NbInstances = 0
	__GongStructShape__000018_Default_A_SPECIFICATIONS.Width = 240.000000
	__GongStructShape__000018_Default_A_SPECIFICATIONS.Height = 63.000000
	__GongStructShape__000018_Default_A_SPECIFICATIONS.IsSelected = false

	__GongStructShape__000019_Default_A_SPEC_OBJECTS.Name = `Default-A_SPEC_OBJECTS`

	//gong:ident [ref_models.A_SPEC_OBJECTS] comment added to overcome the problem with the comment map association
	__GongStructShape__000019_Default_A_SPEC_OBJECTS.Identifier = `ref_models.A_SPEC_OBJECTS`
	__GongStructShape__000019_Default_A_SPEC_OBJECTS.ShowNbInstances = false
	__GongStructShape__000019_Default_A_SPEC_OBJECTS.NbInstances = 0
	__GongStructShape__000019_Default_A_SPEC_OBJECTS.Width = 240.000000
	__GongStructShape__000019_Default_A_SPEC_OBJECTS.Height = 63.000000
	__GongStructShape__000019_Default_A_SPEC_OBJECTS.IsSelected = false

	__GongStructShape__000020_Default_A_SPEC_OBJECT_TYPE_REF.Name = `Default-A_SPEC_OBJECT_TYPE_REF`

	//gong:ident [ref_models.A_SPEC_OBJECT_TYPE_REF] comment added to overcome the problem with the comment map association
	__GongStructShape__000020_Default_A_SPEC_OBJECT_TYPE_REF.Identifier = `ref_models.A_SPEC_OBJECT_TYPE_REF`
	__GongStructShape__000020_Default_A_SPEC_OBJECT_TYPE_REF.ShowNbInstances = false
	__GongStructShape__000020_Default_A_SPEC_OBJECT_TYPE_REF.NbInstances = 0
	__GongStructShape__000020_Default_A_SPEC_OBJECT_TYPE_REF.Width = 240.000000
	__GongStructShape__000020_Default_A_SPEC_OBJECT_TYPE_REF.Height = 63.000000
	__GongStructShape__000020_Default_A_SPEC_OBJECT_TYPE_REF.IsSelected = false

	__GongStructShape__000021_Default_A_SPEC_RELATIONS.Name = `Default-A_SPEC_RELATIONS`

	//gong:ident [ref_models.A_SPEC_RELATIONS] comment added to overcome the problem with the comment map association
	__GongStructShape__000021_Default_A_SPEC_RELATIONS.Identifier = `ref_models.A_SPEC_RELATIONS`
	__GongStructShape__000021_Default_A_SPEC_RELATIONS.ShowNbInstances = false
	__GongStructShape__000021_Default_A_SPEC_RELATIONS.NbInstances = 0
	__GongStructShape__000021_Default_A_SPEC_RELATIONS.Width = 240.000000
	__GongStructShape__000021_Default_A_SPEC_RELATIONS.Height = 63.000000
	__GongStructShape__000021_Default_A_SPEC_RELATIONS.IsSelected = false

	__GongStructShape__000022_Default_A_SPEC_RELATION_TYPE_REF.Name = `Default-A_SPEC_RELATION_TYPE_REF`

	//gong:ident [ref_models.A_SPEC_RELATION_TYPE_REF] comment added to overcome the problem with the comment map association
	__GongStructShape__000022_Default_A_SPEC_RELATION_TYPE_REF.Identifier = `ref_models.A_SPEC_RELATION_TYPE_REF`
	__GongStructShape__000022_Default_A_SPEC_RELATION_TYPE_REF.ShowNbInstances = false
	__GongStructShape__000022_Default_A_SPEC_RELATION_TYPE_REF.NbInstances = 0
	__GongStructShape__000022_Default_A_SPEC_RELATION_TYPE_REF.Width = 240.000000
	__GongStructShape__000022_Default_A_SPEC_RELATION_TYPE_REF.Height = 63.000000
	__GongStructShape__000022_Default_A_SPEC_RELATION_TYPE_REF.IsSelected = false

	__GongStructShape__000023_Default_A_THE_HEADER.Name = `Default-A_THE_HEADER`

	//gong:ident [ref_models.A_THE_HEADER] comment added to overcome the problem with the comment map association
	__GongStructShape__000023_Default_A_THE_HEADER.Identifier = `ref_models.A_THE_HEADER`
	__GongStructShape__000023_Default_A_THE_HEADER.ShowNbInstances = false
	__GongStructShape__000023_Default_A_THE_HEADER.NbInstances = 0
	__GongStructShape__000023_Default_A_THE_HEADER.Width = 240.000000
	__GongStructShape__000023_Default_A_THE_HEADER.Height = 63.000000
	__GongStructShape__000023_Default_A_THE_HEADER.IsSelected = false

	__GongStructShape__000024_Default_REQ_IF.Name = `Default-REQ_IF`

	//gong:ident [ref_models.REQ_IF] comment added to overcome the problem with the comment map association
	__GongStructShape__000024_Default_REQ_IF.Identifier = `ref_models.REQ_IF`
	__GongStructShape__000024_Default_REQ_IF.ShowNbInstances = false
	__GongStructShape__000024_Default_REQ_IF.NbInstances = 0
	__GongStructShape__000024_Default_REQ_IF.Width = 240.000000
	__GongStructShape__000024_Default_REQ_IF.Height = 63.000000
	__GongStructShape__000024_Default_REQ_IF.IsSelected = false

	__GongStructShape__000025_Default_REQ_IF_CONTENT.Name = `Default-REQ_IF_CONTENT`

	//gong:ident [ref_models.REQ_IF_CONTENT] comment added to overcome the problem with the comment map association
	__GongStructShape__000025_Default_REQ_IF_CONTENT.Identifier = `ref_models.REQ_IF_CONTENT`
	__GongStructShape__000025_Default_REQ_IF_CONTENT.ShowNbInstances = false
	__GongStructShape__000025_Default_REQ_IF_CONTENT.NbInstances = 0
	__GongStructShape__000025_Default_REQ_IF_CONTENT.Width = 240.000000
	__GongStructShape__000025_Default_REQ_IF_CONTENT.Height = 63.000000
	__GongStructShape__000025_Default_REQ_IF_CONTENT.IsSelected = false

	__GongStructShape__000026_Default_REQ_IF_HEADER.Name = `Default-REQ_IF_HEADER`

	//gong:ident [ref_models.REQ_IF_HEADER] comment added to overcome the problem with the comment map association
	__GongStructShape__000026_Default_REQ_IF_HEADER.Identifier = `ref_models.REQ_IF_HEADER`
	__GongStructShape__000026_Default_REQ_IF_HEADER.ShowNbInstances = false
	__GongStructShape__000026_Default_REQ_IF_HEADER.NbInstances = 0
	__GongStructShape__000026_Default_REQ_IF_HEADER.Width = 240.000000
	__GongStructShape__000026_Default_REQ_IF_HEADER.Height = 63.000000
	__GongStructShape__000026_Default_REQ_IF_HEADER.IsSelected = false

	__GongStructShape__000027_Default_SPEC_OBJECT.Name = `Default-SPEC_OBJECT`

	//gong:ident [ref_models.SPEC_OBJECT] comment added to overcome the problem with the comment map association
	__GongStructShape__000027_Default_SPEC_OBJECT.Identifier = `ref_models.SPEC_OBJECT`
	__GongStructShape__000027_Default_SPEC_OBJECT.ShowNbInstances = false
	__GongStructShape__000027_Default_SPEC_OBJECT.NbInstances = 0
	__GongStructShape__000027_Default_SPEC_OBJECT.Width = 240.000000
	__GongStructShape__000027_Default_SPEC_OBJECT.Height = 411.000000
	__GongStructShape__000027_Default_SPEC_OBJECT.IsSelected = false

	__GongStructShape__000028_Default_SPEC_RELATION.Name = `Default-SPEC_RELATION`

	//gong:ident [ref_models.SPEC_RELATION] comment added to overcome the problem with the comment map association
	__GongStructShape__000028_Default_SPEC_RELATION.Identifier = `ref_models.SPEC_RELATION`
	__GongStructShape__000028_Default_SPEC_RELATION.ShowNbInstances = false
	__GongStructShape__000028_Default_SPEC_RELATION.NbInstances = 0
	__GongStructShape__000028_Default_SPEC_RELATION.Width = 240.000000
	__GongStructShape__000028_Default_SPEC_RELATION.Height = 63.000000
	__GongStructShape__000028_Default_SPEC_RELATION.IsSelected = false

	__GongStructShape__000029_Default_XHTML_CONTENT.Name = `Default-XHTML_CONTENT`

	//gong:ident [ref_models.XHTML_CONTENT] comment added to overcome the problem with the comment map association
	__GongStructShape__000029_Default_XHTML_CONTENT.Identifier = `ref_models.XHTML_CONTENT`
	__GongStructShape__000029_Default_XHTML_CONTENT.ShowNbInstances = false
	__GongStructShape__000029_Default_XHTML_CONTENT.NbInstances = 0
	__GongStructShape__000029_Default_XHTML_CONTENT.Width = 240.000000
	__GongStructShape__000029_Default_XHTML_CONTENT.Height = 211.000000
	__GongStructShape__000029_Default_XHTML_CONTENT.IsSelected = false

	__Link__000000_ATTRIBUTE_VALUE_BOOLEAN.Name = `ATTRIBUTE_VALUE_BOOLEAN`

	//gong:ident [ref_models.A_ATTRIBUTE_VALUE_XHTML_1.ATTRIBUTE_VALUE_BOOLEAN] comment added to overcome the problem with the comment map association
	__Link__000000_ATTRIBUTE_VALUE_BOOLEAN.Identifier = `ref_models.A_ATTRIBUTE_VALUE_XHTML_1.ATTRIBUTE_VALUE_BOOLEAN`

	//gong:ident [ref_models.ATTRIBUTE_VALUE_BOOLEAN] comment added to overcome the problem with the comment map association
	__Link__000000_ATTRIBUTE_VALUE_BOOLEAN.Fieldtypename = `ref_models.ATTRIBUTE_VALUE_BOOLEAN`
	__Link__000000_ATTRIBUTE_VALUE_BOOLEAN.FieldOffsetX = -50.000000
	__Link__000000_ATTRIBUTE_VALUE_BOOLEAN.FieldOffsetY = -16.000000
	__Link__000000_ATTRIBUTE_VALUE_BOOLEAN.TargetMultiplicity = models.MANY
	__Link__000000_ATTRIBUTE_VALUE_BOOLEAN.TargetMultiplicityOffsetX = -50.000000
	__Link__000000_ATTRIBUTE_VALUE_BOOLEAN.TargetMultiplicityOffsetY = 16.000000
	__Link__000000_ATTRIBUTE_VALUE_BOOLEAN.SourceMultiplicity = models.MANY
	__Link__000000_ATTRIBUTE_VALUE_BOOLEAN.SourceMultiplicityOffsetX = 10.000000
	__Link__000000_ATTRIBUTE_VALUE_BOOLEAN.SourceMultiplicityOffsetY = -50.000000
	__Link__000000_ATTRIBUTE_VALUE_BOOLEAN.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_ATTRIBUTE_VALUE_BOOLEAN.StartRatio = 0.500000
	__Link__000000_ATTRIBUTE_VALUE_BOOLEAN.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_ATTRIBUTE_VALUE_BOOLEAN.EndRatio = 0.500000
	__Link__000000_ATTRIBUTE_VALUE_BOOLEAN.CornerOffsetRatio = 1.380000

	__Link__000001_ATTRIBUTE_VALUE_DATE.Name = `ATTRIBUTE_VALUE_DATE`

	//gong:ident [ref_models.A_ATTRIBUTE_VALUE_XHTML_1.ATTRIBUTE_VALUE_DATE] comment added to overcome the problem with the comment map association
	__Link__000001_ATTRIBUTE_VALUE_DATE.Identifier = `ref_models.A_ATTRIBUTE_VALUE_XHTML_1.ATTRIBUTE_VALUE_DATE`

	//gong:ident [ref_models.ATTRIBUTE_VALUE_DATE] comment added to overcome the problem with the comment map association
	__Link__000001_ATTRIBUTE_VALUE_DATE.Fieldtypename = `ref_models.ATTRIBUTE_VALUE_DATE`
	__Link__000001_ATTRIBUTE_VALUE_DATE.FieldOffsetX = -50.000000
	__Link__000001_ATTRIBUTE_VALUE_DATE.FieldOffsetY = -16.000000
	__Link__000001_ATTRIBUTE_VALUE_DATE.TargetMultiplicity = models.MANY
	__Link__000001_ATTRIBUTE_VALUE_DATE.TargetMultiplicityOffsetX = -50.000000
	__Link__000001_ATTRIBUTE_VALUE_DATE.TargetMultiplicityOffsetY = 16.000000
	__Link__000001_ATTRIBUTE_VALUE_DATE.SourceMultiplicity = models.MANY
	__Link__000001_ATTRIBUTE_VALUE_DATE.SourceMultiplicityOffsetX = 10.000000
	__Link__000001_ATTRIBUTE_VALUE_DATE.SourceMultiplicityOffsetY = -50.000000
	__Link__000001_ATTRIBUTE_VALUE_DATE.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000001_ATTRIBUTE_VALUE_DATE.StartRatio = 0.500000
	__Link__000001_ATTRIBUTE_VALUE_DATE.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000001_ATTRIBUTE_VALUE_DATE.EndRatio = 0.500000
	__Link__000001_ATTRIBUTE_VALUE_DATE.CornerOffsetRatio = 1.380000

	__Link__000002_ATTRIBUTE_VALUE_ENUMERATION.Name = `ATTRIBUTE_VALUE_ENUMERATION`

	//gong:ident [ref_models.A_ATTRIBUTE_VALUE_XHTML_1.ATTRIBUTE_VALUE_ENUMERATION] comment added to overcome the problem with the comment map association
	__Link__000002_ATTRIBUTE_VALUE_ENUMERATION.Identifier = `ref_models.A_ATTRIBUTE_VALUE_XHTML_1.ATTRIBUTE_VALUE_ENUMERATION`

	//gong:ident [ref_models.ATTRIBUTE_VALUE_ENUMERATION] comment added to overcome the problem with the comment map association
	__Link__000002_ATTRIBUTE_VALUE_ENUMERATION.Fieldtypename = `ref_models.ATTRIBUTE_VALUE_ENUMERATION`
	__Link__000002_ATTRIBUTE_VALUE_ENUMERATION.FieldOffsetX = -50.000000
	__Link__000002_ATTRIBUTE_VALUE_ENUMERATION.FieldOffsetY = -16.000000
	__Link__000002_ATTRIBUTE_VALUE_ENUMERATION.TargetMultiplicity = models.MANY
	__Link__000002_ATTRIBUTE_VALUE_ENUMERATION.TargetMultiplicityOffsetX = -50.000000
	__Link__000002_ATTRIBUTE_VALUE_ENUMERATION.TargetMultiplicityOffsetY = 16.000000
	__Link__000002_ATTRIBUTE_VALUE_ENUMERATION.SourceMultiplicity = models.MANY
	__Link__000002_ATTRIBUTE_VALUE_ENUMERATION.SourceMultiplicityOffsetX = 10.000000
	__Link__000002_ATTRIBUTE_VALUE_ENUMERATION.SourceMultiplicityOffsetY = -50.000000
	__Link__000002_ATTRIBUTE_VALUE_ENUMERATION.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000002_ATTRIBUTE_VALUE_ENUMERATION.StartRatio = 0.500000
	__Link__000002_ATTRIBUTE_VALUE_ENUMERATION.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000002_ATTRIBUTE_VALUE_ENUMERATION.EndRatio = 0.500000
	__Link__000002_ATTRIBUTE_VALUE_ENUMERATION.CornerOffsetRatio = 1.380000

	__Link__000003_ATTRIBUTE_VALUE_INTEGER.Name = `ATTRIBUTE_VALUE_INTEGER`

	//gong:ident [ref_models.A_ATTRIBUTE_VALUE_XHTML_1.ATTRIBUTE_VALUE_INTEGER] comment added to overcome the problem with the comment map association
	__Link__000003_ATTRIBUTE_VALUE_INTEGER.Identifier = `ref_models.A_ATTRIBUTE_VALUE_XHTML_1.ATTRIBUTE_VALUE_INTEGER`

	//gong:ident [ref_models.ATTRIBUTE_VALUE_INTEGER] comment added to overcome the problem with the comment map association
	__Link__000003_ATTRIBUTE_VALUE_INTEGER.Fieldtypename = `ref_models.ATTRIBUTE_VALUE_INTEGER`
	__Link__000003_ATTRIBUTE_VALUE_INTEGER.FieldOffsetX = -50.000000
	__Link__000003_ATTRIBUTE_VALUE_INTEGER.FieldOffsetY = -16.000000
	__Link__000003_ATTRIBUTE_VALUE_INTEGER.TargetMultiplicity = models.MANY
	__Link__000003_ATTRIBUTE_VALUE_INTEGER.TargetMultiplicityOffsetX = -50.000000
	__Link__000003_ATTRIBUTE_VALUE_INTEGER.TargetMultiplicityOffsetY = 16.000000
	__Link__000003_ATTRIBUTE_VALUE_INTEGER.SourceMultiplicity = models.MANY
	__Link__000003_ATTRIBUTE_VALUE_INTEGER.SourceMultiplicityOffsetX = 10.000000
	__Link__000003_ATTRIBUTE_VALUE_INTEGER.SourceMultiplicityOffsetY = -50.000000
	__Link__000003_ATTRIBUTE_VALUE_INTEGER.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000003_ATTRIBUTE_VALUE_INTEGER.StartRatio = 0.500000
	__Link__000003_ATTRIBUTE_VALUE_INTEGER.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000003_ATTRIBUTE_VALUE_INTEGER.EndRatio = 0.500000
	__Link__000003_ATTRIBUTE_VALUE_INTEGER.CornerOffsetRatio = 1.380000

	__Link__000004_ATTRIBUTE_VALUE_REAL.Name = `ATTRIBUTE_VALUE_REAL`

	//gong:ident [ref_models.A_ATTRIBUTE_VALUE_XHTML_1.ATTRIBUTE_VALUE_REAL] comment added to overcome the problem with the comment map association
	__Link__000004_ATTRIBUTE_VALUE_REAL.Identifier = `ref_models.A_ATTRIBUTE_VALUE_XHTML_1.ATTRIBUTE_VALUE_REAL`

	//gong:ident [ref_models.ATTRIBUTE_VALUE_REAL] comment added to overcome the problem with the comment map association
	__Link__000004_ATTRIBUTE_VALUE_REAL.Fieldtypename = `ref_models.ATTRIBUTE_VALUE_REAL`
	__Link__000004_ATTRIBUTE_VALUE_REAL.FieldOffsetX = -50.000000
	__Link__000004_ATTRIBUTE_VALUE_REAL.FieldOffsetY = -16.000000
	__Link__000004_ATTRIBUTE_VALUE_REAL.TargetMultiplicity = models.MANY
	__Link__000004_ATTRIBUTE_VALUE_REAL.TargetMultiplicityOffsetX = -50.000000
	__Link__000004_ATTRIBUTE_VALUE_REAL.TargetMultiplicityOffsetY = 16.000000
	__Link__000004_ATTRIBUTE_VALUE_REAL.SourceMultiplicity = models.MANY
	__Link__000004_ATTRIBUTE_VALUE_REAL.SourceMultiplicityOffsetX = 10.000000
	__Link__000004_ATTRIBUTE_VALUE_REAL.SourceMultiplicityOffsetY = -50.000000
	__Link__000004_ATTRIBUTE_VALUE_REAL.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000004_ATTRIBUTE_VALUE_REAL.StartRatio = 0.500000
	__Link__000004_ATTRIBUTE_VALUE_REAL.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000004_ATTRIBUTE_VALUE_REAL.EndRatio = 0.500000
	__Link__000004_ATTRIBUTE_VALUE_REAL.CornerOffsetRatio = 1.380000

	__Link__000005_ATTRIBUTE_VALUE_STRING.Name = `ATTRIBUTE_VALUE_STRING`

	//gong:ident [ref_models.A_ATTRIBUTE_VALUE_XHTML_1.ATTRIBUTE_VALUE_STRING] comment added to overcome the problem with the comment map association
	__Link__000005_ATTRIBUTE_VALUE_STRING.Identifier = `ref_models.A_ATTRIBUTE_VALUE_XHTML_1.ATTRIBUTE_VALUE_STRING`

	//gong:ident [ref_models.ATTRIBUTE_VALUE_STRING] comment added to overcome the problem with the comment map association
	__Link__000005_ATTRIBUTE_VALUE_STRING.Fieldtypename = `ref_models.ATTRIBUTE_VALUE_STRING`
	__Link__000005_ATTRIBUTE_VALUE_STRING.FieldOffsetX = -50.000000
	__Link__000005_ATTRIBUTE_VALUE_STRING.FieldOffsetY = -16.000000
	__Link__000005_ATTRIBUTE_VALUE_STRING.TargetMultiplicity = models.MANY
	__Link__000005_ATTRIBUTE_VALUE_STRING.TargetMultiplicityOffsetX = -50.000000
	__Link__000005_ATTRIBUTE_VALUE_STRING.TargetMultiplicityOffsetY = 16.000000
	__Link__000005_ATTRIBUTE_VALUE_STRING.SourceMultiplicity = models.MANY
	__Link__000005_ATTRIBUTE_VALUE_STRING.SourceMultiplicityOffsetX = 10.000000
	__Link__000005_ATTRIBUTE_VALUE_STRING.SourceMultiplicityOffsetY = -50.000000
	__Link__000005_ATTRIBUTE_VALUE_STRING.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000005_ATTRIBUTE_VALUE_STRING.StartRatio = 0.500000
	__Link__000005_ATTRIBUTE_VALUE_STRING.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000005_ATTRIBUTE_VALUE_STRING.EndRatio = 0.500000
	__Link__000005_ATTRIBUTE_VALUE_STRING.CornerOffsetRatio = 1.380000

	__Link__000006_ATTRIBUTE_VALUE_XHTML.Name = `ATTRIBUTE_VALUE_XHTML`

	//gong:ident [ref_models.A_ATTRIBUTE_VALUE_XHTML_1.ATTRIBUTE_VALUE_XHTML] comment added to overcome the problem with the comment map association
	__Link__000006_ATTRIBUTE_VALUE_XHTML.Identifier = `ref_models.A_ATTRIBUTE_VALUE_XHTML_1.ATTRIBUTE_VALUE_XHTML`

	//gong:ident [ref_models.ATTRIBUTE_VALUE_XHTML] comment added to overcome the problem with the comment map association
	__Link__000006_ATTRIBUTE_VALUE_XHTML.Fieldtypename = `ref_models.ATTRIBUTE_VALUE_XHTML`
	__Link__000006_ATTRIBUTE_VALUE_XHTML.FieldOffsetX = -50.000000
	__Link__000006_ATTRIBUTE_VALUE_XHTML.FieldOffsetY = -16.000000
	__Link__000006_ATTRIBUTE_VALUE_XHTML.TargetMultiplicity = models.MANY
	__Link__000006_ATTRIBUTE_VALUE_XHTML.TargetMultiplicityOffsetX = -50.000000
	__Link__000006_ATTRIBUTE_VALUE_XHTML.TargetMultiplicityOffsetY = 16.000000
	__Link__000006_ATTRIBUTE_VALUE_XHTML.SourceMultiplicity = models.MANY
	__Link__000006_ATTRIBUTE_VALUE_XHTML.SourceMultiplicityOffsetX = 10.000000
	__Link__000006_ATTRIBUTE_VALUE_XHTML.SourceMultiplicityOffsetY = -50.000000
	__Link__000006_ATTRIBUTE_VALUE_XHTML.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000006_ATTRIBUTE_VALUE_XHTML.StartRatio = 0.500000
	__Link__000006_ATTRIBUTE_VALUE_XHTML.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000006_ATTRIBUTE_VALUE_XHTML.EndRatio = 0.500000
	__Link__000006_ATTRIBUTE_VALUE_XHTML.CornerOffsetRatio = 1.380000

	__Link__000007_CORE_CONTENT.Name = `CORE_CONTENT`

	//gong:ident [ref_models.REQ_IF.CORE_CONTENT] comment added to overcome the problem with the comment map association
	__Link__000007_CORE_CONTENT.Identifier = `ref_models.REQ_IF.CORE_CONTENT`

	//gong:ident [ref_models.A_CORE_CONTENT] comment added to overcome the problem with the comment map association
	__Link__000007_CORE_CONTENT.Fieldtypename = `ref_models.A_CORE_CONTENT`
	__Link__000007_CORE_CONTENT.FieldOffsetX = -50.000000
	__Link__000007_CORE_CONTENT.FieldOffsetY = -16.000000
	__Link__000007_CORE_CONTENT.TargetMultiplicity = models.MANY
	__Link__000007_CORE_CONTENT.TargetMultiplicityOffsetX = -50.000000
	__Link__000007_CORE_CONTENT.TargetMultiplicityOffsetY = 16.000000
	__Link__000007_CORE_CONTENT.SourceMultiplicity = models.MANY
	__Link__000007_CORE_CONTENT.SourceMultiplicityOffsetX = 10.000000
	__Link__000007_CORE_CONTENT.SourceMultiplicityOffsetY = -50.000000
	__Link__000007_CORE_CONTENT.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000007_CORE_CONTENT.StartRatio = 0.500000
	__Link__000007_CORE_CONTENT.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000007_CORE_CONTENT.EndRatio = 0.500000
	__Link__000007_CORE_CONTENT.CornerOffsetRatio = 1.380000

	__Link__000008_DEFINITION.Name = `DEFINITION`

	//gong:ident [ref_models.ATTRIBUTE_VALUE_XHTML.DEFINITION] comment added to overcome the problem with the comment map association
	__Link__000008_DEFINITION.Identifier = `ref_models.ATTRIBUTE_VALUE_XHTML.DEFINITION`

	//gong:ident [ref_models.A_ATTRIBUTE_DEFINITION_XHTML_REF] comment added to overcome the problem with the comment map association
	__Link__000008_DEFINITION.Fieldtypename = `ref_models.A_ATTRIBUTE_DEFINITION_XHTML_REF`
	__Link__000008_DEFINITION.FieldOffsetX = -50.000000
	__Link__000008_DEFINITION.FieldOffsetY = -16.000000
	__Link__000008_DEFINITION.TargetMultiplicity = models.MANY
	__Link__000008_DEFINITION.TargetMultiplicityOffsetX = -50.000000
	__Link__000008_DEFINITION.TargetMultiplicityOffsetY = 16.000000
	__Link__000008_DEFINITION.SourceMultiplicity = models.MANY
	__Link__000008_DEFINITION.SourceMultiplicityOffsetX = 10.000000
	__Link__000008_DEFINITION.SourceMultiplicityOffsetY = -50.000000
	__Link__000008_DEFINITION.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000008_DEFINITION.StartRatio = 0.500000
	__Link__000008_DEFINITION.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000008_DEFINITION.EndRatio = 0.500000
	__Link__000008_DEFINITION.CornerOffsetRatio = 1.380000

	__Link__000009_REQ_IF_CONTENT.Name = `REQ_IF_CONTENT`

	//gong:ident [ref_models.A_CORE_CONTENT.REQ_IF_CONTENT] comment added to overcome the problem with the comment map association
	__Link__000009_REQ_IF_CONTENT.Identifier = `ref_models.A_CORE_CONTENT.REQ_IF_CONTENT`

	//gong:ident [ref_models.REQ_IF_CONTENT] comment added to overcome the problem with the comment map association
	__Link__000009_REQ_IF_CONTENT.Fieldtypename = `ref_models.REQ_IF_CONTENT`
	__Link__000009_REQ_IF_CONTENT.FieldOffsetX = -50.000000
	__Link__000009_REQ_IF_CONTENT.FieldOffsetY = -16.000000
	__Link__000009_REQ_IF_CONTENT.TargetMultiplicity = models.MANY
	__Link__000009_REQ_IF_CONTENT.TargetMultiplicityOffsetX = -50.000000
	__Link__000009_REQ_IF_CONTENT.TargetMultiplicityOffsetY = 16.000000
	__Link__000009_REQ_IF_CONTENT.SourceMultiplicity = models.MANY
	__Link__000009_REQ_IF_CONTENT.SourceMultiplicityOffsetX = 10.000000
	__Link__000009_REQ_IF_CONTENT.SourceMultiplicityOffsetY = -50.000000
	__Link__000009_REQ_IF_CONTENT.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000009_REQ_IF_CONTENT.StartRatio = 0.500000
	__Link__000009_REQ_IF_CONTENT.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000009_REQ_IF_CONTENT.EndRatio = 0.500000
	__Link__000009_REQ_IF_CONTENT.CornerOffsetRatio = 1.380000

	__Link__000010_REQ_IF_HEADER.Name = `REQ_IF_HEADER`

	//gong:ident [ref_models.A_THE_HEADER.REQ_IF_HEADER] comment added to overcome the problem with the comment map association
	__Link__000010_REQ_IF_HEADER.Identifier = `ref_models.A_THE_HEADER.REQ_IF_HEADER`

	//gong:ident [ref_models.REQ_IF_HEADER] comment added to overcome the problem with the comment map association
	__Link__000010_REQ_IF_HEADER.Fieldtypename = `ref_models.REQ_IF_HEADER`
	__Link__000010_REQ_IF_HEADER.FieldOffsetX = -50.000000
	__Link__000010_REQ_IF_HEADER.FieldOffsetY = -16.000000
	__Link__000010_REQ_IF_HEADER.TargetMultiplicity = models.MANY
	__Link__000010_REQ_IF_HEADER.TargetMultiplicityOffsetX = -50.000000
	__Link__000010_REQ_IF_HEADER.TargetMultiplicityOffsetY = 16.000000
	__Link__000010_REQ_IF_HEADER.SourceMultiplicity = models.MANY
	__Link__000010_REQ_IF_HEADER.SourceMultiplicityOffsetX = 10.000000
	__Link__000010_REQ_IF_HEADER.SourceMultiplicityOffsetY = -50.000000
	__Link__000010_REQ_IF_HEADER.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000010_REQ_IF_HEADER.StartRatio = 0.500000
	__Link__000010_REQ_IF_HEADER.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000010_REQ_IF_HEADER.EndRatio = 0.500000
	__Link__000010_REQ_IF_HEADER.CornerOffsetRatio = 1.380000

	__Link__000011_SPECIFICATIONS.Name = `SPECIFICATIONS`

	//gong:ident [ref_models.REQ_IF_CONTENT.SPECIFICATIONS] comment added to overcome the problem with the comment map association
	__Link__000011_SPECIFICATIONS.Identifier = `ref_models.REQ_IF_CONTENT.SPECIFICATIONS`

	//gong:ident [ref_models.A_SPECIFICATIONS] comment added to overcome the problem with the comment map association
	__Link__000011_SPECIFICATIONS.Fieldtypename = `ref_models.A_SPECIFICATIONS`
	__Link__000011_SPECIFICATIONS.FieldOffsetX = -50.000000
	__Link__000011_SPECIFICATIONS.FieldOffsetY = -16.000000
	__Link__000011_SPECIFICATIONS.TargetMultiplicity = models.MANY
	__Link__000011_SPECIFICATIONS.TargetMultiplicityOffsetX = -50.000000
	__Link__000011_SPECIFICATIONS.TargetMultiplicityOffsetY = 16.000000
	__Link__000011_SPECIFICATIONS.SourceMultiplicity = models.MANY
	__Link__000011_SPECIFICATIONS.SourceMultiplicityOffsetX = 10.000000
	__Link__000011_SPECIFICATIONS.SourceMultiplicityOffsetY = -50.000000
	__Link__000011_SPECIFICATIONS.StartOrientation = models.ORIENTATION_VERTICAL
	__Link__000011_SPECIFICATIONS.StartRatio = 0.826389
	__Link__000011_SPECIFICATIONS.EndOrientation = models.ORIENTATION_VERTICAL
	__Link__000011_SPECIFICATIONS.EndRatio = 0.763889
	__Link__000011_SPECIFICATIONS.CornerOffsetRatio = 1.732804

	__Link__000012_SPEC_OBJECT.Name = `SPEC_OBJECT`

	//gong:ident [ref_models.A_SPEC_OBJECTS.SPEC_OBJECT] comment added to overcome the problem with the comment map association
	__Link__000012_SPEC_OBJECT.Identifier = `ref_models.A_SPEC_OBJECTS.SPEC_OBJECT`

	//gong:ident [ref_models.SPEC_OBJECT] comment added to overcome the problem with the comment map association
	__Link__000012_SPEC_OBJECT.Fieldtypename = `ref_models.SPEC_OBJECT`
	__Link__000012_SPEC_OBJECT.FieldOffsetX = -50.000000
	__Link__000012_SPEC_OBJECT.FieldOffsetY = -16.000000
	__Link__000012_SPEC_OBJECT.TargetMultiplicity = models.MANY
	__Link__000012_SPEC_OBJECT.TargetMultiplicityOffsetX = -50.000000
	__Link__000012_SPEC_OBJECT.TargetMultiplicityOffsetY = 16.000000
	__Link__000012_SPEC_OBJECT.SourceMultiplicity = models.MANY
	__Link__000012_SPEC_OBJECT.SourceMultiplicityOffsetX = 10.000000
	__Link__000012_SPEC_OBJECT.SourceMultiplicityOffsetY = -50.000000
	__Link__000012_SPEC_OBJECT.StartOrientation = models.ORIENTATION_VERTICAL
	__Link__000012_SPEC_OBJECT.StartRatio = 0.576389
	__Link__000012_SPEC_OBJECT.EndOrientation = models.ORIENTATION_VERTICAL
	__Link__000012_SPEC_OBJECT.EndRatio = 0.597222
	__Link__000012_SPEC_OBJECT.CornerOffsetRatio = 1.796627

	__Link__000013_SPEC_OBJECTS.Name = `SPEC_OBJECTS`

	//gong:ident [ref_models.REQ_IF_CONTENT.SPEC_OBJECTS] comment added to overcome the problem with the comment map association
	__Link__000013_SPEC_OBJECTS.Identifier = `ref_models.REQ_IF_CONTENT.SPEC_OBJECTS`

	//gong:ident [ref_models.A_SPEC_OBJECTS] comment added to overcome the problem with the comment map association
	__Link__000013_SPEC_OBJECTS.Fieldtypename = `ref_models.A_SPEC_OBJECTS`
	__Link__000013_SPEC_OBJECTS.FieldOffsetX = -50.000000
	__Link__000013_SPEC_OBJECTS.FieldOffsetY = -16.000000
	__Link__000013_SPEC_OBJECTS.TargetMultiplicity = models.MANY
	__Link__000013_SPEC_OBJECTS.TargetMultiplicityOffsetX = -50.000000
	__Link__000013_SPEC_OBJECTS.TargetMultiplicityOffsetY = 16.000000
	__Link__000013_SPEC_OBJECTS.SourceMultiplicity = models.MANY
	__Link__000013_SPEC_OBJECTS.SourceMultiplicityOffsetX = 10.000000
	__Link__000013_SPEC_OBJECTS.SourceMultiplicityOffsetY = -50.000000
	__Link__000013_SPEC_OBJECTS.StartOrientation = models.ORIENTATION_VERTICAL
	__Link__000013_SPEC_OBJECTS.StartRatio = 0.580555
	__Link__000013_SPEC_OBJECTS.EndOrientation = models.ORIENTATION_VERTICAL
	__Link__000013_SPEC_OBJECTS.EndRatio = 0.609722
	__Link__000013_SPEC_OBJECTS.CornerOffsetRatio = 2.558201

	__Link__000014_SPEC_RELATIONS.Name = `SPEC_RELATIONS`

	//gong:ident [ref_models.REQ_IF_CONTENT.SPEC_RELATIONS] comment added to overcome the problem with the comment map association
	__Link__000014_SPEC_RELATIONS.Identifier = `ref_models.REQ_IF_CONTENT.SPEC_RELATIONS`

	//gong:ident [ref_models.A_SPEC_RELATIONS] comment added to overcome the problem with the comment map association
	__Link__000014_SPEC_RELATIONS.Fieldtypename = `ref_models.A_SPEC_RELATIONS`
	__Link__000014_SPEC_RELATIONS.FieldOffsetX = -50.000000
	__Link__000014_SPEC_RELATIONS.FieldOffsetY = -16.000000
	__Link__000014_SPEC_RELATIONS.TargetMultiplicity = models.MANY
	__Link__000014_SPEC_RELATIONS.TargetMultiplicityOffsetX = -50.000000
	__Link__000014_SPEC_RELATIONS.TargetMultiplicityOffsetY = 16.000000
	__Link__000014_SPEC_RELATIONS.SourceMultiplicity = models.MANY
	__Link__000014_SPEC_RELATIONS.SourceMultiplicityOffsetX = 10.000000
	__Link__000014_SPEC_RELATIONS.SourceMultiplicityOffsetY = -50.000000
	__Link__000014_SPEC_RELATIONS.StartOrientation = models.ORIENTATION_VERTICAL
	__Link__000014_SPEC_RELATIONS.StartRatio = 0.309722
	__Link__000014_SPEC_RELATIONS.EndOrientation = models.ORIENTATION_VERTICAL
	__Link__000014_SPEC_RELATIONS.EndRatio = 0.680555
	__Link__000014_SPEC_RELATIONS.CornerOffsetRatio = 1.939153

	__Link__000015_THE_HEADER.Name = `THE_HEADER`

	//gong:ident [ref_models.REQ_IF.THE_HEADER] comment added to overcome the problem with the comment map association
	__Link__000015_THE_HEADER.Identifier = `ref_models.REQ_IF.THE_HEADER`

	//gong:ident [ref_models.A_THE_HEADER] comment added to overcome the problem with the comment map association
	__Link__000015_THE_HEADER.Fieldtypename = `ref_models.A_THE_HEADER`
	__Link__000015_THE_HEADER.FieldOffsetX = -50.000000
	__Link__000015_THE_HEADER.FieldOffsetY = -16.000000
	__Link__000015_THE_HEADER.TargetMultiplicity = models.MANY
	__Link__000015_THE_HEADER.TargetMultiplicityOffsetX = -50.000000
	__Link__000015_THE_HEADER.TargetMultiplicityOffsetY = 16.000000
	__Link__000015_THE_HEADER.SourceMultiplicity = models.MANY
	__Link__000015_THE_HEADER.SourceMultiplicityOffsetX = 10.000000
	__Link__000015_THE_HEADER.SourceMultiplicityOffsetY = -50.000000
	__Link__000015_THE_HEADER.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000015_THE_HEADER.StartRatio = 0.500000
	__Link__000015_THE_HEADER.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000015_THE_HEADER.EndRatio = 0.500000
	__Link__000015_THE_HEADER.CornerOffsetRatio = 1.380000

	__Link__000016_THE_ORIGINAL_VALUE.Name = `THE_ORIGINAL_VALUE`

	//gong:ident [ref_models.ATTRIBUTE_VALUE_XHTML.THE_ORIGINAL_VALUE] comment added to overcome the problem with the comment map association
	__Link__000016_THE_ORIGINAL_VALUE.Identifier = `ref_models.ATTRIBUTE_VALUE_XHTML.THE_ORIGINAL_VALUE`

	//gong:ident [ref_models.XHTML_CONTENT] comment added to overcome the problem with the comment map association
	__Link__000016_THE_ORIGINAL_VALUE.Fieldtypename = `ref_models.XHTML_CONTENT`
	__Link__000016_THE_ORIGINAL_VALUE.FieldOffsetX = -50.000000
	__Link__000016_THE_ORIGINAL_VALUE.FieldOffsetY = -16.000000
	__Link__000016_THE_ORIGINAL_VALUE.TargetMultiplicity = models.MANY
	__Link__000016_THE_ORIGINAL_VALUE.TargetMultiplicityOffsetX = -50.000000
	__Link__000016_THE_ORIGINAL_VALUE.TargetMultiplicityOffsetY = 16.000000
	__Link__000016_THE_ORIGINAL_VALUE.SourceMultiplicity = models.MANY
	__Link__000016_THE_ORIGINAL_VALUE.SourceMultiplicityOffsetX = 10.000000
	__Link__000016_THE_ORIGINAL_VALUE.SourceMultiplicityOffsetY = -50.000000
	__Link__000016_THE_ORIGINAL_VALUE.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000016_THE_ORIGINAL_VALUE.StartRatio = 0.500000
	__Link__000016_THE_ORIGINAL_VALUE.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000016_THE_ORIGINAL_VALUE.EndRatio = 0.729969
	__Link__000016_THE_ORIGINAL_VALUE.CornerOffsetRatio = 1.380000

	__Link__000017_THE_VALUE.Name = `THE_VALUE`

	//gong:ident [ref_models.ATTRIBUTE_VALUE_XHTML.THE_VALUE] comment added to overcome the problem with the comment map association
	__Link__000017_THE_VALUE.Identifier = `ref_models.ATTRIBUTE_VALUE_XHTML.THE_VALUE`

	//gong:ident [ref_models.XHTML_CONTENT] comment added to overcome the problem with the comment map association
	__Link__000017_THE_VALUE.Fieldtypename = `ref_models.XHTML_CONTENT`
	__Link__000017_THE_VALUE.FieldOffsetX = -50.000000
	__Link__000017_THE_VALUE.FieldOffsetY = -16.000000
	__Link__000017_THE_VALUE.TargetMultiplicity = models.MANY
	__Link__000017_THE_VALUE.TargetMultiplicityOffsetX = -50.000000
	__Link__000017_THE_VALUE.TargetMultiplicityOffsetY = 16.000000
	__Link__000017_THE_VALUE.SourceMultiplicity = models.MANY
	__Link__000017_THE_VALUE.SourceMultiplicityOffsetX = 10.000000
	__Link__000017_THE_VALUE.SourceMultiplicityOffsetY = -50.000000
	__Link__000017_THE_VALUE.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000017_THE_VALUE.StartRatio = 0.500000
	__Link__000017_THE_VALUE.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000017_THE_VALUE.EndRatio = 0.161249
	__Link__000017_THE_VALUE.CornerOffsetRatio = 1.380000

	__Link__000018_TYPE.Name = `TYPE`

	//gong:ident [ref_models.SPEC_OBJECT.TYPE] comment added to overcome the problem with the comment map association
	__Link__000018_TYPE.Identifier = `ref_models.SPEC_OBJECT.TYPE`

	//gong:ident [ref_models.A_SPEC_OBJECT_TYPE_REF] comment added to overcome the problem with the comment map association
	__Link__000018_TYPE.Fieldtypename = `ref_models.A_SPEC_OBJECT_TYPE_REF`
	__Link__000018_TYPE.FieldOffsetX = -50.000000
	__Link__000018_TYPE.FieldOffsetY = -16.000000
	__Link__000018_TYPE.TargetMultiplicity = models.MANY
	__Link__000018_TYPE.TargetMultiplicityOffsetX = -50.000000
	__Link__000018_TYPE.TargetMultiplicityOffsetY = 16.000000
	__Link__000018_TYPE.SourceMultiplicity = models.MANY
	__Link__000018_TYPE.SourceMultiplicityOffsetX = 10.000000
	__Link__000018_TYPE.SourceMultiplicityOffsetY = -50.000000
	__Link__000018_TYPE.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000018_TYPE.StartRatio = 0.289994
	__Link__000018_TYPE.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000018_TYPE.EndRatio = 0.500000
	__Link__000018_TYPE.CornerOffsetRatio = -0.330729

	__Link__000019_VALUES.Name = `VALUES`

	//gong:ident [ref_models.SPEC_OBJECT.VALUES] comment added to overcome the problem with the comment map association
	__Link__000019_VALUES.Identifier = `ref_models.SPEC_OBJECT.VALUES`

	//gong:ident [ref_models.A_ATTRIBUTE_VALUE_XHTML_1] comment added to overcome the problem with the comment map association
	__Link__000019_VALUES.Fieldtypename = `ref_models.A_ATTRIBUTE_VALUE_XHTML_1`
	__Link__000019_VALUES.FieldOffsetX = -50.000000
	__Link__000019_VALUES.FieldOffsetY = -16.000000
	__Link__000019_VALUES.TargetMultiplicity = models.MANY
	__Link__000019_VALUES.TargetMultiplicityOffsetX = -50.000000
	__Link__000019_VALUES.TargetMultiplicityOffsetY = 16.000000
	__Link__000019_VALUES.SourceMultiplicity = models.MANY
	__Link__000019_VALUES.SourceMultiplicityOffsetX = 10.000000
	__Link__000019_VALUES.SourceMultiplicityOffsetY = -50.000000
	__Link__000019_VALUES.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000019_VALUES.StartRatio = 0.500000
	__Link__000019_VALUES.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000019_VALUES.EndRatio = 0.500000
	__Link__000019_VALUES.CornerOffsetRatio = 1.380000

	__Position__000000_Pos_Default_ATTRIBUTE_DEFINITION_BOOLEAN.X = 103.999939
	__Position__000000_Pos_Default_ATTRIBUTE_DEFINITION_BOOLEAN.Y = 1168.500000
	__Position__000000_Pos_Default_ATTRIBUTE_DEFINITION_BOOLEAN.Name = `Pos-Default-ATTRIBUTE_DEFINITION_BOOLEAN`

	__Position__000001_Pos_Default_ATTRIBUTE_DEFINITION_DATE.X = 114.999939
	__Position__000001_Pos_Default_ATTRIBUTE_DEFINITION_DATE.Y = 1000.000000
	__Position__000001_Pos_Default_ATTRIBUTE_DEFINITION_DATE.Name = `Pos-Default-ATTRIBUTE_DEFINITION_DATE`

	__Position__000002_Pos_Default_ATTRIBUTE_DEFINITION_ENUMERATION.X = 103.999939
	__Position__000002_Pos_Default_ATTRIBUTE_DEFINITION_ENUMERATION.Y = 1080.000000
	__Position__000002_Pos_Default_ATTRIBUTE_DEFINITION_ENUMERATION.Name = `Pos-Default-ATTRIBUTE_DEFINITION_ENUMERATION`

	__Position__000003_Pos_Default_ATTRIBUTE_DEFINITION_INTEGER.X = 103.999939
	__Position__000003_Pos_Default_ATTRIBUTE_DEFINITION_INTEGER.Y = 903.000000
	__Position__000003_Pos_Default_ATTRIBUTE_DEFINITION_INTEGER.Name = `Pos-Default-ATTRIBUTE_DEFINITION_INTEGER`

	__Position__000004_Pos_Default_ATTRIBUTE_DEFINITION_REAL.X = 112.999939
	__Position__000004_Pos_Default_ATTRIBUTE_DEFINITION_REAL.Y = 1371.000000
	__Position__000004_Pos_Default_ATTRIBUTE_DEFINITION_REAL.Name = `Pos-Default-ATTRIBUTE_DEFINITION_REAL`

	__Position__000005_Pos_Default_ATTRIBUTE_DEFINITION_STRING.X = 102.999939
	__Position__000005_Pos_Default_ATTRIBUTE_DEFINITION_STRING.Y = 1281.000000
	__Position__000005_Pos_Default_ATTRIBUTE_DEFINITION_STRING.Name = `Pos-Default-ATTRIBUTE_DEFINITION_STRING`

	__Position__000006_Pos_Default_ATTRIBUTE_DEFINITION_XHTML.X = 1273.000000
	__Position__000006_Pos_Default_ATTRIBUTE_DEFINITION_XHTML.Y = 1056.500000
	__Position__000006_Pos_Default_ATTRIBUTE_DEFINITION_XHTML.Name = `Pos-Default-ATTRIBUTE_DEFINITION_XHTML`

	__Position__000007_Pos_Default_ATTRIBUTE_VALUE_BOOLEAN.X = 1639.999939
	__Position__000007_Pos_Default_ATTRIBUTE_VALUE_BOOLEAN.Y = 710.000000
	__Position__000007_Pos_Default_ATTRIBUTE_VALUE_BOOLEAN.Name = `Pos-Default-ATTRIBUTE_VALUE_BOOLEAN`

	__Position__000008_Pos_Default_ATTRIBUTE_VALUE_DATE.X = 1640.999939
	__Position__000008_Pos_Default_ATTRIBUTE_VALUE_DATE.Y = 624.000000
	__Position__000008_Pos_Default_ATTRIBUTE_VALUE_DATE.Name = `Pos-Default-ATTRIBUTE_VALUE_DATE`

	__Position__000009_Pos_Default_ATTRIBUTE_VALUE_ENUMERATION.X = 1638.999939
	__Position__000009_Pos_Default_ATTRIBUTE_VALUE_ENUMERATION.Y = 537.000000
	__Position__000009_Pos_Default_ATTRIBUTE_VALUE_ENUMERATION.Name = `Pos-Default-ATTRIBUTE_VALUE_ENUMERATION`

	__Position__000010_Pos_Default_ATTRIBUTE_VALUE_INTEGER.X = 1635.999939
	__Position__000010_Pos_Default_ATTRIBUTE_VALUE_INTEGER.Y = 450.000000
	__Position__000010_Pos_Default_ATTRIBUTE_VALUE_INTEGER.Name = `Pos-Default-ATTRIBUTE_VALUE_INTEGER`

	__Position__000011_Pos_Default_ATTRIBUTE_VALUE_REAL.X = 1631.999939
	__Position__000011_Pos_Default_ATTRIBUTE_VALUE_REAL.Y = 367.000000
	__Position__000011_Pos_Default_ATTRIBUTE_VALUE_REAL.Name = `Pos-Default-ATTRIBUTE_VALUE_REAL`

	__Position__000012_Pos_Default_ATTRIBUTE_VALUE_STRING.X = 1629.999939
	__Position__000012_Pos_Default_ATTRIBUTE_VALUE_STRING.Y = 283.000000
	__Position__000012_Pos_Default_ATTRIBUTE_VALUE_STRING.Name = `Pos-Default-ATTRIBUTE_VALUE_STRING`

	__Position__000013_Pos_Default_ATTRIBUTE_VALUE_XHTML.X = 1645.999939
	__Position__000013_Pos_Default_ATTRIBUTE_VALUE_XHTML.Y = 796.000000
	__Position__000013_Pos_Default_ATTRIBUTE_VALUE_XHTML.Name = `Pos-Default-ATTRIBUTE_VALUE_XHTML`

	__Position__000014_Pos_Default_A_ATTRIBUTE_DEFINITION_XHTML_REF.X = 1430.000000
	__Position__000014_Pos_Default_A_ATTRIBUTE_DEFINITION_XHTML_REF.Y = 1237.000000
	__Position__000014_Pos_Default_A_ATTRIBUTE_DEFINITION_XHTML_REF.Name = `Pos-Default-A_ATTRIBUTE_DEFINITION_XHTML_REF`

	__Position__000015_Pos_Default_A_ATTRIBUTE_VALUE_XHTML.X = 987.000000
	__Position__000015_Pos_Default_A_ATTRIBUTE_VALUE_XHTML.Y = 1064.000000
	__Position__000015_Pos_Default_A_ATTRIBUTE_VALUE_XHTML.Name = `Pos-Default-A_ATTRIBUTE_VALUE_XHTML`

	__Position__000016_Pos_Default_A_ATTRIBUTE_VALUE_XHTML_1.X = 1125.999939
	__Position__000016_Pos_Default_A_ATTRIBUTE_VALUE_XHTML_1.Y = 801.000000
	__Position__000016_Pos_Default_A_ATTRIBUTE_VALUE_XHTML_1.Name = `Pos-Default-A_ATTRIBUTE_VALUE_XHTML_1`

	__Position__000017_Pos_Default_A_CORE_CONTENT.X = 486.000000
	__Position__000017_Pos_Default_A_CORE_CONTENT.Y = 177.000000
	__Position__000017_Pos_Default_A_CORE_CONTENT.Name = `Pos-Default-A_CORE_CONTENT`

	__Position__000018_Pos_Default_A_SPECIFICATIONS.X = 850.000000
	__Position__000018_Pos_Default_A_SPECIFICATIONS.Y = 457.000000
	__Position__000018_Pos_Default_A_SPECIFICATIONS.Name = `Pos-Default-A_SPECIFICATIONS`

	__Position__000019_Pos_Default_A_SPEC_OBJECTS.X = 477.000000
	__Position__000019_Pos_Default_A_SPEC_OBJECTS.Y = 451.000000
	__Position__000019_Pos_Default_A_SPEC_OBJECTS.Name = `Pos-Default-A_SPEC_OBJECTS`

	__Position__000020_Pos_Default_A_SPEC_OBJECT_TYPE_REF.X = 61.999939
	__Position__000020_Pos_Default_A_SPEC_OBJECT_TYPE_REF.Y = 713.000000
	__Position__000020_Pos_Default_A_SPEC_OBJECT_TYPE_REF.Name = `Pos-Default-A_SPEC_OBJECT_TYPE_REF`

	__Position__000021_Pos_Default_A_SPEC_RELATIONS.X = 111.000000
	__Position__000021_Pos_Default_A_SPEC_RELATIONS.Y = 463.000000
	__Position__000021_Pos_Default_A_SPEC_RELATIONS.Name = `Pos-Default-A_SPEC_RELATIONS`

	__Position__000022_Pos_Default_A_SPEC_RELATION_TYPE_REF.X = 1107.999939
	__Position__000022_Pos_Default_A_SPEC_RELATION_TYPE_REF.Y = 612.000000
	__Position__000022_Pos_Default_A_SPEC_RELATION_TYPE_REF.Name = `Pos-Default-A_SPEC_RELATION_TYPE_REF`

	__Position__000023_Pos_Default_A_THE_HEADER.X = 482.000000
	__Position__000023_Pos_Default_A_THE_HEADER.Y = 29.000000
	__Position__000023_Pos_Default_A_THE_HEADER.Name = `Pos-Default-A_THE_HEADER`

	__Position__000024_Pos_Default_REQ_IF.X = 57.000000
	__Position__000024_Pos_Default_REQ_IF.Y = 25.000000
	__Position__000024_Pos_Default_REQ_IF.Name = `Pos-Default-REQ_IF`

	__Position__000025_Pos_Default_REQ_IF_CONTENT.X = 835.000000
	__Position__000025_Pos_Default_REQ_IF_CONTENT.Y = 174.000000
	__Position__000025_Pos_Default_REQ_IF_CONTENT.Name = `Pos-Default-REQ_IF_CONTENT`

	__Position__000026_Pos_Default_REQ_IF_HEADER.X = 835.000000
	__Position__000026_Pos_Default_REQ_IF_HEADER.Y = 29.000000
	__Position__000026_Pos_Default_REQ_IF_HEADER.Name = `Pos-Default-REQ_IF_HEADER`

	__Position__000027_Pos_Default_SPEC_OBJECT.X = 475.000000
	__Position__000027_Pos_Default_SPEC_OBJECT.Y = 621.999969
	__Position__000027_Pos_Default_SPEC_OBJECT.Name = `Pos-Default-SPEC_OBJECT`

	__Position__000028_Pos_Default_SPEC_RELATION.X = 110.000000
	__Position__000028_Pos_Default_SPEC_RELATION.Y = 585.000000
	__Position__000028_Pos_Default_SPEC_RELATION.Name = `Pos-Default-SPEC_RELATION`

	__Position__000029_Pos_Default_XHTML_CONTENT.X = 1646.000000
	__Position__000029_Pos_Default_XHTML_CONTENT.Y = 986.000000
	__Position__000029_Pos_Default_XHTML_CONTENT.Name = `Pos-Default-XHTML_CONTENT`

	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_ATTRIBUTE_VALUE_XHTML_and_Default_A_ATTRIBUTE_DEFINITION_XHTML_REF.X = 1897.999969
	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_ATTRIBUTE_VALUE_XHTML_and_Default_A_ATTRIBUTE_DEFINITION_XHTML_REF.Y = 1055.500000
	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_ATTRIBUTE_VALUE_XHTML_and_Default_A_ATTRIBUTE_DEFINITION_XHTML_REF.Name = `Verticle in class diagram Default in middle between Default-ATTRIBUTE_VALUE_XHTML and Default-A_ATTRIBUTE_DEFINITION_XHTML_REF`

	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_ATTRIBUTE_VALUE_XHTML_and_Default_XHTML_CONTENT.X = 2005.999969
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_ATTRIBUTE_VALUE_XHTML_and_Default_XHTML_CONTENT.Y = 931.000000
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_ATTRIBUTE_VALUE_XHTML_and_Default_XHTML_CONTENT.Name = `Verticle in class diagram Default in middle between Default-ATTRIBUTE_VALUE_XHTML and Default-XHTML_CONTENT`

	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_ATTRIBUTE_VALUE_XHTML_and_Default_XHTML_CONTENT.X = 2005.999969
	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_ATTRIBUTE_VALUE_XHTML_and_Default_XHTML_CONTENT.Y = 931.000000
	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_ATTRIBUTE_VALUE_XHTML_and_Default_XHTML_CONTENT.Name = `Verticle in class diagram Default in middle between Default-ATTRIBUTE_VALUE_XHTML and Default-XHTML_CONTENT`

	__Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_A_ATTRIBUTE_VALUE_XHTML_1_and_Default_ATTRIBUTE_VALUE_BOOLEAN.X = 1795.499939
	__Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_A_ATTRIBUTE_VALUE_XHTML_1_and_Default_ATTRIBUTE_VALUE_BOOLEAN.Y = 787.000000
	__Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_A_ATTRIBUTE_VALUE_XHTML_1_and_Default_ATTRIBUTE_VALUE_BOOLEAN.Name = `Verticle in class diagram Default in middle between Default-A_ATTRIBUTE_VALUE_XHTML_1 and Default-ATTRIBUTE_VALUE_BOOLEAN`

	__Vertice__000004_Verticle_in_class_diagram_Default_in_middle_between_Default_A_ATTRIBUTE_VALUE_XHTML_1_and_Default_ATTRIBUTE_VALUE_DATE.X = 1795.999939
	__Vertice__000004_Verticle_in_class_diagram_Default_in_middle_between_Default_A_ATTRIBUTE_VALUE_XHTML_1_and_Default_ATTRIBUTE_VALUE_DATE.Y = 744.000000
	__Vertice__000004_Verticle_in_class_diagram_Default_in_middle_between_Default_A_ATTRIBUTE_VALUE_XHTML_1_and_Default_ATTRIBUTE_VALUE_DATE.Name = `Verticle in class diagram Default in middle between Default-A_ATTRIBUTE_VALUE_XHTML_1 and Default-ATTRIBUTE_VALUE_DATE`

	__Vertice__000005_Verticle_in_class_diagram_Default_in_middle_between_Default_A_ATTRIBUTE_VALUE_XHTML_1_and_Default_ATTRIBUTE_VALUE_ENUMERATION.X = 1794.999939
	__Vertice__000005_Verticle_in_class_diagram_Default_in_middle_between_Default_A_ATTRIBUTE_VALUE_XHTML_1_and_Default_ATTRIBUTE_VALUE_ENUMERATION.Y = 700.500000
	__Vertice__000005_Verticle_in_class_diagram_Default_in_middle_between_Default_A_ATTRIBUTE_VALUE_XHTML_1_and_Default_ATTRIBUTE_VALUE_ENUMERATION.Name = `Verticle in class diagram Default in middle between Default-A_ATTRIBUTE_VALUE_XHTML_1 and Default-ATTRIBUTE_VALUE_ENUMERATION`

	__Vertice__000006_Verticle_in_class_diagram_Default_in_middle_between_Default_A_ATTRIBUTE_VALUE_XHTML_1_and_Default_ATTRIBUTE_VALUE_INTEGER.X = 1793.499939
	__Vertice__000006_Verticle_in_class_diagram_Default_in_middle_between_Default_A_ATTRIBUTE_VALUE_XHTML_1_and_Default_ATTRIBUTE_VALUE_INTEGER.Y = 657.000000
	__Vertice__000006_Verticle_in_class_diagram_Default_in_middle_between_Default_A_ATTRIBUTE_VALUE_XHTML_1_and_Default_ATTRIBUTE_VALUE_INTEGER.Name = `Verticle in class diagram Default in middle between Default-A_ATTRIBUTE_VALUE_XHTML_1 and Default-ATTRIBUTE_VALUE_INTEGER`

	__Vertice__000007_Verticle_in_class_diagram_Default_in_middle_between_Default_A_ATTRIBUTE_VALUE_XHTML_1_and_Default_ATTRIBUTE_VALUE_REAL.X = 1791.499939
	__Vertice__000007_Verticle_in_class_diagram_Default_in_middle_between_Default_A_ATTRIBUTE_VALUE_XHTML_1_and_Default_ATTRIBUTE_VALUE_REAL.Y = 615.500000
	__Vertice__000007_Verticle_in_class_diagram_Default_in_middle_between_Default_A_ATTRIBUTE_VALUE_XHTML_1_and_Default_ATTRIBUTE_VALUE_REAL.Name = `Verticle in class diagram Default in middle between Default-A_ATTRIBUTE_VALUE_XHTML_1 and Default-ATTRIBUTE_VALUE_REAL`

	__Vertice__000008_Verticle_in_class_diagram_Default_in_middle_between_Default_A_ATTRIBUTE_VALUE_XHTML_1_and_Default_ATTRIBUTE_VALUE_STRING.X = 1790.499939
	__Vertice__000008_Verticle_in_class_diagram_Default_in_middle_between_Default_A_ATTRIBUTE_VALUE_XHTML_1_and_Default_ATTRIBUTE_VALUE_STRING.Y = 573.500000
	__Vertice__000008_Verticle_in_class_diagram_Default_in_middle_between_Default_A_ATTRIBUTE_VALUE_XHTML_1_and_Default_ATTRIBUTE_VALUE_STRING.Name = `Verticle in class diagram Default in middle between Default-A_ATTRIBUTE_VALUE_XHTML_1 and Default-ATTRIBUTE_VALUE_STRING`

	__Vertice__000009_Verticle_in_class_diagram_Default_in_middle_between_Default_A_ATTRIBUTE_VALUE_XHTML_1_and_Default_ATTRIBUTE_VALUE_XHTML.X = 1789.999939
	__Vertice__000009_Verticle_in_class_diagram_Default_in_middle_between_Default_A_ATTRIBUTE_VALUE_XHTML_1_and_Default_ATTRIBUTE_VALUE_XHTML.Y = 528.500000
	__Vertice__000009_Verticle_in_class_diagram_Default_in_middle_between_Default_A_ATTRIBUTE_VALUE_XHTML_1_and_Default_ATTRIBUTE_VALUE_XHTML.Name = `Verticle in class diagram Default in middle between Default-A_ATTRIBUTE_VALUE_XHTML_1 and Default-ATTRIBUTE_VALUE_XHTML`

	__Vertice__000010_Verticle_in_class_diagram_Default_in_middle_between_Default_A_CORE_CONTENT_and_Default_REQ_IF_CONTENT.X = 1020.500000
	__Vertice__000010_Verticle_in_class_diagram_Default_in_middle_between_Default_A_CORE_CONTENT_and_Default_REQ_IF_CONTENT.Y = 207.000000
	__Vertice__000010_Verticle_in_class_diagram_Default_in_middle_between_Default_A_CORE_CONTENT_and_Default_REQ_IF_CONTENT.Name = `Verticle in class diagram Default in middle between Default-A_CORE_CONTENT and Default-REQ_IF_CONTENT`

	__Vertice__000011_Verticle_in_class_diagram_Default_in_middle_between_Default_A_SPEC_OBJECTS_and_Default_SPEC_OBJECT.X = 834.500000
	__Vertice__000011_Verticle_in_class_diagram_Default_in_middle_between_Default_A_SPEC_OBJECTS_and_Default_SPEC_OBJECT.Y = 566.499985
	__Vertice__000011_Verticle_in_class_diagram_Default_in_middle_between_Default_A_SPEC_OBJECTS_and_Default_SPEC_OBJECT.Name = `Verticle in class diagram Default in middle between Default-A_SPEC_OBJECTS and Default-SPEC_OBJECT`

	__Vertice__000012_Verticle_in_class_diagram_Default_in_middle_between_Default_A_THE_HEADER_and_Default_REQ_IF_HEADER.X = 1018.500000
	__Vertice__000012_Verticle_in_class_diagram_Default_in_middle_between_Default_A_THE_HEADER_and_Default_REQ_IF_HEADER.Y = 60.500000
	__Vertice__000012_Verticle_in_class_diagram_Default_in_middle_between_Default_A_THE_HEADER_and_Default_REQ_IF_HEADER.Name = `Verticle in class diagram Default in middle between Default-A_THE_HEADER and Default-REQ_IF_HEADER`

	__Vertice__000013_Verticle_in_class_diagram_Default_in_middle_between_Default_REQ_IF_and_Default_A_CORE_CONTENT.X = 631.500000
	__Vertice__000013_Verticle_in_class_diagram_Default_in_middle_between_Default_REQ_IF_and_Default_A_CORE_CONTENT.Y = 132.500000
	__Vertice__000013_Verticle_in_class_diagram_Default_in_middle_between_Default_REQ_IF_and_Default_A_CORE_CONTENT.Name = `Verticle in class diagram Default in middle between Default-REQ_IF and Default-A_CORE_CONTENT`

	__Vertice__000014_Verticle_in_class_diagram_Default_in_middle_between_Default_REQ_IF_and_Default_A_THE_HEADER.X = 442.500000
	__Vertice__000014_Verticle_in_class_diagram_Default_in_middle_between_Default_REQ_IF_and_Default_A_THE_HEADER.Y = 89.500000
	__Vertice__000014_Verticle_in_class_diagram_Default_in_middle_between_Default_REQ_IF_and_Default_A_THE_HEADER.Name = `Verticle in class diagram Default in middle between Default-REQ_IF and Default-A_THE_HEADER`

	__Vertice__000015_Verticle_in_class_diagram_Default_in_middle_between_Default_REQ_IF_CONTENT_and_Default_A_SPECIFICATIONS.X = 1202.500000
	__Vertice__000015_Verticle_in_class_diagram_Default_in_middle_between_Default_REQ_IF_CONTENT_and_Default_A_SPECIFICATIONS.Y = 347.000000
	__Vertice__000015_Verticle_in_class_diagram_Default_in_middle_between_Default_REQ_IF_CONTENT_and_Default_A_SPECIFICATIONS.Name = `Verticle in class diagram Default in middle between Default-REQ_IF_CONTENT and Default-A_SPECIFICATIONS`

	__Vertice__000016_Verticle_in_class_diagram_Default_in_middle_between_Default_REQ_IF_CONTENT_and_Default_A_SPEC_OBJECTS.X = 1016.000000
	__Vertice__000016_Verticle_in_class_diagram_Default_in_middle_between_Default_REQ_IF_CONTENT_and_Default_A_SPEC_OBJECTS.Y = 344.000000
	__Vertice__000016_Verticle_in_class_diagram_Default_in_middle_between_Default_REQ_IF_CONTENT_and_Default_A_SPEC_OBJECTS.Name = `Verticle in class diagram Default in middle between Default-REQ_IF_CONTENT and Default-A_SPEC_OBJECTS`

	__Vertice__000017_Verticle_in_class_diagram_Default_in_middle_between_Default_REQ_IF_CONTENT_and_Default_A_SPEC_RELATIONS.X = 833.000000
	__Vertice__000017_Verticle_in_class_diagram_Default_in_middle_between_Default_REQ_IF_CONTENT_and_Default_A_SPEC_RELATIONS.Y = 350.000000
	__Vertice__000017_Verticle_in_class_diagram_Default_in_middle_between_Default_REQ_IF_CONTENT_and_Default_A_SPEC_RELATIONS.Name = `Verticle in class diagram Default in middle between Default-REQ_IF_CONTENT and Default-A_SPEC_RELATIONS`

	__Vertice__000018_Verticle_in_class_diagram_Default_in_middle_between_Default_SPEC_OBJECT_and_Default_A_ATTRIBUTE_VALUE_XHTML_1.X = 1144.499969
	__Vertice__000018_Verticle_in_class_diagram_Default_in_middle_between_Default_SPEC_OBJECT_and_Default_A_ATTRIBUTE_VALUE_XHTML_1.Y = 934.499985
	__Vertice__000018_Verticle_in_class_diagram_Default_in_middle_between_Default_SPEC_OBJECT_and_Default_A_ATTRIBUTE_VALUE_XHTML_1.Name = `Verticle in class diagram Default in middle between Default-SPEC_OBJECT and Default-A_ATTRIBUTE_VALUE_XHTML_1`

	__Vertice__000019_Verticle_in_class_diagram_Default_in_middle_between_Default_SPEC_OBJECT_and_Default_A_SPEC_OBJECT_TYPE_REF.X = 1150.499969
	__Vertice__000019_Verticle_in_class_diagram_Default_in_middle_between_Default_SPEC_OBJECT_and_Default_A_SPEC_OBJECT_TYPE_REF.Y = 725.999985
	__Vertice__000019_Verticle_in_class_diagram_Default_in_middle_between_Default_SPEC_OBJECT_and_Default_A_SPEC_OBJECT_TYPE_REF.Name = `Verticle in class diagram Default in middle between Default-SPEC_OBJECT and Default-A_SPEC_OBJECT_TYPE_REF`

	// Setup of pointers
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000024_Default_REQ_IF)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000023_Default_A_THE_HEADER)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000026_Default_REQ_IF_HEADER)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000025_Default_REQ_IF_CONTENT)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000017_Default_A_CORE_CONTENT)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000019_Default_A_SPEC_OBJECTS)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000021_Default_A_SPEC_RELATIONS)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000018_Default_A_SPECIFICATIONS)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000027_Default_SPEC_OBJECT)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000022_Default_A_SPEC_RELATION_TYPE_REF)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000020_Default_A_SPEC_OBJECT_TYPE_REF)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000015_Default_A_ATTRIBUTE_VALUE_XHTML)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000016_Default_A_ATTRIBUTE_VALUE_XHTML_1)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000000_Default_ATTRIBUTE_DEFINITION_BOOLEAN)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000001_Default_ATTRIBUTE_DEFINITION_DATE)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000002_Default_ATTRIBUTE_DEFINITION_ENUMERATION)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000003_Default_ATTRIBUTE_DEFINITION_INTEGER)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000004_Default_ATTRIBUTE_DEFINITION_REAL)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000005_Default_ATTRIBUTE_DEFINITION_STRING)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000007_Default_ATTRIBUTE_VALUE_BOOLEAN)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000008_Default_ATTRIBUTE_VALUE_DATE)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000009_Default_ATTRIBUTE_VALUE_ENUMERATION)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000010_Default_ATTRIBUTE_VALUE_INTEGER)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000011_Default_ATTRIBUTE_VALUE_REAL)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000012_Default_ATTRIBUTE_VALUE_STRING)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000013_Default_ATTRIBUTE_VALUE_XHTML)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000006_Default_ATTRIBUTE_DEFINITION_XHTML)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000029_Default_XHTML_CONTENT)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000014_Default_A_ATTRIBUTE_DEFINITION_XHTML_REF)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000028_Default_SPEC_RELATION)
	__GongStructShape__000000_Default_ATTRIBUTE_DEFINITION_BOOLEAN.Position = __Position__000000_Pos_Default_ATTRIBUTE_DEFINITION_BOOLEAN
	__GongStructShape__000001_Default_ATTRIBUTE_DEFINITION_DATE.Position = __Position__000001_Pos_Default_ATTRIBUTE_DEFINITION_DATE
	__GongStructShape__000002_Default_ATTRIBUTE_DEFINITION_ENUMERATION.Position = __Position__000002_Pos_Default_ATTRIBUTE_DEFINITION_ENUMERATION
	__GongStructShape__000003_Default_ATTRIBUTE_DEFINITION_INTEGER.Position = __Position__000003_Pos_Default_ATTRIBUTE_DEFINITION_INTEGER
	__GongStructShape__000004_Default_ATTRIBUTE_DEFINITION_REAL.Position = __Position__000004_Pos_Default_ATTRIBUTE_DEFINITION_REAL
	__GongStructShape__000005_Default_ATTRIBUTE_DEFINITION_STRING.Position = __Position__000005_Pos_Default_ATTRIBUTE_DEFINITION_STRING
	__GongStructShape__000006_Default_ATTRIBUTE_DEFINITION_XHTML.Position = __Position__000006_Pos_Default_ATTRIBUTE_DEFINITION_XHTML
	__GongStructShape__000007_Default_ATTRIBUTE_VALUE_BOOLEAN.Position = __Position__000007_Pos_Default_ATTRIBUTE_VALUE_BOOLEAN
	__GongStructShape__000008_Default_ATTRIBUTE_VALUE_DATE.Position = __Position__000008_Pos_Default_ATTRIBUTE_VALUE_DATE
	__GongStructShape__000009_Default_ATTRIBUTE_VALUE_ENUMERATION.Position = __Position__000009_Pos_Default_ATTRIBUTE_VALUE_ENUMERATION
	__GongStructShape__000010_Default_ATTRIBUTE_VALUE_INTEGER.Position = __Position__000010_Pos_Default_ATTRIBUTE_VALUE_INTEGER
	__GongStructShape__000011_Default_ATTRIBUTE_VALUE_REAL.Position = __Position__000011_Pos_Default_ATTRIBUTE_VALUE_REAL
	__GongStructShape__000012_Default_ATTRIBUTE_VALUE_STRING.Position = __Position__000012_Pos_Default_ATTRIBUTE_VALUE_STRING
	__GongStructShape__000012_Default_ATTRIBUTE_VALUE_STRING.Fields = append(__GongStructShape__000012_Default_ATTRIBUTE_VALUE_STRING.Fields, __Field__000006_THE_VALUE)
	__GongStructShape__000013_Default_ATTRIBUTE_VALUE_XHTML.Position = __Position__000013_Pos_Default_ATTRIBUTE_VALUE_XHTML
	__GongStructShape__000013_Default_ATTRIBUTE_VALUE_XHTML.Fields = append(__GongStructShape__000013_Default_ATTRIBUTE_VALUE_XHTML.Fields, __Field__000003_IS_SIMPLIFIED)
	__GongStructShape__000013_Default_ATTRIBUTE_VALUE_XHTML.Links = append(__GongStructShape__000013_Default_ATTRIBUTE_VALUE_XHTML.Links, __Link__000016_THE_ORIGINAL_VALUE)
	__GongStructShape__000013_Default_ATTRIBUTE_VALUE_XHTML.Links = append(__GongStructShape__000013_Default_ATTRIBUTE_VALUE_XHTML.Links, __Link__000017_THE_VALUE)
	__GongStructShape__000013_Default_ATTRIBUTE_VALUE_XHTML.Links = append(__GongStructShape__000013_Default_ATTRIBUTE_VALUE_XHTML.Links, __Link__000008_DEFINITION)
	__GongStructShape__000014_Default_A_ATTRIBUTE_DEFINITION_XHTML_REF.Position = __Position__000014_Pos_Default_A_ATTRIBUTE_DEFINITION_XHTML_REF
	__GongStructShape__000014_Default_A_ATTRIBUTE_DEFINITION_XHTML_REF.Fields = append(__GongStructShape__000014_Default_A_ATTRIBUTE_DEFINITION_XHTML_REF.Fields, __Field__000000_ATTRIBUTE_DEFINITION_XHTML_REF)
	__GongStructShape__000015_Default_A_ATTRIBUTE_VALUE_XHTML.Position = __Position__000015_Pos_Default_A_ATTRIBUTE_VALUE_XHTML
	__GongStructShape__000016_Default_A_ATTRIBUTE_VALUE_XHTML_1.Position = __Position__000016_Pos_Default_A_ATTRIBUTE_VALUE_XHTML_1
	__GongStructShape__000016_Default_A_ATTRIBUTE_VALUE_XHTML_1.Links = append(__GongStructShape__000016_Default_A_ATTRIBUTE_VALUE_XHTML_1.Links, __Link__000000_ATTRIBUTE_VALUE_BOOLEAN)
	__GongStructShape__000016_Default_A_ATTRIBUTE_VALUE_XHTML_1.Links = append(__GongStructShape__000016_Default_A_ATTRIBUTE_VALUE_XHTML_1.Links, __Link__000001_ATTRIBUTE_VALUE_DATE)
	__GongStructShape__000016_Default_A_ATTRIBUTE_VALUE_XHTML_1.Links = append(__GongStructShape__000016_Default_A_ATTRIBUTE_VALUE_XHTML_1.Links, __Link__000002_ATTRIBUTE_VALUE_ENUMERATION)
	__GongStructShape__000016_Default_A_ATTRIBUTE_VALUE_XHTML_1.Links = append(__GongStructShape__000016_Default_A_ATTRIBUTE_VALUE_XHTML_1.Links, __Link__000003_ATTRIBUTE_VALUE_INTEGER)
	__GongStructShape__000016_Default_A_ATTRIBUTE_VALUE_XHTML_1.Links = append(__GongStructShape__000016_Default_A_ATTRIBUTE_VALUE_XHTML_1.Links, __Link__000004_ATTRIBUTE_VALUE_REAL)
	__GongStructShape__000016_Default_A_ATTRIBUTE_VALUE_XHTML_1.Links = append(__GongStructShape__000016_Default_A_ATTRIBUTE_VALUE_XHTML_1.Links, __Link__000005_ATTRIBUTE_VALUE_STRING)
	__GongStructShape__000016_Default_A_ATTRIBUTE_VALUE_XHTML_1.Links = append(__GongStructShape__000016_Default_A_ATTRIBUTE_VALUE_XHTML_1.Links, __Link__000006_ATTRIBUTE_VALUE_XHTML)
	__GongStructShape__000017_Default_A_CORE_CONTENT.Position = __Position__000017_Pos_Default_A_CORE_CONTENT
	__GongStructShape__000017_Default_A_CORE_CONTENT.Links = append(__GongStructShape__000017_Default_A_CORE_CONTENT.Links, __Link__000009_REQ_IF_CONTENT)
	__GongStructShape__000018_Default_A_SPECIFICATIONS.Position = __Position__000018_Pos_Default_A_SPECIFICATIONS
	__GongStructShape__000019_Default_A_SPEC_OBJECTS.Position = __Position__000019_Pos_Default_A_SPEC_OBJECTS
	__GongStructShape__000019_Default_A_SPEC_OBJECTS.Links = append(__GongStructShape__000019_Default_A_SPEC_OBJECTS.Links, __Link__000012_SPEC_OBJECT)
	__GongStructShape__000020_Default_A_SPEC_OBJECT_TYPE_REF.Position = __Position__000020_Pos_Default_A_SPEC_OBJECT_TYPE_REF
	__GongStructShape__000021_Default_A_SPEC_RELATIONS.Position = __Position__000021_Pos_Default_A_SPEC_RELATIONS
	__GongStructShape__000022_Default_A_SPEC_RELATION_TYPE_REF.Position = __Position__000022_Pos_Default_A_SPEC_RELATION_TYPE_REF
	__GongStructShape__000023_Default_A_THE_HEADER.Position = __Position__000023_Pos_Default_A_THE_HEADER
	__GongStructShape__000023_Default_A_THE_HEADER.Links = append(__GongStructShape__000023_Default_A_THE_HEADER.Links, __Link__000010_REQ_IF_HEADER)
	__GongStructShape__000024_Default_REQ_IF.Position = __Position__000024_Pos_Default_REQ_IF
	__GongStructShape__000024_Default_REQ_IF.Links = append(__GongStructShape__000024_Default_REQ_IF.Links, __Link__000015_THE_HEADER)
	__GongStructShape__000024_Default_REQ_IF.Links = append(__GongStructShape__000024_Default_REQ_IF.Links, __Link__000007_CORE_CONTENT)
	__GongStructShape__000025_Default_REQ_IF_CONTENT.Position = __Position__000025_Pos_Default_REQ_IF_CONTENT
	__GongStructShape__000025_Default_REQ_IF_CONTENT.Links = append(__GongStructShape__000025_Default_REQ_IF_CONTENT.Links, __Link__000013_SPEC_OBJECTS)
	__GongStructShape__000025_Default_REQ_IF_CONTENT.Links = append(__GongStructShape__000025_Default_REQ_IF_CONTENT.Links, __Link__000014_SPEC_RELATIONS)
	__GongStructShape__000025_Default_REQ_IF_CONTENT.Links = append(__GongStructShape__000025_Default_REQ_IF_CONTENT.Links, __Link__000011_SPECIFICATIONS)
	__GongStructShape__000026_Default_REQ_IF_HEADER.Position = __Position__000026_Pos_Default_REQ_IF_HEADER
	__GongStructShape__000027_Default_SPEC_OBJECT.Position = __Position__000027_Pos_Default_SPEC_OBJECT
	__GongStructShape__000027_Default_SPEC_OBJECT.Fields = append(__GongStructShape__000027_Default_SPEC_OBJECT.Fields, __Field__000001_DESC)
	__GongStructShape__000027_Default_SPEC_OBJECT.Fields = append(__GongStructShape__000027_Default_SPEC_OBJECT.Fields, __Field__000002_IDENTIFIER)
	__GongStructShape__000027_Default_SPEC_OBJECT.Fields = append(__GongStructShape__000027_Default_SPEC_OBJECT.Fields, __Field__000004_LAST_CHANGE)
	__GongStructShape__000027_Default_SPEC_OBJECT.Fields = append(__GongStructShape__000027_Default_SPEC_OBJECT.Fields, __Field__000005_LONG_NAME)
	__GongStructShape__000027_Default_SPEC_OBJECT.Links = append(__GongStructShape__000027_Default_SPEC_OBJECT.Links, __Link__000018_TYPE)
	__GongStructShape__000027_Default_SPEC_OBJECT.Links = append(__GongStructShape__000027_Default_SPEC_OBJECT.Links, __Link__000019_VALUES)
	__GongStructShape__000028_Default_SPEC_RELATION.Position = __Position__000028_Pos_Default_SPEC_RELATION
	__GongStructShape__000029_Default_XHTML_CONTENT.Position = __Position__000029_Pos_Default_XHTML_CONTENT
	__Link__000000_ATTRIBUTE_VALUE_BOOLEAN.Middlevertice = __Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_A_ATTRIBUTE_VALUE_XHTML_1_and_Default_ATTRIBUTE_VALUE_BOOLEAN
	__Link__000001_ATTRIBUTE_VALUE_DATE.Middlevertice = __Vertice__000004_Verticle_in_class_diagram_Default_in_middle_between_Default_A_ATTRIBUTE_VALUE_XHTML_1_and_Default_ATTRIBUTE_VALUE_DATE
	__Link__000002_ATTRIBUTE_VALUE_ENUMERATION.Middlevertice = __Vertice__000005_Verticle_in_class_diagram_Default_in_middle_between_Default_A_ATTRIBUTE_VALUE_XHTML_1_and_Default_ATTRIBUTE_VALUE_ENUMERATION
	__Link__000003_ATTRIBUTE_VALUE_INTEGER.Middlevertice = __Vertice__000006_Verticle_in_class_diagram_Default_in_middle_between_Default_A_ATTRIBUTE_VALUE_XHTML_1_and_Default_ATTRIBUTE_VALUE_INTEGER
	__Link__000004_ATTRIBUTE_VALUE_REAL.Middlevertice = __Vertice__000007_Verticle_in_class_diagram_Default_in_middle_between_Default_A_ATTRIBUTE_VALUE_XHTML_1_and_Default_ATTRIBUTE_VALUE_REAL
	__Link__000005_ATTRIBUTE_VALUE_STRING.Middlevertice = __Vertice__000008_Verticle_in_class_diagram_Default_in_middle_between_Default_A_ATTRIBUTE_VALUE_XHTML_1_and_Default_ATTRIBUTE_VALUE_STRING
	__Link__000006_ATTRIBUTE_VALUE_XHTML.Middlevertice = __Vertice__000009_Verticle_in_class_diagram_Default_in_middle_between_Default_A_ATTRIBUTE_VALUE_XHTML_1_and_Default_ATTRIBUTE_VALUE_XHTML
	__Link__000007_CORE_CONTENT.Middlevertice = __Vertice__000013_Verticle_in_class_diagram_Default_in_middle_between_Default_REQ_IF_and_Default_A_CORE_CONTENT
	__Link__000008_DEFINITION.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_ATTRIBUTE_VALUE_XHTML_and_Default_A_ATTRIBUTE_DEFINITION_XHTML_REF
	__Link__000009_REQ_IF_CONTENT.Middlevertice = __Vertice__000010_Verticle_in_class_diagram_Default_in_middle_between_Default_A_CORE_CONTENT_and_Default_REQ_IF_CONTENT
	__Link__000010_REQ_IF_HEADER.Middlevertice = __Vertice__000012_Verticle_in_class_diagram_Default_in_middle_between_Default_A_THE_HEADER_and_Default_REQ_IF_HEADER
	__Link__000011_SPECIFICATIONS.Middlevertice = __Vertice__000015_Verticle_in_class_diagram_Default_in_middle_between_Default_REQ_IF_CONTENT_and_Default_A_SPECIFICATIONS
	__Link__000012_SPEC_OBJECT.Middlevertice = __Vertice__000011_Verticle_in_class_diagram_Default_in_middle_between_Default_A_SPEC_OBJECTS_and_Default_SPEC_OBJECT
	__Link__000013_SPEC_OBJECTS.Middlevertice = __Vertice__000016_Verticle_in_class_diagram_Default_in_middle_between_Default_REQ_IF_CONTENT_and_Default_A_SPEC_OBJECTS
	__Link__000014_SPEC_RELATIONS.Middlevertice = __Vertice__000017_Verticle_in_class_diagram_Default_in_middle_between_Default_REQ_IF_CONTENT_and_Default_A_SPEC_RELATIONS
	__Link__000015_THE_HEADER.Middlevertice = __Vertice__000014_Verticle_in_class_diagram_Default_in_middle_between_Default_REQ_IF_and_Default_A_THE_HEADER
	__Link__000016_THE_ORIGINAL_VALUE.Middlevertice = __Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_ATTRIBUTE_VALUE_XHTML_and_Default_XHTML_CONTENT
	__Link__000017_THE_VALUE.Middlevertice = __Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_ATTRIBUTE_VALUE_XHTML_and_Default_XHTML_CONTENT
	__Link__000018_TYPE.Middlevertice = __Vertice__000019_Verticle_in_class_diagram_Default_in_middle_between_Default_SPEC_OBJECT_and_Default_A_SPEC_OBJECT_TYPE_REF
	__Link__000019_VALUES.Middlevertice = __Vertice__000018_Verticle_in_class_diagram_Default_in_middle_between_Default_SPEC_OBJECT_and_Default_A_ATTRIBUTE_VALUE_XHTML_1
}
