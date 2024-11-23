package cmd

import (
	"beer/core/services"
	"beer/util"
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

			log.Printf("Config path: %v\n", configPath)

			log.Println("Loading apps ...")
			apps := services.GetApps(configPath)

			for _, app := range apps {
				log.Println("-----------------------------------------------------------------")
				log.Println("App name: ", app.Name)
				log.Println("App Process name:", app.Process.Name)
				log.Println("App Process pid:", app.Process.Pid)
				log.Println("App Process cmdline:", app.Process.Cmdline)
				log.Println("App Process status:", app.Process.Status)
				log.Println("App Process connections:", app.Process.Connections)
			}

			fs := http.FileServer(http.Dir("./web/build"))
			http.Handle("/", fs)

			log.Print("Listening on http://localhost:9090/ ...")
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
		log.Fatal(err)
		os.Exit(1)
	}
}
