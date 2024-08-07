package models

type Enumeration struct {
	Name string
	ElementWithAnnotation
	ElementWithValueAttribute
}

type MinInclusive struct {
	Name string
	ElementWithAnnotation
	ElementWithValueAttribute
}

type MaxInclusive struct {
	Name string
	ElementWithAnnotation
	ElementWithValueAttribute
}

type Pattern struct {
	Name string
	ElementWithAnnotation
	ElementWithValueAttribute
}

type WhiteSpace struct {
	Name string
	ElementWithAnnotation
	ElementWithValueAttribute
}

type MinLength struct {
	Name string
	ElementWithAnnotation
	ElementWithValueAttribute
}

type MaxLength struct {
	Name string
	ElementWithAnnotation
	ElementWithValueAttribute
}

type Length struct {
	Name string
	ElementWithAnnotation
	ElementWithValueAttribute
}
