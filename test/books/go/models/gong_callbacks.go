// generated code - do not edit
package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *BookDetailsGroup:
		if stage.OnAfterBookDetailsGroupCreateCallback != nil {
			stage.OnAfterBookDetailsGroupCreateCallback.OnAfterCreate(stage, target)
		}
	case *BookType:
		if stage.OnAfterBookTypeCreateCallback != nil {
			stage.OnAfterBookTypeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Books:
		if stage.OnAfterBooksCreateCallback != nil {
			stage.OnAfterBooksCreateCallback.OnAfterCreate(stage, target)
		}
	case *CommonAttributes:
		if stage.OnAfterCommonAttributesCreateCallback != nil {
			stage.OnAfterCommonAttributesCreateCallback.OnAfterCreate(stage, target)
		}
	case *Credit:
		if stage.OnAfterCreditCreateCallback != nil {
			stage.OnAfterCreditCreateCallback.OnAfterCreate(stage, target)
		}
	case *ExtendedAttributes:
		if stage.OnAfterExtendedAttributesCreateCallback != nil {
			stage.OnAfterExtendedAttributesCreateCallback.OnAfterCreate(stage, target)
		}
	case *Link:
		if stage.OnAfterLinkCreateCallback != nil {
			stage.OnAfterLinkCreateCallback.OnAfterCreate(stage, target)
		}
	default:
		_ = target
	}
}

