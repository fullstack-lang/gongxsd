// generated code - do not edit
package models

import "encoding/xml"

// to avoid compilation error if no xml element is generated
var _ xml.Attr

// Accidental is generated from named complex type "accidental"
type Accidental struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Accidental_mark is generated from named complex type "accidental-mark"
type Accidental_mark struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Accidental_text is generated from named complex type "accidental-text"
type Accidental_text struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Accord is generated from named complex type "accord"
type Accord struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Accordion_registration is generated from named complex type "accordion-registration"
type Accordion_registration struct {
	Name string `xml:"-"`
	
	// insertion point for fields

	// generated from element "accordion-high" of type empty
	Accordion_high []*Empty `xml:"accordion-high"`

	// generated from element "accordion-middle" of type accordion-middle
	Accordion_middle int `xml:"accordion-middle"`

	// generated from element "accordion-low" of type empty
	Accordion_low []*Empty `xml:"accordion-low"`
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
}

// Arrow is generated from named complex type "arrow"
type Arrow struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Articulations is generated from named complex type "articulations"
type Articulations struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Assess is generated from named complex type "assess"
type Assess struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Attributes is generated from named complex type "attributes"
type Attributes struct {
	Name string `xml:"-"`
	
	// insertion point for fields

	// generated from element "divisions" of type positive-divisions
	Divisions float64 `xml:"divisions"`

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

	// generated from element "measure-style" of type measure-style
	Measure_style []*Measure_style `xml:"measure-style"`

	// generated from element "transpose" of type transpose
	Transpose []*Transpose `xml:"transpose"`

	// generated from element "for-part" of type for-part
	For_part []*For_part `xml:"for-part"`
}

// Backup is generated from named complex type "backup"
type Backup struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Bar_style_color is generated from named complex type "bar-style-color"
type Bar_style_color struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Barline is generated from named complex type "barline"
type Barline struct {
	Name string `xml:"-"`
	
	// insertion point for fields

	// generated from element "bar-style" of type bar-style-color
	Bar_style []*Bar_style_color `xml:"bar-style"`

	// generated from element "wavy-line" of type wavy-line
	Wavy_line []*Wavy_line `xml:"wavy-line"`

	// generated from element "segno" of type segno
	Segno []*Segno `xml:"segno"`

	// generated from element "coda" of type coda
	Coda []*Coda `xml:"coda"`

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
}

// Bass is generated from named complex type "bass"
type Bass struct {
	Name string `xml:"-"`
	
	// insertion point for fields

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
}

// Beam is generated from named complex type "beam"
type Beam struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Beat_repeat is generated from named complex type "beat-repeat"
type Beat_repeat struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Beat_unit_tied is generated from named complex type "beat-unit-tied"
type Beat_unit_tied struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Beater is generated from named complex type "beater"
type Beater struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Bend is generated from named complex type "bend"
type Bend struct {
	Name string `xml:"-"`
	
	// insertion point for fields

	// generated from element "bend-alter" of type semitones
	Bend_alter float64 `xml:"bend-alter"`

	// generated from element "with-bar" of type placement-text
	With_bar []*Placement_text `xml:"with-bar"`

	// generated from element "pre-bend" of type empty
	Pre_bend []*Empty `xml:"pre-bend"`

	// generated from element "release" of type release
	Release []*Release `xml:"release"`
}

// Bookmark is generated from named complex type "bookmark"
type Bookmark struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Bracket is generated from named complex type "bracket"
type Bracket struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Breath_mark is generated from named complex type "breath-mark"
type Breath_mark struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Caesura is generated from named complex type "caesura"
type Caesura struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Cancel is generated from named complex type "cancel"
type Cancel struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Clef is generated from named complex type "clef"
type Clef struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Coda is generated from named complex type "coda"
type Coda struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Credit is generated from named complex type "credit"
type Credit struct {
	Name string `xml:"-"`
	
	// insertion point for fields

	// generated from element "credit-type" of type xs:string
	Credit_type string `xml:"credit-type"`

	// generated from element "link" of type link
	Link []*Link `xml:"link"`

	// generated from element "bookmark" of type bookmark
	Bookmark []*Bookmark `xml:"bookmark"`

	// generated from element "credit-image" of type image
	Credit_image []*Image `xml:"credit-image"`

	// generated from element "link" of type link
	Link []*Link `xml:"link"`

	// generated from element "bookmark" of type bookmark
	Bookmark []*Bookmark `xml:"bookmark"`

	// generated from element "credit-words" of type formatted-text-id
	Credit_words []*Formatted_text_id `xml:"credit-words"`

	// generated from element "credit-symbol" of type formatted-symbol-id
	Credit_symbol []*Formatted_symbol_id `xml:"credit-symbol"`

	// generated from element "credit-words" of type formatted-text-id
	Credit_words []*Formatted_text_id `xml:"credit-words"`

	// generated from element "credit-symbol" of type formatted-symbol-id
	Credit_symbol []*Formatted_symbol_id `xml:"credit-symbol"`
}

