build:
	docker build -t go_fiber_app:latest ./go_fiber -f ./go_fiber/Dockerfile
up: build
	docker-compose up -d
build_up:
	docker-compose up -d --build
app_log: up
	docker logs -f fiber-app
db_log: up
	docker logs -f postgres-container
down:
	docker-compose down
