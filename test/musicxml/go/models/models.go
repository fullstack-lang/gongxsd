// generated code - do not edit
package models

import "encoding/xml"

// to avoid compilation error if no xml element is generated
var _ xml.Attr

// A_directive is generated from outer element "directive"
type A_directive struct {

	// insertion point for fields

	// generated from attribute "http://www.w3.org/XML/1998/namespace lang
	Lang string `xml:"http://www.w3.org/XML/1998/namespace lang,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// A_measure_1 is generated from outer element "measure"
type A_measure_1 struct {

	// insertion point for fields

	// generated from attribute group "measure-attributes
	AttributeGroup_measure_attributes

	// generated from anonymous type within outer element "part" of type A_part.
	Part []*A_part `xml:"part,omitempty"`
}

// A_measure is generated from outer element "measure"
type A_measure struct {

	// insertion point for fields

	// generated from attribute group "measure-attributes
	AttributeGroup_measure_attributes

	// generated from group with order 567 depth 3
	Group_music_data
}

// A_part_1 is generated from outer element "part"
type A_part_1 struct {

	// insertion point for fields

	// generated from attribute group "part-attributes
	AttributeGroup_part_attributes

	// generated from anonymous type within outer element "measure" of type A_measure.
	Measure []*A_measure `xml:"measure,omitempty"`
}

// A_part is generated from outer element "part"
type A_part struct {

	// insertion point for fields

	// generated from attribute group "part-attributes
	AttributeGroup_part_attributes

	// generated from group with order 572 depth 3
	Group_music_data
}

// A_score_partwise is generated from outer element "score-partwise"
type A_score_partwise struct {

	// insertion point for fields

	// generated from attribute group "document-attributes
	AttributeGroup_document_attributes

	// generated from group with order 564 depth 1
	Group_score_header

	// generated from anonymous type within outer element "part" of type A_part.
	Part []*A_part_1 `xml:"part,omitempty"`
}

// A_score_timewise is generated from outer element "score-timewise"
type A_score_timewise struct {

	// insertion point for fields

	// generated from attribute group "document-attributes
	AttributeGroup_document_attributes

	// generated from group with order 569 depth 1
	Group_score_header

	// generated from anonymous type within outer element "measure" of type A_measure.
	Measure []*A_measure_1 `xml:"measure,omitempty"`
}

// Accidental is generated from named complex type "accidental"
type Accidental struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "cautionary
	Cautionary string `xml:"cautionary,attr,omitempty"`

	// generated from attribute "editorial
	Editorial string `xml:"editorial,attr,omitempty"`

	// generated from attribute "smufl
	Smufl string `xml:"smufl,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Accidental_mark is generated from named complex type "accidental-mark"
type Accidental_mark struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "smufl
	Smufl string `xml:"smufl,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Accidental_text is generated from named complex type "accidental-text"
type Accidental_text struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "smufl
	Smufl string `xml:"smufl,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Accord is generated from named complex type "accord"
type Accord struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "string
	String int `xml:"string,attr,omitempty"`

	// generated from group with order 89 depth 0
	Group_tuning
}

// Accordion_registration is generated from named complex type "accordion-registration"
type Accordion_registration struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "print-style-align
	AttributeGroup_print_style_align

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from element "accordion-high" of type empty order 90 depth 0
	Accordion_high string `xml:"accordion-high,omitempty"`

	// generated from element "accordion-middle" of type accordion-middle order 91 depth 0
	Accordion_middle int `xml:"accordion-middle,omitempty"`

	// generated from element "accordion-low" of type empty order 92 depth 0
	Accordion_low string `xml:"accordion-low,omitempty"`
}

// Appearance is generated from named complex type "appearance"
type Appearance struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "line-width" of type line-width order 214 depth 0
	Line_width []*Line_width `xml:"line-width,omitempty"`

	// generated from element "note-size" of type note-size order 215 depth 0
	Note_size []*Note_size `xml:"note-size,omitempty"`

	// generated from element "distance" of type distance order 216 depth 0
	Distance []*Distance `xml:"distance,omitempty"`

	// generated from element "glyph" of type glyph order 217 depth 0
	Glyph []*Glyph `xml:"glyph,omitempty"`

	// generated from element "other-appearance" of type other-appearance order 218 depth 0
	Other_appearance []*Other_appearance `xml:"other-appearance,omitempty"`
}

// Arpeggiate is generated from named complex type "arpeggiate"
type Arpeggiate struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "number
	Number int `xml:"number,attr,omitempty"`

	// generated from attribute "direction
	Direction string `xml:"direction,attr,omitempty"`

	// generated from attribute "unbroken
	Unbroken string `xml:"unbroken,attr,omitempty"`

	// generated from attribute group "position
	AttributeGroup_position

	// generated from attribute group "placement
	AttributeGroup_placement

	// generated from attribute group "color
	AttributeGroup_color

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id
}

// Arrow is generated from named complex type "arrow"
type Arrow struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "print-style
	AttributeGroup_print_style

	// generated from attribute group "placement
	AttributeGroup_placement

	// generated from attribute group "smufl
	AttributeGroup_smufl

	// generated from element "arrow-direction" of type arrow-direction order 251 depth 0
	Arrow_direction string `xml:"arrow-direction,omitempty"`

	// generated from element "arrow-style" of type arrow-style order 252 depth 0
	Arrow_style string `xml:"arrow-style,omitempty"`

	// generated from element "arrowhead" of type empty order 253 depth 0
	Arrowhead string `xml:"arrowhead,omitempty"`

	// generated from element "circular-arrow" of type circular-arrow order 254 depth 0
	Circular_arrow string `xml:"circular-arrow,omitempty"`
}

// Articulations is generated from named complex type "articulations"
type Articulations struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from element "accent" of type empty-placement order 234 depth 0
	Accent []*Empty_placement `xml:"accent,omitempty"`

	// generated from element "strong-accent" of type strong-accent order 235 depth 0
	Strong_accent []*Strong_accent `xml:"strong-accent,omitempty"`

	// generated from element "staccato" of type empty-placement order 236 depth 0
	Staccato []*Empty_placement `xml:"staccato,omitempty"`

	// generated from element "tenuto" of type empty-placement order 237 depth 0
	Tenuto []*Empty_placement `xml:"tenuto,omitempty"`

	// generated from element "detached-legato" of type empty-placement order 238 depth 0
	Detached_legato []*Empty_placement `xml:"detached-legato,omitempty"`

	// generated from element "staccatissimo" of type empty-placement order 239 depth 0
	Staccatissimo []*Empty_placement `xml:"staccatissimo,omitempty"`

	// generated from element "spiccato" of type empty-placement order 240 depth 0
	Spiccato []*Empty_placement `xml:"spiccato,omitempty"`

	// generated from element "scoop" of type empty-line order 241 depth 0
	Scoop []*Empty_line `xml:"scoop,omitempty"`

	// generated from element "plop" of type empty-line order 242 depth 0
	Plop []*Empty_line `xml:"plop,omitempty"`

	// generated from element "doit" of type empty-line order 243 depth 0
	Doit []*Empty_line `xml:"doit,omitempty"`

	// generated from element "falloff" of type empty-line order 244 depth 0
	Falloff []*Empty_line `xml:"falloff,omitempty"`

	// generated from element "breath-mark" of type breath-mark order 245 depth 0
	Breath_mark []*Breath_mark `xml:"breath-mark,omitempty"`

	// generated from element "caesura" of type caesura order 246 depth 0
	Caesura []*Caesura `xml:"caesura,omitempty"`

	// generated from element "stress" of type empty-placement order 247 depth 0
	Stress []*Empty_placement `xml:"stress,omitempty"`

	// generated from element "unstress" of type empty-placement order 248 depth 0
	Unstress []*Empty_placement `xml:"unstress,omitempty"`

	// generated from element "soft-accent" of type empty-placement order 249 depth 0
	Soft_accent []*Empty_placement `xml:"soft-accent,omitempty"`

	// generated from element "other-articulation" of type other-placement-text order 250 depth 0
	Other_articulation []*Other_placement_text `xml:"other-articulation,omitempty"`
}

// Assess is generated from named complex type "assess"
type Assess struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute "player
	Player string `xml:"player,attr,omitempty"`

	// generated from attribute "time-only
	Time_only string `xml:"time-only,attr,omitempty"`
}

// Attributes is generated from named complex type "attributes"
type Attributes struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from group with order 41 depth 0
	Group_editorial

	// generated from element "divisions" of type positive-divisions order 42 depth 0
	Divisions string `xml:"divisions,omitempty"`

	// generated from element "key" of type key order 43 depth 0
	Key []*Key `xml:"key,omitempty"`

	// generated from element "time" of type time order 44 depth 0
	Time []*Time `xml:"time,omitempty"`

	// generated from element "staves" of type xs:nonNegativeInteger order 45 depth 0
	Staves int `xml:"staves,omitempty"`

	// generated from element "part-symbol" of type part-symbol order 46 depth 0
	Part_symbol []*Part_symbol `xml:"part-symbol,omitempty"`

	// generated from element "instruments" of type xs:nonNegativeInteger order 47 depth 0
	Instruments int `xml:"instruments,omitempty"`

	// generated from element "clef" of type clef order 48 depth 0
	Clef []*Clef `xml:"clef,omitempty"`

	// generated from element "staff-details" of type staff-details order 49 depth 0
	Staff_details []*Staff_details `xml:"staff-details,omitempty"`

	// generated from element "transpose" of type transpose order 50 depth 0
	Transpose []*Transpose `xml:"transpose,omitempty"`

	// generated from element "for-part" of type for-part order 51 depth 0
	For_part []*For_part `xml:"for-part,omitempty"`

	// generated from anonymous type within outer element "directive" of type A_directive.
	Directive []*A_directive `xml:"directive,omitempty"`

	// generated from element "measure-style" of type measure-style order 53 depth 0
	Measure_style []*Measure_style `xml:"measure-style,omitempty"`
}

// Backup is generated from named complex type "backup"
type Backup struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from group with order 255 depth 0
	Group_duration

	// generated from group with order 256 depth 0
	Group_editorial
}

// Bar_style_color is generated from named complex type "bar-style-color"
type Bar_style_color struct {
	Name string `xml:"-"`

	// insertion point for fields

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Barline is generated from named complex type "barline"
type Barline struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "location
	Location string `xml:"location,attr,omitempty"`

	// generated from attribute "segno
	Segno string `xml:"segno,attr,omitempty"`

	// generated from attribute "coda
	Coda string `xml:"coda,attr,omitempty"`

	// generated from attribute "divisions
	Divisions string `xml:"divisions,attr,omitempty"`

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from element "bar-style" of type bar-style-color order 81 depth 0
	Bar_style []*Bar_style_color `xml:"bar-style,omitempty"`

	// generated from group with order 82 depth 0
	Group_editorial

	// generated from element "wavy-line" of type wavy-line order 83 depth 0
	Wavy_line []*Wavy_line `xml:"wavy-line,omitempty"`

	// generated from element "segno" of type segno order 84 depth 0
	Segno_1 []*Segno `xml:"segno,omitempty"`

	// generated from element "coda" of type coda order 85 depth 0
	Coda_1 []*Coda `xml:"coda,omitempty"`

	// generated from element "fermata" of type fermata order 86 depth 0
	Fermata []*Fermata `xml:"fermata,omitempty"`

	// generated from element "ending" of type ending order 87 depth 0
	Ending []*Ending `xml:"ending,omitempty"`

	// generated from element "repeat" of type repeat order 88 depth 0
	Repeat []*Repeat `xml:"repeat,omitempty"`
}

// Barre is generated from named complex type "barre"
type Barre struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute group "color
	AttributeGroup_color
}

// Bass is generated from named complex type "bass"
type Bass struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "arrangement
	Arrangement string `xml:"arrangement,attr,omitempty"`

	// generated from element "bass-separator" of type style-text order 93 depth 0
	Bass_separator []*Style_text `xml:"bass-separator,omitempty"`

	// generated from element "bass-step" of type bass-step order 94 depth 0
	Bass_step []*Bass_step `xml:"bass-step,omitempty"`

	// generated from element "bass-alter" of type harmony-alter order 95 depth 0
	Bass_alter []*Harmony_alter `xml:"bass-alter,omitempty"`
}

// Bass_step is generated from named complex type "bass-step"
type Bass_step struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "text
	Text string `xml:"text,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Beam is generated from named complex type "beam"
type Beam struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "number
	Number int `xml:"number,attr,omitempty"`

	// generated from attribute "repeater
	Repeater string `xml:"repeater,attr,omitempty"`

	// generated from attribute "fan
	Fan string `xml:"fan,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Beat_repeat is generated from named complex type "beat-repeat"
type Beat_repeat struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute "slashes
	Slashes int `xml:"slashes,attr,omitempty"`

	// generated from attribute "use-dots
	Use_dots string `xml:"use-dots,attr,omitempty"`

	// generated from group with order 54 depth 0
	Group_slash
}

// Beat_unit_tied is generated from named complex type "beat-unit-tied"
type Beat_unit_tied struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from group with order 96 depth 0
	Group_beat_unit
}

// Beater is generated from named complex type "beater"
type Beater struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "tip
	Tip string `xml:"tip,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Bend is generated from named complex type "bend"
type Bend struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "shape
	Shape string `xml:"shape,attr,omitempty"`

	// generated from attribute group "print-style
	AttributeGroup_print_style

	// generated from attribute group "bend-sound
	AttributeGroup_bend_sound

	// generated from element "bend-alter" of type semitones order 257 depth 0
	Bend_alter string `xml:"bend-alter,omitempty"`

	// generated from element "pre-bend" of type empty order 258 depth 0
	Pre_bend string `xml:"pre-bend,omitempty"`

	// generated from element "release" of type release order 259 depth 0
	Release []*Release `xml:"release,omitempty"`

	// generated from element "with-bar" of type placement-text order 260 depth 0
	With_bar []*Placement_text `xml:"with-bar,omitempty"`
}

// Bookmark is generated from named complex type "bookmark"
type Bookmark struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id
	Id string `xml:"id,attr,omitempty"`

	// generated from attribute "name
	NameXSD string `xml:"name,attr,omitempty"`

	// generated from attribute group "element-position
	AttributeGroup_element_position
}

// Bracket is generated from named complex type "bracket"
type Bracket struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute "number
	Number int `xml:"number,attr,omitempty"`

	// generated from attribute "line-end
	Line_end string `xml:"line-end,attr,omitempty"`

	// generated from attribute "end-length
	End_length string `xml:"end-length,attr,omitempty"`

	// generated from attribute group "line-type
	AttributeGroup_line_type

	// generated from attribute group "dashed-formatting
	AttributeGroup_dashed_formatting

	// generated from attribute group "position
	AttributeGroup_position

	// generated from attribute group "color
	AttributeGroup_color

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id
}

// Breath_mark is generated from named complex type "breath-mark"
type Breath_mark struct {
	Name string `xml:"-"`

	// insertion point for fields

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Caesura is generated from named complex type "caesura"
type Caesura struct {
	Name string `xml:"-"`

	// insertion point for fields

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Cancel is generated from named complex type "cancel"
type Cancel struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "location
	Location string `xml:"location,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText int `xml:",chardata"`
}

// Clef is generated from named complex type "clef"
type Clef struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "number
	Number int `xml:"number,attr,omitempty"`

	// generated from attribute "additional
	Additional string `xml:"additional,attr,omitempty"`

	// generated from attribute "size
	Size string `xml:"size,attr,omitempty"`

	// generated from attribute "after-barline
	After_barline string `xml:"after-barline,attr,omitempty"`

	// generated from attribute group "print-style
	AttributeGroup_print_style

	// generated from attribute group "print-object
	AttributeGroup_print_object

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from group with order 55 depth 0
	Group_clef
}

// Coda is generated from named complex type "coda"
type Coda struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "smufl
	Smufl string `xml:"smufl,attr,omitempty"`

	// generated from attribute group "print-style-align
	AttributeGroup_print_style_align

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id
}

