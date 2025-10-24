package rpc_test

import (
	"educationallsp/rpc"
	"testing"
)


type EncodingExample struct {
	Method string
}

func TestEncode(t *testing.T) {
	expected := "Content-Length: 15\r\n\r\n{\"Method\":\"hi\"}"
	example := EncodingExample{Method: "hi"}
	encoded := rpc.EncodeMessage(example)
	if encoded != expected {
		t.Errorf("Expected %s, got %s", expected, encoded)
	}
}

func TestDecode(t *testing.T) {
	message := []byte("Content-Length: 15\r\n\r\n{\"Method\":\"hi\"}")
	method, content, err := rpc.DecodeMessage(message)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if method != "hi" {
		t.Errorf("Expected method hi, got %s", method)
	}
	contentLength := len(content)
	if contentLength != 15 {
		t.Errorf("Expected content length 15, got %d", contentLength)
	}
}