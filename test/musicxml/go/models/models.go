// generated code - do not edit
package models

import "encoding/xml"

// to avoid compilation error if no xml element is generated
var _ xml.Attr

// A_directive is generated from outer element "directive"
type A_directive struct {

	// insertion point for fields

	// generated from attribute "" of type 
	Lang string `xml:",attr"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// A_measure_1 is generated from outer element "measure"
type A_measure_1 struct {

	// insertion point for fields

	// generated from attribute group "measure-attributes
	AttributeGroup_measure_attributes

	// generated from element "note" of type note
	Note []*Note `xml:"note"`

	// generated from element "backup" of type backup
	Backup []*Backup `xml:"backup"`

	// generated from element "forward" of type forward
	Forward []*Forward `xml:"forward"`

	// generated from element "direction" of type direction
	Direction []*Direction `xml:"direction"`

	// generated from element "attributes" of type attributes
	Attributes []*Attributes `xml:"attributes"`

	// generated from element "harmony" of type harmony
	Harmony []*Harmony `xml:"harmony"`

	// generated from element "figured-bass" of type figured-bass
	Figured_bass []*Figured_bass `xml:"figured-bass"`

	// generated from element "print" of type print
	Print []*Print `xml:"print"`

	// generated from element "sound" of type sound
	Sound []*Sound `xml:"sound"`

	// generated from element "listening" of type listening
	Listening []*Listening `xml:"listening"`

	// generated from element "barline" of type barline
	Barline []*Barline `xml:"barline"`

	// generated from element "grouping" of type grouping
	Grouping []*Grouping `xml:"grouping"`

	// generated from element "link" of type link
	Link []*Link `xml:"link"`

	// generated from element "bookmark" of type bookmark
	Bookmark []*Bookmark `xml:"bookmark"`
}

// A_measure is generated from outer element "measure"
type A_measure struct {

	// insertion point for fields

	// generated from attribute group "measure-attributes
	AttributeGroup_measure_attributes

	// generated from anonymous type within outer element "part" of type A_part
	Part []*A_part_1 `xml:"part"`
}

// A_part_1 is generated from outer element "part"
type A_part_1 struct {

	// insertion point for fields

	// generated from attribute group "part-attributes
	AttributeGroup_part_attributes

	// generated from element "note" of type note
	Note []*Note `xml:"note"`

	// generated from element "backup" of type backup
	Backup []*Backup `xml:"backup"`

	// generated from element "forward" of type forward
	Forward []*Forward `xml:"forward"`

	// generated from element "direction" of type direction
	Direction []*Direction `xml:"direction"`

	// generated from element "attributes" of type attributes
	Attributes []*Attributes `xml:"attributes"`

	// generated from element "harmony" of type harmony
	Harmony []*Harmony `xml:"harmony"`

	// generated from element "figured-bass" of type figured-bass
	Figured_bass []*Figured_bass `xml:"figured-bass"`

	// generated from element "print" of type print
	Print []*Print `xml:"print"`

	// generated from element "sound" of type sound
	Sound []*Sound `xml:"sound"`

	// generated from element "listening" of type listening
	Listening []*Listening `xml:"listening"`

	// generated from element "barline" of type barline
	Barline []*Barline `xml:"barline"`

	// generated from element "grouping" of type grouping
	Grouping []*Grouping `xml:"grouping"`

	// generated from element "link" of type link
	Link []*Link `xml:"link"`

	// generated from element "bookmark" of type bookmark
	Bookmark []*Bookmark `xml:"bookmark"`
}

// A_part is generated from outer element "part"
type A_part struct {

	// insertion point for fields

	// generated from attribute group "part-attributes
	AttributeGroup_part_attributes

	// generated from anonymous type within outer element "measure" of type A_measure
	Measure []*A_measure_1 `xml:"measure"`
}

// A_score_partwise is generated from outer element "score-partwise"
type A_score_partwise struct {

	// insertion point for fields

	// generated from attribute group "document-attributes
	AttributeGroup_document_attributes

	// generated from element "work" of type work
	Work []*Work `xml:"work"`

	// generated from element "movement-number" of type xs:string
	Movement_number string `xml:"movement-number"`

	// generated from element "movement-title" of type xs:string
	Movement_title string `xml:"movement-title"`

	// generated from element "identification" of type identification
	Identification []*Identification `xml:"identification"`

	// generated from element "defaults" of type defaults
	Defaults []*Defaults `xml:"defaults"`

	// generated from element "credit" of type credit
	Credit []*Credit `xml:"credit"`

	// generated from element "part-list" of type part-list
	Part_list []*Part_list `xml:"part-list"`

	// generated from anonymous type within outer element "part" of type A_part
	Part []*A_part `xml:"part"`
}

// A_score_timewise is generated from outer element "score-timewise"
type A_score_timewise struct {

	// insertion point for fields

	// generated from attribute group "document-attributes
	AttributeGroup_document_attributes

	// generated from element "work" of type work
	Work []*Work `xml:"work"`

	// generated from element "movement-number" of type xs:string
	Movement_number string `xml:"movement-number"`

	// generated from element "movement-title" of type xs:string
	Movement_title string `xml:"movement-title"`

	// generated from element "identification" of type identification
	Identification []*Identification `xml:"identification"`

	// generated from element "defaults" of type defaults
	Defaults []*Defaults `xml:"defaults"`

	// generated from element "credit" of type credit
	Credit []*Credit `xml:"credit"`

	// generated from element "part-list" of type part-list
	Part_list []*Part_list `xml:"part-list"`

	// generated from anonymous type within outer element "measure" of type A_measure
	Measure []*A_measure `xml:"measure"`
}

// Accidental is generated from named complex type "accidental"
type Accidental struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "cautionary" of type yes-no
	Cautionary string `xml:"cautionary,attr"`

	// generated from attribute "editorial" of type yes-no
	Editorial string `xml:"editorial,attr"`

	// generated from attribute "smufl" of type smufl-accidental-glyph-name
	Smufl string `xml:"smufl,attr"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Accidental_mark is generated from named complex type "accidental-mark"
type Accidental_mark struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "smufl" of type smufl-accidental-glyph-name
	Smufl string `xml:"smufl,attr"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Accidental_text is generated from named complex type "accidental-text"
type Accidental_text struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "smufl" of type smufl-accidental-glyph-name
	Smufl string `xml:"smufl,attr"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Accord is generated from named complex type "accord"
type Accord struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "string" of type string-number
	String int `xml:"string,attr"`

	// generated from element "tuning-step" of type step
	Tuning_step string `xml:"tuning-step"`

	// generated from element "tuning-alter" of type semitones
	Tuning_alter string `xml:"tuning-alter"`

	// generated from element "tuning-octave" of type octave
	Tuning_octave int `xml:"tuning-octave"`
}

// Accordion_registration is generated from named complex type "accordion-registration"
type Accordion_registration struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "print-style-align
	AttributeGroup_print_style_align

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from element "accordion-high" of type empty
	Accordion_high string `xml:"accordion-high"`

	// generated from element "accordion-middle" of type accordion-middle
	Accordion_middle int `xml:"accordion-middle"`

	// generated from element "accordion-low" of type empty
	Accordion_low string `xml:"accordion-low"`
}

// Appearance is generated from named complex type "appearance"
type Appearance struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "line-width" of type line-width
	Line_width []*Line_width `xml:"line-width"`

	// generated from element "note-size" of type note-size
	Note_size []*Note_size `xml:"note-size"`

	// generated from element "distance" of type distance
	Distance []*Distance `xml:"distance"`

	// generated from element "glyph" of type glyph
	Glyph []*Glyph `xml:"glyph"`

	// generated from element "other-appearance" of type other-appearance
	Other_appearance []*Other_appearance `xml:"other-appearance"`
}

// Arpeggiate is generated from named complex type "arpeggiate"
type Arpeggiate struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "number" of type number-level
	Number int `xml:"number,attr"`

	// generated from attribute "direction" of type up-down
	Direction string `xml:"direction,attr"`

	// generated from attribute "unbroken" of type yes-no
	Unbroken string `xml:"unbroken,attr"`

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

	// generated from element "arrow-direction" of type arrow-direction
	Arrow_direction string `xml:"arrow-direction"`

	// generated from element "arrow-style" of type arrow-style
	Arrow_style string `xml:"arrow-style"`

	// generated from element "arrowhead" of type empty
	Arrowhead string `xml:"arrowhead"`

	// generated from element "circular-arrow" of type circular-arrow
	Circular_arrow string `xml:"circular-arrow"`
}

// Articulations is generated from named complex type "articulations"
type Articulations struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from element "accent" of type empty-placement
	Accent []*Empty_placement `xml:"accent"`

	// generated from element "strong-accent" of type strong-accent
	Strong_accent []*Strong_accent `xml:"strong-accent"`

	// generated from element "staccato" of type empty-placement
	Staccato []*Empty_placement `xml:"staccato"`

	// generated from element "tenuto" of type empty-placement
	Tenuto []*Empty_placement `xml:"tenuto"`

	// generated from element "detached-legato" of type empty-placement
	Detached_legato []*Empty_placement `xml:"detached-legato"`

	// generated from element "staccatissimo" of type empty-placement
	Staccatissimo []*Empty_placement `xml:"staccatissimo"`

	// generated from element "spiccato" of type empty-placement
	Spiccato []*Empty_placement `xml:"spiccato"`

	// generated from element "scoop" of type empty-line
	Scoop []*Empty_line `xml:"scoop"`

	// generated from element "plop" of type empty-line
	Plop []*Empty_line `xml:"plop"`

	// generated from element "doit" of type empty-line
	Doit []*Empty_line `xml:"doit"`

	// generated from element "falloff" of type empty-line
	Falloff []*Empty_line `xml:"falloff"`

	// generated from element "breath-mark" of type breath-mark
	Breath_mark []*Breath_mark `xml:"breath-mark"`

	// generated from element "caesura" of type caesura
	Caesura []*Caesura `xml:"caesura"`

	// generated from element "stress" of type empty-placement
	Stress []*Empty_placement `xml:"stress"`

	// generated from element "unstress" of type empty-placement
	Unstress []*Empty_placement `xml:"unstress"`

	// generated from element "soft-accent" of type empty-placement
	Soft_accent []*Empty_placement `xml:"soft-accent"`

	// generated from element "other-articulation" of type other-placement-text
	Other_articulation []*Other_placement_text `xml:"other-articulation"`
}

// Assess is generated from named complex type "assess"
type Assess struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type" of type yes-no
	Type string `xml:"type,attr"`

	// generated from attribute "player" of type xs:IDREF
	Player string `xml:"player,attr"`

	// generated from attribute "time-only" of type time-only
	Time_only string `xml:"time-only,attr"`
}

// Attributes is generated from named complex type "attributes"
type Attributes struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "footnote" of type formatted-text
	Footnote []*Formatted_text `xml:"footnote"`

	// generated from element "level" of type level
	Level []*Level `xml:"level"`

	// generated from element "transpose" of type transpose
	Transpose []*Transpose `xml:"transpose"`

	// generated from element "for-part" of type for-part
	For_part []*For_part `xml:"for-part"`

	// generated from element "divisions" of type positive-divisions
	Divisions string `xml:"divisions"`

	// generated from element "key" of type key
	Key []*Key `xml:"key"`

	// generated from element "time" of type time
	Time []*Time `xml:"time"`

	// generated from element "staves" of type xs:nonNegativeInteger
	Staves int `xml:"staves"`

	// generated from element "part-symbol" of type part-symbol
	Part_symbol []*Part_symbol `xml:"part-symbol"`

	// generated from element "instruments" of type xs:nonNegativeInteger
	Instruments int `xml:"instruments"`

	// generated from element "clef" of type clef
	Clef []*Clef `xml:"clef"`

	// generated from element "staff-details" of type staff-details
	Staff_details []*Staff_details `xml:"staff-details"`

	// generated from anonymous type within outer element "directive" of type A_directive
	Directive []*A_directive `xml:"directive"`

	// generated from element "measure-style" of type measure-style
	Measure_style []*Measure_style `xml:"measure-style"`
}

// Backup is generated from named complex type "backup"
type Backup struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "duration" of type positive-divisions
	Duration string `xml:"duration"`

	// generated from element "footnote" of type formatted-text
	Footnote []*Formatted_text `xml:"footnote"`

	// generated from element "level" of type level
	Level []*Level `xml:"level"`
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

	// generated from attribute "location" of type right-left-middle
	Location string `xml:"location,attr"`

	// generated from attribute "segno" of type xs:token
	Segno string `xml:"segno,attr"`

	// generated from attribute "coda" of type xs:token
	Coda string `xml:"coda,attr"`

	// generated from attribute "divisions" of type divisions
	Divisions string `xml:"divisions,attr"`

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from element "footnote" of type formatted-text
	Footnote []*Formatted_text `xml:"footnote"`

	// generated from element "level" of type level
	Level []*Level `xml:"level"`

	// generated from element "bar-style" of type bar-style-color
	Bar_style []*Bar_style_color `xml:"bar-style"`

	// generated from element "wavy-line" of type wavy-line
	Wavy_line []*Wavy_line `xml:"wavy-line"`

	// generated from element "segno" of type segno
	Segno_1 []*Segno `xml:"segno"`

	// generated from element "coda" of type coda
	Coda_1 []*Coda `xml:"coda"`

	// generated from element "fermata" of type fermata
	Fermata []*Fermata `xml:"fermata"`

	// generated from element "ending" of type ending
	Ending []*Ending `xml:"ending"`

	// generated from element "repeat" of type repeat
	Repeat []*Repeat `xml:"repeat"`
}

// Barre is generated from named complex type "barre"
type Barre struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type" of type start-stop
	Type string `xml:"type,attr"`

	// generated from attribute group "color
	AttributeGroup_color
}

// Bass is generated from named complex type "bass"
type Bass struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "arrangement" of type harmony-arrangement
	Arrangement string `xml:"arrangement,attr"`

	// generated from element "bass-separator" of type style-text
	Bass_separator []*Style_text `xml:"bass-separator"`

	// generated from element "bass-step" of type bass-step
	Bass_step []*Bass_step `xml:"bass-step"`

	// generated from element "bass-alter" of type harmony-alter
	Bass_alter []*Harmony_alter `xml:"bass-alter"`
}

