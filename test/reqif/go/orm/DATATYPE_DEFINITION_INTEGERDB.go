// generated by stacks/gong/go/models/orm_file_per_struct_back_repo.go
package orm

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"time"

	"gorm.io/gorm"

	"github.com/tealeg/xlsx/v3"

	"github.com/fullstack-lang/gongxsd/test/reqif/go/models"
)

// dummy variable to have the import declaration wihthout compile failure (even if no code needing this import is generated)
var dummy_DATATYPE_DEFINITION_INTEGER_sql sql.NullBool
var dummy_DATATYPE_DEFINITION_INTEGER_time time.Duration
var dummy_DATATYPE_DEFINITION_INTEGER_sort sort.Float64Slice

// DATATYPE_DEFINITION_INTEGERAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model datatype_definition_integerAPI
type DATATYPE_DEFINITION_INTEGERAPI struct {
	gorm.Model

	models.DATATYPE_DEFINITION_INTEGER_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	DATATYPE_DEFINITION_INTEGERPointersEncoding DATATYPE_DEFINITION_INTEGERPointersEncoding
}

// DATATYPE_DEFINITION_INTEGERPointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type DATATYPE_DEFINITION_INTEGERPointersEncoding struct {
	// insertion for pointer fields encoding declaration
}

// DATATYPE_DEFINITION_INTEGERDB describes a datatype_definition_integer in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model datatype_definition_integerDB
type DATATYPE_DEFINITION_INTEGERDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field datatype_definition_integerDB.Name
	Name_Data sql.NullString

	// Declation for basic field datatype_definition_integerDB.DESC
	DESC_Data sql.NullString

	// Declation for basic field datatype_definition_integerDB.IDENTIFIER
	IDENTIFIER_Data sql.NullString

	// Declation for basic field datatype_definition_integerDB.LAST_CHANGE
	LAST_CHANGE_Data sql.NullString

	// Declation for basic field datatype_definition_integerDB.LONG_NAME
	LONG_NAME_Data sql.NullString

	// Declation for basic field datatype_definition_integerDB.MAX
	MAX_Data sql.NullInt64

	// Declation for basic field datatype_definition_integerDB.MIN
	MIN_Data sql.NullInt64
	
	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	DATATYPE_DEFINITION_INTEGERPointersEncoding
}

// DATATYPE_DEFINITION_INTEGERDBs arrays datatype_definition_integerDBs
// swagger:response datatype_definition_integerDBsResponse
type DATATYPE_DEFINITION_INTEGERDBs []DATATYPE_DEFINITION_INTEGERDB

// DATATYPE_DEFINITION_INTEGERDBResponse provides response
// swagger:response datatype_definition_integerDBResponse
type DATATYPE_DEFINITION_INTEGERDBResponse struct {
	DATATYPE_DEFINITION_INTEGERDB
}

// DATATYPE_DEFINITION_INTEGERWOP is a DATATYPE_DEFINITION_INTEGER without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type DATATYPE_DEFINITION_INTEGERWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	DESC string `xlsx:"2"`

	IDENTIFIER string `xlsx:"3"`

	LAST_CHANGE string `xlsx:"4"`

	LONG_NAME string `xlsx:"5"`

	MAX int `xlsx:"6"`

	MIN int `xlsx:"7"`
	// insertion for WOP pointer fields
}

var DATATYPE_DEFINITION_INTEGER_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"DESC",
	"IDENTIFIER",
	"LAST_CHANGE",
	"LONG_NAME",
	"MAX",
	"MIN",
}

type BackRepoDATATYPE_DEFINITION_INTEGERStruct struct {
	// stores DATATYPE_DEFINITION_INTEGERDB according to their gorm ID
	Map_DATATYPE_DEFINITION_INTEGERDBID_DATATYPE_DEFINITION_INTEGERDB map[uint]*DATATYPE_DEFINITION_INTEGERDB

	// stores DATATYPE_DEFINITION_INTEGERDB ID according to DATATYPE_DEFINITION_INTEGER address
	Map_DATATYPE_DEFINITION_INTEGERPtr_DATATYPE_DEFINITION_INTEGERDBID map[*models.DATATYPE_DEFINITION_INTEGER]uint

	// stores DATATYPE_DEFINITION_INTEGER according to their gorm ID
	Map_DATATYPE_DEFINITION_INTEGERDBID_DATATYPE_DEFINITION_INTEGERPtr map[uint]*models.DATATYPE_DEFINITION_INTEGER

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoDATATYPE_DEFINITION_INTEGER *BackRepoDATATYPE_DEFINITION_INTEGERStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoDATATYPE_DEFINITION_INTEGER.stage
	return
}

