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
	ALTERNATIVE_IDs           map[*ALTERNATIVE_ID]any
	ALTERNATIVE_IDs_mapString map[string]*ALTERNATIVE_ID

	// insertion point for slice of pointers maps
	OnAfterALTERNATIVE_IDCreateCallback OnAfterCreateInterface[ALTERNATIVE_ID]
	OnAfterALTERNATIVE_IDUpdateCallback OnAfterUpdateInterface[ALTERNATIVE_ID]
	OnAfterALTERNATIVE_IDDeleteCallback OnAfterDeleteInterface[ALTERNATIVE_ID]
	OnAfterALTERNATIVE_IDReadCallback   OnAfterReadInterface[ALTERNATIVE_ID]

	ATTRIBUTE_DEFINITION_BOOLEANs           map[*ATTRIBUTE_DEFINITION_BOOLEAN]any
	ATTRIBUTE_DEFINITION_BOOLEANs_mapString map[string]*ATTRIBUTE_DEFINITION_BOOLEAN

	// insertion point for slice of pointers maps
	OnAfterATTRIBUTE_DEFINITION_BOOLEANCreateCallback OnAfterCreateInterface[ATTRIBUTE_DEFINITION_BOOLEAN]
	OnAfterATTRIBUTE_DEFINITION_BOOLEANUpdateCallback OnAfterUpdateInterface[ATTRIBUTE_DEFINITION_BOOLEAN]
	OnAfterATTRIBUTE_DEFINITION_BOOLEANDeleteCallback OnAfterDeleteInterface[ATTRIBUTE_DEFINITION_BOOLEAN]
	OnAfterATTRIBUTE_DEFINITION_BOOLEANReadCallback   OnAfterReadInterface[ATTRIBUTE_DEFINITION_BOOLEAN]

	ATTRIBUTE_DEFINITION_DATEs           map[*ATTRIBUTE_DEFINITION_DATE]any
	ATTRIBUTE_DEFINITION_DATEs_mapString map[string]*ATTRIBUTE_DEFINITION_DATE

	// insertion point for slice of pointers maps
	OnAfterATTRIBUTE_DEFINITION_DATECreateCallback OnAfterCreateInterface[ATTRIBUTE_DEFINITION_DATE]
	OnAfterATTRIBUTE_DEFINITION_DATEUpdateCallback OnAfterUpdateInterface[ATTRIBUTE_DEFINITION_DATE]
	OnAfterATTRIBUTE_DEFINITION_DATEDeleteCallback OnAfterDeleteInterface[ATTRIBUTE_DEFINITION_DATE]
	OnAfterATTRIBUTE_DEFINITION_DATEReadCallback   OnAfterReadInterface[ATTRIBUTE_DEFINITION_DATE]

	ATTRIBUTE_DEFINITION_ENUMERATIONs           map[*ATTRIBUTE_DEFINITION_ENUMERATION]any
	ATTRIBUTE_DEFINITION_ENUMERATIONs_mapString map[string]*ATTRIBUTE_DEFINITION_ENUMERATION

	// insertion point for slice of pointers maps
	OnAfterATTRIBUTE_DEFINITION_ENUMERATIONCreateCallback OnAfterCreateInterface[ATTRIBUTE_DEFINITION_ENUMERATION]
	OnAfterATTRIBUTE_DEFINITION_ENUMERATIONUpdateCallback OnAfterUpdateInterface[ATTRIBUTE_DEFINITION_ENUMERATION]
	OnAfterATTRIBUTE_DEFINITION_ENUMERATIONDeleteCallback OnAfterDeleteInterface[ATTRIBUTE_DEFINITION_ENUMERATION]
	OnAfterATTRIBUTE_DEFINITION_ENUMERATIONReadCallback   OnAfterReadInterface[ATTRIBUTE_DEFINITION_ENUMERATION]

	ATTRIBUTE_DEFINITION_INTEGERs           map[*ATTRIBUTE_DEFINITION_INTEGER]any
	ATTRIBUTE_DEFINITION_INTEGERs_mapString map[string]*ATTRIBUTE_DEFINITION_INTEGER

	// insertion point for slice of pointers maps
	OnAfterATTRIBUTE_DEFINITION_INTEGERCreateCallback OnAfterCreateInterface[ATTRIBUTE_DEFINITION_INTEGER]
	OnAfterATTRIBUTE_DEFINITION_INTEGERUpdateCallback OnAfterUpdateInterface[ATTRIBUTE_DEFINITION_INTEGER]
	OnAfterATTRIBUTE_DEFINITION_INTEGERDeleteCallback OnAfterDeleteInterface[ATTRIBUTE_DEFINITION_INTEGER]
	OnAfterATTRIBUTE_DEFINITION_INTEGERReadCallback   OnAfterReadInterface[ATTRIBUTE_DEFINITION_INTEGER]

	ATTRIBUTE_DEFINITION_REALs           map[*ATTRIBUTE_DEFINITION_REAL]any
	ATTRIBUTE_DEFINITION_REALs_mapString map[string]*ATTRIBUTE_DEFINITION_REAL

	// insertion point for slice of pointers maps
	OnAfterATTRIBUTE_DEFINITION_REALCreateCallback OnAfterCreateInterface[ATTRIBUTE_DEFINITION_REAL]
	OnAfterATTRIBUTE_DEFINITION_REALUpdateCallback OnAfterUpdateInterface[ATTRIBUTE_DEFINITION_REAL]
	OnAfterATTRIBUTE_DEFINITION_REALDeleteCallback OnAfterDeleteInterface[ATTRIBUTE_DEFINITION_REAL]
	OnAfterATTRIBUTE_DEFINITION_REALReadCallback   OnAfterReadInterface[ATTRIBUTE_DEFINITION_REAL]

	ATTRIBUTE_DEFINITION_STRINGs           map[*ATTRIBUTE_DEFINITION_STRING]any
	ATTRIBUTE_DEFINITION_STRINGs_mapString map[string]*ATTRIBUTE_DEFINITION_STRING

	// insertion point for slice of pointers maps
	OnAfterATTRIBUTE_DEFINITION_STRINGCreateCallback OnAfterCreateInterface[ATTRIBUTE_DEFINITION_STRING]
	OnAfterATTRIBUTE_DEFINITION_STRINGUpdateCallback OnAfterUpdateInterface[ATTRIBUTE_DEFINITION_STRING]
	OnAfterATTRIBUTE_DEFINITION_STRINGDeleteCallback OnAfterDeleteInterface[ATTRIBUTE_DEFINITION_STRING]
	OnAfterATTRIBUTE_DEFINITION_STRINGReadCallback   OnAfterReadInterface[ATTRIBUTE_DEFINITION_STRING]

	ATTRIBUTE_DEFINITION_XHTMLs           map[*ATTRIBUTE_DEFINITION_XHTML]any
	ATTRIBUTE_DEFINITION_XHTMLs_mapString map[string]*ATTRIBUTE_DEFINITION_XHTML

	// insertion point for slice of pointers maps
	OnAfterATTRIBUTE_DEFINITION_XHTMLCreateCallback OnAfterCreateInterface[ATTRIBUTE_DEFINITION_XHTML]
	OnAfterATTRIBUTE_DEFINITION_XHTMLUpdateCallback OnAfterUpdateInterface[ATTRIBUTE_DEFINITION_XHTML]
	OnAfterATTRIBUTE_DEFINITION_XHTMLDeleteCallback OnAfterDeleteInterface[ATTRIBUTE_DEFINITION_XHTML]
	OnAfterATTRIBUTE_DEFINITION_XHTMLReadCallback   OnAfterReadInterface[ATTRIBUTE_DEFINITION_XHTML]

	ATTRIBUTE_VALUE_BOOLEANs           map[*ATTRIBUTE_VALUE_BOOLEAN]any
	ATTRIBUTE_VALUE_BOOLEANs_mapString map[string]*ATTRIBUTE_VALUE_BOOLEAN

	// insertion point for slice of pointers maps
	OnAfterATTRIBUTE_VALUE_BOOLEANCreateCallback OnAfterCreateInterface[ATTRIBUTE_VALUE_BOOLEAN]
	OnAfterATTRIBUTE_VALUE_BOOLEANUpdateCallback OnAfterUpdateInterface[ATTRIBUTE_VALUE_BOOLEAN]
	OnAfterATTRIBUTE_VALUE_BOOLEANDeleteCallback OnAfterDeleteInterface[ATTRIBUTE_VALUE_BOOLEAN]
	OnAfterATTRIBUTE_VALUE_BOOLEANReadCallback   OnAfterReadInterface[ATTRIBUTE_VALUE_BOOLEAN]

	ATTRIBUTE_VALUE_DATEs           map[*ATTRIBUTE_VALUE_DATE]any
	ATTRIBUTE_VALUE_DATEs_mapString map[string]*ATTRIBUTE_VALUE_DATE

	// insertion point for slice of pointers maps
	OnAfterATTRIBUTE_VALUE_DATECreateCallback OnAfterCreateInterface[ATTRIBUTE_VALUE_DATE]
	OnAfterATTRIBUTE_VALUE_DATEUpdateCallback OnAfterUpdateInterface[ATTRIBUTE_VALUE_DATE]
	OnAfterATTRIBUTE_VALUE_DATEDeleteCallback OnAfterDeleteInterface[ATTRIBUTE_VALUE_DATE]
	OnAfterATTRIBUTE_VALUE_DATEReadCallback   OnAfterReadInterface[ATTRIBUTE_VALUE_DATE]

	ATTRIBUTE_VALUE_ENUMERATIONs           map[*ATTRIBUTE_VALUE_ENUMERATION]any
	ATTRIBUTE_VALUE_ENUMERATIONs_mapString map[string]*ATTRIBUTE_VALUE_ENUMERATION

	// insertion point for slice of pointers maps
	OnAfterATTRIBUTE_VALUE_ENUMERATIONCreateCallback OnAfterCreateInterface[ATTRIBUTE_VALUE_ENUMERATION]
	OnAfterATTRIBUTE_VALUE_ENUMERATIONUpdateCallback OnAfterUpdateInterface[ATTRIBUTE_VALUE_ENUMERATION]
	OnAfterATTRIBUTE_VALUE_ENUMERATIONDeleteCallback OnAfterDeleteInterface[ATTRIBUTE_VALUE_ENUMERATION]
	OnAfterATTRIBUTE_VALUE_ENUMERATIONReadCallback   OnAfterReadInterface[ATTRIBUTE_VALUE_ENUMERATION]

	ATTRIBUTE_VALUE_INTEGERs           map[*ATTRIBUTE_VALUE_INTEGER]any
	ATTRIBUTE_VALUE_INTEGERs_mapString map[string]*ATTRIBUTE_VALUE_INTEGER

	// insertion point for slice of pointers maps
	OnAfterATTRIBUTE_VALUE_INTEGERCreateCallback OnAfterCreateInterface[ATTRIBUTE_VALUE_INTEGER]
	OnAfterATTRIBUTE_VALUE_INTEGERUpdateCallback OnAfterUpdateInterface[ATTRIBUTE_VALUE_INTEGER]
	OnAfterATTRIBUTE_VALUE_INTEGERDeleteCallback OnAfterDeleteInterface[ATTRIBUTE_VALUE_INTEGER]
	OnAfterATTRIBUTE_VALUE_INTEGERReadCallback   OnAfterReadInterface[ATTRIBUTE_VALUE_INTEGER]

	ATTRIBUTE_VALUE_REALs           map[*ATTRIBUTE_VALUE_REAL]any
	ATTRIBUTE_VALUE_REALs_mapString map[string]*ATTRIBUTE_VALUE_REAL

	// insertion point for slice of pointers maps
	OnAfterATTRIBUTE_VALUE_REALCreateCallback OnAfterCreateInterface[ATTRIBUTE_VALUE_REAL]
	OnAfterATTRIBUTE_VALUE_REALUpdateCallback OnAfterUpdateInterface[ATTRIBUTE_VALUE_REAL]
	OnAfterATTRIBUTE_VALUE_REALDeleteCallback OnAfterDeleteInterface[ATTRIBUTE_VALUE_REAL]
	OnAfterATTRIBUTE_VALUE_REALReadCallback   OnAfterReadInterface[ATTRIBUTE_VALUE_REAL]

	ATTRIBUTE_VALUE_STRINGs           map[*ATTRIBUTE_VALUE_STRING]any
	ATTRIBUTE_VALUE_STRINGs_mapString map[string]*ATTRIBUTE_VALUE_STRING

	// insertion point for slice of pointers maps
	OnAfterATTRIBUTE_VALUE_STRINGCreateCallback OnAfterCreateInterface[ATTRIBUTE_VALUE_STRING]
	OnAfterATTRIBUTE_VALUE_STRINGUpdateCallback OnAfterUpdateInterface[ATTRIBUTE_VALUE_STRING]
	OnAfterATTRIBUTE_VALUE_STRINGDeleteCallback OnAfterDeleteInterface[ATTRIBUTE_VALUE_STRING]
	OnAfterATTRIBUTE_VALUE_STRINGReadCallback   OnAfterReadInterface[ATTRIBUTE_VALUE_STRING]

	ATTRIBUTE_VALUE_XHTMLs           map[*ATTRIBUTE_VALUE_XHTML]any
	ATTRIBUTE_VALUE_XHTMLs_mapString map[string]*ATTRIBUTE_VALUE_XHTML

	// insertion point for slice of pointers maps
	ATTRIBUTE_VALUE_XHTML_THE_VALUE_reverseMap map[*XHTML_CONTENT]*ATTRIBUTE_VALUE_XHTML

	ATTRIBUTE_VALUE_XHTML_THE_ORIGINAL_VALUE_reverseMap map[*XHTML_CONTENT]*ATTRIBUTE_VALUE_XHTML

	OnAfterATTRIBUTE_VALUE_XHTMLCreateCallback OnAfterCreateInterface[ATTRIBUTE_VALUE_XHTML]
	OnAfterATTRIBUTE_VALUE_XHTMLUpdateCallback OnAfterUpdateInterface[ATTRIBUTE_VALUE_XHTML]
	OnAfterATTRIBUTE_VALUE_XHTMLDeleteCallback OnAfterDeleteInterface[ATTRIBUTE_VALUE_XHTML]
	OnAfterATTRIBUTE_VALUE_XHTMLReadCallback   OnAfterReadInterface[ATTRIBUTE_VALUE_XHTML]

	DATATYPE_DEFINITION_BOOLEANs           map[*DATATYPE_DEFINITION_BOOLEAN]any
	DATATYPE_DEFINITION_BOOLEANs_mapString map[string]*DATATYPE_DEFINITION_BOOLEAN

	// insertion point for slice of pointers maps
	OnAfterDATATYPE_DEFINITION_BOOLEANCreateCallback OnAfterCreateInterface[DATATYPE_DEFINITION_BOOLEAN]
	OnAfterDATATYPE_DEFINITION_BOOLEANUpdateCallback OnAfterUpdateInterface[DATATYPE_DEFINITION_BOOLEAN]
	OnAfterDATATYPE_DEFINITION_BOOLEANDeleteCallback OnAfterDeleteInterface[DATATYPE_DEFINITION_BOOLEAN]
	OnAfterDATATYPE_DEFINITION_BOOLEANReadCallback   OnAfterReadInterface[DATATYPE_DEFINITION_BOOLEAN]

	DATATYPE_DEFINITION_DATEs           map[*DATATYPE_DEFINITION_DATE]any
	DATATYPE_DEFINITION_DATEs_mapString map[string]*DATATYPE_DEFINITION_DATE

	// insertion point for slice of pointers maps
	OnAfterDATATYPE_DEFINITION_DATECreateCallback OnAfterCreateInterface[DATATYPE_DEFINITION_DATE]
	OnAfterDATATYPE_DEFINITION_DATEUpdateCallback OnAfterUpdateInterface[DATATYPE_DEFINITION_DATE]
	OnAfterDATATYPE_DEFINITION_DATEDeleteCallback OnAfterDeleteInterface[DATATYPE_DEFINITION_DATE]
	OnAfterDATATYPE_DEFINITION_DATEReadCallback   OnAfterReadInterface[DATATYPE_DEFINITION_DATE]

	DATATYPE_DEFINITION_ENUMERATIONs           map[*DATATYPE_DEFINITION_ENUMERATION]any
	DATATYPE_DEFINITION_ENUMERATIONs_mapString map[string]*DATATYPE_DEFINITION_ENUMERATION

	// insertion point for slice of pointers maps
	OnAfterDATATYPE_DEFINITION_ENUMERATIONCreateCallback OnAfterCreateInterface[DATATYPE_DEFINITION_ENUMERATION]
	OnAfterDATATYPE_DEFINITION_ENUMERATIONUpdateCallback OnAfterUpdateInterface[DATATYPE_DEFINITION_ENUMERATION]
	OnAfterDATATYPE_DEFINITION_ENUMERATIONDeleteCallback OnAfterDeleteInterface[DATATYPE_DEFINITION_ENUMERATION]
	OnAfterDATATYPE_DEFINITION_ENUMERATIONReadCallback   OnAfterReadInterface[DATATYPE_DEFINITION_ENUMERATION]

	DATATYPE_DEFINITION_INTEGERs           map[*DATATYPE_DEFINITION_INTEGER]any
	DATATYPE_DEFINITION_INTEGERs_mapString map[string]*DATATYPE_DEFINITION_INTEGER

	// insertion point for slice of pointers maps
	OnAfterDATATYPE_DEFINITION_INTEGERCreateCallback OnAfterCreateInterface[DATATYPE_DEFINITION_INTEGER]
	OnAfterDATATYPE_DEFINITION_INTEGERUpdateCallback OnAfterUpdateInterface[DATATYPE_DEFINITION_INTEGER]
	OnAfterDATATYPE_DEFINITION_INTEGERDeleteCallback OnAfterDeleteInterface[DATATYPE_DEFINITION_INTEGER]
	OnAfterDATATYPE_DEFINITION_INTEGERReadCallback   OnAfterReadInterface[DATATYPE_DEFINITION_INTEGER]

	DATATYPE_DEFINITION_REALs           map[*DATATYPE_DEFINITION_REAL]any
	DATATYPE_DEFINITION_REALs_mapString map[string]*DATATYPE_DEFINITION_REAL

	// insertion point for slice of pointers maps
	OnAfterDATATYPE_DEFINITION_REALCreateCallback OnAfterCreateInterface[DATATYPE_DEFINITION_REAL]
	OnAfterDATATYPE_DEFINITION_REALUpdateCallback OnAfterUpdateInterface[DATATYPE_DEFINITION_REAL]
	OnAfterDATATYPE_DEFINITION_REALDeleteCallback OnAfterDeleteInterface[DATATYPE_DEFINITION_REAL]
	OnAfterDATATYPE_DEFINITION_REALReadCallback   OnAfterReadInterface[DATATYPE_DEFINITION_REAL]

	DATATYPE_DEFINITION_STRINGs           map[*DATATYPE_DEFINITION_STRING]any
	DATATYPE_DEFINITION_STRINGs_mapString map[string]*DATATYPE_DEFINITION_STRING

	// insertion point for slice of pointers maps
	OnAfterDATATYPE_DEFINITION_STRINGCreateCallback OnAfterCreateInterface[DATATYPE_DEFINITION_STRING]
	OnAfterDATATYPE_DEFINITION_STRINGUpdateCallback OnAfterUpdateInterface[DATATYPE_DEFINITION_STRING]
	OnAfterDATATYPE_DEFINITION_STRINGDeleteCallback OnAfterDeleteInterface[DATATYPE_DEFINITION_STRING]
	OnAfterDATATYPE_DEFINITION_STRINGReadCallback   OnAfterReadInterface[DATATYPE_DEFINITION_STRING]

	DATATYPE_DEFINITION_XHTMLs           map[*DATATYPE_DEFINITION_XHTML]any
	DATATYPE_DEFINITION_XHTMLs_mapString map[string]*DATATYPE_DEFINITION_XHTML

	// insertion point for slice of pointers maps
	OnAfterDATATYPE_DEFINITION_XHTMLCreateCallback OnAfterCreateInterface[DATATYPE_DEFINITION_XHTML]
	OnAfterDATATYPE_DEFINITION_XHTMLUpdateCallback OnAfterUpdateInterface[DATATYPE_DEFINITION_XHTML]
	OnAfterDATATYPE_DEFINITION_XHTMLDeleteCallback OnAfterDeleteInterface[DATATYPE_DEFINITION_XHTML]
	OnAfterDATATYPE_DEFINITION_XHTMLReadCallback   OnAfterReadInterface[DATATYPE_DEFINITION_XHTML]

	EMBEDDED_VALUEs           map[*EMBEDDED_VALUE]any
	EMBEDDED_VALUEs_mapString map[string]*EMBEDDED_VALUE

	// insertion point for slice of pointers maps
	OnAfterEMBEDDED_VALUECreateCallback OnAfterCreateInterface[EMBEDDED_VALUE]
	OnAfterEMBEDDED_VALUEUpdateCallback OnAfterUpdateInterface[EMBEDDED_VALUE]
	OnAfterEMBEDDED_VALUEDeleteCallback OnAfterDeleteInterface[EMBEDDED_VALUE]
	OnAfterEMBEDDED_VALUEReadCallback   OnAfterReadInterface[EMBEDDED_VALUE]

	ENUM_VALUEs           map[*ENUM_VALUE]any
	ENUM_VALUEs_mapString map[string]*ENUM_VALUE

	// insertion point for slice of pointers maps
	OnAfterENUM_VALUECreateCallback OnAfterCreateInterface[ENUM_VALUE]
	OnAfterENUM_VALUEUpdateCallback OnAfterUpdateInterface[ENUM_VALUE]
	OnAfterENUM_VALUEDeleteCallback OnAfterDeleteInterface[ENUM_VALUE]
	OnAfterENUM_VALUEReadCallback   OnAfterReadInterface[ENUM_VALUE]

	RELATION_GROUPs           map[*RELATION_GROUP]any
	RELATION_GROUPs_mapString map[string]*RELATION_GROUP

	// insertion point for slice of pointers maps
	OnAfterRELATION_GROUPCreateCallback OnAfterCreateInterface[RELATION_GROUP]
	OnAfterRELATION_GROUPUpdateCallback OnAfterUpdateInterface[RELATION_GROUP]
	OnAfterRELATION_GROUPDeleteCallback OnAfterDeleteInterface[RELATION_GROUP]
	OnAfterRELATION_GROUPReadCallback   OnAfterReadInterface[RELATION_GROUP]

	RELATION_GROUP_TYPEs           map[*RELATION_GROUP_TYPE]any
	RELATION_GROUP_TYPEs_mapString map[string]*RELATION_GROUP_TYPE

	// insertion point for slice of pointers maps
	OnAfterRELATION_GROUP_TYPECreateCallback OnAfterCreateInterface[RELATION_GROUP_TYPE]
	OnAfterRELATION_GROUP_TYPEUpdateCallback OnAfterUpdateInterface[RELATION_GROUP_TYPE]
	OnAfterRELATION_GROUP_TYPEDeleteCallback OnAfterDeleteInterface[RELATION_GROUP_TYPE]
	OnAfterRELATION_GROUP_TYPEReadCallback   OnAfterReadInterface[RELATION_GROUP_TYPE]

	REQ_IFs           map[*REQ_IF]any
	REQ_IFs_mapString map[string]*REQ_IF

	// insertion point for slice of pointers maps
	OnAfterREQ_IFCreateCallback OnAfterCreateInterface[REQ_IF]
	OnAfterREQ_IFUpdateCallback OnAfterUpdateInterface[REQ_IF]
	OnAfterREQ_IFDeleteCallback OnAfterDeleteInterface[REQ_IF]
	OnAfterREQ_IFReadCallback   OnAfterReadInterface[REQ_IF]

	REQ_IF_CONTENTs           map[*REQ_IF_CONTENT]any
	REQ_IF_CONTENTs_mapString map[string]*REQ_IF_CONTENT

	// insertion point for slice of pointers maps
	OnAfterREQ_IF_CONTENTCreateCallback OnAfterCreateInterface[REQ_IF_CONTENT]
	OnAfterREQ_IF_CONTENTUpdateCallback OnAfterUpdateInterface[REQ_IF_CONTENT]
	OnAfterREQ_IF_CONTENTDeleteCallback OnAfterDeleteInterface[REQ_IF_CONTENT]
	OnAfterREQ_IF_CONTENTReadCallback   OnAfterReadInterface[REQ_IF_CONTENT]

	REQ_IF_HEADERs           map[*REQ_IF_HEADER]any
	REQ_IF_HEADERs_mapString map[string]*REQ_IF_HEADER

	// insertion point for slice of pointers maps
	OnAfterREQ_IF_HEADERCreateCallback OnAfterCreateInterface[REQ_IF_HEADER]
	OnAfterREQ_IF_HEADERUpdateCallback OnAfterUpdateInterface[REQ_IF_HEADER]
	OnAfterREQ_IF_HEADERDeleteCallback OnAfterDeleteInterface[REQ_IF_HEADER]
	OnAfterREQ_IF_HEADERReadCallback   OnAfterReadInterface[REQ_IF_HEADER]

	REQ_IF_TOOL_EXTENSIONs           map[*REQ_IF_TOOL_EXTENSION]any
	REQ_IF_TOOL_EXTENSIONs_mapString map[string]*REQ_IF_TOOL_EXTENSION

	// insertion point for slice of pointers maps
	OnAfterREQ_IF_TOOL_EXTENSIONCreateCallback OnAfterCreateInterface[REQ_IF_TOOL_EXTENSION]
	OnAfterREQ_IF_TOOL_EXTENSIONUpdateCallback OnAfterUpdateInterface[REQ_IF_TOOL_EXTENSION]
	OnAfterREQ_IF_TOOL_EXTENSIONDeleteCallback OnAfterDeleteInterface[REQ_IF_TOOL_EXTENSION]
	OnAfterREQ_IF_TOOL_EXTENSIONReadCallback   OnAfterReadInterface[REQ_IF_TOOL_EXTENSION]

	SPECIFICATIONs           map[*SPECIFICATION]any
	SPECIFICATIONs_mapString map[string]*SPECIFICATION

	// insertion point for slice of pointers maps
	OnAfterSPECIFICATIONCreateCallback OnAfterCreateInterface[SPECIFICATION]
	OnAfterSPECIFICATIONUpdateCallback OnAfterUpdateInterface[SPECIFICATION]
	OnAfterSPECIFICATIONDeleteCallback OnAfterDeleteInterface[SPECIFICATION]
	OnAfterSPECIFICATIONReadCallback   OnAfterReadInterface[SPECIFICATION]

	SPECIFICATION_TYPEs           map[*SPECIFICATION_TYPE]any
	SPECIFICATION_TYPEs_mapString map[string]*SPECIFICATION_TYPE

	// insertion point for slice of pointers maps
	OnAfterSPECIFICATION_TYPECreateCallback OnAfterCreateInterface[SPECIFICATION_TYPE]
	OnAfterSPECIFICATION_TYPEUpdateCallback OnAfterUpdateInterface[SPECIFICATION_TYPE]
	OnAfterSPECIFICATION_TYPEDeleteCallback OnAfterDeleteInterface[SPECIFICATION_TYPE]
	OnAfterSPECIFICATION_TYPEReadCallback   OnAfterReadInterface[SPECIFICATION_TYPE]

	SPEC_HIERARCHYs           map[*SPEC_HIERARCHY]any
	SPEC_HIERARCHYs_mapString map[string]*SPEC_HIERARCHY

	// insertion point for slice of pointers maps
	OnAfterSPEC_HIERARCHYCreateCallback OnAfterCreateInterface[SPEC_HIERARCHY]
	OnAfterSPEC_HIERARCHYUpdateCallback OnAfterUpdateInterface[SPEC_HIERARCHY]
	OnAfterSPEC_HIERARCHYDeleteCallback OnAfterDeleteInterface[SPEC_HIERARCHY]
	OnAfterSPEC_HIERARCHYReadCallback   OnAfterReadInterface[SPEC_HIERARCHY]

	SPEC_OBJECTs           map[*SPEC_OBJECT]any
	SPEC_OBJECTs_mapString map[string]*SPEC_OBJECT

	// insertion point for slice of pointers maps
	OnAfterSPEC_OBJECTCreateCallback OnAfterCreateInterface[SPEC_OBJECT]
	OnAfterSPEC_OBJECTUpdateCallback OnAfterUpdateInterface[SPEC_OBJECT]
	OnAfterSPEC_OBJECTDeleteCallback OnAfterDeleteInterface[SPEC_OBJECT]
	OnAfterSPEC_OBJECTReadCallback   OnAfterReadInterface[SPEC_OBJECT]

	SPEC_OBJECT_TYPEs           map[*SPEC_OBJECT_TYPE]any
	SPEC_OBJECT_TYPEs_mapString map[string]*SPEC_OBJECT_TYPE

	// insertion point for slice of pointers maps
	OnAfterSPEC_OBJECT_TYPECreateCallback OnAfterCreateInterface[SPEC_OBJECT_TYPE]
	OnAfterSPEC_OBJECT_TYPEUpdateCallback OnAfterUpdateInterface[SPEC_OBJECT_TYPE]
	OnAfterSPEC_OBJECT_TYPEDeleteCallback OnAfterDeleteInterface[SPEC_OBJECT_TYPE]
	OnAfterSPEC_OBJECT_TYPEReadCallback   OnAfterReadInterface[SPEC_OBJECT_TYPE]

	SPEC_RELATIONs           map[*SPEC_RELATION]any
	SPEC_RELATIONs_mapString map[string]*SPEC_RELATION

	// insertion point for slice of pointers maps
	OnAfterSPEC_RELATIONCreateCallback OnAfterCreateInterface[SPEC_RELATION]
	OnAfterSPEC_RELATIONUpdateCallback OnAfterUpdateInterface[SPEC_RELATION]
	OnAfterSPEC_RELATIONDeleteCallback OnAfterDeleteInterface[SPEC_RELATION]
	OnAfterSPEC_RELATIONReadCallback   OnAfterReadInterface[SPEC_RELATION]

	SPEC_RELATION_TYPEs           map[*SPEC_RELATION_TYPE]any
	SPEC_RELATION_TYPEs_mapString map[string]*SPEC_RELATION_TYPE

	// insertion point for slice of pointers maps
	OnAfterSPEC_RELATION_TYPECreateCallback OnAfterCreateInterface[SPEC_RELATION_TYPE]
	OnAfterSPEC_RELATION_TYPEUpdateCallback OnAfterUpdateInterface[SPEC_RELATION_TYPE]
	OnAfterSPEC_RELATION_TYPEDeleteCallback OnAfterDeleteInterface[SPEC_RELATION_TYPE]
	OnAfterSPEC_RELATION_TYPEReadCallback   OnAfterReadInterface[SPEC_RELATION_TYPE]

	XHTML_CONTENTs           map[*XHTML_CONTENT]any
	XHTML_CONTENTs_mapString map[string]*XHTML_CONTENT

	// insertion point for slice of pointers maps
	OnAfterXHTML_CONTENTCreateCallback OnAfterCreateInterface[XHTML_CONTENT]
	OnAfterXHTML_CONTENTUpdateCallback OnAfterUpdateInterface[XHTML_CONTENT]
	OnAfterXHTML_CONTENTDeleteCallback OnAfterDeleteInterface[XHTML_CONTENT]
	OnAfterXHTML_CONTENTReadCallback   OnAfterReadInterface[XHTML_CONTENT]

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
	return "github.com/fullstack-lang/gongxsd/test/reqif/go/models"
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
	CommitALTERNATIVE_ID(alternative_id *ALTERNATIVE_ID)
	CheckoutALTERNATIVE_ID(alternative_id *ALTERNATIVE_ID)
	CommitATTRIBUTE_DEFINITION_BOOLEAN(attribute_definition_boolean *ATTRIBUTE_DEFINITION_BOOLEAN)
	CheckoutATTRIBUTE_DEFINITION_BOOLEAN(attribute_definition_boolean *ATTRIBUTE_DEFINITION_BOOLEAN)
	CommitATTRIBUTE_DEFINITION_DATE(attribute_definition_date *ATTRIBUTE_DEFINITION_DATE)
	CheckoutATTRIBUTE_DEFINITION_DATE(attribute_definition_date *ATTRIBUTE_DEFINITION_DATE)
	CommitATTRIBUTE_DEFINITION_ENUMERATION(attribute_definition_enumeration *ATTRIBUTE_DEFINITION_ENUMERATION)
	CheckoutATTRIBUTE_DEFINITION_ENUMERATION(attribute_definition_enumeration *ATTRIBUTE_DEFINITION_ENUMERATION)
	CommitATTRIBUTE_DEFINITION_INTEGER(attribute_definition_integer *ATTRIBUTE_DEFINITION_INTEGER)
	CheckoutATTRIBUTE_DEFINITION_INTEGER(attribute_definition_integer *ATTRIBUTE_DEFINITION_INTEGER)
	CommitATTRIBUTE_DEFINITION_REAL(attribute_definition_real *ATTRIBUTE_DEFINITION_REAL)
	CheckoutATTRIBUTE_DEFINITION_REAL(attribute_definition_real *ATTRIBUTE_DEFINITION_REAL)
	CommitATTRIBUTE_DEFINITION_STRING(attribute_definition_string *ATTRIBUTE_DEFINITION_STRING)
	CheckoutATTRIBUTE_DEFINITION_STRING(attribute_definition_string *ATTRIBUTE_DEFINITION_STRING)
	CommitATTRIBUTE_DEFINITION_XHTML(attribute_definition_xhtml *ATTRIBUTE_DEFINITION_XHTML)
	CheckoutATTRIBUTE_DEFINITION_XHTML(attribute_definition_xhtml *ATTRIBUTE_DEFINITION_XHTML)
	CommitATTRIBUTE_VALUE_BOOLEAN(attribute_value_boolean *ATTRIBUTE_VALUE_BOOLEAN)
	CheckoutATTRIBUTE_VALUE_BOOLEAN(attribute_value_boolean *ATTRIBUTE_VALUE_BOOLEAN)
	CommitATTRIBUTE_VALUE_DATE(attribute_value_date *ATTRIBUTE_VALUE_DATE)
	CheckoutATTRIBUTE_VALUE_DATE(attribute_value_date *ATTRIBUTE_VALUE_DATE)
	CommitATTRIBUTE_VALUE_ENUMERATION(attribute_value_enumeration *ATTRIBUTE_VALUE_ENUMERATION)
	CheckoutATTRIBUTE_VALUE_ENUMERATION(attribute_value_enumeration *ATTRIBUTE_VALUE_ENUMERATION)
	CommitATTRIBUTE_VALUE_INTEGER(attribute_value_integer *ATTRIBUTE_VALUE_INTEGER)
	CheckoutATTRIBUTE_VALUE_INTEGER(attribute_value_integer *ATTRIBUTE_VALUE_INTEGER)
	CommitATTRIBUTE_VALUE_REAL(attribute_value_real *ATTRIBUTE_VALUE_REAL)
	CheckoutATTRIBUTE_VALUE_REAL(attribute_value_real *ATTRIBUTE_VALUE_REAL)
	CommitATTRIBUTE_VALUE_STRING(attribute_value_string *ATTRIBUTE_VALUE_STRING)
	CheckoutATTRIBUTE_VALUE_STRING(attribute_value_string *ATTRIBUTE_VALUE_STRING)
	CommitATTRIBUTE_VALUE_XHTML(attribute_value_xhtml *ATTRIBUTE_VALUE_XHTML)
	CheckoutATTRIBUTE_VALUE_XHTML(attribute_value_xhtml *ATTRIBUTE_VALUE_XHTML)
	CommitDATATYPE_DEFINITION_BOOLEAN(datatype_definition_boolean *DATATYPE_DEFINITION_BOOLEAN)
	CheckoutDATATYPE_DEFINITION_BOOLEAN(datatype_definition_boolean *DATATYPE_DEFINITION_BOOLEAN)
	CommitDATATYPE_DEFINITION_DATE(datatype_definition_date *DATATYPE_DEFINITION_DATE)
	CheckoutDATATYPE_DEFINITION_DATE(datatype_definition_date *DATATYPE_DEFINITION_DATE)
	CommitDATATYPE_DEFINITION_ENUMERATION(datatype_definition_enumeration *DATATYPE_DEFINITION_ENUMERATION)
	CheckoutDATATYPE_DEFINITION_ENUMERATION(datatype_definition_enumeration *DATATYPE_DEFINITION_ENUMERATION)
	CommitDATATYPE_DEFINITION_INTEGER(datatype_definition_integer *DATATYPE_DEFINITION_INTEGER)
	CheckoutDATATYPE_DEFINITION_INTEGER(datatype_definition_integer *DATATYPE_DEFINITION_INTEGER)
	CommitDATATYPE_DEFINITION_REAL(datatype_definition_real *DATATYPE_DEFINITION_REAL)
	CheckoutDATATYPE_DEFINITION_REAL(datatype_definition_real *DATATYPE_DEFINITION_REAL)
	CommitDATATYPE_DEFINITION_STRING(datatype_definition_string *DATATYPE_DEFINITION_STRING)
	CheckoutDATATYPE_DEFINITION_STRING(datatype_definition_string *DATATYPE_DEFINITION_STRING)
	CommitDATATYPE_DEFINITION_XHTML(datatype_definition_xhtml *DATATYPE_DEFINITION_XHTML)
	CheckoutDATATYPE_DEFINITION_XHTML(datatype_definition_xhtml *DATATYPE_DEFINITION_XHTML)
	CommitEMBEDDED_VALUE(embedded_value *EMBEDDED_VALUE)
	CheckoutEMBEDDED_VALUE(embedded_value *EMBEDDED_VALUE)
	CommitENUM_VALUE(enum_value *ENUM_VALUE)
	CheckoutENUM_VALUE(enum_value *ENUM_VALUE)
	CommitRELATION_GROUP(relation_group *RELATION_GROUP)
	CheckoutRELATION_GROUP(relation_group *RELATION_GROUP)
	CommitRELATION_GROUP_TYPE(relation_group_type *RELATION_GROUP_TYPE)
	CheckoutRELATION_GROUP_TYPE(relation_group_type *RELATION_GROUP_TYPE)
	CommitREQ_IF(req_if *REQ_IF)
	CheckoutREQ_IF(req_if *REQ_IF)
	CommitREQ_IF_CONTENT(req_if_content *REQ_IF_CONTENT)
	CheckoutREQ_IF_CONTENT(req_if_content *REQ_IF_CONTENT)
	CommitREQ_IF_HEADER(req_if_header *REQ_IF_HEADER)
	CheckoutREQ_IF_HEADER(req_if_header *REQ_IF_HEADER)
	CommitREQ_IF_TOOL_EXTENSION(req_if_tool_extension *REQ_IF_TOOL_EXTENSION)
	CheckoutREQ_IF_TOOL_EXTENSION(req_if_tool_extension *REQ_IF_TOOL_EXTENSION)
	CommitSPECIFICATION(specification *SPECIFICATION)
	CheckoutSPECIFICATION(specification *SPECIFICATION)
	CommitSPECIFICATION_TYPE(specification_type *SPECIFICATION_TYPE)
	CheckoutSPECIFICATION_TYPE(specification_type *SPECIFICATION_TYPE)
	CommitSPEC_HIERARCHY(spec_hierarchy *SPEC_HIERARCHY)
	CheckoutSPEC_HIERARCHY(spec_hierarchy *SPEC_HIERARCHY)
	CommitSPEC_OBJECT(spec_object *SPEC_OBJECT)
	CheckoutSPEC_OBJECT(spec_object *SPEC_OBJECT)
	CommitSPEC_OBJECT_TYPE(spec_object_type *SPEC_OBJECT_TYPE)
	CheckoutSPEC_OBJECT_TYPE(spec_object_type *SPEC_OBJECT_TYPE)
	CommitSPEC_RELATION(spec_relation *SPEC_RELATION)
	CheckoutSPEC_RELATION(spec_relation *SPEC_RELATION)
	CommitSPEC_RELATION_TYPE(spec_relation_type *SPEC_RELATION_TYPE)
	CheckoutSPEC_RELATION_TYPE(spec_relation_type *SPEC_RELATION_TYPE)
	CommitXHTML_CONTENT(xhtml_content *XHTML_CONTENT)
	CheckoutXHTML_CONTENT(xhtml_content *XHTML_CONTENT)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

