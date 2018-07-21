build: # build the docker containers
	docker-compose build

clean: # clean up the docker env
	docker-compose down

run: # build and start up the docker containers
	docker-compose build && docker-compose up