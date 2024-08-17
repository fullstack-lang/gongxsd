// generated code - do not edit
package models

import (
	"cmp"
	"errors"
	"fmt"
	"math"
	"slices"
	"time"

	"golang.org/x/exp/maps"
)

func __Gong__Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// errUnkownEnum is returns when a value cannot match enum values
var errUnkownEnum = errors.New("unkown enum")

// needed to avoid when fmt package is not needed by generated code
var __dummy__fmt_variable fmt.Scanner

// idem for math package when not need by generated code
var __dummy_math_variable = math.E

// swagger:ignore
type __void any

// needed for creating set of instances in the stage
var __member __void

// GongStructInterface is the interface met by GongStructs
// It allows runtime reflexion of instances (without the hassle of the "reflect" package)
type GongStructInterface interface {
	GetName() (res string)
	// GetID() (res int)
	// GetFields() (res []string)
	// GetFieldStringValue(fieldName string) (res string)
}

// StageStruct enables storage of staged instances
// swagger:ignore
type StageStruct struct {
	path string

	// insertion point for definition of arrays registering instances
	Annotations           map[*Annotation]any
	Annotations_mapString map[string]*Annotation

	// insertion point for slice of pointers maps
	Annotation_Documentations_reverseMap map[*Documentation]*Annotation

	OnAfterAnnotationCreateCallback OnAfterCreateInterface[Annotation]
	OnAfterAnnotationUpdateCallback OnAfterUpdateInterface[Annotation]
	OnAfterAnnotationDeleteCallback OnAfterDeleteInterface[Annotation]
	OnAfterAnnotationReadCallback   OnAfterReadInterface[Annotation]

	ComplexContents           map[*ComplexContent]any
	ComplexContents_mapString map[string]*ComplexContent

	// insertion point for slice of pointers maps
	OnAfterComplexContentCreateCallback OnAfterCreateInterface[ComplexContent]
	OnAfterComplexContentUpdateCallback OnAfterUpdateInterface[ComplexContent]
	OnAfterComplexContentDeleteCallback OnAfterDeleteInterface[ComplexContent]
	OnAfterComplexContentReadCallback   OnAfterReadInterface[ComplexContent]

	ComplexTypes           map[*ComplexType]any
	ComplexTypes_mapString map[string]*ComplexType

	// insertion point for slice of pointers maps
	OnAfterComplexTypeCreateCallback OnAfterCreateInterface[ComplexType]
	OnAfterComplexTypeUpdateCallback OnAfterUpdateInterface[ComplexType]
	OnAfterComplexTypeDeleteCallback OnAfterDeleteInterface[ComplexType]
	OnAfterComplexTypeReadCallback   OnAfterReadInterface[ComplexType]

	Documentations           map[*Documentation]any
	Documentations_mapString map[string]*Documentation

	// insertion point for slice of pointers maps
	OnAfterDocumentationCreateCallback OnAfterCreateInterface[Documentation]
	OnAfterDocumentationUpdateCallback OnAfterUpdateInterface[Documentation]
	OnAfterDocumentationDeleteCallback OnAfterDeleteInterface[Documentation]
	OnAfterDocumentationReadCallback   OnAfterReadInterface[Documentation]

	Schemas           map[*Schema]any
	Schemas_mapString map[string]*Schema

	// insertion point for slice of pointers maps
	OnAfterSchemaCreateCallback OnAfterCreateInterface[Schema]
	OnAfterSchemaUpdateCallback OnAfterUpdateInterface[Schema]
	OnAfterSchemaDeleteCallback OnAfterDeleteInterface[Schema]
	OnAfterSchemaReadCallback   OnAfterReadInterface[Schema]

	AllModelsStructCreateCallback AllModelsStructCreateInterface

	AllModelsStructDeleteCallback AllModelsStructDeleteInterface

	BackRepo BackRepoInterface

	// if set will be called before each commit to the back repo
	OnInitCommitCallback          OnInitCommitInterface
	OnInitCommitFromFrontCallback OnInitCommitInterface
	OnInitCommitFromBackCallback  OnInitCommitInterface

	// store the number of instance per gongstruct
	Map_GongStructName_InstancesNb map[string]int

	// store meta package import
	MetaPackageImportPath  string
	MetaPackageImportAlias string

	// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
	// map to enable docLink renaming when an identifier is renamed
	Map_DocLink_Renaming map[string]GONG__Identifier
	// the to be removed stops here
}

