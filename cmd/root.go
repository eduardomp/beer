package cmd

import (
	"beer/core/services"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "beer",
	Short: "BEER - BackEnd Engineers Rocks",
	Long:  `A tool to help amazing backend developers to visualize their wonderful applications, revealing all the hidden beatiful of the backend layer to the world.`,
	Run: func(cmd *cobra.Command, args []string) {

		services.GetProcessInfo()

		//TODO: Filter processes by mapped apps in the config file
		//TODO: Return filled apps in the list to contruct the graph

		/*
			interfaces, _ := net.Interfaces()

			for _, i := range interfaces {
				fmt.Println(i.Index)
				fmt.Println(i.Name)
				fmt.Println(i.HardwareAddr)
				fmt.Println(i.Flags)
				fmt.Println(i.MTU)
				fmt.Println("--------")
			}*/

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
