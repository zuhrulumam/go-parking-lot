package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hugo",
	Short: "Hugo for ci cd automation",
	Long:  `ci cd automation`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("begin operation")
	},
}

func init() {
	rootCmd.AddCommand(serverCommand)
	rootCmd.AddCommand(seedCommand)
	rootCmd.AddCommand(cleanerCommand)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
