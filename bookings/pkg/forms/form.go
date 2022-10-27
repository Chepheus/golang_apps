package forms

import (
	"errors"
	"net/url"
	"reflect"
	"strconv"
	"time"
	"unicode"
)

type form interface {
	setErrors(map[string][]string)
	GetErrors() map[string][]string
}

func createForm(f form, form url.Values) (*form, error) {
	t := reflect.TypeOf(f)
	formElement := reflect.ValueOf(&f).Elem()

	var err error
	var formValue string
	var validationRule string
	var formErrors = map[string][]string{}
	for _, fieldName := range formFields(f) {
		// skip private fields
		if unicode.IsLower(rune(fieldName[0])) == true {
			continue
		}

		field, found := t.Elem().FieldByName(fieldName)
		if !found {
			continue
		}

		formValue = form.Get(field.Tag.Get("field"))
		validationRule = field.Tag.Get("validation")

		err = validateFormField(validationRule, formValue)
		if err != nil {
			formErrors[fieldName] = append(formErrors[fieldName], err.Error())
		}

		err = fillFormField(formElement.Elem().Elem().FieldByName(fieldName), formValue)
		if err != nil {
			return nil, err
		}
	}

	f.setErrors(formErrors)

	return &f, nil
}

func validateFormField(validationRule, formValue string) error {
	if validationRule == "not_empty" {
		if formValue == "" {
			return errors.New("field can't be empty")
		}
	}

	return nil
}

func fillFormField(field reflect.Value, formValue string) error {
	switch field.Interface().(type) {
	case string:
		field.SetString(formValue)
	case int:
		formValueInt, err := strconv.Atoi(formValue)
		if err != nil {
			return err
		}
		field.SetInt(int64(formValueInt))

	case time.Time:
		formValueTime, err := time.Parse("2006-01-02", formValue)
		if err != nil {
			return err
		}
		field.Set(reflect.ValueOf(formValueTime))
	}

	return nil
}

func formFields(f form) []string {
	v := reflect.Indirect(reflect.ValueOf(f))
	typeOfS := v.Type()
	numField := v.NumField()

	fields := make([]string, numField)

	for i := 0; i < numField; i++ {
		fields[i] = typeOfS.Field(i).Name
	}

	return fields
}
