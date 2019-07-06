build: # build and install the api binary
	go install ./...

server: # install and start up the api server
	go install ./... && api

cert:
	mkcert localhost
	mv localhost-key.pem key.pem
	mv localhost.pem certificate.pem
