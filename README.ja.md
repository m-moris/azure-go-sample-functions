# Go Azure Functions 

Go を Azure Functions で実行するサンプルで、カスタムハンドラを使って実装している。

[Azure Functions custom handlers | Microsoft Docs](https://docs.microsoft.com/en-us/azure/azure-functions/functions-custom-handlers)


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
% make run
go build -o main main.go
cp ./host.json ./local.settings.json ./main Functions/
cd Functions && func host start

Azure Functions Core Tools
Core Tools Version:       3.0.3477 Commit hash: 5fbb9a76fc00e4168f2cc90d6ff0afe5373afc6d  (64-bit)
Function Runtime Version: 3.0.15584.0


Functions:

        hello: [GET,POST] http://localhost:7071/api/hello

        ping: [GET] http://localhost:7071/api/ping

For detailed output, run func with --verbose flag.
[2021-07-16T07:24:32.452Z] Start Go functions
[2021-07-16T07:24:32.452Z] FUNCTIONS_CUSTOMHANDLER_PORT: 34137
[2021-07-16T07:24:32.452Z] Go server Listening on:  34137
[2021-07-16T07:24:32.485Z] Worker process started and initialized
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
