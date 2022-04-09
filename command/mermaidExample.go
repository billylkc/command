package command

import (
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
)

// PrintMermaidExample prints the demo syntax for different type of graphs from the site
// Should be able to copy to content to the input.mmd and generate the graph
func PrintMermaidExample() error {

	s := `
[[example]]
type = "graph"
content = """
flowchart TD
    Start --> Stop
"""

[[example]]
type = "graph"
content = """
flowchart LR
    A-- text -->B
"""

[[example]]
type = "subgraph"
content = """
flowchart TB
    c1-->a2
    subgraph one
    a1-->a2
    end
    subgraph two
    b1-->b2
    end
    subgraph three
    c1-->c2
    end
"""

[[example]]
type = "decision"
content = """
flowchart LR
    A[Hard edge] -->|Link text| B(Round edge)
    B --> C{Decision}
    C -->|One| D[Result one]
    C -->|Two| E[Result two]
"""

`
	var mermaid MermaidExample
	if _, err := toml.Decode(s, &mermaid); err != nil {
		log.Fatal(err)
	}

	// Print example
	var (
		counter int    // Counter per group
		pType   string // Previous type
	)
	for _, e := range mermaid.Examples {
		if pType != e.Type { // Print type only if it's new
			fmt.Printf("[[%s]]\n", e.Type)
			fmt.Println("--------")
			pType = e.Type
			counter = 0
		} else {
		}
		counter += 1
		fmt.Printf("<%d>\n", counter)
		fmt.Println("---")
		fmt.Println(e.Content)
	}
	return nil
}
