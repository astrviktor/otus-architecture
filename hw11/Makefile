tests:
	go test -v ./...

compose-up:
	docker-compose -f docker-compose/docker-compose.yaml up --remove-orphans -d

compose-down:
	docker-compose -f docker-compose/docker-compose.yaml down -v