func (stage *StageStruct) GetType() string {
	return "github.com/fullstack-lang/gongxsd/alt/go/models"
}

type GONG__Identifier struct {
	Ident string
	Type  GONG__ExpressionType
}

type OnInitCommitInterface interface {
	BeforeCommit(stage *StageStruct)
}

// OnAfterCreateInterface callback when an instance is updated from the front
type OnAfterCreateInterface[Type Gongstruct] interface {
	OnAfterCreate(stage *StageStruct,
		instance *Type)
}

// OnAfterReadInterface callback when an instance is updated from the front
type OnAfterReadInterface[Type Gongstruct] interface {
	OnAfterRead(stage *StageStruct,
		instance *Type)
}

// OnAfterUpdateInterface callback when an instance is updated from the front
type OnAfterUpdateInterface[Type Gongstruct] interface {
	OnAfterUpdate(stage *StageStruct, old, new *Type)
}

// OnAfterDeleteInterface callback when an instance is updated from the front
type OnAfterDeleteInterface[Type Gongstruct] interface {
	OnAfterDelete(stage *StageStruct,
		staged, front *Type)
}

type BackRepoInterface interface {
	Commit(stage *StageStruct)
	Checkout(stage *StageStruct)
	Backup(stage *StageStruct, dirPath string)
	Restore(stage *StageStruct, dirPath string)
	BackupXL(stage *StageStruct, dirPath string)
	RestoreXL(stage *StageStruct, dirPath string)
	// insertion point for Commit and Checkout signatures
	CommitAnnotation(annotation *Annotation)
	CheckoutAnnotation(annotation *Annotation)
	CommitComplexContent(complexcontent *ComplexContent)
	CheckoutComplexContent(complexcontent *ComplexContent)
	CommitComplexType(complextype *ComplexType)
	CheckoutComplexType(complextype *ComplexType)
	CommitDocumentation(documentation *Documentation)
	CheckoutDocumentation(documentation *Documentation)
	CommitSchema(schema *Schema)
	CheckoutSchema(schema *Schema)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

func NewStage(path string) (stage *StageStruct) {

	stage = &StageStruct{ // insertion point for array initiatialisation
		Annotations:           make(map[*Annotation]any),
		Annotations_mapString: make(map[string]*Annotation),

		ComplexContents:           make(map[*ComplexContent]any),
		ComplexContents_mapString: make(map[string]*ComplexContent),

		ComplexTypes:           make(map[*ComplexType]any),
		ComplexTypes_mapString: make(map[string]*ComplexType),

		Documentations:           make(map[*Documentation]any),
		Documentations_mapString: make(map[string]*Documentation),

		Schemas:           make(map[*Schema]any),
		Schemas_mapString: make(map[string]*Schema),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		path: path,

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here
	}

	return
}

func (stage *StageStruct) GetPath() string {
	return stage.path
}

func (stage *StageStruct) CommitWithSuspendedCallbacks() {

	tmp := stage.OnInitCommitFromBackCallback
	stage.OnInitCommitFromBackCallback = nil
	stage.Commit()
	stage.OnInitCommitFromBackCallback = tmp
}

func (stage *StageStruct) Commit() {
	stage.ComputeReverseMaps()

	if stage.BackRepo != nil {
		stage.BackRepo.Commit(stage)
	}

	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["Annotation"] = len(stage.Annotations)
	stage.Map_GongStructName_InstancesNb["ComplexContent"] = len(stage.ComplexContents)
	stage.Map_GongStructName_InstancesNb["ComplexType"] = len(stage.ComplexTypes)
	stage.Map_GongStructName_InstancesNb["Documentation"] = len(stage.Documentations)
	stage.Map_GongStructName_InstancesNb["Schema"] = len(stage.Schemas)

}

func (stage *StageStruct) Checkout() {
	if stage.BackRepo != nil {
		stage.BackRepo.Checkout(stage)
	}

	stage.ComputeReverseMaps()
	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["Annotation"] = len(stage.Annotations)
	stage.Map_GongStructName_InstancesNb["ComplexContent"] = len(stage.ComplexContents)
	stage.Map_GongStructName_InstancesNb["ComplexType"] = len(stage.ComplexTypes)
	stage.Map_GongStructName_InstancesNb["Documentation"] = len(stage.Documentations)
	stage.Map_GongStructName_InstancesNb["Schema"] = len(stage.Schemas)

}

// backup generates backup files in the dirPath
func (stage *StageStruct) Backup(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.Backup(stage, dirPath)
	}
}

