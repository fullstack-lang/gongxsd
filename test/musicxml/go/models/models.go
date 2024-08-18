// generated code - do not edit
package models

import "encoding/xml"

// to avoid compilation error if no xml element is generated
var _ xml.Attr

// A_directive is generated from outer element "directive"
type A_directive struct {

	// insertion point for fields

	// generated from attribute "http://www.w3.org/XML/1998/namespace lang" of type 
	Lang string `xml:"http://www.w3.org/XML/1998/namespace lang,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// A_measure_1 is generated from outer element "measure"
type A_measure_1 struct {

	// insertion point for fields

	// generated from attribute group "measure-attributes
	AttributeGroup_measure_attributes

	// generated from element "note" of type note order 451
	Note []*Note `xml:"note"`

	// generated from element "backup" of type backup order 452
	Backup []*Backup `xml:"backup"`

	// generated from element "forward" of type forward order 453
	Forward []*Forward `xml:"forward"`

	// generated from element "direction" of type direction order 454
	Direction []*Direction `xml:"direction"`

	// generated from element "attributes" of type attributes order 455
	Attributes []*Attributes `xml:"attributes"`

	// generated from element "harmony" of type harmony order 456
	Harmony []*Harmony `xml:"harmony"`

	// generated from element "figured-bass" of type figured-bass order 457
	Figured_bass []*Figured_bass `xml:"figured-bass"`

	// generated from element "print" of type print order 458
	Print []*Print `xml:"print"`

	// generated from element "sound" of type sound order 459
	Sound []*Sound `xml:"sound"`

	// generated from element "listening" of type listening order 460
	Listening []*Listening `xml:"listening"`

	// generated from element "barline" of type barline order 461
	Barline []*Barline `xml:"barline"`

	// generated from element "grouping" of type grouping order 462
	Grouping []*Grouping `xml:"grouping"`

	// generated from element "link" of type link order 463
	Link []*Link `xml:"link"`

	// generated from element "bookmark" of type bookmark order 464
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

// A_part is generated from outer element "part"
type A_part struct {

	// insertion point for fields

	// generated from attribute group "part-attributes
	AttributeGroup_part_attributes

	// generated from anonymous type within outer element "measure" of type A_measure
	Measure []*A_measure_1 `xml:"measure"`
}

// A_part_1 is generated from outer element "part"
type A_part_1 struct {

	// insertion point for fields

	// generated from attribute group "part-attributes
	AttributeGroup_part_attributes

	// generated from element "note" of type note order 451
	Note []*Note `xml:"note"`

	// generated from element "backup" of type backup order 452
	Backup []*Backup `xml:"backup"`

	// generated from element "forward" of type forward order 453
	Forward []*Forward `xml:"forward"`

	// generated from element "direction" of type direction order 454
	Direction []*Direction `xml:"direction"`

	// generated from element "attributes" of type attributes order 455
	Attributes []*Attributes `xml:"attributes"`

	// generated from element "harmony" of type harmony order 456
	Harmony []*Harmony `xml:"harmony"`

	// generated from element "figured-bass" of type figured-bass order 457
	Figured_bass []*Figured_bass `xml:"figured-bass"`

	// generated from element "print" of type print order 458
	Print []*Print `xml:"print"`

	// generated from element "sound" of type sound order 459
	Sound []*Sound `xml:"sound"`

	// generated from element "listening" of type listening order 460
	Listening []*Listening `xml:"listening"`

	// generated from element "barline" of type barline order 461
	Barline []*Barline `xml:"barline"`

	// generated from element "grouping" of type grouping order 462
	Grouping []*Grouping `xml:"grouping"`

	// generated from element "link" of type link order 463
	Link []*Link `xml:"link"`

	// generated from element "bookmark" of type bookmark order 464
	Bookmark []*Bookmark `xml:"bookmark"`
}

// A_score_partwise is generated from outer element "score-partwise"
type A_score_partwise struct {

	// insertion point for fields

	// generated from attribute group "document-attributes
	AttributeGroup_document_attributes

	// generated from element "work" of type work order 466
	Work []*Work `xml:"work"`

	// generated from element "movement-number" of type xs:string order 467
	Movement_number string `xml:"movement-number"`

	// generated from element "movement-title" of type xs:string order 468
	Movement_title string `xml:"movement-title"`

	// generated from element "identification" of type identification order 469
	Identification []*Identification `xml:"identification"`

	// generated from element "defaults" of type defaults order 470
	Defaults []*Defaults `xml:"defaults"`

	// generated from element "credit" of type credit order 471
	Credit []*Credit `xml:"credit"`

	// generated from element "part-list" of type part-list order 472
	Part_list []*Part_list `xml:"part-list"`

	// generated from anonymous type within outer element "part" of type A_part
	Part []*A_part `xml:"part"`
}

// A_score_timewise is generated from outer element "score-timewise"
type A_score_timewise struct {

	// insertion point for fields

	// generated from attribute group "document-attributes
	AttributeGroup_document_attributes

	// generated from element "work" of type work order 466
	Work []*Work `xml:"work"`

	// generated from element "movement-number" of type xs:string order 467
	Movement_number string `xml:"movement-number"`

	// generated from element "movement-title" of type xs:string order 468
	Movement_title string `xml:"movement-title"`

	// generated from element "identification" of type identification order 469
	Identification []*Identification `xml:"identification"`

	// generated from element "defaults" of type defaults order 470
	Defaults []*Defaults `xml:"defaults"`

	// generated from element "credit" of type credit order 471
	Credit []*Credit `xml:"credit"`

	// generated from element "part-list" of type part-list order 472
	Part_list []*Part_list `xml:"part-list"`

	// generated from anonymous type within outer element "measure" of type A_measure
	Measure []*A_measure `xml:"measure"`
}

// Accidental is generated from named complex type "accidental"
type Accidental struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "cautionary" of type yes-no
	Cautionary string `xml:"cautionary,attr,omitempty"`

	// generated from attribute "editorial" of type yes-no
	Editorial string `xml:"editorial,attr,omitempty"`

	// generated from attribute "smufl" of type smufl-accidental-glyph-name
	Smufl string `xml:"smufl,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Accidental_mark is generated from named complex type "accidental-mark"
type Accidental_mark struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "smufl" of type smufl-accidental-glyph-name
	Smufl string `xml:"smufl,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Accidental_text is generated from named complex type "accidental-text"
type Accidental_text struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "smufl" of type smufl-accidental-glyph-name
	Smufl string `xml:"smufl,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Accord is generated from named complex type "accord"
type Accord struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "string" of type string-number
	String int `xml:"string,attr,omitempty"`

	// generated from element "tuning-step" of type step order 402
	Tuning_step string `xml:"tuning-step"`

	// generated from element "tuning-alter" of type semitones order 403
	Tuning_alter string `xml:"tuning-alter"`

	// generated from element "tuning-octave" of type octave order 404
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

	// generated from element "accordion-high" of type empty order 76
	Accordion_high string `xml:"accordion-high"`

	// generated from element "accordion-middle" of type accordion-middle order 77
	Accordion_middle int `xml:"accordion-middle"`

	// generated from element "accordion-low" of type empty order 78
	Accordion_low string `xml:"accordion-low"`
}

// Appearance is generated from named complex type "appearance"
type Appearance struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "line-width" of type line-width order 190
	Line_width []*Line_width `xml:"line-width"`

	// generated from element "note-size" of type note-size order 191
	Note_size []*Note_size `xml:"note-size"`

	// generated from element "distance" of type distance order 192
	Distance []*Distance `xml:"distance"`

	// generated from element "glyph" of type glyph order 193
	Glyph []*Glyph `xml:"glyph"`

	// generated from element "other-appearance" of type other-appearance order 194
	Other_appearance []*Other_appearance `xml:"other-appearance"`
}

// Arpeggiate is generated from named complex type "arpeggiate"
type Arpeggiate struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "number" of type number-level
	Number int `xml:"number,attr,omitempty"`

	// generated from attribute "direction" of type up-down
	Direction string `xml:"direction,attr,omitempty"`

	// generated from attribute "unbroken" of type yes-no
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

	// generated from element "arrow-direction" of type arrow-direction order 225
	Arrow_direction string `xml:"arrow-direction"`

	// generated from element "arrow-style" of type arrow-style order 226
	Arrow_style string `xml:"arrow-style"`

	// generated from element "arrowhead" of type empty order 227
	Arrowhead string `xml:"arrowhead"`

	// generated from element "circular-arrow" of type circular-arrow order 228
	Circular_arrow string `xml:"circular-arrow"`
}

// Articulations is generated from named complex type "articulations"
type Articulations struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from element "accent" of type empty-placement order 208
	Accent []*Empty_placement `xml:"accent"`

	// generated from element "strong-accent" of type strong-accent order 209
	Strong_accent []*Strong_accent `xml:"strong-accent"`

	// generated from element "staccato" of type empty-placement order 210
	Staccato []*Empty_placement `xml:"staccato"`

	// generated from element "tenuto" of type empty-placement order 211
	Tenuto []*Empty_placement `xml:"tenuto"`

	// generated from element "detached-legato" of type empty-placement order 212
	Detached_legato []*Empty_placement `xml:"detached-legato"`

	// generated from element "staccatissimo" of type empty-placement order 213
	Staccatissimo []*Empty_placement `xml:"staccatissimo"`

	// generated from element "spiccato" of type empty-placement order 214
	Spiccato []*Empty_placement `xml:"spiccato"`

	// generated from element "scoop" of type empty-line order 215
	Scoop []*Empty_line `xml:"scoop"`

	// generated from element "plop" of type empty-line order 216
	Plop []*Empty_line `xml:"plop"`

	// generated from element "doit" of type empty-line order 217
	Doit []*Empty_line `xml:"doit"`

	// generated from element "falloff" of type empty-line order 218
	Falloff []*Empty_line `xml:"falloff"`

	// generated from element "breath-mark" of type breath-mark order 219
	Breath_mark []*Breath_mark `xml:"breath-mark"`

	// generated from element "caesura" of type caesura order 220
	Caesura []*Caesura `xml:"caesura"`

	// generated from element "stress" of type empty-placement order 221
	Stress []*Empty_placement `xml:"stress"`

	// generated from element "unstress" of type empty-placement order 222
	Unstress []*Empty_placement `xml:"unstress"`

	// generated from element "soft-accent" of type empty-placement order 223
	Soft_accent []*Empty_placement `xml:"soft-accent"`

	// generated from element "other-articulation" of type other-placement-text order 224
	Other_articulation []*Other_placement_text `xml:"other-articulation"`
}

// Assess is generated from named complex type "assess"
type Assess struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type" of type yes-no
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute "player" of type xs:IDREF
	Player string `xml:"player,attr,omitempty"`

	// generated from attribute "time-only" of type time-only
	Time_only string `xml:"time-only,attr,omitempty"`
}

// Attributes is generated from named complex type "attributes"
type Attributes struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "divisions" of type positive-divisions order 41
	Divisions string `xml:"divisions"`

	// generated from element "key" of type key order 42
	Key []*Key `xml:"key"`

	// generated from element "time" of type time order 43
	Time []*Time `xml:"time"`

	// generated from element "staves" of type xs:nonNegativeInteger order 44
	Staves int `xml:"staves"`

	// generated from element "part-symbol" of type part-symbol order 45
	Part_symbol []*Part_symbol `xml:"part-symbol"`

	// generated from element "instruments" of type xs:nonNegativeInteger order 46
	Instruments int `xml:"instruments"`

	// generated from element "clef" of type clef order 47
	Clef []*Clef `xml:"clef"`

	// generated from element "staff-details" of type staff-details order 48
	Staff_details []*Staff_details `xml:"staff-details"`

	// generated from element "transpose" of type transpose order 49
	Transpose []*Transpose `xml:"transpose"`

	// generated from element "for-part" of type for-part order 50
	For_part []*For_part `xml:"for-part"`

	// generated from anonymous type within outer element "directive" of type A_directive
	Directive []*A_directive `xml:"directive"`

	// generated from element "measure-style" of type measure-style order 52
	Measure_style []*Measure_style `xml:"measure-style"`

	// generated from element "footnote" of type formatted-text order 399
	Footnote []*Formatted_text `xml:"footnote"`

	// generated from element "level" of type level order 400
	Level []*Level `xml:"level"`
}

// Backup is generated from named complex type "backup"
type Backup struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "footnote" of type formatted-text order 399
	Footnote []*Formatted_text `xml:"footnote"`

	// generated from element "level" of type level order 400
	Level []*Level `xml:"level"`

	// generated from element "duration" of type positive-divisions order 444
	Duration string `xml:"duration"`
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
	Location string `xml:"location,attr,omitempty"`

	// generated from attribute "segno" of type xs:token
	Segno string `xml:"segno,attr,omitempty"`

	// generated from attribute "coda" of type xs:token
	Coda string `xml:"coda,attr,omitempty"`

	// generated from attribute "divisions" of type divisions
	Divisions string `xml:"divisions,attr,omitempty"`

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from element "bar-style" of type bar-style-color order 69
	Bar_style []*Bar_style_color `xml:"bar-style"`

	// generated from element "wavy-line" of type wavy-line order 70
	Wavy_line []*Wavy_line `xml:"wavy-line"`

	// generated from element "segno" of type segno order 71
	Segno_1 []*Segno `xml:"segno"`

	// generated from element "coda" of type coda order 72
	Coda_1 []*Coda `xml:"coda"`

	// generated from element "fermata" of type fermata order 73
	Fermata []*Fermata `xml:"fermata"`

	// generated from element "ending" of type ending order 74
	Ending []*Ending `xml:"ending"`

	// generated from element "repeat" of type repeat order 75
	Repeat []*Repeat `xml:"repeat"`

	// generated from element "footnote" of type formatted-text order 399
	Footnote []*Formatted_text `xml:"footnote"`

	// generated from element "level" of type level order 400
	Level []*Level `xml:"level"`
}

// Barre is generated from named complex type "barre"
type Barre struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type" of type start-stop
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute group "color
	AttributeGroup_color
}

// Bass is generated from named complex type "bass"
type Bass struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "arrangement" of type harmony-arrangement
	Arrangement string `xml:"arrangement,attr,omitempty"`

	// generated from element "bass-separator" of type style-text order 79
	Bass_separator []*Style_text `xml:"bass-separator"`

	// generated from element "bass-step" of type bass-step order 80
	Bass_step []*Bass_step `xml:"bass-step"`

	// generated from element "bass-alter" of type harmony-alter order 81
	Bass_alter []*Harmony_alter `xml:"bass-alter"`
}

