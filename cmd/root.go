package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gocutter",
	Short: "A CLI for rapidly scaffolding Go projects with templates or by cloning existing structures.",
	Run: func(cmd *cobra.Command, args []string) {
		destination, err := cmd.Flags().GetString("destination")
		if err != nil {
			fmt.Println("Error retrieving destination flag:", err)
			os.Exit(1)
		}
		if err := cutter(destination); err != nil {
			fmt.Println("Error running cutter:", err)
			os.Exit(1)
		}
		fmt.Println("Successfully created Go project at", destination)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("destination", "d", "", "Destination directory for the new project. For example: ./your/project/path")
	// 标记 destination 为必填
	if err := rootCmd.MarkFlagRequired("destination"); err != nil {
		fmt.Println("Error marking 'destination' flag as required:", err)
		os.Exit(1)
	}

}
