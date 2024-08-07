package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gongxsd-cli",
	Short: "gongxsd-cli explores xsd files",
	Long:  `a cli, based on corba and gong`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

var Verbose bool
var unmarshallFromCode *string
var marshallOnCommit *string
var port *int

func init() {
	rootCmd.AddCommand(readCmd)

	readCmd.Flags().StringP("output", "o", "", "Output file")

	unmarshallFromCode = readCmd.Flags().StringP("unmarshallFromCode", "u", "stage.go", "unmarshall data from go file and '.go' (must be lowercased without spaces), If unmarshallFromCode arg is '', no unmarshalling")
	marshallOnCommit = readCmd.Flags().StringP("marshallOnCommit", "m", "stage.go", "on all commits, marshall staged data to a go file with the marshall name and '.go' (must be lowercased without spaces). If marshall arg is '', no marshalling")
	port = readCmd.Flags().IntP("port", "p", 8080, "port server")

	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")

	rootCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		if cmd.Flags().NFlag() == 0 && len(args) == 0 {
			cmd.Help()
		}
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
