// generated code - do not edit
package models

import (
	"cmp"
	"errors"
	"fmt"
	"math"
	"slices"
	"time"

	"golang.org/x/exp/maps"
)

func __Gong__Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// errUnkownEnum is returns when a value cannot match enum values
var errUnkownEnum = errors.New("unkown enum")

// needed to avoid when fmt package is not needed by generated code
var __dummy__fmt_variable fmt.Scanner

// idem for math package when not need by generated code
var __dummy_math_variable = math.E

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

// StageStruct enables storage of staged instances
// swagger:ignore
type StageStruct struct {
	path string

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
}

func (stage *StageStruct) GetType() string {
	return "github.com/fullstack-lang/gongxsd/test/books/go/models"
}

type GONG__Identifier struct {
	Ident string
	Type  GONG__ExpressionType
}

type OnInitCommitInterface interface {
	BeforeCommit(stage *StageStruct)
}

// OnAfterCreateInterface callback when an instance is updated from the front
type OnAfterCreateInterface[Type Gongstruct] interface {
	OnAfterCreate(stage *StageStruct,
		instance *Type)
}

// OnAfterReadInterface callback when an instance is updated from the front
type OnAfterReadInterface[Type Gongstruct] interface {
	OnAfterRead(stage *StageStruct,
		instance *Type)
}

// OnAfterUpdateInterface callback when an instance is updated from the front
type OnAfterUpdateInterface[Type Gongstruct] interface {
	OnAfterUpdate(stage *StageStruct, old, new *Type)
}

// OnAfterDeleteInterface callback when an instance is updated from the front
type OnAfterDeleteInterface[Type Gongstruct] interface {
	OnAfterDelete(stage *StageStruct,
		staged, front *Type)
}

type BackRepoInterface interface {
	Commit(stage *StageStruct)
	Checkout(stage *StageStruct)
	Backup(stage *StageStruct, dirPath string)
	Restore(stage *StageStruct, dirPath string)
	BackupXL(stage *StageStruct, dirPath string)
	RestoreXL(stage *StageStruct, dirPath string)
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

func NewStage(path string) (stage *StageStruct) {

	stage = &StageStruct{ // insertion point for array initiatialisation
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

		path: path,

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here
	}

	return
}

func (stage *StageStruct) GetPath() string {
	return stage.path
}

func (stage *StageStruct) CommitWithSuspendedCallbacks() {

	tmp := stage.OnInitCommitFromBackCallback
	stage.OnInitCommitFromBackCallback = nil
	stage.Commit()
	stage.OnInitCommitFromBackCallback = tmp
}

func (stage *StageStruct) Commit() {
	stage.ComputeReverseMaps()

	if stage.BackRepo != nil {
		stage.BackRepo.Commit(stage)
	}

	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["BookType"] = len(stage.BookTypes)
	stage.Map_GongStructName_InstancesNb["Books"] = len(stage.Bookss)
	stage.Map_GongStructName_InstancesNb["Credit"] = len(stage.Credits)
	stage.Map_GongStructName_InstancesNb["Link"] = len(stage.Links)

}

func (stage *StageStruct) Checkout() {
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
func (stage *StageStruct) Backup(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.Backup(stage, dirPath)
	}
}

// Restore resets Stage & BackRepo and restores their content from the restore files in dirPath
func (stage *StageStruct) Restore(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.Restore(stage, dirPath)
	}
}

// backup generates backup files in the dirPath
func (stage *StageStruct) BackupXL(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.BackupXL(stage, dirPath)
	}
}

// Restore resets Stage & BackRepo and restores their content from the restore files in dirPath
func (stage *StageStruct) RestoreXL(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.RestoreXL(stage, dirPath)
	}
}

// insertion point for cumulative sub template with model space calls
// Stage puts booktype to the model stage
func (booktype *BookType) Stage(stage *StageStruct) *BookType {
	stage.BookTypes[booktype] = __member
	stage.BookTypes_mapString[booktype.Name] = booktype

	return booktype
}

// Unstage removes booktype off the model stage
func (booktype *BookType) Unstage(stage *StageStruct) *BookType {
	delete(stage.BookTypes, booktype)
	delete(stage.BookTypes_mapString, booktype.Name)
	return booktype
}

// UnstageVoid removes booktype off the model stage
func (booktype *BookType) UnstageVoid(stage *StageStruct) {
	delete(stage.BookTypes, booktype)
	delete(stage.BookTypes_mapString, booktype.Name)
}

// commit booktype to the back repo (if it is already staged)
func (booktype *BookType) Commit(stage *StageStruct) *BookType {
	if _, ok := stage.BookTypes[booktype]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitBookType(booktype)
		}
	}
	return booktype
}

func (booktype *BookType) CommitVoid(stage *StageStruct) {
	booktype.Commit(stage)
}

