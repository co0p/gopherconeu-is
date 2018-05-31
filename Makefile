# back to the 80s Make is the new black
PROJECT?=github.com/co0p/gopherconeu-is
APP?=gopherconeu
PORT?=8080
INTERNAL_PORT?=8001

GOARCH?=amd64
GOOS?=linux

RELEASE=0.0.0
COMMIT?=$(shell git rev-parse HEAD)
BUILD_TIME?=$(shell date --iso-8601=seconds)

CONTAINER_IMAGE?=hub.docker.io/co0p/gopherconeu

clean:
	rm -f ./bin/${APP}

build: clean
	CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} \
	go build \
	-ldflags "-s -w \
	-X ${PROJECT}/version.Release=${RELEASE} \
	-X ${PROJECT}/version.Commit=${COMMIT} \
	-X ${PROJECT}/version.BuildTime=${BUILD_TIME}" \
	-o ./bin/${APP} ${PROJECT}/cmd/gophercon

run: build
	PORT=${PORT} INTERNAL_PORT=${INTERNAL_PORT} ./bin/${APP}

run: container
	docker stop ${CONTAINER_IMAGE}:${RELEASE} || true && docker rm ${CONTAINER_IMAGE}:${RELEASE} || true \
	docker run --name ${APP} -p ${PORT}:${PORT} -p ${INTERNAL_PORT}:${INTERNAL_PORT} --rm \
	-e "PORT=${PORT}" -e "INTERNAL_PORT=${INTERNAL_PORT}" \
	${CONTAINER_IMAGE}:${RELEASE}

test:
	go test -race ./...

container: build
	docker build -t ${CONTAINER_IMAGE}:${RELEASE} .

