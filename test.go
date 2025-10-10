package main

import (
	"github.com/gin-gonic/gin"
	appError "github.com/ginanjar-template-golang/shared-pkg/errors"
	"github.com/ginanjar-template-golang/shared-pkg/logger"
	"github.com/ginanjar-template-golang/shared-pkg/middleware"
	"github.com/ginanjar-template-golang/shared-pkg/response"
)

func configLogger() {
	logger.Init(logger.Config{
		LogglyToken: "",
		LogglyTag:   "service-shared-pkg",
		Enabled:     false,
	})
}

func main() {
	r := gin.Default()

	configLogger()

	r.Use(middleware.CORS())
	r.Use(middleware.RequestLogger())
	r.Use(middleware.Recovery())

	r.GET("/success-get", func(c *gin.Context) {
		user := map[string]any{
			"id":   1,
			"name": "john",
		}

		response.Success(c, "successGet", map[string]any{"user": user})
	})

	r.GET("/success-pagination", func(c *gin.Context) {
		users := map[string]any{
			"id":   1,
			"name": "john",
		}

		pagination := response.Pagination{
			Page:     1,
			Size:     10,
			Limit:    10,
			TotalRow: len(users),
			Results:  users,
		}

		response.PaginationResponse(c, "successGetPagination", pagination)
	})

	r.POST("/success-create", func(c *gin.Context) {
		user := map[string]any{
			"id":   1,
			"name": "john",
		}

		response.Created(c, "successCreate", user)
	})

	r.PATCH("/success-update", func(c *gin.Context) {
		user := map[string]any{
			"id":   1,
			"name": "john",
		}

		response.Updated(c, "successUpdate", user)
	})

	r.DELETE("/success-update", func(c *gin.Context) {
		response.Deleted(c, "successDelete")
	})

	r.GET("/error", func(c *gin.Context) {
		err := appError.ResourceNotFound("user", "error test")
		response.FromInternalError(c, err)
	})

	// Contoh endpoint error
	r.GET("/panic", func(c *gin.Context) {
		panic("unexpected error example")
	})

	r.Run(":8080")
}
