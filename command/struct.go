package command

// Mermaid is a struct to save the mermaid formmated syntax
// Reference. https://mermaid-js.github.io/mermaid/#/usage?id=usage
type Mermaid struct {
	Content string
}

// MermaidExample is the example mermaid markdown
type MermaidExample struct {
	Examples []struct {
		Type    string `toml:"type"`
		Content string `toml:"content"`
	} `toml:"example"`
}

// Furniture are the items from taobao
type Furniture struct {
	Items []struct {
		Type  string `toml:"type"`
		Image string `toml:"image"`
		URL   string `toml:"url"`
	} `toml:"items"`
}
