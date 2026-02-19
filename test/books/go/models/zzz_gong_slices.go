// generated code - do not edit
package models

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

var (
	__GongSliceTemplate_time__dummyDeclaration time.Duration
	_                                          = __GongSliceTemplate_time__dummyDeclaration
)

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct BookType
	// insertion point per field
	stage.BookType_Credit_reverseMap = make(map[*Credit]*BookType)
	for booktype := range stage.BookTypes {
		_ = booktype
		for _, _credit := range booktype.Credit {
			stage.BookType_Credit_reverseMap[_credit] = booktype
		}
	}

	// Compute reverse map for named struct Books
	// insertion point per field
	stage.Books_Book_reverseMap = make(map[*BookType]*Books)
	for books := range stage.Bookss {
		_ = books
		for _, _booktype := range books.Book {
			stage.Books_Book_reverseMap[_booktype] = books
		}
	}

	// Compute reverse map for named struct Credit
	// insertion point per field
	stage.Credit_Link_reverseMap = make(map[*Link]*Credit)
	for credit := range stage.Credits {
		_ = credit
		for _, _link := range credit.Link {
			stage.Credit_Link_reverseMap[_link] = credit
		}
	}

	// Compute reverse map for named struct Link
	// insertion point per field

	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
	for instance := range stage.BookTypes {
		res = append(res, instance)
	}

	for instance := range stage.Bookss {
		res = append(res, instance)
	}

	for instance := range stage.Credits {
		res = append(res, instance)
	}

	for instance := range stage.Links {
		res = append(res, instance)
	}

	return
}

// insertion point per named struct
func (booktype *BookType) GongCopy() GongstructIF {
	newInstance := *booktype
	return &newInstance
}

func (books *Books) GongCopy() GongstructIF {
	newInstance := *books
	return &newInstance
}

func (credit *Credit) GongCopy() GongstructIF {
	newInstance := *credit
	return &newInstance
}

func (link *Link) GongCopy() GongstructIF {
	newInstance := *link
	return &newInstance
}

