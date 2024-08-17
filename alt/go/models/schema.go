package models

type Schema struct {
	Name string
	Xs   string `xml:"xs,attr"`
	Annotated
}
