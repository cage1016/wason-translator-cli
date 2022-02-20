package lib

import (
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/sirupsen/logrus"
	"github.com/watson-developer-cloud/go-sdk/languagetranslatorv3"
)

type DeleteRequest struct {
	Version    string
	APIKey     string
	URL        string
	DocumentID string
}

func Delete(req *DeleteRequest) {
	logrus.Info("Deleting Document...")

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
		logrus.Fatal(languageTranslatorErr)
	}

	languageTranslator.SetServiceURL(req.URL)

	_, responseErr := languageTranslator.DeleteDocument(
		&languagetranslatorv3.DeleteDocumentOptions{
			DocumentID: core.StringPtr(req.DocumentID),
		},
	)
	if responseErr != nil {
		logrus.Fatal(responseErr)
	}

	logrus.Info("Document Deleted")
}