// Bass_step is generated from named complex type "bass-step"
type Bass_step struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "text" of type xs:token
	Text string `xml:"text,attr"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Beam is generated from named complex type "beam"
type Beam struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "number" of type beam-level
	Number int `xml:"number,attr"`

	// generated from attribute "repeater" of type yes-no
	Repeater string `xml:"repeater,attr"`

	// generated from attribute "fan" of type fan
	Fan string `xml:"fan,attr"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Beat_repeat is generated from named complex type "beat-repeat"
type Beat_repeat struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type" of type start-stop
	Type string `xml:"type,attr"`

	// generated from attribute "slashes" of type xs:positiveInteger
	Slashes int `xml:"slashes,attr"`

	// generated from attribute "use-dots" of type yes-no
	Use_dots string `xml:"use-dots,attr"`

	// generated from element "slash-type" of type note-type-value
	Slash_type string `xml:"slash-type"`

	// generated from element "slash-dot" of type empty
	Slash_dot string `xml:"slash-dot"`

	// generated from element "except-voice" of type xs:string
	Except_voice string `xml:"except-voice"`
}

// Beat_unit_tied is generated from named complex type "beat-unit-tied"
type Beat_unit_tied struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "beat-unit" of type note-type-value
	Beat_unit string `xml:"beat-unit"`

	// generated from element "beat-unit-dot" of type empty
	Beat_unit_dot string `xml:"beat-unit-dot"`
}

// Beater is generated from named complex type "beater"
type Beater struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "tip" of type tip-direction
	Tip string `xml:"tip,attr"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Bend is generated from named complex type "bend"
type Bend struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "shape" of type bend-shape
	Shape string `xml:"shape,attr"`

	// generated from attribute group "print-style
	AttributeGroup_print_style

	// generated from attribute group "bend-sound
	AttributeGroup_bend_sound

	// generated from element "pre-bend" of type empty
	Pre_bend string `xml:"pre-bend"`

	// generated from element "release" of type release
	Release []*Release `xml:"release"`

	// generated from element "bend-alter" of type semitones
	Bend_alter string `xml:"bend-alter"`

	// generated from element "with-bar" of type placement-text
	With_bar []*Placement_text `xml:"with-bar"`
}

// Bookmark is generated from named complex type "bookmark"
type Bookmark struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id" of type xs:ID
	Id string `xml:"id,attr"`

	// generated from attribute "name" of type xs:token
	NameXSD string `xml:"name,attr"`

	// generated from attribute group "element-position
	AttributeGroup_element_position
}

// Bracket is generated from named complex type "bracket"
type Bracket struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type" of type start-stop-continue
	Type string `xml:"type,attr"`

	// generated from attribute "number" of type number-level
	Number int `xml:"number,attr"`

	// generated from attribute "line-end" of type line-end
	Line_end string `xml:"line-end,attr"`

	// generated from attribute "end-length" of type tenths
	End_length string `xml:"end-length,attr"`

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

	// generated from attribute "location" of type cancel-location
	Location string `xml:"location,attr"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText int `xml:",chardata"`
}

// Clef is generated from named complex type "clef"
type Clef struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "number" of type staff-number
	Number int `xml:"number,attr"`

	// generated from attribute "additional" of type yes-no
	Additional string `xml:"additional,attr"`

	// generated from attribute "size" of type symbol-size
	Size string `xml:"size,attr"`

	// generated from attribute "after-barline" of type yes-no
	After_barline string `xml:"after-barline,attr"`

	// generated from attribute group "print-style
	AttributeGroup_print_style

	// generated from attribute group "print-object
	AttributeGroup_print_object

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from element "sign" of type clef-sign
	Sign string `xml:"sign"`

	// generated from element "line" of type staff-line-position
	Line int `xml:"line"`

	// generated from element "clef-octave-change" of type xs:integer
	Clef_octave_change int `xml:"clef-octave-change"`
}

// Coda is generated from named complex type "coda"
type Coda struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "smufl" of type smufl-coda-glyph-name
	Smufl string `xml:"smufl,attr"`

	// generated from attribute group "print-style-align
	AttributeGroup_print_style_align

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id
}

// Credit is generated from named complex type "credit"
type Credit struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "page" of type xs:positiveInteger
	Page int `xml:"page,attr"`

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from element "credit-words" of type formatted-text-id
	Credit_words []*Formatted_text_id `xml:"credit-words"`

	// generated from element "credit-symbol" of type formatted-symbol-id
	Credit_symbol []*Formatted_symbol_id `xml:"credit-symbol"`

	// generated from element "link" of type link
	Link []*Link `xml:"link"`

	// generated from element "bookmark" of type bookmark
	Bookmark []*Bookmark `xml:"bookmark"`

	// generated from element "credit-image" of type image
	Credit_image []*Image `xml:"credit-image"`

	// generated from element "credit-type" of type xs:string
	Credit_type string `xml:"credit-type"`
}

// Dashes is generated from named complex type "dashes"
type Dashes struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type" of type start-stop-continue
	Type string `xml:"type,attr"`

	// generated from attribute "number" of type number-level
	Number int `xml:"number,attr"`

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

	// generated from element "page-layout" of type page-layout
	Page_layout []*Page_layout `xml:"page-layout"`

	// generated from element "system-layout" of type system-layout
	System_layout []*System_layout `xml:"system-layout"`

	// generated from element "staff-layout" of type staff-layout
	Staff_layout []*Staff_layout `xml:"staff-layout"`

	// generated from element "scaling" of type scaling
	Scaling []*Scaling `xml:"scaling"`

	// generated from element "concert-score" of type empty
	Concert_score string `xml:"concert-score"`

	// generated from element "appearance" of type appearance
	Appearance []*Appearance `xml:"appearance"`

	// generated from element "music-font" of type empty-font
	Music_font []*Empty_font `xml:"music-font"`

	// generated from element "word-font" of type empty-font
	Word_font []*Empty_font `xml:"word-font"`

	// generated from element "lyric-font" of type lyric-font
	Lyric_font []*Lyric_font `xml:"lyric-font"`

	// generated from element "lyric-language" of type lyric-language
	Lyric_language []*Lyric_language `xml:"lyric-language"`
}

// Degree is generated from named complex type "degree"
type Degree struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "print-object
	AttributeGroup_print_object

	// generated from element "degree-value" of type degree-value
	Degree_value []*Degree_value `xml:"degree-value"`

	// generated from element "degree-alter" of type degree-alter
	Degree_alter []*Degree_alter `xml:"degree-alter"`

	// generated from element "degree-type" of type degree-type
	Degree_type []*Degree_type `xml:"degree-type"`
}

// Degree_alter is generated from named complex type "degree-alter"
type Degree_alter struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "plus-minus" of type yes-no
	Plus_minus string `xml:"plus-minus,attr"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Degree_type is generated from named complex type "degree-type"
type Degree_type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "text" of type xs:token
	Text string `xml:"text,attr"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Degree_value is generated from named complex type "degree-value"
type Degree_value struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "symbol" of type degree-symbol-value
	Symbol string `xml:"symbol,attr"`

	// generated from attribute "text" of type xs:token
	Text string `xml:"text,attr"`

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

	// generated from element "footnote" of type formatted-text
	Footnote []*Formatted_text `xml:"footnote"`

	// generated from element "level" of type level
	Level []*Level `xml:"level"`

	// generated from element "voice" of type xs:string
	Voice string `xml:"voice"`

	// generated from element "staff" of type xs:positiveInteger
	Staff int `xml:"staff"`

	// generated from element "direction-type" of type direction-type
	Direction_type []*Direction_type `xml:"direction-type"`

	// generated from element "offset" of type offset
	Offset []*Offset `xml:"offset"`

	// generated from element "sound" of type sound
	Sound []*Sound `xml:"sound"`

	// generated from element "listening" of type listening
	Listening []*Listening `xml:"listening"`
}

// Direction_type is generated from named complex type "direction-type"
type Direction_type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from element "words" of type formatted-text-id
	Words []*Formatted_text_id `xml:"words"`

	// generated from element "symbol" of type formatted-symbol-id
	Symbol []*Formatted_symbol_id `xml:"symbol"`

	// generated from element "rehearsal" of type formatted-text-id
	Rehearsal []*Formatted_text_id `xml:"rehearsal"`

	// generated from element "segno" of type segno
	Segno []*Segno `xml:"segno"`

	// generated from element "coda" of type coda
	Coda []*Coda `xml:"coda"`

	// generated from element "wedge" of type wedge
	Wedge []*Wedge `xml:"wedge"`

	// generated from element "dynamics" of type dynamics
	Dynamics []*Dynamics `xml:"dynamics"`

	// generated from element "dashes" of type dashes
	Dashes []*Dashes `xml:"dashes"`

	// generated from element "bracket" of type bracket
	Bracket []*Bracket `xml:"bracket"`

	// generated from element "pedal" of type pedal
	Pedal []*Pedal `xml:"pedal"`

	// generated from element "metronome" of type metronome
	Metronome []*Metronome `xml:"metronome"`

	// generated from element "octave-shift" of type octave-shift
	Octave_shift []*Octave_shift `xml:"octave-shift"`

	// generated from element "harp-pedals" of type harp-pedals
	Harp_pedals []*Harp_pedals `xml:"harp-pedals"`

	// generated from element "damp" of type empty-print-style-align-id
	Damp []*Empty_print_style_align_id `xml:"damp"`

	// generated from element "damp-all" of type empty-print-style-align-id
	Damp_all []*Empty_print_style_align_id `xml:"damp-all"`

	// generated from element "eyeglasses" of type empty-print-style-align-id
	Eyeglasses []*Empty_print_style_align_id `xml:"eyeglasses"`

	// generated from element "string-mute" of type string-mute
	String_mute []*String_mute `xml:"string-mute"`

	// generated from element "scordatura" of type scordatura
	Scordatura []*Scordatura `xml:"scordatura"`

	// generated from element "image" of type image
	Image []*Image `xml:"image"`

	// generated from element "principal-voice" of type principal-voice
	Principal_voice []*Principal_voice `xml:"principal-voice"`

	// generated from element "percussion" of type percussion
	Percussion []*Percussion `xml:"percussion"`

	// generated from element "accordion-registration" of type accordion-registration
	Accordion_registration []*Accordion_registration `xml:"accordion-registration"`

	// generated from element "staff-divide" of type staff-divide
	Staff_divide []*Staff_divide `xml:"staff-divide"`

	// generated from element "other-direction" of type other-direction
	Other_direction []*Other_direction `xml:"other-direction"`
}

// Distance is generated from named complex type "distance"
type Distance struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type" of type distance-type
	Type string `xml:"type,attr"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Double is generated from named complex type "double"
type Double struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "above" of type yes-no
	Above string `xml:"above,attr"`
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

	// generated from element "p" of type empty
	P string `xml:"p"`

	// generated from element "pp" of type empty
	Pp string `xml:"pp"`

	// generated from element "ppp" of type empty
	Ppp string `xml:"ppp"`

	// generated from element "pppp" of type empty
	Pppp string `xml:"pppp"`

	// generated from element "ppppp" of type empty
	Ppppp string `xml:"ppppp"`

	// generated from element "pppppp" of type empty
	Pppppp string `xml:"pppppp"`

	// generated from element "f" of type empty
	F string `xml:"f"`

	// generated from element "ff" of type empty
	Ff string `xml:"ff"`

	// generated from element "fff" of type empty
	Fff string `xml:"fff"`

	// generated from element "ffff" of type empty
	Ffff string `xml:"ffff"`

	// generated from element "fffff" of type empty
	Fffff string `xml:"fffff"`

	// generated from element "ffffff" of type empty
	Ffffff string `xml:"ffffff"`

	// generated from element "mp" of type empty
	Mp string `xml:"mp"`

	// generated from element "mf" of type empty
	Mf string `xml:"mf"`

	// generated from element "sf" of type empty
	Sf string `xml:"sf"`

	// generated from element "sfp" of type empty
	Sfp string `xml:"sfp"`

	// generated from element "sfpp" of type empty
	Sfpp string `xml:"sfpp"`

	// generated from element "fp" of type empty
	Fp string `xml:"fp"`

	// generated from element "rf" of type empty
	Rf string `xml:"rf"`

	// generated from element "rfz" of type empty
	Rfz string `xml:"rfz"`

	// generated from element "sfz" of type empty
	Sfz string `xml:"sfz"`

	// generated from element "sffz" of type empty
	Sffz string `xml:"sffz"`

	// generated from element "fz" of type empty
	Fz string `xml:"fz"`

	// generated from element "n" of type empty
	N string `xml:"n"`

	// generated from element "pf" of type empty
	Pf string `xml:"pf"`

	// generated from element "sfzp" of type empty
	Sfzp string `xml:"sfzp"`

	// generated from element "other-dynamics" of type other-text
	Other_dynamics []*Other_text `xml:"other-dynamics"`
}

// Effect is generated from named complex type "effect"
type Effect struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "smufl" of type smufl-pictogram-glyph-name
	Smufl string `xml:"smufl,attr"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Elision is generated from named complex type "elision"
type Elision struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "smufl" of type smufl-lyrics-glyph-name
	Smufl string `xml:"smufl,attr"`

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

	// generated from element "encoder" of type typed-text
	Encoder []*Typed_text `xml:"encoder"`

	// generated from element "software" of type xs:string
	Software string `xml:"software"`

	// generated from element "encoding-description" of type xs:string
	Encoding_description string `xml:"encoding-description"`

	// generated from element "supports" of type supports
	Supports []*Supports `xml:"supports"`
}

// Ending is generated from named complex type "ending"
type Ending struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "number" of type ending-number
	Number string `xml:"number,attr"`

	// generated from attribute "type" of type start-stop-discontinue
	Type string `xml:"type,attr"`

	// generated from attribute "end-length" of type tenths
	End_length string `xml:"end-length,attr"`

	// generated from attribute "text-x" of type tenths
	Text_x string `xml:"text-x,attr"`

	// generated from attribute "text-y" of type tenths
	Text_y string `xml:"text-y,attr"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Extend is generated from named complex type "extend"
type Extend struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type" of type start-stop-continue
	Type string `xml:"type,attr"`

	// generated from attribute group "position
	AttributeGroup_position

	// generated from attribute group "color
	AttributeGroup_color
}

// Feature is generated from named complex type "feature"
type Feature struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type" of type xs:token
	Type string `xml:"type,attr"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Fermata is generated from named complex type "fermata"
