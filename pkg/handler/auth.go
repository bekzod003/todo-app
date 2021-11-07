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
}

func (h *Handler) sighIn(c *gin.Context) {

}
