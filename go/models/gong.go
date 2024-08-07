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

	Lengths           map[*Length]any
	Lengths_mapString map[string]*Length

	// insertion point for slice of pointers maps
	OnAfterLengthCreateCallback OnAfterCreateInterface[Length]
	OnAfterLengthUpdateCallback OnAfterUpdateInterface[Length]
	OnAfterLengthDeleteCallback OnAfterDeleteInterface[Length]
	OnAfterLengthReadCallback   OnAfterReadInterface[Length]

	MaxInclusives           map[*MaxInclusive]any
	MaxInclusives_mapString map[string]*MaxInclusive

	// insertion point for slice of pointers maps
	OnAfterMaxInclusiveCreateCallback OnAfterCreateInterface[MaxInclusive]
	OnAfterMaxInclusiveUpdateCallback OnAfterUpdateInterface[MaxInclusive]
	OnAfterMaxInclusiveDeleteCallback OnAfterDeleteInterface[MaxInclusive]
	OnAfterMaxInclusiveReadCallback   OnAfterReadInterface[MaxInclusive]

	MaxLengths           map[*MaxLength]any
	MaxLengths_mapString map[string]*MaxLength

	// insertion point for slice of pointers maps
	OnAfterMaxLengthCreateCallback OnAfterCreateInterface[MaxLength]
	OnAfterMaxLengthUpdateCallback OnAfterUpdateInterface[MaxLength]
	OnAfterMaxLengthDeleteCallback OnAfterDeleteInterface[MaxLength]
	OnAfterMaxLengthReadCallback   OnAfterReadInterface[MaxLength]

	MinInclusives           map[*MinInclusive]any
	MinInclusives_mapString map[string]*MinInclusive

	// insertion point for slice of pointers maps
	OnAfterMinInclusiveCreateCallback OnAfterCreateInterface[MinInclusive]
	OnAfterMinInclusiveUpdateCallback OnAfterUpdateInterface[MinInclusive]
	OnAfterMinInclusiveDeleteCallback OnAfterDeleteInterface[MinInclusive]
	OnAfterMinInclusiveReadCallback   OnAfterReadInterface[MinInclusive]

	MinLengths           map[*MinLength]any
	MinLengths_mapString map[string]*MinLength

	// insertion point for slice of pointers maps
	OnAfterMinLengthCreateCallback OnAfterCreateInterface[MinLength]
	OnAfterMinLengthUpdateCallback OnAfterUpdateInterface[MinLength]
	OnAfterMinLengthDeleteCallback OnAfterDeleteInterface[MinLength]
	OnAfterMinLengthReadCallback   OnAfterReadInterface[MinLength]

	Patterns           map[*Pattern]any
	Patterns_mapString map[string]*Pattern

	// insertion point for slice of pointers maps
	OnAfterPatternCreateCallback OnAfterCreateInterface[Pattern]
	OnAfterPatternUpdateCallback OnAfterUpdateInterface[Pattern]
	OnAfterPatternDeleteCallback OnAfterDeleteInterface[Pattern]
	OnAfterPatternReadCallback   OnAfterReadInterface[Pattern]

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

	WhiteSpaces           map[*WhiteSpace]any
	WhiteSpaces_mapString map[string]*WhiteSpace

	// insertion point for slice of pointers maps
	OnAfterWhiteSpaceCreateCallback OnAfterCreateInterface[WhiteSpace]
	OnAfterWhiteSpaceUpdateCallback OnAfterUpdateInterface[WhiteSpace]
	OnAfterWhiteSpaceDeleteCallback OnAfterDeleteInterface[WhiteSpace]
	OnAfterWhiteSpaceReadCallback   OnAfterReadInterface[WhiteSpace]

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
	CommitAnnotation(annotation *Annotation)
	CheckoutAnnotation(annotation *Annotation)
	CommitComplexType(complextype *ComplexType)
	CheckoutComplexType(complextype *ComplexType)
	CommitDocumentation(documentation *Documentation)
	CheckoutDocumentation(documentation *Documentation)
	CommitElement(element *Element)
	CheckoutElement(element *Element)
	CommitEnumeration(enumeration *Enumeration)
	CheckoutEnumeration(enumeration *Enumeration)
	CommitLength(length *Length)
	CheckoutLength(length *Length)
	CommitMaxInclusive(maxinclusive *MaxInclusive)
	CheckoutMaxInclusive(maxinclusive *MaxInclusive)
	CommitMaxLength(maxlength *MaxLength)
	CheckoutMaxLength(maxlength *MaxLength)
	CommitMinInclusive(mininclusive *MinInclusive)
	CheckoutMinInclusive(mininclusive *MinInclusive)
	CommitMinLength(minlength *MinLength)
	CheckoutMinLength(minlength *MinLength)
	CommitPattern(pattern *Pattern)
	CheckoutPattern(pattern *Pattern)
	CommitRestriction(restriction *Restriction)
	CheckoutRestriction(restriction *Restriction)
	CommitSchema(schema *Schema)
	CheckoutSchema(schema *Schema)
	CommitSequence(sequence *Sequence)
	CheckoutSequence(sequence *Sequence)
	CommitSimpleType(simpletype *SimpleType)
	CheckoutSimpleType(simpletype *SimpleType)
	CommitWhiteSpace(whitespace *WhiteSpace)
	CheckoutWhiteSpace(whitespace *WhiteSpace)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

