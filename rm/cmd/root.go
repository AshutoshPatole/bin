package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove files and directories",
	Long:  `Remove files and directories`,

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			displayUsage()
			return
		}
		recursive, _ := cmd.Flags().GetBool("recursive")
		for _, arg := range args {
			path, err := filepath.Abs(arg)
			if err != nil {
				fmt.Printf("Error: Could not resolve path for '%s': %v\n", arg, err)
				continue
			}

			// Check if file exists before attempting removal
			if _, err := os.Stat(path); os.IsNotExist(err) {
				fmt.Printf("Error: '%s' does not exist\n", path)
				continue
			}

			if err := remove(path, recursive); err != nil {
				fmt.Printf("Error: Failed to remove '%s': %v\n", path, err)
			} else {
				fmt.Printf("Successfully removed '%s'\n", path)
			}
		}
	},
}

func displayUsage() {
	fmt.Println("Usage: rm [OPTIONS] FILE...")
	fmt.Println("\nOptions:")
	fmt.Println("  -r, --recursive    Remove directories and their contents recursively")
}

func remove(path string, recursive bool) error {
	fi, err := os.Stat(path)
	if err != nil {
		return err
	}

	if fi.IsDir() {
		if recursive {
			// Remove all contents within directory
			entries, err := os.ReadDir(path)
			if err != nil {
				return err
			}

			for _, entry := range entries {
				fullPath := filepath.Join(path, entry.Name())
				if err := remove(fullPath, recursive); err != nil {
					return err
				}
			}
			return os.Remove(path)
		} else {
			return fmt.Errorf("%s is a directory (use -r to remove directories recursively)", path)
		}
	}

	return removeSingleFile(path)
}

func removeSingleFile(path string) error {
	return os.Remove(path)
}

func removeDirectory(path string) error {
	return os.RemoveAll(path)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("recursive", "r", false, "Remove files recursively")
}