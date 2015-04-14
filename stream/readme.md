# stream
--
    import "github.com/singapore-gophers/decode_json/stream"


### Variables
```go
var (
	Secret     = "A very secret secret."
	MagicValue = "Preciousss."
)
```

```go
var (
	Port = "2015"
)
```

#### func  IsMagicValue

```go
func IsMagicValue(value string) bool
```
IsMagicValue takes a Base64 encoded string and checks if it is a magic packet's value. It doesn't return error, in case something fails (e.g the value is not a base64 []byte), the result will be false.

#### func  Serve

```go
func Serve(status chan int)
```
Serve creates a router instance, spawns the goroutine for sending out the
packets then listens for new incoming connections.

#### type Packet

```go
type Packet struct {
	TimeStamp int64  `json:"timestamp"`
	Value     string `json:"value"`
}
```

Packet is the structure of a message in the stream. TimeStamp is a Unix
timestamp, Value is a Base64 encoded []byte.
