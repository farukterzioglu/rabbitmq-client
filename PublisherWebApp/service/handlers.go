package service

import (
	"net/http"
	"encoding/json"
	"strconv"
	"github.com/farukterzioglu/rabbitMqClient/Utilities"
)

func PublishHandler(w http.ResponseWriter, r *http.Request){

}

func HealthCheck(w http.ResponseWriter, r *http.Request){
	var rabbitMqHelper Utilities.IRabbitMqHelper
	rabbitMqHelper = &Utilities.RabbitMqHelper{}

	conn, err := rabbitMqHelper.GetRabbitMqConnection(
		rabbitMqSettings.HostName, rabbitMqSettings.UserName, rabbitMqSettings.Password)

	if err != nil || conn == nil{
		data, _ := json.Marshal(healthCheckResponse{ Status : "RabbitMQ is unaccessible"})
		writeJsonResponse(w, http.StatusServiceUnavailable, data)
	} else {
		data, _ := json.Marshal(healthCheckResponse{ Status : "UP"})
		writeJsonResponse(w, http.StatusOK, data)
	}
}

func writeJsonResponse(w http.ResponseWriter, status int, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(status)
	w.Write(data)
}

type healthCheckResponse struct {
	Status string `json:"status"`
}