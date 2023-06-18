gen:
	 protoc --go_out=. --go_opt=paths=source_relative \
       --go-grpc_out=. --go-grpc_opt=paths=source_relative \
       proto/**/*.proto
server:
	go run cmd/server/main.go
client:
	go build -ldflags="-X 'main.Version=in-makefile' -X 'main.Date=$(shell date)'" -o cmd/client/bin/main cmd/client/main.go
	cmd/client/bin/main
