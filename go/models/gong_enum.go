// generated code - do not edit
package models

// insertion point of enum utility functions
// Utility function for ModelsFileTmplLevel0
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (modelsfiletmpllevel0 ModelsFileTmplLevel0) ToInt() (res int) {

	// migration of former implementation of enum
	switch modelsfiletmpllevel0 {
	// insertion code per enum code
	case ModelsFileTmplLevel0AllGongstructsCode:
		res = 0
	case ModelsFileTmplLevel0Nb:
		res = 1
	}
	return
}

func (modelsfiletmpllevel0 *ModelsFileTmplLevel0) FromInt(input int) (err error) {

	switch input {
	// insertion code per enum code
	case 0:
		*modelsfiletmpllevel0 = ModelsFileTmplLevel0AllGongstructsCode
		return
	case 1:
		*modelsfiletmpllevel0 = ModelsFileTmplLevel0Nb
		return
	default:
		return errUnkownEnum
	}
}

func (modelsfiletmpllevel0 *ModelsFileTmplLevel0) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "ModelsFileTmplLevel0AllGongstructsCode":
		*modelsfiletmpllevel0 = ModelsFileTmplLevel0AllGongstructsCode
	case "ModelsFileTmplLevel0Nb":
		*modelsfiletmpllevel0 = ModelsFileTmplLevel0Nb
	default:
		return errUnkownEnum
	}
	return
}

func (modelsfiletmpllevel0 *ModelsFileTmplLevel0) ToCodeString() (res string) {

	switch *modelsfiletmpllevel0 {
	// insertion code per enum code
	case ModelsFileTmplLevel0AllGongstructsCode:
		res = "ModelsFileTmplLevel0AllGongstructsCode"
	case ModelsFileTmplLevel0Nb:
		res = "ModelsFileTmplLevel0Nb"
	}
	return
}

func (modelsfiletmpllevel0 ModelsFileTmplLevel0) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "ModelsFileTmplLevel0AllGongstructsCode")
	res = append(res, "ModelsFileTmplLevel0Nb")

	return
}

func (modelsfiletmpllevel0 ModelsFileTmplLevel0) CodeValues() (res []int) {

	res = make([]int, 0)

	// insertion code per enum code
	res = append(res, 0)
	res = append(res, 1)

	return
}

// Utility function for ModelsFileTmplLevel1
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (modelsfiletmpllevel1 ModelsFileTmplLevel1) ToInt() (res int) {

	// migration of former implementation of enum
	switch modelsfiletmpllevel1 {
	// insertion code per enum code
	case ModelsFileTmplLevel1OneGongstructCode:
		res = 0
	case ModelsFileTmplLevel1Nb:
		res = 1
	}
	return
}

func (modelsfiletmpllevel1 *ModelsFileTmplLevel1) FromInt(input int) (err error) {

	switch input {
	// insertion code per enum code
	case 0:
		*modelsfiletmpllevel1 = ModelsFileTmplLevel1OneGongstructCode
		return
	case 1:
		*modelsfiletmpllevel1 = ModelsFileTmplLevel1Nb
		return
	default:
		return errUnkownEnum
	}
}

func (modelsfiletmpllevel1 *ModelsFileTmplLevel1) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "ModelsFileTmplLevel1OneGongstructCode":
		*modelsfiletmpllevel1 = ModelsFileTmplLevel1OneGongstructCode
	case "ModelsFileTmplLevel1Nb":
		*modelsfiletmpllevel1 = ModelsFileTmplLevel1Nb
	default:
		return errUnkownEnum
	}
	return
}

func (modelsfiletmpllevel1 *ModelsFileTmplLevel1) ToCodeString() (res string) {

	switch *modelsfiletmpllevel1 {
	// insertion code per enum code
	case ModelsFileTmplLevel1OneGongstructCode:
		res = "ModelsFileTmplLevel1OneGongstructCode"
	case ModelsFileTmplLevel1Nb:
		res = "ModelsFileTmplLevel1Nb"
	}
	return
}

func (modelsfiletmpllevel1 ModelsFileTmplLevel1) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "ModelsFileTmplLevel1OneGongstructCode")
	res = append(res, "ModelsFileTmplLevel1Nb")

	return
}

func (modelsfiletmpllevel1 ModelsFileTmplLevel1) CodeValues() (res []int) {

	res = make([]int, 0)

	// insertion code per enum code
	res = append(res, 0)
	res = append(res, 1)

	return
}

// Utility function for ModelsFileTmplLevel2
// if enum values are string, it is stored with the value
// if enum values are int, they are stored with the code of the value
func (modelsfiletmpllevel2 ModelsFileTmplLevel2) ToInt() (res int) {

	// migration of former implementation of enum
	switch modelsfiletmpllevel2 {
	// insertion code per enum code
	case ModelsFileTmplLevel2Structname:
		res = 0
	case ModelsFileTmplLevel2Nb:
		res = 1
	}
	return
}

func (modelsfiletmpllevel2 *ModelsFileTmplLevel2) FromInt(input int) (err error) {

	switch input {
	// insertion code per enum code
	case 0:
		*modelsfiletmpllevel2 = ModelsFileTmplLevel2Structname
		return
	case 1:
		*modelsfiletmpllevel2 = ModelsFileTmplLevel2Nb
		return
	default:
		return errUnkownEnum
	}
}

func (modelsfiletmpllevel2 *ModelsFileTmplLevel2) FromCodeString(input string) (err error) {

	switch input {
	// insertion code per enum code
	case "ModelsFileTmplLevel2Structname":
		*modelsfiletmpllevel2 = ModelsFileTmplLevel2Structname
	case "ModelsFileTmplLevel2Nb":
		*modelsfiletmpllevel2 = ModelsFileTmplLevel2Nb
	default:
		return errUnkownEnum
	}
	return
}

func (modelsfiletmpllevel2 *ModelsFileTmplLevel2) ToCodeString() (res string) {

	switch *modelsfiletmpllevel2 {
	// insertion code per enum code
	case ModelsFileTmplLevel2Structname:
		res = "ModelsFileTmplLevel2Structname"
	case ModelsFileTmplLevel2Nb:
		res = "ModelsFileTmplLevel2Nb"
	}
	return
}

func (modelsfiletmpllevel2 ModelsFileTmplLevel2) Codes() (res []string) {

	res = make([]string, 0)

	// insertion code per enum code
	res = append(res, "ModelsFileTmplLevel2Structname")
	res = append(res, "ModelsFileTmplLevel2Nb")

	return
}

func (modelsfiletmpllevel2 ModelsFileTmplLevel2) CodeValues() (res []int) {

	res = make([]int, 0)

	// insertion code per enum code
	res = append(res, 0)
	res = append(res, 1)

	return
}

// end of insertion point for enum utility functions

type GongstructEnumStringField interface {
	Codes() []string
	CodeValues() []string
	ToString() string
}

type PointerToGongstructEnumStringField interface {
	FromCodeString(input string) (err error)
}

type GongstructEnumIntField interface {
	int | ModelsFileTmplLevel0 | ModelsFileTmplLevel1 | ModelsFileTmplLevel2
	Codes() []string
	CodeValues() []int
}

type PointerToGongstructEnumIntField interface {
	*ModelsFileTmplLevel0 | *ModelsFileTmplLevel1 | *ModelsFileTmplLevel2
	FromCodeString(input string) (err error)
}

// Last line of the template
