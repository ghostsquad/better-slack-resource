package slackoff

import (
	"gopkg.in/go-playground/validator.v9"
	"io"
	"fmt"
)

func InitValidator() *validator.Validate {
	return validator.New()
}

type ValidatorRegisterable interface {
	RegisterValidations(val *validator.Validate)
}

func Validate(va ValidatorRegisterable, vr *validator.Validate) validator.ValidationErrors {
	va.RegisterValidations(vr)

	errs := vr.Struct(va)
	if errs != nil {
		return errs.(validator.ValidationErrors)
	}

	return nil
}

func WriteValidationErrors(errs validator.ValidationErrors, w io.Writer) {
	if errs != nil {
		for _, err := range errs {
			fmt.Fprintf(w, "%s", err.Namespace())
			fmt.Fprintf(w, "%s", err.Field())
			// can differ when a custom TagNameFunc is registered or
			fmt.Fprintf(w, "%s", err.StructNamespace())
			// by passing alt name to ReportError like below
			fmt.Fprintf(w, "%s", err.StructField())
			fmt.Fprintf(w, "%s", err.Tag())
			fmt.Fprintf(w, "%s", err.ActualTag())
			fmt.Fprintf(w, "%s", err.Kind())
			fmt.Fprintf(w, "%s", err.Type())
			fmt.Fprintf(w, "%s", err.Value())
			fmt.Fprintf(w, "%s", err.Param())
			fmt.Fprintf(w, "")
		}
	}
}
