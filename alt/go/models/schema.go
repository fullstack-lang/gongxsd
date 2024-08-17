package models

type Schema struct {
	Name string
	Xs   string `xml:"xs,attr"`
	Annotated
	ComplexType *ComplexType `xml:"complexType"`
}

func (schema *Schema) Setup() {

	schema.Name = schema.Xs
	schema.ComplexType.OuterSchema = schema
	schema.ComplexType.Setup()

	for _, d := range schema.Annotation.Documentations {
		d.Name = "Documentation  of " + schema.Name
	}

}
