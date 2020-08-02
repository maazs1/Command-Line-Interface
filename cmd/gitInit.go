package cmd

import (
	"fmt"
	"os/exec"
	"github.com/spf13/cobra"
)

// gitInitCmd represents the gitInit command
var gitInitCmd = &cobra.Command{
	Use:   "gitInit",
	Short: "Initializes git repository on a given path",

	Run: func(cmd *cobra.Command, args []string) {
		initializeGit(cmd)
	},
}

func initializeGit (cmd *cobra.Command){
	directoryFolder, _ := cmd.Flags().GetString("directory")


    command := exec.Command("git", "init", directoryFolder)
    err := command.Run()
    if err != nil {
        fmt.Printf("Could not create the directory %v", err)
    }
    fmt.Println("Git repository initialized in " + directoryFolder)

}

func init() {
	rootCmd.AddCommand(gitInitCmd)
	gitInitCmd.Flags().StringP("directory", "d", "", "Directory for git repository initialization")

}