// Restore resets Stage & BackRepo and restores their content from the restore files in dirPath
func (stage *StageStruct) Restore(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.Restore(stage, dirPath)
	}
}

// backup generates backup files in the dirPath
func (stage *StageStruct) BackupXL(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.BackupXL(stage, dirPath)
	}
}

// Restore resets Stage & BackRepo and restores their content from the restore files in dirPath
func (stage *StageStruct) RestoreXL(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.RestoreXL(stage, dirPath)
	}
}

// insertion point for cumulative sub template with model space calls
// Stage puts annotation to the model stage
func (annotation *Annotation) Stage(stage *StageStruct) *Annotation {
	stage.Annotations[annotation] = __member
	stage.Annotations_mapString[annotation.Name] = annotation

	return annotation
}

// Unstage removes annotation off the model stage
func (annotation *Annotation) Unstage(stage *StageStruct) *Annotation {
	delete(stage.Annotations, annotation)
	delete(stage.Annotations_mapString, annotation.Name)
	return annotation
}

// UnstageVoid removes annotation off the model stage
func (annotation *Annotation) UnstageVoid(stage *StageStruct) {
	delete(stage.Annotations, annotation)
	delete(stage.Annotations_mapString, annotation.Name)
}

// commit annotation to the back repo (if it is already staged)
func (annotation *Annotation) Commit(stage *StageStruct) *Annotation {
	if _, ok := stage.Annotations[annotation]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitAnnotation(annotation)
		}
	}
	return annotation
}

func (annotation *Annotation) CommitVoid(stage *StageStruct) {
	annotation.Commit(stage)
}

// Checkout annotation to the back repo (if it is already staged)
func (annotation *Annotation) Checkout(stage *StageStruct) *Annotation {
	if _, ok := stage.Annotations[annotation]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutAnnotation(annotation)
		}
	}
	return annotation
}

// for satisfaction of GongStruct interface
func (annotation *Annotation) GetName() (res string) {
	return annotation.Name
}

// Stage puts complexcontent to the model stage
func (complexcontent *ComplexContent) Stage(stage *StageStruct) *ComplexContent {
	stage.ComplexContents[complexcontent] = __member
	stage.ComplexContents_mapString[complexcontent.Name] = complexcontent

	return complexcontent
}

// Unstage removes complexcontent off the model stage
func (complexcontent *ComplexContent) Unstage(stage *StageStruct) *ComplexContent {
	delete(stage.ComplexContents, complexcontent)
	delete(stage.ComplexContents_mapString, complexcontent.Name)
	return complexcontent
}

// UnstageVoid removes complexcontent off the model stage
func (complexcontent *ComplexContent) UnstageVoid(stage *StageStruct) {
	delete(stage.ComplexContents, complexcontent)
	delete(stage.ComplexContents_mapString, complexcontent.Name)
}

// commit complexcontent to the back repo (if it is already staged)
func (complexcontent *ComplexContent) Commit(stage *StageStruct) *ComplexContent {
	if _, ok := stage.ComplexContents[complexcontent]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitComplexContent(complexcontent)
		}
	}
	return complexcontent
}

func (complexcontent *ComplexContent) CommitVoid(stage *StageStruct) {
	complexcontent.Commit(stage)
}

// Checkout complexcontent to the back repo (if it is already staged)
func (complexcontent *ComplexContent) Checkout(stage *StageStruct) *ComplexContent {
	if _, ok := stage.ComplexContents[complexcontent]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutComplexContent(complexcontent)
		}
	}
	return complexcontent
}

// for satisfaction of GongStruct interface
func (complexcontent *ComplexContent) GetName() (res string) {
	return complexcontent.Name
}

