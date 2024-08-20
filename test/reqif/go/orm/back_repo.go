// do not modify, generated file
package orm

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/fullstack-lang/gongxsd/test/reqif/go/models"

	"github.com/tealeg/xlsx/v3"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// BackRepoStruct supports callback functions
type BackRepoStruct struct {
	// insertion point for per struct back repo declarations
	BackRepoALTERNATIVE_ID BackRepoALTERNATIVE_IDStruct

	BackRepoATTRIBUTE_DEFINITION_BOOLEAN BackRepoATTRIBUTE_DEFINITION_BOOLEANStruct

	BackRepoATTRIBUTE_DEFINITION_DATE BackRepoATTRIBUTE_DEFINITION_DATEStruct

	BackRepoATTRIBUTE_DEFINITION_ENUMERATION BackRepoATTRIBUTE_DEFINITION_ENUMERATIONStruct

	BackRepoATTRIBUTE_DEFINITION_INTEGER BackRepoATTRIBUTE_DEFINITION_INTEGERStruct

	BackRepoATTRIBUTE_DEFINITION_REAL BackRepoATTRIBUTE_DEFINITION_REALStruct

	BackRepoATTRIBUTE_DEFINITION_STRING BackRepoATTRIBUTE_DEFINITION_STRINGStruct

	BackRepoATTRIBUTE_DEFINITION_XHTML BackRepoATTRIBUTE_DEFINITION_XHTMLStruct

	BackRepoATTRIBUTE_VALUE_BOOLEAN BackRepoATTRIBUTE_VALUE_BOOLEANStruct

	BackRepoATTRIBUTE_VALUE_DATE BackRepoATTRIBUTE_VALUE_DATEStruct

	BackRepoATTRIBUTE_VALUE_ENUMERATION BackRepoATTRIBUTE_VALUE_ENUMERATIONStruct

	BackRepoATTRIBUTE_VALUE_INTEGER BackRepoATTRIBUTE_VALUE_INTEGERStruct

	BackRepoATTRIBUTE_VALUE_REAL BackRepoATTRIBUTE_VALUE_REALStruct

	BackRepoATTRIBUTE_VALUE_STRING BackRepoATTRIBUTE_VALUE_STRINGStruct

	BackRepoATTRIBUTE_VALUE_XHTML BackRepoATTRIBUTE_VALUE_XHTMLStruct

	BackRepoDATATYPE_DEFINITION_BOOLEAN BackRepoDATATYPE_DEFINITION_BOOLEANStruct

	BackRepoDATATYPE_DEFINITION_DATE BackRepoDATATYPE_DEFINITION_DATEStruct

	BackRepoDATATYPE_DEFINITION_ENUMERATION BackRepoDATATYPE_DEFINITION_ENUMERATIONStruct

	BackRepoDATATYPE_DEFINITION_INTEGER BackRepoDATATYPE_DEFINITION_INTEGERStruct

	BackRepoDATATYPE_DEFINITION_REAL BackRepoDATATYPE_DEFINITION_REALStruct

	BackRepoDATATYPE_DEFINITION_STRING BackRepoDATATYPE_DEFINITION_STRINGStruct

	BackRepoDATATYPE_DEFINITION_XHTML BackRepoDATATYPE_DEFINITION_XHTMLStruct

	BackRepoEMBEDDED_VALUE BackRepoEMBEDDED_VALUEStruct

	BackRepoENUM_VALUE BackRepoENUM_VALUEStruct

	BackRepoRELATION_GROUP BackRepoRELATION_GROUPStruct

	BackRepoRELATION_GROUP_TYPE BackRepoRELATION_GROUP_TYPEStruct

	BackRepoREQ_IF BackRepoREQ_IFStruct

	BackRepoREQ_IF_CONTENT BackRepoREQ_IF_CONTENTStruct

	BackRepoREQ_IF_HEADER BackRepoREQ_IF_HEADERStruct

	BackRepoREQ_IF_TOOL_EXTENSION BackRepoREQ_IF_TOOL_EXTENSIONStruct

	BackRepoSPECIFICATION BackRepoSPECIFICATIONStruct

	BackRepoSPECIFICATION_TYPE BackRepoSPECIFICATION_TYPEStruct

	BackRepoSPEC_HIERARCHY BackRepoSPEC_HIERARCHYStruct

	BackRepoSPEC_OBJECT BackRepoSPEC_OBJECTStruct

	BackRepoSPEC_OBJECT_TYPE BackRepoSPEC_OBJECT_TYPEStruct

	BackRepoSPEC_RELATION BackRepoSPEC_RELATIONStruct

	BackRepoSPEC_RELATION_TYPE BackRepoSPEC_RELATION_TYPEStruct

	BackRepoXHTML_CONTENT BackRepoXHTML_CONTENTStruct

	CommitFromBackNb uint // records commit increments when performed by the back

	PushFromFrontNb uint // records commit increments when performed by the front

	stage *models.StageStruct

	// the back repo can broadcast the CommitFromBackNb to all interested subscribers
	rwMutex     sync.RWMutex
	subscribers []chan int
}

