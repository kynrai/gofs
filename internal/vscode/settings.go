package vscode

type Settings map[string]any

type Gopls struct {
	FormattingLocal   string `json:"formatting.local"`
	FormattingGofumpt bool   `json:"formatting.gofumpt"`
}

func (s Settings) SetGopls(g Gopls) {
	s["gopls"] = g
}
