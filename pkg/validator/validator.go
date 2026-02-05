package validator

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	entranslations "github.com/go-playground/validator/v10/translations/en"
	"reflect"
	"strings"
)

var (
	Trans ut.Translator
)

func NewValidator() *validator.Validate {
	v := validator.New()

	// 1. TAMBAHKAN INI: Supaya error message pake nama di tag json="...", bukan nama struct
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	// Setup Translator
	english := en.New()
	uni := ut.New(english, english)
	Trans, _ = uni.GetTranslator("en")

	// Register default translations
	_ = entranslations.RegisterDefaultTranslations(v, Trans)

	return v
}
