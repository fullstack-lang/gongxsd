package models

type SimpleType struct {
	Name string
	Annotated
	ElementWithNameAttribute
	Restriction *Restriction `xml:"restriction"`
	Union       *Union       `xml:"union"`
}
