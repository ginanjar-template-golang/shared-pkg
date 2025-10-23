package validator

import (
	"encoding/json"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	grpcCode "github.com/ginanjar-template-golang/shared-pkg/constants/grpc_code"
	httpCode "github.com/ginanjar-template-golang/shared-pkg/constants/http_code"
	internalCode "github.com/ginanjar-template-golang/shared-pkg/constants/internal_code"
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
		if name == "-" || name == "" {
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
	if err := c.ShouldBindJSON(data); err != nil {
		msg := translator.GetMessageGlobal("invalidRequest")
		logger.LogMapLevel("warn", internalCode.InvalidRequestData, msg, err.Error())

		return &appError.AppError{
			HttpCode:   httpCode.BadRequest,
			GrpcCode:   grpcCode.InvalidArgument,
			MessageKey: "invalidRequest",
			Data:       err.Error(),
		}
	}

	return ValidateStruct(data)
}

func ValidateGrpcRequest(data any) *appError.AppError {
	return ValidateStruct(data)
}

func ValidateStruct(data any) *appError.AppError {
	v := GetValidator()
	if err := v.Struct(data); err != nil {
		errMap := make(map[string]string)
		for _, e := range err.(validator.ValidationErrors) {
			field := e.Field()
			tag := e.Tag()
			param := e.Param()
			errMap[field] = buildMessage(field, tag, param)
		}

		jsonMsg, _ := json.Marshal(errMap)
		logger.LogMapLevel("info", internalCode.ValidationError, translator.GetMessageGlobal("validationFailed"), jsonMsg)

		return &appError.AppError{
			HttpCode:   httpCode.BadRequest,
			GrpcCode:   grpcCode.InvalidArgument,
			MessageKey: "validationFailed",
			Data:       errMap,
		}
	}

	return nil
}

// buildMessage constructs a user-friendly validation message.
func buildMessage(field, tag, param string) string {
	template := translator.GetMessageGlobal(tag)
	c := cases.Title(language.Und)
	fieldName := c.String(field)

	template = strings.ReplaceAll(template, "{field}", fieldName)
	template = strings.ReplaceAll(template, "{param}", param)
	return template
}