// Credit is generated from named complex type "credit"
type Credit struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "page
	Page int `xml:"page,attr,omitempty"`

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from element "credit-type" of type xs:string order 400 depth 0
	Credit_type string `xml:"credit-type,omitempty"`

	// generated from element "credit-image" of type image order 403 depth 0
	Credit_image []*Image `xml:"credit-image,omitempty"`

	// generated from element "link" of type link order 406 depth 0
	Link []*Link `xml:"link,omitempty"`

	// generated from element "bookmark" of type bookmark order 407 depth 0
	Bookmark []*Bookmark `xml:"bookmark,omitempty"`

	// generated from element "credit-words" of type formatted-text-id order 408 depth 0
	Credit_words []*Formatted_text_id `xml:"credit-words,omitempty"`

	// generated from element "credit-symbol" of type formatted-symbol-id order 409 depth 0
	Credit_symbol []*Formatted_symbol_id `xml:"credit-symbol,omitempty"`
}

// Dashes is generated from named complex type "dashes"
type Dashes struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute "number
	Number int `xml:"number,attr,omitempty"`

	// generated from attribute group "dashed-formatting
	AttributeGroup_dashed_formatting

	// generated from attribute group "position
	AttributeGroup_position

	// generated from attribute group "color
	AttributeGroup_color

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id
}

// Defaults is generated from named complex type "defaults"
type Defaults struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "scaling" of type scaling order 410 depth 0
	Scaling []*Scaling `xml:"scaling,omitempty"`

	// generated from element "concert-score" of type empty order 411 depth 0
	Concert_score string `xml:"concert-score,omitempty"`

	// generated from group with order 412 depth 0
	Group_layout

	// generated from element "appearance" of type appearance order 413 depth 0
	Appearance []*Appearance `xml:"appearance,omitempty"`

	// generated from element "music-font" of type empty-font order 414 depth 0
	Music_font []*Empty_font `xml:"music-font,omitempty"`

	// generated from element "word-font" of type empty-font order 415 depth 0
	Word_font []*Empty_font `xml:"word-font,omitempty"`

	// generated from element "lyric-font" of type lyric-font order 416 depth 0
	Lyric_font []*Lyric_font `xml:"lyric-font,omitempty"`

	// generated from element "lyric-language" of type lyric-language order 417 depth 0
	Lyric_language []*Lyric_language `xml:"lyric-language,omitempty"`
}

// Degree is generated from named complex type "degree"
type Degree struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "print-object
	AttributeGroup_print_object

	// generated from element "degree-value" of type degree-value order 97 depth 0
	Degree_value []*Degree_value `xml:"degree-value,omitempty"`

	// generated from element "degree-alter" of type degree-alter order 98 depth 0
	Degree_alter []*Degree_alter `xml:"degree-alter,omitempty"`

	// generated from element "degree-type" of type degree-type order 99 depth 0
	Degree_type []*Degree_type `xml:"degree-type,omitempty"`
}

// Degree_alter is generated from named complex type "degree-alter"
type Degree_alter struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "plus-minus
	Plus_minus string `xml:"plus-minus,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Degree_type is generated from named complex type "degree-type"
type Degree_type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "text
	Text string `xml:"text,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Degree_value is generated from named complex type "degree-value"
type Degree_value struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "symbol
	Symbol string `xml:"symbol,attr,omitempty"`

	// generated from attribute "text
	Text string `xml:"text,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText int `xml:",chardata"`
}

// Direction is generated from named complex type "direction"
type Direction struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "placement
	AttributeGroup_placement

	// generated from attribute group "directive
	AttributeGroup_directive

	// generated from attribute group "system-relation
	AttributeGroup_system_relation

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from element "direction-type" of type direction-type order 100 depth 0
	Direction_type []*Direction_type `xml:"direction-type,omitempty"`

	// generated from element "offset" of type offset order 101 depth 0
	Offset []*Offset `xml:"offset,omitempty"`

	// generated from group with order 102 depth 0
	Group_editorial_voice_direction

	// generated from group with order 103 depth 0
	Group_staff

	// generated from element "sound" of type sound order 104 depth 0
	Sound []*Sound `xml:"sound,omitempty"`

	// generated from element "listening" of type listening order 105 depth 0
	Listening []*Listening `xml:"listening,omitempty"`
}

// Direction_type is generated from named complex type "direction-type"
type Direction_type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from element "rehearsal" of type formatted-text-id order 106 depth 0
	Rehearsal []*Formatted_text_id `xml:"rehearsal,omitempty"`

	// generated from element "segno" of type segno order 107 depth 0
	Segno []*Segno `xml:"segno,omitempty"`

	// generated from element "coda" of type coda order 108 depth 0
	Coda []*Coda `xml:"coda,omitempty"`

	// generated from element "words" of type formatted-text-id order 109 depth 0
	Words []*Formatted_text_id `xml:"words,omitempty"`

	// generated from element "symbol" of type formatted-symbol-id order 110 depth 0
	Symbol []*Formatted_symbol_id `xml:"symbol,omitempty"`

	// generated from element "wedge" of type wedge order 111 depth 0
	Wedge []*Wedge `xml:"wedge,omitempty"`

	// generated from element "dynamics" of type dynamics order 112 depth 0
	Dynamics []*Dynamics `xml:"dynamics,omitempty"`

	// generated from element "dashes" of type dashes order 113 depth 0
	Dashes []*Dashes `xml:"dashes,omitempty"`

	// generated from element "bracket" of type bracket order 114 depth 0
	Bracket []*Bracket `xml:"bracket,omitempty"`

	// generated from element "pedal" of type pedal order 115 depth 0
	Pedal []*Pedal `xml:"pedal,omitempty"`

	// generated from element "metronome" of type metronome order 116 depth 0
	Metronome []*Metronome `xml:"metronome,omitempty"`

	// generated from element "octave-shift" of type octave-shift order 117 depth 0
	Octave_shift []*Octave_shift `xml:"octave-shift,omitempty"`

	// generated from element "harp-pedals" of type harp-pedals order 118 depth 0
	Harp_pedals []*Harp_pedals `xml:"harp-pedals,omitempty"`

	// generated from element "damp" of type empty-print-style-align-id order 119 depth 0
	Damp []*Empty_print_style_align_id `xml:"damp,omitempty"`

	// generated from element "damp-all" of type empty-print-style-align-id order 120 depth 0
	Damp_all []*Empty_print_style_align_id `xml:"damp-all,omitempty"`

	// generated from element "eyeglasses" of type empty-print-style-align-id order 121 depth 0
	Eyeglasses []*Empty_print_style_align_id `xml:"eyeglasses,omitempty"`

	// generated from element "string-mute" of type string-mute order 122 depth 0
	String_mute []*String_mute `xml:"string-mute,omitempty"`

	// generated from element "scordatura" of type scordatura order 123 depth 0
	Scordatura []*Scordatura `xml:"scordatura,omitempty"`

	// generated from element "image" of type image order 124 depth 0
	Image []*Image `xml:"image,omitempty"`

	// generated from element "principal-voice" of type principal-voice order 125 depth 0
	Principal_voice []*Principal_voice `xml:"principal-voice,omitempty"`

	// generated from element "percussion" of type percussion order 126 depth 0
	Percussion []*Percussion `xml:"percussion,omitempty"`

	// generated from element "accordion-registration" of type accordion-registration order 127 depth 0
	Accordion_registration []*Accordion_registration `xml:"accordion-registration,omitempty"`

	// generated from element "staff-divide" of type staff-divide order 128 depth 0
	Staff_divide []*Staff_divide `xml:"staff-divide,omitempty"`

	// generated from element "other-direction" of type other-direction order 129 depth 0
	Other_direction []*Other_direction `xml:"other-direction,omitempty"`
}

// Distance is generated from named complex type "distance"
type Distance struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type string `xml:"type,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Double is generated from named complex type "double"
type Double struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "above
	Above string `xml:"above,attr,omitempty"`
}

// Dynamics is generated from named complex type "dynamics"
type Dynamics struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "print-style-align
	AttributeGroup_print_style_align

	// generated from attribute group "placement
	AttributeGroup_placement

	// generated from attribute group "text-decoration
	AttributeGroup_text_decoration

	// generated from attribute group "enclosure
	AttributeGroup_enclosure

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from element "p" of type empty order 0 depth 0
	P string `xml:"p,omitempty"`

	// generated from element "pp" of type empty order 1 depth 0
	Pp string `xml:"pp,omitempty"`

	// generated from element "ppp" of type empty order 2 depth 0
	Ppp string `xml:"ppp,omitempty"`

	// generated from element "pppp" of type empty order 3 depth 0
	Pppp string `xml:"pppp,omitempty"`

	// generated from element "ppppp" of type empty order 4 depth 0
	Ppppp string `xml:"ppppp,omitempty"`

	// generated from element "pppppp" of type empty order 5 depth 0
	Pppppp string `xml:"pppppp,omitempty"`

	// generated from element "f" of type empty order 6 depth 0
	F string `xml:"f,omitempty"`

	// generated from element "ff" of type empty order 7 depth 0
	Ff string `xml:"ff,omitempty"`

	// generated from element "fff" of type empty order 8 depth 0
	Fff string `xml:"fff,omitempty"`

	// generated from element "ffff" of type empty order 9 depth 0
	Ffff string `xml:"ffff,omitempty"`

	// generated from element "fffff" of type empty order 10 depth 0
	Fffff string `xml:"fffff,omitempty"`

	// generated from element "ffffff" of type empty order 11 depth 0
	Ffffff string `xml:"ffffff,omitempty"`

	// generated from element "mp" of type empty order 12 depth 0
	Mp string `xml:"mp,omitempty"`

	// generated from element "mf" of type empty order 13 depth 0
	Mf string `xml:"mf,omitempty"`

	// generated from element "sf" of type empty order 14 depth 0
	Sf string `xml:"sf,omitempty"`

	// generated from element "sfp" of type empty order 15 depth 0
	Sfp string `xml:"sfp,omitempty"`

	// generated from element "sfpp" of type empty order 16 depth 0
	Sfpp string `xml:"sfpp,omitempty"`

	// generated from element "fp" of type empty order 17 depth 0
	Fp string `xml:"fp,omitempty"`

	// generated from element "rf" of type empty order 18 depth 0
	Rf string `xml:"rf,omitempty"`

	// generated from element "rfz" of type empty order 19 depth 0
	Rfz string `xml:"rfz,omitempty"`

	// generated from element "sfz" of type empty order 20 depth 0
	Sfz string `xml:"sfz,omitempty"`

	// generated from element "sffz" of type empty order 21 depth 0
	Sffz string `xml:"sffz,omitempty"`

	// generated from element "fz" of type empty order 22 depth 0
	Fz string `xml:"fz,omitempty"`

	// generated from element "n" of type empty order 23 depth 0
	N string `xml:"n,omitempty"`

	// generated from element "pf" of type empty order 24 depth 0
	Pf string `xml:"pf,omitempty"`

	// generated from element "sfzp" of type empty order 25 depth 0
	Sfzp string `xml:"sfzp,omitempty"`

	// generated from element "other-dynamics" of type other-text order 26 depth 0
	Other_dynamics []*Other_text `xml:"other-dynamics,omitempty"`
}

// Effect is generated from named complex type "effect"
type Effect struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "smufl
	Smufl string `xml:"smufl,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Elision is generated from named complex type "elision"
type Elision struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "smufl
	Smufl string `xml:"smufl,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Empty is generated from named complex type "empty"
type Empty struct {
	Name string `xml:"-"`

	// insertion point for fields
}

// Empty_font is generated from named complex type "empty-font"
type Empty_font struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "font
	AttributeGroup_font
}

// Empty_line is generated from named complex type "empty-line"
type Empty_line struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "line-shape
	AttributeGroup_line_shape

	// generated from attribute group "line-type
	AttributeGroup_line_type

	// generated from attribute group "line-length
	AttributeGroup_line_length

	// generated from attribute group "dashed-formatting
	AttributeGroup_dashed_formatting

	// generated from attribute group "print-style
	AttributeGroup_print_style

	// generated from attribute group "placement
	AttributeGroup_placement
}

// Empty_placement is generated from named complex type "empty-placement"
type Empty_placement struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "print-style
	AttributeGroup_print_style

	// generated from attribute group "placement
	AttributeGroup_placement
}

// Empty_placement_smufl is generated from named complex type "empty-placement-smufl"
type Empty_placement_smufl struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "print-style
	AttributeGroup_print_style

	// generated from attribute group "placement
	AttributeGroup_placement

	// generated from attribute group "smufl
	AttributeGroup_smufl
}

// Empty_print_object_style_align is generated from named complex type "empty-print-object-style-align"
type Empty_print_object_style_align struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "print-object
	AttributeGroup_print_object

	// generated from attribute group "print-style-align
	AttributeGroup_print_style_align
}

// Empty_print_style is generated from named complex type "empty-print-style"
type Empty_print_style struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "print-style
	AttributeGroup_print_style
}

// Empty_print_style_align is generated from named complex type "empty-print-style-align"
type Empty_print_style_align struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "print-style-align
	AttributeGroup_print_style_align
}

// Empty_print_style_align_id is generated from named complex type "empty-print-style-align-id"
type Empty_print_style_align_id struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "print-style-align
	AttributeGroup_print_style_align

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id
}

// Empty_trill_sound is generated from named complex type "empty-trill-sound"
type Empty_trill_sound struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "print-style
	AttributeGroup_print_style

	// generated from attribute group "placement
	AttributeGroup_placement

	// generated from attribute group "trill-sound
	AttributeGroup_trill_sound
}

// Encoding is generated from named complex type "encoding"
type Encoding struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "encoder" of type typed-text order 203 depth 0
	Encoder []*Typed_text `xml:"encoder,omitempty"`

	// generated from element "software" of type xs:string order 204 depth 0
	Software string `xml:"software,omitempty"`

	// generated from element "encoding-description" of type xs:string order 205 depth 0
	Encoding_description string `xml:"encoding-description,omitempty"`

	// generated from element "supports" of type supports order 206 depth 0
	Supports []*Supports `xml:"supports,omitempty"`
}

// Ending is generated from named complex type "ending"
type Ending struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "number
	Number string `xml:"number,attr,omitempty"`

	// generated from attribute "type
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute "end-length
	End_length string `xml:"end-length,attr,omitempty"`

	// generated from attribute "text-x
	Text_x string `xml:"text-x,attr,omitempty"`

	// generated from attribute "text-y
	Text_y string `xml:"text-y,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Extend is generated from named complex type "extend"
type Extend struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute group "position
	AttributeGroup_position

	// generated from attribute group "color
	AttributeGroup_color
}

// Feature is generated from named complex type "feature"
type Feature struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type string `xml:"type,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Fermata is generated from named complex type "fermata"
type Fermata struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type string `xml:"type,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Figure is generated from named complex type "figure"
type Figure struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "prefix" of type style-text order 261 depth 0
	Prefix []*Style_text `xml:"prefix,omitempty"`

	// generated from element "figure-number" of type style-text order 262 depth 0
	Figure_number []*Style_text `xml:"figure-number,omitempty"`

	// generated from element "suffix" of type style-text order 263 depth 0
	Suffix []*Style_text `xml:"suffix,omitempty"`

	// generated from element "extend" of type extend order 264 depth 0
	Extend []*Extend `xml:"extend,omitempty"`

	// generated from group with order 265 depth 0
	Group_editorial
}

// Figured_bass is generated from named complex type "figured-bass"
type Figured_bass struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "parentheses
	Parentheses string `xml:"parentheses,attr,omitempty"`

	// generated from attribute group "print-style-align
	AttributeGroup_print_style_align

	// generated from attribute group "placement
	AttributeGroup_placement

	// generated from attribute group "printout
	AttributeGroup_printout

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from element "figure" of type figure order 266 depth 0
	Figure []*Figure `xml:"figure,omitempty"`

	// generated from group with order 267 depth 0
	Group_duration

	// generated from group with order 268 depth 0
	Group_editorial
}

// Fingering is generated from named complex type "fingering"
type Fingering struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "substitution
	Substitution string `xml:"substitution,attr,omitempty"`

	// generated from attribute "alternate
	Alternate string `xml:"alternate,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// First_fret is generated from named complex type "first-fret"
