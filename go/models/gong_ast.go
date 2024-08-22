// generated code - do not edit
package models

import (
	"bufio"
	"errors"
	"go/ast"
	"go/doc/comment"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var dummy_strconv_import strconv.NumError
var dummy_time_import time.Time

// swagger:ignore
type GONG__ExpressionType string

const (
	GONG__STRUCT_INSTANCE      GONG__ExpressionType = "STRUCT_INSTANCE"
	GONG__FIELD_OR_CONST_VALUE GONG__ExpressionType = "FIELD_OR_CONST_VALUE"
	GONG__FIELD_VALUE          GONG__ExpressionType = "FIELD_VALUE"
	GONG__ENUM_CAST_INT        GONG__ExpressionType = "ENUM_CAST_INT"
	GONG__ENUM_CAST_STRING     GONG__ExpressionType = "ENUM_CAST_STRING"
	GONG__IDENTIFIER_CONST     GONG__ExpressionType = "IDENTIFIER_CONST"
)

// ParseAstFile Parse pathToFile and stages all instances
// declared in the file
func ParseAstFile(stage *StageStruct, pathToFile string) error {

	ReplaceOldDeclarationsInFile(pathToFile)

	fileOfInterest, err := filepath.Abs(pathToFile)
	if err != nil {
		return errors.New("Path does not exist %s ;" + fileOfInterest)
	}

	fset := token.NewFileSet()
	// startParser := time.Now()
	inFile, errParser := parser.ParseFile(fset, fileOfInterest, nil, parser.ParseComments)
	// log.Printf("Parser took %s", time.Since(startParser))

	if errParser != nil {
		return errors.New("Unable to parser " + errParser.Error())
	}

	return ParseAstFileFromAst(stage, inFile, fset)
}

// ParseAstFile Parse pathToFile and stages all instances
// declared in the file
func ParseAstFileFromAst(stage *StageStruct, inFile *ast.File, fset *token.FileSet) error {
	// if there is a meta package import, it is the third import
	if len(inFile.Imports) > 3 {
		log.Fatalln("Too many imports in file", inFile.Name)
	}
	if len(inFile.Imports) == 3 {
		stage.MetaPackageImportAlias = inFile.Imports[2].Name.Name
		stage.MetaPackageImportPath = inFile.Imports[2].Path.Value
	}

	// astCoordinate := "File "
	// log.Println(// astCoordinate)
	for _, decl := range inFile.Decls {
		switch decl := decl.(type) {
		case *ast.FuncDecl:
			funcDecl := decl
			// astCoordinate := // astCoordinate + "\tFunction " + funcDecl.Name.Name
			if name := funcDecl.Name; name != nil {
				isOfInterest := strings.Contains(funcDecl.Name.Name, "_")
				if !isOfInterest {
					continue
				}
				// log.Println(// astCoordinate)
			}
			if doc := funcDecl.Doc; doc != nil {
				// astCoordinate := // astCoordinate + "\tDoc"
				for _, comment := range doc.List {
					_ = comment
					// astCoordinate := // astCoordinate + "\tComment: " + comment.Text
					// log.Println(// astCoordinate)
				}
			}
			if body := funcDecl.Body; body != nil {
				// astCoordinate := // astCoordinate + "\tBody: "
				for _, stmt := range body.List {
					switch stmt := stmt.(type) {
					case *ast.ExprStmt:
						exprStmt := stmt
						// astCoordinate := // astCoordinate + "\tExprStmt: "
						switch expr := exprStmt.X.(type) {
						case *ast.CallExpr:
							// astCoordinate := // astCoordinate + "\tCallExpr: "
							callExpr := expr
							switch fun := callExpr.Fun.(type) {
							case *ast.Ident:
								ident := fun
								_ = ident
								// astCoordinate := // astCoordinate + "\tIdent: " + ident.Name
								// log.Println(// astCoordinate)
							}
						}
					case *ast.AssignStmt:
						// Create an ast.CommentMap from the ast.File's comments.
						// This helps keeping the association between comments
						// and AST nodes.
						cmap := ast.NewCommentMap(fset, inFile, inFile.Comments)
						astCoordinate := "\tAssignStmt: "
						// log.Println(// astCoordinate)
						assignStmt := stmt
						instance, id, gongstruct, fieldName :=
							UnmarshallGongstructStaging(
								stage, &cmap, assignStmt, astCoordinate)
						_ = instance
						_ = id
						_ = gongstruct
						_ = fieldName
					}
				}
			}
		case *ast.GenDecl:
			genDecl := decl
			// log.Println("\tAST GenDecl: ")
			if doc := genDecl.Doc; doc != nil {
				for _, comment := range doc.List {
					_ = comment
					// log.Println("\tAST Comment: ", comment.Text)
				}
			}
			for _, spec := range genDecl.Specs {
				switch spec := spec.(type) {
				case *ast.ImportSpec:
					importSpec := spec
					if path := importSpec.Path; path != nil {
						// log.Println("\t\tAST Path: ", path.Value)
					}
				case *ast.ValueSpec:
					ident := spec.Names[0]
					_ = ident
					if !strings.HasPrefix(ident.Name, "_") {
						continue
					}
					// declaration of a variable without initial value
					if len(spec.Values) == 0 {
						continue
					}
					switch compLit := spec.Values[0].(type) {
					case *ast.CompositeLit:
						var key string
						_ = key
						var value string
						_ = value
						for _, elt := range compLit.Elts {

							// each elt is an expression for struct or for field such as
							// for struct
							//
							//         "dummy.Dummy": &(dummy.Dummy{})
							//
							// or, for field
							//
							//          "dummy.Dummy.Name": (dummy.Dummy{}).Name,
							//
							// first node in the AST is a key value expression
							var ok bool
							var kve *ast.KeyValueExpr
							if kve, ok = elt.(*ast.KeyValueExpr); !ok {
								log.Fatal("Expression should be key value expression" +
									fset.Position(kve.Pos()).String())
							}

							switch bl := kve.Key.(type) {
							case *ast.BasicLit:
								key = bl.Value // "\"dumm.Dummy\"" is the value

								// one remove the ambracing double quotes
								key = strings.TrimPrefix(key, "\"")
								key = strings.TrimSuffix(key, "\"")
							}
							var expressionType GONG__ExpressionType = GONG__STRUCT_INSTANCE
							var docLink GONG__Identifier

							var fieldName string
							var ue *ast.UnaryExpr
							if ue, ok = kve.Value.(*ast.UnaryExpr); !ok {
								expressionType = GONG__FIELD_OR_CONST_VALUE
							}

							var callExpr *ast.CallExpr
							if callExpr, ok = kve.Value.(*ast.CallExpr); ok {

								var se *ast.SelectorExpr
								if se, ok = callExpr.Fun.(*ast.SelectorExpr); !ok {
									log.Fatal("Expression should be a selector expression" +
										fset.Position(callExpr.Pos()).String())
								}

								var id *ast.Ident
								if id, ok = se.X.(*ast.Ident); !ok {
									log.Fatal("Expression should be an ident" +
										fset.Position(se.Pos()).String())
								}

								// check the arg type to select wether this is a int or a string enum
								var bl *ast.BasicLit
								if bl, ok = callExpr.Args[0].(*ast.BasicLit); ok {
									switch bl.Kind {
									case token.STRING:
										expressionType = GONG__ENUM_CAST_STRING
									case token.INT:
										expressionType = GONG__ENUM_CAST_INT
									}
								} else {
									log.Fatal("Expression should be a basic lit" +
										fset.Position(se.Pos()).String())
								}

								docLink.Ident = id.Name + "." + se.Sel.Name
								_ = callExpr
							}

							var se2 *ast.SelectorExpr
							switch expressionType {
							case GONG__FIELD_OR_CONST_VALUE:
								if se2, ok = kve.Value.(*ast.SelectorExpr); ok {

									var ident *ast.Ident
									if _, ok = se2.X.(*ast.ParenExpr); ok {
										expressionType = GONG__FIELD_VALUE
										fieldName = se2.Sel.Name
									} else if ident, ok = se2.X.(*ast.Ident); ok {
										expressionType = GONG__IDENTIFIER_CONST
										docLink.Ident = ident.Name + "." + se2.Sel.Name
									} else {
										log.Fatal("Expression should be a selector expression or an ident" +
											fset.Position(kve.Pos()).String())
									}
								} else {

								}
							}

							var pe *ast.ParenExpr
							switch expressionType {
							case GONG__STRUCT_INSTANCE:
								if pe, ok = ue.X.(*ast.ParenExpr); !ok {
									log.Fatal("Expression should be parenthese expression" +
										fset.Position(ue.Pos()).String())
								}
							case GONG__FIELD_VALUE:
								if pe, ok = se2.X.(*ast.ParenExpr); !ok {
									log.Fatal("Expression should be parenthese expression" +
										fset.Position(ue.Pos()).String())
								}
							}
							switch expressionType {
							case GONG__FIELD_VALUE, GONG__STRUCT_INSTANCE:
								// expect a Composite Litteral with no Element <type>{}
								var cl *ast.CompositeLit
								if cl, ok = pe.X.(*ast.CompositeLit); !ok {
									log.Fatal("Expression should be a composite lit" +
										fset.Position(pe.Pos()).String())
								}

								var se *ast.SelectorExpr
								if se, ok = cl.Type.(*ast.SelectorExpr); !ok {
									log.Fatal("Expression should be a selector" +
										fset.Position(cl.Pos()).String())
								}

								var id *ast.Ident
								if id, ok = se.X.(*ast.Ident); !ok {
									log.Fatal("Expression should be an ident" +
										fset.Position(se.Pos()).String())
								}
								docLink.Ident = id.Name + "." + se.Sel.Name
							}

							switch expressionType {
							case GONG__FIELD_VALUE:
								docLink.Ident += "." + fieldName
							}

							// if map_DocLink_Identifier has the same ident, this means
							// that no renaming has occured since the last processing of the
							// file. But it is neccessary to keep it in memory for the
							// marshalling
							if docLink.Ident == key {
								// continue
							}

							// otherwise, one stores the new ident (after renaming) in the
							// renaming map
							docLink.Type = expressionType
							stage.Map_DocLink_Renaming[key] = docLink
						}
					}
				}
			}
		}

	}
	return nil
}

var __gong__map_Indentifiers_gongstructName = make(map[string]string)

// insertion point for identifiers maps
var __gong__map_All = make(map[string]*All)
var __gong__map_Annotation = make(map[string]*Annotation)
var __gong__map_Attribute = make(map[string]*Attribute)
var __gong__map_AttributeGroup = make(map[string]*AttributeGroup)
var __gong__map_Choice = make(map[string]*Choice)
var __gong__map_ComplexContent = make(map[string]*ComplexContent)
var __gong__map_ComplexType = make(map[string]*ComplexType)
var __gong__map_Documentation = make(map[string]*Documentation)
var __gong__map_Element = make(map[string]*Element)
var __gong__map_Enumeration = make(map[string]*Enumeration)
var __gong__map_Extension = make(map[string]*Extension)
var __gong__map_Group = make(map[string]*Group)
var __gong__map_Length = make(map[string]*Length)
var __gong__map_MaxInclusive = make(map[string]*MaxInclusive)
var __gong__map_MaxLength = make(map[string]*MaxLength)
var __gong__map_MinInclusive = make(map[string]*MinInclusive)
var __gong__map_MinLength = make(map[string]*MinLength)
var __gong__map_Pattern = make(map[string]*Pattern)
var __gong__map_Restriction = make(map[string]*Restriction)
var __gong__map_Schema = make(map[string]*Schema)
var __gong__map_Sequence = make(map[string]*Sequence)
var __gong__map_SimpleContent = make(map[string]*SimpleContent)
var __gong__map_SimpleType = make(map[string]*SimpleType)
var __gong__map_TotalDigit = make(map[string]*TotalDigit)
var __gong__map_Union = make(map[string]*Union)
var __gong__map_WhiteSpace = make(map[string]*WhiteSpace)

// Parser needs to be configured for having the [Name1.Name2] or [pkg.Name1] ...
// to be recognized as a proper identifier.
// While this was introduced in go 1.19, it is not yet implemented in
// gopls (see [issue](https://github.com/golang/go/issues/57559)
func lookupPackage(name string) (importPath string, ok bool) {
	return name, true
}
func lookupSym(recv, name string) (ok bool) {
	if recv == "" {
		return true
	}
	return false
}

// UnmarshallGoStaging unmarshall a go assign statement
func UnmarshallGongstructStaging(stage *StageStruct, cmap *ast.CommentMap, assignStmt *ast.AssignStmt, astCoordinate_ string) (
	instance any,
	identifier string,
	gongstructName string,
	fieldName string) {

	// used for debug purposes
	astCoordinate := "\tAssignStmt: "

	//
	// First parse all comment groups in the assignement
	// if a comment "//gong:ident [DocLink]" is met and is followed by a string assignement.
	// modify the following AST assignement to assigns the DocLink text to the string value
	//

	// get the comment group of the assigStmt
	commentGroups := (*cmap)[assignStmt]
	// get the the prefix
	var hasGongIdentDirective bool
	var commentText string
	var docLinkText string
	for _, commentGroup := range commentGroups {
		for _, comment := range commentGroup.List {
			if strings.HasPrefix(comment.Text, "//gong:ident") {
				hasGongIdentDirective = true
				commentText = comment.Text
			}
		}
	}
	if hasGongIdentDirective {
		// parser configured to find doclinks
		var docLinkFinder comment.Parser
		docLinkFinder.LookupPackage = lookupPackage
		docLinkFinder.LookupSym = lookupSym
		doc := docLinkFinder.Parse(commentText)

		for _, block := range doc.Content {
			switch paragraph := block.(type) {
			case *comment.Paragraph:
				_ = paragraph
				for _, text := range paragraph.Text {
					switch docLink := text.(type) {
					case *comment.DocLink:
						if docLink.Recv == "" {
							docLinkText = docLink.ImportPath + "." + docLink.Name
						} else {
							docLinkText = docLink.ImportPath + "." + docLink.Recv + "." + docLink.Name
						}

						// we check wether the doc link has been renamed
						// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
						if renamed, ok := (stage.Map_DocLink_Renaming)[docLinkText]; ok {
							docLinkText = renamed.Ident
						}
					}
				}
			}
		}
	}

	for rank, expr := range assignStmt.Lhs {
		if rank > 0 {
			continue
		}
		switch expr := expr.(type) {
		case *ast.Ident:
			// we are on a variable declaration
			ident := expr
			// astCoordinate := astCoordinate + "\tLhs" + "." + ident.Name
			// log.Println(astCoordinate)
			identifier = ident.Name
		case *ast.SelectorExpr:
			// we are on a variable assignement
			selectorExpr := expr
			// astCoordinate := astCoordinate + "\tLhs" + "." + selectorExpr.X.(*ast.Ident).Name + "." + selectorExpr.Sel.Name
			// log.Println(astCoordinate)
			identifier = selectorExpr.X.(*ast.Ident).Name
			fieldName = selectorExpr.Sel.Name
		}
	}
	for _, expr := range assignStmt.Rhs {
		// astCoordinate := astCoordinate + "\tRhs"
		switch expr := expr.(type) {
		case *ast.CallExpr:
			callExpr := expr
			// astCoordinate := astCoordinate + "\tFun"
			switch fun := callExpr.Fun.(type) {
			// the is Fun      Expr
			// function expression xxx.Stage()
			case *ast.SelectorExpr:
				selectorExpr := fun
				// astCoordinate := astCoordinate + "\tSelectorExpr"
				switch x := selectorExpr.X.(type) {
				case *ast.ParenExpr:
					// A ParenExpr node represents a parenthesized expression.
					// the is the
					//   { Name : "A1"}
					// astCoordinate := astCoordinate + "\tX"
					parenExpr := x
					switch x := parenExpr.X.(type) {
					case *ast.UnaryExpr:
						unaryExpr := x
						// astCoordinate := astCoordinate + "\tUnaryExpr"
						switch x := unaryExpr.X.(type) {
						case *ast.CompositeLit:
							instanceName := "NoName yet"
							compositeLit := x
							// astCoordinate := astCoordinate + "\tX(CompositeLit)"
							for _, elt := range compositeLit.Elts {
								// astCoordinate := astCoordinate + "\tElt"
								switch elt := elt.(type) {
								case *ast.KeyValueExpr:
									// This is expression
									//     Name: "A1"
									keyValueExpr := elt
									// astCoordinate := astCoordinate + "\tKeyValueExpr"
									switch key := keyValueExpr.Key.(type) {
									case *ast.Ident:
										ident := key
										_ = ident
										// astCoordinate := astCoordinate + "\tKey" + "." + ident.Name
										// log.Println(astCoordinate)
									}
									switch value := keyValueExpr.Value.(type) {
									case *ast.BasicLit:
										basicLit := value
										// astCoordinate := astCoordinate + "\tBasicLit Value" + "." + basicLit.Value
										// log.Println(astCoordinate)
										instanceName = basicLit.Value

										// remove first and last char
										instanceName = instanceName[1 : len(instanceName)-1]
									}
								}
							}
							astCoordinate2 := astCoordinate + "\tType"
							_ = astCoordinate2
							switch type_ := compositeLit.Type.(type) {
							case *ast.SelectorExpr:
								slcExpr := type_
								// astCoordinate := astCoordinate2 + "\tSelectorExpr"
								switch X := slcExpr.X.(type) {
								case *ast.Ident:
									ident := X
									_ = ident
									// astCoordinate := astCoordinate + "\tX" + "." + ident.Name
									// log.Println(astCoordinate)
								}
								if Sel := slcExpr.Sel; Sel != nil {
									// astCoordinate := astCoordinate + "\tSel" + "." + Sel.Name
									// log.Println(astCoordinate)

									gongstructName = Sel.Name
									// this is the place where an instance is created
									switch gongstructName {
									// insertion point for identifiers
									case "All":
										instanceAll := (&All{Name: instanceName}).Stage(stage)
										instance = any(instanceAll)
										__gong__map_All[identifier] = instanceAll
									case "Annotation":
										instanceAnnotation := (&Annotation{Name: instanceName}).Stage(stage)
										instance = any(instanceAnnotation)
										__gong__map_Annotation[identifier] = instanceAnnotation
									case "Attribute":
										instanceAttribute := (&Attribute{Name: instanceName}).Stage(stage)
										instance = any(instanceAttribute)
										__gong__map_Attribute[identifier] = instanceAttribute
									case "AttributeGroup":
										instanceAttributeGroup := (&AttributeGroup{Name: instanceName}).Stage(stage)
										instance = any(instanceAttributeGroup)
										__gong__map_AttributeGroup[identifier] = instanceAttributeGroup
									case "Choice":
										instanceChoice := (&Choice{Name: instanceName}).Stage(stage)
										instance = any(instanceChoice)
										__gong__map_Choice[identifier] = instanceChoice
									case "ComplexContent":
										instanceComplexContent := (&ComplexContent{Name: instanceName}).Stage(stage)
										instance = any(instanceComplexContent)
										__gong__map_ComplexContent[identifier] = instanceComplexContent
									case "ComplexType":
										instanceComplexType := (&ComplexType{Name: instanceName}).Stage(stage)
										instance = any(instanceComplexType)
										__gong__map_ComplexType[identifier] = instanceComplexType
									case "Documentation":
										instanceDocumentation := (&Documentation{Name: instanceName}).Stage(stage)
										instance = any(instanceDocumentation)
										__gong__map_Documentation[identifier] = instanceDocumentation
									case "Element":
										instanceElement := (&Element{Name: instanceName}).Stage(stage)
										instance = any(instanceElement)
										__gong__map_Element[identifier] = instanceElement
									case "Enumeration":
										instanceEnumeration := (&Enumeration{Name: instanceName}).Stage(stage)
										instance = any(instanceEnumeration)
										__gong__map_Enumeration[identifier] = instanceEnumeration
									case "Extension":
										instanceExtension := (&Extension{Name: instanceName}).Stage(stage)
										instance = any(instanceExtension)
										__gong__map_Extension[identifier] = instanceExtension
									case "Group":
										instanceGroup := (&Group{Name: instanceName}).Stage(stage)
										instance = any(instanceGroup)
										__gong__map_Group[identifier] = instanceGroup
									case "Length":
										instanceLength := (&Length{Name: instanceName}).Stage(stage)
										instance = any(instanceLength)
										__gong__map_Length[identifier] = instanceLength
									case "MaxInclusive":
										instanceMaxInclusive := (&MaxInclusive{Name: instanceName}).Stage(stage)
										instance = any(instanceMaxInclusive)
										__gong__map_MaxInclusive[identifier] = instanceMaxInclusive
									case "MaxLength":
										instanceMaxLength := (&MaxLength{Name: instanceName}).Stage(stage)
										instance = any(instanceMaxLength)
										__gong__map_MaxLength[identifier] = instanceMaxLength
									case "MinInclusive":
										instanceMinInclusive := (&MinInclusive{Name: instanceName}).Stage(stage)
										instance = any(instanceMinInclusive)
										__gong__map_MinInclusive[identifier] = instanceMinInclusive
									case "MinLength":
										instanceMinLength := (&MinLength{Name: instanceName}).Stage(stage)
										instance = any(instanceMinLength)
										__gong__map_MinLength[identifier] = instanceMinLength
									case "Pattern":
										instancePattern := (&Pattern{Name: instanceName}).Stage(stage)
										instance = any(instancePattern)
										__gong__map_Pattern[identifier] = instancePattern
									case "Restriction":
										instanceRestriction := (&Restriction{Name: instanceName}).Stage(stage)
										instance = any(instanceRestriction)
										__gong__map_Restriction[identifier] = instanceRestriction
									case "Schema":
										instanceSchema := (&Schema{Name: instanceName}).Stage(stage)
										instance = any(instanceSchema)
										__gong__map_Schema[identifier] = instanceSchema
									case "Sequence":
										instanceSequence := (&Sequence{Name: instanceName}).Stage(stage)
										instance = any(instanceSequence)
										__gong__map_Sequence[identifier] = instanceSequence
									case "SimpleContent":
										instanceSimpleContent := (&SimpleContent{Name: instanceName}).Stage(stage)
										instance = any(instanceSimpleContent)
										__gong__map_SimpleContent[identifier] = instanceSimpleContent
									case "SimpleType":
										instanceSimpleType := (&SimpleType{Name: instanceName}).Stage(stage)
										instance = any(instanceSimpleType)
										__gong__map_SimpleType[identifier] = instanceSimpleType
									case "TotalDigit":
										instanceTotalDigit := (&TotalDigit{Name: instanceName}).Stage(stage)
										instance = any(instanceTotalDigit)
										__gong__map_TotalDigit[identifier] = instanceTotalDigit
									case "Union":
										instanceUnion := (&Union{Name: instanceName}).Stage(stage)
										instance = any(instanceUnion)
										__gong__map_Union[identifier] = instanceUnion
									case "WhiteSpace":
										instanceWhiteSpace := (&WhiteSpace{Name: instanceName}).Stage(stage)
										instance = any(instanceWhiteSpace)
										__gong__map_WhiteSpace[identifier] = instanceWhiteSpace
									}
									__gong__map_Indentifiers_gongstructName[identifier] = gongstructName
									return
								}
							}
						}
					}
				}
				if sel := selectorExpr.Sel; sel != nil {
					// astCoordinate := astCoordinate + "\tSel" + "." + sel.Name
					// log.Println(astCoordinate)
				}
				for iteration, arg := range callExpr.Args {
					// astCoordinate := astCoordinate + "\tArg"
					switch arg := arg.(type) {
					case *ast.BasicLit:
						basicLit := arg
						// astCoordinate := astCoordinate + "\tBasicLit" + "." + basicLit.Value
						// log.Println(astCoordinate)

						// first iteration should be ignored
						if iteration == 0 {
							continue
						}

						// remove first and last char
						date := basicLit.Value[1 : len(basicLit.Value)-1]
						_ = date

						var ok bool
						gongstructName, ok = __gong__map_Indentifiers_gongstructName[identifier]
						if !ok {
							log.Fatalln("gongstructName not found for identifier", identifier)
						}
						switch gongstructName {
						// insertion point for basic lit assignments
						case "All":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Annotation":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Attribute":
							switch fieldName {
							// insertion point for date assign code
							}
						case "AttributeGroup":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Choice":
							switch fieldName {
							// insertion point for date assign code
							}
						case "ComplexContent":
							switch fieldName {
							// insertion point for date assign code
							}
						case "ComplexType":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Documentation":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Element":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Enumeration":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Extension":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Group":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Length":
							switch fieldName {
							// insertion point for date assign code
							}
						case "MaxInclusive":
							switch fieldName {
							// insertion point for date assign code
							}
						case "MaxLength":
							switch fieldName {
							// insertion point for date assign code
							}
						case "MinInclusive":
							switch fieldName {
							// insertion point for date assign code
							}
						case "MinLength":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Pattern":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Restriction":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Schema":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Sequence":
							switch fieldName {
							// insertion point for date assign code
							}
						case "SimpleContent":
							switch fieldName {
							// insertion point for date assign code
							}
						case "SimpleType":
							switch fieldName {
							// insertion point for date assign code
							}
						case "TotalDigit":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Union":
							switch fieldName {
							// insertion point for date assign code
							}
						case "WhiteSpace":
							switch fieldName {
							// insertion point for date assign code
							}
						}
					}
				}
			case *ast.Ident:
				// append function
				ident := fun
				_ = ident
				// astCoordinate := astCoordinate + "\tIdent" + "." + ident.Name
				// log.Println(astCoordinate)
			}
			for _, arg := range callExpr.Args {
				// astCoordinate := astCoordinate + "\tArg"
				switch arg := arg.(type) {
				case *ast.Ident:
					ident := arg
					_ = ident
					// astCoordinate := astCoordinate + "\tIdent" + "." + ident.Name
					// log.Println(astCoordinate)
					var ok bool
					gongstructName, ok = __gong__map_Indentifiers_gongstructName[identifier]
					if !ok {
						log.Fatalln("gongstructName not found for identifier", identifier)
					}
					switch gongstructName {
					// insertion point for slice of pointers assignments
					case "All":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Sequences":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Sequence[targetIdentifier]
							__gong__map_All[identifier].Sequences =
								append(__gong__map_All[identifier].Sequences, target)
						case "Alls":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_All[targetIdentifier]
							__gong__map_All[identifier].Alls =
								append(__gong__map_All[identifier].Alls, target)
						case "Choices":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Choice[targetIdentifier]
							__gong__map_All[identifier].Choices =
								append(__gong__map_All[identifier].Choices, target)
						case "Groups":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Group[targetIdentifier]
							__gong__map_All[identifier].Groups =
								append(__gong__map_All[identifier].Groups, target)
						case "Elements":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Element[targetIdentifier]
							__gong__map_All[identifier].Elements =
								append(__gong__map_All[identifier].Elements, target)
						}
					case "Annotation":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Documentations":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Documentation[targetIdentifier]
							__gong__map_Annotation[identifier].Documentations =
								append(__gong__map_Annotation[identifier].Documentations, target)
						}
					case "Attribute":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "AttributeGroup":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "AttributeGroups":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_AttributeGroup[targetIdentifier]
							__gong__map_AttributeGroup[identifier].AttributeGroups =
								append(__gong__map_AttributeGroup[identifier].AttributeGroups, target)
						case "Attributes":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Attribute[targetIdentifier]
							__gong__map_AttributeGroup[identifier].Attributes =
								append(__gong__map_AttributeGroup[identifier].Attributes, target)
						}
					case "Choice":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Sequences":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Sequence[targetIdentifier]
							__gong__map_Choice[identifier].Sequences =
								append(__gong__map_Choice[identifier].Sequences, target)
						case "Alls":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_All[targetIdentifier]
							__gong__map_Choice[identifier].Alls =
								append(__gong__map_Choice[identifier].Alls, target)
						case "Choices":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Choice[targetIdentifier]
							__gong__map_Choice[identifier].Choices =
								append(__gong__map_Choice[identifier].Choices, target)
						case "Groups":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Group[targetIdentifier]
							__gong__map_Choice[identifier].Groups =
								append(__gong__map_Choice[identifier].Groups, target)
						case "Elements":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Element[targetIdentifier]
							__gong__map_Choice[identifier].Elements =
								append(__gong__map_Choice[identifier].Elements, target)
						}
					case "ComplexContent":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "ComplexType":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Sequences":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Sequence[targetIdentifier]
							__gong__map_ComplexType[identifier].Sequences =
								append(__gong__map_ComplexType[identifier].Sequences, target)
						case "Alls":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_All[targetIdentifier]
							__gong__map_ComplexType[identifier].Alls =
								append(__gong__map_ComplexType[identifier].Alls, target)
						case "Choices":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Choice[targetIdentifier]
							__gong__map_ComplexType[identifier].Choices =
								append(__gong__map_ComplexType[identifier].Choices, target)
						case "Groups":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Group[targetIdentifier]
							__gong__map_ComplexType[identifier].Groups =
								append(__gong__map_ComplexType[identifier].Groups, target)
						case "Elements":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Element[targetIdentifier]
							__gong__map_ComplexType[identifier].Elements =
								append(__gong__map_ComplexType[identifier].Elements, target)
						case "Attributes":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Attribute[targetIdentifier]
							__gong__map_ComplexType[identifier].Attributes =
								append(__gong__map_ComplexType[identifier].Attributes, target)
						case "AttributeGroups":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_AttributeGroup[targetIdentifier]
							__gong__map_ComplexType[identifier].AttributeGroups =
								append(__gong__map_ComplexType[identifier].AttributeGroups, target)
						}
					case "Documentation":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Element":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Groups":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Group[targetIdentifier]
							__gong__map_Element[identifier].Groups =
								append(__gong__map_Element[identifier].Groups, target)
						}
					case "Enumeration":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Extension":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Sequences":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Sequence[targetIdentifier]
							__gong__map_Extension[identifier].Sequences =
								append(__gong__map_Extension[identifier].Sequences, target)
						case "Alls":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_All[targetIdentifier]
							__gong__map_Extension[identifier].Alls =
								append(__gong__map_Extension[identifier].Alls, target)
						case "Choices":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Choice[targetIdentifier]
							__gong__map_Extension[identifier].Choices =
								append(__gong__map_Extension[identifier].Choices, target)
						case "Groups":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Group[targetIdentifier]
							__gong__map_Extension[identifier].Groups =
								append(__gong__map_Extension[identifier].Groups, target)
						case "Elements":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Element[targetIdentifier]
							__gong__map_Extension[identifier].Elements =
								append(__gong__map_Extension[identifier].Elements, target)
						case "Attributes":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Attribute[targetIdentifier]
							__gong__map_Extension[identifier].Attributes =
								append(__gong__map_Extension[identifier].Attributes, target)
						}
					case "Group":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Sequences":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Sequence[targetIdentifier]
							__gong__map_Group[identifier].Sequences =
								append(__gong__map_Group[identifier].Sequences, target)
						case "Alls":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_All[targetIdentifier]
							__gong__map_Group[identifier].Alls =
								append(__gong__map_Group[identifier].Alls, target)
						case "Choices":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Choice[targetIdentifier]
							__gong__map_Group[identifier].Choices =
								append(__gong__map_Group[identifier].Choices, target)
						case "Groups":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Group[targetIdentifier]
							__gong__map_Group[identifier].Groups =
								append(__gong__map_Group[identifier].Groups, target)
						case "Elements":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Element[targetIdentifier]
							__gong__map_Group[identifier].Elements =
								append(__gong__map_Group[identifier].Elements, target)
						}
					case "Length":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "MaxInclusive":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "MaxLength":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "MinInclusive":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "MinLength":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Pattern":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Restriction":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Enumerations":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Enumeration[targetIdentifier]
							__gong__map_Restriction[identifier].Enumerations =
								append(__gong__map_Restriction[identifier].Enumerations, target)
						}
					case "Schema":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Elements":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Element[targetIdentifier]
							__gong__map_Schema[identifier].Elements =
								append(__gong__map_Schema[identifier].Elements, target)
						case "SimpleTypes":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_SimpleType[targetIdentifier]
							__gong__map_Schema[identifier].SimpleTypes =
								append(__gong__map_Schema[identifier].SimpleTypes, target)
						case "ComplexTypes":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_ComplexType[targetIdentifier]
							__gong__map_Schema[identifier].ComplexTypes =
								append(__gong__map_Schema[identifier].ComplexTypes, target)
						case "AttributeGroups":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_AttributeGroup[targetIdentifier]
							__gong__map_Schema[identifier].AttributeGroups =
								append(__gong__map_Schema[identifier].AttributeGroups, target)
						case "Groups":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Group[targetIdentifier]
							__gong__map_Schema[identifier].Groups =
								append(__gong__map_Schema[identifier].Groups, target)
						}
					case "Sequence":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Sequences":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Sequence[targetIdentifier]
							__gong__map_Sequence[identifier].Sequences =
								append(__gong__map_Sequence[identifier].Sequences, target)
						case "Alls":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_All[targetIdentifier]
							__gong__map_Sequence[identifier].Alls =
								append(__gong__map_Sequence[identifier].Alls, target)
						case "Choices":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Choice[targetIdentifier]
							__gong__map_Sequence[identifier].Choices =
								append(__gong__map_Sequence[identifier].Choices, target)
						case "Groups":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Group[targetIdentifier]
							__gong__map_Sequence[identifier].Groups =
								append(__gong__map_Sequence[identifier].Groups, target)
						case "Elements":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Element[targetIdentifier]
							__gong__map_Sequence[identifier].Elements =
								append(__gong__map_Sequence[identifier].Elements, target)
						}
					case "SimpleContent":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "SimpleType":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "TotalDigit":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Union":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "WhiteSpace":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					}
				case *ast.SelectorExpr:
					slcExpr := arg
					// astCoordinate := astCoordinate + "\tSelectorExpr"
					switch X := slcExpr.X.(type) {
					case *ast.Ident:
						ident := X
						_ = ident
						// astCoordinate := astCoordinate + "\tX" + "." + ident.Name
						// log.Println(astCoordinate)

					}
					if Sel := slcExpr.Sel; Sel != nil {
						// astCoordinate := astCoordinate + "\tSel" + "." + Sel.Name
						// log.Println(astCoordinate)
					}
				}
			}
		case *ast.BasicLit, *ast.UnaryExpr:

			var basicLit *ast.BasicLit
			var exprSign = 1.0
			_ = exprSign // in case this is not used

			if bl, ok := expr.(*ast.BasicLit); ok {
				// expression is  for instance ... = 18.000
				basicLit = bl
			} else if ue, ok := expr.(*ast.UnaryExpr); ok {
				// expression is  for instance ... = -18.000
				// we want to extract a *ast.BasicLit from the *ast.UnaryExpr
				basicLit = ue.X.(*ast.BasicLit)
				exprSign = -1
			}

			// astCoordinate := astCoordinate + "\tBasicLit" + "." + basicLit.Value
			// log.Println(astCoordinate)
			var ok bool
			gongstructName, ok = __gong__map_Indentifiers_gongstructName[identifier]
			if !ok {
				log.Fatalln("gongstructName not found for identifier", identifier)
			}

			// substitute the RHS part of the assignment if a //gong:ident directive is met
			if hasGongIdentDirective {
				basicLit.Value = "[" + docLinkText + "]"
			}

			switch gongstructName {
			// insertion point for basic lit assignments
			case "All":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_All[identifier].Name = fielValue
				case "MinOccurs":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_All[identifier].MinOccurs = fielValue
				case "MaxOccurs":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_All[identifier].MaxOccurs = fielValue
				case "OuterElementName":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_All[identifier].OuterElementName = fielValue
				case "Order":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_All[identifier].Order = int(exprSign) * int(fielValue)
				case "Depth":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_All[identifier].Depth = int(exprSign) * int(fielValue)
				}
			case "Annotation":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Annotation[identifier].Name = fielValue
				}
			case "Attribute":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Attribute[identifier].Name = fielValue
				case "NameXSD":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Attribute[identifier].NameXSD = fielValue
				case "Type":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Attribute[identifier].Type = fielValue
				case "GoIdentifier":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Attribute[identifier].GoIdentifier = fielValue
				case "Default":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Attribute[identifier].Default = fielValue
				case "Use":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Attribute[identifier].Use = fielValue
				case "Form":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Attribute[identifier].Form = fielValue
				case "Fixed":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Attribute[identifier].Fixed = fielValue
				case "Ref":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Attribute[identifier].Ref = fielValue
				case "TargetNamespace":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Attribute[identifier].TargetNamespace = fielValue
				case "SimpleType":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Attribute[identifier].SimpleType = fielValue
				case "IDXSD":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Attribute[identifier].IDXSD = fielValue
				}
			case "AttributeGroup":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_AttributeGroup[identifier].Name = fielValue
				case "NameXSD":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_AttributeGroup[identifier].NameXSD = fielValue
				case "GoIdentifier":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_AttributeGroup[identifier].GoIdentifier = fielValue
				case "Ref":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_AttributeGroup[identifier].Ref = fielValue
				}
			case "Choice":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Choice[identifier].Name = fielValue
				case "MinOccurs":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Choice[identifier].MinOccurs = fielValue
				case "MaxOccurs":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Choice[identifier].MaxOccurs = fielValue
				case "OuterElementName":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Choice[identifier].OuterElementName = fielValue
				case "Order":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Choice[identifier].Order = int(exprSign) * int(fielValue)
				case "Depth":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Choice[identifier].Depth = int(exprSign) * int(fielValue)
				}
			case "ComplexContent":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ComplexContent[identifier].Name = fielValue
				}
			case "ComplexType":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ComplexType[identifier].Name = fielValue
				case "GoIdentifier":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ComplexType[identifier].GoIdentifier = fielValue
				case "NameXSD":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ComplexType[identifier].NameXSD = fielValue
				case "OuterElementName":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ComplexType[identifier].OuterElementName = fielValue
				}
			case "Documentation":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Documentation[identifier].Name = fielValue
				case "Text":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Documentation[identifier].Text = fielValue
				case "Source":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Documentation[identifier].Source = fielValue
				case "Lang":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Documentation[identifier].Lang = fielValue
				}
			case "Element":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Element[identifier].Name = fielValue
				case "Order":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Element[identifier].Order = int(exprSign) * int(fielValue)
				case "Depth":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Element[identifier].Depth = int(exprSign) * int(fielValue)
				case "GoIdentifier":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Element[identifier].GoIdentifier = fielValue
				case "NameXSD":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Element[identifier].NameXSD = fielValue
				case "Type":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Element[identifier].Type = fielValue
				case "MinOccurs":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Element[identifier].MinOccurs = fielValue
				case "MaxOccurs":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Element[identifier].MaxOccurs = fielValue
				case "Default":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Element[identifier].Default = fielValue
				case "Fixed":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Element[identifier].Fixed = fielValue
				case "Nillable":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Element[identifier].Nillable = fielValue
				case "Ref":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Element[identifier].Ref = fielValue
				case "Abstract":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Element[identifier].Abstract = fielValue
				case "Form":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Element[identifier].Form = fielValue
				case "Block":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Element[identifier].Block = fielValue
				case "Final":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Element[identifier].Final = fielValue
				}
			case "Enumeration":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Enumeration[identifier].Name = fielValue
				case "Value":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Enumeration[identifier].Value = fielValue
				}
			case "Extension":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Extension[identifier].Name = fielValue
				case "OuterElementName":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Extension[identifier].OuterElementName = fielValue
				case "Base":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Extension[identifier].Base = fielValue
				case "Ref":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Extension[identifier].Ref = fielValue
				}
			case "Group":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Group[identifier].Name = fielValue
				case "NameXSD":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Group[identifier].NameXSD = fielValue
				case "Ref":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Group[identifier].Ref = fielValue
				case "GoIdentifier":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Group[identifier].GoIdentifier = fielValue
				case "OuterElementName":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Group[identifier].OuterElementName = fielValue
				case "Order":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Group[identifier].Order = int(exprSign) * int(fielValue)
				case "Depth":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Group[identifier].Depth = int(exprSign) * int(fielValue)
				}
			case "Length":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Length[identifier].Name = fielValue
				case "Value":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Length[identifier].Value = fielValue
				}
			case "MaxInclusive":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_MaxInclusive[identifier].Name = fielValue
				case "Value":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_MaxInclusive[identifier].Value = fielValue
				}
			case "MaxLength":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_MaxLength[identifier].Name = fielValue
				case "Value":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_MaxLength[identifier].Value = fielValue
				}
			case "MinInclusive":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_MinInclusive[identifier].Name = fielValue
				case "Value":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_MinInclusive[identifier].Value = fielValue
				}
			case "MinLength":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_MinLength[identifier].Name = fielValue
				case "Value":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_MinLength[identifier].Value = fielValue
				}
			case "Pattern":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Pattern[identifier].Name = fielValue
				case "Value":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Pattern[identifier].Value = fielValue
				}
			case "Restriction":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Restriction[identifier].Name = fielValue
				case "Base":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Restriction[identifier].Base = fielValue
				}
			case "Schema":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Schema[identifier].Name = fielValue
				case "Xs":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Schema[identifier].Xs = fielValue
				}
			case "Sequence":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Sequence[identifier].Name = fielValue
				case "MinOccurs":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Sequence[identifier].MinOccurs = fielValue
				case "MaxOccurs":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Sequence[identifier].MaxOccurs = fielValue
				case "OuterElementName":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Sequence[identifier].OuterElementName = fielValue
				case "Order":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Sequence[identifier].Order = int(exprSign) * int(fielValue)
				case "Depth":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Sequence[identifier].Depth = int(exprSign) * int(fielValue)
				}
			case "SimpleContent":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SimpleContent[identifier].Name = fielValue
				}
			case "SimpleType":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SimpleType[identifier].Name = fielValue
				case "NameXSD":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SimpleType[identifier].NameXSD = fielValue
				}
			case "TotalDigit":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_TotalDigit[identifier].Name = fielValue
				case "Value":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_TotalDigit[identifier].Value = fielValue
				}
			case "Union":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Union[identifier].Name = fielValue
				case "MemberTypes":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Union[identifier].MemberTypes = fielValue
				}
			case "WhiteSpace":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_WhiteSpace[identifier].Name = fielValue
				case "Value":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_WhiteSpace[identifier].Value = fielValue
				}
			}
		case *ast.Ident:
			// assignment to boolean field ?
			ident := expr
			_ = ident
			// astCoordinate := astCoordinate + "\tIdent" + "." + ident.Name
			// log.Println(astCoordinate)
			var ok bool
			gongstructName, ok = __gong__map_Indentifiers_gongstructName[identifier]
			if !ok {
				log.Fatalln("gongstructName not found for identifier", identifier)
			}
			switch gongstructName {
			// insertion point for bool & pointers assignments
			case "All":
				switch fieldName {
				// insertion point for field dependant code
				case "Annotation":
					targetIdentifier := ident.Name
					__gong__map_All[identifier].Annotation = __gong__map_Annotation[targetIdentifier]
				}
			case "Annotation":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Attribute":
				switch fieldName {
				// insertion point for field dependant code
				case "Annotation":
					targetIdentifier := ident.Name
					__gong__map_Attribute[identifier].Annotation = __gong__map_Annotation[targetIdentifier]
				case "HasNameConflict":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Attribute[identifier].HasNameConflict = fielValue
				}
			case "AttributeGroup":
				switch fieldName {
				// insertion point for field dependant code
				case "Annotation":
					targetIdentifier := ident.Name
					__gong__map_AttributeGroup[identifier].Annotation = __gong__map_Annotation[targetIdentifier]
				case "HasNameConflict":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_AttributeGroup[identifier].HasNameConflict = fielValue
				}
			case "Choice":
				switch fieldName {
				// insertion point for field dependant code
				case "Annotation":
					targetIdentifier := ident.Name
					__gong__map_Choice[identifier].Annotation = __gong__map_Annotation[targetIdentifier]
				case "IsDuplicatedInXSD":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Choice[identifier].IsDuplicatedInXSD = fielValue
				}
			case "ComplexContent":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "ComplexType":
				switch fieldName {
				// insertion point for field dependant code
				case "HasNameConflict":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_ComplexType[identifier].HasNameConflict = fielValue
				case "IsAnonymous":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_ComplexType[identifier].IsAnonymous = fielValue
				case "OuterElement":
					targetIdentifier := ident.Name
					__gong__map_ComplexType[identifier].OuterElement = __gong__map_Element[targetIdentifier]
				case "Annotation":
					targetIdentifier := ident.Name
					__gong__map_ComplexType[identifier].Annotation = __gong__map_Annotation[targetIdentifier]
				case "Extension":
					targetIdentifier := ident.Name
					__gong__map_ComplexType[identifier].Extension = __gong__map_Extension[targetIdentifier]
				case "SimpleContent":
					targetIdentifier := ident.Name
					__gong__map_ComplexType[identifier].SimpleContent = __gong__map_SimpleContent[targetIdentifier]
				case "ComplexContent":
					targetIdentifier := ident.Name
					__gong__map_ComplexType[identifier].ComplexContent = __gong__map_ComplexContent[targetIdentifier]
				}
			case "Documentation":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Element":
				switch fieldName {
				// insertion point for field dependant code
				case "HasNameConflict":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Element[identifier].HasNameConflict = fielValue
				case "Annotation":
					targetIdentifier := ident.Name
					__gong__map_Element[identifier].Annotation = __gong__map_Annotation[targetIdentifier]
				case "SimpleType":
					targetIdentifier := ident.Name
					__gong__map_Element[identifier].SimpleType = __gong__map_SimpleType[targetIdentifier]
				case "ComplexType":
					targetIdentifier := ident.Name
					__gong__map_Element[identifier].ComplexType = __gong__map_ComplexType[targetIdentifier]
				case "IsDuplicatedInXSD":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Element[identifier].IsDuplicatedInXSD = fielValue
				}
			case "Enumeration":
				switch fieldName {
				// insertion point for field dependant code
				case "Annotation":
					targetIdentifier := ident.Name
					__gong__map_Enumeration[identifier].Annotation = __gong__map_Annotation[targetIdentifier]
				}
			case "Extension":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Group":
				switch fieldName {
				// insertion point for field dependant code
				case "Annotation":
					targetIdentifier := ident.Name
					__gong__map_Group[identifier].Annotation = __gong__map_Annotation[targetIdentifier]
				case "IsAnonymous":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Group[identifier].IsAnonymous = fielValue
				case "OuterElement":
					targetIdentifier := ident.Name
					__gong__map_Group[identifier].OuterElement = __gong__map_Element[targetIdentifier]
				case "HasNameConflict":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Group[identifier].HasNameConflict = fielValue
				}
			case "Length":
				switch fieldName {
				// insertion point for field dependant code
				case "Annotation":
					targetIdentifier := ident.Name
					__gong__map_Length[identifier].Annotation = __gong__map_Annotation[targetIdentifier]
				}
			case "MaxInclusive":
				switch fieldName {
				// insertion point for field dependant code
				case "Annotation":
					targetIdentifier := ident.Name
					__gong__map_MaxInclusive[identifier].Annotation = __gong__map_Annotation[targetIdentifier]
				}
			case "MaxLength":
				switch fieldName {
				// insertion point for field dependant code
				case "Annotation":
					targetIdentifier := ident.Name
					__gong__map_MaxLength[identifier].Annotation = __gong__map_Annotation[targetIdentifier]
				}
			case "MinInclusive":
				switch fieldName {
				// insertion point for field dependant code
				case "Annotation":
					targetIdentifier := ident.Name
					__gong__map_MinInclusive[identifier].Annotation = __gong__map_Annotation[targetIdentifier]
				}
			case "MinLength":
				switch fieldName {
				// insertion point for field dependant code
				case "Annotation":
					targetIdentifier := ident.Name
					__gong__map_MinLength[identifier].Annotation = __gong__map_Annotation[targetIdentifier]
				}
			case "Pattern":
				switch fieldName {
				// insertion point for field dependant code
				case "Annotation":
					targetIdentifier := ident.Name
					__gong__map_Pattern[identifier].Annotation = __gong__map_Annotation[targetIdentifier]
				}
			case "Restriction":
				switch fieldName {
				// insertion point for field dependant code
				case "Annotation":
					targetIdentifier := ident.Name
					__gong__map_Restriction[identifier].Annotation = __gong__map_Annotation[targetIdentifier]
				case "MinInclusive":
					targetIdentifier := ident.Name
					__gong__map_Restriction[identifier].MinInclusive = __gong__map_MinInclusive[targetIdentifier]
				case "MaxInclusive":
					targetIdentifier := ident.Name
					__gong__map_Restriction[identifier].MaxInclusive = __gong__map_MaxInclusive[targetIdentifier]
				case "Pattern":
					targetIdentifier := ident.Name
					__gong__map_Restriction[identifier].Pattern = __gong__map_Pattern[targetIdentifier]
				case "WhiteSpace":
					targetIdentifier := ident.Name
					__gong__map_Restriction[identifier].WhiteSpace = __gong__map_WhiteSpace[targetIdentifier]
				case "MinLength":
					targetIdentifier := ident.Name
					__gong__map_Restriction[identifier].MinLength = __gong__map_MinLength[targetIdentifier]
				case "MaxLength":
					targetIdentifier := ident.Name
					__gong__map_Restriction[identifier].MaxLength = __gong__map_MaxLength[targetIdentifier]
				case "Length":
					targetIdentifier := ident.Name
					__gong__map_Restriction[identifier].Length = __gong__map_Length[targetIdentifier]
				case "TotalDigit":
					targetIdentifier := ident.Name
					__gong__map_Restriction[identifier].TotalDigit = __gong__map_TotalDigit[targetIdentifier]
				}
			case "Schema":
				switch fieldName {
				// insertion point for field dependant code
				case "Annotation":
					targetIdentifier := ident.Name
					__gong__map_Schema[identifier].Annotation = __gong__map_Annotation[targetIdentifier]
				}
			case "Sequence":
				switch fieldName {
				// insertion point for field dependant code
				case "Annotation":
					targetIdentifier := ident.Name
					__gong__map_Sequence[identifier].Annotation = __gong__map_Annotation[targetIdentifier]
				}
			case "SimpleContent":
				switch fieldName {
				// insertion point for field dependant code
				case "Extension":
					targetIdentifier := ident.Name
					__gong__map_SimpleContent[identifier].Extension = __gong__map_Extension[targetIdentifier]
				case "Restriction":
					targetIdentifier := ident.Name
					__gong__map_SimpleContent[identifier].Restriction = __gong__map_Restriction[targetIdentifier]
				}
			case "SimpleType":
				switch fieldName {
				// insertion point for field dependant code
				case "Annotation":
					targetIdentifier := ident.Name
					__gong__map_SimpleType[identifier].Annotation = __gong__map_Annotation[targetIdentifier]
				case "Restriction":
					targetIdentifier := ident.Name
					__gong__map_SimpleType[identifier].Restriction = __gong__map_Restriction[targetIdentifier]
				case "Union":
					targetIdentifier := ident.Name
					__gong__map_SimpleType[identifier].Union = __gong__map_Union[targetIdentifier]
				}
			case "TotalDigit":
				switch fieldName {
				// insertion point for field dependant code
				case "Annotation":
					targetIdentifier := ident.Name
					__gong__map_TotalDigit[identifier].Annotation = __gong__map_Annotation[targetIdentifier]
				}
			case "Union":
				switch fieldName {
				// insertion point for field dependant code
				case "Annotation":
					targetIdentifier := ident.Name
					__gong__map_Union[identifier].Annotation = __gong__map_Annotation[targetIdentifier]
				}
			case "WhiteSpace":
				switch fieldName {
				// insertion point for field dependant code
				case "Annotation":
					targetIdentifier := ident.Name
					__gong__map_WhiteSpace[identifier].Annotation = __gong__map_Annotation[targetIdentifier]
				}
			}
		case *ast.SelectorExpr:
			// assignment to enum field
			selectorExpr := expr
			// astCoordinate := astCoordinate + "\tSelectorExpr"
			switch X := selectorExpr.X.(type) {
			case *ast.Ident:
				ident := X
				_ = ident
				// astCoordinate := astCoordinate + "\tX" + "." + ident.Name
				// log.Println(astCoordinate)
			}
			if Sel := selectorExpr.Sel; Sel != nil {
				// astCoordinate := astCoordinate + "\tSel" + "." + Sel.Name
				// log.Println(astCoordinate)

				// enum field
				var ok bool
				gongstructName, ok = __gong__map_Indentifiers_gongstructName[identifier]
				if !ok {
					log.Fatalln("gongstructName not found for identifier", identifier)
				}

				// remove first and last char
				enumValue := Sel.Name
				_ = enumValue
				switch gongstructName {
				// insertion point for enums assignments
				case "All":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Annotation":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Attribute":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "AttributeGroup":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Choice":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "ComplexContent":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "ComplexType":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Documentation":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Element":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Enumeration":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Extension":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Group":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Length":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "MaxInclusive":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "MaxLength":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "MinInclusive":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "MinLength":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Pattern":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Restriction":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Schema":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Sequence":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "SimpleContent":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "SimpleType":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "TotalDigit":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Union":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "WhiteSpace":
					switch fieldName {
					// insertion point for enum assign code
					}
				}
			}
		}
	}
	return
}

// ReplaceOldDeclarationsInFile replaces specific text in a file at the given path.
func ReplaceOldDeclarationsInFile(pathToFile string) error {
	// Open the file for reading.
	file, err := os.Open(pathToFile)
	if err != nil {
		return err
	}
	defer file.Close()

	// replacing function with Injection
	pattern := regexp.MustCompile(`\b\w*Injection\b`)
	pattern2 := regexp.MustCompile(`\bmap_DocLink_Identifier_\w*\b`)

	// Temporary slice to hold lines from the file.
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Replace the target text with the desired text.
		line := strings.Replace(scanner.Text(), "var ___dummy__Time_stage time.Time", "var _ time.Time", -1)
		line = pattern.ReplaceAllString(line, "_")
		line = pattern2.ReplaceAllString(line, "_")

		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	// Re-open the file for writing.
	file, err = os.Create(pathToFile)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the modified lines back to the file.
	writer := bufio.NewWriter(file)
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	return writer.Flush()
}
