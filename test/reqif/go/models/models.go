// generated code - do not edit
package models

import "encoding/xml"

// to avoid compilation error if no xml element is generated
var _ xml.Attr

// ALTERNATIVE_ID Named source named complex type "ALTERNATIVE-ID"
type ALTERNATIVE_ID struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "IDENTIFIER
	IDENTIFIER string `xml:"IDENTIFIER,attr,omitempty"`
}

// ATTRIBUTE_DEFINITION_BOOLEAN Named source named complex type "ATTRIBUTE-DEFINITION-BOOLEAN"
type ATTRIBUTE_DEFINITION_BOOLEAN struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "DESC
	DESC string `xml:"DESC,attr,omitempty"`

	// generated from attribute "IDENTIFIER
	IDENTIFIER string `xml:"IDENTIFIER,attr,omitempty"`

	// generated from attribute "IS-EDITABLE
	IS_EDITABLE bool `xml:"IS-EDITABLE,attr,omitempty"`

	// generated from attribute "LAST-CHANGE
	LAST_CHANGE string `xml:"LAST-CHANGE,attr,omitempty"`

	// generated from attribute "LONG-NAME
	LONG_NAME string `xml:"LONG-NAME,attr,omitempty"`

	// generated from anonymous type within outer element "ALTERNATIVE-ID" of type A_ALTERNATIVE-ID.
	ALTERNATIVE_ID []*A_ALTERNATIVE_ID `xml:"ALTERNATIVE-ID,omitempty"`

	// generated from anonymous type within outer element "DEFAULT-VALUE" of type A_ATTRIBUTE-VALUE-BOOLEAN.
	DEFAULT_VALUE []*A_ATTRIBUTE_VALUE_BOOLEAN `xml:"DEFAULT-VALUE,omitempty"`

	// generated from anonymous type within outer element "TYPE" of type A_DATATYPE-DEFINITION-BOOLEAN-REF.
	TYPE []*A_DATATYPE_DEFINITION_BOOLEAN_REF `xml:"TYPE,omitempty"`
}

// ATTRIBUTE_DEFINITION_DATE Named source named complex type "ATTRIBUTE-DEFINITION-DATE"
type ATTRIBUTE_DEFINITION_DATE struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "DESC
	DESC string `xml:"DESC,attr,omitempty"`

	// generated from attribute "IDENTIFIER
	IDENTIFIER string `xml:"IDENTIFIER,attr,omitempty"`

	// generated from attribute "IS-EDITABLE
	IS_EDITABLE bool `xml:"IS-EDITABLE,attr,omitempty"`

	// generated from attribute "LAST-CHANGE
	LAST_CHANGE string `xml:"LAST-CHANGE,attr,omitempty"`

	// generated from attribute "LONG-NAME
	LONG_NAME string `xml:"LONG-NAME,attr,omitempty"`

	// generated from anonymous type within outer element "ALTERNATIVE-ID" of type A_ALTERNATIVE-ID.
	ALTERNATIVE_ID []*A_ALTERNATIVE_ID `xml:"ALTERNATIVE-ID,omitempty"`

	// generated from anonymous type within outer element "DEFAULT-VALUE" of type A_ATTRIBUTE-VALUE-DATE.
	DEFAULT_VALUE []*A_ATTRIBUTE_VALUE_DATE `xml:"DEFAULT-VALUE,omitempty"`

	// generated from anonymous type within outer element "TYPE" of type A_DATATYPE-DEFINITION-DATE-REF.
	TYPE []*A_DATATYPE_DEFINITION_DATE_REF `xml:"TYPE,omitempty"`
}

// ATTRIBUTE_DEFINITION_ENUMERATION Named source named complex type "ATTRIBUTE-DEFINITION-ENUMERATION"
type ATTRIBUTE_DEFINITION_ENUMERATION struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "DESC
	DESC string `xml:"DESC,attr,omitempty"`

	// generated from attribute "IDENTIFIER
	IDENTIFIER string `xml:"IDENTIFIER,attr,omitempty"`

	// generated from attribute "IS-EDITABLE
	IS_EDITABLE bool `xml:"IS-EDITABLE,attr,omitempty"`

	// generated from attribute "LAST-CHANGE
	LAST_CHANGE string `xml:"LAST-CHANGE,attr,omitempty"`

	// generated from attribute "LONG-NAME
	LONG_NAME string `xml:"LONG-NAME,attr,omitempty"`

	// generated from attribute "MULTI-VALUED
	MULTI_VALUED bool `xml:"MULTI-VALUED,attr,omitempty"`

	// generated from anonymous type within outer element "ALTERNATIVE-ID" of type A_ALTERNATIVE-ID.
	ALTERNATIVE_ID []*A_ALTERNATIVE_ID `xml:"ALTERNATIVE-ID,omitempty"`

	// generated from anonymous type within outer element "DEFAULT-VALUE" of type A_ATTRIBUTE-VALUE-ENUMERATION.
	DEFAULT_VALUE []*A_ATTRIBUTE_VALUE_ENUMERATION `xml:"DEFAULT-VALUE,omitempty"`

	// generated from anonymous type within outer element "TYPE" of type A_DATATYPE-DEFINITION-ENUMERATION-REF.
	TYPE []*A_DATATYPE_DEFINITION_ENUMERATION_REF `xml:"TYPE,omitempty"`
}

// ATTRIBUTE_DEFINITION_INTEGER Named source named complex type "ATTRIBUTE-DEFINITION-INTEGER"
type ATTRIBUTE_DEFINITION_INTEGER struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "DESC
	DESC string `xml:"DESC,attr,omitempty"`

	// generated from attribute "IDENTIFIER
	IDENTIFIER string `xml:"IDENTIFIER,attr,omitempty"`

	// generated from attribute "IS-EDITABLE
	IS_EDITABLE bool `xml:"IS-EDITABLE,attr,omitempty"`

	// generated from attribute "LAST-CHANGE
	LAST_CHANGE string `xml:"LAST-CHANGE,attr,omitempty"`

	// generated from attribute "LONG-NAME
	LONG_NAME string `xml:"LONG-NAME,attr,omitempty"`

	// generated from anonymous type within outer element "ALTERNATIVE-ID" of type A_ALTERNATIVE-ID.
	ALTERNATIVE_ID []*A_ALTERNATIVE_ID `xml:"ALTERNATIVE-ID,omitempty"`

	// generated from anonymous type within outer element "DEFAULT-VALUE" of type A_ATTRIBUTE-VALUE-INTEGER.
	DEFAULT_VALUE []*A_ATTRIBUTE_VALUE_INTEGER `xml:"DEFAULT-VALUE,omitempty"`

	// generated from anonymous type within outer element "TYPE" of type A_DATATYPE-DEFINITION-INTEGER-REF.
	TYPE []*A_DATATYPE_DEFINITION_INTEGER_REF `xml:"TYPE,omitempty"`
}

// ATTRIBUTE_DEFINITION_REAL Named source named complex type "ATTRIBUTE-DEFINITION-REAL"
type ATTRIBUTE_DEFINITION_REAL struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "DESC
	DESC string `xml:"DESC,attr,omitempty"`

	// generated from attribute "IDENTIFIER
	IDENTIFIER string `xml:"IDENTIFIER,attr,omitempty"`

	// generated from attribute "IS-EDITABLE
	IS_EDITABLE bool `xml:"IS-EDITABLE,attr,omitempty"`

	// generated from attribute "LAST-CHANGE
	LAST_CHANGE string `xml:"LAST-CHANGE,attr,omitempty"`

	// generated from attribute "LONG-NAME
	LONG_NAME string `xml:"LONG-NAME,attr,omitempty"`

	// generated from anonymous type within outer element "ALTERNATIVE-ID" of type A_ALTERNATIVE-ID.
	ALTERNATIVE_ID []*A_ALTERNATIVE_ID `xml:"ALTERNATIVE-ID,omitempty"`

	// generated from anonymous type within outer element "DEFAULT-VALUE" of type A_ATTRIBUTE-VALUE-REAL.
	DEFAULT_VALUE []*A_ATTRIBUTE_VALUE_REAL `xml:"DEFAULT-VALUE,omitempty"`

	// generated from anonymous type within outer element "TYPE" of type A_DATATYPE-DEFINITION-REAL-REF.
	TYPE []*A_DATATYPE_DEFINITION_REAL_REF `xml:"TYPE,omitempty"`
}