type First_fret struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "text
	Text string `xml:"text,attr,omitempty"`

	// generated from attribute "location
	Location string `xml:"location,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText int `xml:",chardata"`
}

// For_part is generated from named complex type "for-part"
type For_part struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "number
	Number int `xml:"number,attr,omitempty"`

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from element "part-clef" of type part-clef order 56 depth 0
	Part_clef []*Part_clef `xml:"part-clef,omitempty"`

	// generated from element "part-transpose" of type part-transpose order 57 depth 0
	Part_transpose []*Part_transpose `xml:"part-transpose,omitempty"`
}

// Formatted_symbol is generated from named complex type "formatted-symbol"
type Formatted_symbol struct {
	Name string `xml:"-"`

	// insertion point for fields

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Formatted_symbol_id is generated from named complex type "formatted-symbol-id"
type Formatted_symbol_id struct {
	Name string `xml:"-"`

	// insertion point for fields

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Formatted_text is generated from named complex type "formatted-text"
type Formatted_text struct {
	Name string `xml:"-"`

	// insertion point for fields

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Formatted_text_id is generated from named complex type "formatted-text-id"
type Formatted_text_id struct {
	Name string `xml:"-"`

	// insertion point for fields

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Forward is generated from named complex type "forward"
type Forward struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from group with order 269 depth 0
	Group_duration

	// generated from group with order 270 depth 0
	Group_editorial_voice

	// generated from group with order 271 depth 0
	Group_staff
}

// Frame is generated from named complex type "frame"
type Frame struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "height
	Height string `xml:"height,attr,omitempty"`

	// generated from attribute "width
	Width string `xml:"width,attr,omitempty"`

	// generated from attribute "unplayed
	Unplayed string `xml:"unplayed,attr,omitempty"`

	// generated from attribute group "position
	AttributeGroup_position

	// generated from attribute group "color
	AttributeGroup_color

	// generated from attribute group "halign
	AttributeGroup_halign

	// generated from attribute group "valign-image
	AttributeGroup_valign_image

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from element "frame-strings" of type xs:positiveInteger order 130 depth 0
	Frame_strings int `xml:"frame-strings,omitempty"`

	// generated from element "frame-frets" of type xs:positiveInteger order 131 depth 0
	Frame_frets int `xml:"frame-frets,omitempty"`

	// generated from element "first-fret" of type first-fret order 132 depth 0
	First_fret []*First_fret `xml:"first-fret,omitempty"`

	// generated from element "frame-note" of type frame-note order 133 depth 0
	Frame_note []*Frame_note `xml:"frame-note,omitempty"`
}

// Frame_note is generated from named complex type "frame-note"
type Frame_note struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "string" of type string-type order 134 depth 0
	String []*String_type `xml:"string,omitempty"`

	// generated from element "fret" of type fret order 135 depth 0
	Fret []*Fret `xml:"fret,omitempty"`

	// generated from element "fingering" of type fingering order 136 depth 0
	Fingering []*Fingering `xml:"fingering,omitempty"`

	// generated from element "barre" of type barre order 137 depth 0
	Barre []*Barre `xml:"barre,omitempty"`
}

// Fret is generated from named complex type "fret"
type Fret struct {
	Name string `xml:"-"`

	// insertion point for fields

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText int `xml:",chardata"`
}

// Glass is generated from named complex type "glass"
type Glass struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "smufl
	Smufl string `xml:"smufl,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Glissando is generated from named complex type "glissando"
type Glissando struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute "number
	Number int `xml:"number,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Glyph is generated from named complex type "glyph"
type Glyph struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type string `xml:"type,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Grace is generated from named complex type "grace"
type Grace struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "steal-time-previous
	Steal_time_previous string `xml:"steal-time-previous,attr,omitempty"`

	// generated from attribute "steal-time-following
	Steal_time_following string `xml:"steal-time-following,attr,omitempty"`

	// generated from attribute "make-time
	Make_time string `xml:"make-time,attr,omitempty"`

	// generated from attribute "slash
	Slash string `xml:"slash,attr,omitempty"`
}

// Group_barline is generated from named complex type "group-barline"
type Group_barline struct {
	Name string `xml:"-"`

	// insertion point for fields

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Group_name is generated from named complex type "group-name"
type Group_name struct {
	Name string `xml:"-"`

	// insertion point for fields

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Group_symbol is generated from named complex type "group-symbol"
type Group_symbol struct {
	Name string `xml:"-"`

	// insertion point for fields

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Grouping is generated from named complex type "grouping"
type Grouping struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute "number
	Number string `xml:"number,attr,omitempty"`

	// generated from attribute "member-of
	Member_of string `xml:"member-of,attr,omitempty"`

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from element "feature" of type feature order 138 depth 0
	Feature []*Feature `xml:"feature,omitempty"`
}

// Hammer_on_pull_off is generated from named complex type "hammer-on-pull-off"
type Hammer_on_pull_off struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute "number
	Number int `xml:"number,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Handbell is generated from named complex type "handbell"
type Handbell struct {
	Name string `xml:"-"`

	// insertion point for fields

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Harmon_closed is generated from named complex type "harmon-closed"
type Harmon_closed struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "location
	Location string `xml:"location,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Harmon_mute is generated from named complex type "harmon-mute"
type Harmon_mute struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "print-style
	AttributeGroup_print_style

	// generated from attribute group "placement
	AttributeGroup_placement

	// generated from element "harmon-closed" of type harmon-closed order 272 depth 0
	Harmon_closed []*Harmon_closed `xml:"harmon-closed,omitempty"`
}

// Harmonic is generated from named complex type "harmonic"
type Harmonic struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "print-object
	AttributeGroup_print_object

	// generated from attribute group "print-style
	AttributeGroup_print_style

	// generated from attribute group "placement
	AttributeGroup_placement

	// generated from element "natural" of type empty order 273 depth 0
	Natural string `xml:"natural,omitempty"`

	// generated from element "artificial" of type empty order 274 depth 0
	Artificial string `xml:"artificial,omitempty"`

	// generated from element "base-pitch" of type empty order 275 depth 0
	Base_pitch string `xml:"base-pitch,omitempty"`

	// generated from element "touching-pitch" of type empty order 276 depth 0
	Touching_pitch string `xml:"touching-pitch,omitempty"`

	// generated from element "sounding-pitch" of type empty order 277 depth 0
	Sounding_pitch string `xml:"sounding-pitch,omitempty"`
}

// Harmony is generated from named complex type "harmony"
type Harmony struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute "print-frame
	Print_frame string `xml:"print-frame,attr,omitempty"`

	// generated from attribute "arrangement
	Arrangement string `xml:"arrangement,attr,omitempty"`

	// generated from attribute group "print-object
	AttributeGroup_print_object

	// generated from attribute group "print-style
	AttributeGroup_print_style

	// generated from attribute group "placement
	AttributeGroup_placement

	// generated from attribute group "system-relation
	AttributeGroup_system_relation

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from group with order 139 depth 0
	Group_harmony_chord

	// generated from element "frame" of type frame order 140 depth 0
	Frame []*Frame `xml:"frame,omitempty"`

	// generated from element "offset" of type offset order 141 depth 0
	Offset []*Offset `xml:"offset,omitempty"`

	// generated from group with order 142 depth 0
	Group_editorial

	// generated from group with order 143 depth 0
	Group_staff
}

// Harmony_alter is generated from named complex type "harmony-alter"
type Harmony_alter struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "location
	Location string `xml:"location,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Harp_pedals is generated from named complex type "harp-pedals"
type Harp_pedals struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "print-style-align
	AttributeGroup_print_style_align

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from element "pedal-tuning" of type pedal-tuning order 144 depth 0
	Pedal_tuning []*Pedal_tuning `xml:"pedal-tuning,omitempty"`
}

// Heel_toe is generated from named complex type "heel-toe"
type Heel_toe struct {
	Name string `xml:"-"`

	// insertion point for fields
}

// Hole is generated from named complex type "hole"
type Hole struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "print-style
	AttributeGroup_print_style

	// generated from attribute group "placement
	AttributeGroup_placement

	// generated from element "hole-type" of type xs:string order 278 depth 0
	Hole_type string `xml:"hole-type,omitempty"`

	// generated from element "hole-closed" of type hole-closed order 279 depth 0
	Hole_closed []*Hole_closed `xml:"hole-closed,omitempty"`

	// generated from element "hole-shape" of type xs:string order 280 depth 0
	Hole_shape string `xml:"hole-shape,omitempty"`
}

// Hole_closed is generated from named complex type "hole-closed"
type Hole_closed struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "location
	Location string `xml:"location,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Horizontal_turn is generated from named complex type "horizontal-turn"
type Horizontal_turn struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "slash
	Slash string `xml:"slash,attr,omitempty"`

	// generated from attribute group "print-style
	AttributeGroup_print_style

	// generated from attribute group "placement
	AttributeGroup_placement

	// generated from attribute group "trill-sound
	AttributeGroup_trill_sound
}

// Identification is generated from named complex type "identification"
type Identification struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "creator" of type typed-text order 207 depth 0
	Creator []*Typed_text `xml:"creator,omitempty"`

	// generated from element "rights" of type typed-text order 208 depth 0
	Rights []*Typed_text `xml:"rights,omitempty"`

	// generated from element "encoding" of type encoding order 209 depth 0
	Encoding []*Encoding `xml:"encoding,omitempty"`

	// generated from element "source" of type xs:string order 210 depth 0
	Source string `xml:"source,omitempty"`

	// generated from element "relation" of type typed-text order 211 depth 0
	Relation []*Typed_text `xml:"relation,omitempty"`

	// generated from element "miscellaneous" of type miscellaneous order 212 depth 0
	Miscellaneous []*Miscellaneous `xml:"miscellaneous,omitempty"`
}

// Image is generated from named complex type "image"
type Image struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "image-attributes
	AttributeGroup_image_attributes

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id
}

// Instrument is generated from named complex type "instrument"
type Instrument struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id
	Id string `xml:"id,attr,omitempty"`
}

// Instrument_change is generated from named complex type "instrument-change"
type Instrument_change struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id
	Id string `xml:"id,attr,omitempty"`

	// generated from group with order 145 depth 0
	Group_virtual_instrument_data
}

// Instrument_link is generated from named complex type "instrument-link"
type Instrument_link struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id
	Id string `xml:"id,attr,omitempty"`
}

// Interchangeable is generated from named complex type "interchangeable"
type Interchangeable struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "symbol
	Symbol string `xml:"symbol,attr,omitempty"`

	// generated from attribute "separator
	Separator string `xml:"separator,attr,omitempty"`

	// generated from element "time-relation" of type time-relation order 58 depth 0
	Time_relation string `xml:"time-relation,omitempty"`

	// generated from group with order 59 depth 0
	Group_time_signature
}

// Inversion is generated from named complex type "inversion"
type Inversion struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "text
	Text string `xml:"text,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText int `xml:",chardata"`
}

// Key is generated from named complex type "key"
type Key struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "number
	Number int `xml:"number,attr,omitempty"`

	// generated from attribute group "print-style
	AttributeGroup_print_style

	// generated from attribute group "print-object
	AttributeGroup_print_object

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from group with order 60 depth 0
	Group_traditional_key

	// generated from group with order 61 depth 0
	Group_non_traditional_key

	// generated from element "key-octave" of type key-octave order 62 depth 0
	Key_octave []*Key_octave `xml:"key-octave,omitempty"`
}

// Key_accidental is generated from named complex type "key-accidental"
type Key_accidental struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "smufl
	Smufl string `xml:"smufl,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Key_octave is generated from named complex type "key-octave"
type Key_octave struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "number
	Number int `xml:"number,attr,omitempty"`

	// generated from attribute "cancel
	Cancel string `xml:"cancel,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText int `xml:",chardata"`
}

// Kind is generated from named complex type "kind"
type Kind struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "use-symbols
	Use_symbols string `xml:"use-symbols,attr,omitempty"`

	// generated from attribute "text
	Text string `xml:"text,attr,omitempty"`

	// generated from attribute "stack-degrees
	Stack_degrees string `xml:"stack-degrees,attr,omitempty"`

	// generated from attribute "parentheses-degrees
	Parentheses_degrees string `xml:"parentheses-degrees,attr,omitempty"`

	// generated from attribute "bracket-degrees
	Bracket_degrees string `xml:"bracket-degrees,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Level is generated from named complex type "level"
type Level struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "reference
	Reference string `xml:"reference,attr,omitempty"`

	// generated from attribute "type
	Type string `xml:"type,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Line_detail is generated from named complex type "line-detail"
type Line_detail struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "line
	Line int `xml:"line,attr,omitempty"`

	// generated from attribute "width
	Width string `xml:"width,attr,omitempty"`

	// generated from attribute group "color
	AttributeGroup_color

	// generated from attribute group "line-type
	AttributeGroup_line_type

	// generated from attribute group "print-object
	AttributeGroup_print_object
}

// Line_width is generated from named complex type "line-width"
type Line_width struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type string `xml:"type,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Link is generated from named complex type "link"
type Link struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "name
	NameXSD string `xml:"name,attr,omitempty"`

	// generated from attribute group "link-attributes
	AttributeGroup_link_attributes

	// generated from attribute group "element-position
	AttributeGroup_element_position

	// generated from attribute group "position
	AttributeGroup_position
}

// Listen is generated from named complex type "listen"
type Listen struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "assess" of type assess order 281 depth 0
	Assess []*Assess `xml:"assess,omitempty"`

	// generated from element "wait" of type wait order 282 depth 0
	Wait []*Wait `xml:"wait,omitempty"`

	// generated from element "other-listen" of type other-listening order 283 depth 0
	Other_listen []*Other_listening `xml:"other-listen,omitempty"`
}

// Listening is generated from named complex type "listening"
type Listening struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "sync" of type sync order 146 depth 0
	Sync []*Sync `xml:"sync,omitempty"`

	// generated from element "other-listening" of type other-listening order 147 depth 0
	Other_listening []*Other_listening `xml:"other-listening,omitempty"`

	// generated from element "offset" of type offset order 148 depth 0
	Offset []*Offset `xml:"offset,omitempty"`
}

// Lyric is generated from named complex type "lyric"
type Lyric struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "number
	Number string `xml:"number,attr,omitempty"`

	// generated from attribute "name
	NameXSD string `xml:"name,attr,omitempty"`

	// generated from attribute "time-only
	Time_only string `xml:"time-only,attr,omitempty"`

	// generated from attribute group "justify
	AttributeGroup_justify

	// generated from attribute group "position
	AttributeGroup_position

	// generated from attribute group "placement
	AttributeGroup_placement

	// generated from attribute group "color
	AttributeGroup_color

	// generated from attribute group "print-object
	AttributeGroup_print_object

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from element "elision" of type elision order 286 depth 0
	Elision []*Elision `xml:"elision,omitempty"`

	// generated from element "syllabic" of type syllabic order 287 depth 0
	Syllabic string `xml:"syllabic,omitempty"`

	// generated from element "text" of type text-element-data order 288 depth 0
	Text []*Text_element_data `xml:"text,omitempty"`

	// generated from element "extend" of type extend order 289 depth 0
	Extend []*Extend `xml:"extend,omitempty"`

	// generated from element "laughing" of type empty order 291 depth 0
	Laughing string `xml:"laughing,omitempty"`

	// generated from element "humming" of type empty order 292 depth 0
	Humming string `xml:"humming,omitempty"`

	// generated from element "end-line" of type empty order 293 depth 0
	End_line string `xml:"end-line,omitempty"`

	// generated from element "end-paragraph" of type empty order 294 depth 0
	End_paragraph string `xml:"end-paragraph,omitempty"`

	// generated from group with order 295 depth 0
	Group_editorial
}

// Lyric_font is generated from named complex type "lyric-font"
type Lyric_font struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "number
	Number string `xml:"number,attr,omitempty"`

	// generated from attribute "name
	NameXSD string `xml:"name,attr,omitempty"`

	// generated from attribute group "font
	AttributeGroup_font
}

// Lyric_language is generated from named complex type "lyric-language"
type Lyric_language struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "number
	Number string `xml:"number,attr,omitempty"`

	// generated from attribute "name
	NameXSD string `xml:"name,attr,omitempty"`

	// generated from attribute "http://www.w3.org/XML/1998/namespace lang
	Lang string `xml:"http://www.w3.org/XML/1998/namespace lang,attr,omitempty"`
}

