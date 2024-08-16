package models

type Schema struct {
	Name string
	Xs   string `xml:"xs,attr"`
	Annotated
	Elements        []*Element        `xml:"element"`
	SimpleTypes     []*SimpleType     `xml:"simpleType"`
	ComplexTypes    []*ComplexType    `xml:"complexType"`
	AttributeGroups []*AttributeGroup `xml:"attributeGroup"`
	Groups          []*Group          `xml:"group"`
}
