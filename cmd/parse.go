
package cmd

import (
	"fmt"
	"strings"
	"github.com/spf13/cobra"
)

// parseCmd represents the parse command
var parseCmd = &cobra.Command{
	Use:   "parse",
	Short: "Parses a given string with a delimeter",

	Run: func(cmd *cobra.Command, args []string) {
		parseString(args)
	},
}

func parseString(args []string){
	line := args[0]
	delimeter := args[1]
	s := strings.Split(line, delimeter)
	fmt.Printf("Parsing of %s is %s", args, s)
}

func init() {
	rootCmd.AddCommand(parseCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// parseCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// parseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
