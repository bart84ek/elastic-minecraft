package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func logUser(c *gin.Context) {
	user := c.MustGet(gin.AuthUserKey).(string)
	log.Printf("User %s loggeding", user)
}

func IndexEndpoint() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html.tmpl", gin.H{
			"title": "Minecraft Poligon Dzieciaków",
		})
	}
}

func StatusEndpoint() gin.HandlerFunc {
	return func(c *gin.Context) {
		status, err := mineStatus()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": status})
	}
}

func StartEndpoint() gin.HandlerFunc {
	return func(c *gin.Context) {
		logUser(c)
		status, err := mineStatus()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if status == "running" || status == "starting" {
			c.JSON(http.StatusConflict, gin.H{"status": "already running"})
		} else {
			go startMine()
			c.JSON(http.StatusOK, gin.H{"status": "starting"})
		}
	}
}

func StopEndpoint() gin.HandlerFunc {
	return func(c *gin.Context) {
		logUser(c)
		go stopMine()

		c.JSON(http.StatusOK, gin.H{"status": "stopping"})
	}
}
