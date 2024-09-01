package cmd

import (
	"encoding/xml"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/fullstack-lang/gongxsd/go/models"
	gongxsd_stack "github.com/fullstack-lang/gongxsd/go/stack"
	gongxsd_static "github.com/fullstack-lang/gongxsd/go/static"

	"github.com/gin-gonic/gin"
)

func process(args []string) (r *gin.Engine, stack *gongxsd_stack.Stack) {
	start := time.Now()
	if Verbose {
		fmt.Printf("generate start\n")
	}

	xsdFilePath := args[0]
	if Verbose {
		fmt.Printf("Reading file: %s\n", xsdFilePath)
	}

	r = gongxsd_static.ServeStaticFiles(false)
	stack = gongxsd_stack.NewStack(r, "gongxsd", *unmarshallFromCode, *marshallOnCommit, "", false, true)
	stack.Stage.Reset()

	content, err := os.ReadFile(xsdFilePath)
	if err != nil {
		fmt.Printf("Error reading XSD file: %v\n", err)
		os.Exit(1)
	}

	err = xml.Unmarshal(content, &models.SchemaSingloton)
	if err != nil {
		fmt.Printf("Error parsing XML: %v\n", err)
		os.Exit(1)
	}

	// suppress duplicates
	if strings.Contains(xsdFilePath, "dtc-11-04-05") {
		models.SchemaSingloton.FactorDuplicates()
		models.SchemaSingloton.RenameTypeAnonymousComplexType()
	}

	stack.Stage.StageBranchSchema(&models.SchemaSingloton)

	stack.Stage.ComputeReverseMaps()

	models.PostProcessing(stack.Stage)

	if Verbose {
		fmt.Printf("generating file %s\n", *outputModelFilePath)
	}

	// and XSD is a acyclic directed graph (ADG) and it can be interesting
	// to navigate this ADG. Therefore, one first set the interface links between
	// the XSD nodes
	models.SchemaSingloton.SetParentAndChildren(nil)
	models.Generate(stack.Stage, *outputModelFilePath)

	if Verbose {
		fmt.Printf("generate took %s\n", time.Since(start))
	}

	return
}
