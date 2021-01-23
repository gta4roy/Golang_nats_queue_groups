package main

import (
	"log"

	pb "nats_event_queuegroups/order"

	"github.com/golang/protobuf/proto"
	"github.com/nats-io/go-nats"
	"github.com/spf13/viper"
)

func main() {
	// Create NATS server connection
	natsConnection, _ := nats.Connect(nats.DefaultURL)
	defer natsConnection.Close()
	log.Println("Connected to " + nats.DefaultURL)
	viper.SetConfigName("app")
	viper.AddConfigPath("config")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Config file not found")
	}

	SystemInfo := pb.GetSystemTime{}
	SystemInfo.Systemtime = viper.GetString("systeminfo.systemtime")
	SystemInfo.Systemdate = viper.GetString("systeminfo.systemdate")
	SystemInfo.Username = viper.GetString("systeminfo.username")
	SystemInfo.Serverip = viper.GetString("systeminfo.serverip")

	log.Println(SystemInfo)

	data1, err := proto.Marshal(&SystemInfo)

	if err == nil {
		log.Println("Messaged been published ")
		natsConnection.Publish("Discovery.GetSystemInfo", data1)

	}

}
