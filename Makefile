build: # build and install the api binary
	go install ./...

cleanup:
	docker-compose -f development/docker-compose.yml down --remove-orphans
	docker system prune -f

reset:
	make cleanup; make dev;

server: # install and start up the api server
	go install ./... && api

cert:
	mkcert localhost
	mv localhost-key.pem key.pem
	mv localhost.pem certificate.pem
