// do not modify, generated file
package orm

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/fullstack-lang/gongxsd/go/models"

	"github.com/tealeg/xlsx/v3"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// BackRepoStruct supports callback functions
type BackRepoStruct struct {
	// insertion point for per struct back repo declarations
	BackRepoAnnotation BackRepoAnnotationStruct

	BackRepoComplexType BackRepoComplexTypeStruct

	BackRepoDocumentation BackRepoDocumentationStruct

	BackRepoElement BackRepoElementStruct

	BackRepoEnumeration BackRepoEnumerationStruct

	BackRepoMaxInclusive BackRepoMaxInclusiveStruct

	BackRepoMinInclusive BackRepoMinInclusiveStruct

	BackRepoPattern BackRepoPatternStruct

	BackRepoRestriction BackRepoRestrictionStruct

	BackRepoSchema BackRepoSchemaStruct

	BackRepoSequence BackRepoSequenceStruct

	BackRepoSimpleType BackRepoSimpleTypeStruct

	CommitFromBackNb uint // records commit increments when performed by the back

	PushFromFrontNb uint // records commit increments when performed by the front

	stage *models.StageStruct

	// the back repo can broadcast the CommitFromBackNb to all interested subscribers
	rwMutex     sync.RWMutex
	subscribers []chan int
}

