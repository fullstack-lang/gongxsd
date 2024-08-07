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
	ComplexTypes           map[*ComplexType]any
	ComplexTypes_mapString map[string]*ComplexType

	// insertion point for slice of pointers maps
	OnAfterComplexTypeCreateCallback OnAfterCreateInterface[ComplexType]
	OnAfterComplexTypeUpdateCallback OnAfterUpdateInterface[ComplexType]
	OnAfterComplexTypeDeleteCallback OnAfterDeleteInterface[ComplexType]
	OnAfterComplexTypeReadCallback   OnAfterReadInterface[ComplexType]

	Elements           map[*Element]any
	Elements_mapString map[string]*Element

	// insertion point for slice of pointers maps
	OnAfterElementCreateCallback OnAfterCreateInterface[Element]
	OnAfterElementUpdateCallback OnAfterUpdateInterface[Element]
	OnAfterElementDeleteCallback OnAfterDeleteInterface[Element]
	OnAfterElementReadCallback   OnAfterReadInterface[Element]

	Enumerations           map[*Enumeration]any
	Enumerations_mapString map[string]*Enumeration

	// insertion point for slice of pointers maps
	OnAfterEnumerationCreateCallback OnAfterCreateInterface[Enumeration]
	OnAfterEnumerationUpdateCallback OnAfterUpdateInterface[Enumeration]
	OnAfterEnumerationDeleteCallback OnAfterDeleteInterface[Enumeration]
	OnAfterEnumerationReadCallback   OnAfterReadInterface[Enumeration]

	Restrictions           map[*Restriction]any
	Restrictions_mapString map[string]*Restriction

	// insertion point for slice of pointers maps
	Restriction_Enumerations_reverseMap map[*Enumeration]*Restriction

	OnAfterRestrictionCreateCallback OnAfterCreateInterface[Restriction]
	OnAfterRestrictionUpdateCallback OnAfterUpdateInterface[Restriction]
	OnAfterRestrictionDeleteCallback OnAfterDeleteInterface[Restriction]
	OnAfterRestrictionReadCallback   OnAfterReadInterface[Restriction]

	Schemas           map[*Schema]any
	Schemas_mapString map[string]*Schema

	// insertion point for slice of pointers maps
	Schema_Elements_reverseMap map[*Element]*Schema

	Schema_SimpleTypes_reverseMap map[*SimpleType]*Schema

	Schema_ComplexTypes_reverseMap map[*ComplexType]*Schema

	OnAfterSchemaCreateCallback OnAfterCreateInterface[Schema]
	OnAfterSchemaUpdateCallback OnAfterUpdateInterface[Schema]
	OnAfterSchemaDeleteCallback OnAfterDeleteInterface[Schema]
	OnAfterSchemaReadCallback   OnAfterReadInterface[Schema]

	Sequences           map[*Sequence]any
	Sequences_mapString map[string]*Sequence

	// insertion point for slice of pointers maps
	Sequence_Elements_reverseMap map[*Element]*Sequence

	OnAfterSequenceCreateCallback OnAfterCreateInterface[Sequence]
	OnAfterSequenceUpdateCallback OnAfterUpdateInterface[Sequence]
	OnAfterSequenceDeleteCallback OnAfterDeleteInterface[Sequence]
	OnAfterSequenceReadCallback   OnAfterReadInterface[Sequence]

	SimpleTypes           map[*SimpleType]any
	SimpleTypes_mapString map[string]*SimpleType

	// insertion point for slice of pointers maps
	OnAfterSimpleTypeCreateCallback OnAfterCreateInterface[SimpleType]
	OnAfterSimpleTypeUpdateCallback OnAfterUpdateInterface[SimpleType]
	OnAfterSimpleTypeDeleteCallback OnAfterDeleteInterface[SimpleType]
	OnAfterSimpleTypeReadCallback   OnAfterReadInterface[SimpleType]

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
	return "github.com/fullstack-lang/gongxsd/go/models"
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
	CommitComplexType(complextype *ComplexType)
	CheckoutComplexType(complextype *ComplexType)
	CommitElement(element *Element)
	CheckoutElement(element *Element)
	CommitEnumeration(enumeration *Enumeration)
	CheckoutEnumeration(enumeration *Enumeration)
	CommitRestriction(restriction *Restriction)
	CheckoutRestriction(restriction *Restriction)
	CommitSchema(schema *Schema)
	CheckoutSchema(schema *Schema)
	CommitSequence(sequence *Sequence)
	CheckoutSequence(sequence *Sequence)
	CommitSimpleType(simpletype *SimpleType)
	CheckoutSimpleType(simpletype *SimpleType)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