// Measure_layout is generated from named complex type "measure-layout"
type Measure_layout struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "measure-distance" of type tenths order 219 depth 0
	Measure_distance string `xml:"measure-distance,omitempty"`
}

// Measure_numbering is generated from named complex type "measure-numbering"
type Measure_numbering struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "system
	System string `xml:"system,attr,omitempty"`

	// generated from attribute "staff
	Staff int `xml:"staff,attr,omitempty"`

	// generated from attribute "multiple-rest-always
	Multiple_rest_always string `xml:"multiple-rest-always,attr,omitempty"`

	// generated from attribute "multiple-rest-range
	Multiple_rest_range string `xml:"multiple-rest-range,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Measure_repeat is generated from named complex type "measure-repeat"
type Measure_repeat struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute "slashes
	Slashes int `xml:"slashes,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Measure_style is generated from named complex type "measure-style"
type Measure_style struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "number
	Number int `xml:"number,attr,omitempty"`

	// generated from attribute group "font
	AttributeGroup_font

	// generated from attribute group "color
	AttributeGroup_color

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from element "multiple-rest" of type multiple-rest order 63 depth 0
	Multiple_rest []*Multiple_rest `xml:"multiple-rest,omitempty"`

	// generated from element "measure-repeat" of type measure-repeat order 64 depth 0
	Measure_repeat []*Measure_repeat `xml:"measure-repeat,omitempty"`

	// generated from element "beat-repeat" of type beat-repeat order 65 depth 0
	Beat_repeat []*Beat_repeat `xml:"beat-repeat,omitempty"`

	// generated from element "slash" of type slash order 66 depth 0
	Slash []*Slash `xml:"slash,omitempty"`
}

// Membrane is generated from named complex type "membrane"
type Membrane struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "smufl
	Smufl string `xml:"smufl,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Metal is generated from named complex type "metal"
type Metal struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "smufl
	Smufl string `xml:"smufl,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Metronome is generated from named complex type "metronome"
type Metronome struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "parentheses
	Parentheses string `xml:"parentheses,attr,omitempty"`

	// generated from attribute group "print-style-align
	AttributeGroup_print_style_align

	// generated from attribute group "print-object
	AttributeGroup_print_object

	// generated from attribute group "justify
	AttributeGroup_justify

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from group with order 149 depth 0
	Group_beat_unit

	// generated from element "per-minute" of type per-minute order 151 depth 0
	Per_minute []*Per_minute `xml:"per-minute,omitempty"`

	// generated from element "beat-unit-tied" of type beat-unit-tied order 153 depth 0
	Beat_unit_tied []*Beat_unit_tied `xml:"beat-unit-tied,omitempty"`

	// generated from element "metronome-arrows" of type empty order 154 depth 0
	Metronome_arrows string `xml:"metronome-arrows,omitempty"`

	// generated from element "metronome-relation" of type xs:string order 156 depth 0
	Metronome_relation string `xml:"metronome-relation,omitempty"`

	// generated from element "metronome-note" of type metronome-note order 157 depth 0
	Metronome_note []*Metronome_note `xml:"metronome-note,omitempty"`
}

// Metronome_beam is generated from named complex type "metronome-beam"
type Metronome_beam struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "number
	Number int `xml:"number,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Metronome_note is generated from named complex type "metronome-note"
type Metronome_note struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "metronome-type" of type note-type-value order 158 depth 0
	Metronome_type string `xml:"metronome-type,omitempty"`

	// generated from element "metronome-dot" of type empty order 159 depth 0
	Metronome_dot string `xml:"metronome-dot,omitempty"`

	// generated from element "metronome-beam" of type metronome-beam order 160 depth 0
	Metronome_beam []*Metronome_beam `xml:"metronome-beam,omitempty"`

	// generated from element "metronome-tied" of type metronome-tied order 161 depth 0
	Metronome_tied []*Metronome_tied `xml:"metronome-tied,omitempty"`

	// generated from element "metronome-tuplet" of type metronome-tuplet order 162 depth 0
	Metronome_tuplet []*Metronome_tuplet `xml:"metronome-tuplet,omitempty"`
}

// Metronome_tied is generated from named complex type "metronome-tied"
type Metronome_tied struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type string `xml:"type,attr,omitempty"`
}

// Metronome_tuplet is generated from named complex type "metronome-tuplet"
type Metronome_tuplet struct {
	Name string `xml:"-"`

	// insertion point for fields
}

// Midi_device is generated from named complex type "midi-device"
type Midi_device struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "port
	Port int `xml:"port,attr,omitempty"`

	// generated from attribute "id
	Id string `xml:"id,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Midi_instrument is generated from named complex type "midi-instrument"
type Midi_instrument struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id
	Id string `xml:"id,attr,omitempty"`

	// generated from element "midi-channel" of type midi-16 order 27 depth 0
	Midi_channel int `xml:"midi-channel,omitempty"`

	// generated from element "midi-name" of type xs:string order 28 depth 0
	Midi_name string `xml:"midi-name,omitempty"`

	// generated from element "midi-bank" of type midi-16384 order 29 depth 0
	Midi_bank int `xml:"midi-bank,omitempty"`

	// generated from element "midi-program" of type midi-128 order 30 depth 0
	Midi_program int `xml:"midi-program,omitempty"`

	// generated from element "midi-unpitched" of type midi-128 order 31 depth 0
	Midi_unpitched int `xml:"midi-unpitched,omitempty"`

	// generated from element "volume" of type percent order 32 depth 0
	Volume string `xml:"volume,omitempty"`

	// generated from element "pan" of type rotation-degrees order 33 depth 0
	Pan string `xml:"pan,omitempty"`

	// generated from element "elevation" of type rotation-degrees order 34 depth 0
	Elevation string `xml:"elevation,omitempty"`
}

// Miscellaneous is generated from named complex type "miscellaneous"
type Miscellaneous struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "miscellaneous-field" of type miscellaneous-field order 213 depth 0
	Miscellaneous_field []*Miscellaneous_field `xml:"miscellaneous-field,omitempty"`
}

// Miscellaneous_field is generated from named complex type "miscellaneous-field"
type Miscellaneous_field struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "name
	NameXSD string `xml:"name,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Mordent is generated from named complex type "mordent"
type Mordent struct {
	Name string `xml:"-"`

	// insertion point for fields
}

// Multiple_rest is generated from named complex type "multiple-rest"
type Multiple_rest struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "use-symbols
	Use_symbols string `xml:"use-symbols,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText int `xml:",chardata"`
}

// Name_display is generated from named complex type "name-display"
type Name_display struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "print-object
	AttributeGroup_print_object

	// generated from element "display-text" of type formatted-text order 35 depth 0
	Display_text []*Formatted_text `xml:"display-text,omitempty"`

	// generated from element "accidental-text" of type accidental-text order 36 depth 0
	Accidental_text []*Accidental_text `xml:"accidental-text,omitempty"`
}

// Non_arpeggiate is generated from named complex type "non-arpeggiate"
type Non_arpeggiate struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute "number
	Number int `xml:"number,attr,omitempty"`

	// generated from attribute group "position
	AttributeGroup_position

	// generated from attribute group "placement
	AttributeGroup_placement

	// generated from attribute group "color
	AttributeGroup_color

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id
}

// Notations is generated from named complex type "notations"
type Notations struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "print-object
	AttributeGroup_print_object

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from group with order 296 depth 0
	Group_editorial

	// generated from element "tied" of type tied order 297 depth 0
	Tied []*Tied `xml:"tied,omitempty"`

	// generated from element "slur" of type slur order 298 depth 0
	Slur []*Slur `xml:"slur,omitempty"`

	// generated from element "tuplet" of type tuplet order 299 depth 0
	Tuplet []*Tuplet `xml:"tuplet,omitempty"`

	// generated from element "glissando" of type glissando order 300 depth 0
	Glissando []*Glissando `xml:"glissando,omitempty"`

	// generated from element "slide" of type slide order 301 depth 0
	Slide []*Slide `xml:"slide,omitempty"`

	// generated from element "ornaments" of type ornaments order 302 depth 0
	Ornaments []*Ornaments `xml:"ornaments,omitempty"`

	// generated from element "technical" of type technical order 303 depth 0
	Technical []*Technical `xml:"technical,omitempty"`

	// generated from element "articulations" of type articulations order 304 depth 0
	Articulations []*Articulations `xml:"articulations,omitempty"`

	// generated from element "dynamics" of type dynamics order 305 depth 0
	Dynamics []*Dynamics `xml:"dynamics,omitempty"`

	// generated from element "fermata" of type fermata order 306 depth 0
	Fermata []*Fermata `xml:"fermata,omitempty"`

	// generated from element "arpeggiate" of type arpeggiate order 307 depth 0
	Arpeggiate []*Arpeggiate `xml:"arpeggiate,omitempty"`

	// generated from element "non-arpeggiate" of type non-arpeggiate order 308 depth 0
	Non_arpeggiate []*Non_arpeggiate `xml:"non-arpeggiate,omitempty"`

	// generated from element "accidental-mark" of type accidental-mark order 309 depth 0
	Accidental_mark []*Accidental_mark `xml:"accidental-mark,omitempty"`

	// generated from element "other-notation" of type other-notation order 310 depth 0
	Other_notation []*Other_notation `xml:"other-notation,omitempty"`
}

// Note is generated from named complex type "note"
type Note struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "print-leger
	Print_leger string `xml:"print-leger,attr,omitempty"`

	// generated from attribute "dynamics
	Dynamics string `xml:"dynamics,attr,omitempty"`

	// generated from attribute "end-dynamics
	End_dynamics string `xml:"end-dynamics,attr,omitempty"`

	// generated from attribute "attack
	Attack string `xml:"attack,attr,omitempty"`

	// generated from attribute "release
	Release string `xml:"release,attr,omitempty"`

	// generated from attribute "time-only
	Time_only string `xml:"time-only,attr,omitempty"`

	// generated from attribute "pizzicato
	Pizzicato string `xml:"pizzicato,attr,omitempty"`

	// generated from attribute group "x-position
	AttributeGroup_x_position

	// generated from attribute group "font
	AttributeGroup_font

	// generated from attribute group "color
	AttributeGroup_color

	// generated from attribute group "printout
	AttributeGroup_printout

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from element "grace" of type grace order 311 depth 0
	Grace []*Grace `xml:"grace,omitempty"`

	// generated from group with order 312 depth 0
	Group_full_note

	// generated from element "tie" of type tie order 313 depth 0
	Tie []*Tie `xml:"tie,omitempty"`

	// generated from element "cue" of type empty order 314 depth 0
	Cue string `xml:"cue,omitempty"`

	// generated from group with order 318 depth 0
	Group_duration

	// generated from element "instrument" of type instrument order 322 depth 0
	Instrument []*Instrument `xml:"instrument,omitempty"`

	// generated from group with order 323 depth 0
	Group_editorial_voice

	// generated from element "type" of type note-type order 324 depth 0
	Type []*Note_type `xml:"type,omitempty"`

	// generated from element "dot" of type empty-placement order 325 depth 0
	Dot []*Empty_placement `xml:"dot,omitempty"`

	// generated from element "accidental" of type accidental order 326 depth 0
	Accidental []*Accidental `xml:"accidental,omitempty"`

	// generated from element "time-modification" of type time-modification order 327 depth 0
	Time_modification []*Time_modification `xml:"time-modification,omitempty"`

	// generated from element "stem" of type stem order 328 depth 0
	Stem []*Stem `xml:"stem,omitempty"`

	// generated from element "notehead" of type notehead order 329 depth 0
	Notehead []*Notehead `xml:"notehead,omitempty"`

	// generated from element "notehead-text" of type notehead-text order 330 depth 0
	Notehead_text []*Notehead_text `xml:"notehead-text,omitempty"`

	// generated from group with order 331 depth 0
	Group_staff

	// generated from element "beam" of type beam order 332 depth 0
	Beam []*Beam `xml:"beam,omitempty"`

	// generated from element "notations" of type notations order 333 depth 0
	Notations []*Notations `xml:"notations,omitempty"`

	// generated from element "lyric" of type lyric order 334 depth 0
	Lyric []*Lyric `xml:"lyric,omitempty"`

	// generated from element "play" of type play order 335 depth 0
	Play []*Play `xml:"play,omitempty"`

	// generated from element "listen" of type listen order 336 depth 0
	Listen []*Listen `xml:"listen,omitempty"`
}

// Note_size is generated from named complex type "note-size"
type Note_size struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type string `xml:"type,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Note_type is generated from named complex type "note-type"
type Note_type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "size
	Size string `xml:"size,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Notehead is generated from named complex type "notehead"
type Notehead struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "filled
	Filled string `xml:"filled,attr,omitempty"`

	// generated from attribute "parentheses
	Parentheses string `xml:"parentheses,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Notehead_text is generated from named complex type "notehead-text"
type Notehead_text struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "display-text" of type formatted-text order 337 depth 0
	Display_text []*Formatted_text `xml:"display-text,omitempty"`

	// generated from element "accidental-text" of type accidental-text order 338 depth 0
	Accidental_text []*Accidental_text `xml:"accidental-text,omitempty"`
}

// Numeral is generated from named complex type "numeral"
type Numeral struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "numeral-root" of type numeral-root order 163 depth 0
	Numeral_root []*Numeral_root `xml:"numeral-root,omitempty"`

	// generated from element "numeral-alter" of type harmony-alter order 164 depth 0
	Numeral_alter []*Harmony_alter `xml:"numeral-alter,omitempty"`

	// generated from element "numeral-key" of type numeral-key order 165 depth 0
	Numeral_key []*Numeral_key `xml:"numeral-key,omitempty"`
}

// Numeral_key is generated from named complex type "numeral-key"
type Numeral_key struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "print-object
	AttributeGroup_print_object

	// generated from element "numeral-fifths" of type fifths order 166 depth 0
	Numeral_fifths int `xml:"numeral-fifths,omitempty"`

	// generated from element "numeral-mode" of type numeral-mode order 167 depth 0
	Numeral_mode string `xml:"numeral-mode,omitempty"`
}

// Numeral_root is generated from named complex type "numeral-root"
type Numeral_root struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "text
	Text string `xml:"text,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText int `xml:",chardata"`
}

// Octave_shift is generated from named complex type "octave-shift"
type Octave_shift struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute "number
	Number int `xml:"number,attr,omitempty"`

	// generated from attribute "size
	Size int `xml:"size,attr,omitempty"`

	// generated from attribute group "dashed-formatting
	AttributeGroup_dashed_formatting

	// generated from attribute group "print-style
	AttributeGroup_print_style

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id
}