// Checkout booktype to the back repo (if it is already staged)
func (booktype *BookType) Checkout(stage *StageStruct) *BookType {
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
func (books *Books) Stage(stage *StageStruct) *Books {
	stage.Bookss[books] = __member
	stage.Bookss_mapString[books.Name] = books

	return books
}

// Unstage removes books off the model stage
func (books *Books) Unstage(stage *StageStruct) *Books {
	delete(stage.Bookss, books)
	delete(stage.Bookss_mapString, books.Name)
	return books
}

// UnstageVoid removes books off the model stage
func (books *Books) UnstageVoid(stage *StageStruct) {
	delete(stage.Bookss, books)
	delete(stage.Bookss_mapString, books.Name)
}

// commit books to the back repo (if it is already staged)
func (books *Books) Commit(stage *StageStruct) *Books {
	if _, ok := stage.Bookss[books]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitBooks(books)
		}
	}
	return books
}

func (books *Books) CommitVoid(stage *StageStruct) {
	books.Commit(stage)
}

// Checkout books to the back repo (if it is already staged)
func (books *Books) Checkout(stage *StageStruct) *Books {
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
func (credit *Credit) Stage(stage *StageStruct) *Credit {
	stage.Credits[credit] = __member
	stage.Credits_mapString[credit.Name] = credit

	return credit
}

// Unstage removes credit off the model stage
func (credit *Credit) Unstage(stage *StageStruct) *Credit {
	delete(stage.Credits, credit)
	delete(stage.Credits_mapString, credit.Name)
	return credit
}

// UnstageVoid removes credit off the model stage
func (credit *Credit) UnstageVoid(stage *StageStruct) {
	delete(stage.Credits, credit)
	delete(stage.Credits_mapString, credit.Name)
}

// commit credit to the back repo (if it is already staged)
func (credit *Credit) Commit(stage *StageStruct) *Credit {
	if _, ok := stage.Credits[credit]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitCredit(credit)
		}
	}
	return credit
}

func (credit *Credit) CommitVoid(stage *StageStruct) {
	credit.Commit(stage)
}

// Checkout credit to the back repo (if it is already staged)
func (credit *Credit) Checkout(stage *StageStruct) *Credit {
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
func (link *Link) Stage(stage *StageStruct) *Link {
	stage.Links[link] = __member
	stage.Links_mapString[link.Name] = link

	return link
}

// Unstage removes link off the model stage
func (link *Link) Unstage(stage *StageStruct) *Link {
	delete(stage.Links, link)
	delete(stage.Links_mapString, link.Name)
	return link
}

// UnstageVoid removes link off the model stage
func (link *Link) UnstageVoid(stage *StageStruct) {
	delete(stage.Links, link)
	delete(stage.Links_mapString, link.Name)
}

// commit link to the back repo (if it is already staged)
func (link *Link) Commit(stage *StageStruct) *Link {
	if _, ok := stage.Links[link]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitLink(link)
		}
	}
	return link
}

func (link *Link) CommitVoid(stage *StageStruct) {
	link.Commit(stage)
}

