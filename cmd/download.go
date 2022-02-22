/*
Copyright © 2022 KAI CHU CHUNG cage.chung@gmail.com

*/
package cmd

import (
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cage1016/document-translator-cli/lib"
)

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download translated document.",
	Run: func(cmd *cobra.Command, args []string) {
		createDownloadPrompt(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)
}

func createDownloadPrompt(cmd *cobra.Command, args []string) {
	res := loadList()

	actions := []action{
		{"Continue Download", 0},
		{"Reload List", 1},
		{"Quit", 2},
	}

	if len(res) == 0 {
		logrus.Info("No documents found")
	} else {
		for {
			doc, err := documentsSelect2(res, "Select Document you want to Download")
			if err != nil {
				logrus.Fatalf("Error get select document: %s", err)
			}

			ext := strings.ToLower(filepath.Ext(doc.Filename))
			if _, ok := lib.AcceptMap[ext]; !ok {
				logrus.Fatalf("Error: %s is not supported", ext)
			}

			pc := promptContent{
				errorMsg: "You must provide the output filename",
				label:    "Output fileName",
			}

			outputFileName := promptGetInput(pc, doc.Filename)
			if outputFileName == "" {
				return
			}

			download(doc, ext, outputFileName)

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

func download(doc *lib.Document, ext string, outputFileName string) {
	req := lib.DownloadRequest{
		Version:        viper.GetString("version"),
		APIKey:         viper.GetString("api_key"),
		URL:            viper.GetString("url"),
		DocumentID:     doc.DocumentID,
		Accept:         lib.AcceptMap[ext],
		OutputFileName: outputFileName,
	}

	lib.DownloadDocument(req)
}
