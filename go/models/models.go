package models

type ElementWithNameAttribute struct {
	NameXSD string `xml:"name,attr"`
}

type ElementWithTypeAttribute struct {
	Type string `xml:"type,attr"`
}

type ElementWithValueAttribute struct {
	Value string `xml:"value,attr"`
}

type Annotated struct {
	Annotation *Annotation `xml:"annotation"`
}

type Annotation struct {
	Name           string
	Documentations []*Documentation `xml:"documentation"`
}

type Union struct {
	Name string
	Annotated
	MemberTypes string `xml:"memberTypes,attr"`
}

type Attribute struct {
	Name string
	ElementWithNameAttribute
	ElementWithTypeAttribute
	Annotated

	WithGoIdentifier

	Default         string `xml:"default,attr"`
	Use             string `xml:"use,attr"`
	Form            string `xml:"form,attr"`
	Fixed           string `xml:"fixed,attr"`
	Ref             string `xml:"ref,attr"`
	TargetNamespace string `xml:"targetNamespace,attr"`
	SimpleType      string `xml:"simpleType,attr"`
	IDXSD           string `xml:"id,attr"`
}