// Bass_step is generated from named complex type "bass-step"
type Bass_step struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "text" of type xs:token
	Text string `xml:"text,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Beam is generated from named complex type "beam"
type Beam struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "number" of type beam-level
	Number int `xml:"number,attr,omitempty"`

	// generated from attribute "repeater" of type yes-no
	Repeater string `xml:"repeater,attr,omitempty"`

	// generated from attribute "fan" of type fan
	Fan string `xml:"fan,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Beat_repeat is generated from named complex type "beat-repeat"
type Beat_repeat struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type" of type start-stop
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute "slashes" of type xs:positiveInteger
	Slashes int `xml:"slashes,attr,omitempty"`

	// generated from attribute "use-dots" of type yes-no
	Use_dots string `xml:"use-dots,attr,omitempty"`

	// generated from element "slash-type" of type note-type-value order 416
	Slash_type string `xml:"slash-type"`

	// generated from element "slash-dot" of type empty order 417
	Slash_dot string `xml:"slash-dot"`

	// generated from element "except-voice" of type xs:string order 418
	Except_voice string `xml:"except-voice"`
}

// Beat_unit_tied is generated from named complex type "beat-unit-tied"
type Beat_unit_tied struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "beat-unit" of type note-type-value order 428
	Beat_unit string `xml:"beat-unit"`

	// generated from element "beat-unit-dot" of type empty order 429
	Beat_unit_dot string `xml:"beat-unit-dot"`
}

// Beater is generated from named complex type "beater"
type Beater struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "tip" of type tip-direction
	Tip string `xml:"tip,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Bend is generated from named complex type "bend"
type Bend struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "shape" of type bend-shape
	Shape string `xml:"shape,attr,omitempty"`

	// generated from attribute group "print-style
	AttributeGroup_print_style

	// generated from attribute group "bend-sound
	AttributeGroup_bend_sound

	// generated from element "bend-alter" of type semitones order 229
	Bend_alter string `xml:"bend-alter"`

	// generated from element "pre-bend" of type empty order 230
	Pre_bend string `xml:"pre-bend"`

	// generated from element "release" of type release order 231
	Release []*Release `xml:"release"`

	// generated from element "with-bar" of type placement-text order 232
	With_bar []*Placement_text `xml:"with-bar"`
}

// Bookmark is generated from named complex type "bookmark"
type Bookmark struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id" of type xs:ID
	Id string `xml:"id,attr,omitempty"`

	// generated from attribute "name" of type xs:token
	NameXSD string `xml:"name,attr,omitempty"`

	// generated from attribute group "element-position
	AttributeGroup_element_position
}

// Bracket is generated from named complex type "bracket"
type Bracket struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type" of type start-stop-continue
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute "number" of type number-level
	Number int `xml:"number,attr,omitempty"`

	// generated from attribute "line-end" of type line-end
	Line_end string `xml:"line-end,attr,omitempty"`

	// generated from attribute "end-length" of type tenths
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

	// generated from attribute "location" of type cancel-location
	Location string `xml:"location,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText int `xml:",chardata"`
}

// Clef is generated from named complex type "clef"
type Clef struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "number" of type staff-number
	Number int `xml:"number,attr,omitempty"`

	// generated from attribute "additional" of type yes-no
	Additional string `xml:"additional,attr,omitempty"`

	// generated from attribute "size" of type symbol-size
	Size string `xml:"size,attr,omitempty"`

	// generated from attribute "after-barline" of type yes-no
	After_barline string `xml:"after-barline,attr,omitempty"`

	// generated from attribute group "print-style
	AttributeGroup_print_style

	// generated from attribute group "print-object
	AttributeGroup_print_object

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from element "sign" of type clef-sign order 410
	Sign string `xml:"sign"`

	// generated from element "line" of type staff-line-position order 411
	Line int `xml:"line"`

	// generated from element "clef-octave-change" of type xs:integer order 412
	Clef_octave_change int `xml:"clef-octave-change"`
}

// Coda is generated from named complex type "coda"
type Coda struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "smufl" of type smufl-coda-glyph-name
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

	// generated from attribute "page" of type xs:positiveInteger
	Page int `xml:"page,attr,omitempty"`

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from element "credit-type" of type xs:string order 354
	Credit_type string `xml:"credit-type"`

	// generated from element "credit-image" of type image order 357
	Credit_image []*Image `xml:"credit-image"`

	// generated from element "link" of type link order 360
	Link []*Link `xml:"link"`

	// generated from element "bookmark" of type bookmark order 361
	Bookmark []*Bookmark `xml:"bookmark"`

	// generated from element "credit-words" of type formatted-text-id order 362
	Credit_words []*Formatted_text_id `xml:"credit-words"`

	// generated from element "credit-symbol" of type formatted-symbol-id order 363
	Credit_symbol []*Formatted_symbol_id `xml:"credit-symbol"`
}

// Dashes is generated from named complex type "dashes"
type Dashes struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type" of type start-stop-continue
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute "number" of type number-level
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

	// generated from element "scaling" of type scaling order 364
	Scaling []*Scaling `xml:"scaling"`

	// generated from element "concert-score" of type empty order 365
	Concert_score string `xml:"concert-score"`

	// generated from element "appearance" of type appearance order 366
	Appearance []*Appearance `xml:"appearance"`

	// generated from element "music-font" of type empty-font order 367
	Music_font []*Empty_font `xml:"music-font"`

	// generated from element "word-font" of type empty-font order 368
	Word_font []*Empty_font `xml:"word-font"`

	// generated from element "lyric-font" of type lyric-font order 369
	Lyric_font []*Lyric_font `xml:"lyric-font"`

	// generated from element "lyric-language" of type lyric-language order 370
	Lyric_language []*Lyric_language `xml:"lyric-language"`

	// generated from element "page-layout" of type page-layout order 439
	Page_layout []*Page_layout `xml:"page-layout"`

	// generated from element "system-layout" of type system-layout order 440
	System_layout []*System_layout `xml:"system-layout"`

	// generated from element "staff-layout" of type staff-layout order 441
	Staff_layout []*Staff_layout `xml:"staff-layout"`
}

// Degree is generated from named complex type "degree"
type Degree struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "print-object
	AttributeGroup_print_object

	// generated from element "degree-value" of type degree-value order 82
	Degree_value []*Degree_value `xml:"degree-value"`

	// generated from element "degree-alter" of type degree-alter order 83
	Degree_alter []*Degree_alter `xml:"degree-alter"`

	// generated from element "degree-type" of type degree-type order 84
	Degree_type []*Degree_type `xml:"degree-type"`
}

// Degree_alter is generated from named complex type "degree-alter"
type Degree_alter struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "plus-minus" of type yes-no
	Plus_minus string `xml:"plus-minus,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Degree_type is generated from named complex type "degree-type"
type Degree_type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "text" of type xs:token
	Text string `xml:"text,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Degree_value is generated from named complex type "degree-value"
type Degree_value struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "symbol" of type degree-symbol-value
	Symbol string `xml:"symbol,attr,omitempty"`

	// generated from attribute "text" of type xs:token
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

	// generated from element "direction-type" of type direction-type order 85
	Direction_type []*Direction_type `xml:"direction-type"`

	// generated from element "offset" of type offset order 86
	Offset []*Offset `xml:"offset"`

	// generated from element "sound" of type sound order 87
	Sound []*Sound `xml:"sound"`

	// generated from element "listening" of type listening order 88
	Listening []*Listening `xml:"listening"`

	// generated from element "footnote" of type formatted-text order 399
	Footnote []*Formatted_text `xml:"footnote"`

	// generated from element "level" of type level order 400
	Level []*Level `xml:"level"`

	// generated from element "staff" of type xs:positiveInteger order 401
	Staff int `xml:"staff"`

	// generated from element "voice" of type xs:string order 409
	Voice string `xml:"voice"`
}

// Direction_type is generated from named complex type "direction-type"
type Direction_type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from element "rehearsal" of type formatted-text-id order 89
	Rehearsal []*Formatted_text_id `xml:"rehearsal"`

	// generated from element "segno" of type segno order 90
	Segno []*Segno `xml:"segno"`

	// generated from element "coda" of type coda order 91
	Coda []*Coda `xml:"coda"`

	// generated from element "words" of type formatted-text-id order 92
	Words []*Formatted_text_id `xml:"words"`

	// generated from element "symbol" of type formatted-symbol-id order 93
	Symbol []*Formatted_symbol_id `xml:"symbol"`

	// generated from element "wedge" of type wedge order 94
	Wedge []*Wedge `xml:"wedge"`

	// generated from element "dynamics" of type dynamics order 95
	Dynamics []*Dynamics `xml:"dynamics"`

	// generated from element "dashes" of type dashes order 96
	Dashes []*Dashes `xml:"dashes"`

	// generated from element "bracket" of type bracket order 97
	Bracket []*Bracket `xml:"bracket"`

	// generated from element "pedal" of type pedal order 98
	Pedal []*Pedal `xml:"pedal"`

	// generated from element "metronome" of type metronome order 99
	Metronome []*Metronome `xml:"metronome"`

	// generated from element "octave-shift" of type octave-shift order 100
	Octave_shift []*Octave_shift `xml:"octave-shift"`

	// generated from element "harp-pedals" of type harp-pedals order 101
	Harp_pedals []*Harp_pedals `xml:"harp-pedals"`

	// generated from element "damp" of type empty-print-style-align-id order 102
	Damp []*Empty_print_style_align_id `xml:"damp"`

	// generated from element "damp-all" of type empty-print-style-align-id order 103
	Damp_all []*Empty_print_style_align_id `xml:"damp-all"`

	// generated from element "eyeglasses" of type empty-print-style-align-id order 104
	Eyeglasses []*Empty_print_style_align_id `xml:"eyeglasses"`

	// generated from element "string-mute" of type string-mute order 105
	String_mute []*String_mute `xml:"string-mute"`

	// generated from element "scordatura" of type scordatura order 106
	Scordatura []*Scordatura `xml:"scordatura"`

	// generated from element "image" of type image order 107
	Image []*Image `xml:"image"`

	// generated from element "principal-voice" of type principal-voice order 108
	Principal_voice []*Principal_voice `xml:"principal-voice"`

	// generated from element "percussion" of type percussion order 109
	Percussion []*Percussion `xml:"percussion"`

	// generated from element "accordion-registration" of type accordion-registration order 110
	Accordion_registration []*Accordion_registration `xml:"accordion-registration"`

	// generated from element "staff-divide" of type staff-divide order 111
	Staff_divide []*Staff_divide `xml:"staff-divide"`

	// generated from element "other-direction" of type other-direction order 112
	Other_direction []*Other_direction `xml:"other-direction"`
}

// Distance is generated from named complex type "distance"
type Distance struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type" of type distance-type
	Type string `xml:"type,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Double is generated from named complex type "double"
type Double struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "above" of type yes-no
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

	// generated from element "p" of type empty order 0
	P string `xml:"p"`

	// generated from element "pp" of type empty order 1
	Pp string `xml:"pp"`

	// generated from element "ppp" of type empty order 2
	Ppp string `xml:"ppp"`

	// generated from element "pppp" of type empty order 3
	Pppp string `xml:"pppp"`

	// generated from element "ppppp" of type empty order 4
	Ppppp string `xml:"ppppp"`

	// generated from element "pppppp" of type empty order 5
	Pppppp string `xml:"pppppp"`

	// generated from element "f" of type empty order 6
	F string `xml:"f"`

	// generated from element "ff" of type empty order 7
	Ff string `xml:"ff"`

	// generated from element "fff" of type empty order 8
	Fff string `xml:"fff"`

	// generated from element "ffff" of type empty order 9
	Ffff string `xml:"ffff"`

	// generated from element "fffff" of type empty order 10
	Fffff string `xml:"fffff"`

	// generated from element "ffffff" of type empty order 11
	Ffffff string `xml:"ffffff"`

	// generated from element "mp" of type empty order 12
	Mp string `xml:"mp"`

	// generated from element "mf" of type empty order 13
	Mf string `xml:"mf"`

	// generated from element "sf" of type empty order 14
	Sf string `xml:"sf"`

	// generated from element "sfp" of type empty order 15
	Sfp string `xml:"sfp"`

	// generated from element "sfpp" of type empty order 16
	Sfpp string `xml:"sfpp"`

	// generated from element "fp" of type empty order 17
	Fp string `xml:"fp"`

	// generated from element "rf" of type empty order 18
	Rf string `xml:"rf"`

	// generated from element "rfz" of type empty order 19
	Rfz string `xml:"rfz"`

	// generated from element "sfz" of type empty order 20
	Sfz string `xml:"sfz"`

	// generated from element "sffz" of type empty order 21
	Sffz string `xml:"sffz"`

	// generated from element "fz" of type empty order 22
	Fz string `xml:"fz"`

	// generated from element "n" of type empty order 23
	N string `xml:"n"`

	// generated from element "pf" of type empty order 24
	Pf string `xml:"pf"`

	// generated from element "sfzp" of type empty order 25
	Sfzp string `xml:"sfzp"`

	// generated from element "other-dynamics" of type other-text order 26
	Other_dynamics []*Other_text `xml:"other-dynamics"`
}

// Effect is generated from named complex type "effect"
type Effect struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "smufl" of type smufl-pictogram-glyph-name
	Smufl string `xml:"smufl,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Elision is generated from named complex type "elision"
type Elision struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "smufl" of type smufl-lyrics-glyph-name
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

	// generated from element "encoder" of type typed-text order 179
	Encoder []*Typed_text `xml:"encoder"`

	// generated from element "software" of type xs:string order 180
	Software string `xml:"software"`

	// generated from element "encoding-description" of type xs:string order 181
	Encoding_description string `xml:"encoding-description"`

	// generated from element "supports" of type supports order 182
	Supports []*Supports `xml:"supports"`
}

// Ending is generated from named complex type "ending"
type Ending struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "number" of type ending-number
	Number string `xml:"number,attr,omitempty"`

	// generated from attribute "type" of type start-stop-discontinue
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute "end-length" of type tenths
	End_length string `xml:"end-length,attr,omitempty"`

	// generated from attribute "text-x" of type tenths
	Text_x string `xml:"text-x,attr,omitempty"`

	// generated from attribute "text-y" of type tenths
	Text_y string `xml:"text-y,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Extend is generated from named complex type "extend"
type Extend struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type" of type start-stop-continue
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

	// generated from attribute "type" of type xs:token
	Type string `xml:"type,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Fermata is generated from named complex type "fermata"
