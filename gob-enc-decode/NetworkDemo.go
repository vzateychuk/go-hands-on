/*
In this example we encoded the data using the client types, sent it to the server, and
dumped out what the server decoded. In what we get back from the server, we can see
it's using different types, that the user has an ID but no name, and that Amount is a 32-bit
float pointer type.
*/

package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"gob-enc-decode/client"
	"gob-enc-decode/server"
	"io"
	"log"
)

func sendToServer(net io.Reader) (*server.TxServer, error) {
	tx := &server.TxServer{}   // variable to be the target for decoding
	dec := gob.NewDecoder(net) // decoder with the network as the source
	err := dec.Decode(tx)      // Decode and capture any errors
	return tx, err
}

func main() {
	// the dummy network, which is a buffer
	var net bytes.Buffer

	// dummy data using the client-side structs
	clientTx := &client.TxClient{
		ID: "123456789",
		User: &client.UserClient{
			ID:   "ABCDEF",
			Name: "James",
		},
		AccountFrom: "Bob",
		AccountTo:   "Jane",
		Amount:      9.99,
	}

	// Encode the data. The target for the encoded data is our dummy network
	enc := gob.NewEncoder(&net)
	if err := enc.Encode(clientTx); err != nil {
		log.Fatal("error encoding: ", err) // Check for errors and exit if any are found
	}

	// Send the data to the server
	serverTx, err := sendToServer(&net)
	if err != nil {
		log.Fatal("server error: ", err)
	}

	// Print the decoded data to the console
	fmt.Printf("%#v\n", serverTx)
}
