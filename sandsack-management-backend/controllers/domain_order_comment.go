package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"team2/sandsack-management-backend/models"
	repo_order "team2/sandsack-management-backend/repository/order"
	"team2/sandsack-management-backend/service"
)

// CommentOrder
// @Description CommentOrder - user can write comments for the order
// @Summary CommentOrder - user can write comments for the order
// @Accept json
// @Produce json
// @Param Authorization header string true " "
// @Param input body models.CommentInput true "Comment input"
// @Success 200 {object} models.Order
// @Failure 500 {object} models.ErrorResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Tags Order
// @Router /order/comment [post]
func (a *App) CommentOrder(c *gin.Context) {
	var input models.CommentInput

	// check whether the structure of request is correct
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println("Fehler: CommentOrder: ", err.Error())
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrCode:    http.StatusBadRequest,
			ErrMessage: "Ungültige Anfrage",
		})
		return
	}

	claims, err := GetClaims(c)
	if err != nil {
		log.Println("Fehler: GetClaims", err.Error())
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			ErrCode:    http.StatusUnauthorized,
			ErrMessage: "Access Token ist ungültig",
		})
		return
	}

	permissions, err := service.GetUserOrderPermissions(a.DB, claims.Id, input.OrderId)

	flag := 0
	for _, i := range permissions {
		if i == models.DictPermissionName["CAN COMMENT"] {
			flag = 1
			break
		}
	}

	if flag == 0 {
		log.Println("Fehler: Kein Zugang: ")
		c.JSON(http.StatusForbidden, models.ErrorResponse{
			ErrCode:    http.StatusForbidden,
			ErrMessage: "Kein Zugang",
		})
		return
	}

	if len(input.Comment) == 0 {
		log.Println("Length of comments is 0", len(input.Comment))
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			ErrCode:    http.StatusBadRequest,
			ErrMessage: "Es muss den Kommentartext hinzugefügt sein",
		})
		return
	}

	var comments []models.Comment
	comments = append(comments, models.Comment{CommentText: input.Comment})
	log.Println("Length of comment list", len(comments))

	if err := service.InsertComments(a.DB, claims.Id, input.OrderId, comments); err != nil {
		return
	}
	user, err := service.GetUserByID(a.DB, claims.Id)
	order, err := service.GetOrder(a.DB, claims.Id, input.OrderId)
	logs := []models.Log{
		{
			OrderId:      input.OrderId,
			ActionTypeId: models.DictActionTypeName["COMMENTED"],
			UpdatedBy:    claims.Id,
			Description:  user.Name + " hat die Bestellung '" + order.Name + "' kommentiert",
		},
	}

	err = repo_order.InsertLogs(a.DB, logs)

	order, err = service.GetOrder(a.DB, claims.Id, input.OrderId)
	c.JSON(http.StatusOK, order)
	return
}
