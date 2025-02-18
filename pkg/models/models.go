package models

// InvokeRequest is the request object for the function invocation
//
// Structure for handling requests from Azure Functions workers.
//
// Azure Functions の worker からのリクエストを処理するための構造体
// .
type InvokeRequest struct {
	Data     map[string]interface{}
	Metadata map[string]interface{}
}

// InvokeResponse is the response object for the function invocation
//
// Structure for handling responses to Azure Functions workers.
//
// Azure Functions の worker へのレスポンスを処理するための構造体
//
type InvokeResponse struct {
	Outputs     map[string]interface{}
	Logs        []string
	ReturnValue interface{}
}