func NewStage(path string) (stage *StageStruct) {

	stage = &StageStruct{ // insertion point for array initiatialisation
		ALTERNATIVE_IDs:           make(map[*ALTERNATIVE_ID]any),
		ALTERNATIVE_IDs_mapString: make(map[string]*ALTERNATIVE_ID),

		ATTRIBUTE_DEFINITION_BOOLEANs:           make(map[*ATTRIBUTE_DEFINITION_BOOLEAN]any),
		ATTRIBUTE_DEFINITION_BOOLEANs_mapString: make(map[string]*ATTRIBUTE_DEFINITION_BOOLEAN),

		ATTRIBUTE_DEFINITION_DATEs:           make(map[*ATTRIBUTE_DEFINITION_DATE]any),
		ATTRIBUTE_DEFINITION_DATEs_mapString: make(map[string]*ATTRIBUTE_DEFINITION_DATE),

		ATTRIBUTE_DEFINITION_ENUMERATIONs:           make(map[*ATTRIBUTE_DEFINITION_ENUMERATION]any),
		ATTRIBUTE_DEFINITION_ENUMERATIONs_mapString: make(map[string]*ATTRIBUTE_DEFINITION_ENUMERATION),

		ATTRIBUTE_DEFINITION_INTEGERs:           make(map[*ATTRIBUTE_DEFINITION_INTEGER]any),
		ATTRIBUTE_DEFINITION_INTEGERs_mapString: make(map[string]*ATTRIBUTE_DEFINITION_INTEGER),

		ATTRIBUTE_DEFINITION_REALs:           make(map[*ATTRIBUTE_DEFINITION_REAL]any),
		ATTRIBUTE_DEFINITION_REALs_mapString: make(map[string]*ATTRIBUTE_DEFINITION_REAL),

		ATTRIBUTE_DEFINITION_STRINGs:           make(map[*ATTRIBUTE_DEFINITION_STRING]any),
		ATTRIBUTE_DEFINITION_STRINGs_mapString: make(map[string]*ATTRIBUTE_DEFINITION_STRING),

		ATTRIBUTE_DEFINITION_XHTMLs:           make(map[*ATTRIBUTE_DEFINITION_XHTML]any),
		ATTRIBUTE_DEFINITION_XHTMLs_mapString: make(map[string]*ATTRIBUTE_DEFINITION_XHTML),

		ATTRIBUTE_VALUE_BOOLEANs:           make(map[*ATTRIBUTE_VALUE_BOOLEAN]any),
		ATTRIBUTE_VALUE_BOOLEANs_mapString: make(map[string]*ATTRIBUTE_VALUE_BOOLEAN),

		ATTRIBUTE_VALUE_DATEs:           make(map[*ATTRIBUTE_VALUE_DATE]any),
		ATTRIBUTE_VALUE_DATEs_mapString: make(map[string]*ATTRIBUTE_VALUE_DATE),

		ATTRIBUTE_VALUE_ENUMERATIONs:           make(map[*ATTRIBUTE_VALUE_ENUMERATION]any),
		ATTRIBUTE_VALUE_ENUMERATIONs_mapString: make(map[string]*ATTRIBUTE_VALUE_ENUMERATION),

		ATTRIBUTE_VALUE_INTEGERs:           make(map[*ATTRIBUTE_VALUE_INTEGER]any),
		ATTRIBUTE_VALUE_INTEGERs_mapString: make(map[string]*ATTRIBUTE_VALUE_INTEGER),

		ATTRIBUTE_VALUE_REALs:           make(map[*ATTRIBUTE_VALUE_REAL]any),
		ATTRIBUTE_VALUE_REALs_mapString: make(map[string]*ATTRIBUTE_VALUE_REAL),

		ATTRIBUTE_VALUE_STRINGs:           make(map[*ATTRIBUTE_VALUE_STRING]any),
		ATTRIBUTE_VALUE_STRINGs_mapString: make(map[string]*ATTRIBUTE_VALUE_STRING),

		ATTRIBUTE_VALUE_XHTMLs:           make(map[*ATTRIBUTE_VALUE_XHTML]any),
		ATTRIBUTE_VALUE_XHTMLs_mapString: make(map[string]*ATTRIBUTE_VALUE_XHTML),

		DATATYPE_DEFINITION_BOOLEANs:           make(map[*DATATYPE_DEFINITION_BOOLEAN]any),
		DATATYPE_DEFINITION_BOOLEANs_mapString: make(map[string]*DATATYPE_DEFINITION_BOOLEAN),

		DATATYPE_DEFINITION_DATEs:           make(map[*DATATYPE_DEFINITION_DATE]any),
		DATATYPE_DEFINITION_DATEs_mapString: make(map[string]*DATATYPE_DEFINITION_DATE),

		DATATYPE_DEFINITION_ENUMERATIONs:           make(map[*DATATYPE_DEFINITION_ENUMERATION]any),
		DATATYPE_DEFINITION_ENUMERATIONs_mapString: make(map[string]*DATATYPE_DEFINITION_ENUMERATION),

		DATATYPE_DEFINITION_INTEGERs:           make(map[*DATATYPE_DEFINITION_INTEGER]any),
		DATATYPE_DEFINITION_INTEGERs_mapString: make(map[string]*DATATYPE_DEFINITION_INTEGER),

		DATATYPE_DEFINITION_REALs:           make(map[*DATATYPE_DEFINITION_REAL]any),
		DATATYPE_DEFINITION_REALs_mapString: make(map[string]*DATATYPE_DEFINITION_REAL),

		DATATYPE_DEFINITION_STRINGs:           make(map[*DATATYPE_DEFINITION_STRING]any),
		DATATYPE_DEFINITION_STRINGs_mapString: make(map[string]*DATATYPE_DEFINITION_STRING),

		DATATYPE_DEFINITION_XHTMLs:           make(map[*DATATYPE_DEFINITION_XHTML]any),
		DATATYPE_DEFINITION_XHTMLs_mapString: make(map[string]*DATATYPE_DEFINITION_XHTML),

		EMBEDDED_VALUEs:           make(map[*EMBEDDED_VALUE]any),
		EMBEDDED_VALUEs_mapString: make(map[string]*EMBEDDED_VALUE),

		ENUM_VALUEs:           make(map[*ENUM_VALUE]any),
		ENUM_VALUEs_mapString: make(map[string]*ENUM_VALUE),

		RELATION_GROUPs:           make(map[*RELATION_GROUP]any),
		RELATION_GROUPs_mapString: make(map[string]*RELATION_GROUP),

		RELATION_GROUP_TYPEs:           make(map[*RELATION_GROUP_TYPE]any),
		RELATION_GROUP_TYPEs_mapString: make(map[string]*RELATION_GROUP_TYPE),

		REQ_IFs:           make(map[*REQ_IF]any),
		REQ_IFs_mapString: make(map[string]*REQ_IF),

		REQ_IF_CONTENTs:           make(map[*REQ_IF_CONTENT]any),
		REQ_IF_CONTENTs_mapString: make(map[string]*REQ_IF_CONTENT),

		REQ_IF_HEADERs:           make(map[*REQ_IF_HEADER]any),
		REQ_IF_HEADERs_mapString: make(map[string]*REQ_IF_HEADER),

		REQ_IF_TOOL_EXTENSIONs:           make(map[*REQ_IF_TOOL_EXTENSION]any),
		REQ_IF_TOOL_EXTENSIONs_mapString: make(map[string]*REQ_IF_TOOL_EXTENSION),

		SPECIFICATIONs:           make(map[*SPECIFICATION]any),
		SPECIFICATIONs_mapString: make(map[string]*SPECIFICATION),

		SPECIFICATION_TYPEs:           make(map[*SPECIFICATION_TYPE]any),
		SPECIFICATION_TYPEs_mapString: make(map[string]*SPECIFICATION_TYPE),

		SPEC_HIERARCHYs:           make(map[*SPEC_HIERARCHY]any),
		SPEC_HIERARCHYs_mapString: make(map[string]*SPEC_HIERARCHY),

		SPEC_OBJECTs:           make(map[*SPEC_OBJECT]any),
		SPEC_OBJECTs_mapString: make(map[string]*SPEC_OBJECT),

		SPEC_OBJECT_TYPEs:           make(map[*SPEC_OBJECT_TYPE]any),
		SPEC_OBJECT_TYPEs_mapString: make(map[string]*SPEC_OBJECT_TYPE),

		SPEC_RELATIONs:           make(map[*SPEC_RELATION]any),
		SPEC_RELATIONs_mapString: make(map[string]*SPEC_RELATION),

		SPEC_RELATION_TYPEs:           make(map[*SPEC_RELATION_TYPE]any),
		SPEC_RELATION_TYPEs_mapString: make(map[string]*SPEC_RELATION_TYPE),

		XHTML_CONTENTs:           make(map[*XHTML_CONTENT]any),
		XHTML_CONTENTs_mapString: make(map[string]*XHTML_CONTENT),

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
	stage.Map_GongStructName_InstancesNb["ALTERNATIVE_ID"] = len(stage.ALTERNATIVE_IDs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTE_DEFINITION_BOOLEAN"] = len(stage.ATTRIBUTE_DEFINITION_BOOLEANs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTE_DEFINITION_DATE"] = len(stage.ATTRIBUTE_DEFINITION_DATEs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTE_DEFINITION_ENUMERATION"] = len(stage.ATTRIBUTE_DEFINITION_ENUMERATIONs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTE_DEFINITION_INTEGER"] = len(stage.ATTRIBUTE_DEFINITION_INTEGERs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTE_DEFINITION_REAL"] = len(stage.ATTRIBUTE_DEFINITION_REALs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTE_DEFINITION_STRING"] = len(stage.ATTRIBUTE_DEFINITION_STRINGs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTE_DEFINITION_XHTML"] = len(stage.ATTRIBUTE_DEFINITION_XHTMLs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTE_VALUE_BOOLEAN"] = len(stage.ATTRIBUTE_VALUE_BOOLEANs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTE_VALUE_DATE"] = len(stage.ATTRIBUTE_VALUE_DATEs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTE_VALUE_ENUMERATION"] = len(stage.ATTRIBUTE_VALUE_ENUMERATIONs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTE_VALUE_INTEGER"] = len(stage.ATTRIBUTE_VALUE_INTEGERs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTE_VALUE_REAL"] = len(stage.ATTRIBUTE_VALUE_REALs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTE_VALUE_STRING"] = len(stage.ATTRIBUTE_VALUE_STRINGs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTE_VALUE_XHTML"] = len(stage.ATTRIBUTE_VALUE_XHTMLs)
	stage.Map_GongStructName_InstancesNb["DATATYPE_DEFINITION_BOOLEAN"] = len(stage.DATATYPE_DEFINITION_BOOLEANs)
	stage.Map_GongStructName_InstancesNb["DATATYPE_DEFINITION_DATE"] = len(stage.DATATYPE_DEFINITION_DATEs)
	stage.Map_GongStructName_InstancesNb["DATATYPE_DEFINITION_ENUMERATION"] = len(stage.DATATYPE_DEFINITION_ENUMERATIONs)
	stage.Map_GongStructName_InstancesNb["DATATYPE_DEFINITION_INTEGER"] = len(stage.DATATYPE_DEFINITION_INTEGERs)
	stage.Map_GongStructName_InstancesNb["DATATYPE_DEFINITION_REAL"] = len(stage.DATATYPE_DEFINITION_REALs)
	stage.Map_GongStructName_InstancesNb["DATATYPE_DEFINITION_STRING"] = len(stage.DATATYPE_DEFINITION_STRINGs)
	stage.Map_GongStructName_InstancesNb["DATATYPE_DEFINITION_XHTML"] = len(stage.DATATYPE_DEFINITION_XHTMLs)
	stage.Map_GongStructName_InstancesNb["EMBEDDED_VALUE"] = len(stage.EMBEDDED_VALUEs)
	stage.Map_GongStructName_InstancesNb["ENUM_VALUE"] = len(stage.ENUM_VALUEs)
	stage.Map_GongStructName_InstancesNb["RELATION_GROUP"] = len(stage.RELATION_GROUPs)
	stage.Map_GongStructName_InstancesNb["RELATION_GROUP_TYPE"] = len(stage.RELATION_GROUP_TYPEs)
	stage.Map_GongStructName_InstancesNb["REQ_IF"] = len(stage.REQ_IFs)
	stage.Map_GongStructName_InstancesNb["REQ_IF_CONTENT"] = len(stage.REQ_IF_CONTENTs)
	stage.Map_GongStructName_InstancesNb["REQ_IF_HEADER"] = len(stage.REQ_IF_HEADERs)
	stage.Map_GongStructName_InstancesNb["REQ_IF_TOOL_EXTENSION"] = len(stage.REQ_IF_TOOL_EXTENSIONs)
	stage.Map_GongStructName_InstancesNb["SPECIFICATION"] = len(stage.SPECIFICATIONs)
	stage.Map_GongStructName_InstancesNb["SPECIFICATION_TYPE"] = len(stage.SPECIFICATION_TYPEs)
	stage.Map_GongStructName_InstancesNb["SPEC_HIERARCHY"] = len(stage.SPEC_HIERARCHYs)
	stage.Map_GongStructName_InstancesNb["SPEC_OBJECT"] = len(stage.SPEC_OBJECTs)
	stage.Map_GongStructName_InstancesNb["SPEC_OBJECT_TYPE"] = len(stage.SPEC_OBJECT_TYPEs)
	stage.Map_GongStructName_InstancesNb["SPEC_RELATION"] = len(stage.SPEC_RELATIONs)
	stage.Map_GongStructName_InstancesNb["SPEC_RELATION_TYPE"] = len(stage.SPEC_RELATION_TYPEs)
	stage.Map_GongStructName_InstancesNb["XHTML_CONTENT"] = len(stage.XHTML_CONTENTs)

}

func (stage *StageStruct) Checkout() {
	if stage.BackRepo != nil {
		stage.BackRepo.Checkout(stage)
	}

	stage.ComputeReverseMaps()
	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["ALTERNATIVE_ID"] = len(stage.ALTERNATIVE_IDs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTE_DEFINITION_BOOLEAN"] = len(stage.ATTRIBUTE_DEFINITION_BOOLEANs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTE_DEFINITION_DATE"] = len(stage.ATTRIBUTE_DEFINITION_DATEs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTE_DEFINITION_ENUMERATION"] = len(stage.ATTRIBUTE_DEFINITION_ENUMERATIONs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTE_DEFINITION_INTEGER"] = len(stage.ATTRIBUTE_DEFINITION_INTEGERs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTE_DEFINITION_REAL"] = len(stage.ATTRIBUTE_DEFINITION_REALs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTE_DEFINITION_STRING"] = len(stage.ATTRIBUTE_DEFINITION_STRINGs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTE_DEFINITION_XHTML"] = len(stage.ATTRIBUTE_DEFINITION_XHTMLs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTE_VALUE_BOOLEAN"] = len(stage.ATTRIBUTE_VALUE_BOOLEANs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTE_VALUE_DATE"] = len(stage.ATTRIBUTE_VALUE_DATEs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTE_VALUE_ENUMERATION"] = len(stage.ATTRIBUTE_VALUE_ENUMERATIONs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTE_VALUE_INTEGER"] = len(stage.ATTRIBUTE_VALUE_INTEGERs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTE_VALUE_REAL"] = len(stage.ATTRIBUTE_VALUE_REALs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTE_VALUE_STRING"] = len(stage.ATTRIBUTE_VALUE_STRINGs)
	stage.Map_GongStructName_InstancesNb["ATTRIBUTE_VALUE_XHTML"] = len(stage.ATTRIBUTE_VALUE_XHTMLs)
	stage.Map_GongStructName_InstancesNb["DATATYPE_DEFINITION_BOOLEAN"] = len(stage.DATATYPE_DEFINITION_BOOLEANs)
	stage.Map_GongStructName_InstancesNb["DATATYPE_DEFINITION_DATE"] = len(stage.DATATYPE_DEFINITION_DATEs)
	stage.Map_GongStructName_InstancesNb["DATATYPE_DEFINITION_ENUMERATION"] = len(stage.DATATYPE_DEFINITION_ENUMERATIONs)
	stage.Map_GongStructName_InstancesNb["DATATYPE_DEFINITION_INTEGER"] = len(stage.DATATYPE_DEFINITION_INTEGERs)
	stage.Map_GongStructName_InstancesNb["DATATYPE_DEFINITION_REAL"] = len(stage.DATATYPE_DEFINITION_REALs)
	stage.Map_GongStructName_InstancesNb["DATATYPE_DEFINITION_STRING"] = len(stage.DATATYPE_DEFINITION_STRINGs)
	stage.Map_GongStructName_InstancesNb["DATATYPE_DEFINITION_XHTML"] = len(stage.DATATYPE_DEFINITION_XHTMLs)
	stage.Map_GongStructName_InstancesNb["EMBEDDED_VALUE"] = len(stage.EMBEDDED_VALUEs)
	stage.Map_GongStructName_InstancesNb["ENUM_VALUE"] = len(stage.ENUM_VALUEs)
	stage.Map_GongStructName_InstancesNb["RELATION_GROUP"] = len(stage.RELATION_GROUPs)
	stage.Map_GongStructName_InstancesNb["RELATION_GROUP_TYPE"] = len(stage.RELATION_GROUP_TYPEs)
	stage.Map_GongStructName_InstancesNb["REQ_IF"] = len(stage.REQ_IFs)
	stage.Map_GongStructName_InstancesNb["REQ_IF_CONTENT"] = len(stage.REQ_IF_CONTENTs)
	stage.Map_GongStructName_InstancesNb["REQ_IF_HEADER"] = len(stage.REQ_IF_HEADERs)
	stage.Map_GongStructName_InstancesNb["REQ_IF_TOOL_EXTENSION"] = len(stage.REQ_IF_TOOL_EXTENSIONs)
	stage.Map_GongStructName_InstancesNb["SPECIFICATION"] = len(stage.SPECIFICATIONs)
	stage.Map_GongStructName_InstancesNb["SPECIFICATION_TYPE"] = len(stage.SPECIFICATION_TYPEs)
	stage.Map_GongStructName_InstancesNb["SPEC_HIERARCHY"] = len(stage.SPEC_HIERARCHYs)
	stage.Map_GongStructName_InstancesNb["SPEC_OBJECT"] = len(stage.SPEC_OBJECTs)
	stage.Map_GongStructName_InstancesNb["SPEC_OBJECT_TYPE"] = len(stage.SPEC_OBJECT_TYPEs)
	stage.Map_GongStructName_InstancesNb["SPEC_RELATION"] = len(stage.SPEC_RELATIONs)
	stage.Map_GongStructName_InstancesNb["SPEC_RELATION_TYPE"] = len(stage.SPEC_RELATION_TYPEs)
	stage.Map_GongStructName_InstancesNb["XHTML_CONTENT"] = len(stage.XHTML_CONTENTs)

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
// Stage puts alternative_id to the model stage
func (alternative_id *ALTERNATIVE_ID) Stage(stage *StageStruct) *ALTERNATIVE_ID {
	stage.ALTERNATIVE_IDs[alternative_id] = __member
	stage.ALTERNATIVE_IDs_mapString[alternative_id.Name] = alternative_id

	return alternative_id
}

// Unstage removes alternative_id off the model stage
func (alternative_id *ALTERNATIVE_ID) Unstage(stage *StageStruct) *ALTERNATIVE_ID {
	delete(stage.ALTERNATIVE_IDs, alternative_id)
	delete(stage.ALTERNATIVE_IDs_mapString, alternative_id.Name)
	return alternative_id
}

// UnstageVoid removes alternative_id off the model stage
func (alternative_id *ALTERNATIVE_ID) UnstageVoid(stage *StageStruct) {
	delete(stage.ALTERNATIVE_IDs, alternative_id)
	delete(stage.ALTERNATIVE_IDs_mapString, alternative_id.Name)
}

// commit alternative_id to the back repo (if it is already staged)
func (alternative_id *ALTERNATIVE_ID) Commit(stage *StageStruct) *ALTERNATIVE_ID {
	if _, ok := stage.ALTERNATIVE_IDs[alternative_id]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitALTERNATIVE_ID(alternative_id)
		}
	}
	return alternative_id
}

func (alternative_id *ALTERNATIVE_ID) CommitVoid(stage *StageStruct) {
	alternative_id.Commit(stage)
}

// Checkout alternative_id to the back repo (if it is already staged)
func (alternative_id *ALTERNATIVE_ID) Checkout(stage *StageStruct) *ALTERNATIVE_ID {
	if _, ok := stage.ALTERNATIVE_IDs[alternative_id]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutALTERNATIVE_ID(alternative_id)
		}
	}
	return alternative_id
}

// for satisfaction of GongStruct interface
func (alternative_id *ALTERNATIVE_ID) GetName() (res string) {
	return alternative_id.Name
}

// Stage puts attribute_definition_boolean to the model stage
func (attribute_definition_boolean *ATTRIBUTE_DEFINITION_BOOLEAN) Stage(stage *StageStruct) *ATTRIBUTE_DEFINITION_BOOLEAN {
	stage.ATTRIBUTE_DEFINITION_BOOLEANs[attribute_definition_boolean] = __member
	stage.ATTRIBUTE_DEFINITION_BOOLEANs_mapString[attribute_definition_boolean.Name] = attribute_definition_boolean

	return attribute_definition_boolean
}

// Unstage removes attribute_definition_boolean off the model stage
func (attribute_definition_boolean *ATTRIBUTE_DEFINITION_BOOLEAN) Unstage(stage *StageStruct) *ATTRIBUTE_DEFINITION_BOOLEAN {
	delete(stage.ATTRIBUTE_DEFINITION_BOOLEANs, attribute_definition_boolean)
	delete(stage.ATTRIBUTE_DEFINITION_BOOLEANs_mapString, attribute_definition_boolean.Name)
	return attribute_definition_boolean
}

// UnstageVoid removes attribute_definition_boolean off the model stage
func (attribute_definition_boolean *ATTRIBUTE_DEFINITION_BOOLEAN) UnstageVoid(stage *StageStruct) {
	delete(stage.ATTRIBUTE_DEFINITION_BOOLEANs, attribute_definition_boolean)
	delete(stage.ATTRIBUTE_DEFINITION_BOOLEANs_mapString, attribute_definition_boolean.Name)
}

// commit attribute_definition_boolean to the back repo (if it is already staged)
func (attribute_definition_boolean *ATTRIBUTE_DEFINITION_BOOLEAN) Commit(stage *StageStruct) *ATTRIBUTE_DEFINITION_BOOLEAN {
	if _, ok := stage.ATTRIBUTE_DEFINITION_BOOLEANs[attribute_definition_boolean]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitATTRIBUTE_DEFINITION_BOOLEAN(attribute_definition_boolean)
		}
	}
	return attribute_definition_boolean
}

func (attribute_definition_boolean *ATTRIBUTE_DEFINITION_BOOLEAN) CommitVoid(stage *StageStruct) {
	attribute_definition_boolean.Commit(stage)
}

// Checkout attribute_definition_boolean to the back repo (if it is already staged)
func (attribute_definition_boolean *ATTRIBUTE_DEFINITION_BOOLEAN) Checkout(stage *StageStruct) *ATTRIBUTE_DEFINITION_BOOLEAN {
	if _, ok := stage.ATTRIBUTE_DEFINITION_BOOLEANs[attribute_definition_boolean]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutATTRIBUTE_DEFINITION_BOOLEAN(attribute_definition_boolean)
		}
	}
	return attribute_definition_boolean
}

// for satisfaction of GongStruct interface
func (attribute_definition_boolean *ATTRIBUTE_DEFINITION_BOOLEAN) GetName() (res string) {
	return attribute_definition_boolean.Name
}

// Stage puts attribute_definition_date to the model stage
func (attribute_definition_date *ATTRIBUTE_DEFINITION_DATE) Stage(stage *StageStruct) *ATTRIBUTE_DEFINITION_DATE {
	stage.ATTRIBUTE_DEFINITION_DATEs[attribute_definition_date] = __member
	stage.ATTRIBUTE_DEFINITION_DATEs_mapString[attribute_definition_date.Name] = attribute_definition_date

	return attribute_definition_date
}

// Unstage removes attribute_definition_date off the model stage
func (attribute_definition_date *ATTRIBUTE_DEFINITION_DATE) Unstage(stage *StageStruct) *ATTRIBUTE_DEFINITION_DATE {
	delete(stage.ATTRIBUTE_DEFINITION_DATEs, attribute_definition_date)
	delete(stage.ATTRIBUTE_DEFINITION_DATEs_mapString, attribute_definition_date.Name)
	return attribute_definition_date
}

// UnstageVoid removes attribute_definition_date off the model stage
func (attribute_definition_date *ATTRIBUTE_DEFINITION_DATE) UnstageVoid(stage *StageStruct) {
	delete(stage.ATTRIBUTE_DEFINITION_DATEs, attribute_definition_date)
	delete(stage.ATTRIBUTE_DEFINITION_DATEs_mapString, attribute_definition_date.Name)
}

// commit attribute_definition_date to the back repo (if it is already staged)
func (attribute_definition_date *ATTRIBUTE_DEFINITION_DATE) Commit(stage *StageStruct) *ATTRIBUTE_DEFINITION_DATE {
	if _, ok := stage.ATTRIBUTE_DEFINITION_DATEs[attribute_definition_date]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitATTRIBUTE_DEFINITION_DATE(attribute_definition_date)
		}
	}
	return attribute_definition_date
}

func (attribute_definition_date *ATTRIBUTE_DEFINITION_DATE) CommitVoid(stage *StageStruct) {
	attribute_definition_date.Commit(stage)
}

// Checkout attribute_definition_date to the back repo (if it is already staged)
func (attribute_definition_date *ATTRIBUTE_DEFINITION_DATE) Checkout(stage *StageStruct) *ATTRIBUTE_DEFINITION_DATE {
	if _, ok := stage.ATTRIBUTE_DEFINITION_DATEs[attribute_definition_date]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutATTRIBUTE_DEFINITION_DATE(attribute_definition_date)
		}
	}
	return attribute_definition_date
}

// for satisfaction of GongStruct interface
func (attribute_definition_date *ATTRIBUTE_DEFINITION_DATE) GetName() (res string) {
	return attribute_definition_date.Name
}

// Stage puts attribute_definition_enumeration to the model stage
func (attribute_definition_enumeration *ATTRIBUTE_DEFINITION_ENUMERATION) Stage(stage *StageStruct) *ATTRIBUTE_DEFINITION_ENUMERATION {
	stage.ATTRIBUTE_DEFINITION_ENUMERATIONs[attribute_definition_enumeration] = __member
	stage.ATTRIBUTE_DEFINITION_ENUMERATIONs_mapString[attribute_definition_enumeration.Name] = attribute_definition_enumeration

	return attribute_definition_enumeration
}

