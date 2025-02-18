package timer

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/m-moris/azure-go-sample-functions/pkg/models"
	"github.com/m-moris/azure-go-sample-functions/pkg/utils"
	"go.uber.org/zap"
)

var logger *zap.Logger

func init() {
	logger = utils.Getlogger()
}

func TimerTriggerHandler(w http.ResponseWriter, r *http.Request) {

	t := time.Now()
	logger.Info("TimerTrigger invoked at: ", zap.Time("time", t))

	// HTTPリクエストから、Function用のリクエストからデータを取得
	var invokeReq models.InvokeRequest
	d := json.NewDecoder(r.Body)
	decodeErr := d.Decode(&invokeReq)
	if decodeErr != nil {
		http.Error(w, decodeErr.Error(), http.StatusBadRequest)
		return
	}

	logger.Info("request", zap.Any("invokeReq", invokeReq))
	response := models.InvokeResponse{Logs: []string{"test log1", "test log2"}}
	js, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}
