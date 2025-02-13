# Go Azure Functions 

Go を Azure Functions で実行するサンプルで、カスタムハンドラを使って実装している。

[Azure Functions custom handlers | Microsoft Docs](https://docs.microsoft.com/en-us/azure/azure-functions/functions-custom-handlers)

## 更新

### :new: 2025/02

- host.json を更新
- go version を 1.23 に更新
- ロギングとエラーハンドリングを追加

## 説明

フォルダ構成は以下の通り

* `functions.json`は、`ファンクション名/Functions` 配下に配置
* `host.json` と `local.settings.json` は、ルート配下にあるものがオリジナルで、実行時に`Functions` へコピーされる。実行ファイル `main` も同様
* 今回は HTTPトリガーのみなので、 `enableForwardingHttpRequest` を `true` にしてある
 

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

## ローカル実行

`az` と `go` がインストールされていれば、ローカル環境で実行できる。`make run` で実行可能。

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

## デプロイ

Azure Functions を以下の設定で作成する。

* Linux
* カスタムハンドラ
* 名前、リージョン等は任意
 
`Makefile` の `FUNCNAME` を、作成した名前に書き換える。`az login` してから`make deploy` デプロイする。

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

`make test` で `curl` を叩く。

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