type Fermata struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type" of type upright-inverted
	Type string `xml:"type,attr"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Figure is generated from named complex type "figure"
type Figure struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "footnote" of type formatted-text
	Footnote []*Formatted_text `xml:"footnote"`

	// generated from element "level" of type level
	Level []*Level `xml:"level"`

	// generated from element "prefix" of type style-text
	Prefix []*Style_text `xml:"prefix"`

	// generated from element "figure-number" of type style-text
	Figure_number []*Style_text `xml:"figure-number"`

	// generated from element "suffix" of type style-text
	Suffix []*Style_text `xml:"suffix"`

	// generated from element "extend" of type extend
	Extend []*Extend `xml:"extend"`
}

// Figured_bass is generated from named complex type "figured-bass"
type Figured_bass struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "parentheses" of type yes-no
	Parentheses string `xml:"parentheses,attr"`

	// generated from attribute group "print-style-align
	AttributeGroup_print_style_align

	// generated from attribute group "placement
	AttributeGroup_placement

	// generated from attribute group "printout
	AttributeGroup_printout

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from element "duration" of type positive-divisions
	Duration string `xml:"duration"`

	// generated from element "footnote" of type formatted-text
	Footnote []*Formatted_text `xml:"footnote"`

	// generated from element "level" of type level
	Level []*Level `xml:"level"`

	// generated from element "figure" of type figure
	Figure []*Figure `xml:"figure"`
}

// Fingering is generated from named complex type "fingering"
type Fingering struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "substitution" of type yes-no
	Substitution string `xml:"substitution,attr"`

	// generated from attribute "alternate" of type yes-no
	Alternate string `xml:"alternate,attr"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// First_fret is generated from named complex type "first-fret"
type First_fret struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "text" of type xs:token
	Text string `xml:"text,attr"`

	// generated from attribute "location" of type left-right
	Location string `xml:"location,attr"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText int `xml:",chardata"`
}

// For_part is generated from named complex type "for-part"
type For_part struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "number" of type staff-number
	Number int `xml:"number,attr"`

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from element "part-clef" of type part-clef
	Part_clef []*Part_clef `xml:"part-clef"`

	// generated from element "part-transpose" of type part-transpose
	Part_transpose []*Part_transpose `xml:"part-transpose"`
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

	// generated from element "duration" of type positive-divisions
	Duration string `xml:"duration"`

	// generated from element "footnote" of type formatted-text
	Footnote []*Formatted_text `xml:"footnote"`

	// generated from element "level" of type level
	Level []*Level `xml:"level"`

	// generated from element "voice" of type xs:string
	Voice string `xml:"voice"`

	// generated from element "staff" of type xs:positiveInteger
	Staff int `xml:"staff"`
}

// Frame is generated from named complex type "frame"
type Frame struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "height" of type tenths
	Height string `xml:"height,attr"`

	// generated from attribute "width" of type tenths
	Width string `xml:"width,attr"`

	// generated from attribute "unplayed" of type xs:token
	Unplayed string `xml:"unplayed,attr"`

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

	// generated from element "frame-strings" of type xs:positiveInteger
	Frame_strings int `xml:"frame-strings"`

	// generated from element "frame-frets" of type xs:positiveInteger
	Frame_frets int `xml:"frame-frets"`

	// generated from element "first-fret" of type first-fret
	First_fret []*First_fret `xml:"first-fret"`

	// generated from element "frame-note" of type frame-note
	Frame_note []*Frame_note `xml:"frame-note"`
}

// Frame_note is generated from named complex type "frame-note"
type Frame_note struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "string" of type string-type
	String []*String_type `xml:"string"`

	// generated from element "fret" of type fret
	Fret []*Fret `xml:"fret"`

	// generated from element "fingering" of type fingering
	Fingering []*Fingering `xml:"fingering"`

	// generated from element "barre" of type barre
	Barre []*Barre `xml:"barre"`
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

	// generated from attribute "smufl" of type smufl-pictogram-glyph-name
	Smufl string `xml:"smufl,attr"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Glissando is generated from named complex type "glissando"
type Glissando struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type" of type start-stop
	Type string `xml:"type,attr"`

	// generated from attribute "number" of type number-level
	Number int `xml:"number,attr"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Glyph is generated from named complex type "glyph"
type Glyph struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type" of type glyph-type
	Type string `xml:"type,attr"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Grace is generated from named complex type "grace"
type Grace struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "steal-time-previous" of type percent
	Steal_time_previous string `xml:"steal-time-previous,attr"`

	// generated from attribute "steal-time-following" of type percent
	Steal_time_following string `xml:"steal-time-following,attr"`

	// generated from attribute "make-time" of type divisions
	Make_time string `xml:"make-time,attr"`

	// generated from attribute "slash" of type yes-no
	Slash string `xml:"slash,attr"`
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

	// generated from attribute "type" of type start-stop-single
	Type string `xml:"type,attr"`

	// generated from attribute "number" of type xs:token
	Number string `xml:"number,attr"`

	// generated from attribute "member-of" of type xs:token
	Member_of string `xml:"member-of,attr"`

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from element "feature" of type feature
	Feature []*Feature `xml:"feature"`
}

// Hammer_on_pull_off is generated from named complex type "hammer-on-pull-off"
type Hammer_on_pull_off struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type" of type start-stop
	Type string `xml:"type,attr"`

	// generated from attribute "number" of type number-level
	Number int `xml:"number,attr"`

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

	// generated from attribute "location" of type harmon-closed-location
	Location string `xml:"location,attr"`

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

	// generated from element "harmon-closed" of type harmon-closed
	Harmon_closed []*Harmon_closed `xml:"harmon-closed"`
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

	// generated from element "natural" of type empty
	Natural string `xml:"natural"`

	// generated from element "artificial" of type empty
	Artificial string `xml:"artificial"`

	// generated from element "base-pitch" of type empty
	Base_pitch string `xml:"base-pitch"`

	// generated from element "touching-pitch" of type empty
	Touching_pitch string `xml:"touching-pitch"`

	// generated from element "sounding-pitch" of type empty
	Sounding_pitch string `xml:"sounding-pitch"`
}

// Harmony is generated from named complex type "harmony"
type Harmony struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type" of type harmony-type
	Type string `xml:"type,attr"`

	// generated from attribute "print-frame" of type yes-no
	Print_frame string `xml:"print-frame,attr"`

	// generated from attribute "arrangement" of type harmony-arrangement
	Arrangement string `xml:"arrangement,attr"`

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

	// generated from element "root" of type root
	Root []*Root `xml:"root"`

	// generated from element "numeral" of type numeral
	Numeral []*Numeral `xml:"numeral"`

	// generated from element "function" of type style-text
	Function []*Style_text `xml:"function"`

	// generated from element "kind" of type kind
	Kind []*Kind `xml:"kind"`

	// generated from element "inversion" of type inversion
	Inversion []*Inversion `xml:"inversion"`

	// generated from element "bass" of type bass
	Bass []*Bass `xml:"bass"`

	// generated from element "degree" of type degree
	Degree []*Degree `xml:"degree"`

	// generated from element "footnote" of type formatted-text
	Footnote []*Formatted_text `xml:"footnote"`

	// generated from element "level" of type level
	Level []*Level `xml:"level"`

	// generated from element "staff" of type xs:positiveInteger
	Staff int `xml:"staff"`

	// generated from element "frame" of type frame
	Frame []*Frame `xml:"frame"`

	// generated from element "offset" of type offset
	Offset []*Offset `xml:"offset"`
}

// Harmony_alter is generated from named complex type "harmony-alter"
type Harmony_alter struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "location" of type left-right
	Location string `xml:"location,attr"`

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

	// generated from element "pedal-tuning" of type pedal-tuning
	Pedal_tuning []*Pedal_tuning `xml:"pedal-tuning"`
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

	// generated from element "hole-type" of type xs:string
	Hole_type string `xml:"hole-type"`

	// generated from element "hole-closed" of type hole-closed
	Hole_closed []*Hole_closed `xml:"hole-closed"`

	// generated from element "hole-shape" of type xs:string
	Hole_shape string `xml:"hole-shape"`
}

// Hole_closed is generated from named complex type "hole-closed"
type Hole_closed struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "location" of type hole-closed-location
	Location string `xml:"location,attr"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Horizontal_turn is generated from named complex type "horizontal-turn"
type Horizontal_turn struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "slash" of type yes-no
	Slash string `xml:"slash,attr"`

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

	// generated from element "creator" of type typed-text
	Creator []*Typed_text `xml:"creator"`

	// generated from element "rights" of type typed-text
	Rights []*Typed_text `xml:"rights"`

	// generated from element "encoding" of type encoding
	Encoding []*Encoding `xml:"encoding"`

	// generated from element "source" of type xs:string
	Source string `xml:"source"`

	// generated from element "relation" of type typed-text
	Relation []*Typed_text `xml:"relation"`

	// generated from element "miscellaneous" of type miscellaneous
	Miscellaneous []*Miscellaneous `xml:"miscellaneous"`
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

	// generated from attribute "id" of type xs:IDREF
	Id string `xml:"id,attr"`
}

// Instrument_change is generated from named complex type "instrument-change"
type Instrument_change struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id" of type xs:IDREF
	Id string `xml:"id,attr"`

	// generated from element "solo" of type empty
	Solo string `xml:"solo"`

	// generated from element "ensemble" of type positive-integer-or-empty
	Ensemble string `xml:"ensemble"`

	// generated from element "instrument-sound" of type xs:string
	Instrument_sound string `xml:"instrument-sound"`

	// generated from element "virtual-instrument" of type virtual-instrument
	Virtual_instrument []*Virtual_instrument `xml:"virtual-instrument"`
}

// Instrument_link is generated from named complex type "instrument-link"
type Instrument_link struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id" of type xs:IDREF
	Id string `xml:"id,attr"`
}

// Interchangeable is generated from named complex type "interchangeable"
type Interchangeable struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "symbol" of type time-symbol
	Symbol string `xml:"symbol,attr"`

	// generated from attribute "separator" of type time-separator
	Separator string `xml:"separator,attr"`

	// generated from element "beats" of type xs:string
	Beats string `xml:"beats"`

	// generated from element "beat-type" of type xs:string
	Beat_type string `xml:"beat-type"`

	// generated from element "time-relation" of type time-relation
	Time_relation string `xml:"time-relation"`
}

// Inversion is generated from named complex type "inversion"
type Inversion struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "text" of type xs:token
	Text string `xml:"text,attr"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText int `xml:",chardata"`
}

// Key is generated from named complex type "key"
type Key struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "number" of type staff-number
	Number int `xml:"number,attr"`

	// generated from attribute group "print-style
	AttributeGroup_print_style

	// generated from attribute group "print-object
	AttributeGroup_print_object

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from element "cancel" of type cancel
	Cancel []*Cancel `xml:"cancel"`

	// generated from element "fifths" of type fifths
	Fifths int `xml:"fifths"`

	// generated from element "mode" of type mode
	Mode string `xml:"mode"`

	// generated from element "key-step" of type step
	Key_step string `xml:"key-step"`

	// generated from element "key-alter" of type semitones
	Key_alter string `xml:"key-alter"`

	// generated from element "key-accidental" of type key-accidental
	Key_accidental []*Key_accidental `xml:"key-accidental"`

	// generated from element "key-octave" of type key-octave
	Key_octave []*Key_octave `xml:"key-octave"`
}

// Key_accidental is generated from named complex type "key-accidental"
type Key_accidental struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "smufl" of type smufl-accidental-glyph-name
	Smufl string `xml:"smufl,attr"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Key_octave is generated from named complex type "key-octave"
type Key_octave struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "number" of type xs:positiveInteger
	Number int `xml:"number,attr"`

	// generated from attribute "cancel" of type yes-no
	Cancel string `xml:"cancel,attr"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText int `xml:",chardata"`
}

// Kind is generated from named complex type "kind"
type Kind struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "use-symbols" of type yes-no
	Use_symbols string `xml:"use-symbols,attr"`

	// generated from attribute "text" of type xs:token
	Text string `xml:"text,attr"`

	// generated from attribute "stack-degrees" of type yes-no
	Stack_degrees string `xml:"stack-degrees,attr"`

	// generated from attribute "parentheses-degrees" of type yes-no
	Parentheses_degrees string `xml:"parentheses-degrees,attr"`

	// generated from attribute "bracket-degrees" of type yes-no
	Bracket_degrees string `xml:"bracket-degrees,attr"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Level is generated from named complex type "level"
type Level struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "reference" of type yes-no
	Reference string `xml:"reference,attr"`

	// generated from attribute "type" of type start-stop-single
	Type string `xml:"type,attr"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Line_detail is generated from named complex type "line-detail"
