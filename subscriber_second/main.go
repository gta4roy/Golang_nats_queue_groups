package main

import (
	"log"
	"runtime"

	pb "nats_event_queuegroups/order"

	"github.com/golang/protobuf/proto"
	"github.com/nats-io/go-nats"
)

func main() {

	// Create server connection
	log.Println("Second Subscriber started ")
	natsConnection, err := nats.Connect(nats.DefaultURL)

	if err != nil {
		log.Println("Not able to Connect NATS Server " + nats.DefaultURL)
	} else {
		log.Println("Able to Connect NATS Server " + nats.DefaultURL)
	}

	natsConnection.QueueSubscribe("Discovery.GetSystemInfo", "worker-group", func(msgSecondRequest *nats.Msg) {
		var systemInfo pb.GetSystemTime
		err := proto.Unmarshal(msgSecondRequest.Data, &systemInfo)
		if err != nil {
			log.Fatalf("Error on unmarshal: %v", err)
		}

		log.Println(" System Info Details recieved by second subscriber:")
		log.Println("System Time ", systemInfo.Systemtime)
		log.Println("System Date ", systemInfo.Systemdate)
		log.Println("System Username ", systemInfo.Username)
		log.Println("System Server Ip ", systemInfo.Serverip)

	})

	// Keep the connection alive
	runtime.Goexit()
}
