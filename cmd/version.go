package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.PersistentFlags().BoolVarP(&VersionFlag, "version", "v", false, "BEER version")
}

var VersionFlag bool
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "BEER version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🍺 BEER \"BackEnd Engineers Rocks\" \nv1.0.0")
	},
}