// Dashes is generated from named complex type "dashes"
type Dashes struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Defaults is generated from named complex type "defaults"
type Defaults struct {
	Name string `xml:"-"`
	
	// insertion point for fields

	// generated from element "scaling" of type scaling
	Scaling []*Scaling `xml:"scaling"`

	// generated from element "concert-score" of type empty
	Concert_score []*Empty `xml:"concert-score"`

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
}

// Degree_type is generated from named complex type "degree-type"
type Degree_type struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Degree_value is generated from named complex type "degree-value"
type Degree_value struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Direction is generated from named complex type "direction"
type Direction struct {
	Name string `xml:"-"`
	
	// insertion point for fields

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
}

// Distance is generated from named complex type "distance"
type Distance struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Double is generated from named complex type "double"
type Double struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Dynamics is generated from named complex type "dynamics"
type Dynamics struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Effect is generated from named complex type "effect"
type Effect struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Elision is generated from named complex type "elision"
type Elision struct {
	Name string `xml:"-"`
	
	// insertion point for fields
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
}

// Empty_line is generated from named complex type "empty-line"
type Empty_line struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Empty_placement is generated from named complex type "empty-placement"
type Empty_placement struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Empty_placement_smufl is generated from named complex type "empty-placement-smufl"
type Empty_placement_smufl struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Empty_print_object_style_align is generated from named complex type "empty-print-object-style-align"
type Empty_print_object_style_align struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Empty_print_style is generated from named complex type "empty-print-style"
type Empty_print_style struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Empty_print_style_align is generated from named complex type "empty-print-style-align"
type Empty_print_style_align struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Empty_print_style_align_id is generated from named complex type "empty-print-style-align-id"
type Empty_print_style_align_id struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Empty_trill_sound is generated from named complex type "empty-trill-sound"
type Empty_trill_sound struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Encoding is generated from named complex type "encoding"
type Encoding struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Ending is generated from named complex type "ending"
type Ending struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Extend is generated from named complex type "extend"
type Extend struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Feature is generated from named complex type "feature"
type Feature struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Fermata is generated from named complex type "fermata"
type Fermata struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Figure is generated from named complex type "figure"
type Figure struct {
	Name string `xml:"-"`
	
	// insertion point for fields

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

	// generated from element "figure" of type figure
	Figure []*Figure `xml:"figure"`
}

// Fingering is generated from named complex type "fingering"
type Fingering struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// First_fret is generated from named complex type "first-fret"
type First_fret struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// For_part is generated from named complex type "for-part"
type For_part struct {
	Name string `xml:"-"`
	
	// insertion point for fields

	// generated from element "part-clef" of type part-clef
	Part_clef []*Part_clef `xml:"part-clef"`

	// generated from element "part-transpose" of type part-transpose
	Part_transpose []*Part_transpose `xml:"part-transpose"`
}

// Formatted_symbol is generated from named complex type "formatted-symbol"
type Formatted_symbol struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Formatted_symbol_id is generated from named complex type "formatted-symbol-id"
type Formatted_symbol_id struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Formatted_text is generated from named complex type "formatted-text"
type Formatted_text struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Formatted_text_id is generated from named complex type "formatted-text-id"
type Formatted_text_id struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Forward is generated from named complex type "forward"
type Forward struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Frame is generated from named complex type "frame"
type Frame struct {
	Name string `xml:"-"`
	
	// insertion point for fields

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

	// generated from element "string" of type string
	String []*String `xml:"string"`

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
}

