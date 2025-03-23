package gossip

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"time"
)

type ID string

type Node struct {
	ID    string            // id of the string
	Peers []string          // id of the peers
	Data  map[string]string //state Node should store
}

func (n *Node) Gossip() {
	for {
		time.Sleep(2 * time.Second) // gossip interval

		if len(n.Peers) == 0 {
			// if no peers then skip
			continue
		}

		peer := n.Peers[rand.Intn(len(n.Peers))]

		//Send own state and exchange others
		n.sendAndExchangeState(peer)
	}
}

func (n *Node) sendAndExchangeState(peer string) {
	// find ip address
	conn, err := net.Dial("udp", peer)
	if err != nil {
		fmt.Println("Failed to contact peer:", err)

		return
	}
	defer conn.Close()

	// Send own state
	data, _ := json.Marshal(n.Data)
	conn.Write(data)

	// Receive peer state
	buffer := make([]byte, 1024)
	nBytes, _ := conn.Read(buffer)
	var peerData map[string]string
	json.Unmarshal(buffer[:nBytes], peerData)

	// merge peer's data with self
	for k, v := range peerData {
		n.Data[k] = v
	}
}

func (n *Node) Listen(port string) {
	addr, _ := net.ResolveUDPAddr("udp", ":"+port)
	conn, _ := net.ListenUDP("udp", addr)

	defer conn.Close()
	// buffer := make([]byte)
}
