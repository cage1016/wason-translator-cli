/*
Copyright © 2022 KAI CHU CHUNG cage.chung@gmail.com

*/
package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete translated document.",
	Run: func(cmd *cobra.Command, args []string) {
		createDeletePrompt(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

func createDeletePrompt(cmd *cobra.Command, args []string) {
	res := loadList()

	actions := []action{
		{"Continue Delete", 0},
		{"Reload List", 1},
		{"Quit", 2},
	}

	if len(res) == 0 {
		logrus.Info("No documents found")
	} else {
		for {
			doc, err := documentsSelect2(res, "Select Document you want to Delete")
			if err != nil {
				logrus.Fatalf("Error get select document: %s", err)
			}

			deleteDocument(doc)

			res = filter(res, func(id string) bool {
				return id != doc.DocumentID
			})

			if len(res) == 0 {
				logrus.Info("No documents found")
				break
			}

			i := promptGetActionSelect(actions)

			if actions[i].Value == 0 {
				continue
			} else if actions[i].Value == 1 {
				res = loadList()
				continue
			} else {
				break
			}
		}
	}
}
