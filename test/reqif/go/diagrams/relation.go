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

	__Classdiagram__000000_relation := (&models.Classdiagram{Name: `relation`}).Stage(stage)

	__Field__000000_DESC := (&models.Field{Name: `DESC`}).Stage(stage)
	__Field__000001_IDENTIFIER := (&models.Field{Name: `IDENTIFIER`}).Stage(stage)
	__Field__000002_LAST_CHANGE := (&models.Field{Name: `LAST_CHANGE`}).Stage(stage)
	__Field__000003_LONG_NAME := (&models.Field{Name: `LONG_NAME`}).Stage(stage)
	__Field__000004_Name := (&models.Field{Name: `Name`}).Stage(stage)

	__GongStructShape__000000_relation_A_ATTRIBUTE_VALUE_XHTML_1 := (&models.GongStructShape{Name: `relation-A_ATTRIBUTE_VALUE_XHTML_1`}).Stage(stage)
	__GongStructShape__000001_relation_A_SOURCE := (&models.GongStructShape{Name: `relation-A_SOURCE`}).Stage(stage)
	__GongStructShape__000002_relation_A_SPEC_RELATIONS := (&models.GongStructShape{Name: `relation-A_SPEC_RELATIONS`}).Stage(stage)
	__GongStructShape__000003_relation_SPECIFICATION := (&models.GongStructShape{Name: `relation-SPECIFICATION`}).Stage(stage)
	__GongStructShape__000004_relation_SPEC_RELATION := (&models.GongStructShape{Name: `relation-SPEC_RELATION`}).Stage(stage)
	__GongStructShape__000005_relation_SPEC_RELATION_TYPE := (&models.GongStructShape{Name: `relation-SPEC_RELATION_TYPE`}).Stage(stage)

	__Link__000000_SOURCE := (&models.Link{Name: `SOURCE`}).Stage(stage)
	__Link__000001_SPEC_RELATION := (&models.Link{Name: `SPEC_RELATION`}).Stage(stage)
	__Link__000002_VALUES := (&models.Link{Name: `VALUES`}).Stage(stage)

	__Position__000000_Pos_relation_A_ATTRIBUTE_VALUE_XHTML_1 := (&models.Position{Name: `Pos-relation-A_ATTRIBUTE_VALUE_XHTML_1`}).Stage(stage)
	__Position__000001_Pos_relation_A_SOURCE := (&models.Position{Name: `Pos-relation-A_SOURCE`}).Stage(stage)
	__Position__000002_Pos_relation_A_SPEC_RELATIONS := (&models.Position{Name: `Pos-relation-A_SPEC_RELATIONS`}).Stage(stage)
	__Position__000003_Pos_relation_SPECIFICATION := (&models.Position{Name: `Pos-relation-SPECIFICATION`}).Stage(stage)
	__Position__000004_Pos_relation_SPEC_RELATION := (&models.Position{Name: `Pos-relation-SPEC_RELATION`}).Stage(stage)
	__Position__000005_Pos_relation_SPEC_RELATION_TYPE := (&models.Position{Name: `Pos-relation-SPEC_RELATION_TYPE`}).Stage(stage)

	__Vertice__000000_Verticle_in_class_diagram_relation_in_middle_between_relation_A_SPEC_RELATIONS_and_relation_SPEC_RELATION := (&models.Vertice{Name: `Verticle in class diagram relation in middle between relation-A_SPEC_RELATIONS and relation-SPEC_RELATION`}).Stage(stage)
	__Vertice__000001_Verticle_in_class_diagram_relation_in_middle_between_relation_SPEC_RELATION_and_relation_A_ATTRIBUTE_VALUE_XHTML_1 := (&models.Vertice{Name: `Verticle in class diagram relation in middle between relation-SPEC_RELATION and relation-A_ATTRIBUTE_VALUE_XHTML_1`}).Stage(stage)
	__Vertice__000002_Verticle_in_class_diagram_relation_in_middle_between_relation_SPEC_RELATION_and_relation_A_SOURCE := (&models.Vertice{Name: `Verticle in class diagram relation in middle between relation-SPEC_RELATION and relation-A_SOURCE`}).Stage(stage)

	// Setup of values

	__Classdiagram__000000_relation.Name = `relation`
	__Classdiagram__000000_relation.IsInDrawMode = false

	__Field__000000_DESC.Name = `DESC`

	//gong:ident [ref_models.SPEC_RELATION.DESC] comment added to overcome the problem with the comment map association
	__Field__000000_DESC.Identifier = `ref_models.SPEC_RELATION.DESC`
	__Field__000000_DESC.FieldTypeAsString = ``
	__Field__000000_DESC.Structname = `SPEC_RELATION`
	__Field__000000_DESC.Fieldtypename = `string`

	__Field__000001_IDENTIFIER.Name = `IDENTIFIER`

	//gong:ident [ref_models.SPEC_RELATION.IDENTIFIER] comment added to overcome the problem with the comment map association
	__Field__000001_IDENTIFIER.Identifier = `ref_models.SPEC_RELATION.IDENTIFIER`
	__Field__000001_IDENTIFIER.FieldTypeAsString = ``
	__Field__000001_IDENTIFIER.Structname = `SPEC_RELATION`
	__Field__000001_IDENTIFIER.Fieldtypename = `string`

	__Field__000002_LAST_CHANGE.Name = `LAST_CHANGE`

	//gong:ident [ref_models.SPEC_RELATION.LAST_CHANGE] comment added to overcome the problem with the comment map association
	__Field__000002_LAST_CHANGE.Identifier = `ref_models.SPEC_RELATION.LAST_CHANGE`
	__Field__000002_LAST_CHANGE.FieldTypeAsString = ``
	__Field__000002_LAST_CHANGE.Structname = `SPEC_RELATION`
	__Field__000002_LAST_CHANGE.Fieldtypename = `string`

	__Field__000003_LONG_NAME.Name = `LONG_NAME`

	//gong:ident [ref_models.SPEC_RELATION.LONG_NAME] comment added to overcome the problem with the comment map association
	__Field__000003_LONG_NAME.Identifier = `ref_models.SPEC_RELATION.LONG_NAME`
	__Field__000003_LONG_NAME.FieldTypeAsString = ``
	__Field__000003_LONG_NAME.Structname = `SPEC_RELATION`
	__Field__000003_LONG_NAME.Fieldtypename = `string`

	__Field__000004_Name.Name = `Name`

	//gong:ident [ref_models.SPEC_RELATION.Name] comment added to overcome the problem with the comment map association
	__Field__000004_Name.Identifier = `ref_models.SPEC_RELATION.Name`
	__Field__000004_Name.FieldTypeAsString = ``
	__Field__000004_Name.Structname = `SPEC_RELATION`
	__Field__000004_Name.Fieldtypename = `string`

	__GongStructShape__000000_relation_A_ATTRIBUTE_VALUE_XHTML_1.Name = `relation-A_ATTRIBUTE_VALUE_XHTML_1`

	//gong:ident [ref_models.A_ATTRIBUTE_VALUE_XHTML_1] comment added to overcome the problem with the comment map association
	__GongStructShape__000000_relation_A_ATTRIBUTE_VALUE_XHTML_1.Identifier = `ref_models.A_ATTRIBUTE_VALUE_XHTML_1`
	__GongStructShape__000000_relation_A_ATTRIBUTE_VALUE_XHTML_1.ShowNbInstances = false
	__GongStructShape__000000_relation_A_ATTRIBUTE_VALUE_XHTML_1.NbInstances = 0
	__GongStructShape__000000_relation_A_ATTRIBUTE_VALUE_XHTML_1.Width = 390.000000
	__GongStructShape__000000_relation_A_ATTRIBUTE_VALUE_XHTML_1.Height = 63.000000
	__GongStructShape__000000_relation_A_ATTRIBUTE_VALUE_XHTML_1.IsSelected = false

	__GongStructShape__000001_relation_A_SOURCE.Name = `relation-A_SOURCE`

	//gong:ident [ref_models.A_SOURCE] comment added to overcome the problem with the comment map association
	__GongStructShape__000001_relation_A_SOURCE.Identifier = `ref_models.A_SOURCE`
	__GongStructShape__000001_relation_A_SOURCE.ShowNbInstances = false
	__GongStructShape__000001_relation_A_SOURCE.NbInstances = 0
	__GongStructShape__000001_relation_A_SOURCE.Width = 240.000000
	__GongStructShape__000001_relation_A_SOURCE.Height = 63.000000
	__GongStructShape__000001_relation_A_SOURCE.IsSelected = false

	__GongStructShape__000002_relation_A_SPEC_RELATIONS.Name = `relation-A_SPEC_RELATIONS`

	//gong:ident [ref_models.A_SPEC_RELATIONS] comment added to overcome the problem with the comment map association
	__GongStructShape__000002_relation_A_SPEC_RELATIONS.Identifier = `ref_models.A_SPEC_RELATIONS`
	__GongStructShape__000002_relation_A_SPEC_RELATIONS.ShowNbInstances = false
	__GongStructShape__000002_relation_A_SPEC_RELATIONS.NbInstances = 0
	__GongStructShape__000002_relation_A_SPEC_RELATIONS.Width = 240.000000
	__GongStructShape__000002_relation_A_SPEC_RELATIONS.Height = 63.000000
	__GongStructShape__000002_relation_A_SPEC_RELATIONS.IsSelected = false

	__GongStructShape__000003_relation_SPECIFICATION.Name = `relation-SPECIFICATION`

	//gong:ident [ref_models.SPECIFICATION] comment added to overcome the problem with the comment map association
	__GongStructShape__000003_relation_SPECIFICATION.Identifier = `ref_models.SPECIFICATION`
	__GongStructShape__000003_relation_SPECIFICATION.ShowNbInstances = false
	__GongStructShape__000003_relation_SPECIFICATION.NbInstances = 0
	__GongStructShape__000003_relation_SPECIFICATION.Width = 240.000000
	__GongStructShape__000003_relation_SPECIFICATION.Height = 63.000000
	__GongStructShape__000003_relation_SPECIFICATION.IsSelected = false

	__GongStructShape__000004_relation_SPEC_RELATION.Name = `relation-SPEC_RELATION`

	//gong:ident [ref_models.SPEC_RELATION] comment added to overcome the problem with the comment map association
	__GongStructShape__000004_relation_SPEC_RELATION.Identifier = `ref_models.SPEC_RELATION`
	__GongStructShape__000004_relation_SPEC_RELATION.ShowNbInstances = false
	__GongStructShape__000004_relation_SPEC_RELATION.NbInstances = 0
	__GongStructShape__000004_relation_SPEC_RELATION.Width = 240.000000
	__GongStructShape__000004_relation_SPEC_RELATION.Height = 406.000000
	__GongStructShape__000004_relation_SPEC_RELATION.IsSelected = false

	__GongStructShape__000005_relation_SPEC_RELATION_TYPE.Name = `relation-SPEC_RELATION_TYPE`

	//gong:ident [ref_models.SPEC_RELATION_TYPE] comment added to overcome the problem with the comment map association
	__GongStructShape__000005_relation_SPEC_RELATION_TYPE.Identifier = `ref_models.SPEC_RELATION_TYPE`
	__GongStructShape__000005_relation_SPEC_RELATION_TYPE.ShowNbInstances = false
	__GongStructShape__000005_relation_SPEC_RELATION_TYPE.NbInstances = 0
	__GongStructShape__000005_relation_SPEC_RELATION_TYPE.Width = 240.000000
	__GongStructShape__000005_relation_SPEC_RELATION_TYPE.Height = 63.000000
	__GongStructShape__000005_relation_SPEC_RELATION_TYPE.IsSelected = false

	__Link__000000_SOURCE.Name = `SOURCE`

	//gong:ident [ref_models.SPEC_RELATION.SOURCE] comment added to overcome the problem with the comment map association
	__Link__000000_SOURCE.Identifier = `ref_models.SPEC_RELATION.SOURCE`

	//gong:ident [ref_models.A_SOURCE] comment added to overcome the problem with the comment map association
	__Link__000000_SOURCE.Fieldtypename = `ref_models.A_SOURCE`
	__Link__000000_SOURCE.FieldOffsetX = -50.000000
	__Link__000000_SOURCE.FieldOffsetY = -16.000000
	__Link__000000_SOURCE.TargetMultiplicity = models.MANY
	__Link__000000_SOURCE.TargetMultiplicityOffsetX = -50.000000
	__Link__000000_SOURCE.TargetMultiplicityOffsetY = 16.000000
	__Link__000000_SOURCE.SourceMultiplicity = models.MANY
	__Link__000000_SOURCE.SourceMultiplicityOffsetX = 10.000000
	__Link__000000_SOURCE.SourceMultiplicityOffsetY = -50.000000
	__Link__000000_SOURCE.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_SOURCE.StartRatio = 0.557605
	__Link__000000_SOURCE.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_SOURCE.EndRatio = 0.500000
	__Link__000000_SOURCE.CornerOffsetRatio = 1.380000

	__Link__000001_SPEC_RELATION.Name = `SPEC_RELATION`

	//gong:ident [ref_models.A_SPEC_RELATIONS.SPEC_RELATION] comment added to overcome the problem with the comment map association
	__Link__000001_SPEC_RELATION.Identifier = `ref_models.A_SPEC_RELATIONS.SPEC_RELATION`

	//gong:ident [ref_models.SPEC_RELATION] comment added to overcome the problem with the comment map association
	__Link__000001_SPEC_RELATION.Fieldtypename = `ref_models.SPEC_RELATION`
	__Link__000001_SPEC_RELATION.FieldOffsetX = -50.000000
	__Link__000001_SPEC_RELATION.FieldOffsetY = -16.000000
	__Link__000001_SPEC_RELATION.TargetMultiplicity = models.MANY
	__Link__000001_SPEC_RELATION.TargetMultiplicityOffsetX = -50.000000
	__Link__000001_SPEC_RELATION.TargetMultiplicityOffsetY = 16.000000
	__Link__000001_SPEC_RELATION.SourceMultiplicity = models.MANY
	__Link__000001_SPEC_RELATION.SourceMultiplicityOffsetX = 10.000000
	__Link__000001_SPEC_RELATION.SourceMultiplicityOffsetY = -50.000000
	__Link__000001_SPEC_RELATION.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000001_SPEC_RELATION.StartRatio = 0.500000
	__Link__000001_SPEC_RELATION.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000001_SPEC_RELATION.EndRatio = 0.500000
	__Link__000001_SPEC_RELATION.CornerOffsetRatio = 1.380000

	__Link__000002_VALUES.Name = `VALUES`

	//gong:ident [ref_models.SPEC_RELATION.VALUES] comment added to overcome the problem with the comment map association
	__Link__000002_VALUES.Identifier = `ref_models.SPEC_RELATION.VALUES`

	//gong:ident [ref_models.A_ATTRIBUTE_VALUE_XHTML_1] comment added to overcome the problem with the comment map association
	__Link__000002_VALUES.Fieldtypename = `ref_models.A_ATTRIBUTE_VALUE_XHTML_1`
	__Link__000002_VALUES.FieldOffsetX = -50.000000
	__Link__000002_VALUES.FieldOffsetY = -16.000000
	__Link__000002_VALUES.TargetMultiplicity = models.MANY
	__Link__000002_VALUES.TargetMultiplicityOffsetX = -50.000000
	__Link__000002_VALUES.TargetMultiplicityOffsetY = 16.000000
	__Link__000002_VALUES.SourceMultiplicity = models.MANY
	__Link__000002_VALUES.SourceMultiplicityOffsetX = 10.000000
	__Link__000002_VALUES.SourceMultiplicityOffsetY = -50.000000
	__Link__000002_VALUES.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000002_VALUES.StartRatio = 0.120265
	__Link__000002_VALUES.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000002_VALUES.EndRatio = 0.500000
	__Link__000002_VALUES.CornerOffsetRatio = 1.380000

	__Position__000000_Pos_relation_A_ATTRIBUTE_VALUE_XHTML_1.X = 1101.000000
	__Position__000000_Pos_relation_A_ATTRIBUTE_VALUE_XHTML_1.Y = 469.000000
	__Position__000000_Pos_relation_A_ATTRIBUTE_VALUE_XHTML_1.Name = `Pos-relation-A_ATTRIBUTE_VALUE_XHTML_1`

	__Position__000001_Pos_relation_A_SOURCE.X = 1120.000000
	__Position__000001_Pos_relation_A_SOURCE.Y = 637.000000
	__Position__000001_Pos_relation_A_SOURCE.Name = `Pos-relation-A_SOURCE`

	__Position__000002_Pos_relation_A_SPEC_RELATIONS.X = 57.000000
	__Position__000002_Pos_relation_A_SPEC_RELATIONS.Y = 237.000000
	__Position__000002_Pos_relation_A_SPEC_RELATIONS.Name = `Pos-relation-A_SPEC_RELATIONS`

	__Position__000003_Pos_relation_SPECIFICATION.X = 892.000000
	__Position__000003_Pos_relation_SPECIFICATION.Y = 114.000000
	__Position__000003_Pos_relation_SPECIFICATION.Name = `Pos-relation-SPECIFICATION`

	__Position__000004_Pos_relation_SPEC_RELATION.X = 599.000000
	__Position__000004_Pos_relation_SPEC_RELATION.Y = 443.000000
	__Position__000004_Pos_relation_SPEC_RELATION.Name = `Pos-relation-SPEC_RELATION`

	__Position__000005_Pos_relation_SPEC_RELATION_TYPE.X = 459.000000
	__Position__000005_Pos_relation_SPEC_RELATION_TYPE.Y = 89.000000
	__Position__000005_Pos_relation_SPEC_RELATION_TYPE.Name = `Pos-relation-SPEC_RELATION_TYPE`

	__Vertice__000000_Verticle_in_class_diagram_relation_in_middle_between_relation_A_SPEC_RELATIONS_and_relation_SPEC_RELATION.X = 399.000000
	__Vertice__000000_Verticle_in_class_diagram_relation_in_middle_between_relation_A_SPEC_RELATIONS_and_relation_SPEC_RELATION.Y = 159.000000
	__Vertice__000000_Verticle_in_class_diagram_relation_in_middle_between_relation_A_SPEC_RELATIONS_and_relation_SPEC_RELATION.Name = `Verticle in class diagram relation in middle between relation-A_SPEC_RELATIONS and relation-SPEC_RELATION`

	__Vertice__000001_Verticle_in_class_diagram_relation_in_middle_between_relation_SPEC_RELATION_and_relation_A_ATTRIBUTE_VALUE_XHTML_1.X = 1193.500000
	__Vertice__000001_Verticle_in_class_diagram_relation_in_middle_between_relation_SPEC_RELATION_and_relation_A_ATTRIBUTE_VALUE_XHTML_1.Y = 505.000000
	__Vertice__000001_Verticle_in_class_diagram_relation_in_middle_between_relation_SPEC_RELATION_and_relation_A_ATTRIBUTE_VALUE_XHTML_1.Name = `Verticle in class diagram relation in middle between relation-SPEC_RELATION and relation-A_ATTRIBUTE_VALUE_XHTML_1`

	__Vertice__000002_Verticle_in_class_diagram_relation_in_middle_between_relation_SPEC_RELATION_and_relation_A_SOURCE.X = 1219.500000
	__Vertice__000002_Verticle_in_class_diagram_relation_in_middle_between_relation_SPEC_RELATION_and_relation_A_SOURCE.Y = 735.500000
	__Vertice__000002_Verticle_in_class_diagram_relation_in_middle_between_relation_SPEC_RELATION_and_relation_A_SOURCE.Name = `Verticle in class diagram relation in middle between relation-SPEC_RELATION and relation-A_SOURCE`

	// Setup of pointers
	__Classdiagram__000000_relation.GongStructShapes = append(__Classdiagram__000000_relation.GongStructShapes, __GongStructShape__000002_relation_A_SPEC_RELATIONS)
	__Classdiagram__000000_relation.GongStructShapes = append(__Classdiagram__000000_relation.GongStructShapes, __GongStructShape__000003_relation_SPECIFICATION)
	__Classdiagram__000000_relation.GongStructShapes = append(__Classdiagram__000000_relation.GongStructShapes, __GongStructShape__000004_relation_SPEC_RELATION)
	__Classdiagram__000000_relation.GongStructShapes = append(__Classdiagram__000000_relation.GongStructShapes, __GongStructShape__000005_relation_SPEC_RELATION_TYPE)
	__Classdiagram__000000_relation.GongStructShapes = append(__Classdiagram__000000_relation.GongStructShapes, __GongStructShape__000000_relation_A_ATTRIBUTE_VALUE_XHTML_1)
	__Classdiagram__000000_relation.GongStructShapes = append(__Classdiagram__000000_relation.GongStructShapes, __GongStructShape__000001_relation_A_SOURCE)
	__GongStructShape__000000_relation_A_ATTRIBUTE_VALUE_XHTML_1.Position = __Position__000000_Pos_relation_A_ATTRIBUTE_VALUE_XHTML_1
	__GongStructShape__000001_relation_A_SOURCE.Position = __Position__000001_Pos_relation_A_SOURCE
	__GongStructShape__000002_relation_A_SPEC_RELATIONS.Position = __Position__000002_Pos_relation_A_SPEC_RELATIONS
	__GongStructShape__000002_relation_A_SPEC_RELATIONS.Links = append(__GongStructShape__000002_relation_A_SPEC_RELATIONS.Links, __Link__000001_SPEC_RELATION)
	__GongStructShape__000003_relation_SPECIFICATION.Position = __Position__000003_Pos_relation_SPECIFICATION
	__GongStructShape__000004_relation_SPEC_RELATION.Position = __Position__000004_Pos_relation_SPEC_RELATION
	__GongStructShape__000004_relation_SPEC_RELATION.Fields = append(__GongStructShape__000004_relation_SPEC_RELATION.Fields, __Field__000004_Name)
	__GongStructShape__000004_relation_SPEC_RELATION.Fields = append(__GongStructShape__000004_relation_SPEC_RELATION.Fields, __Field__000000_DESC)
	__GongStructShape__000004_relation_SPEC_RELATION.Fields = append(__GongStructShape__000004_relation_SPEC_RELATION.Fields, __Field__000001_IDENTIFIER)
	__GongStructShape__000004_relation_SPEC_RELATION.Fields = append(__GongStructShape__000004_relation_SPEC_RELATION.Fields, __Field__000002_LAST_CHANGE)
	__GongStructShape__000004_relation_SPEC_RELATION.Fields = append(__GongStructShape__000004_relation_SPEC_RELATION.Fields, __Field__000003_LONG_NAME)
	__GongStructShape__000004_relation_SPEC_RELATION.Links = append(__GongStructShape__000004_relation_SPEC_RELATION.Links, __Link__000002_VALUES)
	__GongStructShape__000004_relation_SPEC_RELATION.Links = append(__GongStructShape__000004_relation_SPEC_RELATION.Links, __Link__000000_SOURCE)
	__GongStructShape__000005_relation_SPEC_RELATION_TYPE.Position = __Position__000005_Pos_relation_SPEC_RELATION_TYPE
	__Link__000000_SOURCE.Middlevertice = __Vertice__000002_Verticle_in_class_diagram_relation_in_middle_between_relation_SPEC_RELATION_and_relation_A_SOURCE
	__Link__000001_SPEC_RELATION.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_relation_in_middle_between_relation_A_SPEC_RELATIONS_and_relation_SPEC_RELATION
	__Link__000002_VALUES.Middlevertice = __Vertice__000001_Verticle_in_class_diagram_relation_in_middle_between_relation_SPEC_RELATION_and_relation_A_ATTRIBUTE_VALUE_XHTML_1
}