// Glass is generated from named complex type "glass"
type Glass struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Glissando is generated from named complex type "glissando"
type Glissando struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Glyph is generated from named complex type "glyph"
type Glyph struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Grace is generated from named complex type "grace"
type Grace struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Group_barline is generated from named complex type "group-barline"
type Group_barline struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Group_name is generated from named complex type "group-name"
type Group_name struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Group_symbol is generated from named complex type "group-symbol"
type Group_symbol struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Grouping is generated from named complex type "grouping"
type Grouping struct {
	Name string `xml:"-"`
	
	// insertion point for fields

	// generated from element "feature" of type feature
	Feature []*Feature `xml:"feature"`
}

// Hammer_on_pull_off is generated from named complex type "hammer-on-pull-off"
type Hammer_on_pull_off struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Handbell is generated from named complex type "handbell"
type Handbell struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Harmon_closed is generated from named complex type "harmon-closed"
type Harmon_closed struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Harmon_mute is generated from named complex type "harmon-mute"
type Harmon_mute struct {
	Name string `xml:"-"`
	
	// insertion point for fields

	// generated from element "harmon-closed" of type harmon-closed
	Harmon_closed []*Harmon_closed `xml:"harmon-closed"`
}

// Harmonic is generated from named complex type "harmonic"
type Harmonic struct {
	Name string `xml:"-"`
	
	// insertion point for fields

	// generated from element "natural" of type empty
	Natural []*Empty `xml:"natural"`

	// generated from element "artificial" of type empty
	Artificial []*Empty `xml:"artificial"`

	// generated from element "base-pitch" of type empty
	Base_pitch []*Empty `xml:"base-pitch"`

	// generated from element "touching-pitch" of type empty
	Touching_pitch []*Empty `xml:"touching-pitch"`

	// generated from element "sounding-pitch" of type empty
	Sounding_pitch []*Empty `xml:"sounding-pitch"`
}

// Harmony is generated from named complex type "harmony"
type Harmony struct {
	Name string `xml:"-"`
	
	// insertion point for fields

	// generated from element "frame" of type frame
	Frame []*Frame `xml:"frame"`

	// generated from element "offset" of type offset
	Offset []*Offset `xml:"offset"`
}

// Harmony_alter is generated from named complex type "harmony-alter"
type Harmony_alter struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Harp_pedals is generated from named complex type "harp-pedals"
type Harp_pedals struct {
	Name string `xml:"-"`
	
	// insertion point for fields

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
}

// Horizontal_turn is generated from named complex type "horizontal-turn"
type Horizontal_turn struct {
	Name string `xml:"-"`
	
	// insertion point for fields
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
}

// Instrument is generated from named complex type "instrument"
type Instrument struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Instrument_change is generated from named complex type "instrument-change"
type Instrument_change struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Instrument_link is generated from named complex type "instrument-link"
type Instrument_link struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Interchangeable is generated from named complex type "interchangeable"
type Interchangeable struct {
	Name string `xml:"-"`
	
	// insertion point for fields

	// generated from element "time-relation" of type time-relation
	Time_relation string `xml:"time-relation"`
}

// Inversion is generated from named complex type "inversion"
type Inversion struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Key is generated from named complex type "key"
type Key struct {
	Name string `xml:"-"`
	
	// insertion point for fields

	// generated from element "key-octave" of type key-octave
	Key_octave []*Key_octave `xml:"key-octave"`
}

// Key_accidental is generated from named complex type "key-accidental"
type Key_accidental struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Key_octave is generated from named complex type "key-octave"
type Key_octave struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Kind is generated from named complex type "kind"
type Kind struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Level is generated from named complex type "level"
type Level struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Line_detail is generated from named complex type "line-detail"
type Line_detail struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Line_width is generated from named complex type "line-width"
type Line_width struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Link is generated from named complex type "link"
type Link struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Listen is generated from named complex type "listen"
type Listen struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Listening is generated from named complex type "listening"
type Listening struct {
	Name string `xml:"-"`
	
	// insertion point for fields

	// generated from element "offset" of type offset
	Offset []*Offset `xml:"offset"`

	// generated from element "sync" of type sync
	Sync []*Sync `xml:"sync"`

	// generated from element "other-listening" of type other-listening
	Other_listening []*Other_listening `xml:"other-listening"`
}

