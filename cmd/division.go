
package cmd

import (
	"fmt"
	"strconv"
	"github.com/spf13/cobra"
)

// divisionCmd represents the division command
var divisionCmd = &cobra.Command{
	Use:   "division",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		status, _ := cmd.Flags().GetBool("float")
		if status {
			divisionFloat(args)
		}else{
			divisionInt(args)
		}
	},
}

func divisionInt (args []string) {
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
			sumVal = sumVal / temp
		}
		
	}
	fmt.Printf("Division of numbers %s is %d", args, sumVal)
}

func divisionFloat (args []string) {
	var sum float64 =  1

	for _, val := range args {
		temp, err :=strconv.ParseFloat(val,64)

		if err != nil {
			fmt.Println(err)
		}
		sum = sum / temp
	}
	fmt.Printf("Division of numbers %s is %f", args, sum)
}

func init() {
	rootCmd.AddCommand(divisionCmd)
	divisionCmd.Flags().BoolP("float", "f", false, "Dividing Floating Numbers")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// divisionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// divisionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
