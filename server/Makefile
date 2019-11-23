sources = $(wildcard **/*.go)

paper-tracker: $(sources)
	go build -o paper-tracker ./cmd/paper-tracker

run: paper-tracker
	./paper-tracker

test:
	go test ./...

clean:
	- rm paper-tracker

mock:
	mockgen -source=repositories/tracker_repository.go -package=mock -destination=mock/tracker_repository_mock.go
	mockgen -source=repositories/command_repository.go -package=mock -destination=mock/command_repository_mock.go
	mockgen -source=repositories/scan_result_repository.go -package=mock -destination=mock/scan_result_repository_mock.go
	mockgen -source=repositories/room_repository.go -package=mock -destination=mock/room_repository_mock.go