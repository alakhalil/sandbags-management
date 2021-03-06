package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"team2/sandsack-management-backend/models"
	"team2/sandsack-management-backend/service"
)

// ListOrder
// @Description ListOrder - listing all orders
// @Summary ListOrder - listing all orders
// @Accept json
// @Param Authorization header string true " "
// @Success 200 {array} models.Order
// @Failure 500 {object} models.ErrorResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Tags Order
// @Router /order/list [get]
func (a *App) ListOrder(c *gin.Context) {
	claims, err := GetClaims(c)
	if err != nil {
		return
	}
	orders, err := service.GetOrderList(a.DB, claims.Id)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, orders)
	return
}

// GetOrder
// @Description GetOrder - order by id
// @Summary GetOrder - order by id
// @Accept json
// @Param Authorization header string true " "
// @Param id path string true "Id of the order"
// @Success 200 {array} models.Order
// @Failure 500 {object} models.ErrorResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Tags Order
// @Router /order/ [get]
func (a *App) GetOrder(c *gin.Context) {
	id := c.Query("id")
	claims, err := GetClaims(c)
	if err != nil {
		return
	}

	orderId, err := strconv.Atoi(id)
	if err != nil {
		log.Println("Fehler beim Parsen", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrCode:    http.StatusBadRequest,
			ErrMessage: "Ungültiges Eingabeformat",
		})
		return
	}
	order, err := service.GetOrder(a.DB, claims.Id, orderId)
	if err != nil {
		log.Println("Fehler: GetOrder", err.Error())
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			ErrCode:    http.StatusInternalServerError,
			ErrMessage: "Da ist etwas schief gelaufen",
		})
		return
	}
	c.JSON(http.StatusOK, order)
	return

}
