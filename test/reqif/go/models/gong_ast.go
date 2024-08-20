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
var __gong__map_ALTERNATIVE_ID = make(map[string]*ALTERNATIVE_ID)
var __gong__map_ATTRIBUTE_DEFINITION_BOOLEAN = make(map[string]*ATTRIBUTE_DEFINITION_BOOLEAN)
var __gong__map_ATTRIBUTE_DEFINITION_DATE = make(map[string]*ATTRIBUTE_DEFINITION_DATE)
var __gong__map_ATTRIBUTE_DEFINITION_ENUMERATION = make(map[string]*ATTRIBUTE_DEFINITION_ENUMERATION)
var __gong__map_ATTRIBUTE_DEFINITION_INTEGER = make(map[string]*ATTRIBUTE_DEFINITION_INTEGER)
var __gong__map_ATTRIBUTE_DEFINITION_REAL = make(map[string]*ATTRIBUTE_DEFINITION_REAL)
var __gong__map_ATTRIBUTE_DEFINITION_STRING = make(map[string]*ATTRIBUTE_DEFINITION_STRING)
var __gong__map_ATTRIBUTE_DEFINITION_XHTML = make(map[string]*ATTRIBUTE_DEFINITION_XHTML)
var __gong__map_ATTRIBUTE_VALUE_BOOLEAN = make(map[string]*ATTRIBUTE_VALUE_BOOLEAN)
var __gong__map_ATTRIBUTE_VALUE_DATE = make(map[string]*ATTRIBUTE_VALUE_DATE)
var __gong__map_ATTRIBUTE_VALUE_ENUMERATION = make(map[string]*ATTRIBUTE_VALUE_ENUMERATION)
var __gong__map_ATTRIBUTE_VALUE_INTEGER = make(map[string]*ATTRIBUTE_VALUE_INTEGER)
var __gong__map_ATTRIBUTE_VALUE_REAL = make(map[string]*ATTRIBUTE_VALUE_REAL)
var __gong__map_ATTRIBUTE_VALUE_STRING = make(map[string]*ATTRIBUTE_VALUE_STRING)
var __gong__map_ATTRIBUTE_VALUE_XHTML = make(map[string]*ATTRIBUTE_VALUE_XHTML)
var __gong__map_DATATYPE_DEFINITION_BOOLEAN = make(map[string]*DATATYPE_DEFINITION_BOOLEAN)
var __gong__map_DATATYPE_DEFINITION_DATE = make(map[string]*DATATYPE_DEFINITION_DATE)
var __gong__map_DATATYPE_DEFINITION_ENUMERATION = make(map[string]*DATATYPE_DEFINITION_ENUMERATION)
var __gong__map_DATATYPE_DEFINITION_INTEGER = make(map[string]*DATATYPE_DEFINITION_INTEGER)
var __gong__map_DATATYPE_DEFINITION_REAL = make(map[string]*DATATYPE_DEFINITION_REAL)
var __gong__map_DATATYPE_DEFINITION_STRING = make(map[string]*DATATYPE_DEFINITION_STRING)
var __gong__map_DATATYPE_DEFINITION_XHTML = make(map[string]*DATATYPE_DEFINITION_XHTML)
var __gong__map_EMBEDDED_VALUE = make(map[string]*EMBEDDED_VALUE)
var __gong__map_ENUM_VALUE = make(map[string]*ENUM_VALUE)
var __gong__map_RELATION_GROUP = make(map[string]*RELATION_GROUP)
var __gong__map_RELATION_GROUP_TYPE = make(map[string]*RELATION_GROUP_TYPE)
var __gong__map_REQ_IF = make(map[string]*REQ_IF)
var __gong__map_REQ_IF_CONTENT = make(map[string]*REQ_IF_CONTENT)
var __gong__map_REQ_IF_HEADER = make(map[string]*REQ_IF_HEADER)
var __gong__map_REQ_IF_TOOL_EXTENSION = make(map[string]*REQ_IF_TOOL_EXTENSION)
var __gong__map_SPECIFICATION = make(map[string]*SPECIFICATION)
var __gong__map_SPECIFICATION_TYPE = make(map[string]*SPECIFICATION_TYPE)
var __gong__map_SPEC_HIERARCHY = make(map[string]*SPEC_HIERARCHY)
var __gong__map_SPEC_OBJECT = make(map[string]*SPEC_OBJECT)
var __gong__map_SPEC_OBJECT_TYPE = make(map[string]*SPEC_OBJECT_TYPE)
var __gong__map_SPEC_RELATION = make(map[string]*SPEC_RELATION)
var __gong__map_SPEC_RELATION_TYPE = make(map[string]*SPEC_RELATION_TYPE)
var __gong__map_XHTML_CONTENT = make(map[string]*XHTML_CONTENT)

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
									case "ALTERNATIVE_ID":
										instanceALTERNATIVE_ID := (&ALTERNATIVE_ID{Name: instanceName}).Stage(stage)
										instance = any(instanceALTERNATIVE_ID)
										__gong__map_ALTERNATIVE_ID[identifier] = instanceALTERNATIVE_ID
									case "ATTRIBUTE_DEFINITION_BOOLEAN":
										instanceATTRIBUTE_DEFINITION_BOOLEAN := (&ATTRIBUTE_DEFINITION_BOOLEAN{Name: instanceName}).Stage(stage)
										instance = any(instanceATTRIBUTE_DEFINITION_BOOLEAN)
										__gong__map_ATTRIBUTE_DEFINITION_BOOLEAN[identifier] = instanceATTRIBUTE_DEFINITION_BOOLEAN
									case "ATTRIBUTE_DEFINITION_DATE":
										instanceATTRIBUTE_DEFINITION_DATE := (&ATTRIBUTE_DEFINITION_DATE{Name: instanceName}).Stage(stage)
										instance = any(instanceATTRIBUTE_DEFINITION_DATE)
										__gong__map_ATTRIBUTE_DEFINITION_DATE[identifier] = instanceATTRIBUTE_DEFINITION_DATE
									case "ATTRIBUTE_DEFINITION_ENUMERATION":
										instanceATTRIBUTE_DEFINITION_ENUMERATION := (&ATTRIBUTE_DEFINITION_ENUMERATION{Name: instanceName}).Stage(stage)
										instance = any(instanceATTRIBUTE_DEFINITION_ENUMERATION)
										__gong__map_ATTRIBUTE_DEFINITION_ENUMERATION[identifier] = instanceATTRIBUTE_DEFINITION_ENUMERATION
									case "ATTRIBUTE_DEFINITION_INTEGER":
										instanceATTRIBUTE_DEFINITION_INTEGER := (&ATTRIBUTE_DEFINITION_INTEGER{Name: instanceName}).Stage(stage)
										instance = any(instanceATTRIBUTE_DEFINITION_INTEGER)
										__gong__map_ATTRIBUTE_DEFINITION_INTEGER[identifier] = instanceATTRIBUTE_DEFINITION_INTEGER
									case "ATTRIBUTE_DEFINITION_REAL":
										instanceATTRIBUTE_DEFINITION_REAL := (&ATTRIBUTE_DEFINITION_REAL{Name: instanceName}).Stage(stage)
										instance = any(instanceATTRIBUTE_DEFINITION_REAL)
										__gong__map_ATTRIBUTE_DEFINITION_REAL[identifier] = instanceATTRIBUTE_DEFINITION_REAL
									case "ATTRIBUTE_DEFINITION_STRING":
										instanceATTRIBUTE_DEFINITION_STRING := (&ATTRIBUTE_DEFINITION_STRING{Name: instanceName}).Stage(stage)
										instance = any(instanceATTRIBUTE_DEFINITION_STRING)
										__gong__map_ATTRIBUTE_DEFINITION_STRING[identifier] = instanceATTRIBUTE_DEFINITION_STRING
									case "ATTRIBUTE_DEFINITION_XHTML":
										instanceATTRIBUTE_DEFINITION_XHTML := (&ATTRIBUTE_DEFINITION_XHTML{Name: instanceName}).Stage(stage)
										instance = any(instanceATTRIBUTE_DEFINITION_XHTML)
										__gong__map_ATTRIBUTE_DEFINITION_XHTML[identifier] = instanceATTRIBUTE_DEFINITION_XHTML
									case "ATTRIBUTE_VALUE_BOOLEAN":
										instanceATTRIBUTE_VALUE_BOOLEAN := (&ATTRIBUTE_VALUE_BOOLEAN{Name: instanceName}).Stage(stage)
										instance = any(instanceATTRIBUTE_VALUE_BOOLEAN)
										__gong__map_ATTRIBUTE_VALUE_BOOLEAN[identifier] = instanceATTRIBUTE_VALUE_BOOLEAN
									case "ATTRIBUTE_VALUE_DATE":
										instanceATTRIBUTE_VALUE_DATE := (&ATTRIBUTE_VALUE_DATE{Name: instanceName}).Stage(stage)
										instance = any(instanceATTRIBUTE_VALUE_DATE)
										__gong__map_ATTRIBUTE_VALUE_DATE[identifier] = instanceATTRIBUTE_VALUE_DATE
									case "ATTRIBUTE_VALUE_ENUMERATION":
										instanceATTRIBUTE_VALUE_ENUMERATION := (&ATTRIBUTE_VALUE_ENUMERATION{Name: instanceName}).Stage(stage)
										instance = any(instanceATTRIBUTE_VALUE_ENUMERATION)
										__gong__map_ATTRIBUTE_VALUE_ENUMERATION[identifier] = instanceATTRIBUTE_VALUE_ENUMERATION
									case "ATTRIBUTE_VALUE_INTEGER":
										instanceATTRIBUTE_VALUE_INTEGER := (&ATTRIBUTE_VALUE_INTEGER{Name: instanceName}).Stage(stage)
										instance = any(instanceATTRIBUTE_VALUE_INTEGER)
										__gong__map_ATTRIBUTE_VALUE_INTEGER[identifier] = instanceATTRIBUTE_VALUE_INTEGER
									case "ATTRIBUTE_VALUE_REAL":
										instanceATTRIBUTE_VALUE_REAL := (&ATTRIBUTE_VALUE_REAL{Name: instanceName}).Stage(stage)
										instance = any(instanceATTRIBUTE_VALUE_REAL)
										__gong__map_ATTRIBUTE_VALUE_REAL[identifier] = instanceATTRIBUTE_VALUE_REAL
									case "ATTRIBUTE_VALUE_STRING":
										instanceATTRIBUTE_VALUE_STRING := (&ATTRIBUTE_VALUE_STRING{Name: instanceName}).Stage(stage)
										instance = any(instanceATTRIBUTE_VALUE_STRING)
										__gong__map_ATTRIBUTE_VALUE_STRING[identifier] = instanceATTRIBUTE_VALUE_STRING
									case "ATTRIBUTE_VALUE_XHTML":
										instanceATTRIBUTE_VALUE_XHTML := (&ATTRIBUTE_VALUE_XHTML{Name: instanceName}).Stage(stage)
										instance = any(instanceATTRIBUTE_VALUE_XHTML)
										__gong__map_ATTRIBUTE_VALUE_XHTML[identifier] = instanceATTRIBUTE_VALUE_XHTML
									case "DATATYPE_DEFINITION_BOOLEAN":
										instanceDATATYPE_DEFINITION_BOOLEAN := (&DATATYPE_DEFINITION_BOOLEAN{Name: instanceName}).Stage(stage)
										instance = any(instanceDATATYPE_DEFINITION_BOOLEAN)
										__gong__map_DATATYPE_DEFINITION_BOOLEAN[identifier] = instanceDATATYPE_DEFINITION_BOOLEAN
									case "DATATYPE_DEFINITION_DATE":
										instanceDATATYPE_DEFINITION_DATE := (&DATATYPE_DEFINITION_DATE{Name: instanceName}).Stage(stage)
										instance = any(instanceDATATYPE_DEFINITION_DATE)
										__gong__map_DATATYPE_DEFINITION_DATE[identifier] = instanceDATATYPE_DEFINITION_DATE
									case "DATATYPE_DEFINITION_ENUMERATION":
										instanceDATATYPE_DEFINITION_ENUMERATION := (&DATATYPE_DEFINITION_ENUMERATION{Name: instanceName}).Stage(stage)
										instance = any(instanceDATATYPE_DEFINITION_ENUMERATION)
										__gong__map_DATATYPE_DEFINITION_ENUMERATION[identifier] = instanceDATATYPE_DEFINITION_ENUMERATION
									case "DATATYPE_DEFINITION_INTEGER":
										instanceDATATYPE_DEFINITION_INTEGER := (&DATATYPE_DEFINITION_INTEGER{Name: instanceName}).Stage(stage)
										instance = any(instanceDATATYPE_DEFINITION_INTEGER)
										__gong__map_DATATYPE_DEFINITION_INTEGER[identifier] = instanceDATATYPE_DEFINITION_INTEGER
									case "DATATYPE_DEFINITION_REAL":
										instanceDATATYPE_DEFINITION_REAL := (&DATATYPE_DEFINITION_REAL{Name: instanceName}).Stage(stage)
										instance = any(instanceDATATYPE_DEFINITION_REAL)
										__gong__map_DATATYPE_DEFINITION_REAL[identifier] = instanceDATATYPE_DEFINITION_REAL
									case "DATATYPE_DEFINITION_STRING":
										instanceDATATYPE_DEFINITION_STRING := (&DATATYPE_DEFINITION_STRING{Name: instanceName}).Stage(stage)
										instance = any(instanceDATATYPE_DEFINITION_STRING)
										__gong__map_DATATYPE_DEFINITION_STRING[identifier] = instanceDATATYPE_DEFINITION_STRING
									case "DATATYPE_DEFINITION_XHTML":
										instanceDATATYPE_DEFINITION_XHTML := (&DATATYPE_DEFINITION_XHTML{Name: instanceName}).Stage(stage)
										instance = any(instanceDATATYPE_DEFINITION_XHTML)
										__gong__map_DATATYPE_DEFINITION_XHTML[identifier] = instanceDATATYPE_DEFINITION_XHTML
									case "EMBEDDED_VALUE":
										instanceEMBEDDED_VALUE := (&EMBEDDED_VALUE{Name: instanceName}).Stage(stage)
										instance = any(instanceEMBEDDED_VALUE)
										__gong__map_EMBEDDED_VALUE[identifier] = instanceEMBEDDED_VALUE
									case "ENUM_VALUE":
										instanceENUM_VALUE := (&ENUM_VALUE{Name: instanceName}).Stage(stage)
										instance = any(instanceENUM_VALUE)
										__gong__map_ENUM_VALUE[identifier] = instanceENUM_VALUE
									case "RELATION_GROUP":
										instanceRELATION_GROUP := (&RELATION_GROUP{Name: instanceName}).Stage(stage)
										instance = any(instanceRELATION_GROUP)
										__gong__map_RELATION_GROUP[identifier] = instanceRELATION_GROUP
									case "RELATION_GROUP_TYPE":
										instanceRELATION_GROUP_TYPE := (&RELATION_GROUP_TYPE{Name: instanceName}).Stage(stage)
										instance = any(instanceRELATION_GROUP_TYPE)
										__gong__map_RELATION_GROUP_TYPE[identifier] = instanceRELATION_GROUP_TYPE
									case "REQ_IF":
										instanceREQ_IF := (&REQ_IF{Name: instanceName}).Stage(stage)
										instance = any(instanceREQ_IF)
										__gong__map_REQ_IF[identifier] = instanceREQ_IF
									case "REQ_IF_CONTENT":
										instanceREQ_IF_CONTENT := (&REQ_IF_CONTENT{Name: instanceName}).Stage(stage)
										instance = any(instanceREQ_IF_CONTENT)
										__gong__map_REQ_IF_CONTENT[identifier] = instanceREQ_IF_CONTENT
									case "REQ_IF_HEADER":
										instanceREQ_IF_HEADER := (&REQ_IF_HEADER{Name: instanceName}).Stage(stage)
										instance = any(instanceREQ_IF_HEADER)
										__gong__map_REQ_IF_HEADER[identifier] = instanceREQ_IF_HEADER
									case "REQ_IF_TOOL_EXTENSION":
										instanceREQ_IF_TOOL_EXTENSION := (&REQ_IF_TOOL_EXTENSION{Name: instanceName}).Stage(stage)
										instance = any(instanceREQ_IF_TOOL_EXTENSION)
										__gong__map_REQ_IF_TOOL_EXTENSION[identifier] = instanceREQ_IF_TOOL_EXTENSION
									case "SPECIFICATION":
										instanceSPECIFICATION := (&SPECIFICATION{Name: instanceName}).Stage(stage)
										instance = any(instanceSPECIFICATION)
										__gong__map_SPECIFICATION[identifier] = instanceSPECIFICATION
									case "SPECIFICATION_TYPE":
										instanceSPECIFICATION_TYPE := (&SPECIFICATION_TYPE{Name: instanceName}).Stage(stage)
										instance = any(instanceSPECIFICATION_TYPE)
										__gong__map_SPECIFICATION_TYPE[identifier] = instanceSPECIFICATION_TYPE
									case "SPEC_HIERARCHY":
										instanceSPEC_HIERARCHY := (&SPEC_HIERARCHY{Name: instanceName}).Stage(stage)
										instance = any(instanceSPEC_HIERARCHY)
										__gong__map_SPEC_HIERARCHY[identifier] = instanceSPEC_HIERARCHY
									case "SPEC_OBJECT":
										instanceSPEC_OBJECT := (&SPEC_OBJECT{Name: instanceName}).Stage(stage)
										instance = any(instanceSPEC_OBJECT)
										__gong__map_SPEC_OBJECT[identifier] = instanceSPEC_OBJECT
									case "SPEC_OBJECT_TYPE":
										instanceSPEC_OBJECT_TYPE := (&SPEC_OBJECT_TYPE{Name: instanceName}).Stage(stage)
										instance = any(instanceSPEC_OBJECT_TYPE)
										__gong__map_SPEC_OBJECT_TYPE[identifier] = instanceSPEC_OBJECT_TYPE
									case "SPEC_RELATION":
										instanceSPEC_RELATION := (&SPEC_RELATION{Name: instanceName}).Stage(stage)
										instance = any(instanceSPEC_RELATION)
										__gong__map_SPEC_RELATION[identifier] = instanceSPEC_RELATION
									case "SPEC_RELATION_TYPE":
										instanceSPEC_RELATION_TYPE := (&SPEC_RELATION_TYPE{Name: instanceName}).Stage(stage)
										instance = any(instanceSPEC_RELATION_TYPE)
										__gong__map_SPEC_RELATION_TYPE[identifier] = instanceSPEC_RELATION_TYPE
									case "XHTML_CONTENT":
										instanceXHTML_CONTENT := (&XHTML_CONTENT{Name: instanceName}).Stage(stage)
										instance = any(instanceXHTML_CONTENT)
										__gong__map_XHTML_CONTENT[identifier] = instanceXHTML_CONTENT
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
						case "ALTERNATIVE_ID":
							switch fieldName {
							// insertion point for date assign code
							}
						case "ATTRIBUTE_DEFINITION_BOOLEAN":
							switch fieldName {
							// insertion point for date assign code
							}
						case "ATTRIBUTE_DEFINITION_DATE":
							switch fieldName {
							// insertion point for date assign code
							}
						case "ATTRIBUTE_DEFINITION_ENUMERATION":
							switch fieldName {
							// insertion point for date assign code
							}
						case "ATTRIBUTE_DEFINITION_INTEGER":
							switch fieldName {
							// insertion point for date assign code
							}
						case "ATTRIBUTE_DEFINITION_REAL":
							switch fieldName {
							// insertion point for date assign code
							}
						case "ATTRIBUTE_DEFINITION_STRING":
							switch fieldName {
							// insertion point for date assign code
							}
						case "ATTRIBUTE_DEFINITION_XHTML":
							switch fieldName {
							// insertion point for date assign code
							}
						case "ATTRIBUTE_VALUE_BOOLEAN":
							switch fieldName {
							// insertion point for date assign code
							}
						case "ATTRIBUTE_VALUE_DATE":
							switch fieldName {
							// insertion point for date assign code
							}
						case "ATTRIBUTE_VALUE_ENUMERATION":
							switch fieldName {
							// insertion point for date assign code
							}
						case "ATTRIBUTE_VALUE_INTEGER":
							switch fieldName {
							// insertion point for date assign code
							}
						case "ATTRIBUTE_VALUE_REAL":
							switch fieldName {
							// insertion point for date assign code
							}
						case "ATTRIBUTE_VALUE_STRING":
							switch fieldName {
							// insertion point for date assign code
							}
						case "ATTRIBUTE_VALUE_XHTML":
							switch fieldName {
							// insertion point for date assign code
							}
						case "DATATYPE_DEFINITION_BOOLEAN":
							switch fieldName {
							// insertion point for date assign code
							}
						case "DATATYPE_DEFINITION_DATE":
							switch fieldName {
							// insertion point for date assign code
							}
						case "DATATYPE_DEFINITION_ENUMERATION":
							switch fieldName {
							// insertion point for date assign code
							}
						case "DATATYPE_DEFINITION_INTEGER":
							switch fieldName {
							// insertion point for date assign code
							}
						case "DATATYPE_DEFINITION_REAL":
							switch fieldName {
							// insertion point for date assign code
							}
						case "DATATYPE_DEFINITION_STRING":
							switch fieldName {
							// insertion point for date assign code
							}
						case "DATATYPE_DEFINITION_XHTML":
							switch fieldName {
							// insertion point for date assign code
							}
						case "EMBEDDED_VALUE":
							switch fieldName {
							// insertion point for date assign code
							}
						case "ENUM_VALUE":
							switch fieldName {
							// insertion point for date assign code
							}
						case "RELATION_GROUP":
							switch fieldName {
							// insertion point for date assign code
							}
						case "RELATION_GROUP_TYPE":
							switch fieldName {
							// insertion point for date assign code
							}
						case "REQ_IF":
							switch fieldName {
							// insertion point for date assign code
							}
						case "REQ_IF_CONTENT":
							switch fieldName {
							// insertion point for date assign code
							}
						case "REQ_IF_HEADER":
							switch fieldName {
							// insertion point for date assign code
							}
						case "REQ_IF_TOOL_EXTENSION":
							switch fieldName {
							// insertion point for date assign code
							}
						case "SPECIFICATION":
							switch fieldName {
							// insertion point for date assign code
							}
						case "SPECIFICATION_TYPE":
							switch fieldName {
							// insertion point for date assign code
							}
						case "SPEC_HIERARCHY":
							switch fieldName {
							// insertion point for date assign code
							}
						case "SPEC_OBJECT":
							switch fieldName {
							// insertion point for date assign code
							}
						case "SPEC_OBJECT_TYPE":
							switch fieldName {
							// insertion point for date assign code
							}
						case "SPEC_RELATION":
							switch fieldName {
							// insertion point for date assign code
							}
						case "SPEC_RELATION_TYPE":
							switch fieldName {
							// insertion point for date assign code
							}
						case "XHTML_CONTENT":
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
					case "ALTERNATIVE_ID":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "ATTRIBUTE_DEFINITION_BOOLEAN":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "ATTRIBUTE_DEFINITION_DATE":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "ATTRIBUTE_DEFINITION_ENUMERATION":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "ATTRIBUTE_DEFINITION_INTEGER":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "ATTRIBUTE_DEFINITION_REAL":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "ATTRIBUTE_DEFINITION_STRING":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "ATTRIBUTE_DEFINITION_XHTML":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "ATTRIBUTE_VALUE_BOOLEAN":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "ATTRIBUTE_VALUE_DATE":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "ATTRIBUTE_VALUE_ENUMERATION":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "ATTRIBUTE_VALUE_INTEGER":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "ATTRIBUTE_VALUE_REAL":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "ATTRIBUTE_VALUE_STRING":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "ATTRIBUTE_VALUE_XHTML":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "THE_VALUE":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_XHTML_CONTENT[targetIdentifier]
							__gong__map_ATTRIBUTE_VALUE_XHTML[identifier].THE_VALUE =
								append(__gong__map_ATTRIBUTE_VALUE_XHTML[identifier].THE_VALUE, target)
						case "THE_ORIGINAL_VALUE":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_XHTML_CONTENT[targetIdentifier]
							__gong__map_ATTRIBUTE_VALUE_XHTML[identifier].THE_ORIGINAL_VALUE =
								append(__gong__map_ATTRIBUTE_VALUE_XHTML[identifier].THE_ORIGINAL_VALUE, target)
						}
					case "DATATYPE_DEFINITION_BOOLEAN":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "DATATYPE_DEFINITION_DATE":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "DATATYPE_DEFINITION_ENUMERATION":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "DATATYPE_DEFINITION_INTEGER":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "DATATYPE_DEFINITION_REAL":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "DATATYPE_DEFINITION_STRING":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "DATATYPE_DEFINITION_XHTML":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "EMBEDDED_VALUE":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "ENUM_VALUE":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "RELATION_GROUP":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "RELATION_GROUP_TYPE":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "REQ_IF":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "REQ_IF_CONTENT":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "REQ_IF_HEADER":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "REQ_IF_TOOL_EXTENSION":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "SPECIFICATION":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "SPECIFICATION_TYPE":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "SPEC_HIERARCHY":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "SPEC_OBJECT":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "SPEC_OBJECT_TYPE":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "SPEC_RELATION":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "SPEC_RELATION_TYPE":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "XHTML_CONTENT":
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
			case "ALTERNATIVE_ID":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ALTERNATIVE_ID[identifier].Name = fielValue
				case "IDENTIFIER":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ALTERNATIVE_ID[identifier].IDENTIFIER = fielValue
				}
			case "ATTRIBUTE_DEFINITION_BOOLEAN":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_DEFINITION_BOOLEAN[identifier].Name = fielValue
				case "DESC":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_DEFINITION_BOOLEAN[identifier].DESC = fielValue
				case "IDENTIFIER":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_DEFINITION_BOOLEAN[identifier].IDENTIFIER = fielValue
				case "LAST_CHANGE":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_DEFINITION_BOOLEAN[identifier].LAST_CHANGE = fielValue
				case "LONG_NAME":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_DEFINITION_BOOLEAN[identifier].LONG_NAME = fielValue
				}
			case "ATTRIBUTE_DEFINITION_DATE":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_DEFINITION_DATE[identifier].Name = fielValue
				case "DESC":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_DEFINITION_DATE[identifier].DESC = fielValue
				case "IDENTIFIER":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_DEFINITION_DATE[identifier].IDENTIFIER = fielValue
				case "LAST_CHANGE":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_DEFINITION_DATE[identifier].LAST_CHANGE = fielValue
				case "LONG_NAME":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_DEFINITION_DATE[identifier].LONG_NAME = fielValue
				}
			case "ATTRIBUTE_DEFINITION_ENUMERATION":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_DEFINITION_ENUMERATION[identifier].Name = fielValue
				case "DESC":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_DEFINITION_ENUMERATION[identifier].DESC = fielValue
				case "IDENTIFIER":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_DEFINITION_ENUMERATION[identifier].IDENTIFIER = fielValue
				case "LAST_CHANGE":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_DEFINITION_ENUMERATION[identifier].LAST_CHANGE = fielValue
				case "LONG_NAME":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_DEFINITION_ENUMERATION[identifier].LONG_NAME = fielValue
				}
			case "ATTRIBUTE_DEFINITION_INTEGER":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_DEFINITION_INTEGER[identifier].Name = fielValue
				case "DESC":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_DEFINITION_INTEGER[identifier].DESC = fielValue
				case "IDENTIFIER":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_DEFINITION_INTEGER[identifier].IDENTIFIER = fielValue
				case "LAST_CHANGE":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_DEFINITION_INTEGER[identifier].LAST_CHANGE = fielValue
				case "LONG_NAME":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_DEFINITION_INTEGER[identifier].LONG_NAME = fielValue
				}
			case "ATTRIBUTE_DEFINITION_REAL":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_DEFINITION_REAL[identifier].Name = fielValue
				case "DESC":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_DEFINITION_REAL[identifier].DESC = fielValue
				case "IDENTIFIER":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_DEFINITION_REAL[identifier].IDENTIFIER = fielValue
				case "LAST_CHANGE":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_DEFINITION_REAL[identifier].LAST_CHANGE = fielValue
				case "LONG_NAME":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_DEFINITION_REAL[identifier].LONG_NAME = fielValue
				}
			case "ATTRIBUTE_DEFINITION_STRING":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_DEFINITION_STRING[identifier].Name = fielValue
				case "DESC":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_DEFINITION_STRING[identifier].DESC = fielValue
				case "IDENTIFIER":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_DEFINITION_STRING[identifier].IDENTIFIER = fielValue
				case "LAST_CHANGE":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_DEFINITION_STRING[identifier].LAST_CHANGE = fielValue
				case "LONG_NAME":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_DEFINITION_STRING[identifier].LONG_NAME = fielValue
				}
			case "ATTRIBUTE_DEFINITION_XHTML":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_DEFINITION_XHTML[identifier].Name = fielValue
				case "DESC":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_DEFINITION_XHTML[identifier].DESC = fielValue
				case "IDENTIFIER":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_DEFINITION_XHTML[identifier].IDENTIFIER = fielValue
				case "LAST_CHANGE":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_DEFINITION_XHTML[identifier].LAST_CHANGE = fielValue
				case "LONG_NAME":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_DEFINITION_XHTML[identifier].LONG_NAME = fielValue
				}
			case "ATTRIBUTE_VALUE_BOOLEAN":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_VALUE_BOOLEAN[identifier].Name = fielValue
				}
			case "ATTRIBUTE_VALUE_DATE":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_VALUE_DATE[identifier].Name = fielValue
				case "THE_VALUE":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_VALUE_DATE[identifier].THE_VALUE = fielValue
				}
			case "ATTRIBUTE_VALUE_ENUMERATION":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_VALUE_ENUMERATION[identifier].Name = fielValue
				}
			case "ATTRIBUTE_VALUE_INTEGER":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_VALUE_INTEGER[identifier].Name = fielValue
				case "THE_VALUE":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_ATTRIBUTE_VALUE_INTEGER[identifier].THE_VALUE = int(exprSign) * int(fielValue)
				}
			case "ATTRIBUTE_VALUE_REAL":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_VALUE_REAL[identifier].Name = fielValue
				case "THE_VALUE":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_ATTRIBUTE_VALUE_REAL[identifier].THE_VALUE = exprSign * fielValue
				}
			case "ATTRIBUTE_VALUE_STRING":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_VALUE_STRING[identifier].Name = fielValue
				case "THE_VALUE":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_VALUE_STRING[identifier].THE_VALUE = fielValue
				}
			case "ATTRIBUTE_VALUE_XHTML":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ATTRIBUTE_VALUE_XHTML[identifier].Name = fielValue
				}
			case "DATATYPE_DEFINITION_BOOLEAN":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPE_DEFINITION_BOOLEAN[identifier].Name = fielValue
				case "DESC":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPE_DEFINITION_BOOLEAN[identifier].DESC = fielValue
				case "IDENTIFIER":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPE_DEFINITION_BOOLEAN[identifier].IDENTIFIER = fielValue
				case "LAST_CHANGE":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPE_DEFINITION_BOOLEAN[identifier].LAST_CHANGE = fielValue
				case "LONG_NAME":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPE_DEFINITION_BOOLEAN[identifier].LONG_NAME = fielValue
				}
			case "DATATYPE_DEFINITION_DATE":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPE_DEFINITION_DATE[identifier].Name = fielValue
				case "DESC":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPE_DEFINITION_DATE[identifier].DESC = fielValue
				case "IDENTIFIER":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPE_DEFINITION_DATE[identifier].IDENTIFIER = fielValue
				case "LAST_CHANGE":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPE_DEFINITION_DATE[identifier].LAST_CHANGE = fielValue
				case "LONG_NAME":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPE_DEFINITION_DATE[identifier].LONG_NAME = fielValue
				}
			case "DATATYPE_DEFINITION_ENUMERATION":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPE_DEFINITION_ENUMERATION[identifier].Name = fielValue
				case "DESC":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPE_DEFINITION_ENUMERATION[identifier].DESC = fielValue
				case "IDENTIFIER":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPE_DEFINITION_ENUMERATION[identifier].IDENTIFIER = fielValue
				case "LAST_CHANGE":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPE_DEFINITION_ENUMERATION[identifier].LAST_CHANGE = fielValue
				case "LONG_NAME":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPE_DEFINITION_ENUMERATION[identifier].LONG_NAME = fielValue
				}
			case "DATATYPE_DEFINITION_INTEGER":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPE_DEFINITION_INTEGER[identifier].Name = fielValue
				case "DESC":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPE_DEFINITION_INTEGER[identifier].DESC = fielValue
				case "IDENTIFIER":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPE_DEFINITION_INTEGER[identifier].IDENTIFIER = fielValue
				case "LAST_CHANGE":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPE_DEFINITION_INTEGER[identifier].LAST_CHANGE = fielValue
				case "LONG_NAME":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPE_DEFINITION_INTEGER[identifier].LONG_NAME = fielValue
				case "MAX":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_DATATYPE_DEFINITION_INTEGER[identifier].MAX = int(exprSign) * int(fielValue)
				case "MIN":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_DATATYPE_DEFINITION_INTEGER[identifier].MIN = int(exprSign) * int(fielValue)
				}
			case "DATATYPE_DEFINITION_REAL":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPE_DEFINITION_REAL[identifier].Name = fielValue
				case "ACCURACY":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_DATATYPE_DEFINITION_REAL[identifier].ACCURACY = int(exprSign) * int(fielValue)
				case "DESC":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPE_DEFINITION_REAL[identifier].DESC = fielValue
				case "IDENTIFIER":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPE_DEFINITION_REAL[identifier].IDENTIFIER = fielValue
				case "LAST_CHANGE":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPE_DEFINITION_REAL[identifier].LAST_CHANGE = fielValue
				case "LONG_NAME":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPE_DEFINITION_REAL[identifier].LONG_NAME = fielValue
				case "MAX":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_DATATYPE_DEFINITION_REAL[identifier].MAX = exprSign * fielValue
				case "MIN":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_DATATYPE_DEFINITION_REAL[identifier].MIN = exprSign * fielValue
				}
			case "DATATYPE_DEFINITION_STRING":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPE_DEFINITION_STRING[identifier].Name = fielValue
				case "DESC":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPE_DEFINITION_STRING[identifier].DESC = fielValue
				case "IDENTIFIER":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPE_DEFINITION_STRING[identifier].IDENTIFIER = fielValue
				case "LAST_CHANGE":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPE_DEFINITION_STRING[identifier].LAST_CHANGE = fielValue
				case "LONG_NAME":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPE_DEFINITION_STRING[identifier].LONG_NAME = fielValue
				case "MAX_LENGTH":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_DATATYPE_DEFINITION_STRING[identifier].MAX_LENGTH = int(exprSign) * int(fielValue)
				}
			case "DATATYPE_DEFINITION_XHTML":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPE_DEFINITION_XHTML[identifier].Name = fielValue
				case "DESC":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPE_DEFINITION_XHTML[identifier].DESC = fielValue
				case "IDENTIFIER":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPE_DEFINITION_XHTML[identifier].IDENTIFIER = fielValue
				case "LAST_CHANGE":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPE_DEFINITION_XHTML[identifier].LAST_CHANGE = fielValue
				case "LONG_NAME":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_DATATYPE_DEFINITION_XHTML[identifier].LONG_NAME = fielValue
				}
			case "EMBEDDED_VALUE":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_EMBEDDED_VALUE[identifier].Name = fielValue
				case "KEY":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_EMBEDDED_VALUE[identifier].KEY = int(exprSign) * int(fielValue)
				case "OTHER_CONTENT":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_EMBEDDED_VALUE[identifier].OTHER_CONTENT = fielValue
				}
			case "ENUM_VALUE":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ENUM_VALUE[identifier].Name = fielValue
				case "DESC":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ENUM_VALUE[identifier].DESC = fielValue
				case "IDENTIFIER":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ENUM_VALUE[identifier].IDENTIFIER = fielValue
				case "LAST_CHANGE":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ENUM_VALUE[identifier].LAST_CHANGE = fielValue
				case "LONG_NAME":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ENUM_VALUE[identifier].LONG_NAME = fielValue
				}
			case "RELATION_GROUP":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_RELATION_GROUP[identifier].Name = fielValue
				case "DESC":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_RELATION_GROUP[identifier].DESC = fielValue
				case "IDENTIFIER":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_RELATION_GROUP[identifier].IDENTIFIER = fielValue
				case "LAST_CHANGE":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_RELATION_GROUP[identifier].LAST_CHANGE = fielValue
				case "LONG_NAME":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_RELATION_GROUP[identifier].LONG_NAME = fielValue
				}
			case "RELATION_GROUP_TYPE":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_RELATION_GROUP_TYPE[identifier].Name = fielValue
				case "DESC":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_RELATION_GROUP_TYPE[identifier].DESC = fielValue
				case "IDENTIFIER":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_RELATION_GROUP_TYPE[identifier].IDENTIFIER = fielValue
				case "LAST_CHANGE":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_RELATION_GROUP_TYPE[identifier].LAST_CHANGE = fielValue
				case "LONG_NAME":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_RELATION_GROUP_TYPE[identifier].LONG_NAME = fielValue
				}
			case "REQ_IF":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_REQ_IF[identifier].Name = fielValue
				case "Lang":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_REQ_IF[identifier].Lang = fielValue
				}
			case "REQ_IF_CONTENT":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_REQ_IF_CONTENT[identifier].Name = fielValue
				}
			case "REQ_IF_HEADER":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_REQ_IF_HEADER[identifier].Name = fielValue
				case "IDENTIFIER":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_REQ_IF_HEADER[identifier].IDENTIFIER = fielValue
				case "COMMENT":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_REQ_IF_HEADER[identifier].COMMENT = fielValue
				case "CREATION_TIME":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_REQ_IF_HEADER[identifier].CREATION_TIME = fielValue
				case "REPOSITORY_ID":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_REQ_IF_HEADER[identifier].REPOSITORY_ID = fielValue
				case "REQ_IF_TOOL_ID":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_REQ_IF_HEADER[identifier].REQ_IF_TOOL_ID = fielValue
				case "REQ_IF_VERSION":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_REQ_IF_HEADER[identifier].REQ_IF_VERSION = fielValue
				case "SOURCE_TOOL_ID":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_REQ_IF_HEADER[identifier].SOURCE_TOOL_ID = fielValue
				case "TITLE":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_REQ_IF_HEADER[identifier].TITLE = fielValue
				}
			case "REQ_IF_TOOL_EXTENSION":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_REQ_IF_TOOL_EXTENSION[identifier].Name = fielValue
				}
			case "SPECIFICATION":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPECIFICATION[identifier].Name = fielValue
				case "DESC":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPECIFICATION[identifier].DESC = fielValue
				case "IDENTIFIER":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPECIFICATION[identifier].IDENTIFIER = fielValue
				case "LAST_CHANGE":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPECIFICATION[identifier].LAST_CHANGE = fielValue
				case "LONG_NAME":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPECIFICATION[identifier].LONG_NAME = fielValue
				}
			case "SPECIFICATION_TYPE":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPECIFICATION_TYPE[identifier].Name = fielValue
				case "DESC":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPECIFICATION_TYPE[identifier].DESC = fielValue
				case "IDENTIFIER":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPECIFICATION_TYPE[identifier].IDENTIFIER = fielValue
				case "LAST_CHANGE":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPECIFICATION_TYPE[identifier].LAST_CHANGE = fielValue
				case "LONG_NAME":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPECIFICATION_TYPE[identifier].LONG_NAME = fielValue
				}
			case "SPEC_HIERARCHY":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPEC_HIERARCHY[identifier].Name = fielValue
				case "DESC":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPEC_HIERARCHY[identifier].DESC = fielValue
				case "IDENTIFIER":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPEC_HIERARCHY[identifier].IDENTIFIER = fielValue
				case "LAST_CHANGE":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPEC_HIERARCHY[identifier].LAST_CHANGE = fielValue
				case "LONG_NAME":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPEC_HIERARCHY[identifier].LONG_NAME = fielValue
				}
			case "SPEC_OBJECT":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPEC_OBJECT[identifier].Name = fielValue
				case "DESC":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPEC_OBJECT[identifier].DESC = fielValue
				case "IDENTIFIER":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPEC_OBJECT[identifier].IDENTIFIER = fielValue
				case "LAST_CHANGE":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPEC_OBJECT[identifier].LAST_CHANGE = fielValue
				case "LONG_NAME":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPEC_OBJECT[identifier].LONG_NAME = fielValue
				}
			case "SPEC_OBJECT_TYPE":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPEC_OBJECT_TYPE[identifier].Name = fielValue
				case "DESC":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPEC_OBJECT_TYPE[identifier].DESC = fielValue
				case "IDENTIFIER":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPEC_OBJECT_TYPE[identifier].IDENTIFIER = fielValue
				case "LAST_CHANGE":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPEC_OBJECT_TYPE[identifier].LAST_CHANGE = fielValue
				case "LONG_NAME":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPEC_OBJECT_TYPE[identifier].LONG_NAME = fielValue
				}
			case "SPEC_RELATION":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPEC_RELATION[identifier].Name = fielValue
				case "DESC":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPEC_RELATION[identifier].DESC = fielValue
				case "IDENTIFIER":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPEC_RELATION[identifier].IDENTIFIER = fielValue
				case "LAST_CHANGE":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPEC_RELATION[identifier].LAST_CHANGE = fielValue
				case "LONG_NAME":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPEC_RELATION[identifier].LONG_NAME = fielValue
				}
			case "SPEC_RELATION_TYPE":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPEC_RELATION_TYPE[identifier].Name = fielValue
				case "DESC":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPEC_RELATION_TYPE[identifier].DESC = fielValue
				case "IDENTIFIER":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPEC_RELATION_TYPE[identifier].IDENTIFIER = fielValue
				case "LAST_CHANGE":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPEC_RELATION_TYPE[identifier].LAST_CHANGE = fielValue
				case "LONG_NAME":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SPEC_RELATION_TYPE[identifier].LONG_NAME = fielValue
				}
			case "XHTML_CONTENT":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_XHTML_CONTENT[identifier].Name = fielValue
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
			case "ALTERNATIVE_ID":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "ATTRIBUTE_DEFINITION_BOOLEAN":
				switch fieldName {
				// insertion point for field dependant code
				case "IS_EDITABLE":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_ATTRIBUTE_DEFINITION_BOOLEAN[identifier].IS_EDITABLE = fielValue
				}
			case "ATTRIBUTE_DEFINITION_DATE":
				switch fieldName {
				// insertion point for field dependant code
				case "IS_EDITABLE":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_ATTRIBUTE_DEFINITION_DATE[identifier].IS_EDITABLE = fielValue
				}
			case "ATTRIBUTE_DEFINITION_ENUMERATION":
				switch fieldName {
				// insertion point for field dependant code
				case "IS_EDITABLE":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_ATTRIBUTE_DEFINITION_ENUMERATION[identifier].IS_EDITABLE = fielValue
				case "MULTI_VALUED":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_ATTRIBUTE_DEFINITION_ENUMERATION[identifier].MULTI_VALUED = fielValue
				}
			case "ATTRIBUTE_DEFINITION_INTEGER":
				switch fieldName {
				// insertion point for field dependant code
				case "IS_EDITABLE":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_ATTRIBUTE_DEFINITION_INTEGER[identifier].IS_EDITABLE = fielValue
				}
			case "ATTRIBUTE_DEFINITION_REAL":
				switch fieldName {
				// insertion point for field dependant code
				case "IS_EDITABLE":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_ATTRIBUTE_DEFINITION_REAL[identifier].IS_EDITABLE = fielValue
				}
			case "ATTRIBUTE_DEFINITION_STRING":
				switch fieldName {
				// insertion point for field dependant code
				case "IS_EDITABLE":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_ATTRIBUTE_DEFINITION_STRING[identifier].IS_EDITABLE = fielValue
				}
			case "ATTRIBUTE_DEFINITION_XHTML":
				switch fieldName {
				// insertion point for field dependant code
				case "IS_EDITABLE":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_ATTRIBUTE_DEFINITION_XHTML[identifier].IS_EDITABLE = fielValue
				}
			case "ATTRIBUTE_VALUE_BOOLEAN":
				switch fieldName {
				// insertion point for field dependant code
				case "THE_VALUE":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_ATTRIBUTE_VALUE_BOOLEAN[identifier].THE_VALUE = fielValue
				}
			case "ATTRIBUTE_VALUE_DATE":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "ATTRIBUTE_VALUE_ENUMERATION":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "ATTRIBUTE_VALUE_INTEGER":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "ATTRIBUTE_VALUE_REAL":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "ATTRIBUTE_VALUE_STRING":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "ATTRIBUTE_VALUE_XHTML":
				switch fieldName {
				// insertion point for field dependant code
				case "IS_SIMPLIFIED":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_ATTRIBUTE_VALUE_XHTML[identifier].IS_SIMPLIFIED = fielValue
				}
			case "DATATYPE_DEFINITION_BOOLEAN":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "DATATYPE_DEFINITION_DATE":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "DATATYPE_DEFINITION_ENUMERATION":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "DATATYPE_DEFINITION_INTEGER":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "DATATYPE_DEFINITION_REAL":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "DATATYPE_DEFINITION_STRING":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "DATATYPE_DEFINITION_XHTML":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "EMBEDDED_VALUE":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "ENUM_VALUE":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "RELATION_GROUP":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "RELATION_GROUP_TYPE":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "REQ_IF":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "REQ_IF_CONTENT":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "REQ_IF_HEADER":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "REQ_IF_TOOL_EXTENSION":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "SPECIFICATION":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "SPECIFICATION_TYPE":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "SPEC_HIERARCHY":
				switch fieldName {
				// insertion point for field dependant code
				case "IS_EDITABLE":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_SPEC_HIERARCHY[identifier].IS_EDITABLE = fielValue
				case "IS_TABLE_INTERNAL":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_SPEC_HIERARCHY[identifier].IS_TABLE_INTERNAL = fielValue
				}
			case "SPEC_OBJECT":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "SPEC_OBJECT_TYPE":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "SPEC_RELATION":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "SPEC_RELATION_TYPE":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "XHTML_CONTENT":
				switch fieldName {
				// insertion point for field dependant code
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
				case "ALTERNATIVE_ID":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "ATTRIBUTE_DEFINITION_BOOLEAN":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "ATTRIBUTE_DEFINITION_DATE":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "ATTRIBUTE_DEFINITION_ENUMERATION":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "ATTRIBUTE_DEFINITION_INTEGER":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "ATTRIBUTE_DEFINITION_REAL":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "ATTRIBUTE_DEFINITION_STRING":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "ATTRIBUTE_DEFINITION_XHTML":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "ATTRIBUTE_VALUE_BOOLEAN":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "ATTRIBUTE_VALUE_DATE":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "ATTRIBUTE_VALUE_ENUMERATION":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "ATTRIBUTE_VALUE_INTEGER":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "ATTRIBUTE_VALUE_REAL":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "ATTRIBUTE_VALUE_STRING":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "ATTRIBUTE_VALUE_XHTML":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "DATATYPE_DEFINITION_BOOLEAN":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "DATATYPE_DEFINITION_DATE":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "DATATYPE_DEFINITION_ENUMERATION":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "DATATYPE_DEFINITION_INTEGER":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "DATATYPE_DEFINITION_REAL":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "DATATYPE_DEFINITION_STRING":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "DATATYPE_DEFINITION_XHTML":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "EMBEDDED_VALUE":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "ENUM_VALUE":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "RELATION_GROUP":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "RELATION_GROUP_TYPE":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "REQ_IF":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "REQ_IF_CONTENT":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "REQ_IF_HEADER":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "REQ_IF_TOOL_EXTENSION":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "SPECIFICATION":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "SPECIFICATION_TYPE":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "SPEC_HIERARCHY":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "SPEC_OBJECT":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "SPEC_OBJECT_TYPE":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "SPEC_RELATION":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "SPEC_RELATION_TYPE":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "XHTML_CONTENT":
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
