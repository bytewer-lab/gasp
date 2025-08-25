/*
Copyright © 2025 Lucas Soares <lucas@bytewer.com>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a hook file",
	Long:  `This command will run all the commands in the hook file.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Error: hook filename required")
			os.Exit(1)
		}
		hookFilename := args[0]
		rootDir, err := FindGitRoot()
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			os.Exit(1)
		}
		hookPath := filepath.Join(rootDir, "hooks", hookFilename)

		if _, err := os.Stat(hookPath); os.IsNotExist(err) {
			os.Exit(0)
		}

		fmt.Printf("Executing hook file %s\n", hookFilename)
		var execCmd *exec.Cmd
		if runtime.GOOS == "windows" {
			execCmd = exec.Command("C:\\Program Files\\Git\\usr\\bin\\bash.exe", hookPath)
		} else {
			execCmd = exec.Command("/bin/sh", hookPath)
		}
		execCmd.Dir = rootDir
		execCmd.Stdout = os.Stdout
		execCmd.Stderr = os.Stderr
		err = execCmd.Run()
		if err != nil {
			fmt.Printf("Hook execution failed: %s\n", err)
			os.Exit(1)
		}
		fmt.Println("Hook execution completed successfully")
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
