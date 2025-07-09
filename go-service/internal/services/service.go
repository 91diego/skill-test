package services

import (
	"context"
	"net/http"

	"github.com/sevengit-wq/skill-test/config"
	"github.com/sevengit-wq/skill-test/pkg/errors"
	"github.com/sevengit-wq/skill-test/pkg/models"
	"github.com/sevengit-wq/skill-test/pkg/utils"
)

type Service interface {
	// auth
	GetCredentials(ctx context.Context) (models.LoginResponse, errors.ErrorHandler)
	// students
	GetStudentByID(ctx context.Context, id int) (models.StudentResponse, errors.ErrorHandler)
	GetReportByStudentID(ctx context.Context, id int) ([]byte, errors.ErrorHandler)
}

type service struct {
	logger     utils.Logger
	apiService config.ApiService
	httpClient http.Client
}

func NewService(logger utils.Logger, apiService config.ApiService) Service {
	return service{
		logger:     logger,
		apiService: apiService,
	}
}