type Fermata struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type" of type upright-inverted
	Type string `xml:"type,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Figure is generated from named complex type "figure"
type Figure struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "prefix" of type style-text order 233
	Prefix []*Style_text `xml:"prefix"`

	// generated from element "figure-number" of type style-text order 234
	Figure_number []*Style_text `xml:"figure-number"`

	// generated from element "suffix" of type style-text order 235
	Suffix []*Style_text `xml:"suffix"`

	// generated from element "extend" of type extend order 236
	Extend []*Extend `xml:"extend"`

	// generated from element "footnote" of type formatted-text order 399
	Footnote []*Formatted_text `xml:"footnote"`

	// generated from element "level" of type level order 400
	Level []*Level `xml:"level"`
}

// Figured_bass is generated from named complex type "figured-bass"
type Figured_bass struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "parentheses" of type yes-no
	Parentheses string `xml:"parentheses,attr,omitempty"`

	// generated from attribute group "print-style-align
	AttributeGroup_print_style_align

	// generated from attribute group "placement
	AttributeGroup_placement

	// generated from attribute group "printout
	AttributeGroup_printout

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from element "figure" of type figure order 237
	Figure []*Figure `xml:"figure"`

	// generated from element "footnote" of type formatted-text order 399
	Footnote []*Formatted_text `xml:"footnote"`

	// generated from element "level" of type level order 400
	Level []*Level `xml:"level"`

	// generated from element "duration" of type positive-divisions order 444
	Duration string `xml:"duration"`
}

// Fingering is generated from named complex type "fingering"
type Fingering struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "substitution" of type yes-no
	Substitution string `xml:"substitution,attr,omitempty"`

	// generated from attribute "alternate" of type yes-no
	Alternate string `xml:"alternate,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// First_fret is generated from named complex type "first-fret"
type First_fret struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "text" of type xs:token
	Text string `xml:"text,attr,omitempty"`

	// generated from attribute "location" of type left-right
	Location string `xml:"location,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText int `xml:",chardata"`
}

// For_part is generated from named complex type "for-part"
type For_part struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "number" of type staff-number
	Number int `xml:"number,attr,omitempty"`

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from element "part-clef" of type part-clef order 53
	Part_clef []*Part_clef `xml:"part-clef"`

	// generated from element "part-transpose" of type part-transpose order 54
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

	// generated from element "footnote" of type formatted-text order 399
	Footnote []*Formatted_text `xml:"footnote"`

	// generated from element "level" of type level order 400
	Level []*Level `xml:"level"`

	// generated from element "staff" of type xs:positiveInteger order 401
	Staff int `xml:"staff"`

	// generated from element "voice" of type xs:string order 409
	Voice string `xml:"voice"`

	// generated from element "duration" of type positive-divisions order 444
	Duration string `xml:"duration"`
}

// Frame is generated from named complex type "frame"
type Frame struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "height" of type tenths
	Height string `xml:"height,attr,omitempty"`

	// generated from attribute "width" of type tenths
	Width string `xml:"width,attr,omitempty"`

	// generated from attribute "unplayed" of type xs:token
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

	// generated from element "frame-strings" of type xs:positiveInteger order 113
	Frame_strings int `xml:"frame-strings"`

	// generated from element "frame-frets" of type xs:positiveInteger order 114
	Frame_frets int `xml:"frame-frets"`

	// generated from element "first-fret" of type first-fret order 115
	First_fret []*First_fret `xml:"first-fret"`

	// generated from element "frame-note" of type frame-note order 116
	Frame_note []*Frame_note `xml:"frame-note"`
}

// Frame_note is generated from named complex type "frame-note"
type Frame_note struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "string" of type string-type order 117
	String []*String_type `xml:"string"`

	// generated from element "fret" of type fret order 118
	Fret []*Fret `xml:"fret"`

	// generated from element "fingering" of type fingering order 119
	Fingering []*Fingering `xml:"fingering"`

	// generated from element "barre" of type barre order 120
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
	Smufl string `xml:"smufl,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Glissando is generated from named complex type "glissando"
type Glissando struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type" of type start-stop
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute "number" of type number-level
	Number int `xml:"number,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Glyph is generated from named complex type "glyph"
type Glyph struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type" of type glyph-type
	Type string `xml:"type,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Grace is generated from named complex type "grace"
type Grace struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "steal-time-previous" of type percent
	Steal_time_previous string `xml:"steal-time-previous,attr,omitempty"`

	// generated from attribute "steal-time-following" of type percent
	Steal_time_following string `xml:"steal-time-following,attr,omitempty"`

	// generated from attribute "make-time" of type divisions
	Make_time string `xml:"make-time,attr,omitempty"`

	// generated from attribute "slash" of type yes-no
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

	// generated from attribute "type" of type start-stop-single
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute "number" of type xs:token
	Number string `xml:"number,attr,omitempty"`

	// generated from attribute "member-of" of type xs:token
	Member_of string `xml:"member-of,attr,omitempty"`

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from element "feature" of type feature order 121
	Feature []*Feature `xml:"feature"`
}

// Hammer_on_pull_off is generated from named complex type "hammer-on-pull-off"
type Hammer_on_pull_off struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type" of type start-stop
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute "number" of type number-level
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

	// generated from attribute "location" of type harmon-closed-location
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

	// generated from element "harmon-closed" of type harmon-closed order 238
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

	// generated from element "natural" of type empty order 239
	Natural string `xml:"natural"`

	// generated from element "artificial" of type empty order 240
	Artificial string `xml:"artificial"`

	// generated from element "base-pitch" of type empty order 241
	Base_pitch string `xml:"base-pitch"`

	// generated from element "touching-pitch" of type empty order 242
	Touching_pitch string `xml:"touching-pitch"`

	// generated from element "sounding-pitch" of type empty order 243
	Sounding_pitch string `xml:"sounding-pitch"`
}

// Harmony is generated from named complex type "harmony"
type Harmony struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type" of type harmony-type
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute "print-frame" of type yes-no
	Print_frame string `xml:"print-frame,attr,omitempty"`

	// generated from attribute "arrangement" of type harmony-arrangement
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

	// generated from element "frame" of type frame order 122
	Frame []*Frame `xml:"frame"`

	// generated from element "offset" of type offset order 123
	Offset []*Offset `xml:"offset"`

	// generated from element "footnote" of type formatted-text order 399
	Footnote []*Formatted_text `xml:"footnote"`

	// generated from element "level" of type level order 400
	Level []*Level `xml:"level"`

	// generated from element "staff" of type xs:positiveInteger order 401
	Staff int `xml:"staff"`

	// generated from element "root" of type root order 430
	Root []*Root `xml:"root"`

	// generated from element "numeral" of type numeral order 431
	Numeral []*Numeral `xml:"numeral"`

	// generated from element "function" of type style-text order 432
	Function []*Style_text `xml:"function"`

	// generated from element "kind" of type kind order 433
	Kind []*Kind `xml:"kind"`

	// generated from element "inversion" of type inversion order 434
	Inversion []*Inversion `xml:"inversion"`

	// generated from element "bass" of type bass order 435
	Bass []*Bass `xml:"bass"`

	// generated from element "degree" of type degree order 436
	Degree []*Degree `xml:"degree"`
}

// Harmony_alter is generated from named complex type "harmony-alter"
type Harmony_alter struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "location" of type left-right
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

	// generated from element "pedal-tuning" of type pedal-tuning order 124
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

	// generated from element "hole-type" of type xs:string order 244
	Hole_type string `xml:"hole-type"`

	// generated from element "hole-closed" of type hole-closed order 245
	Hole_closed []*Hole_closed `xml:"hole-closed"`

	// generated from element "hole-shape" of type xs:string order 246
	Hole_shape string `xml:"hole-shape"`
}

// Hole_closed is generated from named complex type "hole-closed"
type Hole_closed struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "location" of type hole-closed-location
	Location string `xml:"location,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Horizontal_turn is generated from named complex type "horizontal-turn"
type Horizontal_turn struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "slash" of type yes-no
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

	// generated from element "creator" of type typed-text order 183
	Creator []*Typed_text `xml:"creator"`

	// generated from element "rights" of type typed-text order 184
	Rights []*Typed_text `xml:"rights"`

	// generated from element "encoding" of type encoding order 185
	Encoding []*Encoding `xml:"encoding"`

	// generated from element "source" of type xs:string order 186
	Source string `xml:"source"`

	// generated from element "relation" of type typed-text order 187
	Relation []*Typed_text `xml:"relation"`

	// generated from element "miscellaneous" of type miscellaneous order 188
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
	Id string `xml:"id,attr,omitempty"`
}

// Instrument_change is generated from named complex type "instrument-change"
type Instrument_change struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id" of type xs:IDREF
	Id string `xml:"id,attr,omitempty"`

	// generated from element "instrument-sound" of type xs:string order 405
	Instrument_sound string `xml:"instrument-sound"`

	// generated from element "solo" of type empty order 406
	Solo string `xml:"solo"`

	// generated from element "ensemble" of type positive-integer-or-empty order 407
	Ensemble string `xml:"ensemble"`

	// generated from element "virtual-instrument" of type virtual-instrument order 408
	Virtual_instrument []*Virtual_instrument `xml:"virtual-instrument"`
}

// Instrument_link is generated from named complex type "instrument-link"
type Instrument_link struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id" of type xs:IDREF
	Id string `xml:"id,attr,omitempty"`
}

// Interchangeable is generated from named complex type "interchangeable"
type Interchangeable struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "symbol" of type time-symbol
	Symbol string `xml:"symbol,attr,omitempty"`

	// generated from attribute "separator" of type time-separator
	Separator string `xml:"separator,attr,omitempty"`

	// generated from element "time-relation" of type time-relation order 55
	Time_relation string `xml:"time-relation"`

	// generated from element "beats" of type xs:string order 419
	Beats string `xml:"beats"`

	// generated from element "beat-type" of type xs:string order 420
	Beat_type string `xml:"beat-type"`
}

// Inversion is generated from named complex type "inversion"
type Inversion struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "text" of type xs:token
	Text string `xml:"text,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText int `xml:",chardata"`
}

// Key is generated from named complex type "key"
type Key struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "number" of type staff-number
	Number int `xml:"number,attr,omitempty"`

	// generated from attribute group "print-style
	AttributeGroup_print_style

	// generated from attribute group "print-object
	AttributeGroup_print_object

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from element "key-octave" of type key-octave order 56
	Key_octave []*Key_octave `xml:"key-octave"`

	// generated from element "key-step" of type step order 413
	Key_step string `xml:"key-step"`

	// generated from element "key-alter" of type semitones order 414
	Key_alter string `xml:"key-alter"`

	// generated from element "key-accidental" of type key-accidental order 415
	Key_accidental []*Key_accidental `xml:"key-accidental"`

	// generated from element "cancel" of type cancel order 421
	Cancel []*Cancel `xml:"cancel"`

	// generated from element "fifths" of type fifths order 422
	Fifths int `xml:"fifths"`

	// generated from element "mode" of type mode order 423
	Mode string `xml:"mode"`
}

// Key_accidental is generated from named complex type "key-accidental"
type Key_accidental struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "smufl" of type smufl-accidental-glyph-name
	Smufl string `xml:"smufl,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Key_octave is generated from named complex type "key-octave"
type Key_octave struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "number" of type xs:positiveInteger
	Number int `xml:"number,attr,omitempty"`

	// generated from attribute "cancel" of type yes-no
	Cancel string `xml:"cancel,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText int `xml:",chardata"`
}

// Kind is generated from named complex type "kind"
type Kind struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "use-symbols" of type yes-no
	Use_symbols string `xml:"use-symbols,attr,omitempty"`

	// generated from attribute "text" of type xs:token
	Text string `xml:"text,attr,omitempty"`

	// generated from attribute "stack-degrees" of type yes-no
	Stack_degrees string `xml:"stack-degrees,attr,omitempty"`

	// generated from attribute "parentheses-degrees" of type yes-no
	Parentheses_degrees string `xml:"parentheses-degrees,attr,omitempty"`

	// generated from attribute "bracket-degrees" of type yes-no
	Bracket_degrees string `xml:"bracket-degrees,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Level is generated from named complex type "level"
type Level struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "reference" of type yes-no
	Reference string `xml:"reference,attr,omitempty"`

	// generated from attribute "type" of type start-stop-single
	Type string `xml:"type,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Line_detail is generated from named complex type "line-detail"
type Line_detail struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "line" of type staff-line
	Line int `xml:"line,attr,omitempty"`

	// generated from attribute "width" of type tenths
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

	// generated from attribute "type" of type line-width-type
	Type string `xml:"type,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Link is generated from named complex type "link"
type Link struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "name" of type xs:token
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

	// generated from element "assess" of type assess order 247
	Assess []*Assess `xml:"assess"`

	// generated from element "wait" of type wait order 248
	Wait []*Wait `xml:"wait"`

	// generated from element "other-listen" of type other-listening order 249
	Other_listen []*Other_listening `xml:"other-listen"`
}

// Listening is generated from named complex type "listening"
type Listening struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "sync" of type sync order 125
	Sync []*Sync `xml:"sync"`

	// generated from element "other-listening" of type other-listening order 126
	Other_listening []*Other_listening `xml:"other-listening"`

	// generated from element "offset" of type offset order 127
	Offset []*Offset `xml:"offset"`
}

// Lyric is generated from named complex type "lyric"
type Lyric struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "number" of type xs:NMTOKEN
	Number string `xml:"number,attr,omitempty"`

	// generated from attribute "name" of type xs:token
	NameXSD string `xml:"name,attr,omitempty"`

	// generated from attribute "time-only" of type time-only
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

	// generated from element "elision" of type elision order 252
	Elision []*Elision `xml:"elision"`

	// generated from element "syllabic" of type syllabic order 253
	Syllabic string `xml:"syllabic"`

	// generated from element "text" of type text-element-data order 254
	Text []*Text_element_data `xml:"text"`

	// generated from element "extend" of type extend order 255
	Extend []*Extend `xml:"extend"`

	// generated from element "laughing" of type empty order 257
	Laughing string `xml:"laughing"`

	// generated from element "humming" of type empty order 258
	Humming string `xml:"humming"`

	// generated from element "end-line" of type empty order 259
	End_line string `xml:"end-line"`

	// generated from element "end-paragraph" of type empty order 260
	End_paragraph string `xml:"end-paragraph"`

	// generated from element "footnote" of type formatted-text order 399
	Footnote []*Formatted_text `xml:"footnote"`

	// generated from element "level" of type level order 400
	Level []*Level `xml:"level"`
}

// Lyric_font is generated from named complex type "lyric-font"
type Lyric_font struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "number" of type xs:NMTOKEN
	Number string `xml:"number,attr,omitempty"`

	// generated from attribute "name" of type xs:token
	NameXSD string `xml:"name,attr,omitempty"`

	// generated from attribute group "font
	AttributeGroup_font
}

// Lyric_language is generated from named complex type "lyric-language"
type Lyric_language struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "number" of type xs:NMTOKEN
	Number string `xml:"number,attr,omitempty"`

	// generated from attribute "name" of type xs:token
	NameXSD string `xml:"name,attr,omitempty"`

	// generated from attribute "http://www.w3.org/XML/1998/namespace lang" of type 
	Lang string `xml:"http://www.w3.org/XML/1998/namespace lang,attr,omitempty"`
}

// Measure_layout is generated from named complex type "measure-layout"
type Measure_layout struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "measure-distance" of type tenths order 195
	Measure_distance string `xml:"measure-distance"`
}

// Measure_numbering is generated from named complex type "measure-numbering"
type Measure_numbering struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "system" of type system-relation-number
	System string `xml:"system,attr,omitempty"`

	// generated from attribute "staff" of type staff-number
	Staff int `xml:"staff,attr,omitempty"`

	// generated from attribute "multiple-rest-always" of type yes-no
	Multiple_rest_always string `xml:"multiple-rest-always,attr,omitempty"`

	// generated from attribute "multiple-rest-range" of type yes-no
	Multiple_rest_range string `xml:"multiple-rest-range,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Measure_repeat is generated from named complex type "measure-repeat"
