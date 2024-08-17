import { Injectable } from '@angular/core'
import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http'

import { Observable, combineLatest, BehaviorSubject, of } from 'rxjs'

// insertion point sub template for services imports
import { AnnotationAPI } from './annotation-api'
import { Annotation, CopyAnnotationAPIToAnnotation } from './annotation'
import { AnnotationService } from './annotation.service'

import { DocumentationAPI } from './documentation-api'
import { Documentation, CopyDocumentationAPIToDocumentation } from './documentation'
import { DocumentationService } from './documentation.service'

import { SchemaAPI } from './schema-api'
import { Schema, CopySchemaAPIToSchema } from './schema'
import { SchemaService } from './schema.service'


import { BackRepoData } from './back-repo-data'

export const StackType = "github.com/fullstack-lang/gongxsd/alt/go/models"

// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template
	array_Annotations = new Array<Annotation>() // array of front instances
	map_ID_Annotation = new Map<number, Annotation>() // map of front instances

	array_Documentations = new Array<Documentation>() // array of front instances
	map_ID_Documentation = new Map<number, Documentation>() // map of front instances

	array_Schemas = new Array<Schema>() // array of front instances
	map_ID_Schema = new Map<number, Schema>() // map of front instances


	// getFrontArray allows for a get function that is robust to refactoring of the named struct name
	// for instance frontRepo.getArray<Astruct>( Astruct.GONGSTRUCT_NAME), is robust to a refactoring of Astruct identifier
	// contrary to frontRepo.Astructs_array which is not refactored when Astruct identifier is modified
	getFrontArray<Type>(gongStructName: string): Array<Type> {
		switch (gongStructName) {
			// insertion point
			case 'Annotation':
				return this.array_Annotations as unknown as Array<Type>
			case 'Documentation':
				return this.array_Documentations as unknown as Array<Type>
			case 'Schema':
				return this.array_Schemas as unknown as Array<Type>
			default:
				throw new Error("Type not recognized");
		}
	}

	getFrontMap<Type>(gongStructName: string): Map<number, Type> {
		switch (gongStructName) {
			// insertion point
			case 'Annotation':
				return this.map_ID_Annotation as unknown as Map<number, Type>
			case 'Documentation':
				return this.map_ID_Documentation as unknown as Map<number, Type>
			case 'Schema':
				return this.map_ID_Schema as unknown as Map<number, Type>
			default:
				throw new Error("Type not recognized");
		}
	}
}

// the table component is called in different ways
//
// DISPLAY or ASSOCIATION MODE
//
// in ASSOCIATION MODE, it is invoked within a diaglo and a Dialog Data item is used to
// configure the component
// DialogData define the interface for information that is forwarded from the calling instance to 
// the select table
export class DialogData {
	ID: number = 0 // ID of the calling instance

	// the reverse pointer is the name of the generated field on the destination
	// struct of the ONE-MANY association
	ReversePointer: string = "" // field of {{Structname}} that serve as reverse pointer
	OrderingMode: boolean = false // if true, this is for ordering items

	// there are different selection mode : ONE_MANY or MANY_MANY
	SelectionMode: SelectionMode = SelectionMode.ONE_MANY_ASSOCIATION_MODE

	// used if SelectionMode is MANY_MANY_ASSOCIATION_MODE
	//
	// In Gong, a MANY-MANY association is implemented as a ONE-ZERO/ONE followed by a ONE_MANY association
	// 
	// in the MANY_MANY_ASSOCIATION_MODE case, we need also the Struct and the FieldName that are
	// at the end of the ONE-MANY association
	SourceStruct: string = ""	// The "Aclass"
	SourceField: string = "" // the "AnarrayofbUse"
	IntermediateStruct: string = "" // the "AclassBclassUse" 
	IntermediateStructField: string = "" // the "Bclass" as field
	NextAssociationStruct: string = "" // the "Bclass"

	GONG__StackPath: string = ""
}

export enum SelectionMode {
	ONE_MANY_ASSOCIATION_MODE = "ONE_MANY_ASSOCIATION_MODE",
	MANY_MANY_ASSOCIATION_MODE = "MANY_MANY_ASSOCIATION_MODE",
}

//
// observable that fetch all elements of the stack and store them in the FrontRepo
//
@Injectable({
	providedIn: 'root'
})
export class FrontRepoService {

	GONG__StackPath: string = ""
	private socket: WebSocket | undefined

	httpOptions = {
		headers: new HttpHeaders({ 'Content-Type': 'application/json' })
	};

	//
	// Store of all instances of the stack
	//
	frontRepo = new (FrontRepo)

	constructor(
		private http: HttpClient, // insertion point sub template 
		private annotationService: AnnotationService,
		private documentationService: DocumentationService,
		private schemaService: SchemaService,
	) { }

