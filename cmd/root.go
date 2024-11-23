package cmd

import (
	"beer/core/services"
	"beer/util"
	"fmt"
	"log"
	"net/http"
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

			configPath := util.Config("BEER_CONFIG_PATH", "../../config.beer.yml")

			if cfgFile != "" {
				configPath = cfgFile
			}

			fmt.Printf("Config path: %v\n", configPath)

			fmt.Println("Loading apps ...")
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

			fs := http.FileServer(http.Dir("./web/build"))
			http.Handle("/", fs)

			fmt.Print("\n\nListening on http://localhost:9090 ...\n")
			err := http.ListenAndServe(":9090", nil)
			if err != nil {
				log.Fatal(err)
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
