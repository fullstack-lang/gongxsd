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

	"github.com/fullstack-lang/gongxsd/go/db"
	"github.com/fullstack-lang/gongxsd/go/models"
)

// dummy variable to have the import declaration wihthout compile failure (even if no code needing this import is generated)
var dummy_MinLength_sql sql.NullBool
var dummy_MinLength_time time.Duration
var dummy_MinLength_sort sort.Float64Slice

// MinLengthAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model minlengthAPI
type MinLengthAPI struct {
	gorm.Model

	models.MinLength_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	MinLengthPointersEncoding MinLengthPointersEncoding
}

// MinLengthPointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type MinLengthPointersEncoding struct {
	// insertion for pointer fields encoding declaration

	// field Annotation is a pointer to another Struct (optional or 0..1)
	// This field is generated into another field to enable AS ONE association
	AnnotationID sql.NullInt64
}

// MinLengthDB describes a minlength in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model minlengthDB
type MinLengthDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field minlengthDB.Name
	Name_Data sql.NullString

	// Declation for basic field minlengthDB.Value
	Value_Data sql.NullString

	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	MinLengthPointersEncoding
}

// MinLengthDBs arrays minlengthDBs
// swagger:response minlengthDBsResponse
type MinLengthDBs []MinLengthDB

// MinLengthDBResponse provides response
// swagger:response minlengthDBResponse
type MinLengthDBResponse struct {
	MinLengthDB
}

// MinLengthWOP is a MinLength without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type MinLengthWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	Value string `xlsx:"2"`
	// insertion for WOP pointer fields
}

var MinLength_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"Value",
}

type BackRepoMinLengthStruct struct {
	// stores MinLengthDB according to their gorm ID
	Map_MinLengthDBID_MinLengthDB map[uint]*MinLengthDB

	// stores MinLengthDB ID according to MinLength address
	Map_MinLengthPtr_MinLengthDBID map[*models.MinLength]uint

	// stores MinLength according to their gorm ID
	Map_MinLengthDBID_MinLengthPtr map[uint]*models.MinLength

	db db.DBInterface

	stage *models.Stage
}

func (backRepoMinLength *BackRepoMinLengthStruct) GetStage() (stage *models.Stage) {
	stage = backRepoMinLength.stage
	return
}

func (backRepoMinLength *BackRepoMinLengthStruct) GetDB() db.DBInterface {
	return backRepoMinLength.db
}

// GetMinLengthDBFromMinLengthPtr is a handy function to access the back repo instance from the stage instance
func (backRepoMinLength *BackRepoMinLengthStruct) GetMinLengthDBFromMinLengthPtr(minlength *models.MinLength) (minlengthDB *MinLengthDB) {
	id := backRepoMinLength.Map_MinLengthPtr_MinLengthDBID[minlength]
	minlengthDB = backRepoMinLength.Map_MinLengthDBID_MinLengthDB[id]
	return
}

