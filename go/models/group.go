package models

type Group struct {
	Name string
	ElementWithAnnotation
	ElementWithNameAttribute
	Ref string `xml:"ref,attr"`

	Composer
}
