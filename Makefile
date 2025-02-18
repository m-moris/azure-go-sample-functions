FUNCNAME = somefunctionsname
GO_FILES = $(shell find ./pkg -name '*.go') main.go

main:	$(GO_FILES)
	go build -o main

x:
	echo $(GO_FILES)
run:	main
	docker compose up -d
	cp ./host.json ./local.settings.json ./main Functions/
	cd Functions && func host start
	docker compose donw 

deploy:	main
	cp ./host.json ./local.settings.json ./main Functions/
	cd Functions && func azure functionapp publish $(FUNCNAME)
	
clean:
	rm -f ./main

test:
	@echo "\n-----"
	curl $(FUNCNAME).azurewebsites.net/api/ping
	@echo "\n-----"
	curl $(FUNCNAME).azurewebsites.net/api/hello?name=auzre	
	@echo "\n-----"
	curl $(FUNCNAME).azurewebsites.net/api/hello -X POST -H 'Content-Type: application/json' -d '{"name" : "azure2"} '


queue:
	az storage queue create --name input --account-name devstoreaccount1 --account-key 'Eby8vdM02xNOcqFlqUwJPLlmEtlCDXJ1OUzFT50uSRZ6IFsuFq2UVErCz4I6tq/K1SZFPTOtr/KBHBeksoGMGw==' --queue-endpoint http://127.0.0.1:10001/devstoreaccount1
	az storage queue create --name output --account-name devstoreaccount1 --account-key 'Eby8vdM02xNOcqFlqUwJPLlmEtlCDXJ1OUzFT50uSRZ6IFsuFq2UVErCz4I6tq/K1SZFPTOtr/KBHBeksoGMGw==' --queue-endpoint http://127.0.0.1:10001/devstoreaccount1