// ATTRIBUTE_DEFINITION_STRING Named source named complex type "ATTRIBUTE-DEFINITION-STRING"
type ATTRIBUTE_DEFINITION_STRING struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "DESC
	DESC string `xml:"DESC,attr,omitempty"`

	// generated from attribute "IDENTIFIER
	IDENTIFIER string `xml:"IDENTIFIER,attr,omitempty"`

	// generated from attribute "IS-EDITABLE
	IS_EDITABLE bool `xml:"IS-EDITABLE,attr,omitempty"`

	// generated from attribute "LAST-CHANGE
	LAST_CHANGE string `xml:"LAST-CHANGE,attr,omitempty"`

	// generated from attribute "LONG-NAME
	LONG_NAME string `xml:"LONG-NAME,attr,omitempty"`

	// generated from anonymous type within outer element "ALTERNATIVE-ID" of type A_ALTERNATIVE-ID.
	ALTERNATIVE_ID []*A_ALTERNATIVE_ID `xml:"ALTERNATIVE-ID,omitempty"`

	// generated from anonymous type within outer element "DEFAULT-VALUE" of type A_ATTRIBUTE-VALUE-STRING.
	DEFAULT_VALUE []*A_ATTRIBUTE_VALUE_STRING `xml:"DEFAULT-VALUE,omitempty"`

	// generated from anonymous type within outer element "TYPE" of type A_DATATYPE-DEFINITION-STRING-REF.
	TYPE []*A_DATATYPE_DEFINITION_STRING_REF `xml:"TYPE,omitempty"`
}

// ATTRIBUTE_DEFINITION_XHTML Named source named complex type "ATTRIBUTE-DEFINITION-XHTML"
type ATTRIBUTE_DEFINITION_XHTML struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "DESC
	DESC string `xml:"DESC,attr,omitempty"`

	// generated from attribute "IDENTIFIER
	IDENTIFIER string `xml:"IDENTIFIER,attr,omitempty"`

	// generated from attribute "IS-EDITABLE
	IS_EDITABLE bool `xml:"IS-EDITABLE,attr,omitempty"`

	// generated from attribute "LAST-CHANGE
	LAST_CHANGE string `xml:"LAST-CHANGE,attr,omitempty"`

	// generated from attribute "LONG-NAME
	LONG_NAME string `xml:"LONG-NAME,attr,omitempty"`

	// generated from anonymous type within outer element "ALTERNATIVE-ID" of type A_ALTERNATIVE-ID.
	ALTERNATIVE_ID []*A_ALTERNATIVE_ID `xml:"ALTERNATIVE-ID,omitempty"`

	// generated from anonymous type within outer element "DEFAULT-VALUE" of type A_ATTRIBUTE-VALUE-XHTML.
	DEFAULT_VALUE []*A_ATTRIBUTE_VALUE_XHTML `xml:"DEFAULT-VALUE,omitempty"`

	// generated from anonymous type within outer element "TYPE" of type A_DATATYPE-DEFINITION-XHTML-REF.
	TYPE []*A_DATATYPE_DEFINITION_XHTML_REF `xml:"TYPE,omitempty"`
}

// ATTRIBUTE_VALUE_BOOLEAN Named source named complex type "ATTRIBUTE-VALUE-BOOLEAN"
type ATTRIBUTE_VALUE_BOOLEAN struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "THE-VALUE
	THE_VALUE bool `xml:"THE-VALUE,attr,omitempty"`

	// generated from anonymous type within outer element "DEFINITION" of type A_ATTRIBUTE-DEFINITION-BOOLEAN-REF.
	DEFINITION []*A_ATTRIBUTE_DEFINITION_BOOLEAN_REF `xml:"DEFINITION,omitempty"`
}

// ATTRIBUTE_VALUE_DATE Named source named complex type "ATTRIBUTE-VALUE-DATE"
type ATTRIBUTE_VALUE_DATE struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "THE-VALUE
	THE_VALUE string `xml:"THE-VALUE,attr,omitempty"`

	// generated from anonymous type within outer element "DEFINITION" of type A_ATTRIBUTE-DEFINITION-DATE-REF.
	DEFINITION []*A_ATTRIBUTE_DEFINITION_DATE_REF `xml:"DEFINITION,omitempty"`
}

// ATTRIBUTE_VALUE_ENUMERATION Named source named complex type "ATTRIBUTE-VALUE-ENUMERATION"
type ATTRIBUTE_VALUE_ENUMERATION struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from anonymous type within outer element "DEFINITION" of type A_ATTRIBUTE-DEFINITION-ENUMERATION-REF.
	DEFINITION []*A_ATTRIBUTE_DEFINITION_ENUMERATION_REF `xml:"DEFINITION,omitempty"`

	// generated from anonymous type within outer element "VALUES" of type A_ENUM-VALUE-REF.
	VALUES []*A_ENUM_VALUE_REF `xml:"VALUES,omitempty"`
}

// ATTRIBUTE_VALUE_INTEGER Named source named complex type "ATTRIBUTE-VALUE-INTEGER"
type ATTRIBUTE_VALUE_INTEGER struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "THE-VALUE
	THE_VALUE int `xml:"THE-VALUE,attr,omitempty"`

	// generated from anonymous type within outer element "DEFINITION" of type A_ATTRIBUTE-DEFINITION-INTEGER-REF.
	DEFINITION []*A_ATTRIBUTE_DEFINITION_INTEGER_REF `xml:"DEFINITION,omitempty"`
}

// ATTRIBUTE_VALUE_REAL Named source named complex type "ATTRIBUTE-VALUE-REAL"
type ATTRIBUTE_VALUE_REAL struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "THE-VALUE
	THE_VALUE float64 `xml:"THE-VALUE,attr,omitempty"`

	// generated from anonymous type within outer element "DEFINITION" of type A_ATTRIBUTE-DEFINITION-REAL-REF.
	DEFINITION []*A_ATTRIBUTE_DEFINITION_REAL_REF `xml:"DEFINITION,omitempty"`
}

// ATTRIBUTE_VALUE_STRING Named source named complex type "ATTRIBUTE-VALUE-STRING"
type ATTRIBUTE_VALUE_STRING struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "THE-VALUE
	THE_VALUE string `xml:"THE-VALUE,attr,omitempty"`

	// generated from anonymous type within outer element "DEFINITION" of type A_ATTRIBUTE-DEFINITION-STRING-REF.
	DEFINITION []*A_ATTRIBUTE_DEFINITION_STRING_REF `xml:"DEFINITION,omitempty"`
}

// ATTRIBUTE_VALUE_XHTML Named source named complex type "ATTRIBUTE-VALUE-XHTML"
type ATTRIBUTE_VALUE_XHTML struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "IS-SIMPLIFIED
	IS_SIMPLIFIED bool `xml:"IS-SIMPLIFIED,attr,omitempty"`

	// generated from element "THE-VALUE" of type XHTML-CONTENT order 57 depth 0
	THE_VALUE []*XHTML_CONTENT `xml:"THE-VALUE,omitempty"`

	// generated from element "THE-ORIGINAL-VALUE" of type XHTML-CONTENT order 58 depth 0
	THE_ORIGINAL_VALUE []*XHTML_CONTENT `xml:"THE-ORIGINAL-VALUE,omitempty"`

	// generated from anonymous type within outer element "DEFINITION" of type A_ATTRIBUTE-DEFINITION-XHTML-REF.
	DEFINITION []*A_ATTRIBUTE_DEFINITION_XHTML_REF `xml:"DEFINITION,omitempty"`
}

// A_ALTERNATIVE_ID Named source within outer element "ALTERNATIVE-ID"
type A_ALTERNATIVE_ID struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "ALTERNATIVE-ID" of type ALTERNATIVE-ID order 2 depth 1
	ALTERNATIVE_ID []*ALTERNATIVE_ID `xml:"ALTERNATIVE-ID,omitempty"`
}

// A_ATTRIBUTE_DEFINITION_BOOLEAN_REF Named source within outer element "DEFINITION"
type A_ATTRIBUTE_DEFINITION_BOOLEAN_REF struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "ATTRIBUTE-DEFINITION-BOOLEAN-REF" of type LOCAL-REF order 44 depth 1
	ATTRIBUTE_DEFINITION_BOOLEAN_REF string `xml:"ATTRIBUTE-DEFINITION-BOOLEAN-REF,omitempty"`
}

