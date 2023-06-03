compose-up:
	docker-compose -f docker-compose/docker-compose.yml up --remove-orphans -d

compose-down:
	docker-compose -f docker-compose/docker-compose.yml down -v

tests:
	docker exec -it newman newman run hw09.postman_collection.json --verbose
