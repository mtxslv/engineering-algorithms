package main

import (
    "fmt"
    "regexp"
)

type nodeNameAndLabel struct {
    nodeName string
    nodeLabel string
} 

func getNodeNameAndLabel(sample string) nodeNameAndLabel {
    var re = regexp.MustCompile(`(node\d+) \[label="(.+)"];`)
    var submatch = re.FindStringSubmatch(sample)

    return nodeNameAndLabel{
        nodeName: submatch[1],
        nodeLabel: submatch[2],
    }
}

func main(){
    var samples = []string{
        `node1 [label="add flour \n to bowl"];`,
        `node2 [label="indent \n center"];`,
        `node3 [label="add salt & \n baking powder"];`,
        `node4 [label="warm water"];`,
        `node5 [label="warm butter"];`,
        `node6 [label="add liquids \n to mix"];`,
        `node7 [label="pull together \n dough"];`,
        `node8 [label="knead dough"];`,
        `node9 [label="15 min rest"];`,
        `node10 [label="Divide \n in pieces"];`,
        `node11 [label="Roll \n into disks"];`,
        `node12 [label="Cook \n each"];`, 
    };

    // Regex stuff
    for _, sample := range samples {
        getNodeNameAndLabel(sample)
        result := getNodeNameAndLabel(sample)
        fmt.Printf("Node Name: %s, Node Label: %s\n", result.nodeName, result.nodeLabel)
    }
}