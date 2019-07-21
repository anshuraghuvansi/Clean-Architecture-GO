package formatter

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response :
type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

//HandleResponse :
func HandleResponse(c *gin.Context, data interface{}, err error) {

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, Response{Status: -1, Message: err.Error()})
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, Response{Status: 0, Message: "Success", Data: data})
}
