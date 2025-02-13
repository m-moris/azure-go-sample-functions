package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"go.uber.org/zap"
)

var logger *zap.Logger

func init() {
	logger, _ = zap.NewProduction()
}

func main() {
	logger.Info("Start Go functions")
	customHandlerPort, exists := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT")
	if exists {
		logger.Info("FUNCTIONS_CUSTOMHANDLER_PORT: " + customHandlerPort)
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/api/ping", pingHandler)
	mux.HandleFunc("/api/hello", helloHandler)
	fmt.Println("Go server Listening on: ", customHandlerPort)
	log.Fatal(http.ListenAndServe(":"+customHandlerPort, mux))
}

type pingResponse struct {
	Message  string    `json:"message"`
	DateTime time.Time `json:"dateTime"`
}

// pingHandler  provides a simple http get response
func pingHandler(w http.ResponseWriter, r *http.Request) {
	response := pingResponse{
		Message:  "Hello go function",
		DateTime: time.Now(),
	}
	fmt.Println(response)

	w.Header().Set("Content-Type", "application/json")
	responseJson, _ := json.Marshal(response)
	w.Write(responseJson)
}

type helloRequest struct {
	Name string `json:"name"`
}

type helloResponse struct {
	Message string `json:"message"`
}

// helloHandler provides a simple http post response with request body
func helloHandler(w http.ResponseWriter, r *http.Request) {

	name := "anon"
	if r.Method == "POST" {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			logger.Error("Error reading request body", zap.Error(err))
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		logger.Info("Request body", zap.String("body", string(body)))
		req := new(helloRequest)
		err = json.Unmarshal(body, &req)
		if err != nil {
			logger.Error("Error unmarshalling request body", zap.Error(err))
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		name = req.Name
	} else {
		name = r.URL.Query().Get("name")
	}

	response := helloResponse{
		Message: "Hello " + name,
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