type Measure_repeat struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type" of type start-stop
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute "slashes" of type xs:positiveInteger
	Slashes int `xml:"slashes,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Measure_style is generated from named complex type "measure-style"
type Measure_style struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "number" of type staff-number
	Number int `xml:"number,attr,omitempty"`

	// generated from attribute group "font
	AttributeGroup_font

	// generated from attribute group "color
	AttributeGroup_color

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from element "multiple-rest" of type multiple-rest order 57
	Multiple_rest []*Multiple_rest `xml:"multiple-rest"`

	// generated from element "measure-repeat" of type measure-repeat order 58
	Measure_repeat []*Measure_repeat `xml:"measure-repeat"`

	// generated from element "beat-repeat" of type beat-repeat order 59
	Beat_repeat []*Beat_repeat `xml:"beat-repeat"`

	// generated from element "slash" of type slash order 60
	Slash []*Slash `xml:"slash"`
}

// Membrane is generated from named complex type "membrane"
type Membrane struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "smufl" of type smufl-pictogram-glyph-name
	Smufl string `xml:"smufl,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Metal is generated from named complex type "metal"
type Metal struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "smufl" of type smufl-pictogram-glyph-name
	Smufl string `xml:"smufl,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Metronome is generated from named complex type "metronome"
type Metronome struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "parentheses" of type yes-no
	Parentheses string `xml:"parentheses,attr,omitempty"`

	// generated from attribute group "print-style-align
	AttributeGroup_print_style_align

	// generated from attribute group "print-object
	AttributeGroup_print_object

	// generated from attribute group "justify
	AttributeGroup_justify

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from element "per-minute" of type per-minute order 129
	Per_minute []*Per_minute `xml:"per-minute"`

	// generated from element "beat-unit-tied" of type beat-unit-tied order 130
	Beat_unit_tied []*Beat_unit_tied `xml:"beat-unit-tied"`

	// generated from element "metronome-arrows" of type empty order 131
	Metronome_arrows string `xml:"metronome-arrows"`

	// generated from element "metronome-relation" of type xs:string order 133
	Metronome_relation string `xml:"metronome-relation"`

	// generated from element "metronome-note" of type metronome-note order 134
	Metronome_note []*Metronome_note `xml:"metronome-note"`

	// generated from element "beat-unit" of type note-type-value order 428
	Beat_unit string `xml:"beat-unit"`

	// generated from element "beat-unit-dot" of type empty order 429
	Beat_unit_dot string `xml:"beat-unit-dot"`
}

// Metronome_beam is generated from named complex type "metronome-beam"
type Metronome_beam struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "number" of type beam-level
	Number int `xml:"number,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Metronome_note is generated from named complex type "metronome-note"
type Metronome_note struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "metronome-type" of type note-type-value order 135
	Metronome_type string `xml:"metronome-type"`

	// generated from element "metronome-dot" of type empty order 136
	Metronome_dot string `xml:"metronome-dot"`

	// generated from element "metronome-beam" of type metronome-beam order 137
	Metronome_beam []*Metronome_beam `xml:"metronome-beam"`

	// generated from element "metronome-tied" of type metronome-tied order 138
	Metronome_tied []*Metronome_tied `xml:"metronome-tied"`

	// generated from element "metronome-tuplet" of type metronome-tuplet order 139
	Metronome_tuplet []*Metronome_tuplet `xml:"metronome-tuplet"`
}

// Metronome_tied is generated from named complex type "metronome-tied"
type Metronome_tied struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type" of type start-stop
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

	// generated from attribute "port" of type midi-16
	Port int `xml:"port,attr,omitempty"`

	// generated from attribute "id" of type xs:IDREF
	Id string `xml:"id,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Midi_instrument is generated from named complex type "midi-instrument"
type Midi_instrument struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id" of type xs:IDREF
	Id string `xml:"id,attr,omitempty"`

	// generated from element "midi-channel" of type midi-16 order 27
	Midi_channel int `xml:"midi-channel"`

	// generated from element "midi-name" of type xs:string order 28
	Midi_name string `xml:"midi-name"`

	// generated from element "midi-bank" of type midi-16384 order 29
	Midi_bank int `xml:"midi-bank"`

	// generated from element "midi-program" of type midi-128 order 30
	Midi_program int `xml:"midi-program"`

	// generated from element "midi-unpitched" of type midi-128 order 31
	Midi_unpitched int `xml:"midi-unpitched"`

	// generated from element "volume" of type percent order 32
	Volume string `xml:"volume"`

	// generated from element "pan" of type rotation-degrees order 33
	Pan string `xml:"pan"`

	// generated from element "elevation" of type rotation-degrees order 34
	Elevation string `xml:"elevation"`
}

// Miscellaneous is generated from named complex type "miscellaneous"
type Miscellaneous struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "miscellaneous-field" of type miscellaneous-field order 189
	Miscellaneous_field []*Miscellaneous_field `xml:"miscellaneous-field"`
}

// Miscellaneous_field is generated from named complex type "miscellaneous-field"
type Miscellaneous_field struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "name" of type xs:token
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

	// generated from attribute "use-symbols" of type yes-no
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

	// generated from element "display-text" of type formatted-text order 35
	Display_text []*Formatted_text `xml:"display-text"`

	// generated from element "accidental-text" of type accidental-text order 36
	Accidental_text []*Accidental_text `xml:"accidental-text"`
}

// Non_arpeggiate is generated from named complex type "non-arpeggiate"
type Non_arpeggiate struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type" of type top-bottom
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute "number" of type number-level
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

	// generated from element "tied" of type tied order 261
	Tied []*Tied `xml:"tied"`

	// generated from element "slur" of type slur order 262
	Slur []*Slur `xml:"slur"`

	// generated from element "tuplet" of type tuplet order 263
	Tuplet []*Tuplet `xml:"tuplet"`

	// generated from element "glissando" of type glissando order 264
	Glissando []*Glissando `xml:"glissando"`

	// generated from element "slide" of type slide order 265
	Slide []*Slide `xml:"slide"`

	// generated from element "ornaments" of type ornaments order 266
	Ornaments []*Ornaments `xml:"ornaments"`

	// generated from element "technical" of type technical order 267
	Technical []*Technical `xml:"technical"`

	// generated from element "articulations" of type articulations order 268
	Articulations []*Articulations `xml:"articulations"`

	// generated from element "dynamics" of type dynamics order 269
	Dynamics []*Dynamics `xml:"dynamics"`

	// generated from element "fermata" of type fermata order 270
	Fermata []*Fermata `xml:"fermata"`

	// generated from element "arpeggiate" of type arpeggiate order 271
	Arpeggiate []*Arpeggiate `xml:"arpeggiate"`

	// generated from element "non-arpeggiate" of type non-arpeggiate order 272
	Non_arpeggiate []*Non_arpeggiate `xml:"non-arpeggiate"`

	// generated from element "accidental-mark" of type accidental-mark order 273
	Accidental_mark []*Accidental_mark `xml:"accidental-mark"`

	// generated from element "other-notation" of type other-notation order 274
	Other_notation []*Other_notation `xml:"other-notation"`

	// generated from element "footnote" of type formatted-text order 399
	Footnote []*Formatted_text `xml:"footnote"`

	// generated from element "level" of type level order 400
	Level []*Level `xml:"level"`
}

// Note is generated from named complex type "note"
type Note struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "print-leger" of type yes-no
	Print_leger string `xml:"print-leger,attr,omitempty"`

	// generated from attribute "dynamics" of type non-negative-decimal
	Dynamics string `xml:"dynamics,attr,omitempty"`

	// generated from attribute "end-dynamics" of type non-negative-decimal
	End_dynamics string `xml:"end-dynamics,attr,omitempty"`

	// generated from attribute "attack" of type divisions
	Attack string `xml:"attack,attr,omitempty"`

	// generated from attribute "release" of type divisions
	Release string `xml:"release,attr,omitempty"`

	// generated from attribute "time-only" of type time-only
	Time_only string `xml:"time-only,attr,omitempty"`

	// generated from attribute "pizzicato" of type yes-no
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

	// generated from element "grace" of type grace order 275
	Grace []*Grace `xml:"grace"`

	// generated from element "tie" of type tie order 276
	Tie []*Tie `xml:"tie"`

	// generated from element "cue" of type empty order 277
	Cue string `xml:"cue"`

	// generated from element "instrument" of type instrument order 280
	Instrument []*Instrument `xml:"instrument"`

	// generated from element "type" of type note-type order 281
	Type []*Note_type `xml:"type"`

	// generated from element "dot" of type empty-placement order 282
	Dot []*Empty_placement `xml:"dot"`

	// generated from element "accidental" of type accidental order 283
	Accidental []*Accidental `xml:"accidental"`

	// generated from element "time-modification" of type time-modification order 284
	Time_modification []*Time_modification `xml:"time-modification"`

	// generated from element "stem" of type stem order 285
	Stem []*Stem `xml:"stem"`

	// generated from element "notehead" of type notehead order 286
	Notehead []*Notehead `xml:"notehead"`

	// generated from element "notehead-text" of type notehead-text order 287
	Notehead_text []*Notehead_text `xml:"notehead-text"`

	// generated from element "beam" of type beam order 288
	Beam []*Beam `xml:"beam"`

	// generated from element "notations" of type notations order 289
	Notations []*Notations `xml:"notations"`

	// generated from element "lyric" of type lyric order 290
	Lyric []*Lyric `xml:"lyric"`

	// generated from element "play" of type play order 291
	Play []*Play `xml:"play"`

	// generated from element "listen" of type listen order 292
	Listen []*Listen `xml:"listen"`

	// generated from element "footnote" of type formatted-text order 399
	Footnote []*Formatted_text `xml:"footnote"`

	// generated from element "level" of type level order 400
	Level []*Level `xml:"level"`

	// generated from element "staff" of type xs:positiveInteger order 401
	Staff int `xml:"staff"`

	// generated from element "voice" of type xs:string order 409
	Voice string `xml:"voice"`

	// generated from element "duration" of type positive-divisions order 444
	Duration string `xml:"duration"`

	// generated from element "chord" of type empty order 447
	Chord string `xml:"chord"`

	// generated from element "pitch" of type pitch order 448
	Pitch []*Pitch `xml:"pitch"`

	// generated from element "unpitched" of type unpitched order 449
	Unpitched []*Unpitched `xml:"unpitched"`

	// generated from element "rest" of type rest order 450
	Rest []*Rest `xml:"rest"`
}

// Note_size is generated from named complex type "note-size"
type Note_size struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type" of type note-size-type
	Type string `xml:"type,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Note_type is generated from named complex type "note-type"
type Note_type struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "size" of type symbol-size
	Size string `xml:"size,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Notehead is generated from named complex type "notehead"
type Notehead struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "filled" of type yes-no
	Filled string `xml:"filled,attr,omitempty"`

	// generated from attribute "parentheses" of type yes-no
	Parentheses string `xml:"parentheses,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Notehead_text is generated from named complex type "notehead-text"
type Notehead_text struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "display-text" of type formatted-text order 293
	Display_text []*Formatted_text `xml:"display-text"`

	// generated from element "accidental-text" of type accidental-text order 294
	Accidental_text []*Accidental_text `xml:"accidental-text"`
}

// Numeral is generated from named complex type "numeral"
type Numeral struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "numeral-root" of type numeral-root order 140
	Numeral_root []*Numeral_root `xml:"numeral-root"`

	// generated from element "numeral-alter" of type harmony-alter order 141
	Numeral_alter []*Harmony_alter `xml:"numeral-alter"`

	// generated from element "numeral-key" of type numeral-key order 142
	Numeral_key []*Numeral_key `xml:"numeral-key"`
}

// Numeral_key is generated from named complex type "numeral-key"
type Numeral_key struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "print-object
	AttributeGroup_print_object

	// generated from element "numeral-fifths" of type fifths order 143
	Numeral_fifths int `xml:"numeral-fifths"`

	// generated from element "numeral-mode" of type numeral-mode order 144
	Numeral_mode string `xml:"numeral-mode"`
}

