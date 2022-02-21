/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/cage1016/document-translator-cli/lib"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a document.",
	Run: func(cmd *cobra.Command, args []string) {
		createDeletePrompt(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

func createDeletePrompt(cmd *cobra.Command, args []string) {
	logrus.Info("Fetching Documents list...")

	res, err := lib.List(lib.ListRequest{
		Version: viper.GetString("version"),
		APIKey:  viper.GetString("api_key"),
		URL:     viper.GetString("url"),
	})
	if err != nil {
		logrus.Errorf("Error Fetching documents: %s", err)
		return
	}

	doc, err := documentsSelect(res, "Select Document you want to Delete")
	if err != nil {
		logrus.Fatalf("Error get select document: %s", err)
	}

	req := &lib.DeleteRequest{
		Version:    viper.GetString("version"),
		APIKey:     viper.GetString("api_key"),
		URL:        viper.GetString("url"),
		DocumentID: doc.DocumentID,
	}

	lib.Delete(req)
}
