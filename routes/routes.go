package routes

import (
	"github.com/extimsu/JsonRestApi/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	server.GET("/events", getEvents)                              // GET, POST, PUT, DELETE,	PATCH
	server.GET("/users", middlewares.Authenticate, getUsers)      // GET, POST, PUT, DELETE,	PATCH
	server.GET("/events/:id", middlewares.Authenticate, getEvent) // /events/1, /events/2, /events/3
	server.POST("/events", middlewares.Authenticate, createEvent)
	server.PUT("/events/:id", middlewares.Authenticate, updateEvent)
	server.DELETE("/events/:id", middlewares.Authenticate, deleteEvent)
	server.POST("/events/:id/register", middlewares.Authenticate, registerForEvent)
	server.DELETE("/events/:id/register", middlewares.Authenticate, cancelRegistration)
	server.POST("/signup", signup)
	server.POST("/login", login)
}
