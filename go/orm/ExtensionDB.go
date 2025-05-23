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
var dummy_Extension_sql sql.NullBool
var dummy_Extension_time time.Duration
var dummy_Extension_sort sort.Float64Slice

// ExtensionAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model extensionAPI
type ExtensionAPI struct {
	gorm.Model

	models.Extension_WOP

	// encoding of pointers
	// for API, it cannot be embedded
	ExtensionPointersEncoding ExtensionPointersEncoding
}

// ExtensionPointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type ExtensionPointersEncoding struct {
	// insertion for pointer fields encoding declaration

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

	// field Attributes is a slice of pointers to another Struct (optional or 0..1)
	Attributes IntSlice `gorm:"type:TEXT"`

	// field AttributeGroups is a slice of pointers to another Struct (optional or 0..1)
	AttributeGroups IntSlice `gorm:"type:TEXT"`
}

// ExtensionDB describes a extension in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model extensionDB
type ExtensionDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field extensionDB.Name
	Name_Data sql.NullString

	// Declation for basic field extensionDB.OuterElementName
	OuterElementName_Data sql.NullString

	// Declation for basic field extensionDB.Order
	Order_Data sql.NullInt64

	// Declation for basic field extensionDB.Depth
	Depth_Data sql.NullInt64

	// Declation for basic field extensionDB.MinOccurs
	MinOccurs_Data sql.NullString

	// Declation for basic field extensionDB.MaxOccurs
	MaxOccurs_Data sql.NullString

	// Declation for basic field extensionDB.Base
	Base_Data sql.NullString

	// Declation for basic field extensionDB.Ref
	Ref_Data sql.NullString

	// encoding of pointers
	// for GORM serialization, it is necessary to embed to Pointer Encoding declaration
	ExtensionPointersEncoding
}

// ExtensionDBs arrays extensionDBs
// swagger:response extensionDBsResponse
type ExtensionDBs []ExtensionDB

// ExtensionDBResponse provides response
// swagger:response extensionDBResponse
type ExtensionDBResponse struct {
	ExtensionDB
}

// ExtensionWOP is a Extension without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type ExtensionWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	OuterElementName string `xlsx:"2"`

	Order int `xlsx:"3"`

	Depth int `xlsx:"4"`

	MinOccurs string `xlsx:"5"`

	MaxOccurs string `xlsx:"6"`

	Base string `xlsx:"7"`

	Ref string `xlsx:"8"`
	// insertion for WOP pointer fields
}

var Extension_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"OuterElementName",
	"Order",
	"Depth",
	"MinOccurs",
	"MaxOccurs",
	"Base",
	"Ref",
}

type BackRepoExtensionStruct struct {
	// stores ExtensionDB according to their gorm ID
	Map_ExtensionDBID_ExtensionDB map[uint]*ExtensionDB

	// stores ExtensionDB ID according to Extension address
	Map_ExtensionPtr_ExtensionDBID map[*models.Extension]uint

	// stores Extension according to their gorm ID
	Map_ExtensionDBID_ExtensionPtr map[uint]*models.Extension

	db db.DBInterface

	stage *models.Stage
}

func (backRepoExtension *BackRepoExtensionStruct) GetStage() (stage *models.Stage) {
	stage = backRepoExtension.stage
	return
}

func (backRepoExtension *BackRepoExtensionStruct) GetDB() db.DBInterface {
	return backRepoExtension.db
}

// GetExtensionDBFromExtensionPtr is a handy function to access the back repo instance from the stage instance
func (backRepoExtension *BackRepoExtensionStruct) GetExtensionDBFromExtensionPtr(extension *models.Extension) (extensionDB *ExtensionDB) {
	id := backRepoExtension.Map_ExtensionPtr_ExtensionDBID[extension]
	extensionDB = backRepoExtension.Map_ExtensionDBID_ExtensionDB[id]
	return
}

