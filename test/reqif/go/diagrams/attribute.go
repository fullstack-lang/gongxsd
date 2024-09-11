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

	__Classdiagram__000000_attribute := (&models.Classdiagram{Name: `attribute`}).Stage(stage)

	__Field__000000_DATATYPE_DEFINITION_ENUMERATION_REF := (&models.Field{Name: `DATATYPE_DEFINITION_ENUMERATION_REF`}).Stage(stage)
	__Field__000001_DATATYPE_DEFINITION_XHTML_REF := (&models.Field{Name: `DATATYPE_DEFINITION_XHTML_REF`}).Stage(stage)
	__Field__000002_IDENTIFIER := (&models.Field{Name: `IDENTIFIER`}).Stage(stage)
	__Field__000003_IDENTIFIER := (&models.Field{Name: `IDENTIFIER`}).Stage(stage)
	__Field__000004_IDENTIFIER := (&models.Field{Name: `IDENTIFIER`}).Stage(stage)
	__Field__000005_IDENTIFIER := (&models.Field{Name: `IDENTIFIER`}).Stage(stage)
	__Field__000006_IDENTIFIER := (&models.Field{Name: `IDENTIFIER`}).Stage(stage)
	__Field__000007_IS_EDITABLE := (&models.Field{Name: `IS_EDITABLE`}).Stage(stage)
	__Field__000008_IS_EDITABLE := (&models.Field{Name: `IS_EDITABLE`}).Stage(stage)
	__Field__000009_KEY := (&models.Field{Name: `KEY`}).Stage(stage)
	__Field__000010_LAST_CHANGE := (&models.Field{Name: `LAST_CHANGE`}).Stage(stage)
	__Field__000011_LAST_CHANGE := (&models.Field{Name: `LAST_CHANGE`}).Stage(stage)
	__Field__000012_LAST_CHANGE := (&models.Field{Name: `LAST_CHANGE`}).Stage(stage)
	__Field__000013_LAST_CHANGE := (&models.Field{Name: `LAST_CHANGE`}).Stage(stage)
	__Field__000014_LAST_CHANGE := (&models.Field{Name: `LAST_CHANGE`}).Stage(stage)
	__Field__000015_LONG_NAME := (&models.Field{Name: `LONG_NAME`}).Stage(stage)
	__Field__000016_LONG_NAME := (&models.Field{Name: `LONG_NAME`}).Stage(stage)
	__Field__000017_LONG_NAME := (&models.Field{Name: `LONG_NAME`}).Stage(stage)
	__Field__000018_LONG_NAME := (&models.Field{Name: `LONG_NAME`}).Stage(stage)
	__Field__000019_LONG_NAME := (&models.Field{Name: `LONG_NAME`}).Stage(stage)
	__Field__000020_MULTI_VALUED := (&models.Field{Name: `MULTI_VALUED`}).Stage(stage)
	__Field__000021_OTHER_CONTENT := (&models.Field{Name: `OTHER_CONTENT`}).Stage(stage)

	__GongStructShape__000000_attribute_ATTRIBUTE_DEFINITION_ENUMERATION := (&models.GongStructShape{Name: `attribute-ATTRIBUTE_DEFINITION_ENUMERATION`}).Stage(stage)
	__GongStructShape__000001_attribute_ATTRIBUTE_DEFINITION_XHTML := (&models.GongStructShape{Name: `attribute-ATTRIBUTE_DEFINITION_XHTML`}).Stage(stage)
	__GongStructShape__000002_attribute_A_DATATYPES := (&models.GongStructShape{Name: `attribute-A_DATATYPES`}).Stage(stage)
	__GongStructShape__000003_attribute_A_DATATYPE_DEFINITION_ENUMERATION_REF := (&models.GongStructShape{Name: `attribute-A_DATATYPE_DEFINITION_ENUMERATION_REF`}).Stage(stage)
	__GongStructShape__000004_attribute_A_DATATYPE_DEFINITION_XHTML_REF := (&models.GongStructShape{Name: `attribute-A_DATATYPE_DEFINITION_XHTML_REF`}).Stage(stage)
	__GongStructShape__000005_attribute_A_PROPERTIES := (&models.GongStructShape{Name: `attribute-A_PROPERTIES`}).Stage(stage)
	__GongStructShape__000006_attribute_A_SPECIFIED_VALUES := (&models.GongStructShape{Name: `attribute-A_SPECIFIED_VALUES`}).Stage(stage)
	__GongStructShape__000007_attribute_A_SPEC_ATTRIBUTES := (&models.GongStructShape{Name: `attribute-A_SPEC_ATTRIBUTES`}).Stage(stage)
	__GongStructShape__000008_attribute_A_SPEC_TYPES := (&models.GongStructShape{Name: `attribute-A_SPEC_TYPES`}).Stage(stage)
	__GongStructShape__000009_attribute_DATATYPE_DEFINITION_ENUMERATION := (&models.GongStructShape{Name: `attribute-DATATYPE_DEFINITION_ENUMERATION`}).Stage(stage)
	__GongStructShape__000010_attribute_EMBEDDED_VALUE := (&models.GongStructShape{Name: `attribute-EMBEDDED_VALUE`}).Stage(stage)
	__GongStructShape__000011_attribute_ENUM_VALUE := (&models.GongStructShape{Name: `attribute-ENUM_VALUE`}).Stage(stage)
	__GongStructShape__000012_attribute_REQ_IF_CONTENT := (&models.GongStructShape{Name: `attribute-REQ_IF_CONTENT`}).Stage(stage)
	__GongStructShape__000013_attribute_SPEC_OBJECT_TYPE := (&models.GongStructShape{Name: `attribute-SPEC_OBJECT_TYPE`}).Stage(stage)

	__Link__000000_ATTRIBUTE_DEFINITION_ENUMERATION := (&models.Link{Name: `ATTRIBUTE_DEFINITION_ENUMERATION`}).Stage(stage)
	__Link__000001_ATTRIBUTE_DEFINITION_XHTML := (&models.Link{Name: `ATTRIBUTE_DEFINITION_XHTML`}).Stage(stage)
	__Link__000002_DATATYPES := (&models.Link{Name: `DATATYPES`}).Stage(stage)
	__Link__000003_DATATYPE_DEFINITION_ENUMERATION := (&models.Link{Name: `DATATYPE_DEFINITION_ENUMERATION`}).Stage(stage)
	__Link__000004_EMBEDDED_VALUE := (&models.Link{Name: `EMBEDDED_VALUE`}).Stage(stage)
	__Link__000005_ENUM_VALUE := (&models.Link{Name: `ENUM_VALUE`}).Stage(stage)
	__Link__000006_PROPERTIES := (&models.Link{Name: `PROPERTIES`}).Stage(stage)
	__Link__000007_SPECIFIED_VALUES := (&models.Link{Name: `SPECIFIED_VALUES`}).Stage(stage)
	__Link__000008_SPEC_ATTRIBUTES := (&models.Link{Name: `SPEC_ATTRIBUTES`}).Stage(stage)
	__Link__000009_SPEC_OBJECT_TYPE := (&models.Link{Name: `SPEC_OBJECT_TYPE`}).Stage(stage)
	__Link__000010_SPEC_TYPES := (&models.Link{Name: `SPEC_TYPES`}).Stage(stage)
	__Link__000011_TYPE := (&models.Link{Name: `TYPE`}).Stage(stage)
	__Link__000012_TYPE := (&models.Link{Name: `TYPE`}).Stage(stage)

	__Position__000000_Pos_attribute_ATTRIBUTE_DEFINITION_ENUMERATION := (&models.Position{Name: `Pos-attribute-ATTRIBUTE_DEFINITION_ENUMERATION`}).Stage(stage)
	__Position__000001_Pos_attribute_ATTRIBUTE_DEFINITION_XHTML := (&models.Position{Name: `Pos-attribute-ATTRIBUTE_DEFINITION_XHTML`}).Stage(stage)
	__Position__000002_Pos_attribute_A_DATATYPES := (&models.Position{Name: `Pos-attribute-A_DATATYPES`}).Stage(stage)
	__Position__000003_Pos_attribute_A_DATATYPE_DEFINITION_ENUMERATION_REF := (&models.Position{Name: `Pos-attribute-A_DATATYPE_DEFINITION_ENUMERATION_REF`}).Stage(stage)
	__Position__000004_Pos_attribute_A_DATATYPE_DEFINITION_XHTML_REF := (&models.Position{Name: `Pos-attribute-A_DATATYPE_DEFINITION_XHTML_REF`}).Stage(stage)
	__Position__000005_Pos_attribute_A_PROPERTIES := (&models.Position{Name: `Pos-attribute-A_PROPERTIES`}).Stage(stage)
	__Position__000006_Pos_attribute_A_SPECIFIED_VALUES := (&models.Position{Name: `Pos-attribute-A_SPECIFIED_VALUES`}).Stage(stage)
	__Position__000007_Pos_attribute_A_SPEC_ATTRIBUTES := (&models.Position{Name: `Pos-attribute-A_SPEC_ATTRIBUTES`}).Stage(stage)
	__Position__000008_Pos_attribute_A_SPEC_TYPES := (&models.Position{Name: `Pos-attribute-A_SPEC_TYPES`}).Stage(stage)
	__Position__000009_Pos_attribute_DATATYPE_DEFINITION_ENUMERATION := (&models.Position{Name: `Pos-attribute-DATATYPE_DEFINITION_ENUMERATION`}).Stage(stage)
	__Position__000010_Pos_attribute_EMBEDDED_VALUE := (&models.Position{Name: `Pos-attribute-EMBEDDED_VALUE`}).Stage(stage)
	__Position__000011_Pos_attribute_ENUM_VALUE := (&models.Position{Name: `Pos-attribute-ENUM_VALUE`}).Stage(stage)
	__Position__000012_Pos_attribute_REQ_IF_CONTENT := (&models.Position{Name: `Pos-attribute-REQ_IF_CONTENT`}).Stage(stage)
	__Position__000013_Pos_attribute_SPEC_OBJECT_TYPE := (&models.Position{Name: `Pos-attribute-SPEC_OBJECT_TYPE`}).Stage(stage)

	__Vertice__000000_Verticle_in_class_diagram_attribute_in_middle_between_attribute_ATTRIBUTE_DEFINITION_ENUMERATION_and_attribute_A_DATATYPE_DEFINITION_ENUMERATION_REF := (&models.Vertice{Name: `Verticle in class diagram attribute in middle between attribute-ATTRIBUTE_DEFINITION_ENUMERATION and attribute-A_DATATYPE_DEFINITION_ENUMERATION_REF`}).Stage(stage)
	__Vertice__000001_Verticle_in_class_diagram_attribute_in_middle_between_attribute_ATTRIBUTE_DEFINITION_XHTML_and_attribute_A_DATATYPE_DEFINITION_XHTML_REF := (&models.Vertice{Name: `Verticle in class diagram attribute in middle between attribute-ATTRIBUTE_DEFINITION_XHTML and attribute-A_DATATYPE_DEFINITION_XHTML_REF`}).Stage(stage)
	__Vertice__000002_Verticle_in_class_diagram_attribute_in_middle_between_attribute_A_DATATYPES_and_attribute_DATATYPE_DEFINITION_ENUMERATION := (&models.Vertice{Name: `Verticle in class diagram attribute in middle between attribute-A_DATATYPES and attribute-DATATYPE_DEFINITION_ENUMERATION`}).Stage(stage)
	__Vertice__000003_Verticle_in_class_diagram_attribute_in_middle_between_attribute_A_PROPERTIES_and_attribute_EMBEDDED_VALUE := (&models.Vertice{Name: `Verticle in class diagram attribute in middle between attribute-A_PROPERTIES and attribute-EMBEDDED_VALUE`}).Stage(stage)
	__Vertice__000004_Verticle_in_class_diagram_attribute_in_middle_between_attribute_A_SPECIFIED_VALUES_and_attribute_ENUM_VALUE := (&models.Vertice{Name: `Verticle in class diagram attribute in middle between attribute-A_SPECIFIED_VALUES and attribute-ENUM_VALUE`}).Stage(stage)
	__Vertice__000005_Verticle_in_class_diagram_attribute_in_middle_between_attribute_A_SPEC_ATTRIBUTES_and_attribute_ATTRIBUTE_DEFINITION_ENUMERATION := (&models.Vertice{Name: `Verticle in class diagram attribute in middle between attribute-A_SPEC_ATTRIBUTES and attribute-ATTRIBUTE_DEFINITION_ENUMERATION`}).Stage(stage)
	__Vertice__000006_Verticle_in_class_diagram_attribute_in_middle_between_attribute_A_SPEC_ATTRIBUTES_and_attribute_ATTRIBUTE_DEFINITION_XHTML := (&models.Vertice{Name: `Verticle in class diagram attribute in middle between attribute-A_SPEC_ATTRIBUTES and attribute-ATTRIBUTE_DEFINITION_XHTML`}).Stage(stage)
	__Vertice__000007_Verticle_in_class_diagram_attribute_in_middle_between_attribute_A_SPEC_TYPES_and_attribute_SPEC_OBJECT_TYPE := (&models.Vertice{Name: `Verticle in class diagram attribute in middle between attribute-A_SPEC_TYPES and attribute-SPEC_OBJECT_TYPE`}).Stage(stage)
	__Vertice__000008_Verticle_in_class_diagram_attribute_in_middle_between_attribute_DATATYPE_DEFINITION_ENUMERATION_and_attribute_A_SPECIFIED_VALUES := (&models.Vertice{Name: `Verticle in class diagram attribute in middle between attribute-DATATYPE_DEFINITION_ENUMERATION and attribute-A_SPECIFIED_VALUES`}).Stage(stage)
	__Vertice__000009_Verticle_in_class_diagram_attribute_in_middle_between_attribute_ENUM_VALUE_and_attribute_A_PROPERTIES := (&models.Vertice{Name: `Verticle in class diagram attribute in middle between attribute-ENUM_VALUE and attribute-A_PROPERTIES`}).Stage(stage)
	__Vertice__000010_Verticle_in_class_diagram_attribute_in_middle_between_attribute_REQ_IF_CONTENT_and_attribute_A_DATATYPES := (&models.Vertice{Name: `Verticle in class diagram attribute in middle between attribute-REQ_IF_CONTENT and attribute-A_DATATYPES`}).Stage(stage)
	__Vertice__000011_Verticle_in_class_diagram_attribute_in_middle_between_attribute_REQ_IF_CONTENT_and_attribute_A_SPEC_TYPES := (&models.Vertice{Name: `Verticle in class diagram attribute in middle between attribute-REQ_IF_CONTENT and attribute-A_SPEC_TYPES`}).Stage(stage)
	__Vertice__000012_Verticle_in_class_diagram_attribute_in_middle_between_attribute_SPEC_OBJECT_TYPE_and_attribute_A_SPEC_ATTRIBUTES := (&models.Vertice{Name: `Verticle in class diagram attribute in middle between attribute-SPEC_OBJECT_TYPE and attribute-A_SPEC_ATTRIBUTES`}).Stage(stage)

	// Setup of values

	__Classdiagram__000000_attribute.Name = `attribute`
	__Classdiagram__000000_attribute.IsInDrawMode = false

	__Field__000000_DATATYPE_DEFINITION_ENUMERATION_REF.Name = `DATATYPE_DEFINITION_ENUMERATION_REF`

	//gong:ident [ref_models.A_DATATYPE_DEFINITION_ENUMERATION_REF.DATATYPE_DEFINITION_ENUMERATION_REF] comment added to overcome the problem with the comment map association
	__Field__000000_DATATYPE_DEFINITION_ENUMERATION_REF.Identifier = `ref_models.A_DATATYPE_DEFINITION_ENUMERATION_REF.DATATYPE_DEFINITION_ENUMERATION_REF`
	__Field__000000_DATATYPE_DEFINITION_ENUMERATION_REF.FieldTypeAsString = ``
	__Field__000000_DATATYPE_DEFINITION_ENUMERATION_REF.Structname = `A_DATATYPE_DEFINITION_ENUMERATION_REF`
	__Field__000000_DATATYPE_DEFINITION_ENUMERATION_REF.Fieldtypename = `string`

	__Field__000001_DATATYPE_DEFINITION_XHTML_REF.Name = `DATATYPE_DEFINITION_XHTML_REF`

	//gong:ident [ref_models.A_DATATYPE_DEFINITION_XHTML_REF.DATATYPE_DEFINITION_XHTML_REF] comment added to overcome the problem with the comment map association
	__Field__000001_DATATYPE_DEFINITION_XHTML_REF.Identifier = `ref_models.A_DATATYPE_DEFINITION_XHTML_REF.DATATYPE_DEFINITION_XHTML_REF`
	__Field__000001_DATATYPE_DEFINITION_XHTML_REF.FieldTypeAsString = ``
	__Field__000001_DATATYPE_DEFINITION_XHTML_REF.Structname = `A_DATATYPE_DEFINITION_XHTML_REF`
	__Field__000001_DATATYPE_DEFINITION_XHTML_REF.Fieldtypename = `string`

	__Field__000002_IDENTIFIER.Name = `IDENTIFIER`

	//gong:ident [ref_models.SPEC_OBJECT_TYPE.IDENTIFIER] comment added to overcome the problem with the comment map association
	__Field__000002_IDENTIFIER.Identifier = `ref_models.SPEC_OBJECT_TYPE.IDENTIFIER`
	__Field__000002_IDENTIFIER.FieldTypeAsString = ``
	__Field__000002_IDENTIFIER.Structname = `SPEC_OBJECT_TYPE`
	__Field__000002_IDENTIFIER.Fieldtypename = `string`

	__Field__000003_IDENTIFIER.Name = `IDENTIFIER`

	//gong:ident [ref_models.ATTRIBUTE_DEFINITION_XHTML.IDENTIFIER] comment added to overcome the problem with the comment map association
	__Field__000003_IDENTIFIER.Identifier = `ref_models.ATTRIBUTE_DEFINITION_XHTML.IDENTIFIER`
	__Field__000003_IDENTIFIER.FieldTypeAsString = ``
	__Field__000003_IDENTIFIER.Structname = `ATTRIBUTE_DEFINITION_XHTML`
	__Field__000003_IDENTIFIER.Fieldtypename = `string`

	__Field__000004_IDENTIFIER.Name = `IDENTIFIER`

	//gong:ident [ref_models.ENUM_VALUE.IDENTIFIER] comment added to overcome the problem with the comment map association
	__Field__000004_IDENTIFIER.Identifier = `ref_models.ENUM_VALUE.IDENTIFIER`
	__Field__000004_IDENTIFIER.FieldTypeAsString = ``
	__Field__000004_IDENTIFIER.Structname = `ENUM_VALUE`
	__Field__000004_IDENTIFIER.Fieldtypename = `string`

	__Field__000005_IDENTIFIER.Name = `IDENTIFIER`

	//gong:ident [ref_models.ATTRIBUTE_DEFINITION_ENUMERATION.IDENTIFIER] comment added to overcome the problem with the comment map association
	__Field__000005_IDENTIFIER.Identifier = `ref_models.ATTRIBUTE_DEFINITION_ENUMERATION.IDENTIFIER`
	__Field__000005_IDENTIFIER.FieldTypeAsString = ``
	__Field__000005_IDENTIFIER.Structname = `ATTRIBUTE_DEFINITION_ENUMERATION`
	__Field__000005_IDENTIFIER.Fieldtypename = `string`

	__Field__000006_IDENTIFIER.Name = `IDENTIFIER`

	//gong:ident [ref_models.DATATYPE_DEFINITION_ENUMERATION.IDENTIFIER] comment added to overcome the problem with the comment map association
	__Field__000006_IDENTIFIER.Identifier = `ref_models.DATATYPE_DEFINITION_ENUMERATION.IDENTIFIER`
	__Field__000006_IDENTIFIER.FieldTypeAsString = ``
	__Field__000006_IDENTIFIER.Structname = `DATATYPE_DEFINITION_ENUMERATION`
	__Field__000006_IDENTIFIER.Fieldtypename = `string`

	__Field__000007_IS_EDITABLE.Name = `IS_EDITABLE`

	//gong:ident [ref_models.ATTRIBUTE_DEFINITION_ENUMERATION.IS_EDITABLE] comment added to overcome the problem with the comment map association
	__Field__000007_IS_EDITABLE.Identifier = `ref_models.ATTRIBUTE_DEFINITION_ENUMERATION.IS_EDITABLE`
	__Field__000007_IS_EDITABLE.FieldTypeAsString = ``
	__Field__000007_IS_EDITABLE.Structname = `ATTRIBUTE_DEFINITION_ENUMERATION`
	__Field__000007_IS_EDITABLE.Fieldtypename = `bool`

	__Field__000008_IS_EDITABLE.Name = `IS_EDITABLE`

	//gong:ident [ref_models.ATTRIBUTE_DEFINITION_XHTML.IS_EDITABLE] comment added to overcome the problem with the comment map association
	__Field__000008_IS_EDITABLE.Identifier = `ref_models.ATTRIBUTE_DEFINITION_XHTML.IS_EDITABLE`
	__Field__000008_IS_EDITABLE.FieldTypeAsString = ``
	__Field__000008_IS_EDITABLE.Structname = `ATTRIBUTE_DEFINITION_XHTML`
	__Field__000008_IS_EDITABLE.Fieldtypename = `bool`

	__Field__000009_KEY.Name = `KEY`

	//gong:ident [ref_models.EMBEDDED_VALUE.KEY] comment added to overcome the problem with the comment map association
	__Field__000009_KEY.Identifier = `ref_models.EMBEDDED_VALUE.KEY`
	__Field__000009_KEY.FieldTypeAsString = ``
	__Field__000009_KEY.Structname = `EMBEDDED_VALUE`
	__Field__000009_KEY.Fieldtypename = `int`

	__Field__000010_LAST_CHANGE.Name = `LAST_CHANGE`

	//gong:ident [ref_models.ATTRIBUTE_DEFINITION_ENUMERATION.LAST_CHANGE] comment added to overcome the problem with the comment map association
	__Field__000010_LAST_CHANGE.Identifier = `ref_models.ATTRIBUTE_DEFINITION_ENUMERATION.LAST_CHANGE`
	__Field__000010_LAST_CHANGE.FieldTypeAsString = ``
	__Field__000010_LAST_CHANGE.Structname = `ATTRIBUTE_DEFINITION_ENUMERATION`
	__Field__000010_LAST_CHANGE.Fieldtypename = `string`

	__Field__000011_LAST_CHANGE.Name = `LAST_CHANGE`

	//gong:ident [ref_models.ATTRIBUTE_DEFINITION_XHTML.LAST_CHANGE] comment added to overcome the problem with the comment map association
	__Field__000011_LAST_CHANGE.Identifier = `ref_models.ATTRIBUTE_DEFINITION_XHTML.LAST_CHANGE`
	__Field__000011_LAST_CHANGE.FieldTypeAsString = ``
	__Field__000011_LAST_CHANGE.Structname = `ATTRIBUTE_DEFINITION_XHTML`
	__Field__000011_LAST_CHANGE.Fieldtypename = `string`

	__Field__000012_LAST_CHANGE.Name = `LAST_CHANGE`

	//gong:ident [ref_models.SPEC_OBJECT_TYPE.LAST_CHANGE] comment added to overcome the problem with the comment map association
	__Field__000012_LAST_CHANGE.Identifier = `ref_models.SPEC_OBJECT_TYPE.LAST_CHANGE`
	__Field__000012_LAST_CHANGE.FieldTypeAsString = ``
	__Field__000012_LAST_CHANGE.Structname = `SPEC_OBJECT_TYPE`
	__Field__000012_LAST_CHANGE.Fieldtypename = `string`

	__Field__000013_LAST_CHANGE.Name = `LAST_CHANGE`

	//gong:ident [ref_models.ENUM_VALUE.LAST_CHANGE] comment added to overcome the problem with the comment map association
	__Field__000013_LAST_CHANGE.Identifier = `ref_models.ENUM_VALUE.LAST_CHANGE`
	__Field__000013_LAST_CHANGE.FieldTypeAsString = ``
	__Field__000013_LAST_CHANGE.Structname = `ENUM_VALUE`
	__Field__000013_LAST_CHANGE.Fieldtypename = `string`

	__Field__000014_LAST_CHANGE.Name = `LAST_CHANGE`

	//gong:ident [ref_models.DATATYPE_DEFINITION_ENUMERATION.LAST_CHANGE] comment added to overcome the problem with the comment map association
	__Field__000014_LAST_CHANGE.Identifier = `ref_models.DATATYPE_DEFINITION_ENUMERATION.LAST_CHANGE`
	__Field__000014_LAST_CHANGE.FieldTypeAsString = ``
	__Field__000014_LAST_CHANGE.Structname = `DATATYPE_DEFINITION_ENUMERATION`
	__Field__000014_LAST_CHANGE.Fieldtypename = `string`

	__Field__000015_LONG_NAME.Name = `LONG_NAME`

	//gong:ident [ref_models.ENUM_VALUE.LONG_NAME] comment added to overcome the problem with the comment map association
	__Field__000015_LONG_NAME.Identifier = `ref_models.ENUM_VALUE.LONG_NAME`
	__Field__000015_LONG_NAME.FieldTypeAsString = ``
	__Field__000015_LONG_NAME.Structname = `ENUM_VALUE`
	__Field__000015_LONG_NAME.Fieldtypename = `string`

	__Field__000016_LONG_NAME.Name = `LONG_NAME`

	//gong:ident [ref_models.ATTRIBUTE_DEFINITION_ENUMERATION.LONG_NAME] comment added to overcome the problem with the comment map association
	__Field__000016_LONG_NAME.Identifier = `ref_models.ATTRIBUTE_DEFINITION_ENUMERATION.LONG_NAME`
	__Field__000016_LONG_NAME.FieldTypeAsString = ``
	__Field__000016_LONG_NAME.Structname = `ATTRIBUTE_DEFINITION_ENUMERATION`
	__Field__000016_LONG_NAME.Fieldtypename = `string`

	__Field__000017_LONG_NAME.Name = `LONG_NAME`

	//gong:ident [ref_models.ATTRIBUTE_DEFINITION_XHTML.LONG_NAME] comment added to overcome the problem with the comment map association
	__Field__000017_LONG_NAME.Identifier = `ref_models.ATTRIBUTE_DEFINITION_XHTML.LONG_NAME`
	__Field__000017_LONG_NAME.FieldTypeAsString = ``
	__Field__000017_LONG_NAME.Structname = `ATTRIBUTE_DEFINITION_XHTML`
	__Field__000017_LONG_NAME.Fieldtypename = `string`

	__Field__000018_LONG_NAME.Name = `LONG_NAME`

	//gong:ident [ref_models.DATATYPE_DEFINITION_ENUMERATION.LONG_NAME] comment added to overcome the problem with the comment map association
	__Field__000018_LONG_NAME.Identifier = `ref_models.DATATYPE_DEFINITION_ENUMERATION.LONG_NAME`
	__Field__000018_LONG_NAME.FieldTypeAsString = ``
	__Field__000018_LONG_NAME.Structname = `DATATYPE_DEFINITION_ENUMERATION`
	__Field__000018_LONG_NAME.Fieldtypename = `string`

	__Field__000019_LONG_NAME.Name = `LONG_NAME`

	//gong:ident [ref_models.SPEC_OBJECT_TYPE.LONG_NAME] comment added to overcome the problem with the comment map association
	__Field__000019_LONG_NAME.Identifier = `ref_models.SPEC_OBJECT_TYPE.LONG_NAME`
	__Field__000019_LONG_NAME.FieldTypeAsString = ``
	__Field__000019_LONG_NAME.Structname = `SPEC_OBJECT_TYPE`
	__Field__000019_LONG_NAME.Fieldtypename = `string`

	__Field__000020_MULTI_VALUED.Name = `MULTI_VALUED`

	//gong:ident [ref_models.ATTRIBUTE_DEFINITION_ENUMERATION.MULTI_VALUED] comment added to overcome the problem with the comment map association
	__Field__000020_MULTI_VALUED.Identifier = `ref_models.ATTRIBUTE_DEFINITION_ENUMERATION.MULTI_VALUED`
	__Field__000020_MULTI_VALUED.FieldTypeAsString = ``
	__Field__000020_MULTI_VALUED.Structname = `ATTRIBUTE_DEFINITION_ENUMERATION`
	__Field__000020_MULTI_VALUED.Fieldtypename = `bool`

	__Field__000021_OTHER_CONTENT.Name = `OTHER_CONTENT`

	//gong:ident [ref_models.EMBEDDED_VALUE.OTHER_CONTENT] comment added to overcome the problem with the comment map association
	__Field__000021_OTHER_CONTENT.Identifier = `ref_models.EMBEDDED_VALUE.OTHER_CONTENT`
	__Field__000021_OTHER_CONTENT.FieldTypeAsString = ``
	__Field__000021_OTHER_CONTENT.Structname = `EMBEDDED_VALUE`
	__Field__000021_OTHER_CONTENT.Fieldtypename = `string`

	__GongStructShape__000000_attribute_ATTRIBUTE_DEFINITION_ENUMERATION.Name = `attribute-ATTRIBUTE_DEFINITION_ENUMERATION`

	//gong:ident [ref_models.ATTRIBUTE_DEFINITION_ENUMERATION] comment added to overcome the problem with the comment map association
	__GongStructShape__000000_attribute_ATTRIBUTE_DEFINITION_ENUMERATION.Identifier = `ref_models.ATTRIBUTE_DEFINITION_ENUMERATION`
	__GongStructShape__000000_attribute_ATTRIBUTE_DEFINITION_ENUMERATION.ShowNbInstances = false
	__GongStructShape__000000_attribute_ATTRIBUTE_DEFINITION_ENUMERATION.NbInstances = 0
	__GongStructShape__000000_attribute_ATTRIBUTE_DEFINITION_ENUMERATION.Width = 330.999939
	__GongStructShape__000000_attribute_ATTRIBUTE_DEFINITION_ENUMERATION.Height = 138.000000
	__GongStructShape__000000_attribute_ATTRIBUTE_DEFINITION_ENUMERATION.IsSelected = false

	__GongStructShape__000001_attribute_ATTRIBUTE_DEFINITION_XHTML.Name = `attribute-ATTRIBUTE_DEFINITION_XHTML`

	//gong:ident [ref_models.ATTRIBUTE_DEFINITION_XHTML] comment added to overcome the problem with the comment map association
	__GongStructShape__000001_attribute_ATTRIBUTE_DEFINITION_XHTML.Identifier = `ref_models.ATTRIBUTE_DEFINITION_XHTML`
	__GongStructShape__000001_attribute_ATTRIBUTE_DEFINITION_XHTML.ShowNbInstances = false
	__GongStructShape__000001_attribute_ATTRIBUTE_DEFINITION_XHTML.NbInstances = 0
	__GongStructShape__000001_attribute_ATTRIBUTE_DEFINITION_XHTML.Width = 299.999939
	__GongStructShape__000001_attribute_ATTRIBUTE_DEFINITION_XHTML.Height = 123.000000
	__GongStructShape__000001_attribute_ATTRIBUTE_DEFINITION_XHTML.IsSelected = false

	__GongStructShape__000002_attribute_A_DATATYPES.Name = `attribute-A_DATATYPES`

	//gong:ident [ref_models.A_DATATYPES] comment added to overcome the problem with the comment map association
	__GongStructShape__000002_attribute_A_DATATYPES.Identifier = `ref_models.A_DATATYPES`
	__GongStructShape__000002_attribute_A_DATATYPES.ShowNbInstances = false
	__GongStructShape__000002_attribute_A_DATATYPES.NbInstances = 0
	__GongStructShape__000002_attribute_A_DATATYPES.Width = 240.000000
	__GongStructShape__000002_attribute_A_DATATYPES.Height = 63.000000
	__GongStructShape__000002_attribute_A_DATATYPES.IsSelected = false

	__GongStructShape__000003_attribute_A_DATATYPE_DEFINITION_ENUMERATION_REF.Name = `attribute-A_DATATYPE_DEFINITION_ENUMERATION_REF`

	//gong:ident [ref_models.A_DATATYPE_DEFINITION_ENUMERATION_REF] comment added to overcome the problem with the comment map association
	__GongStructShape__000003_attribute_A_DATATYPE_DEFINITION_ENUMERATION_REF.Identifier = `ref_models.A_DATATYPE_DEFINITION_ENUMERATION_REF`
	__GongStructShape__000003_attribute_A_DATATYPE_DEFINITION_ENUMERATION_REF.ShowNbInstances = false
	__GongStructShape__000003_attribute_A_DATATYPE_DEFINITION_ENUMERATION_REF.NbInstances = 0
	__GongStructShape__000003_attribute_A_DATATYPE_DEFINITION_ENUMERATION_REF.Width = 413.000000
	__GongStructShape__000003_attribute_A_DATATYPE_DEFINITION_ENUMERATION_REF.Height = 78.000000
	__GongStructShape__000003_attribute_A_DATATYPE_DEFINITION_ENUMERATION_REF.IsSelected = false

	__GongStructShape__000004_attribute_A_DATATYPE_DEFINITION_XHTML_REF.Name = `attribute-A_DATATYPE_DEFINITION_XHTML_REF`

	//gong:ident [ref_models.A_DATATYPE_DEFINITION_XHTML_REF] comment added to overcome the problem with the comment map association
	__GongStructShape__000004_attribute_A_DATATYPE_DEFINITION_XHTML_REF.Identifier = `ref_models.A_DATATYPE_DEFINITION_XHTML_REF`
	__GongStructShape__000004_attribute_A_DATATYPE_DEFINITION_XHTML_REF.ShowNbInstances = false
	__GongStructShape__000004_attribute_A_DATATYPE_DEFINITION_XHTML_REF.NbInstances = 0
	__GongStructShape__000004_attribute_A_DATATYPE_DEFINITION_XHTML_REF.Width = 406.000000
	__GongStructShape__000004_attribute_A_DATATYPE_DEFINITION_XHTML_REF.Height = 78.000000
	__GongStructShape__000004_attribute_A_DATATYPE_DEFINITION_XHTML_REF.IsSelected = false

	__GongStructShape__000005_attribute_A_PROPERTIES.Name = `attribute-A_PROPERTIES`

	//gong:ident [ref_models.A_PROPERTIES] comment added to overcome the problem with the comment map association
	__GongStructShape__000005_attribute_A_PROPERTIES.Identifier = `ref_models.A_PROPERTIES`
	__GongStructShape__000005_attribute_A_PROPERTIES.ShowNbInstances = false
	__GongStructShape__000005_attribute_A_PROPERTIES.NbInstances = 0
	__GongStructShape__000005_attribute_A_PROPERTIES.Width = 240.000000
	__GongStructShape__000005_attribute_A_PROPERTIES.Height = 63.000000
	__GongStructShape__000005_attribute_A_PROPERTIES.IsSelected = false

	__GongStructShape__000006_attribute_A_SPECIFIED_VALUES.Name = `attribute-A_SPECIFIED_VALUES`

	//gong:ident [ref_models.A_SPECIFIED_VALUES] comment added to overcome the problem with the comment map association
	__GongStructShape__000006_attribute_A_SPECIFIED_VALUES.Identifier = `ref_models.A_SPECIFIED_VALUES`
	__GongStructShape__000006_attribute_A_SPECIFIED_VALUES.ShowNbInstances = false
	__GongStructShape__000006_attribute_A_SPECIFIED_VALUES.NbInstances = 0
	__GongStructShape__000006_attribute_A_SPECIFIED_VALUES.Width = 240.000000
	__GongStructShape__000006_attribute_A_SPECIFIED_VALUES.Height = 63.000000
	__GongStructShape__000006_attribute_A_SPECIFIED_VALUES.IsSelected = false

	__GongStructShape__000007_attribute_A_SPEC_ATTRIBUTES.Name = `attribute-A_SPEC_ATTRIBUTES`

	//gong:ident [ref_models.A_SPEC_ATTRIBUTES] comment added to overcome the problem with the comment map association
	__GongStructShape__000007_attribute_A_SPEC_ATTRIBUTES.Identifier = `ref_models.A_SPEC_ATTRIBUTES`
	__GongStructShape__000007_attribute_A_SPEC_ATTRIBUTES.ShowNbInstances = false
	__GongStructShape__000007_attribute_A_SPEC_ATTRIBUTES.NbInstances = 0
	__GongStructShape__000007_attribute_A_SPEC_ATTRIBUTES.Width = 240.000000
	__GongStructShape__000007_attribute_A_SPEC_ATTRIBUTES.Height = 307.000000
	__GongStructShape__000007_attribute_A_SPEC_ATTRIBUTES.IsSelected = false

	__GongStructShape__000008_attribute_A_SPEC_TYPES.Name = `attribute-A_SPEC_TYPES`

	//gong:ident [ref_models.A_SPEC_TYPES] comment added to overcome the problem with the comment map association
	__GongStructShape__000008_attribute_A_SPEC_TYPES.Identifier = `ref_models.A_SPEC_TYPES`
	__GongStructShape__000008_attribute_A_SPEC_TYPES.ShowNbInstances = false
	__GongStructShape__000008_attribute_A_SPEC_TYPES.NbInstances = 0
	__GongStructShape__000008_attribute_A_SPEC_TYPES.Width = 240.000000
	__GongStructShape__000008_attribute_A_SPEC_TYPES.Height = 63.000000
	__GongStructShape__000008_attribute_A_SPEC_TYPES.IsSelected = false

	__GongStructShape__000009_attribute_DATATYPE_DEFINITION_ENUMERATION.Name = `attribute-DATATYPE_DEFINITION_ENUMERATION`

	//gong:ident [ref_models.DATATYPE_DEFINITION_ENUMERATION] comment added to overcome the problem with the comment map association
	__GongStructShape__000009_attribute_DATATYPE_DEFINITION_ENUMERATION.Identifier = `ref_models.DATATYPE_DEFINITION_ENUMERATION`
	__GongStructShape__000009_attribute_DATATYPE_DEFINITION_ENUMERATION.ShowNbInstances = false
	__GongStructShape__000009_attribute_DATATYPE_DEFINITION_ENUMERATION.NbInstances = 0
	__GongStructShape__000009_attribute_DATATYPE_DEFINITION_ENUMERATION.Width = 337.000000
	__GongStructShape__000009_attribute_DATATYPE_DEFINITION_ENUMERATION.Height = 108.000000
	__GongStructShape__000009_attribute_DATATYPE_DEFINITION_ENUMERATION.IsSelected = false

	__GongStructShape__000010_attribute_EMBEDDED_VALUE.Name = `attribute-EMBEDDED_VALUE`

	//gong:ident [ref_models.EMBEDDED_VALUE] comment added to overcome the problem with the comment map association
	__GongStructShape__000010_attribute_EMBEDDED_VALUE.Identifier = `ref_models.EMBEDDED_VALUE`
	__GongStructShape__000010_attribute_EMBEDDED_VALUE.ShowNbInstances = false
	__GongStructShape__000010_attribute_EMBEDDED_VALUE.NbInstances = 0
	__GongStructShape__000010_attribute_EMBEDDED_VALUE.Width = 240.000000
	__GongStructShape__000010_attribute_EMBEDDED_VALUE.Height = 93.000000
	__GongStructShape__000010_attribute_EMBEDDED_VALUE.IsSelected = false

	__GongStructShape__000011_attribute_ENUM_VALUE.Name = `attribute-ENUM_VALUE`

	//gong:ident [ref_models.ENUM_VALUE] comment added to overcome the problem with the comment map association
	__GongStructShape__000011_attribute_ENUM_VALUE.Identifier = `ref_models.ENUM_VALUE`
	__GongStructShape__000011_attribute_ENUM_VALUE.ShowNbInstances = false
	__GongStructShape__000011_attribute_ENUM_VALUE.NbInstances = 0
	__GongStructShape__000011_attribute_ENUM_VALUE.Width = 240.000000
	__GongStructShape__000011_attribute_ENUM_VALUE.Height = 108.000000
	__GongStructShape__000011_attribute_ENUM_VALUE.IsSelected = false

	__GongStructShape__000012_attribute_REQ_IF_CONTENT.Name = `attribute-REQ_IF_CONTENT`

	//gong:ident [ref_models.REQ_IF_CONTENT] comment added to overcome the problem with the comment map association
	__GongStructShape__000012_attribute_REQ_IF_CONTENT.Identifier = `ref_models.REQ_IF_CONTENT`
	__GongStructShape__000012_attribute_REQ_IF_CONTENT.ShowNbInstances = false
	__GongStructShape__000012_attribute_REQ_IF_CONTENT.NbInstances = 0
	__GongStructShape__000012_attribute_REQ_IF_CONTENT.Width = 240.000000
	__GongStructShape__000012_attribute_REQ_IF_CONTENT.Height = 63.000000
	__GongStructShape__000012_attribute_REQ_IF_CONTENT.IsSelected = false

	__GongStructShape__000013_attribute_SPEC_OBJECT_TYPE.Name = `attribute-SPEC_OBJECT_TYPE`

	//gong:ident [ref_models.SPEC_OBJECT_TYPE] comment added to overcome the problem with the comment map association
	__GongStructShape__000013_attribute_SPEC_OBJECT_TYPE.Identifier = `ref_models.SPEC_OBJECT_TYPE`
	__GongStructShape__000013_attribute_SPEC_OBJECT_TYPE.ShowNbInstances = false
	__GongStructShape__000013_attribute_SPEC_OBJECT_TYPE.NbInstances = 0
	__GongStructShape__000013_attribute_SPEC_OBJECT_TYPE.Width = 240.000000
	__GongStructShape__000013_attribute_SPEC_OBJECT_TYPE.Height = 108.000000
	__GongStructShape__000013_attribute_SPEC_OBJECT_TYPE.IsSelected = false

	__Link__000000_ATTRIBUTE_DEFINITION_ENUMERATION.Name = `ATTRIBUTE_DEFINITION_ENUMERATION`

	//gong:ident [ref_models.A_SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION] comment added to overcome the problem with the comment map association
	__Link__000000_ATTRIBUTE_DEFINITION_ENUMERATION.Identifier = `ref_models.A_SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION`

	//gong:ident [ref_models.ATTRIBUTE_DEFINITION_ENUMERATION] comment added to overcome the problem with the comment map association
	__Link__000000_ATTRIBUTE_DEFINITION_ENUMERATION.Fieldtypename = `ref_models.ATTRIBUTE_DEFINITION_ENUMERATION`
	__Link__000000_ATTRIBUTE_DEFINITION_ENUMERATION.FieldOffsetX = -50.000000
	__Link__000000_ATTRIBUTE_DEFINITION_ENUMERATION.FieldOffsetY = -16.000000
	__Link__000000_ATTRIBUTE_DEFINITION_ENUMERATION.TargetMultiplicity = models.MANY
	__Link__000000_ATTRIBUTE_DEFINITION_ENUMERATION.TargetMultiplicityOffsetX = -50.000000
	__Link__000000_ATTRIBUTE_DEFINITION_ENUMERATION.TargetMultiplicityOffsetY = 16.000000
	__Link__000000_ATTRIBUTE_DEFINITION_ENUMERATION.SourceMultiplicity = models.MANY
	__Link__000000_ATTRIBUTE_DEFINITION_ENUMERATION.SourceMultiplicityOffsetX = 10.000000
	__Link__000000_ATTRIBUTE_DEFINITION_ENUMERATION.SourceMultiplicityOffsetY = -50.000000
	__Link__000000_ATTRIBUTE_DEFINITION_ENUMERATION.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_ATTRIBUTE_DEFINITION_ENUMERATION.StartRatio = 0.785559
	__Link__000000_ATTRIBUTE_DEFINITION_ENUMERATION.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_ATTRIBUTE_DEFINITION_ENUMERATION.EndRatio = 0.500000
	__Link__000000_ATTRIBUTE_DEFINITION_ENUMERATION.CornerOffsetRatio = 1.380000

	__Link__000001_ATTRIBUTE_DEFINITION_XHTML.Name = `ATTRIBUTE_DEFINITION_XHTML`

	//gong:ident [ref_models.A_SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML] comment added to overcome the problem with the comment map association
	__Link__000001_ATTRIBUTE_DEFINITION_XHTML.Identifier = `ref_models.A_SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML`

	//gong:ident [ref_models.ATTRIBUTE_DEFINITION_XHTML] comment added to overcome the problem with the comment map association
	__Link__000001_ATTRIBUTE_DEFINITION_XHTML.Fieldtypename = `ref_models.ATTRIBUTE_DEFINITION_XHTML`
	__Link__000001_ATTRIBUTE_DEFINITION_XHTML.FieldOffsetX = -50.000000
	__Link__000001_ATTRIBUTE_DEFINITION_XHTML.FieldOffsetY = -16.000000
	__Link__000001_ATTRIBUTE_DEFINITION_XHTML.TargetMultiplicity = models.MANY
	__Link__000001_ATTRIBUTE_DEFINITION_XHTML.TargetMultiplicityOffsetX = -50.000000
	__Link__000001_ATTRIBUTE_DEFINITION_XHTML.TargetMultiplicityOffsetY = 16.000000
	__Link__000001_ATTRIBUTE_DEFINITION_XHTML.SourceMultiplicity = models.MANY
	__Link__000001_ATTRIBUTE_DEFINITION_XHTML.SourceMultiplicityOffsetX = 10.000000
	__Link__000001_ATTRIBUTE_DEFINITION_XHTML.SourceMultiplicityOffsetY = -50.000000
	__Link__000001_ATTRIBUTE_DEFINITION_XHTML.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000001_ATTRIBUTE_DEFINITION_XHTML.StartRatio = 0.156895
	__Link__000001_ATTRIBUTE_DEFINITION_XHTML.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000001_ATTRIBUTE_DEFINITION_XHTML.EndRatio = 0.500000
	__Link__000001_ATTRIBUTE_DEFINITION_XHTML.CornerOffsetRatio = 1.380000

	__Link__000002_DATATYPES.Name = `DATATYPES`

	//gong:ident [ref_models.REQ_IF_CONTENT.DATATYPES] comment added to overcome the problem with the comment map association
	__Link__000002_DATATYPES.Identifier = `ref_models.REQ_IF_CONTENT.DATATYPES`

	//gong:ident [ref_models.A_DATATYPES] comment added to overcome the problem with the comment map association
	__Link__000002_DATATYPES.Fieldtypename = `ref_models.A_DATATYPES`
	__Link__000002_DATATYPES.FieldOffsetX = -50.000000
	__Link__000002_DATATYPES.FieldOffsetY = -16.000000
	__Link__000002_DATATYPES.TargetMultiplicity = models.ZERO_ONE
	__Link__000002_DATATYPES.TargetMultiplicityOffsetX = -50.000000
	__Link__000002_DATATYPES.TargetMultiplicityOffsetY = 16.000000
	__Link__000002_DATATYPES.SourceMultiplicity = models.MANY
	__Link__000002_DATATYPES.SourceMultiplicityOffsetX = 10.000000
	__Link__000002_DATATYPES.SourceMultiplicityOffsetY = -50.000000
	__Link__000002_DATATYPES.StartOrientation = models.ORIENTATION_VERTICAL
	__Link__000002_DATATYPES.StartRatio = 0.176389
	__Link__000002_DATATYPES.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000002_DATATYPES.EndRatio = 0.500000
	__Link__000002_DATATYPES.CornerOffsetRatio = 1.018518

	__Link__000003_DATATYPE_DEFINITION_ENUMERATION.Name = `DATATYPE_DEFINITION_ENUMERATION`

	//gong:ident [ref_models.A_DATATYPES.DATATYPE_DEFINITION_ENUMERATION] comment added to overcome the problem with the comment map association
	__Link__000003_DATATYPE_DEFINITION_ENUMERATION.Identifier = `ref_models.A_DATATYPES.DATATYPE_DEFINITION_ENUMERATION`

	//gong:ident [ref_models.DATATYPE_DEFINITION_ENUMERATION] comment added to overcome the problem with the comment map association
	__Link__000003_DATATYPE_DEFINITION_ENUMERATION.Fieldtypename = `ref_models.DATATYPE_DEFINITION_ENUMERATION`
	__Link__000003_DATATYPE_DEFINITION_ENUMERATION.FieldOffsetX = -50.000000
	__Link__000003_DATATYPE_DEFINITION_ENUMERATION.FieldOffsetY = -16.000000
	__Link__000003_DATATYPE_DEFINITION_ENUMERATION.TargetMultiplicity = models.MANY
	__Link__000003_DATATYPE_DEFINITION_ENUMERATION.TargetMultiplicityOffsetX = -50.000000
	__Link__000003_DATATYPE_DEFINITION_ENUMERATION.TargetMultiplicityOffsetY = 16.000000
	__Link__000003_DATATYPE_DEFINITION_ENUMERATION.SourceMultiplicity = models.MANY
	__Link__000003_DATATYPE_DEFINITION_ENUMERATION.SourceMultiplicityOffsetX = 10.000000
	__Link__000003_DATATYPE_DEFINITION_ENUMERATION.SourceMultiplicityOffsetY = -50.000000
	__Link__000003_DATATYPE_DEFINITION_ENUMERATION.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000003_DATATYPE_DEFINITION_ENUMERATION.StartRatio = 0.500000
	__Link__000003_DATATYPE_DEFINITION_ENUMERATION.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000003_DATATYPE_DEFINITION_ENUMERATION.EndRatio = 0.500000
	__Link__000003_DATATYPE_DEFINITION_ENUMERATION.CornerOffsetRatio = 1.380000

	__Link__000004_EMBEDDED_VALUE.Name = `EMBEDDED_VALUE`

	//gong:ident [ref_models.A_PROPERTIES.EMBEDDED_VALUE] comment added to overcome the problem with the comment map association
	__Link__000004_EMBEDDED_VALUE.Identifier = `ref_models.A_PROPERTIES.EMBEDDED_VALUE`

	//gong:ident [ref_models.EMBEDDED_VALUE] comment added to overcome the problem with the comment map association
	__Link__000004_EMBEDDED_VALUE.Fieldtypename = `ref_models.EMBEDDED_VALUE`
	__Link__000004_EMBEDDED_VALUE.FieldOffsetX = -50.000000
	__Link__000004_EMBEDDED_VALUE.FieldOffsetY = -16.000000
	__Link__000004_EMBEDDED_VALUE.TargetMultiplicity = models.ZERO_ONE
	__Link__000004_EMBEDDED_VALUE.TargetMultiplicityOffsetX = -50.000000
	__Link__000004_EMBEDDED_VALUE.TargetMultiplicityOffsetY = 16.000000
	__Link__000004_EMBEDDED_VALUE.SourceMultiplicity = models.MANY
	__Link__000004_EMBEDDED_VALUE.SourceMultiplicityOffsetX = 10.000000
	__Link__000004_EMBEDDED_VALUE.SourceMultiplicityOffsetY = -50.000000
	__Link__000004_EMBEDDED_VALUE.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000004_EMBEDDED_VALUE.StartRatio = 0.500000
	__Link__000004_EMBEDDED_VALUE.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000004_EMBEDDED_VALUE.EndRatio = 0.500000
	__Link__000004_EMBEDDED_VALUE.CornerOffsetRatio = 1.380000

	__Link__000005_ENUM_VALUE.Name = `ENUM_VALUE`

	//gong:ident [ref_models.A_SPECIFIED_VALUES.ENUM_VALUE] comment added to overcome the problem with the comment map association
	__Link__000005_ENUM_VALUE.Identifier = `ref_models.A_SPECIFIED_VALUES.ENUM_VALUE`

	//gong:ident [ref_models.ENUM_VALUE] comment added to overcome the problem with the comment map association
	__Link__000005_ENUM_VALUE.Fieldtypename = `ref_models.ENUM_VALUE`
	__Link__000005_ENUM_VALUE.FieldOffsetX = -50.000000
	__Link__000005_ENUM_VALUE.FieldOffsetY = -16.000000
	__Link__000005_ENUM_VALUE.TargetMultiplicity = models.MANY
	__Link__000005_ENUM_VALUE.TargetMultiplicityOffsetX = -50.000000
	__Link__000005_ENUM_VALUE.TargetMultiplicityOffsetY = 16.000000
	__Link__000005_ENUM_VALUE.SourceMultiplicity = models.MANY
	__Link__000005_ENUM_VALUE.SourceMultiplicityOffsetX = 10.000000
	__Link__000005_ENUM_VALUE.SourceMultiplicityOffsetY = -50.000000
	__Link__000005_ENUM_VALUE.StartOrientation = models.ORIENTATION_VERTICAL
	__Link__000005_ENUM_VALUE.StartRatio = 0.451389
	__Link__000005_ENUM_VALUE.EndOrientation = models.ORIENTATION_VERTICAL
	__Link__000005_ENUM_VALUE.EndRatio = 0.834722
	__Link__000005_ENUM_VALUE.CornerOffsetRatio = 1.664020

	__Link__000006_PROPERTIES.Name = `PROPERTIES`

	//gong:ident [ref_models.ENUM_VALUE.PROPERTIES] comment added to overcome the problem with the comment map association
	__Link__000006_PROPERTIES.Identifier = `ref_models.ENUM_VALUE.PROPERTIES`

	//gong:ident [ref_models.A_PROPERTIES] comment added to overcome the problem with the comment map association
	__Link__000006_PROPERTIES.Fieldtypename = `ref_models.A_PROPERTIES`
	__Link__000006_PROPERTIES.FieldOffsetX = -50.000000
	__Link__000006_PROPERTIES.FieldOffsetY = -16.000000
	__Link__000006_PROPERTIES.TargetMultiplicity = models.ZERO_ONE
	__Link__000006_PROPERTIES.TargetMultiplicityOffsetX = -50.000000
	__Link__000006_PROPERTIES.TargetMultiplicityOffsetY = 16.000000
	__Link__000006_PROPERTIES.SourceMultiplicity = models.MANY
	__Link__000006_PROPERTIES.SourceMultiplicityOffsetX = 10.000000
	__Link__000006_PROPERTIES.SourceMultiplicityOffsetY = -50.000000
	__Link__000006_PROPERTIES.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000006_PROPERTIES.StartRatio = 0.500000
	__Link__000006_PROPERTIES.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000006_PROPERTIES.EndRatio = 0.473545
	__Link__000006_PROPERTIES.CornerOffsetRatio = 1.959722

	__Link__000007_SPECIFIED_VALUES.Name = `SPECIFIED_VALUES`

	//gong:ident [ref_models.DATATYPE_DEFINITION_ENUMERATION.SPECIFIED_VALUES] comment added to overcome the problem with the comment map association
	__Link__000007_SPECIFIED_VALUES.Identifier = `ref_models.DATATYPE_DEFINITION_ENUMERATION.SPECIFIED_VALUES`

	//gong:ident [ref_models.A_SPECIFIED_VALUES] comment added to overcome the problem with the comment map association
	__Link__000007_SPECIFIED_VALUES.Fieldtypename = `ref_models.A_SPECIFIED_VALUES`
	__Link__000007_SPECIFIED_VALUES.FieldOffsetX = -50.000000
	__Link__000007_SPECIFIED_VALUES.FieldOffsetY = -16.000000
	__Link__000007_SPECIFIED_VALUES.TargetMultiplicity = models.ZERO_ONE
	__Link__000007_SPECIFIED_VALUES.TargetMultiplicityOffsetX = -50.000000
	__Link__000007_SPECIFIED_VALUES.TargetMultiplicityOffsetY = 16.000000
	__Link__000007_SPECIFIED_VALUES.SourceMultiplicity = models.MANY
	__Link__000007_SPECIFIED_VALUES.SourceMultiplicityOffsetX = 10.000000
	__Link__000007_SPECIFIED_VALUES.SourceMultiplicityOffsetY = -50.000000
	__Link__000007_SPECIFIED_VALUES.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000007_SPECIFIED_VALUES.StartRatio = 0.500000
	__Link__000007_SPECIFIED_VALUES.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000007_SPECIFIED_VALUES.EndRatio = 0.500000
	__Link__000007_SPECIFIED_VALUES.CornerOffsetRatio = 1.380000

	__Link__000008_SPEC_ATTRIBUTES.Name = `SPEC_ATTRIBUTES`

	//gong:ident [ref_models.SPEC_OBJECT_TYPE.SPEC_ATTRIBUTES] comment added to overcome the problem with the comment map association
	__Link__000008_SPEC_ATTRIBUTES.Identifier = `ref_models.SPEC_OBJECT_TYPE.SPEC_ATTRIBUTES`

	//gong:ident [ref_models.A_SPEC_ATTRIBUTES] comment added to overcome the problem with the comment map association
	__Link__000008_SPEC_ATTRIBUTES.Fieldtypename = `ref_models.A_SPEC_ATTRIBUTES`
	__Link__000008_SPEC_ATTRIBUTES.FieldOffsetX = -50.000000
	__Link__000008_SPEC_ATTRIBUTES.FieldOffsetY = -16.000000
	__Link__000008_SPEC_ATTRIBUTES.TargetMultiplicity = models.ZERO_ONE
	__Link__000008_SPEC_ATTRIBUTES.TargetMultiplicityOffsetX = -50.000000
	__Link__000008_SPEC_ATTRIBUTES.TargetMultiplicityOffsetY = 16.000000
	__Link__000008_SPEC_ATTRIBUTES.SourceMultiplicity = models.MANY
	__Link__000008_SPEC_ATTRIBUTES.SourceMultiplicityOffsetX = 10.000000
	__Link__000008_SPEC_ATTRIBUTES.SourceMultiplicityOffsetY = -50.000000
	__Link__000008_SPEC_ATTRIBUTES.StartOrientation = models.ORIENTATION_VERTICAL
	__Link__000008_SPEC_ATTRIBUTES.StartRatio = 0.830556
	__Link__000008_SPEC_ATTRIBUTES.EndOrientation = models.ORIENTATION_VERTICAL
	__Link__000008_SPEC_ATTRIBUTES.EndRatio = 0.813889
	__Link__000008_SPEC_ATTRIBUTES.CornerOffsetRatio = 1.473765

	__Link__000009_SPEC_OBJECT_TYPE.Name = `SPEC_OBJECT_TYPE`

	//gong:ident [ref_models.A_SPEC_TYPES.SPEC_OBJECT_TYPE] comment added to overcome the problem with the comment map association
	__Link__000009_SPEC_OBJECT_TYPE.Identifier = `ref_models.A_SPEC_TYPES.SPEC_OBJECT_TYPE`

	//gong:ident [ref_models.SPEC_OBJECT_TYPE] comment added to overcome the problem with the comment map association
	__Link__000009_SPEC_OBJECT_TYPE.Fieldtypename = `ref_models.SPEC_OBJECT_TYPE`
	__Link__000009_SPEC_OBJECT_TYPE.FieldOffsetX = -50.000000
	__Link__000009_SPEC_OBJECT_TYPE.FieldOffsetY = -16.000000
	__Link__000009_SPEC_OBJECT_TYPE.TargetMultiplicity = models.MANY
	__Link__000009_SPEC_OBJECT_TYPE.TargetMultiplicityOffsetX = -50.000000
	__Link__000009_SPEC_OBJECT_TYPE.TargetMultiplicityOffsetY = 16.000000
	__Link__000009_SPEC_OBJECT_TYPE.SourceMultiplicity = models.MANY
	__Link__000009_SPEC_OBJECT_TYPE.SourceMultiplicityOffsetX = 10.000000
	__Link__000009_SPEC_OBJECT_TYPE.SourceMultiplicityOffsetY = -50.000000
	__Link__000009_SPEC_OBJECT_TYPE.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000009_SPEC_OBJECT_TYPE.StartRatio = 0.500000
	__Link__000009_SPEC_OBJECT_TYPE.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000009_SPEC_OBJECT_TYPE.EndRatio = 0.500000
	__Link__000009_SPEC_OBJECT_TYPE.CornerOffsetRatio = 1.380000

	__Link__000010_SPEC_TYPES.Name = `SPEC_TYPES`

	//gong:ident [ref_models.REQ_IF_CONTENT.SPEC_TYPES] comment added to overcome the problem with the comment map association
	__Link__000010_SPEC_TYPES.Identifier = `ref_models.REQ_IF_CONTENT.SPEC_TYPES`

	//gong:ident [ref_models.A_SPEC_TYPES] comment added to overcome the problem with the comment map association
	__Link__000010_SPEC_TYPES.Fieldtypename = `ref_models.A_SPEC_TYPES`
	__Link__000010_SPEC_TYPES.FieldOffsetX = -50.000000
	__Link__000010_SPEC_TYPES.FieldOffsetY = -16.000000
	__Link__000010_SPEC_TYPES.TargetMultiplicity = models.ZERO_ONE
	__Link__000010_SPEC_TYPES.TargetMultiplicityOffsetX = -50.000000
	__Link__000010_SPEC_TYPES.TargetMultiplicityOffsetY = 16.000000
	__Link__000010_SPEC_TYPES.SourceMultiplicity = models.MANY
	__Link__000010_SPEC_TYPES.SourceMultiplicityOffsetX = 10.000000
	__Link__000010_SPEC_TYPES.SourceMultiplicityOffsetY = -50.000000
	__Link__000010_SPEC_TYPES.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000010_SPEC_TYPES.StartRatio = 0.500000
	__Link__000010_SPEC_TYPES.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000010_SPEC_TYPES.EndRatio = 0.500000
	__Link__000010_SPEC_TYPES.CornerOffsetRatio = 1.380000

	__Link__000011_TYPE.Name = `TYPE`

	//gong:ident [ref_models.ATTRIBUTE_DEFINITION_XHTML.TYPE] comment added to overcome the problem with the comment map association
	__Link__000011_TYPE.Identifier = `ref_models.ATTRIBUTE_DEFINITION_XHTML.TYPE`

	//gong:ident [ref_models.A_DATATYPE_DEFINITION_XHTML_REF] comment added to overcome the problem with the comment map association
	__Link__000011_TYPE.Fieldtypename = `ref_models.A_DATATYPE_DEFINITION_XHTML_REF`
	__Link__000011_TYPE.FieldOffsetX = -50.000000
	__Link__000011_TYPE.FieldOffsetY = -16.000000
	__Link__000011_TYPE.TargetMultiplicity = models.ZERO_ONE
	__Link__000011_TYPE.TargetMultiplicityOffsetX = -50.000000
	__Link__000011_TYPE.TargetMultiplicityOffsetY = 16.000000
	__Link__000011_TYPE.SourceMultiplicity = models.MANY
	__Link__000011_TYPE.SourceMultiplicityOffsetX = 10.000000
	__Link__000011_TYPE.SourceMultiplicityOffsetY = -50.000000
	__Link__000011_TYPE.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000011_TYPE.StartRatio = 0.500000
	__Link__000011_TYPE.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000011_TYPE.EndRatio = 0.500000
	__Link__000011_TYPE.CornerOffsetRatio = 1.380000

	__Link__000012_TYPE.Name = `TYPE`

	//gong:ident [ref_models.ATTRIBUTE_DEFINITION_ENUMERATION.TYPE] comment added to overcome the problem with the comment map association
	__Link__000012_TYPE.Identifier = `ref_models.ATTRIBUTE_DEFINITION_ENUMERATION.TYPE`

	//gong:ident [ref_models.A_DATATYPE_DEFINITION_ENUMERATION_REF] comment added to overcome the problem with the comment map association
	__Link__000012_TYPE.Fieldtypename = `ref_models.A_DATATYPE_DEFINITION_ENUMERATION_REF`
	__Link__000012_TYPE.FieldOffsetX = -50.000000
	__Link__000012_TYPE.FieldOffsetY = -16.000000
	__Link__000012_TYPE.TargetMultiplicity = models.ZERO_ONE
	__Link__000012_TYPE.TargetMultiplicityOffsetX = -50.000000
	__Link__000012_TYPE.TargetMultiplicityOffsetY = 16.000000
	__Link__000012_TYPE.SourceMultiplicity = models.MANY
	__Link__000012_TYPE.SourceMultiplicityOffsetX = 10.000000
	__Link__000012_TYPE.SourceMultiplicityOffsetY = -50.000000
	__Link__000012_TYPE.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000012_TYPE.StartRatio = 0.500000
	__Link__000012_TYPE.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000012_TYPE.EndRatio = 0.500000
	__Link__000012_TYPE.CornerOffsetRatio = 1.380000

	__Position__000000_Pos_attribute_ATTRIBUTE_DEFINITION_ENUMERATION.X = 705.000061
	__Position__000000_Pos_attribute_ATTRIBUTE_DEFINITION_ENUMERATION.Y = 489.999969
	__Position__000000_Pos_attribute_ATTRIBUTE_DEFINITION_ENUMERATION.Name = `Pos-attribute-ATTRIBUTE_DEFINITION_ENUMERATION`

	__Position__000001_Pos_attribute_ATTRIBUTE_DEFINITION_XHTML.X = 716.000000
	__Position__000001_Pos_attribute_ATTRIBUTE_DEFINITION_XHTML.Y = 306.000000
	__Position__000001_Pos_attribute_ATTRIBUTE_DEFINITION_XHTML.Name = `Pos-attribute-ATTRIBUTE_DEFINITION_XHTML`

	__Position__000002_Pos_attribute_A_DATATYPES.X = 212.000000
	__Position__000002_Pos_attribute_A_DATATYPES.Y = 926.999969
	__Position__000002_Pos_attribute_A_DATATYPES.Name = `Pos-attribute-A_DATATYPES`

	__Position__000003_Pos_attribute_A_DATATYPE_DEFINITION_ENUMERATION_REF.X = 1224.999939
	__Position__000003_Pos_attribute_A_DATATYPE_DEFINITION_ENUMERATION_REF.Y = 524.999969
	__Position__000003_Pos_attribute_A_DATATYPE_DEFINITION_ENUMERATION_REF.Name = `Pos-attribute-A_DATATYPE_DEFINITION_ENUMERATION_REF`

	__Position__000004_Pos_attribute_A_DATATYPE_DEFINITION_XHTML_REF.X = 1219.999939
	__Position__000004_Pos_attribute_A_DATATYPE_DEFINITION_XHTML_REF.Y = 340.000000
	__Position__000004_Pos_attribute_A_DATATYPE_DEFINITION_XHTML_REF.Name = `Pos-attribute-A_DATATYPE_DEFINITION_XHTML_REF`

	__Position__000005_Pos_attribute_A_PROPERTIES.X = 704.000000
	__Position__000005_Pos_attribute_A_PROPERTIES.Y = 1117.999908
	__Position__000005_Pos_attribute_A_PROPERTIES.Name = `Pos-attribute-A_PROPERTIES`

	__Position__000006_Pos_attribute_A_SPECIFIED_VALUES.X = 1351.999939
	__Position__000006_Pos_attribute_A_SPECIFIED_VALUES.Y = 924.999969
	__Position__000006_Pos_attribute_A_SPECIFIED_VALUES.Name = `Pos-attribute-A_SPECIFIED_VALUES`

	__Position__000007_Pos_attribute_A_SPEC_ATTRIBUTES.X = 125.000000
	__Position__000007_Pos_attribute_A_SPEC_ATTRIBUTES.Y = 317.000000
	__Position__000007_Pos_attribute_A_SPEC_ATTRIBUTES.Name = `Pos-attribute-A_SPEC_ATTRIBUTES`

	__Position__000008_Pos_attribute_A_SPEC_TYPES.X = 469.000000
	__Position__000008_Pos_attribute_A_SPEC_TYPES.Y = 97.000000
	__Position__000008_Pos_attribute_A_SPEC_TYPES.Name = `Pos-attribute-A_SPEC_TYPES`

	__Position__000009_Pos_attribute_DATATYPE_DEFINITION_ENUMERATION.X = 785.000061
	__Position__000009_Pos_attribute_DATATYPE_DEFINITION_ENUMERATION.Y = 899.999969
	__Position__000009_Pos_attribute_DATATYPE_DEFINITION_ENUMERATION.Name = `Pos-attribute-DATATYPE_DEFINITION_ENUMERATION`

	__Position__000010_Pos_attribute_EMBEDDED_VALUE.X = 1293.999939
	__Position__000010_Pos_attribute_EMBEDDED_VALUE.Y = 1109.999908
	__Position__000010_Pos_attribute_EMBEDDED_VALUE.Name = `Pos-attribute-EMBEDDED_VALUE`

	__Position__000011_Pos_attribute_ENUM_VALUE.X = 224.000000
	__Position__000011_Pos_attribute_ENUM_VALUE.Y = 1096.999908
	__Position__000011_Pos_attribute_ENUM_VALUE.Name = `Pos-attribute-ENUM_VALUE`

	__Position__000012_Pos_attribute_REQ_IF_CONTENT.X = 28.000000
	__Position__000012_Pos_attribute_REQ_IF_CONTENT.Y = 95.000000
	__Position__000012_Pos_attribute_REQ_IF_CONTENT.Name = `Pos-attribute-REQ_IF_CONTENT`

	__Position__000013_Pos_attribute_SPEC_OBJECT_TYPE.X = 967.999878
	__Position__000013_Pos_attribute_SPEC_OBJECT_TYPE.Y = 81.000000
	__Position__000013_Pos_attribute_SPEC_OBJECT_TYPE.Name = `Pos-attribute-SPEC_OBJECT_TYPE`

	__Vertice__000000_Verticle_in_class_diagram_attribute_in_middle_between_attribute_ATTRIBUTE_DEFINITION_ENUMERATION_and_attribute_A_DATATYPE_DEFINITION_ENUMERATION_REF.X = 1456.999908
	__Vertice__000000_Verticle_in_class_diagram_attribute_in_middle_between_attribute_ATTRIBUTE_DEFINITION_ENUMERATION_and_attribute_A_DATATYPE_DEFINITION_ENUMERATION_REF.Y = 561.999969
	__Vertice__000000_Verticle_in_class_diagram_attribute_in_middle_between_attribute_ATTRIBUTE_DEFINITION_ENUMERATION_and_attribute_A_DATATYPE_DEFINITION_ENUMERATION_REF.Name = `Verticle in class diagram attribute in middle between attribute-ATTRIBUTE_DEFINITION_ENUMERATION and attribute-A_DATATYPE_DEFINITION_ENUMERATION_REF`

	__Vertice__000001_Verticle_in_class_diagram_attribute_in_middle_between_attribute_ATTRIBUTE_DEFINITION_XHTML_and_attribute_A_DATATYPE_DEFINITION_XHTML_REF.X = 1417.999878
	__Vertice__000001_Verticle_in_class_diagram_attribute_in_middle_between_attribute_ATTRIBUTE_DEFINITION_XHTML_and_attribute_A_DATATYPE_DEFINITION_XHTML_REF.Y = 376.500000
	__Vertice__000001_Verticle_in_class_diagram_attribute_in_middle_between_attribute_ATTRIBUTE_DEFINITION_XHTML_and_attribute_A_DATATYPE_DEFINITION_XHTML_REF.Name = `Verticle in class diagram attribute in middle between attribute-ATTRIBUTE_DEFINITION_XHTML and attribute-A_DATATYPE_DEFINITION_XHTML_REF`

	__Vertice__000002_Verticle_in_class_diagram_attribute_in_middle_between_attribute_A_DATATYPES_and_attribute_DATATYPE_DEFINITION_ENUMERATION.X = 1058.000000
	__Vertice__000002_Verticle_in_class_diagram_attribute_in_middle_between_attribute_A_DATATYPES_and_attribute_DATATYPE_DEFINITION_ENUMERATION.Y = 953.499969
	__Vertice__000002_Verticle_in_class_diagram_attribute_in_middle_between_attribute_A_DATATYPES_and_attribute_DATATYPE_DEFINITION_ENUMERATION.Name = `Verticle in class diagram attribute in middle between attribute-A_DATATYPES and attribute-DATATYPE_DEFINITION_ENUMERATION`

	__Vertice__000003_Verticle_in_class_diagram_attribute_in_middle_between_attribute_A_PROPERTIES_and_attribute_EMBEDDED_VALUE.X = 1358.999969
	__Vertice__000003_Verticle_in_class_diagram_attribute_in_middle_between_attribute_A_PROPERTIES_and_attribute_EMBEDDED_VALUE.Y = 1145.499908
	__Vertice__000003_Verticle_in_class_diagram_attribute_in_middle_between_attribute_A_PROPERTIES_and_attribute_EMBEDDED_VALUE.Name = `Verticle in class diagram attribute in middle between attribute-A_PROPERTIES and attribute-EMBEDDED_VALUE`

	__Vertice__000004_Verticle_in_class_diagram_attribute_in_middle_between_attribute_A_SPECIFIED_VALUES_and_attribute_ENUM_VALUE.X = 1708.499939
	__Vertice__000004_Verticle_in_class_diagram_attribute_in_middle_between_attribute_A_SPECIFIED_VALUES_and_attribute_ENUM_VALUE.Y = 884.499969
	__Vertice__000004_Verticle_in_class_diagram_attribute_in_middle_between_attribute_A_SPECIFIED_VALUES_and_attribute_ENUM_VALUE.Name = `Verticle in class diagram attribute in middle between attribute-A_SPECIFIED_VALUES and attribute-ENUM_VALUE`

	__Vertice__000005_Verticle_in_class_diagram_attribute_in_middle_between_attribute_A_SPEC_ATTRIBUTES_and_attribute_ATTRIBUTE_DEFINITION_ENUMERATION.X = 727.500031
	__Vertice__000005_Verticle_in_class_diagram_attribute_in_middle_between_attribute_A_SPEC_ATTRIBUTES_and_attribute_ATTRIBUTE_DEFINITION_ENUMERATION.Y = 690.499985
	__Vertice__000005_Verticle_in_class_diagram_attribute_in_middle_between_attribute_A_SPEC_ATTRIBUTES_and_attribute_ATTRIBUTE_DEFINITION_ENUMERATION.Name = `Verticle in class diagram attribute in middle between attribute-A_SPEC_ATTRIBUTES and attribute-ATTRIBUTE_DEFINITION_ENUMERATION`

	__Vertice__000006_Verticle_in_class_diagram_attribute_in_middle_between_attribute_A_SPEC_ATTRIBUTES_and_attribute_ATTRIBUTE_DEFINITION_XHTML.X = 748.000000
	__Vertice__000006_Verticle_in_class_diagram_attribute_in_middle_between_attribute_A_SPEC_ATTRIBUTES_and_attribute_ATTRIBUTE_DEFINITION_XHTML.Y = 601.000000
	__Vertice__000006_Verticle_in_class_diagram_attribute_in_middle_between_attribute_A_SPEC_ATTRIBUTES_and_attribute_ATTRIBUTE_DEFINITION_XHTML.Name = `Verticle in class diagram attribute in middle between attribute-A_SPEC_ATTRIBUTES and attribute-ATTRIBUTE_DEFINITION_XHTML`

	__Vertice__000007_Verticle_in_class_diagram_attribute_in_middle_between_attribute_A_SPEC_TYPES_and_attribute_SPEC_OBJECT_TYPE.X = 1031.999969
	__Vertice__000007_Verticle_in_class_diagram_attribute_in_middle_between_attribute_A_SPEC_TYPES_and_attribute_SPEC_OBJECT_TYPE.Y = 129.500000
	__Vertice__000007_Verticle_in_class_diagram_attribute_in_middle_between_attribute_A_SPEC_TYPES_and_attribute_SPEC_OBJECT_TYPE.Name = `Verticle in class diagram attribute in middle between attribute-A_SPEC_TYPES and attribute-SPEC_OBJECT_TYPE`

	__Vertice__000008_Verticle_in_class_diagram_attribute_in_middle_between_attribute_DATATYPE_DEFINITION_ENUMERATION_and_attribute_A_SPECIFIED_VALUES.X = 1524.000000
	__Vertice__000008_Verticle_in_class_diagram_attribute_in_middle_between_attribute_DATATYPE_DEFINITION_ENUMERATION_and_attribute_A_SPECIFIED_VALUES.Y = 964.499969
	__Vertice__000008_Verticle_in_class_diagram_attribute_in_middle_between_attribute_DATATYPE_DEFINITION_ENUMERATION_and_attribute_A_SPECIFIED_VALUES.Name = `Verticle in class diagram attribute in middle between attribute-DATATYPE_DEFINITION_ENUMERATION and attribute-A_SPECIFIED_VALUES`

	__Vertice__000009_Verticle_in_class_diagram_attribute_in_middle_between_attribute_ENUM_VALUE_and_attribute_A_PROPERTIES.X = 1709.999939
	__Vertice__000009_Verticle_in_class_diagram_attribute_in_middle_between_attribute_ENUM_VALUE_and_attribute_A_PROPERTIES.Y = 793.499969
	__Vertice__000009_Verticle_in_class_diagram_attribute_in_middle_between_attribute_ENUM_VALUE_and_attribute_A_PROPERTIES.Name = `Verticle in class diagram attribute in middle between attribute-ENUM_VALUE and attribute-A_PROPERTIES`

	__Vertice__000010_Verticle_in_class_diagram_attribute_in_middle_between_attribute_REQ_IF_CONTENT_and_attribute_A_DATATYPES.X = 600.500000
	__Vertice__000010_Verticle_in_class_diagram_attribute_in_middle_between_attribute_REQ_IF_CONTENT_and_attribute_A_DATATYPES.Y = 540.499985
	__Vertice__000010_Verticle_in_class_diagram_attribute_in_middle_between_attribute_REQ_IF_CONTENT_and_attribute_A_DATATYPES.Name = `Verticle in class diagram attribute in middle between attribute-REQ_IF_CONTENT and attribute-A_DATATYPES`

	__Vertice__000011_Verticle_in_class_diagram_attribute_in_middle_between_attribute_REQ_IF_CONTENT_and_attribute_A_SPEC_TYPES.X = 563.500000
	__Vertice__000011_Verticle_in_class_diagram_attribute_in_middle_between_attribute_REQ_IF_CONTENT_and_attribute_A_SPEC_TYPES.Y = 126.500000
	__Vertice__000011_Verticle_in_class_diagram_attribute_in_middle_between_attribute_REQ_IF_CONTENT_and_attribute_A_SPEC_TYPES.Name = `Verticle in class diagram attribute in middle between attribute-REQ_IF_CONTENT and attribute-A_SPEC_TYPES`

	__Vertice__000012_Verticle_in_class_diagram_attribute_in_middle_between_attribute_SPEC_OBJECT_TYPE_and_attribute_A_SPEC_ATTRIBUTES.X = 863.499939
	__Vertice__000012_Verticle_in_class_diagram_attribute_in_middle_between_attribute_SPEC_OBJECT_TYPE_and_attribute_A_SPEC_ATTRIBUTES.Y = 102.000000
	__Vertice__000012_Verticle_in_class_diagram_attribute_in_middle_between_attribute_SPEC_OBJECT_TYPE_and_attribute_A_SPEC_ATTRIBUTES.Name = `Verticle in class diagram attribute in middle between attribute-SPEC_OBJECT_TYPE and attribute-A_SPEC_ATTRIBUTES`

	// Setup of pointers
	__Classdiagram__000000_attribute.GongStructShapes = append(__Classdiagram__000000_attribute.GongStructShapes, __GongStructShape__000012_attribute_REQ_IF_CONTENT)
	__Classdiagram__000000_attribute.GongStructShapes = append(__Classdiagram__000000_attribute.GongStructShapes, __GongStructShape__000008_attribute_A_SPEC_TYPES)
	__Classdiagram__000000_attribute.GongStructShapes = append(__Classdiagram__000000_attribute.GongStructShapes, __GongStructShape__000013_attribute_SPEC_OBJECT_TYPE)
	__Classdiagram__000000_attribute.GongStructShapes = append(__Classdiagram__000000_attribute.GongStructShapes, __GongStructShape__000007_attribute_A_SPEC_ATTRIBUTES)
	__Classdiagram__000000_attribute.GongStructShapes = append(__Classdiagram__000000_attribute.GongStructShapes, __GongStructShape__000001_attribute_ATTRIBUTE_DEFINITION_XHTML)
	__Classdiagram__000000_attribute.GongStructShapes = append(__Classdiagram__000000_attribute.GongStructShapes, __GongStructShape__000004_attribute_A_DATATYPE_DEFINITION_XHTML_REF)
	__Classdiagram__000000_attribute.GongStructShapes = append(__Classdiagram__000000_attribute.GongStructShapes, __GongStructShape__000000_attribute_ATTRIBUTE_DEFINITION_ENUMERATION)
	__Classdiagram__000000_attribute.GongStructShapes = append(__Classdiagram__000000_attribute.GongStructShapes, __GongStructShape__000003_attribute_A_DATATYPE_DEFINITION_ENUMERATION_REF)
	__Classdiagram__000000_attribute.GongStructShapes = append(__Classdiagram__000000_attribute.GongStructShapes, __GongStructShape__000002_attribute_A_DATATYPES)
	__Classdiagram__000000_attribute.GongStructShapes = append(__Classdiagram__000000_attribute.GongStructShapes, __GongStructShape__000009_attribute_DATATYPE_DEFINITION_ENUMERATION)
	__Classdiagram__000000_attribute.GongStructShapes = append(__Classdiagram__000000_attribute.GongStructShapes, __GongStructShape__000006_attribute_A_SPECIFIED_VALUES)
	__Classdiagram__000000_attribute.GongStructShapes = append(__Classdiagram__000000_attribute.GongStructShapes, __GongStructShape__000011_attribute_ENUM_VALUE)
	__Classdiagram__000000_attribute.GongStructShapes = append(__Classdiagram__000000_attribute.GongStructShapes, __GongStructShape__000005_attribute_A_PROPERTIES)
	__Classdiagram__000000_attribute.GongStructShapes = append(__Classdiagram__000000_attribute.GongStructShapes, __GongStructShape__000010_attribute_EMBEDDED_VALUE)
	__GongStructShape__000000_attribute_ATTRIBUTE_DEFINITION_ENUMERATION.Position = __Position__000000_Pos_attribute_ATTRIBUTE_DEFINITION_ENUMERATION
	__GongStructShape__000000_attribute_ATTRIBUTE_DEFINITION_ENUMERATION.Fields = append(__GongStructShape__000000_attribute_ATTRIBUTE_DEFINITION_ENUMERATION.Fields, __Field__000005_IDENTIFIER)
	__GongStructShape__000000_attribute_ATTRIBUTE_DEFINITION_ENUMERATION.Fields = append(__GongStructShape__000000_attribute_ATTRIBUTE_DEFINITION_ENUMERATION.Fields, __Field__000007_IS_EDITABLE)
	__GongStructShape__000000_attribute_ATTRIBUTE_DEFINITION_ENUMERATION.Fields = append(__GongStructShape__000000_attribute_ATTRIBUTE_DEFINITION_ENUMERATION.Fields, __Field__000010_LAST_CHANGE)
	__GongStructShape__000000_attribute_ATTRIBUTE_DEFINITION_ENUMERATION.Fields = append(__GongStructShape__000000_attribute_ATTRIBUTE_DEFINITION_ENUMERATION.Fields, __Field__000016_LONG_NAME)
	__GongStructShape__000000_attribute_ATTRIBUTE_DEFINITION_ENUMERATION.Fields = append(__GongStructShape__000000_attribute_ATTRIBUTE_DEFINITION_ENUMERATION.Fields, __Field__000020_MULTI_VALUED)
	__GongStructShape__000000_attribute_ATTRIBUTE_DEFINITION_ENUMERATION.Links = append(__GongStructShape__000000_attribute_ATTRIBUTE_DEFINITION_ENUMERATION.Links, __Link__000012_TYPE)
	__GongStructShape__000001_attribute_ATTRIBUTE_DEFINITION_XHTML.Position = __Position__000001_Pos_attribute_ATTRIBUTE_DEFINITION_XHTML
	__GongStructShape__000001_attribute_ATTRIBUTE_DEFINITION_XHTML.Fields = append(__GongStructShape__000001_attribute_ATTRIBUTE_DEFINITION_XHTML.Fields, __Field__000003_IDENTIFIER)
	__GongStructShape__000001_attribute_ATTRIBUTE_DEFINITION_XHTML.Fields = append(__GongStructShape__000001_attribute_ATTRIBUTE_DEFINITION_XHTML.Fields, __Field__000008_IS_EDITABLE)
	__GongStructShape__000001_attribute_ATTRIBUTE_DEFINITION_XHTML.Fields = append(__GongStructShape__000001_attribute_ATTRIBUTE_DEFINITION_XHTML.Fields, __Field__000011_LAST_CHANGE)
	__GongStructShape__000001_attribute_ATTRIBUTE_DEFINITION_XHTML.Fields = append(__GongStructShape__000001_attribute_ATTRIBUTE_DEFINITION_XHTML.Fields, __Field__000017_LONG_NAME)
	__GongStructShape__000001_attribute_ATTRIBUTE_DEFINITION_XHTML.Links = append(__GongStructShape__000001_attribute_ATTRIBUTE_DEFINITION_XHTML.Links, __Link__000011_TYPE)
	__GongStructShape__000002_attribute_A_DATATYPES.Position = __Position__000002_Pos_attribute_A_DATATYPES
	__GongStructShape__000002_attribute_A_DATATYPES.Links = append(__GongStructShape__000002_attribute_A_DATATYPES.Links, __Link__000003_DATATYPE_DEFINITION_ENUMERATION)
	__GongStructShape__000003_attribute_A_DATATYPE_DEFINITION_ENUMERATION_REF.Position = __Position__000003_Pos_attribute_A_DATATYPE_DEFINITION_ENUMERATION_REF
	__GongStructShape__000003_attribute_A_DATATYPE_DEFINITION_ENUMERATION_REF.Fields = append(__GongStructShape__000003_attribute_A_DATATYPE_DEFINITION_ENUMERATION_REF.Fields, __Field__000000_DATATYPE_DEFINITION_ENUMERATION_REF)
	__GongStructShape__000004_attribute_A_DATATYPE_DEFINITION_XHTML_REF.Position = __Position__000004_Pos_attribute_A_DATATYPE_DEFINITION_XHTML_REF
	__GongStructShape__000004_attribute_A_DATATYPE_DEFINITION_XHTML_REF.Fields = append(__GongStructShape__000004_attribute_A_DATATYPE_DEFINITION_XHTML_REF.Fields, __Field__000001_DATATYPE_DEFINITION_XHTML_REF)
	__GongStructShape__000005_attribute_A_PROPERTIES.Position = __Position__000005_Pos_attribute_A_PROPERTIES
	__GongStructShape__000005_attribute_A_PROPERTIES.Links = append(__GongStructShape__000005_attribute_A_PROPERTIES.Links, __Link__000004_EMBEDDED_VALUE)
	__GongStructShape__000006_attribute_A_SPECIFIED_VALUES.Position = __Position__000006_Pos_attribute_A_SPECIFIED_VALUES
	__GongStructShape__000006_attribute_A_SPECIFIED_VALUES.Links = append(__GongStructShape__000006_attribute_A_SPECIFIED_VALUES.Links, __Link__000005_ENUM_VALUE)
	__GongStructShape__000007_attribute_A_SPEC_ATTRIBUTES.Position = __Position__000007_Pos_attribute_A_SPEC_ATTRIBUTES
	__GongStructShape__000007_attribute_A_SPEC_ATTRIBUTES.Links = append(__GongStructShape__000007_attribute_A_SPEC_ATTRIBUTES.Links, __Link__000001_ATTRIBUTE_DEFINITION_XHTML)
	__GongStructShape__000007_attribute_A_SPEC_ATTRIBUTES.Links = append(__GongStructShape__000007_attribute_A_SPEC_ATTRIBUTES.Links, __Link__000000_ATTRIBUTE_DEFINITION_ENUMERATION)
	__GongStructShape__000008_attribute_A_SPEC_TYPES.Position = __Position__000008_Pos_attribute_A_SPEC_TYPES
	__GongStructShape__000008_attribute_A_SPEC_TYPES.Links = append(__GongStructShape__000008_attribute_A_SPEC_TYPES.Links, __Link__000009_SPEC_OBJECT_TYPE)
	__GongStructShape__000009_attribute_DATATYPE_DEFINITION_ENUMERATION.Position = __Position__000009_Pos_attribute_DATATYPE_DEFINITION_ENUMERATION
	__GongStructShape__000009_attribute_DATATYPE_DEFINITION_ENUMERATION.Fields = append(__GongStructShape__000009_attribute_DATATYPE_DEFINITION_ENUMERATION.Fields, __Field__000006_IDENTIFIER)
	__GongStructShape__000009_attribute_DATATYPE_DEFINITION_ENUMERATION.Fields = append(__GongStructShape__000009_attribute_DATATYPE_DEFINITION_ENUMERATION.Fields, __Field__000014_LAST_CHANGE)
	__GongStructShape__000009_attribute_DATATYPE_DEFINITION_ENUMERATION.Fields = append(__GongStructShape__000009_attribute_DATATYPE_DEFINITION_ENUMERATION.Fields, __Field__000018_LONG_NAME)
	__GongStructShape__000009_attribute_DATATYPE_DEFINITION_ENUMERATION.Links = append(__GongStructShape__000009_attribute_DATATYPE_DEFINITION_ENUMERATION.Links, __Link__000007_SPECIFIED_VALUES)
	__GongStructShape__000010_attribute_EMBEDDED_VALUE.Position = __Position__000010_Pos_attribute_EMBEDDED_VALUE
	__GongStructShape__000010_attribute_EMBEDDED_VALUE.Fields = append(__GongStructShape__000010_attribute_EMBEDDED_VALUE.Fields, __Field__000009_KEY)
	__GongStructShape__000010_attribute_EMBEDDED_VALUE.Fields = append(__GongStructShape__000010_attribute_EMBEDDED_VALUE.Fields, __Field__000021_OTHER_CONTENT)
	__GongStructShape__000011_attribute_ENUM_VALUE.Position = __Position__000011_Pos_attribute_ENUM_VALUE
	__GongStructShape__000011_attribute_ENUM_VALUE.Fields = append(__GongStructShape__000011_attribute_ENUM_VALUE.Fields, __Field__000004_IDENTIFIER)
	__GongStructShape__000011_attribute_ENUM_VALUE.Fields = append(__GongStructShape__000011_attribute_ENUM_VALUE.Fields, __Field__000013_LAST_CHANGE)
	__GongStructShape__000011_attribute_ENUM_VALUE.Fields = append(__GongStructShape__000011_attribute_ENUM_VALUE.Fields, __Field__000015_LONG_NAME)
	__GongStructShape__000011_attribute_ENUM_VALUE.Links = append(__GongStructShape__000011_attribute_ENUM_VALUE.Links, __Link__000006_PROPERTIES)
	__GongStructShape__000012_attribute_REQ_IF_CONTENT.Position = __Position__000012_Pos_attribute_REQ_IF_CONTENT
	__GongStructShape__000012_attribute_REQ_IF_CONTENT.Links = append(__GongStructShape__000012_attribute_REQ_IF_CONTENT.Links, __Link__000010_SPEC_TYPES)
	__GongStructShape__000012_attribute_REQ_IF_CONTENT.Links = append(__GongStructShape__000012_attribute_REQ_IF_CONTENT.Links, __Link__000002_DATATYPES)
	__GongStructShape__000013_attribute_SPEC_OBJECT_TYPE.Position = __Position__000013_Pos_attribute_SPEC_OBJECT_TYPE
	__GongStructShape__000013_attribute_SPEC_OBJECT_TYPE.Fields = append(__GongStructShape__000013_attribute_SPEC_OBJECT_TYPE.Fields, __Field__000002_IDENTIFIER)
	__GongStructShape__000013_attribute_SPEC_OBJECT_TYPE.Fields = append(__GongStructShape__000013_attribute_SPEC_OBJECT_TYPE.Fields, __Field__000012_LAST_CHANGE)
	__GongStructShape__000013_attribute_SPEC_OBJECT_TYPE.Fields = append(__GongStructShape__000013_attribute_SPEC_OBJECT_TYPE.Fields, __Field__000019_LONG_NAME)
	__GongStructShape__000013_attribute_SPEC_OBJECT_TYPE.Links = append(__GongStructShape__000013_attribute_SPEC_OBJECT_TYPE.Links, __Link__000008_SPEC_ATTRIBUTES)
	__Link__000000_ATTRIBUTE_DEFINITION_ENUMERATION.Middlevertice = __Vertice__000005_Verticle_in_class_diagram_attribute_in_middle_between_attribute_A_SPEC_ATTRIBUTES_and_attribute_ATTRIBUTE_DEFINITION_ENUMERATION
	__Link__000001_ATTRIBUTE_DEFINITION_XHTML.Middlevertice = __Vertice__000006_Verticle_in_class_diagram_attribute_in_middle_between_attribute_A_SPEC_ATTRIBUTES_and_attribute_ATTRIBUTE_DEFINITION_XHTML
	__Link__000002_DATATYPES.Middlevertice = __Vertice__000010_Verticle_in_class_diagram_attribute_in_middle_between_attribute_REQ_IF_CONTENT_and_attribute_A_DATATYPES
	__Link__000003_DATATYPE_DEFINITION_ENUMERATION.Middlevertice = __Vertice__000002_Verticle_in_class_diagram_attribute_in_middle_between_attribute_A_DATATYPES_and_attribute_DATATYPE_DEFINITION_ENUMERATION
	__Link__000004_EMBEDDED_VALUE.Middlevertice = __Vertice__000003_Verticle_in_class_diagram_attribute_in_middle_between_attribute_A_PROPERTIES_and_attribute_EMBEDDED_VALUE
	__Link__000005_ENUM_VALUE.Middlevertice = __Vertice__000004_Verticle_in_class_diagram_attribute_in_middle_between_attribute_A_SPECIFIED_VALUES_and_attribute_ENUM_VALUE
	__Link__000006_PROPERTIES.Middlevertice = __Vertice__000009_Verticle_in_class_diagram_attribute_in_middle_between_attribute_ENUM_VALUE_and_attribute_A_PROPERTIES
	__Link__000007_SPECIFIED_VALUES.Middlevertice = __Vertice__000008_Verticle_in_class_diagram_attribute_in_middle_between_attribute_DATATYPE_DEFINITION_ENUMERATION_and_attribute_A_SPECIFIED_VALUES
	__Link__000008_SPEC_ATTRIBUTES.Middlevertice = __Vertice__000012_Verticle_in_class_diagram_attribute_in_middle_between_attribute_SPEC_OBJECT_TYPE_and_attribute_A_SPEC_ATTRIBUTES
	__Link__000009_SPEC_OBJECT_TYPE.Middlevertice = __Vertice__000007_Verticle_in_class_diagram_attribute_in_middle_between_attribute_A_SPEC_TYPES_and_attribute_SPEC_OBJECT_TYPE
	__Link__000010_SPEC_TYPES.Middlevertice = __Vertice__000011_Verticle_in_class_diagram_attribute_in_middle_between_attribute_REQ_IF_CONTENT_and_attribute_A_SPEC_TYPES
	__Link__000011_TYPE.Middlevertice = __Vertice__000001_Verticle_in_class_diagram_attribute_in_middle_between_attribute_ATTRIBUTE_DEFINITION_XHTML_and_attribute_A_DATATYPE_DEFINITION_XHTML_REF
	__Link__000012_TYPE.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_attribute_in_middle_between_attribute_ATTRIBUTE_DEFINITION_ENUMERATION_and_attribute_A_DATATYPE_DEFINITION_ENUMERATION_REF
}
