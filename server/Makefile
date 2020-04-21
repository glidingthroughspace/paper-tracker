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
	mockgen -source=managers/export_manager.go -package=mock -destination=mock/export_manager.go
	mockgen -source=managers/learning_manager.go -package=mock -destination=mock/learning_manager.go
	mockgen -source=managers/room_manager.go -package=mock -destination=mock/room_manager.go
	mockgen -source=managers/tracker_manager.go -package=mock -destination=mock/tracker_manager.go
	mockgen -source=managers/tracking_manager.go -package=mock -destination=mock/tracking_manager.go
	mockgen -source=managers/workflow_exec_manager.go -package=mock -destination=mock/workflow_exec_manager.go
	mockgen -source=managers/workflow_template_manager.go -package=mock -destination=mock/workflow_template_manager.go
	mockgen -source=repositories/tracker_repository.go -package=mock -destination=mock/tracker_repository_mock.go
	mockgen -source=repositories/scan_result_repository.go -package=mock -destination=mock/scan_result_repository_mock.go
	mockgen -source=repositories/room_repository.go -package=mock -destination=mock/room_repository_mock.go
	mockgen -source=repositories/workflow_template_repository.go -package=mock -destination=mock/workflow_template_repository_mock.go
	mockgen -source=repositories/workflow_exec_repository.go -package=mock -destination=mock/workflow_exec_repository_mock.go

new-tracker:
	- @coap-client -m post coap://localhost:5688/tracker/new -o tracker.tmp
	- @echo "Tracker ID: $$(cat tracker.tmp | cbor-diag | jq '.id')"
	- @rm tracker.tmp