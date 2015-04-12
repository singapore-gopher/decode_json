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
	TeamName string   `json:"team"`
	Values   []string `json:"values"`
}

// for your convenience
func (self *response) addValue(value string) {
	self.Lock()
	defer self.Unlock()

	self.Values = append(self.Values, value)
}

// again, for your convenience. Use this to submit your solution.
func (self *response) submit(url, teamName string) error {

	self.TeamName = teamName
	data, _ := json.Marshal(self)
	res, err := http.Post(url, "application/json", bytes.NewReader(data))
	if err != nil {
		return err
	}

	responseData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return fmt.Errorf("Err %v - %s", res.StatusCode, string(responseData))
	}

	fmt.Println(string(responseData))
	return nil
}

func main() {
	conn, err := net.Dial("tcp", ":"+stream.Port)
	if err != nil {
		log.Fatal(err)
	}

	err = handleConn(conn)
	if err != nil {
		log.Println(err)
	}

	// done
}

func handleConn(conn io.ReadWriteCloser) error {
	// create stream.Packet objects from the TCP stream so you can examine
	// their value. hint: import "encoding/json"

	// create the json stream decoder
	dec := json.NewDecoder(conn)

	// you will collect 3 magic values in this struct
	solution := &response{
		Values: []string{},
	}

	for {
		p := &stream.Packet{}
		err := dec.Decode(p)
		if err != nil {
			if err == io.EOF {
				return err
			}
			dec = json.NewDecoder(conn)
			continue
		}

		// For checking a packet's value use stream.IsMagicValue()
		if stream.IsMagicValue(p.Value) {
			fmt.Println("yeah!")

			// hint: you can use solution.addValue()
			solution.addValue(p.Value)
			if len(solution.Values) == 3 {
				break
			}
		}

	}

	url := "http://127.0.0.1" + ":4000/stage3/submit.json" // the server's url
	teamName := "test"                                     // your team's name
	err := solution.submit(url, teamName)                  //
	if err != nil {
		fmt.Println(err)
		return nil
	}

	// done
	return nil
}
