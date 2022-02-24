package lib

import (
	"encoding/json"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/sirupsen/logrus"
	"github.com/watson-developer-cloud/go-sdk/languagetranslatorv3"
)

var AcceptMap = map[string]string{
	".ppt":  "application/powerpoint",
	".json": "application/json",
	".xml":  "application/xml",
	".xls":  "application/vnd.ms-excel",
	".xlsx": "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
	".pptx": "application/vnd.openxmlformats-officedocument.presentationml.presentation",
	".doc":  "application/msword",
	".docx": "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
	".ods":  "application/vnd.oasis.opendocument.spreadsheet",
	".odp":  "application/vnd.oasis.opendocument.presentation",
	".odt":  "application/vnd.oasis.opendocument.text",
	".pdf":  "application/pdf",
	".rtf":  "application/rtf",
	".html": "text/html",
	".htm":  "text/html",
	".txt":  "text/plain",
}

type ListRequest struct {
	Version string
	APIKey  string
	URL     string
}

func ListDocument(req ListRequest) ([]byte, error) {

	authenticator := &core.IamAuthenticator{
		ApiKey: req.APIKey,
	}

	options := &languagetranslatorv3.LanguageTranslatorV3Options{
		Version:       req.Version,
		Authenticator: authenticator,
	}

	languageTranslator, languageTranslatorErr := languagetranslatorv3.
		NewLanguageTranslatorV3(options)

	if languageTranslatorErr != nil {
		logrus.Errorf("Error creating Language Translator service: %s", languageTranslatorErr)
		return []byte{}, languageTranslatorErr
	}

	languageTranslator.SetServiceURL(req.URL)

	result, _, responseErr := languageTranslator.ListDocuments(
		&languagetranslatorv3.ListDocumentsOptions{},
	)
	if responseErr != nil {
		logrus.Errorf("Error getting documents: %s", responseErr)
		return []byte{}, responseErr
	}
	b, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		logrus.Errorf("Error marshaling documents: %s", err)
		return []byte{}, err
	}
	return b, nil
}
