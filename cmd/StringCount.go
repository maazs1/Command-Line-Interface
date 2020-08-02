
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// StringCountCmd represents the StringCount command
var StringCountCmd = &cobra.Command{
	Use:   "StringCount",
	Short: "Counts the number of indexes in a string",

	Run: func(cmd *cobra.Command, args []string) {
		countString(args)
	},
}

func countString(args []string){
	var count int
	for _, val:=range args {
		count =count +len(val)
	}
	fmt.Printf("The number of items in %s is %d", args, count)
}

func init() {
	rootCmd.AddCommand(StringCountCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// StringCountCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// StringCountCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
