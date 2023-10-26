package handlers

import (
	"job-application-api/internal/auth"
	"job-application-api/internal/middleware"
	"job-application-api/internal/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func API(a *auth.Auth, sc service.UserService) *gin.Engine {
	r := gin.New()
	m, err := middleware.NewMid(a)
	if err != nil {
		log.Panic("middleware not setup")
	}
	h := handler{
		service: sc,
	}
	r.Use(middleware.Log(), gin.Recovery())
	r.GET("/check", m.Authenticate((check)))
	r.POST("/signup", h.SignUp)

	return r
}
func check(c *gin.Context) {
	c.JSON(http.StatusOK, "Msg :ok")

}
