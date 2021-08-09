ifndef $(GOPATH)
	GOPATH=$(shell go env GOPATH)
	export GOPATH
endif

up-server: build-docker-server
	docker network ls | grep tcppow > /dev/null || docker network create tcppow
	docker run --rm -p 3032:3032  --name tcp-pow-server --network=tcppow tcp-pow/server --server=:3032

up-client: build-docker-client
	docker run --network=tcppow tcp-pow/client --server=tcp-pow-server:3032

build-docker-server:
	docker build -t tcp-pow/server -f Dockerfile.server .

build-docker-client:
	docker build -t tcp-pow/client -f Dockerfile.client .