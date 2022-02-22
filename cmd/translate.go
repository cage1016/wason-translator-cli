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

// translateCmd represents the translate command
var translateCmd = &cobra.Command{
	Use:   "translate",
	Short: "Submit a document for translation.",
	Long: `You can submit the document contents in the file parameter, or you can reference a previously submitted document by document ID. The maximum file size for document translation is:

- 20 MB for service instances on the Standard, Advanced, and Premium plans
- 2 MB for service instances on the Lite plan..`,
	Run: func(cmd *cobra.Command, args []string) {
		createTranslatePrompt(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(translateCmd)
}

func createTranslatePrompt(cmd *cobra.Command, args []string) {
	actions := []action{
		{"Continue Translate", 0},
		{"Quit", 1},
	}

	for {
		translate()

		i := promptGetActionSelect(actions)

		if actions[i].Value == 0 {
			continue
		} else {
			break
		}
	}

}

func translate() {
	pc := promptContent{
		errorMsg: "You must provide the filename",
		label:    "File Name",
	}

	filename := promptGetInput(pc, "")
	if filename == "" {
		return
	}

	ext := strings.ToLower(filepath.Ext(filename))
	if _, ok := lib.AcceptMap[ext]; !ok {
		logrus.Fatalf("Error: %s is not supported content type", ext)
	}

	pc = promptContent{
		errorMsg: "You must provide the Source Language",
		label:    "Source Language",
	}

	source := promptGetSelect(pc, []string{"en", "zh", "ja"})
	if source == "" {
		logrus.Error("You must provide the Source Language")
		return
	}

	pc = promptContent{
		errorMsg: "You must provide the Source Language",
		label:    "Target Language",
	}

	target := promptGetSelect(pc, []string{"zh-TW"})
	if source == "" {
		logrus.Error("You must provide the target Language")
		return
	}

	// DocumentID, _ := cmd.Flags().GetString("documentId")
	// if DocumentID == "" {
	// }

	lib.TranslateDocument(lib.TranslateRequest{
		Version:  viper.GetString("version"),
		APIKey:   viper.GetString("api_key"),
		URL:      viper.GetString("url"),
		FileName: filename,
		Accept:   lib.AcceptMap[ext],
		Source:   source,
		Target:   target,
	})
}
