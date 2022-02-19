/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// translateCmd represents the translate command
var translateCmd = &cobra.Command{
	Use:   "translate",
	Short: "Submit a document for translation.",
	Long: `You can submit the document contents in the file parameter, or you can reference a previously submitted document by document ID. The maximum file size for document translation is:

- 20 MB for service instances on the Standard, Advanced, and Premium plans
- 2 MB for service instances on the Lite plan..`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("translate called")
	},
}

func init() {
	rootCmd.AddCommand(translateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// translateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// translateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
