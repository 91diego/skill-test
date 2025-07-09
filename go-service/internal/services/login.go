package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/sevengit-wq/skill-test/pkg/constants"
	"github.com/sevengit-wq/skill-test/pkg/errors"
	"github.com/sevengit-wq/skill-test/pkg/mappers"
	"github.com/sevengit-wq/skill-test/pkg/models"
	"github.com/sevengit-wq/skill-test/pkg/utils"
)

func (s service) GetCredentials(ctx context.Context) (models.LoginResponse, errors.ErrorHandler) {
	credentialsResponse := models.LoginResponse{}
	credentials := &models.Login{}
	mappers.LoginPostMapper(credentials, s.apiService)

	jsonBytes, err := json.Marshal(credentials)
	if err != nil {
		return credentialsResponse, &errors.CustomError{
			ErrorCode:    http.StatusInternalServerError,
			ErrorMessage: err.Error(),
		}
	}
	payload := strings.NewReader(string(jsonBytes))

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/auth/login", s.apiService.URL), payload)
	if err != nil {
		s.logger.WithError(err).Error("request error on [GetCredentials method]")
		return credentialsResponse, &errors.CustomError{
			ErrorCode:    http.StatusInternalServerError,
			ErrorMessage: err.Error(),
		}
	}

	headers := map[string]string{
		"Content-Type": constants.JSONContentType,
	}
	utils.AddHeadersToHTTPRequest(req, headers)

	res, err := s.httpClient.Do(req)
	if err != nil {
		s.logger.WithError(err).Error("http client error on [GetCredentials method]")
		return credentialsResponse, &errors.CustomError{
			ErrorCode:    http.StatusInternalServerError,
			ErrorMessage: err.Error(),
		}
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		s.logger.WithError(err).Error("body read error on [GetCredentials method]")
		return credentialsResponse, &errors.CustomError{
			ErrorCode:    http.StatusInternalServerError,
			ErrorMessage: err.Error(),
		}
	}

	if res.StatusCode != http.StatusOK {
		s.logger.Errorf("validation error on [GetCredentials method] with status code [%v]", res.StatusCode)
		return credentialsResponse, errors.ValidationError(body, res.StatusCode)
	}

	mappers.LoginResponseMapper(res, &credentialsResponse)
	return credentialsResponse, nil
}
