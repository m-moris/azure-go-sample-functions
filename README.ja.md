# Go Azure Functions 

Go を Azure Functions で実行するサンプルで、カスタムハンドラを使って実装している。

[Azure Functions custom handlers | Microsoft Docs](https://docs.microsoft.com/en-us/azure/azure-functions/functions-custom-handlers)

## 更新

### :new: 2025/02

- host.json を更新
- go version を 1.23 に更新
- ロギングとエラーハンドリングを追加
- サンプル追加

## サンプル

| フォルダ | Input         | Ouput                |
| -------- | ------------- | -------------------- |
| hello    | http trigger  | -                    |
| ping     | http trigger  | -                    |
| timer    | timer trigger |                      |
| queue    | queue trigger | queue output binding |

### フォルダ構成

フォルダ構成は以下の通り

- `functions.json`は、`ファンクション名/Functions` 配下に配置
- `host.json` と `local.settings.json` は、ルート配下にあるものがオリジナルで、実行時に`Functions` へコピーされる。実行ファイル `main` も同様
- HTTPトリガーは、HTTPリクエストをそのまま転送するので `enableForwardingHttpRequest` を `true` にしてある

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

## ローカル実行

`az` と `go` がインストールされていれば、ローカル環境で実行できる。`make run` で実行可能。

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

## トリガの入出力

### HTTP

`enableForwardingHttpRequest` を `true` にしてあるので、http リクエストをそのまま扱う。任意の型にデシリアライズして、任意の型をシリアライズでレスポンスを返すことができる。

### その他

QueueTrigger や Timer Triggerは、HTTP リクエスト、レスポンスを受け取るが、以下の型にバインドして処理する。内容は各トリガで異なるので注意が必要。

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

例えば QueueTrigger は以下のような JSON が入力されので、適宜キューメッセージへアクセスできる。

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