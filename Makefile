# back to the 80s Make is the new black
PROJECT?=github.com/co0p/gopherconeu-is
APP?=gopherconeu
PORT?=8080

GOARCH?=linux
GOOS?=amd64

clean:
	rm -f ./bin/${APP}

build: clean
	# CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} \ 
	go build \
	-o ./bin/${APP} ${PROJECT}/cmd/gophercon

run: build
	PORT=${PORT} ./bin/${APP}