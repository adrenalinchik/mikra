build:
#		protoc --go_out=plugins=micro:. /media/taras/work/GolangProjectsLinux/go/src/github.com/adrenalinchik/mikra/consignment-service/proto/consignment
		docker build -t consignment-service .

run:
		docker run -p 50051:50051 \
                -e MICRO_SERVER_ADDRESS=:50051 \
                -e MICRO_REGISTRY=mdns consignment-service