func NewStage(path string) (stage *StageStruct) {

	stage = &StageStruct{ // insertion point for array initiatialisation
		ComplexTypes:           make(map[*ComplexType]any),
		ComplexTypes_mapString: make(map[string]*ComplexType),

		Elements:           make(map[*Element]any),
		Elements_mapString: make(map[string]*Element),

		Enumerations:           make(map[*Enumeration]any),
		Enumerations_mapString: make(map[string]*Enumeration),

		Restrictions:           make(map[*Restriction]any),
		Restrictions_mapString: make(map[string]*Restriction),

		Schemas:           make(map[*Schema]any),
		Schemas_mapString: make(map[string]*Schema),

		Sequences:           make(map[*Sequence]any),
		Sequences_mapString: make(map[string]*Sequence),

		SimpleTypes:           make(map[*SimpleType]any),
		SimpleTypes_mapString: make(map[string]*SimpleType),

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
	stage.Map_GongStructName_InstancesNb["ComplexType"] = len(stage.ComplexTypes)
	stage.Map_GongStructName_InstancesNb["Element"] = len(stage.Elements)
	stage.Map_GongStructName_InstancesNb["Enumeration"] = len(stage.Enumerations)
	stage.Map_GongStructName_InstancesNb["Restriction"] = len(stage.Restrictions)
	stage.Map_GongStructName_InstancesNb["Schema"] = len(stage.Schemas)
	stage.Map_GongStructName_InstancesNb["Sequence"] = len(stage.Sequences)
	stage.Map_GongStructName_InstancesNb["SimpleType"] = len(stage.SimpleTypes)

}

func (stage *StageStruct) Checkout() {
	if stage.BackRepo != nil {
		stage.BackRepo.Checkout(stage)
	}

	stage.ComputeReverseMaps()
	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["ComplexType"] = len(stage.ComplexTypes)
	stage.Map_GongStructName_InstancesNb["Element"] = len(stage.Elements)
	stage.Map_GongStructName_InstancesNb["Enumeration"] = len(stage.Enumerations)
	stage.Map_GongStructName_InstancesNb["Restriction"] = len(stage.Restrictions)
	stage.Map_GongStructName_InstancesNb["Schema"] = len(stage.Schemas)
	stage.Map_GongStructName_InstancesNb["Sequence"] = len(stage.Sequences)
	stage.Map_GongStructName_InstancesNb["SimpleType"] = len(stage.SimpleTypes)

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

// Stage puts element to the model stage
func (element *Element) Stage(stage *StageStruct) *Element {
	stage.Elements[element] = __member
	stage.Elements_mapString[element.Name] = element

	return element
}

// Unstage removes element off the model stage
func (element *Element) Unstage(stage *StageStruct) *Element {
	delete(stage.Elements, element)
	delete(stage.Elements_mapString, element.Name)
	return element
}

// UnstageVoid removes element off the model stage
func (element *Element) UnstageVoid(stage *StageStruct) {
	delete(stage.Elements, element)
	delete(stage.Elements_mapString, element.Name)
}

// commit element to the back repo (if it is already staged)
func (element *Element) Commit(stage *StageStruct) *Element {
	if _, ok := stage.Elements[element]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitElement(element)
		}
	}
	return element
}

func (element *Element) CommitVoid(stage *StageStruct) {
	element.Commit(stage)
}

