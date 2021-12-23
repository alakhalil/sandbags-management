package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"team2/sandsack-management-backend/models"
	"team2/sandsack-management-backend/service"
)

// GetPriority
// @Description GetPriority - array of priorities in system
// @Summary GetPriority - array of priorities in system
// @Accept json
// @Param Authorization header string true "Bearer "
// @Success 200 {array} models.Priority
// @Failure 500 {object} models.ErrorResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Tags Core
// @Router /core/priority [get]
func (a *App) GetPriority(c *gin.Context) {
	priority, err := service.GetPriority(a.DB)
	if err != nil {
		log.Println("GetPriority error", err.Error())
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			ErrMessage: "something went wrong",
			ErrCode: http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, priority)
	return
}