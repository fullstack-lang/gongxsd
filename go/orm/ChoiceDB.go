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

	"github.com/fullstack-lang/gongxsd/go/models"
)

// dummy variable to have the import declaration wihthout compile failure (even if no code needing this import is generated)
var dummy_Choice_sql sql.NullBool
var dummy_Choice_time time.Duration
var dummy_Choice_sort sort.Float64Slice

// ChoiceAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model choiceAPI
type ChoiceAPI struct {
	gorm.Model

	models.Choice_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	ChoicePointersEncoding ChoicePointersEncoding
}

// ChoicePointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type ChoicePointersEncoding struct {
	// insertion for pointer fields encoding declaration

	// field Annotation is a pointer to another Struct (optional or 0..1)
	// This field is generated into another field to enable AS ONE association
	AnnotationID sql.NullInt64

	// field Elements is a slice of pointers to another Struct (optional or 0..1)
	Elements IntSlice `gorm:"type:TEXT"`

	// field Groups is a slice of pointers to another Struct (optional or 0..1)
	Groups IntSlice `gorm:"type:TEXT"`
}

// ChoiceDB describes a choice in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model choiceDB
type ChoiceDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field choiceDB.Name
	Name_Data sql.NullString

	// Declation for basic field choiceDB.MinOccurs
	MinOccurs_Data sql.NullString

	// Declation for basic field choiceDB.MaxOccurs
	MaxOccurs_Data sql.NullString
	
	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	ChoicePointersEncoding
}

// ChoiceDBs arrays choiceDBs
// swagger:response choiceDBsResponse
type ChoiceDBs []ChoiceDB

// ChoiceDBResponse provides response
// swagger:response choiceDBResponse
type ChoiceDBResponse struct {
	ChoiceDB
}

// ChoiceWOP is a Choice without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type ChoiceWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	MinOccurs string `xlsx:"2"`

	MaxOccurs string `xlsx:"3"`
	// insertion for WOP pointer fields
}

var Choice_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"MinOccurs",
	"MaxOccurs",
}

type BackRepoChoiceStruct struct {
	// stores ChoiceDB according to their gorm ID
	Map_ChoiceDBID_ChoiceDB map[uint]*ChoiceDB

	// stores ChoiceDB ID according to Choice address
	Map_ChoicePtr_ChoiceDBID map[*models.Choice]uint

	// stores Choice according to their gorm ID
	Map_ChoiceDBID_ChoicePtr map[uint]*models.Choice

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoChoice *BackRepoChoiceStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoChoice.stage
	return
}

func (backRepoChoice *BackRepoChoiceStruct) GetDB() *gorm.DB {
	return backRepoChoice.db
}

// GetChoiceDBFromChoicePtr is a handy function to access the back repo instance from the stage instance
func (backRepoChoice *BackRepoChoiceStruct) GetChoiceDBFromChoicePtr(choice *models.Choice) (choiceDB *ChoiceDB) {
	id := backRepoChoice.Map_ChoicePtr_ChoiceDBID[choice]
	choiceDB = backRepoChoice.Map_ChoiceDBID_ChoiceDB[id]
	return
}

