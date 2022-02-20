package lib

import (
	"encoding/json"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/sirupsen/logrus"
	"github.com/watson-developer-cloud/go-sdk/v2/languagetranslatorv3"
)

type TranslateRequest struct {
	Version  string
	APIKey   string
	URL      string
	FileName string
	Accept   string
	// ModelID  string
	Source string
	Target string
}

func Translate(req TranslateRequest) {
	logrus.Infof("Translate file %s uploading", req.FileName)

	authenticator := &core.IamAuthenticator{
		ApiKey: req.APIKey,
	}

	options := &languagetranslatorv3.LanguageTranslatorV3Options{
		Version:       &req.Version,
		Authenticator: authenticator,
	}

	languageTranslator, languageTranslatorErr := languagetranslatorv3.
		NewLanguageTranslatorV3(options)

	if languageTranslatorErr != nil {
		logrus.Fatal(languageTranslatorErr)
	}

	languageTranslator.SetServiceURL(req.URL)

	file, fileErr := os.Open(req.FileName)
	if fileErr != nil {
		logrus.Fatal(fileErr)
	}
	defer file.Close()

	result, _, responseErr := languageTranslator.TranslateDocument(
		&languagetranslatorv3.TranslateDocumentOptions{
			File:            file,
			FileContentType: core.StringPtr(req.Accept),
			Filename:        core.StringPtr(req.FileName),
			// ModelID:         core.StringPtr(req.ModelID),
			Source: core.StringPtr(req.Source),
			Target: core.StringPtr(req.Target),
		},
	)
	if responseErr != nil {
		logrus.Fatal(responseErr)
	}
	b, _ := json.MarshalIndent(result, "", "  ")
	logrus.Infof("%s", string(b))
	logrus.Infof("Translated file %s upload Done", req.FileName)
}
