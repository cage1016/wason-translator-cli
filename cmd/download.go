/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"path/filepath"
	"strings"

	"github.com/cage1016/wason-translate/lib"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Gets the translated document associated with the given document ID.",
	Run: func(cmd *cobra.Command, args []string) {
		createDownload(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)
}

func createDownload(cmd *cobra.Command, args []string) {
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

	// doc, err := documentsSelect([]byte(`{ "documents": [ { "document_id": "e942182b-15c0-4b5e-a4a0-3361a3019833", "filename": "sgc-summary-for-tsmc-1.pptx", "status": "available", "model_id": "en-zh-TW", "source": "en", "target": "zh-TW", "created": "2022-02-19T16:34:29.000Z", "completed": "2022-02-19T16:34:35.000Z", "word_count": 1169, "character_count": 7411 }, { "document_id": "9170b005-b620-4377-bdff-b45ba83df20b", "filename": "sgc-summary-for-tsmc-2.pptx", "status": "available", "model_id": "en-zh-TW", "source": "en", "target": "zh-TW", "created": "2022-02-20T01:43:32.000Z", "completed": "2022-02-20T01:43:37.000Z", "word_count": 696, "character_count": 4228 } ] } `))
	doc, err := documentsSelect(res, "Select Document you want to Download")
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

	req := lib.DownloadRequest{
		Version:        viper.GetString("version"),
		APIKey:         viper.GetString("api_key"),
		URL:            viper.GetString("url"),
		DocumentID:     doc.DocumentID,
		Accept:         lib.AcceptMap[ext],
		OutputFileName: outputFileName,
	}

	lib.Download(req)
}