// Stage puts complextype to the model stage
func (complextype *ComplexType) Stage(stage *StageStruct) *ComplexType {
	stage.ComplexTypes[complextype] = __member
	stage.ComplexTypes_mapString[complextype.Name] = complextype

	return complextype
}

// Unstage removes complextype off the model stage
func (complextype *ComplexType) Unstage(stage *StageStruct) *ComplexType {
	delete(stage.ComplexTypes, complextype)
	delete(stage.ComplexTypes_mapString, complextype.Name)
	return complextype
}

// UnstageVoid removes complextype off the model stage
func (complextype *ComplexType) UnstageVoid(stage *StageStruct) {
	delete(stage.ComplexTypes, complextype)
	delete(stage.ComplexTypes_mapString, complextype.Name)
}

// commit complextype to the back repo (if it is already staged)
func (complextype *ComplexType) Commit(stage *StageStruct) *ComplexType {
	if _, ok := stage.ComplexTypes[complextype]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitComplexType(complextype)
		}
	}
	return complextype
}

func (complextype *ComplexType) CommitVoid(stage *StageStruct) {
	complextype.Commit(stage)
}

// Checkout complextype to the back repo (if it is already staged)
func (complextype *ComplexType) Checkout(stage *StageStruct) *ComplexType {
	if _, ok := stage.ComplexTypes[complextype]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutComplexType(complextype)
		}
	}
	return complextype
}

// for satisfaction of GongStruct interface
func (complextype *ComplexType) GetName() (res string) {
	return complextype.Name
}

// Stage puts documentation to the model stage
func (documentation *Documentation) Stage(stage *StageStruct) *Documentation {
	stage.Documentations[documentation] = __member
	stage.Documentations_mapString[documentation.Name] = documentation

	return documentation
}

// Unstage removes documentation off the model stage
func (documentation *Documentation) Unstage(stage *StageStruct) *Documentation {
	delete(stage.Documentations, documentation)
	delete(stage.Documentations_mapString, documentation.Name)
	return documentation
}

// UnstageVoid removes documentation off the model stage
func (documentation *Documentation) UnstageVoid(stage *StageStruct) {
	delete(stage.Documentations, documentation)
	delete(stage.Documentations_mapString, documentation.Name)
}

// commit documentation to the back repo (if it is already staged)
func (documentation *Documentation) Commit(stage *StageStruct) *Documentation {
	if _, ok := stage.Documentations[documentation]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitDocumentation(documentation)
		}
	}
	return documentation
}

func (documentation *Documentation) CommitVoid(stage *StageStruct) {
	documentation.Commit(stage)
}

// Checkout documentation to the back repo (if it is already staged)
func (documentation *Documentation) Checkout(stage *StageStruct) *Documentation {
	if _, ok := stage.Documentations[documentation]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutDocumentation(documentation)
		}
	}
	return documentation
}

// for satisfaction of GongStruct interface
func (documentation *Documentation) GetName() (res string) {
	return documentation.Name
}

// Stage puts schema to the model stage
func (schema *Schema) Stage(stage *StageStruct) *Schema {
	stage.Schemas[schema] = __member
	stage.Schemas_mapString[schema.Name] = schema

	return schema
}

// Unstage removes schema off the model stage
func (schema *Schema) Unstage(stage *StageStruct) *Schema {
	delete(stage.Schemas, schema)
	delete(stage.Schemas_mapString, schema.Name)
	return schema
}

// UnstageVoid removes schema off the model stage
func (schema *Schema) UnstageVoid(stage *StageStruct) {
	delete(stage.Schemas, schema)
	delete(stage.Schemas_mapString, schema.Name)
}

// commit schema to the back repo (if it is already staged)
func (schema *Schema) Commit(stage *StageStruct) *Schema {
	if _, ok := stage.Schemas[schema]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSchema(schema)
		}
	}
	return schema
}

func (schema *Schema) CommitVoid(stage *StageStruct) {
	schema.Commit(stage)
}

// Checkout schema to the back repo (if it is already staged)
func (schema *Schema) Checkout(stage *StageStruct) *Schema {
	if _, ok := stage.Schemas[schema]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutSchema(schema)
		}
	}
	return schema
}

