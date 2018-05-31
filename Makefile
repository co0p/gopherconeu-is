# back to the 80s Make is the new black
PROJECT?=github.com/co0p/gopherconeu-is
APP?=gopherconeu
PORT?=8080

GOARCH?=linux
GOOS?=amd64

RELEASE=0.0.0
COMMIT?=$(shell git rev-parse HEAD)
BUILD_TIME?=$(shell date --iso-8601=seconds)

clean:
	rm -f ./bin/${APP}

build: clean
	# CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} \ 
	go build \
	-ldflags "-s -w \
	-X ${PROJECT}/version.Release=${RELEASE} \
	-X ${PROJECT}/version.Commit=${COMMIT} \
	-X ${PROJECT}/version.BuildTime=${BUILD_TIME}" \
	-o ./bin/${APP} ${PROJECT}/cmd/gophercon
	

run: build
	PORT=${PORT} ./bin/${APP}

test:
	go test -race ./...