package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NoRouteHandler(c *gin.Context) {
    c.HTML(http.StatusNotFound, "404.html", nil)
}