// AfterUpdateFromFront is called after a update from front
func AfterUpdateFromFront[Type Gongstruct](stage *StageStruct, old, new *Type) {

	switch oldTarget := any(old).(type) {
	// insertion point
	case *BookDetailsGroup:
		newTarget := any(new).(*BookDetailsGroup)
		if stage.OnAfterBookDetailsGroupUpdateCallback != nil {
			stage.OnAfterBookDetailsGroupUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *BookType:
		newTarget := any(new).(*BookType)
		if stage.OnAfterBookTypeUpdateCallback != nil {
			stage.OnAfterBookTypeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Books:
		newTarget := any(new).(*Books)
		if stage.OnAfterBooksUpdateCallback != nil {
			stage.OnAfterBooksUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *CommonAttributes:
		newTarget := any(new).(*CommonAttributes)
		if stage.OnAfterCommonAttributesUpdateCallback != nil {
			stage.OnAfterCommonAttributesUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Credit:
		newTarget := any(new).(*Credit)
		if stage.OnAfterCreditUpdateCallback != nil {
			stage.OnAfterCreditUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ExtendedAttributes:
		newTarget := any(new).(*ExtendedAttributes)
		if stage.OnAfterExtendedAttributesUpdateCallback != nil {
			stage.OnAfterExtendedAttributesUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Link:
		newTarget := any(new).(*Link)
		if stage.OnAfterLinkUpdateCallback != nil {
			stage.OnAfterLinkUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	default:
		_ = oldTarget
	}
}

// AfterDeleteFromFront is called after a delete from front
func AfterDeleteFromFront[Type Gongstruct](stage *StageStruct, staged, front *Type) {

	switch front := any(front).(type) {
	// insertion point
	case *BookDetailsGroup:
		if stage.OnAfterBookDetailsGroupDeleteCallback != nil {
			staged := any(staged).(*BookDetailsGroup)
			stage.OnAfterBookDetailsGroupDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *BookType:
		if stage.OnAfterBookTypeDeleteCallback != nil {
			staged := any(staged).(*BookType)
			stage.OnAfterBookTypeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Books:
		if stage.OnAfterBooksDeleteCallback != nil {
			staged := any(staged).(*Books)
			stage.OnAfterBooksDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *CommonAttributes:
		if stage.OnAfterCommonAttributesDeleteCallback != nil {
			staged := any(staged).(*CommonAttributes)
			stage.OnAfterCommonAttributesDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Credit:
		if stage.OnAfterCreditDeleteCallback != nil {
			staged := any(staged).(*Credit)
			stage.OnAfterCreditDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ExtendedAttributes:
		if stage.OnAfterExtendedAttributesDeleteCallback != nil {
			staged := any(staged).(*ExtendedAttributes)
			stage.OnAfterExtendedAttributesDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Link:
		if stage.OnAfterLinkDeleteCallback != nil {
			staged := any(staged).(*Link)
			stage.OnAfterLinkDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	default:
		_ = front
	}
}

// AfterReadFromFront is called after a Read from front
func AfterReadFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *BookDetailsGroup:
		if stage.OnAfterBookDetailsGroupReadCallback != nil {
			stage.OnAfterBookDetailsGroupReadCallback.OnAfterRead(stage, target)
		}
	case *BookType:
		if stage.OnAfterBookTypeReadCallback != nil {
			stage.OnAfterBookTypeReadCallback.OnAfterRead(stage, target)
		}
	case *Books:
		if stage.OnAfterBooksReadCallback != nil {
			stage.OnAfterBooksReadCallback.OnAfterRead(stage, target)
		}
	case *CommonAttributes:
		if stage.OnAfterCommonAttributesReadCallback != nil {
			stage.OnAfterCommonAttributesReadCallback.OnAfterRead(stage, target)
		}
	case *Credit:
		if stage.OnAfterCreditReadCallback != nil {
			stage.OnAfterCreditReadCallback.OnAfterRead(stage, target)
		}
	case *ExtendedAttributes:
		if stage.OnAfterExtendedAttributesReadCallback != nil {
			stage.OnAfterExtendedAttributesReadCallback.OnAfterRead(stage, target)
		}
	case *Link:
		if stage.OnAfterLinkReadCallback != nil {
			stage.OnAfterLinkReadCallback.OnAfterRead(stage, target)
		}
	default:
		_ = target
	}
}

// SetCallbackAfterUpdateFromFront is a function to set up callback that is robust to refactoring
func SetCallbackAfterUpdateFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterUpdateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *BookDetailsGroup:
		stage.OnAfterBookDetailsGroupUpdateCallback = any(callback).(OnAfterUpdateInterface[BookDetailsGroup])
	
	case *BookType:
		stage.OnAfterBookTypeUpdateCallback = any(callback).(OnAfterUpdateInterface[BookType])
	
	case *Books:
		stage.OnAfterBooksUpdateCallback = any(callback).(OnAfterUpdateInterface[Books])
	
	case *CommonAttributes:
		stage.OnAfterCommonAttributesUpdateCallback = any(callback).(OnAfterUpdateInterface[CommonAttributes])
	
	case *Credit:
		stage.OnAfterCreditUpdateCallback = any(callback).(OnAfterUpdateInterface[Credit])
	
	case *ExtendedAttributes:
		stage.OnAfterExtendedAttributesUpdateCallback = any(callback).(OnAfterUpdateInterface[ExtendedAttributes])
	
	case *Link:
		stage.OnAfterLinkUpdateCallback = any(callback).(OnAfterUpdateInterface[Link])
	
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *BookDetailsGroup:
		stage.OnAfterBookDetailsGroupCreateCallback = any(callback).(OnAfterCreateInterface[BookDetailsGroup])
	
	case *BookType:
		stage.OnAfterBookTypeCreateCallback = any(callback).(OnAfterCreateInterface[BookType])
	
	case *Books:
		stage.OnAfterBooksCreateCallback = any(callback).(OnAfterCreateInterface[Books])
	
	case *CommonAttributes:
		stage.OnAfterCommonAttributesCreateCallback = any(callback).(OnAfterCreateInterface[CommonAttributes])
	
	case *Credit:
		stage.OnAfterCreditCreateCallback = any(callback).(OnAfterCreateInterface[Credit])
	
	case *ExtendedAttributes:
		stage.OnAfterExtendedAttributesCreateCallback = any(callback).(OnAfterCreateInterface[ExtendedAttributes])
	
	case *Link:
		stage.OnAfterLinkCreateCallback = any(callback).(OnAfterCreateInterface[Link])
	
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *BookDetailsGroup:
		stage.OnAfterBookDetailsGroupDeleteCallback = any(callback).(OnAfterDeleteInterface[BookDetailsGroup])
	
	case *BookType:
		stage.OnAfterBookTypeDeleteCallback = any(callback).(OnAfterDeleteInterface[BookType])
	
	case *Books:
		stage.OnAfterBooksDeleteCallback = any(callback).(OnAfterDeleteInterface[Books])
	
	case *CommonAttributes:
		stage.OnAfterCommonAttributesDeleteCallback = any(callback).(OnAfterDeleteInterface[CommonAttributes])
	
	case *Credit:
		stage.OnAfterCreditDeleteCallback = any(callback).(OnAfterDeleteInterface[Credit])
	
	case *ExtendedAttributes:
		stage.OnAfterExtendedAttributesDeleteCallback = any(callback).(OnAfterDeleteInterface[ExtendedAttributes])
	
	case *Link:
		stage.OnAfterLinkDeleteCallback = any(callback).(OnAfterDeleteInterface[Link])
	
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *BookDetailsGroup:
		stage.OnAfterBookDetailsGroupReadCallback = any(callback).(OnAfterReadInterface[BookDetailsGroup])
	
	case *BookType:
		stage.OnAfterBookTypeReadCallback = any(callback).(OnAfterReadInterface[BookType])
	
	case *Books:
		stage.OnAfterBooksReadCallback = any(callback).(OnAfterReadInterface[Books])
	
	case *CommonAttributes:
		stage.OnAfterCommonAttributesReadCallback = any(callback).(OnAfterReadInterface[CommonAttributes])
	
	case *Credit:
		stage.OnAfterCreditReadCallback = any(callback).(OnAfterReadInterface[Credit])
	
	case *ExtendedAttributes:
		stage.OnAfterExtendedAttributesReadCallback = any(callback).(OnAfterReadInterface[ExtendedAttributes])
	
	case *Link:
		stage.OnAfterLinkReadCallback = any(callback).(OnAfterReadInterface[Link])
	
	}
}
