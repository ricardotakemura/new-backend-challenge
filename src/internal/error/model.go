package error

import (
	"encoding/json"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

type ErrorModel struct {
	bundle *i18n.Bundle
}

func NewErrorModel() *ErrorModel {
	var bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	bundle.LoadMessageFile("../resources/en.json")
	bundle.LoadMessageFile("../resources/pt-BR.json")
	return &ErrorModel{bundle: bundle}
}

func (model ErrorModel) GetById(id string, lang string, params map[string]string) Error {
	localizer := i18n.NewLocalizer(model.bundle, lang, language.English.String())
	code := ""
	switch id {
	case "product_not_found":
		code = "0001"
	case "invalid_quantity":
		code = "0002"
	case "insufficient_stock":
		code = "0003"
	case "invalid_body":
		code = "0004"
	case "invalid_path":
		code = "0005"
	default:
		code = "0000"
	}
	localizeConfig := i18n.LocalizeConfig{
		MessageID: id,
	}
	if params != nil {
		localizeConfig = i18n.LocalizeConfig{
			MessageID:    id,
			TemplateData: params,
		}
	}
	localization, _ := localizer.Localize(&localizeConfig)
	return Error{Code: code, Message: localization}
}
