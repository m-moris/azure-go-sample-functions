# Go Azure Functions 

[日本語 (Japanese)](./README.ja.md)

This sample executes Azure Functions in the Go language and is implemented using custom handlers for Azure Functions.

[Azure Functions custom handlers | Microsoft Docs](https://docs.microsoft.com/en-us/azure/azure-functions/functions-custom-handlers)

## Update 

### :new: 2025/02
- Update host.json
- Update go version to 1.23
- Update core tools version
- Add logging and error handling
- Add queue trigger


### Sample
| folder | Input         | Ouput                |
| ------ | ------------- | -------------------- |
| hello  | http trigger  | -                    |
| ping   | http trigger  | -                    |
| timer  | timer trigger |                      |
| queue  | queue trigger | queue output binding |


## Description

The folder structure is as follows

- `functions.json` is placed under `FunctionName/Functions`.
- `host.json` and `local.settings.json` are originally located in the root directory and are copied to `Functions` at runtime. The executable file `main` is also handled in the same way.
- HTTP triggers forward HTTP requests directly, so `enableForwardingHttpRequest` is set to `true`.
 

```sh
.
├── host.json
├── local.settings.json
├── main
├── hello
│   └── function.json
├── ping
│   └── function.json
├── queue
│   └── function.json
└── timer
    └── function.json
```

## Local execution

If `az` and `go` are installed, it can be run in a local environment. You can run it with `make run`.

```sh
$ make run
cp ./host.json ./local.settings.json ./main Functions/
cd Functions && func host start

Azure Functions Core Tools
Core Tools Version:       4.0.6610 Commit hash: N/A +0d55b5d7efe83d85d2b5c6e0b0a9c1b213e96256 (64-bit)
Function Runtime Version: 4.1036.1.23224

[2025-02-18T03:01:31.970Z] 2025-02-18T12:01:31.967+0900 INFO    azure-go-sample-functions/main.go:24    Start Go functions
[2025-02-18T03:01:31.972Z] 2025-02-18T12:01:31.967+0900 INFO    azure-go-sample-functions/main.go:27    FUNCTIONS_CUSTOMHANDLER_PORT: 46865
[2025-02-18T03:01:31.972Z] Go server Listening on:  46865
[2025-02-18T03:01:31.999Z] Worker process started and initialized.

Functions:

        hello: [POST] http://localhost:7071/api/hello

        ping: [GET] http://localhost:7071/api/ping

        queue: queueTrigger

        timer: timerTrigger

For detailed output, run func with --verbose flag.
```

## Deploy

Create Azure Functions with the following configuration

* Linux
* Custom handlers
* Name, region, etc. optional

 
 Rewrite `FUNCNAME` in the `Makefile` to the name you created. `az login` and deploy make deploy.

```sh
$ make deploy
cp ./host.json ./local.settings.json ./main Functions/
cd Functions && func azure functionapp publish somefunctionsname
Getting site publishing info...
Uploading package...
Uploading 3.29 MB [###############################################################################]
Upload completed successfully.
Deployment completed successfully.
Syncing triggers...
Functions in somefunctionsname:
    hello - [httpTrigger]
        Invoke url: https://somefunctionsname.azurewebsites.net/api/hello

    ping - [httpTrigger]
        Invoke url: https://somefunctionsname.azurewebsites.net/api/ping
```

When you `make test`, Azure Functions will be called with `curl`.

```
$ make test

-----
curl somefunctionsname.azurewebsites.net/api/ping
{"message":"Hello go function","dateTime":"2021-07-16T07:47:50.8568901Z"}
-----
curl somefunctionsname.azurewebsites.net/api/hello?name=auzre
{"message":"Hello auzre"}
-----
curl somefunctionsname.azurewebsites.net/api/hello -X POST -H 'Content-Type: application/json' -d '{"name" : "azure2"} '
{"message":"Hello azure2"}%
moris@mypc /work/go/functions/go-func-app2

```
## Trigger Input and Output

### HTTP

Since `enableForwardingHttpRequest` is set to `true`, HTTP requests are handled as they are. You can deserialize them into any type and serialize any type to return a response.

### Others

QueueTrigger and TimerTrigger receive HTTP requests and responses, but they are bound to the following types for processing. Note that the content varies for each trigger.

```go
type InvokeRequest struct {
    Data     map[string]interface{}
    Metadata map[string]interface{}
}

type InvokeResponse struct {
    Outputs     map[string]interface{}
    Logs        []string
    ReturnValue interface{}
}
```

For example, QueueTrigger receives the following JSON input, allowing you to access the queue message as needed.

```json
{
    "Data": {
      "queue": "\"hello!!\""
    },
    "Metadata": {
      "DequeueCount": "1",
      "ExpirationTime": "2025-02-25T03:06:19+00:00",
      "Id": "\"52d53c5f-e212-487e-998b-abbe3b9d51e4\"",
      "InsertionTime": "2025-02-18T03:06:19+00:00",
      "NextVisibleTime": "2025-02-18T03:16:20+00:00",
      "PopReceipt": "\"MThGZWIyMDI1MDM6MDY6MjA3NjM1\"",
      "sys": {
        "MethodName": "queue",
        "RandGuid": "e59556b5-892a-4465-a19d-4396f0978dab",
        "UtcNow": "2025-02-18T03:06:20.7562691Z"
      }
    }
}
```
