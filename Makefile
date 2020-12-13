TAG = 0.0.1

all: start

start:
	go run ./... > backend.log 2>&1 &

build-local:
	go build -o backend && chmod +x backend

build-arm:
	GOOS=linux GOARCH=arm CGO_ENABLED=0 go build -o backend && chmod +x backend

upload:
	docker build -t jensdriller/istio-test-backen:v${TAG} .
	docker push jensdriller/istio-test-backen:v${TAG}

.PHONY: start build