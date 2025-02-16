package error

import (
	"errors"
	"fmt"
	"strings"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type ValidationResponse struct {
	Field  string `json:"field,omitempty"`
	Message string `json:"message,omitempty"`
}

var ErrValidator = map[string]string{}

func ErrValidatonResponse(err error) (validationResonse []ValidationResponse) {
	var fieldErrors validator.ValidationErrors
	if errors.As(err, &fieldErrors) {
		for _, err := range fieldErrors {
			switch err.Tag() {
			case "required":
				validationResonse = append(validationResonse, ValidationResponse{
					Field:  err.Field(),
					Message: fmt.Sprintf("%s is required", err.Field()),
				})
			case "email":
				validationResonse = append(validationResonse, ValidationResponse{
					Field:  err.Field(),
					Message: fmt.Sprintf("%s is not valid email", err.Field()),
				})
			default:
				ErrValidator, ok := ErrValidator[err.Tag()]
				if ok {
					count := strings.Count(ErrValidator, "%s")
					if count == 1 {
						validationResonse = append(validationResonse, ValidationResponse{
							Field:  err.Field(),
							Message: fmt.Sprintf(ErrValidator, err.Field()),
						})
					} else {
						validationResonse = append(validationResonse, ValidationResponse{
							Field:  err.Field(),
							Message: fmt.Sprintf(ErrValidator, err.Field(), err.Param()),
						})
					}
				} else {
					validationResonse = append(validationResonse, ValidationResponse{
						Field:  err.Field(),
						Message: fmt.Sprintf("somethig wrong on %s; %s", err.Field(), err.Tag()),
					})
				}
			}
		}
	}
	return validationResonse
}

func WrapError(err error) error {
	logrus.Errorf("error: %v", err)
	return err
}