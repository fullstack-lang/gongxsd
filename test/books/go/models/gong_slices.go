// generated code - do not edit
package models

// EvictInOtherSlices allows for adherance between
// the gong association model and go.
//
// Says you have a Astruct struct with a slice field "anarrayofb []*Bstruct"
//
// go allows many Astruct instance to have the anarrayofb field that have the
// same pointers. go slices are MANY-MANY association.
//
// With gong it is only ZERO-ONE-MANY associations, a Bstruct can be pointed only
// once by an Astruct instance through a given field. This follows the requirement
// that gong is suited for full stack programming and therefore the association
// is encoded as a reverse pointer (not as a joint table). In gong, a named struct
// is translated in a table and each table is a named struct.
//
// EvictInOtherSlices removes the fields instances from other
// fields of other instance
//
// Note : algo is in O(N)log(N) of nb of Astruct and Bstruct instances
func EvictInOtherSlices[OwningType PointerToGongstruct, FieldType PointerToGongstruct](
	stage *StageStruct,
	owningInstance OwningType,
	sliceField []FieldType,
	fieldName string) {

	// create a map of the field elements
	setOfFieldInstances := make(map[FieldType]any, 0)
	for _, fieldInstance := range sliceField {
		setOfFieldInstances[fieldInstance] = true
	}

	switch owningInstanceInfered := any(owningInstance).(type) {
	// insertion point
	case *BookType:
		// insertion point per field
		if fieldName == "Credit" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*BookType) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*BookType)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Credit).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Credit = _inferedTypeInstance.Credit[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Credit =
								append(_inferedTypeInstance.Credit, any(fieldInstance).(*Credit))
						}
					}
				}
			}
		}

	case *Books:
		// insertion point per field
		if fieldName == "Book" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Books) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Books)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Book).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Book = _inferedTypeInstance.Book[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Book =
								append(_inferedTypeInstance.Book, any(fieldInstance).(*BookType))
						}
					}
				}
			}
		}

	case *Credit:
		// insertion point per field
		if fieldName == "Link" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Credit) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Credit)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Link).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Link = _inferedTypeInstance.Link[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Link =
								append(_inferedTypeInstance.Link, any(fieldInstance).(*Link))
						}
					}
				}
			}
		}

	case *Link:
		// insertion point per field

	default:
		_ = owningInstanceInfered // to avoid "declared and not used" error if no named struct has slices
	}
}

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *StageStruct) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct BookType
	// insertion point per field
	clear(stage.BookType_Credit_reverseMap)
	stage.BookType_Credit_reverseMap = make(map[*Credit]*BookType)
	for booktype := range stage.BookTypes {
		_ = booktype
		for _, _credit := range booktype.Credit {
			stage.BookType_Credit_reverseMap[_credit] = booktype
		}
	}

	// Compute reverse map for named struct Books
	// insertion point per field
	clear(stage.Books_Book_reverseMap)
	stage.Books_Book_reverseMap = make(map[*BookType]*Books)
	for books := range stage.Bookss {
		_ = books
		for _, _booktype := range books.Book {
			stage.Books_Book_reverseMap[_booktype] = books
		}
	}

	// Compute reverse map for named struct Credit
	// insertion point per field
	clear(stage.Credit_Link_reverseMap)
	stage.Credit_Link_reverseMap = make(map[*Link]*Credit)
	for credit := range stage.Credits {
		_ = credit
		for _, _link := range credit.Link {
			stage.Credit_Link_reverseMap[_link] = credit
		}
	}

	// Compute reverse map for named struct Link
	// insertion point per field

}
