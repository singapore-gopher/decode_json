package stream

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"testing"
)

func TestNewNormalValue(t *testing.T) {
	val, err := newNormalValue()
	if err != nil || len(val) == 0 {
		t.Fail()
		return
	}
}

func TestNewMagicValue(t *testing.T) {
	val, err := newMagicValue()
	if err != nil || len(val) == 0 {
		t.Fail()
		return
	}

}

func TestNewPacket(t *testing.T) {
	data, err := newPacket(false)
	if err != nil || len(data) == 0 {
		t.Fail()
		return
	}

	p := &Packet{}
	err = json.Unmarshal(data, p)
	if err != nil ||
		p.TimeStamp == 0 ||
		len(p.Value) == 0 {
		t.Fatal(fmt.Errorf("Failed to create a valid packet"))
		return
	}

	if IsMagicValue(p.Value) {
		t.Fail()
	}
}

func TestEncryptDecrypt(t *testing.T) {
	type text struct {
		clear     string
		encrypted string // base64 []byte
	}

	testData := []*text{
		&text{clear: "hello world"},
		&text{clear: "hello world"},
	}

	for _, txt := range testData {
		encrypted, err := encrypt(txt.clear)
		if err != nil {
			t.Fatal(err)
			return
		}

		txt.encrypted = base64.StdEncoding.EncodeToString(encrypted)
	}

	// for an extra check (are encrypted values really unique?)
	unique := map[string]interface{}{}

	for _, txt := range testData {
		unique[txt.encrypted] = nil

		data, err := base64.StdEncoding.DecodeString(txt.encrypted)
		if err != nil {
			t.Fatal(err)
			return
		}

		decrypted, err := decrypt(data)
		if err != nil {
			t.Fatal(err)
			return
		}

		if decrypted != txt.clear {
			t.Fatal(fmt.Errorf(
				"Failed to decrypt. Expected:%s Got:%s",
				txt.clear, decrypted))
		}
	}

	if len(unique) != len(testData) {
		t.Fail()
		return
	}
}

func TestNewMagicPacket(t *testing.T) {
	data, err := newPacket(true)
	if err != nil || len(data) == 0 {
		t.Fail()
		return
	}

	p := &Packet{}
	err = json.Unmarshal(data, p)
	if err != nil ||
		p.TimeStamp == 0 ||
		len(p.Value) == 0 {
		t.Fail()
		return
	}

	if !IsMagicValue(p.Value) {
		t.Fail()
	}

}