func NewBackRepo(stage *models.StageStruct, filename string) (backRepo *BackRepoStruct) {

	// adjust naming strategy to the stack
	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "github_com_fullstack_lang_gong_test_go_", // table name prefix
		},
	}
	db, err := gorm.Open(sqlite.Open(filename), gormConfig)

	// since testsim is a multi threaded application. It is important to set up
	// only one open connexion at a time
	dbDB_inMemory, err := db.DB()
	if err != nil {
		panic("cannot access DB of db" + err.Error())
	}
	// it is mandatory to allow parallel access, otherwise, bizarre errors occurs
	dbDB_inMemory.SetMaxOpenConns(1)

	if err != nil {
		panic("Failed to connect to database!")
	}

	// adjust naming strategy to the stack
	db.Config.NamingStrategy = &schema.NamingStrategy{
		TablePrefix: "github_com_fullstack_lang_gong_test_go_", // table name prefix
	}

	err = db.AutoMigrate( // insertion point for reference to structs
		&AnnotationDB{},
		&ComplexTypeDB{},
		&DocumentationDB{},
		&ElementDB{},
		&EnumerationDB{},
		&MaxInclusiveDB{},
		&MinInclusiveDB{},
		&PatternDB{},
		&RestrictionDB{},
		&SchemaDB{},
		&SequenceDB{},
		&SimpleTypeDB{},
	)

	if err != nil {
		msg := err.Error()
		panic("problem with migration " + msg + " on package github.com/fullstack-lang/gong/test/go")
	}

	backRepo = new(BackRepoStruct)

	// insertion point for per struct back repo declarations
	backRepo.BackRepoAnnotation = BackRepoAnnotationStruct{
		Map_AnnotationDBID_AnnotationPtr: make(map[uint]*models.Annotation, 0),
		Map_AnnotationDBID_AnnotationDB:  make(map[uint]*AnnotationDB, 0),
		Map_AnnotationPtr_AnnotationDBID: make(map[*models.Annotation]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoComplexType = BackRepoComplexTypeStruct{
		Map_ComplexTypeDBID_ComplexTypePtr: make(map[uint]*models.ComplexType, 0),
		Map_ComplexTypeDBID_ComplexTypeDB:  make(map[uint]*ComplexTypeDB, 0),
		Map_ComplexTypePtr_ComplexTypeDBID: make(map[*models.ComplexType]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoDocumentation = BackRepoDocumentationStruct{
		Map_DocumentationDBID_DocumentationPtr: make(map[uint]*models.Documentation, 0),
		Map_DocumentationDBID_DocumentationDB:  make(map[uint]*DocumentationDB, 0),
		Map_DocumentationPtr_DocumentationDBID: make(map[*models.Documentation]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoElement = BackRepoElementStruct{
		Map_ElementDBID_ElementPtr: make(map[uint]*models.Element, 0),
		Map_ElementDBID_ElementDB:  make(map[uint]*ElementDB, 0),
		Map_ElementPtr_ElementDBID: make(map[*models.Element]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoEnumeration = BackRepoEnumerationStruct{
		Map_EnumerationDBID_EnumerationPtr: make(map[uint]*models.Enumeration, 0),
		Map_EnumerationDBID_EnumerationDB:  make(map[uint]*EnumerationDB, 0),
		Map_EnumerationPtr_EnumerationDBID: make(map[*models.Enumeration]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoMaxInclusive = BackRepoMaxInclusiveStruct{
		Map_MaxInclusiveDBID_MaxInclusivePtr: make(map[uint]*models.MaxInclusive, 0),
		Map_MaxInclusiveDBID_MaxInclusiveDB:  make(map[uint]*MaxInclusiveDB, 0),
		Map_MaxInclusivePtr_MaxInclusiveDBID: make(map[*models.MaxInclusive]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoMinInclusive = BackRepoMinInclusiveStruct{
		Map_MinInclusiveDBID_MinInclusivePtr: make(map[uint]*models.MinInclusive, 0),
		Map_MinInclusiveDBID_MinInclusiveDB:  make(map[uint]*MinInclusiveDB, 0),
		Map_MinInclusivePtr_MinInclusiveDBID: make(map[*models.MinInclusive]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoPattern = BackRepoPatternStruct{
		Map_PatternDBID_PatternPtr: make(map[uint]*models.Pattern, 0),
		Map_PatternDBID_PatternDB:  make(map[uint]*PatternDB, 0),
		Map_PatternPtr_PatternDBID: make(map[*models.Pattern]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoRestriction = BackRepoRestrictionStruct{
		Map_RestrictionDBID_RestrictionPtr: make(map[uint]*models.Restriction, 0),
		Map_RestrictionDBID_RestrictionDB:  make(map[uint]*RestrictionDB, 0),
		Map_RestrictionPtr_RestrictionDBID: make(map[*models.Restriction]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoSchema = BackRepoSchemaStruct{
		Map_SchemaDBID_SchemaPtr: make(map[uint]*models.Schema, 0),
		Map_SchemaDBID_SchemaDB:  make(map[uint]*SchemaDB, 0),
		Map_SchemaPtr_SchemaDBID: make(map[*models.Schema]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoSequence = BackRepoSequenceStruct{
		Map_SequenceDBID_SequencePtr: make(map[uint]*models.Sequence, 0),
		Map_SequenceDBID_SequenceDB:  make(map[uint]*SequenceDB, 0),
		Map_SequencePtr_SequenceDBID: make(map[*models.Sequence]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoSimpleType = BackRepoSimpleTypeStruct{
		Map_SimpleTypeDBID_SimpleTypePtr: make(map[uint]*models.SimpleType, 0),
		Map_SimpleTypeDBID_SimpleTypeDB:  make(map[uint]*SimpleTypeDB, 0),
		Map_SimpleTypePtr_SimpleTypeDBID: make(map[*models.SimpleType]uint, 0),

		db:    db,
		stage: stage,
	}

	stage.BackRepo = backRepo
	backRepo.stage = stage

	return
}

func (backRepo *BackRepoStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepo.stage
	return
}

func (backRepo *BackRepoStruct) GetLastCommitFromBackNb() uint {
	return backRepo.CommitFromBackNb
}

func (backRepo *BackRepoStruct) GetLastPushFromFrontNb() uint {
	return backRepo.PushFromFrontNb
}

func (backRepo *BackRepoStruct) IncrementCommitFromBackNb() uint {
	if backRepo.stage.OnInitCommitCallback != nil {
		backRepo.stage.OnInitCommitCallback.BeforeCommit(backRepo.stage)
	}
	if backRepo.stage.OnInitCommitFromBackCallback != nil {
		backRepo.stage.OnInitCommitFromBackCallback.BeforeCommit(backRepo.stage)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1

	backRepo.broadcastNbCommitToBack()
	
	return backRepo.CommitFromBackNb
}

func (backRepo *BackRepoStruct) IncrementPushFromFrontNb() uint {
	if backRepo.stage.OnInitCommitCallback != nil {
		backRepo.stage.OnInitCommitCallback.BeforeCommit(backRepo.stage)
	}
	if backRepo.stage.OnInitCommitFromFrontCallback != nil {
		backRepo.stage.OnInitCommitFromFrontCallback.BeforeCommit(backRepo.stage)
	}
	backRepo.PushFromFrontNb = backRepo.PushFromFrontNb + 1
	return backRepo.CommitFromBackNb
}

// Commit the BackRepoStruct inner variables and link to the database
func (backRepo *BackRepoStruct) Commit(stage *models.StageStruct) {
	// insertion point for per struct back repo phase one commit
	backRepo.BackRepoAnnotation.CommitPhaseOne(stage)
	backRepo.BackRepoComplexType.CommitPhaseOne(stage)
	backRepo.BackRepoDocumentation.CommitPhaseOne(stage)
	backRepo.BackRepoElement.CommitPhaseOne(stage)
	backRepo.BackRepoEnumeration.CommitPhaseOne(stage)
	backRepo.BackRepoMaxInclusive.CommitPhaseOne(stage)
	backRepo.BackRepoMinInclusive.CommitPhaseOne(stage)
	backRepo.BackRepoPattern.CommitPhaseOne(stage)
	backRepo.BackRepoRestriction.CommitPhaseOne(stage)
	backRepo.BackRepoSchema.CommitPhaseOne(stage)
	backRepo.BackRepoSequence.CommitPhaseOne(stage)
	backRepo.BackRepoSimpleType.CommitPhaseOne(stage)

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoAnnotation.CommitPhaseTwo(backRepo)
	backRepo.BackRepoComplexType.CommitPhaseTwo(backRepo)
	backRepo.BackRepoDocumentation.CommitPhaseTwo(backRepo)
	backRepo.BackRepoElement.CommitPhaseTwo(backRepo)
	backRepo.BackRepoEnumeration.CommitPhaseTwo(backRepo)
	backRepo.BackRepoMaxInclusive.CommitPhaseTwo(backRepo)
	backRepo.BackRepoMinInclusive.CommitPhaseTwo(backRepo)
	backRepo.BackRepoPattern.CommitPhaseTwo(backRepo)
	backRepo.BackRepoRestriction.CommitPhaseTwo(backRepo)
	backRepo.BackRepoSchema.CommitPhaseTwo(backRepo)
	backRepo.BackRepoSequence.CommitPhaseTwo(backRepo)
	backRepo.BackRepoSimpleType.CommitPhaseTwo(backRepo)

	backRepo.IncrementCommitFromBackNb()
}

// Checkout the database into the stage
func (backRepo *BackRepoStruct) Checkout(stage *models.StageStruct) {
	// insertion point for per struct back repo phase one commit
	backRepo.BackRepoAnnotation.CheckoutPhaseOne()
	backRepo.BackRepoComplexType.CheckoutPhaseOne()
	backRepo.BackRepoDocumentation.CheckoutPhaseOne()
	backRepo.BackRepoElement.CheckoutPhaseOne()
	backRepo.BackRepoEnumeration.CheckoutPhaseOne()
	backRepo.BackRepoMaxInclusive.CheckoutPhaseOne()
	backRepo.BackRepoMinInclusive.CheckoutPhaseOne()
	backRepo.BackRepoPattern.CheckoutPhaseOne()
	backRepo.BackRepoRestriction.CheckoutPhaseOne()
	backRepo.BackRepoSchema.CheckoutPhaseOne()
	backRepo.BackRepoSequence.CheckoutPhaseOne()
	backRepo.BackRepoSimpleType.CheckoutPhaseOne()

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoAnnotation.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoComplexType.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoDocumentation.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoElement.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoEnumeration.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoMaxInclusive.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoMinInclusive.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoPattern.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoRestriction.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoSchema.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoSequence.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoSimpleType.CheckoutPhaseTwo(backRepo)
}

// Backup the BackRepoStruct
func (backRepo *BackRepoStruct) Backup(stage *models.StageStruct, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// insertion point for per struct backup
	backRepo.BackRepoAnnotation.Backup(dirPath)
	backRepo.BackRepoComplexType.Backup(dirPath)
	backRepo.BackRepoDocumentation.Backup(dirPath)
	backRepo.BackRepoElement.Backup(dirPath)
	backRepo.BackRepoEnumeration.Backup(dirPath)
	backRepo.BackRepoMaxInclusive.Backup(dirPath)
	backRepo.BackRepoMinInclusive.Backup(dirPath)
	backRepo.BackRepoPattern.Backup(dirPath)
	backRepo.BackRepoRestriction.Backup(dirPath)
	backRepo.BackRepoSchema.Backup(dirPath)
	backRepo.BackRepoSequence.Backup(dirPath)
	backRepo.BackRepoSimpleType.Backup(dirPath)
}

// Backup in XL the BackRepoStruct
func (backRepo *BackRepoStruct) BackupXL(stage *models.StageStruct, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// open an existing file
	file := xlsx.NewFile()

	// insertion point for per struct backup
	backRepo.BackRepoAnnotation.BackupXL(file)
	backRepo.BackRepoComplexType.BackupXL(file)
	backRepo.BackRepoDocumentation.BackupXL(file)
	backRepo.BackRepoElement.BackupXL(file)
	backRepo.BackRepoEnumeration.BackupXL(file)
	backRepo.BackRepoMaxInclusive.BackupXL(file)
	backRepo.BackRepoMinInclusive.BackupXL(file)
	backRepo.BackRepoPattern.BackupXL(file)
	backRepo.BackRepoRestriction.BackupXL(file)
	backRepo.BackRepoSchema.BackupXL(file)
	backRepo.BackRepoSequence.BackupXL(file)
	backRepo.BackRepoSimpleType.BackupXL(file)

	var b bytes.Buffer
	writer := bufio.NewWriter(&b)
	file.Write(writer)
	theBytes := b.Bytes()

	filename := filepath.Join(dirPath, "bckp.xlsx")
	err := ioutil.WriteFile(filename, theBytes, 0644)
	if err != nil {
		log.Panic("Cannot write the XL file", err.Error())
	}
}

// Restore the database into the back repo
func (backRepo *BackRepoStruct) Restore(stage *models.StageStruct, dirPath string) {
	backRepo.stage.Commit()
	backRepo.stage.Reset()
	backRepo.stage.Checkout()

	//
	// restauration first phase (create DB instance with new IDs)
	//

	// insertion point for per struct backup
	backRepo.BackRepoAnnotation.RestorePhaseOne(dirPath)
	backRepo.BackRepoComplexType.RestorePhaseOne(dirPath)
	backRepo.BackRepoDocumentation.RestorePhaseOne(dirPath)
	backRepo.BackRepoElement.RestorePhaseOne(dirPath)
	backRepo.BackRepoEnumeration.RestorePhaseOne(dirPath)
	backRepo.BackRepoMaxInclusive.RestorePhaseOne(dirPath)
	backRepo.BackRepoMinInclusive.RestorePhaseOne(dirPath)
	backRepo.BackRepoPattern.RestorePhaseOne(dirPath)
	backRepo.BackRepoRestriction.RestorePhaseOne(dirPath)
	backRepo.BackRepoSchema.RestorePhaseOne(dirPath)
	backRepo.BackRepoSequence.RestorePhaseOne(dirPath)
	backRepo.BackRepoSimpleType.RestorePhaseOne(dirPath)

	//
	// restauration second phase (reindex pointers with the new ID)
	//

	// insertion point for per struct backup
	backRepo.BackRepoAnnotation.RestorePhaseTwo()
	backRepo.BackRepoComplexType.RestorePhaseTwo()
	backRepo.BackRepoDocumentation.RestorePhaseTwo()
	backRepo.BackRepoElement.RestorePhaseTwo()
	backRepo.BackRepoEnumeration.RestorePhaseTwo()
	backRepo.BackRepoMaxInclusive.RestorePhaseTwo()
	backRepo.BackRepoMinInclusive.RestorePhaseTwo()
	backRepo.BackRepoPattern.RestorePhaseTwo()
	backRepo.BackRepoRestriction.RestorePhaseTwo()
	backRepo.BackRepoSchema.RestorePhaseTwo()
	backRepo.BackRepoSequence.RestorePhaseTwo()
	backRepo.BackRepoSimpleType.RestorePhaseTwo()

	backRepo.stage.Checkout()
}

// Restore the database into the back repo
func (backRepo *BackRepoStruct) RestoreXL(stage *models.StageStruct, dirPath string) {

	// clean the stage
	backRepo.stage.Reset()

	// commit the cleaned stage
	backRepo.stage.Commit()

	// open an existing file
	filename := filepath.Join(dirPath, "bckp.xlsx")
	file, err := xlsx.OpenFile(filename)
	_ = file

	if err != nil {
		log.Panic("Cannot read the XL file", err.Error())
	}

	//
	// restauration first phase (create DB instance with new IDs)
	//

	// insertion point for per struct backup
	backRepo.BackRepoAnnotation.RestoreXLPhaseOne(file)
	backRepo.BackRepoComplexType.RestoreXLPhaseOne(file)
	backRepo.BackRepoDocumentation.RestoreXLPhaseOne(file)
	backRepo.BackRepoElement.RestoreXLPhaseOne(file)
	backRepo.BackRepoEnumeration.RestoreXLPhaseOne(file)
	backRepo.BackRepoMaxInclusive.RestoreXLPhaseOne(file)
	backRepo.BackRepoMinInclusive.RestoreXLPhaseOne(file)
	backRepo.BackRepoPattern.RestoreXLPhaseOne(file)
	backRepo.BackRepoRestriction.RestoreXLPhaseOne(file)
	backRepo.BackRepoSchema.RestoreXLPhaseOne(file)
	backRepo.BackRepoSequence.RestoreXLPhaseOne(file)
	backRepo.BackRepoSimpleType.RestoreXLPhaseOne(file)

	// commit the restored stage
	backRepo.stage.Commit()
}

func (backRepoStruct *BackRepoStruct) SubscribeToCommitNb() <-chan int {
	backRepoStruct.rwMutex.Lock()
	defer backRepoStruct.rwMutex.Unlock()

	ch := make(chan int)
	backRepoStruct.subscribers = append(backRepoStruct.subscribers, ch)
	return ch
}

func (backRepoStruct *BackRepoStruct) broadcastNbCommitToBack() {
	backRepoStruct.rwMutex.RLock()
	defer backRepoStruct.rwMutex.RUnlock()

	activeChannels := make([]chan int, 0)

	for _, ch := range backRepoStruct.subscribers {
		select {
		case ch <- int(backRepoStruct.CommitFromBackNb):
			activeChannels = append(activeChannels, ch)
		default:
			// Assume channel is no longer active; don't add to activeChannels
			log.Println("Channel no longer active")
			close(ch)
		}
	}
	backRepoStruct.subscribers = activeChannels
}
