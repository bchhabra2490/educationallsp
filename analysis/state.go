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

func getDiagnosticsForFile(text string) []lsp.Diagnostic {
	diagnostics := []lsp.Diagnostic{}
	for row, line := range strings.Split(text, "\n") {
		idx := strings.Index(line, "VS Code")
		if idx != -1 {
			diagnostics = append(diagnostics, lsp.Diagnostic{
				Range: LineRange(row, idx, idx + len("VS Code")),
				Message: "Please make sure we use good language in this code",
				Severity: 1,
				Source: "educationallsp",
			})
		}
	}
	return diagnostics
}

func (s *State) OpenDocument(uri, text string) []lsp.Diagnostic {
	s.Documents[uri] = text
	return getDiagnosticsForFile(text)
}

func (s *State) UpdateDocument(uri, text string) []lsp.Diagnostic {
	s.Documents[uri] = text
	return getDiagnosticsForFile(text)
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

func (s *State) Definition(logger *log.Logger, id int, uri string, position lsp.Position) lsp.DefinitionResponse{
	
	return lsp.DefinitionResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID: &id,
		},
		Result: lsp.Location{
			URI: uri,
			Range: lsp.Range{
				Start: lsp.Position{
					Line: position.Line - 1,
					Character: 0,
				},
				End: lsp.Position{
					Line: position.Line - 1,
					Character: 0,
				},
			},
		},
	}
}

func (s *State) CodeAction(logger *log.Logger, id int, uri string) lsp.TextDocumentCodeActionResponse{
	text := s.Documents[uri]

	actions := []lsp.CodeAction{}
	for row, line := range strings.Split(text, "\n") {
		idx := strings.Index(line, "VS Code")
		if idx != -1 {
			replaceChange := map[string][]lsp.TextEdit{}
			replaceChange[uri] = []lsp.TextEdit{
				{
					Range: LineRange(row, idx, idx + len("VS Code")),
					NewText: "Neovim",
				},
			}

			actions = append(actions, lsp.CodeAction{
				Title: "Replace VS Code with Neovim",
				Edit: &lsp.WorkspaceEdit{
					Changes: replaceChange,
				},
			})

			censorChange := map[string][]lsp.TextEdit{}
			censorChange[uri] = []lsp.TextEdit{
				{
					Range: LineRange(row, idx, idx + len("VS Code")),
					NewText: "********",
				},
			}
			
			actions = append(actions, lsp.CodeAction{
				Title: "Censor VS Code",
				Edit: &lsp.WorkspaceEdit{
					Changes: censorChange,
				},
			})
		}

	}

	return lsp.TextDocumentCodeActionResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID: &id,
		},
		Result: actions,
	}
}

func (s *State) Completion(logger *log.Logger, id int, uri string, position lsp.Position) lsp.CompletionResponse{
	return lsp.CompletionResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID: &id,
		},
		Result: lsp.CompletionResult{
			Items: []lsp.CompletionItem{
				{
					Label: "VS Code",
					Detail: "VS Code is a code editor",
					Documentation: "VS Code is a code editor",
				},
			},
		},
	}
}

func LineRange(row, start, end int) lsp.Range{
	return lsp.Range{
		Start: lsp.Position{
			Line: row,
			Character: start,
		},
		End: lsp.Position{
			Line: row,
			Character: end,
		},
	}
}