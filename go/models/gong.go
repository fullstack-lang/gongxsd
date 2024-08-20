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
	Alls           map[*All]any
	Alls_mapString map[string]*All

	// insertion point for slice of pointers maps
	All_Sequences_reverseMap map[*Sequence]*All

	All_Alls_reverseMap map[*All]*All

	All_Choices_reverseMap map[*Choice]*All

	All_Groups_reverseMap map[*Group]*All

	All_Elements_reverseMap map[*Element]*All

	OnAfterAllCreateCallback OnAfterCreateInterface[All]
	OnAfterAllUpdateCallback OnAfterUpdateInterface[All]
	OnAfterAllDeleteCallback OnAfterDeleteInterface[All]
	OnAfterAllReadCallback   OnAfterReadInterface[All]

	Annotations           map[*Annotation]any
	Annotations_mapString map[string]*Annotation

	// insertion point for slice of pointers maps
	Annotation_Documentations_reverseMap map[*Documentation]*Annotation

	OnAfterAnnotationCreateCallback OnAfterCreateInterface[Annotation]
	OnAfterAnnotationUpdateCallback OnAfterUpdateInterface[Annotation]
	OnAfterAnnotationDeleteCallback OnAfterDeleteInterface[Annotation]
	OnAfterAnnotationReadCallback   OnAfterReadInterface[Annotation]

	Attributes           map[*Attribute]any
	Attributes_mapString map[string]*Attribute

	// insertion point for slice of pointers maps
	OnAfterAttributeCreateCallback OnAfterCreateInterface[Attribute]
	OnAfterAttributeUpdateCallback OnAfterUpdateInterface[Attribute]
	OnAfterAttributeDeleteCallback OnAfterDeleteInterface[Attribute]
	OnAfterAttributeReadCallback   OnAfterReadInterface[Attribute]

	AttributeGroups           map[*AttributeGroup]any
	AttributeGroups_mapString map[string]*AttributeGroup

	// insertion point for slice of pointers maps
	AttributeGroup_AttributeGroups_reverseMap map[*AttributeGroup]*AttributeGroup

	AttributeGroup_Attributes_reverseMap map[*Attribute]*AttributeGroup

	OnAfterAttributeGroupCreateCallback OnAfterCreateInterface[AttributeGroup]
	OnAfterAttributeGroupUpdateCallback OnAfterUpdateInterface[AttributeGroup]
	OnAfterAttributeGroupDeleteCallback OnAfterDeleteInterface[AttributeGroup]
	OnAfterAttributeGroupReadCallback   OnAfterReadInterface[AttributeGroup]

	Choices           map[*Choice]any
	Choices_mapString map[string]*Choice

	// insertion point for slice of pointers maps
	Choice_Sequences_reverseMap map[*Sequence]*Choice

	Choice_Alls_reverseMap map[*All]*Choice

	Choice_Choices_reverseMap map[*Choice]*Choice

	Choice_Groups_reverseMap map[*Group]*Choice

	Choice_Elements_reverseMap map[*Element]*Choice

	OnAfterChoiceCreateCallback OnAfterCreateInterface[Choice]
	OnAfterChoiceUpdateCallback OnAfterUpdateInterface[Choice]
	OnAfterChoiceDeleteCallback OnAfterDeleteInterface[Choice]
	OnAfterChoiceReadCallback   OnAfterReadInterface[Choice]

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
	ComplexType_Sequences_reverseMap map[*Sequence]*ComplexType

	ComplexType_Alls_reverseMap map[*All]*ComplexType

	ComplexType_Choices_reverseMap map[*Choice]*ComplexType

	ComplexType_Groups_reverseMap map[*Group]*ComplexType

	ComplexType_Elements_reverseMap map[*Element]*ComplexType

	ComplexType_Attributes_reverseMap map[*Attribute]*ComplexType

	ComplexType_AttributeGroups_reverseMap map[*AttributeGroup]*ComplexType

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
	Element_Groups_reverseMap map[*Group]*Element

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

	Extensions           map[*Extension]any
	Extensions_mapString map[string]*Extension

	// insertion point for slice of pointers maps
	Extension_Sequences_reverseMap map[*Sequence]*Extension

	Extension_Alls_reverseMap map[*All]*Extension

	Extension_Choices_reverseMap map[*Choice]*Extension

	Extension_Groups_reverseMap map[*Group]*Extension

	Extension_Elements_reverseMap map[*Element]*Extension

	Extension_Attributes_reverseMap map[*Attribute]*Extension

	OnAfterExtensionCreateCallback OnAfterCreateInterface[Extension]
	OnAfterExtensionUpdateCallback OnAfterUpdateInterface[Extension]
	OnAfterExtensionDeleteCallback OnAfterDeleteInterface[Extension]
	OnAfterExtensionReadCallback   OnAfterReadInterface[Extension]

	Groups           map[*Group]any
	Groups_mapString map[string]*Group

	// insertion point for slice of pointers maps
	Group_Sequences_reverseMap map[*Sequence]*Group

	Group_Alls_reverseMap map[*All]*Group

	Group_Choices_reverseMap map[*Choice]*Group

	Group_Groups_reverseMap map[*Group]*Group

	Group_Elements_reverseMap map[*Element]*Group

	OnAfterGroupCreateCallback OnAfterCreateInterface[Group]
	OnAfterGroupUpdateCallback OnAfterUpdateInterface[Group]
	OnAfterGroupDeleteCallback OnAfterDeleteInterface[Group]
	OnAfterGroupReadCallback   OnAfterReadInterface[Group]

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

	Schema_AttributeGroups_reverseMap map[*AttributeGroup]*Schema

	Schema_Groups_reverseMap map[*Group]*Schema

	OnAfterSchemaCreateCallback OnAfterCreateInterface[Schema]
	OnAfterSchemaUpdateCallback OnAfterUpdateInterface[Schema]
	OnAfterSchemaDeleteCallback OnAfterDeleteInterface[Schema]
	OnAfterSchemaReadCallback   OnAfterReadInterface[Schema]

	Sequences           map[*Sequence]any
	Sequences_mapString map[string]*Sequence

	// insertion point for slice of pointers maps
	Sequence_Sequences_reverseMap map[*Sequence]*Sequence

	Sequence_Alls_reverseMap map[*All]*Sequence

	Sequence_Choices_reverseMap map[*Choice]*Sequence

	Sequence_Groups_reverseMap map[*Group]*Sequence

	Sequence_Elements_reverseMap map[*Element]*Sequence

	OnAfterSequenceCreateCallback OnAfterCreateInterface[Sequence]
	OnAfterSequenceUpdateCallback OnAfterUpdateInterface[Sequence]
	OnAfterSequenceDeleteCallback OnAfterDeleteInterface[Sequence]
	OnAfterSequenceReadCallback   OnAfterReadInterface[Sequence]

	SimpleContents           map[*SimpleContent]any
	SimpleContents_mapString map[string]*SimpleContent

	// insertion point for slice of pointers maps
	OnAfterSimpleContentCreateCallback OnAfterCreateInterface[SimpleContent]
	OnAfterSimpleContentUpdateCallback OnAfterUpdateInterface[SimpleContent]
	OnAfterSimpleContentDeleteCallback OnAfterDeleteInterface[SimpleContent]
	OnAfterSimpleContentReadCallback   OnAfterReadInterface[SimpleContent]

	SimpleTypes           map[*SimpleType]any
	SimpleTypes_mapString map[string]*SimpleType

	// insertion point for slice of pointers maps
	OnAfterSimpleTypeCreateCallback OnAfterCreateInterface[SimpleType]
	OnAfterSimpleTypeUpdateCallback OnAfterUpdateInterface[SimpleType]
	OnAfterSimpleTypeDeleteCallback OnAfterDeleteInterface[SimpleType]
	OnAfterSimpleTypeReadCallback   OnAfterReadInterface[SimpleType]

	TotalDigits           map[*TotalDigit]any
	TotalDigits_mapString map[string]*TotalDigit

	// insertion point for slice of pointers maps
	OnAfterTotalDigitCreateCallback OnAfterCreateInterface[TotalDigit]
	OnAfterTotalDigitUpdateCallback OnAfterUpdateInterface[TotalDigit]
	OnAfterTotalDigitDeleteCallback OnAfterDeleteInterface[TotalDigit]
	OnAfterTotalDigitReadCallback   OnAfterReadInterface[TotalDigit]

	Unions           map[*Union]any
	Unions_mapString map[string]*Union

	// insertion point for slice of pointers maps
	OnAfterUnionCreateCallback OnAfterCreateInterface[Union]
	OnAfterUnionUpdateCallback OnAfterUpdateInterface[Union]
	OnAfterUnionDeleteCallback OnAfterDeleteInterface[Union]
	OnAfterUnionReadCallback   OnAfterReadInterface[Union]

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
	CommitAll(all *All)
	CheckoutAll(all *All)
	CommitAnnotation(annotation *Annotation)
	CheckoutAnnotation(annotation *Annotation)
	CommitAttribute(attribute *Attribute)
	CheckoutAttribute(attribute *Attribute)
	CommitAttributeGroup(attributegroup *AttributeGroup)
	CheckoutAttributeGroup(attributegroup *AttributeGroup)
	CommitChoice(choice *Choice)
	CheckoutChoice(choice *Choice)
	CommitComplexContent(complexcontent *ComplexContent)
	CheckoutComplexContent(complexcontent *ComplexContent)
	CommitComplexType(complextype *ComplexType)
	CheckoutComplexType(complextype *ComplexType)
	CommitDocumentation(documentation *Documentation)
	CheckoutDocumentation(documentation *Documentation)
	CommitElement(element *Element)
	CheckoutElement(element *Element)
	CommitEnumeration(enumeration *Enumeration)
	CheckoutEnumeration(enumeration *Enumeration)
	CommitExtension(extension *Extension)
	CheckoutExtension(extension *Extension)
	CommitGroup(group *Group)
	CheckoutGroup(group *Group)
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
	CommitSimpleContent(simplecontent *SimpleContent)
	CheckoutSimpleContent(simplecontent *SimpleContent)
	CommitSimpleType(simpletype *SimpleType)
	CheckoutSimpleType(simpletype *SimpleType)
	CommitTotalDigit(totaldigit *TotalDigit)
	CheckoutTotalDigit(totaldigit *TotalDigit)
	CommitUnion(union *Union)
	CheckoutUnion(union *Union)
	CommitWhiteSpace(whitespace *WhiteSpace)
	CheckoutWhiteSpace(whitespace *WhiteSpace)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

