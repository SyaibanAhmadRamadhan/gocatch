package gvalidation

import (
	"reflect"
	"strings"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type Transalation uint

const (
	En Transalation = iota
	Id
)

type Validation struct {
	Validate      *validator.Validate
	Transalation  ut.Translator
	IdTranslation ut.Translator
	EnTranslation ut.Translator
}

func New(opt ...validator.Option) *Validation {
	validate := validator.New(opt...)

	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}

		return name
	})

	return &Validation{
		Validate: validate,
		// Transalation: trans,
	}
}