// BackRepoChoice.CommitPhaseOne commits all staged instances of Choice to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoChoice *BackRepoChoiceStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for choice := range stage.Choices {
		backRepoChoice.CommitPhaseOneInstance(choice)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, choice := range backRepoChoice.Map_ChoiceDBID_ChoicePtr {
		if _, ok := stage.Choices[choice]; !ok {
			backRepoChoice.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoChoice.CommitDeleteInstance commits deletion of Choice to the BackRepo
func (backRepoChoice *BackRepoChoiceStruct) CommitDeleteInstance(id uint) (Error error) {

	choice := backRepoChoice.Map_ChoiceDBID_ChoicePtr[id]

	// choice is not staged anymore, remove choiceDB
	choiceDB := backRepoChoice.Map_ChoiceDBID_ChoiceDB[id]
	query := backRepoChoice.db.Unscoped().Delete(&choiceDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	delete(backRepoChoice.Map_ChoicePtr_ChoiceDBID, choice)
	delete(backRepoChoice.Map_ChoiceDBID_ChoicePtr, id)
	delete(backRepoChoice.Map_ChoiceDBID_ChoiceDB, id)

	return
}

// BackRepoChoice.CommitPhaseOneInstance commits choice staged instances of Choice to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoChoice *BackRepoChoiceStruct) CommitPhaseOneInstance(choice *models.Choice) (Error error) {

	// check if the choice is not commited yet
	if _, ok := backRepoChoice.Map_ChoicePtr_ChoiceDBID[choice]; ok {
		return
	}

	// initiate choice
	var choiceDB ChoiceDB
	choiceDB.CopyBasicFieldsFromChoice(choice)

	query := backRepoChoice.db.Create(&choiceDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	backRepoChoice.Map_ChoicePtr_ChoiceDBID[choice] = choiceDB.ID
	backRepoChoice.Map_ChoiceDBID_ChoicePtr[choiceDB.ID] = choice
	backRepoChoice.Map_ChoiceDBID_ChoiceDB[choiceDB.ID] = &choiceDB

	return
}

// BackRepoChoice.CommitPhaseTwo commits all staged instances of Choice to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoChoice *BackRepoChoiceStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, choice := range backRepoChoice.Map_ChoiceDBID_ChoicePtr {
		backRepoChoice.CommitPhaseTwoInstance(backRepo, idx, choice)
	}

	return
}

// BackRepoChoice.CommitPhaseTwoInstance commits {{structname }} of models.Choice to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoChoice *BackRepoChoiceStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, choice *models.Choice) (Error error) {

	// fetch matching choiceDB
	if choiceDB, ok := backRepoChoice.Map_ChoiceDBID_ChoiceDB[idx]; ok {

		choiceDB.CopyBasicFieldsFromChoice(choice)

		// insertion point for translating pointers encodings into actual pointers
		// commit pointer value choice.Annotation translates to updating the choice.AnnotationID
		choiceDB.AnnotationID.Valid = true // allow for a 0 value (nil association)
		if choice.Annotation != nil {
			if AnnotationId, ok := backRepo.BackRepoAnnotation.Map_AnnotationPtr_AnnotationDBID[choice.Annotation]; ok {
				choiceDB.AnnotationID.Int64 = int64(AnnotationId)
				choiceDB.AnnotationID.Valid = true
			}
		} else {
			choiceDB.AnnotationID.Int64 = 0
			choiceDB.AnnotationID.Valid = true
		}

		// 1. reset
		choiceDB.ChoicePointersEncoding.Elements = make([]int, 0)
		// 2. encode
		for _, elementAssocEnd := range choice.Elements {
			elementAssocEnd_DB :=
				backRepo.BackRepoElement.GetElementDBFromElementPtr(elementAssocEnd)
			
			// the stage might be inconsistant, meaning that the elementAssocEnd_DB might
			// be missing from the stage. In this case, the commit operation is robust
			// An alternative would be to crash here to reveal the missing element.
			if elementAssocEnd_DB == nil {
				continue
			}
			
			choiceDB.ChoicePointersEncoding.Elements =
				append(choiceDB.ChoicePointersEncoding.Elements, int(elementAssocEnd_DB.ID))
		}

		// 1. reset
		choiceDB.ChoicePointersEncoding.Groups = make([]int, 0)
		// 2. encode
		for _, groupAssocEnd := range choice.Groups {
			groupAssocEnd_DB :=
				backRepo.BackRepoGroup.GetGroupDBFromGroupPtr(groupAssocEnd)
			
			// the stage might be inconsistant, meaning that the groupAssocEnd_DB might
			// be missing from the stage. In this case, the commit operation is robust
			// An alternative would be to crash here to reveal the missing element.
			if groupAssocEnd_DB == nil {
				continue
			}
			
			choiceDB.ChoicePointersEncoding.Groups =
				append(choiceDB.ChoicePointersEncoding.Groups, int(groupAssocEnd_DB.ID))
		}

		query := backRepoChoice.db.Save(&choiceDB)
		if query.Error != nil {
			log.Fatalln(query.Error)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown Choice intance %s", choice.Name))
		return err
	}

	return
}

// BackRepoChoice.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoChoice *BackRepoChoiceStruct) CheckoutPhaseOne() (Error error) {

	choiceDBArray := make([]ChoiceDB, 0)
	query := backRepoChoice.db.Find(&choiceDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	choiceInstancesToBeRemovedFromTheStage := make(map[*models.Choice]any)
	for key, value := range backRepoChoice.stage.Choices {
		choiceInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, choiceDB := range choiceDBArray {
		backRepoChoice.CheckoutPhaseOneInstance(&choiceDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		choice, ok := backRepoChoice.Map_ChoiceDBID_ChoicePtr[choiceDB.ID]
		if ok {
			delete(choiceInstancesToBeRemovedFromTheStage, choice)
		}
	}

	// remove from stage and back repo's 3 maps all choices that are not in the checkout
	for choice := range choiceInstancesToBeRemovedFromTheStage {
		choice.Unstage(backRepoChoice.GetStage())

		// remove instance from the back repo 3 maps
		choiceID := backRepoChoice.Map_ChoicePtr_ChoiceDBID[choice]
		delete(backRepoChoice.Map_ChoicePtr_ChoiceDBID, choice)
		delete(backRepoChoice.Map_ChoiceDBID_ChoiceDB, choiceID)
		delete(backRepoChoice.Map_ChoiceDBID_ChoicePtr, choiceID)
	}

	return
}

// CheckoutPhaseOneInstance takes a choiceDB that has been found in the DB, updates the backRepo and stages the
// models version of the choiceDB
func (backRepoChoice *BackRepoChoiceStruct) CheckoutPhaseOneInstance(choiceDB *ChoiceDB) (Error error) {

	choice, ok := backRepoChoice.Map_ChoiceDBID_ChoicePtr[choiceDB.ID]
	if !ok {
		choice = new(models.Choice)

		backRepoChoice.Map_ChoiceDBID_ChoicePtr[choiceDB.ID] = choice
		backRepoChoice.Map_ChoicePtr_ChoiceDBID[choice] = choiceDB.ID

		// append model store with the new element
		choice.Name = choiceDB.Name_Data.String
		choice.Stage(backRepoChoice.GetStage())
	}
	choiceDB.CopyBasicFieldsToChoice(choice)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	choice.Stage(backRepoChoice.GetStage())

	// preserve pointer to choiceDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_ChoiceDBID_ChoiceDB)[choiceDB hold variable pointers
	choiceDB_Data := *choiceDB
	preservedPtrToChoice := &choiceDB_Data
	backRepoChoice.Map_ChoiceDBID_ChoiceDB[choiceDB.ID] = preservedPtrToChoice

	return
}

// BackRepoChoice.CheckoutPhaseTwo Checkouts all staged instances of Choice to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoChoice *BackRepoChoiceStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, choiceDB := range backRepoChoice.Map_ChoiceDBID_ChoiceDB {
		backRepoChoice.CheckoutPhaseTwoInstance(backRepo, choiceDB)
	}
	return
}

// BackRepoChoice.CheckoutPhaseTwoInstance Checkouts staged instances of Choice to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoChoice *BackRepoChoiceStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, choiceDB *ChoiceDB) (Error error) {

	choice := backRepoChoice.Map_ChoiceDBID_ChoicePtr[choiceDB.ID]

	choiceDB.DecodePointers(backRepo, choice)

	return
}

