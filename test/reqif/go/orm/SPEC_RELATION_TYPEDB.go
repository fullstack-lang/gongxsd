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
var dummy_SPEC_RELATION_TYPE_sql sql.NullBool
var dummy_SPEC_RELATION_TYPE_time time.Duration
var dummy_SPEC_RELATION_TYPE_sort sort.Float64Slice

// SPEC_RELATION_TYPEAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model spec_relation_typeAPI
type SPEC_RELATION_TYPEAPI struct {
	gorm.Model

	models.SPEC_RELATION_TYPE_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	SPEC_RELATION_TYPEPointersEncoding SPEC_RELATION_TYPEPointersEncoding
}

// SPEC_RELATION_TYPEPointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type SPEC_RELATION_TYPEPointersEncoding struct {
	// insertion for pointer fields encoding declaration
}

// SPEC_RELATION_TYPEDB describes a spec_relation_type in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model spec_relation_typeDB
type SPEC_RELATION_TYPEDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field spec_relation_typeDB.Name
	Name_Data sql.NullString

	// Declation for basic field spec_relation_typeDB.DESC
	DESC_Data sql.NullString

	// Declation for basic field spec_relation_typeDB.IDENTIFIER
	IDENTIFIER_Data sql.NullString

	// Declation for basic field spec_relation_typeDB.LAST_CHANGE
	LAST_CHANGE_Data sql.NullString

	// Declation for basic field spec_relation_typeDB.LONG_NAME
	LONG_NAME_Data sql.NullString
	
	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	SPEC_RELATION_TYPEPointersEncoding
}

// SPEC_RELATION_TYPEDBs arrays spec_relation_typeDBs
// swagger:response spec_relation_typeDBsResponse
type SPEC_RELATION_TYPEDBs []SPEC_RELATION_TYPEDB

// SPEC_RELATION_TYPEDBResponse provides response
// swagger:response spec_relation_typeDBResponse
type SPEC_RELATION_TYPEDBResponse struct {
	SPEC_RELATION_TYPEDB
}

// SPEC_RELATION_TYPEWOP is a SPEC_RELATION_TYPE without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type SPEC_RELATION_TYPEWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	DESC string `xlsx:"2"`

	IDENTIFIER string `xlsx:"3"`

	LAST_CHANGE string `xlsx:"4"`

	LONG_NAME string `xlsx:"5"`
	// insertion for WOP pointer fields
}

var SPEC_RELATION_TYPE_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"DESC",
	"IDENTIFIER",
	"LAST_CHANGE",
	"LONG_NAME",
}

type BackRepoSPEC_RELATION_TYPEStruct struct {
	// stores SPEC_RELATION_TYPEDB according to their gorm ID
	Map_SPEC_RELATION_TYPEDBID_SPEC_RELATION_TYPEDB map[uint]*SPEC_RELATION_TYPEDB

	// stores SPEC_RELATION_TYPEDB ID according to SPEC_RELATION_TYPE address
	Map_SPEC_RELATION_TYPEPtr_SPEC_RELATION_TYPEDBID map[*models.SPEC_RELATION_TYPE]uint

	// stores SPEC_RELATION_TYPE according to their gorm ID
	Map_SPEC_RELATION_TYPEDBID_SPEC_RELATION_TYPEPtr map[uint]*models.SPEC_RELATION_TYPE

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoSPEC_RELATION_TYPE *BackRepoSPEC_RELATION_TYPEStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoSPEC_RELATION_TYPE.stage
	return
}

func (backRepoSPEC_RELATION_TYPE *BackRepoSPEC_RELATION_TYPEStruct) GetDB() *gorm.DB {
	return backRepoSPEC_RELATION_TYPE.db
}

// GetSPEC_RELATION_TYPEDBFromSPEC_RELATION_TYPEPtr is a handy function to access the back repo instance from the stage instance
func (backRepoSPEC_RELATION_TYPE *BackRepoSPEC_RELATION_TYPEStruct) GetSPEC_RELATION_TYPEDBFromSPEC_RELATION_TYPEPtr(spec_relation_type *models.SPEC_RELATION_TYPE) (spec_relation_typeDB *SPEC_RELATION_TYPEDB) {
	id := backRepoSPEC_RELATION_TYPE.Map_SPEC_RELATION_TYPEPtr_SPEC_RELATION_TYPEDBID[spec_relation_type]
	spec_relation_typeDB = backRepoSPEC_RELATION_TYPE.Map_SPEC_RELATION_TYPEDBID_SPEC_RELATION_TYPEDB[id]
	return
}

