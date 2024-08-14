// generated code - do not edit

import { BooksAPI } from './books-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { BookType } from './booktype'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Books {

	static GONGSTRUCT_NAME = "Books"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for pointers and slices of pointers declarations
	Book: Array<BookType> = []
}

export function CopyBooksToBooksAPI(books: Books, booksAPI: BooksAPI) {

	booksAPI.CreatedAt = books.CreatedAt
	booksAPI.DeletedAt = books.DeletedAt
	booksAPI.ID = books.ID

	// insertion point for basic fields copy operations
	booksAPI.Name = books.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	booksAPI.BooksPointersEncoding.Book = []
	for (let _booktype of books.Book) {
		booksAPI.BooksPointersEncoding.Book.push(_booktype.ID)
	}

}

// CopyBooksAPIToBooks update basic, pointers and slice of pointers fields of books
// from respectively the basic fields and encoded fields of pointers and slices of pointers of booksAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyBooksAPIToBooks(booksAPI: BooksAPI, books: Books, frontRepo: FrontRepo) {

	books.CreatedAt = booksAPI.CreatedAt
	books.DeletedAt = booksAPI.DeletedAt
	books.ID = booksAPI.ID

	// insertion point for basic fields copy operations
	books.Name = booksAPI.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	books.Book = new Array<BookType>()
	for (let _id of booksAPI.BooksPointersEncoding.Book) {
		let _booktype = frontRepo.map_ID_BookType.get(_id)
		if (_booktype != undefined) {
			books.Book.push(_booktype!)
		}
	}
}