func (choiceDB *ChoiceDB) DecodePointers(backRepo *BackRepoStruct, choice *models.Choice) {

	// insertion point for checkout of pointer encoding
	// Annotation field
	choice.Annotation = nil
	if choiceDB.AnnotationID.Int64 != 0 {
		choice.Annotation = backRepo.BackRepoAnnotation.Map_AnnotationDBID_AnnotationPtr[uint(choiceDB.AnnotationID.Int64)]
	}
	// This loop redeem choice.Elements in the stage from the encode in the back repo
	// It parses all ElementDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	choice.Elements = choice.Elements[:0]
	for _, _Elementid := range choiceDB.ChoicePointersEncoding.Elements {
		choice.Elements = append(choice.Elements, backRepo.BackRepoElement.Map_ElementDBID_ElementPtr[uint(_Elementid)])
	}

	// This loop redeem choice.Groups in the stage from the encode in the back repo
	// It parses all GroupDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	choice.Groups = choice.Groups[:0]
	for _, _Groupid := range choiceDB.ChoicePointersEncoding.Groups {
		choice.Groups = append(choice.Groups, backRepo.BackRepoGroup.Map_GroupDBID_GroupPtr[uint(_Groupid)])
	}

	return
}

// CommitChoice allows commit of a single choice (if already staged)
func (backRepo *BackRepoStruct) CommitChoice(choice *models.Choice) {
	backRepo.BackRepoChoice.CommitPhaseOneInstance(choice)
	if id, ok := backRepo.BackRepoChoice.Map_ChoicePtr_ChoiceDBID[choice]; ok {
		backRepo.BackRepoChoice.CommitPhaseTwoInstance(backRepo, id, choice)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitChoice allows checkout of a single choice (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutChoice(choice *models.Choice) {
	// check if the choice is staged
	if _, ok := backRepo.BackRepoChoice.Map_ChoicePtr_ChoiceDBID[choice]; ok {

		if id, ok := backRepo.BackRepoChoice.Map_ChoicePtr_ChoiceDBID[choice]; ok {
			var choiceDB ChoiceDB
			choiceDB.ID = id

			if err := backRepo.BackRepoChoice.db.First(&choiceDB, id).Error; err != nil {
				log.Fatalln("CheckoutChoice : Problem with getting object with id:", id)
			}
			backRepo.BackRepoChoice.CheckoutPhaseOneInstance(&choiceDB)
			backRepo.BackRepoChoice.CheckoutPhaseTwoInstance(backRepo, &choiceDB)
		}
	}
}

// CopyBasicFieldsFromChoice
func (choiceDB *ChoiceDB) CopyBasicFieldsFromChoice(choice *models.Choice) {
	// insertion point for fields commit

	choiceDB.Name_Data.String = choice.Name
	choiceDB.Name_Data.Valid = true

	choiceDB.MinOccurs_Data.String = choice.MinOccurs
	choiceDB.MinOccurs_Data.Valid = true

	choiceDB.MaxOccurs_Data.String = choice.MaxOccurs
	choiceDB.MaxOccurs_Data.Valid = true
}

// CopyBasicFieldsFromChoice_WOP
func (choiceDB *ChoiceDB) CopyBasicFieldsFromChoice_WOP(choice *models.Choice_WOP) {
	// insertion point for fields commit

	choiceDB.Name_Data.String = choice.Name
	choiceDB.Name_Data.Valid = true

	choiceDB.MinOccurs_Data.String = choice.MinOccurs
	choiceDB.MinOccurs_Data.Valid = true

	choiceDB.MaxOccurs_Data.String = choice.MaxOccurs
	choiceDB.MaxOccurs_Data.Valid = true
}

// CopyBasicFieldsFromChoiceWOP
func (choiceDB *ChoiceDB) CopyBasicFieldsFromChoiceWOP(choice *ChoiceWOP) {
	// insertion point for fields commit

	choiceDB.Name_Data.String = choice.Name
	choiceDB.Name_Data.Valid = true

	choiceDB.MinOccurs_Data.String = choice.MinOccurs
	choiceDB.MinOccurs_Data.Valid = true

	choiceDB.MaxOccurs_Data.String = choice.MaxOccurs
	choiceDB.MaxOccurs_Data.Valid = true
}

// CopyBasicFieldsToChoice
func (choiceDB *ChoiceDB) CopyBasicFieldsToChoice(choice *models.Choice) {
	// insertion point for checkout of basic fields (back repo to stage)
	choice.Name = choiceDB.Name_Data.String
	choice.MinOccurs = choiceDB.MinOccurs_Data.String
	choice.MaxOccurs = choiceDB.MaxOccurs_Data.String
}

// CopyBasicFieldsToChoice_WOP
func (choiceDB *ChoiceDB) CopyBasicFieldsToChoice_WOP(choice *models.Choice_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	choice.Name = choiceDB.Name_Data.String
	choice.MinOccurs = choiceDB.MinOccurs_Data.String
	choice.MaxOccurs = choiceDB.MaxOccurs_Data.String
}

// CopyBasicFieldsToChoiceWOP
func (choiceDB *ChoiceDB) CopyBasicFieldsToChoiceWOP(choice *ChoiceWOP) {
	choice.ID = int(choiceDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	choice.Name = choiceDB.Name_Data.String
	choice.MinOccurs = choiceDB.MinOccurs_Data.String
	choice.MaxOccurs = choiceDB.MaxOccurs_Data.String
}

// Backup generates a json file from a slice of all ChoiceDB instances in the backrepo
func (backRepoChoice *BackRepoChoiceStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "ChoiceDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*ChoiceDB, 0)
	for _, choiceDB := range backRepoChoice.Map_ChoiceDBID_ChoiceDB {
		forBackup = append(forBackup, choiceDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json Choice ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json Choice file", err.Error())
	}
}

// Backup generates a json file from a slice of all ChoiceDB instances in the backrepo
func (backRepoChoice *BackRepoChoiceStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*ChoiceDB, 0)
	for _, choiceDB := range backRepoChoice.Map_ChoiceDBID_ChoiceDB {
		forBackup = append(forBackup, choiceDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("Choice")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&Choice_Fields, -1)
	for _, choiceDB := range forBackup {

		var choiceWOP ChoiceWOP
		choiceDB.CopyBasicFieldsToChoiceWOP(&choiceWOP)

		row := sh.AddRow()
		row.WriteStruct(&choiceWOP, -1)
	}
}

// RestoreXL from the "Choice" sheet all ChoiceDB instances
func (backRepoChoice *BackRepoChoiceStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoChoiceid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["Choice"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoChoice.rowVisitorChoice)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoChoice *BackRepoChoiceStruct) rowVisitorChoice(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var choiceWOP ChoiceWOP
		row.ReadStruct(&choiceWOP)

		// add the unmarshalled struct to the stage
		choiceDB := new(ChoiceDB)
		choiceDB.CopyBasicFieldsFromChoiceWOP(&choiceWOP)

		choiceDB_ID_atBackupTime := choiceDB.ID
		choiceDB.ID = 0
		query := backRepoChoice.db.Create(choiceDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoChoice.Map_ChoiceDBID_ChoiceDB[choiceDB.ID] = choiceDB
		BackRepoChoiceid_atBckpTime_newID[choiceDB_ID_atBackupTime] = choiceDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "ChoiceDB.json" in dirPath that stores an array
// of ChoiceDB and stores it in the database
// the map BackRepoChoiceid_atBckpTime_newID is updated accordingly
func (backRepoChoice *BackRepoChoiceStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoChoiceid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "ChoiceDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json Choice file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*ChoiceDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_ChoiceDBID_ChoiceDB
	for _, choiceDB := range forRestore {

		choiceDB_ID_atBackupTime := choiceDB.ID
		choiceDB.ID = 0
		query := backRepoChoice.db.Create(choiceDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoChoice.Map_ChoiceDBID_ChoiceDB[choiceDB.ID] = choiceDB
		BackRepoChoiceid_atBckpTime_newID[choiceDB_ID_atBackupTime] = choiceDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json Choice file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<Choice>id_atBckpTime_newID
// to compute new index
func (backRepoChoice *BackRepoChoiceStruct) RestorePhaseTwo() {

	for _, choiceDB := range backRepoChoice.Map_ChoiceDBID_ChoiceDB {

		// next line of code is to avert unused variable compilation error
		_ = choiceDB

		// insertion point for reindexing pointers encoding
		// reindexing Annotation field
		if choiceDB.AnnotationID.Int64 != 0 {
			choiceDB.AnnotationID.Int64 = int64(BackRepoAnnotationid_atBckpTime_newID[uint(choiceDB.AnnotationID.Int64)])
			choiceDB.AnnotationID.Valid = true
		}

		// update databse with new index encoding
		query := backRepoChoice.db.Model(choiceDB).Updates(*choiceDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
	}

}

// BackRepoChoice.ResetReversePointers commits all staged instances of Choice to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoChoice *BackRepoChoiceStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, choice := range backRepoChoice.Map_ChoiceDBID_ChoicePtr {
		backRepoChoice.ResetReversePointersInstance(backRepo, idx, choice)
	}

	return
}

func (backRepoChoice *BackRepoChoiceStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, choice *models.Choice) (Error error) {

	// fetch matching choiceDB
	if choiceDB, ok := backRepoChoice.Map_ChoiceDBID_ChoiceDB[idx]; ok {
		_ = choiceDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoChoiceid_atBckpTime_newID map[uint]uint