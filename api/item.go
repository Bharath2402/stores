package API

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetItem() {
	//TBD
}

func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
