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
var dummy_DATATYPE_DEFINITION_STRING_sql sql.NullBool
var dummy_DATATYPE_DEFINITION_STRING_time time.Duration
var dummy_DATATYPE_DEFINITION_STRING_sort sort.Float64Slice

// DATATYPE_DEFINITION_STRINGAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model datatype_definition_stringAPI
type DATATYPE_DEFINITION_STRINGAPI struct {
	gorm.Model

	models.DATATYPE_DEFINITION_STRING_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	DATATYPE_DEFINITION_STRINGPointersEncoding DATATYPE_DEFINITION_STRINGPointersEncoding
}

// DATATYPE_DEFINITION_STRINGPointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type DATATYPE_DEFINITION_STRINGPointersEncoding struct {
	// insertion for pointer fields encoding declaration
}

// DATATYPE_DEFINITION_STRINGDB describes a datatype_definition_string in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model datatype_definition_stringDB
type DATATYPE_DEFINITION_STRINGDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field datatype_definition_stringDB.Name
	Name_Data sql.NullString

	// Declation for basic field datatype_definition_stringDB.DESC
	DESC_Data sql.NullString

	// Declation for basic field datatype_definition_stringDB.IDENTIFIER
	IDENTIFIER_Data sql.NullString

	// Declation for basic field datatype_definition_stringDB.LAST_CHANGE
	LAST_CHANGE_Data sql.NullString

	// Declation for basic field datatype_definition_stringDB.LONG_NAME
	LONG_NAME_Data sql.NullString

	// Declation for basic field datatype_definition_stringDB.MAX_LENGTH
	MAX_LENGTH_Data sql.NullInt64
	
	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	DATATYPE_DEFINITION_STRINGPointersEncoding
}

// DATATYPE_DEFINITION_STRINGDBs arrays datatype_definition_stringDBs
// swagger:response datatype_definition_stringDBsResponse
type DATATYPE_DEFINITION_STRINGDBs []DATATYPE_DEFINITION_STRINGDB

// DATATYPE_DEFINITION_STRINGDBResponse provides response
// swagger:response datatype_definition_stringDBResponse
type DATATYPE_DEFINITION_STRINGDBResponse struct {
	DATATYPE_DEFINITION_STRINGDB
}

// DATATYPE_DEFINITION_STRINGWOP is a DATATYPE_DEFINITION_STRING without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type DATATYPE_DEFINITION_STRINGWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	DESC string `xlsx:"2"`

	IDENTIFIER string `xlsx:"3"`

	LAST_CHANGE string `xlsx:"4"`

	LONG_NAME string `xlsx:"5"`

	MAX_LENGTH int `xlsx:"6"`
	// insertion for WOP pointer fields
}

var DATATYPE_DEFINITION_STRING_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"DESC",
	"IDENTIFIER",
	"LAST_CHANGE",
	"LONG_NAME",
	"MAX_LENGTH",
}

type BackRepoDATATYPE_DEFINITION_STRINGStruct struct {
	// stores DATATYPE_DEFINITION_STRINGDB according to their gorm ID
	Map_DATATYPE_DEFINITION_STRINGDBID_DATATYPE_DEFINITION_STRINGDB map[uint]*DATATYPE_DEFINITION_STRINGDB

	// stores DATATYPE_DEFINITION_STRINGDB ID according to DATATYPE_DEFINITION_STRING address
	Map_DATATYPE_DEFINITION_STRINGPtr_DATATYPE_DEFINITION_STRINGDBID map[*models.DATATYPE_DEFINITION_STRING]uint

	// stores DATATYPE_DEFINITION_STRING according to their gorm ID
	Map_DATATYPE_DEFINITION_STRINGDBID_DATATYPE_DEFINITION_STRINGPtr map[uint]*models.DATATYPE_DEFINITION_STRING

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoDATATYPE_DEFINITION_STRING *BackRepoDATATYPE_DEFINITION_STRINGStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoDATATYPE_DEFINITION_STRING.stage
	return
}

func (backRepoDATATYPE_DEFINITION_STRING *BackRepoDATATYPE_DEFINITION_STRINGStruct) GetDB() *gorm.DB {
	return backRepoDATATYPE_DEFINITION_STRING.db
}

