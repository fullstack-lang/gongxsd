// generated code - do not edit
package models

import (
	"cmp"
	"embed"
	"errors"
	"fmt"
	"log"
	"math"
	"slices"
	"sort"
	"time"

	books_go "github.com/fullstack-lang/gongxsd/test/books/go"
)

// can be used for
//
//	days := __Gong__Abs(int(int(inferedInstance.ComputedDuration.Hours()) / 24))
func __Gong__Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

var _ = __Gong__Abs

const ProbeTreeSidebarSuffix = ":sidebar of the probe"
const ProbeTableSuffix = ":table of the probe"
const ProbeFormSuffix = ":form of the probe"
const ProbeSplitSuffix = ":probe of the probe"

func (stage *Stage) GetProbeTreeSidebarStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeTreeSidebarSuffix
}

func (stage *Stage) GetProbeFormStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeFormSuffix
}

func (stage *Stage) GetProbeTableStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeTableSuffix
}

func (stage *Stage) GetProbeSplitStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeSplitSuffix
}

// errUnkownEnum is returns when a value cannot match enum values
var errUnkownEnum = errors.New("unkown enum")

// needed to avoid when fmt package is not needed by generated code
var __dummy__fmt_variable fmt.Scanner

var _ = __dummy__fmt_variable

// idem for math package when not need by generated code
var __dummy_math_variable = math.E

var _ = __dummy_math_variable

// swagger:ignore
type __void any

// needed for creating set of instances in the stage
var __member __void

// GongStructInterface is the interface met by GongStructs
// It allows runtime reflexion of instances (without the hassle of the "reflect" package)
type GongStructInterface interface {
	GetName() (res string)
	// GetID() (res int)
	// GetFields() (res []string)
	// GetFieldStringValue(fieldName string) (res string)
}

// Stage enables storage of staged instances
// swagger:ignore
type Stage struct {
	name               string
	commitId           uint // commitId is updated at each commit
	commitTimeStamp    time.Time
	contentWhenParsed  string
	commitIdWhenParsed uint
	generatesDiff      bool

	// insertion point for definition of arrays registering instances
	BookTypes           map[*BookType]any
	BookTypes_mapString map[string]*BookType

	// insertion point for slice of pointers maps
	BookType_Credit_reverseMap map[*Credit]*BookType

	OnAfterBookTypeCreateCallback OnAfterCreateInterface[BookType]
	OnAfterBookTypeUpdateCallback OnAfterUpdateInterface[BookType]
	OnAfterBookTypeDeleteCallback OnAfterDeleteInterface[BookType]
	OnAfterBookTypeReadCallback   OnAfterReadInterface[BookType]

	Bookss           map[*Books]any
	Bookss_mapString map[string]*Books

	// insertion point for slice of pointers maps
	Books_Book_reverseMap map[*BookType]*Books

	OnAfterBooksCreateCallback OnAfterCreateInterface[Books]
	OnAfterBooksUpdateCallback OnAfterUpdateInterface[Books]
	OnAfterBooksDeleteCallback OnAfterDeleteInterface[Books]
	OnAfterBooksReadCallback   OnAfterReadInterface[Books]

	Credits           map[*Credit]any
	Credits_mapString map[string]*Credit

	// insertion point for slice of pointers maps
	Credit_Link_reverseMap map[*Link]*Credit

	OnAfterCreditCreateCallback OnAfterCreateInterface[Credit]
	OnAfterCreditUpdateCallback OnAfterUpdateInterface[Credit]
	OnAfterCreditDeleteCallback OnAfterDeleteInterface[Credit]
	OnAfterCreditReadCallback   OnAfterReadInterface[Credit]

	Links           map[*Link]any
	Links_mapString map[string]*Link

	// insertion point for slice of pointers maps
	OnAfterLinkCreateCallback OnAfterCreateInterface[Link]
	OnAfterLinkUpdateCallback OnAfterUpdateInterface[Link]
	OnAfterLinkDeleteCallback OnAfterDeleteInterface[Link]
	OnAfterLinkReadCallback   OnAfterReadInterface[Link]

	AllModelsStructCreateCallback AllModelsStructCreateInterface

	AllModelsStructDeleteCallback AllModelsStructDeleteInterface

	BackRepo BackRepoInterface

	// if set will be called before each commit to the back repo
	OnInitCommitCallback          OnInitCommitInterface
	OnInitCommitFromFrontCallback OnInitCommitInterface
	OnInitCommitFromBackCallback  OnInitCommitInterface

	// store the number of instance per gongstruct
	Map_GongStructName_InstancesNb map[string]int

	// store meta package import
	MetaPackageImportPath  string
	MetaPackageImportAlias string

	// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
	// map to enable docLink renaming when an identifier is renamed
	Map_DocLink_Renaming map[string]GONG__Identifier
	// the to be removed stops here

	// store the stage order of each instance in order to
	// preserve this order when serializing them
	// insertion point for order fields declaration
	BookTypeOrder            uint
	BookTypeMap_Staged_Order map[*BookType]uint

	BooksOrder            uint
	BooksMap_Staged_Order map[*Books]uint

	CreditOrder            uint
	CreditMap_Staged_Order map[*Credit]uint

	LinkOrder            uint
	LinkMap_Staged_Order map[*Link]uint

	// end of insertion point

	NamedStructs []*NamedStruct
}

