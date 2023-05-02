package dataWork

import (
	"encoding/json"
	"github.com/plitn/wb_l2/develop/dev11/models"
	"log"
)

func ParseSliceToJson(data []models.EventModel) []byte {
	resp := make(map[string][]models.EventModel)
	resp["result"] = data
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	return jsonResp
}

func ParseStringToJson(data string, ok bool) []byte {
	if ok {
		resp := make(map[string]string)
		resp["result"] = data
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		return jsonResp
	} else {
		resp := make(map[string]string)
		resp["error"] = data
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		return jsonResp
	}
}
