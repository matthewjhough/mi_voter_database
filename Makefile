.PHONY: setup init build clean pack upload deploy run ensure serve stop

DIR := ${CURDIR}

setup:
	docker run -d -p 5000:5000 --restart=always --name registry registry:2

init:
	docker run --rm -v ${DIR}/src:/go/src/skaioskit -w /go/src/skaioskit lushdigital/docker-golang-dep init

build:
	docker run --rm -v ${DIR}/src:/go/src/skaioskit -w /go/src/skaioskit lushdigital/docker-golang-dep ensure
	docker run --rm -v ${DIR}/src:/go/src/skaioskit -w /go/src/skaioskit golang:latest go build -ldflags "-linkmode external -extldflags -static" -o voter

clean:
	rm ${DIR}/src/voter

pack:
	docker build -f ./Dockerfile -t localhost:5000/skaioskit/voter-service .

upload:
	docker push localhost:5000/skaioskit/voter-service

deploy:
	envsubst < deployment.yaml | kubectl apply -f -

run:
	docker run -it -v ${DIR}/config:/etc/skaioskit -v ${DIR}/data:/data -p 8081:80 localhost:5000/skaioskit/voter-service /voter

ensure:
	docker run -it -v ${DIR}/config:/etc/skaioskit -v ${DIR}/data:/data -p 8081:80 localhost:5000/skaioskit/voter-service /voter ensure

serve:
	docker run -it -v ${DIR}/config:/etc/skaioskit -v ${DIR}/data:/data -p 8081:80 localhost:5000/skaioskit/voter-service /voter serve

stop:
	kubectl delete deployments,services,pods,pvc,cronjob --all
