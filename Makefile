TOP_PATH = $(shell pwd) 

run:
	rm -rf $(shell pwd)/protoorder/*.go
	protoc --proto_path=$(shell pwd)/proto/ --micro_out=$(shell pwd)/protoorder --go_out=$(shell pwd)/protoorder $(shell pwd)/proto/*proto
	go run main.go
co:
	go run $(shell pwd)/client/cliord.go
