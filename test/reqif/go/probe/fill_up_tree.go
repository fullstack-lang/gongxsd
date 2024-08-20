package probe

import (
	"fmt"
	"sort"
	"strings"

	gongtree_buttons "github.com/fullstack-lang/gongtree/go/buttons"
	tree "github.com/fullstack-lang/gongtree/go/models"

	gong_models "github.com/fullstack-lang/gong/go/models"

	"github.com/fullstack-lang/gongxsd/test/reqif/go/models"
)

func fillUpTree(
	probe *Probe,
) {
	// keep in memory which nodes have been unfolded / folded
	expandedNodesSet := make(map[string]any, 0)
	var _sidebar *tree.Tree
	for __sidebar := range probe.treeStage.Trees {
		_sidebar = __sidebar
	}
	if _sidebar != nil {
		for _, node := range _sidebar.RootNodes {
			if node.IsExpanded {
				expandedNodesSet[strings.Fields(node.Name)[0]] = true
			}
		}
	}

	probe.treeStage.Reset()

	// create tree
	sidebar := (&tree.Tree{Name: "gong"}).Stage(probe.treeStage)

	// collect all gong struct to construe the true
	setOfGongStructs := *gong_models.GetGongstructInstancesSet[gong_models.GongStruct](probe.gongStage)

	sliceOfGongStructsSorted := make([]*gong_models.GongStruct, len(setOfGongStructs))
	i := 0
	for k := range setOfGongStructs {
		sliceOfGongStructsSorted[i] = k
		i++
	}
	sort.Slice(sliceOfGongStructsSorted, func(i, j int) bool {
		return sliceOfGongStructsSorted[i].Name < sliceOfGongStructsSorted[j].Name
	})

	for _, gongStruct := range sliceOfGongStructsSorted {

		name := gongStruct.Name + " (" +
			fmt.Sprintf("%d", probe.stageOfInterest.Map_GongStructName_InstancesNb[gongStruct.Name]) + ")"

		nodeGongstruct := (&tree.Node{Name: name}).Stage(probe.treeStage)

		nodeGongstruct.IsExpanded = false
		if _, ok := expandedNodesSet[strings.Fields(name)[0]]; ok {
			nodeGongstruct.IsExpanded = true
		}

		switch gongStruct.Name {
		// insertion point
		case "ALTERNATIVE_ID":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.ALTERNATIVE_ID](probe.stageOfInterest)
			for _alternative_id := range set {
				nodeInstance := (&tree.Node{Name: _alternative_id.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_alternative_id, "ALTERNATIVE_ID", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "ATTRIBUTE_DEFINITION_BOOLEAN":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.ATTRIBUTE_DEFINITION_BOOLEAN](probe.stageOfInterest)
			for _attribute_definition_boolean := range set {
				nodeInstance := (&tree.Node{Name: _attribute_definition_boolean.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_attribute_definition_boolean, "ATTRIBUTE_DEFINITION_BOOLEAN", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "ATTRIBUTE_DEFINITION_DATE":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.ATTRIBUTE_DEFINITION_DATE](probe.stageOfInterest)
			for _attribute_definition_date := range set {
				nodeInstance := (&tree.Node{Name: _attribute_definition_date.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_attribute_definition_date, "ATTRIBUTE_DEFINITION_DATE", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "ATTRIBUTE_DEFINITION_ENUMERATION":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.ATTRIBUTE_DEFINITION_ENUMERATION](probe.stageOfInterest)
			for _attribute_definition_enumeration := range set {
				nodeInstance := (&tree.Node{Name: _attribute_definition_enumeration.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_attribute_definition_enumeration, "ATTRIBUTE_DEFINITION_ENUMERATION", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "ATTRIBUTE_DEFINITION_INTEGER":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.ATTRIBUTE_DEFINITION_INTEGER](probe.stageOfInterest)
			for _attribute_definition_integer := range set {
				nodeInstance := (&tree.Node{Name: _attribute_definition_integer.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_attribute_definition_integer, "ATTRIBUTE_DEFINITION_INTEGER", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "ATTRIBUTE_DEFINITION_REAL":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.ATTRIBUTE_DEFINITION_REAL](probe.stageOfInterest)
			for _attribute_definition_real := range set {
				nodeInstance := (&tree.Node{Name: _attribute_definition_real.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_attribute_definition_real, "ATTRIBUTE_DEFINITION_REAL", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "ATTRIBUTE_DEFINITION_STRING":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.ATTRIBUTE_DEFINITION_STRING](probe.stageOfInterest)
			for _attribute_definition_string := range set {
				nodeInstance := (&tree.Node{Name: _attribute_definition_string.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_attribute_definition_string, "ATTRIBUTE_DEFINITION_STRING", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "ATTRIBUTE_DEFINITION_XHTML":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.ATTRIBUTE_DEFINITION_XHTML](probe.stageOfInterest)
			for _attribute_definition_xhtml := range set {
				nodeInstance := (&tree.Node{Name: _attribute_definition_xhtml.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_attribute_definition_xhtml, "ATTRIBUTE_DEFINITION_XHTML", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "ATTRIBUTE_VALUE_BOOLEAN":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.ATTRIBUTE_VALUE_BOOLEAN](probe.stageOfInterest)
			for _attribute_value_boolean := range set {
				nodeInstance := (&tree.Node{Name: _attribute_value_boolean.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_attribute_value_boolean, "ATTRIBUTE_VALUE_BOOLEAN", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "ATTRIBUTE_VALUE_DATE":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.ATTRIBUTE_VALUE_DATE](probe.stageOfInterest)
			for _attribute_value_date := range set {
				nodeInstance := (&tree.Node{Name: _attribute_value_date.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_attribute_value_date, "ATTRIBUTE_VALUE_DATE", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "ATTRIBUTE_VALUE_ENUMERATION":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.ATTRIBUTE_VALUE_ENUMERATION](probe.stageOfInterest)
			for _attribute_value_enumeration := range set {
				nodeInstance := (&tree.Node{Name: _attribute_value_enumeration.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_attribute_value_enumeration, "ATTRIBUTE_VALUE_ENUMERATION", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "ATTRIBUTE_VALUE_INTEGER":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.ATTRIBUTE_VALUE_INTEGER](probe.stageOfInterest)
			for _attribute_value_integer := range set {
				nodeInstance := (&tree.Node{Name: _attribute_value_integer.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_attribute_value_integer, "ATTRIBUTE_VALUE_INTEGER", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "ATTRIBUTE_VALUE_REAL":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.ATTRIBUTE_VALUE_REAL](probe.stageOfInterest)
			for _attribute_value_real := range set {
				nodeInstance := (&tree.Node{Name: _attribute_value_real.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_attribute_value_real, "ATTRIBUTE_VALUE_REAL", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "ATTRIBUTE_VALUE_STRING":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.ATTRIBUTE_VALUE_STRING](probe.stageOfInterest)
			for _attribute_value_string := range set {
				nodeInstance := (&tree.Node{Name: _attribute_value_string.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_attribute_value_string, "ATTRIBUTE_VALUE_STRING", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "ATTRIBUTE_VALUE_XHTML":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.ATTRIBUTE_VALUE_XHTML](probe.stageOfInterest)
			for _attribute_value_xhtml := range set {
				nodeInstance := (&tree.Node{Name: _attribute_value_xhtml.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_attribute_value_xhtml, "ATTRIBUTE_VALUE_XHTML", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "DATATYPE_DEFINITION_BOOLEAN":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.DATATYPE_DEFINITION_BOOLEAN](probe.stageOfInterest)
			for _datatype_definition_boolean := range set {
				nodeInstance := (&tree.Node{Name: _datatype_definition_boolean.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_datatype_definition_boolean, "DATATYPE_DEFINITION_BOOLEAN", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "DATATYPE_DEFINITION_DATE":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.DATATYPE_DEFINITION_DATE](probe.stageOfInterest)
			for _datatype_definition_date := range set {
				nodeInstance := (&tree.Node{Name: _datatype_definition_date.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_datatype_definition_date, "DATATYPE_DEFINITION_DATE", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "DATATYPE_DEFINITION_ENUMERATION":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.DATATYPE_DEFINITION_ENUMERATION](probe.stageOfInterest)
			for _datatype_definition_enumeration := range set {
				nodeInstance := (&tree.Node{Name: _datatype_definition_enumeration.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_datatype_definition_enumeration, "DATATYPE_DEFINITION_ENUMERATION", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "DATATYPE_DEFINITION_INTEGER":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.DATATYPE_DEFINITION_INTEGER](probe.stageOfInterest)
			for _datatype_definition_integer := range set {
				nodeInstance := (&tree.Node{Name: _datatype_definition_integer.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_datatype_definition_integer, "DATATYPE_DEFINITION_INTEGER", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "DATATYPE_DEFINITION_REAL":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.DATATYPE_DEFINITION_REAL](probe.stageOfInterest)
			for _datatype_definition_real := range set {
				nodeInstance := (&tree.Node{Name: _datatype_definition_real.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_datatype_definition_real, "DATATYPE_DEFINITION_REAL", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "DATATYPE_DEFINITION_STRING":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.DATATYPE_DEFINITION_STRING](probe.stageOfInterest)
			for _datatype_definition_string := range set {
				nodeInstance := (&tree.Node{Name: _datatype_definition_string.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_datatype_definition_string, "DATATYPE_DEFINITION_STRING", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "DATATYPE_DEFINITION_XHTML":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.DATATYPE_DEFINITION_XHTML](probe.stageOfInterest)
			for _datatype_definition_xhtml := range set {
				nodeInstance := (&tree.Node{Name: _datatype_definition_xhtml.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_datatype_definition_xhtml, "DATATYPE_DEFINITION_XHTML", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "EMBEDDED_VALUE":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.EMBEDDED_VALUE](probe.stageOfInterest)
			for _embedded_value := range set {
				nodeInstance := (&tree.Node{Name: _embedded_value.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_embedded_value, "EMBEDDED_VALUE", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "ENUM_VALUE":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.ENUM_VALUE](probe.stageOfInterest)
			for _enum_value := range set {
				nodeInstance := (&tree.Node{Name: _enum_value.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_enum_value, "ENUM_VALUE", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "RELATION_GROUP":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.RELATION_GROUP](probe.stageOfInterest)
			for _relation_group := range set {
				nodeInstance := (&tree.Node{Name: _relation_group.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_relation_group, "RELATION_GROUP", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "RELATION_GROUP_TYPE":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.RELATION_GROUP_TYPE](probe.stageOfInterest)
			for _relation_group_type := range set {
				nodeInstance := (&tree.Node{Name: _relation_group_type.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_relation_group_type, "RELATION_GROUP_TYPE", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "REQ_IF":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.REQ_IF](probe.stageOfInterest)
			for _req_if := range set {
				nodeInstance := (&tree.Node{Name: _req_if.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_req_if, "REQ_IF", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "REQ_IF_CONTENT":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.REQ_IF_CONTENT](probe.stageOfInterest)
			for _req_if_content := range set {
				nodeInstance := (&tree.Node{Name: _req_if_content.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_req_if_content, "REQ_IF_CONTENT", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "REQ_IF_HEADER":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.REQ_IF_HEADER](probe.stageOfInterest)
			for _req_if_header := range set {
				nodeInstance := (&tree.Node{Name: _req_if_header.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_req_if_header, "REQ_IF_HEADER", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "REQ_IF_TOOL_EXTENSION":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.REQ_IF_TOOL_EXTENSION](probe.stageOfInterest)
			for _req_if_tool_extension := range set {
				nodeInstance := (&tree.Node{Name: _req_if_tool_extension.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_req_if_tool_extension, "REQ_IF_TOOL_EXTENSION", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "SPECIFICATION":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.SPECIFICATION](probe.stageOfInterest)
			for _specification := range set {
				nodeInstance := (&tree.Node{Name: _specification.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_specification, "SPECIFICATION", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "SPECIFICATION_TYPE":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.SPECIFICATION_TYPE](probe.stageOfInterest)
			for _specification_type := range set {
				nodeInstance := (&tree.Node{Name: _specification_type.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_specification_type, "SPECIFICATION_TYPE", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "SPEC_HIERARCHY":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.SPEC_HIERARCHY](probe.stageOfInterest)
			for _spec_hierarchy := range set {
				nodeInstance := (&tree.Node{Name: _spec_hierarchy.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_spec_hierarchy, "SPEC_HIERARCHY", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "SPEC_OBJECT":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.SPEC_OBJECT](probe.stageOfInterest)
			for _spec_object := range set {
				nodeInstance := (&tree.Node{Name: _spec_object.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_spec_object, "SPEC_OBJECT", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "SPEC_OBJECT_TYPE":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.SPEC_OBJECT_TYPE](probe.stageOfInterest)
			for _spec_object_type := range set {
				nodeInstance := (&tree.Node{Name: _spec_object_type.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_spec_object_type, "SPEC_OBJECT_TYPE", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "SPEC_RELATION":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.SPEC_RELATION](probe.stageOfInterest)
			for _spec_relation := range set {
				nodeInstance := (&tree.Node{Name: _spec_relation.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_spec_relation, "SPEC_RELATION", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "SPEC_RELATION_TYPE":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.SPEC_RELATION_TYPE](probe.stageOfInterest)
			for _spec_relation_type := range set {
				nodeInstance := (&tree.Node{Name: _spec_relation_type.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_spec_relation_type, "SPEC_RELATION_TYPE", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "XHTML_CONTENT":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.XHTML_CONTENT](probe.stageOfInterest)
			for _xhtml_content := range set {
				nodeInstance := (&tree.Node{Name: _xhtml_content.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_xhtml_content, "XHTML_CONTENT", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		}

		nodeGongstruct.IsNodeClickable = true
		nodeGongstruct.Impl = NewTreeNodeImplGongstruct(gongStruct, probe)

		// add add button
		addButton := (&tree.Button{
			Name: gongStruct.Name + " " + string(gongtree_buttons.BUTTON_add),
			Icon: string(gongtree_buttons.BUTTON_add)}).Stage(probe.treeStage)
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)
		addButton.Impl = NewButtonImplGongstruct(
			gongStruct,
			gongtree_buttons.BUTTON_add,
			probe,
		)

		sidebar.RootNodes = append(sidebar.RootNodes, nodeGongstruct)
	}

	// Add a refresh button
	nodeRefreshButton := (&tree.Node{Name: ""}).Stage(probe.treeStage)
	sidebar.RootNodes = append(sidebar.RootNodes, nodeRefreshButton)
	refreshButton := (&tree.Button{
		Name: "RefreshButton" + " " + string(gongtree_buttons.BUTTON_refresh),
		Icon: string(gongtree_buttons.BUTTON_refresh)}).Stage(probe.treeStage)
	nodeRefreshButton.Buttons = append(nodeRefreshButton.Buttons, refreshButton)
	refreshButton.Impl = NewButtonImplRefresh(probe)

	probe.treeStage.Commit()
}

type InstanceNodeCallback[T models.Gongstruct] struct {
	Instance       *T
	gongstructName string
	probe          *Probe
}

func NewInstanceNodeCallback[T models.Gongstruct](
	instance *T,
	gongstructName string,
	probe *Probe) (
	instanceNodeCallback *InstanceNodeCallback[T],
) {
	instanceNodeCallback = new(InstanceNodeCallback[T])

	instanceNodeCallback.probe = probe
	instanceNodeCallback.gongstructName = gongstructName
	instanceNodeCallback.Instance = instance

	return
}

func (instanceNodeCallback *InstanceNodeCallback[T]) OnAfterUpdate(
	gongtreeStage *tree.StageStruct,
	stagedNode, frontNode *tree.Node) {

	FillUpFormFromGongstruct(
		instanceNodeCallback.Instance,
		instanceNodeCallback.probe,
	)
}