// Checkout link to the back repo (if it is already staged)
func (link *Link) Checkout(stage *StageStruct) *Link {
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

func (stage *StageStruct) Reset() { // insertion point for array reset
	stage.BookTypes = make(map[*BookType]any)
	stage.BookTypes_mapString = make(map[string]*BookType)

	stage.Bookss = make(map[*Books]any)
	stage.Bookss_mapString = make(map[string]*Books)

	stage.Credits = make(map[*Credit]any)
	stage.Credits_mapString = make(map[string]*Credit)

	stage.Links = make(map[*Link]any)
	stage.Links_mapString = make(map[string]*Link)

}

func (stage *StageStruct) Nil() { // insertion point for array nil
	stage.BookTypes = nil
	stage.BookTypes_mapString = nil

	stage.Bookss = nil
	stage.Bookss_mapString = nil

	stage.Credits = nil
	stage.Credits_mapString = nil

	stage.Links = nil
	stage.Links_mapString = nil

}

func (stage *StageStruct) Unstage() { // insertion point for array nil
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
	CommitVoid(*StageStruct)
	UnstageVoid(stage *StageStruct)
	comparable
}

func CompareGongstructByName[T PointerToGongstruct](a, b T) int {
	return cmp.Compare(a.GetName(), b.GetName())
}

func SortGongstructSetByName[T PointerToGongstruct](set map[T]any) (sortedSlice []T) {

	sortedSlice = maps.Keys(set)
	slices.SortFunc(sortedSlice, CompareGongstructByName)

	return
}

func GetGongstrucsSorted[T PointerToGongstruct](stage *StageStruct) (sortedSlice []T) {

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
func GongGetSet[Type GongstructSet](stage *StageStruct) *Type {
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
func GongGetMap[Type GongstructMapString](stage *StageStruct) *Type {
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
func GetGongstructInstancesSet[Type Gongstruct](stage *StageStruct) *map[*Type]any {
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
func GetGongstructInstancesSetFromPointerType[Type PointerToGongstruct](stage *StageStruct) *map[Type]any {
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
func GetGongstructInstancesMap[Type Gongstruct](stage *StageStruct) *map[string]*Type {
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
func GetPointerReverseMap[Start, End Gongstruct](fieldname string, stage *StageStruct) map[*End][]*Start {

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
func GetSliceOfPointersReverseMap[Start, End Gongstruct](fieldname string, stage *StageStruct) map[*End]*Start {

	var ret Start

	switch any(ret).(type) {
	// insertion point of functions that provide maps for reverse associations
	// reverse maps of direct associations of BookType
	case BookType:
		switch fieldname {
		// insertion point for per direct association field
		case "Credit":
			res := make(map[*Credit]*BookType)
			for booktype := range stage.BookTypes {
				for _, credit_ := range booktype.Credit {
					res[credit_] = booktype
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of Books
	case Books:
		switch fieldname {
		// insertion point for per direct association field
		case "Book":
			res := make(map[*BookType]*Books)
			for books := range stage.Bookss {
				for _, booktype_ := range books.Book {
					res[booktype_] = books
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of Credit
	case Credit:
		switch fieldname {
		// insertion point for per direct association field
		case "Link":
			res := make(map[*Link]*Credit)
			for credit := range stage.Credits {
				for _, link_ := range credit.Link {
					res[link_] = credit
				}
			}
			return any(res).(map[*End]*Start)
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
		res = []string{"Name", "Edition", "Isbn", "Bestseller", "Credit", "Title", "Author", "Year", "Format"}
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
		res = []string{"Name", "Edition", "Isbn", "Bestseller", "Credit", "Title", "Author", "Year", "Format"}
	case *Books:
		res = []string{"Name", "Book"}
	case *Credit:
		res = []string{"Name", "Page", "Credit_type", "Link", "Credit_words", "Credit_symbol"}
	case *Link:
		res = []string{"Name", "NameXSD", "EnclosedText"}
	}
	return
}

func GetFieldStringValueFromPointer[Type PointerToGongstruct](instance Type, fieldName string) (res string) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
	case *BookType:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Edition":
			res = inferedInstance.Edition
		case "Isbn":
			res = inferedInstance.Isbn
		case "Bestseller":
			res = fmt.Sprintf("%t", inferedInstance.Bestseller)
		case "Credit":
			for idx, __instance__ := range inferedInstance.Credit {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "Title":
			res = inferedInstance.Title
		case "Author":
			res = inferedInstance.Author
		case "Year":
			res = fmt.Sprintf("%d", inferedInstance.Year)
		case "Format":
			res = inferedInstance.Format
		}
	case *Books:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Book":
			for idx, __instance__ := range inferedInstance.Book {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case *Credit:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Page":
			res = fmt.Sprintf("%d", inferedInstance.Page)
		case "Credit_type":
			res = inferedInstance.Credit_type
		case "Link":
			for idx, __instance__ := range inferedInstance.Link {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "Credit_words":
			res = inferedInstance.Credit_words
		case "Credit_symbol":
			res = inferedInstance.Credit_symbol
		}
	case *Link:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "NameXSD":
			res = inferedInstance.NameXSD
		case "EnclosedText":
			res = inferedInstance.EnclosedText
		}
	default:
		_ = inferedInstance
	}
	return
}

func GetFieldStringValue[Type Gongstruct](instance Type, fieldName string) (res string) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
	case BookType:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Edition":
			res = inferedInstance.Edition
		case "Isbn":
			res = inferedInstance.Isbn
		case "Bestseller":
			res = fmt.Sprintf("%t", inferedInstance.Bestseller)
		case "Credit":
			for idx, __instance__ := range inferedInstance.Credit {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "Title":
			res = inferedInstance.Title
		case "Author":
			res = inferedInstance.Author
		case "Year":
			res = fmt.Sprintf("%d", inferedInstance.Year)
		case "Format":
			res = inferedInstance.Format
		}
	case Books:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Book":
			for idx, __instance__ := range inferedInstance.Book {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		}
	case Credit:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Page":
			res = fmt.Sprintf("%d", inferedInstance.Page)
		case "Credit_type":
			res = inferedInstance.Credit_type
		case "Link":
			for idx, __instance__ := range inferedInstance.Link {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "Credit_words":
			res = inferedInstance.Credit_words
		case "Credit_symbol":
			res = inferedInstance.Credit_symbol
		}
	case Link:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "NameXSD":
			res = inferedInstance.NameXSD
		case "EnclosedText":
			res = inferedInstance.EnclosedText
		}
	default:
		_ = inferedInstance
	}
	return
}

// Last line of the template
