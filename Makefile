echoserver:
	go run echo/server.go

echoclient:
	go run echo/client.go

commandserver:
	go run command/main.go cat

ws:
	go run serverclient/main.go