func (stage *Stage) ComputeForwardAndBackwardCommits() {
	var lenNewInstances int
	var lenModifiedInstances int
	var lenDeletedInstances int

	var newInstancesSlice []string
	var fieldsEditSlice []string
	var deletedInstancesSlice []string

	var newInstancesReverseSlice []string
	var fieldsEditReverseSlice []string
	var deletedInstancesReverseSlice []string

	// first clean the staging area to remove non staged instances
	// from pointers fields and slices of pointers fields
	stage.Clean()

	// insertion point per named struct
	var booktypes_newInstances []*BookType
	var booktypes_deletedInstances []*BookType

	// parse all staged instances and check if they have a reference
	for booktype := range stage.BookTypes {
		if ref, ok := stage.BookTypes_reference[booktype]; !ok {
			booktypes_newInstances = append(booktypes_newInstances, booktype)
			newInstancesSlice = append(newInstancesSlice, booktype.GongMarshallIdentifier(stage))
			if stage.BookTypes_referenceOrder == nil {
				stage.BookTypes_referenceOrder = make(map[*BookType]uint)
			}
			stage.BookTypes_referenceOrder[booktype] = stage.BookTypeMap_Staged_Order[booktype]
			newInstancesReverseSlice = append(newInstancesReverseSlice, booktype.GongMarshallUnstaging(stage))
			delete(stage.BookTypes_referenceOrder, booktype)
			fieldInitializers, pointersInitializations := booktype.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.BookTypeMap_Staged_Order[ref] = stage.BookTypeMap_Staged_Order[booktype]
			diffs := booktype.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, booktype)
			delete(stage.BookTypeMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", booktype.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.BookTypes_reference {
		if _, ok := stage.BookTypes[ref]; !ok {
			booktypes_deletedInstances = append(booktypes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(booktypes_newInstances)
	lenDeletedInstances += len(booktypes_deletedInstances)
	var bookss_newInstances []*Books
	var bookss_deletedInstances []*Books

	// parse all staged instances and check if they have a reference
	for books := range stage.Bookss {
		if ref, ok := stage.Bookss_reference[books]; !ok {
			bookss_newInstances = append(bookss_newInstances, books)
			newInstancesSlice = append(newInstancesSlice, books.GongMarshallIdentifier(stage))
			if stage.Bookss_referenceOrder == nil {
				stage.Bookss_referenceOrder = make(map[*Books]uint)
			}
			stage.Bookss_referenceOrder[books] = stage.BooksMap_Staged_Order[books]
			newInstancesReverseSlice = append(newInstancesReverseSlice, books.GongMarshallUnstaging(stage))
			delete(stage.Bookss_referenceOrder, books)
			fieldInitializers, pointersInitializations := books.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.BooksMap_Staged_Order[ref] = stage.BooksMap_Staged_Order[books]
			diffs := books.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, books)
			delete(stage.BooksMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", books.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.Bookss_reference {
		if _, ok := stage.Bookss[ref]; !ok {
			bookss_deletedInstances = append(bookss_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(bookss_newInstances)
	lenDeletedInstances += len(bookss_deletedInstances)
	var credits_newInstances []*Credit
	var credits_deletedInstances []*Credit

	// parse all staged instances and check if they have a reference
	for credit := range stage.Credits {
		if ref, ok := stage.Credits_reference[credit]; !ok {
			credits_newInstances = append(credits_newInstances, credit)
			newInstancesSlice = append(newInstancesSlice, credit.GongMarshallIdentifier(stage))
			if stage.Credits_referenceOrder == nil {
				stage.Credits_referenceOrder = make(map[*Credit]uint)
			}
			stage.Credits_referenceOrder[credit] = stage.CreditMap_Staged_Order[credit]
			newInstancesReverseSlice = append(newInstancesReverseSlice, credit.GongMarshallUnstaging(stage))
			delete(stage.Credits_referenceOrder, credit)
			fieldInitializers, pointersInitializations := credit.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.CreditMap_Staged_Order[ref] = stage.CreditMap_Staged_Order[credit]
			diffs := credit.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, credit)
			delete(stage.CreditMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", credit.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.Credits_reference {
		if _, ok := stage.Credits[ref]; !ok {
			credits_deletedInstances = append(credits_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(credits_newInstances)
	lenDeletedInstances += len(credits_deletedInstances)
	var links_newInstances []*Link
	var links_deletedInstances []*Link

	// parse all staged instances and check if they have a reference
	for link := range stage.Links {
		if ref, ok := stage.Links_reference[link]; !ok {
			links_newInstances = append(links_newInstances, link)
			newInstancesSlice = append(newInstancesSlice, link.GongMarshallIdentifier(stage))
			if stage.Links_referenceOrder == nil {
				stage.Links_referenceOrder = make(map[*Link]uint)
			}
			stage.Links_referenceOrder[link] = stage.LinkMap_Staged_Order[link]
			newInstancesReverseSlice = append(newInstancesReverseSlice, link.GongMarshallUnstaging(stage))
			delete(stage.Links_referenceOrder, link)
			fieldInitializers, pointersInitializations := link.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.LinkMap_Staged_Order[ref] = stage.LinkMap_Staged_Order[link]
			diffs := link.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, link)
			delete(stage.LinkMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", link.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.Links_reference {
		if _, ok := stage.Links[ref]; !ok {
			links_deletedInstances = append(links_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(links_newInstances)
	lenDeletedInstances += len(links_deletedInstances)

	if lenNewInstances > 0 || lenDeletedInstances > 0 || lenModifiedInstances > 0 {

		// sort the stmt to have reproductible forward/backward commit
		sort.Strings(newInstancesSlice)
		newInstancesStmt := strings.Join(newInstancesSlice, "")
		sort.Strings(fieldsEditSlice)
		fieldsEditStmt := strings.Join(fieldsEditSlice, "")
		sort.Strings(deletedInstancesSlice)
		deletedInstancesStmt := strings.Join(deletedInstancesSlice, "")

		sort.Strings(newInstancesReverseSlice)
		newInstancesReverseStmt := strings.Join(newInstancesReverseSlice, "")
		sort.Strings(fieldsEditReverseSlice)
		fieldsEditReverseStmt := strings.Join(fieldsEditReverseSlice, "")
		sort.Strings(deletedInstancesReverseSlice)
		deletedInstancesReverseStmt := strings.Join(deletedInstancesReverseSlice, "")

		forwardCommit := newInstancesStmt + fieldsEditStmt + deletedInstancesStmt
		forwardCommit += "\n\tstage.Commit()"
		stage.forwardCommits = append(stage.forwardCommits, forwardCommit)

		backwardCommit := deletedInstancesReverseStmt + fieldsEditReverseStmt + newInstancesReverseStmt
		backwardCommit += "\n\tstage.Commit()"
		// append to the end of the backward commits slice
		stage.backwardCommits = append(stage.backwardCommits, backwardCommit)
	}
}

// ComputeReferenceAndOrders will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReferenceAndOrders() {
	// insertion point per named struct
	stage.BookTypes_reference = make(map[*BookType]*BookType)
	stage.BookTypes_referenceOrder = make(map[*BookType]uint) // diff Unstage needs the reference order
	for instance := range stage.BookTypes {
		stage.BookTypes_reference[instance] = instance.GongCopy().(*BookType)
		stage.BookTypes_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Bookss_reference = make(map[*Books]*Books)
	stage.Bookss_referenceOrder = make(map[*Books]uint) // diff Unstage needs the reference order
	for instance := range stage.Bookss {
		stage.Bookss_reference[instance] = instance.GongCopy().(*Books)
		stage.Bookss_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Credits_reference = make(map[*Credit]*Credit)
	stage.Credits_referenceOrder = make(map[*Credit]uint) // diff Unstage needs the reference order
	for instance := range stage.Credits {
		stage.Credits_reference[instance] = instance.GongCopy().(*Credit)
		stage.Credits_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Links_reference = make(map[*Link]*Link)
	stage.Links_referenceOrder = make(map[*Link]uint) // diff Unstage needs the reference order
	for instance := range stage.Links {
		stage.Links_reference[instance] = instance.GongCopy().(*Link)
		stage.Links_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.recomputeOrders()
}

// GongGetOrder returns the order of the instance in the staging area
// This order is set at staging time, and reflects the order of creation of the instances
// in the staging area
// It is used when rendering slices of GongstructIF to keep a deterministic order
// which is important for frontends such as web frontends
// to avoid unnecessary re-renderings
// insertion point per named struct
func (booktype *BookType) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.BookTypeMap_Staged_Order[booktype]; ok {
		return order
	}
	return stage.BookTypes_referenceOrder[booktype]
}

func (booktype *BookType) GongGetReferenceOrder(stage *Stage) uint {
	return stage.BookTypes_referenceOrder[booktype]
}

func (books *Books) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.BooksMap_Staged_Order[books]; ok {
		return order
	}
	return stage.Bookss_referenceOrder[books]
}

func (books *Books) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Bookss_referenceOrder[books]
}

func (credit *Credit) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.CreditMap_Staged_Order[credit]; ok {
		return order
	}
	return stage.Credits_referenceOrder[credit]
}

func (credit *Credit) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Credits_referenceOrder[credit]
}

func (link *Link) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.LinkMap_Staged_Order[link]; ok {
		return order
	}
	return stage.Links_referenceOrder[link]
}

func (link *Link) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Links_referenceOrder[link]
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (booktype *BookType) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", booktype.GongGetGongstructName(), booktype.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (booktype *BookType) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", booktype.GongGetGongstructName(), booktype.GongGetReferenceOrder(stage))
}

func (books *Books) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", books.GongGetGongstructName(), books.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (books *Books) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", books.GongGetGongstructName(), books.GongGetReferenceOrder(stage))
}

func (credit *Credit) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", credit.GongGetGongstructName(), credit.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (credit *Credit) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", credit.GongGetGongstructName(), credit.GongGetReferenceOrder(stage))
}

func (link *Link) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", link.GongGetGongstructName(), link.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (link *Link) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", link.GongGetGongstructName(), link.GongGetReferenceOrder(stage))
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (booktype *BookType) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", booktype.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "BookType")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", booktype.Name)
	return
}

func (books *Books) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", books.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Books")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", books.Name)
	return
}

func (credit *Credit) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", credit.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Credit")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", credit.Name)
	return
}

func (link *Link) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", link.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Link")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", link.Name)
	return
}

// insertion point for unstaging
func (booktype *BookType) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", booktype.GongGetReferenceIdentifier(stage))
	return
}

func (books *Books) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", books.GongGetReferenceIdentifier(stage))
	return
}

func (credit *Credit) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", credit.GongGetReferenceIdentifier(stage))
	return
}

func (link *Link) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", link.GongGetReferenceIdentifier(stage))
	return
}

// end of template