func (backRepoDATATYPE_DEFINITION_INTEGER *BackRepoDATATYPE_DEFINITION_INTEGERStruct) GetDB() *gorm.DB {
	return backRepoDATATYPE_DEFINITION_INTEGER.db
}

// GetDATATYPE_DEFINITION_INTEGERDBFromDATATYPE_DEFINITION_INTEGERPtr is a handy function to access the back repo instance from the stage instance
func (backRepoDATATYPE_DEFINITION_INTEGER *BackRepoDATATYPE_DEFINITION_INTEGERStruct) GetDATATYPE_DEFINITION_INTEGERDBFromDATATYPE_DEFINITION_INTEGERPtr(datatype_definition_integer *models.DATATYPE_DEFINITION_INTEGER) (datatype_definition_integerDB *DATATYPE_DEFINITION_INTEGERDB) {
	id := backRepoDATATYPE_DEFINITION_INTEGER.Map_DATATYPE_DEFINITION_INTEGERPtr_DATATYPE_DEFINITION_INTEGERDBID[datatype_definition_integer]
	datatype_definition_integerDB = backRepoDATATYPE_DEFINITION_INTEGER.Map_DATATYPE_DEFINITION_INTEGERDBID_DATATYPE_DEFINITION_INTEGERDB[id]
	return
}

