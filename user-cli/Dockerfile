FROM golang:1.9.0 as builder

WORKDIR /go/src/github.com/adrenalinchik/mikra/user-cli

COPY . .

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep init && dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .


FROM ubuntu:16.04

RUN mkdir -p /app
WORKDIR /app

COPY --from=builder /go/src/github.com/adrenalinchik/mikra/user-cli/user-cli .

ENTRYPOINT ["./user-cli"]

CMD ["./user-cli"]