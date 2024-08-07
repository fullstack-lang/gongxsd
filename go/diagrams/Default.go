package diagrams

import (
	"time"

	"github.com/fullstack-lang/gongdoc/go/models"

	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gongxsd/go/models"
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

	__Field__000000_Base := (&models.Field{Name: `Base`}).Stage(stage)
	__Field__000001_Type := (&models.Field{Name: `Type`}).Stage(stage)
	__Field__000002_Value := (&models.Field{Name: `Value`}).Stage(stage)

	__GongStructShape__000000_Default_ComplexType := (&models.GongStructShape{Name: `Default-ComplexType`}).Stage(stage)
	__GongStructShape__000001_Default_Element := (&models.GongStructShape{Name: `Default-Element`}).Stage(stage)
	__GongStructShape__000002_Default_Enumeration := (&models.GongStructShape{Name: `Default-Enumeration`}).Stage(stage)
	__GongStructShape__000003_Default_MaxInclusive := (&models.GongStructShape{Name: `Default-MaxInclusive`}).Stage(stage)
	__GongStructShape__000004_Default_MinInclusive := (&models.GongStructShape{Name: `Default-MinInclusive`}).Stage(stage)
	__GongStructShape__000005_Default_Restriction := (&models.GongStructShape{Name: `Default-Restriction`}).Stage(stage)
	__GongStructShape__000006_Default_Schema := (&models.GongStructShape{Name: `Default-Schema`}).Stage(stage)
	__GongStructShape__000007_Default_Sequence := (&models.GongStructShape{Name: `Default-Sequence`}).Stage(stage)
	__GongStructShape__000008_Default_SimpleType := (&models.GongStructShape{Name: `Default-SimpleType`}).Stage(stage)

	__Link__000000_ComplexType := (&models.Link{Name: `ComplexType`}).Stage(stage)
	__Link__000001_ComplexTypes := (&models.Link{Name: `ComplexTypes`}).Stage(stage)
	__Link__000002_Elements := (&models.Link{Name: `Elements`}).Stage(stage)
	__Link__000003_Elements := (&models.Link{Name: `Elements`}).Stage(stage)
	__Link__000004_Enumerations := (&models.Link{Name: `Enumerations`}).Stage(stage)
	__Link__000005_MaxInclusive := (&models.Link{Name: `MaxInclusive`}).Stage(stage)
	__Link__000006_MinInclusive := (&models.Link{Name: `MinInclusive`}).Stage(stage)
	__Link__000007_Restriction := (&models.Link{Name: `Restriction`}).Stage(stage)
	__Link__000008_Sequence := (&models.Link{Name: `Sequence`}).Stage(stage)
	__Link__000009_SimpleType := (&models.Link{Name: `SimpleType`}).Stage(stage)
	__Link__000010_SimpleTypes := (&models.Link{Name: `SimpleTypes`}).Stage(stage)

	__Position__000000_Pos_Default_ComplexType := (&models.Position{Name: `Pos-Default-ComplexType`}).Stage(stage)
	__Position__000001_Pos_Default_Element := (&models.Position{Name: `Pos-Default-Element`}).Stage(stage)
	__Position__000002_Pos_Default_Enumeration := (&models.Position{Name: `Pos-Default-Enumeration`}).Stage(stage)
	__Position__000003_Pos_Default_MaxInclusive := (&models.Position{Name: `Pos-Default-MaxInclusive`}).Stage(stage)
	__Position__000004_Pos_Default_MinInclusive := (&models.Position{Name: `Pos-Default-MinInclusive`}).Stage(stage)
	__Position__000005_Pos_Default_Restriction := (&models.Position{Name: `Pos-Default-Restriction`}).Stage(stage)
	__Position__000006_Pos_Default_Schema := (&models.Position{Name: `Pos-Default-Schema`}).Stage(stage)
	__Position__000007_Pos_Default_Sequence := (&models.Position{Name: `Pos-Default-Sequence`}).Stage(stage)
	__Position__000008_Pos_Default_SimpleType := (&models.Position{Name: `Pos-Default-SimpleType`}).Stage(stage)

	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_ComplexType_and_Default_Sequence := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-ComplexType and Default-Sequence`}).Stage(stage)
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Element_and_Default_ComplexType := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-Element and Default-ComplexType`}).Stage(stage)
	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_Element_and_Default_SimpleType := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-Element and Default-SimpleType`}).Stage(stage)
	__Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_Restriction_and_Default_Enumeration := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-Restriction and Default-Enumeration`}).Stage(stage)
	__Vertice__000004_Verticle_in_class_diagram_Default_in_middle_between_Default_Restriction_and_Default_MaxInclusive := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-Restriction and Default-MaxInclusive`}).Stage(stage)
	__Vertice__000005_Verticle_in_class_diagram_Default_in_middle_between_Default_Restriction_and_Default_MinInclusive := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-Restriction and Default-MinInclusive`}).Stage(stage)
	__Vertice__000006_Verticle_in_class_diagram_Default_in_middle_between_Default_Schema_and_Default_ComplexType := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-Schema and Default-ComplexType`}).Stage(stage)
	__Vertice__000007_Verticle_in_class_diagram_Default_in_middle_between_Default_Schema_and_Default_Element := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-Schema and Default-Element`}).Stage(stage)
	__Vertice__000008_Verticle_in_class_diagram_Default_in_middle_between_Default_Schema_and_Default_SimpleType := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-Schema and Default-SimpleType`}).Stage(stage)
	__Vertice__000009_Verticle_in_class_diagram_Default_in_middle_between_Default_Sequence_and_Default_Element := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-Sequence and Default-Element`}).Stage(stage)
	__Vertice__000010_Verticle_in_class_diagram_Default_in_middle_between_Default_SimpleType_and_Default_Restriction := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-SimpleType and Default-Restriction`}).Stage(stage)

	// Setup of values

	__Classdiagram__000000_Default.Name = `Default`
	__Classdiagram__000000_Default.IsInDrawMode = false

	__Field__000000_Base.Name = `Base`

	//gong:ident [ref_models.Restriction.Base] comment added to overcome the problem with the comment map association
	__Field__000000_Base.Identifier = `ref_models.Restriction.Base`
	__Field__000000_Base.FieldTypeAsString = ``
	__Field__000000_Base.Structname = `Restriction`
	__Field__000000_Base.Fieldtypename = `string`

	__Field__000001_Type.Name = `Type`

	//gong:ident [ref_models.Element.Type] comment added to overcome the problem with the comment map association
	__Field__000001_Type.Identifier = `ref_models.Element.Type`
	__Field__000001_Type.FieldTypeAsString = ``
	__Field__000001_Type.Structname = `Element`
	__Field__000001_Type.Fieldtypename = `string`

	__Field__000002_Value.Name = `Value`

	//gong:ident [ref_models.Enumeration.Value] comment added to overcome the problem with the comment map association
	__Field__000002_Value.Identifier = `ref_models.Enumeration.Value`
	__Field__000002_Value.FieldTypeAsString = ``
	__Field__000002_Value.Structname = `Enumeration`
	__Field__000002_Value.Fieldtypename = `string`

	__GongStructShape__000000_Default_ComplexType.Name = `Default-ComplexType`

	//gong:ident [ref_models.ComplexType] comment added to overcome the problem with the comment map association
	__GongStructShape__000000_Default_ComplexType.Identifier = `ref_models.ComplexType`
	__GongStructShape__000000_Default_ComplexType.ShowNbInstances = false
	__GongStructShape__000000_Default_ComplexType.NbInstances = 2
	__GongStructShape__000000_Default_ComplexType.Width = 269.000000
	__GongStructShape__000000_Default_ComplexType.Height = 63.000000
	__GongStructShape__000000_Default_ComplexType.IsSelected = false

	__GongStructShape__000001_Default_Element.Name = `Default-Element`

	//gong:ident [ref_models.Element] comment added to overcome the problem with the comment map association
	__GongStructShape__000001_Default_Element.Identifier = `ref_models.Element`
	__GongStructShape__000001_Default_Element.ShowNbInstances = false
	__GongStructShape__000001_Default_Element.NbInstances = 6
	__GongStructShape__000001_Default_Element.Width = 240.000000
	__GongStructShape__000001_Default_Element.Height = 297.000000
	__GongStructShape__000001_Default_Element.IsSelected = false

	__GongStructShape__000002_Default_Enumeration.Name = `Default-Enumeration`

	//gong:ident [ref_models.Enumeration] comment added to overcome the problem with the comment map association
	__GongStructShape__000002_Default_Enumeration.Identifier = `ref_models.Enumeration`
	__GongStructShape__000002_Default_Enumeration.ShowNbInstances = false
	__GongStructShape__000002_Default_Enumeration.NbInstances = 2
	__GongStructShape__000002_Default_Enumeration.Width = 240.000000
	__GongStructShape__000002_Default_Enumeration.Height = 78.000000
	__GongStructShape__000002_Default_Enumeration.IsSelected = false

	__GongStructShape__000003_Default_MaxInclusive.Name = `Default-MaxInclusive`

	//gong:ident [ref_models.MaxInclusive] comment added to overcome the problem with the comment map association
	__GongStructShape__000003_Default_MaxInclusive.Identifier = `ref_models.MaxInclusive`
	__GongStructShape__000003_Default_MaxInclusive.ShowNbInstances = false
	__GongStructShape__000003_Default_MaxInclusive.NbInstances = 0
	__GongStructShape__000003_Default_MaxInclusive.Width = 240.000000
	__GongStructShape__000003_Default_MaxInclusive.Height = 63.000000
	__GongStructShape__000003_Default_MaxInclusive.IsSelected = false

	__GongStructShape__000004_Default_MinInclusive.Name = `Default-MinInclusive`

	//gong:ident [ref_models.MinInclusive] comment added to overcome the problem with the comment map association
	__GongStructShape__000004_Default_MinInclusive.Identifier = `ref_models.MinInclusive`
	__GongStructShape__000004_Default_MinInclusive.ShowNbInstances = false
	__GongStructShape__000004_Default_MinInclusive.NbInstances = 0
	__GongStructShape__000004_Default_MinInclusive.Width = 240.000000
	__GongStructShape__000004_Default_MinInclusive.Height = 63.000000
	__GongStructShape__000004_Default_MinInclusive.IsSelected = false

	__GongStructShape__000005_Default_Restriction.Name = `Default-Restriction`

	//gong:ident [ref_models.Restriction] comment added to overcome the problem with the comment map association
	__GongStructShape__000005_Default_Restriction.Identifier = `ref_models.Restriction`
	__GongStructShape__000005_Default_Restriction.ShowNbInstances = false
	__GongStructShape__000005_Default_Restriction.NbInstances = 2
	__GongStructShape__000005_Default_Restriction.Width = 240.000000
	__GongStructShape__000005_Default_Restriction.Height = 78.000000
	__GongStructShape__000005_Default_Restriction.IsSelected = false

	__GongStructShape__000006_Default_Schema.Name = `Default-Schema`

	//gong:ident [ref_models.Schema] comment added to overcome the problem with the comment map association
	__GongStructShape__000006_Default_Schema.Identifier = `ref_models.Schema`
	__GongStructShape__000006_Default_Schema.ShowNbInstances = false
	__GongStructShape__000006_Default_Schema.NbInstances = 1
	__GongStructShape__000006_Default_Schema.Width = 240.000000
	__GongStructShape__000006_Default_Schema.Height = 170.000000
	__GongStructShape__000006_Default_Schema.IsSelected = false

	__GongStructShape__000007_Default_Sequence.Name = `Default-Sequence`

	//gong:ident [ref_models.Sequence] comment added to overcome the problem with the comment map association
	__GongStructShape__000007_Default_Sequence.Identifier = `ref_models.Sequence`
	__GongStructShape__000007_Default_Sequence.ShowNbInstances = false
	__GongStructShape__000007_Default_Sequence.NbInstances = 2
	__GongStructShape__000007_Default_Sequence.Width = 240.000000
	__GongStructShape__000007_Default_Sequence.Height = 63.000000
	__GongStructShape__000007_Default_Sequence.IsSelected = false

	__GongStructShape__000008_Default_SimpleType.Name = `Default-SimpleType`

	//gong:ident [ref_models.SimpleType] comment added to overcome the problem with the comment map association
	__GongStructShape__000008_Default_SimpleType.Identifier = `ref_models.SimpleType`
	__GongStructShape__000008_Default_SimpleType.ShowNbInstances = false
	__GongStructShape__000008_Default_SimpleType.NbInstances = 2
	__GongStructShape__000008_Default_SimpleType.Width = 287.000000
	__GongStructShape__000008_Default_SimpleType.Height = 63.000000
	__GongStructShape__000008_Default_SimpleType.IsSelected = false

	__Link__000000_ComplexType.Name = `ComplexType`

	//gong:ident [ref_models.Element.ComplexType] comment added to overcome the problem with the comment map association
	__Link__000000_ComplexType.Identifier = `ref_models.Element.ComplexType`

	//gong:ident [ref_models.ComplexType] comment added to overcome the problem with the comment map association
	__Link__000000_ComplexType.Fieldtypename = `ref_models.ComplexType`
	__Link__000000_ComplexType.FieldOffsetX = -50.000000
	__Link__000000_ComplexType.FieldOffsetY = -16.000000
	__Link__000000_ComplexType.TargetMultiplicity = models.ZERO_ONE
	__Link__000000_ComplexType.TargetMultiplicityOffsetX = -50.000000
	__Link__000000_ComplexType.TargetMultiplicityOffsetY = 16.000000
	__Link__000000_ComplexType.SourceMultiplicity = models.MANY
	__Link__000000_ComplexType.SourceMultiplicityOffsetX = 10.000000
	__Link__000000_ComplexType.SourceMultiplicityOffsetY = -50.000000
	__Link__000000_ComplexType.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_ComplexType.StartRatio = 0.500000
	__Link__000000_ComplexType.EndOrientation = models.ORIENTATION_VERTICAL
	__Link__000000_ComplexType.EndRatio = 0.134306
	__Link__000000_ComplexType.CornerOffsetRatio = 0.821751

	__Link__000001_ComplexTypes.Name = `ComplexTypes`

	//gong:ident [ref_models.Schema.ComplexTypes] comment added to overcome the problem with the comment map association
	__Link__000001_ComplexTypes.Identifier = `ref_models.Schema.ComplexTypes`

	//gong:ident [ref_models.ComplexType] comment added to overcome the problem with the comment map association
	__Link__000001_ComplexTypes.Fieldtypename = `ref_models.ComplexType`
	__Link__000001_ComplexTypes.FieldOffsetX = -50.000000
	__Link__000001_ComplexTypes.FieldOffsetY = -16.000000
	__Link__000001_ComplexTypes.TargetMultiplicity = models.MANY
	__Link__000001_ComplexTypes.TargetMultiplicityOffsetX = -50.000000
	__Link__000001_ComplexTypes.TargetMultiplicityOffsetY = 16.000000
	__Link__000001_ComplexTypes.SourceMultiplicity = models.MANY
	__Link__000001_ComplexTypes.SourceMultiplicityOffsetX = 10.000000
	__Link__000001_ComplexTypes.SourceMultiplicityOffsetY = -50.000000
	__Link__000001_ComplexTypes.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000001_ComplexTypes.StartRatio = 0.527845
	__Link__000001_ComplexTypes.EndOrientation = models.ORIENTATION_VERTICAL
	__Link__000001_ComplexTypes.EndRatio = 0.905139
	__Link__000001_ComplexTypes.CornerOffsetRatio = -0.086528

	__Link__000002_Elements.Name = `Elements`

	//gong:ident [ref_models.Sequence.Elements] comment added to overcome the problem with the comment map association
	__Link__000002_Elements.Identifier = `ref_models.Sequence.Elements`

	//gong:ident [ref_models.Element] comment added to overcome the problem with the comment map association
	__Link__000002_Elements.Fieldtypename = `ref_models.Element`
	__Link__000002_Elements.FieldOffsetX = -50.000000
	__Link__000002_Elements.FieldOffsetY = -16.000000
	__Link__000002_Elements.TargetMultiplicity = models.MANY
	__Link__000002_Elements.TargetMultiplicityOffsetX = -50.000000
	__Link__000002_Elements.TargetMultiplicityOffsetY = 16.000000
	__Link__000002_Elements.SourceMultiplicity = models.MANY
	__Link__000002_Elements.SourceMultiplicityOffsetX = 10.000000
	__Link__000002_Elements.SourceMultiplicityOffsetY = -50.000000
	__Link__000002_Elements.StartOrientation = models.ORIENTATION_VERTICAL
	__Link__000002_Elements.StartRatio = 0.338472
	__Link__000002_Elements.EndOrientation = models.ORIENTATION_VERTICAL
	__Link__000002_Elements.EndRatio = 0.871806
	__Link__000002_Elements.CornerOffsetRatio = 2.233867

	__Link__000003_Elements.Name = `Elements`

	//gong:ident [ref_models.Schema.Elements] comment added to overcome the problem with the comment map association
	__Link__000003_Elements.Identifier = `ref_models.Schema.Elements`

	//gong:ident [ref_models.Element] comment added to overcome the problem with the comment map association
	__Link__000003_Elements.Fieldtypename = `ref_models.Element`
	__Link__000003_Elements.FieldOffsetX = -50.000000
	__Link__000003_Elements.FieldOffsetY = -16.000000
	__Link__000003_Elements.TargetMultiplicity = models.MANY
	__Link__000003_Elements.TargetMultiplicityOffsetX = -50.000000
	__Link__000003_Elements.TargetMultiplicityOffsetY = 16.000000
	__Link__000003_Elements.SourceMultiplicity = models.MANY
	__Link__000003_Elements.SourceMultiplicityOffsetX = 10.000000
	__Link__000003_Elements.SourceMultiplicityOffsetY = -50.000000
	__Link__000003_Elements.StartOrientation = models.ORIENTATION_VERTICAL
	__Link__000003_Elements.StartRatio = 0.605139
	__Link__000003_Elements.EndOrientation = models.ORIENTATION_VERTICAL
	__Link__000003_Elements.EndRatio = 0.296806
	__Link__000003_Elements.CornerOffsetRatio = 1.840687

	__Link__000004_Enumerations.Name = `Enumerations`

	//gong:ident [ref_models.Restriction.Enumerations] comment added to overcome the problem with the comment map association
	__Link__000004_Enumerations.Identifier = `ref_models.Restriction.Enumerations`

	//gong:ident [ref_models.Enumeration] comment added to overcome the problem with the comment map association
	__Link__000004_Enumerations.Fieldtypename = `ref_models.Enumeration`
	__Link__000004_Enumerations.FieldOffsetX = -50.000000
	__Link__000004_Enumerations.FieldOffsetY = -16.000000
	__Link__000004_Enumerations.TargetMultiplicity = models.MANY
	__Link__000004_Enumerations.TargetMultiplicityOffsetX = -50.000000
	__Link__000004_Enumerations.TargetMultiplicityOffsetY = 16.000000
	__Link__000004_Enumerations.SourceMultiplicity = models.MANY
	__Link__000004_Enumerations.SourceMultiplicityOffsetX = 10.000000
	__Link__000004_Enumerations.SourceMultiplicityOffsetY = -50.000000
	__Link__000004_Enumerations.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000004_Enumerations.StartRatio = 0.500000
	__Link__000004_Enumerations.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000004_Enumerations.EndRatio = 0.500000
	__Link__000004_Enumerations.CornerOffsetRatio = 1.380000

	__Link__000005_MaxInclusive.Name = `MaxInclusive`

	//gong:ident [ref_models.Restriction.MaxInclusive] comment added to overcome the problem with the comment map association
	__Link__000005_MaxInclusive.Identifier = `ref_models.Restriction.MaxInclusive`

	//gong:ident [ref_models.MaxInclusive] comment added to overcome the problem with the comment map association
	__Link__000005_MaxInclusive.Fieldtypename = `ref_models.MaxInclusive`
	__Link__000005_MaxInclusive.FieldOffsetX = -50.000000
	__Link__000005_MaxInclusive.FieldOffsetY = -16.000000
	__Link__000005_MaxInclusive.TargetMultiplicity = models.ZERO_ONE
	__Link__000005_MaxInclusive.TargetMultiplicityOffsetX = -50.000000
	__Link__000005_MaxInclusive.TargetMultiplicityOffsetY = 16.000000
	__Link__000005_MaxInclusive.SourceMultiplicity = models.MANY
	__Link__000005_MaxInclusive.SourceMultiplicityOffsetX = 10.000000
	__Link__000005_MaxInclusive.SourceMultiplicityOffsetY = -50.000000
	__Link__000005_MaxInclusive.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000005_MaxInclusive.StartRatio = 0.500000
	__Link__000005_MaxInclusive.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000005_MaxInclusive.EndRatio = 0.500000
	__Link__000005_MaxInclusive.CornerOffsetRatio = 1.380000

	__Link__000006_MinInclusive.Name = `MinInclusive`

	//gong:ident [ref_models.Restriction.MinInclusive] comment added to overcome the problem with the comment map association
	__Link__000006_MinInclusive.Identifier = `ref_models.Restriction.MinInclusive`

	//gong:ident [ref_models.MinInclusive] comment added to overcome the problem with the comment map association
	__Link__000006_MinInclusive.Fieldtypename = `ref_models.MinInclusive`
	__Link__000006_MinInclusive.FieldOffsetX = -50.000000
	__Link__000006_MinInclusive.FieldOffsetY = -16.000000
	__Link__000006_MinInclusive.TargetMultiplicity = models.ZERO_ONE
	__Link__000006_MinInclusive.TargetMultiplicityOffsetX = -50.000000
	__Link__000006_MinInclusive.TargetMultiplicityOffsetY = 16.000000
	__Link__000006_MinInclusive.SourceMultiplicity = models.MANY
	__Link__000006_MinInclusive.SourceMultiplicityOffsetX = 10.000000
	__Link__000006_MinInclusive.SourceMultiplicityOffsetY = -50.000000
	__Link__000006_MinInclusive.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000006_MinInclusive.StartRatio = 0.500000
	__Link__000006_MinInclusive.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000006_MinInclusive.EndRatio = 0.500000
	__Link__000006_MinInclusive.CornerOffsetRatio = 1.380000

	__Link__000007_Restriction.Name = `Restriction`

	//gong:ident [ref_models.SimpleType.Restriction] comment added to overcome the problem with the comment map association
	__Link__000007_Restriction.Identifier = `ref_models.SimpleType.Restriction`

	//gong:ident [ref_models.Restriction] comment added to overcome the problem with the comment map association
	__Link__000007_Restriction.Fieldtypename = `ref_models.Restriction`
	__Link__000007_Restriction.FieldOffsetX = -50.000000
	__Link__000007_Restriction.FieldOffsetY = -16.000000
	__Link__000007_Restriction.TargetMultiplicity = models.ZERO_ONE
	__Link__000007_Restriction.TargetMultiplicityOffsetX = -50.000000
	__Link__000007_Restriction.TargetMultiplicityOffsetY = 16.000000
	__Link__000007_Restriction.SourceMultiplicity = models.MANY
	__Link__000007_Restriction.SourceMultiplicityOffsetX = 10.000000
	__Link__000007_Restriction.SourceMultiplicityOffsetY = -50.000000
	__Link__000007_Restriction.StartOrientation = models.ORIENTATION_VERTICAL
	__Link__000007_Restriction.StartRatio = 0.223810
	__Link__000007_Restriction.EndOrientation = models.ORIENTATION_VERTICAL
	__Link__000007_Restriction.EndRatio = 0.196806
	__Link__000007_Restriction.CornerOffsetRatio = -0.495764

	__Link__000008_Sequence.Name = `Sequence`

	//gong:ident [ref_models.ComplexType.Sequence] comment added to overcome the problem with the comment map association
	__Link__000008_Sequence.Identifier = `ref_models.ComplexType.Sequence`

	//gong:ident [ref_models.Sequence] comment added to overcome the problem with the comment map association
	__Link__000008_Sequence.Fieldtypename = `ref_models.Sequence`
	__Link__000008_Sequence.FieldOffsetX = -50.000000
	__Link__000008_Sequence.FieldOffsetY = -16.000000
	__Link__000008_Sequence.TargetMultiplicity = models.ZERO_ONE
	__Link__000008_Sequence.TargetMultiplicityOffsetX = -50.000000
	__Link__000008_Sequence.TargetMultiplicityOffsetY = 16.000000
	__Link__000008_Sequence.SourceMultiplicity = models.MANY
	__Link__000008_Sequence.SourceMultiplicityOffsetX = 10.000000
	__Link__000008_Sequence.SourceMultiplicityOffsetY = -50.000000
	__Link__000008_Sequence.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000008_Sequence.StartRatio = 0.535455
	__Link__000008_Sequence.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000008_Sequence.EndRatio = 0.500000
	__Link__000008_Sequence.CornerOffsetRatio = 1.380000

	__Link__000009_SimpleType.Name = `SimpleType`

	//gong:ident [ref_models.Element.SimpleType] comment added to overcome the problem with the comment map association
	__Link__000009_SimpleType.Identifier = `ref_models.Element.SimpleType`

	//gong:ident [ref_models.SimpleType] comment added to overcome the problem with the comment map association
	__Link__000009_SimpleType.Fieldtypename = `ref_models.SimpleType`
	__Link__000009_SimpleType.FieldOffsetX = -50.000000
	__Link__000009_SimpleType.FieldOffsetY = -16.000000
	__Link__000009_SimpleType.TargetMultiplicity = models.ZERO_ONE
	__Link__000009_SimpleType.TargetMultiplicityOffsetX = -50.000000
	__Link__000009_SimpleType.TargetMultiplicityOffsetY = 16.000000
	__Link__000009_SimpleType.SourceMultiplicity = models.MANY
	__Link__000009_SimpleType.SourceMultiplicityOffsetX = 10.000000
	__Link__000009_SimpleType.SourceMultiplicityOffsetY = -50.000000
	__Link__000009_SimpleType.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000009_SimpleType.StartRatio = 0.431680
	__Link__000009_SimpleType.EndOrientation = models.ORIENTATION_VERTICAL
	__Link__000009_SimpleType.EndRatio = 0.155139
	__Link__000009_SimpleType.CornerOffsetRatio = 0.233098

	__Link__000010_SimpleTypes.Name = `SimpleTypes`

	//gong:ident [ref_models.Schema.SimpleTypes] comment added to overcome the problem with the comment map association
	__Link__000010_SimpleTypes.Identifier = `ref_models.Schema.SimpleTypes`

	//gong:ident [ref_models.SimpleType] comment added to overcome the problem with the comment map association
	__Link__000010_SimpleTypes.Fieldtypename = `ref_models.SimpleType`
	__Link__000010_SimpleTypes.FieldOffsetX = -50.000000
	__Link__000010_SimpleTypes.FieldOffsetY = -16.000000
	__Link__000010_SimpleTypes.TargetMultiplicity = models.MANY
	__Link__000010_SimpleTypes.TargetMultiplicityOffsetX = -50.000000
	__Link__000010_SimpleTypes.TargetMultiplicityOffsetY = 16.000000
	__Link__000010_SimpleTypes.SourceMultiplicity = models.MANY
	__Link__000010_SimpleTypes.SourceMultiplicityOffsetX = 10.000000
	__Link__000010_SimpleTypes.SourceMultiplicityOffsetY = -50.000000
	__Link__000010_SimpleTypes.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000010_SimpleTypes.StartRatio = 0.270058
	__Link__000010_SimpleTypes.EndOrientation = models.ORIENTATION_VERTICAL
	__Link__000010_SimpleTypes.EndRatio = 0.867639
	__Link__000010_SimpleTypes.CornerOffsetRatio = -0.496291

	__Position__000000_Pos_Default_ComplexType.X = 503.000000
	__Position__000000_Pos_Default_ComplexType.Y = 712.000000
	__Position__000000_Pos_Default_ComplexType.Name = `Pos-Default-ComplexType`

	__Position__000001_Pos_Default_Element.X = 23.000000
	__Position__000001_Pos_Default_Element.Y = 479.000000
	__Position__000001_Pos_Default_Element.Name = `Pos-Default-Element`

	__Position__000002_Pos_Default_Enumeration.X = 1009.000061
	__Position__000002_Pos_Default_Enumeration.Y = 75.000000
	__Position__000002_Pos_Default_Enumeration.Name = `Pos-Default-Enumeration`

	__Position__000003_Pos_Default_MaxInclusive.X = 1012.000061
	__Position__000003_Pos_Default_MaxInclusive.Y = 298.000000
	__Position__000003_Pos_Default_MaxInclusive.Name = `Pos-Default-MaxInclusive`

	__Position__000004_Pos_Default_MinInclusive.X = 1009.000061
	__Position__000004_Pos_Default_MinInclusive.Y = 194.000000
	__Position__000004_Pos_Default_MinInclusive.Name = `Pos-Default-MinInclusive`

	__Position__000005_Pos_Default_Restriction.X = 511.000000
	__Position__000005_Pos_Default_Restriction.Y = 66.000000
	__Position__000005_Pos_Default_Restriction.Name = `Pos-Default-Restriction`

	__Position__000006_Pos_Default_Schema.X = 1386.000061
	__Position__000006_Pos_Default_Schema.Y = 554.000000
	__Position__000006_Pos_Default_Schema.Name = `Pos-Default-Schema`

	__Position__000007_Pos_Default_Sequence.X = 1001.000061
	__Position__000007_Pos_Default_Sequence.Y = 702.000000
	__Position__000007_Pos_Default_Sequence.Name = `Pos-Default-Sequence`

	__Position__000008_Pos_Default_SimpleType.X = 498.000000
	__Position__000008_Pos_Default_SimpleType.Y = 474.000000
	__Position__000008_Pos_Default_SimpleType.Name = `Pos-Default-SimpleType`

	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_ComplexType_and_Default_Sequence.X = 664.500000
	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_ComplexType_and_Default_Sequence.Y = 268.500000
	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_ComplexType_and_Default_Sequence.Name = `Verticle in class diagram Default in middle between Default-ComplexType and Default-Sequence`

	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Element_and_Default_ComplexType.X = 635.000000
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Element_and_Default_ComplexType.Y = 529.500000
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Element_and_Default_ComplexType.Name = `Verticle in class diagram Default in middle between Default-Element and Default-ComplexType`

	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_Element_and_Default_SimpleType.X = 632.500000
	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_Element_and_Default_SimpleType.Y = 410.500000
	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_Element_and_Default_SimpleType.Name = `Verticle in class diagram Default in middle between Default-Element and Default-SimpleType`

	__Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_Restriction_and_Default_Enumeration.X = 883.000031
	__Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_Restriction_and_Default_Enumeration.Y = 205.000000
	__Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_Restriction_and_Default_Enumeration.Name = `Verticle in class diagram Default in middle between Default-Restriction and Default-Enumeration`

	__Vertice__000004_Verticle_in_class_diagram_Default_in_middle_between_Default_Restriction_and_Default_MaxInclusive.X = 1121.500030
	__Vertice__000004_Verticle_in_class_diagram_Default_in_middle_between_Default_Restriction_and_Default_MaxInclusive.Y = 221.000000
	__Vertice__000004_Verticle_in_class_diagram_Default_in_middle_between_Default_Restriction_and_Default_MaxInclusive.Name = `Verticle in class diagram Default in middle between Default-Restriction and Default-MaxInclusive`

	__Vertice__000005_Verticle_in_class_diagram_Default_in_middle_between_Default_Restriction_and_Default_MinInclusive.X = 1120.000030
	__Vertice__000005_Verticle_in_class_diagram_Default_in_middle_between_Default_Restriction_and_Default_MinInclusive.Y = 169.000000
	__Vertice__000005_Verticle_in_class_diagram_Default_in_middle_between_Default_Restriction_and_Default_MinInclusive.Name = `Verticle in class diagram Default in middle between Default-Restriction and Default-MinInclusive`

	__Vertice__000006_Verticle_in_class_diagram_Default_in_middle_between_Default_Schema_and_Default_ComplexType.X = 536.000000
	__Vertice__000006_Verticle_in_class_diagram_Default_in_middle_between_Default_Schema_and_Default_ComplexType.Y = 239.000000
	__Vertice__000006_Verticle_in_class_diagram_Default_in_middle_between_Default_Schema_and_Default_ComplexType.Name = `Verticle in class diagram Default in middle between Default-Schema and Default-ComplexType`

	__Vertice__000007_Verticle_in_class_diagram_Default_in_middle_between_Default_Schema_and_Default_Element.X = 1076.500031
	__Vertice__000007_Verticle_in_class_diagram_Default_in_middle_between_Default_Schema_and_Default_Element.Y = 427.500000
	__Vertice__000007_Verticle_in_class_diagram_Default_in_middle_between_Default_Schema_and_Default_Element.Name = `Verticle in class diagram Default in middle between Default-Schema and Default-Element`

	__Vertice__000008_Verticle_in_class_diagram_Default_in_middle_between_Default_Schema_and_Default_SimpleType.X = 528.500000
	__Vertice__000008_Verticle_in_class_diagram_Default_in_middle_between_Default_Schema_and_Default_SimpleType.Y = 177.500000
	__Vertice__000008_Verticle_in_class_diagram_Default_in_middle_between_Default_Schema_and_Default_SimpleType.Name = `Verticle in class diagram Default in middle between Default-Schema and Default-SimpleType`

	__Vertice__000009_Verticle_in_class_diagram_Default_in_middle_between_Default_Sequence_and_Default_Element.X = 1119.500031
	__Vertice__000009_Verticle_in_class_diagram_Default_in_middle_between_Default_Sequence_and_Default_Element.Y = 301.500000
	__Vertice__000009_Verticle_in_class_diagram_Default_in_middle_between_Default_Sequence_and_Default_Element.Name = `Verticle in class diagram Default in middle between Default-Sequence and Default-Element`

	__Vertice__000010_Verticle_in_class_diagram_Default_in_middle_between_Default_SimpleType_and_Default_Restriction.X = 1126.000031
	__Vertice__000010_Verticle_in_class_diagram_Default_in_middle_between_Default_SimpleType_and_Default_Restriction.Y = 320.000000
	__Vertice__000010_Verticle_in_class_diagram_Default_in_middle_between_Default_SimpleType_and_Default_Restriction.Name = `Verticle in class diagram Default in middle between Default-SimpleType and Default-Restriction`

	// Setup of pointers
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000006_Default_Schema)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000001_Default_Element)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000000_Default_ComplexType)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000008_Default_SimpleType)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000005_Default_Restriction)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000007_Default_Sequence)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000002_Default_Enumeration)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000003_Default_MaxInclusive)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000004_Default_MinInclusive)
	__GongStructShape__000000_Default_ComplexType.Position = __Position__000000_Pos_Default_ComplexType
	__GongStructShape__000000_Default_ComplexType.Links = append(__GongStructShape__000000_Default_ComplexType.Links, __Link__000008_Sequence)
	__GongStructShape__000001_Default_Element.Position = __Position__000001_Pos_Default_Element
	__GongStructShape__000001_Default_Element.Fields = append(__GongStructShape__000001_Default_Element.Fields, __Field__000001_Type)
	__GongStructShape__000001_Default_Element.Links = append(__GongStructShape__000001_Default_Element.Links, __Link__000009_SimpleType)
	__GongStructShape__000001_Default_Element.Links = append(__GongStructShape__000001_Default_Element.Links, __Link__000000_ComplexType)
	__GongStructShape__000002_Default_Enumeration.Position = __Position__000002_Pos_Default_Enumeration
	__GongStructShape__000002_Default_Enumeration.Fields = append(__GongStructShape__000002_Default_Enumeration.Fields, __Field__000002_Value)
	__GongStructShape__000003_Default_MaxInclusive.Position = __Position__000003_Pos_Default_MaxInclusive
	__GongStructShape__000004_Default_MinInclusive.Position = __Position__000004_Pos_Default_MinInclusive
	__GongStructShape__000005_Default_Restriction.Position = __Position__000005_Pos_Default_Restriction
	__GongStructShape__000005_Default_Restriction.Fields = append(__GongStructShape__000005_Default_Restriction.Fields, __Field__000000_Base)
	__GongStructShape__000005_Default_Restriction.Links = append(__GongStructShape__000005_Default_Restriction.Links, __Link__000004_Enumerations)
	__GongStructShape__000005_Default_Restriction.Links = append(__GongStructShape__000005_Default_Restriction.Links, __Link__000006_MinInclusive)
	__GongStructShape__000005_Default_Restriction.Links = append(__GongStructShape__000005_Default_Restriction.Links, __Link__000005_MaxInclusive)
	__GongStructShape__000006_Default_Schema.Position = __Position__000006_Pos_Default_Schema
	__GongStructShape__000006_Default_Schema.Links = append(__GongStructShape__000006_Default_Schema.Links, __Link__000010_SimpleTypes)
	__GongStructShape__000006_Default_Schema.Links = append(__GongStructShape__000006_Default_Schema.Links, __Link__000001_ComplexTypes)
	__GongStructShape__000006_Default_Schema.Links = append(__GongStructShape__000006_Default_Schema.Links, __Link__000003_Elements)
	__GongStructShape__000007_Default_Sequence.Position = __Position__000007_Pos_Default_Sequence
	__GongStructShape__000007_Default_Sequence.Links = append(__GongStructShape__000007_Default_Sequence.Links, __Link__000002_Elements)
	__GongStructShape__000008_Default_SimpleType.Position = __Position__000008_Pos_Default_SimpleType
	__GongStructShape__000008_Default_SimpleType.Links = append(__GongStructShape__000008_Default_SimpleType.Links, __Link__000007_Restriction)
	__Link__000000_ComplexType.Middlevertice = __Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Element_and_Default_ComplexType
	__Link__000001_ComplexTypes.Middlevertice = __Vertice__000006_Verticle_in_class_diagram_Default_in_middle_between_Default_Schema_and_Default_ComplexType
	__Link__000002_Elements.Middlevertice = __Vertice__000009_Verticle_in_class_diagram_Default_in_middle_between_Default_Sequence_and_Default_Element
	__Link__000003_Elements.Middlevertice = __Vertice__000007_Verticle_in_class_diagram_Default_in_middle_between_Default_Schema_and_Default_Element
	__Link__000004_Enumerations.Middlevertice = __Vertice__000003_Verticle_in_class_diagram_Default_in_middle_between_Default_Restriction_and_Default_Enumeration
	__Link__000005_MaxInclusive.Middlevertice = __Vertice__000004_Verticle_in_class_diagram_Default_in_middle_between_Default_Restriction_and_Default_MaxInclusive
	__Link__000006_MinInclusive.Middlevertice = __Vertice__000005_Verticle_in_class_diagram_Default_in_middle_between_Default_Restriction_and_Default_MinInclusive
	__Link__000007_Restriction.Middlevertice = __Vertice__000010_Verticle_in_class_diagram_Default_in_middle_between_Default_SimpleType_and_Default_Restriction
	__Link__000008_Sequence.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_ComplexType_and_Default_Sequence
	__Link__000009_SimpleType.Middlevertice = __Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_Element_and_Default_SimpleType
	__Link__000010_SimpleTypes.Middlevertice = __Vertice__000008_Verticle_in_class_diagram_Default_in_middle_between_Default_Schema_and_Default_SimpleType
}