// GetDATATYPE_DEFINITION_STRINGDBFromDATATYPE_DEFINITION_STRINGPtr is a handy function to access the back repo instance from the stage instance
func (backRepoDATATYPE_DEFINITION_STRING *BackRepoDATATYPE_DEFINITION_STRINGStruct) GetDATATYPE_DEFINITION_STRINGDBFromDATATYPE_DEFINITION_STRINGPtr(datatype_definition_string *models.DATATYPE_DEFINITION_STRING) (datatype_definition_stringDB *DATATYPE_DEFINITION_STRINGDB) {
	id := backRepoDATATYPE_DEFINITION_STRING.Map_DATATYPE_DEFINITION_STRINGPtr_DATATYPE_DEFINITION_STRINGDBID[datatype_definition_string]
	datatype_definition_stringDB = backRepoDATATYPE_DEFINITION_STRING.Map_DATATYPE_DEFINITION_STRINGDBID_DATATYPE_DEFINITION_STRINGDB[id]
	return
}

// BackRepoDATATYPE_DEFINITION_STRING.CommitPhaseOne commits all staged instances of DATATYPE_DEFINITION_STRING to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoDATATYPE_DEFINITION_STRING *BackRepoDATATYPE_DEFINITION_STRINGStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for datatype_definition_string := range stage.DATATYPE_DEFINITION_STRINGs {
		backRepoDATATYPE_DEFINITION_STRING.CommitPhaseOneInstance(datatype_definition_string)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, datatype_definition_string := range backRepoDATATYPE_DEFINITION_STRING.Map_DATATYPE_DEFINITION_STRINGDBID_DATATYPE_DEFINITION_STRINGPtr {
		if _, ok := stage.DATATYPE_DEFINITION_STRINGs[datatype_definition_string]; !ok {
			backRepoDATATYPE_DEFINITION_STRING.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoDATATYPE_DEFINITION_STRING.CommitDeleteInstance commits deletion of DATATYPE_DEFINITION_STRING to the BackRepo
func (backRepoDATATYPE_DEFINITION_STRING *BackRepoDATATYPE_DEFINITION_STRINGStruct) CommitDeleteInstance(id uint) (Error error) {

	datatype_definition_string := backRepoDATATYPE_DEFINITION_STRING.Map_DATATYPE_DEFINITION_STRINGDBID_DATATYPE_DEFINITION_STRINGPtr[id]

	// datatype_definition_string is not staged anymore, remove datatype_definition_stringDB
	datatype_definition_stringDB := backRepoDATATYPE_DEFINITION_STRING.Map_DATATYPE_DEFINITION_STRINGDBID_DATATYPE_DEFINITION_STRINGDB[id]
	query := backRepoDATATYPE_DEFINITION_STRING.db.Unscoped().Delete(&datatype_definition_stringDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	delete(backRepoDATATYPE_DEFINITION_STRING.Map_DATATYPE_DEFINITION_STRINGPtr_DATATYPE_DEFINITION_STRINGDBID, datatype_definition_string)
	delete(backRepoDATATYPE_DEFINITION_STRING.Map_DATATYPE_DEFINITION_STRINGDBID_DATATYPE_DEFINITION_STRINGPtr, id)
	delete(backRepoDATATYPE_DEFINITION_STRING.Map_DATATYPE_DEFINITION_STRINGDBID_DATATYPE_DEFINITION_STRINGDB, id)

	return
}

// BackRepoDATATYPE_DEFINITION_STRING.CommitPhaseOneInstance commits datatype_definition_string staged instances of DATATYPE_DEFINITION_STRING to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoDATATYPE_DEFINITION_STRING *BackRepoDATATYPE_DEFINITION_STRINGStruct) CommitPhaseOneInstance(datatype_definition_string *models.DATATYPE_DEFINITION_STRING) (Error error) {

	// check if the datatype_definition_string is not commited yet
	if _, ok := backRepoDATATYPE_DEFINITION_STRING.Map_DATATYPE_DEFINITION_STRINGPtr_DATATYPE_DEFINITION_STRINGDBID[datatype_definition_string]; ok {
		return
	}

	// initiate datatype_definition_string
	var datatype_definition_stringDB DATATYPE_DEFINITION_STRINGDB
	datatype_definition_stringDB.CopyBasicFieldsFromDATATYPE_DEFINITION_STRING(datatype_definition_string)

	query := backRepoDATATYPE_DEFINITION_STRING.db.Create(&datatype_definition_stringDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	backRepoDATATYPE_DEFINITION_STRING.Map_DATATYPE_DEFINITION_STRINGPtr_DATATYPE_DEFINITION_STRINGDBID[datatype_definition_string] = datatype_definition_stringDB.ID
	backRepoDATATYPE_DEFINITION_STRING.Map_DATATYPE_DEFINITION_STRINGDBID_DATATYPE_DEFINITION_STRINGPtr[datatype_definition_stringDB.ID] = datatype_definition_string
	backRepoDATATYPE_DEFINITION_STRING.Map_DATATYPE_DEFINITION_STRINGDBID_DATATYPE_DEFINITION_STRINGDB[datatype_definition_stringDB.ID] = &datatype_definition_stringDB

	return
}

// BackRepoDATATYPE_DEFINITION_STRING.CommitPhaseTwo commits all staged instances of DATATYPE_DEFINITION_STRING to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoDATATYPE_DEFINITION_STRING *BackRepoDATATYPE_DEFINITION_STRINGStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, datatype_definition_string := range backRepoDATATYPE_DEFINITION_STRING.Map_DATATYPE_DEFINITION_STRINGDBID_DATATYPE_DEFINITION_STRINGPtr {
		backRepoDATATYPE_DEFINITION_STRING.CommitPhaseTwoInstance(backRepo, idx, datatype_definition_string)
	}

	return
}

// BackRepoDATATYPE_DEFINITION_STRING.CommitPhaseTwoInstance commits {{structname }} of models.DATATYPE_DEFINITION_STRING to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoDATATYPE_DEFINITION_STRING *BackRepoDATATYPE_DEFINITION_STRINGStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, datatype_definition_string *models.DATATYPE_DEFINITION_STRING) (Error error) {

	// fetch matching datatype_definition_stringDB
	if datatype_definition_stringDB, ok := backRepoDATATYPE_DEFINITION_STRING.Map_DATATYPE_DEFINITION_STRINGDBID_DATATYPE_DEFINITION_STRINGDB[idx]; ok {

		datatype_definition_stringDB.CopyBasicFieldsFromDATATYPE_DEFINITION_STRING(datatype_definition_string)

		// insertion point for translating pointers encodings into actual pointers
		query := backRepoDATATYPE_DEFINITION_STRING.db.Save(&datatype_definition_stringDB)
		if query.Error != nil {
			log.Fatalln(query.Error)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown DATATYPE_DEFINITION_STRING intance %s", datatype_definition_string.Name))
		return err
	}

	return
}

// BackRepoDATATYPE_DEFINITION_STRING.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoDATATYPE_DEFINITION_STRING *BackRepoDATATYPE_DEFINITION_STRINGStruct) CheckoutPhaseOne() (Error error) {

	datatype_definition_stringDBArray := make([]DATATYPE_DEFINITION_STRINGDB, 0)
	query := backRepoDATATYPE_DEFINITION_STRING.db.Find(&datatype_definition_stringDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	datatype_definition_stringInstancesToBeRemovedFromTheStage := make(map[*models.DATATYPE_DEFINITION_STRING]any)
	for key, value := range backRepoDATATYPE_DEFINITION_STRING.stage.DATATYPE_DEFINITION_STRINGs {
		datatype_definition_stringInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, datatype_definition_stringDB := range datatype_definition_stringDBArray {
		backRepoDATATYPE_DEFINITION_STRING.CheckoutPhaseOneInstance(&datatype_definition_stringDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		datatype_definition_string, ok := backRepoDATATYPE_DEFINITION_STRING.Map_DATATYPE_DEFINITION_STRINGDBID_DATATYPE_DEFINITION_STRINGPtr[datatype_definition_stringDB.ID]
		if ok {
			delete(datatype_definition_stringInstancesToBeRemovedFromTheStage, datatype_definition_string)
		}
	}

	// remove from stage and back repo's 3 maps all datatype_definition_strings that are not in the checkout
	for datatype_definition_string := range datatype_definition_stringInstancesToBeRemovedFromTheStage {
		datatype_definition_string.Unstage(backRepoDATATYPE_DEFINITION_STRING.GetStage())

		// remove instance from the back repo 3 maps
		datatype_definition_stringID := backRepoDATATYPE_DEFINITION_STRING.Map_DATATYPE_DEFINITION_STRINGPtr_DATATYPE_DEFINITION_STRINGDBID[datatype_definition_string]
		delete(backRepoDATATYPE_DEFINITION_STRING.Map_DATATYPE_DEFINITION_STRINGPtr_DATATYPE_DEFINITION_STRINGDBID, datatype_definition_string)
		delete(backRepoDATATYPE_DEFINITION_STRING.Map_DATATYPE_DEFINITION_STRINGDBID_DATATYPE_DEFINITION_STRINGDB, datatype_definition_stringID)
		delete(backRepoDATATYPE_DEFINITION_STRING.Map_DATATYPE_DEFINITION_STRINGDBID_DATATYPE_DEFINITION_STRINGPtr, datatype_definition_stringID)
	}

	return
}

// CheckoutPhaseOneInstance takes a datatype_definition_stringDB that has been found in the DB, updates the backRepo and stages the
// models version of the datatype_definition_stringDB
func (backRepoDATATYPE_DEFINITION_STRING *BackRepoDATATYPE_DEFINITION_STRINGStruct) CheckoutPhaseOneInstance(datatype_definition_stringDB *DATATYPE_DEFINITION_STRINGDB) (Error error) {

	datatype_definition_string, ok := backRepoDATATYPE_DEFINITION_STRING.Map_DATATYPE_DEFINITION_STRINGDBID_DATATYPE_DEFINITION_STRINGPtr[datatype_definition_stringDB.ID]
	if !ok {
		datatype_definition_string = new(models.DATATYPE_DEFINITION_STRING)

		backRepoDATATYPE_DEFINITION_STRING.Map_DATATYPE_DEFINITION_STRINGDBID_DATATYPE_DEFINITION_STRINGPtr[datatype_definition_stringDB.ID] = datatype_definition_string
		backRepoDATATYPE_DEFINITION_STRING.Map_DATATYPE_DEFINITION_STRINGPtr_DATATYPE_DEFINITION_STRINGDBID[datatype_definition_string] = datatype_definition_stringDB.ID

		// append model store with the new element
		datatype_definition_string.Name = datatype_definition_stringDB.Name_Data.String
		datatype_definition_string.Stage(backRepoDATATYPE_DEFINITION_STRING.GetStage())
	}
	datatype_definition_stringDB.CopyBasicFieldsToDATATYPE_DEFINITION_STRING(datatype_definition_string)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	datatype_definition_string.Stage(backRepoDATATYPE_DEFINITION_STRING.GetStage())

	// preserve pointer to datatype_definition_stringDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_DATATYPE_DEFINITION_STRINGDBID_DATATYPE_DEFINITION_STRINGDB)[datatype_definition_stringDB hold variable pointers
	datatype_definition_stringDB_Data := *datatype_definition_stringDB
	preservedPtrToDATATYPE_DEFINITION_STRING := &datatype_definition_stringDB_Data
	backRepoDATATYPE_DEFINITION_STRING.Map_DATATYPE_DEFINITION_STRINGDBID_DATATYPE_DEFINITION_STRINGDB[datatype_definition_stringDB.ID] = preservedPtrToDATATYPE_DEFINITION_STRING

	return
}

// BackRepoDATATYPE_DEFINITION_STRING.CheckoutPhaseTwo Checkouts all staged instances of DATATYPE_DEFINITION_STRING to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoDATATYPE_DEFINITION_STRING *BackRepoDATATYPE_DEFINITION_STRINGStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, datatype_definition_stringDB := range backRepoDATATYPE_DEFINITION_STRING.Map_DATATYPE_DEFINITION_STRINGDBID_DATATYPE_DEFINITION_STRINGDB {
		backRepoDATATYPE_DEFINITION_STRING.CheckoutPhaseTwoInstance(backRepo, datatype_definition_stringDB)
	}
	return
}

// BackRepoDATATYPE_DEFINITION_STRING.CheckoutPhaseTwoInstance Checkouts staged instances of DATATYPE_DEFINITION_STRING to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoDATATYPE_DEFINITION_STRING *BackRepoDATATYPE_DEFINITION_STRINGStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, datatype_definition_stringDB *DATATYPE_DEFINITION_STRINGDB) (Error error) {

	datatype_definition_string := backRepoDATATYPE_DEFINITION_STRING.Map_DATATYPE_DEFINITION_STRINGDBID_DATATYPE_DEFINITION_STRINGPtr[datatype_definition_stringDB.ID]

	datatype_definition_stringDB.DecodePointers(backRepo, datatype_definition_string)

	return
}

