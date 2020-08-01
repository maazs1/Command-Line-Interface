
package cmd

import (
	"fmt"
	"strconv"
	"github.com/spf13/cobra"
)

// multiplyCmd represents the multiply command
var multiplyCmd = &cobra.Command{
	Use:   "multiply",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		status, _ := cmd.Flags().GetBool("float")
		if status {
			multiplyFloat(args)
		}else{
			multiplyInt(args)
		}
	},
}

func multiplyInt (args []string) {
	var sumVal int = 1

	for _, val := range args {
		temp, err := strconv.Atoi(val)
       
    	if err != nil {
           	fmt.Println(err)
       	}
       	sumVal = sumVal * temp
	}
	fmt.Printf("Multiplication of numbers %s is %d", args, sumVal)
}

func multiplyFloat (args []string) {
	var sum float64 =  1

	for _, val := range args {
		temp, err :=strconv.ParseFloat(val,64)

		if err != nil {
			fmt.Println(err)
		}
		sum = sum * temp
	}
	fmt.Printf("Multiplication of numbers %s is %f", args, sum)
}

func init() {
	rootCmd.AddCommand(multiplyCmd)
	multiplyCmd.Flags().BoolP("float", "f", false, "Multiply Floating Numbers")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// multiplyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// multiplyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
