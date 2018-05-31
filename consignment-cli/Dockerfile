FROM golang:1.9.0 as builder

WORKDIR /go/src/github.com/adrenalinchik/mikra/consignment-cli

COPY . .

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep init && dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .


FROM ubuntu:16.04

RUN mkdir -p /app
WORKDIR /app

ADD consignment.json /app/consignment.json
COPY --from=builder /go/src/github.com/adrenalinchik/mikra/consignment-cli .

ENTRYPOINT ["./consignment-cli", "consignment.json"]
CMD ["./consignment-cli"]