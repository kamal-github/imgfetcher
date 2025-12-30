package cmd

import (
	"github.com/kamal-github/imgfetcher/internal"
	"github.com/spf13/cobra"
)

var (
	inputFile     string
	outputDir     string
	imagesPerItem int
	workers       int
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate image cards from YAML input",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := internal.LoadYAML(inputFile)
		if err != nil {
			return err
		}
		return internal.GenerateCards(cfg, outputDir, imagesPerItem, workers)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().StringVarP(&inputFile, "input", "i", "", "YAML input file")
	generateCmd.Flags().StringVarP(&outputDir, "out", "o", "images", "Output directory")
	generateCmd.Flags().IntVar(&imagesPerItem, "images-per-item", 2, "Images per item")
	generateCmd.Flags().IntVar(&workers, "workers", 5, "Concurrent workers")

	_ = generateCmd.MarkFlagRequired("input")
}