// Unstage removes attribute_definition_enumeration off the model stage
func (attribute_definition_enumeration *ATTRIBUTE_DEFINITION_ENUMERATION) Unstage(stage *StageStruct) *ATTRIBUTE_DEFINITION_ENUMERATION {
	delete(stage.ATTRIBUTE_DEFINITION_ENUMERATIONs, attribute_definition_enumeration)
	delete(stage.ATTRIBUTE_DEFINITION_ENUMERATIONs_mapString, attribute_definition_enumeration.Name)
	return attribute_definition_enumeration
}

// UnstageVoid removes attribute_definition_enumeration off the model stage
func (attribute_definition_enumeration *ATTRIBUTE_DEFINITION_ENUMERATION) UnstageVoid(stage *StageStruct) {
	delete(stage.ATTRIBUTE_DEFINITION_ENUMERATIONs, attribute_definition_enumeration)
	delete(stage.ATTRIBUTE_DEFINITION_ENUMERATIONs_mapString, attribute_definition_enumeration.Name)
}

// commit attribute_definition_enumeration to the back repo (if it is already staged)
func (attribute_definition_enumeration *ATTRIBUTE_DEFINITION_ENUMERATION) Commit(stage *StageStruct) *ATTRIBUTE_DEFINITION_ENUMERATION {
	if _, ok := stage.ATTRIBUTE_DEFINITION_ENUMERATIONs[attribute_definition_enumeration]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitATTRIBUTE_DEFINITION_ENUMERATION(attribute_definition_enumeration)
		}
	}
	return attribute_definition_enumeration
}

func (attribute_definition_enumeration *ATTRIBUTE_DEFINITION_ENUMERATION) CommitVoid(stage *StageStruct) {
	attribute_definition_enumeration.Commit(stage)
}

// Checkout attribute_definition_enumeration to the back repo (if it is already staged)
func (attribute_definition_enumeration *ATTRIBUTE_DEFINITION_ENUMERATION) Checkout(stage *StageStruct) *ATTRIBUTE_DEFINITION_ENUMERATION {
	if _, ok := stage.ATTRIBUTE_DEFINITION_ENUMERATIONs[attribute_definition_enumeration]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutATTRIBUTE_DEFINITION_ENUMERATION(attribute_definition_enumeration)
		}
	}
	return attribute_definition_enumeration
}

// for satisfaction of GongStruct interface
func (attribute_definition_enumeration *ATTRIBUTE_DEFINITION_ENUMERATION) GetName() (res string) {
	return attribute_definition_enumeration.Name
}

// Stage puts attribute_definition_integer to the model stage
func (attribute_definition_integer *ATTRIBUTE_DEFINITION_INTEGER) Stage(stage *StageStruct) *ATTRIBUTE_DEFINITION_INTEGER {
	stage.ATTRIBUTE_DEFINITION_INTEGERs[attribute_definition_integer] = __member
	stage.ATTRIBUTE_DEFINITION_INTEGERs_mapString[attribute_definition_integer.Name] = attribute_definition_integer

	return attribute_definition_integer
}

// Unstage removes attribute_definition_integer off the model stage
func (attribute_definition_integer *ATTRIBUTE_DEFINITION_INTEGER) Unstage(stage *StageStruct) *ATTRIBUTE_DEFINITION_INTEGER {
	delete(stage.ATTRIBUTE_DEFINITION_INTEGERs, attribute_definition_integer)
	delete(stage.ATTRIBUTE_DEFINITION_INTEGERs_mapString, attribute_definition_integer.Name)
	return attribute_definition_integer
}

// UnstageVoid removes attribute_definition_integer off the model stage
func (attribute_definition_integer *ATTRIBUTE_DEFINITION_INTEGER) UnstageVoid(stage *StageStruct) {
	delete(stage.ATTRIBUTE_DEFINITION_INTEGERs, attribute_definition_integer)
	delete(stage.ATTRIBUTE_DEFINITION_INTEGERs_mapString, attribute_definition_integer.Name)
}

// commit attribute_definition_integer to the back repo (if it is already staged)
func (attribute_definition_integer *ATTRIBUTE_DEFINITION_INTEGER) Commit(stage *StageStruct) *ATTRIBUTE_DEFINITION_INTEGER {
	if _, ok := stage.ATTRIBUTE_DEFINITION_INTEGERs[attribute_definition_integer]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitATTRIBUTE_DEFINITION_INTEGER(attribute_definition_integer)
		}
	}
	return attribute_definition_integer
}

func (attribute_definition_integer *ATTRIBUTE_DEFINITION_INTEGER) CommitVoid(stage *StageStruct) {
	attribute_definition_integer.Commit(stage)
}

// Checkout attribute_definition_integer to the back repo (if it is already staged)
func (attribute_definition_integer *ATTRIBUTE_DEFINITION_INTEGER) Checkout(stage *StageStruct) *ATTRIBUTE_DEFINITION_INTEGER {
	if _, ok := stage.ATTRIBUTE_DEFINITION_INTEGERs[attribute_definition_integer]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutATTRIBUTE_DEFINITION_INTEGER(attribute_definition_integer)
		}
	}
	return attribute_definition_integer
}

// for satisfaction of GongStruct interface
func (attribute_definition_integer *ATTRIBUTE_DEFINITION_INTEGER) GetName() (res string) {
	return attribute_definition_integer.Name
}

// Stage puts attribute_definition_real to the model stage
func (attribute_definition_real *ATTRIBUTE_DEFINITION_REAL) Stage(stage *StageStruct) *ATTRIBUTE_DEFINITION_REAL {
	stage.ATTRIBUTE_DEFINITION_REALs[attribute_definition_real] = __member
	stage.ATTRIBUTE_DEFINITION_REALs_mapString[attribute_definition_real.Name] = attribute_definition_real

	return attribute_definition_real
}

// Unstage removes attribute_definition_real off the model stage
func (attribute_definition_real *ATTRIBUTE_DEFINITION_REAL) Unstage(stage *StageStruct) *ATTRIBUTE_DEFINITION_REAL {
	delete(stage.ATTRIBUTE_DEFINITION_REALs, attribute_definition_real)
	delete(stage.ATTRIBUTE_DEFINITION_REALs_mapString, attribute_definition_real.Name)
	return attribute_definition_real
}

// UnstageVoid removes attribute_definition_real off the model stage
func (attribute_definition_real *ATTRIBUTE_DEFINITION_REAL) UnstageVoid(stage *StageStruct) {
	delete(stage.ATTRIBUTE_DEFINITION_REALs, attribute_definition_real)
	delete(stage.ATTRIBUTE_DEFINITION_REALs_mapString, attribute_definition_real.Name)
}

// commit attribute_definition_real to the back repo (if it is already staged)
func (attribute_definition_real *ATTRIBUTE_DEFINITION_REAL) Commit(stage *StageStruct) *ATTRIBUTE_DEFINITION_REAL {
	if _, ok := stage.ATTRIBUTE_DEFINITION_REALs[attribute_definition_real]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitATTRIBUTE_DEFINITION_REAL(attribute_definition_real)
		}
	}
	return attribute_definition_real
}

func (attribute_definition_real *ATTRIBUTE_DEFINITION_REAL) CommitVoid(stage *StageStruct) {
	attribute_definition_real.Commit(stage)
}

// Checkout attribute_definition_real to the back repo (if it is already staged)
func (attribute_definition_real *ATTRIBUTE_DEFINITION_REAL) Checkout(stage *StageStruct) *ATTRIBUTE_DEFINITION_REAL {
	if _, ok := stage.ATTRIBUTE_DEFINITION_REALs[attribute_definition_real]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutATTRIBUTE_DEFINITION_REAL(attribute_definition_real)
		}
	}
	return attribute_definition_real
}

// for satisfaction of GongStruct interface
func (attribute_definition_real *ATTRIBUTE_DEFINITION_REAL) GetName() (res string) {
	return attribute_definition_real.Name
}

// Stage puts attribute_definition_string to the model stage
func (attribute_definition_string *ATTRIBUTE_DEFINITION_STRING) Stage(stage *StageStruct) *ATTRIBUTE_DEFINITION_STRING {
	stage.ATTRIBUTE_DEFINITION_STRINGs[attribute_definition_string] = __member
	stage.ATTRIBUTE_DEFINITION_STRINGs_mapString[attribute_definition_string.Name] = attribute_definition_string

	return attribute_definition_string
}

// Unstage removes attribute_definition_string off the model stage
func (attribute_definition_string *ATTRIBUTE_DEFINITION_STRING) Unstage(stage *StageStruct) *ATTRIBUTE_DEFINITION_STRING {
	delete(stage.ATTRIBUTE_DEFINITION_STRINGs, attribute_definition_string)
	delete(stage.ATTRIBUTE_DEFINITION_STRINGs_mapString, attribute_definition_string.Name)
	return attribute_definition_string
}

// UnstageVoid removes attribute_definition_string off the model stage
func (attribute_definition_string *ATTRIBUTE_DEFINITION_STRING) UnstageVoid(stage *StageStruct) {
	delete(stage.ATTRIBUTE_DEFINITION_STRINGs, attribute_definition_string)
	delete(stage.ATTRIBUTE_DEFINITION_STRINGs_mapString, attribute_definition_string.Name)
}

// commit attribute_definition_string to the back repo (if it is already staged)
func (attribute_definition_string *ATTRIBUTE_DEFINITION_STRING) Commit(stage *StageStruct) *ATTRIBUTE_DEFINITION_STRING {
	if _, ok := stage.ATTRIBUTE_DEFINITION_STRINGs[attribute_definition_string]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitATTRIBUTE_DEFINITION_STRING(attribute_definition_string)
		}
	}
	return attribute_definition_string
}

func (attribute_definition_string *ATTRIBUTE_DEFINITION_STRING) CommitVoid(stage *StageStruct) {
	attribute_definition_string.Commit(stage)
}

// Checkout attribute_definition_string to the back repo (if it is already staged)
func (attribute_definition_string *ATTRIBUTE_DEFINITION_STRING) Checkout(stage *StageStruct) *ATTRIBUTE_DEFINITION_STRING {
	if _, ok := stage.ATTRIBUTE_DEFINITION_STRINGs[attribute_definition_string]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutATTRIBUTE_DEFINITION_STRING(attribute_definition_string)
		}
	}
	return attribute_definition_string
}

// for satisfaction of GongStruct interface
func (attribute_definition_string *ATTRIBUTE_DEFINITION_STRING) GetName() (res string) {
	return attribute_definition_string.Name
}

// Stage puts attribute_definition_xhtml to the model stage
func (attribute_definition_xhtml *ATTRIBUTE_DEFINITION_XHTML) Stage(stage *StageStruct) *ATTRIBUTE_DEFINITION_XHTML {
	stage.ATTRIBUTE_DEFINITION_XHTMLs[attribute_definition_xhtml] = __member
	stage.ATTRIBUTE_DEFINITION_XHTMLs_mapString[attribute_definition_xhtml.Name] = attribute_definition_xhtml

	return attribute_definition_xhtml
}

// Unstage removes attribute_definition_xhtml off the model stage
func (attribute_definition_xhtml *ATTRIBUTE_DEFINITION_XHTML) Unstage(stage *StageStruct) *ATTRIBUTE_DEFINITION_XHTML {
	delete(stage.ATTRIBUTE_DEFINITION_XHTMLs, attribute_definition_xhtml)
	delete(stage.ATTRIBUTE_DEFINITION_XHTMLs_mapString, attribute_definition_xhtml.Name)
	return attribute_definition_xhtml
}

// UnstageVoid removes attribute_definition_xhtml off the model stage
func (attribute_definition_xhtml *ATTRIBUTE_DEFINITION_XHTML) UnstageVoid(stage *StageStruct) {
	delete(stage.ATTRIBUTE_DEFINITION_XHTMLs, attribute_definition_xhtml)
	delete(stage.ATTRIBUTE_DEFINITION_XHTMLs_mapString, attribute_definition_xhtml.Name)
}

// commit attribute_definition_xhtml to the back repo (if it is already staged)
func (attribute_definition_xhtml *ATTRIBUTE_DEFINITION_XHTML) Commit(stage *StageStruct) *ATTRIBUTE_DEFINITION_XHTML {
	if _, ok := stage.ATTRIBUTE_DEFINITION_XHTMLs[attribute_definition_xhtml]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitATTRIBUTE_DEFINITION_XHTML(attribute_definition_xhtml)
		}
	}
	return attribute_definition_xhtml
}

func (attribute_definition_xhtml *ATTRIBUTE_DEFINITION_XHTML) CommitVoid(stage *StageStruct) {
	attribute_definition_xhtml.Commit(stage)
}

// Checkout attribute_definition_xhtml to the back repo (if it is already staged)
func (attribute_definition_xhtml *ATTRIBUTE_DEFINITION_XHTML) Checkout(stage *StageStruct) *ATTRIBUTE_DEFINITION_XHTML {
	if _, ok := stage.ATTRIBUTE_DEFINITION_XHTMLs[attribute_definition_xhtml]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutATTRIBUTE_DEFINITION_XHTML(attribute_definition_xhtml)
		}
	}
	return attribute_definition_xhtml
}

// for satisfaction of GongStruct interface
func (attribute_definition_xhtml *ATTRIBUTE_DEFINITION_XHTML) GetName() (res string) {
	return attribute_definition_xhtml.Name
}

// Stage puts attribute_value_boolean to the model stage
func (attribute_value_boolean *ATTRIBUTE_VALUE_BOOLEAN) Stage(stage *StageStruct) *ATTRIBUTE_VALUE_BOOLEAN {
	stage.ATTRIBUTE_VALUE_BOOLEANs[attribute_value_boolean] = __member
	stage.ATTRIBUTE_VALUE_BOOLEANs_mapString[attribute_value_boolean.Name] = attribute_value_boolean

	return attribute_value_boolean
}

// Unstage removes attribute_value_boolean off the model stage
func (attribute_value_boolean *ATTRIBUTE_VALUE_BOOLEAN) Unstage(stage *StageStruct) *ATTRIBUTE_VALUE_BOOLEAN {
	delete(stage.ATTRIBUTE_VALUE_BOOLEANs, attribute_value_boolean)
	delete(stage.ATTRIBUTE_VALUE_BOOLEANs_mapString, attribute_value_boolean.Name)
	return attribute_value_boolean
}

// UnstageVoid removes attribute_value_boolean off the model stage
func (attribute_value_boolean *ATTRIBUTE_VALUE_BOOLEAN) UnstageVoid(stage *StageStruct) {
	delete(stage.ATTRIBUTE_VALUE_BOOLEANs, attribute_value_boolean)
	delete(stage.ATTRIBUTE_VALUE_BOOLEANs_mapString, attribute_value_boolean.Name)
}

// commit attribute_value_boolean to the back repo (if it is already staged)
func (attribute_value_boolean *ATTRIBUTE_VALUE_BOOLEAN) Commit(stage *StageStruct) *ATTRIBUTE_VALUE_BOOLEAN {
	if _, ok := stage.ATTRIBUTE_VALUE_BOOLEANs[attribute_value_boolean]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitATTRIBUTE_VALUE_BOOLEAN(attribute_value_boolean)
		}
	}
	return attribute_value_boolean
}

func (attribute_value_boolean *ATTRIBUTE_VALUE_BOOLEAN) CommitVoid(stage *StageStruct) {
	attribute_value_boolean.Commit(stage)
}

// Checkout attribute_value_boolean to the back repo (if it is already staged)
func (attribute_value_boolean *ATTRIBUTE_VALUE_BOOLEAN) Checkout(stage *StageStruct) *ATTRIBUTE_VALUE_BOOLEAN {
	if _, ok := stage.ATTRIBUTE_VALUE_BOOLEANs[attribute_value_boolean]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutATTRIBUTE_VALUE_BOOLEAN(attribute_value_boolean)
		}
	}
	return attribute_value_boolean
}

// for satisfaction of GongStruct interface
func (attribute_value_boolean *ATTRIBUTE_VALUE_BOOLEAN) GetName() (res string) {
	return attribute_value_boolean.Name
}

// Stage puts attribute_value_date to the model stage
func (attribute_value_date *ATTRIBUTE_VALUE_DATE) Stage(stage *StageStruct) *ATTRIBUTE_VALUE_DATE {
	stage.ATTRIBUTE_VALUE_DATEs[attribute_value_date] = __member
	stage.ATTRIBUTE_VALUE_DATEs_mapString[attribute_value_date.Name] = attribute_value_date

	return attribute_value_date
}

// Unstage removes attribute_value_date off the model stage
func (attribute_value_date *ATTRIBUTE_VALUE_DATE) Unstage(stage *StageStruct) *ATTRIBUTE_VALUE_DATE {
	delete(stage.ATTRIBUTE_VALUE_DATEs, attribute_value_date)
	delete(stage.ATTRIBUTE_VALUE_DATEs_mapString, attribute_value_date.Name)
	return attribute_value_date
}

// UnstageVoid removes attribute_value_date off the model stage
func (attribute_value_date *ATTRIBUTE_VALUE_DATE) UnstageVoid(stage *StageStruct) {
	delete(stage.ATTRIBUTE_VALUE_DATEs, attribute_value_date)
	delete(stage.ATTRIBUTE_VALUE_DATEs_mapString, attribute_value_date.Name)
}

// commit attribute_value_date to the back repo (if it is already staged)
func (attribute_value_date *ATTRIBUTE_VALUE_DATE) Commit(stage *StageStruct) *ATTRIBUTE_VALUE_DATE {
	if _, ok := stage.ATTRIBUTE_VALUE_DATEs[attribute_value_date]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitATTRIBUTE_VALUE_DATE(attribute_value_date)
		}
	}
	return attribute_value_date
}

func (attribute_value_date *ATTRIBUTE_VALUE_DATE) CommitVoid(stage *StageStruct) {
	attribute_value_date.Commit(stage)
}

// Checkout attribute_value_date to the back repo (if it is already staged)
func (attribute_value_date *ATTRIBUTE_VALUE_DATE) Checkout(stage *StageStruct) *ATTRIBUTE_VALUE_DATE {
	if _, ok := stage.ATTRIBUTE_VALUE_DATEs[attribute_value_date]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutATTRIBUTE_VALUE_DATE(attribute_value_date)
		}
	}
	return attribute_value_date
}

// for satisfaction of GongStruct interface
func (attribute_value_date *ATTRIBUTE_VALUE_DATE) GetName() (res string) {
	return attribute_value_date.Name
}

// Stage puts attribute_value_enumeration to the model stage
func (attribute_value_enumeration *ATTRIBUTE_VALUE_ENUMERATION) Stage(stage *StageStruct) *ATTRIBUTE_VALUE_ENUMERATION {
	stage.ATTRIBUTE_VALUE_ENUMERATIONs[attribute_value_enumeration] = __member
	stage.ATTRIBUTE_VALUE_ENUMERATIONs_mapString[attribute_value_enumeration.Name] = attribute_value_enumeration

	return attribute_value_enumeration
}

// Unstage removes attribute_value_enumeration off the model stage
func (attribute_value_enumeration *ATTRIBUTE_VALUE_ENUMERATION) Unstage(stage *StageStruct) *ATTRIBUTE_VALUE_ENUMERATION {
	delete(stage.ATTRIBUTE_VALUE_ENUMERATIONs, attribute_value_enumeration)
	delete(stage.ATTRIBUTE_VALUE_ENUMERATIONs_mapString, attribute_value_enumeration.Name)
	return attribute_value_enumeration
}

// UnstageVoid removes attribute_value_enumeration off the model stage
func (attribute_value_enumeration *ATTRIBUTE_VALUE_ENUMERATION) UnstageVoid(stage *StageStruct) {
	delete(stage.ATTRIBUTE_VALUE_ENUMERATIONs, attribute_value_enumeration)
	delete(stage.ATTRIBUTE_VALUE_ENUMERATIONs_mapString, attribute_value_enumeration.Name)
}

// commit attribute_value_enumeration to the back repo (if it is already staged)
func (attribute_value_enumeration *ATTRIBUTE_VALUE_ENUMERATION) Commit(stage *StageStruct) *ATTRIBUTE_VALUE_ENUMERATION {
	if _, ok := stage.ATTRIBUTE_VALUE_ENUMERATIONs[attribute_value_enumeration]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitATTRIBUTE_VALUE_ENUMERATION(attribute_value_enumeration)
		}
	}
	return attribute_value_enumeration
}

func (attribute_value_enumeration *ATTRIBUTE_VALUE_ENUMERATION) CommitVoid(stage *StageStruct) {
	attribute_value_enumeration.Commit(stage)
}

// Checkout attribute_value_enumeration to the back repo (if it is already staged)
func (attribute_value_enumeration *ATTRIBUTE_VALUE_ENUMERATION) Checkout(stage *StageStruct) *ATTRIBUTE_VALUE_ENUMERATION {
	if _, ok := stage.ATTRIBUTE_VALUE_ENUMERATIONs[attribute_value_enumeration]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutATTRIBUTE_VALUE_ENUMERATION(attribute_value_enumeration)
		}
	}
	return attribute_value_enumeration
}

// for satisfaction of GongStruct interface
func (attribute_value_enumeration *ATTRIBUTE_VALUE_ENUMERATION) GetName() (res string) {
	return attribute_value_enumeration.Name
}

// Stage puts attribute_value_integer to the model stage
func (attribute_value_integer *ATTRIBUTE_VALUE_INTEGER) Stage(stage *StageStruct) *ATTRIBUTE_VALUE_INTEGER {
	stage.ATTRIBUTE_VALUE_INTEGERs[attribute_value_integer] = __member
	stage.ATTRIBUTE_VALUE_INTEGERs_mapString[attribute_value_integer.Name] = attribute_value_integer

	return attribute_value_integer
}

// Unstage removes attribute_value_integer off the model stage
func (attribute_value_integer *ATTRIBUTE_VALUE_INTEGER) Unstage(stage *StageStruct) *ATTRIBUTE_VALUE_INTEGER {
	delete(stage.ATTRIBUTE_VALUE_INTEGERs, attribute_value_integer)
	delete(stage.ATTRIBUTE_VALUE_INTEGERs_mapString, attribute_value_integer.Name)
	return attribute_value_integer
}

// UnstageVoid removes attribute_value_integer off the model stage
func (attribute_value_integer *ATTRIBUTE_VALUE_INTEGER) UnstageVoid(stage *StageStruct) {
	delete(stage.ATTRIBUTE_VALUE_INTEGERs, attribute_value_integer)
	delete(stage.ATTRIBUTE_VALUE_INTEGERs_mapString, attribute_value_integer.Name)
}

// commit attribute_value_integer to the back repo (if it is already staged)
func (attribute_value_integer *ATTRIBUTE_VALUE_INTEGER) Commit(stage *StageStruct) *ATTRIBUTE_VALUE_INTEGER {
	if _, ok := stage.ATTRIBUTE_VALUE_INTEGERs[attribute_value_integer]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitATTRIBUTE_VALUE_INTEGER(attribute_value_integer)
		}
	}
	return attribute_value_integer
}

func (attribute_value_integer *ATTRIBUTE_VALUE_INTEGER) CommitVoid(stage *StageStruct) {
	attribute_value_integer.Commit(stage)
}

// Checkout attribute_value_integer to the back repo (if it is already staged)
func (attribute_value_integer *ATTRIBUTE_VALUE_INTEGER) Checkout(stage *StageStruct) *ATTRIBUTE_VALUE_INTEGER {
	if _, ok := stage.ATTRIBUTE_VALUE_INTEGERs[attribute_value_integer]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutATTRIBUTE_VALUE_INTEGER(attribute_value_integer)
		}
	}
	return attribute_value_integer
}

// for satisfaction of GongStruct interface
func (attribute_value_integer *ATTRIBUTE_VALUE_INTEGER) GetName() (res string) {
	return attribute_value_integer.Name
}

// Stage puts attribute_value_real to the model stage
func (attribute_value_real *ATTRIBUTE_VALUE_REAL) Stage(stage *StageStruct) *ATTRIBUTE_VALUE_REAL {
	stage.ATTRIBUTE_VALUE_REALs[attribute_value_real] = __member
	stage.ATTRIBUTE_VALUE_REALs_mapString[attribute_value_real.Name] = attribute_value_real

	return attribute_value_real
}

// Unstage removes attribute_value_real off the model stage
func (attribute_value_real *ATTRIBUTE_VALUE_REAL) Unstage(stage *StageStruct) *ATTRIBUTE_VALUE_REAL {
	delete(stage.ATTRIBUTE_VALUE_REALs, attribute_value_real)
	delete(stage.ATTRIBUTE_VALUE_REALs_mapString, attribute_value_real.Name)
	return attribute_value_real
}

// UnstageVoid removes attribute_value_real off the model stage
func (attribute_value_real *ATTRIBUTE_VALUE_REAL) UnstageVoid(stage *StageStruct) {
	delete(stage.ATTRIBUTE_VALUE_REALs, attribute_value_real)
	delete(stage.ATTRIBUTE_VALUE_REALs_mapString, attribute_value_real.Name)
}

// commit attribute_value_real to the back repo (if it is already staged)
func (attribute_value_real *ATTRIBUTE_VALUE_REAL) Commit(stage *StageStruct) *ATTRIBUTE_VALUE_REAL {
	if _, ok := stage.ATTRIBUTE_VALUE_REALs[attribute_value_real]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitATTRIBUTE_VALUE_REAL(attribute_value_real)
		}
	}
	return attribute_value_real
}

func (attribute_value_real *ATTRIBUTE_VALUE_REAL) CommitVoid(stage *StageStruct) {
	attribute_value_real.Commit(stage)
}

// Checkout attribute_value_real to the back repo (if it is already staged)
func (attribute_value_real *ATTRIBUTE_VALUE_REAL) Checkout(stage *StageStruct) *ATTRIBUTE_VALUE_REAL {
	if _, ok := stage.ATTRIBUTE_VALUE_REALs[attribute_value_real]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutATTRIBUTE_VALUE_REAL(attribute_value_real)
		}
	}
	return attribute_value_real
}

// for satisfaction of GongStruct interface
func (attribute_value_real *ATTRIBUTE_VALUE_REAL) GetName() (res string) {
	return attribute_value_real.Name
}

// Stage puts attribute_value_string to the model stage
func (attribute_value_string *ATTRIBUTE_VALUE_STRING) Stage(stage *StageStruct) *ATTRIBUTE_VALUE_STRING {
	stage.ATTRIBUTE_VALUE_STRINGs[attribute_value_string] = __member
	stage.ATTRIBUTE_VALUE_STRINGs_mapString[attribute_value_string.Name] = attribute_value_string

	return attribute_value_string
}

// Unstage removes attribute_value_string off the model stage
func (attribute_value_string *ATTRIBUTE_VALUE_STRING) Unstage(stage *StageStruct) *ATTRIBUTE_VALUE_STRING {
	delete(stage.ATTRIBUTE_VALUE_STRINGs, attribute_value_string)
	delete(stage.ATTRIBUTE_VALUE_STRINGs_mapString, attribute_value_string.Name)
	return attribute_value_string
}

// UnstageVoid removes attribute_value_string off the model stage
func (attribute_value_string *ATTRIBUTE_VALUE_STRING) UnstageVoid(stage *StageStruct) {
	delete(stage.ATTRIBUTE_VALUE_STRINGs, attribute_value_string)
	delete(stage.ATTRIBUTE_VALUE_STRINGs_mapString, attribute_value_string.Name)
}

// commit attribute_value_string to the back repo (if it is already staged)
func (attribute_value_string *ATTRIBUTE_VALUE_STRING) Commit(stage *StageStruct) *ATTRIBUTE_VALUE_STRING {
	if _, ok := stage.ATTRIBUTE_VALUE_STRINGs[attribute_value_string]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitATTRIBUTE_VALUE_STRING(attribute_value_string)
		}
	}
	return attribute_value_string
}

func (attribute_value_string *ATTRIBUTE_VALUE_STRING) CommitVoid(stage *StageStruct) {
	attribute_value_string.Commit(stage)
}

// Checkout attribute_value_string to the back repo (if it is already staged)
func (attribute_value_string *ATTRIBUTE_VALUE_STRING) Checkout(stage *StageStruct) *ATTRIBUTE_VALUE_STRING {
	if _, ok := stage.ATTRIBUTE_VALUE_STRINGs[attribute_value_string]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutATTRIBUTE_VALUE_STRING(attribute_value_string)
		}
	}
	return attribute_value_string
}

// for satisfaction of GongStruct interface
func (attribute_value_string *ATTRIBUTE_VALUE_STRING) GetName() (res string) {
	return attribute_value_string.Name
}

// Stage puts attribute_value_xhtml to the model stage
func (attribute_value_xhtml *ATTRIBUTE_VALUE_XHTML) Stage(stage *StageStruct) *ATTRIBUTE_VALUE_XHTML {
	stage.ATTRIBUTE_VALUE_XHTMLs[attribute_value_xhtml] = __member
	stage.ATTRIBUTE_VALUE_XHTMLs_mapString[attribute_value_xhtml.Name] = attribute_value_xhtml

	return attribute_value_xhtml
}

// Unstage removes attribute_value_xhtml off the model stage
func (attribute_value_xhtml *ATTRIBUTE_VALUE_XHTML) Unstage(stage *StageStruct) *ATTRIBUTE_VALUE_XHTML {
	delete(stage.ATTRIBUTE_VALUE_XHTMLs, attribute_value_xhtml)
	delete(stage.ATTRIBUTE_VALUE_XHTMLs_mapString, attribute_value_xhtml.Name)
	return attribute_value_xhtml
}

// UnstageVoid removes attribute_value_xhtml off the model stage
func (attribute_value_xhtml *ATTRIBUTE_VALUE_XHTML) UnstageVoid(stage *StageStruct) {
	delete(stage.ATTRIBUTE_VALUE_XHTMLs, attribute_value_xhtml)
	delete(stage.ATTRIBUTE_VALUE_XHTMLs_mapString, attribute_value_xhtml.Name)
}

// commit attribute_value_xhtml to the back repo (if it is already staged)
func (attribute_value_xhtml *ATTRIBUTE_VALUE_XHTML) Commit(stage *StageStruct) *ATTRIBUTE_VALUE_XHTML {
	if _, ok := stage.ATTRIBUTE_VALUE_XHTMLs[attribute_value_xhtml]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitATTRIBUTE_VALUE_XHTML(attribute_value_xhtml)
		}
	}
	return attribute_value_xhtml
}

func (attribute_value_xhtml *ATTRIBUTE_VALUE_XHTML) CommitVoid(stage *StageStruct) {
	attribute_value_xhtml.Commit(stage)
}

// Checkout attribute_value_xhtml to the back repo (if it is already staged)
func (attribute_value_xhtml *ATTRIBUTE_VALUE_XHTML) Checkout(stage *StageStruct) *ATTRIBUTE_VALUE_XHTML {
	if _, ok := stage.ATTRIBUTE_VALUE_XHTMLs[attribute_value_xhtml]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutATTRIBUTE_VALUE_XHTML(attribute_value_xhtml)
		}
	}
	return attribute_value_xhtml
}

// for satisfaction of GongStruct interface
func (attribute_value_xhtml *ATTRIBUTE_VALUE_XHTML) GetName() (res string) {
	return attribute_value_xhtml.Name
}

// Stage puts datatype_definition_boolean to the model stage
func (datatype_definition_boolean *DATATYPE_DEFINITION_BOOLEAN) Stage(stage *StageStruct) *DATATYPE_DEFINITION_BOOLEAN {
	stage.DATATYPE_DEFINITION_BOOLEANs[datatype_definition_boolean] = __member
	stage.DATATYPE_DEFINITION_BOOLEANs_mapString[datatype_definition_boolean.Name] = datatype_definition_boolean

	return datatype_definition_boolean
}

// Unstage removes datatype_definition_boolean off the model stage
func (datatype_definition_boolean *DATATYPE_DEFINITION_BOOLEAN) Unstage(stage *StageStruct) *DATATYPE_DEFINITION_BOOLEAN {
	delete(stage.DATATYPE_DEFINITION_BOOLEANs, datatype_definition_boolean)
	delete(stage.DATATYPE_DEFINITION_BOOLEANs_mapString, datatype_definition_boolean.Name)
	return datatype_definition_boolean
}

// UnstageVoid removes datatype_definition_boolean off the model stage
func (datatype_definition_boolean *DATATYPE_DEFINITION_BOOLEAN) UnstageVoid(stage *StageStruct) {
	delete(stage.DATATYPE_DEFINITION_BOOLEANs, datatype_definition_boolean)
	delete(stage.DATATYPE_DEFINITION_BOOLEANs_mapString, datatype_definition_boolean.Name)
}

// commit datatype_definition_boolean to the back repo (if it is already staged)
func (datatype_definition_boolean *DATATYPE_DEFINITION_BOOLEAN) Commit(stage *StageStruct) *DATATYPE_DEFINITION_BOOLEAN {
	if _, ok := stage.DATATYPE_DEFINITION_BOOLEANs[datatype_definition_boolean]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitDATATYPE_DEFINITION_BOOLEAN(datatype_definition_boolean)
		}
	}
	return datatype_definition_boolean
}

func (datatype_definition_boolean *DATATYPE_DEFINITION_BOOLEAN) CommitVoid(stage *StageStruct) {
	datatype_definition_boolean.Commit(stage)
}

