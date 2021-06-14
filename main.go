package main

import (
    "fmt"
    "io/ioutil"
    "strings"
)

type node struct {
    deps []string
}

func findEffected(nodes map[string]node, effectedNode node, result *[]string) {
    for _, name := range effectedNode.deps {
        *result = append(*result, name)
        findEffected(nodes, nodes[name], result)
    }
}

func main() {
    data, err := ioutil.ReadFile("input.txt")
    if err != nil {
        panic(err)
    }

    nodes := map[string]node {}

    for _, line := range strings.Split(string(data), "\n") {
        leftRight := strings.Split(line, "->")
        if len(leftRight) != 2 {
            continue
        }

        left, right := strings.TrimSpace(leftRight[0]), strings.TrimSpace(leftRight[1])
        currentNode := nodes[left]
        currentNode.deps = append(currentNode.deps, right)
        nodes[left] = currentNode
    }
    var effectedNode string
    fmt.Println("name of effected node: ")
    fmt.Scanln(&effectedNode)
    result := []string {effectedNode}
    fmt.Println(result)
    findEffected(nodes, nodes[effectedNode], &result)
    fmt.Println(result)
}
