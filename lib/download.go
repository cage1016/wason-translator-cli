package lib

import (
	"bytes"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/sirupsen/logrus"
	"github.com/watson-developer-cloud/go-sdk/languagetranslatorv3"
)

type DownloadRequest struct {
	Version        string
	APIKey         string
	URL            string
	DocumentID     string
	Accept         string
	OutputFileName string
}

func Download(req DownloadRequest) {
	logrus.Infof("Downloading document %s as %s", req.DocumentID, req.OutputFileName)

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

	result, _, responseErr := languageTranslator.GetTranslatedDocument(
		&languagetranslatorv3.GetTranslatedDocumentOptions{
			DocumentID: core.StringPtr(req.DocumentID),
			Accept:     core.StringPtr(req.Accept),
		},
	)
	if responseErr != nil {
		logrus.Fatal(responseErr)
	}
	if result != nil {
		buff := new(bytes.Buffer)
		buff.ReadFrom(result)
		file, _ := os.Create(req.OutputFileName)
		file.Write(buff.Bytes())
		file.Close()
	}
	result.Close()

	logrus.Infof("Downloaded document %s as %s", req.DocumentID, req.OutputFileName)
}
