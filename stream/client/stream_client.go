package main

import (
	"github.com/singapore-gophers/decode_json/stream"
	"io"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", stream.Port)
	if err != nil {
		log.Fatal(err)
	}

	err = handleConn(conn)
	if err != nil {
		log.Println(err)
	}

	// done
}

func handleConn(conn io.ReadWriteCloser) {
	// create stream.Packet objects from the stream

	// check the packet's value with stream.IsMagicValue()

	// collect 3 magic values in this
	reponse := []string{}

}
