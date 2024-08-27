// generated code - do not edit
package probe

import (
	"log"

	gongtable "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongxsd/test/reqif/go/models"
)

func NewCellDeleteIconImpl[T models.Gongstruct](
	Instance *T,
	probe *Probe,
) (cellDeleteIconImpl *CellDeleteIconImpl[T]) {
	cellDeleteIconImpl = new(CellDeleteIconImpl[T])
	cellDeleteIconImpl.Instance = Instance
	cellDeleteIconImpl.probe = probe
	return
}

type CellDeleteIconImpl[T models.Gongstruct] struct {
	Instance   *T
	probe *Probe
}

func (cellDeleteIconImpl *CellDeleteIconImpl[T]) CellIconUpdated(stage *gongtable.StageStruct,
	row, updatedCellIcon *gongtable.CellIcon) {
	log.Println("CellIconUpdate: CellIconUpdated", updatedCellIcon.Name)

	switch instancesTyped := any(cellDeleteIconImpl.Instance).(type) {
	// insertion point
	case *models.ALTERNATIVE_ID:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.ATTRIBUTE_DEFINITION_BOOLEAN:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.ATTRIBUTE_DEFINITION_DATE:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.ATTRIBUTE_DEFINITION_ENUMERATION:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.ATTRIBUTE_DEFINITION_INTEGER:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.ATTRIBUTE_DEFINITION_REAL:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.ATTRIBUTE_DEFINITION_STRING:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.ATTRIBUTE_DEFINITION_XHTML:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.ATTRIBUTE_VALUE_BOOLEAN:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.ATTRIBUTE_VALUE_DATE:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.ATTRIBUTE_VALUE_ENUMERATION:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.ATTRIBUTE_VALUE_INTEGER:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.ATTRIBUTE_VALUE_REAL:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.ATTRIBUTE_VALUE_STRING:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.ATTRIBUTE_VALUE_XHTML:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.A_ALTERNATIVE_ID:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.A_ATTRIBUTE_DEFINITION_BOOLEAN_REF:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.A_ATTRIBUTE_DEFINITION_DATE_REF:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.A_ATTRIBUTE_DEFINITION_ENUMERATION_REF:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.A_ATTRIBUTE_DEFINITION_INTEGER_REF:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.A_ATTRIBUTE_DEFINITION_REAL_REF:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.A_ATTRIBUTE_DEFINITION_STRING_REF:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.A_ATTRIBUTE_DEFINITION_XHTML_REF:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.A_ATTRIBUTE_VALUE_BOOLEAN:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.A_ATTRIBUTE_VALUE_DATE:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.A_ATTRIBUTE_VALUE_ENUMERATION:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.A_ATTRIBUTE_VALUE_INTEGER:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.A_ATTRIBUTE_VALUE_REAL:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.A_ATTRIBUTE_VALUE_STRING:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.A_ATTRIBUTE_VALUE_XHTML:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.A_ATTRIBUTE_VALUE_XHTML_1:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.A_CHILDREN:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.A_CORE_CONTENT:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.A_DATATYPES:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.A_DATATYPE_DEFINITION_BOOLEAN_REF:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.A_DATATYPE_DEFINITION_DATE_REF:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.A_DATATYPE_DEFINITION_ENUMERATION_REF:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.A_DATATYPE_DEFINITION_INTEGER_REF:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.A_DATATYPE_DEFINITION_REAL_REF:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.A_DATATYPE_DEFINITION_STRING_REF:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.A_DATATYPE_DEFINITION_XHTML_REF:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.A_EDITABLE_ATTS:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.A_ENUM_VALUE_REF:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.A_OBJECT:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.A_PROPERTIES:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.A_RELATION_GROUP_TYPE_REF:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.A_SOURCE:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.A_SOURCE_SPECIFICATION:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.A_SPECIFICATIONS:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.A_SPECIFICATION_TYPE_REF:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.A_SPECIFIED_VALUES:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.A_SPEC_ATTRIBUTES:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.A_SPEC_OBJECTS:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.A_SPEC_OBJECT_TYPE_REF:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.A_SPEC_RELATIONS:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.A_SPEC_RELATION_GROUPS:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.A_SPEC_RELATION_REF:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.A_SPEC_RELATION_TYPE_REF:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.A_SPEC_TYPES:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.A_THE_HEADER:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.A_TOOL_EXTENSIONS:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.DATATYPE_DEFINITION_BOOLEAN:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.DATATYPE_DEFINITION_DATE:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.DATATYPE_DEFINITION_ENUMERATION:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.DATATYPE_DEFINITION_INTEGER:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.DATATYPE_DEFINITION_REAL:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.DATATYPE_DEFINITION_STRING:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.DATATYPE_DEFINITION_XHTML:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.EMBEDDED_VALUE:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.ENUM_VALUE:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.RELATION_GROUP:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.RELATION_GROUP_TYPE:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.REQ_IF:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.REQ_IF_CONTENT:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.REQ_IF_HEADER:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.REQ_IF_TOOL_EXTENSION:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.SPECIFICATION:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.SPECIFICATION_TYPE:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.SPEC_HIERARCHY:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.SPEC_OBJECT:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.SPEC_OBJECT_TYPE:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.SPEC_RELATION:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.SPEC_RELATION_TYPE:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.XHTML_CONTENT:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	default:
		_ = instancesTyped
	}
	cellDeleteIconImpl.probe.stageOfInterest.Commit()

	fillUpTable[T](cellDeleteIconImpl.probe)
	fillUpTree(cellDeleteIconImpl.probe)
	cellDeleteIconImpl.probe.tableStage.Commit()
}