// BackRepoExtension.CommitPhaseOne commits all staged instances of Extension to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoExtension *BackRepoExtensionStruct) CommitPhaseOne(stage *models.Stage) (Error error) {

	var extensions []*models.Extension
	for extension := range stage.Extensions {
		extensions = append(extensions, extension)
	}

	// Sort by the order stored in Map_Staged_Order.
	sort.Slice(extensions, func(i, j int) bool {
		return stage.ExtensionMap_Staged_Order[extensions[i]] < stage.ExtensionMap_Staged_Order[extensions[j]]
	})

	for _, extension := range extensions {
		backRepoExtension.CommitPhaseOneInstance(extension)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, extension := range backRepoExtension.Map_ExtensionDBID_ExtensionPtr {
		if _, ok := stage.Extensions[extension]; !ok {
			backRepoExtension.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoExtension.CommitDeleteInstance commits deletion of Extension to the BackRepo
func (backRepoExtension *BackRepoExtensionStruct) CommitDeleteInstance(id uint) (Error error) {

	extension := backRepoExtension.Map_ExtensionDBID_ExtensionPtr[id]

	// extension is not staged anymore, remove extensionDB
	extensionDB := backRepoExtension.Map_ExtensionDBID_ExtensionDB[id]
	db, _ := backRepoExtension.db.Unscoped()
	_, err := db.Delete(extensionDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	delete(backRepoExtension.Map_ExtensionPtr_ExtensionDBID, extension)
	delete(backRepoExtension.Map_ExtensionDBID_ExtensionPtr, id)
	delete(backRepoExtension.Map_ExtensionDBID_ExtensionDB, id)

	return
}

// BackRepoExtension.CommitPhaseOneInstance commits extension staged instances of Extension to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoExtension *BackRepoExtensionStruct) CommitPhaseOneInstance(extension *models.Extension) (Error error) {

	// check if the extension is not commited yet
	if _, ok := backRepoExtension.Map_ExtensionPtr_ExtensionDBID[extension]; ok {
		return
	}

	// initiate extension
	var extensionDB ExtensionDB
	extensionDB.CopyBasicFieldsFromExtension(extension)

	_, err := backRepoExtension.db.Create(&extensionDB)
	if err != nil {
		log.Fatal(err)
	}

	// update stores
	backRepoExtension.Map_ExtensionPtr_ExtensionDBID[extension] = extensionDB.ID
	backRepoExtension.Map_ExtensionDBID_ExtensionPtr[extensionDB.ID] = extension
	backRepoExtension.Map_ExtensionDBID_ExtensionDB[extensionDB.ID] = &extensionDB

	return
}

// BackRepoExtension.CommitPhaseTwo commits all staged instances of Extension to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoExtension *BackRepoExtensionStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, extension := range backRepoExtension.Map_ExtensionDBID_ExtensionPtr {
		backRepoExtension.CommitPhaseTwoInstance(backRepo, idx, extension)
	}

	return
}

// BackRepoExtension.CommitPhaseTwoInstance commits {{structname }} of models.Extension to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoExtension *BackRepoExtensionStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, extension *models.Extension) (Error error) {

	// fetch matching extensionDB
	if extensionDB, ok := backRepoExtension.Map_ExtensionDBID_ExtensionDB[idx]; ok {

		extensionDB.CopyBasicFieldsFromExtension(extension)

		// insertion point for translating pointers encodings into actual pointers
		// 1. reset
		extensionDB.ExtensionPointersEncoding.Sequences = make([]int, 0)
		// 2. encode
		for _, sequenceAssocEnd := range extension.Sequences {
			sequenceAssocEnd_DB :=
				backRepo.BackRepoSequence.GetSequenceDBFromSequencePtr(sequenceAssocEnd)
			
			// the stage might be inconsistant, meaning that the sequenceAssocEnd_DB might
			// be missing from the stage. In this case, the commit operation is robust
			// An alternative would be to crash here to reveal the missing element.
			if sequenceAssocEnd_DB == nil {
				continue
			}
			
			extensionDB.ExtensionPointersEncoding.Sequences =
				append(extensionDB.ExtensionPointersEncoding.Sequences, int(sequenceAssocEnd_DB.ID))
		}

		// 1. reset
		extensionDB.ExtensionPointersEncoding.Alls = make([]int, 0)
		// 2. encode
		for _, allAssocEnd := range extension.Alls {
			allAssocEnd_DB :=
				backRepo.BackRepoAll.GetAllDBFromAllPtr(allAssocEnd)
			
			// the stage might be inconsistant, meaning that the allAssocEnd_DB might
			// be missing from the stage. In this case, the commit operation is robust
			// An alternative would be to crash here to reveal the missing element.
			if allAssocEnd_DB == nil {
				continue
			}
			
			extensionDB.ExtensionPointersEncoding.Alls =
				append(extensionDB.ExtensionPointersEncoding.Alls, int(allAssocEnd_DB.ID))
		}

		// 1. reset
		extensionDB.ExtensionPointersEncoding.Choices = make([]int, 0)
		// 2. encode
		for _, choiceAssocEnd := range extension.Choices {
			choiceAssocEnd_DB :=
				backRepo.BackRepoChoice.GetChoiceDBFromChoicePtr(choiceAssocEnd)
			
			// the stage might be inconsistant, meaning that the choiceAssocEnd_DB might
			// be missing from the stage. In this case, the commit operation is robust
			// An alternative would be to crash here to reveal the missing element.
			if choiceAssocEnd_DB == nil {
				continue
			}
			
			extensionDB.ExtensionPointersEncoding.Choices =
				append(extensionDB.ExtensionPointersEncoding.Choices, int(choiceAssocEnd_DB.ID))
		}

		// 1. reset
		extensionDB.ExtensionPointersEncoding.Groups = make([]int, 0)
		// 2. encode
		for _, groupAssocEnd := range extension.Groups {
			groupAssocEnd_DB :=
				backRepo.BackRepoGroup.GetGroupDBFromGroupPtr(groupAssocEnd)
			
			// the stage might be inconsistant, meaning that the groupAssocEnd_DB might
			// be missing from the stage. In this case, the commit operation is robust
			// An alternative would be to crash here to reveal the missing element.
			if groupAssocEnd_DB == nil {
				continue
			}
			
			extensionDB.ExtensionPointersEncoding.Groups =
				append(extensionDB.ExtensionPointersEncoding.Groups, int(groupAssocEnd_DB.ID))
		}

		// 1. reset
		extensionDB.ExtensionPointersEncoding.Elements = make([]int, 0)
		// 2. encode
		for _, elementAssocEnd := range extension.Elements {
			elementAssocEnd_DB :=
				backRepo.BackRepoElement.GetElementDBFromElementPtr(elementAssocEnd)
			
			// the stage might be inconsistant, meaning that the elementAssocEnd_DB might
			// be missing from the stage. In this case, the commit operation is robust
			// An alternative would be to crash here to reveal the missing element.
			if elementAssocEnd_DB == nil {
				continue
			}
			
			extensionDB.ExtensionPointersEncoding.Elements =
				append(extensionDB.ExtensionPointersEncoding.Elements, int(elementAssocEnd_DB.ID))
		}

		// 1. reset
		extensionDB.ExtensionPointersEncoding.Attributes = make([]int, 0)
		// 2. encode
		for _, attributeAssocEnd := range extension.Attributes {
			attributeAssocEnd_DB :=
				backRepo.BackRepoAttribute.GetAttributeDBFromAttributePtr(attributeAssocEnd)
			
			// the stage might be inconsistant, meaning that the attributeAssocEnd_DB might
			// be missing from the stage. In this case, the commit operation is robust
			// An alternative would be to crash here to reveal the missing element.
			if attributeAssocEnd_DB == nil {
				continue
			}
			
			extensionDB.ExtensionPointersEncoding.Attributes =
				append(extensionDB.ExtensionPointersEncoding.Attributes, int(attributeAssocEnd_DB.ID))
		}

		// 1. reset
		extensionDB.ExtensionPointersEncoding.AttributeGroups = make([]int, 0)
		// 2. encode
		for _, attributegroupAssocEnd := range extension.AttributeGroups {
			attributegroupAssocEnd_DB :=
				backRepo.BackRepoAttributeGroup.GetAttributeGroupDBFromAttributeGroupPtr(attributegroupAssocEnd)
			
			// the stage might be inconsistant, meaning that the attributegroupAssocEnd_DB might
			// be missing from the stage. In this case, the commit operation is robust
			// An alternative would be to crash here to reveal the missing element.
			if attributegroupAssocEnd_DB == nil {
				continue
			}
			
			extensionDB.ExtensionPointersEncoding.AttributeGroups =
				append(extensionDB.ExtensionPointersEncoding.AttributeGroups, int(attributegroupAssocEnd_DB.ID))
		}

		_, err := backRepoExtension.db.Save(extensionDB)
		if err != nil {
			log.Fatal(err)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown Extension intance %s", extension.Name))
		return err
	}

	return
}

// BackRepoExtension.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoExtension *BackRepoExtensionStruct) CheckoutPhaseOne() (Error error) {

	extensionDBArray := make([]ExtensionDB, 0)
	_, err := backRepoExtension.db.Find(&extensionDBArray)
	if err != nil {
		return err
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	extensionInstancesToBeRemovedFromTheStage := make(map[*models.Extension]any)
	for key, value := range backRepoExtension.stage.Extensions {
		extensionInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, extensionDB := range extensionDBArray {
		backRepoExtension.CheckoutPhaseOneInstance(&extensionDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		extension, ok := backRepoExtension.Map_ExtensionDBID_ExtensionPtr[extensionDB.ID]
		if ok {
			delete(extensionInstancesToBeRemovedFromTheStage, extension)
		}
	}

	// remove from stage and back repo's 3 maps all extensions that are not in the checkout
	for extension := range extensionInstancesToBeRemovedFromTheStage {
		extension.Unstage(backRepoExtension.GetStage())

		// remove instance from the back repo 3 maps
		extensionID := backRepoExtension.Map_ExtensionPtr_ExtensionDBID[extension]
		delete(backRepoExtension.Map_ExtensionPtr_ExtensionDBID, extension)
		delete(backRepoExtension.Map_ExtensionDBID_ExtensionDB, extensionID)
		delete(backRepoExtension.Map_ExtensionDBID_ExtensionPtr, extensionID)
	}

	return
}

// CheckoutPhaseOneInstance takes a extensionDB that has been found in the DB, updates the backRepo and stages the
// models version of the extensionDB
func (backRepoExtension *BackRepoExtensionStruct) CheckoutPhaseOneInstance(extensionDB *ExtensionDB) (Error error) {

	extension, ok := backRepoExtension.Map_ExtensionDBID_ExtensionPtr[extensionDB.ID]
	if !ok {
		extension = new(models.Extension)

		backRepoExtension.Map_ExtensionDBID_ExtensionPtr[extensionDB.ID] = extension
		backRepoExtension.Map_ExtensionPtr_ExtensionDBID[extension] = extensionDB.ID

		// append model store with the new element
		extension.Name = extensionDB.Name_Data.String
		extension.Stage(backRepoExtension.GetStage())
	}
	extensionDB.CopyBasicFieldsToExtension(extension)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	extension.Stage(backRepoExtension.GetStage())

	// preserve pointer to extensionDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_ExtensionDBID_ExtensionDB)[extensionDB hold variable pointers
	extensionDB_Data := *extensionDB
	preservedPtrToExtension := &extensionDB_Data
	backRepoExtension.Map_ExtensionDBID_ExtensionDB[extensionDB.ID] = preservedPtrToExtension

	return
}

// BackRepoExtension.CheckoutPhaseTwo Checkouts all staged instances of Extension to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoExtension *BackRepoExtensionStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, extensionDB := range backRepoExtension.Map_ExtensionDBID_ExtensionDB {
		backRepoExtension.CheckoutPhaseTwoInstance(backRepo, extensionDB)
	}
	return
}

// BackRepoExtension.CheckoutPhaseTwoInstance Checkouts staged instances of Extension to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoExtension *BackRepoExtensionStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, extensionDB *ExtensionDB) (Error error) {

	extension := backRepoExtension.Map_ExtensionDBID_ExtensionPtr[extensionDB.ID]

	extensionDB.DecodePointers(backRepo, extension)

	return
}

