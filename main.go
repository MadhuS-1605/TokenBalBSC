package main

import (
	"fmt"
	"log"
	"net/http"

	Handlers "TokenBalBSC/handler"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
)

func main() {
	// Create a client instance to connect to our providr
	client, err := ethclient.Dial("https://speedy-nodes-nyc.moralis.io/2abeb0fd5d1beb3b221a0f05/bsc/mainnet")

	if err != nil {
		fmt.Println(err)
	}

	// Create a mux router
	r := mux.NewRouter()

	// We will define a single endpoint
	r.Handle("/api/v2/bsc/{module}", Handlers.ClientHandler{client})
	log.Fatal(http.ListenAndServe(":8080", r))
}