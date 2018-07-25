.PHONY: init build clean pack upload ensure serve

DIR := ${CURDIR}

init:
	docker run --rm -v ${DIR}/src:/go/src/skaioskit -w /go/src/skaioskit lushdigital/docker-golang-dep init

build:
	docker run --rm -v ${DIR}/src:/go/src/skaioskit -w /go/src/skaioskit lushdigital/docker-golang-dep ensure
	docker run --rm -v ${DIR}/src:/go/src/skaioskit -w /go/src/skaioskit golang:latest go build -ldflags "-linkmode external -extldflags -static" -o voter

clean:
	rm ${DIR}/src/voter

pack:
	docker build -f ./Dockerfile -t skaioskit/voter-service .

run:
	docker run -it -v ${DIR}/data:/data -p 8081:80 -e APP_PORT_NUMBER="80" -e APP_MYSQL_CONN_STR="root:password@tcp(docker.for.mac.localhost)/voter?charset=utf8&parseTime=True&loc=Local" skaioskit/voter-service /voter

ensure:
	docker run -it -v ${DIR}/data:/data -p 8081:80 -e APP_PORT_NUMBER="80" -e APP_MYSQL_CONN_STR="root:password@tcp(docker.for.mac.localhost)/voter?charset=utf8&parseTime=True&loc=Local" skaioskit/voter-service /voter ensure

serve:
	docker run -it -v ${DIR}/data:/data -p 8081:80 -e APP_PORT_NUMBER="80" -e APP_MYSQL_CONN_STR="root:password@tcp(docker.for.mac.localhost)/voter?charset=utf8&parseTime=True&loc=Local" skaioskit/voter-service /voter serve