func (stage *Stage) GetCommitId() uint {
	return stage.commitId
}

func (stage *Stage) GetCommitTS() time.Time {
	return stage.commitTimeStamp
}

func (stage *Stage) SetGeneratesDiff(generatesDiff bool) {
	stage.generatesDiff = generatesDiff
}

// GetNamedStructs implements models.ProbebStage.
func (stage *Stage) GetNamedStructsNames() (res []string) {

	for _, namedStruct := range stage.NamedStructs {
		res = append(res, namedStruct.name)
	}

	return
}

func GetNamedStructInstances[T PointerToGongstruct](set map[T]any, order map[T]uint) (res []string) {

	orderedSet := []T{}
	for instance := range set {
		orderedSet = append(orderedSet, instance)
	}
	sort.Slice(orderedSet[:], func(i, j int) bool {
		instancei := orderedSet[i]
		instancej := orderedSet[j]
		i_order, oki := order[instancei]
		j_order, okj := order[instancej]
		if !oki || !okj {
			log.Fatalf("GetNamedStructInstances: pointer not found")
		}
		return i_order < j_order
	})

	for _, instance := range orderedSet {
		res = append(res, instance.GetName())
	}

	return
}

func GetStructInstancesByOrderAuto[T PointerToGongstruct](stage *Stage) (res []T) {
	var t T
	switch any(t).(type) {
		// insertion point for case
	case *BookType:
		tmp := GetStructInstancesByOrder(stage.BookTypes, stage.BookTypeMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *BookType implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Books:
		tmp := GetStructInstancesByOrder(stage.Bookss, stage.BooksMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Books implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Credit:
		tmp := GetStructInstancesByOrder(stage.Credits, stage.CreditMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Credit implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Link:
		tmp := GetStructInstancesByOrder(stage.Links, stage.LinkMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Link implements.
			res = append(res, any(v).(T))
		}
		return res

	}
	return
}

func GetStructInstancesByOrder[T PointerToGongstruct](set map[T]any, order map[T]uint) (res []T) {

	orderedSet := []T{}
	for instance := range set {
		orderedSet = append(orderedSet, instance)
	}
	sort.Slice(orderedSet[:], func(i, j int) bool {
		instancei := orderedSet[i]
		instancej := orderedSet[j]
		i_order, oki := order[instancei]
		j_order, okj := order[instancej]
		if !oki || !okj {
			log.Fatalf("GetNamedStructInstances: pointer not found")
		}
		return i_order < j_order
	})

	for _, instance := range orderedSet {
		res = append(res, instance)
	}

	return
}

func (stage *Stage) GetNamedStructNamesByOrder(namedStructName string) (res []string) {

	switch namedStructName {
	// insertion point for case
	case "BookType":
		res = GetNamedStructInstances(stage.BookTypes, stage.BookTypeMap_Staged_Order)
	case "Books":
		res = GetNamedStructInstances(stage.Bookss, stage.BooksMap_Staged_Order)
	case "Credit":
		res = GetNamedStructInstances(stage.Credits, stage.CreditMap_Staged_Order)
	case "Link":
		res = GetNamedStructInstances(stage.Links, stage.LinkMap_Staged_Order)
	}

	return
}

type NamedStruct struct {
	name string
}

func (namedStruct *NamedStruct) GetName() string {
	return namedStruct.name
}

func (stage *Stage) GetType() string {
	return "github.com/fullstack-lang/gongxsd/test/books/go/models"
}

func (stage *Stage) GetMap_GongStructName_InstancesNb() map[string]int {
	return stage.Map_GongStructName_InstancesNb
}

func (stage *Stage) GetModelsEmbededDir() embed.FS {
	return books_go.GoModelsDir
}

func (stage *Stage) GetDigramsEmbededDir() embed.FS {
	return books_go.GoDiagramsDir
}

type GONG__Identifier struct {
	Ident string
	Type  GONG__ExpressionType
}

type OnInitCommitInterface interface {
	BeforeCommit(stage *Stage)
}

// OnAfterCreateInterface callback when an instance is updated from the front
type OnAfterCreateInterface[Type Gongstruct] interface {
	OnAfterCreate(stage *Stage,
		instance *Type)
}

// OnAfterReadInterface callback when an instance is updated from the front
type OnAfterReadInterface[Type Gongstruct] interface {
	OnAfterRead(stage *Stage,
		instance *Type)
}

// OnAfterUpdateInterface callback when an instance is updated from the front
type OnAfterUpdateInterface[Type Gongstruct] interface {
	OnAfterUpdate(stage *Stage, old, new *Type)
}

// OnAfterDeleteInterface callback when an instance is updated from the front
type OnAfterDeleteInterface[Type Gongstruct] interface {
	OnAfterDelete(stage *Stage,
		staged, front *Type)
}

type BackRepoInterface interface {
	Commit(stage *Stage)
	Checkout(stage *Stage)
	Backup(stage *Stage, dirPath string)
	Restore(stage *Stage, dirPath string)
	BackupXL(stage *Stage, dirPath string)
	RestoreXL(stage *Stage, dirPath string)
	// insertion point for Commit and Checkout signatures
	CommitBookType(booktype *BookType)
	CheckoutBookType(booktype *BookType)
	CommitBooks(books *Books)
	CheckoutBooks(books *Books)
	CommitCredit(credit *Credit)
	CheckoutCredit(credit *Credit)
	CommitLink(link *Link)
	CheckoutLink(link *Link)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

func NewStage(name string) (stage *Stage) {

	stage = &Stage{ // insertion point for array initiatialisation
		BookTypes:           make(map[*BookType]any),
		BookTypes_mapString: make(map[string]*BookType),

		Bookss:           make(map[*Books]any),
		Bookss_mapString: make(map[string]*Books),

		Credits:           make(map[*Credit]any),
		Credits_mapString: make(map[string]*Credit),

		Links:           make(map[*Link]any),
		Links_mapString: make(map[string]*Link),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		name: name,

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		// insertion point for order map initialisations
		BookTypeMap_Staged_Order: make(map[*BookType]uint),

		BooksMap_Staged_Order: make(map[*Books]uint),

		CreditMap_Staged_Order: make(map[*Credit]uint),

		LinkMap_Staged_Order: make(map[*Link]uint),

		// end of insertion point

		NamedStructs: []*NamedStruct{ // insertion point for order map initialisations
			{name: "BookType"},
			{name: "Books"},
			{name: "Credit"},
			{name: "Link"},
		}, // end of insertion point
	}

	return
}

func GetOrder[Type Gongstruct](stage *Stage, instance *Type) uint {

	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *BookType:
		return stage.BookTypeMap_Staged_Order[instance]
	case *Books:
		return stage.BooksMap_Staged_Order[instance]
	case *Credit:
		return stage.CreditMap_Staged_Order[instance]
	case *Link:
		return stage.LinkMap_Staged_Order[instance]
	default:
		return 0 // should not happen
	}
}

func GetOrderPointerGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) uint {

	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *BookType:
		return stage.BookTypeMap_Staged_Order[instance]
	case *Books:
		return stage.BooksMap_Staged_Order[instance]
	case *Credit:
		return stage.CreditMap_Staged_Order[instance]
	case *Link:
		return stage.LinkMap_Staged_Order[instance]
	default:
		return 0 // should not happen
	}
}

func (stage *Stage) GetName() string {
	return stage.name
}

func (stage *Stage) CommitWithSuspendedCallbacks() {

	tmp := stage.OnInitCommitFromBackCallback
	stage.OnInitCommitFromBackCallback = nil
	stage.Commit()
	stage.OnInitCommitFromBackCallback = tmp
}

func (stage *Stage) Commit() {
	stage.ComputeReverseMaps()
	stage.commitId++
	stage.commitTimeStamp = time.Now()

	if stage.BackRepo != nil {
		stage.BackRepo.Commit(stage)
	}

	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["BookType"] = len(stage.BookTypes)
	stage.Map_GongStructName_InstancesNb["Books"] = len(stage.Bookss)
	stage.Map_GongStructName_InstancesNb["Credit"] = len(stage.Credits)
	stage.Map_GongStructName_InstancesNb["Link"] = len(stage.Links)

}

func (stage *Stage) Checkout() {
	if stage.BackRepo != nil {
		stage.BackRepo.Checkout(stage)
	}

	stage.ComputeReverseMaps()
	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["BookType"] = len(stage.BookTypes)
	stage.Map_GongStructName_InstancesNb["Books"] = len(stage.Bookss)
	stage.Map_GongStructName_InstancesNb["Credit"] = len(stage.Credits)
	stage.Map_GongStructName_InstancesNb["Link"] = len(stage.Links)

}

// backup generates backup files in the dirPath
func (stage *Stage) Backup(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.Backup(stage, dirPath)
	}
}

