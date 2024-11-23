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

		fmt.Println("Loading apps ...")

		apps := services.GetApps()

		for _, app := range apps {
			fmt.Println("-----------------------------------------------------------------")
			fmt.Println("App name: ", app.Name)
			fmt.Println("App Process name:", app.Process.Name)
			fmt.Println("App Process pid:", app.Process.Pid)
			fmt.Println("App Process cmdline:", app.Process.Cmdline)
			fmt.Println("App Process status:", app.Process.Status)
			fmt.Println("App Process connections:", app.Process.Connections)
		}

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
