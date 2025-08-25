/*
Copyright © 2025 Lucas Soares <lucas@bytewer.com>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove the git hooks",
	Long:  `This command will remove the git hooks from the .git/hooks directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		rootDir, err := FindGitRoot()
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			os.Exit(1)
		}
		gitHooksDir := filepath.Join(rootDir, ".git", "hooks")
		err = CleanTargetDirectory(gitHooksDir)
		if err != nil {
			fmt.Printf("Error while cleaning target hooks directory: %s\n", err)
			os.Exit(1)
		}
		fmt.Printf("Git hooks removed successfully\n")
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// removeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// removeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