// for satisfaction of GongStruct interface
func (schema *Schema) GetName() (res string) {
	return schema.Name
}

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMAnnotation(Annotation *Annotation)
	CreateORMComplexContent(ComplexContent *ComplexContent)
	CreateORMComplexType(ComplexType *ComplexType)
	CreateORMDocumentation(Documentation *Documentation)
	CreateORMSchema(Schema *Schema)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMAnnotation(Annotation *Annotation)
	DeleteORMComplexContent(ComplexContent *ComplexContent)
	DeleteORMComplexType(ComplexType *ComplexType)
	DeleteORMDocumentation(Documentation *Documentation)
	DeleteORMSchema(Schema *Schema)
}

func (stage *StageStruct) Reset() { // insertion point for array reset
	stage.Annotations = make(map[*Annotation]any)
	stage.Annotations_mapString = make(map[string]*Annotation)

	stage.ComplexContents = make(map[*ComplexContent]any)
	stage.ComplexContents_mapString = make(map[string]*ComplexContent)

	stage.ComplexTypes = make(map[*ComplexType]any)
	stage.ComplexTypes_mapString = make(map[string]*ComplexType)

	stage.Documentations = make(map[*Documentation]any)
	stage.Documentations_mapString = make(map[string]*Documentation)

	stage.Schemas = make(map[*Schema]any)
	stage.Schemas_mapString = make(map[string]*Schema)

}

func (stage *StageStruct) Nil() { // insertion point for array nil
	stage.Annotations = nil
	stage.Annotations_mapString = nil

	stage.ComplexContents = nil
	stage.ComplexContents_mapString = nil

	stage.ComplexTypes = nil
	stage.ComplexTypes_mapString = nil

	stage.Documentations = nil
	stage.Documentations_mapString = nil

	stage.Schemas = nil
	stage.Schemas_mapString = nil

}

func (stage *StageStruct) Unstage() { // insertion point for array nil
	for annotation := range stage.Annotations {
		annotation.Unstage(stage)
	}

	for complexcontent := range stage.ComplexContents {
		complexcontent.Unstage(stage)
	}

	for complextype := range stage.ComplexTypes {
		complextype.Unstage(stage)
	}

	for documentation := range stage.Documentations {
		documentation.Unstage(stage)
	}

	for schema := range stage.Schemas {
		schema.Unstage(stage)
	}

}

// Gongstruct is the type parameter for generated generic function that allows
// - access to staged instances
// - navigation between staged instances by going backward association links between gongstruct
// - full refactoring of Gongstruct identifiers / fields
type Gongstruct interface {

}

type GongtructBasicField interface {
	int | float64 | bool | string | time.Time | time.Duration
}

// Gongstruct is the type parameter for generated generic function that allows
// - access to staged instances
// - navigation between staged instances by going backward association links between gongstruct
// - full refactoring of Gongstruct identifiers / fields
type PointerToGongstruct interface {
	GetName() string
	CommitVoid(*StageStruct)
	UnstageVoid(stage *StageStruct)
	comparable
}

func CompareGongstructByName[T PointerToGongstruct](a, b T) int {
	return cmp.Compare(a.GetName(), b.GetName())
}

func SortGongstructSetByName[T PointerToGongstruct](set map[T]any) (sortedSlice []T) {

	sortedSlice = maps.Keys(set)
	slices.SortFunc(sortedSlice, CompareGongstructByName)

	return
}

func GetGongstrucsSorted[T PointerToGongstruct](stage *StageStruct) (sortedSlice []T) {

	set := GetGongstructInstancesSetFromPointerType[T](stage)
	sortedSlice = SortGongstructSetByName(*set)

	return
}

type GongstructSet interface {
	map[any]any
}

type GongstructMapString interface {
	map[any]any
}

// GongGetSet returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GongGetSet[Type GongstructSet](stage *StageStruct) *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case map[*Annotation]any:
		return any(&stage.Annotations).(*Type)
	case map[*ComplexContent]any:
		return any(&stage.ComplexContents).(*Type)
	case map[*ComplexType]any:
		return any(&stage.ComplexTypes).(*Type)
	case map[*Documentation]any:
		return any(&stage.Documentations).(*Type)
	case map[*Schema]any:
		return any(&stage.Schemas).(*Type)
	default:
		return nil
	}
}