func NewBackRepo(stage *models.StageStruct, filename string) (backRepo *BackRepoStruct) {

	// adjust naming strategy to the stack
	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "github_com_fullstack_lang_gong_test_go_", // table name prefix
		},
	}
	db, err := gorm.Open(sqlite.Open(filename), gormConfig)

	// since testsim is a multi threaded application. It is important to set up
	// only one open connexion at a time
	dbDB_inMemory, err := db.DB()
	if err != nil {
		panic("cannot access DB of db" + err.Error())
	}
	// it is mandatory to allow parallel access, otherwise, bizarre errors occurs
	dbDB_inMemory.SetMaxOpenConns(1)

	if err != nil {
		panic("Failed to connect to database!")
	}

	// adjust naming strategy to the stack
	db.Config.NamingStrategy = &schema.NamingStrategy{
		TablePrefix: "github_com_fullstack_lang_gong_test_go_", // table name prefix
	}

	err = db.AutoMigrate( // insertion point for reference to structs
		&ALTERNATIVE_IDDB{},
		&ATTRIBUTE_DEFINITION_BOOLEANDB{},
		&ATTRIBUTE_DEFINITION_DATEDB{},
		&ATTRIBUTE_DEFINITION_ENUMERATIONDB{},
		&ATTRIBUTE_DEFINITION_INTEGERDB{},
		&ATTRIBUTE_DEFINITION_REALDB{},
		&ATTRIBUTE_DEFINITION_STRINGDB{},
		&ATTRIBUTE_DEFINITION_XHTMLDB{},
		&ATTRIBUTE_VALUE_BOOLEANDB{},
		&ATTRIBUTE_VALUE_DATEDB{},
		&ATTRIBUTE_VALUE_ENUMERATIONDB{},
		&ATTRIBUTE_VALUE_INTEGERDB{},
		&ATTRIBUTE_VALUE_REALDB{},
		&ATTRIBUTE_VALUE_STRINGDB{},
		&ATTRIBUTE_VALUE_XHTMLDB{},
		&DATATYPE_DEFINITION_BOOLEANDB{},
		&DATATYPE_DEFINITION_DATEDB{},
		&DATATYPE_DEFINITION_ENUMERATIONDB{},
		&DATATYPE_DEFINITION_INTEGERDB{},
		&DATATYPE_DEFINITION_REALDB{},
		&DATATYPE_DEFINITION_STRINGDB{},
		&DATATYPE_DEFINITION_XHTMLDB{},
		&EMBEDDED_VALUEDB{},
		&ENUM_VALUEDB{},
		&RELATION_GROUPDB{},
		&RELATION_GROUP_TYPEDB{},
		&REQ_IFDB{},
		&REQ_IF_CONTENTDB{},
		&REQ_IF_HEADERDB{},
		&REQ_IF_TOOL_EXTENSIONDB{},
		&SPECIFICATIONDB{},
		&SPECIFICATION_TYPEDB{},
		&SPEC_HIERARCHYDB{},
		&SPEC_OBJECTDB{},
		&SPEC_OBJECT_TYPEDB{},
		&SPEC_RELATIONDB{},
		&SPEC_RELATION_TYPEDB{},
		&XHTML_CONTENTDB{},
	)

	if err != nil {
		msg := err.Error()
		panic("problem with migration " + msg + " on package github.com/fullstack-lang/gong/test/go")
	}

	backRepo = new(BackRepoStruct)

	// insertion point for per struct back repo declarations
	backRepo.BackRepoALTERNATIVE_ID = BackRepoALTERNATIVE_IDStruct{
		Map_ALTERNATIVE_IDDBID_ALTERNATIVE_IDPtr: make(map[uint]*models.ALTERNATIVE_ID, 0),
		Map_ALTERNATIVE_IDDBID_ALTERNATIVE_IDDB:  make(map[uint]*ALTERNATIVE_IDDB, 0),
		Map_ALTERNATIVE_IDPtr_ALTERNATIVE_IDDBID: make(map[*models.ALTERNATIVE_ID]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoATTRIBUTE_DEFINITION_BOOLEAN = BackRepoATTRIBUTE_DEFINITION_BOOLEANStruct{
		Map_ATTRIBUTE_DEFINITION_BOOLEANDBID_ATTRIBUTE_DEFINITION_BOOLEANPtr: make(map[uint]*models.ATTRIBUTE_DEFINITION_BOOLEAN, 0),
		Map_ATTRIBUTE_DEFINITION_BOOLEANDBID_ATTRIBUTE_DEFINITION_BOOLEANDB:  make(map[uint]*ATTRIBUTE_DEFINITION_BOOLEANDB, 0),
		Map_ATTRIBUTE_DEFINITION_BOOLEANPtr_ATTRIBUTE_DEFINITION_BOOLEANDBID: make(map[*models.ATTRIBUTE_DEFINITION_BOOLEAN]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoATTRIBUTE_DEFINITION_DATE = BackRepoATTRIBUTE_DEFINITION_DATEStruct{
		Map_ATTRIBUTE_DEFINITION_DATEDBID_ATTRIBUTE_DEFINITION_DATEPtr: make(map[uint]*models.ATTRIBUTE_DEFINITION_DATE, 0),
		Map_ATTRIBUTE_DEFINITION_DATEDBID_ATTRIBUTE_DEFINITION_DATEDB:  make(map[uint]*ATTRIBUTE_DEFINITION_DATEDB, 0),
		Map_ATTRIBUTE_DEFINITION_DATEPtr_ATTRIBUTE_DEFINITION_DATEDBID: make(map[*models.ATTRIBUTE_DEFINITION_DATE]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoATTRIBUTE_DEFINITION_ENUMERATION = BackRepoATTRIBUTE_DEFINITION_ENUMERATIONStruct{
		Map_ATTRIBUTE_DEFINITION_ENUMERATIONDBID_ATTRIBUTE_DEFINITION_ENUMERATIONPtr: make(map[uint]*models.ATTRIBUTE_DEFINITION_ENUMERATION, 0),
		Map_ATTRIBUTE_DEFINITION_ENUMERATIONDBID_ATTRIBUTE_DEFINITION_ENUMERATIONDB:  make(map[uint]*ATTRIBUTE_DEFINITION_ENUMERATIONDB, 0),
		Map_ATTRIBUTE_DEFINITION_ENUMERATIONPtr_ATTRIBUTE_DEFINITION_ENUMERATIONDBID: make(map[*models.ATTRIBUTE_DEFINITION_ENUMERATION]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoATTRIBUTE_DEFINITION_INTEGER = BackRepoATTRIBUTE_DEFINITION_INTEGERStruct{
		Map_ATTRIBUTE_DEFINITION_INTEGERDBID_ATTRIBUTE_DEFINITION_INTEGERPtr: make(map[uint]*models.ATTRIBUTE_DEFINITION_INTEGER, 0),
		Map_ATTRIBUTE_DEFINITION_INTEGERDBID_ATTRIBUTE_DEFINITION_INTEGERDB:  make(map[uint]*ATTRIBUTE_DEFINITION_INTEGERDB, 0),
		Map_ATTRIBUTE_DEFINITION_INTEGERPtr_ATTRIBUTE_DEFINITION_INTEGERDBID: make(map[*models.ATTRIBUTE_DEFINITION_INTEGER]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoATTRIBUTE_DEFINITION_REAL = BackRepoATTRIBUTE_DEFINITION_REALStruct{
		Map_ATTRIBUTE_DEFINITION_REALDBID_ATTRIBUTE_DEFINITION_REALPtr: make(map[uint]*models.ATTRIBUTE_DEFINITION_REAL, 0),
		Map_ATTRIBUTE_DEFINITION_REALDBID_ATTRIBUTE_DEFINITION_REALDB:  make(map[uint]*ATTRIBUTE_DEFINITION_REALDB, 0),
		Map_ATTRIBUTE_DEFINITION_REALPtr_ATTRIBUTE_DEFINITION_REALDBID: make(map[*models.ATTRIBUTE_DEFINITION_REAL]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoATTRIBUTE_DEFINITION_STRING = BackRepoATTRIBUTE_DEFINITION_STRINGStruct{
		Map_ATTRIBUTE_DEFINITION_STRINGDBID_ATTRIBUTE_DEFINITION_STRINGPtr: make(map[uint]*models.ATTRIBUTE_DEFINITION_STRING, 0),
		Map_ATTRIBUTE_DEFINITION_STRINGDBID_ATTRIBUTE_DEFINITION_STRINGDB:  make(map[uint]*ATTRIBUTE_DEFINITION_STRINGDB, 0),
		Map_ATTRIBUTE_DEFINITION_STRINGPtr_ATTRIBUTE_DEFINITION_STRINGDBID: make(map[*models.ATTRIBUTE_DEFINITION_STRING]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoATTRIBUTE_DEFINITION_XHTML = BackRepoATTRIBUTE_DEFINITION_XHTMLStruct{
		Map_ATTRIBUTE_DEFINITION_XHTMLDBID_ATTRIBUTE_DEFINITION_XHTMLPtr: make(map[uint]*models.ATTRIBUTE_DEFINITION_XHTML, 0),
		Map_ATTRIBUTE_DEFINITION_XHTMLDBID_ATTRIBUTE_DEFINITION_XHTMLDB:  make(map[uint]*ATTRIBUTE_DEFINITION_XHTMLDB, 0),
		Map_ATTRIBUTE_DEFINITION_XHTMLPtr_ATTRIBUTE_DEFINITION_XHTMLDBID: make(map[*models.ATTRIBUTE_DEFINITION_XHTML]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoATTRIBUTE_VALUE_BOOLEAN = BackRepoATTRIBUTE_VALUE_BOOLEANStruct{
		Map_ATTRIBUTE_VALUE_BOOLEANDBID_ATTRIBUTE_VALUE_BOOLEANPtr: make(map[uint]*models.ATTRIBUTE_VALUE_BOOLEAN, 0),
		Map_ATTRIBUTE_VALUE_BOOLEANDBID_ATTRIBUTE_VALUE_BOOLEANDB:  make(map[uint]*ATTRIBUTE_VALUE_BOOLEANDB, 0),
		Map_ATTRIBUTE_VALUE_BOOLEANPtr_ATTRIBUTE_VALUE_BOOLEANDBID: make(map[*models.ATTRIBUTE_VALUE_BOOLEAN]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoATTRIBUTE_VALUE_DATE = BackRepoATTRIBUTE_VALUE_DATEStruct{
		Map_ATTRIBUTE_VALUE_DATEDBID_ATTRIBUTE_VALUE_DATEPtr: make(map[uint]*models.ATTRIBUTE_VALUE_DATE, 0),
		Map_ATTRIBUTE_VALUE_DATEDBID_ATTRIBUTE_VALUE_DATEDB:  make(map[uint]*ATTRIBUTE_VALUE_DATEDB, 0),
		Map_ATTRIBUTE_VALUE_DATEPtr_ATTRIBUTE_VALUE_DATEDBID: make(map[*models.ATTRIBUTE_VALUE_DATE]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoATTRIBUTE_VALUE_ENUMERATION = BackRepoATTRIBUTE_VALUE_ENUMERATIONStruct{
		Map_ATTRIBUTE_VALUE_ENUMERATIONDBID_ATTRIBUTE_VALUE_ENUMERATIONPtr: make(map[uint]*models.ATTRIBUTE_VALUE_ENUMERATION, 0),
		Map_ATTRIBUTE_VALUE_ENUMERATIONDBID_ATTRIBUTE_VALUE_ENUMERATIONDB:  make(map[uint]*ATTRIBUTE_VALUE_ENUMERATIONDB, 0),
		Map_ATTRIBUTE_VALUE_ENUMERATIONPtr_ATTRIBUTE_VALUE_ENUMERATIONDBID: make(map[*models.ATTRIBUTE_VALUE_ENUMERATION]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoATTRIBUTE_VALUE_INTEGER = BackRepoATTRIBUTE_VALUE_INTEGERStruct{
		Map_ATTRIBUTE_VALUE_INTEGERDBID_ATTRIBUTE_VALUE_INTEGERPtr: make(map[uint]*models.ATTRIBUTE_VALUE_INTEGER, 0),
		Map_ATTRIBUTE_VALUE_INTEGERDBID_ATTRIBUTE_VALUE_INTEGERDB:  make(map[uint]*ATTRIBUTE_VALUE_INTEGERDB, 0),
		Map_ATTRIBUTE_VALUE_INTEGERPtr_ATTRIBUTE_VALUE_INTEGERDBID: make(map[*models.ATTRIBUTE_VALUE_INTEGER]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoATTRIBUTE_VALUE_REAL = BackRepoATTRIBUTE_VALUE_REALStruct{
		Map_ATTRIBUTE_VALUE_REALDBID_ATTRIBUTE_VALUE_REALPtr: make(map[uint]*models.ATTRIBUTE_VALUE_REAL, 0),
		Map_ATTRIBUTE_VALUE_REALDBID_ATTRIBUTE_VALUE_REALDB:  make(map[uint]*ATTRIBUTE_VALUE_REALDB, 0),
		Map_ATTRIBUTE_VALUE_REALPtr_ATTRIBUTE_VALUE_REALDBID: make(map[*models.ATTRIBUTE_VALUE_REAL]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoATTRIBUTE_VALUE_STRING = BackRepoATTRIBUTE_VALUE_STRINGStruct{
		Map_ATTRIBUTE_VALUE_STRINGDBID_ATTRIBUTE_VALUE_STRINGPtr: make(map[uint]*models.ATTRIBUTE_VALUE_STRING, 0),
		Map_ATTRIBUTE_VALUE_STRINGDBID_ATTRIBUTE_VALUE_STRINGDB:  make(map[uint]*ATTRIBUTE_VALUE_STRINGDB, 0),
		Map_ATTRIBUTE_VALUE_STRINGPtr_ATTRIBUTE_VALUE_STRINGDBID: make(map[*models.ATTRIBUTE_VALUE_STRING]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoATTRIBUTE_VALUE_XHTML = BackRepoATTRIBUTE_VALUE_XHTMLStruct{
		Map_ATTRIBUTE_VALUE_XHTMLDBID_ATTRIBUTE_VALUE_XHTMLPtr: make(map[uint]*models.ATTRIBUTE_VALUE_XHTML, 0),
		Map_ATTRIBUTE_VALUE_XHTMLDBID_ATTRIBUTE_VALUE_XHTMLDB:  make(map[uint]*ATTRIBUTE_VALUE_XHTMLDB, 0),
		Map_ATTRIBUTE_VALUE_XHTMLPtr_ATTRIBUTE_VALUE_XHTMLDBID: make(map[*models.ATTRIBUTE_VALUE_XHTML]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoDATATYPE_DEFINITION_BOOLEAN = BackRepoDATATYPE_DEFINITION_BOOLEANStruct{
		Map_DATATYPE_DEFINITION_BOOLEANDBID_DATATYPE_DEFINITION_BOOLEANPtr: make(map[uint]*models.DATATYPE_DEFINITION_BOOLEAN, 0),
		Map_DATATYPE_DEFINITION_BOOLEANDBID_DATATYPE_DEFINITION_BOOLEANDB:  make(map[uint]*DATATYPE_DEFINITION_BOOLEANDB, 0),
		Map_DATATYPE_DEFINITION_BOOLEANPtr_DATATYPE_DEFINITION_BOOLEANDBID: make(map[*models.DATATYPE_DEFINITION_BOOLEAN]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoDATATYPE_DEFINITION_DATE = BackRepoDATATYPE_DEFINITION_DATEStruct{
		Map_DATATYPE_DEFINITION_DATEDBID_DATATYPE_DEFINITION_DATEPtr: make(map[uint]*models.DATATYPE_DEFINITION_DATE, 0),
		Map_DATATYPE_DEFINITION_DATEDBID_DATATYPE_DEFINITION_DATEDB:  make(map[uint]*DATATYPE_DEFINITION_DATEDB, 0),
		Map_DATATYPE_DEFINITION_DATEPtr_DATATYPE_DEFINITION_DATEDBID: make(map[*models.DATATYPE_DEFINITION_DATE]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoDATATYPE_DEFINITION_ENUMERATION = BackRepoDATATYPE_DEFINITION_ENUMERATIONStruct{
		Map_DATATYPE_DEFINITION_ENUMERATIONDBID_DATATYPE_DEFINITION_ENUMERATIONPtr: make(map[uint]*models.DATATYPE_DEFINITION_ENUMERATION, 0),
		Map_DATATYPE_DEFINITION_ENUMERATIONDBID_DATATYPE_DEFINITION_ENUMERATIONDB:  make(map[uint]*DATATYPE_DEFINITION_ENUMERATIONDB, 0),
		Map_DATATYPE_DEFINITION_ENUMERATIONPtr_DATATYPE_DEFINITION_ENUMERATIONDBID: make(map[*models.DATATYPE_DEFINITION_ENUMERATION]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoDATATYPE_DEFINITION_INTEGER = BackRepoDATATYPE_DEFINITION_INTEGERStruct{
		Map_DATATYPE_DEFINITION_INTEGERDBID_DATATYPE_DEFINITION_INTEGERPtr: make(map[uint]*models.DATATYPE_DEFINITION_INTEGER, 0),
		Map_DATATYPE_DEFINITION_INTEGERDBID_DATATYPE_DEFINITION_INTEGERDB:  make(map[uint]*DATATYPE_DEFINITION_INTEGERDB, 0),
		Map_DATATYPE_DEFINITION_INTEGERPtr_DATATYPE_DEFINITION_INTEGERDBID: make(map[*models.DATATYPE_DEFINITION_INTEGER]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoDATATYPE_DEFINITION_REAL = BackRepoDATATYPE_DEFINITION_REALStruct{
		Map_DATATYPE_DEFINITION_REALDBID_DATATYPE_DEFINITION_REALPtr: make(map[uint]*models.DATATYPE_DEFINITION_REAL, 0),
		Map_DATATYPE_DEFINITION_REALDBID_DATATYPE_DEFINITION_REALDB:  make(map[uint]*DATATYPE_DEFINITION_REALDB, 0),
		Map_DATATYPE_DEFINITION_REALPtr_DATATYPE_DEFINITION_REALDBID: make(map[*models.DATATYPE_DEFINITION_REAL]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoDATATYPE_DEFINITION_STRING = BackRepoDATATYPE_DEFINITION_STRINGStruct{
		Map_DATATYPE_DEFINITION_STRINGDBID_DATATYPE_DEFINITION_STRINGPtr: make(map[uint]*models.DATATYPE_DEFINITION_STRING, 0),
		Map_DATATYPE_DEFINITION_STRINGDBID_DATATYPE_DEFINITION_STRINGDB:  make(map[uint]*DATATYPE_DEFINITION_STRINGDB, 0),
		Map_DATATYPE_DEFINITION_STRINGPtr_DATATYPE_DEFINITION_STRINGDBID: make(map[*models.DATATYPE_DEFINITION_STRING]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoDATATYPE_DEFINITION_XHTML = BackRepoDATATYPE_DEFINITION_XHTMLStruct{
		Map_DATATYPE_DEFINITION_XHTMLDBID_DATATYPE_DEFINITION_XHTMLPtr: make(map[uint]*models.DATATYPE_DEFINITION_XHTML, 0),
		Map_DATATYPE_DEFINITION_XHTMLDBID_DATATYPE_DEFINITION_XHTMLDB:  make(map[uint]*DATATYPE_DEFINITION_XHTMLDB, 0),
		Map_DATATYPE_DEFINITION_XHTMLPtr_DATATYPE_DEFINITION_XHTMLDBID: make(map[*models.DATATYPE_DEFINITION_XHTML]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoEMBEDDED_VALUE = BackRepoEMBEDDED_VALUEStruct{
		Map_EMBEDDED_VALUEDBID_EMBEDDED_VALUEPtr: make(map[uint]*models.EMBEDDED_VALUE, 0),
		Map_EMBEDDED_VALUEDBID_EMBEDDED_VALUEDB:  make(map[uint]*EMBEDDED_VALUEDB, 0),
		Map_EMBEDDED_VALUEPtr_EMBEDDED_VALUEDBID: make(map[*models.EMBEDDED_VALUE]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoENUM_VALUE = BackRepoENUM_VALUEStruct{
		Map_ENUM_VALUEDBID_ENUM_VALUEPtr: make(map[uint]*models.ENUM_VALUE, 0),
		Map_ENUM_VALUEDBID_ENUM_VALUEDB:  make(map[uint]*ENUM_VALUEDB, 0),
		Map_ENUM_VALUEPtr_ENUM_VALUEDBID: make(map[*models.ENUM_VALUE]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoRELATION_GROUP = BackRepoRELATION_GROUPStruct{
		Map_RELATION_GROUPDBID_RELATION_GROUPPtr: make(map[uint]*models.RELATION_GROUP, 0),
		Map_RELATION_GROUPDBID_RELATION_GROUPDB:  make(map[uint]*RELATION_GROUPDB, 0),
		Map_RELATION_GROUPPtr_RELATION_GROUPDBID: make(map[*models.RELATION_GROUP]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoRELATION_GROUP_TYPE = BackRepoRELATION_GROUP_TYPEStruct{
		Map_RELATION_GROUP_TYPEDBID_RELATION_GROUP_TYPEPtr: make(map[uint]*models.RELATION_GROUP_TYPE, 0),
		Map_RELATION_GROUP_TYPEDBID_RELATION_GROUP_TYPEDB:  make(map[uint]*RELATION_GROUP_TYPEDB, 0),
		Map_RELATION_GROUP_TYPEPtr_RELATION_GROUP_TYPEDBID: make(map[*models.RELATION_GROUP_TYPE]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoREQ_IF = BackRepoREQ_IFStruct{
		Map_REQ_IFDBID_REQ_IFPtr: make(map[uint]*models.REQ_IF, 0),
		Map_REQ_IFDBID_REQ_IFDB:  make(map[uint]*REQ_IFDB, 0),
		Map_REQ_IFPtr_REQ_IFDBID: make(map[*models.REQ_IF]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoREQ_IF_CONTENT = BackRepoREQ_IF_CONTENTStruct{
		Map_REQ_IF_CONTENTDBID_REQ_IF_CONTENTPtr: make(map[uint]*models.REQ_IF_CONTENT, 0),
		Map_REQ_IF_CONTENTDBID_REQ_IF_CONTENTDB:  make(map[uint]*REQ_IF_CONTENTDB, 0),
		Map_REQ_IF_CONTENTPtr_REQ_IF_CONTENTDBID: make(map[*models.REQ_IF_CONTENT]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoREQ_IF_HEADER = BackRepoREQ_IF_HEADERStruct{
		Map_REQ_IF_HEADERDBID_REQ_IF_HEADERPtr: make(map[uint]*models.REQ_IF_HEADER, 0),
		Map_REQ_IF_HEADERDBID_REQ_IF_HEADERDB:  make(map[uint]*REQ_IF_HEADERDB, 0),
		Map_REQ_IF_HEADERPtr_REQ_IF_HEADERDBID: make(map[*models.REQ_IF_HEADER]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoREQ_IF_TOOL_EXTENSION = BackRepoREQ_IF_TOOL_EXTENSIONStruct{
		Map_REQ_IF_TOOL_EXTENSIONDBID_REQ_IF_TOOL_EXTENSIONPtr: make(map[uint]*models.REQ_IF_TOOL_EXTENSION, 0),
		Map_REQ_IF_TOOL_EXTENSIONDBID_REQ_IF_TOOL_EXTENSIONDB:  make(map[uint]*REQ_IF_TOOL_EXTENSIONDB, 0),
		Map_REQ_IF_TOOL_EXTENSIONPtr_REQ_IF_TOOL_EXTENSIONDBID: make(map[*models.REQ_IF_TOOL_EXTENSION]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoSPECIFICATION = BackRepoSPECIFICATIONStruct{
		Map_SPECIFICATIONDBID_SPECIFICATIONPtr: make(map[uint]*models.SPECIFICATION, 0),
		Map_SPECIFICATIONDBID_SPECIFICATIONDB:  make(map[uint]*SPECIFICATIONDB, 0),
		Map_SPECIFICATIONPtr_SPECIFICATIONDBID: make(map[*models.SPECIFICATION]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoSPECIFICATION_TYPE = BackRepoSPECIFICATION_TYPEStruct{
		Map_SPECIFICATION_TYPEDBID_SPECIFICATION_TYPEPtr: make(map[uint]*models.SPECIFICATION_TYPE, 0),
		Map_SPECIFICATION_TYPEDBID_SPECIFICATION_TYPEDB:  make(map[uint]*SPECIFICATION_TYPEDB, 0),
		Map_SPECIFICATION_TYPEPtr_SPECIFICATION_TYPEDBID: make(map[*models.SPECIFICATION_TYPE]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoSPEC_HIERARCHY = BackRepoSPEC_HIERARCHYStruct{
		Map_SPEC_HIERARCHYDBID_SPEC_HIERARCHYPtr: make(map[uint]*models.SPEC_HIERARCHY, 0),
		Map_SPEC_HIERARCHYDBID_SPEC_HIERARCHYDB:  make(map[uint]*SPEC_HIERARCHYDB, 0),
		Map_SPEC_HIERARCHYPtr_SPEC_HIERARCHYDBID: make(map[*models.SPEC_HIERARCHY]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoSPEC_OBJECT = BackRepoSPEC_OBJECTStruct{
		Map_SPEC_OBJECTDBID_SPEC_OBJECTPtr: make(map[uint]*models.SPEC_OBJECT, 0),
		Map_SPEC_OBJECTDBID_SPEC_OBJECTDB:  make(map[uint]*SPEC_OBJECTDB, 0),
		Map_SPEC_OBJECTPtr_SPEC_OBJECTDBID: make(map[*models.SPEC_OBJECT]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoSPEC_OBJECT_TYPE = BackRepoSPEC_OBJECT_TYPEStruct{
		Map_SPEC_OBJECT_TYPEDBID_SPEC_OBJECT_TYPEPtr: make(map[uint]*models.SPEC_OBJECT_TYPE, 0),
		Map_SPEC_OBJECT_TYPEDBID_SPEC_OBJECT_TYPEDB:  make(map[uint]*SPEC_OBJECT_TYPEDB, 0),
		Map_SPEC_OBJECT_TYPEPtr_SPEC_OBJECT_TYPEDBID: make(map[*models.SPEC_OBJECT_TYPE]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoSPEC_RELATION = BackRepoSPEC_RELATIONStruct{
		Map_SPEC_RELATIONDBID_SPEC_RELATIONPtr: make(map[uint]*models.SPEC_RELATION, 0),
		Map_SPEC_RELATIONDBID_SPEC_RELATIONDB:  make(map[uint]*SPEC_RELATIONDB, 0),
		Map_SPEC_RELATIONPtr_SPEC_RELATIONDBID: make(map[*models.SPEC_RELATION]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoSPEC_RELATION_TYPE = BackRepoSPEC_RELATION_TYPEStruct{
		Map_SPEC_RELATION_TYPEDBID_SPEC_RELATION_TYPEPtr: make(map[uint]*models.SPEC_RELATION_TYPE, 0),
		Map_SPEC_RELATION_TYPEDBID_SPEC_RELATION_TYPEDB:  make(map[uint]*SPEC_RELATION_TYPEDB, 0),
		Map_SPEC_RELATION_TYPEPtr_SPEC_RELATION_TYPEDBID: make(map[*models.SPEC_RELATION_TYPE]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoXHTML_CONTENT = BackRepoXHTML_CONTENTStruct{
		Map_XHTML_CONTENTDBID_XHTML_CONTENTPtr: make(map[uint]*models.XHTML_CONTENT, 0),
		Map_XHTML_CONTENTDBID_XHTML_CONTENTDB:  make(map[uint]*XHTML_CONTENTDB, 0),
		Map_XHTML_CONTENTPtr_XHTML_CONTENTDBID: make(map[*models.XHTML_CONTENT]uint, 0),

		db:    db,
		stage: stage,
	}

	stage.BackRepo = backRepo
	backRepo.stage = stage

	return
}

func (backRepo *BackRepoStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepo.stage
	return
}

func (backRepo *BackRepoStruct) GetLastCommitFromBackNb() uint {
	return backRepo.CommitFromBackNb
}

func (backRepo *BackRepoStruct) GetLastPushFromFrontNb() uint {
	return backRepo.PushFromFrontNb
}

func (backRepo *BackRepoStruct) IncrementCommitFromBackNb() uint {
	if backRepo.stage.OnInitCommitCallback != nil {
		backRepo.stage.OnInitCommitCallback.BeforeCommit(backRepo.stage)
	}
	if backRepo.stage.OnInitCommitFromBackCallback != nil {
		backRepo.stage.OnInitCommitFromBackCallback.BeforeCommit(backRepo.stage)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1

	backRepo.broadcastNbCommitToBack()
	
	return backRepo.CommitFromBackNb
}

func (backRepo *BackRepoStruct) IncrementPushFromFrontNb() uint {
	if backRepo.stage.OnInitCommitCallback != nil {
		backRepo.stage.OnInitCommitCallback.BeforeCommit(backRepo.stage)
	}
	if backRepo.stage.OnInitCommitFromFrontCallback != nil {
		backRepo.stage.OnInitCommitFromFrontCallback.BeforeCommit(backRepo.stage)
	}
	backRepo.PushFromFrontNb = backRepo.PushFromFrontNb + 1
	return backRepo.CommitFromBackNb
}

// Commit the BackRepoStruct inner variables and link to the database
func (backRepo *BackRepoStruct) Commit(stage *models.StageStruct) {
	// insertion point for per struct back repo phase one commit
	backRepo.BackRepoALTERNATIVE_ID.CommitPhaseOne(stage)
	backRepo.BackRepoATTRIBUTE_DEFINITION_BOOLEAN.CommitPhaseOne(stage)
	backRepo.BackRepoATTRIBUTE_DEFINITION_DATE.CommitPhaseOne(stage)
	backRepo.BackRepoATTRIBUTE_DEFINITION_ENUMERATION.CommitPhaseOne(stage)
	backRepo.BackRepoATTRIBUTE_DEFINITION_INTEGER.CommitPhaseOne(stage)
	backRepo.BackRepoATTRIBUTE_DEFINITION_REAL.CommitPhaseOne(stage)
	backRepo.BackRepoATTRIBUTE_DEFINITION_STRING.CommitPhaseOne(stage)
	backRepo.BackRepoATTRIBUTE_DEFINITION_XHTML.CommitPhaseOne(stage)
	backRepo.BackRepoATTRIBUTE_VALUE_BOOLEAN.CommitPhaseOne(stage)
	backRepo.BackRepoATTRIBUTE_VALUE_DATE.CommitPhaseOne(stage)
	backRepo.BackRepoATTRIBUTE_VALUE_ENUMERATION.CommitPhaseOne(stage)
	backRepo.BackRepoATTRIBUTE_VALUE_INTEGER.CommitPhaseOne(stage)
	backRepo.BackRepoATTRIBUTE_VALUE_REAL.CommitPhaseOne(stage)
	backRepo.BackRepoATTRIBUTE_VALUE_STRING.CommitPhaseOne(stage)
	backRepo.BackRepoATTRIBUTE_VALUE_XHTML.CommitPhaseOne(stage)
	backRepo.BackRepoDATATYPE_DEFINITION_BOOLEAN.CommitPhaseOne(stage)
	backRepo.BackRepoDATATYPE_DEFINITION_DATE.CommitPhaseOne(stage)
	backRepo.BackRepoDATATYPE_DEFINITION_ENUMERATION.CommitPhaseOne(stage)
	backRepo.BackRepoDATATYPE_DEFINITION_INTEGER.CommitPhaseOne(stage)
	backRepo.BackRepoDATATYPE_DEFINITION_REAL.CommitPhaseOne(stage)
	backRepo.BackRepoDATATYPE_DEFINITION_STRING.CommitPhaseOne(stage)
	backRepo.BackRepoDATATYPE_DEFINITION_XHTML.CommitPhaseOne(stage)
	backRepo.BackRepoEMBEDDED_VALUE.CommitPhaseOne(stage)
	backRepo.BackRepoENUM_VALUE.CommitPhaseOne(stage)
	backRepo.BackRepoRELATION_GROUP.CommitPhaseOne(stage)
	backRepo.BackRepoRELATION_GROUP_TYPE.CommitPhaseOne(stage)
	backRepo.BackRepoREQ_IF.CommitPhaseOne(stage)
	backRepo.BackRepoREQ_IF_CONTENT.CommitPhaseOne(stage)
	backRepo.BackRepoREQ_IF_HEADER.CommitPhaseOne(stage)
	backRepo.BackRepoREQ_IF_TOOL_EXTENSION.CommitPhaseOne(stage)
	backRepo.BackRepoSPECIFICATION.CommitPhaseOne(stage)
	backRepo.BackRepoSPECIFICATION_TYPE.CommitPhaseOne(stage)
	backRepo.BackRepoSPEC_HIERARCHY.CommitPhaseOne(stage)
	backRepo.BackRepoSPEC_OBJECT.CommitPhaseOne(stage)
	backRepo.BackRepoSPEC_OBJECT_TYPE.CommitPhaseOne(stage)
	backRepo.BackRepoSPEC_RELATION.CommitPhaseOne(stage)
	backRepo.BackRepoSPEC_RELATION_TYPE.CommitPhaseOne(stage)
	backRepo.BackRepoXHTML_CONTENT.CommitPhaseOne(stage)

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoALTERNATIVE_ID.CommitPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTE_DEFINITION_BOOLEAN.CommitPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTE_DEFINITION_DATE.CommitPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTE_DEFINITION_ENUMERATION.CommitPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTE_DEFINITION_INTEGER.CommitPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTE_DEFINITION_REAL.CommitPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTE_DEFINITION_STRING.CommitPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTE_DEFINITION_XHTML.CommitPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTE_VALUE_BOOLEAN.CommitPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTE_VALUE_DATE.CommitPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTE_VALUE_ENUMERATION.CommitPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTE_VALUE_INTEGER.CommitPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTE_VALUE_REAL.CommitPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTE_VALUE_STRING.CommitPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTE_VALUE_XHTML.CommitPhaseTwo(backRepo)
	backRepo.BackRepoDATATYPE_DEFINITION_BOOLEAN.CommitPhaseTwo(backRepo)
	backRepo.BackRepoDATATYPE_DEFINITION_DATE.CommitPhaseTwo(backRepo)
	backRepo.BackRepoDATATYPE_DEFINITION_ENUMERATION.CommitPhaseTwo(backRepo)
	backRepo.BackRepoDATATYPE_DEFINITION_INTEGER.CommitPhaseTwo(backRepo)
	backRepo.BackRepoDATATYPE_DEFINITION_REAL.CommitPhaseTwo(backRepo)
	backRepo.BackRepoDATATYPE_DEFINITION_STRING.CommitPhaseTwo(backRepo)
	backRepo.BackRepoDATATYPE_DEFINITION_XHTML.CommitPhaseTwo(backRepo)
	backRepo.BackRepoEMBEDDED_VALUE.CommitPhaseTwo(backRepo)
	backRepo.BackRepoENUM_VALUE.CommitPhaseTwo(backRepo)
	backRepo.BackRepoRELATION_GROUP.CommitPhaseTwo(backRepo)
	backRepo.BackRepoRELATION_GROUP_TYPE.CommitPhaseTwo(backRepo)
	backRepo.BackRepoREQ_IF.CommitPhaseTwo(backRepo)
	backRepo.BackRepoREQ_IF_CONTENT.CommitPhaseTwo(backRepo)
	backRepo.BackRepoREQ_IF_HEADER.CommitPhaseTwo(backRepo)
	backRepo.BackRepoREQ_IF_TOOL_EXTENSION.CommitPhaseTwo(backRepo)
	backRepo.BackRepoSPECIFICATION.CommitPhaseTwo(backRepo)
	backRepo.BackRepoSPECIFICATION_TYPE.CommitPhaseTwo(backRepo)
	backRepo.BackRepoSPEC_HIERARCHY.CommitPhaseTwo(backRepo)
	backRepo.BackRepoSPEC_OBJECT.CommitPhaseTwo(backRepo)
	backRepo.BackRepoSPEC_OBJECT_TYPE.CommitPhaseTwo(backRepo)
	backRepo.BackRepoSPEC_RELATION.CommitPhaseTwo(backRepo)
	backRepo.BackRepoSPEC_RELATION_TYPE.CommitPhaseTwo(backRepo)
	backRepo.BackRepoXHTML_CONTENT.CommitPhaseTwo(backRepo)

	backRepo.IncrementCommitFromBackNb()
}

// Checkout the database into the stage
func (backRepo *BackRepoStruct) Checkout(stage *models.StageStruct) {
	// insertion point for per struct back repo phase one commit
	backRepo.BackRepoALTERNATIVE_ID.CheckoutPhaseOne()
	backRepo.BackRepoATTRIBUTE_DEFINITION_BOOLEAN.CheckoutPhaseOne()
	backRepo.BackRepoATTRIBUTE_DEFINITION_DATE.CheckoutPhaseOne()
	backRepo.BackRepoATTRIBUTE_DEFINITION_ENUMERATION.CheckoutPhaseOne()
	backRepo.BackRepoATTRIBUTE_DEFINITION_INTEGER.CheckoutPhaseOne()
	backRepo.BackRepoATTRIBUTE_DEFINITION_REAL.CheckoutPhaseOne()
	backRepo.BackRepoATTRIBUTE_DEFINITION_STRING.CheckoutPhaseOne()
	backRepo.BackRepoATTRIBUTE_DEFINITION_XHTML.CheckoutPhaseOne()
	backRepo.BackRepoATTRIBUTE_VALUE_BOOLEAN.CheckoutPhaseOne()
	backRepo.BackRepoATTRIBUTE_VALUE_DATE.CheckoutPhaseOne()
	backRepo.BackRepoATTRIBUTE_VALUE_ENUMERATION.CheckoutPhaseOne()
	backRepo.BackRepoATTRIBUTE_VALUE_INTEGER.CheckoutPhaseOne()
	backRepo.BackRepoATTRIBUTE_VALUE_REAL.CheckoutPhaseOne()
	backRepo.BackRepoATTRIBUTE_VALUE_STRING.CheckoutPhaseOne()
	backRepo.BackRepoATTRIBUTE_VALUE_XHTML.CheckoutPhaseOne()
	backRepo.BackRepoDATATYPE_DEFINITION_BOOLEAN.CheckoutPhaseOne()
	backRepo.BackRepoDATATYPE_DEFINITION_DATE.CheckoutPhaseOne()
	backRepo.BackRepoDATATYPE_DEFINITION_ENUMERATION.CheckoutPhaseOne()
	backRepo.BackRepoDATATYPE_DEFINITION_INTEGER.CheckoutPhaseOne()
	backRepo.BackRepoDATATYPE_DEFINITION_REAL.CheckoutPhaseOne()
	backRepo.BackRepoDATATYPE_DEFINITION_STRING.CheckoutPhaseOne()
	backRepo.BackRepoDATATYPE_DEFINITION_XHTML.CheckoutPhaseOne()
	backRepo.BackRepoEMBEDDED_VALUE.CheckoutPhaseOne()
	backRepo.BackRepoENUM_VALUE.CheckoutPhaseOne()
	backRepo.BackRepoRELATION_GROUP.CheckoutPhaseOne()
	backRepo.BackRepoRELATION_GROUP_TYPE.CheckoutPhaseOne()
	backRepo.BackRepoREQ_IF.CheckoutPhaseOne()
	backRepo.BackRepoREQ_IF_CONTENT.CheckoutPhaseOne()
	backRepo.BackRepoREQ_IF_HEADER.CheckoutPhaseOne()
	backRepo.BackRepoREQ_IF_TOOL_EXTENSION.CheckoutPhaseOne()
	backRepo.BackRepoSPECIFICATION.CheckoutPhaseOne()
	backRepo.BackRepoSPECIFICATION_TYPE.CheckoutPhaseOne()
	backRepo.BackRepoSPEC_HIERARCHY.CheckoutPhaseOne()
	backRepo.BackRepoSPEC_OBJECT.CheckoutPhaseOne()
	backRepo.BackRepoSPEC_OBJECT_TYPE.CheckoutPhaseOne()
	backRepo.BackRepoSPEC_RELATION.CheckoutPhaseOne()
	backRepo.BackRepoSPEC_RELATION_TYPE.CheckoutPhaseOne()
	backRepo.BackRepoXHTML_CONTENT.CheckoutPhaseOne()

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoALTERNATIVE_ID.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTE_DEFINITION_BOOLEAN.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTE_DEFINITION_DATE.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTE_DEFINITION_ENUMERATION.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTE_DEFINITION_INTEGER.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTE_DEFINITION_REAL.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTE_DEFINITION_STRING.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTE_DEFINITION_XHTML.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTE_VALUE_BOOLEAN.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTE_VALUE_DATE.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTE_VALUE_ENUMERATION.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTE_VALUE_INTEGER.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTE_VALUE_REAL.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTE_VALUE_STRING.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoATTRIBUTE_VALUE_XHTML.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoDATATYPE_DEFINITION_BOOLEAN.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoDATATYPE_DEFINITION_DATE.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoDATATYPE_DEFINITION_ENUMERATION.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoDATATYPE_DEFINITION_INTEGER.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoDATATYPE_DEFINITION_REAL.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoDATATYPE_DEFINITION_STRING.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoDATATYPE_DEFINITION_XHTML.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoEMBEDDED_VALUE.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoENUM_VALUE.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoRELATION_GROUP.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoRELATION_GROUP_TYPE.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoREQ_IF.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoREQ_IF_CONTENT.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoREQ_IF_HEADER.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoREQ_IF_TOOL_EXTENSION.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoSPECIFICATION.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoSPECIFICATION_TYPE.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoSPEC_HIERARCHY.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoSPEC_OBJECT.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoSPEC_OBJECT_TYPE.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoSPEC_RELATION.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoSPEC_RELATION_TYPE.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoXHTML_CONTENT.CheckoutPhaseTwo(backRepo)
}

// Backup the BackRepoStruct
func (backRepo *BackRepoStruct) Backup(stage *models.StageStruct, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// insertion point for per struct backup
	backRepo.BackRepoALTERNATIVE_ID.Backup(dirPath)
	backRepo.BackRepoATTRIBUTE_DEFINITION_BOOLEAN.Backup(dirPath)
	backRepo.BackRepoATTRIBUTE_DEFINITION_DATE.Backup(dirPath)
	backRepo.BackRepoATTRIBUTE_DEFINITION_ENUMERATION.Backup(dirPath)
	backRepo.BackRepoATTRIBUTE_DEFINITION_INTEGER.Backup(dirPath)
	backRepo.BackRepoATTRIBUTE_DEFINITION_REAL.Backup(dirPath)
	backRepo.BackRepoATTRIBUTE_DEFINITION_STRING.Backup(dirPath)
	backRepo.BackRepoATTRIBUTE_DEFINITION_XHTML.Backup(dirPath)
	backRepo.BackRepoATTRIBUTE_VALUE_BOOLEAN.Backup(dirPath)
	backRepo.BackRepoATTRIBUTE_VALUE_DATE.Backup(dirPath)
	backRepo.BackRepoATTRIBUTE_VALUE_ENUMERATION.Backup(dirPath)
	backRepo.BackRepoATTRIBUTE_VALUE_INTEGER.Backup(dirPath)
	backRepo.BackRepoATTRIBUTE_VALUE_REAL.Backup(dirPath)
	backRepo.BackRepoATTRIBUTE_VALUE_STRING.Backup(dirPath)
	backRepo.BackRepoATTRIBUTE_VALUE_XHTML.Backup(dirPath)
	backRepo.BackRepoDATATYPE_DEFINITION_BOOLEAN.Backup(dirPath)
	backRepo.BackRepoDATATYPE_DEFINITION_DATE.Backup(dirPath)
	backRepo.BackRepoDATATYPE_DEFINITION_ENUMERATION.Backup(dirPath)
	backRepo.BackRepoDATATYPE_DEFINITION_INTEGER.Backup(dirPath)
	backRepo.BackRepoDATATYPE_DEFINITION_REAL.Backup(dirPath)
	backRepo.BackRepoDATATYPE_DEFINITION_STRING.Backup(dirPath)
	backRepo.BackRepoDATATYPE_DEFINITION_XHTML.Backup(dirPath)
	backRepo.BackRepoEMBEDDED_VALUE.Backup(dirPath)
	backRepo.BackRepoENUM_VALUE.Backup(dirPath)
	backRepo.BackRepoRELATION_GROUP.Backup(dirPath)
	backRepo.BackRepoRELATION_GROUP_TYPE.Backup(dirPath)
	backRepo.BackRepoREQ_IF.Backup(dirPath)
	backRepo.BackRepoREQ_IF_CONTENT.Backup(dirPath)
	backRepo.BackRepoREQ_IF_HEADER.Backup(dirPath)
	backRepo.BackRepoREQ_IF_TOOL_EXTENSION.Backup(dirPath)
	backRepo.BackRepoSPECIFICATION.Backup(dirPath)
	backRepo.BackRepoSPECIFICATION_TYPE.Backup(dirPath)
	backRepo.BackRepoSPEC_HIERARCHY.Backup(dirPath)
	backRepo.BackRepoSPEC_OBJECT.Backup(dirPath)
	backRepo.BackRepoSPEC_OBJECT_TYPE.Backup(dirPath)
	backRepo.BackRepoSPEC_RELATION.Backup(dirPath)
	backRepo.BackRepoSPEC_RELATION_TYPE.Backup(dirPath)
	backRepo.BackRepoXHTML_CONTENT.Backup(dirPath)
}

// Backup in XL the BackRepoStruct
func (backRepo *BackRepoStruct) BackupXL(stage *models.StageStruct, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// open an existing file
	file := xlsx.NewFile()

	// insertion point for per struct backup
	backRepo.BackRepoALTERNATIVE_ID.BackupXL(file)
	backRepo.BackRepoATTRIBUTE_DEFINITION_BOOLEAN.BackupXL(file)
	backRepo.BackRepoATTRIBUTE_DEFINITION_DATE.BackupXL(file)
	backRepo.BackRepoATTRIBUTE_DEFINITION_ENUMERATION.BackupXL(file)
	backRepo.BackRepoATTRIBUTE_DEFINITION_INTEGER.BackupXL(file)
	backRepo.BackRepoATTRIBUTE_DEFINITION_REAL.BackupXL(file)
	backRepo.BackRepoATTRIBUTE_DEFINITION_STRING.BackupXL(file)
	backRepo.BackRepoATTRIBUTE_DEFINITION_XHTML.BackupXL(file)
	backRepo.BackRepoATTRIBUTE_VALUE_BOOLEAN.BackupXL(file)
	backRepo.BackRepoATTRIBUTE_VALUE_DATE.BackupXL(file)
	backRepo.BackRepoATTRIBUTE_VALUE_ENUMERATION.BackupXL(file)
	backRepo.BackRepoATTRIBUTE_VALUE_INTEGER.BackupXL(file)
	backRepo.BackRepoATTRIBUTE_VALUE_REAL.BackupXL(file)
	backRepo.BackRepoATTRIBUTE_VALUE_STRING.BackupXL(file)
	backRepo.BackRepoATTRIBUTE_VALUE_XHTML.BackupXL(file)
	backRepo.BackRepoDATATYPE_DEFINITION_BOOLEAN.BackupXL(file)
	backRepo.BackRepoDATATYPE_DEFINITION_DATE.BackupXL(file)
	backRepo.BackRepoDATATYPE_DEFINITION_ENUMERATION.BackupXL(file)
	backRepo.BackRepoDATATYPE_DEFINITION_INTEGER.BackupXL(file)
	backRepo.BackRepoDATATYPE_DEFINITION_REAL.BackupXL(file)
	backRepo.BackRepoDATATYPE_DEFINITION_STRING.BackupXL(file)
	backRepo.BackRepoDATATYPE_DEFINITION_XHTML.BackupXL(file)
	backRepo.BackRepoEMBEDDED_VALUE.BackupXL(file)
	backRepo.BackRepoENUM_VALUE.BackupXL(file)
	backRepo.BackRepoRELATION_GROUP.BackupXL(file)
	backRepo.BackRepoRELATION_GROUP_TYPE.BackupXL(file)
	backRepo.BackRepoREQ_IF.BackupXL(file)
	backRepo.BackRepoREQ_IF_CONTENT.BackupXL(file)
	backRepo.BackRepoREQ_IF_HEADER.BackupXL(file)
	backRepo.BackRepoREQ_IF_TOOL_EXTENSION.BackupXL(file)
	backRepo.BackRepoSPECIFICATION.BackupXL(file)
	backRepo.BackRepoSPECIFICATION_TYPE.BackupXL(file)
	backRepo.BackRepoSPEC_HIERARCHY.BackupXL(file)
	backRepo.BackRepoSPEC_OBJECT.BackupXL(file)
	backRepo.BackRepoSPEC_OBJECT_TYPE.BackupXL(file)
	backRepo.BackRepoSPEC_RELATION.BackupXL(file)
	backRepo.BackRepoSPEC_RELATION_TYPE.BackupXL(file)
	backRepo.BackRepoXHTML_CONTENT.BackupXL(file)

	var b bytes.Buffer
	writer := bufio.NewWriter(&b)
	file.Write(writer)
	theBytes := b.Bytes()

	filename := filepath.Join(dirPath, "bckp.xlsx")
	err := ioutil.WriteFile(filename, theBytes, 0644)
	if err != nil {
		log.Panic("Cannot write the XL file", err.Error())
	}
}

// Restore the database into the back repo
func (backRepo *BackRepoStruct) Restore(stage *models.StageStruct, dirPath string) {
	backRepo.stage.Commit()
	backRepo.stage.Reset()
	backRepo.stage.Checkout()

	//
	// restauration first phase (create DB instance with new IDs)
	//

	// insertion point for per struct backup
	backRepo.BackRepoALTERNATIVE_ID.RestorePhaseOne(dirPath)
	backRepo.BackRepoATTRIBUTE_DEFINITION_BOOLEAN.RestorePhaseOne(dirPath)
	backRepo.BackRepoATTRIBUTE_DEFINITION_DATE.RestorePhaseOne(dirPath)
	backRepo.BackRepoATTRIBUTE_DEFINITION_ENUMERATION.RestorePhaseOne(dirPath)
	backRepo.BackRepoATTRIBUTE_DEFINITION_INTEGER.RestorePhaseOne(dirPath)
	backRepo.BackRepoATTRIBUTE_DEFINITION_REAL.RestorePhaseOne(dirPath)
	backRepo.BackRepoATTRIBUTE_DEFINITION_STRING.RestorePhaseOne(dirPath)
	backRepo.BackRepoATTRIBUTE_DEFINITION_XHTML.RestorePhaseOne(dirPath)
	backRepo.BackRepoATTRIBUTE_VALUE_BOOLEAN.RestorePhaseOne(dirPath)
	backRepo.BackRepoATTRIBUTE_VALUE_DATE.RestorePhaseOne(dirPath)
	backRepo.BackRepoATTRIBUTE_VALUE_ENUMERATION.RestorePhaseOne(dirPath)
	backRepo.BackRepoATTRIBUTE_VALUE_INTEGER.RestorePhaseOne(dirPath)
	backRepo.BackRepoATTRIBUTE_VALUE_REAL.RestorePhaseOne(dirPath)
	backRepo.BackRepoATTRIBUTE_VALUE_STRING.RestorePhaseOne(dirPath)
	backRepo.BackRepoATTRIBUTE_VALUE_XHTML.RestorePhaseOne(dirPath)
	backRepo.BackRepoDATATYPE_DEFINITION_BOOLEAN.RestorePhaseOne(dirPath)
	backRepo.BackRepoDATATYPE_DEFINITION_DATE.RestorePhaseOne(dirPath)
	backRepo.BackRepoDATATYPE_DEFINITION_ENUMERATION.RestorePhaseOne(dirPath)
	backRepo.BackRepoDATATYPE_DEFINITION_INTEGER.RestorePhaseOne(dirPath)
	backRepo.BackRepoDATATYPE_DEFINITION_REAL.RestorePhaseOne(dirPath)
	backRepo.BackRepoDATATYPE_DEFINITION_STRING.RestorePhaseOne(dirPath)
	backRepo.BackRepoDATATYPE_DEFINITION_XHTML.RestorePhaseOne(dirPath)
	backRepo.BackRepoEMBEDDED_VALUE.RestorePhaseOne(dirPath)
	backRepo.BackRepoENUM_VALUE.RestorePhaseOne(dirPath)
	backRepo.BackRepoRELATION_GROUP.RestorePhaseOne(dirPath)
	backRepo.BackRepoRELATION_GROUP_TYPE.RestorePhaseOne(dirPath)
	backRepo.BackRepoREQ_IF.RestorePhaseOne(dirPath)
	backRepo.BackRepoREQ_IF_CONTENT.RestorePhaseOne(dirPath)
	backRepo.BackRepoREQ_IF_HEADER.RestorePhaseOne(dirPath)
	backRepo.BackRepoREQ_IF_TOOL_EXTENSION.RestorePhaseOne(dirPath)
	backRepo.BackRepoSPECIFICATION.RestorePhaseOne(dirPath)
	backRepo.BackRepoSPECIFICATION_TYPE.RestorePhaseOne(dirPath)
	backRepo.BackRepoSPEC_HIERARCHY.RestorePhaseOne(dirPath)
	backRepo.BackRepoSPEC_OBJECT.RestorePhaseOne(dirPath)
	backRepo.BackRepoSPEC_OBJECT_TYPE.RestorePhaseOne(dirPath)
	backRepo.BackRepoSPEC_RELATION.RestorePhaseOne(dirPath)
	backRepo.BackRepoSPEC_RELATION_TYPE.RestorePhaseOne(dirPath)
	backRepo.BackRepoXHTML_CONTENT.RestorePhaseOne(dirPath)

	//
	// restauration second phase (reindex pointers with the new ID)
	//

	// insertion point for per struct backup
	backRepo.BackRepoALTERNATIVE_ID.RestorePhaseTwo()
	backRepo.BackRepoATTRIBUTE_DEFINITION_BOOLEAN.RestorePhaseTwo()
	backRepo.BackRepoATTRIBUTE_DEFINITION_DATE.RestorePhaseTwo()
	backRepo.BackRepoATTRIBUTE_DEFINITION_ENUMERATION.RestorePhaseTwo()
	backRepo.BackRepoATTRIBUTE_DEFINITION_INTEGER.RestorePhaseTwo()
	backRepo.BackRepoATTRIBUTE_DEFINITION_REAL.RestorePhaseTwo()
	backRepo.BackRepoATTRIBUTE_DEFINITION_STRING.RestorePhaseTwo()
	backRepo.BackRepoATTRIBUTE_DEFINITION_XHTML.RestorePhaseTwo()
	backRepo.BackRepoATTRIBUTE_VALUE_BOOLEAN.RestorePhaseTwo()
	backRepo.BackRepoATTRIBUTE_VALUE_DATE.RestorePhaseTwo()
	backRepo.BackRepoATTRIBUTE_VALUE_ENUMERATION.RestorePhaseTwo()
	backRepo.BackRepoATTRIBUTE_VALUE_INTEGER.RestorePhaseTwo()
	backRepo.BackRepoATTRIBUTE_VALUE_REAL.RestorePhaseTwo()
	backRepo.BackRepoATTRIBUTE_VALUE_STRING.RestorePhaseTwo()
	backRepo.BackRepoATTRIBUTE_VALUE_XHTML.RestorePhaseTwo()
	backRepo.BackRepoDATATYPE_DEFINITION_BOOLEAN.RestorePhaseTwo()
	backRepo.BackRepoDATATYPE_DEFINITION_DATE.RestorePhaseTwo()
	backRepo.BackRepoDATATYPE_DEFINITION_ENUMERATION.RestorePhaseTwo()
	backRepo.BackRepoDATATYPE_DEFINITION_INTEGER.RestorePhaseTwo()
	backRepo.BackRepoDATATYPE_DEFINITION_REAL.RestorePhaseTwo()
	backRepo.BackRepoDATATYPE_DEFINITION_STRING.RestorePhaseTwo()
	backRepo.BackRepoDATATYPE_DEFINITION_XHTML.RestorePhaseTwo()
	backRepo.BackRepoEMBEDDED_VALUE.RestorePhaseTwo()
	backRepo.BackRepoENUM_VALUE.RestorePhaseTwo()
	backRepo.BackRepoRELATION_GROUP.RestorePhaseTwo()
	backRepo.BackRepoRELATION_GROUP_TYPE.RestorePhaseTwo()
	backRepo.BackRepoREQ_IF.RestorePhaseTwo()
	backRepo.BackRepoREQ_IF_CONTENT.RestorePhaseTwo()
	backRepo.BackRepoREQ_IF_HEADER.RestorePhaseTwo()
	backRepo.BackRepoREQ_IF_TOOL_EXTENSION.RestorePhaseTwo()
	backRepo.BackRepoSPECIFICATION.RestorePhaseTwo()
	backRepo.BackRepoSPECIFICATION_TYPE.RestorePhaseTwo()
	backRepo.BackRepoSPEC_HIERARCHY.RestorePhaseTwo()
	backRepo.BackRepoSPEC_OBJECT.RestorePhaseTwo()
	backRepo.BackRepoSPEC_OBJECT_TYPE.RestorePhaseTwo()
	backRepo.BackRepoSPEC_RELATION.RestorePhaseTwo()
	backRepo.BackRepoSPEC_RELATION_TYPE.RestorePhaseTwo()
	backRepo.BackRepoXHTML_CONTENT.RestorePhaseTwo()

	backRepo.stage.Checkout()
}

// Restore the database into the back repo
func (backRepo *BackRepoStruct) RestoreXL(stage *models.StageStruct, dirPath string) {

	// clean the stage
	backRepo.stage.Reset()

	// commit the cleaned stage
	backRepo.stage.Commit()

	// open an existing file
	filename := filepath.Join(dirPath, "bckp.xlsx")
	file, err := xlsx.OpenFile(filename)
	_ = file

	if err != nil {
		log.Panic("Cannot read the XL file", err.Error())
	}

	//
	// restauration first phase (create DB instance with new IDs)
	//

	// insertion point for per struct backup
	backRepo.BackRepoALTERNATIVE_ID.RestoreXLPhaseOne(file)
	backRepo.BackRepoATTRIBUTE_DEFINITION_BOOLEAN.RestoreXLPhaseOne(file)
	backRepo.BackRepoATTRIBUTE_DEFINITION_DATE.RestoreXLPhaseOne(file)
	backRepo.BackRepoATTRIBUTE_DEFINITION_ENUMERATION.RestoreXLPhaseOne(file)
	backRepo.BackRepoATTRIBUTE_DEFINITION_INTEGER.RestoreXLPhaseOne(file)
	backRepo.BackRepoATTRIBUTE_DEFINITION_REAL.RestoreXLPhaseOne(file)
	backRepo.BackRepoATTRIBUTE_DEFINITION_STRING.RestoreXLPhaseOne(file)
	backRepo.BackRepoATTRIBUTE_DEFINITION_XHTML.RestoreXLPhaseOne(file)
	backRepo.BackRepoATTRIBUTE_VALUE_BOOLEAN.RestoreXLPhaseOne(file)
	backRepo.BackRepoATTRIBUTE_VALUE_DATE.RestoreXLPhaseOne(file)
	backRepo.BackRepoATTRIBUTE_VALUE_ENUMERATION.RestoreXLPhaseOne(file)
	backRepo.BackRepoATTRIBUTE_VALUE_INTEGER.RestoreXLPhaseOne(file)
	backRepo.BackRepoATTRIBUTE_VALUE_REAL.RestoreXLPhaseOne(file)
	backRepo.BackRepoATTRIBUTE_VALUE_STRING.RestoreXLPhaseOne(file)
	backRepo.BackRepoATTRIBUTE_VALUE_XHTML.RestoreXLPhaseOne(file)
	backRepo.BackRepoDATATYPE_DEFINITION_BOOLEAN.RestoreXLPhaseOne(file)
	backRepo.BackRepoDATATYPE_DEFINITION_DATE.RestoreXLPhaseOne(file)
	backRepo.BackRepoDATATYPE_DEFINITION_ENUMERATION.RestoreXLPhaseOne(file)
	backRepo.BackRepoDATATYPE_DEFINITION_INTEGER.RestoreXLPhaseOne(file)
	backRepo.BackRepoDATATYPE_DEFINITION_REAL.RestoreXLPhaseOne(file)
	backRepo.BackRepoDATATYPE_DEFINITION_STRING.RestoreXLPhaseOne(file)
	backRepo.BackRepoDATATYPE_DEFINITION_XHTML.RestoreXLPhaseOne(file)
	backRepo.BackRepoEMBEDDED_VALUE.RestoreXLPhaseOne(file)
	backRepo.BackRepoENUM_VALUE.RestoreXLPhaseOne(file)
	backRepo.BackRepoRELATION_GROUP.RestoreXLPhaseOne(file)
	backRepo.BackRepoRELATION_GROUP_TYPE.RestoreXLPhaseOne(file)
	backRepo.BackRepoREQ_IF.RestoreXLPhaseOne(file)
	backRepo.BackRepoREQ_IF_CONTENT.RestoreXLPhaseOne(file)
	backRepo.BackRepoREQ_IF_HEADER.RestoreXLPhaseOne(file)
	backRepo.BackRepoREQ_IF_TOOL_EXTENSION.RestoreXLPhaseOne(file)
	backRepo.BackRepoSPECIFICATION.RestoreXLPhaseOne(file)
	backRepo.BackRepoSPECIFICATION_TYPE.RestoreXLPhaseOne(file)
	backRepo.BackRepoSPEC_HIERARCHY.RestoreXLPhaseOne(file)
	backRepo.BackRepoSPEC_OBJECT.RestoreXLPhaseOne(file)
	backRepo.BackRepoSPEC_OBJECT_TYPE.RestoreXLPhaseOne(file)
	backRepo.BackRepoSPEC_RELATION.RestoreXLPhaseOne(file)
	backRepo.BackRepoSPEC_RELATION_TYPE.RestoreXLPhaseOne(file)
	backRepo.BackRepoXHTML_CONTENT.RestoreXLPhaseOne(file)

	// commit the restored stage
	backRepo.stage.Commit()
}

func (backRepoStruct *BackRepoStruct) SubscribeToCommitNb() <-chan int {
	backRepoStruct.rwMutex.Lock()
	defer backRepoStruct.rwMutex.Unlock()

	ch := make(chan int)
	backRepoStruct.subscribers = append(backRepoStruct.subscribers, ch)
	return ch
}

func (backRepoStruct *BackRepoStruct) broadcastNbCommitToBack() {
	backRepoStruct.rwMutex.RLock()
	defer backRepoStruct.rwMutex.RUnlock()

	activeChannels := make([]chan int, 0)

	for _, ch := range backRepoStruct.subscribers {
		select {
		case ch <- int(backRepoStruct.CommitFromBackNb):
			activeChannels = append(activeChannels, ch)
		default:
			// Assume channel is no longer active; don't add to activeChannels
			log.Println("Channel no longer active")
			close(ch)
		}
	}
	backRepoStruct.subscribers = activeChannels
}
