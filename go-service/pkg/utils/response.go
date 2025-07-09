package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/sevengit-wq/skill-test/pkg/constants"
)

type Response struct {
	Message string      `json:"message,omitempty"`
	Body    interface{} `json:"body,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

func SuccessPDFResponse(c *gin.Context, statusCode int, data []byte) {
	headers := map[string]string{
		"Content-Type":        constants.PDFContentType,
		"Content-Disposition": `inline; filename="student.pdf"`,
	}
	AddHeadersToGinResponse(c, headers)
	c.Data(statusCode, constants.PDFContentType, data)
}

func ErrorResponse(c *gin.Context, statusCode int, err interface{}) {
	response := Response{
		Error: err,
	}
	c.JSON(statusCode, response)
}
