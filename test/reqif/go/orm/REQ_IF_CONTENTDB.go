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
var dummy_REQ_IF_CONTENT_sql sql.NullBool
var dummy_REQ_IF_CONTENT_time time.Duration
var dummy_REQ_IF_CONTENT_sort sort.Float64Slice

// REQ_IF_CONTENTAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model req_if_contentAPI
type REQ_IF_CONTENTAPI struct {
	gorm.Model

	models.REQ_IF_CONTENT_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	REQ_IF_CONTENTPointersEncoding REQ_IF_CONTENTPointersEncoding
}

// REQ_IF_CONTENTPointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type REQ_IF_CONTENTPointersEncoding struct {
	// insertion for pointer fields encoding declaration
}

// REQ_IF_CONTENTDB describes a req_if_content in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model req_if_contentDB
type REQ_IF_CONTENTDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field req_if_contentDB.Name
	Name_Data sql.NullString
	
	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	REQ_IF_CONTENTPointersEncoding
}

// REQ_IF_CONTENTDBs arrays req_if_contentDBs
// swagger:response req_if_contentDBsResponse
type REQ_IF_CONTENTDBs []REQ_IF_CONTENTDB

// REQ_IF_CONTENTDBResponse provides response
// swagger:response req_if_contentDBResponse
type REQ_IF_CONTENTDBResponse struct {
	REQ_IF_CONTENTDB
}

// REQ_IF_CONTENTWOP is a REQ_IF_CONTENT without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type REQ_IF_CONTENTWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`
	// insertion for WOP pointer fields
}

var REQ_IF_CONTENT_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
}

type BackRepoREQ_IF_CONTENTStruct struct {
	// stores REQ_IF_CONTENTDB according to their gorm ID
	Map_REQ_IF_CONTENTDBID_REQ_IF_CONTENTDB map[uint]*REQ_IF_CONTENTDB

	// stores REQ_IF_CONTENTDB ID according to REQ_IF_CONTENT address
	Map_REQ_IF_CONTENTPtr_REQ_IF_CONTENTDBID map[*models.REQ_IF_CONTENT]uint

	// stores REQ_IF_CONTENT according to their gorm ID
	Map_REQ_IF_CONTENTDBID_REQ_IF_CONTENTPtr map[uint]*models.REQ_IF_CONTENT

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoREQ_IF_CONTENT *BackRepoREQ_IF_CONTENTStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoREQ_IF_CONTENT.stage
	return
}

func (backRepoREQ_IF_CONTENT *BackRepoREQ_IF_CONTENTStruct) GetDB() *gorm.DB {
	return backRepoREQ_IF_CONTENT.db
}

// GetREQ_IF_CONTENTDBFromREQ_IF_CONTENTPtr is a handy function to access the back repo instance from the stage instance
func (backRepoREQ_IF_CONTENT *BackRepoREQ_IF_CONTENTStruct) GetREQ_IF_CONTENTDBFromREQ_IF_CONTENTPtr(req_if_content *models.REQ_IF_CONTENT) (req_if_contentDB *REQ_IF_CONTENTDB) {
	id := backRepoREQ_IF_CONTENT.Map_REQ_IF_CONTENTPtr_REQ_IF_CONTENTDBID[req_if_content]
	req_if_contentDB = backRepoREQ_IF_CONTENT.Map_REQ_IF_CONTENTDBID_REQ_IF_CONTENTDB[id]
	return
}