// Checkout element to the back repo (if it is already staged)
func (element *Element) Checkout(stage *StageStruct) *Element {
	if _, ok := stage.Elements[element]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutElement(element)
		}
	}
	return element
}

// for satisfaction of GongStruct interface
func (element *Element) GetName() (res string) {
	return element.Name
}

// Stage puts enumeration to the model stage
func (enumeration *Enumeration) Stage(stage *StageStruct) *Enumeration {
	stage.Enumerations[enumeration] = __member
	stage.Enumerations_mapString[enumeration.Name] = enumeration

	return enumeration
}

// Unstage removes enumeration off the model stage
func (enumeration *Enumeration) Unstage(stage *StageStruct) *Enumeration {
	delete(stage.Enumerations, enumeration)
	delete(stage.Enumerations_mapString, enumeration.Name)
	return enumeration
}

// UnstageVoid removes enumeration off the model stage
func (enumeration *Enumeration) UnstageVoid(stage *StageStruct) {
	delete(stage.Enumerations, enumeration)
	delete(stage.Enumerations_mapString, enumeration.Name)
}

// commit enumeration to the back repo (if it is already staged)
func (enumeration *Enumeration) Commit(stage *StageStruct) *Enumeration {
	if _, ok := stage.Enumerations[enumeration]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitEnumeration(enumeration)
		}
	}
	return enumeration
}

func (enumeration *Enumeration) CommitVoid(stage *StageStruct) {
	enumeration.Commit(stage)
}

// Checkout enumeration to the back repo (if it is already staged)
func (enumeration *Enumeration) Checkout(stage *StageStruct) *Enumeration {
	if _, ok := stage.Enumerations[enumeration]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutEnumeration(enumeration)
		}
	}
	return enumeration
}

// for satisfaction of GongStruct interface
func (enumeration *Enumeration) GetName() (res string) {
	return enumeration.Name
}

// Stage puts restriction to the model stage
func (restriction *Restriction) Stage(stage *StageStruct) *Restriction {
	stage.Restrictions[restriction] = __member
	stage.Restrictions_mapString[restriction.Name] = restriction

	return restriction
}

// Unstage removes restriction off the model stage
func (restriction *Restriction) Unstage(stage *StageStruct) *Restriction {
	delete(stage.Restrictions, restriction)
	delete(stage.Restrictions_mapString, restriction.Name)
	return restriction
}

// UnstageVoid removes restriction off the model stage
func (restriction *Restriction) UnstageVoid(stage *StageStruct) {
	delete(stage.Restrictions, restriction)
	delete(stage.Restrictions_mapString, restriction.Name)
}

// commit restriction to the back repo (if it is already staged)
func (restriction *Restriction) Commit(stage *StageStruct) *Restriction {
	if _, ok := stage.Restrictions[restriction]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitRestriction(restriction)
		}
	}
	return restriction
}

func (restriction *Restriction) CommitVoid(stage *StageStruct) {
	restriction.Commit(stage)
}

// Checkout restriction to the back repo (if it is already staged)
func (restriction *Restriction) Checkout(stage *StageStruct) *Restriction {
	if _, ok := stage.Restrictions[restriction]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutRestriction(restriction)
		}
	}
	return restriction
}

// for satisfaction of GongStruct interface
func (restriction *Restriction) GetName() (res string) {
	return restriction.Name
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

// Stage puts sequence to the model stage
func (sequence *Sequence) Stage(stage *StageStruct) *Sequence {
	stage.Sequences[sequence] = __member
	stage.Sequences_mapString[sequence.Name] = sequence

	return sequence
}

// Unstage removes sequence off the model stage
func (sequence *Sequence) Unstage(stage *StageStruct) *Sequence {
	delete(stage.Sequences, sequence)
	delete(stage.Sequences_mapString, sequence.Name)
	return sequence
}

// UnstageVoid removes sequence off the model stage
func (sequence *Sequence) UnstageVoid(stage *StageStruct) {
	delete(stage.Sequences, sequence)
	delete(stage.Sequences_mapString, sequence.Name)
}

// commit sequence to the back repo (if it is already staged)
func (sequence *Sequence) Commit(stage *StageStruct) *Sequence {
	if _, ok := stage.Sequences[sequence]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSequence(sequence)
		}
	}
	return sequence
}