// Numeral_root is generated from named complex type "numeral-root"
type Numeral_root struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "text" of type xs:token
	Text string `xml:"text,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText int `xml:",chardata"`
}

// Octave_shift is generated from named complex type "octave-shift"
type Octave_shift struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type" of type up-down-stop-continue
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute "number" of type number-level
	Number int `xml:"number,attr,omitempty"`

	// generated from attribute "size" of type xs:positiveInteger
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

	// generated from attribute "sound" of type yes-no
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

	// generated from element "trill-mark" of type empty-trill-sound order 295
	Trill_mark []*Empty_trill_sound `xml:"trill-mark"`

	// generated from element "turn" of type horizontal-turn order 296
	Turn []*Horizontal_turn `xml:"turn"`

	// generated from element "delayed-turn" of type horizontal-turn order 297
	Delayed_turn []*Horizontal_turn `xml:"delayed-turn"`

	// generated from element "inverted-turn" of type horizontal-turn order 298
	Inverted_turn []*Horizontal_turn `xml:"inverted-turn"`

	// generated from element "delayed-inverted-turn" of type horizontal-turn order 299
	Delayed_inverted_turn []*Horizontal_turn `xml:"delayed-inverted-turn"`

	// generated from element "vertical-turn" of type empty-trill-sound order 300
	Vertical_turn []*Empty_trill_sound `xml:"vertical-turn"`

	// generated from element "inverted-vertical-turn" of type empty-trill-sound order 301
	Inverted_vertical_turn []*Empty_trill_sound `xml:"inverted-vertical-turn"`

	// generated from element "shake" of type empty-trill-sound order 302
	Shake []*Empty_trill_sound `xml:"shake"`

	// generated from element "wavy-line" of type wavy-line order 303
	Wavy_line []*Wavy_line `xml:"wavy-line"`

	// generated from element "mordent" of type mordent order 304
	Mordent []*Mordent `xml:"mordent"`

	// generated from element "inverted-mordent" of type mordent order 305
	Inverted_mordent []*Mordent `xml:"inverted-mordent"`

	// generated from element "schleifer" of type empty-placement order 306
	Schleifer []*Empty_placement `xml:"schleifer"`

	// generated from element "tremolo" of type tremolo order 307
	Tremolo []*Tremolo `xml:"tremolo"`

	// generated from element "haydn" of type empty-trill-sound order 308
	Haydn []*Empty_trill_sound `xml:"haydn"`

	// generated from element "other-ornament" of type other-placement-text order 309
	Other_ornament []*Other_placement_text `xml:"other-ornament"`

	// generated from element "accidental-mark" of type accidental-mark order 310
	Accidental_mark []*Accidental_mark `xml:"accidental-mark"`
}

// Other_appearance is generated from named complex type "other-appearance"
type Other_appearance struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type" of type xs:token
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

	// generated from attribute "type" of type xs:token
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute "player" of type xs:IDREF
	Player string `xml:"player,attr,omitempty"`

	// generated from attribute "time-only" of type time-only
	Time_only string `xml:"time-only,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Other_notation is generated from named complex type "other-notation"
type Other_notation struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type" of type start-stop-single
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute "number" of type number-level
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

	// generated from attribute "type" of type xs:token
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

	// generated from element "page-height" of type tenths order 196
	Page_height string `xml:"page-height"`

	// generated from element "page-width" of type tenths order 197
	Page_width string `xml:"page-width"`

	// generated from element "page-margins" of type page-margins order 198
	Page_margins []*Page_margins `xml:"page-margins"`
}

// Page_margins is generated from named complex type "page-margins"
type Page_margins struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type" of type margin-type
	Type string `xml:"type,attr,omitempty"`

	// generated from element "top-margin" of type tenths order 437
	Top_margin string `xml:"top-margin"`

	// generated from element "bottom-margin" of type tenths order 438
	Bottom_margin string `xml:"bottom-margin"`

	// generated from element "left-margin" of type tenths order 442
	Left_margin string `xml:"left-margin"`

	// generated from element "right-margin" of type tenths order 443
	Right_margin string `xml:"right-margin"`
}

// Part_clef is generated from named complex type "part-clef"
type Part_clef struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "sign" of type clef-sign order 410
	Sign string `xml:"sign"`

	// generated from element "line" of type staff-line-position order 411
	Line int `xml:"line"`

	// generated from element "clef-octave-change" of type xs:integer order 412
	Clef_octave_change int `xml:"clef-octave-change"`
}

// Part_group is generated from named complex type "part-group"
type Part_group struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type" of type start-stop
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute "number" of type xs:token
	Number string `xml:"number,attr,omitempty"`

	// generated from element "group-name" of type group-name order 371
	Group_name []*Group_name `xml:"group-name"`

	// generated from element "group-name-display" of type name-display order 372
	Group_name_display []*Name_display `xml:"group-name-display"`

	// generated from element "group-abbreviation" of type group-name order 373
	Group_abbreviation []*Group_name `xml:"group-abbreviation"`

	// generated from element "group-abbreviation-display" of type name-display order 374
	Group_abbreviation_display []*Name_display `xml:"group-abbreviation-display"`

	// generated from element "group-symbol" of type group-symbol order 375
	Group_symbol []*Group_symbol `xml:"group-symbol"`

	// generated from element "group-barline" of type group-barline order 376
	Group_barline []*Group_barline `xml:"group-barline"`

	// generated from element "group-time" of type empty order 377
	Group_time string `xml:"group-time"`

	// generated from element "footnote" of type formatted-text order 399
	Footnote []*Formatted_text `xml:"footnote"`

	// generated from element "level" of type level order 400
	Level []*Level `xml:"level"`
}

// Part_link is generated from named complex type "part-link"
type Part_link struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "link-attributes
	AttributeGroup_link_attributes

	// generated from element "instrument-link" of type instrument-link order 378
	Instrument_link []*Instrument_link `xml:"instrument-link"`

	// generated from element "group-link" of type xs:string order 379
	Group_link string `xml:"group-link"`
}

// Part_list is generated from named complex type "part-list"
type Part_list struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "part-group" of type part-group order 465
	Part_group []*Part_group `xml:"part-group"`

	// generated from element "score-part" of type score-part order 473
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
	Top_staff int `xml:"top-staff,attr,omitempty"`

	// generated from attribute "bottom-staff" of type staff-number
	Bottom_staff int `xml:"bottom-staff,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Part_transpose is generated from named complex type "part-transpose"
type Part_transpose struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "diatonic" of type xs:integer order 424
	Diatonic int `xml:"diatonic"`

	// generated from element "chromatic" of type semitones order 425
	Chromatic string `xml:"chromatic"`

	// generated from element "octave-change" of type xs:integer order 426
	Octave_change int `xml:"octave-change"`

	// generated from element "double" of type double order 427
	Double []*Double `xml:"double"`
}

// Pedal is generated from named complex type "pedal"
type Pedal struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type" of type pedal-type
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute "number" of type number-level
	Number int `xml:"number,attr,omitempty"`

	// generated from attribute "line" of type yes-no
	Line string `xml:"line,attr,omitempty"`

	// generated from attribute "sign" of type yes-no
	Sign string `xml:"sign,attr,omitempty"`

	// generated from attribute "abbreviated" of type yes-no
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

	// generated from element "pedal-step" of type step order 145
	Pedal_step string `xml:"pedal-step"`

	// generated from element "pedal-alter" of type semitones order 146
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

	// generated from element "glass" of type glass order 147
	Glass []*Glass `xml:"glass"`

	// generated from element "metal" of type metal order 148
	Metal []*Metal `xml:"metal"`

	// generated from element "wood" of type wood order 149
	Wood []*Wood `xml:"wood"`

	// generated from element "pitched" of type pitched order 150
	Pitched []*Pitched `xml:"pitched"`

	// generated from element "membrane" of type membrane order 151
	Membrane []*Membrane `xml:"membrane"`

	// generated from element "effect" of type effect order 152
	Effect []*Effect `xml:"effect"`

	// generated from element "timpani" of type timpani order 153
	Timpani []*Timpani `xml:"timpani"`

	// generated from element "beater" of type beater order 154
	Beater []*Beater `xml:"beater"`

	// generated from element "stick" of type stick order 155
	Stick []*Stick `xml:"stick"`

	// generated from element "stick-location" of type stick-location order 156
	Stick_location string `xml:"stick-location"`

	// generated from element "other-percussion" of type other-text order 157
	Other_percussion []*Other_text `xml:"other-percussion"`
}

// Pitch is generated from named complex type "pitch"
type Pitch struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "step" of type step order 311
	Step string `xml:"step"`

	// generated from element "alter" of type semitones order 312
	Alter string `xml:"alter"`

	// generated from element "octave" of type octave order 313
	Octave int `xml:"octave"`
}

// Pitched is generated from named complex type "pitched"
type Pitched struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "smufl" of type smufl-pictogram-glyph-name
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

	// generated from attribute "id" of type xs:IDREF
	Id string `xml:"id,attr,omitempty"`

	// generated from element "ipa" of type xs:string order 37
	Ipa string `xml:"ipa"`

	// generated from element "mute" of type mute order 38
	Mute string `xml:"mute"`

	// generated from element "semi-pitched" of type semi-pitched order 39
	Semi_pitched string `xml:"semi-pitched"`

	// generated from element "other-play" of type other-play order 40
	Other_play []*Other_play `xml:"other-play"`
}

// Player is generated from named complex type "player"
type Player struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id" of type xs:ID
	Id string `xml:"id,attr,omitempty"`

	// generated from element "player-name" of type xs:string order 380
	Player_name string `xml:"player-name"`
}

// Principal_voice is generated from named complex type "principal-voice"
type Principal_voice struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type" of type start-stop
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute "symbol" of type principal-voice-symbol
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

	// generated from element "measure-layout" of type measure-layout order 158
	Measure_layout []*Measure_layout `xml:"measure-layout"`

	// generated from element "measure-numbering" of type measure-numbering order 159
	Measure_numbering []*Measure_numbering `xml:"measure-numbering"`

	// generated from element "part-name-display" of type name-display order 160
	Part_name_display []*Name_display `xml:"part-name-display"`

	// generated from element "part-abbreviation-display" of type name-display order 161
	Part_abbreviation_display []*Name_display `xml:"part-abbreviation-display"`

	// generated from element "page-layout" of type page-layout order 439
	Page_layout []*Page_layout `xml:"page-layout"`

	// generated from element "system-layout" of type system-layout order 440
	System_layout []*System_layout `xml:"system-layout"`

	// generated from element "staff-layout" of type staff-layout order 441
	Staff_layout []*Staff_layout `xml:"staff-layout"`
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
	Direction string `xml:"direction,attr,omitempty"`

	// generated from attribute "times" of type xs:nonNegativeInteger
	Times int `xml:"times,attr,omitempty"`

	// generated from attribute "after-jump" of type yes-no
	After_jump string `xml:"after-jump,attr,omitempty"`

	// generated from attribute "winged" of type winged
	Winged string `xml:"winged,attr,omitempty"`
}

// Rest is generated from named complex type "rest"
type Rest struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "measure" of type yes-no
	Measure string `xml:"measure,attr,omitempty"`

	// generated from element "display-step" of type step order 445
	Display_step string `xml:"display-step"`

	// generated from element "display-octave" of type octave order 446
	Display_octave int `xml:"display-octave"`
}

// Root is generated from named complex type "root"
type Root struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "root-step" of type root-step order 162
	Root_step []*Root_step `xml:"root-step"`

	// generated from element "root-alter" of type harmony-alter order 163
	Root_alter []*Harmony_alter `xml:"root-alter"`
}

// Root_step is generated from named complex type "root-step"
type Root_step struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "text" of type xs:token
	Text string `xml:"text,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Scaling is generated from named complex type "scaling"
type Scaling struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "millimeters" of type millimeters order 199
	Millimeters string `xml:"millimeters"`

	// generated from element "tenths" of type tenths order 200
	Tenths string `xml:"tenths"`
}

// Scordatura is generated from named complex type "scordatura"
type Scordatura struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from element "accord" of type accord order 164
	Accord []*Accord `xml:"accord"`
}

// Score_instrument is generated from named complex type "score-instrument"
type Score_instrument struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id" of type xs:ID
	Id string `xml:"id,attr,omitempty"`

	// generated from element "instrument-name" of type xs:string order 381
	Instrument_name string `xml:"instrument-name"`

	// generated from element "instrument-abbreviation" of type xs:string order 382
	Instrument_abbreviation string `xml:"instrument-abbreviation"`

	// generated from element "instrument-sound" of type xs:string order 405
	Instrument_sound string `xml:"instrument-sound"`

	// generated from element "solo" of type empty order 406
	Solo string `xml:"solo"`

	// generated from element "ensemble" of type positive-integer-or-empty order 407
	Ensemble string `xml:"ensemble"`

	// generated from element "virtual-instrument" of type virtual-instrument order 408
	Virtual_instrument []*Virtual_instrument `xml:"virtual-instrument"`
}

// Score_part is generated from named complex type "score-part"
type Score_part struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "id" of type xs:ID
	Id string `xml:"id,attr,omitempty"`

	// generated from element "identification" of type identification order 383
	Identification []*Identification `xml:"identification"`

	// generated from element "part-link" of type part-link order 384
	Part_link []*Part_link `xml:"part-link"`

	// generated from element "part-name" of type part-name order 385
	Part_name []*Part_name `xml:"part-name"`

	// generated from element "part-name-display" of type name-display order 386
	Part_name_display []*Name_display `xml:"part-name-display"`

	// generated from element "part-abbreviation" of type part-name order 387
	Part_abbreviation []*Part_name `xml:"part-abbreviation"`

	// generated from element "part-abbreviation-display" of type name-display order 388
	Part_abbreviation_display []*Name_display `xml:"part-abbreviation-display"`

	// generated from element "group" of type xs:string order 389
	Group string `xml:"group"`

	// generated from element "score-instrument" of type score-instrument order 390
	Score_instrument []*Score_instrument `xml:"score-instrument"`

	// generated from element "player" of type player order 391
	Player []*Player `xml:"player"`

	// generated from element "midi-device" of type midi-device order 392
	Midi_device []*Midi_device `xml:"midi-device"`

	// generated from element "midi-instrument" of type midi-instrument order 393
	Midi_instrument []*Midi_instrument `xml:"midi-instrument"`
}

// Segno is generated from named complex type "segno"
type Segno struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "smufl" of type smufl-segno-glyph-name
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

	// generated from attribute "type" of type start-stop
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute "use-dots" of type yes-no
	Use_dots string `xml:"use-dots,attr,omitempty"`

	// generated from attribute "use-stems" of type yes-no
	Use_stems string `xml:"use-stems,attr,omitempty"`

	// generated from element "slash-type" of type note-type-value order 416
	Slash_type string `xml:"slash-type"`

	// generated from element "slash-dot" of type empty order 417
	Slash_dot string `xml:"slash-dot"`

	// generated from element "except-voice" of type xs:string order 418
	Except_voice string `xml:"except-voice"`
}