	// postService provides a post function for each struct name
	postService(structName: string, instanceToBePosted: any) {
		let service = this[structName.toLowerCase() + "Service" + "Service" as keyof FrontRepoService]
		let servicePostFunction = service[("post" + structName) as keyof typeof service] as (instance: typeof instanceToBePosted) => Observable<typeof instanceToBePosted>

		servicePostFunction(instanceToBePosted).subscribe(
			instance => {
				let behaviorSubject = instanceToBePosted[(structName + "ServiceChanged") as keyof typeof instanceToBePosted] as unknown as BehaviorSubject<string>
				behaviorSubject.next("post")
			}
		);
	}

	// deleteService provides a delete function for each struct name
	deleteService(structName: string, instanceToBeDeleted: any) {
		let service = this[structName.toLowerCase() + "Service" as keyof FrontRepoService]
		let serviceDeleteFunction = service["delete" + structName as keyof typeof service] as (instance: typeof instanceToBeDeleted) => Observable<typeof instanceToBeDeleted>

		serviceDeleteFunction(instanceToBeDeleted).subscribe(
			instance => {
				let behaviorSubject = instanceToBeDeleted[(structName + "ServiceChanged") as keyof typeof instanceToBeDeleted] as unknown as BehaviorSubject<string>
				behaviorSubject.next("delete")
			}
		);
	}

	// typing of observable can be messy in typescript. Therefore, one force the type
	observableFrontRepo: [
		Observable<null>, // see below for the of(null) observable
		// insertion point sub template 
		Observable<AnnotationAPI[]>,
		Observable<DocumentationAPI[]>,
		Observable<SchemaAPI[]>,
	] = [
			// Using "combineLatest" with a placeholder observable.
			//
			// This allows the typescript compiler to pass when no GongStruct is present in the front API
			//
			// The "of(null)" is a "meaningless" observable that emits a single value (null) and completes.
			// This is used as a workaround to satisfy TypeScript requirements and the "combineLatest" 
			// expectation for a non-empty array of observables.
			of(null), // 
			// insertion point sub template
			this.annotationService.getAnnotations(this.GONG__StackPath, this.frontRepo),
			this.documentationService.getDocumentations(this.GONG__StackPath, this.frontRepo),
			this.schemaService.getSchemas(this.GONG__StackPath, this.frontRepo),
		];

	//
	// pull performs a GET on all struct of the stack and redeem association pointers 
	//
	// This is an observable. Therefore, the control flow forks with
	// - pull() return immediatly the observable
	// - the observable observer, if it subscribe, is called when all GET calls are performs
	pull(GONG__StackPath: string = ""): Observable<FrontRepo> {

		this.GONG__StackPath = GONG__StackPath

		this.observableFrontRepo = [
			of(null), // see above for justification
			// insertion point sub template
			this.annotationService.getAnnotations(this.GONG__StackPath, this.frontRepo),
			this.documentationService.getDocumentations(this.GONG__StackPath, this.frontRepo),
			this.schemaService.getSchemas(this.GONG__StackPath, this.frontRepo),
		]

		return new Observable<FrontRepo>(
			(observer) => {
				combineLatest(
					this.observableFrontRepo
				).subscribe(
					([
						___of_null, // see above for the explanation about of
						// insertion point sub template for declarations 
						annotations_,
						documentations_,
						schemas_,
					]) => {
						let _this = this
						// Typing can be messy with many items. Therefore, type casting is necessary here
						// insertion point sub template for type casting 
						var annotations: AnnotationAPI[]
						annotations = annotations_ as AnnotationAPI[]
						var documentations: DocumentationAPI[]
						documentations = documentations_ as DocumentationAPI[]
						var schemas: SchemaAPI[]
						schemas = schemas_ as SchemaAPI[]

						// 
						// First Step: init map of instances
						// insertion point sub template for init 
						// init the arrays
						this.frontRepo.array_Annotations = []
						this.frontRepo.map_ID_Annotation.clear()

						annotations.forEach(
							annotationAPI => {
								let annotation = new Annotation
								this.frontRepo.array_Annotations.push(annotation)
								this.frontRepo.map_ID_Annotation.set(annotationAPI.ID, annotation)
							}
						)

						// init the arrays
						this.frontRepo.array_Documentations = []
						this.frontRepo.map_ID_Documentation.clear()

						documentations.forEach(
							documentationAPI => {
								let documentation = new Documentation
								this.frontRepo.array_Documentations.push(documentation)
								this.frontRepo.map_ID_Documentation.set(documentationAPI.ID, documentation)
							}
						)

						// init the arrays
						this.frontRepo.array_Schemas = []
						this.frontRepo.map_ID_Schema.clear()

						schemas.forEach(
							schemaAPI => {
								let schema = new Schema
								this.frontRepo.array_Schemas.push(schema)
								this.frontRepo.map_ID_Schema.set(schemaAPI.ID, schema)
							}
						)


						// 
						// Second Step: reddeem front objects
						// insertion point sub template for redeem 
						// fill up front objects
						annotations.forEach(
							annotationAPI => {
								let annotation = this.frontRepo.map_ID_Annotation.get(annotationAPI.ID)
								CopyAnnotationAPIToAnnotation(annotationAPI, annotation!, this.frontRepo)
							}
						)

						// fill up front objects
						documentations.forEach(
							documentationAPI => {
								let documentation = this.frontRepo.map_ID_Documentation.get(documentationAPI.ID)
								CopyDocumentationAPIToDocumentation(documentationAPI, documentation!, this.frontRepo)
							}
						)

						// fill up front objects
						schemas.forEach(
							schemaAPI => {
								let schema = this.frontRepo.map_ID_Schema.get(schemaAPI.ID)
								CopySchemaAPIToSchema(schemaAPI, schema!, this.frontRepo)
							}
						)


						// hand over control flow to observer
						observer.next(this.frontRepo)
					}
				)
			}
		)
	}

