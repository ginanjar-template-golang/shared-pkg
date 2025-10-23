package main

import (
	"github.com/gin-gonic/gin"
	appError "github.com/ginanjar-template-golang/shared-pkg/errors"
	"github.com/ginanjar-template-golang/shared-pkg/logger"
	"github.com/ginanjar-template-golang/shared-pkg/middleware"
	httpResponse "github.com/ginanjar-template-golang/shared-pkg/response/http_response"
	"github.com/ginanjar-template-golang/shared-pkg/validator"
)

func configLogger() {
	logger.Init(logger.Config{
		LogglyUrl:   "https://logs-01.loggly.com/inputs/%s/tag/%s",
		LogglyToken: "", //your-loggly-token
		LogglyTag:   "service-shared-pkg",
		Environment: "dev", // dev (TRACE,DEBUG,INFO,WARN,ERROR) | staging (TRACE,INFO,WARN,ERROR) | prod (WARN,ERROR)
		AllLogLevel: false,
	})
}

// http test
func main() {
	r := gin.Default()

	configLogger()

	r.Use(middleware.CORS())
	r.Use(middleware.RequestLogger())
	r.Use(middleware.Recovery())

	logger.Trace("Trace message", map[string]any{"foo": "bar"})
	logger.Debug("Debug message", map[string]any{"foo": "bar"})
	logger.Info("Info message", map[string]any{"foo": "bar"})
	logger.Warn("Warning message", map[string]any{"foo": "bar"})
	logger.Error("Error message", map[string]any{"foo": "bar"})

	r.GET("/success-get", func(c *gin.Context) {
		user := map[string]any{
			"id":   1,
			"name": "john",
		}

		httpResponse.Success(c, "successGet", map[string]any{"user": user})
	})

	r.GET("/success-pagination", func(c *gin.Context) {
		users := map[string]any{
			"id":   1,
			"name": "john",
		}

		pagination := httpResponse.Pagination{
			Page:     1,
			Size:     10,
			Limit:    10,
			TotalRow: len(users),
			Results:  users,
		}

		httpResponse.PaginationResponse(c, "successGetPagination", pagination)
	})

	r.POST("/success-create", func(c *gin.Context) {
		user := map[string]any{
			"id":   1,
			"name": "john",
		}

		httpResponse.Created(c, "successCreate", user)
	})

	r.PATCH("/success-update", func(c *gin.Context) {
		user := map[string]any{
			"id":   1,
			"name": "john",
		}

		httpResponse.Updated(c, "successUpdate", user)
	})

	r.DELETE("/success-update", func(c *gin.Context) {
		httpResponse.Deleted(c, "successDelete")
	})

	r.GET("/error", func(c *gin.Context) {
		err := appError.AlreadyUsedError("user", nil)
		httpResponse.FromAppError(c, err)
	})

	// Contoh endpoint error
	r.GET("/panic", func(c *gin.Context) {
		panic("unexpected error example")
	})

	r.POST("/validation", func(c *gin.Context) {
		type RegisterDto struct {
			Username string `json:"username" validate:"required"`
			Email    string `json:"email" validate:"required,email,notexample"`
			Password string `json:"password" validate:"required,min=6,strongpassword"`
		}

		var params RegisterDto

		if appErr := validator.ValidateRequest(c, &params); appErr != nil {
			httpResponse.FromAppError(c, *appErr)
			return
		}
	})

	r.Run(":8080")

	// go TestGrpcServer()
}
