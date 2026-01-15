package routes

import (
	"net/http"
	"strconv"

	"github.com/abhijitpattar/gin-rest-go/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {

	userid_from_context := context.GetInt64("userId")

	// verify if event id in the request URL is integer type
	eventId_from_context, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id"})
		return
	}
	// verify if event is present in DB
	event, err := models.GetEventByID(eventId_from_context)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event"})
		return
	}

	event.RegisterEvent(userid_from_context)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not register user for the event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "user registered for the event successfully"})
}

func cancelRegistration(context *gin.Context) {
	userid_from_context := context.GetInt64("userId")

	// verify if event id in the request URL is integer type
	eventId_from_context, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id"})
		return
	}

	// create a event instance to call the cancle event
	var event models.Event
	event.ID = eventId_from_context
	// cancel the registration
	err = event.CancelRegistration(userid_from_context)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not cancel registration"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "successfully cancelled registration"})
}
