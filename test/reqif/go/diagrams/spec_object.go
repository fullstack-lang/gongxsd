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

	__Classdiagram__000000_spec_object := (&models.Classdiagram{Name: `spec_object`}).Stage(stage)

	__Field__000000_ATTRIBUTE_DEFINITION_XHTML_REF := (&models.Field{Name: `ATTRIBUTE_DEFINITION_XHTML_REF`}).Stage(stage)
	__Field__000001_EnclosedText := (&models.Field{Name: `EnclosedText`}).Stage(stage)
	__Field__000002_IDENTIFIER := (&models.Field{Name: `IDENTIFIER`}).Stage(stage)
	__Field__000003_LAST_CHANGE := (&models.Field{Name: `LAST_CHANGE`}).Stage(stage)
	__Field__000004_LONG_NAME := (&models.Field{Name: `LONG_NAME`}).Stage(stage)

	__GongStructShape__000000_spec_object_ATTRIBUTE_VALUE_ENUMERATION := (&models.GongStructShape{Name: `spec_object-ATTRIBUTE_VALUE_ENUMERATION`}).Stage(stage)
	__GongStructShape__000001_spec_object_ATTRIBUTE_VALUE_XHTML := (&models.GongStructShape{Name: `spec_object-ATTRIBUTE_VALUE_XHTML`}).Stage(stage)
	__GongStructShape__000002_spec_object_A_ATTRIBUTE_DEFINITION_ENUMERATION_REF := (&models.GongStructShape{Name: `spec_object-A_ATTRIBUTE_DEFINITION_ENUMERATION_REF`}).Stage(stage)
	__GongStructShape__000003_spec_object_A_ATTRIBUTE_DEFINITION_XHTML_REF := (&models.GongStructShape{Name: `spec_object-A_ATTRIBUTE_DEFINITION_XHTML_REF`}).Stage(stage)
	__GongStructShape__000004_spec_object_A_ATTRIBUTE_VALUE_XHTML_1 := (&models.GongStructShape{Name: `spec_object-A_ATTRIBUTE_VALUE_XHTML_1`}).Stage(stage)
	__GongStructShape__000005_spec_object_A_ENUM_VALUE_REF := (&models.GongStructShape{Name: `spec_object-A_ENUM_VALUE_REF`}).Stage(stage)
	__GongStructShape__000006_spec_object_A_SPEC_OBJECT_TYPE_REF := (&models.GongStructShape{Name: `spec_object-A_SPEC_OBJECT_TYPE_REF`}).Stage(stage)
	__GongStructShape__000007_spec_object_SPEC_OBJECT := (&models.GongStructShape{Name: `spec_object-SPEC_OBJECT`}).Stage(stage)
	__GongStructShape__000008_spec_object_XHTML_CONTENT := (&models.GongStructShape{Name: `spec_object-XHTML_CONTENT`}).Stage(stage)

	__Link__000000_ATTRIBUTE_VALUE_ENUMERATION := (&models.Link{Name: `ATTRIBUTE_VALUE_ENUMERATION`}).Stage(stage)
	__Link__000001_ATTRIBUTE_VALUE_XHTML := (&models.Link{Name: `ATTRIBUTE_VALUE_XHTML`}).Stage(stage)
	__Link__000002_DEFINITION := (&models.Link{Name: `DEFINITION`}).Stage(stage)
	__Link__000003_DEFINITION := (&models.Link{Name: `DEFINITION`}).Stage(stage)
	__Link__000004_THE_ORIGINAL_VALUE := (&models.Link{Name: `THE_ORIGINAL_VALUE`}).Stage(stage)
	__Link__000005_THE_VALUE := (&models.Link{Name: `THE_VALUE`}).Stage(stage)
	__Link__000006_TYPE := (&models.Link{Name: `TYPE`}).Stage(stage)
	__Link__000007_VALUES := (&models.Link{Name: `VALUES`}).Stage(stage)
	__Link__000008_VALUES := (&models.Link{Name: `VALUES`}).Stage(stage)

	__Position__000000_Pos_spec_object_ATTRIBUTE_VALUE_ENUMERATION := (&models.Position{Name: `Pos-spec_object-ATTRIBUTE_VALUE_ENUMERATION`}).Stage(stage)
	__Position__000001_Pos_spec_object_ATTRIBUTE_VALUE_XHTML := (&models.Position{Name: `Pos-spec_object-ATTRIBUTE_VALUE_XHTML`}).Stage(stage)
	__Position__000002_Pos_spec_object_A_ATTRIBUTE_DEFINITION_ENUMERATION_REF := (&models.Position{Name: `Pos-spec_object-A_ATTRIBUTE_DEFINITION_ENUMERATION_REF`}).Stage(stage)
	__Position__000003_Pos_spec_object_A_ATTRIBUTE_DEFINITION_XHTML_REF := (&models.Position{Name: `Pos-spec_object-A_ATTRIBUTE_DEFINITION_XHTML_REF`}).Stage(stage)
	__Position__000004_Pos_spec_object_A_ATTRIBUTE_VALUE_XHTML_1 := (&models.Position{Name: `Pos-spec_object-A_ATTRIBUTE_VALUE_XHTML_1`}).Stage(stage)
	__Position__000005_Pos_spec_object_A_ENUM_VALUE_REF := (&models.Position{Name: `Pos-spec_object-A_ENUM_VALUE_REF`}).Stage(stage)
	__Position__000006_Pos_spec_object_A_SPEC_OBJECT_TYPE_REF := (&models.Position{Name: `Pos-spec_object-A_SPEC_OBJECT_TYPE_REF`}).Stage(stage)
	__Position__000007_Pos_spec_object_SPEC_OBJECT := (&models.Position{Name: `Pos-spec_object-SPEC_OBJECT`}).Stage(stage)
	__Position__000008_Pos_spec_object_XHTML_CONTENT := (&models.Position{Name: `Pos-spec_object-XHTML_CONTENT`}).Stage(stage)

	__Vertice__000000_Verticle_in_class_diagram_spec_object_in_middle_between_spec_object_ATTRIBUTE_VALUE_ENUMERATION_and_spec_object_A_ATTRIBUTE_DEFINITION_ENUMERATION_REF := (&models.Vertice{Name: `Verticle in class diagram spec_object in middle between spec_object-ATTRIBUTE_VALUE_ENUMERATION and spec_object-A_ATTRIBUTE_DEFINITION_ENUMERATION_REF`}).Stage(stage)
	__Vertice__000001_Verticle_in_class_diagram_spec_object_in_middle_between_spec_object_ATTRIBUTE_VALUE_ENUMERATION_and_spec_object_A_ENUM_VALUE_REF := (&models.Vertice{Name: `Verticle in class diagram spec_object in middle between spec_object-ATTRIBUTE_VALUE_ENUMERATION and spec_object-A_ENUM_VALUE_REF`}).Stage(stage)
	__Vertice__000002_Verticle_in_class_diagram_spec_object_in_middle_between_spec_object_ATTRIBUTE_VALUE_XHTML_and_spec_object_A_ATTRIBUTE_DEFINITION_XHTML_REF := (&models.Vertice{Name: `Verticle in class diagram spec_object in middle between spec_object-ATTRIBUTE_VALUE_XHTML and spec_object-A_ATTRIBUTE_DEFINITION_XHTML_REF`}).Stage(stage)
	__Vertice__000003_Verticle_in_class_diagram_spec_object_in_middle_between_spec_object_ATTRIBUTE_VALUE_XHTML_and_spec_object_XHTML_CONTENT := (&models.Vertice{Name: `Verticle in class diagram spec_object in middle between spec_object-ATTRIBUTE_VALUE_XHTML and spec_object-XHTML_CONTENT`}).Stage(stage)
	__Vertice__000004_Verticle_in_class_diagram_spec_object_in_middle_between_spec_object_ATTRIBUTE_VALUE_XHTML_and_spec_object_XHTML_CONTENT := (&models.Vertice{Name: `Verticle in class diagram spec_object in middle between spec_object-ATTRIBUTE_VALUE_XHTML and spec_object-XHTML_CONTENT`}).Stage(stage)
	__Vertice__000005_Verticle_in_class_diagram_spec_object_in_middle_between_spec_object_A_ATTRIBUTE_VALUE_XHTML_1_and_spec_object_ATTRIBUTE_VALUE_ENUMERATION := (&models.Vertice{Name: `Verticle in class diagram spec_object in middle between spec_object-A_ATTRIBUTE_VALUE_XHTML_1 and spec_object-ATTRIBUTE_VALUE_ENUMERATION`}).Stage(stage)
	__Vertice__000006_Verticle_in_class_diagram_spec_object_in_middle_between_spec_object_A_ATTRIBUTE_VALUE_XHTML_1_and_spec_object_ATTRIBUTE_VALUE_XHTML := (&models.Vertice{Name: `Verticle in class diagram spec_object in middle between spec_object-A_ATTRIBUTE_VALUE_XHTML_1 and spec_object-ATTRIBUTE_VALUE_XHTML`}).Stage(stage)
	__Vertice__000007_Verticle_in_class_diagram_spec_object_in_middle_between_spec_object_SPEC_OBJECT_and_spec_object_A_ATTRIBUTE_VALUE_XHTML_1 := (&models.Vertice{Name: `Verticle in class diagram spec_object in middle between spec_object-SPEC_OBJECT and spec_object-A_ATTRIBUTE_VALUE_XHTML_1`}).Stage(stage)
	__Vertice__000008_Verticle_in_class_diagram_spec_object_in_middle_between_spec_object_SPEC_OBJECT_and_spec_object_A_SPEC_OBJECT_TYPE_REF := (&models.Vertice{Name: `Verticle in class diagram spec_object in middle between spec_object-SPEC_OBJECT and spec_object-A_SPEC_OBJECT_TYPE_REF`}).Stage(stage)

	// Setup of values

	__Classdiagram__000000_spec_object.Name = `spec_object`
	__Classdiagram__000000_spec_object.IsInDrawMode = false

	__Field__000000_ATTRIBUTE_DEFINITION_XHTML_REF.Name = `ATTRIBUTE_DEFINITION_XHTML_REF`

	//gong:ident [ref_models.A_ATTRIBUTE_DEFINITION_XHTML_REF.ATTRIBUTE_DEFINITION_XHTML_REF] comment added to overcome the problem with the comment map association
	__Field__000000_ATTRIBUTE_DEFINITION_XHTML_REF.Identifier = `ref_models.A_ATTRIBUTE_DEFINITION_XHTML_REF.ATTRIBUTE_DEFINITION_XHTML_REF`
	__Field__000000_ATTRIBUTE_DEFINITION_XHTML_REF.FieldTypeAsString = ``
	__Field__000000_ATTRIBUTE_DEFINITION_XHTML_REF.Structname = `A_ATTRIBUTE_DEFINITION_XHTML_REF`
	__Field__000000_ATTRIBUTE_DEFINITION_XHTML_REF.Fieldtypename = `string`

	__Field__000001_EnclosedText.Name = `EnclosedText`

	//gong:ident [ref_models.XHTML_CONTENT.EnclosedText] comment added to overcome the problem with the comment map association
	__Field__000001_EnclosedText.Identifier = `ref_models.XHTML_CONTENT.EnclosedText`
	__Field__000001_EnclosedText.FieldTypeAsString = ``
	__Field__000001_EnclosedText.Structname = `XHTML_CONTENT`
	__Field__000001_EnclosedText.Fieldtypename = `string`

	__Field__000002_IDENTIFIER.Name = `IDENTIFIER`

	//gong:ident [ref_models.SPEC_OBJECT.IDENTIFIER] comment added to overcome the problem with the comment map association
	__Field__000002_IDENTIFIER.Identifier = `ref_models.SPEC_OBJECT.IDENTIFIER`
	__Field__000002_IDENTIFIER.FieldTypeAsString = ``
	__Field__000002_IDENTIFIER.Structname = `SPEC_OBJECT`
	__Field__000002_IDENTIFIER.Fieldtypename = `string`

	__Field__000003_LAST_CHANGE.Name = `LAST_CHANGE`

	//gong:ident [ref_models.SPEC_OBJECT.LAST_CHANGE] comment added to overcome the problem with the comment map association
	__Field__000003_LAST_CHANGE.Identifier = `ref_models.SPEC_OBJECT.LAST_CHANGE`
	__Field__000003_LAST_CHANGE.FieldTypeAsString = ``
	__Field__000003_LAST_CHANGE.Structname = `SPEC_OBJECT`
	__Field__000003_LAST_CHANGE.Fieldtypename = `string`

	__Field__000004_LONG_NAME.Name = `LONG_NAME`

	//gong:ident [ref_models.SPEC_OBJECT.LONG_NAME] comment added to overcome the problem with the comment map association
	__Field__000004_LONG_NAME.Identifier = `ref_models.SPEC_OBJECT.LONG_NAME`
	__Field__000004_LONG_NAME.FieldTypeAsString = ``
	__Field__000004_LONG_NAME.Structname = `SPEC_OBJECT`
	__Field__000004_LONG_NAME.Fieldtypename = `string`

	__GongStructShape__000000_spec_object_ATTRIBUTE_VALUE_ENUMERATION.Name = `spec_object-ATTRIBUTE_VALUE_ENUMERATION`

	//gong:ident [ref_models.ATTRIBUTE_VALUE_ENUMERATION] comment added to overcome the problem with the comment map association
	__GongStructShape__000000_spec_object_ATTRIBUTE_VALUE_ENUMERATION.Identifier = `ref_models.ATTRIBUTE_VALUE_ENUMERATION`
	__GongStructShape__000000_spec_object_ATTRIBUTE_VALUE_ENUMERATION.ShowNbInstances = false
	__GongStructShape__000000_spec_object_ATTRIBUTE_VALUE_ENUMERATION.NbInstances = 0
	__GongStructShape__000000_spec_object_ATTRIBUTE_VALUE_ENUMERATION.Width = 387.999939
	__GongStructShape__000000_spec_object_ATTRIBUTE_VALUE_ENUMERATION.Height = 160.000000
	__GongStructShape__000000_spec_object_ATTRIBUTE_VALUE_ENUMERATION.IsSelected = false

	__GongStructShape__000001_spec_object_ATTRIBUTE_VALUE_XHTML.Name = `spec_object-ATTRIBUTE_VALUE_XHTML`

	//gong:ident [ref_models.ATTRIBUTE_VALUE_XHTML] comment added to overcome the problem with the comment map association
	__GongStructShape__000001_spec_object_ATTRIBUTE_VALUE_XHTML.Identifier = `ref_models.ATTRIBUTE_VALUE_XHTML`
	__GongStructShape__000001_spec_object_ATTRIBUTE_VALUE_XHTML.ShowNbInstances = false
	__GongStructShape__000001_spec_object_ATTRIBUTE_VALUE_XHTML.NbInstances = 0
	__GongStructShape__000001_spec_object_ATTRIBUTE_VALUE_XHTML.Width = 240.000000
	__GongStructShape__000001_spec_object_ATTRIBUTE_VALUE_XHTML.Height = 190.000000
	__GongStructShape__000001_spec_object_ATTRIBUTE_VALUE_XHTML.IsSelected = false

	__GongStructShape__000002_spec_object_A_ATTRIBUTE_DEFINITION_ENUMERATION_REF.Name = `spec_object-A_ATTRIBUTE_DEFINITION_ENUMERATION_REF`

	//gong:ident [ref_models.A_ATTRIBUTE_DEFINITION_ENUMERATION_REF] comment added to overcome the problem with the comment map association
	__GongStructShape__000002_spec_object_A_ATTRIBUTE_DEFINITION_ENUMERATION_REF.Identifier = `ref_models.A_ATTRIBUTE_DEFINITION_ENUMERATION_REF`
	__GongStructShape__000002_spec_object_A_ATTRIBUTE_DEFINITION_ENUMERATION_REF.ShowNbInstances = false
	__GongStructShape__000002_spec_object_A_ATTRIBUTE_DEFINITION_ENUMERATION_REF.NbInstances = 0
	__GongStructShape__000002_spec_object_A_ATTRIBUTE_DEFINITION_ENUMERATION_REF.Width = 444.000000
	__GongStructShape__000002_spec_object_A_ATTRIBUTE_DEFINITION_ENUMERATION_REF.Height = 63.000000
	__GongStructShape__000002_spec_object_A_ATTRIBUTE_DEFINITION_ENUMERATION_REF.IsSelected = false

	__GongStructShape__000003_spec_object_A_ATTRIBUTE_DEFINITION_XHTML_REF.Name = `spec_object-A_ATTRIBUTE_DEFINITION_XHTML_REF`

	//gong:ident [ref_models.A_ATTRIBUTE_DEFINITION_XHTML_REF] comment added to overcome the problem with the comment map association
	__GongStructShape__000003_spec_object_A_ATTRIBUTE_DEFINITION_XHTML_REF.Identifier = `ref_models.A_ATTRIBUTE_DEFINITION_XHTML_REF`
	__GongStructShape__000003_spec_object_A_ATTRIBUTE_DEFINITION_XHTML_REF.ShowNbInstances = false
	__GongStructShape__000003_spec_object_A_ATTRIBUTE_DEFINITION_XHTML_REF.NbInstances = 0
	__GongStructShape__000003_spec_object_A_ATTRIBUTE_DEFINITION_XHTML_REF.Width = 446.000000
	__GongStructShape__000003_spec_object_A_ATTRIBUTE_DEFINITION_XHTML_REF.Height = 78.000000
	__GongStructShape__000003_spec_object_A_ATTRIBUTE_DEFINITION_XHTML_REF.IsSelected = false

	__GongStructShape__000004_spec_object_A_ATTRIBUTE_VALUE_XHTML_1.Name = `spec_object-A_ATTRIBUTE_VALUE_XHTML_1`

	//gong:ident [ref_models.A_ATTRIBUTE_VALUE_XHTML_1] comment added to overcome the problem with the comment map association
	__GongStructShape__000004_spec_object_A_ATTRIBUTE_VALUE_XHTML_1.Identifier = `ref_models.A_ATTRIBUTE_VALUE_XHTML_1`
	__GongStructShape__000004_spec_object_A_ATTRIBUTE_VALUE_XHTML_1.ShowNbInstances = false
	__GongStructShape__000004_spec_object_A_ATTRIBUTE_VALUE_XHTML_1.NbInstances = 0
	__GongStructShape__000004_spec_object_A_ATTRIBUTE_VALUE_XHTML_1.Width = 328.000000
	__GongStructShape__000004_spec_object_A_ATTRIBUTE_VALUE_XHTML_1.Height = 686.999969
	__GongStructShape__000004_spec_object_A_ATTRIBUTE_VALUE_XHTML_1.IsSelected = false

	__GongStructShape__000005_spec_object_A_ENUM_VALUE_REF.Name = `spec_object-A_ENUM_VALUE_REF`

	//gong:ident [ref_models.A_ENUM_VALUE_REF] comment added to overcome the problem with the comment map association
	__GongStructShape__000005_spec_object_A_ENUM_VALUE_REF.Identifier = `ref_models.A_ENUM_VALUE_REF`
	__GongStructShape__000005_spec_object_A_ENUM_VALUE_REF.ShowNbInstances = false
	__GongStructShape__000005_spec_object_A_ENUM_VALUE_REF.NbInstances = 0
	__GongStructShape__000005_spec_object_A_ENUM_VALUE_REF.Width = 240.000000
	__GongStructShape__000005_spec_object_A_ENUM_VALUE_REF.Height = 63.000000
	__GongStructShape__000005_spec_object_A_ENUM_VALUE_REF.IsSelected = false

	__GongStructShape__000006_spec_object_A_SPEC_OBJECT_TYPE_REF.Name = `spec_object-A_SPEC_OBJECT_TYPE_REF`

	//gong:ident [ref_models.A_SPEC_OBJECT_TYPE_REF] comment added to overcome the problem with the comment map association
	__GongStructShape__000006_spec_object_A_SPEC_OBJECT_TYPE_REF.Identifier = `ref_models.A_SPEC_OBJECT_TYPE_REF`
	__GongStructShape__000006_spec_object_A_SPEC_OBJECT_TYPE_REF.ShowNbInstances = false
	__GongStructShape__000006_spec_object_A_SPEC_OBJECT_TYPE_REF.NbInstances = 0
	__GongStructShape__000006_spec_object_A_SPEC_OBJECT_TYPE_REF.Width = 240.000000
	__GongStructShape__000006_spec_object_A_SPEC_OBJECT_TYPE_REF.Height = 63.000000
	__GongStructShape__000006_spec_object_A_SPEC_OBJECT_TYPE_REF.IsSelected = false

	__GongStructShape__000007_spec_object_SPEC_OBJECT.Name = `spec_object-SPEC_OBJECT`

	//gong:ident [ref_models.SPEC_OBJECT] comment added to overcome the problem with the comment map association
	__GongStructShape__000007_spec_object_SPEC_OBJECT.Identifier = `ref_models.SPEC_OBJECT`
	__GongStructShape__000007_spec_object_SPEC_OBJECT.ShowNbInstances = false
	__GongStructShape__000007_spec_object_SPEC_OBJECT.NbInstances = 0
	__GongStructShape__000007_spec_object_SPEC_OBJECT.Width = 240.000000
	__GongStructShape__000007_spec_object_SPEC_OBJECT.Height = 108.000000
	__GongStructShape__000007_spec_object_SPEC_OBJECT.IsSelected = false

	__GongStructShape__000008_spec_object_XHTML_CONTENT.Name = `spec_object-XHTML_CONTENT`

	//gong:ident [ref_models.XHTML_CONTENT] comment added to overcome the problem with the comment map association
	__GongStructShape__000008_spec_object_XHTML_CONTENT.Identifier = `ref_models.XHTML_CONTENT`
	__GongStructShape__000008_spec_object_XHTML_CONTENT.ShowNbInstances = false
	__GongStructShape__000008_spec_object_XHTML_CONTENT.NbInstances = 0
	__GongStructShape__000008_spec_object_XHTML_CONTENT.Width = 240.000000
	__GongStructShape__000008_spec_object_XHTML_CONTENT.Height = 172.000000
	__GongStructShape__000008_spec_object_XHTML_CONTENT.IsSelected = false

	__Link__000000_ATTRIBUTE_VALUE_ENUMERATION.Name = `ATTRIBUTE_VALUE_ENUMERATION`

	//gong:ident [ref_models.A_ATTRIBUTE_VALUE_XHTML_1.ATTRIBUTE_VALUE_ENUMERATION] comment added to overcome the problem with the comment map association
	__Link__000000_ATTRIBUTE_VALUE_ENUMERATION.Identifier = `ref_models.A_ATTRIBUTE_VALUE_XHTML_1.ATTRIBUTE_VALUE_ENUMERATION`

	//gong:ident [ref_models.ATTRIBUTE_VALUE_ENUMERATION] comment added to overcome the problem with the comment map association
	__Link__000000_ATTRIBUTE_VALUE_ENUMERATION.Fieldtypename = `ref_models.ATTRIBUTE_VALUE_ENUMERATION`
	__Link__000000_ATTRIBUTE_VALUE_ENUMERATION.FieldOffsetX = -50.000000
	__Link__000000_ATTRIBUTE_VALUE_ENUMERATION.FieldOffsetY = -16.000000
	__Link__000000_ATTRIBUTE_VALUE_ENUMERATION.TargetMultiplicity = models.MANY
	__Link__000000_ATTRIBUTE_VALUE_ENUMERATION.TargetMultiplicityOffsetX = -50.000000
	__Link__000000_ATTRIBUTE_VALUE_ENUMERATION.TargetMultiplicityOffsetY = 16.000000
	__Link__000000_ATTRIBUTE_VALUE_ENUMERATION.SourceMultiplicity = models.MANY
	__Link__000000_ATTRIBUTE_VALUE_ENUMERATION.SourceMultiplicityOffsetX = 10.000000
	__Link__000000_ATTRIBUTE_VALUE_ENUMERATION.SourceMultiplicityOffsetY = -50.000000
	__Link__000000_ATTRIBUTE_VALUE_ENUMERATION.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_ATTRIBUTE_VALUE_ENUMERATION.StartRatio = 0.706211
	__Link__000000_ATTRIBUTE_VALUE_ENUMERATION.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_ATTRIBUTE_VALUE_ENUMERATION.EndRatio = 0.194316
	__Link__000000_ATTRIBUTE_VALUE_ENUMERATION.CornerOffsetRatio = 1.380000

	__Link__000001_ATTRIBUTE_VALUE_XHTML.Name = `ATTRIBUTE_VALUE_XHTML`

	//gong:ident [ref_models.A_ATTRIBUTE_VALUE_XHTML_1.ATTRIBUTE_VALUE_XHTML] comment added to overcome the problem with the comment map association
	__Link__000001_ATTRIBUTE_VALUE_XHTML.Identifier = `ref_models.A_ATTRIBUTE_VALUE_XHTML_1.ATTRIBUTE_VALUE_XHTML`

	//gong:ident [ref_models.ATTRIBUTE_VALUE_XHTML] comment added to overcome the problem with the comment map association
	__Link__000001_ATTRIBUTE_VALUE_XHTML.Fieldtypename = `ref_models.ATTRIBUTE_VALUE_XHTML`
	__Link__000001_ATTRIBUTE_VALUE_XHTML.FieldOffsetX = -50.000000
	__Link__000001_ATTRIBUTE_VALUE_XHTML.FieldOffsetY = -16.000000
	__Link__000001_ATTRIBUTE_VALUE_XHTML.TargetMultiplicity = models.MANY
	__Link__000001_ATTRIBUTE_VALUE_XHTML.TargetMultiplicityOffsetX = -50.000000
	__Link__000001_ATTRIBUTE_VALUE_XHTML.TargetMultiplicityOffsetY = 16.000000
	__Link__000001_ATTRIBUTE_VALUE_XHTML.SourceMultiplicity = models.MANY
	__Link__000001_ATTRIBUTE_VALUE_XHTML.SourceMultiplicityOffsetX = 10.000000
	__Link__000001_ATTRIBUTE_VALUE_XHTML.SourceMultiplicityOffsetY = -50.000000
	__Link__000001_ATTRIBUTE_VALUE_XHTML.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000001_ATTRIBUTE_VALUE_XHTML.StartRatio = 0.153081
	__Link__000001_ATTRIBUTE_VALUE_XHTML.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000001_ATTRIBUTE_VALUE_XHTML.EndRatio = 0.500000
	__Link__000001_ATTRIBUTE_VALUE_XHTML.CornerOffsetRatio = 1.380000

	__Link__000002_DEFINITION.Name = `DEFINITION`

	//gong:ident [ref_models.ATTRIBUTE_VALUE_ENUMERATION.DEFINITION] comment added to overcome the problem with the comment map association
	__Link__000002_DEFINITION.Identifier = `ref_models.ATTRIBUTE_VALUE_ENUMERATION.DEFINITION`

	//gong:ident [ref_models.A_ATTRIBUTE_DEFINITION_ENUMERATION_REF] comment added to overcome the problem with the comment map association
	__Link__000002_DEFINITION.Fieldtypename = `ref_models.A_ATTRIBUTE_DEFINITION_ENUMERATION_REF`
	__Link__000002_DEFINITION.FieldOffsetX = -50.000000
	__Link__000002_DEFINITION.FieldOffsetY = -16.000000
	__Link__000002_DEFINITION.TargetMultiplicity = models.ZERO_ONE
	__Link__000002_DEFINITION.TargetMultiplicityOffsetX = -50.000000
	__Link__000002_DEFINITION.TargetMultiplicityOffsetY = 16.000000
	__Link__000002_DEFINITION.SourceMultiplicity = models.MANY
	__Link__000002_DEFINITION.SourceMultiplicityOffsetX = 10.000000
	__Link__000002_DEFINITION.SourceMultiplicityOffsetY = -50.000000
	__Link__000002_DEFINITION.StartOrientation = models.ORIENTATION_VERTICAL
	__Link__000002_DEFINITION.StartRatio = 0.366838
	__Link__000002_DEFINITION.EndOrientation = models.ORIENTATION_VERTICAL
	__Link__000002_DEFINITION.EndRatio = 0.307057
	__Link__000002_DEFINITION.CornerOffsetRatio = 1.384375

	__Link__000003_DEFINITION.Name = `DEFINITION`

	//gong:ident [ref_models.ATTRIBUTE_VALUE_XHTML.DEFINITION] comment added to overcome the problem with the comment map association
	__Link__000003_DEFINITION.Identifier = `ref_models.ATTRIBUTE_VALUE_XHTML.DEFINITION`

	//gong:ident [ref_models.A_ATTRIBUTE_DEFINITION_XHTML_REF] comment added to overcome the problem with the comment map association
	__Link__000003_DEFINITION.Fieldtypename = `ref_models.A_ATTRIBUTE_DEFINITION_XHTML_REF`
	__Link__000003_DEFINITION.FieldOffsetX = -50.000000
	__Link__000003_DEFINITION.FieldOffsetY = -16.000000
	__Link__000003_DEFINITION.TargetMultiplicity = models.ZERO_ONE
	__Link__000003_DEFINITION.TargetMultiplicityOffsetX = -50.000000
	__Link__000003_DEFINITION.TargetMultiplicityOffsetY = 16.000000
	__Link__000003_DEFINITION.SourceMultiplicity = models.MANY
	__Link__000003_DEFINITION.SourceMultiplicityOffsetX = 10.000000
	__Link__000003_DEFINITION.SourceMultiplicityOffsetY = -50.000000
	__Link__000003_DEFINITION.StartOrientation = models.ORIENTATION_VERTICAL
	__Link__000003_DEFINITION.StartRatio = 0.297222
	__Link__000003_DEFINITION.EndOrientation = models.ORIENTATION_VERTICAL
	__Link__000003_DEFINITION.EndRatio = 0.168909
	__Link__000003_DEFINITION.CornerOffsetRatio = 1.232456

	__Link__000004_THE_ORIGINAL_VALUE.Name = `THE_ORIGINAL_VALUE`

	//gong:ident [ref_models.ATTRIBUTE_VALUE_XHTML.THE_ORIGINAL_VALUE] comment added to overcome the problem with the comment map association
	__Link__000004_THE_ORIGINAL_VALUE.Identifier = `ref_models.ATTRIBUTE_VALUE_XHTML.THE_ORIGINAL_VALUE`

	//gong:ident [ref_models.XHTML_CONTENT] comment added to overcome the problem with the comment map association
	__Link__000004_THE_ORIGINAL_VALUE.Fieldtypename = `ref_models.XHTML_CONTENT`
	__Link__000004_THE_ORIGINAL_VALUE.FieldOffsetX = -50.000000
	__Link__000004_THE_ORIGINAL_VALUE.FieldOffsetY = -16.000000
	__Link__000004_THE_ORIGINAL_VALUE.TargetMultiplicity = models.ZERO_ONE
	__Link__000004_THE_ORIGINAL_VALUE.TargetMultiplicityOffsetX = -50.000000
	__Link__000004_THE_ORIGINAL_VALUE.TargetMultiplicityOffsetY = 16.000000
	__Link__000004_THE_ORIGINAL_VALUE.SourceMultiplicity = models.MANY
	__Link__000004_THE_ORIGINAL_VALUE.SourceMultiplicityOffsetX = 10.000000
	__Link__000004_THE_ORIGINAL_VALUE.SourceMultiplicityOffsetY = -50.000000
	__Link__000004_THE_ORIGINAL_VALUE.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000004_THE_ORIGINAL_VALUE.StartRatio = 0.764035
	__Link__000004_THE_ORIGINAL_VALUE.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000004_THE_ORIGINAL_VALUE.EndRatio = 0.820736
	__Link__000004_THE_ORIGINAL_VALUE.CornerOffsetRatio = 1.380000

	__Link__000005_THE_VALUE.Name = `THE_VALUE`

	//gong:ident [ref_models.ATTRIBUTE_VALUE_XHTML.THE_VALUE] comment added to overcome the problem with the comment map association
	__Link__000005_THE_VALUE.Identifier = `ref_models.ATTRIBUTE_VALUE_XHTML.THE_VALUE`

	//gong:ident [ref_models.XHTML_CONTENT] comment added to overcome the problem with the comment map association
	__Link__000005_THE_VALUE.Fieldtypename = `ref_models.XHTML_CONTENT`
	__Link__000005_THE_VALUE.FieldOffsetX = -50.000000
	__Link__000005_THE_VALUE.FieldOffsetY = -16.000000
	__Link__000005_THE_VALUE.TargetMultiplicity = models.ZERO_ONE
	__Link__000005_THE_VALUE.TargetMultiplicityOffsetX = -50.000000
	__Link__000005_THE_VALUE.TargetMultiplicityOffsetY = 16.000000
	__Link__000005_THE_VALUE.SourceMultiplicity = models.MANY
	__Link__000005_THE_VALUE.SourceMultiplicityOffsetX = 10.000000
	__Link__000005_THE_VALUE.SourceMultiplicityOffsetY = -50.000000
	__Link__000005_THE_VALUE.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000005_THE_VALUE.StartRatio = 0.311403
	__Link__000005_THE_VALUE.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000005_THE_VALUE.EndRatio = 0.367248
	__Link__000005_THE_VALUE.CornerOffsetRatio = 1.380000

	__Link__000006_TYPE.Name = `TYPE`

	//gong:ident [ref_models.SPEC_OBJECT.TYPE] comment added to overcome the problem with the comment map association
	__Link__000006_TYPE.Identifier = `ref_models.SPEC_OBJECT.TYPE`

	//gong:ident [ref_models.A_SPEC_OBJECT_TYPE_REF] comment added to overcome the problem with the comment map association
	__Link__000006_TYPE.Fieldtypename = `ref_models.A_SPEC_OBJECT_TYPE_REF`
	__Link__000006_TYPE.FieldOffsetX = -50.000000
	__Link__000006_TYPE.FieldOffsetY = -16.000000
	__Link__000006_TYPE.TargetMultiplicity = models.ZERO_ONE
	__Link__000006_TYPE.TargetMultiplicityOffsetX = -50.000000
	__Link__000006_TYPE.TargetMultiplicityOffsetY = 16.000000
	__Link__000006_TYPE.SourceMultiplicity = models.MANY
	__Link__000006_TYPE.SourceMultiplicityOffsetX = 10.000000
	__Link__000006_TYPE.SourceMultiplicityOffsetY = -50.000000
	__Link__000006_TYPE.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000006_TYPE.StartRatio = 0.500000
	__Link__000006_TYPE.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000006_TYPE.EndRatio = 0.500000
	__Link__000006_TYPE.CornerOffsetRatio = 1.380000

	__Link__000007_VALUES.Name = `VALUES`

	//gong:ident [ref_models.SPEC_OBJECT.VALUES] comment added to overcome the problem with the comment map association
	__Link__000007_VALUES.Identifier = `ref_models.SPEC_OBJECT.VALUES`

	//gong:ident [ref_models.A_ATTRIBUTE_VALUE_XHTML_1] comment added to overcome the problem with the comment map association
	__Link__000007_VALUES.Fieldtypename = `ref_models.A_ATTRIBUTE_VALUE_XHTML_1`
	__Link__000007_VALUES.FieldOffsetX = -50.000000
	__Link__000007_VALUES.FieldOffsetY = -16.000000
	__Link__000007_VALUES.TargetMultiplicity = models.ZERO_ONE
	__Link__000007_VALUES.TargetMultiplicityOffsetX = -50.000000
	__Link__000007_VALUES.TargetMultiplicityOffsetY = 16.000000
	__Link__000007_VALUES.SourceMultiplicity = models.MANY
	__Link__000007_VALUES.SourceMultiplicityOffsetX = 10.000000
	__Link__000007_VALUES.SourceMultiplicityOffsetY = -50.000000
	__Link__000007_VALUES.StartOrientation = models.ORIENTATION_VERTICAL
	__Link__000007_VALUES.StartRatio = 0.497222
	__Link__000007_VALUES.EndOrientation = models.ORIENTATION_VERTICAL
	__Link__000007_VALUES.EndRatio = 0.372222
	__Link__000007_VALUES.CornerOffsetRatio = 1.233025

	__Link__000008_VALUES.Name = `VALUES`

	//gong:ident [ref_models.ATTRIBUTE_VALUE_ENUMERATION.VALUES] comment added to overcome the problem with the comment map association
	__Link__000008_VALUES.Identifier = `ref_models.ATTRIBUTE_VALUE_ENUMERATION.VALUES`

	//gong:ident [ref_models.A_ENUM_VALUE_REF] comment added to overcome the problem with the comment map association
	__Link__000008_VALUES.Fieldtypename = `ref_models.A_ENUM_VALUE_REF`
	__Link__000008_VALUES.FieldOffsetX = -50.000000
	__Link__000008_VALUES.FieldOffsetY = -16.000000
	__Link__000008_VALUES.TargetMultiplicity = models.ZERO_ONE
	__Link__000008_VALUES.TargetMultiplicityOffsetX = -50.000000
	__Link__000008_VALUES.TargetMultiplicityOffsetY = 16.000000
	__Link__000008_VALUES.SourceMultiplicity = models.MANY
	__Link__000008_VALUES.SourceMultiplicityOffsetX = 10.000000
	__Link__000008_VALUES.SourceMultiplicityOffsetY = -50.000000
	__Link__000008_VALUES.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000008_VALUES.StartRatio = 0.500000
	__Link__000008_VALUES.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000008_VALUES.EndRatio = 0.500000
	__Link__000008_VALUES.CornerOffsetRatio = 1.380000

	__Position__000000_Pos_spec_object_ATTRIBUTE_VALUE_ENUMERATION.X = 781.000000
	__Position__000000_Pos_spec_object_ATTRIBUTE_VALUE_ENUMERATION.Y = 740.999969
	__Position__000000_Pos_spec_object_ATTRIBUTE_VALUE_ENUMERATION.Name = `Pos-spec_object-ATTRIBUTE_VALUE_ENUMERATION`

	__Position__000001_Pos_spec_object_ATTRIBUTE_VALUE_XHTML.X = 808.000000
	__Position__000001_Pos_spec_object_ATTRIBUTE_VALUE_XHTML.Y = 303.999992
	__Position__000001_Pos_spec_object_ATTRIBUTE_VALUE_XHTML.Name = `Pos-spec_object-ATTRIBUTE_VALUE_XHTML`

	__Position__000002_Pos_spec_object_A_ATTRIBUTE_DEFINITION_ENUMERATION_REF.X = 781.000000
	__Position__000002_Pos_spec_object_A_ATTRIBUTE_DEFINITION_ENUMERATION_REF.Y = 1006.333344
	__Position__000002_Pos_spec_object_A_ATTRIBUTE_DEFINITION_ENUMERATION_REF.Name = `Pos-spec_object-A_ATTRIBUTE_DEFINITION_ENUMERATION_REF`

	__Position__000003_Pos_spec_object_A_ATTRIBUTE_DEFINITION_XHTML_REF.X = 804.000000
	__Position__000003_Pos_spec_object_A_ATTRIBUTE_DEFINITION_XHTML_REF.Y = 605.999962
	__Position__000003_Pos_spec_object_A_ATTRIBUTE_DEFINITION_XHTML_REF.Name = `Pos-spec_object-A_ATTRIBUTE_DEFINITION_XHTML_REF`

	__Position__000004_Pos_spec_object_A_ATTRIBUTE_VALUE_XHTML_1.X = 100.000000
	__Position__000004_Pos_spec_object_A_ATTRIBUTE_VALUE_XHTML_1.Y = 291.999992
	__Position__000004_Pos_spec_object_A_ATTRIBUTE_VALUE_XHTML_1.Name = `Pos-spec_object-A_ATTRIBUTE_VALUE_XHTML_1`

	__Position__000005_Pos_spec_object_A_ENUM_VALUE_REF.X = 1436.999939
	__Position__000005_Pos_spec_object_A_ENUM_VALUE_REF.Y = 778.999969
	__Position__000005_Pos_spec_object_A_ENUM_VALUE_REF.Name = `Pos-spec_object-A_ENUM_VALUE_REF`

	__Position__000006_Pos_spec_object_A_SPEC_OBJECT_TYPE_REF.X = 566.000000
	__Position__000006_Pos_spec_object_A_SPEC_OBJECT_TYPE_REF.Y = 49.000000
	__Position__000006_Pos_spec_object_A_SPEC_OBJECT_TYPE_REF.Name = `Pos-spec_object-A_SPEC_OBJECT_TYPE_REF`

	__Position__000007_Pos_spec_object_SPEC_OBJECT.X = 103.000000
	__Position__000007_Pos_spec_object_SPEC_OBJECT.Y = 37.000000
	__Position__000007_Pos_spec_object_SPEC_OBJECT.Name = `Pos-spec_object-SPEC_OBJECT`

	__Position__000008_Pos_spec_object_XHTML_CONTENT.X = 1353.999939
	__Position__000008_Pos_spec_object_XHTML_CONTENT.Y = 305.999992
	__Position__000008_Pos_spec_object_XHTML_CONTENT.Name = `Pos-spec_object-XHTML_CONTENT`

	__Vertice__000000_Verticle_in_class_diagram_spec_object_in_middle_between_spec_object_ATTRIBUTE_VALUE_ENUMERATION_and_spec_object_A_ATTRIBUTE_DEFINITION_ENUMERATION_REF.X = 1378.499908
	__Vertice__000000_Verticle_in_class_diagram_spec_object_in_middle_between_spec_object_ATTRIBUTE_VALUE_ENUMERATION_and_spec_object_A_ATTRIBUTE_DEFINITION_ENUMERATION_REF.Y = 952.166656
	__Vertice__000000_Verticle_in_class_diagram_spec_object_in_middle_between_spec_object_ATTRIBUTE_VALUE_ENUMERATION_and_spec_object_A_ATTRIBUTE_DEFINITION_ENUMERATION_REF.Name = `Verticle in class diagram spec_object in middle between spec_object-ATTRIBUTE_VALUE_ENUMERATION and spec_object-A_ATTRIBUTE_DEFINITION_ENUMERATION_REF`

	__Vertice__000001_Verticle_in_class_diagram_spec_object_in_middle_between_spec_object_ATTRIBUTE_VALUE_ENUMERATION_and_spec_object_A_ENUM_VALUE_REF.X = 1652.999878
	__Vertice__000001_Verticle_in_class_diagram_spec_object_in_middle_between_spec_object_ATTRIBUTE_VALUE_ENUMERATION_and_spec_object_A_ENUM_VALUE_REF.Y = 815.999969
	__Vertice__000001_Verticle_in_class_diagram_spec_object_in_middle_between_spec_object_ATTRIBUTE_VALUE_ENUMERATION_and_spec_object_A_ENUM_VALUE_REF.Name = `Verticle in class diagram spec_object in middle between spec_object-ATTRIBUTE_VALUE_ENUMERATION and spec_object-A_ENUM_VALUE_REF`

	__Vertice__000002_Verticle_in_class_diagram_spec_object_in_middle_between_spec_object_ATTRIBUTE_VALUE_XHTML_and_spec_object_A_ATTRIBUTE_DEFINITION_XHTML_REF.X = 1166.000000
	__Vertice__000002_Verticle_in_class_diagram_spec_object_in_middle_between_spec_object_ATTRIBUTE_VALUE_XHTML_and_spec_object_A_ATTRIBUTE_DEFINITION_XHTML_REF.Y = 549.999977
	__Vertice__000002_Verticle_in_class_diagram_spec_object_in_middle_between_spec_object_ATTRIBUTE_VALUE_XHTML_and_spec_object_A_ATTRIBUTE_DEFINITION_XHTML_REF.Name = `Verticle in class diagram spec_object in middle between spec_object-ATTRIBUTE_VALUE_XHTML and spec_object-A_ATTRIBUTE_DEFINITION_XHTML_REF`

	__Vertice__000003_Verticle_in_class_diagram_spec_object_in_middle_between_spec_object_ATTRIBUTE_VALUE_XHTML_and_spec_object_XHTML_CONTENT.X = 1439.999969
	__Vertice__000003_Verticle_in_class_diagram_spec_object_in_middle_between_spec_object_ATTRIBUTE_VALUE_XHTML_and_spec_object_XHTML_CONTENT.Y = 340.499992
	__Vertice__000003_Verticle_in_class_diagram_spec_object_in_middle_between_spec_object_ATTRIBUTE_VALUE_XHTML_and_spec_object_XHTML_CONTENT.Name = `Verticle in class diagram spec_object in middle between spec_object-ATTRIBUTE_VALUE_XHTML and spec_object-XHTML_CONTENT`

	__Vertice__000004_Verticle_in_class_diagram_spec_object_in_middle_between_spec_object_ATTRIBUTE_VALUE_XHTML_and_spec_object_XHTML_CONTENT.X = 1439.999969
	__Vertice__000004_Verticle_in_class_diagram_spec_object_in_middle_between_spec_object_ATTRIBUTE_VALUE_XHTML_and_spec_object_XHTML_CONTENT.Y = 340.499992
	__Vertice__000004_Verticle_in_class_diagram_spec_object_in_middle_between_spec_object_ATTRIBUTE_VALUE_XHTML_and_spec_object_XHTML_CONTENT.Name = `Verticle in class diagram spec_object in middle between spec_object-ATTRIBUTE_VALUE_XHTML and spec_object-XHTML_CONTENT`

	__Vertice__000005_Verticle_in_class_diagram_spec_object_in_middle_between_spec_object_A_ATTRIBUTE_VALUE_XHTML_1_and_spec_object_ATTRIBUTE_VALUE_ENUMERATION.X = 932.500000
	__Vertice__000005_Verticle_in_class_diagram_spec_object_in_middle_between_spec_object_A_ATTRIBUTE_VALUE_XHTML_1_and_spec_object_ATTRIBUTE_VALUE_ENUMERATION.Y = 860.499966
	__Vertice__000005_Verticle_in_class_diagram_spec_object_in_middle_between_spec_object_A_ATTRIBUTE_VALUE_XHTML_1_and_spec_object_ATTRIBUTE_VALUE_ENUMERATION.Name = `Verticle in class diagram spec_object in middle between spec_object-A_ATTRIBUTE_VALUE_XHTML_1 and spec_object-ATTRIBUTE_VALUE_ENUMERATION`

	__Vertice__000006_Verticle_in_class_diagram_spec_object_in_middle_between_spec_object_A_ATTRIBUTE_VALUE_XHTML_1_and_spec_object_ATTRIBUTE_VALUE_XHTML.X = 858.500000
	__Vertice__000006_Verticle_in_class_diagram_spec_object_in_middle_between_spec_object_A_ATTRIBUTE_VALUE_XHTML_1_and_spec_object_ATTRIBUTE_VALUE_XHTML.Y = 521.499977
	__Vertice__000006_Verticle_in_class_diagram_spec_object_in_middle_between_spec_object_A_ATTRIBUTE_VALUE_XHTML_1_and_spec_object_ATTRIBUTE_VALUE_XHTML.Name = `Verticle in class diagram spec_object in middle between spec_object-A_ATTRIBUTE_VALUE_XHTML_1 and spec_object-ATTRIBUTE_VALUE_XHTML`

	__Vertice__000007_Verticle_in_class_diagram_spec_object_in_middle_between_spec_object_SPEC_OBJECT_and_spec_object_A_ATTRIBUTE_VALUE_XHTML_1.X = 465.000000
	__Vertice__000007_Verticle_in_class_diagram_spec_object_in_middle_between_spec_object_SPEC_OBJECT_and_spec_object_A_ATTRIBUTE_VALUE_XHTML_1.Y = 189.999996
	__Vertice__000007_Verticle_in_class_diagram_spec_object_in_middle_between_spec_object_SPEC_OBJECT_and_spec_object_A_ATTRIBUTE_VALUE_XHTML_1.Name = `Verticle in class diagram spec_object in middle between spec_object-SPEC_OBJECT and spec_object-A_ATTRIBUTE_VALUE_XHTML_1`

	__Vertice__000008_Verticle_in_class_diagram_spec_object_in_middle_between_spec_object_SPEC_OBJECT_and_spec_object_A_SPEC_OBJECT_TYPE_REF.X = 694.500000
	__Vertice__000008_Verticle_in_class_diagram_spec_object_in_middle_between_spec_object_SPEC_OBJECT_and_spec_object_A_SPEC_OBJECT_TYPE_REF.Y = 97.000000
	__Vertice__000008_Verticle_in_class_diagram_spec_object_in_middle_between_spec_object_SPEC_OBJECT_and_spec_object_A_SPEC_OBJECT_TYPE_REF.Name = `Verticle in class diagram spec_object in middle between spec_object-SPEC_OBJECT and spec_object-A_SPEC_OBJECT_TYPE_REF`

	// Setup of pointers
	__Classdiagram__000000_spec_object.GongStructShapes = append(__Classdiagram__000000_spec_object.GongStructShapes, __GongStructShape__000007_spec_object_SPEC_OBJECT)
	__Classdiagram__000000_spec_object.GongStructShapes = append(__Classdiagram__000000_spec_object.GongStructShapes, __GongStructShape__000006_spec_object_A_SPEC_OBJECT_TYPE_REF)
	__Classdiagram__000000_spec_object.GongStructShapes = append(__Classdiagram__000000_spec_object.GongStructShapes, __GongStructShape__000004_spec_object_A_ATTRIBUTE_VALUE_XHTML_1)
	__Classdiagram__000000_spec_object.GongStructShapes = append(__Classdiagram__000000_spec_object.GongStructShapes, __GongStructShape__000001_spec_object_ATTRIBUTE_VALUE_XHTML)
	__Classdiagram__000000_spec_object.GongStructShapes = append(__Classdiagram__000000_spec_object.GongStructShapes, __GongStructShape__000008_spec_object_XHTML_CONTENT)
	__Classdiagram__000000_spec_object.GongStructShapes = append(__Classdiagram__000000_spec_object.GongStructShapes, __GongStructShape__000003_spec_object_A_ATTRIBUTE_DEFINITION_XHTML_REF)
	__Classdiagram__000000_spec_object.GongStructShapes = append(__Classdiagram__000000_spec_object.GongStructShapes, __GongStructShape__000000_spec_object_ATTRIBUTE_VALUE_ENUMERATION)
	__Classdiagram__000000_spec_object.GongStructShapes = append(__Classdiagram__000000_spec_object.GongStructShapes, __GongStructShape__000005_spec_object_A_ENUM_VALUE_REF)
	__Classdiagram__000000_spec_object.GongStructShapes = append(__Classdiagram__000000_spec_object.GongStructShapes, __GongStructShape__000002_spec_object_A_ATTRIBUTE_DEFINITION_ENUMERATION_REF)
	__GongStructShape__000000_spec_object_ATTRIBUTE_VALUE_ENUMERATION.Position = __Position__000000_Pos_spec_object_ATTRIBUTE_VALUE_ENUMERATION
	__GongStructShape__000000_spec_object_ATTRIBUTE_VALUE_ENUMERATION.Links = append(__GongStructShape__000000_spec_object_ATTRIBUTE_VALUE_ENUMERATION.Links, __Link__000008_VALUES)
	__GongStructShape__000000_spec_object_ATTRIBUTE_VALUE_ENUMERATION.Links = append(__GongStructShape__000000_spec_object_ATTRIBUTE_VALUE_ENUMERATION.Links, __Link__000002_DEFINITION)
	__GongStructShape__000001_spec_object_ATTRIBUTE_VALUE_XHTML.Position = __Position__000001_Pos_spec_object_ATTRIBUTE_VALUE_XHTML
	__GongStructShape__000001_spec_object_ATTRIBUTE_VALUE_XHTML.Links = append(__GongStructShape__000001_spec_object_ATTRIBUTE_VALUE_XHTML.Links, __Link__000005_THE_VALUE)
	__GongStructShape__000001_spec_object_ATTRIBUTE_VALUE_XHTML.Links = append(__GongStructShape__000001_spec_object_ATTRIBUTE_VALUE_XHTML.Links, __Link__000004_THE_ORIGINAL_VALUE)
	__GongStructShape__000001_spec_object_ATTRIBUTE_VALUE_XHTML.Links = append(__GongStructShape__000001_spec_object_ATTRIBUTE_VALUE_XHTML.Links, __Link__000003_DEFINITION)
	__GongStructShape__000002_spec_object_A_ATTRIBUTE_DEFINITION_ENUMERATION_REF.Position = __Position__000002_Pos_spec_object_A_ATTRIBUTE_DEFINITION_ENUMERATION_REF
	__GongStructShape__000003_spec_object_A_ATTRIBUTE_DEFINITION_XHTML_REF.Position = __Position__000003_Pos_spec_object_A_ATTRIBUTE_DEFINITION_XHTML_REF
	__GongStructShape__000003_spec_object_A_ATTRIBUTE_DEFINITION_XHTML_REF.Fields = append(__GongStructShape__000003_spec_object_A_ATTRIBUTE_DEFINITION_XHTML_REF.Fields, __Field__000000_ATTRIBUTE_DEFINITION_XHTML_REF)
	__GongStructShape__000004_spec_object_A_ATTRIBUTE_VALUE_XHTML_1.Position = __Position__000004_Pos_spec_object_A_ATTRIBUTE_VALUE_XHTML_1
	__GongStructShape__000004_spec_object_A_ATTRIBUTE_VALUE_XHTML_1.Links = append(__GongStructShape__000004_spec_object_A_ATTRIBUTE_VALUE_XHTML_1.Links, __Link__000001_ATTRIBUTE_VALUE_XHTML)
	__GongStructShape__000004_spec_object_A_ATTRIBUTE_VALUE_XHTML_1.Links = append(__GongStructShape__000004_spec_object_A_ATTRIBUTE_VALUE_XHTML_1.Links, __Link__000000_ATTRIBUTE_VALUE_ENUMERATION)
	__GongStructShape__000005_spec_object_A_ENUM_VALUE_REF.Position = __Position__000005_Pos_spec_object_A_ENUM_VALUE_REF
	__GongStructShape__000006_spec_object_A_SPEC_OBJECT_TYPE_REF.Position = __Position__000006_Pos_spec_object_A_SPEC_OBJECT_TYPE_REF
	__GongStructShape__000007_spec_object_SPEC_OBJECT.Position = __Position__000007_Pos_spec_object_SPEC_OBJECT
	__GongStructShape__000007_spec_object_SPEC_OBJECT.Fields = append(__GongStructShape__000007_spec_object_SPEC_OBJECT.Fields, __Field__000002_IDENTIFIER)
	__GongStructShape__000007_spec_object_SPEC_OBJECT.Fields = append(__GongStructShape__000007_spec_object_SPEC_OBJECT.Fields, __Field__000003_LAST_CHANGE)
	__GongStructShape__000007_spec_object_SPEC_OBJECT.Fields = append(__GongStructShape__000007_spec_object_SPEC_OBJECT.Fields, __Field__000004_LONG_NAME)
	__GongStructShape__000007_spec_object_SPEC_OBJECT.Links = append(__GongStructShape__000007_spec_object_SPEC_OBJECT.Links, __Link__000006_TYPE)
	__GongStructShape__000007_spec_object_SPEC_OBJECT.Links = append(__GongStructShape__000007_spec_object_SPEC_OBJECT.Links, __Link__000007_VALUES)
	__GongStructShape__000008_spec_object_XHTML_CONTENT.Position = __Position__000008_Pos_spec_object_XHTML_CONTENT
	__GongStructShape__000008_spec_object_XHTML_CONTENT.Fields = append(__GongStructShape__000008_spec_object_XHTML_CONTENT.Fields, __Field__000001_EnclosedText)
	__Link__000000_ATTRIBUTE_VALUE_ENUMERATION.Middlevertice = __Vertice__000005_Verticle_in_class_diagram_spec_object_in_middle_between_spec_object_A_ATTRIBUTE_VALUE_XHTML_1_and_spec_object_ATTRIBUTE_VALUE_ENUMERATION
	__Link__000001_ATTRIBUTE_VALUE_XHTML.Middlevertice = __Vertice__000006_Verticle_in_class_diagram_spec_object_in_middle_between_spec_object_A_ATTRIBUTE_VALUE_XHTML_1_and_spec_object_ATTRIBUTE_VALUE_XHTML
	__Link__000002_DEFINITION.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_spec_object_in_middle_between_spec_object_ATTRIBUTE_VALUE_ENUMERATION_and_spec_object_A_ATTRIBUTE_DEFINITION_ENUMERATION_REF
	__Link__000003_DEFINITION.Middlevertice = __Vertice__000002_Verticle_in_class_diagram_spec_object_in_middle_between_spec_object_ATTRIBUTE_VALUE_XHTML_and_spec_object_A_ATTRIBUTE_DEFINITION_XHTML_REF
	__Link__000004_THE_ORIGINAL_VALUE.Middlevertice = __Vertice__000004_Verticle_in_class_diagram_spec_object_in_middle_between_spec_object_ATTRIBUTE_VALUE_XHTML_and_spec_object_XHTML_CONTENT
	__Link__000005_THE_VALUE.Middlevertice = __Vertice__000003_Verticle_in_class_diagram_spec_object_in_middle_between_spec_object_ATTRIBUTE_VALUE_XHTML_and_spec_object_XHTML_CONTENT
	__Link__000006_TYPE.Middlevertice = __Vertice__000008_Verticle_in_class_diagram_spec_object_in_middle_between_spec_object_SPEC_OBJECT_and_spec_object_A_SPEC_OBJECT_TYPE_REF
	__Link__000007_VALUES.Middlevertice = __Vertice__000007_Verticle_in_class_diagram_spec_object_in_middle_between_spec_object_SPEC_OBJECT_and_spec_object_A_ATTRIBUTE_VALUE_XHTML_1
	__Link__000008_VALUES.Middlevertice = __Vertice__000001_Verticle_in_class_diagram_spec_object_in_middle_between_spec_object_ATTRIBUTE_VALUE_ENUMERATION_and_spec_object_A_ENUM_VALUE_REF
}