// A_ATTRIBUTE_DEFINITION_DATE_REF Named source within outer element "DEFINITION"
type A_ATTRIBUTE_DEFINITION_DATE_REF struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "ATTRIBUTE-DEFINITION-DATE-REF" of type LOCAL-REF order 46 depth 1
	ATTRIBUTE_DEFINITION_DATE_REF string `xml:"ATTRIBUTE-DEFINITION-DATE-REF,omitempty"`
}

// A_ATTRIBUTE_DEFINITION_ENUMERATION_REF Named source within outer element "DEFINITION"
type A_ATTRIBUTE_DEFINITION_ENUMERATION_REF struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "ATTRIBUTE-DEFINITION-ENUMERATION-REF" of type LOCAL-REF order 48 depth 1
	ATTRIBUTE_DEFINITION_ENUMERATION_REF string `xml:"ATTRIBUTE-DEFINITION-ENUMERATION-REF,omitempty"`
}

// A_ATTRIBUTE_DEFINITION_INTEGER_REF Named source within outer element "DEFINITION"
type A_ATTRIBUTE_DEFINITION_INTEGER_REF struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "ATTRIBUTE-DEFINITION-INTEGER-REF" of type LOCAL-REF order 52 depth 1
	ATTRIBUTE_DEFINITION_INTEGER_REF string `xml:"ATTRIBUTE-DEFINITION-INTEGER-REF,omitempty"`
}

// A_ATTRIBUTE_DEFINITION_REAL_REF Named source within outer element "DEFINITION"
type A_ATTRIBUTE_DEFINITION_REAL_REF struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "ATTRIBUTE-DEFINITION-REAL-REF" of type LOCAL-REF order 54 depth 1
	ATTRIBUTE_DEFINITION_REAL_REF string `xml:"ATTRIBUTE-DEFINITION-REAL-REF,omitempty"`
}

// A_ATTRIBUTE_DEFINITION_STRING_REF Named source within outer element "DEFINITION"
type A_ATTRIBUTE_DEFINITION_STRING_REF struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "ATTRIBUTE-DEFINITION-STRING-REF" of type LOCAL-REF order 56 depth 1
	ATTRIBUTE_DEFINITION_STRING_REF string `xml:"ATTRIBUTE-DEFINITION-STRING-REF,omitempty"`
}

// A_ATTRIBUTE_DEFINITION_XHTML_REF Named source within outer element "DEFINITION"
type A_ATTRIBUTE_DEFINITION_XHTML_REF struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "ATTRIBUTE-DEFINITION-XHTML-REF" of type LOCAL-REF order 60 depth 1
	ATTRIBUTE_DEFINITION_XHTML_REF string `xml:"ATTRIBUTE-DEFINITION-XHTML-REF,omitempty"`
}

// A_ATTRIBUTE_VALUE_BOOLEAN Named source within outer element "DEFAULT-VALUE"
type A_ATTRIBUTE_VALUE_BOOLEAN struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "ATTRIBUTE-VALUE-BOOLEAN" of type ATTRIBUTE-VALUE-BOOLEAN order 4 depth 1
	ATTRIBUTE_VALUE_BOOLEAN []*ATTRIBUTE_VALUE_BOOLEAN `xml:"ATTRIBUTE-VALUE-BOOLEAN,omitempty"`
}

// A_ATTRIBUTE_VALUE_DATE Named source within outer element "DEFAULT-VALUE"
type A_ATTRIBUTE_VALUE_DATE struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "ATTRIBUTE-VALUE-DATE" of type ATTRIBUTE-VALUE-DATE order 10 depth 1
	ATTRIBUTE_VALUE_DATE []*ATTRIBUTE_VALUE_DATE `xml:"ATTRIBUTE-VALUE-DATE,omitempty"`
}

// A_ATTRIBUTE_VALUE_ENUMERATION Named source within outer element "DEFAULT-VALUE"
type A_ATTRIBUTE_VALUE_ENUMERATION struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "ATTRIBUTE-VALUE-ENUMERATION" of type ATTRIBUTE-VALUE-ENUMERATION order 14 depth 1
	ATTRIBUTE_VALUE_ENUMERATION []*ATTRIBUTE_VALUE_ENUMERATION `xml:"ATTRIBUTE-VALUE-ENUMERATION,omitempty"`
}

// A_ATTRIBUTE_VALUE_INTEGER Named source within outer element "DEFAULT-VALUE"
type A_ATTRIBUTE_VALUE_INTEGER struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "ATTRIBUTE-VALUE-INTEGER" of type ATTRIBUTE-VALUE-INTEGER order 22 depth 1
	ATTRIBUTE_VALUE_INTEGER []*ATTRIBUTE_VALUE_INTEGER `xml:"ATTRIBUTE-VALUE-INTEGER,omitempty"`
}

// A_ATTRIBUTE_VALUE_REAL Named source within outer element "DEFAULT-VALUE"
type A_ATTRIBUTE_VALUE_REAL struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "ATTRIBUTE-VALUE-REAL" of type ATTRIBUTE-VALUE-REAL order 28 depth 1
	ATTRIBUTE_VALUE_REAL []*ATTRIBUTE_VALUE_REAL `xml:"ATTRIBUTE-VALUE-REAL,omitempty"`
}

// A_ATTRIBUTE_VALUE_STRING Named source within outer element "DEFAULT-VALUE"
type A_ATTRIBUTE_VALUE_STRING struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "ATTRIBUTE-VALUE-STRING" of type ATTRIBUTE-VALUE-STRING order 34 depth 1
	ATTRIBUTE_VALUE_STRING []*ATTRIBUTE_VALUE_STRING `xml:"ATTRIBUTE-VALUE-STRING,omitempty"`
}

// A_ATTRIBUTE_VALUE_XHTML Named source within outer element "DEFAULT-VALUE"
type A_ATTRIBUTE_VALUE_XHTML struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "ATTRIBUTE-VALUE-XHTML" of type ATTRIBUTE-VALUE-XHTML order 40 depth 1
	ATTRIBUTE_VALUE_XHTML []*ATTRIBUTE_VALUE_XHTML `xml:"ATTRIBUTE-VALUE-XHTML,omitempty"`
}

// A_ATTRIBUTE_VALUE_XHTML_1 Named source within outer element "VALUES"
type A_ATTRIBUTE_VALUE_XHTML_1 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "ATTRIBUTE-VALUE-BOOLEAN" of type ATTRIBUTE-VALUE-BOOLEAN order 4 depth 1
	ATTRIBUTE_VALUE_BOOLEAN []*ATTRIBUTE_VALUE_BOOLEAN `xml:"ATTRIBUTE-VALUE-BOOLEAN,omitempty"`

	// generated from element "ATTRIBUTE-VALUE-DATE" of type ATTRIBUTE-VALUE-DATE order 10 depth 1
	ATTRIBUTE_VALUE_DATE []*ATTRIBUTE_VALUE_DATE `xml:"ATTRIBUTE-VALUE-DATE,omitempty"`

	// generated from element "ATTRIBUTE-VALUE-ENUMERATION" of type ATTRIBUTE-VALUE-ENUMERATION order 14 depth 1
	ATTRIBUTE_VALUE_ENUMERATION []*ATTRIBUTE_VALUE_ENUMERATION `xml:"ATTRIBUTE-VALUE-ENUMERATION,omitempty"`

	// generated from element "ATTRIBUTE-VALUE-INTEGER" of type ATTRIBUTE-VALUE-INTEGER order 22 depth 1
	ATTRIBUTE_VALUE_INTEGER []*ATTRIBUTE_VALUE_INTEGER `xml:"ATTRIBUTE-VALUE-INTEGER,omitempty"`

	// generated from element "ATTRIBUTE-VALUE-REAL" of type ATTRIBUTE-VALUE-REAL order 28 depth 1
	ATTRIBUTE_VALUE_REAL []*ATTRIBUTE_VALUE_REAL `xml:"ATTRIBUTE-VALUE-REAL,omitempty"`

	// generated from element "ATTRIBUTE-VALUE-STRING" of type ATTRIBUTE-VALUE-STRING order 34 depth 1
	ATTRIBUTE_VALUE_STRING []*ATTRIBUTE_VALUE_STRING `xml:"ATTRIBUTE-VALUE-STRING,omitempty"`

	// generated from element "ATTRIBUTE-VALUE-XHTML" of type ATTRIBUTE-VALUE-XHTML order 40 depth 1
	ATTRIBUTE_VALUE_XHTML []*ATTRIBUTE_VALUE_XHTML `xml:"ATTRIBUTE-VALUE-XHTML,omitempty"`
}

