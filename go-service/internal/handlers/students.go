package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sevengit-wq/skill-test/pkg/models"
	"github.com/sevengit-wq/skill-test/pkg/utils"
)

// GetReportByStudentID godoc
//
// @Summary      Generate student report
// @Description  Generates and returns a PDF report for a student by their ID.
// @Tags         Students
// @Accept       json
// @Produce      application/pdf
//
// @Success      200  {file}  application/pdf  "PDF report generated successfully"
// @Failure      400  {object}  utils.Response  "Bad request - Invalid ID or parameters"
// @Failure      404  {object}  utils.Response  "Not found"
// @Failure      500  {object}  utils.Response  "Internal server error"
//
// @Router       /api/v1/students/{id}/report [get]

func (h *handler) GetReportByStudentID(c *gin.Context) {
	var params models.StudentByIDPathParams

	if err := c.ShouldBindUri(&params); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	pdf, err := h.service.GetReportByStudentID(c.Request.Context(), params.ID)
	if err != nil {
		utils.ErrorResponse(c, err.Code(), err.Message())
		return
	}
	utils.SuccessPDFResponse(c, http.StatusOK, pdf)
}