// Offset is generated from named complex type "offset"
type Offset struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "sound
	Sound string `xml:"sound,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Opus is generated from named complex type "opus"
type Opus struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "link-attributes
	AttributeGroup_link_attributes
}

// Ornaments is generated from named complex type "ornaments"
type Ornaments struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from element "trill-mark" of type empty-trill-sound order 339 depth 0
	Trill_mark []*Empty_trill_sound `xml:"trill-mark,omitempty"`

	// generated from element "turn" of type horizontal-turn order 340 depth 0
	Turn []*Horizontal_turn `xml:"turn,omitempty"`

	// generated from element "delayed-turn" of type horizontal-turn order 341 depth 0
	Delayed_turn []*Horizontal_turn `xml:"delayed-turn,omitempty"`

	// generated from element "inverted-turn" of type horizontal-turn order 342 depth 0
	Inverted_turn []*Horizontal_turn `xml:"inverted-turn,omitempty"`

	// generated from element "delayed-inverted-turn" of type horizontal-turn order 343 depth 0
	Delayed_inverted_turn []*Horizontal_turn `xml:"delayed-inverted-turn,omitempty"`

	// generated from element "vertical-turn" of type empty-trill-sound order 344 depth 0
	Vertical_turn []*Empty_trill_sound `xml:"vertical-turn,omitempty"`

	// generated from element "inverted-vertical-turn" of type empty-trill-sound order 345 depth 0
	Inverted_vertical_turn []*Empty_trill_sound `xml:"inverted-vertical-turn,omitempty"`

	// generated from element "shake" of type empty-trill-sound order 346 depth 0
	Shake []*Empty_trill_sound `xml:"shake,omitempty"`

	// generated from element "wavy-line" of type wavy-line order 347 depth 0
	Wavy_line []*Wavy_line `xml:"wavy-line,omitempty"`

	// generated from element "mordent" of type mordent order 348 depth 0
	Mordent []*Mordent `xml:"mordent,omitempty"`

	// generated from element "inverted-mordent" of type mordent order 349 depth 0
	Inverted_mordent []*Mordent `xml:"inverted-mordent,omitempty"`

	// generated from element "schleifer" of type empty-placement order 350 depth 0
	Schleifer []*Empty_placement `xml:"schleifer,omitempty"`

	// generated from element "tremolo" of type tremolo order 351 depth 0
	Tremolo []*Tremolo `xml:"tremolo,omitempty"`

	// generated from element "haydn" of type empty-trill-sound order 352 depth 0
	Haydn []*Empty_trill_sound `xml:"haydn,omitempty"`

	// generated from element "other-ornament" of type other-placement-text order 353 depth 0
	Other_ornament []*Other_placement_text `xml:"other-ornament,omitempty"`

	// generated from element "accidental-mark" of type accidental-mark order 354 depth 0
	Accidental_mark []*Accidental_mark `xml:"accidental-mark,omitempty"`
}

// Other_appearance is generated from named complex type "other-appearance"
type Other_appearance struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type string `xml:"type,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Other_direction is generated from named complex type "other-direction"
type Other_direction struct {
	Name string `xml:"-"`

	// insertion point for fields

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Other_listening is generated from named complex type "other-listening"
type Other_listening struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute "player
	Player string `xml:"player,attr,omitempty"`

	// generated from attribute "time-only
	Time_only string `xml:"time-only,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Other_notation is generated from named complex type "other-notation"
type Other_notation struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute "number
	Number int `xml:"number,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Other_placement_text is generated from named complex type "other-placement-text"
type Other_placement_text struct {
	Name string `xml:"-"`

	// insertion point for fields

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Other_play is generated from named complex type "other-play"
type Other_play struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type string `xml:"type,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Other_text is generated from named complex type "other-text"
type Other_text struct {
	Name string `xml:"-"`

	// insertion point for fields

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Page_layout is generated from named complex type "page-layout"
type Page_layout struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "page-height" of type tenths order 220 depth 0
	Page_height string `xml:"page-height,omitempty"`

	// generated from element "page-width" of type tenths order 221 depth 0
	Page_width string `xml:"page-width,omitempty"`

	// generated from element "page-margins" of type page-margins order 222 depth 0
	Page_margins []*Page_margins `xml:"page-margins,omitempty"`
}

// Page_margins is generated from named complex type "page-margins"
type Page_margins struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type string `xml:"type,attr,omitempty"`

	// generated from group with order 223 depth 0
	Group_all_margins
}

// Part_clef is generated from named complex type "part-clef"
type Part_clef struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from group with order 67 depth 0
	Group_clef
}

// Part_group is generated from named complex type "part-group"
type Part_group struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute "number
	Number string `xml:"number,attr,omitempty"`

	// generated from element "group-name" of type group-name order 418 depth 0
	Group_name []*Group_name `xml:"group-name,omitempty"`

	// generated from element "group-name-display" of type name-display order 419 depth 0
	Group_name_display []*Name_display `xml:"group-name-display,omitempty"`

	// generated from element "group-abbreviation" of type group-name order 420 depth 0
	Group_abbreviation []*Group_name `xml:"group-abbreviation,omitempty"`

	// generated from element "group-abbreviation-display" of type name-display order 421 depth 0
	Group_abbreviation_display []*Name_display `xml:"group-abbreviation-display,omitempty"`

	// generated from element "group-symbol" of type group-symbol order 422 depth 0
	Group_symbol []*Group_symbol `xml:"group-symbol,omitempty"`

	// generated from element "group-barline" of type group-barline order 423 depth 0
	Group_barline []*Group_barline `xml:"group-barline,omitempty"`

	// generated from element "group-time" of type empty order 424 depth 0
	Group_time string `xml:"group-time,omitempty"`

	// generated from group with order 425 depth 0
	Group_editorial
}

// Part_link is generated from named complex type "part-link"
type Part_link struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "link-attributes
	AttributeGroup_link_attributes

	// generated from element "instrument-link" of type instrument-link order 426 depth 0
	Instrument_link []*Instrument_link `xml:"instrument-link,omitempty"`

	// generated from element "group-link" of type xs:string order 427 depth 0
	Group_link string `xml:"group-link,omitempty"`
}

// Part_list is generated from named complex type "part-list"
type Part_list struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from group with order 428 depth 0
	Group_part_group

	// generated from group with order 429 depth 0
	Group_score_part
}

// Part_name is generated from named complex type "part-name"
type Part_name struct {
	Name string `xml:"-"`

	// insertion point for fields

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Part_symbol is generated from named complex type "part-symbol"
type Part_symbol struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "top-staff
	Top_staff int `xml:"top-staff,attr,omitempty"`

	// generated from attribute "bottom-staff
	Bottom_staff int `xml:"bottom-staff,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Part_transpose is generated from named complex type "part-transpose"
type Part_transpose struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from group with order 68 depth 0
	Group_transpose
}

// Pedal is generated from named complex type "pedal"
type Pedal struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute "number
	Number int `xml:"number,attr,omitempty"`

	// generated from attribute "line
	Line string `xml:"line,attr,omitempty"`

	// generated from attribute "sign
	Sign string `xml:"sign,attr,omitempty"`

	// generated from attribute "abbreviated
	Abbreviated string `xml:"abbreviated,attr,omitempty"`

	// generated from attribute group "print-style-align
	AttributeGroup_print_style_align

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id
}

// Pedal_tuning is generated from named complex type "pedal-tuning"
type Pedal_tuning struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "pedal-step" of type step order 168 depth 0
	Pedal_step string `xml:"pedal-step,omitempty"`

	// generated from element "pedal-alter" of type semitones order 169 depth 0
	Pedal_alter string `xml:"pedal-alter,omitempty"`
}

// Per_minute is generated from named complex type "per-minute"
type Per_minute struct {
	Name string `xml:"-"`

	// insertion point for fields

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Percussion is generated from named complex type "percussion"
type Percussion struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "print-style-align
	AttributeGroup_print_style_align

	// generated from attribute group "enclosure
	AttributeGroup_enclosure

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from element "glass" of type glass order 170 depth 0
	Glass []*Glass `xml:"glass,omitempty"`

	// generated from element "metal" of type metal order 171 depth 0
	Metal []*Metal `xml:"metal,omitempty"`

	// generated from element "wood" of type wood order 172 depth 0
	Wood []*Wood `xml:"wood,omitempty"`

	// generated from element "pitched" of type pitched order 173 depth 0
	Pitched []*Pitched `xml:"pitched,omitempty"`

	// generated from element "membrane" of type membrane order 174 depth 0
	Membrane []*Membrane `xml:"membrane,omitempty"`

	// generated from element "effect" of type effect order 175 depth 0
	Effect []*Effect `xml:"effect,omitempty"`

	// generated from element "timpani" of type timpani order 176 depth 0
	Timpani []*Timpani `xml:"timpani,omitempty"`

	// generated from element "beater" of type beater order 177 depth 0
	Beater []*Beater `xml:"beater,omitempty"`

	// generated from element "stick" of type stick order 178 depth 0
	Stick []*Stick `xml:"stick,omitempty"`

	// generated from element "stick-location" of type stick-location order 179 depth 0
	Stick_location string `xml:"stick-location,omitempty"`

	// generated from element "other-percussion" of type other-text order 180 depth 0
	Other_percussion []*Other_text `xml:"other-percussion,omitempty"`
}

// Pitch is generated from named complex type "pitch"
type Pitch struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "step" of type step order 355 depth 0
	Step string `xml:"step,omitempty"`

	// generated from element "alter" of type semitones order 356 depth 0
	Alter string `xml:"alter,omitempty"`

	// generated from element "octave" of type octave order 357 depth 0
	Octave int `xml:"octave,omitempty"`
}

// Pitched is generated from named complex type "pitched"
type Pitched struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "smufl
	Smufl string `xml:"smufl,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Placement_text is generated from named complex type "placement-text"
type Placement_text struct {
	Name string `xml:"-"`

	// insertion point for fields

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Play is generated from named complex type "play"
type Play struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id
	Id string `xml:"id,attr,omitempty"`

	// generated from element "ipa" of type xs:string order 37 depth 0
	Ipa string `xml:"ipa,omitempty"`

	// generated from element "mute" of type mute order 38 depth 0
	Mute string `xml:"mute,omitempty"`

	// generated from element "semi-pitched" of type semi-pitched order 39 depth 0
	Semi_pitched string `xml:"semi-pitched,omitempty"`

	// generated from element "other-play" of type other-play order 40 depth 0
	Other_play []*Other_play `xml:"other-play,omitempty"`
}

// Player is generated from named complex type "player"
type Player struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id
	Id string `xml:"id,attr,omitempty"`

	// generated from element "player-name" of type xs:string order 432 depth 0
	Player_name string `xml:"player-name,omitempty"`
}

// Principal_voice is generated from named complex type "principal-voice"
type Principal_voice struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute "symbol
	Symbol string `xml:"symbol,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Print is generated from named complex type "print"
type Print struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "print-attributes
	AttributeGroup_print_attributes

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from group with order 181 depth 0
	Group_layout

	// generated from element "measure-layout" of type measure-layout order 182 depth 0
	Measure_layout []*Measure_layout `xml:"measure-layout,omitempty"`

	// generated from element "measure-numbering" of type measure-numbering order 183 depth 0
	Measure_numbering []*Measure_numbering `xml:"measure-numbering,omitempty"`

	// generated from element "part-name-display" of type name-display order 184 depth 0
	Part_name_display []*Name_display `xml:"part-name-display,omitempty"`

	// generated from element "part-abbreviation-display" of type name-display order 185 depth 0
	Part_abbreviation_display []*Name_display `xml:"part-abbreviation-display,omitempty"`
}

// Release is generated from named complex type "release"
type Release struct {
	Name string `xml:"-"`

	// insertion point for fields
}

// Repeat is generated from named complex type "repeat"
type Repeat struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "direction
	Direction string `xml:"direction,attr,omitempty"`

	// generated from attribute "times
	Times int `xml:"times,attr,omitempty"`

	// generated from attribute "after-jump
	After_jump string `xml:"after-jump,attr,omitempty"`

	// generated from attribute "winged
	Winged string `xml:"winged,attr,omitempty"`
}

// Rest is generated from named complex type "rest"
type Rest struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "measure
	Measure string `xml:"measure,attr,omitempty"`

	// generated from group with order 358 depth 0
	Group_display_step_octave
}

// Root is generated from named complex type "root"
type Root struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "root-step" of type root-step order 186 depth 0
	Root_step []*Root_step `xml:"root-step,omitempty"`

	// generated from element "root-alter" of type harmony-alter order 187 depth 0
	Root_alter []*Harmony_alter `xml:"root-alter,omitempty"`
}

// Root_step is generated from named complex type "root-step"
type Root_step struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "text
	Text string `xml:"text,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Scaling is generated from named complex type "scaling"
type Scaling struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "millimeters" of type millimeters order 224 depth 0
	Millimeters string `xml:"millimeters,omitempty"`

	// generated from element "tenths" of type tenths order 225 depth 0
	Tenths string `xml:"tenths,omitempty"`
}

// Scordatura is generated from named complex type "scordatura"
type Scordatura struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from element "accord" of type accord order 188 depth 0
	Accord []*Accord `xml:"accord,omitempty"`
}

// Score_instrument is generated from named complex type "score-instrument"
type Score_instrument struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id
	Id string `xml:"id,attr,omitempty"`

	// generated from element "instrument-name" of type xs:string order 433 depth 0
	Instrument_name string `xml:"instrument-name,omitempty"`

	// generated from element "instrument-abbreviation" of type xs:string order 434 depth 0
	Instrument_abbreviation string `xml:"instrument-abbreviation,omitempty"`

	// generated from group with order 435 depth 0
	Group_virtual_instrument_data
}

// Score_part is generated from named complex type "score-part"
type Score_part struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id
	Id string `xml:"id,attr,omitempty"`

	// generated from element "identification" of type identification order 436 depth 0
	Identification []*Identification `xml:"identification,omitempty"`

	// generated from element "part-link" of type part-link order 437 depth 0
	Part_link []*Part_link `xml:"part-link,omitempty"`

	// generated from element "part-name" of type part-name order 438 depth 0
	Part_name []*Part_name `xml:"part-name,omitempty"`

	// generated from element "part-name-display" of type name-display order 439 depth 0
	Part_name_display []*Name_display `xml:"part-name-display,omitempty"`

	// generated from element "part-abbreviation" of type part-name order 440 depth 0
	Part_abbreviation []*Part_name `xml:"part-abbreviation,omitempty"`

	// generated from element "part-abbreviation-display" of type name-display order 441 depth 0
	Part_abbreviation_display []*Name_display `xml:"part-abbreviation-display,omitempty"`

	// generated from element "group" of type xs:string order 442 depth 0
	Group string `xml:"group,omitempty"`

	// generated from element "score-instrument" of type score-instrument order 443 depth 0
	Score_instrument []*Score_instrument `xml:"score-instrument,omitempty"`

	// generated from element "player" of type player order 444 depth 0
	Player []*Player `xml:"player,omitempty"`

	// generated from element "midi-device" of type midi-device order 445 depth 0
	Midi_device []*Midi_device `xml:"midi-device,omitempty"`

	// generated from element "midi-instrument" of type midi-instrument order 446 depth 0
	Midi_instrument []*Midi_instrument `xml:"midi-instrument,omitempty"`
}

// Segno is generated from named complex type "segno"
type Segno struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "smufl
	Smufl string `xml:"smufl,attr,omitempty"`

	// generated from attribute group "print-style-align
	AttributeGroup_print_style_align

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id
}

// Slash is generated from named complex type "slash"
type Slash struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute "use-dots
	Use_dots string `xml:"use-dots,attr,omitempty"`

	// generated from attribute "use-stems
	Use_stems string `xml:"use-stems,attr,omitempty"`

	// generated from group with order 69 depth 0
	Group_slash
}

// Slide is generated from named complex type "slide"
type Slide struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute "number
	Number int `xml:"number,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Slur is generated from named complex type "slur"
type Slur struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute "number
	Number int `xml:"number,attr,omitempty"`

	// generated from attribute group "line-type
	AttributeGroup_line_type

	// generated from attribute group "dashed-formatting
	AttributeGroup_dashed_formatting

	// generated from attribute group "position
	AttributeGroup_position

	// generated from attribute group "placement
	AttributeGroup_placement

	// generated from attribute group "orientation
	AttributeGroup_orientation

	// generated from attribute group "bezier
	AttributeGroup_bezier

	// generated from attribute group "color
	AttributeGroup_color

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id
}

// Sound is generated from named complex type "sound"
type Sound struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "tempo
	Tempo string `xml:"tempo,attr,omitempty"`

	// generated from attribute "dynamics
	Dynamics string `xml:"dynamics,attr,omitempty"`

	// generated from attribute "dacapo
	Dacapo string `xml:"dacapo,attr,omitempty"`

	// generated from attribute "segno
	Segno string `xml:"segno,attr,omitempty"`

	// generated from attribute "dalsegno
	Dalsegno string `xml:"dalsegno,attr,omitempty"`

	// generated from attribute "coda
	Coda string `xml:"coda,attr,omitempty"`

	// generated from attribute "tocoda
	Tocoda string `xml:"tocoda,attr,omitempty"`

	// generated from attribute "divisions
	Divisions string `xml:"divisions,attr,omitempty"`

	// generated from attribute "forward-repeat
	Forward_repeat string `xml:"forward-repeat,attr,omitempty"`

	// generated from attribute "fine
	Fine string `xml:"fine,attr,omitempty"`

	// generated from attribute "time-only
	Time_only string `xml:"time-only,attr,omitempty"`

	// generated from attribute "pizzicato
	Pizzicato string `xml:"pizzicato,attr,omitempty"`

	// generated from attribute "pan
	Pan string `xml:"pan,attr,omitempty"`

	// generated from attribute "elevation
	Elevation string `xml:"elevation,attr,omitempty"`

	// generated from attribute "damper-pedal
	Damper_pedal string `xml:"damper-pedal,attr,omitempty"`

	// generated from attribute "soft-pedal
	Soft_pedal string `xml:"soft-pedal,attr,omitempty"`

	// generated from attribute "sostenuto-pedal
	Sostenuto_pedal string `xml:"sostenuto-pedal,attr,omitempty"`

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from element "instrument-change" of type instrument-change order 189 depth 0
	Instrument_change []*Instrument_change `xml:"instrument-change,omitempty"`

	// generated from element "midi-device" of type midi-device order 190 depth 0
	Midi_device []*Midi_device `xml:"midi-device,omitempty"`

	// generated from element "midi-instrument" of type midi-instrument order 191 depth 0
	Midi_instrument []*Midi_instrument `xml:"midi-instrument,omitempty"`

	// generated from element "play" of type play order 192 depth 0
	Play []*Play `xml:"play,omitempty"`

	// generated from element "swing" of type swing order 193 depth 0
	Swing []*Swing `xml:"swing,omitempty"`

	// generated from element "offset" of type offset order 194 depth 0
	Offset []*Offset `xml:"offset,omitempty"`
}

