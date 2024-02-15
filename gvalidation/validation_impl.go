package gvalidation

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	entranslation "github.com/go-playground/validator/v10/translations/en"
	idtranslations "github.com/go-playground/validator/v10/translations/id"

	"github.com/SyaibanAhmadRamadhan/gocatch/gcommon"
)

type ErrValidate struct {
	Err map[string]string
}

func (e *ErrValidate) Error() string {
	str := ""
	for _, v := range e.Err {
		str += v + ". "
	}
	return str
}

func (v *Validation) StructM(s interface{}) error {
	err := v.Validate.Struct(s)
	if err != nil {
		var validationErrors validator.ValidationErrors
		ok := errors.As(err, &validationErrors)
		if !ok {
			return err
		}

		res := make(map[string]string)
		for _, validationError := range validationErrors {
			res[v.GetField(validationError.Namespace())] = validationError.Translate(v.Transalation)
		}
		return &ErrValidate{Err: res}
	}

	return nil
}

func (v *Validation) GetField(namespace string) string {
	strSplit := strings.Split(namespace, ".")
	if len(strSplit) > 1 {
		return strings.Join(strSplit[1:], ".")
	}
	return namespace
}

func (v *Validation) PrintFieldValidationError(field validator.FieldError) {
	fmt.Println("Field: ", field.Field())
	fmt.Println("Tag: ", field.Tag())
	fmt.Println("Value: ", field.Value())
	fmt.Println("StructField: ", field.StructField())
	fmt.Println("ActualTag: ", field.ActualTag())
	fmt.Println("Kind String: ", field.Kind().String())
	fmt.Println("Namespace: ", field.Namespace())
	fmt.Println("Param: ", field.Param())
	fmt.Println("StructNamespace: ", field.StructNamespace())
	fmt.Println("Type String: ", field.Type().String())
}

func (v *Validation) InitIdTranslation() {
	idLocale := id.New()
	uni := ut.New(idLocale, idLocale)

	trans, ok := uni.GetTranslator("id")
	if !ok {
		panic("invalid locale")
	}

	err := idtranslations.RegisterDefaultTranslations(v.Validate, trans)
	gcommon.PanicIfError(err)

	v.IdTranslation = trans

}

func (v *Validation) InitEnTranslation() {
	enLocale := en.New()
	uni := ut.New(enLocale, enLocale)

	trans, ok := uni.GetTranslator("en")
	if !ok {
		panic("invalid locale")
	}

	err := entranslation.RegisterDefaultTranslations(v.Validate, trans)
	gcommon.PanicIfError(err)

	v.EnTranslation = trans
}

func (v *Validation) SetTranslation(t Transalation) {
	switch t {
	case En:
		v.Transalation = v.EnTranslation
	case Id:
		v.Transalation = v.IdTranslation
	default:
		panic("invalid translation")
	}
}
