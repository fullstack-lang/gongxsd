package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/fullstack-lang/gongxsd/test/reqif/go/processing"
	reqif_stack "github.com/fullstack-lang/gongxsd/test/reqif/go/stack"

	split_static "github.com/fullstack-lang/gong/lib/split/go/static"

	models "github.com/fullstack-lang/gongxsd/test/reqif/go/models"
)

var (
	logGINFlag = flag.Bool("logGIN", false, "log mode for gin")

	unmarshallFromCode = flag.String("unmarshallFromCode", "", "unmarshall data from go file and '.go' (must be lowercased without spaces), If unmarshallFromCode arg is '', no unmarshalling")
	marshallOnCommit   = flag.String("marshallOnCommit", "", "on all commits, marshall staged data to a go file with the marshall name and '.go' (must be lowercased without spaces). If marshall arg is '', no marshalling")

	diagrams         = flag.Bool("diagrams", true, "parse/analysis go/models and go/diagrams")
	embeddedDiagrams = flag.Bool("embeddedDiagrams", false, "parse/analysis go/models and go/embeddedDiagrams")

	port = flag.Int("port", 8080, "port server")

	reqifFile = flag.String("reqifFile", "", "")
)

func main() {

	log.SetPrefix("reqif: ")
	log.SetFlags(0)

	// parse program arguments
	flag.Parse()

	// setup the static file server and get the controller
	r := split_static.ServeStaticFiles(*logGINFlag)

	// setup stack
	stack := reqif_stack.NewStack(r, "reqif", *unmarshallFromCode, *marshallOnCommit, "", *embeddedDiagrams, true)
	stack.Probe.Refresh()

	// Open the XML file
	xmlFile, err := os.Open(*reqifFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer xmlFile.Close()

	// Read the XML file
	byteValue, err := io.ReadAll(xmlFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Unmarshal the XML into the Reqif struct
	var req_if models.REQ_IF
	err = xml.Unmarshal(byteValue, &req_if)
	if err != nil {
		fmt.Println("Error unmarshalling XML:", err)
		return
	}

	stack.Stage.StageBranchREQ_IF(&req_if)

	processing.PostProcessing(stack.Stage)

	stack.Stage.Commit()
	stack.Probe.Refresh()

	log.Printf("Server ready serve on localhost:" + strconv.Itoa(*port))
	err = r.Run(":" + strconv.Itoa(*port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