// Restore resets Stage & BackRepo and restores their content from the restore files in dirPath
func (stage *Stage) Restore(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.Restore(stage, dirPath)
	}
}

// backup generates backup files in the dirPath
func (stage *Stage) BackupXL(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.BackupXL(stage, dirPath)
	}
}

// Restore resets Stage & BackRepo and restores their content from the restore files in dirPath
func (stage *Stage) RestoreXL(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.RestoreXL(stage, dirPath)
	}
}

// insertion point for cumulative sub template with model space calls
// Stage puts booktype to the model stage
func (booktype *BookType) Stage(stage *Stage) *BookType {

	if _, ok := stage.BookTypes[booktype]; !ok {
		stage.BookTypes[booktype] = __member
		stage.BookTypeMap_Staged_Order[booktype] = stage.BookTypeOrder
		stage.BookTypeOrder++
	}
	stage.BookTypes_mapString[booktype.Name] = booktype

	return booktype
}

// Unstage removes booktype off the model stage
func (booktype *BookType) Unstage(stage *Stage) *BookType {
	delete(stage.BookTypes, booktype)
	delete(stage.BookTypes_mapString, booktype.Name)
	return booktype
}

// UnstageVoid removes booktype off the model stage
func (booktype *BookType) UnstageVoid(stage *Stage) {
	delete(stage.BookTypes, booktype)
	delete(stage.BookTypes_mapString, booktype.Name)
}