// Lyric is generated from named complex type "lyric"
type Lyric struct {
	Name string `xml:"-"`
	
	// insertion point for fields

	// generated from element "end-line" of type empty
	End_line []*Empty `xml:"end-line"`

	// generated from element "end-paragraph" of type empty
	End_paragraph []*Empty `xml:"end-paragraph"`

	// generated from element "extend" of type extend
	Extend []*Extend `xml:"extend"`

	// generated from element "laughing" of type empty
	Laughing []*Empty `xml:"laughing"`

	// generated from element "humming" of type empty
	Humming []*Empty `xml:"humming"`

	// generated from element "syllabic" of type syllabic
	Syllabic string `xml:"syllabic"`

	// generated from element "text" of type text-element-data
	Text []*Text_element_data `xml:"text"`

	// generated from element "extend" of type extend
	Extend []*Extend `xml:"extend"`

	// generated from element "text" of type text-element-data
	Text []*Text_element_data `xml:"text"`

	// generated from element "elision" of type elision
	Elision []*Elision `xml:"elision"`

	// generated from element "syllabic" of type syllabic
	Syllabic string `xml:"syllabic"`
}

// Lyric_font is generated from named complex type "lyric-font"
type Lyric_font struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Lyric_language is generated from named complex type "lyric-language"
type Lyric_language struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Measure_layout is generated from named complex type "measure-layout"
type Measure_layout struct {
	Name string `xml:"-"`
	
	// insertion point for fields

	// generated from element "measure-distance" of type tenths
	Measure_distance float64 `xml:"measure-distance"`
}

// Measure_numbering is generated from named complex type "measure-numbering"
type Measure_numbering struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Measure_repeat is generated from named complex type "measure-repeat"
type Measure_repeat struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Measure_style is generated from named complex type "measure-style"
type Measure_style struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Membrane is generated from named complex type "membrane"
type Membrane struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Metal is generated from named complex type "metal"
type Metal struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Metronome is generated from named complex type "metronome"
type Metronome struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Metronome_beam is generated from named complex type "metronome-beam"
type Metronome_beam struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Metronome_note is generated from named complex type "metronome-note"
type Metronome_note struct {
	Name string `xml:"-"`
	
	// insertion point for fields

	// generated from element "metronome-type" of type note-type-value
	Metronome_type string `xml:"metronome-type"`

	// generated from element "metronome-dot" of type empty
	Metronome_dot []*Empty `xml:"metronome-dot"`

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
}

// Midi_instrument is generated from named complex type "midi-instrument"
type Midi_instrument struct {
	Name string `xml:"-"`
	
	// insertion point for fields

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
	Volume float64 `xml:"volume"`

	// generated from element "pan" of type rotation-degrees
	Pan float64 `xml:"pan"`

	// generated from element "elevation" of type rotation-degrees
	Elevation float64 `xml:"elevation"`
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
}

// Name_display is generated from named complex type "name-display"
type Name_display struct {
	Name string `xml:"-"`
	
	// insertion point for fields

	// generated from element "display-text" of type formatted-text
	Display_text []*Formatted_text `xml:"display-text"`

	// generated from element "accidental-text" of type accidental-text
	Accidental_text []*Accidental_text `xml:"accidental-text"`
}

// Non_arpeggiate is generated from named complex type "non-arpeggiate"
type Non_arpeggiate struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Notations is generated from named complex type "notations"
type Notations struct {
	Name string `xml:"-"`
	
	// insertion point for fields

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

	// generated from element "grace" of type grace
	Grace []*Grace `xml:"grace"`

	// generated from element "cue" of type empty
	Cue []*Empty `xml:"cue"`

	// generated from element "tie" of type tie
	Tie []*Tie `xml:"tie"`

	// generated from element "tie" of type tie
	Tie []*Tie `xml:"tie"`

	// generated from element "cue" of type empty
	Cue []*Empty `xml:"cue"`
}