// Checkout datatype_definition_boolean to the back repo (if it is already staged)
func (datatype_definition_boolean *DATATYPE_DEFINITION_BOOLEAN) Checkout(stage *StageStruct) *DATATYPE_DEFINITION_BOOLEAN {
	if _, ok := stage.DATATYPE_DEFINITION_BOOLEANs[datatype_definition_boolean]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutDATATYPE_DEFINITION_BOOLEAN(datatype_definition_boolean)
		}
	}
	return datatype_definition_boolean
}

// for satisfaction of GongStruct interface
func (datatype_definition_boolean *DATATYPE_DEFINITION_BOOLEAN) GetName() (res string) {
	return datatype_definition_boolean.Name
}

// Stage puts datatype_definition_date to the model stage
func (datatype_definition_date *DATATYPE_DEFINITION_DATE) Stage(stage *StageStruct) *DATATYPE_DEFINITION_DATE {
	stage.DATATYPE_DEFINITION_DATEs[datatype_definition_date] = __member
	stage.DATATYPE_DEFINITION_DATEs_mapString[datatype_definition_date.Name] = datatype_definition_date

	return datatype_definition_date
}

// Unstage removes datatype_definition_date off the model stage
func (datatype_definition_date *DATATYPE_DEFINITION_DATE) Unstage(stage *StageStruct) *DATATYPE_DEFINITION_DATE {
	delete(stage.DATATYPE_DEFINITION_DATEs, datatype_definition_date)
	delete(stage.DATATYPE_DEFINITION_DATEs_mapString, datatype_definition_date.Name)
	return datatype_definition_date
}

// UnstageVoid removes datatype_definition_date off the model stage
func (datatype_definition_date *DATATYPE_DEFINITION_DATE) UnstageVoid(stage *StageStruct) {
	delete(stage.DATATYPE_DEFINITION_DATEs, datatype_definition_date)
	delete(stage.DATATYPE_DEFINITION_DATEs_mapString, datatype_definition_date.Name)
}

// commit datatype_definition_date to the back repo (if it is already staged)
func (datatype_definition_date *DATATYPE_DEFINITION_DATE) Commit(stage *StageStruct) *DATATYPE_DEFINITION_DATE {
	if _, ok := stage.DATATYPE_DEFINITION_DATEs[datatype_definition_date]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitDATATYPE_DEFINITION_DATE(datatype_definition_date)
		}
	}
	return datatype_definition_date
}

func (datatype_definition_date *DATATYPE_DEFINITION_DATE) CommitVoid(stage *StageStruct) {
	datatype_definition_date.Commit(stage)
}

// Checkout datatype_definition_date to the back repo (if it is already staged)
func (datatype_definition_date *DATATYPE_DEFINITION_DATE) Checkout(stage *StageStruct) *DATATYPE_DEFINITION_DATE {
	if _, ok := stage.DATATYPE_DEFINITION_DATEs[datatype_definition_date]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutDATATYPE_DEFINITION_DATE(datatype_definition_date)
		}
	}
	return datatype_definition_date
}

// for satisfaction of GongStruct interface
func (datatype_definition_date *DATATYPE_DEFINITION_DATE) GetName() (res string) {
	return datatype_definition_date.Name
}

// Stage puts datatype_definition_enumeration to the model stage
func (datatype_definition_enumeration *DATATYPE_DEFINITION_ENUMERATION) Stage(stage *StageStruct) *DATATYPE_DEFINITION_ENUMERATION {
	stage.DATATYPE_DEFINITION_ENUMERATIONs[datatype_definition_enumeration] = __member
	stage.DATATYPE_DEFINITION_ENUMERATIONs_mapString[datatype_definition_enumeration.Name] = datatype_definition_enumeration

	return datatype_definition_enumeration
}

// Unstage removes datatype_definition_enumeration off the model stage
func (datatype_definition_enumeration *DATATYPE_DEFINITION_ENUMERATION) Unstage(stage *StageStruct) *DATATYPE_DEFINITION_ENUMERATION {
	delete(stage.DATATYPE_DEFINITION_ENUMERATIONs, datatype_definition_enumeration)
	delete(stage.DATATYPE_DEFINITION_ENUMERATIONs_mapString, datatype_definition_enumeration.Name)
	return datatype_definition_enumeration
}

// UnstageVoid removes datatype_definition_enumeration off the model stage
func (datatype_definition_enumeration *DATATYPE_DEFINITION_ENUMERATION) UnstageVoid(stage *StageStruct) {
	delete(stage.DATATYPE_DEFINITION_ENUMERATIONs, datatype_definition_enumeration)
	delete(stage.DATATYPE_DEFINITION_ENUMERATIONs_mapString, datatype_definition_enumeration.Name)
}

// commit datatype_definition_enumeration to the back repo (if it is already staged)
func (datatype_definition_enumeration *DATATYPE_DEFINITION_ENUMERATION) Commit(stage *StageStruct) *DATATYPE_DEFINITION_ENUMERATION {
	if _, ok := stage.DATATYPE_DEFINITION_ENUMERATIONs[datatype_definition_enumeration]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitDATATYPE_DEFINITION_ENUMERATION(datatype_definition_enumeration)
		}
	}
	return datatype_definition_enumeration
}

func (datatype_definition_enumeration *DATATYPE_DEFINITION_ENUMERATION) CommitVoid(stage *StageStruct) {
	datatype_definition_enumeration.Commit(stage)
}

// Checkout datatype_definition_enumeration to the back repo (if it is already staged)
func (datatype_definition_enumeration *DATATYPE_DEFINITION_ENUMERATION) Checkout(stage *StageStruct) *DATATYPE_DEFINITION_ENUMERATION {
	if _, ok := stage.DATATYPE_DEFINITION_ENUMERATIONs[datatype_definition_enumeration]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutDATATYPE_DEFINITION_ENUMERATION(datatype_definition_enumeration)
		}
	}
	return datatype_definition_enumeration
}

// for satisfaction of GongStruct interface
func (datatype_definition_enumeration *DATATYPE_DEFINITION_ENUMERATION) GetName() (res string) {
	return datatype_definition_enumeration.Name
}

// Stage puts datatype_definition_integer to the model stage
func (datatype_definition_integer *DATATYPE_DEFINITION_INTEGER) Stage(stage *StageStruct) *DATATYPE_DEFINITION_INTEGER {
	stage.DATATYPE_DEFINITION_INTEGERs[datatype_definition_integer] = __member
	stage.DATATYPE_DEFINITION_INTEGERs_mapString[datatype_definition_integer.Name] = datatype_definition_integer

	return datatype_definition_integer
}

// Unstage removes datatype_definition_integer off the model stage
func (datatype_definition_integer *DATATYPE_DEFINITION_INTEGER) Unstage(stage *StageStruct) *DATATYPE_DEFINITION_INTEGER {
	delete(stage.DATATYPE_DEFINITION_INTEGERs, datatype_definition_integer)
	delete(stage.DATATYPE_DEFINITION_INTEGERs_mapString, datatype_definition_integer.Name)
	return datatype_definition_integer
}

// UnstageVoid removes datatype_definition_integer off the model stage
func (datatype_definition_integer *DATATYPE_DEFINITION_INTEGER) UnstageVoid(stage *StageStruct) {
	delete(stage.DATATYPE_DEFINITION_INTEGERs, datatype_definition_integer)
	delete(stage.DATATYPE_DEFINITION_INTEGERs_mapString, datatype_definition_integer.Name)
}

// commit datatype_definition_integer to the back repo (if it is already staged)
func (datatype_definition_integer *DATATYPE_DEFINITION_INTEGER) Commit(stage *StageStruct) *DATATYPE_DEFINITION_INTEGER {
	if _, ok := stage.DATATYPE_DEFINITION_INTEGERs[datatype_definition_integer]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitDATATYPE_DEFINITION_INTEGER(datatype_definition_integer)
		}
	}
	return datatype_definition_integer
}

func (datatype_definition_integer *DATATYPE_DEFINITION_INTEGER) CommitVoid(stage *StageStruct) {
	datatype_definition_integer.Commit(stage)
}

// Checkout datatype_definition_integer to the back repo (if it is already staged)
func (datatype_definition_integer *DATATYPE_DEFINITION_INTEGER) Checkout(stage *StageStruct) *DATATYPE_DEFINITION_INTEGER {
	if _, ok := stage.DATATYPE_DEFINITION_INTEGERs[datatype_definition_integer]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutDATATYPE_DEFINITION_INTEGER(datatype_definition_integer)
		}
	}
	return datatype_definition_integer
}

// for satisfaction of GongStruct interface
func (datatype_definition_integer *DATATYPE_DEFINITION_INTEGER) GetName() (res string) {
	return datatype_definition_integer.Name
}

// Stage puts datatype_definition_real to the model stage
func (datatype_definition_real *DATATYPE_DEFINITION_REAL) Stage(stage *StageStruct) *DATATYPE_DEFINITION_REAL {
	stage.DATATYPE_DEFINITION_REALs[datatype_definition_real] = __member
	stage.DATATYPE_DEFINITION_REALs_mapString[datatype_definition_real.Name] = datatype_definition_real

	return datatype_definition_real
}

// Unstage removes datatype_definition_real off the model stage
func (datatype_definition_real *DATATYPE_DEFINITION_REAL) Unstage(stage *StageStruct) *DATATYPE_DEFINITION_REAL {
	delete(stage.DATATYPE_DEFINITION_REALs, datatype_definition_real)
	delete(stage.DATATYPE_DEFINITION_REALs_mapString, datatype_definition_real.Name)
	return datatype_definition_real
}

// UnstageVoid removes datatype_definition_real off the model stage
func (datatype_definition_real *DATATYPE_DEFINITION_REAL) UnstageVoid(stage *StageStruct) {
	delete(stage.DATATYPE_DEFINITION_REALs, datatype_definition_real)
	delete(stage.DATATYPE_DEFINITION_REALs_mapString, datatype_definition_real.Name)
}

// commit datatype_definition_real to the back repo (if it is already staged)
func (datatype_definition_real *DATATYPE_DEFINITION_REAL) Commit(stage *StageStruct) *DATATYPE_DEFINITION_REAL {
	if _, ok := stage.DATATYPE_DEFINITION_REALs[datatype_definition_real]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitDATATYPE_DEFINITION_REAL(datatype_definition_real)
		}
	}
	return datatype_definition_real
}

func (datatype_definition_real *DATATYPE_DEFINITION_REAL) CommitVoid(stage *StageStruct) {
	datatype_definition_real.Commit(stage)
}

// Checkout datatype_definition_real to the back repo (if it is already staged)
func (datatype_definition_real *DATATYPE_DEFINITION_REAL) Checkout(stage *StageStruct) *DATATYPE_DEFINITION_REAL {
	if _, ok := stage.DATATYPE_DEFINITION_REALs[datatype_definition_real]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutDATATYPE_DEFINITION_REAL(datatype_definition_real)
		}
	}
	return datatype_definition_real
}

// for satisfaction of GongStruct interface
func (datatype_definition_real *DATATYPE_DEFINITION_REAL) GetName() (res string) {
	return datatype_definition_real.Name
}

// Stage puts datatype_definition_string to the model stage
func (datatype_definition_string *DATATYPE_DEFINITION_STRING) Stage(stage *StageStruct) *DATATYPE_DEFINITION_STRING {
	stage.DATATYPE_DEFINITION_STRINGs[datatype_definition_string] = __member
	stage.DATATYPE_DEFINITION_STRINGs_mapString[datatype_definition_string.Name] = datatype_definition_string

	return datatype_definition_string
}

// Unstage removes datatype_definition_string off the model stage
func (datatype_definition_string *DATATYPE_DEFINITION_STRING) Unstage(stage *StageStruct) *DATATYPE_DEFINITION_STRING {
	delete(stage.DATATYPE_DEFINITION_STRINGs, datatype_definition_string)
	delete(stage.DATATYPE_DEFINITION_STRINGs_mapString, datatype_definition_string.Name)
	return datatype_definition_string
}

// UnstageVoid removes datatype_definition_string off the model stage
func (datatype_definition_string *DATATYPE_DEFINITION_STRING) UnstageVoid(stage *StageStruct) {
	delete(stage.DATATYPE_DEFINITION_STRINGs, datatype_definition_string)
	delete(stage.DATATYPE_DEFINITION_STRINGs_mapString, datatype_definition_string.Name)
}

// commit datatype_definition_string to the back repo (if it is already staged)
func (datatype_definition_string *DATATYPE_DEFINITION_STRING) Commit(stage *StageStruct) *DATATYPE_DEFINITION_STRING {
	if _, ok := stage.DATATYPE_DEFINITION_STRINGs[datatype_definition_string]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitDATATYPE_DEFINITION_STRING(datatype_definition_string)
		}
	}
	return datatype_definition_string
}

func (datatype_definition_string *DATATYPE_DEFINITION_STRING) CommitVoid(stage *StageStruct) {
	datatype_definition_string.Commit(stage)
}

// Checkout datatype_definition_string to the back repo (if it is already staged)
func (datatype_definition_string *DATATYPE_DEFINITION_STRING) Checkout(stage *StageStruct) *DATATYPE_DEFINITION_STRING {
	if _, ok := stage.DATATYPE_DEFINITION_STRINGs[datatype_definition_string]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutDATATYPE_DEFINITION_STRING(datatype_definition_string)
		}
	}
	return datatype_definition_string
}

// for satisfaction of GongStruct interface
func (datatype_definition_string *DATATYPE_DEFINITION_STRING) GetName() (res string) {
	return datatype_definition_string.Name
}

// Stage puts datatype_definition_xhtml to the model stage
func (datatype_definition_xhtml *DATATYPE_DEFINITION_XHTML) Stage(stage *StageStruct) *DATATYPE_DEFINITION_XHTML {
	stage.DATATYPE_DEFINITION_XHTMLs[datatype_definition_xhtml] = __member
	stage.DATATYPE_DEFINITION_XHTMLs_mapString[datatype_definition_xhtml.Name] = datatype_definition_xhtml

	return datatype_definition_xhtml
}

// Unstage removes datatype_definition_xhtml off the model stage
func (datatype_definition_xhtml *DATATYPE_DEFINITION_XHTML) Unstage(stage *StageStruct) *DATATYPE_DEFINITION_XHTML {
	delete(stage.DATATYPE_DEFINITION_XHTMLs, datatype_definition_xhtml)
	delete(stage.DATATYPE_DEFINITION_XHTMLs_mapString, datatype_definition_xhtml.Name)
	return datatype_definition_xhtml
}

// UnstageVoid removes datatype_definition_xhtml off the model stage
func (datatype_definition_xhtml *DATATYPE_DEFINITION_XHTML) UnstageVoid(stage *StageStruct) {
	delete(stage.DATATYPE_DEFINITION_XHTMLs, datatype_definition_xhtml)
	delete(stage.DATATYPE_DEFINITION_XHTMLs_mapString, datatype_definition_xhtml.Name)
}

// commit datatype_definition_xhtml to the back repo (if it is already staged)
func (datatype_definition_xhtml *DATATYPE_DEFINITION_XHTML) Commit(stage *StageStruct) *DATATYPE_DEFINITION_XHTML {
	if _, ok := stage.DATATYPE_DEFINITION_XHTMLs[datatype_definition_xhtml]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitDATATYPE_DEFINITION_XHTML(datatype_definition_xhtml)
		}
	}
	return datatype_definition_xhtml
}

func (datatype_definition_xhtml *DATATYPE_DEFINITION_XHTML) CommitVoid(stage *StageStruct) {
	datatype_definition_xhtml.Commit(stage)
}

// Checkout datatype_definition_xhtml to the back repo (if it is already staged)
func (datatype_definition_xhtml *DATATYPE_DEFINITION_XHTML) Checkout(stage *StageStruct) *DATATYPE_DEFINITION_XHTML {
	if _, ok := stage.DATATYPE_DEFINITION_XHTMLs[datatype_definition_xhtml]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutDATATYPE_DEFINITION_XHTML(datatype_definition_xhtml)
		}
	}
	return datatype_definition_xhtml
}

// for satisfaction of GongStruct interface
func (datatype_definition_xhtml *DATATYPE_DEFINITION_XHTML) GetName() (res string) {
	return datatype_definition_xhtml.Name
}

// Stage puts embedded_value to the model stage
func (embedded_value *EMBEDDED_VALUE) Stage(stage *StageStruct) *EMBEDDED_VALUE {
	stage.EMBEDDED_VALUEs[embedded_value] = __member
	stage.EMBEDDED_VALUEs_mapString[embedded_value.Name] = embedded_value

	return embedded_value
}

// Unstage removes embedded_value off the model stage
func (embedded_value *EMBEDDED_VALUE) Unstage(stage *StageStruct) *EMBEDDED_VALUE {
	delete(stage.EMBEDDED_VALUEs, embedded_value)
	delete(stage.EMBEDDED_VALUEs_mapString, embedded_value.Name)
	return embedded_value
}

// UnstageVoid removes embedded_value off the model stage
func (embedded_value *EMBEDDED_VALUE) UnstageVoid(stage *StageStruct) {
	delete(stage.EMBEDDED_VALUEs, embedded_value)
	delete(stage.EMBEDDED_VALUEs_mapString, embedded_value.Name)
}

// commit embedded_value to the back repo (if it is already staged)
func (embedded_value *EMBEDDED_VALUE) Commit(stage *StageStruct) *EMBEDDED_VALUE {
	if _, ok := stage.EMBEDDED_VALUEs[embedded_value]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitEMBEDDED_VALUE(embedded_value)
		}
	}
	return embedded_value
}

func (embedded_value *EMBEDDED_VALUE) CommitVoid(stage *StageStruct) {
	embedded_value.Commit(stage)
}

// Checkout embedded_value to the back repo (if it is already staged)
func (embedded_value *EMBEDDED_VALUE) Checkout(stage *StageStruct) *EMBEDDED_VALUE {
	if _, ok := stage.EMBEDDED_VALUEs[embedded_value]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutEMBEDDED_VALUE(embedded_value)
		}
	}
	return embedded_value
}

// for satisfaction of GongStruct interface
func (embedded_value *EMBEDDED_VALUE) GetName() (res string) {
	return embedded_value.Name
}

// Stage puts enum_value to the model stage
func (enum_value *ENUM_VALUE) Stage(stage *StageStruct) *ENUM_VALUE {
	stage.ENUM_VALUEs[enum_value] = __member
	stage.ENUM_VALUEs_mapString[enum_value.Name] = enum_value

	return enum_value
}

// Unstage removes enum_value off the model stage
func (enum_value *ENUM_VALUE) Unstage(stage *StageStruct) *ENUM_VALUE {
	delete(stage.ENUM_VALUEs, enum_value)
	delete(stage.ENUM_VALUEs_mapString, enum_value.Name)
	return enum_value
}

// UnstageVoid removes enum_value off the model stage
func (enum_value *ENUM_VALUE) UnstageVoid(stage *StageStruct) {
	delete(stage.ENUM_VALUEs, enum_value)
	delete(stage.ENUM_VALUEs_mapString, enum_value.Name)
}

// commit enum_value to the back repo (if it is already staged)
func (enum_value *ENUM_VALUE) Commit(stage *StageStruct) *ENUM_VALUE {
	if _, ok := stage.ENUM_VALUEs[enum_value]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitENUM_VALUE(enum_value)
		}
	}
	return enum_value
}

func (enum_value *ENUM_VALUE) CommitVoid(stage *StageStruct) {
	enum_value.Commit(stage)
}

// Checkout enum_value to the back repo (if it is already staged)
func (enum_value *ENUM_VALUE) Checkout(stage *StageStruct) *ENUM_VALUE {
	if _, ok := stage.ENUM_VALUEs[enum_value]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutENUM_VALUE(enum_value)
		}
	}
	return enum_value
}

// for satisfaction of GongStruct interface
func (enum_value *ENUM_VALUE) GetName() (res string) {
	return enum_value.Name
}

// Stage puts relation_group to the model stage
func (relation_group *RELATION_GROUP) Stage(stage *StageStruct) *RELATION_GROUP {
	stage.RELATION_GROUPs[relation_group] = __member
	stage.RELATION_GROUPs_mapString[relation_group.Name] = relation_group

	return relation_group
}

// Unstage removes relation_group off the model stage
func (relation_group *RELATION_GROUP) Unstage(stage *StageStruct) *RELATION_GROUP {
	delete(stage.RELATION_GROUPs, relation_group)
	delete(stage.RELATION_GROUPs_mapString, relation_group.Name)
	return relation_group
}

// UnstageVoid removes relation_group off the model stage
func (relation_group *RELATION_GROUP) UnstageVoid(stage *StageStruct) {
	delete(stage.RELATION_GROUPs, relation_group)
	delete(stage.RELATION_GROUPs_mapString, relation_group.Name)
}

// commit relation_group to the back repo (if it is already staged)
func (relation_group *RELATION_GROUP) Commit(stage *StageStruct) *RELATION_GROUP {
	if _, ok := stage.RELATION_GROUPs[relation_group]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitRELATION_GROUP(relation_group)
		}
	}
	return relation_group
}

func (relation_group *RELATION_GROUP) CommitVoid(stage *StageStruct) {
	relation_group.Commit(stage)
}

// Checkout relation_group to the back repo (if it is already staged)
func (relation_group *RELATION_GROUP) Checkout(stage *StageStruct) *RELATION_GROUP {
	if _, ok := stage.RELATION_GROUPs[relation_group]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutRELATION_GROUP(relation_group)
		}
	}
	return relation_group
}

// for satisfaction of GongStruct interface
func (relation_group *RELATION_GROUP) GetName() (res string) {
	return relation_group.Name
}

// Stage puts relation_group_type to the model stage
func (relation_group_type *RELATION_GROUP_TYPE) Stage(stage *StageStruct) *RELATION_GROUP_TYPE {
	stage.RELATION_GROUP_TYPEs[relation_group_type] = __member
	stage.RELATION_GROUP_TYPEs_mapString[relation_group_type.Name] = relation_group_type

	return relation_group_type
}

// Unstage removes relation_group_type off the model stage
func (relation_group_type *RELATION_GROUP_TYPE) Unstage(stage *StageStruct) *RELATION_GROUP_TYPE {
	delete(stage.RELATION_GROUP_TYPEs, relation_group_type)
	delete(stage.RELATION_GROUP_TYPEs_mapString, relation_group_type.Name)
	return relation_group_type
}

// UnstageVoid removes relation_group_type off the model stage
func (relation_group_type *RELATION_GROUP_TYPE) UnstageVoid(stage *StageStruct) {
	delete(stage.RELATION_GROUP_TYPEs, relation_group_type)
	delete(stage.RELATION_GROUP_TYPEs_mapString, relation_group_type.Name)
}

// commit relation_group_type to the back repo (if it is already staged)
func (relation_group_type *RELATION_GROUP_TYPE) Commit(stage *StageStruct) *RELATION_GROUP_TYPE {
	if _, ok := stage.RELATION_GROUP_TYPEs[relation_group_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitRELATION_GROUP_TYPE(relation_group_type)
		}
	}
	return relation_group_type
}

func (relation_group_type *RELATION_GROUP_TYPE) CommitVoid(stage *StageStruct) {
	relation_group_type.Commit(stage)
}

// Checkout relation_group_type to the back repo (if it is already staged)
func (relation_group_type *RELATION_GROUP_TYPE) Checkout(stage *StageStruct) *RELATION_GROUP_TYPE {
	if _, ok := stage.RELATION_GROUP_TYPEs[relation_group_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutRELATION_GROUP_TYPE(relation_group_type)
		}
	}
	return relation_group_type
}

// for satisfaction of GongStruct interface
func (relation_group_type *RELATION_GROUP_TYPE) GetName() (res string) {
	return relation_group_type.Name
}

// Stage puts req_if to the model stage
func (req_if *REQ_IF) Stage(stage *StageStruct) *REQ_IF {
	stage.REQ_IFs[req_if] = __member
	stage.REQ_IFs_mapString[req_if.Name] = req_if

	return req_if
}

// Unstage removes req_if off the model stage
func (req_if *REQ_IF) Unstage(stage *StageStruct) *REQ_IF {
	delete(stage.REQ_IFs, req_if)
	delete(stage.REQ_IFs_mapString, req_if.Name)
	return req_if
}

// UnstageVoid removes req_if off the model stage
func (req_if *REQ_IF) UnstageVoid(stage *StageStruct) {
	delete(stage.REQ_IFs, req_if)
	delete(stage.REQ_IFs_mapString, req_if.Name)
}

// commit req_if to the back repo (if it is already staged)
func (req_if *REQ_IF) Commit(stage *StageStruct) *REQ_IF {
	if _, ok := stage.REQ_IFs[req_if]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitREQ_IF(req_if)
		}
	}
	return req_if
}

func (req_if *REQ_IF) CommitVoid(stage *StageStruct) {
	req_if.Commit(stage)
}

// Checkout req_if to the back repo (if it is already staged)
func (req_if *REQ_IF) Checkout(stage *StageStruct) *REQ_IF {
	if _, ok := stage.REQ_IFs[req_if]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutREQ_IF(req_if)
		}
	}
	return req_if
}

// for satisfaction of GongStruct interface
func (req_if *REQ_IF) GetName() (res string) {
	return req_if.Name
}

// Stage puts req_if_content to the model stage
func (req_if_content *REQ_IF_CONTENT) Stage(stage *StageStruct) *REQ_IF_CONTENT {
	stage.REQ_IF_CONTENTs[req_if_content] = __member
	stage.REQ_IF_CONTENTs_mapString[req_if_content.Name] = req_if_content

	return req_if_content
}

// Unstage removes req_if_content off the model stage
func (req_if_content *REQ_IF_CONTENT) Unstage(stage *StageStruct) *REQ_IF_CONTENT {
	delete(stage.REQ_IF_CONTENTs, req_if_content)
	delete(stage.REQ_IF_CONTENTs_mapString, req_if_content.Name)
	return req_if_content
}

// UnstageVoid removes req_if_content off the model stage
func (req_if_content *REQ_IF_CONTENT) UnstageVoid(stage *StageStruct) {
	delete(stage.REQ_IF_CONTENTs, req_if_content)
	delete(stage.REQ_IF_CONTENTs_mapString, req_if_content.Name)
}

// commit req_if_content to the back repo (if it is already staged)
func (req_if_content *REQ_IF_CONTENT) Commit(stage *StageStruct) *REQ_IF_CONTENT {
	if _, ok := stage.REQ_IF_CONTENTs[req_if_content]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitREQ_IF_CONTENT(req_if_content)
		}
	}
	return req_if_content
}

func (req_if_content *REQ_IF_CONTENT) CommitVoid(stage *StageStruct) {
	req_if_content.Commit(stage)
}

// Checkout req_if_content to the back repo (if it is already staged)
func (req_if_content *REQ_IF_CONTENT) Checkout(stage *StageStruct) *REQ_IF_CONTENT {
	if _, ok := stage.REQ_IF_CONTENTs[req_if_content]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutREQ_IF_CONTENT(req_if_content)
		}
	}
	return req_if_content
}

// for satisfaction of GongStruct interface
func (req_if_content *REQ_IF_CONTENT) GetName() (res string) {
	return req_if_content.Name
}

// Stage puts req_if_header to the model stage
func (req_if_header *REQ_IF_HEADER) Stage(stage *StageStruct) *REQ_IF_HEADER {
	stage.REQ_IF_HEADERs[req_if_header] = __member
	stage.REQ_IF_HEADERs_mapString[req_if_header.Name] = req_if_header

	return req_if_header
}

// Unstage removes req_if_header off the model stage
func (req_if_header *REQ_IF_HEADER) Unstage(stage *StageStruct) *REQ_IF_HEADER {
	delete(stage.REQ_IF_HEADERs, req_if_header)
	delete(stage.REQ_IF_HEADERs_mapString, req_if_header.Name)
	return req_if_header
}

// UnstageVoid removes req_if_header off the model stage
func (req_if_header *REQ_IF_HEADER) UnstageVoid(stage *StageStruct) {
	delete(stage.REQ_IF_HEADERs, req_if_header)
	delete(stage.REQ_IF_HEADERs_mapString, req_if_header.Name)
}

// commit req_if_header to the back repo (if it is already staged)
func (req_if_header *REQ_IF_HEADER) Commit(stage *StageStruct) *REQ_IF_HEADER {
	if _, ok := stage.REQ_IF_HEADERs[req_if_header]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitREQ_IF_HEADER(req_if_header)
		}
	}
	return req_if_header
}

func (req_if_header *REQ_IF_HEADER) CommitVoid(stage *StageStruct) {
	req_if_header.Commit(stage)
}

// Checkout req_if_header to the back repo (if it is already staged)
func (req_if_header *REQ_IF_HEADER) Checkout(stage *StageStruct) *REQ_IF_HEADER {
	if _, ok := stage.REQ_IF_HEADERs[req_if_header]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutREQ_IF_HEADER(req_if_header)
		}
	}
	return req_if_header
}

// for satisfaction of GongStruct interface
func (req_if_header *REQ_IF_HEADER) GetName() (res string) {
	return req_if_header.Name
}

// Stage puts req_if_tool_extension to the model stage
func (req_if_tool_extension *REQ_IF_TOOL_EXTENSION) Stage(stage *StageStruct) *REQ_IF_TOOL_EXTENSION {
	stage.REQ_IF_TOOL_EXTENSIONs[req_if_tool_extension] = __member
	stage.REQ_IF_TOOL_EXTENSIONs_mapString[req_if_tool_extension.Name] = req_if_tool_extension

	return req_if_tool_extension
}

// Unstage removes req_if_tool_extension off the model stage
func (req_if_tool_extension *REQ_IF_TOOL_EXTENSION) Unstage(stage *StageStruct) *REQ_IF_TOOL_EXTENSION {
	delete(stage.REQ_IF_TOOL_EXTENSIONs, req_if_tool_extension)
	delete(stage.REQ_IF_TOOL_EXTENSIONs_mapString, req_if_tool_extension.Name)
	return req_if_tool_extension
}

// UnstageVoid removes req_if_tool_extension off the model stage
func (req_if_tool_extension *REQ_IF_TOOL_EXTENSION) UnstageVoid(stage *StageStruct) {
	delete(stage.REQ_IF_TOOL_EXTENSIONs, req_if_tool_extension)
	delete(stage.REQ_IF_TOOL_EXTENSIONs_mapString, req_if_tool_extension.Name)
}

// commit req_if_tool_extension to the back repo (if it is already staged)
func (req_if_tool_extension *REQ_IF_TOOL_EXTENSION) Commit(stage *StageStruct) *REQ_IF_TOOL_EXTENSION {
	if _, ok := stage.REQ_IF_TOOL_EXTENSIONs[req_if_tool_extension]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitREQ_IF_TOOL_EXTENSION(req_if_tool_extension)
		}
	}
	return req_if_tool_extension
}

func (req_if_tool_extension *REQ_IF_TOOL_EXTENSION) CommitVoid(stage *StageStruct) {
	req_if_tool_extension.Commit(stage)
}

// Checkout req_if_tool_extension to the back repo (if it is already staged)
func (req_if_tool_extension *REQ_IF_TOOL_EXTENSION) Checkout(stage *StageStruct) *REQ_IF_TOOL_EXTENSION {
	if _, ok := stage.REQ_IF_TOOL_EXTENSIONs[req_if_tool_extension]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutREQ_IF_TOOL_EXTENSION(req_if_tool_extension)
		}
	}
	return req_if_tool_extension
}

// for satisfaction of GongStruct interface
func (req_if_tool_extension *REQ_IF_TOOL_EXTENSION) GetName() (res string) {
	return req_if_tool_extension.Name
}

// Stage puts specification to the model stage
func (specification *SPECIFICATION) Stage(stage *StageStruct) *SPECIFICATION {
	stage.SPECIFICATIONs[specification] = __member
	stage.SPECIFICATIONs_mapString[specification.Name] = specification

	return specification
}

// Unstage removes specification off the model stage
func (specification *SPECIFICATION) Unstage(stage *StageStruct) *SPECIFICATION {
	delete(stage.SPECIFICATIONs, specification)
	delete(stage.SPECIFICATIONs_mapString, specification.Name)
	return specification
}

// UnstageVoid removes specification off the model stage
func (specification *SPECIFICATION) UnstageVoid(stage *StageStruct) {
	delete(stage.SPECIFICATIONs, specification)
	delete(stage.SPECIFICATIONs_mapString, specification.Name)
}

// commit specification to the back repo (if it is already staged)
func (specification *SPECIFICATION) Commit(stage *StageStruct) *SPECIFICATION {
	if _, ok := stage.SPECIFICATIONs[specification]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSPECIFICATION(specification)
		}
	}
	return specification
}

func (specification *SPECIFICATION) CommitVoid(stage *StageStruct) {
	specification.Commit(stage)
}

// Checkout specification to the back repo (if it is already staged)
func (specification *SPECIFICATION) Checkout(stage *StageStruct) *SPECIFICATION {
	if _, ok := stage.SPECIFICATIONs[specification]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutSPECIFICATION(specification)
		}
	}
	return specification
}

// for satisfaction of GongStruct interface
func (specification *SPECIFICATION) GetName() (res string) {
	return specification.Name
}

// Stage puts specification_type to the model stage
func (specification_type *SPECIFICATION_TYPE) Stage(stage *StageStruct) *SPECIFICATION_TYPE {
	stage.SPECIFICATION_TYPEs[specification_type] = __member
	stage.SPECIFICATION_TYPEs_mapString[specification_type.Name] = specification_type

	return specification_type
}

