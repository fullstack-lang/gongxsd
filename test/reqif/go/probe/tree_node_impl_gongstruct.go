// generated code - do not edit
package probe

import (
	"log"

	gong_models "github.com/fullstack-lang/gong/go/models"
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"

	"github.com/fullstack-lang/gongxsd/test/reqif/go/models"
)

type TreeNodeImplGongstruct struct {
	gongStruct *gong_models.GongStruct
	probe *Probe
}

func NewTreeNodeImplGongstruct(
	gongStruct *gong_models.GongStruct,
	probe *Probe,
) (nodeImplGongstruct *TreeNodeImplGongstruct) {

	nodeImplGongstruct = new(TreeNodeImplGongstruct)
	nodeImplGongstruct.gongStruct = gongStruct
	nodeImplGongstruct.probe = probe
	return
}

func (nodeImplGongstruct *TreeNodeImplGongstruct) OnAfterUpdate(
	gongtreeStage *gongtree_models.StageStruct,
	stagedNode, frontNode *gongtree_models.Node) {

	// setting the value of the staged node	to the new value
	// otherwise, the expansion memory is lost
	if stagedNode.IsExpanded != frontNode.IsExpanded {
		stagedNode.IsExpanded = frontNode.IsExpanded
		return
	}

	// if node is unchecked
	if stagedNode.IsChecked && !frontNode.IsChecked {

	}

	// if node is checked, add gongstructshape
	if !stagedNode.IsChecked && frontNode.IsChecked {

	}

	// the node was selected. Therefore, one request the
	// table to route to the table
	log.Println("NodeImplGongstruct:OnAfterUpdate with: ", nodeImplGongstruct.gongStruct.GetName())

	// insertion point
	if nodeImplGongstruct.gongStruct.GetName() == "ALTERNATIVE_ID" {
		fillUpTable[models.ALTERNATIVE_ID](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ATTRIBUTE_DEFINITION_BOOLEAN" {
		fillUpTable[models.ATTRIBUTE_DEFINITION_BOOLEAN](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ATTRIBUTE_DEFINITION_DATE" {
		fillUpTable[models.ATTRIBUTE_DEFINITION_DATE](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ATTRIBUTE_DEFINITION_ENUMERATION" {
		fillUpTable[models.ATTRIBUTE_DEFINITION_ENUMERATION](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ATTRIBUTE_DEFINITION_INTEGER" {
		fillUpTable[models.ATTRIBUTE_DEFINITION_INTEGER](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ATTRIBUTE_DEFINITION_REAL" {
		fillUpTable[models.ATTRIBUTE_DEFINITION_REAL](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ATTRIBUTE_DEFINITION_STRING" {
		fillUpTable[models.ATTRIBUTE_DEFINITION_STRING](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ATTRIBUTE_DEFINITION_XHTML" {
		fillUpTable[models.ATTRIBUTE_DEFINITION_XHTML](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ATTRIBUTE_VALUE_BOOLEAN" {
		fillUpTable[models.ATTRIBUTE_VALUE_BOOLEAN](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ATTRIBUTE_VALUE_DATE" {
		fillUpTable[models.ATTRIBUTE_VALUE_DATE](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ATTRIBUTE_VALUE_ENUMERATION" {
		fillUpTable[models.ATTRIBUTE_VALUE_ENUMERATION](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ATTRIBUTE_VALUE_INTEGER" {
		fillUpTable[models.ATTRIBUTE_VALUE_INTEGER](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ATTRIBUTE_VALUE_REAL" {
		fillUpTable[models.ATTRIBUTE_VALUE_REAL](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ATTRIBUTE_VALUE_STRING" {
		fillUpTable[models.ATTRIBUTE_VALUE_STRING](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ATTRIBUTE_VALUE_XHTML" {
		fillUpTable[models.ATTRIBUTE_VALUE_XHTML](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_ALTERNATIVE_ID" {
		fillUpTable[models.A_ALTERNATIVE_ID](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_ATTRIBUTE_DEFINITION_BOOLEAN_REF" {
		fillUpTable[models.A_ATTRIBUTE_DEFINITION_BOOLEAN_REF](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_ATTRIBUTE_DEFINITION_DATE_REF" {
		fillUpTable[models.A_ATTRIBUTE_DEFINITION_DATE_REF](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_ATTRIBUTE_DEFINITION_ENUMERATION_REF" {
		fillUpTable[models.A_ATTRIBUTE_DEFINITION_ENUMERATION_REF](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_ATTRIBUTE_DEFINITION_INTEGER_REF" {
		fillUpTable[models.A_ATTRIBUTE_DEFINITION_INTEGER_REF](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_ATTRIBUTE_DEFINITION_REAL_REF" {
		fillUpTable[models.A_ATTRIBUTE_DEFINITION_REAL_REF](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_ATTRIBUTE_DEFINITION_STRING_REF" {
		fillUpTable[models.A_ATTRIBUTE_DEFINITION_STRING_REF](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_ATTRIBUTE_DEFINITION_XHTML_REF" {
		fillUpTable[models.A_ATTRIBUTE_DEFINITION_XHTML_REF](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_ATTRIBUTE_VALUE_BOOLEAN" {
		fillUpTable[models.A_ATTRIBUTE_VALUE_BOOLEAN](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_ATTRIBUTE_VALUE_DATE" {
		fillUpTable[models.A_ATTRIBUTE_VALUE_DATE](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_ATTRIBUTE_VALUE_ENUMERATION" {
		fillUpTable[models.A_ATTRIBUTE_VALUE_ENUMERATION](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_ATTRIBUTE_VALUE_INTEGER" {
		fillUpTable[models.A_ATTRIBUTE_VALUE_INTEGER](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_ATTRIBUTE_VALUE_REAL" {
		fillUpTable[models.A_ATTRIBUTE_VALUE_REAL](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_ATTRIBUTE_VALUE_STRING" {
		fillUpTable[models.A_ATTRIBUTE_VALUE_STRING](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_ATTRIBUTE_VALUE_XHTML" {
		fillUpTable[models.A_ATTRIBUTE_VALUE_XHTML](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_ATTRIBUTE_VALUE_XHTML_1" {
		fillUpTable[models.A_ATTRIBUTE_VALUE_XHTML_1](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_CHILDREN" {
		fillUpTable[models.A_CHILDREN](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_CORE_CONTENT" {
		fillUpTable[models.A_CORE_CONTENT](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_DATATYPES" {
		fillUpTable[models.A_DATATYPES](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_DATATYPE_DEFINITION_BOOLEAN_REF" {
		fillUpTable[models.A_DATATYPE_DEFINITION_BOOLEAN_REF](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_DATATYPE_DEFINITION_DATE_REF" {
		fillUpTable[models.A_DATATYPE_DEFINITION_DATE_REF](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_DATATYPE_DEFINITION_ENUMERATION_REF" {
		fillUpTable[models.A_DATATYPE_DEFINITION_ENUMERATION_REF](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_DATATYPE_DEFINITION_INTEGER_REF" {
		fillUpTable[models.A_DATATYPE_DEFINITION_INTEGER_REF](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_DATATYPE_DEFINITION_REAL_REF" {
		fillUpTable[models.A_DATATYPE_DEFINITION_REAL_REF](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_DATATYPE_DEFINITION_STRING_REF" {
		fillUpTable[models.A_DATATYPE_DEFINITION_STRING_REF](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_DATATYPE_DEFINITION_XHTML_REF" {
		fillUpTable[models.A_DATATYPE_DEFINITION_XHTML_REF](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_EDITABLE_ATTS" {
		fillUpTable[models.A_EDITABLE_ATTS](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_ENUM_VALUE_REF" {
		fillUpTable[models.A_ENUM_VALUE_REF](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_OBJECT" {
		fillUpTable[models.A_OBJECT](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_PROPERTIES" {
		fillUpTable[models.A_PROPERTIES](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_RELATION_GROUP_TYPE_REF" {
		fillUpTable[models.A_RELATION_GROUP_TYPE_REF](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_SOURCE" {
		fillUpTable[models.A_SOURCE](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_SOURCE_SPECIFICATION" {
		fillUpTable[models.A_SOURCE_SPECIFICATION](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_SPECIFICATIONS" {
		fillUpTable[models.A_SPECIFICATIONS](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_SPECIFICATION_TYPE_REF" {
		fillUpTable[models.A_SPECIFICATION_TYPE_REF](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_SPECIFIED_VALUES" {
		fillUpTable[models.A_SPECIFIED_VALUES](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_SPEC_ATTRIBUTES" {
		fillUpTable[models.A_SPEC_ATTRIBUTES](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_SPEC_OBJECTS" {
		fillUpTable[models.A_SPEC_OBJECTS](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_SPEC_OBJECT_TYPE_REF" {
		fillUpTable[models.A_SPEC_OBJECT_TYPE_REF](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_SPEC_RELATIONS" {
		fillUpTable[models.A_SPEC_RELATIONS](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_SPEC_RELATION_GROUPS" {
		fillUpTable[models.A_SPEC_RELATION_GROUPS](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_SPEC_RELATION_REF" {
		fillUpTable[models.A_SPEC_RELATION_REF](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_SPEC_RELATION_TYPE_REF" {
		fillUpTable[models.A_SPEC_RELATION_TYPE_REF](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_SPEC_TYPES" {
		fillUpTable[models.A_SPEC_TYPES](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_THE_HEADER" {
		fillUpTable[models.A_THE_HEADER](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "A_TOOL_EXTENSIONS" {
		fillUpTable[models.A_TOOL_EXTENSIONS](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "DATATYPE_DEFINITION_BOOLEAN" {
		fillUpTable[models.DATATYPE_DEFINITION_BOOLEAN](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "DATATYPE_DEFINITION_DATE" {
		fillUpTable[models.DATATYPE_DEFINITION_DATE](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "DATATYPE_DEFINITION_ENUMERATION" {
		fillUpTable[models.DATATYPE_DEFINITION_ENUMERATION](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "DATATYPE_DEFINITION_INTEGER" {
		fillUpTable[models.DATATYPE_DEFINITION_INTEGER](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "DATATYPE_DEFINITION_REAL" {
		fillUpTable[models.DATATYPE_DEFINITION_REAL](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "DATATYPE_DEFINITION_STRING" {
		fillUpTable[models.DATATYPE_DEFINITION_STRING](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "DATATYPE_DEFINITION_XHTML" {
		fillUpTable[models.DATATYPE_DEFINITION_XHTML](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "EMBEDDED_VALUE" {
		fillUpTable[models.EMBEDDED_VALUE](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ENUM_VALUE" {
		fillUpTable[models.ENUM_VALUE](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "RELATION_GROUP" {
		fillUpTable[models.RELATION_GROUP](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "RELATION_GROUP_TYPE" {
		fillUpTable[models.RELATION_GROUP_TYPE](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "REQ_IF" {
		fillUpTable[models.REQ_IF](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "REQ_IF_CONTENT" {
		fillUpTable[models.REQ_IF_CONTENT](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "REQ_IF_HEADER" {
		fillUpTable[models.REQ_IF_HEADER](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "REQ_IF_TOOL_EXTENSION" {
		fillUpTable[models.REQ_IF_TOOL_EXTENSION](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SPECIFICATION" {
		fillUpTable[models.SPECIFICATION](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SPECIFICATION_TYPE" {
		fillUpTable[models.SPECIFICATION_TYPE](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SPEC_HIERARCHY" {
		fillUpTable[models.SPEC_HIERARCHY](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SPEC_OBJECT" {
		fillUpTable[models.SPEC_OBJECT](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SPEC_OBJECT_TYPE" {
		fillUpTable[models.SPEC_OBJECT_TYPE](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SPEC_RELATION" {
		fillUpTable[models.SPEC_RELATION](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SPEC_RELATION_TYPE" {
		fillUpTable[models.SPEC_RELATION_TYPE](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "XHTML_CONTENT" {
		fillUpTable[models.XHTML_CONTENT](nodeImplGongstruct.probe)
	}

	// set color for node and reset all other nodes color
	for node := range *gongtree_models.GetGongstructInstancesSet[gongtree_models.Node](gongtreeStage) {
		node.BackgroundColor = ""
	}
	stagedNode.BackgroundColor = "lightgrey"
	gongtreeStage.Commit()

	nodeImplGongstruct.probe.tableStage.Commit()
}
