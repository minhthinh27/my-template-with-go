
up:
	docker-compose --env-file .env.develop build --no-cache && docker-compose --env-file .env.develop up -d && docker image prune -f

down:
	docker rmi tga-service

log:
	docker logs -f go-tga-server