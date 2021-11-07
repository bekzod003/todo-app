package handler

import (
	"net/http"
	todoapp "todo-app"

	"github.com/gin-gonic/gin"
)

func (h *Handler) sighUp(c *gin.Context) {
	var input todoapp.User

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.CreateUser(input)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) sighIn(c *gin.Context) {

}