func (datatype_definition_stringDB *DATATYPE_DEFINITION_STRINGDB) DecodePointers(backRepo *BackRepoStruct, datatype_definition_string *models.DATATYPE_DEFINITION_STRING) {

	// insertion point for checkout of pointer encoding
	return
}

// CommitDATATYPE_DEFINITION_STRING allows commit of a single datatype_definition_string (if already staged)
func (backRepo *BackRepoStruct) CommitDATATYPE_DEFINITION_STRING(datatype_definition_string *models.DATATYPE_DEFINITION_STRING) {
	backRepo.BackRepoDATATYPE_DEFINITION_STRING.CommitPhaseOneInstance(datatype_definition_string)
	if id, ok := backRepo.BackRepoDATATYPE_DEFINITION_STRING.Map_DATATYPE_DEFINITION_STRINGPtr_DATATYPE_DEFINITION_STRINGDBID[datatype_definition_string]; ok {
		backRepo.BackRepoDATATYPE_DEFINITION_STRING.CommitPhaseTwoInstance(backRepo, id, datatype_definition_string)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitDATATYPE_DEFINITION_STRING allows checkout of a single datatype_definition_string (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutDATATYPE_DEFINITION_STRING(datatype_definition_string *models.DATATYPE_DEFINITION_STRING) {
	// check if the datatype_definition_string is staged
	if _, ok := backRepo.BackRepoDATATYPE_DEFINITION_STRING.Map_DATATYPE_DEFINITION_STRINGPtr_DATATYPE_DEFINITION_STRINGDBID[datatype_definition_string]; ok {

		if id, ok := backRepo.BackRepoDATATYPE_DEFINITION_STRING.Map_DATATYPE_DEFINITION_STRINGPtr_DATATYPE_DEFINITION_STRINGDBID[datatype_definition_string]; ok {
			var datatype_definition_stringDB DATATYPE_DEFINITION_STRINGDB
			datatype_definition_stringDB.ID = id

			if err := backRepo.BackRepoDATATYPE_DEFINITION_STRING.db.First(&datatype_definition_stringDB, id).Error; err != nil {
				log.Fatalln("CheckoutDATATYPE_DEFINITION_STRING : Problem with getting object with id:", id)
			}
			backRepo.BackRepoDATATYPE_DEFINITION_STRING.CheckoutPhaseOneInstance(&datatype_definition_stringDB)
			backRepo.BackRepoDATATYPE_DEFINITION_STRING.CheckoutPhaseTwoInstance(backRepo, &datatype_definition_stringDB)
		}
	}
}

// CopyBasicFieldsFromDATATYPE_DEFINITION_STRING
func (datatype_definition_stringDB *DATATYPE_DEFINITION_STRINGDB) CopyBasicFieldsFromDATATYPE_DEFINITION_STRING(datatype_definition_string *models.DATATYPE_DEFINITION_STRING) {
	// insertion point for fields commit

	datatype_definition_stringDB.Name_Data.String = datatype_definition_string.Name
	datatype_definition_stringDB.Name_Data.Valid = true

	datatype_definition_stringDB.DESC_Data.String = datatype_definition_string.DESC
	datatype_definition_stringDB.DESC_Data.Valid = true

	datatype_definition_stringDB.IDENTIFIER_Data.String = datatype_definition_string.IDENTIFIER
	datatype_definition_stringDB.IDENTIFIER_Data.Valid = true

	datatype_definition_stringDB.LAST_CHANGE_Data.String = datatype_definition_string.LAST_CHANGE
	datatype_definition_stringDB.LAST_CHANGE_Data.Valid = true

	datatype_definition_stringDB.LONG_NAME_Data.String = datatype_definition_string.LONG_NAME
	datatype_definition_stringDB.LONG_NAME_Data.Valid = true

	datatype_definition_stringDB.MAX_LENGTH_Data.Int64 = int64(datatype_definition_string.MAX_LENGTH)
	datatype_definition_stringDB.MAX_LENGTH_Data.Valid = true
}

// CopyBasicFieldsFromDATATYPE_DEFINITION_STRING_WOP
func (datatype_definition_stringDB *DATATYPE_DEFINITION_STRINGDB) CopyBasicFieldsFromDATATYPE_DEFINITION_STRING_WOP(datatype_definition_string *models.DATATYPE_DEFINITION_STRING_WOP) {
	// insertion point for fields commit

	datatype_definition_stringDB.Name_Data.String = datatype_definition_string.Name
	datatype_definition_stringDB.Name_Data.Valid = true

	datatype_definition_stringDB.DESC_Data.String = datatype_definition_string.DESC
	datatype_definition_stringDB.DESC_Data.Valid = true

	datatype_definition_stringDB.IDENTIFIER_Data.String = datatype_definition_string.IDENTIFIER
	datatype_definition_stringDB.IDENTIFIER_Data.Valid = true

	datatype_definition_stringDB.LAST_CHANGE_Data.String = datatype_definition_string.LAST_CHANGE
	datatype_definition_stringDB.LAST_CHANGE_Data.Valid = true

	datatype_definition_stringDB.LONG_NAME_Data.String = datatype_definition_string.LONG_NAME
	datatype_definition_stringDB.LONG_NAME_Data.Valid = true

	datatype_definition_stringDB.MAX_LENGTH_Data.Int64 = int64(datatype_definition_string.MAX_LENGTH)
	datatype_definition_stringDB.MAX_LENGTH_Data.Valid = true
}

// CopyBasicFieldsFromDATATYPE_DEFINITION_STRINGWOP
func (datatype_definition_stringDB *DATATYPE_DEFINITION_STRINGDB) CopyBasicFieldsFromDATATYPE_DEFINITION_STRINGWOP(datatype_definition_string *DATATYPE_DEFINITION_STRINGWOP) {
	// insertion point for fields commit

	datatype_definition_stringDB.Name_Data.String = datatype_definition_string.Name
	datatype_definition_stringDB.Name_Data.Valid = true

	datatype_definition_stringDB.DESC_Data.String = datatype_definition_string.DESC
	datatype_definition_stringDB.DESC_Data.Valid = true

	datatype_definition_stringDB.IDENTIFIER_Data.String = datatype_definition_string.IDENTIFIER
	datatype_definition_stringDB.IDENTIFIER_Data.Valid = true

	datatype_definition_stringDB.LAST_CHANGE_Data.String = datatype_definition_string.LAST_CHANGE
	datatype_definition_stringDB.LAST_CHANGE_Data.Valid = true

	datatype_definition_stringDB.LONG_NAME_Data.String = datatype_definition_string.LONG_NAME
	datatype_definition_stringDB.LONG_NAME_Data.Valid = true

	datatype_definition_stringDB.MAX_LENGTH_Data.Int64 = int64(datatype_definition_string.MAX_LENGTH)
	datatype_definition_stringDB.MAX_LENGTH_Data.Valid = true
}

// CopyBasicFieldsToDATATYPE_DEFINITION_STRING
func (datatype_definition_stringDB *DATATYPE_DEFINITION_STRINGDB) CopyBasicFieldsToDATATYPE_DEFINITION_STRING(datatype_definition_string *models.DATATYPE_DEFINITION_STRING) {
	// insertion point for checkout of basic fields (back repo to stage)
	datatype_definition_string.Name = datatype_definition_stringDB.Name_Data.String
	datatype_definition_string.DESC = datatype_definition_stringDB.DESC_Data.String
	datatype_definition_string.IDENTIFIER = datatype_definition_stringDB.IDENTIFIER_Data.String
	datatype_definition_string.LAST_CHANGE = datatype_definition_stringDB.LAST_CHANGE_Data.String
	datatype_definition_string.LONG_NAME = datatype_definition_stringDB.LONG_NAME_Data.String
	datatype_definition_string.MAX_LENGTH = int(datatype_definition_stringDB.MAX_LENGTH_Data.Int64)
}

// CopyBasicFieldsToDATATYPE_DEFINITION_STRING_WOP
func (datatype_definition_stringDB *DATATYPE_DEFINITION_STRINGDB) CopyBasicFieldsToDATATYPE_DEFINITION_STRING_WOP(datatype_definition_string *models.DATATYPE_DEFINITION_STRING_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	datatype_definition_string.Name = datatype_definition_stringDB.Name_Data.String
	datatype_definition_string.DESC = datatype_definition_stringDB.DESC_Data.String
	datatype_definition_string.IDENTIFIER = datatype_definition_stringDB.IDENTIFIER_Data.String
	datatype_definition_string.LAST_CHANGE = datatype_definition_stringDB.LAST_CHANGE_Data.String
	datatype_definition_string.LONG_NAME = datatype_definition_stringDB.LONG_NAME_Data.String
	datatype_definition_string.MAX_LENGTH = int(datatype_definition_stringDB.MAX_LENGTH_Data.Int64)
}

// CopyBasicFieldsToDATATYPE_DEFINITION_STRINGWOP
func (datatype_definition_stringDB *DATATYPE_DEFINITION_STRINGDB) CopyBasicFieldsToDATATYPE_DEFINITION_STRINGWOP(datatype_definition_string *DATATYPE_DEFINITION_STRINGWOP) {
	datatype_definition_string.ID = int(datatype_definition_stringDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	datatype_definition_string.Name = datatype_definition_stringDB.Name_Data.String
	datatype_definition_string.DESC = datatype_definition_stringDB.DESC_Data.String
	datatype_definition_string.IDENTIFIER = datatype_definition_stringDB.IDENTIFIER_Data.String
	datatype_definition_string.LAST_CHANGE = datatype_definition_stringDB.LAST_CHANGE_Data.String
	datatype_definition_string.LONG_NAME = datatype_definition_stringDB.LONG_NAME_Data.String
	datatype_definition_string.MAX_LENGTH = int(datatype_definition_stringDB.MAX_LENGTH_Data.Int64)
}

// Backup generates a json file from a slice of all DATATYPE_DEFINITION_STRINGDB instances in the backrepo
func (backRepoDATATYPE_DEFINITION_STRING *BackRepoDATATYPE_DEFINITION_STRINGStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "DATATYPE_DEFINITION_STRINGDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*DATATYPE_DEFINITION_STRINGDB, 0)
	for _, datatype_definition_stringDB := range backRepoDATATYPE_DEFINITION_STRING.Map_DATATYPE_DEFINITION_STRINGDBID_DATATYPE_DEFINITION_STRINGDB {
		forBackup = append(forBackup, datatype_definition_stringDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json DATATYPE_DEFINITION_STRING ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json DATATYPE_DEFINITION_STRING file", err.Error())
	}
}

// Backup generates a json file from a slice of all DATATYPE_DEFINITION_STRINGDB instances in the backrepo
func (backRepoDATATYPE_DEFINITION_STRING *BackRepoDATATYPE_DEFINITION_STRINGStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*DATATYPE_DEFINITION_STRINGDB, 0)
	for _, datatype_definition_stringDB := range backRepoDATATYPE_DEFINITION_STRING.Map_DATATYPE_DEFINITION_STRINGDBID_DATATYPE_DEFINITION_STRINGDB {
		forBackup = append(forBackup, datatype_definition_stringDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("DATATYPE_DEFINITION_STRING")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&DATATYPE_DEFINITION_STRING_Fields, -1)
	for _, datatype_definition_stringDB := range forBackup {

		var datatype_definition_stringWOP DATATYPE_DEFINITION_STRINGWOP
		datatype_definition_stringDB.CopyBasicFieldsToDATATYPE_DEFINITION_STRINGWOP(&datatype_definition_stringWOP)

		row := sh.AddRow()
		row.WriteStruct(&datatype_definition_stringWOP, -1)
	}
}

// RestoreXL from the "DATATYPE_DEFINITION_STRING" sheet all DATATYPE_DEFINITION_STRINGDB instances
func (backRepoDATATYPE_DEFINITION_STRING *BackRepoDATATYPE_DEFINITION_STRINGStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoDATATYPE_DEFINITION_STRINGid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["DATATYPE_DEFINITION_STRING"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoDATATYPE_DEFINITION_STRING.rowVisitorDATATYPE_DEFINITION_STRING)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoDATATYPE_DEFINITION_STRING *BackRepoDATATYPE_DEFINITION_STRINGStruct) rowVisitorDATATYPE_DEFINITION_STRING(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var datatype_definition_stringWOP DATATYPE_DEFINITION_STRINGWOP
		row.ReadStruct(&datatype_definition_stringWOP)

		// add the unmarshalled struct to the stage
		datatype_definition_stringDB := new(DATATYPE_DEFINITION_STRINGDB)
		datatype_definition_stringDB.CopyBasicFieldsFromDATATYPE_DEFINITION_STRINGWOP(&datatype_definition_stringWOP)

		datatype_definition_stringDB_ID_atBackupTime := datatype_definition_stringDB.ID
		datatype_definition_stringDB.ID = 0
		query := backRepoDATATYPE_DEFINITION_STRING.db.Create(datatype_definition_stringDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoDATATYPE_DEFINITION_STRING.Map_DATATYPE_DEFINITION_STRINGDBID_DATATYPE_DEFINITION_STRINGDB[datatype_definition_stringDB.ID] = datatype_definition_stringDB
		BackRepoDATATYPE_DEFINITION_STRINGid_atBckpTime_newID[datatype_definition_stringDB_ID_atBackupTime] = datatype_definition_stringDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "DATATYPE_DEFINITION_STRINGDB.json" in dirPath that stores an array
// of DATATYPE_DEFINITION_STRINGDB and stores it in the database
// the map BackRepoDATATYPE_DEFINITION_STRINGid_atBckpTime_newID is updated accordingly
func (backRepoDATATYPE_DEFINITION_STRING *BackRepoDATATYPE_DEFINITION_STRINGStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoDATATYPE_DEFINITION_STRINGid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "DATATYPE_DEFINITION_STRINGDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json DATATYPE_DEFINITION_STRING file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*DATATYPE_DEFINITION_STRINGDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_DATATYPE_DEFINITION_STRINGDBID_DATATYPE_DEFINITION_STRINGDB
	for _, datatype_definition_stringDB := range forRestore {

		datatype_definition_stringDB_ID_atBackupTime := datatype_definition_stringDB.ID
		datatype_definition_stringDB.ID = 0
		query := backRepoDATATYPE_DEFINITION_STRING.db.Create(datatype_definition_stringDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoDATATYPE_DEFINITION_STRING.Map_DATATYPE_DEFINITION_STRINGDBID_DATATYPE_DEFINITION_STRINGDB[datatype_definition_stringDB.ID] = datatype_definition_stringDB
		BackRepoDATATYPE_DEFINITION_STRINGid_atBckpTime_newID[datatype_definition_stringDB_ID_atBackupTime] = datatype_definition_stringDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json DATATYPE_DEFINITION_STRING file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<DATATYPE_DEFINITION_STRING>id_atBckpTime_newID
// to compute new index
func (backRepoDATATYPE_DEFINITION_STRING *BackRepoDATATYPE_DEFINITION_STRINGStruct) RestorePhaseTwo() {

	for _, datatype_definition_stringDB := range backRepoDATATYPE_DEFINITION_STRING.Map_DATATYPE_DEFINITION_STRINGDBID_DATATYPE_DEFINITION_STRINGDB {

		// next line of code is to avert unused variable compilation error
		_ = datatype_definition_stringDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		query := backRepoDATATYPE_DEFINITION_STRING.db.Model(datatype_definition_stringDB).Updates(*datatype_definition_stringDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
	}

}

// BackRepoDATATYPE_DEFINITION_STRING.ResetReversePointers commits all staged instances of DATATYPE_DEFINITION_STRING to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoDATATYPE_DEFINITION_STRING *BackRepoDATATYPE_DEFINITION_STRINGStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, datatype_definition_string := range backRepoDATATYPE_DEFINITION_STRING.Map_DATATYPE_DEFINITION_STRINGDBID_DATATYPE_DEFINITION_STRINGPtr {
		backRepoDATATYPE_DEFINITION_STRING.ResetReversePointersInstance(backRepo, idx, datatype_definition_string)
	}

	return
}

func (backRepoDATATYPE_DEFINITION_STRING *BackRepoDATATYPE_DEFINITION_STRINGStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, datatype_definition_string *models.DATATYPE_DEFINITION_STRING) (Error error) {

	// fetch matching datatype_definition_stringDB
	if datatype_definition_stringDB, ok := backRepoDATATYPE_DEFINITION_STRING.Map_DATATYPE_DEFINITION_STRINGDBID_DATATYPE_DEFINITION_STRINGDB[idx]; ok {
		_ = datatype_definition_stringDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoDATATYPE_DEFINITION_STRINGid_atBckpTime_newID map[uint]uint