// BackRepoSPEC_RELATION_TYPE.CommitPhaseOne commits all staged instances of SPEC_RELATION_TYPE to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoSPEC_RELATION_TYPE *BackRepoSPEC_RELATION_TYPEStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for spec_relation_type := range stage.SPEC_RELATION_TYPEs {
		backRepoSPEC_RELATION_TYPE.CommitPhaseOneInstance(spec_relation_type)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, spec_relation_type := range backRepoSPEC_RELATION_TYPE.Map_SPEC_RELATION_TYPEDBID_SPEC_RELATION_TYPEPtr {
		if _, ok := stage.SPEC_RELATION_TYPEs[spec_relation_type]; !ok {
			backRepoSPEC_RELATION_TYPE.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoSPEC_RELATION_TYPE.CommitDeleteInstance commits deletion of SPEC_RELATION_TYPE to the BackRepo
func (backRepoSPEC_RELATION_TYPE *BackRepoSPEC_RELATION_TYPEStruct) CommitDeleteInstance(id uint) (Error error) {

	spec_relation_type := backRepoSPEC_RELATION_TYPE.Map_SPEC_RELATION_TYPEDBID_SPEC_RELATION_TYPEPtr[id]

	// spec_relation_type is not staged anymore, remove spec_relation_typeDB
	spec_relation_typeDB := backRepoSPEC_RELATION_TYPE.Map_SPEC_RELATION_TYPEDBID_SPEC_RELATION_TYPEDB[id]
	query := backRepoSPEC_RELATION_TYPE.db.Unscoped().Delete(&spec_relation_typeDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	delete(backRepoSPEC_RELATION_TYPE.Map_SPEC_RELATION_TYPEPtr_SPEC_RELATION_TYPEDBID, spec_relation_type)
	delete(backRepoSPEC_RELATION_TYPE.Map_SPEC_RELATION_TYPEDBID_SPEC_RELATION_TYPEPtr, id)
	delete(backRepoSPEC_RELATION_TYPE.Map_SPEC_RELATION_TYPEDBID_SPEC_RELATION_TYPEDB, id)

	return
}

// BackRepoSPEC_RELATION_TYPE.CommitPhaseOneInstance commits spec_relation_type staged instances of SPEC_RELATION_TYPE to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoSPEC_RELATION_TYPE *BackRepoSPEC_RELATION_TYPEStruct) CommitPhaseOneInstance(spec_relation_type *models.SPEC_RELATION_TYPE) (Error error) {

	// check if the spec_relation_type is not commited yet
	if _, ok := backRepoSPEC_RELATION_TYPE.Map_SPEC_RELATION_TYPEPtr_SPEC_RELATION_TYPEDBID[spec_relation_type]; ok {
		return
	}

	// initiate spec_relation_type
	var spec_relation_typeDB SPEC_RELATION_TYPEDB
	spec_relation_typeDB.CopyBasicFieldsFromSPEC_RELATION_TYPE(spec_relation_type)

	query := backRepoSPEC_RELATION_TYPE.db.Create(&spec_relation_typeDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	backRepoSPEC_RELATION_TYPE.Map_SPEC_RELATION_TYPEPtr_SPEC_RELATION_TYPEDBID[spec_relation_type] = spec_relation_typeDB.ID
	backRepoSPEC_RELATION_TYPE.Map_SPEC_RELATION_TYPEDBID_SPEC_RELATION_TYPEPtr[spec_relation_typeDB.ID] = spec_relation_type
	backRepoSPEC_RELATION_TYPE.Map_SPEC_RELATION_TYPEDBID_SPEC_RELATION_TYPEDB[spec_relation_typeDB.ID] = &spec_relation_typeDB

	return
}

// BackRepoSPEC_RELATION_TYPE.CommitPhaseTwo commits all staged instances of SPEC_RELATION_TYPE to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSPEC_RELATION_TYPE *BackRepoSPEC_RELATION_TYPEStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, spec_relation_type := range backRepoSPEC_RELATION_TYPE.Map_SPEC_RELATION_TYPEDBID_SPEC_RELATION_TYPEPtr {
		backRepoSPEC_RELATION_TYPE.CommitPhaseTwoInstance(backRepo, idx, spec_relation_type)
	}

	return
}

// BackRepoSPEC_RELATION_TYPE.CommitPhaseTwoInstance commits {{structname }} of models.SPEC_RELATION_TYPE to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSPEC_RELATION_TYPE *BackRepoSPEC_RELATION_TYPEStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, spec_relation_type *models.SPEC_RELATION_TYPE) (Error error) {

	// fetch matching spec_relation_typeDB
	if spec_relation_typeDB, ok := backRepoSPEC_RELATION_TYPE.Map_SPEC_RELATION_TYPEDBID_SPEC_RELATION_TYPEDB[idx]; ok {

		spec_relation_typeDB.CopyBasicFieldsFromSPEC_RELATION_TYPE(spec_relation_type)

		// insertion point for translating pointers encodings into actual pointers
		query := backRepoSPEC_RELATION_TYPE.db.Save(&spec_relation_typeDB)
		if query.Error != nil {
			log.Fatalln(query.Error)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown SPEC_RELATION_TYPE intance %s", spec_relation_type.Name))
		return err
	}

	return
}

// BackRepoSPEC_RELATION_TYPE.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoSPEC_RELATION_TYPE *BackRepoSPEC_RELATION_TYPEStruct) CheckoutPhaseOne() (Error error) {

	spec_relation_typeDBArray := make([]SPEC_RELATION_TYPEDB, 0)
	query := backRepoSPEC_RELATION_TYPE.db.Find(&spec_relation_typeDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	spec_relation_typeInstancesToBeRemovedFromTheStage := make(map[*models.SPEC_RELATION_TYPE]any)
	for key, value := range backRepoSPEC_RELATION_TYPE.stage.SPEC_RELATION_TYPEs {
		spec_relation_typeInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, spec_relation_typeDB := range spec_relation_typeDBArray {
		backRepoSPEC_RELATION_TYPE.CheckoutPhaseOneInstance(&spec_relation_typeDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		spec_relation_type, ok := backRepoSPEC_RELATION_TYPE.Map_SPEC_RELATION_TYPEDBID_SPEC_RELATION_TYPEPtr[spec_relation_typeDB.ID]
		if ok {
			delete(spec_relation_typeInstancesToBeRemovedFromTheStage, spec_relation_type)
		}
	}

	// remove from stage and back repo's 3 maps all spec_relation_types that are not in the checkout
	for spec_relation_type := range spec_relation_typeInstancesToBeRemovedFromTheStage {
		spec_relation_type.Unstage(backRepoSPEC_RELATION_TYPE.GetStage())

		// remove instance from the back repo 3 maps
		spec_relation_typeID := backRepoSPEC_RELATION_TYPE.Map_SPEC_RELATION_TYPEPtr_SPEC_RELATION_TYPEDBID[spec_relation_type]
		delete(backRepoSPEC_RELATION_TYPE.Map_SPEC_RELATION_TYPEPtr_SPEC_RELATION_TYPEDBID, spec_relation_type)
		delete(backRepoSPEC_RELATION_TYPE.Map_SPEC_RELATION_TYPEDBID_SPEC_RELATION_TYPEDB, spec_relation_typeID)
		delete(backRepoSPEC_RELATION_TYPE.Map_SPEC_RELATION_TYPEDBID_SPEC_RELATION_TYPEPtr, spec_relation_typeID)
	}

	return
}

// CheckoutPhaseOneInstance takes a spec_relation_typeDB that has been found in the DB, updates the backRepo and stages the
// models version of the spec_relation_typeDB
func (backRepoSPEC_RELATION_TYPE *BackRepoSPEC_RELATION_TYPEStruct) CheckoutPhaseOneInstance(spec_relation_typeDB *SPEC_RELATION_TYPEDB) (Error error) {

	spec_relation_type, ok := backRepoSPEC_RELATION_TYPE.Map_SPEC_RELATION_TYPEDBID_SPEC_RELATION_TYPEPtr[spec_relation_typeDB.ID]
	if !ok {
		spec_relation_type = new(models.SPEC_RELATION_TYPE)

		backRepoSPEC_RELATION_TYPE.Map_SPEC_RELATION_TYPEDBID_SPEC_RELATION_TYPEPtr[spec_relation_typeDB.ID] = spec_relation_type
		backRepoSPEC_RELATION_TYPE.Map_SPEC_RELATION_TYPEPtr_SPEC_RELATION_TYPEDBID[spec_relation_type] = spec_relation_typeDB.ID

		// append model store with the new element
		spec_relation_type.Name = spec_relation_typeDB.Name_Data.String
		spec_relation_type.Stage(backRepoSPEC_RELATION_TYPE.GetStage())
	}
	spec_relation_typeDB.CopyBasicFieldsToSPEC_RELATION_TYPE(spec_relation_type)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	spec_relation_type.Stage(backRepoSPEC_RELATION_TYPE.GetStage())

	// preserve pointer to spec_relation_typeDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_SPEC_RELATION_TYPEDBID_SPEC_RELATION_TYPEDB)[spec_relation_typeDB hold variable pointers
	spec_relation_typeDB_Data := *spec_relation_typeDB
	preservedPtrToSPEC_RELATION_TYPE := &spec_relation_typeDB_Data
	backRepoSPEC_RELATION_TYPE.Map_SPEC_RELATION_TYPEDBID_SPEC_RELATION_TYPEDB[spec_relation_typeDB.ID] = preservedPtrToSPEC_RELATION_TYPE

	return
}

// BackRepoSPEC_RELATION_TYPE.CheckoutPhaseTwo Checkouts all staged instances of SPEC_RELATION_TYPE to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSPEC_RELATION_TYPE *BackRepoSPEC_RELATION_TYPEStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, spec_relation_typeDB := range backRepoSPEC_RELATION_TYPE.Map_SPEC_RELATION_TYPEDBID_SPEC_RELATION_TYPEDB {
		backRepoSPEC_RELATION_TYPE.CheckoutPhaseTwoInstance(backRepo, spec_relation_typeDB)
	}
	return
}

// BackRepoSPEC_RELATION_TYPE.CheckoutPhaseTwoInstance Checkouts staged instances of SPEC_RELATION_TYPE to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSPEC_RELATION_TYPE *BackRepoSPEC_RELATION_TYPEStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, spec_relation_typeDB *SPEC_RELATION_TYPEDB) (Error error) {

	spec_relation_type := backRepoSPEC_RELATION_TYPE.Map_SPEC_RELATION_TYPEDBID_SPEC_RELATION_TYPEPtr[spec_relation_typeDB.ID]

	spec_relation_typeDB.DecodePointers(backRepo, spec_relation_type)

	return
}

