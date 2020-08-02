
package cmd

import (
	"fmt"
	"strconv"
	"github.com/spf13/cobra"
)

// subtractCmd represents the subtract command
var subtractCmd = &cobra.Command{
	Use:   "subtract",
	Short: "Subtracts float or integers",

	Run: func(cmd *cobra.Command, args []string) {
		status, _ := cmd.Flags().GetBool("float")
		if status {
			subtractFloat(args)
		}else{
			subtractInt(args)
		}
	},
}

func subtractInt (args []string) {
	var sumVal int 
	var flag int =1

	for _, val := range args {
		temp, err := strconv.Atoi(val)
       
    	if err != nil {
           	fmt.Println(err)
		}
		if flag ==1{
			sumVal =temp
			flag=2
		}else{
			sumVal = sumVal - temp
		}
		
	}
	fmt.Printf("Subtraction of numbers %s is %d", args, sumVal)
}

func subtractFloat (args []string) {
	var sumVal float64 
	var flag int =1

	for _, val := range args {
		temp, err :=strconv.ParseFloat(val,64)

		if err != nil {
			fmt.Println(err)
		}
		if flag ==1{
			sumVal =temp
			flag=2
		}else{
			sumVal = sumVal - temp
		}
	}
	fmt.Printf("Subtraction of numbers %s is %f", args, sumVal)
}

func init() {
	rootCmd.AddCommand(subtractCmd)
	subtractCmd.Flags().BoolP("float", "f", false, "Subtracting Floating Numbers")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// subtractCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// subtractCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
