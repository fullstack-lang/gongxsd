// generated code - do not edit

import { LinkAPI } from './link-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Link {

	static GONGSTRUCT_NAME = "Link"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	EnclosedText: string = ""
	NameXSD: string = ""

	// insertion point for pointers and slices of pointers declarations
}

export function CopyLinkToLinkAPI(link: Link, linkAPI: LinkAPI) {

	linkAPI.CreatedAt = link.CreatedAt
	linkAPI.DeletedAt = link.DeletedAt
	linkAPI.ID = link.ID

	// insertion point for basic fields copy operations
	linkAPI.Name = link.Name
	linkAPI.EnclosedText = link.EnclosedText
	linkAPI.NameXSD = link.NameXSD

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyLinkAPIToLink update basic, pointers and slice of pointers fields of link
// from respectively the basic fields and encoded fields of pointers and slices of pointers of linkAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyLinkAPIToLink(linkAPI: LinkAPI, link: Link, frontRepo: FrontRepo) {

	link.CreatedAt = linkAPI.CreatedAt
	link.DeletedAt = linkAPI.DeletedAt
	link.ID = linkAPI.ID

	// insertion point for basic fields copy operations
	link.Name = linkAPI.Name
	link.EnclosedText = linkAPI.EnclosedText
	link.NameXSD = linkAPI.NameXSD

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