// Note_size is generated from named complex type "note-size"
type Note_size struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Note_type is generated from named complex type "note-type"
type Note_type struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Notehead is generated from named complex type "notehead"
type Notehead struct {
	Name string `xml:"-"`
	
	// insertion point for fields
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

	// generated from element "numeral-fifths" of type fifths
	Numeral_fifths int `xml:"numeral-fifths"`

	// generated from element "numeral-mode" of type numeral-mode
	Numeral_mode string `xml:"numeral-mode"`
}

// Numeral_root is generated from named complex type "numeral-root"
type Numeral_root struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Octave_shift is generated from named complex type "octave-shift"
type Octave_shift struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Offset is generated from named complex type "offset"
type Offset struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Opus is generated from named complex type "opus"
type Opus struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Ornaments is generated from named complex type "ornaments"
type Ornaments struct {
	Name string `xml:"-"`
	
	// insertion point for fields

	// generated from element "accidental-mark" of type accidental-mark
	Accidental_mark []*Accidental_mark `xml:"accidental-mark"`

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
}

// Other_appearance is generated from named complex type "other-appearance"
type Other_appearance struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Other_direction is generated from named complex type "other-direction"
type Other_direction struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Other_listening is generated from named complex type "other-listening"
type Other_listening struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Other_notation is generated from named complex type "other-notation"
type Other_notation struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Other_placement_text is generated from named complex type "other-placement-text"
type Other_placement_text struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Other_play is generated from named complex type "other-play"
type Other_play struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Other_text is generated from named complex type "other-text"
type Other_text struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Page_layout is generated from named complex type "page-layout"
type Page_layout struct {
	Name string `xml:"-"`
	
	// insertion point for fields

	// generated from element "page-margins" of type page-margins
	Page_margins []*Page_margins `xml:"page-margins"`

	// generated from element "page-height" of type tenths
	Page_height float64 `xml:"page-height"`

	// generated from element "page-width" of type tenths
	Page_width float64 `xml:"page-width"`
}

// Page_margins is generated from named complex type "page-margins"
type Page_margins struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Part_clef is generated from named complex type "part-clef"
type Part_clef struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Part_group is generated from named complex type "part-group"
type Part_group struct {
	Name string `xml:"-"`
	
	// insertion point for fields

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
	Group_time []*Empty `xml:"group-time"`
}

// Part_link is generated from named complex type "part-link"
type Part_link struct {
	Name string `xml:"-"`
	
	// insertion point for fields

	// generated from element "instrument-link" of type instrument-link
	Instrument_link []*Instrument_link `xml:"instrument-link"`

	// generated from element "group-link" of type xs:string
	Group_link string `xml:"group-link"`
}

// Part_list is generated from named complex type "part-list"
type Part_list struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Part_name is generated from named complex type "part-name"
type Part_name struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Part_symbol is generated from named complex type "part-symbol"
type Part_symbol struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Part_transpose is generated from named complex type "part-transpose"
type Part_transpose struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Pedal is generated from named complex type "pedal"
type Pedal struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Pedal_tuning is generated from named complex type "pedal-tuning"
type Pedal_tuning struct {
	Name string `xml:"-"`
	
	// insertion point for fields

	// generated from element "pedal-step" of type step
	Pedal_step string `xml:"pedal-step"`

	// generated from element "pedal-alter" of type semitones
	Pedal_alter float64 `xml:"pedal-alter"`
}

// Per_minute is generated from named complex type "per-minute"
type Per_minute struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Percussion is generated from named complex type "percussion"
type Percussion struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Pitch is generated from named complex type "pitch"
type Pitch struct {
	Name string `xml:"-"`
	
	// insertion point for fields

	// generated from element "step" of type step
	Step string `xml:"step"`

	// generated from element "alter" of type semitones
	Alter float64 `xml:"alter"`

	// generated from element "octave" of type octave
	Octave int `xml:"octave"`
}

// Pitched is generated from named complex type "pitched"
type Pitched struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Placement_text is generated from named complex type "placement-text"
type Placement_text struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Play is generated from named complex type "play"
type Play struct {
	Name string `xml:"-"`
	
	// insertion point for fields

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

	// generated from element "player-name" of type xs:string
	Player_name string `xml:"player-name"`
}

// Principal_voice is generated from named complex type "principal-voice"
type Principal_voice struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Print is generated from named complex type "print"
type Print struct {
	Name string `xml:"-"`
	
	// insertion point for fields

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
}

