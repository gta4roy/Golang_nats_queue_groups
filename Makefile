all:
	compile, build

compile:
	protoc -I order/ order/order.proto --go_out=plugins=grpc:order

build:
	go build -o nats_event_queuegroups/publisher nats_event_queuegroups/publisher
	go build -o nats_event_queuegroups/subscriber_first nats_event_queuegroups/subscriber_first
	go build -o nats_event_queuegroups/subscriber_second nats_event_queuegroups/subscriber_second
	cp -r publisher/config nats_event_queuegroups/

