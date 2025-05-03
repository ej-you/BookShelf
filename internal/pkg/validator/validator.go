// Package validator provides interface for validate struct data
// by tags
package validator

import (
	"errors"
	"strings"

	validatorModule "github.com/go-playground/validator/v10"

	enLocale "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	enTrans "github.com/go-playground/validator/v10/translations/en"
)

var _ Validator = (*validator)(nil)

type Validator interface {
	Validate(s any) error
}

// Validator implementation
type validator struct {
	validatorInstance *validatorModule.Validate
	translator        ut.Translator
}

// Validator constructor
func NewValidator() Validator {
	en := enLocale.New()
	uni := ut.New(en, en)
	trans, _ := uni.GetTranslator("en")

	validate := validatorModule.New(validatorModule.WithRequiredStructEnabled())
	err := enTrans.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		panic(err)
	}

	return &validator{validate, trans}
}

// Validate given struct s (using pointer to this struct) with error handling to HTTPError
func (v validator) Validate(s any) error {
	err := v.validatorInstance.Struct(s)
	if err == nil { // NOT err
		return nil
	}

	// assert error to validatorModule.ValidationErrors
	var validateErrors validatorModule.ValidationErrors
	if !errors.As(err, &validateErrors) {
		return err
	}
	// handle error messages
	rawTranstaledMap := validateErrors.Translate(v.translator)

	// sort out errors and concat them into string
	transtaledStringSlice := make([]string, 0, len(rawTranstaledMap))
	for _, v := range rawTranstaledMap {
		transtaledStringSlice = append(transtaledStringSlice, strings.ToLower(v))
	}

	return errors.New(strings.Join(transtaledStringSlice, " && "))
}
