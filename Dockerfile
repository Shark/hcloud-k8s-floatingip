FROM golang:1.12-alpine
RUN apk add --no-cache git
WORKDIR /go/src/app
COPY main.go .

RUN go get -d -v ./...
RUN go install -v ./...

FROM alpine:latest
RUN apk add --no-cache ca-certificates
COPY --from=0 /go/bin/app /usr/local/bin/hcloud-k8s-floatingip
CMD ["/usr/local/bin/hcloud-k8s-floatingip"]
