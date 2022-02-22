package lib

import (
	"encoding/json"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/sirupsen/logrus"
	"github.com/watson-developer-cloud/go-sdk/languagetranslatorv3"
)

var AcceptMap = map[string]string{
	// "": "application/powerpoint",
	// "": "application/mspowerpoint",
	// "": "application/x-rtf",
	// "": "application/json",
	// "": "application/xml",
	// "": "application/vnd.ms-excel",
	".xlsx": "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
	// "": "application/vnd.ms-powerpoint",
	".pptx": "application/vnd.openxmlformats-officedocument.presentationml.presentation",
	// "": "application/msword",
	".docx": "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
	// "": "application/vnd.oasis.opendocument.spreadsheet",
	// "": "application/vnd.oasis.opendocument.presentation",
	// "": "application/vnd.oasis.opendocument.text",
	".pdf": "application/pdf",
	// "": "application/rtf",
	".html": "text/html",
	// "": "text/json",
	// "": "text/plain",
	// "": "text/richtext",
	// "": "text/rtf",
	// "": "text/sbv",
	// "": "text/srt",
	// "": "text/vtt",
	// "": "text/xml",
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
