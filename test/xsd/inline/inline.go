package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type InlineStruct struct {
	FieldA string `xml:"FieldA"`
	FieldB string `xml:"FieldB"`
}

type Root struct {
	InlineStruct `xml:",inline"`
}

func main() {
	// Create an instance of Root with some data
	root := new(Root)
	root.FieldA = "A"
	root.FieldB = "B"

	// Marshal the Root struct to XML
	output, err := xml.MarshalIndent(root, "", "  ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	// Print the XML output
	fmt.Println(string(output))

	// Optionally, save the XML to a file
	file, err := os.Create("output.xml")
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	defer file.Close()

	file.WriteString(xml.Header)
	file.Write(output)
}