func (extensionDB *ExtensionDB) DecodePointers(backRepo *BackRepoStruct, extension *models.Extension) {

	// insertion point for checkout of pointer encoding
	// This loop redeem extension.Sequences in the stage from the encode in the back repo
	// It parses all SequenceDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	extension.Sequences = extension.Sequences[:0]
	for _, _Sequenceid := range extensionDB.ExtensionPointersEncoding.Sequences {
		extension.Sequences = append(extension.Sequences, backRepo.BackRepoSequence.Map_SequenceDBID_SequencePtr[uint(_Sequenceid)])
	}

	// This loop redeem extension.Alls in the stage from the encode in the back repo
	// It parses all AllDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	extension.Alls = extension.Alls[:0]
	for _, _Allid := range extensionDB.ExtensionPointersEncoding.Alls {
		extension.Alls = append(extension.Alls, backRepo.BackRepoAll.Map_AllDBID_AllPtr[uint(_Allid)])
	}

	// This loop redeem extension.Choices in the stage from the encode in the back repo
	// It parses all ChoiceDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	extension.Choices = extension.Choices[:0]
	for _, _Choiceid := range extensionDB.ExtensionPointersEncoding.Choices {
		extension.Choices = append(extension.Choices, backRepo.BackRepoChoice.Map_ChoiceDBID_ChoicePtr[uint(_Choiceid)])
	}

	// This loop redeem extension.Groups in the stage from the encode in the back repo
	// It parses all GroupDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	extension.Groups = extension.Groups[:0]
	for _, _Groupid := range extensionDB.ExtensionPointersEncoding.Groups {
		extension.Groups = append(extension.Groups, backRepo.BackRepoGroup.Map_GroupDBID_GroupPtr[uint(_Groupid)])
	}

	// This loop redeem extension.Elements in the stage from the encode in the back repo
	// It parses all ElementDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	extension.Elements = extension.Elements[:0]
	for _, _Elementid := range extensionDB.ExtensionPointersEncoding.Elements {
		extension.Elements = append(extension.Elements, backRepo.BackRepoElement.Map_ElementDBID_ElementPtr[uint(_Elementid)])
	}

	// This loop redeem extension.Attributes in the stage from the encode in the back repo
	// It parses all AttributeDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	extension.Attributes = extension.Attributes[:0]
	for _, _Attributeid := range extensionDB.ExtensionPointersEncoding.Attributes {
		extension.Attributes = append(extension.Attributes, backRepo.BackRepoAttribute.Map_AttributeDBID_AttributePtr[uint(_Attributeid)])
	}

	// This loop redeem extension.AttributeGroups in the stage from the encode in the back repo
	// It parses all AttributeGroupDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	extension.AttributeGroups = extension.AttributeGroups[:0]
	for _, _AttributeGroupid := range extensionDB.ExtensionPointersEncoding.AttributeGroups {
		extension.AttributeGroups = append(extension.AttributeGroups, backRepo.BackRepoAttributeGroup.Map_AttributeGroupDBID_AttributeGroupPtr[uint(_AttributeGroupid)])
	}

	return
}

