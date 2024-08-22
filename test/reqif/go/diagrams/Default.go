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

// Injection point for meta package dummy declaration
var _ ref_models.StageStruct

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.StageStruct) {

	// Declaration of instances to stage

	__Classdiagram__000000_Default := (&models.Classdiagram{Name: `Default`}).Stage(stage)

	__Field__000000_ATTRIBUTE_DEFINITION_BOOLEAN_REF := (&models.Field{Name: `ATTRIBUTE_DEFINITION_BOOLEAN_REF`}).Stage(stage)
	__Field__000001_ATTRIBUTE_DEFINITION_DATE_REF := (&models.Field{Name: `ATTRIBUTE_DEFINITION_DATE_REF`}).Stage(stage)
	__Field__000002_ATTRIBUTE_DEFINITION_ENUMERATION_REF := (&models.Field{Name: `ATTRIBUTE_DEFINITION_ENUMERATION_REF`}).Stage(stage)
	__Field__000003_ATTRIBUTE_DEFINITION_INTEGER_REF := (&models.Field{Name: `ATTRIBUTE_DEFINITION_INTEGER_REF`}).Stage(stage)
	__Field__000004_ATTRIBUTE_DEFINITION_REAL_REF := (&models.Field{Name: `ATTRIBUTE_DEFINITION_REAL_REF`}).Stage(stage)
	__Field__000005_ATTRIBUTE_DEFINITION_STRING_REF := (&models.Field{Name: `ATTRIBUTE_DEFINITION_STRING_REF`}).Stage(stage)
	__Field__000006_ATTRIBUTE_DEFINITION_XHTML_REF := (&models.Field{Name: `ATTRIBUTE_DEFINITION_XHTML_REF`}).Stage(stage)
	__Field__000007_COMMENT := (&models.Field{Name: `COMMENT`}).Stage(stage)
	__Field__000008_CREATION_TIME := (&models.Field{Name: `CREATION_TIME`}).Stage(stage)
	__Field__000009_IDENTIFIER := (&models.Field{Name: `IDENTIFIER`}).Stage(stage)
	__Field__000010_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000011_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000012_REPOSITORY_ID := (&models.Field{Name: `REPOSITORY_ID`}).Stage(stage)
	__Field__000013_REQ_IF_TOOL_ID := (&models.Field{Name: `REQ_IF_TOOL_ID`}).Stage(stage)
	__Field__000014_REQ_IF_VERSION := (&models.Field{Name: `REQ_IF_VERSION`}).Stage(stage)
	__Field__000015_SOURCE_TOOL_ID := (&models.Field{Name: `SOURCE_TOOL_ID`}).Stage(stage)
	__Field__000016_TITLE := (&models.Field{Name: `TITLE`}).Stage(stage)

	__GongStructShape__000000_Default_A_CHILDREN := (&models.GongStructShape{Name: `Default-A_CHILDREN`}).Stage(stage)
	__GongStructShape__000001_Default_A_CORE_CONTENT := (&models.GongStructShape{Name: `Default-A_CORE_CONTENT`}).Stage(stage)
	__GongStructShape__000002_Default_A_DATATYPES := (&models.GongStructShape{Name: `Default-A_DATATYPES`}).Stage(stage)
	__GongStructShape__000003_Default_A_DEFINITION := (&models.GongStructShape{Name: `Default-A_DEFINITION`}).Stage(stage)
	__GongStructShape__000004_Default_A_DEFINITION_1 := (&models.GongStructShape{Name: `Default-A_DEFINITION_1`}).Stage(stage)
	__GongStructShape__000005_Default_A_DEFINITION_2 := (&models.GongStructShape{Name: `Default-A_DEFINITION_2`}).Stage(stage)
	__GongStructShape__000006_Default_A_DEFINITION_3 := (&models.GongStructShape{Name: `Default-A_DEFINITION_3`}).Stage(stage)
	__GongStructShape__000007_Default_A_DEFINITION_4 := (&models.GongStructShape{Name: `Default-A_DEFINITION_4`}).Stage(stage)
	__GongStructShape__000008_Default_A_DEFINITION_5 := (&models.GongStructShape{Name: `Default-A_DEFINITION_5`}).Stage(stage)
	__GongStructShape__000009_Default_A_DEFINITION_6 := (&models.GongStructShape{Name: `Default-A_DEFINITION_6`}).Stage(stage)
	__GongStructShape__000010_Default_A_THE_HEADER := (&models.GongStructShape{Name: `Default-A_THE_HEADER`}).Stage(stage)
	__GongStructShape__000011_Default_A_TOOL_EXTENSIONS := (&models.GongStructShape{Name: `Default-A_TOOL_EXTENSIONS`}).Stage(stage)
	__GongStructShape__000012_Default_REQ_IF := (&models.GongStructShape{Name: `Default-REQ_IF`}).Stage(stage)
	__GongStructShape__000013_Default_REQ_IF_CONTENT := (&models.GongStructShape{Name: `Default-REQ_IF_CONTENT`}).Stage(stage)
	__GongStructShape__000014_Default_REQ_IF_HEADER := (&models.GongStructShape{Name: `Default-REQ_IF_HEADER`}).Stage(stage)
	__GongStructShape__000015_Default_REQ_IF_TOOL_EXTENSION := (&models.GongStructShape{Name: `Default-REQ_IF_TOOL_EXTENSION`}).Stage(stage)

	__Link__000000_CORE_CONTENT := (&models.Link{Name: `CORE_CONTENT`}).Stage(stage)
	__Link__000001_REQ_IF_CONTENT := (&models.Link{Name: `REQ_IF_CONTENT`}).Stage(stage)
	__Link__000002_REQ_IF_HEADER := (&models.Link{Name: `REQ_IF_HEADER`}).Stage(stage)
	__Link__000003_REQ_IF_TOOL_EXTENSION := (&models.Link{Name: `REQ_IF_TOOL_EXTENSION`}).Stage(stage)
	__Link__000004_THE_HEADER := (&models.Link{Name: `THE_HEADER`}).Stage(stage)
	__Link__000005_TOOL_EXTENSIONS := (&models.Link{Name: `TOOL_EXTENSIONS`}).Stage(stage)

	__Position__000000_Pos_Default_A_CHILDREN := (&models.Position{Name: `Pos-Default-A_CHILDREN`}).Stage(stage)
	__Position__000001_Pos_Default_A_CORE_CONTENT := (&models.Position{Name: `Pos-Default-A_CORE_CONTENT`}).Stage(stage)
	__Position__000002_Pos_Default_A_DATATYPES := (&models.Position{Name: `Pos-Default-A_DATATYPES`}).Stage(stage)
	__Position__000003_Pos_Default_A_DEFINITION := (&models.Position{Name: `Pos-Default-A_DEFINITION`}).Stage(stage)
	__Position__000004_Pos_Default_A_DEFINITION_1 := (&models.Position{Name: `Pos-Default-A_DEFINITION_1`}).Stage(stage)
	__Position__000005_Pos_Default_A_DEFINITION_2 := (&models.Position{Name: `Pos-Default-A_DEFINITION_2`}).Stage(stage)
	__Position__000006_Pos_Default_A_DEFINITION_3 := (&models.Position{Name: `Pos-Default-A_DEFINITION_3`}).Stage(stage)
	__Position__000007_Pos_Default_A_DEFINITION_4 := (&models.Position{Name: `Pos-Default-A_DEFINITION_4`}).Stage(stage)
	__Position__000008_Pos_Default_A_DEFINITION_5 := (&models.Position{Name: `Pos-Default-A_DEFINITION_5`}).Stage(stage)
	__Position__000009_Pos_Default_A_DEFINITION_6 := (&models.Position{Name: `Pos-Default-A_DEFINITION_6`}).Stage(stage)
	__Position__000010_Pos_Default_A_THE_HEADER := (&models.Position{Name: `Pos-Default-A_THE_HEADER`}).Stage(stage)
	__Position__000011_Pos_Default_A_TOOL_EXTENSIONS := (&models.Position{Name: `Pos-Default-A_TOOL_EXTENSIONS`}).Stage(stage)
	__Position__000012_Pos_Default_REQ_IF := (&models.Position{Name: `Pos-Default-REQ_IF`}).Stage(stage)
	__Position__000013_Pos_Default_REQ_IF_CONTENT := (&models.Position{Name: `Pos-Default-REQ_IF_CONTENT`}).Stage(stage)
	__Position__000014_Pos_Default_REQ_IF_HEADER := (&models.Position{Name: `Pos-Default-REQ_IF_HEADER`}).Stage(stage)
	__Position__000015_Pos_Default_REQ_IF_TOOL_EXTENSION := (&models.Position{Name: `Pos-Default-REQ_IF_TOOL_EXTENSION`}).Stage(stage)

	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_A_CORE_CONTENT_and_Default_REQ_IF_CONTENT := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-A_CORE_CONTENT and Default-REQ_IF_CONTENT`}).Stage(stage)
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_A_THE_HEADER_and_Default_REQ_IF_HEADER := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-A_THE_HEADER and Default-REQ_IF_HEADER`}).Stage(stage)
	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_A_TOOL_EXTENSIONS_and_Default_REQ_IF_TOOL_EXTENSION := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-A_TOOL_EXTENSIONS and Default-REQ_IF_TOOL_EXTENSION`}).Stage(stage)
	__Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_REQ_IF_and_Default_A_CORE_CONTENT := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-REQ_IF and Default-A_CORE_CONTENT`}).Stage(stage)
	__Vertice__000004_Verticle_in_class_diagram_Default_in_middle_between_Default_REQ_IF_and_Default_A_THE_HEADER := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-REQ_IF and Default-A_THE_HEADER`}).Stage(stage)
	__Vertice__000005_Verticle_in_class_diagram_Default_in_middle_between_Default_REQ_IF_and_Default_A_TOOL_EXTENSIONS := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-REQ_IF and Default-A_TOOL_EXTENSIONS`}).Stage(stage)

	// Setup of values

	__Classdiagram__000000_Default.Name = `Default`
	__Classdiagram__000000_Default.IsInDrawMode = false

	__Field__000000_ATTRIBUTE_DEFINITION_BOOLEAN_REF.Name = `ATTRIBUTE_DEFINITION_BOOLEAN_REF`

	//gong:ident [ref_models.A_DEFINITION_4.ATTRIBUTE_DEFINITION_BOOLEAN_REF] comment added to overcome the problem with the comment map association
	__Field__000000_ATTRIBUTE_DEFINITION_BOOLEAN_REF.Identifier = `ref_models.A_DEFINITION_4.ATTRIBUTE_DEFINITION_BOOLEAN_REF`
	__Field__000000_ATTRIBUTE_DEFINITION_BOOLEAN_REF.FieldTypeAsString = ``
	__Field__000000_ATTRIBUTE_DEFINITION_BOOLEAN_REF.Structname = `A_DEFINITION_4`
	__Field__000000_ATTRIBUTE_DEFINITION_BOOLEAN_REF.Fieldtypename = `string`

	__Field__000001_ATTRIBUTE_DEFINITION_DATE_REF.Name = `ATTRIBUTE_DEFINITION_DATE_REF`

	//gong:ident [ref_models.A_DEFINITION_2.ATTRIBUTE_DEFINITION_DATE_REF] comment added to overcome the problem with the comment map association
	__Field__000001_ATTRIBUTE_DEFINITION_DATE_REF.Identifier = `ref_models.A_DEFINITION_2.ATTRIBUTE_DEFINITION_DATE_REF`
	__Field__000001_ATTRIBUTE_DEFINITION_DATE_REF.FieldTypeAsString = ``
	__Field__000001_ATTRIBUTE_DEFINITION_DATE_REF.Structname = `A_DEFINITION_2`
	__Field__000001_ATTRIBUTE_DEFINITION_DATE_REF.Fieldtypename = `string`

	__Field__000002_ATTRIBUTE_DEFINITION_ENUMERATION_REF.Name = `ATTRIBUTE_DEFINITION_ENUMERATION_REF`

	//gong:ident [ref_models.A_DEFINITION_5.ATTRIBUTE_DEFINITION_ENUMERATION_REF] comment added to overcome the problem with the comment map association
	__Field__000002_ATTRIBUTE_DEFINITION_ENUMERATION_REF.Identifier = `ref_models.A_DEFINITION_5.ATTRIBUTE_DEFINITION_ENUMERATION_REF`
	__Field__000002_ATTRIBUTE_DEFINITION_ENUMERATION_REF.FieldTypeAsString = ``
	__Field__000002_ATTRIBUTE_DEFINITION_ENUMERATION_REF.Structname = `A_DEFINITION_5`
	__Field__000002_ATTRIBUTE_DEFINITION_ENUMERATION_REF.Fieldtypename = `string`

	__Field__000003_ATTRIBUTE_DEFINITION_INTEGER_REF.Name = `ATTRIBUTE_DEFINITION_INTEGER_REF`

	//gong:ident [ref_models.A_DEFINITION.ATTRIBUTE_DEFINITION_INTEGER_REF] comment added to overcome the problem with the comment map association
	__Field__000003_ATTRIBUTE_DEFINITION_INTEGER_REF.Identifier = `ref_models.A_DEFINITION.ATTRIBUTE_DEFINITION_INTEGER_REF`
	__Field__000003_ATTRIBUTE_DEFINITION_INTEGER_REF.FieldTypeAsString = ``
	__Field__000003_ATTRIBUTE_DEFINITION_INTEGER_REF.Structname = `A_DEFINITION`
	__Field__000003_ATTRIBUTE_DEFINITION_INTEGER_REF.Fieldtypename = `string`

	__Field__000004_ATTRIBUTE_DEFINITION_REAL_REF.Name = `ATTRIBUTE_DEFINITION_REAL_REF`

	//gong:ident [ref_models.A_DEFINITION_6.ATTRIBUTE_DEFINITION_REAL_REF] comment added to overcome the problem with the comment map association
	__Field__000004_ATTRIBUTE_DEFINITION_REAL_REF.Identifier = `ref_models.A_DEFINITION_6.ATTRIBUTE_DEFINITION_REAL_REF`
	__Field__000004_ATTRIBUTE_DEFINITION_REAL_REF.FieldTypeAsString = ``
	__Field__000004_ATTRIBUTE_DEFINITION_REAL_REF.Structname = `A_DEFINITION_6`
	__Field__000004_ATTRIBUTE_DEFINITION_REAL_REF.Fieldtypename = `string`

	__Field__000005_ATTRIBUTE_DEFINITION_STRING_REF.Name = `ATTRIBUTE_DEFINITION_STRING_REF`

	//gong:ident [ref_models.A_DEFINITION_3.ATTRIBUTE_DEFINITION_STRING_REF] comment added to overcome the problem with the comment map association
	__Field__000005_ATTRIBUTE_DEFINITION_STRING_REF.Identifier = `ref_models.A_DEFINITION_3.ATTRIBUTE_DEFINITION_STRING_REF`
	__Field__000005_ATTRIBUTE_DEFINITION_STRING_REF.FieldTypeAsString = ``
	__Field__000005_ATTRIBUTE_DEFINITION_STRING_REF.Structname = `A_DEFINITION_3`
	__Field__000005_ATTRIBUTE_DEFINITION_STRING_REF.Fieldtypename = `string`

	__Field__000006_ATTRIBUTE_DEFINITION_XHTML_REF.Name = `ATTRIBUTE_DEFINITION_XHTML_REF`

	//gong:ident [ref_models.A_DEFINITION_1.ATTRIBUTE_DEFINITION_XHTML_REF] comment added to overcome the problem with the comment map association
	__Field__000006_ATTRIBUTE_DEFINITION_XHTML_REF.Identifier = `ref_models.A_DEFINITION_1.ATTRIBUTE_DEFINITION_XHTML_REF`
	__Field__000006_ATTRIBUTE_DEFINITION_XHTML_REF.FieldTypeAsString = ``
	__Field__000006_ATTRIBUTE_DEFINITION_XHTML_REF.Structname = `A_DEFINITION_1`
	__Field__000006_ATTRIBUTE_DEFINITION_XHTML_REF.Fieldtypename = `string`

	__Field__000007_COMMENT.Name = `COMMENT`

	//gong:ident [ref_models.REQ_IF_HEADER.COMMENT] comment added to overcome the problem with the comment map association
	__Field__000007_COMMENT.Identifier = `ref_models.REQ_IF_HEADER.COMMENT`
	__Field__000007_COMMENT.FieldTypeAsString = ``
	__Field__000007_COMMENT.Structname = `REQ_IF_HEADER`
	__Field__000007_COMMENT.Fieldtypename = `string`

	__Field__000008_CREATION_TIME.Name = `CREATION_TIME`

	//gong:ident [ref_models.REQ_IF_HEADER.CREATION_TIME] comment added to overcome the problem with the comment map association
	__Field__000008_CREATION_TIME.Identifier = `ref_models.REQ_IF_HEADER.CREATION_TIME`
	__Field__000008_CREATION_TIME.FieldTypeAsString = ``
	__Field__000008_CREATION_TIME.Structname = `REQ_IF_HEADER`
	__Field__000008_CREATION_TIME.Fieldtypename = `string`

	__Field__000009_IDENTIFIER.Name = `IDENTIFIER`

	//gong:ident [ref_models.REQ_IF_HEADER.IDENTIFIER] comment added to overcome the problem with the comment map association
	__Field__000009_IDENTIFIER.Identifier = `ref_models.REQ_IF_HEADER.IDENTIFIER`
	__Field__000009_IDENTIFIER.FieldTypeAsString = ``
	__Field__000009_IDENTIFIER.Structname = `REQ_IF_HEADER`
	__Field__000009_IDENTIFIER.Fieldtypename = `string`

	__Field__000010_Name.Name = `Name`

	//gong:ident [ref_models.REQ_IF_CONTENT.Name] comment added to overcome the problem with the comment map association
	__Field__000010_Name.Identifier = `ref_models.REQ_IF_CONTENT.Name`
	__Field__000010_Name.FieldTypeAsString = ``
	__Field__000010_Name.Structname = `REQ_IF_CONTENT`
	__Field__000010_Name.Fieldtypename = `string`

	__Field__000011_Name.Name = `Name`

	//gong:ident [ref_models.REQ_IF.Name] comment added to overcome the problem with the comment map association
	__Field__000011_Name.Identifier = `ref_models.REQ_IF.Name`
	__Field__000011_Name.FieldTypeAsString = ``
	__Field__000011_Name.Structname = `REQ_IF`
	__Field__000011_Name.Fieldtypename = `string`

	__Field__000012_REPOSITORY_ID.Name = `REPOSITORY_ID`

	//gong:ident [ref_models.REQ_IF_HEADER.REPOSITORY_ID] comment added to overcome the problem with the comment map association
	__Field__000012_REPOSITORY_ID.Identifier = `ref_models.REQ_IF_HEADER.REPOSITORY_ID`
	__Field__000012_REPOSITORY_ID.FieldTypeAsString = ``
	__Field__000012_REPOSITORY_ID.Structname = `REQ_IF_HEADER`
	__Field__000012_REPOSITORY_ID.Fieldtypename = `string`

	__Field__000013_REQ_IF_TOOL_ID.Name = `REQ_IF_TOOL_ID`

	//gong:ident [ref_models.REQ_IF_HEADER.REQ_IF_TOOL_ID] comment added to overcome the problem with the comment map association
	__Field__000013_REQ_IF_TOOL_ID.Identifier = `ref_models.REQ_IF_HEADER.REQ_IF_TOOL_ID`
	__Field__000013_REQ_IF_TOOL_ID.FieldTypeAsString = ``
	__Field__000013_REQ_IF_TOOL_ID.Structname = `REQ_IF_HEADER`
	__Field__000013_REQ_IF_TOOL_ID.Fieldtypename = `string`

	__Field__000014_REQ_IF_VERSION.Name = `REQ_IF_VERSION`

	//gong:ident [ref_models.REQ_IF_HEADER.REQ_IF_VERSION] comment added to overcome the problem with the comment map association
	__Field__000014_REQ_IF_VERSION.Identifier = `ref_models.REQ_IF_HEADER.REQ_IF_VERSION`
	__Field__000014_REQ_IF_VERSION.FieldTypeAsString = ``
	__Field__000014_REQ_IF_VERSION.Structname = `REQ_IF_HEADER`
	__Field__000014_REQ_IF_VERSION.Fieldtypename = `string`

	__Field__000015_SOURCE_TOOL_ID.Name = `SOURCE_TOOL_ID`

	//gong:ident [ref_models.REQ_IF_HEADER.SOURCE_TOOL_ID] comment added to overcome the problem with the comment map association
	__Field__000015_SOURCE_TOOL_ID.Identifier = `ref_models.REQ_IF_HEADER.SOURCE_TOOL_ID`
	__Field__000015_SOURCE_TOOL_ID.FieldTypeAsString = ``
	__Field__000015_SOURCE_TOOL_ID.Structname = `REQ_IF_HEADER`
	__Field__000015_SOURCE_TOOL_ID.Fieldtypename = `string`

	__Field__000016_TITLE.Name = `TITLE`

	//gong:ident [ref_models.REQ_IF_HEADER.TITLE] comment added to overcome the problem with the comment map association
	__Field__000016_TITLE.Identifier = `ref_models.REQ_IF_HEADER.TITLE`
	__Field__000016_TITLE.FieldTypeAsString = ``
	__Field__000016_TITLE.Structname = `REQ_IF_HEADER`
	__Field__000016_TITLE.Fieldtypename = `string`

	__GongStructShape__000000_Default_A_CHILDREN.Name = `Default-A_CHILDREN`

	//gong:ident [ref_models.A_CHILDREN] comment added to overcome the problem with the comment map association
	__GongStructShape__000000_Default_A_CHILDREN.Identifier = `ref_models.A_CHILDREN`
	__GongStructShape__000000_Default_A_CHILDREN.ShowNbInstances = false
	__GongStructShape__000000_Default_A_CHILDREN.NbInstances = 0
	__GongStructShape__000000_Default_A_CHILDREN.Width = 240.000000
	__GongStructShape__000000_Default_A_CHILDREN.Height = 63.000000
	__GongStructShape__000000_Default_A_CHILDREN.IsSelected = false

	__GongStructShape__000001_Default_A_CORE_CONTENT.Name = `Default-A_CORE_CONTENT`

	//gong:ident [ref_models.A_CORE_CONTENT] comment added to overcome the problem with the comment map association
	__GongStructShape__000001_Default_A_CORE_CONTENT.Identifier = `ref_models.A_CORE_CONTENT`
	__GongStructShape__000001_Default_A_CORE_CONTENT.ShowNbInstances = false
	__GongStructShape__000001_Default_A_CORE_CONTENT.NbInstances = 0
	__GongStructShape__000001_Default_A_CORE_CONTENT.Width = 240.000000
	__GongStructShape__000001_Default_A_CORE_CONTENT.Height = 63.000000
	__GongStructShape__000001_Default_A_CORE_CONTENT.IsSelected = false

	__GongStructShape__000002_Default_A_DATATYPES.Name = `Default-A_DATATYPES`

	//gong:ident [ref_models.A_DATATYPES] comment added to overcome the problem with the comment map association
	__GongStructShape__000002_Default_A_DATATYPES.Identifier = `ref_models.A_DATATYPES`
	__GongStructShape__000002_Default_A_DATATYPES.ShowNbInstances = false
	__GongStructShape__000002_Default_A_DATATYPES.NbInstances = 0
	__GongStructShape__000002_Default_A_DATATYPES.Width = 240.000000
	__GongStructShape__000002_Default_A_DATATYPES.Height = 63.000000
	__GongStructShape__000002_Default_A_DATATYPES.IsSelected = false

	__GongStructShape__000003_Default_A_DEFINITION.Name = `Default-A_DEFINITION`

	//gong:ident [ref_models.A_DEFINITION] comment added to overcome the problem with the comment map association
	__GongStructShape__000003_Default_A_DEFINITION.Identifier = `ref_models.A_DEFINITION`
	__GongStructShape__000003_Default_A_DEFINITION.ShowNbInstances = false
	__GongStructShape__000003_Default_A_DEFINITION.NbInstances = 0
	__GongStructShape__000003_Default_A_DEFINITION.Width = 240.000000
	__GongStructShape__000003_Default_A_DEFINITION.Height = 78.000000
	__GongStructShape__000003_Default_A_DEFINITION.IsSelected = false

	__GongStructShape__000004_Default_A_DEFINITION_1.Name = `Default-A_DEFINITION_1`

	//gong:ident [ref_models.A_DEFINITION_1] comment added to overcome the problem with the comment map association
	__GongStructShape__000004_Default_A_DEFINITION_1.Identifier = `ref_models.A_DEFINITION_1`
	__GongStructShape__000004_Default_A_DEFINITION_1.ShowNbInstances = false
	__GongStructShape__000004_Default_A_DEFINITION_1.NbInstances = 0
	__GongStructShape__000004_Default_A_DEFINITION_1.Width = 240.000000
	__GongStructShape__000004_Default_A_DEFINITION_1.Height = 78.000000
	__GongStructShape__000004_Default_A_DEFINITION_1.IsSelected = false

	__GongStructShape__000005_Default_A_DEFINITION_2.Name = `Default-A_DEFINITION_2`

	//gong:ident [ref_models.A_DEFINITION_2] comment added to overcome the problem with the comment map association
	__GongStructShape__000005_Default_A_DEFINITION_2.Identifier = `ref_models.A_DEFINITION_2`
	__GongStructShape__000005_Default_A_DEFINITION_2.ShowNbInstances = false
	__GongStructShape__000005_Default_A_DEFINITION_2.NbInstances = 0
	__GongStructShape__000005_Default_A_DEFINITION_2.Width = 240.000000
	__GongStructShape__000005_Default_A_DEFINITION_2.Height = 78.000000
	__GongStructShape__000005_Default_A_DEFINITION_2.IsSelected = false

	__GongStructShape__000006_Default_A_DEFINITION_3.Name = `Default-A_DEFINITION_3`

	//gong:ident [ref_models.A_DEFINITION_3] comment added to overcome the problem with the comment map association
	__GongStructShape__000006_Default_A_DEFINITION_3.Identifier = `ref_models.A_DEFINITION_3`
	__GongStructShape__000006_Default_A_DEFINITION_3.ShowNbInstances = false
	__GongStructShape__000006_Default_A_DEFINITION_3.NbInstances = 0
	__GongStructShape__000006_Default_A_DEFINITION_3.Width = 240.000000
	__GongStructShape__000006_Default_A_DEFINITION_3.Height = 78.000000
	__GongStructShape__000006_Default_A_DEFINITION_3.IsSelected = false

	__GongStructShape__000007_Default_A_DEFINITION_4.Name = `Default-A_DEFINITION_4`

	//gong:ident [ref_models.A_DEFINITION_4] comment added to overcome the problem with the comment map association
	__GongStructShape__000007_Default_A_DEFINITION_4.Identifier = `ref_models.A_DEFINITION_4`
	__GongStructShape__000007_Default_A_DEFINITION_4.ShowNbInstances = false
	__GongStructShape__000007_Default_A_DEFINITION_4.NbInstances = 0
	__GongStructShape__000007_Default_A_DEFINITION_4.Width = 240.000000
	__GongStructShape__000007_Default_A_DEFINITION_4.Height = 78.000000
	__GongStructShape__000007_Default_A_DEFINITION_4.IsSelected = false

	__GongStructShape__000008_Default_A_DEFINITION_5.Name = `Default-A_DEFINITION_5`

	//gong:ident [ref_models.A_DEFINITION_5] comment added to overcome the problem with the comment map association
	__GongStructShape__000008_Default_A_DEFINITION_5.Identifier = `ref_models.A_DEFINITION_5`
	__GongStructShape__000008_Default_A_DEFINITION_5.ShowNbInstances = false
	__GongStructShape__000008_Default_A_DEFINITION_5.NbInstances = 0
	__GongStructShape__000008_Default_A_DEFINITION_5.Width = 240.000000
	__GongStructShape__000008_Default_A_DEFINITION_5.Height = 78.000000
	__GongStructShape__000008_Default_A_DEFINITION_5.IsSelected = false

	__GongStructShape__000009_Default_A_DEFINITION_6.Name = `Default-A_DEFINITION_6`

	//gong:ident [ref_models.A_DEFINITION_6] comment added to overcome the problem with the comment map association
	__GongStructShape__000009_Default_A_DEFINITION_6.Identifier = `ref_models.A_DEFINITION_6`
	__GongStructShape__000009_Default_A_DEFINITION_6.ShowNbInstances = false
	__GongStructShape__000009_Default_A_DEFINITION_6.NbInstances = 0
	__GongStructShape__000009_Default_A_DEFINITION_6.Width = 240.000000
	__GongStructShape__000009_Default_A_DEFINITION_6.Height = 78.000000
	__GongStructShape__000009_Default_A_DEFINITION_6.IsSelected = false

	__GongStructShape__000010_Default_A_THE_HEADER.Name = `Default-A_THE_HEADER`

	//gong:ident [ref_models.A_THE_HEADER] comment added to overcome the problem with the comment map association
	__GongStructShape__000010_Default_A_THE_HEADER.Identifier = `ref_models.A_THE_HEADER`
	__GongStructShape__000010_Default_A_THE_HEADER.ShowNbInstances = false
	__GongStructShape__000010_Default_A_THE_HEADER.NbInstances = 0
	__GongStructShape__000010_Default_A_THE_HEADER.Width = 240.000000
	__GongStructShape__000010_Default_A_THE_HEADER.Height = 63.000000
	__GongStructShape__000010_Default_A_THE_HEADER.IsSelected = false

	__GongStructShape__000011_Default_A_TOOL_EXTENSIONS.Name = `Default-A_TOOL_EXTENSIONS`

	//gong:ident [ref_models.A_TOOL_EXTENSIONS] comment added to overcome the problem with the comment map association
	__GongStructShape__000011_Default_A_TOOL_EXTENSIONS.Identifier = `ref_models.A_TOOL_EXTENSIONS`
	__GongStructShape__000011_Default_A_TOOL_EXTENSIONS.ShowNbInstances = false
	__GongStructShape__000011_Default_A_TOOL_EXTENSIONS.NbInstances = 0
	__GongStructShape__000011_Default_A_TOOL_EXTENSIONS.Width = 240.000000
	__GongStructShape__000011_Default_A_TOOL_EXTENSIONS.Height = 63.000000
	__GongStructShape__000011_Default_A_TOOL_EXTENSIONS.IsSelected = false

	__GongStructShape__000012_Default_REQ_IF.Name = `Default-REQ_IF`

	//gong:ident [ref_models.REQ_IF] comment added to overcome the problem with the comment map association
	__GongStructShape__000012_Default_REQ_IF.Identifier = `ref_models.REQ_IF`
	__GongStructShape__000012_Default_REQ_IF.ShowNbInstances = false
	__GongStructShape__000012_Default_REQ_IF.NbInstances = 0
	__GongStructShape__000012_Default_REQ_IF.Width = 240.000000
	__GongStructShape__000012_Default_REQ_IF.Height = 78.000000
	__GongStructShape__000012_Default_REQ_IF.IsSelected = false

	__GongStructShape__000013_Default_REQ_IF_CONTENT.Name = `Default-REQ_IF_CONTENT`

	//gong:ident [ref_models.REQ_IF_CONTENT] comment added to overcome the problem with the comment map association
	__GongStructShape__000013_Default_REQ_IF_CONTENT.Identifier = `ref_models.REQ_IF_CONTENT`
	__GongStructShape__000013_Default_REQ_IF_CONTENT.ShowNbInstances = false
	__GongStructShape__000013_Default_REQ_IF_CONTENT.NbInstances = 0
	__GongStructShape__000013_Default_REQ_IF_CONTENT.Width = 240.000000
	__GongStructShape__000013_Default_REQ_IF_CONTENT.Height = 78.000000
	__GongStructShape__000013_Default_REQ_IF_CONTENT.IsSelected = false

	__GongStructShape__000014_Default_REQ_IF_HEADER.Name = `Default-REQ_IF_HEADER`

	//gong:ident [ref_models.REQ_IF_HEADER] comment added to overcome the problem with the comment map association
	__GongStructShape__000014_Default_REQ_IF_HEADER.Identifier = `ref_models.REQ_IF_HEADER`
	__GongStructShape__000014_Default_REQ_IF_HEADER.ShowNbInstances = false
	__GongStructShape__000014_Default_REQ_IF_HEADER.NbInstances = 0
	__GongStructShape__000014_Default_REQ_IF_HEADER.Width = 240.000000
	__GongStructShape__000014_Default_REQ_IF_HEADER.Height = 183.000000
	__GongStructShape__000014_Default_REQ_IF_HEADER.IsSelected = false

	__GongStructShape__000015_Default_REQ_IF_TOOL_EXTENSION.Name = `Default-REQ_IF_TOOL_EXTENSION`

	//gong:ident [ref_models.REQ_IF_TOOL_EXTENSION] comment added to overcome the problem with the comment map association
	__GongStructShape__000015_Default_REQ_IF_TOOL_EXTENSION.Identifier = `ref_models.REQ_IF_TOOL_EXTENSION`
	__GongStructShape__000015_Default_REQ_IF_TOOL_EXTENSION.ShowNbInstances = false
	__GongStructShape__000015_Default_REQ_IF_TOOL_EXTENSION.NbInstances = 0
	__GongStructShape__000015_Default_REQ_IF_TOOL_EXTENSION.Width = 240.000000
	__GongStructShape__000015_Default_REQ_IF_TOOL_EXTENSION.Height = 63.000000
	__GongStructShape__000015_Default_REQ_IF_TOOL_EXTENSION.IsSelected = false

	__Link__000000_CORE_CONTENT.Name = `CORE_CONTENT`

	//gong:ident [ref_models.REQ_IF.CORE_CONTENT] comment added to overcome the problem with the comment map association
	__Link__000000_CORE_CONTENT.Identifier = `ref_models.REQ_IF.CORE_CONTENT`

	//gong:ident [ref_models.A_CORE_CONTENT] comment added to overcome the problem with the comment map association
	__Link__000000_CORE_CONTENT.Fieldtypename = `ref_models.A_CORE_CONTENT`
	__Link__000000_CORE_CONTENT.FieldOffsetX = -50.000000
	__Link__000000_CORE_CONTENT.FieldOffsetY = -16.000000
	__Link__000000_CORE_CONTENT.TargetMultiplicity = models.MANY
	__Link__000000_CORE_CONTENT.TargetMultiplicityOffsetX = -50.000000
	__Link__000000_CORE_CONTENT.TargetMultiplicityOffsetY = 16.000000
	__Link__000000_CORE_CONTENT.SourceMultiplicity = models.MANY
	__Link__000000_CORE_CONTENT.SourceMultiplicityOffsetX = 10.000000
	__Link__000000_CORE_CONTENT.SourceMultiplicityOffsetY = -50.000000
	__Link__000000_CORE_CONTENT.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_CORE_CONTENT.StartRatio = 0.500000
	__Link__000000_CORE_CONTENT.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_CORE_CONTENT.EndRatio = 0.500000
	__Link__000000_CORE_CONTENT.CornerOffsetRatio = 1.380000

	__Link__000001_REQ_IF_CONTENT.Name = `REQ_IF_CONTENT`

	//gong:ident [ref_models.A_CORE_CONTENT.REQ_IF_CONTENT] comment added to overcome the problem with the comment map association
	__Link__000001_REQ_IF_CONTENT.Identifier = `ref_models.A_CORE_CONTENT.REQ_IF_CONTENT`

	//gong:ident [ref_models.REQ_IF_CONTENT] comment added to overcome the problem with the comment map association
	__Link__000001_REQ_IF_CONTENT.Fieldtypename = `ref_models.REQ_IF_CONTENT`
	__Link__000001_REQ_IF_CONTENT.FieldOffsetX = -50.000000
	__Link__000001_REQ_IF_CONTENT.FieldOffsetY = -16.000000
	__Link__000001_REQ_IF_CONTENT.TargetMultiplicity = models.MANY
	__Link__000001_REQ_IF_CONTENT.TargetMultiplicityOffsetX = -50.000000
	__Link__000001_REQ_IF_CONTENT.TargetMultiplicityOffsetY = 16.000000
	__Link__000001_REQ_IF_CONTENT.SourceMultiplicity = models.MANY
	__Link__000001_REQ_IF_CONTENT.SourceMultiplicityOffsetX = 10.000000
	__Link__000001_REQ_IF_CONTENT.SourceMultiplicityOffsetY = -50.000000
	__Link__000001_REQ_IF_CONTENT.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000001_REQ_IF_CONTENT.StartRatio = 0.500000
	__Link__000001_REQ_IF_CONTENT.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000001_REQ_IF_CONTENT.EndRatio = 0.500000
	__Link__000001_REQ_IF_CONTENT.CornerOffsetRatio = 1.380000

	__Link__000002_REQ_IF_HEADER.Name = `REQ_IF_HEADER`

	//gong:ident [ref_models.A_THE_HEADER.REQ_IF_HEADER] comment added to overcome the problem with the comment map association
	__Link__000002_REQ_IF_HEADER.Identifier = `ref_models.A_THE_HEADER.REQ_IF_HEADER`

	//gong:ident [ref_models.REQ_IF_HEADER] comment added to overcome the problem with the comment map association
	__Link__000002_REQ_IF_HEADER.Fieldtypename = `ref_models.REQ_IF_HEADER`
	__Link__000002_REQ_IF_HEADER.FieldOffsetX = -50.000000
	__Link__000002_REQ_IF_HEADER.FieldOffsetY = -16.000000
	__Link__000002_REQ_IF_HEADER.TargetMultiplicity = models.MANY
	__Link__000002_REQ_IF_HEADER.TargetMultiplicityOffsetX = -50.000000
	__Link__000002_REQ_IF_HEADER.TargetMultiplicityOffsetY = 16.000000
	__Link__000002_REQ_IF_HEADER.SourceMultiplicity = models.MANY
	__Link__000002_REQ_IF_HEADER.SourceMultiplicityOffsetX = 10.000000
	__Link__000002_REQ_IF_HEADER.SourceMultiplicityOffsetY = -50.000000
	__Link__000002_REQ_IF_HEADER.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000002_REQ_IF_HEADER.StartRatio = 0.500000
	__Link__000002_REQ_IF_HEADER.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000002_REQ_IF_HEADER.EndRatio = 0.500000
	__Link__000002_REQ_IF_HEADER.CornerOffsetRatio = 1.380000

	__Link__000003_REQ_IF_TOOL_EXTENSION.Name = `REQ_IF_TOOL_EXTENSION`

	//gong:ident [ref_models.A_TOOL_EXTENSIONS.REQ_IF_TOOL_EXTENSION] comment added to overcome the problem with the comment map association
	__Link__000003_REQ_IF_TOOL_EXTENSION.Identifier = `ref_models.A_TOOL_EXTENSIONS.REQ_IF_TOOL_EXTENSION`

	//gong:ident [ref_models.REQ_IF_TOOL_EXTENSION] comment added to overcome the problem with the comment map association
	__Link__000003_REQ_IF_TOOL_EXTENSION.Fieldtypename = `ref_models.REQ_IF_TOOL_EXTENSION`
	__Link__000003_REQ_IF_TOOL_EXTENSION.FieldOffsetX = -50.000000
	__Link__000003_REQ_IF_TOOL_EXTENSION.FieldOffsetY = -16.000000
	__Link__000003_REQ_IF_TOOL_EXTENSION.TargetMultiplicity = models.MANY
	__Link__000003_REQ_IF_TOOL_EXTENSION.TargetMultiplicityOffsetX = -50.000000
	__Link__000003_REQ_IF_TOOL_EXTENSION.TargetMultiplicityOffsetY = 16.000000
	__Link__000003_REQ_IF_TOOL_EXTENSION.SourceMultiplicity = models.MANY
	__Link__000003_REQ_IF_TOOL_EXTENSION.SourceMultiplicityOffsetX = 10.000000
	__Link__000003_REQ_IF_TOOL_EXTENSION.SourceMultiplicityOffsetY = -50.000000
	__Link__000003_REQ_IF_TOOL_EXTENSION.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000003_REQ_IF_TOOL_EXTENSION.StartRatio = 0.500000
	__Link__000003_REQ_IF_TOOL_EXTENSION.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000003_REQ_IF_TOOL_EXTENSION.EndRatio = 0.500000
	__Link__000003_REQ_IF_TOOL_EXTENSION.CornerOffsetRatio = 1.380000

	__Link__000004_THE_HEADER.Name = `THE_HEADER`

	//gong:ident [ref_models.REQ_IF.THE_HEADER] comment added to overcome the problem with the comment map association
	__Link__000004_THE_HEADER.Identifier = `ref_models.REQ_IF.THE_HEADER`

	//gong:ident [ref_models.A_THE_HEADER] comment added to overcome the problem with the comment map association
	__Link__000004_THE_HEADER.Fieldtypename = `ref_models.A_THE_HEADER`
	__Link__000004_THE_HEADER.FieldOffsetX = -50.000000
	__Link__000004_THE_HEADER.FieldOffsetY = -16.000000
	__Link__000004_THE_HEADER.TargetMultiplicity = models.MANY
	__Link__000004_THE_HEADER.TargetMultiplicityOffsetX = -50.000000
	__Link__000004_THE_HEADER.TargetMultiplicityOffsetY = 16.000000
	__Link__000004_THE_HEADER.SourceMultiplicity = models.MANY
	__Link__000004_THE_HEADER.SourceMultiplicityOffsetX = 10.000000
	__Link__000004_THE_HEADER.SourceMultiplicityOffsetY = -50.000000
	__Link__000004_THE_HEADER.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000004_THE_HEADER.StartRatio = 0.500000
	__Link__000004_THE_HEADER.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000004_THE_HEADER.EndRatio = 0.500000
	__Link__000004_THE_HEADER.CornerOffsetRatio = 1.380000

	__Link__000005_TOOL_EXTENSIONS.Name = `TOOL_EXTENSIONS`

	//gong:ident [ref_models.REQ_IF.TOOL_EXTENSIONS] comment added to overcome the problem with the comment map association
	__Link__000005_TOOL_EXTENSIONS.Identifier = `ref_models.REQ_IF.TOOL_EXTENSIONS`

	//gong:ident [ref_models.A_TOOL_EXTENSIONS] comment added to overcome the problem with the comment map association
	__Link__000005_TOOL_EXTENSIONS.Fieldtypename = `ref_models.A_TOOL_EXTENSIONS`
	__Link__000005_TOOL_EXTENSIONS.FieldOffsetX = -50.000000
	__Link__000005_TOOL_EXTENSIONS.FieldOffsetY = -16.000000
	__Link__000005_TOOL_EXTENSIONS.TargetMultiplicity = models.MANY
	__Link__000005_TOOL_EXTENSIONS.TargetMultiplicityOffsetX = -50.000000
	__Link__000005_TOOL_EXTENSIONS.TargetMultiplicityOffsetY = 16.000000
	__Link__000005_TOOL_EXTENSIONS.SourceMultiplicity = models.MANY
	__Link__000005_TOOL_EXTENSIONS.SourceMultiplicityOffsetX = 10.000000
	__Link__000005_TOOL_EXTENSIONS.SourceMultiplicityOffsetY = -50.000000
	__Link__000005_TOOL_EXTENSIONS.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000005_TOOL_EXTENSIONS.StartRatio = 0.500000
	__Link__000005_TOOL_EXTENSIONS.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000005_TOOL_EXTENSIONS.EndRatio = 0.500000
	__Link__000005_TOOL_EXTENSIONS.CornerOffsetRatio = 1.380000

	__Position__000000_Pos_Default_A_CHILDREN.X = 776.000000
	__Position__000000_Pos_Default_A_CHILDREN.Y = 967.999969
	__Position__000000_Pos_Default_A_CHILDREN.Name = `Pos-Default-A_CHILDREN`

	__Position__000001_Pos_Default_A_CORE_CONTENT.X = 592.000000
	__Position__000001_Pos_Default_A_CORE_CONTENT.Y = 663.999969
	__Position__000001_Pos_Default_A_CORE_CONTENT.Name = `Pos-Default-A_CORE_CONTENT`

	__Position__000002_Pos_Default_A_DATATYPES.X = 172.000000
	__Position__000002_Pos_Default_A_DATATYPES.Y = 918.999969
	__Position__000002_Pos_Default_A_DATATYPES.Name = `Pos-Default-A_DATATYPES`

	__Position__000003_Pos_Default_A_DEFINITION.X = 396.000000
	__Position__000003_Pos_Default_A_DEFINITION.Y = 823.999969
	__Position__000003_Pos_Default_A_DEFINITION.Name = `Pos-Default-A_DEFINITION`

	__Position__000004_Pos_Default_A_DEFINITION_1.X = 799.000000
	__Position__000004_Pos_Default_A_DEFINITION_1.Y = 850.999969
	__Position__000004_Pos_Default_A_DEFINITION_1.Name = `Pos-Default-A_DEFINITION_1`

	__Position__000005_Pos_Default_A_DEFINITION_2.X = 429.000000
	__Position__000005_Pos_Default_A_DEFINITION_2.Y = 951.999969
	__Position__000005_Pos_Default_A_DEFINITION_2.Name = `Pos-Default-A_DEFINITION_2`

	__Position__000006_Pos_Default_A_DEFINITION_3.X = 149.000000
	__Position__000006_Pos_Default_A_DEFINITION_3.Y = 1134.999969
	__Position__000006_Pos_Default_A_DEFINITION_3.Name = `Pos-Default-A_DEFINITION_3`

	__Position__000007_Pos_Default_A_DEFINITION_4.X = 164.000000
	__Position__000007_Pos_Default_A_DEFINITION_4.Y = 1246.999969
	__Position__000007_Pos_Default_A_DEFINITION_4.Name = `Pos-Default-A_DEFINITION_4`

	__Position__000008_Pos_Default_A_DEFINITION_5.X = 180.000000
	__Position__000008_Pos_Default_A_DEFINITION_5.Y = 1357.999969
	__Position__000008_Pos_Default_A_DEFINITION_5.Name = `Pos-Default-A_DEFINITION_5`

	__Position__000009_Pos_Default_A_DEFINITION_6.X = 1063.999939
	__Position__000009_Pos_Default_A_DEFINITION_6.Y = 995.333344
	__Position__000009_Pos_Default_A_DEFINITION_6.Name = `Pos-Default-A_DEFINITION_6`

	__Position__000010_Pos_Default_A_THE_HEADER.X = 575.000000
	__Position__000010_Pos_Default_A_THE_HEADER.Y = 74.000000
	__Position__000010_Pos_Default_A_THE_HEADER.Name = `Pos-Default-A_THE_HEADER`

	__Position__000011_Pos_Default_A_TOOL_EXTENSIONS.X = 567.000000
	__Position__000011_Pos_Default_A_TOOL_EXTENSIONS.Y = 436.000000
	__Position__000011_Pos_Default_A_TOOL_EXTENSIONS.Name = `Pos-Default-A_TOOL_EXTENSIONS`

	__Position__000012_Pos_Default_REQ_IF.X = 89.000000
	__Position__000012_Pos_Default_REQ_IF.Y = 65.000000
	__Position__000012_Pos_Default_REQ_IF.Name = `Pos-Default-REQ_IF`

	__Position__000013_Pos_Default_REQ_IF_CONTENT.X = 1068.999939
	__Position__000013_Pos_Default_REQ_IF_CONTENT.Y = 657.999969
	__Position__000013_Pos_Default_REQ_IF_CONTENT.Name = `Pos-Default-REQ_IF_CONTENT`

	__Position__000014_Pos_Default_REQ_IF_HEADER.X = 1063.999939
	__Position__000014_Pos_Default_REQ_IF_HEADER.Y = 14.000000
	__Position__000014_Pos_Default_REQ_IF_HEADER.Name = `Pos-Default-REQ_IF_HEADER`

	__Position__000015_Pos_Default_REQ_IF_TOOL_EXTENSION.X = 1061.999939
	__Position__000015_Pos_Default_REQ_IF_TOOL_EXTENSION.Y = 429.000000
	__Position__000015_Pos_Default_REQ_IF_TOOL_EXTENSION.Name = `Pos-Default-REQ_IF_TOOL_EXTENSION`

	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_A_CORE_CONTENT_and_Default_REQ_IF_CONTENT.X = 1176.499969
	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_A_CORE_CONTENT_and_Default_REQ_IF_CONTENT.Y = 220.500000
	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_A_CORE_CONTENT_and_Default_REQ_IF_CONTENT.Name = `Verticle in class diagram Default in middle between Default-A_CORE_CONTENT and Default-REQ_IF_CONTENT`

	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_A_THE_HEADER_and_Default_REQ_IF_HEADER.X = 1056.999969
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_A_THE_HEADER_and_Default_REQ_IF_HEADER.Y = 97.000000
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_A_THE_HEADER_and_Default_REQ_IF_HEADER.Name = `Verticle in class diagram Default in middle between Default-A_THE_HEADER and Default-REQ_IF_HEADER`

	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_A_TOOL_EXTENSIONS_and_Default_REQ_IF_TOOL_EXTENSION.X = 1178.499969
	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_A_TOOL_EXTENSIONS_and_Default_REQ_IF_TOOL_EXTENSION.Y = 356.000000
	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_A_TOOL_EXTENSIONS_and_Default_REQ_IF_TOOL_EXTENSION.Name = `Verticle in class diagram Default in middle between Default-A_TOOL_EXTENSIONS and Default-REQ_IF_TOOL_EXTENSION`

	__Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_REQ_IF_and_Default_A_CORE_CONTENT.X = 689.000000
	__Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_REQ_IF_and_Default_A_CORE_CONTENT.Y = 171.000000
	__Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_REQ_IF_and_Default_A_CORE_CONTENT.Name = `Verticle in class diagram Default in middle between Default-REQ_IF and Default-A_CORE_CONTENT`

	__Vertice__000004_Verticle_in_class_diagram_Default_in_middle_between_Default_REQ_IF_and_Default_A_THE_HEADER.X = 646.500000
	__Vertice__000004_Verticle_in_class_diagram_Default_in_middle_between_Default_REQ_IF_and_Default_A_THE_HEADER.Y = 107.000000
	__Vertice__000004_Verticle_in_class_diagram_Default_in_middle_between_Default_REQ_IF_and_Default_A_THE_HEADER.Name = `Verticle in class diagram Default in middle between Default-REQ_IF and Default-A_THE_HEADER`

	__Vertice__000005_Verticle_in_class_diagram_Default_in_middle_between_Default_REQ_IF_and_Default_A_TOOL_EXTENSIONS.X = 689.000000
	__Vertice__000005_Verticle_in_class_diagram_Default_in_middle_between_Default_REQ_IF_and_Default_A_TOOL_EXTENSIONS.Y = 234.000000
	__Vertice__000005_Verticle_in_class_diagram_Default_in_middle_between_Default_REQ_IF_and_Default_A_TOOL_EXTENSIONS.Name = `Verticle in class diagram Default in middle between Default-REQ_IF and Default-A_TOOL_EXTENSIONS`

	// Setup of pointers
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000012_Default_REQ_IF)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000013_Default_REQ_IF_CONTENT)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000014_Default_REQ_IF_HEADER)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000010_Default_A_THE_HEADER)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000001_Default_A_CORE_CONTENT)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000000_Default_A_CHILDREN)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000011_Default_A_TOOL_EXTENSIONS)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000015_Default_REQ_IF_TOOL_EXTENSION)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000002_Default_A_DATATYPES)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000003_Default_A_DEFINITION)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000004_Default_A_DEFINITION_1)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000005_Default_A_DEFINITION_2)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000006_Default_A_DEFINITION_3)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000007_Default_A_DEFINITION_4)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000008_Default_A_DEFINITION_5)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000009_Default_A_DEFINITION_6)
	__GongStructShape__000000_Default_A_CHILDREN.Position = __Position__000000_Pos_Default_A_CHILDREN
	__GongStructShape__000001_Default_A_CORE_CONTENT.Position = __Position__000001_Pos_Default_A_CORE_CONTENT
	__GongStructShape__000001_Default_A_CORE_CONTENT.Links = append(__GongStructShape__000001_Default_A_CORE_CONTENT.Links, __Link__000001_REQ_IF_CONTENT)
	__GongStructShape__000002_Default_A_DATATYPES.Position = __Position__000002_Pos_Default_A_DATATYPES
	__GongStructShape__000003_Default_A_DEFINITION.Position = __Position__000003_Pos_Default_A_DEFINITION
	__GongStructShape__000003_Default_A_DEFINITION.Fields = append(__GongStructShape__000003_Default_A_DEFINITION.Fields, __Field__000003_ATTRIBUTE_DEFINITION_INTEGER_REF)
	__GongStructShape__000004_Default_A_DEFINITION_1.Position = __Position__000004_Pos_Default_A_DEFINITION_1
	__GongStructShape__000004_Default_A_DEFINITION_1.Fields = append(__GongStructShape__000004_Default_A_DEFINITION_1.Fields, __Field__000006_ATTRIBUTE_DEFINITION_XHTML_REF)
	__GongStructShape__000005_Default_A_DEFINITION_2.Position = __Position__000005_Pos_Default_A_DEFINITION_2
	__GongStructShape__000005_Default_A_DEFINITION_2.Fields = append(__GongStructShape__000005_Default_A_DEFINITION_2.Fields, __Field__000001_ATTRIBUTE_DEFINITION_DATE_REF)
	__GongStructShape__000006_Default_A_DEFINITION_3.Position = __Position__000006_Pos_Default_A_DEFINITION_3
	__GongStructShape__000006_Default_A_DEFINITION_3.Fields = append(__GongStructShape__000006_Default_A_DEFINITION_3.Fields, __Field__000005_ATTRIBUTE_DEFINITION_STRING_REF)
	__GongStructShape__000007_Default_A_DEFINITION_4.Position = __Position__000007_Pos_Default_A_DEFINITION_4
	__GongStructShape__000007_Default_A_DEFINITION_4.Fields = append(__GongStructShape__000007_Default_A_DEFINITION_4.Fields, __Field__000000_ATTRIBUTE_DEFINITION_BOOLEAN_REF)
	__GongStructShape__000008_Default_A_DEFINITION_5.Position = __Position__000008_Pos_Default_A_DEFINITION_5
	__GongStructShape__000008_Default_A_DEFINITION_5.Fields = append(__GongStructShape__000008_Default_A_DEFINITION_5.Fields, __Field__000002_ATTRIBUTE_DEFINITION_ENUMERATION_REF)
	__GongStructShape__000009_Default_A_DEFINITION_6.Position = __Position__000009_Pos_Default_A_DEFINITION_6
	__GongStructShape__000009_Default_A_DEFINITION_6.Fields = append(__GongStructShape__000009_Default_A_DEFINITION_6.Fields, __Field__000004_ATTRIBUTE_DEFINITION_REAL_REF)
	__GongStructShape__000010_Default_A_THE_HEADER.Position = __Position__000010_Pos_Default_A_THE_HEADER
	__GongStructShape__000010_Default_A_THE_HEADER.Links = append(__GongStructShape__000010_Default_A_THE_HEADER.Links, __Link__000002_REQ_IF_HEADER)
	__GongStructShape__000011_Default_A_TOOL_EXTENSIONS.Position = __Position__000011_Pos_Default_A_TOOL_EXTENSIONS
	__GongStructShape__000011_Default_A_TOOL_EXTENSIONS.Links = append(__GongStructShape__000011_Default_A_TOOL_EXTENSIONS.Links, __Link__000003_REQ_IF_TOOL_EXTENSION)
	__GongStructShape__000012_Default_REQ_IF.Position = __Position__000012_Pos_Default_REQ_IF
	__GongStructShape__000012_Default_REQ_IF.Fields = append(__GongStructShape__000012_Default_REQ_IF.Fields, __Field__000011_Name)
	__GongStructShape__000012_Default_REQ_IF.Links = append(__GongStructShape__000012_Default_REQ_IF.Links, __Link__000004_THE_HEADER)
	__GongStructShape__000012_Default_REQ_IF.Links = append(__GongStructShape__000012_Default_REQ_IF.Links, __Link__000005_TOOL_EXTENSIONS)
	__GongStructShape__000012_Default_REQ_IF.Links = append(__GongStructShape__000012_Default_REQ_IF.Links, __Link__000000_CORE_CONTENT)
	__GongStructShape__000013_Default_REQ_IF_CONTENT.Position = __Position__000013_Pos_Default_REQ_IF_CONTENT
	__GongStructShape__000013_Default_REQ_IF_CONTENT.Fields = append(__GongStructShape__000013_Default_REQ_IF_CONTENT.Fields, __Field__000010_Name)
	__GongStructShape__000014_Default_REQ_IF_HEADER.Position = __Position__000014_Pos_Default_REQ_IF_HEADER
	__GongStructShape__000014_Default_REQ_IF_HEADER.Fields = append(__GongStructShape__000014_Default_REQ_IF_HEADER.Fields, __Field__000009_IDENTIFIER)
	__GongStructShape__000014_Default_REQ_IF_HEADER.Fields = append(__GongStructShape__000014_Default_REQ_IF_HEADER.Fields, __Field__000007_COMMENT)
	__GongStructShape__000014_Default_REQ_IF_HEADER.Fields = append(__GongStructShape__000014_Default_REQ_IF_HEADER.Fields, __Field__000008_CREATION_TIME)
	__GongStructShape__000014_Default_REQ_IF_HEADER.Fields = append(__GongStructShape__000014_Default_REQ_IF_HEADER.Fields, __Field__000012_REPOSITORY_ID)
	__GongStructShape__000014_Default_REQ_IF_HEADER.Fields = append(__GongStructShape__000014_Default_REQ_IF_HEADER.Fields, __Field__000013_REQ_IF_TOOL_ID)
	__GongStructShape__000014_Default_REQ_IF_HEADER.Fields = append(__GongStructShape__000014_Default_REQ_IF_HEADER.Fields, __Field__000014_REQ_IF_VERSION)
	__GongStructShape__000014_Default_REQ_IF_HEADER.Fields = append(__GongStructShape__000014_Default_REQ_IF_HEADER.Fields, __Field__000015_SOURCE_TOOL_ID)
	__GongStructShape__000014_Default_REQ_IF_HEADER.Fields = append(__GongStructShape__000014_Default_REQ_IF_HEADER.Fields, __Field__000016_TITLE)
	__GongStructShape__000015_Default_REQ_IF_TOOL_EXTENSION.Position = __Position__000015_Pos_Default_REQ_IF_TOOL_EXTENSION
	__Link__000000_CORE_CONTENT.Middlevertice = __Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_REQ_IF_and_Default_A_CORE_CONTENT
	__Link__000001_REQ_IF_CONTENT.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_A_CORE_CONTENT_and_Default_REQ_IF_CONTENT
	__Link__000002_REQ_IF_HEADER.Middlevertice = __Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_A_THE_HEADER_and_Default_REQ_IF_HEADER
	__Link__000003_REQ_IF_TOOL_EXTENSION.Middlevertice = __Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_A_TOOL_EXTENSIONS_and_Default_REQ_IF_TOOL_EXTENSION
	__Link__000004_THE_HEADER.Middlevertice = __Vertice__000004_Verticle_in_class_diagram_Default_in_middle_between_Default_REQ_IF_and_Default_A_THE_HEADER
	__Link__000005_TOOL_EXTENSIONS.Middlevertice = __Vertice__000005_Verticle_in_class_diagram_Default_in_middle_between_Default_REQ_IF_and_Default_A_TOOL_EXTENSIONS
}
