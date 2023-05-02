package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/plitn/wb_l2/develop/dev11/models"
	"time"
)

type EventsDB struct {
	db *sqlx.DB
}

func NewEventsDB(db *sqlx.DB) *EventsDB {
	return &EventsDB{
		db: db,
	}
}

func (edb *EventsDB) CreateNewEvent(name string, date time.Time) error {
	_, err := edb.db.Exec("INSERT INTO events (name, date) VALUES ($1, $2)", name, date)
	if err != nil {
		return err
	}
	return nil
}

func (edb *EventsDB) DeleteEvent(name string, date time.Time) error {
	query := fmt.Sprintf("delete from events where name = $1 and date = $2 ")
	_, err := edb.db.Exec(query, name, date)
	return err
}

func (edb *EventsDB) UpdateEvent(oldName string, oldDate time.Time, name string, date time.Time) error {
	query := fmt.Sprintf("update events set name = $1, date = $2 where name = $3 and date = $4")
	_, err := edb.db.Exec(query, name, date, oldName, oldDate)
	return err
}

func (edb *EventsDB) GetDateInRange(start time.Time, end time.Time) ([]models.EventModel, error) {
	var events []models.EventModel
	getDatesQuery := fmt.Sprintf("select name, date from events where date >= $1 and date <= $2")
	err := edb.db.Select(&events, getDatesQuery, start, end)
	return events, err
}
