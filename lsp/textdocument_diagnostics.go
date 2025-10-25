package lsp

type PublishDiagnosticsNotification struct{
	Notification
	Params PublishDiagnosticsParams `json:"params"`
}

type PublishDiagnosticsParams struct{
	URI string `json:"uri"`
	Diagnostics []Diagnostic `json:"diagnostics"`
}

type Diagnostic struct{
	Range Range `json:"range"`
	Message string `json:"message"`
	Source string `json:"source"`
	Severity DiagnosticSeverity `json:"severity"`
}

type DiagnosticSeverity int