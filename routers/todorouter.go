package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	Description string `json:"description"`
}

func GetAllTodo(c *gin.Context) {
	todos := []todo{
		{Description: "Initial todo"},
	}
	c.IndentedJSON(http.StatusOK, todos)
}
