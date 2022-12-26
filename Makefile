dev:
	go run main.go

test:
	go test -coverprofile=cover.out ./...

coverage:
	go tool cover -html=cover.out