// A_CHILDREN Named source within outer element "CHILDREN"
type A_CHILDREN struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "SPEC-HIERARCHY" of type SPEC-HIERARCHY order 138 depth 1
	SPEC_HIERARCHY []*SPEC_HIERARCHY `xml:"SPEC-HIERARCHY,omitempty"`
}

// A_CORE_CONTENT Named source within outer element "CORE-CONTENT"
type A_CORE_CONTENT struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "REQ-IF-CONTENT" of type REQ-IF-CONTENT order 104 depth 1
	REQ_IF_CONTENT []*REQ_IF_CONTENT `xml:"REQ-IF-CONTENT,omitempty"`
}

// A_DATATYPE_DEFINITION_BOOLEAN_REF Named source within outer element "TYPE"
type A_DATATYPE_DEFINITION_BOOLEAN_REF struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "DATATYPE-DEFINITION-BOOLEAN-REF" of type LOCAL-REF order 6 depth 1
	DATATYPE_DEFINITION_BOOLEAN_REF string `xml:"DATATYPE-DEFINITION-BOOLEAN-REF,omitempty"`
}

// A_DATATYPE_DEFINITION_DATE_REF Named source within outer element "TYPE"
type A_DATATYPE_DEFINITION_DATE_REF struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "DATATYPE-DEFINITION-DATE-REF" of type LOCAL-REF order 12 depth 1
	DATATYPE_DEFINITION_DATE_REF string `xml:"DATATYPE-DEFINITION-DATE-REF,omitempty"`
}

// A_DATATYPE_DEFINITION_ENUMERATION_REF Named source within outer element "TYPE"
type A_DATATYPE_DEFINITION_ENUMERATION_REF struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "DATATYPE-DEFINITION-ENUMERATION-REF" of type LOCAL-REF order 18 depth 1
	DATATYPE_DEFINITION_ENUMERATION_REF string `xml:"DATATYPE-DEFINITION-ENUMERATION-REF,omitempty"`
}

// A_DATATYPE_DEFINITION_INTEGER_REF Named source within outer element "TYPE"
type A_DATATYPE_DEFINITION_INTEGER_REF struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "DATATYPE-DEFINITION-INTEGER-REF" of type LOCAL-REF order 24 depth 1
	DATATYPE_DEFINITION_INTEGER_REF string `xml:"DATATYPE-DEFINITION-INTEGER-REF,omitempty"`
}

// A_DATATYPE_DEFINITION_REAL_REF Named source within outer element "TYPE"
type A_DATATYPE_DEFINITION_REAL_REF struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "DATATYPE-DEFINITION-REAL-REF" of type LOCAL-REF order 30 depth 1
	DATATYPE_DEFINITION_REAL_REF string `xml:"DATATYPE-DEFINITION-REAL-REF,omitempty"`
}

// A_DATATYPE_DEFINITION_STRING_REF Named source within outer element "TYPE"
type A_DATATYPE_DEFINITION_STRING_REF struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "DATATYPE-DEFINITION-STRING-REF" of type LOCAL-REF order 36 depth 1
	DATATYPE_DEFINITION_STRING_REF string `xml:"DATATYPE-DEFINITION-STRING-REF,omitempty"`
}

// A_DATATYPE_DEFINITION_XHTML_REF Named source within outer element "TYPE"
type A_DATATYPE_DEFINITION_XHTML_REF struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "DATATYPE-DEFINITION-XHTML-REF" of type LOCAL-REF order 42 depth 1
	DATATYPE_DEFINITION_XHTML_REF string `xml:"DATATYPE-DEFINITION-XHTML-REF,omitempty"`
}

// A_DATATYPES Named source within outer element "DATATYPES"
type A_DATATYPES struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "DATATYPE-DEFINITION-BOOLEAN" of type DATATYPE-DEFINITION-BOOLEAN order 108 depth 1
	DATATYPE_DEFINITION_BOOLEAN []*DATATYPE_DEFINITION_BOOLEAN `xml:"DATATYPE-DEFINITION-BOOLEAN,omitempty"`

	// generated from element "DATATYPE-DEFINITION-DATE" of type DATATYPE-DEFINITION-DATE order 109 depth 1
	DATATYPE_DEFINITION_DATE []*DATATYPE_DEFINITION_DATE `xml:"DATATYPE-DEFINITION-DATE,omitempty"`

	// generated from element "DATATYPE-DEFINITION-ENUMERATION" of type DATATYPE-DEFINITION-ENUMERATION order 110 depth 1
	DATATYPE_DEFINITION_ENUMERATION []*DATATYPE_DEFINITION_ENUMERATION `xml:"DATATYPE-DEFINITION-ENUMERATION,omitempty"`

	// generated from element "DATATYPE-DEFINITION-INTEGER" of type DATATYPE-DEFINITION-INTEGER order 111 depth 1
	DATATYPE_DEFINITION_INTEGER []*DATATYPE_DEFINITION_INTEGER `xml:"DATATYPE-DEFINITION-INTEGER,omitempty"`

	// generated from element "DATATYPE-DEFINITION-REAL" of type DATATYPE-DEFINITION-REAL order 112 depth 1
	DATATYPE_DEFINITION_REAL []*DATATYPE_DEFINITION_REAL `xml:"DATATYPE-DEFINITION-REAL,omitempty"`

	// generated from element "DATATYPE-DEFINITION-STRING" of type DATATYPE-DEFINITION-STRING order 113 depth 1
	DATATYPE_DEFINITION_STRING []*DATATYPE_DEFINITION_STRING `xml:"DATATYPE-DEFINITION-STRING,omitempty"`

	// generated from element "DATATYPE-DEFINITION-XHTML" of type DATATYPE-DEFINITION-XHTML order 114 depth 1
	DATATYPE_DEFINITION_XHTML []*DATATYPE_DEFINITION_XHTML `xml:"DATATYPE-DEFINITION-XHTML,omitempty"`
}

// A_EDITABLE_ATTS Named source within outer element "EDITABLE-ATTS"
type A_EDITABLE_ATTS struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "ATTRIBUTE-DEFINITION-BOOLEAN-REF" of type LOCAL-REF order 44 depth 1
	ATTRIBUTE_DEFINITION_BOOLEAN_REF string `xml:"ATTRIBUTE-DEFINITION-BOOLEAN-REF,omitempty"`

	// generated from element "ATTRIBUTE-DEFINITION-DATE-REF" of type LOCAL-REF order 46 depth 1
	ATTRIBUTE_DEFINITION_DATE_REF string `xml:"ATTRIBUTE-DEFINITION-DATE-REF,omitempty"`

	// generated from element "ATTRIBUTE-DEFINITION-ENUMERATION-REF" of type LOCAL-REF order 48 depth 1
	ATTRIBUTE_DEFINITION_ENUMERATION_REF string `xml:"ATTRIBUTE-DEFINITION-ENUMERATION-REF,omitempty"`

	// generated from element "ATTRIBUTE-DEFINITION-INTEGER-REF" of type LOCAL-REF order 52 depth 1
	ATTRIBUTE_DEFINITION_INTEGER_REF string `xml:"ATTRIBUTE-DEFINITION-INTEGER-REF,omitempty"`

	// generated from element "ATTRIBUTE-DEFINITION-REAL-REF" of type LOCAL-REF order 54 depth 1
	ATTRIBUTE_DEFINITION_REAL_REF string `xml:"ATTRIBUTE-DEFINITION-REAL-REF,omitempty"`

	// generated from element "ATTRIBUTE-DEFINITION-STRING-REF" of type LOCAL-REF order 56 depth 1
	ATTRIBUTE_DEFINITION_STRING_REF string `xml:"ATTRIBUTE-DEFINITION-STRING-REF,omitempty"`

	// generated from element "ATTRIBUTE-DEFINITION-XHTML-REF" of type LOCAL-REF order 60 depth 1
	ATTRIBUTE_DEFINITION_XHTML_REF string `xml:"ATTRIBUTE-DEFINITION-XHTML-REF,omitempty"`
}

// A_ENUM_VALUE_REF Named source within outer element "VALUES"
type A_ENUM_VALUE_REF struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "ENUM-VALUE-REF" of type LOCAL-REF order 50 depth 1
	ENUM_VALUE_REF string `xml:"ENUM-VALUE-REF,omitempty"`
}