// GongGetMap returns the map of staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GongGetMap[Type GongstructMapString](stage *StageStruct) *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case map[string]*Annotation:
		return any(&stage.Annotations_mapString).(*Type)
	case map[string]*ComplexContent:
		return any(&stage.ComplexContents_mapString).(*Type)
	case map[string]*ComplexType:
		return any(&stage.ComplexTypes_mapString).(*Type)
	case map[string]*Documentation:
		return any(&stage.Documentations_mapString).(*Type)
	case map[string]*Schema:
		return any(&stage.Schemas_mapString).(*Type)
	default:
		return nil
	}
}

// GetGongstructInstancesSet returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesSet[Type Gongstruct](stage *StageStruct) *map[*Type]any {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case Annotation:
		return any(&stage.Annotations).(*map[*Type]any)
	case ComplexContent:
		return any(&stage.ComplexContents).(*map[*Type]any)
	case ComplexType:
		return any(&stage.ComplexTypes).(*map[*Type]any)
	case Documentation:
		return any(&stage.Documentations).(*map[*Type]any)
	case Schema:
		return any(&stage.Schemas).(*map[*Type]any)
	default:
		return nil
	}
}

// GetGongstructInstancesSetFromPointerType returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesSetFromPointerType[Type PointerToGongstruct](stage *StageStruct) *map[Type]any {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case *Annotation:
		return any(&stage.Annotations).(*map[Type]any)
	case *ComplexContent:
		return any(&stage.ComplexContents).(*map[Type]any)
	case *ComplexType:
		return any(&stage.ComplexTypes).(*map[Type]any)
	case *Documentation:
		return any(&stage.Documentations).(*map[Type]any)
	case *Schema:
		return any(&stage.Schemas).(*map[Type]any)
	default:
		return nil
	}
}

// GetGongstructInstancesMap returns the map of staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GetGongstructInstancesMap[Type Gongstruct](stage *StageStruct) *map[string]*Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case Annotation:
		return any(&stage.Annotations_mapString).(*map[string]*Type)
	case ComplexContent:
		return any(&stage.ComplexContents_mapString).(*map[string]*Type)
	case ComplexType:
		return any(&stage.ComplexTypes_mapString).(*map[string]*Type)
	case Documentation:
		return any(&stage.Documentations_mapString).(*map[string]*Type)
	case Schema:
		return any(&stage.Schemas_mapString).(*map[string]*Type)
	default:
		return nil
	}
}

// GetAssociationName is a generic function that returns an instance of Type
// where each association is filled with an instance whose name is the name of the association
//
// This function can be handy for generating navigation function that are refactorable
func GetAssociationName[Type Gongstruct]() *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for instance with special fields
	case Annotation:
		return any(&Annotation{
			// Initialisation of associations
			// field is initialized with an instance of Documentation with the name of the field
			Documentations: []*Documentation{{Name: "Documentations"}},
		}).(*Type)
	case ComplexContent:
		return any(&ComplexContent{
			// Initialisation of associations
		}).(*Type)
	case ComplexType:
		return any(&ComplexType{
			// Initialisation of associations
		}).(*Type)
	case Documentation:
		return any(&Documentation{
			// Initialisation of associations
		}).(*Type)
	case Schema:
		return any(&Schema{
			// Initialisation of associations
			// field is initialized with Annotated problem with composites
			
		}).(*Type)
	default:
		return nil
	}
}

