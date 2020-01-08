package graphcall

import (
	"strings"
)

type BadSmellGraphCall struct {
}

var totalPath []string

func NewBadSmellGraphCall() *BadSmellGraphCall {
	return &BadSmellGraphCall{}
}

func (j *BadSmellGraphCall) AnalysisGraphCallPath(nodes map[string][]string) []string {
	for k := range nodes {
		getConnectedPath(k, nodes)
	}
	return totalPath;
}

func getConnectedPath(startNode string, nodes map[string][]string) {
	relatedNodes := nodes[startNode]
	var currentPath []string
	for i := range relatedNodes {
		for j := i + 1; j < len(relatedNodes); j++ {
			getPath(startNode, nodes, currentPath, relatedNodes[i], relatedNodes[j]);
			getPath(startNode, nodes, currentPath, relatedNodes[j], relatedNodes[i]);
		}
	}
}

func getPath(startNode string, nodes map[string][]string, currentPath []string, currentNode string, endNode string) {
	nextNodes := nodes[currentNode]
	if len(nextNodes) == 0 || currentNode == startNode || currentNode == endNode {
		return;
	}
	if contains(nextNodes, endNode) {
		path := strings.Join(append([]string{startNode}, append(currentPath, currentNode, endNode)...), "->")
		totalPath = append(totalPath, path+";"+startNode+"->"+endNode+"");
	}
	for _, node := range nextNodes {
		if contains(currentPath, node) {
			continue;
		}
		getPath(startNode, nodes, append(currentPath, currentNode), node, endNode);
	}
}

func contains(list []string, element string) bool {
	for _, e := range list {
		if e == element {
			return true
		}
	}
	return false
}
