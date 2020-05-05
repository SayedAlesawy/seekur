package main

import(
	"github.com/SayedAlesawy/seekur/drivers/tcp"
	"github.com/pebbe/zmq4"
	"log"
)

func main() {
	ip := "127.0.0.1"
	port := "9902"
	logSign := "[Receiver]"

	// Construct the communicaiton endpoint with receiver
	endpoint := tcp.BuildConnectionString(ip, port)

	// Acquire a request connection 
	connection, err := tcp.NewConnection(zmq4.SUB, "")
	if err {
		log.Fatal(logSign, "Failed to init a reply socket")
	}

	// Connect to the receiver endpoint
	connection.Connect(endpoint)

	for {
		msg, err := connection.RecvString(0)

		if err{
			log.Println(logSign, "Failed to receive message")
			continue
		}

		log.Println(logSign, "Received: ", msg)
	}
}