// Slide is generated from named complex type "slide"
type Slide struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type" of type start-stop
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute "number" of type number-level
	Number int `xml:"number,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Slur is generated from named complex type "slur"
type Slur struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type" of type start-stop-continue
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute "number" of type number-level
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

	// generated from attribute "tempo" of type non-negative-decimal
	Tempo string `xml:"tempo,attr,omitempty"`

	// generated from attribute "dynamics" of type non-negative-decimal
	Dynamics string `xml:"dynamics,attr,omitempty"`

	// generated from attribute "dacapo" of type yes-no
	Dacapo string `xml:"dacapo,attr,omitempty"`

	// generated from attribute "segno" of type xs:token
	Segno string `xml:"segno,attr,omitempty"`

	// generated from attribute "dalsegno" of type xs:token
	Dalsegno string `xml:"dalsegno,attr,omitempty"`

	// generated from attribute "coda" of type xs:token
	Coda string `xml:"coda,attr,omitempty"`

	// generated from attribute "tocoda" of type xs:token
	Tocoda string `xml:"tocoda,attr,omitempty"`

	// generated from attribute "divisions" of type divisions
	Divisions string `xml:"divisions,attr,omitempty"`

	// generated from attribute "forward-repeat" of type yes-no
	Forward_repeat string `xml:"forward-repeat,attr,omitempty"`

	// generated from attribute "fine" of type xs:token
	Fine string `xml:"fine,attr,omitempty"`

	// generated from attribute "time-only" of type time-only
	Time_only string `xml:"time-only,attr,omitempty"`

	// generated from attribute "pizzicato" of type yes-no
	Pizzicato string `xml:"pizzicato,attr,omitempty"`

	// generated from attribute "pan" of type rotation-degrees
	Pan string `xml:"pan,attr,omitempty"`

	// generated from attribute "elevation" of type rotation-degrees
	Elevation string `xml:"elevation,attr,omitempty"`

	// generated from attribute "damper-pedal" of type yes-no-number
	Damper_pedal string `xml:"damper-pedal,attr,omitempty"`

	// generated from attribute "soft-pedal" of type yes-no-number
	Soft_pedal string `xml:"soft-pedal,attr,omitempty"`

	// generated from attribute "sostenuto-pedal" of type yes-no-number
	Sostenuto_pedal string `xml:"sostenuto-pedal,attr,omitempty"`

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from element "instrument-change" of type instrument-change order 165
	Instrument_change []*Instrument_change `xml:"instrument-change"`

	// generated from element "midi-device" of type midi-device order 166
	Midi_device []*Midi_device `xml:"midi-device"`

	// generated from element "midi-instrument" of type midi-instrument order 167
	Midi_instrument []*Midi_instrument `xml:"midi-instrument"`

	// generated from element "play" of type play order 168
	Play []*Play `xml:"play"`

	// generated from element "swing" of type swing order 169
	Swing []*Swing `xml:"swing"`

	// generated from element "offset" of type offset order 170
	Offset []*Offset `xml:"offset"`
}

// Staff_details is generated from named complex type "staff-details"
type Staff_details struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "number" of type staff-number
	Number int `xml:"number,attr,omitempty"`

	// generated from attribute "show-frets" of type show-frets
	Show_frets string `xml:"show-frets,attr,omitempty"`

	// generated from attribute group "print-object
	AttributeGroup_print_object

	// generated from attribute group "print-spacing
	AttributeGroup_print_spacing

	// generated from element "staff-type" of type staff-type order 61
	Staff_type string `xml:"staff-type"`

	// generated from element "staff-lines" of type xs:nonNegativeInteger order 62
	Staff_lines int `xml:"staff-lines"`

	// generated from element "line-detail" of type line-detail order 63
	Line_detail []*Line_detail `xml:"line-detail"`

	// generated from element "staff-tuning" of type staff-tuning order 64
	Staff_tuning []*Staff_tuning `xml:"staff-tuning"`

	// generated from element "capo" of type xs:nonNegativeInteger order 65
	Capo int `xml:"capo"`

	// generated from element "staff-size" of type staff-size order 66
	Staff_size []*Staff_size `xml:"staff-size"`
}

// Staff_divide is generated from named complex type "staff-divide"
type Staff_divide struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type" of type staff-divide-symbol
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

	// generated from attribute "number" of type staff-number
	Number int `xml:"number,attr,omitempty"`

	// generated from element "staff-distance" of type tenths order 201
	Staff_distance string `xml:"staff-distance"`
}

// Staff_size is generated from named complex type "staff-size"
type Staff_size struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "scaling" of type non-negative-decimal
	Scaling string `xml:"scaling,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Staff_tuning is generated from named complex type "staff-tuning"
type Staff_tuning struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "line" of type staff-line
	Line int `xml:"line,attr,omitempty"`

	// generated from element "tuning-step" of type step order 402
	Tuning_step string `xml:"tuning-step"`

	// generated from element "tuning-alter" of type semitones order 403
	Tuning_alter string `xml:"tuning-alter"`

	// generated from element "tuning-octave" of type octave order 404
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
	Tip string `xml:"tip,attr,omitempty"`

	// generated from attribute "parentheses" of type yes-no
	Parentheses string `xml:"parentheses,attr,omitempty"`

	// generated from attribute "dashed-circle" of type yes-no
	Dashed_circle string `xml:"dashed-circle,attr,omitempty"`

	// generated from element "stick-type" of type stick-type order 171
	Stick_type string `xml:"stick-type"`

	// generated from element "stick-material" of type stick-material order 172
	Stick_material string `xml:"stick-material"`
}

// String_mute is generated from named complex type "string-mute"
type String_mute struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type" of type on-off
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

	// generated from attribute "type" of type yes-no
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute "element" of type xs:NMTOKEN
	Element string `xml:"element,attr,omitempty"`

	// generated from attribute "attribute" of type xs:NMTOKEN
	Attribute string `xml:"attribute,attr,omitempty"`

	// generated from attribute "value" of type xs:token
	Value string `xml:"value,attr,omitempty"`
}

// Swing is generated from named complex type "swing"
type Swing struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "straight" of type empty order 173
	Straight string `xml:"straight"`

	// generated from element "first" of type xs:positiveInteger order 174
	First int `xml:"first"`

	// generated from element "second" of type xs:positiveInteger order 175
	Second int `xml:"second"`

	// generated from element "swing-type" of type swing-type-value order 176
	Swing_type string `xml:"swing-type"`

	// generated from element "swing-style" of type xs:string order 177
	Swing_style string `xml:"swing-style"`
}

// Sync is generated from named complex type "sync"
type Sync struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type" of type sync-type
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute "latency" of type milliseconds
	Latency int `xml:"latency,attr,omitempty"`

	// generated from attribute "player" of type xs:IDREF
	Player string `xml:"player,attr,omitempty"`

	// generated from attribute "time-only" of type time-only
	Time_only string `xml:"time-only,attr,omitempty"`
}

// System_dividers is generated from named complex type "system-dividers"
type System_dividers struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "left-divider" of type empty-print-object-style-align order 202
	Left_divider []*Empty_print_object_style_align `xml:"left-divider"`

	// generated from element "right-divider" of type empty-print-object-style-align order 203
	Right_divider []*Empty_print_object_style_align `xml:"right-divider"`
}

// System_layout is generated from named complex type "system-layout"
type System_layout struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "system-margins" of type system-margins order 204
	System_margins []*System_margins `xml:"system-margins"`

	// generated from element "system-distance" of type tenths order 205
	System_distance string `xml:"system-distance"`

	// generated from element "top-system-distance" of type tenths order 206
	Top_system_distance string `xml:"top-system-distance"`

	// generated from element "system-dividers" of type system-dividers order 207
	System_dividers []*System_dividers `xml:"system-dividers"`
}

// System_margins is generated from named complex type "system-margins"
type System_margins struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "left-margin" of type tenths order 442
	Left_margin string `xml:"left-margin"`

	// generated from element "right-margin" of type tenths order 443
	Right_margin string `xml:"right-margin"`
}

// Tap is generated from named complex type "tap"
type Tap struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "hand" of type tap-hand
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

	// generated from element "up-bow" of type empty-placement order 314
	Up_bow []*Empty_placement `xml:"up-bow"`

	// generated from element "down-bow" of type empty-placement order 315
	Down_bow []*Empty_placement `xml:"down-bow"`

	// generated from element "harmonic" of type harmonic order 316
	Harmonic []*Harmonic `xml:"harmonic"`

	// generated from element "open-string" of type empty-placement order 317
	Open_string []*Empty_placement `xml:"open-string"`

	// generated from element "thumb-position" of type empty-placement order 318
	Thumb_position []*Empty_placement `xml:"thumb-position"`

	// generated from element "fingering" of type fingering order 319
	Fingering []*Fingering `xml:"fingering"`

	// generated from element "pluck" of type placement-text order 320
	Pluck []*Placement_text `xml:"pluck"`

	// generated from element "double-tongue" of type empty-placement order 321
	Double_tongue []*Empty_placement `xml:"double-tongue"`

	// generated from element "triple-tongue" of type empty-placement order 322
	Triple_tongue []*Empty_placement `xml:"triple-tongue"`

	// generated from element "stopped" of type empty-placement-smufl order 323
	Stopped []*Empty_placement_smufl `xml:"stopped"`

	// generated from element "snap-pizzicato" of type empty-placement order 324
	Snap_pizzicato []*Empty_placement `xml:"snap-pizzicato"`

	// generated from element "fret" of type fret order 325
	Fret []*Fret `xml:"fret"`

	// generated from element "string" of type string-type order 326
	String []*String_type `xml:"string"`

	// generated from element "hammer-on" of type hammer-on-pull-off order 327
	Hammer_on []*Hammer_on_pull_off `xml:"hammer-on"`

	// generated from element "pull-off" of type hammer-on-pull-off order 328
	Pull_off []*Hammer_on_pull_off `xml:"pull-off"`

	// generated from element "bend" of type bend order 329
	Bend []*Bend `xml:"bend"`

	// generated from element "tap" of type tap order 330
	Tap []*Tap `xml:"tap"`

	// generated from element "heel" of type heel-toe order 331
	Heel []*Heel_toe `xml:"heel"`

	// generated from element "toe" of type heel-toe order 332
	Toe []*Heel_toe `xml:"toe"`

	// generated from element "fingernails" of type empty-placement order 333
	Fingernails []*Empty_placement `xml:"fingernails"`

	// generated from element "hole" of type hole order 334
	Hole []*Hole `xml:"hole"`

	// generated from element "arrow" of type arrow order 335
	Arrow []*Arrow `xml:"arrow"`

	// generated from element "handbell" of type handbell order 336
	Handbell []*Handbell `xml:"handbell"`

	// generated from element "brass-bend" of type empty-placement order 337
	Brass_bend []*Empty_placement `xml:"brass-bend"`

	// generated from element "flip" of type empty-placement order 338
	Flip []*Empty_placement `xml:"flip"`

	// generated from element "smear" of type empty-placement order 339
	Smear []*Empty_placement `xml:"smear"`

	// generated from element "open" of type empty-placement-smufl order 340
	Open []*Empty_placement_smufl `xml:"open"`

	// generated from element "half-muted" of type empty-placement-smufl order 341
	Half_muted []*Empty_placement_smufl `xml:"half-muted"`

	// generated from element "harmon-mute" of type harmon-mute order 342
	Harmon_mute []*Harmon_mute `xml:"harmon-mute"`

	// generated from element "golpe" of type empty-placement order 343
	Golpe []*Empty_placement `xml:"golpe"`

	// generated from element "other-technical" of type other-placement-text order 344
	Other_technical []*Other_placement_text `xml:"other-technical"`
}

// Text_element_data is generated from named complex type "text-element-data"
type Text_element_data struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "http://www.w3.org/XML/1998/namespace lang" of type 
	Lang string `xml:"http://www.w3.org/XML/1998/namespace lang,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Tie is generated from named complex type "tie"
type Tie struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type" of type start-stop
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute "time-only" of type time-only
	Time_only string `xml:"time-only,attr,omitempty"`
}

// Tied is generated from named complex type "tied"
type Tied struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type" of type tied-type
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute "number" of type number-level
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

	// generated from attribute "number" of type staff-number
	Number int `xml:"number,attr,omitempty"`

	// generated from attribute "symbol" of type time-symbol
	Symbol string `xml:"symbol,attr,omitempty"`

	// generated from attribute "separator" of type time-separator
	Separator string `xml:"separator,attr,omitempty"`

	// generated from attribute group "print-style-align
	AttributeGroup_print_style_align

	// generated from attribute group "print-object
	AttributeGroup_print_object

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from element "interchangeable" of type interchangeable order 67
	Interchangeable []*Interchangeable `xml:"interchangeable"`

	// generated from element "senza-misura" of type xs:string order 68
	Senza_misura string `xml:"senza-misura"`

	// generated from element "beats" of type xs:string order 419
	Beats string `xml:"beats"`

	// generated from element "beat-type" of type xs:string order 420
	Beat_type string `xml:"beat-type"`
}

// Time_modification is generated from named complex type "time-modification"
type Time_modification struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "actual-notes" of type xs:nonNegativeInteger order 345
	Actual_notes int `xml:"actual-notes"`

	// generated from element "normal-notes" of type xs:nonNegativeInteger order 346
	Normal_notes int `xml:"normal-notes"`

	// generated from element "normal-type" of type note-type-value order 347
	Normal_type string `xml:"normal-type"`

	// generated from element "normal-dot" of type empty order 348
	Normal_dot string `xml:"normal-dot"`
}

// Timpani is generated from named complex type "timpani"
type Timpani struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "smufl" of type smufl-pictogram-glyph-name
	Smufl string `xml:"smufl,attr,omitempty"`
}

// Transpose is generated from named complex type "transpose"
type Transpose struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "number" of type staff-number
	Number int `xml:"number,attr,omitempty"`

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from element "diatonic" of type xs:integer order 424
	Diatonic int `xml:"diatonic"`

	// generated from element "chromatic" of type semitones order 425
	Chromatic string `xml:"chromatic"`

	// generated from element "octave-change" of type xs:integer order 426
	Octave_change int `xml:"octave-change"`

	// generated from element "double" of type double order 427
	Double []*Double `xml:"double"`
}

// Tremolo is generated from named complex type "tremolo"
type Tremolo struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type" of type tremolo-type
	Type string `xml:"type,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText int `xml:",chardata"`
}

