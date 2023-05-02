package main

import (
	"github.com/plitn/wb_l2/develop/dev11/database"
	"github.com/plitn/wb_l2/develop/dev11/handler"
	"github.com/plitn/wb_l2/develop/dev11/middleware"
	"log"
	"net/http"
	"time"
)

func main() {
	server := http.Server{
		Addr:         ":8000",
		Handler:      middleware.LoggerMiddleware(http.DefaultServeMux),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	db, err := database.NewPostgresql(database.ConfigDB{
		Host:     "localhost",
		Port:     "5432",
		Username: "postgresql",
		Password: "admin",
		DBName:   "notes-db",
		SSlMode:  "disable",
	})
	if err != nil {
		log.Fatalf("db err: %s", err)
	}
	rq := handler.NewRequestHandler(db)
	http.HandleFunc("/create_event", rq.CreateEventHandler)
	http.HandleFunc("/update_event", rq.UpdateEventHandler)
	http.HandleFunc("/delete_event", rq.DeleteEventHandler)
	http.HandleFunc("/events_for_day", rq.EventsForDayEventHandler)
	http.HandleFunc("/events_for_week", rq.EventsForWeekEventHandler)
	http.HandleFunc("/events_for_month", rq.EventsForMonthEventHandler)
	log.Println("starting server")
	server.ListenAndServe()
}
