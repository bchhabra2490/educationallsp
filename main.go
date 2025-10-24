package main

import (
	"bufio"
	"educationallsp/lsp"
	"educationallsp/rpc"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main(){
	fmt.Println("Hello, World!")

	logger := getLogger("educationallsp.log")
	logger.Println("Starting educationallsp")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)
	for scanner.Scan() {
		msg := scanner.Bytes()
		method, content, err := rpc.DecodeMessage(msg)
		if err != nil {
			logger.Println("Error decoding message: ", err)
			continue
		}
		handleMessage(logger, method, content)
	}
}

func handleMessage(logger *log.Logger, method string, content []byte){
	logger.Println("Received message: ", method)
	switch method {
	case "initialize":
		var request lsp.InitializeRequest
		if err := json.Unmarshal(content, &request); err != nil {
			logger.Println("Error unmarshalling initialize request: ", err)
			return
		}
		logger.Printf("Connect to: %s %s", request.Params.ClientInfo.Name, request.Params.ClientInfo.Version)

		response := lsp.NewInitializeResponse(request.ID)
		writer := os.Stdout
		writer.Write([]byte(rpc.EncodeMessage(response)))

		logger.Print("Sent the reply.")
	case "shutdown":
		logger.Println("Shutting down")
	case "exit":
		logger.Println("Exiting")
	}
}

func getLogger(filename string) *log.Logger{
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	return log.New(logfile, "[educationallsp] ", log.Ldate|log.Ltime|log.Lshortfile)
}