func NewStage(path string) (stage *StageStruct) {

	stage = &StageStruct{ // insertion point for array initiatialisation
		Annotations:           make(map[*Annotation]any),
		Annotations_mapString: make(map[string]*Annotation),

		ComplexTypes:           make(map[*ComplexType]any),
		ComplexTypes_mapString: make(map[string]*ComplexType),

		Documentations:           make(map[*Documentation]any),
		Documentations_mapString: make(map[string]*Documentation),

		Elements:           make(map[*Element]any),
		Elements_mapString: make(map[string]*Element),

		Enumerations:           make(map[*Enumeration]any),
		Enumerations_mapString: make(map[string]*Enumeration),

		Lengths:           make(map[*Length]any),
		Lengths_mapString: make(map[string]*Length),

		MaxInclusives:           make(map[*MaxInclusive]any),
		MaxInclusives_mapString: make(map[string]*MaxInclusive),

		MaxLengths:           make(map[*MaxLength]any),
		MaxLengths_mapString: make(map[string]*MaxLength),

		MinInclusives:           make(map[*MinInclusive]any),
		MinInclusives_mapString: make(map[string]*MinInclusive),

		MinLengths:           make(map[*MinLength]any),
		MinLengths_mapString: make(map[string]*MinLength),

		Patterns:           make(map[*Pattern]any),
		Patterns_mapString: make(map[string]*Pattern),

		Restrictions:           make(map[*Restriction]any),
		Restrictions_mapString: make(map[string]*Restriction),

		Schemas:           make(map[*Schema]any),
		Schemas_mapString: make(map[string]*Schema),

		Sequences:           make(map[*Sequence]any),
		Sequences_mapString: make(map[string]*Sequence),

		SimpleTypes:           make(map[*SimpleType]any),
		SimpleTypes_mapString: make(map[string]*SimpleType),

		WhiteSpaces:           make(map[*WhiteSpace]any),
		WhiteSpaces_mapString: make(map[string]*WhiteSpace),

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
	stage.Map_GongStructName_InstancesNb["ComplexType"] = len(stage.ComplexTypes)
	stage.Map_GongStructName_InstancesNb["Documentation"] = len(stage.Documentations)
	stage.Map_GongStructName_InstancesNb["Element"] = len(stage.Elements)
	stage.Map_GongStructName_InstancesNb["Enumeration"] = len(stage.Enumerations)
	stage.Map_GongStructName_InstancesNb["Length"] = len(stage.Lengths)
	stage.Map_GongStructName_InstancesNb["MaxInclusive"] = len(stage.MaxInclusives)
	stage.Map_GongStructName_InstancesNb["MaxLength"] = len(stage.MaxLengths)
	stage.Map_GongStructName_InstancesNb["MinInclusive"] = len(stage.MinInclusives)
	stage.Map_GongStructName_InstancesNb["MinLength"] = len(stage.MinLengths)
	stage.Map_GongStructName_InstancesNb["Pattern"] = len(stage.Patterns)
	stage.Map_GongStructName_InstancesNb["Restriction"] = len(stage.Restrictions)
	stage.Map_GongStructName_InstancesNb["Schema"] = len(stage.Schemas)
	stage.Map_GongStructName_InstancesNb["Sequence"] = len(stage.Sequences)
	stage.Map_GongStructName_InstancesNb["SimpleType"] = len(stage.SimpleTypes)
	stage.Map_GongStructName_InstancesNb["WhiteSpace"] = len(stage.WhiteSpaces)

}

func (stage *StageStruct) Checkout() {
	if stage.BackRepo != nil {
		stage.BackRepo.Checkout(stage)
	}

	stage.ComputeReverseMaps()
	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["Annotation"] = len(stage.Annotations)
	stage.Map_GongStructName_InstancesNb["ComplexType"] = len(stage.ComplexTypes)
	stage.Map_GongStructName_InstancesNb["Documentation"] = len(stage.Documentations)
	stage.Map_GongStructName_InstancesNb["Element"] = len(stage.Elements)
	stage.Map_GongStructName_InstancesNb["Enumeration"] = len(stage.Enumerations)
	stage.Map_GongStructName_InstancesNb["Length"] = len(stage.Lengths)
	stage.Map_GongStructName_InstancesNb["MaxInclusive"] = len(stage.MaxInclusives)
	stage.Map_GongStructName_InstancesNb["MaxLength"] = len(stage.MaxLengths)
	stage.Map_GongStructName_InstancesNb["MinInclusive"] = len(stage.MinInclusives)
	stage.Map_GongStructName_InstancesNb["MinLength"] = len(stage.MinLengths)
	stage.Map_GongStructName_InstancesNb["Pattern"] = len(stage.Patterns)
	stage.Map_GongStructName_InstancesNb["Restriction"] = len(stage.Restrictions)
	stage.Map_GongStructName_InstancesNb["Schema"] = len(stage.Schemas)
	stage.Map_GongStructName_InstancesNb["Sequence"] = len(stage.Sequences)
	stage.Map_GongStructName_InstancesNb["SimpleType"] = len(stage.SimpleTypes)
	stage.Map_GongStructName_InstancesNb["WhiteSpace"] = len(stage.WhiteSpaces)

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

// Stage puts length to the model stage
func (length *Length) Stage(stage *StageStruct) *Length {
	stage.Lengths[length] = __member
	stage.Lengths_mapString[length.Name] = length

	return length
}

// Unstage removes length off the model stage
func (length *Length) Unstage(stage *StageStruct) *Length {
	delete(stage.Lengths, length)
	delete(stage.Lengths_mapString, length.Name)
	return length
}

// UnstageVoid removes length off the model stage
func (length *Length) UnstageVoid(stage *StageStruct) {
	delete(stage.Lengths, length)
	delete(stage.Lengths_mapString, length.Name)
}

// commit length to the back repo (if it is already staged)
func (length *Length) Commit(stage *StageStruct) *Length {
	if _, ok := stage.Lengths[length]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitLength(length)
		}
	}
	return length
}

func (length *Length) CommitVoid(stage *StageStruct) {
	length.Commit(stage)
}

// Checkout length to the back repo (if it is already staged)
func (length *Length) Checkout(stage *StageStruct) *Length {
	if _, ok := stage.Lengths[length]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutLength(length)
		}
	}
	return length
}

// for satisfaction of GongStruct interface
func (length *Length) GetName() (res string) {
	return length.Name
}

// Stage puts maxinclusive to the model stage
func (maxinclusive *MaxInclusive) Stage(stage *StageStruct) *MaxInclusive {
	stage.MaxInclusives[maxinclusive] = __member
	stage.MaxInclusives_mapString[maxinclusive.Name] = maxinclusive

	return maxinclusive
}

// Unstage removes maxinclusive off the model stage
func (maxinclusive *MaxInclusive) Unstage(stage *StageStruct) *MaxInclusive {
	delete(stage.MaxInclusives, maxinclusive)
	delete(stage.MaxInclusives_mapString, maxinclusive.Name)
	return maxinclusive
}

// UnstageVoid removes maxinclusive off the model stage
func (maxinclusive *MaxInclusive) UnstageVoid(stage *StageStruct) {
	delete(stage.MaxInclusives, maxinclusive)
	delete(stage.MaxInclusives_mapString, maxinclusive.Name)
}

// commit maxinclusive to the back repo (if it is already staged)
func (maxinclusive *MaxInclusive) Commit(stage *StageStruct) *MaxInclusive {
	if _, ok := stage.MaxInclusives[maxinclusive]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitMaxInclusive(maxinclusive)
		}
	}
	return maxinclusive
}

func (maxinclusive *MaxInclusive) CommitVoid(stage *StageStruct) {
	maxinclusive.Commit(stage)
}

// Checkout maxinclusive to the back repo (if it is already staged)
func (maxinclusive *MaxInclusive) Checkout(stage *StageStruct) *MaxInclusive {
	if _, ok := stage.MaxInclusives[maxinclusive]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutMaxInclusive(maxinclusive)
		}
	}
	return maxinclusive
}

// for satisfaction of GongStruct interface
func (maxinclusive *MaxInclusive) GetName() (res string) {
	return maxinclusive.Name
}

// Stage puts maxlength to the model stage
func (maxlength *MaxLength) Stage(stage *StageStruct) *MaxLength {
	stage.MaxLengths[maxlength] = __member
	stage.MaxLengths_mapString[maxlength.Name] = maxlength

	return maxlength
}

// Unstage removes maxlength off the model stage
func (maxlength *MaxLength) Unstage(stage *StageStruct) *MaxLength {
	delete(stage.MaxLengths, maxlength)
	delete(stage.MaxLengths_mapString, maxlength.Name)
	return maxlength
}

// UnstageVoid removes maxlength off the model stage
func (maxlength *MaxLength) UnstageVoid(stage *StageStruct) {
	delete(stage.MaxLengths, maxlength)
	delete(stage.MaxLengths_mapString, maxlength.Name)
}

// commit maxlength to the back repo (if it is already staged)
func (maxlength *MaxLength) Commit(stage *StageStruct) *MaxLength {
	if _, ok := stage.MaxLengths[maxlength]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitMaxLength(maxlength)
		}
	}
	return maxlength
}

func (maxlength *MaxLength) CommitVoid(stage *StageStruct) {
	maxlength.Commit(stage)
}

// Checkout maxlength to the back repo (if it is already staged)
func (maxlength *MaxLength) Checkout(stage *StageStruct) *MaxLength {
	if _, ok := stage.MaxLengths[maxlength]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutMaxLength(maxlength)
		}
	}
	return maxlength
}

// for satisfaction of GongStruct interface
func (maxlength *MaxLength) GetName() (res string) {
	return maxlength.Name
}

// Stage puts mininclusive to the model stage
func (mininclusive *MinInclusive) Stage(stage *StageStruct) *MinInclusive {
	stage.MinInclusives[mininclusive] = __member
	stage.MinInclusives_mapString[mininclusive.Name] = mininclusive

	return mininclusive
}

// Unstage removes mininclusive off the model stage
func (mininclusive *MinInclusive) Unstage(stage *StageStruct) *MinInclusive {
	delete(stage.MinInclusives, mininclusive)
	delete(stage.MinInclusives_mapString, mininclusive.Name)
	return mininclusive
}

// UnstageVoid removes mininclusive off the model stage
func (mininclusive *MinInclusive) UnstageVoid(stage *StageStruct) {
	delete(stage.MinInclusives, mininclusive)
	delete(stage.MinInclusives_mapString, mininclusive.Name)
}

// commit mininclusive to the back repo (if it is already staged)
func (mininclusive *MinInclusive) Commit(stage *StageStruct) *MinInclusive {
	if _, ok := stage.MinInclusives[mininclusive]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitMinInclusive(mininclusive)
		}
	}
	return mininclusive
}

