// generated code - do not edit
package models

import (
	"fmt"
	"log"
	"unicode/utf8"

	"github.com/xuri/excelize/v2"
)

func SerializeStage(stage *StageStruct, filename string) {

	f := excelize.NewFile()
	{
		// insertion point
		SerializeExcelize[ALTERNATIVE_ID](stage, f)
		SerializeExcelize[ATTRIBUTE_DEFINITION_BOOLEAN](stage, f)
		SerializeExcelize[ATTRIBUTE_DEFINITION_DATE](stage, f)
		SerializeExcelize[ATTRIBUTE_DEFINITION_ENUMERATION](stage, f)
		SerializeExcelize[ATTRIBUTE_DEFINITION_INTEGER](stage, f)
		SerializeExcelize[ATTRIBUTE_DEFINITION_REAL](stage, f)
		SerializeExcelize[ATTRIBUTE_DEFINITION_STRING](stage, f)
		SerializeExcelize[ATTRIBUTE_DEFINITION_XHTML](stage, f)
		SerializeExcelize[ATTRIBUTE_VALUE_BOOLEAN](stage, f)
		SerializeExcelize[ATTRIBUTE_VALUE_DATE](stage, f)
		SerializeExcelize[ATTRIBUTE_VALUE_ENUMERATION](stage, f)
		SerializeExcelize[ATTRIBUTE_VALUE_INTEGER](stage, f)
		SerializeExcelize[ATTRIBUTE_VALUE_REAL](stage, f)
		SerializeExcelize[ATTRIBUTE_VALUE_STRING](stage, f)
		SerializeExcelize[ATTRIBUTE_VALUE_XHTML](stage, f)
		SerializeExcelize[A_ALTERNATIVE_ID](stage, f)
		SerializeExcelize[A_CHILDREN](stage, f)
		SerializeExcelize[A_CORE_CONTENT](stage, f)
		SerializeExcelize[A_DATATYPES](stage, f)
		SerializeExcelize[A_DEFAULT_VALUE](stage, f)
		SerializeExcelize[A_DEFAULT_VALUE_1](stage, f)
		SerializeExcelize[A_DEFAULT_VALUE_2](stage, f)
		SerializeExcelize[A_DEFAULT_VALUE_3](stage, f)
		SerializeExcelize[A_DEFAULT_VALUE_4](stage, f)
		SerializeExcelize[A_DEFAULT_VALUE_5](stage, f)
		SerializeExcelize[A_DEFAULT_VALUE_6](stage, f)
		SerializeExcelize[A_DEFINITION](stage, f)
		SerializeExcelize[A_DEFINITION_1](stage, f)
		SerializeExcelize[A_DEFINITION_2](stage, f)
		SerializeExcelize[A_DEFINITION_3](stage, f)
		SerializeExcelize[A_DEFINITION_4](stage, f)
		SerializeExcelize[A_DEFINITION_5](stage, f)
		SerializeExcelize[A_DEFINITION_6](stage, f)
		SerializeExcelize[A_EDITABLE_ATTS](stage, f)
		SerializeExcelize[A_OBJECT](stage, f)
		SerializeExcelize[A_PROPERTIES](stage, f)
		SerializeExcelize[A_SOURCE](stage, f)
		SerializeExcelize[A_SOURCE_SPECIFICATION](stage, f)
		SerializeExcelize[A_SPECIFICATIONS](stage, f)
		SerializeExcelize[A_SPECIFIED_VALUES](stage, f)
		SerializeExcelize[A_SPEC_ATTRIBUTES](stage, f)
		SerializeExcelize[A_SPEC_OBJECTS](stage, f)
		SerializeExcelize[A_SPEC_RELATIONS](stage, f)
		SerializeExcelize[A_SPEC_RELATIONS_1](stage, f)
		SerializeExcelize[A_SPEC_RELATION_GROUPS](stage, f)
		SerializeExcelize[A_SPEC_TYPES](stage, f)
		SerializeExcelize[A_THE_HEADER](stage, f)
		SerializeExcelize[A_TOOL_EXTENSIONS](stage, f)
		SerializeExcelize[A_TYPE](stage, f)
		SerializeExcelize[A_TYPE_1](stage, f)
		SerializeExcelize[A_TYPE_10](stage, f)
		SerializeExcelize[A_TYPE_2](stage, f)
		SerializeExcelize[A_TYPE_3](stage, f)
		SerializeExcelize[A_TYPE_4](stage, f)
		SerializeExcelize[A_TYPE_5](stage, f)
		SerializeExcelize[A_TYPE_6](stage, f)
		SerializeExcelize[A_TYPE_7](stage, f)
		SerializeExcelize[A_TYPE_8](stage, f)
		SerializeExcelize[A_TYPE_9](stage, f)
		SerializeExcelize[A_VALUES](stage, f)
		SerializeExcelize[A_VALUES_1](stage, f)
		SerializeExcelize[DATATYPE_DEFINITION_BOOLEAN](stage, f)
		SerializeExcelize[DATATYPE_DEFINITION_DATE](stage, f)
		SerializeExcelize[DATATYPE_DEFINITION_ENUMERATION](stage, f)
		SerializeExcelize[DATATYPE_DEFINITION_INTEGER](stage, f)
		SerializeExcelize[DATATYPE_DEFINITION_REAL](stage, f)
		SerializeExcelize[DATATYPE_DEFINITION_STRING](stage, f)
		SerializeExcelize[DATATYPE_DEFINITION_XHTML](stage, f)
		SerializeExcelize[EMBEDDED_VALUE](stage, f)
		SerializeExcelize[ENUM_VALUE](stage, f)
		SerializeExcelize[RELATION_GROUP](stage, f)
		SerializeExcelize[RELATION_GROUP_TYPE](stage, f)
		SerializeExcelize[REQ_IF](stage, f)
		SerializeExcelize[REQ_IF_CONTENT](stage, f)
		SerializeExcelize[REQ_IF_HEADER](stage, f)
		SerializeExcelize[REQ_IF_TOOL_EXTENSION](stage, f)
		SerializeExcelize[SPECIFICATION](stage, f)
		SerializeExcelize[SPECIFICATION_TYPE](stage, f)
		SerializeExcelize[SPEC_HIERARCHY](stage, f)
		SerializeExcelize[SPEC_OBJECT](stage, f)
		SerializeExcelize[SPEC_OBJECT_TYPE](stage, f)
		SerializeExcelize[SPEC_RELATION](stage, f)
		SerializeExcelize[SPEC_RELATION_TYPE](stage, f)
		SerializeExcelize[XHTML_CONTENT](stage, f)
	}

	var tab ExcelizeTabulator
	tab.SetExcelizeFile(f)
	{
		f.DeleteSheet("Sheet1")
		if err := f.SaveAs(filename); err != nil {
			fmt.Println("cannot write xl file : ", err)
		}
	}

}

