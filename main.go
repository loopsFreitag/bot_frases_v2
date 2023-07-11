package main

import (
	"example/bot-frases/api"
	"flag"
	"fmt"
	"log"
)

func main() {
	listenAddr := flag.String("listenaddr", ":3000", "Server address")
	flag.Parse()

	server := api.NewServer(*listenAddr)
	fmt.Println("Server on port", *listenAddr)
	log.Fatal(server.Start())
}