func (mininclusive *MinInclusive) CommitVoid(stage *StageStruct) {
	mininclusive.Commit(stage)
}

// Checkout mininclusive to the back repo (if it is already staged)
func (mininclusive *MinInclusive) Checkout(stage *StageStruct) *MinInclusive {
	if _, ok := stage.MinInclusives[mininclusive]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutMinInclusive(mininclusive)
		}
	}
	return mininclusive
}

// for satisfaction of GongStruct interface
func (mininclusive *MinInclusive) GetName() (res string) {
	return mininclusive.Name
}

// Stage puts minlength to the model stage
func (minlength *MinLength) Stage(stage *StageStruct) *MinLength {
	stage.MinLengths[minlength] = __member
	stage.MinLengths_mapString[minlength.Name] = minlength

	return minlength
}

// Unstage removes minlength off the model stage
func (minlength *MinLength) Unstage(stage *StageStruct) *MinLength {
	delete(stage.MinLengths, minlength)
	delete(stage.MinLengths_mapString, minlength.Name)
	return minlength
}

// UnstageVoid removes minlength off the model stage
func (minlength *MinLength) UnstageVoid(stage *StageStruct) {
	delete(stage.MinLengths, minlength)
	delete(stage.MinLengths_mapString, minlength.Name)
}

// commit minlength to the back repo (if it is already staged)
func (minlength *MinLength) Commit(stage *StageStruct) *MinLength {
	if _, ok := stage.MinLengths[minlength]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitMinLength(minlength)
		}
	}
	return minlength
}

func (minlength *MinLength) CommitVoid(stage *StageStruct) {
	minlength.Commit(stage)
}

// Checkout minlength to the back repo (if it is already staged)
func (minlength *MinLength) Checkout(stage *StageStruct) *MinLength {
	if _, ok := stage.MinLengths[minlength]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutMinLength(minlength)
		}
	}
	return minlength
}

// for satisfaction of GongStruct interface
func (minlength *MinLength) GetName() (res string) {
	return minlength.Name
}

// Stage puts pattern to the model stage
func (pattern *Pattern) Stage(stage *StageStruct) *Pattern {
	stage.Patterns[pattern] = __member
	stage.Patterns_mapString[pattern.Name] = pattern

	return pattern
}

// Unstage removes pattern off the model stage
func (pattern *Pattern) Unstage(stage *StageStruct) *Pattern {
	delete(stage.Patterns, pattern)
	delete(stage.Patterns_mapString, pattern.Name)
	return pattern
}

// UnstageVoid removes pattern off the model stage
func (pattern *Pattern) UnstageVoid(stage *StageStruct) {
	delete(stage.Patterns, pattern)
	delete(stage.Patterns_mapString, pattern.Name)
}

// commit pattern to the back repo (if it is already staged)
func (pattern *Pattern) Commit(stage *StageStruct) *Pattern {
	if _, ok := stage.Patterns[pattern]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitPattern(pattern)
		}
	}
	return pattern
}

func (pattern *Pattern) CommitVoid(stage *StageStruct) {
	pattern.Commit(stage)
}

// Checkout pattern to the back repo (if it is already staged)
func (pattern *Pattern) Checkout(stage *StageStruct) *Pattern {
	if _, ok := stage.Patterns[pattern]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutPattern(pattern)
		}
	}
	return pattern
}

// for satisfaction of GongStruct interface
func (pattern *Pattern) GetName() (res string) {
	return pattern.Name
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

// Stage puts whitespace to the model stage
func (whitespace *WhiteSpace) Stage(stage *StageStruct) *WhiteSpace {
	stage.WhiteSpaces[whitespace] = __member
	stage.WhiteSpaces_mapString[whitespace.Name] = whitespace

	return whitespace
}

// Unstage removes whitespace off the model stage
func (whitespace *WhiteSpace) Unstage(stage *StageStruct) *WhiteSpace {
	delete(stage.WhiteSpaces, whitespace)
	delete(stage.WhiteSpaces_mapString, whitespace.Name)
	return whitespace
}

// UnstageVoid removes whitespace off the model stage
func (whitespace *WhiteSpace) UnstageVoid(stage *StageStruct) {
	delete(stage.WhiteSpaces, whitespace)
	delete(stage.WhiteSpaces_mapString, whitespace.Name)
}

// commit whitespace to the back repo (if it is already staged)
func (whitespace *WhiteSpace) Commit(stage *StageStruct) *WhiteSpace {
	if _, ok := stage.WhiteSpaces[whitespace]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitWhiteSpace(whitespace)
		}
	}
	return whitespace
}

func (whitespace *WhiteSpace) CommitVoid(stage *StageStruct) {
	whitespace.Commit(stage)
}

// Checkout whitespace to the back repo (if it is already staged)
func (whitespace *WhiteSpace) Checkout(stage *StageStruct) *WhiteSpace {
	if _, ok := stage.WhiteSpaces[whitespace]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutWhiteSpace(whitespace)
		}
	}
	return whitespace
}

// for satisfaction of GongStruct interface
func (whitespace *WhiteSpace) GetName() (res string) {
	return whitespace.Name
}

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMAnnotation(Annotation *Annotation)
	CreateORMComplexType(ComplexType *ComplexType)
	CreateORMDocumentation(Documentation *Documentation)
	CreateORMElement(Element *Element)
	CreateORMEnumeration(Enumeration *Enumeration)
	CreateORMLength(Length *Length)
	CreateORMMaxInclusive(MaxInclusive *MaxInclusive)
	CreateORMMaxLength(MaxLength *MaxLength)
	CreateORMMinInclusive(MinInclusive *MinInclusive)
	CreateORMMinLength(MinLength *MinLength)
	CreateORMPattern(Pattern *Pattern)
	CreateORMRestriction(Restriction *Restriction)
	CreateORMSchema(Schema *Schema)
	CreateORMSequence(Sequence *Sequence)
	CreateORMSimpleType(SimpleType *SimpleType)
	CreateORMWhiteSpace(WhiteSpace *WhiteSpace)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMAnnotation(Annotation *Annotation)
	DeleteORMComplexType(ComplexType *ComplexType)
	DeleteORMDocumentation(Documentation *Documentation)
	DeleteORMElement(Element *Element)
	DeleteORMEnumeration(Enumeration *Enumeration)
	DeleteORMLength(Length *Length)
	DeleteORMMaxInclusive(MaxInclusive *MaxInclusive)
	DeleteORMMaxLength(MaxLength *MaxLength)
	DeleteORMMinInclusive(MinInclusive *MinInclusive)
	DeleteORMMinLength(MinLength *MinLength)
	DeleteORMPattern(Pattern *Pattern)
	DeleteORMRestriction(Restriction *Restriction)
	DeleteORMSchema(Schema *Schema)
	DeleteORMSequence(Sequence *Sequence)
	DeleteORMSimpleType(SimpleType *SimpleType)
	DeleteORMWhiteSpace(WhiteSpace *WhiteSpace)
}

func (stage *StageStruct) Reset() { // insertion point for array reset
	stage.Annotations = make(map[*Annotation]any)
	stage.Annotations_mapString = make(map[string]*Annotation)

	stage.ComplexTypes = make(map[*ComplexType]any)
	stage.ComplexTypes_mapString = make(map[string]*ComplexType)

	stage.Documentations = make(map[*Documentation]any)
	stage.Documentations_mapString = make(map[string]*Documentation)

	stage.Elements = make(map[*Element]any)
	stage.Elements_mapString = make(map[string]*Element)

	stage.Enumerations = make(map[*Enumeration]any)
	stage.Enumerations_mapString = make(map[string]*Enumeration)

	stage.Lengths = make(map[*Length]any)
	stage.Lengths_mapString = make(map[string]*Length)

	stage.MaxInclusives = make(map[*MaxInclusive]any)
	stage.MaxInclusives_mapString = make(map[string]*MaxInclusive)

	stage.MaxLengths = make(map[*MaxLength]any)
	stage.MaxLengths_mapString = make(map[string]*MaxLength)

	stage.MinInclusives = make(map[*MinInclusive]any)
	stage.MinInclusives_mapString = make(map[string]*MinInclusive)

	stage.MinLengths = make(map[*MinLength]any)
	stage.MinLengths_mapString = make(map[string]*MinLength)

	stage.Patterns = make(map[*Pattern]any)
	stage.Patterns_mapString = make(map[string]*Pattern)

	stage.Restrictions = make(map[*Restriction]any)
	stage.Restrictions_mapString = make(map[string]*Restriction)

	stage.Schemas = make(map[*Schema]any)
	stage.Schemas_mapString = make(map[string]*Schema)

	stage.Sequences = make(map[*Sequence]any)
	stage.Sequences_mapString = make(map[string]*Sequence)

	stage.SimpleTypes = make(map[*SimpleType]any)
	stage.SimpleTypes_mapString = make(map[string]*SimpleType)

	stage.WhiteSpaces = make(map[*WhiteSpace]any)
	stage.WhiteSpaces_mapString = make(map[string]*WhiteSpace)

}

