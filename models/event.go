package models

import (
	"time"

	"github.com/go-rest-api/db"
)

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	DateTime    time.Time `json:"date_time" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	UserID      string    `json:"user_id"`
}

var events = []Event{}

func (e *Event) Save() error {
	stmt, err := db.DB.Prepare(`
			INSERT INTO events (name, description, date_time, location, user_id) 
			VALUES (?, ?, ?, ?, ?)
			`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.DateTime, e.Location, e.UserID)
	if err != nil {
		return err
	}

	e.ID, err = result.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}

func GetAllEvents() ([]Event, error) {
	rows, err := db.DB.Query("SELECT * FROM events")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var e Event
		err = rows.Scan(&e.ID, &e.Name, &e.Description, &e.DateTime, &e.Location, &e.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, e)
	}

	return events, nil
}

func GetEvent(id int64) (*Event, error) {
	row := db.DB.QueryRow("SELECT * FROM events WHERE id = ?", id)

	var e Event
	err := row.Scan(&e.ID, &e.Name, &e.Description, &e.DateTime, &e.Location, &e.UserID)
	if err != nil {
		return nil, err
	}

	return &e, nil
}

func (e *Event) Update() error {
	stmt, err := db.DB.Prepare(`
		UPDATE events 
		SET name = ?, description = ?, date_time = ?, location = ?, user_id = ?
		WHERE id = ?
		`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.Name, e.Description, e.DateTime, e.Location, e.UserID, e.ID)
	if err != nil {
		return err
	}

	return nil
}

func (e *Event) Delete() error {
	stmt, err := db.DB.Prepare(`
		DELETE FROM events WHERE id = ?
		`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.ID)
	if err != nil {
		return err
	}

	return nil
}