// commit booktype to the back repo (if it is already staged)
func (booktype *BookType) Commit(stage *Stage) *BookType {
	if _, ok := stage.BookTypes[booktype]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitBookType(booktype)
		}
	}
	return booktype
}

func (booktype *BookType) CommitVoid(stage *Stage) {
	booktype.Commit(stage)
}

// Checkout booktype to the back repo (if it is already staged)
func (booktype *BookType) Checkout(stage *Stage) *BookType {
	if _, ok := stage.BookTypes[booktype]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutBookType(booktype)
		}
	}
	return booktype
}

// for satisfaction of GongStruct interface
func (booktype *BookType) GetName() (res string) {
	return booktype.Name
}

// Stage puts books to the model stage
func (books *Books) Stage(stage *Stage) *Books {

	if _, ok := stage.Bookss[books]; !ok {
		stage.Bookss[books] = __member
		stage.BooksMap_Staged_Order[books] = stage.BooksOrder
		stage.BooksOrder++
	}
	stage.Bookss_mapString[books.Name] = books

	return books
}

// Unstage removes books off the model stage
func (books *Books) Unstage(stage *Stage) *Books {
	delete(stage.Bookss, books)
	delete(stage.Bookss_mapString, books.Name)
	return books
}

// UnstageVoid removes books off the model stage
func (books *Books) UnstageVoid(stage *Stage) {
	delete(stage.Bookss, books)
	delete(stage.Bookss_mapString, books.Name)
}

// commit books to the back repo (if it is already staged)
func (books *Books) Commit(stage *Stage) *Books {
	if _, ok := stage.Bookss[books]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitBooks(books)
		}
	}
	return books
}

func (books *Books) CommitVoid(stage *Stage) {
	books.Commit(stage)
}

// Checkout books to the back repo (if it is already staged)
func (books *Books) Checkout(stage *Stage) *Books {
	if _, ok := stage.Bookss[books]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutBooks(books)
		}
	}
	return books
}

// for satisfaction of GongStruct interface
func (books *Books) GetName() (res string) {
	return books.Name
}

// Stage puts credit to the model stage
func (credit *Credit) Stage(stage *Stage) *Credit {

	if _, ok := stage.Credits[credit]; !ok {
		stage.Credits[credit] = __member
		stage.CreditMap_Staged_Order[credit] = stage.CreditOrder
		stage.CreditOrder++
	}
	stage.Credits_mapString[credit.Name] = credit

	return credit
}