// BackRepoDATATYPE_DEFINITION_INTEGER.CommitPhaseOne commits all staged instances of DATATYPE_DEFINITION_INTEGER to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoDATATYPE_DEFINITION_INTEGER *BackRepoDATATYPE_DEFINITION_INTEGERStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for datatype_definition_integer := range stage.DATATYPE_DEFINITION_INTEGERs {
		backRepoDATATYPE_DEFINITION_INTEGER.CommitPhaseOneInstance(datatype_definition_integer)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, datatype_definition_integer := range backRepoDATATYPE_DEFINITION_INTEGER.Map_DATATYPE_DEFINITION_INTEGERDBID_DATATYPE_DEFINITION_INTEGERPtr {
		if _, ok := stage.DATATYPE_DEFINITION_INTEGERs[datatype_definition_integer]; !ok {
			backRepoDATATYPE_DEFINITION_INTEGER.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoDATATYPE_DEFINITION_INTEGER.CommitDeleteInstance commits deletion of DATATYPE_DEFINITION_INTEGER to the BackRepo
func (backRepoDATATYPE_DEFINITION_INTEGER *BackRepoDATATYPE_DEFINITION_INTEGERStruct) CommitDeleteInstance(id uint) (Error error) {

	datatype_definition_integer := backRepoDATATYPE_DEFINITION_INTEGER.Map_DATATYPE_DEFINITION_INTEGERDBID_DATATYPE_DEFINITION_INTEGERPtr[id]

	// datatype_definition_integer is not staged anymore, remove datatype_definition_integerDB
	datatype_definition_integerDB := backRepoDATATYPE_DEFINITION_INTEGER.Map_DATATYPE_DEFINITION_INTEGERDBID_DATATYPE_DEFINITION_INTEGERDB[id]
	query := backRepoDATATYPE_DEFINITION_INTEGER.db.Unscoped().Delete(&datatype_definition_integerDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	delete(backRepoDATATYPE_DEFINITION_INTEGER.Map_DATATYPE_DEFINITION_INTEGERPtr_DATATYPE_DEFINITION_INTEGERDBID, datatype_definition_integer)
	delete(backRepoDATATYPE_DEFINITION_INTEGER.Map_DATATYPE_DEFINITION_INTEGERDBID_DATATYPE_DEFINITION_INTEGERPtr, id)
	delete(backRepoDATATYPE_DEFINITION_INTEGER.Map_DATATYPE_DEFINITION_INTEGERDBID_DATATYPE_DEFINITION_INTEGERDB, id)

	return
}

// BackRepoDATATYPE_DEFINITION_INTEGER.CommitPhaseOneInstance commits datatype_definition_integer staged instances of DATATYPE_DEFINITION_INTEGER to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoDATATYPE_DEFINITION_INTEGER *BackRepoDATATYPE_DEFINITION_INTEGERStruct) CommitPhaseOneInstance(datatype_definition_integer *models.DATATYPE_DEFINITION_INTEGER) (Error error) {

	// check if the datatype_definition_integer is not commited yet
	if _, ok := backRepoDATATYPE_DEFINITION_INTEGER.Map_DATATYPE_DEFINITION_INTEGERPtr_DATATYPE_DEFINITION_INTEGERDBID[datatype_definition_integer]; ok {
		return
	}

	// initiate datatype_definition_integer
	var datatype_definition_integerDB DATATYPE_DEFINITION_INTEGERDB
	datatype_definition_integerDB.CopyBasicFieldsFromDATATYPE_DEFINITION_INTEGER(datatype_definition_integer)

	query := backRepoDATATYPE_DEFINITION_INTEGER.db.Create(&datatype_definition_integerDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	backRepoDATATYPE_DEFINITION_INTEGER.Map_DATATYPE_DEFINITION_INTEGERPtr_DATATYPE_DEFINITION_INTEGERDBID[datatype_definition_integer] = datatype_definition_integerDB.ID
	backRepoDATATYPE_DEFINITION_INTEGER.Map_DATATYPE_DEFINITION_INTEGERDBID_DATATYPE_DEFINITION_INTEGERPtr[datatype_definition_integerDB.ID] = datatype_definition_integer
	backRepoDATATYPE_DEFINITION_INTEGER.Map_DATATYPE_DEFINITION_INTEGERDBID_DATATYPE_DEFINITION_INTEGERDB[datatype_definition_integerDB.ID] = &datatype_definition_integerDB

	return
}

// BackRepoDATATYPE_DEFINITION_INTEGER.CommitPhaseTwo commits all staged instances of DATATYPE_DEFINITION_INTEGER to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoDATATYPE_DEFINITION_INTEGER *BackRepoDATATYPE_DEFINITION_INTEGERStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, datatype_definition_integer := range backRepoDATATYPE_DEFINITION_INTEGER.Map_DATATYPE_DEFINITION_INTEGERDBID_DATATYPE_DEFINITION_INTEGERPtr {
		backRepoDATATYPE_DEFINITION_INTEGER.CommitPhaseTwoInstance(backRepo, idx, datatype_definition_integer)
	}

	return
}

// BackRepoDATATYPE_DEFINITION_INTEGER.CommitPhaseTwoInstance commits {{structname }} of models.DATATYPE_DEFINITION_INTEGER to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoDATATYPE_DEFINITION_INTEGER *BackRepoDATATYPE_DEFINITION_INTEGERStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, datatype_definition_integer *models.DATATYPE_DEFINITION_INTEGER) (Error error) {

	// fetch matching datatype_definition_integerDB
	if datatype_definition_integerDB, ok := backRepoDATATYPE_DEFINITION_INTEGER.Map_DATATYPE_DEFINITION_INTEGERDBID_DATATYPE_DEFINITION_INTEGERDB[idx]; ok {

		datatype_definition_integerDB.CopyBasicFieldsFromDATATYPE_DEFINITION_INTEGER(datatype_definition_integer)

		// insertion point for translating pointers encodings into actual pointers
		query := backRepoDATATYPE_DEFINITION_INTEGER.db.Save(&datatype_definition_integerDB)
		if query.Error != nil {
			log.Fatalln(query.Error)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown DATATYPE_DEFINITION_INTEGER intance %s", datatype_definition_integer.Name))
		return err
	}

	return
}

// BackRepoDATATYPE_DEFINITION_INTEGER.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoDATATYPE_DEFINITION_INTEGER *BackRepoDATATYPE_DEFINITION_INTEGERStruct) CheckoutPhaseOne() (Error error) {

	datatype_definition_integerDBArray := make([]DATATYPE_DEFINITION_INTEGERDB, 0)
	query := backRepoDATATYPE_DEFINITION_INTEGER.db.Find(&datatype_definition_integerDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	datatype_definition_integerInstancesToBeRemovedFromTheStage := make(map[*models.DATATYPE_DEFINITION_INTEGER]any)
	for key, value := range backRepoDATATYPE_DEFINITION_INTEGER.stage.DATATYPE_DEFINITION_INTEGERs {
		datatype_definition_integerInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, datatype_definition_integerDB := range datatype_definition_integerDBArray {
		backRepoDATATYPE_DEFINITION_INTEGER.CheckoutPhaseOneInstance(&datatype_definition_integerDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		datatype_definition_integer, ok := backRepoDATATYPE_DEFINITION_INTEGER.Map_DATATYPE_DEFINITION_INTEGERDBID_DATATYPE_DEFINITION_INTEGERPtr[datatype_definition_integerDB.ID]
		if ok {
			delete(datatype_definition_integerInstancesToBeRemovedFromTheStage, datatype_definition_integer)
		}
	}

	// remove from stage and back repo's 3 maps all datatype_definition_integers that are not in the checkout
	for datatype_definition_integer := range datatype_definition_integerInstancesToBeRemovedFromTheStage {
		datatype_definition_integer.Unstage(backRepoDATATYPE_DEFINITION_INTEGER.GetStage())

		// remove instance from the back repo 3 maps
		datatype_definition_integerID := backRepoDATATYPE_DEFINITION_INTEGER.Map_DATATYPE_DEFINITION_INTEGERPtr_DATATYPE_DEFINITION_INTEGERDBID[datatype_definition_integer]
		delete(backRepoDATATYPE_DEFINITION_INTEGER.Map_DATATYPE_DEFINITION_INTEGERPtr_DATATYPE_DEFINITION_INTEGERDBID, datatype_definition_integer)
		delete(backRepoDATATYPE_DEFINITION_INTEGER.Map_DATATYPE_DEFINITION_INTEGERDBID_DATATYPE_DEFINITION_INTEGERDB, datatype_definition_integerID)
		delete(backRepoDATATYPE_DEFINITION_INTEGER.Map_DATATYPE_DEFINITION_INTEGERDBID_DATATYPE_DEFINITION_INTEGERPtr, datatype_definition_integerID)
	}

	return
}

// CheckoutPhaseOneInstance takes a datatype_definition_integerDB that has been found in the DB, updates the backRepo and stages the
// models version of the datatype_definition_integerDB
func (backRepoDATATYPE_DEFINITION_INTEGER *BackRepoDATATYPE_DEFINITION_INTEGERStruct) CheckoutPhaseOneInstance(datatype_definition_integerDB *DATATYPE_DEFINITION_INTEGERDB) (Error error) {

	datatype_definition_integer, ok := backRepoDATATYPE_DEFINITION_INTEGER.Map_DATATYPE_DEFINITION_INTEGERDBID_DATATYPE_DEFINITION_INTEGERPtr[datatype_definition_integerDB.ID]
	if !ok {
		datatype_definition_integer = new(models.DATATYPE_DEFINITION_INTEGER)

		backRepoDATATYPE_DEFINITION_INTEGER.Map_DATATYPE_DEFINITION_INTEGERDBID_DATATYPE_DEFINITION_INTEGERPtr[datatype_definition_integerDB.ID] = datatype_definition_integer
		backRepoDATATYPE_DEFINITION_INTEGER.Map_DATATYPE_DEFINITION_INTEGERPtr_DATATYPE_DEFINITION_INTEGERDBID[datatype_definition_integer] = datatype_definition_integerDB.ID

		// append model store with the new element
		datatype_definition_integer.Name = datatype_definition_integerDB.Name_Data.String
		datatype_definition_integer.Stage(backRepoDATATYPE_DEFINITION_INTEGER.GetStage())
	}
	datatype_definition_integerDB.CopyBasicFieldsToDATATYPE_DEFINITION_INTEGER(datatype_definition_integer)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	datatype_definition_integer.Stage(backRepoDATATYPE_DEFINITION_INTEGER.GetStage())

	// preserve pointer to datatype_definition_integerDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_DATATYPE_DEFINITION_INTEGERDBID_DATATYPE_DEFINITION_INTEGERDB)[datatype_definition_integerDB hold variable pointers
	datatype_definition_integerDB_Data := *datatype_definition_integerDB
	preservedPtrToDATATYPE_DEFINITION_INTEGER := &datatype_definition_integerDB_Data
	backRepoDATATYPE_DEFINITION_INTEGER.Map_DATATYPE_DEFINITION_INTEGERDBID_DATATYPE_DEFINITION_INTEGERDB[datatype_definition_integerDB.ID] = preservedPtrToDATATYPE_DEFINITION_INTEGER

	return
}

// BackRepoDATATYPE_DEFINITION_INTEGER.CheckoutPhaseTwo Checkouts all staged instances of DATATYPE_DEFINITION_INTEGER to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoDATATYPE_DEFINITION_INTEGER *BackRepoDATATYPE_DEFINITION_INTEGERStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, datatype_definition_integerDB := range backRepoDATATYPE_DEFINITION_INTEGER.Map_DATATYPE_DEFINITION_INTEGERDBID_DATATYPE_DEFINITION_INTEGERDB {
		backRepoDATATYPE_DEFINITION_INTEGER.CheckoutPhaseTwoInstance(backRepo, datatype_definition_integerDB)
	}
	return
}

// BackRepoDATATYPE_DEFINITION_INTEGER.CheckoutPhaseTwoInstance Checkouts staged instances of DATATYPE_DEFINITION_INTEGER to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoDATATYPE_DEFINITION_INTEGER *BackRepoDATATYPE_DEFINITION_INTEGERStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, datatype_definition_integerDB *DATATYPE_DEFINITION_INTEGERDB) (Error error) {

	datatype_definition_integer := backRepoDATATYPE_DEFINITION_INTEGER.Map_DATATYPE_DEFINITION_INTEGERDBID_DATATYPE_DEFINITION_INTEGERPtr[datatype_definition_integerDB.ID]

	datatype_definition_integerDB.DecodePointers(backRepo, datatype_definition_integer)

	return
}

