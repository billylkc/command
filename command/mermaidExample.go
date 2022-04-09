package command

import (
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
)

// PrintMermaidExample prints the demo syntax for different type of graphs from the site
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
	fmt.Println(PrettyPrint(mermaid))
	return nil
}
