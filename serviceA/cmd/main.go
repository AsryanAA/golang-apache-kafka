package main

import (
	"fmt"
	"os"
	"log"
	"encoding/json"
	"io/ioutil"
	//"models/config"
	//"context"
	//"time"

	//kafka "github.com/segmentio/kafka-go"
)

type Config struct {
	PortDev string `json:"port_dev"`
}

// bin/kafka-console-consumer.sh --topic payment_service_topic --from-beginning --bootstrap-server localhost:9092

//read config file function
func ReadConfig() (string, error) {
	config_json, err := os.Open("./config.json")
	defer config_json.Close()
	if err != nil {
		log.Fatal("Error while open config file", err)
		return "", err
	}

	options, err1 := ioutil.ReadAll(config_json)
	if err1 != nil {
		log.Fatal("Error while read config file", err)
		return "", err1
	}

	var config Config
	json.Unmarshal(options, &config)
	fmt.Println(config.PortDev)
	return config.PortDev, nil
	
}

func main() {
	/*
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "payment_service_topic", 0)

	if err != nil {
		fmt.Println(err)
	}

	conn.SetWriteDeadline(time.Now().Add(time.Second * 10))

	conn.WriteMessages(kafka.Message{Value: []byte("afka again")})
	*/
	port, err := ReadConfig()
	if err != nil {
		log.Fatal("Server port is empty", err)
	}
	
	http.HandleFunc("/api/v1/fund_transfer", controllers.AppRequest)
	err1 := http.ListenAndServe(port, nil)
	if err1 != nil {
		log.Fatal("Error while start server", err)
	}
}