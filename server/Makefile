sources = $(wildcard **/*.go)

.PHONY: clean test mock paper-tracker

paper-tracker: $(sources)
	go build -o paper-tracker ./cmd/paper-tracker

run: paper-tracker
	./paper-tracker

test:
	go test ./...

clean:
	- rm paper-tracker
	- rm paper-tracker.db

mock:
	mockgen -source=repositories/tracker_repository.go -package=mock -destination=mock/tracker_repository_mock.go
	mockgen -source=repositories/command_repository.go -package=mock -destination=mock/command_repository_mock.go
	mockgen -source=repositories/scan_result_repository.go -package=mock -destination=mock/scan_result_repository_mock.go
	mockgen -source=repositories/room_repository.go -package=mock -destination=mock/room_repository_mock.go
	mockgen -source=repositories/workflow_repository.go -package=mock -destination=mock/workflow_repository_mock.go

new-tracker:
	- @coap-client -m post coap://localhost:5688/tracker/new -o tracker.tmp
	- @echo "Tracker ID: $$(cat tracker.tmp | cbor-diag | jq '.id')"
	- @rm tracker.tmp