// Tabulator is an interface for writing to a table strings
type Tabulator interface {
	AddSheet(sheetName string)
	AddRow(sheetName string) int
	AddCell(sheetName string, rowId, columnIndex int, value string)
}

func Serialize[Type Gongstruct](stage *StageStruct, tab Tabulator) {
	sheetName := GetGongstructName[Type]()

	// Create a new sheet.
	tab.AddSheet(sheetName)

	headerRowIndex := tab.AddRow(sheetName)
	for colIndex, fieldName := range GetFields[Type]() {
		tab.AddCell(sheetName, headerRowIndex, colIndex, fieldName)
		// f.SetCellStr(sheetName, fmt.Sprintf("%s%d", IntToLetters(int32(index+1)), line), fieldName)
	}

	set := *GetGongstructInstancesSet[Type](stage)
	for instance := range set {
		line := tab.AddRow(sheetName)
		for index, fieldName := range GetFields[Type]() {
			tab.AddCell(sheetName, line, index, GetFieldStringValue(
				any(*instance).(Type), fieldName))
			// f.SetCellStr(sheetName, fmt.Sprintf("%s%d", IntToLetters(int32(index+1)), line), GetFieldStringValue(
			// 	any(*instance).(Type), fieldName))
		}
	}
}

type ExcelizeTabulator struct {
	f *excelize.File
}

func (tab *ExcelizeTabulator) SetExcelizeFile(f *excelize.File) {
	tab.f = f
}

func (tab *ExcelizeTabulator) AddSheet(sheetName string) {

}

func (tab *ExcelizeTabulator) AddRow(sheetName string) (rowId int) {
	return
}

func (tab *ExcelizeTabulator) AddCell(sheetName string, rowId, columnIndex int, value string) {

}

func SerializeExcelize[Type Gongstruct](stage *StageStruct, f *excelize.File) {
	sheetName := GetGongstructName[Type]()

	// Create a new sheet.
	f.NewSheet(sheetName)

	set := *GetGongstructInstancesSet[Type](stage)
	line := 1

	for index, fieldName := range GetFields[Type]() {
		f.SetCellStr(sheetName, fmt.Sprintf("%s%d", IntToLetters(int32(index+1)), line), fieldName)
	}
	f.AutoFilter(sheetName,
		fmt.Sprintf("%s%d", IntToLetters(int32(1)), line),
		[]excelize.AutoFilterOptions{})

	for instance := range set {
		line = line + 1
		for index, fieldName := range GetFields[Type]() {
			f.SetCellStr(sheetName, fmt.Sprintf("%s%d", IntToLetters(int32(index+1)), line), GetFieldStringValue(
				any(*instance).(Type), fieldName))
		}
	}

	// Autofit all columns according to their text content
	cols, err := f.GetCols(sheetName)
	if err != nil {
		log.Panicln("SerializeExcelize")
	}
	for idx, col := range cols {
		largestWidth := 0
		for _, rowCell := range col {
			cellWidth := utf8.RuneCountInString(rowCell) + 2 // + 2 for margin
			if cellWidth > largestWidth {
				largestWidth = cellWidth
			}
		}
		name, err := excelize.ColumnNumberToName(idx + 1)
		if err != nil {
			log.Panicln("SerializeExcelize")
		}
		f.SetColWidth(sheetName, name, name, float64(largestWidth))
	}
}

func IntToLetters(number int32) (letters string) {
	number--
	if firstLetter := number / 26; firstLetter > 0 {
		letters += IntToLetters(firstLetter)
		letters += string('A' + number%26)
	} else {
		letters += string('A' + number)
	}

	return
}