// Staff_details is generated from named complex type "staff-details"
type Staff_details struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "number
	Number int `xml:"number,attr,omitempty"`

	// generated from attribute "show-frets
	Show_frets string `xml:"show-frets,attr,omitempty"`

	// generated from attribute group "print-object
	AttributeGroup_print_object

	// generated from attribute group "print-spacing
	AttributeGroup_print_spacing

	// generated from element "staff-type" of type staff-type order 70 depth 0
	Staff_type string `xml:"staff-type,omitempty"`

	// generated from element "staff-lines" of type xs:nonNegativeInteger order 71 depth 0
	Staff_lines int `xml:"staff-lines,omitempty"`

	// generated from element "line-detail" of type line-detail order 72 depth 0
	Line_detail []*Line_detail `xml:"line-detail,omitempty"`

	// generated from element "staff-tuning" of type staff-tuning order 73 depth 0
	Staff_tuning []*Staff_tuning `xml:"staff-tuning,omitempty"`

	// generated from element "capo" of type xs:nonNegativeInteger order 74 depth 0
	Capo int `xml:"capo,omitempty"`

	// generated from element "staff-size" of type staff-size order 75 depth 0
	Staff_size []*Staff_size `xml:"staff-size,omitempty"`
}

// Staff_divide is generated from named complex type "staff-divide"
type Staff_divide struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute group "print-style-align
	AttributeGroup_print_style_align

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id
}

// Staff_layout is generated from named complex type "staff-layout"
type Staff_layout struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "number
	Number int `xml:"number,attr,omitempty"`

	// generated from element "staff-distance" of type tenths order 226 depth 0
	Staff_distance string `xml:"staff-distance,omitempty"`
}

// Staff_size is generated from named complex type "staff-size"
type Staff_size struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "scaling
	Scaling string `xml:"scaling,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Staff_tuning is generated from named complex type "staff-tuning"
type Staff_tuning struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "line
	Line int `xml:"line,attr,omitempty"`

	// generated from group with order 76 depth 0
	Group_tuning
}

// Stem is generated from named complex type "stem"
type Stem struct {
	Name string `xml:"-"`

	// insertion point for fields

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Stick is generated from named complex type "stick"
type Stick struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "tip
	Tip string `xml:"tip,attr,omitempty"`

	// generated from attribute "parentheses
	Parentheses string `xml:"parentheses,attr,omitempty"`

	// generated from attribute "dashed-circle
	Dashed_circle string `xml:"dashed-circle,attr,omitempty"`

	// generated from element "stick-type" of type stick-type order 195 depth 0
	Stick_type string `xml:"stick-type,omitempty"`

	// generated from element "stick-material" of type stick-material order 196 depth 0
	Stick_material string `xml:"stick-material,omitempty"`
}

// String_mute is generated from named complex type "string-mute"
type String_mute struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute group "print-style-align
	AttributeGroup_print_style_align

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id
}

// String_type is generated from named complex type "string-type"
type String_type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText int `xml:",chardata"`
}

// Strong_accent is generated from named complex type "strong-accent"
type Strong_accent struct {
	Name string `xml:"-"`

	// insertion point for fields
}

// Style_text is generated from named complex type "style-text"
type Style_text struct {
	Name string `xml:"-"`

	// insertion point for fields

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Supports is generated from named complex type "supports"
type Supports struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute "element
	Element string `xml:"element,attr,omitempty"`

	// generated from attribute "attribute
	Attribute string `xml:"attribute,attr,omitempty"`

	// generated from attribute "value
	Value string `xml:"value,attr,omitempty"`
}

// Swing is generated from named complex type "swing"
type Swing struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "straight" of type empty order 197 depth 0
	Straight string `xml:"straight,omitempty"`

	// generated from element "first" of type xs:positiveInteger order 198 depth 0
	First int `xml:"first,omitempty"`

	// generated from element "second" of type xs:positiveInteger order 199 depth 0
	Second int `xml:"second,omitempty"`

	// generated from element "swing-type" of type swing-type-value order 200 depth 0
	Swing_type string `xml:"swing-type,omitempty"`

	// generated from element "swing-style" of type xs:string order 201 depth 0
	Swing_style string `xml:"swing-style,omitempty"`
}

// Sync is generated from named complex type "sync"
type Sync struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute "latency
	Latency int `xml:"latency,attr,omitempty"`

	// generated from attribute "player
	Player string `xml:"player,attr,omitempty"`

	// generated from attribute "time-only
	Time_only string `xml:"time-only,attr,omitempty"`
}

// System_dividers is generated from named complex type "system-dividers"
type System_dividers struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "left-divider" of type empty-print-object-style-align order 227 depth 0
	Left_divider []*Empty_print_object_style_align `xml:"left-divider,omitempty"`

	// generated from element "right-divider" of type empty-print-object-style-align order 228 depth 0
	Right_divider []*Empty_print_object_style_align `xml:"right-divider,omitempty"`
}

// System_layout is generated from named complex type "system-layout"
type System_layout struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "system-margins" of type system-margins order 229 depth 0
	System_margins []*System_margins `xml:"system-margins,omitempty"`

	// generated from element "system-distance" of type tenths order 230 depth 0
	System_distance string `xml:"system-distance,omitempty"`

	// generated from element "top-system-distance" of type tenths order 231 depth 0
	Top_system_distance string `xml:"top-system-distance,omitempty"`

	// generated from element "system-dividers" of type system-dividers order 232 depth 0
	System_dividers []*System_dividers `xml:"system-dividers,omitempty"`
}

// System_margins is generated from named complex type "system-margins"
type System_margins struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from group with order 233 depth 0
	Group_left_right_margins
}

// Tap is generated from named complex type "tap"
type Tap struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "hand
	Hand string `xml:"hand,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Technical is generated from named complex type "technical"
type Technical struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from element "up-bow" of type empty-placement order 359 depth 0
	Up_bow []*Empty_placement `xml:"up-bow,omitempty"`

	// generated from element "down-bow" of type empty-placement order 360 depth 0
	Down_bow []*Empty_placement `xml:"down-bow,omitempty"`

	// generated from element "harmonic" of type harmonic order 361 depth 0
	Harmonic []*Harmonic `xml:"harmonic,omitempty"`

	// generated from element "open-string" of type empty-placement order 362 depth 0
	Open_string []*Empty_placement `xml:"open-string,omitempty"`

	// generated from element "thumb-position" of type empty-placement order 363 depth 0
	Thumb_position []*Empty_placement `xml:"thumb-position,omitempty"`

	// generated from element "fingering" of type fingering order 364 depth 0
	Fingering []*Fingering `xml:"fingering,omitempty"`

	// generated from element "pluck" of type placement-text order 365 depth 0
	Pluck []*Placement_text `xml:"pluck,omitempty"`

	// generated from element "double-tongue" of type empty-placement order 366 depth 0
	Double_tongue []*Empty_placement `xml:"double-tongue,omitempty"`

	// generated from element "triple-tongue" of type empty-placement order 367 depth 0
	Triple_tongue []*Empty_placement `xml:"triple-tongue,omitempty"`

	// generated from element "stopped" of type empty-placement-smufl order 368 depth 0
	Stopped []*Empty_placement_smufl `xml:"stopped,omitempty"`

	// generated from element "snap-pizzicato" of type empty-placement order 369 depth 0
	Snap_pizzicato []*Empty_placement `xml:"snap-pizzicato,omitempty"`

	// generated from element "fret" of type fret order 370 depth 0
	Fret []*Fret `xml:"fret,omitempty"`

	// generated from element "string" of type string-type order 371 depth 0
	String []*String_type `xml:"string,omitempty"`

	// generated from element "hammer-on" of type hammer-on-pull-off order 372 depth 0
	Hammer_on []*Hammer_on_pull_off `xml:"hammer-on,omitempty"`

	// generated from element "pull-off" of type hammer-on-pull-off order 373 depth 0
	Pull_off []*Hammer_on_pull_off `xml:"pull-off,omitempty"`

	// generated from element "bend" of type bend order 374 depth 0
	Bend []*Bend `xml:"bend,omitempty"`

	// generated from element "tap" of type tap order 375 depth 0
	Tap []*Tap `xml:"tap,omitempty"`

	// generated from element "heel" of type heel-toe order 376 depth 0
	Heel []*Heel_toe `xml:"heel,omitempty"`

	// generated from element "toe" of type heel-toe order 377 depth 0
	Toe []*Heel_toe `xml:"toe,omitempty"`

	// generated from element "fingernails" of type empty-placement order 378 depth 0
	Fingernails []*Empty_placement `xml:"fingernails,omitempty"`

	// generated from element "hole" of type hole order 379 depth 0
	Hole []*Hole `xml:"hole,omitempty"`

	// generated from element "arrow" of type arrow order 380 depth 0
	Arrow []*Arrow `xml:"arrow,omitempty"`

	// generated from element "handbell" of type handbell order 381 depth 0
	Handbell []*Handbell `xml:"handbell,omitempty"`

	// generated from element "brass-bend" of type empty-placement order 382 depth 0
	Brass_bend []*Empty_placement `xml:"brass-bend,omitempty"`

	// generated from element "flip" of type empty-placement order 383 depth 0
	Flip []*Empty_placement `xml:"flip,omitempty"`

	// generated from element "smear" of type empty-placement order 384 depth 0
	Smear []*Empty_placement `xml:"smear,omitempty"`

	// generated from element "open" of type empty-placement-smufl order 385 depth 0
	Open []*Empty_placement_smufl `xml:"open,omitempty"`

	// generated from element "half-muted" of type empty-placement-smufl order 386 depth 0
	Half_muted []*Empty_placement_smufl `xml:"half-muted,omitempty"`

	// generated from element "harmon-mute" of type harmon-mute order 387 depth 0
	Harmon_mute []*Harmon_mute `xml:"harmon-mute,omitempty"`

	// generated from element "golpe" of type empty-placement order 388 depth 0
	Golpe []*Empty_placement `xml:"golpe,omitempty"`

	// generated from element "other-technical" of type other-placement-text order 389 depth 0
	Other_technical []*Other_placement_text `xml:"other-technical,omitempty"`
}

// Text_element_data is generated from named complex type "text-element-data"
type Text_element_data struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "http://www.w3.org/XML/1998/namespace lang
	Lang string `xml:"http://www.w3.org/XML/1998/namespace lang,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Tie is generated from named complex type "tie"
type Tie struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute "time-only
	Time_only string `xml:"time-only,attr,omitempty"`
}

// Tied is generated from named complex type "tied"
type Tied struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute "number
	Number int `xml:"number,attr,omitempty"`

	// generated from attribute group "line-type
	AttributeGroup_line_type

	// generated from attribute group "dashed-formatting
	AttributeGroup_dashed_formatting

	// generated from attribute group "position
	AttributeGroup_position

	// generated from attribute group "placement
	AttributeGroup_placement

	// generated from attribute group "orientation
	AttributeGroup_orientation

	// generated from attribute group "bezier
	AttributeGroup_bezier

	// generated from attribute group "color
	AttributeGroup_color

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id
}

// Time is generated from named complex type "time"
type Time struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "number
	Number int `xml:"number,attr,omitempty"`

	// generated from attribute "symbol
	Symbol string `xml:"symbol,attr,omitempty"`

	// generated from attribute "separator
	Separator string `xml:"separator,attr,omitempty"`

	// generated from attribute group "print-style-align
	AttributeGroup_print_style_align

	// generated from attribute group "print-object
	AttributeGroup_print_object

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from group with order 77 depth 0
	Group_time_signature

	// generated from element "interchangeable" of type interchangeable order 78 depth 0
	Interchangeable []*Interchangeable `xml:"interchangeable,omitempty"`

	// generated from element "senza-misura" of type xs:string order 79 depth 0
	Senza_misura string `xml:"senza-misura,omitempty"`
}

// Time_modification is generated from named complex type "time-modification"
type Time_modification struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "actual-notes" of type xs:nonNegativeInteger order 390 depth 0
	Actual_notes int `xml:"actual-notes,omitempty"`

	// generated from element "normal-notes" of type xs:nonNegativeInteger order 391 depth 0
	Normal_notes int `xml:"normal-notes,omitempty"`

	// generated from element "normal-type" of type note-type-value order 392 depth 0
	Normal_type string `xml:"normal-type,omitempty"`

	// generated from element "normal-dot" of type empty order 393 depth 0
	Normal_dot string `xml:"normal-dot,omitempty"`
}

// Timpani is generated from named complex type "timpani"
type Timpani struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "smufl
	Smufl string `xml:"smufl,attr,omitempty"`
}

// Transpose is generated from named complex type "transpose"
type Transpose struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "number
	Number int `xml:"number,attr,omitempty"`

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from group with order 80 depth 0
	Group_transpose
}

// Tremolo is generated from named complex type "tremolo"
type Tremolo struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type string `xml:"type,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText int `xml:",chardata"`
}

// Tuplet is generated from named complex type "tuplet"
type Tuplet struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute "number
	Number int `xml:"number,attr,omitempty"`

	// generated from attribute "bracket
	Bracket string `xml:"bracket,attr,omitempty"`

	// generated from attribute "show-number
	Show_number string `xml:"show-number,attr,omitempty"`

	// generated from attribute "show-type
	Show_type string `xml:"show-type,attr,omitempty"`

	// generated from attribute group "line-shape
	AttributeGroup_line_shape

	// generated from attribute group "position
	AttributeGroup_position

	// generated from attribute group "placement
	AttributeGroup_placement

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from element "tuplet-actual" of type tuplet-portion order 394 depth 0
	Tuplet_actual []*Tuplet_portion `xml:"tuplet-actual,omitempty"`

	// generated from element "tuplet-normal" of type tuplet-portion order 395 depth 0
	Tuplet_normal []*Tuplet_portion `xml:"tuplet-normal,omitempty"`
}

// Tuplet_dot is generated from named complex type "tuplet-dot"
type Tuplet_dot struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "font
	AttributeGroup_font

	// generated from attribute group "color
	AttributeGroup_color
}

// Tuplet_number is generated from named complex type "tuplet-number"
type Tuplet_number struct {
	Name string `xml:"-"`

	// insertion point for fields

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText int `xml:",chardata"`
}

// Tuplet_portion is generated from named complex type "tuplet-portion"
type Tuplet_portion struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "tuplet-number" of type tuplet-number order 396 depth 0
	Tuplet_number []*Tuplet_number `xml:"tuplet-number,omitempty"`

	// generated from element "tuplet-type" of type tuplet-type order 397 depth 0
	Tuplet_type []*Tuplet_type `xml:"tuplet-type,omitempty"`

	// generated from element "tuplet-dot" of type tuplet-dot order 398 depth 0
	Tuplet_dot []*Tuplet_dot `xml:"tuplet-dot,omitempty"`
}