// A_OBJECT Named source within outer element "OBJECT"
type A_OBJECT struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "SPEC-OBJECT-REF" of type LOCAL-REF order 148 depth 1
	SPEC_OBJECT_REF string `xml:"SPEC-OBJECT-REF,omitempty"`
}

// A_PROPERTIES Named source within outer element "PROPERTIES"
type A_PROPERTIES struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "EMBEDDED-VALUE" of type EMBEDDED-VALUE order 80 depth 1
	EMBEDDED_VALUE []*EMBEDDED_VALUE `xml:"EMBEDDED-VALUE,omitempty"`
}

// A_RELATION_GROUP_TYPE_REF Named source within outer element "TYPE"
type A_RELATION_GROUP_TYPE_REF struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "RELATION-GROUP-TYPE-REF" of type LOCAL-REF order 90 depth 1
	RELATION_GROUP_TYPE_REF string `xml:"RELATION-GROUP-TYPE-REF,omitempty"`
}

// A_SPEC_ATTRIBUTES Named source within outer element "SPEC-ATTRIBUTES"
type A_SPEC_ATTRIBUTES struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "ATTRIBUTE-DEFINITION-BOOLEAN" of type ATTRIBUTE-DEFINITION-BOOLEAN order 94 depth 1
	ATTRIBUTE_DEFINITION_BOOLEAN []*ATTRIBUTE_DEFINITION_BOOLEAN `xml:"ATTRIBUTE-DEFINITION-BOOLEAN,omitempty"`

	// generated from element "ATTRIBUTE-DEFINITION-DATE" of type ATTRIBUTE-DEFINITION-DATE order 95 depth 1
	ATTRIBUTE_DEFINITION_DATE []*ATTRIBUTE_DEFINITION_DATE `xml:"ATTRIBUTE-DEFINITION-DATE,omitempty"`

	// generated from element "ATTRIBUTE-DEFINITION-ENUMERATION" of type ATTRIBUTE-DEFINITION-ENUMERATION order 96 depth 1
	ATTRIBUTE_DEFINITION_ENUMERATION []*ATTRIBUTE_DEFINITION_ENUMERATION `xml:"ATTRIBUTE-DEFINITION-ENUMERATION,omitempty"`

	// generated from element "ATTRIBUTE-DEFINITION-INTEGER" of type ATTRIBUTE-DEFINITION-INTEGER order 97 depth 1
	ATTRIBUTE_DEFINITION_INTEGER []*ATTRIBUTE_DEFINITION_INTEGER `xml:"ATTRIBUTE-DEFINITION-INTEGER,omitempty"`

	// generated from element "ATTRIBUTE-DEFINITION-REAL" of type ATTRIBUTE-DEFINITION-REAL order 98 depth 1
	ATTRIBUTE_DEFINITION_REAL []*ATTRIBUTE_DEFINITION_REAL `xml:"ATTRIBUTE-DEFINITION-REAL,omitempty"`

	// generated from element "ATTRIBUTE-DEFINITION-STRING" of type ATTRIBUTE-DEFINITION-STRING order 99 depth 1
	ATTRIBUTE_DEFINITION_STRING []*ATTRIBUTE_DEFINITION_STRING `xml:"ATTRIBUTE-DEFINITION-STRING,omitempty"`

	// generated from element "ATTRIBUTE-DEFINITION-XHTML" of type ATTRIBUTE-DEFINITION-XHTML order 100 depth 1
	ATTRIBUTE_DEFINITION_XHTML []*ATTRIBUTE_DEFINITION_XHTML `xml:"ATTRIBUTE-DEFINITION-XHTML,omitempty"`
}

// A_SPEC_OBJECT_TYPE_REF Named source within outer element "TYPE"
type A_SPEC_OBJECT_TYPE_REF struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "SPEC-OBJECT-TYPE-REF" of type LOCAL-REF order 160 depth 1
	SPEC_OBJECT_TYPE_REF string `xml:"SPEC-OBJECT-TYPE-REF,omitempty"`
}

// A_SPEC_OBJECTS Named source within outer element "SPEC-OBJECTS"
type A_SPEC_OBJECTS struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "SPEC-OBJECT" of type SPEC-OBJECT order 121 depth 1
	SPEC_OBJECT []*SPEC_OBJECT `xml:"SPEC-OBJECT,omitempty"`
}

// A_SPEC_RELATION_GROUPS Named source within outer element "SPEC-RELATION-GROUPS"
type A_SPEC_RELATION_GROUPS struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "RELATION-GROUP" of type RELATION-GROUP order 127 depth 1
	RELATION_GROUP []*RELATION_GROUP `xml:"RELATION-GROUP,omitempty"`
}

// A_SPEC_RELATION_REF Named source within outer element "SPEC-RELATIONS"
type A_SPEC_RELATION_REF struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "SPEC-RELATION-REF" of type LOCAL-REF order 86 depth 1
	SPEC_RELATION_REF string `xml:"SPEC-RELATION-REF,omitempty"`
}

// A_SPEC_RELATION_TYPE_REF Named source within outer element "TYPE"
type A_SPEC_RELATION_TYPE_REF struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "SPEC-RELATION-TYPE-REF" of type LOCAL-REF order 186 depth 1
	SPEC_RELATION_TYPE_REF string `xml:"SPEC-RELATION-TYPE-REF,omitempty"`
}

// A_SPEC_RELATIONS Named source within outer element "SPEC-RELATIONS"
type A_SPEC_RELATIONS struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "SPEC-RELATION" of type SPEC-RELATION order 123 depth 1
	SPEC_RELATION []*SPEC_RELATION `xml:"SPEC-RELATION,omitempty"`
}

// A_SPEC_TYPES Named source within outer element "SPEC-TYPES"
type A_SPEC_TYPES struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "RELATION-GROUP-TYPE" of type RELATION-GROUP-TYPE order 116 depth 1
	RELATION_GROUP_TYPE []*RELATION_GROUP_TYPE `xml:"RELATION-GROUP-TYPE,omitempty"`

	// generated from element "SPEC-OBJECT-TYPE" of type SPEC-OBJECT-TYPE order 117 depth 1
	SPEC_OBJECT_TYPE []*SPEC_OBJECT_TYPE `xml:"SPEC-OBJECT-TYPE,omitempty"`

	// generated from element "SPEC-RELATION-TYPE" of type SPEC-RELATION-TYPE order 118 depth 1
	SPEC_RELATION_TYPE []*SPEC_RELATION_TYPE `xml:"SPEC-RELATION-TYPE,omitempty"`

	// generated from element "SPECIFICATION-TYPE" of type SPECIFICATION-TYPE order 119 depth 1
	SPECIFICATION_TYPE []*SPECIFICATION_TYPE `xml:"SPECIFICATION-TYPE,omitempty"`
}

// A_SPECIFICATION_TYPE_REF Named source within outer element "TYPE"
type A_SPECIFICATION_TYPE_REF struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "SPECIFICATION-TYPE-REF" of type LOCAL-REF order 210 depth 1
	SPECIFICATION_TYPE_REF string `xml:"SPECIFICATION-TYPE-REF,omitempty"`
}

// A_SPECIFICATIONS Named source within outer element "SPECIFICATIONS"
type A_SPECIFICATIONS struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "SPECIFICATION" of type SPECIFICATION order 125 depth 1
	SPECIFICATION []*SPECIFICATION `xml:"SPECIFICATION,omitempty"`
}

// A_SPECIFIED_VALUES Named source within outer element "SPECIFIED-VALUES"
type A_SPECIFIED_VALUES struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "ENUM-VALUE" of type ENUM-VALUE order 68 depth 1
	ENUM_VALUE []*ENUM_VALUE `xml:"ENUM-VALUE,omitempty"`
}

// A_TARGET_1 Named source within outer element "SOURCE"
type A_TARGET_1 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "SPEC-OBJECT-REF" of type GLOBAL-REF order 182 depth 1
	SPEC_OBJECT_REF string `xml:"SPEC-OBJECT-REF,omitempty"`
}

// A_TARGET_SPECIFICATION_1 Named source within outer element "SOURCE-SPECIFICATION"
type A_TARGET_SPECIFICATION_1 struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "SPECIFICATION-REF" of type GLOBAL-REF order 84 depth 1
	SPECIFICATION_REF string `xml:"SPECIFICATION-REF,omitempty"`
}

// A_THE_HEADER Named source within out	er element "THE-HEADER"
type A_THE_HEADER struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "REQ-IF-HEADER" of type REQ-IF-HEADER order 102 depth 1
	REQ_IF_HEADER []*REQ_IF_HEADER `xml:"REQ-IF-HEADER,omitempty"`
}

