package handlers

// import (
// 	"encoding/json"
// 	"job-application-api/internal/auth"
// 	"job-application-api/internal/middleware"
// 	"job-application-api/internal/models"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"github.com/golang-jwt/jwt/v5"
// 	"github.com/rs/zerolog/log"
// )

// func (h *handler) Process(c *gin.Context) {
// 	ctx := c.Request.Context()
// 	traceid, ok := ctx.Value(middleware.TraceIdkey).(string)
// 	if !ok {
// 		log.Error().Msg("traceid missing from context")
// 		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
// 			"error": http.StatusText(http.StatusInternalServerError),
// 		})
// 		return
// 	}
// 	_, ok = ctx.Value(auth.Ctxkey).(jwt.RegisteredClaims)
// 	if !ok {
// 		log.Error().Str("Trace Id", traceid).Msg("login required")
// 		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": http.StatusText(http.StatusUnauthorized)})
// 		return
// 	}
// 	var processData []models.UserApplicant

// 	err := json.NewDecoder(c.Request.Body).Decode(&processData)
// 	if err != nil {
// 		log.Error().Err(err).Str("trace id", traceid)
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
// 			"error": "provide valid details",
// 		})
// 		return
// 	}
// 	applicantData, err := h.service.ProcessApplications(ctx, processData)
// 	if err != nil {
// 		log.Error().Err(err).Str("trace id", traceid)
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
// 			"error": err,
// 		})
// 		return
// 	}
// 	c.JSON(http.StatusOK, applicantData)
// }
