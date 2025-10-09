package main

import (
	"github.com/gin-gonic/gin"
	appError "github.com/ginanjar-template-golang/shared-pkg/errors"
	"github.com/ginanjar-template-golang/shared-pkg/logger"
	"github.com/ginanjar-template-golang/shared-pkg/middleware"
	"github.com/ginanjar-template-golang/shared-pkg/response"
	"github.com/ginanjar-template-golang/shared-pkg/utils"
)

func configLogger() {
	logger.Init(logger.Config{
		LogglyToken: "",
		LogglyTag:   "service-shared-pkg",
		Enabled:     false, // true = kirim ke Loggly, false = hanya print di console
	})
}

func main() {
	r := gin.Default()

	configLogger()

	r.Use(middleware.CORS())
	r.Use(middleware.RequestLogger())
	r.Use(middleware.Recovery())

	r.GET("/success-get", func(c *gin.Context) {
		reqID := utils.NewRequestID()
		c.Set("request_id", reqID)

		user := map[string]any{
			"id":   1,
			"name": "john",
		}

		response.Success(c, "Success get data", map[string]any{"user": user})
	})

	r.GET("/success-pagination", func(c *gin.Context) {
		reqID := utils.NewRequestID()
		c.Set("request_id", reqID)

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

		response.PaginationResponse(c, "success get users", pagination)
	})

	r.POST("/success-create", func(c *gin.Context) {
		reqID := utils.NewRequestID()
		c.Set("request_id", reqID)

		user := map[string]any{
			"id":   1,
			"name": "john",
		}

		response.Created(c, "User created", user)
	})

	r.PATCH("/success-update", func(c *gin.Context) {
		reqID := utils.NewRequestID()
		c.Set("request_id", reqID)

		user := map[string]any{
			"id":   1,
			"name": "john",
		}

		response.Updated(c, "User updated", user)
	})

	r.DELETE("/success-update", func(c *gin.Context) {
		reqID := utils.NewRequestID()
		c.Set("request_id", reqID)

		response.Deleted(c, "User deleted")
	})

	r.GET("/error", func(c *gin.Context) {
		appError.ResourceNotFound("user", "error test")
		// return nil, err

		// kalau tidak error
		// response.Success(c, "Success get user", nil)
	})

	// Contoh endpoint error
	r.GET("/panic", func(c *gin.Context) {
		panic("unexpected error example")
	})

	// Setelah itu tinggal pakai di mana pun:
	// logger.Info("User created successfully", map[string]interface{}{
	// 	"user_id": 42,
	// 	"email":   "test@example.com",
	// })

	// logger.Error("Payment failed", map[string]interface{}{
	// 	"request_id": "req-456",
	// 	"amount":     150000,
	// 	"error":      "timeout",
	// })

	r.Run(":8080")
}
