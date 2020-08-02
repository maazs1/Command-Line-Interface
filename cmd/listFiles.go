package cmd

import (
	"fmt"
	"log"
	"os"
	"github.com/spf13/cobra"
)

// listFilesCmd represents the listFiles command
var listFilesCmd = &cobra.Command{
	Use:   "listFiles",
	Short: "Listing Files in a Folder",

	Run: func(cmd *cobra.Command, args []string) {
		listFiles()
	},
}
func listFiles(){
	dirname := "."

    f, err := os.Open(dirname)
    if err != nil {
        log.Fatal(err)
    }
    files, err := f.Readdir(-1)
    f.Close()
    if err != nil {
        log.Fatal(err)
    }

    for _, file := range files {
        fmt.Println(file.Name())
    }
}
func init() {
	rootCmd.AddCommand(listFilesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listFilesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listFilesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
