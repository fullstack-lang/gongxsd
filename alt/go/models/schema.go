package models

/*
<xs:element name="schema" id="schema">
<xs:annotation>

	<xs:documentation source="../structures/structures.html#element-schema" />

</xs:annotation>
<xs:complexType>

	<xs:complexContent>
		<xs:extension base="xs:openAttrs">
			<xs:sequence>
				<xs:group ref="xs:composition" minOccurs="0" maxOccurs="unbounded" />
				<xs:sequence minOccurs="0">
					<xs:element ref="xs:defaultOpenContent" />
					<xs:element ref="xs:annotation" minOccurs="0" maxOccurs="unbounded" />
				</xs:sequence>
				<xs:sequence minOccurs="0" maxOccurs="unbounded">
					<xs:group ref="xs:schemaTop" />
					<xs:element ref="xs:annotation" minOccurs="0" maxOccurs="unbounded" />
				</xs:sequence>
			</xs:sequence>
			<xs:attribute name="targetNamespace" type="xs:anyURI" />
			<xs:attribute name="version" type="xs:token" />
			<xs:attribute name="finalDefault" type="xs:fullDerivationSet" default=""
				use="optional" />
			<xs:attribute name="blockDefault" type="xs:blockSet" default="" use="optional" />
			<xs:attribute name="attributeFormDefault" type="xs:formChoice"
				default="unqualified" use="optional" />
			<xs:attribute name="elementFormDefault" type="xs:formChoice"
				default="unqualified" use="optional" />
			<xs:attribute name="defaultAttributes" type="xs:QName" />
			<xs:attribute name="xpathDefaultNamespace" type="xs:xpathDefaultNamespace"
				default="##local" use="optional" />
			<xs:attribute name="id" type="xs:ID" />
			<xs:attribute ref="xml:lang" />
		</xs:extension>
	</xs:complexContent>

</xs:complexType>
..
</xs:element>
*/
type Schema struct {
	Name string
	Xs   string `xml:"xs,attr"`
	Annotated
	Schema_A_ComplexType
}

type Schema_A_ComplexType struct {
	Schema_A_ComplexType_A_ComplexContentDummy int
	Schema_A_ComplexType_A_ComplexContent
}

type Schema_A_ComplexType_A_ComplexContent struct {
	Schema_A_ComplexType_A_ComplexContent_A_Extension
}

type Schema_A_ComplexType_A_ComplexContent_A_Extension struct {
	Schema_A_ComplexType_A_ComplexContent_A_Extension_SequenceDummy int
	Schema_A_ComplexType_A_ComplexContent_A_Extension_Sequence
}

type Schema_A_ComplexType_A_ComplexContent_A_Extension_Sequence struct {
	Schema_A_ComplexType_A_ComplexContent_A_Extension_Sequence_Sequence1

	// this line below can read the complex type
	// Redefinables []*ComplexType `xml:"complexType"`

	Sequence2 struct {
		ComplexType *ComplexType `xml:"complexType"`
	} `xml:",inline"`
}

type Foo struct {
}

/*
<xs:sequence minOccurs="0">

	<xs:element ref="xs:defaultOpenContent" />
	<xs:element ref="xs:annotation" minOccurs="0" maxOccurs="unbounded" />

</xs:sequence>
*/

type Schema_A_ComplexType_A_ComplexContent_A_Extension_Sequence_Sequence1 struct {
	Schema_A_ComplexType_A_ComplexContent_A_Extension_Sequence_Sequence1Dummy int
	// shall i have a pointer or just a composition to reflect
	// that minOccurs
	// A pointer would help to have a Annotation at each level
	// composition is easier to understand
	DefaultOpenContent
}

/*
	<xs:sequence minOccurs="0" maxOccurs="unbounded">
		<xs:group ref="xs:schemaTop" />
		<xs:element ref="xs:annotation" minOccurs="0" maxOccurs="unbounded" />
	</xs:sequence>
*/

type Schema_A_ComplexType_A_ComplexContent_A_Extension_Sequence_Sequence2 struct {
	SchemaTop
}

func (schema *Schema) Setup() {

	schema.Name = schema.Xs

	for _, d := range schema.Annotation.Documentations {
		d.Name = "Documentation  of " + schema.Name
	}

}
