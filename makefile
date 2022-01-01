watch:
	nodemon --exec "go run" main.go --signal SIGKILL
run:
	go run main.go
build:
	go build main.go