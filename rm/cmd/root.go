package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/TwiN/go-color"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove files and directories",
	Long:  `Remove files and directories`,

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			fmt.Println(color.InRed("requires atleast one parameter"))
			os.Exit(1)
		}
		recursive, _ := cmd.Flags().GetBool("recursive")
		for _, arg := range args {
			path, err := filepath.Abs(arg)
			if err != nil {
				fmt.Println(color.InRed(err))
				continue
			}
			if err := remove(path, recursive); err != nil {
				fmt.Println(color.InRed(err))
			} else {
				fmt.Println(color.InGreen(path + " removed"))
			}

		}
	},
}

func remove(path string, recursive bool) error {
	fi, err := os.Stat(path)
	if err != nil {
		return err
	}

	if fi.IsDir() {
		if recursive {
			// Only remove directories if recursive flag is set
			return removeDirectory(path)
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
