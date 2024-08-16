// generated code - do not edit
package models

func IsStaged[Type Gongstruct](stage *StageStruct, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *BookType:
		ok = stage.IsStagedBookType(target)

	case *Books:
		ok = stage.IsStagedBooks(target)

	case *Credit:
		ok = stage.IsStagedCredit(target)

	case *Link:
		ok = stage.IsStagedLink(target)

	default:
		_ = target
	}
	return
}

// insertion point for stage per struct
func (stage *StageStruct) IsStagedBookType(booktype *BookType) (ok bool) {

	_, ok = stage.BookTypes[booktype]

	return
}

func (stage *StageStruct) IsStagedBooks(books *Books) (ok bool) {

	_, ok = stage.Bookss[books]

	return
}

func (stage *StageStruct) IsStagedCredit(credit *Credit) (ok bool) {

	_, ok = stage.Credits[credit]

	return
}

func (stage *StageStruct) IsStagedLink(link *Link) (ok bool) {

	_, ok = stage.Links[link]

	return
}

// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for stage branch
	case *BookType:
		stage.StageBranchBookType(target)

	case *Books:
		stage.StageBranchBooks(target)

	case *Credit:
		stage.StageBranchCredit(target)

	case *Link:
		stage.StageBranchLink(target)

	default:
		_ = target
	}
}

// insertion point for stage branch per struct
func (stage *StageStruct) StageBranchBookType(booktype *BookType) {

	// check if instance is already staged
	if IsStaged(stage, booktype) {
		return
	}

	booktype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _credit := range booktype.Credit {
		StageBranch(stage, _credit)
	}

}

func (stage *StageStruct) StageBranchBooks(books *Books) {

	// check if instance is already staged
	if IsStaged(stage, books) {
		return
	}

	books.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _booktype := range books.Book {
		StageBranch(stage, _booktype)
	}

}

func (stage *StageStruct) StageBranchCredit(credit *Credit) {

	// check if instance is already staged
	if IsStaged(stage, credit) {
		return
	}

	credit.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _link := range credit.Link {
		StageBranch(stage, _link)
	}

}

func (stage *StageStruct) StageBranchLink(link *Link) {

	// check if instance is already staged
	if IsStaged(stage, link) {
		return
	}

	link.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

// CopyBranch stages instance and apply CopyBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func CopyBranch[Type Gongstruct](from *Type) (to *Type) {

	mapOrigCopy := make(map[any]any)
	_ = mapOrigCopy

	switch fromT := any(from).(type) {
	// insertion point for stage branch
	case *BookType:
		toT := CopyBranchBookType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Books:
		toT := CopyBranchBooks(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Credit:
		toT := CopyBranchCredit(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Link:
		toT := CopyBranchLink(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func CopyBranchBookType(mapOrigCopy map[any]any, booktypeFrom *BookType) (booktypeTo *BookType) {

	// booktypeFrom has already been copied
	if _booktypeTo, ok := mapOrigCopy[booktypeFrom]; ok {
		booktypeTo = _booktypeTo.(*BookType)
		return
	}

	booktypeTo = new(BookType)
	mapOrigCopy[booktypeFrom] = booktypeTo
	booktypeFrom.CopyBasicFields(booktypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _credit := range booktypeFrom.Credit {
		booktypeTo.Credit = append(booktypeTo.Credit, CopyBranchCredit(mapOrigCopy, _credit))
	}

	return
}

func CopyBranchBooks(mapOrigCopy map[any]any, booksFrom *Books) (booksTo *Books) {

	// booksFrom has already been copied
	if _booksTo, ok := mapOrigCopy[booksFrom]; ok {
		booksTo = _booksTo.(*Books)
		return
	}

	booksTo = new(Books)
	mapOrigCopy[booksFrom] = booksTo
	booksFrom.CopyBasicFields(booksTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _booktype := range booksFrom.Book {
		booksTo.Book = append(booksTo.Book, CopyBranchBookType(mapOrigCopy, _booktype))
	}

	return
}

func CopyBranchCredit(mapOrigCopy map[any]any, creditFrom *Credit) (creditTo *Credit) {

	// creditFrom has already been copied
	if _creditTo, ok := mapOrigCopy[creditFrom]; ok {
		creditTo = _creditTo.(*Credit)
		return
	}

	creditTo = new(Credit)
	mapOrigCopy[creditFrom] = creditTo
	creditFrom.CopyBasicFields(creditTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _link := range creditFrom.Link {
		creditTo.Link = append(creditTo.Link, CopyBranchLink(mapOrigCopy, _link))
	}

	return
}

func CopyBranchLink(mapOrigCopy map[any]any, linkFrom *Link) (linkTo *Link) {

	// linkFrom has already been copied
	if _linkTo, ok := mapOrigCopy[linkFrom]; ok {
		linkTo = _linkTo.(*Link)
		return
	}

	linkTo = new(Link)
	mapOrigCopy[linkFrom] = linkTo
	linkFrom.CopyBasicFields(linkTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
func UnstageBranch[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for unstage branch
	case *BookType:
		stage.UnstageBranchBookType(target)

	case *Books:
		stage.UnstageBranchBooks(target)

	case *Credit:
		stage.UnstageBranchCredit(target)

	case *Link:
		stage.UnstageBranchLink(target)

	default:
		_ = target
	}
}

// insertion point for unstage branch per struct
func (stage *StageStruct) UnstageBranchBookType(booktype *BookType) {

	// check if instance is already staged
	if !IsStaged(stage, booktype) {
		return
	}

	booktype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _credit := range booktype.Credit {
		UnstageBranch(stage, _credit)
	}

}

func (stage *StageStruct) UnstageBranchBooks(books *Books) {

	// check if instance is already staged
	if !IsStaged(stage, books) {
		return
	}

	books.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _booktype := range books.Book {
		UnstageBranch(stage, _booktype)
	}

}

func (stage *StageStruct) UnstageBranchCredit(credit *Credit) {

	// check if instance is already staged
	if !IsStaged(stage, credit) {
		return
	}

	credit.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _link := range credit.Link {
		UnstageBranch(stage, _link)
	}

}

func (stage *StageStruct) UnstageBranchLink(link *Link) {

	// check if instance is already staged
	if !IsStaged(stage, link) {
		return
	}

	link.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}
