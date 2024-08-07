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

		content, err := os.ReadFile(xsdFilePath)
		if err != nil {
			fmt.Printf("Error reading XSD file: %v\n", err)
			os.Exit(1)
		}

		var xsd Schema
		err = xml.Unmarshal(content, &xsd)
		if err != nil {
			fmt.Printf("Error parsing XML: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("XSD File Content:")
		// fmt.Println(string(content))

		r := gongxsd_static.ServeStaticFiles(false)
		stack := gongxsd_stack.NewStack(r, "gongxsd", *unmarshallFromCode, *marshallOnCommit, "", false, true)
		stack.Probe.Refresh()

		log.Printf("Server ready serve on localhost:" + strconv.Itoa(*port))
		err = r.Run(":" + strconv.Itoa(*port))
		if err != nil {
			log.Fatalln(err.Error())
		}
	},
}

type Schema struct {
	Elements     []*Element     `xml:"element"`
	SimpleTypes  []*SimpleType  `xml:"simpleType"`
	ComplexTypes []*ComplexType `xml:"complexType"`
}

type Element struct {
	Name        string       `xml:"name,attr"`
	Type        string       `xml:"type,attr"`
	SimpleType  *SimpleType  `xml:"simpleType"`
	ComplexType *ComplexType `xml:"complexType"`
}

type SimpleType struct {
	Name        string       `xml:"name,attr"`
	Restriction *Restriction `xml:"restriction"`
}

type Restriction struct {
	Base         string         `xml:"base,attr"`
	Enumerations []*Enumeration `xml:"enumeration"`
}

type Enumeration struct {
	Value string `xml:"value,attr"`
}

type ComplexType struct {
	Name     string    `xml:"name,attr"`
	Sequence *Sequence `xml:"sequence"`
}

type Sequence struct {
	Elements []*Element `xml:"element"`
}
