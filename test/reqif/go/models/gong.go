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
	OnAfterATTRIBUTE_VALUE_XHTMLCreateCallback OnAfterCreateInterface[ATTRIBUTE_VALUE_XHTML]
	OnAfterATTRIBUTE_VALUE_XHTMLUpdateCallback OnAfterUpdateInterface[ATTRIBUTE_VALUE_XHTML]
	OnAfterATTRIBUTE_VALUE_XHTMLDeleteCallback OnAfterDeleteInterface[ATTRIBUTE_VALUE_XHTML]
	OnAfterATTRIBUTE_VALUE_XHTMLReadCallback   OnAfterReadInterface[ATTRIBUTE_VALUE_XHTML]

	A_ALTERNATIVE_IDs           map[*A_ALTERNATIVE_ID]any
	A_ALTERNATIVE_IDs_mapString map[string]*A_ALTERNATIVE_ID

	// insertion point for slice of pointers maps
	OnAfterA_ALTERNATIVE_IDCreateCallback OnAfterCreateInterface[A_ALTERNATIVE_ID]
	OnAfterA_ALTERNATIVE_IDUpdateCallback OnAfterUpdateInterface[A_ALTERNATIVE_ID]
	OnAfterA_ALTERNATIVE_IDDeleteCallback OnAfterDeleteInterface[A_ALTERNATIVE_ID]
	OnAfterA_ALTERNATIVE_IDReadCallback   OnAfterReadInterface[A_ALTERNATIVE_ID]

	A_ATTRIBUTE_DEFINITION_BOOLEAN_REFs           map[*A_ATTRIBUTE_DEFINITION_BOOLEAN_REF]any
	A_ATTRIBUTE_DEFINITION_BOOLEAN_REFs_mapString map[string]*A_ATTRIBUTE_DEFINITION_BOOLEAN_REF

	// insertion point for slice of pointers maps
	OnAfterA_ATTRIBUTE_DEFINITION_BOOLEAN_REFCreateCallback OnAfterCreateInterface[A_ATTRIBUTE_DEFINITION_BOOLEAN_REF]
	OnAfterA_ATTRIBUTE_DEFINITION_BOOLEAN_REFUpdateCallback OnAfterUpdateInterface[A_ATTRIBUTE_DEFINITION_BOOLEAN_REF]
	OnAfterA_ATTRIBUTE_DEFINITION_BOOLEAN_REFDeleteCallback OnAfterDeleteInterface[A_ATTRIBUTE_DEFINITION_BOOLEAN_REF]
	OnAfterA_ATTRIBUTE_DEFINITION_BOOLEAN_REFReadCallback   OnAfterReadInterface[A_ATTRIBUTE_DEFINITION_BOOLEAN_REF]

	A_ATTRIBUTE_DEFINITION_DATE_REFs           map[*A_ATTRIBUTE_DEFINITION_DATE_REF]any
	A_ATTRIBUTE_DEFINITION_DATE_REFs_mapString map[string]*A_ATTRIBUTE_DEFINITION_DATE_REF

	// insertion point for slice of pointers maps
	OnAfterA_ATTRIBUTE_DEFINITION_DATE_REFCreateCallback OnAfterCreateInterface[A_ATTRIBUTE_DEFINITION_DATE_REF]
	OnAfterA_ATTRIBUTE_DEFINITION_DATE_REFUpdateCallback OnAfterUpdateInterface[A_ATTRIBUTE_DEFINITION_DATE_REF]
	OnAfterA_ATTRIBUTE_DEFINITION_DATE_REFDeleteCallback OnAfterDeleteInterface[A_ATTRIBUTE_DEFINITION_DATE_REF]
	OnAfterA_ATTRIBUTE_DEFINITION_DATE_REFReadCallback   OnAfterReadInterface[A_ATTRIBUTE_DEFINITION_DATE_REF]

	A_ATTRIBUTE_DEFINITION_ENUMERATION_REFs           map[*A_ATTRIBUTE_DEFINITION_ENUMERATION_REF]any
	A_ATTRIBUTE_DEFINITION_ENUMERATION_REFs_mapString map[string]*A_ATTRIBUTE_DEFINITION_ENUMERATION_REF

	// insertion point for slice of pointers maps
	OnAfterA_ATTRIBUTE_DEFINITION_ENUMERATION_REFCreateCallback OnAfterCreateInterface[A_ATTRIBUTE_DEFINITION_ENUMERATION_REF]
	OnAfterA_ATTRIBUTE_DEFINITION_ENUMERATION_REFUpdateCallback OnAfterUpdateInterface[A_ATTRIBUTE_DEFINITION_ENUMERATION_REF]
	OnAfterA_ATTRIBUTE_DEFINITION_ENUMERATION_REFDeleteCallback OnAfterDeleteInterface[A_ATTRIBUTE_DEFINITION_ENUMERATION_REF]
	OnAfterA_ATTRIBUTE_DEFINITION_ENUMERATION_REFReadCallback   OnAfterReadInterface[A_ATTRIBUTE_DEFINITION_ENUMERATION_REF]

	A_ATTRIBUTE_DEFINITION_INTEGER_REFs           map[*A_ATTRIBUTE_DEFINITION_INTEGER_REF]any
	A_ATTRIBUTE_DEFINITION_INTEGER_REFs_mapString map[string]*A_ATTRIBUTE_DEFINITION_INTEGER_REF

	// insertion point for slice of pointers maps
	OnAfterA_ATTRIBUTE_DEFINITION_INTEGER_REFCreateCallback OnAfterCreateInterface[A_ATTRIBUTE_DEFINITION_INTEGER_REF]
	OnAfterA_ATTRIBUTE_DEFINITION_INTEGER_REFUpdateCallback OnAfterUpdateInterface[A_ATTRIBUTE_DEFINITION_INTEGER_REF]
	OnAfterA_ATTRIBUTE_DEFINITION_INTEGER_REFDeleteCallback OnAfterDeleteInterface[A_ATTRIBUTE_DEFINITION_INTEGER_REF]
	OnAfterA_ATTRIBUTE_DEFINITION_INTEGER_REFReadCallback   OnAfterReadInterface[A_ATTRIBUTE_DEFINITION_INTEGER_REF]

	A_ATTRIBUTE_DEFINITION_REAL_REFs           map[*A_ATTRIBUTE_DEFINITION_REAL_REF]any
	A_ATTRIBUTE_DEFINITION_REAL_REFs_mapString map[string]*A_ATTRIBUTE_DEFINITION_REAL_REF

	// insertion point for slice of pointers maps
	OnAfterA_ATTRIBUTE_DEFINITION_REAL_REFCreateCallback OnAfterCreateInterface[A_ATTRIBUTE_DEFINITION_REAL_REF]
	OnAfterA_ATTRIBUTE_DEFINITION_REAL_REFUpdateCallback OnAfterUpdateInterface[A_ATTRIBUTE_DEFINITION_REAL_REF]
	OnAfterA_ATTRIBUTE_DEFINITION_REAL_REFDeleteCallback OnAfterDeleteInterface[A_ATTRIBUTE_DEFINITION_REAL_REF]
	OnAfterA_ATTRIBUTE_DEFINITION_REAL_REFReadCallback   OnAfterReadInterface[A_ATTRIBUTE_DEFINITION_REAL_REF]

	A_ATTRIBUTE_DEFINITION_STRING_REFs           map[*A_ATTRIBUTE_DEFINITION_STRING_REF]any
	A_ATTRIBUTE_DEFINITION_STRING_REFs_mapString map[string]*A_ATTRIBUTE_DEFINITION_STRING_REF

	// insertion point for slice of pointers maps
	OnAfterA_ATTRIBUTE_DEFINITION_STRING_REFCreateCallback OnAfterCreateInterface[A_ATTRIBUTE_DEFINITION_STRING_REF]
	OnAfterA_ATTRIBUTE_DEFINITION_STRING_REFUpdateCallback OnAfterUpdateInterface[A_ATTRIBUTE_DEFINITION_STRING_REF]
	OnAfterA_ATTRIBUTE_DEFINITION_STRING_REFDeleteCallback OnAfterDeleteInterface[A_ATTRIBUTE_DEFINITION_STRING_REF]
	OnAfterA_ATTRIBUTE_DEFINITION_STRING_REFReadCallback   OnAfterReadInterface[A_ATTRIBUTE_DEFINITION_STRING_REF]

	A_ATTRIBUTE_DEFINITION_XHTML_REFs           map[*A_ATTRIBUTE_DEFINITION_XHTML_REF]any
	A_ATTRIBUTE_DEFINITION_XHTML_REFs_mapString map[string]*A_ATTRIBUTE_DEFINITION_XHTML_REF

	// insertion point for slice of pointers maps
	OnAfterA_ATTRIBUTE_DEFINITION_XHTML_REFCreateCallback OnAfterCreateInterface[A_ATTRIBUTE_DEFINITION_XHTML_REF]
	OnAfterA_ATTRIBUTE_DEFINITION_XHTML_REFUpdateCallback OnAfterUpdateInterface[A_ATTRIBUTE_DEFINITION_XHTML_REF]
	OnAfterA_ATTRIBUTE_DEFINITION_XHTML_REFDeleteCallback OnAfterDeleteInterface[A_ATTRIBUTE_DEFINITION_XHTML_REF]
	OnAfterA_ATTRIBUTE_DEFINITION_XHTML_REFReadCallback   OnAfterReadInterface[A_ATTRIBUTE_DEFINITION_XHTML_REF]

	A_ATTRIBUTE_VALUE_BOOLEANs           map[*A_ATTRIBUTE_VALUE_BOOLEAN]any
	A_ATTRIBUTE_VALUE_BOOLEANs_mapString map[string]*A_ATTRIBUTE_VALUE_BOOLEAN

	// insertion point for slice of pointers maps
	OnAfterA_ATTRIBUTE_VALUE_BOOLEANCreateCallback OnAfterCreateInterface[A_ATTRIBUTE_VALUE_BOOLEAN]
	OnAfterA_ATTRIBUTE_VALUE_BOOLEANUpdateCallback OnAfterUpdateInterface[A_ATTRIBUTE_VALUE_BOOLEAN]
	OnAfterA_ATTRIBUTE_VALUE_BOOLEANDeleteCallback OnAfterDeleteInterface[A_ATTRIBUTE_VALUE_BOOLEAN]
	OnAfterA_ATTRIBUTE_VALUE_BOOLEANReadCallback   OnAfterReadInterface[A_ATTRIBUTE_VALUE_BOOLEAN]

	A_ATTRIBUTE_VALUE_DATEs           map[*A_ATTRIBUTE_VALUE_DATE]any
	A_ATTRIBUTE_VALUE_DATEs_mapString map[string]*A_ATTRIBUTE_VALUE_DATE

	// insertion point for slice of pointers maps
	OnAfterA_ATTRIBUTE_VALUE_DATECreateCallback OnAfterCreateInterface[A_ATTRIBUTE_VALUE_DATE]
	OnAfterA_ATTRIBUTE_VALUE_DATEUpdateCallback OnAfterUpdateInterface[A_ATTRIBUTE_VALUE_DATE]
	OnAfterA_ATTRIBUTE_VALUE_DATEDeleteCallback OnAfterDeleteInterface[A_ATTRIBUTE_VALUE_DATE]
	OnAfterA_ATTRIBUTE_VALUE_DATEReadCallback   OnAfterReadInterface[A_ATTRIBUTE_VALUE_DATE]

	A_ATTRIBUTE_VALUE_ENUMERATIONs           map[*A_ATTRIBUTE_VALUE_ENUMERATION]any
	A_ATTRIBUTE_VALUE_ENUMERATIONs_mapString map[string]*A_ATTRIBUTE_VALUE_ENUMERATION

	// insertion point for slice of pointers maps
	OnAfterA_ATTRIBUTE_VALUE_ENUMERATIONCreateCallback OnAfterCreateInterface[A_ATTRIBUTE_VALUE_ENUMERATION]
	OnAfterA_ATTRIBUTE_VALUE_ENUMERATIONUpdateCallback OnAfterUpdateInterface[A_ATTRIBUTE_VALUE_ENUMERATION]
	OnAfterA_ATTRIBUTE_VALUE_ENUMERATIONDeleteCallback OnAfterDeleteInterface[A_ATTRIBUTE_VALUE_ENUMERATION]
	OnAfterA_ATTRIBUTE_VALUE_ENUMERATIONReadCallback   OnAfterReadInterface[A_ATTRIBUTE_VALUE_ENUMERATION]

	A_ATTRIBUTE_VALUE_INTEGERs           map[*A_ATTRIBUTE_VALUE_INTEGER]any
	A_ATTRIBUTE_VALUE_INTEGERs_mapString map[string]*A_ATTRIBUTE_VALUE_INTEGER

	// insertion point for slice of pointers maps
	OnAfterA_ATTRIBUTE_VALUE_INTEGERCreateCallback OnAfterCreateInterface[A_ATTRIBUTE_VALUE_INTEGER]
	OnAfterA_ATTRIBUTE_VALUE_INTEGERUpdateCallback OnAfterUpdateInterface[A_ATTRIBUTE_VALUE_INTEGER]
	OnAfterA_ATTRIBUTE_VALUE_INTEGERDeleteCallback OnAfterDeleteInterface[A_ATTRIBUTE_VALUE_INTEGER]
	OnAfterA_ATTRIBUTE_VALUE_INTEGERReadCallback   OnAfterReadInterface[A_ATTRIBUTE_VALUE_INTEGER]

	A_ATTRIBUTE_VALUE_REALs           map[*A_ATTRIBUTE_VALUE_REAL]any
	A_ATTRIBUTE_VALUE_REALs_mapString map[string]*A_ATTRIBUTE_VALUE_REAL

	// insertion point for slice of pointers maps
	OnAfterA_ATTRIBUTE_VALUE_REALCreateCallback OnAfterCreateInterface[A_ATTRIBUTE_VALUE_REAL]
	OnAfterA_ATTRIBUTE_VALUE_REALUpdateCallback OnAfterUpdateInterface[A_ATTRIBUTE_VALUE_REAL]
	OnAfterA_ATTRIBUTE_VALUE_REALDeleteCallback OnAfterDeleteInterface[A_ATTRIBUTE_VALUE_REAL]
	OnAfterA_ATTRIBUTE_VALUE_REALReadCallback   OnAfterReadInterface[A_ATTRIBUTE_VALUE_REAL]

	A_ATTRIBUTE_VALUE_STRINGs           map[*A_ATTRIBUTE_VALUE_STRING]any
	A_ATTRIBUTE_VALUE_STRINGs_mapString map[string]*A_ATTRIBUTE_VALUE_STRING

	// insertion point for slice of pointers maps
	OnAfterA_ATTRIBUTE_VALUE_STRINGCreateCallback OnAfterCreateInterface[A_ATTRIBUTE_VALUE_STRING]
	OnAfterA_ATTRIBUTE_VALUE_STRINGUpdateCallback OnAfterUpdateInterface[A_ATTRIBUTE_VALUE_STRING]
	OnAfterA_ATTRIBUTE_VALUE_STRINGDeleteCallback OnAfterDeleteInterface[A_ATTRIBUTE_VALUE_STRING]
	OnAfterA_ATTRIBUTE_VALUE_STRINGReadCallback   OnAfterReadInterface[A_ATTRIBUTE_VALUE_STRING]

	A_ATTRIBUTE_VALUE_XHTMLs           map[*A_ATTRIBUTE_VALUE_XHTML]any
	A_ATTRIBUTE_VALUE_XHTMLs_mapString map[string]*A_ATTRIBUTE_VALUE_XHTML

	// insertion point for slice of pointers maps
	OnAfterA_ATTRIBUTE_VALUE_XHTMLCreateCallback OnAfterCreateInterface[A_ATTRIBUTE_VALUE_XHTML]
	OnAfterA_ATTRIBUTE_VALUE_XHTMLUpdateCallback OnAfterUpdateInterface[A_ATTRIBUTE_VALUE_XHTML]
	OnAfterA_ATTRIBUTE_VALUE_XHTMLDeleteCallback OnAfterDeleteInterface[A_ATTRIBUTE_VALUE_XHTML]
	OnAfterA_ATTRIBUTE_VALUE_XHTMLReadCallback   OnAfterReadInterface[A_ATTRIBUTE_VALUE_XHTML]

	A_ATTRIBUTE_VALUE_XHTML_1s           map[*A_ATTRIBUTE_VALUE_XHTML_1]any
	A_ATTRIBUTE_VALUE_XHTML_1s_mapString map[string]*A_ATTRIBUTE_VALUE_XHTML_1

	// insertion point for slice of pointers maps
	OnAfterA_ATTRIBUTE_VALUE_XHTML_1CreateCallback OnAfterCreateInterface[A_ATTRIBUTE_VALUE_XHTML_1]
	OnAfterA_ATTRIBUTE_VALUE_XHTML_1UpdateCallback OnAfterUpdateInterface[A_ATTRIBUTE_VALUE_XHTML_1]
	OnAfterA_ATTRIBUTE_VALUE_XHTML_1DeleteCallback OnAfterDeleteInterface[A_ATTRIBUTE_VALUE_XHTML_1]
	OnAfterA_ATTRIBUTE_VALUE_XHTML_1ReadCallback   OnAfterReadInterface[A_ATTRIBUTE_VALUE_XHTML_1]

	A_CHILDRENs           map[*A_CHILDREN]any
	A_CHILDRENs_mapString map[string]*A_CHILDREN

	// insertion point for slice of pointers maps
	OnAfterA_CHILDRENCreateCallback OnAfterCreateInterface[A_CHILDREN]
	OnAfterA_CHILDRENUpdateCallback OnAfterUpdateInterface[A_CHILDREN]
	OnAfterA_CHILDRENDeleteCallback OnAfterDeleteInterface[A_CHILDREN]
	OnAfterA_CHILDRENReadCallback   OnAfterReadInterface[A_CHILDREN]

	A_CORE_CONTENTs           map[*A_CORE_CONTENT]any
	A_CORE_CONTENTs_mapString map[string]*A_CORE_CONTENT

	// insertion point for slice of pointers maps
	OnAfterA_CORE_CONTENTCreateCallback OnAfterCreateInterface[A_CORE_CONTENT]
	OnAfterA_CORE_CONTENTUpdateCallback OnAfterUpdateInterface[A_CORE_CONTENT]
	OnAfterA_CORE_CONTENTDeleteCallback OnAfterDeleteInterface[A_CORE_CONTENT]
	OnAfterA_CORE_CONTENTReadCallback   OnAfterReadInterface[A_CORE_CONTENT]

	A_DATATYPESs           map[*A_DATATYPES]any
	A_DATATYPESs_mapString map[string]*A_DATATYPES

	// insertion point for slice of pointers maps
	OnAfterA_DATATYPESCreateCallback OnAfterCreateInterface[A_DATATYPES]
	OnAfterA_DATATYPESUpdateCallback OnAfterUpdateInterface[A_DATATYPES]
	OnAfterA_DATATYPESDeleteCallback OnAfterDeleteInterface[A_DATATYPES]
	OnAfterA_DATATYPESReadCallback   OnAfterReadInterface[A_DATATYPES]

	A_DATATYPE_DEFINITION_BOOLEAN_REFs           map[*A_DATATYPE_DEFINITION_BOOLEAN_REF]any
	A_DATATYPE_DEFINITION_BOOLEAN_REFs_mapString map[string]*A_DATATYPE_DEFINITION_BOOLEAN_REF

	// insertion point for slice of pointers maps
	OnAfterA_DATATYPE_DEFINITION_BOOLEAN_REFCreateCallback OnAfterCreateInterface[A_DATATYPE_DEFINITION_BOOLEAN_REF]
	OnAfterA_DATATYPE_DEFINITION_BOOLEAN_REFUpdateCallback OnAfterUpdateInterface[A_DATATYPE_DEFINITION_BOOLEAN_REF]
	OnAfterA_DATATYPE_DEFINITION_BOOLEAN_REFDeleteCallback OnAfterDeleteInterface[A_DATATYPE_DEFINITION_BOOLEAN_REF]
	OnAfterA_DATATYPE_DEFINITION_BOOLEAN_REFReadCallback   OnAfterReadInterface[A_DATATYPE_DEFINITION_BOOLEAN_REF]

	A_DATATYPE_DEFINITION_DATE_REFs           map[*A_DATATYPE_DEFINITION_DATE_REF]any
	A_DATATYPE_DEFINITION_DATE_REFs_mapString map[string]*A_DATATYPE_DEFINITION_DATE_REF

	// insertion point for slice of pointers maps
	OnAfterA_DATATYPE_DEFINITION_DATE_REFCreateCallback OnAfterCreateInterface[A_DATATYPE_DEFINITION_DATE_REF]
	OnAfterA_DATATYPE_DEFINITION_DATE_REFUpdateCallback OnAfterUpdateInterface[A_DATATYPE_DEFINITION_DATE_REF]
	OnAfterA_DATATYPE_DEFINITION_DATE_REFDeleteCallback OnAfterDeleteInterface[A_DATATYPE_DEFINITION_DATE_REF]
	OnAfterA_DATATYPE_DEFINITION_DATE_REFReadCallback   OnAfterReadInterface[A_DATATYPE_DEFINITION_DATE_REF]

	A_DATATYPE_DEFINITION_ENUMERATION_REFs           map[*A_DATATYPE_DEFINITION_ENUMERATION_REF]any
	A_DATATYPE_DEFINITION_ENUMERATION_REFs_mapString map[string]*A_DATATYPE_DEFINITION_ENUMERATION_REF

	// insertion point for slice of pointers maps
	OnAfterA_DATATYPE_DEFINITION_ENUMERATION_REFCreateCallback OnAfterCreateInterface[A_DATATYPE_DEFINITION_ENUMERATION_REF]
	OnAfterA_DATATYPE_DEFINITION_ENUMERATION_REFUpdateCallback OnAfterUpdateInterface[A_DATATYPE_DEFINITION_ENUMERATION_REF]
	OnAfterA_DATATYPE_DEFINITION_ENUMERATION_REFDeleteCallback OnAfterDeleteInterface[A_DATATYPE_DEFINITION_ENUMERATION_REF]
	OnAfterA_DATATYPE_DEFINITION_ENUMERATION_REFReadCallback   OnAfterReadInterface[A_DATATYPE_DEFINITION_ENUMERATION_REF]

	A_DATATYPE_DEFINITION_INTEGER_REFs           map[*A_DATATYPE_DEFINITION_INTEGER_REF]any
	A_DATATYPE_DEFINITION_INTEGER_REFs_mapString map[string]*A_DATATYPE_DEFINITION_INTEGER_REF

	// insertion point for slice of pointers maps
	OnAfterA_DATATYPE_DEFINITION_INTEGER_REFCreateCallback OnAfterCreateInterface[A_DATATYPE_DEFINITION_INTEGER_REF]
	OnAfterA_DATATYPE_DEFINITION_INTEGER_REFUpdateCallback OnAfterUpdateInterface[A_DATATYPE_DEFINITION_INTEGER_REF]
	OnAfterA_DATATYPE_DEFINITION_INTEGER_REFDeleteCallback OnAfterDeleteInterface[A_DATATYPE_DEFINITION_INTEGER_REF]
	OnAfterA_DATATYPE_DEFINITION_INTEGER_REFReadCallback   OnAfterReadInterface[A_DATATYPE_DEFINITION_INTEGER_REF]

	A_DATATYPE_DEFINITION_REAL_REFs           map[*A_DATATYPE_DEFINITION_REAL_REF]any
	A_DATATYPE_DEFINITION_REAL_REFs_mapString map[string]*A_DATATYPE_DEFINITION_REAL_REF

	// insertion point for slice of pointers maps
	OnAfterA_DATATYPE_DEFINITION_REAL_REFCreateCallback OnAfterCreateInterface[A_DATATYPE_DEFINITION_REAL_REF]
	OnAfterA_DATATYPE_DEFINITION_REAL_REFUpdateCallback OnAfterUpdateInterface[A_DATATYPE_DEFINITION_REAL_REF]
	OnAfterA_DATATYPE_DEFINITION_REAL_REFDeleteCallback OnAfterDeleteInterface[A_DATATYPE_DEFINITION_REAL_REF]
	OnAfterA_DATATYPE_DEFINITION_REAL_REFReadCallback   OnAfterReadInterface[A_DATATYPE_DEFINITION_REAL_REF]

	A_DATATYPE_DEFINITION_STRING_REFs           map[*A_DATATYPE_DEFINITION_STRING_REF]any
	A_DATATYPE_DEFINITION_STRING_REFs_mapString map[string]*A_DATATYPE_DEFINITION_STRING_REF

	// insertion point for slice of pointers maps
	OnAfterA_DATATYPE_DEFINITION_STRING_REFCreateCallback OnAfterCreateInterface[A_DATATYPE_DEFINITION_STRING_REF]
	OnAfterA_DATATYPE_DEFINITION_STRING_REFUpdateCallback OnAfterUpdateInterface[A_DATATYPE_DEFINITION_STRING_REF]
	OnAfterA_DATATYPE_DEFINITION_STRING_REFDeleteCallback OnAfterDeleteInterface[A_DATATYPE_DEFINITION_STRING_REF]
	OnAfterA_DATATYPE_DEFINITION_STRING_REFReadCallback   OnAfterReadInterface[A_DATATYPE_DEFINITION_STRING_REF]

	A_DATATYPE_DEFINITION_XHTML_REFs           map[*A_DATATYPE_DEFINITION_XHTML_REF]any
	A_DATATYPE_DEFINITION_XHTML_REFs_mapString map[string]*A_DATATYPE_DEFINITION_XHTML_REF

	// insertion point for slice of pointers maps
	OnAfterA_DATATYPE_DEFINITION_XHTML_REFCreateCallback OnAfterCreateInterface[A_DATATYPE_DEFINITION_XHTML_REF]
	OnAfterA_DATATYPE_DEFINITION_XHTML_REFUpdateCallback OnAfterUpdateInterface[A_DATATYPE_DEFINITION_XHTML_REF]
	OnAfterA_DATATYPE_DEFINITION_XHTML_REFDeleteCallback OnAfterDeleteInterface[A_DATATYPE_DEFINITION_XHTML_REF]
	OnAfterA_DATATYPE_DEFINITION_XHTML_REFReadCallback   OnAfterReadInterface[A_DATATYPE_DEFINITION_XHTML_REF]

	A_EDITABLE_ATTSs           map[*A_EDITABLE_ATTS]any
	A_EDITABLE_ATTSs_mapString map[string]*A_EDITABLE_ATTS

	// insertion point for slice of pointers maps
	OnAfterA_EDITABLE_ATTSCreateCallback OnAfterCreateInterface[A_EDITABLE_ATTS]
	OnAfterA_EDITABLE_ATTSUpdateCallback OnAfterUpdateInterface[A_EDITABLE_ATTS]
	OnAfterA_EDITABLE_ATTSDeleteCallback OnAfterDeleteInterface[A_EDITABLE_ATTS]
	OnAfterA_EDITABLE_ATTSReadCallback   OnAfterReadInterface[A_EDITABLE_ATTS]

	A_ENUM_VALUE_REFs           map[*A_ENUM_VALUE_REF]any
	A_ENUM_VALUE_REFs_mapString map[string]*A_ENUM_VALUE_REF

	// insertion point for slice of pointers maps
	OnAfterA_ENUM_VALUE_REFCreateCallback OnAfterCreateInterface[A_ENUM_VALUE_REF]
	OnAfterA_ENUM_VALUE_REFUpdateCallback OnAfterUpdateInterface[A_ENUM_VALUE_REF]
	OnAfterA_ENUM_VALUE_REFDeleteCallback OnAfterDeleteInterface[A_ENUM_VALUE_REF]
	OnAfterA_ENUM_VALUE_REFReadCallback   OnAfterReadInterface[A_ENUM_VALUE_REF]

	A_OBJECTs           map[*A_OBJECT]any
	A_OBJECTs_mapString map[string]*A_OBJECT

	// insertion point for slice of pointers maps
	OnAfterA_OBJECTCreateCallback OnAfterCreateInterface[A_OBJECT]
	OnAfterA_OBJECTUpdateCallback OnAfterUpdateInterface[A_OBJECT]
	OnAfterA_OBJECTDeleteCallback OnAfterDeleteInterface[A_OBJECT]
	OnAfterA_OBJECTReadCallback   OnAfterReadInterface[A_OBJECT]

	A_PROPERTIESs           map[*A_PROPERTIES]any
	A_PROPERTIESs_mapString map[string]*A_PROPERTIES

	// insertion point for slice of pointers maps
	OnAfterA_PROPERTIESCreateCallback OnAfterCreateInterface[A_PROPERTIES]
	OnAfterA_PROPERTIESUpdateCallback OnAfterUpdateInterface[A_PROPERTIES]
	OnAfterA_PROPERTIESDeleteCallback OnAfterDeleteInterface[A_PROPERTIES]
	OnAfterA_PROPERTIESReadCallback   OnAfterReadInterface[A_PROPERTIES]

	A_RELATION_GROUP_TYPE_REFs           map[*A_RELATION_GROUP_TYPE_REF]any
	A_RELATION_GROUP_TYPE_REFs_mapString map[string]*A_RELATION_GROUP_TYPE_REF

	// insertion point for slice of pointers maps
	OnAfterA_RELATION_GROUP_TYPE_REFCreateCallback OnAfterCreateInterface[A_RELATION_GROUP_TYPE_REF]
	OnAfterA_RELATION_GROUP_TYPE_REFUpdateCallback OnAfterUpdateInterface[A_RELATION_GROUP_TYPE_REF]
	OnAfterA_RELATION_GROUP_TYPE_REFDeleteCallback OnAfterDeleteInterface[A_RELATION_GROUP_TYPE_REF]
	OnAfterA_RELATION_GROUP_TYPE_REFReadCallback   OnAfterReadInterface[A_RELATION_GROUP_TYPE_REF]

	A_SOURCE_1s           map[*A_SOURCE_1]any
	A_SOURCE_1s_mapString map[string]*A_SOURCE_1

	// insertion point for slice of pointers maps
	OnAfterA_SOURCE_1CreateCallback OnAfterCreateInterface[A_SOURCE_1]
	OnAfterA_SOURCE_1UpdateCallback OnAfterUpdateInterface[A_SOURCE_1]
	OnAfterA_SOURCE_1DeleteCallback OnAfterDeleteInterface[A_SOURCE_1]
	OnAfterA_SOURCE_1ReadCallback   OnAfterReadInterface[A_SOURCE_1]

	A_SPECIFICATIONSs           map[*A_SPECIFICATIONS]any
	A_SPECIFICATIONSs_mapString map[string]*A_SPECIFICATIONS

	// insertion point for slice of pointers maps
	OnAfterA_SPECIFICATIONSCreateCallback OnAfterCreateInterface[A_SPECIFICATIONS]
	OnAfterA_SPECIFICATIONSUpdateCallback OnAfterUpdateInterface[A_SPECIFICATIONS]
	OnAfterA_SPECIFICATIONSDeleteCallback OnAfterDeleteInterface[A_SPECIFICATIONS]
	OnAfterA_SPECIFICATIONSReadCallback   OnAfterReadInterface[A_SPECIFICATIONS]

	A_SPECIFICATION_TYPE_REFs           map[*A_SPECIFICATION_TYPE_REF]any
	A_SPECIFICATION_TYPE_REFs_mapString map[string]*A_SPECIFICATION_TYPE_REF

	// insertion point for slice of pointers maps
	OnAfterA_SPECIFICATION_TYPE_REFCreateCallback OnAfterCreateInterface[A_SPECIFICATION_TYPE_REF]
	OnAfterA_SPECIFICATION_TYPE_REFUpdateCallback OnAfterUpdateInterface[A_SPECIFICATION_TYPE_REF]
	OnAfterA_SPECIFICATION_TYPE_REFDeleteCallback OnAfterDeleteInterface[A_SPECIFICATION_TYPE_REF]
	OnAfterA_SPECIFICATION_TYPE_REFReadCallback   OnAfterReadInterface[A_SPECIFICATION_TYPE_REF]

	A_SPECIFIED_VALUESs           map[*A_SPECIFIED_VALUES]any
	A_SPECIFIED_VALUESs_mapString map[string]*A_SPECIFIED_VALUES

	// insertion point for slice of pointers maps
	OnAfterA_SPECIFIED_VALUESCreateCallback OnAfterCreateInterface[A_SPECIFIED_VALUES]
	OnAfterA_SPECIFIED_VALUESUpdateCallback OnAfterUpdateInterface[A_SPECIFIED_VALUES]
	OnAfterA_SPECIFIED_VALUESDeleteCallback OnAfterDeleteInterface[A_SPECIFIED_VALUES]
	OnAfterA_SPECIFIED_VALUESReadCallback   OnAfterReadInterface[A_SPECIFIED_VALUES]

	A_SPEC_ATTRIBUTESs           map[*A_SPEC_ATTRIBUTES]any
	A_SPEC_ATTRIBUTESs_mapString map[string]*A_SPEC_ATTRIBUTES

	// insertion point for slice of pointers maps
	OnAfterA_SPEC_ATTRIBUTESCreateCallback OnAfterCreateInterface[A_SPEC_ATTRIBUTES]
	OnAfterA_SPEC_ATTRIBUTESUpdateCallback OnAfterUpdateInterface[A_SPEC_ATTRIBUTES]
	OnAfterA_SPEC_ATTRIBUTESDeleteCallback OnAfterDeleteInterface[A_SPEC_ATTRIBUTES]
	OnAfterA_SPEC_ATTRIBUTESReadCallback   OnAfterReadInterface[A_SPEC_ATTRIBUTES]

	A_SPEC_OBJECTSs           map[*A_SPEC_OBJECTS]any
	A_SPEC_OBJECTSs_mapString map[string]*A_SPEC_OBJECTS

	// insertion point for slice of pointers maps
	OnAfterA_SPEC_OBJECTSCreateCallback OnAfterCreateInterface[A_SPEC_OBJECTS]
	OnAfterA_SPEC_OBJECTSUpdateCallback OnAfterUpdateInterface[A_SPEC_OBJECTS]
	OnAfterA_SPEC_OBJECTSDeleteCallback OnAfterDeleteInterface[A_SPEC_OBJECTS]
	OnAfterA_SPEC_OBJECTSReadCallback   OnAfterReadInterface[A_SPEC_OBJECTS]

	A_SPEC_OBJECT_TYPE_REFs           map[*A_SPEC_OBJECT_TYPE_REF]any
	A_SPEC_OBJECT_TYPE_REFs_mapString map[string]*A_SPEC_OBJECT_TYPE_REF

	// insertion point for slice of pointers maps
	OnAfterA_SPEC_OBJECT_TYPE_REFCreateCallback OnAfterCreateInterface[A_SPEC_OBJECT_TYPE_REF]
	OnAfterA_SPEC_OBJECT_TYPE_REFUpdateCallback OnAfterUpdateInterface[A_SPEC_OBJECT_TYPE_REF]
	OnAfterA_SPEC_OBJECT_TYPE_REFDeleteCallback OnAfterDeleteInterface[A_SPEC_OBJECT_TYPE_REF]
	OnAfterA_SPEC_OBJECT_TYPE_REFReadCallback   OnAfterReadInterface[A_SPEC_OBJECT_TYPE_REF]

	A_SPEC_RELATIONSs           map[*A_SPEC_RELATIONS]any
	A_SPEC_RELATIONSs_mapString map[string]*A_SPEC_RELATIONS

	// insertion point for slice of pointers maps
	OnAfterA_SPEC_RELATIONSCreateCallback OnAfterCreateInterface[A_SPEC_RELATIONS]
	OnAfterA_SPEC_RELATIONSUpdateCallback OnAfterUpdateInterface[A_SPEC_RELATIONS]
	OnAfterA_SPEC_RELATIONSDeleteCallback OnAfterDeleteInterface[A_SPEC_RELATIONS]
	OnAfterA_SPEC_RELATIONSReadCallback   OnAfterReadInterface[A_SPEC_RELATIONS]

	A_SPEC_RELATION_GROUPSs           map[*A_SPEC_RELATION_GROUPS]any
	A_SPEC_RELATION_GROUPSs_mapString map[string]*A_SPEC_RELATION_GROUPS

	// insertion point for slice of pointers maps
	OnAfterA_SPEC_RELATION_GROUPSCreateCallback OnAfterCreateInterface[A_SPEC_RELATION_GROUPS]
	OnAfterA_SPEC_RELATION_GROUPSUpdateCallback OnAfterUpdateInterface[A_SPEC_RELATION_GROUPS]
	OnAfterA_SPEC_RELATION_GROUPSDeleteCallback OnAfterDeleteInterface[A_SPEC_RELATION_GROUPS]
	OnAfterA_SPEC_RELATION_GROUPSReadCallback   OnAfterReadInterface[A_SPEC_RELATION_GROUPS]

	A_SPEC_RELATION_REFs           map[*A_SPEC_RELATION_REF]any
	A_SPEC_RELATION_REFs_mapString map[string]*A_SPEC_RELATION_REF

	// insertion point for slice of pointers maps
	OnAfterA_SPEC_RELATION_REFCreateCallback OnAfterCreateInterface[A_SPEC_RELATION_REF]
	OnAfterA_SPEC_RELATION_REFUpdateCallback OnAfterUpdateInterface[A_SPEC_RELATION_REF]
	OnAfterA_SPEC_RELATION_REFDeleteCallback OnAfterDeleteInterface[A_SPEC_RELATION_REF]
	OnAfterA_SPEC_RELATION_REFReadCallback   OnAfterReadInterface[A_SPEC_RELATION_REF]

	A_SPEC_RELATION_TYPE_REFs           map[*A_SPEC_RELATION_TYPE_REF]any
	A_SPEC_RELATION_TYPE_REFs_mapString map[string]*A_SPEC_RELATION_TYPE_REF

	// insertion point for slice of pointers maps
	OnAfterA_SPEC_RELATION_TYPE_REFCreateCallback OnAfterCreateInterface[A_SPEC_RELATION_TYPE_REF]
	OnAfterA_SPEC_RELATION_TYPE_REFUpdateCallback OnAfterUpdateInterface[A_SPEC_RELATION_TYPE_REF]
	OnAfterA_SPEC_RELATION_TYPE_REFDeleteCallback OnAfterDeleteInterface[A_SPEC_RELATION_TYPE_REF]
	OnAfterA_SPEC_RELATION_TYPE_REFReadCallback   OnAfterReadInterface[A_SPEC_RELATION_TYPE_REF]

	A_SPEC_TYPESs           map[*A_SPEC_TYPES]any
	A_SPEC_TYPESs_mapString map[string]*A_SPEC_TYPES

	// insertion point for slice of pointers maps
	OnAfterA_SPEC_TYPESCreateCallback OnAfterCreateInterface[A_SPEC_TYPES]
	OnAfterA_SPEC_TYPESUpdateCallback OnAfterUpdateInterface[A_SPEC_TYPES]
	OnAfterA_SPEC_TYPESDeleteCallback OnAfterDeleteInterface[A_SPEC_TYPES]
	OnAfterA_SPEC_TYPESReadCallback   OnAfterReadInterface[A_SPEC_TYPES]

	A_TARGET_SPECIFICATION_1s           map[*A_TARGET_SPECIFICATION_1]any
	A_TARGET_SPECIFICATION_1s_mapString map[string]*A_TARGET_SPECIFICATION_1

	// insertion point for slice of pointers maps
	OnAfterA_TARGET_SPECIFICATION_1CreateCallback OnAfterCreateInterface[A_TARGET_SPECIFICATION_1]
	OnAfterA_TARGET_SPECIFICATION_1UpdateCallback OnAfterUpdateInterface[A_TARGET_SPECIFICATION_1]
	OnAfterA_TARGET_SPECIFICATION_1DeleteCallback OnAfterDeleteInterface[A_TARGET_SPECIFICATION_1]
	OnAfterA_TARGET_SPECIFICATION_1ReadCallback   OnAfterReadInterface[A_TARGET_SPECIFICATION_1]

	A_THE_HEADERs           map[*A_THE_HEADER]any
	A_THE_HEADERs_mapString map[string]*A_THE_HEADER

	// insertion point for slice of pointers maps
	OnAfterA_THE_HEADERCreateCallback OnAfterCreateInterface[A_THE_HEADER]
	OnAfterA_THE_HEADERUpdateCallback OnAfterUpdateInterface[A_THE_HEADER]
	OnAfterA_THE_HEADERDeleteCallback OnAfterDeleteInterface[A_THE_HEADER]
	OnAfterA_THE_HEADERReadCallback   OnAfterReadInterface[A_THE_HEADER]

	A_TOOL_EXTENSIONSs           map[*A_TOOL_EXTENSIONS]any
	A_TOOL_EXTENSIONSs_mapString map[string]*A_TOOL_EXTENSIONS

	// insertion point for slice of pointers maps
	OnAfterA_TOOL_EXTENSIONSCreateCallback OnAfterCreateInterface[A_TOOL_EXTENSIONS]
	OnAfterA_TOOL_EXTENSIONSUpdateCallback OnAfterUpdateInterface[A_TOOL_EXTENSIONS]
	OnAfterA_TOOL_EXTENSIONSDeleteCallback OnAfterDeleteInterface[A_TOOL_EXTENSIONS]
	OnAfterA_TOOL_EXTENSIONSReadCallback   OnAfterReadInterface[A_TOOL_EXTENSIONS]

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
	CommitA_ALTERNATIVE_ID(a_alternative_id *A_ALTERNATIVE_ID)
	CheckoutA_ALTERNATIVE_ID(a_alternative_id *A_ALTERNATIVE_ID)
	CommitA_ATTRIBUTE_DEFINITION_BOOLEAN_REF(a_attribute_definition_boolean_ref *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF)
	CheckoutA_ATTRIBUTE_DEFINITION_BOOLEAN_REF(a_attribute_definition_boolean_ref *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF)
	CommitA_ATTRIBUTE_DEFINITION_DATE_REF(a_attribute_definition_date_ref *A_ATTRIBUTE_DEFINITION_DATE_REF)
	CheckoutA_ATTRIBUTE_DEFINITION_DATE_REF(a_attribute_definition_date_ref *A_ATTRIBUTE_DEFINITION_DATE_REF)
	CommitA_ATTRIBUTE_DEFINITION_ENUMERATION_REF(a_attribute_definition_enumeration_ref *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF)
	CheckoutA_ATTRIBUTE_DEFINITION_ENUMERATION_REF(a_attribute_definition_enumeration_ref *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF)
	CommitA_ATTRIBUTE_DEFINITION_INTEGER_REF(a_attribute_definition_integer_ref *A_ATTRIBUTE_DEFINITION_INTEGER_REF)
	CheckoutA_ATTRIBUTE_DEFINITION_INTEGER_REF(a_attribute_definition_integer_ref *A_ATTRIBUTE_DEFINITION_INTEGER_REF)
	CommitA_ATTRIBUTE_DEFINITION_REAL_REF(a_attribute_definition_real_ref *A_ATTRIBUTE_DEFINITION_REAL_REF)
	CheckoutA_ATTRIBUTE_DEFINITION_REAL_REF(a_attribute_definition_real_ref *A_ATTRIBUTE_DEFINITION_REAL_REF)
	CommitA_ATTRIBUTE_DEFINITION_STRING_REF(a_attribute_definition_string_ref *A_ATTRIBUTE_DEFINITION_STRING_REF)
	CheckoutA_ATTRIBUTE_DEFINITION_STRING_REF(a_attribute_definition_string_ref *A_ATTRIBUTE_DEFINITION_STRING_REF)
	CommitA_ATTRIBUTE_DEFINITION_XHTML_REF(a_attribute_definition_xhtml_ref *A_ATTRIBUTE_DEFINITION_XHTML_REF)
	CheckoutA_ATTRIBUTE_DEFINITION_XHTML_REF(a_attribute_definition_xhtml_ref *A_ATTRIBUTE_DEFINITION_XHTML_REF)
	CommitA_ATTRIBUTE_VALUE_BOOLEAN(a_attribute_value_boolean *A_ATTRIBUTE_VALUE_BOOLEAN)
	CheckoutA_ATTRIBUTE_VALUE_BOOLEAN(a_attribute_value_boolean *A_ATTRIBUTE_VALUE_BOOLEAN)
	CommitA_ATTRIBUTE_VALUE_DATE(a_attribute_value_date *A_ATTRIBUTE_VALUE_DATE)
	CheckoutA_ATTRIBUTE_VALUE_DATE(a_attribute_value_date *A_ATTRIBUTE_VALUE_DATE)
	CommitA_ATTRIBUTE_VALUE_ENUMERATION(a_attribute_value_enumeration *A_ATTRIBUTE_VALUE_ENUMERATION)
	CheckoutA_ATTRIBUTE_VALUE_ENUMERATION(a_attribute_value_enumeration *A_ATTRIBUTE_VALUE_ENUMERATION)
	CommitA_ATTRIBUTE_VALUE_INTEGER(a_attribute_value_integer *A_ATTRIBUTE_VALUE_INTEGER)
	CheckoutA_ATTRIBUTE_VALUE_INTEGER(a_attribute_value_integer *A_ATTRIBUTE_VALUE_INTEGER)
	CommitA_ATTRIBUTE_VALUE_REAL(a_attribute_value_real *A_ATTRIBUTE_VALUE_REAL)
	CheckoutA_ATTRIBUTE_VALUE_REAL(a_attribute_value_real *A_ATTRIBUTE_VALUE_REAL)
	CommitA_ATTRIBUTE_VALUE_STRING(a_attribute_value_string *A_ATTRIBUTE_VALUE_STRING)
	CheckoutA_ATTRIBUTE_VALUE_STRING(a_attribute_value_string *A_ATTRIBUTE_VALUE_STRING)
	CommitA_ATTRIBUTE_VALUE_XHTML(a_attribute_value_xhtml *A_ATTRIBUTE_VALUE_XHTML)
	CheckoutA_ATTRIBUTE_VALUE_XHTML(a_attribute_value_xhtml *A_ATTRIBUTE_VALUE_XHTML)
	CommitA_ATTRIBUTE_VALUE_XHTML_1(a_attribute_value_xhtml_1 *A_ATTRIBUTE_VALUE_XHTML_1)
	CheckoutA_ATTRIBUTE_VALUE_XHTML_1(a_attribute_value_xhtml_1 *A_ATTRIBUTE_VALUE_XHTML_1)
	CommitA_CHILDREN(a_children *A_CHILDREN)
	CheckoutA_CHILDREN(a_children *A_CHILDREN)
	CommitA_CORE_CONTENT(a_core_content *A_CORE_CONTENT)
	CheckoutA_CORE_CONTENT(a_core_content *A_CORE_CONTENT)
	CommitA_DATATYPES(a_datatypes *A_DATATYPES)
	CheckoutA_DATATYPES(a_datatypes *A_DATATYPES)
	CommitA_DATATYPE_DEFINITION_BOOLEAN_REF(a_datatype_definition_boolean_ref *A_DATATYPE_DEFINITION_BOOLEAN_REF)
	CheckoutA_DATATYPE_DEFINITION_BOOLEAN_REF(a_datatype_definition_boolean_ref *A_DATATYPE_DEFINITION_BOOLEAN_REF)
	CommitA_DATATYPE_DEFINITION_DATE_REF(a_datatype_definition_date_ref *A_DATATYPE_DEFINITION_DATE_REF)
	CheckoutA_DATATYPE_DEFINITION_DATE_REF(a_datatype_definition_date_ref *A_DATATYPE_DEFINITION_DATE_REF)
	CommitA_DATATYPE_DEFINITION_ENUMERATION_REF(a_datatype_definition_enumeration_ref *A_DATATYPE_DEFINITION_ENUMERATION_REF)
	CheckoutA_DATATYPE_DEFINITION_ENUMERATION_REF(a_datatype_definition_enumeration_ref *A_DATATYPE_DEFINITION_ENUMERATION_REF)
	CommitA_DATATYPE_DEFINITION_INTEGER_REF(a_datatype_definition_integer_ref *A_DATATYPE_DEFINITION_INTEGER_REF)
	CheckoutA_DATATYPE_DEFINITION_INTEGER_REF(a_datatype_definition_integer_ref *A_DATATYPE_DEFINITION_INTEGER_REF)
	CommitA_DATATYPE_DEFINITION_REAL_REF(a_datatype_definition_real_ref *A_DATATYPE_DEFINITION_REAL_REF)
	CheckoutA_DATATYPE_DEFINITION_REAL_REF(a_datatype_definition_real_ref *A_DATATYPE_DEFINITION_REAL_REF)
	CommitA_DATATYPE_DEFINITION_STRING_REF(a_datatype_definition_string_ref *A_DATATYPE_DEFINITION_STRING_REF)
	CheckoutA_DATATYPE_DEFINITION_STRING_REF(a_datatype_definition_string_ref *A_DATATYPE_DEFINITION_STRING_REF)
	CommitA_DATATYPE_DEFINITION_XHTML_REF(a_datatype_definition_xhtml_ref *A_DATATYPE_DEFINITION_XHTML_REF)
	CheckoutA_DATATYPE_DEFINITION_XHTML_REF(a_datatype_definition_xhtml_ref *A_DATATYPE_DEFINITION_XHTML_REF)
	CommitA_EDITABLE_ATTS(a_editable_atts *A_EDITABLE_ATTS)
	CheckoutA_EDITABLE_ATTS(a_editable_atts *A_EDITABLE_ATTS)
	CommitA_ENUM_VALUE_REF(a_enum_value_ref *A_ENUM_VALUE_REF)
	CheckoutA_ENUM_VALUE_REF(a_enum_value_ref *A_ENUM_VALUE_REF)
	CommitA_OBJECT(a_object *A_OBJECT)
	CheckoutA_OBJECT(a_object *A_OBJECT)
	CommitA_PROPERTIES(a_properties *A_PROPERTIES)
	CheckoutA_PROPERTIES(a_properties *A_PROPERTIES)
	CommitA_RELATION_GROUP_TYPE_REF(a_relation_group_type_ref *A_RELATION_GROUP_TYPE_REF)
	CheckoutA_RELATION_GROUP_TYPE_REF(a_relation_group_type_ref *A_RELATION_GROUP_TYPE_REF)
	CommitA_SOURCE_1(a_source_1 *A_SOURCE_1)
	CheckoutA_SOURCE_1(a_source_1 *A_SOURCE_1)
	CommitA_SPECIFICATIONS(a_specifications *A_SPECIFICATIONS)
	CheckoutA_SPECIFICATIONS(a_specifications *A_SPECIFICATIONS)
	CommitA_SPECIFICATION_TYPE_REF(a_specification_type_ref *A_SPECIFICATION_TYPE_REF)
	CheckoutA_SPECIFICATION_TYPE_REF(a_specification_type_ref *A_SPECIFICATION_TYPE_REF)
	CommitA_SPECIFIED_VALUES(a_specified_values *A_SPECIFIED_VALUES)
	CheckoutA_SPECIFIED_VALUES(a_specified_values *A_SPECIFIED_VALUES)
	CommitA_SPEC_ATTRIBUTES(a_spec_attributes *A_SPEC_ATTRIBUTES)
	CheckoutA_SPEC_ATTRIBUTES(a_spec_attributes *A_SPEC_ATTRIBUTES)
	CommitA_SPEC_OBJECTS(a_spec_objects *A_SPEC_OBJECTS)
	CheckoutA_SPEC_OBJECTS(a_spec_objects *A_SPEC_OBJECTS)
	CommitA_SPEC_OBJECT_TYPE_REF(a_spec_object_type_ref *A_SPEC_OBJECT_TYPE_REF)
	CheckoutA_SPEC_OBJECT_TYPE_REF(a_spec_object_type_ref *A_SPEC_OBJECT_TYPE_REF)
	CommitA_SPEC_RELATIONS(a_spec_relations *A_SPEC_RELATIONS)
	CheckoutA_SPEC_RELATIONS(a_spec_relations *A_SPEC_RELATIONS)
	CommitA_SPEC_RELATION_GROUPS(a_spec_relation_groups *A_SPEC_RELATION_GROUPS)
	CheckoutA_SPEC_RELATION_GROUPS(a_spec_relation_groups *A_SPEC_RELATION_GROUPS)
	CommitA_SPEC_RELATION_REF(a_spec_relation_ref *A_SPEC_RELATION_REF)
	CheckoutA_SPEC_RELATION_REF(a_spec_relation_ref *A_SPEC_RELATION_REF)
	CommitA_SPEC_RELATION_TYPE_REF(a_spec_relation_type_ref *A_SPEC_RELATION_TYPE_REF)
	CheckoutA_SPEC_RELATION_TYPE_REF(a_spec_relation_type_ref *A_SPEC_RELATION_TYPE_REF)
	CommitA_SPEC_TYPES(a_spec_types *A_SPEC_TYPES)
	CheckoutA_SPEC_TYPES(a_spec_types *A_SPEC_TYPES)
	CommitA_TARGET_SPECIFICATION_1(a_target_specification_1 *A_TARGET_SPECIFICATION_1)
	CheckoutA_TARGET_SPECIFICATION_1(a_target_specification_1 *A_TARGET_SPECIFICATION_1)
	CommitA_THE_HEADER(a_the_header *A_THE_HEADER)
	CheckoutA_THE_HEADER(a_the_header *A_THE_HEADER)
	CommitA_TOOL_EXTENSIONS(a_tool_extensions *A_TOOL_EXTENSIONS)
	CheckoutA_TOOL_EXTENSIONS(a_tool_extensions *A_TOOL_EXTENSIONS)
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

		A_ALTERNATIVE_IDs:           make(map[*A_ALTERNATIVE_ID]any),
		A_ALTERNATIVE_IDs_mapString: make(map[string]*A_ALTERNATIVE_ID),

		A_ATTRIBUTE_DEFINITION_BOOLEAN_REFs:           make(map[*A_ATTRIBUTE_DEFINITION_BOOLEAN_REF]any),
		A_ATTRIBUTE_DEFINITION_BOOLEAN_REFs_mapString: make(map[string]*A_ATTRIBUTE_DEFINITION_BOOLEAN_REF),

		A_ATTRIBUTE_DEFINITION_DATE_REFs:           make(map[*A_ATTRIBUTE_DEFINITION_DATE_REF]any),
		A_ATTRIBUTE_DEFINITION_DATE_REFs_mapString: make(map[string]*A_ATTRIBUTE_DEFINITION_DATE_REF),

		A_ATTRIBUTE_DEFINITION_ENUMERATION_REFs:           make(map[*A_ATTRIBUTE_DEFINITION_ENUMERATION_REF]any),
		A_ATTRIBUTE_DEFINITION_ENUMERATION_REFs_mapString: make(map[string]*A_ATTRIBUTE_DEFINITION_ENUMERATION_REF),

		A_ATTRIBUTE_DEFINITION_INTEGER_REFs:           make(map[*A_ATTRIBUTE_DEFINITION_INTEGER_REF]any),
		A_ATTRIBUTE_DEFINITION_INTEGER_REFs_mapString: make(map[string]*A_ATTRIBUTE_DEFINITION_INTEGER_REF),

		A_ATTRIBUTE_DEFINITION_REAL_REFs:           make(map[*A_ATTRIBUTE_DEFINITION_REAL_REF]any),
		A_ATTRIBUTE_DEFINITION_REAL_REFs_mapString: make(map[string]*A_ATTRIBUTE_DEFINITION_REAL_REF),

		A_ATTRIBUTE_DEFINITION_STRING_REFs:           make(map[*A_ATTRIBUTE_DEFINITION_STRING_REF]any),
		A_ATTRIBUTE_DEFINITION_STRING_REFs_mapString: make(map[string]*A_ATTRIBUTE_DEFINITION_STRING_REF),

		A_ATTRIBUTE_DEFINITION_XHTML_REFs:           make(map[*A_ATTRIBUTE_DEFINITION_XHTML_REF]any),
		A_ATTRIBUTE_DEFINITION_XHTML_REFs_mapString: make(map[string]*A_ATTRIBUTE_DEFINITION_XHTML_REF),

		A_ATTRIBUTE_VALUE_BOOLEANs:           make(map[*A_ATTRIBUTE_VALUE_BOOLEAN]any),
		A_ATTRIBUTE_VALUE_BOOLEANs_mapString: make(map[string]*A_ATTRIBUTE_VALUE_BOOLEAN),

		A_ATTRIBUTE_VALUE_DATEs:           make(map[*A_ATTRIBUTE_VALUE_DATE]any),
		A_ATTRIBUTE_VALUE_DATEs_mapString: make(map[string]*A_ATTRIBUTE_VALUE_DATE),

		A_ATTRIBUTE_VALUE_ENUMERATIONs:           make(map[*A_ATTRIBUTE_VALUE_ENUMERATION]any),
		A_ATTRIBUTE_VALUE_ENUMERATIONs_mapString: make(map[string]*A_ATTRIBUTE_VALUE_ENUMERATION),

		A_ATTRIBUTE_VALUE_INTEGERs:           make(map[*A_ATTRIBUTE_VALUE_INTEGER]any),
		A_ATTRIBUTE_VALUE_INTEGERs_mapString: make(map[string]*A_ATTRIBUTE_VALUE_INTEGER),

		A_ATTRIBUTE_VALUE_REALs:           make(map[*A_ATTRIBUTE_VALUE_REAL]any),
		A_ATTRIBUTE_VALUE_REALs_mapString: make(map[string]*A_ATTRIBUTE_VALUE_REAL),

		A_ATTRIBUTE_VALUE_STRINGs:           make(map[*A_ATTRIBUTE_VALUE_STRING]any),
		A_ATTRIBUTE_VALUE_STRINGs_mapString: make(map[string]*A_ATTRIBUTE_VALUE_STRING),

		A_ATTRIBUTE_VALUE_XHTMLs:           make(map[*A_ATTRIBUTE_VALUE_XHTML]any),
		A_ATTRIBUTE_VALUE_XHTMLs_mapString: make(map[string]*A_ATTRIBUTE_VALUE_XHTML),

		A_ATTRIBUTE_VALUE_XHTML_1s:           make(map[*A_ATTRIBUTE_VALUE_XHTML_1]any),
		A_ATTRIBUTE_VALUE_XHTML_1s_mapString: make(map[string]*A_ATTRIBUTE_VALUE_XHTML_1),

		A_CHILDRENs:           make(map[*A_CHILDREN]any),
		A_CHILDRENs_mapString: make(map[string]*A_CHILDREN),

		A_CORE_CONTENTs:           make(map[*A_CORE_CONTENT]any),
		A_CORE_CONTENTs_mapString: make(map[string]*A_CORE_CONTENT),

		A_DATATYPESs:           make(map[*A_DATATYPES]any),
		A_DATATYPESs_mapString: make(map[string]*A_DATATYPES),

		A_DATATYPE_DEFINITION_BOOLEAN_REFs:           make(map[*A_DATATYPE_DEFINITION_BOOLEAN_REF]any),
		A_DATATYPE_DEFINITION_BOOLEAN_REFs_mapString: make(map[string]*A_DATATYPE_DEFINITION_BOOLEAN_REF),

		A_DATATYPE_DEFINITION_DATE_REFs:           make(map[*A_DATATYPE_DEFINITION_DATE_REF]any),
		A_DATATYPE_DEFINITION_DATE_REFs_mapString: make(map[string]*A_DATATYPE_DEFINITION_DATE_REF),

		A_DATATYPE_DEFINITION_ENUMERATION_REFs:           make(map[*A_DATATYPE_DEFINITION_ENUMERATION_REF]any),
		A_DATATYPE_DEFINITION_ENUMERATION_REFs_mapString: make(map[string]*A_DATATYPE_DEFINITION_ENUMERATION_REF),

		A_DATATYPE_DEFINITION_INTEGER_REFs:           make(map[*A_DATATYPE_DEFINITION_INTEGER_REF]any),
		A_DATATYPE_DEFINITION_INTEGER_REFs_mapString: make(map[string]*A_DATATYPE_DEFINITION_INTEGER_REF),

		A_DATATYPE_DEFINITION_REAL_REFs:           make(map[*A_DATATYPE_DEFINITION_REAL_REF]any),
		A_DATATYPE_DEFINITION_REAL_REFs_mapString: make(map[string]*A_DATATYPE_DEFINITION_REAL_REF),

		A_DATATYPE_DEFINITION_STRING_REFs:           make(map[*A_DATATYPE_DEFINITION_STRING_REF]any),
		A_DATATYPE_DEFINITION_STRING_REFs_mapString: make(map[string]*A_DATATYPE_DEFINITION_STRING_REF),

		A_DATATYPE_DEFINITION_XHTML_REFs:           make(map[*A_DATATYPE_DEFINITION_XHTML_REF]any),
		A_DATATYPE_DEFINITION_XHTML_REFs_mapString: make(map[string]*A_DATATYPE_DEFINITION_XHTML_REF),

		A_EDITABLE_ATTSs:           make(map[*A_EDITABLE_ATTS]any),
		A_EDITABLE_ATTSs_mapString: make(map[string]*A_EDITABLE_ATTS),

		A_ENUM_VALUE_REFs:           make(map[*A_ENUM_VALUE_REF]any),
		A_ENUM_VALUE_REFs_mapString: make(map[string]*A_ENUM_VALUE_REF),

		A_OBJECTs:           make(map[*A_OBJECT]any),
		A_OBJECTs_mapString: make(map[string]*A_OBJECT),

		A_PROPERTIESs:           make(map[*A_PROPERTIES]any),
		A_PROPERTIESs_mapString: make(map[string]*A_PROPERTIES),

		A_RELATION_GROUP_TYPE_REFs:           make(map[*A_RELATION_GROUP_TYPE_REF]any),
		A_RELATION_GROUP_TYPE_REFs_mapString: make(map[string]*A_RELATION_GROUP_TYPE_REF),

		A_SOURCE_1s:           make(map[*A_SOURCE_1]any),
		A_SOURCE_1s_mapString: make(map[string]*A_SOURCE_1),

		A_SPECIFICATIONSs:           make(map[*A_SPECIFICATIONS]any),
		A_SPECIFICATIONSs_mapString: make(map[string]*A_SPECIFICATIONS),

		A_SPECIFICATION_TYPE_REFs:           make(map[*A_SPECIFICATION_TYPE_REF]any),
		A_SPECIFICATION_TYPE_REFs_mapString: make(map[string]*A_SPECIFICATION_TYPE_REF),

		A_SPECIFIED_VALUESs:           make(map[*A_SPECIFIED_VALUES]any),
		A_SPECIFIED_VALUESs_mapString: make(map[string]*A_SPECIFIED_VALUES),

		A_SPEC_ATTRIBUTESs:           make(map[*A_SPEC_ATTRIBUTES]any),
		A_SPEC_ATTRIBUTESs_mapString: make(map[string]*A_SPEC_ATTRIBUTES),

		A_SPEC_OBJECTSs:           make(map[*A_SPEC_OBJECTS]any),
		A_SPEC_OBJECTSs_mapString: make(map[string]*A_SPEC_OBJECTS),

		A_SPEC_OBJECT_TYPE_REFs:           make(map[*A_SPEC_OBJECT_TYPE_REF]any),
		A_SPEC_OBJECT_TYPE_REFs_mapString: make(map[string]*A_SPEC_OBJECT_TYPE_REF),

		A_SPEC_RELATIONSs:           make(map[*A_SPEC_RELATIONS]any),
		A_SPEC_RELATIONSs_mapString: make(map[string]*A_SPEC_RELATIONS),

		A_SPEC_RELATION_GROUPSs:           make(map[*A_SPEC_RELATION_GROUPS]any),
		A_SPEC_RELATION_GROUPSs_mapString: make(map[string]*A_SPEC_RELATION_GROUPS),

		A_SPEC_RELATION_REFs:           make(map[*A_SPEC_RELATION_REF]any),
		A_SPEC_RELATION_REFs_mapString: make(map[string]*A_SPEC_RELATION_REF),

		A_SPEC_RELATION_TYPE_REFs:           make(map[*A_SPEC_RELATION_TYPE_REF]any),
		A_SPEC_RELATION_TYPE_REFs_mapString: make(map[string]*A_SPEC_RELATION_TYPE_REF),

		A_SPEC_TYPESs:           make(map[*A_SPEC_TYPES]any),
		A_SPEC_TYPESs_mapString: make(map[string]*A_SPEC_TYPES),

		A_TARGET_SPECIFICATION_1s:           make(map[*A_TARGET_SPECIFICATION_1]any),
		A_TARGET_SPECIFICATION_1s_mapString: make(map[string]*A_TARGET_SPECIFICATION_1),

		A_THE_HEADERs:           make(map[*A_THE_HEADER]any),
		A_THE_HEADERs_mapString: make(map[string]*A_THE_HEADER),

		A_TOOL_EXTENSIONSs:           make(map[*A_TOOL_EXTENSIONS]any),
		A_TOOL_EXTENSIONSs_mapString: make(map[string]*A_TOOL_EXTENSIONS),

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
	stage.Map_GongStructName_InstancesNb["A_ALTERNATIVE_ID"] = len(stage.A_ALTERNATIVE_IDs)
	stage.Map_GongStructName_InstancesNb["A_ATTRIBUTE_DEFINITION_BOOLEAN_REF"] = len(stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFs)
	stage.Map_GongStructName_InstancesNb["A_ATTRIBUTE_DEFINITION_DATE_REF"] = len(stage.A_ATTRIBUTE_DEFINITION_DATE_REFs)
	stage.Map_GongStructName_InstancesNb["A_ATTRIBUTE_DEFINITION_ENUMERATION_REF"] = len(stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFs)
	stage.Map_GongStructName_InstancesNb["A_ATTRIBUTE_DEFINITION_INTEGER_REF"] = len(stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFs)
	stage.Map_GongStructName_InstancesNb["A_ATTRIBUTE_DEFINITION_REAL_REF"] = len(stage.A_ATTRIBUTE_DEFINITION_REAL_REFs)
	stage.Map_GongStructName_InstancesNb["A_ATTRIBUTE_DEFINITION_STRING_REF"] = len(stage.A_ATTRIBUTE_DEFINITION_STRING_REFs)
	stage.Map_GongStructName_InstancesNb["A_ATTRIBUTE_DEFINITION_XHTML_REF"] = len(stage.A_ATTRIBUTE_DEFINITION_XHTML_REFs)
	stage.Map_GongStructName_InstancesNb["A_ATTRIBUTE_VALUE_BOOLEAN"] = len(stage.A_ATTRIBUTE_VALUE_BOOLEANs)
	stage.Map_GongStructName_InstancesNb["A_ATTRIBUTE_VALUE_DATE"] = len(stage.A_ATTRIBUTE_VALUE_DATEs)
	stage.Map_GongStructName_InstancesNb["A_ATTRIBUTE_VALUE_ENUMERATION"] = len(stage.A_ATTRIBUTE_VALUE_ENUMERATIONs)
	stage.Map_GongStructName_InstancesNb["A_ATTRIBUTE_VALUE_INTEGER"] = len(stage.A_ATTRIBUTE_VALUE_INTEGERs)
	stage.Map_GongStructName_InstancesNb["A_ATTRIBUTE_VALUE_REAL"] = len(stage.A_ATTRIBUTE_VALUE_REALs)
	stage.Map_GongStructName_InstancesNb["A_ATTRIBUTE_VALUE_STRING"] = len(stage.A_ATTRIBUTE_VALUE_STRINGs)
	stage.Map_GongStructName_InstancesNb["A_ATTRIBUTE_VALUE_XHTML"] = len(stage.A_ATTRIBUTE_VALUE_XHTMLs)
	stage.Map_GongStructName_InstancesNb["A_ATTRIBUTE_VALUE_XHTML_1"] = len(stage.A_ATTRIBUTE_VALUE_XHTML_1s)
	stage.Map_GongStructName_InstancesNb["A_CHILDREN"] = len(stage.A_CHILDRENs)
	stage.Map_GongStructName_InstancesNb["A_CORE_CONTENT"] = len(stage.A_CORE_CONTENTs)
	stage.Map_GongStructName_InstancesNb["A_DATATYPES"] = len(stage.A_DATATYPESs)
	stage.Map_GongStructName_InstancesNb["A_DATATYPE_DEFINITION_BOOLEAN_REF"] = len(stage.A_DATATYPE_DEFINITION_BOOLEAN_REFs)
	stage.Map_GongStructName_InstancesNb["A_DATATYPE_DEFINITION_DATE_REF"] = len(stage.A_DATATYPE_DEFINITION_DATE_REFs)
	stage.Map_GongStructName_InstancesNb["A_DATATYPE_DEFINITION_ENUMERATION_REF"] = len(stage.A_DATATYPE_DEFINITION_ENUMERATION_REFs)
	stage.Map_GongStructName_InstancesNb["A_DATATYPE_DEFINITION_INTEGER_REF"] = len(stage.A_DATATYPE_DEFINITION_INTEGER_REFs)
	stage.Map_GongStructName_InstancesNb["A_DATATYPE_DEFINITION_REAL_REF"] = len(stage.A_DATATYPE_DEFINITION_REAL_REFs)
	stage.Map_GongStructName_InstancesNb["A_DATATYPE_DEFINITION_STRING_REF"] = len(stage.A_DATATYPE_DEFINITION_STRING_REFs)
	stage.Map_GongStructName_InstancesNb["A_DATATYPE_DEFINITION_XHTML_REF"] = len(stage.A_DATATYPE_DEFINITION_XHTML_REFs)
	stage.Map_GongStructName_InstancesNb["A_EDITABLE_ATTS"] = len(stage.A_EDITABLE_ATTSs)
	stage.Map_GongStructName_InstancesNb["A_ENUM_VALUE_REF"] = len(stage.A_ENUM_VALUE_REFs)
	stage.Map_GongStructName_InstancesNb["A_OBJECT"] = len(stage.A_OBJECTs)
	stage.Map_GongStructName_InstancesNb["A_PROPERTIES"] = len(stage.A_PROPERTIESs)
	stage.Map_GongStructName_InstancesNb["A_RELATION_GROUP_TYPE_REF"] = len(stage.A_RELATION_GROUP_TYPE_REFs)
	stage.Map_GongStructName_InstancesNb["A_SOURCE_1"] = len(stage.A_SOURCE_1s)
	stage.Map_GongStructName_InstancesNb["A_SPECIFICATIONS"] = len(stage.A_SPECIFICATIONSs)
	stage.Map_GongStructName_InstancesNb["A_SPECIFICATION_TYPE_REF"] = len(stage.A_SPECIFICATION_TYPE_REFs)
	stage.Map_GongStructName_InstancesNb["A_SPECIFIED_VALUES"] = len(stage.A_SPECIFIED_VALUESs)
	stage.Map_GongStructName_InstancesNb["A_SPEC_ATTRIBUTES"] = len(stage.A_SPEC_ATTRIBUTESs)
	stage.Map_GongStructName_InstancesNb["A_SPEC_OBJECTS"] = len(stage.A_SPEC_OBJECTSs)
	stage.Map_GongStructName_InstancesNb["A_SPEC_OBJECT_TYPE_REF"] = len(stage.A_SPEC_OBJECT_TYPE_REFs)
	stage.Map_GongStructName_InstancesNb["A_SPEC_RELATIONS"] = len(stage.A_SPEC_RELATIONSs)
	stage.Map_GongStructName_InstancesNb["A_SPEC_RELATION_GROUPS"] = len(stage.A_SPEC_RELATION_GROUPSs)
	stage.Map_GongStructName_InstancesNb["A_SPEC_RELATION_REF"] = len(stage.A_SPEC_RELATION_REFs)
	stage.Map_GongStructName_InstancesNb["A_SPEC_RELATION_TYPE_REF"] = len(stage.A_SPEC_RELATION_TYPE_REFs)
	stage.Map_GongStructName_InstancesNb["A_SPEC_TYPES"] = len(stage.A_SPEC_TYPESs)
	stage.Map_GongStructName_InstancesNb["A_TARGET_SPECIFICATION_1"] = len(stage.A_TARGET_SPECIFICATION_1s)
	stage.Map_GongStructName_InstancesNb["A_THE_HEADER"] = len(stage.A_THE_HEADERs)
	stage.Map_GongStructName_InstancesNb["A_TOOL_EXTENSIONS"] = len(stage.A_TOOL_EXTENSIONSs)
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
	stage.Map_GongStructName_InstancesNb["A_ALTERNATIVE_ID"] = len(stage.A_ALTERNATIVE_IDs)
	stage.Map_GongStructName_InstancesNb["A_ATTRIBUTE_DEFINITION_BOOLEAN_REF"] = len(stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFs)
	stage.Map_GongStructName_InstancesNb["A_ATTRIBUTE_DEFINITION_DATE_REF"] = len(stage.A_ATTRIBUTE_DEFINITION_DATE_REFs)
	stage.Map_GongStructName_InstancesNb["A_ATTRIBUTE_DEFINITION_ENUMERATION_REF"] = len(stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFs)
	stage.Map_GongStructName_InstancesNb["A_ATTRIBUTE_DEFINITION_INTEGER_REF"] = len(stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFs)
	stage.Map_GongStructName_InstancesNb["A_ATTRIBUTE_DEFINITION_REAL_REF"] = len(stage.A_ATTRIBUTE_DEFINITION_REAL_REFs)
	stage.Map_GongStructName_InstancesNb["A_ATTRIBUTE_DEFINITION_STRING_REF"] = len(stage.A_ATTRIBUTE_DEFINITION_STRING_REFs)
	stage.Map_GongStructName_InstancesNb["A_ATTRIBUTE_DEFINITION_XHTML_REF"] = len(stage.A_ATTRIBUTE_DEFINITION_XHTML_REFs)
	stage.Map_GongStructName_InstancesNb["A_ATTRIBUTE_VALUE_BOOLEAN"] = len(stage.A_ATTRIBUTE_VALUE_BOOLEANs)
	stage.Map_GongStructName_InstancesNb["A_ATTRIBUTE_VALUE_DATE"] = len(stage.A_ATTRIBUTE_VALUE_DATEs)
	stage.Map_GongStructName_InstancesNb["A_ATTRIBUTE_VALUE_ENUMERATION"] = len(stage.A_ATTRIBUTE_VALUE_ENUMERATIONs)
	stage.Map_GongStructName_InstancesNb["A_ATTRIBUTE_VALUE_INTEGER"] = len(stage.A_ATTRIBUTE_VALUE_INTEGERs)
	stage.Map_GongStructName_InstancesNb["A_ATTRIBUTE_VALUE_REAL"] = len(stage.A_ATTRIBUTE_VALUE_REALs)
	stage.Map_GongStructName_InstancesNb["A_ATTRIBUTE_VALUE_STRING"] = len(stage.A_ATTRIBUTE_VALUE_STRINGs)
	stage.Map_GongStructName_InstancesNb["A_ATTRIBUTE_VALUE_XHTML"] = len(stage.A_ATTRIBUTE_VALUE_XHTMLs)
	stage.Map_GongStructName_InstancesNb["A_ATTRIBUTE_VALUE_XHTML_1"] = len(stage.A_ATTRIBUTE_VALUE_XHTML_1s)
	stage.Map_GongStructName_InstancesNb["A_CHILDREN"] = len(stage.A_CHILDRENs)
	stage.Map_GongStructName_InstancesNb["A_CORE_CONTENT"] = len(stage.A_CORE_CONTENTs)
	stage.Map_GongStructName_InstancesNb["A_DATATYPES"] = len(stage.A_DATATYPESs)
	stage.Map_GongStructName_InstancesNb["A_DATATYPE_DEFINITION_BOOLEAN_REF"] = len(stage.A_DATATYPE_DEFINITION_BOOLEAN_REFs)
	stage.Map_GongStructName_InstancesNb["A_DATATYPE_DEFINITION_DATE_REF"] = len(stage.A_DATATYPE_DEFINITION_DATE_REFs)
	stage.Map_GongStructName_InstancesNb["A_DATATYPE_DEFINITION_ENUMERATION_REF"] = len(stage.A_DATATYPE_DEFINITION_ENUMERATION_REFs)
	stage.Map_GongStructName_InstancesNb["A_DATATYPE_DEFINITION_INTEGER_REF"] = len(stage.A_DATATYPE_DEFINITION_INTEGER_REFs)
	stage.Map_GongStructName_InstancesNb["A_DATATYPE_DEFINITION_REAL_REF"] = len(stage.A_DATATYPE_DEFINITION_REAL_REFs)
	stage.Map_GongStructName_InstancesNb["A_DATATYPE_DEFINITION_STRING_REF"] = len(stage.A_DATATYPE_DEFINITION_STRING_REFs)
	stage.Map_GongStructName_InstancesNb["A_DATATYPE_DEFINITION_XHTML_REF"] = len(stage.A_DATATYPE_DEFINITION_XHTML_REFs)
	stage.Map_GongStructName_InstancesNb["A_EDITABLE_ATTS"] = len(stage.A_EDITABLE_ATTSs)
	stage.Map_GongStructName_InstancesNb["A_ENUM_VALUE_REF"] = len(stage.A_ENUM_VALUE_REFs)
	stage.Map_GongStructName_InstancesNb["A_OBJECT"] = len(stage.A_OBJECTs)
	stage.Map_GongStructName_InstancesNb["A_PROPERTIES"] = len(stage.A_PROPERTIESs)
	stage.Map_GongStructName_InstancesNb["A_RELATION_GROUP_TYPE_REF"] = len(stage.A_RELATION_GROUP_TYPE_REFs)
	stage.Map_GongStructName_InstancesNb["A_SOURCE_1"] = len(stage.A_SOURCE_1s)
	stage.Map_GongStructName_InstancesNb["A_SPECIFICATIONS"] = len(stage.A_SPECIFICATIONSs)
	stage.Map_GongStructName_InstancesNb["A_SPECIFICATION_TYPE_REF"] = len(stage.A_SPECIFICATION_TYPE_REFs)
	stage.Map_GongStructName_InstancesNb["A_SPECIFIED_VALUES"] = len(stage.A_SPECIFIED_VALUESs)
	stage.Map_GongStructName_InstancesNb["A_SPEC_ATTRIBUTES"] = len(stage.A_SPEC_ATTRIBUTESs)
	stage.Map_GongStructName_InstancesNb["A_SPEC_OBJECTS"] = len(stage.A_SPEC_OBJECTSs)
	stage.Map_GongStructName_InstancesNb["A_SPEC_OBJECT_TYPE_REF"] = len(stage.A_SPEC_OBJECT_TYPE_REFs)
	stage.Map_GongStructName_InstancesNb["A_SPEC_RELATIONS"] = len(stage.A_SPEC_RELATIONSs)
	stage.Map_GongStructName_InstancesNb["A_SPEC_RELATION_GROUPS"] = len(stage.A_SPEC_RELATION_GROUPSs)
	stage.Map_GongStructName_InstancesNb["A_SPEC_RELATION_REF"] = len(stage.A_SPEC_RELATION_REFs)
	stage.Map_GongStructName_InstancesNb["A_SPEC_RELATION_TYPE_REF"] = len(stage.A_SPEC_RELATION_TYPE_REFs)
	stage.Map_GongStructName_InstancesNb["A_SPEC_TYPES"] = len(stage.A_SPEC_TYPESs)
	stage.Map_GongStructName_InstancesNb["A_TARGET_SPECIFICATION_1"] = len(stage.A_TARGET_SPECIFICATION_1s)
	stage.Map_GongStructName_InstancesNb["A_THE_HEADER"] = len(stage.A_THE_HEADERs)
	stage.Map_GongStructName_InstancesNb["A_TOOL_EXTENSIONS"] = len(stage.A_TOOL_EXTENSIONSs)
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

// Stage puts a_alternative_id to the model stage
func (a_alternative_id *A_ALTERNATIVE_ID) Stage(stage *StageStruct) *A_ALTERNATIVE_ID {
	stage.A_ALTERNATIVE_IDs[a_alternative_id] = __member
	stage.A_ALTERNATIVE_IDs_mapString[a_alternative_id.Name] = a_alternative_id

	return a_alternative_id
}

// Unstage removes a_alternative_id off the model stage
func (a_alternative_id *A_ALTERNATIVE_ID) Unstage(stage *StageStruct) *A_ALTERNATIVE_ID {
	delete(stage.A_ALTERNATIVE_IDs, a_alternative_id)
	delete(stage.A_ALTERNATIVE_IDs_mapString, a_alternative_id.Name)
	return a_alternative_id
}

// UnstageVoid removes a_alternative_id off the model stage
func (a_alternative_id *A_ALTERNATIVE_ID) UnstageVoid(stage *StageStruct) {
	delete(stage.A_ALTERNATIVE_IDs, a_alternative_id)
	delete(stage.A_ALTERNATIVE_IDs_mapString, a_alternative_id.Name)
}

// commit a_alternative_id to the back repo (if it is already staged)
func (a_alternative_id *A_ALTERNATIVE_ID) Commit(stage *StageStruct) *A_ALTERNATIVE_ID {
	if _, ok := stage.A_ALTERNATIVE_IDs[a_alternative_id]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitA_ALTERNATIVE_ID(a_alternative_id)
		}
	}
	return a_alternative_id
}

func (a_alternative_id *A_ALTERNATIVE_ID) CommitVoid(stage *StageStruct) {
	a_alternative_id.Commit(stage)
}

// Checkout a_alternative_id to the back repo (if it is already staged)
func (a_alternative_id *A_ALTERNATIVE_ID) Checkout(stage *StageStruct) *A_ALTERNATIVE_ID {
	if _, ok := stage.A_ALTERNATIVE_IDs[a_alternative_id]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutA_ALTERNATIVE_ID(a_alternative_id)
		}
	}
	return a_alternative_id
}

// for satisfaction of GongStruct interface
func (a_alternative_id *A_ALTERNATIVE_ID) GetName() (res string) {
	return a_alternative_id.Name
}

// Stage puts a_attribute_definition_boolean_ref to the model stage
func (a_attribute_definition_boolean_ref *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF) Stage(stage *StageStruct) *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF {
	stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFs[a_attribute_definition_boolean_ref] = __member
	stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFs_mapString[a_attribute_definition_boolean_ref.Name] = a_attribute_definition_boolean_ref

	return a_attribute_definition_boolean_ref
}

// Unstage removes a_attribute_definition_boolean_ref off the model stage
func (a_attribute_definition_boolean_ref *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF) Unstage(stage *StageStruct) *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF {
	delete(stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFs, a_attribute_definition_boolean_ref)
	delete(stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFs_mapString, a_attribute_definition_boolean_ref.Name)
	return a_attribute_definition_boolean_ref
}

// UnstageVoid removes a_attribute_definition_boolean_ref off the model stage
func (a_attribute_definition_boolean_ref *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF) UnstageVoid(stage *StageStruct) {
	delete(stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFs, a_attribute_definition_boolean_ref)
	delete(stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFs_mapString, a_attribute_definition_boolean_ref.Name)
}

// commit a_attribute_definition_boolean_ref to the back repo (if it is already staged)
func (a_attribute_definition_boolean_ref *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF) Commit(stage *StageStruct) *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF {
	if _, ok := stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFs[a_attribute_definition_boolean_ref]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitA_ATTRIBUTE_DEFINITION_BOOLEAN_REF(a_attribute_definition_boolean_ref)
		}
	}
	return a_attribute_definition_boolean_ref
}

func (a_attribute_definition_boolean_ref *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF) CommitVoid(stage *StageStruct) {
	a_attribute_definition_boolean_ref.Commit(stage)
}

// Checkout a_attribute_definition_boolean_ref to the back repo (if it is already staged)
func (a_attribute_definition_boolean_ref *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF) Checkout(stage *StageStruct) *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF {
	if _, ok := stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFs[a_attribute_definition_boolean_ref]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutA_ATTRIBUTE_DEFINITION_BOOLEAN_REF(a_attribute_definition_boolean_ref)
		}
	}
	return a_attribute_definition_boolean_ref
}

// for satisfaction of GongStruct interface
func (a_attribute_definition_boolean_ref *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF) GetName() (res string) {
	return a_attribute_definition_boolean_ref.Name
}

// Stage puts a_attribute_definition_date_ref to the model stage
func (a_attribute_definition_date_ref *A_ATTRIBUTE_DEFINITION_DATE_REF) Stage(stage *StageStruct) *A_ATTRIBUTE_DEFINITION_DATE_REF {
	stage.A_ATTRIBUTE_DEFINITION_DATE_REFs[a_attribute_definition_date_ref] = __member
	stage.A_ATTRIBUTE_DEFINITION_DATE_REFs_mapString[a_attribute_definition_date_ref.Name] = a_attribute_definition_date_ref

	return a_attribute_definition_date_ref
}

// Unstage removes a_attribute_definition_date_ref off the model stage
func (a_attribute_definition_date_ref *A_ATTRIBUTE_DEFINITION_DATE_REF) Unstage(stage *StageStruct) *A_ATTRIBUTE_DEFINITION_DATE_REF {
	delete(stage.A_ATTRIBUTE_DEFINITION_DATE_REFs, a_attribute_definition_date_ref)
	delete(stage.A_ATTRIBUTE_DEFINITION_DATE_REFs_mapString, a_attribute_definition_date_ref.Name)
	return a_attribute_definition_date_ref
}

// UnstageVoid removes a_attribute_definition_date_ref off the model stage
func (a_attribute_definition_date_ref *A_ATTRIBUTE_DEFINITION_DATE_REF) UnstageVoid(stage *StageStruct) {
	delete(stage.A_ATTRIBUTE_DEFINITION_DATE_REFs, a_attribute_definition_date_ref)
	delete(stage.A_ATTRIBUTE_DEFINITION_DATE_REFs_mapString, a_attribute_definition_date_ref.Name)
}

// commit a_attribute_definition_date_ref to the back repo (if it is already staged)
func (a_attribute_definition_date_ref *A_ATTRIBUTE_DEFINITION_DATE_REF) Commit(stage *StageStruct) *A_ATTRIBUTE_DEFINITION_DATE_REF {
	if _, ok := stage.A_ATTRIBUTE_DEFINITION_DATE_REFs[a_attribute_definition_date_ref]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitA_ATTRIBUTE_DEFINITION_DATE_REF(a_attribute_definition_date_ref)
		}
	}
	return a_attribute_definition_date_ref
}

func (a_attribute_definition_date_ref *A_ATTRIBUTE_DEFINITION_DATE_REF) CommitVoid(stage *StageStruct) {
	a_attribute_definition_date_ref.Commit(stage)
}

// Checkout a_attribute_definition_date_ref to the back repo (if it is already staged)
func (a_attribute_definition_date_ref *A_ATTRIBUTE_DEFINITION_DATE_REF) Checkout(stage *StageStruct) *A_ATTRIBUTE_DEFINITION_DATE_REF {
	if _, ok := stage.A_ATTRIBUTE_DEFINITION_DATE_REFs[a_attribute_definition_date_ref]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutA_ATTRIBUTE_DEFINITION_DATE_REF(a_attribute_definition_date_ref)
		}
	}
	return a_attribute_definition_date_ref
}

// for satisfaction of GongStruct interface
func (a_attribute_definition_date_ref *A_ATTRIBUTE_DEFINITION_DATE_REF) GetName() (res string) {
	return a_attribute_definition_date_ref.Name
}

// Stage puts a_attribute_definition_enumeration_ref to the model stage
func (a_attribute_definition_enumeration_ref *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF) Stage(stage *StageStruct) *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF {
	stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFs[a_attribute_definition_enumeration_ref] = __member
	stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFs_mapString[a_attribute_definition_enumeration_ref.Name] = a_attribute_definition_enumeration_ref

	return a_attribute_definition_enumeration_ref
}

// Unstage removes a_attribute_definition_enumeration_ref off the model stage
func (a_attribute_definition_enumeration_ref *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF) Unstage(stage *StageStruct) *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF {
	delete(stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFs, a_attribute_definition_enumeration_ref)
	delete(stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFs_mapString, a_attribute_definition_enumeration_ref.Name)
	return a_attribute_definition_enumeration_ref
}

// UnstageVoid removes a_attribute_definition_enumeration_ref off the model stage
func (a_attribute_definition_enumeration_ref *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF) UnstageVoid(stage *StageStruct) {
	delete(stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFs, a_attribute_definition_enumeration_ref)
	delete(stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFs_mapString, a_attribute_definition_enumeration_ref.Name)
}

// commit a_attribute_definition_enumeration_ref to the back repo (if it is already staged)
func (a_attribute_definition_enumeration_ref *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF) Commit(stage *StageStruct) *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF {
	if _, ok := stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFs[a_attribute_definition_enumeration_ref]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitA_ATTRIBUTE_DEFINITION_ENUMERATION_REF(a_attribute_definition_enumeration_ref)
		}
	}
	return a_attribute_definition_enumeration_ref
}

func (a_attribute_definition_enumeration_ref *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF) CommitVoid(stage *StageStruct) {
	a_attribute_definition_enumeration_ref.Commit(stage)
}

// Checkout a_attribute_definition_enumeration_ref to the back repo (if it is already staged)
func (a_attribute_definition_enumeration_ref *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF) Checkout(stage *StageStruct) *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF {
	if _, ok := stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFs[a_attribute_definition_enumeration_ref]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutA_ATTRIBUTE_DEFINITION_ENUMERATION_REF(a_attribute_definition_enumeration_ref)
		}
	}
	return a_attribute_definition_enumeration_ref
}

// for satisfaction of GongStruct interface
func (a_attribute_definition_enumeration_ref *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF) GetName() (res string) {
	return a_attribute_definition_enumeration_ref.Name
}

// Stage puts a_attribute_definition_integer_ref to the model stage
func (a_attribute_definition_integer_ref *A_ATTRIBUTE_DEFINITION_INTEGER_REF) Stage(stage *StageStruct) *A_ATTRIBUTE_DEFINITION_INTEGER_REF {
	stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFs[a_attribute_definition_integer_ref] = __member
	stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFs_mapString[a_attribute_definition_integer_ref.Name] = a_attribute_definition_integer_ref

	return a_attribute_definition_integer_ref
}

// Unstage removes a_attribute_definition_integer_ref off the model stage
func (a_attribute_definition_integer_ref *A_ATTRIBUTE_DEFINITION_INTEGER_REF) Unstage(stage *StageStruct) *A_ATTRIBUTE_DEFINITION_INTEGER_REF {
	delete(stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFs, a_attribute_definition_integer_ref)
	delete(stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFs_mapString, a_attribute_definition_integer_ref.Name)
	return a_attribute_definition_integer_ref
}

// UnstageVoid removes a_attribute_definition_integer_ref off the model stage
func (a_attribute_definition_integer_ref *A_ATTRIBUTE_DEFINITION_INTEGER_REF) UnstageVoid(stage *StageStruct) {
	delete(stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFs, a_attribute_definition_integer_ref)
	delete(stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFs_mapString, a_attribute_definition_integer_ref.Name)
}

// commit a_attribute_definition_integer_ref to the back repo (if it is already staged)
func (a_attribute_definition_integer_ref *A_ATTRIBUTE_DEFINITION_INTEGER_REF) Commit(stage *StageStruct) *A_ATTRIBUTE_DEFINITION_INTEGER_REF {
	if _, ok := stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFs[a_attribute_definition_integer_ref]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitA_ATTRIBUTE_DEFINITION_INTEGER_REF(a_attribute_definition_integer_ref)
		}
	}
	return a_attribute_definition_integer_ref
}

func (a_attribute_definition_integer_ref *A_ATTRIBUTE_DEFINITION_INTEGER_REF) CommitVoid(stage *StageStruct) {
	a_attribute_definition_integer_ref.Commit(stage)
}

// Checkout a_attribute_definition_integer_ref to the back repo (if it is already staged)
func (a_attribute_definition_integer_ref *A_ATTRIBUTE_DEFINITION_INTEGER_REF) Checkout(stage *StageStruct) *A_ATTRIBUTE_DEFINITION_INTEGER_REF {
	if _, ok := stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFs[a_attribute_definition_integer_ref]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutA_ATTRIBUTE_DEFINITION_INTEGER_REF(a_attribute_definition_integer_ref)
		}
	}
	return a_attribute_definition_integer_ref
}

// for satisfaction of GongStruct interface
func (a_attribute_definition_integer_ref *A_ATTRIBUTE_DEFINITION_INTEGER_REF) GetName() (res string) {
	return a_attribute_definition_integer_ref.Name
}

// Stage puts a_attribute_definition_real_ref to the model stage
func (a_attribute_definition_real_ref *A_ATTRIBUTE_DEFINITION_REAL_REF) Stage(stage *StageStruct) *A_ATTRIBUTE_DEFINITION_REAL_REF {
	stage.A_ATTRIBUTE_DEFINITION_REAL_REFs[a_attribute_definition_real_ref] = __member
	stage.A_ATTRIBUTE_DEFINITION_REAL_REFs_mapString[a_attribute_definition_real_ref.Name] = a_attribute_definition_real_ref

	return a_attribute_definition_real_ref
}

// Unstage removes a_attribute_definition_real_ref off the model stage
func (a_attribute_definition_real_ref *A_ATTRIBUTE_DEFINITION_REAL_REF) Unstage(stage *StageStruct) *A_ATTRIBUTE_DEFINITION_REAL_REF {
	delete(stage.A_ATTRIBUTE_DEFINITION_REAL_REFs, a_attribute_definition_real_ref)
	delete(stage.A_ATTRIBUTE_DEFINITION_REAL_REFs_mapString, a_attribute_definition_real_ref.Name)
	return a_attribute_definition_real_ref
}

// UnstageVoid removes a_attribute_definition_real_ref off the model stage
func (a_attribute_definition_real_ref *A_ATTRIBUTE_DEFINITION_REAL_REF) UnstageVoid(stage *StageStruct) {
	delete(stage.A_ATTRIBUTE_DEFINITION_REAL_REFs, a_attribute_definition_real_ref)
	delete(stage.A_ATTRIBUTE_DEFINITION_REAL_REFs_mapString, a_attribute_definition_real_ref.Name)
}

// commit a_attribute_definition_real_ref to the back repo (if it is already staged)
func (a_attribute_definition_real_ref *A_ATTRIBUTE_DEFINITION_REAL_REF) Commit(stage *StageStruct) *A_ATTRIBUTE_DEFINITION_REAL_REF {
	if _, ok := stage.A_ATTRIBUTE_DEFINITION_REAL_REFs[a_attribute_definition_real_ref]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitA_ATTRIBUTE_DEFINITION_REAL_REF(a_attribute_definition_real_ref)
		}
	}
	return a_attribute_definition_real_ref
}

func (a_attribute_definition_real_ref *A_ATTRIBUTE_DEFINITION_REAL_REF) CommitVoid(stage *StageStruct) {
	a_attribute_definition_real_ref.Commit(stage)
}

// Checkout a_attribute_definition_real_ref to the back repo (if it is already staged)
func (a_attribute_definition_real_ref *A_ATTRIBUTE_DEFINITION_REAL_REF) Checkout(stage *StageStruct) *A_ATTRIBUTE_DEFINITION_REAL_REF {
	if _, ok := stage.A_ATTRIBUTE_DEFINITION_REAL_REFs[a_attribute_definition_real_ref]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutA_ATTRIBUTE_DEFINITION_REAL_REF(a_attribute_definition_real_ref)
		}
	}
	return a_attribute_definition_real_ref
}

// for satisfaction of GongStruct interface
func (a_attribute_definition_real_ref *A_ATTRIBUTE_DEFINITION_REAL_REF) GetName() (res string) {
	return a_attribute_definition_real_ref.Name
}

// Stage puts a_attribute_definition_string_ref to the model stage
func (a_attribute_definition_string_ref *A_ATTRIBUTE_DEFINITION_STRING_REF) Stage(stage *StageStruct) *A_ATTRIBUTE_DEFINITION_STRING_REF {
	stage.A_ATTRIBUTE_DEFINITION_STRING_REFs[a_attribute_definition_string_ref] = __member
	stage.A_ATTRIBUTE_DEFINITION_STRING_REFs_mapString[a_attribute_definition_string_ref.Name] = a_attribute_definition_string_ref

	return a_attribute_definition_string_ref
}

// Unstage removes a_attribute_definition_string_ref off the model stage
func (a_attribute_definition_string_ref *A_ATTRIBUTE_DEFINITION_STRING_REF) Unstage(stage *StageStruct) *A_ATTRIBUTE_DEFINITION_STRING_REF {
	delete(stage.A_ATTRIBUTE_DEFINITION_STRING_REFs, a_attribute_definition_string_ref)
	delete(stage.A_ATTRIBUTE_DEFINITION_STRING_REFs_mapString, a_attribute_definition_string_ref.Name)
	return a_attribute_definition_string_ref
}

// UnstageVoid removes a_attribute_definition_string_ref off the model stage
func (a_attribute_definition_string_ref *A_ATTRIBUTE_DEFINITION_STRING_REF) UnstageVoid(stage *StageStruct) {
	delete(stage.A_ATTRIBUTE_DEFINITION_STRING_REFs, a_attribute_definition_string_ref)
	delete(stage.A_ATTRIBUTE_DEFINITION_STRING_REFs_mapString, a_attribute_definition_string_ref.Name)
}

// commit a_attribute_definition_string_ref to the back repo (if it is already staged)
func (a_attribute_definition_string_ref *A_ATTRIBUTE_DEFINITION_STRING_REF) Commit(stage *StageStruct) *A_ATTRIBUTE_DEFINITION_STRING_REF {
	if _, ok := stage.A_ATTRIBUTE_DEFINITION_STRING_REFs[a_attribute_definition_string_ref]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitA_ATTRIBUTE_DEFINITION_STRING_REF(a_attribute_definition_string_ref)
		}
	}
	return a_attribute_definition_string_ref
}

func (a_attribute_definition_string_ref *A_ATTRIBUTE_DEFINITION_STRING_REF) CommitVoid(stage *StageStruct) {
	a_attribute_definition_string_ref.Commit(stage)
}

// Checkout a_attribute_definition_string_ref to the back repo (if it is already staged)
func (a_attribute_definition_string_ref *A_ATTRIBUTE_DEFINITION_STRING_REF) Checkout(stage *StageStruct) *A_ATTRIBUTE_DEFINITION_STRING_REF {
	if _, ok := stage.A_ATTRIBUTE_DEFINITION_STRING_REFs[a_attribute_definition_string_ref]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutA_ATTRIBUTE_DEFINITION_STRING_REF(a_attribute_definition_string_ref)
		}
	}
	return a_attribute_definition_string_ref
}

// for satisfaction of GongStruct interface
func (a_attribute_definition_string_ref *A_ATTRIBUTE_DEFINITION_STRING_REF) GetName() (res string) {
	return a_attribute_definition_string_ref.Name
}

// Stage puts a_attribute_definition_xhtml_ref to the model stage
func (a_attribute_definition_xhtml_ref *A_ATTRIBUTE_DEFINITION_XHTML_REF) Stage(stage *StageStruct) *A_ATTRIBUTE_DEFINITION_XHTML_REF {
	stage.A_ATTRIBUTE_DEFINITION_XHTML_REFs[a_attribute_definition_xhtml_ref] = __member
	stage.A_ATTRIBUTE_DEFINITION_XHTML_REFs_mapString[a_attribute_definition_xhtml_ref.Name] = a_attribute_definition_xhtml_ref

	return a_attribute_definition_xhtml_ref
}

// Unstage removes a_attribute_definition_xhtml_ref off the model stage
func (a_attribute_definition_xhtml_ref *A_ATTRIBUTE_DEFINITION_XHTML_REF) Unstage(stage *StageStruct) *A_ATTRIBUTE_DEFINITION_XHTML_REF {
	delete(stage.A_ATTRIBUTE_DEFINITION_XHTML_REFs, a_attribute_definition_xhtml_ref)
	delete(stage.A_ATTRIBUTE_DEFINITION_XHTML_REFs_mapString, a_attribute_definition_xhtml_ref.Name)
	return a_attribute_definition_xhtml_ref
}

// UnstageVoid removes a_attribute_definition_xhtml_ref off the model stage
func (a_attribute_definition_xhtml_ref *A_ATTRIBUTE_DEFINITION_XHTML_REF) UnstageVoid(stage *StageStruct) {
	delete(stage.A_ATTRIBUTE_DEFINITION_XHTML_REFs, a_attribute_definition_xhtml_ref)
	delete(stage.A_ATTRIBUTE_DEFINITION_XHTML_REFs_mapString, a_attribute_definition_xhtml_ref.Name)
}

// commit a_attribute_definition_xhtml_ref to the back repo (if it is already staged)
func (a_attribute_definition_xhtml_ref *A_ATTRIBUTE_DEFINITION_XHTML_REF) Commit(stage *StageStruct) *A_ATTRIBUTE_DEFINITION_XHTML_REF {
	if _, ok := stage.A_ATTRIBUTE_DEFINITION_XHTML_REFs[a_attribute_definition_xhtml_ref]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitA_ATTRIBUTE_DEFINITION_XHTML_REF(a_attribute_definition_xhtml_ref)
		}
	}
	return a_attribute_definition_xhtml_ref
}

func (a_attribute_definition_xhtml_ref *A_ATTRIBUTE_DEFINITION_XHTML_REF) CommitVoid(stage *StageStruct) {
	a_attribute_definition_xhtml_ref.Commit(stage)
}

// Checkout a_attribute_definition_xhtml_ref to the back repo (if it is already staged)
func (a_attribute_definition_xhtml_ref *A_ATTRIBUTE_DEFINITION_XHTML_REF) Checkout(stage *StageStruct) *A_ATTRIBUTE_DEFINITION_XHTML_REF {
	if _, ok := stage.A_ATTRIBUTE_DEFINITION_XHTML_REFs[a_attribute_definition_xhtml_ref]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutA_ATTRIBUTE_DEFINITION_XHTML_REF(a_attribute_definition_xhtml_ref)
		}
	}
	return a_attribute_definition_xhtml_ref
}

// for satisfaction of GongStruct interface
func (a_attribute_definition_xhtml_ref *A_ATTRIBUTE_DEFINITION_XHTML_REF) GetName() (res string) {
	return a_attribute_definition_xhtml_ref.Name
}

// Stage puts a_attribute_value_boolean to the model stage
func (a_attribute_value_boolean *A_ATTRIBUTE_VALUE_BOOLEAN) Stage(stage *StageStruct) *A_ATTRIBUTE_VALUE_BOOLEAN {
	stage.A_ATTRIBUTE_VALUE_BOOLEANs[a_attribute_value_boolean] = __member
	stage.A_ATTRIBUTE_VALUE_BOOLEANs_mapString[a_attribute_value_boolean.Name] = a_attribute_value_boolean

	return a_attribute_value_boolean
}

// Unstage removes a_attribute_value_boolean off the model stage
func (a_attribute_value_boolean *A_ATTRIBUTE_VALUE_BOOLEAN) Unstage(stage *StageStruct) *A_ATTRIBUTE_VALUE_BOOLEAN {
	delete(stage.A_ATTRIBUTE_VALUE_BOOLEANs, a_attribute_value_boolean)
	delete(stage.A_ATTRIBUTE_VALUE_BOOLEANs_mapString, a_attribute_value_boolean.Name)
	return a_attribute_value_boolean
}

// UnstageVoid removes a_attribute_value_boolean off the model stage
func (a_attribute_value_boolean *A_ATTRIBUTE_VALUE_BOOLEAN) UnstageVoid(stage *StageStruct) {
	delete(stage.A_ATTRIBUTE_VALUE_BOOLEANs, a_attribute_value_boolean)
	delete(stage.A_ATTRIBUTE_VALUE_BOOLEANs_mapString, a_attribute_value_boolean.Name)
}

// commit a_attribute_value_boolean to the back repo (if it is already staged)
func (a_attribute_value_boolean *A_ATTRIBUTE_VALUE_BOOLEAN) Commit(stage *StageStruct) *A_ATTRIBUTE_VALUE_BOOLEAN {
	if _, ok := stage.A_ATTRIBUTE_VALUE_BOOLEANs[a_attribute_value_boolean]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitA_ATTRIBUTE_VALUE_BOOLEAN(a_attribute_value_boolean)
		}
	}
	return a_attribute_value_boolean
}

func (a_attribute_value_boolean *A_ATTRIBUTE_VALUE_BOOLEAN) CommitVoid(stage *StageStruct) {
	a_attribute_value_boolean.Commit(stage)
}

// Checkout a_attribute_value_boolean to the back repo (if it is already staged)
func (a_attribute_value_boolean *A_ATTRIBUTE_VALUE_BOOLEAN) Checkout(stage *StageStruct) *A_ATTRIBUTE_VALUE_BOOLEAN {
	if _, ok := stage.A_ATTRIBUTE_VALUE_BOOLEANs[a_attribute_value_boolean]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutA_ATTRIBUTE_VALUE_BOOLEAN(a_attribute_value_boolean)
		}
	}
	return a_attribute_value_boolean
}

// for satisfaction of GongStruct interface
func (a_attribute_value_boolean *A_ATTRIBUTE_VALUE_BOOLEAN) GetName() (res string) {
	return a_attribute_value_boolean.Name
}

// Stage puts a_attribute_value_date to the model stage
func (a_attribute_value_date *A_ATTRIBUTE_VALUE_DATE) Stage(stage *StageStruct) *A_ATTRIBUTE_VALUE_DATE {
	stage.A_ATTRIBUTE_VALUE_DATEs[a_attribute_value_date] = __member
	stage.A_ATTRIBUTE_VALUE_DATEs_mapString[a_attribute_value_date.Name] = a_attribute_value_date

	return a_attribute_value_date
}

// Unstage removes a_attribute_value_date off the model stage
func (a_attribute_value_date *A_ATTRIBUTE_VALUE_DATE) Unstage(stage *StageStruct) *A_ATTRIBUTE_VALUE_DATE {
	delete(stage.A_ATTRIBUTE_VALUE_DATEs, a_attribute_value_date)
	delete(stage.A_ATTRIBUTE_VALUE_DATEs_mapString, a_attribute_value_date.Name)
	return a_attribute_value_date
}

// UnstageVoid removes a_attribute_value_date off the model stage
func (a_attribute_value_date *A_ATTRIBUTE_VALUE_DATE) UnstageVoid(stage *StageStruct) {
	delete(stage.A_ATTRIBUTE_VALUE_DATEs, a_attribute_value_date)
	delete(stage.A_ATTRIBUTE_VALUE_DATEs_mapString, a_attribute_value_date.Name)
}

// commit a_attribute_value_date to the back repo (if it is already staged)
func (a_attribute_value_date *A_ATTRIBUTE_VALUE_DATE) Commit(stage *StageStruct) *A_ATTRIBUTE_VALUE_DATE {
	if _, ok := stage.A_ATTRIBUTE_VALUE_DATEs[a_attribute_value_date]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitA_ATTRIBUTE_VALUE_DATE(a_attribute_value_date)
		}
	}
	return a_attribute_value_date
}

func (a_attribute_value_date *A_ATTRIBUTE_VALUE_DATE) CommitVoid(stage *StageStruct) {
	a_attribute_value_date.Commit(stage)
}

// Checkout a_attribute_value_date to the back repo (if it is already staged)
func (a_attribute_value_date *A_ATTRIBUTE_VALUE_DATE) Checkout(stage *StageStruct) *A_ATTRIBUTE_VALUE_DATE {
	if _, ok := stage.A_ATTRIBUTE_VALUE_DATEs[a_attribute_value_date]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutA_ATTRIBUTE_VALUE_DATE(a_attribute_value_date)
		}
	}
	return a_attribute_value_date
}

// for satisfaction of GongStruct interface
func (a_attribute_value_date *A_ATTRIBUTE_VALUE_DATE) GetName() (res string) {
	return a_attribute_value_date.Name
}

// Stage puts a_attribute_value_enumeration to the model stage
func (a_attribute_value_enumeration *A_ATTRIBUTE_VALUE_ENUMERATION) Stage(stage *StageStruct) *A_ATTRIBUTE_VALUE_ENUMERATION {
	stage.A_ATTRIBUTE_VALUE_ENUMERATIONs[a_attribute_value_enumeration] = __member
	stage.A_ATTRIBUTE_VALUE_ENUMERATIONs_mapString[a_attribute_value_enumeration.Name] = a_attribute_value_enumeration

	return a_attribute_value_enumeration
}

// Unstage removes a_attribute_value_enumeration off the model stage
func (a_attribute_value_enumeration *A_ATTRIBUTE_VALUE_ENUMERATION) Unstage(stage *StageStruct) *A_ATTRIBUTE_VALUE_ENUMERATION {
	delete(stage.A_ATTRIBUTE_VALUE_ENUMERATIONs, a_attribute_value_enumeration)
	delete(stage.A_ATTRIBUTE_VALUE_ENUMERATIONs_mapString, a_attribute_value_enumeration.Name)
	return a_attribute_value_enumeration
}

// UnstageVoid removes a_attribute_value_enumeration off the model stage
func (a_attribute_value_enumeration *A_ATTRIBUTE_VALUE_ENUMERATION) UnstageVoid(stage *StageStruct) {
	delete(stage.A_ATTRIBUTE_VALUE_ENUMERATIONs, a_attribute_value_enumeration)
	delete(stage.A_ATTRIBUTE_VALUE_ENUMERATIONs_mapString, a_attribute_value_enumeration.Name)
}

// commit a_attribute_value_enumeration to the back repo (if it is already staged)
func (a_attribute_value_enumeration *A_ATTRIBUTE_VALUE_ENUMERATION) Commit(stage *StageStruct) *A_ATTRIBUTE_VALUE_ENUMERATION {
	if _, ok := stage.A_ATTRIBUTE_VALUE_ENUMERATIONs[a_attribute_value_enumeration]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitA_ATTRIBUTE_VALUE_ENUMERATION(a_attribute_value_enumeration)
		}
	}
	return a_attribute_value_enumeration
}

func (a_attribute_value_enumeration *A_ATTRIBUTE_VALUE_ENUMERATION) CommitVoid(stage *StageStruct) {
	a_attribute_value_enumeration.Commit(stage)
}

// Checkout a_attribute_value_enumeration to the back repo (if it is already staged)
func (a_attribute_value_enumeration *A_ATTRIBUTE_VALUE_ENUMERATION) Checkout(stage *StageStruct) *A_ATTRIBUTE_VALUE_ENUMERATION {
	if _, ok := stage.A_ATTRIBUTE_VALUE_ENUMERATIONs[a_attribute_value_enumeration]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutA_ATTRIBUTE_VALUE_ENUMERATION(a_attribute_value_enumeration)
		}
	}
	return a_attribute_value_enumeration
}

// for satisfaction of GongStruct interface
func (a_attribute_value_enumeration *A_ATTRIBUTE_VALUE_ENUMERATION) GetName() (res string) {
	return a_attribute_value_enumeration.Name
}

// Stage puts a_attribute_value_integer to the model stage
func (a_attribute_value_integer *A_ATTRIBUTE_VALUE_INTEGER) Stage(stage *StageStruct) *A_ATTRIBUTE_VALUE_INTEGER {
	stage.A_ATTRIBUTE_VALUE_INTEGERs[a_attribute_value_integer] = __member
	stage.A_ATTRIBUTE_VALUE_INTEGERs_mapString[a_attribute_value_integer.Name] = a_attribute_value_integer

	return a_attribute_value_integer
}

// Unstage removes a_attribute_value_integer off the model stage
func (a_attribute_value_integer *A_ATTRIBUTE_VALUE_INTEGER) Unstage(stage *StageStruct) *A_ATTRIBUTE_VALUE_INTEGER {
	delete(stage.A_ATTRIBUTE_VALUE_INTEGERs, a_attribute_value_integer)
	delete(stage.A_ATTRIBUTE_VALUE_INTEGERs_mapString, a_attribute_value_integer.Name)
	return a_attribute_value_integer
}

// UnstageVoid removes a_attribute_value_integer off the model stage
func (a_attribute_value_integer *A_ATTRIBUTE_VALUE_INTEGER) UnstageVoid(stage *StageStruct) {
	delete(stage.A_ATTRIBUTE_VALUE_INTEGERs, a_attribute_value_integer)
	delete(stage.A_ATTRIBUTE_VALUE_INTEGERs_mapString, a_attribute_value_integer.Name)
}

// commit a_attribute_value_integer to the back repo (if it is already staged)
func (a_attribute_value_integer *A_ATTRIBUTE_VALUE_INTEGER) Commit(stage *StageStruct) *A_ATTRIBUTE_VALUE_INTEGER {
	if _, ok := stage.A_ATTRIBUTE_VALUE_INTEGERs[a_attribute_value_integer]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitA_ATTRIBUTE_VALUE_INTEGER(a_attribute_value_integer)
		}
	}
	return a_attribute_value_integer
}

func (a_attribute_value_integer *A_ATTRIBUTE_VALUE_INTEGER) CommitVoid(stage *StageStruct) {
	a_attribute_value_integer.Commit(stage)
}

// Checkout a_attribute_value_integer to the back repo (if it is already staged)
func (a_attribute_value_integer *A_ATTRIBUTE_VALUE_INTEGER) Checkout(stage *StageStruct) *A_ATTRIBUTE_VALUE_INTEGER {
	if _, ok := stage.A_ATTRIBUTE_VALUE_INTEGERs[a_attribute_value_integer]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutA_ATTRIBUTE_VALUE_INTEGER(a_attribute_value_integer)
		}
	}
	return a_attribute_value_integer
}

// for satisfaction of GongStruct interface
func (a_attribute_value_integer *A_ATTRIBUTE_VALUE_INTEGER) GetName() (res string) {
	return a_attribute_value_integer.Name
}

// Stage puts a_attribute_value_real to the model stage
func (a_attribute_value_real *A_ATTRIBUTE_VALUE_REAL) Stage(stage *StageStruct) *A_ATTRIBUTE_VALUE_REAL {
	stage.A_ATTRIBUTE_VALUE_REALs[a_attribute_value_real] = __member
	stage.A_ATTRIBUTE_VALUE_REALs_mapString[a_attribute_value_real.Name] = a_attribute_value_real

	return a_attribute_value_real
}

// Unstage removes a_attribute_value_real off the model stage
func (a_attribute_value_real *A_ATTRIBUTE_VALUE_REAL) Unstage(stage *StageStruct) *A_ATTRIBUTE_VALUE_REAL {
	delete(stage.A_ATTRIBUTE_VALUE_REALs, a_attribute_value_real)
	delete(stage.A_ATTRIBUTE_VALUE_REALs_mapString, a_attribute_value_real.Name)
	return a_attribute_value_real
}

// UnstageVoid removes a_attribute_value_real off the model stage
func (a_attribute_value_real *A_ATTRIBUTE_VALUE_REAL) UnstageVoid(stage *StageStruct) {
	delete(stage.A_ATTRIBUTE_VALUE_REALs, a_attribute_value_real)
	delete(stage.A_ATTRIBUTE_VALUE_REALs_mapString, a_attribute_value_real.Name)
}

// commit a_attribute_value_real to the back repo (if it is already staged)
func (a_attribute_value_real *A_ATTRIBUTE_VALUE_REAL) Commit(stage *StageStruct) *A_ATTRIBUTE_VALUE_REAL {
	if _, ok := stage.A_ATTRIBUTE_VALUE_REALs[a_attribute_value_real]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitA_ATTRIBUTE_VALUE_REAL(a_attribute_value_real)
		}
	}
	return a_attribute_value_real
}

func (a_attribute_value_real *A_ATTRIBUTE_VALUE_REAL) CommitVoid(stage *StageStruct) {
	a_attribute_value_real.Commit(stage)
}

// Checkout a_attribute_value_real to the back repo (if it is already staged)
func (a_attribute_value_real *A_ATTRIBUTE_VALUE_REAL) Checkout(stage *StageStruct) *A_ATTRIBUTE_VALUE_REAL {
	if _, ok := stage.A_ATTRIBUTE_VALUE_REALs[a_attribute_value_real]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutA_ATTRIBUTE_VALUE_REAL(a_attribute_value_real)
		}
	}
	return a_attribute_value_real
}

// for satisfaction of GongStruct interface
func (a_attribute_value_real *A_ATTRIBUTE_VALUE_REAL) GetName() (res string) {
	return a_attribute_value_real.Name
}

// Stage puts a_attribute_value_string to the model stage
func (a_attribute_value_string *A_ATTRIBUTE_VALUE_STRING) Stage(stage *StageStruct) *A_ATTRIBUTE_VALUE_STRING {
	stage.A_ATTRIBUTE_VALUE_STRINGs[a_attribute_value_string] = __member
	stage.A_ATTRIBUTE_VALUE_STRINGs_mapString[a_attribute_value_string.Name] = a_attribute_value_string

	return a_attribute_value_string
}

// Unstage removes a_attribute_value_string off the model stage
func (a_attribute_value_string *A_ATTRIBUTE_VALUE_STRING) Unstage(stage *StageStruct) *A_ATTRIBUTE_VALUE_STRING {
	delete(stage.A_ATTRIBUTE_VALUE_STRINGs, a_attribute_value_string)
	delete(stage.A_ATTRIBUTE_VALUE_STRINGs_mapString, a_attribute_value_string.Name)
	return a_attribute_value_string
}

// UnstageVoid removes a_attribute_value_string off the model stage
func (a_attribute_value_string *A_ATTRIBUTE_VALUE_STRING) UnstageVoid(stage *StageStruct) {
	delete(stage.A_ATTRIBUTE_VALUE_STRINGs, a_attribute_value_string)
	delete(stage.A_ATTRIBUTE_VALUE_STRINGs_mapString, a_attribute_value_string.Name)
}

// commit a_attribute_value_string to the back repo (if it is already staged)
func (a_attribute_value_string *A_ATTRIBUTE_VALUE_STRING) Commit(stage *StageStruct) *A_ATTRIBUTE_VALUE_STRING {
	if _, ok := stage.A_ATTRIBUTE_VALUE_STRINGs[a_attribute_value_string]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitA_ATTRIBUTE_VALUE_STRING(a_attribute_value_string)
		}
	}
	return a_attribute_value_string
}

func (a_attribute_value_string *A_ATTRIBUTE_VALUE_STRING) CommitVoid(stage *StageStruct) {
	a_attribute_value_string.Commit(stage)
}

// Checkout a_attribute_value_string to the back repo (if it is already staged)
func (a_attribute_value_string *A_ATTRIBUTE_VALUE_STRING) Checkout(stage *StageStruct) *A_ATTRIBUTE_VALUE_STRING {
	if _, ok := stage.A_ATTRIBUTE_VALUE_STRINGs[a_attribute_value_string]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutA_ATTRIBUTE_VALUE_STRING(a_attribute_value_string)
		}
	}
	return a_attribute_value_string
}

// for satisfaction of GongStruct interface
func (a_attribute_value_string *A_ATTRIBUTE_VALUE_STRING) GetName() (res string) {
	return a_attribute_value_string.Name
}

// Stage puts a_attribute_value_xhtml to the model stage
func (a_attribute_value_xhtml *A_ATTRIBUTE_VALUE_XHTML) Stage(stage *StageStruct) *A_ATTRIBUTE_VALUE_XHTML {
	stage.A_ATTRIBUTE_VALUE_XHTMLs[a_attribute_value_xhtml] = __member
	stage.A_ATTRIBUTE_VALUE_XHTMLs_mapString[a_attribute_value_xhtml.Name] = a_attribute_value_xhtml

	return a_attribute_value_xhtml
}

// Unstage removes a_attribute_value_xhtml off the model stage
func (a_attribute_value_xhtml *A_ATTRIBUTE_VALUE_XHTML) Unstage(stage *StageStruct) *A_ATTRIBUTE_VALUE_XHTML {
	delete(stage.A_ATTRIBUTE_VALUE_XHTMLs, a_attribute_value_xhtml)
	delete(stage.A_ATTRIBUTE_VALUE_XHTMLs_mapString, a_attribute_value_xhtml.Name)
	return a_attribute_value_xhtml
}

// UnstageVoid removes a_attribute_value_xhtml off the model stage
func (a_attribute_value_xhtml *A_ATTRIBUTE_VALUE_XHTML) UnstageVoid(stage *StageStruct) {
	delete(stage.A_ATTRIBUTE_VALUE_XHTMLs, a_attribute_value_xhtml)
	delete(stage.A_ATTRIBUTE_VALUE_XHTMLs_mapString, a_attribute_value_xhtml.Name)
}

// commit a_attribute_value_xhtml to the back repo (if it is already staged)
func (a_attribute_value_xhtml *A_ATTRIBUTE_VALUE_XHTML) Commit(stage *StageStruct) *A_ATTRIBUTE_VALUE_XHTML {
	if _, ok := stage.A_ATTRIBUTE_VALUE_XHTMLs[a_attribute_value_xhtml]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitA_ATTRIBUTE_VALUE_XHTML(a_attribute_value_xhtml)
		}
	}
	return a_attribute_value_xhtml
}

func (a_attribute_value_xhtml *A_ATTRIBUTE_VALUE_XHTML) CommitVoid(stage *StageStruct) {
	a_attribute_value_xhtml.Commit(stage)
}

// Checkout a_attribute_value_xhtml to the back repo (if it is already staged)
func (a_attribute_value_xhtml *A_ATTRIBUTE_VALUE_XHTML) Checkout(stage *StageStruct) *A_ATTRIBUTE_VALUE_XHTML {
	if _, ok := stage.A_ATTRIBUTE_VALUE_XHTMLs[a_attribute_value_xhtml]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutA_ATTRIBUTE_VALUE_XHTML(a_attribute_value_xhtml)
		}
	}
	return a_attribute_value_xhtml
}

// for satisfaction of GongStruct interface
func (a_attribute_value_xhtml *A_ATTRIBUTE_VALUE_XHTML) GetName() (res string) {
	return a_attribute_value_xhtml.Name
}

// Stage puts a_attribute_value_xhtml_1 to the model stage
func (a_attribute_value_xhtml_1 *A_ATTRIBUTE_VALUE_XHTML_1) Stage(stage *StageStruct) *A_ATTRIBUTE_VALUE_XHTML_1 {
	stage.A_ATTRIBUTE_VALUE_XHTML_1s[a_attribute_value_xhtml_1] = __member
	stage.A_ATTRIBUTE_VALUE_XHTML_1s_mapString[a_attribute_value_xhtml_1.Name] = a_attribute_value_xhtml_1

	return a_attribute_value_xhtml_1
}

// Unstage removes a_attribute_value_xhtml_1 off the model stage
func (a_attribute_value_xhtml_1 *A_ATTRIBUTE_VALUE_XHTML_1) Unstage(stage *StageStruct) *A_ATTRIBUTE_VALUE_XHTML_1 {
	delete(stage.A_ATTRIBUTE_VALUE_XHTML_1s, a_attribute_value_xhtml_1)
	delete(stage.A_ATTRIBUTE_VALUE_XHTML_1s_mapString, a_attribute_value_xhtml_1.Name)
	return a_attribute_value_xhtml_1
}

// UnstageVoid removes a_attribute_value_xhtml_1 off the model stage
func (a_attribute_value_xhtml_1 *A_ATTRIBUTE_VALUE_XHTML_1) UnstageVoid(stage *StageStruct) {
	delete(stage.A_ATTRIBUTE_VALUE_XHTML_1s, a_attribute_value_xhtml_1)
	delete(stage.A_ATTRIBUTE_VALUE_XHTML_1s_mapString, a_attribute_value_xhtml_1.Name)
}

// commit a_attribute_value_xhtml_1 to the back repo (if it is already staged)
func (a_attribute_value_xhtml_1 *A_ATTRIBUTE_VALUE_XHTML_1) Commit(stage *StageStruct) *A_ATTRIBUTE_VALUE_XHTML_1 {
	if _, ok := stage.A_ATTRIBUTE_VALUE_XHTML_1s[a_attribute_value_xhtml_1]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitA_ATTRIBUTE_VALUE_XHTML_1(a_attribute_value_xhtml_1)
		}
	}
	return a_attribute_value_xhtml_1
}

func (a_attribute_value_xhtml_1 *A_ATTRIBUTE_VALUE_XHTML_1) CommitVoid(stage *StageStruct) {
	a_attribute_value_xhtml_1.Commit(stage)
}

// Checkout a_attribute_value_xhtml_1 to the back repo (if it is already staged)
func (a_attribute_value_xhtml_1 *A_ATTRIBUTE_VALUE_XHTML_1) Checkout(stage *StageStruct) *A_ATTRIBUTE_VALUE_XHTML_1 {
	if _, ok := stage.A_ATTRIBUTE_VALUE_XHTML_1s[a_attribute_value_xhtml_1]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutA_ATTRIBUTE_VALUE_XHTML_1(a_attribute_value_xhtml_1)
		}
	}
	return a_attribute_value_xhtml_1
}

// for satisfaction of GongStruct interface
func (a_attribute_value_xhtml_1 *A_ATTRIBUTE_VALUE_XHTML_1) GetName() (res string) {
	return a_attribute_value_xhtml_1.Name
}

// Stage puts a_children to the model stage
func (a_children *A_CHILDREN) Stage(stage *StageStruct) *A_CHILDREN {
	stage.A_CHILDRENs[a_children] = __member
	stage.A_CHILDRENs_mapString[a_children.Name] = a_children

	return a_children
}

// Unstage removes a_children off the model stage
func (a_children *A_CHILDREN) Unstage(stage *StageStruct) *A_CHILDREN {
	delete(stage.A_CHILDRENs, a_children)
	delete(stage.A_CHILDRENs_mapString, a_children.Name)
	return a_children
}

// UnstageVoid removes a_children off the model stage
func (a_children *A_CHILDREN) UnstageVoid(stage *StageStruct) {
	delete(stage.A_CHILDRENs, a_children)
	delete(stage.A_CHILDRENs_mapString, a_children.Name)
}

// commit a_children to the back repo (if it is already staged)
func (a_children *A_CHILDREN) Commit(stage *StageStruct) *A_CHILDREN {
	if _, ok := stage.A_CHILDRENs[a_children]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitA_CHILDREN(a_children)
		}
	}
	return a_children
}

func (a_children *A_CHILDREN) CommitVoid(stage *StageStruct) {
	a_children.Commit(stage)
}

// Checkout a_children to the back repo (if it is already staged)
func (a_children *A_CHILDREN) Checkout(stage *StageStruct) *A_CHILDREN {
	if _, ok := stage.A_CHILDRENs[a_children]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutA_CHILDREN(a_children)
		}
	}
	return a_children
}

// for satisfaction of GongStruct interface
func (a_children *A_CHILDREN) GetName() (res string) {
	return a_children.Name
}

// Stage puts a_core_content to the model stage
func (a_core_content *A_CORE_CONTENT) Stage(stage *StageStruct) *A_CORE_CONTENT {
	stage.A_CORE_CONTENTs[a_core_content] = __member
	stage.A_CORE_CONTENTs_mapString[a_core_content.Name] = a_core_content

	return a_core_content
}

// Unstage removes a_core_content off the model stage
func (a_core_content *A_CORE_CONTENT) Unstage(stage *StageStruct) *A_CORE_CONTENT {
	delete(stage.A_CORE_CONTENTs, a_core_content)
	delete(stage.A_CORE_CONTENTs_mapString, a_core_content.Name)
	return a_core_content
}

// UnstageVoid removes a_core_content off the model stage
func (a_core_content *A_CORE_CONTENT) UnstageVoid(stage *StageStruct) {
	delete(stage.A_CORE_CONTENTs, a_core_content)
	delete(stage.A_CORE_CONTENTs_mapString, a_core_content.Name)
}

// commit a_core_content to the back repo (if it is already staged)
func (a_core_content *A_CORE_CONTENT) Commit(stage *StageStruct) *A_CORE_CONTENT {
	if _, ok := stage.A_CORE_CONTENTs[a_core_content]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitA_CORE_CONTENT(a_core_content)
		}
	}
	return a_core_content
}

func (a_core_content *A_CORE_CONTENT) CommitVoid(stage *StageStruct) {
	a_core_content.Commit(stage)
}

// Checkout a_core_content to the back repo (if it is already staged)
func (a_core_content *A_CORE_CONTENT) Checkout(stage *StageStruct) *A_CORE_CONTENT {
	if _, ok := stage.A_CORE_CONTENTs[a_core_content]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutA_CORE_CONTENT(a_core_content)
		}
	}
	return a_core_content
}

// for satisfaction of GongStruct interface
func (a_core_content *A_CORE_CONTENT) GetName() (res string) {
	return a_core_content.Name
}

// Stage puts a_datatypes to the model stage
func (a_datatypes *A_DATATYPES) Stage(stage *StageStruct) *A_DATATYPES {
	stage.A_DATATYPESs[a_datatypes] = __member
	stage.A_DATATYPESs_mapString[a_datatypes.Name] = a_datatypes

	return a_datatypes
}

// Unstage removes a_datatypes off the model stage
func (a_datatypes *A_DATATYPES) Unstage(stage *StageStruct) *A_DATATYPES {
	delete(stage.A_DATATYPESs, a_datatypes)
	delete(stage.A_DATATYPESs_mapString, a_datatypes.Name)
	return a_datatypes
}

// UnstageVoid removes a_datatypes off the model stage
func (a_datatypes *A_DATATYPES) UnstageVoid(stage *StageStruct) {
	delete(stage.A_DATATYPESs, a_datatypes)
	delete(stage.A_DATATYPESs_mapString, a_datatypes.Name)
}

// commit a_datatypes to the back repo (if it is already staged)
func (a_datatypes *A_DATATYPES) Commit(stage *StageStruct) *A_DATATYPES {
	if _, ok := stage.A_DATATYPESs[a_datatypes]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitA_DATATYPES(a_datatypes)
		}
	}
	return a_datatypes
}

func (a_datatypes *A_DATATYPES) CommitVoid(stage *StageStruct) {
	a_datatypes.Commit(stage)
}

// Checkout a_datatypes to the back repo (if it is already staged)
func (a_datatypes *A_DATATYPES) Checkout(stage *StageStruct) *A_DATATYPES {
	if _, ok := stage.A_DATATYPESs[a_datatypes]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutA_DATATYPES(a_datatypes)
		}
	}
	return a_datatypes
}

// for satisfaction of GongStruct interface
func (a_datatypes *A_DATATYPES) GetName() (res string) {
	return a_datatypes.Name
}

// Stage puts a_datatype_definition_boolean_ref to the model stage
func (a_datatype_definition_boolean_ref *A_DATATYPE_DEFINITION_BOOLEAN_REF) Stage(stage *StageStruct) *A_DATATYPE_DEFINITION_BOOLEAN_REF {
	stage.A_DATATYPE_DEFINITION_BOOLEAN_REFs[a_datatype_definition_boolean_ref] = __member
	stage.A_DATATYPE_DEFINITION_BOOLEAN_REFs_mapString[a_datatype_definition_boolean_ref.Name] = a_datatype_definition_boolean_ref

	return a_datatype_definition_boolean_ref
}

// Unstage removes a_datatype_definition_boolean_ref off the model stage
func (a_datatype_definition_boolean_ref *A_DATATYPE_DEFINITION_BOOLEAN_REF) Unstage(stage *StageStruct) *A_DATATYPE_DEFINITION_BOOLEAN_REF {
	delete(stage.A_DATATYPE_DEFINITION_BOOLEAN_REFs, a_datatype_definition_boolean_ref)
	delete(stage.A_DATATYPE_DEFINITION_BOOLEAN_REFs_mapString, a_datatype_definition_boolean_ref.Name)
	return a_datatype_definition_boolean_ref
}

// UnstageVoid removes a_datatype_definition_boolean_ref off the model stage
func (a_datatype_definition_boolean_ref *A_DATATYPE_DEFINITION_BOOLEAN_REF) UnstageVoid(stage *StageStruct) {
	delete(stage.A_DATATYPE_DEFINITION_BOOLEAN_REFs, a_datatype_definition_boolean_ref)
	delete(stage.A_DATATYPE_DEFINITION_BOOLEAN_REFs_mapString, a_datatype_definition_boolean_ref.Name)
}

// commit a_datatype_definition_boolean_ref to the back repo (if it is already staged)
func (a_datatype_definition_boolean_ref *A_DATATYPE_DEFINITION_BOOLEAN_REF) Commit(stage *StageStruct) *A_DATATYPE_DEFINITION_BOOLEAN_REF {
	if _, ok := stage.A_DATATYPE_DEFINITION_BOOLEAN_REFs[a_datatype_definition_boolean_ref]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitA_DATATYPE_DEFINITION_BOOLEAN_REF(a_datatype_definition_boolean_ref)
		}
	}
	return a_datatype_definition_boolean_ref
}

func (a_datatype_definition_boolean_ref *A_DATATYPE_DEFINITION_BOOLEAN_REF) CommitVoid(stage *StageStruct) {
	a_datatype_definition_boolean_ref.Commit(stage)
}

// Checkout a_datatype_definition_boolean_ref to the back repo (if it is already staged)
func (a_datatype_definition_boolean_ref *A_DATATYPE_DEFINITION_BOOLEAN_REF) Checkout(stage *StageStruct) *A_DATATYPE_DEFINITION_BOOLEAN_REF {
	if _, ok := stage.A_DATATYPE_DEFINITION_BOOLEAN_REFs[a_datatype_definition_boolean_ref]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutA_DATATYPE_DEFINITION_BOOLEAN_REF(a_datatype_definition_boolean_ref)
		}
	}
	return a_datatype_definition_boolean_ref
}

// for satisfaction of GongStruct interface
func (a_datatype_definition_boolean_ref *A_DATATYPE_DEFINITION_BOOLEAN_REF) GetName() (res string) {
	return a_datatype_definition_boolean_ref.Name
}

// Stage puts a_datatype_definition_date_ref to the model stage
func (a_datatype_definition_date_ref *A_DATATYPE_DEFINITION_DATE_REF) Stage(stage *StageStruct) *A_DATATYPE_DEFINITION_DATE_REF {
	stage.A_DATATYPE_DEFINITION_DATE_REFs[a_datatype_definition_date_ref] = __member
	stage.A_DATATYPE_DEFINITION_DATE_REFs_mapString[a_datatype_definition_date_ref.Name] = a_datatype_definition_date_ref

	return a_datatype_definition_date_ref
}

// Unstage removes a_datatype_definition_date_ref off the model stage
func (a_datatype_definition_date_ref *A_DATATYPE_DEFINITION_DATE_REF) Unstage(stage *StageStruct) *A_DATATYPE_DEFINITION_DATE_REF {
	delete(stage.A_DATATYPE_DEFINITION_DATE_REFs, a_datatype_definition_date_ref)
	delete(stage.A_DATATYPE_DEFINITION_DATE_REFs_mapString, a_datatype_definition_date_ref.Name)
	return a_datatype_definition_date_ref
}

// UnstageVoid removes a_datatype_definition_date_ref off the model stage
func (a_datatype_definition_date_ref *A_DATATYPE_DEFINITION_DATE_REF) UnstageVoid(stage *StageStruct) {
	delete(stage.A_DATATYPE_DEFINITION_DATE_REFs, a_datatype_definition_date_ref)
	delete(stage.A_DATATYPE_DEFINITION_DATE_REFs_mapString, a_datatype_definition_date_ref.Name)
}

// commit a_datatype_definition_date_ref to the back repo (if it is already staged)
func (a_datatype_definition_date_ref *A_DATATYPE_DEFINITION_DATE_REF) Commit(stage *StageStruct) *A_DATATYPE_DEFINITION_DATE_REF {
	if _, ok := stage.A_DATATYPE_DEFINITION_DATE_REFs[a_datatype_definition_date_ref]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitA_DATATYPE_DEFINITION_DATE_REF(a_datatype_definition_date_ref)
		}
	}
	return a_datatype_definition_date_ref
}

func (a_datatype_definition_date_ref *A_DATATYPE_DEFINITION_DATE_REF) CommitVoid(stage *StageStruct) {
	a_datatype_definition_date_ref.Commit(stage)
}

// Checkout a_datatype_definition_date_ref to the back repo (if it is already staged)
func (a_datatype_definition_date_ref *A_DATATYPE_DEFINITION_DATE_REF) Checkout(stage *StageStruct) *A_DATATYPE_DEFINITION_DATE_REF {
	if _, ok := stage.A_DATATYPE_DEFINITION_DATE_REFs[a_datatype_definition_date_ref]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutA_DATATYPE_DEFINITION_DATE_REF(a_datatype_definition_date_ref)
		}
	}
	return a_datatype_definition_date_ref
}

// for satisfaction of GongStruct interface
func (a_datatype_definition_date_ref *A_DATATYPE_DEFINITION_DATE_REF) GetName() (res string) {
	return a_datatype_definition_date_ref.Name
}

// Stage puts a_datatype_definition_enumeration_ref to the model stage
func (a_datatype_definition_enumeration_ref *A_DATATYPE_DEFINITION_ENUMERATION_REF) Stage(stage *StageStruct) *A_DATATYPE_DEFINITION_ENUMERATION_REF {
	stage.A_DATATYPE_DEFINITION_ENUMERATION_REFs[a_datatype_definition_enumeration_ref] = __member
	stage.A_DATATYPE_DEFINITION_ENUMERATION_REFs_mapString[a_datatype_definition_enumeration_ref.Name] = a_datatype_definition_enumeration_ref

	return a_datatype_definition_enumeration_ref
}

// Unstage removes a_datatype_definition_enumeration_ref off the model stage
func (a_datatype_definition_enumeration_ref *A_DATATYPE_DEFINITION_ENUMERATION_REF) Unstage(stage *StageStruct) *A_DATATYPE_DEFINITION_ENUMERATION_REF {
	delete(stage.A_DATATYPE_DEFINITION_ENUMERATION_REFs, a_datatype_definition_enumeration_ref)
	delete(stage.A_DATATYPE_DEFINITION_ENUMERATION_REFs_mapString, a_datatype_definition_enumeration_ref.Name)
	return a_datatype_definition_enumeration_ref
}

// UnstageVoid removes a_datatype_definition_enumeration_ref off the model stage
func (a_datatype_definition_enumeration_ref *A_DATATYPE_DEFINITION_ENUMERATION_REF) UnstageVoid(stage *StageStruct) {
	delete(stage.A_DATATYPE_DEFINITION_ENUMERATION_REFs, a_datatype_definition_enumeration_ref)
	delete(stage.A_DATATYPE_DEFINITION_ENUMERATION_REFs_mapString, a_datatype_definition_enumeration_ref.Name)
}

// commit a_datatype_definition_enumeration_ref to the back repo (if it is already staged)
func (a_datatype_definition_enumeration_ref *A_DATATYPE_DEFINITION_ENUMERATION_REF) Commit(stage *StageStruct) *A_DATATYPE_DEFINITION_ENUMERATION_REF {
	if _, ok := stage.A_DATATYPE_DEFINITION_ENUMERATION_REFs[a_datatype_definition_enumeration_ref]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitA_DATATYPE_DEFINITION_ENUMERATION_REF(a_datatype_definition_enumeration_ref)
		}
	}
	return a_datatype_definition_enumeration_ref
}

func (a_datatype_definition_enumeration_ref *A_DATATYPE_DEFINITION_ENUMERATION_REF) CommitVoid(stage *StageStruct) {
	a_datatype_definition_enumeration_ref.Commit(stage)
}

// Checkout a_datatype_definition_enumeration_ref to the back repo (if it is already staged)
func (a_datatype_definition_enumeration_ref *A_DATATYPE_DEFINITION_ENUMERATION_REF) Checkout(stage *StageStruct) *A_DATATYPE_DEFINITION_ENUMERATION_REF {
	if _, ok := stage.A_DATATYPE_DEFINITION_ENUMERATION_REFs[a_datatype_definition_enumeration_ref]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutA_DATATYPE_DEFINITION_ENUMERATION_REF(a_datatype_definition_enumeration_ref)
		}
	}
	return a_datatype_definition_enumeration_ref
}

// for satisfaction of GongStruct interface
func (a_datatype_definition_enumeration_ref *A_DATATYPE_DEFINITION_ENUMERATION_REF) GetName() (res string) {
	return a_datatype_definition_enumeration_ref.Name
}

// Stage puts a_datatype_definition_integer_ref to the model stage
func (a_datatype_definition_integer_ref *A_DATATYPE_DEFINITION_INTEGER_REF) Stage(stage *StageStruct) *A_DATATYPE_DEFINITION_INTEGER_REF {
	stage.A_DATATYPE_DEFINITION_INTEGER_REFs[a_datatype_definition_integer_ref] = __member
	stage.A_DATATYPE_DEFINITION_INTEGER_REFs_mapString[a_datatype_definition_integer_ref.Name] = a_datatype_definition_integer_ref

	return a_datatype_definition_integer_ref
}

// Unstage removes a_datatype_definition_integer_ref off the model stage
func (a_datatype_definition_integer_ref *A_DATATYPE_DEFINITION_INTEGER_REF) Unstage(stage *StageStruct) *A_DATATYPE_DEFINITION_INTEGER_REF {
	delete(stage.A_DATATYPE_DEFINITION_INTEGER_REFs, a_datatype_definition_integer_ref)
	delete(stage.A_DATATYPE_DEFINITION_INTEGER_REFs_mapString, a_datatype_definition_integer_ref.Name)
	return a_datatype_definition_integer_ref
}

// UnstageVoid removes a_datatype_definition_integer_ref off the model stage
func (a_datatype_definition_integer_ref *A_DATATYPE_DEFINITION_INTEGER_REF) UnstageVoid(stage *StageStruct) {
	delete(stage.A_DATATYPE_DEFINITION_INTEGER_REFs, a_datatype_definition_integer_ref)
	delete(stage.A_DATATYPE_DEFINITION_INTEGER_REFs_mapString, a_datatype_definition_integer_ref.Name)
}

// commit a_datatype_definition_integer_ref to the back repo (if it is already staged)
func (a_datatype_definition_integer_ref *A_DATATYPE_DEFINITION_INTEGER_REF) Commit(stage *StageStruct) *A_DATATYPE_DEFINITION_INTEGER_REF {
	if _, ok := stage.A_DATATYPE_DEFINITION_INTEGER_REFs[a_datatype_definition_integer_ref]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitA_DATATYPE_DEFINITION_INTEGER_REF(a_datatype_definition_integer_ref)
		}
	}
	return a_datatype_definition_integer_ref
}

func (a_datatype_definition_integer_ref *A_DATATYPE_DEFINITION_INTEGER_REF) CommitVoid(stage *StageStruct) {
	a_datatype_definition_integer_ref.Commit(stage)
}

// Checkout a_datatype_definition_integer_ref to the back repo (if it is already staged)
func (a_datatype_definition_integer_ref *A_DATATYPE_DEFINITION_INTEGER_REF) Checkout(stage *StageStruct) *A_DATATYPE_DEFINITION_INTEGER_REF {
	if _, ok := stage.A_DATATYPE_DEFINITION_INTEGER_REFs[a_datatype_definition_integer_ref]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutA_DATATYPE_DEFINITION_INTEGER_REF(a_datatype_definition_integer_ref)
		}
	}
	return a_datatype_definition_integer_ref
}

// for satisfaction of GongStruct interface
func (a_datatype_definition_integer_ref *A_DATATYPE_DEFINITION_INTEGER_REF) GetName() (res string) {
	return a_datatype_definition_integer_ref.Name
}

// Stage puts a_datatype_definition_real_ref to the model stage
func (a_datatype_definition_real_ref *A_DATATYPE_DEFINITION_REAL_REF) Stage(stage *StageStruct) *A_DATATYPE_DEFINITION_REAL_REF {
	stage.A_DATATYPE_DEFINITION_REAL_REFs[a_datatype_definition_real_ref] = __member
	stage.A_DATATYPE_DEFINITION_REAL_REFs_mapString[a_datatype_definition_real_ref.Name] = a_datatype_definition_real_ref

	return a_datatype_definition_real_ref
}

// Unstage removes a_datatype_definition_real_ref off the model stage
func (a_datatype_definition_real_ref *A_DATATYPE_DEFINITION_REAL_REF) Unstage(stage *StageStruct) *A_DATATYPE_DEFINITION_REAL_REF {
	delete(stage.A_DATATYPE_DEFINITION_REAL_REFs, a_datatype_definition_real_ref)
	delete(stage.A_DATATYPE_DEFINITION_REAL_REFs_mapString, a_datatype_definition_real_ref.Name)
	return a_datatype_definition_real_ref
}

// UnstageVoid removes a_datatype_definition_real_ref off the model stage
func (a_datatype_definition_real_ref *A_DATATYPE_DEFINITION_REAL_REF) UnstageVoid(stage *StageStruct) {
	delete(stage.A_DATATYPE_DEFINITION_REAL_REFs, a_datatype_definition_real_ref)
	delete(stage.A_DATATYPE_DEFINITION_REAL_REFs_mapString, a_datatype_definition_real_ref.Name)
}

// commit a_datatype_definition_real_ref to the back repo (if it is already staged)
func (a_datatype_definition_real_ref *A_DATATYPE_DEFINITION_REAL_REF) Commit(stage *StageStruct) *A_DATATYPE_DEFINITION_REAL_REF {
	if _, ok := stage.A_DATATYPE_DEFINITION_REAL_REFs[a_datatype_definition_real_ref]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitA_DATATYPE_DEFINITION_REAL_REF(a_datatype_definition_real_ref)
		}
	}
	return a_datatype_definition_real_ref
}

func (a_datatype_definition_real_ref *A_DATATYPE_DEFINITION_REAL_REF) CommitVoid(stage *StageStruct) {
	a_datatype_definition_real_ref.Commit(stage)
}

// Checkout a_datatype_definition_real_ref to the back repo (if it is already staged)
func (a_datatype_definition_real_ref *A_DATATYPE_DEFINITION_REAL_REF) Checkout(stage *StageStruct) *A_DATATYPE_DEFINITION_REAL_REF {
	if _, ok := stage.A_DATATYPE_DEFINITION_REAL_REFs[a_datatype_definition_real_ref]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutA_DATATYPE_DEFINITION_REAL_REF(a_datatype_definition_real_ref)
		}
	}
	return a_datatype_definition_real_ref
}

// for satisfaction of GongStruct interface
func (a_datatype_definition_real_ref *A_DATATYPE_DEFINITION_REAL_REF) GetName() (res string) {
	return a_datatype_definition_real_ref.Name
}

// Stage puts a_datatype_definition_string_ref to the model stage
func (a_datatype_definition_string_ref *A_DATATYPE_DEFINITION_STRING_REF) Stage(stage *StageStruct) *A_DATATYPE_DEFINITION_STRING_REF {
	stage.A_DATATYPE_DEFINITION_STRING_REFs[a_datatype_definition_string_ref] = __member
	stage.A_DATATYPE_DEFINITION_STRING_REFs_mapString[a_datatype_definition_string_ref.Name] = a_datatype_definition_string_ref

	return a_datatype_definition_string_ref
}

// Unstage removes a_datatype_definition_string_ref off the model stage
func (a_datatype_definition_string_ref *A_DATATYPE_DEFINITION_STRING_REF) Unstage(stage *StageStruct) *A_DATATYPE_DEFINITION_STRING_REF {
	delete(stage.A_DATATYPE_DEFINITION_STRING_REFs, a_datatype_definition_string_ref)
	delete(stage.A_DATATYPE_DEFINITION_STRING_REFs_mapString, a_datatype_definition_string_ref.Name)
	return a_datatype_definition_string_ref
}

// UnstageVoid removes a_datatype_definition_string_ref off the model stage
func (a_datatype_definition_string_ref *A_DATATYPE_DEFINITION_STRING_REF) UnstageVoid(stage *StageStruct) {
	delete(stage.A_DATATYPE_DEFINITION_STRING_REFs, a_datatype_definition_string_ref)
	delete(stage.A_DATATYPE_DEFINITION_STRING_REFs_mapString, a_datatype_definition_string_ref.Name)
}

// commit a_datatype_definition_string_ref to the back repo (if it is already staged)
func (a_datatype_definition_string_ref *A_DATATYPE_DEFINITION_STRING_REF) Commit(stage *StageStruct) *A_DATATYPE_DEFINITION_STRING_REF {
	if _, ok := stage.A_DATATYPE_DEFINITION_STRING_REFs[a_datatype_definition_string_ref]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitA_DATATYPE_DEFINITION_STRING_REF(a_datatype_definition_string_ref)
		}
	}
	return a_datatype_definition_string_ref
}

func (a_datatype_definition_string_ref *A_DATATYPE_DEFINITION_STRING_REF) CommitVoid(stage *StageStruct) {
	a_datatype_definition_string_ref.Commit(stage)
}

// Checkout a_datatype_definition_string_ref to the back repo (if it is already staged)
func (a_datatype_definition_string_ref *A_DATATYPE_DEFINITION_STRING_REF) Checkout(stage *StageStruct) *A_DATATYPE_DEFINITION_STRING_REF {
	if _, ok := stage.A_DATATYPE_DEFINITION_STRING_REFs[a_datatype_definition_string_ref]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutA_DATATYPE_DEFINITION_STRING_REF(a_datatype_definition_string_ref)
		}
	}
	return a_datatype_definition_string_ref
}

// for satisfaction of GongStruct interface
func (a_datatype_definition_string_ref *A_DATATYPE_DEFINITION_STRING_REF) GetName() (res string) {
	return a_datatype_definition_string_ref.Name
}

// Stage puts a_datatype_definition_xhtml_ref to the model stage
func (a_datatype_definition_xhtml_ref *A_DATATYPE_DEFINITION_XHTML_REF) Stage(stage *StageStruct) *A_DATATYPE_DEFINITION_XHTML_REF {
	stage.A_DATATYPE_DEFINITION_XHTML_REFs[a_datatype_definition_xhtml_ref] = __member
	stage.A_DATATYPE_DEFINITION_XHTML_REFs_mapString[a_datatype_definition_xhtml_ref.Name] = a_datatype_definition_xhtml_ref

	return a_datatype_definition_xhtml_ref
}

// Unstage removes a_datatype_definition_xhtml_ref off the model stage
func (a_datatype_definition_xhtml_ref *A_DATATYPE_DEFINITION_XHTML_REF) Unstage(stage *StageStruct) *A_DATATYPE_DEFINITION_XHTML_REF {
	delete(stage.A_DATATYPE_DEFINITION_XHTML_REFs, a_datatype_definition_xhtml_ref)
	delete(stage.A_DATATYPE_DEFINITION_XHTML_REFs_mapString, a_datatype_definition_xhtml_ref.Name)
	return a_datatype_definition_xhtml_ref
}

// UnstageVoid removes a_datatype_definition_xhtml_ref off the model stage
func (a_datatype_definition_xhtml_ref *A_DATATYPE_DEFINITION_XHTML_REF) UnstageVoid(stage *StageStruct) {
	delete(stage.A_DATATYPE_DEFINITION_XHTML_REFs, a_datatype_definition_xhtml_ref)
	delete(stage.A_DATATYPE_DEFINITION_XHTML_REFs_mapString, a_datatype_definition_xhtml_ref.Name)
}

// commit a_datatype_definition_xhtml_ref to the back repo (if it is already staged)
func (a_datatype_definition_xhtml_ref *A_DATATYPE_DEFINITION_XHTML_REF) Commit(stage *StageStruct) *A_DATATYPE_DEFINITION_XHTML_REF {
	if _, ok := stage.A_DATATYPE_DEFINITION_XHTML_REFs[a_datatype_definition_xhtml_ref]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitA_DATATYPE_DEFINITION_XHTML_REF(a_datatype_definition_xhtml_ref)
		}
	}
	return a_datatype_definition_xhtml_ref
}

func (a_datatype_definition_xhtml_ref *A_DATATYPE_DEFINITION_XHTML_REF) CommitVoid(stage *StageStruct) {
	a_datatype_definition_xhtml_ref.Commit(stage)
}

// Checkout a_datatype_definition_xhtml_ref to the back repo (if it is already staged)
func (a_datatype_definition_xhtml_ref *A_DATATYPE_DEFINITION_XHTML_REF) Checkout(stage *StageStruct) *A_DATATYPE_DEFINITION_XHTML_REF {
	if _, ok := stage.A_DATATYPE_DEFINITION_XHTML_REFs[a_datatype_definition_xhtml_ref]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutA_DATATYPE_DEFINITION_XHTML_REF(a_datatype_definition_xhtml_ref)
		}
	}
	return a_datatype_definition_xhtml_ref
}

// for satisfaction of GongStruct interface
func (a_datatype_definition_xhtml_ref *A_DATATYPE_DEFINITION_XHTML_REF) GetName() (res string) {
	return a_datatype_definition_xhtml_ref.Name
}

// Stage puts a_editable_atts to the model stage
func (a_editable_atts *A_EDITABLE_ATTS) Stage(stage *StageStruct) *A_EDITABLE_ATTS {
	stage.A_EDITABLE_ATTSs[a_editable_atts] = __member
	stage.A_EDITABLE_ATTSs_mapString[a_editable_atts.Name] = a_editable_atts

	return a_editable_atts
}

// Unstage removes a_editable_atts off the model stage
func (a_editable_atts *A_EDITABLE_ATTS) Unstage(stage *StageStruct) *A_EDITABLE_ATTS {
	delete(stage.A_EDITABLE_ATTSs, a_editable_atts)
	delete(stage.A_EDITABLE_ATTSs_mapString, a_editable_atts.Name)
	return a_editable_atts
}

// UnstageVoid removes a_editable_atts off the model stage
func (a_editable_atts *A_EDITABLE_ATTS) UnstageVoid(stage *StageStruct) {
	delete(stage.A_EDITABLE_ATTSs, a_editable_atts)
	delete(stage.A_EDITABLE_ATTSs_mapString, a_editable_atts.Name)
}

// commit a_editable_atts to the back repo (if it is already staged)
func (a_editable_atts *A_EDITABLE_ATTS) Commit(stage *StageStruct) *A_EDITABLE_ATTS {
	if _, ok := stage.A_EDITABLE_ATTSs[a_editable_atts]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitA_EDITABLE_ATTS(a_editable_atts)
		}
	}
	return a_editable_atts
}

func (a_editable_atts *A_EDITABLE_ATTS) CommitVoid(stage *StageStruct) {
	a_editable_atts.Commit(stage)
}

// Checkout a_editable_atts to the back repo (if it is already staged)
func (a_editable_atts *A_EDITABLE_ATTS) Checkout(stage *StageStruct) *A_EDITABLE_ATTS {
	if _, ok := stage.A_EDITABLE_ATTSs[a_editable_atts]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutA_EDITABLE_ATTS(a_editable_atts)
		}
	}
	return a_editable_atts
}

// for satisfaction of GongStruct interface
func (a_editable_atts *A_EDITABLE_ATTS) GetName() (res string) {
	return a_editable_atts.Name
}

// Stage puts a_enum_value_ref to the model stage
func (a_enum_value_ref *A_ENUM_VALUE_REF) Stage(stage *StageStruct) *A_ENUM_VALUE_REF {
	stage.A_ENUM_VALUE_REFs[a_enum_value_ref] = __member
	stage.A_ENUM_VALUE_REFs_mapString[a_enum_value_ref.Name] = a_enum_value_ref

	return a_enum_value_ref
}

// Unstage removes a_enum_value_ref off the model stage
func (a_enum_value_ref *A_ENUM_VALUE_REF) Unstage(stage *StageStruct) *A_ENUM_VALUE_REF {
	delete(stage.A_ENUM_VALUE_REFs, a_enum_value_ref)
	delete(stage.A_ENUM_VALUE_REFs_mapString, a_enum_value_ref.Name)
	return a_enum_value_ref
}

// UnstageVoid removes a_enum_value_ref off the model stage
func (a_enum_value_ref *A_ENUM_VALUE_REF) UnstageVoid(stage *StageStruct) {
	delete(stage.A_ENUM_VALUE_REFs, a_enum_value_ref)
	delete(stage.A_ENUM_VALUE_REFs_mapString, a_enum_value_ref.Name)
}

// commit a_enum_value_ref to the back repo (if it is already staged)
func (a_enum_value_ref *A_ENUM_VALUE_REF) Commit(stage *StageStruct) *A_ENUM_VALUE_REF {
	if _, ok := stage.A_ENUM_VALUE_REFs[a_enum_value_ref]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitA_ENUM_VALUE_REF(a_enum_value_ref)
		}
	}
	return a_enum_value_ref
}

func (a_enum_value_ref *A_ENUM_VALUE_REF) CommitVoid(stage *StageStruct) {
	a_enum_value_ref.Commit(stage)
}

// Checkout a_enum_value_ref to the back repo (if it is already staged)
func (a_enum_value_ref *A_ENUM_VALUE_REF) Checkout(stage *StageStruct) *A_ENUM_VALUE_REF {
	if _, ok := stage.A_ENUM_VALUE_REFs[a_enum_value_ref]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutA_ENUM_VALUE_REF(a_enum_value_ref)
		}
	}
	return a_enum_value_ref
}

// for satisfaction of GongStruct interface
func (a_enum_value_ref *A_ENUM_VALUE_REF) GetName() (res string) {
	return a_enum_value_ref.Name
}

// Stage puts a_object to the model stage
func (a_object *A_OBJECT) Stage(stage *StageStruct) *A_OBJECT {
	stage.A_OBJECTs[a_object] = __member
	stage.A_OBJECTs_mapString[a_object.Name] = a_object

	return a_object
}

// Unstage removes a_object off the model stage
func (a_object *A_OBJECT) Unstage(stage *StageStruct) *A_OBJECT {
	delete(stage.A_OBJECTs, a_object)
	delete(stage.A_OBJECTs_mapString, a_object.Name)
	return a_object
}

// UnstageVoid removes a_object off the model stage
func (a_object *A_OBJECT) UnstageVoid(stage *StageStruct) {
	delete(stage.A_OBJECTs, a_object)
	delete(stage.A_OBJECTs_mapString, a_object.Name)
}

// commit a_object to the back repo (if it is already staged)
func (a_object *A_OBJECT) Commit(stage *StageStruct) *A_OBJECT {
	if _, ok := stage.A_OBJECTs[a_object]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitA_OBJECT(a_object)
		}
	}
	return a_object
}

func (a_object *A_OBJECT) CommitVoid(stage *StageStruct) {
	a_object.Commit(stage)
}

// Checkout a_object to the back repo (if it is already staged)
func (a_object *A_OBJECT) Checkout(stage *StageStruct) *A_OBJECT {
	if _, ok := stage.A_OBJECTs[a_object]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutA_OBJECT(a_object)
		}
	}
	return a_object
}

// for satisfaction of GongStruct interface
func (a_object *A_OBJECT) GetName() (res string) {
	return a_object.Name
}

// Stage puts a_properties to the model stage
func (a_properties *A_PROPERTIES) Stage(stage *StageStruct) *A_PROPERTIES {
	stage.A_PROPERTIESs[a_properties] = __member
	stage.A_PROPERTIESs_mapString[a_properties.Name] = a_properties

	return a_properties
}

// Unstage removes a_properties off the model stage
func (a_properties *A_PROPERTIES) Unstage(stage *StageStruct) *A_PROPERTIES {
	delete(stage.A_PROPERTIESs, a_properties)
	delete(stage.A_PROPERTIESs_mapString, a_properties.Name)
	return a_properties
}

// UnstageVoid removes a_properties off the model stage
func (a_properties *A_PROPERTIES) UnstageVoid(stage *StageStruct) {
	delete(stage.A_PROPERTIESs, a_properties)
	delete(stage.A_PROPERTIESs_mapString, a_properties.Name)
}

// commit a_properties to the back repo (if it is already staged)
func (a_properties *A_PROPERTIES) Commit(stage *StageStruct) *A_PROPERTIES {
	if _, ok := stage.A_PROPERTIESs[a_properties]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitA_PROPERTIES(a_properties)
		}
	}
	return a_properties
}

func (a_properties *A_PROPERTIES) CommitVoid(stage *StageStruct) {
	a_properties.Commit(stage)
}

// Checkout a_properties to the back repo (if it is already staged)
func (a_properties *A_PROPERTIES) Checkout(stage *StageStruct) *A_PROPERTIES {
	if _, ok := stage.A_PROPERTIESs[a_properties]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutA_PROPERTIES(a_properties)
		}
	}
	return a_properties
}

// for satisfaction of GongStruct interface
func (a_properties *A_PROPERTIES) GetName() (res string) {
	return a_properties.Name
}

// Stage puts a_relation_group_type_ref to the model stage
func (a_relation_group_type_ref *A_RELATION_GROUP_TYPE_REF) Stage(stage *StageStruct) *A_RELATION_GROUP_TYPE_REF {
	stage.A_RELATION_GROUP_TYPE_REFs[a_relation_group_type_ref] = __member
	stage.A_RELATION_GROUP_TYPE_REFs_mapString[a_relation_group_type_ref.Name] = a_relation_group_type_ref

	return a_relation_group_type_ref
}

// Unstage removes a_relation_group_type_ref off the model stage
func (a_relation_group_type_ref *A_RELATION_GROUP_TYPE_REF) Unstage(stage *StageStruct) *A_RELATION_GROUP_TYPE_REF {
	delete(stage.A_RELATION_GROUP_TYPE_REFs, a_relation_group_type_ref)
	delete(stage.A_RELATION_GROUP_TYPE_REFs_mapString, a_relation_group_type_ref.Name)
	return a_relation_group_type_ref
}

// UnstageVoid removes a_relation_group_type_ref off the model stage
func (a_relation_group_type_ref *A_RELATION_GROUP_TYPE_REF) UnstageVoid(stage *StageStruct) {
	delete(stage.A_RELATION_GROUP_TYPE_REFs, a_relation_group_type_ref)
	delete(stage.A_RELATION_GROUP_TYPE_REFs_mapString, a_relation_group_type_ref.Name)
}

// commit a_relation_group_type_ref to the back repo (if it is already staged)
func (a_relation_group_type_ref *A_RELATION_GROUP_TYPE_REF) Commit(stage *StageStruct) *A_RELATION_GROUP_TYPE_REF {
	if _, ok := stage.A_RELATION_GROUP_TYPE_REFs[a_relation_group_type_ref]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitA_RELATION_GROUP_TYPE_REF(a_relation_group_type_ref)
		}
	}
	return a_relation_group_type_ref
}

func (a_relation_group_type_ref *A_RELATION_GROUP_TYPE_REF) CommitVoid(stage *StageStruct) {
	a_relation_group_type_ref.Commit(stage)
}

// Checkout a_relation_group_type_ref to the back repo (if it is already staged)
func (a_relation_group_type_ref *A_RELATION_GROUP_TYPE_REF) Checkout(stage *StageStruct) *A_RELATION_GROUP_TYPE_REF {
	if _, ok := stage.A_RELATION_GROUP_TYPE_REFs[a_relation_group_type_ref]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutA_RELATION_GROUP_TYPE_REF(a_relation_group_type_ref)
		}
	}
	return a_relation_group_type_ref
}

// for satisfaction of GongStruct interface
func (a_relation_group_type_ref *A_RELATION_GROUP_TYPE_REF) GetName() (res string) {
	return a_relation_group_type_ref.Name
}

// Stage puts a_source_1 to the model stage
func (a_source_1 *A_SOURCE_1) Stage(stage *StageStruct) *A_SOURCE_1 {
	stage.A_SOURCE_1s[a_source_1] = __member
	stage.A_SOURCE_1s_mapString[a_source_1.Name] = a_source_1

	return a_source_1
}

// Unstage removes a_source_1 off the model stage
func (a_source_1 *A_SOURCE_1) Unstage(stage *StageStruct) *A_SOURCE_1 {
	delete(stage.A_SOURCE_1s, a_source_1)
	delete(stage.A_SOURCE_1s_mapString, a_source_1.Name)
	return a_source_1
}

// UnstageVoid removes a_source_1 off the model stage
func (a_source_1 *A_SOURCE_1) UnstageVoid(stage *StageStruct) {
	delete(stage.A_SOURCE_1s, a_source_1)
	delete(stage.A_SOURCE_1s_mapString, a_source_1.Name)
}

// commit a_source_1 to the back repo (if it is already staged)
func (a_source_1 *A_SOURCE_1) Commit(stage *StageStruct) *A_SOURCE_1 {
	if _, ok := stage.A_SOURCE_1s[a_source_1]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitA_SOURCE_1(a_source_1)
		}
	}
	return a_source_1
}

func (a_source_1 *A_SOURCE_1) CommitVoid(stage *StageStruct) {
	a_source_1.Commit(stage)
}

// Checkout a_source_1 to the back repo (if it is already staged)
func (a_source_1 *A_SOURCE_1) Checkout(stage *StageStruct) *A_SOURCE_1 {
	if _, ok := stage.A_SOURCE_1s[a_source_1]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutA_SOURCE_1(a_source_1)
		}
	}
	return a_source_1
}

// for satisfaction of GongStruct interface
func (a_source_1 *A_SOURCE_1) GetName() (res string) {
	return a_source_1.Name
}

// Stage puts a_specifications to the model stage
func (a_specifications *A_SPECIFICATIONS) Stage(stage *StageStruct) *A_SPECIFICATIONS {
	stage.A_SPECIFICATIONSs[a_specifications] = __member
	stage.A_SPECIFICATIONSs_mapString[a_specifications.Name] = a_specifications

	return a_specifications
}

// Unstage removes a_specifications off the model stage
func (a_specifications *A_SPECIFICATIONS) Unstage(stage *StageStruct) *A_SPECIFICATIONS {
	delete(stage.A_SPECIFICATIONSs, a_specifications)
	delete(stage.A_SPECIFICATIONSs_mapString, a_specifications.Name)
	return a_specifications
}

// UnstageVoid removes a_specifications off the model stage
func (a_specifications *A_SPECIFICATIONS) UnstageVoid(stage *StageStruct) {
	delete(stage.A_SPECIFICATIONSs, a_specifications)
	delete(stage.A_SPECIFICATIONSs_mapString, a_specifications.Name)
}

// commit a_specifications to the back repo (if it is already staged)
func (a_specifications *A_SPECIFICATIONS) Commit(stage *StageStruct) *A_SPECIFICATIONS {
	if _, ok := stage.A_SPECIFICATIONSs[a_specifications]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitA_SPECIFICATIONS(a_specifications)
		}
	}
	return a_specifications
}

func (a_specifications *A_SPECIFICATIONS) CommitVoid(stage *StageStruct) {
	a_specifications.Commit(stage)
}

// Checkout a_specifications to the back repo (if it is already staged)
func (a_specifications *A_SPECIFICATIONS) Checkout(stage *StageStruct) *A_SPECIFICATIONS {
	if _, ok := stage.A_SPECIFICATIONSs[a_specifications]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutA_SPECIFICATIONS(a_specifications)
		}
	}
	return a_specifications
}

// for satisfaction of GongStruct interface
func (a_specifications *A_SPECIFICATIONS) GetName() (res string) {
	return a_specifications.Name
}

// Stage puts a_specification_type_ref to the model stage
func (a_specification_type_ref *A_SPECIFICATION_TYPE_REF) Stage(stage *StageStruct) *A_SPECIFICATION_TYPE_REF {
	stage.A_SPECIFICATION_TYPE_REFs[a_specification_type_ref] = __member
	stage.A_SPECIFICATION_TYPE_REFs_mapString[a_specification_type_ref.Name] = a_specification_type_ref

	return a_specification_type_ref
}

// Unstage removes a_specification_type_ref off the model stage
func (a_specification_type_ref *A_SPECIFICATION_TYPE_REF) Unstage(stage *StageStruct) *A_SPECIFICATION_TYPE_REF {
	delete(stage.A_SPECIFICATION_TYPE_REFs, a_specification_type_ref)
	delete(stage.A_SPECIFICATION_TYPE_REFs_mapString, a_specification_type_ref.Name)
	return a_specification_type_ref
}

// UnstageVoid removes a_specification_type_ref off the model stage
func (a_specification_type_ref *A_SPECIFICATION_TYPE_REF) UnstageVoid(stage *StageStruct) {
	delete(stage.A_SPECIFICATION_TYPE_REFs, a_specification_type_ref)
	delete(stage.A_SPECIFICATION_TYPE_REFs_mapString, a_specification_type_ref.Name)
}

// commit a_specification_type_ref to the back repo (if it is already staged)
func (a_specification_type_ref *A_SPECIFICATION_TYPE_REF) Commit(stage *StageStruct) *A_SPECIFICATION_TYPE_REF {
	if _, ok := stage.A_SPECIFICATION_TYPE_REFs[a_specification_type_ref]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitA_SPECIFICATION_TYPE_REF(a_specification_type_ref)
		}
	}
	return a_specification_type_ref
}

func (a_specification_type_ref *A_SPECIFICATION_TYPE_REF) CommitVoid(stage *StageStruct) {
	a_specification_type_ref.Commit(stage)
}

// Checkout a_specification_type_ref to the back repo (if it is already staged)
func (a_specification_type_ref *A_SPECIFICATION_TYPE_REF) Checkout(stage *StageStruct) *A_SPECIFICATION_TYPE_REF {
	if _, ok := stage.A_SPECIFICATION_TYPE_REFs[a_specification_type_ref]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutA_SPECIFICATION_TYPE_REF(a_specification_type_ref)
		}
	}
	return a_specification_type_ref
}

// for satisfaction of GongStruct interface
func (a_specification_type_ref *A_SPECIFICATION_TYPE_REF) GetName() (res string) {
	return a_specification_type_ref.Name
}

// Stage puts a_specified_values to the model stage
func (a_specified_values *A_SPECIFIED_VALUES) Stage(stage *StageStruct) *A_SPECIFIED_VALUES {
	stage.A_SPECIFIED_VALUESs[a_specified_values] = __member
	stage.A_SPECIFIED_VALUESs_mapString[a_specified_values.Name] = a_specified_values

	return a_specified_values
}

// Unstage removes a_specified_values off the model stage
func (a_specified_values *A_SPECIFIED_VALUES) Unstage(stage *StageStruct) *A_SPECIFIED_VALUES {
	delete(stage.A_SPECIFIED_VALUESs, a_specified_values)
	delete(stage.A_SPECIFIED_VALUESs_mapString, a_specified_values.Name)
	return a_specified_values
}

// UnstageVoid removes a_specified_values off the model stage
func (a_specified_values *A_SPECIFIED_VALUES) UnstageVoid(stage *StageStruct) {
	delete(stage.A_SPECIFIED_VALUESs, a_specified_values)
	delete(stage.A_SPECIFIED_VALUESs_mapString, a_specified_values.Name)
}

// commit a_specified_values to the back repo (if it is already staged)
func (a_specified_values *A_SPECIFIED_VALUES) Commit(stage *StageStruct) *A_SPECIFIED_VALUES {
	if _, ok := stage.A_SPECIFIED_VALUESs[a_specified_values]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitA_SPECIFIED_VALUES(a_specified_values)
		}
	}
	return a_specified_values
}

func (a_specified_values *A_SPECIFIED_VALUES) CommitVoid(stage *StageStruct) {
	a_specified_values.Commit(stage)
}

// Checkout a_specified_values to the back repo (if it is already staged)
func (a_specified_values *A_SPECIFIED_VALUES) Checkout(stage *StageStruct) *A_SPECIFIED_VALUES {
	if _, ok := stage.A_SPECIFIED_VALUESs[a_specified_values]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutA_SPECIFIED_VALUES(a_specified_values)
		}
	}
	return a_specified_values
}

// for satisfaction of GongStruct interface
func (a_specified_values *A_SPECIFIED_VALUES) GetName() (res string) {
	return a_specified_values.Name
}

// Stage puts a_spec_attributes to the model stage
func (a_spec_attributes *A_SPEC_ATTRIBUTES) Stage(stage *StageStruct) *A_SPEC_ATTRIBUTES {
	stage.A_SPEC_ATTRIBUTESs[a_spec_attributes] = __member
	stage.A_SPEC_ATTRIBUTESs_mapString[a_spec_attributes.Name] = a_spec_attributes

	return a_spec_attributes
}

// Unstage removes a_spec_attributes off the model stage
func (a_spec_attributes *A_SPEC_ATTRIBUTES) Unstage(stage *StageStruct) *A_SPEC_ATTRIBUTES {
	delete(stage.A_SPEC_ATTRIBUTESs, a_spec_attributes)
	delete(stage.A_SPEC_ATTRIBUTESs_mapString, a_spec_attributes.Name)
	return a_spec_attributes
}

// UnstageVoid removes a_spec_attributes off the model stage
func (a_spec_attributes *A_SPEC_ATTRIBUTES) UnstageVoid(stage *StageStruct) {
	delete(stage.A_SPEC_ATTRIBUTESs, a_spec_attributes)
	delete(stage.A_SPEC_ATTRIBUTESs_mapString, a_spec_attributes.Name)
}

// commit a_spec_attributes to the back repo (if it is already staged)
func (a_spec_attributes *A_SPEC_ATTRIBUTES) Commit(stage *StageStruct) *A_SPEC_ATTRIBUTES {
	if _, ok := stage.A_SPEC_ATTRIBUTESs[a_spec_attributes]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitA_SPEC_ATTRIBUTES(a_spec_attributes)
		}
	}
	return a_spec_attributes
}

func (a_spec_attributes *A_SPEC_ATTRIBUTES) CommitVoid(stage *StageStruct) {
	a_spec_attributes.Commit(stage)
}

// Checkout a_spec_attributes to the back repo (if it is already staged)
func (a_spec_attributes *A_SPEC_ATTRIBUTES) Checkout(stage *StageStruct) *A_SPEC_ATTRIBUTES {
	if _, ok := stage.A_SPEC_ATTRIBUTESs[a_spec_attributes]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutA_SPEC_ATTRIBUTES(a_spec_attributes)
		}
	}
	return a_spec_attributes
}

// for satisfaction of GongStruct interface
func (a_spec_attributes *A_SPEC_ATTRIBUTES) GetName() (res string) {
	return a_spec_attributes.Name
}

// Stage puts a_spec_objects to the model stage
func (a_spec_objects *A_SPEC_OBJECTS) Stage(stage *StageStruct) *A_SPEC_OBJECTS {
	stage.A_SPEC_OBJECTSs[a_spec_objects] = __member
	stage.A_SPEC_OBJECTSs_mapString[a_spec_objects.Name] = a_spec_objects

	return a_spec_objects
}

// Unstage removes a_spec_objects off the model stage
func (a_spec_objects *A_SPEC_OBJECTS) Unstage(stage *StageStruct) *A_SPEC_OBJECTS {
	delete(stage.A_SPEC_OBJECTSs, a_spec_objects)
	delete(stage.A_SPEC_OBJECTSs_mapString, a_spec_objects.Name)
	return a_spec_objects
}

// UnstageVoid removes a_spec_objects off the model stage
func (a_spec_objects *A_SPEC_OBJECTS) UnstageVoid(stage *StageStruct) {
	delete(stage.A_SPEC_OBJECTSs, a_spec_objects)
	delete(stage.A_SPEC_OBJECTSs_mapString, a_spec_objects.Name)
}

// commit a_spec_objects to the back repo (if it is already staged)
func (a_spec_objects *A_SPEC_OBJECTS) Commit(stage *StageStruct) *A_SPEC_OBJECTS {
	if _, ok := stage.A_SPEC_OBJECTSs[a_spec_objects]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitA_SPEC_OBJECTS(a_spec_objects)
		}
	}
	return a_spec_objects
}

func (a_spec_objects *A_SPEC_OBJECTS) CommitVoid(stage *StageStruct) {
	a_spec_objects.Commit(stage)
}

// Checkout a_spec_objects to the back repo (if it is already staged)
func (a_spec_objects *A_SPEC_OBJECTS) Checkout(stage *StageStruct) *A_SPEC_OBJECTS {
	if _, ok := stage.A_SPEC_OBJECTSs[a_spec_objects]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutA_SPEC_OBJECTS(a_spec_objects)
		}
	}
	return a_spec_objects
}

// for satisfaction of GongStruct interface
func (a_spec_objects *A_SPEC_OBJECTS) GetName() (res string) {
	return a_spec_objects.Name
}

// Stage puts a_spec_object_type_ref to the model stage
func (a_spec_object_type_ref *A_SPEC_OBJECT_TYPE_REF) Stage(stage *StageStruct) *A_SPEC_OBJECT_TYPE_REF {
	stage.A_SPEC_OBJECT_TYPE_REFs[a_spec_object_type_ref] = __member
	stage.A_SPEC_OBJECT_TYPE_REFs_mapString[a_spec_object_type_ref.Name] = a_spec_object_type_ref

	return a_spec_object_type_ref
}

// Unstage removes a_spec_object_type_ref off the model stage
func (a_spec_object_type_ref *A_SPEC_OBJECT_TYPE_REF) Unstage(stage *StageStruct) *A_SPEC_OBJECT_TYPE_REF {
	delete(stage.A_SPEC_OBJECT_TYPE_REFs, a_spec_object_type_ref)
	delete(stage.A_SPEC_OBJECT_TYPE_REFs_mapString, a_spec_object_type_ref.Name)
	return a_spec_object_type_ref
}

// UnstageVoid removes a_spec_object_type_ref off the model stage
func (a_spec_object_type_ref *A_SPEC_OBJECT_TYPE_REF) UnstageVoid(stage *StageStruct) {
	delete(stage.A_SPEC_OBJECT_TYPE_REFs, a_spec_object_type_ref)
	delete(stage.A_SPEC_OBJECT_TYPE_REFs_mapString, a_spec_object_type_ref.Name)
}

// commit a_spec_object_type_ref to the back repo (if it is already staged)
func (a_spec_object_type_ref *A_SPEC_OBJECT_TYPE_REF) Commit(stage *StageStruct) *A_SPEC_OBJECT_TYPE_REF {
	if _, ok := stage.A_SPEC_OBJECT_TYPE_REFs[a_spec_object_type_ref]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitA_SPEC_OBJECT_TYPE_REF(a_spec_object_type_ref)
		}
	}
	return a_spec_object_type_ref
}

func (a_spec_object_type_ref *A_SPEC_OBJECT_TYPE_REF) CommitVoid(stage *StageStruct) {
	a_spec_object_type_ref.Commit(stage)
}

// Checkout a_spec_object_type_ref to the back repo (if it is already staged)
func (a_spec_object_type_ref *A_SPEC_OBJECT_TYPE_REF) Checkout(stage *StageStruct) *A_SPEC_OBJECT_TYPE_REF {
	if _, ok := stage.A_SPEC_OBJECT_TYPE_REFs[a_spec_object_type_ref]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutA_SPEC_OBJECT_TYPE_REF(a_spec_object_type_ref)
		}
	}
	return a_spec_object_type_ref
}

// for satisfaction of GongStruct interface
func (a_spec_object_type_ref *A_SPEC_OBJECT_TYPE_REF) GetName() (res string) {
	return a_spec_object_type_ref.Name
}

// Stage puts a_spec_relations to the model stage
func (a_spec_relations *A_SPEC_RELATIONS) Stage(stage *StageStruct) *A_SPEC_RELATIONS {
	stage.A_SPEC_RELATIONSs[a_spec_relations] = __member
	stage.A_SPEC_RELATIONSs_mapString[a_spec_relations.Name] = a_spec_relations

	return a_spec_relations
}

// Unstage removes a_spec_relations off the model stage
func (a_spec_relations *A_SPEC_RELATIONS) Unstage(stage *StageStruct) *A_SPEC_RELATIONS {
	delete(stage.A_SPEC_RELATIONSs, a_spec_relations)
	delete(stage.A_SPEC_RELATIONSs_mapString, a_spec_relations.Name)
	return a_spec_relations
}

// UnstageVoid removes a_spec_relations off the model stage
func (a_spec_relations *A_SPEC_RELATIONS) UnstageVoid(stage *StageStruct) {
	delete(stage.A_SPEC_RELATIONSs, a_spec_relations)
	delete(stage.A_SPEC_RELATIONSs_mapString, a_spec_relations.Name)
}

// commit a_spec_relations to the back repo (if it is already staged)
func (a_spec_relations *A_SPEC_RELATIONS) Commit(stage *StageStruct) *A_SPEC_RELATIONS {
	if _, ok := stage.A_SPEC_RELATIONSs[a_spec_relations]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitA_SPEC_RELATIONS(a_spec_relations)
		}
	}
	return a_spec_relations
}

func (a_spec_relations *A_SPEC_RELATIONS) CommitVoid(stage *StageStruct) {
	a_spec_relations.Commit(stage)
}

// Checkout a_spec_relations to the back repo (if it is already staged)
func (a_spec_relations *A_SPEC_RELATIONS) Checkout(stage *StageStruct) *A_SPEC_RELATIONS {
	if _, ok := stage.A_SPEC_RELATIONSs[a_spec_relations]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutA_SPEC_RELATIONS(a_spec_relations)
		}
	}
	return a_spec_relations
}

// for satisfaction of GongStruct interface
func (a_spec_relations *A_SPEC_RELATIONS) GetName() (res string) {
	return a_spec_relations.Name
}

// Stage puts a_spec_relation_groups to the model stage
func (a_spec_relation_groups *A_SPEC_RELATION_GROUPS) Stage(stage *StageStruct) *A_SPEC_RELATION_GROUPS {
	stage.A_SPEC_RELATION_GROUPSs[a_spec_relation_groups] = __member
	stage.A_SPEC_RELATION_GROUPSs_mapString[a_spec_relation_groups.Name] = a_spec_relation_groups

	return a_spec_relation_groups
}

// Unstage removes a_spec_relation_groups off the model stage
func (a_spec_relation_groups *A_SPEC_RELATION_GROUPS) Unstage(stage *StageStruct) *A_SPEC_RELATION_GROUPS {
	delete(stage.A_SPEC_RELATION_GROUPSs, a_spec_relation_groups)
	delete(stage.A_SPEC_RELATION_GROUPSs_mapString, a_spec_relation_groups.Name)
	return a_spec_relation_groups
}

// UnstageVoid removes a_spec_relation_groups off the model stage
func (a_spec_relation_groups *A_SPEC_RELATION_GROUPS) UnstageVoid(stage *StageStruct) {
	delete(stage.A_SPEC_RELATION_GROUPSs, a_spec_relation_groups)
	delete(stage.A_SPEC_RELATION_GROUPSs_mapString, a_spec_relation_groups.Name)
}

// commit a_spec_relation_groups to the back repo (if it is already staged)
func (a_spec_relation_groups *A_SPEC_RELATION_GROUPS) Commit(stage *StageStruct) *A_SPEC_RELATION_GROUPS {
	if _, ok := stage.A_SPEC_RELATION_GROUPSs[a_spec_relation_groups]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitA_SPEC_RELATION_GROUPS(a_spec_relation_groups)
		}
	}
	return a_spec_relation_groups
}

func (a_spec_relation_groups *A_SPEC_RELATION_GROUPS) CommitVoid(stage *StageStruct) {
	a_spec_relation_groups.Commit(stage)
}

// Checkout a_spec_relation_groups to the back repo (if it is already staged)
func (a_spec_relation_groups *A_SPEC_RELATION_GROUPS) Checkout(stage *StageStruct) *A_SPEC_RELATION_GROUPS {
	if _, ok := stage.A_SPEC_RELATION_GROUPSs[a_spec_relation_groups]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutA_SPEC_RELATION_GROUPS(a_spec_relation_groups)
		}
	}
	return a_spec_relation_groups
}

// for satisfaction of GongStruct interface
func (a_spec_relation_groups *A_SPEC_RELATION_GROUPS) GetName() (res string) {
	return a_spec_relation_groups.Name
}

// Stage puts a_spec_relation_ref to the model stage
func (a_spec_relation_ref *A_SPEC_RELATION_REF) Stage(stage *StageStruct) *A_SPEC_RELATION_REF {
	stage.A_SPEC_RELATION_REFs[a_spec_relation_ref] = __member
	stage.A_SPEC_RELATION_REFs_mapString[a_spec_relation_ref.Name] = a_spec_relation_ref

	return a_spec_relation_ref
}

// Unstage removes a_spec_relation_ref off the model stage
func (a_spec_relation_ref *A_SPEC_RELATION_REF) Unstage(stage *StageStruct) *A_SPEC_RELATION_REF {
	delete(stage.A_SPEC_RELATION_REFs, a_spec_relation_ref)
	delete(stage.A_SPEC_RELATION_REFs_mapString, a_spec_relation_ref.Name)
	return a_spec_relation_ref
}

// UnstageVoid removes a_spec_relation_ref off the model stage
func (a_spec_relation_ref *A_SPEC_RELATION_REF) UnstageVoid(stage *StageStruct) {
	delete(stage.A_SPEC_RELATION_REFs, a_spec_relation_ref)
	delete(stage.A_SPEC_RELATION_REFs_mapString, a_spec_relation_ref.Name)
}

// commit a_spec_relation_ref to the back repo (if it is already staged)
func (a_spec_relation_ref *A_SPEC_RELATION_REF) Commit(stage *StageStruct) *A_SPEC_RELATION_REF {
	if _, ok := stage.A_SPEC_RELATION_REFs[a_spec_relation_ref]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitA_SPEC_RELATION_REF(a_spec_relation_ref)
		}
	}
	return a_spec_relation_ref
}

func (a_spec_relation_ref *A_SPEC_RELATION_REF) CommitVoid(stage *StageStruct) {
	a_spec_relation_ref.Commit(stage)
}

// Checkout a_spec_relation_ref to the back repo (if it is already staged)
func (a_spec_relation_ref *A_SPEC_RELATION_REF) Checkout(stage *StageStruct) *A_SPEC_RELATION_REF {
	if _, ok := stage.A_SPEC_RELATION_REFs[a_spec_relation_ref]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutA_SPEC_RELATION_REF(a_spec_relation_ref)
		}
	}
	return a_spec_relation_ref
}

// for satisfaction of GongStruct interface
func (a_spec_relation_ref *A_SPEC_RELATION_REF) GetName() (res string) {
	return a_spec_relation_ref.Name
}

// Stage puts a_spec_relation_type_ref to the model stage
func (a_spec_relation_type_ref *A_SPEC_RELATION_TYPE_REF) Stage(stage *StageStruct) *A_SPEC_RELATION_TYPE_REF {
	stage.A_SPEC_RELATION_TYPE_REFs[a_spec_relation_type_ref] = __member
	stage.A_SPEC_RELATION_TYPE_REFs_mapString[a_spec_relation_type_ref.Name] = a_spec_relation_type_ref

	return a_spec_relation_type_ref
}

// Unstage removes a_spec_relation_type_ref off the model stage
func (a_spec_relation_type_ref *A_SPEC_RELATION_TYPE_REF) Unstage(stage *StageStruct) *A_SPEC_RELATION_TYPE_REF {
	delete(stage.A_SPEC_RELATION_TYPE_REFs, a_spec_relation_type_ref)
	delete(stage.A_SPEC_RELATION_TYPE_REFs_mapString, a_spec_relation_type_ref.Name)
	return a_spec_relation_type_ref
}

// UnstageVoid removes a_spec_relation_type_ref off the model stage
func (a_spec_relation_type_ref *A_SPEC_RELATION_TYPE_REF) UnstageVoid(stage *StageStruct) {
	delete(stage.A_SPEC_RELATION_TYPE_REFs, a_spec_relation_type_ref)
	delete(stage.A_SPEC_RELATION_TYPE_REFs_mapString, a_spec_relation_type_ref.Name)
}

// commit a_spec_relation_type_ref to the back repo (if it is already staged)
func (a_spec_relation_type_ref *A_SPEC_RELATION_TYPE_REF) Commit(stage *StageStruct) *A_SPEC_RELATION_TYPE_REF {
	if _, ok := stage.A_SPEC_RELATION_TYPE_REFs[a_spec_relation_type_ref]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitA_SPEC_RELATION_TYPE_REF(a_spec_relation_type_ref)
		}
	}
	return a_spec_relation_type_ref
}

func (a_spec_relation_type_ref *A_SPEC_RELATION_TYPE_REF) CommitVoid(stage *StageStruct) {
	a_spec_relation_type_ref.Commit(stage)
}

// Checkout a_spec_relation_type_ref to the back repo (if it is already staged)
func (a_spec_relation_type_ref *A_SPEC_RELATION_TYPE_REF) Checkout(stage *StageStruct) *A_SPEC_RELATION_TYPE_REF {
	if _, ok := stage.A_SPEC_RELATION_TYPE_REFs[a_spec_relation_type_ref]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutA_SPEC_RELATION_TYPE_REF(a_spec_relation_type_ref)
		}
	}
	return a_spec_relation_type_ref
}

// for satisfaction of GongStruct interface
func (a_spec_relation_type_ref *A_SPEC_RELATION_TYPE_REF) GetName() (res string) {
	return a_spec_relation_type_ref.Name
}

// Stage puts a_spec_types to the model stage
func (a_spec_types *A_SPEC_TYPES) Stage(stage *StageStruct) *A_SPEC_TYPES {
	stage.A_SPEC_TYPESs[a_spec_types] = __member
	stage.A_SPEC_TYPESs_mapString[a_spec_types.Name] = a_spec_types

	return a_spec_types
}

// Unstage removes a_spec_types off the model stage
func (a_spec_types *A_SPEC_TYPES) Unstage(stage *StageStruct) *A_SPEC_TYPES {
	delete(stage.A_SPEC_TYPESs, a_spec_types)
	delete(stage.A_SPEC_TYPESs_mapString, a_spec_types.Name)
	return a_spec_types
}

// UnstageVoid removes a_spec_types off the model stage
func (a_spec_types *A_SPEC_TYPES) UnstageVoid(stage *StageStruct) {
	delete(stage.A_SPEC_TYPESs, a_spec_types)
	delete(stage.A_SPEC_TYPESs_mapString, a_spec_types.Name)
}

// commit a_spec_types to the back repo (if it is already staged)
func (a_spec_types *A_SPEC_TYPES) Commit(stage *StageStruct) *A_SPEC_TYPES {
	if _, ok := stage.A_SPEC_TYPESs[a_spec_types]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitA_SPEC_TYPES(a_spec_types)
		}
	}
	return a_spec_types
}

func (a_spec_types *A_SPEC_TYPES) CommitVoid(stage *StageStruct) {
	a_spec_types.Commit(stage)
}

// Checkout a_spec_types to the back repo (if it is already staged)
func (a_spec_types *A_SPEC_TYPES) Checkout(stage *StageStruct) *A_SPEC_TYPES {
	if _, ok := stage.A_SPEC_TYPESs[a_spec_types]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutA_SPEC_TYPES(a_spec_types)
		}
	}
	return a_spec_types
}

// for satisfaction of GongStruct interface
func (a_spec_types *A_SPEC_TYPES) GetName() (res string) {
	return a_spec_types.Name
}

// Stage puts a_target_specification_1 to the model stage
func (a_target_specification_1 *A_TARGET_SPECIFICATION_1) Stage(stage *StageStruct) *A_TARGET_SPECIFICATION_1 {
	stage.A_TARGET_SPECIFICATION_1s[a_target_specification_1] = __member
	stage.A_TARGET_SPECIFICATION_1s_mapString[a_target_specification_1.Name] = a_target_specification_1

	return a_target_specification_1
}

// Unstage removes a_target_specification_1 off the model stage
func (a_target_specification_1 *A_TARGET_SPECIFICATION_1) Unstage(stage *StageStruct) *A_TARGET_SPECIFICATION_1 {
	delete(stage.A_TARGET_SPECIFICATION_1s, a_target_specification_1)
	delete(stage.A_TARGET_SPECIFICATION_1s_mapString, a_target_specification_1.Name)
	return a_target_specification_1
}

// UnstageVoid removes a_target_specification_1 off the model stage
func (a_target_specification_1 *A_TARGET_SPECIFICATION_1) UnstageVoid(stage *StageStruct) {
	delete(stage.A_TARGET_SPECIFICATION_1s, a_target_specification_1)
	delete(stage.A_TARGET_SPECIFICATION_1s_mapString, a_target_specification_1.Name)
}

// commit a_target_specification_1 to the back repo (if it is already staged)
func (a_target_specification_1 *A_TARGET_SPECIFICATION_1) Commit(stage *StageStruct) *A_TARGET_SPECIFICATION_1 {
	if _, ok := stage.A_TARGET_SPECIFICATION_1s[a_target_specification_1]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitA_TARGET_SPECIFICATION_1(a_target_specification_1)
		}
	}
	return a_target_specification_1
}

func (a_target_specification_1 *A_TARGET_SPECIFICATION_1) CommitVoid(stage *StageStruct) {
	a_target_specification_1.Commit(stage)
}

// Checkout a_target_specification_1 to the back repo (if it is already staged)
func (a_target_specification_1 *A_TARGET_SPECIFICATION_1) Checkout(stage *StageStruct) *A_TARGET_SPECIFICATION_1 {
	if _, ok := stage.A_TARGET_SPECIFICATION_1s[a_target_specification_1]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutA_TARGET_SPECIFICATION_1(a_target_specification_1)
		}
	}
	return a_target_specification_1
}

// for satisfaction of GongStruct interface
func (a_target_specification_1 *A_TARGET_SPECIFICATION_1) GetName() (res string) {
	return a_target_specification_1.Name
}

// Stage puts a_the_header to the model stage
func (a_the_header *A_THE_HEADER) Stage(stage *StageStruct) *A_THE_HEADER {
	stage.A_THE_HEADERs[a_the_header] = __member
	stage.A_THE_HEADERs_mapString[a_the_header.Name] = a_the_header

	return a_the_header
}

// Unstage removes a_the_header off the model stage
func (a_the_header *A_THE_HEADER) Unstage(stage *StageStruct) *A_THE_HEADER {
	delete(stage.A_THE_HEADERs, a_the_header)
	delete(stage.A_THE_HEADERs_mapString, a_the_header.Name)
	return a_the_header
}

// UnstageVoid removes a_the_header off the model stage
func (a_the_header *A_THE_HEADER) UnstageVoid(stage *StageStruct) {
	delete(stage.A_THE_HEADERs, a_the_header)
	delete(stage.A_THE_HEADERs_mapString, a_the_header.Name)
}

// commit a_the_header to the back repo (if it is already staged)
func (a_the_header *A_THE_HEADER) Commit(stage *StageStruct) *A_THE_HEADER {
	if _, ok := stage.A_THE_HEADERs[a_the_header]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitA_THE_HEADER(a_the_header)
		}
	}
	return a_the_header
}

func (a_the_header *A_THE_HEADER) CommitVoid(stage *StageStruct) {
	a_the_header.Commit(stage)
}

// Checkout a_the_header to the back repo (if it is already staged)
func (a_the_header *A_THE_HEADER) Checkout(stage *StageStruct) *A_THE_HEADER {
	if _, ok := stage.A_THE_HEADERs[a_the_header]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutA_THE_HEADER(a_the_header)
		}
	}
	return a_the_header
}

// for satisfaction of GongStruct interface
func (a_the_header *A_THE_HEADER) GetName() (res string) {
	return a_the_header.Name
}

// Stage puts a_tool_extensions to the model stage
func (a_tool_extensions *A_TOOL_EXTENSIONS) Stage(stage *StageStruct) *A_TOOL_EXTENSIONS {
	stage.A_TOOL_EXTENSIONSs[a_tool_extensions] = __member
	stage.A_TOOL_EXTENSIONSs_mapString[a_tool_extensions.Name] = a_tool_extensions

	return a_tool_extensions
}

// Unstage removes a_tool_extensions off the model stage
func (a_tool_extensions *A_TOOL_EXTENSIONS) Unstage(stage *StageStruct) *A_TOOL_EXTENSIONS {
	delete(stage.A_TOOL_EXTENSIONSs, a_tool_extensions)
	delete(stage.A_TOOL_EXTENSIONSs_mapString, a_tool_extensions.Name)
	return a_tool_extensions
}

// UnstageVoid removes a_tool_extensions off the model stage
func (a_tool_extensions *A_TOOL_EXTENSIONS) UnstageVoid(stage *StageStruct) {
	delete(stage.A_TOOL_EXTENSIONSs, a_tool_extensions)
	delete(stage.A_TOOL_EXTENSIONSs_mapString, a_tool_extensions.Name)
}

// commit a_tool_extensions to the back repo (if it is already staged)
func (a_tool_extensions *A_TOOL_EXTENSIONS) Commit(stage *StageStruct) *A_TOOL_EXTENSIONS {
	if _, ok := stage.A_TOOL_EXTENSIONSs[a_tool_extensions]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitA_TOOL_EXTENSIONS(a_tool_extensions)
		}
	}
	return a_tool_extensions
}

func (a_tool_extensions *A_TOOL_EXTENSIONS) CommitVoid(stage *StageStruct) {
	a_tool_extensions.Commit(stage)
}

// Checkout a_tool_extensions to the back repo (if it is already staged)
func (a_tool_extensions *A_TOOL_EXTENSIONS) Checkout(stage *StageStruct) *A_TOOL_EXTENSIONS {
	if _, ok := stage.A_TOOL_EXTENSIONSs[a_tool_extensions]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutA_TOOL_EXTENSIONS(a_tool_extensions)
		}
	}
	return a_tool_extensions
}

// for satisfaction of GongStruct interface
func (a_tool_extensions *A_TOOL_EXTENSIONS) GetName() (res string) {
	return a_tool_extensions.Name
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
	CreateORMA_ALTERNATIVE_ID(A_ALTERNATIVE_ID *A_ALTERNATIVE_ID)
	CreateORMA_ATTRIBUTE_DEFINITION_BOOLEAN_REF(A_ATTRIBUTE_DEFINITION_BOOLEAN_REF *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF)
	CreateORMA_ATTRIBUTE_DEFINITION_DATE_REF(A_ATTRIBUTE_DEFINITION_DATE_REF *A_ATTRIBUTE_DEFINITION_DATE_REF)
	CreateORMA_ATTRIBUTE_DEFINITION_ENUMERATION_REF(A_ATTRIBUTE_DEFINITION_ENUMERATION_REF *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF)
	CreateORMA_ATTRIBUTE_DEFINITION_INTEGER_REF(A_ATTRIBUTE_DEFINITION_INTEGER_REF *A_ATTRIBUTE_DEFINITION_INTEGER_REF)
	CreateORMA_ATTRIBUTE_DEFINITION_REAL_REF(A_ATTRIBUTE_DEFINITION_REAL_REF *A_ATTRIBUTE_DEFINITION_REAL_REF)
	CreateORMA_ATTRIBUTE_DEFINITION_STRING_REF(A_ATTRIBUTE_DEFINITION_STRING_REF *A_ATTRIBUTE_DEFINITION_STRING_REF)
	CreateORMA_ATTRIBUTE_DEFINITION_XHTML_REF(A_ATTRIBUTE_DEFINITION_XHTML_REF *A_ATTRIBUTE_DEFINITION_XHTML_REF)
	CreateORMA_ATTRIBUTE_VALUE_BOOLEAN(A_ATTRIBUTE_VALUE_BOOLEAN *A_ATTRIBUTE_VALUE_BOOLEAN)
	CreateORMA_ATTRIBUTE_VALUE_DATE(A_ATTRIBUTE_VALUE_DATE *A_ATTRIBUTE_VALUE_DATE)
	CreateORMA_ATTRIBUTE_VALUE_ENUMERATION(A_ATTRIBUTE_VALUE_ENUMERATION *A_ATTRIBUTE_VALUE_ENUMERATION)
	CreateORMA_ATTRIBUTE_VALUE_INTEGER(A_ATTRIBUTE_VALUE_INTEGER *A_ATTRIBUTE_VALUE_INTEGER)
	CreateORMA_ATTRIBUTE_VALUE_REAL(A_ATTRIBUTE_VALUE_REAL *A_ATTRIBUTE_VALUE_REAL)
	CreateORMA_ATTRIBUTE_VALUE_STRING(A_ATTRIBUTE_VALUE_STRING *A_ATTRIBUTE_VALUE_STRING)
	CreateORMA_ATTRIBUTE_VALUE_XHTML(A_ATTRIBUTE_VALUE_XHTML *A_ATTRIBUTE_VALUE_XHTML)
	CreateORMA_ATTRIBUTE_VALUE_XHTML_1(A_ATTRIBUTE_VALUE_XHTML_1 *A_ATTRIBUTE_VALUE_XHTML_1)
	CreateORMA_CHILDREN(A_CHILDREN *A_CHILDREN)
	CreateORMA_CORE_CONTENT(A_CORE_CONTENT *A_CORE_CONTENT)
	CreateORMA_DATATYPES(A_DATATYPES *A_DATATYPES)
	CreateORMA_DATATYPE_DEFINITION_BOOLEAN_REF(A_DATATYPE_DEFINITION_BOOLEAN_REF *A_DATATYPE_DEFINITION_BOOLEAN_REF)
	CreateORMA_DATATYPE_DEFINITION_DATE_REF(A_DATATYPE_DEFINITION_DATE_REF *A_DATATYPE_DEFINITION_DATE_REF)
	CreateORMA_DATATYPE_DEFINITION_ENUMERATION_REF(A_DATATYPE_DEFINITION_ENUMERATION_REF *A_DATATYPE_DEFINITION_ENUMERATION_REF)
	CreateORMA_DATATYPE_DEFINITION_INTEGER_REF(A_DATATYPE_DEFINITION_INTEGER_REF *A_DATATYPE_DEFINITION_INTEGER_REF)
	CreateORMA_DATATYPE_DEFINITION_REAL_REF(A_DATATYPE_DEFINITION_REAL_REF *A_DATATYPE_DEFINITION_REAL_REF)
	CreateORMA_DATATYPE_DEFINITION_STRING_REF(A_DATATYPE_DEFINITION_STRING_REF *A_DATATYPE_DEFINITION_STRING_REF)
	CreateORMA_DATATYPE_DEFINITION_XHTML_REF(A_DATATYPE_DEFINITION_XHTML_REF *A_DATATYPE_DEFINITION_XHTML_REF)
	CreateORMA_EDITABLE_ATTS(A_EDITABLE_ATTS *A_EDITABLE_ATTS)
	CreateORMA_ENUM_VALUE_REF(A_ENUM_VALUE_REF *A_ENUM_VALUE_REF)
	CreateORMA_OBJECT(A_OBJECT *A_OBJECT)
	CreateORMA_PROPERTIES(A_PROPERTIES *A_PROPERTIES)
	CreateORMA_RELATION_GROUP_TYPE_REF(A_RELATION_GROUP_TYPE_REF *A_RELATION_GROUP_TYPE_REF)
	CreateORMA_SOURCE_1(A_SOURCE_1 *A_SOURCE_1)
	CreateORMA_SPECIFICATIONS(A_SPECIFICATIONS *A_SPECIFICATIONS)
	CreateORMA_SPECIFICATION_TYPE_REF(A_SPECIFICATION_TYPE_REF *A_SPECIFICATION_TYPE_REF)
	CreateORMA_SPECIFIED_VALUES(A_SPECIFIED_VALUES *A_SPECIFIED_VALUES)
	CreateORMA_SPEC_ATTRIBUTES(A_SPEC_ATTRIBUTES *A_SPEC_ATTRIBUTES)
	CreateORMA_SPEC_OBJECTS(A_SPEC_OBJECTS *A_SPEC_OBJECTS)
	CreateORMA_SPEC_OBJECT_TYPE_REF(A_SPEC_OBJECT_TYPE_REF *A_SPEC_OBJECT_TYPE_REF)
	CreateORMA_SPEC_RELATIONS(A_SPEC_RELATIONS *A_SPEC_RELATIONS)
	CreateORMA_SPEC_RELATION_GROUPS(A_SPEC_RELATION_GROUPS *A_SPEC_RELATION_GROUPS)
	CreateORMA_SPEC_RELATION_REF(A_SPEC_RELATION_REF *A_SPEC_RELATION_REF)
	CreateORMA_SPEC_RELATION_TYPE_REF(A_SPEC_RELATION_TYPE_REF *A_SPEC_RELATION_TYPE_REF)
	CreateORMA_SPEC_TYPES(A_SPEC_TYPES *A_SPEC_TYPES)
	CreateORMA_TARGET_SPECIFICATION_1(A_TARGET_SPECIFICATION_1 *A_TARGET_SPECIFICATION_1)
	CreateORMA_THE_HEADER(A_THE_HEADER *A_THE_HEADER)
	CreateORMA_TOOL_EXTENSIONS(A_TOOL_EXTENSIONS *A_TOOL_EXTENSIONS)
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
	DeleteORMA_ALTERNATIVE_ID(A_ALTERNATIVE_ID *A_ALTERNATIVE_ID)
	DeleteORMA_ATTRIBUTE_DEFINITION_BOOLEAN_REF(A_ATTRIBUTE_DEFINITION_BOOLEAN_REF *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF)
	DeleteORMA_ATTRIBUTE_DEFINITION_DATE_REF(A_ATTRIBUTE_DEFINITION_DATE_REF *A_ATTRIBUTE_DEFINITION_DATE_REF)
	DeleteORMA_ATTRIBUTE_DEFINITION_ENUMERATION_REF(A_ATTRIBUTE_DEFINITION_ENUMERATION_REF *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF)
	DeleteORMA_ATTRIBUTE_DEFINITION_INTEGER_REF(A_ATTRIBUTE_DEFINITION_INTEGER_REF *A_ATTRIBUTE_DEFINITION_INTEGER_REF)
	DeleteORMA_ATTRIBUTE_DEFINITION_REAL_REF(A_ATTRIBUTE_DEFINITION_REAL_REF *A_ATTRIBUTE_DEFINITION_REAL_REF)
	DeleteORMA_ATTRIBUTE_DEFINITION_STRING_REF(A_ATTRIBUTE_DEFINITION_STRING_REF *A_ATTRIBUTE_DEFINITION_STRING_REF)
	DeleteORMA_ATTRIBUTE_DEFINITION_XHTML_REF(A_ATTRIBUTE_DEFINITION_XHTML_REF *A_ATTRIBUTE_DEFINITION_XHTML_REF)
	DeleteORMA_ATTRIBUTE_VALUE_BOOLEAN(A_ATTRIBUTE_VALUE_BOOLEAN *A_ATTRIBUTE_VALUE_BOOLEAN)
	DeleteORMA_ATTRIBUTE_VALUE_DATE(A_ATTRIBUTE_VALUE_DATE *A_ATTRIBUTE_VALUE_DATE)
	DeleteORMA_ATTRIBUTE_VALUE_ENUMERATION(A_ATTRIBUTE_VALUE_ENUMERATION *A_ATTRIBUTE_VALUE_ENUMERATION)
	DeleteORMA_ATTRIBUTE_VALUE_INTEGER(A_ATTRIBUTE_VALUE_INTEGER *A_ATTRIBUTE_VALUE_INTEGER)
	DeleteORMA_ATTRIBUTE_VALUE_REAL(A_ATTRIBUTE_VALUE_REAL *A_ATTRIBUTE_VALUE_REAL)
	DeleteORMA_ATTRIBUTE_VALUE_STRING(A_ATTRIBUTE_VALUE_STRING *A_ATTRIBUTE_VALUE_STRING)
	DeleteORMA_ATTRIBUTE_VALUE_XHTML(A_ATTRIBUTE_VALUE_XHTML *A_ATTRIBUTE_VALUE_XHTML)
	DeleteORMA_ATTRIBUTE_VALUE_XHTML_1(A_ATTRIBUTE_VALUE_XHTML_1 *A_ATTRIBUTE_VALUE_XHTML_1)
	DeleteORMA_CHILDREN(A_CHILDREN *A_CHILDREN)
	DeleteORMA_CORE_CONTENT(A_CORE_CONTENT *A_CORE_CONTENT)
	DeleteORMA_DATATYPES(A_DATATYPES *A_DATATYPES)
	DeleteORMA_DATATYPE_DEFINITION_BOOLEAN_REF(A_DATATYPE_DEFINITION_BOOLEAN_REF *A_DATATYPE_DEFINITION_BOOLEAN_REF)
	DeleteORMA_DATATYPE_DEFINITION_DATE_REF(A_DATATYPE_DEFINITION_DATE_REF *A_DATATYPE_DEFINITION_DATE_REF)
	DeleteORMA_DATATYPE_DEFINITION_ENUMERATION_REF(A_DATATYPE_DEFINITION_ENUMERATION_REF *A_DATATYPE_DEFINITION_ENUMERATION_REF)
	DeleteORMA_DATATYPE_DEFINITION_INTEGER_REF(A_DATATYPE_DEFINITION_INTEGER_REF *A_DATATYPE_DEFINITION_INTEGER_REF)
	DeleteORMA_DATATYPE_DEFINITION_REAL_REF(A_DATATYPE_DEFINITION_REAL_REF *A_DATATYPE_DEFINITION_REAL_REF)
	DeleteORMA_DATATYPE_DEFINITION_STRING_REF(A_DATATYPE_DEFINITION_STRING_REF *A_DATATYPE_DEFINITION_STRING_REF)
	DeleteORMA_DATATYPE_DEFINITION_XHTML_REF(A_DATATYPE_DEFINITION_XHTML_REF *A_DATATYPE_DEFINITION_XHTML_REF)
	DeleteORMA_EDITABLE_ATTS(A_EDITABLE_ATTS *A_EDITABLE_ATTS)
	DeleteORMA_ENUM_VALUE_REF(A_ENUM_VALUE_REF *A_ENUM_VALUE_REF)
	DeleteORMA_OBJECT(A_OBJECT *A_OBJECT)
	DeleteORMA_PROPERTIES(A_PROPERTIES *A_PROPERTIES)
	DeleteORMA_RELATION_GROUP_TYPE_REF(A_RELATION_GROUP_TYPE_REF *A_RELATION_GROUP_TYPE_REF)
	DeleteORMA_SOURCE_1(A_SOURCE_1 *A_SOURCE_1)
	DeleteORMA_SPECIFICATIONS(A_SPECIFICATIONS *A_SPECIFICATIONS)
	DeleteORMA_SPECIFICATION_TYPE_REF(A_SPECIFICATION_TYPE_REF *A_SPECIFICATION_TYPE_REF)
	DeleteORMA_SPECIFIED_VALUES(A_SPECIFIED_VALUES *A_SPECIFIED_VALUES)
	DeleteORMA_SPEC_ATTRIBUTES(A_SPEC_ATTRIBUTES *A_SPEC_ATTRIBUTES)
	DeleteORMA_SPEC_OBJECTS(A_SPEC_OBJECTS *A_SPEC_OBJECTS)
	DeleteORMA_SPEC_OBJECT_TYPE_REF(A_SPEC_OBJECT_TYPE_REF *A_SPEC_OBJECT_TYPE_REF)
	DeleteORMA_SPEC_RELATIONS(A_SPEC_RELATIONS *A_SPEC_RELATIONS)
	DeleteORMA_SPEC_RELATION_GROUPS(A_SPEC_RELATION_GROUPS *A_SPEC_RELATION_GROUPS)
	DeleteORMA_SPEC_RELATION_REF(A_SPEC_RELATION_REF *A_SPEC_RELATION_REF)
	DeleteORMA_SPEC_RELATION_TYPE_REF(A_SPEC_RELATION_TYPE_REF *A_SPEC_RELATION_TYPE_REF)
	DeleteORMA_SPEC_TYPES(A_SPEC_TYPES *A_SPEC_TYPES)
	DeleteORMA_TARGET_SPECIFICATION_1(A_TARGET_SPECIFICATION_1 *A_TARGET_SPECIFICATION_1)
	DeleteORMA_THE_HEADER(A_THE_HEADER *A_THE_HEADER)
	DeleteORMA_TOOL_EXTENSIONS(A_TOOL_EXTENSIONS *A_TOOL_EXTENSIONS)
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

	stage.A_ALTERNATIVE_IDs = make(map[*A_ALTERNATIVE_ID]any)
	stage.A_ALTERNATIVE_IDs_mapString = make(map[string]*A_ALTERNATIVE_ID)

	stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFs = make(map[*A_ATTRIBUTE_DEFINITION_BOOLEAN_REF]any)
	stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFs_mapString = make(map[string]*A_ATTRIBUTE_DEFINITION_BOOLEAN_REF)

	stage.A_ATTRIBUTE_DEFINITION_DATE_REFs = make(map[*A_ATTRIBUTE_DEFINITION_DATE_REF]any)
	stage.A_ATTRIBUTE_DEFINITION_DATE_REFs_mapString = make(map[string]*A_ATTRIBUTE_DEFINITION_DATE_REF)

	stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFs = make(map[*A_ATTRIBUTE_DEFINITION_ENUMERATION_REF]any)
	stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFs_mapString = make(map[string]*A_ATTRIBUTE_DEFINITION_ENUMERATION_REF)

	stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFs = make(map[*A_ATTRIBUTE_DEFINITION_INTEGER_REF]any)
	stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFs_mapString = make(map[string]*A_ATTRIBUTE_DEFINITION_INTEGER_REF)

	stage.A_ATTRIBUTE_DEFINITION_REAL_REFs = make(map[*A_ATTRIBUTE_DEFINITION_REAL_REF]any)
	stage.A_ATTRIBUTE_DEFINITION_REAL_REFs_mapString = make(map[string]*A_ATTRIBUTE_DEFINITION_REAL_REF)

	stage.A_ATTRIBUTE_DEFINITION_STRING_REFs = make(map[*A_ATTRIBUTE_DEFINITION_STRING_REF]any)
	stage.A_ATTRIBUTE_DEFINITION_STRING_REFs_mapString = make(map[string]*A_ATTRIBUTE_DEFINITION_STRING_REF)

	stage.A_ATTRIBUTE_DEFINITION_XHTML_REFs = make(map[*A_ATTRIBUTE_DEFINITION_XHTML_REF]any)
	stage.A_ATTRIBUTE_DEFINITION_XHTML_REFs_mapString = make(map[string]*A_ATTRIBUTE_DEFINITION_XHTML_REF)

	stage.A_ATTRIBUTE_VALUE_BOOLEANs = make(map[*A_ATTRIBUTE_VALUE_BOOLEAN]any)
	stage.A_ATTRIBUTE_VALUE_BOOLEANs_mapString = make(map[string]*A_ATTRIBUTE_VALUE_BOOLEAN)

	stage.A_ATTRIBUTE_VALUE_DATEs = make(map[*A_ATTRIBUTE_VALUE_DATE]any)
	stage.A_ATTRIBUTE_VALUE_DATEs_mapString = make(map[string]*A_ATTRIBUTE_VALUE_DATE)

	stage.A_ATTRIBUTE_VALUE_ENUMERATIONs = make(map[*A_ATTRIBUTE_VALUE_ENUMERATION]any)
	stage.A_ATTRIBUTE_VALUE_ENUMERATIONs_mapString = make(map[string]*A_ATTRIBUTE_VALUE_ENUMERATION)

	stage.A_ATTRIBUTE_VALUE_INTEGERs = make(map[*A_ATTRIBUTE_VALUE_INTEGER]any)
	stage.A_ATTRIBUTE_VALUE_INTEGERs_mapString = make(map[string]*A_ATTRIBUTE_VALUE_INTEGER)

	stage.A_ATTRIBUTE_VALUE_REALs = make(map[*A_ATTRIBUTE_VALUE_REAL]any)
	stage.A_ATTRIBUTE_VALUE_REALs_mapString = make(map[string]*A_ATTRIBUTE_VALUE_REAL)

	stage.A_ATTRIBUTE_VALUE_STRINGs = make(map[*A_ATTRIBUTE_VALUE_STRING]any)
	stage.A_ATTRIBUTE_VALUE_STRINGs_mapString = make(map[string]*A_ATTRIBUTE_VALUE_STRING)

	stage.A_ATTRIBUTE_VALUE_XHTMLs = make(map[*A_ATTRIBUTE_VALUE_XHTML]any)
	stage.A_ATTRIBUTE_VALUE_XHTMLs_mapString = make(map[string]*A_ATTRIBUTE_VALUE_XHTML)

	stage.A_ATTRIBUTE_VALUE_XHTML_1s = make(map[*A_ATTRIBUTE_VALUE_XHTML_1]any)
	stage.A_ATTRIBUTE_VALUE_XHTML_1s_mapString = make(map[string]*A_ATTRIBUTE_VALUE_XHTML_1)

	stage.A_CHILDRENs = make(map[*A_CHILDREN]any)
	stage.A_CHILDRENs_mapString = make(map[string]*A_CHILDREN)

	stage.A_CORE_CONTENTs = make(map[*A_CORE_CONTENT]any)
	stage.A_CORE_CONTENTs_mapString = make(map[string]*A_CORE_CONTENT)

	stage.A_DATATYPESs = make(map[*A_DATATYPES]any)
	stage.A_DATATYPESs_mapString = make(map[string]*A_DATATYPES)

	stage.A_DATATYPE_DEFINITION_BOOLEAN_REFs = make(map[*A_DATATYPE_DEFINITION_BOOLEAN_REF]any)
	stage.A_DATATYPE_DEFINITION_BOOLEAN_REFs_mapString = make(map[string]*A_DATATYPE_DEFINITION_BOOLEAN_REF)

	stage.A_DATATYPE_DEFINITION_DATE_REFs = make(map[*A_DATATYPE_DEFINITION_DATE_REF]any)
	stage.A_DATATYPE_DEFINITION_DATE_REFs_mapString = make(map[string]*A_DATATYPE_DEFINITION_DATE_REF)

	stage.A_DATATYPE_DEFINITION_ENUMERATION_REFs = make(map[*A_DATATYPE_DEFINITION_ENUMERATION_REF]any)
	stage.A_DATATYPE_DEFINITION_ENUMERATION_REFs_mapString = make(map[string]*A_DATATYPE_DEFINITION_ENUMERATION_REF)

	stage.A_DATATYPE_DEFINITION_INTEGER_REFs = make(map[*A_DATATYPE_DEFINITION_INTEGER_REF]any)
	stage.A_DATATYPE_DEFINITION_INTEGER_REFs_mapString = make(map[string]*A_DATATYPE_DEFINITION_INTEGER_REF)

	stage.A_DATATYPE_DEFINITION_REAL_REFs = make(map[*A_DATATYPE_DEFINITION_REAL_REF]any)
	stage.A_DATATYPE_DEFINITION_REAL_REFs_mapString = make(map[string]*A_DATATYPE_DEFINITION_REAL_REF)

	stage.A_DATATYPE_DEFINITION_STRING_REFs = make(map[*A_DATATYPE_DEFINITION_STRING_REF]any)
	stage.A_DATATYPE_DEFINITION_STRING_REFs_mapString = make(map[string]*A_DATATYPE_DEFINITION_STRING_REF)

	stage.A_DATATYPE_DEFINITION_XHTML_REFs = make(map[*A_DATATYPE_DEFINITION_XHTML_REF]any)
	stage.A_DATATYPE_DEFINITION_XHTML_REFs_mapString = make(map[string]*A_DATATYPE_DEFINITION_XHTML_REF)

	stage.A_EDITABLE_ATTSs = make(map[*A_EDITABLE_ATTS]any)
	stage.A_EDITABLE_ATTSs_mapString = make(map[string]*A_EDITABLE_ATTS)

	stage.A_ENUM_VALUE_REFs = make(map[*A_ENUM_VALUE_REF]any)
	stage.A_ENUM_VALUE_REFs_mapString = make(map[string]*A_ENUM_VALUE_REF)

	stage.A_OBJECTs = make(map[*A_OBJECT]any)
	stage.A_OBJECTs_mapString = make(map[string]*A_OBJECT)

	stage.A_PROPERTIESs = make(map[*A_PROPERTIES]any)
	stage.A_PROPERTIESs_mapString = make(map[string]*A_PROPERTIES)

	stage.A_RELATION_GROUP_TYPE_REFs = make(map[*A_RELATION_GROUP_TYPE_REF]any)
	stage.A_RELATION_GROUP_TYPE_REFs_mapString = make(map[string]*A_RELATION_GROUP_TYPE_REF)

	stage.A_SOURCE_1s = make(map[*A_SOURCE_1]any)
	stage.A_SOURCE_1s_mapString = make(map[string]*A_SOURCE_1)

	stage.A_SPECIFICATIONSs = make(map[*A_SPECIFICATIONS]any)
	stage.A_SPECIFICATIONSs_mapString = make(map[string]*A_SPECIFICATIONS)

	stage.A_SPECIFICATION_TYPE_REFs = make(map[*A_SPECIFICATION_TYPE_REF]any)
	stage.A_SPECIFICATION_TYPE_REFs_mapString = make(map[string]*A_SPECIFICATION_TYPE_REF)

	stage.A_SPECIFIED_VALUESs = make(map[*A_SPECIFIED_VALUES]any)
	stage.A_SPECIFIED_VALUESs_mapString = make(map[string]*A_SPECIFIED_VALUES)

	stage.A_SPEC_ATTRIBUTESs = make(map[*A_SPEC_ATTRIBUTES]any)
	stage.A_SPEC_ATTRIBUTESs_mapString = make(map[string]*A_SPEC_ATTRIBUTES)

	stage.A_SPEC_OBJECTSs = make(map[*A_SPEC_OBJECTS]any)
	stage.A_SPEC_OBJECTSs_mapString = make(map[string]*A_SPEC_OBJECTS)

	stage.A_SPEC_OBJECT_TYPE_REFs = make(map[*A_SPEC_OBJECT_TYPE_REF]any)
	stage.A_SPEC_OBJECT_TYPE_REFs_mapString = make(map[string]*A_SPEC_OBJECT_TYPE_REF)

	stage.A_SPEC_RELATIONSs = make(map[*A_SPEC_RELATIONS]any)
	stage.A_SPEC_RELATIONSs_mapString = make(map[string]*A_SPEC_RELATIONS)

	stage.A_SPEC_RELATION_GROUPSs = make(map[*A_SPEC_RELATION_GROUPS]any)
	stage.A_SPEC_RELATION_GROUPSs_mapString = make(map[string]*A_SPEC_RELATION_GROUPS)

	stage.A_SPEC_RELATION_REFs = make(map[*A_SPEC_RELATION_REF]any)
	stage.A_SPEC_RELATION_REFs_mapString = make(map[string]*A_SPEC_RELATION_REF)

	stage.A_SPEC_RELATION_TYPE_REFs = make(map[*A_SPEC_RELATION_TYPE_REF]any)
	stage.A_SPEC_RELATION_TYPE_REFs_mapString = make(map[string]*A_SPEC_RELATION_TYPE_REF)

	stage.A_SPEC_TYPESs = make(map[*A_SPEC_TYPES]any)
	stage.A_SPEC_TYPESs_mapString = make(map[string]*A_SPEC_TYPES)

	stage.A_TARGET_SPECIFICATION_1s = make(map[*A_TARGET_SPECIFICATION_1]any)
	stage.A_TARGET_SPECIFICATION_1s_mapString = make(map[string]*A_TARGET_SPECIFICATION_1)

	stage.A_THE_HEADERs = make(map[*A_THE_HEADER]any)
	stage.A_THE_HEADERs_mapString = make(map[string]*A_THE_HEADER)

	stage.A_TOOL_EXTENSIONSs = make(map[*A_TOOL_EXTENSIONS]any)
	stage.A_TOOL_EXTENSIONSs_mapString = make(map[string]*A_TOOL_EXTENSIONS)

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

	stage.A_ALTERNATIVE_IDs = nil
	stage.A_ALTERNATIVE_IDs_mapString = nil

	stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFs = nil
	stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFs_mapString = nil

	stage.A_ATTRIBUTE_DEFINITION_DATE_REFs = nil
	stage.A_ATTRIBUTE_DEFINITION_DATE_REFs_mapString = nil

	stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFs = nil
	stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFs_mapString = nil

	stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFs = nil
	stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFs_mapString = nil

	stage.A_ATTRIBUTE_DEFINITION_REAL_REFs = nil
	stage.A_ATTRIBUTE_DEFINITION_REAL_REFs_mapString = nil

	stage.A_ATTRIBUTE_DEFINITION_STRING_REFs = nil
	stage.A_ATTRIBUTE_DEFINITION_STRING_REFs_mapString = nil

	stage.A_ATTRIBUTE_DEFINITION_XHTML_REFs = nil
	stage.A_ATTRIBUTE_DEFINITION_XHTML_REFs_mapString = nil

	stage.A_ATTRIBUTE_VALUE_BOOLEANs = nil
	stage.A_ATTRIBUTE_VALUE_BOOLEANs_mapString = nil

	stage.A_ATTRIBUTE_VALUE_DATEs = nil
	stage.A_ATTRIBUTE_VALUE_DATEs_mapString = nil

	stage.A_ATTRIBUTE_VALUE_ENUMERATIONs = nil
	stage.A_ATTRIBUTE_VALUE_ENUMERATIONs_mapString = nil

	stage.A_ATTRIBUTE_VALUE_INTEGERs = nil
	stage.A_ATTRIBUTE_VALUE_INTEGERs_mapString = nil

	stage.A_ATTRIBUTE_VALUE_REALs = nil
	stage.A_ATTRIBUTE_VALUE_REALs_mapString = nil

	stage.A_ATTRIBUTE_VALUE_STRINGs = nil
	stage.A_ATTRIBUTE_VALUE_STRINGs_mapString = nil

	stage.A_ATTRIBUTE_VALUE_XHTMLs = nil
	stage.A_ATTRIBUTE_VALUE_XHTMLs_mapString = nil

	stage.A_ATTRIBUTE_VALUE_XHTML_1s = nil
	stage.A_ATTRIBUTE_VALUE_XHTML_1s_mapString = nil

	stage.A_CHILDRENs = nil
	stage.A_CHILDRENs_mapString = nil

	stage.A_CORE_CONTENTs = nil
	stage.A_CORE_CONTENTs_mapString = nil

	stage.A_DATATYPESs = nil
	stage.A_DATATYPESs_mapString = nil

	stage.A_DATATYPE_DEFINITION_BOOLEAN_REFs = nil
	stage.A_DATATYPE_DEFINITION_BOOLEAN_REFs_mapString = nil

	stage.A_DATATYPE_DEFINITION_DATE_REFs = nil
	stage.A_DATATYPE_DEFINITION_DATE_REFs_mapString = nil

	stage.A_DATATYPE_DEFINITION_ENUMERATION_REFs = nil
	stage.A_DATATYPE_DEFINITION_ENUMERATION_REFs_mapString = nil

	stage.A_DATATYPE_DEFINITION_INTEGER_REFs = nil
	stage.A_DATATYPE_DEFINITION_INTEGER_REFs_mapString = nil

	stage.A_DATATYPE_DEFINITION_REAL_REFs = nil
	stage.A_DATATYPE_DEFINITION_REAL_REFs_mapString = nil

	stage.A_DATATYPE_DEFINITION_STRING_REFs = nil
	stage.A_DATATYPE_DEFINITION_STRING_REFs_mapString = nil

	stage.A_DATATYPE_DEFINITION_XHTML_REFs = nil
	stage.A_DATATYPE_DEFINITION_XHTML_REFs_mapString = nil

	stage.A_EDITABLE_ATTSs = nil
	stage.A_EDITABLE_ATTSs_mapString = nil

	stage.A_ENUM_VALUE_REFs = nil
	stage.A_ENUM_VALUE_REFs_mapString = nil

	stage.A_OBJECTs = nil
	stage.A_OBJECTs_mapString = nil

	stage.A_PROPERTIESs = nil
	stage.A_PROPERTIESs_mapString = nil

	stage.A_RELATION_GROUP_TYPE_REFs = nil
	stage.A_RELATION_GROUP_TYPE_REFs_mapString = nil

	stage.A_SOURCE_1s = nil
	stage.A_SOURCE_1s_mapString = nil

	stage.A_SPECIFICATIONSs = nil
	stage.A_SPECIFICATIONSs_mapString = nil

	stage.A_SPECIFICATION_TYPE_REFs = nil
	stage.A_SPECIFICATION_TYPE_REFs_mapString = nil

	stage.A_SPECIFIED_VALUESs = nil
	stage.A_SPECIFIED_VALUESs_mapString = nil

	stage.A_SPEC_ATTRIBUTESs = nil
	stage.A_SPEC_ATTRIBUTESs_mapString = nil

	stage.A_SPEC_OBJECTSs = nil
	stage.A_SPEC_OBJECTSs_mapString = nil

	stage.A_SPEC_OBJECT_TYPE_REFs = nil
	stage.A_SPEC_OBJECT_TYPE_REFs_mapString = nil

	stage.A_SPEC_RELATIONSs = nil
	stage.A_SPEC_RELATIONSs_mapString = nil

	stage.A_SPEC_RELATION_GROUPSs = nil
	stage.A_SPEC_RELATION_GROUPSs_mapString = nil

	stage.A_SPEC_RELATION_REFs = nil
	stage.A_SPEC_RELATION_REFs_mapString = nil

	stage.A_SPEC_RELATION_TYPE_REFs = nil
	stage.A_SPEC_RELATION_TYPE_REFs_mapString = nil

	stage.A_SPEC_TYPESs = nil
	stage.A_SPEC_TYPESs_mapString = nil

	stage.A_TARGET_SPECIFICATION_1s = nil
	stage.A_TARGET_SPECIFICATION_1s_mapString = nil

	stage.A_THE_HEADERs = nil
	stage.A_THE_HEADERs_mapString = nil

	stage.A_TOOL_EXTENSIONSs = nil
	stage.A_TOOL_EXTENSIONSs_mapString = nil

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

	for a_alternative_id := range stage.A_ALTERNATIVE_IDs {
		a_alternative_id.Unstage(stage)
	}

	for a_attribute_definition_boolean_ref := range stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFs {
		a_attribute_definition_boolean_ref.Unstage(stage)
	}

	for a_attribute_definition_date_ref := range stage.A_ATTRIBUTE_DEFINITION_DATE_REFs {
		a_attribute_definition_date_ref.Unstage(stage)
	}

	for a_attribute_definition_enumeration_ref := range stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFs {
		a_attribute_definition_enumeration_ref.Unstage(stage)
	}

	for a_attribute_definition_integer_ref := range stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFs {
		a_attribute_definition_integer_ref.Unstage(stage)
	}

	for a_attribute_definition_real_ref := range stage.A_ATTRIBUTE_DEFINITION_REAL_REFs {
		a_attribute_definition_real_ref.Unstage(stage)
	}

	for a_attribute_definition_string_ref := range stage.A_ATTRIBUTE_DEFINITION_STRING_REFs {
		a_attribute_definition_string_ref.Unstage(stage)
	}

	for a_attribute_definition_xhtml_ref := range stage.A_ATTRIBUTE_DEFINITION_XHTML_REFs {
		a_attribute_definition_xhtml_ref.Unstage(stage)
	}

	for a_attribute_value_boolean := range stage.A_ATTRIBUTE_VALUE_BOOLEANs {
		a_attribute_value_boolean.Unstage(stage)
	}

	for a_attribute_value_date := range stage.A_ATTRIBUTE_VALUE_DATEs {
		a_attribute_value_date.Unstage(stage)
	}

	for a_attribute_value_enumeration := range stage.A_ATTRIBUTE_VALUE_ENUMERATIONs {
		a_attribute_value_enumeration.Unstage(stage)
	}

	for a_attribute_value_integer := range stage.A_ATTRIBUTE_VALUE_INTEGERs {
		a_attribute_value_integer.Unstage(stage)
	}

	for a_attribute_value_real := range stage.A_ATTRIBUTE_VALUE_REALs {
		a_attribute_value_real.Unstage(stage)
	}

	for a_attribute_value_string := range stage.A_ATTRIBUTE_VALUE_STRINGs {
		a_attribute_value_string.Unstage(stage)
	}

	for a_attribute_value_xhtml := range stage.A_ATTRIBUTE_VALUE_XHTMLs {
		a_attribute_value_xhtml.Unstage(stage)
	}

	for a_attribute_value_xhtml_1 := range stage.A_ATTRIBUTE_VALUE_XHTML_1s {
		a_attribute_value_xhtml_1.Unstage(stage)
	}

	for a_children := range stage.A_CHILDRENs {
		a_children.Unstage(stage)
	}

	for a_core_content := range stage.A_CORE_CONTENTs {
		a_core_content.Unstage(stage)
	}

	for a_datatypes := range stage.A_DATATYPESs {
		a_datatypes.Unstage(stage)
	}

	for a_datatype_definition_boolean_ref := range stage.A_DATATYPE_DEFINITION_BOOLEAN_REFs {
		a_datatype_definition_boolean_ref.Unstage(stage)
	}

	for a_datatype_definition_date_ref := range stage.A_DATATYPE_DEFINITION_DATE_REFs {
		a_datatype_definition_date_ref.Unstage(stage)
	}

	for a_datatype_definition_enumeration_ref := range stage.A_DATATYPE_DEFINITION_ENUMERATION_REFs {
		a_datatype_definition_enumeration_ref.Unstage(stage)
	}

	for a_datatype_definition_integer_ref := range stage.A_DATATYPE_DEFINITION_INTEGER_REFs {
		a_datatype_definition_integer_ref.Unstage(stage)
	}

	for a_datatype_definition_real_ref := range stage.A_DATATYPE_DEFINITION_REAL_REFs {
		a_datatype_definition_real_ref.Unstage(stage)
	}

	for a_datatype_definition_string_ref := range stage.A_DATATYPE_DEFINITION_STRING_REFs {
		a_datatype_definition_string_ref.Unstage(stage)
	}

	for a_datatype_definition_xhtml_ref := range stage.A_DATATYPE_DEFINITION_XHTML_REFs {
		a_datatype_definition_xhtml_ref.Unstage(stage)
	}

	for a_editable_atts := range stage.A_EDITABLE_ATTSs {
		a_editable_atts.Unstage(stage)
	}

	for a_enum_value_ref := range stage.A_ENUM_VALUE_REFs {
		a_enum_value_ref.Unstage(stage)
	}

	for a_object := range stage.A_OBJECTs {
		a_object.Unstage(stage)
	}

	for a_properties := range stage.A_PROPERTIESs {
		a_properties.Unstage(stage)
	}

	for a_relation_group_type_ref := range stage.A_RELATION_GROUP_TYPE_REFs {
		a_relation_group_type_ref.Unstage(stage)
	}

	for a_source_1 := range stage.A_SOURCE_1s {
		a_source_1.Unstage(stage)
	}

	for a_specifications := range stage.A_SPECIFICATIONSs {
		a_specifications.Unstage(stage)
	}

	for a_specification_type_ref := range stage.A_SPECIFICATION_TYPE_REFs {
		a_specification_type_ref.Unstage(stage)
	}

	for a_specified_values := range stage.A_SPECIFIED_VALUESs {
		a_specified_values.Unstage(stage)
	}

	for a_spec_attributes := range stage.A_SPEC_ATTRIBUTESs {
		a_spec_attributes.Unstage(stage)
	}

	for a_spec_objects := range stage.A_SPEC_OBJECTSs {
		a_spec_objects.Unstage(stage)
	}

	for a_spec_object_type_ref := range stage.A_SPEC_OBJECT_TYPE_REFs {
		a_spec_object_type_ref.Unstage(stage)
	}

	for a_spec_relations := range stage.A_SPEC_RELATIONSs {
		a_spec_relations.Unstage(stage)
	}

	for a_spec_relation_groups := range stage.A_SPEC_RELATION_GROUPSs {
		a_spec_relation_groups.Unstage(stage)
	}

	for a_spec_relation_ref := range stage.A_SPEC_RELATION_REFs {
		a_spec_relation_ref.Unstage(stage)
	}

	for a_spec_relation_type_ref := range stage.A_SPEC_RELATION_TYPE_REFs {
		a_spec_relation_type_ref.Unstage(stage)
	}

	for a_spec_types := range stage.A_SPEC_TYPESs {
		a_spec_types.Unstage(stage)
	}

	for a_target_specification_1 := range stage.A_TARGET_SPECIFICATION_1s {
		a_target_specification_1.Unstage(stage)
	}

	for a_the_header := range stage.A_THE_HEADERs {
		a_the_header.Unstage(stage)
	}

	for a_tool_extensions := range stage.A_TOOL_EXTENSIONSs {
		a_tool_extensions.Unstage(stage)
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
	case map[*A_ALTERNATIVE_ID]any:
		return any(&stage.A_ALTERNATIVE_IDs).(*Type)
	case map[*A_ATTRIBUTE_DEFINITION_BOOLEAN_REF]any:
		return any(&stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFs).(*Type)
	case map[*A_ATTRIBUTE_DEFINITION_DATE_REF]any:
		return any(&stage.A_ATTRIBUTE_DEFINITION_DATE_REFs).(*Type)
	case map[*A_ATTRIBUTE_DEFINITION_ENUMERATION_REF]any:
		return any(&stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFs).(*Type)
	case map[*A_ATTRIBUTE_DEFINITION_INTEGER_REF]any:
		return any(&stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFs).(*Type)
	case map[*A_ATTRIBUTE_DEFINITION_REAL_REF]any:
		return any(&stage.A_ATTRIBUTE_DEFINITION_REAL_REFs).(*Type)
	case map[*A_ATTRIBUTE_DEFINITION_STRING_REF]any:
		return any(&stage.A_ATTRIBUTE_DEFINITION_STRING_REFs).(*Type)
	case map[*A_ATTRIBUTE_DEFINITION_XHTML_REF]any:
		return any(&stage.A_ATTRIBUTE_DEFINITION_XHTML_REFs).(*Type)
	case map[*A_ATTRIBUTE_VALUE_BOOLEAN]any:
		return any(&stage.A_ATTRIBUTE_VALUE_BOOLEANs).(*Type)
	case map[*A_ATTRIBUTE_VALUE_DATE]any:
		return any(&stage.A_ATTRIBUTE_VALUE_DATEs).(*Type)
	case map[*A_ATTRIBUTE_VALUE_ENUMERATION]any:
		return any(&stage.A_ATTRIBUTE_VALUE_ENUMERATIONs).(*Type)
	case map[*A_ATTRIBUTE_VALUE_INTEGER]any:
		return any(&stage.A_ATTRIBUTE_VALUE_INTEGERs).(*Type)
	case map[*A_ATTRIBUTE_VALUE_REAL]any:
		return any(&stage.A_ATTRIBUTE_VALUE_REALs).(*Type)
	case map[*A_ATTRIBUTE_VALUE_STRING]any:
		return any(&stage.A_ATTRIBUTE_VALUE_STRINGs).(*Type)
	case map[*A_ATTRIBUTE_VALUE_XHTML]any:
		return any(&stage.A_ATTRIBUTE_VALUE_XHTMLs).(*Type)
	case map[*A_ATTRIBUTE_VALUE_XHTML_1]any:
		return any(&stage.A_ATTRIBUTE_VALUE_XHTML_1s).(*Type)
	case map[*A_CHILDREN]any:
		return any(&stage.A_CHILDRENs).(*Type)
	case map[*A_CORE_CONTENT]any:
		return any(&stage.A_CORE_CONTENTs).(*Type)
	case map[*A_DATATYPES]any:
		return any(&stage.A_DATATYPESs).(*Type)
	case map[*A_DATATYPE_DEFINITION_BOOLEAN_REF]any:
		return any(&stage.A_DATATYPE_DEFINITION_BOOLEAN_REFs).(*Type)
	case map[*A_DATATYPE_DEFINITION_DATE_REF]any:
		return any(&stage.A_DATATYPE_DEFINITION_DATE_REFs).(*Type)
	case map[*A_DATATYPE_DEFINITION_ENUMERATION_REF]any:
		return any(&stage.A_DATATYPE_DEFINITION_ENUMERATION_REFs).(*Type)
	case map[*A_DATATYPE_DEFINITION_INTEGER_REF]any:
		return any(&stage.A_DATATYPE_DEFINITION_INTEGER_REFs).(*Type)
	case map[*A_DATATYPE_DEFINITION_REAL_REF]any:
		return any(&stage.A_DATATYPE_DEFINITION_REAL_REFs).(*Type)
	case map[*A_DATATYPE_DEFINITION_STRING_REF]any:
		return any(&stage.A_DATATYPE_DEFINITION_STRING_REFs).(*Type)
	case map[*A_DATATYPE_DEFINITION_XHTML_REF]any:
		return any(&stage.A_DATATYPE_DEFINITION_XHTML_REFs).(*Type)
	case map[*A_EDITABLE_ATTS]any:
		return any(&stage.A_EDITABLE_ATTSs).(*Type)
	case map[*A_ENUM_VALUE_REF]any:
		return any(&stage.A_ENUM_VALUE_REFs).(*Type)
	case map[*A_OBJECT]any:
		return any(&stage.A_OBJECTs).(*Type)
	case map[*A_PROPERTIES]any:
		return any(&stage.A_PROPERTIESs).(*Type)
	case map[*A_RELATION_GROUP_TYPE_REF]any:
		return any(&stage.A_RELATION_GROUP_TYPE_REFs).(*Type)
	case map[*A_SOURCE_1]any:
		return any(&stage.A_SOURCE_1s).(*Type)
	case map[*A_SPECIFICATIONS]any:
		return any(&stage.A_SPECIFICATIONSs).(*Type)
	case map[*A_SPECIFICATION_TYPE_REF]any:
		return any(&stage.A_SPECIFICATION_TYPE_REFs).(*Type)
	case map[*A_SPECIFIED_VALUES]any:
		return any(&stage.A_SPECIFIED_VALUESs).(*Type)
	case map[*A_SPEC_ATTRIBUTES]any:
		return any(&stage.A_SPEC_ATTRIBUTESs).(*Type)
	case map[*A_SPEC_OBJECTS]any:
		return any(&stage.A_SPEC_OBJECTSs).(*Type)
	case map[*A_SPEC_OBJECT_TYPE_REF]any:
		return any(&stage.A_SPEC_OBJECT_TYPE_REFs).(*Type)
	case map[*A_SPEC_RELATIONS]any:
		return any(&stage.A_SPEC_RELATIONSs).(*Type)
	case map[*A_SPEC_RELATION_GROUPS]any:
		return any(&stage.A_SPEC_RELATION_GROUPSs).(*Type)
	case map[*A_SPEC_RELATION_REF]any:
		return any(&stage.A_SPEC_RELATION_REFs).(*Type)
	case map[*A_SPEC_RELATION_TYPE_REF]any:
		return any(&stage.A_SPEC_RELATION_TYPE_REFs).(*Type)
	case map[*A_SPEC_TYPES]any:
		return any(&stage.A_SPEC_TYPESs).(*Type)
	case map[*A_TARGET_SPECIFICATION_1]any:
		return any(&stage.A_TARGET_SPECIFICATION_1s).(*Type)
	case map[*A_THE_HEADER]any:
		return any(&stage.A_THE_HEADERs).(*Type)
	case map[*A_TOOL_EXTENSIONS]any:
		return any(&stage.A_TOOL_EXTENSIONSs).(*Type)
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
	case map[string]*A_ALTERNATIVE_ID:
		return any(&stage.A_ALTERNATIVE_IDs_mapString).(*Type)
	case map[string]*A_ATTRIBUTE_DEFINITION_BOOLEAN_REF:
		return any(&stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFs_mapString).(*Type)
	case map[string]*A_ATTRIBUTE_DEFINITION_DATE_REF:
		return any(&stage.A_ATTRIBUTE_DEFINITION_DATE_REFs_mapString).(*Type)
	case map[string]*A_ATTRIBUTE_DEFINITION_ENUMERATION_REF:
		return any(&stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFs_mapString).(*Type)
	case map[string]*A_ATTRIBUTE_DEFINITION_INTEGER_REF:
		return any(&stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFs_mapString).(*Type)
	case map[string]*A_ATTRIBUTE_DEFINITION_REAL_REF:
		return any(&stage.A_ATTRIBUTE_DEFINITION_REAL_REFs_mapString).(*Type)
	case map[string]*A_ATTRIBUTE_DEFINITION_STRING_REF:
		return any(&stage.A_ATTRIBUTE_DEFINITION_STRING_REFs_mapString).(*Type)
	case map[string]*A_ATTRIBUTE_DEFINITION_XHTML_REF:
		return any(&stage.A_ATTRIBUTE_DEFINITION_XHTML_REFs_mapString).(*Type)
	case map[string]*A_ATTRIBUTE_VALUE_BOOLEAN:
		return any(&stage.A_ATTRIBUTE_VALUE_BOOLEANs_mapString).(*Type)
	case map[string]*A_ATTRIBUTE_VALUE_DATE:
		return any(&stage.A_ATTRIBUTE_VALUE_DATEs_mapString).(*Type)
	case map[string]*A_ATTRIBUTE_VALUE_ENUMERATION:
		return any(&stage.A_ATTRIBUTE_VALUE_ENUMERATIONs_mapString).(*Type)
	case map[string]*A_ATTRIBUTE_VALUE_INTEGER:
		return any(&stage.A_ATTRIBUTE_VALUE_INTEGERs_mapString).(*Type)
	case map[string]*A_ATTRIBUTE_VALUE_REAL:
		return any(&stage.A_ATTRIBUTE_VALUE_REALs_mapString).(*Type)
	case map[string]*A_ATTRIBUTE_VALUE_STRING:
		return any(&stage.A_ATTRIBUTE_VALUE_STRINGs_mapString).(*Type)
	case map[string]*A_ATTRIBUTE_VALUE_XHTML:
		return any(&stage.A_ATTRIBUTE_VALUE_XHTMLs_mapString).(*Type)
	case map[string]*A_ATTRIBUTE_VALUE_XHTML_1:
		return any(&stage.A_ATTRIBUTE_VALUE_XHTML_1s_mapString).(*Type)
	case map[string]*A_CHILDREN:
		return any(&stage.A_CHILDRENs_mapString).(*Type)
	case map[string]*A_CORE_CONTENT:
		return any(&stage.A_CORE_CONTENTs_mapString).(*Type)
	case map[string]*A_DATATYPES:
		return any(&stage.A_DATATYPESs_mapString).(*Type)
	case map[string]*A_DATATYPE_DEFINITION_BOOLEAN_REF:
		return any(&stage.A_DATATYPE_DEFINITION_BOOLEAN_REFs_mapString).(*Type)
	case map[string]*A_DATATYPE_DEFINITION_DATE_REF:
		return any(&stage.A_DATATYPE_DEFINITION_DATE_REFs_mapString).(*Type)
	case map[string]*A_DATATYPE_DEFINITION_ENUMERATION_REF:
		return any(&stage.A_DATATYPE_DEFINITION_ENUMERATION_REFs_mapString).(*Type)
	case map[string]*A_DATATYPE_DEFINITION_INTEGER_REF:
		return any(&stage.A_DATATYPE_DEFINITION_INTEGER_REFs_mapString).(*Type)
	case map[string]*A_DATATYPE_DEFINITION_REAL_REF:
		return any(&stage.A_DATATYPE_DEFINITION_REAL_REFs_mapString).(*Type)
	case map[string]*A_DATATYPE_DEFINITION_STRING_REF:
		return any(&stage.A_DATATYPE_DEFINITION_STRING_REFs_mapString).(*Type)
	case map[string]*A_DATATYPE_DEFINITION_XHTML_REF:
		return any(&stage.A_DATATYPE_DEFINITION_XHTML_REFs_mapString).(*Type)
	case map[string]*A_EDITABLE_ATTS:
		return any(&stage.A_EDITABLE_ATTSs_mapString).(*Type)
	case map[string]*A_ENUM_VALUE_REF:
		return any(&stage.A_ENUM_VALUE_REFs_mapString).(*Type)
	case map[string]*A_OBJECT:
		return any(&stage.A_OBJECTs_mapString).(*Type)
	case map[string]*A_PROPERTIES:
		return any(&stage.A_PROPERTIESs_mapString).(*Type)
	case map[string]*A_RELATION_GROUP_TYPE_REF:
		return any(&stage.A_RELATION_GROUP_TYPE_REFs_mapString).(*Type)
	case map[string]*A_SOURCE_1:
		return any(&stage.A_SOURCE_1s_mapString).(*Type)
	case map[string]*A_SPECIFICATIONS:
		return any(&stage.A_SPECIFICATIONSs_mapString).(*Type)
	case map[string]*A_SPECIFICATION_TYPE_REF:
		return any(&stage.A_SPECIFICATION_TYPE_REFs_mapString).(*Type)
	case map[string]*A_SPECIFIED_VALUES:
		return any(&stage.A_SPECIFIED_VALUESs_mapString).(*Type)
	case map[string]*A_SPEC_ATTRIBUTES:
		return any(&stage.A_SPEC_ATTRIBUTESs_mapString).(*Type)
	case map[string]*A_SPEC_OBJECTS:
		return any(&stage.A_SPEC_OBJECTSs_mapString).(*Type)
	case map[string]*A_SPEC_OBJECT_TYPE_REF:
		return any(&stage.A_SPEC_OBJECT_TYPE_REFs_mapString).(*Type)
	case map[string]*A_SPEC_RELATIONS:
		return any(&stage.A_SPEC_RELATIONSs_mapString).(*Type)
	case map[string]*A_SPEC_RELATION_GROUPS:
		return any(&stage.A_SPEC_RELATION_GROUPSs_mapString).(*Type)
	case map[string]*A_SPEC_RELATION_REF:
		return any(&stage.A_SPEC_RELATION_REFs_mapString).(*Type)
	case map[string]*A_SPEC_RELATION_TYPE_REF:
		return any(&stage.A_SPEC_RELATION_TYPE_REFs_mapString).(*Type)
	case map[string]*A_SPEC_TYPES:
		return any(&stage.A_SPEC_TYPESs_mapString).(*Type)
	case map[string]*A_TARGET_SPECIFICATION_1:
		return any(&stage.A_TARGET_SPECIFICATION_1s_mapString).(*Type)
	case map[string]*A_THE_HEADER:
		return any(&stage.A_THE_HEADERs_mapString).(*Type)
	case map[string]*A_TOOL_EXTENSIONS:
		return any(&stage.A_TOOL_EXTENSIONSs_mapString).(*Type)
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
	case A_ALTERNATIVE_ID:
		return any(&stage.A_ALTERNATIVE_IDs).(*map[*Type]any)
	case A_ATTRIBUTE_DEFINITION_BOOLEAN_REF:
		return any(&stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFs).(*map[*Type]any)
	case A_ATTRIBUTE_DEFINITION_DATE_REF:
		return any(&stage.A_ATTRIBUTE_DEFINITION_DATE_REFs).(*map[*Type]any)
	case A_ATTRIBUTE_DEFINITION_ENUMERATION_REF:
		return any(&stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFs).(*map[*Type]any)
	case A_ATTRIBUTE_DEFINITION_INTEGER_REF:
		return any(&stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFs).(*map[*Type]any)
	case A_ATTRIBUTE_DEFINITION_REAL_REF:
		return any(&stage.A_ATTRIBUTE_DEFINITION_REAL_REFs).(*map[*Type]any)
	case A_ATTRIBUTE_DEFINITION_STRING_REF:
		return any(&stage.A_ATTRIBUTE_DEFINITION_STRING_REFs).(*map[*Type]any)
	case A_ATTRIBUTE_DEFINITION_XHTML_REF:
		return any(&stage.A_ATTRIBUTE_DEFINITION_XHTML_REFs).(*map[*Type]any)
	case A_ATTRIBUTE_VALUE_BOOLEAN:
		return any(&stage.A_ATTRIBUTE_VALUE_BOOLEANs).(*map[*Type]any)
	case A_ATTRIBUTE_VALUE_DATE:
		return any(&stage.A_ATTRIBUTE_VALUE_DATEs).(*map[*Type]any)
	case A_ATTRIBUTE_VALUE_ENUMERATION:
		return any(&stage.A_ATTRIBUTE_VALUE_ENUMERATIONs).(*map[*Type]any)
	case A_ATTRIBUTE_VALUE_INTEGER:
		return any(&stage.A_ATTRIBUTE_VALUE_INTEGERs).(*map[*Type]any)
	case A_ATTRIBUTE_VALUE_REAL:
		return any(&stage.A_ATTRIBUTE_VALUE_REALs).(*map[*Type]any)
	case A_ATTRIBUTE_VALUE_STRING:
		return any(&stage.A_ATTRIBUTE_VALUE_STRINGs).(*map[*Type]any)
	case A_ATTRIBUTE_VALUE_XHTML:
		return any(&stage.A_ATTRIBUTE_VALUE_XHTMLs).(*map[*Type]any)
	case A_ATTRIBUTE_VALUE_XHTML_1:
		return any(&stage.A_ATTRIBUTE_VALUE_XHTML_1s).(*map[*Type]any)
	case A_CHILDREN:
		return any(&stage.A_CHILDRENs).(*map[*Type]any)
	case A_CORE_CONTENT:
		return any(&stage.A_CORE_CONTENTs).(*map[*Type]any)
	case A_DATATYPES:
		return any(&stage.A_DATATYPESs).(*map[*Type]any)
	case A_DATATYPE_DEFINITION_BOOLEAN_REF:
		return any(&stage.A_DATATYPE_DEFINITION_BOOLEAN_REFs).(*map[*Type]any)
	case A_DATATYPE_DEFINITION_DATE_REF:
		return any(&stage.A_DATATYPE_DEFINITION_DATE_REFs).(*map[*Type]any)
	case A_DATATYPE_DEFINITION_ENUMERATION_REF:
		return any(&stage.A_DATATYPE_DEFINITION_ENUMERATION_REFs).(*map[*Type]any)
	case A_DATATYPE_DEFINITION_INTEGER_REF:
		return any(&stage.A_DATATYPE_DEFINITION_INTEGER_REFs).(*map[*Type]any)
	case A_DATATYPE_DEFINITION_REAL_REF:
		return any(&stage.A_DATATYPE_DEFINITION_REAL_REFs).(*map[*Type]any)
	case A_DATATYPE_DEFINITION_STRING_REF:
		return any(&stage.A_DATATYPE_DEFINITION_STRING_REFs).(*map[*Type]any)
	case A_DATATYPE_DEFINITION_XHTML_REF:
		return any(&stage.A_DATATYPE_DEFINITION_XHTML_REFs).(*map[*Type]any)
	case A_EDITABLE_ATTS:
		return any(&stage.A_EDITABLE_ATTSs).(*map[*Type]any)
	case A_ENUM_VALUE_REF:
		return any(&stage.A_ENUM_VALUE_REFs).(*map[*Type]any)
	case A_OBJECT:
		return any(&stage.A_OBJECTs).(*map[*Type]any)
	case A_PROPERTIES:
		return any(&stage.A_PROPERTIESs).(*map[*Type]any)
	case A_RELATION_GROUP_TYPE_REF:
		return any(&stage.A_RELATION_GROUP_TYPE_REFs).(*map[*Type]any)
	case A_SOURCE_1:
		return any(&stage.A_SOURCE_1s).(*map[*Type]any)
	case A_SPECIFICATIONS:
		return any(&stage.A_SPECIFICATIONSs).(*map[*Type]any)
	case A_SPECIFICATION_TYPE_REF:
		return any(&stage.A_SPECIFICATION_TYPE_REFs).(*map[*Type]any)
	case A_SPECIFIED_VALUES:
		return any(&stage.A_SPECIFIED_VALUESs).(*map[*Type]any)
	case A_SPEC_ATTRIBUTES:
		return any(&stage.A_SPEC_ATTRIBUTESs).(*map[*Type]any)
	case A_SPEC_OBJECTS:
		return any(&stage.A_SPEC_OBJECTSs).(*map[*Type]any)
	case A_SPEC_OBJECT_TYPE_REF:
		return any(&stage.A_SPEC_OBJECT_TYPE_REFs).(*map[*Type]any)
	case A_SPEC_RELATIONS:
		return any(&stage.A_SPEC_RELATIONSs).(*map[*Type]any)
	case A_SPEC_RELATION_GROUPS:
		return any(&stage.A_SPEC_RELATION_GROUPSs).(*map[*Type]any)
	case A_SPEC_RELATION_REF:
		return any(&stage.A_SPEC_RELATION_REFs).(*map[*Type]any)
	case A_SPEC_RELATION_TYPE_REF:
		return any(&stage.A_SPEC_RELATION_TYPE_REFs).(*map[*Type]any)
	case A_SPEC_TYPES:
		return any(&stage.A_SPEC_TYPESs).(*map[*Type]any)
	case A_TARGET_SPECIFICATION_1:
		return any(&stage.A_TARGET_SPECIFICATION_1s).(*map[*Type]any)
	case A_THE_HEADER:
		return any(&stage.A_THE_HEADERs).(*map[*Type]any)
	case A_TOOL_EXTENSIONS:
		return any(&stage.A_TOOL_EXTENSIONSs).(*map[*Type]any)
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
	case *A_ALTERNATIVE_ID:
		return any(&stage.A_ALTERNATIVE_IDs).(*map[Type]any)
	case *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF:
		return any(&stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFs).(*map[Type]any)
	case *A_ATTRIBUTE_DEFINITION_DATE_REF:
		return any(&stage.A_ATTRIBUTE_DEFINITION_DATE_REFs).(*map[Type]any)
	case *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF:
		return any(&stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFs).(*map[Type]any)
	case *A_ATTRIBUTE_DEFINITION_INTEGER_REF:
		return any(&stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFs).(*map[Type]any)
	case *A_ATTRIBUTE_DEFINITION_REAL_REF:
		return any(&stage.A_ATTRIBUTE_DEFINITION_REAL_REFs).(*map[Type]any)
	case *A_ATTRIBUTE_DEFINITION_STRING_REF:
		return any(&stage.A_ATTRIBUTE_DEFINITION_STRING_REFs).(*map[Type]any)
	case *A_ATTRIBUTE_DEFINITION_XHTML_REF:
		return any(&stage.A_ATTRIBUTE_DEFINITION_XHTML_REFs).(*map[Type]any)
	case *A_ATTRIBUTE_VALUE_BOOLEAN:
		return any(&stage.A_ATTRIBUTE_VALUE_BOOLEANs).(*map[Type]any)
	case *A_ATTRIBUTE_VALUE_DATE:
		return any(&stage.A_ATTRIBUTE_VALUE_DATEs).(*map[Type]any)
	case *A_ATTRIBUTE_VALUE_ENUMERATION:
		return any(&stage.A_ATTRIBUTE_VALUE_ENUMERATIONs).(*map[Type]any)
	case *A_ATTRIBUTE_VALUE_INTEGER:
		return any(&stage.A_ATTRIBUTE_VALUE_INTEGERs).(*map[Type]any)
	case *A_ATTRIBUTE_VALUE_REAL:
		return any(&stage.A_ATTRIBUTE_VALUE_REALs).(*map[Type]any)
	case *A_ATTRIBUTE_VALUE_STRING:
		return any(&stage.A_ATTRIBUTE_VALUE_STRINGs).(*map[Type]any)
	case *A_ATTRIBUTE_VALUE_XHTML:
		return any(&stage.A_ATTRIBUTE_VALUE_XHTMLs).(*map[Type]any)
	case *A_ATTRIBUTE_VALUE_XHTML_1:
		return any(&stage.A_ATTRIBUTE_VALUE_XHTML_1s).(*map[Type]any)
	case *A_CHILDREN:
		return any(&stage.A_CHILDRENs).(*map[Type]any)
	case *A_CORE_CONTENT:
		return any(&stage.A_CORE_CONTENTs).(*map[Type]any)
	case *A_DATATYPES:
		return any(&stage.A_DATATYPESs).(*map[Type]any)
	case *A_DATATYPE_DEFINITION_BOOLEAN_REF:
		return any(&stage.A_DATATYPE_DEFINITION_BOOLEAN_REFs).(*map[Type]any)
	case *A_DATATYPE_DEFINITION_DATE_REF:
		return any(&stage.A_DATATYPE_DEFINITION_DATE_REFs).(*map[Type]any)
	case *A_DATATYPE_DEFINITION_ENUMERATION_REF:
		return any(&stage.A_DATATYPE_DEFINITION_ENUMERATION_REFs).(*map[Type]any)
	case *A_DATATYPE_DEFINITION_INTEGER_REF:
		return any(&stage.A_DATATYPE_DEFINITION_INTEGER_REFs).(*map[Type]any)
	case *A_DATATYPE_DEFINITION_REAL_REF:
		return any(&stage.A_DATATYPE_DEFINITION_REAL_REFs).(*map[Type]any)
	case *A_DATATYPE_DEFINITION_STRING_REF:
		return any(&stage.A_DATATYPE_DEFINITION_STRING_REFs).(*map[Type]any)
	case *A_DATATYPE_DEFINITION_XHTML_REF:
		return any(&stage.A_DATATYPE_DEFINITION_XHTML_REFs).(*map[Type]any)
	case *A_EDITABLE_ATTS:
		return any(&stage.A_EDITABLE_ATTSs).(*map[Type]any)
	case *A_ENUM_VALUE_REF:
		return any(&stage.A_ENUM_VALUE_REFs).(*map[Type]any)
	case *A_OBJECT:
		return any(&stage.A_OBJECTs).(*map[Type]any)
	case *A_PROPERTIES:
		return any(&stage.A_PROPERTIESs).(*map[Type]any)
	case *A_RELATION_GROUP_TYPE_REF:
		return any(&stage.A_RELATION_GROUP_TYPE_REFs).(*map[Type]any)
	case *A_SOURCE_1:
		return any(&stage.A_SOURCE_1s).(*map[Type]any)
	case *A_SPECIFICATIONS:
		return any(&stage.A_SPECIFICATIONSs).(*map[Type]any)
	case *A_SPECIFICATION_TYPE_REF:
		return any(&stage.A_SPECIFICATION_TYPE_REFs).(*map[Type]any)
	case *A_SPECIFIED_VALUES:
		return any(&stage.A_SPECIFIED_VALUESs).(*map[Type]any)
	case *A_SPEC_ATTRIBUTES:
		return any(&stage.A_SPEC_ATTRIBUTESs).(*map[Type]any)
	case *A_SPEC_OBJECTS:
		return any(&stage.A_SPEC_OBJECTSs).(*map[Type]any)
	case *A_SPEC_OBJECT_TYPE_REF:
		return any(&stage.A_SPEC_OBJECT_TYPE_REFs).(*map[Type]any)
	case *A_SPEC_RELATIONS:
		return any(&stage.A_SPEC_RELATIONSs).(*map[Type]any)
	case *A_SPEC_RELATION_GROUPS:
		return any(&stage.A_SPEC_RELATION_GROUPSs).(*map[Type]any)
	case *A_SPEC_RELATION_REF:
		return any(&stage.A_SPEC_RELATION_REFs).(*map[Type]any)
	case *A_SPEC_RELATION_TYPE_REF:
		return any(&stage.A_SPEC_RELATION_TYPE_REFs).(*map[Type]any)
	case *A_SPEC_TYPES:
		return any(&stage.A_SPEC_TYPESs).(*map[Type]any)
	case *A_TARGET_SPECIFICATION_1:
		return any(&stage.A_TARGET_SPECIFICATION_1s).(*map[Type]any)
	case *A_THE_HEADER:
		return any(&stage.A_THE_HEADERs).(*map[Type]any)
	case *A_TOOL_EXTENSIONS:
		return any(&stage.A_TOOL_EXTENSIONSs).(*map[Type]any)
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
	case A_ALTERNATIVE_ID:
		return any(&stage.A_ALTERNATIVE_IDs_mapString).(*map[string]*Type)
	case A_ATTRIBUTE_DEFINITION_BOOLEAN_REF:
		return any(&stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFs_mapString).(*map[string]*Type)
	case A_ATTRIBUTE_DEFINITION_DATE_REF:
		return any(&stage.A_ATTRIBUTE_DEFINITION_DATE_REFs_mapString).(*map[string]*Type)
	case A_ATTRIBUTE_DEFINITION_ENUMERATION_REF:
		return any(&stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFs_mapString).(*map[string]*Type)
	case A_ATTRIBUTE_DEFINITION_INTEGER_REF:
		return any(&stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFs_mapString).(*map[string]*Type)
	case A_ATTRIBUTE_DEFINITION_REAL_REF:
		return any(&stage.A_ATTRIBUTE_DEFINITION_REAL_REFs_mapString).(*map[string]*Type)
	case A_ATTRIBUTE_DEFINITION_STRING_REF:
		return any(&stage.A_ATTRIBUTE_DEFINITION_STRING_REFs_mapString).(*map[string]*Type)
	case A_ATTRIBUTE_DEFINITION_XHTML_REF:
		return any(&stage.A_ATTRIBUTE_DEFINITION_XHTML_REFs_mapString).(*map[string]*Type)
	case A_ATTRIBUTE_VALUE_BOOLEAN:
		return any(&stage.A_ATTRIBUTE_VALUE_BOOLEANs_mapString).(*map[string]*Type)
	case A_ATTRIBUTE_VALUE_DATE:
		return any(&stage.A_ATTRIBUTE_VALUE_DATEs_mapString).(*map[string]*Type)
	case A_ATTRIBUTE_VALUE_ENUMERATION:
		return any(&stage.A_ATTRIBUTE_VALUE_ENUMERATIONs_mapString).(*map[string]*Type)
	case A_ATTRIBUTE_VALUE_INTEGER:
		return any(&stage.A_ATTRIBUTE_VALUE_INTEGERs_mapString).(*map[string]*Type)
	case A_ATTRIBUTE_VALUE_REAL:
		return any(&stage.A_ATTRIBUTE_VALUE_REALs_mapString).(*map[string]*Type)
	case A_ATTRIBUTE_VALUE_STRING:
		return any(&stage.A_ATTRIBUTE_VALUE_STRINGs_mapString).(*map[string]*Type)
	case A_ATTRIBUTE_VALUE_XHTML:
		return any(&stage.A_ATTRIBUTE_VALUE_XHTMLs_mapString).(*map[string]*Type)
	case A_ATTRIBUTE_VALUE_XHTML_1:
		return any(&stage.A_ATTRIBUTE_VALUE_XHTML_1s_mapString).(*map[string]*Type)
	case A_CHILDREN:
		return any(&stage.A_CHILDRENs_mapString).(*map[string]*Type)
	case A_CORE_CONTENT:
		return any(&stage.A_CORE_CONTENTs_mapString).(*map[string]*Type)
	case A_DATATYPES:
		return any(&stage.A_DATATYPESs_mapString).(*map[string]*Type)
	case A_DATATYPE_DEFINITION_BOOLEAN_REF:
		return any(&stage.A_DATATYPE_DEFINITION_BOOLEAN_REFs_mapString).(*map[string]*Type)
	case A_DATATYPE_DEFINITION_DATE_REF:
		return any(&stage.A_DATATYPE_DEFINITION_DATE_REFs_mapString).(*map[string]*Type)
	case A_DATATYPE_DEFINITION_ENUMERATION_REF:
		return any(&stage.A_DATATYPE_DEFINITION_ENUMERATION_REFs_mapString).(*map[string]*Type)
	case A_DATATYPE_DEFINITION_INTEGER_REF:
		return any(&stage.A_DATATYPE_DEFINITION_INTEGER_REFs_mapString).(*map[string]*Type)
	case A_DATATYPE_DEFINITION_REAL_REF:
		return any(&stage.A_DATATYPE_DEFINITION_REAL_REFs_mapString).(*map[string]*Type)
	case A_DATATYPE_DEFINITION_STRING_REF:
		return any(&stage.A_DATATYPE_DEFINITION_STRING_REFs_mapString).(*map[string]*Type)
	case A_DATATYPE_DEFINITION_XHTML_REF:
		return any(&stage.A_DATATYPE_DEFINITION_XHTML_REFs_mapString).(*map[string]*Type)
	case A_EDITABLE_ATTS:
		return any(&stage.A_EDITABLE_ATTSs_mapString).(*map[string]*Type)
	case A_ENUM_VALUE_REF:
		return any(&stage.A_ENUM_VALUE_REFs_mapString).(*map[string]*Type)
	case A_OBJECT:
		return any(&stage.A_OBJECTs_mapString).(*map[string]*Type)
	case A_PROPERTIES:
		return any(&stage.A_PROPERTIESs_mapString).(*map[string]*Type)
	case A_RELATION_GROUP_TYPE_REF:
		return any(&stage.A_RELATION_GROUP_TYPE_REFs_mapString).(*map[string]*Type)
	case A_SOURCE_1:
		return any(&stage.A_SOURCE_1s_mapString).(*map[string]*Type)
	case A_SPECIFICATIONS:
		return any(&stage.A_SPECIFICATIONSs_mapString).(*map[string]*Type)
	case A_SPECIFICATION_TYPE_REF:
		return any(&stage.A_SPECIFICATION_TYPE_REFs_mapString).(*map[string]*Type)
	case A_SPECIFIED_VALUES:
		return any(&stage.A_SPECIFIED_VALUESs_mapString).(*map[string]*Type)
	case A_SPEC_ATTRIBUTES:
		return any(&stage.A_SPEC_ATTRIBUTESs_mapString).(*map[string]*Type)
	case A_SPEC_OBJECTS:
		return any(&stage.A_SPEC_OBJECTSs_mapString).(*map[string]*Type)
	case A_SPEC_OBJECT_TYPE_REF:
		return any(&stage.A_SPEC_OBJECT_TYPE_REFs_mapString).(*map[string]*Type)
	case A_SPEC_RELATIONS:
		return any(&stage.A_SPEC_RELATIONSs_mapString).(*map[string]*Type)
	case A_SPEC_RELATION_GROUPS:
		return any(&stage.A_SPEC_RELATION_GROUPSs_mapString).(*map[string]*Type)
	case A_SPEC_RELATION_REF:
		return any(&stage.A_SPEC_RELATION_REFs_mapString).(*map[string]*Type)
	case A_SPEC_RELATION_TYPE_REF:
		return any(&stage.A_SPEC_RELATION_TYPE_REFs_mapString).(*map[string]*Type)
	case A_SPEC_TYPES:
		return any(&stage.A_SPEC_TYPESs_mapString).(*map[string]*Type)
	case A_TARGET_SPECIFICATION_1:
		return any(&stage.A_TARGET_SPECIFICATION_1s_mapString).(*map[string]*Type)
	case A_THE_HEADER:
		return any(&stage.A_THE_HEADERs_mapString).(*map[string]*Type)
	case A_TOOL_EXTENSIONS:
		return any(&stage.A_TOOL_EXTENSIONSs_mapString).(*map[string]*Type)
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
			// field is initialized with an instance of A_ALTERNATIVE_ID with the name of the field
			ALTERNATIVE_ID: &A_ALTERNATIVE_ID{Name: "ALTERNATIVE_ID"},
			// field is initialized with an instance of A_ATTRIBUTE_VALUE_BOOLEAN with the name of the field
			DEFAULT_VALUE: &A_ATTRIBUTE_VALUE_BOOLEAN{Name: "DEFAULT_VALUE"},
			// field is initialized with an instance of A_DATATYPE_DEFINITION_BOOLEAN_REF with the name of the field
			TYPE: &A_DATATYPE_DEFINITION_BOOLEAN_REF{Name: "TYPE"},
		}).(*Type)
	case ATTRIBUTE_DEFINITION_DATE:
		return any(&ATTRIBUTE_DEFINITION_DATE{
			// Initialisation of associations
			// field is initialized with an instance of A_ALTERNATIVE_ID with the name of the field
			ALTERNATIVE_ID: &A_ALTERNATIVE_ID{Name: "ALTERNATIVE_ID"},
			// field is initialized with an instance of A_ATTRIBUTE_VALUE_DATE with the name of the field
			DEFAULT_VALUE: &A_ATTRIBUTE_VALUE_DATE{Name: "DEFAULT_VALUE"},
			// field is initialized with an instance of A_DATATYPE_DEFINITION_DATE_REF with the name of the field
			TYPE: &A_DATATYPE_DEFINITION_DATE_REF{Name: "TYPE"},
		}).(*Type)
	case ATTRIBUTE_DEFINITION_ENUMERATION:
		return any(&ATTRIBUTE_DEFINITION_ENUMERATION{
			// Initialisation of associations
			// field is initialized with an instance of A_ALTERNATIVE_ID with the name of the field
			ALTERNATIVE_ID: &A_ALTERNATIVE_ID{Name: "ALTERNATIVE_ID"},
			// field is initialized with an instance of A_ATTRIBUTE_VALUE_ENUMERATION with the name of the field
			DEFAULT_VALUE: &A_ATTRIBUTE_VALUE_ENUMERATION{Name: "DEFAULT_VALUE"},
			// field is initialized with an instance of A_DATATYPE_DEFINITION_ENUMERATION_REF with the name of the field
			TYPE: &A_DATATYPE_DEFINITION_ENUMERATION_REF{Name: "TYPE"},
		}).(*Type)
	case ATTRIBUTE_DEFINITION_INTEGER:
		return any(&ATTRIBUTE_DEFINITION_INTEGER{
			// Initialisation of associations
			// field is initialized with an instance of A_ALTERNATIVE_ID with the name of the field
			ALTERNATIVE_ID: &A_ALTERNATIVE_ID{Name: "ALTERNATIVE_ID"},
			// field is initialized with an instance of A_ATTRIBUTE_VALUE_INTEGER with the name of the field
			DEFAULT_VALUE: &A_ATTRIBUTE_VALUE_INTEGER{Name: "DEFAULT_VALUE"},
			// field is initialized with an instance of A_DATATYPE_DEFINITION_INTEGER_REF with the name of the field
			TYPE: &A_DATATYPE_DEFINITION_INTEGER_REF{Name: "TYPE"},
		}).(*Type)
	case ATTRIBUTE_DEFINITION_REAL:
		return any(&ATTRIBUTE_DEFINITION_REAL{
			// Initialisation of associations
			// field is initialized with an instance of A_ALTERNATIVE_ID with the name of the field
			ALTERNATIVE_ID: &A_ALTERNATIVE_ID{Name: "ALTERNATIVE_ID"},
			// field is initialized with an instance of A_ATTRIBUTE_VALUE_REAL with the name of the field
			DEFAULT_VALUE: &A_ATTRIBUTE_VALUE_REAL{Name: "DEFAULT_VALUE"},
			// field is initialized with an instance of A_DATATYPE_DEFINITION_REAL_REF with the name of the field
			TYPE: &A_DATATYPE_DEFINITION_REAL_REF{Name: "TYPE"},
		}).(*Type)
	case ATTRIBUTE_DEFINITION_STRING:
		return any(&ATTRIBUTE_DEFINITION_STRING{
			// Initialisation of associations
			// field is initialized with an instance of A_ALTERNATIVE_ID with the name of the field
			ALTERNATIVE_ID: &A_ALTERNATIVE_ID{Name: "ALTERNATIVE_ID"},
			// field is initialized with an instance of A_ATTRIBUTE_VALUE_STRING with the name of the field
			DEFAULT_VALUE: &A_ATTRIBUTE_VALUE_STRING{Name: "DEFAULT_VALUE"},
			// field is initialized with an instance of A_DATATYPE_DEFINITION_STRING_REF with the name of the field
			TYPE: &A_DATATYPE_DEFINITION_STRING_REF{Name: "TYPE"},
		}).(*Type)
	case ATTRIBUTE_DEFINITION_XHTML:
		return any(&ATTRIBUTE_DEFINITION_XHTML{
			// Initialisation of associations
			// field is initialized with an instance of A_ALTERNATIVE_ID with the name of the field
			ALTERNATIVE_ID: &A_ALTERNATIVE_ID{Name: "ALTERNATIVE_ID"},
			// field is initialized with an instance of A_ATTRIBUTE_VALUE_XHTML_1 with the name of the field
			DEFAULT_VALUE: &A_ATTRIBUTE_VALUE_XHTML_1{Name: "DEFAULT_VALUE"},
			// field is initialized with an instance of A_DATATYPE_DEFINITION_XHTML_REF with the name of the field
			TYPE: &A_DATATYPE_DEFINITION_XHTML_REF{Name: "TYPE"},
		}).(*Type)
	case ATTRIBUTE_VALUE_BOOLEAN:
		return any(&ATTRIBUTE_VALUE_BOOLEAN{
			// Initialisation of associations
			// field is initialized with an instance of A_ATTRIBUTE_DEFINITION_BOOLEAN_REF with the name of the field
			DEFINITION: &A_ATTRIBUTE_DEFINITION_BOOLEAN_REF{Name: "DEFINITION"},
		}).(*Type)
	case ATTRIBUTE_VALUE_DATE:
		return any(&ATTRIBUTE_VALUE_DATE{
			// Initialisation of associations
			// field is initialized with an instance of A_ATTRIBUTE_DEFINITION_DATE_REF with the name of the field
			DEFINITION: &A_ATTRIBUTE_DEFINITION_DATE_REF{Name: "DEFINITION"},
		}).(*Type)
	case ATTRIBUTE_VALUE_ENUMERATION:
		return any(&ATTRIBUTE_VALUE_ENUMERATION{
			// Initialisation of associations
			// field is initialized with an instance of A_ATTRIBUTE_DEFINITION_ENUMERATION_REF with the name of the field
			DEFINITION: &A_ATTRIBUTE_DEFINITION_ENUMERATION_REF{Name: "DEFINITION"},
			// field is initialized with an instance of A_ENUM_VALUE_REF with the name of the field
			VALUES: &A_ENUM_VALUE_REF{Name: "VALUES"},
		}).(*Type)
	case ATTRIBUTE_VALUE_INTEGER:
		return any(&ATTRIBUTE_VALUE_INTEGER{
			// Initialisation of associations
			// field is initialized with an instance of A_ATTRIBUTE_DEFINITION_INTEGER_REF with the name of the field
			DEFINITION: &A_ATTRIBUTE_DEFINITION_INTEGER_REF{Name: "DEFINITION"},
		}).(*Type)
	case ATTRIBUTE_VALUE_REAL:
		return any(&ATTRIBUTE_VALUE_REAL{
			// Initialisation of associations
			// field is initialized with an instance of A_ATTRIBUTE_DEFINITION_REAL_REF with the name of the field
			DEFINITION: &A_ATTRIBUTE_DEFINITION_REAL_REF{Name: "DEFINITION"},
		}).(*Type)
	case ATTRIBUTE_VALUE_STRING:
		return any(&ATTRIBUTE_VALUE_STRING{
			// Initialisation of associations
			// field is initialized with an instance of A_ATTRIBUTE_DEFINITION_STRING_REF with the name of the field
			DEFINITION: &A_ATTRIBUTE_DEFINITION_STRING_REF{Name: "DEFINITION"},
		}).(*Type)
	case ATTRIBUTE_VALUE_XHTML:
		return any(&ATTRIBUTE_VALUE_XHTML{
			// Initialisation of associations
			// field is initialized with an instance of XHTML_CONTENT with the name of the field
			THE_VALUE: &XHTML_CONTENT{Name: "THE_VALUE"},
			// field is initialized with an instance of XHTML_CONTENT with the name of the field
			THE_ORIGINAL_VALUE: &XHTML_CONTENT{Name: "THE_ORIGINAL_VALUE"},
			// field is initialized with an instance of A_ATTRIBUTE_DEFINITION_XHTML_REF with the name of the field
			DEFINITION: &A_ATTRIBUTE_DEFINITION_XHTML_REF{Name: "DEFINITION"},
		}).(*Type)
	case A_ALTERNATIVE_ID:
		return any(&A_ALTERNATIVE_ID{
			// Initialisation of associations
			// field is initialized with an instance of ALTERNATIVE_ID with the name of the field
			ALTERNATIVE_ID: &ALTERNATIVE_ID{Name: "ALTERNATIVE_ID"},
		}).(*Type)
	case A_ATTRIBUTE_DEFINITION_BOOLEAN_REF:
		return any(&A_ATTRIBUTE_DEFINITION_BOOLEAN_REF{
			// Initialisation of associations
		}).(*Type)
	case A_ATTRIBUTE_DEFINITION_DATE_REF:
		return any(&A_ATTRIBUTE_DEFINITION_DATE_REF{
			// Initialisation of associations
		}).(*Type)
	case A_ATTRIBUTE_DEFINITION_ENUMERATION_REF:
		return any(&A_ATTRIBUTE_DEFINITION_ENUMERATION_REF{
			// Initialisation of associations
		}).(*Type)
	case A_ATTRIBUTE_DEFINITION_INTEGER_REF:
		return any(&A_ATTRIBUTE_DEFINITION_INTEGER_REF{
			// Initialisation of associations
		}).(*Type)
	case A_ATTRIBUTE_DEFINITION_REAL_REF:
		return any(&A_ATTRIBUTE_DEFINITION_REAL_REF{
			// Initialisation of associations
		}).(*Type)
	case A_ATTRIBUTE_DEFINITION_STRING_REF:
		return any(&A_ATTRIBUTE_DEFINITION_STRING_REF{
			// Initialisation of associations
		}).(*Type)
	case A_ATTRIBUTE_DEFINITION_XHTML_REF:
		return any(&A_ATTRIBUTE_DEFINITION_XHTML_REF{
			// Initialisation of associations
		}).(*Type)
	case A_ATTRIBUTE_VALUE_BOOLEAN:
		return any(&A_ATTRIBUTE_VALUE_BOOLEAN{
			// Initialisation of associations
			// field is initialized with an instance of ATTRIBUTE_VALUE_BOOLEAN with the name of the field
			ATTRIBUTE_VALUE_BOOLEAN: &ATTRIBUTE_VALUE_BOOLEAN{Name: "ATTRIBUTE_VALUE_BOOLEAN"},
		}).(*Type)
	case A_ATTRIBUTE_VALUE_DATE:
		return any(&A_ATTRIBUTE_VALUE_DATE{
			// Initialisation of associations
			// field is initialized with an instance of ATTRIBUTE_VALUE_DATE with the name of the field
			ATTRIBUTE_VALUE_DATE: &ATTRIBUTE_VALUE_DATE{Name: "ATTRIBUTE_VALUE_DATE"},
		}).(*Type)
	case A_ATTRIBUTE_VALUE_ENUMERATION:
		return any(&A_ATTRIBUTE_VALUE_ENUMERATION{
			// Initialisation of associations
			// field is initialized with an instance of ATTRIBUTE_VALUE_ENUMERATION with the name of the field
			ATTRIBUTE_VALUE_ENUMERATION: &ATTRIBUTE_VALUE_ENUMERATION{Name: "ATTRIBUTE_VALUE_ENUMERATION"},
		}).(*Type)
	case A_ATTRIBUTE_VALUE_INTEGER:
		return any(&A_ATTRIBUTE_VALUE_INTEGER{
			// Initialisation of associations
			// field is initialized with an instance of ATTRIBUTE_VALUE_INTEGER with the name of the field
			ATTRIBUTE_VALUE_INTEGER: &ATTRIBUTE_VALUE_INTEGER{Name: "ATTRIBUTE_VALUE_INTEGER"},
		}).(*Type)
	case A_ATTRIBUTE_VALUE_REAL:
		return any(&A_ATTRIBUTE_VALUE_REAL{
			// Initialisation of associations
			// field is initialized with an instance of ATTRIBUTE_VALUE_REAL with the name of the field
			ATTRIBUTE_VALUE_REAL: &ATTRIBUTE_VALUE_REAL{Name: "ATTRIBUTE_VALUE_REAL"},
		}).(*Type)
	case A_ATTRIBUTE_VALUE_STRING:
		return any(&A_ATTRIBUTE_VALUE_STRING{
			// Initialisation of associations
			// field is initialized with an instance of ATTRIBUTE_VALUE_STRING with the name of the field
			ATTRIBUTE_VALUE_STRING: &ATTRIBUTE_VALUE_STRING{Name: "ATTRIBUTE_VALUE_STRING"},
		}).(*Type)
	case A_ATTRIBUTE_VALUE_XHTML:
		return any(&A_ATTRIBUTE_VALUE_XHTML{
			// Initialisation of associations
			// field is initialized with an instance of ATTRIBUTE_VALUE_BOOLEAN with the name of the field
			ATTRIBUTE_VALUE_BOOLEAN: &ATTRIBUTE_VALUE_BOOLEAN{Name: "ATTRIBUTE_VALUE_BOOLEAN"},
			// field is initialized with an instance of ATTRIBUTE_VALUE_DATE with the name of the field
			ATTRIBUTE_VALUE_DATE: &ATTRIBUTE_VALUE_DATE{Name: "ATTRIBUTE_VALUE_DATE"},
			// field is initialized with an instance of ATTRIBUTE_VALUE_ENUMERATION with the name of the field
			ATTRIBUTE_VALUE_ENUMERATION: &ATTRIBUTE_VALUE_ENUMERATION{Name: "ATTRIBUTE_VALUE_ENUMERATION"},
			// field is initialized with an instance of ATTRIBUTE_VALUE_INTEGER with the name of the field
			ATTRIBUTE_VALUE_INTEGER: &ATTRIBUTE_VALUE_INTEGER{Name: "ATTRIBUTE_VALUE_INTEGER"},
			// field is initialized with an instance of ATTRIBUTE_VALUE_REAL with the name of the field
			ATTRIBUTE_VALUE_REAL: &ATTRIBUTE_VALUE_REAL{Name: "ATTRIBUTE_VALUE_REAL"},
			// field is initialized with an instance of ATTRIBUTE_VALUE_STRING with the name of the field
			ATTRIBUTE_VALUE_STRING: &ATTRIBUTE_VALUE_STRING{Name: "ATTRIBUTE_VALUE_STRING"},
			// field is initialized with an instance of ATTRIBUTE_VALUE_XHTML with the name of the field
			ATTRIBUTE_VALUE_XHTML: &ATTRIBUTE_VALUE_XHTML{Name: "ATTRIBUTE_VALUE_XHTML"},
		}).(*Type)
	case A_ATTRIBUTE_VALUE_XHTML_1:
		return any(&A_ATTRIBUTE_VALUE_XHTML_1{
			// Initialisation of associations
			// field is initialized with an instance of ATTRIBUTE_VALUE_XHTML with the name of the field
			ATTRIBUTE_VALUE_XHTML: &ATTRIBUTE_VALUE_XHTML{Name: "ATTRIBUTE_VALUE_XHTML"},
		}).(*Type)
	case A_CHILDREN:
		return any(&A_CHILDREN{
			// Initialisation of associations
			// field is initialized with an instance of SPEC_HIERARCHY with the name of the field
			SPEC_HIERARCHY: &SPEC_HIERARCHY{Name: "SPEC_HIERARCHY"},
		}).(*Type)
	case A_CORE_CONTENT:
		return any(&A_CORE_CONTENT{
			// Initialisation of associations
			// field is initialized with an instance of REQ_IF_CONTENT with the name of the field
			REQ_IF_CONTENT: &REQ_IF_CONTENT{Name: "REQ_IF_CONTENT"},
		}).(*Type)
	case A_DATATYPES:
		return any(&A_DATATYPES{
			// Initialisation of associations
			// field is initialized with an instance of DATATYPE_DEFINITION_BOOLEAN with the name of the field
			DATATYPE_DEFINITION_BOOLEAN: &DATATYPE_DEFINITION_BOOLEAN{Name: "DATATYPE_DEFINITION_BOOLEAN"},
			// field is initialized with an instance of DATATYPE_DEFINITION_DATE with the name of the field
			DATATYPE_DEFINITION_DATE: &DATATYPE_DEFINITION_DATE{Name: "DATATYPE_DEFINITION_DATE"},
			// field is initialized with an instance of DATATYPE_DEFINITION_ENUMERATION with the name of the field
			DATATYPE_DEFINITION_ENUMERATION: &DATATYPE_DEFINITION_ENUMERATION{Name: "DATATYPE_DEFINITION_ENUMERATION"},
			// field is initialized with an instance of DATATYPE_DEFINITION_INTEGER with the name of the field
			DATATYPE_DEFINITION_INTEGER: &DATATYPE_DEFINITION_INTEGER{Name: "DATATYPE_DEFINITION_INTEGER"},
			// field is initialized with an instance of DATATYPE_DEFINITION_REAL with the name of the field
			DATATYPE_DEFINITION_REAL: &DATATYPE_DEFINITION_REAL{Name: "DATATYPE_DEFINITION_REAL"},
			// field is initialized with an instance of DATATYPE_DEFINITION_STRING with the name of the field
			DATATYPE_DEFINITION_STRING: &DATATYPE_DEFINITION_STRING{Name: "DATATYPE_DEFINITION_STRING"},
			// field is initialized with an instance of DATATYPE_DEFINITION_XHTML with the name of the field
			DATATYPE_DEFINITION_XHTML: &DATATYPE_DEFINITION_XHTML{Name: "DATATYPE_DEFINITION_XHTML"},
		}).(*Type)
	case A_DATATYPE_DEFINITION_BOOLEAN_REF:
		return any(&A_DATATYPE_DEFINITION_BOOLEAN_REF{
			// Initialisation of associations
		}).(*Type)
	case A_DATATYPE_DEFINITION_DATE_REF:
		return any(&A_DATATYPE_DEFINITION_DATE_REF{
			// Initialisation of associations
		}).(*Type)
	case A_DATATYPE_DEFINITION_ENUMERATION_REF:
		return any(&A_DATATYPE_DEFINITION_ENUMERATION_REF{
			// Initialisation of associations
		}).(*Type)
	case A_DATATYPE_DEFINITION_INTEGER_REF:
		return any(&A_DATATYPE_DEFINITION_INTEGER_REF{
			// Initialisation of associations
		}).(*Type)
	case A_DATATYPE_DEFINITION_REAL_REF:
		return any(&A_DATATYPE_DEFINITION_REAL_REF{
			// Initialisation of associations
		}).(*Type)
	case A_DATATYPE_DEFINITION_STRING_REF:
		return any(&A_DATATYPE_DEFINITION_STRING_REF{
			// Initialisation of associations
		}).(*Type)
	case A_DATATYPE_DEFINITION_XHTML_REF:
		return any(&A_DATATYPE_DEFINITION_XHTML_REF{
			// Initialisation of associations
		}).(*Type)
	case A_EDITABLE_ATTS:
		return any(&A_EDITABLE_ATTS{
			// Initialisation of associations
		}).(*Type)
	case A_ENUM_VALUE_REF:
		return any(&A_ENUM_VALUE_REF{
			// Initialisation of associations
		}).(*Type)
	case A_OBJECT:
		return any(&A_OBJECT{
			// Initialisation of associations
		}).(*Type)
	case A_PROPERTIES:
		return any(&A_PROPERTIES{
			// Initialisation of associations
			// field is initialized with an instance of EMBEDDED_VALUE with the name of the field
			EMBEDDED_VALUE: &EMBEDDED_VALUE{Name: "EMBEDDED_VALUE"},
		}).(*Type)
	case A_RELATION_GROUP_TYPE_REF:
		return any(&A_RELATION_GROUP_TYPE_REF{
			// Initialisation of associations
		}).(*Type)
	case A_SOURCE_1:
		return any(&A_SOURCE_1{
			// Initialisation of associations
		}).(*Type)
	case A_SPECIFICATIONS:
		return any(&A_SPECIFICATIONS{
			// Initialisation of associations
			// field is initialized with an instance of SPECIFICATION with the name of the field
			SPECIFICATION: &SPECIFICATION{Name: "SPECIFICATION"},
		}).(*Type)
	case A_SPECIFICATION_TYPE_REF:
		return any(&A_SPECIFICATION_TYPE_REF{
			// Initialisation of associations
		}).(*Type)
	case A_SPECIFIED_VALUES:
		return any(&A_SPECIFIED_VALUES{
			// Initialisation of associations
			// field is initialized with an instance of ENUM_VALUE with the name of the field
			ENUM_VALUE: &ENUM_VALUE{Name: "ENUM_VALUE"},
		}).(*Type)
	case A_SPEC_ATTRIBUTES:
		return any(&A_SPEC_ATTRIBUTES{
			// Initialisation of associations
			// field is initialized with an instance of ATTRIBUTE_DEFINITION_BOOLEAN with the name of the field
			ATTRIBUTE_DEFINITION_BOOLEAN: &ATTRIBUTE_DEFINITION_BOOLEAN{Name: "ATTRIBUTE_DEFINITION_BOOLEAN"},
			// field is initialized with an instance of ATTRIBUTE_DEFINITION_DATE with the name of the field
			ATTRIBUTE_DEFINITION_DATE: &ATTRIBUTE_DEFINITION_DATE{Name: "ATTRIBUTE_DEFINITION_DATE"},
			// field is initialized with an instance of ATTRIBUTE_DEFINITION_ENUMERATION with the name of the field
			ATTRIBUTE_DEFINITION_ENUMERATION: &ATTRIBUTE_DEFINITION_ENUMERATION{Name: "ATTRIBUTE_DEFINITION_ENUMERATION"},
			// field is initialized with an instance of ATTRIBUTE_DEFINITION_INTEGER with the name of the field
			ATTRIBUTE_DEFINITION_INTEGER: &ATTRIBUTE_DEFINITION_INTEGER{Name: "ATTRIBUTE_DEFINITION_INTEGER"},
			// field is initialized with an instance of ATTRIBUTE_DEFINITION_REAL with the name of the field
			ATTRIBUTE_DEFINITION_REAL: &ATTRIBUTE_DEFINITION_REAL{Name: "ATTRIBUTE_DEFINITION_REAL"},
			// field is initialized with an instance of ATTRIBUTE_DEFINITION_STRING with the name of the field
			ATTRIBUTE_DEFINITION_STRING: &ATTRIBUTE_DEFINITION_STRING{Name: "ATTRIBUTE_DEFINITION_STRING"},
			// field is initialized with an instance of ATTRIBUTE_DEFINITION_XHTML with the name of the field
			ATTRIBUTE_DEFINITION_XHTML: &ATTRIBUTE_DEFINITION_XHTML{Name: "ATTRIBUTE_DEFINITION_XHTML"},
		}).(*Type)
	case A_SPEC_OBJECTS:
		return any(&A_SPEC_OBJECTS{
			// Initialisation of associations
			// field is initialized with an instance of SPEC_OBJECT with the name of the field
			SPEC_OBJECT: &SPEC_OBJECT{Name: "SPEC_OBJECT"},
		}).(*Type)
	case A_SPEC_OBJECT_TYPE_REF:
		return any(&A_SPEC_OBJECT_TYPE_REF{
			// Initialisation of associations
		}).(*Type)
	case A_SPEC_RELATIONS:
		return any(&A_SPEC_RELATIONS{
			// Initialisation of associations
			// field is initialized with an instance of SPEC_RELATION with the name of the field
			SPEC_RELATION: &SPEC_RELATION{Name: "SPEC_RELATION"},
		}).(*Type)
	case A_SPEC_RELATION_GROUPS:
		return any(&A_SPEC_RELATION_GROUPS{
			// Initialisation of associations
			// field is initialized with an instance of RELATION_GROUP with the name of the field
			RELATION_GROUP: &RELATION_GROUP{Name: "RELATION_GROUP"},
		}).(*Type)
	case A_SPEC_RELATION_REF:
		return any(&A_SPEC_RELATION_REF{
			// Initialisation of associations
		}).(*Type)
	case A_SPEC_RELATION_TYPE_REF:
		return any(&A_SPEC_RELATION_TYPE_REF{
			// Initialisation of associations
		}).(*Type)
	case A_SPEC_TYPES:
		return any(&A_SPEC_TYPES{
			// Initialisation of associations
			// field is initialized with an instance of RELATION_GROUP_TYPE with the name of the field
			RELATION_GROUP_TYPE: &RELATION_GROUP_TYPE{Name: "RELATION_GROUP_TYPE"},
			// field is initialized with an instance of SPEC_OBJECT_TYPE with the name of the field
			SPEC_OBJECT_TYPE: &SPEC_OBJECT_TYPE{Name: "SPEC_OBJECT_TYPE"},
			// field is initialized with an instance of SPEC_RELATION_TYPE with the name of the field
			SPEC_RELATION_TYPE: &SPEC_RELATION_TYPE{Name: "SPEC_RELATION_TYPE"},
			// field is initialized with an instance of SPECIFICATION_TYPE with the name of the field
			SPECIFICATION_TYPE: &SPECIFICATION_TYPE{Name: "SPECIFICATION_TYPE"},
		}).(*Type)
	case A_TARGET_SPECIFICATION_1:
		return any(&A_TARGET_SPECIFICATION_1{
			// Initialisation of associations
		}).(*Type)
	case A_THE_HEADER:
		return any(&A_THE_HEADER{
			// Initialisation of associations
			// field is initialized with an instance of REQ_IF_HEADER with the name of the field
			REQ_IF_HEADER: &REQ_IF_HEADER{Name: "REQ_IF_HEADER"},
		}).(*Type)
	case A_TOOL_EXTENSIONS:
		return any(&A_TOOL_EXTENSIONS{
			// Initialisation of associations
			// field is initialized with an instance of REQ_IF_TOOL_EXTENSION with the name of the field
			REQ_IF_TOOL_EXTENSION: &REQ_IF_TOOL_EXTENSION{Name: "REQ_IF_TOOL_EXTENSION"},
		}).(*Type)
	case DATATYPE_DEFINITION_BOOLEAN:
		return any(&DATATYPE_DEFINITION_BOOLEAN{
			// Initialisation of associations
			// field is initialized with an instance of A_ALTERNATIVE_ID with the name of the field
			ALTERNATIVE_ID: &A_ALTERNATIVE_ID{Name: "ALTERNATIVE_ID"},
		}).(*Type)
	case DATATYPE_DEFINITION_DATE:
		return any(&DATATYPE_DEFINITION_DATE{
			// Initialisation of associations
			// field is initialized with an instance of A_ALTERNATIVE_ID with the name of the field
			ALTERNATIVE_ID: &A_ALTERNATIVE_ID{Name: "ALTERNATIVE_ID"},
		}).(*Type)
	case DATATYPE_DEFINITION_ENUMERATION:
		return any(&DATATYPE_DEFINITION_ENUMERATION{
			// Initialisation of associations
			// field is initialized with an instance of A_ALTERNATIVE_ID with the name of the field
			ALTERNATIVE_ID: &A_ALTERNATIVE_ID{Name: "ALTERNATIVE_ID"},
			// field is initialized with an instance of A_SPECIFIED_VALUES with the name of the field
			SPECIFIED_VALUES: &A_SPECIFIED_VALUES{Name: "SPECIFIED_VALUES"},
		}).(*Type)
	case DATATYPE_DEFINITION_INTEGER:
		return any(&DATATYPE_DEFINITION_INTEGER{
			// Initialisation of associations
			// field is initialized with an instance of A_ALTERNATIVE_ID with the name of the field
			ALTERNATIVE_ID: &A_ALTERNATIVE_ID{Name: "ALTERNATIVE_ID"},
		}).(*Type)
	case DATATYPE_DEFINITION_REAL:
		return any(&DATATYPE_DEFINITION_REAL{
			// Initialisation of associations
			// field is initialized with an instance of A_ALTERNATIVE_ID with the name of the field
			ALTERNATIVE_ID: &A_ALTERNATIVE_ID{Name: "ALTERNATIVE_ID"},
		}).(*Type)
	case DATATYPE_DEFINITION_STRING:
		return any(&DATATYPE_DEFINITION_STRING{
			// Initialisation of associations
			// field is initialized with an instance of A_ALTERNATIVE_ID with the name of the field
			ALTERNATIVE_ID: &A_ALTERNATIVE_ID{Name: "ALTERNATIVE_ID"},
		}).(*Type)
	case DATATYPE_DEFINITION_XHTML:
		return any(&DATATYPE_DEFINITION_XHTML{
			// Initialisation of associations
			// field is initialized with an instance of A_ALTERNATIVE_ID with the name of the field
			ALTERNATIVE_ID: &A_ALTERNATIVE_ID{Name: "ALTERNATIVE_ID"},
		}).(*Type)
	case EMBEDDED_VALUE:
		return any(&EMBEDDED_VALUE{
			// Initialisation of associations
		}).(*Type)
	case ENUM_VALUE:
		return any(&ENUM_VALUE{
			// Initialisation of associations
			// field is initialized with an instance of A_ALTERNATIVE_ID with the name of the field
			ALTERNATIVE_ID: &A_ALTERNATIVE_ID{Name: "ALTERNATIVE_ID"},
			// field is initialized with an instance of A_PROPERTIES with the name of the field
			PROPERTIES: &A_PROPERTIES{Name: "PROPERTIES"},
		}).(*Type)
	case RELATION_GROUP:
		return any(&RELATION_GROUP{
			// Initialisation of associations
			// field is initialized with an instance of A_ALTERNATIVE_ID with the name of the field
			ALTERNATIVE_ID: &A_ALTERNATIVE_ID{Name: "ALTERNATIVE_ID"},
			// field is initialized with an instance of A_TARGET_SPECIFICATION_1 with the name of the field
			SOURCE_SPECIFICATION: &A_TARGET_SPECIFICATION_1{Name: "SOURCE_SPECIFICATION"},
			// field is initialized with an instance of A_SPEC_RELATION_REF with the name of the field
			SPEC_RELATIONS: &A_SPEC_RELATION_REF{Name: "SPEC_RELATIONS"},
			// field is initialized with an instance of A_TARGET_SPECIFICATION_1 with the name of the field
			TARGET_SPECIFICATION: &A_TARGET_SPECIFICATION_1{Name: "TARGET_SPECIFICATION"},
			// field is initialized with an instance of A_RELATION_GROUP_TYPE_REF with the name of the field
			TYPE: &A_RELATION_GROUP_TYPE_REF{Name: "TYPE"},
		}).(*Type)
	case RELATION_GROUP_TYPE:
		return any(&RELATION_GROUP_TYPE{
			// Initialisation of associations
			// field is initialized with an instance of A_ALTERNATIVE_ID with the name of the field
			ALTERNATIVE_ID: &A_ALTERNATIVE_ID{Name: "ALTERNATIVE_ID"},
			// field is initialized with an instance of A_SPEC_ATTRIBUTES with the name of the field
			SPEC_ATTRIBUTES: &A_SPEC_ATTRIBUTES{Name: "SPEC_ATTRIBUTES"},
		}).(*Type)
	case REQ_IF:
		return any(&REQ_IF{
			// Initialisation of associations
			// field is initialized with an instance of A_THE_HEADER with the name of the field
			THE_HEADER: &A_THE_HEADER{Name: "THE_HEADER"},
			// field is initialized with an instance of A_CORE_CONTENT with the name of the field
			CORE_CONTENT: &A_CORE_CONTENT{Name: "CORE_CONTENT"},
			// field is initialized with an instance of A_TOOL_EXTENSIONS with the name of the field
			TOOL_EXTENSIONS: &A_TOOL_EXTENSIONS{Name: "TOOL_EXTENSIONS"},
		}).(*Type)
	case REQ_IF_CONTENT:
		return any(&REQ_IF_CONTENT{
			// Initialisation of associations
			// field is initialized with an instance of A_DATATYPES with the name of the field
			DATATYPES: &A_DATATYPES{Name: "DATATYPES"},
			// field is initialized with an instance of A_SPEC_TYPES with the name of the field
			SPEC_TYPES: &A_SPEC_TYPES{Name: "SPEC_TYPES"},
			// field is initialized with an instance of A_SPEC_OBJECTS with the name of the field
			SPEC_OBJECTS: &A_SPEC_OBJECTS{Name: "SPEC_OBJECTS"},
			// field is initialized with an instance of A_SPEC_RELATIONS with the name of the field
			SPEC_RELATIONS: &A_SPEC_RELATIONS{Name: "SPEC_RELATIONS"},
			// field is initialized with an instance of A_SPECIFICATIONS with the name of the field
			SPECIFICATIONS: &A_SPECIFICATIONS{Name: "SPECIFICATIONS"},
			// field is initialized with an instance of A_SPEC_RELATION_GROUPS with the name of the field
			SPEC_RELATION_GROUPS: &A_SPEC_RELATION_GROUPS{Name: "SPEC_RELATION_GROUPS"},
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
			// field is initialized with an instance of A_ALTERNATIVE_ID with the name of the field
			ALTERNATIVE_ID: &A_ALTERNATIVE_ID{Name: "ALTERNATIVE_ID"},
			// field is initialized with an instance of A_CHILDREN with the name of the field
			CHILDREN: &A_CHILDREN{Name: "CHILDREN"},
			// field is initialized with an instance of A_ATTRIBUTE_VALUE_XHTML with the name of the field
			VALUES: &A_ATTRIBUTE_VALUE_XHTML{Name: "VALUES"},
			// field is initialized with an instance of A_SPECIFICATION_TYPE_REF with the name of the field
			TYPE: &A_SPECIFICATION_TYPE_REF{Name: "TYPE"},
		}).(*Type)
	case SPECIFICATION_TYPE:
		return any(&SPECIFICATION_TYPE{
			// Initialisation of associations
			// field is initialized with an instance of A_ALTERNATIVE_ID with the name of the field
			ALTERNATIVE_ID: &A_ALTERNATIVE_ID{Name: "ALTERNATIVE_ID"},
			// field is initialized with an instance of A_SPEC_ATTRIBUTES with the name of the field
			SPEC_ATTRIBUTES: &A_SPEC_ATTRIBUTES{Name: "SPEC_ATTRIBUTES"},
		}).(*Type)
	case SPEC_HIERARCHY:
		return any(&SPEC_HIERARCHY{
			// Initialisation of associations
			// field is initialized with an instance of A_ALTERNATIVE_ID with the name of the field
			ALTERNATIVE_ID: &A_ALTERNATIVE_ID{Name: "ALTERNATIVE_ID"},
			// field is initialized with an instance of A_CHILDREN with the name of the field
			CHILDREN: &A_CHILDREN{Name: "CHILDREN"},
			// field is initialized with an instance of A_EDITABLE_ATTS with the name of the field
			EDITABLE_ATTS: &A_EDITABLE_ATTS{Name: "EDITABLE_ATTS"},
			// field is initialized with an instance of A_OBJECT with the name of the field
			OBJECT: &A_OBJECT{Name: "OBJECT"},
		}).(*Type)
	case SPEC_OBJECT:
		return any(&SPEC_OBJECT{
			// Initialisation of associations
			// field is initialized with an instance of A_ALTERNATIVE_ID with the name of the field
			ALTERNATIVE_ID: &A_ALTERNATIVE_ID{Name: "ALTERNATIVE_ID"},
			// field is initialized with an instance of A_ATTRIBUTE_VALUE_XHTML with the name of the field
			VALUES: &A_ATTRIBUTE_VALUE_XHTML{Name: "VALUES"},
			// field is initialized with an instance of A_SPEC_OBJECT_TYPE_REF with the name of the field
			TYPE: &A_SPEC_OBJECT_TYPE_REF{Name: "TYPE"},
		}).(*Type)
	case SPEC_OBJECT_TYPE:
		return any(&SPEC_OBJECT_TYPE{
			// Initialisation of associations
			// field is initialized with an instance of A_ALTERNATIVE_ID with the name of the field
			ALTERNATIVE_ID: &A_ALTERNATIVE_ID{Name: "ALTERNATIVE_ID"},
			// field is initialized with an instance of A_SPEC_ATTRIBUTES with the name of the field
			SPEC_ATTRIBUTES: &A_SPEC_ATTRIBUTES{Name: "SPEC_ATTRIBUTES"},
		}).(*Type)
	case SPEC_RELATION:
		return any(&SPEC_RELATION{
			// Initialisation of associations
			// field is initialized with an instance of A_ALTERNATIVE_ID with the name of the field
			ALTERNATIVE_ID: &A_ALTERNATIVE_ID{Name: "ALTERNATIVE_ID"},
			// field is initialized with an instance of A_ATTRIBUTE_VALUE_XHTML with the name of the field
			VALUES: &A_ATTRIBUTE_VALUE_XHTML{Name: "VALUES"},
			// field is initialized with an instance of A_SOURCE_1 with the name of the field
			SOURCE: &A_SOURCE_1{Name: "SOURCE"},
			// field is initialized with an instance of A_SOURCE_1 with the name of the field
			TARGET: &A_SOURCE_1{Name: "TARGET"},
			// field is initialized with an instance of A_SPEC_RELATION_TYPE_REF with the name of the field
			TYPE: &A_SPEC_RELATION_TYPE_REF{Name: "TYPE"},
		}).(*Type)
	case SPEC_RELATION_TYPE:
		return any(&SPEC_RELATION_TYPE{
			// Initialisation of associations
			// field is initialized with an instance of A_ALTERNATIVE_ID with the name of the field
			ALTERNATIVE_ID: &A_ALTERNATIVE_ID{Name: "ALTERNATIVE_ID"},
			// field is initialized with an instance of A_SPEC_ATTRIBUTES with the name of the field
			SPEC_ATTRIBUTES: &A_SPEC_ATTRIBUTES{Name: "SPEC_ATTRIBUTES"},
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
		case "ALTERNATIVE_ID":
			res := make(map[*A_ALTERNATIVE_ID][]*ATTRIBUTE_DEFINITION_BOOLEAN)
			for attribute_definition_boolean := range stage.ATTRIBUTE_DEFINITION_BOOLEANs {
				if attribute_definition_boolean.ALTERNATIVE_ID != nil {
					a_alternative_id_ := attribute_definition_boolean.ALTERNATIVE_ID
					var attribute_definition_booleans []*ATTRIBUTE_DEFINITION_BOOLEAN
					_, ok := res[a_alternative_id_]
					if ok {
						attribute_definition_booleans = res[a_alternative_id_]
					} else {
						attribute_definition_booleans = make([]*ATTRIBUTE_DEFINITION_BOOLEAN, 0)
					}
					attribute_definition_booleans = append(attribute_definition_booleans, attribute_definition_boolean)
					res[a_alternative_id_] = attribute_definition_booleans
				}
			}
			return any(res).(map[*End][]*Start)
		case "DEFAULT_VALUE":
			res := make(map[*A_ATTRIBUTE_VALUE_BOOLEAN][]*ATTRIBUTE_DEFINITION_BOOLEAN)
			for attribute_definition_boolean := range stage.ATTRIBUTE_DEFINITION_BOOLEANs {
				if attribute_definition_boolean.DEFAULT_VALUE != nil {
					a_attribute_value_boolean_ := attribute_definition_boolean.DEFAULT_VALUE
					var attribute_definition_booleans []*ATTRIBUTE_DEFINITION_BOOLEAN
					_, ok := res[a_attribute_value_boolean_]
					if ok {
						attribute_definition_booleans = res[a_attribute_value_boolean_]
					} else {
						attribute_definition_booleans = make([]*ATTRIBUTE_DEFINITION_BOOLEAN, 0)
					}
					attribute_definition_booleans = append(attribute_definition_booleans, attribute_definition_boolean)
					res[a_attribute_value_boolean_] = attribute_definition_booleans
				}
			}
			return any(res).(map[*End][]*Start)
		case "TYPE":
			res := make(map[*A_DATATYPE_DEFINITION_BOOLEAN_REF][]*ATTRIBUTE_DEFINITION_BOOLEAN)
			for attribute_definition_boolean := range stage.ATTRIBUTE_DEFINITION_BOOLEANs {
				if attribute_definition_boolean.TYPE != nil {
					a_datatype_definition_boolean_ref_ := attribute_definition_boolean.TYPE
					var attribute_definition_booleans []*ATTRIBUTE_DEFINITION_BOOLEAN
					_, ok := res[a_datatype_definition_boolean_ref_]
					if ok {
						attribute_definition_booleans = res[a_datatype_definition_boolean_ref_]
					} else {
						attribute_definition_booleans = make([]*ATTRIBUTE_DEFINITION_BOOLEAN, 0)
					}
					attribute_definition_booleans = append(attribute_definition_booleans, attribute_definition_boolean)
					res[a_datatype_definition_boolean_ref_] = attribute_definition_booleans
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ATTRIBUTE_DEFINITION_DATE
	case ATTRIBUTE_DEFINITION_DATE:
		switch fieldname {
		// insertion point for per direct association field
		case "ALTERNATIVE_ID":
			res := make(map[*A_ALTERNATIVE_ID][]*ATTRIBUTE_DEFINITION_DATE)
			for attribute_definition_date := range stage.ATTRIBUTE_DEFINITION_DATEs {
				if attribute_definition_date.ALTERNATIVE_ID != nil {
					a_alternative_id_ := attribute_definition_date.ALTERNATIVE_ID
					var attribute_definition_dates []*ATTRIBUTE_DEFINITION_DATE
					_, ok := res[a_alternative_id_]
					if ok {
						attribute_definition_dates = res[a_alternative_id_]
					} else {
						attribute_definition_dates = make([]*ATTRIBUTE_DEFINITION_DATE, 0)
					}
					attribute_definition_dates = append(attribute_definition_dates, attribute_definition_date)
					res[a_alternative_id_] = attribute_definition_dates
				}
			}
			return any(res).(map[*End][]*Start)
		case "DEFAULT_VALUE":
			res := make(map[*A_ATTRIBUTE_VALUE_DATE][]*ATTRIBUTE_DEFINITION_DATE)
			for attribute_definition_date := range stage.ATTRIBUTE_DEFINITION_DATEs {
				if attribute_definition_date.DEFAULT_VALUE != nil {
					a_attribute_value_date_ := attribute_definition_date.DEFAULT_VALUE
					var attribute_definition_dates []*ATTRIBUTE_DEFINITION_DATE
					_, ok := res[a_attribute_value_date_]
					if ok {
						attribute_definition_dates = res[a_attribute_value_date_]
					} else {
						attribute_definition_dates = make([]*ATTRIBUTE_DEFINITION_DATE, 0)
					}
					attribute_definition_dates = append(attribute_definition_dates, attribute_definition_date)
					res[a_attribute_value_date_] = attribute_definition_dates
				}
			}
			return any(res).(map[*End][]*Start)
		case "TYPE":
			res := make(map[*A_DATATYPE_DEFINITION_DATE_REF][]*ATTRIBUTE_DEFINITION_DATE)
			for attribute_definition_date := range stage.ATTRIBUTE_DEFINITION_DATEs {
				if attribute_definition_date.TYPE != nil {
					a_datatype_definition_date_ref_ := attribute_definition_date.TYPE
					var attribute_definition_dates []*ATTRIBUTE_DEFINITION_DATE
					_, ok := res[a_datatype_definition_date_ref_]
					if ok {
						attribute_definition_dates = res[a_datatype_definition_date_ref_]
					} else {
						attribute_definition_dates = make([]*ATTRIBUTE_DEFINITION_DATE, 0)
					}
					attribute_definition_dates = append(attribute_definition_dates, attribute_definition_date)
					res[a_datatype_definition_date_ref_] = attribute_definition_dates
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ATTRIBUTE_DEFINITION_ENUMERATION
	case ATTRIBUTE_DEFINITION_ENUMERATION:
		switch fieldname {
		// insertion point for per direct association field
		case "ALTERNATIVE_ID":
			res := make(map[*A_ALTERNATIVE_ID][]*ATTRIBUTE_DEFINITION_ENUMERATION)
			for attribute_definition_enumeration := range stage.ATTRIBUTE_DEFINITION_ENUMERATIONs {
				if attribute_definition_enumeration.ALTERNATIVE_ID != nil {
					a_alternative_id_ := attribute_definition_enumeration.ALTERNATIVE_ID
					var attribute_definition_enumerations []*ATTRIBUTE_DEFINITION_ENUMERATION
					_, ok := res[a_alternative_id_]
					if ok {
						attribute_definition_enumerations = res[a_alternative_id_]
					} else {
						attribute_definition_enumerations = make([]*ATTRIBUTE_DEFINITION_ENUMERATION, 0)
					}
					attribute_definition_enumerations = append(attribute_definition_enumerations, attribute_definition_enumeration)
					res[a_alternative_id_] = attribute_definition_enumerations
				}
			}
			return any(res).(map[*End][]*Start)
		case "DEFAULT_VALUE":
			res := make(map[*A_ATTRIBUTE_VALUE_ENUMERATION][]*ATTRIBUTE_DEFINITION_ENUMERATION)
			for attribute_definition_enumeration := range stage.ATTRIBUTE_DEFINITION_ENUMERATIONs {
				if attribute_definition_enumeration.DEFAULT_VALUE != nil {
					a_attribute_value_enumeration_ := attribute_definition_enumeration.DEFAULT_VALUE
					var attribute_definition_enumerations []*ATTRIBUTE_DEFINITION_ENUMERATION
					_, ok := res[a_attribute_value_enumeration_]
					if ok {
						attribute_definition_enumerations = res[a_attribute_value_enumeration_]
					} else {
						attribute_definition_enumerations = make([]*ATTRIBUTE_DEFINITION_ENUMERATION, 0)
					}
					attribute_definition_enumerations = append(attribute_definition_enumerations, attribute_definition_enumeration)
					res[a_attribute_value_enumeration_] = attribute_definition_enumerations
				}
			}
			return any(res).(map[*End][]*Start)
		case "TYPE":
			res := make(map[*A_DATATYPE_DEFINITION_ENUMERATION_REF][]*ATTRIBUTE_DEFINITION_ENUMERATION)
			for attribute_definition_enumeration := range stage.ATTRIBUTE_DEFINITION_ENUMERATIONs {
				if attribute_definition_enumeration.TYPE != nil {
					a_datatype_definition_enumeration_ref_ := attribute_definition_enumeration.TYPE
					var attribute_definition_enumerations []*ATTRIBUTE_DEFINITION_ENUMERATION
					_, ok := res[a_datatype_definition_enumeration_ref_]
					if ok {
						attribute_definition_enumerations = res[a_datatype_definition_enumeration_ref_]
					} else {
						attribute_definition_enumerations = make([]*ATTRIBUTE_DEFINITION_ENUMERATION, 0)
					}
					attribute_definition_enumerations = append(attribute_definition_enumerations, attribute_definition_enumeration)
					res[a_datatype_definition_enumeration_ref_] = attribute_definition_enumerations
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ATTRIBUTE_DEFINITION_INTEGER
	case ATTRIBUTE_DEFINITION_INTEGER:
		switch fieldname {
		// insertion point for per direct association field
		case "ALTERNATIVE_ID":
			res := make(map[*A_ALTERNATIVE_ID][]*ATTRIBUTE_DEFINITION_INTEGER)
			for attribute_definition_integer := range stage.ATTRIBUTE_DEFINITION_INTEGERs {
				if attribute_definition_integer.ALTERNATIVE_ID != nil {
					a_alternative_id_ := attribute_definition_integer.ALTERNATIVE_ID
					var attribute_definition_integers []*ATTRIBUTE_DEFINITION_INTEGER
					_, ok := res[a_alternative_id_]
					if ok {
						attribute_definition_integers = res[a_alternative_id_]
					} else {
						attribute_definition_integers = make([]*ATTRIBUTE_DEFINITION_INTEGER, 0)
					}
					attribute_definition_integers = append(attribute_definition_integers, attribute_definition_integer)
					res[a_alternative_id_] = attribute_definition_integers
				}
			}
			return any(res).(map[*End][]*Start)
		case "DEFAULT_VALUE":
			res := make(map[*A_ATTRIBUTE_VALUE_INTEGER][]*ATTRIBUTE_DEFINITION_INTEGER)
			for attribute_definition_integer := range stage.ATTRIBUTE_DEFINITION_INTEGERs {
				if attribute_definition_integer.DEFAULT_VALUE != nil {
					a_attribute_value_integer_ := attribute_definition_integer.DEFAULT_VALUE
					var attribute_definition_integers []*ATTRIBUTE_DEFINITION_INTEGER
					_, ok := res[a_attribute_value_integer_]
					if ok {
						attribute_definition_integers = res[a_attribute_value_integer_]
					} else {
						attribute_definition_integers = make([]*ATTRIBUTE_DEFINITION_INTEGER, 0)
					}
					attribute_definition_integers = append(attribute_definition_integers, attribute_definition_integer)
					res[a_attribute_value_integer_] = attribute_definition_integers
				}
			}
			return any(res).(map[*End][]*Start)
		case "TYPE":
			res := make(map[*A_DATATYPE_DEFINITION_INTEGER_REF][]*ATTRIBUTE_DEFINITION_INTEGER)
			for attribute_definition_integer := range stage.ATTRIBUTE_DEFINITION_INTEGERs {
				if attribute_definition_integer.TYPE != nil {
					a_datatype_definition_integer_ref_ := attribute_definition_integer.TYPE
					var attribute_definition_integers []*ATTRIBUTE_DEFINITION_INTEGER
					_, ok := res[a_datatype_definition_integer_ref_]
					if ok {
						attribute_definition_integers = res[a_datatype_definition_integer_ref_]
					} else {
						attribute_definition_integers = make([]*ATTRIBUTE_DEFINITION_INTEGER, 0)
					}
					attribute_definition_integers = append(attribute_definition_integers, attribute_definition_integer)
					res[a_datatype_definition_integer_ref_] = attribute_definition_integers
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ATTRIBUTE_DEFINITION_REAL
	case ATTRIBUTE_DEFINITION_REAL:
		switch fieldname {
		// insertion point for per direct association field
		case "ALTERNATIVE_ID":
			res := make(map[*A_ALTERNATIVE_ID][]*ATTRIBUTE_DEFINITION_REAL)
			for attribute_definition_real := range stage.ATTRIBUTE_DEFINITION_REALs {
				if attribute_definition_real.ALTERNATIVE_ID != nil {
					a_alternative_id_ := attribute_definition_real.ALTERNATIVE_ID
					var attribute_definition_reals []*ATTRIBUTE_DEFINITION_REAL
					_, ok := res[a_alternative_id_]
					if ok {
						attribute_definition_reals = res[a_alternative_id_]
					} else {
						attribute_definition_reals = make([]*ATTRIBUTE_DEFINITION_REAL, 0)
					}
					attribute_definition_reals = append(attribute_definition_reals, attribute_definition_real)
					res[a_alternative_id_] = attribute_definition_reals
				}
			}
			return any(res).(map[*End][]*Start)
		case "DEFAULT_VALUE":
			res := make(map[*A_ATTRIBUTE_VALUE_REAL][]*ATTRIBUTE_DEFINITION_REAL)
			for attribute_definition_real := range stage.ATTRIBUTE_DEFINITION_REALs {
				if attribute_definition_real.DEFAULT_VALUE != nil {
					a_attribute_value_real_ := attribute_definition_real.DEFAULT_VALUE
					var attribute_definition_reals []*ATTRIBUTE_DEFINITION_REAL
					_, ok := res[a_attribute_value_real_]
					if ok {
						attribute_definition_reals = res[a_attribute_value_real_]
					} else {
						attribute_definition_reals = make([]*ATTRIBUTE_DEFINITION_REAL, 0)
					}
					attribute_definition_reals = append(attribute_definition_reals, attribute_definition_real)
					res[a_attribute_value_real_] = attribute_definition_reals
				}
			}
			return any(res).(map[*End][]*Start)
		case "TYPE":
			res := make(map[*A_DATATYPE_DEFINITION_REAL_REF][]*ATTRIBUTE_DEFINITION_REAL)
			for attribute_definition_real := range stage.ATTRIBUTE_DEFINITION_REALs {
				if attribute_definition_real.TYPE != nil {
					a_datatype_definition_real_ref_ := attribute_definition_real.TYPE
					var attribute_definition_reals []*ATTRIBUTE_DEFINITION_REAL
					_, ok := res[a_datatype_definition_real_ref_]
					if ok {
						attribute_definition_reals = res[a_datatype_definition_real_ref_]
					} else {
						attribute_definition_reals = make([]*ATTRIBUTE_DEFINITION_REAL, 0)
					}
					attribute_definition_reals = append(attribute_definition_reals, attribute_definition_real)
					res[a_datatype_definition_real_ref_] = attribute_definition_reals
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ATTRIBUTE_DEFINITION_STRING
	case ATTRIBUTE_DEFINITION_STRING:
		switch fieldname {
		// insertion point for per direct association field
		case "ALTERNATIVE_ID":
			res := make(map[*A_ALTERNATIVE_ID][]*ATTRIBUTE_DEFINITION_STRING)
			for attribute_definition_string := range stage.ATTRIBUTE_DEFINITION_STRINGs {
				if attribute_definition_string.ALTERNATIVE_ID != nil {
					a_alternative_id_ := attribute_definition_string.ALTERNATIVE_ID
					var attribute_definition_strings []*ATTRIBUTE_DEFINITION_STRING
					_, ok := res[a_alternative_id_]
					if ok {
						attribute_definition_strings = res[a_alternative_id_]
					} else {
						attribute_definition_strings = make([]*ATTRIBUTE_DEFINITION_STRING, 0)
					}
					attribute_definition_strings = append(attribute_definition_strings, attribute_definition_string)
					res[a_alternative_id_] = attribute_definition_strings
				}
			}
			return any(res).(map[*End][]*Start)
		case "DEFAULT_VALUE":
			res := make(map[*A_ATTRIBUTE_VALUE_STRING][]*ATTRIBUTE_DEFINITION_STRING)
			for attribute_definition_string := range stage.ATTRIBUTE_DEFINITION_STRINGs {
				if attribute_definition_string.DEFAULT_VALUE != nil {
					a_attribute_value_string_ := attribute_definition_string.DEFAULT_VALUE
					var attribute_definition_strings []*ATTRIBUTE_DEFINITION_STRING
					_, ok := res[a_attribute_value_string_]
					if ok {
						attribute_definition_strings = res[a_attribute_value_string_]
					} else {
						attribute_definition_strings = make([]*ATTRIBUTE_DEFINITION_STRING, 0)
					}
					attribute_definition_strings = append(attribute_definition_strings, attribute_definition_string)
					res[a_attribute_value_string_] = attribute_definition_strings
				}
			}
			return any(res).(map[*End][]*Start)
		case "TYPE":
			res := make(map[*A_DATATYPE_DEFINITION_STRING_REF][]*ATTRIBUTE_DEFINITION_STRING)
			for attribute_definition_string := range stage.ATTRIBUTE_DEFINITION_STRINGs {
				if attribute_definition_string.TYPE != nil {
					a_datatype_definition_string_ref_ := attribute_definition_string.TYPE
					var attribute_definition_strings []*ATTRIBUTE_DEFINITION_STRING
					_, ok := res[a_datatype_definition_string_ref_]
					if ok {
						attribute_definition_strings = res[a_datatype_definition_string_ref_]
					} else {
						attribute_definition_strings = make([]*ATTRIBUTE_DEFINITION_STRING, 0)
					}
					attribute_definition_strings = append(attribute_definition_strings, attribute_definition_string)
					res[a_datatype_definition_string_ref_] = attribute_definition_strings
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ATTRIBUTE_DEFINITION_XHTML
	case ATTRIBUTE_DEFINITION_XHTML:
		switch fieldname {
		// insertion point for per direct association field
		case "ALTERNATIVE_ID":
			res := make(map[*A_ALTERNATIVE_ID][]*ATTRIBUTE_DEFINITION_XHTML)
			for attribute_definition_xhtml := range stage.ATTRIBUTE_DEFINITION_XHTMLs {
				if attribute_definition_xhtml.ALTERNATIVE_ID != nil {
					a_alternative_id_ := attribute_definition_xhtml.ALTERNATIVE_ID
					var attribute_definition_xhtmls []*ATTRIBUTE_DEFINITION_XHTML
					_, ok := res[a_alternative_id_]
					if ok {
						attribute_definition_xhtmls = res[a_alternative_id_]
					} else {
						attribute_definition_xhtmls = make([]*ATTRIBUTE_DEFINITION_XHTML, 0)
					}
					attribute_definition_xhtmls = append(attribute_definition_xhtmls, attribute_definition_xhtml)
					res[a_alternative_id_] = attribute_definition_xhtmls
				}
			}
			return any(res).(map[*End][]*Start)
		case "DEFAULT_VALUE":
			res := make(map[*A_ATTRIBUTE_VALUE_XHTML_1][]*ATTRIBUTE_DEFINITION_XHTML)
			for attribute_definition_xhtml := range stage.ATTRIBUTE_DEFINITION_XHTMLs {
				if attribute_definition_xhtml.DEFAULT_VALUE != nil {
					a_attribute_value_xhtml_1_ := attribute_definition_xhtml.DEFAULT_VALUE
					var attribute_definition_xhtmls []*ATTRIBUTE_DEFINITION_XHTML
					_, ok := res[a_attribute_value_xhtml_1_]
					if ok {
						attribute_definition_xhtmls = res[a_attribute_value_xhtml_1_]
					} else {
						attribute_definition_xhtmls = make([]*ATTRIBUTE_DEFINITION_XHTML, 0)
					}
					attribute_definition_xhtmls = append(attribute_definition_xhtmls, attribute_definition_xhtml)
					res[a_attribute_value_xhtml_1_] = attribute_definition_xhtmls
				}
			}
			return any(res).(map[*End][]*Start)
		case "TYPE":
			res := make(map[*A_DATATYPE_DEFINITION_XHTML_REF][]*ATTRIBUTE_DEFINITION_XHTML)
			for attribute_definition_xhtml := range stage.ATTRIBUTE_DEFINITION_XHTMLs {
				if attribute_definition_xhtml.TYPE != nil {
					a_datatype_definition_xhtml_ref_ := attribute_definition_xhtml.TYPE
					var attribute_definition_xhtmls []*ATTRIBUTE_DEFINITION_XHTML
					_, ok := res[a_datatype_definition_xhtml_ref_]
					if ok {
						attribute_definition_xhtmls = res[a_datatype_definition_xhtml_ref_]
					} else {
						attribute_definition_xhtmls = make([]*ATTRIBUTE_DEFINITION_XHTML, 0)
					}
					attribute_definition_xhtmls = append(attribute_definition_xhtmls, attribute_definition_xhtml)
					res[a_datatype_definition_xhtml_ref_] = attribute_definition_xhtmls
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ATTRIBUTE_VALUE_BOOLEAN
	case ATTRIBUTE_VALUE_BOOLEAN:
		switch fieldname {
		// insertion point for per direct association field
		case "DEFINITION":
			res := make(map[*A_ATTRIBUTE_DEFINITION_BOOLEAN_REF][]*ATTRIBUTE_VALUE_BOOLEAN)
			for attribute_value_boolean := range stage.ATTRIBUTE_VALUE_BOOLEANs {
				if attribute_value_boolean.DEFINITION != nil {
					a_attribute_definition_boolean_ref_ := attribute_value_boolean.DEFINITION
					var attribute_value_booleans []*ATTRIBUTE_VALUE_BOOLEAN
					_, ok := res[a_attribute_definition_boolean_ref_]
					if ok {
						attribute_value_booleans = res[a_attribute_definition_boolean_ref_]
					} else {
						attribute_value_booleans = make([]*ATTRIBUTE_VALUE_BOOLEAN, 0)
					}
					attribute_value_booleans = append(attribute_value_booleans, attribute_value_boolean)
					res[a_attribute_definition_boolean_ref_] = attribute_value_booleans
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ATTRIBUTE_VALUE_DATE
	case ATTRIBUTE_VALUE_DATE:
		switch fieldname {
		// insertion point for per direct association field
		case "DEFINITION":
			res := make(map[*A_ATTRIBUTE_DEFINITION_DATE_REF][]*ATTRIBUTE_VALUE_DATE)
			for attribute_value_date := range stage.ATTRIBUTE_VALUE_DATEs {
				if attribute_value_date.DEFINITION != nil {
					a_attribute_definition_date_ref_ := attribute_value_date.DEFINITION
					var attribute_value_dates []*ATTRIBUTE_VALUE_DATE
					_, ok := res[a_attribute_definition_date_ref_]
					if ok {
						attribute_value_dates = res[a_attribute_definition_date_ref_]
					} else {
						attribute_value_dates = make([]*ATTRIBUTE_VALUE_DATE, 0)
					}
					attribute_value_dates = append(attribute_value_dates, attribute_value_date)
					res[a_attribute_definition_date_ref_] = attribute_value_dates
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ATTRIBUTE_VALUE_ENUMERATION
	case ATTRIBUTE_VALUE_ENUMERATION:
		switch fieldname {
		// insertion point for per direct association field
		case "DEFINITION":
			res := make(map[*A_ATTRIBUTE_DEFINITION_ENUMERATION_REF][]*ATTRIBUTE_VALUE_ENUMERATION)
			for attribute_value_enumeration := range stage.ATTRIBUTE_VALUE_ENUMERATIONs {
				if attribute_value_enumeration.DEFINITION != nil {
					a_attribute_definition_enumeration_ref_ := attribute_value_enumeration.DEFINITION
					var attribute_value_enumerations []*ATTRIBUTE_VALUE_ENUMERATION
					_, ok := res[a_attribute_definition_enumeration_ref_]
					if ok {
						attribute_value_enumerations = res[a_attribute_definition_enumeration_ref_]
					} else {
						attribute_value_enumerations = make([]*ATTRIBUTE_VALUE_ENUMERATION, 0)
					}
					attribute_value_enumerations = append(attribute_value_enumerations, attribute_value_enumeration)
					res[a_attribute_definition_enumeration_ref_] = attribute_value_enumerations
				}
			}
			return any(res).(map[*End][]*Start)
		case "VALUES":
			res := make(map[*A_ENUM_VALUE_REF][]*ATTRIBUTE_VALUE_ENUMERATION)
			for attribute_value_enumeration := range stage.ATTRIBUTE_VALUE_ENUMERATIONs {
				if attribute_value_enumeration.VALUES != nil {
					a_enum_value_ref_ := attribute_value_enumeration.VALUES
					var attribute_value_enumerations []*ATTRIBUTE_VALUE_ENUMERATION
					_, ok := res[a_enum_value_ref_]
					if ok {
						attribute_value_enumerations = res[a_enum_value_ref_]
					} else {
						attribute_value_enumerations = make([]*ATTRIBUTE_VALUE_ENUMERATION, 0)
					}
					attribute_value_enumerations = append(attribute_value_enumerations, attribute_value_enumeration)
					res[a_enum_value_ref_] = attribute_value_enumerations
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ATTRIBUTE_VALUE_INTEGER
	case ATTRIBUTE_VALUE_INTEGER:
		switch fieldname {
		// insertion point for per direct association field
		case "DEFINITION":
			res := make(map[*A_ATTRIBUTE_DEFINITION_INTEGER_REF][]*ATTRIBUTE_VALUE_INTEGER)
			for attribute_value_integer := range stage.ATTRIBUTE_VALUE_INTEGERs {
				if attribute_value_integer.DEFINITION != nil {
					a_attribute_definition_integer_ref_ := attribute_value_integer.DEFINITION
					var attribute_value_integers []*ATTRIBUTE_VALUE_INTEGER
					_, ok := res[a_attribute_definition_integer_ref_]
					if ok {
						attribute_value_integers = res[a_attribute_definition_integer_ref_]
					} else {
						attribute_value_integers = make([]*ATTRIBUTE_VALUE_INTEGER, 0)
					}
					attribute_value_integers = append(attribute_value_integers, attribute_value_integer)
					res[a_attribute_definition_integer_ref_] = attribute_value_integers
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ATTRIBUTE_VALUE_REAL
	case ATTRIBUTE_VALUE_REAL:
		switch fieldname {
		// insertion point for per direct association field
		case "DEFINITION":
			res := make(map[*A_ATTRIBUTE_DEFINITION_REAL_REF][]*ATTRIBUTE_VALUE_REAL)
			for attribute_value_real := range stage.ATTRIBUTE_VALUE_REALs {
				if attribute_value_real.DEFINITION != nil {
					a_attribute_definition_real_ref_ := attribute_value_real.DEFINITION
					var attribute_value_reals []*ATTRIBUTE_VALUE_REAL
					_, ok := res[a_attribute_definition_real_ref_]
					if ok {
						attribute_value_reals = res[a_attribute_definition_real_ref_]
					} else {
						attribute_value_reals = make([]*ATTRIBUTE_VALUE_REAL, 0)
					}
					attribute_value_reals = append(attribute_value_reals, attribute_value_real)
					res[a_attribute_definition_real_ref_] = attribute_value_reals
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ATTRIBUTE_VALUE_STRING
	case ATTRIBUTE_VALUE_STRING:
		switch fieldname {
		// insertion point for per direct association field
		case "DEFINITION":
			res := make(map[*A_ATTRIBUTE_DEFINITION_STRING_REF][]*ATTRIBUTE_VALUE_STRING)
			for attribute_value_string := range stage.ATTRIBUTE_VALUE_STRINGs {
				if attribute_value_string.DEFINITION != nil {
					a_attribute_definition_string_ref_ := attribute_value_string.DEFINITION
					var attribute_value_strings []*ATTRIBUTE_VALUE_STRING
					_, ok := res[a_attribute_definition_string_ref_]
					if ok {
						attribute_value_strings = res[a_attribute_definition_string_ref_]
					} else {
						attribute_value_strings = make([]*ATTRIBUTE_VALUE_STRING, 0)
					}
					attribute_value_strings = append(attribute_value_strings, attribute_value_string)
					res[a_attribute_definition_string_ref_] = attribute_value_strings
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of ATTRIBUTE_VALUE_XHTML
	case ATTRIBUTE_VALUE_XHTML:
		switch fieldname {
		// insertion point for per direct association field
		case "THE_VALUE":
			res := make(map[*XHTML_CONTENT][]*ATTRIBUTE_VALUE_XHTML)
			for attribute_value_xhtml := range stage.ATTRIBUTE_VALUE_XHTMLs {
				if attribute_value_xhtml.THE_VALUE != nil {
					xhtml_content_ := attribute_value_xhtml.THE_VALUE
					var attribute_value_xhtmls []*ATTRIBUTE_VALUE_XHTML
					_, ok := res[xhtml_content_]
					if ok {
						attribute_value_xhtmls = res[xhtml_content_]
					} else {
						attribute_value_xhtmls = make([]*ATTRIBUTE_VALUE_XHTML, 0)
					}
					attribute_value_xhtmls = append(attribute_value_xhtmls, attribute_value_xhtml)
					res[xhtml_content_] = attribute_value_xhtmls
				}
			}
			return any(res).(map[*End][]*Start)
		case "THE_ORIGINAL_VALUE":
			res := make(map[*XHTML_CONTENT][]*ATTRIBUTE_VALUE_XHTML)
			for attribute_value_xhtml := range stage.ATTRIBUTE_VALUE_XHTMLs {
				if attribute_value_xhtml.THE_ORIGINAL_VALUE != nil {
					xhtml_content_ := attribute_value_xhtml.THE_ORIGINAL_VALUE
					var attribute_value_xhtmls []*ATTRIBUTE_VALUE_XHTML
					_, ok := res[xhtml_content_]
					if ok {
						attribute_value_xhtmls = res[xhtml_content_]
					} else {
						attribute_value_xhtmls = make([]*ATTRIBUTE_VALUE_XHTML, 0)
					}
					attribute_value_xhtmls = append(attribute_value_xhtmls, attribute_value_xhtml)
					res[xhtml_content_] = attribute_value_xhtmls
				}
			}
			return any(res).(map[*End][]*Start)
		case "DEFINITION":
			res := make(map[*A_ATTRIBUTE_DEFINITION_XHTML_REF][]*ATTRIBUTE_VALUE_XHTML)
			for attribute_value_xhtml := range stage.ATTRIBUTE_VALUE_XHTMLs {
				if attribute_value_xhtml.DEFINITION != nil {
					a_attribute_definition_xhtml_ref_ := attribute_value_xhtml.DEFINITION
					var attribute_value_xhtmls []*ATTRIBUTE_VALUE_XHTML
					_, ok := res[a_attribute_definition_xhtml_ref_]
					if ok {
						attribute_value_xhtmls = res[a_attribute_definition_xhtml_ref_]
					} else {
						attribute_value_xhtmls = make([]*ATTRIBUTE_VALUE_XHTML, 0)
					}
					attribute_value_xhtmls = append(attribute_value_xhtmls, attribute_value_xhtml)
					res[a_attribute_definition_xhtml_ref_] = attribute_value_xhtmls
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of A_ALTERNATIVE_ID
	case A_ALTERNATIVE_ID:
		switch fieldname {
		// insertion point for per direct association field
		case "ALTERNATIVE_ID":
			res := make(map[*ALTERNATIVE_ID][]*A_ALTERNATIVE_ID)
			for a_alternative_id := range stage.A_ALTERNATIVE_IDs {
				if a_alternative_id.ALTERNATIVE_ID != nil {
					alternative_id_ := a_alternative_id.ALTERNATIVE_ID
					var a_alternative_ids []*A_ALTERNATIVE_ID
					_, ok := res[alternative_id_]
					if ok {
						a_alternative_ids = res[alternative_id_]
					} else {
						a_alternative_ids = make([]*A_ALTERNATIVE_ID, 0)
					}
					a_alternative_ids = append(a_alternative_ids, a_alternative_id)
					res[alternative_id_] = a_alternative_ids
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of A_ATTRIBUTE_DEFINITION_BOOLEAN_REF
	case A_ATTRIBUTE_DEFINITION_BOOLEAN_REF:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_ATTRIBUTE_DEFINITION_DATE_REF
	case A_ATTRIBUTE_DEFINITION_DATE_REF:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_ATTRIBUTE_DEFINITION_ENUMERATION_REF
	case A_ATTRIBUTE_DEFINITION_ENUMERATION_REF:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_ATTRIBUTE_DEFINITION_INTEGER_REF
	case A_ATTRIBUTE_DEFINITION_INTEGER_REF:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_ATTRIBUTE_DEFINITION_REAL_REF
	case A_ATTRIBUTE_DEFINITION_REAL_REF:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_ATTRIBUTE_DEFINITION_STRING_REF
	case A_ATTRIBUTE_DEFINITION_STRING_REF:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_ATTRIBUTE_DEFINITION_XHTML_REF
	case A_ATTRIBUTE_DEFINITION_XHTML_REF:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_ATTRIBUTE_VALUE_BOOLEAN
	case A_ATTRIBUTE_VALUE_BOOLEAN:
		switch fieldname {
		// insertion point for per direct association field
		case "ATTRIBUTE_VALUE_BOOLEAN":
			res := make(map[*ATTRIBUTE_VALUE_BOOLEAN][]*A_ATTRIBUTE_VALUE_BOOLEAN)
			for a_attribute_value_boolean := range stage.A_ATTRIBUTE_VALUE_BOOLEANs {
				if a_attribute_value_boolean.ATTRIBUTE_VALUE_BOOLEAN != nil {
					attribute_value_boolean_ := a_attribute_value_boolean.ATTRIBUTE_VALUE_BOOLEAN
					var a_attribute_value_booleans []*A_ATTRIBUTE_VALUE_BOOLEAN
					_, ok := res[attribute_value_boolean_]
					if ok {
						a_attribute_value_booleans = res[attribute_value_boolean_]
					} else {
						a_attribute_value_booleans = make([]*A_ATTRIBUTE_VALUE_BOOLEAN, 0)
					}
					a_attribute_value_booleans = append(a_attribute_value_booleans, a_attribute_value_boolean)
					res[attribute_value_boolean_] = a_attribute_value_booleans
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of A_ATTRIBUTE_VALUE_DATE
	case A_ATTRIBUTE_VALUE_DATE:
		switch fieldname {
		// insertion point for per direct association field
		case "ATTRIBUTE_VALUE_DATE":
			res := make(map[*ATTRIBUTE_VALUE_DATE][]*A_ATTRIBUTE_VALUE_DATE)
			for a_attribute_value_date := range stage.A_ATTRIBUTE_VALUE_DATEs {
				if a_attribute_value_date.ATTRIBUTE_VALUE_DATE != nil {
					attribute_value_date_ := a_attribute_value_date.ATTRIBUTE_VALUE_DATE
					var a_attribute_value_dates []*A_ATTRIBUTE_VALUE_DATE
					_, ok := res[attribute_value_date_]
					if ok {
						a_attribute_value_dates = res[attribute_value_date_]
					} else {
						a_attribute_value_dates = make([]*A_ATTRIBUTE_VALUE_DATE, 0)
					}
					a_attribute_value_dates = append(a_attribute_value_dates, a_attribute_value_date)
					res[attribute_value_date_] = a_attribute_value_dates
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of A_ATTRIBUTE_VALUE_ENUMERATION
	case A_ATTRIBUTE_VALUE_ENUMERATION:
		switch fieldname {
		// insertion point for per direct association field
		case "ATTRIBUTE_VALUE_ENUMERATION":
			res := make(map[*ATTRIBUTE_VALUE_ENUMERATION][]*A_ATTRIBUTE_VALUE_ENUMERATION)
			for a_attribute_value_enumeration := range stage.A_ATTRIBUTE_VALUE_ENUMERATIONs {
				if a_attribute_value_enumeration.ATTRIBUTE_VALUE_ENUMERATION != nil {
					attribute_value_enumeration_ := a_attribute_value_enumeration.ATTRIBUTE_VALUE_ENUMERATION
					var a_attribute_value_enumerations []*A_ATTRIBUTE_VALUE_ENUMERATION
					_, ok := res[attribute_value_enumeration_]
					if ok {
						a_attribute_value_enumerations = res[attribute_value_enumeration_]
					} else {
						a_attribute_value_enumerations = make([]*A_ATTRIBUTE_VALUE_ENUMERATION, 0)
					}
					a_attribute_value_enumerations = append(a_attribute_value_enumerations, a_attribute_value_enumeration)
					res[attribute_value_enumeration_] = a_attribute_value_enumerations
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of A_ATTRIBUTE_VALUE_INTEGER
	case A_ATTRIBUTE_VALUE_INTEGER:
		switch fieldname {
		// insertion point for per direct association field
		case "ATTRIBUTE_VALUE_INTEGER":
			res := make(map[*ATTRIBUTE_VALUE_INTEGER][]*A_ATTRIBUTE_VALUE_INTEGER)
			for a_attribute_value_integer := range stage.A_ATTRIBUTE_VALUE_INTEGERs {
				if a_attribute_value_integer.ATTRIBUTE_VALUE_INTEGER != nil {
					attribute_value_integer_ := a_attribute_value_integer.ATTRIBUTE_VALUE_INTEGER
					var a_attribute_value_integers []*A_ATTRIBUTE_VALUE_INTEGER
					_, ok := res[attribute_value_integer_]
					if ok {
						a_attribute_value_integers = res[attribute_value_integer_]
					} else {
						a_attribute_value_integers = make([]*A_ATTRIBUTE_VALUE_INTEGER, 0)
					}
					a_attribute_value_integers = append(a_attribute_value_integers, a_attribute_value_integer)
					res[attribute_value_integer_] = a_attribute_value_integers
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of A_ATTRIBUTE_VALUE_REAL
	case A_ATTRIBUTE_VALUE_REAL:
		switch fieldname {
		// insertion point for per direct association field
		case "ATTRIBUTE_VALUE_REAL":
			res := make(map[*ATTRIBUTE_VALUE_REAL][]*A_ATTRIBUTE_VALUE_REAL)
			for a_attribute_value_real := range stage.A_ATTRIBUTE_VALUE_REALs {
				if a_attribute_value_real.ATTRIBUTE_VALUE_REAL != nil {
					attribute_value_real_ := a_attribute_value_real.ATTRIBUTE_VALUE_REAL
					var a_attribute_value_reals []*A_ATTRIBUTE_VALUE_REAL
					_, ok := res[attribute_value_real_]
					if ok {
						a_attribute_value_reals = res[attribute_value_real_]
					} else {
						a_attribute_value_reals = make([]*A_ATTRIBUTE_VALUE_REAL, 0)
					}
					a_attribute_value_reals = append(a_attribute_value_reals, a_attribute_value_real)
					res[attribute_value_real_] = a_attribute_value_reals
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of A_ATTRIBUTE_VALUE_STRING
	case A_ATTRIBUTE_VALUE_STRING:
		switch fieldname {
		// insertion point for per direct association field
		case "ATTRIBUTE_VALUE_STRING":
			res := make(map[*ATTRIBUTE_VALUE_STRING][]*A_ATTRIBUTE_VALUE_STRING)
			for a_attribute_value_string := range stage.A_ATTRIBUTE_VALUE_STRINGs {
				if a_attribute_value_string.ATTRIBUTE_VALUE_STRING != nil {
					attribute_value_string_ := a_attribute_value_string.ATTRIBUTE_VALUE_STRING
					var a_attribute_value_strings []*A_ATTRIBUTE_VALUE_STRING
					_, ok := res[attribute_value_string_]
					if ok {
						a_attribute_value_strings = res[attribute_value_string_]
					} else {
						a_attribute_value_strings = make([]*A_ATTRIBUTE_VALUE_STRING, 0)
					}
					a_attribute_value_strings = append(a_attribute_value_strings, a_attribute_value_string)
					res[attribute_value_string_] = a_attribute_value_strings
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of A_ATTRIBUTE_VALUE_XHTML
	case A_ATTRIBUTE_VALUE_XHTML:
		switch fieldname {
		// insertion point for per direct association field
		case "ATTRIBUTE_VALUE_BOOLEAN":
			res := make(map[*ATTRIBUTE_VALUE_BOOLEAN][]*A_ATTRIBUTE_VALUE_XHTML)
			for a_attribute_value_xhtml := range stage.A_ATTRIBUTE_VALUE_XHTMLs {
				if a_attribute_value_xhtml.ATTRIBUTE_VALUE_BOOLEAN != nil {
					attribute_value_boolean_ := a_attribute_value_xhtml.ATTRIBUTE_VALUE_BOOLEAN
					var a_attribute_value_xhtmls []*A_ATTRIBUTE_VALUE_XHTML
					_, ok := res[attribute_value_boolean_]
					if ok {
						a_attribute_value_xhtmls = res[attribute_value_boolean_]
					} else {
						a_attribute_value_xhtmls = make([]*A_ATTRIBUTE_VALUE_XHTML, 0)
					}
					a_attribute_value_xhtmls = append(a_attribute_value_xhtmls, a_attribute_value_xhtml)
					res[attribute_value_boolean_] = a_attribute_value_xhtmls
				}
			}
			return any(res).(map[*End][]*Start)
		case "ATTRIBUTE_VALUE_DATE":
			res := make(map[*ATTRIBUTE_VALUE_DATE][]*A_ATTRIBUTE_VALUE_XHTML)
			for a_attribute_value_xhtml := range stage.A_ATTRIBUTE_VALUE_XHTMLs {
				if a_attribute_value_xhtml.ATTRIBUTE_VALUE_DATE != nil {
					attribute_value_date_ := a_attribute_value_xhtml.ATTRIBUTE_VALUE_DATE
					var a_attribute_value_xhtmls []*A_ATTRIBUTE_VALUE_XHTML
					_, ok := res[attribute_value_date_]
					if ok {
						a_attribute_value_xhtmls = res[attribute_value_date_]
					} else {
						a_attribute_value_xhtmls = make([]*A_ATTRIBUTE_VALUE_XHTML, 0)
					}
					a_attribute_value_xhtmls = append(a_attribute_value_xhtmls, a_attribute_value_xhtml)
					res[attribute_value_date_] = a_attribute_value_xhtmls
				}
			}
			return any(res).(map[*End][]*Start)
		case "ATTRIBUTE_VALUE_ENUMERATION":
			res := make(map[*ATTRIBUTE_VALUE_ENUMERATION][]*A_ATTRIBUTE_VALUE_XHTML)
			for a_attribute_value_xhtml := range stage.A_ATTRIBUTE_VALUE_XHTMLs {
				if a_attribute_value_xhtml.ATTRIBUTE_VALUE_ENUMERATION != nil {
					attribute_value_enumeration_ := a_attribute_value_xhtml.ATTRIBUTE_VALUE_ENUMERATION
					var a_attribute_value_xhtmls []*A_ATTRIBUTE_VALUE_XHTML
					_, ok := res[attribute_value_enumeration_]
					if ok {
						a_attribute_value_xhtmls = res[attribute_value_enumeration_]
					} else {
						a_attribute_value_xhtmls = make([]*A_ATTRIBUTE_VALUE_XHTML, 0)
					}
					a_attribute_value_xhtmls = append(a_attribute_value_xhtmls, a_attribute_value_xhtml)
					res[attribute_value_enumeration_] = a_attribute_value_xhtmls
				}
			}
			return any(res).(map[*End][]*Start)
		case "ATTRIBUTE_VALUE_INTEGER":
			res := make(map[*ATTRIBUTE_VALUE_INTEGER][]*A_ATTRIBUTE_VALUE_XHTML)
			for a_attribute_value_xhtml := range stage.A_ATTRIBUTE_VALUE_XHTMLs {
				if a_attribute_value_xhtml.ATTRIBUTE_VALUE_INTEGER != nil {
					attribute_value_integer_ := a_attribute_value_xhtml.ATTRIBUTE_VALUE_INTEGER
					var a_attribute_value_xhtmls []*A_ATTRIBUTE_VALUE_XHTML
					_, ok := res[attribute_value_integer_]
					if ok {
						a_attribute_value_xhtmls = res[attribute_value_integer_]
					} else {
						a_attribute_value_xhtmls = make([]*A_ATTRIBUTE_VALUE_XHTML, 0)
					}
					a_attribute_value_xhtmls = append(a_attribute_value_xhtmls, a_attribute_value_xhtml)
					res[attribute_value_integer_] = a_attribute_value_xhtmls
				}
			}
			return any(res).(map[*End][]*Start)
		case "ATTRIBUTE_VALUE_REAL":
			res := make(map[*ATTRIBUTE_VALUE_REAL][]*A_ATTRIBUTE_VALUE_XHTML)
			for a_attribute_value_xhtml := range stage.A_ATTRIBUTE_VALUE_XHTMLs {
				if a_attribute_value_xhtml.ATTRIBUTE_VALUE_REAL != nil {
					attribute_value_real_ := a_attribute_value_xhtml.ATTRIBUTE_VALUE_REAL
					var a_attribute_value_xhtmls []*A_ATTRIBUTE_VALUE_XHTML
					_, ok := res[attribute_value_real_]
					if ok {
						a_attribute_value_xhtmls = res[attribute_value_real_]
					} else {
						a_attribute_value_xhtmls = make([]*A_ATTRIBUTE_VALUE_XHTML, 0)
					}
					a_attribute_value_xhtmls = append(a_attribute_value_xhtmls, a_attribute_value_xhtml)
					res[attribute_value_real_] = a_attribute_value_xhtmls
				}
			}
			return any(res).(map[*End][]*Start)
		case "ATTRIBUTE_VALUE_STRING":
			res := make(map[*ATTRIBUTE_VALUE_STRING][]*A_ATTRIBUTE_VALUE_XHTML)
			for a_attribute_value_xhtml := range stage.A_ATTRIBUTE_VALUE_XHTMLs {
				if a_attribute_value_xhtml.ATTRIBUTE_VALUE_STRING != nil {
					attribute_value_string_ := a_attribute_value_xhtml.ATTRIBUTE_VALUE_STRING
					var a_attribute_value_xhtmls []*A_ATTRIBUTE_VALUE_XHTML
					_, ok := res[attribute_value_string_]
					if ok {
						a_attribute_value_xhtmls = res[attribute_value_string_]
					} else {
						a_attribute_value_xhtmls = make([]*A_ATTRIBUTE_VALUE_XHTML, 0)
					}
					a_attribute_value_xhtmls = append(a_attribute_value_xhtmls, a_attribute_value_xhtml)
					res[attribute_value_string_] = a_attribute_value_xhtmls
				}
			}
			return any(res).(map[*End][]*Start)
		case "ATTRIBUTE_VALUE_XHTML":
			res := make(map[*ATTRIBUTE_VALUE_XHTML][]*A_ATTRIBUTE_VALUE_XHTML)
			for a_attribute_value_xhtml := range stage.A_ATTRIBUTE_VALUE_XHTMLs {
				if a_attribute_value_xhtml.ATTRIBUTE_VALUE_XHTML != nil {
					attribute_value_xhtml_ := a_attribute_value_xhtml.ATTRIBUTE_VALUE_XHTML
					var a_attribute_value_xhtmls []*A_ATTRIBUTE_VALUE_XHTML
					_, ok := res[attribute_value_xhtml_]
					if ok {
						a_attribute_value_xhtmls = res[attribute_value_xhtml_]
					} else {
						a_attribute_value_xhtmls = make([]*A_ATTRIBUTE_VALUE_XHTML, 0)
					}
					a_attribute_value_xhtmls = append(a_attribute_value_xhtmls, a_attribute_value_xhtml)
					res[attribute_value_xhtml_] = a_attribute_value_xhtmls
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of A_ATTRIBUTE_VALUE_XHTML_1
	case A_ATTRIBUTE_VALUE_XHTML_1:
		switch fieldname {
		// insertion point for per direct association field
		case "ATTRIBUTE_VALUE_XHTML":
			res := make(map[*ATTRIBUTE_VALUE_XHTML][]*A_ATTRIBUTE_VALUE_XHTML_1)
			for a_attribute_value_xhtml_1 := range stage.A_ATTRIBUTE_VALUE_XHTML_1s {
				if a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_XHTML != nil {
					attribute_value_xhtml_ := a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_XHTML
					var a_attribute_value_xhtml_1s []*A_ATTRIBUTE_VALUE_XHTML_1
					_, ok := res[attribute_value_xhtml_]
					if ok {
						a_attribute_value_xhtml_1s = res[attribute_value_xhtml_]
					} else {
						a_attribute_value_xhtml_1s = make([]*A_ATTRIBUTE_VALUE_XHTML_1, 0)
					}
					a_attribute_value_xhtml_1s = append(a_attribute_value_xhtml_1s, a_attribute_value_xhtml_1)
					res[attribute_value_xhtml_] = a_attribute_value_xhtml_1s
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of A_CHILDREN
	case A_CHILDREN:
		switch fieldname {
		// insertion point for per direct association field
		case "SPEC_HIERARCHY":
			res := make(map[*SPEC_HIERARCHY][]*A_CHILDREN)
			for a_children := range stage.A_CHILDRENs {
				if a_children.SPEC_HIERARCHY != nil {
					spec_hierarchy_ := a_children.SPEC_HIERARCHY
					var a_childrens []*A_CHILDREN
					_, ok := res[spec_hierarchy_]
					if ok {
						a_childrens = res[spec_hierarchy_]
					} else {
						a_childrens = make([]*A_CHILDREN, 0)
					}
					a_childrens = append(a_childrens, a_children)
					res[spec_hierarchy_] = a_childrens
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of A_CORE_CONTENT
	case A_CORE_CONTENT:
		switch fieldname {
		// insertion point for per direct association field
		case "REQ_IF_CONTENT":
			res := make(map[*REQ_IF_CONTENT][]*A_CORE_CONTENT)
			for a_core_content := range stage.A_CORE_CONTENTs {
				if a_core_content.REQ_IF_CONTENT != nil {
					req_if_content_ := a_core_content.REQ_IF_CONTENT
					var a_core_contents []*A_CORE_CONTENT
					_, ok := res[req_if_content_]
					if ok {
						a_core_contents = res[req_if_content_]
					} else {
						a_core_contents = make([]*A_CORE_CONTENT, 0)
					}
					a_core_contents = append(a_core_contents, a_core_content)
					res[req_if_content_] = a_core_contents
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of A_DATATYPES
	case A_DATATYPES:
		switch fieldname {
		// insertion point for per direct association field
		case "DATATYPE_DEFINITION_BOOLEAN":
			res := make(map[*DATATYPE_DEFINITION_BOOLEAN][]*A_DATATYPES)
			for a_datatypes := range stage.A_DATATYPESs {
				if a_datatypes.DATATYPE_DEFINITION_BOOLEAN != nil {
					datatype_definition_boolean_ := a_datatypes.DATATYPE_DEFINITION_BOOLEAN
					var a_datatypess []*A_DATATYPES
					_, ok := res[datatype_definition_boolean_]
					if ok {
						a_datatypess = res[datatype_definition_boolean_]
					} else {
						a_datatypess = make([]*A_DATATYPES, 0)
					}
					a_datatypess = append(a_datatypess, a_datatypes)
					res[datatype_definition_boolean_] = a_datatypess
				}
			}
			return any(res).(map[*End][]*Start)
		case "DATATYPE_DEFINITION_DATE":
			res := make(map[*DATATYPE_DEFINITION_DATE][]*A_DATATYPES)
			for a_datatypes := range stage.A_DATATYPESs {
				if a_datatypes.DATATYPE_DEFINITION_DATE != nil {
					datatype_definition_date_ := a_datatypes.DATATYPE_DEFINITION_DATE
					var a_datatypess []*A_DATATYPES
					_, ok := res[datatype_definition_date_]
					if ok {
						a_datatypess = res[datatype_definition_date_]
					} else {
						a_datatypess = make([]*A_DATATYPES, 0)
					}
					a_datatypess = append(a_datatypess, a_datatypes)
					res[datatype_definition_date_] = a_datatypess
				}
			}
			return any(res).(map[*End][]*Start)
		case "DATATYPE_DEFINITION_ENUMERATION":
			res := make(map[*DATATYPE_DEFINITION_ENUMERATION][]*A_DATATYPES)
			for a_datatypes := range stage.A_DATATYPESs {
				if a_datatypes.DATATYPE_DEFINITION_ENUMERATION != nil {
					datatype_definition_enumeration_ := a_datatypes.DATATYPE_DEFINITION_ENUMERATION
					var a_datatypess []*A_DATATYPES
					_, ok := res[datatype_definition_enumeration_]
					if ok {
						a_datatypess = res[datatype_definition_enumeration_]
					} else {
						a_datatypess = make([]*A_DATATYPES, 0)
					}
					a_datatypess = append(a_datatypess, a_datatypes)
					res[datatype_definition_enumeration_] = a_datatypess
				}
			}
			return any(res).(map[*End][]*Start)
		case "DATATYPE_DEFINITION_INTEGER":
			res := make(map[*DATATYPE_DEFINITION_INTEGER][]*A_DATATYPES)
			for a_datatypes := range stage.A_DATATYPESs {
				if a_datatypes.DATATYPE_DEFINITION_INTEGER != nil {
					datatype_definition_integer_ := a_datatypes.DATATYPE_DEFINITION_INTEGER
					var a_datatypess []*A_DATATYPES
					_, ok := res[datatype_definition_integer_]
					if ok {
						a_datatypess = res[datatype_definition_integer_]
					} else {
						a_datatypess = make([]*A_DATATYPES, 0)
					}
					a_datatypess = append(a_datatypess, a_datatypes)
					res[datatype_definition_integer_] = a_datatypess
				}
			}
			return any(res).(map[*End][]*Start)
		case "DATATYPE_DEFINITION_REAL":
			res := make(map[*DATATYPE_DEFINITION_REAL][]*A_DATATYPES)
			for a_datatypes := range stage.A_DATATYPESs {
				if a_datatypes.DATATYPE_DEFINITION_REAL != nil {
					datatype_definition_real_ := a_datatypes.DATATYPE_DEFINITION_REAL
					var a_datatypess []*A_DATATYPES
					_, ok := res[datatype_definition_real_]
					if ok {
						a_datatypess = res[datatype_definition_real_]
					} else {
						a_datatypess = make([]*A_DATATYPES, 0)
					}
					a_datatypess = append(a_datatypess, a_datatypes)
					res[datatype_definition_real_] = a_datatypess
				}
			}
			return any(res).(map[*End][]*Start)
		case "DATATYPE_DEFINITION_STRING":
			res := make(map[*DATATYPE_DEFINITION_STRING][]*A_DATATYPES)
			for a_datatypes := range stage.A_DATATYPESs {
				if a_datatypes.DATATYPE_DEFINITION_STRING != nil {
					datatype_definition_string_ := a_datatypes.DATATYPE_DEFINITION_STRING
					var a_datatypess []*A_DATATYPES
					_, ok := res[datatype_definition_string_]
					if ok {
						a_datatypess = res[datatype_definition_string_]
					} else {
						a_datatypess = make([]*A_DATATYPES, 0)
					}
					a_datatypess = append(a_datatypess, a_datatypes)
					res[datatype_definition_string_] = a_datatypess
				}
			}
			return any(res).(map[*End][]*Start)
		case "DATATYPE_DEFINITION_XHTML":
			res := make(map[*DATATYPE_DEFINITION_XHTML][]*A_DATATYPES)
			for a_datatypes := range stage.A_DATATYPESs {
				if a_datatypes.DATATYPE_DEFINITION_XHTML != nil {
					datatype_definition_xhtml_ := a_datatypes.DATATYPE_DEFINITION_XHTML
					var a_datatypess []*A_DATATYPES
					_, ok := res[datatype_definition_xhtml_]
					if ok {
						a_datatypess = res[datatype_definition_xhtml_]
					} else {
						a_datatypess = make([]*A_DATATYPES, 0)
					}
					a_datatypess = append(a_datatypess, a_datatypes)
					res[datatype_definition_xhtml_] = a_datatypess
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of A_DATATYPE_DEFINITION_BOOLEAN_REF
	case A_DATATYPE_DEFINITION_BOOLEAN_REF:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_DATATYPE_DEFINITION_DATE_REF
	case A_DATATYPE_DEFINITION_DATE_REF:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_DATATYPE_DEFINITION_ENUMERATION_REF
	case A_DATATYPE_DEFINITION_ENUMERATION_REF:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_DATATYPE_DEFINITION_INTEGER_REF
	case A_DATATYPE_DEFINITION_INTEGER_REF:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_DATATYPE_DEFINITION_REAL_REF
	case A_DATATYPE_DEFINITION_REAL_REF:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_DATATYPE_DEFINITION_STRING_REF
	case A_DATATYPE_DEFINITION_STRING_REF:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_DATATYPE_DEFINITION_XHTML_REF
	case A_DATATYPE_DEFINITION_XHTML_REF:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_EDITABLE_ATTS
	case A_EDITABLE_ATTS:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_ENUM_VALUE_REF
	case A_ENUM_VALUE_REF:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_OBJECT
	case A_OBJECT:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_PROPERTIES
	case A_PROPERTIES:
		switch fieldname {
		// insertion point for per direct association field
		case "EMBEDDED_VALUE":
			res := make(map[*EMBEDDED_VALUE][]*A_PROPERTIES)
			for a_properties := range stage.A_PROPERTIESs {
				if a_properties.EMBEDDED_VALUE != nil {
					embedded_value_ := a_properties.EMBEDDED_VALUE
					var a_propertiess []*A_PROPERTIES
					_, ok := res[embedded_value_]
					if ok {
						a_propertiess = res[embedded_value_]
					} else {
						a_propertiess = make([]*A_PROPERTIES, 0)
					}
					a_propertiess = append(a_propertiess, a_properties)
					res[embedded_value_] = a_propertiess
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of A_RELATION_GROUP_TYPE_REF
	case A_RELATION_GROUP_TYPE_REF:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_SOURCE_1
	case A_SOURCE_1:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_SPECIFICATIONS
	case A_SPECIFICATIONS:
		switch fieldname {
		// insertion point for per direct association field
		case "SPECIFICATION":
			res := make(map[*SPECIFICATION][]*A_SPECIFICATIONS)
			for a_specifications := range stage.A_SPECIFICATIONSs {
				if a_specifications.SPECIFICATION != nil {
					specification_ := a_specifications.SPECIFICATION
					var a_specificationss []*A_SPECIFICATIONS
					_, ok := res[specification_]
					if ok {
						a_specificationss = res[specification_]
					} else {
						a_specificationss = make([]*A_SPECIFICATIONS, 0)
					}
					a_specificationss = append(a_specificationss, a_specifications)
					res[specification_] = a_specificationss
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of A_SPECIFICATION_TYPE_REF
	case A_SPECIFICATION_TYPE_REF:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_SPECIFIED_VALUES
	case A_SPECIFIED_VALUES:
		switch fieldname {
		// insertion point for per direct association field
		case "ENUM_VALUE":
			res := make(map[*ENUM_VALUE][]*A_SPECIFIED_VALUES)
			for a_specified_values := range stage.A_SPECIFIED_VALUESs {
				if a_specified_values.ENUM_VALUE != nil {
					enum_value_ := a_specified_values.ENUM_VALUE
					var a_specified_valuess []*A_SPECIFIED_VALUES
					_, ok := res[enum_value_]
					if ok {
						a_specified_valuess = res[enum_value_]
					} else {
						a_specified_valuess = make([]*A_SPECIFIED_VALUES, 0)
					}
					a_specified_valuess = append(a_specified_valuess, a_specified_values)
					res[enum_value_] = a_specified_valuess
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of A_SPEC_ATTRIBUTES
	case A_SPEC_ATTRIBUTES:
		switch fieldname {
		// insertion point for per direct association field
		case "ATTRIBUTE_DEFINITION_BOOLEAN":
			res := make(map[*ATTRIBUTE_DEFINITION_BOOLEAN][]*A_SPEC_ATTRIBUTES)
			for a_spec_attributes := range stage.A_SPEC_ATTRIBUTESs {
				if a_spec_attributes.ATTRIBUTE_DEFINITION_BOOLEAN != nil {
					attribute_definition_boolean_ := a_spec_attributes.ATTRIBUTE_DEFINITION_BOOLEAN
					var a_spec_attributess []*A_SPEC_ATTRIBUTES
					_, ok := res[attribute_definition_boolean_]
					if ok {
						a_spec_attributess = res[attribute_definition_boolean_]
					} else {
						a_spec_attributess = make([]*A_SPEC_ATTRIBUTES, 0)
					}
					a_spec_attributess = append(a_spec_attributess, a_spec_attributes)
					res[attribute_definition_boolean_] = a_spec_attributess
				}
			}
			return any(res).(map[*End][]*Start)
		case "ATTRIBUTE_DEFINITION_DATE":
			res := make(map[*ATTRIBUTE_DEFINITION_DATE][]*A_SPEC_ATTRIBUTES)
			for a_spec_attributes := range stage.A_SPEC_ATTRIBUTESs {
				if a_spec_attributes.ATTRIBUTE_DEFINITION_DATE != nil {
					attribute_definition_date_ := a_spec_attributes.ATTRIBUTE_DEFINITION_DATE
					var a_spec_attributess []*A_SPEC_ATTRIBUTES
					_, ok := res[attribute_definition_date_]
					if ok {
						a_spec_attributess = res[attribute_definition_date_]
					} else {
						a_spec_attributess = make([]*A_SPEC_ATTRIBUTES, 0)
					}
					a_spec_attributess = append(a_spec_attributess, a_spec_attributes)
					res[attribute_definition_date_] = a_spec_attributess
				}
			}
			return any(res).(map[*End][]*Start)
		case "ATTRIBUTE_DEFINITION_ENUMERATION":
			res := make(map[*ATTRIBUTE_DEFINITION_ENUMERATION][]*A_SPEC_ATTRIBUTES)
			for a_spec_attributes := range stage.A_SPEC_ATTRIBUTESs {
				if a_spec_attributes.ATTRIBUTE_DEFINITION_ENUMERATION != nil {
					attribute_definition_enumeration_ := a_spec_attributes.ATTRIBUTE_DEFINITION_ENUMERATION
					var a_spec_attributess []*A_SPEC_ATTRIBUTES
					_, ok := res[attribute_definition_enumeration_]
					if ok {
						a_spec_attributess = res[attribute_definition_enumeration_]
					} else {
						a_spec_attributess = make([]*A_SPEC_ATTRIBUTES, 0)
					}
					a_spec_attributess = append(a_spec_attributess, a_spec_attributes)
					res[attribute_definition_enumeration_] = a_spec_attributess
				}
			}
			return any(res).(map[*End][]*Start)
		case "ATTRIBUTE_DEFINITION_INTEGER":
			res := make(map[*ATTRIBUTE_DEFINITION_INTEGER][]*A_SPEC_ATTRIBUTES)
			for a_spec_attributes := range stage.A_SPEC_ATTRIBUTESs {
				if a_spec_attributes.ATTRIBUTE_DEFINITION_INTEGER != nil {
					attribute_definition_integer_ := a_spec_attributes.ATTRIBUTE_DEFINITION_INTEGER
					var a_spec_attributess []*A_SPEC_ATTRIBUTES
					_, ok := res[attribute_definition_integer_]
					if ok {
						a_spec_attributess = res[attribute_definition_integer_]
					} else {
						a_spec_attributess = make([]*A_SPEC_ATTRIBUTES, 0)
					}
					a_spec_attributess = append(a_spec_attributess, a_spec_attributes)
					res[attribute_definition_integer_] = a_spec_attributess
				}
			}
			return any(res).(map[*End][]*Start)
		case "ATTRIBUTE_DEFINITION_REAL":
			res := make(map[*ATTRIBUTE_DEFINITION_REAL][]*A_SPEC_ATTRIBUTES)
			for a_spec_attributes := range stage.A_SPEC_ATTRIBUTESs {
				if a_spec_attributes.ATTRIBUTE_DEFINITION_REAL != nil {
					attribute_definition_real_ := a_spec_attributes.ATTRIBUTE_DEFINITION_REAL
					var a_spec_attributess []*A_SPEC_ATTRIBUTES
					_, ok := res[attribute_definition_real_]
					if ok {
						a_spec_attributess = res[attribute_definition_real_]
					} else {
						a_spec_attributess = make([]*A_SPEC_ATTRIBUTES, 0)
					}
					a_spec_attributess = append(a_spec_attributess, a_spec_attributes)
					res[attribute_definition_real_] = a_spec_attributess
				}
			}
			return any(res).(map[*End][]*Start)
		case "ATTRIBUTE_DEFINITION_STRING":
			res := make(map[*ATTRIBUTE_DEFINITION_STRING][]*A_SPEC_ATTRIBUTES)
			for a_spec_attributes := range stage.A_SPEC_ATTRIBUTESs {
				if a_spec_attributes.ATTRIBUTE_DEFINITION_STRING != nil {
					attribute_definition_string_ := a_spec_attributes.ATTRIBUTE_DEFINITION_STRING
					var a_spec_attributess []*A_SPEC_ATTRIBUTES
					_, ok := res[attribute_definition_string_]
					if ok {
						a_spec_attributess = res[attribute_definition_string_]
					} else {
						a_spec_attributess = make([]*A_SPEC_ATTRIBUTES, 0)
					}
					a_spec_attributess = append(a_spec_attributess, a_spec_attributes)
					res[attribute_definition_string_] = a_spec_attributess
				}
			}
			return any(res).(map[*End][]*Start)
		case "ATTRIBUTE_DEFINITION_XHTML":
			res := make(map[*ATTRIBUTE_DEFINITION_XHTML][]*A_SPEC_ATTRIBUTES)
			for a_spec_attributes := range stage.A_SPEC_ATTRIBUTESs {
				if a_spec_attributes.ATTRIBUTE_DEFINITION_XHTML != nil {
					attribute_definition_xhtml_ := a_spec_attributes.ATTRIBUTE_DEFINITION_XHTML
					var a_spec_attributess []*A_SPEC_ATTRIBUTES
					_, ok := res[attribute_definition_xhtml_]
					if ok {
						a_spec_attributess = res[attribute_definition_xhtml_]
					} else {
						a_spec_attributess = make([]*A_SPEC_ATTRIBUTES, 0)
					}
					a_spec_attributess = append(a_spec_attributess, a_spec_attributes)
					res[attribute_definition_xhtml_] = a_spec_attributess
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of A_SPEC_OBJECTS
	case A_SPEC_OBJECTS:
		switch fieldname {
		// insertion point for per direct association field
		case "SPEC_OBJECT":
			res := make(map[*SPEC_OBJECT][]*A_SPEC_OBJECTS)
			for a_spec_objects := range stage.A_SPEC_OBJECTSs {
				if a_spec_objects.SPEC_OBJECT != nil {
					spec_object_ := a_spec_objects.SPEC_OBJECT
					var a_spec_objectss []*A_SPEC_OBJECTS
					_, ok := res[spec_object_]
					if ok {
						a_spec_objectss = res[spec_object_]
					} else {
						a_spec_objectss = make([]*A_SPEC_OBJECTS, 0)
					}
					a_spec_objectss = append(a_spec_objectss, a_spec_objects)
					res[spec_object_] = a_spec_objectss
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of A_SPEC_OBJECT_TYPE_REF
	case A_SPEC_OBJECT_TYPE_REF:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_SPEC_RELATIONS
	case A_SPEC_RELATIONS:
		switch fieldname {
		// insertion point for per direct association field
		case "SPEC_RELATION":
			res := make(map[*SPEC_RELATION][]*A_SPEC_RELATIONS)
			for a_spec_relations := range stage.A_SPEC_RELATIONSs {
				if a_spec_relations.SPEC_RELATION != nil {
					spec_relation_ := a_spec_relations.SPEC_RELATION
					var a_spec_relationss []*A_SPEC_RELATIONS
					_, ok := res[spec_relation_]
					if ok {
						a_spec_relationss = res[spec_relation_]
					} else {
						a_spec_relationss = make([]*A_SPEC_RELATIONS, 0)
					}
					a_spec_relationss = append(a_spec_relationss, a_spec_relations)
					res[spec_relation_] = a_spec_relationss
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of A_SPEC_RELATION_GROUPS
	case A_SPEC_RELATION_GROUPS:
		switch fieldname {
		// insertion point for per direct association field
		case "RELATION_GROUP":
			res := make(map[*RELATION_GROUP][]*A_SPEC_RELATION_GROUPS)
			for a_spec_relation_groups := range stage.A_SPEC_RELATION_GROUPSs {
				if a_spec_relation_groups.RELATION_GROUP != nil {
					relation_group_ := a_spec_relation_groups.RELATION_GROUP
					var a_spec_relation_groupss []*A_SPEC_RELATION_GROUPS
					_, ok := res[relation_group_]
					if ok {
						a_spec_relation_groupss = res[relation_group_]
					} else {
						a_spec_relation_groupss = make([]*A_SPEC_RELATION_GROUPS, 0)
					}
					a_spec_relation_groupss = append(a_spec_relation_groupss, a_spec_relation_groups)
					res[relation_group_] = a_spec_relation_groupss
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of A_SPEC_RELATION_REF
	case A_SPEC_RELATION_REF:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_SPEC_RELATION_TYPE_REF
	case A_SPEC_RELATION_TYPE_REF:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_SPEC_TYPES
	case A_SPEC_TYPES:
		switch fieldname {
		// insertion point for per direct association field
		case "RELATION_GROUP_TYPE":
			res := make(map[*RELATION_GROUP_TYPE][]*A_SPEC_TYPES)
			for a_spec_types := range stage.A_SPEC_TYPESs {
				if a_spec_types.RELATION_GROUP_TYPE != nil {
					relation_group_type_ := a_spec_types.RELATION_GROUP_TYPE
					var a_spec_typess []*A_SPEC_TYPES
					_, ok := res[relation_group_type_]
					if ok {
						a_spec_typess = res[relation_group_type_]
					} else {
						a_spec_typess = make([]*A_SPEC_TYPES, 0)
					}
					a_spec_typess = append(a_spec_typess, a_spec_types)
					res[relation_group_type_] = a_spec_typess
				}
			}
			return any(res).(map[*End][]*Start)
		case "SPEC_OBJECT_TYPE":
			res := make(map[*SPEC_OBJECT_TYPE][]*A_SPEC_TYPES)
			for a_spec_types := range stage.A_SPEC_TYPESs {
				if a_spec_types.SPEC_OBJECT_TYPE != nil {
					spec_object_type_ := a_spec_types.SPEC_OBJECT_TYPE
					var a_spec_typess []*A_SPEC_TYPES
					_, ok := res[spec_object_type_]
					if ok {
						a_spec_typess = res[spec_object_type_]
					} else {
						a_spec_typess = make([]*A_SPEC_TYPES, 0)
					}
					a_spec_typess = append(a_spec_typess, a_spec_types)
					res[spec_object_type_] = a_spec_typess
				}
			}
			return any(res).(map[*End][]*Start)
		case "SPEC_RELATION_TYPE":
			res := make(map[*SPEC_RELATION_TYPE][]*A_SPEC_TYPES)
			for a_spec_types := range stage.A_SPEC_TYPESs {
				if a_spec_types.SPEC_RELATION_TYPE != nil {
					spec_relation_type_ := a_spec_types.SPEC_RELATION_TYPE
					var a_spec_typess []*A_SPEC_TYPES
					_, ok := res[spec_relation_type_]
					if ok {
						a_spec_typess = res[spec_relation_type_]
					} else {
						a_spec_typess = make([]*A_SPEC_TYPES, 0)
					}
					a_spec_typess = append(a_spec_typess, a_spec_types)
					res[spec_relation_type_] = a_spec_typess
				}
			}
			return any(res).(map[*End][]*Start)
		case "SPECIFICATION_TYPE":
			res := make(map[*SPECIFICATION_TYPE][]*A_SPEC_TYPES)
			for a_spec_types := range stage.A_SPEC_TYPESs {
				if a_spec_types.SPECIFICATION_TYPE != nil {
					specification_type_ := a_spec_types.SPECIFICATION_TYPE
					var a_spec_typess []*A_SPEC_TYPES
					_, ok := res[specification_type_]
					if ok {
						a_spec_typess = res[specification_type_]
					} else {
						a_spec_typess = make([]*A_SPEC_TYPES, 0)
					}
					a_spec_typess = append(a_spec_typess, a_spec_types)
					res[specification_type_] = a_spec_typess
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of A_TARGET_SPECIFICATION_1
	case A_TARGET_SPECIFICATION_1:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_THE_HEADER
	case A_THE_HEADER:
		switch fieldname {
		// insertion point for per direct association field
		case "REQ_IF_HEADER":
			res := make(map[*REQ_IF_HEADER][]*A_THE_HEADER)
			for a_the_header := range stage.A_THE_HEADERs {
				if a_the_header.REQ_IF_HEADER != nil {
					req_if_header_ := a_the_header.REQ_IF_HEADER
					var a_the_headers []*A_THE_HEADER
					_, ok := res[req_if_header_]
					if ok {
						a_the_headers = res[req_if_header_]
					} else {
						a_the_headers = make([]*A_THE_HEADER, 0)
					}
					a_the_headers = append(a_the_headers, a_the_header)
					res[req_if_header_] = a_the_headers
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of A_TOOL_EXTENSIONS
	case A_TOOL_EXTENSIONS:
		switch fieldname {
		// insertion point for per direct association field
		case "REQ_IF_TOOL_EXTENSION":
			res := make(map[*REQ_IF_TOOL_EXTENSION][]*A_TOOL_EXTENSIONS)
			for a_tool_extensions := range stage.A_TOOL_EXTENSIONSs {
				if a_tool_extensions.REQ_IF_TOOL_EXTENSION != nil {
					req_if_tool_extension_ := a_tool_extensions.REQ_IF_TOOL_EXTENSION
					var a_tool_extensionss []*A_TOOL_EXTENSIONS
					_, ok := res[req_if_tool_extension_]
					if ok {
						a_tool_extensionss = res[req_if_tool_extension_]
					} else {
						a_tool_extensionss = make([]*A_TOOL_EXTENSIONS, 0)
					}
					a_tool_extensionss = append(a_tool_extensionss, a_tool_extensions)
					res[req_if_tool_extension_] = a_tool_extensionss
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of DATATYPE_DEFINITION_BOOLEAN
	case DATATYPE_DEFINITION_BOOLEAN:
		switch fieldname {
		// insertion point for per direct association field
		case "ALTERNATIVE_ID":
			res := make(map[*A_ALTERNATIVE_ID][]*DATATYPE_DEFINITION_BOOLEAN)
			for datatype_definition_boolean := range stage.DATATYPE_DEFINITION_BOOLEANs {
				if datatype_definition_boolean.ALTERNATIVE_ID != nil {
					a_alternative_id_ := datatype_definition_boolean.ALTERNATIVE_ID
					var datatype_definition_booleans []*DATATYPE_DEFINITION_BOOLEAN
					_, ok := res[a_alternative_id_]
					if ok {
						datatype_definition_booleans = res[a_alternative_id_]
					} else {
						datatype_definition_booleans = make([]*DATATYPE_DEFINITION_BOOLEAN, 0)
					}
					datatype_definition_booleans = append(datatype_definition_booleans, datatype_definition_boolean)
					res[a_alternative_id_] = datatype_definition_booleans
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of DATATYPE_DEFINITION_DATE
	case DATATYPE_DEFINITION_DATE:
		switch fieldname {
		// insertion point for per direct association field
		case "ALTERNATIVE_ID":
			res := make(map[*A_ALTERNATIVE_ID][]*DATATYPE_DEFINITION_DATE)
			for datatype_definition_date := range stage.DATATYPE_DEFINITION_DATEs {
				if datatype_definition_date.ALTERNATIVE_ID != nil {
					a_alternative_id_ := datatype_definition_date.ALTERNATIVE_ID
					var datatype_definition_dates []*DATATYPE_DEFINITION_DATE
					_, ok := res[a_alternative_id_]
					if ok {
						datatype_definition_dates = res[a_alternative_id_]
					} else {
						datatype_definition_dates = make([]*DATATYPE_DEFINITION_DATE, 0)
					}
					datatype_definition_dates = append(datatype_definition_dates, datatype_definition_date)
					res[a_alternative_id_] = datatype_definition_dates
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of DATATYPE_DEFINITION_ENUMERATION
	case DATATYPE_DEFINITION_ENUMERATION:
		switch fieldname {
		// insertion point for per direct association field
		case "ALTERNATIVE_ID":
			res := make(map[*A_ALTERNATIVE_ID][]*DATATYPE_DEFINITION_ENUMERATION)
			for datatype_definition_enumeration := range stage.DATATYPE_DEFINITION_ENUMERATIONs {
				if datatype_definition_enumeration.ALTERNATIVE_ID != nil {
					a_alternative_id_ := datatype_definition_enumeration.ALTERNATIVE_ID
					var datatype_definition_enumerations []*DATATYPE_DEFINITION_ENUMERATION
					_, ok := res[a_alternative_id_]
					if ok {
						datatype_definition_enumerations = res[a_alternative_id_]
					} else {
						datatype_definition_enumerations = make([]*DATATYPE_DEFINITION_ENUMERATION, 0)
					}
					datatype_definition_enumerations = append(datatype_definition_enumerations, datatype_definition_enumeration)
					res[a_alternative_id_] = datatype_definition_enumerations
				}
			}
			return any(res).(map[*End][]*Start)
		case "SPECIFIED_VALUES":
			res := make(map[*A_SPECIFIED_VALUES][]*DATATYPE_DEFINITION_ENUMERATION)
			for datatype_definition_enumeration := range stage.DATATYPE_DEFINITION_ENUMERATIONs {
				if datatype_definition_enumeration.SPECIFIED_VALUES != nil {
					a_specified_values_ := datatype_definition_enumeration.SPECIFIED_VALUES
					var datatype_definition_enumerations []*DATATYPE_DEFINITION_ENUMERATION
					_, ok := res[a_specified_values_]
					if ok {
						datatype_definition_enumerations = res[a_specified_values_]
					} else {
						datatype_definition_enumerations = make([]*DATATYPE_DEFINITION_ENUMERATION, 0)
					}
					datatype_definition_enumerations = append(datatype_definition_enumerations, datatype_definition_enumeration)
					res[a_specified_values_] = datatype_definition_enumerations
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of DATATYPE_DEFINITION_INTEGER
	case DATATYPE_DEFINITION_INTEGER:
		switch fieldname {
		// insertion point for per direct association field
		case "ALTERNATIVE_ID":
			res := make(map[*A_ALTERNATIVE_ID][]*DATATYPE_DEFINITION_INTEGER)
			for datatype_definition_integer := range stage.DATATYPE_DEFINITION_INTEGERs {
				if datatype_definition_integer.ALTERNATIVE_ID != nil {
					a_alternative_id_ := datatype_definition_integer.ALTERNATIVE_ID
					var datatype_definition_integers []*DATATYPE_DEFINITION_INTEGER
					_, ok := res[a_alternative_id_]
					if ok {
						datatype_definition_integers = res[a_alternative_id_]
					} else {
						datatype_definition_integers = make([]*DATATYPE_DEFINITION_INTEGER, 0)
					}
					datatype_definition_integers = append(datatype_definition_integers, datatype_definition_integer)
					res[a_alternative_id_] = datatype_definition_integers
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of DATATYPE_DEFINITION_REAL
	case DATATYPE_DEFINITION_REAL:
		switch fieldname {
		// insertion point for per direct association field
		case "ALTERNATIVE_ID":
			res := make(map[*A_ALTERNATIVE_ID][]*DATATYPE_DEFINITION_REAL)
			for datatype_definition_real := range stage.DATATYPE_DEFINITION_REALs {
				if datatype_definition_real.ALTERNATIVE_ID != nil {
					a_alternative_id_ := datatype_definition_real.ALTERNATIVE_ID
					var datatype_definition_reals []*DATATYPE_DEFINITION_REAL
					_, ok := res[a_alternative_id_]
					if ok {
						datatype_definition_reals = res[a_alternative_id_]
					} else {
						datatype_definition_reals = make([]*DATATYPE_DEFINITION_REAL, 0)
					}
					datatype_definition_reals = append(datatype_definition_reals, datatype_definition_real)
					res[a_alternative_id_] = datatype_definition_reals
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of DATATYPE_DEFINITION_STRING
	case DATATYPE_DEFINITION_STRING:
		switch fieldname {
		// insertion point for per direct association field
		case "ALTERNATIVE_ID":
			res := make(map[*A_ALTERNATIVE_ID][]*DATATYPE_DEFINITION_STRING)
			for datatype_definition_string := range stage.DATATYPE_DEFINITION_STRINGs {
				if datatype_definition_string.ALTERNATIVE_ID != nil {
					a_alternative_id_ := datatype_definition_string.ALTERNATIVE_ID
					var datatype_definition_strings []*DATATYPE_DEFINITION_STRING
					_, ok := res[a_alternative_id_]
					if ok {
						datatype_definition_strings = res[a_alternative_id_]
					} else {
						datatype_definition_strings = make([]*DATATYPE_DEFINITION_STRING, 0)
					}
					datatype_definition_strings = append(datatype_definition_strings, datatype_definition_string)
					res[a_alternative_id_] = datatype_definition_strings
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of DATATYPE_DEFINITION_XHTML
	case DATATYPE_DEFINITION_XHTML:
		switch fieldname {
		// insertion point for per direct association field
		case "ALTERNATIVE_ID":
			res := make(map[*A_ALTERNATIVE_ID][]*DATATYPE_DEFINITION_XHTML)
			for datatype_definition_xhtml := range stage.DATATYPE_DEFINITION_XHTMLs {
				if datatype_definition_xhtml.ALTERNATIVE_ID != nil {
					a_alternative_id_ := datatype_definition_xhtml.ALTERNATIVE_ID
					var datatype_definition_xhtmls []*DATATYPE_DEFINITION_XHTML
					_, ok := res[a_alternative_id_]
					if ok {
						datatype_definition_xhtmls = res[a_alternative_id_]
					} else {
						datatype_definition_xhtmls = make([]*DATATYPE_DEFINITION_XHTML, 0)
					}
					datatype_definition_xhtmls = append(datatype_definition_xhtmls, datatype_definition_xhtml)
					res[a_alternative_id_] = datatype_definition_xhtmls
				}
			}
			return any(res).(map[*End][]*Start)
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
		case "ALTERNATIVE_ID":
			res := make(map[*A_ALTERNATIVE_ID][]*ENUM_VALUE)
			for enum_value := range stage.ENUM_VALUEs {
				if enum_value.ALTERNATIVE_ID != nil {
					a_alternative_id_ := enum_value.ALTERNATIVE_ID
					var enum_values []*ENUM_VALUE
					_, ok := res[a_alternative_id_]
					if ok {
						enum_values = res[a_alternative_id_]
					} else {
						enum_values = make([]*ENUM_VALUE, 0)
					}
					enum_values = append(enum_values, enum_value)
					res[a_alternative_id_] = enum_values
				}
			}
			return any(res).(map[*End][]*Start)
		case "PROPERTIES":
			res := make(map[*A_PROPERTIES][]*ENUM_VALUE)
			for enum_value := range stage.ENUM_VALUEs {
				if enum_value.PROPERTIES != nil {
					a_properties_ := enum_value.PROPERTIES
					var enum_values []*ENUM_VALUE
					_, ok := res[a_properties_]
					if ok {
						enum_values = res[a_properties_]
					} else {
						enum_values = make([]*ENUM_VALUE, 0)
					}
					enum_values = append(enum_values, enum_value)
					res[a_properties_] = enum_values
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of RELATION_GROUP
	case RELATION_GROUP:
		switch fieldname {
		// insertion point for per direct association field
		case "ALTERNATIVE_ID":
			res := make(map[*A_ALTERNATIVE_ID][]*RELATION_GROUP)
			for relation_group := range stage.RELATION_GROUPs {
				if relation_group.ALTERNATIVE_ID != nil {
					a_alternative_id_ := relation_group.ALTERNATIVE_ID
					var relation_groups []*RELATION_GROUP
					_, ok := res[a_alternative_id_]
					if ok {
						relation_groups = res[a_alternative_id_]
					} else {
						relation_groups = make([]*RELATION_GROUP, 0)
					}
					relation_groups = append(relation_groups, relation_group)
					res[a_alternative_id_] = relation_groups
				}
			}
			return any(res).(map[*End][]*Start)
		case "SOURCE_SPECIFICATION":
			res := make(map[*A_TARGET_SPECIFICATION_1][]*RELATION_GROUP)
			for relation_group := range stage.RELATION_GROUPs {
				if relation_group.SOURCE_SPECIFICATION != nil {
					a_target_specification_1_ := relation_group.SOURCE_SPECIFICATION
					var relation_groups []*RELATION_GROUP
					_, ok := res[a_target_specification_1_]
					if ok {
						relation_groups = res[a_target_specification_1_]
					} else {
						relation_groups = make([]*RELATION_GROUP, 0)
					}
					relation_groups = append(relation_groups, relation_group)
					res[a_target_specification_1_] = relation_groups
				}
			}
			return any(res).(map[*End][]*Start)
		case "SPEC_RELATIONS":
			res := make(map[*A_SPEC_RELATION_REF][]*RELATION_GROUP)
			for relation_group := range stage.RELATION_GROUPs {
				if relation_group.SPEC_RELATIONS != nil {
					a_spec_relation_ref_ := relation_group.SPEC_RELATIONS
					var relation_groups []*RELATION_GROUP
					_, ok := res[a_spec_relation_ref_]
					if ok {
						relation_groups = res[a_spec_relation_ref_]
					} else {
						relation_groups = make([]*RELATION_GROUP, 0)
					}
					relation_groups = append(relation_groups, relation_group)
					res[a_spec_relation_ref_] = relation_groups
				}
			}
			return any(res).(map[*End][]*Start)
		case "TARGET_SPECIFICATION":
			res := make(map[*A_TARGET_SPECIFICATION_1][]*RELATION_GROUP)
			for relation_group := range stage.RELATION_GROUPs {
				if relation_group.TARGET_SPECIFICATION != nil {
					a_target_specification_1_ := relation_group.TARGET_SPECIFICATION
					var relation_groups []*RELATION_GROUP
					_, ok := res[a_target_specification_1_]
					if ok {
						relation_groups = res[a_target_specification_1_]
					} else {
						relation_groups = make([]*RELATION_GROUP, 0)
					}
					relation_groups = append(relation_groups, relation_group)
					res[a_target_specification_1_] = relation_groups
				}
			}
			return any(res).(map[*End][]*Start)
		case "TYPE":
			res := make(map[*A_RELATION_GROUP_TYPE_REF][]*RELATION_GROUP)
			for relation_group := range stage.RELATION_GROUPs {
				if relation_group.TYPE != nil {
					a_relation_group_type_ref_ := relation_group.TYPE
					var relation_groups []*RELATION_GROUP
					_, ok := res[a_relation_group_type_ref_]
					if ok {
						relation_groups = res[a_relation_group_type_ref_]
					} else {
						relation_groups = make([]*RELATION_GROUP, 0)
					}
					relation_groups = append(relation_groups, relation_group)
					res[a_relation_group_type_ref_] = relation_groups
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of RELATION_GROUP_TYPE
	case RELATION_GROUP_TYPE:
		switch fieldname {
		// insertion point for per direct association field
		case "ALTERNATIVE_ID":
			res := make(map[*A_ALTERNATIVE_ID][]*RELATION_GROUP_TYPE)
			for relation_group_type := range stage.RELATION_GROUP_TYPEs {
				if relation_group_type.ALTERNATIVE_ID != nil {
					a_alternative_id_ := relation_group_type.ALTERNATIVE_ID
					var relation_group_types []*RELATION_GROUP_TYPE
					_, ok := res[a_alternative_id_]
					if ok {
						relation_group_types = res[a_alternative_id_]
					} else {
						relation_group_types = make([]*RELATION_GROUP_TYPE, 0)
					}
					relation_group_types = append(relation_group_types, relation_group_type)
					res[a_alternative_id_] = relation_group_types
				}
			}
			return any(res).(map[*End][]*Start)
		case "SPEC_ATTRIBUTES":
			res := make(map[*A_SPEC_ATTRIBUTES][]*RELATION_GROUP_TYPE)
			for relation_group_type := range stage.RELATION_GROUP_TYPEs {
				if relation_group_type.SPEC_ATTRIBUTES != nil {
					a_spec_attributes_ := relation_group_type.SPEC_ATTRIBUTES
					var relation_group_types []*RELATION_GROUP_TYPE
					_, ok := res[a_spec_attributes_]
					if ok {
						relation_group_types = res[a_spec_attributes_]
					} else {
						relation_group_types = make([]*RELATION_GROUP_TYPE, 0)
					}
					relation_group_types = append(relation_group_types, relation_group_type)
					res[a_spec_attributes_] = relation_group_types
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of REQ_IF
	case REQ_IF:
		switch fieldname {
		// insertion point for per direct association field
		case "THE_HEADER":
			res := make(map[*A_THE_HEADER][]*REQ_IF)
			for req_if := range stage.REQ_IFs {
				if req_if.THE_HEADER != nil {
					a_the_header_ := req_if.THE_HEADER
					var req_ifs []*REQ_IF
					_, ok := res[a_the_header_]
					if ok {
						req_ifs = res[a_the_header_]
					} else {
						req_ifs = make([]*REQ_IF, 0)
					}
					req_ifs = append(req_ifs, req_if)
					res[a_the_header_] = req_ifs
				}
			}
			return any(res).(map[*End][]*Start)
		case "CORE_CONTENT":
			res := make(map[*A_CORE_CONTENT][]*REQ_IF)
			for req_if := range stage.REQ_IFs {
				if req_if.CORE_CONTENT != nil {
					a_core_content_ := req_if.CORE_CONTENT
					var req_ifs []*REQ_IF
					_, ok := res[a_core_content_]
					if ok {
						req_ifs = res[a_core_content_]
					} else {
						req_ifs = make([]*REQ_IF, 0)
					}
					req_ifs = append(req_ifs, req_if)
					res[a_core_content_] = req_ifs
				}
			}
			return any(res).(map[*End][]*Start)
		case "TOOL_EXTENSIONS":
			res := make(map[*A_TOOL_EXTENSIONS][]*REQ_IF)
			for req_if := range stage.REQ_IFs {
				if req_if.TOOL_EXTENSIONS != nil {
					a_tool_extensions_ := req_if.TOOL_EXTENSIONS
					var req_ifs []*REQ_IF
					_, ok := res[a_tool_extensions_]
					if ok {
						req_ifs = res[a_tool_extensions_]
					} else {
						req_ifs = make([]*REQ_IF, 0)
					}
					req_ifs = append(req_ifs, req_if)
					res[a_tool_extensions_] = req_ifs
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of REQ_IF_CONTENT
	case REQ_IF_CONTENT:
		switch fieldname {
		// insertion point for per direct association field
		case "DATATYPES":
			res := make(map[*A_DATATYPES][]*REQ_IF_CONTENT)
			for req_if_content := range stage.REQ_IF_CONTENTs {
				if req_if_content.DATATYPES != nil {
					a_datatypes_ := req_if_content.DATATYPES
					var req_if_contents []*REQ_IF_CONTENT
					_, ok := res[a_datatypes_]
					if ok {
						req_if_contents = res[a_datatypes_]
					} else {
						req_if_contents = make([]*REQ_IF_CONTENT, 0)
					}
					req_if_contents = append(req_if_contents, req_if_content)
					res[a_datatypes_] = req_if_contents
				}
			}
			return any(res).(map[*End][]*Start)
		case "SPEC_TYPES":
			res := make(map[*A_SPEC_TYPES][]*REQ_IF_CONTENT)
			for req_if_content := range stage.REQ_IF_CONTENTs {
				if req_if_content.SPEC_TYPES != nil {
					a_spec_types_ := req_if_content.SPEC_TYPES
					var req_if_contents []*REQ_IF_CONTENT
					_, ok := res[a_spec_types_]
					if ok {
						req_if_contents = res[a_spec_types_]
					} else {
						req_if_contents = make([]*REQ_IF_CONTENT, 0)
					}
					req_if_contents = append(req_if_contents, req_if_content)
					res[a_spec_types_] = req_if_contents
				}
			}
			return any(res).(map[*End][]*Start)
		case "SPEC_OBJECTS":
			res := make(map[*A_SPEC_OBJECTS][]*REQ_IF_CONTENT)
			for req_if_content := range stage.REQ_IF_CONTENTs {
				if req_if_content.SPEC_OBJECTS != nil {
					a_spec_objects_ := req_if_content.SPEC_OBJECTS
					var req_if_contents []*REQ_IF_CONTENT
					_, ok := res[a_spec_objects_]
					if ok {
						req_if_contents = res[a_spec_objects_]
					} else {
						req_if_contents = make([]*REQ_IF_CONTENT, 0)
					}
					req_if_contents = append(req_if_contents, req_if_content)
					res[a_spec_objects_] = req_if_contents
				}
			}
			return any(res).(map[*End][]*Start)
		case "SPEC_RELATIONS":
			res := make(map[*A_SPEC_RELATIONS][]*REQ_IF_CONTENT)
			for req_if_content := range stage.REQ_IF_CONTENTs {
				if req_if_content.SPEC_RELATIONS != nil {
					a_spec_relations_ := req_if_content.SPEC_RELATIONS
					var req_if_contents []*REQ_IF_CONTENT
					_, ok := res[a_spec_relations_]
					if ok {
						req_if_contents = res[a_spec_relations_]
					} else {
						req_if_contents = make([]*REQ_IF_CONTENT, 0)
					}
					req_if_contents = append(req_if_contents, req_if_content)
					res[a_spec_relations_] = req_if_contents
				}
			}
			return any(res).(map[*End][]*Start)
		case "SPECIFICATIONS":
			res := make(map[*A_SPECIFICATIONS][]*REQ_IF_CONTENT)
			for req_if_content := range stage.REQ_IF_CONTENTs {
				if req_if_content.SPECIFICATIONS != nil {
					a_specifications_ := req_if_content.SPECIFICATIONS
					var req_if_contents []*REQ_IF_CONTENT
					_, ok := res[a_specifications_]
					if ok {
						req_if_contents = res[a_specifications_]
					} else {
						req_if_contents = make([]*REQ_IF_CONTENT, 0)
					}
					req_if_contents = append(req_if_contents, req_if_content)
					res[a_specifications_] = req_if_contents
				}
			}
			return any(res).(map[*End][]*Start)
		case "SPEC_RELATION_GROUPS":
			res := make(map[*A_SPEC_RELATION_GROUPS][]*REQ_IF_CONTENT)
			for req_if_content := range stage.REQ_IF_CONTENTs {
				if req_if_content.SPEC_RELATION_GROUPS != nil {
					a_spec_relation_groups_ := req_if_content.SPEC_RELATION_GROUPS
					var req_if_contents []*REQ_IF_CONTENT
					_, ok := res[a_spec_relation_groups_]
					if ok {
						req_if_contents = res[a_spec_relation_groups_]
					} else {
						req_if_contents = make([]*REQ_IF_CONTENT, 0)
					}
					req_if_contents = append(req_if_contents, req_if_content)
					res[a_spec_relation_groups_] = req_if_contents
				}
			}
			return any(res).(map[*End][]*Start)
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
		case "ALTERNATIVE_ID":
			res := make(map[*A_ALTERNATIVE_ID][]*SPECIFICATION)
			for specification := range stage.SPECIFICATIONs {
				if specification.ALTERNATIVE_ID != nil {
					a_alternative_id_ := specification.ALTERNATIVE_ID
					var specifications []*SPECIFICATION
					_, ok := res[a_alternative_id_]
					if ok {
						specifications = res[a_alternative_id_]
					} else {
						specifications = make([]*SPECIFICATION, 0)
					}
					specifications = append(specifications, specification)
					res[a_alternative_id_] = specifications
				}
			}
			return any(res).(map[*End][]*Start)
		case "CHILDREN":
			res := make(map[*A_CHILDREN][]*SPECIFICATION)
			for specification := range stage.SPECIFICATIONs {
				if specification.CHILDREN != nil {
					a_children_ := specification.CHILDREN
					var specifications []*SPECIFICATION
					_, ok := res[a_children_]
					if ok {
						specifications = res[a_children_]
					} else {
						specifications = make([]*SPECIFICATION, 0)
					}
					specifications = append(specifications, specification)
					res[a_children_] = specifications
				}
			}
			return any(res).(map[*End][]*Start)
		case "VALUES":
			res := make(map[*A_ATTRIBUTE_VALUE_XHTML][]*SPECIFICATION)
			for specification := range stage.SPECIFICATIONs {
				if specification.VALUES != nil {
					a_attribute_value_xhtml_ := specification.VALUES
					var specifications []*SPECIFICATION
					_, ok := res[a_attribute_value_xhtml_]
					if ok {
						specifications = res[a_attribute_value_xhtml_]
					} else {
						specifications = make([]*SPECIFICATION, 0)
					}
					specifications = append(specifications, specification)
					res[a_attribute_value_xhtml_] = specifications
				}
			}
			return any(res).(map[*End][]*Start)
		case "TYPE":
			res := make(map[*A_SPECIFICATION_TYPE_REF][]*SPECIFICATION)
			for specification := range stage.SPECIFICATIONs {
				if specification.TYPE != nil {
					a_specification_type_ref_ := specification.TYPE
					var specifications []*SPECIFICATION
					_, ok := res[a_specification_type_ref_]
					if ok {
						specifications = res[a_specification_type_ref_]
					} else {
						specifications = make([]*SPECIFICATION, 0)
					}
					specifications = append(specifications, specification)
					res[a_specification_type_ref_] = specifications
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of SPECIFICATION_TYPE
	case SPECIFICATION_TYPE:
		switch fieldname {
		// insertion point for per direct association field
		case "ALTERNATIVE_ID":
			res := make(map[*A_ALTERNATIVE_ID][]*SPECIFICATION_TYPE)
			for specification_type := range stage.SPECIFICATION_TYPEs {
				if specification_type.ALTERNATIVE_ID != nil {
					a_alternative_id_ := specification_type.ALTERNATIVE_ID
					var specification_types []*SPECIFICATION_TYPE
					_, ok := res[a_alternative_id_]
					if ok {
						specification_types = res[a_alternative_id_]
					} else {
						specification_types = make([]*SPECIFICATION_TYPE, 0)
					}
					specification_types = append(specification_types, specification_type)
					res[a_alternative_id_] = specification_types
				}
			}
			return any(res).(map[*End][]*Start)
		case "SPEC_ATTRIBUTES":
			res := make(map[*A_SPEC_ATTRIBUTES][]*SPECIFICATION_TYPE)
			for specification_type := range stage.SPECIFICATION_TYPEs {
				if specification_type.SPEC_ATTRIBUTES != nil {
					a_spec_attributes_ := specification_type.SPEC_ATTRIBUTES
					var specification_types []*SPECIFICATION_TYPE
					_, ok := res[a_spec_attributes_]
					if ok {
						specification_types = res[a_spec_attributes_]
					} else {
						specification_types = make([]*SPECIFICATION_TYPE, 0)
					}
					specification_types = append(specification_types, specification_type)
					res[a_spec_attributes_] = specification_types
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of SPEC_HIERARCHY
	case SPEC_HIERARCHY:
		switch fieldname {
		// insertion point for per direct association field
		case "ALTERNATIVE_ID":
			res := make(map[*A_ALTERNATIVE_ID][]*SPEC_HIERARCHY)
			for spec_hierarchy := range stage.SPEC_HIERARCHYs {
				if spec_hierarchy.ALTERNATIVE_ID != nil {
					a_alternative_id_ := spec_hierarchy.ALTERNATIVE_ID
					var spec_hierarchys []*SPEC_HIERARCHY
					_, ok := res[a_alternative_id_]
					if ok {
						spec_hierarchys = res[a_alternative_id_]
					} else {
						spec_hierarchys = make([]*SPEC_HIERARCHY, 0)
					}
					spec_hierarchys = append(spec_hierarchys, spec_hierarchy)
					res[a_alternative_id_] = spec_hierarchys
				}
			}
			return any(res).(map[*End][]*Start)
		case "CHILDREN":
			res := make(map[*A_CHILDREN][]*SPEC_HIERARCHY)
			for spec_hierarchy := range stage.SPEC_HIERARCHYs {
				if spec_hierarchy.CHILDREN != nil {
					a_children_ := spec_hierarchy.CHILDREN
					var spec_hierarchys []*SPEC_HIERARCHY
					_, ok := res[a_children_]
					if ok {
						spec_hierarchys = res[a_children_]
					} else {
						spec_hierarchys = make([]*SPEC_HIERARCHY, 0)
					}
					spec_hierarchys = append(spec_hierarchys, spec_hierarchy)
					res[a_children_] = spec_hierarchys
				}
			}
			return any(res).(map[*End][]*Start)
		case "EDITABLE_ATTS":
			res := make(map[*A_EDITABLE_ATTS][]*SPEC_HIERARCHY)
			for spec_hierarchy := range stage.SPEC_HIERARCHYs {
				if spec_hierarchy.EDITABLE_ATTS != nil {
					a_editable_atts_ := spec_hierarchy.EDITABLE_ATTS
					var spec_hierarchys []*SPEC_HIERARCHY
					_, ok := res[a_editable_atts_]
					if ok {
						spec_hierarchys = res[a_editable_atts_]
					} else {
						spec_hierarchys = make([]*SPEC_HIERARCHY, 0)
					}
					spec_hierarchys = append(spec_hierarchys, spec_hierarchy)
					res[a_editable_atts_] = spec_hierarchys
				}
			}
			return any(res).(map[*End][]*Start)
		case "OBJECT":
			res := make(map[*A_OBJECT][]*SPEC_HIERARCHY)
			for spec_hierarchy := range stage.SPEC_HIERARCHYs {
				if spec_hierarchy.OBJECT != nil {
					a_object_ := spec_hierarchy.OBJECT
					var spec_hierarchys []*SPEC_HIERARCHY
					_, ok := res[a_object_]
					if ok {
						spec_hierarchys = res[a_object_]
					} else {
						spec_hierarchys = make([]*SPEC_HIERARCHY, 0)
					}
					spec_hierarchys = append(spec_hierarchys, spec_hierarchy)
					res[a_object_] = spec_hierarchys
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of SPEC_OBJECT
	case SPEC_OBJECT:
		switch fieldname {
		// insertion point for per direct association field
		case "ALTERNATIVE_ID":
			res := make(map[*A_ALTERNATIVE_ID][]*SPEC_OBJECT)
			for spec_object := range stage.SPEC_OBJECTs {
				if spec_object.ALTERNATIVE_ID != nil {
					a_alternative_id_ := spec_object.ALTERNATIVE_ID
					var spec_objects []*SPEC_OBJECT
					_, ok := res[a_alternative_id_]
					if ok {
						spec_objects = res[a_alternative_id_]
					} else {
						spec_objects = make([]*SPEC_OBJECT, 0)
					}
					spec_objects = append(spec_objects, spec_object)
					res[a_alternative_id_] = spec_objects
				}
			}
			return any(res).(map[*End][]*Start)
		case "VALUES":
			res := make(map[*A_ATTRIBUTE_VALUE_XHTML][]*SPEC_OBJECT)
			for spec_object := range stage.SPEC_OBJECTs {
				if spec_object.VALUES != nil {
					a_attribute_value_xhtml_ := spec_object.VALUES
					var spec_objects []*SPEC_OBJECT
					_, ok := res[a_attribute_value_xhtml_]
					if ok {
						spec_objects = res[a_attribute_value_xhtml_]
					} else {
						spec_objects = make([]*SPEC_OBJECT, 0)
					}
					spec_objects = append(spec_objects, spec_object)
					res[a_attribute_value_xhtml_] = spec_objects
				}
			}
			return any(res).(map[*End][]*Start)
		case "TYPE":
			res := make(map[*A_SPEC_OBJECT_TYPE_REF][]*SPEC_OBJECT)
			for spec_object := range stage.SPEC_OBJECTs {
				if spec_object.TYPE != nil {
					a_spec_object_type_ref_ := spec_object.TYPE
					var spec_objects []*SPEC_OBJECT
					_, ok := res[a_spec_object_type_ref_]
					if ok {
						spec_objects = res[a_spec_object_type_ref_]
					} else {
						spec_objects = make([]*SPEC_OBJECT, 0)
					}
					spec_objects = append(spec_objects, spec_object)
					res[a_spec_object_type_ref_] = spec_objects
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of SPEC_OBJECT_TYPE
	case SPEC_OBJECT_TYPE:
		switch fieldname {
		// insertion point for per direct association field
		case "ALTERNATIVE_ID":
			res := make(map[*A_ALTERNATIVE_ID][]*SPEC_OBJECT_TYPE)
			for spec_object_type := range stage.SPEC_OBJECT_TYPEs {
				if spec_object_type.ALTERNATIVE_ID != nil {
					a_alternative_id_ := spec_object_type.ALTERNATIVE_ID
					var spec_object_types []*SPEC_OBJECT_TYPE
					_, ok := res[a_alternative_id_]
					if ok {
						spec_object_types = res[a_alternative_id_]
					} else {
						spec_object_types = make([]*SPEC_OBJECT_TYPE, 0)
					}
					spec_object_types = append(spec_object_types, spec_object_type)
					res[a_alternative_id_] = spec_object_types
				}
			}
			return any(res).(map[*End][]*Start)
		case "SPEC_ATTRIBUTES":
			res := make(map[*A_SPEC_ATTRIBUTES][]*SPEC_OBJECT_TYPE)
			for spec_object_type := range stage.SPEC_OBJECT_TYPEs {
				if spec_object_type.SPEC_ATTRIBUTES != nil {
					a_spec_attributes_ := spec_object_type.SPEC_ATTRIBUTES
					var spec_object_types []*SPEC_OBJECT_TYPE
					_, ok := res[a_spec_attributes_]
					if ok {
						spec_object_types = res[a_spec_attributes_]
					} else {
						spec_object_types = make([]*SPEC_OBJECT_TYPE, 0)
					}
					spec_object_types = append(spec_object_types, spec_object_type)
					res[a_spec_attributes_] = spec_object_types
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of SPEC_RELATION
	case SPEC_RELATION:
		switch fieldname {
		// insertion point for per direct association field
		case "ALTERNATIVE_ID":
			res := make(map[*A_ALTERNATIVE_ID][]*SPEC_RELATION)
			for spec_relation := range stage.SPEC_RELATIONs {
				if spec_relation.ALTERNATIVE_ID != nil {
					a_alternative_id_ := spec_relation.ALTERNATIVE_ID
					var spec_relations []*SPEC_RELATION
					_, ok := res[a_alternative_id_]
					if ok {
						spec_relations = res[a_alternative_id_]
					} else {
						spec_relations = make([]*SPEC_RELATION, 0)
					}
					spec_relations = append(spec_relations, spec_relation)
					res[a_alternative_id_] = spec_relations
				}
			}
			return any(res).(map[*End][]*Start)
		case "VALUES":
			res := make(map[*A_ATTRIBUTE_VALUE_XHTML][]*SPEC_RELATION)
			for spec_relation := range stage.SPEC_RELATIONs {
				if spec_relation.VALUES != nil {
					a_attribute_value_xhtml_ := spec_relation.VALUES
					var spec_relations []*SPEC_RELATION
					_, ok := res[a_attribute_value_xhtml_]
					if ok {
						spec_relations = res[a_attribute_value_xhtml_]
					} else {
						spec_relations = make([]*SPEC_RELATION, 0)
					}
					spec_relations = append(spec_relations, spec_relation)
					res[a_attribute_value_xhtml_] = spec_relations
				}
			}
			return any(res).(map[*End][]*Start)
		case "SOURCE":
			res := make(map[*A_SOURCE_1][]*SPEC_RELATION)
			for spec_relation := range stage.SPEC_RELATIONs {
				if spec_relation.SOURCE != nil {
					a_source_1_ := spec_relation.SOURCE
					var spec_relations []*SPEC_RELATION
					_, ok := res[a_source_1_]
					if ok {
						spec_relations = res[a_source_1_]
					} else {
						spec_relations = make([]*SPEC_RELATION, 0)
					}
					spec_relations = append(spec_relations, spec_relation)
					res[a_source_1_] = spec_relations
				}
			}
			return any(res).(map[*End][]*Start)
		case "TARGET":
			res := make(map[*A_SOURCE_1][]*SPEC_RELATION)
			for spec_relation := range stage.SPEC_RELATIONs {
				if spec_relation.TARGET != nil {
					a_source_1_ := spec_relation.TARGET
					var spec_relations []*SPEC_RELATION
					_, ok := res[a_source_1_]
					if ok {
						spec_relations = res[a_source_1_]
					} else {
						spec_relations = make([]*SPEC_RELATION, 0)
					}
					spec_relations = append(spec_relations, spec_relation)
					res[a_source_1_] = spec_relations
				}
			}
			return any(res).(map[*End][]*Start)
		case "TYPE":
			res := make(map[*A_SPEC_RELATION_TYPE_REF][]*SPEC_RELATION)
			for spec_relation := range stage.SPEC_RELATIONs {
				if spec_relation.TYPE != nil {
					a_spec_relation_type_ref_ := spec_relation.TYPE
					var spec_relations []*SPEC_RELATION
					_, ok := res[a_spec_relation_type_ref_]
					if ok {
						spec_relations = res[a_spec_relation_type_ref_]
					} else {
						spec_relations = make([]*SPEC_RELATION, 0)
					}
					spec_relations = append(spec_relations, spec_relation)
					res[a_spec_relation_type_ref_] = spec_relations
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of SPEC_RELATION_TYPE
	case SPEC_RELATION_TYPE:
		switch fieldname {
		// insertion point for per direct association field
		case "ALTERNATIVE_ID":
			res := make(map[*A_ALTERNATIVE_ID][]*SPEC_RELATION_TYPE)
			for spec_relation_type := range stage.SPEC_RELATION_TYPEs {
				if spec_relation_type.ALTERNATIVE_ID != nil {
					a_alternative_id_ := spec_relation_type.ALTERNATIVE_ID
					var spec_relation_types []*SPEC_RELATION_TYPE
					_, ok := res[a_alternative_id_]
					if ok {
						spec_relation_types = res[a_alternative_id_]
					} else {
						spec_relation_types = make([]*SPEC_RELATION_TYPE, 0)
					}
					spec_relation_types = append(spec_relation_types, spec_relation_type)
					res[a_alternative_id_] = spec_relation_types
				}
			}
			return any(res).(map[*End][]*Start)
		case "SPEC_ATTRIBUTES":
			res := make(map[*A_SPEC_ATTRIBUTES][]*SPEC_RELATION_TYPE)
			for spec_relation_type := range stage.SPEC_RELATION_TYPEs {
				if spec_relation_type.SPEC_ATTRIBUTES != nil {
					a_spec_attributes_ := spec_relation_type.SPEC_ATTRIBUTES
					var spec_relation_types []*SPEC_RELATION_TYPE
					_, ok := res[a_spec_attributes_]
					if ok {
						spec_relation_types = res[a_spec_attributes_]
					} else {
						spec_relation_types = make([]*SPEC_RELATION_TYPE, 0)
					}
					spec_relation_types = append(spec_relation_types, spec_relation_type)
					res[a_spec_attributes_] = spec_relation_types
				}
			}
			return any(res).(map[*End][]*Start)
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
		}
	// reverse maps of direct associations of A_ALTERNATIVE_ID
	case A_ALTERNATIVE_ID:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_ATTRIBUTE_DEFINITION_BOOLEAN_REF
	case A_ATTRIBUTE_DEFINITION_BOOLEAN_REF:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_ATTRIBUTE_DEFINITION_DATE_REF
	case A_ATTRIBUTE_DEFINITION_DATE_REF:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_ATTRIBUTE_DEFINITION_ENUMERATION_REF
	case A_ATTRIBUTE_DEFINITION_ENUMERATION_REF:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_ATTRIBUTE_DEFINITION_INTEGER_REF
	case A_ATTRIBUTE_DEFINITION_INTEGER_REF:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_ATTRIBUTE_DEFINITION_REAL_REF
	case A_ATTRIBUTE_DEFINITION_REAL_REF:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_ATTRIBUTE_DEFINITION_STRING_REF
	case A_ATTRIBUTE_DEFINITION_STRING_REF:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_ATTRIBUTE_DEFINITION_XHTML_REF
	case A_ATTRIBUTE_DEFINITION_XHTML_REF:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_ATTRIBUTE_VALUE_BOOLEAN
	case A_ATTRIBUTE_VALUE_BOOLEAN:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_ATTRIBUTE_VALUE_DATE
	case A_ATTRIBUTE_VALUE_DATE:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_ATTRIBUTE_VALUE_ENUMERATION
	case A_ATTRIBUTE_VALUE_ENUMERATION:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_ATTRIBUTE_VALUE_INTEGER
	case A_ATTRIBUTE_VALUE_INTEGER:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_ATTRIBUTE_VALUE_REAL
	case A_ATTRIBUTE_VALUE_REAL:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_ATTRIBUTE_VALUE_STRING
	case A_ATTRIBUTE_VALUE_STRING:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_ATTRIBUTE_VALUE_XHTML
	case A_ATTRIBUTE_VALUE_XHTML:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_ATTRIBUTE_VALUE_XHTML_1
	case A_ATTRIBUTE_VALUE_XHTML_1:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_CHILDREN
	case A_CHILDREN:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_CORE_CONTENT
	case A_CORE_CONTENT:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_DATATYPES
	case A_DATATYPES:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_DATATYPE_DEFINITION_BOOLEAN_REF
	case A_DATATYPE_DEFINITION_BOOLEAN_REF:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_DATATYPE_DEFINITION_DATE_REF
	case A_DATATYPE_DEFINITION_DATE_REF:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_DATATYPE_DEFINITION_ENUMERATION_REF
	case A_DATATYPE_DEFINITION_ENUMERATION_REF:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_DATATYPE_DEFINITION_INTEGER_REF
	case A_DATATYPE_DEFINITION_INTEGER_REF:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_DATATYPE_DEFINITION_REAL_REF
	case A_DATATYPE_DEFINITION_REAL_REF:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_DATATYPE_DEFINITION_STRING_REF
	case A_DATATYPE_DEFINITION_STRING_REF:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_DATATYPE_DEFINITION_XHTML_REF
	case A_DATATYPE_DEFINITION_XHTML_REF:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_EDITABLE_ATTS
	case A_EDITABLE_ATTS:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_ENUM_VALUE_REF
	case A_ENUM_VALUE_REF:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_OBJECT
	case A_OBJECT:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_PROPERTIES
	case A_PROPERTIES:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_RELATION_GROUP_TYPE_REF
	case A_RELATION_GROUP_TYPE_REF:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_SOURCE_1
	case A_SOURCE_1:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_SPECIFICATIONS
	case A_SPECIFICATIONS:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_SPECIFICATION_TYPE_REF
	case A_SPECIFICATION_TYPE_REF:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_SPECIFIED_VALUES
	case A_SPECIFIED_VALUES:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_SPEC_ATTRIBUTES
	case A_SPEC_ATTRIBUTES:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_SPEC_OBJECTS
	case A_SPEC_OBJECTS:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_SPEC_OBJECT_TYPE_REF
	case A_SPEC_OBJECT_TYPE_REF:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_SPEC_RELATIONS
	case A_SPEC_RELATIONS:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_SPEC_RELATION_GROUPS
	case A_SPEC_RELATION_GROUPS:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_SPEC_RELATION_REF
	case A_SPEC_RELATION_REF:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_SPEC_RELATION_TYPE_REF
	case A_SPEC_RELATION_TYPE_REF:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_SPEC_TYPES
	case A_SPEC_TYPES:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_TARGET_SPECIFICATION_1
	case A_TARGET_SPECIFICATION_1:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_THE_HEADER
	case A_THE_HEADER:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of A_TOOL_EXTENSIONS
	case A_TOOL_EXTENSIONS:
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
	case A_ALTERNATIVE_ID:
		res = "A_ALTERNATIVE_ID"
	case A_ATTRIBUTE_DEFINITION_BOOLEAN_REF:
		res = "A_ATTRIBUTE_DEFINITION_BOOLEAN_REF"
	case A_ATTRIBUTE_DEFINITION_DATE_REF:
		res = "A_ATTRIBUTE_DEFINITION_DATE_REF"
	case A_ATTRIBUTE_DEFINITION_ENUMERATION_REF:
		res = "A_ATTRIBUTE_DEFINITION_ENUMERATION_REF"
	case A_ATTRIBUTE_DEFINITION_INTEGER_REF:
		res = "A_ATTRIBUTE_DEFINITION_INTEGER_REF"
	case A_ATTRIBUTE_DEFINITION_REAL_REF:
		res = "A_ATTRIBUTE_DEFINITION_REAL_REF"
	case A_ATTRIBUTE_DEFINITION_STRING_REF:
		res = "A_ATTRIBUTE_DEFINITION_STRING_REF"
	case A_ATTRIBUTE_DEFINITION_XHTML_REF:
		res = "A_ATTRIBUTE_DEFINITION_XHTML_REF"
	case A_ATTRIBUTE_VALUE_BOOLEAN:
		res = "A_ATTRIBUTE_VALUE_BOOLEAN"
	case A_ATTRIBUTE_VALUE_DATE:
		res = "A_ATTRIBUTE_VALUE_DATE"
	case A_ATTRIBUTE_VALUE_ENUMERATION:
		res = "A_ATTRIBUTE_VALUE_ENUMERATION"
	case A_ATTRIBUTE_VALUE_INTEGER:
		res = "A_ATTRIBUTE_VALUE_INTEGER"
	case A_ATTRIBUTE_VALUE_REAL:
		res = "A_ATTRIBUTE_VALUE_REAL"
	case A_ATTRIBUTE_VALUE_STRING:
		res = "A_ATTRIBUTE_VALUE_STRING"
	case A_ATTRIBUTE_VALUE_XHTML:
		res = "A_ATTRIBUTE_VALUE_XHTML"
	case A_ATTRIBUTE_VALUE_XHTML_1:
		res = "A_ATTRIBUTE_VALUE_XHTML_1"
	case A_CHILDREN:
		res = "A_CHILDREN"
	case A_CORE_CONTENT:
		res = "A_CORE_CONTENT"
	case A_DATATYPES:
		res = "A_DATATYPES"
	case A_DATATYPE_DEFINITION_BOOLEAN_REF:
		res = "A_DATATYPE_DEFINITION_BOOLEAN_REF"
	case A_DATATYPE_DEFINITION_DATE_REF:
		res = "A_DATATYPE_DEFINITION_DATE_REF"
	case A_DATATYPE_DEFINITION_ENUMERATION_REF:
		res = "A_DATATYPE_DEFINITION_ENUMERATION_REF"
	case A_DATATYPE_DEFINITION_INTEGER_REF:
		res = "A_DATATYPE_DEFINITION_INTEGER_REF"
	case A_DATATYPE_DEFINITION_REAL_REF:
		res = "A_DATATYPE_DEFINITION_REAL_REF"
	case A_DATATYPE_DEFINITION_STRING_REF:
		res = "A_DATATYPE_DEFINITION_STRING_REF"
	case A_DATATYPE_DEFINITION_XHTML_REF:
		res = "A_DATATYPE_DEFINITION_XHTML_REF"
	case A_EDITABLE_ATTS:
		res = "A_EDITABLE_ATTS"
	case A_ENUM_VALUE_REF:
		res = "A_ENUM_VALUE_REF"
	case A_OBJECT:
		res = "A_OBJECT"
	case A_PROPERTIES:
		res = "A_PROPERTIES"
	case A_RELATION_GROUP_TYPE_REF:
		res = "A_RELATION_GROUP_TYPE_REF"
	case A_SOURCE_1:
		res = "A_SOURCE_1"
	case A_SPECIFICATIONS:
		res = "A_SPECIFICATIONS"
	case A_SPECIFICATION_TYPE_REF:
		res = "A_SPECIFICATION_TYPE_REF"
	case A_SPECIFIED_VALUES:
		res = "A_SPECIFIED_VALUES"
	case A_SPEC_ATTRIBUTES:
		res = "A_SPEC_ATTRIBUTES"
	case A_SPEC_OBJECTS:
		res = "A_SPEC_OBJECTS"
	case A_SPEC_OBJECT_TYPE_REF:
		res = "A_SPEC_OBJECT_TYPE_REF"
	case A_SPEC_RELATIONS:
		res = "A_SPEC_RELATIONS"
	case A_SPEC_RELATION_GROUPS:
		res = "A_SPEC_RELATION_GROUPS"
	case A_SPEC_RELATION_REF:
		res = "A_SPEC_RELATION_REF"
	case A_SPEC_RELATION_TYPE_REF:
		res = "A_SPEC_RELATION_TYPE_REF"
	case A_SPEC_TYPES:
		res = "A_SPEC_TYPES"
	case A_TARGET_SPECIFICATION_1:
		res = "A_TARGET_SPECIFICATION_1"
	case A_THE_HEADER:
		res = "A_THE_HEADER"
	case A_TOOL_EXTENSIONS:
		res = "A_TOOL_EXTENSIONS"
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
	case *A_ALTERNATIVE_ID:
		res = "A_ALTERNATIVE_ID"
	case *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF:
		res = "A_ATTRIBUTE_DEFINITION_BOOLEAN_REF"
	case *A_ATTRIBUTE_DEFINITION_DATE_REF:
		res = "A_ATTRIBUTE_DEFINITION_DATE_REF"
	case *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF:
		res = "A_ATTRIBUTE_DEFINITION_ENUMERATION_REF"
	case *A_ATTRIBUTE_DEFINITION_INTEGER_REF:
		res = "A_ATTRIBUTE_DEFINITION_INTEGER_REF"
	case *A_ATTRIBUTE_DEFINITION_REAL_REF:
		res = "A_ATTRIBUTE_DEFINITION_REAL_REF"
	case *A_ATTRIBUTE_DEFINITION_STRING_REF:
		res = "A_ATTRIBUTE_DEFINITION_STRING_REF"
	case *A_ATTRIBUTE_DEFINITION_XHTML_REF:
		res = "A_ATTRIBUTE_DEFINITION_XHTML_REF"
	case *A_ATTRIBUTE_VALUE_BOOLEAN:
		res = "A_ATTRIBUTE_VALUE_BOOLEAN"
	case *A_ATTRIBUTE_VALUE_DATE:
		res = "A_ATTRIBUTE_VALUE_DATE"
	case *A_ATTRIBUTE_VALUE_ENUMERATION:
		res = "A_ATTRIBUTE_VALUE_ENUMERATION"
	case *A_ATTRIBUTE_VALUE_INTEGER:
		res = "A_ATTRIBUTE_VALUE_INTEGER"
	case *A_ATTRIBUTE_VALUE_REAL:
		res = "A_ATTRIBUTE_VALUE_REAL"
	case *A_ATTRIBUTE_VALUE_STRING:
		res = "A_ATTRIBUTE_VALUE_STRING"
	case *A_ATTRIBUTE_VALUE_XHTML:
		res = "A_ATTRIBUTE_VALUE_XHTML"
	case *A_ATTRIBUTE_VALUE_XHTML_1:
		res = "A_ATTRIBUTE_VALUE_XHTML_1"
	case *A_CHILDREN:
		res = "A_CHILDREN"
	case *A_CORE_CONTENT:
		res = "A_CORE_CONTENT"
	case *A_DATATYPES:
		res = "A_DATATYPES"
	case *A_DATATYPE_DEFINITION_BOOLEAN_REF:
		res = "A_DATATYPE_DEFINITION_BOOLEAN_REF"
	case *A_DATATYPE_DEFINITION_DATE_REF:
		res = "A_DATATYPE_DEFINITION_DATE_REF"
	case *A_DATATYPE_DEFINITION_ENUMERATION_REF:
		res = "A_DATATYPE_DEFINITION_ENUMERATION_REF"
	case *A_DATATYPE_DEFINITION_INTEGER_REF:
		res = "A_DATATYPE_DEFINITION_INTEGER_REF"
	case *A_DATATYPE_DEFINITION_REAL_REF:
		res = "A_DATATYPE_DEFINITION_REAL_REF"
	case *A_DATATYPE_DEFINITION_STRING_REF:
		res = "A_DATATYPE_DEFINITION_STRING_REF"
	case *A_DATATYPE_DEFINITION_XHTML_REF:
		res = "A_DATATYPE_DEFINITION_XHTML_REF"
	case *A_EDITABLE_ATTS:
		res = "A_EDITABLE_ATTS"
	case *A_ENUM_VALUE_REF:
		res = "A_ENUM_VALUE_REF"
	case *A_OBJECT:
		res = "A_OBJECT"
	case *A_PROPERTIES:
		res = "A_PROPERTIES"
	case *A_RELATION_GROUP_TYPE_REF:
		res = "A_RELATION_GROUP_TYPE_REF"
	case *A_SOURCE_1:
		res = "A_SOURCE_1"
	case *A_SPECIFICATIONS:
		res = "A_SPECIFICATIONS"
	case *A_SPECIFICATION_TYPE_REF:
		res = "A_SPECIFICATION_TYPE_REF"
	case *A_SPECIFIED_VALUES:
		res = "A_SPECIFIED_VALUES"
	case *A_SPEC_ATTRIBUTES:
		res = "A_SPEC_ATTRIBUTES"
	case *A_SPEC_OBJECTS:
		res = "A_SPEC_OBJECTS"
	case *A_SPEC_OBJECT_TYPE_REF:
		res = "A_SPEC_OBJECT_TYPE_REF"
	case *A_SPEC_RELATIONS:
		res = "A_SPEC_RELATIONS"
	case *A_SPEC_RELATION_GROUPS:
		res = "A_SPEC_RELATION_GROUPS"
	case *A_SPEC_RELATION_REF:
		res = "A_SPEC_RELATION_REF"
	case *A_SPEC_RELATION_TYPE_REF:
		res = "A_SPEC_RELATION_TYPE_REF"
	case *A_SPEC_TYPES:
		res = "A_SPEC_TYPES"
	case *A_TARGET_SPECIFICATION_1:
		res = "A_TARGET_SPECIFICATION_1"
	case *A_THE_HEADER:
		res = "A_THE_HEADER"
	case *A_TOOL_EXTENSIONS:
		res = "A_TOOL_EXTENSIONS"
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
		res = []string{"Name", "DESC", "IDENTIFIER", "IS_EDITABLE", "LAST_CHANGE", "LONG_NAME", "ALTERNATIVE_ID", "DEFAULT_VALUE", "TYPE"}
	case ATTRIBUTE_DEFINITION_DATE:
		res = []string{"Name", "DESC", "IDENTIFIER", "IS_EDITABLE", "LAST_CHANGE", "LONG_NAME", "ALTERNATIVE_ID", "DEFAULT_VALUE", "TYPE"}
	case ATTRIBUTE_DEFINITION_ENUMERATION:
		res = []string{"Name", "DESC", "IDENTIFIER", "IS_EDITABLE", "LAST_CHANGE", "LONG_NAME", "MULTI_VALUED", "ALTERNATIVE_ID", "DEFAULT_VALUE", "TYPE"}
	case ATTRIBUTE_DEFINITION_INTEGER:
		res = []string{"Name", "DESC", "IDENTIFIER", "IS_EDITABLE", "LAST_CHANGE", "LONG_NAME", "ALTERNATIVE_ID", "DEFAULT_VALUE", "TYPE"}
	case ATTRIBUTE_DEFINITION_REAL:
		res = []string{"Name", "DESC", "IDENTIFIER", "IS_EDITABLE", "LAST_CHANGE", "LONG_NAME", "ALTERNATIVE_ID", "DEFAULT_VALUE", "TYPE"}
	case ATTRIBUTE_DEFINITION_STRING:
		res = []string{"Name", "DESC", "IDENTIFIER", "IS_EDITABLE", "LAST_CHANGE", "LONG_NAME", "ALTERNATIVE_ID", "DEFAULT_VALUE", "TYPE"}
	case ATTRIBUTE_DEFINITION_XHTML:
		res = []string{"Name", "DESC", "IDENTIFIER", "IS_EDITABLE", "LAST_CHANGE", "LONG_NAME", "ALTERNATIVE_ID", "DEFAULT_VALUE", "TYPE"}
	case ATTRIBUTE_VALUE_BOOLEAN:
		res = []string{"Name", "THE_VALUE", "DEFINITION"}
	case ATTRIBUTE_VALUE_DATE:
		res = []string{"Name", "THE_VALUE", "DEFINITION"}
	case ATTRIBUTE_VALUE_ENUMERATION:
		res = []string{"Name", "DEFINITION", "VALUES"}
	case ATTRIBUTE_VALUE_INTEGER:
		res = []string{"Name", "THE_VALUE", "DEFINITION"}
	case ATTRIBUTE_VALUE_REAL:
		res = []string{"Name", "THE_VALUE", "DEFINITION"}
	case ATTRIBUTE_VALUE_STRING:
		res = []string{"Name", "THE_VALUE", "DEFINITION"}
	case ATTRIBUTE_VALUE_XHTML:
		res = []string{"Name", "IS_SIMPLIFIED", "THE_VALUE", "THE_ORIGINAL_VALUE", "DEFINITION"}
	case A_ALTERNATIVE_ID:
		res = []string{"Name", "ALTERNATIVE_ID"}
	case A_ATTRIBUTE_DEFINITION_BOOLEAN_REF:
		res = []string{"Name", "ATTRIBUTE_DEFINITION_BOOLEAN_REF"}
	case A_ATTRIBUTE_DEFINITION_DATE_REF:
		res = []string{"Name", "ATTRIBUTE_DEFINITION_DATE_REF"}
	case A_ATTRIBUTE_DEFINITION_ENUMERATION_REF:
		res = []string{"Name", "ATTRIBUTE_DEFINITION_ENUMERATION_REF"}
	case A_ATTRIBUTE_DEFINITION_INTEGER_REF:
		res = []string{"Name", "ATTRIBUTE_DEFINITION_INTEGER_REF"}
	case A_ATTRIBUTE_DEFINITION_REAL_REF:
		res = []string{"Name", "ATTRIBUTE_DEFINITION_REAL_REF"}
	case A_ATTRIBUTE_DEFINITION_STRING_REF:
		res = []string{"Name", "ATTRIBUTE_DEFINITION_STRING_REF"}
	case A_ATTRIBUTE_DEFINITION_XHTML_REF:
		res = []string{"Name", "ATTRIBUTE_DEFINITION_XHTML_REF"}
	case A_ATTRIBUTE_VALUE_BOOLEAN:
		res = []string{"Name", "ATTRIBUTE_VALUE_BOOLEAN"}
	case A_ATTRIBUTE_VALUE_DATE:
		res = []string{"Name", "ATTRIBUTE_VALUE_DATE"}
	case A_ATTRIBUTE_VALUE_ENUMERATION:
		res = []string{"Name", "ATTRIBUTE_VALUE_ENUMERATION"}
	case A_ATTRIBUTE_VALUE_INTEGER:
		res = []string{"Name", "ATTRIBUTE_VALUE_INTEGER"}
	case A_ATTRIBUTE_VALUE_REAL:
		res = []string{"Name", "ATTRIBUTE_VALUE_REAL"}
	case A_ATTRIBUTE_VALUE_STRING:
		res = []string{"Name", "ATTRIBUTE_VALUE_STRING"}
	case A_ATTRIBUTE_VALUE_XHTML:
		res = []string{"Name", "ATTRIBUTE_VALUE_BOOLEAN", "ATTRIBUTE_VALUE_DATE", "ATTRIBUTE_VALUE_ENUMERATION", "ATTRIBUTE_VALUE_INTEGER", "ATTRIBUTE_VALUE_REAL", "ATTRIBUTE_VALUE_STRING", "ATTRIBUTE_VALUE_XHTML"}
	case A_ATTRIBUTE_VALUE_XHTML_1:
		res = []string{"Name", "ATTRIBUTE_VALUE_XHTML"}
	case A_CHILDREN:
		res = []string{"Name", "SPEC_HIERARCHY"}
	case A_CORE_CONTENT:
		res = []string{"Name", "REQ_IF_CONTENT"}
	case A_DATATYPES:
		res = []string{"Name", "DATATYPE_DEFINITION_BOOLEAN", "DATATYPE_DEFINITION_DATE", "DATATYPE_DEFINITION_ENUMERATION", "DATATYPE_DEFINITION_INTEGER", "DATATYPE_DEFINITION_REAL", "DATATYPE_DEFINITION_STRING", "DATATYPE_DEFINITION_XHTML"}
	case A_DATATYPE_DEFINITION_BOOLEAN_REF:
		res = []string{"Name", "DATATYPE_DEFINITION_BOOLEAN_REF"}
	case A_DATATYPE_DEFINITION_DATE_REF:
		res = []string{"Name", "DATATYPE_DEFINITION_DATE_REF"}
	case A_DATATYPE_DEFINITION_ENUMERATION_REF:
		res = []string{"Name", "DATATYPE_DEFINITION_ENUMERATION_REF"}
	case A_DATATYPE_DEFINITION_INTEGER_REF:
		res = []string{"Name", "DATATYPE_DEFINITION_INTEGER_REF"}
	case A_DATATYPE_DEFINITION_REAL_REF:
		res = []string{"Name", "DATATYPE_DEFINITION_REAL_REF"}
	case A_DATATYPE_DEFINITION_STRING_REF:
		res = []string{"Name", "DATATYPE_DEFINITION_STRING_REF"}
	case A_DATATYPE_DEFINITION_XHTML_REF:
		res = []string{"Name", "DATATYPE_DEFINITION_XHTML_REF"}
	case A_EDITABLE_ATTS:
		res = []string{"Name", "ATTRIBUTE_DEFINITION_BOOLEAN_REF", "ATTRIBUTE_DEFINITION_DATE_REF", "ATTRIBUTE_DEFINITION_ENUMERATION_REF", "ATTRIBUTE_DEFINITION_INTEGER_REF", "ATTRIBUTE_DEFINITION_REAL_REF", "ATTRIBUTE_DEFINITION_STRING_REF", "ATTRIBUTE_DEFINITION_XHTML_REF"}
	case A_ENUM_VALUE_REF:
		res = []string{"Name", "ENUM_VALUE_REF"}
	case A_OBJECT:
		res = []string{"Name", "SPEC_OBJECT_REF"}
	case A_PROPERTIES:
		res = []string{"Name", "EMBEDDED_VALUE"}
	case A_RELATION_GROUP_TYPE_REF:
		res = []string{"Name", "RELATION_GROUP_TYPE_REF"}
	case A_SOURCE_1:
		res = []string{"Name", "SPEC_OBJECT_REF"}
	case A_SPECIFICATIONS:
		res = []string{"Name", "SPECIFICATION"}
	case A_SPECIFICATION_TYPE_REF:
		res = []string{"Name", "SPECIFICATION_TYPE_REF"}
	case A_SPECIFIED_VALUES:
		res = []string{"Name", "ENUM_VALUE"}
	case A_SPEC_ATTRIBUTES:
		res = []string{"Name", "ATTRIBUTE_DEFINITION_BOOLEAN", "ATTRIBUTE_DEFINITION_DATE", "ATTRIBUTE_DEFINITION_ENUMERATION", "ATTRIBUTE_DEFINITION_INTEGER", "ATTRIBUTE_DEFINITION_REAL", "ATTRIBUTE_DEFINITION_STRING", "ATTRIBUTE_DEFINITION_XHTML"}
	case A_SPEC_OBJECTS:
		res = []string{"Name", "SPEC_OBJECT"}
	case A_SPEC_OBJECT_TYPE_REF:
		res = []string{"Name", "SPEC_OBJECT_TYPE_REF"}
	case A_SPEC_RELATIONS:
		res = []string{"Name", "SPEC_RELATION"}
	case A_SPEC_RELATION_GROUPS:
		res = []string{"Name", "RELATION_GROUP"}
	case A_SPEC_RELATION_REF:
		res = []string{"Name", "SPEC_RELATION_REF"}
	case A_SPEC_RELATION_TYPE_REF:
		res = []string{"Name", "SPEC_RELATION_TYPE_REF"}
	case A_SPEC_TYPES:
		res = []string{"Name", "RELATION_GROUP_TYPE", "SPEC_OBJECT_TYPE", "SPEC_RELATION_TYPE", "SPECIFICATION_TYPE"}
	case A_TARGET_SPECIFICATION_1:
		res = []string{"Name", "SPECIFICATION_REF"}
	case A_THE_HEADER:
		res = []string{"Name", "REQ_IF_HEADER"}
	case A_TOOL_EXTENSIONS:
		res = []string{"Name", "REQ_IF_TOOL_EXTENSION"}
	case DATATYPE_DEFINITION_BOOLEAN:
		res = []string{"Name", "DESC", "IDENTIFIER", "LAST_CHANGE", "LONG_NAME", "ALTERNATIVE_ID"}
	case DATATYPE_DEFINITION_DATE:
		res = []string{"Name", "DESC", "IDENTIFIER", "LAST_CHANGE", "LONG_NAME", "ALTERNATIVE_ID"}
	case DATATYPE_DEFINITION_ENUMERATION:
		res = []string{"Name", "DESC", "IDENTIFIER", "LAST_CHANGE", "LONG_NAME", "ALTERNATIVE_ID", "SPECIFIED_VALUES"}
	case DATATYPE_DEFINITION_INTEGER:
		res = []string{"Name", "DESC", "IDENTIFIER", "LAST_CHANGE", "LONG_NAME", "MAX", "MIN", "ALTERNATIVE_ID"}
	case DATATYPE_DEFINITION_REAL:
		res = []string{"Name", "ACCURACY", "DESC", "IDENTIFIER", "LAST_CHANGE", "LONG_NAME", "MAX", "MIN", "ALTERNATIVE_ID"}
	case DATATYPE_DEFINITION_STRING:
		res = []string{"Name", "DESC", "IDENTIFIER", "LAST_CHANGE", "LONG_NAME", "MAX_LENGTH", "ALTERNATIVE_ID"}
	case DATATYPE_DEFINITION_XHTML:
		res = []string{"Name", "DESC", "IDENTIFIER", "LAST_CHANGE", "LONG_NAME", "ALTERNATIVE_ID"}
	case EMBEDDED_VALUE:
		res = []string{"Name", "KEY", "OTHER_CONTENT"}
	case ENUM_VALUE:
		res = []string{"Name", "DESC", "IDENTIFIER", "LAST_CHANGE", "LONG_NAME", "ALTERNATIVE_ID", "PROPERTIES"}
	case RELATION_GROUP:
		res = []string{"Name", "DESC", "IDENTIFIER", "LAST_CHANGE", "LONG_NAME", "ALTERNATIVE_ID", "SOURCE_SPECIFICATION", "SPEC_RELATIONS", "TARGET_SPECIFICATION", "TYPE"}
	case RELATION_GROUP_TYPE:
		res = []string{"Name", "DESC", "IDENTIFIER", "LAST_CHANGE", "LONG_NAME", "ALTERNATIVE_ID", "SPEC_ATTRIBUTES"}
	case REQ_IF:
		res = []string{"Name", "Lang", "THE_HEADER", "CORE_CONTENT", "TOOL_EXTENSIONS"}
	case REQ_IF_CONTENT:
		res = []string{"Name", "DATATYPES", "SPEC_TYPES", "SPEC_OBJECTS", "SPEC_RELATIONS", "SPECIFICATIONS", "SPEC_RELATION_GROUPS"}
	case REQ_IF_HEADER:
		res = []string{"Name", "IDENTIFIER", "COMMENT", "CREATION_TIME", "REPOSITORY_ID", "REQ_IF_TOOL_ID", "REQ_IF_VERSION", "SOURCE_TOOL_ID", "TITLE"}
	case REQ_IF_TOOL_EXTENSION:
		res = []string{"Name"}
	case SPECIFICATION:
		res = []string{"Name", "DESC", "IDENTIFIER", "LAST_CHANGE", "LONG_NAME", "ALTERNATIVE_ID", "CHILDREN", "VALUES", "TYPE"}
	case SPECIFICATION_TYPE:
		res = []string{"Name", "DESC", "IDENTIFIER", "LAST_CHANGE", "LONG_NAME", "ALTERNATIVE_ID", "SPEC_ATTRIBUTES"}
	case SPEC_HIERARCHY:
		res = []string{"Name", "DESC", "IDENTIFIER", "IS_EDITABLE", "IS_TABLE_INTERNAL", "LAST_CHANGE", "LONG_NAME", "ALTERNATIVE_ID", "CHILDREN", "EDITABLE_ATTS", "OBJECT"}
	case SPEC_OBJECT:
		res = []string{"Name", "DESC", "IDENTIFIER", "LAST_CHANGE", "LONG_NAME", "ALTERNATIVE_ID", "VALUES", "TYPE"}
	case SPEC_OBJECT_TYPE:
		res = []string{"Name", "DESC", "IDENTIFIER", "LAST_CHANGE", "LONG_NAME", "ALTERNATIVE_ID", "SPEC_ATTRIBUTES"}
	case SPEC_RELATION:
		res = []string{"Name", "DESC", "IDENTIFIER", "LAST_CHANGE", "LONG_NAME", "ALTERNATIVE_ID", "VALUES", "SOURCE", "TARGET", "TYPE"}
	case SPEC_RELATION_TYPE:
		res = []string{"Name", "DESC", "IDENTIFIER", "LAST_CHANGE", "LONG_NAME", "ALTERNATIVE_ID", "SPEC_ATTRIBUTES"}
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
	case A_ALTERNATIVE_ID:
		var rf ReverseField
		_ = rf
	case A_ATTRIBUTE_DEFINITION_BOOLEAN_REF:
		var rf ReverseField
		_ = rf
	case A_ATTRIBUTE_DEFINITION_DATE_REF:
		var rf ReverseField
		_ = rf
	case A_ATTRIBUTE_DEFINITION_ENUMERATION_REF:
		var rf ReverseField
		_ = rf
	case A_ATTRIBUTE_DEFINITION_INTEGER_REF:
		var rf ReverseField
		_ = rf
	case A_ATTRIBUTE_DEFINITION_REAL_REF:
		var rf ReverseField
		_ = rf
	case A_ATTRIBUTE_DEFINITION_STRING_REF:
		var rf ReverseField
		_ = rf
	case A_ATTRIBUTE_DEFINITION_XHTML_REF:
		var rf ReverseField
		_ = rf
	case A_ATTRIBUTE_VALUE_BOOLEAN:
		var rf ReverseField
		_ = rf
	case A_ATTRIBUTE_VALUE_DATE:
		var rf ReverseField
		_ = rf
	case A_ATTRIBUTE_VALUE_ENUMERATION:
		var rf ReverseField
		_ = rf
	case A_ATTRIBUTE_VALUE_INTEGER:
		var rf ReverseField
		_ = rf
	case A_ATTRIBUTE_VALUE_REAL:
		var rf ReverseField
		_ = rf
	case A_ATTRIBUTE_VALUE_STRING:
		var rf ReverseField
		_ = rf
	case A_ATTRIBUTE_VALUE_XHTML:
		var rf ReverseField
		_ = rf
	case A_ATTRIBUTE_VALUE_XHTML_1:
		var rf ReverseField
		_ = rf
	case A_CHILDREN:
		var rf ReverseField
		_ = rf
	case A_CORE_CONTENT:
		var rf ReverseField
		_ = rf
	case A_DATATYPES:
		var rf ReverseField
		_ = rf
	case A_DATATYPE_DEFINITION_BOOLEAN_REF:
		var rf ReverseField
		_ = rf
	case A_DATATYPE_DEFINITION_DATE_REF:
		var rf ReverseField
		_ = rf
	case A_DATATYPE_DEFINITION_ENUMERATION_REF:
		var rf ReverseField
		_ = rf
	case A_DATATYPE_DEFINITION_INTEGER_REF:
		var rf ReverseField
		_ = rf
	case A_DATATYPE_DEFINITION_REAL_REF:
		var rf ReverseField
		_ = rf
	case A_DATATYPE_DEFINITION_STRING_REF:
		var rf ReverseField
		_ = rf
	case A_DATATYPE_DEFINITION_XHTML_REF:
		var rf ReverseField
		_ = rf
	case A_EDITABLE_ATTS:
		var rf ReverseField
		_ = rf
	case A_ENUM_VALUE_REF:
		var rf ReverseField
		_ = rf
	case A_OBJECT:
		var rf ReverseField
		_ = rf
	case A_PROPERTIES:
		var rf ReverseField
		_ = rf
	case A_RELATION_GROUP_TYPE_REF:
		var rf ReverseField
		_ = rf
	case A_SOURCE_1:
		var rf ReverseField
		_ = rf
	case A_SPECIFICATIONS:
		var rf ReverseField
		_ = rf
	case A_SPECIFICATION_TYPE_REF:
		var rf ReverseField
		_ = rf
	case A_SPECIFIED_VALUES:
		var rf ReverseField
		_ = rf
	case A_SPEC_ATTRIBUTES:
		var rf ReverseField
		_ = rf
	case A_SPEC_OBJECTS:
		var rf ReverseField
		_ = rf
	case A_SPEC_OBJECT_TYPE_REF:
		var rf ReverseField
		_ = rf
	case A_SPEC_RELATIONS:
		var rf ReverseField
		_ = rf
	case A_SPEC_RELATION_GROUPS:
		var rf ReverseField
		_ = rf
	case A_SPEC_RELATION_REF:
		var rf ReverseField
		_ = rf
	case A_SPEC_RELATION_TYPE_REF:
		var rf ReverseField
		_ = rf
	case A_SPEC_TYPES:
		var rf ReverseField
		_ = rf
	case A_TARGET_SPECIFICATION_1:
		var rf ReverseField
		_ = rf
	case A_THE_HEADER:
		var rf ReverseField
		_ = rf
	case A_TOOL_EXTENSIONS:
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
		res = []string{"Name", "DESC", "IDENTIFIER", "IS_EDITABLE", "LAST_CHANGE", "LONG_NAME", "ALTERNATIVE_ID", "DEFAULT_VALUE", "TYPE"}
	case *ATTRIBUTE_DEFINITION_DATE:
		res = []string{"Name", "DESC", "IDENTIFIER", "IS_EDITABLE", "LAST_CHANGE", "LONG_NAME", "ALTERNATIVE_ID", "DEFAULT_VALUE", "TYPE"}
	case *ATTRIBUTE_DEFINITION_ENUMERATION:
		res = []string{"Name", "DESC", "IDENTIFIER", "IS_EDITABLE", "LAST_CHANGE", "LONG_NAME", "MULTI_VALUED", "ALTERNATIVE_ID", "DEFAULT_VALUE", "TYPE"}
	case *ATTRIBUTE_DEFINITION_INTEGER:
		res = []string{"Name", "DESC", "IDENTIFIER", "IS_EDITABLE", "LAST_CHANGE", "LONG_NAME", "ALTERNATIVE_ID", "DEFAULT_VALUE", "TYPE"}
	case *ATTRIBUTE_DEFINITION_REAL:
		res = []string{"Name", "DESC", "IDENTIFIER", "IS_EDITABLE", "LAST_CHANGE", "LONG_NAME", "ALTERNATIVE_ID", "DEFAULT_VALUE", "TYPE"}
	case *ATTRIBUTE_DEFINITION_STRING:
		res = []string{"Name", "DESC", "IDENTIFIER", "IS_EDITABLE", "LAST_CHANGE", "LONG_NAME", "ALTERNATIVE_ID", "DEFAULT_VALUE", "TYPE"}
	case *ATTRIBUTE_DEFINITION_XHTML:
		res = []string{"Name", "DESC", "IDENTIFIER", "IS_EDITABLE", "LAST_CHANGE", "LONG_NAME", "ALTERNATIVE_ID", "DEFAULT_VALUE", "TYPE"}
	case *ATTRIBUTE_VALUE_BOOLEAN:
		res = []string{"Name", "THE_VALUE", "DEFINITION"}
	case *ATTRIBUTE_VALUE_DATE:
		res = []string{"Name", "THE_VALUE", "DEFINITION"}
	case *ATTRIBUTE_VALUE_ENUMERATION:
		res = []string{"Name", "DEFINITION", "VALUES"}
	case *ATTRIBUTE_VALUE_INTEGER:
		res = []string{"Name", "THE_VALUE", "DEFINITION"}
	case *ATTRIBUTE_VALUE_REAL:
		res = []string{"Name", "THE_VALUE", "DEFINITION"}
	case *ATTRIBUTE_VALUE_STRING:
		res = []string{"Name", "THE_VALUE", "DEFINITION"}
	case *ATTRIBUTE_VALUE_XHTML:
		res = []string{"Name", "IS_SIMPLIFIED", "THE_VALUE", "THE_ORIGINAL_VALUE", "DEFINITION"}
	case *A_ALTERNATIVE_ID:
		res = []string{"Name", "ALTERNATIVE_ID"}
	case *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF:
		res = []string{"Name", "ATTRIBUTE_DEFINITION_BOOLEAN_REF"}
	case *A_ATTRIBUTE_DEFINITION_DATE_REF:
		res = []string{"Name", "ATTRIBUTE_DEFINITION_DATE_REF"}
	case *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF:
		res = []string{"Name", "ATTRIBUTE_DEFINITION_ENUMERATION_REF"}
	case *A_ATTRIBUTE_DEFINITION_INTEGER_REF:
		res = []string{"Name", "ATTRIBUTE_DEFINITION_INTEGER_REF"}
	case *A_ATTRIBUTE_DEFINITION_REAL_REF:
		res = []string{"Name", "ATTRIBUTE_DEFINITION_REAL_REF"}
	case *A_ATTRIBUTE_DEFINITION_STRING_REF:
		res = []string{"Name", "ATTRIBUTE_DEFINITION_STRING_REF"}
	case *A_ATTRIBUTE_DEFINITION_XHTML_REF:
		res = []string{"Name", "ATTRIBUTE_DEFINITION_XHTML_REF"}
	case *A_ATTRIBUTE_VALUE_BOOLEAN:
		res = []string{"Name", "ATTRIBUTE_VALUE_BOOLEAN"}
	case *A_ATTRIBUTE_VALUE_DATE:
		res = []string{"Name", "ATTRIBUTE_VALUE_DATE"}
	case *A_ATTRIBUTE_VALUE_ENUMERATION:
		res = []string{"Name", "ATTRIBUTE_VALUE_ENUMERATION"}
	case *A_ATTRIBUTE_VALUE_INTEGER:
		res = []string{"Name", "ATTRIBUTE_VALUE_INTEGER"}
	case *A_ATTRIBUTE_VALUE_REAL:
		res = []string{"Name", "ATTRIBUTE_VALUE_REAL"}
	case *A_ATTRIBUTE_VALUE_STRING:
		res = []string{"Name", "ATTRIBUTE_VALUE_STRING"}
	case *A_ATTRIBUTE_VALUE_XHTML:
		res = []string{"Name", "ATTRIBUTE_VALUE_BOOLEAN", "ATTRIBUTE_VALUE_DATE", "ATTRIBUTE_VALUE_ENUMERATION", "ATTRIBUTE_VALUE_INTEGER", "ATTRIBUTE_VALUE_REAL", "ATTRIBUTE_VALUE_STRING", "ATTRIBUTE_VALUE_XHTML"}
	case *A_ATTRIBUTE_VALUE_XHTML_1:
		res = []string{"Name", "ATTRIBUTE_VALUE_XHTML"}
	case *A_CHILDREN:
		res = []string{"Name", "SPEC_HIERARCHY"}
	case *A_CORE_CONTENT:
		res = []string{"Name", "REQ_IF_CONTENT"}
	case *A_DATATYPES:
		res = []string{"Name", "DATATYPE_DEFINITION_BOOLEAN", "DATATYPE_DEFINITION_DATE", "DATATYPE_DEFINITION_ENUMERATION", "DATATYPE_DEFINITION_INTEGER", "DATATYPE_DEFINITION_REAL", "DATATYPE_DEFINITION_STRING", "DATATYPE_DEFINITION_XHTML"}
	case *A_DATATYPE_DEFINITION_BOOLEAN_REF:
		res = []string{"Name", "DATATYPE_DEFINITION_BOOLEAN_REF"}
	case *A_DATATYPE_DEFINITION_DATE_REF:
		res = []string{"Name", "DATATYPE_DEFINITION_DATE_REF"}
	case *A_DATATYPE_DEFINITION_ENUMERATION_REF:
		res = []string{"Name", "DATATYPE_DEFINITION_ENUMERATION_REF"}
	case *A_DATATYPE_DEFINITION_INTEGER_REF:
		res = []string{"Name", "DATATYPE_DEFINITION_INTEGER_REF"}
	case *A_DATATYPE_DEFINITION_REAL_REF:
		res = []string{"Name", "DATATYPE_DEFINITION_REAL_REF"}
	case *A_DATATYPE_DEFINITION_STRING_REF:
		res = []string{"Name", "DATATYPE_DEFINITION_STRING_REF"}
	case *A_DATATYPE_DEFINITION_XHTML_REF:
		res = []string{"Name", "DATATYPE_DEFINITION_XHTML_REF"}
	case *A_EDITABLE_ATTS:
		res = []string{"Name", "ATTRIBUTE_DEFINITION_BOOLEAN_REF", "ATTRIBUTE_DEFINITION_DATE_REF", "ATTRIBUTE_DEFINITION_ENUMERATION_REF", "ATTRIBUTE_DEFINITION_INTEGER_REF", "ATTRIBUTE_DEFINITION_REAL_REF", "ATTRIBUTE_DEFINITION_STRING_REF", "ATTRIBUTE_DEFINITION_XHTML_REF"}
	case *A_ENUM_VALUE_REF:
		res = []string{"Name", "ENUM_VALUE_REF"}
	case *A_OBJECT:
		res = []string{"Name", "SPEC_OBJECT_REF"}
	case *A_PROPERTIES:
		res = []string{"Name", "EMBEDDED_VALUE"}
	case *A_RELATION_GROUP_TYPE_REF:
		res = []string{"Name", "RELATION_GROUP_TYPE_REF"}
	case *A_SOURCE_1:
		res = []string{"Name", "SPEC_OBJECT_REF"}
	case *A_SPECIFICATIONS:
		res = []string{"Name", "SPECIFICATION"}
	case *A_SPECIFICATION_TYPE_REF:
		res = []string{"Name", "SPECIFICATION_TYPE_REF"}
	case *A_SPECIFIED_VALUES:
		res = []string{"Name", "ENUM_VALUE"}
	case *A_SPEC_ATTRIBUTES:
		res = []string{"Name", "ATTRIBUTE_DEFINITION_BOOLEAN", "ATTRIBUTE_DEFINITION_DATE", "ATTRIBUTE_DEFINITION_ENUMERATION", "ATTRIBUTE_DEFINITION_INTEGER", "ATTRIBUTE_DEFINITION_REAL", "ATTRIBUTE_DEFINITION_STRING", "ATTRIBUTE_DEFINITION_XHTML"}
	case *A_SPEC_OBJECTS:
		res = []string{"Name", "SPEC_OBJECT"}
	case *A_SPEC_OBJECT_TYPE_REF:
		res = []string{"Name", "SPEC_OBJECT_TYPE_REF"}
	case *A_SPEC_RELATIONS:
		res = []string{"Name", "SPEC_RELATION"}
	case *A_SPEC_RELATION_GROUPS:
		res = []string{"Name", "RELATION_GROUP"}
	case *A_SPEC_RELATION_REF:
		res = []string{"Name", "SPEC_RELATION_REF"}
	case *A_SPEC_RELATION_TYPE_REF:
		res = []string{"Name", "SPEC_RELATION_TYPE_REF"}
	case *A_SPEC_TYPES:
		res = []string{"Name", "RELATION_GROUP_TYPE", "SPEC_OBJECT_TYPE", "SPEC_RELATION_TYPE", "SPECIFICATION_TYPE"}
	case *A_TARGET_SPECIFICATION_1:
		res = []string{"Name", "SPECIFICATION_REF"}
	case *A_THE_HEADER:
		res = []string{"Name", "REQ_IF_HEADER"}
	case *A_TOOL_EXTENSIONS:
		res = []string{"Name", "REQ_IF_TOOL_EXTENSION"}
	case *DATATYPE_DEFINITION_BOOLEAN:
		res = []string{"Name", "DESC", "IDENTIFIER", "LAST_CHANGE", "LONG_NAME", "ALTERNATIVE_ID"}
	case *DATATYPE_DEFINITION_DATE:
		res = []string{"Name", "DESC", "IDENTIFIER", "LAST_CHANGE", "LONG_NAME", "ALTERNATIVE_ID"}
	case *DATATYPE_DEFINITION_ENUMERATION:
		res = []string{"Name", "DESC", "IDENTIFIER", "LAST_CHANGE", "LONG_NAME", "ALTERNATIVE_ID", "SPECIFIED_VALUES"}
	case *DATATYPE_DEFINITION_INTEGER:
		res = []string{"Name", "DESC", "IDENTIFIER", "LAST_CHANGE", "LONG_NAME", "MAX", "MIN", "ALTERNATIVE_ID"}
	case *DATATYPE_DEFINITION_REAL:
		res = []string{"Name", "ACCURACY", "DESC", "IDENTIFIER", "LAST_CHANGE", "LONG_NAME", "MAX", "MIN", "ALTERNATIVE_ID"}
	case *DATATYPE_DEFINITION_STRING:
		res = []string{"Name", "DESC", "IDENTIFIER", "LAST_CHANGE", "LONG_NAME", "MAX_LENGTH", "ALTERNATIVE_ID"}
	case *DATATYPE_DEFINITION_XHTML:
		res = []string{"Name", "DESC", "IDENTIFIER", "LAST_CHANGE", "LONG_NAME", "ALTERNATIVE_ID"}
	case *EMBEDDED_VALUE:
		res = []string{"Name", "KEY", "OTHER_CONTENT"}
	case *ENUM_VALUE:
		res = []string{"Name", "DESC", "IDENTIFIER", "LAST_CHANGE", "LONG_NAME", "ALTERNATIVE_ID", "PROPERTIES"}
	case *RELATION_GROUP:
		res = []string{"Name", "DESC", "IDENTIFIER", "LAST_CHANGE", "LONG_NAME", "ALTERNATIVE_ID", "SOURCE_SPECIFICATION", "SPEC_RELATIONS", "TARGET_SPECIFICATION", "TYPE"}
	case *RELATION_GROUP_TYPE:
		res = []string{"Name", "DESC", "IDENTIFIER", "LAST_CHANGE", "LONG_NAME", "ALTERNATIVE_ID", "SPEC_ATTRIBUTES"}
	case *REQ_IF:
		res = []string{"Name", "Lang", "THE_HEADER", "CORE_CONTENT", "TOOL_EXTENSIONS"}
	case *REQ_IF_CONTENT:
		res = []string{"Name", "DATATYPES", "SPEC_TYPES", "SPEC_OBJECTS", "SPEC_RELATIONS", "SPECIFICATIONS", "SPEC_RELATION_GROUPS"}
	case *REQ_IF_HEADER:
		res = []string{"Name", "IDENTIFIER", "COMMENT", "CREATION_TIME", "REPOSITORY_ID", "REQ_IF_TOOL_ID", "REQ_IF_VERSION", "SOURCE_TOOL_ID", "TITLE"}
	case *REQ_IF_TOOL_EXTENSION:
		res = []string{"Name"}
	case *SPECIFICATION:
		res = []string{"Name", "DESC", "IDENTIFIER", "LAST_CHANGE", "LONG_NAME", "ALTERNATIVE_ID", "CHILDREN", "VALUES", "TYPE"}
	case *SPECIFICATION_TYPE:
		res = []string{"Name", "DESC", "IDENTIFIER", "LAST_CHANGE", "LONG_NAME", "ALTERNATIVE_ID", "SPEC_ATTRIBUTES"}
	case *SPEC_HIERARCHY:
		res = []string{"Name", "DESC", "IDENTIFIER", "IS_EDITABLE", "IS_TABLE_INTERNAL", "LAST_CHANGE", "LONG_NAME", "ALTERNATIVE_ID", "CHILDREN", "EDITABLE_ATTS", "OBJECT"}
	case *SPEC_OBJECT:
		res = []string{"Name", "DESC", "IDENTIFIER", "LAST_CHANGE", "LONG_NAME", "ALTERNATIVE_ID", "VALUES", "TYPE"}
	case *SPEC_OBJECT_TYPE:
		res = []string{"Name", "DESC", "IDENTIFIER", "LAST_CHANGE", "LONG_NAME", "ALTERNATIVE_ID", "SPEC_ATTRIBUTES"}
	case *SPEC_RELATION:
		res = []string{"Name", "DESC", "IDENTIFIER", "LAST_CHANGE", "LONG_NAME", "ALTERNATIVE_ID", "VALUES", "SOURCE", "TARGET", "TYPE"}
	case *SPEC_RELATION_TYPE:
		res = []string{"Name", "DESC", "IDENTIFIER", "LAST_CHANGE", "LONG_NAME", "ALTERNATIVE_ID", "SPEC_ATTRIBUTES"}
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
		case "ALTERNATIVE_ID":
			if inferedInstance.ALTERNATIVE_ID != nil {
				res = inferedInstance.ALTERNATIVE_ID.Name
			}
		case "DEFAULT_VALUE":
			if inferedInstance.DEFAULT_VALUE != nil {
				res = inferedInstance.DEFAULT_VALUE.Name
			}
		case "TYPE":
			if inferedInstance.TYPE != nil {
				res = inferedInstance.TYPE.Name
			}
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
		case "ALTERNATIVE_ID":
			if inferedInstance.ALTERNATIVE_ID != nil {
				res = inferedInstance.ALTERNATIVE_ID.Name
			}
		case "DEFAULT_VALUE":
			if inferedInstance.DEFAULT_VALUE != nil {
				res = inferedInstance.DEFAULT_VALUE.Name
			}
		case "TYPE":
			if inferedInstance.TYPE != nil {
				res = inferedInstance.TYPE.Name
			}
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
		case "ALTERNATIVE_ID":
			if inferedInstance.ALTERNATIVE_ID != nil {
				res = inferedInstance.ALTERNATIVE_ID.Name
			}
		case "DEFAULT_VALUE":
			if inferedInstance.DEFAULT_VALUE != nil {
				res = inferedInstance.DEFAULT_VALUE.Name
			}
		case "TYPE":
			if inferedInstance.TYPE != nil {
				res = inferedInstance.TYPE.Name
			}
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
		case "ALTERNATIVE_ID":
			if inferedInstance.ALTERNATIVE_ID != nil {
				res = inferedInstance.ALTERNATIVE_ID.Name
			}
		case "DEFAULT_VALUE":
			if inferedInstance.DEFAULT_VALUE != nil {
				res = inferedInstance.DEFAULT_VALUE.Name
			}
		case "TYPE":
			if inferedInstance.TYPE != nil {
				res = inferedInstance.TYPE.Name
			}
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
		case "ALTERNATIVE_ID":
			if inferedInstance.ALTERNATIVE_ID != nil {
				res = inferedInstance.ALTERNATIVE_ID.Name
			}
		case "DEFAULT_VALUE":
			if inferedInstance.DEFAULT_VALUE != nil {
				res = inferedInstance.DEFAULT_VALUE.Name
			}
		case "TYPE":
			if inferedInstance.TYPE != nil {
				res = inferedInstance.TYPE.Name
			}
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
		case "ALTERNATIVE_ID":
			if inferedInstance.ALTERNATIVE_ID != nil {
				res = inferedInstance.ALTERNATIVE_ID.Name
			}
		case "DEFAULT_VALUE":
			if inferedInstance.DEFAULT_VALUE != nil {
				res = inferedInstance.DEFAULT_VALUE.Name
			}
		case "TYPE":
			if inferedInstance.TYPE != nil {
				res = inferedInstance.TYPE.Name
			}
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
		case "ALTERNATIVE_ID":
			if inferedInstance.ALTERNATIVE_ID != nil {
				res = inferedInstance.ALTERNATIVE_ID.Name
			}
		case "DEFAULT_VALUE":
			if inferedInstance.DEFAULT_VALUE != nil {
				res = inferedInstance.DEFAULT_VALUE.Name
			}
		case "TYPE":
			if inferedInstance.TYPE != nil {
				res = inferedInstance.TYPE.Name
			}
		}
	case *ATTRIBUTE_VALUE_BOOLEAN:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "THE_VALUE":
			res = fmt.Sprintf("%t", inferedInstance.THE_VALUE)
		case "DEFINITION":
			if inferedInstance.DEFINITION != nil {
				res = inferedInstance.DEFINITION.Name
			}
		}
	case *ATTRIBUTE_VALUE_DATE:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "THE_VALUE":
			res = inferedInstance.THE_VALUE
		case "DEFINITION":
			if inferedInstance.DEFINITION != nil {
				res = inferedInstance.DEFINITION.Name
			}
		}
	case *ATTRIBUTE_VALUE_ENUMERATION:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DEFINITION":
			if inferedInstance.DEFINITION != nil {
				res = inferedInstance.DEFINITION.Name
			}
		case "VALUES":
			if inferedInstance.VALUES != nil {
				res = inferedInstance.VALUES.Name
			}
		}
	case *ATTRIBUTE_VALUE_INTEGER:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "THE_VALUE":
			res = fmt.Sprintf("%d", inferedInstance.THE_VALUE)
		case "DEFINITION":
			if inferedInstance.DEFINITION != nil {
				res = inferedInstance.DEFINITION.Name
			}
		}
	case *ATTRIBUTE_VALUE_REAL:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "THE_VALUE":
			res = fmt.Sprintf("%f", inferedInstance.THE_VALUE)
		case "DEFINITION":
			if inferedInstance.DEFINITION != nil {
				res = inferedInstance.DEFINITION.Name
			}
		}
	case *ATTRIBUTE_VALUE_STRING:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "THE_VALUE":
			res = inferedInstance.THE_VALUE
		case "DEFINITION":
			if inferedInstance.DEFINITION != nil {
				res = inferedInstance.DEFINITION.Name
			}
		}
	case *ATTRIBUTE_VALUE_XHTML:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "IS_SIMPLIFIED":
			res = fmt.Sprintf("%t", inferedInstance.IS_SIMPLIFIED)
		case "THE_VALUE":
			if inferedInstance.THE_VALUE != nil {
				res = inferedInstance.THE_VALUE.Name
			}
		case "THE_ORIGINAL_VALUE":
			if inferedInstance.THE_ORIGINAL_VALUE != nil {
				res = inferedInstance.THE_ORIGINAL_VALUE.Name
			}
		case "DEFINITION":
			if inferedInstance.DEFINITION != nil {
				res = inferedInstance.DEFINITION.Name
			}
		}
	case *A_ALTERNATIVE_ID:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "ALTERNATIVE_ID":
			if inferedInstance.ALTERNATIVE_ID != nil {
				res = inferedInstance.ALTERNATIVE_ID.Name
			}
		}
	case *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "ATTRIBUTE_DEFINITION_BOOLEAN_REF":
			res = inferedInstance.ATTRIBUTE_DEFINITION_BOOLEAN_REF
		}
	case *A_ATTRIBUTE_DEFINITION_DATE_REF:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "ATTRIBUTE_DEFINITION_DATE_REF":
			res = inferedInstance.ATTRIBUTE_DEFINITION_DATE_REF
		}
	case *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "ATTRIBUTE_DEFINITION_ENUMERATION_REF":
			res = inferedInstance.ATTRIBUTE_DEFINITION_ENUMERATION_REF
		}
	case *A_ATTRIBUTE_DEFINITION_INTEGER_REF:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "ATTRIBUTE_DEFINITION_INTEGER_REF":
			res = inferedInstance.ATTRIBUTE_DEFINITION_INTEGER_REF
		}
	case *A_ATTRIBUTE_DEFINITION_REAL_REF:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "ATTRIBUTE_DEFINITION_REAL_REF":
			res = inferedInstance.ATTRIBUTE_DEFINITION_REAL_REF
		}
	case *A_ATTRIBUTE_DEFINITION_STRING_REF:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "ATTRIBUTE_DEFINITION_STRING_REF":
			res = inferedInstance.ATTRIBUTE_DEFINITION_STRING_REF
		}
	case *A_ATTRIBUTE_DEFINITION_XHTML_REF:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "ATTRIBUTE_DEFINITION_XHTML_REF":
			res = inferedInstance.ATTRIBUTE_DEFINITION_XHTML_REF
		}
	case *A_ATTRIBUTE_VALUE_BOOLEAN:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "ATTRIBUTE_VALUE_BOOLEAN":
			if inferedInstance.ATTRIBUTE_VALUE_BOOLEAN != nil {
				res = inferedInstance.ATTRIBUTE_VALUE_BOOLEAN.Name
			}
		}
	case *A_ATTRIBUTE_VALUE_DATE:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "ATTRIBUTE_VALUE_DATE":
			if inferedInstance.ATTRIBUTE_VALUE_DATE != nil {
				res = inferedInstance.ATTRIBUTE_VALUE_DATE.Name
			}
		}
	case *A_ATTRIBUTE_VALUE_ENUMERATION:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "ATTRIBUTE_VALUE_ENUMERATION":
			if inferedInstance.ATTRIBUTE_VALUE_ENUMERATION != nil {
				res = inferedInstance.ATTRIBUTE_VALUE_ENUMERATION.Name
			}
		}
	case *A_ATTRIBUTE_VALUE_INTEGER:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "ATTRIBUTE_VALUE_INTEGER":
			if inferedInstance.ATTRIBUTE_VALUE_INTEGER != nil {
				res = inferedInstance.ATTRIBUTE_VALUE_INTEGER.Name
			}
		}
	case *A_ATTRIBUTE_VALUE_REAL:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "ATTRIBUTE_VALUE_REAL":
			if inferedInstance.ATTRIBUTE_VALUE_REAL != nil {
				res = inferedInstance.ATTRIBUTE_VALUE_REAL.Name
			}
		}
	case *A_ATTRIBUTE_VALUE_STRING:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "ATTRIBUTE_VALUE_STRING":
			if inferedInstance.ATTRIBUTE_VALUE_STRING != nil {
				res = inferedInstance.ATTRIBUTE_VALUE_STRING.Name
			}
		}
	case *A_ATTRIBUTE_VALUE_XHTML:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "ATTRIBUTE_VALUE_BOOLEAN":
			if inferedInstance.ATTRIBUTE_VALUE_BOOLEAN != nil {
				res = inferedInstance.ATTRIBUTE_VALUE_BOOLEAN.Name
			}
		case "ATTRIBUTE_VALUE_DATE":
			if inferedInstance.ATTRIBUTE_VALUE_DATE != nil {
				res = inferedInstance.ATTRIBUTE_VALUE_DATE.Name
			}
		case "ATTRIBUTE_VALUE_ENUMERATION":
			if inferedInstance.ATTRIBUTE_VALUE_ENUMERATION != nil {
				res = inferedInstance.ATTRIBUTE_VALUE_ENUMERATION.Name
			}
		case "ATTRIBUTE_VALUE_INTEGER":
			if inferedInstance.ATTRIBUTE_VALUE_INTEGER != nil {
				res = inferedInstance.ATTRIBUTE_VALUE_INTEGER.Name
			}
		case "ATTRIBUTE_VALUE_REAL":
			if inferedInstance.ATTRIBUTE_VALUE_REAL != nil {
				res = inferedInstance.ATTRIBUTE_VALUE_REAL.Name
			}
		case "ATTRIBUTE_VALUE_STRING":
			if inferedInstance.ATTRIBUTE_VALUE_STRING != nil {
				res = inferedInstance.ATTRIBUTE_VALUE_STRING.Name
			}
		case "ATTRIBUTE_VALUE_XHTML":
			if inferedInstance.ATTRIBUTE_VALUE_XHTML != nil {
				res = inferedInstance.ATTRIBUTE_VALUE_XHTML.Name
			}
		}
	case *A_ATTRIBUTE_VALUE_XHTML_1:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "ATTRIBUTE_VALUE_XHTML":
			if inferedInstance.ATTRIBUTE_VALUE_XHTML != nil {
				res = inferedInstance.ATTRIBUTE_VALUE_XHTML.Name
			}
		}
	case *A_CHILDREN:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "SPEC_HIERARCHY":
			if inferedInstance.SPEC_HIERARCHY != nil {
				res = inferedInstance.SPEC_HIERARCHY.Name
			}
		}
	case *A_CORE_CONTENT:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "REQ_IF_CONTENT":
			if inferedInstance.REQ_IF_CONTENT != nil {
				res = inferedInstance.REQ_IF_CONTENT.Name
			}
		}
	case *A_DATATYPES:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DATATYPE_DEFINITION_BOOLEAN":
			if inferedInstance.DATATYPE_DEFINITION_BOOLEAN != nil {
				res = inferedInstance.DATATYPE_DEFINITION_BOOLEAN.Name
			}
		case "DATATYPE_DEFINITION_DATE":
			if inferedInstance.DATATYPE_DEFINITION_DATE != nil {
				res = inferedInstance.DATATYPE_DEFINITION_DATE.Name
			}
		case "DATATYPE_DEFINITION_ENUMERATION":
			if inferedInstance.DATATYPE_DEFINITION_ENUMERATION != nil {
				res = inferedInstance.DATATYPE_DEFINITION_ENUMERATION.Name
			}
		case "DATATYPE_DEFINITION_INTEGER":
			if inferedInstance.DATATYPE_DEFINITION_INTEGER != nil {
				res = inferedInstance.DATATYPE_DEFINITION_INTEGER.Name
			}
		case "DATATYPE_DEFINITION_REAL":
			if inferedInstance.DATATYPE_DEFINITION_REAL != nil {
				res = inferedInstance.DATATYPE_DEFINITION_REAL.Name
			}
		case "DATATYPE_DEFINITION_STRING":
			if inferedInstance.DATATYPE_DEFINITION_STRING != nil {
				res = inferedInstance.DATATYPE_DEFINITION_STRING.Name
			}
		case "DATATYPE_DEFINITION_XHTML":
			if inferedInstance.DATATYPE_DEFINITION_XHTML != nil {
				res = inferedInstance.DATATYPE_DEFINITION_XHTML.Name
			}
		}
	case *A_DATATYPE_DEFINITION_BOOLEAN_REF:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DATATYPE_DEFINITION_BOOLEAN_REF":
			res = inferedInstance.DATATYPE_DEFINITION_BOOLEAN_REF
		}
	case *A_DATATYPE_DEFINITION_DATE_REF:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DATATYPE_DEFINITION_DATE_REF":
			res = inferedInstance.DATATYPE_DEFINITION_DATE_REF
		}
	case *A_DATATYPE_DEFINITION_ENUMERATION_REF:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DATATYPE_DEFINITION_ENUMERATION_REF":
			res = inferedInstance.DATATYPE_DEFINITION_ENUMERATION_REF
		}
	case *A_DATATYPE_DEFINITION_INTEGER_REF:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DATATYPE_DEFINITION_INTEGER_REF":
			res = inferedInstance.DATATYPE_DEFINITION_INTEGER_REF
		}
	case *A_DATATYPE_DEFINITION_REAL_REF:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DATATYPE_DEFINITION_REAL_REF":
			res = inferedInstance.DATATYPE_DEFINITION_REAL_REF
		}
	case *A_DATATYPE_DEFINITION_STRING_REF:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DATATYPE_DEFINITION_STRING_REF":
			res = inferedInstance.DATATYPE_DEFINITION_STRING_REF
		}
	case *A_DATATYPE_DEFINITION_XHTML_REF:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DATATYPE_DEFINITION_XHTML_REF":
			res = inferedInstance.DATATYPE_DEFINITION_XHTML_REF
		}
	case *A_EDITABLE_ATTS:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "ATTRIBUTE_DEFINITION_BOOLEAN_REF":
			res = inferedInstance.ATTRIBUTE_DEFINITION_BOOLEAN_REF
		case "ATTRIBUTE_DEFINITION_DATE_REF":
			res = inferedInstance.ATTRIBUTE_DEFINITION_DATE_REF
		case "ATTRIBUTE_DEFINITION_ENUMERATION_REF":
			res = inferedInstance.ATTRIBUTE_DEFINITION_ENUMERATION_REF
		case "ATTRIBUTE_DEFINITION_INTEGER_REF":
			res = inferedInstance.ATTRIBUTE_DEFINITION_INTEGER_REF
		case "ATTRIBUTE_DEFINITION_REAL_REF":
			res = inferedInstance.ATTRIBUTE_DEFINITION_REAL_REF
		case "ATTRIBUTE_DEFINITION_STRING_REF":
			res = inferedInstance.ATTRIBUTE_DEFINITION_STRING_REF
		case "ATTRIBUTE_DEFINITION_XHTML_REF":
			res = inferedInstance.ATTRIBUTE_DEFINITION_XHTML_REF
		}
	case *A_ENUM_VALUE_REF:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "ENUM_VALUE_REF":
			res = inferedInstance.ENUM_VALUE_REF
		}
	case *A_OBJECT:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "SPEC_OBJECT_REF":
			res = inferedInstance.SPEC_OBJECT_REF
		}
	case *A_PROPERTIES:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "EMBEDDED_VALUE":
			if inferedInstance.EMBEDDED_VALUE != nil {
				res = inferedInstance.EMBEDDED_VALUE.Name
			}
		}
	case *A_RELATION_GROUP_TYPE_REF:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "RELATION_GROUP_TYPE_REF":
			res = inferedInstance.RELATION_GROUP_TYPE_REF
		}
	case *A_SOURCE_1:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "SPEC_OBJECT_REF":
			res = inferedInstance.SPEC_OBJECT_REF
		}
	case *A_SPECIFICATIONS:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "SPECIFICATION":
			if inferedInstance.SPECIFICATION != nil {
				res = inferedInstance.SPECIFICATION.Name
			}
		}
	case *A_SPECIFICATION_TYPE_REF:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "SPECIFICATION_TYPE_REF":
			res = inferedInstance.SPECIFICATION_TYPE_REF
		}
	case *A_SPECIFIED_VALUES:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "ENUM_VALUE":
			if inferedInstance.ENUM_VALUE != nil {
				res = inferedInstance.ENUM_VALUE.Name
			}
		}
	case *A_SPEC_ATTRIBUTES:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "ATTRIBUTE_DEFINITION_BOOLEAN":
			if inferedInstance.ATTRIBUTE_DEFINITION_BOOLEAN != nil {
				res = inferedInstance.ATTRIBUTE_DEFINITION_BOOLEAN.Name
			}
		case "ATTRIBUTE_DEFINITION_DATE":
			if inferedInstance.ATTRIBUTE_DEFINITION_DATE != nil {
				res = inferedInstance.ATTRIBUTE_DEFINITION_DATE.Name
			}
		case "ATTRIBUTE_DEFINITION_ENUMERATION":
			if inferedInstance.ATTRIBUTE_DEFINITION_ENUMERATION != nil {
				res = inferedInstance.ATTRIBUTE_DEFINITION_ENUMERATION.Name
			}
		case "ATTRIBUTE_DEFINITION_INTEGER":
			if inferedInstance.ATTRIBUTE_DEFINITION_INTEGER != nil {
				res = inferedInstance.ATTRIBUTE_DEFINITION_INTEGER.Name
			}
		case "ATTRIBUTE_DEFINITION_REAL":
			if inferedInstance.ATTRIBUTE_DEFINITION_REAL != nil {
				res = inferedInstance.ATTRIBUTE_DEFINITION_REAL.Name
			}
		case "ATTRIBUTE_DEFINITION_STRING":
			if inferedInstance.ATTRIBUTE_DEFINITION_STRING != nil {
				res = inferedInstance.ATTRIBUTE_DEFINITION_STRING.Name
			}
		case "ATTRIBUTE_DEFINITION_XHTML":
			if inferedInstance.ATTRIBUTE_DEFINITION_XHTML != nil {
				res = inferedInstance.ATTRIBUTE_DEFINITION_XHTML.Name
			}
		}
	case *A_SPEC_OBJECTS:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "SPEC_OBJECT":
			if inferedInstance.SPEC_OBJECT != nil {
				res = inferedInstance.SPEC_OBJECT.Name
			}
		}
	case *A_SPEC_OBJECT_TYPE_REF:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "SPEC_OBJECT_TYPE_REF":
			res = inferedInstance.SPEC_OBJECT_TYPE_REF
		}
	case *A_SPEC_RELATIONS:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "SPEC_RELATION":
			if inferedInstance.SPEC_RELATION != nil {
				res = inferedInstance.SPEC_RELATION.Name
			}
		}
	case *A_SPEC_RELATION_GROUPS:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "RELATION_GROUP":
			if inferedInstance.RELATION_GROUP != nil {
				res = inferedInstance.RELATION_GROUP.Name
			}
		}
	case *A_SPEC_RELATION_REF:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "SPEC_RELATION_REF":
			res = inferedInstance.SPEC_RELATION_REF
		}
	case *A_SPEC_RELATION_TYPE_REF:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "SPEC_RELATION_TYPE_REF":
			res = inferedInstance.SPEC_RELATION_TYPE_REF
		}
	case *A_SPEC_TYPES:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "RELATION_GROUP_TYPE":
			if inferedInstance.RELATION_GROUP_TYPE != nil {
				res = inferedInstance.RELATION_GROUP_TYPE.Name
			}
		case "SPEC_OBJECT_TYPE":
			if inferedInstance.SPEC_OBJECT_TYPE != nil {
				res = inferedInstance.SPEC_OBJECT_TYPE.Name
			}
		case "SPEC_RELATION_TYPE":
			if inferedInstance.SPEC_RELATION_TYPE != nil {
				res = inferedInstance.SPEC_RELATION_TYPE.Name
			}
		case "SPECIFICATION_TYPE":
			if inferedInstance.SPECIFICATION_TYPE != nil {
				res = inferedInstance.SPECIFICATION_TYPE.Name
			}
		}
	case *A_TARGET_SPECIFICATION_1:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "SPECIFICATION_REF":
			res = inferedInstance.SPECIFICATION_REF
		}
	case *A_THE_HEADER:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "REQ_IF_HEADER":
			if inferedInstance.REQ_IF_HEADER != nil {
				res = inferedInstance.REQ_IF_HEADER.Name
			}
		}
	case *A_TOOL_EXTENSIONS:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "REQ_IF_TOOL_EXTENSION":
			if inferedInstance.REQ_IF_TOOL_EXTENSION != nil {
				res = inferedInstance.REQ_IF_TOOL_EXTENSION.Name
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
		case "ALTERNATIVE_ID":
			if inferedInstance.ALTERNATIVE_ID != nil {
				res = inferedInstance.ALTERNATIVE_ID.Name
			}
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
		case "ALTERNATIVE_ID":
			if inferedInstance.ALTERNATIVE_ID != nil {
				res = inferedInstance.ALTERNATIVE_ID.Name
			}
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
		case "ALTERNATIVE_ID":
			if inferedInstance.ALTERNATIVE_ID != nil {
				res = inferedInstance.ALTERNATIVE_ID.Name
			}
		case "SPECIFIED_VALUES":
			if inferedInstance.SPECIFIED_VALUES != nil {
				res = inferedInstance.SPECIFIED_VALUES.Name
			}
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
		case "ALTERNATIVE_ID":
			if inferedInstance.ALTERNATIVE_ID != nil {
				res = inferedInstance.ALTERNATIVE_ID.Name
			}
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
		case "ALTERNATIVE_ID":
			if inferedInstance.ALTERNATIVE_ID != nil {
				res = inferedInstance.ALTERNATIVE_ID.Name
			}
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
		case "ALTERNATIVE_ID":
			if inferedInstance.ALTERNATIVE_ID != nil {
				res = inferedInstance.ALTERNATIVE_ID.Name
			}
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
		case "ALTERNATIVE_ID":
			if inferedInstance.ALTERNATIVE_ID != nil {
				res = inferedInstance.ALTERNATIVE_ID.Name
			}
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
		case "ALTERNATIVE_ID":
			if inferedInstance.ALTERNATIVE_ID != nil {
				res = inferedInstance.ALTERNATIVE_ID.Name
			}
		case "PROPERTIES":
			if inferedInstance.PROPERTIES != nil {
				res = inferedInstance.PROPERTIES.Name
			}
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
		case "ALTERNATIVE_ID":
			if inferedInstance.ALTERNATIVE_ID != nil {
				res = inferedInstance.ALTERNATIVE_ID.Name
			}
		case "SOURCE_SPECIFICATION":
			if inferedInstance.SOURCE_SPECIFICATION != nil {
				res = inferedInstance.SOURCE_SPECIFICATION.Name
			}
		case "SPEC_RELATIONS":
			if inferedInstance.SPEC_RELATIONS != nil {
				res = inferedInstance.SPEC_RELATIONS.Name
			}
		case "TARGET_SPECIFICATION":
			if inferedInstance.TARGET_SPECIFICATION != nil {
				res = inferedInstance.TARGET_SPECIFICATION.Name
			}
		case "TYPE":
			if inferedInstance.TYPE != nil {
				res = inferedInstance.TYPE.Name
			}
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
		case "ALTERNATIVE_ID":
			if inferedInstance.ALTERNATIVE_ID != nil {
				res = inferedInstance.ALTERNATIVE_ID.Name
			}
		case "SPEC_ATTRIBUTES":
			if inferedInstance.SPEC_ATTRIBUTES != nil {
				res = inferedInstance.SPEC_ATTRIBUTES.Name
			}
		}
	case *REQ_IF:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Lang":
			res = inferedInstance.Lang
		case "THE_HEADER":
			if inferedInstance.THE_HEADER != nil {
				res = inferedInstance.THE_HEADER.Name
			}
		case "CORE_CONTENT":
			if inferedInstance.CORE_CONTENT != nil {
				res = inferedInstance.CORE_CONTENT.Name
			}
		case "TOOL_EXTENSIONS":
			if inferedInstance.TOOL_EXTENSIONS != nil {
				res = inferedInstance.TOOL_EXTENSIONS.Name
			}
		}
	case *REQ_IF_CONTENT:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DATATYPES":
			if inferedInstance.DATATYPES != nil {
				res = inferedInstance.DATATYPES.Name
			}
		case "SPEC_TYPES":
			if inferedInstance.SPEC_TYPES != nil {
				res = inferedInstance.SPEC_TYPES.Name
			}
		case "SPEC_OBJECTS":
			if inferedInstance.SPEC_OBJECTS != nil {
				res = inferedInstance.SPEC_OBJECTS.Name
			}
		case "SPEC_RELATIONS":
			if inferedInstance.SPEC_RELATIONS != nil {
				res = inferedInstance.SPEC_RELATIONS.Name
			}
		case "SPECIFICATIONS":
			if inferedInstance.SPECIFICATIONS != nil {
				res = inferedInstance.SPECIFICATIONS.Name
			}
		case "SPEC_RELATION_GROUPS":
			if inferedInstance.SPEC_RELATION_GROUPS != nil {
				res = inferedInstance.SPEC_RELATION_GROUPS.Name
			}
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
		case "ALTERNATIVE_ID":
			if inferedInstance.ALTERNATIVE_ID != nil {
				res = inferedInstance.ALTERNATIVE_ID.Name
			}
		case "CHILDREN":
			if inferedInstance.CHILDREN != nil {
				res = inferedInstance.CHILDREN.Name
			}
		case "VALUES":
			if inferedInstance.VALUES != nil {
				res = inferedInstance.VALUES.Name
			}
		case "TYPE":
			if inferedInstance.TYPE != nil {
				res = inferedInstance.TYPE.Name
			}
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
		case "ALTERNATIVE_ID":
			if inferedInstance.ALTERNATIVE_ID != nil {
				res = inferedInstance.ALTERNATIVE_ID.Name
			}
		case "SPEC_ATTRIBUTES":
			if inferedInstance.SPEC_ATTRIBUTES != nil {
				res = inferedInstance.SPEC_ATTRIBUTES.Name
			}
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
		case "ALTERNATIVE_ID":
			if inferedInstance.ALTERNATIVE_ID != nil {
				res = inferedInstance.ALTERNATIVE_ID.Name
			}
		case "CHILDREN":
			if inferedInstance.CHILDREN != nil {
				res = inferedInstance.CHILDREN.Name
			}
		case "EDITABLE_ATTS":
			if inferedInstance.EDITABLE_ATTS != nil {
				res = inferedInstance.EDITABLE_ATTS.Name
			}
		case "OBJECT":
			if inferedInstance.OBJECT != nil {
				res = inferedInstance.OBJECT.Name
			}
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
		case "ALTERNATIVE_ID":
			if inferedInstance.ALTERNATIVE_ID != nil {
				res = inferedInstance.ALTERNATIVE_ID.Name
			}
		case "VALUES":
			if inferedInstance.VALUES != nil {
				res = inferedInstance.VALUES.Name
			}
		case "TYPE":
			if inferedInstance.TYPE != nil {
				res = inferedInstance.TYPE.Name
			}
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
		case "ALTERNATIVE_ID":
			if inferedInstance.ALTERNATIVE_ID != nil {
				res = inferedInstance.ALTERNATIVE_ID.Name
			}
		case "SPEC_ATTRIBUTES":
			if inferedInstance.SPEC_ATTRIBUTES != nil {
				res = inferedInstance.SPEC_ATTRIBUTES.Name
			}
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
		case "ALTERNATIVE_ID":
			if inferedInstance.ALTERNATIVE_ID != nil {
				res = inferedInstance.ALTERNATIVE_ID.Name
			}
		case "VALUES":
			if inferedInstance.VALUES != nil {
				res = inferedInstance.VALUES.Name
			}
		case "SOURCE":
			if inferedInstance.SOURCE != nil {
				res = inferedInstance.SOURCE.Name
			}
		case "TARGET":
			if inferedInstance.TARGET != nil {
				res = inferedInstance.TARGET.Name
			}
		case "TYPE":
			if inferedInstance.TYPE != nil {
				res = inferedInstance.TYPE.Name
			}
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
		case "ALTERNATIVE_ID":
			if inferedInstance.ALTERNATIVE_ID != nil {
				res = inferedInstance.ALTERNATIVE_ID.Name
			}
		case "SPEC_ATTRIBUTES":
			if inferedInstance.SPEC_ATTRIBUTES != nil {
				res = inferedInstance.SPEC_ATTRIBUTES.Name
			}
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
		case "ALTERNATIVE_ID":
			if inferedInstance.ALTERNATIVE_ID != nil {
				res = inferedInstance.ALTERNATIVE_ID.Name
			}
		case "DEFAULT_VALUE":
			if inferedInstance.DEFAULT_VALUE != nil {
				res = inferedInstance.DEFAULT_VALUE.Name
			}
		case "TYPE":
			if inferedInstance.TYPE != nil {
				res = inferedInstance.TYPE.Name
			}
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
		case "ALTERNATIVE_ID":
			if inferedInstance.ALTERNATIVE_ID != nil {
				res = inferedInstance.ALTERNATIVE_ID.Name
			}
		case "DEFAULT_VALUE":
			if inferedInstance.DEFAULT_VALUE != nil {
				res = inferedInstance.DEFAULT_VALUE.Name
			}
		case "TYPE":
			if inferedInstance.TYPE != nil {
				res = inferedInstance.TYPE.Name
			}
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
		case "ALTERNATIVE_ID":
			if inferedInstance.ALTERNATIVE_ID != nil {
				res = inferedInstance.ALTERNATIVE_ID.Name
			}
		case "DEFAULT_VALUE":
			if inferedInstance.DEFAULT_VALUE != nil {
				res = inferedInstance.DEFAULT_VALUE.Name
			}
		case "TYPE":
			if inferedInstance.TYPE != nil {
				res = inferedInstance.TYPE.Name
			}
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
		case "ALTERNATIVE_ID":
			if inferedInstance.ALTERNATIVE_ID != nil {
				res = inferedInstance.ALTERNATIVE_ID.Name
			}
		case "DEFAULT_VALUE":
			if inferedInstance.DEFAULT_VALUE != nil {
				res = inferedInstance.DEFAULT_VALUE.Name
			}
		case "TYPE":
			if inferedInstance.TYPE != nil {
				res = inferedInstance.TYPE.Name
			}
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
		case "ALTERNATIVE_ID":
			if inferedInstance.ALTERNATIVE_ID != nil {
				res = inferedInstance.ALTERNATIVE_ID.Name
			}
		case "DEFAULT_VALUE":
			if inferedInstance.DEFAULT_VALUE != nil {
				res = inferedInstance.DEFAULT_VALUE.Name
			}
		case "TYPE":
			if inferedInstance.TYPE != nil {
				res = inferedInstance.TYPE.Name
			}
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
		case "ALTERNATIVE_ID":
			if inferedInstance.ALTERNATIVE_ID != nil {
				res = inferedInstance.ALTERNATIVE_ID.Name
			}
		case "DEFAULT_VALUE":
			if inferedInstance.DEFAULT_VALUE != nil {
				res = inferedInstance.DEFAULT_VALUE.Name
			}
		case "TYPE":
			if inferedInstance.TYPE != nil {
				res = inferedInstance.TYPE.Name
			}
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
		case "ALTERNATIVE_ID":
			if inferedInstance.ALTERNATIVE_ID != nil {
				res = inferedInstance.ALTERNATIVE_ID.Name
			}
		case "DEFAULT_VALUE":
			if inferedInstance.DEFAULT_VALUE != nil {
				res = inferedInstance.DEFAULT_VALUE.Name
			}
		case "TYPE":
			if inferedInstance.TYPE != nil {
				res = inferedInstance.TYPE.Name
			}
		}
	case ATTRIBUTE_VALUE_BOOLEAN:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "THE_VALUE":
			res = fmt.Sprintf("%t", inferedInstance.THE_VALUE)
		case "DEFINITION":
			if inferedInstance.DEFINITION != nil {
				res = inferedInstance.DEFINITION.Name
			}
		}
	case ATTRIBUTE_VALUE_DATE:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "THE_VALUE":
			res = inferedInstance.THE_VALUE
		case "DEFINITION":
			if inferedInstance.DEFINITION != nil {
				res = inferedInstance.DEFINITION.Name
			}
		}
	case ATTRIBUTE_VALUE_ENUMERATION:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DEFINITION":
			if inferedInstance.DEFINITION != nil {
				res = inferedInstance.DEFINITION.Name
			}
		case "VALUES":
			if inferedInstance.VALUES != nil {
				res = inferedInstance.VALUES.Name
			}
		}
	case ATTRIBUTE_VALUE_INTEGER:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "THE_VALUE":
			res = fmt.Sprintf("%d", inferedInstance.THE_VALUE)
		case "DEFINITION":
			if inferedInstance.DEFINITION != nil {
				res = inferedInstance.DEFINITION.Name
			}
		}
	case ATTRIBUTE_VALUE_REAL:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "THE_VALUE":
			res = fmt.Sprintf("%f", inferedInstance.THE_VALUE)
		case "DEFINITION":
			if inferedInstance.DEFINITION != nil {
				res = inferedInstance.DEFINITION.Name
			}
		}
	case ATTRIBUTE_VALUE_STRING:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "THE_VALUE":
			res = inferedInstance.THE_VALUE
		case "DEFINITION":
			if inferedInstance.DEFINITION != nil {
				res = inferedInstance.DEFINITION.Name
			}
		}
	case ATTRIBUTE_VALUE_XHTML:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "IS_SIMPLIFIED":
			res = fmt.Sprintf("%t", inferedInstance.IS_SIMPLIFIED)
		case "THE_VALUE":
			if inferedInstance.THE_VALUE != nil {
				res = inferedInstance.THE_VALUE.Name
			}
		case "THE_ORIGINAL_VALUE":
			if inferedInstance.THE_ORIGINAL_VALUE != nil {
				res = inferedInstance.THE_ORIGINAL_VALUE.Name
			}
		case "DEFINITION":
			if inferedInstance.DEFINITION != nil {
				res = inferedInstance.DEFINITION.Name
			}
		}
	case A_ALTERNATIVE_ID:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "ALTERNATIVE_ID":
			if inferedInstance.ALTERNATIVE_ID != nil {
				res = inferedInstance.ALTERNATIVE_ID.Name
			}
		}
	case A_ATTRIBUTE_DEFINITION_BOOLEAN_REF:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "ATTRIBUTE_DEFINITION_BOOLEAN_REF":
			res = inferedInstance.ATTRIBUTE_DEFINITION_BOOLEAN_REF
		}
	case A_ATTRIBUTE_DEFINITION_DATE_REF:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "ATTRIBUTE_DEFINITION_DATE_REF":
			res = inferedInstance.ATTRIBUTE_DEFINITION_DATE_REF
		}
	case A_ATTRIBUTE_DEFINITION_ENUMERATION_REF:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "ATTRIBUTE_DEFINITION_ENUMERATION_REF":
			res = inferedInstance.ATTRIBUTE_DEFINITION_ENUMERATION_REF
		}
	case A_ATTRIBUTE_DEFINITION_INTEGER_REF:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "ATTRIBUTE_DEFINITION_INTEGER_REF":
			res = inferedInstance.ATTRIBUTE_DEFINITION_INTEGER_REF
		}
	case A_ATTRIBUTE_DEFINITION_REAL_REF:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "ATTRIBUTE_DEFINITION_REAL_REF":
			res = inferedInstance.ATTRIBUTE_DEFINITION_REAL_REF
		}
	case A_ATTRIBUTE_DEFINITION_STRING_REF:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "ATTRIBUTE_DEFINITION_STRING_REF":
			res = inferedInstance.ATTRIBUTE_DEFINITION_STRING_REF
		}
	case A_ATTRIBUTE_DEFINITION_XHTML_REF:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "ATTRIBUTE_DEFINITION_XHTML_REF":
			res = inferedInstance.ATTRIBUTE_DEFINITION_XHTML_REF
		}
	case A_ATTRIBUTE_VALUE_BOOLEAN:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "ATTRIBUTE_VALUE_BOOLEAN":
			if inferedInstance.ATTRIBUTE_VALUE_BOOLEAN != nil {
				res = inferedInstance.ATTRIBUTE_VALUE_BOOLEAN.Name
			}
		}
	case A_ATTRIBUTE_VALUE_DATE:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "ATTRIBUTE_VALUE_DATE":
			if inferedInstance.ATTRIBUTE_VALUE_DATE != nil {
				res = inferedInstance.ATTRIBUTE_VALUE_DATE.Name
			}
		}
	case A_ATTRIBUTE_VALUE_ENUMERATION:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "ATTRIBUTE_VALUE_ENUMERATION":
			if inferedInstance.ATTRIBUTE_VALUE_ENUMERATION != nil {
				res = inferedInstance.ATTRIBUTE_VALUE_ENUMERATION.Name
			}
		}
	case A_ATTRIBUTE_VALUE_INTEGER:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "ATTRIBUTE_VALUE_INTEGER":
			if inferedInstance.ATTRIBUTE_VALUE_INTEGER != nil {
				res = inferedInstance.ATTRIBUTE_VALUE_INTEGER.Name
			}
		}
	case A_ATTRIBUTE_VALUE_REAL:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "ATTRIBUTE_VALUE_REAL":
			if inferedInstance.ATTRIBUTE_VALUE_REAL != nil {
				res = inferedInstance.ATTRIBUTE_VALUE_REAL.Name
			}
		}
	case A_ATTRIBUTE_VALUE_STRING:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "ATTRIBUTE_VALUE_STRING":
			if inferedInstance.ATTRIBUTE_VALUE_STRING != nil {
				res = inferedInstance.ATTRIBUTE_VALUE_STRING.Name
			}
		}
	case A_ATTRIBUTE_VALUE_XHTML:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "ATTRIBUTE_VALUE_BOOLEAN":
			if inferedInstance.ATTRIBUTE_VALUE_BOOLEAN != nil {
				res = inferedInstance.ATTRIBUTE_VALUE_BOOLEAN.Name
			}
		case "ATTRIBUTE_VALUE_DATE":
			if inferedInstance.ATTRIBUTE_VALUE_DATE != nil {
				res = inferedInstance.ATTRIBUTE_VALUE_DATE.Name
			}
		case "ATTRIBUTE_VALUE_ENUMERATION":
			if inferedInstance.ATTRIBUTE_VALUE_ENUMERATION != nil {
				res = inferedInstance.ATTRIBUTE_VALUE_ENUMERATION.Name
			}
		case "ATTRIBUTE_VALUE_INTEGER":
			if inferedInstance.ATTRIBUTE_VALUE_INTEGER != nil {
				res = inferedInstance.ATTRIBUTE_VALUE_INTEGER.Name
			}
		case "ATTRIBUTE_VALUE_REAL":
			if inferedInstance.ATTRIBUTE_VALUE_REAL != nil {
				res = inferedInstance.ATTRIBUTE_VALUE_REAL.Name
			}
		case "ATTRIBUTE_VALUE_STRING":
			if inferedInstance.ATTRIBUTE_VALUE_STRING != nil {
				res = inferedInstance.ATTRIBUTE_VALUE_STRING.Name
			}
		case "ATTRIBUTE_VALUE_XHTML":
			if inferedInstance.ATTRIBUTE_VALUE_XHTML != nil {
				res = inferedInstance.ATTRIBUTE_VALUE_XHTML.Name
			}
		}
	case A_ATTRIBUTE_VALUE_XHTML_1:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "ATTRIBUTE_VALUE_XHTML":
			if inferedInstance.ATTRIBUTE_VALUE_XHTML != nil {
				res = inferedInstance.ATTRIBUTE_VALUE_XHTML.Name
			}
		}
	case A_CHILDREN:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "SPEC_HIERARCHY":
			if inferedInstance.SPEC_HIERARCHY != nil {
				res = inferedInstance.SPEC_HIERARCHY.Name
			}
		}
	case A_CORE_CONTENT:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "REQ_IF_CONTENT":
			if inferedInstance.REQ_IF_CONTENT != nil {
				res = inferedInstance.REQ_IF_CONTENT.Name
			}
		}
	case A_DATATYPES:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DATATYPE_DEFINITION_BOOLEAN":
			if inferedInstance.DATATYPE_DEFINITION_BOOLEAN != nil {
				res = inferedInstance.DATATYPE_DEFINITION_BOOLEAN.Name
			}
		case "DATATYPE_DEFINITION_DATE":
			if inferedInstance.DATATYPE_DEFINITION_DATE != nil {
				res = inferedInstance.DATATYPE_DEFINITION_DATE.Name
			}
		case "DATATYPE_DEFINITION_ENUMERATION":
			if inferedInstance.DATATYPE_DEFINITION_ENUMERATION != nil {
				res = inferedInstance.DATATYPE_DEFINITION_ENUMERATION.Name
			}
		case "DATATYPE_DEFINITION_INTEGER":
			if inferedInstance.DATATYPE_DEFINITION_INTEGER != nil {
				res = inferedInstance.DATATYPE_DEFINITION_INTEGER.Name
			}
		case "DATATYPE_DEFINITION_REAL":
			if inferedInstance.DATATYPE_DEFINITION_REAL != nil {
				res = inferedInstance.DATATYPE_DEFINITION_REAL.Name
			}
		case "DATATYPE_DEFINITION_STRING":
			if inferedInstance.DATATYPE_DEFINITION_STRING != nil {
				res = inferedInstance.DATATYPE_DEFINITION_STRING.Name
			}
		case "DATATYPE_DEFINITION_XHTML":
			if inferedInstance.DATATYPE_DEFINITION_XHTML != nil {
				res = inferedInstance.DATATYPE_DEFINITION_XHTML.Name
			}
		}
	case A_DATATYPE_DEFINITION_BOOLEAN_REF:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DATATYPE_DEFINITION_BOOLEAN_REF":
			res = inferedInstance.DATATYPE_DEFINITION_BOOLEAN_REF
		}
	case A_DATATYPE_DEFINITION_DATE_REF:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DATATYPE_DEFINITION_DATE_REF":
			res = inferedInstance.DATATYPE_DEFINITION_DATE_REF
		}
	case A_DATATYPE_DEFINITION_ENUMERATION_REF:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DATATYPE_DEFINITION_ENUMERATION_REF":
			res = inferedInstance.DATATYPE_DEFINITION_ENUMERATION_REF
		}
	case A_DATATYPE_DEFINITION_INTEGER_REF:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DATATYPE_DEFINITION_INTEGER_REF":
			res = inferedInstance.DATATYPE_DEFINITION_INTEGER_REF
		}
	case A_DATATYPE_DEFINITION_REAL_REF:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DATATYPE_DEFINITION_REAL_REF":
			res = inferedInstance.DATATYPE_DEFINITION_REAL_REF
		}
	case A_DATATYPE_DEFINITION_STRING_REF:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DATATYPE_DEFINITION_STRING_REF":
			res = inferedInstance.DATATYPE_DEFINITION_STRING_REF
		}
	case A_DATATYPE_DEFINITION_XHTML_REF:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DATATYPE_DEFINITION_XHTML_REF":
			res = inferedInstance.DATATYPE_DEFINITION_XHTML_REF
		}
	case A_EDITABLE_ATTS:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "ATTRIBUTE_DEFINITION_BOOLEAN_REF":
			res = inferedInstance.ATTRIBUTE_DEFINITION_BOOLEAN_REF
		case "ATTRIBUTE_DEFINITION_DATE_REF":
			res = inferedInstance.ATTRIBUTE_DEFINITION_DATE_REF
		case "ATTRIBUTE_DEFINITION_ENUMERATION_REF":
			res = inferedInstance.ATTRIBUTE_DEFINITION_ENUMERATION_REF
		case "ATTRIBUTE_DEFINITION_INTEGER_REF":
			res = inferedInstance.ATTRIBUTE_DEFINITION_INTEGER_REF
		case "ATTRIBUTE_DEFINITION_REAL_REF":
			res = inferedInstance.ATTRIBUTE_DEFINITION_REAL_REF
		case "ATTRIBUTE_DEFINITION_STRING_REF":
			res = inferedInstance.ATTRIBUTE_DEFINITION_STRING_REF
		case "ATTRIBUTE_DEFINITION_XHTML_REF":
			res = inferedInstance.ATTRIBUTE_DEFINITION_XHTML_REF
		}
	case A_ENUM_VALUE_REF:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "ENUM_VALUE_REF":
			res = inferedInstance.ENUM_VALUE_REF
		}
	case A_OBJECT:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "SPEC_OBJECT_REF":
			res = inferedInstance.SPEC_OBJECT_REF
		}
	case A_PROPERTIES:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "EMBEDDED_VALUE":
			if inferedInstance.EMBEDDED_VALUE != nil {
				res = inferedInstance.EMBEDDED_VALUE.Name
			}
		}
	case A_RELATION_GROUP_TYPE_REF:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "RELATION_GROUP_TYPE_REF":
			res = inferedInstance.RELATION_GROUP_TYPE_REF
		}
	case A_SOURCE_1:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "SPEC_OBJECT_REF":
			res = inferedInstance.SPEC_OBJECT_REF
		}
	case A_SPECIFICATIONS:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "SPECIFICATION":
			if inferedInstance.SPECIFICATION != nil {
				res = inferedInstance.SPECIFICATION.Name
			}
		}
	case A_SPECIFICATION_TYPE_REF:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "SPECIFICATION_TYPE_REF":
			res = inferedInstance.SPECIFICATION_TYPE_REF
		}
	case A_SPECIFIED_VALUES:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "ENUM_VALUE":
			if inferedInstance.ENUM_VALUE != nil {
				res = inferedInstance.ENUM_VALUE.Name
			}
		}
	case A_SPEC_ATTRIBUTES:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "ATTRIBUTE_DEFINITION_BOOLEAN":
			if inferedInstance.ATTRIBUTE_DEFINITION_BOOLEAN != nil {
				res = inferedInstance.ATTRIBUTE_DEFINITION_BOOLEAN.Name
			}
		case "ATTRIBUTE_DEFINITION_DATE":
			if inferedInstance.ATTRIBUTE_DEFINITION_DATE != nil {
				res = inferedInstance.ATTRIBUTE_DEFINITION_DATE.Name
			}
		case "ATTRIBUTE_DEFINITION_ENUMERATION":
			if inferedInstance.ATTRIBUTE_DEFINITION_ENUMERATION != nil {
				res = inferedInstance.ATTRIBUTE_DEFINITION_ENUMERATION.Name
			}
		case "ATTRIBUTE_DEFINITION_INTEGER":
			if inferedInstance.ATTRIBUTE_DEFINITION_INTEGER != nil {
				res = inferedInstance.ATTRIBUTE_DEFINITION_INTEGER.Name
			}
		case "ATTRIBUTE_DEFINITION_REAL":
			if inferedInstance.ATTRIBUTE_DEFINITION_REAL != nil {
				res = inferedInstance.ATTRIBUTE_DEFINITION_REAL.Name
			}
		case "ATTRIBUTE_DEFINITION_STRING":
			if inferedInstance.ATTRIBUTE_DEFINITION_STRING != nil {
				res = inferedInstance.ATTRIBUTE_DEFINITION_STRING.Name
			}
		case "ATTRIBUTE_DEFINITION_XHTML":
			if inferedInstance.ATTRIBUTE_DEFINITION_XHTML != nil {
				res = inferedInstance.ATTRIBUTE_DEFINITION_XHTML.Name
			}
		}
	case A_SPEC_OBJECTS:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "SPEC_OBJECT":
			if inferedInstance.SPEC_OBJECT != nil {
				res = inferedInstance.SPEC_OBJECT.Name
			}
		}
	case A_SPEC_OBJECT_TYPE_REF:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "SPEC_OBJECT_TYPE_REF":
			res = inferedInstance.SPEC_OBJECT_TYPE_REF
		}
	case A_SPEC_RELATIONS:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "SPEC_RELATION":
			if inferedInstance.SPEC_RELATION != nil {
				res = inferedInstance.SPEC_RELATION.Name
			}
		}
	case A_SPEC_RELATION_GROUPS:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "RELATION_GROUP":
			if inferedInstance.RELATION_GROUP != nil {
				res = inferedInstance.RELATION_GROUP.Name
			}
		}
	case A_SPEC_RELATION_REF:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "SPEC_RELATION_REF":
			res = inferedInstance.SPEC_RELATION_REF
		}
	case A_SPEC_RELATION_TYPE_REF:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "SPEC_RELATION_TYPE_REF":
			res = inferedInstance.SPEC_RELATION_TYPE_REF
		}
	case A_SPEC_TYPES:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "RELATION_GROUP_TYPE":
			if inferedInstance.RELATION_GROUP_TYPE != nil {
				res = inferedInstance.RELATION_GROUP_TYPE.Name
			}
		case "SPEC_OBJECT_TYPE":
			if inferedInstance.SPEC_OBJECT_TYPE != nil {
				res = inferedInstance.SPEC_OBJECT_TYPE.Name
			}
		case "SPEC_RELATION_TYPE":
			if inferedInstance.SPEC_RELATION_TYPE != nil {
				res = inferedInstance.SPEC_RELATION_TYPE.Name
			}
		case "SPECIFICATION_TYPE":
			if inferedInstance.SPECIFICATION_TYPE != nil {
				res = inferedInstance.SPECIFICATION_TYPE.Name
			}
		}
	case A_TARGET_SPECIFICATION_1:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "SPECIFICATION_REF":
			res = inferedInstance.SPECIFICATION_REF
		}
	case A_THE_HEADER:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "REQ_IF_HEADER":
			if inferedInstance.REQ_IF_HEADER != nil {
				res = inferedInstance.REQ_IF_HEADER.Name
			}
		}
	case A_TOOL_EXTENSIONS:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "REQ_IF_TOOL_EXTENSION":
			if inferedInstance.REQ_IF_TOOL_EXTENSION != nil {
				res = inferedInstance.REQ_IF_TOOL_EXTENSION.Name
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
		case "ALTERNATIVE_ID":
			if inferedInstance.ALTERNATIVE_ID != nil {
				res = inferedInstance.ALTERNATIVE_ID.Name
			}
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
		case "ALTERNATIVE_ID":
			if inferedInstance.ALTERNATIVE_ID != nil {
				res = inferedInstance.ALTERNATIVE_ID.Name
			}
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
		case "ALTERNATIVE_ID":
			if inferedInstance.ALTERNATIVE_ID != nil {
				res = inferedInstance.ALTERNATIVE_ID.Name
			}
		case "SPECIFIED_VALUES":
			if inferedInstance.SPECIFIED_VALUES != nil {
				res = inferedInstance.SPECIFIED_VALUES.Name
			}
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
		case "ALTERNATIVE_ID":
			if inferedInstance.ALTERNATIVE_ID != nil {
				res = inferedInstance.ALTERNATIVE_ID.Name
			}
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
		case "ALTERNATIVE_ID":
			if inferedInstance.ALTERNATIVE_ID != nil {
				res = inferedInstance.ALTERNATIVE_ID.Name
			}
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
		case "ALTERNATIVE_ID":
			if inferedInstance.ALTERNATIVE_ID != nil {
				res = inferedInstance.ALTERNATIVE_ID.Name
			}
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
		case "ALTERNATIVE_ID":
			if inferedInstance.ALTERNATIVE_ID != nil {
				res = inferedInstance.ALTERNATIVE_ID.Name
			}
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
		case "ALTERNATIVE_ID":
			if inferedInstance.ALTERNATIVE_ID != nil {
				res = inferedInstance.ALTERNATIVE_ID.Name
			}
		case "PROPERTIES":
			if inferedInstance.PROPERTIES != nil {
				res = inferedInstance.PROPERTIES.Name
			}
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
		case "ALTERNATIVE_ID":
			if inferedInstance.ALTERNATIVE_ID != nil {
				res = inferedInstance.ALTERNATIVE_ID.Name
			}
		case "SOURCE_SPECIFICATION":
			if inferedInstance.SOURCE_SPECIFICATION != nil {
				res = inferedInstance.SOURCE_SPECIFICATION.Name
			}
		case "SPEC_RELATIONS":
			if inferedInstance.SPEC_RELATIONS != nil {
				res = inferedInstance.SPEC_RELATIONS.Name
			}
		case "TARGET_SPECIFICATION":
			if inferedInstance.TARGET_SPECIFICATION != nil {
				res = inferedInstance.TARGET_SPECIFICATION.Name
			}
		case "TYPE":
			if inferedInstance.TYPE != nil {
				res = inferedInstance.TYPE.Name
			}
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
		case "ALTERNATIVE_ID":
			if inferedInstance.ALTERNATIVE_ID != nil {
				res = inferedInstance.ALTERNATIVE_ID.Name
			}
		case "SPEC_ATTRIBUTES":
			if inferedInstance.SPEC_ATTRIBUTES != nil {
				res = inferedInstance.SPEC_ATTRIBUTES.Name
			}
		}
	case REQ_IF:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Lang":
			res = inferedInstance.Lang
		case "THE_HEADER":
			if inferedInstance.THE_HEADER != nil {
				res = inferedInstance.THE_HEADER.Name
			}
		case "CORE_CONTENT":
			if inferedInstance.CORE_CONTENT != nil {
				res = inferedInstance.CORE_CONTENT.Name
			}
		case "TOOL_EXTENSIONS":
			if inferedInstance.TOOL_EXTENSIONS != nil {
				res = inferedInstance.TOOL_EXTENSIONS.Name
			}
		}
	case REQ_IF_CONTENT:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DATATYPES":
			if inferedInstance.DATATYPES != nil {
				res = inferedInstance.DATATYPES.Name
			}
		case "SPEC_TYPES":
			if inferedInstance.SPEC_TYPES != nil {
				res = inferedInstance.SPEC_TYPES.Name
			}
		case "SPEC_OBJECTS":
			if inferedInstance.SPEC_OBJECTS != nil {
				res = inferedInstance.SPEC_OBJECTS.Name
			}
		case "SPEC_RELATIONS":
			if inferedInstance.SPEC_RELATIONS != nil {
				res = inferedInstance.SPEC_RELATIONS.Name
			}
		case "SPECIFICATIONS":
			if inferedInstance.SPECIFICATIONS != nil {
				res = inferedInstance.SPECIFICATIONS.Name
			}
		case "SPEC_RELATION_GROUPS":
			if inferedInstance.SPEC_RELATION_GROUPS != nil {
				res = inferedInstance.SPEC_RELATION_GROUPS.Name
			}
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
		case "ALTERNATIVE_ID":
			if inferedInstance.ALTERNATIVE_ID != nil {
				res = inferedInstance.ALTERNATIVE_ID.Name
			}
		case "CHILDREN":
			if inferedInstance.CHILDREN != nil {
				res = inferedInstance.CHILDREN.Name
			}
		case "VALUES":
			if inferedInstance.VALUES != nil {
				res = inferedInstance.VALUES.Name
			}
		case "TYPE":
			if inferedInstance.TYPE != nil {
				res = inferedInstance.TYPE.Name
			}
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
		case "ALTERNATIVE_ID":
			if inferedInstance.ALTERNATIVE_ID != nil {
				res = inferedInstance.ALTERNATIVE_ID.Name
			}
		case "SPEC_ATTRIBUTES":
			if inferedInstance.SPEC_ATTRIBUTES != nil {
				res = inferedInstance.SPEC_ATTRIBUTES.Name
			}
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
		case "ALTERNATIVE_ID":
			if inferedInstance.ALTERNATIVE_ID != nil {
				res = inferedInstance.ALTERNATIVE_ID.Name
			}
		case "CHILDREN":
			if inferedInstance.CHILDREN != nil {
				res = inferedInstance.CHILDREN.Name
			}
		case "EDITABLE_ATTS":
			if inferedInstance.EDITABLE_ATTS != nil {
				res = inferedInstance.EDITABLE_ATTS.Name
			}
		case "OBJECT":
			if inferedInstance.OBJECT != nil {
				res = inferedInstance.OBJECT.Name
			}
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
		case "ALTERNATIVE_ID":
			if inferedInstance.ALTERNATIVE_ID != nil {
				res = inferedInstance.ALTERNATIVE_ID.Name
			}
		case "VALUES":
			if inferedInstance.VALUES != nil {
				res = inferedInstance.VALUES.Name
			}
		case "TYPE":
			if inferedInstance.TYPE != nil {
				res = inferedInstance.TYPE.Name
			}
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
		case "ALTERNATIVE_ID":
			if inferedInstance.ALTERNATIVE_ID != nil {
				res = inferedInstance.ALTERNATIVE_ID.Name
			}
		case "SPEC_ATTRIBUTES":
			if inferedInstance.SPEC_ATTRIBUTES != nil {
				res = inferedInstance.SPEC_ATTRIBUTES.Name
			}
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
		case "ALTERNATIVE_ID":
			if inferedInstance.ALTERNATIVE_ID != nil {
				res = inferedInstance.ALTERNATIVE_ID.Name
			}
		case "VALUES":
			if inferedInstance.VALUES != nil {
				res = inferedInstance.VALUES.Name
			}
		case "SOURCE":
			if inferedInstance.SOURCE != nil {
				res = inferedInstance.SOURCE.Name
			}
		case "TARGET":
			if inferedInstance.TARGET != nil {
				res = inferedInstance.TARGET.Name
			}
		case "TYPE":
			if inferedInstance.TYPE != nil {
				res = inferedInstance.TYPE.Name
			}
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
		case "ALTERNATIVE_ID":
			if inferedInstance.ALTERNATIVE_ID != nil {
				res = inferedInstance.ALTERNATIVE_ID.Name
			}
		case "SPEC_ATTRIBUTES":
			if inferedInstance.SPEC_ATTRIBUTES != nil {
				res = inferedInstance.SPEC_ATTRIBUTES.Name
			}
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
