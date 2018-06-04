FROM golang:latest as builder

WORKDIR /go/src/github.com/adrenalinchik/mikra/consignment-service

COPY . .

#RUN go get
#RUN CGO_ENABLED=0 GOOS=linux go build  -o consignment-service -a -installsuffix cgo main.go

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep init && dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .


FROM ubuntu:16.04

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/adrenalinchik/mikra/consignment-service/consignment-service ./

CMD ["./consignment-service"]
