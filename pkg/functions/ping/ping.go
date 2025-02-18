package ping

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/m-moris/azure-go-sample-functions/pkg/utils"
	"go.uber.org/zap"
)

type pingResponse struct {
	Message  string    `json:"message"`
	DateTime time.Time `json:"dateTime"`
}

var logger *zap.Logger

func init() {
	logger = utils.Getlogger()
}

// pingHandler  provides a simple http get response
func PingHandler(w http.ResponseWriter, r *http.Request) {

	logger.Info("PingHandler invoked")

	name := r.URL.Query().Get("name")
	response := pingResponse{
		Message:  "Hello go function : " + name,
		DateTime: time.Now(),
	}

	logger.Info("Response", zap.Any("response", response))
	responseJson, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJson)
}
