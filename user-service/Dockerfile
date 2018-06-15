FROM golang:1.9.0 as builder

WORKDIR /go/src/github.com/adrenalinchik/mikra/user-service

COPY . .

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep init && dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build -o user-service -a -installsuffix cgo main.go database.go handler.go repository.go


FROM ubuntu:16.04

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/adrenalinchik/mikra/user-service/user-service .

CMD ["./user-service"]