// Unstage removes credit off the model stage
func (credit *Credit) Unstage(stage *Stage) *Credit {
	delete(stage.Credits, credit)
	delete(stage.Credits_mapString, credit.Name)
	return credit
}

// UnstageVoid removes credit off the model stage
func (credit *Credit) UnstageVoid(stage *Stage) {
	delete(stage.Credits, credit)
	delete(stage.Credits_mapString, credit.Name)
}

// commit credit to the back repo (if it is already staged)
func (credit *Credit) Commit(stage *Stage) *Credit {
	if _, ok := stage.Credits[credit]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitCredit(credit)
		}
	}
	return credit
}

func (credit *Credit) CommitVoid(stage *Stage) {
	credit.Commit(stage)
}

// Checkout credit to the back repo (if it is already staged)
func (credit *Credit) Checkout(stage *Stage) *Credit {
	if _, ok := stage.Credits[credit]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutCredit(credit)
		}
	}
	return credit
}

// for satisfaction of GongStruct interface
func (credit *Credit) GetName() (res string) {
	return credit.Name
}

// Stage puts link to the model stage
func (link *Link) Stage(stage *Stage) *Link {

	if _, ok := stage.Links[link]; !ok {
		stage.Links[link] = __member
		stage.LinkMap_Staged_Order[link] = stage.LinkOrder
		stage.LinkOrder++
	}
	stage.Links_mapString[link.Name] = link

	return link
}

// Unstage removes link off the model stage
func (link *Link) Unstage(stage *Stage) *Link {
	delete(stage.Links, link)
	delete(stage.Links_mapString, link.Name)
	return link
}

// UnstageVoid removes link off the model stage
func (link *Link) UnstageVoid(stage *Stage) {
	delete(stage.Links, link)
	delete(stage.Links_mapString, link.Name)
}

// commit link to the back repo (if it is already staged)
func (link *Link) Commit(stage *Stage) *Link {
	if _, ok := stage.Links[link]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitLink(link)
		}
	}
	return link
}

func (link *Link) CommitVoid(stage *Stage) {
	link.Commit(stage)
}

// Checkout link to the back repo (if it is already staged)
func (link *Link) Checkout(stage *Stage) *Link {
	if _, ok := stage.Links[link]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutLink(link)
		}
	}
	return link
}

// for satisfaction of GongStruct interface
func (link *Link) GetName() (res string) {
	return link.Name
}

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMBookType(BookType *BookType)
	CreateORMBooks(Books *Books)
	CreateORMCredit(Credit *Credit)
	CreateORMLink(Link *Link)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMBookType(BookType *BookType)
	DeleteORMBooks(Books *Books)
	DeleteORMCredit(Credit *Credit)
	DeleteORMLink(Link *Link)
}

func (stage *Stage) Reset() { // insertion point for array reset
	stage.BookTypes = make(map[*BookType]any)
	stage.BookTypes_mapString = make(map[string]*BookType)
	stage.BookTypeMap_Staged_Order = make(map[*BookType]uint)
	stage.BookTypeOrder = 0

	stage.Bookss = make(map[*Books]any)
	stage.Bookss_mapString = make(map[string]*Books)
	stage.BooksMap_Staged_Order = make(map[*Books]uint)
	stage.BooksOrder = 0

	stage.Credits = make(map[*Credit]any)
	stage.Credits_mapString = make(map[string]*Credit)
	stage.CreditMap_Staged_Order = make(map[*Credit]uint)
	stage.CreditOrder = 0

	stage.Links = make(map[*Link]any)
	stage.Links_mapString = make(map[string]*Link)
	stage.LinkMap_Staged_Order = make(map[*Link]uint)
	stage.LinkOrder = 0

}

func (stage *Stage) Nil() { // insertion point for array nil
	stage.BookTypes = nil
	stage.BookTypes_mapString = nil

	stage.Bookss = nil
	stage.Bookss_mapString = nil

	stage.Credits = nil
	stage.Credits_mapString = nil

	stage.Links = nil
	stage.Links_mapString = nil

}

func (stage *Stage) Unstage() { // insertion point for array nil
	for booktype := range stage.BookTypes {
		booktype.Unstage(stage)
	}

	for books := range stage.Bookss {
		books.Unstage(stage)
	}

	for credit := range stage.Credits {
		credit.Unstage(stage)
	}

	for link := range stage.Links {
		link.Unstage(stage)
	}

}

