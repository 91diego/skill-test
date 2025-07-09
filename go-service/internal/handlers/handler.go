package handlers

import (
	"github.com/sevengit-wq/skill-test/internal/services"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	// students
	GetReportByStudentID(c *gin.Context)
}

type handler struct {
	service services.Service
}

func NewHandler(service services.Service) Handler {
	return &handler{
		service: service,
	}
}

var _ Handler = &handler{}
