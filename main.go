package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ginanjar-template-golang/shared-pkg/errors"
	"github.com/ginanjar-template-golang/shared-pkg/response"
	"github.com/ginanjar-template-golang/shared-pkg/translator"
	"github.com/ginanjar-template-golang/shared-pkg/utils"
)

func main() {
	r := gin.Default()

	// Middleware untuk set request_id dan translator sesuai header lang
	r.Use(func(c *gin.Context) {
		// set request_id
		c.Set("request_id", "req-123456")

		// baca header lang
		lang := c.GetHeader("Accept-Language")
		if lang == "" {
			lang = "en"
		}

		// load translator sesuai lang
		var t translator.Translator
		switch lang {
		case "id":
			t = translator.NewTranslator("./translator/messages/id.json")
		default:
			t = translator.NewTranslator("./translator/messages/en.json")
		}

		c.Set("translator", t)
		c.Next()
	})

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

		response.PaginationResponse(c, 1, 10, 100, users)
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
		t := c.MustGet("translator").(translator.Translator)
		err := errors.ResourceNotFound(t, "user", "error test")

		fmt.Println(err)

		if err != (errors.InternalError{}) {
			// langsung response ke client
			response.FromInternalError(c, err)
			return
		}

		// kalau tidak error
		response.Success(c, "Success get user", nil)
	})

	r.Run(":8080")
}