	public connectToWebSocket(GONG__StackPath: string): Observable<FrontRepo> {

		this.GONG__StackPath = GONG__StackPath


		let params = new HttpParams().set("GONG__StackPath", this.GONG__StackPath)
		let basePath = 'ws://localhost:8080/api/github.com/fullstack-lang/gongxsd/alt/go/v1/ws/stage'
		let paramString = params.toString()
		let url = `${basePath}?${paramString}`
		this.socket = new WebSocket(url)

		return new Observable(observer => {
			this.socket!.onmessage = event => {
				let _this = this

				const backRepoData = new BackRepoData(JSON.parse(event.data))

				// 
				// First Step: init map of instances
				// insertion point sub template for init 
				// init the arrays
				// insertion point sub template for init 
				// init the arrays
				this.frontRepo.array_Annotations = []
				this.frontRepo.map_ID_Annotation.clear()

				backRepoData.AnnotationAPIs.forEach(
					annotationAPI => {
						let annotation = new Annotation
						this.frontRepo.array_Annotations.push(annotation)
						this.frontRepo.map_ID_Annotation.set(annotationAPI.ID, annotation)
					}
				)

				// init the arrays
				this.frontRepo.array_Documentations = []
				this.frontRepo.map_ID_Documentation.clear()

				backRepoData.DocumentationAPIs.forEach(
					documentationAPI => {
						let documentation = new Documentation
						this.frontRepo.array_Documentations.push(documentation)
						this.frontRepo.map_ID_Documentation.set(documentationAPI.ID, documentation)
					}
				)

				// init the arrays
				this.frontRepo.array_Schemas = []
				this.frontRepo.map_ID_Schema.clear()

				backRepoData.SchemaAPIs.forEach(
					schemaAPI => {
						let schema = new Schema
						this.frontRepo.array_Schemas.push(schema)
						this.frontRepo.map_ID_Schema.set(schemaAPI.ID, schema)
					}
				)


				// 
				// Second Step: reddeem front objects
				// insertion point sub template for redeem 
				// fill up front objects
				// insertion point sub template for redeem 
				// fill up front objects
				backRepoData.AnnotationAPIs.forEach(
					annotationAPI => {
						let annotation = this.frontRepo.map_ID_Annotation.get(annotationAPI.ID)
						CopyAnnotationAPIToAnnotation(annotationAPI, annotation!, this.frontRepo)
					}
				)

				// fill up front objects
				backRepoData.DocumentationAPIs.forEach(
					documentationAPI => {
						let documentation = this.frontRepo.map_ID_Documentation.get(documentationAPI.ID)
						CopyDocumentationAPIToDocumentation(documentationAPI, documentation!, this.frontRepo)
					}
				)

				// fill up front objects
				backRepoData.SchemaAPIs.forEach(
					schemaAPI => {
						let schema = this.frontRepo.map_ID_Schema.get(schemaAPI.ID)
						CopySchemaAPIToSchema(schemaAPI, schema!, this.frontRepo)
					}
				)



				observer.next(this.frontRepo)
			}
			this.socket!.onerror = event => {
				observer.error(event)
			}
			this.socket!.onclose = event => {
				observer.complete()
			}

			return () => {
				this.socket!.close()
			}
		})
	}
}

// insertion point for get unique ID per struct 
export function getAnnotationUniqueID(id: number): number {
	return 31 * id
}
export function getDocumentationUniqueID(id: number): number {
	return 37 * id
}
export function getSchemaUniqueID(id: number): number {
	return 41 * id
}
