/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/cage1016/wason-translate/lib"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists documents that have been submitted for translation.",
	Run: func(cmd *cobra.Command, args []string) {
		createList(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func createList(cmd *cobra.Command, args []string) {
	logrus.Info("Fetching Documents list...")

	req := lib.ListRequest{
		Version: viper.GetString("version"),
		APIKey:  viper.GetString("api_key"),
		URL:     viper.GetString("url"),
	}

	res, err := lib.List(req)
	if err != nil {
		logrus.Fatalf("Failed to list documents: %s", err)
		return
	}

	doc, err := documentsSelect(res, "Dump Choose Item Detail")
	if err != nil {
		logrus.Fatalf("Error get select document: %s", err)
	}

	logrus.Infof("DocumentID: %s", doc.DocumentID)
	logrus.Infof("Filename: %s", doc.Filename)
	logrus.Infof("Status: %s", doc.Status)
	logrus.Infof("ModelID: %s", doc.ModelID)
	logrus.Infof("Source: %s", doc.Source)
	logrus.Infof("Target: %s", doc.Target)
	logrus.Infof("WordCount: %s", doc.WordCount)
	logrus.Infof("CharacterCount: %s", doc.CharacterCount)
	logrus.Infof("Created: %s", doc.Created)
	logrus.Infof("Completed: %s", doc.Completed)
	logrus.Infof("")

	logrus.Info("Fetching Documents list Done...")
}
