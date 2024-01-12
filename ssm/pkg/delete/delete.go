package delete

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	c "github.com/AshutoshPatole/ssh-manager/utils"
	"github.com/TwiN/go-color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var server string

var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete server from configuration",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		deleteServer(server)
	},
}

func init() {

	DeleteCmd.Flags().StringVarP(&server, "server", "s", "", "Server name to delete")
	DeleteCmd.MarkFlagRequired("server")
}

func deleteServer(serverName string) {
	var config c.Config

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalln(err)
	}
	found := 0
	for _, grp := range config.Groups {
		for _, env := range grp.Environment {
			for i, server := range env.Servers {
				if server.HostName == serverName {
					found = 1
					fmt.Println(color.InBlackOverRed(server.HostName + " found in " + env.Name + " in " + grp.Name))
					reader := bufio.NewReader(os.Stdin)
					fmt.Print("Are you sure you want to delete this server? (y/n): ")
					response, err := reader.ReadString('\n')
					if err != nil {
						log.Fatalln(err)
					}

					if strings.TrimSpace(response) == "y" || strings.TrimSpace(response) == "yes" {

						env.Servers = append(env.Servers[:i], env.Servers[i+1:]...)
						if err := viper.WriteConfig(); err != nil { // Write updated config to file
							log.Fatalln(err)
						}
						fmt.Println(color.InGreen("Server deleted successfully!"))
						break // Exit the inner loop since the server is deleted
					} else {
						fmt.Println(color.InYellow("Server deletion aborted."))
					}

				}
			}

		}
	}
	if found == 1 {
		// Save the information in the config file
		viper.Set("groups", config.Groups)
		if err := viper.WriteConfig(); err != nil {
			log.Fatalln(err)
		}
	} else if found == 0 {
		fmt.Println(color.InRed("Server " + serverName + " was not found in configuration"))
	}

}
