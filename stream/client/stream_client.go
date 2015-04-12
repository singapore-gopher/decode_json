package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/singapore-gophers/decode_json/stream"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"sync"
)

type response struct {
	sync.RWMutex
	Values []string `json:"values"`
}

// for your convenience
func (self *response) addValue(value string) {
	self.Lock()
	defer self.Unlock()

	self.Values = append(self.values, value)
}

// again, for your convenience.
func (self *response) submit(url, teamName string) error {

	data, _ := json.Marshal(self)
	res, err := http.Post(url, "application/json", bytes.NewReader(data))
	if err != nil {
		return err
	}

	responseData, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return fmt.Errorf("Err %v - %s", res.StatusCode, string(responseData))
	}

	return nil
}

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
	// create stream.Packet objects from the TCP stream so you can examine
	// their value. hint: import "encoding/json"

	// For checking a packet's value use stream.IsMagicValue()

	// collect 3 magic values in this struct
	solution := &response{
		values: []string{},
	}

	// hint: you can use solution.addValue()

	url := ""                             // the server's url
	teamName := ""                        // your team's name
	err := solution.submit(url, teamName) //
	if err != nil {
		fmt.Println(err)
		return
	}

	// done
}
