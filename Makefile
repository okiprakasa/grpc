gen:
	@protoc -I=proto --go_out=./ proto/*.proto

clean:
	@del pb\*.go

run:
	@go run .