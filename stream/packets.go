package stream

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"time"
)

var (
	Secret     = "A very secret secret."
	MagicValue = "Preciousss."
)

// Packet is the structure of a message in the stream. TimeStamp is
// a Unix timestamp, Value is a Base64 encoded []byte.
type Packet struct {
	TimeStamp int64  `json:"timestamp"`
	Value     string `json:"value"`
}

func newPacket(magic bool) ([]byte, error) {
	var value string
	var err error

	if magic {
		value, err = newMagicValue()
		if err != nil {
			return nil, err
		}
	} else {
		value, err = newNormalValue()
		if err != nil {
			return nil, err
		}
	}

	p := &Packet{
		TimeStamp: time.Now().UTC().Unix(),
		Value:     value,
	}

	return json.Marshal(p)
}

func newMagicValue() (string, error) {
	data, err := encrypt(MagicValue)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(data), nil
}

func newNormalValue() (string, error) {
	data := make([]byte, len(MagicValue)+aes.BlockSize) //we want them the same size
	_, err := rand.Read(data)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(data), nil
}

// because encrypting things is fun
func encrypt(text string) ([]byte, error) {
	key := sha256.Sum256([]byte(Secret))
	block, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, err
	}

	cipherText := make([]byte, len(text)+aes.BlockSize)
	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	enc := cipher.NewCFBEncrypter(block, iv)
	enc.XORKeyStream(cipherText[aes.BlockSize:], []byte(text))
	return cipherText, nil
}

// decrypting things is even more fun
func decrypt(cipherData []byte) (string, error) {
	key := sha256.Sum256([]byte(Secret))
	block, err := aes.NewCipher(key[:])
	if err != nil {
		return "", err
	}

	if len(cipherData) < aes.BlockSize {
		return "", fmt.Errorf("Invalid text")
	}

	data := cipherData[aes.BlockSize:]
	iv := cipherData[:aes.BlockSize]

	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(data, data)

	return string(data), nil
}

// IsMagicValue takes a Base64 encoded string and checks if it is an
// a magic packet value. It doesn't return errors, in case something
// fails (e.g the value is not a base64 []byte), the result is false.
func IsMagicValue(value string) bool {
	data, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		return false
	}

	dataString, _ := decrypt(data)
	if err != nil {
		return false // you might want proper error handling ;)
	}

	if dataString == MagicValue {
		return true
	}

	return false
}