// Tuplet_type is generated from named complex type "tuplet-type"
type Tuplet_type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Typed_text is generated from named complex type "typed-text"
type Typed_text struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type string `xml:"type,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Unpitched is generated from named complex type "unpitched"
type Unpitched struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from group with order 399 depth 0
	Group_display_step_octave
}

// Virtual_instrument is generated from named complex type "virtual-instrument"
type Virtual_instrument struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "virtual-library" of type xs:string order 447 depth 0
	Virtual_library string `xml:"virtual-library,omitempty"`

	// generated from element "virtual-name" of type xs:string order 448 depth 0
	Virtual_name string `xml:"virtual-name,omitempty"`
}

// Wait is generated from named complex type "wait"
type Wait struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "player
	Player string `xml:"player,attr,omitempty"`

	// generated from attribute "time-only
	Time_only string `xml:"time-only,attr,omitempty"`
}

// Wavy_line is generated from named complex type "wavy-line"
type Wavy_line struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute "number
	Number int `xml:"number,attr,omitempty"`

	// generated from attribute "smufl
	Smufl string `xml:"smufl,attr,omitempty"`

	// generated from attribute group "position
	AttributeGroup_position

	// generated from attribute group "placement
	AttributeGroup_placement

	// generated from attribute group "color
	AttributeGroup_color

	// generated from attribute group "trill-sound
	AttributeGroup_trill_sound
}

// Wedge is generated from named complex type "wedge"
type Wedge struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute "number
	Number int `xml:"number,attr,omitempty"`

	// generated from attribute "spread
	Spread string `xml:"spread,attr,omitempty"`

	// generated from attribute "niente
	Niente string `xml:"niente,attr,omitempty"`

	// generated from attribute group "line-type
	AttributeGroup_line_type

	// generated from attribute group "dashed-formatting
	AttributeGroup_dashed_formatting

	// generated from attribute group "position
	AttributeGroup_position

	// generated from attribute group "color
	AttributeGroup_color

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id
}

// Wood is generated from named complex type "wood"
type Wood struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "smufl
	Smufl string `xml:"smufl,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Work is generated from named complex type "work"
type Work struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "work-number" of type xs:string order 449 depth 0
	Work_number string `xml:"work-number,omitempty"`

	// generated from element "work-title" of type xs:string order 450 depth 0
	Work_title string `xml:"work-title,omitempty"`

	// generated from element "opus" of type opus order 451 depth 0
	Opus []*Opus `xml:"opus,omitempty"`
}

// Group_all_margins is generated from named group "all-margins"
type Group_all_margins struct {

	// insertion point for fields

	// generated from group with order 516 depth 1
	Group_left_right_margins

	// generated from element "top-margin" of type tenths order 517 depth 1
	Top_margin string `xml:"top-margin,omitempty"`

	// generated from element "bottom-margin" of type tenths order 518 depth 1
	Bottom_margin string `xml:"bottom-margin,omitempty"`
}

// Group_beat_unit is generated from named group "beat-unit"
type Group_beat_unit struct {

	// insertion point for fields

	// generated from element "beat-unit" of type note-type-value order 505 depth 1
	Beat_unit string `xml:"beat-unit,omitempty"`

	// generated from element "beat-unit-dot" of type empty order 506 depth 1
	Beat_unit_dot string `xml:"beat-unit-dot,omitempty"`
}

// Group_clef is generated from named group "clef"
type Group_clef struct {

	// insertion point for fields

	// generated from element "sign" of type clef-sign order 481 depth 1
	Sign string `xml:"sign,omitempty"`

	// generated from element "line" of type staff-line-position order 482 depth 1
	Line int `xml:"line,omitempty"`

	// generated from element "clef-octave-change" of type xs:integer order 483 depth 1
	Clef_octave_change int `xml:"clef-octave-change,omitempty"`
}

// Group_display_step_octave is generated from named group "display-step-octave"
type Group_display_step_octave struct {

	// insertion point for fields

	// generated from element "display-step" of type step order 529 depth 1
	Display_step string `xml:"display-step,omitempty"`

	// generated from element "display-octave" of type octave order 530 depth 1
	Display_octave int `xml:"display-octave,omitempty"`
}

// Group_duration is generated from named group "duration"
type Group_duration struct {

	// insertion point for fields

	// generated from element "duration" of type positive-divisions order 527 depth 1
	Duration string `xml:"duration,omitempty"`
}

// Group_editorial is generated from named group "editorial"
type Group_editorial struct {

	// insertion point for fields

	// generated from group with order 453 depth 1
	Group_footnote

	// generated from group with order 454 depth 1
	Group_level
}

// Group_editorial_voice is generated from named group "editorial-voice"
type Group_editorial_voice struct {

	// insertion point for fields

	// generated from group with order 456 depth 1
	Group_footnote

	// generated from group with order 457 depth 1
	Group_level

	// generated from group with order 458 depth 1
	Group_voice
}

// Group_editorial_voice_direction is generated from named group "editorial-voice-direction"
type Group_editorial_voice_direction struct {

	// insertion point for fields

	// generated from group with order 460 depth 1
	Group_footnote

	// generated from group with order 461 depth 1
	Group_level

	// generated from group with order 462 depth 1
	Group_voice
}

// Group_footnote is generated from named group "footnote"
type Group_footnote struct {

	// insertion point for fields

	// generated from element "footnote" of type formatted-text order 464 depth 1
	Footnote []*Formatted_text `xml:"footnote,omitempty"`
}

// Group_full_note is generated from named group "full-note"
type Group_full_note struct {

	// insertion point for fields

	// generated from element "chord" of type empty order 532 depth 1
	Chord string `xml:"chord,omitempty"`

	// generated from element "pitch" of type pitch order 533 depth 1
	Pitch []*Pitch `xml:"pitch,omitempty"`

	// generated from element "unpitched" of type unpitched order 534 depth 1
	Unpitched []*Unpitched `xml:"unpitched,omitempty"`

	// generated from element "rest" of type rest order 535 depth 1
	Rest []*Rest `xml:"rest,omitempty"`
}

// Group_harmony_chord is generated from named group "harmony-chord"
type Group_harmony_chord struct {

	// insertion point for fields

	// generated from element "root" of type root order 508 depth 1
	Root []*Root `xml:"root,omitempty"`

	// generated from element "numeral" of type numeral order 509 depth 1
	Numeral []*Numeral `xml:"numeral,omitempty"`

	// generated from element "function" of type style-text order 510 depth 1
	Function []*Style_text `xml:"function,omitempty"`

	// generated from element "kind" of type kind order 511 depth 1
	Kind []*Kind `xml:"kind,omitempty"`

	// generated from element "inversion" of type inversion order 512 depth 1
	Inversion []*Inversion `xml:"inversion,omitempty"`

	// generated from element "bass" of type bass order 513 depth 1
	Bass []*Bass `xml:"bass,omitempty"`

	// generated from element "degree" of type degree order 514 depth 1
	Degree []*Degree `xml:"degree,omitempty"`
}

// Group_layout is generated from named group "layout"
type Group_layout struct {

	// insertion point for fields

	// generated from element "page-layout" of type page-layout order 520 depth 1
	Page_layout []*Page_layout `xml:"page-layout,omitempty"`

	// generated from element "system-layout" of type system-layout order 521 depth 1
	System_layout []*System_layout `xml:"system-layout,omitempty"`

	// generated from element "staff-layout" of type staff-layout order 522 depth 1
	Staff_layout []*Staff_layout `xml:"staff-layout,omitempty"`
}

// Group_left_right_margins is generated from named group "left-right-margins"
type Group_left_right_margins struct {

	// insertion point for fields

	// generated from element "left-margin" of type tenths order 524 depth 1
	Left_margin string `xml:"left-margin,omitempty"`

	// generated from element "right-margin" of type tenths order 525 depth 1
	Right_margin string `xml:"right-margin,omitempty"`
}

// Group_level is generated from named group "level"
type Group_level struct {

	// insertion point for fields

	// generated from element "level" of type level order 466 depth 1
	Level []*Level `xml:"level,omitempty"`
}

// Group_music_data is generated from named group "music-data"
type Group_music_data struct {

	// insertion point for fields

	// generated from element "note" of type note order 537 depth 1
	Note []*Note `xml:"note,omitempty"`

	// generated from element "backup" of type backup order 538 depth 1
	Backup []*Backup `xml:"backup,omitempty"`

	// generated from element "forward" of type forward order 539 depth 1
	Forward []*Forward `xml:"forward,omitempty"`

	// generated from element "direction" of type direction order 540 depth 1
	Direction []*Direction `xml:"direction,omitempty"`

	// generated from element "attributes" of type attributes order 541 depth 1
	Attributes []*Attributes `xml:"attributes,omitempty"`

	// generated from element "harmony" of type harmony order 542 depth 1
	Harmony []*Harmony `xml:"harmony,omitempty"`

	// generated from element "figured-bass" of type figured-bass order 543 depth 1
	Figured_bass []*Figured_bass `xml:"figured-bass,omitempty"`

	// generated from element "print" of type print order 544 depth 1
	Print []*Print `xml:"print,omitempty"`

	// generated from element "sound" of type sound order 545 depth 1
	Sound []*Sound `xml:"sound,omitempty"`

	// generated from element "listening" of type listening order 546 depth 1
	Listening []*Listening `xml:"listening,omitempty"`

	// generated from element "barline" of type barline order 547 depth 1
	Barline []*Barline `xml:"barline,omitempty"`

	// generated from element "grouping" of type grouping order 548 depth 1
	Grouping []*Grouping `xml:"grouping,omitempty"`

	// generated from element "link" of type link order 549 depth 1
	Link []*Link `xml:"link,omitempty"`

	// generated from element "bookmark" of type bookmark order 550 depth 1
	Bookmark []*Bookmark `xml:"bookmark,omitempty"`
}

// Group_non_traditional_key is generated from named group "non-traditional-key"
type Group_non_traditional_key struct {

	// insertion point for fields

	// generated from element "key-step" of type step order 485 depth 1
	Key_step string `xml:"key-step,omitempty"`

	// generated from element "key-alter" of type semitones order 486 depth 1
	Key_alter string `xml:"key-alter,omitempty"`

	// generated from element "key-accidental" of type key-accidental order 487 depth 1
	Key_accidental []*Key_accidental `xml:"key-accidental,omitempty"`
}

// Group_part_group is generated from named group "part-group"
type Group_part_group struct {

	// insertion point for fields

	// generated from element "part-group" of type part-group order 552 depth 1
	Part_group []*Part_group `xml:"part-group,omitempty"`
}

// Group_score_header is generated from named group "score-header"
type Group_score_header struct {

	// insertion point for fields

	// generated from element "work" of type work order 554 depth 1
	Work []*Work `xml:"work,omitempty"`

	// generated from element "movement-number" of type xs:string order 555 depth 1
	Movement_number string `xml:"movement-number,omitempty"`

	// generated from element "movement-title" of type xs:string order 556 depth 1
	Movement_title string `xml:"movement-title,omitempty"`

	// generated from element "identification" of type identification order 557 depth 1
	Identification []*Identification `xml:"identification,omitempty"`

	// generated from element "defaults" of type defaults order 558 depth 1
	Defaults []*Defaults `xml:"defaults,omitempty"`

	// generated from element "credit" of type credit order 559 depth 1
	Credit []*Credit `xml:"credit,omitempty"`

	// generated from element "part-list" of type part-list order 560 depth 1
	Part_list []*Part_list `xml:"part-list,omitempty"`
}

// Group_score_part is generated from named group "score-part"
type Group_score_part struct {

	// insertion point for fields

	// generated from element "score-part" of type score-part order 562 depth 1
	Score_part []*Score_part `xml:"score-part,omitempty"`
}

// Group_slash is generated from named group "slash"
type Group_slash struct {

	// insertion point for fields

	// generated from element "slash-type" of type note-type-value order 489 depth 1
	Slash_type string `xml:"slash-type,omitempty"`

	// generated from element "slash-dot" of type empty order 490 depth 1
	Slash_dot string `xml:"slash-dot,omitempty"`

	// generated from element "except-voice" of type xs:string order 491 depth 1
	Except_voice string `xml:"except-voice,omitempty"`
}

// Group_staff is generated from named group "staff"
type Group_staff struct {

	// insertion point for fields

	// generated from element "staff" of type xs:positiveInteger order 468 depth 1
	Staff int `xml:"staff,omitempty"`
}

// Group_time_signature is generated from named group "time-signature"
type Group_time_signature struct {

	// insertion point for fields

	// generated from element "beats" of type xs:string order 493 depth 1
	Beats string `xml:"beats,omitempty"`

	// generated from element "beat-type" of type xs:string order 494 depth 1
	Beat_type string `xml:"beat-type,omitempty"`
}

// Group_traditional_key is generated from named group "traditional-key"
type Group_traditional_key struct {

	// insertion point for fields

	// generated from element "cancel" of type cancel order 496 depth 1
	Cancel []*Cancel `xml:"cancel,omitempty"`

	// generated from element "fifths" of type fifths order 497 depth 1
	Fifths int `xml:"fifths,omitempty"`

	// generated from element "mode" of type mode order 498 depth 1
	Mode string `xml:"mode,omitempty"`
}

// Group_transpose is generated from named group "transpose"
type Group_transpose struct {

	// insertion point for fields

	// generated from element "diatonic" of type xs:integer order 500 depth 1
	Diatonic int `xml:"diatonic,omitempty"`

	// generated from element "chromatic" of type semitones order 501 depth 1
	Chromatic string `xml:"chromatic,omitempty"`

	// generated from element "octave-change" of type xs:integer order 502 depth 1
	Octave_change int `xml:"octave-change,omitempty"`

	// generated from element "double" of type double order 503 depth 1
	Double []*Double `xml:"double,omitempty"`
}

// Group_tuning is generated from named group "tuning"
type Group_tuning struct {

	// insertion point for fields

	// generated from element "tuning-step" of type step order 470 depth 1
	Tuning_step string `xml:"tuning-step,omitempty"`

	// generated from element "tuning-alter" of type semitones order 471 depth 1
	Tuning_alter string `xml:"tuning-alter,omitempty"`

	// generated from element "tuning-octave" of type octave order 472 depth 1
	Tuning_octave int `xml:"tuning-octave,omitempty"`
}

// Group_virtual_instrument_data is generated from named group "virtual-instrument-data"
type Group_virtual_instrument_data struct {

	// insertion point for fields

	// generated from element "instrument-sound" of type xs:string order 474 depth 1
	Instrument_sound string `xml:"instrument-sound,omitempty"`

	// generated from element "solo" of type empty order 475 depth 1
	Solo string `xml:"solo,omitempty"`

	// generated from element "ensemble" of type positive-integer-or-empty order 476 depth 1
	Ensemble string `xml:"ensemble,omitempty"`

	// generated from element "virtual-instrument" of type virtual-instrument order 477 depth 1
	Virtual_instrument []*Virtual_instrument `xml:"virtual-instrument,omitempty"`
}

// Group_voice is generated from named group "voice"
type Group_voice struct {

	// insertion point for fields

	// generated from element "voice" of type xs:string order 479 depth 1
	Voice string `xml:"voice,omitempty"`
}

// AttributeGroup_bend_sound is generated from named attribute group "bend-sound"
type AttributeGroup_bend_sound struct {

	// insertion point for fields

	// generated from attribute "accelerate
	Accelerate string `xml:"accelerate,attr,omitempty"`

	// generated from attribute "beats
	Beats string `xml:"beats,attr,omitempty"`

	// generated from attribute "first-beat
	First_beat string `xml:"first-beat,attr,omitempty"`

	// generated from attribute "last-beat
	Last_beat string `xml:"last-beat,attr,omitempty"`
}

// AttributeGroup_bezier is generated from named attribute group "bezier"
type AttributeGroup_bezier struct {

	// insertion point for fields

	// generated from attribute "bezier-x
	Bezier_x string `xml:"bezier-x,attr,omitempty"`

	// generated from attribute "bezier-y
	Bezier_y string `xml:"bezier-y,attr,omitempty"`

	// generated from attribute "bezier-x2
	Bezier_x2 string `xml:"bezier-x2,attr,omitempty"`

	// generated from attribute "bezier-y2
	Bezier_y2 string `xml:"bezier-y2,attr,omitempty"`

	// generated from attribute "bezier-offset
	Bezier_offset string `xml:"bezier-offset,attr,omitempty"`

	// generated from attribute "bezier-offset2
	Bezier_offset2 string `xml:"bezier-offset2,attr,omitempty"`
}

