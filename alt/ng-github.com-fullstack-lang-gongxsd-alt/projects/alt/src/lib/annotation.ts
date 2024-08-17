// generated code - do not edit

import { AnnotationAPI } from './annotation-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { Documentation } from './documentation'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Annotation {

	static GONGSTRUCT_NAME = "Annotation"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for pointers and slices of pointers declarations
	Documentations: Array<Documentation> = []
}

export function CopyAnnotationToAnnotationAPI(annotation: Annotation, annotationAPI: AnnotationAPI) {

	annotationAPI.CreatedAt = annotation.CreatedAt
	annotationAPI.DeletedAt = annotation.DeletedAt
	annotationAPI.ID = annotation.ID

	// insertion point for basic fields copy operations
	annotationAPI.Name = annotation.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	annotationAPI.AnnotationPointersEncoding.Documentations = []
	for (let _documentation of annotation.Documentations) {
		annotationAPI.AnnotationPointersEncoding.Documentations.push(_documentation.ID)
	}

}

// CopyAnnotationAPIToAnnotation update basic, pointers and slice of pointers fields of annotation
// from respectively the basic fields and encoded fields of pointers and slices of pointers of annotationAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyAnnotationAPIToAnnotation(annotationAPI: AnnotationAPI, annotation: Annotation, frontRepo: FrontRepo) {

	annotation.CreatedAt = annotationAPI.CreatedAt
	annotation.DeletedAt = annotationAPI.DeletedAt
	annotation.ID = annotationAPI.ID

	// insertion point for basic fields copy operations
	annotation.Name = annotationAPI.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	annotation.Documentations = new Array<Documentation>()
	for (let _id of annotationAPI.AnnotationPointersEncoding.Documentations) {
		let _documentation = frontRepo.map_ID_Documentation.get(_id)
		if (_documentation != undefined) {
			annotation.Documentations.push(_documentation!)
		}
	}
}
