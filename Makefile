FUNCNAME = somefunctionsname

main:	main.go
	go build -o main main.go

run:	main
	cp ./host.json ./local.settings.json ./main Functions/
	cd Functions && func host start

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