// A_TOOL_EXTENSIONS Named source within outer element "TOOL-EXTENSIONS"
type A_TOOL_EXTENSIONS struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "REQ-IF-TOOL-EXTENSION" of type REQ-IF-TOOL-EXTENSION order 106 depth 1
	REQ_IF_TOOL_EXTENSION []*REQ_IF_TOOL_EXTENSION `xml:"REQ-IF-TOOL-EXTENSION,omitempty"`
}

// DATATYPE_DEFINITION_BOOLEAN Named source named complex type "DATATYPE-DEFINITION-BOOLEAN"
type DATATYPE_DEFINITION_BOOLEAN struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "DESC
	DESC string `xml:"DESC,attr,omitempty"`

	// generated from attribute "IDENTIFIER
	IDENTIFIER string `xml:"IDENTIFIER,attr,omitempty"`

	// generated from attribute "LAST-CHANGE
	LAST_CHANGE string `xml:"LAST-CHANGE,attr,omitempty"`

	// generated from attribute "LONG-NAME
	LONG_NAME string `xml:"LONG-NAME,attr,omitempty"`

	// generated from anonymous type within outer element "ALTERNATIVE-ID" of type A_ALTERNATIVE-ID.
	ALTERNATIVE_ID []*A_ALTERNATIVE_ID `xml:"ALTERNATIVE-ID,omitempty"`
}

// DATATYPE_DEFINITION_DATE Named source named complex type "DATATYPE-DEFINITION-DATE"
type DATATYPE_DEFINITION_DATE struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "DESC
	DESC string `xml:"DESC,attr,omitempty"`

	// generated from attribute "IDENTIFIER
	IDENTIFIER string `xml:"IDENTIFIER,attr,omitempty"`

	// generated from attribute "LAST-CHANGE
	LAST_CHANGE string `xml:"LAST-CHANGE,attr,omitempty"`

	// generated from attribute "LONG-NAME
	LONG_NAME string `xml:"LONG-NAME,attr,omitempty"`

	// generated from anonymous type within outer element "ALTERNATIVE-ID" of type A_ALTERNATIVE-ID.
	ALTERNATIVE_ID []*A_ALTERNATIVE_ID `xml:"ALTERNATIVE-ID,omitempty"`
}

// DATATYPE_DEFINITION_ENUMERATION Named source named complex type "DATATYPE-DEFINITION-ENUMERATION"
type DATATYPE_DEFINITION_ENUMERATION struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "DESC
	DESC string `xml:"DESC,attr,omitempty"`

	// generated from attribute "IDENTIFIER
	IDENTIFIER string `xml:"IDENTIFIER,attr,omitempty"`

	// generated from attribute "LAST-CHANGE
	LAST_CHANGE string `xml:"LAST-CHANGE,attr,omitempty"`

	// generated from attribute "LONG-NAME
	LONG_NAME string `xml:"LONG-NAME,attr,omitempty"`

	// generated from anonymous type within outer element "ALTERNATIVE-ID" of type A_ALTERNATIVE-ID.
	ALTERNATIVE_ID []*A_ALTERNATIVE_ID `xml:"ALTERNATIVE-ID,omitempty"`

	// generated from anonymous type within outer element "SPECIFIED-VALUES" of type A_SPECIFIED-VALUES.
	SPECIFIED_VALUES []*A_SPECIFIED_VALUES `xml:"SPECIFIED-VALUES,omitempty"`
}

// DATATYPE_DEFINITION_INTEGER Named source named complex type "DATATYPE-DEFINITION-INTEGER"
type DATATYPE_DEFINITION_INTEGER struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "DESC
	DESC string `xml:"DESC,attr,omitempty"`

	// generated from attribute "IDENTIFIER
	IDENTIFIER string `xml:"IDENTIFIER,attr,omitempty"`

	// generated from attribute "LAST-CHANGE
	LAST_CHANGE string `xml:"LAST-CHANGE,attr,omitempty"`

	// generated from attribute "LONG-NAME
	LONG_NAME string `xml:"LONG-NAME,attr,omitempty"`

	// generated from attribute "MAX
	MAX int `xml:"MAX,attr,omitempty"`

	// generated from attribute "MIN
	MIN int `xml:"MIN,attr,omitempty"`

	// generated from anonymous type within outer element "ALTERNATIVE-ID" of type A_ALTERNATIVE-ID.
	ALTERNATIVE_ID []*A_ALTERNATIVE_ID `xml:"ALTERNATIVE-ID,omitempty"`
}

// DATATYPE_DEFINITION_REAL Named source named complex type "DATATYPE-DEFINITION-REAL"
type DATATYPE_DEFINITION_REAL struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "ACCURACY
	ACCURACY int `xml:"ACCURACY,attr,omitempty"`

	// generated from attribute "DESC
	DESC string `xml:"DESC,attr,omitempty"`

	// generated from attribute "IDENTIFIER
	IDENTIFIER string `xml:"IDENTIFIER,attr,omitempty"`

	// generated from attribute "LAST-CHANGE
	LAST_CHANGE string `xml:"LAST-CHANGE,attr,omitempty"`

	// generated from attribute "LONG-NAME
	LONG_NAME string `xml:"LONG-NAME,attr,omitempty"`

	// generated from attribute "MAX
	MAX float64 `xml:"MAX,attr,omitempty"`

	// generated from attribute "MIN
	MIN float64 `xml:"MIN,attr,omitempty"`

	// generated from anonymous type within outer element "ALTERNATIVE-ID" of type A_ALTERNATIVE-ID.
	ALTERNATIVE_ID []*A_ALTERNATIVE_ID `xml:"ALTERNATIVE-ID,omitempty"`
}

// DATATYPE_DEFINITION_STRING Named source named complex type "DATATYPE-DEFINITION-STRING"
type DATATYPE_DEFINITION_STRING struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "DESC
	DESC string `xml:"DESC,attr,omitempty"`

	// generated from attribute "IDENTIFIER
	IDENTIFIER string `xml:"IDENTIFIER,attr,omitempty"`

	// generated from attribute "LAST-CHANGE
	LAST_CHANGE string `xml:"LAST-CHANGE,attr,omitempty"`

	// generated from attribute "LONG-NAME
	LONG_NAME string `xml:"LONG-NAME,attr,omitempty"`

	// generated from attribute "MAX-LENGTH
	MAX_LENGTH int `xml:"MAX-LENGTH,attr,omitempty"`

	// generated from anonymous type within outer element "ALTERNATIVE-ID" of type A_ALTERNATIVE-ID.
	ALTERNATIVE_ID []*A_ALTERNATIVE_ID `xml:"ALTERNATIVE-ID,omitempty"`
}

// DATATYPE_DEFINITION_XHTML Named source named complex type "DATATYPE-DEFINITION-XHTML"
type DATATYPE_DEFINITION_XHTML struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "DESC
	DESC string `xml:"DESC,attr,omitempty"`

	// generated from attribute "IDENTIFIER
	IDENTIFIER string `xml:"IDENTIFIER,attr,omitempty"`

	// generated from attribute "LAST-CHANGE
	LAST_CHANGE string `xml:"LAST-CHANGE,attr,omitempty"`

	// generated from attribute "LONG-NAME
	LONG_NAME string `xml:"LONG-NAME,attr,omitempty"`

	// generated from anonymous type within outer element "ALTERNATIVE-ID" of type A_ALTERNATIVE-ID.
	ALTERNATIVE_ID []*A_ALTERNATIVE_ID `xml:"ALTERNATIVE-ID,omitempty"`
}

// EMBEDDED_VALUE Named source named complex type "EMBEDDED-VALUE"
type EMBEDDED_VALUE struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "KEY
	KEY int `xml:"KEY,attr,omitempty"`

	// generated from attribute "OTHER-CONTENT
	OTHER_CONTENT string `xml:"OTHER-CONTENT,attr,omitempty"`
}

// ENUM_VALUE Named source named complex type "ENUM-VALUE"
type ENUM_VALUE struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "DESC
	DESC string `xml:"DESC,attr,omitempty"`

	// generated from attribute "IDENTIFIER
	IDENTIFIER string `xml:"IDENTIFIER,attr,omitempty"`

	// generated from attribute "LAST-CHANGE
	LAST_CHANGE string `xml:"LAST-CHANGE,attr,omitempty"`

	// generated from attribute "LONG-NAME
	LONG_NAME string `xml:"LONG-NAME,attr,omitempty"`

	// generated from anonymous type within outer element "ALTERNATIVE-ID" of type A_ALTERNATIVE-ID.
	ALTERNATIVE_ID []*A_ALTERNATIVE_ID `xml:"ALTERNATIVE-ID,omitempty"`

	// generated from anonymous type within outer element "PROPERTIES" of type A_PROPERTIES.
	PROPERTIES []*A_PROPERTIES `xml:"PROPERTIES,omitempty"`
}