// Rest is generated from named complex type "rest"
type Rest struct {
	Name string `xml:"-"`
	
	// insertion point for fields
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
}

// Scaling is generated from named complex type "scaling"
type Scaling struct {
	Name string `xml:"-"`
	
	// insertion point for fields

	// generated from element "millimeters" of type millimeters
	Millimeters float64 `xml:"millimeters"`

	// generated from element "tenths" of type tenths
	Tenths float64 `xml:"tenths"`
}

// Scordatura is generated from named complex type "scordatura"
type Scordatura struct {
	Name string `xml:"-"`
	
	// insertion point for fields

	// generated from element "accord" of type accord
	Accord []*Accord `xml:"accord"`
}

// Score_instrument is generated from named complex type "score-instrument"
type Score_instrument struct {
	Name string `xml:"-"`
	
	// insertion point for fields

	// generated from element "instrument-name" of type xs:string
	Instrument_name string `xml:"instrument-name"`

	// generated from element "instrument-abbreviation" of type xs:string
	Instrument_abbreviation string `xml:"instrument-abbreviation"`
}

// Score_part is generated from named complex type "score-part"
type Score_part struct {
	Name string `xml:"-"`
	
	// insertion point for fields

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

	// generated from element "midi-device" of type midi-device
	Midi_device []*Midi_device `xml:"midi-device"`

	// generated from element "midi-instrument" of type midi-instrument
	Midi_instrument []*Midi_instrument `xml:"midi-instrument"`
}

// Segno is generated from named complex type "segno"
type Segno struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Slash is generated from named complex type "slash"
type Slash struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Slide is generated from named complex type "slide"
type Slide struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Slur is generated from named complex type "slur"
type Slur struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Sound is generated from named complex type "sound"
type Sound struct {
	Name string `xml:"-"`
	
	// insertion point for fields

	// generated from element "swing" of type swing
	Swing []*Swing `xml:"swing"`

	// generated from element "offset" of type offset
	Offset []*Offset `xml:"offset"`

	// generated from element "instrument-change" of type instrument-change
	Instrument_change []*Instrument_change `xml:"instrument-change"`

	// generated from element "midi-device" of type midi-device
	Midi_device []*Midi_device `xml:"midi-device"`

	// generated from element "midi-instrument" of type midi-instrument
	Midi_instrument []*Midi_instrument `xml:"midi-instrument"`

	// generated from element "play" of type play
	Play []*Play `xml:"play"`
}

// Staff_details is generated from named complex type "staff-details"
type Staff_details struct {
	Name string `xml:"-"`
	
	// insertion point for fields

	// generated from element "staff-type" of type staff-type
	Staff_type string `xml:"staff-type"`

	// generated from element "staff-tuning" of type staff-tuning
	Staff_tuning []*Staff_tuning `xml:"staff-tuning"`

	// generated from element "capo" of type xs:nonNegativeInteger
	Capo int `xml:"capo"`

	// generated from element "staff-size" of type staff-size
	Staff_size []*Staff_size `xml:"staff-size"`

	// generated from element "staff-lines" of type xs:nonNegativeInteger
	Staff_lines int `xml:"staff-lines"`

	// generated from element "line-detail" of type line-detail
	Line_detail []*Line_detail `xml:"line-detail"`
}

// Staff_divide is generated from named complex type "staff-divide"
type Staff_divide struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Staff_layout is generated from named complex type "staff-layout"
type Staff_layout struct {
	Name string `xml:"-"`
	
	// insertion point for fields

	// generated from element "staff-distance" of type tenths
	Staff_distance float64 `xml:"staff-distance"`
}

// Staff_size is generated from named complex type "staff-size"
type Staff_size struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Staff_tuning is generated from named complex type "staff-tuning"
type Staff_tuning struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Stem is generated from named complex type "stem"
type Stem struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Stick is generated from named complex type "stick"
type Stick struct {
	Name string `xml:"-"`
	
	// insertion point for fields

	// generated from element "stick-type" of type stick-type
	Stick_type string `xml:"stick-type"`

	// generated from element "stick-material" of type stick-material
	Stick_material string `xml:"stick-material"`
}

// String is generated from named complex type "string"
type String struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// String_mute is generated from named complex type "string-mute"
type String_mute struct {
	Name string `xml:"-"`
	
	// insertion point for fields
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
}