func (spec_relation_typeDB *SPEC_RELATION_TYPEDB) DecodePointers(backRepo *BackRepoStruct, spec_relation_type *models.SPEC_RELATION_TYPE) {

	// insertion point for checkout of pointer encoding
	return
}

// CommitSPEC_RELATION_TYPE allows commit of a single spec_relation_type (if already staged)
func (backRepo *BackRepoStruct) CommitSPEC_RELATION_TYPE(spec_relation_type *models.SPEC_RELATION_TYPE) {
	backRepo.BackRepoSPEC_RELATION_TYPE.CommitPhaseOneInstance(spec_relation_type)
	if id, ok := backRepo.BackRepoSPEC_RELATION_TYPE.Map_SPEC_RELATION_TYPEPtr_SPEC_RELATION_TYPEDBID[spec_relation_type]; ok {
		backRepo.BackRepoSPEC_RELATION_TYPE.CommitPhaseTwoInstance(backRepo, id, spec_relation_type)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitSPEC_RELATION_TYPE allows checkout of a single spec_relation_type (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutSPEC_RELATION_TYPE(spec_relation_type *models.SPEC_RELATION_TYPE) {
	// check if the spec_relation_type is staged
	if _, ok := backRepo.BackRepoSPEC_RELATION_TYPE.Map_SPEC_RELATION_TYPEPtr_SPEC_RELATION_TYPEDBID[spec_relation_type]; ok {

		if id, ok := backRepo.BackRepoSPEC_RELATION_TYPE.Map_SPEC_RELATION_TYPEPtr_SPEC_RELATION_TYPEDBID[spec_relation_type]; ok {
			var spec_relation_typeDB SPEC_RELATION_TYPEDB
			spec_relation_typeDB.ID = id

			if err := backRepo.BackRepoSPEC_RELATION_TYPE.db.First(&spec_relation_typeDB, id).Error; err != nil {
				log.Fatalln("CheckoutSPEC_RELATION_TYPE : Problem with getting object with id:", id)
			}
			backRepo.BackRepoSPEC_RELATION_TYPE.CheckoutPhaseOneInstance(&spec_relation_typeDB)
			backRepo.BackRepoSPEC_RELATION_TYPE.CheckoutPhaseTwoInstance(backRepo, &spec_relation_typeDB)
		}
	}
}

// CopyBasicFieldsFromSPEC_RELATION_TYPE
func (spec_relation_typeDB *SPEC_RELATION_TYPEDB) CopyBasicFieldsFromSPEC_RELATION_TYPE(spec_relation_type *models.SPEC_RELATION_TYPE) {
	// insertion point for fields commit

	spec_relation_typeDB.Name_Data.String = spec_relation_type.Name
	spec_relation_typeDB.Name_Data.Valid = true

	spec_relation_typeDB.DESC_Data.String = spec_relation_type.DESC
	spec_relation_typeDB.DESC_Data.Valid = true

	spec_relation_typeDB.IDENTIFIER_Data.String = spec_relation_type.IDENTIFIER
	spec_relation_typeDB.IDENTIFIER_Data.Valid = true

	spec_relation_typeDB.LAST_CHANGE_Data.String = spec_relation_type.LAST_CHANGE
	spec_relation_typeDB.LAST_CHANGE_Data.Valid = true

	spec_relation_typeDB.LONG_NAME_Data.String = spec_relation_type.LONG_NAME
	spec_relation_typeDB.LONG_NAME_Data.Valid = true
}

// CopyBasicFieldsFromSPEC_RELATION_TYPE_WOP
func (spec_relation_typeDB *SPEC_RELATION_TYPEDB) CopyBasicFieldsFromSPEC_RELATION_TYPE_WOP(spec_relation_type *models.SPEC_RELATION_TYPE_WOP) {
	// insertion point for fields commit

	spec_relation_typeDB.Name_Data.String = spec_relation_type.Name
	spec_relation_typeDB.Name_Data.Valid = true

	spec_relation_typeDB.DESC_Data.String = spec_relation_type.DESC
	spec_relation_typeDB.DESC_Data.Valid = true

	spec_relation_typeDB.IDENTIFIER_Data.String = spec_relation_type.IDENTIFIER
	spec_relation_typeDB.IDENTIFIER_Data.Valid = true

	spec_relation_typeDB.LAST_CHANGE_Data.String = spec_relation_type.LAST_CHANGE
	spec_relation_typeDB.LAST_CHANGE_Data.Valid = true

	spec_relation_typeDB.LONG_NAME_Data.String = spec_relation_type.LONG_NAME
	spec_relation_typeDB.LONG_NAME_Data.Valid = true
}

// CopyBasicFieldsFromSPEC_RELATION_TYPEWOP
func (spec_relation_typeDB *SPEC_RELATION_TYPEDB) CopyBasicFieldsFromSPEC_RELATION_TYPEWOP(spec_relation_type *SPEC_RELATION_TYPEWOP) {
	// insertion point for fields commit

	spec_relation_typeDB.Name_Data.String = spec_relation_type.Name
	spec_relation_typeDB.Name_Data.Valid = true

	spec_relation_typeDB.DESC_Data.String = spec_relation_type.DESC
	spec_relation_typeDB.DESC_Data.Valid = true

	spec_relation_typeDB.IDENTIFIER_Data.String = spec_relation_type.IDENTIFIER
	spec_relation_typeDB.IDENTIFIER_Data.Valid = true

	spec_relation_typeDB.LAST_CHANGE_Data.String = spec_relation_type.LAST_CHANGE
	spec_relation_typeDB.LAST_CHANGE_Data.Valid = true

	spec_relation_typeDB.LONG_NAME_Data.String = spec_relation_type.LONG_NAME
	spec_relation_typeDB.LONG_NAME_Data.Valid = true
}

// CopyBasicFieldsToSPEC_RELATION_TYPE
func (spec_relation_typeDB *SPEC_RELATION_TYPEDB) CopyBasicFieldsToSPEC_RELATION_TYPE(spec_relation_type *models.SPEC_RELATION_TYPE) {
	// insertion point for checkout of basic fields (back repo to stage)
	spec_relation_type.Name = spec_relation_typeDB.Name_Data.String
	spec_relation_type.DESC = spec_relation_typeDB.DESC_Data.String
	spec_relation_type.IDENTIFIER = spec_relation_typeDB.IDENTIFIER_Data.String
	spec_relation_type.LAST_CHANGE = spec_relation_typeDB.LAST_CHANGE_Data.String
	spec_relation_type.LONG_NAME = spec_relation_typeDB.LONG_NAME_Data.String
}

// CopyBasicFieldsToSPEC_RELATION_TYPE_WOP
func (spec_relation_typeDB *SPEC_RELATION_TYPEDB) CopyBasicFieldsToSPEC_RELATION_TYPE_WOP(spec_relation_type *models.SPEC_RELATION_TYPE_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	spec_relation_type.Name = spec_relation_typeDB.Name_Data.String
	spec_relation_type.DESC = spec_relation_typeDB.DESC_Data.String
	spec_relation_type.IDENTIFIER = spec_relation_typeDB.IDENTIFIER_Data.String
	spec_relation_type.LAST_CHANGE = spec_relation_typeDB.LAST_CHANGE_Data.String
	spec_relation_type.LONG_NAME = spec_relation_typeDB.LONG_NAME_Data.String
}

// CopyBasicFieldsToSPEC_RELATION_TYPEWOP
func (spec_relation_typeDB *SPEC_RELATION_TYPEDB) CopyBasicFieldsToSPEC_RELATION_TYPEWOP(spec_relation_type *SPEC_RELATION_TYPEWOP) {
	spec_relation_type.ID = int(spec_relation_typeDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	spec_relation_type.Name = spec_relation_typeDB.Name_Data.String
	spec_relation_type.DESC = spec_relation_typeDB.DESC_Data.String
	spec_relation_type.IDENTIFIER = spec_relation_typeDB.IDENTIFIER_Data.String
	spec_relation_type.LAST_CHANGE = spec_relation_typeDB.LAST_CHANGE_Data.String
	spec_relation_type.LONG_NAME = spec_relation_typeDB.LONG_NAME_Data.String
}

// Backup generates a json file from a slice of all SPEC_RELATION_TYPEDB instances in the backrepo
func (backRepoSPEC_RELATION_TYPE *BackRepoSPEC_RELATION_TYPEStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "SPEC_RELATION_TYPEDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*SPEC_RELATION_TYPEDB, 0)
	for _, spec_relation_typeDB := range backRepoSPEC_RELATION_TYPE.Map_SPEC_RELATION_TYPEDBID_SPEC_RELATION_TYPEDB {
		forBackup = append(forBackup, spec_relation_typeDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json SPEC_RELATION_TYPE ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json SPEC_RELATION_TYPE file", err.Error())
	}
}

// Backup generates a json file from a slice of all SPEC_RELATION_TYPEDB instances in the backrepo
func (backRepoSPEC_RELATION_TYPE *BackRepoSPEC_RELATION_TYPEStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*SPEC_RELATION_TYPEDB, 0)
	for _, spec_relation_typeDB := range backRepoSPEC_RELATION_TYPE.Map_SPEC_RELATION_TYPEDBID_SPEC_RELATION_TYPEDB {
		forBackup = append(forBackup, spec_relation_typeDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("SPEC_RELATION_TYPE")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&SPEC_RELATION_TYPE_Fields, -1)
	for _, spec_relation_typeDB := range forBackup {

		var spec_relation_typeWOP SPEC_RELATION_TYPEWOP
		spec_relation_typeDB.CopyBasicFieldsToSPEC_RELATION_TYPEWOP(&spec_relation_typeWOP)

		row := sh.AddRow()
		row.WriteStruct(&spec_relation_typeWOP, -1)
	}
}

// RestoreXL from the "SPEC_RELATION_TYPE" sheet all SPEC_RELATION_TYPEDB instances
func (backRepoSPEC_RELATION_TYPE *BackRepoSPEC_RELATION_TYPEStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoSPEC_RELATION_TYPEid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["SPEC_RELATION_TYPE"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoSPEC_RELATION_TYPE.rowVisitorSPEC_RELATION_TYPE)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoSPEC_RELATION_TYPE *BackRepoSPEC_RELATION_TYPEStruct) rowVisitorSPEC_RELATION_TYPE(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var spec_relation_typeWOP SPEC_RELATION_TYPEWOP
		row.ReadStruct(&spec_relation_typeWOP)

		// add the unmarshalled struct to the stage
		spec_relation_typeDB := new(SPEC_RELATION_TYPEDB)
		spec_relation_typeDB.CopyBasicFieldsFromSPEC_RELATION_TYPEWOP(&spec_relation_typeWOP)

		spec_relation_typeDB_ID_atBackupTime := spec_relation_typeDB.ID
		spec_relation_typeDB.ID = 0
		query := backRepoSPEC_RELATION_TYPE.db.Create(spec_relation_typeDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoSPEC_RELATION_TYPE.Map_SPEC_RELATION_TYPEDBID_SPEC_RELATION_TYPEDB[spec_relation_typeDB.ID] = spec_relation_typeDB
		BackRepoSPEC_RELATION_TYPEid_atBckpTime_newID[spec_relation_typeDB_ID_atBackupTime] = spec_relation_typeDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "SPEC_RELATION_TYPEDB.json" in dirPath that stores an array
// of SPEC_RELATION_TYPEDB and stores it in the database
// the map BackRepoSPEC_RELATION_TYPEid_atBckpTime_newID is updated accordingly
func (backRepoSPEC_RELATION_TYPE *BackRepoSPEC_RELATION_TYPEStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoSPEC_RELATION_TYPEid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "SPEC_RELATION_TYPEDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json SPEC_RELATION_TYPE file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*SPEC_RELATION_TYPEDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_SPEC_RELATION_TYPEDBID_SPEC_RELATION_TYPEDB
	for _, spec_relation_typeDB := range forRestore {

		spec_relation_typeDB_ID_atBackupTime := spec_relation_typeDB.ID
		spec_relation_typeDB.ID = 0
		query := backRepoSPEC_RELATION_TYPE.db.Create(spec_relation_typeDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoSPEC_RELATION_TYPE.Map_SPEC_RELATION_TYPEDBID_SPEC_RELATION_TYPEDB[spec_relation_typeDB.ID] = spec_relation_typeDB
		BackRepoSPEC_RELATION_TYPEid_atBckpTime_newID[spec_relation_typeDB_ID_atBackupTime] = spec_relation_typeDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json SPEC_RELATION_TYPE file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<SPEC_RELATION_TYPE>id_atBckpTime_newID
// to compute new index
func (backRepoSPEC_RELATION_TYPE *BackRepoSPEC_RELATION_TYPEStruct) RestorePhaseTwo() {

	for _, spec_relation_typeDB := range backRepoSPEC_RELATION_TYPE.Map_SPEC_RELATION_TYPEDBID_SPEC_RELATION_TYPEDB {

		// next line of code is to avert unused variable compilation error
		_ = spec_relation_typeDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		query := backRepoSPEC_RELATION_TYPE.db.Model(spec_relation_typeDB).Updates(*spec_relation_typeDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
	}

}

// BackRepoSPEC_RELATION_TYPE.ResetReversePointers commits all staged instances of SPEC_RELATION_TYPE to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSPEC_RELATION_TYPE *BackRepoSPEC_RELATION_TYPEStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, spec_relation_type := range backRepoSPEC_RELATION_TYPE.Map_SPEC_RELATION_TYPEDBID_SPEC_RELATION_TYPEPtr {
		backRepoSPEC_RELATION_TYPE.ResetReversePointersInstance(backRepo, idx, spec_relation_type)
	}

	return
}

func (backRepoSPEC_RELATION_TYPE *BackRepoSPEC_RELATION_TYPEStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, spec_relation_type *models.SPEC_RELATION_TYPE) (Error error) {

	// fetch matching spec_relation_typeDB
	if spec_relation_typeDB, ok := backRepoSPEC_RELATION_TYPE.Map_SPEC_RELATION_TYPEDBID_SPEC_RELATION_TYPEDB[idx]; ok {
		_ = spec_relation_typeDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoSPEC_RELATION_TYPEid_atBckpTime_newID map[uint]uint