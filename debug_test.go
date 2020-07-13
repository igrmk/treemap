package treemap

import "fmt"

// nolint: megacheck
func graphvizNodes(node *node) string {
	res := ""
	if node.isBlack {
		res += fmt.Sprintf("%d\n", node.key)
	} else {
		res += fmt.Sprintf("%d [fillcolor=lightpink]\n", node.key)
	}
	if node.left != nil {
		res += graphvizNodes(node.left)
	}
	if node.right != nil {
		res += graphvizNodes(node.right)
	}
	return res
}

// nolint: megacheck
func graphvizEdges(node *node) string {
	res := ""
	if node.left != nil {
		res += fmt.Sprintf("%d:sw -> %d\n", node.key, node.left.key)
		res += graphvizEdges(node.left)
	}
	if node.right != nil {
		res += fmt.Sprintf("%d:se -> %d\n", node.key, node.right.key)
		res += graphvizEdges(node.right)
	}
	return res
}

// nolint: megacheck
//noinspection GoUnusedFunction
func graphviz(node *node) string {
	return "digraph {\nnode [style=filled]\n" + graphvizNodes(node) + graphvizEdges(node) + "}\n"
}
