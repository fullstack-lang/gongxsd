/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/xml"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"

	"github.com/fullstack-lang/gongxsd/go/models"
	gongxsd_stack "github.com/fullstack-lang/gongxsd/go/stack"
	gongxsd_static "github.com/fullstack-lang/gongxsd/go/static"
)

// generateCmd represents the read command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "read an xsd file and generate the gong models",
	Args:  cobra.ExactArgs(1),
	Long:  `read an xsd file and generate the gong models`,
	Run: func(cmd *cobra.Command, args []string) {

		start := time.Now()
		if Verbose {
			fmt.Printf("generate start\n")
		}

		xsdFilePath := args[0]
		if Verbose {
			fmt.Printf("Reading file: %s\n", xsdFilePath)
		}

		r := gongxsd_static.ServeStaticFiles(false)
		stack := gongxsd_stack.NewStack(r, "gongxsd", *unmarshallFromCode, *marshallOnCommit, "", false, true)
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

		stack.Stage.StageBranchSchema(&models.SchemaSingloton)

		stack.Stage.ComputeReverseMaps()

		models.PostProcessing(stack.Stage)

		if Verbose {
			fmt.Printf("generating file %s\n", *outputModelFilePath)
		}

		models.Generate(stack.Stage, *outputModelFilePath)

		if Verbose {
			fmt.Printf("generate took %s\n", time.Since(start))
		}

	},
}