// Supports is generated from named complex type "supports"
type Supports struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Swing is generated from named complex type "swing"
type Swing struct {
	Name string `xml:"-"`
	
	// insertion point for fields

	// generated from element "swing-style" of type xs:string
	Swing_style string `xml:"swing-style"`

	// generated from element "straight" of type empty
	Straight []*Empty `xml:"straight"`

	// generated from element "first" of type xs:positiveInteger
	First int `xml:"first"`

	// generated from element "second" of type xs:positiveInteger
	Second int `xml:"second"`

	// generated from element "swing-type" of type swing-type-value
	Swing_type string `xml:"swing-type"`
}

// Sync is generated from named complex type "sync"
type Sync struct {
	Name string `xml:"-"`
	
	// insertion point for fields
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
	System_distance float64 `xml:"system-distance"`

	// generated from element "top-system-distance" of type tenths
	Top_system_distance float64 `xml:"top-system-distance"`

	// generated from element "system-dividers" of type system-dividers
	System_dividers []*System_dividers `xml:"system-dividers"`
}

// System_margins is generated from named complex type "system-margins"
type System_margins struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Tap is generated from named complex type "tap"
type Tap struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Technical is generated from named complex type "technical"
type Technical struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Text_element_data is generated from named complex type "text-element-data"
type Text_element_data struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Tie is generated from named complex type "tie"
type Tie struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Tied is generated from named complex type "tied"
type Tied struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Time is generated from named complex type "time"
type Time struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Time_modification is generated from named complex type "time-modification"
type Time_modification struct {
	Name string `xml:"-"`
	
	// insertion point for fields

	// generated from element "actual-notes" of type xs:nonNegativeInteger
	Actual_notes int `xml:"actual-notes"`

	// generated from element "normal-notes" of type xs:nonNegativeInteger
	Normal_notes int `xml:"normal-notes"`

	// generated from element "normal-type" of type note-type-value
	Normal_type string `xml:"normal-type"`

	// generated from element "normal-dot" of type empty
	Normal_dot []*Empty `xml:"normal-dot"`
}

// Timpani is generated from named complex type "timpani"
type Timpani struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Transpose is generated from named complex type "transpose"
type Transpose struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Tremolo is generated from named complex type "tremolo"
type Tremolo struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Tuplet is generated from named complex type "tuplet"
type Tuplet struct {
	Name string `xml:"-"`
	
	// insertion point for fields

	// generated from element "tuplet-actual" of type tuplet-portion
	Tuplet_actual []*Tuplet_portion `xml:"tuplet-actual"`

	// generated from element "tuplet-normal" of type tuplet-portion
	Tuplet_normal []*Tuplet_portion `xml:"tuplet-normal"`
}

// Tuplet_dot is generated from named complex type "tuplet-dot"
type Tuplet_dot struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Tuplet_number is generated from named complex type "tuplet-number"
type Tuplet_number struct {
	Name string `xml:"-"`
	
	// insertion point for fields
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
}

// Typed_text is generated from named complex type "typed-text"
type Typed_text struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Unpitched is generated from named complex type "unpitched"
type Unpitched struct {
	Name string `xml:"-"`
	
	// insertion point for fields
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
}

// Wavy_line is generated from named complex type "wavy-line"
type Wavy_line struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Wedge is generated from named complex type "wedge"
type Wedge struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Wood is generated from named complex type "wood"
type Wood struct {
	Name string `xml:"-"`
	
	// insertion point for fields
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

// Directive is generated from inlined complex type within element "directive"
type Directive struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Measure_1 is generated from inlined complex type within element "measure"
// Identifier is post fixed because more than one xsd element has the name "measure"
type Measure_1 struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Measure is generated from inlined complex type within element "measure"
type Measure struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Part_1 is generated from inlined complex type within element "part"
// Identifier is post fixed because more than one xsd element has the name "part"
type Part_1 struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Part is generated from inlined complex type within element "part"
type Part struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Score_partwise is generated from inlined complex type within element "score-partwise"
type Score_partwise struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}

// Score_timewise is generated from inlined complex type within element "score-timewise"
type Score_timewise struct {
	Name string `xml:"-"`
	
	// insertion point for fields
}