// RELATION_GROUP Named source named complex type "RELATION-GROUP"
type RELATION_GROUP struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "DESC
	DESC string `xml:"DESC,attr,omitempty"`

	// generated from attribute "IDENTIFIER
	IDENTIFIER string `xml:"IDENTIFIER,attr,omitempty"`

	// generated from attribute "LAST-CHANGE
	LAST_CHANGE string `xml:"LAST-CHANGE,attr,omitempty"`

	// generated from attribute "LONG-NAME
	LONG_NAME string `xml:"LONG-NAME,attr,omitempty"`

	// generated from anonymous type within outer element "ALTERNATIVE-ID" of type A_ALTERNATIVE-ID.
	ALTERNATIVE_ID []*A_ALTERNATIVE_ID `xml:"ALTERNATIVE-ID,omitempty"`

	// generated from anonymous type within outer element "SOURCE-SPECIFICATION" of type A_TARGET-SPECIFICATION.
	SOURCE_SPECIFICATION []*A_TARGET_SPECIFICATION_1 `xml:"SOURCE-SPECIFICATION,omitempty"`

	// generated from anonymous type within outer element "SPEC-RELATIONS" of type A_SPEC-RELATION-REF.
	SPEC_RELATIONS []*A_SPEC_RELATION_REF `xml:"SPEC-RELATIONS,omitempty"`

	// generated from anonymous type within outer element "TARGET-SPECIFICATION" of type A_TARGET-SPECIFICATION.
	TARGET_SPECIFICATION []*A_TARGET_SPECIFICATION_1 `xml:"TARGET-SPECIFICATION,omitempty"`

	// generated from anonymous type within outer element "TYPE" of type A_RELATION-GROUP-TYPE-REF.
	TYPE []*A_RELATION_GROUP_TYPE_REF `xml:"TYPE,omitempty"`
}

// RELATION_GROUP_TYPE Named source named complex type "RELATION-GROUP-TYPE"
type RELATION_GROUP_TYPE struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "DESC
	DESC string `xml:"DESC,attr,omitempty"`

	// generated from attribute "IDENTIFIER
	IDENTIFIER string `xml:"IDENTIFIER,attr,omitempty"`

	// generated from attribute "LAST-CHANGE
	LAST_CHANGE string `xml:"LAST-CHANGE,attr,omitempty"`

	// generated from attribute "LONG-NAME
	LONG_NAME string `xml:"LONG-NAME,attr,omitempty"`

	// generated from anonymous type within outer element "ALTERNATIVE-ID" of type A_ALTERNATIVE-ID.
	ALTERNATIVE_ID []*A_ALTERNATIVE_ID `xml:"ALTERNATIVE-ID,omitempty"`

	// generated from anonymous type within outer element "SPEC-ATTRIBUTES" of type A_SPEC-ATTRIBUTES.
	SPEC_ATTRIBUTES []*A_SPEC_ATTRIBUTES `xml:"SPEC-ATTRIBUTES,omitempty"`
}

// REQ_IF Named source named complex type "REQ-IF"
type REQ_IF struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "http://www.w3.org/XML/1998/namespace lang
	Lang string `xml:"http://www.w3.org/XML/1998/namespace lang,attr,omitempty"`

	// generated from anonymous type within outer element "THE-HEADER" of type A_THE-HEADER.
	THE_HEADER []*A_THE_HEADER `xml:"THE-HEADER,omitempty"`

	// generated from anonymous type within outer element "CORE-CONTENT" of type A_CORE-CONTENT.
	CORE_CONTENT []*A_CORE_CONTENT `xml:"CORE-CONTENT,omitempty"`

	// generated from anonymous type within outer element "TOOL-EXTENSIONS" of type A_TOOL-EXTENSIONS.
	TOOL_EXTENSIONS []*A_TOOL_EXTENSIONS `xml:"TOOL-EXTENSIONS,omitempty"`
}

// REQ_IF_CONTENT Named source named complex type "REQ-IF-CONTENT"
type REQ_IF_CONTENT struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from anonymous type within outer element "DATATYPES" of type A_DATATYPES.
	DATATYPES []*A_DATATYPES `xml:"DATATYPES,omitempty"`

	// generated from anonymous type within outer element "SPEC-TYPES" of type A_SPEC-TYPES.
	SPEC_TYPES []*A_SPEC_TYPES `xml:"SPEC-TYPES,omitempty"`

	// generated from anonymous type within outer element "SPEC-OBJECTS" of type A_SPEC-OBJECTS.
	SPEC_OBJECTS []*A_SPEC_OBJECTS `xml:"SPEC-OBJECTS,omitempty"`

	// generated from anonymous type within outer element "SPEC-RELATIONS" of type A_SPEC-RELATIONS.
	SPEC_RELATIONS []*A_SPEC_RELATIONS `xml:"SPEC-RELATIONS,omitempty"`

	// generated from anonymous type within outer element "SPECIFICATIONS" of type A_SPECIFICATIONS.
	SPECIFICATIONS []*A_SPECIFICATIONS `xml:"SPECIFICATIONS,omitempty"`

	// generated from anonymous type within outer element "SPEC-RELATION-GROUPS" of type A_SPEC-RELATION-GROUPS.
	SPEC_RELATION_GROUPS []*A_SPEC_RELATION_GROUPS `xml:"SPEC-RELATION-GROUPS,omitempty"`
}

// REQ_IF_HEADER Named source named complex type "REQ-IF-HEADER"
type REQ_IF_HEADER struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "IDENTIFIER
	IDENTIFIER string `xml:"IDENTIFIER,attr,omitempty"`

	// generated from element "COMMENT" of type string order 128 depth 0
	COMMENT string `xml:"COMMENT,omitempty"`

	// generated from element "CREATION-TIME" of type dateTime order 129 depth 0
	CREATION_TIME string `xml:"CREATION-TIME,omitempty"`

	// generated from element "REPOSITORY-ID" of type string order 130 depth 0
	REPOSITORY_ID string `xml:"REPOSITORY-ID,omitempty"`

	// generated from element "REQ-IF-TOOL-ID" of type string order 131 depth 0
	REQ_IF_TOOL_ID string `xml:"REQ-IF-TOOL-ID,omitempty"`

	// generated from element "REQ-IF-VERSION" of type string order 132 depth 0
	REQ_IF_VERSION string `xml:"REQ-IF-VERSION,omitempty"`

	// generated from element "SOURCE-TOOL-ID" of type string order 133 depth 0
	SOURCE_TOOL_ID string `xml:"SOURCE-TOOL-ID,omitempty"`

	// generated from element "TITLE" of type string order 134 depth 0
	TITLE string `xml:"TITLE,omitempty"`
}

// REQ_IF_TOOL_EXTENSION Named source named complex type "REQ-IF-TOOL-EXTENSION"
type REQ_IF_TOOL_EXTENSION struct {
	Name string `xml:"-"`

	// insertion point for fields
}

// SPEC_HIERARCHY Named source named complex type "SPEC-HIERARCHY"
type SPEC_HIERARCHY struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "DESC
	DESC string `xml:"DESC,attr,omitempty"`

	// generated from attribute "IDENTIFIER
	IDENTIFIER string `xml:"IDENTIFIER,attr,omitempty"`

	// generated from attribute "IS-EDITABLE
	IS_EDITABLE bool `xml:"IS-EDITABLE,attr,omitempty"`

	// generated from attribute "IS-TABLE-INTERNAL
	IS_TABLE_INTERNAL bool `xml:"IS-TABLE-INTERNAL,attr,omitempty"`

	// generated from attribute "LAST-CHANGE
	LAST_CHANGE string `xml:"LAST-CHANGE,attr,omitempty"`

	// generated from attribute "LONG-NAME
	LONG_NAME string `xml:"LONG-NAME,attr,omitempty"`

	// generated from anonymous type within outer element "ALTERNATIVE-ID" of type A_ALTERNATIVE-ID.
	ALTERNATIVE_ID []*A_ALTERNATIVE_ID `xml:"ALTERNATIVE-ID,omitempty"`

	// generated from anonymous type within outer element "CHILDREN" of type A_CHILDREN.
	CHILDREN []*A_CHILDREN `xml:"CHILDREN,omitempty"`

	// generated from anonymous type within outer element "EDITABLE-ATTS" of type A_EDITABLE-ATTS.
	EDITABLE_ATTS []*A_EDITABLE_ATTS `xml:"EDITABLE-ATTS,omitempty"`

	// generated from anonymous type within outer element "OBJECT" of type A_OBJECT.
	OBJECT []*A_OBJECT `xml:"OBJECT,omitempty"`
}

