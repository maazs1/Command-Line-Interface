package cmd

import (
	"fmt"
	"errors"
	"os"
	"github.com/spf13/cobra"
	"log"
)

// deleteFileCmd represents the deleteFile command
var deleteFileCmd = &cobra.Command{
	Use:   "deleteFile",
	Short: "Deleting selected file",

	Run: func(cmd *cobra.Command, args []string) {
		deletingFile(cmd)
	},
}

func deletingFile(cmd *cobra.Command) error {
	nameofFile, _ := cmd.Flags().GetString("remove")

	if nameofFile == "" {
        return errors.New("File needs a name")
	}
	e := os.Remove(nameofFile)
	if e != nil { 
        log.Fatal(e) 
	} 
	fmt.Println(nameofFile + " has been deleted")
	return nil
}

func init() {
	rootCmd.AddCommand(deleteFileCmd)
	deleteFileCmd.Flags().StringP("remove", "r","", "Name of the file you want to delete")

}
