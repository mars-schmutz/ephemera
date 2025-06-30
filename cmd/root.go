package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:  "ephemera",
	Long: `A tool to extract/download archives/git repositories, and then easily clean up afterwards`,
	Run:  DoEphemera,
}

var archivePath string
var repoPath string

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&archivePath, "archive", "a", "", "path to the archive file (.zip, .tar.gz, etc.)")
	rootCmd.Flags().StringVarP(&repoPath, "repository", "r", "", "link to GitHub repository (https or ssh)")
}

func DoEphemera(cmd *cobra.Command, args []string) {
	if archivePath != "" && repoPath != "" {
		fmt.Println("You should use --archive or --repository, but not both.")
		fmt.Println()
		cmd.Help()
	}

	// TODO: Fix no flags given
	// Should it create empty temp shell or error out?
	if archivePath == "" && repoPath == "" {
		fmt.Println("You didn't supply a flag.")
		fmt.Println()
		cmd.Help()
		os.Exit(1)
	}

	temp, err := os.MkdirTemp("", "ephemera_")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(temp)

	if repoPath != "" {
		err := cloneRepo(temp, repoPath)
		if err != nil {
			os.RemoveAll(temp)
			os.Exit(1)
		}
	}

	if archivePath != "" {
		Unarchive(temp, archivePath)
	}

	shell := exec.Command("zsh")
	shell.Dir = temp
	shell.Stdin = os.Stdin
	shell.Stderr = os.Stderr
	shell.Stdout = os.Stdout

	fmt.Println("Entering temporary shell. Tpye 'exit' when done.")
	if err := shell.Run(); err != nil {
		fmt.Println("Error running shell: ", err)
		return
	}

	fmt.Println("Exited shell. Cleaning up...")
}

func cloneRepo(dir, repo string) error {
	clone := exec.Command("git", "clone", repo, dir)
	clone.Stderr = os.Stderr
	clone.Stdout = os.Stdout

	if err := clone.Run(); err != nil {
		fmt.Printf("Clone failed: %s\n", err)
		return err
	}

	return nil
}
