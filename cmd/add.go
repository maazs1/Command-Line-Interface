
package cmd

import (
	"fmt"
	"strconv"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds float or integers",
	Run: func(cmd *cobra.Command, args []string) {
		status, _ := cmd.Flags().GetBool("float")
		if status {
			addFloat(args)
		}else{
			addInt(args)
		}
		
	},
}

func addInt (args []string) {
	var sumVal int

	for _, val := range args {
		temp, err := strconv.Atoi(val)
       
    	if err != nil {
           	fmt.Println(err)
       	}
       	sumVal = sumVal + temp
	}
	fmt.Printf("Addition of numbers %s is %d", args, sumVal)
}

func addFloat (args []string) {
	var sum float64

	for _, val := range args {
		temp, err :=strconv.ParseFloat(val,64)

		if err != nil {
			fmt.Println(err)
		}
		sum = sum +temp
	}
	fmt.Printf("Addition of numbers %s is %f", args, sum)
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().BoolP("float", "f", false, "Add Floating Numbers")
}