// AttributeGroup_color is generated from named attribute group "color"
type AttributeGroup_color struct {

	// insertion point for fields

	// generated from attribute "color
	Color string `xml:"color,attr,omitempty"`
}

// AttributeGroup_dashed_formatting is generated from named attribute group "dashed-formatting"
type AttributeGroup_dashed_formatting struct {

	// insertion point for fields

	// generated from attribute "dash-length
	Dash_length string `xml:"dash-length,attr,omitempty"`

	// generated from attribute "space-length
	Space_length string `xml:"space-length,attr,omitempty"`
}

// AttributeGroup_directive is generated from named attribute group "directive"
type AttributeGroup_directive struct {

	// insertion point for fields

	// generated from attribute "directive
	Directive string `xml:"directive,attr,omitempty"`
}

// AttributeGroup_document_attributes is generated from named attribute group "document-attributes"
type AttributeGroup_document_attributes struct {

	// insertion point for fields

	// generated from attribute "version
	Version string `xml:"version,attr,omitempty"`
}

// AttributeGroup_element_position is generated from named attribute group "element-position"
type AttributeGroup_element_position struct {

	// insertion point for fields

	// generated from attribute "element
	Element string `xml:"element,attr,omitempty"`

	// generated from attribute "position
	Position int `xml:"position,attr,omitempty"`
}

// AttributeGroup_enclosure is generated from named attribute group "enclosure"
type AttributeGroup_enclosure struct {

	// insertion point for fields

	// generated from attribute "enclosure
	Enclosure string `xml:"enclosure,attr,omitempty"`
}

// AttributeGroup_font is generated from named attribute group "font"
type AttributeGroup_font struct {

	// insertion point for fields

	// generated from attribute "font-family
	Font_family string `xml:"font-family,attr,omitempty"`

	// generated from attribute "font-style
	Font_style string `xml:"font-style,attr,omitempty"`

	// generated from attribute "font-size
	Font_size string `xml:"font-size,attr,omitempty"`

	// generated from attribute "font-weight
	Font_weight string `xml:"font-weight,attr,omitempty"`
}

// AttributeGroup_group_name_text is generated from named attribute group "group-name-text"
type AttributeGroup_group_name_text struct {

	// insertion point for fields

	// generated from attribute group "print-style
	AttributeGroup_print_style

	// generated from attribute group "justify
	AttributeGroup_justify
}

// AttributeGroup_halign is generated from named attribute group "halign"
type AttributeGroup_halign struct {

	// insertion point for fields

	// generated from attribute "halign
	Halign string `xml:"halign,attr,omitempty"`
}

// AttributeGroup_image_attributes is generated from named attribute group "image-attributes"
type AttributeGroup_image_attributes struct {

	// insertion point for fields

	// generated from attribute "source
	Source string `xml:"source,attr,omitempty"`

	// generated from attribute "type
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute "height
	Height string `xml:"height,attr,omitempty"`

	// generated from attribute "width
	Width string `xml:"width,attr,omitempty"`

	// generated from attribute group "position
	AttributeGroup_position

	// generated from attribute group "halign
	AttributeGroup_halign

	// generated from attribute group "valign-image
	AttributeGroup_valign_image
}

// AttributeGroup_justify is generated from named attribute group "justify"
type AttributeGroup_justify struct {

	// insertion point for fields

	// generated from attribute "justify
	Justify string `xml:"justify,attr,omitempty"`
}

// AttributeGroup_letter_spacing is generated from named attribute group "letter-spacing"
type AttributeGroup_letter_spacing struct {

	// insertion point for fields

	// generated from attribute "letter-spacing
	Letter_spacing string `xml:"letter-spacing,attr,omitempty"`
}

// AttributeGroup_level_display is generated from named attribute group "level-display"
type AttributeGroup_level_display struct {

	// insertion point for fields

	// generated from attribute "parentheses
	Parentheses string `xml:"parentheses,attr,omitempty"`

	// generated from attribute "bracket
	Bracket string `xml:"bracket,attr,omitempty"`

	// generated from attribute "size
	Size string `xml:"size,attr,omitempty"`
}

// AttributeGroup_line_height is generated from named attribute group "line-height"
type AttributeGroup_line_height struct {

	// insertion point for fields

	// generated from attribute "line-height
	Line_height string `xml:"line-height,attr,omitempty"`
}

// AttributeGroup_line_length is generated from named attribute group "line-length"
type AttributeGroup_line_length struct {

	// insertion point for fields

	// generated from attribute "line-length
	Line_length string `xml:"line-length,attr,omitempty"`
}

// AttributeGroup_line_shape is generated from named attribute group "line-shape"
type AttributeGroup_line_shape struct {

	// insertion point for fields

	// generated from attribute "line-shape
	Line_shape string `xml:"line-shape,attr,omitempty"`
}

// AttributeGroup_line_type is generated from named attribute group "line-type"
type AttributeGroup_line_type struct {

	// insertion point for fields

	// generated from attribute "line-type
	Line_type string `xml:"line-type,attr,omitempty"`
}

// AttributeGroup_link_attributes is generated from named attribute group "link-attributes"
type AttributeGroup_link_attributes struct {

	// insertion point for fields

	// generated from attribute "http://www.w3.org/1999/xlink href
	Href string `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`

	// generated from attribute "http://www.w3.org/1999/xlink type
	Type string `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`

	// generated from attribute "http://www.w3.org/1999/xlink role
	Role string `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`

	// generated from attribute "http://www.w3.org/1999/xlink title
	Title string `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`

	// generated from attribute "http://www.w3.org/1999/xlink show
	Show string `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`

	// generated from attribute "http://www.w3.org/1999/xlink actuate
	Actuate string `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

// AttributeGroup_measure_attributes is generated from named attribute group "measure-attributes"
type AttributeGroup_measure_attributes struct {

	// insertion point for fields

	// generated from attribute "number
	Number string `xml:"number,attr,omitempty"`

	// generated from attribute "text
	Text string `xml:"text,attr,omitempty"`

	// generated from attribute "implicit
	Implicit string `xml:"implicit,attr,omitempty"`

	// generated from attribute "non-controlling
	Non_controlling string `xml:"non-controlling,attr,omitempty"`

	// generated from attribute "width
	Width string `xml:"width,attr,omitempty"`

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id
}

// AttributeGroup_optional_unique_id is generated from named attribute group "optional-unique-id"
type AttributeGroup_optional_unique_id struct {

	// insertion point for fields

	// generated from attribute "id
	Id string `xml:"id,attr,omitempty"`
}

// AttributeGroup_orientation is generated from named attribute group "orientation"
type AttributeGroup_orientation struct {

	// insertion point for fields

	// generated from attribute "orientation
	Orientation string `xml:"orientation,attr,omitempty"`
}

// AttributeGroup_part_attributes is generated from named attribute group "part-attributes"
type AttributeGroup_part_attributes struct {

	// insertion point for fields

	// generated from attribute "id
	Id string `xml:"id,attr,omitempty"`
}

// AttributeGroup_part_name_text is generated from named attribute group "part-name-text"
type AttributeGroup_part_name_text struct {

	// insertion point for fields

	// generated from attribute group "print-style
	AttributeGroup_print_style

	// generated from attribute group "print-object
	AttributeGroup_print_object

	// generated from attribute group "justify
	AttributeGroup_justify
}

// AttributeGroup_placement is generated from named attribute group "placement"
type AttributeGroup_placement struct {

	// insertion point for fields

	// generated from attribute "placement
	Placement string `xml:"placement,attr,omitempty"`
}

// AttributeGroup_position is generated from named attribute group "position"
type AttributeGroup_position struct {

	// insertion point for fields

	// generated from attribute "default-x
	Default_x string `xml:"default-x,attr,omitempty"`

	// generated from attribute "default-y
	Default_y string `xml:"default-y,attr,omitempty"`

	// generated from attribute "relative-x
	Relative_x string `xml:"relative-x,attr,omitempty"`

	// generated from attribute "relative-y
	Relative_y string `xml:"relative-y,attr,omitempty"`
}

// AttributeGroup_print_attributes is generated from named attribute group "print-attributes"
type AttributeGroup_print_attributes struct {

	// insertion point for fields

	// generated from attribute "staff-spacing
	Staff_spacing string `xml:"staff-spacing,attr,omitempty"`

	// generated from attribute "new-system
	New_system string `xml:"new-system,attr,omitempty"`

	// generated from attribute "new-page
	New_page string `xml:"new-page,attr,omitempty"`

	// generated from attribute "blank-page
	Blank_page int `xml:"blank-page,attr,omitempty"`

	// generated from attribute "page-number
	Page_number string `xml:"page-number,attr,omitempty"`
}

// AttributeGroup_print_object is generated from named attribute group "print-object"
type AttributeGroup_print_object struct {

	// insertion point for fields

	// generated from attribute "print-object
	Print_object string `xml:"print-object,attr,omitempty"`
}

// AttributeGroup_print_spacing is generated from named attribute group "print-spacing"
type AttributeGroup_print_spacing struct {

	// insertion point for fields

	// generated from attribute "print-spacing
	Print_spacing string `xml:"print-spacing,attr,omitempty"`
}

// AttributeGroup_print_style is generated from named attribute group "print-style"
type AttributeGroup_print_style struct {

	// insertion point for fields

	// generated from attribute group "position
	AttributeGroup_position

	// generated from attribute group "font
	AttributeGroup_font

	// generated from attribute group "color
	AttributeGroup_color
}

// AttributeGroup_print_style_align is generated from named attribute group "print-style-align"
type AttributeGroup_print_style_align struct {

	// insertion point for fields

	// generated from attribute group "print-style
	AttributeGroup_print_style

	// generated from attribute group "halign
	AttributeGroup_halign

	// generated from attribute group "valign
	AttributeGroup_valign
}

// AttributeGroup_printout is generated from named attribute group "printout"
type AttributeGroup_printout struct {

	// insertion point for fields

	// generated from attribute "print-dot
	Print_dot string `xml:"print-dot,attr,omitempty"`

	// generated from attribute "print-lyric
	Print_lyric string `xml:"print-lyric,attr,omitempty"`

	// generated from attribute group "print-object
	AttributeGroup_print_object

	// generated from attribute group "print-spacing
	AttributeGroup_print_spacing
}

// AttributeGroup_smufl is generated from named attribute group "smufl"
type AttributeGroup_smufl struct {

	// insertion point for fields

	// generated from attribute "smufl
	Smufl string `xml:"smufl,attr,omitempty"`
}

// AttributeGroup_symbol_formatting is generated from named attribute group "symbol-formatting"
type AttributeGroup_symbol_formatting struct {

	// insertion point for fields

	// generated from attribute group "justify
	AttributeGroup_justify

	// generated from attribute group "print-style-align
	AttributeGroup_print_style_align

	// generated from attribute group "text-decoration
	AttributeGroup_text_decoration

	// generated from attribute group "text-rotation
	AttributeGroup_text_rotation

	// generated from attribute group "letter-spacing
	AttributeGroup_letter_spacing

	// generated from attribute group "line-height
	AttributeGroup_line_height

	// generated from attribute group "text-direction
	AttributeGroup_text_direction

	// generated from attribute group "enclosure
	AttributeGroup_enclosure
}

// AttributeGroup_system_relation is generated from named attribute group "system-relation"
type AttributeGroup_system_relation struct {

	// insertion point for fields

	// generated from attribute "system
	System string `xml:"system,attr,omitempty"`
}

// AttributeGroup_text_decoration is generated from named attribute group "text-decoration"
type AttributeGroup_text_decoration struct {

	// insertion point for fields

	// generated from attribute "underline
	Underline int `xml:"underline,attr,omitempty"`

	// generated from attribute "overline
	Overline int `xml:"overline,attr,omitempty"`

	// generated from attribute "line-through
	Line_through int `xml:"line-through,attr,omitempty"`
}

// AttributeGroup_text_direction is generated from named attribute group "text-direction"
type AttributeGroup_text_direction struct {

	// insertion point for fields

	// generated from attribute "dir
	Dir string `xml:"dir,attr,omitempty"`
}

// AttributeGroup_text_formatting is generated from named attribute group "text-formatting"
type AttributeGroup_text_formatting struct {

	// insertion point for fields

	// generated from attribute "http://www.w3.org/XML/1998/namespace lang
	Lang string `xml:"http://www.w3.org/XML/1998/namespace lang,attr,omitempty"`

	// generated from attribute "http://www.w3.org/XML/1998/namespace space
	Space string `xml:"http://www.w3.org/XML/1998/namespace space,attr,omitempty"`

	// generated from attribute group "justify
	AttributeGroup_justify

	// generated from attribute group "print-style-align
	AttributeGroup_print_style_align

	// generated from attribute group "text-decoration
	AttributeGroup_text_decoration

	// generated from attribute group "text-rotation
	AttributeGroup_text_rotation

	// generated from attribute group "letter-spacing
	AttributeGroup_letter_spacing

	// generated from attribute group "line-height
	AttributeGroup_line_height

	// generated from attribute group "text-direction
	AttributeGroup_text_direction

	// generated from attribute group "enclosure
	AttributeGroup_enclosure
}

// AttributeGroup_text_rotation is generated from named attribute group "text-rotation"
type AttributeGroup_text_rotation struct {

	// insertion point for fields

	// generated from attribute "rotation
	Rotation string `xml:"rotation,attr,omitempty"`
}

// AttributeGroup_trill_sound is generated from named attribute group "trill-sound"
type AttributeGroup_trill_sound struct {

	// insertion point for fields

	// generated from attribute "start-note
	Start_note string `xml:"start-note,attr,omitempty"`

	// generated from attribute "trill-step
	Trill_step string `xml:"trill-step,attr,omitempty"`

	// generated from attribute "two-note-turn
	Two_note_turn string `xml:"two-note-turn,attr,omitempty"`

	// generated from attribute "accelerate
	Accelerate string `xml:"accelerate,attr,omitempty"`

	// generated from attribute "beats
	Beats string `xml:"beats,attr,omitempty"`

	// generated from attribute "second-beat
	Second_beat string `xml:"second-beat,attr,omitempty"`

	// generated from attribute "last-beat
	Last_beat string `xml:"last-beat,attr,omitempty"`
}

// AttributeGroup_valign is generated from named attribute group "valign"
type AttributeGroup_valign struct {

	// insertion point for fields

	// generated from attribute "valign
	Valign string `xml:"valign,attr,omitempty"`
}

// AttributeGroup_valign_image is generated from named attribute group "valign-image"
type AttributeGroup_valign_image struct {

	// insertion point for fields

	// generated from attribute "valign
	Valign string `xml:"valign,attr,omitempty"`
}

// AttributeGroup_x_position is generated from named attribute group "x-position"
type AttributeGroup_x_position struct {

	// insertion point for fields

	// generated from attribute "default-x
	Default_x string `xml:"default-x,attr,omitempty"`

	// generated from attribute "default-y
	Default_y string `xml:"default-y,attr,omitempty"`

	// generated from attribute "relative-x
	Relative_x string `xml:"relative-x,attr,omitempty"`

	// generated from attribute "relative-y
	Relative_y string `xml:"relative-y,attr,omitempty"`
}

// AttributeGroup_y_position is generated from named attribute group "y-position"
type AttributeGroup_y_position struct {

	// insertion point for fields

	// generated from attribute "default-x
	Default_x string `xml:"default-x,attr,omitempty"`

	// generated from attribute "default-y
	Default_y string `xml:"default-y,attr,omitempty"`

	// generated from attribute "relative-x
	Relative_x string `xml:"relative-x,attr,omitempty"`

	// generated from attribute "relative-y
	Relative_y string `xml:"relative-y,attr,omitempty"`
}

// Score_partwise is generated from element score-partwise within root schema
type Score_partwise struct {
	Name string `xml:"-"`

	// insertion point for fields

	// necessary since it is a root element
	XMLName xml.Name `xml:"score-partwise"`

	// generated from inline complex type
	A_score_partwise
}

// Score_timewise is generated from element score-timewise within root schema
type Score_timewise struct {
	Name string `xml:"-"`

	// insertion point for fields

	// necessary since it is a root element
	XMLName xml.Name `xml:"score-timewise"`

	// generated from inline complex type
	A_score_timewise
}