// BackRepoREQ_IF_CONTENT.CommitPhaseOne commits all staged instances of REQ_IF_CONTENT to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoREQ_IF_CONTENT *BackRepoREQ_IF_CONTENTStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for req_if_content := range stage.REQ_IF_CONTENTs {
		backRepoREQ_IF_CONTENT.CommitPhaseOneInstance(req_if_content)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, req_if_content := range backRepoREQ_IF_CONTENT.Map_REQ_IF_CONTENTDBID_REQ_IF_CONTENTPtr {
		if _, ok := stage.REQ_IF_CONTENTs[req_if_content]; !ok {
			backRepoREQ_IF_CONTENT.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoREQ_IF_CONTENT.CommitDeleteInstance commits deletion of REQ_IF_CONTENT to the BackRepo
func (backRepoREQ_IF_CONTENT *BackRepoREQ_IF_CONTENTStruct) CommitDeleteInstance(id uint) (Error error) {

	req_if_content := backRepoREQ_IF_CONTENT.Map_REQ_IF_CONTENTDBID_REQ_IF_CONTENTPtr[id]

	// req_if_content is not staged anymore, remove req_if_contentDB
	req_if_contentDB := backRepoREQ_IF_CONTENT.Map_REQ_IF_CONTENTDBID_REQ_IF_CONTENTDB[id]
	query := backRepoREQ_IF_CONTENT.db.Unscoped().Delete(&req_if_contentDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	delete(backRepoREQ_IF_CONTENT.Map_REQ_IF_CONTENTPtr_REQ_IF_CONTENTDBID, req_if_content)
	delete(backRepoREQ_IF_CONTENT.Map_REQ_IF_CONTENTDBID_REQ_IF_CONTENTPtr, id)
	delete(backRepoREQ_IF_CONTENT.Map_REQ_IF_CONTENTDBID_REQ_IF_CONTENTDB, id)

	return
}

// BackRepoREQ_IF_CONTENT.CommitPhaseOneInstance commits req_if_content staged instances of REQ_IF_CONTENT to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoREQ_IF_CONTENT *BackRepoREQ_IF_CONTENTStruct) CommitPhaseOneInstance(req_if_content *models.REQ_IF_CONTENT) (Error error) {

	// check if the req_if_content is not commited yet
	if _, ok := backRepoREQ_IF_CONTENT.Map_REQ_IF_CONTENTPtr_REQ_IF_CONTENTDBID[req_if_content]; ok {
		return
	}

	// initiate req_if_content
	var req_if_contentDB REQ_IF_CONTENTDB
	req_if_contentDB.CopyBasicFieldsFromREQ_IF_CONTENT(req_if_content)

	query := backRepoREQ_IF_CONTENT.db.Create(&req_if_contentDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	backRepoREQ_IF_CONTENT.Map_REQ_IF_CONTENTPtr_REQ_IF_CONTENTDBID[req_if_content] = req_if_contentDB.ID
	backRepoREQ_IF_CONTENT.Map_REQ_IF_CONTENTDBID_REQ_IF_CONTENTPtr[req_if_contentDB.ID] = req_if_content
	backRepoREQ_IF_CONTENT.Map_REQ_IF_CONTENTDBID_REQ_IF_CONTENTDB[req_if_contentDB.ID] = &req_if_contentDB

	return
}

// BackRepoREQ_IF_CONTENT.CommitPhaseTwo commits all staged instances of REQ_IF_CONTENT to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoREQ_IF_CONTENT *BackRepoREQ_IF_CONTENTStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, req_if_content := range backRepoREQ_IF_CONTENT.Map_REQ_IF_CONTENTDBID_REQ_IF_CONTENTPtr {
		backRepoREQ_IF_CONTENT.CommitPhaseTwoInstance(backRepo, idx, req_if_content)
	}

	return
}

// BackRepoREQ_IF_CONTENT.CommitPhaseTwoInstance commits {{structname }} of models.REQ_IF_CONTENT to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoREQ_IF_CONTENT *BackRepoREQ_IF_CONTENTStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, req_if_content *models.REQ_IF_CONTENT) (Error error) {

	// fetch matching req_if_contentDB
	if req_if_contentDB, ok := backRepoREQ_IF_CONTENT.Map_REQ_IF_CONTENTDBID_REQ_IF_CONTENTDB[idx]; ok {

		req_if_contentDB.CopyBasicFieldsFromREQ_IF_CONTENT(req_if_content)

		// insertion point for translating pointers encodings into actual pointers
		query := backRepoREQ_IF_CONTENT.db.Save(&req_if_contentDB)
		if query.Error != nil {
			log.Fatalln(query.Error)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown REQ_IF_CONTENT intance %s", req_if_content.Name))
		return err
	}

	return
}

// BackRepoREQ_IF_CONTENT.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoREQ_IF_CONTENT *BackRepoREQ_IF_CONTENTStruct) CheckoutPhaseOne() (Error error) {

	req_if_contentDBArray := make([]REQ_IF_CONTENTDB, 0)
	query := backRepoREQ_IF_CONTENT.db.Find(&req_if_contentDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	req_if_contentInstancesToBeRemovedFromTheStage := make(map[*models.REQ_IF_CONTENT]any)
	for key, value := range backRepoREQ_IF_CONTENT.stage.REQ_IF_CONTENTs {
		req_if_contentInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, req_if_contentDB := range req_if_contentDBArray {
		backRepoREQ_IF_CONTENT.CheckoutPhaseOneInstance(&req_if_contentDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		req_if_content, ok := backRepoREQ_IF_CONTENT.Map_REQ_IF_CONTENTDBID_REQ_IF_CONTENTPtr[req_if_contentDB.ID]
		if ok {
			delete(req_if_contentInstancesToBeRemovedFromTheStage, req_if_content)
		}
	}

	// remove from stage and back repo's 3 maps all req_if_contents that are not in the checkout
	for req_if_content := range req_if_contentInstancesToBeRemovedFromTheStage {
		req_if_content.Unstage(backRepoREQ_IF_CONTENT.GetStage())

		// remove instance from the back repo 3 maps
		req_if_contentID := backRepoREQ_IF_CONTENT.Map_REQ_IF_CONTENTPtr_REQ_IF_CONTENTDBID[req_if_content]
		delete(backRepoREQ_IF_CONTENT.Map_REQ_IF_CONTENTPtr_REQ_IF_CONTENTDBID, req_if_content)
		delete(backRepoREQ_IF_CONTENT.Map_REQ_IF_CONTENTDBID_REQ_IF_CONTENTDB, req_if_contentID)
		delete(backRepoREQ_IF_CONTENT.Map_REQ_IF_CONTENTDBID_REQ_IF_CONTENTPtr, req_if_contentID)
	}

	return
}

// CheckoutPhaseOneInstance takes a req_if_contentDB that has been found in the DB, updates the backRepo and stages the
// models version of the req_if_contentDB
func (backRepoREQ_IF_CONTENT *BackRepoREQ_IF_CONTENTStruct) CheckoutPhaseOneInstance(req_if_contentDB *REQ_IF_CONTENTDB) (Error error) {

	req_if_content, ok := backRepoREQ_IF_CONTENT.Map_REQ_IF_CONTENTDBID_REQ_IF_CONTENTPtr[req_if_contentDB.ID]
	if !ok {
		req_if_content = new(models.REQ_IF_CONTENT)

		backRepoREQ_IF_CONTENT.Map_REQ_IF_CONTENTDBID_REQ_IF_CONTENTPtr[req_if_contentDB.ID] = req_if_content
		backRepoREQ_IF_CONTENT.Map_REQ_IF_CONTENTPtr_REQ_IF_CONTENTDBID[req_if_content] = req_if_contentDB.ID

		// append model store with the new element
		req_if_content.Name = req_if_contentDB.Name_Data.String
		req_if_content.Stage(backRepoREQ_IF_CONTENT.GetStage())
	}
	req_if_contentDB.CopyBasicFieldsToREQ_IF_CONTENT(req_if_content)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	req_if_content.Stage(backRepoREQ_IF_CONTENT.GetStage())

	// preserve pointer to req_if_contentDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_REQ_IF_CONTENTDBID_REQ_IF_CONTENTDB)[req_if_contentDB hold variable pointers
	req_if_contentDB_Data := *req_if_contentDB
	preservedPtrToREQ_IF_CONTENT := &req_if_contentDB_Data
	backRepoREQ_IF_CONTENT.Map_REQ_IF_CONTENTDBID_REQ_IF_CONTENTDB[req_if_contentDB.ID] = preservedPtrToREQ_IF_CONTENT

	return
}

// BackRepoREQ_IF_CONTENT.CheckoutPhaseTwo Checkouts all staged instances of REQ_IF_CONTENT to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoREQ_IF_CONTENT *BackRepoREQ_IF_CONTENTStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, req_if_contentDB := range backRepoREQ_IF_CONTENT.Map_REQ_IF_CONTENTDBID_REQ_IF_CONTENTDB {
		backRepoREQ_IF_CONTENT.CheckoutPhaseTwoInstance(backRepo, req_if_contentDB)
	}
	return
}

// BackRepoREQ_IF_CONTENT.CheckoutPhaseTwoInstance Checkouts staged instances of REQ_IF_CONTENT to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoREQ_IF_CONTENT *BackRepoREQ_IF_CONTENTStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, req_if_contentDB *REQ_IF_CONTENTDB) (Error error) {

	req_if_content := backRepoREQ_IF_CONTENT.Map_REQ_IF_CONTENTDBID_REQ_IF_CONTENTPtr[req_if_contentDB.ID]

	req_if_contentDB.DecodePointers(backRepo, req_if_content)

	return
}

