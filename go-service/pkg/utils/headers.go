package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddHeadersToHTTPRequest(req *http.Request, headers map[string]string) {
	for key, value := range headers {
		req.Header.Add(key, value)
	}
}

func AddHeadersToGinResponse(c *gin.Context, headers map[string]string) {
	for key, value := range headers {
		c.Writer.Header().Add(key, value)
	}
}