func (datatype_definition_integerDB *DATATYPE_DEFINITION_INTEGERDB) DecodePointers(backRepo *BackRepoStruct, datatype_definition_integer *models.DATATYPE_DEFINITION_INTEGER) {

	// insertion point for checkout of pointer encoding
	return
}

// CommitDATATYPE_DEFINITION_INTEGER allows commit of a single datatype_definition_integer (if already staged)
func (backRepo *BackRepoStruct) CommitDATATYPE_DEFINITION_INTEGER(datatype_definition_integer *models.DATATYPE_DEFINITION_INTEGER) {
	backRepo.BackRepoDATATYPE_DEFINITION_INTEGER.CommitPhaseOneInstance(datatype_definition_integer)
	if id, ok := backRepo.BackRepoDATATYPE_DEFINITION_INTEGER.Map_DATATYPE_DEFINITION_INTEGERPtr_DATATYPE_DEFINITION_INTEGERDBID[datatype_definition_integer]; ok {
		backRepo.BackRepoDATATYPE_DEFINITION_INTEGER.CommitPhaseTwoInstance(backRepo, id, datatype_definition_integer)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitDATATYPE_DEFINITION_INTEGER allows checkout of a single datatype_definition_integer (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutDATATYPE_DEFINITION_INTEGER(datatype_definition_integer *models.DATATYPE_DEFINITION_INTEGER) {
	// check if the datatype_definition_integer is staged
	if _, ok := backRepo.BackRepoDATATYPE_DEFINITION_INTEGER.Map_DATATYPE_DEFINITION_INTEGERPtr_DATATYPE_DEFINITION_INTEGERDBID[datatype_definition_integer]; ok {

		if id, ok := backRepo.BackRepoDATATYPE_DEFINITION_INTEGER.Map_DATATYPE_DEFINITION_INTEGERPtr_DATATYPE_DEFINITION_INTEGERDBID[datatype_definition_integer]; ok {
			var datatype_definition_integerDB DATATYPE_DEFINITION_INTEGERDB
			datatype_definition_integerDB.ID = id

			if err := backRepo.BackRepoDATATYPE_DEFINITION_INTEGER.db.First(&datatype_definition_integerDB, id).Error; err != nil {
				log.Fatalln("CheckoutDATATYPE_DEFINITION_INTEGER : Problem with getting object with id:", id)
			}
			backRepo.BackRepoDATATYPE_DEFINITION_INTEGER.CheckoutPhaseOneInstance(&datatype_definition_integerDB)
			backRepo.BackRepoDATATYPE_DEFINITION_INTEGER.CheckoutPhaseTwoInstance(backRepo, &datatype_definition_integerDB)
		}
	}
}

// CopyBasicFieldsFromDATATYPE_DEFINITION_INTEGER
func (datatype_definition_integerDB *DATATYPE_DEFINITION_INTEGERDB) CopyBasicFieldsFromDATATYPE_DEFINITION_INTEGER(datatype_definition_integer *models.DATATYPE_DEFINITION_INTEGER) {
	// insertion point for fields commit

	datatype_definition_integerDB.Name_Data.String = datatype_definition_integer.Name
	datatype_definition_integerDB.Name_Data.Valid = true

	datatype_definition_integerDB.DESC_Data.String = datatype_definition_integer.DESC
	datatype_definition_integerDB.DESC_Data.Valid = true

	datatype_definition_integerDB.IDENTIFIER_Data.String = datatype_definition_integer.IDENTIFIER
	datatype_definition_integerDB.IDENTIFIER_Data.Valid = true

	datatype_definition_integerDB.LAST_CHANGE_Data.String = datatype_definition_integer.LAST_CHANGE
	datatype_definition_integerDB.LAST_CHANGE_Data.Valid = true

	datatype_definition_integerDB.LONG_NAME_Data.String = datatype_definition_integer.LONG_NAME
	datatype_definition_integerDB.LONG_NAME_Data.Valid = true

	datatype_definition_integerDB.MAX_Data.Int64 = int64(datatype_definition_integer.MAX)
	datatype_definition_integerDB.MAX_Data.Valid = true

	datatype_definition_integerDB.MIN_Data.Int64 = int64(datatype_definition_integer.MIN)
	datatype_definition_integerDB.MIN_Data.Valid = true
}

// CopyBasicFieldsFromDATATYPE_DEFINITION_INTEGER_WOP
func (datatype_definition_integerDB *DATATYPE_DEFINITION_INTEGERDB) CopyBasicFieldsFromDATATYPE_DEFINITION_INTEGER_WOP(datatype_definition_integer *models.DATATYPE_DEFINITION_INTEGER_WOP) {
	// insertion point for fields commit

	datatype_definition_integerDB.Name_Data.String = datatype_definition_integer.Name
	datatype_definition_integerDB.Name_Data.Valid = true

	datatype_definition_integerDB.DESC_Data.String = datatype_definition_integer.DESC
	datatype_definition_integerDB.DESC_Data.Valid = true

	datatype_definition_integerDB.IDENTIFIER_Data.String = datatype_definition_integer.IDENTIFIER
	datatype_definition_integerDB.IDENTIFIER_Data.Valid = true

	datatype_definition_integerDB.LAST_CHANGE_Data.String = datatype_definition_integer.LAST_CHANGE
	datatype_definition_integerDB.LAST_CHANGE_Data.Valid = true

	datatype_definition_integerDB.LONG_NAME_Data.String = datatype_definition_integer.LONG_NAME
	datatype_definition_integerDB.LONG_NAME_Data.Valid = true

	datatype_definition_integerDB.MAX_Data.Int64 = int64(datatype_definition_integer.MAX)
	datatype_definition_integerDB.MAX_Data.Valid = true

	datatype_definition_integerDB.MIN_Data.Int64 = int64(datatype_definition_integer.MIN)
	datatype_definition_integerDB.MIN_Data.Valid = true
}

// CopyBasicFieldsFromDATATYPE_DEFINITION_INTEGERWOP
func (datatype_definition_integerDB *DATATYPE_DEFINITION_INTEGERDB) CopyBasicFieldsFromDATATYPE_DEFINITION_INTEGERWOP(datatype_definition_integer *DATATYPE_DEFINITION_INTEGERWOP) {
	// insertion point for fields commit

	datatype_definition_integerDB.Name_Data.String = datatype_definition_integer.Name
	datatype_definition_integerDB.Name_Data.Valid = true

	datatype_definition_integerDB.DESC_Data.String = datatype_definition_integer.DESC
	datatype_definition_integerDB.DESC_Data.Valid = true

	datatype_definition_integerDB.IDENTIFIER_Data.String = datatype_definition_integer.IDENTIFIER
	datatype_definition_integerDB.IDENTIFIER_Data.Valid = true

	datatype_definition_integerDB.LAST_CHANGE_Data.String = datatype_definition_integer.LAST_CHANGE
	datatype_definition_integerDB.LAST_CHANGE_Data.Valid = true

	datatype_definition_integerDB.LONG_NAME_Data.String = datatype_definition_integer.LONG_NAME
	datatype_definition_integerDB.LONG_NAME_Data.Valid = true

	datatype_definition_integerDB.MAX_Data.Int64 = int64(datatype_definition_integer.MAX)
	datatype_definition_integerDB.MAX_Data.Valid = true

	datatype_definition_integerDB.MIN_Data.Int64 = int64(datatype_definition_integer.MIN)
	datatype_definition_integerDB.MIN_Data.Valid = true
}

// CopyBasicFieldsToDATATYPE_DEFINITION_INTEGER
func (datatype_definition_integerDB *DATATYPE_DEFINITION_INTEGERDB) CopyBasicFieldsToDATATYPE_DEFINITION_INTEGER(datatype_definition_integer *models.DATATYPE_DEFINITION_INTEGER) {
	// insertion point for checkout of basic fields (back repo to stage)
	datatype_definition_integer.Name = datatype_definition_integerDB.Name_Data.String
	datatype_definition_integer.DESC = datatype_definition_integerDB.DESC_Data.String
	datatype_definition_integer.IDENTIFIER = datatype_definition_integerDB.IDENTIFIER_Data.String
	datatype_definition_integer.LAST_CHANGE = datatype_definition_integerDB.LAST_CHANGE_Data.String
	datatype_definition_integer.LONG_NAME = datatype_definition_integerDB.LONG_NAME_Data.String
	datatype_definition_integer.MAX = int(datatype_definition_integerDB.MAX_Data.Int64)
	datatype_definition_integer.MIN = int(datatype_definition_integerDB.MIN_Data.Int64)
}

// CopyBasicFieldsToDATATYPE_DEFINITION_INTEGER_WOP
func (datatype_definition_integerDB *DATATYPE_DEFINITION_INTEGERDB) CopyBasicFieldsToDATATYPE_DEFINITION_INTEGER_WOP(datatype_definition_integer *models.DATATYPE_DEFINITION_INTEGER_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	datatype_definition_integer.Name = datatype_definition_integerDB.Name_Data.String
	datatype_definition_integer.DESC = datatype_definition_integerDB.DESC_Data.String
	datatype_definition_integer.IDENTIFIER = datatype_definition_integerDB.IDENTIFIER_Data.String
	datatype_definition_integer.LAST_CHANGE = datatype_definition_integerDB.LAST_CHANGE_Data.String
	datatype_definition_integer.LONG_NAME = datatype_definition_integerDB.LONG_NAME_Data.String
	datatype_definition_integer.MAX = int(datatype_definition_integerDB.MAX_Data.Int64)
	datatype_definition_integer.MIN = int(datatype_definition_integerDB.MIN_Data.Int64)
}

// CopyBasicFieldsToDATATYPE_DEFINITION_INTEGERWOP
func (datatype_definition_integerDB *DATATYPE_DEFINITION_INTEGERDB) CopyBasicFieldsToDATATYPE_DEFINITION_INTEGERWOP(datatype_definition_integer *DATATYPE_DEFINITION_INTEGERWOP) {
	datatype_definition_integer.ID = int(datatype_definition_integerDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	datatype_definition_integer.Name = datatype_definition_integerDB.Name_Data.String
	datatype_definition_integer.DESC = datatype_definition_integerDB.DESC_Data.String
	datatype_definition_integer.IDENTIFIER = datatype_definition_integerDB.IDENTIFIER_Data.String
	datatype_definition_integer.LAST_CHANGE = datatype_definition_integerDB.LAST_CHANGE_Data.String
	datatype_definition_integer.LONG_NAME = datatype_definition_integerDB.LONG_NAME_Data.String
	datatype_definition_integer.MAX = int(datatype_definition_integerDB.MAX_Data.Int64)
	datatype_definition_integer.MIN = int(datatype_definition_integerDB.MIN_Data.Int64)
}

// Backup generates a json file from a slice of all DATATYPE_DEFINITION_INTEGERDB instances in the backrepo
func (backRepoDATATYPE_DEFINITION_INTEGER *BackRepoDATATYPE_DEFINITION_INTEGERStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "DATATYPE_DEFINITION_INTEGERDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*DATATYPE_DEFINITION_INTEGERDB, 0)
	for _, datatype_definition_integerDB := range backRepoDATATYPE_DEFINITION_INTEGER.Map_DATATYPE_DEFINITION_INTEGERDBID_DATATYPE_DEFINITION_INTEGERDB {
		forBackup = append(forBackup, datatype_definition_integerDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json DATATYPE_DEFINITION_INTEGER ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json DATATYPE_DEFINITION_INTEGER file", err.Error())
	}
}

// Backup generates a json file from a slice of all DATATYPE_DEFINITION_INTEGERDB instances in the backrepo
func (backRepoDATATYPE_DEFINITION_INTEGER *BackRepoDATATYPE_DEFINITION_INTEGERStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*DATATYPE_DEFINITION_INTEGERDB, 0)
	for _, datatype_definition_integerDB := range backRepoDATATYPE_DEFINITION_INTEGER.Map_DATATYPE_DEFINITION_INTEGERDBID_DATATYPE_DEFINITION_INTEGERDB {
		forBackup = append(forBackup, datatype_definition_integerDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("DATATYPE_DEFINITION_INTEGER")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&DATATYPE_DEFINITION_INTEGER_Fields, -1)
	for _, datatype_definition_integerDB := range forBackup {

		var datatype_definition_integerWOP DATATYPE_DEFINITION_INTEGERWOP
		datatype_definition_integerDB.CopyBasicFieldsToDATATYPE_DEFINITION_INTEGERWOP(&datatype_definition_integerWOP)

		row := sh.AddRow()
		row.WriteStruct(&datatype_definition_integerWOP, -1)
	}
}

// RestoreXL from the "DATATYPE_DEFINITION_INTEGER" sheet all DATATYPE_DEFINITION_INTEGERDB instances
func (backRepoDATATYPE_DEFINITION_INTEGER *BackRepoDATATYPE_DEFINITION_INTEGERStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoDATATYPE_DEFINITION_INTEGERid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["DATATYPE_DEFINITION_INTEGER"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoDATATYPE_DEFINITION_INTEGER.rowVisitorDATATYPE_DEFINITION_INTEGER)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoDATATYPE_DEFINITION_INTEGER *BackRepoDATATYPE_DEFINITION_INTEGERStruct) rowVisitorDATATYPE_DEFINITION_INTEGER(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var datatype_definition_integerWOP DATATYPE_DEFINITION_INTEGERWOP
		row.ReadStruct(&datatype_definition_integerWOP)

		// add the unmarshalled struct to the stage
		datatype_definition_integerDB := new(DATATYPE_DEFINITION_INTEGERDB)
		datatype_definition_integerDB.CopyBasicFieldsFromDATATYPE_DEFINITION_INTEGERWOP(&datatype_definition_integerWOP)

		datatype_definition_integerDB_ID_atBackupTime := datatype_definition_integerDB.ID
		datatype_definition_integerDB.ID = 0
		query := backRepoDATATYPE_DEFINITION_INTEGER.db.Create(datatype_definition_integerDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoDATATYPE_DEFINITION_INTEGER.Map_DATATYPE_DEFINITION_INTEGERDBID_DATATYPE_DEFINITION_INTEGERDB[datatype_definition_integerDB.ID] = datatype_definition_integerDB
		BackRepoDATATYPE_DEFINITION_INTEGERid_atBckpTime_newID[datatype_definition_integerDB_ID_atBackupTime] = datatype_definition_integerDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "DATATYPE_DEFINITION_INTEGERDB.json" in dirPath that stores an array
// of DATATYPE_DEFINITION_INTEGERDB and stores it in the database
// the map BackRepoDATATYPE_DEFINITION_INTEGERid_atBckpTime_newID is updated accordingly
func (backRepoDATATYPE_DEFINITION_INTEGER *BackRepoDATATYPE_DEFINITION_INTEGERStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoDATATYPE_DEFINITION_INTEGERid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "DATATYPE_DEFINITION_INTEGERDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json DATATYPE_DEFINITION_INTEGER file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*DATATYPE_DEFINITION_INTEGERDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_DATATYPE_DEFINITION_INTEGERDBID_DATATYPE_DEFINITION_INTEGERDB
	for _, datatype_definition_integerDB := range forRestore {

		datatype_definition_integerDB_ID_atBackupTime := datatype_definition_integerDB.ID
		datatype_definition_integerDB.ID = 0
		query := backRepoDATATYPE_DEFINITION_INTEGER.db.Create(datatype_definition_integerDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoDATATYPE_DEFINITION_INTEGER.Map_DATATYPE_DEFINITION_INTEGERDBID_DATATYPE_DEFINITION_INTEGERDB[datatype_definition_integerDB.ID] = datatype_definition_integerDB
		BackRepoDATATYPE_DEFINITION_INTEGERid_atBckpTime_newID[datatype_definition_integerDB_ID_atBackupTime] = datatype_definition_integerDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json DATATYPE_DEFINITION_INTEGER file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<DATATYPE_DEFINITION_INTEGER>id_atBckpTime_newID
// to compute new index
func (backRepoDATATYPE_DEFINITION_INTEGER *BackRepoDATATYPE_DEFINITION_INTEGERStruct) RestorePhaseTwo() {

	for _, datatype_definition_integerDB := range backRepoDATATYPE_DEFINITION_INTEGER.Map_DATATYPE_DEFINITION_INTEGERDBID_DATATYPE_DEFINITION_INTEGERDB {

		// next line of code is to avert unused variable compilation error
		_ = datatype_definition_integerDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		query := backRepoDATATYPE_DEFINITION_INTEGER.db.Model(datatype_definition_integerDB).Updates(*datatype_definition_integerDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
	}

}

// BackRepoDATATYPE_DEFINITION_INTEGER.ResetReversePointers commits all staged instances of DATATYPE_DEFINITION_INTEGER to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoDATATYPE_DEFINITION_INTEGER *BackRepoDATATYPE_DEFINITION_INTEGERStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, datatype_definition_integer := range backRepoDATATYPE_DEFINITION_INTEGER.Map_DATATYPE_DEFINITION_INTEGERDBID_DATATYPE_DEFINITION_INTEGERPtr {
		backRepoDATATYPE_DEFINITION_INTEGER.ResetReversePointersInstance(backRepo, idx, datatype_definition_integer)
	}

	return
}

func (backRepoDATATYPE_DEFINITION_INTEGER *BackRepoDATATYPE_DEFINITION_INTEGERStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, datatype_definition_integer *models.DATATYPE_DEFINITION_INTEGER) (Error error) {

	// fetch matching datatype_definition_integerDB
	if datatype_definition_integerDB, ok := backRepoDATATYPE_DEFINITION_INTEGER.Map_DATATYPE_DEFINITION_INTEGERDBID_DATATYPE_DEFINITION_INTEGERDB[idx]; ok {
		_ = datatype_definition_integerDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoDATATYPE_DEFINITION_INTEGERid_atBckpTime_newID map[uint]uint