type Line_detail struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "line" of type staff-line
	Line int `xml:"line,attr"`

	// generated from attribute "width" of type tenths
	Width string `xml:"width,attr"`

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

	// generated from attribute "type" of type line-width-type
	Type string `xml:"type,attr"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Link is generated from named complex type "link"
type Link struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "name" of type xs:token
	NameXSD string `xml:"name,attr"`

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

	// generated from element "assess" of type assess
	Assess []*Assess `xml:"assess"`

	// generated from element "wait" of type wait
	Wait []*Wait `xml:"wait"`

	// generated from element "other-listen" of type other-listening
	Other_listen []*Other_listening `xml:"other-listen"`
}

// Listening is generated from named complex type "listening"
type Listening struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "sync" of type sync
	Sync []*Sync `xml:"sync"`

	// generated from element "other-listening" of type other-listening
	Other_listening []*Other_listening `xml:"other-listening"`

	// generated from element "offset" of type offset
	Offset []*Offset `xml:"offset"`
}

// Lyric is generated from named complex type "lyric"
type Lyric struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "number" of type xs:NMTOKEN
	Number string `xml:"number,attr"`

	// generated from attribute "name" of type xs:token
	NameXSD string `xml:"name,attr"`

	// generated from attribute "time-only" of type time-only
	Time_only string `xml:"time-only,attr"`

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

	// generated from element "footnote" of type formatted-text
	Footnote []*Formatted_text `xml:"footnote"`

	// generated from element "level" of type level
	Level []*Level `xml:"level"`

	// generated from element "elision" of type elision
	Elision []*Elision `xml:"elision"`

	// generated from element "syllabic" of type syllabic
	Syllabic string `xml:"syllabic"`

	// generated from element "text" of type text-element-data
	Text []*Text_element_data `xml:"text"`

	// generated from element "extend" of type extend
	Extend []*Extend `xml:"extend"`

	// generated from element "laughing" of type empty
	Laughing string `xml:"laughing"`

	// generated from element "humming" of type empty
	Humming string `xml:"humming"`

	// generated from element "end-line" of type empty
	End_line string `xml:"end-line"`

	// generated from element "end-paragraph" of type empty
	End_paragraph string `xml:"end-paragraph"`
}

// Lyric_font is generated from named complex type "lyric-font"
type Lyric_font struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "number" of type xs:NMTOKEN
	Number string `xml:"number,attr"`

	// generated from attribute "name" of type xs:token
	NameXSD string `xml:"name,attr"`

	// generated from attribute group "font
	AttributeGroup_font
}

// Lyric_language is generated from named complex type "lyric-language"
type Lyric_language struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "number" of type xs:NMTOKEN
	Number string `xml:"number,attr"`

	// generated from attribute "name" of type xs:token
	NameXSD string `xml:"name,attr"`

	// generated from attribute "" of type 
	Lang string `xml:",attr"`
}

// Measure_layout is generated from named complex type "measure-layout"
type Measure_layout struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "measure-distance" of type tenths
	Measure_distance string `xml:"measure-distance"`
}

// Measure_numbering is generated from named complex type "measure-numbering"
type Measure_numbering struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "system" of type system-relation-number
	System string `xml:"system,attr"`

	// generated from attribute "staff" of type staff-number
	Staff int `xml:"staff,attr"`

	// generated from attribute "multiple-rest-always" of type yes-no
	Multiple_rest_always string `xml:"multiple-rest-always,attr"`

	// generated from attribute "multiple-rest-range" of type yes-no
	Multiple_rest_range string `xml:"multiple-rest-range,attr"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Measure_repeat is generated from named complex type "measure-repeat"
type Measure_repeat struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type" of type start-stop
	Type string `xml:"type,attr"`

	// generated from attribute "slashes" of type xs:positiveInteger
	Slashes int `xml:"slashes,attr"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Measure_style is generated from named complex type "measure-style"
type Measure_style struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "number" of type staff-number
	Number int `xml:"number,attr"`

	// generated from attribute group "font
	AttributeGroup_font

	// generated from attribute group "color
	AttributeGroup_color

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from element "multiple-rest" of type multiple-rest
	Multiple_rest []*Multiple_rest `xml:"multiple-rest"`

	// generated from element "measure-repeat" of type measure-repeat
	Measure_repeat []*Measure_repeat `xml:"measure-repeat"`

	// generated from element "beat-repeat" of type beat-repeat
	Beat_repeat []*Beat_repeat `xml:"beat-repeat"`

	// generated from element "slash" of type slash
	Slash []*Slash `xml:"slash"`
}

// Membrane is generated from named complex type "membrane"
type Membrane struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "smufl" of type smufl-pictogram-glyph-name
	Smufl string `xml:"smufl,attr"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Metal is generated from named complex type "metal"
type Metal struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "smufl" of type smufl-pictogram-glyph-name
	Smufl string `xml:"smufl,attr"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Metronome is generated from named complex type "metronome"
type Metronome struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "parentheses" of type yes-no
	Parentheses string `xml:"parentheses,attr"`

	// generated from attribute group "print-style-align
	AttributeGroup_print_style_align

	// generated from attribute group "print-object
	AttributeGroup_print_object

	// generated from attribute group "justify
	AttributeGroup_justify

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from element "beat-unit" of type note-type-value
	Beat_unit string `xml:"beat-unit"`

	// generated from element "beat-unit-dot" of type empty
	Beat_unit_dot string `xml:"beat-unit-dot"`

	// generated from element "beat-unit-tied" of type beat-unit-tied
	Beat_unit_tied []*Beat_unit_tied `xml:"beat-unit-tied"`

	// generated from element "per-minute" of type per-minute
	Per_minute []*Per_minute `xml:"per-minute"`

	// generated from element "metronome-relation" of type xs:string
	Metronome_relation string `xml:"metronome-relation"`

	// generated from element "metronome-note" of type metronome-note
	Metronome_note []*Metronome_note `xml:"metronome-note"`

	// generated from element "metronome-arrows" of type empty
	Metronome_arrows string `xml:"metronome-arrows"`
}

// Metronome_beam is generated from named complex type "metronome-beam"
type Metronome_beam struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "number" of type beam-level
	Number int `xml:"number,attr"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Metronome_note is generated from named complex type "metronome-note"
type Metronome_note struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "metronome-type" of type note-type-value
	Metronome_type string `xml:"metronome-type"`

	// generated from element "metronome-dot" of type empty
	Metronome_dot string `xml:"metronome-dot"`

	// generated from element "metronome-beam" of type metronome-beam
	Metronome_beam []*Metronome_beam `xml:"metronome-beam"`

	// generated from element "metronome-tied" of type metronome-tied
	Metronome_tied []*Metronome_tied `xml:"metronome-tied"`

	// generated from element "metronome-tuplet" of type metronome-tuplet
	Metronome_tuplet []*Metronome_tuplet `xml:"metronome-tuplet"`
}

// Metronome_tied is generated from named complex type "metronome-tied"
type Metronome_tied struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type" of type start-stop
	Type string `xml:"type,attr"`
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

	// generated from attribute "port" of type midi-16
	Port int `xml:"port,attr"`

	// generated from attribute "id" of type xs:IDREF
	Id string `xml:"id,attr"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Midi_instrument is generated from named complex type "midi-instrument"
type Midi_instrument struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id" of type xs:IDREF
	Id string `xml:"id,attr"`

	// generated from element "midi-channel" of type midi-16
	Midi_channel int `xml:"midi-channel"`

	// generated from element "midi-name" of type xs:string
	Midi_name string `xml:"midi-name"`

	// generated from element "midi-bank" of type midi-16384
	Midi_bank int `xml:"midi-bank"`

	// generated from element "midi-program" of type midi-128
	Midi_program int `xml:"midi-program"`

	// generated from element "midi-unpitched" of type midi-128
	Midi_unpitched int `xml:"midi-unpitched"`

	// generated from element "volume" of type percent
	Volume string `xml:"volume"`

	// generated from element "pan" of type rotation-degrees
	Pan string `xml:"pan"`

	// generated from element "elevation" of type rotation-degrees
	Elevation string `xml:"elevation"`
}

// Miscellaneous is generated from named complex type "miscellaneous"
type Miscellaneous struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "miscellaneous-field" of type miscellaneous-field
	Miscellaneous_field []*Miscellaneous_field `xml:"miscellaneous-field"`
}

// Miscellaneous_field is generated from named complex type "miscellaneous-field"
type Miscellaneous_field struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "name" of type xs:token
	NameXSD string `xml:"name,attr"`

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

	// generated from attribute "use-symbols" of type yes-no
	Use_symbols string `xml:"use-symbols,attr"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText int `xml:",chardata"`
}

// Name_display is generated from named complex type "name-display"
type Name_display struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "print-object
	AttributeGroup_print_object

	// generated from element "display-text" of type formatted-text
	Display_text []*Formatted_text `xml:"display-text"`

	// generated from element "accidental-text" of type accidental-text
	Accidental_text []*Accidental_text `xml:"accidental-text"`
}

// Non_arpeggiate is generated from named complex type "non-arpeggiate"
type Non_arpeggiate struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type" of type top-bottom
	Type string `xml:"type,attr"`

	// generated from attribute "number" of type number-level
	Number int `xml:"number,attr"`

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

	// generated from element "footnote" of type formatted-text
	Footnote []*Formatted_text `xml:"footnote"`

	// generated from element "level" of type level
	Level []*Level `xml:"level"`

	// generated from element "tied" of type tied
	Tied []*Tied `xml:"tied"`

	// generated from element "slur" of type slur
	Slur []*Slur `xml:"slur"`

	// generated from element "tuplet" of type tuplet
	Tuplet []*Tuplet `xml:"tuplet"`

	// generated from element "glissando" of type glissando
	Glissando []*Glissando `xml:"glissando"`

	// generated from element "slide" of type slide
	Slide []*Slide `xml:"slide"`

	// generated from element "ornaments" of type ornaments
	Ornaments []*Ornaments `xml:"ornaments"`

	// generated from element "technical" of type technical
	Technical []*Technical `xml:"technical"`

	// generated from element "articulations" of type articulations
	Articulations []*Articulations `xml:"articulations"`

	// generated from element "dynamics" of type dynamics
	Dynamics []*Dynamics `xml:"dynamics"`

	// generated from element "fermata" of type fermata
	Fermata []*Fermata `xml:"fermata"`

	// generated from element "arpeggiate" of type arpeggiate
	Arpeggiate []*Arpeggiate `xml:"arpeggiate"`

	// generated from element "non-arpeggiate" of type non-arpeggiate
	Non_arpeggiate []*Non_arpeggiate `xml:"non-arpeggiate"`

	// generated from element "accidental-mark" of type accidental-mark
	Accidental_mark []*Accidental_mark `xml:"accidental-mark"`

	// generated from element "other-notation" of type other-notation
	Other_notation []*Other_notation `xml:"other-notation"`
}

// Note is generated from named complex type "note"
type Note struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "print-leger" of type yes-no
	Print_leger string `xml:"print-leger,attr"`

	// generated from attribute "dynamics" of type non-negative-decimal
	Dynamics string `xml:"dynamics,attr"`

	// generated from attribute "end-dynamics" of type non-negative-decimal
	End_dynamics string `xml:"end-dynamics,attr"`

	// generated from attribute "attack" of type divisions
	Attack string `xml:"attack,attr"`

	// generated from attribute "release" of type divisions
	Release string `xml:"release,attr"`

	// generated from attribute "time-only" of type time-only
	Time_only string `xml:"time-only,attr"`

	// generated from attribute "pizzicato" of type yes-no
	Pizzicato string `xml:"pizzicato,attr"`

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

	// generated from element "footnote" of type formatted-text
	Footnote []*Formatted_text `xml:"footnote"`

	// generated from element "level" of type level
	Level []*Level `xml:"level"`

	// generated from element "voice" of type xs:string
	Voice string `xml:"voice"`

	// generated from element "staff" of type xs:positiveInteger
	Staff int `xml:"staff"`

	// generated from element "pitch" of type pitch
	Pitch []*Pitch `xml:"pitch"`

	// generated from element "unpitched" of type unpitched
	Unpitched []*Unpitched `xml:"unpitched"`

	// generated from element "rest" of type rest
	Rest []*Rest `xml:"rest"`

	// generated from element "chord" of type empty
	Chord string `xml:"chord"`

	// generated from element "tie" of type tie
	Tie []*Tie `xml:"tie"`

	// generated from element "cue" of type empty
	Cue string `xml:"cue"`

	// generated from element "grace" of type grace
	Grace []*Grace `xml:"grace"`

	// generated from element "duration" of type positive-divisions
	Duration string `xml:"duration"`

	// generated from element "instrument" of type instrument
	Instrument []*Instrument `xml:"instrument"`

	// generated from element "type" of type note-type
	Type []*Note_type `xml:"type"`

	// generated from element "dot" of type empty-placement
	Dot []*Empty_placement `xml:"dot"`

	// generated from element "accidental" of type accidental
	Accidental []*Accidental `xml:"accidental"`

	// generated from element "time-modification" of type time-modification
	Time_modification []*Time_modification `xml:"time-modification"`

	// generated from element "stem" of type stem
	Stem []*Stem `xml:"stem"`

	// generated from element "notehead" of type notehead
	Notehead []*Notehead `xml:"notehead"`

	// generated from element "notehead-text" of type notehead-text
	Notehead_text []*Notehead_text `xml:"notehead-text"`

	// generated from element "beam" of type beam
	Beam []*Beam `xml:"beam"`

	// generated from element "notations" of type notations
	Notations []*Notations `xml:"notations"`

	// generated from element "lyric" of type lyric
	Lyric []*Lyric `xml:"lyric"`

	// generated from element "play" of type play
	Play []*Play `xml:"play"`

	// generated from element "listen" of type listen
	Listen []*Listen `xml:"listen"`
}

// Note_size is generated from named complex type "note-size"
type Note_size struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type" of type note-size-type
	Type string `xml:"type,attr"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Note_type is generated from named complex type "note-type"
type Note_type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "size" of type symbol-size
	Size string `xml:"size,attr"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Notehead is generated from named complex type "notehead"
type Notehead struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "filled" of type yes-no
	Filled string `xml:"filled,attr"`

	// generated from attribute "parentheses" of type yes-no
	Parentheses string `xml:"parentheses,attr"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Notehead_text is generated from named complex type "notehead-text"
type Notehead_text struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "display-text" of type formatted-text
	Display_text []*Formatted_text `xml:"display-text"`

	// generated from element "accidental-text" of type accidental-text
	Accidental_text []*Accidental_text `xml:"accidental-text"`
}

// Numeral is generated from named complex type "numeral"
type Numeral struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "numeral-root" of type numeral-root
	Numeral_root []*Numeral_root `xml:"numeral-root"`

	// generated from element "numeral-alter" of type harmony-alter
	Numeral_alter []*Harmony_alter `xml:"numeral-alter"`

	// generated from element "numeral-key" of type numeral-key
	Numeral_key []*Numeral_key `xml:"numeral-key"`
}

// Numeral_key is generated from named complex type "numeral-key"
type Numeral_key struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "print-object
	AttributeGroup_print_object

	// generated from element "numeral-fifths" of type fifths
	Numeral_fifths int `xml:"numeral-fifths"`

	// generated from element "numeral-mode" of type numeral-mode
	Numeral_mode string `xml:"numeral-mode"`
}

// Numeral_root is generated from named complex type "numeral-root"
type Numeral_root struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "text" of type xs:token
	Text string `xml:"text,attr"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText int `xml:",chardata"`
}

