
package cmd

import (
	"fmt"
	
	"github.com/spf13/cobra"
)

// ConcatenateStringsCmd represents the ConcatenateStrings command
var ConcatenateStringsCmd = &cobra.Command{
	Use:   "ConcatenateStrings",
	Short: "Concatenates Strings",
	Long: `Concatenates string together seperated by space`,
	Run: func(cmd *cobra.Command, args []string) {
		concatenate(args)
	},
}
func concatenate(args []string){
	var addedString string

	for _, val := range args {
		
       	addedString = addedString + val
	}
	fmt.Printf("Concatenation of %s is %q", args, addedString)
}

func init() {
	rootCmd.AddCommand(ConcatenateStringsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ConcatenateStringsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ConcatenateStringsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
