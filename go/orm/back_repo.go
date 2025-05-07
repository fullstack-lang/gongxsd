// do not modify, generated file
package orm

import (
	"bufio"
	"bytes"
	"context"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/fullstack-lang/gongxsd/go/db"
	"github.com/fullstack-lang/gongxsd/go/models"

	/* THIS IS REMOVED BY GONG COMPILER IF TARGET IS gorm
	"github.com/fullstack-lang/gongxsd/go/orm/dbgorm"
	THIS IS REMOVED BY GONG COMPILER IF TARGET IS gorm */

	"github.com/tealeg/xlsx/v3"
)

// BackRepoStruct supports callback functions
type BackRepoStruct struct {
	// insertion point for per struct back repo declarations
	BackRepoAll BackRepoAllStruct

	BackRepoAnnotation BackRepoAnnotationStruct

	BackRepoAttribute BackRepoAttributeStruct

	BackRepoAttributeGroup BackRepoAttributeGroupStruct

	BackRepoChoice BackRepoChoiceStruct

	BackRepoComplexContent BackRepoComplexContentStruct

	BackRepoComplexType BackRepoComplexTypeStruct

	BackRepoDocumentation BackRepoDocumentationStruct

	BackRepoElement BackRepoElementStruct

	BackRepoEnumeration BackRepoEnumerationStruct

	BackRepoExtension BackRepoExtensionStruct

	BackRepoGroup BackRepoGroupStruct

	BackRepoLength BackRepoLengthStruct

	BackRepoMaxInclusive BackRepoMaxInclusiveStruct

	BackRepoMaxLength BackRepoMaxLengthStruct

	BackRepoMinInclusive BackRepoMinInclusiveStruct

	BackRepoMinLength BackRepoMinLengthStruct

	BackRepoPattern BackRepoPatternStruct

	BackRepoRestriction BackRepoRestrictionStruct

	BackRepoSchema BackRepoSchemaStruct

	BackRepoSequence BackRepoSequenceStruct

	BackRepoSimpleContent BackRepoSimpleContentStruct

	BackRepoSimpleType BackRepoSimpleTypeStruct

	BackRepoTotalDigit BackRepoTotalDigitStruct

	BackRepoUnion BackRepoUnionStruct

	BackRepoWhiteSpace BackRepoWhiteSpaceStruct

	CommitFromBackNb uint // records commit increments when performed by the back

	PushFromFrontNb uint // records commit increments when performed by the front

	stage *models.Stage

	// the back repo can broadcast the CommitFromBackNb to all interested subscribers
	rwMutex     sync.RWMutex

	subscribersRwMutex sync.RWMutex
	subscribers []chan int
}