// Octave_shift is generated from named complex type "octave-shift"
type Octave_shift struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type" of type up-down-stop-continue
	Type string `xml:"type,attr"`

	// generated from attribute "number" of type number-level
	Number int `xml:"number,attr"`

	// generated from attribute "size" of type xs:positiveInteger
	Size int `xml:"size,attr"`

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

	// generated from attribute "sound" of type yes-no
	Sound string `xml:"sound,attr"`

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

	// generated from element "trill-mark" of type empty-trill-sound
	Trill_mark []*Empty_trill_sound `xml:"trill-mark"`

	// generated from element "turn" of type horizontal-turn
	Turn []*Horizontal_turn `xml:"turn"`

	// generated from element "delayed-turn" of type horizontal-turn
	Delayed_turn []*Horizontal_turn `xml:"delayed-turn"`

	// generated from element "inverted-turn" of type horizontal-turn
	Inverted_turn []*Horizontal_turn `xml:"inverted-turn"`

	// generated from element "delayed-inverted-turn" of type horizontal-turn
	Delayed_inverted_turn []*Horizontal_turn `xml:"delayed-inverted-turn"`

	// generated from element "vertical-turn" of type empty-trill-sound
	Vertical_turn []*Empty_trill_sound `xml:"vertical-turn"`

	// generated from element "inverted-vertical-turn" of type empty-trill-sound
	Inverted_vertical_turn []*Empty_trill_sound `xml:"inverted-vertical-turn"`

	// generated from element "shake" of type empty-trill-sound
	Shake []*Empty_trill_sound `xml:"shake"`

	// generated from element "wavy-line" of type wavy-line
	Wavy_line []*Wavy_line `xml:"wavy-line"`

	// generated from element "mordent" of type mordent
	Mordent []*Mordent `xml:"mordent"`

	// generated from element "inverted-mordent" of type mordent
	Inverted_mordent []*Mordent `xml:"inverted-mordent"`

	// generated from element "schleifer" of type empty-placement
	Schleifer []*Empty_placement `xml:"schleifer"`

	// generated from element "tremolo" of type tremolo
	Tremolo []*Tremolo `xml:"tremolo"`

	// generated from element "haydn" of type empty-trill-sound
	Haydn []*Empty_trill_sound `xml:"haydn"`

	// generated from element "other-ornament" of type other-placement-text
	Other_ornament []*Other_placement_text `xml:"other-ornament"`

	// generated from element "accidental-mark" of type accidental-mark
	Accidental_mark []*Accidental_mark `xml:"accidental-mark"`
}

// Other_appearance is generated from named complex type "other-appearance"
type Other_appearance struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type" of type xs:token
	Type string `xml:"type,attr"`

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

	// generated from attribute "type" of type xs:token
	Type string `xml:"type,attr"`

	// generated from attribute "player" of type xs:IDREF
	Player string `xml:"player,attr"`

	// generated from attribute "time-only" of type time-only
	Time_only string `xml:"time-only,attr"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Other_notation is generated from named complex type "other-notation"
type Other_notation struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type" of type start-stop-single
	Type string `xml:"type,attr"`

	// generated from attribute "number" of type number-level
	Number int `xml:"number,attr"`

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

	// generated from attribute "type" of type xs:token
	Type string `xml:"type,attr"`

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

	// generated from element "page-height" of type tenths
	Page_height string `xml:"page-height"`

	// generated from element "page-width" of type tenths
	Page_width string `xml:"page-width"`

	// generated from element "page-margins" of type page-margins
	Page_margins []*Page_margins `xml:"page-margins"`
}

// Page_margins is generated from named complex type "page-margins"
type Page_margins struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type" of type margin-type
	Type string `xml:"type,attr"`

	// generated from element "left-margin" of type tenths
	Left_margin string `xml:"left-margin"`

	// generated from element "right-margin" of type tenths
	Right_margin string `xml:"right-margin"`

	// generated from element "top-margin" of type tenths
	Top_margin string `xml:"top-margin"`

	// generated from element "bottom-margin" of type tenths
	Bottom_margin string `xml:"bottom-margin"`
}

// Part_clef is generated from named complex type "part-clef"
type Part_clef struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "sign" of type clef-sign
	Sign string `xml:"sign"`

	// generated from element "line" of type staff-line-position
	Line int `xml:"line"`

	// generated from element "clef-octave-change" of type xs:integer
	Clef_octave_change int `xml:"clef-octave-change"`
}

// Part_group is generated from named complex type "part-group"
type Part_group struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type" of type start-stop
	Type string `xml:"type,attr"`

	// generated from attribute "number" of type xs:token
	Number string `xml:"number,attr"`

	// generated from element "footnote" of type formatted-text
	Footnote []*Formatted_text `xml:"footnote"`

	// generated from element "level" of type level
	Level []*Level `xml:"level"`

	// generated from element "group-name" of type group-name
	Group_name []*Group_name `xml:"group-name"`

	// generated from element "group-name-display" of type name-display
	Group_name_display []*Name_display `xml:"group-name-display"`

	// generated from element "group-abbreviation" of type group-name
	Group_abbreviation []*Group_name `xml:"group-abbreviation"`

	// generated from element "group-abbreviation-display" of type name-display
	Group_abbreviation_display []*Name_display `xml:"group-abbreviation-display"`

	// generated from element "group-symbol" of type group-symbol
	Group_symbol []*Group_symbol `xml:"group-symbol"`

	// generated from element "group-barline" of type group-barline
	Group_barline []*Group_barline `xml:"group-barline"`

	// generated from element "group-time" of type empty
	Group_time string `xml:"group-time"`
}

// Part_link is generated from named complex type "part-link"
type Part_link struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "link-attributes
	AttributeGroup_link_attributes

	// generated from element "instrument-link" of type instrument-link
	Instrument_link []*Instrument_link `xml:"instrument-link"`

	// generated from element "group-link" of type xs:string
	Group_link string `xml:"group-link"`
}

// Part_list is generated from named complex type "part-list"
type Part_list struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "part-group" of type part-group
	Part_group []*Part_group `xml:"part-group"`

	// generated from element "score-part" of type score-part
	Score_part []*Score_part `xml:"score-part"`
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

	// generated from attribute "top-staff" of type staff-number
	Top_staff int `xml:"top-staff,attr"`

	// generated from attribute "bottom-staff" of type staff-number
	Bottom_staff int `xml:"bottom-staff,attr"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Part_transpose is generated from named complex type "part-transpose"
type Part_transpose struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "diatonic" of type xs:integer
	Diatonic int `xml:"diatonic"`

	// generated from element "chromatic" of type semitones
	Chromatic string `xml:"chromatic"`

	// generated from element "octave-change" of type xs:integer
	Octave_change int `xml:"octave-change"`

	// generated from element "double" of type double
	Double []*Double `xml:"double"`
}

// Pedal is generated from named complex type "pedal"
type Pedal struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type" of type pedal-type
	Type string `xml:"type,attr"`

	// generated from attribute "number" of type number-level
	Number int `xml:"number,attr"`

	// generated from attribute "line" of type yes-no
	Line string `xml:"line,attr"`

	// generated from attribute "sign" of type yes-no
	Sign string `xml:"sign,attr"`

	// generated from attribute "abbreviated" of type yes-no
	Abbreviated string `xml:"abbreviated,attr"`

	// generated from attribute group "print-style-align
	AttributeGroup_print_style_align

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id
}

// Pedal_tuning is generated from named complex type "pedal-tuning"
type Pedal_tuning struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "pedal-step" of type step
	Pedal_step string `xml:"pedal-step"`

	// generated from element "pedal-alter" of type semitones
	Pedal_alter string `xml:"pedal-alter"`
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

	// generated from element "glass" of type glass
	Glass []*Glass `xml:"glass"`

	// generated from element "metal" of type metal
	Metal []*Metal `xml:"metal"`

	// generated from element "wood" of type wood
	Wood []*Wood `xml:"wood"`

	// generated from element "pitched" of type pitched
	Pitched []*Pitched `xml:"pitched"`

	// generated from element "membrane" of type membrane
	Membrane []*Membrane `xml:"membrane"`

	// generated from element "effect" of type effect
	Effect []*Effect `xml:"effect"`

	// generated from element "timpani" of type timpani
	Timpani []*Timpani `xml:"timpani"`

	// generated from element "beater" of type beater
	Beater []*Beater `xml:"beater"`

	// generated from element "stick" of type stick
	Stick []*Stick `xml:"stick"`

	// generated from element "stick-location" of type stick-location
	Stick_location string `xml:"stick-location"`

	// generated from element "other-percussion" of type other-text
	Other_percussion []*Other_text `xml:"other-percussion"`
}

// Pitch is generated from named complex type "pitch"
type Pitch struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "step" of type step
	Step string `xml:"step"`

	// generated from element "alter" of type semitones
	Alter string `xml:"alter"`

	// generated from element "octave" of type octave
	Octave int `xml:"octave"`
}

// Pitched is generated from named complex type "pitched"
type Pitched struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "smufl" of type smufl-pictogram-glyph-name
	Smufl string `xml:"smufl,attr"`

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

	// generated from attribute "id" of type xs:IDREF
	Id string `xml:"id,attr"`

	// generated from element "ipa" of type xs:string
	Ipa string `xml:"ipa"`

	// generated from element "mute" of type mute
	Mute string `xml:"mute"`

	// generated from element "semi-pitched" of type semi-pitched
	Semi_pitched string `xml:"semi-pitched"`

	// generated from element "other-play" of type other-play
	Other_play []*Other_play `xml:"other-play"`
}

// Player is generated from named complex type "player"
type Player struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id" of type xs:ID
	Id string `xml:"id,attr"`

	// generated from element "player-name" of type xs:string
	Player_name string `xml:"player-name"`
}

// Principal_voice is generated from named complex type "principal-voice"
type Principal_voice struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type" of type start-stop
	Type string `xml:"type,attr"`

	// generated from attribute "symbol" of type principal-voice-symbol
	Symbol string `xml:"symbol,attr"`

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

	// generated from element "page-layout" of type page-layout
	Page_layout []*Page_layout `xml:"page-layout"`

	// generated from element "system-layout" of type system-layout
	System_layout []*System_layout `xml:"system-layout"`

	// generated from element "staff-layout" of type staff-layout
	Staff_layout []*Staff_layout `xml:"staff-layout"`

	// generated from element "measure-layout" of type measure-layout
	Measure_layout []*Measure_layout `xml:"measure-layout"`

	// generated from element "measure-numbering" of type measure-numbering
	Measure_numbering []*Measure_numbering `xml:"measure-numbering"`

	// generated from element "part-name-display" of type name-display
	Part_name_display []*Name_display `xml:"part-name-display"`

	// generated from element "part-abbreviation-display" of type name-display
	Part_abbreviation_display []*Name_display `xml:"part-abbreviation-display"`
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

	// generated from attribute "direction" of type backward-forward
	Direction string `xml:"direction,attr"`

	// generated from attribute "times" of type xs:nonNegativeInteger
	Times int `xml:"times,attr"`

	// generated from attribute "after-jump" of type yes-no
	After_jump string `xml:"after-jump,attr"`

	// generated from attribute "winged" of type winged
	Winged string `xml:"winged,attr"`
}

// Rest is generated from named complex type "rest"
type Rest struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "measure" of type yes-no
	Measure string `xml:"measure,attr"`

	// generated from element "display-step" of type step
	Display_step string `xml:"display-step"`

	// generated from element "display-octave" of type octave
	Display_octave int `xml:"display-octave"`
}

// Root is generated from named complex type "root"
type Root struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "root-step" of type root-step
	Root_step []*Root_step `xml:"root-step"`

	// generated from element "root-alter" of type harmony-alter
	Root_alter []*Harmony_alter `xml:"root-alter"`
}

// Root_step is generated from named complex type "root-step"
type Root_step struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "text" of type xs:token
	Text string `xml:"text,attr"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Scaling is generated from named complex type "scaling"
type Scaling struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "millimeters" of type millimeters
	Millimeters string `xml:"millimeters"`

	// generated from element "tenths" of type tenths
	Tenths string `xml:"tenths"`
}

// Scordatura is generated from named complex type "scordatura"
type Scordatura struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from element "accord" of type accord
	Accord []*Accord `xml:"accord"`
}

// Score_instrument is generated from named complex type "score-instrument"
type Score_instrument struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id" of type xs:ID
	Id string `xml:"id,attr"`

	// generated from element "solo" of type empty
	Solo string `xml:"solo"`

	// generated from element "ensemble" of type positive-integer-or-empty
	Ensemble string `xml:"ensemble"`

	// generated from element "instrument-sound" of type xs:string
	Instrument_sound string `xml:"instrument-sound"`

	// generated from element "virtual-instrument" of type virtual-instrument
	Virtual_instrument []*Virtual_instrument `xml:"virtual-instrument"`

	// generated from element "instrument-name" of type xs:string
	Instrument_name string `xml:"instrument-name"`

	// generated from element "instrument-abbreviation" of type xs:string
	Instrument_abbreviation string `xml:"instrument-abbreviation"`
}

// Score_part is generated from named complex type "score-part"
type Score_part struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id" of type xs:ID
	Id string `xml:"id,attr"`

	// generated from element "midi-device" of type midi-device
	Midi_device []*Midi_device `xml:"midi-device"`

	// generated from element "midi-instrument" of type midi-instrument
	Midi_instrument []*Midi_instrument `xml:"midi-instrument"`

	// generated from element "identification" of type identification
	Identification []*Identification `xml:"identification"`

	// generated from element "part-link" of type part-link
	Part_link []*Part_link `xml:"part-link"`

	// generated from element "part-name" of type part-name
	Part_name []*Part_name `xml:"part-name"`

	// generated from element "part-name-display" of type name-display
	Part_name_display []*Name_display `xml:"part-name-display"`

	// generated from element "part-abbreviation" of type part-name
	Part_abbreviation []*Part_name `xml:"part-abbreviation"`

	// generated from element "part-abbreviation-display" of type name-display
	Part_abbreviation_display []*Name_display `xml:"part-abbreviation-display"`

	// generated from element "group" of type xs:string
	Group string `xml:"group"`

	// generated from element "score-instrument" of type score-instrument
	Score_instrument []*Score_instrument `xml:"score-instrument"`

	// generated from element "player" of type player
	Player []*Player `xml:"player"`
}

// Segno is generated from named complex type "segno"
type Segno struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "smufl" of type smufl-segno-glyph-name
	Smufl string `xml:"smufl,attr"`

	// generated from attribute group "print-style-align
	AttributeGroup_print_style_align

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id
}

// Slash is generated from named complex type "slash"
type Slash struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type" of type start-stop
	Type string `xml:"type,attr"`

	// generated from attribute "use-dots" of type yes-no
	Use_dots string `xml:"use-dots,attr"`

	// generated from attribute "use-stems" of type yes-no
	Use_stems string `xml:"use-stems,attr"`

	// generated from element "slash-type" of type note-type-value
	Slash_type string `xml:"slash-type"`

	// generated from element "slash-dot" of type empty
	Slash_dot string `xml:"slash-dot"`

	// generated from element "except-voice" of type xs:string
	Except_voice string `xml:"except-voice"`
}

// Slide is generated from named complex type "slide"
type Slide struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type" of type start-stop
	Type string `xml:"type,attr"`

	// generated from attribute "number" of type number-level
	Number int `xml:"number,attr"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Slur is generated from named complex type "slur"
