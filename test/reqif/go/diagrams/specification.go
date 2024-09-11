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

	__Classdiagram__000000_specification := (&models.Classdiagram{Name: `specification`}).Stage(stage)

	__Field__000000_DESC := (&models.Field{Name: `DESC`}).Stage(stage)
	__Field__000001_DESC := (&models.Field{Name: `DESC`}).Stage(stage)
	__Field__000002_IDENTIFIER := (&models.Field{Name: `IDENTIFIER`}).Stage(stage)
	__Field__000003_IDENTIFIER := (&models.Field{Name: `IDENTIFIER`}).Stage(stage)
	__Field__000004_IDENTIFIER := (&models.Field{Name: `IDENTIFIER`}).Stage(stage)
	__Field__000005_IS_EDITABLE := (&models.Field{Name: `IS_EDITABLE`}).Stage(stage)
	__Field__000006_IS_TABLE_INTERNAL := (&models.Field{Name: `IS_TABLE_INTERNAL`}).Stage(stage)
	__Field__000007_LAST_CHANGE := (&models.Field{Name: `LAST_CHANGE`}).Stage(stage)
	__Field__000008_LAST_CHANGE := (&models.Field{Name: `LAST_CHANGE`}).Stage(stage)
	__Field__000009_LONG_NAME := (&models.Field{Name: `LONG_NAME`}).Stage(stage)
	__Field__000010_LONG_NAME := (&models.Field{Name: `LONG_NAME`}).Stage(stage)
	__Field__000011_LONG_NAME := (&models.Field{Name: `LONG_NAME`}).Stage(stage)
	__Field__000012_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000013_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000014_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000015_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000016_SPECIFICATION_TYPE_REF := (&models.Field{Name: `SPECIFICATION_TYPE_REF`}).Stage(stage)
	__Field__000017_SPEC_OBJECT_REF := (&models.Field{Name: `SPEC_OBJECT_REF`}).Stage(stage)

	__GongStructShape__000000_specification_A_CHILDREN := (&models.GongStructShape{Name: `specification-A_CHILDREN`}).Stage(stage)
	__GongStructShape__000001_specification_A_OBJECT := (&models.GongStructShape{Name: `specification-A_OBJECT`}).Stage(stage)
	__GongStructShape__000002_specification_A_SPECIFICATION_TYPE_REF := (&models.GongStructShape{Name: `specification-A_SPECIFICATION_TYPE_REF`}).Stage(stage)
	__GongStructShape__000003_specification_SPECIFICATION := (&models.GongStructShape{Name: `specification-SPECIFICATION`}).Stage(stage)
	__GongStructShape__000004_specification_SPECIFICATION_TYPE := (&models.GongStructShape{Name: `specification-SPECIFICATION_TYPE`}).Stage(stage)
	__GongStructShape__000005_specification_SPEC_HIERARCHY := (&models.GongStructShape{Name: `specification-SPEC_HIERARCHY`}).Stage(stage)
	__GongStructShape__000006_specification_SPEC_OBJECT := (&models.GongStructShape{Name: `specification-SPEC_OBJECT`}).Stage(stage)

	__Link__000000_CHILDREN := (&models.Link{Name: `CHILDREN`}).Stage(stage)
	__Link__000001_CHILDREN := (&models.Link{Name: `CHILDREN`}).Stage(stage)
	__Link__000002_OBJECT := (&models.Link{Name: `OBJECT`}).Stage(stage)
	__Link__000003_SPEC_HIERARCHY := (&models.Link{Name: `SPEC_HIERARCHY`}).Stage(stage)
	__Link__000004_TYPE := (&models.Link{Name: `TYPE`}).Stage(stage)

	__Position__000000_Pos_specification_A_CHILDREN := (&models.Position{Name: `Pos-specification-A_CHILDREN`}).Stage(stage)
	__Position__000001_Pos_specification_A_OBJECT := (&models.Position{Name: `Pos-specification-A_OBJECT`}).Stage(stage)
	__Position__000002_Pos_specification_A_SPECIFICATION_TYPE_REF := (&models.Position{Name: `Pos-specification-A_SPECIFICATION_TYPE_REF`}).Stage(stage)
	__Position__000003_Pos_specification_SPECIFICATION := (&models.Position{Name: `Pos-specification-SPECIFICATION`}).Stage(stage)
	__Position__000004_Pos_specification_SPECIFICATION_TYPE := (&models.Position{Name: `Pos-specification-SPECIFICATION_TYPE`}).Stage(stage)
	__Position__000005_Pos_specification_SPEC_HIERARCHY := (&models.Position{Name: `Pos-specification-SPEC_HIERARCHY`}).Stage(stage)
	__Position__000006_Pos_specification_SPEC_OBJECT := (&models.Position{Name: `Pos-specification-SPEC_OBJECT`}).Stage(stage)

	__Vertice__000000_Verticle_in_class_diagram_specification_in_middle_between_specification_A_CHILDREN_and_specification_SPEC_HIERARCHY := (&models.Vertice{Name: `Verticle in class diagram specification in middle between specification-A_CHILDREN and specification-SPEC_HIERARCHY`}).Stage(stage)
	__Vertice__000001_Verticle_in_class_diagram_specification_in_middle_between_specification_SPECIFICATION_and_specification_A_CHILDREN := (&models.Vertice{Name: `Verticle in class diagram specification in middle between specification-SPECIFICATION and specification-A_CHILDREN`}).Stage(stage)
	__Vertice__000002_Verticle_in_class_diagram_specification_in_middle_between_specification_SPECIFICATION_and_specification_A_SPECIFICATION_TYPE_REF := (&models.Vertice{Name: `Verticle in class diagram specification in middle between specification-SPECIFICATION and specification-A_SPECIFICATION_TYPE_REF`}).Stage(stage)
	__Vertice__000003_Verticle_in_class_diagram_specification_in_middle_between_specification_SPEC_HIERARCHY_and_specification_A_CHILDREN := (&models.Vertice{Name: `Verticle in class diagram specification in middle between specification-SPEC_HIERARCHY and specification-A_CHILDREN`}).Stage(stage)
	__Vertice__000004_Verticle_in_class_diagram_specification_in_middle_between_specification_SPEC_HIERARCHY_and_specification_A_OBJECT := (&models.Vertice{Name: `Verticle in class diagram specification in middle between specification-SPEC_HIERARCHY and specification-A_OBJECT`}).Stage(stage)

	// Setup of values

	__Classdiagram__000000_specification.Name = `specification`
	__Classdiagram__000000_specification.IsInDrawMode = false

	__Field__000000_DESC.Name = `DESC`

	//gong:ident [ref_models.SPEC_OBJECT.DESC] comment added to overcome the problem with the comment map association
	__Field__000000_DESC.Identifier = `ref_models.SPEC_OBJECT.DESC`
	__Field__000000_DESC.FieldTypeAsString = ``
	__Field__000000_DESC.Structname = `SPEC_OBJECT`
	__Field__000000_DESC.Fieldtypename = `string`

	__Field__000001_DESC.Name = `DESC`

	//gong:ident [ref_models.SPEC_HIERARCHY.DESC] comment added to overcome the problem with the comment map association
	__Field__000001_DESC.Identifier = `ref_models.SPEC_HIERARCHY.DESC`
	__Field__000001_DESC.FieldTypeAsString = ``
	__Field__000001_DESC.Structname = `SPEC_HIERARCHY`
	__Field__000001_DESC.Fieldtypename = `string`

	__Field__000002_IDENTIFIER.Name = `IDENTIFIER`

	//gong:ident [ref_models.SPEC_OBJECT.IDENTIFIER] comment added to overcome the problem with the comment map association
	__Field__000002_IDENTIFIER.Identifier = `ref_models.SPEC_OBJECT.IDENTIFIER`
	__Field__000002_IDENTIFIER.FieldTypeAsString = ``
	__Field__000002_IDENTIFIER.Structname = `SPEC_OBJECT`
	__Field__000002_IDENTIFIER.Fieldtypename = `string`

	__Field__000003_IDENTIFIER.Name = `IDENTIFIER`

	//gong:ident [ref_models.SPEC_HIERARCHY.IDENTIFIER] comment added to overcome the problem with the comment map association
	__Field__000003_IDENTIFIER.Identifier = `ref_models.SPEC_HIERARCHY.IDENTIFIER`
	__Field__000003_IDENTIFIER.FieldTypeAsString = ``
	__Field__000003_IDENTIFIER.Structname = `SPEC_HIERARCHY`
	__Field__000003_IDENTIFIER.Fieldtypename = `string`

	__Field__000004_IDENTIFIER.Name = `IDENTIFIER`

	//gong:ident [ref_models.SPECIFICATION.IDENTIFIER] comment added to overcome the problem with the comment map association
	__Field__000004_IDENTIFIER.Identifier = `ref_models.SPECIFICATION.IDENTIFIER`
	__Field__000004_IDENTIFIER.FieldTypeAsString = ``
	__Field__000004_IDENTIFIER.Structname = `SPECIFICATION`
	__Field__000004_IDENTIFIER.Fieldtypename = `string`

	__Field__000005_IS_EDITABLE.Name = `IS_EDITABLE`

	//gong:ident [ref_models.SPEC_HIERARCHY.IS_EDITABLE] comment added to overcome the problem with the comment map association
	__Field__000005_IS_EDITABLE.Identifier = `ref_models.SPEC_HIERARCHY.IS_EDITABLE`
	__Field__000005_IS_EDITABLE.FieldTypeAsString = ``
	__Field__000005_IS_EDITABLE.Structname = `SPEC_HIERARCHY`
	__Field__000005_IS_EDITABLE.Fieldtypename = `bool`

	__Field__000006_IS_TABLE_INTERNAL.Name = `IS_TABLE_INTERNAL`

	//gong:ident [ref_models.SPEC_HIERARCHY.IS_TABLE_INTERNAL] comment added to overcome the problem with the comment map association
	__Field__000006_IS_TABLE_INTERNAL.Identifier = `ref_models.SPEC_HIERARCHY.IS_TABLE_INTERNAL`
	__Field__000006_IS_TABLE_INTERNAL.FieldTypeAsString = ``
	__Field__000006_IS_TABLE_INTERNAL.Structname = `SPEC_HIERARCHY`
	__Field__000006_IS_TABLE_INTERNAL.Fieldtypename = `bool`

	__Field__000007_LAST_CHANGE.Name = `LAST_CHANGE`

	//gong:ident [ref_models.SPEC_OBJECT.LAST_CHANGE] comment added to overcome the problem with the comment map association
	__Field__000007_LAST_CHANGE.Identifier = `ref_models.SPEC_OBJECT.LAST_CHANGE`
	__Field__000007_LAST_CHANGE.FieldTypeAsString = ``
	__Field__000007_LAST_CHANGE.Structname = `SPEC_OBJECT`
	__Field__000007_LAST_CHANGE.Fieldtypename = `string`

	__Field__000008_LAST_CHANGE.Name = `LAST_CHANGE`

	//gong:ident [ref_models.SPECIFICATION.LAST_CHANGE] comment added to overcome the problem with the comment map association
	__Field__000008_LAST_CHANGE.Identifier = `ref_models.SPECIFICATION.LAST_CHANGE`
	__Field__000008_LAST_CHANGE.FieldTypeAsString = ``
	__Field__000008_LAST_CHANGE.Structname = `SPECIFICATION`
	__Field__000008_LAST_CHANGE.Fieldtypename = `string`

	__Field__000009_LONG_NAME.Name = `LONG_NAME`

	//gong:ident [ref_models.SPEC_OBJECT.LONG_NAME] comment added to overcome the problem with the comment map association
	__Field__000009_LONG_NAME.Identifier = `ref_models.SPEC_OBJECT.LONG_NAME`
	__Field__000009_LONG_NAME.FieldTypeAsString = ``
	__Field__000009_LONG_NAME.Structname = `SPEC_OBJECT`
	__Field__000009_LONG_NAME.Fieldtypename = `string`

	__Field__000010_LONG_NAME.Name = `LONG_NAME`

	//gong:ident [ref_models.SPEC_HIERARCHY.LONG_NAME] comment added to overcome the problem with the comment map association
	__Field__000010_LONG_NAME.Identifier = `ref_models.SPEC_HIERARCHY.LONG_NAME`
	__Field__000010_LONG_NAME.FieldTypeAsString = ``
	__Field__000010_LONG_NAME.Structname = `SPEC_HIERARCHY`
	__Field__000010_LONG_NAME.Fieldtypename = `string`

	__Field__000011_LONG_NAME.Name = `LONG_NAME`

	//gong:ident [ref_models.SPECIFICATION.LONG_NAME] comment added to overcome the problem with the comment map association
	__Field__000011_LONG_NAME.Identifier = `ref_models.SPECIFICATION.LONG_NAME`
	__Field__000011_LONG_NAME.FieldTypeAsString = ``
	__Field__000011_LONG_NAME.Structname = `SPECIFICATION`
	__Field__000011_LONG_NAME.Fieldtypename = `string`

	__Field__000012_Name.Name = `Name`

	//gong:ident [ref_models.A_OBJECT.Name] comment added to overcome the problem with the comment map association
	__Field__000012_Name.Identifier = `ref_models.A_OBJECT.Name`
	__Field__000012_Name.FieldTypeAsString = ``
	__Field__000012_Name.Structname = `A_OBJECT`
	__Field__000012_Name.Fieldtypename = `string`

	__Field__000013_Name.Name = `Name`

	//gong:ident [ref_models.A_CHILDREN.Name] comment added to overcome the problem with the comment map association
	__Field__000013_Name.Identifier = `ref_models.A_CHILDREN.Name`
	__Field__000013_Name.FieldTypeAsString = ``
	__Field__000013_Name.Structname = `A_CHILDREN`
	__Field__000013_Name.Fieldtypename = `string`

	__Field__000014_Name.Name = `Name`

	//gong:ident [ref_models.SPEC_HIERARCHY.Name] comment added to overcome the problem with the comment map association
	__Field__000014_Name.Identifier = `ref_models.SPEC_HIERARCHY.Name`
	__Field__000014_Name.FieldTypeAsString = ``
	__Field__000014_Name.Structname = `SPEC_HIERARCHY`
	__Field__000014_Name.Fieldtypename = `string`

	__Field__000015_Name.Name = `Name`

	//gong:ident [ref_models.SPEC_OBJECT.Name] comment added to overcome the problem with the comment map association
	__Field__000015_Name.Identifier = `ref_models.SPEC_OBJECT.Name`
	__Field__000015_Name.FieldTypeAsString = ``
	__Field__000015_Name.Structname = `SPEC_OBJECT`
	__Field__000015_Name.Fieldtypename = `string`

	__Field__000016_SPECIFICATION_TYPE_REF.Name = `SPECIFICATION_TYPE_REF`

	//gong:ident [ref_models.A_SPECIFICATION_TYPE_REF.SPECIFICATION_TYPE_REF] comment added to overcome the problem with the comment map association
	__Field__000016_SPECIFICATION_TYPE_REF.Identifier = `ref_models.A_SPECIFICATION_TYPE_REF.SPECIFICATION_TYPE_REF`
	__Field__000016_SPECIFICATION_TYPE_REF.FieldTypeAsString = ``
	__Field__000016_SPECIFICATION_TYPE_REF.Structname = `A_SPECIFICATION_TYPE_REF`
	__Field__000016_SPECIFICATION_TYPE_REF.Fieldtypename = `string`

	__Field__000017_SPEC_OBJECT_REF.Name = `SPEC_OBJECT_REF`

	//gong:ident [ref_models.A_OBJECT.SPEC_OBJECT_REF] comment added to overcome the problem with the comment map association
	__Field__000017_SPEC_OBJECT_REF.Identifier = `ref_models.A_OBJECT.SPEC_OBJECT_REF`
	__Field__000017_SPEC_OBJECT_REF.FieldTypeAsString = ``
	__Field__000017_SPEC_OBJECT_REF.Structname = `A_OBJECT`
	__Field__000017_SPEC_OBJECT_REF.Fieldtypename = `string`

	__GongStructShape__000000_specification_A_CHILDREN.Name = `specification-A_CHILDREN`

	//gong:ident [ref_models.A_CHILDREN] comment added to overcome the problem with the comment map association
	__GongStructShape__000000_specification_A_CHILDREN.Identifier = `ref_models.A_CHILDREN`
	__GongStructShape__000000_specification_A_CHILDREN.ShowNbInstances = false
	__GongStructShape__000000_specification_A_CHILDREN.NbInstances = 0
	__GongStructShape__000000_specification_A_CHILDREN.Width = 240.000000
	__GongStructShape__000000_specification_A_CHILDREN.Height = 78.000000
	__GongStructShape__000000_specification_A_CHILDREN.IsSelected = false

	__GongStructShape__000001_specification_A_OBJECT.Name = `specification-A_OBJECT`

	//gong:ident [ref_models.A_OBJECT] comment added to overcome the problem with the comment map association
	__GongStructShape__000001_specification_A_OBJECT.Identifier = `ref_models.A_OBJECT`
	__GongStructShape__000001_specification_A_OBJECT.ShowNbInstances = false
	__GongStructShape__000001_specification_A_OBJECT.NbInstances = 0
	__GongStructShape__000001_specification_A_OBJECT.Width = 240.000000
	__GongStructShape__000001_specification_A_OBJECT.Height = 93.000000
	__GongStructShape__000001_specification_A_OBJECT.IsSelected = false

	__GongStructShape__000002_specification_A_SPECIFICATION_TYPE_REF.Name = `specification-A_SPECIFICATION_TYPE_REF`

	//gong:ident [ref_models.A_SPECIFICATION_TYPE_REF] comment added to overcome the problem with the comment map association
	__GongStructShape__000002_specification_A_SPECIFICATION_TYPE_REF.Identifier = `ref_models.A_SPECIFICATION_TYPE_REF`
	__GongStructShape__000002_specification_A_SPECIFICATION_TYPE_REF.ShowNbInstances = false
	__GongStructShape__000002_specification_A_SPECIFICATION_TYPE_REF.NbInstances = 0
	__GongStructShape__000002_specification_A_SPECIFICATION_TYPE_REF.Width = 240.000000
	__GongStructShape__000002_specification_A_SPECIFICATION_TYPE_REF.Height = 78.000000
	__GongStructShape__000002_specification_A_SPECIFICATION_TYPE_REF.IsSelected = false

	__GongStructShape__000003_specification_SPECIFICATION.Name = `specification-SPECIFICATION`

	//gong:ident [ref_models.SPECIFICATION] comment added to overcome the problem with the comment map association
	__GongStructShape__000003_specification_SPECIFICATION.Identifier = `ref_models.SPECIFICATION`
	__GongStructShape__000003_specification_SPECIFICATION.ShowNbInstances = false
	__GongStructShape__000003_specification_SPECIFICATION.NbInstances = 0
	__GongStructShape__000003_specification_SPECIFICATION.Width = 240.000000
	__GongStructShape__000003_specification_SPECIFICATION.Height = 108.000000
	__GongStructShape__000003_specification_SPECIFICATION.IsSelected = false

	__GongStructShape__000004_specification_SPECIFICATION_TYPE.Name = `specification-SPECIFICATION_TYPE`

	//gong:ident [ref_models.SPECIFICATION_TYPE] comment added to overcome the problem with the comment map association
	__GongStructShape__000004_specification_SPECIFICATION_TYPE.Identifier = `ref_models.SPECIFICATION_TYPE`
	__GongStructShape__000004_specification_SPECIFICATION_TYPE.ShowNbInstances = false
	__GongStructShape__000004_specification_SPECIFICATION_TYPE.NbInstances = 0
	__GongStructShape__000004_specification_SPECIFICATION_TYPE.Width = 240.000000
	__GongStructShape__000004_specification_SPECIFICATION_TYPE.Height = 63.000000
	__GongStructShape__000004_specification_SPECIFICATION_TYPE.IsSelected = false

	__GongStructShape__000005_specification_SPEC_HIERARCHY.Name = `specification-SPEC_HIERARCHY`

	//gong:ident [ref_models.SPEC_HIERARCHY] comment added to overcome the problem with the comment map association
	__GongStructShape__000005_specification_SPEC_HIERARCHY.Identifier = `ref_models.SPEC_HIERARCHY`
	__GongStructShape__000005_specification_SPEC_HIERARCHY.ShowNbInstances = false
	__GongStructShape__000005_specification_SPEC_HIERARCHY.NbInstances = 0
	__GongStructShape__000005_specification_SPEC_HIERARCHY.Width = 240.000000
	__GongStructShape__000005_specification_SPEC_HIERARCHY.Height = 153.000000
	__GongStructShape__000005_specification_SPEC_HIERARCHY.IsSelected = false

	__GongStructShape__000006_specification_SPEC_OBJECT.Name = `specification-SPEC_OBJECT`

	//gong:ident [ref_models.SPEC_OBJECT] comment added to overcome the problem with the comment map association
	__GongStructShape__000006_specification_SPEC_OBJECT.Identifier = `ref_models.SPEC_OBJECT`
	__GongStructShape__000006_specification_SPEC_OBJECT.ShowNbInstances = false
	__GongStructShape__000006_specification_SPEC_OBJECT.NbInstances = 0
	__GongStructShape__000006_specification_SPEC_OBJECT.Width = 240.000000
	__GongStructShape__000006_specification_SPEC_OBJECT.Height = 138.000000
	__GongStructShape__000006_specification_SPEC_OBJECT.IsSelected = false

	__Link__000000_CHILDREN.Name = `CHILDREN`

	//gong:ident [ref_models.SPEC_HIERARCHY.CHILDREN] comment added to overcome the problem with the comment map association
	__Link__000000_CHILDREN.Identifier = `ref_models.SPEC_HIERARCHY.CHILDREN`

	//gong:ident [ref_models.A_CHILDREN] comment added to overcome the problem with the comment map association
	__Link__000000_CHILDREN.Fieldtypename = `ref_models.A_CHILDREN`
	__Link__000000_CHILDREN.FieldOffsetX = -50.000000
	__Link__000000_CHILDREN.FieldOffsetY = -16.000000
	__Link__000000_CHILDREN.TargetMultiplicity = models.ZERO_ONE
	__Link__000000_CHILDREN.TargetMultiplicityOffsetX = -50.000000
	__Link__000000_CHILDREN.TargetMultiplicityOffsetY = 16.000000
	__Link__000000_CHILDREN.SourceMultiplicity = models.MANY
	__Link__000000_CHILDREN.SourceMultiplicityOffsetX = 10.000000
	__Link__000000_CHILDREN.SourceMultiplicityOffsetY = -50.000000
	__Link__000000_CHILDREN.StartOrientation = models.ORIENTATION_VERTICAL
	__Link__000000_CHILDREN.StartRatio = 0.349861
	__Link__000000_CHILDREN.EndOrientation = models.ORIENTATION_VERTICAL
	__Link__000000_CHILDREN.EndRatio = 0.412361
	__Link__000000_CHILDREN.CornerOffsetRatio = 1.851634

	__Link__000001_CHILDREN.Name = `CHILDREN`

	//gong:ident [ref_models.SPECIFICATION.CHILDREN] comment added to overcome the problem with the comment map association
	__Link__000001_CHILDREN.Identifier = `ref_models.SPECIFICATION.CHILDREN`

	//gong:ident [ref_models.A_CHILDREN] comment added to overcome the problem with the comment map association
	__Link__000001_CHILDREN.Fieldtypename = `ref_models.A_CHILDREN`
	__Link__000001_CHILDREN.FieldOffsetX = -50.000000
	__Link__000001_CHILDREN.FieldOffsetY = -16.000000
	__Link__000001_CHILDREN.TargetMultiplicity = models.ZERO_ONE
	__Link__000001_CHILDREN.TargetMultiplicityOffsetX = -50.000000
	__Link__000001_CHILDREN.TargetMultiplicityOffsetY = 16.000000
	__Link__000001_CHILDREN.SourceMultiplicity = models.MANY
	__Link__000001_CHILDREN.SourceMultiplicityOffsetX = 10.000000
	__Link__000001_CHILDREN.SourceMultiplicityOffsetY = -50.000000
	__Link__000001_CHILDREN.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000001_CHILDREN.StartRatio = 0.500000
	__Link__000001_CHILDREN.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000001_CHILDREN.EndRatio = 0.500000
	__Link__000001_CHILDREN.CornerOffsetRatio = 1.380000

	__Link__000002_OBJECT.Name = `OBJECT`

	//gong:ident [ref_models.SPEC_HIERARCHY.OBJECT] comment added to overcome the problem with the comment map association
	__Link__000002_OBJECT.Identifier = `ref_models.SPEC_HIERARCHY.OBJECT`

	//gong:ident [ref_models.A_OBJECT] comment added to overcome the problem with the comment map association
	__Link__000002_OBJECT.Fieldtypename = `ref_models.A_OBJECT`
	__Link__000002_OBJECT.FieldOffsetX = -50.000000
	__Link__000002_OBJECT.FieldOffsetY = -16.000000
	__Link__000002_OBJECT.TargetMultiplicity = models.ZERO_ONE
	__Link__000002_OBJECT.TargetMultiplicityOffsetX = -50.000000
	__Link__000002_OBJECT.TargetMultiplicityOffsetY = 16.000000
	__Link__000002_OBJECT.SourceMultiplicity = models.MANY
	__Link__000002_OBJECT.SourceMultiplicityOffsetX = 10.000000
	__Link__000002_OBJECT.SourceMultiplicityOffsetY = -50.000000
	__Link__000002_OBJECT.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000002_OBJECT.StartRatio = 0.500000
	__Link__000002_OBJECT.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000002_OBJECT.EndRatio = 0.500000
	__Link__000002_OBJECT.CornerOffsetRatio = 1.380000

	__Link__000003_SPEC_HIERARCHY.Name = `SPEC_HIERARCHY`

	//gong:ident [ref_models.A_CHILDREN.SPEC_HIERARCHY] comment added to overcome the problem with the comment map association
	__Link__000003_SPEC_HIERARCHY.Identifier = `ref_models.A_CHILDREN.SPEC_HIERARCHY`

	//gong:ident [ref_models.SPEC_HIERARCHY] comment added to overcome the problem with the comment map association
	__Link__000003_SPEC_HIERARCHY.Fieldtypename = `ref_models.SPEC_HIERARCHY`
	__Link__000003_SPEC_HIERARCHY.FieldOffsetX = -50.000000
	__Link__000003_SPEC_HIERARCHY.FieldOffsetY = -16.000000
	__Link__000003_SPEC_HIERARCHY.TargetMultiplicity = models.MANY
	__Link__000003_SPEC_HIERARCHY.TargetMultiplicityOffsetX = -50.000000
	__Link__000003_SPEC_HIERARCHY.TargetMultiplicityOffsetY = 16.000000
	__Link__000003_SPEC_HIERARCHY.SourceMultiplicity = models.MANY
	__Link__000003_SPEC_HIERARCHY.SourceMultiplicityOffsetX = 10.000000
	__Link__000003_SPEC_HIERARCHY.SourceMultiplicityOffsetY = -50.000000
	__Link__000003_SPEC_HIERARCHY.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000003_SPEC_HIERARCHY.StartRatio = 0.500000
	__Link__000003_SPEC_HIERARCHY.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000003_SPEC_HIERARCHY.EndRatio = 0.500000
	__Link__000003_SPEC_HIERARCHY.CornerOffsetRatio = 1.380000

	__Link__000004_TYPE.Name = `TYPE`

	//gong:ident [ref_models.SPECIFICATION.TYPE] comment added to overcome the problem with the comment map association
	__Link__000004_TYPE.Identifier = `ref_models.SPECIFICATION.TYPE`

	//gong:ident [ref_models.A_SPECIFICATION_TYPE_REF] comment added to overcome the problem with the comment map association
	__Link__000004_TYPE.Fieldtypename = `ref_models.A_SPECIFICATION_TYPE_REF`
	__Link__000004_TYPE.FieldOffsetX = -50.000000
	__Link__000004_TYPE.FieldOffsetY = -16.000000
	__Link__000004_TYPE.TargetMultiplicity = models.ZERO_ONE
	__Link__000004_TYPE.TargetMultiplicityOffsetX = -50.000000
	__Link__000004_TYPE.TargetMultiplicityOffsetY = 16.000000
	__Link__000004_TYPE.SourceMultiplicity = models.MANY
	__Link__000004_TYPE.SourceMultiplicityOffsetX = 10.000000
	__Link__000004_TYPE.SourceMultiplicityOffsetY = -50.000000
	__Link__000004_TYPE.StartOrientation = models.ORIENTATION_VERTICAL
	__Link__000004_TYPE.StartRatio = 0.284722
	__Link__000004_TYPE.EndOrientation = models.ORIENTATION_VERTICAL
	__Link__000004_TYPE.EndRatio = 0.255555
	__Link__000004_TYPE.CornerOffsetRatio = 1.510802

	__Position__000000_Pos_specification_A_CHILDREN.X = 434.000000
	__Position__000000_Pos_specification_A_CHILDREN.Y = 74.000000
	__Position__000000_Pos_specification_A_CHILDREN.Name = `Pos-specification-A_CHILDREN`

	__Position__000001_Pos_specification_A_OBJECT.X = 1317.999939
	__Position__000001_Pos_specification_A_OBJECT.Y = 79.000031
	__Position__000001_Pos_specification_A_OBJECT.Name = `Pos-specification-A_OBJECT`

	__Position__000002_Pos_specification_A_SPECIFICATION_TYPE_REF.X = 37.000000
	__Position__000002_Pos_specification_A_SPECIFICATION_TYPE_REF.Y = 255.000000
	__Position__000002_Pos_specification_A_SPECIFICATION_TYPE_REF.Name = `Pos-specification-A_SPECIFICATION_TYPE_REF`

	__Position__000003_Pos_specification_SPECIFICATION.X = 28.000000
	__Position__000003_Pos_specification_SPECIFICATION.Y = 55.000000
	__Position__000003_Pos_specification_SPECIFICATION.Name = `Pos-specification-SPECIFICATION`

	__Position__000004_Pos_specification_SPECIFICATION_TYPE.X = 35.000000
	__Position__000004_Pos_specification_SPECIFICATION_TYPE.Y = 443.000000
	__Position__000004_Pos_specification_SPECIFICATION_TYPE.Name = `Pos-specification-SPECIFICATION_TYPE`

	__Position__000005_Pos_specification_SPEC_HIERARCHY.X = 913.000000
	__Position__000005_Pos_specification_SPEC_HIERARCHY.Y = 47.000000
	__Position__000005_Pos_specification_SPEC_HIERARCHY.Name = `Pos-specification-SPEC_HIERARCHY`

	__Position__000006_Pos_specification_SPEC_OBJECT.X = 38.000000
	__Position__000006_Pos_specification_SPEC_OBJECT.Y = 617.000000
	__Position__000006_Pos_specification_SPEC_OBJECT.Name = `Pos-specification-SPEC_OBJECT`

	__Vertice__000000_Verticle_in_class_diagram_specification_in_middle_between_specification_A_CHILDREN_and_specification_SPEC_HIERARCHY.X = 989.500000
	__Vertice__000000_Verticle_in_class_diagram_specification_in_middle_between_specification_A_CHILDREN_and_specification_SPEC_HIERARCHY.Y = 116.500000
	__Vertice__000000_Verticle_in_class_diagram_specification_in_middle_between_specification_A_CHILDREN_and_specification_SPEC_HIERARCHY.Name = `Verticle in class diagram specification in middle between specification-A_CHILDREN and specification-SPEC_HIERARCHY`

	__Vertice__000001_Verticle_in_class_diagram_specification_in_middle_between_specification_SPECIFICATION_and_specification_A_CHILDREN.X = 561.500000
	__Vertice__000001_Verticle_in_class_diagram_specification_in_middle_between_specification_SPECIFICATION_and_specification_A_CHILDREN.Y = 111.000000
	__Vertice__000001_Verticle_in_class_diagram_specification_in_middle_between_specification_SPECIFICATION_and_specification_A_CHILDREN.Name = `Verticle in class diagram specification in middle between specification-SPECIFICATION and specification-A_CHILDREN`

	__Vertice__000002_Verticle_in_class_diagram_specification_in_middle_between_specification_SPECIFICATION_and_specification_A_SPECIFICATION_TYPE_REF.X = 391.500000
	__Vertice__000002_Verticle_in_class_diagram_specification_in_middle_between_specification_SPECIFICATION_and_specification_A_SPECIFICATION_TYPE_REF.Y = 296.500000
	__Vertice__000002_Verticle_in_class_diagram_specification_in_middle_between_specification_SPECIFICATION_and_specification_A_SPECIFICATION_TYPE_REF.Name = `Verticle in class diagram specification in middle between specification-SPECIFICATION and specification-A_SPECIFICATION_TYPE_REF`

	__Vertice__000003_Verticle_in_class_diagram_specification_in_middle_between_specification_SPEC_HIERARCHY_and_specification_A_CHILDREN.X = 989.500000
	__Vertice__000003_Verticle_in_class_diagram_specification_in_middle_between_specification_SPEC_HIERARCHY_and_specification_A_CHILDREN.Y = 154.000000
	__Vertice__000003_Verticle_in_class_diagram_specification_in_middle_between_specification_SPEC_HIERARCHY_and_specification_A_CHILDREN.Name = `Verticle in class diagram specification in middle between specification-SPEC_HIERARCHY and specification-A_CHILDREN`

	__Vertice__000004_Verticle_in_class_diagram_specification_in_middle_between_specification_SPEC_HIERARCHY_and_specification_A_OBJECT.X = 1199.500000
	__Vertice__000004_Verticle_in_class_diagram_specification_in_middle_between_specification_SPEC_HIERARCHY_and_specification_A_OBJECT.Y = 377.500000
	__Vertice__000004_Verticle_in_class_diagram_specification_in_middle_between_specification_SPEC_HIERARCHY_and_specification_A_OBJECT.Name = `Verticle in class diagram specification in middle between specification-SPEC_HIERARCHY and specification-A_OBJECT`

	// Setup of pointers
	__Classdiagram__000000_specification.GongStructShapes = append(__Classdiagram__000000_specification.GongStructShapes, __GongStructShape__000003_specification_SPECIFICATION)
	__Classdiagram__000000_specification.GongStructShapes = append(__Classdiagram__000000_specification.GongStructShapes, __GongStructShape__000000_specification_A_CHILDREN)
	__Classdiagram__000000_specification.GongStructShapes = append(__Classdiagram__000000_specification.GongStructShapes, __GongStructShape__000005_specification_SPEC_HIERARCHY)
	__Classdiagram__000000_specification.GongStructShapes = append(__Classdiagram__000000_specification.GongStructShapes, __GongStructShape__000006_specification_SPEC_OBJECT)
	__Classdiagram__000000_specification.GongStructShapes = append(__Classdiagram__000000_specification.GongStructShapes, __GongStructShape__000001_specification_A_OBJECT)
	__Classdiagram__000000_specification.GongStructShapes = append(__Classdiagram__000000_specification.GongStructShapes, __GongStructShape__000004_specification_SPECIFICATION_TYPE)
	__Classdiagram__000000_specification.GongStructShapes = append(__Classdiagram__000000_specification.GongStructShapes, __GongStructShape__000002_specification_A_SPECIFICATION_TYPE_REF)
	__GongStructShape__000000_specification_A_CHILDREN.Position = __Position__000000_Pos_specification_A_CHILDREN
	__GongStructShape__000000_specification_A_CHILDREN.Fields = append(__GongStructShape__000000_specification_A_CHILDREN.Fields, __Field__000013_Name)
	__GongStructShape__000000_specification_A_CHILDREN.Links = append(__GongStructShape__000000_specification_A_CHILDREN.Links, __Link__000003_SPEC_HIERARCHY)
	__GongStructShape__000001_specification_A_OBJECT.Position = __Position__000001_Pos_specification_A_OBJECT
	__GongStructShape__000001_specification_A_OBJECT.Fields = append(__GongStructShape__000001_specification_A_OBJECT.Fields, __Field__000012_Name)
	__GongStructShape__000001_specification_A_OBJECT.Fields = append(__GongStructShape__000001_specification_A_OBJECT.Fields, __Field__000017_SPEC_OBJECT_REF)
	__GongStructShape__000002_specification_A_SPECIFICATION_TYPE_REF.Position = __Position__000002_Pos_specification_A_SPECIFICATION_TYPE_REF
	__GongStructShape__000002_specification_A_SPECIFICATION_TYPE_REF.Fields = append(__GongStructShape__000002_specification_A_SPECIFICATION_TYPE_REF.Fields, __Field__000016_SPECIFICATION_TYPE_REF)
	__GongStructShape__000003_specification_SPECIFICATION.Position = __Position__000003_Pos_specification_SPECIFICATION
	__GongStructShape__000003_specification_SPECIFICATION.Fields = append(__GongStructShape__000003_specification_SPECIFICATION.Fields, __Field__000004_IDENTIFIER)
	__GongStructShape__000003_specification_SPECIFICATION.Fields = append(__GongStructShape__000003_specification_SPECIFICATION.Fields, __Field__000008_LAST_CHANGE)
	__GongStructShape__000003_specification_SPECIFICATION.Fields = append(__GongStructShape__000003_specification_SPECIFICATION.Fields, __Field__000011_LONG_NAME)
	__GongStructShape__000003_specification_SPECIFICATION.Links = append(__GongStructShape__000003_specification_SPECIFICATION.Links, __Link__000001_CHILDREN)
	__GongStructShape__000003_specification_SPECIFICATION.Links = append(__GongStructShape__000003_specification_SPECIFICATION.Links, __Link__000004_TYPE)
	__GongStructShape__000004_specification_SPECIFICATION_TYPE.Position = __Position__000004_Pos_specification_SPECIFICATION_TYPE
	__GongStructShape__000005_specification_SPEC_HIERARCHY.Position = __Position__000005_Pos_specification_SPEC_HIERARCHY
	__GongStructShape__000005_specification_SPEC_HIERARCHY.Fields = append(__GongStructShape__000005_specification_SPEC_HIERARCHY.Fields, __Field__000014_Name)
	__GongStructShape__000005_specification_SPEC_HIERARCHY.Fields = append(__GongStructShape__000005_specification_SPEC_HIERARCHY.Fields, __Field__000001_DESC)
	__GongStructShape__000005_specification_SPEC_HIERARCHY.Fields = append(__GongStructShape__000005_specification_SPEC_HIERARCHY.Fields, __Field__000003_IDENTIFIER)
	__GongStructShape__000005_specification_SPEC_HIERARCHY.Fields = append(__GongStructShape__000005_specification_SPEC_HIERARCHY.Fields, __Field__000005_IS_EDITABLE)
	__GongStructShape__000005_specification_SPEC_HIERARCHY.Fields = append(__GongStructShape__000005_specification_SPEC_HIERARCHY.Fields, __Field__000006_IS_TABLE_INTERNAL)
	__GongStructShape__000005_specification_SPEC_HIERARCHY.Fields = append(__GongStructShape__000005_specification_SPEC_HIERARCHY.Fields, __Field__000010_LONG_NAME)
	__GongStructShape__000005_specification_SPEC_HIERARCHY.Links = append(__GongStructShape__000005_specification_SPEC_HIERARCHY.Links, __Link__000000_CHILDREN)
	__GongStructShape__000005_specification_SPEC_HIERARCHY.Links = append(__GongStructShape__000005_specification_SPEC_HIERARCHY.Links, __Link__000002_OBJECT)
	__GongStructShape__000006_specification_SPEC_OBJECT.Position = __Position__000006_Pos_specification_SPEC_OBJECT
	__GongStructShape__000006_specification_SPEC_OBJECT.Fields = append(__GongStructShape__000006_specification_SPEC_OBJECT.Fields, __Field__000015_Name)
	__GongStructShape__000006_specification_SPEC_OBJECT.Fields = append(__GongStructShape__000006_specification_SPEC_OBJECT.Fields, __Field__000000_DESC)
	__GongStructShape__000006_specification_SPEC_OBJECT.Fields = append(__GongStructShape__000006_specification_SPEC_OBJECT.Fields, __Field__000002_IDENTIFIER)
	__GongStructShape__000006_specification_SPEC_OBJECT.Fields = append(__GongStructShape__000006_specification_SPEC_OBJECT.Fields, __Field__000007_LAST_CHANGE)
	__GongStructShape__000006_specification_SPEC_OBJECT.Fields = append(__GongStructShape__000006_specification_SPEC_OBJECT.Fields, __Field__000009_LONG_NAME)
	__Link__000000_CHILDREN.Middlevertice = __Vertice__000003_Verticle_in_class_diagram_specification_in_middle_between_specification_SPEC_HIERARCHY_and_specification_A_CHILDREN
	__Link__000001_CHILDREN.Middlevertice = __Vertice__000001_Verticle_in_class_diagram_specification_in_middle_between_specification_SPECIFICATION_and_specification_A_CHILDREN
	__Link__000002_OBJECT.Middlevertice = __Vertice__000004_Verticle_in_class_diagram_specification_in_middle_between_specification_SPEC_HIERARCHY_and_specification_A_OBJECT
	__Link__000003_SPEC_HIERARCHY.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_specification_in_middle_between_specification_A_CHILDREN_and_specification_SPEC_HIERARCHY
	__Link__000004_TYPE.Middlevertice = __Vertice__000002_Verticle_in_class_diagram_specification_in_middle_between_specification_SPECIFICATION_and_specification_A_SPECIFICATION_TYPE_REF
}