// Unstage removes specification_type off the model stage
func (specification_type *SPECIFICATION_TYPE) Unstage(stage *StageStruct) *SPECIFICATION_TYPE {
	delete(stage.SPECIFICATION_TYPEs, specification_type)
	delete(stage.SPECIFICATION_TYPEs_mapString, specification_type.Name)
	return specification_type
}

// UnstageVoid removes specification_type off the model stage
func (specification_type *SPECIFICATION_TYPE) UnstageVoid(stage *StageStruct) {
	delete(stage.SPECIFICATION_TYPEs, specification_type)
	delete(stage.SPECIFICATION_TYPEs_mapString, specification_type.Name)
}

// commit specification_type to the back repo (if it is already staged)
func (specification_type *SPECIFICATION_TYPE) Commit(stage *StageStruct) *SPECIFICATION_TYPE {
	if _, ok := stage.SPECIFICATION_TYPEs[specification_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSPECIFICATION_TYPE(specification_type)
		}
	}
	return specification_type
}

func (specification_type *SPECIFICATION_TYPE) CommitVoid(stage *StageStruct) {
	specification_type.Commit(stage)
}

// Checkout specification_type to the back repo (if it is already staged)
func (specification_type *SPECIFICATION_TYPE) Checkout(stage *StageStruct) *SPECIFICATION_TYPE {
	if _, ok := stage.SPECIFICATION_TYPEs[specification_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutSPECIFICATION_TYPE(specification_type)
		}
	}
	return specification_type
}

// for satisfaction of GongStruct interface
func (specification_type *SPECIFICATION_TYPE) GetName() (res string) {
	return specification_type.Name
}

// Stage puts spec_hierarchy to the model stage
func (spec_hierarchy *SPEC_HIERARCHY) Stage(stage *StageStruct) *SPEC_HIERARCHY {
	stage.SPEC_HIERARCHYs[spec_hierarchy] = __member
	stage.SPEC_HIERARCHYs_mapString[spec_hierarchy.Name] = spec_hierarchy

	return spec_hierarchy
}

// Unstage removes spec_hierarchy off the model stage
func (spec_hierarchy *SPEC_HIERARCHY) Unstage(stage *StageStruct) *SPEC_HIERARCHY {
	delete(stage.SPEC_HIERARCHYs, spec_hierarchy)
	delete(stage.SPEC_HIERARCHYs_mapString, spec_hierarchy.Name)
	return spec_hierarchy
}

// UnstageVoid removes spec_hierarchy off the model stage
func (spec_hierarchy *SPEC_HIERARCHY) UnstageVoid(stage *StageStruct) {
	delete(stage.SPEC_HIERARCHYs, spec_hierarchy)
	delete(stage.SPEC_HIERARCHYs_mapString, spec_hierarchy.Name)
}

// commit spec_hierarchy to the back repo (if it is already staged)
func (spec_hierarchy *SPEC_HIERARCHY) Commit(stage *StageStruct) *SPEC_HIERARCHY {
	if _, ok := stage.SPEC_HIERARCHYs[spec_hierarchy]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSPEC_HIERARCHY(spec_hierarchy)
		}
	}
	return spec_hierarchy
}

func (spec_hierarchy *SPEC_HIERARCHY) CommitVoid(stage *StageStruct) {
	spec_hierarchy.Commit(stage)
}

// Checkout spec_hierarchy to the back repo (if it is already staged)
func (spec_hierarchy *SPEC_HIERARCHY) Checkout(stage *StageStruct) *SPEC_HIERARCHY {
	if _, ok := stage.SPEC_HIERARCHYs[spec_hierarchy]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutSPEC_HIERARCHY(spec_hierarchy)
		}
	}
	return spec_hierarchy
}

// for satisfaction of GongStruct interface
func (spec_hierarchy *SPEC_HIERARCHY) GetName() (res string) {
	return spec_hierarchy.Name
}

// Stage puts spec_object to the model stage
func (spec_object *SPEC_OBJECT) Stage(stage *StageStruct) *SPEC_OBJECT {
	stage.SPEC_OBJECTs[spec_object] = __member
	stage.SPEC_OBJECTs_mapString[spec_object.Name] = spec_object

	return spec_object
}

// Unstage removes spec_object off the model stage
func (spec_object *SPEC_OBJECT) Unstage(stage *StageStruct) *SPEC_OBJECT {
	delete(stage.SPEC_OBJECTs, spec_object)
	delete(stage.SPEC_OBJECTs_mapString, spec_object.Name)
	return spec_object
}

// UnstageVoid removes spec_object off the model stage
func (spec_object *SPEC_OBJECT) UnstageVoid(stage *StageStruct) {
	delete(stage.SPEC_OBJECTs, spec_object)
	delete(stage.SPEC_OBJECTs_mapString, spec_object.Name)
}

// commit spec_object to the back repo (if it is already staged)
func (spec_object *SPEC_OBJECT) Commit(stage *StageStruct) *SPEC_OBJECT {
	if _, ok := stage.SPEC_OBJECTs[spec_object]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSPEC_OBJECT(spec_object)
		}
	}
	return spec_object
}

func (spec_object *SPEC_OBJECT) CommitVoid(stage *StageStruct) {
	spec_object.Commit(stage)
}

// Checkout spec_object to the back repo (if it is already staged)
func (spec_object *SPEC_OBJECT) Checkout(stage *StageStruct) *SPEC_OBJECT {
	if _, ok := stage.SPEC_OBJECTs[spec_object]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutSPEC_OBJECT(spec_object)
		}
	}
	return spec_object
}

// for satisfaction of GongStruct interface
func (spec_object *SPEC_OBJECT) GetName() (res string) {
	return spec_object.Name
}

// Stage puts spec_object_type to the model stage
func (spec_object_type *SPEC_OBJECT_TYPE) Stage(stage *StageStruct) *SPEC_OBJECT_TYPE {
	stage.SPEC_OBJECT_TYPEs[spec_object_type] = __member
	stage.SPEC_OBJECT_TYPEs_mapString[spec_object_type.Name] = spec_object_type

	return spec_object_type
}

// Unstage removes spec_object_type off the model stage
func (spec_object_type *SPEC_OBJECT_TYPE) Unstage(stage *StageStruct) *SPEC_OBJECT_TYPE {
	delete(stage.SPEC_OBJECT_TYPEs, spec_object_type)
	delete(stage.SPEC_OBJECT_TYPEs_mapString, spec_object_type.Name)
	return spec_object_type
}

// UnstageVoid removes spec_object_type off the model stage
func (spec_object_type *SPEC_OBJECT_TYPE) UnstageVoid(stage *StageStruct) {
	delete(stage.SPEC_OBJECT_TYPEs, spec_object_type)
	delete(stage.SPEC_OBJECT_TYPEs_mapString, spec_object_type.Name)
}

// commit spec_object_type to the back repo (if it is already staged)
func (spec_object_type *SPEC_OBJECT_TYPE) Commit(stage *StageStruct) *SPEC_OBJECT_TYPE {
	if _, ok := stage.SPEC_OBJECT_TYPEs[spec_object_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSPEC_OBJECT_TYPE(spec_object_type)
		}
	}
	return spec_object_type
}

func (spec_object_type *SPEC_OBJECT_TYPE) CommitVoid(stage *StageStruct) {
	spec_object_type.Commit(stage)
}

// Checkout spec_object_type to the back repo (if it is already staged)
func (spec_object_type *SPEC_OBJECT_TYPE) Checkout(stage *StageStruct) *SPEC_OBJECT_TYPE {
	if _, ok := stage.SPEC_OBJECT_TYPEs[spec_object_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutSPEC_OBJECT_TYPE(spec_object_type)
		}
	}
	return spec_object_type
}

// for satisfaction of GongStruct interface
func (spec_object_type *SPEC_OBJECT_TYPE) GetName() (res string) {
	return spec_object_type.Name
}

// Stage puts spec_relation to the model stage
func (spec_relation *SPEC_RELATION) Stage(stage *StageStruct) *SPEC_RELATION {
	stage.SPEC_RELATIONs[spec_relation] = __member
	stage.SPEC_RELATIONs_mapString[spec_relation.Name] = spec_relation

	return spec_relation
}

// Unstage removes spec_relation off the model stage
func (spec_relation *SPEC_RELATION) Unstage(stage *StageStruct) *SPEC_RELATION {
	delete(stage.SPEC_RELATIONs, spec_relation)
	delete(stage.SPEC_RELATIONs_mapString, spec_relation.Name)
	return spec_relation
}

// UnstageVoid removes spec_relation off the model stage
func (spec_relation *SPEC_RELATION) UnstageVoid(stage *StageStruct) {
	delete(stage.SPEC_RELATIONs, spec_relation)
	delete(stage.SPEC_RELATIONs_mapString, spec_relation.Name)
}

// commit spec_relation to the back repo (if it is already staged)
func (spec_relation *SPEC_RELATION) Commit(stage *StageStruct) *SPEC_RELATION {
	if _, ok := stage.SPEC_RELATIONs[spec_relation]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSPEC_RELATION(spec_relation)
		}
	}
	return spec_relation
}

func (spec_relation *SPEC_RELATION) CommitVoid(stage *StageStruct) {
	spec_relation.Commit(stage)
}

// Checkout spec_relation to the back repo (if it is already staged)
func (spec_relation *SPEC_RELATION) Checkout(stage *StageStruct) *SPEC_RELATION {
	if _, ok := stage.SPEC_RELATIONs[spec_relation]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutSPEC_RELATION(spec_relation)
		}
	}
	return spec_relation
}

// for satisfaction of GongStruct interface
func (spec_relation *SPEC_RELATION) GetName() (res string) {
	return spec_relation.Name
}

// Stage puts spec_relation_type to the model stage
func (spec_relation_type *SPEC_RELATION_TYPE) Stage(stage *StageStruct) *SPEC_RELATION_TYPE {
	stage.SPEC_RELATION_TYPEs[spec_relation_type] = __member
	stage.SPEC_RELATION_TYPEs_mapString[spec_relation_type.Name] = spec_relation_type

	return spec_relation_type
}

// Unstage removes spec_relation_type off the model stage
func (spec_relation_type *SPEC_RELATION_TYPE) Unstage(stage *StageStruct) *SPEC_RELATION_TYPE {
	delete(stage.SPEC_RELATION_TYPEs, spec_relation_type)
	delete(stage.SPEC_RELATION_TYPEs_mapString, spec_relation_type.Name)
	return spec_relation_type
}

// UnstageVoid removes spec_relation_type off the model stage
func (spec_relation_type *SPEC_RELATION_TYPE) UnstageVoid(stage *StageStruct) {
	delete(stage.SPEC_RELATION_TYPEs, spec_relation_type)
	delete(stage.SPEC_RELATION_TYPEs_mapString, spec_relation_type.Name)
}

// commit spec_relation_type to the back repo (if it is already staged)
func (spec_relation_type *SPEC_RELATION_TYPE) Commit(stage *StageStruct) *SPEC_RELATION_TYPE {
	if _, ok := stage.SPEC_RELATION_TYPEs[spec_relation_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSPEC_RELATION_TYPE(spec_relation_type)
		}
	}
	return spec_relation_type
}

func (spec_relation_type *SPEC_RELATION_TYPE) CommitVoid(stage *StageStruct) {
	spec_relation_type.Commit(stage)
}

// Checkout spec_relation_type to the back repo (if it is already staged)
func (spec_relation_type *SPEC_RELATION_TYPE) Checkout(stage *StageStruct) *SPEC_RELATION_TYPE {
	if _, ok := stage.SPEC_RELATION_TYPEs[spec_relation_type]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutSPEC_RELATION_TYPE(spec_relation_type)
		}
	}
	return spec_relation_type
}

// for satisfaction of GongStruct interface
func (spec_relation_type *SPEC_RELATION_TYPE) GetName() (res string) {
	return spec_relation_type.Name
}

// Stage puts xhtml_content to the model stage
func (xhtml_content *XHTML_CONTENT) Stage(stage *StageStruct) *XHTML_CONTENT {
	stage.XHTML_CONTENTs[xhtml_content] = __member
	stage.XHTML_CONTENTs_mapString[xhtml_content.Name] = xhtml_content

	return xhtml_content
}

// Unstage removes xhtml_content off the model stage
func (xhtml_content *XHTML_CONTENT) Unstage(stage *StageStruct) *XHTML_CONTENT {
	delete(stage.XHTML_CONTENTs, xhtml_content)
	delete(stage.XHTML_CONTENTs_mapString, xhtml_content.Name)
	return xhtml_content
}

// UnstageVoid removes xhtml_content off the model stage
func (xhtml_content *XHTML_CONTENT) UnstageVoid(stage *StageStruct) {
	delete(stage.XHTML_CONTENTs, xhtml_content)
	delete(stage.XHTML_CONTENTs_mapString, xhtml_content.Name)
}

// commit xhtml_content to the back repo (if it is already staged)
func (xhtml_content *XHTML_CONTENT) Commit(stage *StageStruct) *XHTML_CONTENT {
	if _, ok := stage.XHTML_CONTENTs[xhtml_content]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitXHTML_CONTENT(xhtml_content)
		}
	}
	return xhtml_content
}

func (xhtml_content *XHTML_CONTENT) CommitVoid(stage *StageStruct) {
	xhtml_content.Commit(stage)
}

// Checkout xhtml_content to the back repo (if it is already staged)
func (xhtml_content *XHTML_CONTENT) Checkout(stage *StageStruct) *XHTML_CONTENT {
	if _, ok := stage.XHTML_CONTENTs[xhtml_content]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutXHTML_CONTENT(xhtml_content)
		}
	}
	return xhtml_content
}

// for satisfaction of GongStruct interface
func (xhtml_content *XHTML_CONTENT) GetName() (res string) {
	return xhtml_content.Name
}

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMALTERNATIVE_ID(ALTERNATIVE_ID *ALTERNATIVE_ID)
	CreateORMATTRIBUTE_DEFINITION_BOOLEAN(ATTRIBUTE_DEFINITION_BOOLEAN *ATTRIBUTE_DEFINITION_BOOLEAN)
	CreateORMATTRIBUTE_DEFINITION_DATE(ATTRIBUTE_DEFINITION_DATE *ATTRIBUTE_DEFINITION_DATE)
	CreateORMATTRIBUTE_DEFINITION_ENUMERATION(ATTRIBUTE_DEFINITION_ENUMERATION *ATTRIBUTE_DEFINITION_ENUMERATION)
	CreateORMATTRIBUTE_DEFINITION_INTEGER(ATTRIBUTE_DEFINITION_INTEGER *ATTRIBUTE_DEFINITION_INTEGER)
	CreateORMATTRIBUTE_DEFINITION_REAL(ATTRIBUTE_DEFINITION_REAL *ATTRIBUTE_DEFINITION_REAL)
	CreateORMATTRIBUTE_DEFINITION_STRING(ATTRIBUTE_DEFINITION_STRING *ATTRIBUTE_DEFINITION_STRING)
	CreateORMATTRIBUTE_DEFINITION_XHTML(ATTRIBUTE_DEFINITION_XHTML *ATTRIBUTE_DEFINITION_XHTML)
	CreateORMATTRIBUTE_VALUE_BOOLEAN(ATTRIBUTE_VALUE_BOOLEAN *ATTRIBUTE_VALUE_BOOLEAN)
	CreateORMATTRIBUTE_VALUE_DATE(ATTRIBUTE_VALUE_DATE *ATTRIBUTE_VALUE_DATE)
	CreateORMATTRIBUTE_VALUE_ENUMERATION(ATTRIBUTE_VALUE_ENUMERATION *ATTRIBUTE_VALUE_ENUMERATION)
	CreateORMATTRIBUTE_VALUE_INTEGER(ATTRIBUTE_VALUE_INTEGER *ATTRIBUTE_VALUE_INTEGER)
	CreateORMATTRIBUTE_VALUE_REAL(ATTRIBUTE_VALUE_REAL *ATTRIBUTE_VALUE_REAL)
	CreateORMATTRIBUTE_VALUE_STRING(ATTRIBUTE_VALUE_STRING *ATTRIBUTE_VALUE_STRING)
	CreateORMATTRIBUTE_VALUE_XHTML(ATTRIBUTE_VALUE_XHTML *ATTRIBUTE_VALUE_XHTML)
	CreateORMDATATYPE_DEFINITION_BOOLEAN(DATATYPE_DEFINITION_BOOLEAN *DATATYPE_DEFINITION_BOOLEAN)
	CreateORMDATATYPE_DEFINITION_DATE(DATATYPE_DEFINITION_DATE *DATATYPE_DEFINITION_DATE)
	CreateORMDATATYPE_DEFINITION_ENUMERATION(DATATYPE_DEFINITION_ENUMERATION *DATATYPE_DEFINITION_ENUMERATION)
	CreateORMDATATYPE_DEFINITION_INTEGER(DATATYPE_DEFINITION_INTEGER *DATATYPE_DEFINITION_INTEGER)
	CreateORMDATATYPE_DEFINITION_REAL(DATATYPE_DEFINITION_REAL *DATATYPE_DEFINITION_REAL)
	CreateORMDATATYPE_DEFINITION_STRING(DATATYPE_DEFINITION_STRING *DATATYPE_DEFINITION_STRING)
	CreateORMDATATYPE_DEFINITION_XHTML(DATATYPE_DEFINITION_XHTML *DATATYPE_DEFINITION_XHTML)
	CreateORMEMBEDDED_VALUE(EMBEDDED_VALUE *EMBEDDED_VALUE)
	CreateORMENUM_VALUE(ENUM_VALUE *ENUM_VALUE)
	CreateORMRELATION_GROUP(RELATION_GROUP *RELATION_GROUP)
	CreateORMRELATION_GROUP_TYPE(RELATION_GROUP_TYPE *RELATION_GROUP_TYPE)
	CreateORMREQ_IF(REQ_IF *REQ_IF)
	CreateORMREQ_IF_CONTENT(REQ_IF_CONTENT *REQ_IF_CONTENT)
	CreateORMREQ_IF_HEADER(REQ_IF_HEADER *REQ_IF_HEADER)
	CreateORMREQ_IF_TOOL_EXTENSION(REQ_IF_TOOL_EXTENSION *REQ_IF_TOOL_EXTENSION)
	CreateORMSPECIFICATION(SPECIFICATION *SPECIFICATION)
	CreateORMSPECIFICATION_TYPE(SPECIFICATION_TYPE *SPECIFICATION_TYPE)
	CreateORMSPEC_HIERARCHY(SPEC_HIERARCHY *SPEC_HIERARCHY)
	CreateORMSPEC_OBJECT(SPEC_OBJECT *SPEC_OBJECT)
	CreateORMSPEC_OBJECT_TYPE(SPEC_OBJECT_TYPE *SPEC_OBJECT_TYPE)
	CreateORMSPEC_RELATION(SPEC_RELATION *SPEC_RELATION)
	CreateORMSPEC_RELATION_TYPE(SPEC_RELATION_TYPE *SPEC_RELATION_TYPE)
	CreateORMXHTML_CONTENT(XHTML_CONTENT *XHTML_CONTENT)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMALTERNATIVE_ID(ALTERNATIVE_ID *ALTERNATIVE_ID)
	DeleteORMATTRIBUTE_DEFINITION_BOOLEAN(ATTRIBUTE_DEFINITION_BOOLEAN *ATTRIBUTE_DEFINITION_BOOLEAN)
	DeleteORMATTRIBUTE_DEFINITION_DATE(ATTRIBUTE_DEFINITION_DATE *ATTRIBUTE_DEFINITION_DATE)
	DeleteORMATTRIBUTE_DEFINITION_ENUMERATION(ATTRIBUTE_DEFINITION_ENUMERATION *ATTRIBUTE_DEFINITION_ENUMERATION)
	DeleteORMATTRIBUTE_DEFINITION_INTEGER(ATTRIBUTE_DEFINITION_INTEGER *ATTRIBUTE_DEFINITION_INTEGER)
	DeleteORMATTRIBUTE_DEFINITION_REAL(ATTRIBUTE_DEFINITION_REAL *ATTRIBUTE_DEFINITION_REAL)
	DeleteORMATTRIBUTE_DEFINITION_STRING(ATTRIBUTE_DEFINITION_STRING *ATTRIBUTE_DEFINITION_STRING)
	DeleteORMATTRIBUTE_DEFINITION_XHTML(ATTRIBUTE_DEFINITION_XHTML *ATTRIBUTE_DEFINITION_XHTML)
	DeleteORMATTRIBUTE_VALUE_BOOLEAN(ATTRIBUTE_VALUE_BOOLEAN *ATTRIBUTE_VALUE_BOOLEAN)
	DeleteORMATTRIBUTE_VALUE_DATE(ATTRIBUTE_VALUE_DATE *ATTRIBUTE_VALUE_DATE)
	DeleteORMATTRIBUTE_VALUE_ENUMERATION(ATTRIBUTE_VALUE_ENUMERATION *ATTRIBUTE_VALUE_ENUMERATION)
	DeleteORMATTRIBUTE_VALUE_INTEGER(ATTRIBUTE_VALUE_INTEGER *ATTRIBUTE_VALUE_INTEGER)
	DeleteORMATTRIBUTE_VALUE_REAL(ATTRIBUTE_VALUE_REAL *ATTRIBUTE_VALUE_REAL)
	DeleteORMATTRIBUTE_VALUE_STRING(ATTRIBUTE_VALUE_STRING *ATTRIBUTE_VALUE_STRING)
	DeleteORMATTRIBUTE_VALUE_XHTML(ATTRIBUTE_VALUE_XHTML *ATTRIBUTE_VALUE_XHTML)
	DeleteORMDATATYPE_DEFINITION_BOOLEAN(DATATYPE_DEFINITION_BOOLEAN *DATATYPE_DEFINITION_BOOLEAN)
	DeleteORMDATATYPE_DEFINITION_DATE(DATATYPE_DEFINITION_DATE *DATATYPE_DEFINITION_DATE)
	DeleteORMDATATYPE_DEFINITION_ENUMERATION(DATATYPE_DEFINITION_ENUMERATION *DATATYPE_DEFINITION_ENUMERATION)
	DeleteORMDATATYPE_DEFINITION_INTEGER(DATATYPE_DEFINITION_INTEGER *DATATYPE_DEFINITION_INTEGER)
	DeleteORMDATATYPE_DEFINITION_REAL(DATATYPE_DEFINITION_REAL *DATATYPE_DEFINITION_REAL)
	DeleteORMDATATYPE_DEFINITION_STRING(DATATYPE_DEFINITION_STRING *DATATYPE_DEFINITION_STRING)
	DeleteORMDATATYPE_DEFINITION_XHTML(DATATYPE_DEFINITION_XHTML *DATATYPE_DEFINITION_XHTML)
	DeleteORMEMBEDDED_VALUE(EMBEDDED_VALUE *EMBEDDED_VALUE)
	DeleteORMENUM_VALUE(ENUM_VALUE *ENUM_VALUE)
	DeleteORMRELATION_GROUP(RELATION_GROUP *RELATION_GROUP)
	DeleteORMRELATION_GROUP_TYPE(RELATION_GROUP_TYPE *RELATION_GROUP_TYPE)
	DeleteORMREQ_IF(REQ_IF *REQ_IF)
	DeleteORMREQ_IF_CONTENT(REQ_IF_CONTENT *REQ_IF_CONTENT)
	DeleteORMREQ_IF_HEADER(REQ_IF_HEADER *REQ_IF_HEADER)
	DeleteORMREQ_IF_TOOL_EXTENSION(REQ_IF_TOOL_EXTENSION *REQ_IF_TOOL_EXTENSION)
	DeleteORMSPECIFICATION(SPECIFICATION *SPECIFICATION)
	DeleteORMSPECIFICATION_TYPE(SPECIFICATION_TYPE *SPECIFICATION_TYPE)
	DeleteORMSPEC_HIERARCHY(SPEC_HIERARCHY *SPEC_HIERARCHY)
	DeleteORMSPEC_OBJECT(SPEC_OBJECT *SPEC_OBJECT)
	DeleteORMSPEC_OBJECT_TYPE(SPEC_OBJECT_TYPE *SPEC_OBJECT_TYPE)
	DeleteORMSPEC_RELATION(SPEC_RELATION *SPEC_RELATION)
	DeleteORMSPEC_RELATION_TYPE(SPEC_RELATION_TYPE *SPEC_RELATION_TYPE)
	DeleteORMXHTML_CONTENT(XHTML_CONTENT *XHTML_CONTENT)
}

func (stage *StageStruct) Reset() { // insertion point for array reset
	stage.ALTERNATIVE_IDs = make(map[*ALTERNATIVE_ID]any)
	stage.ALTERNATIVE_IDs_mapString = make(map[string]*ALTERNATIVE_ID)

	stage.ATTRIBUTE_DEFINITION_BOOLEANs = make(map[*ATTRIBUTE_DEFINITION_BOOLEAN]any)
	stage.ATTRIBUTE_DEFINITION_BOOLEANs_mapString = make(map[string]*ATTRIBUTE_DEFINITION_BOOLEAN)

	stage.ATTRIBUTE_DEFINITION_DATEs = make(map[*ATTRIBUTE_DEFINITION_DATE]any)
	stage.ATTRIBUTE_DEFINITION_DATEs_mapString = make(map[string]*ATTRIBUTE_DEFINITION_DATE)

	stage.ATTRIBUTE_DEFINITION_ENUMERATIONs = make(map[*ATTRIBUTE_DEFINITION_ENUMERATION]any)
	stage.ATTRIBUTE_DEFINITION_ENUMERATIONs_mapString = make(map[string]*ATTRIBUTE_DEFINITION_ENUMERATION)

	stage.ATTRIBUTE_DEFINITION_INTEGERs = make(map[*ATTRIBUTE_DEFINITION_INTEGER]any)
	stage.ATTRIBUTE_DEFINITION_INTEGERs_mapString = make(map[string]*ATTRIBUTE_DEFINITION_INTEGER)

	stage.ATTRIBUTE_DEFINITION_REALs = make(map[*ATTRIBUTE_DEFINITION_REAL]any)
	stage.ATTRIBUTE_DEFINITION_REALs_mapString = make(map[string]*ATTRIBUTE_DEFINITION_REAL)

	stage.ATTRIBUTE_DEFINITION_STRINGs = make(map[*ATTRIBUTE_DEFINITION_STRING]any)
	stage.ATTRIBUTE_DEFINITION_STRINGs_mapString = make(map[string]*ATTRIBUTE_DEFINITION_STRING)

	stage.ATTRIBUTE_DEFINITION_XHTMLs = make(map[*ATTRIBUTE_DEFINITION_XHTML]any)
	stage.ATTRIBUTE_DEFINITION_XHTMLs_mapString = make(map[string]*ATTRIBUTE_DEFINITION_XHTML)

	stage.ATTRIBUTE_VALUE_BOOLEANs = make(map[*ATTRIBUTE_VALUE_BOOLEAN]any)
	stage.ATTRIBUTE_VALUE_BOOLEANs_mapString = make(map[string]*ATTRIBUTE_VALUE_BOOLEAN)

	stage.ATTRIBUTE_VALUE_DATEs = make(map[*ATTRIBUTE_VALUE_DATE]any)
	stage.ATTRIBUTE_VALUE_DATEs_mapString = make(map[string]*ATTRIBUTE_VALUE_DATE)

	stage.ATTRIBUTE_VALUE_ENUMERATIONs = make(map[*ATTRIBUTE_VALUE_ENUMERATION]any)
	stage.ATTRIBUTE_VALUE_ENUMERATIONs_mapString = make(map[string]*ATTRIBUTE_VALUE_ENUMERATION)

	stage.ATTRIBUTE_VALUE_INTEGERs = make(map[*ATTRIBUTE_VALUE_INTEGER]any)
	stage.ATTRIBUTE_VALUE_INTEGERs_mapString = make(map[string]*ATTRIBUTE_VALUE_INTEGER)

	stage.ATTRIBUTE_VALUE_REALs = make(map[*ATTRIBUTE_VALUE_REAL]any)
	stage.ATTRIBUTE_VALUE_REALs_mapString = make(map[string]*ATTRIBUTE_VALUE_REAL)

	stage.ATTRIBUTE_VALUE_STRINGs = make(map[*ATTRIBUTE_VALUE_STRING]any)
	stage.ATTRIBUTE_VALUE_STRINGs_mapString = make(map[string]*ATTRIBUTE_VALUE_STRING)

	stage.ATTRIBUTE_VALUE_XHTMLs = make(map[*ATTRIBUTE_VALUE_XHTML]any)
	stage.ATTRIBUTE_VALUE_XHTMLs_mapString = make(map[string]*ATTRIBUTE_VALUE_XHTML)

	stage.DATATYPE_DEFINITION_BOOLEANs = make(map[*DATATYPE_DEFINITION_BOOLEAN]any)
	stage.DATATYPE_DEFINITION_BOOLEANs_mapString = make(map[string]*DATATYPE_DEFINITION_BOOLEAN)

	stage.DATATYPE_DEFINITION_DATEs = make(map[*DATATYPE_DEFINITION_DATE]any)
	stage.DATATYPE_DEFINITION_DATEs_mapString = make(map[string]*DATATYPE_DEFINITION_DATE)

	stage.DATATYPE_DEFINITION_ENUMERATIONs = make(map[*DATATYPE_DEFINITION_ENUMERATION]any)
	stage.DATATYPE_DEFINITION_ENUMERATIONs_mapString = make(map[string]*DATATYPE_DEFINITION_ENUMERATION)

	stage.DATATYPE_DEFINITION_INTEGERs = make(map[*DATATYPE_DEFINITION_INTEGER]any)
	stage.DATATYPE_DEFINITION_INTEGERs_mapString = make(map[string]*DATATYPE_DEFINITION_INTEGER)

	stage.DATATYPE_DEFINITION_REALs = make(map[*DATATYPE_DEFINITION_REAL]any)
	stage.DATATYPE_DEFINITION_REALs_mapString = make(map[string]*DATATYPE_DEFINITION_REAL)

	stage.DATATYPE_DEFINITION_STRINGs = make(map[*DATATYPE_DEFINITION_STRING]any)
	stage.DATATYPE_DEFINITION_STRINGs_mapString = make(map[string]*DATATYPE_DEFINITION_STRING)

	stage.DATATYPE_DEFINITION_XHTMLs = make(map[*DATATYPE_DEFINITION_XHTML]any)
	stage.DATATYPE_DEFINITION_XHTMLs_mapString = make(map[string]*DATATYPE_DEFINITION_XHTML)

	stage.EMBEDDED_VALUEs = make(map[*EMBEDDED_VALUE]any)
	stage.EMBEDDED_VALUEs_mapString = make(map[string]*EMBEDDED_VALUE)

	stage.ENUM_VALUEs = make(map[*ENUM_VALUE]any)
	stage.ENUM_VALUEs_mapString = make(map[string]*ENUM_VALUE)

	stage.RELATION_GROUPs = make(map[*RELATION_GROUP]any)
	stage.RELATION_GROUPs_mapString = make(map[string]*RELATION_GROUP)

	stage.RELATION_GROUP_TYPEs = make(map[*RELATION_GROUP_TYPE]any)
	stage.RELATION_GROUP_TYPEs_mapString = make(map[string]*RELATION_GROUP_TYPE)

	stage.REQ_IFs = make(map[*REQ_IF]any)
	stage.REQ_IFs_mapString = make(map[string]*REQ_IF)

	stage.REQ_IF_CONTENTs = make(map[*REQ_IF_CONTENT]any)
	stage.REQ_IF_CONTENTs_mapString = make(map[string]*REQ_IF_CONTENT)

	stage.REQ_IF_HEADERs = make(map[*REQ_IF_HEADER]any)
	stage.REQ_IF_HEADERs_mapString = make(map[string]*REQ_IF_HEADER)

	stage.REQ_IF_TOOL_EXTENSIONs = make(map[*REQ_IF_TOOL_EXTENSION]any)
	stage.REQ_IF_TOOL_EXTENSIONs_mapString = make(map[string]*REQ_IF_TOOL_EXTENSION)

	stage.SPECIFICATIONs = make(map[*SPECIFICATION]any)
	stage.SPECIFICATIONs_mapString = make(map[string]*SPECIFICATION)

	stage.SPECIFICATION_TYPEs = make(map[*SPECIFICATION_TYPE]any)
	stage.SPECIFICATION_TYPEs_mapString = make(map[string]*SPECIFICATION_TYPE)

	stage.SPEC_HIERARCHYs = make(map[*SPEC_HIERARCHY]any)
	stage.SPEC_HIERARCHYs_mapString = make(map[string]*SPEC_HIERARCHY)

	stage.SPEC_OBJECTs = make(map[*SPEC_OBJECT]any)
	stage.SPEC_OBJECTs_mapString = make(map[string]*SPEC_OBJECT)

	stage.SPEC_OBJECT_TYPEs = make(map[*SPEC_OBJECT_TYPE]any)
	stage.SPEC_OBJECT_TYPEs_mapString = make(map[string]*SPEC_OBJECT_TYPE)

	stage.SPEC_RELATIONs = make(map[*SPEC_RELATION]any)
	stage.SPEC_RELATIONs_mapString = make(map[string]*SPEC_RELATION)

	stage.SPEC_RELATION_TYPEs = make(map[*SPEC_RELATION_TYPE]any)
	stage.SPEC_RELATION_TYPEs_mapString = make(map[string]*SPEC_RELATION_TYPE)

	stage.XHTML_CONTENTs = make(map[*XHTML_CONTENT]any)
	stage.XHTML_CONTENTs_mapString = make(map[string]*XHTML_CONTENT)

}