type Slur struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type" of type start-stop-continue
	Type string `xml:"type,attr"`

	// generated from attribute "number" of type number-level
	Number int `xml:"number,attr"`

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

	// generated from attribute "tempo" of type non-negative-decimal
	Tempo string `xml:"tempo,attr"`

	// generated from attribute "dynamics" of type non-negative-decimal
	Dynamics string `xml:"dynamics,attr"`

	// generated from attribute "dacapo" of type yes-no
	Dacapo string `xml:"dacapo,attr"`

	// generated from attribute "segno" of type xs:token
	Segno string `xml:"segno,attr"`

	// generated from attribute "dalsegno" of type xs:token
	Dalsegno string `xml:"dalsegno,attr"`

	// generated from attribute "coda" of type xs:token
	Coda string `xml:"coda,attr"`

	// generated from attribute "tocoda" of type xs:token
	Tocoda string `xml:"tocoda,attr"`

	// generated from attribute "divisions" of type divisions
	Divisions string `xml:"divisions,attr"`

	// generated from attribute "forward-repeat" of type yes-no
	Forward_repeat string `xml:"forward-repeat,attr"`

	// generated from attribute "fine" of type xs:token
	Fine string `xml:"fine,attr"`

	// generated from attribute "time-only" of type time-only
	Time_only string `xml:"time-only,attr"`

	// generated from attribute "pizzicato" of type yes-no
	Pizzicato string `xml:"pizzicato,attr"`

	// generated from attribute "pan" of type rotation-degrees
	Pan string `xml:"pan,attr"`

	// generated from attribute "elevation" of type rotation-degrees
	Elevation string `xml:"elevation,attr"`

	// generated from attribute "damper-pedal" of type yes-no-number
	Damper_pedal string `xml:"damper-pedal,attr"`

	// generated from attribute "soft-pedal" of type yes-no-number
	Soft_pedal string `xml:"soft-pedal,attr"`

	// generated from attribute "sostenuto-pedal" of type yes-no-number
	Sostenuto_pedal string `xml:"sostenuto-pedal,attr"`

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from element "instrument-change" of type instrument-change
	Instrument_change []*Instrument_change `xml:"instrument-change"`

	// generated from element "midi-device" of type midi-device
	Midi_device []*Midi_device `xml:"midi-device"`

	// generated from element "midi-instrument" of type midi-instrument
	Midi_instrument []*Midi_instrument `xml:"midi-instrument"`

	// generated from element "play" of type play
	Play []*Play `xml:"play"`

	// generated from element "swing" of type swing
	Swing []*Swing `xml:"swing"`

	// generated from element "offset" of type offset
	Offset []*Offset `xml:"offset"`
}

// Staff_details is generated from named complex type "staff-details"
type Staff_details struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "number" of type staff-number
	Number int `xml:"number,attr"`

	// generated from attribute "show-frets" of type show-frets
	Show_frets string `xml:"show-frets,attr"`

	// generated from attribute group "print-object
	AttributeGroup_print_object

	// generated from attribute group "print-spacing
	AttributeGroup_print_spacing

	// generated from element "staff-lines" of type xs:nonNegativeInteger
	Staff_lines int `xml:"staff-lines"`

	// generated from element "line-detail" of type line-detail
	Line_detail []*Line_detail `xml:"line-detail"`

	// generated from element "staff-type" of type staff-type
	Staff_type string `xml:"staff-type"`

	// generated from element "staff-tuning" of type staff-tuning
	Staff_tuning []*Staff_tuning `xml:"staff-tuning"`

	// generated from element "capo" of type xs:nonNegativeInteger
	Capo int `xml:"capo"`

	// generated from element "staff-size" of type staff-size
	Staff_size []*Staff_size `xml:"staff-size"`
}

// Staff_divide is generated from named complex type "staff-divide"
type Staff_divide struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type" of type staff-divide-symbol
	Type string `xml:"type,attr"`

	// generated from attribute group "print-style-align
	AttributeGroup_print_style_align

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id
}

// Staff_layout is generated from named complex type "staff-layout"
type Staff_layout struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "number" of type staff-number
	Number int `xml:"number,attr"`

	// generated from element "staff-distance" of type tenths
	Staff_distance string `xml:"staff-distance"`
}

// Staff_size is generated from named complex type "staff-size"
type Staff_size struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "scaling" of type non-negative-decimal
	Scaling string `xml:"scaling,attr"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Staff_tuning is generated from named complex type "staff-tuning"
type Staff_tuning struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "line" of type staff-line
	Line int `xml:"line,attr"`

	// generated from element "tuning-step" of type step
	Tuning_step string `xml:"tuning-step"`

	// generated from element "tuning-alter" of type semitones
	Tuning_alter string `xml:"tuning-alter"`

	// generated from element "tuning-octave" of type octave
	Tuning_octave int `xml:"tuning-octave"`
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

	// generated from attribute "tip" of type tip-direction
	Tip string `xml:"tip,attr"`

	// generated from attribute "parentheses" of type yes-no
	Parentheses string `xml:"parentheses,attr"`

	// generated from attribute "dashed-circle" of type yes-no
	Dashed_circle string `xml:"dashed-circle,attr"`

	// generated from element "stick-type" of type stick-type
	Stick_type string `xml:"stick-type"`

	// generated from element "stick-material" of type stick-material
	Stick_material string `xml:"stick-material"`
}

// String_mute is generated from named complex type "string-mute"
type String_mute struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type" of type on-off
	Type string `xml:"type,attr"`

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

	// generated from attribute "type" of type yes-no
	Type string `xml:"type,attr"`

	// generated from attribute "element" of type xs:NMTOKEN
	Element string `xml:"element,attr"`

	// generated from attribute "attribute" of type xs:NMTOKEN
	Attribute string `xml:"attribute,attr"`

	// generated from attribute "value" of type xs:token
	Value string `xml:"value,attr"`
}

// Swing is generated from named complex type "swing"
type Swing struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "first" of type xs:positiveInteger
	First int `xml:"first"`

	// generated from element "second" of type xs:positiveInteger
	Second int `xml:"second"`

	// generated from element "swing-type" of type swing-type-value
	Swing_type string `xml:"swing-type"`

	// generated from element "straight" of type empty
	Straight string `xml:"straight"`

	// generated from element "swing-style" of type xs:string
	Swing_style string `xml:"swing-style"`
}

// Sync is generated from named complex type "sync"
type Sync struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type" of type sync-type
	Type string `xml:"type,attr"`

	// generated from attribute "latency" of type milliseconds
	Latency int `xml:"latency,attr"`

	// generated from attribute "player" of type xs:IDREF
	Player string `xml:"player,attr"`

	// generated from attribute "time-only" of type time-only
	Time_only string `xml:"time-only,attr"`
}

// System_dividers is generated from named complex type "system-dividers"
type System_dividers struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "left-divider" of type empty-print-object-style-align
	Left_divider []*Empty_print_object_style_align `xml:"left-divider"`

	// generated from element "right-divider" of type empty-print-object-style-align
	Right_divider []*Empty_print_object_style_align `xml:"right-divider"`
}

// System_layout is generated from named complex type "system-layout"
type System_layout struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "system-margins" of type system-margins
	System_margins []*System_margins `xml:"system-margins"`

	// generated from element "system-distance" of type tenths
	System_distance string `xml:"system-distance"`

	// generated from element "top-system-distance" of type tenths
	Top_system_distance string `xml:"top-system-distance"`

	// generated from element "system-dividers" of type system-dividers
	System_dividers []*System_dividers `xml:"system-dividers"`
}

// System_margins is generated from named complex type "system-margins"
type System_margins struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "left-margin" of type tenths
	Left_margin string `xml:"left-margin"`

	// generated from element "right-margin" of type tenths
	Right_margin string `xml:"right-margin"`
}

// Tap is generated from named complex type "tap"
type Tap struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "hand" of type tap-hand
	Hand string `xml:"hand,attr"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Technical is generated from named complex type "technical"
type Technical struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from element "up-bow" of type empty-placement
	Up_bow []*Empty_placement `xml:"up-bow"`

	// generated from element "down-bow" of type empty-placement
	Down_bow []*Empty_placement `xml:"down-bow"`

	// generated from element "harmonic" of type harmonic
	Harmonic []*Harmonic `xml:"harmonic"`

	// generated from element "open-string" of type empty-placement
	Open_string []*Empty_placement `xml:"open-string"`

	// generated from element "thumb-position" of type empty-placement
	Thumb_position []*Empty_placement `xml:"thumb-position"`

	// generated from element "fingering" of type fingering
	Fingering []*Fingering `xml:"fingering"`

	// generated from element "pluck" of type placement-text
	Pluck []*Placement_text `xml:"pluck"`

	// generated from element "double-tongue" of type empty-placement
	Double_tongue []*Empty_placement `xml:"double-tongue"`

	// generated from element "triple-tongue" of type empty-placement
	Triple_tongue []*Empty_placement `xml:"triple-tongue"`

	// generated from element "stopped" of type empty-placement-smufl
	Stopped []*Empty_placement_smufl `xml:"stopped"`

	// generated from element "snap-pizzicato" of type empty-placement
	Snap_pizzicato []*Empty_placement `xml:"snap-pizzicato"`

	// generated from element "fret" of type fret
	Fret []*Fret `xml:"fret"`

	// generated from element "string" of type string-type
	String []*String_type `xml:"string"`

	// generated from element "hammer-on" of type hammer-on-pull-off
	Hammer_on []*Hammer_on_pull_off `xml:"hammer-on"`

	// generated from element "pull-off" of type hammer-on-pull-off
	Pull_off []*Hammer_on_pull_off `xml:"pull-off"`

	// generated from element "bend" of type bend
	Bend []*Bend `xml:"bend"`

	// generated from element "tap" of type tap
	Tap []*Tap `xml:"tap"`

	// generated from element "heel" of type heel-toe
	Heel []*Heel_toe `xml:"heel"`

	// generated from element "toe" of type heel-toe
	Toe []*Heel_toe `xml:"toe"`

	// generated from element "fingernails" of type empty-placement
	Fingernails []*Empty_placement `xml:"fingernails"`

	// generated from element "hole" of type hole
	Hole []*Hole `xml:"hole"`

	// generated from element "arrow" of type arrow
	Arrow []*Arrow `xml:"arrow"`

	// generated from element "handbell" of type handbell
	Handbell []*Handbell `xml:"handbell"`

	// generated from element "brass-bend" of type empty-placement
	Brass_bend []*Empty_placement `xml:"brass-bend"`

	// generated from element "flip" of type empty-placement
	Flip []*Empty_placement `xml:"flip"`

	// generated from element "smear" of type empty-placement
	Smear []*Empty_placement `xml:"smear"`

	// generated from element "open" of type empty-placement-smufl
	Open []*Empty_placement_smufl `xml:"open"`

	// generated from element "half-muted" of type empty-placement-smufl
	Half_muted []*Empty_placement_smufl `xml:"half-muted"`

	// generated from element "harmon-mute" of type harmon-mute
	Harmon_mute []*Harmon_mute `xml:"harmon-mute"`

	// generated from element "golpe" of type empty-placement
	Golpe []*Empty_placement `xml:"golpe"`

	// generated from element "other-technical" of type other-placement-text
	Other_technical []*Other_placement_text `xml:"other-technical"`
}

// Text_element_data is generated from named complex type "text-element-data"
type Text_element_data struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "" of type 
	Lang string `xml:",attr"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Tie is generated from named complex type "tie"
type Tie struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type" of type start-stop
	Type string `xml:"type,attr"`

	// generated from attribute "time-only" of type time-only
	Time_only string `xml:"time-only,attr"`
}

// Tied is generated from named complex type "tied"
type Tied struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type" of type tied-type
	Type string `xml:"type,attr"`

	// generated from attribute "number" of type number-level
	Number int `xml:"number,attr"`

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

	// generated from attribute "number" of type staff-number
	Number int `xml:"number,attr"`

	// generated from attribute "symbol" of type time-symbol
	Symbol string `xml:"symbol,attr"`

	// generated from attribute "separator" of type time-separator
	Separator string `xml:"separator,attr"`

	// generated from attribute group "print-style-align
	AttributeGroup_print_style_align

	// generated from attribute group "print-object
	AttributeGroup_print_object

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from element "beats" of type xs:string
	Beats string `xml:"beats"`

	// generated from element "beat-type" of type xs:string
	Beat_type string `xml:"beat-type"`

	// generated from element "interchangeable" of type interchangeable
	Interchangeable []*Interchangeable `xml:"interchangeable"`

	// generated from element "senza-misura" of type xs:string
	Senza_misura string `xml:"senza-misura"`
}

// Time_modification is generated from named complex type "time-modification"
type Time_modification struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "normal-type" of type note-type-value
	Normal_type string `xml:"normal-type"`

	// generated from element "normal-dot" of type empty
	Normal_dot string `xml:"normal-dot"`

	// generated from element "actual-notes" of type xs:nonNegativeInteger
	Actual_notes int `xml:"actual-notes"`

	// generated from element "normal-notes" of type xs:nonNegativeInteger
	Normal_notes int `xml:"normal-notes"`
}

// Timpani is generated from named complex type "timpani"
type Timpani struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "smufl" of type smufl-pictogram-glyph-name
	Smufl string `xml:"smufl,attr"`
}

// Transpose is generated from named complex type "transpose"
type Transpose struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "number" of type staff-number
	Number int `xml:"number,attr"`

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from element "diatonic" of type xs:integer
	Diatonic int `xml:"diatonic"`

	// generated from element "chromatic" of type semitones
	Chromatic string `xml:"chromatic"`

	// generated from element "octave-change" of type xs:integer
	Octave_change int `xml:"octave-change"`

	// generated from element "double" of type double
	Double []*Double `xml:"double"`
}

// Tremolo is generated from named complex type "tremolo"
type Tremolo struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type" of type tremolo-type
	Type string `xml:"type,attr"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText int `xml:",chardata"`
}

