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
var dummy_A_ATTRIBUTE_VALUE_ENUMERATION_sql sql.NullBool
var dummy_A_ATTRIBUTE_VALUE_ENUMERATION_time time.Duration
var dummy_A_ATTRIBUTE_VALUE_ENUMERATION_sort sort.Float64Slice

// A_ATTRIBUTE_VALUE_ENUMERATIONAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model a_attribute_value_enumerationAPI
type A_ATTRIBUTE_VALUE_ENUMERATIONAPI struct {
	gorm.Model

	models.A_ATTRIBUTE_VALUE_ENUMERATION_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	A_ATTRIBUTE_VALUE_ENUMERATIONPointersEncoding A_ATTRIBUTE_VALUE_ENUMERATIONPointersEncoding
}

// A_ATTRIBUTE_VALUE_ENUMERATIONPointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type A_ATTRIBUTE_VALUE_ENUMERATIONPointersEncoding struct {
	// insertion for pointer fields encoding declaration

	// field ATTRIBUTE_VALUE_ENUMERATION is a slice of pointers to another Struct (optional or 0..1)
	ATTRIBUTE_VALUE_ENUMERATION IntSlice `gorm:"type:TEXT"`
}

// A_ATTRIBUTE_VALUE_ENUMERATIONDB describes a a_attribute_value_enumeration in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model a_attribute_value_enumerationDB
type A_ATTRIBUTE_VALUE_ENUMERATIONDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field a_attribute_value_enumerationDB.Name
	Name_Data sql.NullString
	
	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	A_ATTRIBUTE_VALUE_ENUMERATIONPointersEncoding
}

// A_ATTRIBUTE_VALUE_ENUMERATIONDBs arrays a_attribute_value_enumerationDBs
// swagger:response a_attribute_value_enumerationDBsResponse
type A_ATTRIBUTE_VALUE_ENUMERATIONDBs []A_ATTRIBUTE_VALUE_ENUMERATIONDB

// A_ATTRIBUTE_VALUE_ENUMERATIONDBResponse provides response
// swagger:response a_attribute_value_enumerationDBResponse
type A_ATTRIBUTE_VALUE_ENUMERATIONDBResponse struct {
	A_ATTRIBUTE_VALUE_ENUMERATIONDB
}

// A_ATTRIBUTE_VALUE_ENUMERATIONWOP is a A_ATTRIBUTE_VALUE_ENUMERATION without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type A_ATTRIBUTE_VALUE_ENUMERATIONWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`
	// insertion for WOP pointer fields
}

var A_ATTRIBUTE_VALUE_ENUMERATION_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
}

type BackRepoA_ATTRIBUTE_VALUE_ENUMERATIONStruct struct {
	// stores A_ATTRIBUTE_VALUE_ENUMERATIONDB according to their gorm ID
	Map_A_ATTRIBUTE_VALUE_ENUMERATIONDBID_A_ATTRIBUTE_VALUE_ENUMERATIONDB map[uint]*A_ATTRIBUTE_VALUE_ENUMERATIONDB

	// stores A_ATTRIBUTE_VALUE_ENUMERATIONDB ID according to A_ATTRIBUTE_VALUE_ENUMERATION address
	Map_A_ATTRIBUTE_VALUE_ENUMERATIONPtr_A_ATTRIBUTE_VALUE_ENUMERATIONDBID map[*models.A_ATTRIBUTE_VALUE_ENUMERATION]uint

	// stores A_ATTRIBUTE_VALUE_ENUMERATION according to their gorm ID
	Map_A_ATTRIBUTE_VALUE_ENUMERATIONDBID_A_ATTRIBUTE_VALUE_ENUMERATIONPtr map[uint]*models.A_ATTRIBUTE_VALUE_ENUMERATION

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoA_ATTRIBUTE_VALUE_ENUMERATION *BackRepoA_ATTRIBUTE_VALUE_ENUMERATIONStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoA_ATTRIBUTE_VALUE_ENUMERATION.stage
	return
}

func (backRepoA_ATTRIBUTE_VALUE_ENUMERATION *BackRepoA_ATTRIBUTE_VALUE_ENUMERATIONStruct) GetDB() *gorm.DB {
	return backRepoA_ATTRIBUTE_VALUE_ENUMERATION.db
}

