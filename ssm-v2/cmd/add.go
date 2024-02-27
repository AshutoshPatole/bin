/*
Copyright © 2024 AshutoshPatole
*/
package cmd

import (
	"fmt"
	"os"
	"syscall"
	"time"

	"github.com/TwiN/go-color"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh"
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
			AddServer(args[0])
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

func AddServer(server string) {
	password := AskPass()
	InitServer(server, aGroup, aUser, aEnv, aAlias, password)
}

func InitServer(server, group, user, env, alias, password string) {
	const PORT = "22"

	config := &ssh.ClientConfig{
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		User:    user,
		Timeout: 5 * time.Second,
		HostKeyAlgorithms: []string{
			ssh.KeyAlgoRSA,
			ssh.KeyAlgoDSA,
			ssh.KeyAlgoECDSA256,
			ssh.KeyAlgoECDSA384,
			ssh.KeyAlgoECDSA521,
			ssh.KeyAlgoED25519,
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	client, err := ssh.Dial("tcp", server+":"+PORT, config)
	if err != nil {
		fmt.Println(color.InRed(err.Error()))
		return
	}
	defer client.Close()
	// start session
	session, err := client.NewSession()
	if err != nil {
		fmt.Println(color.InRed(err.Error()))
		return
	}
	defer session.Close()
	// setup standard out and error
	// uses writer interface
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	AddPubKeysToServer(session)

}

func AddPubKeysToServer(session *ssh.Session) bool {
	home, _ := os.UserHomeDir()

	pubKeyPath := home + "/.ssh/id_ed25519.pub"

	pubKey, err := os.ReadFile(pubKeyPath)

	if err != nil {
		fmt.Println(color.InRed("Could not read public key " + pubKeyPath))
		return false
	}

	command := fmt.Sprintf("mkdir -p ~/.ssh/; chmod 700 -R ~/.ssh; echo '%s' >> ~/.ssh/authorized_keys; chmod 600 ~/.ssh/authorized_keys", pubKey)
	if err := session.Run(command); err != nil {
		fmt.Println(color.InRed("Failed to add public key " + err.Error()))
		return false
	} else {
		fmt.Println(color.InGreen("Public keys are added"))
	}

	defer session.Close()
	return true
}