// BackRepoMinLength.CommitPhaseOne commits all staged instances of MinLength to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoMinLength *BackRepoMinLengthStruct) CommitPhaseOne(stage *models.Stage) (Error error) {

	var minlengths []*models.MinLength
	for minlength := range stage.MinLengths {
		minlengths = append(minlengths, minlength)
	}

	// Sort by the order stored in Map_Staged_Order.
	sort.Slice(minlengths, func(i, j int) bool {
		return stage.MinLengthMap_Staged_Order[minlengths[i]] < stage.MinLengthMap_Staged_Order[minlengths[j]]
	})

	for _, minlength := range minlengths {
		backRepoMinLength.CommitPhaseOneInstance(minlength)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, minlength := range backRepoMinLength.Map_MinLengthDBID_MinLengthPtr {
		if _, ok := stage.MinLengths[minlength]; !ok {
			backRepoMinLength.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoMinLength.CommitDeleteInstance commits deletion of MinLength to the BackRepo
func (backRepoMinLength *BackRepoMinLengthStruct) CommitDeleteInstance(id uint) (Error error) {

	minlength := backRepoMinLength.Map_MinLengthDBID_MinLengthPtr[id]

	// minlength is not staged anymore, remove minlengthDB
	minlengthDB := backRepoMinLength.Map_MinLengthDBID_MinLengthDB[id]
	db, _ := backRepoMinLength.db.Unscoped()
	_, err := db.Delete(minlengthDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	delete(backRepoMinLength.Map_MinLengthPtr_MinLengthDBID, minlength)
	delete(backRepoMinLength.Map_MinLengthDBID_MinLengthPtr, id)
	delete(backRepoMinLength.Map_MinLengthDBID_MinLengthDB, id)

	return
}

// BackRepoMinLength.CommitPhaseOneInstance commits minlength staged instances of MinLength to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoMinLength *BackRepoMinLengthStruct) CommitPhaseOneInstance(minlength *models.MinLength) (Error error) {

	// check if the minlength is not commited yet
	if _, ok := backRepoMinLength.Map_MinLengthPtr_MinLengthDBID[minlength]; ok {
		return
	}

	// initiate minlength
	var minlengthDB MinLengthDB
	minlengthDB.CopyBasicFieldsFromMinLength(minlength)

	_, err := backRepoMinLength.db.Create(&minlengthDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	backRepoMinLength.Map_MinLengthPtr_MinLengthDBID[minlength] = minlengthDB.ID
	backRepoMinLength.Map_MinLengthDBID_MinLengthPtr[minlengthDB.ID] = minlength
	backRepoMinLength.Map_MinLengthDBID_MinLengthDB[minlengthDB.ID] = &minlengthDB

	return
}

// BackRepoMinLength.CommitPhaseTwo commits all staged instances of MinLength to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoMinLength *BackRepoMinLengthStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, minlength := range backRepoMinLength.Map_MinLengthDBID_MinLengthPtr {
		backRepoMinLength.CommitPhaseTwoInstance(backRepo, idx, minlength)
	}

	return
}

// BackRepoMinLength.CommitPhaseTwoInstance commits {{structname }} of models.MinLength to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoMinLength *BackRepoMinLengthStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, minlength *models.MinLength) (Error error) {

	// fetch matching minlengthDB
	if minlengthDB, ok := backRepoMinLength.Map_MinLengthDBID_MinLengthDB[idx]; ok {

		minlengthDB.CopyBasicFieldsFromMinLength(minlength)

		// insertion point for translating pointers encodings into actual pointers
		// commit pointer value minlength.Annotation translates to updating the minlength.AnnotationID
		minlengthDB.AnnotationID.Valid = true // allow for a 0 value (nil association)
		if minlength.Annotation != nil {
			if AnnotationId, ok := backRepo.BackRepoAnnotation.Map_AnnotationPtr_AnnotationDBID[minlength.Annotation]; ok {
				minlengthDB.AnnotationID.Int64 = int64(AnnotationId)
				minlengthDB.AnnotationID.Valid = true
			}
		} else {
			minlengthDB.AnnotationID.Int64 = 0
			minlengthDB.AnnotationID.Valid = true
		}

		_, err := backRepoMinLength.db.Save(minlengthDB)
		if err != nil {
			log.Fatal(err)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown MinLength intance %s", minlength.Name))
		return err
	}

	return
}

// BackRepoMinLength.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoMinLength *BackRepoMinLengthStruct) CheckoutPhaseOne() (Error error) {

	minlengthDBArray := make([]MinLengthDB, 0)
	_, err := backRepoMinLength.db.Find(&minlengthDBArray)
	if err != nil {
		return err
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	minlengthInstancesToBeRemovedFromTheStage := make(map[*models.MinLength]any)
	for key, value := range backRepoMinLength.stage.MinLengths {
		minlengthInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, minlengthDB := range minlengthDBArray {
		backRepoMinLength.CheckoutPhaseOneInstance(&minlengthDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		minlength, ok := backRepoMinLength.Map_MinLengthDBID_MinLengthPtr[minlengthDB.ID]
		if ok {
			delete(minlengthInstancesToBeRemovedFromTheStage, minlength)
		}
	}

	// remove from stage and back repo's 3 maps all minlengths that are not in the checkout
	for minlength := range minlengthInstancesToBeRemovedFromTheStage {
		minlength.Unstage(backRepoMinLength.GetStage())

		// remove instance from the back repo 3 maps
		minlengthID := backRepoMinLength.Map_MinLengthPtr_MinLengthDBID[minlength]
		delete(backRepoMinLength.Map_MinLengthPtr_MinLengthDBID, minlength)
		delete(backRepoMinLength.Map_MinLengthDBID_MinLengthDB, minlengthID)
		delete(backRepoMinLength.Map_MinLengthDBID_MinLengthPtr, minlengthID)
	}

	return
}

// CheckoutPhaseOneInstance takes a minlengthDB that has been found in the DB, updates the backRepo and stages the
// models version of the minlengthDB
func (backRepoMinLength *BackRepoMinLengthStruct) CheckoutPhaseOneInstance(minlengthDB *MinLengthDB) (Error error) {

	minlength, ok := backRepoMinLength.Map_MinLengthDBID_MinLengthPtr[minlengthDB.ID]
	if !ok {
		minlength = new(models.MinLength)

		backRepoMinLength.Map_MinLengthDBID_MinLengthPtr[minlengthDB.ID] = minlength
		backRepoMinLength.Map_MinLengthPtr_MinLengthDBID[minlength] = minlengthDB.ID

		// append model store with the new element
		minlength.Name = minlengthDB.Name_Data.String
		minlength.Stage(backRepoMinLength.GetStage())
	}
	minlengthDB.CopyBasicFieldsToMinLength(minlength)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	minlength.Stage(backRepoMinLength.GetStage())

	// preserve pointer to minlengthDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_MinLengthDBID_MinLengthDB)[minlengthDB hold variable pointers
	minlengthDB_Data := *minlengthDB
	preservedPtrToMinLength := &minlengthDB_Data
	backRepoMinLength.Map_MinLengthDBID_MinLengthDB[minlengthDB.ID] = preservedPtrToMinLength

	return
}

// BackRepoMinLength.CheckoutPhaseTwo Checkouts all staged instances of MinLength to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoMinLength *BackRepoMinLengthStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, minlengthDB := range backRepoMinLength.Map_MinLengthDBID_MinLengthDB {
		backRepoMinLength.CheckoutPhaseTwoInstance(backRepo, minlengthDB)
	}
	return
}

// BackRepoMinLength.CheckoutPhaseTwoInstance Checkouts staged instances of MinLength to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoMinLength *BackRepoMinLengthStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, minlengthDB *MinLengthDB) (Error error) {

	minlength := backRepoMinLength.Map_MinLengthDBID_MinLengthPtr[minlengthDB.ID]

	minlengthDB.DecodePointers(backRepo, minlength)

	return
}

func (minlengthDB *MinLengthDB) DecodePointers(backRepo *BackRepoStruct, minlength *models.MinLength) {

	// insertion point for checkout of pointer encoding
	// Annotation field	
	{
		id := minlengthDB.AnnotationID.Int64
		if id != 0 {
			tmp, ok := backRepo.BackRepoAnnotation.Map_AnnotationDBID_AnnotationPtr[uint(id)]

			// if the pointer id is unknown, it is not a problem, maybe the target was removed from the front
			if !ok {
				log.Println("DecodePointers: minlength.Annotation, unknown pointer id", id)
				minlength.Annotation = nil
			} else {
				// updates only if field has changed
				if minlength.Annotation == nil || minlength.Annotation != tmp {
					minlength.Annotation = tmp
				}
			}
		} else {
			minlength.Annotation = nil
		}
	}
	
	return
}

// CommitMinLength allows commit of a single minlength (if already staged)
func (backRepo *BackRepoStruct) CommitMinLength(minlength *models.MinLength) {
	backRepo.BackRepoMinLength.CommitPhaseOneInstance(minlength)
	if id, ok := backRepo.BackRepoMinLength.Map_MinLengthPtr_MinLengthDBID[minlength]; ok {
		backRepo.BackRepoMinLength.CommitPhaseTwoInstance(backRepo, id, minlength)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitMinLength allows checkout of a single minlength (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutMinLength(minlength *models.MinLength) {
	// check if the minlength is staged
	if _, ok := backRepo.BackRepoMinLength.Map_MinLengthPtr_MinLengthDBID[minlength]; ok {

		if id, ok := backRepo.BackRepoMinLength.Map_MinLengthPtr_MinLengthDBID[minlength]; ok {
			var minlengthDB MinLengthDB
			minlengthDB.ID = id

			if _, err := backRepo.BackRepoMinLength.db.First(&minlengthDB, id); err != nil {
				log.Fatalln("CheckoutMinLength : Problem with getting object with id:", id)
			}
			backRepo.BackRepoMinLength.CheckoutPhaseOneInstance(&minlengthDB)
			backRepo.BackRepoMinLength.CheckoutPhaseTwoInstance(backRepo, &minlengthDB)
		}
	}
}

// CopyBasicFieldsFromMinLength
func (minlengthDB *MinLengthDB) CopyBasicFieldsFromMinLength(minlength *models.MinLength) {
	// insertion point for fields commit

	minlengthDB.Name_Data.String = minlength.Name
	minlengthDB.Name_Data.Valid = true

	minlengthDB.Value_Data.String = minlength.Value
	minlengthDB.Value_Data.Valid = true
}

// CopyBasicFieldsFromMinLength_WOP
func (minlengthDB *MinLengthDB) CopyBasicFieldsFromMinLength_WOP(minlength *models.MinLength_WOP) {
	// insertion point for fields commit

	minlengthDB.Name_Data.String = minlength.Name
	minlengthDB.Name_Data.Valid = true

	minlengthDB.Value_Data.String = minlength.Value
	minlengthDB.Value_Data.Valid = true
}

// CopyBasicFieldsFromMinLengthWOP
func (minlengthDB *MinLengthDB) CopyBasicFieldsFromMinLengthWOP(minlength *MinLengthWOP) {
	// insertion point for fields commit

	minlengthDB.Name_Data.String = minlength.Name
	minlengthDB.Name_Data.Valid = true

	minlengthDB.Value_Data.String = minlength.Value
	minlengthDB.Value_Data.Valid = true
}

// CopyBasicFieldsToMinLength
func (minlengthDB *MinLengthDB) CopyBasicFieldsToMinLength(minlength *models.MinLength) {
	// insertion point for checkout of basic fields (back repo to stage)
	minlength.Name = minlengthDB.Name_Data.String
	minlength.Value = minlengthDB.Value_Data.String
}

// CopyBasicFieldsToMinLength_WOP
func (minlengthDB *MinLengthDB) CopyBasicFieldsToMinLength_WOP(minlength *models.MinLength_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	minlength.Name = minlengthDB.Name_Data.String
	minlength.Value = minlengthDB.Value_Data.String
}

// CopyBasicFieldsToMinLengthWOP
func (minlengthDB *MinLengthDB) CopyBasicFieldsToMinLengthWOP(minlength *MinLengthWOP) {
	minlength.ID = int(minlengthDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	minlength.Name = minlengthDB.Name_Data.String
	minlength.Value = minlengthDB.Value_Data.String
}

// Backup generates a json file from a slice of all MinLengthDB instances in the backrepo
func (backRepoMinLength *BackRepoMinLengthStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "MinLengthDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*MinLengthDB, 0)
	for _, minlengthDB := range backRepoMinLength.Map_MinLengthDBID_MinLengthDB {
		forBackup = append(forBackup, minlengthDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json MinLength ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json MinLength file", err.Error())
	}
}

// Backup generates a json file from a slice of all MinLengthDB instances in the backrepo
func (backRepoMinLength *BackRepoMinLengthStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*MinLengthDB, 0)
	for _, minlengthDB := range backRepoMinLength.Map_MinLengthDBID_MinLengthDB {
		forBackup = append(forBackup, minlengthDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("MinLength")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&MinLength_Fields, -1)
	for _, minlengthDB := range forBackup {

		var minlengthWOP MinLengthWOP
		minlengthDB.CopyBasicFieldsToMinLengthWOP(&minlengthWOP)

		row := sh.AddRow()
		row.WriteStruct(&minlengthWOP, -1)
	}
}

// RestoreXL from the "MinLength" sheet all MinLengthDB instances
func (backRepoMinLength *BackRepoMinLengthStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoMinLengthid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["MinLength"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoMinLength.rowVisitorMinLength)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoMinLength *BackRepoMinLengthStruct) rowVisitorMinLength(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var minlengthWOP MinLengthWOP
		row.ReadStruct(&minlengthWOP)

		// add the unmarshalled struct to the stage
		minlengthDB := new(MinLengthDB)
		minlengthDB.CopyBasicFieldsFromMinLengthWOP(&minlengthWOP)

		minlengthDB_ID_atBackupTime := minlengthDB.ID
		minlengthDB.ID = 0
		_, err := backRepoMinLength.db.Create(minlengthDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoMinLength.Map_MinLengthDBID_MinLengthDB[minlengthDB.ID] = minlengthDB
		BackRepoMinLengthid_atBckpTime_newID[minlengthDB_ID_atBackupTime] = minlengthDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "MinLengthDB.json" in dirPath that stores an array
// of MinLengthDB and stores it in the database
// the map BackRepoMinLengthid_atBckpTime_newID is updated accordingly
func (backRepoMinLength *BackRepoMinLengthStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoMinLengthid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "MinLengthDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json MinLength file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*MinLengthDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_MinLengthDBID_MinLengthDB
	for _, minlengthDB := range forRestore {

		minlengthDB_ID_atBackupTime := minlengthDB.ID
		minlengthDB.ID = 0
		_, err := backRepoMinLength.db.Create(minlengthDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoMinLength.Map_MinLengthDBID_MinLengthDB[minlengthDB.ID] = minlengthDB
		BackRepoMinLengthid_atBckpTime_newID[minlengthDB_ID_atBackupTime] = minlengthDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json MinLength file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<MinLength>id_atBckpTime_newID
// to compute new index
func (backRepoMinLength *BackRepoMinLengthStruct) RestorePhaseTwo() {

	for _, minlengthDB := range backRepoMinLength.Map_MinLengthDBID_MinLengthDB {

		// next line of code is to avert unused variable compilation error
		_ = minlengthDB

		// insertion point for reindexing pointers encoding
		// reindexing Annotation field
		if minlengthDB.AnnotationID.Int64 != 0 {
			minlengthDB.AnnotationID.Int64 = int64(BackRepoAnnotationid_atBckpTime_newID[uint(minlengthDB.AnnotationID.Int64)])
			minlengthDB.AnnotationID.Valid = true
		}

		// update databse with new index encoding
		db, _ := backRepoMinLength.db.Model(minlengthDB)
		_, err := db.Updates(*minlengthDB)
		if err != nil {
			log.Fatal(err)
		}
	}

}

// BackRepoMinLength.ResetReversePointers commits all staged instances of MinLength to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoMinLength *BackRepoMinLengthStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, minlength := range backRepoMinLength.Map_MinLengthDBID_MinLengthPtr {
		backRepoMinLength.ResetReversePointersInstance(backRepo, idx, minlength)
	}

	return
}

func (backRepoMinLength *BackRepoMinLengthStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, minlength *models.MinLength) (Error error) {

	// fetch matching minlengthDB
	if minlengthDB, ok := backRepoMinLength.Map_MinLengthDBID_MinLengthDB[idx]; ok {
		_ = minlengthDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoMinLengthid_atBckpTime_newID map[uint]uint
