package models

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mustaphalimar/event-booking/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	Datetime    time.Time `binding:"required"`
	UserID      int
}

var events = []Event{}

func (event Event) Save() error {
	query := `INSERT INTO events(name,description,location,date_time,user_id)
				VALUES (?,?,?,?,?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(event.Name, event.Description, event.Location, event.Datetime, event.UserID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	event.ID = id
	return err
}

func GetEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)
	defer rows.Close()

	if err != nil {
		return nil, err
	}
	// context.JSON(http.StatusBadRequest, gin.H{
	// 	"success": false,
	// 	"data":     rows,
	// })

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.Datetime, &event.UserID)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func CreateEvent(context *gin.Context) {
	var event Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"err":     err,
		})
		return
	}
	event.ID = 1
	event.UserID = 1

	event.Save()

	context.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    event,
	})
}
