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
var dummy_Sequence_sql sql.NullBool
var dummy_Sequence_time time.Duration
var dummy_Sequence_sort sort.Float64Slice

// SequenceAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model sequenceAPI
type SequenceAPI struct {
	gorm.Model

	models.Sequence_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	SequencePointersEncoding SequencePointersEncoding
}

// SequencePointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type SequencePointersEncoding struct {
	// insertion for pointer fields encoding declaration

	// field Annotation is a pointer to another Struct (optional or 0..1)
	// This field is generated into another field to enable AS ONE association
	AnnotationID sql.NullInt64

	// field Sequences is a slice of pointers to another Struct (optional or 0..1)
	Sequences IntSlice `gorm:"type:TEXT"`

	// field Alls is a slice of pointers to another Struct (optional or 0..1)
	Alls IntSlice `gorm:"type:TEXT"`

	// field Choices is a slice of pointers to another Struct (optional or 0..1)
	Choices IntSlice `gorm:"type:TEXT"`

	// field Groups is a slice of pointers to another Struct (optional or 0..1)
	Groups IntSlice `gorm:"type:TEXT"`

	// field Elements is a slice of pointers to another Struct (optional or 0..1)
	Elements IntSlice `gorm:"type:TEXT"`
}

// SequenceDB describes a sequence in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model sequenceDB
type SequenceDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field sequenceDB.Name
	Name_Data sql.NullString

	// Declation for basic field sequenceDB.OuterElementName
	OuterElementName_Data sql.NullString

	// Declation for basic field sequenceDB.Order
	Order_Data sql.NullInt64

	// Declation for basic field sequenceDB.Depth
	Depth_Data sql.NullInt64

	// Declation for basic field sequenceDB.MinOccurs
	MinOccurs_Data sql.NullString

	// Declation for basic field sequenceDB.MaxOccurs
	MaxOccurs_Data sql.NullString

	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	SequencePointersEncoding
}

// SequenceDBs arrays sequenceDBs
// swagger:response sequenceDBsResponse
type SequenceDBs []SequenceDB

// SequenceDBResponse provides response
// swagger:response sequenceDBResponse
type SequenceDBResponse struct {
	SequenceDB
}

// SequenceWOP is a Sequence without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type SequenceWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	OuterElementName string `xlsx:"2"`

	Order int `xlsx:"3"`

	Depth int `xlsx:"4"`

	MinOccurs string `xlsx:"5"`

	MaxOccurs string `xlsx:"6"`
	// insertion for WOP pointer fields
}

var Sequence_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"OuterElementName",
	"Order",
	"Depth",
	"MinOccurs",
	"MaxOccurs",
}

type BackRepoSequenceStruct struct {
	// stores SequenceDB according to their gorm ID
	Map_SequenceDBID_SequenceDB map[uint]*SequenceDB

	// stores SequenceDB ID according to Sequence address
	Map_SequencePtr_SequenceDBID map[*models.Sequence]uint

	// stores Sequence according to their gorm ID
	Map_SequenceDBID_SequencePtr map[uint]*models.Sequence

	db db.DBInterface

	stage *models.Stage
}

func (backRepoSequence *BackRepoSequenceStruct) GetStage() (stage *models.Stage) {
	stage = backRepoSequence.stage
	return
}

func (backRepoSequence *BackRepoSequenceStruct) GetDB() db.DBInterface {
	return backRepoSequence.db
}

// GetSequenceDBFromSequencePtr is a handy function to access the back repo instance from the stage instance
func (backRepoSequence *BackRepoSequenceStruct) GetSequenceDBFromSequencePtr(sequence *models.Sequence) (sequenceDB *SequenceDB) {
	id := backRepoSequence.Map_SequencePtr_SequenceDBID[sequence]
	sequenceDB = backRepoSequence.Map_SequenceDBID_SequenceDB[id]
	return
}

