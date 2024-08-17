package models

/*
<xs:group name="schemaTop">

	<xs:annotation>
	    <xs:documentation> This group is for the elements which occur freely at the top level of
	        schemas. All of their types are based on the "annotated" type by extension. </xs:documentation>
	</xs:annotation>
	<xs:choice>
	    <xs:group ref="xs:redefinable" />
	    <xs:element ref="xs:element" />
	    <xs:element ref="xs:attribute" />
	    <xs:element ref="xs:notation" />
	</xs:choice>

</xs:group>
*/
type SchemaTop struct {
	SchemaTop_Choice
}

type SchemaTop_Choice struct {
	Redefinable
}
