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

	"github.com/fullstack-lang/gong/go/db"
	"github.com/fullstack-lang/gong/go/models"
)

// dummy variable to have the import declaration wihthout compile failure (even if no code needing this import is generated)
var dummy_GongStruct_sql sql.NullBool
var dummy_GongStruct_time time.Duration
var dummy_GongStruct_sort sort.Float64Slice

// GongStructAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model gongstructAPI
type GongStructAPI struct {
	gorm.Model

	models.GongStruct_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	GongStructPointersEncoding GongStructPointersEncoding
}

// GongStructPointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type GongStructPointersEncoding struct {
	// insertion for pointer fields encoding declaration

	// field GongBasicFields is a slice of pointers to another Struct (optional or 0..1)
	GongBasicFields IntSlice `gorm:"type:TEXT"`

	// field GongTimeFields is a slice of pointers to another Struct (optional or 0..1)
	GongTimeFields IntSlice `gorm:"type:TEXT"`

	// field PointerToGongStructFields is a slice of pointers to another Struct (optional or 0..1)
	PointerToGongStructFields IntSlice `gorm:"type:TEXT"`

	// field SliceOfPointerToGongStructFields is a slice of pointers to another Struct (optional or 0..1)
	SliceOfPointerToGongStructFields IntSlice `gorm:"type:TEXT"`
}

// GongStructDB describes a gongstruct in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model gongstructDB
type GongStructDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field gongstructDB.Name
	Name_Data sql.NullString

	// Declation for basic field gongstructDB.HasOnAfterUpdateSignature
	// provide the sql storage for the boolan
	HasOnAfterUpdateSignature_Data sql.NullBool

	// Declation for basic field gongstructDB.IsIgnoredForFront
	// provide the sql storage for the boolan
	IsIgnoredForFront_Data sql.NullBool

	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	GongStructPointersEncoding
}

// GongStructDBs arrays gongstructDBs
// swagger:response gongstructDBsResponse
type GongStructDBs []GongStructDB

// GongStructDBResponse provides response
// swagger:response gongstructDBResponse
type GongStructDBResponse struct {
	GongStructDB
}

// GongStructWOP is a GongStruct without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type GongStructWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	HasOnAfterUpdateSignature bool `xlsx:"2"`

	IsIgnoredForFront bool `xlsx:"3"`
	// insertion for WOP pointer fields
}

var GongStruct_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"HasOnAfterUpdateSignature",
	"IsIgnoredForFront",
}

type BackRepoGongStructStruct struct {
	// stores GongStructDB according to their gorm ID
	Map_GongStructDBID_GongStructDB map[uint]*GongStructDB

	// stores GongStructDB ID according to GongStruct address
	Map_GongStructPtr_GongStructDBID map[*models.GongStruct]uint

	// stores GongStruct according to their gorm ID
	Map_GongStructDBID_GongStructPtr map[uint]*models.GongStruct

	db db.DBInterface

	stage *models.Stage
}

func (backRepoGongStruct *BackRepoGongStructStruct) GetStage() (stage *models.Stage) {
	stage = backRepoGongStruct.stage
	return
}

func (backRepoGongStruct *BackRepoGongStructStruct) GetDB() db.DBInterface {
	return backRepoGongStruct.db
}

// GetGongStructDBFromGongStructPtr is a handy function to access the back repo instance from the stage instance
func (backRepoGongStruct *BackRepoGongStructStruct) GetGongStructDBFromGongStructPtr(gongstruct *models.GongStruct) (gongstructDB *GongStructDB) {
	id := backRepoGongStruct.Map_GongStructPtr_GongStructDBID[gongstruct]
	gongstructDB = backRepoGongStruct.Map_GongStructDBID_GongStructDB[id]
	return
}

