package handler

import (
	"github.com/jmoiron/sqlx"
	"github.com/plitn/wb_l2/develop/dev11/dataWork"
	"github.com/plitn/wb_l2/develop/dev11/database"
	"log"
	"net/http"
	"time"
)

type RequestHandler struct {
	ndb *database.EventsDB
}

func NewRequestHandler(db *sqlx.DB) *RequestHandler {
	return &RequestHandler{
		ndb: database.NewEventsDB(db),
	}
}

func writeResultOK(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(dataWork.ParseStringToJson(msg, true))
}

func writeResultBadRequest(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusBadRequest)
	w.Header().Set("Content-Type", "application/json")
	w.Write(dataWork.ParseStringToJson(msg, false))
}

func writeResultServiceUnavailable(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusServiceUnavailable)
	w.Header().Set("Content-Type", "application/json")
	w.Write(dataWork.ParseStringToJson(msg, false))
}

// POST
func (rh *RequestHandler) CreateEventHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("inside create handler")

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parsing request body", http.StatusBadRequest)
		return
	}

	// Retrieve the value of the "date" form key
	date := r.FormValue("date")
	name := r.FormValue("name")
	requestTime, err := dataWork.ParseDate(date)
	if err != nil {
		writeResultBadRequest(w, "invalid input")
		log.Fatalf(err.Error())
	}
	err = rh.ndb.CreateNewEvent(name, requestTime)
	if err != nil {
		writeResultServiceUnavailable(w, "creation error")
		log.Fatalf(err.Error())
	}

	writeResultOK(w, "event created")
}

func (rh *RequestHandler) UpdateEventHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("inside delete handler")

	if err := r.ParseForm(); err != nil {
		writeResultBadRequest(w, "request body parsing error")
		return
	}

	// Retrieve the value of the "date" form key
	date := r.FormValue("date")
	name := r.FormValue("name")
	newDate := r.FormValue("newDate")
	newName := r.FormValue("newName")
	requestTime, err := dataWork.ParseDate(date)
	if err != nil {
		writeResultBadRequest(w, "invalid input")
		log.Fatalf(err.Error())
	}
	newTime, err := dataWork.ParseDate(newDate)
	if err != nil {
		writeResultBadRequest(w, "invalid input")
		log.Fatalf(err.Error())
	}
	err = rh.ndb.UpdateEvent(name, requestTime, newName, newTime)
	if err != nil {
		writeResultServiceUnavailable(w, "update event error")
		log.Fatalf(err.Error())
	}

	writeResultOK(w, "event updated")
}

// POST
func (rh *RequestHandler) DeleteEventHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("inside delete handler")

	if err := r.ParseForm(); err != nil {
		writeResultBadRequest(w, "request body parsing error")

		return
	}

	// Retrieve the value of the "date" form key
	date := r.FormValue("date")
	name := r.FormValue("name")
	requestTime, err := dataWork.ParseDate(date)
	if err != nil {
		writeResultBadRequest(w, "invalid input")
		log.Fatalf(err.Error())
	}
	err = rh.ndb.DeleteEvent(name, requestTime)
	if err != nil {
		writeResultServiceUnavailable(w, "delete event error")
		log.Fatalf(err.Error())
	}

	writeResultOK(w, "event deleted")
}

// GET
func (rh *RequestHandler) EventsForDayEventHandler(w http.ResponseWriter, r *http.Request) {
	// queryString for get requests
	stringDate := r.URL.Query()["date"][0]
	requestTime, err := dataWork.ParseDate(stringDate)
	if err != nil {
		writeResultBadRequest(w, "invalid input")
		log.Fatalf(err.Error())
	}
	datesRange, err := rh.ndb.GetDateInRange(requestTime, requestTime.Add(time.Hour*24))
	if err != nil {
		writeResultServiceUnavailable(w, "dates in range of day error")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(dataWork.ParseSliceToJson(datesRange))
}

// GET
func (rh *RequestHandler) EventsForWeekEventHandler(w http.ResponseWriter, r *http.Request) {
	// queryString for get requests
	stringDate := r.URL.Query()["date"][0]
	requestTime, err := dataWork.ParseDate(stringDate)
	if err != nil {
		writeResultBadRequest(w, "invalid input")
		log.Fatalf(err.Error())
	}
	datesRange, err := rh.ndb.GetDateInRange(requestTime, requestTime.Add(time.Hour*24*7))
	if err != nil {
		writeResultServiceUnavailable(w, "dates in range of week error")

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(dataWork.ParseSliceToJson(datesRange))
}

// GET
func (rh *RequestHandler) EventsForMonthEventHandler(w http.ResponseWriter, r *http.Request) {
	// queryString for get requests
	stringDate := r.URL.Query()["date"][0]
	requestTime, err := dataWork.ParseDate(stringDate)
	if err != nil {
		writeResultBadRequest(w, "invalid input")
		log.Fatalf(err.Error())
	}
	datesRange, err := rh.ndb.GetDateInRange(requestTime, requestTime.AddDate(0, 1, 0))
	if err != nil {
		writeResultServiceUnavailable(w, "dates in range of month error")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(dataWork.ParseSliceToJson(datesRange))
}
