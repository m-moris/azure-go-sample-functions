package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	fmt.Println("Start Go functions")
	customHandlerPort, exists := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT")
	if exists {
		fmt.Println("FUNCTIONS_CUSTOMHANDLER_PORT: " + customHandlerPort)
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

func helloHandler(w http.ResponseWriter, r *http.Request) {

	name := "anon"
	if r.Method == "POST" {
		body, _ := ioutil.ReadAll(r.Body)
		req := new(helloRequest)
		_ = json.Unmarshal(body, &req)
		name = req.Name
	} else {
		name = r.URL.Query().Get("name")
	}

	response := helloResponse{
		Message: "Hello " + name,
	}

	w.Header().Set("Content-Type", "application/json")
	responseJson, _ := json.Marshal(response)
	w.Write(responseJson)
}
