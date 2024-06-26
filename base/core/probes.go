package core

import (
	"app/base/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Liveness(c *gin.Context) {
	c.JSON(http.StatusOK, "ok")
}

func Readiness(c *gin.Context) {
	sqlDB, err := database.DB.DB()
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"err": err.Error()})
		return
	}
	err = sqlDB.Ping()
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "ok")
}

func InitProbes(app *gin.Engine) {
	// public routes - deprecated
	app.GET("/liveness", Liveness)
	app.GET("/readiness", Readiness)

	// public routes
	app.GET("/healthz", Liveness)
	app.GET("/livez", Liveness)
	app.GET("/readyz", Readiness)
}