// BackRepoSequence.CommitPhaseOne commits all staged instances of Sequence to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoSequence *BackRepoSequenceStruct) CommitPhaseOne(stage *models.Stage) (Error error) {

	var sequences []*models.Sequence
	for sequence := range stage.Sequences {
		sequences = append(sequences, sequence)
	}

	// Sort by the order stored in Map_Staged_Order.
	sort.Slice(sequences, func(i, j int) bool {
		return stage.SequenceMap_Staged_Order[sequences[i]] < stage.SequenceMap_Staged_Order[sequences[j]]
	})

	for _, sequence := range sequences {
		backRepoSequence.CommitPhaseOneInstance(sequence)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, sequence := range backRepoSequence.Map_SequenceDBID_SequencePtr {
		if _, ok := stage.Sequences[sequence]; !ok {
			backRepoSequence.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoSequence.CommitDeleteInstance commits deletion of Sequence to the BackRepo
func (backRepoSequence *BackRepoSequenceStruct) CommitDeleteInstance(id uint) (Error error) {

	sequence := backRepoSequence.Map_SequenceDBID_SequencePtr[id]

	// sequence is not staged anymore, remove sequenceDB
	sequenceDB := backRepoSequence.Map_SequenceDBID_SequenceDB[id]
	db, _ := backRepoSequence.db.Unscoped()
	_, err := db.Delete(sequenceDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	delete(backRepoSequence.Map_SequencePtr_SequenceDBID, sequence)
	delete(backRepoSequence.Map_SequenceDBID_SequencePtr, id)
	delete(backRepoSequence.Map_SequenceDBID_SequenceDB, id)

	return
}

// BackRepoSequence.CommitPhaseOneInstance commits sequence staged instances of Sequence to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoSequence *BackRepoSequenceStruct) CommitPhaseOneInstance(sequence *models.Sequence) (Error error) {

	// check if the sequence is not commited yet
	if _, ok := backRepoSequence.Map_SequencePtr_SequenceDBID[sequence]; ok {
		return
	}

	// initiate sequence
	var sequenceDB SequenceDB
	sequenceDB.CopyBasicFieldsFromSequence(sequence)

	_, err := backRepoSequence.db.Create(&sequenceDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	backRepoSequence.Map_SequencePtr_SequenceDBID[sequence] = sequenceDB.ID
	backRepoSequence.Map_SequenceDBID_SequencePtr[sequenceDB.ID] = sequence
	backRepoSequence.Map_SequenceDBID_SequenceDB[sequenceDB.ID] = &sequenceDB

	return
}

// BackRepoSequence.CommitPhaseTwo commits all staged instances of Sequence to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSequence *BackRepoSequenceStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, sequence := range backRepoSequence.Map_SequenceDBID_SequencePtr {
		backRepoSequence.CommitPhaseTwoInstance(backRepo, idx, sequence)
	}

	return
}

// BackRepoSequence.CommitPhaseTwoInstance commits {{structname }} of models.Sequence to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSequence *BackRepoSequenceStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, sequence *models.Sequence) (Error error) {

	// fetch matching sequenceDB
	if sequenceDB, ok := backRepoSequence.Map_SequenceDBID_SequenceDB[idx]; ok {

		sequenceDB.CopyBasicFieldsFromSequence(sequence)

		// insertion point for translating pointers encodings into actual pointers
		// commit pointer value sequence.Annotation translates to updating the sequence.AnnotationID
		sequenceDB.AnnotationID.Valid = true // allow for a 0 value (nil association)
		if sequence.Annotation != nil {
			if AnnotationId, ok := backRepo.BackRepoAnnotation.Map_AnnotationPtr_AnnotationDBID[sequence.Annotation]; ok {
				sequenceDB.AnnotationID.Int64 = int64(AnnotationId)
				sequenceDB.AnnotationID.Valid = true
			}
		} else {
			sequenceDB.AnnotationID.Int64 = 0
			sequenceDB.AnnotationID.Valid = true
		}

		// 1. reset
		sequenceDB.SequencePointersEncoding.Sequences = make([]int, 0)
		// 2. encode
		for _, sequenceAssocEnd := range sequence.Sequences {
			sequenceAssocEnd_DB :=
				backRepo.BackRepoSequence.GetSequenceDBFromSequencePtr(sequenceAssocEnd)
			
			// the stage might be inconsistant, meaning that the sequenceAssocEnd_DB might
			// be missing from the stage. In this case, the commit operation is robust
			// An alternative would be to crash here to reveal the missing element.
			if sequenceAssocEnd_DB == nil {
				continue
			}
			
			sequenceDB.SequencePointersEncoding.Sequences =
				append(sequenceDB.SequencePointersEncoding.Sequences, int(sequenceAssocEnd_DB.ID))
		}

		// 1. reset
		sequenceDB.SequencePointersEncoding.Alls = make([]int, 0)
		// 2. encode
		for _, allAssocEnd := range sequence.Alls {
			allAssocEnd_DB :=
				backRepo.BackRepoAll.GetAllDBFromAllPtr(allAssocEnd)
			
			// the stage might be inconsistant, meaning that the allAssocEnd_DB might
			// be missing from the stage. In this case, the commit operation is robust
			// An alternative would be to crash here to reveal the missing element.
			if allAssocEnd_DB == nil {
				continue
			}
			
			sequenceDB.SequencePointersEncoding.Alls =
				append(sequenceDB.SequencePointersEncoding.Alls, int(allAssocEnd_DB.ID))
		}

		// 1. reset
		sequenceDB.SequencePointersEncoding.Choices = make([]int, 0)
		// 2. encode
		for _, choiceAssocEnd := range sequence.Choices {
			choiceAssocEnd_DB :=
				backRepo.BackRepoChoice.GetChoiceDBFromChoicePtr(choiceAssocEnd)
			
			// the stage might be inconsistant, meaning that the choiceAssocEnd_DB might
			// be missing from the stage. In this case, the commit operation is robust
			// An alternative would be to crash here to reveal the missing element.
			if choiceAssocEnd_DB == nil {
				continue
			}
			
			sequenceDB.SequencePointersEncoding.Choices =
				append(sequenceDB.SequencePointersEncoding.Choices, int(choiceAssocEnd_DB.ID))
		}

		// 1. reset
		sequenceDB.SequencePointersEncoding.Groups = make([]int, 0)
		// 2. encode
		for _, groupAssocEnd := range sequence.Groups {
			groupAssocEnd_DB :=
				backRepo.BackRepoGroup.GetGroupDBFromGroupPtr(groupAssocEnd)
			
			// the stage might be inconsistant, meaning that the groupAssocEnd_DB might
			// be missing from the stage. In this case, the commit operation is robust
			// An alternative would be to crash here to reveal the missing element.
			if groupAssocEnd_DB == nil {
				continue
			}
			
			sequenceDB.SequencePointersEncoding.Groups =
				append(sequenceDB.SequencePointersEncoding.Groups, int(groupAssocEnd_DB.ID))
		}

		// 1. reset
		sequenceDB.SequencePointersEncoding.Elements = make([]int, 0)
		// 2. encode
		for _, elementAssocEnd := range sequence.Elements {
			elementAssocEnd_DB :=
				backRepo.BackRepoElement.GetElementDBFromElementPtr(elementAssocEnd)
			
			// the stage might be inconsistant, meaning that the elementAssocEnd_DB might
			// be missing from the stage. In this case, the commit operation is robust
			// An alternative would be to crash here to reveal the missing element.
			if elementAssocEnd_DB == nil {
				continue
			}
			
			sequenceDB.SequencePointersEncoding.Elements =
				append(sequenceDB.SequencePointersEncoding.Elements, int(elementAssocEnd_DB.ID))
		}

		_, err := backRepoSequence.db.Save(sequenceDB)
		if err != nil {
			log.Fatal(err)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown Sequence intance %s", sequence.Name))
		return err
	}

	return
}

// BackRepoSequence.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoSequence *BackRepoSequenceStruct) CheckoutPhaseOne() (Error error) {

	sequenceDBArray := make([]SequenceDB, 0)
	_, err := backRepoSequence.db.Find(&sequenceDBArray)
	if err != nil {
		return err
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	sequenceInstancesToBeRemovedFromTheStage := make(map[*models.Sequence]any)
	for key, value := range backRepoSequence.stage.Sequences {
		sequenceInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, sequenceDB := range sequenceDBArray {
		backRepoSequence.CheckoutPhaseOneInstance(&sequenceDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		sequence, ok := backRepoSequence.Map_SequenceDBID_SequencePtr[sequenceDB.ID]
		if ok {
			delete(sequenceInstancesToBeRemovedFromTheStage, sequence)
		}
	}

	// remove from stage and back repo's 3 maps all sequences that are not in the checkout
	for sequence := range sequenceInstancesToBeRemovedFromTheStage {
		sequence.Unstage(backRepoSequence.GetStage())

		// remove instance from the back repo 3 maps
		sequenceID := backRepoSequence.Map_SequencePtr_SequenceDBID[sequence]
		delete(backRepoSequence.Map_SequencePtr_SequenceDBID, sequence)
		delete(backRepoSequence.Map_SequenceDBID_SequenceDB, sequenceID)
		delete(backRepoSequence.Map_SequenceDBID_SequencePtr, sequenceID)
	}

	return
}

// CheckoutPhaseOneInstance takes a sequenceDB that has been found in the DB, updates the backRepo and stages the
// models version of the sequenceDB
func (backRepoSequence *BackRepoSequenceStruct) CheckoutPhaseOneInstance(sequenceDB *SequenceDB) (Error error) {

	sequence, ok := backRepoSequence.Map_SequenceDBID_SequencePtr[sequenceDB.ID]
	if !ok {
		sequence = new(models.Sequence)

		backRepoSequence.Map_SequenceDBID_SequencePtr[sequenceDB.ID] = sequence
		backRepoSequence.Map_SequencePtr_SequenceDBID[sequence] = sequenceDB.ID

		// append model store with the new element
		sequence.Name = sequenceDB.Name_Data.String
		sequence.Stage(backRepoSequence.GetStage())
	}
	sequenceDB.CopyBasicFieldsToSequence(sequence)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	sequence.Stage(backRepoSequence.GetStage())

	// preserve pointer to sequenceDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_SequenceDBID_SequenceDB)[sequenceDB hold variable pointers
	sequenceDB_Data := *sequenceDB
	preservedPtrToSequence := &sequenceDB_Data
	backRepoSequence.Map_SequenceDBID_SequenceDB[sequenceDB.ID] = preservedPtrToSequence

	return
}

// BackRepoSequence.CheckoutPhaseTwo Checkouts all staged instances of Sequence to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSequence *BackRepoSequenceStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, sequenceDB := range backRepoSequence.Map_SequenceDBID_SequenceDB {
		backRepoSequence.CheckoutPhaseTwoInstance(backRepo, sequenceDB)
	}
	return
}

// BackRepoSequence.CheckoutPhaseTwoInstance Checkouts staged instances of Sequence to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSequence *BackRepoSequenceStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, sequenceDB *SequenceDB) (Error error) {

	sequence := backRepoSequence.Map_SequenceDBID_SequencePtr[sequenceDB.ID]

	sequenceDB.DecodePointers(backRepo, sequence)

	return
}

func (sequenceDB *SequenceDB) DecodePointers(backRepo *BackRepoStruct, sequence *models.Sequence) {

	// insertion point for checkout of pointer encoding
	// Annotation field	
	{
		id := sequenceDB.AnnotationID.Int64
		if id != 0 {
			tmp, ok := backRepo.BackRepoAnnotation.Map_AnnotationDBID_AnnotationPtr[uint(id)]

			// if the pointer id is unknown, it is not a problem, maybe the target was removed from the front
			if !ok {
				log.Println("DecodePointers: sequence.Annotation, unknown pointer id", id)
				sequence.Annotation = nil
			} else {
				// updates only if field has changed
				if sequence.Annotation == nil || sequence.Annotation != tmp {
					sequence.Annotation = tmp
				}
			}
		} else {
			sequence.Annotation = nil
		}
	}
	
	// This loop redeem sequence.Sequences in the stage from the encode in the back repo
	// It parses all SequenceDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	sequence.Sequences = sequence.Sequences[:0]
	for _, _Sequenceid := range sequenceDB.SequencePointersEncoding.Sequences {
		sequence.Sequences = append(sequence.Sequences, backRepo.BackRepoSequence.Map_SequenceDBID_SequencePtr[uint(_Sequenceid)])
	}

	// This loop redeem sequence.Alls in the stage from the encode in the back repo
	// It parses all AllDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	sequence.Alls = sequence.Alls[:0]
	for _, _Allid := range sequenceDB.SequencePointersEncoding.Alls {
		sequence.Alls = append(sequence.Alls, backRepo.BackRepoAll.Map_AllDBID_AllPtr[uint(_Allid)])
	}

	// This loop redeem sequence.Choices in the stage from the encode in the back repo
	// It parses all ChoiceDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	sequence.Choices = sequence.Choices[:0]
	for _, _Choiceid := range sequenceDB.SequencePointersEncoding.Choices {
		sequence.Choices = append(sequence.Choices, backRepo.BackRepoChoice.Map_ChoiceDBID_ChoicePtr[uint(_Choiceid)])
	}

	// This loop redeem sequence.Groups in the stage from the encode in the back repo
	// It parses all GroupDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	sequence.Groups = sequence.Groups[:0]
	for _, _Groupid := range sequenceDB.SequencePointersEncoding.Groups {
		sequence.Groups = append(sequence.Groups, backRepo.BackRepoGroup.Map_GroupDBID_GroupPtr[uint(_Groupid)])
	}

	// This loop redeem sequence.Elements in the stage from the encode in the back repo
	// It parses all ElementDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	sequence.Elements = sequence.Elements[:0]
	for _, _Elementid := range sequenceDB.SequencePointersEncoding.Elements {
		sequence.Elements = append(sequence.Elements, backRepo.BackRepoElement.Map_ElementDBID_ElementPtr[uint(_Elementid)])
	}

	return
}

// CommitSequence allows commit of a single sequence (if already staged)
func (backRepo *BackRepoStruct) CommitSequence(sequence *models.Sequence) {
	backRepo.BackRepoSequence.CommitPhaseOneInstance(sequence)
	if id, ok := backRepo.BackRepoSequence.Map_SequencePtr_SequenceDBID[sequence]; ok {
		backRepo.BackRepoSequence.CommitPhaseTwoInstance(backRepo, id, sequence)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitSequence allows checkout of a single sequence (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutSequence(sequence *models.Sequence) {
	// check if the sequence is staged
	if _, ok := backRepo.BackRepoSequence.Map_SequencePtr_SequenceDBID[sequence]; ok {

		if id, ok := backRepo.BackRepoSequence.Map_SequencePtr_SequenceDBID[sequence]; ok {
			var sequenceDB SequenceDB
			sequenceDB.ID = id

			if _, err := backRepo.BackRepoSequence.db.First(&sequenceDB, id); err != nil {
				log.Fatalln("CheckoutSequence : Problem with getting object with id:", id)
			}
			backRepo.BackRepoSequence.CheckoutPhaseOneInstance(&sequenceDB)
			backRepo.BackRepoSequence.CheckoutPhaseTwoInstance(backRepo, &sequenceDB)
		}
	}
}

// CopyBasicFieldsFromSequence
func (sequenceDB *SequenceDB) CopyBasicFieldsFromSequence(sequence *models.Sequence) {
	// insertion point for fields commit

	sequenceDB.Name_Data.String = sequence.Name
	sequenceDB.Name_Data.Valid = true

	sequenceDB.OuterElementName_Data.String = sequence.OuterElementName
	sequenceDB.OuterElementName_Data.Valid = true

	sequenceDB.Order_Data.Int64 = int64(sequence.Order)
	sequenceDB.Order_Data.Valid = true

	sequenceDB.Depth_Data.Int64 = int64(sequence.Depth)
	sequenceDB.Depth_Data.Valid = true

	sequenceDB.MinOccurs_Data.String = sequence.MinOccurs
	sequenceDB.MinOccurs_Data.Valid = true

	sequenceDB.MaxOccurs_Data.String = sequence.MaxOccurs
	sequenceDB.MaxOccurs_Data.Valid = true
}

// CopyBasicFieldsFromSequence_WOP
func (sequenceDB *SequenceDB) CopyBasicFieldsFromSequence_WOP(sequence *models.Sequence_WOP) {
	// insertion point for fields commit

	sequenceDB.Name_Data.String = sequence.Name
	sequenceDB.Name_Data.Valid = true

	sequenceDB.OuterElementName_Data.String = sequence.OuterElementName
	sequenceDB.OuterElementName_Data.Valid = true

	sequenceDB.Order_Data.Int64 = int64(sequence.Order)
	sequenceDB.Order_Data.Valid = true

	sequenceDB.Depth_Data.Int64 = int64(sequence.Depth)
	sequenceDB.Depth_Data.Valid = true

	sequenceDB.MinOccurs_Data.String = sequence.MinOccurs
	sequenceDB.MinOccurs_Data.Valid = true

	sequenceDB.MaxOccurs_Data.String = sequence.MaxOccurs
	sequenceDB.MaxOccurs_Data.Valid = true
}

// CopyBasicFieldsFromSequenceWOP
func (sequenceDB *SequenceDB) CopyBasicFieldsFromSequenceWOP(sequence *SequenceWOP) {
	// insertion point for fields commit

	sequenceDB.Name_Data.String = sequence.Name
	sequenceDB.Name_Data.Valid = true

	sequenceDB.OuterElementName_Data.String = sequence.OuterElementName
	sequenceDB.OuterElementName_Data.Valid = true

	sequenceDB.Order_Data.Int64 = int64(sequence.Order)
	sequenceDB.Order_Data.Valid = true

	sequenceDB.Depth_Data.Int64 = int64(sequence.Depth)
	sequenceDB.Depth_Data.Valid = true

	sequenceDB.MinOccurs_Data.String = sequence.MinOccurs
	sequenceDB.MinOccurs_Data.Valid = true

	sequenceDB.MaxOccurs_Data.String = sequence.MaxOccurs
	sequenceDB.MaxOccurs_Data.Valid = true
}

// CopyBasicFieldsToSequence
func (sequenceDB *SequenceDB) CopyBasicFieldsToSequence(sequence *models.Sequence) {
	// insertion point for checkout of basic fields (back repo to stage)
	sequence.Name = sequenceDB.Name_Data.String
	sequence.OuterElementName = sequenceDB.OuterElementName_Data.String
	sequence.Order = int(sequenceDB.Order_Data.Int64)
	sequence.Depth = int(sequenceDB.Depth_Data.Int64)
	sequence.MinOccurs = sequenceDB.MinOccurs_Data.String
	sequence.MaxOccurs = sequenceDB.MaxOccurs_Data.String
}

// CopyBasicFieldsToSequence_WOP
func (sequenceDB *SequenceDB) CopyBasicFieldsToSequence_WOP(sequence *models.Sequence_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	sequence.Name = sequenceDB.Name_Data.String
	sequence.OuterElementName = sequenceDB.OuterElementName_Data.String
	sequence.Order = int(sequenceDB.Order_Data.Int64)
	sequence.Depth = int(sequenceDB.Depth_Data.Int64)
	sequence.MinOccurs = sequenceDB.MinOccurs_Data.String
	sequence.MaxOccurs = sequenceDB.MaxOccurs_Data.String
}

// CopyBasicFieldsToSequenceWOP
func (sequenceDB *SequenceDB) CopyBasicFieldsToSequenceWOP(sequence *SequenceWOP) {
	sequence.ID = int(sequenceDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	sequence.Name = sequenceDB.Name_Data.String
	sequence.OuterElementName = sequenceDB.OuterElementName_Data.String
	sequence.Order = int(sequenceDB.Order_Data.Int64)
	sequence.Depth = int(sequenceDB.Depth_Data.Int64)
	sequence.MinOccurs = sequenceDB.MinOccurs_Data.String
	sequence.MaxOccurs = sequenceDB.MaxOccurs_Data.String
}

// Backup generates a json file from a slice of all SequenceDB instances in the backrepo
func (backRepoSequence *BackRepoSequenceStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "SequenceDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*SequenceDB, 0)
	for _, sequenceDB := range backRepoSequence.Map_SequenceDBID_SequenceDB {
		forBackup = append(forBackup, sequenceDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json Sequence ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json Sequence file", err.Error())
	}
}

// Backup generates a json file from a slice of all SequenceDB instances in the backrepo
func (backRepoSequence *BackRepoSequenceStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*SequenceDB, 0)
	for _, sequenceDB := range backRepoSequence.Map_SequenceDBID_SequenceDB {
		forBackup = append(forBackup, sequenceDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("Sequence")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&Sequence_Fields, -1)
	for _, sequenceDB := range forBackup {

		var sequenceWOP SequenceWOP
		sequenceDB.CopyBasicFieldsToSequenceWOP(&sequenceWOP)

		row := sh.AddRow()
		row.WriteStruct(&sequenceWOP, -1)
	}
}

// RestoreXL from the "Sequence" sheet all SequenceDB instances
func (backRepoSequence *BackRepoSequenceStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoSequenceid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["Sequence"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoSequence.rowVisitorSequence)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoSequence *BackRepoSequenceStruct) rowVisitorSequence(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var sequenceWOP SequenceWOP
		row.ReadStruct(&sequenceWOP)

		// add the unmarshalled struct to the stage
		sequenceDB := new(SequenceDB)
		sequenceDB.CopyBasicFieldsFromSequenceWOP(&sequenceWOP)

		sequenceDB_ID_atBackupTime := sequenceDB.ID
		sequenceDB.ID = 0
		_, err := backRepoSequence.db.Create(sequenceDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoSequence.Map_SequenceDBID_SequenceDB[sequenceDB.ID] = sequenceDB
		BackRepoSequenceid_atBckpTime_newID[sequenceDB_ID_atBackupTime] = sequenceDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "SequenceDB.json" in dirPath that stores an array
// of SequenceDB and stores it in the database
// the map BackRepoSequenceid_atBckpTime_newID is updated accordingly
func (backRepoSequence *BackRepoSequenceStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoSequenceid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "SequenceDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json Sequence file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*SequenceDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_SequenceDBID_SequenceDB
	for _, sequenceDB := range forRestore {

		sequenceDB_ID_atBackupTime := sequenceDB.ID
		sequenceDB.ID = 0
		_, err := backRepoSequence.db.Create(sequenceDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoSequence.Map_SequenceDBID_SequenceDB[sequenceDB.ID] = sequenceDB
		BackRepoSequenceid_atBckpTime_newID[sequenceDB_ID_atBackupTime] = sequenceDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json Sequence file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<Sequence>id_atBckpTime_newID
// to compute new index
func (backRepoSequence *BackRepoSequenceStruct) RestorePhaseTwo() {

	for _, sequenceDB := range backRepoSequence.Map_SequenceDBID_SequenceDB {

		// next line of code is to avert unused variable compilation error
		_ = sequenceDB

		// insertion point for reindexing pointers encoding
		// reindexing Annotation field
		if sequenceDB.AnnotationID.Int64 != 0 {
			sequenceDB.AnnotationID.Int64 = int64(BackRepoAnnotationid_atBckpTime_newID[uint(sequenceDB.AnnotationID.Int64)])
			sequenceDB.AnnotationID.Valid = true
		}

		// update databse with new index encoding
		db, _ := backRepoSequence.db.Model(sequenceDB)
		_, err := db.Updates(*sequenceDB)
		if err != nil {
			log.Fatal(err)
		}
	}

}

// BackRepoSequence.ResetReversePointers commits all staged instances of Sequence to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoSequence *BackRepoSequenceStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, sequence := range backRepoSequence.Map_SequenceDBID_SequencePtr {
		backRepoSequence.ResetReversePointersInstance(backRepo, idx, sequence)
	}

	return
}

func (backRepoSequence *BackRepoSequenceStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, sequence *models.Sequence) (Error error) {

	// fetch matching sequenceDB
	if sequenceDB, ok := backRepoSequence.Map_SequenceDBID_SequenceDB[idx]; ok {
		_ = sequenceDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoSequenceid_atBckpTime_newID map[uint]uint