// GetA_ATTRIBUTE_VALUE_ENUMERATIONDBFromA_ATTRIBUTE_VALUE_ENUMERATIONPtr is a handy function to access the back repo instance from the stage instance
func (backRepoA_ATTRIBUTE_VALUE_ENUMERATION *BackRepoA_ATTRIBUTE_VALUE_ENUMERATIONStruct) GetA_ATTRIBUTE_VALUE_ENUMERATIONDBFromA_ATTRIBUTE_VALUE_ENUMERATIONPtr(a_attribute_value_enumeration *models.A_ATTRIBUTE_VALUE_ENUMERATION) (a_attribute_value_enumerationDB *A_ATTRIBUTE_VALUE_ENUMERATIONDB) {
	id := backRepoA_ATTRIBUTE_VALUE_ENUMERATION.Map_A_ATTRIBUTE_VALUE_ENUMERATIONPtr_A_ATTRIBUTE_VALUE_ENUMERATIONDBID[a_attribute_value_enumeration]
	a_attribute_value_enumerationDB = backRepoA_ATTRIBUTE_VALUE_ENUMERATION.Map_A_ATTRIBUTE_VALUE_ENUMERATIONDBID_A_ATTRIBUTE_VALUE_ENUMERATIONDB[id]
	return
}

// BackRepoA_ATTRIBUTE_VALUE_ENUMERATION.CommitPhaseOne commits all staged instances of A_ATTRIBUTE_VALUE_ENUMERATION to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoA_ATTRIBUTE_VALUE_ENUMERATION *BackRepoA_ATTRIBUTE_VALUE_ENUMERATIONStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for a_attribute_value_enumeration := range stage.A_ATTRIBUTE_VALUE_ENUMERATIONs {
		backRepoA_ATTRIBUTE_VALUE_ENUMERATION.CommitPhaseOneInstance(a_attribute_value_enumeration)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, a_attribute_value_enumeration := range backRepoA_ATTRIBUTE_VALUE_ENUMERATION.Map_A_ATTRIBUTE_VALUE_ENUMERATIONDBID_A_ATTRIBUTE_VALUE_ENUMERATIONPtr {
		if _, ok := stage.A_ATTRIBUTE_VALUE_ENUMERATIONs[a_attribute_value_enumeration]; !ok {
			backRepoA_ATTRIBUTE_VALUE_ENUMERATION.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoA_ATTRIBUTE_VALUE_ENUMERATION.CommitDeleteInstance commits deletion of A_ATTRIBUTE_VALUE_ENUMERATION to the BackRepo
func (backRepoA_ATTRIBUTE_VALUE_ENUMERATION *BackRepoA_ATTRIBUTE_VALUE_ENUMERATIONStruct) CommitDeleteInstance(id uint) (Error error) {

	a_attribute_value_enumeration := backRepoA_ATTRIBUTE_VALUE_ENUMERATION.Map_A_ATTRIBUTE_VALUE_ENUMERATIONDBID_A_ATTRIBUTE_VALUE_ENUMERATIONPtr[id]

	// a_attribute_value_enumeration is not staged anymore, remove a_attribute_value_enumerationDB
	a_attribute_value_enumerationDB := backRepoA_ATTRIBUTE_VALUE_ENUMERATION.Map_A_ATTRIBUTE_VALUE_ENUMERATIONDBID_A_ATTRIBUTE_VALUE_ENUMERATIONDB[id]
	query := backRepoA_ATTRIBUTE_VALUE_ENUMERATION.db.Unscoped().Delete(&a_attribute_value_enumerationDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	delete(backRepoA_ATTRIBUTE_VALUE_ENUMERATION.Map_A_ATTRIBUTE_VALUE_ENUMERATIONPtr_A_ATTRIBUTE_VALUE_ENUMERATIONDBID, a_attribute_value_enumeration)
	delete(backRepoA_ATTRIBUTE_VALUE_ENUMERATION.Map_A_ATTRIBUTE_VALUE_ENUMERATIONDBID_A_ATTRIBUTE_VALUE_ENUMERATIONPtr, id)
	delete(backRepoA_ATTRIBUTE_VALUE_ENUMERATION.Map_A_ATTRIBUTE_VALUE_ENUMERATIONDBID_A_ATTRIBUTE_VALUE_ENUMERATIONDB, id)

	return
}

// BackRepoA_ATTRIBUTE_VALUE_ENUMERATION.CommitPhaseOneInstance commits a_attribute_value_enumeration staged instances of A_ATTRIBUTE_VALUE_ENUMERATION to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoA_ATTRIBUTE_VALUE_ENUMERATION *BackRepoA_ATTRIBUTE_VALUE_ENUMERATIONStruct) CommitPhaseOneInstance(a_attribute_value_enumeration *models.A_ATTRIBUTE_VALUE_ENUMERATION) (Error error) {

	// check if the a_attribute_value_enumeration is not commited yet
	if _, ok := backRepoA_ATTRIBUTE_VALUE_ENUMERATION.Map_A_ATTRIBUTE_VALUE_ENUMERATIONPtr_A_ATTRIBUTE_VALUE_ENUMERATIONDBID[a_attribute_value_enumeration]; ok {
		return
	}

	// initiate a_attribute_value_enumeration
	var a_attribute_value_enumerationDB A_ATTRIBUTE_VALUE_ENUMERATIONDB
	a_attribute_value_enumerationDB.CopyBasicFieldsFromA_ATTRIBUTE_VALUE_ENUMERATION(a_attribute_value_enumeration)

	query := backRepoA_ATTRIBUTE_VALUE_ENUMERATION.db.Create(&a_attribute_value_enumerationDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	backRepoA_ATTRIBUTE_VALUE_ENUMERATION.Map_A_ATTRIBUTE_VALUE_ENUMERATIONPtr_A_ATTRIBUTE_VALUE_ENUMERATIONDBID[a_attribute_value_enumeration] = a_attribute_value_enumerationDB.ID
	backRepoA_ATTRIBUTE_VALUE_ENUMERATION.Map_A_ATTRIBUTE_VALUE_ENUMERATIONDBID_A_ATTRIBUTE_VALUE_ENUMERATIONPtr[a_attribute_value_enumerationDB.ID] = a_attribute_value_enumeration
	backRepoA_ATTRIBUTE_VALUE_ENUMERATION.Map_A_ATTRIBUTE_VALUE_ENUMERATIONDBID_A_ATTRIBUTE_VALUE_ENUMERATIONDB[a_attribute_value_enumerationDB.ID] = &a_attribute_value_enumerationDB

	return
}

// BackRepoA_ATTRIBUTE_VALUE_ENUMERATION.CommitPhaseTwo commits all staged instances of A_ATTRIBUTE_VALUE_ENUMERATION to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoA_ATTRIBUTE_VALUE_ENUMERATION *BackRepoA_ATTRIBUTE_VALUE_ENUMERATIONStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, a_attribute_value_enumeration := range backRepoA_ATTRIBUTE_VALUE_ENUMERATION.Map_A_ATTRIBUTE_VALUE_ENUMERATIONDBID_A_ATTRIBUTE_VALUE_ENUMERATIONPtr {
		backRepoA_ATTRIBUTE_VALUE_ENUMERATION.CommitPhaseTwoInstance(backRepo, idx, a_attribute_value_enumeration)
	}

	return
}

// BackRepoA_ATTRIBUTE_VALUE_ENUMERATION.CommitPhaseTwoInstance commits {{structname }} of models.A_ATTRIBUTE_VALUE_ENUMERATION to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoA_ATTRIBUTE_VALUE_ENUMERATION *BackRepoA_ATTRIBUTE_VALUE_ENUMERATIONStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, a_attribute_value_enumeration *models.A_ATTRIBUTE_VALUE_ENUMERATION) (Error error) {

	// fetch matching a_attribute_value_enumerationDB
	if a_attribute_value_enumerationDB, ok := backRepoA_ATTRIBUTE_VALUE_ENUMERATION.Map_A_ATTRIBUTE_VALUE_ENUMERATIONDBID_A_ATTRIBUTE_VALUE_ENUMERATIONDB[idx]; ok {

		a_attribute_value_enumerationDB.CopyBasicFieldsFromA_ATTRIBUTE_VALUE_ENUMERATION(a_attribute_value_enumeration)

		// insertion point for translating pointers encodings into actual pointers
		// 1. reset
		a_attribute_value_enumerationDB.A_ATTRIBUTE_VALUE_ENUMERATIONPointersEncoding.ATTRIBUTE_VALUE_ENUMERATION = make([]int, 0)
		// 2. encode
		for _, attribute_value_enumerationAssocEnd := range a_attribute_value_enumeration.ATTRIBUTE_VALUE_ENUMERATION {
			attribute_value_enumerationAssocEnd_DB :=
				backRepo.BackRepoATTRIBUTE_VALUE_ENUMERATION.GetATTRIBUTE_VALUE_ENUMERATIONDBFromATTRIBUTE_VALUE_ENUMERATIONPtr(attribute_value_enumerationAssocEnd)
			
			// the stage might be inconsistant, meaning that the attribute_value_enumerationAssocEnd_DB might
			// be missing from the stage. In this case, the commit operation is robust
			// An alternative would be to crash here to reveal the missing element.
			if attribute_value_enumerationAssocEnd_DB == nil {
				continue
			}
			
			a_attribute_value_enumerationDB.A_ATTRIBUTE_VALUE_ENUMERATIONPointersEncoding.ATTRIBUTE_VALUE_ENUMERATION =
				append(a_attribute_value_enumerationDB.A_ATTRIBUTE_VALUE_ENUMERATIONPointersEncoding.ATTRIBUTE_VALUE_ENUMERATION, int(attribute_value_enumerationAssocEnd_DB.ID))
		}

		query := backRepoA_ATTRIBUTE_VALUE_ENUMERATION.db.Save(&a_attribute_value_enumerationDB)
		if query.Error != nil {
			log.Fatalln(query.Error)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown A_ATTRIBUTE_VALUE_ENUMERATION intance %s", a_attribute_value_enumeration.Name))
		return err
	}

	return
}

// BackRepoA_ATTRIBUTE_VALUE_ENUMERATION.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoA_ATTRIBUTE_VALUE_ENUMERATION *BackRepoA_ATTRIBUTE_VALUE_ENUMERATIONStruct) CheckoutPhaseOne() (Error error) {

	a_attribute_value_enumerationDBArray := make([]A_ATTRIBUTE_VALUE_ENUMERATIONDB, 0)
	query := backRepoA_ATTRIBUTE_VALUE_ENUMERATION.db.Find(&a_attribute_value_enumerationDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	a_attribute_value_enumerationInstancesToBeRemovedFromTheStage := make(map[*models.A_ATTRIBUTE_VALUE_ENUMERATION]any)
	for key, value := range backRepoA_ATTRIBUTE_VALUE_ENUMERATION.stage.A_ATTRIBUTE_VALUE_ENUMERATIONs {
		a_attribute_value_enumerationInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, a_attribute_value_enumerationDB := range a_attribute_value_enumerationDBArray {
		backRepoA_ATTRIBUTE_VALUE_ENUMERATION.CheckoutPhaseOneInstance(&a_attribute_value_enumerationDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		a_attribute_value_enumeration, ok := backRepoA_ATTRIBUTE_VALUE_ENUMERATION.Map_A_ATTRIBUTE_VALUE_ENUMERATIONDBID_A_ATTRIBUTE_VALUE_ENUMERATIONPtr[a_attribute_value_enumerationDB.ID]
		if ok {
			delete(a_attribute_value_enumerationInstancesToBeRemovedFromTheStage, a_attribute_value_enumeration)
		}
	}

	// remove from stage and back repo's 3 maps all a_attribute_value_enumerations that are not in the checkout
	for a_attribute_value_enumeration := range a_attribute_value_enumerationInstancesToBeRemovedFromTheStage {
		a_attribute_value_enumeration.Unstage(backRepoA_ATTRIBUTE_VALUE_ENUMERATION.GetStage())

		// remove instance from the back repo 3 maps
		a_attribute_value_enumerationID := backRepoA_ATTRIBUTE_VALUE_ENUMERATION.Map_A_ATTRIBUTE_VALUE_ENUMERATIONPtr_A_ATTRIBUTE_VALUE_ENUMERATIONDBID[a_attribute_value_enumeration]
		delete(backRepoA_ATTRIBUTE_VALUE_ENUMERATION.Map_A_ATTRIBUTE_VALUE_ENUMERATIONPtr_A_ATTRIBUTE_VALUE_ENUMERATIONDBID, a_attribute_value_enumeration)
		delete(backRepoA_ATTRIBUTE_VALUE_ENUMERATION.Map_A_ATTRIBUTE_VALUE_ENUMERATIONDBID_A_ATTRIBUTE_VALUE_ENUMERATIONDB, a_attribute_value_enumerationID)
		delete(backRepoA_ATTRIBUTE_VALUE_ENUMERATION.Map_A_ATTRIBUTE_VALUE_ENUMERATIONDBID_A_ATTRIBUTE_VALUE_ENUMERATIONPtr, a_attribute_value_enumerationID)
	}

	return
}

// CheckoutPhaseOneInstance takes a a_attribute_value_enumerationDB that has been found in the DB, updates the backRepo and stages the
// models version of the a_attribute_value_enumerationDB
func (backRepoA_ATTRIBUTE_VALUE_ENUMERATION *BackRepoA_ATTRIBUTE_VALUE_ENUMERATIONStruct) CheckoutPhaseOneInstance(a_attribute_value_enumerationDB *A_ATTRIBUTE_VALUE_ENUMERATIONDB) (Error error) {

	a_attribute_value_enumeration, ok := backRepoA_ATTRIBUTE_VALUE_ENUMERATION.Map_A_ATTRIBUTE_VALUE_ENUMERATIONDBID_A_ATTRIBUTE_VALUE_ENUMERATIONPtr[a_attribute_value_enumerationDB.ID]
	if !ok {
		a_attribute_value_enumeration = new(models.A_ATTRIBUTE_VALUE_ENUMERATION)

		backRepoA_ATTRIBUTE_VALUE_ENUMERATION.Map_A_ATTRIBUTE_VALUE_ENUMERATIONDBID_A_ATTRIBUTE_VALUE_ENUMERATIONPtr[a_attribute_value_enumerationDB.ID] = a_attribute_value_enumeration
		backRepoA_ATTRIBUTE_VALUE_ENUMERATION.Map_A_ATTRIBUTE_VALUE_ENUMERATIONPtr_A_ATTRIBUTE_VALUE_ENUMERATIONDBID[a_attribute_value_enumeration] = a_attribute_value_enumerationDB.ID

		// append model store with the new element
		a_attribute_value_enumeration.Name = a_attribute_value_enumerationDB.Name_Data.String
		a_attribute_value_enumeration.Stage(backRepoA_ATTRIBUTE_VALUE_ENUMERATION.GetStage())
	}
	a_attribute_value_enumerationDB.CopyBasicFieldsToA_ATTRIBUTE_VALUE_ENUMERATION(a_attribute_value_enumeration)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	a_attribute_value_enumeration.Stage(backRepoA_ATTRIBUTE_VALUE_ENUMERATION.GetStage())

	// preserve pointer to a_attribute_value_enumerationDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_A_ATTRIBUTE_VALUE_ENUMERATIONDBID_A_ATTRIBUTE_VALUE_ENUMERATIONDB)[a_attribute_value_enumerationDB hold variable pointers
	a_attribute_value_enumerationDB_Data := *a_attribute_value_enumerationDB
	preservedPtrToA_ATTRIBUTE_VALUE_ENUMERATION := &a_attribute_value_enumerationDB_Data
	backRepoA_ATTRIBUTE_VALUE_ENUMERATION.Map_A_ATTRIBUTE_VALUE_ENUMERATIONDBID_A_ATTRIBUTE_VALUE_ENUMERATIONDB[a_attribute_value_enumerationDB.ID] = preservedPtrToA_ATTRIBUTE_VALUE_ENUMERATION

	return
}

// BackRepoA_ATTRIBUTE_VALUE_ENUMERATION.CheckoutPhaseTwo Checkouts all staged instances of A_ATTRIBUTE_VALUE_ENUMERATION to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoA_ATTRIBUTE_VALUE_ENUMERATION *BackRepoA_ATTRIBUTE_VALUE_ENUMERATIONStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, a_attribute_value_enumerationDB := range backRepoA_ATTRIBUTE_VALUE_ENUMERATION.Map_A_ATTRIBUTE_VALUE_ENUMERATIONDBID_A_ATTRIBUTE_VALUE_ENUMERATIONDB {
		backRepoA_ATTRIBUTE_VALUE_ENUMERATION.CheckoutPhaseTwoInstance(backRepo, a_attribute_value_enumerationDB)
	}
	return
}

// BackRepoA_ATTRIBUTE_VALUE_ENUMERATION.CheckoutPhaseTwoInstance Checkouts staged instances of A_ATTRIBUTE_VALUE_ENUMERATION to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoA_ATTRIBUTE_VALUE_ENUMERATION *BackRepoA_ATTRIBUTE_VALUE_ENUMERATIONStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, a_attribute_value_enumerationDB *A_ATTRIBUTE_VALUE_ENUMERATIONDB) (Error error) {

	a_attribute_value_enumeration := backRepoA_ATTRIBUTE_VALUE_ENUMERATION.Map_A_ATTRIBUTE_VALUE_ENUMERATIONDBID_A_ATTRIBUTE_VALUE_ENUMERATIONPtr[a_attribute_value_enumerationDB.ID]

	a_attribute_value_enumerationDB.DecodePointers(backRepo, a_attribute_value_enumeration)

	return
}

