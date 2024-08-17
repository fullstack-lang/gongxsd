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
	
	case *Documentation:
		stage.OnAfterDocumentationReadCallback = any(callback).(OnAfterReadInterface[Documentation])
	
	case *Schema:
		stage.OnAfterSchemaReadCallback = any(callback).(OnAfterReadInterface[Schema])
	
	}
}
