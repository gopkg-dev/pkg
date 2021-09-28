package fiberx

import (
	"log"
	"reflect"

	"github.com/gopkg-dev/pkg/errors"

	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

var (
	Validator = validator.New()
	trans     ut.Translator

	ErrInvalidArgument = errors.BadRequest("INVALID_PARAMETER", "参数错误")
	ErrValidate        = errors.BadRequest("PARAMETER_VALIDATE_ERROR", "参数校验错误")
)

func init() {
	Validator.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := fld.Tag.Get("query")
		if name == "" {
			name = fld.Tag.Get("label")
		}
		return name
	})
	uni := ut.New(zh.New())
	trans, _ = uni.GetTranslator("zh")
	if err := zhTranslations.RegisterDefaultTranslations(Validator, trans); err != nil {
		log.Fatal("Error while registering english translations")
	}
}

// Validate validates the input struct
func Validate(payload interface{}) error {

	err := Validator.Struct(payload)
	if err == nil {
		return nil
	}

	if e, ok := err.(*validator.InvalidValidationError); ok {
		return ErrInvalidArgument.WithMetadata(map[string]string{
			"error": e.Error(),
		})
	}

	if errs, ok := err.(validator.ValidationErrors); ok {
		metadata := make(map[string]string, 0)
		for _, fieldError := range errs {
			ns := fieldError.Field()
			metadata[ns] = fieldError.Translate(trans)
		}
		return ErrValidate.WithMetadata(metadata)
	}

	return nil
}
