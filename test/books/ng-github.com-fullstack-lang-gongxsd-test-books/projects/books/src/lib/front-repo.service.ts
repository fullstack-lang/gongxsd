import { Injectable } from '@angular/core'
import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http'

import { Observable, combineLatest, BehaviorSubject, of } from 'rxjs'

// insertion point sub template for services imports
import { BookDetailsGroupAPI } from './bookdetailsgroup-api'
import { BookDetailsGroup, CopyBookDetailsGroupAPIToBookDetailsGroup } from './bookdetailsgroup'
import { BookDetailsGroupService } from './bookdetailsgroup.service'

import { BookTypeAPI } from './booktype-api'
import { BookType, CopyBookTypeAPIToBookType } from './booktype'
import { BookTypeService } from './booktype.service'

import { BooksAPI } from './books-api'
import { Books, CopyBooksAPIToBooks } from './books'
import { BooksService } from './books.service'

import { CreditAPI } from './credit-api'
import { Credit, CopyCreditAPIToCredit } from './credit'
import { CreditService } from './credit.service'

import { LinkAPI } from './link-api'
import { Link, CopyLinkAPIToLink } from './link'
import { LinkService } from './link.service'


import { BackRepoData } from './back-repo-data'

export const StackType = "github.com/fullstack-lang/gongxsd/test/books/go/models"

// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template
	array_BookDetailsGroups = new Array<BookDetailsGroup>() // array of front instances
	map_ID_BookDetailsGroup = new Map<number, BookDetailsGroup>() // map of front instances

	array_BookTypes = new Array<BookType>() // array of front instances
	map_ID_BookType = new Map<number, BookType>() // map of front instances

	array_Bookss = new Array<Books>() // array of front instances
	map_ID_Books = new Map<number, Books>() // map of front instances

	array_Credits = new Array<Credit>() // array of front instances
	map_ID_Credit = new Map<number, Credit>() // map of front instances

	array_Links = new Array<Link>() // array of front instances
	map_ID_Link = new Map<number, Link>() // map of front instances


	// getFrontArray allows for a get function that is robust to refactoring of the named struct name
	// for instance frontRepo.getArray<Astruct>( Astruct.GONGSTRUCT_NAME), is robust to a refactoring of Astruct identifier
	// contrary to frontRepo.Astructs_array which is not refactored when Astruct identifier is modified
	getFrontArray<Type>(gongStructName: string): Array<Type> {
		switch (gongStructName) {
			// insertion point
			case 'BookDetailsGroup':
				return this.array_BookDetailsGroups as unknown as Array<Type>
			case 'BookType':
				return this.array_BookTypes as unknown as Array<Type>
			case 'Books':
				return this.array_Bookss as unknown as Array<Type>
			case 'Credit':
				return this.array_Credits as unknown as Array<Type>
			case 'Link':
				return this.array_Links as unknown as Array<Type>
			default:
				throw new Error("Type not recognized");
		}
	}

	getFrontMap<Type>(gongStructName: string): Map<number, Type> {
		switch (gongStructName) {
			// insertion point
			case 'BookDetailsGroup':
				return this.map_ID_BookDetailsGroup as unknown as Map<number, Type>
			case 'BookType':
				return this.map_ID_BookType as unknown as Map<number, Type>
			case 'Books':
				return this.map_ID_Books as unknown as Map<number, Type>
			case 'Credit':
				return this.map_ID_Credit as unknown as Map<number, Type>
			case 'Link':
				return this.map_ID_Link as unknown as Map<number, Type>
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
		private bookdetailsgroupService: BookDetailsGroupService,
		private booktypeService: BookTypeService,
		private booksService: BooksService,
		private creditService: CreditService,
		private linkService: LinkService,
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
		Observable<BookDetailsGroupAPI[]>,
		Observable<BookTypeAPI[]>,
		Observable<BooksAPI[]>,
		Observable<CreditAPI[]>,
		Observable<LinkAPI[]>,
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
			this.bookdetailsgroupService.getBookDetailsGroups(this.GONG__StackPath, this.frontRepo),
			this.booktypeService.getBookTypes(this.GONG__StackPath, this.frontRepo),
			this.booksService.getBookss(this.GONG__StackPath, this.frontRepo),
			this.creditService.getCredits(this.GONG__StackPath, this.frontRepo),
			this.linkService.getLinks(this.GONG__StackPath, this.frontRepo),
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
			this.bookdetailsgroupService.getBookDetailsGroups(this.GONG__StackPath, this.frontRepo),
			this.booktypeService.getBookTypes(this.GONG__StackPath, this.frontRepo),
			this.booksService.getBookss(this.GONG__StackPath, this.frontRepo),
			this.creditService.getCredits(this.GONG__StackPath, this.frontRepo),
			this.linkService.getLinks(this.GONG__StackPath, this.frontRepo),
		]

		return new Observable<FrontRepo>(
			(observer) => {
				combineLatest(
					this.observableFrontRepo
				).subscribe(
					([
						___of_null, // see above for the explanation about of
						// insertion point sub template for declarations 
						bookdetailsgroups_,
						booktypes_,
						bookss_,
						credits_,
						links_,
					]) => {
						let _this = this
						// Typing can be messy with many items. Therefore, type casting is necessary here
						// insertion point sub template for type casting 
						var bookdetailsgroups: BookDetailsGroupAPI[]
						bookdetailsgroups = bookdetailsgroups_ as BookDetailsGroupAPI[]
						var booktypes: BookTypeAPI[]
						booktypes = booktypes_ as BookTypeAPI[]
						var bookss: BooksAPI[]
						bookss = bookss_ as BooksAPI[]
						var credits: CreditAPI[]
						credits = credits_ as CreditAPI[]
						var links: LinkAPI[]
						links = links_ as LinkAPI[]

						// 
						// First Step: init map of instances
						// insertion point sub template for init 
						// init the arrays
						this.frontRepo.array_BookDetailsGroups = []
						this.frontRepo.map_ID_BookDetailsGroup.clear()

						bookdetailsgroups.forEach(
							bookdetailsgroupAPI => {
								let bookdetailsgroup = new BookDetailsGroup
								this.frontRepo.array_BookDetailsGroups.push(bookdetailsgroup)
								this.frontRepo.map_ID_BookDetailsGroup.set(bookdetailsgroupAPI.ID, bookdetailsgroup)
							}
						)

						// init the arrays
						this.frontRepo.array_BookTypes = []
						this.frontRepo.map_ID_BookType.clear()

						booktypes.forEach(
							booktypeAPI => {
								let booktype = new BookType
								this.frontRepo.array_BookTypes.push(booktype)
								this.frontRepo.map_ID_BookType.set(booktypeAPI.ID, booktype)
							}
						)

						// init the arrays
						this.frontRepo.array_Bookss = []
						this.frontRepo.map_ID_Books.clear()

						bookss.forEach(
							booksAPI => {
								let books = new Books
								this.frontRepo.array_Bookss.push(books)
								this.frontRepo.map_ID_Books.set(booksAPI.ID, books)
							}
						)

						// init the arrays
						this.frontRepo.array_Credits = []
						this.frontRepo.map_ID_Credit.clear()

						credits.forEach(
							creditAPI => {
								let credit = new Credit
								this.frontRepo.array_Credits.push(credit)
								this.frontRepo.map_ID_Credit.set(creditAPI.ID, credit)
							}
						)

						// init the arrays
						this.frontRepo.array_Links = []
						this.frontRepo.map_ID_Link.clear()

						links.forEach(
							linkAPI => {
								let link = new Link
								this.frontRepo.array_Links.push(link)
								this.frontRepo.map_ID_Link.set(linkAPI.ID, link)
							}
						)


						// 
						// Second Step: reddeem front objects
						// insertion point sub template for redeem 
						// fill up front objects
						bookdetailsgroups.forEach(
							bookdetailsgroupAPI => {
								let bookdetailsgroup = this.frontRepo.map_ID_BookDetailsGroup.get(bookdetailsgroupAPI.ID)
								CopyBookDetailsGroupAPIToBookDetailsGroup(bookdetailsgroupAPI, bookdetailsgroup!, this.frontRepo)
							}
						)

						// fill up front objects
						booktypes.forEach(
							booktypeAPI => {
								let booktype = this.frontRepo.map_ID_BookType.get(booktypeAPI.ID)
								CopyBookTypeAPIToBookType(booktypeAPI, booktype!, this.frontRepo)
							}
						)

						// fill up front objects
						bookss.forEach(
							booksAPI => {
								let books = this.frontRepo.map_ID_Books.get(booksAPI.ID)
								CopyBooksAPIToBooks(booksAPI, books!, this.frontRepo)
							}
						)

						// fill up front objects
						credits.forEach(
							creditAPI => {
								let credit = this.frontRepo.map_ID_Credit.get(creditAPI.ID)
								CopyCreditAPIToCredit(creditAPI, credit!, this.frontRepo)
							}
						)

						// fill up front objects
						links.forEach(
							linkAPI => {
								let link = this.frontRepo.map_ID_Link.get(linkAPI.ID)
								CopyLinkAPIToLink(linkAPI, link!, this.frontRepo)
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
		let basePath = 'ws://localhost:8080/api/github.com/fullstack-lang/gongxsd/test/books/go/v1/ws/stage'
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
				this.frontRepo.array_BookDetailsGroups = []
				this.frontRepo.map_ID_BookDetailsGroup.clear()

				backRepoData.BookDetailsGroupAPIs.forEach(
					bookdetailsgroupAPI => {
						let bookdetailsgroup = new BookDetailsGroup
						this.frontRepo.array_BookDetailsGroups.push(bookdetailsgroup)
						this.frontRepo.map_ID_BookDetailsGroup.set(bookdetailsgroupAPI.ID, bookdetailsgroup)
					}
				)

				// init the arrays
				this.frontRepo.array_BookTypes = []
				this.frontRepo.map_ID_BookType.clear()

				backRepoData.BookTypeAPIs.forEach(
					booktypeAPI => {
						let booktype = new BookType
						this.frontRepo.array_BookTypes.push(booktype)
						this.frontRepo.map_ID_BookType.set(booktypeAPI.ID, booktype)
					}
				)

				// init the arrays
				this.frontRepo.array_Bookss = []
				this.frontRepo.map_ID_Books.clear()

				backRepoData.BooksAPIs.forEach(
					booksAPI => {
						let books = new Books
						this.frontRepo.array_Bookss.push(books)
						this.frontRepo.map_ID_Books.set(booksAPI.ID, books)
					}
				)

				// init the arrays
				this.frontRepo.array_Credits = []
				this.frontRepo.map_ID_Credit.clear()

				backRepoData.CreditAPIs.forEach(
					creditAPI => {
						let credit = new Credit
						this.frontRepo.array_Credits.push(credit)
						this.frontRepo.map_ID_Credit.set(creditAPI.ID, credit)
					}
				)

				// init the arrays
				this.frontRepo.array_Links = []
				this.frontRepo.map_ID_Link.clear()

				backRepoData.LinkAPIs.forEach(
					linkAPI => {
						let link = new Link
						this.frontRepo.array_Links.push(link)
						this.frontRepo.map_ID_Link.set(linkAPI.ID, link)
					}
				)


				// 
				// Second Step: reddeem front objects
				// insertion point sub template for redeem 
				// fill up front objects
				// insertion point sub template for redeem 
				// fill up front objects
				backRepoData.BookDetailsGroupAPIs.forEach(
					bookdetailsgroupAPI => {
						let bookdetailsgroup = this.frontRepo.map_ID_BookDetailsGroup.get(bookdetailsgroupAPI.ID)
						CopyBookDetailsGroupAPIToBookDetailsGroup(bookdetailsgroupAPI, bookdetailsgroup!, this.frontRepo)
					}
				)

				// fill up front objects
				backRepoData.BookTypeAPIs.forEach(
					booktypeAPI => {
						let booktype = this.frontRepo.map_ID_BookType.get(booktypeAPI.ID)
						CopyBookTypeAPIToBookType(booktypeAPI, booktype!, this.frontRepo)
					}
				)

				// fill up front objects
				backRepoData.BooksAPIs.forEach(
					booksAPI => {
						let books = this.frontRepo.map_ID_Books.get(booksAPI.ID)
						CopyBooksAPIToBooks(booksAPI, books!, this.frontRepo)
					}
				)

				// fill up front objects
				backRepoData.CreditAPIs.forEach(
					creditAPI => {
						let credit = this.frontRepo.map_ID_Credit.get(creditAPI.ID)
						CopyCreditAPIToCredit(creditAPI, credit!, this.frontRepo)
					}
				)

				// fill up front objects
				backRepoData.LinkAPIs.forEach(
					linkAPI => {
						let link = this.frontRepo.map_ID_Link.get(linkAPI.ID)
						CopyLinkAPIToLink(linkAPI, link!, this.frontRepo)
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
export function getBookDetailsGroupUniqueID(id: number): number {
	return 31 * id
}
export function getBookTypeUniqueID(id: number): number {
	return 37 * id
}
export function getBooksUniqueID(id: number): number {
	return 41 * id
}
export function getCreditUniqueID(id: number): number {
	return 43 * id
}
export function getLinkUniqueID(id: number): number {
	return 47 * id
}