func NewStage(path string) (stage *StageStruct) {

	stage = &StageStruct{ // insertion point for array initiatialisation
		Alls:           make(map[*All]any),
		Alls_mapString: make(map[string]*All),

		Annotations:           make(map[*Annotation]any),
		Annotations_mapString: make(map[string]*Annotation),

		Attributes:           make(map[*Attribute]any),
		Attributes_mapString: make(map[string]*Attribute),

		AttributeGroups:           make(map[*AttributeGroup]any),
		AttributeGroups_mapString: make(map[string]*AttributeGroup),

		Choices:           make(map[*Choice]any),
		Choices_mapString: make(map[string]*Choice),

		ComplexContents:           make(map[*ComplexContent]any),
		ComplexContents_mapString: make(map[string]*ComplexContent),

		ComplexTypes:           make(map[*ComplexType]any),
		ComplexTypes_mapString: make(map[string]*ComplexType),

		Documentations:           make(map[*Documentation]any),
		Documentations_mapString: make(map[string]*Documentation),

		Elements:           make(map[*Element]any),
		Elements_mapString: make(map[string]*Element),

		Enumerations:           make(map[*Enumeration]any),
		Enumerations_mapString: make(map[string]*Enumeration),

		Extensions:           make(map[*Extension]any),
		Extensions_mapString: make(map[string]*Extension),

		Groups:           make(map[*Group]any),
		Groups_mapString: make(map[string]*Group),

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

		SimpleContents:           make(map[*SimpleContent]any),
		SimpleContents_mapString: make(map[string]*SimpleContent),

		SimpleTypes:           make(map[*SimpleType]any),
		SimpleTypes_mapString: make(map[string]*SimpleType),

		TotalDigits:           make(map[*TotalDigit]any),
		TotalDigits_mapString: make(map[string]*TotalDigit),

		Unions:           make(map[*Union]any),
		Unions_mapString: make(map[string]*Union),

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
	stage.Map_GongStructName_InstancesNb["All"] = len(stage.Alls)
	stage.Map_GongStructName_InstancesNb["Annotation"] = len(stage.Annotations)
	stage.Map_GongStructName_InstancesNb["Attribute"] = len(stage.Attributes)
	stage.Map_GongStructName_InstancesNb["AttributeGroup"] = len(stage.AttributeGroups)
	stage.Map_GongStructName_InstancesNb["Choice"] = len(stage.Choices)
	stage.Map_GongStructName_InstancesNb["ComplexContent"] = len(stage.ComplexContents)
	stage.Map_GongStructName_InstancesNb["ComplexType"] = len(stage.ComplexTypes)
	stage.Map_GongStructName_InstancesNb["Documentation"] = len(stage.Documentations)
	stage.Map_GongStructName_InstancesNb["Element"] = len(stage.Elements)
	stage.Map_GongStructName_InstancesNb["Enumeration"] = len(stage.Enumerations)
	stage.Map_GongStructName_InstancesNb["Extension"] = len(stage.Extensions)
	stage.Map_GongStructName_InstancesNb["Group"] = len(stage.Groups)
	stage.Map_GongStructName_InstancesNb["Length"] = len(stage.Lengths)
	stage.Map_GongStructName_InstancesNb["MaxInclusive"] = len(stage.MaxInclusives)
	stage.Map_GongStructName_InstancesNb["MaxLength"] = len(stage.MaxLengths)
	stage.Map_GongStructName_InstancesNb["MinInclusive"] = len(stage.MinInclusives)
	stage.Map_GongStructName_InstancesNb["MinLength"] = len(stage.MinLengths)
	stage.Map_GongStructName_InstancesNb["Pattern"] = len(stage.Patterns)
	stage.Map_GongStructName_InstancesNb["Restriction"] = len(stage.Restrictions)
	stage.Map_GongStructName_InstancesNb["Schema"] = len(stage.Schemas)
	stage.Map_GongStructName_InstancesNb["Sequence"] = len(stage.Sequences)
	stage.Map_GongStructName_InstancesNb["SimpleContent"] = len(stage.SimpleContents)
	stage.Map_GongStructName_InstancesNb["SimpleType"] = len(stage.SimpleTypes)
	stage.Map_GongStructName_InstancesNb["TotalDigit"] = len(stage.TotalDigits)
	stage.Map_GongStructName_InstancesNb["Union"] = len(stage.Unions)
	stage.Map_GongStructName_InstancesNb["WhiteSpace"] = len(stage.WhiteSpaces)

}

func (stage *StageStruct) Checkout() {
	if stage.BackRepo != nil {
		stage.BackRepo.Checkout(stage)
	}

	stage.ComputeReverseMaps()
	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["All"] = len(stage.Alls)
	stage.Map_GongStructName_InstancesNb["Annotation"] = len(stage.Annotations)
	stage.Map_GongStructName_InstancesNb["Attribute"] = len(stage.Attributes)
	stage.Map_GongStructName_InstancesNb["AttributeGroup"] = len(stage.AttributeGroups)
	stage.Map_GongStructName_InstancesNb["Choice"] = len(stage.Choices)
	stage.Map_GongStructName_InstancesNb["ComplexContent"] = len(stage.ComplexContents)
	stage.Map_GongStructName_InstancesNb["ComplexType"] = len(stage.ComplexTypes)
	stage.Map_GongStructName_InstancesNb["Documentation"] = len(stage.Documentations)
	stage.Map_GongStructName_InstancesNb["Element"] = len(stage.Elements)
	stage.Map_GongStructName_InstancesNb["Enumeration"] = len(stage.Enumerations)
	stage.Map_GongStructName_InstancesNb["Extension"] = len(stage.Extensions)
	stage.Map_GongStructName_InstancesNb["Group"] = len(stage.Groups)
	stage.Map_GongStructName_InstancesNb["Length"] = len(stage.Lengths)
	stage.Map_GongStructName_InstancesNb["MaxInclusive"] = len(stage.MaxInclusives)
	stage.Map_GongStructName_InstancesNb["MaxLength"] = len(stage.MaxLengths)
	stage.Map_GongStructName_InstancesNb["MinInclusive"] = len(stage.MinInclusives)
	stage.Map_GongStructName_InstancesNb["MinLength"] = len(stage.MinLengths)
	stage.Map_GongStructName_InstancesNb["Pattern"] = len(stage.Patterns)
	stage.Map_GongStructName_InstancesNb["Restriction"] = len(stage.Restrictions)
	stage.Map_GongStructName_InstancesNb["Schema"] = len(stage.Schemas)
	stage.Map_GongStructName_InstancesNb["Sequence"] = len(stage.Sequences)
	stage.Map_GongStructName_InstancesNb["SimpleContent"] = len(stage.SimpleContents)
	stage.Map_GongStructName_InstancesNb["SimpleType"] = len(stage.SimpleTypes)
	stage.Map_GongStructName_InstancesNb["TotalDigit"] = len(stage.TotalDigits)
	stage.Map_GongStructName_InstancesNb["Union"] = len(stage.Unions)
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
// Stage puts all to the model stage
func (all *All) Stage(stage *StageStruct) *All {
	stage.Alls[all] = __member
	stage.Alls_mapString[all.Name] = all

	return all
}

// Unstage removes all off the model stage
func (all *All) Unstage(stage *StageStruct) *All {
	delete(stage.Alls, all)
	delete(stage.Alls_mapString, all.Name)
	return all
}

// UnstageVoid removes all off the model stage
func (all *All) UnstageVoid(stage *StageStruct) {
	delete(stage.Alls, all)
	delete(stage.Alls_mapString, all.Name)
}

// commit all to the back repo (if it is already staged)
func (all *All) Commit(stage *StageStruct) *All {
	if _, ok := stage.Alls[all]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitAll(all)
		}
	}
	return all
}

func (all *All) CommitVoid(stage *StageStruct) {
	all.Commit(stage)
}

// Checkout all to the back repo (if it is already staged)
func (all *All) Checkout(stage *StageStruct) *All {
	if _, ok := stage.Alls[all]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutAll(all)
		}
	}
	return all
}

// for satisfaction of GongStruct interface
func (all *All) GetName() (res string) {
	return all.Name
}

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

// Stage puts attribute to the model stage
func (attribute *Attribute) Stage(stage *StageStruct) *Attribute {
	stage.Attributes[attribute] = __member
	stage.Attributes_mapString[attribute.Name] = attribute

	return attribute
}

// Unstage removes attribute off the model stage
func (attribute *Attribute) Unstage(stage *StageStruct) *Attribute {
	delete(stage.Attributes, attribute)
	delete(stage.Attributes_mapString, attribute.Name)
	return attribute
}

// UnstageVoid removes attribute off the model stage
func (attribute *Attribute) UnstageVoid(stage *StageStruct) {
	delete(stage.Attributes, attribute)
	delete(stage.Attributes_mapString, attribute.Name)
}

// commit attribute to the back repo (if it is already staged)
func (attribute *Attribute) Commit(stage *StageStruct) *Attribute {
	if _, ok := stage.Attributes[attribute]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitAttribute(attribute)
		}
	}
	return attribute
}

func (attribute *Attribute) CommitVoid(stage *StageStruct) {
	attribute.Commit(stage)
}

// Checkout attribute to the back repo (if it is already staged)
func (attribute *Attribute) Checkout(stage *StageStruct) *Attribute {
	if _, ok := stage.Attributes[attribute]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutAttribute(attribute)
		}
	}
	return attribute
}

// for satisfaction of GongStruct interface
func (attribute *Attribute) GetName() (res string) {
	return attribute.Name
}

// Stage puts attributegroup to the model stage
func (attributegroup *AttributeGroup) Stage(stage *StageStruct) *AttributeGroup {
	stage.AttributeGroups[attributegroup] = __member
	stage.AttributeGroups_mapString[attributegroup.Name] = attributegroup

	return attributegroup
}

// Unstage removes attributegroup off the model stage
func (attributegroup *AttributeGroup) Unstage(stage *StageStruct) *AttributeGroup {
	delete(stage.AttributeGroups, attributegroup)
	delete(stage.AttributeGroups_mapString, attributegroup.Name)
	return attributegroup
}

// UnstageVoid removes attributegroup off the model stage
func (attributegroup *AttributeGroup) UnstageVoid(stage *StageStruct) {
	delete(stage.AttributeGroups, attributegroup)
	delete(stage.AttributeGroups_mapString, attributegroup.Name)
}

// commit attributegroup to the back repo (if it is already staged)
func (attributegroup *AttributeGroup) Commit(stage *StageStruct) *AttributeGroup {
	if _, ok := stage.AttributeGroups[attributegroup]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitAttributeGroup(attributegroup)
		}
	}
	return attributegroup
}

func (attributegroup *AttributeGroup) CommitVoid(stage *StageStruct) {
	attributegroup.Commit(stage)
}

// Checkout attributegroup to the back repo (if it is already staged)
func (attributegroup *AttributeGroup) Checkout(stage *StageStruct) *AttributeGroup {
	if _, ok := stage.AttributeGroups[attributegroup]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutAttributeGroup(attributegroup)
		}
	}
	return attributegroup
}

// for satisfaction of GongStruct interface
func (attributegroup *AttributeGroup) GetName() (res string) {
	return attributegroup.Name
}

// Stage puts choice to the model stage
func (choice *Choice) Stage(stage *StageStruct) *Choice {
	stage.Choices[choice] = __member
	stage.Choices_mapString[choice.Name] = choice

	return choice
}

// Unstage removes choice off the model stage
func (choice *Choice) Unstage(stage *StageStruct) *Choice {
	delete(stage.Choices, choice)
	delete(stage.Choices_mapString, choice.Name)
	return choice
}

// UnstageVoid removes choice off the model stage
func (choice *Choice) UnstageVoid(stage *StageStruct) {
	delete(stage.Choices, choice)
	delete(stage.Choices_mapString, choice.Name)
}

// commit choice to the back repo (if it is already staged)
func (choice *Choice) Commit(stage *StageStruct) *Choice {
	if _, ok := stage.Choices[choice]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitChoice(choice)
		}
	}
	return choice
}

func (choice *Choice) CommitVoid(stage *StageStruct) {
	choice.Commit(stage)
}

// Checkout choice to the back repo (if it is already staged)
func (choice *Choice) Checkout(stage *StageStruct) *Choice {
	if _, ok := stage.Choices[choice]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutChoice(choice)
		}
	}
	return choice
}

// for satisfaction of GongStruct interface
func (choice *Choice) GetName() (res string) {
	return choice.Name
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

// Stage puts extension to the model stage
func (extension *Extension) Stage(stage *StageStruct) *Extension {
	stage.Extensions[extension] = __member
	stage.Extensions_mapString[extension.Name] = extension

	return extension
}

// Unstage removes extension off the model stage
func (extension *Extension) Unstage(stage *StageStruct) *Extension {
	delete(stage.Extensions, extension)
	delete(stage.Extensions_mapString, extension.Name)
	return extension
}

// UnstageVoid removes extension off the model stage
func (extension *Extension) UnstageVoid(stage *StageStruct) {
	delete(stage.Extensions, extension)
	delete(stage.Extensions_mapString, extension.Name)
}

// commit extension to the back repo (if it is already staged)
func (extension *Extension) Commit(stage *StageStruct) *Extension {
	if _, ok := stage.Extensions[extension]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitExtension(extension)
		}
	}
	return extension
}

func (extension *Extension) CommitVoid(stage *StageStruct) {
	extension.Commit(stage)
}

// Checkout extension to the back repo (if it is already staged)
func (extension *Extension) Checkout(stage *StageStruct) *Extension {
	if _, ok := stage.Extensions[extension]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutExtension(extension)
		}
	}
	return extension
}

// for satisfaction of GongStruct interface
func (extension *Extension) GetName() (res string) {
	return extension.Name
}

// Stage puts group to the model stage
func (group *Group) Stage(stage *StageStruct) *Group {
	stage.Groups[group] = __member
	stage.Groups_mapString[group.Name] = group

	return group
}

// Unstage removes group off the model stage
func (group *Group) Unstage(stage *StageStruct) *Group {
	delete(stage.Groups, group)
	delete(stage.Groups_mapString, group.Name)
	return group
}

// UnstageVoid removes group off the model stage
func (group *Group) UnstageVoid(stage *StageStruct) {
	delete(stage.Groups, group)
	delete(stage.Groups_mapString, group.Name)
}

// commit group to the back repo (if it is already staged)
func (group *Group) Commit(stage *StageStruct) *Group {
	if _, ok := stage.Groups[group]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitGroup(group)
		}
	}
	return group
}

func (group *Group) CommitVoid(stage *StageStruct) {
	group.Commit(stage)
}

