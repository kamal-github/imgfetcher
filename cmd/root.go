package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "imgfetcher",
	Short: "Generate printable image cards for kids",
	Long:  "imgfetcher downloads free images, labels them, and organizes them into printable cards",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
