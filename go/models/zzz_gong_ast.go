// generated code - do not edit
package models

import (
	"bufio"
	"embed"
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
var _ = dummy_strconv_import
var dummy_time_import time.Time
var _ = dummy_time_import

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
func ParseAstFile(stage *Stage, pathToFile string) error {

	ReplaceOldDeclarationsInFile(pathToFile)

	fileOfInterest, err := filepath.Abs(pathToFile)
	if err != nil {
		return errors.New("Path does not exist %s ;" + fileOfInterest)
	}

	// Read the file content using os.ReadFile
	content, err := os.ReadFile(fileOfInterest)
	if err != nil {
		return errors.New("Unable to read file " + err.Error())
	}

	// Assign the content to stage.contentWhenParsed
	stage.contentWhenParsed = string(content)

	fset := token.NewFileSet()
	// startParser := time.Now()
	inFile, errParser := parser.ParseFile(fset, fileOfInterest, nil, parser.ParseComments)
	// log.Printf("Parser took %s", time.Since(startParser))

	if errParser != nil {
		return errors.New("Unable to parser " + errParser.Error())
	}

	return ParseAstFileFromAst(stage, inFile, fset)
}

// ParseAstEmbeddedFile parses the Go source code from an embedded file
// specified by pathToFile within the provided embed.FS directory and
// stages instances declared in the file using the provided Stage.
//
// Parameters:
//
//	stage:      The staging area to populate.
//	directory:  The embedded filesystem containing the file.
//	pathToFile: The path to the Go source file within the embedded filesystem.
//
// Returns:
//
//	An error if reading or parsing the file fails, or if ParseAstFileFromAst fails.
func ParseAstEmbeddedFile(stage *Stage, directory embed.FS, pathToFile string) error {

	// 1. Read the content from the embedded filesystem.
	//    We don't need filepath.Abs as embed.FS uses relative paths.
	//    We also skip ReplaceOldDeclarationsInFile as embedded files are read-only.
	fileContentBytes, err := directory.ReadFile(pathToFile)
	if err != nil {
		// Return a specific error if the file can't be read from the embed.FS
		return errors.New(stage.GetName() + "; Unable to read embedded file " + err.Error())
	}

	// 2. Create a FileSet to manage position information.
	fset := token.NewFileSet()

	// 3. Parse the file content.
	//    Instead of passing a filename for the OS to read, we pass the pathToFile
	//    as the identifier and the actual file content (fileContentBytes) as the source.
	//    parser.ParseComments is included to match the original function's behavior.
	//    The type *ast.File is returned by parser.ParseFile.
	inFile, errParser := parser.ParseFile(fset, pathToFile, fileContentBytes, parser.ParseComments)
	if errParser != nil {
		// Return a specific error if parsing fails
		return errors.New("Unable to parse embedded file '" + pathToFile + "': " + errParser.Error())
	}

	// 4. Call the common AST processing logic.
	//    Pass the parsed AST (*ast.File), the FileSet, and the stage.
	return ParseAstFileFromAst(stage, inFile, fset)
}

// ParseAstFile Parse pathToFile and stages all instances
// declared in the file
func ParseAstFileFromAst(stage *Stage, inFile *ast.File, fset *token.FileSet) error {
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
					case *ast.DeclStmt:
						if genDecl, ok := stmt.Decl.(*ast.GenDecl); ok && genDecl.Tok == token.CONST {
							for _, spec := range genDecl.Specs {
								if valueSpec, ok := spec.(*ast.ValueSpec); ok {
									for i, name := range valueSpec.Names {
										if i < len(valueSpec.Values) {
											if basicLit, ok := valueSpec.Values[i].(*ast.BasicLit); ok && basicLit.Kind == token.STRING {
												// Remove quotes from string literal
												value := strings.Trim(basicLit.Value, `"`)

												switch name.Name {
												case "__commitId__":
													if parsedUint, err := strconv.ParseUint(value, 10, 64); err == nil {
														stage.commitId = uint(parsedUint)
														stage.commitIdWhenParsed = stage.commitId
													}
												}
											}
										}
									}
								}
							}
						}
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
						// cmap := ast.NewCommentMap(fset, inFile, inFile.Comments)
						var cmap ast.CommentMap
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

func lookupSym(recv, name string) bool {
	return recv == ""
}

// UnmarshallGoStaging unmarshall a go assign statement
func UnmarshallGongstructStaging(stage *Stage, cmap *ast.CommentMap, assignStmt *ast.AssignStmt, astCoordinate_ string) (
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
			var basicLit *ast.BasicLit

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
										instanceAll := new(All)
										instanceAll.Name = instanceName
										instanceAll.Stage(stage)
										instance = any(instanceAll)
										__gong__map_All[identifier] = instanceAll
									case "Annotation":
										instanceAnnotation := new(Annotation)
										instanceAnnotation.Name = instanceName
										instanceAnnotation.Stage(stage)
										instance = any(instanceAnnotation)
										__gong__map_Annotation[identifier] = instanceAnnotation
									case "Attribute":
										instanceAttribute := new(Attribute)
										instanceAttribute.Name = instanceName
										instanceAttribute.Stage(stage)
										instance = any(instanceAttribute)
										__gong__map_Attribute[identifier] = instanceAttribute
									case "AttributeGroup":
										instanceAttributeGroup := new(AttributeGroup)
										instanceAttributeGroup.Name = instanceName
										instanceAttributeGroup.Stage(stage)
										instance = any(instanceAttributeGroup)
										__gong__map_AttributeGroup[identifier] = instanceAttributeGroup
									case "Choice":
										instanceChoice := new(Choice)
										instanceChoice.Name = instanceName
										instanceChoice.Stage(stage)
										instance = any(instanceChoice)
										__gong__map_Choice[identifier] = instanceChoice
									case "ComplexContent":
										instanceComplexContent := new(ComplexContent)
										instanceComplexContent.Name = instanceName
										instanceComplexContent.Stage(stage)
										instance = any(instanceComplexContent)
										__gong__map_ComplexContent[identifier] = instanceComplexContent
									case "ComplexType":
										instanceComplexType := new(ComplexType)
										instanceComplexType.Name = instanceName
										instanceComplexType.Stage(stage)
										instance = any(instanceComplexType)
										__gong__map_ComplexType[identifier] = instanceComplexType
									case "Documentation":
										instanceDocumentation := new(Documentation)
										instanceDocumentation.Name = instanceName
										instanceDocumentation.Stage(stage)
										instance = any(instanceDocumentation)
										__gong__map_Documentation[identifier] = instanceDocumentation
									case "Element":
										instanceElement := new(Element)
										instanceElement.Name = instanceName
										instanceElement.Stage(stage)
										instance = any(instanceElement)
										__gong__map_Element[identifier] = instanceElement
									case "Enumeration":
										instanceEnumeration := new(Enumeration)
										instanceEnumeration.Name = instanceName
										instanceEnumeration.Stage(stage)
										instance = any(instanceEnumeration)
										__gong__map_Enumeration[identifier] = instanceEnumeration
									case "Extension":
										instanceExtension := new(Extension)
										instanceExtension.Name = instanceName
										instanceExtension.Stage(stage)
										instance = any(instanceExtension)
										__gong__map_Extension[identifier] = instanceExtension
									case "Group":
										instanceGroup := new(Group)
										instanceGroup.Name = instanceName
										instanceGroup.Stage(stage)
										instance = any(instanceGroup)
										__gong__map_Group[identifier] = instanceGroup
									case "Length":
										instanceLength := new(Length)
										instanceLength.Name = instanceName
										instanceLength.Stage(stage)
										instance = any(instanceLength)
										__gong__map_Length[identifier] = instanceLength
									case "MaxInclusive":
										instanceMaxInclusive := new(MaxInclusive)
										instanceMaxInclusive.Name = instanceName
										instanceMaxInclusive.Stage(stage)
										instance = any(instanceMaxInclusive)
										__gong__map_MaxInclusive[identifier] = instanceMaxInclusive
									case "MaxLength":
										instanceMaxLength := new(MaxLength)
										instanceMaxLength.Name = instanceName
										instanceMaxLength.Stage(stage)
										instance = any(instanceMaxLength)
										__gong__map_MaxLength[identifier] = instanceMaxLength
									case "MinInclusive":
										instanceMinInclusive := new(MinInclusive)
										instanceMinInclusive.Name = instanceName
										instanceMinInclusive.Stage(stage)
										instance = any(instanceMinInclusive)
										__gong__map_MinInclusive[identifier] = instanceMinInclusive
									case "MinLength":
										instanceMinLength := new(MinLength)
										instanceMinLength.Name = instanceName
										instanceMinLength.Stage(stage)
										instance = any(instanceMinLength)
										__gong__map_MinLength[identifier] = instanceMinLength
									case "Pattern":
										instancePattern := new(Pattern)
										instancePattern.Name = instanceName
										instancePattern.Stage(stage)
										instance = any(instancePattern)
										__gong__map_Pattern[identifier] = instancePattern
									case "Restriction":
										instanceRestriction := new(Restriction)
										instanceRestriction.Name = instanceName
										instanceRestriction.Stage(stage)
										instance = any(instanceRestriction)
										__gong__map_Restriction[identifier] = instanceRestriction
									case "Schema":
										instanceSchema := new(Schema)
										instanceSchema.Name = instanceName
										instanceSchema.Stage(stage)
										instance = any(instanceSchema)
										__gong__map_Schema[identifier] = instanceSchema
									case "Sequence":
										instanceSequence := new(Sequence)
										instanceSequence.Name = instanceName
										instanceSequence.Stage(stage)
										instance = any(instanceSequence)
										__gong__map_Sequence[identifier] = instanceSequence
									case "SimpleContent":
										instanceSimpleContent := new(SimpleContent)
										instanceSimpleContent.Name = instanceName
										instanceSimpleContent.Stage(stage)
										instance = any(instanceSimpleContent)
										__gong__map_SimpleContent[identifier] = instanceSimpleContent
									case "SimpleType":
										instanceSimpleType := new(SimpleType)
										instanceSimpleType.Name = instanceName
										instanceSimpleType.Stage(stage)
										instance = any(instanceSimpleType)
										__gong__map_SimpleType[identifier] = instanceSimpleType
									case "TotalDigit":
										instanceTotalDigit := new(TotalDigit)
										instanceTotalDigit.Name = instanceName
										instanceTotalDigit.Stage(stage)
										instance = any(instanceTotalDigit)
										__gong__map_TotalDigit[identifier] = instanceTotalDigit
									case "Union":
										instanceUnion := new(Union)
										instanceUnion.Name = instanceName
										instanceUnion.Stage(stage)
										instance = any(instanceUnion)
										__gong__map_Union[identifier] = instanceUnion
									case "WhiteSpace":
										instanceWhiteSpace := new(WhiteSpace)
										instanceWhiteSpace.Name = instanceName
										instanceWhiteSpace.Stage(stage)
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
							log.Println("gongstructName not found for identifier", identifier)
							break
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
				// pick up the first arg
				if len(callExpr.Args) != 1 {
					break
				}
				arg0 := callExpr.Args[0]

				var se *ast.SelectorExpr
				var ok bool
				if se, ok = arg0.(*ast.SelectorExpr); !ok {
					break
				}

				var seXident *ast.Ident
				if seXident = se.X.(*ast.Ident); !ok {
					break
				}

				basicLit = new(ast.BasicLit)
				// For a "fake" literal, Kind might be set to something like token.STRING or a custom indicator
				basicLit.Kind = token.STRING // Or another appropriate token.Kind
				basicLit.Value = "new(" + seXident.Name + "." + se.Sel.Name + ")"
				// following lines are here to avoid warning "unused write to field..."
				_ = basicLit.Kind
				_ = basicLit.Value
				_ = basicLit
			}
			for argNb, arg := range callExpr.Args {
				_ = argNb
				// astCoordinate := astCoordinate + "\tArg"
				switch arg := arg.(type) {
				case *ast.Ident, *ast.SelectorExpr:
					var ident *ast.Ident
					var ok bool
					_ = ok
					if ident, ok = arg.(*ast.Ident); !ok {
						// log.Println("we are in the case of new(....)")
					}

					var se *ast.SelectorExpr
					if se, ok = arg.(*ast.SelectorExpr); ok {
						if ident, ok = se.X.(*ast.Ident); !ok {
							// log.Println("we are in the case of append(....)")
						}
					}
					_ = ident

					gongstructName, ok = __gong__map_Indentifiers_gongstructName[identifier]
					if !ok {
						log.Println("gongstructName not found for identifier", identifier)
						break
					}
					switch gongstructName {
					// insertion point for slice of pointers assignments
					case "All":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Sequences":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Sequence[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_All[identifier]
								instanceWhoseFieldIsAppended.Sequences = append(instanceWhoseFieldIsAppended.Sequences, instanceToAppend)
							}
						case "Alls":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_All[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_All[identifier]
								instanceWhoseFieldIsAppended.Alls = append(instanceWhoseFieldIsAppended.Alls, instanceToAppend)
							}
						case "Choices":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Choice[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_All[identifier]
								instanceWhoseFieldIsAppended.Choices = append(instanceWhoseFieldIsAppended.Choices, instanceToAppend)
							}
						case "Groups":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Group[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_All[identifier]
								instanceWhoseFieldIsAppended.Groups = append(instanceWhoseFieldIsAppended.Groups, instanceToAppend)
							}
						case "Elements":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Element[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_All[identifier]
								instanceWhoseFieldIsAppended.Elements = append(instanceWhoseFieldIsAppended.Elements, instanceToAppend)
							}
						}
					case "Annotation":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Documentations":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Documentation[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_Annotation[identifier]
								instanceWhoseFieldIsAppended.Documentations = append(instanceWhoseFieldIsAppended.Documentations, instanceToAppend)
							}
						}
					case "Attribute":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "AttributeGroup":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "AttributeGroups":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_AttributeGroup[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_AttributeGroup[identifier]
								instanceWhoseFieldIsAppended.AttributeGroups = append(instanceWhoseFieldIsAppended.AttributeGroups, instanceToAppend)
							}
						case "Attributes":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Attribute[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_AttributeGroup[identifier]
								instanceWhoseFieldIsAppended.Attributes = append(instanceWhoseFieldIsAppended.Attributes, instanceToAppend)
							}
						}
					case "Choice":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Sequences":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Sequence[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_Choice[identifier]
								instanceWhoseFieldIsAppended.Sequences = append(instanceWhoseFieldIsAppended.Sequences, instanceToAppend)
							}
						case "Alls":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_All[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_Choice[identifier]
								instanceWhoseFieldIsAppended.Alls = append(instanceWhoseFieldIsAppended.Alls, instanceToAppend)
							}
						case "Choices":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Choice[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_Choice[identifier]
								instanceWhoseFieldIsAppended.Choices = append(instanceWhoseFieldIsAppended.Choices, instanceToAppend)
							}
						case "Groups":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Group[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_Choice[identifier]
								instanceWhoseFieldIsAppended.Groups = append(instanceWhoseFieldIsAppended.Groups, instanceToAppend)
							}
						case "Elements":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Element[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_Choice[identifier]
								instanceWhoseFieldIsAppended.Elements = append(instanceWhoseFieldIsAppended.Elements, instanceToAppend)
							}
						}
					case "ComplexContent":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "ComplexType":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Sequences":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Sequence[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_ComplexType[identifier]
								instanceWhoseFieldIsAppended.Sequences = append(instanceWhoseFieldIsAppended.Sequences, instanceToAppend)
							}
						case "Alls":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_All[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_ComplexType[identifier]
								instanceWhoseFieldIsAppended.Alls = append(instanceWhoseFieldIsAppended.Alls, instanceToAppend)
							}
						case "Choices":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Choice[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_ComplexType[identifier]
								instanceWhoseFieldIsAppended.Choices = append(instanceWhoseFieldIsAppended.Choices, instanceToAppend)
							}
						case "Groups":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Group[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_ComplexType[identifier]
								instanceWhoseFieldIsAppended.Groups = append(instanceWhoseFieldIsAppended.Groups, instanceToAppend)
							}
						case "Elements":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Element[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_ComplexType[identifier]
								instanceWhoseFieldIsAppended.Elements = append(instanceWhoseFieldIsAppended.Elements, instanceToAppend)
							}
						case "Attributes":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Attribute[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_ComplexType[identifier]
								instanceWhoseFieldIsAppended.Attributes = append(instanceWhoseFieldIsAppended.Attributes, instanceToAppend)
							}
						case "AttributeGroups":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_AttributeGroup[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_ComplexType[identifier]
								instanceWhoseFieldIsAppended.AttributeGroups = append(instanceWhoseFieldIsAppended.AttributeGroups, instanceToAppend)
							}
						}
					case "Documentation":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Element":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Groups":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Group[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_Element[identifier]
								instanceWhoseFieldIsAppended.Groups = append(instanceWhoseFieldIsAppended.Groups, instanceToAppend)
							}
						}
					case "Enumeration":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Extension":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Sequences":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Sequence[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_Extension[identifier]
								instanceWhoseFieldIsAppended.Sequences = append(instanceWhoseFieldIsAppended.Sequences, instanceToAppend)
							}
						case "Alls":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_All[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_Extension[identifier]
								instanceWhoseFieldIsAppended.Alls = append(instanceWhoseFieldIsAppended.Alls, instanceToAppend)
							}
						case "Choices":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Choice[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_Extension[identifier]
								instanceWhoseFieldIsAppended.Choices = append(instanceWhoseFieldIsAppended.Choices, instanceToAppend)
							}
						case "Groups":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Group[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_Extension[identifier]
								instanceWhoseFieldIsAppended.Groups = append(instanceWhoseFieldIsAppended.Groups, instanceToAppend)
							}
						case "Elements":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Element[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_Extension[identifier]
								instanceWhoseFieldIsAppended.Elements = append(instanceWhoseFieldIsAppended.Elements, instanceToAppend)
							}
						case "Attributes":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Attribute[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_Extension[identifier]
								instanceWhoseFieldIsAppended.Attributes = append(instanceWhoseFieldIsAppended.Attributes, instanceToAppend)
							}
						case "AttributeGroups":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_AttributeGroup[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_Extension[identifier]
								instanceWhoseFieldIsAppended.AttributeGroups = append(instanceWhoseFieldIsAppended.AttributeGroups, instanceToAppend)
							}
						}
					case "Group":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Sequences":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Sequence[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_Group[identifier]
								instanceWhoseFieldIsAppended.Sequences = append(instanceWhoseFieldIsAppended.Sequences, instanceToAppend)
							}
						case "Alls":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_All[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_Group[identifier]
								instanceWhoseFieldIsAppended.Alls = append(instanceWhoseFieldIsAppended.Alls, instanceToAppend)
							}
						case "Choices":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Choice[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_Group[identifier]
								instanceWhoseFieldIsAppended.Choices = append(instanceWhoseFieldIsAppended.Choices, instanceToAppend)
							}
						case "Groups":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Group[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_Group[identifier]
								instanceWhoseFieldIsAppended.Groups = append(instanceWhoseFieldIsAppended.Groups, instanceToAppend)
							}
						case "Elements":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Element[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_Group[identifier]
								instanceWhoseFieldIsAppended.Elements = append(instanceWhoseFieldIsAppended.Elements, instanceToAppend)
							}
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
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Enumeration[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_Restriction[identifier]
								instanceWhoseFieldIsAppended.Enumerations = append(instanceWhoseFieldIsAppended.Enumerations, instanceToAppend)
							}
						}
					case "Schema":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Elements":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Element[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_Schema[identifier]
								instanceWhoseFieldIsAppended.Elements = append(instanceWhoseFieldIsAppended.Elements, instanceToAppend)
							}
						case "SimpleTypes":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_SimpleType[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_Schema[identifier]
								instanceWhoseFieldIsAppended.SimpleTypes = append(instanceWhoseFieldIsAppended.SimpleTypes, instanceToAppend)
							}
						case "ComplexTypes":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_ComplexType[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_Schema[identifier]
								instanceWhoseFieldIsAppended.ComplexTypes = append(instanceWhoseFieldIsAppended.ComplexTypes, instanceToAppend)
							}
						case "AttributeGroups":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_AttributeGroup[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_Schema[identifier]
								instanceWhoseFieldIsAppended.AttributeGroups = append(instanceWhoseFieldIsAppended.AttributeGroups, instanceToAppend)
							}
						case "Groups":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Group[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_Schema[identifier]
								instanceWhoseFieldIsAppended.Groups = append(instanceWhoseFieldIsAppended.Groups, instanceToAppend)
							}
						}
					case "Sequence":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Sequences":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Sequence[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_Sequence[identifier]
								instanceWhoseFieldIsAppended.Sequences = append(instanceWhoseFieldIsAppended.Sequences, instanceToAppend)
							}
						case "Alls":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_All[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_Sequence[identifier]
								instanceWhoseFieldIsAppended.Alls = append(instanceWhoseFieldIsAppended.Alls, instanceToAppend)
							}
						case "Choices":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Choice[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_Sequence[identifier]
								instanceWhoseFieldIsAppended.Choices = append(instanceWhoseFieldIsAppended.Choices, instanceToAppend)
							}
						case "Groups":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Group[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_Sequence[identifier]
								instanceWhoseFieldIsAppended.Groups = append(instanceWhoseFieldIsAppended.Groups, instanceToAppend)
							}
						case "Elements":
							// perform the append only when the loop is processing the second argument
							if argNb == 0 {
								break
							}
							identifierOfInstanceToAppend := ident.Name
							if instanceToAppend, ok := __gong__map_Element[identifierOfInstanceToAppend]; ok {
								instanceWhoseFieldIsAppended := __gong__map_Sequence[identifier]
								instanceWhoseFieldIsAppended.Elements = append(instanceWhoseFieldIsAppended.Elements, instanceToAppend)
							}
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
				}
			}
		case *ast.BasicLit, *ast.UnaryExpr, *ast.CompositeLit:

			var basicLit *ast.BasicLit
			var exprSign = 1.0
			_ = exprSign // in case this is not used
			switch v := expr.(type) {
			case *ast.BasicLit:
				// expression is for instance ... = 18.000
				basicLit = v
			case *ast.UnaryExpr:
				// expression is for instance ... = -18.000
				// we want to extract a *ast.BasicLit from the *ast.UnaryExpr
				if bl, ok := v.X.(*ast.BasicLit); ok {
					basicLit = bl
					// Check the operator to set the sign
					if v.Op == token.SUB { // token.SUB is for '-'
						exprSign = -1
					} else if v.Op == token.ADD { // token.ADD is for '+'
						exprSign = 1
					}
				}
			case *ast.CompositeLit:
				var sl *ast.SelectorExpr
				var ident *ast.Ident
				var ok bool

				if sl, ok = v.Type.(*ast.SelectorExpr); !ok {
					break // Exits the switch case
				}

				if ident, ok = sl.X.(*ast.Ident); !ok {
					break // Exits the switch case
				}

				basicLit = new(ast.BasicLit)
				// For a "fake" literal, Kind might be set to something like token.STRING or a custom indicator
				basicLit.Kind = token.STRING // Or another appropriate token.Kind
				basicLit.Value = ident.Name + "." + sl.Sel.Name + "{}"
			}

			// astCoordinate := astCoordinate + "\tBasicLit" + "." + basicLit.Value
			// log.Println(astCoordinate)
			var ok bool
			gongstructName, ok = __gong__map_Indentifiers_gongstructName[identifier]
			if !ok {
				log.Println("gongstructName not found for identifier", identifier)
				break
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
				case "MinOccurs":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_All[identifier].MinOccurs = fielValue
				case "MaxOccurs":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_All[identifier].MaxOccurs = fielValue
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
				case "Order":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_AttributeGroup[identifier].Order = int(exprSign) * int(fielValue)
				case "Depth":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_AttributeGroup[identifier].Depth = int(exprSign) * int(fielValue)
				}
			case "Choice":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Choice[identifier].Name = fielValue
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
				case "MinOccurs":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Choice[identifier].MinOccurs = fielValue
				case "MaxOccurs":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Choice[identifier].MaxOccurs = fielValue
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
				case "Order":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_ComplexType[identifier].Order = int(exprSign) * int(fielValue)
				case "Depth":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_ComplexType[identifier].Depth = int(exprSign) * int(fielValue)
				case "MinOccurs":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ComplexType[identifier].MinOccurs = fielValue
				case "MaxOccurs":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_ComplexType[identifier].MaxOccurs = fielValue
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
				case "Order":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Extension[identifier].Order = int(exprSign) * int(fielValue)
				case "Depth":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Extension[identifier].Depth = int(exprSign) * int(fielValue)
				case "MinOccurs":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Extension[identifier].MinOccurs = fielValue
				case "MaxOccurs":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Extension[identifier].MaxOccurs = fielValue
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
				case "MinOccurs":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Group[identifier].MinOccurs = fielValue
				case "MaxOccurs":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Group[identifier].MaxOccurs = fielValue
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
				case "Order":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Schema[identifier].Order = int(exprSign) * int(fielValue)
				case "Depth":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Schema[identifier].Depth = int(exprSign) * int(fielValue)
				}
			case "Sequence":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Sequence[identifier].Name = fielValue
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
				case "MinOccurs":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Sequence[identifier].MinOccurs = fielValue
				case "MaxOccurs":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Sequence[identifier].MaxOccurs = fielValue
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
				case "Order":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_SimpleType[identifier].Order = int(exprSign) * int(fielValue)
				case "Depth":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_SimpleType[identifier].Depth = int(exprSign) * int(fielValue)
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
				log.Println("gongstructName not found for identifier", identifier)
				break
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
				case "IsDuplicatedInXSD":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_ComplexType[identifier].IsDuplicatedInXSD = fielValue
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
			var basicLit *ast.BasicLit
			var ident *ast.Ident

			// assignment to enum field
			selectorExpr := expr
			// astCoordinate := astCoordinate + "\tSelectorExpr"
			switch X := selectorExpr.X.(type) {
			case *ast.Ident:
				ident := X
				_ = ident
				// astCoordinate := astCoordinate + "\tX" + "." + ident.Name
				// log.Println(astCoordinate)
			case *ast.CompositeLit:
				var ok bool
				var sl *ast.SelectorExpr

				if sl, ok = X.Type.(*ast.SelectorExpr); !ok {
					break // Exits the switch case
				}

				if ident, ok = sl.X.(*ast.Ident); !ok {
					break // Exits the switch case
				}

				basicLit = new(ast.BasicLit)
				// For a "fake" literal, Kind might be set to something like token.STRING or a custom indicator
				basicLit.Kind = token.STRING // Or another appropriate token.Kind
				basicLit.Value = ident.Name + "." + sl.Sel.Name + "{}." + selectorExpr.Sel.Name
			}

			if Sel := selectorExpr.Sel; Sel != nil {
				// astCoordinate := astCoordinate + "\tSel" + "." + Sel.Name
				// log.Println(astCoordinate)

				// enum field
				var ok bool
				gongstructName, ok = __gong__map_Indentifiers_gongstructName[identifier]
				if !ok {
					log.Println("gongstructName not found for identifier", identifier)
					break
				}

				if basicLit == nil {
					// for the meta field written as ref_models.ENUM_VALUE1
					basicLit = new(ast.BasicLit)
					basicLit.Kind = token.STRING // Or another appropriate token.Kind
					basicLit.Value = selectorExpr.X.(*ast.Ident).Name + "." + Sel.Name
					_ = basicLit.Kind
					_ = basicLit.Value
				}

				// remove first and last char
				enumValue := Sel.Name
				_ = enumValue
				switch gongstructName {
				// insertion point for selector expr assignments
				case "All":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "Annotation":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "Attribute":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "AttributeGroup":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "Choice":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "ComplexContent":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "ComplexType":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "Documentation":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "Element":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "Enumeration":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "Extension":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "Group":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "Length":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "MaxInclusive":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "MaxLength":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "MinInclusive":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "MinLength":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "Pattern":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "Restriction":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "Schema":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "Sequence":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "SimpleContent":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "SimpleType":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "TotalDigit":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "Union":
					switch fieldName {
					// insertion point for selector expr assign code
					}
				case "WhiteSpace":
					switch fieldName {
					// insertion point for selector expr assign code
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