// Checkout group to the back repo (if it is already staged)
func (group *Group) Checkout(stage *StageStruct) *Group {
	if _, ok := stage.Groups[group]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutGroup(group)
		}
	}
	return group
}

// for satisfaction of GongStruct interface
func (group *Group) GetName() (res string) {
	return group.Name
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

// Stage puts simplecontent to the model stage
func (simplecontent *SimpleContent) Stage(stage *StageStruct) *SimpleContent {
	stage.SimpleContents[simplecontent] = __member
	stage.SimpleContents_mapString[simplecontent.Name] = simplecontent

	return simplecontent
}

// Unstage removes simplecontent off the model stage
func (simplecontent *SimpleContent) Unstage(stage *StageStruct) *SimpleContent {
	delete(stage.SimpleContents, simplecontent)
	delete(stage.SimpleContents_mapString, simplecontent.Name)
	return simplecontent
}

// UnstageVoid removes simplecontent off the model stage
func (simplecontent *SimpleContent) UnstageVoid(stage *StageStruct) {
	delete(stage.SimpleContents, simplecontent)
	delete(stage.SimpleContents_mapString, simplecontent.Name)
}

// commit simplecontent to the back repo (if it is already staged)
func (simplecontent *SimpleContent) Commit(stage *StageStruct) *SimpleContent {
	if _, ok := stage.SimpleContents[simplecontent]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSimpleContent(simplecontent)
		}
	}
	return simplecontent
}

func (simplecontent *SimpleContent) CommitVoid(stage *StageStruct) {
	simplecontent.Commit(stage)
}

// Checkout simplecontent to the back repo (if it is already staged)
func (simplecontent *SimpleContent) Checkout(stage *StageStruct) *SimpleContent {
	if _, ok := stage.SimpleContents[simplecontent]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutSimpleContent(simplecontent)
		}
	}
	return simplecontent
}

// for satisfaction of GongStruct interface
func (simplecontent *SimpleContent) GetName() (res string) {
	return simplecontent.Name
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

// Stage puts totaldigit to the model stage
func (totaldigit *TotalDigit) Stage(stage *StageStruct) *TotalDigit {
	stage.TotalDigits[totaldigit] = __member
	stage.TotalDigits_mapString[totaldigit.Name] = totaldigit

	return totaldigit
}

// Unstage removes totaldigit off the model stage
func (totaldigit *TotalDigit) Unstage(stage *StageStruct) *TotalDigit {
	delete(stage.TotalDigits, totaldigit)
	delete(stage.TotalDigits_mapString, totaldigit.Name)
	return totaldigit
}

// UnstageVoid removes totaldigit off the model stage
func (totaldigit *TotalDigit) UnstageVoid(stage *StageStruct) {
	delete(stage.TotalDigits, totaldigit)
	delete(stage.TotalDigits_mapString, totaldigit.Name)
}

// commit totaldigit to the back repo (if it is already staged)
func (totaldigit *TotalDigit) Commit(stage *StageStruct) *TotalDigit {
	if _, ok := stage.TotalDigits[totaldigit]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitTotalDigit(totaldigit)
		}
	}
	return totaldigit
}

func (totaldigit *TotalDigit) CommitVoid(stage *StageStruct) {
	totaldigit.Commit(stage)
}

// Checkout totaldigit to the back repo (if it is already staged)
func (totaldigit *TotalDigit) Checkout(stage *StageStruct) *TotalDigit {
	if _, ok := stage.TotalDigits[totaldigit]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutTotalDigit(totaldigit)
		}
	}
	return totaldigit
}

// for satisfaction of GongStruct interface
func (totaldigit *TotalDigit) GetName() (res string) {
	return totaldigit.Name
}

// Stage puts union to the model stage
func (union *Union) Stage(stage *StageStruct) *Union {
	stage.Unions[union] = __member
	stage.Unions_mapString[union.Name] = union

	return union
}

// Unstage removes union off the model stage
func (union *Union) Unstage(stage *StageStruct) *Union {
	delete(stage.Unions, union)
	delete(stage.Unions_mapString, union.Name)
	return union
}

// UnstageVoid removes union off the model stage
func (union *Union) UnstageVoid(stage *StageStruct) {
	delete(stage.Unions, union)
	delete(stage.Unions_mapString, union.Name)
}

// commit union to the back repo (if it is already staged)
func (union *Union) Commit(stage *StageStruct) *Union {
	if _, ok := stage.Unions[union]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitUnion(union)
		}
	}
	return union
}

func (union *Union) CommitVoid(stage *StageStruct) {
	union.Commit(stage)
}

// Checkout union to the back repo (if it is already staged)
func (union *Union) Checkout(stage *StageStruct) *Union {
	if _, ok := stage.Unions[union]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutUnion(union)
		}
	}
	return union
}

// for satisfaction of GongStruct interface
func (union *Union) GetName() (res string) {
	return union.Name
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
	CreateORMAll(All *All)
	CreateORMAnnotation(Annotation *Annotation)
	CreateORMAttribute(Attribute *Attribute)
	CreateORMAttributeGroup(AttributeGroup *AttributeGroup)
	CreateORMChoice(Choice *Choice)
	CreateORMComplexContent(ComplexContent *ComplexContent)
	CreateORMComplexType(ComplexType *ComplexType)
	CreateORMDocumentation(Documentation *Documentation)
	CreateORMElement(Element *Element)
	CreateORMEnumeration(Enumeration *Enumeration)
	CreateORMExtension(Extension *Extension)
	CreateORMGroup(Group *Group)
	CreateORMLength(Length *Length)
	CreateORMMaxInclusive(MaxInclusive *MaxInclusive)
	CreateORMMaxLength(MaxLength *MaxLength)
	CreateORMMinInclusive(MinInclusive *MinInclusive)
	CreateORMMinLength(MinLength *MinLength)
	CreateORMPattern(Pattern *Pattern)
	CreateORMRestriction(Restriction *Restriction)
	CreateORMSchema(Schema *Schema)
	CreateORMSequence(Sequence *Sequence)
	CreateORMSimpleContent(SimpleContent *SimpleContent)
	CreateORMSimpleType(SimpleType *SimpleType)
	CreateORMTotalDigit(TotalDigit *TotalDigit)
	CreateORMUnion(Union *Union)
	CreateORMWhiteSpace(WhiteSpace *WhiteSpace)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMAll(All *All)
	DeleteORMAnnotation(Annotation *Annotation)
	DeleteORMAttribute(Attribute *Attribute)
	DeleteORMAttributeGroup(AttributeGroup *AttributeGroup)
	DeleteORMChoice(Choice *Choice)
	DeleteORMComplexContent(ComplexContent *ComplexContent)
	DeleteORMComplexType(ComplexType *ComplexType)
	DeleteORMDocumentation(Documentation *Documentation)
	DeleteORMElement(Element *Element)
	DeleteORMEnumeration(Enumeration *Enumeration)
	DeleteORMExtension(Extension *Extension)
	DeleteORMGroup(Group *Group)
	DeleteORMLength(Length *Length)
	DeleteORMMaxInclusive(MaxInclusive *MaxInclusive)
	DeleteORMMaxLength(MaxLength *MaxLength)
	DeleteORMMinInclusive(MinInclusive *MinInclusive)
	DeleteORMMinLength(MinLength *MinLength)
	DeleteORMPattern(Pattern *Pattern)
	DeleteORMRestriction(Restriction *Restriction)
	DeleteORMSchema(Schema *Schema)
	DeleteORMSequence(Sequence *Sequence)
	DeleteORMSimpleContent(SimpleContent *SimpleContent)
	DeleteORMSimpleType(SimpleType *SimpleType)
	DeleteORMTotalDigit(TotalDigit *TotalDigit)
	DeleteORMUnion(Union *Union)
	DeleteORMWhiteSpace(WhiteSpace *WhiteSpace)
}

func (stage *StageStruct) Reset() { // insertion point for array reset
	stage.Alls = make(map[*All]any)
	stage.Alls_mapString = make(map[string]*All)

	stage.Annotations = make(map[*Annotation]any)
	stage.Annotations_mapString = make(map[string]*Annotation)

	stage.Attributes = make(map[*Attribute]any)
	stage.Attributes_mapString = make(map[string]*Attribute)

	stage.AttributeGroups = make(map[*AttributeGroup]any)
	stage.AttributeGroups_mapString = make(map[string]*AttributeGroup)

	stage.Choices = make(map[*Choice]any)
	stage.Choices_mapString = make(map[string]*Choice)

	stage.ComplexContents = make(map[*ComplexContent]any)
	stage.ComplexContents_mapString = make(map[string]*ComplexContent)

	stage.ComplexTypes = make(map[*ComplexType]any)
	stage.ComplexTypes_mapString = make(map[string]*ComplexType)

	stage.Documentations = make(map[*Documentation]any)
	stage.Documentations_mapString = make(map[string]*Documentation)

	stage.Elements = make(map[*Element]any)
	stage.Elements_mapString = make(map[string]*Element)

	stage.Enumerations = make(map[*Enumeration]any)
	stage.Enumerations_mapString = make(map[string]*Enumeration)

	stage.Extensions = make(map[*Extension]any)
	stage.Extensions_mapString = make(map[string]*Extension)

	stage.Groups = make(map[*Group]any)
	stage.Groups_mapString = make(map[string]*Group)

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

	stage.SimpleContents = make(map[*SimpleContent]any)
	stage.SimpleContents_mapString = make(map[string]*SimpleContent)

	stage.SimpleTypes = make(map[*SimpleType]any)
	stage.SimpleTypes_mapString = make(map[string]*SimpleType)

	stage.TotalDigits = make(map[*TotalDigit]any)
	stage.TotalDigits_mapString = make(map[string]*TotalDigit)

	stage.Unions = make(map[*Union]any)
	stage.Unions_mapString = make(map[string]*Union)

	stage.WhiteSpaces = make(map[*WhiteSpace]any)
	stage.WhiteSpaces_mapString = make(map[string]*WhiteSpace)

}

func (stage *StageStruct) Nil() { // insertion point for array nil
	stage.Alls = nil
	stage.Alls_mapString = nil

	stage.Annotations = nil
	stage.Annotations_mapString = nil

	stage.Attributes = nil
	stage.Attributes_mapString = nil

	stage.AttributeGroups = nil
	stage.AttributeGroups_mapString = nil

	stage.Choices = nil
	stage.Choices_mapString = nil

	stage.ComplexContents = nil
	stage.ComplexContents_mapString = nil

	stage.ComplexTypes = nil
	stage.ComplexTypes_mapString = nil

	stage.Documentations = nil
	stage.Documentations_mapString = nil

	stage.Elements = nil
	stage.Elements_mapString = nil

	stage.Enumerations = nil
	stage.Enumerations_mapString = nil

	stage.Extensions = nil
	stage.Extensions_mapString = nil

	stage.Groups = nil
	stage.Groups_mapString = nil

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

	stage.SimpleContents = nil
	stage.SimpleContents_mapString = nil

	stage.SimpleTypes = nil
	stage.SimpleTypes_mapString = nil

	stage.TotalDigits = nil
	stage.TotalDigits_mapString = nil

	stage.Unions = nil
	stage.Unions_mapString = nil

	stage.WhiteSpaces = nil
	stage.WhiteSpaces_mapString = nil

}

