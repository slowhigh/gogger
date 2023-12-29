package main

import (
	"log"

	"github.com/Slowhigh/gogger/producer/infrastructure"
	"github.com/Slowhigh/gogger/producer/infrastructure/router"
	"github.com/Slowhigh/gogger/producer/internal/adapter/controller/http"
	"github.com/memphisdev/memphis.go"
)

func main() {
	conn, err := memphis.Connect("localhost", "producer_1", memphis.Password("#B8T2oA-mZ"), memphis.AccountId(1))
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	producer, err := infrastructure.NewMessageProducer(conn)
	if err != nil {
		log.Fatalln(err)
	}

	controller := http.NewController(producer)

	router := router.NewRouter(controller)

	if err := router.Serve(); err != nil {
		log.Fatalln()
	}
}
