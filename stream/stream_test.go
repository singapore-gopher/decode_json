package stream

import (
	"bytes"
	"encoding/json"
	"io"
	"net"
	"sync"
	"testing"
	"time"
)

type fakeConn struct {
	buf *bytes.Buffer
}

func newFakeConn() *fakeConn {
	fc := &fakeConn{
		buf: bytes.NewBuffer([]byte{}),
	}

	return fc
}

// Close function, so fakeConn implements io.ReadWriteCloser
func (self *fakeConn) Close() error {
	return nil
}

func (self *fakeConn) Write(p []byte) (int, error) {
	return self.buf.Write(p)
}

func (self *fakeConn) Read(p []byte) (int, error) {
	return self.buf.Read(p)
}

func TestFakeConn(t *testing.T) {
	text := "test text"

	conn := newFakeConn()
	conn.Write([]byte(text))

	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		t.Fatal(err)
	}
	buf = buf[:n]

	if string(buf) != text {
		t.Fatalf("expected: '%s' got: '%s'",
			text, string(buf))
		return
	}

}

func TestFanout(t *testing.T) {
	r := &router{
		clients: map[int64]io.ReadWriteCloser{},
	}

	clientCount := 3

	for i := 0; i < clientCount; i++ {
		c := newFakeConn()
		r.add(c)
	}

	if len(r.clients) != clientCount {
		t.Fatal("Failed to create clients")
	}

	text := "test text"
	r.fanout([]byte(text))

	for _, client := range r.clients {
		b := make([]byte, 4096)
		n, err := client.Read(b)
		if err != nil {
			t.Fatal(err)
		}
		b = bytes.TrimSpace(b[:n])
		got := string(b)

		if got != text {
			t.Fatalf("expected: '%s' got: '%s'",
				text, string(b))
			return
		}

	}
}

func TestServe(t *testing.T) {
	serverStatus := make(chan int) // in case starting a server takes time
	go Serve(serverStatus)
	_ = <-serverStatus

	clientCount := 20
	clientData := map[int64]*bytes.Buffer{}
	var wg sync.WaitGroup

	for i := 0; i < clientCount; i++ {
		wg.Add(1)
		c, err := net.Dial("tcp", ":2015")
		if err != nil {
			t.Fatal(err)
			return
		}

		go func(c net.Conn) {
			id := time.Now().UTC().UnixNano()
			buf := bytes.NewBuffer([]byte{})
			defer c.Close()

			dec := json.NewDecoder(c)
			for i := 0; i < 100; i++ {
				p := json.RawMessage{}

				err := dec.Decode(&p)
				if err != nil {
					t.Fatal(err)
				}

				_, err = buf.Write(p)
				if err != nil {
					t.Fatal(err)
				}
			}

			clientData[id] = buf // not thread safe.. don't do this in prod. code ;)
			wg.Done()
		}(c)
	}

	wg.Wait()

	result := ""
	for _, data := range clientData {
		if len(data.Bytes()) == 0 {
			t.Fatal("empty result")
		}
		if result == "" {
			result = data.String()
		} else if data.String() != result {
			t.Fatal("client result doesn't match")
			return
		}
	}
}
