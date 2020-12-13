FROM golang:1.15.2-alpine  AS build
WORKDIR /go/src/github.com/jensdriller/tekton-apiserver-build-and-push
ENV CGO_ENABLED=0
COPY go.* .
COPY *.go .
RUN go mod download \
  && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=build /go/src/github.com/jensdriller/tekton-apiserver-build-and-push/app .
CMD ["./app"]
