sources = $(wildcard **/*.go)

paper-tracker: $(sources)
	go build -o paper-tracker ./cmd/paper-tracker

run: paper-tracker
	./paper-tracker

test:
	go test ./...

clean:
	- rm paper-tracker