// BackRepoGongStruct.CommitPhaseOne commits all staged instances of GongStruct to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoGongStruct *BackRepoGongStructStruct) CommitPhaseOne(stage *models.Stage) (Error error) {

	var gongstructs []*models.GongStruct
	for gongstruct := range stage.GongStructs {
		gongstructs = append(gongstructs, gongstruct)
	}

	// Sort by the order stored in Map_Staged_Order.
	sort.Slice(gongstructs, func(i, j int) bool {
		return stage.GongStructMap_Staged_Order[gongstructs[i]] < stage.GongStructMap_Staged_Order[gongstructs[j]]
	})

	for _, gongstruct := range gongstructs {
		backRepoGongStruct.CommitPhaseOneInstance(gongstruct)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, gongstruct := range backRepoGongStruct.Map_GongStructDBID_GongStructPtr {
		if _, ok := stage.GongStructs[gongstruct]; !ok {
			backRepoGongStruct.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoGongStruct.CommitDeleteInstance commits deletion of GongStruct to the BackRepo
func (backRepoGongStruct *BackRepoGongStructStruct) CommitDeleteInstance(id uint) (Error error) {

	gongstruct := backRepoGongStruct.Map_GongStructDBID_GongStructPtr[id]

	// gongstruct is not staged anymore, remove gongstructDB
	gongstructDB := backRepoGongStruct.Map_GongStructDBID_GongStructDB[id]
	db, _ := backRepoGongStruct.db.Unscoped()
	_, err := db.Delete(gongstructDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	delete(backRepoGongStruct.Map_GongStructPtr_GongStructDBID, gongstruct)
	delete(backRepoGongStruct.Map_GongStructDBID_GongStructPtr, id)
	delete(backRepoGongStruct.Map_GongStructDBID_GongStructDB, id)

	return
}

// BackRepoGongStruct.CommitPhaseOneInstance commits gongstruct staged instances of GongStruct to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoGongStruct *BackRepoGongStructStruct) CommitPhaseOneInstance(gongstruct *models.GongStruct) (Error error) {

	// check if the gongstruct is not commited yet
	if _, ok := backRepoGongStruct.Map_GongStructPtr_GongStructDBID[gongstruct]; ok {
		return
	}

	// initiate gongstruct
	var gongstructDB GongStructDB
	gongstructDB.CopyBasicFieldsFromGongStruct(gongstruct)

	_, err := backRepoGongStruct.db.Create(&gongstructDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	backRepoGongStruct.Map_GongStructPtr_GongStructDBID[gongstruct] = gongstructDB.ID
	backRepoGongStruct.Map_GongStructDBID_GongStructPtr[gongstructDB.ID] = gongstruct
	backRepoGongStruct.Map_GongStructDBID_GongStructDB[gongstructDB.ID] = &gongstructDB

	return
}

// BackRepoGongStruct.CommitPhaseTwo commits all staged instances of GongStruct to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoGongStruct *BackRepoGongStructStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, gongstruct := range backRepoGongStruct.Map_GongStructDBID_GongStructPtr {
		backRepoGongStruct.CommitPhaseTwoInstance(backRepo, idx, gongstruct)
	}

	return
}

// BackRepoGongStruct.CommitPhaseTwoInstance commits {{structname }} of models.GongStruct to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoGongStruct *BackRepoGongStructStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, gongstruct *models.GongStruct) (Error error) {

	// fetch matching gongstructDB
	if gongstructDB, ok := backRepoGongStruct.Map_GongStructDBID_GongStructDB[idx]; ok {

		gongstructDB.CopyBasicFieldsFromGongStruct(gongstruct)

		// insertion point for translating pointers encodings into actual pointers
		// 1. reset
		gongstructDB.GongStructPointersEncoding.GongBasicFields = make([]int, 0)
		// 2. encode
		for _, gongbasicfieldAssocEnd := range gongstruct.GongBasicFields {
			gongbasicfieldAssocEnd_DB :=
				backRepo.BackRepoGongBasicField.GetGongBasicFieldDBFromGongBasicFieldPtr(gongbasicfieldAssocEnd)
			
			// the stage might be inconsistant, meaning that the gongbasicfieldAssocEnd_DB might
			// be missing from the stage. In this case, the commit operation is robust
			// An alternative would be to crash here to reveal the missing element.
			if gongbasicfieldAssocEnd_DB == nil {
				continue
			}
			
			gongstructDB.GongStructPointersEncoding.GongBasicFields =
				append(gongstructDB.GongStructPointersEncoding.GongBasicFields, int(gongbasicfieldAssocEnd_DB.ID))
		}

		// 1. reset
		gongstructDB.GongStructPointersEncoding.GongTimeFields = make([]int, 0)
		// 2. encode
		for _, gongtimefieldAssocEnd := range gongstruct.GongTimeFields {
			gongtimefieldAssocEnd_DB :=
				backRepo.BackRepoGongTimeField.GetGongTimeFieldDBFromGongTimeFieldPtr(gongtimefieldAssocEnd)
			
			// the stage might be inconsistant, meaning that the gongtimefieldAssocEnd_DB might
			// be missing from the stage. In this case, the commit operation is robust
			// An alternative would be to crash here to reveal the missing element.
			if gongtimefieldAssocEnd_DB == nil {
				continue
			}
			
			gongstructDB.GongStructPointersEncoding.GongTimeFields =
				append(gongstructDB.GongStructPointersEncoding.GongTimeFields, int(gongtimefieldAssocEnd_DB.ID))
		}

		// 1. reset
		gongstructDB.GongStructPointersEncoding.PointerToGongStructFields = make([]int, 0)
		// 2. encode
		for _, pointertogongstructfieldAssocEnd := range gongstruct.PointerToGongStructFields {
			pointertogongstructfieldAssocEnd_DB :=
				backRepo.BackRepoPointerToGongStructField.GetPointerToGongStructFieldDBFromPointerToGongStructFieldPtr(pointertogongstructfieldAssocEnd)
			
			// the stage might be inconsistant, meaning that the pointertogongstructfieldAssocEnd_DB might
			// be missing from the stage. In this case, the commit operation is robust
			// An alternative would be to crash here to reveal the missing element.
			if pointertogongstructfieldAssocEnd_DB == nil {
				continue
			}
			
			gongstructDB.GongStructPointersEncoding.PointerToGongStructFields =
				append(gongstructDB.GongStructPointersEncoding.PointerToGongStructFields, int(pointertogongstructfieldAssocEnd_DB.ID))
		}

		// 1. reset
		gongstructDB.GongStructPointersEncoding.SliceOfPointerToGongStructFields = make([]int, 0)
		// 2. encode
		for _, sliceofpointertogongstructfieldAssocEnd := range gongstruct.SliceOfPointerToGongStructFields {
			sliceofpointertogongstructfieldAssocEnd_DB :=
				backRepo.BackRepoSliceOfPointerToGongStructField.GetSliceOfPointerToGongStructFieldDBFromSliceOfPointerToGongStructFieldPtr(sliceofpointertogongstructfieldAssocEnd)
			
			// the stage might be inconsistant, meaning that the sliceofpointertogongstructfieldAssocEnd_DB might
			// be missing from the stage. In this case, the commit operation is robust
			// An alternative would be to crash here to reveal the missing element.
			if sliceofpointertogongstructfieldAssocEnd_DB == nil {
				continue
			}
			
			gongstructDB.GongStructPointersEncoding.SliceOfPointerToGongStructFields =
				append(gongstructDB.GongStructPointersEncoding.SliceOfPointerToGongStructFields, int(sliceofpointertogongstructfieldAssocEnd_DB.ID))
		}

		_, err := backRepoGongStruct.db.Save(gongstructDB)
		if err != nil {
			log.Fatal(err)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown GongStruct intance %s", gongstruct.Name))
		return err
	}

	return
}

// BackRepoGongStruct.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoGongStruct *BackRepoGongStructStruct) CheckoutPhaseOne() (Error error) {

	gongstructDBArray := make([]GongStructDB, 0)
	_, err := backRepoGongStruct.db.Find(&gongstructDBArray)
	if err != nil {
		return err
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	gongstructInstancesToBeRemovedFromTheStage := make(map[*models.GongStruct]any)
	for key, value := range backRepoGongStruct.stage.GongStructs {
		gongstructInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, gongstructDB := range gongstructDBArray {
		backRepoGongStruct.CheckoutPhaseOneInstance(&gongstructDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		gongstruct, ok := backRepoGongStruct.Map_GongStructDBID_GongStructPtr[gongstructDB.ID]
		if ok {
			delete(gongstructInstancesToBeRemovedFromTheStage, gongstruct)
		}
	}

	// remove from stage and back repo's 3 maps all gongstructs that are not in the checkout
	for gongstruct := range gongstructInstancesToBeRemovedFromTheStage {
		gongstruct.Unstage(backRepoGongStruct.GetStage())

		// remove instance from the back repo 3 maps
		gongstructID := backRepoGongStruct.Map_GongStructPtr_GongStructDBID[gongstruct]
		delete(backRepoGongStruct.Map_GongStructPtr_GongStructDBID, gongstruct)
		delete(backRepoGongStruct.Map_GongStructDBID_GongStructDB, gongstructID)
		delete(backRepoGongStruct.Map_GongStructDBID_GongStructPtr, gongstructID)
	}

	return
}

// CheckoutPhaseOneInstance takes a gongstructDB that has been found in the DB, updates the backRepo and stages the
// models version of the gongstructDB
func (backRepoGongStruct *BackRepoGongStructStruct) CheckoutPhaseOneInstance(gongstructDB *GongStructDB) (Error error) {

	gongstruct, ok := backRepoGongStruct.Map_GongStructDBID_GongStructPtr[gongstructDB.ID]
	if !ok {
		gongstruct = new(models.GongStruct)

		backRepoGongStruct.Map_GongStructDBID_GongStructPtr[gongstructDB.ID] = gongstruct
		backRepoGongStruct.Map_GongStructPtr_GongStructDBID[gongstruct] = gongstructDB.ID

		// append model store with the new element
		gongstruct.Name = gongstructDB.Name_Data.String
		gongstruct.Stage(backRepoGongStruct.GetStage())
	}
	gongstructDB.CopyBasicFieldsToGongStruct(gongstruct)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	gongstruct.Stage(backRepoGongStruct.GetStage())

	// preserve pointer to gongstructDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_GongStructDBID_GongStructDB)[gongstructDB hold variable pointers
	gongstructDB_Data := *gongstructDB
	preservedPtrToGongStruct := &gongstructDB_Data
	backRepoGongStruct.Map_GongStructDBID_GongStructDB[gongstructDB.ID] = preservedPtrToGongStruct

	return
}

// BackRepoGongStruct.CheckoutPhaseTwo Checkouts all staged instances of GongStruct to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoGongStruct *BackRepoGongStructStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, gongstructDB := range backRepoGongStruct.Map_GongStructDBID_GongStructDB {
		backRepoGongStruct.CheckoutPhaseTwoInstance(backRepo, gongstructDB)
	}
	return
}

// BackRepoGongStruct.CheckoutPhaseTwoInstance Checkouts staged instances of GongStruct to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoGongStruct *BackRepoGongStructStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, gongstructDB *GongStructDB) (Error error) {

	gongstruct := backRepoGongStruct.Map_GongStructDBID_GongStructPtr[gongstructDB.ID]

	gongstructDB.DecodePointers(backRepo, gongstruct)

	return
}

func (gongstructDB *GongStructDB) DecodePointers(backRepo *BackRepoStruct, gongstruct *models.GongStruct) {

	// insertion point for checkout of pointer encoding
	// This loop redeem gongstruct.GongBasicFields in the stage from the encode in the back repo
	// It parses all GongBasicFieldDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	gongstruct.GongBasicFields = gongstruct.GongBasicFields[:0]
	for _, _GongBasicFieldid := range gongstructDB.GongStructPointersEncoding.GongBasicFields {
		gongstruct.GongBasicFields = append(gongstruct.GongBasicFields, backRepo.BackRepoGongBasicField.Map_GongBasicFieldDBID_GongBasicFieldPtr[uint(_GongBasicFieldid)])
	}

	// This loop redeem gongstruct.GongTimeFields in the stage from the encode in the back repo
	// It parses all GongTimeFieldDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	gongstruct.GongTimeFields = gongstruct.GongTimeFields[:0]
	for _, _GongTimeFieldid := range gongstructDB.GongStructPointersEncoding.GongTimeFields {
		gongstruct.GongTimeFields = append(gongstruct.GongTimeFields, backRepo.BackRepoGongTimeField.Map_GongTimeFieldDBID_GongTimeFieldPtr[uint(_GongTimeFieldid)])
	}

	// This loop redeem gongstruct.PointerToGongStructFields in the stage from the encode in the back repo
	// It parses all PointerToGongStructFieldDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	gongstruct.PointerToGongStructFields = gongstruct.PointerToGongStructFields[:0]
	for _, _PointerToGongStructFieldid := range gongstructDB.GongStructPointersEncoding.PointerToGongStructFields {
		gongstruct.PointerToGongStructFields = append(gongstruct.PointerToGongStructFields, backRepo.BackRepoPointerToGongStructField.Map_PointerToGongStructFieldDBID_PointerToGongStructFieldPtr[uint(_PointerToGongStructFieldid)])
	}

	// This loop redeem gongstruct.SliceOfPointerToGongStructFields in the stage from the encode in the back repo
	// It parses all SliceOfPointerToGongStructFieldDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	gongstruct.SliceOfPointerToGongStructFields = gongstruct.SliceOfPointerToGongStructFields[:0]
	for _, _SliceOfPointerToGongStructFieldid := range gongstructDB.GongStructPointersEncoding.SliceOfPointerToGongStructFields {
		gongstruct.SliceOfPointerToGongStructFields = append(gongstruct.SliceOfPointerToGongStructFields, backRepo.BackRepoSliceOfPointerToGongStructField.Map_SliceOfPointerToGongStructFieldDBID_SliceOfPointerToGongStructFieldPtr[uint(_SliceOfPointerToGongStructFieldid)])
	}

	return
}

// CommitGongStruct allows commit of a single gongstruct (if already staged)
func (backRepo *BackRepoStruct) CommitGongStruct(gongstruct *models.GongStruct) {
	backRepo.BackRepoGongStruct.CommitPhaseOneInstance(gongstruct)
	if id, ok := backRepo.BackRepoGongStruct.Map_GongStructPtr_GongStructDBID[gongstruct]; ok {
		backRepo.BackRepoGongStruct.CommitPhaseTwoInstance(backRepo, id, gongstruct)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitGongStruct allows checkout of a single gongstruct (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutGongStruct(gongstruct *models.GongStruct) {
	// check if the gongstruct is staged
	if _, ok := backRepo.BackRepoGongStruct.Map_GongStructPtr_GongStructDBID[gongstruct]; ok {

		if id, ok := backRepo.BackRepoGongStruct.Map_GongStructPtr_GongStructDBID[gongstruct]; ok {
			var gongstructDB GongStructDB
			gongstructDB.ID = id

			if _, err := backRepo.BackRepoGongStruct.db.First(&gongstructDB, id); err != nil {
				log.Fatalln("CheckoutGongStruct : Problem with getting object with id:", id)
			}
			backRepo.BackRepoGongStruct.CheckoutPhaseOneInstance(&gongstructDB)
			backRepo.BackRepoGongStruct.CheckoutPhaseTwoInstance(backRepo, &gongstructDB)
		}
	}
}

// CopyBasicFieldsFromGongStruct
func (gongstructDB *GongStructDB) CopyBasicFieldsFromGongStruct(gongstruct *models.GongStruct) {
	// insertion point for fields commit

	gongstructDB.Name_Data.String = gongstruct.Name
	gongstructDB.Name_Data.Valid = true

	gongstructDB.HasOnAfterUpdateSignature_Data.Bool = gongstruct.HasOnAfterUpdateSignature
	gongstructDB.HasOnAfterUpdateSignature_Data.Valid = true

	gongstructDB.IsIgnoredForFront_Data.Bool = gongstruct.IsIgnoredForFront
	gongstructDB.IsIgnoredForFront_Data.Valid = true
}

// CopyBasicFieldsFromGongStruct_WOP
func (gongstructDB *GongStructDB) CopyBasicFieldsFromGongStruct_WOP(gongstruct *models.GongStruct_WOP) {
	// insertion point for fields commit

	gongstructDB.Name_Data.String = gongstruct.Name
	gongstructDB.Name_Data.Valid = true

	gongstructDB.HasOnAfterUpdateSignature_Data.Bool = gongstruct.HasOnAfterUpdateSignature
	gongstructDB.HasOnAfterUpdateSignature_Data.Valid = true

	gongstructDB.IsIgnoredForFront_Data.Bool = gongstruct.IsIgnoredForFront
	gongstructDB.IsIgnoredForFront_Data.Valid = true
}

// CopyBasicFieldsFromGongStructWOP
func (gongstructDB *GongStructDB) CopyBasicFieldsFromGongStructWOP(gongstruct *GongStructWOP) {
	// insertion point for fields commit

	gongstructDB.Name_Data.String = gongstruct.Name
	gongstructDB.Name_Data.Valid = true

	gongstructDB.HasOnAfterUpdateSignature_Data.Bool = gongstruct.HasOnAfterUpdateSignature
	gongstructDB.HasOnAfterUpdateSignature_Data.Valid = true

	gongstructDB.IsIgnoredForFront_Data.Bool = gongstruct.IsIgnoredForFront
	gongstructDB.IsIgnoredForFront_Data.Valid = true
}

// CopyBasicFieldsToGongStruct
func (gongstructDB *GongStructDB) CopyBasicFieldsToGongStruct(gongstruct *models.GongStruct) {
	// insertion point for checkout of basic fields (back repo to stage)
	gongstruct.Name = gongstructDB.Name_Data.String
	gongstruct.HasOnAfterUpdateSignature = gongstructDB.HasOnAfterUpdateSignature_Data.Bool
	gongstruct.IsIgnoredForFront = gongstructDB.IsIgnoredForFront_Data.Bool
}

// CopyBasicFieldsToGongStruct_WOP
func (gongstructDB *GongStructDB) CopyBasicFieldsToGongStruct_WOP(gongstruct *models.GongStruct_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	gongstruct.Name = gongstructDB.Name_Data.String
	gongstruct.HasOnAfterUpdateSignature = gongstructDB.HasOnAfterUpdateSignature_Data.Bool
	gongstruct.IsIgnoredForFront = gongstructDB.IsIgnoredForFront_Data.Bool
}

// CopyBasicFieldsToGongStructWOP
func (gongstructDB *GongStructDB) CopyBasicFieldsToGongStructWOP(gongstruct *GongStructWOP) {
	gongstruct.ID = int(gongstructDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	gongstruct.Name = gongstructDB.Name_Data.String
	gongstruct.HasOnAfterUpdateSignature = gongstructDB.HasOnAfterUpdateSignature_Data.Bool
	gongstruct.IsIgnoredForFront = gongstructDB.IsIgnoredForFront_Data.Bool
}

// Backup generates a json file from a slice of all GongStructDB instances in the backrepo
func (backRepoGongStruct *BackRepoGongStructStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "GongStructDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*GongStructDB, 0)
	for _, gongstructDB := range backRepoGongStruct.Map_GongStructDBID_GongStructDB {
		forBackup = append(forBackup, gongstructDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json GongStruct ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json GongStruct file", err.Error())
	}
}

// Backup generates a json file from a slice of all GongStructDB instances in the backrepo
func (backRepoGongStruct *BackRepoGongStructStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*GongStructDB, 0)
	for _, gongstructDB := range backRepoGongStruct.Map_GongStructDBID_GongStructDB {
		forBackup = append(forBackup, gongstructDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("GongStruct")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&GongStruct_Fields, -1)
	for _, gongstructDB := range forBackup {

		var gongstructWOP GongStructWOP
		gongstructDB.CopyBasicFieldsToGongStructWOP(&gongstructWOP)

		row := sh.AddRow()
		row.WriteStruct(&gongstructWOP, -1)
	}
}

// RestoreXL from the "GongStruct" sheet all GongStructDB instances
func (backRepoGongStruct *BackRepoGongStructStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoGongStructid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["GongStruct"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoGongStruct.rowVisitorGongStruct)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoGongStruct *BackRepoGongStructStruct) rowVisitorGongStruct(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var gongstructWOP GongStructWOP
		row.ReadStruct(&gongstructWOP)

		// add the unmarshalled struct to the stage
		gongstructDB := new(GongStructDB)
		gongstructDB.CopyBasicFieldsFromGongStructWOP(&gongstructWOP)

		gongstructDB_ID_atBackupTime := gongstructDB.ID
		gongstructDB.ID = 0
		_, err := backRepoGongStruct.db.Create(gongstructDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoGongStruct.Map_GongStructDBID_GongStructDB[gongstructDB.ID] = gongstructDB
		BackRepoGongStructid_atBckpTime_newID[gongstructDB_ID_atBackupTime] = gongstructDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "GongStructDB.json" in dirPath that stores an array
// of GongStructDB and stores it in the database
// the map BackRepoGongStructid_atBckpTime_newID is updated accordingly
func (backRepoGongStruct *BackRepoGongStructStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoGongStructid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "GongStructDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json GongStruct file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*GongStructDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_GongStructDBID_GongStructDB
	for _, gongstructDB := range forRestore {

		gongstructDB_ID_atBackupTime := gongstructDB.ID
		gongstructDB.ID = 0
		_, err := backRepoGongStruct.db.Create(gongstructDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoGongStruct.Map_GongStructDBID_GongStructDB[gongstructDB.ID] = gongstructDB
		BackRepoGongStructid_atBckpTime_newID[gongstructDB_ID_atBackupTime] = gongstructDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json GongStruct file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<GongStruct>id_atBckpTime_newID
// to compute new index
func (backRepoGongStruct *BackRepoGongStructStruct) RestorePhaseTwo() {

	for _, gongstructDB := range backRepoGongStruct.Map_GongStructDBID_GongStructDB {

		// next line of code is to avert unused variable compilation error
		_ = gongstructDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		db, _ := backRepoGongStruct.db.Model(gongstructDB)
		_, err := db.Updates(*gongstructDB)
		if err != nil {
			log.Fatal(err)
		}
	}

}

// BackRepoGongStruct.ResetReversePointers commits all staged instances of GongStruct to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoGongStruct *BackRepoGongStructStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, gongstruct := range backRepoGongStruct.Map_GongStructDBID_GongStructPtr {
		backRepoGongStruct.ResetReversePointersInstance(backRepo, idx, gongstruct)
	}

	return
}

func (backRepoGongStruct *BackRepoGongStructStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, gongstruct *models.GongStruct) (Error error) {

	// fetch matching gongstructDB
	if gongstructDB, ok := backRepoGongStruct.Map_GongStructDBID_GongStructDB[idx]; ok {
		_ = gongstructDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoGongStructid_atBckpTime_newID map[uint]uint