func (a_attribute_value_enumerationDB *A_ATTRIBUTE_VALUE_ENUMERATIONDB) DecodePointers(backRepo *BackRepoStruct, a_attribute_value_enumeration *models.A_ATTRIBUTE_VALUE_ENUMERATION) {

	// insertion point for checkout of pointer encoding
	// This loop redeem a_attribute_value_enumeration.ATTRIBUTE_VALUE_ENUMERATION in the stage from the encode in the back repo
	// It parses all ATTRIBUTE_VALUE_ENUMERATIONDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	a_attribute_value_enumeration.ATTRIBUTE_VALUE_ENUMERATION = a_attribute_value_enumeration.ATTRIBUTE_VALUE_ENUMERATION[:0]
	for _, _ATTRIBUTE_VALUE_ENUMERATIONid := range a_attribute_value_enumerationDB.A_ATTRIBUTE_VALUE_ENUMERATIONPointersEncoding.ATTRIBUTE_VALUE_ENUMERATION {
		a_attribute_value_enumeration.ATTRIBUTE_VALUE_ENUMERATION = append(a_attribute_value_enumeration.ATTRIBUTE_VALUE_ENUMERATION, backRepo.BackRepoATTRIBUTE_VALUE_ENUMERATION.Map_ATTRIBUTE_VALUE_ENUMERATIONDBID_ATTRIBUTE_VALUE_ENUMERATIONPtr[uint(_ATTRIBUTE_VALUE_ENUMERATIONid)])
	}

	return
}

