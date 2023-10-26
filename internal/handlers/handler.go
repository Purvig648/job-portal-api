package handlers

import (
	"job-application-api/internal/auth"
	"job-application-api/internal/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func API(a *auth.Auth) *gin.Engine {
	r := gin.New()
	r.Use(middleware.Log(), gin.Recovery())
	r.GET("/home", check)

	return r
}
func check(c *gin.Context) {
	c.JSON(http.StatusOK, "Msg : I am purvi All ok")

}
