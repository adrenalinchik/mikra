FROM golang:latest as builder

WORKDIR /go/src/github.com/adrenalinchik/mikra/consignment-service

COPY . .

#RUN go get github.com/derekparker/delve/cmd/dlv
RUN go get -u github.com/golang/dep/cmd/dep

RUN dep init && dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build  -o consignment-service -a -installsuffix cgo main.go repository.go handler.go datastore.go

#EXPOSE 50051 40000

FROM ubuntu:16.04

RUN mkdir /app

WORKDIR /app

COPY --from=builder /go/src/github.com/adrenalinchik/mikra/consignment-service/consignment-service ./
#COPY --from=builder /go/bin/dlv /

#CMD ["/dlv", "--listen=:40000", "--headless=true", "--api-version=2", "exec", "./consignment-service"]
CMD ["./consignment-service"]
