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

	__GongStructShape__000000_Default_A_THE_HEADER := (&models.GongStructShape{Name: `Default-A_THE_HEADER`}).Stage(stage)
	__GongStructShape__000001_Default_REQ_IF := (&models.GongStructShape{Name: `Default-REQ_IF`}).Stage(stage)

	__Link__000000_THE_HEADER := (&models.Link{Name: `THE_HEADER`}).Stage(stage)

	__Position__000000_Pos_Default_A_THE_HEADER := (&models.Position{Name: `Pos-Default-A_THE_HEADER`}).Stage(stage)
	__Position__000001_Pos_Default_REQ_IF := (&models.Position{Name: `Pos-Default-REQ_IF`}).Stage(stage)

	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_REQ_IF_and_Default_A_THE_HEADER := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-REQ_IF and Default-A_THE_HEADER`}).Stage(stage)

	// Setup of values

	__Classdiagram__000000_Default.Name = `Default`
	__Classdiagram__000000_Default.IsInDrawMode = false

	__GongStructShape__000000_Default_A_THE_HEADER.Name = `Default-A_THE_HEADER`

	//gong:ident [ref_models.A_THE_HEADER] comment added to overcome the problem with the comment map association
	__GongStructShape__000000_Default_A_THE_HEADER.Identifier = `ref_models.A_THE_HEADER`
	__GongStructShape__000000_Default_A_THE_HEADER.ShowNbInstances = false
	__GongStructShape__000000_Default_A_THE_HEADER.NbInstances = 0
	__GongStructShape__000000_Default_A_THE_HEADER.Width = 240.000000
	__GongStructShape__000000_Default_A_THE_HEADER.Height = 63.000000
	__GongStructShape__000000_Default_A_THE_HEADER.IsSelected = false

	__GongStructShape__000001_Default_REQ_IF.Name = `Default-REQ_IF`

	//gong:ident [ref_models.REQ_IF] comment added to overcome the problem with the comment map association
	__GongStructShape__000001_Default_REQ_IF.Identifier = `ref_models.REQ_IF`
	__GongStructShape__000001_Default_REQ_IF.ShowNbInstances = false
	__GongStructShape__000001_Default_REQ_IF.NbInstances = 0
	__GongStructShape__000001_Default_REQ_IF.Width = 240.000000
	__GongStructShape__000001_Default_REQ_IF.Height = 63.000000
	__GongStructShape__000001_Default_REQ_IF.IsSelected = false

	__Link__000000_THE_HEADER.Name = `THE_HEADER`

	//gong:ident [ref_models.REQ_IF.THE_HEADER] comment added to overcome the problem with the comment map association
	__Link__000000_THE_HEADER.Identifier = `ref_models.REQ_IF.THE_HEADER`

	//gong:ident [ref_models.A_THE_HEADER] comment added to overcome the problem with the comment map association
	__Link__000000_THE_HEADER.Fieldtypename = `ref_models.A_THE_HEADER`
	__Link__000000_THE_HEADER.FieldOffsetX = -50.000000
	__Link__000000_THE_HEADER.FieldOffsetY = -16.000000
	__Link__000000_THE_HEADER.TargetMultiplicity = models.MANY
	__Link__000000_THE_HEADER.TargetMultiplicityOffsetX = -50.000000
	__Link__000000_THE_HEADER.TargetMultiplicityOffsetY = 16.000000
	__Link__000000_THE_HEADER.SourceMultiplicity = models.MANY
	__Link__000000_THE_HEADER.SourceMultiplicityOffsetX = 10.000000
	__Link__000000_THE_HEADER.SourceMultiplicityOffsetY = -50.000000
	__Link__000000_THE_HEADER.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_THE_HEADER.StartRatio = 0.500000
	__Link__000000_THE_HEADER.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_THE_HEADER.EndRatio = 0.500000
	__Link__000000_THE_HEADER.CornerOffsetRatio = 1.380000

	__Position__000000_Pos_Default_A_THE_HEADER.X = 108.000000
	__Position__000000_Pos_Default_A_THE_HEADER.Y = 91.000000
	__Position__000000_Pos_Default_A_THE_HEADER.Name = `Pos-Default-A_THE_HEADER`

	__Position__000001_Pos_Default_REQ_IF.X = 57.000000
	__Position__000001_Pos_Default_REQ_IF.Y = 25.000000
	__Position__000001_Pos_Default_REQ_IF.Name = `Pos-Default-REQ_IF`

	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_REQ_IF_and_Default_A_THE_HEADER.X = 442.500000
	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_REQ_IF_and_Default_A_THE_HEADER.Y = 89.500000
	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_REQ_IF_and_Default_A_THE_HEADER.Name = `Verticle in class diagram Default in middle between Default-REQ_IF and Default-A_THE_HEADER`

	// Setup of pointers
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000001_Default_REQ_IF)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000000_Default_A_THE_HEADER)
	__GongStructShape__000000_Default_A_THE_HEADER.Position = __Position__000000_Pos_Default_A_THE_HEADER
	__GongStructShape__000001_Default_REQ_IF.Position = __Position__000001_Pos_Default_REQ_IF
	__GongStructShape__000001_Default_REQ_IF.Links = append(__GongStructShape__000001_Default_REQ_IF.Links, __Link__000000_THE_HEADER)
	__Link__000000_THE_HEADER.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_REQ_IF_and_Default_A_THE_HEADER
}
