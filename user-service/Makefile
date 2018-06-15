build:
#		protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/adrenalinchik/mikra/user-service/proto/user/user.proto
		docker build -t vessel-service .

run:
		docker run -p 50052:50051 \
				-e MICRO_SERVER_ADDRESS=:50051 \
				-e MICRO_REGISTRY=mdns user-service
