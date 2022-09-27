dockerup:
	docker-compose up -d --build --scale worker=4
dockerstop:
	docker-compose stop