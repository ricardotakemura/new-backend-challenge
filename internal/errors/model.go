package errors

import (
	"encoding/json"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

type ErrorModel struct {
	localizer *i18n.Localizer
}

func NewErrorModel() *ErrorModel {
	var bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	bundle.LoadMessageFile("resources/en.json")
	bundle.LoadMessageFile("resources/pt.json")
	return &ErrorModel{localizer: i18n.NewLocalizer(bundle, language.English.String(), language.BrazilianPortuguese.String())}
}

func (model ErrorModel) PRODUCT_NOT_FOUND() *Error {
	localizeConfig := i18n.LocalizeConfig{
		MessageID: `product_not_found`,
	}
	localization, _ := model.localizer.Localize(&localizeConfig)
	return &Error{code: `0001`, message: localization}
}

func (model ErrorModel) INVALID_QUANTITY() *Error {
	localizeConfig := i18n.LocalizeConfig{
		MessageID: `invalid_quantity`,
	}
	localization, _ := model.localizer.Localize(&localizeConfig)
	return &Error{code: `0002`, message: localization}
}

func (model ErrorModel) INSUFFICIENT_STOCK() *Error {
	localizeConfig := i18n.LocalizeConfig{ //5
		MessageID: "insufficient_stock",
	}
	localization, _ := model.localizer.Localize(&localizeConfig)
	return &Error{code: `0003`, message: localization}
}
