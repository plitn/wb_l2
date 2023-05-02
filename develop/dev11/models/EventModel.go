package models

import "time"

type EventModel struct {
	Name string    `json:"name"`
	Date time.Time `json:"date"`
}
