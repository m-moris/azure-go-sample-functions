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


## Description

The folder structure is as follows

* Place `functions.json` under `{FunctioName}/Functions` 
* The `host.json` and `local.settings.json` are the originals under root and are copied to `Functions` at runtime. The same goes for the executable `main`.
* In this case, enableForwardingHttpRequest is set to true since this is an HTTP trigger only.
 

```
[root]
│  host.json
│  local.settings.json
│  main.go
|  main
└─Functions
    │  host.json
    │  local.settings.json
    │  main
    │
    ├─hello
    │      function.json
    │
    └─ping
            function.json
```

## Local execution

If `az` and `go` are installed, it can be run in a local environment. You can run it with `make run`.

```sh
$ make run
go build -o main main.go
cp ./host.json ./local.settings.json ./main Functions/
cd Functions && func host start

Azure Functions Core Tools
Core Tools Version:       4.0.6610 Commit hash: N/A +0d55b5d7efe83d85d2b5c6e0b0a9c1b213e96256 (64-bit)
Function Runtime Version: 4.1036.1.23224

[2025-02-13T07:03:12.277Z] Go server Listening on:  41049
[2025-02-13T07:03:12.281Z] Worker process started and initialized.
[2025-02-13T07:03:12.282Z] {"level":"info","ts":1739430192.27376,"caller":"azure-go-sample-functions/main.go:22","msg":"Start Go functions"}
[2025-02-13T07:03:12.285Z] {"level":"info","ts":1739430192.2738538,"caller":"azure-go-sample-functions/main.go:25","msg":"FUNCTIONS_CUSTOMHANDLER_PORT: 41049"}

Functions:

        hello: [GET,POST] http://localhost:7071/api/hello

        ping: [GET] http://localhost:7071/api/ping

For detailed output, run func with --verbose flag.
[2025-02-13T07:03:17.258Z] Host lock lease acquired by instance ID '0000000000000000000000002BE192DD'.
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
