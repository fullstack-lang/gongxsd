package diagrams

import (
	"time"

	"github.com/fullstack-lang/gong/lib/doc2/go/models"
	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gongxsd/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ time.Time

// _ point for meta package dummy declaration
var _ ref_models.Stage

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.Stage) {

	const __write__local_time = "2025-08-17 17:42:08.896454 CEST"
	const __write__utc_time__ = "2025-08-17 15:42:08.896454 UTC"

	const __commitId__ = "0000000007"

	// Declaration of instances to stage

	__Classdiagram__000000_Default := (&models.Classdiagram{}).Stage(stage)

	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_07T07_09_27Z := (&models.DiagramPackage{}).Stage(stage)

	__GongStructShape__000000_Default_ComplexType := (&models.GongStructShape{}).Stage(stage)

	// Setup of values

	__Classdiagram__000000_Default.Name = `Default`
	__Classdiagram__000000_Default.Description = ``
	__Classdiagram__000000_Default.IsIncludedInStaticWebSite = false
	__Classdiagram__000000_Default.IsInRenameMode = false
	__Classdiagram__000000_Default.IsExpanded = true
	__Classdiagram__000000_Default.NodeGongStructsIsExpanded = true
	__Classdiagram__000000_Default.NodeGongStructNodeExpansion = `[false,false,false,false,false,false,true]`
	__Classdiagram__000000_Default.NodeGongEnumsIsExpanded = false
	__Classdiagram__000000_Default.NodeGongEnumNodeExpansion = ``
	__Classdiagram__000000_Default.NodeGongNotesIsExpanded = false
	__Classdiagram__000000_Default.NodeGongNoteNodeExpansion = ``

	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_07T07_09_27Z.Name = `Diagram Package created the 2025-05-07T07:09:27Z`
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_07T07_09_27Z.Path = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_07T07_09_27Z.GongModelPath = ``
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_07T07_09_27Z.AbsolutePathToDiagramPackage = ``

	__GongStructShape__000000_Default_ComplexType.Name = `Default-ComplexType`
	__GongStructShape__000000_Default_ComplexType.X = 58.000000
	__GongStructShape__000000_Default_ComplexType.Y = 70.000000
	__GongStructShape__000000_Default_ComplexType.IdentifierMeta = ref_models.ComplexType{}
	__GongStructShape__000000_Default_ComplexType.Width = 240.000000
	__GongStructShape__000000_Default_ComplexType.Height = 63.000000
	__GongStructShape__000000_Default_ComplexType.IsSelected = false

	// Setup of pointers
	// setup of Classdiagram instances pointers
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000000_Default_ComplexType)
	// setup of DiagramPackage instances pointers
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_07T07_09_27Z.Classdiagrams = append(__DiagramPackage__000000_Diagram_Package_created_the_2025_05_07T07_09_27Z.Classdiagrams, __Classdiagram__000000_Default)
	__DiagramPackage__000000_Diagram_Package_created_the_2025_05_07T07_09_27Z.SelectedClassdiagram = __Classdiagram__000000_Default
	// setup of GongStructShape instances pointers
}

