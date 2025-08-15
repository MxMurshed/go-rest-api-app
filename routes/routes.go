package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-rest-api/middlewares"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/events", getEvents)
	router.GET("/events/:id", getEvent)

	// middleware to authenticate all routes in the group
	authenticated := router.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)

	router.POST("/signup", signup)
	router.POST("/login", login)
}