func (stage *StageStruct) Nil() { // insertion point for array nil
	stage.ALTERNATIVE_IDs = nil
	stage.ALTERNATIVE_IDs_mapString = nil

	stage.ATTRIBUTE_DEFINITION_BOOLEANs = nil
	stage.ATTRIBUTE_DEFINITION_BOOLEANs_mapString = nil

	stage.ATTRIBUTE_DEFINITION_DATEs = nil
	stage.ATTRIBUTE_DEFINITION_DATEs_mapString = nil

	stage.ATTRIBUTE_DEFINITION_ENUMERATIONs = nil
	stage.ATTRIBUTE_DEFINITION_ENUMERATIONs_mapString = nil

	stage.ATTRIBUTE_DEFINITION_INTEGERs = nil
	stage.ATTRIBUTE_DEFINITION_INTEGERs_mapString = nil

	stage.ATTRIBUTE_DEFINITION_REALs = nil
	stage.ATTRIBUTE_DEFINITION_REALs_mapString = nil

	stage.ATTRIBUTE_DEFINITION_STRINGs = nil
	stage.ATTRIBUTE_DEFINITION_STRINGs_mapString = nil

	stage.ATTRIBUTE_DEFINITION_XHTMLs = nil
	stage.ATTRIBUTE_DEFINITION_XHTMLs_mapString = nil

	stage.ATTRIBUTE_VALUE_BOOLEANs = nil
	stage.ATTRIBUTE_VALUE_BOOLEANs_mapString = nil

	stage.ATTRIBUTE_VALUE_DATEs = nil
	stage.ATTRIBUTE_VALUE_DATEs_mapString = nil

	stage.ATTRIBUTE_VALUE_ENUMERATIONs = nil
	stage.ATTRIBUTE_VALUE_ENUMERATIONs_mapString = nil

	stage.ATTRIBUTE_VALUE_INTEGERs = nil
	stage.ATTRIBUTE_VALUE_INTEGERs_mapString = nil

	stage.ATTRIBUTE_VALUE_REALs = nil
	stage.ATTRIBUTE_VALUE_REALs_mapString = nil

	stage.ATTRIBUTE_VALUE_STRINGs = nil
	stage.ATTRIBUTE_VALUE_STRINGs_mapString = nil

	stage.ATTRIBUTE_VALUE_XHTMLs = nil
	stage.ATTRIBUTE_VALUE_XHTMLs_mapString = nil

	stage.DATATYPE_DEFINITION_BOOLEANs = nil
	stage.DATATYPE_DEFINITION_BOOLEANs_mapString = nil

	stage.DATATYPE_DEFINITION_DATEs = nil
	stage.DATATYPE_DEFINITION_DATEs_mapString = nil

	stage.DATATYPE_DEFINITION_ENUMERATIONs = nil
	stage.DATATYPE_DEFINITION_ENUMERATIONs_mapString = nil

	stage.DATATYPE_DEFINITION_INTEGERs = nil
	stage.DATATYPE_DEFINITION_INTEGERs_mapString = nil

	stage.DATATYPE_DEFINITION_REALs = nil
	stage.DATATYPE_DEFINITION_REALs_mapString = nil

	stage.DATATYPE_DEFINITION_STRINGs = nil
	stage.DATATYPE_DEFINITION_STRINGs_mapString = nil

	stage.DATATYPE_DEFINITION_XHTMLs = nil
	stage.DATATYPE_DEFINITION_XHTMLs_mapString = nil

	stage.EMBEDDED_VALUEs = nil
	stage.EMBEDDED_VALUEs_mapString = nil

	stage.ENUM_VALUEs = nil
	stage.ENUM_VALUEs_mapString = nil

	stage.RELATION_GROUPs = nil
	stage.RELATION_GROUPs_mapString = nil

	stage.RELATION_GROUP_TYPEs = nil
	stage.RELATION_GROUP_TYPEs_mapString = nil

	stage.REQ_IFs = nil
	stage.REQ_IFs_mapString = nil

	stage.REQ_IF_CONTENTs = nil
	stage.REQ_IF_CONTENTs_mapString = nil

	stage.REQ_IF_HEADERs = nil
	stage.REQ_IF_HEADERs_mapString = nil

	stage.REQ_IF_TOOL_EXTENSIONs = nil
	stage.REQ_IF_TOOL_EXTENSIONs_mapString = nil

	stage.SPECIFICATIONs = nil
	stage.SPECIFICATIONs_mapString = nil

	stage.SPECIFICATION_TYPEs = nil
	stage.SPECIFICATION_TYPEs_mapString = nil

	stage.SPEC_HIERARCHYs = nil
	stage.SPEC_HIERARCHYs_mapString = nil

	stage.SPEC_OBJECTs = nil
	stage.SPEC_OBJECTs_mapString = nil

	stage.SPEC_OBJECT_TYPEs = nil
	stage.SPEC_OBJECT_TYPEs_mapString = nil

	stage.SPEC_RELATIONs = nil
	stage.SPEC_RELATIONs_mapString = nil

	stage.SPEC_RELATION_TYPEs = nil
	stage.SPEC_RELATION_TYPEs_mapString = nil

	stage.XHTML_CONTENTs = nil
	stage.XHTML_CONTENTs_mapString = nil

}

