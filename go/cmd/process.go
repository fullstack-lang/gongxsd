package cmd

import (
	"encoding/xml"
	"fmt"
	"os"
	"strings"
	"time"

	gongxsd_level1stack "github.com/fullstack-lang/gongxsd/go/level1stack"

	// insertion point for models import
	gongxsd_models "github.com/fullstack-lang/gongxsd/go/models"

	"github.com/gin-gonic/gin"
)

func process(args []string) (r *gin.Engine, stack *gongxsd_level1stack.Level1Stack) {
	start := time.Now()
	if Verbose {
		fmt.Printf("generate start\n")
	}

	xsdFilePath := args[0]
	if Verbose {
		fmt.Printf("Reading file: %s\n", xsdFilePath)
	}

	// setup model stack with its probe
	stack = gongxsd_level1stack.NewLevel1Stack("gongxsd", *unmarshallFromCode, *marshallOnCommit, false, true)
	stack.Probe.Refresh()

	// insertion point for call to stager
	gongxsd_models.NewStager(r, stack.Stage, stack.Probe)

	content, err := os.ReadFile(xsdFilePath)
	if err != nil {
		fmt.Printf("Error reading XSD file: %v\n", err)
		os.Exit(1)
	}

	err = xml.Unmarshal(content, &gongxsd_models.SchemaSingloton)
	if err != nil {
		fmt.Printf("Error parsing XML: %v\n", err)
		os.Exit(1)
	}

	// suppress duplicates
	if strings.Contains(xsdFilePath, "dtc-11-04-05") {
		gongxsd_models.SchemaSingloton.FactorDuplicates()
		gongxsd_models.SchemaSingloton.RenameTypeAnonymousComplexType()
	}

	stack.Stage.StageBranchSchema(&gongxsd_models.SchemaSingloton)

	stack.Stage.ComputeReverseMaps()

	gongxsd_models.PostProcessing(stack.Stage)

	if Verbose {
		fmt.Printf("generating file %s\n", *outputModelFilePath)
	}

	// and XSD is a acyclic directed graph (ADG) and it can be interesting
	// to navigate this ADG. Therefore, one first set the interface links between
	// the XSD nodes
	gongxsd_models.SchemaSingloton.SetParentAndChildren(nil)
	gongxsd_models.Generate(stack.Stage, *outputModelFilePath)

	if Verbose {
		fmt.Printf("generate took %s\n", time.Since(start))
	}

	return
}
