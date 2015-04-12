package stream

import (
	"io"
	"log"
	"net"
	"sync"
	"time"
)

var (
	Port = "4001"
)

// for this json-challenge this would do the job
// for a more flexible solution, look at channels
type router struct {
	sync.RWMutex
	clients map[int64]io.ReadWriteCloser
}

// add new client to router
func (self *router) add(c io.ReadWriteCloser) {
	id := time.Now().UTC().UnixNano()
	log.Println("new client:", id)
	self.Lock()
	self.clients[id] = c
	log.Println("clients:", len(self.clients))
	self.Unlock()
}

// delete a client (called on write errors)
func (self *router) remove(id int64) {
	self.Lock()
	log.Println("client disconnected:", id)
	delete(self.clients, id)
	log.Println("clients:", len(self.clients))
	self.Unlock()
}

// send to everyone
// does not wait for writes to finish
func (self *router) fanout(data []byte) {

	self.RLock()
	clients := self.clients // copy the clients so we can unlock the map
	self.RUnlock()

	for clientId, c := range clients {
		go func(clientId int64, data []byte, c io.ReadWriteCloser) {

			_, err := c.Write(data)
			if err != nil {
				self.remove(clientId)
				log.Println(clientId, err.Error())
				c.Close()
				return
			}

		}(clientId, data, c)
	}

	return

}

// Serve creates a router instance, spawns the goroutine for sending out the packets
// then listens for new incoming connections.
func Serve(status chan int) {

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	r := &router{
		clients: map[int64]io.ReadWriteCloser{},
	}

	// sending out the goodies.
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

	// ready
	status <- 0

	l, err := net.Listen("tcp", ":"+Port)
	if err != nil {
		log.Fatal(err)
	}

	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		log.Println("connection accepted from IP:", conn.RemoteAddr())

		r.add(conn)
	}
}
