FROM golang:latest as builder

WORKDIR /go/src/github.com/adrenalinchik/mikra/consignment-service

COPY . .

RUN go get
RUN CGO_ENABLED=0 GOOS=linux go build  -o consignment-service -a -installsuffix cgo main.go


FROM ubuntu:16.04

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/adrenalinchik/mikra/consignment-service/consignment-service ./

CMD ["./consignment-service"]