func (req_if_contentDB *REQ_IF_CONTENTDB) DecodePointers(backRepo *BackRepoStruct, req_if_content *models.REQ_IF_CONTENT) {

	// insertion point for checkout of pointer encoding
	return
}

// CommitREQ_IF_CONTENT allows commit of a single req_if_content (if already staged)
func (backRepo *BackRepoStruct) CommitREQ_IF_CONTENT(req_if_content *models.REQ_IF_CONTENT) {
	backRepo.BackRepoREQ_IF_CONTENT.CommitPhaseOneInstance(req_if_content)
	if id, ok := backRepo.BackRepoREQ_IF_CONTENT.Map_REQ_IF_CONTENTPtr_REQ_IF_CONTENTDBID[req_if_content]; ok {
		backRepo.BackRepoREQ_IF_CONTENT.CommitPhaseTwoInstance(backRepo, id, req_if_content)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitREQ_IF_CONTENT allows checkout of a single req_if_content (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutREQ_IF_CONTENT(req_if_content *models.REQ_IF_CONTENT) {
	// check if the req_if_content is staged
	if _, ok := backRepo.BackRepoREQ_IF_CONTENT.Map_REQ_IF_CONTENTPtr_REQ_IF_CONTENTDBID[req_if_content]; ok {

		if id, ok := backRepo.BackRepoREQ_IF_CONTENT.Map_REQ_IF_CONTENTPtr_REQ_IF_CONTENTDBID[req_if_content]; ok {
			var req_if_contentDB REQ_IF_CONTENTDB
			req_if_contentDB.ID = id

			if err := backRepo.BackRepoREQ_IF_CONTENT.db.First(&req_if_contentDB, id).Error; err != nil {
				log.Fatalln("CheckoutREQ_IF_CONTENT : Problem with getting object with id:", id)
			}
			backRepo.BackRepoREQ_IF_CONTENT.CheckoutPhaseOneInstance(&req_if_contentDB)
			backRepo.BackRepoREQ_IF_CONTENT.CheckoutPhaseTwoInstance(backRepo, &req_if_contentDB)
		}
	}
}

// CopyBasicFieldsFromREQ_IF_CONTENT
func (req_if_contentDB *REQ_IF_CONTENTDB) CopyBasicFieldsFromREQ_IF_CONTENT(req_if_content *models.REQ_IF_CONTENT) {
	// insertion point for fields commit

	req_if_contentDB.Name_Data.String = req_if_content.Name
	req_if_contentDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromREQ_IF_CONTENT_WOP
func (req_if_contentDB *REQ_IF_CONTENTDB) CopyBasicFieldsFromREQ_IF_CONTENT_WOP(req_if_content *models.REQ_IF_CONTENT_WOP) {
	// insertion point for fields commit

	req_if_contentDB.Name_Data.String = req_if_content.Name
	req_if_contentDB.Name_Data.Valid = true
}

// CopyBasicFieldsFromREQ_IF_CONTENTWOP
func (req_if_contentDB *REQ_IF_CONTENTDB) CopyBasicFieldsFromREQ_IF_CONTENTWOP(req_if_content *REQ_IF_CONTENTWOP) {
	// insertion point for fields commit

	req_if_contentDB.Name_Data.String = req_if_content.Name
	req_if_contentDB.Name_Data.Valid = true
}

// CopyBasicFieldsToREQ_IF_CONTENT
func (req_if_contentDB *REQ_IF_CONTENTDB) CopyBasicFieldsToREQ_IF_CONTENT(req_if_content *models.REQ_IF_CONTENT) {
	// insertion point for checkout of basic fields (back repo to stage)
	req_if_content.Name = req_if_contentDB.Name_Data.String
}

// CopyBasicFieldsToREQ_IF_CONTENT_WOP
func (req_if_contentDB *REQ_IF_CONTENTDB) CopyBasicFieldsToREQ_IF_CONTENT_WOP(req_if_content *models.REQ_IF_CONTENT_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	req_if_content.Name = req_if_contentDB.Name_Data.String
}

// CopyBasicFieldsToREQ_IF_CONTENTWOP
func (req_if_contentDB *REQ_IF_CONTENTDB) CopyBasicFieldsToREQ_IF_CONTENTWOP(req_if_content *REQ_IF_CONTENTWOP) {
	req_if_content.ID = int(req_if_contentDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	req_if_content.Name = req_if_contentDB.Name_Data.String
}

// Backup generates a json file from a slice of all REQ_IF_CONTENTDB instances in the backrepo
func (backRepoREQ_IF_CONTENT *BackRepoREQ_IF_CONTENTStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "REQ_IF_CONTENTDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*REQ_IF_CONTENTDB, 0)
	for _, req_if_contentDB := range backRepoREQ_IF_CONTENT.Map_REQ_IF_CONTENTDBID_REQ_IF_CONTENTDB {
		forBackup = append(forBackup, req_if_contentDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json REQ_IF_CONTENT ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json REQ_IF_CONTENT file", err.Error())
	}
}

// Backup generates a json file from a slice of all REQ_IF_CONTENTDB instances in the backrepo
func (backRepoREQ_IF_CONTENT *BackRepoREQ_IF_CONTENTStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*REQ_IF_CONTENTDB, 0)
	for _, req_if_contentDB := range backRepoREQ_IF_CONTENT.Map_REQ_IF_CONTENTDBID_REQ_IF_CONTENTDB {
		forBackup = append(forBackup, req_if_contentDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("REQ_IF_CONTENT")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&REQ_IF_CONTENT_Fields, -1)
	for _, req_if_contentDB := range forBackup {

		var req_if_contentWOP REQ_IF_CONTENTWOP
		req_if_contentDB.CopyBasicFieldsToREQ_IF_CONTENTWOP(&req_if_contentWOP)

		row := sh.AddRow()
		row.WriteStruct(&req_if_contentWOP, -1)
	}
}

// RestoreXL from the "REQ_IF_CONTENT" sheet all REQ_IF_CONTENTDB instances
func (backRepoREQ_IF_CONTENT *BackRepoREQ_IF_CONTENTStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoREQ_IF_CONTENTid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["REQ_IF_CONTENT"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoREQ_IF_CONTENT.rowVisitorREQ_IF_CONTENT)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoREQ_IF_CONTENT *BackRepoREQ_IF_CONTENTStruct) rowVisitorREQ_IF_CONTENT(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var req_if_contentWOP REQ_IF_CONTENTWOP
		row.ReadStruct(&req_if_contentWOP)

		// add the unmarshalled struct to the stage
		req_if_contentDB := new(REQ_IF_CONTENTDB)
		req_if_contentDB.CopyBasicFieldsFromREQ_IF_CONTENTWOP(&req_if_contentWOP)

		req_if_contentDB_ID_atBackupTime := req_if_contentDB.ID
		req_if_contentDB.ID = 0
		query := backRepoREQ_IF_CONTENT.db.Create(req_if_contentDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoREQ_IF_CONTENT.Map_REQ_IF_CONTENTDBID_REQ_IF_CONTENTDB[req_if_contentDB.ID] = req_if_contentDB
		BackRepoREQ_IF_CONTENTid_atBckpTime_newID[req_if_contentDB_ID_atBackupTime] = req_if_contentDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "REQ_IF_CONTENTDB.json" in dirPath that stores an array
// of REQ_IF_CONTENTDB and stores it in the database
// the map BackRepoREQ_IF_CONTENTid_atBckpTime_newID is updated accordingly
func (backRepoREQ_IF_CONTENT *BackRepoREQ_IF_CONTENTStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoREQ_IF_CONTENTid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "REQ_IF_CONTENTDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json REQ_IF_CONTENT file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*REQ_IF_CONTENTDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_REQ_IF_CONTENTDBID_REQ_IF_CONTENTDB
	for _, req_if_contentDB := range forRestore {

		req_if_contentDB_ID_atBackupTime := req_if_contentDB.ID
		req_if_contentDB.ID = 0
		query := backRepoREQ_IF_CONTENT.db.Create(req_if_contentDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoREQ_IF_CONTENT.Map_REQ_IF_CONTENTDBID_REQ_IF_CONTENTDB[req_if_contentDB.ID] = req_if_contentDB
		BackRepoREQ_IF_CONTENTid_atBckpTime_newID[req_if_contentDB_ID_atBackupTime] = req_if_contentDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json REQ_IF_CONTENT file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<REQ_IF_CONTENT>id_atBckpTime_newID
// to compute new index
func (backRepoREQ_IF_CONTENT *BackRepoREQ_IF_CONTENTStruct) RestorePhaseTwo() {

	for _, req_if_contentDB := range backRepoREQ_IF_CONTENT.Map_REQ_IF_CONTENTDBID_REQ_IF_CONTENTDB {

		// next line of code is to avert unused variable compilation error
		_ = req_if_contentDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		query := backRepoREQ_IF_CONTENT.db.Model(req_if_contentDB).Updates(*req_if_contentDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
	}

}

// BackRepoREQ_IF_CONTENT.ResetReversePointers commits all staged instances of REQ_IF_CONTENT to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoREQ_IF_CONTENT *BackRepoREQ_IF_CONTENTStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, req_if_content := range backRepoREQ_IF_CONTENT.Map_REQ_IF_CONTENTDBID_REQ_IF_CONTENTPtr {
		backRepoREQ_IF_CONTENT.ResetReversePointersInstance(backRepo, idx, req_if_content)
	}

	return
}

func (backRepoREQ_IF_CONTENT *BackRepoREQ_IF_CONTENTStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, req_if_content *models.REQ_IF_CONTENT) (Error error) {

	// fetch matching req_if_contentDB
	if req_if_contentDB, ok := backRepoREQ_IF_CONTENT.Map_REQ_IF_CONTENTDBID_REQ_IF_CONTENTDB[idx]; ok {
		_ = req_if_contentDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoREQ_IF_CONTENTid_atBckpTime_newID map[uint]uint