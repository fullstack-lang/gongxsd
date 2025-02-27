// generated code - do not edit
package probe

import (
	"fmt"
	"log"
	"sort"

	gongtable "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/maticons/maticons"

	"github.com/fullstack-lang/gongxsd/test/reqif/go/models"
	"github.com/fullstack-lang/gongxsd/test/reqif/go/orm"
)

func fillUpTablePointerToGongstruct[T models.PointerToGongstruct](
	probe *Probe,
) {
	var typedInstance T
	switch any(typedInstance).(type) {
	// insertion point
	case *models.ALTERNATIVE_ID:
		fillUpTable[models.ALTERNATIVE_ID](probe)
	case *models.ATTRIBUTE_DEFINITION_BOOLEAN:
		fillUpTable[models.ATTRIBUTE_DEFINITION_BOOLEAN](probe)
	case *models.ATTRIBUTE_DEFINITION_DATE:
		fillUpTable[models.ATTRIBUTE_DEFINITION_DATE](probe)
	case *models.ATTRIBUTE_DEFINITION_ENUMERATION:
		fillUpTable[models.ATTRIBUTE_DEFINITION_ENUMERATION](probe)
	case *models.ATTRIBUTE_DEFINITION_INTEGER:
		fillUpTable[models.ATTRIBUTE_DEFINITION_INTEGER](probe)
	case *models.ATTRIBUTE_DEFINITION_REAL:
		fillUpTable[models.ATTRIBUTE_DEFINITION_REAL](probe)
	case *models.ATTRIBUTE_DEFINITION_STRING:
		fillUpTable[models.ATTRIBUTE_DEFINITION_STRING](probe)
	case *models.ATTRIBUTE_DEFINITION_XHTML:
		fillUpTable[models.ATTRIBUTE_DEFINITION_XHTML](probe)
	case *models.ATTRIBUTE_VALUE_BOOLEAN:
		fillUpTable[models.ATTRIBUTE_VALUE_BOOLEAN](probe)
	case *models.ATTRIBUTE_VALUE_DATE:
		fillUpTable[models.ATTRIBUTE_VALUE_DATE](probe)
	case *models.ATTRIBUTE_VALUE_ENUMERATION:
		fillUpTable[models.ATTRIBUTE_VALUE_ENUMERATION](probe)
	case *models.ATTRIBUTE_VALUE_INTEGER:
		fillUpTable[models.ATTRIBUTE_VALUE_INTEGER](probe)
	case *models.ATTRIBUTE_VALUE_REAL:
		fillUpTable[models.ATTRIBUTE_VALUE_REAL](probe)
	case *models.ATTRIBUTE_VALUE_STRING:
		fillUpTable[models.ATTRIBUTE_VALUE_STRING](probe)
	case *models.ATTRIBUTE_VALUE_XHTML:
		fillUpTable[models.ATTRIBUTE_VALUE_XHTML](probe)
	case *models.A_ALTERNATIVE_ID:
		fillUpTable[models.A_ALTERNATIVE_ID](probe)
	case *models.A_ATTRIBUTE_DEFINITION_BOOLEAN_REF:
		fillUpTable[models.A_ATTRIBUTE_DEFINITION_BOOLEAN_REF](probe)
	case *models.A_ATTRIBUTE_DEFINITION_DATE_REF:
		fillUpTable[models.A_ATTRIBUTE_DEFINITION_DATE_REF](probe)
	case *models.A_ATTRIBUTE_DEFINITION_ENUMERATION_REF:
		fillUpTable[models.A_ATTRIBUTE_DEFINITION_ENUMERATION_REF](probe)
	case *models.A_ATTRIBUTE_DEFINITION_INTEGER_REF:
		fillUpTable[models.A_ATTRIBUTE_DEFINITION_INTEGER_REF](probe)
	case *models.A_ATTRIBUTE_DEFINITION_REAL_REF:
		fillUpTable[models.A_ATTRIBUTE_DEFINITION_REAL_REF](probe)
	case *models.A_ATTRIBUTE_DEFINITION_STRING_REF:
		fillUpTable[models.A_ATTRIBUTE_DEFINITION_STRING_REF](probe)
	case *models.A_ATTRIBUTE_DEFINITION_XHTML_REF:
		fillUpTable[models.A_ATTRIBUTE_DEFINITION_XHTML_REF](probe)
	case *models.A_ATTRIBUTE_VALUE_BOOLEAN:
		fillUpTable[models.A_ATTRIBUTE_VALUE_BOOLEAN](probe)
	case *models.A_ATTRIBUTE_VALUE_DATE:
		fillUpTable[models.A_ATTRIBUTE_VALUE_DATE](probe)
	case *models.A_ATTRIBUTE_VALUE_ENUMERATION:
		fillUpTable[models.A_ATTRIBUTE_VALUE_ENUMERATION](probe)
	case *models.A_ATTRIBUTE_VALUE_INTEGER:
		fillUpTable[models.A_ATTRIBUTE_VALUE_INTEGER](probe)
	case *models.A_ATTRIBUTE_VALUE_REAL:
		fillUpTable[models.A_ATTRIBUTE_VALUE_REAL](probe)
	case *models.A_ATTRIBUTE_VALUE_STRING:
		fillUpTable[models.A_ATTRIBUTE_VALUE_STRING](probe)
	case *models.A_ATTRIBUTE_VALUE_XHTML:
		fillUpTable[models.A_ATTRIBUTE_VALUE_XHTML](probe)
	case *models.A_ATTRIBUTE_VALUE_XHTML_1:
		fillUpTable[models.A_ATTRIBUTE_VALUE_XHTML_1](probe)
	case *models.A_CHILDREN:
		fillUpTable[models.A_CHILDREN](probe)
	case *models.A_CORE_CONTENT:
		fillUpTable[models.A_CORE_CONTENT](probe)
	case *models.A_DATATYPES:
		fillUpTable[models.A_DATATYPES](probe)
	case *models.A_DATATYPE_DEFINITION_BOOLEAN_REF:
		fillUpTable[models.A_DATATYPE_DEFINITION_BOOLEAN_REF](probe)
	case *models.A_DATATYPE_DEFINITION_DATE_REF:
		fillUpTable[models.A_DATATYPE_DEFINITION_DATE_REF](probe)
	case *models.A_DATATYPE_DEFINITION_ENUMERATION_REF:
		fillUpTable[models.A_DATATYPE_DEFINITION_ENUMERATION_REF](probe)
	case *models.A_DATATYPE_DEFINITION_INTEGER_REF:
		fillUpTable[models.A_DATATYPE_DEFINITION_INTEGER_REF](probe)
	case *models.A_DATATYPE_DEFINITION_REAL_REF:
		fillUpTable[models.A_DATATYPE_DEFINITION_REAL_REF](probe)
	case *models.A_DATATYPE_DEFINITION_STRING_REF:
		fillUpTable[models.A_DATATYPE_DEFINITION_STRING_REF](probe)
	case *models.A_DATATYPE_DEFINITION_XHTML_REF:
		fillUpTable[models.A_DATATYPE_DEFINITION_XHTML_REF](probe)
	case *models.A_EDITABLE_ATTS:
		fillUpTable[models.A_EDITABLE_ATTS](probe)
	case *models.A_ENUM_VALUE_REF:
		fillUpTable[models.A_ENUM_VALUE_REF](probe)
	case *models.A_OBJECT:
		fillUpTable[models.A_OBJECT](probe)
	case *models.A_PROPERTIES:
		fillUpTable[models.A_PROPERTIES](probe)
	case *models.A_RELATION_GROUP_TYPE_REF:
		fillUpTable[models.A_RELATION_GROUP_TYPE_REF](probe)
	case *models.A_SOURCE_1:
		fillUpTable[models.A_SOURCE_1](probe)
	case *models.A_SOURCE_SPECIFICATION_1:
		fillUpTable[models.A_SOURCE_SPECIFICATION_1](probe)
	case *models.A_SPECIFICATIONS:
		fillUpTable[models.A_SPECIFICATIONS](probe)
	case *models.A_SPECIFICATION_TYPE_REF:
		fillUpTable[models.A_SPECIFICATION_TYPE_REF](probe)
	case *models.A_SPECIFIED_VALUES:
		fillUpTable[models.A_SPECIFIED_VALUES](probe)
	case *models.A_SPEC_ATTRIBUTES:
		fillUpTable[models.A_SPEC_ATTRIBUTES](probe)
	case *models.A_SPEC_OBJECTS:
		fillUpTable[models.A_SPEC_OBJECTS](probe)
	case *models.A_SPEC_OBJECT_TYPE_REF:
		fillUpTable[models.A_SPEC_OBJECT_TYPE_REF](probe)
	case *models.A_SPEC_RELATIONS:
		fillUpTable[models.A_SPEC_RELATIONS](probe)
	case *models.A_SPEC_RELATION_GROUPS:
		fillUpTable[models.A_SPEC_RELATION_GROUPS](probe)
	case *models.A_SPEC_RELATION_REF:
		fillUpTable[models.A_SPEC_RELATION_REF](probe)
	case *models.A_SPEC_RELATION_TYPE_REF:
		fillUpTable[models.A_SPEC_RELATION_TYPE_REF](probe)
	case *models.A_SPEC_TYPES:
		fillUpTable[models.A_SPEC_TYPES](probe)
	case *models.A_THE_HEADER:
		fillUpTable[models.A_THE_HEADER](probe)
	case *models.A_TOOL_EXTENSIONS:
		fillUpTable[models.A_TOOL_EXTENSIONS](probe)
	case *models.DATATYPE_DEFINITION_BOOLEAN:
		fillUpTable[models.DATATYPE_DEFINITION_BOOLEAN](probe)
	case *models.DATATYPE_DEFINITION_DATE:
		fillUpTable[models.DATATYPE_DEFINITION_DATE](probe)
	case *models.DATATYPE_DEFINITION_ENUMERATION:
		fillUpTable[models.DATATYPE_DEFINITION_ENUMERATION](probe)
	case *models.DATATYPE_DEFINITION_INTEGER:
		fillUpTable[models.DATATYPE_DEFINITION_INTEGER](probe)
	case *models.DATATYPE_DEFINITION_REAL:
		fillUpTable[models.DATATYPE_DEFINITION_REAL](probe)
	case *models.DATATYPE_DEFINITION_STRING:
		fillUpTable[models.DATATYPE_DEFINITION_STRING](probe)
	case *models.DATATYPE_DEFINITION_XHTML:
		fillUpTable[models.DATATYPE_DEFINITION_XHTML](probe)
	case *models.EMBEDDED_VALUE:
		fillUpTable[models.EMBEDDED_VALUE](probe)
	case *models.ENUM_VALUE:
		fillUpTable[models.ENUM_VALUE](probe)
	case *models.RELATION_GROUP:
		fillUpTable[models.RELATION_GROUP](probe)
	case *models.RELATION_GROUP_TYPE:
		fillUpTable[models.RELATION_GROUP_TYPE](probe)
	case *models.REQ_IF:
		fillUpTable[models.REQ_IF](probe)
	case *models.REQ_IF_CONTENT:
		fillUpTable[models.REQ_IF_CONTENT](probe)
	case *models.REQ_IF_HEADER:
		fillUpTable[models.REQ_IF_HEADER](probe)
	case *models.REQ_IF_TOOL_EXTENSION:
		fillUpTable[models.REQ_IF_TOOL_EXTENSION](probe)
	case *models.SPECIFICATION:
		fillUpTable[models.SPECIFICATION](probe)
	case *models.SPECIFICATION_TYPE:
		fillUpTable[models.SPECIFICATION_TYPE](probe)
	case *models.SPEC_HIERARCHY:
		fillUpTable[models.SPEC_HIERARCHY](probe)
	case *models.SPEC_OBJECT:
		fillUpTable[models.SPEC_OBJECT](probe)
	case *models.SPEC_OBJECT_TYPE:
		fillUpTable[models.SPEC_OBJECT_TYPE](probe)
	case *models.SPEC_RELATION:
		fillUpTable[models.SPEC_RELATION](probe)
	case *models.SPEC_RELATION_TYPE:
		fillUpTable[models.SPEC_RELATION_TYPE](probe)
	case *models.XHTML_CONTENT:
		fillUpTable[models.XHTML_CONTENT](probe)
	default:
		log.Println("unknow type")
	}
}

