package routes

import (
	"net/http"
	"strconv"

	"github.com/extimsu/JsonRestApi/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not " + err.Error()})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch an event " + err.Error()})
		return
	}

	if err := event.Register(userId); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not register user to an event " + err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Registered."})

}

func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not " + err.Error()})
		return
	}

	var event models.Event
	event.ID = eventId

	if err := event.CancelRegistration(userId); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not cancel registration to an event " + err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Registration canceled."})
}
