package queue

import (
	"encoding/json"

	"net/http"

	"github.com/m-moris/azure-go-sample-functions/pkg/models"
	"github.com/m-moris/azure-go-sample-functions/pkg/utils"
	"go.uber.org/zap"
)

var logger *zap.Logger

func init() {
	logger = utils.Getlogger()
}

func QueueTriggerHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info("QueueTriggerHandler invoked")
	var invokeReq models.InvokeRequest
	d := json.NewDecoder(r.Body)
	err := d.Decode(&invokeReq)
	if err != nil {
		logger.Error("Error decoding request body", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	logger.Info("Request", zap.Any("request", invokeReq))

	result := "Hello " + invokeReq.Data["queue"].(string)
	invokeResponse := models.InvokeResponse{Logs: []string{"test log1", "test log2"}, ReturnValue: result}
	js, err := json.Marshal(invokeResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
