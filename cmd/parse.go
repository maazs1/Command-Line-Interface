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
	delimiter := args[1]
	if (delimiter==" "){
		s := strings.Split(line, " ")
		fmt.Printf("Parsing of %s is %s", args, s)
	}else{
		s := strings.Split(line, delimiter)
		fmt.Printf("Parsing of %s is %s", args, s)
	}
	
	
}

func init() {
	rootCmd.AddCommand(parseCmd)
}
