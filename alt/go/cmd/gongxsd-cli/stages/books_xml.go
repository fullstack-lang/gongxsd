package main

import (
	"time"

	"github.com/fullstack-lang/gongxsd/alt/go/models"

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

	__Annotation__000000_ := (&models.Annotation{Name: ``}).Stage(stage)

	__Documentation__000000_Documentation_of_http_www_w3_org_2001_XMLSchema := (&models.Documentation{Name: `Documentation  of http://www.w3.org/2001/XMLSchema`}).Stage(stage)
	__Documentation__000001_Documentation_of_http_www_w3_org_2001_XMLSchema := (&models.Documentation{Name: `Documentation  of http://www.w3.org/2001/XMLSchema`}).Stage(stage)

	__Schema__000000_http_www_w3_org_2001_XMLSchema := (&models.Schema{Name: `http://www.w3.org/2001/XMLSchema`}).Stage(stage)

	// Setup of values

	__Annotation__000000_.Name = ``

	__Documentation__000000_Documentation_of_http_www_w3_org_2001_XMLSchema.Name = `Documentation  of http://www.w3.org/2001/XMLSchema`
	__Documentation__000000_Documentation_of_http_www_w3_org_2001_XMLSchema.Text = ` This schema defines
            the structure of a simple book collection. It includes types for book details, such as
            title, author, year, and format. `
	__Documentation__000000_Documentation_of_http_www_w3_org_2001_XMLSchema.Source = `http://example.com/schema-docs`
	__Documentation__000000_Documentation_of_http_www_w3_org_2001_XMLSchema.Lang = `en`

	__Documentation__000001_Documentation_of_http_www_w3_org_2001_XMLSchema.Name = `Documentation  of http://www.w3.org/2001/XMLSchema`
	__Documentation__000001_Documentation_of_http_www_w3_org_2001_XMLSchema.Text = ` Ce schéma définit
            la structure d'une collection de livres simple. Il inclut des types pour les détails du
            livre, tels que le titre, l'auteur, l'année et le format. `
	__Documentation__000001_Documentation_of_http_www_w3_org_2001_XMLSchema.Source = `http://example.com/schema-docs`
	__Documentation__000001_Documentation_of_http_www_w3_org_2001_XMLSchema.Lang = `fr`

	__Schema__000000_http_www_w3_org_2001_XMLSchema.Name = `http://www.w3.org/2001/XMLSchema`
	__Schema__000000_http_www_w3_org_2001_XMLSchema.Xs = `http://www.w3.org/2001/XMLSchema`
	__Schema__000000_http_www_w3_org_2001_XMLSchema.Schema_A_ComplexType_A_ComplexContentDummy = 0
	__Schema__000000_http_www_w3_org_2001_XMLSchema.Schema_A_ComplexType_A_ComplexContent_A_Extension_SequenceDummy = 0
	__Schema__000000_http_www_w3_org_2001_XMLSchema.Schema_A_ComplexType_A_ComplexContent_A_Extension_Sequence_Sequence1Dummy = 0

	// Setup of pointers
	__Annotation__000000_.Documentations = append(__Annotation__000000_.Documentations, __Documentation__000000_Documentation_of_http_www_w3_org_2001_XMLSchema)
	__Annotation__000000_.Documentations = append(__Annotation__000000_.Documentations, __Documentation__000001_Documentation_of_http_www_w3_org_2001_XMLSchema)
	__Schema__000000_http_www_w3_org_2001_XMLSchema.Annotation = __Annotation__000000_
}
