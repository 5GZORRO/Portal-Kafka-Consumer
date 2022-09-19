package api

import (
	"encoding/json"
	"log"
	"net/http"

	"portal-kafka-consumer/model"
	"portal-kafka-consumer/redis"
	"portal-kafka-consumer/utils"
)

func GetViolations(w http.ResponseWriter, r *http.Request) {
	log.Println("Fetching Violations..")

	// List of violations to be retrieved
	listViolations := []model.Violation{}

	// Get All stored violations
	keys, _ := redis.GetAllKeys()

	// Get value for each key, bundle them up to send to React and delete them from DB which is basically a Stack
	for _, key := range keys {
		violationJson, _ := redis.Get(key) // Get specific violation

		// Bundle it
		violation := model.Violation{}
		utils.JsonToStruct(&violation, violationJson)
		listViolations = append(listViolations, violation)

		redis.Delete(key) // Delete that violation entry from stack
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(listViolations)
}
