package models

type ComplexType struct {
	Name string

	ComplexType_A_ComplexContent

	//
	//
	//

	OuterSchema *Schema
}

type ComplexType_A_ComplexContent struct {
	Annotated
}

func (ct *ComplexType) Setup() {

	if ct.OuterSchema != nil {
		ct.Name = "CT_A_" + ct.OuterSchema.Name
	}

}