func (sequence *Sequence) CommitVoid(stage *StageStruct) {
	sequence.Commit(stage)
}

// Checkout sequence to the back repo (if it is already staged)
func (sequence *Sequence) Checkout(stage *StageStruct) *Sequence {
	if _, ok := stage.Sequences[sequence]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutSequence(sequence)
		}
	}
	return sequence
}

// for satisfaction of GongStruct interface
func (sequence *Sequence) GetName() (res string) {
	return sequence.Name
}

// Stage puts simpletype to the model stage
func (simpletype *SimpleType) Stage(stage *StageStruct) *SimpleType {
	stage.SimpleTypes[simpletype] = __member
	stage.SimpleTypes_mapString[simpletype.Name] = simpletype

	return simpletype
}

// Unstage removes simpletype off the model stage
func (simpletype *SimpleType) Unstage(stage *StageStruct) *SimpleType {
	delete(stage.SimpleTypes, simpletype)
	delete(stage.SimpleTypes_mapString, simpletype.Name)
	return simpletype
}

// UnstageVoid removes simpletype off the model stage
func (simpletype *SimpleType) UnstageVoid(stage *StageStruct) {
	delete(stage.SimpleTypes, simpletype)
	delete(stage.SimpleTypes_mapString, simpletype.Name)
}

// commit simpletype to the back repo (if it is already staged)
func (simpletype *SimpleType) Commit(stage *StageStruct) *SimpleType {
	if _, ok := stage.SimpleTypes[simpletype]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSimpleType(simpletype)
		}
	}
	return simpletype
}

func (simpletype *SimpleType) CommitVoid(stage *StageStruct) {
	simpletype.Commit(stage)
}

// Checkout simpletype to the back repo (if it is already staged)
func (simpletype *SimpleType) Checkout(stage *StageStruct) *SimpleType {
	if _, ok := stage.SimpleTypes[simpletype]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutSimpleType(simpletype)
		}
	}
	return simpletype
}

// for satisfaction of GongStruct interface
func (simpletype *SimpleType) GetName() (res string) {
	return simpletype.Name
}

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMComplexType(ComplexType *ComplexType)
	CreateORMElement(Element *Element)
	CreateORMEnumeration(Enumeration *Enumeration)
	CreateORMRestriction(Restriction *Restriction)
	CreateORMSchema(Schema *Schema)
	CreateORMSequence(Sequence *Sequence)
	CreateORMSimpleType(SimpleType *SimpleType)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMComplexType(ComplexType *ComplexType)
	DeleteORMElement(Element *Element)
	DeleteORMEnumeration(Enumeration *Enumeration)
	DeleteORMRestriction(Restriction *Restriction)
	DeleteORMSchema(Schema *Schema)
	DeleteORMSequence(Sequence *Sequence)
	DeleteORMSimpleType(SimpleType *SimpleType)
}

func (stage *StageStruct) Reset() { // insertion point for array reset
	stage.ComplexTypes = make(map[*ComplexType]any)
	stage.ComplexTypes_mapString = make(map[string]*ComplexType)

	stage.Elements = make(map[*Element]any)
	stage.Elements_mapString = make(map[string]*Element)

	stage.Enumerations = make(map[*Enumeration]any)
	stage.Enumerations_mapString = make(map[string]*Enumeration)

	stage.Restrictions = make(map[*Restriction]any)
	stage.Restrictions_mapString = make(map[string]*Restriction)

	stage.Schemas = make(map[*Schema]any)
	stage.Schemas_mapString = make(map[string]*Schema)

	stage.Sequences = make(map[*Sequence]any)
	stage.Sequences_mapString = make(map[string]*Sequence)

	stage.SimpleTypes = make(map[*SimpleType]any)
	stage.SimpleTypes_mapString = make(map[string]*SimpleType)

}

