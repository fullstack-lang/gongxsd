package models

const ModelsFileTemplate = `// generated code - do not edit
package models

import "encoding/xml"

// to avoid compilation error if no xml element is generated
var _ xml.Attr

{{` + string(rune(ModelsFileTmplLevel0AllGongstructsCode)) + `}}
`

type ModelsFileTmplLevel0 int

const (
	ModelsFileTmplLevel0AllGongstructsCode ModelsFileTmplLevel0 = iota
	ModelsFileTmplLevel0Nb
)

var ModelsFileTmplLevel0Code map[ModelsFileTmplLevel0]string = // new line
map[ModelsFileTmplLevel0]string{
	ModelsFileTmplLevel0AllGongstructsCode: ``,
}

type ModelsFileTmplLevel1 int

const (
	ModelsFileTmplLevel1OneGongstructCode ModelsFileTmplLevel1 = iota
	ModelsFileTmplLevel1Nb
)

var ModelsFileTmplLevel1Code map[ModelsFileTmplLevel1]string = // new line
map[ModelsFileTmplLevel1]string{
	ModelsFileTmplLevel1OneGongstructCode: `// {{` + string(rune(ModelsFileTmplLevel2Structname)) +
		`}} is generated from {{` + string(rune(ModelsFileTmplLevel2Source)) +
		`}}
type {{` + string(rune(ModelsFileTmplLevel2Structname)) + `}} struct {
	Name string
	
	// insertion point for fields{{` + string(rune(ModelsFileTmplLevel2Fields)) + `}}
}

`,
}

type ModelsFileTmplLevel2 int

const (
	ModelsFileTmplLevel2Structname ModelsFileTmplLevel2 = iota
	ModelsFileTmplLevel2Source
	ModelsFileTmplLevel2Fields
	ModelsFileTmplLevel2Nb
)

var ModelsFileTmplLevel2Code map[ModelsFileTmplLevel2]string = // new line
map[ModelsFileTmplLevel2]string{
	ModelsFileTmplLevel2Structname: ``,
	ModelsFileTmplLevel2Source:     ``,
	ModelsFileTmplLevel2Fields:     ``,
}