// Tuplet is generated from named complex type "tuplet"
type Tuplet struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type" of type start-stop
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute "number" of type number-level
	Number int `xml:"number,attr,omitempty"`

	// generated from attribute "bracket" of type yes-no
	Bracket string `xml:"bracket,attr,omitempty"`

	// generated from attribute "show-number" of type show-tuplet
	Show_number string `xml:"show-number,attr,omitempty"`

	// generated from attribute "show-type" of type show-tuplet
	Show_type string `xml:"show-type,attr,omitempty"`

	// generated from attribute group "line-shape
	AttributeGroup_line_shape

	// generated from attribute group "position
	AttributeGroup_position

	// generated from attribute group "placement
	AttributeGroup_placement

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id

	// generated from element "tuplet-actual" of type tuplet-portion order 349
	Tuplet_actual []*Tuplet_portion `xml:"tuplet-actual"`

	// generated from element "tuplet-normal" of type tuplet-portion order 350
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

	// generated from element "tuplet-number" of type tuplet-number order 351
	Tuplet_number []*Tuplet_number `xml:"tuplet-number"`

	// generated from element "tuplet-type" of type tuplet-type order 352
	Tuplet_type []*Tuplet_type `xml:"tuplet-type"`

	// generated from element "tuplet-dot" of type tuplet-dot order 353
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
	Type string `xml:"type,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Unpitched is generated from named complex type "unpitched"
type Unpitched struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "display-step" of type step order 445
	Display_step string `xml:"display-step"`

	// generated from element "display-octave" of type octave order 446
	Display_octave int `xml:"display-octave"`
}

// Virtual_instrument is generated from named complex type "virtual-instrument"
type Virtual_instrument struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "virtual-library" of type xs:string order 394
	Virtual_library string `xml:"virtual-library"`

	// generated from element "virtual-name" of type xs:string order 395
	Virtual_name string `xml:"virtual-name"`
}

// Wait is generated from named complex type "wait"
type Wait struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "player" of type xs:IDREF
	Player string `xml:"player,attr,omitempty"`

	// generated from attribute "time-only" of type time-only
	Time_only string `xml:"time-only,attr,omitempty"`
}

// Wavy_line is generated from named complex type "wavy-line"
type Wavy_line struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "type" of type start-stop-continue
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute "number" of type number-level
	Number int `xml:"number,attr,omitempty"`

	// generated from attribute "smufl" of type smufl-wavy-line-glyph-name
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

	// generated from attribute "type" of type wedge-type
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute "number" of type number-level
	Number int `xml:"number,attr,omitempty"`

	// generated from attribute "spread" of type tenths
	Spread string `xml:"spread,attr,omitempty"`

	// generated from attribute "niente" of type yes-no
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

	// generated from attribute "smufl" of type smufl-pictogram-glyph-name
	Smufl string `xml:"smufl,attr,omitempty"`

	// in case the extension has base type xs:string, one has to had the chardata stuff
	EnclosedText string `xml:",chardata"`
}

// Work is generated from named complex type "work"
type Work struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from element "work-number" of type xs:string order 396
	Work_number string `xml:"work-number"`

	// generated from element "work-title" of type xs:string order 397
	Work_title string `xml:"work-title"`

	// generated from element "opus" of type opus order 398
	Opus []*Opus `xml:"opus"`
}

// Group_all_margins is generated from named group "all-margins"
type Group_all_margins struct {

	// insertion point for fields

	// generated from element "top-margin" of type tenths order 437
	Top_margin string `xml:"top-margin"`

	// generated from element "bottom-margin" of type tenths order 438
	Bottom_margin string `xml:"bottom-margin"`

	// generated from element "left-margin" of type tenths order 442
	Left_margin string `xml:"left-margin"`

	// generated from element "right-margin" of type tenths order 443
	Right_margin string `xml:"right-margin"`
}

// Group_beat_unit is generated from named group "beat-unit"
type Group_beat_unit struct {

	// insertion point for fields

	// generated from element "beat-unit" of type note-type-value order 428
	Beat_unit string `xml:"beat-unit"`

	// generated from element "beat-unit-dot" of type empty order 429
	Beat_unit_dot string `xml:"beat-unit-dot"`
}

// Group_clef is generated from named group "clef"
type Group_clef struct {

	// insertion point for fields

	// generated from element "sign" of type clef-sign order 410
	Sign string `xml:"sign"`

	// generated from element "line" of type staff-line-position order 411
	Line int `xml:"line"`

	// generated from element "clef-octave-change" of type xs:integer order 412
	Clef_octave_change int `xml:"clef-octave-change"`
}

// Group_display_step_octave is generated from named group "display-step-octave"
type Group_display_step_octave struct {

	// insertion point for fields

	// generated from element "display-step" of type step order 445
	Display_step string `xml:"display-step"`

	// generated from element "display-octave" of type octave order 446
	Display_octave int `xml:"display-octave"`
}

// Group_duration is generated from named group "duration"
type Group_duration struct {

	// insertion point for fields

	// generated from element "duration" of type positive-divisions order 444
	Duration string `xml:"duration"`
}

// Group_editorial is generated from named group "editorial"
type Group_editorial struct {

	// insertion point for fields

	// generated from element "footnote" of type formatted-text order 399
	Footnote []*Formatted_text `xml:"footnote"`

	// generated from element "level" of type level order 400
	Level []*Level `xml:"level"`
}

// Group_editorial_voice is generated from named group "editorial-voice"
type Group_editorial_voice struct {

	// insertion point for fields

	// generated from element "footnote" of type formatted-text order 399
	Footnote []*Formatted_text `xml:"footnote"`

	// generated from element "level" of type level order 400
	Level []*Level `xml:"level"`

	// generated from element "voice" of type xs:string order 409
	Voice string `xml:"voice"`
}

// Group_editorial_voice_direction is generated from named group "editorial-voice-direction"
type Group_editorial_voice_direction struct {

	// insertion point for fields

	// generated from element "footnote" of type formatted-text order 399
	Footnote []*Formatted_text `xml:"footnote"`

	// generated from element "level" of type level order 400
	Level []*Level `xml:"level"`

	// generated from element "voice" of type xs:string order 409
	Voice string `xml:"voice"`
}

// Group_footnote is generated from named group "footnote"
type Group_footnote struct {

	// insertion point for fields

	// generated from element "footnote" of type formatted-text order 399
	Footnote []*Formatted_text `xml:"footnote"`
}

// Group_full_note is generated from named group "full-note"
type Group_full_note struct {

	// insertion point for fields

	// generated from element "chord" of type empty order 447
	Chord string `xml:"chord"`

	// generated from element "pitch" of type pitch order 448
	Pitch []*Pitch `xml:"pitch"`

	// generated from element "unpitched" of type unpitched order 449
	Unpitched []*Unpitched `xml:"unpitched"`

	// generated from element "rest" of type rest order 450
	Rest []*Rest `xml:"rest"`
}

// Group_harmony_chord is generated from named group "harmony-chord"
type Group_harmony_chord struct {

	// insertion point for fields

	// generated from element "root" of type root order 430
	Root []*Root `xml:"root"`

	// generated from element "numeral" of type numeral order 431
	Numeral []*Numeral `xml:"numeral"`

	// generated from element "function" of type style-text order 432
	Function []*Style_text `xml:"function"`

	// generated from element "kind" of type kind order 433
	Kind []*Kind `xml:"kind"`

	// generated from element "inversion" of type inversion order 434
	Inversion []*Inversion `xml:"inversion"`

	// generated from element "bass" of type bass order 435
	Bass []*Bass `xml:"bass"`

	// generated from element "degree" of type degree order 436
	Degree []*Degree `xml:"degree"`
}

// Group_layout is generated from named group "layout"
type Group_layout struct {

	// insertion point for fields

	// generated from element "page-layout" of type page-layout order 439
	Page_layout []*Page_layout `xml:"page-layout"`

	// generated from element "system-layout" of type system-layout order 440
	System_layout []*System_layout `xml:"system-layout"`

	// generated from element "staff-layout" of type staff-layout order 441
	Staff_layout []*Staff_layout `xml:"staff-layout"`
}

// Group_left_right_margins is generated from named group "left-right-margins"
type Group_left_right_margins struct {

	// insertion point for fields

	// generated from element "left-margin" of type tenths order 442
	Left_margin string `xml:"left-margin"`

	// generated from element "right-margin" of type tenths order 443
	Right_margin string `xml:"right-margin"`
}

// Group_level is generated from named group "level"
type Group_level struct {

	// insertion point for fields

	// generated from element "level" of type level order 400
	Level []*Level `xml:"level"`
}

// Group_music_data is generated from named group "music-data"
type Group_music_data struct {

	// insertion point for fields

	// generated from element "note" of type note order 451
	Note []*Note `xml:"note"`

	// generated from element "backup" of type backup order 452
	Backup []*Backup `xml:"backup"`

	// generated from element "forward" of type forward order 453
	Forward []*Forward `xml:"forward"`

	// generated from element "direction" of type direction order 454
	Direction []*Direction `xml:"direction"`

	// generated from element "attributes" of type attributes order 455
	Attributes []*Attributes `xml:"attributes"`

	// generated from element "harmony" of type harmony order 456
	Harmony []*Harmony `xml:"harmony"`

	// generated from element "figured-bass" of type figured-bass order 457
	Figured_bass []*Figured_bass `xml:"figured-bass"`

	// generated from element "print" of type print order 458
	Print []*Print `xml:"print"`

	// generated from element "sound" of type sound order 459
	Sound []*Sound `xml:"sound"`

	// generated from element "listening" of type listening order 460
	Listening []*Listening `xml:"listening"`

	// generated from element "barline" of type barline order 461
	Barline []*Barline `xml:"barline"`

	// generated from element "grouping" of type grouping order 462
	Grouping []*Grouping `xml:"grouping"`

	// generated from element "link" of type link order 463
	Link []*Link `xml:"link"`

	// generated from element "bookmark" of type bookmark order 464
	Bookmark []*Bookmark `xml:"bookmark"`
}

// Group_non_traditional_key is generated from named group "non-traditional-key"
type Group_non_traditional_key struct {

	// insertion point for fields

	// generated from element "key-step" of type step order 413
	Key_step string `xml:"key-step"`

	// generated from element "key-alter" of type semitones order 414
	Key_alter string `xml:"key-alter"`

	// generated from element "key-accidental" of type key-accidental order 415
	Key_accidental []*Key_accidental `xml:"key-accidental"`
}

// Group_part_group is generated from named group "part-group"
type Group_part_group struct {

	// insertion point for fields

	// generated from element "part-group" of type part-group order 465
	Part_group []*Part_group `xml:"part-group"`
}

// Group_score_header is generated from named group "score-header"
type Group_score_header struct {

	// insertion point for fields

	// generated from element "work" of type work order 466
	Work []*Work `xml:"work"`

	// generated from element "movement-number" of type xs:string order 467
	Movement_number string `xml:"movement-number"`

	// generated from element "movement-title" of type xs:string order 468
	Movement_title string `xml:"movement-title"`

	// generated from element "identification" of type identification order 469
	Identification []*Identification `xml:"identification"`

	// generated from element "defaults" of type defaults order 470
	Defaults []*Defaults `xml:"defaults"`

	// generated from element "credit" of type credit order 471
	Credit []*Credit `xml:"credit"`

	// generated from element "part-list" of type part-list order 472
	Part_list []*Part_list `xml:"part-list"`
}

// Group_score_part is generated from named group "score-part"
type Group_score_part struct {

	// insertion point for fields

	// generated from element "score-part" of type score-part order 473
	Score_part []*Score_part `xml:"score-part"`
}

// Group_slash is generated from named group "slash"
type Group_slash struct {

	// insertion point for fields

	// generated from element "slash-type" of type note-type-value order 416
	Slash_type string `xml:"slash-type"`

	// generated from element "slash-dot" of type empty order 417
	Slash_dot string `xml:"slash-dot"`

	// generated from element "except-voice" of type xs:string order 418
	Except_voice string `xml:"except-voice"`
}

// Group_staff is generated from named group "staff"
type Group_staff struct {

	// insertion point for fields

	// generated from element "staff" of type xs:positiveInteger order 401
	Staff int `xml:"staff"`
}

// Group_time_signature is generated from named group "time-signature"
type Group_time_signature struct {

	// insertion point for fields

	// generated from element "beats" of type xs:string order 419
	Beats string `xml:"beats"`

	// generated from element "beat-type" of type xs:string order 420
	Beat_type string `xml:"beat-type"`
}

// Group_traditional_key is generated from named group "traditional-key"
type Group_traditional_key struct {

	// insertion point for fields

	// generated from element "cancel" of type cancel order 421
	Cancel []*Cancel `xml:"cancel"`

	// generated from element "fifths" of type fifths order 422
	Fifths int `xml:"fifths"`

	// generated from element "mode" of type mode order 423
	Mode string `xml:"mode"`
}

// Group_transpose is generated from named group "transpose"
type Group_transpose struct {

	// insertion point for fields

	// generated from element "diatonic" of type xs:integer order 424
	Diatonic int `xml:"diatonic"`

	// generated from element "chromatic" of type semitones order 425
	Chromatic string `xml:"chromatic"`

	// generated from element "octave-change" of type xs:integer order 426
	Octave_change int `xml:"octave-change"`

	// generated from element "double" of type double order 427
	Double []*Double `xml:"double"`
}

// Group_tuning is generated from named group "tuning"
type Group_tuning struct {

	// insertion point for fields

	// generated from element "tuning-step" of type step order 402
	Tuning_step string `xml:"tuning-step"`

	// generated from element "tuning-alter" of type semitones order 403
	Tuning_alter string `xml:"tuning-alter"`

	// generated from element "tuning-octave" of type octave order 404
	Tuning_octave int `xml:"tuning-octave"`
}

// Group_virtual_instrument_data is generated from named group "virtual-instrument-data"
type Group_virtual_instrument_data struct {

	// insertion point for fields

	// generated from element "instrument-sound" of type xs:string order 405
	Instrument_sound string `xml:"instrument-sound"`

	// generated from element "solo" of type empty order 406
	Solo string `xml:"solo"`

	// generated from element "ensemble" of type positive-integer-or-empty order 407
	Ensemble string `xml:"ensemble"`

	// generated from element "virtual-instrument" of type virtual-instrument order 408
	Virtual_instrument []*Virtual_instrument `xml:"virtual-instrument"`
}

// Group_voice is generated from named group "voice"
type Group_voice struct {

	// insertion point for fields

	// generated from element "voice" of type xs:string order 409
	Voice string `xml:"voice"`
}

// AttributeGroup_bend_sound is generated from named attribute group "bend-sound"
type AttributeGroup_bend_sound struct {

	// insertion point for fields

	// generated from attribute "accelerate" of type yes-no
	Accelerate string `xml:"accelerate,attr,omitempty"`

	// generated from attribute "beats" of type trill-beats
	Beats string `xml:"beats,attr,omitempty"`

	// generated from attribute "first-beat" of type percent
	First_beat string `xml:"first-beat,attr,omitempty"`

	// generated from attribute "last-beat" of type percent
	Last_beat string `xml:"last-beat,attr,omitempty"`
}

// AttributeGroup_bezier is generated from named attribute group "bezier"
type AttributeGroup_bezier struct {

	// insertion point for fields

	// generated from attribute "bezier-x" of type tenths
	Bezier_x string `xml:"bezier-x,attr,omitempty"`

	// generated from attribute "bezier-y" of type tenths
	Bezier_y string `xml:"bezier-y,attr,omitempty"`

	// generated from attribute "bezier-x2" of type tenths
	Bezier_x2 string `xml:"bezier-x2,attr,omitempty"`

	// generated from attribute "bezier-y2" of type tenths
	Bezier_y2 string `xml:"bezier-y2,attr,omitempty"`

	// generated from attribute "bezier-offset" of type divisions
	Bezier_offset string `xml:"bezier-offset,attr,omitempty"`

	// generated from attribute "bezier-offset2" of type divisions
	Bezier_offset2 string `xml:"bezier-offset2,attr,omitempty"`
}

// AttributeGroup_color is generated from named attribute group "color"
type AttributeGroup_color struct {

	// insertion point for fields

	// generated from attribute "color" of type color
	Color string `xml:"color,attr,omitempty"`
}

// AttributeGroup_dashed_formatting is generated from named attribute group "dashed-formatting"
type AttributeGroup_dashed_formatting struct {

	// insertion point for fields

	// generated from attribute "dash-length" of type tenths
	Dash_length string `xml:"dash-length,attr,omitempty"`

	// generated from attribute "space-length" of type tenths
	Space_length string `xml:"space-length,attr,omitempty"`
}

// AttributeGroup_directive is generated from named attribute group "directive"
type AttributeGroup_directive struct {

	// insertion point for fields

	// generated from attribute "directive" of type yes-no
	Directive string `xml:"directive,attr,omitempty"`
}

// AttributeGroup_document_attributes is generated from named attribute group "document-attributes"
type AttributeGroup_document_attributes struct {

	// insertion point for fields

	// generated from attribute "version" of type xs:token
	Version string `xml:"version,attr,omitempty"`
}

// AttributeGroup_element_position is generated from named attribute group "element-position"
type AttributeGroup_element_position struct {

	// insertion point for fields

	// generated from attribute "element" of type xs:NMTOKEN
	Element string `xml:"element,attr,omitempty"`

	// generated from attribute "position" of type xs:positiveInteger
	Position int `xml:"position,attr,omitempty"`
}

// AttributeGroup_enclosure is generated from named attribute group "enclosure"
type AttributeGroup_enclosure struct {

	// insertion point for fields

	// generated from attribute "enclosure" of type enclosure-shape
	Enclosure string `xml:"enclosure,attr,omitempty"`
}

// AttributeGroup_font is generated from named attribute group "font"
type AttributeGroup_font struct {

	// insertion point for fields

	// generated from attribute "font-family" of type font-family
	Font_family string `xml:"font-family,attr,omitempty"`

	// generated from attribute "font-style" of type font-style
	Font_style string `xml:"font-style,attr,omitempty"`

	// generated from attribute "font-size" of type font-size
	Font_size string `xml:"font-size,attr,omitempty"`

	// generated from attribute "font-weight" of type font-weight
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

	// generated from attribute "halign" of type left-center-right
	Halign string `xml:"halign,attr,omitempty"`
}

// AttributeGroup_image_attributes is generated from named attribute group "image-attributes"
type AttributeGroup_image_attributes struct {

	// insertion point for fields

	// generated from attribute "source" of type xs:anyURI
	Source string `xml:"source,attr,omitempty"`

	// generated from attribute "type" of type xs:token
	Type string `xml:"type,attr,omitempty"`

	// generated from attribute "height" of type tenths
	Height string `xml:"height,attr,omitempty"`

	// generated from attribute "width" of type tenths
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

	// generated from attribute "justify" of type left-center-right
	Justify string `xml:"justify,attr,omitempty"`
}

// AttributeGroup_letter_spacing is generated from named attribute group "letter-spacing"
type AttributeGroup_letter_spacing struct {

	// insertion point for fields

	// generated from attribute "letter-spacing" of type number-or-normal
	Letter_spacing string `xml:"letter-spacing,attr,omitempty"`
}

// AttributeGroup_level_display is generated from named attribute group "level-display"
type AttributeGroup_level_display struct {

	// insertion point for fields

	// generated from attribute "parentheses" of type yes-no
	Parentheses string `xml:"parentheses,attr,omitempty"`

	// generated from attribute "bracket" of type yes-no
	Bracket string `xml:"bracket,attr,omitempty"`

	// generated from attribute "size" of type symbol-size
	Size string `xml:"size,attr,omitempty"`
}

// AttributeGroup_line_height is generated from named attribute group "line-height"
type AttributeGroup_line_height struct {

	// insertion point for fields

	// generated from attribute "line-height" of type number-or-normal
	Line_height string `xml:"line-height,attr,omitempty"`
}

// AttributeGroup_line_length is generated from named attribute group "line-length"
type AttributeGroup_line_length struct {

	// insertion point for fields

	// generated from attribute "line-length" of type line-length
	Line_length string `xml:"line-length,attr,omitempty"`
}

// AttributeGroup_line_shape is generated from named attribute group "line-shape"
type AttributeGroup_line_shape struct {

	// insertion point for fields

	// generated from attribute "line-shape" of type line-shape
	Line_shape string `xml:"line-shape,attr,omitempty"`
}

// AttributeGroup_line_type is generated from named attribute group "line-type"
type AttributeGroup_line_type struct {

	// insertion point for fields

	// generated from attribute "line-type" of type line-type
	Line_type string `xml:"line-type,attr,omitempty"`
}

// AttributeGroup_link_attributes is generated from named attribute group "link-attributes"
type AttributeGroup_link_attributes struct {

	// insertion point for fields

	// generated from attribute "http://www.w3.org/1999/xlink href" of type 
	Href string `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`

	// generated from attribute "http://www.w3.org/1999/xlink type" of type 
	Type string `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`

	// generated from attribute "http://www.w3.org/1999/xlink role" of type 
	Role string `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`

	// generated from attribute "http://www.w3.org/1999/xlink title" of type 
	Title string `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`

	// generated from attribute "http://www.w3.org/1999/xlink show" of type 
	Show string `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`

	// generated from attribute "http://www.w3.org/1999/xlink actuate" of type 
	Actuate string `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

