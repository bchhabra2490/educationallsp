package main

import (
	"bufio"
	"educationallsp/analysis"
	"educationallsp/lsp"
	"educationallsp/rpc"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

func main(){
	fmt.Println("Hello, World!")

	logger := getLogger("educationallsp.log")
	logger.Println("Starting educationallsp")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	state := analysis.NewState()
	writer := os.Stdout
	for scanner.Scan() {
		msg := scanner.Bytes()
		method, content, err := rpc.DecodeMessage(msg)
		if err != nil {
			logger.Println("Error decoding message: ", err)
			continue
		}
		handleMessage(logger, writer, state, method, content)
	}
}

func handleMessage(logger *log.Logger, writer io.Writer, state analysis.State, method string, content []byte){
	logger.Println("Received message: ", method)
	switch method {
	case "initialize":
		var request lsp.InitializeRequest
		if err := json.Unmarshal(content, &request); err != nil {
			logger.Println("Error unmarshalling initialize request: ", err)
			return
		}
		logger.Printf("Connect to: %s %s", request.Params.ClientInfo.Name, request.Params.ClientInfo.Version)

		writeResponse(writer, lsp.NewInitializeResponse(request.ID))
	case "textDocument/didOpen":
		var notification lsp.DidOpenTextDocumentNotification
		if err := json.Unmarshal(content, &notification); err != nil {
			logger.Println("Error unmarshalling did open text document notification: ", err)
			return
		}
		state.OpenDocument(notification.Params.TextDocument.URI, notification.Params.TextDocument.Text)
	case "textDocument/didChange":
		var notification lsp.TextDocumentDidChangeNotification
		if err := json.Unmarshal(content, &notification); err != nil {
			logger.Println("Error unmarshalling did change text document notification: ", err)
			return
		}
		logger.Printf("Did change text document: %s %s", notification.Params.TextDocument.URI, notification.Params.ContentChanges)
		for _, contentChange := range notification.Params.ContentChanges {
			logger.Printf("Content change: %s", contentChange.Text)
			state.UpdateDocument(notification.Params.TextDocument.URI, contentChange.Text)
		}
	case "textDocument/hover":
		var request lsp.HoverRequest
		if err := json.Unmarshal(content, &request); err != nil {
			logger.Println("Error unmarshalling hover request: ", err)
			return
		}
		writeResponse(writer, state.Hover(logger, request.ID, request.Params.TextDocument.URI, request.Params.Position))
	case "shutdown":
		logger.Println("Shutting down")
	case "exit":
		logger.Println("Exiting")
	}
}

func writeResponse(writer io.Writer, msg any){
	reply := rpc.EncodeMessage(msg)
	writer.Write([]byte(reply))
}

func getLogger(filename string) *log.Logger{
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	return log.New(logfile, "[educationallsp] ", log.Ldate|log.Ltime|log.Lshortfile)
}