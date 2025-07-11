package routers

import (
	"github.com/go-playground/validator/v10"
	"github.com/ian995/UniqueBank/pkg/utils"
)

var validCurrency validator.Func = (func(fl validator.FieldLevel) bool {
	if currency, ok := fl.Field().Interface().(string); ok {
		return utils.IsSupportedCurrency(currency)
	}	
	return false
})