// GetPointerReverseMap allows backtrack navigation of any Start.Fieldname
// associations (0..1) that is a pointer from one staged Gongstruct (type Start)
// instances to another (type End)
//
// The function provides a map with keys as instances of End and values to arrays of *Start
// the map is construed by iterating over all Start instances and populationg keys with End instances
// and values with slice of Start instances
func GetPointerReverseMap[Start, End Gongstruct](fieldname string, stage *StageStruct) map[*End][]*Start {

	var ret Start

	switch any(ret).(type) {
	// insertion point of functions that provide maps for reverse associations
	// reverse maps of direct associations of Annotation
	case Annotation:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ComplexContent
	case ComplexContent:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ComplexType
	case ComplexType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Documentation
	case Documentation:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Schema
	case Schema:
		switch fieldname {
		// insertion point for per direct association field
		case "Annotation":
			res := make(map[*Annotation][]*Schema)
			for schema := range stage.Schemas {
				if schema.Annotation != nil {
					annotation_ := schema.Annotation
					var schemas []*Schema
					_, ok := res[annotation_]
					if ok {
						schemas = res[annotation_]
					} else {
						schemas = make([]*Schema, 0)
					}
					schemas = append(schemas, schema)
					res[annotation_] = schemas
				}
			}
			return any(res).(map[*End][]*Start)
		case "Sequence2.ComplexType":
			res := make(map[*ComplexType][]*Schema)
			for schema := range stage.Schemas {
				if schema.Sequence2.ComplexType != nil {
					complextype_ := schema.Sequence2.ComplexType
					var schemas []*Schema
					_, ok := res[complextype_]
					if ok {
						schemas = res[complextype_]
					} else {
						schemas = make([]*Schema, 0)
					}
					schemas = append(schemas, schema)
					res[complextype_] = schemas
				}
			}
			return any(res).(map[*End][]*Start)
		}
	}
	return nil
}

// GetSliceOfPointersReverseMap allows backtrack navigation of any Start.Fieldname
// associations (0..N) between one staged Gongstruct instances and many others
//
// The function provides a map with keys as instances of End and values to *Start instances
// the map is construed by iterating over all Start instances and populating keys with End instances
// and values with the Start instances
func GetSliceOfPointersReverseMap[Start, End Gongstruct](fieldname string, stage *StageStruct) map[*End]*Start {

	var ret Start

	switch any(ret).(type) {
	// insertion point of functions that provide maps for reverse associations
	// reverse maps of direct associations of Annotation
	case Annotation:
		switch fieldname {
		// insertion point for per direct association field
		case "Documentations":
			res := make(map[*Documentation]*Annotation)
			for annotation := range stage.Annotations {
				for _, documentation_ := range annotation.Documentations {
					res[documentation_] = annotation
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of ComplexContent
	case ComplexContent:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ComplexType
	case ComplexType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Documentation
	case Documentation:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Schema
	case Schema:
		switch fieldname {
		// insertion point for per direct association field
		}
	}
	return nil
}

// GetGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetGongstructName[Type Gongstruct]() (res string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case Annotation:
		res = "Annotation"
	case ComplexContent:
		res = "ComplexContent"
	case ComplexType:
		res = "ComplexType"
	case Documentation:
		res = "Documentation"
	case Schema:
		res = "Schema"
	}
	return res
}

// GetPointerToGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetPointerToGongstructName[Type PointerToGongstruct]() (res string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case *Annotation:
		res = "Annotation"
	case *ComplexContent:
		res = "ComplexContent"
	case *ComplexType:
		res = "ComplexType"
	case *Documentation:
		res = "Documentation"
	case *Schema:
		res = "Schema"
	}
	return res
}

// GetFields return the array of the fields
func GetFields[Type Gongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case Annotation:
		res = []string{"Name", "Documentations"}
	case ComplexContent:
		res = []string{"Name"}
	case ComplexType:
		res = []string{"Name"}
	case Documentation:
		res = []string{"Name", "Text", "Source", "Lang"}
	case Schema:
		res = []string{"Name", "Xs", "Annotation", "Schema_A_ComplexType_A_ComplexContentDummy", "Schema_A_ComplexType_A_ComplexContent_A_Extension_SequenceDummy", "Schema_A_ComplexType_A_ComplexContent_A_Extension_Sequence_Sequence1Dummy", "Sequence2.ComplexType"}
	}
	return
}

type ReverseField struct {
	GongstructName string
	Fieldname      string
}

func GetReverseFields[Type Gongstruct]() (res []ReverseField) {

	res = make([]ReverseField, 0)

	var ret Type

	switch any(ret).(type) {

	// insertion point for generic get gongstruct name
	case Annotation:
		var rf ReverseField
		_ = rf
	case ComplexContent:
		var rf ReverseField
		_ = rf
	case ComplexType:
		var rf ReverseField
		_ = rf
	case Documentation:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Annotation"
		rf.Fieldname = "Documentations"
		res = append(res, rf)
	case Schema:
		var rf ReverseField
		_ = rf
	}
	return
}

