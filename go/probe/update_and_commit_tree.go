package probe

import (
	"fmt"
	"sort"
	"strings"
	"time"

	gongtree_buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"

	gong_models "github.com/fullstack-lang/gong/go/models"

	"github.com/fullstack-lang/gongxsd/go/models"
)

func updateAndCommitTree(
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
	sidebar := &tree.Tree{Name: "Sidebar"}

	nodeRefreshButton := &tree.Node{Name: fmt.Sprintf("Stage %s, # %d, %s",
		probe.stageOfInterest.GetName(),
		probe.stageOfInterest.GetCommitId(),
		probe.stageOfInterest.GetCommitTS().Local().Format(time.Kitchen))}
	sidebar.RootNodes = append(sidebar.RootNodes, nodeRefreshButton)
	refreshButton := &tree.Button{
		Name:            "RefreshButton" + " " + string(gongtree_buttons.BUTTON_refresh),
		Icon:            string(gongtree_buttons.BUTTON_refresh),
		HasToolTip:      true,
		ToolTipText:     "Refresh probe",
		ToolTipPosition: tree.Left,
	}

	nodeRefreshButton.Buttons = append(nodeRefreshButton.Buttons, refreshButton)
	refreshButton.Impl = NewButtonImplRefresh(probe)

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

		nodeGongstruct := &tree.Node{Name: name}
		nodeGongstruct.HasToolTip = true
		nodeGongstruct.ToolTipText = "Display table of all " + name + " instances"
		nodeGongstruct.ToolTipPosition = tree.Right

		nodeGongstruct.IsExpanded = false
		if _, ok := expandedNodesSet[strings.Fields(name)[0]]; ok {
			nodeGongstruct.IsExpanded = true
		}

		switch gongStruct.Name {
		// insertion point
		case "All":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.All](probe.stageOfInterest)
			for _all := range set {
				nodeInstance := &tree.Node{Name: _all.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_all, "All", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Annotation":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Annotation](probe.stageOfInterest)
			for _annotation := range set {
				nodeInstance := &tree.Node{Name: _annotation.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_annotation, "Annotation", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Attribute":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Attribute](probe.stageOfInterest)
			for _attribute := range set {
				nodeInstance := &tree.Node{Name: _attribute.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_attribute, "Attribute", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "AttributeGroup":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.AttributeGroup](probe.stageOfInterest)
			for _attributegroup := range set {
				nodeInstance := &tree.Node{Name: _attributegroup.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_attributegroup, "AttributeGroup", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Choice":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Choice](probe.stageOfInterest)
			for _choice := range set {
				nodeInstance := &tree.Node{Name: _choice.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_choice, "Choice", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "ComplexContent":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.ComplexContent](probe.stageOfInterest)
			for _complexcontent := range set {
				nodeInstance := &tree.Node{Name: _complexcontent.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_complexcontent, "ComplexContent", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "ComplexType":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.ComplexType](probe.stageOfInterest)
			for _complextype := range set {
				nodeInstance := &tree.Node{Name: _complextype.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_complextype, "ComplexType", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Documentation":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Documentation](probe.stageOfInterest)
			for _documentation := range set {
				nodeInstance := &tree.Node{Name: _documentation.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_documentation, "Documentation", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Element":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Element](probe.stageOfInterest)
			for _element := range set {
				nodeInstance := &tree.Node{Name: _element.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_element, "Element", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Enumeration":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Enumeration](probe.stageOfInterest)
			for _enumeration := range set {
				nodeInstance := &tree.Node{Name: _enumeration.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_enumeration, "Enumeration", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Extension":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Extension](probe.stageOfInterest)
			for _extension := range set {
				nodeInstance := &tree.Node{Name: _extension.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_extension, "Extension", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Group":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Group](probe.stageOfInterest)
			for _group := range set {
				nodeInstance := &tree.Node{Name: _group.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_group, "Group", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Length":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Length](probe.stageOfInterest)
			for _length := range set {
				nodeInstance := &tree.Node{Name: _length.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_length, "Length", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "MaxInclusive":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.MaxInclusive](probe.stageOfInterest)
			for _maxinclusive := range set {
				nodeInstance := &tree.Node{Name: _maxinclusive.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_maxinclusive, "MaxInclusive", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "MaxLength":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.MaxLength](probe.stageOfInterest)
			for _maxlength := range set {
				nodeInstance := &tree.Node{Name: _maxlength.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_maxlength, "MaxLength", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "MinInclusive":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.MinInclusive](probe.stageOfInterest)
			for _mininclusive := range set {
				nodeInstance := &tree.Node{Name: _mininclusive.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_mininclusive, "MinInclusive", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "MinLength":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.MinLength](probe.stageOfInterest)
			for _minlength := range set {
				nodeInstance := &tree.Node{Name: _minlength.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_minlength, "MinLength", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Pattern":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Pattern](probe.stageOfInterest)
			for _pattern := range set {
				nodeInstance := &tree.Node{Name: _pattern.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_pattern, "Pattern", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Restriction":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Restriction](probe.stageOfInterest)
			for _restriction := range set {
				nodeInstance := &tree.Node{Name: _restriction.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_restriction, "Restriction", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Schema":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Schema](probe.stageOfInterest)
			for _schema := range set {
				nodeInstance := &tree.Node{Name: _schema.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_schema, "Schema", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Sequence":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Sequence](probe.stageOfInterest)
			for _sequence := range set {
				nodeInstance := &tree.Node{Name: _sequence.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_sequence, "Sequence", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "SimpleContent":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.SimpleContent](probe.stageOfInterest)
			for _simplecontent := range set {
				nodeInstance := &tree.Node{Name: _simplecontent.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_simplecontent, "SimpleContent", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "SimpleType":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.SimpleType](probe.stageOfInterest)
			for _simpletype := range set {
				nodeInstance := &tree.Node{Name: _simpletype.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_simpletype, "SimpleType", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "TotalDigit":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.TotalDigit](probe.stageOfInterest)
			for _totaldigit := range set {
				nodeInstance := &tree.Node{Name: _totaldigit.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_totaldigit, "TotalDigit", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "Union":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.Union](probe.stageOfInterest)
			for _union := range set {
				nodeInstance := &tree.Node{Name: _union.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_union, "Union", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		case "WhiteSpace":
			nodeGongstruct.Name = name
			set := *models.GetGongstructInstancesSet[models.WhiteSpace](probe.stageOfInterest)
			for _whitespace := range set {
				nodeInstance := &tree.Node{Name: _whitespace.GetName()}
				nodeInstance.IsNodeClickable = true
				nodeInstance.Impl = NewInstanceNodeCallback(_whitespace, "WhiteSpace", probe)

				nodeGongstruct.Children = append(nodeGongstruct.Children, nodeInstance)
			}
		}

		nodeGongstruct.IsNodeClickable = true
		nodeGongstruct.Impl = NewTreeNodeImplGongstruct(gongStruct, probe)

		// add add button
		addButton := &tree.Button{
			Name:            gongStruct.Name + " " + string(gongtree_buttons.BUTTON_add),
			Icon:            string(gongtree_buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipText:     "Add an instance of " + gongStruct.GetName(),
			ToolTipPosition: tree.Right,
		}
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)
		addButton.Impl = NewButtonImplGongstruct(
			gongStruct,
			gongtree_buttons.BUTTON_add,
			probe,
		)

		sidebar.RootNodes = append(sidebar.RootNodes, nodeGongstruct)
	}

	tree.StageBranch(probe.treeStage, sidebar)

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
	gongtreeStage *tree.Stage,
	stagedNode, frontNode *tree.Node) {

	FillUpFormFromGongstruct(
		instanceNodeCallback.Instance,
		instanceNodeCallback.probe,
	)
}
