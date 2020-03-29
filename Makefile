default: build

clean:
	rm -f ./esphomekit

build:
	CGO_ENABLED=0 \
	go build -a -o esphomekit *.go

build-arm:
	CGO_ENABLED=0 \
	GOARCH=arm \
	GOARM=5 \
	GOOS=linux \
	go build -ldflags "-X main.Version=${DRONE_TAG} -X main.BuildDate=$(date +"%d.%m.%Y")" -a -o esphomekit *.go

build-linux:
	CGO_ENABLED=0 \
	GOARCH=amd64 \
	GOOS=linux \
	go build -ldflags "-X main.Version=${DRONE_TAG} -X main.BuildDate=$(date +"%d.%m.%Y")" -a -o esphomekit *.go

test:
	go test ./
