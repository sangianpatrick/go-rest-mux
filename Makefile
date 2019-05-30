dev:
	go run main.go

test:
	go test -v -cover -covermode=atomic ./...

dep: 
	dep ensure

update:
	dep ensure -update

secret-key:
	./secret-key-gen.sh