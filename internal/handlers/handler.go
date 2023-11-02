package handlers

import (
	"job-application-api/internal/auth"
	"job-application-api/internal/middleware"
	"job-application-api/internal/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func API(a auth.TokenAuth, sc service.UserService) *gin.Engine {
	r := gin.New()
	m, err := middleware.NewMid(a)
	if err != nil {
		log.Panic("middleware not setup")
		return nil
	}
	h, err := NewHandlerFunc(sc)
	if err != nil {
		log.Panic("handler not setup")
		return nil
	}
	r.Use(middleware.Log(), gin.Recovery())
	r.GET("/check", m.Authenticate((check)))
	r.POST("/signup", h.SignUp)
	r.POST("/login", h.Login)
	r.POST("/addcompany", m.Authenticate(h.AddCompany))
	r.GET("/viewcompany/:id", m.Authenticate(h.ViewCompany))
	r.GET("/viewAllcompany", m.Authenticate(h.ViewAllCompanies))
	r.POST("/addjob", m.Authenticate(h.AddJob))
	r.GET("/viewJobByCid/:cid", m.Authenticate(h.ViewJobByCompanyId))
	r.GET("/viewAllJobPostings", m.Authenticate(h.ViewAllJobs))
	r.GET("/viewJobById/:id", m.Authenticate(h.ViewJobById))

	return r
}
func check(c *gin.Context) {
	c.JSON(http.StatusOK, "Msg :ok")

}
