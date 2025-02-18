package hello

import (
	"encoding/json"
	"net/http"

	"github.com/m-moris/azure-go-sample-functions/pkg/utils"
	"go.uber.org/zap"
)

type helloRequest struct {
	Name string `json:"name"`
}

type helloResponse struct {
	Message string `json:"message"`
}

var logger *zap.Logger

func init() {
	logger = utils.Getlogger()
}

// helloHandler provides a simple http post response with request body
func HelloHandler(w http.ResponseWriter, r *http.Request) {

	logger.Info("HelloHandler invoked")

	var req helloRequest
	d := json.NewDecoder(r.Body)
	err := d.Decode(&req)
	if err != nil {
		logger.Error("Error decoding request body", zap.Error(err))
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	response := helloResponse{
		Message: "Hello " + req.Name,
	}

	responseJson, err := json.Marshal(response)

	if err != nil {
		logger.Error("Error marshalling response", zap.Error(err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJson)
}