// Tuplet is generated from named complex type "tuplet"
type Tuplet struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type" of type start-stop
	Type string `xml:"type,attr"`

	// generated from attribute "number" of type number-level
	Number int `xml:"number,attr"`

	// generated from attribute "bracket" of type yes-no
	Bracket string `xml:"bracket,attr"`

	// generated from attribute "show-number" of type show-tuplet
	Show_number string `xml:"show-number,attr"`

	// generated from attribute "show-type" of type show-tuplet
	Show_type string `xml:"show-type,attr"`

	// generated from attribute group "line-shape
	AttributeGroup_line_shape

	// generated from attribute group "position
	AttributeGroup_position

	// generated from attribute group "placement
	AttributeGroup_placement

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from element "tuplet-actual" of type tuplet-portion
	Tuplet_actual []*Tuplet_portion `xml:"tuplet-actual"`

	// generated from element "tuplet-normal" of type tuplet-portion
	Tuplet_normal []*Tuplet_portion `xml:"tuplet-normal"`
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

	// generated from element "tuplet-number" of type tuplet-number
	Tuplet_number []*Tuplet_number `xml:"tuplet-number"`

	// generated from element "tuplet-type" of type tuplet-type
	Tuplet_type []*Tuplet_type `xml:"tuplet-type"`

	// generated from element "tuplet-dot" of type tuplet-dot
	Tuplet_dot []*Tuplet_dot `xml:"tuplet-dot"`
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

	// generated from attribute "type" of type xs:token
	Type string `xml:"type,attr"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Unpitched is generated from named complex type "unpitched"
type Unpitched struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "display-step" of type step
	Display_step string `xml:"display-step"`

	// generated from element "display-octave" of type octave
	Display_octave int `xml:"display-octave"`
}

// Virtual_instrument is generated from named complex type "virtual-instrument"
type Virtual_instrument struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "virtual-library" of type xs:string
	Virtual_library string `xml:"virtual-library"`

	// generated from element "virtual-name" of type xs:string
	Virtual_name string `xml:"virtual-name"`
}

// Wait is generated from named complex type "wait"
type Wait struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "player" of type xs:IDREF
	Player string `xml:"player,attr"`

	// generated from attribute "time-only" of type time-only
	Time_only string `xml:"time-only,attr"`
}

// Wavy_line is generated from named complex type "wavy-line"
type Wavy_line struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type" of type start-stop-continue
	Type string `xml:"type,attr"`

	// generated from attribute "number" of type number-level
	Number int `xml:"number,attr"`

	// generated from attribute "smufl" of type smufl-wavy-line-glyph-name
	Smufl string `xml:"smufl,attr"`

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

	// generated from attribute "type" of type wedge-type
	Type string `xml:"type,attr"`

	// generated from attribute "number" of type number-level
	Number int `xml:"number,attr"`

	// generated from attribute "spread" of type tenths
	Spread string `xml:"spread,attr"`

	// generated from attribute "niente" of type yes-no
	Niente string `xml:"niente,attr"`

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

	// generated from attribute "smufl" of type smufl-pictogram-glyph-name
	Smufl string `xml:"smufl,attr"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Work is generated from named complex type "work"
type Work struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "work-number" of type xs:string
	Work_number string `xml:"work-number"`

	// generated from element "work-title" of type xs:string
	Work_title string `xml:"work-title"`

	// generated from element "opus" of type opus
	Opus []*Opus `xml:"opus"`
}

// Group_all_margins is generated from named group "all-margins"
type Group_all_margins struct {

	// insertion point for fields

	// generated from element "left-margin" of type tenths
	Left_margin string `xml:"left-margin"`

	// generated from element "right-margin" of type tenths
	Right_margin string `xml:"right-margin"`

	// generated from element "top-margin" of type tenths
	Top_margin string `xml:"top-margin"`

	// generated from element "bottom-margin" of type tenths
	Bottom_margin string `xml:"bottom-margin"`
}

// Group_beat_unit is generated from named group "beat-unit"
type Group_beat_unit struct {

	// insertion point for fields

	// generated from element "beat-unit" of type note-type-value
	Beat_unit string `xml:"beat-unit"`

	// generated from element "beat-unit-dot" of type empty
	Beat_unit_dot string `xml:"beat-unit-dot"`
}

// Group_clef is generated from named group "clef"
type Group_clef struct {

	// insertion point for fields

	// generated from element "sign" of type clef-sign
	Sign string `xml:"sign"`

	// generated from element "line" of type staff-line-position
	Line int `xml:"line"`

	// generated from element "clef-octave-change" of type xs:integer
	Clef_octave_change int `xml:"clef-octave-change"`
}

// Group_display_step_octave is generated from named group "display-step-octave"
type Group_display_step_octave struct {

	// insertion point for fields

	// generated from element "display-step" of type step
	Display_step string `xml:"display-step"`

	// generated from element "display-octave" of type octave
	Display_octave int `xml:"display-octave"`
}

// Group_duration is generated from named group "duration"
type Group_duration struct {

	// insertion point for fields

	// generated from element "duration" of type positive-divisions
	Duration string `xml:"duration"`
}

// Group_editorial is generated from named group "editorial"
type Group_editorial struct {

	// insertion point for fields

	// generated from element "footnote" of type formatted-text
	Footnote []*Formatted_text `xml:"footnote"`

	// generated from element "level" of type level
	Level []*Level `xml:"level"`
}

// Group_editorial_voice is generated from named group "editorial-voice"
type Group_editorial_voice struct {

	// insertion point for fields

	// generated from element "footnote" of type formatted-text
	Footnote []*Formatted_text `xml:"footnote"`

	// generated from element "level" of type level
	Level []*Level `xml:"level"`

	// generated from element "voice" of type xs:string
	Voice string `xml:"voice"`
}

// Group_editorial_voice_direction is generated from named group "editorial-voice-direction"
type Group_editorial_voice_direction struct {

	// insertion point for fields

	// generated from element "footnote" of type formatted-text
	Footnote []*Formatted_text `xml:"footnote"`

	// generated from element "level" of type level
	Level []*Level `xml:"level"`

	// generated from element "voice" of type xs:string
	Voice string `xml:"voice"`
}

// Group_footnote is generated from named group "footnote"
type Group_footnote struct {

	// insertion point for fields

	// generated from element "footnote" of type formatted-text
	Footnote []*Formatted_text `xml:"footnote"`
}

// Group_full_note is generated from named group "full-note"
type Group_full_note struct {

	// insertion point for fields

	// generated from element "pitch" of type pitch
	Pitch []*Pitch `xml:"pitch"`

	// generated from element "unpitched" of type unpitched
	Unpitched []*Unpitched `xml:"unpitched"`

	// generated from element "rest" of type rest
	Rest []*Rest `xml:"rest"`

	// generated from element "chord" of type empty
	Chord string `xml:"chord"`
}

// Group_harmony_chord is generated from named group "harmony-chord"
type Group_harmony_chord struct {

	// insertion point for fields

	// generated from element "root" of type root
	Root []*Root `xml:"root"`

	// generated from element "numeral" of type numeral
	Numeral []*Numeral `xml:"numeral"`

	// generated from element "function" of type style-text
	Function []*Style_text `xml:"function"`

	// generated from element "kind" of type kind
	Kind []*Kind `xml:"kind"`

	// generated from element "inversion" of type inversion
	Inversion []*Inversion `xml:"inversion"`

	// generated from element "bass" of type bass
	Bass []*Bass `xml:"bass"`

	// generated from element "degree" of type degree
	Degree []*Degree `xml:"degree"`
}

// Group_layout is generated from named group "layout"
type Group_layout struct {

	// insertion point for fields

	// generated from element "page-layout" of type page-layout
	Page_layout []*Page_layout `xml:"page-layout"`

	// generated from element "system-layout" of type system-layout
	System_layout []*System_layout `xml:"system-layout"`

	// generated from element "staff-layout" of type staff-layout
	Staff_layout []*Staff_layout `xml:"staff-layout"`
}

// Group_left_right_margins is generated from named group "left-right-margins"
type Group_left_right_margins struct {

	// insertion point for fields

	// generated from element "left-margin" of type tenths
	Left_margin string `xml:"left-margin"`

	// generated from element "right-margin" of type tenths
	Right_margin string `xml:"right-margin"`
}

// Group_level is generated from named group "level"
type Group_level struct {

	// insertion point for fields

	// generated from element "level" of type level
	Level []*Level `xml:"level"`
}

// Group_music_data is generated from named group "music-data"
type Group_music_data struct {

	// insertion point for fields

	// generated from element "note" of type note
	Note []*Note `xml:"note"`

	// generated from element "backup" of type backup
	Backup []*Backup `xml:"backup"`

	// generated from element "forward" of type forward
	Forward []*Forward `xml:"forward"`

	// generated from element "direction" of type direction
	Direction []*Direction `xml:"direction"`

	// generated from element "attributes" of type attributes
	Attributes []*Attributes `xml:"attributes"`

	// generated from element "harmony" of type harmony
	Harmony []*Harmony `xml:"harmony"`

	// generated from element "figured-bass" of type figured-bass
	Figured_bass []*Figured_bass `xml:"figured-bass"`

	// generated from element "print" of type print
	Print []*Print `xml:"print"`

	// generated from element "sound" of type sound
	Sound []*Sound `xml:"sound"`

	// generated from element "listening" of type listening
	Listening []*Listening `xml:"listening"`

	// generated from element "barline" of type barline
	Barline []*Barline `xml:"barline"`

	// generated from element "grouping" of type grouping
	Grouping []*Grouping `xml:"grouping"`

	// generated from element "link" of type link
	Link []*Link `xml:"link"`

	// generated from element "bookmark" of type bookmark
	Bookmark []*Bookmark `xml:"bookmark"`
}

// Group_non_traditional_key is generated from named group "non-traditional-key"
type Group_non_traditional_key struct {

	// insertion point for fields

	// generated from element "key-step" of type step
	Key_step string `xml:"key-step"`

	// generated from element "key-alter" of type semitones
	Key_alter string `xml:"key-alter"`

	// generated from element "key-accidental" of type key-accidental
	Key_accidental []*Key_accidental `xml:"key-accidental"`
}

// Group_part_group is generated from named group "part-group"
type Group_part_group struct {

	// insertion point for fields

	// generated from element "part-group" of type part-group
	Part_group []*Part_group `xml:"part-group"`
}

// Group_score_header is generated from named group "score-header"
type Group_score_header struct {

	// insertion point for fields

	// generated from element "work" of type work
	Work []*Work `xml:"work"`

	// generated from element "movement-number" of type xs:string
	Movement_number string `xml:"movement-number"`

	// generated from element "movement-title" of type xs:string
	Movement_title string `xml:"movement-title"`

	// generated from element "identification" of type identification
	Identification []*Identification `xml:"identification"`

	// generated from element "defaults" of type defaults
	Defaults []*Defaults `xml:"defaults"`

	// generated from element "credit" of type credit
	Credit []*Credit `xml:"credit"`

	// generated from element "part-list" of type part-list
	Part_list []*Part_list `xml:"part-list"`
}

// Group_score_part is generated from named group "score-part"
type Group_score_part struct {

	// insertion point for fields

	// generated from element "score-part" of type score-part
	Score_part []*Score_part `xml:"score-part"`
}

// Group_slash is generated from named group "slash"
type Group_slash struct {

	// insertion point for fields

	// generated from element "slash-type" of type note-type-value
	Slash_type string `xml:"slash-type"`

	// generated from element "slash-dot" of type empty
	Slash_dot string `xml:"slash-dot"`

	// generated from element "except-voice" of type xs:string
	Except_voice string `xml:"except-voice"`
}

// Group_staff is generated from named group "staff"
type Group_staff struct {

	// insertion point for fields

	// generated from element "staff" of type xs:positiveInteger
	Staff int `xml:"staff"`
}

// Group_time_signature is generated from named group "time-signature"
type Group_time_signature struct {

	// insertion point for fields

	// generated from element "beats" of type xs:string
	Beats string `xml:"beats"`

	// generated from element "beat-type" of type xs:string
	Beat_type string `xml:"beat-type"`
}

// Group_traditional_key is generated from named group "traditional-key"
type Group_traditional_key struct {

	// insertion point for fields

	// generated from element "cancel" of type cancel
	Cancel []*Cancel `xml:"cancel"`

	// generated from element "fifths" of type fifths
	Fifths int `xml:"fifths"`

	// generated from element "mode" of type mode
	Mode string `xml:"mode"`
}

// Group_transpose is generated from named group "transpose"
type Group_transpose struct {

	// insertion point for fields

	// generated from element "diatonic" of type xs:integer
	Diatonic int `xml:"diatonic"`

	// generated from element "chromatic" of type semitones
	Chromatic string `xml:"chromatic"`

	// generated from element "octave-change" of type xs:integer
	Octave_change int `xml:"octave-change"`

	// generated from element "double" of type double
	Double []*Double `xml:"double"`
}

// Group_tuning is generated from named group "tuning"
type Group_tuning struct {

	// insertion point for fields

	// generated from element "tuning-step" of type step
	Tuning_step string `xml:"tuning-step"`

	// generated from element "tuning-alter" of type semitones
	Tuning_alter string `xml:"tuning-alter"`

	// generated from element "tuning-octave" of type octave
	Tuning_octave int `xml:"tuning-octave"`
}