func (stage *StageStruct) Unstage() { // insertion point for array nil
	for all := range stage.Alls {
		all.Unstage(stage)
	}

	for annotation := range stage.Annotations {
		annotation.Unstage(stage)
	}

	for attribute := range stage.Attributes {
		attribute.Unstage(stage)
	}

	for attributegroup := range stage.AttributeGroups {
		attributegroup.Unstage(stage)
	}

	for choice := range stage.Choices {
		choice.Unstage(stage)
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

	for element := range stage.Elements {
		element.Unstage(stage)
	}

	for enumeration := range stage.Enumerations {
		enumeration.Unstage(stage)
	}

	for extension := range stage.Extensions {
		extension.Unstage(stage)
	}

	for group := range stage.Groups {
		group.Unstage(stage)
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

	for simplecontent := range stage.SimpleContents {
		simplecontent.Unstage(stage)
	}

	for simpletype := range stage.SimpleTypes {
		simpletype.Unstage(stage)
	}

	for totaldigit := range stage.TotalDigits {
		totaldigit.Unstage(stage)
	}

	for union := range stage.Unions {
		union.Unstage(stage)
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
	case map[*All]any:
		return any(&stage.Alls).(*Type)
	case map[*Annotation]any:
		return any(&stage.Annotations).(*Type)
	case map[*Attribute]any:
		return any(&stage.Attributes).(*Type)
	case map[*AttributeGroup]any:
		return any(&stage.AttributeGroups).(*Type)
	case map[*Choice]any:
		return any(&stage.Choices).(*Type)
	case map[*ComplexContent]any:
		return any(&stage.ComplexContents).(*Type)
	case map[*ComplexType]any:
		return any(&stage.ComplexTypes).(*Type)
	case map[*Documentation]any:
		return any(&stage.Documentations).(*Type)
	case map[*Element]any:
		return any(&stage.Elements).(*Type)
	case map[*Enumeration]any:
		return any(&stage.Enumerations).(*Type)
	case map[*Extension]any:
		return any(&stage.Extensions).(*Type)
	case map[*Group]any:
		return any(&stage.Groups).(*Type)
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
	case map[*SimpleContent]any:
		return any(&stage.SimpleContents).(*Type)
	case map[*SimpleType]any:
		return any(&stage.SimpleTypes).(*Type)
	case map[*TotalDigit]any:
		return any(&stage.TotalDigits).(*Type)
	case map[*Union]any:
		return any(&stage.Unions).(*Type)
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
	case map[string]*All:
		return any(&stage.Alls_mapString).(*Type)
	case map[string]*Annotation:
		return any(&stage.Annotations_mapString).(*Type)
	case map[string]*Attribute:
		return any(&stage.Attributes_mapString).(*Type)
	case map[string]*AttributeGroup:
		return any(&stage.AttributeGroups_mapString).(*Type)
	case map[string]*Choice:
		return any(&stage.Choices_mapString).(*Type)
	case map[string]*ComplexContent:
		return any(&stage.ComplexContents_mapString).(*Type)
	case map[string]*ComplexType:
		return any(&stage.ComplexTypes_mapString).(*Type)
	case map[string]*Documentation:
		return any(&stage.Documentations_mapString).(*Type)
	case map[string]*Element:
		return any(&stage.Elements_mapString).(*Type)
	case map[string]*Enumeration:
		return any(&stage.Enumerations_mapString).(*Type)
	case map[string]*Extension:
		return any(&stage.Extensions_mapString).(*Type)
	case map[string]*Group:
		return any(&stage.Groups_mapString).(*Type)
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
	case map[string]*SimpleContent:
		return any(&stage.SimpleContents_mapString).(*Type)
	case map[string]*SimpleType:
		return any(&stage.SimpleTypes_mapString).(*Type)
	case map[string]*TotalDigit:
		return any(&stage.TotalDigits_mapString).(*Type)
	case map[string]*Union:
		return any(&stage.Unions_mapString).(*Type)
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
	case All:
		return any(&stage.Alls).(*map[*Type]any)
	case Annotation:
		return any(&stage.Annotations).(*map[*Type]any)
	case Attribute:
		return any(&stage.Attributes).(*map[*Type]any)
	case AttributeGroup:
		return any(&stage.AttributeGroups).(*map[*Type]any)
	case Choice:
		return any(&stage.Choices).(*map[*Type]any)
	case ComplexContent:
		return any(&stage.ComplexContents).(*map[*Type]any)
	case ComplexType:
		return any(&stage.ComplexTypes).(*map[*Type]any)
	case Documentation:
		return any(&stage.Documentations).(*map[*Type]any)
	case Element:
		return any(&stage.Elements).(*map[*Type]any)
	case Enumeration:
		return any(&stage.Enumerations).(*map[*Type]any)
	case Extension:
		return any(&stage.Extensions).(*map[*Type]any)
	case Group:
		return any(&stage.Groups).(*map[*Type]any)
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
	case SimpleContent:
		return any(&stage.SimpleContents).(*map[*Type]any)
	case SimpleType:
		return any(&stage.SimpleTypes).(*map[*Type]any)
	case TotalDigit:
		return any(&stage.TotalDigits).(*map[*Type]any)
	case Union:
		return any(&stage.Unions).(*map[*Type]any)
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
	case *All:
		return any(&stage.Alls).(*map[Type]any)
	case *Annotation:
		return any(&stage.Annotations).(*map[Type]any)
	case *Attribute:
		return any(&stage.Attributes).(*map[Type]any)
	case *AttributeGroup:
		return any(&stage.AttributeGroups).(*map[Type]any)
	case *Choice:
		return any(&stage.Choices).(*map[Type]any)
	case *ComplexContent:
		return any(&stage.ComplexContents).(*map[Type]any)
	case *ComplexType:
		return any(&stage.ComplexTypes).(*map[Type]any)
	case *Documentation:
		return any(&stage.Documentations).(*map[Type]any)
	case *Element:
		return any(&stage.Elements).(*map[Type]any)
	case *Enumeration:
		return any(&stage.Enumerations).(*map[Type]any)
	case *Extension:
		return any(&stage.Extensions).(*map[Type]any)
	case *Group:
		return any(&stage.Groups).(*map[Type]any)
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
	case *SimpleContent:
		return any(&stage.SimpleContents).(*map[Type]any)
	case *SimpleType:
		return any(&stage.SimpleTypes).(*map[Type]any)
	case *TotalDigit:
		return any(&stage.TotalDigits).(*map[Type]any)
	case *Union:
		return any(&stage.Unions).(*map[Type]any)
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
	case All:
		return any(&stage.Alls_mapString).(*map[string]*Type)
	case Annotation:
		return any(&stage.Annotations_mapString).(*map[string]*Type)
	case Attribute:
		return any(&stage.Attributes_mapString).(*map[string]*Type)
	case AttributeGroup:
		return any(&stage.AttributeGroups_mapString).(*map[string]*Type)
	case Choice:
		return any(&stage.Choices_mapString).(*map[string]*Type)
	case ComplexContent:
		return any(&stage.ComplexContents_mapString).(*map[string]*Type)
	case ComplexType:
		return any(&stage.ComplexTypes_mapString).(*map[string]*Type)
	case Documentation:
		return any(&stage.Documentations_mapString).(*map[string]*Type)
	case Element:
		return any(&stage.Elements_mapString).(*map[string]*Type)
	case Enumeration:
		return any(&stage.Enumerations_mapString).(*map[string]*Type)
	case Extension:
		return any(&stage.Extensions_mapString).(*map[string]*Type)
	case Group:
		return any(&stage.Groups_mapString).(*map[string]*Type)
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
	case SimpleContent:
		return any(&stage.SimpleContents_mapString).(*map[string]*Type)
	case SimpleType:
		return any(&stage.SimpleTypes_mapString).(*map[string]*Type)
	case TotalDigit:
		return any(&stage.TotalDigits_mapString).(*map[string]*Type)
	case Union:
		return any(&stage.Unions_mapString).(*map[string]*Type)
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
	case All:
		return any(&All{
			// Initialisation of associations
			// field is initialized with Annotated as it is a composite
			Annotated: Annotated{
				// per field init
				//
				Annotation: &Annotation{Name: "Annotation"},
			},
		}).(*Type)
	case Annotation:
		return any(&Annotation{
			// Initialisation of associations
			// field is initialized with an instance of Documentation with the name of the field
			Documentations: []*Documentation{{Name: "Documentations"}},
		}).(*Type)
	case Attribute:
		return any(&Attribute{
			// Initialisation of associations
			// field is initialized with Annotated as it is a composite
			Annotated: Annotated{
				// per field init
				//
				Annotation: &Annotation{Name: "Annotation"},
			},
		}).(*Type)
	case AttributeGroup:
		return any(&AttributeGroup{
			// Initialisation of associations
			// field is initialized with an instance of AttributeGroup with the name of the field
			AttributeGroups: []*AttributeGroup{{Name: "AttributeGroups"}},
			// field is initialized with an instance of Attribute with the name of the field
			Attributes: []*Attribute{{Name: "Attributes"}},
			// field is initialized with Annotated as it is a composite
			Annotated: Annotated{
				// per field init
				//
				Annotation: &Annotation{Name: "Annotation"},
			},
		}).(*Type)
	case Choice:
		return any(&Choice{
			// Initialisation of associations
			// field is initialized with Annotated as it is a composite
			Annotated: Annotated{
				// per field init
				//
				Annotation: &Annotation{Name: "Annotation"},
			},
		}).(*Type)
	case ComplexContent:
		return any(&ComplexContent{
			// Initialisation of associations
		}).(*Type)
	case ComplexType:
		return any(&ComplexType{
			// Initialisation of associations
			// field is initialized with an instance of Element with the name of the field
			OuterElement: &Element{Name: "OuterElement"},
			// field is initialized with an instance of Extension with the name of the field
			Extension: &Extension{Name: "Extension"},
			// field is initialized with an instance of SimpleContent with the name of the field
			SimpleContent: &SimpleContent{Name: "SimpleContent"},
			// field is initialized with an instance of ComplexContent with the name of the field
			ComplexContent: &ComplexContent{Name: "ComplexContent"},
			// field is initialized with an instance of Attribute with the name of the field
			Attributes: []*Attribute{{Name: "Attributes"}},
			// field is initialized with an instance of AttributeGroup with the name of the field
			AttributeGroups: []*AttributeGroup{{Name: "AttributeGroups"}},
			// field is initialized with Annotated as it is a composite
			Annotated: Annotated{
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
			// field is initialized with an instance of Group with the name of the field
			Groups: []*Group{{Name: "Groups"}},
			// field is initialized with Annotated as it is a composite
			Annotated: Annotated{
				// per field init
				//
				Annotation: &Annotation{Name: "Annotation"},
			},
		}).(*Type)
	case Enumeration:
		return any(&Enumeration{
			// Initialisation of associations
			// field is initialized with Annotated as it is a composite
			Annotated: Annotated{
				// per field init
				//
				Annotation: &Annotation{Name: "Annotation"},
			},
		}).(*Type)
	case Extension:
		return any(&Extension{
			// Initialisation of associations
			// field is initialized with an instance of Attribute with the name of the field
			Attributes: []*Attribute{{Name: "Attributes"}},
		}).(*Type)
	case Group:
		return any(&Group{
			// Initialisation of associations
			// field is initialized with an instance of Element with the name of the field
			OuterElement: &Element{Name: "OuterElement"},
			// field is initialized with Annotated as it is a composite
			Annotated: Annotated{
				// per field init
				//
				Annotation: &Annotation{Name: "Annotation"},
			},
		}).(*Type)
	case Length:
		return any(&Length{
			// Initialisation of associations
			// field is initialized with Annotated as it is a composite
			Annotated: Annotated{
				// per field init
				//
				Annotation: &Annotation{Name: "Annotation"},
			},
		}).(*Type)
	case MaxInclusive:
		return any(&MaxInclusive{
			// Initialisation of associations
			// field is initialized with Annotated as it is a composite
			Annotated: Annotated{
				// per field init
				//
				Annotation: &Annotation{Name: "Annotation"},
			},
		}).(*Type)
	case MaxLength:
		return any(&MaxLength{
			// Initialisation of associations
			// field is initialized with Annotated as it is a composite
			Annotated: Annotated{
				// per field init
				//
				Annotation: &Annotation{Name: "Annotation"},
			},
		}).(*Type)
	case MinInclusive:
		return any(&MinInclusive{
			// Initialisation of associations
			// field is initialized with Annotated as it is a composite
			Annotated: Annotated{
				// per field init
				//
				Annotation: &Annotation{Name: "Annotation"},
			},
		}).(*Type)
	case MinLength:
		return any(&MinLength{
			// Initialisation of associations
			// field is initialized with Annotated as it is a composite
			Annotated: Annotated{
				// per field init
				//
				Annotation: &Annotation{Name: "Annotation"},
			},
		}).(*Type)
	case Pattern:
		return any(&Pattern{
			// Initialisation of associations
			// field is initialized with Annotated as it is a composite
			Annotated: Annotated{
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
			// field is initialized with an instance of TotalDigit with the name of the field
			TotalDigit: &TotalDigit{Name: "TotalDigit"},
			// field is initialized with Annotated as it is a composite
			Annotated: Annotated{
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
			// field is initialized with an instance of AttributeGroup with the name of the field
			AttributeGroups: []*AttributeGroup{{Name: "AttributeGroups"}},
			// field is initialized with an instance of Group with the name of the field
			Groups: []*Group{{Name: "Groups"}},
			// field is initialized with Annotated as it is a composite
			Annotated: Annotated{
				// per field init
				//
				Annotation: &Annotation{Name: "Annotation"},
			},
		}).(*Type)
	case Sequence:
		return any(&Sequence{
			// Initialisation of associations
			// field is initialized with Annotated as it is a composite
			Annotated: Annotated{
				// per field init
				//
				Annotation: &Annotation{Name: "Annotation"},
			},
		}).(*Type)
	case SimpleContent:
		return any(&SimpleContent{
			// Initialisation of associations
			// field is initialized with an instance of Extension with the name of the field
			Extension: &Extension{Name: "Extension"},
			// field is initialized with an instance of Restriction with the name of the field
			Restriction: &Restriction{Name: "Restriction"},
		}).(*Type)
	case SimpleType:
		return any(&SimpleType{
			// Initialisation of associations
			// field is initialized with an instance of Restriction with the name of the field
			Restriction: &Restriction{Name: "Restriction"},
			// field is initialized with an instance of Union with the name of the field
			Union: &Union{Name: "Union"},
			// field is initialized with Annotated as it is a composite
			Annotated: Annotated{
				// per field init
				//
				Annotation: &Annotation{Name: "Annotation"},
			},
		}).(*Type)
	case TotalDigit:
		return any(&TotalDigit{
			// Initialisation of associations
			// field is initialized with Annotated as it is a composite
			Annotated: Annotated{
				// per field init
				//
				Annotation: &Annotation{Name: "Annotation"},
			},
		}).(*Type)
	case Union:
		return any(&Union{
			// Initialisation of associations
			// field is initialized with Annotated as it is a composite
			Annotated: Annotated{
				// per field init
				//
				Annotation: &Annotation{Name: "Annotation"},
			},
		}).(*Type)
	case WhiteSpace:
		return any(&WhiteSpace{
			// Initialisation of associations
			// field is initialized with Annotated as it is a composite
			Annotated: Annotated{
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
	// reverse maps of direct associations of All
	case All:
		switch fieldname {
		// insertion point for per direct association field
		case "Annotation":
			res := make(map[*Annotation][]*All)
			for all := range stage.Alls {
				if all.Annotation != nil {
					annotation_ := all.Annotation
					var alls []*All
					_, ok := res[annotation_]
					if ok {
						alls = res[annotation_]
					} else {
						alls = make([]*All, 0)
					}
					alls = append(alls, all)
					res[annotation_] = alls
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Annotation
	case Annotation:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Attribute
	case Attribute:
		switch fieldname {
		// insertion point for per direct association field
		case "Annotation":
			res := make(map[*Annotation][]*Attribute)
			for attribute := range stage.Attributes {
				if attribute.Annotation != nil {
					annotation_ := attribute.Annotation
					var attributes []*Attribute
					_, ok := res[annotation_]
					if ok {
						attributes = res[annotation_]
					} else {
						attributes = make([]*Attribute, 0)
					}
					attributes = append(attributes, attribute)
					res[annotation_] = attributes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of AttributeGroup
	case AttributeGroup:
		switch fieldname {
		// insertion point for per direct association field
		case "Annotation":
			res := make(map[*Annotation][]*AttributeGroup)
			for attributegroup := range stage.AttributeGroups {
				if attributegroup.Annotation != nil {
					annotation_ := attributegroup.Annotation
					var attributegroups []*AttributeGroup
					_, ok := res[annotation_]
					if ok {
						attributegroups = res[annotation_]
					} else {
						attributegroups = make([]*AttributeGroup, 0)
					}
					attributegroups = append(attributegroups, attributegroup)
					res[annotation_] = attributegroups
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Choice
	case Choice:
		switch fieldname {
		// insertion point for per direct association field
		case "Annotation":
			res := make(map[*Annotation][]*Choice)
			for choice := range stage.Choices {
				if choice.Annotation != nil {
					annotation_ := choice.Annotation
					var choices []*Choice
					_, ok := res[annotation_]
					if ok {
						choices = res[annotation_]
					} else {
						choices = make([]*Choice, 0)
					}
					choices = append(choices, choice)
					res[annotation_] = choices
				}
			}
			return any(res).(map[*End][]*Start)
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
		case "OuterElement":
			res := make(map[*Element][]*ComplexType)
			for complextype := range stage.ComplexTypes {
				if complextype.OuterElement != nil {
					element_ := complextype.OuterElement
					var complextypes []*ComplexType
					_, ok := res[element_]
					if ok {
						complextypes = res[element_]
					} else {
						complextypes = make([]*ComplexType, 0)
					}
					complextypes = append(complextypes, complextype)
					res[element_] = complextypes
				}
			}
			return any(res).(map[*End][]*Start)
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
		case "Extension":
			res := make(map[*Extension][]*ComplexType)
			for complextype := range stage.ComplexTypes {
				if complextype.Extension != nil {
					extension_ := complextype.Extension
					var complextypes []*ComplexType
					_, ok := res[extension_]
					if ok {
						complextypes = res[extension_]
					} else {
						complextypes = make([]*ComplexType, 0)
					}
					complextypes = append(complextypes, complextype)
					res[extension_] = complextypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "SimpleContent":
			res := make(map[*SimpleContent][]*ComplexType)
			for complextype := range stage.ComplexTypes {
				if complextype.SimpleContent != nil {
					simplecontent_ := complextype.SimpleContent
					var complextypes []*ComplexType
					_, ok := res[simplecontent_]
					if ok {
						complextypes = res[simplecontent_]
					} else {
						complextypes = make([]*ComplexType, 0)
					}
					complextypes = append(complextypes, complextype)
					res[simplecontent_] = complextypes
				}
			}
			return any(res).(map[*End][]*Start)
		case "ComplexContent":
			res := make(map[*ComplexContent][]*ComplexType)
			for complextype := range stage.ComplexTypes {
				if complextype.ComplexContent != nil {
					complexcontent_ := complextype.ComplexContent
					var complextypes []*ComplexType
					_, ok := res[complexcontent_]
					if ok {
						complextypes = res[complexcontent_]
					} else {
						complextypes = make([]*ComplexType, 0)
					}
					complextypes = append(complextypes, complextype)
					res[complexcontent_] = complextypes
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
	// reverse maps of direct associations of Extension
	case Extension:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Group
	case Group:
		switch fieldname {
		// insertion point for per direct association field
		case "Annotation":
			res := make(map[*Annotation][]*Group)
			for group := range stage.Groups {
				if group.Annotation != nil {
					annotation_ := group.Annotation
					var groups []*Group
					_, ok := res[annotation_]
					if ok {
						groups = res[annotation_]
					} else {
						groups = make([]*Group, 0)
					}
					groups = append(groups, group)
					res[annotation_] = groups
				}
			}
			return any(res).(map[*End][]*Start)
		case "OuterElement":
			res := make(map[*Element][]*Group)
			for group := range stage.Groups {
				if group.OuterElement != nil {
					element_ := group.OuterElement
					var groups []*Group
					_, ok := res[element_]
					if ok {
						groups = res[element_]
					} else {
						groups = make([]*Group, 0)
					}
					groups = append(groups, group)
					res[element_] = groups
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
		case "TotalDigit":
			res := make(map[*TotalDigit][]*Restriction)
			for restriction := range stage.Restrictions {
				if restriction.TotalDigit != nil {
					totaldigit_ := restriction.TotalDigit
					var restrictions []*Restriction
					_, ok := res[totaldigit_]
					if ok {
						restrictions = res[totaldigit_]
					} else {
						restrictions = make([]*Restriction, 0)
					}
					restrictions = append(restrictions, restriction)
					res[totaldigit_] = restrictions
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
	// reverse maps of direct associations of SimpleContent
	case SimpleContent:
		switch fieldname {
		// insertion point for per direct association field
		case "Extension":
			res := make(map[*Extension][]*SimpleContent)
			for simplecontent := range stage.SimpleContents {
				if simplecontent.Extension != nil {
					extension_ := simplecontent.Extension
					var simplecontents []*SimpleContent
					_, ok := res[extension_]
					if ok {
						simplecontents = res[extension_]
					} else {
						simplecontents = make([]*SimpleContent, 0)
					}
					simplecontents = append(simplecontents, simplecontent)
					res[extension_] = simplecontents
				}
			}
			return any(res).(map[*End][]*Start)
		case "Restriction":
			res := make(map[*Restriction][]*SimpleContent)
			for simplecontent := range stage.SimpleContents {
				if simplecontent.Restriction != nil {
					restriction_ := simplecontent.Restriction
					var simplecontents []*SimpleContent
					_, ok := res[restriction_]
					if ok {
						simplecontents = res[restriction_]
					} else {
						simplecontents = make([]*SimpleContent, 0)
					}
					simplecontents = append(simplecontents, simplecontent)
					res[restriction_] = simplecontents
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
		case "Union":
			res := make(map[*Union][]*SimpleType)
			for simpletype := range stage.SimpleTypes {
				if simpletype.Union != nil {
					union_ := simpletype.Union
					var simpletypes []*SimpleType
					_, ok := res[union_]
					if ok {
						simpletypes = res[union_]
					} else {
						simpletypes = make([]*SimpleType, 0)
					}
					simpletypes = append(simpletypes, simpletype)
					res[union_] = simpletypes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of TotalDigit
	case TotalDigit:
		switch fieldname {
		// insertion point for per direct association field
		case "Annotation":
			res := make(map[*Annotation][]*TotalDigit)
			for totaldigit := range stage.TotalDigits {
				if totaldigit.Annotation != nil {
					annotation_ := totaldigit.Annotation
					var totaldigits []*TotalDigit
					_, ok := res[annotation_]
					if ok {
						totaldigits = res[annotation_]
					} else {
						totaldigits = make([]*TotalDigit, 0)
					}
					totaldigits = append(totaldigits, totaldigit)
					res[annotation_] = totaldigits
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Union
	case Union:
		switch fieldname {
		// insertion point for per direct association field
		case "Annotation":
			res := make(map[*Annotation][]*Union)
			for union := range stage.Unions {
				if union.Annotation != nil {
					annotation_ := union.Annotation
					var unions []*Union
					_, ok := res[annotation_]
					if ok {
						unions = res[annotation_]
					} else {
						unions = make([]*Union, 0)
					}
					unions = append(unions, union)
					res[annotation_] = unions
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
	// reverse maps of direct associations of All
	case All:
		switch fieldname {
		// insertion point for per direct association field
		case "Sequences":
			res := make(map[*Sequence]*All)
			for all := range stage.Alls {
				for _, sequence_ := range all.Sequences {
					res[sequence_] = all
				}
			}
			return any(res).(map[*End]*Start)
		case "Alls":
			res := make(map[*All]*All)
			for all := range stage.Alls {
				for _, all_ := range all.Alls {
					res[all_] = all
				}
			}
			return any(res).(map[*End]*Start)
		case "Choices":
			res := make(map[*Choice]*All)
			for all := range stage.Alls {
				for _, choice_ := range all.Choices {
					res[choice_] = all
				}
			}
			return any(res).(map[*End]*Start)
		case "Groups":
			res := make(map[*Group]*All)
			for all := range stage.Alls {
				for _, group_ := range all.Groups {
					res[group_] = all
				}
			}
			return any(res).(map[*End]*Start)
		case "Elements":
			res := make(map[*Element]*All)
			for all := range stage.Alls {
				for _, element_ := range all.Elements {
					res[element_] = all
				}
			}
			return any(res).(map[*End]*Start)
		}
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
	// reverse maps of direct associations of Attribute
	case Attribute:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of AttributeGroup
	case AttributeGroup:
		switch fieldname {
		// insertion point for per direct association field
		case "AttributeGroups":
			res := make(map[*AttributeGroup]*AttributeGroup)
			for attributegroup := range stage.AttributeGroups {
				for _, attributegroup_ := range attributegroup.AttributeGroups {
					res[attributegroup_] = attributegroup
				}
			}
			return any(res).(map[*End]*Start)
		case "Attributes":
			res := make(map[*Attribute]*AttributeGroup)
			for attributegroup := range stage.AttributeGroups {
				for _, attribute_ := range attributegroup.Attributes {
					res[attribute_] = attributegroup
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of Choice
	case Choice:
		switch fieldname {
		// insertion point for per direct association field
		case "Sequences":
			res := make(map[*Sequence]*Choice)
			for choice := range stage.Choices {
				for _, sequence_ := range choice.Sequences {
					res[sequence_] = choice
				}
			}
			return any(res).(map[*End]*Start)
		case "Alls":
			res := make(map[*All]*Choice)
			for choice := range stage.Choices {
				for _, all_ := range choice.Alls {
					res[all_] = choice
				}
			}
			return any(res).(map[*End]*Start)
		case "Choices":
			res := make(map[*Choice]*Choice)
			for choice := range stage.Choices {
				for _, choice_ := range choice.Choices {
					res[choice_] = choice
				}
			}
			return any(res).(map[*End]*Start)
		case "Groups":
			res := make(map[*Group]*Choice)
			for choice := range stage.Choices {
				for _, group_ := range choice.Groups {
					res[group_] = choice
				}
			}
			return any(res).(map[*End]*Start)
		case "Elements":
			res := make(map[*Element]*Choice)
			for choice := range stage.Choices {
				for _, element_ := range choice.Elements {
					res[element_] = choice
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
		case "Sequences":
			res := make(map[*Sequence]*ComplexType)
			for complextype := range stage.ComplexTypes {
				for _, sequence_ := range complextype.Sequences {
					res[sequence_] = complextype
				}
			}
			return any(res).(map[*End]*Start)
		case "Alls":
			res := make(map[*All]*ComplexType)
			for complextype := range stage.ComplexTypes {
				for _, all_ := range complextype.Alls {
					res[all_] = complextype
				}
			}
			return any(res).(map[*End]*Start)
		case "Choices":
			res := make(map[*Choice]*ComplexType)
			for complextype := range stage.ComplexTypes {
				for _, choice_ := range complextype.Choices {
					res[choice_] = complextype
				}
			}
			return any(res).(map[*End]*Start)
		case "Groups":
			res := make(map[*Group]*ComplexType)
			for complextype := range stage.ComplexTypes {
				for _, group_ := range complextype.Groups {
					res[group_] = complextype
				}
			}
			return any(res).(map[*End]*Start)
		case "Elements":
			res := make(map[*Element]*ComplexType)
			for complextype := range stage.ComplexTypes {
				for _, element_ := range complextype.Elements {
					res[element_] = complextype
				}
			}
			return any(res).(map[*End]*Start)
		case "Attributes":
			res := make(map[*Attribute]*ComplexType)
			for complextype := range stage.ComplexTypes {
				for _, attribute_ := range complextype.Attributes {
					res[attribute_] = complextype
				}
			}
			return any(res).(map[*End]*Start)
		case "AttributeGroups":
			res := make(map[*AttributeGroup]*ComplexType)
			for complextype := range stage.ComplexTypes {
				for _, attributegroup_ := range complextype.AttributeGroups {
					res[attributegroup_] = complextype
				}
			}
			return any(res).(map[*End]*Start)
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
		case "Groups":
			res := make(map[*Group]*Element)
			for element := range stage.Elements {
				for _, group_ := range element.Groups {
					res[group_] = element
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of Enumeration
	case Enumeration:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Extension
	case Extension:
		switch fieldname {
		// insertion point for per direct association field
		case "Sequences":
			res := make(map[*Sequence]*Extension)
			for extension := range stage.Extensions {
				for _, sequence_ := range extension.Sequences {
					res[sequence_] = extension
				}
			}
			return any(res).(map[*End]*Start)
		case "Alls":
			res := make(map[*All]*Extension)
			for extension := range stage.Extensions {
				for _, all_ := range extension.Alls {
					res[all_] = extension
				}
			}
			return any(res).(map[*End]*Start)
		case "Choices":
			res := make(map[*Choice]*Extension)
			for extension := range stage.Extensions {
				for _, choice_ := range extension.Choices {
					res[choice_] = extension
				}
			}
			return any(res).(map[*End]*Start)
		case "Groups":
			res := make(map[*Group]*Extension)
			for extension := range stage.Extensions {
				for _, group_ := range extension.Groups {
					res[group_] = extension
				}
			}
			return any(res).(map[*End]*Start)
		case "Elements":
			res := make(map[*Element]*Extension)
			for extension := range stage.Extensions {
				for _, element_ := range extension.Elements {
					res[element_] = extension
				}
			}
			return any(res).(map[*End]*Start)
		case "Attributes":
			res := make(map[*Attribute]*Extension)
			for extension := range stage.Extensions {
				for _, attribute_ := range extension.Attributes {
					res[attribute_] = extension
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of Group
	case Group:
		switch fieldname {
		// insertion point for per direct association field
		case "Sequences":
			res := make(map[*Sequence]*Group)
			for group := range stage.Groups {
				for _, sequence_ := range group.Sequences {
					res[sequence_] = group
				}
			}
			return any(res).(map[*End]*Start)
		case "Alls":
			res := make(map[*All]*Group)
			for group := range stage.Groups {
				for _, all_ := range group.Alls {
					res[all_] = group
				}
			}
			return any(res).(map[*End]*Start)
		case "Choices":
			res := make(map[*Choice]*Group)
			for group := range stage.Groups {
				for _, choice_ := range group.Choices {
					res[choice_] = group
				}
			}
			return any(res).(map[*End]*Start)
		case "Groups":
			res := make(map[*Group]*Group)
			for group := range stage.Groups {
				for _, group_ := range group.Groups {
					res[group_] = group
				}
			}
			return any(res).(map[*End]*Start)
		case "Elements":
			res := make(map[*Element]*Group)
			for group := range stage.Groups {
				for _, element_ := range group.Elements {
					res[element_] = group
				}
			}
			return any(res).(map[*End]*Start)
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
		case "AttributeGroups":
			res := make(map[*AttributeGroup]*Schema)
			for schema := range stage.Schemas {
				for _, attributegroup_ := range schema.AttributeGroups {
					res[attributegroup_] = schema
				}
			}
			return any(res).(map[*End]*Start)
		case "Groups":
			res := make(map[*Group]*Schema)
			for schema := range stage.Schemas {
				for _, group_ := range schema.Groups {
					res[group_] = schema
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of Sequence
	case Sequence:
		switch fieldname {
		// insertion point for per direct association field
		case "Sequences":
			res := make(map[*Sequence]*Sequence)
			for sequence := range stage.Sequences {
				for _, sequence_ := range sequence.Sequences {
					res[sequence_] = sequence
				}
			}
			return any(res).(map[*End]*Start)
		case "Alls":
			res := make(map[*All]*Sequence)
			for sequence := range stage.Sequences {
				for _, all_ := range sequence.Alls {
					res[all_] = sequence
				}
			}
			return any(res).(map[*End]*Start)
		case "Choices":
			res := make(map[*Choice]*Sequence)
			for sequence := range stage.Sequences {
				for _, choice_ := range sequence.Choices {
					res[choice_] = sequence
				}
			}
			return any(res).(map[*End]*Start)
		case "Groups":
			res := make(map[*Group]*Sequence)
			for sequence := range stage.Sequences {
				for _, group_ := range sequence.Groups {
					res[group_] = sequence
				}
			}
			return any(res).(map[*End]*Start)
		case "Elements":
			res := make(map[*Element]*Sequence)
			for sequence := range stage.Sequences {
				for _, element_ := range sequence.Elements {
					res[element_] = sequence
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of SimpleContent
	case SimpleContent:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SimpleType
	case SimpleType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of TotalDigit
	case TotalDigit:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Union
	case Union:
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
	case All:
		res = "All"
	case Annotation:
		res = "Annotation"
	case Attribute:
		res = "Attribute"
	case AttributeGroup:
		res = "AttributeGroup"
	case Choice:
		res = "Choice"
	case ComplexContent:
		res = "ComplexContent"
	case ComplexType:
		res = "ComplexType"
	case Documentation:
		res = "Documentation"
	case Element:
		res = "Element"
	case Enumeration:
		res = "Enumeration"
	case Extension:
		res = "Extension"
	case Group:
		res = "Group"
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
	case SimpleContent:
		res = "SimpleContent"
	case SimpleType:
		res = "SimpleType"
	case TotalDigit:
		res = "TotalDigit"
	case Union:
		res = "Union"
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
	case *All:
		res = "All"
	case *Annotation:
		res = "Annotation"
	case *Attribute:
		res = "Attribute"
	case *AttributeGroup:
		res = "AttributeGroup"
	case *Choice:
		res = "Choice"
	case *ComplexContent:
		res = "ComplexContent"
	case *ComplexType:
		res = "ComplexType"
	case *Documentation:
		res = "Documentation"
	case *Element:
		res = "Element"
	case *Enumeration:
		res = "Enumeration"
	case *Extension:
		res = "Extension"
	case *Group:
		res = "Group"
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
	case *SimpleContent:
		res = "SimpleContent"
	case *SimpleType:
		res = "SimpleType"
	case *TotalDigit:
		res = "TotalDigit"
	case *Union:
		res = "Union"
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
	case All:
		res = []string{"Name", "Annotation", "MinOccurs", "MaxOccurs", "OuterElementName", "Sequences", "Alls", "Choices", "Groups", "Elements"}
	case Annotation:
		res = []string{"Name", "Documentations"}
	case Attribute:
		res = []string{"Name", "NameXSD", "Type", "Annotation", "HasNameConflict", "GoIdentifier", "Default", "Use", "Form", "Fixed", "Ref", "TargetNamespace", "SimpleType", "IDXSD"}
	case AttributeGroup:
		res = []string{"Name", "NameXSD", "Annotation", "HasNameConflict", "GoIdentifier", "AttributeGroups", "Ref", "Attributes"}
	case Choice:
		res = []string{"Name", "Annotation", "MinOccurs", "MaxOccurs", "OuterElementName", "Sequences", "Alls", "Choices", "Groups", "Elements"}
	case ComplexContent:
		res = []string{"Name"}
	case ComplexType:
		res = []string{"Name", "HasNameConflict", "GoIdentifier", "IsAnonymous", "OuterElement", "Annotation", "NameXSD", "OuterElementName", "Sequences", "Alls", "Choices", "Groups", "Elements", "Extension", "SimpleContent", "ComplexContent", "Attributes", "AttributeGroups"}
	case Documentation:
		res = []string{"Name", "Text", "Source", "Lang"}
	case Element:
		res = []string{"Name", "Order", "Depth", "HasNameConflict", "GoIdentifier", "Annotation", "NameXSD", "Type", "MinOccurs", "MaxOccurs", "Default", "Fixed", "Nillable", "Ref", "Abstract", "Form", "Block", "Final", "SimpleType", "ComplexType", "Groups"}
	case Enumeration:
		res = []string{"Name", "Annotation", "Value"}
	case Extension:
		res = []string{"Name", "OuterElementName", "Sequences", "Alls", "Choices", "Groups", "Elements", "Base", "Ref", "Attributes"}
	case Group:
		res = []string{"Name", "Annotation", "NameXSD", "Ref", "IsAnonymous", "OuterElement", "HasNameConflict", "GoIdentifier", "OuterElementName", "Sequences", "Alls", "Choices", "Groups", "Elements", "Order", "Depth"}
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
		res = []string{"Name", "Annotation", "Base", "Enumerations", "MinInclusive", "MaxInclusive", "Pattern", "WhiteSpace", "MinLength", "MaxLength", "Length", "TotalDigit"}
	case Schema:
		res = []string{"Name", "Xs", "Annotation", "Elements", "SimpleTypes", "ComplexTypes", "AttributeGroups", "Groups"}
	case Sequence:
		res = []string{"Name", "Annotation", "MinOccurs", "MaxOccurs", "OuterElementName", "Sequences", "Alls", "Choices", "Groups", "Elements"}
	case SimpleContent:
		res = []string{"Name", "Extension", "Restriction"}
	case SimpleType:
		res = []string{"Name", "Annotation", "NameXSD", "Restriction", "Union"}
	case TotalDigit:
		res = []string{"Name", "Annotation", "Value"}
	case Union:
		res = []string{"Name", "Annotation", "MemberTypes"}
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
	case All:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "All"
		rf.Fieldname = "Alls"
		res = append(res, rf)
		rf.GongstructName = "Choice"
		rf.Fieldname = "Alls"
		res = append(res, rf)
		rf.GongstructName = "ComplexType"
		rf.Fieldname = "Alls"
		res = append(res, rf)
		rf.GongstructName = "Extension"
		rf.Fieldname = "Alls"
		res = append(res, rf)
		rf.GongstructName = "Group"
		rf.Fieldname = "Alls"
		res = append(res, rf)
		rf.GongstructName = "Sequence"
		rf.Fieldname = "Alls"
		res = append(res, rf)
	case Annotation:
		var rf ReverseField
		_ = rf
	case Attribute:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "AttributeGroup"
		rf.Fieldname = "Attributes"
		res = append(res, rf)
		rf.GongstructName = "ComplexType"
		rf.Fieldname = "Attributes"
		res = append(res, rf)
		rf.GongstructName = "Extension"
		rf.Fieldname = "Attributes"
		res = append(res, rf)
	case AttributeGroup:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "AttributeGroup"
		rf.Fieldname = "AttributeGroups"
		res = append(res, rf)
		rf.GongstructName = "ComplexType"
		rf.Fieldname = "AttributeGroups"
		res = append(res, rf)
		rf.GongstructName = "Schema"
		rf.Fieldname = "AttributeGroups"
		res = append(res, rf)
	case Choice:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "All"
		rf.Fieldname = "Choices"
		res = append(res, rf)
		rf.GongstructName = "Choice"
		rf.Fieldname = "Choices"
		res = append(res, rf)
		rf.GongstructName = "ComplexType"
		rf.Fieldname = "Choices"
		res = append(res, rf)
		rf.GongstructName = "Extension"
		rf.Fieldname = "Choices"
		res = append(res, rf)
		rf.GongstructName = "Group"
		rf.Fieldname = "Choices"
		res = append(res, rf)
		rf.GongstructName = "Sequence"
		rf.Fieldname = "Choices"
		res = append(res, rf)
	case ComplexContent:
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
		rf.GongstructName = "All"
		rf.Fieldname = "Elements"
		res = append(res, rf)
		rf.GongstructName = "Choice"
		rf.Fieldname = "Elements"
		res = append(res, rf)
		rf.GongstructName = "ComplexType"
		rf.Fieldname = "Elements"
		res = append(res, rf)
		rf.GongstructName = "Extension"
		rf.Fieldname = "Elements"
		res = append(res, rf)
		rf.GongstructName = "Group"
		rf.Fieldname = "Elements"
		res = append(res, rf)
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
	case Extension:
		var rf ReverseField
		_ = rf
	case Group:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "All"
		rf.Fieldname = "Groups"
		res = append(res, rf)
		rf.GongstructName = "Choice"
		rf.Fieldname = "Groups"
		res = append(res, rf)
		rf.GongstructName = "ComplexType"
		rf.Fieldname = "Groups"
		res = append(res, rf)
		rf.GongstructName = "Element"
		rf.Fieldname = "Groups"
		res = append(res, rf)
		rf.GongstructName = "Extension"
		rf.Fieldname = "Groups"
		res = append(res, rf)
		rf.GongstructName = "Group"
		rf.Fieldname = "Groups"
		res = append(res, rf)
		rf.GongstructName = "Schema"
		rf.Fieldname = "Groups"
		res = append(res, rf)
		rf.GongstructName = "Sequence"
		rf.Fieldname = "Groups"
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
		rf.GongstructName = "All"
		rf.Fieldname = "Sequences"
		res = append(res, rf)
		rf.GongstructName = "Choice"
		rf.Fieldname = "Sequences"
		res = append(res, rf)
		rf.GongstructName = "ComplexType"
		rf.Fieldname = "Sequences"
		res = append(res, rf)
		rf.GongstructName = "Extension"
		rf.Fieldname = "Sequences"
		res = append(res, rf)
		rf.GongstructName = "Group"
		rf.Fieldname = "Sequences"
		res = append(res, rf)
		rf.GongstructName = "Sequence"
		rf.Fieldname = "Sequences"
		res = append(res, rf)
	case SimpleContent:
		var rf ReverseField
		_ = rf
	case SimpleType:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Schema"
		rf.Fieldname = "SimpleTypes"
		res = append(res, rf)
	case TotalDigit:
		var rf ReverseField
		_ = rf
	case Union:
		var rf ReverseField
		_ = rf
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
	case *All:
		res = []string{"Name", "Annotation", "MinOccurs", "MaxOccurs", "OuterElementName", "Sequences", "Alls", "Choices", "Groups", "Elements"}
	case *Annotation:
		res = []string{"Name", "Documentations"}
	case *Attribute:
		res = []string{"Name", "NameXSD", "Type", "Annotation", "HasNameConflict", "GoIdentifier", "Default", "Use", "Form", "Fixed", "Ref", "TargetNamespace", "SimpleType", "IDXSD"}
	case *AttributeGroup:
		res = []string{"Name", "NameXSD", "Annotation", "HasNameConflict", "GoIdentifier", "AttributeGroups", "Ref", "Attributes"}
	case *Choice:
		res = []string{"Name", "Annotation", "MinOccurs", "MaxOccurs", "OuterElementName", "Sequences", "Alls", "Choices", "Groups", "Elements"}
	case *ComplexContent:
		res = []string{"Name"}
	case *ComplexType:
		res = []string{"Name", "HasNameConflict", "GoIdentifier", "IsAnonymous", "OuterElement", "Annotation", "NameXSD", "OuterElementName", "Sequences", "Alls", "Choices", "Groups", "Elements", "Extension", "SimpleContent", "ComplexContent", "Attributes", "AttributeGroups"}
	case *Documentation:
		res = []string{"Name", "Text", "Source", "Lang"}
	case *Element:
		res = []string{"Name", "Order", "Depth", "HasNameConflict", "GoIdentifier", "Annotation", "NameXSD", "Type", "MinOccurs", "MaxOccurs", "Default", "Fixed", "Nillable", "Ref", "Abstract", "Form", "Block", "Final", "SimpleType", "ComplexType", "Groups"}
	case *Enumeration:
		res = []string{"Name", "Annotation", "Value"}
	case *Extension:
		res = []string{"Name", "OuterElementName", "Sequences", "Alls", "Choices", "Groups", "Elements", "Base", "Ref", "Attributes"}
	case *Group:
		res = []string{"Name", "Annotation", "NameXSD", "Ref", "IsAnonymous", "OuterElement", "HasNameConflict", "GoIdentifier", "OuterElementName", "Sequences", "Alls", "Choices", "Groups", "Elements", "Order", "Depth"}
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
		res = []string{"Name", "Annotation", "Base", "Enumerations", "MinInclusive", "MaxInclusive", "Pattern", "WhiteSpace", "MinLength", "MaxLength", "Length", "TotalDigit"}
	case *Schema:
		res = []string{"Name", "Xs", "Annotation", "Elements", "SimpleTypes", "ComplexTypes", "AttributeGroups", "Groups"}
	case *Sequence:
		res = []string{"Name", "Annotation", "MinOccurs", "MaxOccurs", "OuterElementName", "Sequences", "Alls", "Choices", "Groups", "Elements"}
	case *SimpleContent:
		res = []string{"Name", "Extension", "Restriction"}
	case *SimpleType:
		res = []string{"Name", "Annotation", "NameXSD", "Restriction", "Union"}
	case *TotalDigit:
		res = []string{"Name", "Annotation", "Value"}
	case *Union:
		res = []string{"Name", "Annotation", "MemberTypes"}
	case *WhiteSpace:
		res = []string{"Name", "Annotation", "Value"}
	}
	return
}

func GetFieldStringValueFromPointer[Type PointerToGongstruct](instance Type, fieldName string) (res string) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
	case *All:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Annotation":
			if inferedInstance.Annotation != nil {
				res = inferedInstance.Annotation.Name
			}
		case "MinOccurs":
			res = inferedInstance.MinOccurs
		case "MaxOccurs":
			res = inferedInstance.MaxOccurs
		case "OuterElementName":
			res = inferedInstance.OuterElementName
		case "Sequences":
			for idx, __instance__ := range inferedInstance.Sequences {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "Alls":
			for idx, __instance__ := range inferedInstance.Alls {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "Choices":
			for idx, __instance__ := range inferedInstance.Choices {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "Groups":
			for idx, __instance__ := range inferedInstance.Groups {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "Elements":
			for idx, __instance__ := range inferedInstance.Elements {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
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
	case *Attribute:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "NameXSD":
			res = inferedInstance.NameXSD
		case "Type":
			res = inferedInstance.Type
		case "Annotation":
			if inferedInstance.Annotation != nil {
				res = inferedInstance.Annotation.Name
			}
		case "HasNameConflict":
			res = fmt.Sprintf("%t", inferedInstance.HasNameConflict)
		case "GoIdentifier":
			res = inferedInstance.GoIdentifier
		case "Default":
			res = inferedInstance.Default
		case "Use":
			res = inferedInstance.Use
		case "Form":
			res = inferedInstance.Form
		case "Fixed":
			res = inferedInstance.Fixed
		case "Ref":
			res = inferedInstance.Ref
		case "TargetNamespace":
			res = inferedInstance.TargetNamespace
		case "SimpleType":
			res = inferedInstance.SimpleType
		case "IDXSD":
			res = inferedInstance.IDXSD
		}
	case *AttributeGroup:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "NameXSD":
			res = inferedInstance.NameXSD
		case "Annotation":
			if inferedInstance.Annotation != nil {
				res = inferedInstance.Annotation.Name
			}
		case "HasNameConflict":
			res = fmt.Sprintf("%t", inferedInstance.HasNameConflict)
		case "GoIdentifier":
			res = inferedInstance.GoIdentifier
		case "AttributeGroups":
			for idx, __instance__ := range inferedInstance.AttributeGroups {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "Ref":
			res = inferedInstance.Ref
		case "Attributes":
			for idx, __instance__ := range inferedInstance.Attributes {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case *Choice:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Annotation":
			if inferedInstance.Annotation != nil {
				res = inferedInstance.Annotation.Name
			}
		case "MinOccurs":
			res = inferedInstance.MinOccurs
		case "MaxOccurs":
			res = inferedInstance.MaxOccurs
		case "OuterElementName":
			res = inferedInstance.OuterElementName
		case "Sequences":
			for idx, __instance__ := range inferedInstance.Sequences {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "Alls":
			for idx, __instance__ := range inferedInstance.Alls {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "Choices":
			for idx, __instance__ := range inferedInstance.Choices {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "Groups":
			for idx, __instance__ := range inferedInstance.Groups {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "Elements":
			for idx, __instance__ := range inferedInstance.Elements {
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
		case "HasNameConflict":
			res = fmt.Sprintf("%t", inferedInstance.HasNameConflict)
		case "GoIdentifier":
			res = inferedInstance.GoIdentifier
		case "IsAnonymous":
			res = fmt.Sprintf("%t", inferedInstance.IsAnonymous)
		case "OuterElement":
			if inferedInstance.OuterElement != nil {
				res = inferedInstance.OuterElement.Name
			}
		case "Annotation":
			if inferedInstance.Annotation != nil {
				res = inferedInstance.Annotation.Name
			}
		case "NameXSD":
			res = inferedInstance.NameXSD
		case "OuterElementName":
			res = inferedInstance.OuterElementName
		case "Sequences":
			for idx, __instance__ := range inferedInstance.Sequences {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "Alls":
			for idx, __instance__ := range inferedInstance.Alls {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "Choices":
			for idx, __instance__ := range inferedInstance.Choices {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "Groups":
			for idx, __instance__ := range inferedInstance.Groups {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "Elements":
			for idx, __instance__ := range inferedInstance.Elements {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "Extension":
			if inferedInstance.Extension != nil {
				res = inferedInstance.Extension.Name
			}
		case "SimpleContent":
			if inferedInstance.SimpleContent != nil {
				res = inferedInstance.SimpleContent.Name
			}
		case "ComplexContent":
			if inferedInstance.ComplexContent != nil {
				res = inferedInstance.ComplexContent.Name
			}
		case "Attributes":
			for idx, __instance__ := range inferedInstance.Attributes {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "AttributeGroups":
			for idx, __instance__ := range inferedInstance.AttributeGroups {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
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
	case *Element:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Order":
			res = fmt.Sprintf("%d", inferedInstance.Order)
		case "Depth":
			res = fmt.Sprintf("%d", inferedInstance.Depth)
		case "HasNameConflict":
			res = fmt.Sprintf("%t", inferedInstance.HasNameConflict)
		case "GoIdentifier":
			res = inferedInstance.GoIdentifier
		case "Annotation":
			if inferedInstance.Annotation != nil {
				res = inferedInstance.Annotation.Name
			}
		case "NameXSD":
			res = inferedInstance.NameXSD
		case "Type":
			res = inferedInstance.Type
		case "MinOccurs":
			res = inferedInstance.MinOccurs
		case "MaxOccurs":
			res = inferedInstance.MaxOccurs
		case "Default":
			res = inferedInstance.Default
		case "Fixed":
			res = inferedInstance.Fixed
		case "Nillable":
			res = inferedInstance.Nillable
		case "Ref":
			res = inferedInstance.Ref
		case "Abstract":
			res = inferedInstance.Abstract
		case "Form":
			res = inferedInstance.Form
		case "Block":
			res = inferedInstance.Block
		case "Final":
			res = inferedInstance.Final
		case "SimpleType":
			if inferedInstance.SimpleType != nil {
				res = inferedInstance.SimpleType.Name
			}
		case "ComplexType":
			if inferedInstance.ComplexType != nil {
				res = inferedInstance.ComplexType.Name
			}
		case "Groups":
			for idx, __instance__ := range inferedInstance.Groups {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
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
	case *Extension:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "OuterElementName":
			res = inferedInstance.OuterElementName
		case "Sequences":
			for idx, __instance__ := range inferedInstance.Sequences {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "Alls":
			for idx, __instance__ := range inferedInstance.Alls {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "Choices":
			for idx, __instance__ := range inferedInstance.Choices {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "Groups":
			for idx, __instance__ := range inferedInstance.Groups {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "Elements":
			for idx, __instance__ := range inferedInstance.Elements {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "Base":
			res = inferedInstance.Base
		case "Ref":
			res = inferedInstance.Ref
		case "Attributes":
			for idx, __instance__ := range inferedInstance.Attributes {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case *Group:
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
		case "Ref":
			res = inferedInstance.Ref
		case "IsAnonymous":
			res = fmt.Sprintf("%t", inferedInstance.IsAnonymous)
		case "OuterElement":
			if inferedInstance.OuterElement != nil {
				res = inferedInstance.OuterElement.Name
			}
		case "HasNameConflict":
			res = fmt.Sprintf("%t", inferedInstance.HasNameConflict)
		case "GoIdentifier":
			res = inferedInstance.GoIdentifier
		case "OuterElementName":
			res = inferedInstance.OuterElementName
		case "Sequences":
			for idx, __instance__ := range inferedInstance.Sequences {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "Alls":
			for idx, __instance__ := range inferedInstance.Alls {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "Choices":
			for idx, __instance__ := range inferedInstance.Choices {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "Groups":
			for idx, __instance__ := range inferedInstance.Groups {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "Elements":
			for idx, __instance__ := range inferedInstance.Elements {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "Order":
			res = fmt.Sprintf("%d", inferedInstance.Order)
		case "Depth":
			res = fmt.Sprintf("%d", inferedInstance.Depth)
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
		case "TotalDigit":
			if inferedInstance.TotalDigit != nil {
				res = inferedInstance.TotalDigit.Name
			}
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
		case "AttributeGroups":
			for idx, __instance__ := range inferedInstance.AttributeGroups {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "Groups":
			for idx, __instance__ := range inferedInstance.Groups {
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
		case "MinOccurs":
			res = inferedInstance.MinOccurs
		case "MaxOccurs":
			res = inferedInstance.MaxOccurs
		case "OuterElementName":
			res = inferedInstance.OuterElementName
		case "Sequences":
			for idx, __instance__ := range inferedInstance.Sequences {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "Alls":
			for idx, __instance__ := range inferedInstance.Alls {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "Choices":
			for idx, __instance__ := range inferedInstance.Choices {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "Groups":
			for idx, __instance__ := range inferedInstance.Groups {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "Elements":
			for idx, __instance__ := range inferedInstance.Elements {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case *SimpleContent:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Extension":
			if inferedInstance.Extension != nil {
				res = inferedInstance.Extension.Name
			}
		case "Restriction":
			if inferedInstance.Restriction != nil {
				res = inferedInstance.Restriction.Name
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
		case "Union":
			if inferedInstance.Union != nil {
				res = inferedInstance.Union.Name
			}
		}
	case *TotalDigit:
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
	case *Union:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Annotation":
			if inferedInstance.Annotation != nil {
				res = inferedInstance.Annotation.Name
			}
		case "MemberTypes":
			res = inferedInstance.MemberTypes
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
	case All:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Annotation":
			if inferedInstance.Annotation != nil {
				res = inferedInstance.Annotation.Name
			}
		case "MinOccurs":
			res = inferedInstance.MinOccurs
		case "MaxOccurs":
			res = inferedInstance.MaxOccurs
		case "OuterElementName":
			res = inferedInstance.OuterElementName
		case "Sequences":
			for idx, __instance__ := range inferedInstance.Sequences {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "Alls":
			for idx, __instance__ := range inferedInstance.Alls {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "Choices":
			for idx, __instance__ := range inferedInstance.Choices {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "Groups":
			for idx, __instance__ := range inferedInstance.Groups {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "Elements":
			for idx, __instance__ := range inferedInstance.Elements {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
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
	case Attribute:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "NameXSD":
			res = inferedInstance.NameXSD
		case "Type":
			res = inferedInstance.Type
		case "Annotation":
			if inferedInstance.Annotation != nil {
				res = inferedInstance.Annotation.Name
			}
		case "HasNameConflict":
			res = fmt.Sprintf("%t", inferedInstance.HasNameConflict)
		case "GoIdentifier":
			res = inferedInstance.GoIdentifier
		case "Default":
			res = inferedInstance.Default
		case "Use":
			res = inferedInstance.Use
		case "Form":
			res = inferedInstance.Form
		case "Fixed":
			res = inferedInstance.Fixed
		case "Ref":
			res = inferedInstance.Ref
		case "TargetNamespace":
			res = inferedInstance.TargetNamespace
		case "SimpleType":
			res = inferedInstance.SimpleType
		case "IDXSD":
			res = inferedInstance.IDXSD
		}
	case AttributeGroup:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "NameXSD":
			res = inferedInstance.NameXSD
		case "Annotation":
			if inferedInstance.Annotation != nil {
				res = inferedInstance.Annotation.Name
			}
		case "HasNameConflict":
			res = fmt.Sprintf("%t", inferedInstance.HasNameConflict)
		case "GoIdentifier":
			res = inferedInstance.GoIdentifier
		case "AttributeGroups":
			for idx, __instance__ := range inferedInstance.AttributeGroups {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "Ref":
			res = inferedInstance.Ref
		case "Attributes":
			for idx, __instance__ := range inferedInstance.Attributes {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case Choice:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Annotation":
			if inferedInstance.Annotation != nil {
				res = inferedInstance.Annotation.Name
			}
		case "MinOccurs":
			res = inferedInstance.MinOccurs
		case "MaxOccurs":
			res = inferedInstance.MaxOccurs
		case "OuterElementName":
			res = inferedInstance.OuterElementName
		case "Sequences":
			for idx, __instance__ := range inferedInstance.Sequences {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "Alls":
			for idx, __instance__ := range inferedInstance.Alls {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "Choices":
			for idx, __instance__ := range inferedInstance.Choices {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "Groups":
			for idx, __instance__ := range inferedInstance.Groups {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "Elements":
			for idx, __instance__ := range inferedInstance.Elements {
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
		case "HasNameConflict":
			res = fmt.Sprintf("%t", inferedInstance.HasNameConflict)
		case "GoIdentifier":
			res = inferedInstance.GoIdentifier
		case "IsAnonymous":
			res = fmt.Sprintf("%t", inferedInstance.IsAnonymous)
		case "OuterElement":
			if inferedInstance.OuterElement != nil {
				res = inferedInstance.OuterElement.Name
			}
		case "Annotation":
			if inferedInstance.Annotation != nil {
				res = inferedInstance.Annotation.Name
			}
		case "NameXSD":
			res = inferedInstance.NameXSD
		case "OuterElementName":
			res = inferedInstance.OuterElementName
		case "Sequences":
			for idx, __instance__ := range inferedInstance.Sequences {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "Alls":
			for idx, __instance__ := range inferedInstance.Alls {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "Choices":
			for idx, __instance__ := range inferedInstance.Choices {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "Groups":
			for idx, __instance__ := range inferedInstance.Groups {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "Elements":
			for idx, __instance__ := range inferedInstance.Elements {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "Extension":
			if inferedInstance.Extension != nil {
				res = inferedInstance.Extension.Name
			}
		case "SimpleContent":
			if inferedInstance.SimpleContent != nil {
				res = inferedInstance.SimpleContent.Name
			}
		case "ComplexContent":
			if inferedInstance.ComplexContent != nil {
				res = inferedInstance.ComplexContent.Name
			}
		case "Attributes":
			for idx, __instance__ := range inferedInstance.Attributes {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "AttributeGroups":
			for idx, __instance__ := range inferedInstance.AttributeGroups {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
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
	case Element:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Order":
			res = fmt.Sprintf("%d", inferedInstance.Order)
		case "Depth":
			res = fmt.Sprintf("%d", inferedInstance.Depth)
		case "HasNameConflict":
			res = fmt.Sprintf("%t", inferedInstance.HasNameConflict)
		case "GoIdentifier":
			res = inferedInstance.GoIdentifier
		case "Annotation":
			if inferedInstance.Annotation != nil {
				res = inferedInstance.Annotation.Name
			}
		case "NameXSD":
			res = inferedInstance.NameXSD
		case "Type":
			res = inferedInstance.Type
		case "MinOccurs":
			res = inferedInstance.MinOccurs
		case "MaxOccurs":
			res = inferedInstance.MaxOccurs
		case "Default":
			res = inferedInstance.Default
		case "Fixed":
			res = inferedInstance.Fixed
		case "Nillable":
			res = inferedInstance.Nillable
		case "Ref":
			res = inferedInstance.Ref
		case "Abstract":
			res = inferedInstance.Abstract
		case "Form":
			res = inferedInstance.Form
		case "Block":
			res = inferedInstance.Block
		case "Final":
			res = inferedInstance.Final
		case "SimpleType":
			if inferedInstance.SimpleType != nil {
				res = inferedInstance.SimpleType.Name
			}
		case "ComplexType":
			if inferedInstance.ComplexType != nil {
				res = inferedInstance.ComplexType.Name
			}
		case "Groups":
			for idx, __instance__ := range inferedInstance.Groups {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
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
	case Extension:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "OuterElementName":
			res = inferedInstance.OuterElementName
		case "Sequences":
			for idx, __instance__ := range inferedInstance.Sequences {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "Alls":
			for idx, __instance__ := range inferedInstance.Alls {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "Choices":
			for idx, __instance__ := range inferedInstance.Choices {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "Groups":
			for idx, __instance__ := range inferedInstance.Groups {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "Elements":
			for idx, __instance__ := range inferedInstance.Elements {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "Base":
			res = inferedInstance.Base
		case "Ref":
			res = inferedInstance.Ref
		case "Attributes":
			for idx, __instance__ := range inferedInstance.Attributes {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case Group:
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
		case "Ref":
			res = inferedInstance.Ref
		case "IsAnonymous":
			res = fmt.Sprintf("%t", inferedInstance.IsAnonymous)
		case "OuterElement":
			if inferedInstance.OuterElement != nil {
				res = inferedInstance.OuterElement.Name
			}
		case "HasNameConflict":
			res = fmt.Sprintf("%t", inferedInstance.HasNameConflict)
		case "GoIdentifier":
			res = inferedInstance.GoIdentifier
		case "OuterElementName":
			res = inferedInstance.OuterElementName
		case "Sequences":
			for idx, __instance__ := range inferedInstance.Sequences {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "Alls":
			for idx, __instance__ := range inferedInstance.Alls {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "Choices":
			for idx, __instance__ := range inferedInstance.Choices {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "Groups":
			for idx, __instance__ := range inferedInstance.Groups {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "Elements":
			for idx, __instance__ := range inferedInstance.Elements {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "Order":
			res = fmt.Sprintf("%d", inferedInstance.Order)
		case "Depth":
			res = fmt.Sprintf("%d", inferedInstance.Depth)
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
		case "TotalDigit":
			if inferedInstance.TotalDigit != nil {
				res = inferedInstance.TotalDigit.Name
			}
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
		case "AttributeGroups":
			for idx, __instance__ := range inferedInstance.AttributeGroups {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "Groups":
			for idx, __instance__ := range inferedInstance.Groups {
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
		case "MinOccurs":
			res = inferedInstance.MinOccurs
		case "MaxOccurs":
			res = inferedInstance.MaxOccurs
		case "OuterElementName":
			res = inferedInstance.OuterElementName
		case "Sequences":
			for idx, __instance__ := range inferedInstance.Sequences {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "Alls":
			for idx, __instance__ := range inferedInstance.Alls {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "Choices":
			for idx, __instance__ := range inferedInstance.Choices {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "Groups":
			for idx, __instance__ := range inferedInstance.Groups {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "Elements":
			for idx, __instance__ := range inferedInstance.Elements {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case SimpleContent:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Extension":
			if inferedInstance.Extension != nil {
				res = inferedInstance.Extension.Name
			}
		case "Restriction":
			if inferedInstance.Restriction != nil {
				res = inferedInstance.Restriction.Name
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
		case "Union":
			if inferedInstance.Union != nil {
				res = inferedInstance.Union.Name
			}
		}
	case TotalDigit:
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
	case Union:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Annotation":
			if inferedInstance.Annotation != nil {
				res = inferedInstance.Annotation.Name
			}
		case "MemberTypes":
			res = inferedInstance.MemberTypes
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
