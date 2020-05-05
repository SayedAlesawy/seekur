package main

import(
	"log"
	"time"

	"github.com/pebbe/zmq4"
	"github.com/SayedAlesawy/seekur/drivers/tcp"
	"github.com/SayedAlesawy/seekur/rsa"
)

var logSign = "[Sender]"
var ip = "127.0.0.1"
var port = "9902"

func main() {
	// Generate a key pair for the encyrption
	log.Println(logSign, "Generating key pair")
	pubKey, privKey, err := rsa.GenerateKeyPair(1024)
	if err {
		log.Fatal(logSign, "Failed to generate encryption keys")
	}

	// Construct the communicaiton endpoint with receiver
	endpoint := tcp.BuildConnectionString(ip, port)

	// Acquire a request connection
	connection, err := tcp.NewConnection(zmq4.PUB, "")
	if err {
		log.Fatal(logSign, "Failed to init a request socket")
	}

	// Connect to the receiver endpoint
	connection.Bind(endpoint)

	// Wait until conneciton is stable
	time.Sleep(1 * time.Second)

	// Do key exchange - assume secure key exchange
	log.Println(logSign, "Sending private key")

	encodedPrivKey, err := privKey.Encode()
	if err {
		log.Fatal(logSign, "Failed to serialize private key")
	}

	err = connection.Send(encodedPrivKey, 0)
	if err{
		log.Fatal(logSign, "Failed to send private key")
	}

	// Wait until keys are exchanged
	time.Sleep(1 * time.Second)

	//msg := []byte("I use 2 programming languages, Go and rails")
	msg := "I use 2 programming languages, Go and rails"

	// Encrypt the message before sending
	encryptedMsg := rsa.Encrypt(pubKey, []byte(msg))

	// Send encrypted message
	log.Println(logSign, "Sending message: ", msg)
	err = connection.Send(encryptedMsg.Bytes(), 0)
	if err{
		log.Fatal(logSign, "Failed to send message")
	}

	// Wait for message recveial
	time.Sleep(1 * time.Second)
}