// Gongstruct is the type parameter for generated generic function that allows
// - access to staged instances
// - navigation between staged instances by going backward association links between gongstruct
// - full refactoring of Gongstruct identifiers / fields
type Gongstruct interface {
}

type GongtructBasicField interface {
	int | float64 | bool | string | time.Time | time.Duration
}

// Gongstruct is the type parameter for generated generic function that allows
// - access to staged instances
// - navigation between staged instances by going backward association links between gongstruct
// - full refactoring of Gongstruct identifiers / fields
type PointerToGongstruct interface {
	GetName() string
	CommitVoid(*Stage)
	UnstageVoid(stage *Stage)
	comparable
}

func CompareGongstructByName[T PointerToGongstruct](a, b T) int {
	return cmp.Compare(a.GetName(), b.GetName())
}

func SortGongstructSetByName[T PointerToGongstruct](set map[T]any) (sortedSlice []T) {

	for key := range set {
		sortedSlice = append(sortedSlice, key)
	}
	slices.SortFunc(sortedSlice, CompareGongstructByName)

	return
}

func GetGongstrucsSorted[T PointerToGongstruct](stage *Stage) (sortedSlice []T) {

	set := GetGongstructInstancesSetFromPointerType[T](stage)
	sortedSlice = SortGongstructSetByName(*set)

	return
}

type GongstructSet interface {
	map[any]any
}

type GongstructMapString interface {
	map[any]any
}

// GongGetSet returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GongGetSet[Type GongstructSet](stage *Stage) *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case map[*BookType]any:
		return any(&stage.BookTypes).(*Type)
	case map[*Books]any:
		return any(&stage.Bookss).(*Type)
	case map[*Credit]any:
		return any(&stage.Credits).(*Type)
	case map[*Link]any:
		return any(&stage.Links).(*Type)
	default:
		return nil
	}
}

// GongGetMap returns the map of staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GongGetMap[Type GongstructMapString](stage *Stage) *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case map[string]*BookType:
		return any(&stage.BookTypes_mapString).(*Type)
	case map[string]*Books:
		return any(&stage.Bookss_mapString).(*Type)
	case map[string]*Credit:
		return any(&stage.Credits_mapString).(*Type)
	case map[string]*Link:
		return any(&stage.Links_mapString).(*Type)
	default:
		return nil
	}
}

// GetGongstructInstancesSet returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesSet[Type Gongstruct](stage *Stage) *map[*Type]any {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case BookType:
		return any(&stage.BookTypes).(*map[*Type]any)
	case Books:
		return any(&stage.Bookss).(*map[*Type]any)
	case Credit:
		return any(&stage.Credits).(*map[*Type]any)
	case Link:
		return any(&stage.Links).(*map[*Type]any)
	default:
		return nil
	}
}

// GetGongstructInstancesSetFromPointerType returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesSetFromPointerType[Type PointerToGongstruct](stage *Stage) *map[Type]any {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case *BookType:
		return any(&stage.BookTypes).(*map[Type]any)
	case *Books:
		return any(&stage.Bookss).(*map[Type]any)
	case *Credit:
		return any(&stage.Credits).(*map[Type]any)
	case *Link:
		return any(&stage.Links).(*map[Type]any)
	default:
		return nil
	}
}

// GetGongstructInstancesMap returns the map of staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GetGongstructInstancesMap[Type Gongstruct](stage *Stage) *map[string]*Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case BookType:
		return any(&stage.BookTypes_mapString).(*map[string]*Type)
	case Books:
		return any(&stage.Bookss_mapString).(*map[string]*Type)
	case Credit:
		return any(&stage.Credits_mapString).(*map[string]*Type)
	case Link:
		return any(&stage.Links_mapString).(*map[string]*Type)
	default:
		return nil
	}
}

// GetAssociationName is a generic function that returns an instance of Type
// where each association is filled with an instance whose name is the name of the association
//
// This function can be handy for generating navigation function that are refactorable
func GetAssociationName[Type Gongstruct]() *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for instance with special fields
	case BookType:
		return any(&BookType{
			// Initialisation of associations
			// field is initialized with an instance of Credit with the name of the field
			Credit: []*Credit{{Name: "Credit"}},
		}).(*Type)
	case Books:
		return any(&Books{
			// Initialisation of associations
		}).(*Type)
	case Credit:
		return any(&Credit{
			// Initialisation of associations
			// field is initialized with an instance of Link with the name of the field
			Link: []*Link{{Name: "Link"}},
		}).(*Type)
	case Link:
		return any(&Link{
			// Initialisation of associations
		}).(*Type)
	default:
		return nil
	}
}

