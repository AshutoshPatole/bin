package cmd

import (
	"fmt"
	"os"
	"runtime"

	"github.com/spf13/cobra"
	"github.com/zcalusic/sysinfo"
)

var platform, architecture string

var rootCmd = &cobra.Command{
	Use:   "saygo",
	Short: "Initator cli for linux machines",
	Run: func(cmd *cobra.Command, args []string) {
		platform = runtime.GOOS
		architecture = runtime.GOARCH
		fmt.Println(platform)
		fmt.Println(architecture)

		var si sysinfo.BIOS

		// si.GetSysInfo()
		fmt.Println(si.Vendor)

		// data, err := json.Marshal(&si)
		// if err != nil {
		// 	log.Fatal(err)
		// }

		// fmt.Println(string(data))

	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

}
