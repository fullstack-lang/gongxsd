package diagrams

import (
	"time"

	"github.com/fullstack-lang/gongdoc/go/models"

	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gongxsd/test/books/go/models"
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

	__GongStructShape__000000_Default_BookDetailsGroup := (&models.GongStructShape{Name: `Default-BookDetailsGroup`}).Stage(stage)
	__GongStructShape__000001_Default_BookType := (&models.GongStructShape{Name: `Default-BookType`}).Stage(stage)
	__GongStructShape__000002_Default_Books := (&models.GongStructShape{Name: `Default-Books`}).Stage(stage)

	__Position__000000_Pos_Default_BookDetailsGroup := (&models.Position{Name: `Pos-Default-BookDetailsGroup`}).Stage(stage)
	__Position__000001_Pos_Default_BookType := (&models.Position{Name: `Pos-Default-BookType`}).Stage(stage)
	__Position__000002_Pos_Default_Books := (&models.Position{Name: `Pos-Default-Books`}).Stage(stage)

	// Setup of values

	__Classdiagram__000000_Default.Name = `Default`
	__Classdiagram__000000_Default.IsInDrawMode = false

	__Field__000000_Name.Name = `Name`

	//gong:ident [ref_models.Books.Name] comment added to overcome the problem with the comment map association
	__Field__000000_Name.Identifier = `ref_models.Books.Name`
	__Field__000000_Name.FieldTypeAsString = ``
	__Field__000000_Name.Structname = `Books`
	__Field__000000_Name.Fieldtypename = `string`

	__GongStructShape__000000_Default_BookDetailsGroup.Name = `Default-BookDetailsGroup`

	//gong:ident [ref_models.BookDetailsGroup] comment added to overcome the problem with the comment map association
	__GongStructShape__000000_Default_BookDetailsGroup.Identifier = `ref_models.BookDetailsGroup`
	__GongStructShape__000000_Default_BookDetailsGroup.ShowNbInstances = false
	__GongStructShape__000000_Default_BookDetailsGroup.NbInstances = 0
	__GongStructShape__000000_Default_BookDetailsGroup.Width = 240.000000
	__GongStructShape__000000_Default_BookDetailsGroup.Height = 63.000000
	__GongStructShape__000000_Default_BookDetailsGroup.IsSelected = false

	__GongStructShape__000001_Default_BookType.Name = `Default-BookType`

	//gong:ident [ref_models.BookType] comment added to overcome the problem with the comment map association
	__GongStructShape__000001_Default_BookType.Identifier = `ref_models.BookType`
	__GongStructShape__000001_Default_BookType.ShowNbInstances = false
	__GongStructShape__000001_Default_BookType.NbInstances = 0
	__GongStructShape__000001_Default_BookType.Width = 240.000000
	__GongStructShape__000001_Default_BookType.Height = 63.000000
	__GongStructShape__000001_Default_BookType.IsSelected = false

	__GongStructShape__000002_Default_Books.Name = `Default-Books`

	//gong:ident [ref_models.Books] comment added to overcome the problem with the comment map association
	__GongStructShape__000002_Default_Books.Identifier = `ref_models.Books`
	__GongStructShape__000002_Default_Books.ShowNbInstances = false
	__GongStructShape__000002_Default_Books.NbInstances = 0
	__GongStructShape__000002_Default_Books.Width = 240.000000
	__GongStructShape__000002_Default_Books.Height = 78.000000
	__GongStructShape__000002_Default_Books.IsSelected = false

	__Position__000000_Pos_Default_BookDetailsGroup.X = 100.000000
	__Position__000000_Pos_Default_BookDetailsGroup.Y = 36.000000
	__Position__000000_Pos_Default_BookDetailsGroup.Name = `Pos-Default-BookDetailsGroup`

	__Position__000001_Pos_Default_BookType.X = 92.000000
	__Position__000001_Pos_Default_BookType.Y = 48.000000
	__Position__000001_Pos_Default_BookType.Name = `Pos-Default-BookType`

	__Position__000002_Pos_Default_Books.X = 28.000000
	__Position__000002_Pos_Default_Books.Y = 49.000000
	__Position__000002_Pos_Default_Books.Name = `Pos-Default-Books`

	// Setup of pointers
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000002_Default_Books)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000001_Default_BookType)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000000_Default_BookDetailsGroup)
	__GongStructShape__000000_Default_BookDetailsGroup.Position = __Position__000000_Pos_Default_BookDetailsGroup
	__GongStructShape__000001_Default_BookType.Position = __Position__000001_Pos_Default_BookType
	__GongStructShape__000002_Default_Books.Position = __Position__000002_Pos_Default_Books
	__GongStructShape__000002_Default_Books.Fields = append(__GongStructShape__000002_Default_Books.Fields, __Field__000000_Name)
}