func (stage *StageStruct) Nil() { // insertion point for array nil
	stage.ComplexTypes = nil
	stage.ComplexTypes_mapString = nil

	stage.Elements = nil
	stage.Elements_mapString = nil

	stage.Enumerations = nil
	stage.Enumerations_mapString = nil

	stage.Restrictions = nil
	stage.Restrictions_mapString = nil

	stage.Schemas = nil
	stage.Schemas_mapString = nil

	stage.Sequences = nil
	stage.Sequences_mapString = nil

	stage.SimpleTypes = nil
	stage.SimpleTypes_mapString = nil

}

func (stage *StageStruct) Unstage() { // insertion point for array nil
	for complextype := range stage.ComplexTypes {
		complextype.Unstage(stage)
	}

	for element := range stage.Elements {
		element.Unstage(stage)
	}

	for enumeration := range stage.Enumerations {
		enumeration.Unstage(stage)
	}

	for restriction := range stage.Restrictions {
		restriction.Unstage(stage)
	}

	for schema := range stage.Schemas {
		schema.Unstage(stage)
	}

	for sequence := range stage.Sequences {
		sequence.Unstage(stage)
	}

	for simpletype := range stage.SimpleTypes {
		simpletype.Unstage(stage)
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
	case map[*ComplexType]any:
		return any(&stage.ComplexTypes).(*Type)
	case map[*Element]any:
		return any(&stage.Elements).(*Type)
	case map[*Enumeration]any:
		return any(&stage.Enumerations).(*Type)
	case map[*Restriction]any:
		return any(&stage.Restrictions).(*Type)
	case map[*Schema]any:
		return any(&stage.Schemas).(*Type)
	case map[*Sequence]any:
		return any(&stage.Sequences).(*Type)
	case map[*SimpleType]any:
		return any(&stage.SimpleTypes).(*Type)
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
	case map[string]*ComplexType:
		return any(&stage.ComplexTypes_mapString).(*Type)
	case map[string]*Element:
		return any(&stage.Elements_mapString).(*Type)
	case map[string]*Enumeration:
		return any(&stage.Enumerations_mapString).(*Type)
	case map[string]*Restriction:
		return any(&stage.Restrictions_mapString).(*Type)
	case map[string]*Schema:
		return any(&stage.Schemas_mapString).(*Type)
	case map[string]*Sequence:
		return any(&stage.Sequences_mapString).(*Type)
	case map[string]*SimpleType:
		return any(&stage.SimpleTypes_mapString).(*Type)
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
	case ComplexType:
		return any(&stage.ComplexTypes).(*map[*Type]any)
	case Element:
		return any(&stage.Elements).(*map[*Type]any)
	case Enumeration:
		return any(&stage.Enumerations).(*map[*Type]any)
	case Restriction:
		return any(&stage.Restrictions).(*map[*Type]any)
	case Schema:
		return any(&stage.Schemas).(*map[*Type]any)
	case Sequence:
		return any(&stage.Sequences).(*map[*Type]any)
	case SimpleType:
		return any(&stage.SimpleTypes).(*map[*Type]any)
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
	case *ComplexType:
		return any(&stage.ComplexTypes).(*map[Type]any)
	case *Element:
		return any(&stage.Elements).(*map[Type]any)
	case *Enumeration:
		return any(&stage.Enumerations).(*map[Type]any)
	case *Restriction:
		return any(&stage.Restrictions).(*map[Type]any)
	case *Schema:
		return any(&stage.Schemas).(*map[Type]any)
	case *Sequence:
		return any(&stage.Sequences).(*map[Type]any)
	case *SimpleType:
		return any(&stage.SimpleTypes).(*map[Type]any)
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
	case ComplexType:
		return any(&stage.ComplexTypes_mapString).(*map[string]*Type)
	case Element:
		return any(&stage.Elements_mapString).(*map[string]*Type)
	case Enumeration:
		return any(&stage.Enumerations_mapString).(*map[string]*Type)
	case Restriction:
		return any(&stage.Restrictions_mapString).(*map[string]*Type)
	case Schema:
		return any(&stage.Schemas_mapString).(*map[string]*Type)
	case Sequence:
		return any(&stage.Sequences_mapString).(*map[string]*Type)
	case SimpleType:
		return any(&stage.SimpleTypes_mapString).(*map[string]*Type)
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
	case ComplexType:
		return any(&ComplexType{
			// Initialisation of associations
			// field is initialized with an instance of Sequence with the name of the field
			Sequence: &Sequence{Name: "Sequence"},
		}).(*Type)
	case Element:
		return any(&Element{
			// Initialisation of associations
			// field is initialized with an instance of SimpleType with the name of the field
			SimpleType: &SimpleType{Name: "SimpleType"},
			// field is initialized with an instance of ComplexType with the name of the field
			ComplexType: &ComplexType{Name: "ComplexType"},
		}).(*Type)
	case Enumeration:
		return any(&Enumeration{
			// Initialisation of associations
		}).(*Type)
	case Restriction:
		return any(&Restriction{
			// Initialisation of associations
			// field is initialized with an instance of Enumeration with the name of the field
			Enumerations: []*Enumeration{{Name: "Enumerations"}},
		}).(*Type)
	case Schema:
		return any(&Schema{
			// Initialisation of associations
			// field is initialized with an instance of Element with the name of the field
			Elements: []*Element{{Name: "Elements"}},
			// field is initialized with an instance of SimpleType with the name of the field
			SimpleTypes: []*SimpleType{{Name: "SimpleTypes"}},
			// field is initialized with an instance of ComplexType with the name of the field
			ComplexTypes: []*ComplexType{{Name: "ComplexTypes"}},
		}).(*Type)
	case Sequence:
		return any(&Sequence{
			// Initialisation of associations
			// field is initialized with an instance of Element with the name of the field
			Elements: []*Element{{Name: "Elements"}},
		}).(*Type)
	case SimpleType:
		return any(&SimpleType{
			// Initialisation of associations
			// field is initialized with an instance of Restriction with the name of the field
			Restriction: &Restriction{Name: "Restriction"},
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
	// reverse maps of direct associations of ComplexType
	case ComplexType:
		switch fieldname {
		// insertion point for per direct association field
		case "Sequence":
			res := make(map[*Sequence][]*ComplexType)
			for complextype := range stage.ComplexTypes {
				if complextype.Sequence != nil {
					sequence_ := complextype.Sequence
					var complextypes []*ComplexType
					_, ok := res[sequence_]
					if ok {
						complextypes = res[sequence_]
					} else {
						complextypes = make([]*ComplexType, 0)
					}
					complextypes = append(complextypes, complextype)
					res[sequence_] = complextypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Element
	case Element:
		switch fieldname {
		// insertion point for per direct association field
		case "SimpleType":
			res := make(map[*SimpleType][]*Element)
			for element := range stage.Elements {
				if element.SimpleType != nil {
					simpletype_ := element.SimpleType
					var elements []*Element
					_, ok := res[simpletype_]
					if ok {
						elements = res[simpletype_]
					} else {
						elements = make([]*Element, 0)
					}
					elements = append(elements, element)
					res[simpletype_] = elements
				}
			}
			return any(res).(map[*End][]*Start)
		case "ComplexType":
			res := make(map[*ComplexType][]*Element)
			for element := range stage.Elements {
				if element.ComplexType != nil {
					complextype_ := element.ComplexType
					var elements []*Element
					_, ok := res[complextype_]
					if ok {
						elements = res[complextype_]
					} else {
						elements = make([]*Element, 0)
					}
					elements = append(elements, element)
					res[complextype_] = elements
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Enumeration
	case Enumeration:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Restriction
	case Restriction:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Schema
	case Schema:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Sequence
	case Sequence:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SimpleType
	case SimpleType:
		switch fieldname {
		// insertion point for per direct association field
		case "Restriction":
			res := make(map[*Restriction][]*SimpleType)
			for simpletype := range stage.SimpleTypes {
				if simpletype.Restriction != nil {
					restriction_ := simpletype.Restriction
					var simpletypes []*SimpleType
					_, ok := res[restriction_]
					if ok {
						simpletypes = res[restriction_]
					} else {
						simpletypes = make([]*SimpleType, 0)
					}
					simpletypes = append(simpletypes, simpletype)
					res[restriction_] = simpletypes
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
	// reverse maps of direct associations of ComplexType
	case ComplexType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Element
	case Element:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Enumeration
	case Enumeration:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Restriction
	case Restriction:
		switch fieldname {
		// insertion point for per direct association field
		case "Enumerations":
			res := make(map[*Enumeration]*Restriction)
			for restriction := range stage.Restrictions {
				for _, enumeration_ := range restriction.Enumerations {
					res[enumeration_] = restriction
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of Schema
	case Schema:
		switch fieldname {
		// insertion point for per direct association field
		case "Elements":
			res := make(map[*Element]*Schema)
			for schema := range stage.Schemas {
				for _, element_ := range schema.Elements {
					res[element_] = schema
				}
			}
			return any(res).(map[*End]*Start)
		case "SimpleTypes":
			res := make(map[*SimpleType]*Schema)
			for schema := range stage.Schemas {
				for _, simpletype_ := range schema.SimpleTypes {
					res[simpletype_] = schema
				}
			}
			return any(res).(map[*End]*Start)
		case "ComplexTypes":
			res := make(map[*ComplexType]*Schema)
			for schema := range stage.Schemas {
				for _, complextype_ := range schema.ComplexTypes {
					res[complextype_] = schema
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of Sequence
	case Sequence:
		switch fieldname {
		// insertion point for per direct association field
		case "Elements":
			res := make(map[*Element]*Sequence)
			for sequence := range stage.Sequences {
				for _, element_ := range sequence.Elements {
					res[element_] = sequence
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of SimpleType
	case SimpleType:
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
	case ComplexType:
		res = "ComplexType"
	case Element:
		res = "Element"
	case Enumeration:
		res = "Enumeration"
	case Restriction:
		res = "Restriction"
	case Schema:
		res = "Schema"
	case Sequence:
		res = "Sequence"
	case SimpleType:
		res = "SimpleType"
	}
	return res
}

// GetPointerToGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetPointerToGongstructName[Type PointerToGongstruct]() (res string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case *ComplexType:
		res = "ComplexType"
	case *Element:
		res = "Element"
	case *Enumeration:
		res = "Enumeration"
	case *Restriction:
		res = "Restriction"
	case *Schema:
		res = "Schema"
	case *Sequence:
		res = "Sequence"
	case *SimpleType:
		res = "SimpleType"
	}
	return res
}

// GetFields return the array of the fields
func GetFields[Type Gongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case ComplexType:
		res = []string{"Name", "NameXML", "Sequence"}
	case Element:
		res = []string{"Name", "NameXSD", "Type", "SimpleType", "ComplexType"}
	case Enumeration:
		res = []string{"Name", "Value"}
	case Restriction:
		res = []string{"Name", "Base", "Enumerations"}
	case Schema:
		res = []string{"Name", "Elements", "SimpleTypes", "ComplexTypes"}
	case Sequence:
		res = []string{"Name", "Elements"}
	case SimpleType:
		res = []string{"Name", "NameXSD", "Restriction"}
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
	case ComplexType:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Schema"
		rf.Fieldname = "ComplexTypes"
		res = append(res, rf)
	case Element:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Schema"
		rf.Fieldname = "Elements"
		res = append(res, rf)
		rf.GongstructName = "Sequence"
		rf.Fieldname = "Elements"
		res = append(res, rf)
	case Enumeration:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Restriction"
		rf.Fieldname = "Enumerations"
		res = append(res, rf)
	case Restriction:
		var rf ReverseField
		_ = rf
	case Schema:
		var rf ReverseField
		_ = rf
	case Sequence:
		var rf ReverseField
		_ = rf
	case SimpleType:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Schema"
		rf.Fieldname = "SimpleTypes"
		res = append(res, rf)
	}
	return
}

// GetFieldsFromPointer return the array of the fields
func GetFieldsFromPointer[Type PointerToGongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case *ComplexType:
		res = []string{"Name", "NameXML", "Sequence"}
	case *Element:
		res = []string{"Name", "NameXSD", "Type", "SimpleType", "ComplexType"}
	case *Enumeration:
		res = []string{"Name", "Value"}
	case *Restriction:
		res = []string{"Name", "Base", "Enumerations"}
	case *Schema:
		res = []string{"Name", "Elements", "SimpleTypes", "ComplexTypes"}
	case *Sequence:
		res = []string{"Name", "Elements"}
	case *SimpleType:
		res = []string{"Name", "NameXSD", "Restriction"}
	}
	return
}

func GetFieldStringValueFromPointer[Type PointerToGongstruct](instance Type, fieldName string) (res string) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
	case *ComplexType:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "NameXML":
			res = inferedInstance.NameXML
		case "Sequence":
			if inferedInstance.Sequence != nil {
				res = inferedInstance.Sequence.Name
			}
		}
	case *Element:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "NameXSD":
			res = inferedInstance.NameXSD
		case "Type":
			res = inferedInstance.Type
		case "SimpleType":
			if inferedInstance.SimpleType != nil {
				res = inferedInstance.SimpleType.Name
			}
		case "ComplexType":
			if inferedInstance.ComplexType != nil {
				res = inferedInstance.ComplexType.Name
			}
		}
	case *Enumeration:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Value":
			res = inferedInstance.Value
		}
	case *Restriction:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Base":
			res = inferedInstance.Base
		case "Enumerations":
			for idx, __instance__ := range inferedInstance.Enumerations {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case *Schema:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Elements":
			for idx, __instance__ := range inferedInstance.Elements {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "SimpleTypes":
			for idx, __instance__ := range inferedInstance.SimpleTypes {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "ComplexTypes":
			for idx, __instance__ := range inferedInstance.ComplexTypes {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case *Sequence:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Elements":
			for idx, __instance__ := range inferedInstance.Elements {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case *SimpleType:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "NameXSD":
			res = inferedInstance.NameXSD
		case "Restriction":
			if inferedInstance.Restriction != nil {
				res = inferedInstance.Restriction.Name
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
	case ComplexType:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "NameXML":
			res = inferedInstance.NameXML
		case "Sequence":
			if inferedInstance.Sequence != nil {
				res = inferedInstance.Sequence.Name
			}
		}
	case Element:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "NameXSD":
			res = inferedInstance.NameXSD
		case "Type":
			res = inferedInstance.Type
		case "SimpleType":
			if inferedInstance.SimpleType != nil {
				res = inferedInstance.SimpleType.Name
			}
		case "ComplexType":
			if inferedInstance.ComplexType != nil {
				res = inferedInstance.ComplexType.Name
			}
		}
	case Enumeration:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Value":
			res = inferedInstance.Value
		}
	case Restriction:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Base":
			res = inferedInstance.Base
		case "Enumerations":
			for idx, __instance__ := range inferedInstance.Enumerations {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case Schema:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Elements":
			for idx, __instance__ := range inferedInstance.Elements {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "SimpleTypes":
			for idx, __instance__ := range inferedInstance.SimpleTypes {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "ComplexTypes":
			for idx, __instance__ := range inferedInstance.ComplexTypes {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case Sequence:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Elements":
			for idx, __instance__ := range inferedInstance.Elements {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case SimpleType:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "NameXSD":
			res = inferedInstance.NameXSD
		case "Restriction":
			if inferedInstance.Restriction != nil {
				res = inferedInstance.Restriction.Name
			}
		}
	default:
		_ = inferedInstance
	}
	return
}

// Last line of the template
