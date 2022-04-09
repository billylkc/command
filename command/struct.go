package command

// Mermaid is a struct to save the mermaid formmated syntax
// Reference. https://mermaid-js.github.io/mermaid/#/usage?id=usage
type Mermaid struct {
	Content string
}

type MermaidExample struct {
	Examples []struct {
		Type    string `toml:"type"`
		Content string `toml:"content"`
	} `toml:"example"`
}