// SPEC_OBJECT Named source named complex type "SPEC-OBJECT"
type SPEC_OBJECT struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "DESC
	DESC string `xml:"DESC,attr,omitempty"`

	// generated from attribute "IDENTIFIER
	IDENTIFIER string `xml:"IDENTIFIER,attr,omitempty"`

	// generated from attribute "LAST-CHANGE
	LAST_CHANGE string `xml:"LAST-CHANGE,attr,omitempty"`

	// generated from attribute "LONG-NAME
	LONG_NAME string `xml:"LONG-NAME,attr,omitempty"`

	// generated from anonymous type within outer element "ALTERNATIVE-ID" of type A_ALTERNATIVE-ID.
	ALTERNATIVE_ID []*A_ALTERNATIVE_ID `xml:"ALTERNATIVE-ID,omitempty"`

	// generated from anonymous type within outer element "VALUES" of type A_ATTRIBUTE-VALUE-XHTML.
	VALUES []*A_ATTRIBUTE_VALUE_XHTML_1 `xml:"VALUES,omitempty"`

	// generated from anonymous type within outer element "TYPE" of type A_SPEC-OBJECT-TYPE-REF.
	TYPE []*A_SPEC_OBJECT_TYPE_REF `xml:"TYPE,omitempty"`
}

// SPEC_OBJECT_TYPE Named source named complex type "SPEC-OBJECT-TYPE"
type SPEC_OBJECT_TYPE struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "DESC
	DESC string `xml:"DESC,attr,omitempty"`

	// generated from attribute "IDENTIFIER
	IDENTIFIER string `xml:"IDENTIFIER,attr,omitempty"`

	// generated from attribute "LAST-CHANGE
	LAST_CHANGE string `xml:"LAST-CHANGE,attr,omitempty"`

	// generated from attribute "LONG-NAME
	LONG_NAME string `xml:"LONG-NAME,attr,omitempty"`

	// generated from anonymous type within outer element "ALTERNATIVE-ID" of type A_ALTERNATIVE-ID.
	ALTERNATIVE_ID []*A_ALTERNATIVE_ID `xml:"ALTERNATIVE-ID,omitempty"`

	// generated from anonymous type within outer element "SPEC-ATTRIBUTES" of type A_SPEC-ATTRIBUTES.
	SPEC_ATTRIBUTES []*A_SPEC_ATTRIBUTES `xml:"SPEC-ATTRIBUTES,omitempty"`
}

// SPEC_RELATION Named source named complex type "SPEC-RELATION"
type SPEC_RELATION struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "DESC
	DESC string `xml:"DESC,attr,omitempty"`

	// generated from attribute "IDENTIFIER
	IDENTIFIER string `xml:"IDENTIFIER,attr,omitempty"`

	// generated from attribute "LAST-CHANGE
	LAST_CHANGE string `xml:"LAST-CHANGE,attr,omitempty"`

	// generated from attribute "LONG-NAME
	LONG_NAME string `xml:"LONG-NAME,attr,omitempty"`

	// generated from anonymous type within outer element "ALTERNATIVE-ID" of type A_ALTERNATIVE-ID.
	ALTERNATIVE_ID []*A_ALTERNATIVE_ID `xml:"ALTERNATIVE-ID,omitempty"`

	// generated from anonymous type within outer element "VALUES" of type A_ATTRIBUTE-VALUE-XHTML.
	VALUES []*A_ATTRIBUTE_VALUE_XHTML_1 `xml:"VALUES,omitempty"`

	// generated from anonymous type within outer element "SOURCE" of type A_TARGET.
	SOURCE []*A_TARGET_1 `xml:"SOURCE,omitempty"`

	// generated from anonymous type within outer element "TARGET" of type A_TARGET.
	TARGET []*A_TARGET_1 `xml:"TARGET,omitempty"`

	// generated from anonymous type within outer element "TYPE" of type A_SPEC-RELATION-TYPE-REF.
	TYPE []*A_SPEC_RELATION_TYPE_REF `xml:"TYPE,omitempty"`
}

// SPEC_RELATION_TYPE Named source named complex type "SPEC-RELATION-TYPE"
type SPEC_RELATION_TYPE struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "DESC
	DESC string `xml:"DESC,attr,omitempty"`

	// generated from attribute "IDENTIFIER
	IDENTIFIER string `xml:"IDENTIFIER,attr,omitempty"`

	// generated from attribute "LAST-CHANGE
	LAST_CHANGE string `xml:"LAST-CHANGE,attr,omitempty"`

	// generated from attribute "LONG-NAME
	LONG_NAME string `xml:"LONG-NAME,attr,omitempty"`

	// generated from anonymous type within outer element "ALTERNATIVE-ID" of type A_ALTERNATIVE-ID.
	ALTERNATIVE_ID []*A_ALTERNATIVE_ID `xml:"ALTERNATIVE-ID,omitempty"`

	// generated from anonymous type within outer element "SPEC-ATTRIBUTES" of type A_SPEC-ATTRIBUTES.
	SPEC_ATTRIBUTES []*A_SPEC_ATTRIBUTES `xml:"SPEC-ATTRIBUTES,omitempty"`
}

// SPECIFICATION Named source named complex type "SPECIFICATION"
type SPECIFICATION struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "DESC
	DESC string `xml:"DESC,attr,omitempty"`

	// generated from attribute "IDENTIFIER
	IDENTIFIER string `xml:"IDENTIFIER,attr,omitempty"`

	// generated from attribute "LAST-CHANGE
	LAST_CHANGE string `xml:"LAST-CHANGE,attr,omitempty"`

	// generated from attribute "LONG-NAME
	LONG_NAME string `xml:"LONG-NAME,attr,omitempty"`

	// generated from anonymous type within outer element "ALTERNATIVE-ID" of type A_ALTERNATIVE-ID.
	ALTERNATIVE_ID []*A_ALTERNATIVE_ID `xml:"ALTERNATIVE-ID,omitempty"`

	// generated from anonymous type within outer element "CHILDREN" of type A_CHILDREN.
	CHILDREN []*A_CHILDREN `xml:"CHILDREN,omitempty"`

	// generated from anonymous type within outer element "VALUES" of type A_ATTRIBUTE-VALUE-XHTML.
	VALUES []*A_ATTRIBUTE_VALUE_XHTML_1 `xml:"VALUES,omitempty"`

	// generated from anonymous type within outer element "TYPE" of type A_SPECIFICATION-TYPE-REF.
	TYPE []*A_SPECIFICATION_TYPE_REF `xml:"TYPE,omitempty"`
}

// SPECIFICATION_TYPE Named source named complex type "SPECIFICATION-TYPE"
type SPECIFICATION_TYPE struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "DESC
	DESC string `xml:"DESC,attr,omitempty"`

	// generated from attribute "IDENTIFIER
	IDENTIFIER string `xml:"IDENTIFIER,attr,omitempty"`

	// generated from attribute "LAST-CHANGE
	LAST_CHANGE string `xml:"LAST-CHANGE,attr,omitempty"`

	// generated from attribute "LONG-NAME
	LONG_NAME string `xml:"LONG-NAME,attr,omitempty"`

	// generated from anonymous type within outer element "ALTERNATIVE-ID" of type A_ALTERNATIVE-ID.
	ALTERNATIVE_ID []*A_ALTERNATIVE_ID `xml:"ALTERNATIVE-ID,omitempty"`

	// generated from anonymous type within outer element "SPEC-ATTRIBUTES" of type A_SPEC-ATTRIBUTES.
	SPEC_ATTRIBUTES []*A_SPEC_ATTRIBUTES `xml:"SPEC-ATTRIBUTES,omitempty"`
}

// XHTML_CONTENT Named source named complex type "XHTML-CONTENT"
type XHTML_CONTENT struct {
	Name string `xml:"-"`

	// insertion point for fields
}
