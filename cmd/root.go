package cmd

import (
	"beer/core/services"
	"beer/util"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "beer",
		Short: "BEER - BackEnd Engineers Rocks",
		Long:  `A tool to help amazing backend developers to visualize their wonderful applications, revealing all the hidden beatiful of the backend layer to the world.`,
		Run: func(cmd *cobra.Command, args []string) {

			fmt.Println("Loading apps ...")

			configPath := util.Config("BEER_CONFIG_PATH", "../../config.beer.yml")

			if cfgFile != "" {
				configPath = cfgFile
			}

			fmt.Printf("Config path: %v\n", configPath)

			apps := services.GetApps(configPath)

			for _, app := range apps {
				fmt.Println("-----------------------------------------------------------------")
				fmt.Println("App name: ", app.Name)
				fmt.Println("App Process name:", app.Process.Name)
				fmt.Println("App Process pid:", app.Process.Pid)
				fmt.Println("App Process cmdline:", app.Process.Cmdline)
				fmt.Println("App Process status:", app.Process.Status)
				fmt.Println("App Process connections:", app.Process.Connections)
			}

		},
	}
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $PWD/config.beer.yml)")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
