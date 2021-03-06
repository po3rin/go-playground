package handler

import (
	l "go-playground/try-logrus-stackdriver/driver"

	"github.com/gin-gonic/gin"
)

// AllHandler try logging all
func AllHandler(c *gin.Context) {
	l.Info("ii")
	l.Warn("ww")
	l.Error("ee")
	c.JSON(200, gin.H{
		"message": "ok",
	})
}

// InfoHandler try logging all
func InfoHandler(c *gin.Context) {
	l.Info("ii")
	c.JSON(200, gin.H{
		"message": "ok",
	})
}

// WarnHandler try logging all
func WarnHandler(c *gin.Context) {
	l.Warn("ee")
	c.JSON(200, gin.H{
		"message": "ok",
	})
}

// ErrorHandler try logging all
func ErrorHandler(c *gin.Context) {
	l.Error("ee")
	c.JSON(200, gin.H{
		"message": "ok",
	})
}
