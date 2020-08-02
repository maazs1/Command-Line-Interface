package cmd

import (
	"fmt"
	"errors"
	"os"
	"github.com/spf13/cobra"
)

// createFolderCmd represents the createFolder command
var createFolderCmd = &cobra.Command{
	Use:   "createFolder",
	Short: "Creates a folder with a name as parameter`",
	Run: func(cmd *cobra.Command, args []string) {
		createFolder(cmd, args)
	},
}

func createFolder(cmd *cobra.Command, args []string) error {
	nameofFile, _ := cmd.Flags().GetString("name")

    if nameofFile == "" {
        return errors.New("Folder needs a name")
    }

    err := os.MkdirAll(nameofFile, os.ModePerm)
    if err != nil {
        fmt.Printf("Could not create the directory %v", err)
    }
    fmt.Println("Folder " + nameofFile + " created")

    return nil
}

func init() {
	rootCmd.AddCommand(createFolderCmd)
	createFolderCmd.Flags().StringP("name", "n","", "Name of the folder you want to create.")
}
