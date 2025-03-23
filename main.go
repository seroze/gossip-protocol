package main

import (
	"fmt"
	"gossip/internal/gossip"
)

func main() {
	fmt.Println("Main in gossip protocol")

	node := &gossip.Node{
		ID:    "node1",
		Peers: []string{"localhost:5001", "localhost:5002"}, // Other known peers
		Data:  map[string]string{"hello": "world"},
	}

	go node.Listen("5000") // Listen on port 5000
	go node.Gossip()       // Start gossiping

	select {} // Keep the program running
}
