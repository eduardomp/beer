package cmd

import (
	"beer/api"
	"beer/util"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	cfgFile  string
	headless bool
	rootCmd  = &cobra.Command{
		Use:   "beer",
		Short: "BEER - BackEnd Engineers Rocks",
		Long:  `A tool to help amazing backend developers to visualize their wonderful applications, revealing all the hidden beatiful of the backend layer to the world.`,
		Run: func(cmd *cobra.Command, args []string) {

			log.Printf("Config path: %v\n", util.GetBeerConfigFile())

			api.RunServer(headless)

		},
	}
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $PWD/config.beer.yml)")
	rootCmd.PersistentFlags().BoolVarP(&headless, "headless", "H", false, "headless mode (just run the REST API)")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