// CommitExtension allows commit of a single extension (if already staged)
func (backRepo *BackRepoStruct) CommitExtension(extension *models.Extension) {
	backRepo.BackRepoExtension.CommitPhaseOneInstance(extension)
	if id, ok := backRepo.BackRepoExtension.Map_ExtensionPtr_ExtensionDBID[extension]; ok {
		backRepo.BackRepoExtension.CommitPhaseTwoInstance(backRepo, id, extension)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitExtension allows checkout of a single extension (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutExtension(extension *models.Extension) {
	// check if the extension is staged
	if _, ok := backRepo.BackRepoExtension.Map_ExtensionPtr_ExtensionDBID[extension]; ok {

		if id, ok := backRepo.BackRepoExtension.Map_ExtensionPtr_ExtensionDBID[extension]; ok {
			var extensionDB ExtensionDB
			extensionDB.ID = id

			if _, err := backRepo.BackRepoExtension.db.First(&extensionDB, id); err != nil {
				log.Fatalln("CheckoutExtension : Problem with getting object with id:", id)
			}
			backRepo.BackRepoExtension.CheckoutPhaseOneInstance(&extensionDB)
			backRepo.BackRepoExtension.CheckoutPhaseTwoInstance(backRepo, &extensionDB)
		}
	}
}

// CopyBasicFieldsFromExtension
func (extensionDB *ExtensionDB) CopyBasicFieldsFromExtension(extension *models.Extension) {
	// insertion point for fields commit

	extensionDB.Name_Data.String = extension.Name
	extensionDB.Name_Data.Valid = true

	extensionDB.OuterElementName_Data.String = extension.OuterElementName
	extensionDB.OuterElementName_Data.Valid = true

	extensionDB.Order_Data.Int64 = int64(extension.Order)
	extensionDB.Order_Data.Valid = true

	extensionDB.Depth_Data.Int64 = int64(extension.Depth)
	extensionDB.Depth_Data.Valid = true

	extensionDB.MinOccurs_Data.String = extension.MinOccurs
	extensionDB.MinOccurs_Data.Valid = true

	extensionDB.MaxOccurs_Data.String = extension.MaxOccurs
	extensionDB.MaxOccurs_Data.Valid = true

	extensionDB.Base_Data.String = extension.Base
	extensionDB.Base_Data.Valid = true

	extensionDB.Ref_Data.String = extension.Ref
	extensionDB.Ref_Data.Valid = true
}

// CopyBasicFieldsFromExtension_WOP
func (extensionDB *ExtensionDB) CopyBasicFieldsFromExtension_WOP(extension *models.Extension_WOP) {
	// insertion point for fields commit

	extensionDB.Name_Data.String = extension.Name
	extensionDB.Name_Data.Valid = true

	extensionDB.OuterElementName_Data.String = extension.OuterElementName
	extensionDB.OuterElementName_Data.Valid = true

	extensionDB.Order_Data.Int64 = int64(extension.Order)
	extensionDB.Order_Data.Valid = true

	extensionDB.Depth_Data.Int64 = int64(extension.Depth)
	extensionDB.Depth_Data.Valid = true

	extensionDB.MinOccurs_Data.String = extension.MinOccurs
	extensionDB.MinOccurs_Data.Valid = true

	extensionDB.MaxOccurs_Data.String = extension.MaxOccurs
	extensionDB.MaxOccurs_Data.Valid = true

	extensionDB.Base_Data.String = extension.Base
	extensionDB.Base_Data.Valid = true

	extensionDB.Ref_Data.String = extension.Ref
	extensionDB.Ref_Data.Valid = true
}

// CopyBasicFieldsFromExtensionWOP
func (extensionDB *ExtensionDB) CopyBasicFieldsFromExtensionWOP(extension *ExtensionWOP) {
	// insertion point for fields commit

	extensionDB.Name_Data.String = extension.Name
	extensionDB.Name_Data.Valid = true

	extensionDB.OuterElementName_Data.String = extension.OuterElementName
	extensionDB.OuterElementName_Data.Valid = true

	extensionDB.Order_Data.Int64 = int64(extension.Order)
	extensionDB.Order_Data.Valid = true

	extensionDB.Depth_Data.Int64 = int64(extension.Depth)
	extensionDB.Depth_Data.Valid = true

	extensionDB.MinOccurs_Data.String = extension.MinOccurs
	extensionDB.MinOccurs_Data.Valid = true

	extensionDB.MaxOccurs_Data.String = extension.MaxOccurs
	extensionDB.MaxOccurs_Data.Valid = true

	extensionDB.Base_Data.String = extension.Base
	extensionDB.Base_Data.Valid = true

	extensionDB.Ref_Data.String = extension.Ref
	extensionDB.Ref_Data.Valid = true
}

// CopyBasicFieldsToExtension
func (extensionDB *ExtensionDB) CopyBasicFieldsToExtension(extension *models.Extension) {
	// insertion point for checkout of basic fields (back repo to stage)
	extension.Name = extensionDB.Name_Data.String
	extension.OuterElementName = extensionDB.OuterElementName_Data.String
	extension.Order = int(extensionDB.Order_Data.Int64)
	extension.Depth = int(extensionDB.Depth_Data.Int64)
	extension.MinOccurs = extensionDB.MinOccurs_Data.String
	extension.MaxOccurs = extensionDB.MaxOccurs_Data.String
	extension.Base = extensionDB.Base_Data.String
	extension.Ref = extensionDB.Ref_Data.String
}

// CopyBasicFieldsToExtension_WOP
func (extensionDB *ExtensionDB) CopyBasicFieldsToExtension_WOP(extension *models.Extension_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	extension.Name = extensionDB.Name_Data.String
	extension.OuterElementName = extensionDB.OuterElementName_Data.String
	extension.Order = int(extensionDB.Order_Data.Int64)
	extension.Depth = int(extensionDB.Depth_Data.Int64)
	extension.MinOccurs = extensionDB.MinOccurs_Data.String
	extension.MaxOccurs = extensionDB.MaxOccurs_Data.String
	extension.Base = extensionDB.Base_Data.String
	extension.Ref = extensionDB.Ref_Data.String
}

// CopyBasicFieldsToExtensionWOP
func (extensionDB *ExtensionDB) CopyBasicFieldsToExtensionWOP(extension *ExtensionWOP) {
	extension.ID = int(extensionDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	extension.Name = extensionDB.Name_Data.String
	extension.OuterElementName = extensionDB.OuterElementName_Data.String
	extension.Order = int(extensionDB.Order_Data.Int64)
	extension.Depth = int(extensionDB.Depth_Data.Int64)
	extension.MinOccurs = extensionDB.MinOccurs_Data.String
	extension.MaxOccurs = extensionDB.MaxOccurs_Data.String
	extension.Base = extensionDB.Base_Data.String
	extension.Ref = extensionDB.Ref_Data.String
}

// Backup generates a json file from a slice of all ExtensionDB instances in the backrepo
func (backRepoExtension *BackRepoExtensionStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "ExtensionDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*ExtensionDB, 0)
	for _, extensionDB := range backRepoExtension.Map_ExtensionDBID_ExtensionDB {
		forBackup = append(forBackup, extensionDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json Extension ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json Extension file", err.Error())
	}
}

// Backup generates a json file from a slice of all ExtensionDB instances in the backrepo
func (backRepoExtension *BackRepoExtensionStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*ExtensionDB, 0)
	for _, extensionDB := range backRepoExtension.Map_ExtensionDBID_ExtensionDB {
		forBackup = append(forBackup, extensionDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("Extension")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&Extension_Fields, -1)
	for _, extensionDB := range forBackup {

		var extensionWOP ExtensionWOP
		extensionDB.CopyBasicFieldsToExtensionWOP(&extensionWOP)

		row := sh.AddRow()
		row.WriteStruct(&extensionWOP, -1)
	}
}

// RestoreXL from the "Extension" sheet all ExtensionDB instances
func (backRepoExtension *BackRepoExtensionStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoExtensionid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["Extension"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoExtension.rowVisitorExtension)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoExtension *BackRepoExtensionStruct) rowVisitorExtension(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var extensionWOP ExtensionWOP
		row.ReadStruct(&extensionWOP)

		// add the unmarshalled struct to the stage
		extensionDB := new(ExtensionDB)
		extensionDB.CopyBasicFieldsFromExtensionWOP(&extensionWOP)

		extensionDB_ID_atBackupTime := extensionDB.ID
		extensionDB.ID = 0
		_, err := backRepoExtension.db.Create(extensionDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoExtension.Map_ExtensionDBID_ExtensionDB[extensionDB.ID] = extensionDB
		BackRepoExtensionid_atBckpTime_newID[extensionDB_ID_atBackupTime] = extensionDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "ExtensionDB.json" in dirPath that stores an array
// of ExtensionDB and stores it in the database
// the map BackRepoExtensionid_atBckpTime_newID is updated accordingly
func (backRepoExtension *BackRepoExtensionStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoExtensionid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "ExtensionDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json Extension file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*ExtensionDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_ExtensionDBID_ExtensionDB
	for _, extensionDB := range forRestore {

		extensionDB_ID_atBackupTime := extensionDB.ID
		extensionDB.ID = 0
		_, err := backRepoExtension.db.Create(extensionDB)
		if err != nil {
			log.Fatal(err)
		}
		backRepoExtension.Map_ExtensionDBID_ExtensionDB[extensionDB.ID] = extensionDB
		BackRepoExtensionid_atBckpTime_newID[extensionDB_ID_atBackupTime] = extensionDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json Extension file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<Extension>id_atBckpTime_newID
// to compute new index
func (backRepoExtension *BackRepoExtensionStruct) RestorePhaseTwo() {

	for _, extensionDB := range backRepoExtension.Map_ExtensionDBID_ExtensionDB {

		// next line of code is to avert unused variable compilation error
		_ = extensionDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		db, _ := backRepoExtension.db.Model(extensionDB)
		_, err := db.Updates(*extensionDB)
		if err != nil {
			log.Fatal(err)
		}
	}

}

// BackRepoExtension.ResetReversePointers commits all staged instances of Extension to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoExtension *BackRepoExtensionStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, extension := range backRepoExtension.Map_ExtensionDBID_ExtensionPtr {
		backRepoExtension.ResetReversePointersInstance(backRepo, idx, extension)
	}

	return
}

func (backRepoExtension *BackRepoExtensionStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, extension *models.Extension) (Error error) {

	// fetch matching extensionDB
	if extensionDB, ok := backRepoExtension.Map_ExtensionDBID_ExtensionDB[idx]; ok {
		_ = extensionDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoExtensionid_atBckpTime_newID map[uint]uint
