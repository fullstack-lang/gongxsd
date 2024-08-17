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

	"github.com/fullstack-lang/gongxsd/alt/go/models"

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

	BackRepoSchema BackRepoSchemaStruct

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
		&SchemaDB{},
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
	backRepo.BackRepoSchema = BackRepoSchemaStruct{
		Map_SchemaDBID_SchemaPtr: make(map[uint]*models.Schema, 0),
		Map_SchemaDBID_SchemaDB:  make(map[uint]*SchemaDB, 0),
		Map_SchemaPtr_SchemaDBID: make(map[*models.Schema]uint, 0),

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
	backRepo.BackRepoSchema.CommitPhaseOne(stage)

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoAnnotation.CommitPhaseTwo(backRepo)
	backRepo.BackRepoComplexType.CommitPhaseTwo(backRepo)
	backRepo.BackRepoDocumentation.CommitPhaseTwo(backRepo)
	backRepo.BackRepoSchema.CommitPhaseTwo(backRepo)

	backRepo.IncrementCommitFromBackNb()
}

// Checkout the database into the stage
func (backRepo *BackRepoStruct) Checkout(stage *models.StageStruct) {
	// insertion point for per struct back repo phase one commit
	backRepo.BackRepoAnnotation.CheckoutPhaseOne()
	backRepo.BackRepoComplexType.CheckoutPhaseOne()
	backRepo.BackRepoDocumentation.CheckoutPhaseOne()
	backRepo.BackRepoSchema.CheckoutPhaseOne()

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoAnnotation.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoComplexType.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoDocumentation.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoSchema.CheckoutPhaseTwo(backRepo)
}

// Backup the BackRepoStruct
func (backRepo *BackRepoStruct) Backup(stage *models.StageStruct, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// insertion point for per struct backup
	backRepo.BackRepoAnnotation.Backup(dirPath)
	backRepo.BackRepoComplexType.Backup(dirPath)
	backRepo.BackRepoDocumentation.Backup(dirPath)
	backRepo.BackRepoSchema.Backup(dirPath)
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
	backRepo.BackRepoSchema.BackupXL(file)

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
	backRepo.BackRepoSchema.RestorePhaseOne(dirPath)

	//
	// restauration second phase (reindex pointers with the new ID)
	//

	// insertion point for per struct backup
	backRepo.BackRepoAnnotation.RestorePhaseTwo()
	backRepo.BackRepoComplexType.RestorePhaseTwo()
	backRepo.BackRepoDocumentation.RestorePhaseTwo()
	backRepo.BackRepoSchema.RestorePhaseTwo()

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
	backRepo.BackRepoSchema.RestoreXLPhaseOne(file)

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