// GetPointerReverseMap allows backtrack navigation of any Start.Fieldname
// associations (0..1) that is a pointer from one staged Gongstruct (type Start)
// instances to another (type End)
//
// The function provides a map with keys as instances of End and values to arrays of *Start
// the map is construed by iterating over all Start instances and populationg keys with End instances
// and values with slice of Start instances
func GetPointerReverseMap[Start, End Gongstruct](fieldname string, stage *Stage) map[*End][]*Start {

	var ret Start

	switch any(ret).(type) {
	// insertion point of functions that provide maps for reverse associations
	// reverse maps of direct associations of BookType
	case BookType:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Books
	case Books:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Credit
	case Credit:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Link
	case Link:
		switch fieldname {
		// insertion point for per direct association field
		}
	}
	return nil
}

// GetSliceOfPointersReverseMap allows backtrack navigation of any Start.Fieldname
// associations (0..N) between one staged Gongstruct instances and many others
//
// The function provides a map with keys as instances of End and values to *Start instances
// the map is construed by iterating over all Start instances and populating keys with End instances
// and values with the Start instances
func GetSliceOfPointersReverseMap[Start, End Gongstruct](fieldname string, stage *Stage) map[*End][]*Start {

	var ret Start

	switch any(ret).(type) {
	// insertion point of functions that provide maps for reverse associations
	// reverse maps of direct associations of BookType
	case BookType:
		switch fieldname {
		// insertion point for per direct association field
		case "Credit":
			res := make(map[*Credit][]*BookType)
			for booktype := range stage.BookTypes {
				for _, credit_ := range booktype.Credit {
					res[credit_] = append(res[credit_], booktype)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Books
	case Books:
		switch fieldname {
		// insertion point for per direct association field
		case "Book":
			res := make(map[*BookType][]*Books)
			for books := range stage.Bookss {
				for _, booktype_ := range books.Book {
					res[booktype_] = append(res[booktype_], books)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Credit
	case Credit:
		switch fieldname {
		// insertion point for per direct association field
		case "Link":
			res := make(map[*Link][]*Credit)
			for credit := range stage.Credits {
				for _, link_ := range credit.Link {
					res[link_] = append(res[link_], credit)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Link
	case Link:
		switch fieldname {
		// insertion point for per direct association field
		}
	}
	return nil
}

// GetGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetGongstructName[Type Gongstruct]() (res string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case BookType:
		res = "BookType"
	case Books:
		res = "Books"
	case Credit:
		res = "Credit"
	case Link:
		res = "Link"
	}
	return res
}

// GetPointerToGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetPointerToGongstructName[Type PointerToGongstruct]() (res string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case *BookType:
		res = "BookType"
	case *Books:
		res = "Books"
	case *Credit:
		res = "Credit"
	case *Link:
		res = "Link"
	}
	return res
}

// GetFields return the array of the fields
func GetFields[Type Gongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case BookType:
		res = []string{"Name", "Edition", "Isbn", "Bestseller", "Title", "Author", "Year", "Format", "Credit"}
	case Books:
		res = []string{"Name", "Book"}
	case Credit:
		res = []string{"Name", "Page", "Credit_type", "Link", "Credit_words", "Credit_symbol"}
	case Link:
		res = []string{"Name", "NameXSD", "EnclosedText"}
	}
	return
}

type ReverseField struct {
	GongstructName string
	Fieldname      string
}

func GetReverseFields[Type Gongstruct]() (res []ReverseField) {

	res = make([]ReverseField, 0)

	var ret Type

	switch any(ret).(type) {

	// insertion point for generic get gongstruct name
	case BookType:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Books"
		rf.Fieldname = "Book"
		res = append(res, rf)
	case Books:
		var rf ReverseField
		_ = rf
	case Credit:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "BookType"
		rf.Fieldname = "Credit"
		res = append(res, rf)
	case Link:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Credit"
		rf.Fieldname = "Link"
		res = append(res, rf)
	}
	return
}

// GetFieldsFromPointer return the array of the fields
func GetFieldsFromPointer[Type PointerToGongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case *BookType:
		res = []string{"Name", "Edition", "Isbn", "Bestseller", "Title", "Author", "Year", "Format", "Credit"}
	case *Books:
		res = []string{"Name", "Book"}
	case *Credit:
		res = []string{"Name", "Page", "Credit_type", "Link", "Credit_words", "Credit_symbol"}
	case *Link:
		res = []string{"Name", "NameXSD", "EnclosedText"}
	}
	return
}

type GongFieldValueType string

const (
	GongFieldValueTypeInt    GongFieldValueType = "GongFieldValueTypeInt"
	GongFieldValueTypeFloat  GongFieldValueType = "GongFieldValueTypeFloat"
	GongFieldValueTypeBool   GongFieldValueType = "GongFieldValueTypeBool"
	GongFieldValueTypeOthers GongFieldValueType = "GongFieldValueTypeOthers"
)

type GongFieldValue struct {
	valueString string
	GongFieldValueType
	valueInt   int
	valueFloat float64
	valueBool  bool
}

func (gongValueField *GongFieldValue) GetValueString() string {
	return gongValueField.valueString
}

func (gongValueField *GongFieldValue) GetValueInt() int {
	return gongValueField.valueInt
}

func (gongValueField *GongFieldValue) GetValueFloat() float64 {
	return gongValueField.valueFloat
}

func (gongValueField *GongFieldValue) GetValueBool() bool {
	return gongValueField.valueBool
}

func GetFieldStringValueFromPointer(instance any, fieldName string) (res GongFieldValue) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
	case *BookType:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Edition":
			res.valueString = inferedInstance.Edition
		case "Isbn":
			res.valueString = inferedInstance.Isbn
		case "Bestseller":
			res.valueString = fmt.Sprintf("%t", inferedInstance.Bestseller)
			res.valueBool = inferedInstance.Bestseller
			res.GongFieldValueType = GongFieldValueTypeBool
		case "Title":
			res.valueString = inferedInstance.Title
		case "Author":
			res.valueString = inferedInstance.Author
		case "Year":
			res.valueString = fmt.Sprintf("%d", inferedInstance.Year)
			res.valueInt = inferedInstance.Year
			res.GongFieldValueType = GongFieldValueTypeInt
		case "Format":
			res.valueString = inferedInstance.Format
		case "Credit":
			for idx, __instance__ := range inferedInstance.Credit {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case *Books:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Book":
			for idx, __instance__ := range inferedInstance.Book {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case *Credit:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Page":
			res.valueString = fmt.Sprintf("%d", inferedInstance.Page)
			res.valueInt = inferedInstance.Page
			res.GongFieldValueType = GongFieldValueTypeInt
		case "Credit_type":
			res.valueString = inferedInstance.Credit_type
		case "Link":
			for idx, __instance__ := range inferedInstance.Link {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "Credit_words":
			res.valueString = inferedInstance.Credit_words
		case "Credit_symbol":
			res.valueString = inferedInstance.Credit_symbol
		}
	case *Link:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "NameXSD":
			res.valueString = inferedInstance.NameXSD
		case "EnclosedText":
			res.valueString = inferedInstance.EnclosedText
		}
	default:
		_ = inferedInstance
	}
	return
}

func GetFieldStringValue(instance any, fieldName string) (res GongFieldValue) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
	case BookType:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Edition":
			res.valueString = inferedInstance.Edition
		case "Isbn":
			res.valueString = inferedInstance.Isbn
		case "Bestseller":
			res.valueString = fmt.Sprintf("%t", inferedInstance.Bestseller)
			res.valueBool = inferedInstance.Bestseller
			res.GongFieldValueType = GongFieldValueTypeBool
		case "Title":
			res.valueString = inferedInstance.Title
		case "Author":
			res.valueString = inferedInstance.Author
		case "Year":
			res.valueString = fmt.Sprintf("%d", inferedInstance.Year)
			res.valueInt = inferedInstance.Year
			res.GongFieldValueType = GongFieldValueTypeInt
		case "Format":
			res.valueString = inferedInstance.Format
		case "Credit":
			for idx, __instance__ := range inferedInstance.Credit {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case Books:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Book":
			for idx, __instance__ := range inferedInstance.Book {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case Credit:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Page":
			res.valueString = fmt.Sprintf("%d", inferedInstance.Page)
			res.valueInt = inferedInstance.Page
			res.GongFieldValueType = GongFieldValueTypeInt
		case "Credit_type":
			res.valueString = inferedInstance.Credit_type
		case "Link":
			for idx, __instance__ := range inferedInstance.Link {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "Credit_words":
			res.valueString = inferedInstance.Credit_words
		case "Credit_symbol":
			res.valueString = inferedInstance.Credit_symbol
		}
	case Link:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "NameXSD":
			res.valueString = inferedInstance.NameXSD
		case "EnclosedText":
			res.valueString = inferedInstance.EnclosedText
		}
	default:
		_ = inferedInstance
	}
	return
}

// Last line of the template
