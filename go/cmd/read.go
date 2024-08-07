/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"

	gongxsd_models "github.com/fullstack-lang/gongxsd/go/models"
	gongxsd_stack "github.com/fullstack-lang/gongxsd/go/stack"
	gongxsd_static "github.com/fullstack-lang/gongxsd/go/static"
)

// readCmd represents the read command
var readCmd = &cobra.Command{
	Use:   "read",
	Short: "read an xsd file",
	Args:  cobra.ExactArgs(1),
	Long:  `...`,
	Run: func(cmd *cobra.Command, args []string) {

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

		var schema gongxsd_models.Schema
		err = xml.Unmarshal(content, &schema)
		if err != nil {
			fmt.Printf("Error parsing XML: %v\n", err)
			os.Exit(1)
		}

		stack.Stage.StageBranchSchema(&schema)

		stack.Stage.Commit()
		fmt.Println("XSD File Content:")
		// fmt.Println(string(content))

		stack.Probe.Refresh()

		log.Printf("Server ready serve on localhost:" + strconv.Itoa(*port))
		err = r.Run(":" + strconv.Itoa(*port))
		if err != nil {
			log.Fatalln(err.Error())
		}
	},
}
