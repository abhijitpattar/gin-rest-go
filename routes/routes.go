package routes

import (
	"github.com/abhijitpattar/gin-rest-go/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	// non authenticated route
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	// create a group of paths/routes, that need to be authenticated with a token before execution
	authenticated := server.Group("/")
	// add the middleware authentication process
	// this ensures that the middleware is always run before any of the routes in this group is executed
	authenticated.Use(middlewares.Authenticate)
	// add the routes to the group
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)

	// non authenticated route
	server.POST("/signup", signup)
	server.POST("/login", login)
}
