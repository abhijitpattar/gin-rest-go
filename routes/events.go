package routes

import (
	"net/http"
	"strconv"

	"github.com/abhijitpattar/gin-rest-go/models"
	"github.com/gin-gonic/gin"
)

func getEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}

	event, err := models.GetEventByID(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "no event for this id present"})
		return
	}

	context.JSON(http.StatusFound, event)
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fecth events, try again later"})
		return
	}
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {

	var event models.Event
	//fmt.Println("in create event", event)
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse the request data"})
		return
	}

	event.UserID = context.GetInt64("userId")
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not CREATE events, try again later"})
	}
	context.JSON(http.StatusCreated, gin.H{"message": " event created", "event": event})
}

func updateEvent(context *gin.Context) {
	// check if ID is int value
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "ID should be a number value"})
		return
	}

	// check if event is present in DB
	event, err := models.GetEventByID(id)
	if err != nil {
		context.JSON(http.StatusNoContent, gin.H{"message": "ID is not present in DB"})
		return
	}

	// prevent other users from updating events that are not created by them
	user_id_from_context := context.GetInt64("userId")
	if event.UserID != user_id_from_context {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "user not allowed to update"})
		return
	}

	var updateEvent models.Event
	err = context.ShouldBindJSON(&updateEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse the request data"})
		return
	}

	updateEvent.ID = id

	err = updateEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not update event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "event updated successfully"})
}

func deleteEvent(context *gin.Context) {
	// check if ID is int value
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "ID should be a number value"})
		return
	}

	// check if event is present in DB
	event, err := models.GetEventByID(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "ID is not present in DB"})
		return
	}

	// prevent other users from deleting events that are not created by them
	user_id_from_context := context.GetInt64("userId")
	if event.UserID != user_id_from_context {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "user not allowed to delete"})
		return
	}

	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not delete event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "event deleted successfully"})
}
