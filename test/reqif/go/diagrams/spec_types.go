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

	__Classdiagram__000000_spec_types := (&models.Classdiagram{Name: `spec_types`}).Stage(stage)

	__GongStructShape__000000_spec_types_SPEC_OBJECT_TYPE := (&models.GongStructShape{Name: `spec_types-SPEC_OBJECT_TYPE`}).Stage(stage)

	__Position__000000_Pos_spec_types_SPEC_OBJECT_TYPE := (&models.Position{Name: `Pos-spec_types-SPEC_OBJECT_TYPE`}).Stage(stage)

	// Setup of values

	__Classdiagram__000000_spec_types.Name = `spec_types`
	__Classdiagram__000000_spec_types.IsInDrawMode = false

	__GongStructShape__000000_spec_types_SPEC_OBJECT_TYPE.Name = `spec_types-SPEC_OBJECT_TYPE`

	//gong:ident [ref_models.SPEC_OBJECT_TYPE] comment added to overcome the problem with the comment map association
	__GongStructShape__000000_spec_types_SPEC_OBJECT_TYPE.Identifier = `ref_models.SPEC_OBJECT_TYPE`
	__GongStructShape__000000_spec_types_SPEC_OBJECT_TYPE.ShowNbInstances = false
	__GongStructShape__000000_spec_types_SPEC_OBJECT_TYPE.NbInstances = 0
	__GongStructShape__000000_spec_types_SPEC_OBJECT_TYPE.Width = 240.000000
	__GongStructShape__000000_spec_types_SPEC_OBJECT_TYPE.Height = 63.000000
	__GongStructShape__000000_spec_types_SPEC_OBJECT_TYPE.IsSelected = false

	__Position__000000_Pos_spec_types_SPEC_OBJECT_TYPE.X = 10.000000
	__Position__000000_Pos_spec_types_SPEC_OBJECT_TYPE.Y = 22.000000
	__Position__000000_Pos_spec_types_SPEC_OBJECT_TYPE.Name = `Pos-spec_types-SPEC_OBJECT_TYPE`

	// Setup of pointers
	__Classdiagram__000000_spec_types.GongStructShapes = append(__Classdiagram__000000_spec_types.GongStructShapes, __GongStructShape__000000_spec_types_SPEC_OBJECT_TYPE)
	__GongStructShape__000000_spec_types_SPEC_OBJECT_TYPE.Position = __Position__000000_Pos_spec_types_SPEC_OBJECT_TYPE
}
