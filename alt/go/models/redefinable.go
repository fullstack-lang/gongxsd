package models

/*
<xs:group name="redefinable">

	<xs:annotation>
	    <xs:documentation> This group is for the elements which can self-redefine (see redefine
	        below). </xs:documentation>
	</xs:annotation>
	<xs:choice>
	    <xs:element ref="xs:simpleType" />
	    <xs:element ref="xs:complexType" />
	    <xs:element ref="xs:group" />
	    <xs:element ref="xs:attributeGroup" />
	</xs:choice>

</xs:group>
*/
type Redefinable struct {
	ComplexType *ComplexType `xml:"xs:complexType"`
}
