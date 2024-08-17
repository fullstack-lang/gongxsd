package models

/*
<xs:element name="defaultOpenContent" id="defaultOpenContent">

	<xs:annotation>
	    <xs:documentation source="../structures/structures.html#element-defaultOpenContent" />
	</xs:annotation>
	<xs:complexType>
	    <xs:complexContent>
	        <xs:extension base="xs:annotated">
	            <xs:sequence>
	                <xs:element name="any" type="xs:wildcard" />
	            </xs:sequence>

*/

// this is for any attributes
type DefaultOpenContent struct {
	DefaultOpenContent_A_ComplexType
}

type DefaultOpenContent_A_ComplexType struct {
}

/*
<xs:complexType name="wildcard">

	<xs:complexContent>
	    <xs:extension base="xs:annotated">
	        <xs:attributeGroup ref="xs:anyAttrGroup" />
	    </xs:extension>
	</xs:complexContent>

</xs:complexType>
*/
type DefaultOpenContent_A_ComplexType_A_Sequence struct {
}
