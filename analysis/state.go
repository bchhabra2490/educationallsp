package analysis

import (
	"educationallsp/lsp"
	"log"
	"strings"
)

type State struct {
	Documents map[string]string
}

func NewState() State {
	return State{Documents: map[string]string{}}
}

func (s *State) OpenDocument(uri, text string) {
	s.Documents[uri] = text
}

func (s *State) UpdateDocument(uri, text string) {
	s.Documents[uri] = text
}

func (s *State) Hover(logger *log.Logger,id int, uri string, position lsp.Position) lsp.HoverResponse{

	document, ok := s.Documents[uri]
	if !ok {
		return lsp.HoverResponse{
			Response: lsp.Response{
				RPC: "2.0",
				ID: &id,
			},
		}
	}

	lines := strings.Split(document, "\n")
	line := lines[position.Line]
	logger.Printf("Hovering over line: %s", line)
	explanation, err := ExplainCode(logger, line, document)
	if err != nil {
		return lsp.HoverResponse{
			Response: lsp.Response{
				RPC: "2.0",
				ID: &id,
			},
		}
	}
	return lsp.HoverResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID: &id,
		},
		Result: lsp.HoverResult{
			Contents: explanation,
		},
	}
}