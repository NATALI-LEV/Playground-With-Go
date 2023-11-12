package main

import (
	"fmt"
	"os"
	"os/exec"
)

// updateCommitPush automates the process of adding all files, committing changes,
// and pushing them to the remote repository.
func updateCommitPush() {
	// Command 1: Add all files recursively to git repo
	addCommand := exec.Command("git", "add", "-A")
	addCommand.Stdout = os.Stdout
	addCommand.Stderr = os.Stderr

	// Run the 'git add' command
	if err := addCommand.Run(); err != nil {
		// Print an error message and exit if the command fails
		fmt.Fprintln(os.Stderr, "Error: Failed to add files to the git repo.")
		os.Exit(1)
	}

	// Command 2: Check if there are changes to commit
	statusCommand := exec.Command("git", "status", "--porcelain")
	statusOutput, err := statusCommand.Output()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: Failed to check git status.")
		os.Exit(1)
	}

	// If there are changes, proceed with commit and push
	if len(statusOutput) > 0 {
		// Command 3: Commit all changes with a predefined message
		commitCommand := exec.Command("git", "commit", "-m", "automated commit")
		commitCommand.Stdout = os.Stdout
		commitCommand.Stderr = os.Stderr

		// Run the 'git commit' command
		if err := commitCommand.Run(); err != nil {
			// Print an error message and exit if the command fails
			fmt.Fprintln(os.Stderr, "Error: Failed to commit changes.")
			os.Exit(1)
		}

		// Command 4: Push to remote (origin main)
		pushCommand := exec.Command("git", "push", "origin", "main")
		pushCommand.Stdout = os.Stdout
		pushCommand.Stderr = os.Stderr

		// Run the 'git push' command
		if err := pushCommand.Run(); err != nil {
			// Print an error message and exit if the command fails
			fmt.Fprintln(os.Stderr, "Error: Failed to push changes to remote.")
			os.Exit(1)
		}

		// Print a success message if all commands are successful
		fmt.Println("Successfully added, committed, and pushed changes!")
	} else {
		// Print a message if there are no changes to commit
		fmt.Println("No changes to commit.")
	}
}

func main() {
	// Call the updateCommitPush function to perform the automated Git workflow
	updateCommitPush()
}