func NewBackRepo(stage *models.Stage, filename string) (backRepo *BackRepoStruct) {

	var db db.DBInterface

	db = NewDBLite()

	/* THIS IS REMOVED BY GONG COMPILER IF TARGET IS gorm
	db = dbgorm.NewDBWrapper(filename, "github_com_fullstack_lang_gongxsd_go",
		&AllDB{},
		&AnnotationDB{},
		&AttributeDB{},
		&AttributeGroupDB{},
		&ChoiceDB{},
		&ComplexContentDB{},
		&ComplexTypeDB{},
		&DocumentationDB{},
		&ElementDB{},
		&EnumerationDB{},
		&ExtensionDB{},
		&GroupDB{},
		&LengthDB{},
		&MaxInclusiveDB{},
		&MaxLengthDB{},
		&MinInclusiveDB{},
		&MinLengthDB{},
		&PatternDB{},
		&RestrictionDB{},
		&SchemaDB{},
		&SequenceDB{},
		&SimpleContentDB{},
		&SimpleTypeDB{},
		&TotalDigitDB{},
		&UnionDB{},
		&WhiteSpaceDB{},
	)
	THIS IS REMOVED BY GONG COMPILER IF TARGET IS gorm */

	backRepo = new(BackRepoStruct)

	// insertion point for per struct back repo declarations
	backRepo.BackRepoAll = BackRepoAllStruct{
		Map_AllDBID_AllPtr: make(map[uint]*models.All, 0),
		Map_AllDBID_AllDB:  make(map[uint]*AllDB, 0),
		Map_AllPtr_AllDBID: make(map[*models.All]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoAnnotation = BackRepoAnnotationStruct{
		Map_AnnotationDBID_AnnotationPtr: make(map[uint]*models.Annotation, 0),
		Map_AnnotationDBID_AnnotationDB:  make(map[uint]*AnnotationDB, 0),
		Map_AnnotationPtr_AnnotationDBID: make(map[*models.Annotation]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoAttribute = BackRepoAttributeStruct{
		Map_AttributeDBID_AttributePtr: make(map[uint]*models.Attribute, 0),
		Map_AttributeDBID_AttributeDB:  make(map[uint]*AttributeDB, 0),
		Map_AttributePtr_AttributeDBID: make(map[*models.Attribute]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoAttributeGroup = BackRepoAttributeGroupStruct{
		Map_AttributeGroupDBID_AttributeGroupPtr: make(map[uint]*models.AttributeGroup, 0),
		Map_AttributeGroupDBID_AttributeGroupDB:  make(map[uint]*AttributeGroupDB, 0),
		Map_AttributeGroupPtr_AttributeGroupDBID: make(map[*models.AttributeGroup]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoChoice = BackRepoChoiceStruct{
		Map_ChoiceDBID_ChoicePtr: make(map[uint]*models.Choice, 0),
		Map_ChoiceDBID_ChoiceDB:  make(map[uint]*ChoiceDB, 0),
		Map_ChoicePtr_ChoiceDBID: make(map[*models.Choice]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoComplexContent = BackRepoComplexContentStruct{
		Map_ComplexContentDBID_ComplexContentPtr: make(map[uint]*models.ComplexContent, 0),
		Map_ComplexContentDBID_ComplexContentDB:  make(map[uint]*ComplexContentDB, 0),
		Map_ComplexContentPtr_ComplexContentDBID: make(map[*models.ComplexContent]uint, 0),

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
	backRepo.BackRepoExtension = BackRepoExtensionStruct{
		Map_ExtensionDBID_ExtensionPtr: make(map[uint]*models.Extension, 0),
		Map_ExtensionDBID_ExtensionDB:  make(map[uint]*ExtensionDB, 0),
		Map_ExtensionPtr_ExtensionDBID: make(map[*models.Extension]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoGroup = BackRepoGroupStruct{
		Map_GroupDBID_GroupPtr: make(map[uint]*models.Group, 0),
		Map_GroupDBID_GroupDB:  make(map[uint]*GroupDB, 0),
		Map_GroupPtr_GroupDBID: make(map[*models.Group]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoLength = BackRepoLengthStruct{
		Map_LengthDBID_LengthPtr: make(map[uint]*models.Length, 0),
		Map_LengthDBID_LengthDB:  make(map[uint]*LengthDB, 0),
		Map_LengthPtr_LengthDBID: make(map[*models.Length]uint, 0),

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
	backRepo.BackRepoMaxLength = BackRepoMaxLengthStruct{
		Map_MaxLengthDBID_MaxLengthPtr: make(map[uint]*models.MaxLength, 0),
		Map_MaxLengthDBID_MaxLengthDB:  make(map[uint]*MaxLengthDB, 0),
		Map_MaxLengthPtr_MaxLengthDBID: make(map[*models.MaxLength]uint, 0),

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
	backRepo.BackRepoMinLength = BackRepoMinLengthStruct{
		Map_MinLengthDBID_MinLengthPtr: make(map[uint]*models.MinLength, 0),
		Map_MinLengthDBID_MinLengthDB:  make(map[uint]*MinLengthDB, 0),
		Map_MinLengthPtr_MinLengthDBID: make(map[*models.MinLength]uint, 0),

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
	backRepo.BackRepoSimpleContent = BackRepoSimpleContentStruct{
		Map_SimpleContentDBID_SimpleContentPtr: make(map[uint]*models.SimpleContent, 0),
		Map_SimpleContentDBID_SimpleContentDB:  make(map[uint]*SimpleContentDB, 0),
		Map_SimpleContentPtr_SimpleContentDBID: make(map[*models.SimpleContent]uint, 0),

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
	backRepo.BackRepoTotalDigit = BackRepoTotalDigitStruct{
		Map_TotalDigitDBID_TotalDigitPtr: make(map[uint]*models.TotalDigit, 0),
		Map_TotalDigitDBID_TotalDigitDB:  make(map[uint]*TotalDigitDB, 0),
		Map_TotalDigitPtr_TotalDigitDBID: make(map[*models.TotalDigit]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoUnion = BackRepoUnionStruct{
		Map_UnionDBID_UnionPtr: make(map[uint]*models.Union, 0),
		Map_UnionDBID_UnionDB:  make(map[uint]*UnionDB, 0),
		Map_UnionPtr_UnionDBID: make(map[*models.Union]uint, 0),

		db:    db,
		stage: stage,
	}
	backRepo.BackRepoWhiteSpace = BackRepoWhiteSpaceStruct{
		Map_WhiteSpaceDBID_WhiteSpacePtr: make(map[uint]*models.WhiteSpace, 0),
		Map_WhiteSpaceDBID_WhiteSpaceDB:  make(map[uint]*WhiteSpaceDB, 0),
		Map_WhiteSpacePtr_WhiteSpaceDBID: make(map[*models.WhiteSpace]uint, 0),

		db:    db,
		stage: stage,
	}

	stage.BackRepo = backRepo
	backRepo.stage = stage

	return
}

func (backRepo *BackRepoStruct) GetStage() (stage *models.Stage) {
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
func (backRepo *BackRepoStruct) Commit(stage *models.Stage) {

	// forbid read of back repo during commit
	backRepo.rwMutex.Lock()

	// insertion point for per struct back repo phase one commit
	backRepo.BackRepoAll.CommitPhaseOne(stage)
	backRepo.BackRepoAnnotation.CommitPhaseOne(stage)
	backRepo.BackRepoAttribute.CommitPhaseOne(stage)
	backRepo.BackRepoAttributeGroup.CommitPhaseOne(stage)
	backRepo.BackRepoChoice.CommitPhaseOne(stage)
	backRepo.BackRepoComplexContent.CommitPhaseOne(stage)
	backRepo.BackRepoComplexType.CommitPhaseOne(stage)
	backRepo.BackRepoDocumentation.CommitPhaseOne(stage)
	backRepo.BackRepoElement.CommitPhaseOne(stage)
	backRepo.BackRepoEnumeration.CommitPhaseOne(stage)
	backRepo.BackRepoExtension.CommitPhaseOne(stage)
	backRepo.BackRepoGroup.CommitPhaseOne(stage)
	backRepo.BackRepoLength.CommitPhaseOne(stage)
	backRepo.BackRepoMaxInclusive.CommitPhaseOne(stage)
	backRepo.BackRepoMaxLength.CommitPhaseOne(stage)
	backRepo.BackRepoMinInclusive.CommitPhaseOne(stage)
	backRepo.BackRepoMinLength.CommitPhaseOne(stage)
	backRepo.BackRepoPattern.CommitPhaseOne(stage)
	backRepo.BackRepoRestriction.CommitPhaseOne(stage)
	backRepo.BackRepoSchema.CommitPhaseOne(stage)
	backRepo.BackRepoSequence.CommitPhaseOne(stage)
	backRepo.BackRepoSimpleContent.CommitPhaseOne(stage)
	backRepo.BackRepoSimpleType.CommitPhaseOne(stage)
	backRepo.BackRepoTotalDigit.CommitPhaseOne(stage)
	backRepo.BackRepoUnion.CommitPhaseOne(stage)
	backRepo.BackRepoWhiteSpace.CommitPhaseOne(stage)

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoAll.CommitPhaseTwo(backRepo)
	backRepo.BackRepoAnnotation.CommitPhaseTwo(backRepo)
	backRepo.BackRepoAttribute.CommitPhaseTwo(backRepo)
	backRepo.BackRepoAttributeGroup.CommitPhaseTwo(backRepo)
	backRepo.BackRepoChoice.CommitPhaseTwo(backRepo)
	backRepo.BackRepoComplexContent.CommitPhaseTwo(backRepo)
	backRepo.BackRepoComplexType.CommitPhaseTwo(backRepo)
	backRepo.BackRepoDocumentation.CommitPhaseTwo(backRepo)
	backRepo.BackRepoElement.CommitPhaseTwo(backRepo)
	backRepo.BackRepoEnumeration.CommitPhaseTwo(backRepo)
	backRepo.BackRepoExtension.CommitPhaseTwo(backRepo)
	backRepo.BackRepoGroup.CommitPhaseTwo(backRepo)
	backRepo.BackRepoLength.CommitPhaseTwo(backRepo)
	backRepo.BackRepoMaxInclusive.CommitPhaseTwo(backRepo)
	backRepo.BackRepoMaxLength.CommitPhaseTwo(backRepo)
	backRepo.BackRepoMinInclusive.CommitPhaseTwo(backRepo)
	backRepo.BackRepoMinLength.CommitPhaseTwo(backRepo)
	backRepo.BackRepoPattern.CommitPhaseTwo(backRepo)
	backRepo.BackRepoRestriction.CommitPhaseTwo(backRepo)
	backRepo.BackRepoSchema.CommitPhaseTwo(backRepo)
	backRepo.BackRepoSequence.CommitPhaseTwo(backRepo)
	backRepo.BackRepoSimpleContent.CommitPhaseTwo(backRepo)
	backRepo.BackRepoSimpleType.CommitPhaseTwo(backRepo)
	backRepo.BackRepoTotalDigit.CommitPhaseTwo(backRepo)
	backRepo.BackRepoUnion.CommitPhaseTwo(backRepo)
	backRepo.BackRepoWhiteSpace.CommitPhaseTwo(backRepo)

	// important to release the mutex before calls to IncrementCommitFromBackNb
	// because it will block otherwise
	backRepo.rwMutex.Unlock()

	backRepo.IncrementCommitFromBackNb()
}

// Checkout the database into the stage
func (backRepo *BackRepoStruct) Checkout(stage *models.Stage) {

	backRepo.rwMutex.Lock()
	defer backRepo.rwMutex.Unlock()
	// insertion point for per struct back repo phase one commit
	backRepo.BackRepoAll.CheckoutPhaseOne()
	backRepo.BackRepoAnnotation.CheckoutPhaseOne()
	backRepo.BackRepoAttribute.CheckoutPhaseOne()
	backRepo.BackRepoAttributeGroup.CheckoutPhaseOne()
	backRepo.BackRepoChoice.CheckoutPhaseOne()
	backRepo.BackRepoComplexContent.CheckoutPhaseOne()
	backRepo.BackRepoComplexType.CheckoutPhaseOne()
	backRepo.BackRepoDocumentation.CheckoutPhaseOne()
	backRepo.BackRepoElement.CheckoutPhaseOne()
	backRepo.BackRepoEnumeration.CheckoutPhaseOne()
	backRepo.BackRepoExtension.CheckoutPhaseOne()
	backRepo.BackRepoGroup.CheckoutPhaseOne()
	backRepo.BackRepoLength.CheckoutPhaseOne()
	backRepo.BackRepoMaxInclusive.CheckoutPhaseOne()
	backRepo.BackRepoMaxLength.CheckoutPhaseOne()
	backRepo.BackRepoMinInclusive.CheckoutPhaseOne()
	backRepo.BackRepoMinLength.CheckoutPhaseOne()
	backRepo.BackRepoPattern.CheckoutPhaseOne()
	backRepo.BackRepoRestriction.CheckoutPhaseOne()
	backRepo.BackRepoSchema.CheckoutPhaseOne()
	backRepo.BackRepoSequence.CheckoutPhaseOne()
	backRepo.BackRepoSimpleContent.CheckoutPhaseOne()
	backRepo.BackRepoSimpleType.CheckoutPhaseOne()
	backRepo.BackRepoTotalDigit.CheckoutPhaseOne()
	backRepo.BackRepoUnion.CheckoutPhaseOne()
	backRepo.BackRepoWhiteSpace.CheckoutPhaseOne()

	// insertion point for per struct back repo phase two commit
	backRepo.BackRepoAll.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoAnnotation.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoAttribute.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoAttributeGroup.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoChoice.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoComplexContent.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoComplexType.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoDocumentation.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoElement.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoEnumeration.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoExtension.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoGroup.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoLength.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoMaxInclusive.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoMaxLength.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoMinInclusive.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoMinLength.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoPattern.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoRestriction.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoSchema.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoSequence.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoSimpleContent.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoSimpleType.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoTotalDigit.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoUnion.CheckoutPhaseTwo(backRepo)
	backRepo.BackRepoWhiteSpace.CheckoutPhaseTwo(backRepo)
}

// Backup the BackRepoStruct
func (backRepo *BackRepoStruct) Backup(stage *models.Stage, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// insertion point for per struct backup
	backRepo.BackRepoAll.Backup(dirPath)
	backRepo.BackRepoAnnotation.Backup(dirPath)
	backRepo.BackRepoAttribute.Backup(dirPath)
	backRepo.BackRepoAttributeGroup.Backup(dirPath)
	backRepo.BackRepoChoice.Backup(dirPath)
	backRepo.BackRepoComplexContent.Backup(dirPath)
	backRepo.BackRepoComplexType.Backup(dirPath)
	backRepo.BackRepoDocumentation.Backup(dirPath)
	backRepo.BackRepoElement.Backup(dirPath)
	backRepo.BackRepoEnumeration.Backup(dirPath)
	backRepo.BackRepoExtension.Backup(dirPath)
	backRepo.BackRepoGroup.Backup(dirPath)
	backRepo.BackRepoLength.Backup(dirPath)
	backRepo.BackRepoMaxInclusive.Backup(dirPath)
	backRepo.BackRepoMaxLength.Backup(dirPath)
	backRepo.BackRepoMinInclusive.Backup(dirPath)
	backRepo.BackRepoMinLength.Backup(dirPath)
	backRepo.BackRepoPattern.Backup(dirPath)
	backRepo.BackRepoRestriction.Backup(dirPath)
	backRepo.BackRepoSchema.Backup(dirPath)
	backRepo.BackRepoSequence.Backup(dirPath)
	backRepo.BackRepoSimpleContent.Backup(dirPath)
	backRepo.BackRepoSimpleType.Backup(dirPath)
	backRepo.BackRepoTotalDigit.Backup(dirPath)
	backRepo.BackRepoUnion.Backup(dirPath)
	backRepo.BackRepoWhiteSpace.Backup(dirPath)
}

// Backup in XL the BackRepoStruct
func (backRepo *BackRepoStruct) BackupXL(stage *models.Stage, dirPath string) {
	os.MkdirAll(dirPath, os.ModePerm)

	// open an existing file
	file := xlsx.NewFile()

	// insertion point for per struct backup
	backRepo.BackRepoAll.BackupXL(file)
	backRepo.BackRepoAnnotation.BackupXL(file)
	backRepo.BackRepoAttribute.BackupXL(file)
	backRepo.BackRepoAttributeGroup.BackupXL(file)
	backRepo.BackRepoChoice.BackupXL(file)
	backRepo.BackRepoComplexContent.BackupXL(file)
	backRepo.BackRepoComplexType.BackupXL(file)
	backRepo.BackRepoDocumentation.BackupXL(file)
	backRepo.BackRepoElement.BackupXL(file)
	backRepo.BackRepoEnumeration.BackupXL(file)
	backRepo.BackRepoExtension.BackupXL(file)
	backRepo.BackRepoGroup.BackupXL(file)
	backRepo.BackRepoLength.BackupXL(file)
	backRepo.BackRepoMaxInclusive.BackupXL(file)
	backRepo.BackRepoMaxLength.BackupXL(file)
	backRepo.BackRepoMinInclusive.BackupXL(file)
	backRepo.BackRepoMinLength.BackupXL(file)
	backRepo.BackRepoPattern.BackupXL(file)
	backRepo.BackRepoRestriction.BackupXL(file)
	backRepo.BackRepoSchema.BackupXL(file)
	backRepo.BackRepoSequence.BackupXL(file)
	backRepo.BackRepoSimpleContent.BackupXL(file)
	backRepo.BackRepoSimpleType.BackupXL(file)
	backRepo.BackRepoTotalDigit.BackupXL(file)
	backRepo.BackRepoUnion.BackupXL(file)
	backRepo.BackRepoWhiteSpace.BackupXL(file)

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
func (backRepo *BackRepoStruct) Restore(stage *models.Stage, dirPath string) {
	backRepo.stage.Commit()
	backRepo.stage.Reset()
	backRepo.stage.Checkout()

	//
	// restauration first phase (create DB instance with new IDs)
	//

	// insertion point for per struct backup
	backRepo.BackRepoAll.RestorePhaseOne(dirPath)
	backRepo.BackRepoAnnotation.RestorePhaseOne(dirPath)
	backRepo.BackRepoAttribute.RestorePhaseOne(dirPath)
	backRepo.BackRepoAttributeGroup.RestorePhaseOne(dirPath)
	backRepo.BackRepoChoice.RestorePhaseOne(dirPath)
	backRepo.BackRepoComplexContent.RestorePhaseOne(dirPath)
	backRepo.BackRepoComplexType.RestorePhaseOne(dirPath)
	backRepo.BackRepoDocumentation.RestorePhaseOne(dirPath)
	backRepo.BackRepoElement.RestorePhaseOne(dirPath)
	backRepo.BackRepoEnumeration.RestorePhaseOne(dirPath)
	backRepo.BackRepoExtension.RestorePhaseOne(dirPath)
	backRepo.BackRepoGroup.RestorePhaseOne(dirPath)
	backRepo.BackRepoLength.RestorePhaseOne(dirPath)
	backRepo.BackRepoMaxInclusive.RestorePhaseOne(dirPath)
	backRepo.BackRepoMaxLength.RestorePhaseOne(dirPath)
	backRepo.BackRepoMinInclusive.RestorePhaseOne(dirPath)
	backRepo.BackRepoMinLength.RestorePhaseOne(dirPath)
	backRepo.BackRepoPattern.RestorePhaseOne(dirPath)
	backRepo.BackRepoRestriction.RestorePhaseOne(dirPath)
	backRepo.BackRepoSchema.RestorePhaseOne(dirPath)
	backRepo.BackRepoSequence.RestorePhaseOne(dirPath)
	backRepo.BackRepoSimpleContent.RestorePhaseOne(dirPath)
	backRepo.BackRepoSimpleType.RestorePhaseOne(dirPath)
	backRepo.BackRepoTotalDigit.RestorePhaseOne(dirPath)
	backRepo.BackRepoUnion.RestorePhaseOne(dirPath)
	backRepo.BackRepoWhiteSpace.RestorePhaseOne(dirPath)

	//
	// restauration second phase (reindex pointers with the new ID)
	//

	// insertion point for per struct backup
	backRepo.BackRepoAll.RestorePhaseTwo()
	backRepo.BackRepoAnnotation.RestorePhaseTwo()
	backRepo.BackRepoAttribute.RestorePhaseTwo()
	backRepo.BackRepoAttributeGroup.RestorePhaseTwo()
	backRepo.BackRepoChoice.RestorePhaseTwo()
	backRepo.BackRepoComplexContent.RestorePhaseTwo()
	backRepo.BackRepoComplexType.RestorePhaseTwo()
	backRepo.BackRepoDocumentation.RestorePhaseTwo()
	backRepo.BackRepoElement.RestorePhaseTwo()
	backRepo.BackRepoEnumeration.RestorePhaseTwo()
	backRepo.BackRepoExtension.RestorePhaseTwo()
	backRepo.BackRepoGroup.RestorePhaseTwo()
	backRepo.BackRepoLength.RestorePhaseTwo()
	backRepo.BackRepoMaxInclusive.RestorePhaseTwo()
	backRepo.BackRepoMaxLength.RestorePhaseTwo()
	backRepo.BackRepoMinInclusive.RestorePhaseTwo()
	backRepo.BackRepoMinLength.RestorePhaseTwo()
	backRepo.BackRepoPattern.RestorePhaseTwo()
	backRepo.BackRepoRestriction.RestorePhaseTwo()
	backRepo.BackRepoSchema.RestorePhaseTwo()
	backRepo.BackRepoSequence.RestorePhaseTwo()
	backRepo.BackRepoSimpleContent.RestorePhaseTwo()
	backRepo.BackRepoSimpleType.RestorePhaseTwo()
	backRepo.BackRepoTotalDigit.RestorePhaseTwo()
	backRepo.BackRepoUnion.RestorePhaseTwo()
	backRepo.BackRepoWhiteSpace.RestorePhaseTwo()

	backRepo.stage.Checkout()
}

// Restore the database into the back repo
func (backRepo *BackRepoStruct) RestoreXL(stage *models.Stage, dirPath string) {

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
	backRepo.BackRepoAll.RestoreXLPhaseOne(file)
	backRepo.BackRepoAnnotation.RestoreXLPhaseOne(file)
	backRepo.BackRepoAttribute.RestoreXLPhaseOne(file)
	backRepo.BackRepoAttributeGroup.RestoreXLPhaseOne(file)
	backRepo.BackRepoChoice.RestoreXLPhaseOne(file)
	backRepo.BackRepoComplexContent.RestoreXLPhaseOne(file)
	backRepo.BackRepoComplexType.RestoreXLPhaseOne(file)
	backRepo.BackRepoDocumentation.RestoreXLPhaseOne(file)
	backRepo.BackRepoElement.RestoreXLPhaseOne(file)
	backRepo.BackRepoEnumeration.RestoreXLPhaseOne(file)
	backRepo.BackRepoExtension.RestoreXLPhaseOne(file)
	backRepo.BackRepoGroup.RestoreXLPhaseOne(file)
	backRepo.BackRepoLength.RestoreXLPhaseOne(file)
	backRepo.BackRepoMaxInclusive.RestoreXLPhaseOne(file)
	backRepo.BackRepoMaxLength.RestoreXLPhaseOne(file)
	backRepo.BackRepoMinInclusive.RestoreXLPhaseOne(file)
	backRepo.BackRepoMinLength.RestoreXLPhaseOne(file)
	backRepo.BackRepoPattern.RestoreXLPhaseOne(file)
	backRepo.BackRepoRestriction.RestoreXLPhaseOne(file)
	backRepo.BackRepoSchema.RestoreXLPhaseOne(file)
	backRepo.BackRepoSequence.RestoreXLPhaseOne(file)
	backRepo.BackRepoSimpleContent.RestoreXLPhaseOne(file)
	backRepo.BackRepoSimpleType.RestoreXLPhaseOne(file)
	backRepo.BackRepoTotalDigit.RestoreXLPhaseOne(file)
	backRepo.BackRepoUnion.RestoreXLPhaseOne(file)
	backRepo.BackRepoWhiteSpace.RestoreXLPhaseOne(file)

	// commit the restored stage
	backRepo.stage.Commit()
}

func (backRepoStruct *BackRepoStruct) SubscribeToCommitNb(ctx context.Context) <-chan int {
	ch := make(chan int)

	backRepoStruct.subscribersRwMutex.Lock()
	backRepoStruct.subscribers = append(backRepoStruct.subscribers, ch)
	backRepoStruct.subscribersRwMutex.Unlock()

	// Goroutine to remove subscriber when context is done
	go func() {
		<-ctx.Done()
		backRepoStruct.unsubscribe(ch)
	}()
	return ch
}

// unsubscribe removes a subscriber's channel from the subscribers slice.
func (backRepoStruct *BackRepoStruct) unsubscribe(ch chan int) {
	backRepoStruct.subscribersRwMutex.Lock()
	defer backRepoStruct.subscribersRwMutex.Unlock()
	for i, subscriber := range backRepoStruct.subscribers {
		if subscriber == ch {
			backRepoStruct.subscribers =
				append(backRepoStruct.subscribers[:i],
					backRepoStruct.subscribers[i+1:]...)
			close(ch) // Close the channel to signal completion
			break
		}
	}
}

func (backRepoStruct *BackRepoStruct) broadcastNbCommitToBack() {
	backRepoStruct.subscribersRwMutex.RLock()
	subscribers := make([]chan int, len(backRepoStruct.subscribers))
	copy(subscribers, backRepoStruct.subscribers)
	backRepoStruct.subscribersRwMutex.RUnlock()

	for _, ch := range subscribers {
		select {
		case ch <- int(backRepoStruct.CommitFromBackNb):
			// Successfully sent commit from back
		default:
			// Subscriber is not ready to receive; skip to avoid blocking
		}
	}
}
