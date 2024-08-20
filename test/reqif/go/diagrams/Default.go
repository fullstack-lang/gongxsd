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

	__Field__000000_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000001_Name := (&models.Field{Name: `Name`}).Stage(stage)

	__GongStructShape__000000_Default_REQ_IF := (&models.GongStructShape{Name: `Default-REQ_IF`}).Stage(stage)
	__GongStructShape__000001_Default_REQ_IF_CONTENT := (&models.GongStructShape{Name: `Default-REQ_IF_CONTENT`}).Stage(stage)
	__GongStructShape__000002_Default_REQ_IF_HEADER := (&models.GongStructShape{Name: `Default-REQ_IF_HEADER`}).Stage(stage)

	__Position__000000_Pos_Default_REQ_IF := (&models.Position{Name: `Pos-Default-REQ_IF`}).Stage(stage)
	__Position__000001_Pos_Default_REQ_IF_CONTENT := (&models.Position{Name: `Pos-Default-REQ_IF_CONTENT`}).Stage(stage)
	__Position__000002_Pos_Default_REQ_IF_HEADER := (&models.Position{Name: `Pos-Default-REQ_IF_HEADER`}).Stage(stage)

	// Setup of values

	__Classdiagram__000000_Default.Name = `Default`
	__Classdiagram__000000_Default.IsInDrawMode = false

	__Field__000000_Name.Name = `Name`

	//gong:ident [ref_models.REQ_IF.Name] comment added to overcome the problem with the comment map association
	__Field__000000_Name.Identifier = `ref_models.REQ_IF.Name`
	__Field__000000_Name.FieldTypeAsString = ``
	__Field__000000_Name.Structname = `REQ_IF`
	__Field__000000_Name.Fieldtypename = `string`

	__Field__000001_Name.Name = `Name`

	//gong:ident [ref_models.REQ_IF_CONTENT.Name] comment added to overcome the problem with the comment map association
	__Field__000001_Name.Identifier = `ref_models.REQ_IF_CONTENT.Name`
	__Field__000001_Name.FieldTypeAsString = ``
	__Field__000001_Name.Structname = `REQ_IF_CONTENT`
	__Field__000001_Name.Fieldtypename = `string`

	__GongStructShape__000000_Default_REQ_IF.Name = `Default-REQ_IF`

	//gong:ident [ref_models.REQ_IF] comment added to overcome the problem with the comment map association
	__GongStructShape__000000_Default_REQ_IF.Identifier = `ref_models.REQ_IF`
	__GongStructShape__000000_Default_REQ_IF.ShowNbInstances = false
	__GongStructShape__000000_Default_REQ_IF.NbInstances = 0
	__GongStructShape__000000_Default_REQ_IF.Width = 240.000000
	__GongStructShape__000000_Default_REQ_IF.Height = 78.000000
	__GongStructShape__000000_Default_REQ_IF.IsSelected = false

	__GongStructShape__000001_Default_REQ_IF_CONTENT.Name = `Default-REQ_IF_CONTENT`

	//gong:ident [ref_models.REQ_IF_CONTENT] comment added to overcome the problem with the comment map association
	__GongStructShape__000001_Default_REQ_IF_CONTENT.Identifier = `ref_models.REQ_IF_CONTENT`
	__GongStructShape__000001_Default_REQ_IF_CONTENT.ShowNbInstances = false
	__GongStructShape__000001_Default_REQ_IF_CONTENT.NbInstances = 0
	__GongStructShape__000001_Default_REQ_IF_CONTENT.Width = 240.000000
	__GongStructShape__000001_Default_REQ_IF_CONTENT.Height = 78.000000
	__GongStructShape__000001_Default_REQ_IF_CONTENT.IsSelected = false

	__GongStructShape__000002_Default_REQ_IF_HEADER.Name = `Default-REQ_IF_HEADER`

	//gong:ident [ref_models.REQ_IF_HEADER] comment added to overcome the problem with the comment map association
	__GongStructShape__000002_Default_REQ_IF_HEADER.Identifier = `ref_models.REQ_IF_HEADER`
	__GongStructShape__000002_Default_REQ_IF_HEADER.ShowNbInstances = false
	__GongStructShape__000002_Default_REQ_IF_HEADER.NbInstances = 0
	__GongStructShape__000002_Default_REQ_IF_HEADER.Width = 240.000000
	__GongStructShape__000002_Default_REQ_IF_HEADER.Height = 63.000000
	__GongStructShape__000002_Default_REQ_IF_HEADER.IsSelected = false

	__Position__000000_Pos_Default_REQ_IF.X = 89.000000
	__Position__000000_Pos_Default_REQ_IF.Y = 65.000000
	__Position__000000_Pos_Default_REQ_IF.Name = `Pos-Default-REQ_IF`

	__Position__000001_Pos_Default_REQ_IF_CONTENT.X = 86.000000
	__Position__000001_Pos_Default_REQ_IF_CONTENT.Y = 221.000000
	__Position__000001_Pos_Default_REQ_IF_CONTENT.Name = `Pos-Default-REQ_IF_CONTENT`

	__Position__000002_Pos_Default_REQ_IF_HEADER.X = 482.000000
	__Position__000002_Pos_Default_REQ_IF_HEADER.Y = 76.000000
	__Position__000002_Pos_Default_REQ_IF_HEADER.Name = `Pos-Default-REQ_IF_HEADER`

	// Setup of pointers
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000000_Default_REQ_IF)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000001_Default_REQ_IF_CONTENT)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000002_Default_REQ_IF_HEADER)
	__GongStructShape__000000_Default_REQ_IF.Position = __Position__000000_Pos_Default_REQ_IF
	__GongStructShape__000000_Default_REQ_IF.Fields = append(__GongStructShape__000000_Default_REQ_IF.Fields, __Field__000000_Name)
	__GongStructShape__000001_Default_REQ_IF_CONTENT.Position = __Position__000001_Pos_Default_REQ_IF_CONTENT
	__GongStructShape__000001_Default_REQ_IF_CONTENT.Fields = append(__GongStructShape__000001_Default_REQ_IF_CONTENT.Fields, __Field__000001_Name)
	__GongStructShape__000002_Default_REQ_IF_HEADER.Position = __Position__000002_Pos_Default_REQ_IF_HEADER
}
