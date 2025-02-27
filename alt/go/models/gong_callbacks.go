// generated code - do not edit
package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Annotation:
		if stage.OnAfterAnnotationCreateCallback != nil {
			stage.OnAfterAnnotationCreateCallback.OnAfterCreate(stage, target)
		}
	case *ComplexContent:
		if stage.OnAfterComplexContentCreateCallback != nil {
			stage.OnAfterComplexContentCreateCallback.OnAfterCreate(stage, target)
		}
	case *ComplexType:
		if stage.OnAfterComplexTypeCreateCallback != nil {
			stage.OnAfterComplexTypeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Documentation:
		if stage.OnAfterDocumentationCreateCallback != nil {
			stage.OnAfterDocumentationCreateCallback.OnAfterCreate(stage, target)
		}
	case *Schema:
		if stage.OnAfterSchemaCreateCallback != nil {
			stage.OnAfterSchemaCreateCallback.OnAfterCreate(stage, target)
		}
	default:
		_ = target
	}
}

// AfterUpdateFromFront is called after a update from front
func AfterUpdateFromFront[Type Gongstruct](stage *StageStruct, old, new *Type) {

	switch oldTarget := any(old).(type) {
	// insertion point
	case *Annotation:
		newTarget := any(new).(*Annotation)
		if stage.OnAfterAnnotationUpdateCallback != nil {
			stage.OnAfterAnnotationUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ComplexContent:
		newTarget := any(new).(*ComplexContent)
		if stage.OnAfterComplexContentUpdateCallback != nil {
			stage.OnAfterComplexContentUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ComplexType:
		newTarget := any(new).(*ComplexType)
		if stage.OnAfterComplexTypeUpdateCallback != nil {
			stage.OnAfterComplexTypeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Documentation:
		newTarget := any(new).(*Documentation)
		if stage.OnAfterDocumentationUpdateCallback != nil {
			stage.OnAfterDocumentationUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Schema:
		newTarget := any(new).(*Schema)
		if stage.OnAfterSchemaUpdateCallback != nil {
			stage.OnAfterSchemaUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	default:
		_ = oldTarget
	}
}

// AfterDeleteFromFront is called after a delete from front
func AfterDeleteFromFront[Type Gongstruct](stage *StageStruct, staged, front *Type) {

	switch front := any(front).(type) {
	// insertion point
	case *Annotation:
		if stage.OnAfterAnnotationDeleteCallback != nil {
			staged := any(staged).(*Annotation)
			stage.OnAfterAnnotationDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ComplexContent:
		if stage.OnAfterComplexContentDeleteCallback != nil {
			staged := any(staged).(*ComplexContent)
			stage.OnAfterComplexContentDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ComplexType:
		if stage.OnAfterComplexTypeDeleteCallback != nil {
			staged := any(staged).(*ComplexType)
			stage.OnAfterComplexTypeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Documentation:
		if stage.OnAfterDocumentationDeleteCallback != nil {
			staged := any(staged).(*Documentation)
			stage.OnAfterDocumentationDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Schema:
		if stage.OnAfterSchemaDeleteCallback != nil {
			staged := any(staged).(*Schema)
			stage.OnAfterSchemaDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	default:
		_ = front
	}
}

// AfterReadFromFront is called after a Read from front
func AfterReadFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Annotation:
		if stage.OnAfterAnnotationReadCallback != nil {
			stage.OnAfterAnnotationReadCallback.OnAfterRead(stage, target)
		}
	case *ComplexContent:
		if stage.OnAfterComplexContentReadCallback != nil {
			stage.OnAfterComplexContentReadCallback.OnAfterRead(stage, target)
		}
	case *ComplexType:
		if stage.OnAfterComplexTypeReadCallback != nil {
			stage.OnAfterComplexTypeReadCallback.OnAfterRead(stage, target)
		}
	case *Documentation:
		if stage.OnAfterDocumentationReadCallback != nil {
			stage.OnAfterDocumentationReadCallback.OnAfterRead(stage, target)
		}
	case *Schema:
		if stage.OnAfterSchemaReadCallback != nil {
			stage.OnAfterSchemaReadCallback.OnAfterRead(stage, target)
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
	case *Annotation:
		stage.OnAfterAnnotationUpdateCallback = any(callback).(OnAfterUpdateInterface[Annotation])
	
	case *ComplexContent:
		stage.OnAfterComplexContentUpdateCallback = any(callback).(OnAfterUpdateInterface[ComplexContent])
	
	case *ComplexType:
		stage.OnAfterComplexTypeUpdateCallback = any(callback).(OnAfterUpdateInterface[ComplexType])
	
	case *Documentation:
		stage.OnAfterDocumentationUpdateCallback = any(callback).(OnAfterUpdateInterface[Documentation])
	
	case *Schema:
		stage.OnAfterSchemaUpdateCallback = any(callback).(OnAfterUpdateInterface[Schema])
	
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Annotation:
		stage.OnAfterAnnotationCreateCallback = any(callback).(OnAfterCreateInterface[Annotation])
	
	case *ComplexContent:
		stage.OnAfterComplexContentCreateCallback = any(callback).(OnAfterCreateInterface[ComplexContent])
	
	case *ComplexType:
		stage.OnAfterComplexTypeCreateCallback = any(callback).(OnAfterCreateInterface[ComplexType])
	
	case *Documentation:
		stage.OnAfterDocumentationCreateCallback = any(callback).(OnAfterCreateInterface[Documentation])
	
	case *Schema:
		stage.OnAfterSchemaCreateCallback = any(callback).(OnAfterCreateInterface[Schema])
	
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Annotation:
		stage.OnAfterAnnotationDeleteCallback = any(callback).(OnAfterDeleteInterface[Annotation])
	
	case *ComplexContent:
		stage.OnAfterComplexContentDeleteCallback = any(callback).(OnAfterDeleteInterface[ComplexContent])
	
	case *ComplexType:
		stage.OnAfterComplexTypeDeleteCallback = any(callback).(OnAfterDeleteInterface[ComplexType])
	
	case *Documentation:
		stage.OnAfterDocumentationDeleteCallback = any(callback).(OnAfterDeleteInterface[Documentation])
	
	case *Schema:
		stage.OnAfterSchemaDeleteCallback = any(callback).(OnAfterDeleteInterface[Schema])
	
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Annotation:
		stage.OnAfterAnnotationReadCallback = any(callback).(OnAfterReadInterface[Annotation])
	
	case *ComplexContent:
		stage.OnAfterComplexContentReadCallback = any(callback).(OnAfterReadInterface[ComplexContent])
	
	case *ComplexType:
		stage.OnAfterComplexTypeReadCallback = any(callback).(OnAfterReadInterface[ComplexType])
	
	case *Documentation:
		stage.OnAfterDocumentationReadCallback = any(callback).(OnAfterReadInterface[Documentation])
	
	case *Schema:
		stage.OnAfterSchemaReadCallback = any(callback).(OnAfterReadInterface[Schema])
	
	}
}