func (stage *StageStruct) Nil() { // insertion point for array nil
	stage.Annotations = nil
	stage.Annotations_mapString = nil

	stage.ComplexTypes = nil
	stage.ComplexTypes_mapString = nil

	stage.Documentations = nil
	stage.Documentations_mapString = nil

	stage.Elements = nil
	stage.Elements_mapString = nil

	stage.Enumerations = nil
	stage.Enumerations_mapString = nil

	stage.Lengths = nil
	stage.Lengths_mapString = nil

	stage.MaxInclusives = nil
	stage.MaxInclusives_mapString = nil

	stage.MaxLengths = nil
	stage.MaxLengths_mapString = nil

	stage.MinInclusives = nil
	stage.MinInclusives_mapString = nil

	stage.MinLengths = nil
	stage.MinLengths_mapString = nil

	stage.Patterns = nil
	stage.Patterns_mapString = nil

	stage.Restrictions = nil
	stage.Restrictions_mapString = nil

	stage.Schemas = nil
	stage.Schemas_mapString = nil

	stage.Sequences = nil
	stage.Sequences_mapString = nil

	stage.SimpleTypes = nil
	stage.SimpleTypes_mapString = nil

	stage.WhiteSpaces = nil
	stage.WhiteSpaces_mapString = nil

}