func (stage *StageStruct) Unstage() { // insertion point for array nil
	for alternative_id := range stage.ALTERNATIVE_IDs {
		alternative_id.Unstage(stage)
	}

	for attribute_definition_boolean := range stage.ATTRIBUTE_DEFINITION_BOOLEANs {
		attribute_definition_boolean.Unstage(stage)
	}

	for attribute_definition_date := range stage.ATTRIBUTE_DEFINITION_DATEs {
		attribute_definition_date.Unstage(stage)
	}

	for attribute_definition_enumeration := range stage.ATTRIBUTE_DEFINITION_ENUMERATIONs {
		attribute_definition_enumeration.Unstage(stage)
	}

	for attribute_definition_integer := range stage.ATTRIBUTE_DEFINITION_INTEGERs {
		attribute_definition_integer.Unstage(stage)
	}

	for attribute_definition_real := range stage.ATTRIBUTE_DEFINITION_REALs {
		attribute_definition_real.Unstage(stage)
	}

	for attribute_definition_string := range stage.ATTRIBUTE_DEFINITION_STRINGs {
		attribute_definition_string.Unstage(stage)
	}

	for attribute_definition_xhtml := range stage.ATTRIBUTE_DEFINITION_XHTMLs {
		attribute_definition_xhtml.Unstage(stage)
	}

	for attribute_value_boolean := range stage.ATTRIBUTE_VALUE_BOOLEANs {
		attribute_value_boolean.Unstage(stage)
	}

	for attribute_value_date := range stage.ATTRIBUTE_VALUE_DATEs {
		attribute_value_date.Unstage(stage)
	}

	for attribute_value_enumeration := range stage.ATTRIBUTE_VALUE_ENUMERATIONs {
		attribute_value_enumeration.Unstage(stage)
	}

	for attribute_value_integer := range stage.ATTRIBUTE_VALUE_INTEGERs {
		attribute_value_integer.Unstage(stage)
	}

	for attribute_value_real := range stage.ATTRIBUTE_VALUE_REALs {
		attribute_value_real.Unstage(stage)
	}

	for attribute_value_string := range stage.ATTRIBUTE_VALUE_STRINGs {
		attribute_value_string.Unstage(stage)
	}

	for attribute_value_xhtml := range stage.ATTRIBUTE_VALUE_XHTMLs {
		attribute_value_xhtml.Unstage(stage)
	}

	for datatype_definition_boolean := range stage.DATATYPE_DEFINITION_BOOLEANs {
		datatype_definition_boolean.Unstage(stage)
	}

	for datatype_definition_date := range stage.DATATYPE_DEFINITION_DATEs {
		datatype_definition_date.Unstage(stage)
	}

	for datatype_definition_enumeration := range stage.DATATYPE_DEFINITION_ENUMERATIONs {
		datatype_definition_enumeration.Unstage(stage)
	}

	for datatype_definition_integer := range stage.DATATYPE_DEFINITION_INTEGERs {
		datatype_definition_integer.Unstage(stage)
	}

	for datatype_definition_real := range stage.DATATYPE_DEFINITION_REALs {
		datatype_definition_real.Unstage(stage)
	}

	for datatype_definition_string := range stage.DATATYPE_DEFINITION_STRINGs {
		datatype_definition_string.Unstage(stage)
	}

	for datatype_definition_xhtml := range stage.DATATYPE_DEFINITION_XHTMLs {
		datatype_definition_xhtml.Unstage(stage)
	}

	for embedded_value := range stage.EMBEDDED_VALUEs {
		embedded_value.Unstage(stage)
	}

	for enum_value := range stage.ENUM_VALUEs {
		enum_value.Unstage(stage)
	}

	for relation_group := range stage.RELATION_GROUPs {
		relation_group.Unstage(stage)
	}

	for relation_group_type := range stage.RELATION_GROUP_TYPEs {
		relation_group_type.Unstage(stage)
	}

	for req_if := range stage.REQ_IFs {
		req_if.Unstage(stage)
	}

	for req_if_content := range stage.REQ_IF_CONTENTs {
		req_if_content.Unstage(stage)
	}

	for req_if_header := range stage.REQ_IF_HEADERs {
		req_if_header.Unstage(stage)
	}

	for req_if_tool_extension := range stage.REQ_IF_TOOL_EXTENSIONs {
		req_if_tool_extension.Unstage(stage)
	}

	for specification := range stage.SPECIFICATIONs {
		specification.Unstage(stage)
	}

	for specification_type := range stage.SPECIFICATION_TYPEs {
		specification_type.Unstage(stage)
	}

	for spec_hierarchy := range stage.SPEC_HIERARCHYs {
		spec_hierarchy.Unstage(stage)
	}

	for spec_object := range stage.SPEC_OBJECTs {
		spec_object.Unstage(stage)
	}

	for spec_object_type := range stage.SPEC_OBJECT_TYPEs {
		spec_object_type.Unstage(stage)
	}

	for spec_relation := range stage.SPEC_RELATIONs {
		spec_relation.Unstage(stage)
	}

	for spec_relation_type := range stage.SPEC_RELATION_TYPEs {
		spec_relation_type.Unstage(stage)
	}

	for xhtml_content := range stage.XHTML_CONTENTs {
		xhtml_content.Unstage(stage)
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
	case map[*ALTERNATIVE_ID]any:
		return any(&stage.ALTERNATIVE_IDs).(*Type)
	case map[*ATTRIBUTE_DEFINITION_BOOLEAN]any:
		return any(&stage.ATTRIBUTE_DEFINITION_BOOLEANs).(*Type)
	case map[*ATTRIBUTE_DEFINITION_DATE]any:
		return any(&stage.ATTRIBUTE_DEFINITION_DATEs).(*Type)
	case map[*ATTRIBUTE_DEFINITION_ENUMERATION]any:
		return any(&stage.ATTRIBUTE_DEFINITION_ENUMERATIONs).(*Type)
	case map[*ATTRIBUTE_DEFINITION_INTEGER]any:
		return any(&stage.ATTRIBUTE_DEFINITION_INTEGERs).(*Type)
	case map[*ATTRIBUTE_DEFINITION_REAL]any:
		return any(&stage.ATTRIBUTE_DEFINITION_REALs).(*Type)
	case map[*ATTRIBUTE_DEFINITION_STRING]any:
		return any(&stage.ATTRIBUTE_DEFINITION_STRINGs).(*Type)
	case map[*ATTRIBUTE_DEFINITION_XHTML]any:
		return any(&stage.ATTRIBUTE_DEFINITION_XHTMLs).(*Type)
	case map[*ATTRIBUTE_VALUE_BOOLEAN]any:
		return any(&stage.ATTRIBUTE_VALUE_BOOLEANs).(*Type)
	case map[*ATTRIBUTE_VALUE_DATE]any:
		return any(&stage.ATTRIBUTE_VALUE_DATEs).(*Type)
	case map[*ATTRIBUTE_VALUE_ENUMERATION]any:
		return any(&stage.ATTRIBUTE_VALUE_ENUMERATIONs).(*Type)
	case map[*ATTRIBUTE_VALUE_INTEGER]any:
		return any(&stage.ATTRIBUTE_VALUE_INTEGERs).(*Type)
	case map[*ATTRIBUTE_VALUE_REAL]any:
		return any(&stage.ATTRIBUTE_VALUE_REALs).(*Type)
	case map[*ATTRIBUTE_VALUE_STRING]any:
		return any(&stage.ATTRIBUTE_VALUE_STRINGs).(*Type)
	case map[*ATTRIBUTE_VALUE_XHTML]any:
		return any(&stage.ATTRIBUTE_VALUE_XHTMLs).(*Type)
	case map[*DATATYPE_DEFINITION_BOOLEAN]any:
		return any(&stage.DATATYPE_DEFINITION_BOOLEANs).(*Type)
	case map[*DATATYPE_DEFINITION_DATE]any:
		return any(&stage.DATATYPE_DEFINITION_DATEs).(*Type)
	case map[*DATATYPE_DEFINITION_ENUMERATION]any:
		return any(&stage.DATATYPE_DEFINITION_ENUMERATIONs).(*Type)
	case map[*DATATYPE_DEFINITION_INTEGER]any:
		return any(&stage.DATATYPE_DEFINITION_INTEGERs).(*Type)
	case map[*DATATYPE_DEFINITION_REAL]any:
		return any(&stage.DATATYPE_DEFINITION_REALs).(*Type)
	case map[*DATATYPE_DEFINITION_STRING]any:
		return any(&stage.DATATYPE_DEFINITION_STRINGs).(*Type)
	case map[*DATATYPE_DEFINITION_XHTML]any:
		return any(&stage.DATATYPE_DEFINITION_XHTMLs).(*Type)
	case map[*EMBEDDED_VALUE]any:
		return any(&stage.EMBEDDED_VALUEs).(*Type)
	case map[*ENUM_VALUE]any:
		return any(&stage.ENUM_VALUEs).(*Type)
	case map[*RELATION_GROUP]any:
		return any(&stage.RELATION_GROUPs).(*Type)
	case map[*RELATION_GROUP_TYPE]any:
		return any(&stage.RELATION_GROUP_TYPEs).(*Type)
	case map[*REQ_IF]any:
		return any(&stage.REQ_IFs).(*Type)
	case map[*REQ_IF_CONTENT]any:
		return any(&stage.REQ_IF_CONTENTs).(*Type)
	case map[*REQ_IF_HEADER]any:
		return any(&stage.REQ_IF_HEADERs).(*Type)
	case map[*REQ_IF_TOOL_EXTENSION]any:
		return any(&stage.REQ_IF_TOOL_EXTENSIONs).(*Type)
	case map[*SPECIFICATION]any:
		return any(&stage.SPECIFICATIONs).(*Type)
	case map[*SPECIFICATION_TYPE]any:
		return any(&stage.SPECIFICATION_TYPEs).(*Type)
	case map[*SPEC_HIERARCHY]any:
		return any(&stage.SPEC_HIERARCHYs).(*Type)
	case map[*SPEC_OBJECT]any:
		return any(&stage.SPEC_OBJECTs).(*Type)
	case map[*SPEC_OBJECT_TYPE]any:
		return any(&stage.SPEC_OBJECT_TYPEs).(*Type)
	case map[*SPEC_RELATION]any:
		return any(&stage.SPEC_RELATIONs).(*Type)
	case map[*SPEC_RELATION_TYPE]any:
		return any(&stage.SPEC_RELATION_TYPEs).(*Type)
	case map[*XHTML_CONTENT]any:
		return any(&stage.XHTML_CONTENTs).(*Type)
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
	case map[string]*ALTERNATIVE_ID:
		return any(&stage.ALTERNATIVE_IDs_mapString).(*Type)
	case map[string]*ATTRIBUTE_DEFINITION_BOOLEAN:
		return any(&stage.ATTRIBUTE_DEFINITION_BOOLEANs_mapString).(*Type)
	case map[string]*ATTRIBUTE_DEFINITION_DATE:
		return any(&stage.ATTRIBUTE_DEFINITION_DATEs_mapString).(*Type)
	case map[string]*ATTRIBUTE_DEFINITION_ENUMERATION:
		return any(&stage.ATTRIBUTE_DEFINITION_ENUMERATIONs_mapString).(*Type)
	case map[string]*ATTRIBUTE_DEFINITION_INTEGER:
		return any(&stage.ATTRIBUTE_DEFINITION_INTEGERs_mapString).(*Type)
	case map[string]*ATTRIBUTE_DEFINITION_REAL:
		return any(&stage.ATTRIBUTE_DEFINITION_REALs_mapString).(*Type)
	case map[string]*ATTRIBUTE_DEFINITION_STRING:
		return any(&stage.ATTRIBUTE_DEFINITION_STRINGs_mapString).(*Type)
	case map[string]*ATTRIBUTE_DEFINITION_XHTML:
		return any(&stage.ATTRIBUTE_DEFINITION_XHTMLs_mapString).(*Type)
	case map[string]*ATTRIBUTE_VALUE_BOOLEAN:
		return any(&stage.ATTRIBUTE_VALUE_BOOLEANs_mapString).(*Type)
	case map[string]*ATTRIBUTE_VALUE_DATE:
		return any(&stage.ATTRIBUTE_VALUE_DATEs_mapString).(*Type)
	case map[string]*ATTRIBUTE_VALUE_ENUMERATION:
		return any(&stage.ATTRIBUTE_VALUE_ENUMERATIONs_mapString).(*Type)
	case map[string]*ATTRIBUTE_VALUE_INTEGER:
		return any(&stage.ATTRIBUTE_VALUE_INTEGERs_mapString).(*Type)
	case map[string]*ATTRIBUTE_VALUE_REAL:
		return any(&stage.ATTRIBUTE_VALUE_REALs_mapString).(*Type)
	case map[string]*ATTRIBUTE_VALUE_STRING:
		return any(&stage.ATTRIBUTE_VALUE_STRINGs_mapString).(*Type)
	case map[string]*ATTRIBUTE_VALUE_XHTML:
		return any(&stage.ATTRIBUTE_VALUE_XHTMLs_mapString).(*Type)
	case map[string]*DATATYPE_DEFINITION_BOOLEAN:
		return any(&stage.DATATYPE_DEFINITION_BOOLEANs_mapString).(*Type)
	case map[string]*DATATYPE_DEFINITION_DATE:
		return any(&stage.DATATYPE_DEFINITION_DATEs_mapString).(*Type)
	case map[string]*DATATYPE_DEFINITION_ENUMERATION:
		return any(&stage.DATATYPE_DEFINITION_ENUMERATIONs_mapString).(*Type)
	case map[string]*DATATYPE_DEFINITION_INTEGER:
		return any(&stage.DATATYPE_DEFINITION_INTEGERs_mapString).(*Type)
	case map[string]*DATATYPE_DEFINITION_REAL:
		return any(&stage.DATATYPE_DEFINITION_REALs_mapString).(*Type)
	case map[string]*DATATYPE_DEFINITION_STRING:
		return any(&stage.DATATYPE_DEFINITION_STRINGs_mapString).(*Type)
	case map[string]*DATATYPE_DEFINITION_XHTML:
		return any(&stage.DATATYPE_DEFINITION_XHTMLs_mapString).(*Type)
	case map[string]*EMBEDDED_VALUE:
		return any(&stage.EMBEDDED_VALUEs_mapString).(*Type)
	case map[string]*ENUM_VALUE:
		return any(&stage.ENUM_VALUEs_mapString).(*Type)
	case map[string]*RELATION_GROUP:
		return any(&stage.RELATION_GROUPs_mapString).(*Type)
	case map[string]*RELATION_GROUP_TYPE:
		return any(&stage.RELATION_GROUP_TYPEs_mapString).(*Type)
	case map[string]*REQ_IF:
		return any(&stage.REQ_IFs_mapString).(*Type)
	case map[string]*REQ_IF_CONTENT:
		return any(&stage.REQ_IF_CONTENTs_mapString).(*Type)
	case map[string]*REQ_IF_HEADER:
		return any(&stage.REQ_IF_HEADERs_mapString).(*Type)
	case map[string]*REQ_IF_TOOL_EXTENSION:
		return any(&stage.REQ_IF_TOOL_EXTENSIONs_mapString).(*Type)
	case map[string]*SPECIFICATION:
		return any(&stage.SPECIFICATIONs_mapString).(*Type)
	case map[string]*SPECIFICATION_TYPE:
		return any(&stage.SPECIFICATION_TYPEs_mapString).(*Type)
	case map[string]*SPEC_HIERARCHY:
		return any(&stage.SPEC_HIERARCHYs_mapString).(*Type)
	case map[string]*SPEC_OBJECT:
		return any(&stage.SPEC_OBJECTs_mapString).(*Type)
	case map[string]*SPEC_OBJECT_TYPE:
		return any(&stage.SPEC_OBJECT_TYPEs_mapString).(*Type)
	case map[string]*SPEC_RELATION:
		return any(&stage.SPEC_RELATIONs_mapString).(*Type)
	case map[string]*SPEC_RELATION_TYPE:
		return any(&stage.SPEC_RELATION_TYPEs_mapString).(*Type)
	case map[string]*XHTML_CONTENT:
		return any(&stage.XHTML_CONTENTs_mapString).(*Type)
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
	case ALTERNATIVE_ID:
		return any(&stage.ALTERNATIVE_IDs).(*map[*Type]any)
	case ATTRIBUTE_DEFINITION_BOOLEAN:
		return any(&stage.ATTRIBUTE_DEFINITION_BOOLEANs).(*map[*Type]any)
	case ATTRIBUTE_DEFINITION_DATE:
		return any(&stage.ATTRIBUTE_DEFINITION_DATEs).(*map[*Type]any)
	case ATTRIBUTE_DEFINITION_ENUMERATION:
		return any(&stage.ATTRIBUTE_DEFINITION_ENUMERATIONs).(*map[*Type]any)
	case ATTRIBUTE_DEFINITION_INTEGER:
		return any(&stage.ATTRIBUTE_DEFINITION_INTEGERs).(*map[*Type]any)
	case ATTRIBUTE_DEFINITION_REAL:
		return any(&stage.ATTRIBUTE_DEFINITION_REALs).(*map[*Type]any)
	case ATTRIBUTE_DEFINITION_STRING:
		return any(&stage.ATTRIBUTE_DEFINITION_STRINGs).(*map[*Type]any)
	case ATTRIBUTE_DEFINITION_XHTML:
		return any(&stage.ATTRIBUTE_DEFINITION_XHTMLs).(*map[*Type]any)
	case ATTRIBUTE_VALUE_BOOLEAN:
		return any(&stage.ATTRIBUTE_VALUE_BOOLEANs).(*map[*Type]any)
	case ATTRIBUTE_VALUE_DATE:
		return any(&stage.ATTRIBUTE_VALUE_DATEs).(*map[*Type]any)
	case ATTRIBUTE_VALUE_ENUMERATION:
		return any(&stage.ATTRIBUTE_VALUE_ENUMERATIONs).(*map[*Type]any)
	case ATTRIBUTE_VALUE_INTEGER:
		return any(&stage.ATTRIBUTE_VALUE_INTEGERs).(*map[*Type]any)
	case ATTRIBUTE_VALUE_REAL:
		return any(&stage.ATTRIBUTE_VALUE_REALs).(*map[*Type]any)
	case ATTRIBUTE_VALUE_STRING:
		return any(&stage.ATTRIBUTE_VALUE_STRINGs).(*map[*Type]any)
	case ATTRIBUTE_VALUE_XHTML:
		return any(&stage.ATTRIBUTE_VALUE_XHTMLs).(*map[*Type]any)
	case DATATYPE_DEFINITION_BOOLEAN:
		return any(&stage.DATATYPE_DEFINITION_BOOLEANs).(*map[*Type]any)
	case DATATYPE_DEFINITION_DATE:
		return any(&stage.DATATYPE_DEFINITION_DATEs).(*map[*Type]any)
	case DATATYPE_DEFINITION_ENUMERATION:
		return any(&stage.DATATYPE_DEFINITION_ENUMERATIONs).(*map[*Type]any)
	case DATATYPE_DEFINITION_INTEGER:
		return any(&stage.DATATYPE_DEFINITION_INTEGERs).(*map[*Type]any)
	case DATATYPE_DEFINITION_REAL:
		return any(&stage.DATATYPE_DEFINITION_REALs).(*map[*Type]any)
	case DATATYPE_DEFINITION_STRING:
		return any(&stage.DATATYPE_DEFINITION_STRINGs).(*map[*Type]any)
	case DATATYPE_DEFINITION_XHTML:
		return any(&stage.DATATYPE_DEFINITION_XHTMLs).(*map[*Type]any)
	case EMBEDDED_VALUE:
		return any(&stage.EMBEDDED_VALUEs).(*map[*Type]any)
	case ENUM_VALUE:
		return any(&stage.ENUM_VALUEs).(*map[*Type]any)
	case RELATION_GROUP:
		return any(&stage.RELATION_GROUPs).(*map[*Type]any)
	case RELATION_GROUP_TYPE:
		return any(&stage.RELATION_GROUP_TYPEs).(*map[*Type]any)
	case REQ_IF:
		return any(&stage.REQ_IFs).(*map[*Type]any)
	case REQ_IF_CONTENT:
		return any(&stage.REQ_IF_CONTENTs).(*map[*Type]any)
	case REQ_IF_HEADER:
		return any(&stage.REQ_IF_HEADERs).(*map[*Type]any)
	case REQ_IF_TOOL_EXTENSION:
		return any(&stage.REQ_IF_TOOL_EXTENSIONs).(*map[*Type]any)
	case SPECIFICATION:
		return any(&stage.SPECIFICATIONs).(*map[*Type]any)
	case SPECIFICATION_TYPE:
		return any(&stage.SPECIFICATION_TYPEs).(*map[*Type]any)
	case SPEC_HIERARCHY:
		return any(&stage.SPEC_HIERARCHYs).(*map[*Type]any)
	case SPEC_OBJECT:
		return any(&stage.SPEC_OBJECTs).(*map[*Type]any)
	case SPEC_OBJECT_TYPE:
		return any(&stage.SPEC_OBJECT_TYPEs).(*map[*Type]any)
	case SPEC_RELATION:
		return any(&stage.SPEC_RELATIONs).(*map[*Type]any)
	case SPEC_RELATION_TYPE:
		return any(&stage.SPEC_RELATION_TYPEs).(*map[*Type]any)
	case XHTML_CONTENT:
		return any(&stage.XHTML_CONTENTs).(*map[*Type]any)
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
	case *ALTERNATIVE_ID:
		return any(&stage.ALTERNATIVE_IDs).(*map[Type]any)
	case *ATTRIBUTE_DEFINITION_BOOLEAN:
		return any(&stage.ATTRIBUTE_DEFINITION_BOOLEANs).(*map[Type]any)
	case *ATTRIBUTE_DEFINITION_DATE:
		return any(&stage.ATTRIBUTE_DEFINITION_DATEs).(*map[Type]any)
	case *ATTRIBUTE_DEFINITION_ENUMERATION:
		return any(&stage.ATTRIBUTE_DEFINITION_ENUMERATIONs).(*map[Type]any)
	case *ATTRIBUTE_DEFINITION_INTEGER:
		return any(&stage.ATTRIBUTE_DEFINITION_INTEGERs).(*map[Type]any)
	case *ATTRIBUTE_DEFINITION_REAL:
		return any(&stage.ATTRIBUTE_DEFINITION_REALs).(*map[Type]any)
	case *ATTRIBUTE_DEFINITION_STRING:
		return any(&stage.ATTRIBUTE_DEFINITION_STRINGs).(*map[Type]any)
	case *ATTRIBUTE_DEFINITION_XHTML:
		return any(&stage.ATTRIBUTE_DEFINITION_XHTMLs).(*map[Type]any)
	case *ATTRIBUTE_VALUE_BOOLEAN:
		return any(&stage.ATTRIBUTE_VALUE_BOOLEANs).(*map[Type]any)
	case *ATTRIBUTE_VALUE_DATE:
		return any(&stage.ATTRIBUTE_VALUE_DATEs).(*map[Type]any)
	case *ATTRIBUTE_VALUE_ENUMERATION:
		return any(&stage.ATTRIBUTE_VALUE_ENUMERATIONs).(*map[Type]any)
	case *ATTRIBUTE_VALUE_INTEGER:
		return any(&stage.ATTRIBUTE_VALUE_INTEGERs).(*map[Type]any)
	case *ATTRIBUTE_VALUE_REAL:
		return any(&stage.ATTRIBUTE_VALUE_REALs).(*map[Type]any)
	case *ATTRIBUTE_VALUE_STRING:
		return any(&stage.ATTRIBUTE_VALUE_STRINGs).(*map[Type]any)
	case *ATTRIBUTE_VALUE_XHTML:
		return any(&stage.ATTRIBUTE_VALUE_XHTMLs).(*map[Type]any)
	case *DATATYPE_DEFINITION_BOOLEAN:
		return any(&stage.DATATYPE_DEFINITION_BOOLEANs).(*map[Type]any)
	case *DATATYPE_DEFINITION_DATE:
		return any(&stage.DATATYPE_DEFINITION_DATEs).(*map[Type]any)
	case *DATATYPE_DEFINITION_ENUMERATION:
		return any(&stage.DATATYPE_DEFINITION_ENUMERATIONs).(*map[Type]any)
	case *DATATYPE_DEFINITION_INTEGER:
		return any(&stage.DATATYPE_DEFINITION_INTEGERs).(*map[Type]any)
	case *DATATYPE_DEFINITION_REAL:
		return any(&stage.DATATYPE_DEFINITION_REALs).(*map[Type]any)
	case *DATATYPE_DEFINITION_STRING:
		return any(&stage.DATATYPE_DEFINITION_STRINGs).(*map[Type]any)
	case *DATATYPE_DEFINITION_XHTML:
		return any(&stage.DATATYPE_DEFINITION_XHTMLs).(*map[Type]any)
	case *EMBEDDED_VALUE:
		return any(&stage.EMBEDDED_VALUEs).(*map[Type]any)
	case *ENUM_VALUE:
		return any(&stage.ENUM_VALUEs).(*map[Type]any)
	case *RELATION_GROUP:
		return any(&stage.RELATION_GROUPs).(*map[Type]any)
	case *RELATION_GROUP_TYPE:
		return any(&stage.RELATION_GROUP_TYPEs).(*map[Type]any)
	case *REQ_IF:
		return any(&stage.REQ_IFs).(*map[Type]any)
	case *REQ_IF_CONTENT:
		return any(&stage.REQ_IF_CONTENTs).(*map[Type]any)
	case *REQ_IF_HEADER:
		return any(&stage.REQ_IF_HEADERs).(*map[Type]any)
	case *REQ_IF_TOOL_EXTENSION:
		return any(&stage.REQ_IF_TOOL_EXTENSIONs).(*map[Type]any)
	case *SPECIFICATION:
		return any(&stage.SPECIFICATIONs).(*map[Type]any)
	case *SPECIFICATION_TYPE:
		return any(&stage.SPECIFICATION_TYPEs).(*map[Type]any)
	case *SPEC_HIERARCHY:
		return any(&stage.SPEC_HIERARCHYs).(*map[Type]any)
	case *SPEC_OBJECT:
		return any(&stage.SPEC_OBJECTs).(*map[Type]any)
	case *SPEC_OBJECT_TYPE:
		return any(&stage.SPEC_OBJECT_TYPEs).(*map[Type]any)
	case *SPEC_RELATION:
		return any(&stage.SPEC_RELATIONs).(*map[Type]any)
	case *SPEC_RELATION_TYPE:
		return any(&stage.SPEC_RELATION_TYPEs).(*map[Type]any)
	case *XHTML_CONTENT:
		return any(&stage.XHTML_CONTENTs).(*map[Type]any)
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
	case ALTERNATIVE_ID:
		return any(&stage.ALTERNATIVE_IDs_mapString).(*map[string]*Type)
	case ATTRIBUTE_DEFINITION_BOOLEAN:
		return any(&stage.ATTRIBUTE_DEFINITION_BOOLEANs_mapString).(*map[string]*Type)
	case ATTRIBUTE_DEFINITION_DATE:
		return any(&stage.ATTRIBUTE_DEFINITION_DATEs_mapString).(*map[string]*Type)
	case ATTRIBUTE_DEFINITION_ENUMERATION:
		return any(&stage.ATTRIBUTE_DEFINITION_ENUMERATIONs_mapString).(*map[string]*Type)
	case ATTRIBUTE_DEFINITION_INTEGER:
		return any(&stage.ATTRIBUTE_DEFINITION_INTEGERs_mapString).(*map[string]*Type)
	case ATTRIBUTE_DEFINITION_REAL:
		return any(&stage.ATTRIBUTE_DEFINITION_REALs_mapString).(*map[string]*Type)
	case ATTRIBUTE_DEFINITION_STRING:
		return any(&stage.ATTRIBUTE_DEFINITION_STRINGs_mapString).(*map[string]*Type)
	case ATTRIBUTE_DEFINITION_XHTML:
		return any(&stage.ATTRIBUTE_DEFINITION_XHTMLs_mapString).(*map[string]*Type)
	case ATTRIBUTE_VALUE_BOOLEAN:
		return any(&stage.ATTRIBUTE_VALUE_BOOLEANs_mapString).(*map[string]*Type)
	case ATTRIBUTE_VALUE_DATE:
		return any(&stage.ATTRIBUTE_VALUE_DATEs_mapString).(*map[string]*Type)
	case ATTRIBUTE_VALUE_ENUMERATION:
		return any(&stage.ATTRIBUTE_VALUE_ENUMERATIONs_mapString).(*map[string]*Type)
	case ATTRIBUTE_VALUE_INTEGER:
		return any(&stage.ATTRIBUTE_VALUE_INTEGERs_mapString).(*map[string]*Type)
	case ATTRIBUTE_VALUE_REAL:
		return any(&stage.ATTRIBUTE_VALUE_REALs_mapString).(*map[string]*Type)
	case ATTRIBUTE_VALUE_STRING:
		return any(&stage.ATTRIBUTE_VALUE_STRINGs_mapString).(*map[string]*Type)
	case ATTRIBUTE_VALUE_XHTML:
		return any(&stage.ATTRIBUTE_VALUE_XHTMLs_mapString).(*map[string]*Type)
	case DATATYPE_DEFINITION_BOOLEAN:
		return any(&stage.DATATYPE_DEFINITION_BOOLEANs_mapString).(*map[string]*Type)
	case DATATYPE_DEFINITION_DATE:
		return any(&stage.DATATYPE_DEFINITION_DATEs_mapString).(*map[string]*Type)
	case DATATYPE_DEFINITION_ENUMERATION:
		return any(&stage.DATATYPE_DEFINITION_ENUMERATIONs_mapString).(*map[string]*Type)
	case DATATYPE_DEFINITION_INTEGER:
		return any(&stage.DATATYPE_DEFINITION_INTEGERs_mapString).(*map[string]*Type)
	case DATATYPE_DEFINITION_REAL:
		return any(&stage.DATATYPE_DEFINITION_REALs_mapString).(*map[string]*Type)
	case DATATYPE_DEFINITION_STRING:
		return any(&stage.DATATYPE_DEFINITION_STRINGs_mapString).(*map[string]*Type)
	case DATATYPE_DEFINITION_XHTML:
		return any(&stage.DATATYPE_DEFINITION_XHTMLs_mapString).(*map[string]*Type)
	case EMBEDDED_VALUE:
		return any(&stage.EMBEDDED_VALUEs_mapString).(*map[string]*Type)
	case ENUM_VALUE:
		return any(&stage.ENUM_VALUEs_mapString).(*map[string]*Type)
	case RELATION_GROUP:
		return any(&stage.RELATION_GROUPs_mapString).(*map[string]*Type)
	case RELATION_GROUP_TYPE:
		return any(&stage.RELATION_GROUP_TYPEs_mapString).(*map[string]*Type)
	case REQ_IF:
		return any(&stage.REQ_IFs_mapString).(*map[string]*Type)
	case REQ_IF_CONTENT:
		return any(&stage.REQ_IF_CONTENTs_mapString).(*map[string]*Type)
	case REQ_IF_HEADER:
		return any(&stage.REQ_IF_HEADERs_mapString).(*map[string]*Type)
	case REQ_IF_TOOL_EXTENSION:
		return any(&stage.REQ_IF_TOOL_EXTENSIONs_mapString).(*map[string]*Type)
	case SPECIFICATION:
		return any(&stage.SPECIFICATIONs_mapString).(*map[string]*Type)
	case SPECIFICATION_TYPE:
		return any(&stage.SPECIFICATION_TYPEs_mapString).(*map[string]*Type)
	case SPEC_HIERARCHY:
		return any(&stage.SPEC_HIERARCHYs_mapString).(*map[string]*Type)
	case SPEC_OBJECT:
		return any(&stage.SPEC_OBJECTs_mapString).(*map[string]*Type)
	case SPEC_OBJECT_TYPE:
		return any(&stage.SPEC_OBJECT_TYPEs_mapString).(*map[string]*Type)
	case SPEC_RELATION:
		return any(&stage.SPEC_RELATIONs_mapString).(*map[string]*Type)
	case SPEC_RELATION_TYPE:
		return any(&stage.SPEC_RELATION_TYPEs_mapString).(*map[string]*Type)
	case XHTML_CONTENT:
		return any(&stage.XHTML_CONTENTs_mapString).(*map[string]*Type)
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
	case ALTERNATIVE_ID:
		return any(&ALTERNATIVE_ID{
			// Initialisation of associations
		}).(*Type)
	case ATTRIBUTE_DEFINITION_BOOLEAN:
		return any(&ATTRIBUTE_DEFINITION_BOOLEAN{
			// Initialisation of associations
		}).(*Type)
	case ATTRIBUTE_DEFINITION_DATE:
		return any(&ATTRIBUTE_DEFINITION_DATE{
			// Initialisation of associations
		}).(*Type)
	case ATTRIBUTE_DEFINITION_ENUMERATION:
		return any(&ATTRIBUTE_DEFINITION_ENUMERATION{
			// Initialisation of associations
		}).(*Type)
	case ATTRIBUTE_DEFINITION_INTEGER:
		return any(&ATTRIBUTE_DEFINITION_INTEGER{
			// Initialisation of associations
		}).(*Type)
	case ATTRIBUTE_DEFINITION_REAL:
		return any(&ATTRIBUTE_DEFINITION_REAL{
			// Initialisation of associations
		}).(*Type)
	case ATTRIBUTE_DEFINITION_STRING:
		return any(&ATTRIBUTE_DEFINITION_STRING{
			// Initialisation of associations
		}).(*Type)
	case ATTRIBUTE_DEFINITION_XHTML:
		return any(&ATTRIBUTE_DEFINITION_XHTML{
			// Initialisation of associations
		}).(*Type)
	case ATTRIBUTE_VALUE_BOOLEAN:
		return any(&ATTRIBUTE_VALUE_BOOLEAN{
			// Initialisation of associations
		}).(*Type)
	case ATTRIBUTE_VALUE_DATE:
		return any(&ATTRIBUTE_VALUE_DATE{
			// Initialisation of associations
		}).(*Type)
	case ATTRIBUTE_VALUE_ENUMERATION:
		return any(&ATTRIBUTE_VALUE_ENUMERATION{
			// Initialisation of associations
		}).(*Type)
	case ATTRIBUTE_VALUE_INTEGER:
		return any(&ATTRIBUTE_VALUE_INTEGER{
			// Initialisation of associations
		}).(*Type)
	case ATTRIBUTE_VALUE_REAL:
		return any(&ATTRIBUTE_VALUE_REAL{
			// Initialisation of associations
		}).(*Type)
	case ATTRIBUTE_VALUE_STRING:
		return any(&ATTRIBUTE_VALUE_STRING{
			// Initialisation of associations
		}).(*Type)
	case ATTRIBUTE_VALUE_XHTML:
		return any(&ATTRIBUTE_VALUE_XHTML{
			// Initialisation of associations
			// field is initialized with an instance of XHTML_CONTENT with the name of the field
			THE_VALUE: []*XHTML_CONTENT{{Name: "THE_VALUE"}},
			// field is initialized with an instance of XHTML_CONTENT with the name of the field
			THE_ORIGINAL_VALUE: []*XHTML_CONTENT{{Name: "THE_ORIGINAL_VALUE"}},
		}).(*Type)
	case DATATYPE_DEFINITION_BOOLEAN:
		return any(&DATATYPE_DEFINITION_BOOLEAN{
			// Initialisation of associations
		}).(*Type)
	case DATATYPE_DEFINITION_DATE:
		return any(&DATATYPE_DEFINITION_DATE{
			// Initialisation of associations
		}).(*Type)
	case DATATYPE_DEFINITION_ENUMERATION:
		return any(&DATATYPE_DEFINITION_ENUMERATION{
			// Initialisation of associations
		}).(*Type)
	case DATATYPE_DEFINITION_INTEGER:
		return any(&DATATYPE_DEFINITION_INTEGER{
			// Initialisation of associations
		}).(*Type)
	case DATATYPE_DEFINITION_REAL:
		return any(&DATATYPE_DEFINITION_REAL{
			// Initialisation of associations
		}).(*Type)
	case DATATYPE_DEFINITION_STRING:
		return any(&DATATYPE_DEFINITION_STRING{
			// Initialisation of associations
		}).(*Type)
	case DATATYPE_DEFINITION_XHTML:
		return any(&DATATYPE_DEFINITION_XHTML{
			// Initialisation of associations
		}).(*Type)
	case EMBEDDED_VALUE:
		return any(&EMBEDDED_VALUE{
			// Initialisation of associations
		}).(*Type)
	case ENUM_VALUE:
		return any(&ENUM_VALUE{
			// Initialisation of associations
		}).(*Type)
	case RELATION_GROUP:
		return any(&RELATION_GROUP{
			// Initialisation of associations
		}).(*Type)
	case RELATION_GROUP_TYPE:
		return any(&RELATION_GROUP_TYPE{
			// Initialisation of associations
		}).(*Type)
	case REQ_IF:
		return any(&REQ_IF{
			// Initialisation of associations
		}).(*Type)
	case REQ_IF_CONTENT:
		return any(&REQ_IF_CONTENT{
			// Initialisation of associations
		}).(*Type)
	case REQ_IF_HEADER:
		return any(&REQ_IF_HEADER{
			// Initialisation of associations
		}).(*Type)
	case REQ_IF_TOOL_EXTENSION:
		return any(&REQ_IF_TOOL_EXTENSION{
			// Initialisation of associations
		}).(*Type)
	case SPECIFICATION:
		return any(&SPECIFICATION{
			// Initialisation of associations
		}).(*Type)
	case SPECIFICATION_TYPE:
		return any(&SPECIFICATION_TYPE{
			// Initialisation of associations
		}).(*Type)
	case SPEC_HIERARCHY:
		return any(&SPEC_HIERARCHY{
			// Initialisation of associations
		}).(*Type)
	case SPEC_OBJECT:
		return any(&SPEC_OBJECT{
			// Initialisation of associations
		}).(*Type)
	case SPEC_OBJECT_TYPE:
		return any(&SPEC_OBJECT_TYPE{
			// Initialisation of associations
		}).(*Type)
	case SPEC_RELATION:
		return any(&SPEC_RELATION{
			// Initialisation of associations
		}).(*Type)
	case SPEC_RELATION_TYPE:
		return any(&SPEC_RELATION_TYPE{
			// Initialisation of associations
		}).(*Type)
	case XHTML_CONTENT:
		return any(&XHTML_CONTENT{
			// Initialisation of associations
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
	// reverse maps of direct associations of ALTERNATIVE_ID
	case ALTERNATIVE_ID:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ATTRIBUTE_DEFINITION_BOOLEAN
	case ATTRIBUTE_DEFINITION_BOOLEAN:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ATTRIBUTE_DEFINITION_DATE
	case ATTRIBUTE_DEFINITION_DATE:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ATTRIBUTE_DEFINITION_ENUMERATION
	case ATTRIBUTE_DEFINITION_ENUMERATION:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ATTRIBUTE_DEFINITION_INTEGER
	case ATTRIBUTE_DEFINITION_INTEGER:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ATTRIBUTE_DEFINITION_REAL
	case ATTRIBUTE_DEFINITION_REAL:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ATTRIBUTE_DEFINITION_STRING
	case ATTRIBUTE_DEFINITION_STRING:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ATTRIBUTE_DEFINITION_XHTML
	case ATTRIBUTE_DEFINITION_XHTML:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ATTRIBUTE_VALUE_BOOLEAN
	case ATTRIBUTE_VALUE_BOOLEAN:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ATTRIBUTE_VALUE_DATE
	case ATTRIBUTE_VALUE_DATE:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ATTRIBUTE_VALUE_ENUMERATION
	case ATTRIBUTE_VALUE_ENUMERATION:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ATTRIBUTE_VALUE_INTEGER
	case ATTRIBUTE_VALUE_INTEGER:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ATTRIBUTE_VALUE_REAL
	case ATTRIBUTE_VALUE_REAL:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ATTRIBUTE_VALUE_STRING
	case ATTRIBUTE_VALUE_STRING:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ATTRIBUTE_VALUE_XHTML
	case ATTRIBUTE_VALUE_XHTML:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DATATYPE_DEFINITION_BOOLEAN
	case DATATYPE_DEFINITION_BOOLEAN:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DATATYPE_DEFINITION_DATE
	case DATATYPE_DEFINITION_DATE:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DATATYPE_DEFINITION_ENUMERATION
	case DATATYPE_DEFINITION_ENUMERATION:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DATATYPE_DEFINITION_INTEGER
	case DATATYPE_DEFINITION_INTEGER:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DATATYPE_DEFINITION_REAL
	case DATATYPE_DEFINITION_REAL:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DATATYPE_DEFINITION_STRING
	case DATATYPE_DEFINITION_STRING:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DATATYPE_DEFINITION_XHTML
	case DATATYPE_DEFINITION_XHTML:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of EMBEDDED_VALUE
	case EMBEDDED_VALUE:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ENUM_VALUE
	case ENUM_VALUE:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of RELATION_GROUP
	case RELATION_GROUP:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of RELATION_GROUP_TYPE
	case RELATION_GROUP_TYPE:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of REQ_IF
	case REQ_IF:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of REQ_IF_CONTENT
	case REQ_IF_CONTENT:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of REQ_IF_HEADER
	case REQ_IF_HEADER:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of REQ_IF_TOOL_EXTENSION
	case REQ_IF_TOOL_EXTENSION:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SPECIFICATION
	case SPECIFICATION:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SPECIFICATION_TYPE
	case SPECIFICATION_TYPE:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SPEC_HIERARCHY
	case SPEC_HIERARCHY:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SPEC_OBJECT
	case SPEC_OBJECT:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SPEC_OBJECT_TYPE
	case SPEC_OBJECT_TYPE:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SPEC_RELATION
	case SPEC_RELATION:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SPEC_RELATION_TYPE
	case SPEC_RELATION_TYPE:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of XHTML_CONTENT
	case XHTML_CONTENT:
		switch fieldname {
		// insertion point for per direct association field
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
	// reverse maps of direct associations of ALTERNATIVE_ID
	case ALTERNATIVE_ID:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ATTRIBUTE_DEFINITION_BOOLEAN
	case ATTRIBUTE_DEFINITION_BOOLEAN:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ATTRIBUTE_DEFINITION_DATE
	case ATTRIBUTE_DEFINITION_DATE:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ATTRIBUTE_DEFINITION_ENUMERATION
	case ATTRIBUTE_DEFINITION_ENUMERATION:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ATTRIBUTE_DEFINITION_INTEGER
	case ATTRIBUTE_DEFINITION_INTEGER:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ATTRIBUTE_DEFINITION_REAL
	case ATTRIBUTE_DEFINITION_REAL:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ATTRIBUTE_DEFINITION_STRING
	case ATTRIBUTE_DEFINITION_STRING:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ATTRIBUTE_DEFINITION_XHTML
	case ATTRIBUTE_DEFINITION_XHTML:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ATTRIBUTE_VALUE_BOOLEAN
	case ATTRIBUTE_VALUE_BOOLEAN:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ATTRIBUTE_VALUE_DATE
	case ATTRIBUTE_VALUE_DATE:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ATTRIBUTE_VALUE_ENUMERATION
	case ATTRIBUTE_VALUE_ENUMERATION:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ATTRIBUTE_VALUE_INTEGER
	case ATTRIBUTE_VALUE_INTEGER:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ATTRIBUTE_VALUE_REAL
	case ATTRIBUTE_VALUE_REAL:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ATTRIBUTE_VALUE_STRING
	case ATTRIBUTE_VALUE_STRING:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ATTRIBUTE_VALUE_XHTML
	case ATTRIBUTE_VALUE_XHTML:
		switch fieldname {
		// insertion point for per direct association field
		case "THE_VALUE":
			res := make(map[*XHTML_CONTENT]*ATTRIBUTE_VALUE_XHTML)
			for attribute_value_xhtml := range stage.ATTRIBUTE_VALUE_XHTMLs {
				for _, xhtml_content_ := range attribute_value_xhtml.THE_VALUE {
					res[xhtml_content_] = attribute_value_xhtml
				}
			}
			return any(res).(map[*End]*Start)
		case "THE_ORIGINAL_VALUE":
			res := make(map[*XHTML_CONTENT]*ATTRIBUTE_VALUE_XHTML)
			for attribute_value_xhtml := range stage.ATTRIBUTE_VALUE_XHTMLs {
				for _, xhtml_content_ := range attribute_value_xhtml.THE_ORIGINAL_VALUE {
					res[xhtml_content_] = attribute_value_xhtml
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of DATATYPE_DEFINITION_BOOLEAN
	case DATATYPE_DEFINITION_BOOLEAN:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DATATYPE_DEFINITION_DATE
	case DATATYPE_DEFINITION_DATE:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DATATYPE_DEFINITION_ENUMERATION
	case DATATYPE_DEFINITION_ENUMERATION:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DATATYPE_DEFINITION_INTEGER
	case DATATYPE_DEFINITION_INTEGER:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DATATYPE_DEFINITION_REAL
	case DATATYPE_DEFINITION_REAL:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DATATYPE_DEFINITION_STRING
	case DATATYPE_DEFINITION_STRING:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DATATYPE_DEFINITION_XHTML
	case DATATYPE_DEFINITION_XHTML:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of EMBEDDED_VALUE
	case EMBEDDED_VALUE:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of ENUM_VALUE
	case ENUM_VALUE:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of RELATION_GROUP
	case RELATION_GROUP:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of RELATION_GROUP_TYPE
	case RELATION_GROUP_TYPE:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of REQ_IF
	case REQ_IF:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of REQ_IF_CONTENT
	case REQ_IF_CONTENT:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of REQ_IF_HEADER
	case REQ_IF_HEADER:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of REQ_IF_TOOL_EXTENSION
	case REQ_IF_TOOL_EXTENSION:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SPECIFICATION
	case SPECIFICATION:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SPECIFICATION_TYPE
	case SPECIFICATION_TYPE:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SPEC_HIERARCHY
	case SPEC_HIERARCHY:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SPEC_OBJECT
	case SPEC_OBJECT:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SPEC_OBJECT_TYPE
	case SPEC_OBJECT_TYPE:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SPEC_RELATION
	case SPEC_RELATION:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SPEC_RELATION_TYPE
	case SPEC_RELATION_TYPE:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of XHTML_CONTENT
	case XHTML_CONTENT:
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
	case ALTERNATIVE_ID:
		res = "ALTERNATIVE_ID"
	case ATTRIBUTE_DEFINITION_BOOLEAN:
		res = "ATTRIBUTE_DEFINITION_BOOLEAN"
	case ATTRIBUTE_DEFINITION_DATE:
		res = "ATTRIBUTE_DEFINITION_DATE"
	case ATTRIBUTE_DEFINITION_ENUMERATION:
		res = "ATTRIBUTE_DEFINITION_ENUMERATION"
	case ATTRIBUTE_DEFINITION_INTEGER:
		res = "ATTRIBUTE_DEFINITION_INTEGER"
	case ATTRIBUTE_DEFINITION_REAL:
		res = "ATTRIBUTE_DEFINITION_REAL"
	case ATTRIBUTE_DEFINITION_STRING:
		res = "ATTRIBUTE_DEFINITION_STRING"
	case ATTRIBUTE_DEFINITION_XHTML:
		res = "ATTRIBUTE_DEFINITION_XHTML"
	case ATTRIBUTE_VALUE_BOOLEAN:
		res = "ATTRIBUTE_VALUE_BOOLEAN"
	case ATTRIBUTE_VALUE_DATE:
		res = "ATTRIBUTE_VALUE_DATE"
	case ATTRIBUTE_VALUE_ENUMERATION:
		res = "ATTRIBUTE_VALUE_ENUMERATION"
	case ATTRIBUTE_VALUE_INTEGER:
		res = "ATTRIBUTE_VALUE_INTEGER"
	case ATTRIBUTE_VALUE_REAL:
		res = "ATTRIBUTE_VALUE_REAL"
	case ATTRIBUTE_VALUE_STRING:
		res = "ATTRIBUTE_VALUE_STRING"
	case ATTRIBUTE_VALUE_XHTML:
		res = "ATTRIBUTE_VALUE_XHTML"
	case DATATYPE_DEFINITION_BOOLEAN:
		res = "DATATYPE_DEFINITION_BOOLEAN"
	case DATATYPE_DEFINITION_DATE:
		res = "DATATYPE_DEFINITION_DATE"
	case DATATYPE_DEFINITION_ENUMERATION:
		res = "DATATYPE_DEFINITION_ENUMERATION"
	case DATATYPE_DEFINITION_INTEGER:
		res = "DATATYPE_DEFINITION_INTEGER"
	case DATATYPE_DEFINITION_REAL:
		res = "DATATYPE_DEFINITION_REAL"
	case DATATYPE_DEFINITION_STRING:
		res = "DATATYPE_DEFINITION_STRING"
	case DATATYPE_DEFINITION_XHTML:
		res = "DATATYPE_DEFINITION_XHTML"
	case EMBEDDED_VALUE:
		res = "EMBEDDED_VALUE"
	case ENUM_VALUE:
		res = "ENUM_VALUE"
	case RELATION_GROUP:
		res = "RELATION_GROUP"
	case RELATION_GROUP_TYPE:
		res = "RELATION_GROUP_TYPE"
	case REQ_IF:
		res = "REQ_IF"
	case REQ_IF_CONTENT:
		res = "REQ_IF_CONTENT"
	case REQ_IF_HEADER:
		res = "REQ_IF_HEADER"
	case REQ_IF_TOOL_EXTENSION:
		res = "REQ_IF_TOOL_EXTENSION"
	case SPECIFICATION:
		res = "SPECIFICATION"
	case SPECIFICATION_TYPE:
		res = "SPECIFICATION_TYPE"
	case SPEC_HIERARCHY:
		res = "SPEC_HIERARCHY"
	case SPEC_OBJECT:
		res = "SPEC_OBJECT"
	case SPEC_OBJECT_TYPE:
		res = "SPEC_OBJECT_TYPE"
	case SPEC_RELATION:
		res = "SPEC_RELATION"
	case SPEC_RELATION_TYPE:
		res = "SPEC_RELATION_TYPE"
	case XHTML_CONTENT:
		res = "XHTML_CONTENT"
	}
	return res
}

// GetPointerToGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetPointerToGongstructName[Type PointerToGongstruct]() (res string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case *ALTERNATIVE_ID:
		res = "ALTERNATIVE_ID"
	case *ATTRIBUTE_DEFINITION_BOOLEAN:
		res = "ATTRIBUTE_DEFINITION_BOOLEAN"
	case *ATTRIBUTE_DEFINITION_DATE:
		res = "ATTRIBUTE_DEFINITION_DATE"
	case *ATTRIBUTE_DEFINITION_ENUMERATION:
		res = "ATTRIBUTE_DEFINITION_ENUMERATION"
	case *ATTRIBUTE_DEFINITION_INTEGER:
		res = "ATTRIBUTE_DEFINITION_INTEGER"
	case *ATTRIBUTE_DEFINITION_REAL:
		res = "ATTRIBUTE_DEFINITION_REAL"
	case *ATTRIBUTE_DEFINITION_STRING:
		res = "ATTRIBUTE_DEFINITION_STRING"
	case *ATTRIBUTE_DEFINITION_XHTML:
		res = "ATTRIBUTE_DEFINITION_XHTML"
	case *ATTRIBUTE_VALUE_BOOLEAN:
		res = "ATTRIBUTE_VALUE_BOOLEAN"
	case *ATTRIBUTE_VALUE_DATE:
		res = "ATTRIBUTE_VALUE_DATE"
	case *ATTRIBUTE_VALUE_ENUMERATION:
		res = "ATTRIBUTE_VALUE_ENUMERATION"
	case *ATTRIBUTE_VALUE_INTEGER:
		res = "ATTRIBUTE_VALUE_INTEGER"
	case *ATTRIBUTE_VALUE_REAL:
		res = "ATTRIBUTE_VALUE_REAL"
	case *ATTRIBUTE_VALUE_STRING:
		res = "ATTRIBUTE_VALUE_STRING"
	case *ATTRIBUTE_VALUE_XHTML:
		res = "ATTRIBUTE_VALUE_XHTML"
	case *DATATYPE_DEFINITION_BOOLEAN:
		res = "DATATYPE_DEFINITION_BOOLEAN"
	case *DATATYPE_DEFINITION_DATE:
		res = "DATATYPE_DEFINITION_DATE"
	case *DATATYPE_DEFINITION_ENUMERATION:
		res = "DATATYPE_DEFINITION_ENUMERATION"
	case *DATATYPE_DEFINITION_INTEGER:
		res = "DATATYPE_DEFINITION_INTEGER"
	case *DATATYPE_DEFINITION_REAL:
		res = "DATATYPE_DEFINITION_REAL"
	case *DATATYPE_DEFINITION_STRING:
		res = "DATATYPE_DEFINITION_STRING"
	case *DATATYPE_DEFINITION_XHTML:
		res = "DATATYPE_DEFINITION_XHTML"
	case *EMBEDDED_VALUE:
		res = "EMBEDDED_VALUE"
	case *ENUM_VALUE:
		res = "ENUM_VALUE"
	case *RELATION_GROUP:
		res = "RELATION_GROUP"
	case *RELATION_GROUP_TYPE:
		res = "RELATION_GROUP_TYPE"
	case *REQ_IF:
		res = "REQ_IF"
	case *REQ_IF_CONTENT:
		res = "REQ_IF_CONTENT"
	case *REQ_IF_HEADER:
		res = "REQ_IF_HEADER"
	case *REQ_IF_TOOL_EXTENSION:
		res = "REQ_IF_TOOL_EXTENSION"
	case *SPECIFICATION:
		res = "SPECIFICATION"
	case *SPECIFICATION_TYPE:
		res = "SPECIFICATION_TYPE"
	case *SPEC_HIERARCHY:
		res = "SPEC_HIERARCHY"
	case *SPEC_OBJECT:
		res = "SPEC_OBJECT"
	case *SPEC_OBJECT_TYPE:
		res = "SPEC_OBJECT_TYPE"
	case *SPEC_RELATION:
		res = "SPEC_RELATION"
	case *SPEC_RELATION_TYPE:
		res = "SPEC_RELATION_TYPE"
	case *XHTML_CONTENT:
		res = "XHTML_CONTENT"
	}
	return res
}

// GetFields return the array of the fields
func GetFields[Type Gongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case ALTERNATIVE_ID:
		res = []string{"Name", "IDENTIFIER"}
	case ATTRIBUTE_DEFINITION_BOOLEAN:
		res = []string{"Name", "DESC", "IDENTIFIER", "IS_EDITABLE", "LAST_CHANGE", "LONG_NAME"}
	case ATTRIBUTE_DEFINITION_DATE:
		res = []string{"Name", "DESC", "IDENTIFIER", "IS_EDITABLE", "LAST_CHANGE", "LONG_NAME"}
	case ATTRIBUTE_DEFINITION_ENUMERATION:
		res = []string{"Name", "DESC", "IDENTIFIER", "IS_EDITABLE", "LAST_CHANGE", "LONG_NAME", "MULTI_VALUED"}
	case ATTRIBUTE_DEFINITION_INTEGER:
		res = []string{"Name", "DESC", "IDENTIFIER", "IS_EDITABLE", "LAST_CHANGE", "LONG_NAME"}
	case ATTRIBUTE_DEFINITION_REAL:
		res = []string{"Name", "DESC", "IDENTIFIER", "IS_EDITABLE", "LAST_CHANGE", "LONG_NAME"}
	case ATTRIBUTE_DEFINITION_STRING:
		res = []string{"Name", "DESC", "IDENTIFIER", "IS_EDITABLE", "LAST_CHANGE", "LONG_NAME"}
	case ATTRIBUTE_DEFINITION_XHTML:
		res = []string{"Name", "DESC", "IDENTIFIER", "IS_EDITABLE", "LAST_CHANGE", "LONG_NAME"}
	case ATTRIBUTE_VALUE_BOOLEAN:
		res = []string{"Name", "THE_VALUE"}
	case ATTRIBUTE_VALUE_DATE:
		res = []string{"Name", "THE_VALUE"}
	case ATTRIBUTE_VALUE_ENUMERATION:
		res = []string{"Name"}
	case ATTRIBUTE_VALUE_INTEGER:
		res = []string{"Name", "THE_VALUE"}
	case ATTRIBUTE_VALUE_REAL:
		res = []string{"Name", "THE_VALUE"}
	case ATTRIBUTE_VALUE_STRING:
		res = []string{"Name", "THE_VALUE"}
	case ATTRIBUTE_VALUE_XHTML:
		res = []string{"Name", "IS_SIMPLIFIED", "THE_VALUE", "THE_ORIGINAL_VALUE"}
	case DATATYPE_DEFINITION_BOOLEAN:
		res = []string{"Name", "DESC", "IDENTIFIER", "LAST_CHANGE", "LONG_NAME"}
	case DATATYPE_DEFINITION_DATE:
		res = []string{"Name", "DESC", "IDENTIFIER", "LAST_CHANGE", "LONG_NAME"}
	case DATATYPE_DEFINITION_ENUMERATION:
		res = []string{"Name", "DESC", "IDENTIFIER", "LAST_CHANGE", "LONG_NAME"}
	case DATATYPE_DEFINITION_INTEGER:
		res = []string{"Name", "DESC", "IDENTIFIER", "LAST_CHANGE", "LONG_NAME", "MAX", "MIN"}
	case DATATYPE_DEFINITION_REAL:
		res = []string{"Name", "ACCURACY", "DESC", "IDENTIFIER", "LAST_CHANGE", "LONG_NAME", "MAX", "MIN"}
	case DATATYPE_DEFINITION_STRING:
		res = []string{"Name", "DESC", "IDENTIFIER", "LAST_CHANGE", "LONG_NAME", "MAX_LENGTH"}
	case DATATYPE_DEFINITION_XHTML:
		res = []string{"Name", "DESC", "IDENTIFIER", "LAST_CHANGE", "LONG_NAME"}
	case EMBEDDED_VALUE:
		res = []string{"Name", "KEY", "OTHER_CONTENT"}
	case ENUM_VALUE:
		res = []string{"Name", "DESC", "IDENTIFIER", "LAST_CHANGE", "LONG_NAME"}
	case RELATION_GROUP:
		res = []string{"Name", "DESC", "IDENTIFIER", "LAST_CHANGE", "LONG_NAME"}
	case RELATION_GROUP_TYPE:
		res = []string{"Name", "DESC", "IDENTIFIER", "LAST_CHANGE", "LONG_NAME"}
	case REQ_IF:
		res = []string{"Name", "Lang"}
	case REQ_IF_CONTENT:
		res = []string{"Name"}
	case REQ_IF_HEADER:
		res = []string{"Name", "IDENTIFIER", "COMMENT", "CREATION_TIME", "REPOSITORY_ID", "REQ_IF_TOOL_ID", "REQ_IF_VERSION", "SOURCE_TOOL_ID", "TITLE"}
	case REQ_IF_TOOL_EXTENSION:
		res = []string{"Name"}
	case SPECIFICATION:
		res = []string{"Name", "DESC", "IDENTIFIER", "LAST_CHANGE", "LONG_NAME"}
	case SPECIFICATION_TYPE:
		res = []string{"Name", "DESC", "IDENTIFIER", "LAST_CHANGE", "LONG_NAME"}
	case SPEC_HIERARCHY:
		res = []string{"Name", "DESC", "IDENTIFIER", "IS_EDITABLE", "IS_TABLE_INTERNAL", "LAST_CHANGE", "LONG_NAME"}
	case SPEC_OBJECT:
		res = []string{"Name", "DESC", "IDENTIFIER", "LAST_CHANGE", "LONG_NAME"}
	case SPEC_OBJECT_TYPE:
		res = []string{"Name", "DESC", "IDENTIFIER", "LAST_CHANGE", "LONG_NAME"}
	case SPEC_RELATION:
		res = []string{"Name", "DESC", "IDENTIFIER", "LAST_CHANGE", "LONG_NAME"}
	case SPEC_RELATION_TYPE:
		res = []string{"Name", "DESC", "IDENTIFIER", "LAST_CHANGE", "LONG_NAME"}
	case XHTML_CONTENT:
		res = []string{"Name"}
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
	case ALTERNATIVE_ID:
		var rf ReverseField
		_ = rf
	case ATTRIBUTE_DEFINITION_BOOLEAN:
		var rf ReverseField
		_ = rf
	case ATTRIBUTE_DEFINITION_DATE:
		var rf ReverseField
		_ = rf
	case ATTRIBUTE_DEFINITION_ENUMERATION:
		var rf ReverseField
		_ = rf
	case ATTRIBUTE_DEFINITION_INTEGER:
		var rf ReverseField
		_ = rf
	case ATTRIBUTE_DEFINITION_REAL:
		var rf ReverseField
		_ = rf
	case ATTRIBUTE_DEFINITION_STRING:
		var rf ReverseField
		_ = rf
	case ATTRIBUTE_DEFINITION_XHTML:
		var rf ReverseField
		_ = rf
	case ATTRIBUTE_VALUE_BOOLEAN:
		var rf ReverseField
		_ = rf
	case ATTRIBUTE_VALUE_DATE:
		var rf ReverseField
		_ = rf
	case ATTRIBUTE_VALUE_ENUMERATION:
		var rf ReverseField
		_ = rf
	case ATTRIBUTE_VALUE_INTEGER:
		var rf ReverseField
		_ = rf
	case ATTRIBUTE_VALUE_REAL:
		var rf ReverseField
		_ = rf
	case ATTRIBUTE_VALUE_STRING:
		var rf ReverseField
		_ = rf
	case ATTRIBUTE_VALUE_XHTML:
		var rf ReverseField
		_ = rf
	case DATATYPE_DEFINITION_BOOLEAN:
		var rf ReverseField
		_ = rf
	case DATATYPE_DEFINITION_DATE:
		var rf ReverseField
		_ = rf
	case DATATYPE_DEFINITION_ENUMERATION:
		var rf ReverseField
		_ = rf
	case DATATYPE_DEFINITION_INTEGER:
		var rf ReverseField
		_ = rf
	case DATATYPE_DEFINITION_REAL:
		var rf ReverseField
		_ = rf
	case DATATYPE_DEFINITION_STRING:
		var rf ReverseField
		_ = rf
	case DATATYPE_DEFINITION_XHTML:
		var rf ReverseField
		_ = rf
	case EMBEDDED_VALUE:
		var rf ReverseField
		_ = rf
	case ENUM_VALUE:
		var rf ReverseField
		_ = rf
	case RELATION_GROUP:
		var rf ReverseField
		_ = rf
	case RELATION_GROUP_TYPE:
		var rf ReverseField
		_ = rf
	case REQ_IF:
		var rf ReverseField
		_ = rf
	case REQ_IF_CONTENT:
		var rf ReverseField
		_ = rf
	case REQ_IF_HEADER:
		var rf ReverseField
		_ = rf
	case REQ_IF_TOOL_EXTENSION:
		var rf ReverseField
		_ = rf
	case SPECIFICATION:
		var rf ReverseField
		_ = rf
	case SPECIFICATION_TYPE:
		var rf ReverseField
		_ = rf
	case SPEC_HIERARCHY:
		var rf ReverseField
		_ = rf
	case SPEC_OBJECT:
		var rf ReverseField
		_ = rf
	case SPEC_OBJECT_TYPE:
		var rf ReverseField
		_ = rf
	case SPEC_RELATION:
		var rf ReverseField
		_ = rf
	case SPEC_RELATION_TYPE:
		var rf ReverseField
		_ = rf
	case XHTML_CONTENT:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "ATTRIBUTE_VALUE_XHTML"
		rf.Fieldname = "THE_VALUE"
		res = append(res, rf)
		rf.GongstructName = "ATTRIBUTE_VALUE_XHTML"
		rf.Fieldname = "THE_ORIGINAL_VALUE"
		res = append(res, rf)
	}
	return
}

// GetFieldsFromPointer return the array of the fields
func GetFieldsFromPointer[Type PointerToGongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case *ALTERNATIVE_ID:
		res = []string{"Name", "IDENTIFIER"}
	case *ATTRIBUTE_DEFINITION_BOOLEAN:
		res = []string{"Name", "DESC", "IDENTIFIER", "IS_EDITABLE", "LAST_CHANGE", "LONG_NAME"}
	case *ATTRIBUTE_DEFINITION_DATE:
		res = []string{"Name", "DESC", "IDENTIFIER", "IS_EDITABLE", "LAST_CHANGE", "LONG_NAME"}
	case *ATTRIBUTE_DEFINITION_ENUMERATION:
		res = []string{"Name", "DESC", "IDENTIFIER", "IS_EDITABLE", "LAST_CHANGE", "LONG_NAME", "MULTI_VALUED"}
	case *ATTRIBUTE_DEFINITION_INTEGER:
		res = []string{"Name", "DESC", "IDENTIFIER", "IS_EDITABLE", "LAST_CHANGE", "LONG_NAME"}
	case *ATTRIBUTE_DEFINITION_REAL:
		res = []string{"Name", "DESC", "IDENTIFIER", "IS_EDITABLE", "LAST_CHANGE", "LONG_NAME"}
	case *ATTRIBUTE_DEFINITION_STRING:
		res = []string{"Name", "DESC", "IDENTIFIER", "IS_EDITABLE", "LAST_CHANGE", "LONG_NAME"}
	case *ATTRIBUTE_DEFINITION_XHTML:
		res = []string{"Name", "DESC", "IDENTIFIER", "IS_EDITABLE", "LAST_CHANGE", "LONG_NAME"}
	case *ATTRIBUTE_VALUE_BOOLEAN:
		res = []string{"Name", "THE_VALUE"}
	case *ATTRIBUTE_VALUE_DATE:
		res = []string{"Name", "THE_VALUE"}
	case *ATTRIBUTE_VALUE_ENUMERATION:
		res = []string{"Name"}
	case *ATTRIBUTE_VALUE_INTEGER:
		res = []string{"Name", "THE_VALUE"}
	case *ATTRIBUTE_VALUE_REAL:
		res = []string{"Name", "THE_VALUE"}
	case *ATTRIBUTE_VALUE_STRING:
		res = []string{"Name", "THE_VALUE"}
	case *ATTRIBUTE_VALUE_XHTML:
		res = []string{"Name", "IS_SIMPLIFIED", "THE_VALUE", "THE_ORIGINAL_VALUE"}
	case *DATATYPE_DEFINITION_BOOLEAN:
		res = []string{"Name", "DESC", "IDENTIFIER", "LAST_CHANGE", "LONG_NAME"}
	case *DATATYPE_DEFINITION_DATE:
		res = []string{"Name", "DESC", "IDENTIFIER", "LAST_CHANGE", "LONG_NAME"}
	case *DATATYPE_DEFINITION_ENUMERATION:
		res = []string{"Name", "DESC", "IDENTIFIER", "LAST_CHANGE", "LONG_NAME"}
	case *DATATYPE_DEFINITION_INTEGER:
		res = []string{"Name", "DESC", "IDENTIFIER", "LAST_CHANGE", "LONG_NAME", "MAX", "MIN"}
	case *DATATYPE_DEFINITION_REAL:
		res = []string{"Name", "ACCURACY", "DESC", "IDENTIFIER", "LAST_CHANGE", "LONG_NAME", "MAX", "MIN"}
	case *DATATYPE_DEFINITION_STRING:
		res = []string{"Name", "DESC", "IDENTIFIER", "LAST_CHANGE", "LONG_NAME", "MAX_LENGTH"}
	case *DATATYPE_DEFINITION_XHTML:
		res = []string{"Name", "DESC", "IDENTIFIER", "LAST_CHANGE", "LONG_NAME"}
	case *EMBEDDED_VALUE:
		res = []string{"Name", "KEY", "OTHER_CONTENT"}
	case *ENUM_VALUE:
		res = []string{"Name", "DESC", "IDENTIFIER", "LAST_CHANGE", "LONG_NAME"}
	case *RELATION_GROUP:
		res = []string{"Name", "DESC", "IDENTIFIER", "LAST_CHANGE", "LONG_NAME"}
	case *RELATION_GROUP_TYPE:
		res = []string{"Name", "DESC", "IDENTIFIER", "LAST_CHANGE", "LONG_NAME"}
	case *REQ_IF:
		res = []string{"Name", "Lang"}
	case *REQ_IF_CONTENT:
		res = []string{"Name"}
	case *REQ_IF_HEADER:
		res = []string{"Name", "IDENTIFIER", "COMMENT", "CREATION_TIME", "REPOSITORY_ID", "REQ_IF_TOOL_ID", "REQ_IF_VERSION", "SOURCE_TOOL_ID", "TITLE"}
	case *REQ_IF_TOOL_EXTENSION:
		res = []string{"Name"}
	case *SPECIFICATION:
		res = []string{"Name", "DESC", "IDENTIFIER", "LAST_CHANGE", "LONG_NAME"}
	case *SPECIFICATION_TYPE:
		res = []string{"Name", "DESC", "IDENTIFIER", "LAST_CHANGE", "LONG_NAME"}
	case *SPEC_HIERARCHY:
		res = []string{"Name", "DESC", "IDENTIFIER", "IS_EDITABLE", "IS_TABLE_INTERNAL", "LAST_CHANGE", "LONG_NAME"}
	case *SPEC_OBJECT:
		res = []string{"Name", "DESC", "IDENTIFIER", "LAST_CHANGE", "LONG_NAME"}
	case *SPEC_OBJECT_TYPE:
		res = []string{"Name", "DESC", "IDENTIFIER", "LAST_CHANGE", "LONG_NAME"}
	case *SPEC_RELATION:
		res = []string{"Name", "DESC", "IDENTIFIER", "LAST_CHANGE", "LONG_NAME"}
	case *SPEC_RELATION_TYPE:
		res = []string{"Name", "DESC", "IDENTIFIER", "LAST_CHANGE", "LONG_NAME"}
	case *XHTML_CONTENT:
		res = []string{"Name"}
	}
	return
}

func GetFieldStringValueFromPointer[Type PointerToGongstruct](instance Type, fieldName string) (res string) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
	case *ALTERNATIVE_ID:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "IDENTIFIER":
			res = inferedInstance.IDENTIFIER
		}
	case *ATTRIBUTE_DEFINITION_BOOLEAN:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESC":
			res = inferedInstance.DESC
		case "IDENTIFIER":
			res = inferedInstance.IDENTIFIER
		case "IS_EDITABLE":
			res = fmt.Sprintf("%t", inferedInstance.IS_EDITABLE)
		case "LAST_CHANGE":
			res = inferedInstance.LAST_CHANGE
		case "LONG_NAME":
			res = inferedInstance.LONG_NAME
		}
	case *ATTRIBUTE_DEFINITION_DATE:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESC":
			res = inferedInstance.DESC
		case "IDENTIFIER":
			res = inferedInstance.IDENTIFIER
		case "IS_EDITABLE":
			res = fmt.Sprintf("%t", inferedInstance.IS_EDITABLE)
		case "LAST_CHANGE":
			res = inferedInstance.LAST_CHANGE
		case "LONG_NAME":
			res = inferedInstance.LONG_NAME
		}
	case *ATTRIBUTE_DEFINITION_ENUMERATION:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESC":
			res = inferedInstance.DESC
		case "IDENTIFIER":
			res = inferedInstance.IDENTIFIER
		case "IS_EDITABLE":
			res = fmt.Sprintf("%t", inferedInstance.IS_EDITABLE)
		case "LAST_CHANGE":
			res = inferedInstance.LAST_CHANGE
		case "LONG_NAME":
			res = inferedInstance.LONG_NAME
		case "MULTI_VALUED":
			res = fmt.Sprintf("%t", inferedInstance.MULTI_VALUED)
		}
	case *ATTRIBUTE_DEFINITION_INTEGER:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESC":
			res = inferedInstance.DESC
		case "IDENTIFIER":
			res = inferedInstance.IDENTIFIER
		case "IS_EDITABLE":
			res = fmt.Sprintf("%t", inferedInstance.IS_EDITABLE)
		case "LAST_CHANGE":
			res = inferedInstance.LAST_CHANGE
		case "LONG_NAME":
			res = inferedInstance.LONG_NAME
		}
	case *ATTRIBUTE_DEFINITION_REAL:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESC":
			res = inferedInstance.DESC
		case "IDENTIFIER":
			res = inferedInstance.IDENTIFIER
		case "IS_EDITABLE":
			res = fmt.Sprintf("%t", inferedInstance.IS_EDITABLE)
		case "LAST_CHANGE":
			res = inferedInstance.LAST_CHANGE
		case "LONG_NAME":
			res = inferedInstance.LONG_NAME
		}
	case *ATTRIBUTE_DEFINITION_STRING:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESC":
			res = inferedInstance.DESC
		case "IDENTIFIER":
			res = inferedInstance.IDENTIFIER
		case "IS_EDITABLE":
			res = fmt.Sprintf("%t", inferedInstance.IS_EDITABLE)
		case "LAST_CHANGE":
			res = inferedInstance.LAST_CHANGE
		case "LONG_NAME":
			res = inferedInstance.LONG_NAME
		}
	case *ATTRIBUTE_DEFINITION_XHTML:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESC":
			res = inferedInstance.DESC
		case "IDENTIFIER":
			res = inferedInstance.IDENTIFIER
		case "IS_EDITABLE":
			res = fmt.Sprintf("%t", inferedInstance.IS_EDITABLE)
		case "LAST_CHANGE":
			res = inferedInstance.LAST_CHANGE
		case "LONG_NAME":
			res = inferedInstance.LONG_NAME
		}
	case *ATTRIBUTE_VALUE_BOOLEAN:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "THE_VALUE":
			res = fmt.Sprintf("%t", inferedInstance.THE_VALUE)
		}
	case *ATTRIBUTE_VALUE_DATE:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "THE_VALUE":
			res = inferedInstance.THE_VALUE
		}
	case *ATTRIBUTE_VALUE_ENUMERATION:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		}
	case *ATTRIBUTE_VALUE_INTEGER:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "THE_VALUE":
			res = fmt.Sprintf("%d", inferedInstance.THE_VALUE)
		}
	case *ATTRIBUTE_VALUE_REAL:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "THE_VALUE":
			res = fmt.Sprintf("%f", inferedInstance.THE_VALUE)
		}
	case *ATTRIBUTE_VALUE_STRING:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "THE_VALUE":
			res = inferedInstance.THE_VALUE
		}
	case *ATTRIBUTE_VALUE_XHTML:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "IS_SIMPLIFIED":
			res = fmt.Sprintf("%t", inferedInstance.IS_SIMPLIFIED)
		case "THE_VALUE":
			for idx, __instance__ := range inferedInstance.THE_VALUE {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "THE_ORIGINAL_VALUE":
			for idx, __instance__ := range inferedInstance.THE_ORIGINAL_VALUE {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case *DATATYPE_DEFINITION_BOOLEAN:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESC":
			res = inferedInstance.DESC
		case "IDENTIFIER":
			res = inferedInstance.IDENTIFIER
		case "LAST_CHANGE":
			res = inferedInstance.LAST_CHANGE
		case "LONG_NAME":
			res = inferedInstance.LONG_NAME
		}
	case *DATATYPE_DEFINITION_DATE:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESC":
			res = inferedInstance.DESC
		case "IDENTIFIER":
			res = inferedInstance.IDENTIFIER
		case "LAST_CHANGE":
			res = inferedInstance.LAST_CHANGE
		case "LONG_NAME":
			res = inferedInstance.LONG_NAME
		}
	case *DATATYPE_DEFINITION_ENUMERATION:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESC":
			res = inferedInstance.DESC
		case "IDENTIFIER":
			res = inferedInstance.IDENTIFIER
		case "LAST_CHANGE":
			res = inferedInstance.LAST_CHANGE
		case "LONG_NAME":
			res = inferedInstance.LONG_NAME
		}
	case *DATATYPE_DEFINITION_INTEGER:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESC":
			res = inferedInstance.DESC
		case "IDENTIFIER":
			res = inferedInstance.IDENTIFIER
		case "LAST_CHANGE":
			res = inferedInstance.LAST_CHANGE
		case "LONG_NAME":
			res = inferedInstance.LONG_NAME
		case "MAX":
			res = fmt.Sprintf("%d", inferedInstance.MAX)
		case "MIN":
			res = fmt.Sprintf("%d", inferedInstance.MIN)
		}
	case *DATATYPE_DEFINITION_REAL:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "ACCURACY":
			res = fmt.Sprintf("%d", inferedInstance.ACCURACY)
		case "DESC":
			res = inferedInstance.DESC
		case "IDENTIFIER":
			res = inferedInstance.IDENTIFIER
		case "LAST_CHANGE":
			res = inferedInstance.LAST_CHANGE
		case "LONG_NAME":
			res = inferedInstance.LONG_NAME
		case "MAX":
			res = fmt.Sprintf("%f", inferedInstance.MAX)
		case "MIN":
			res = fmt.Sprintf("%f", inferedInstance.MIN)
		}
	case *DATATYPE_DEFINITION_STRING:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESC":
			res = inferedInstance.DESC
		case "IDENTIFIER":
			res = inferedInstance.IDENTIFIER
		case "LAST_CHANGE":
			res = inferedInstance.LAST_CHANGE
		case "LONG_NAME":
			res = inferedInstance.LONG_NAME
		case "MAX_LENGTH":
			res = fmt.Sprintf("%d", inferedInstance.MAX_LENGTH)
		}
	case *DATATYPE_DEFINITION_XHTML:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESC":
			res = inferedInstance.DESC
		case "IDENTIFIER":
			res = inferedInstance.IDENTIFIER
		case "LAST_CHANGE":
			res = inferedInstance.LAST_CHANGE
		case "LONG_NAME":
			res = inferedInstance.LONG_NAME
		}
	case *EMBEDDED_VALUE:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "KEY":
			res = fmt.Sprintf("%d", inferedInstance.KEY)
		case "OTHER_CONTENT":
			res = inferedInstance.OTHER_CONTENT
		}
	case *ENUM_VALUE:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESC":
			res = inferedInstance.DESC
		case "IDENTIFIER":
			res = inferedInstance.IDENTIFIER
		case "LAST_CHANGE":
			res = inferedInstance.LAST_CHANGE
		case "LONG_NAME":
			res = inferedInstance.LONG_NAME
		}
	case *RELATION_GROUP:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESC":
			res = inferedInstance.DESC
		case "IDENTIFIER":
			res = inferedInstance.IDENTIFIER
		case "LAST_CHANGE":
			res = inferedInstance.LAST_CHANGE
		case "LONG_NAME":
			res = inferedInstance.LONG_NAME
		}
	case *RELATION_GROUP_TYPE:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESC":
			res = inferedInstance.DESC
		case "IDENTIFIER":
			res = inferedInstance.IDENTIFIER
		case "LAST_CHANGE":
			res = inferedInstance.LAST_CHANGE
		case "LONG_NAME":
			res = inferedInstance.LONG_NAME
		}
	case *REQ_IF:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Lang":
			res = inferedInstance.Lang
		}
	case *REQ_IF_CONTENT:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		}
	case *REQ_IF_HEADER:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "IDENTIFIER":
			res = inferedInstance.IDENTIFIER
		case "COMMENT":
			res = inferedInstance.COMMENT
		case "CREATION_TIME":
			res = inferedInstance.CREATION_TIME
		case "REPOSITORY_ID":
			res = inferedInstance.REPOSITORY_ID
		case "REQ_IF_TOOL_ID":
			res = inferedInstance.REQ_IF_TOOL_ID
		case "REQ_IF_VERSION":
			res = inferedInstance.REQ_IF_VERSION
		case "SOURCE_TOOL_ID":
			res = inferedInstance.SOURCE_TOOL_ID
		case "TITLE":
			res = inferedInstance.TITLE
		}
	case *REQ_IF_TOOL_EXTENSION:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		}
	case *SPECIFICATION:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESC":
			res = inferedInstance.DESC
		case "IDENTIFIER":
			res = inferedInstance.IDENTIFIER
		case "LAST_CHANGE":
			res = inferedInstance.LAST_CHANGE
		case "LONG_NAME":
			res = inferedInstance.LONG_NAME
		}
	case *SPECIFICATION_TYPE:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESC":
			res = inferedInstance.DESC
		case "IDENTIFIER":
			res = inferedInstance.IDENTIFIER
		case "LAST_CHANGE":
			res = inferedInstance.LAST_CHANGE
		case "LONG_NAME":
			res = inferedInstance.LONG_NAME
		}
	case *SPEC_HIERARCHY:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESC":
			res = inferedInstance.DESC
		case "IDENTIFIER":
			res = inferedInstance.IDENTIFIER
		case "IS_EDITABLE":
			res = fmt.Sprintf("%t", inferedInstance.IS_EDITABLE)
		case "IS_TABLE_INTERNAL":
			res = fmt.Sprintf("%t", inferedInstance.IS_TABLE_INTERNAL)
		case "LAST_CHANGE":
			res = inferedInstance.LAST_CHANGE
		case "LONG_NAME":
			res = inferedInstance.LONG_NAME
		}
	case *SPEC_OBJECT:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESC":
			res = inferedInstance.DESC
		case "IDENTIFIER":
			res = inferedInstance.IDENTIFIER
		case "LAST_CHANGE":
			res = inferedInstance.LAST_CHANGE
		case "LONG_NAME":
			res = inferedInstance.LONG_NAME
		}
	case *SPEC_OBJECT_TYPE:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESC":
			res = inferedInstance.DESC
		case "IDENTIFIER":
			res = inferedInstance.IDENTIFIER
		case "LAST_CHANGE":
			res = inferedInstance.LAST_CHANGE
		case "LONG_NAME":
			res = inferedInstance.LONG_NAME
		}
	case *SPEC_RELATION:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESC":
			res = inferedInstance.DESC
		case "IDENTIFIER":
			res = inferedInstance.IDENTIFIER
		case "LAST_CHANGE":
			res = inferedInstance.LAST_CHANGE
		case "LONG_NAME":
			res = inferedInstance.LONG_NAME
		}
	case *SPEC_RELATION_TYPE:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESC":
			res = inferedInstance.DESC
		case "IDENTIFIER":
			res = inferedInstance.IDENTIFIER
		case "LAST_CHANGE":
			res = inferedInstance.LAST_CHANGE
		case "LONG_NAME":
			res = inferedInstance.LONG_NAME
		}
	case *XHTML_CONTENT:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		}
	default:
		_ = inferedInstance
	}
	return
}

