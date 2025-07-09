package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/sevengit-wq/skill-test/pkg/constants"
	"github.com/sevengit-wq/skill-test/pkg/errors"
	"github.com/sevengit-wq/skill-test/pkg/models"
	"github.com/sevengit-wq/skill-test/pkg/utils"
)

func (s service) GetStudentByID(ctx context.Context, id int) (models.StudentResponse, errors.ErrorHandler) {
	student := models.StudentResponse{}
	credentials, credentialsErr := s.GetCredentials(ctx)
	if credentialsErr != nil {
		s.logger.WithError(credentialsErr).Error("login error")
		return student, &errors.CustomError{
			ErrorCode:    credentialsErr.Code(),
			ErrorMessage: credentialsErr.Error(),
		}
	}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/students/%v", s.apiService.URL, id), nil)
	if err != nil {
		s.logger.WithError(err).Error("request error on [GetStudentByID method]")
		return student, &errors.CustomError{
			ErrorCode:    credentialsErr.Code(),
			ErrorMessage: credentialsErr.Error(),
		}
	}

	headers := map[string]string{
		"Content-Type":  constants.JSONContentType,
		"Authorization": credentials.AccessToken,
		"X-CSRF-Token":  credentials.CsrfToken,
		"Cookie": fmt.Sprintf("accessToken=%s; csrfToken=%s; refreshToken=%s;",
			credentials.AccessToken, credentials.CsrfToken, credentials.RefreshToken),
	}
	utils.AddHeadersToHTTPRequest(req, headers)

	res, err := s.httpClient.Do(req)
	if err != nil {
		s.logger.WithError(err).Error("http client error on [GetStudentByID method]")
		return student, &errors.CustomError{
			ErrorCode:    http.StatusInternalServerError,
			ErrorMessage: err.Error(),
		}
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		s.logger.WithError(err).Error("body read error on [GetStudentByID method]")
		return student, &errors.CustomError{
			ErrorCode:    http.StatusInternalServerError,
			ErrorMessage: err.Error(),
		}
	}

	if res.StatusCode != http.StatusOK {
		s.logger.Errorf("validation error on [GetStudentByID method] with status code [%v]", res.StatusCode)
		return student, errors.ValidationError(body, res.StatusCode)
	}

	if err = json.Unmarshal(body, &student); err != nil {
		s.logger.WithError(err).Error("json Unmarshal error on [GetStudentByID method]")
		return student, &errors.CustomError{
			ErrorCode:    http.StatusInternalServerError,
			ErrorMessage: err.Error(),
		}
	}

	return student, nil
}

func (s service) GetReportByStudentID(ctx context.Context, id int) ([]byte, errors.ErrorHandler) {
	student, err := s.GetStudentByID(ctx, id)
	if err != nil {
		s.logger.WithError(err).Error("error on [GetReportByStudentID method]")
		return []byte{}, &errors.CustomError{
			ErrorCode:    err.Code(),
			ErrorMessage: err.Error(),
		}
	}
	pdf, pdfErr := utils.GeneratePDF(student.Data, "Student Information Summary")
	if pdfErr != nil {
		return pdf, &errors.CustomError{
			ErrorCode:    http.StatusInternalServerError,
			ErrorMessage: err.Error(),
		}
	}
	return pdf, nil
}
