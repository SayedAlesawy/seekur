package main

import(
	"github.com/SayedAlesawy/seekur/drivers/tcp"
	"github.com/pebbe/zmq4"
	"log"
)

func main() {
	ip := "127.0.0.1"
	port := "9902"
	logSign := "[Sender]"

	// Construct the communicaiton endpoint with receiver
	endpoint := tcp.BuildConnectionString(ip, port)

	// Acquire a request connection 
	connection, err := tcp.NewConnection(zmq4.PUB, "")
	if err {
		log.Fatal(logSign, "Failed to init a request socket")
	}

	// Connect to the receiver endpoint
	connection.Bind(endpoint)

	for {
		log.Println(logSign, "Sending: Hello")

		err = connection.Send("Hello", 0)
		if err{
			log.Println(logSign, "Failed to send message")
		}
	}
}
