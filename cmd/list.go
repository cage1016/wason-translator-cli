/*
Copyright © 2022 KAI CHU CHUNG cage.chung@gmail.com

*/
package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists documents that have been submitted for translation.",
	Run: func(cmd *cobra.Command, args []string) {
		createListPrompt(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func createListPrompt(cmd *cobra.Command, args []string) {
	res := loadList()

	doc, err := documentsSelect2(res, "Dump Choose Item Detail")
	if err != nil {
		logrus.Fatalf("Error get select document: %s", err)
	}

	logrus.Infof("DocumentID: %s", doc.DocumentID)
	logrus.Infof("Filename: %s", doc.Filename)
	logrus.Infof("Status: %s", doc.Status)
	logrus.Infof("ModelID: %s", doc.ModelID)
	logrus.Infof("Source: %s", doc.Source)
	logrus.Infof("Target: %s", doc.Target)
	logrus.Infof("WordCount: %d", doc.WordCount)
	logrus.Infof("CharacterCount: %d", doc.CharacterCount)
	logrus.Infof("Created: %s", doc.Created)
	logrus.Infof("Completed: %s", doc.Completed)
	logrus.Infof("")

	logrus.Info("Fetching Documents list Done...")
}