func (stage *StageStruct) Unstage() { // insertion point for array nil
	for annotation := range stage.Annotations {
		annotation.Unstage(stage)
	}

	for complextype := range stage.ComplexTypes {
		complextype.Unstage(stage)
	}

	for documentation := range stage.Documentations {
		documentation.Unstage(stage)
	}

	for element := range stage.Elements {
		element.Unstage(stage)
	}

	for enumeration := range stage.Enumerations {
		enumeration.Unstage(stage)
	}

	for length := range stage.Lengths {
		length.Unstage(stage)
	}

	for maxinclusive := range stage.MaxInclusives {
		maxinclusive.Unstage(stage)
	}

	for maxlength := range stage.MaxLengths {
		maxlength.Unstage(stage)
	}

	for mininclusive := range stage.MinInclusives {
		mininclusive.Unstage(stage)
	}

	for minlength := range stage.MinLengths {
		minlength.Unstage(stage)
	}

	for pattern := range stage.Patterns {
		pattern.Unstage(stage)
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

	for whitespace := range stage.WhiteSpaces {
		whitespace.Unstage(stage)
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
	case map[*ComplexType]any:
		return any(&stage.ComplexTypes).(*Type)
	case map[*Documentation]any:
		return any(&stage.Documentations).(*Type)
	case map[*Element]any:
		return any(&stage.Elements).(*Type)
	case map[*Enumeration]any:
		return any(&stage.Enumerations).(*Type)
	case map[*Length]any:
		return any(&stage.Lengths).(*Type)
	case map[*MaxInclusive]any:
		return any(&stage.MaxInclusives).(*Type)
	case map[*MaxLength]any:
		return any(&stage.MaxLengths).(*Type)
	case map[*MinInclusive]any:
		return any(&stage.MinInclusives).(*Type)
	case map[*MinLength]any:
		return any(&stage.MinLengths).(*Type)
	case map[*Pattern]any:
		return any(&stage.Patterns).(*Type)
	case map[*Restriction]any:
		return any(&stage.Restrictions).(*Type)
	case map[*Schema]any:
		return any(&stage.Schemas).(*Type)
	case map[*Sequence]any:
		return any(&stage.Sequences).(*Type)
	case map[*SimpleType]any:
		return any(&stage.SimpleTypes).(*Type)
	case map[*WhiteSpace]any:
		return any(&stage.WhiteSpaces).(*Type)
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
	case map[string]*ComplexType:
		return any(&stage.ComplexTypes_mapString).(*Type)
	case map[string]*Documentation:
		return any(&stage.Documentations_mapString).(*Type)
	case map[string]*Element:
		return any(&stage.Elements_mapString).(*Type)
	case map[string]*Enumeration:
		return any(&stage.Enumerations_mapString).(*Type)
	case map[string]*Length:
		return any(&stage.Lengths_mapString).(*Type)
	case map[string]*MaxInclusive:
		return any(&stage.MaxInclusives_mapString).(*Type)
	case map[string]*MaxLength:
		return any(&stage.MaxLengths_mapString).(*Type)
	case map[string]*MinInclusive:
		return any(&stage.MinInclusives_mapString).(*Type)
	case map[string]*MinLength:
		return any(&stage.MinLengths_mapString).(*Type)
	case map[string]*Pattern:
		return any(&stage.Patterns_mapString).(*Type)
	case map[string]*Restriction:
		return any(&stage.Restrictions_mapString).(*Type)
	case map[string]*Schema:
		return any(&stage.Schemas_mapString).(*Type)
	case map[string]*Sequence:
		return any(&stage.Sequences_mapString).(*Type)
	case map[string]*SimpleType:
		return any(&stage.SimpleTypes_mapString).(*Type)
	case map[string]*WhiteSpace:
		return any(&stage.WhiteSpaces_mapString).(*Type)
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
	case ComplexType:
		return any(&stage.ComplexTypes).(*map[*Type]any)
	case Documentation:
		return any(&stage.Documentations).(*map[*Type]any)
	case Element:
		return any(&stage.Elements).(*map[*Type]any)
	case Enumeration:
		return any(&stage.Enumerations).(*map[*Type]any)
	case Length:
		return any(&stage.Lengths).(*map[*Type]any)
	case MaxInclusive:
		return any(&stage.MaxInclusives).(*map[*Type]any)
	case MaxLength:
		return any(&stage.MaxLengths).(*map[*Type]any)
	case MinInclusive:
		return any(&stage.MinInclusives).(*map[*Type]any)
	case MinLength:
		return any(&stage.MinLengths).(*map[*Type]any)
	case Pattern:
		return any(&stage.Patterns).(*map[*Type]any)
	case Restriction:
		return any(&stage.Restrictions).(*map[*Type]any)
	case Schema:
		return any(&stage.Schemas).(*map[*Type]any)
	case Sequence:
		return any(&stage.Sequences).(*map[*Type]any)
	case SimpleType:
		return any(&stage.SimpleTypes).(*map[*Type]any)
	case WhiteSpace:
		return any(&stage.WhiteSpaces).(*map[*Type]any)
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
	case *ComplexType:
		return any(&stage.ComplexTypes).(*map[Type]any)
	case *Documentation:
		return any(&stage.Documentations).(*map[Type]any)
	case *Element:
		return any(&stage.Elements).(*map[Type]any)
	case *Enumeration:
		return any(&stage.Enumerations).(*map[Type]any)
	case *Length:
		return any(&stage.Lengths).(*map[Type]any)
	case *MaxInclusive:
		return any(&stage.MaxInclusives).(*map[Type]any)
	case *MaxLength:
		return any(&stage.MaxLengths).(*map[Type]any)
	case *MinInclusive:
		return any(&stage.MinInclusives).(*map[Type]any)
	case *MinLength:
		return any(&stage.MinLengths).(*map[Type]any)
	case *Pattern:
		return any(&stage.Patterns).(*map[Type]any)
	case *Restriction:
		return any(&stage.Restrictions).(*map[Type]any)
	case *Schema:
		return any(&stage.Schemas).(*map[Type]any)
	case *Sequence:
		return any(&stage.Sequences).(*map[Type]any)
	case *SimpleType:
		return any(&stage.SimpleTypes).(*map[Type]any)
	case *WhiteSpace:
		return any(&stage.WhiteSpaces).(*map[Type]any)
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
	case ComplexType:
		return any(&stage.ComplexTypes_mapString).(*map[string]*Type)
	case Documentation:
		return any(&stage.Documentations_mapString).(*map[string]*Type)
	case Element:
		return any(&stage.Elements_mapString).(*map[string]*Type)
	case Enumeration:
		return any(&stage.Enumerations_mapString).(*map[string]*Type)
	case Length:
		return any(&stage.Lengths_mapString).(*map[string]*Type)
	case MaxInclusive:
		return any(&stage.MaxInclusives_mapString).(*map[string]*Type)
	case MaxLength:
		return any(&stage.MaxLengths_mapString).(*map[string]*Type)
	case MinInclusive:
		return any(&stage.MinInclusives_mapString).(*map[string]*Type)
	case MinLength:
		return any(&stage.MinLengths_mapString).(*map[string]*Type)
	case Pattern:
		return any(&stage.Patterns_mapString).(*map[string]*Type)
	case Restriction:
		return any(&stage.Restrictions_mapString).(*map[string]*Type)
	case Schema:
		return any(&stage.Schemas_mapString).(*map[string]*Type)
	case Sequence:
		return any(&stage.Sequences_mapString).(*map[string]*Type)
	case SimpleType:
		return any(&stage.SimpleTypes_mapString).(*map[string]*Type)
	case WhiteSpace:
		return any(&stage.WhiteSpaces_mapString).(*map[string]*Type)
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
	case ComplexType:
		return any(&ComplexType{
			// Initialisation of associations
			// field is initialized with an instance of Sequence with the name of the field
			Sequence: &Sequence{Name: "Sequence"},
			// field is initialized with ElementWithAnnotation as it is a composite
			ElementWithAnnotation: ElementWithAnnotation{
				// per field init
				//
				Annotation: &Annotation{Name: "Annotation"},
			},
		}).(*Type)
	case Documentation:
		return any(&Documentation{
			// Initialisation of associations
		}).(*Type)
	case Element:
		return any(&Element{
			// Initialisation of associations
			// field is initialized with an instance of SimpleType with the name of the field
			SimpleType: &SimpleType{Name: "SimpleType"},
			// field is initialized with an instance of ComplexType with the name of the field
			ComplexType: &ComplexType{Name: "ComplexType"},
			// field is initialized with ElementWithAnnotation as it is a composite
			ElementWithAnnotation: ElementWithAnnotation{
				// per field init
				//
				Annotation: &Annotation{Name: "Annotation"},
			},
		}).(*Type)
	case Enumeration:
		return any(&Enumeration{
			// Initialisation of associations
			// field is initialized with ElementWithAnnotation as it is a composite
			ElementWithAnnotation: ElementWithAnnotation{
				// per field init
				//
				Annotation: &Annotation{Name: "Annotation"},
			},
		}).(*Type)
	case Length:
		return any(&Length{
			// Initialisation of associations
			// field is initialized with ElementWithAnnotation as it is a composite
			ElementWithAnnotation: ElementWithAnnotation{
				// per field init
				//
				Annotation: &Annotation{Name: "Annotation"},
			},
		}).(*Type)
	case MaxInclusive:
		return any(&MaxInclusive{
			// Initialisation of associations
			// field is initialized with ElementWithAnnotation as it is a composite
			ElementWithAnnotation: ElementWithAnnotation{
				// per field init
				//
				Annotation: &Annotation{Name: "Annotation"},
			},
		}).(*Type)
	case MaxLength:
		return any(&MaxLength{
			// Initialisation of associations
			// field is initialized with ElementWithAnnotation as it is a composite
			ElementWithAnnotation: ElementWithAnnotation{
				// per field init
				//
				Annotation: &Annotation{Name: "Annotation"},
			},
		}).(*Type)
	case MinInclusive:
		return any(&MinInclusive{
			// Initialisation of associations
			// field is initialized with ElementWithAnnotation as it is a composite
			ElementWithAnnotation: ElementWithAnnotation{
				// per field init
				//
				Annotation: &Annotation{Name: "Annotation"},
			},
		}).(*Type)
	case MinLength:
		return any(&MinLength{
			// Initialisation of associations
			// field is initialized with ElementWithAnnotation as it is a composite
			ElementWithAnnotation: ElementWithAnnotation{
				// per field init
				//
				Annotation: &Annotation{Name: "Annotation"},
			},
		}).(*Type)
	case Pattern:
		return any(&Pattern{
			// Initialisation of associations
			// field is initialized with ElementWithAnnotation as it is a composite
			ElementWithAnnotation: ElementWithAnnotation{
				// per field init
				//
				Annotation: &Annotation{Name: "Annotation"},
			},
		}).(*Type)
	case Restriction:
		return any(&Restriction{
			// Initialisation of associations
			// field is initialized with an instance of Enumeration with the name of the field
			Enumerations: []*Enumeration{{Name: "Enumerations"}},
			// field is initialized with an instance of MinInclusive with the name of the field
			MinInclusive: &MinInclusive{Name: "MinInclusive"},
			// field is initialized with an instance of MaxInclusive with the name of the field
			MaxInclusive: &MaxInclusive{Name: "MaxInclusive"},
			// field is initialized with an instance of Pattern with the name of the field
			Pattern: &Pattern{Name: "Pattern"},
			// field is initialized with an instance of WhiteSpace with the name of the field
			WhiteSpace: &WhiteSpace{Name: "WhiteSpace"},
			// field is initialized with an instance of MinLength with the name of the field
			MinLength: &MinLength{Name: "MinLength"},
			// field is initialized with an instance of MaxLength with the name of the field
			MaxLength: &MaxLength{Name: "MaxLength"},
			// field is initialized with an instance of Length with the name of the field
			Length: &Length{Name: "Length"},
			// field is initialized with ElementWithAnnotation as it is a composite
			ElementWithAnnotation: ElementWithAnnotation{
				// per field init
				//
				Annotation: &Annotation{Name: "Annotation"},
			},
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
			// field is initialized with ElementWithAnnotation as it is a composite
			ElementWithAnnotation: ElementWithAnnotation{
				// per field init
				//
				Annotation: &Annotation{Name: "Annotation"},
			},
		}).(*Type)
	case Sequence:
		return any(&Sequence{
			// Initialisation of associations
			// field is initialized with an instance of Element with the name of the field
			Elements: []*Element{{Name: "Elements"}},
			// field is initialized with ElementWithAnnotation as it is a composite
			ElementWithAnnotation: ElementWithAnnotation{
				// per field init
				//
				Annotation: &Annotation{Name: "Annotation"},
			},
		}).(*Type)
	case SimpleType:
		return any(&SimpleType{
			// Initialisation of associations
			// field is initialized with an instance of Restriction with the name of the field
			Restriction: &Restriction{Name: "Restriction"},
			// field is initialized with ElementWithAnnotation as it is a composite
			ElementWithAnnotation: ElementWithAnnotation{
				// per field init
				//
				Annotation: &Annotation{Name: "Annotation"},
			},
		}).(*Type)
	case WhiteSpace:
		return any(&WhiteSpace{
			// Initialisation of associations
			// field is initialized with ElementWithAnnotation as it is a composite
			ElementWithAnnotation: ElementWithAnnotation{
				// per field init
				//
				Annotation: &Annotation{Name: "Annotation"},
			},
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
	// reverse maps of direct associations of ComplexType
	case ComplexType:
		switch fieldname {
		// insertion point for per direct association field
		case "Annotation":
			res := make(map[*Annotation][]*ComplexType)
			for complextype := range stage.ComplexTypes {
				if complextype.Annotation != nil {
					annotation_ := complextype.Annotation
					var complextypes []*ComplexType
					_, ok := res[annotation_]
					if ok {
						complextypes = res[annotation_]
					} else {
						complextypes = make([]*ComplexType, 0)
					}
					complextypes = append(complextypes, complextype)
					res[annotation_] = complextypes
				}
			}
			return any(res).(map[*End][]*Start)
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
	// reverse maps of direct associations of Documentation
	case Documentation:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Element
	case Element:
		switch fieldname {
		// insertion point for per direct association field
		case "Annotation":
			res := make(map[*Annotation][]*Element)
			for element := range stage.Elements {
				if element.Annotation != nil {
					annotation_ := element.Annotation
					var elements []*Element
					_, ok := res[annotation_]
					if ok {
						elements = res[annotation_]
					} else {
						elements = make([]*Element, 0)
					}
					elements = append(elements, element)
					res[annotation_] = elements
				}
			}
			return any(res).(map[*End][]*Start)
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
		case "Annotation":
			res := make(map[*Annotation][]*Enumeration)
			for enumeration := range stage.Enumerations {
				if enumeration.Annotation != nil {
					annotation_ := enumeration.Annotation
					var enumerations []*Enumeration
					_, ok := res[annotation_]
					if ok {
						enumerations = res[annotation_]
					} else {
						enumerations = make([]*Enumeration, 0)
					}
					enumerations = append(enumerations, enumeration)
					res[annotation_] = enumerations
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Length
	case Length:
		switch fieldname {
		// insertion point for per direct association field
		case "Annotation":
			res := make(map[*Annotation][]*Length)
			for length := range stage.Lengths {
				if length.Annotation != nil {
					annotation_ := length.Annotation
					var lengths []*Length
					_, ok := res[annotation_]
					if ok {
						lengths = res[annotation_]
					} else {
						lengths = make([]*Length, 0)
					}
					lengths = append(lengths, length)
					res[annotation_] = lengths
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of MaxInclusive
	case MaxInclusive:
		switch fieldname {
		// insertion point for per direct association field
		case "Annotation":
			res := make(map[*Annotation][]*MaxInclusive)
			for maxinclusive := range stage.MaxInclusives {
				if maxinclusive.Annotation != nil {
					annotation_ := maxinclusive.Annotation
					var maxinclusives []*MaxInclusive
					_, ok := res[annotation_]
					if ok {
						maxinclusives = res[annotation_]
					} else {
						maxinclusives = make([]*MaxInclusive, 0)
					}
					maxinclusives = append(maxinclusives, maxinclusive)
					res[annotation_] = maxinclusives
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of MaxLength
	case MaxLength:
		switch fieldname {
		// insertion point for per direct association field
		case "Annotation":
			res := make(map[*Annotation][]*MaxLength)
			for maxlength := range stage.MaxLengths {
				if maxlength.Annotation != nil {
					annotation_ := maxlength.Annotation
					var maxlengths []*MaxLength
					_, ok := res[annotation_]
					if ok {
						maxlengths = res[annotation_]
					} else {
						maxlengths = make([]*MaxLength, 0)
					}
					maxlengths = append(maxlengths, maxlength)
					res[annotation_] = maxlengths
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of MinInclusive
	case MinInclusive:
		switch fieldname {
		// insertion point for per direct association field
		case "Annotation":
			res := make(map[*Annotation][]*MinInclusive)
			for mininclusive := range stage.MinInclusives {
				if mininclusive.Annotation != nil {
					annotation_ := mininclusive.Annotation
					var mininclusives []*MinInclusive
					_, ok := res[annotation_]
					if ok {
						mininclusives = res[annotation_]
					} else {
						mininclusives = make([]*MinInclusive, 0)
					}
					mininclusives = append(mininclusives, mininclusive)
					res[annotation_] = mininclusives
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of MinLength
	case MinLength:
		switch fieldname {
		// insertion point for per direct association field
		case "Annotation":
			res := make(map[*Annotation][]*MinLength)
			for minlength := range stage.MinLengths {
				if minlength.Annotation != nil {
					annotation_ := minlength.Annotation
					var minlengths []*MinLength
					_, ok := res[annotation_]
					if ok {
						minlengths = res[annotation_]
					} else {
						minlengths = make([]*MinLength, 0)
					}
					minlengths = append(minlengths, minlength)
					res[annotation_] = minlengths
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Pattern
	case Pattern:
		switch fieldname {
		// insertion point for per direct association field
		case "Annotation":
			res := make(map[*Annotation][]*Pattern)
			for pattern := range stage.Patterns {
				if pattern.Annotation != nil {
					annotation_ := pattern.Annotation
					var patterns []*Pattern
					_, ok := res[annotation_]
					if ok {
						patterns = res[annotation_]
					} else {
						patterns = make([]*Pattern, 0)
					}
					patterns = append(patterns, pattern)
					res[annotation_] = patterns
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Restriction
	case Restriction:
		switch fieldname {
		// insertion point for per direct association field
		case "Annotation":
			res := make(map[*Annotation][]*Restriction)
			for restriction := range stage.Restrictions {
				if restriction.Annotation != nil {
					annotation_ := restriction.Annotation
					var restrictions []*Restriction
					_, ok := res[annotation_]
					if ok {
						restrictions = res[annotation_]
					} else {
						restrictions = make([]*Restriction, 0)
					}
					restrictions = append(restrictions, restriction)
					res[annotation_] = restrictions
				}
			}
			return any(res).(map[*End][]*Start)
		case "MinInclusive":
			res := make(map[*MinInclusive][]*Restriction)
			for restriction := range stage.Restrictions {
				if restriction.MinInclusive != nil {
					mininclusive_ := restriction.MinInclusive
					var restrictions []*Restriction
					_, ok := res[mininclusive_]
					if ok {
						restrictions = res[mininclusive_]
					} else {
						restrictions = make([]*Restriction, 0)
					}
					restrictions = append(restrictions, restriction)
					res[mininclusive_] = restrictions
				}
			}
			return any(res).(map[*End][]*Start)
		case "MaxInclusive":
			res := make(map[*MaxInclusive][]*Restriction)
			for restriction := range stage.Restrictions {
				if restriction.MaxInclusive != nil {
					maxinclusive_ := restriction.MaxInclusive
					var restrictions []*Restriction
					_, ok := res[maxinclusive_]
					if ok {
						restrictions = res[maxinclusive_]
					} else {
						restrictions = make([]*Restriction, 0)
					}
					restrictions = append(restrictions, restriction)
					res[maxinclusive_] = restrictions
				}
			}
			return any(res).(map[*End][]*Start)
		case "Pattern":
			res := make(map[*Pattern][]*Restriction)
			for restriction := range stage.Restrictions {
				if restriction.Pattern != nil {
					pattern_ := restriction.Pattern
					var restrictions []*Restriction
					_, ok := res[pattern_]
					if ok {
						restrictions = res[pattern_]
					} else {
						restrictions = make([]*Restriction, 0)
					}
					restrictions = append(restrictions, restriction)
					res[pattern_] = restrictions
				}
			}
			return any(res).(map[*End][]*Start)
		case "WhiteSpace":
			res := make(map[*WhiteSpace][]*Restriction)
			for restriction := range stage.Restrictions {
				if restriction.WhiteSpace != nil {
					whitespace_ := restriction.WhiteSpace
					var restrictions []*Restriction
					_, ok := res[whitespace_]
					if ok {
						restrictions = res[whitespace_]
					} else {
						restrictions = make([]*Restriction, 0)
					}
					restrictions = append(restrictions, restriction)
					res[whitespace_] = restrictions
				}
			}
			return any(res).(map[*End][]*Start)
		case "MinLength":
			res := make(map[*MinLength][]*Restriction)
			for restriction := range stage.Restrictions {
				if restriction.MinLength != nil {
					minlength_ := restriction.MinLength
					var restrictions []*Restriction
					_, ok := res[minlength_]
					if ok {
						restrictions = res[minlength_]
					} else {
						restrictions = make([]*Restriction, 0)
					}
					restrictions = append(restrictions, restriction)
					res[minlength_] = restrictions
				}
			}
			return any(res).(map[*End][]*Start)
		case "MaxLength":
			res := make(map[*MaxLength][]*Restriction)
			for restriction := range stage.Restrictions {
				if restriction.MaxLength != nil {
					maxlength_ := restriction.MaxLength
					var restrictions []*Restriction
					_, ok := res[maxlength_]
					if ok {
						restrictions = res[maxlength_]
					} else {
						restrictions = make([]*Restriction, 0)
					}
					restrictions = append(restrictions, restriction)
					res[maxlength_] = restrictions
				}
			}
			return any(res).(map[*End][]*Start)
		case "Length":
			res := make(map[*Length][]*Restriction)
			for restriction := range stage.Restrictions {
				if restriction.Length != nil {
					length_ := restriction.Length
					var restrictions []*Restriction
					_, ok := res[length_]
					if ok {
						restrictions = res[length_]
					} else {
						restrictions = make([]*Restriction, 0)
					}
					restrictions = append(restrictions, restriction)
					res[length_] = restrictions
				}
			}
			return any(res).(map[*End][]*Start)
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
		}
	// reverse maps of direct associations of Sequence
	case Sequence:
		switch fieldname {
		// insertion point for per direct association field
		case "Annotation":
			res := make(map[*Annotation][]*Sequence)
			for sequence := range stage.Sequences {
				if sequence.Annotation != nil {
					annotation_ := sequence.Annotation
					var sequences []*Sequence
					_, ok := res[annotation_]
					if ok {
						sequences = res[annotation_]
					} else {
						sequences = make([]*Sequence, 0)
					}
					sequences = append(sequences, sequence)
					res[annotation_] = sequences
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of SimpleType
	case SimpleType:
		switch fieldname {
		// insertion point for per direct association field
		case "Annotation":
			res := make(map[*Annotation][]*SimpleType)
			for simpletype := range stage.SimpleTypes {
				if simpletype.Annotation != nil {
					annotation_ := simpletype.Annotation
					var simpletypes []*SimpleType
					_, ok := res[annotation_]
					if ok {
						simpletypes = res[annotation_]
					} else {
						simpletypes = make([]*SimpleType, 0)
					}
					simpletypes = append(simpletypes, simpletype)
					res[annotation_] = simpletypes
				}
			}
			return any(res).(map[*End][]*Start)
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
	// reverse maps of direct associations of WhiteSpace
	case WhiteSpace:
		switch fieldname {
		// insertion point for per direct association field
		case "Annotation":
			res := make(map[*Annotation][]*WhiteSpace)
			for whitespace := range stage.WhiteSpaces {
				if whitespace.Annotation != nil {
					annotation_ := whitespace.Annotation
					var whitespaces []*WhiteSpace
					_, ok := res[annotation_]
					if ok {
						whitespaces = res[annotation_]
					} else {
						whitespaces = make([]*WhiteSpace, 0)
					}
					whitespaces = append(whitespaces, whitespace)
					res[annotation_] = whitespaces
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
	// reverse maps of direct associations of Length
	case Length:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MaxInclusive
	case MaxInclusive:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MaxLength
	case MaxLength:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MinInclusive
	case MinInclusive:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of MinLength
	case MinLength:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Pattern
	case Pattern:
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
	// reverse maps of direct associations of WhiteSpace
	case WhiteSpace:
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
	case ComplexType:
		res = "ComplexType"
	case Documentation:
		res = "Documentation"
	case Element:
		res = "Element"
	case Enumeration:
		res = "Enumeration"
	case Length:
		res = "Length"
	case MaxInclusive:
		res = "MaxInclusive"
	case MaxLength:
		res = "MaxLength"
	case MinInclusive:
		res = "MinInclusive"
	case MinLength:
		res = "MinLength"
	case Pattern:
		res = "Pattern"
	case Restriction:
		res = "Restriction"
	case Schema:
		res = "Schema"
	case Sequence:
		res = "Sequence"
	case SimpleType:
		res = "SimpleType"
	case WhiteSpace:
		res = "WhiteSpace"
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
	case *ComplexType:
		res = "ComplexType"
	case *Documentation:
		res = "Documentation"
	case *Element:
		res = "Element"
	case *Enumeration:
		res = "Enumeration"
	case *Length:
		res = "Length"
	case *MaxInclusive:
		res = "MaxInclusive"
	case *MaxLength:
		res = "MaxLength"
	case *MinInclusive:
		res = "MinInclusive"
	case *MinLength:
		res = "MinLength"
	case *Pattern:
		res = "Pattern"
	case *Restriction:
		res = "Restriction"
	case *Schema:
		res = "Schema"
	case *Sequence:
		res = "Sequence"
	case *SimpleType:
		res = "SimpleType"
	case *WhiteSpace:
		res = "WhiteSpace"
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
	case ComplexType:
		res = []string{"Name", "Annotation", "NameXSD", "Sequence"}
	case Documentation:
		res = []string{"Name", "Source", "Lang"}
	case Element:
		res = []string{"Name", "Annotation", "NameXSD", "Type", "SimpleType", "ComplexType"}
	case Enumeration:
		res = []string{"Name", "Annotation", "Value"}
	case Length:
		res = []string{"Name", "Annotation", "Value"}
	case MaxInclusive:
		res = []string{"Name", "Annotation", "Value"}
	case MaxLength:
		res = []string{"Name", "Annotation", "Value"}
	case MinInclusive:
		res = []string{"Name", "Annotation", "Value"}
	case MinLength:
		res = []string{"Name", "Annotation", "Value"}
	case Pattern:
		res = []string{"Name", "Annotation", "Value"}
	case Restriction:
		res = []string{"Name", "Annotation", "Base", "Enumerations", "MinInclusive", "MaxInclusive", "Pattern", "WhiteSpace", "MinLength", "MaxLength", "Length"}
	case Schema:
		res = []string{"Name", "Annotation", "Elements", "SimpleTypes", "ComplexTypes"}
	case Sequence:
		res = []string{"Name", "Annotation", "Elements"}
	case SimpleType:
		res = []string{"Name", "Annotation", "NameXSD", "Restriction"}
	case WhiteSpace:
		res = []string{"Name", "Annotation", "Value"}
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
	case ComplexType:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Schema"
		rf.Fieldname = "ComplexTypes"
		res = append(res, rf)
	case Documentation:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Annotation"
		rf.Fieldname = "Documentations"
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
	case Length:
		var rf ReverseField
		_ = rf
	case MaxInclusive:
		var rf ReverseField
		_ = rf
	case MaxLength:
		var rf ReverseField
		_ = rf
	case MinInclusive:
		var rf ReverseField
		_ = rf
	case MinLength:
		var rf ReverseField
		_ = rf
	case Pattern:
		var rf ReverseField
		_ = rf
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
	case WhiteSpace:
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
	case *ComplexType:
		res = []string{"Name", "Annotation", "NameXSD", "Sequence"}
	case *Documentation:
		res = []string{"Name", "Source", "Lang"}
	case *Element:
		res = []string{"Name", "Annotation", "NameXSD", "Type", "SimpleType", "ComplexType"}
	case *Enumeration:
		res = []string{"Name", "Annotation", "Value"}
	case *Length:
		res = []string{"Name", "Annotation", "Value"}
	case *MaxInclusive:
		res = []string{"Name", "Annotation", "Value"}
	case *MaxLength:
		res = []string{"Name", "Annotation", "Value"}
	case *MinInclusive:
		res = []string{"Name", "Annotation", "Value"}
	case *MinLength:
		res = []string{"Name", "Annotation", "Value"}
	case *Pattern:
		res = []string{"Name", "Annotation", "Value"}
	case *Restriction:
		res = []string{"Name", "Annotation", "Base", "Enumerations", "MinInclusive", "MaxInclusive", "Pattern", "WhiteSpace", "MinLength", "MaxLength", "Length"}
	case *Schema:
		res = []string{"Name", "Annotation", "Elements", "SimpleTypes", "ComplexTypes"}
	case *Sequence:
		res = []string{"Name", "Annotation", "Elements"}
	case *SimpleType:
		res = []string{"Name", "Annotation", "NameXSD", "Restriction"}
	case *WhiteSpace:
		res = []string{"Name", "Annotation", "Value"}
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
	case *ComplexType:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Annotation":
			if inferedInstance.Annotation != nil {
				res = inferedInstance.Annotation.Name
			}
		case "NameXSD":
			res = inferedInstance.NameXSD
		case "Sequence":
			if inferedInstance.Sequence != nil {
				res = inferedInstance.Sequence.Name
			}
		}
	case *Documentation:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Source":
			res = inferedInstance.Source
		case "Lang":
			res = inferedInstance.Lang
		}
	case *Element:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Annotation":
			if inferedInstance.Annotation != nil {
				res = inferedInstance.Annotation.Name
			}
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
		case "Annotation":
			if inferedInstance.Annotation != nil {
				res = inferedInstance.Annotation.Name
			}
		case "Value":
			res = inferedInstance.Value
		}
	case *Length:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Annotation":
			if inferedInstance.Annotation != nil {
				res = inferedInstance.Annotation.Name
			}
		case "Value":
			res = inferedInstance.Value
		}
	case *MaxInclusive:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Annotation":
			if inferedInstance.Annotation != nil {
				res = inferedInstance.Annotation.Name
			}
		case "Value":
			res = inferedInstance.Value
		}
	case *MaxLength:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Annotation":
			if inferedInstance.Annotation != nil {
				res = inferedInstance.Annotation.Name
			}
		case "Value":
			res = inferedInstance.Value
		}
	case *MinInclusive:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Annotation":
			if inferedInstance.Annotation != nil {
				res = inferedInstance.Annotation.Name
			}
		case "Value":
			res = inferedInstance.Value
		}
	case *MinLength:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Annotation":
			if inferedInstance.Annotation != nil {
				res = inferedInstance.Annotation.Name
			}
		case "Value":
			res = inferedInstance.Value
		}
	case *Pattern:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Annotation":
			if inferedInstance.Annotation != nil {
				res = inferedInstance.Annotation.Name
			}
		case "Value":
			res = inferedInstance.Value
		}
	case *Restriction:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Annotation":
			if inferedInstance.Annotation != nil {
				res = inferedInstance.Annotation.Name
			}
		case "Base":
			res = inferedInstance.Base
		case "Enumerations":
			for idx, __instance__ := range inferedInstance.Enumerations {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "MinInclusive":
			if inferedInstance.MinInclusive != nil {
				res = inferedInstance.MinInclusive.Name
			}
		case "MaxInclusive":
			if inferedInstance.MaxInclusive != nil {
				res = inferedInstance.MaxInclusive.Name
			}
		case "Pattern":
			if inferedInstance.Pattern != nil {
				res = inferedInstance.Pattern.Name
			}
		case "WhiteSpace":
			if inferedInstance.WhiteSpace != nil {
				res = inferedInstance.WhiteSpace.Name
			}
		case "MinLength":
			if inferedInstance.MinLength != nil {
				res = inferedInstance.MinLength.Name
			}
		case "MaxLength":
			if inferedInstance.MaxLength != nil {
				res = inferedInstance.MaxLength.Name
			}
		case "Length":
			if inferedInstance.Length != nil {
				res = inferedInstance.Length.Name
			}
		}
	case *Schema:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Annotation":
			if inferedInstance.Annotation != nil {
				res = inferedInstance.Annotation.Name
			}
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
		case "Annotation":
			if inferedInstance.Annotation != nil {
				res = inferedInstance.Annotation.Name
			}
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
		case "Annotation":
			if inferedInstance.Annotation != nil {
				res = inferedInstance.Annotation.Name
			}
		case "NameXSD":
			res = inferedInstance.NameXSD
		case "Restriction":
			if inferedInstance.Restriction != nil {
				res = inferedInstance.Restriction.Name
			}
		}
	case *WhiteSpace:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Annotation":
			if inferedInstance.Annotation != nil {
				res = inferedInstance.Annotation.Name
			}
		case "Value":
			res = inferedInstance.Value
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
	case ComplexType:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Annotation":
			if inferedInstance.Annotation != nil {
				res = inferedInstance.Annotation.Name
			}
		case "NameXSD":
			res = inferedInstance.NameXSD
		case "Sequence":
			if inferedInstance.Sequence != nil {
				res = inferedInstance.Sequence.Name
			}
		}
	case Documentation:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Source":
			res = inferedInstance.Source
		case "Lang":
			res = inferedInstance.Lang
		}
	case Element:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Annotation":
			if inferedInstance.Annotation != nil {
				res = inferedInstance.Annotation.Name
			}
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
		case "Annotation":
			if inferedInstance.Annotation != nil {
				res = inferedInstance.Annotation.Name
			}
		case "Value":
			res = inferedInstance.Value
		}
	case Length:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Annotation":
			if inferedInstance.Annotation != nil {
				res = inferedInstance.Annotation.Name
			}
		case "Value":
			res = inferedInstance.Value
		}
	case MaxInclusive:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Annotation":
			if inferedInstance.Annotation != nil {
				res = inferedInstance.Annotation.Name
			}
		case "Value":
			res = inferedInstance.Value
		}
	case MaxLength:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Annotation":
			if inferedInstance.Annotation != nil {
				res = inferedInstance.Annotation.Name
			}
		case "Value":
			res = inferedInstance.Value
		}
	case MinInclusive:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Annotation":
			if inferedInstance.Annotation != nil {
				res = inferedInstance.Annotation.Name
			}
		case "Value":
			res = inferedInstance.Value
		}
	case MinLength:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Annotation":
			if inferedInstance.Annotation != nil {
				res = inferedInstance.Annotation.Name
			}
		case "Value":
			res = inferedInstance.Value
		}
	case Pattern:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Annotation":
			if inferedInstance.Annotation != nil {
				res = inferedInstance.Annotation.Name
			}
		case "Value":
			res = inferedInstance.Value
		}
	case Restriction:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Annotation":
			if inferedInstance.Annotation != nil {
				res = inferedInstance.Annotation.Name
			}
		case "Base":
			res = inferedInstance.Base
		case "Enumerations":
			for idx, __instance__ := range inferedInstance.Enumerations {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "MinInclusive":
			if inferedInstance.MinInclusive != nil {
				res = inferedInstance.MinInclusive.Name
			}
		case "MaxInclusive":
			if inferedInstance.MaxInclusive != nil {
				res = inferedInstance.MaxInclusive.Name
			}
		case "Pattern":
			if inferedInstance.Pattern != nil {
				res = inferedInstance.Pattern.Name
			}
		case "WhiteSpace":
			if inferedInstance.WhiteSpace != nil {
				res = inferedInstance.WhiteSpace.Name
			}
		case "MinLength":
			if inferedInstance.MinLength != nil {
				res = inferedInstance.MinLength.Name
			}
		case "MaxLength":
			if inferedInstance.MaxLength != nil {
				res = inferedInstance.MaxLength.Name
			}
		case "Length":
			if inferedInstance.Length != nil {
				res = inferedInstance.Length.Name
			}
		}
	case Schema:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Annotation":
			if inferedInstance.Annotation != nil {
				res = inferedInstance.Annotation.Name
			}
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
		case "Annotation":
			if inferedInstance.Annotation != nil {
				res = inferedInstance.Annotation.Name
			}
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
		case "Annotation":
			if inferedInstance.Annotation != nil {
				res = inferedInstance.Annotation.Name
			}
		case "NameXSD":
			res = inferedInstance.NameXSD
		case "Restriction":
			if inferedInstance.Restriction != nil {
				res = inferedInstance.Restriction.Name
			}
		}
	case WhiteSpace:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Annotation":
			if inferedInstance.Annotation != nil {
				res = inferedInstance.Annotation.Name
			}
		case "Value":
			res = inferedInstance.Value
		}
	default:
		_ = inferedInstance
	}
	return
}

// Last line of the template