// CommitA_ATTRIBUTE_VALUE_ENUMERATION allows commit of a single a_attribute_value_enumeration (if already staged)
func (backRepo *BackRepoStruct) CommitA_ATTRIBUTE_VALUE_ENUMERATION(a_attribute_value_enumeration *models.A_ATTRIBUTE_VALUE_ENUMERATION) {
	backRepo.BackRepoA_ATTRIBUTE_VALUE_ENUMERATION.CommitPhaseOneInstance(a_attribute_value_enumeration)
	if id, ok := backRepo.BackRepoA_ATTRIBUTE_VALUE_ENUMERATION.Map_A_ATTRIBUTE_VALUE_ENUMERATIONPtr_A_ATTRIBUTE_VALUE_ENUMERATIONDBID[a_attribute_value_enumeration]; ok {
		backRepo.BackRepoA_ATTRIBUTE_VALUE_ENUMERATION.CommitPhaseTwoInstance(backRepo, id, a_attribute_value_enumeration)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitA_ATTRIBUTE_VALUE_ENUMERATION allows checkout of a single a_attribute_value_enumeration (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutA_ATTRIBUTE_VALUE_ENUMERATION(a_attribute_value_enumeration *models.A_ATTRIBUTE_VALUE_ENUMERATION) {
	// check if the a_attribute_value_enumeration is staged
	if _, ok := backRepo.BackRepoA_ATTRIBUTE_VALUE_ENUMERATION.Map_A_ATTRIBUTE_VALUE_ENUMERATIONPtr_A_ATTRIBUTE_VALUE_ENUMERATIONDBID[a_attribute_value_enumeration]; ok {

		if id, ok := backRepo.BackRepoA_ATTRIBUTE_VALUE_ENUMERATION.Map_A_ATTRIBUTE_VALUE_ENUMERATIONPtr_A_ATTRIBUTE_VALUE_ENUMERATIONDBID[a_attribute_value_enumeration]; ok {
			var a_attribute_value_enumerationDB A_ATTRIBUTE_VALUE_ENUMERATIONDB
			a_attribute_value_enumerationDB.ID = id

			if err := backRepo.BackRepoA_ATTRIBUTE_VALUE_ENUMERATION.db.First(&a_attribute_value_enumerationDB, id).Error; err != nil {
				log.Fatalln("CheckoutA_ATTRIBUTE_VALUE_ENUMERATION : Problem with getting object with id:", id)
			}
			backRepo.BackRepoA_ATTRIBUTE_VALUE_ENUMERATION.CheckoutPhaseOneInstance(&a_attribute_value_enumerationDB)
			backRepo.BackRepoA_ATTRIBUTE_VALUE_ENUMERATION.CheckoutPhaseTwoInstance(backRepo, &a_attribute_value_enumerationDB)
		}
	}
}

// CopyBasicFieldsFromA_ATTRIBUTE_VALUE_ENUMERATION
func (a_attribute_value_enumerationDB *A_ATTRIBUTE_VALUE_ENUMERATIONDB) CopyBasicFieldsFromA_ATTRIBUTE_VALUE_ENUMERATION(a_attribute_value_enumeration *models.A_ATTRIBUTE_VALUE_ENUMERATION) {
	// insertion point for fields commit

	a_attribute_value_enumerationDB.Name_Data.String = a_attribute_value_enumeration.Name
	a_attribute_value_enumerationDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromA_ATTRIBUTE_VALUE_ENUMERATION_WOP
func (a_attribute_value_enumerationDB *A_ATTRIBUTE_VALUE_ENUMERATIONDB) CopyBasicFieldsFromA_ATTRIBUTE_VALUE_ENUMERATION_WOP(a_attribute_value_enumeration *models.A_ATTRIBUTE_VALUE_ENUMERATION_WOP) {
	// insertion point for fields commit

	a_attribute_value_enumerationDB.Name_Data.String = a_attribute_value_enumeration.Name
	a_attribute_value_enumerationDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromA_ATTRIBUTE_VALUE_ENUMERATIONWOP
func (a_attribute_value_enumerationDB *A_ATTRIBUTE_VALUE_ENUMERATIONDB) CopyBasicFieldsFromA_ATTRIBUTE_VALUE_ENUMERATIONWOP(a_attribute_value_enumeration *A_ATTRIBUTE_VALUE_ENUMERATIONWOP) {
	// insertion point for fields commit

	a_attribute_value_enumerationDB.Name_Data.String = a_attribute_value_enumeration.Name
	a_attribute_value_enumerationDB.Name_Data.Valid = true
}

// CopyBasicFieldsToA_ATTRIBUTE_VALUE_ENUMERATION
func (a_attribute_value_enumerationDB *A_ATTRIBUTE_VALUE_ENUMERATIONDB) CopyBasicFieldsToA_ATTRIBUTE_VALUE_ENUMERATION(a_attribute_value_enumeration *models.A_ATTRIBUTE_VALUE_ENUMERATION) {
	// insertion point for checkout of basic fields (back repo to stage)
	a_attribute_value_enumeration.Name = a_attribute_value_enumerationDB.Name_Data.String
}

// CopyBasicFieldsToA_ATTRIBUTE_VALUE_ENUMERATION_WOP
func (a_attribute_value_enumerationDB *A_ATTRIBUTE_VALUE_ENUMERATIONDB) CopyBasicFieldsToA_ATTRIBUTE_VALUE_ENUMERATION_WOP(a_attribute_value_enumeration *models.A_ATTRIBUTE_VALUE_ENUMERATION_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	a_attribute_value_enumeration.Name = a_attribute_value_enumerationDB.Name_Data.String
}

// CopyBasicFieldsToA_ATTRIBUTE_VALUE_ENUMERATIONWOP
func (a_attribute_value_enumerationDB *A_ATTRIBUTE_VALUE_ENUMERATIONDB) CopyBasicFieldsToA_ATTRIBUTE_VALUE_ENUMERATIONWOP(a_attribute_value_enumeration *A_ATTRIBUTE_VALUE_ENUMERATIONWOP) {
	a_attribute_value_enumeration.ID = int(a_attribute_value_enumerationDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	a_attribute_value_enumeration.Name = a_attribute_value_enumerationDB.Name_Data.String
}

// Backup generates a json file from a slice of all A_ATTRIBUTE_VALUE_ENUMERATIONDB instances in the backrepo
func (backRepoA_ATTRIBUTE_VALUE_ENUMERATION *BackRepoA_ATTRIBUTE_VALUE_ENUMERATIONStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "A_ATTRIBUTE_VALUE_ENUMERATIONDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*A_ATTRIBUTE_VALUE_ENUMERATIONDB, 0)
	for _, a_attribute_value_enumerationDB := range backRepoA_ATTRIBUTE_VALUE_ENUMERATION.Map_A_ATTRIBUTE_VALUE_ENUMERATIONDBID_A_ATTRIBUTE_VALUE_ENUMERATIONDB {
		forBackup = append(forBackup, a_attribute_value_enumerationDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json A_ATTRIBUTE_VALUE_ENUMERATION ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json A_ATTRIBUTE_VALUE_ENUMERATION file", err.Error())
	}
}

// Backup generates a json file from a slice of all A_ATTRIBUTE_VALUE_ENUMERATIONDB instances in the backrepo
func (backRepoA_ATTRIBUTE_VALUE_ENUMERATION *BackRepoA_ATTRIBUTE_VALUE_ENUMERATIONStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*A_ATTRIBUTE_VALUE_ENUMERATIONDB, 0)
	for _, a_attribute_value_enumerationDB := range backRepoA_ATTRIBUTE_VALUE_ENUMERATION.Map_A_ATTRIBUTE_VALUE_ENUMERATIONDBID_A_ATTRIBUTE_VALUE_ENUMERATIONDB {
		forBackup = append(forBackup, a_attribute_value_enumerationDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("A_ATTRIBUTE_VALUE_ENUMERATION")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&A_ATTRIBUTE_VALUE_ENUMERATION_Fields, -1)
	for _, a_attribute_value_enumerationDB := range forBackup {

		var a_attribute_value_enumerationWOP A_ATTRIBUTE_VALUE_ENUMERATIONWOP
		a_attribute_value_enumerationDB.CopyBasicFieldsToA_ATTRIBUTE_VALUE_ENUMERATIONWOP(&a_attribute_value_enumerationWOP)

		row := sh.AddRow()
		row.WriteStruct(&a_attribute_value_enumerationWOP, -1)
	}
}

// RestoreXL from the "A_ATTRIBUTE_VALUE_ENUMERATION" sheet all A_ATTRIBUTE_VALUE_ENUMERATIONDB instances
func (backRepoA_ATTRIBUTE_VALUE_ENUMERATION *BackRepoA_ATTRIBUTE_VALUE_ENUMERATIONStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoA_ATTRIBUTE_VALUE_ENUMERATIONid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["A_ATTRIBUTE_VALUE_ENUMERATION"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoA_ATTRIBUTE_VALUE_ENUMERATION.rowVisitorA_ATTRIBUTE_VALUE_ENUMERATION)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoA_ATTRIBUTE_VALUE_ENUMERATION *BackRepoA_ATTRIBUTE_VALUE_ENUMERATIONStruct) rowVisitorA_ATTRIBUTE_VALUE_ENUMERATION(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var a_attribute_value_enumerationWOP A_ATTRIBUTE_VALUE_ENUMERATIONWOP
		row.ReadStruct(&a_attribute_value_enumerationWOP)

		// add the unmarshalled struct to the stage
		a_attribute_value_enumerationDB := new(A_ATTRIBUTE_VALUE_ENUMERATIONDB)
		a_attribute_value_enumerationDB.CopyBasicFieldsFromA_ATTRIBUTE_VALUE_ENUMERATIONWOP(&a_attribute_value_enumerationWOP)

		a_attribute_value_enumerationDB_ID_atBackupTime := a_attribute_value_enumerationDB.ID
		a_attribute_value_enumerationDB.ID = 0
		query := backRepoA_ATTRIBUTE_VALUE_ENUMERATION.db.Create(a_attribute_value_enumerationDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoA_ATTRIBUTE_VALUE_ENUMERATION.Map_A_ATTRIBUTE_VALUE_ENUMERATIONDBID_A_ATTRIBUTE_VALUE_ENUMERATIONDB[a_attribute_value_enumerationDB.ID] = a_attribute_value_enumerationDB
		BackRepoA_ATTRIBUTE_VALUE_ENUMERATIONid_atBckpTime_newID[a_attribute_value_enumerationDB_ID_atBackupTime] = a_attribute_value_enumerationDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "A_ATTRIBUTE_VALUE_ENUMERATIONDB.json" in dirPath that stores an array
// of A_ATTRIBUTE_VALUE_ENUMERATIONDB and stores it in the database
// the map BackRepoA_ATTRIBUTE_VALUE_ENUMERATIONid_atBckpTime_newID is updated accordingly
func (backRepoA_ATTRIBUTE_VALUE_ENUMERATION *BackRepoA_ATTRIBUTE_VALUE_ENUMERATIONStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoA_ATTRIBUTE_VALUE_ENUMERATIONid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "A_ATTRIBUTE_VALUE_ENUMERATIONDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json A_ATTRIBUTE_VALUE_ENUMERATION file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*A_ATTRIBUTE_VALUE_ENUMERATIONDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_A_ATTRIBUTE_VALUE_ENUMERATIONDBID_A_ATTRIBUTE_VALUE_ENUMERATIONDB
	for _, a_attribute_value_enumerationDB := range forRestore {

		a_attribute_value_enumerationDB_ID_atBackupTime := a_attribute_value_enumerationDB.ID
		a_attribute_value_enumerationDB.ID = 0
		query := backRepoA_ATTRIBUTE_VALUE_ENUMERATION.db.Create(a_attribute_value_enumerationDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoA_ATTRIBUTE_VALUE_ENUMERATION.Map_A_ATTRIBUTE_VALUE_ENUMERATIONDBID_A_ATTRIBUTE_VALUE_ENUMERATIONDB[a_attribute_value_enumerationDB.ID] = a_attribute_value_enumerationDB
		BackRepoA_ATTRIBUTE_VALUE_ENUMERATIONid_atBckpTime_newID[a_attribute_value_enumerationDB_ID_atBackupTime] = a_attribute_value_enumerationDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json A_ATTRIBUTE_VALUE_ENUMERATION file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<A_ATTRIBUTE_VALUE_ENUMERATION>id_atBckpTime_newID
// to compute new index
func (backRepoA_ATTRIBUTE_VALUE_ENUMERATION *BackRepoA_ATTRIBUTE_VALUE_ENUMERATIONStruct) RestorePhaseTwo() {

	for _, a_attribute_value_enumerationDB := range backRepoA_ATTRIBUTE_VALUE_ENUMERATION.Map_A_ATTRIBUTE_VALUE_ENUMERATIONDBID_A_ATTRIBUTE_VALUE_ENUMERATIONDB {

		// next line of code is to avert unused variable compilation error
		_ = a_attribute_value_enumerationDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		query := backRepoA_ATTRIBUTE_VALUE_ENUMERATION.db.Model(a_attribute_value_enumerationDB).Updates(*a_attribute_value_enumerationDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
	}

}

// BackRepoA_ATTRIBUTE_VALUE_ENUMERATION.ResetReversePointers commits all staged instances of A_ATTRIBUTE_VALUE_ENUMERATION to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoA_ATTRIBUTE_VALUE_ENUMERATION *BackRepoA_ATTRIBUTE_VALUE_ENUMERATIONStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, a_attribute_value_enumeration := range backRepoA_ATTRIBUTE_VALUE_ENUMERATION.Map_A_ATTRIBUTE_VALUE_ENUMERATIONDBID_A_ATTRIBUTE_VALUE_ENUMERATIONPtr {
		backRepoA_ATTRIBUTE_VALUE_ENUMERATION.ResetReversePointersInstance(backRepo, idx, a_attribute_value_enumeration)
	}

	return
}

func (backRepoA_ATTRIBUTE_VALUE_ENUMERATION *BackRepoA_ATTRIBUTE_VALUE_ENUMERATIONStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, a_attribute_value_enumeration *models.A_ATTRIBUTE_VALUE_ENUMERATION) (Error error) {

	// fetch matching a_attribute_value_enumerationDB
	if a_attribute_value_enumerationDB, ok := backRepoA_ATTRIBUTE_VALUE_ENUMERATION.Map_A_ATTRIBUTE_VALUE_ENUMERATIONDBID_A_ATTRIBUTE_VALUE_ENUMERATIONDB[idx]; ok {
		_ = a_attribute_value_enumerationDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoA_ATTRIBUTE_VALUE_ENUMERATIONid_atBckpTime_newID map[uint]uint