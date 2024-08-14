// generated code - do not edit

//insertion point for imports
import { BookDetailsGroupAPI } from './bookdetailsgroup-api'

import { BookTypeAPI } from './booktype-api'

import { BooksAPI } from './books-api'

import { CreditAPI } from './credit-api'

import { LinkAPI } from './link-api'


export class BackRepoData {
	// insertion point for declarations
	BookDetailsGroupAPIs = new Array<BookDetailsGroupAPI>()

	BookTypeAPIs = new Array<BookTypeAPI>()

	BooksAPIs = new Array<BooksAPI>()

	CreditAPIs = new Array<CreditAPI>()

	LinkAPIs = new Array<LinkAPI>()



	constructor(data?: Partial<BackRepoData>) {
		// insertion point for copies
		this.BookDetailsGroupAPIs = data?.BookDetailsGroupAPIs || [];

		this.BookTypeAPIs = data?.BookTypeAPIs || [];

		this.BooksAPIs = data?.BooksAPIs || [];

		this.CreditAPIs = data?.CreditAPIs || [];

		this.LinkAPIs = data?.LinkAPIs || [];

	}

}