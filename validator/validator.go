package validator

import (
	"encoding/json"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ginanjar-template-golang/shared-pkg/constants"
	appError "github.com/ginanjar-template-golang/shared-pkg/errors"
	"github.com/ginanjar-template-golang/shared-pkg/logger"
	"github.com/ginanjar-template-golang/shared-pkg/translator"
	"github.com/go-playground/validator/v10"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var validate *validator.Validate

func Init() *validator.Validate {
	validate = validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return fld.Name
		}
		return name
	})
	return validate
}

func GetValidator() *validator.Validate {
	if validate == nil {
		validate = Init()
	}
	return validate
}

func ValidateRequest(c *gin.Context, data any) *appError.AppError {
	v := GetValidator()

	if err := c.ShouldBindJSON(data); err != nil {
		msg := translator.GetMessageGlobal("invalidRequest")
		logger.LogMapLevel("warn", constants.InvalidRequestData, msg, err.Error())
		return &appError.AppError{
			Code:       constants.BadRequest,
			MessageKey: "invalidRequest",
			Data:       err.Error(),
		}
	}

	if err := v.Struct(data); err != nil {
		errMap := make(map[string]string)
		for _, e := range err.(validator.ValidationErrors) {
			field := e.Field()
			tag := e.Tag()
			param := e.Param()
			errMap[field] = buildMessage(field, tag, param)
		}

		jsonMsg, _ := json.Marshal(errMap)
		logger.LogMapLevel("info", constants.ValidationError, translator.GetMessageGlobal("validationFailed"), jsonMsg)

		return &appError.AppError{
			Code:       constants.BadRequest,
			MessageKey: "validationFailed",
			Data:       errMap,
		}
	}

	return nil
}

func buildMessage(field, tag, param string) string {
	template := translator.GetMessageGlobal(tag)
	c := cases.Title(language.Und)
	fieldName := c.String(field)
	template = strings.ReplaceAll(template, "{field}", fieldName)
	template = strings.ReplaceAll(template, "{param}", param)
	return template
}
