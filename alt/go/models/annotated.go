package models

type Annotated struct {
	Annotation *Annotation `xml:"annotation"`
}

type Annotation struct {
	Name           string
	Documentations []*Documentation `xml:"documentation"`
}