// Group_virtual_instrument_data is generated from named group "virtual-instrument-data"
type Group_virtual_instrument_data struct {

	// insertion point for fields

	// generated from element "solo" of type empty
	Solo string `xml:"solo"`

	// generated from element "ensemble" of type positive-integer-or-empty
	Ensemble string `xml:"ensemble"`

	// generated from element "instrument-sound" of type xs:string
	Instrument_sound string `xml:"instrument-sound"`

	// generated from element "virtual-instrument" of type virtual-instrument
	Virtual_instrument []*Virtual_instrument `xml:"virtual-instrument"`
}

// Group_voice is generated from named group "voice"
type Group_voice struct {

	// insertion point for fields

	// generated from element "voice" of type xs:string
	Voice string `xml:"voice"`
}

// AttributeGroup_bend_sound is generated from named attribute group "bend-sound"
type AttributeGroup_bend_sound struct {

	// insertion point for fields

	// generated from attribute "accelerate" of type yes-no
	Accelerate string `xml:"accelerate,attr"`

	// generated from attribute "beats" of type trill-beats
	Beats string `xml:"beats,attr"`

	// generated from attribute "first-beat" of type percent
	First_beat string `xml:"first-beat,attr"`

	// generated from attribute "last-beat" of type percent
	Last_beat string `xml:"last-beat,attr"`
}

// AttributeGroup_bezier is generated from named attribute group "bezier"
type AttributeGroup_bezier struct {

	// insertion point for fields

	// generated from attribute "bezier-x" of type tenths
	Bezier_x string `xml:"bezier-x,attr"`

	// generated from attribute "bezier-y" of type tenths
	Bezier_y string `xml:"bezier-y,attr"`

	// generated from attribute "bezier-x2" of type tenths
	Bezier_x2 string `xml:"bezier-x2,attr"`

	// generated from attribute "bezier-y2" of type tenths
	Bezier_y2 string `xml:"bezier-y2,attr"`

	// generated from attribute "bezier-offset" of type divisions
	Bezier_offset string `xml:"bezier-offset,attr"`

	// generated from attribute "bezier-offset2" of type divisions
	Bezier_offset2 string `xml:"bezier-offset2,attr"`
}

// AttributeGroup_color is generated from named attribute group "color"
type AttributeGroup_color struct {

	// insertion point for fields

	// generated from attribute "color" of type color
	Color string `xml:"color,attr"`
}

// AttributeGroup_dashed_formatting is generated from named attribute group "dashed-formatting"
type AttributeGroup_dashed_formatting struct {

	// insertion point for fields

	// generated from attribute "dash-length" of type tenths
	Dash_length string `xml:"dash-length,attr"`

	// generated from attribute "space-length" of type tenths
	Space_length string `xml:"space-length,attr"`
}

// AttributeGroup_directive is generated from named attribute group "directive"
type AttributeGroup_directive struct {

	// insertion point for fields

	// generated from attribute "directive" of type yes-no
	Directive string `xml:"directive,attr"`
}

// AttributeGroup_document_attributes is generated from named attribute group "document-attributes"
type AttributeGroup_document_attributes struct {

	// insertion point for fields

	// generated from attribute "version" of type xs:token
	Version string `xml:"version,attr"`
}

// AttributeGroup_element_position is generated from named attribute group "element-position"
type AttributeGroup_element_position struct {

	// insertion point for fields

	// generated from attribute "element" of type xs:NMTOKEN
	Element string `xml:"element,attr"`

	// generated from attribute "position" of type xs:positiveInteger
	Position int `xml:"position,attr"`
}

// AttributeGroup_enclosure is generated from named attribute group "enclosure"
type AttributeGroup_enclosure struct {

	// insertion point for fields

	// generated from attribute "enclosure" of type enclosure-shape
	Enclosure string `xml:"enclosure,attr"`
}

// AttributeGroup_font is generated from named attribute group "font"
type AttributeGroup_font struct {

	// insertion point for fields

	// generated from attribute "font-family" of type font-family
	Font_family string `xml:"font-family,attr"`

	// generated from attribute "font-style" of type font-style
	Font_style string `xml:"font-style,attr"`

	// generated from attribute "font-size" of type font-size
	Font_size string `xml:"font-size,attr"`

	// generated from attribute "font-weight" of type font-weight
	Font_weight string `xml:"font-weight,attr"`
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

	// generated from attribute "halign" of type left-center-right
	Halign string `xml:"halign,attr"`
}

// AttributeGroup_image_attributes is generated from named attribute group "image-attributes"
type AttributeGroup_image_attributes struct {

	// insertion point for fields

	// generated from attribute "source" of type xs:anyURI
	Source string `xml:"source,attr"`

	// generated from attribute "type" of type xs:token
	Type string `xml:"type,attr"`

	// generated from attribute "height" of type tenths
	Height string `xml:"height,attr"`

	// generated from attribute "width" of type tenths
	Width string `xml:"width,attr"`

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

	// generated from attribute "justify" of type left-center-right
	Justify string `xml:"justify,attr"`
}

// AttributeGroup_letter_spacing is generated from named attribute group "letter-spacing"
type AttributeGroup_letter_spacing struct {

	// insertion point for fields

	// generated from attribute "letter-spacing" of type number-or-normal
	Letter_spacing string `xml:"letter-spacing,attr"`
}

// AttributeGroup_level_display is generated from named attribute group "level-display"
type AttributeGroup_level_display struct {

	// insertion point for fields

	// generated from attribute "parentheses" of type yes-no
	Parentheses string `xml:"parentheses,attr"`

	// generated from attribute "bracket" of type yes-no
	Bracket string `xml:"bracket,attr"`

	// generated from attribute "size" of type symbol-size
	Size string `xml:"size,attr"`
}

// AttributeGroup_line_height is generated from named attribute group "line-height"
type AttributeGroup_line_height struct {

	// insertion point for fields

	// generated from attribute "line-height" of type number-or-normal
	Line_height string `xml:"line-height,attr"`
}

// AttributeGroup_line_length is generated from named attribute group "line-length"
type AttributeGroup_line_length struct {

	// insertion point for fields

	// generated from attribute "line-length" of type line-length
	Line_length string `xml:"line-length,attr"`
}

// AttributeGroup_line_shape is generated from named attribute group "line-shape"
type AttributeGroup_line_shape struct {

	// insertion point for fields

	// generated from attribute "line-shape" of type line-shape
	Line_shape string `xml:"line-shape,attr"`
}

// AttributeGroup_line_type is generated from named attribute group "line-type"
type AttributeGroup_line_type struct {

	// insertion point for fields

	// generated from attribute "line-type" of type line-type
	Line_type string `xml:"line-type,attr"`
}

// AttributeGroup_link_attributes is generated from named attribute group "link-attributes"
type AttributeGroup_link_attributes struct {

	// insertion point for fields

	// generated from attribute "" of type 
	Href string `xml:",attr"`

	// generated from attribute "" of type 
	Type string `xml:",attr"`

	// generated from attribute "" of type 
	Role string `xml:",attr"`

	// generated from attribute "" of type 
	Title string `xml:",attr"`

	// generated from attribute "" of type 
	Show string `xml:",attr"`

	// generated from attribute "" of type 
	Actuate string `xml:",attr"`
}

// AttributeGroup_measure_attributes is generated from named attribute group "measure-attributes"
type AttributeGroup_measure_attributes struct {

	// insertion point for fields

	// generated from attribute "number" of type xs:token
	Number string `xml:"number,attr"`

	// generated from attribute "text" of type measure-text
	Text string `xml:"text,attr"`

	// generated from attribute "implicit" of type yes-no
	Implicit string `xml:"implicit,attr"`

	// generated from attribute "non-controlling" of type yes-no
	Non_controlling string `xml:"non-controlling,attr"`

	// generated from attribute "width" of type tenths
	Width string `xml:"width,attr"`

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id
}

// AttributeGroup_optional_unique_id is generated from named attribute group "optional-unique-id"
type AttributeGroup_optional_unique_id struct {

	// insertion point for fields

	// generated from attribute "id" of type xs:ID
	Id string `xml:"id,attr"`
}

// AttributeGroup_orientation is generated from named attribute group "orientation"
type AttributeGroup_orientation struct {

	// insertion point for fields

	// generated from attribute "orientation" of type over-under
	Orientation string `xml:"orientation,attr"`
}

// AttributeGroup_part_attributes is generated from named attribute group "part-attributes"
type AttributeGroup_part_attributes struct {

	// insertion point for fields

	// generated from attribute "id" of type xs:IDREF
	Id string `xml:"id,attr"`
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

	// generated from attribute "placement" of type above-below
	Placement string `xml:"placement,attr"`
}

// AttributeGroup_position is generated from named attribute group "position"
type AttributeGroup_position struct {

	// insertion point for fields

	// generated from attribute "default-x" of type tenths
	Default_x string `xml:"default-x,attr"`

	// generated from attribute "default-y" of type tenths
	Default_y string `xml:"default-y,attr"`

	// generated from attribute "relative-x" of type tenths
	Relative_x string `xml:"relative-x,attr"`

	// generated from attribute "relative-y" of type tenths
	Relative_y string `xml:"relative-y,attr"`
}

// AttributeGroup_print_attributes is generated from named attribute group "print-attributes"
type AttributeGroup_print_attributes struct {

	// insertion point for fields

	// generated from attribute "staff-spacing" of type tenths
	Staff_spacing string `xml:"staff-spacing,attr"`

	// generated from attribute "new-system" of type yes-no
	New_system string `xml:"new-system,attr"`

	// generated from attribute "new-page" of type yes-no
	New_page string `xml:"new-page,attr"`

	// generated from attribute "blank-page" of type xs:positiveInteger
	Blank_page int `xml:"blank-page,attr"`

	// generated from attribute "page-number" of type xs:token
	Page_number string `xml:"page-number,attr"`
}

// AttributeGroup_print_object is generated from named attribute group "print-object"
type AttributeGroup_print_object struct {

	// insertion point for fields

	// generated from attribute "print-object" of type yes-no
	Print_object string `xml:"print-object,attr"`
}

// AttributeGroup_print_spacing is generated from named attribute group "print-spacing"
type AttributeGroup_print_spacing struct {

	// insertion point for fields

	// generated from attribute "print-spacing" of type yes-no
	Print_spacing string `xml:"print-spacing,attr"`
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

	// generated from attribute "print-dot" of type yes-no
	Print_dot string `xml:"print-dot,attr"`

	// generated from attribute "print-lyric" of type yes-no
	Print_lyric string `xml:"print-lyric,attr"`

	// generated from attribute group "print-object
	AttributeGroup_print_object

	// generated from attribute group "print-spacing
	AttributeGroup_print_spacing
}

// AttributeGroup_smufl is generated from named attribute group "smufl"
type AttributeGroup_smufl struct {

	// insertion point for fields

	// generated from attribute "smufl" of type smufl-glyph-name
	Smufl string `xml:"smufl,attr"`
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

	// generated from attribute "system" of type system-relation
	System string `xml:"system,attr"`
}

// AttributeGroup_text_decoration is generated from named attribute group "text-decoration"
type AttributeGroup_text_decoration struct {

	// insertion point for fields

	// generated from attribute "underline" of type number-of-lines
	Underline int `xml:"underline,attr"`

	// generated from attribute "overline" of type number-of-lines
	Overline int `xml:"overline,attr"`

	// generated from attribute "line-through" of type number-of-lines
	Line_through int `xml:"line-through,attr"`
}

// AttributeGroup_text_direction is generated from named attribute group "text-direction"
type AttributeGroup_text_direction struct {

	// insertion point for fields

	// generated from attribute "dir" of type text-direction
	Dir string `xml:"dir,attr"`
}

// AttributeGroup_text_formatting is generated from named attribute group "text-formatting"
type AttributeGroup_text_formatting struct {

	// insertion point for fields

	// generated from attribute "" of type 
	Lang string `xml:",attr"`

	// generated from attribute "" of type 
	Space string `xml:",attr"`

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

	// generated from attribute "rotation" of type rotation-degrees
	Rotation string `xml:"rotation,attr"`
}

// AttributeGroup_trill_sound is generated from named attribute group "trill-sound"
type AttributeGroup_trill_sound struct {

	// insertion point for fields

	// generated from attribute "start-note" of type start-note
	Start_note string `xml:"start-note,attr"`

	// generated from attribute "trill-step" of type trill-step
	Trill_step string `xml:"trill-step,attr"`

	// generated from attribute "two-note-turn" of type two-note-turn
	Two_note_turn string `xml:"two-note-turn,attr"`

	// generated from attribute "accelerate" of type yes-no
	Accelerate string `xml:"accelerate,attr"`

	// generated from attribute "beats" of type trill-beats
	Beats string `xml:"beats,attr"`

	// generated from attribute "second-beat" of type percent
	Second_beat string `xml:"second-beat,attr"`

	// generated from attribute "last-beat" of type percent
	Last_beat string `xml:"last-beat,attr"`
}

// AttributeGroup_valign is generated from named attribute group "valign"
type AttributeGroup_valign struct {

	// insertion point for fields

	// generated from attribute "valign" of type valign
	Valign string `xml:"valign,attr"`
}

// AttributeGroup_valign_image is generated from named attribute group "valign-image"
type AttributeGroup_valign_image struct {

	// insertion point for fields

	// generated from attribute "valign" of type valign-image
	Valign string `xml:"valign,attr"`
}

// AttributeGroup_x_position is generated from named attribute group "x-position"
type AttributeGroup_x_position struct {

	// insertion point for fields

	// generated from attribute "default-x" of type tenths
	Default_x string `xml:"default-x,attr"`

	// generated from attribute "default-y" of type tenths
	Default_y string `xml:"default-y,attr"`

	// generated from attribute "relative-x" of type tenths
	Relative_x string `xml:"relative-x,attr"`

	// generated from attribute "relative-y" of type tenths
	Relative_y string `xml:"relative-y,attr"`
}

// AttributeGroup_y_position is generated from named attribute group "y-position"
type AttributeGroup_y_position struct {

	// insertion point for fields

	// generated from attribute "default-x" of type tenths
	Default_x string `xml:"default-x,attr"`

	// generated from attribute "default-y" of type tenths
	Default_y string `xml:"default-y,attr"`

	// generated from attribute "relative-x" of type tenths
	Relative_x string `xml:"relative-x,attr"`

	// generated from attribute "relative-y" of type tenths
	Relative_y string `xml:"relative-y,attr"`
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
