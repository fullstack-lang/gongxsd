package main

import (
	"time"

	"github.com/fullstack-lang/gongxsd/go/models"

	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ time.Time

// Injection point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.StageStruct) {

	// Declaration of instances to stage

	__ComplexType__000000_ := (&models.ComplexType{Name: ``}).Stage(stage)
	__ComplexType__000001_ := (&models.ComplexType{Name: ``}).Stage(stage)

	__Element__000000_ := (&models.Element{Name: ``}).Stage(stage)
	__Element__000001_ := (&models.Element{Name: ``}).Stage(stage)
	__Element__000002_ := (&models.Element{Name: ``}).Stage(stage)
	__Element__000003_ := (&models.Element{Name: ``}).Stage(stage)
	__Element__000004_ := (&models.Element{Name: ``}).Stage(stage)

	__Restriction__000000_ := (&models.Restriction{Name: ``}).Stage(stage)

	__Schema__000000_ := (&models.Schema{Name: ``}).Stage(stage)

	__Sequence__000000_ := (&models.Sequence{Name: ``}).Stage(stage)
	__Sequence__000001_ := (&models.Sequence{Name: ``}).Stage(stage)

	__SimpleType__000000_ := (&models.SimpleType{Name: ``}).Stage(stage)

	// Setup of values

	__ComplexType__000000_.Name = ``
	__ComplexType__000000_.NameXML = ``

	__ComplexType__000001_.Name = ``
	__ComplexType__000001_.NameXML = `bookType`

	__Element__000000_.Name = ``
	__Element__000000_.NameXSD = `books`
	__Element__000000_.Type = ``

	__Element__000001_.Name = ``
	__Element__000001_.NameXSD = `book`
	__Element__000001_.Type = `bookType`

	__Element__000002_.Name = ``
	__Element__000002_.NameXSD = `title`
	__Element__000002_.Type = `xs:string`

	__Element__000003_.Name = ``
	__Element__000003_.NameXSD = `author`
	__Element__000003_.Type = `xs:string`

	__Element__000004_.Name = ``
	__Element__000004_.NameXSD = `year`
	__Element__000004_.Type = `yearType`

	__Restriction__000000_.Name = ``
	__Restriction__000000_.Base = `xs:integer`

	__Schema__000000_.Name = ``

	__Sequence__000000_.Name = ``

	__Sequence__000001_.Name = ``

	__SimpleType__000000_.Name = ``
	__SimpleType__000000_.NameXSD = `yearType`

	// Setup of pointers
	__ComplexType__000000_.Sequence = __Sequence__000000_
	__ComplexType__000001_.Sequence = __Sequence__000001_
	__Element__000000_.ComplexType = __ComplexType__000000_
	__Schema__000000_.Elements = append(__Schema__000000_.Elements, __Element__000000_)
	__Schema__000000_.SimpleTypes = append(__Schema__000000_.SimpleTypes, __SimpleType__000000_)
	__Schema__000000_.ComplexTypes = append(__Schema__000000_.ComplexTypes, __ComplexType__000001_)
	__Sequence__000000_.Elements = append(__Sequence__000000_.Elements, __Element__000001_)
	__Sequence__000001_.Elements = append(__Sequence__000001_.Elements, __Element__000002_)
	__Sequence__000001_.Elements = append(__Sequence__000001_.Elements, __Element__000003_)
	__Sequence__000001_.Elements = append(__Sequence__000001_.Elements, __Element__000004_)
	__SimpleType__000000_.Restriction = __Restriction__000000_
}