// GetFieldsFromPointer return the array of the fields
func GetFieldsFromPointer[Type PointerToGongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case *Annotation:
		res = []string{"Name", "Documentations"}
	case *ComplexContent:
		res = []string{"Name"}
	case *ComplexType:
		res = []string{"Name"}
	case *Documentation:
		res = []string{"Name", "Text", "Source", "Lang"}
	case *Schema:
		res = []string{"Name", "Xs", "Annotation", "Schema_A_ComplexType_A_ComplexContentDummy", "Schema_A_ComplexType_A_ComplexContent_A_Extension_SequenceDummy", "Schema_A_ComplexType_A_ComplexContent_A_Extension_Sequence_Sequence1Dummy", "Sequence2.ComplexType"}
	}
	return
}

func GetFieldStringValueFromPointer[Type PointerToGongstruct](instance Type, fieldName string) (res string) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
	case *Annotation:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Documentations":
			for idx, __instance__ := range inferedInstance.Documentations {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case *ComplexContent:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		}
	case *ComplexType:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		}
	case *Documentation:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Text":
			res = inferedInstance.Text
		case "Source":
			res = inferedInstance.Source
		case "Lang":
			res = inferedInstance.Lang
		}
	case *Schema:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Xs":
			res = inferedInstance.Xs
		case "Annotation":
			if inferedInstance.Annotation != nil {
				res = inferedInstance.Annotation.Name
			}
		case "Schema_A_ComplexType_A_ComplexContentDummy":
			res = fmt.Sprintf("%d", inferedInstance.Schema_A_ComplexType_A_ComplexContentDummy)
		case "Schema_A_ComplexType_A_ComplexContent_A_Extension_SequenceDummy":
			res = fmt.Sprintf("%d", inferedInstance.Schema_A_ComplexType_A_ComplexContent_A_Extension_SequenceDummy)
		case "Schema_A_ComplexType_A_ComplexContent_A_Extension_Sequence_Sequence1Dummy":
			res = fmt.Sprintf("%d", inferedInstance.Schema_A_ComplexType_A_ComplexContent_A_Extension_Sequence_Sequence1Dummy)
		case "Sequence2.ComplexType":
			if inferedInstance.Sequence2.ComplexType != nil {
				res = inferedInstance.Sequence2.ComplexType.Name
			}
		}
	default:
		_ = inferedInstance
	}
	return
}

func GetFieldStringValue[Type Gongstruct](instance Type, fieldName string) (res string) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
	case Annotation:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Documentations":
			for idx, __instance__ := range inferedInstance.Documentations {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case ComplexContent:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		}
	case ComplexType:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		}
	case Documentation:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Text":
			res = inferedInstance.Text
		case "Source":
			res = inferedInstance.Source
		case "Lang":
			res = inferedInstance.Lang
		}
	case Schema:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Xs":
			res = inferedInstance.Xs
		case "Annotation":
			if inferedInstance.Annotation != nil {
				res = inferedInstance.Annotation.Name
			}
		case "Schema_A_ComplexType_A_ComplexContentDummy":
			res = fmt.Sprintf("%d", inferedInstance.Schema_A_ComplexType_A_ComplexContentDummy)
		case "Schema_A_ComplexType_A_ComplexContent_A_Extension_SequenceDummy":
			res = fmt.Sprintf("%d", inferedInstance.Schema_A_ComplexType_A_ComplexContent_A_Extension_SequenceDummy)
		case "Schema_A_ComplexType_A_ComplexContent_A_Extension_Sequence_Sequence1Dummy":
			res = fmt.Sprintf("%d", inferedInstance.Schema_A_ComplexType_A_ComplexContent_A_Extension_Sequence_Sequence1Dummy)
		case "Sequence2.ComplexType":
			if inferedInstance.Sequence2.ComplexType != nil {
				res = inferedInstance.Sequence2.ComplexType.Name
			}
		}
	default:
		_ = inferedInstance
	}
	return
}

// Last line of the template
