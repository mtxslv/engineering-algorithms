package utils

import (
	"io"
	"log"
	"os"
    "regexp"
    "strings"
)

func LoadText(textPath string) string {
	// adapted from https://stackoverflow.com/questions/36111777/how-to-read-a-text-file
	// and from https://stackoverflow.com/questions/9644139/from-io-reader-to-string-in-go
    file, err := os.Open(textPath)
    if err != nil {
        log.Fatal(err)
    }
    defer func() {
        if err = file.Close(); err != nil {
            log.Fatal(err)
        }
    }()
  b, err := io.ReadAll(file)
  text := string(b)
  return text
}

func BreakTextInNewLines(text string) []string {
    splittedString := strings.Split(text,"\n")
    var lines []string
    for _, splitted := range splittedString {
        // Remove four-whitespaces prefix
        splitted = strings.TrimPrefix(splitted, "    ")
        // Remove empty characters
        if len(splitted) > 0 {
            lines = append(lines, splitted)
        }
    }
    return lines
}

func LoadGraphDefinition(graphPath string) []string {
    graphCode := LoadText(graphPath)
    commands := BreakTextInNewLines(graphCode)
    return commands
}

func ExtractNodesAndEdges(graphCommands []string) ([]nodeNameAndLabel, []graphEdge) {

    var nodes []nodeNameAndLabel
    var edges []graphEdge

    var reNodeDefinition = regexp.MustCompile(`(node\d+) \[label="(.+)"];`)
    var reEdgeDefinition = regexp.MustCompile(`(node\d+) -> (node\d+);`)

    // Check if command is a Node Definition
    for _, command := range graphCommands {
        
        var purpotedNodeDef = reNodeDefinition.FindStringSubmatch(command)

        if len(purpotedNodeDef) > 0 {
            node := nodeNameAndLabel{
                nodeName: purpotedNodeDef[1],
                nodeLabel: purpotedNodeDef[2],
            }
            nodes = append(nodes, node)
        } else {
            var purpotedEdgeDef = reEdgeDefinition.FindStringSubmatch(command)
            if len(purpotedEdgeDef) > 0 {
                edge := graphEdge{
                    nodeNameOrigin: purpotedEdgeDef[1],
                    nodeNameDestiny: purpotedEdgeDef[2],
                }
                edges = append(edges, edge)
            } 
        }
    }

    return nodes, edges
}