func fillUpTable[T models.Gongstruct](
	probe *Probe,
) {

	probe.tableStage.Reset()
	probe.tableStage.Commit()

	table := new(gongtable.Table).Stage(probe.tableStage)
	table.Name = "Table"
	table.HasColumnSorting = true
	table.HasFiltering = true
	table.HasPaginator = true
	table.HasCheckableRows = false
	table.HasSaveButton = false

	fields := models.GetFields[T]()
	reverseFields := models.GetReverseFields[T]()

	table.NbOfStickyColumns = 3

	// refresh the stage of interest
	probe.stageOfInterest.Checkout()

	setOfStructs := (*models.GetGongstructInstancesSet[T](probe.stageOfInterest))
	sliceOfGongStructsSorted := make([]*T, len(setOfStructs))
	i := 0
	for k := range setOfStructs {
		sliceOfGongStructsSorted[i] = k
		i++
	}
	sort.Slice(sliceOfGongStructsSorted, func(i, j int) bool {
		return orm.GetID(
			probe.stageOfInterest,
			probe.backRepoOfInterest,
			sliceOfGongStructsSorted[i],
		) <
			orm.GetID(
				probe.stageOfInterest,
				probe.backRepoOfInterest,
				sliceOfGongStructsSorted[j],
			)
	})

	column := new(gongtable.DisplayedColumn).Stage(probe.tableStage)
	column.Name = "ID"
	table.DisplayedColumns = append(table.DisplayedColumns, column)

	column = new(gongtable.DisplayedColumn).Stage(probe.tableStage)
	column.Name = "Delete"
	table.DisplayedColumns = append(table.DisplayedColumns, column)

	for _, fieldName := range fields {
		column := new(gongtable.DisplayedColumn).Stage(probe.tableStage)
		column.Name = fieldName
		table.DisplayedColumns = append(table.DisplayedColumns, column)
	}
	for _, reverseField := range reverseFields {
		column := new(gongtable.DisplayedColumn).Stage(probe.tableStage)
		column.Name = "(" + reverseField.GongstructName + ") -> " + reverseField.Fieldname
		table.DisplayedColumns = append(table.DisplayedColumns, column)
	}

	fieldIndex := 0
	for _, structInstance := range sliceOfGongStructsSorted {
		row := new(gongtable.Row).Stage(probe.tableStage)
		row.Name = models.GetFieldStringValue[T](*structInstance, "Name")

		updater := NewRowUpdate[T](structInstance, probe)
		updater.Instance = structInstance
		row.Impl = updater

		table.Rows = append(table.Rows, row)

		cell := (&gongtable.Cell{
			Name: "ID",
		}).Stage(probe.tableStage)
		row.Cells = append(row.Cells, cell)
		cellInt := (&gongtable.CellInt{
			Name: "ID",
			Value: orm.GetID(
				probe.stageOfInterest,
				probe.backRepoOfInterest,
				structInstance,
			),
		}).Stage(probe.tableStage)
		cell.CellInt = cellInt

		cell = (&gongtable.Cell{
			Name: "Delete Icon",
		}).Stage(probe.tableStage)
		row.Cells = append(row.Cells, cell)
		cellIcon := (&gongtable.CellIcon{
			Name: "Delete Icon",
			Icon: string(maticons.BUTTON_delete),
		}).Stage(probe.tableStage)
		cellIcon.Impl = NewCellDeleteIconImpl[T](structInstance, probe)
		cell.CellIcon = cellIcon

		for _, fieldName := range fields {
			value := models.GetFieldStringValue[T](*structInstance, fieldName)
			name := fmt.Sprintf("%d", fieldIndex) + " " + value
			fieldIndex++
			// log.Println(fieldName, value)
			cell := (&gongtable.Cell{
				Name: name,
			}).Stage(probe.tableStage)
			row.Cells = append(row.Cells, cell)

			cellString := (&gongtable.CellString{
				Name:  name,
				Value: value,
			}).Stage(probe.tableStage)
			cell.CellString = cellString
		}
		for _, reverseField := range reverseFields {

			value := orm.GetReverseFieldOwnerName[T](
				probe.stageOfInterest,
				probe.backRepoOfInterest,
				structInstance,
				&reverseField)
			name := fmt.Sprintf("%d", fieldIndex) + " " + value
			fieldIndex++
			// log.Println(fieldName, value)
			cell := (&gongtable.Cell{
				Name: name,
			}).Stage(probe.tableStage)
			row.Cells = append(row.Cells, cell)

			cellString := (&gongtable.CellString{
				Name:  name,
				Value: value,
			}).Stage(probe.tableStage)
			cell.CellString = cellString
		}
	}
}

func NewRowUpdate[T models.Gongstruct](
	Instance *T,
	probe *Probe,
) (rowUpdate *RowUpdate[T]) {
	rowUpdate = new(RowUpdate[T])
	rowUpdate.Instance = Instance
	rowUpdate.probe = probe
	return
}

type RowUpdate[T models.Gongstruct] struct {
	Instance   *T
	probe *Probe
}

func (rowUpdate *RowUpdate[T]) RowUpdated(stage *gongtable.StageStruct, row, updatedRow *gongtable.Row) {
	log.Println("RowUpdate: RowUpdated", updatedRow.Name)

	FillUpFormFromGongstruct(rowUpdate.Instance, rowUpdate.probe)
}
