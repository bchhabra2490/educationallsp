package rpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

func EncodeMessage(msg any) string {
	content, err := json.Marshal(msg)
	if err != nil{
		panic(err)
	}

    return fmt.Sprintf("Content-Length: %d\r\n\r\n%s", len(content), content)
}

type BaseMessage struct {
	Method string `json:"method"`
}

func DecodeMessage(message []byte) (string, []byte, error) {
	header, content, found := bytes.Cut(message, []byte{'\r', '\n', '\r', '\n'})
	if !found {
		return "", nil, errors.New("did not find separator")
	}

	// Content-Length: <number>

	contentLengthBytes := header[len("Content-Length: "):]
	contentLength, err := strconv.Atoi(string(contentLengthBytes))
	if err != nil {
		return "", nil, errors.New("invalid content length")
	}

	if len(content) != contentLength {
		return "", nil, errors.New("content length mismatch")
	}

	var baseMessage BaseMessage
	if err := json.Unmarshal(content[:contentLength], &baseMessage); err != nil {
		return "", nil, err
	}

	return baseMessage.Method, content[:contentLength], nil
}

// type SplitFunc func(data []byte, atEOF bool) (advance int, token []byte, err error)
func Split(data []byte, _ bool) (advance int, token []byte, err error){
	header, content, found := bytes.Cut(data, []byte{'\r', '\n', '\r', '\n'})
	if !found {
		return 0, nil, nil
	}

	// Content-Length: <number>

	contentLengthBytes := header[len("Content-Length: "):]
	contentLength, err := strconv.Atoi(string(contentLengthBytes))
	if err != nil {
		return 0, nil, errors.New("invalid content length: " + err.Error())
	}

	if len(content) < contentLength {
		return 0, nil, nil
	}

	return len(header) + contentLength + 4, data[:len(header) + contentLength + 4], nil // 4 for \r\n\r\n
}