/*
Copyright © 2024 AshutoshPatole
*/
package cmd

import (
	"fmt"
	"os"
	"syscall"

	"github.com/TwiN/go-color"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

var aGroup, aUser, aEnv, aAlias string

var allowedEnvironments = []string{"dev", "uat", "sit", "ppd", "prd", "test"}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: func(cmd *cobra.Command, args []string) error {

		// check if a server name is provided
		if len(args) != 1 {
			return fmt.Errorf(color.InRed("Requires hostname or ip\nUsage: ssm add hostname [...options]"))
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {

		// check if a valid env value is passed
		isValid := false
		for _, env := range allowedEnvironments {
			if aEnv == env {
				isValid = true
				break
			}
		}
		if isValid {
			fmt.Println("Valid Environment")
		} else {
			fmt.Print(color.InRed("Unknown environment: Allowed values are "))
			fmt.Println(allowedEnvironments)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&aGroup, "group", "g", "", "Group name in which this server should be added (required)")
	addCmd.MarkFlagRequired("group")
	addCmd.Flags().StringVarP(&aUser, "user", "u", "", "User name to connect (required)")
	addCmd.MarkFlagRequired("user")
	addCmd.Flags().StringVarP(&aEnv, "env", "e", "dev", "Enviornment name to store this server. Allowed values are [prd, ppd, uat, sit, dev]")
	addCmd.Flags().StringVarP(&aAlias, "alias", "a", "", "Alias for the server (required)")
	addCmd.MarkFlagRequired("alias")
}

// Read password from terminal and return a string
func AskPass() string {
	fmt.Print("Password: ")
	pwbyte, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		fmt.Print(color.InRedOverBlack(err))
		os.Exit(1)
	}
	fmt.Println("")
	password := string(pwbyte)
	return password
}
