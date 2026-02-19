// generated code - do not edit
package probe

import (
	gong_models "github.com/fullstack-lang/gong/go/models"
	gongtree_models "github.com/fullstack-lang/gong/lib/tree/go/models"

	"github.com/fullstack-lang/gongxsd/go/models"
)

type TreeNodeImplGongstruct struct {
	gongStruct *gong_models.GongStruct
	probe      *Probe
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
	gongtreeStage *gongtree_models.Stage,
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
	// log.Println("NodeImplGongstruct:OnAfterUpdate with: ", nodeImplGongstruct.gongStruct.GetName())

	// insertion point
	if nodeImplGongstruct.gongStruct.GetName() == "All" {
		updateProbeTable[*models.All](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Annotation" {
		updateProbeTable[*models.Annotation](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Attribute" {
		updateProbeTable[*models.Attribute](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "AttributeGroup" {
		updateProbeTable[*models.AttributeGroup](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Choice" {
		updateProbeTable[*models.Choice](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ComplexContent" {
		updateProbeTable[*models.ComplexContent](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "ComplexType" {
		updateProbeTable[*models.ComplexType](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Documentation" {
		updateProbeTable[*models.Documentation](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Element" {
		updateProbeTable[*models.Element](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Enumeration" {
		updateProbeTable[*models.Enumeration](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Extension" {
		updateProbeTable[*models.Extension](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Group" {
		updateProbeTable[*models.Group](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Length" {
		updateProbeTable[*models.Length](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "MaxInclusive" {
		updateProbeTable[*models.MaxInclusive](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "MaxLength" {
		updateProbeTable[*models.MaxLength](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "MinInclusive" {
		updateProbeTable[*models.MinInclusive](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "MinLength" {
		updateProbeTable[*models.MinLength](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Pattern" {
		updateProbeTable[*models.Pattern](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Restriction" {
		updateProbeTable[*models.Restriction](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Schema" {
		updateProbeTable[*models.Schema](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Sequence" {
		updateProbeTable[*models.Sequence](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SimpleContent" {
		updateProbeTable[*models.SimpleContent](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "SimpleType" {
		updateProbeTable[*models.SimpleType](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "TotalDigit" {
		updateProbeTable[*models.TotalDigit](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "Union" {
		updateProbeTable[*models.Union](nodeImplGongstruct.probe)
	}
	if nodeImplGongstruct.gongStruct.GetName() == "WhiteSpace" {
		updateProbeTable[*models.WhiteSpace](nodeImplGongstruct.probe)
	}

	// set color for node and reset all other nodes color
	for node := range *gongtree_models.GetGongstructInstancesSet[gongtree_models.Node](gongtreeStage) {
		node.BackgroundColor = ""
	}
	stagedNode.BackgroundColor = "lightgrey"
	gongtreeStage.Commit()
}