// AttributeGroup_measure_attributes is generated from named attribute group "measure-attributes"
type AttributeGroup_measure_attributes struct {

	// insertion point for fields

	// generated from attribute "number" of type xs:token
	Number string `xml:"number,attr,omitempty"`

	// generated from attribute "text" of type measure-text
	Text string `xml:"text,attr,omitempty"`

	// generated from attribute "implicit" of type yes-no
	Implicit string `xml:"implicit,attr,omitempty"`

	// generated from attribute "non-controlling" of type yes-no
	Non_controlling string `xml:"non-controlling,attr,omitempty"`

	// generated from attribute "width" of type tenths
	Width string `xml:"width,attr,omitempty"`

	// generated from attribute group "optional-unique-id
	AttributeGroup_optional_unique_id
}

// AttributeGroup_optional_unique_id is generated from named attribute group "optional-unique-id"
type AttributeGroup_optional_unique_id struct {

	// insertion point for fields

	// generated from attribute "id" of type xs:ID
	Id string `xml:"id,attr,omitempty"`
}

// AttributeGroup_orientation is generated from named attribute group "orientation"
type AttributeGroup_orientation struct {

	// insertion point for fields

	// generated from attribute "orientation" of type over-under
	Orientation string `xml:"orientation,attr,omitempty"`
}

// AttributeGroup_part_attributes is generated from named attribute group "part-attributes"
type AttributeGroup_part_attributes struct {

	// insertion point for fields

	// generated from attribute "id" of type xs:IDREF
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

	// generated from attribute "placement" of type above-below
	Placement string `xml:"placement,attr,omitempty"`
}

// AttributeGroup_position is generated from named attribute group "position"
type AttributeGroup_position struct {

	// insertion point for fields

	// generated from attribute "default-x" of type tenths
	Default_x string `xml:"default-x,attr,omitempty"`

	// generated from attribute "default-y" of type tenths
	Default_y string `xml:"default-y,attr,omitempty"`

	// generated from attribute "relative-x" of type tenths
	Relative_x string `xml:"relative-x,attr,omitempty"`

	// generated from attribute "relative-y" of type tenths
	Relative_y string `xml:"relative-y,attr,omitempty"`
}

// AttributeGroup_print_attributes is generated from named attribute group "print-attributes"
type AttributeGroup_print_attributes struct {

	// insertion point for fields

	// generated from attribute "staff-spacing" of type tenths
	Staff_spacing string `xml:"staff-spacing,attr,omitempty"`

	// generated from attribute "new-system" of type yes-no
	New_system string `xml:"new-system,attr,omitempty"`

	// generated from attribute "new-page" of type yes-no
	New_page string `xml:"new-page,attr,omitempty"`

	// generated from attribute "blank-page" of type xs:positiveInteger
	Blank_page int `xml:"blank-page,attr,omitempty"`

	// generated from attribute "page-number" of type xs:token
	Page_number string `xml:"page-number,attr,omitempty"`
}

// AttributeGroup_print_object is generated from named attribute group "print-object"
type AttributeGroup_print_object struct {

	// insertion point for fields

	// generated from attribute "print-object" of type yes-no
	Print_object string `xml:"print-object,attr,omitempty"`
}

// AttributeGroup_print_spacing is generated from named attribute group "print-spacing"
type AttributeGroup_print_spacing struct {

	// insertion point for fields

	// generated from attribute "print-spacing" of type yes-no
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

	// generated from attribute "print-dot" of type yes-no
	Print_dot string `xml:"print-dot,attr,omitempty"`

	// generated from attribute "print-lyric" of type yes-no
	Print_lyric string `xml:"print-lyric,attr,omitempty"`

	// generated from attribute group "print-object
	AttributeGroup_print_object

	// generated from attribute group "print-spacing
	AttributeGroup_print_spacing
}

// AttributeGroup_smufl is generated from named attribute group "smufl"
type AttributeGroup_smufl struct {

	// insertion point for fields

	// generated from attribute "smufl" of type smufl-glyph-name
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

	// generated from attribute "system" of type system-relation
	System string `xml:"system,attr,omitempty"`
}

// AttributeGroup_text_decoration is generated from named attribute group "text-decoration"
type AttributeGroup_text_decoration struct {

	// insertion point for fields

	// generated from attribute "underline" of type number-of-lines
	Underline int `xml:"underline,attr,omitempty"`

	// generated from attribute "overline" of type number-of-lines
	Overline int `xml:"overline,attr,omitempty"`

	// generated from attribute "line-through" of type number-of-lines
	Line_through int `xml:"line-through,attr,omitempty"`
}

// AttributeGroup_text_direction is generated from named attribute group "text-direction"
type AttributeGroup_text_direction struct {

	// insertion point for fields

	// generated from attribute "dir" of type text-direction
	Dir string `xml:"dir,attr,omitempty"`
}

// AttributeGroup_text_formatting is generated from named attribute group "text-formatting"
type AttributeGroup_text_formatting struct {

	// insertion point for fields

	// generated from attribute "http://www.w3.org/XML/1998/namespace lang" of type 
	Lang string `xml:"http://www.w3.org/XML/1998/namespace lang,attr,omitempty"`

	// generated from attribute "http://www.w3.org/XML/1998/namespace space" of type 
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

	// generated from attribute "rotation" of type rotation-degrees
	Rotation string `xml:"rotation,attr,omitempty"`
}

// AttributeGroup_trill_sound is generated from named attribute group "trill-sound"
type AttributeGroup_trill_sound struct {

	// insertion point for fields

	// generated from attribute "start-note" of type start-note
	Start_note string `xml:"start-note,attr,omitempty"`

	// generated from attribute "trill-step" of type trill-step
	Trill_step string `xml:"trill-step,attr,omitempty"`

	// generated from attribute "two-note-turn" of type two-note-turn
	Two_note_turn string `xml:"two-note-turn,attr,omitempty"`

	// generated from attribute "accelerate" of type yes-no
	Accelerate string `xml:"accelerate,attr,omitempty"`

	// generated from attribute "beats" of type trill-beats
	Beats string `xml:"beats,attr,omitempty"`

	// generated from attribute "second-beat" of type percent
	Second_beat string `xml:"second-beat,attr,omitempty"`

	// generated from attribute "last-beat" of type percent
	Last_beat string `xml:"last-beat,attr,omitempty"`
}

// AttributeGroup_valign is generated from named attribute group "valign"
type AttributeGroup_valign struct {

	// insertion point for fields

	// generated from attribute "valign" of type valign
	Valign string `xml:"valign,attr,omitempty"`
}

// AttributeGroup_valign_image is generated from named attribute group "valign-image"
type AttributeGroup_valign_image struct {

	// insertion point for fields

	// generated from attribute "valign" of type valign-image
	Valign string `xml:"valign,attr,omitempty"`
}

// AttributeGroup_x_position is generated from named attribute group "x-position"
type AttributeGroup_x_position struct {

	// insertion point for fields

	// generated from attribute "default-x" of type tenths
	Default_x string `xml:"default-x,attr,omitempty"`

	// generated from attribute "default-y" of type tenths
	Default_y string `xml:"default-y,attr,omitempty"`

	// generated from attribute "relative-x" of type tenths
	Relative_x string `xml:"relative-x,attr,omitempty"`

	// generated from attribute "relative-y" of type tenths
	Relative_y string `xml:"relative-y,attr,omitempty"`
}

// AttributeGroup_y_position is generated from named attribute group "y-position"
type AttributeGroup_y_position struct {

	// insertion point for fields

	// generated from attribute "default-x" of type tenths
	Default_x string `xml:"default-x,attr,omitempty"`

	// generated from attribute "default-y" of type tenths
	Default_y string `xml:"default-y,attr,omitempty"`

	// generated from attribute "relative-x" of type tenths
	Relative_x string `xml:"relative-x,attr,omitempty"`

	// generated from attribute "relative-y" of type tenths
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
