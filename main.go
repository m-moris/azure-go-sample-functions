package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/m-moris/azure-go-sample-functions/pkg/functions/hello"
	"github.com/m-moris/azure-go-sample-functions/pkg/functions/ping"
	"github.com/m-moris/azure-go-sample-functions/pkg/functions/queue"
	"github.com/m-moris/azure-go-sample-functions/pkg/functions/timer"
	"github.com/m-moris/azure-go-sample-functions/pkg/utils"
	"go.uber.org/zap"
)

var logger *zap.Logger

func init() {
	logger = utils.Getlogger()
}

func main() {
	logger.Info("Start Go functions")
	customHandlerPort, exists := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT")
	if exists {
		logger.Info("FUNCTIONS_CUSTOMHANDLER_PORT: " + customHandlerPort)
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/api/ping", ping.PingHandler)
	mux.HandleFunc("/api/hello", hello.HelloHandler)
	mux.HandleFunc("/timer", timer.TimerTriggerHandler)
	mux.HandleFunc("/queue", queue.QueueTriggerHandler)
	fmt.Println("Go server Listening on: ", customHandlerPort)
	log.Fatal(http.ListenAndServe(":"+customHandlerPort, mux))
}
