FROM golang:1.21.12

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod ./

COPY . .
RUN go build -v -o /usr/bin/k8s-keepalive ./...

EXPOSE 5000

CMD ["/usr/bin/k8s-keepalive"]
