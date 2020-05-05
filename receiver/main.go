package main

import(
	"log"
	"bytes"

	"github.com/pebbe/zmq4"
	"github.com/SayedAlesawy/seekur/drivers/tcp"
	"github.com/SayedAlesawy/seekur/rsa"
)

var logSign = "[Receiver]"
var ip = "127.0.0.1"
var port = "9902"

func main() {
	// Construct the communicaiton endpoint with receiver
	endpoint := tcp.BuildConnectionString(ip, port)

	// Acquire a request connection
	connection, err := tcp.NewConnection(zmq4.SUB, "")
	if err {
		log.Fatal(logSign, "Failed to init a reply socket")
	}

	// Connect to the receiver endpoint
	connection.Connect(endpoint)
	log.Println(logSign, "Listening...")

	// Receive private key
	priv, err := connection.RecvBytes(0)
	if err{
		log.Println(logSign, "Failed to receive message")
	}

	log.Println(logSign, "Received Private key")

	// Decode private key
	privKey, err := rsa.DecodePrivKey(priv)
	if err{
		log.Fatal(logSign, "Error while decoding private key")
	}

	// Receive the encrypted message
	recvMsg, err := connection.RecvBytes(0)
	if err{
		log.Println(logSign, "Failed to receive message")
	}

	// Decrypt the received message
	decyrptedMsg := rsa.Decrypt(&privKey, recvMsg)

	// Print plain text after decryption
	plainTextMsg := bytes.NewBuffer(decyrptedMsg.Bytes()).String()
	log.Println(logSign, "Received message: ", plainTextMsg)
}
