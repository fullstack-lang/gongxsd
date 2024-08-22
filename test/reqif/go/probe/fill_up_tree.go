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
		case "A_ALTERNATIVE_ID":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_ALTERNATIVE_ID](probe.stageOfInterest)
			for _a_alternative_id := range set {
				nodeInstance := (&tree.Node{Name: _a_alternative_id.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_alternative_id, "A_ALTERNATIVE_ID", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_CHILDREN":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_CHILDREN](probe.stageOfInterest)
			for _a_children := range set {
				nodeInstance := (&tree.Node{Name: _a_children.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_children, "A_CHILDREN", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_CORE_CONTENT":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_CORE_CONTENT](probe.stageOfInterest)
			for _a_core_content := range set {
				nodeInstance := (&tree.Node{Name: _a_core_content.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_core_content, "A_CORE_CONTENT", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_DATATYPES":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_DATATYPES](probe.stageOfInterest)
			for _a_datatypes := range set {
				nodeInstance := (&tree.Node{Name: _a_datatypes.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_datatypes, "A_DATATYPES", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_DEFAULT_VALUE":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_DEFAULT_VALUE](probe.stageOfInterest)
			for _a_default_value := range set {
				nodeInstance := (&tree.Node{Name: _a_default_value.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_default_value, "A_DEFAULT_VALUE", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_DEFAULT_VALUE_1":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_DEFAULT_VALUE_1](probe.stageOfInterest)
			for _a_default_value_1 := range set {
				nodeInstance := (&tree.Node{Name: _a_default_value_1.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_default_value_1, "A_DEFAULT_VALUE_1", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_DEFAULT_VALUE_2":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_DEFAULT_VALUE_2](probe.stageOfInterest)
			for _a_default_value_2 := range set {
				nodeInstance := (&tree.Node{Name: _a_default_value_2.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_default_value_2, "A_DEFAULT_VALUE_2", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_DEFAULT_VALUE_3":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_DEFAULT_VALUE_3](probe.stageOfInterest)
			for _a_default_value_3 := range set {
				nodeInstance := (&tree.Node{Name: _a_default_value_3.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_default_value_3, "A_DEFAULT_VALUE_3", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_DEFAULT_VALUE_4":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_DEFAULT_VALUE_4](probe.stageOfInterest)
			for _a_default_value_4 := range set {
				nodeInstance := (&tree.Node{Name: _a_default_value_4.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_default_value_4, "A_DEFAULT_VALUE_4", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_DEFAULT_VALUE_5":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_DEFAULT_VALUE_5](probe.stageOfInterest)
			for _a_default_value_5 := range set {
				nodeInstance := (&tree.Node{Name: _a_default_value_5.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_default_value_5, "A_DEFAULT_VALUE_5", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_DEFAULT_VALUE_6":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_DEFAULT_VALUE_6](probe.stageOfInterest)
			for _a_default_value_6 := range set {
				nodeInstance := (&tree.Node{Name: _a_default_value_6.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_default_value_6, "A_DEFAULT_VALUE_6", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_DEFINITION":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_DEFINITION](probe.stageOfInterest)
			for _a_definition := range set {
				nodeInstance := (&tree.Node{Name: _a_definition.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_definition, "A_DEFINITION", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_DEFINITION_1":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_DEFINITION_1](probe.stageOfInterest)
			for _a_definition_1 := range set {
				nodeInstance := (&tree.Node{Name: _a_definition_1.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_definition_1, "A_DEFINITION_1", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_DEFINITION_2":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_DEFINITION_2](probe.stageOfInterest)
			for _a_definition_2 := range set {
				nodeInstance := (&tree.Node{Name: _a_definition_2.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_definition_2, "A_DEFINITION_2", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_DEFINITION_3":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_DEFINITION_3](probe.stageOfInterest)
			for _a_definition_3 := range set {
				nodeInstance := (&tree.Node{Name: _a_definition_3.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_definition_3, "A_DEFINITION_3", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_DEFINITION_4":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_DEFINITION_4](probe.stageOfInterest)
			for _a_definition_4 := range set {
				nodeInstance := (&tree.Node{Name: _a_definition_4.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_definition_4, "A_DEFINITION_4", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_DEFINITION_5":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_DEFINITION_5](probe.stageOfInterest)
			for _a_definition_5 := range set {
				nodeInstance := (&tree.Node{Name: _a_definition_5.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_definition_5, "A_DEFINITION_5", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_DEFINITION_6":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_DEFINITION_6](probe.stageOfInterest)
			for _a_definition_6 := range set {
				nodeInstance := (&tree.Node{Name: _a_definition_6.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_definition_6, "A_DEFINITION_6", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_EDITABLE_ATTS":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_EDITABLE_ATTS](probe.stageOfInterest)
			for _a_editable_atts := range set {
				nodeInstance := (&tree.Node{Name: _a_editable_atts.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_editable_atts, "A_EDITABLE_ATTS", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_OBJECT":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_OBJECT](probe.stageOfInterest)
			for _a_object := range set {
				nodeInstance := (&tree.Node{Name: _a_object.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_object, "A_OBJECT", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_PROPERTIES":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_PROPERTIES](probe.stageOfInterest)
			for _a_properties := range set {
				nodeInstance := (&tree.Node{Name: _a_properties.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_properties, "A_PROPERTIES", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_SOURCE":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_SOURCE](probe.stageOfInterest)
			for _a_source := range set {
				nodeInstance := (&tree.Node{Name: _a_source.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_source, "A_SOURCE", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_SOURCE_SPECIFICATION":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_SOURCE_SPECIFICATION](probe.stageOfInterest)
			for _a_source_specification := range set {
				nodeInstance := (&tree.Node{Name: _a_source_specification.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_source_specification, "A_SOURCE_SPECIFICATION", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_SPECIFICATIONS":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_SPECIFICATIONS](probe.stageOfInterest)
			for _a_specifications := range set {
				nodeInstance := (&tree.Node{Name: _a_specifications.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_specifications, "A_SPECIFICATIONS", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_SPECIFIED_VALUES":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_SPECIFIED_VALUES](probe.stageOfInterest)
			for _a_specified_values := range set {
				nodeInstance := (&tree.Node{Name: _a_specified_values.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_specified_values, "A_SPECIFIED_VALUES", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_SPEC_ATTRIBUTES":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_SPEC_ATTRIBUTES](probe.stageOfInterest)
			for _a_spec_attributes := range set {
				nodeInstance := (&tree.Node{Name: _a_spec_attributes.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_spec_attributes, "A_SPEC_ATTRIBUTES", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_SPEC_OBJECTS":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_SPEC_OBJECTS](probe.stageOfInterest)
			for _a_spec_objects := range set {
				nodeInstance := (&tree.Node{Name: _a_spec_objects.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_spec_objects, "A_SPEC_OBJECTS", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_SPEC_RELATIONS":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_SPEC_RELATIONS](probe.stageOfInterest)
			for _a_spec_relations := range set {
				nodeInstance := (&tree.Node{Name: _a_spec_relations.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_spec_relations, "A_SPEC_RELATIONS", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_SPEC_RELATIONS_1":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_SPEC_RELATIONS_1](probe.stageOfInterest)
			for _a_spec_relations_1 := range set {
				nodeInstance := (&tree.Node{Name: _a_spec_relations_1.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_spec_relations_1, "A_SPEC_RELATIONS_1", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_SPEC_RELATION_GROUPS":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_SPEC_RELATION_GROUPS](probe.stageOfInterest)
			for _a_spec_relation_groups := range set {
				nodeInstance := (&tree.Node{Name: _a_spec_relation_groups.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_spec_relation_groups, "A_SPEC_RELATION_GROUPS", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_SPEC_TYPES":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_SPEC_TYPES](probe.stageOfInterest)
			for _a_spec_types := range set {
				nodeInstance := (&tree.Node{Name: _a_spec_types.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_spec_types, "A_SPEC_TYPES", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_THE_HEADER":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_THE_HEADER](probe.stageOfInterest)
			for _a_the_header := range set {
				nodeInstance := (&tree.Node{Name: _a_the_header.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_the_header, "A_THE_HEADER", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_TOOL_EXTENSIONS":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_TOOL_EXTENSIONS](probe.stageOfInterest)
			for _a_tool_extensions := range set {
				nodeInstance := (&tree.Node{Name: _a_tool_extensions.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_tool_extensions, "A_TOOL_EXTENSIONS", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_TYPE":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_TYPE](probe.stageOfInterest)
			for _a_type := range set {
				nodeInstance := (&tree.Node{Name: _a_type.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_type, "A_TYPE", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_TYPE_1":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_TYPE_1](probe.stageOfInterest)
			for _a_type_1 := range set {
				nodeInstance := (&tree.Node{Name: _a_type_1.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_type_1, "A_TYPE_1", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_TYPE_10":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_TYPE_10](probe.stageOfInterest)
			for _a_type_10 := range set {
				nodeInstance := (&tree.Node{Name: _a_type_10.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_type_10, "A_TYPE_10", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_TYPE_2":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_TYPE_2](probe.stageOfInterest)
			for _a_type_2 := range set {
				nodeInstance := (&tree.Node{Name: _a_type_2.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_type_2, "A_TYPE_2", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_TYPE_3":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_TYPE_3](probe.stageOfInterest)
			for _a_type_3 := range set {
				nodeInstance := (&tree.Node{Name: _a_type_3.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_type_3, "A_TYPE_3", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_TYPE_4":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_TYPE_4](probe.stageOfInterest)
			for _a_type_4 := range set {
				nodeInstance := (&tree.Node{Name: _a_type_4.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_type_4, "A_TYPE_4", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_TYPE_5":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_TYPE_5](probe.stageOfInterest)
			for _a_type_5 := range set {
				nodeInstance := (&tree.Node{Name: _a_type_5.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_type_5, "A_TYPE_5", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_TYPE_6":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_TYPE_6](probe.stageOfInterest)
			for _a_type_6 := range set {
				nodeInstance := (&tree.Node{Name: _a_type_6.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_type_6, "A_TYPE_6", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_TYPE_7":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_TYPE_7](probe.stageOfInterest)
			for _a_type_7 := range set {
				nodeInstance := (&tree.Node{Name: _a_type_7.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_type_7, "A_TYPE_7", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_TYPE_8":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_TYPE_8](probe.stageOfInterest)
			for _a_type_8 := range set {
				nodeInstance := (&tree.Node{Name: _a_type_8.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_type_8, "A_TYPE_8", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_TYPE_9":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_TYPE_9](probe.stageOfInterest)
			for _a_type_9 := range set {
				nodeInstance := (&tree.Node{Name: _a_type_9.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_type_9, "A_TYPE_9", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_VALUES":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_VALUES](probe.stageOfInterest)
			for _a_values := range set {
				nodeInstance := (&tree.Node{Name: _a_values.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_values, "A_VALUES", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "A_VALUES_1":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.A_VALUES_1](probe.stageOfInterest)
			for _a_values_1 := range set {
				nodeInstance := (&tree.Node{Name: _a_values_1.GetName()}).Stage(probe.treeStage)
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_a_values_1, "A_VALUES_1", probe)

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
