package stream

import (
	"fmt"
	"io"
	"log"
	"net"
	"sync"
	"time"
)

var (
	clientId = 0
)

// for this json-challenge this would do the job
// for more flexible solution, look at channels
type router struct {
	sync.RWMutex
	clients map[int64]io.ReadWriteCloser
}

// add new client
func (self *router) add(c io.ReadWriteCloser) {
	id := time.Now().UTC().UnixNano()
	fmt.Println(id)
	self.Lock()
	self.clients[id] = c
	self.Unlock()
}

// delete a client (called on write errors)
func (self *router) remove(id int64) {
	self.Lock()
	delete(self.clients, id)
	self.Unlock()
}

// send to everyone
// waits for writes to finish
func (self *router) fanout(data []byte) {

	var wg sync.WaitGroup

	self.RLock()
	defer self.RUnlock()
	for clientId, c := range self.clients {
		wg.Add(1)
		go func(clientId int64, data []byte, c io.ReadWriteCloser) {
			defer wg.Done()

			_, err := c.Write(data)
			if err != nil {
				log.Println(err)
				c.Close()
				self.remove(clientId)
				return
			}

		}(clientId, data, c)
	}

	wg.Wait()
	return

}

type streamHandler struct {
	r *router
}

// Serve creates a router instance, spawns the goroutine for sending out the packets
// then listens for new incoming connections.
func Serve(status chan int) {

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	r := &router{
		clients: map[int64]io.ReadWriteCloser{},
	}

	go func() {
		magicTick := time.NewTicker(1017 * time.Millisecond)
		normalTick := time.NewTicker(57 * time.Millisecond)
		for {
			select {
			case <-magicTick.C:
				data, err := newPacket(true)
				//fmt.Println(">>>", string(data))
				if err != nil {
					log.Println(err)
				}

				r.fanout(data)
			case <-normalTick.C:
				data, err := newPacket(false)
				//fmt.Println(">>>", string(data))
				if err != nil {
					log.Println(err)
				}

				r.fanout(data)
			}
		}

	}()

	status <- 0

	l, err := net.Listen("tcp", ":2015")
	if err != nil {
		log.Fatal(err)
	}

	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		r.add(conn)
	}
}