func GetFieldStringValue[Type Gongstruct](instance Type, fieldName string) (res string) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
	case ALTERNATIVE_ID:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "IDENTIFIER":
			res = inferedInstance.IDENTIFIER
		}
	case ATTRIBUTE_DEFINITION_BOOLEAN:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESC":
			res = inferedInstance.DESC
		case "IDENTIFIER":
			res = inferedInstance.IDENTIFIER
		case "IS_EDITABLE":
			res = fmt.Sprintf("%t", inferedInstance.IS_EDITABLE)
		case "LAST_CHANGE":
			res = inferedInstance.LAST_CHANGE
		case "LONG_NAME":
			res = inferedInstance.LONG_NAME
		}
	case ATTRIBUTE_DEFINITION_DATE:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESC":
			res = inferedInstance.DESC
		case "IDENTIFIER":
			res = inferedInstance.IDENTIFIER
		case "IS_EDITABLE":
			res = fmt.Sprintf("%t", inferedInstance.IS_EDITABLE)
		case "LAST_CHANGE":
			res = inferedInstance.LAST_CHANGE
		case "LONG_NAME":
			res = inferedInstance.LONG_NAME
		}
	case ATTRIBUTE_DEFINITION_ENUMERATION:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESC":
			res = inferedInstance.DESC
		case "IDENTIFIER":
			res = inferedInstance.IDENTIFIER
		case "IS_EDITABLE":
			res = fmt.Sprintf("%t", inferedInstance.IS_EDITABLE)
		case "LAST_CHANGE":
			res = inferedInstance.LAST_CHANGE
		case "LONG_NAME":
			res = inferedInstance.LONG_NAME
		case "MULTI_VALUED":
			res = fmt.Sprintf("%t", inferedInstance.MULTI_VALUED)
		}
	case ATTRIBUTE_DEFINITION_INTEGER:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESC":
			res = inferedInstance.DESC
		case "IDENTIFIER":
			res = inferedInstance.IDENTIFIER
		case "IS_EDITABLE":
			res = fmt.Sprintf("%t", inferedInstance.IS_EDITABLE)
		case "LAST_CHANGE":
			res = inferedInstance.LAST_CHANGE
		case "LONG_NAME":
			res = inferedInstance.LONG_NAME
		}
	case ATTRIBUTE_DEFINITION_REAL:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESC":
			res = inferedInstance.DESC
		case "IDENTIFIER":
			res = inferedInstance.IDENTIFIER
		case "IS_EDITABLE":
			res = fmt.Sprintf("%t", inferedInstance.IS_EDITABLE)
		case "LAST_CHANGE":
			res = inferedInstance.LAST_CHANGE
		case "LONG_NAME":
			res = inferedInstance.LONG_NAME
		}
	case ATTRIBUTE_DEFINITION_STRING:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESC":
			res = inferedInstance.DESC
		case "IDENTIFIER":
			res = inferedInstance.IDENTIFIER
		case "IS_EDITABLE":
			res = fmt.Sprintf("%t", inferedInstance.IS_EDITABLE)
		case "LAST_CHANGE":
			res = inferedInstance.LAST_CHANGE
		case "LONG_NAME":
			res = inferedInstance.LONG_NAME
		}
	case ATTRIBUTE_DEFINITION_XHTML:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESC":
			res = inferedInstance.DESC
		case "IDENTIFIER":
			res = inferedInstance.IDENTIFIER
		case "IS_EDITABLE":
			res = fmt.Sprintf("%t", inferedInstance.IS_EDITABLE)
		case "LAST_CHANGE":
			res = inferedInstance.LAST_CHANGE
		case "LONG_NAME":
			res = inferedInstance.LONG_NAME
		}
	case ATTRIBUTE_VALUE_BOOLEAN:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "THE_VALUE":
			res = fmt.Sprintf("%t", inferedInstance.THE_VALUE)
		}
	case ATTRIBUTE_VALUE_DATE:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "THE_VALUE":
			res = inferedInstance.THE_VALUE
		}
	case ATTRIBUTE_VALUE_ENUMERATION:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		}
	case ATTRIBUTE_VALUE_INTEGER:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "THE_VALUE":
			res = fmt.Sprintf("%d", inferedInstance.THE_VALUE)
		}
	case ATTRIBUTE_VALUE_REAL:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "THE_VALUE":
			res = fmt.Sprintf("%f", inferedInstance.THE_VALUE)
		}
	case ATTRIBUTE_VALUE_STRING:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "THE_VALUE":
			res = inferedInstance.THE_VALUE
		}
	case ATTRIBUTE_VALUE_XHTML:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "IS_SIMPLIFIED":
			res = fmt.Sprintf("%t", inferedInstance.IS_SIMPLIFIED)
		case "THE_VALUE":
			for idx, __instance__ := range inferedInstance.THE_VALUE {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "THE_ORIGINAL_VALUE":
			for idx, __instance__ := range inferedInstance.THE_ORIGINAL_VALUE {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case DATATYPE_DEFINITION_BOOLEAN:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESC":
			res = inferedInstance.DESC
		case "IDENTIFIER":
			res = inferedInstance.IDENTIFIER
		case "LAST_CHANGE":
			res = inferedInstance.LAST_CHANGE
		case "LONG_NAME":
			res = inferedInstance.LONG_NAME
		}
	case DATATYPE_DEFINITION_DATE:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESC":
			res = inferedInstance.DESC
		case "IDENTIFIER":
			res = inferedInstance.IDENTIFIER
		case "LAST_CHANGE":
			res = inferedInstance.LAST_CHANGE
		case "LONG_NAME":
			res = inferedInstance.LONG_NAME
		}
	case DATATYPE_DEFINITION_ENUMERATION:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESC":
			res = inferedInstance.DESC
		case "IDENTIFIER":
			res = inferedInstance.IDENTIFIER
		case "LAST_CHANGE":
			res = inferedInstance.LAST_CHANGE
		case "LONG_NAME":
			res = inferedInstance.LONG_NAME
		}
	case DATATYPE_DEFINITION_INTEGER:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESC":
			res = inferedInstance.DESC
		case "IDENTIFIER":
			res = inferedInstance.IDENTIFIER
		case "LAST_CHANGE":
			res = inferedInstance.LAST_CHANGE
		case "LONG_NAME":
			res = inferedInstance.LONG_NAME
		case "MAX":
			res = fmt.Sprintf("%d", inferedInstance.MAX)
		case "MIN":
			res = fmt.Sprintf("%d", inferedInstance.MIN)
		}
	case DATATYPE_DEFINITION_REAL:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "ACCURACY":
			res = fmt.Sprintf("%d", inferedInstance.ACCURACY)
		case "DESC":
			res = inferedInstance.DESC
		case "IDENTIFIER":
			res = inferedInstance.IDENTIFIER
		case "LAST_CHANGE":
			res = inferedInstance.LAST_CHANGE
		case "LONG_NAME":
			res = inferedInstance.LONG_NAME
		case "MAX":
			res = fmt.Sprintf("%f", inferedInstance.MAX)
		case "MIN":
			res = fmt.Sprintf("%f", inferedInstance.MIN)
		}
	case DATATYPE_DEFINITION_STRING:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESC":
			res = inferedInstance.DESC
		case "IDENTIFIER":
			res = inferedInstance.IDENTIFIER
		case "LAST_CHANGE":
			res = inferedInstance.LAST_CHANGE
		case "LONG_NAME":
			res = inferedInstance.LONG_NAME
		case "MAX_LENGTH":
			res = fmt.Sprintf("%d", inferedInstance.MAX_LENGTH)
		}
	case DATATYPE_DEFINITION_XHTML:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESC":
			res = inferedInstance.DESC
		case "IDENTIFIER":
			res = inferedInstance.IDENTIFIER
		case "LAST_CHANGE":
			res = inferedInstance.LAST_CHANGE
		case "LONG_NAME":
			res = inferedInstance.LONG_NAME
		}
	case EMBEDDED_VALUE:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "KEY":
			res = fmt.Sprintf("%d", inferedInstance.KEY)
		case "OTHER_CONTENT":
			res = inferedInstance.OTHER_CONTENT
		}
	case ENUM_VALUE:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESC":
			res = inferedInstance.DESC
		case "IDENTIFIER":
			res = inferedInstance.IDENTIFIER
		case "LAST_CHANGE":
			res = inferedInstance.LAST_CHANGE
		case "LONG_NAME":
			res = inferedInstance.LONG_NAME
		}
	case RELATION_GROUP:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESC":
			res = inferedInstance.DESC
		case "IDENTIFIER":
			res = inferedInstance.IDENTIFIER
		case "LAST_CHANGE":
			res = inferedInstance.LAST_CHANGE
		case "LONG_NAME":
			res = inferedInstance.LONG_NAME
		}
	case RELATION_GROUP_TYPE:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESC":
			res = inferedInstance.DESC
		case "IDENTIFIER":
			res = inferedInstance.IDENTIFIER
		case "LAST_CHANGE":
			res = inferedInstance.LAST_CHANGE
		case "LONG_NAME":
			res = inferedInstance.LONG_NAME
		}
	case REQ_IF:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Lang":
			res = inferedInstance.Lang
		}
	case REQ_IF_CONTENT:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		}
	case REQ_IF_HEADER:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "IDENTIFIER":
			res = inferedInstance.IDENTIFIER
		case "COMMENT":
			res = inferedInstance.COMMENT
		case "CREATION_TIME":
			res = inferedInstance.CREATION_TIME
		case "REPOSITORY_ID":
			res = inferedInstance.REPOSITORY_ID
		case "REQ_IF_TOOL_ID":
			res = inferedInstance.REQ_IF_TOOL_ID
		case "REQ_IF_VERSION":
			res = inferedInstance.REQ_IF_VERSION
		case "SOURCE_TOOL_ID":
			res = inferedInstance.SOURCE_TOOL_ID
		case "TITLE":
			res = inferedInstance.TITLE
		}
	case REQ_IF_TOOL_EXTENSION:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		}
	case SPECIFICATION:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESC":
			res = inferedInstance.DESC
		case "IDENTIFIER":
			res = inferedInstance.IDENTIFIER
		case "LAST_CHANGE":
			res = inferedInstance.LAST_CHANGE
		case "LONG_NAME":
			res = inferedInstance.LONG_NAME
		}
	case SPECIFICATION_TYPE:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESC":
			res = inferedInstance.DESC
		case "IDENTIFIER":
			res = inferedInstance.IDENTIFIER
		case "LAST_CHANGE":
			res = inferedInstance.LAST_CHANGE
		case "LONG_NAME":
			res = inferedInstance.LONG_NAME
		}
	case SPEC_HIERARCHY:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESC":
			res = inferedInstance.DESC
		case "IDENTIFIER":
			res = inferedInstance.IDENTIFIER
		case "IS_EDITABLE":
			res = fmt.Sprintf("%t", inferedInstance.IS_EDITABLE)
		case "IS_TABLE_INTERNAL":
			res = fmt.Sprintf("%t", inferedInstance.IS_TABLE_INTERNAL)
		case "LAST_CHANGE":
			res = inferedInstance.LAST_CHANGE
		case "LONG_NAME":
			res = inferedInstance.LONG_NAME
		}
	case SPEC_OBJECT:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESC":
			res = inferedInstance.DESC
		case "IDENTIFIER":
			res = inferedInstance.IDENTIFIER
		case "LAST_CHANGE":
			res = inferedInstance.LAST_CHANGE
		case "LONG_NAME":
			res = inferedInstance.LONG_NAME
		}
	case SPEC_OBJECT_TYPE:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESC":
			res = inferedInstance.DESC
		case "IDENTIFIER":
			res = inferedInstance.IDENTIFIER
		case "LAST_CHANGE":
			res = inferedInstance.LAST_CHANGE
		case "LONG_NAME":
			res = inferedInstance.LONG_NAME
		}
	case SPEC_RELATION:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESC":
			res = inferedInstance.DESC
		case "IDENTIFIER":
			res = inferedInstance.IDENTIFIER
		case "LAST_CHANGE":
			res = inferedInstance.LAST_CHANGE
		case "LONG_NAME":
			res = inferedInstance.LONG_NAME
		}
	case SPEC_RELATION_TYPE:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DESC":
			res = inferedInstance.DESC
		case "IDENTIFIER":
			res = inferedInstance.IDENTIFIER
		case "LAST_CHANGE":
			res = inferedInstance.LAST_CHANGE
		case "LONG_NAME":
			res = inferedInstance.LONG_NAME
		}
	case XHTML_CONTENT:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		}
	default:
		_ = inferedInstance
	}
	return
}

// Last line of the template
