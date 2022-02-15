package treemap

import "fmt"

//nolint:unused,deadcode
func graphvizNodes[Key, Value any](
	node *node[Key, Value],
) string {
	res := ""
	if node.isBlack {
		res += fmt.Sprintf("%v\n", node.key)
	} else {
		res += fmt.Sprintf("%v [fillcolor=lightpink]\n", node.key)
	}
	if node.left != nil {
		res += graphvizNodes(node.left)
	}
	if node.right != nil {
		res += graphvizNodes(node.right)
	}
	return res
}

//nolint:unused,deadcode
func graphvizEdges[Key, Value any](
	node *node[Key, Value],
) string {
	res := ""
	if node.left != nil {
		res += fmt.Sprintf("%v:sw -> %v\n", node.key, node.left.key)
		res += graphvizEdges(node.left)
	}
	if node.right != nil {
		res += fmt.Sprintf("%v:se -> %v\n", node.key, node.right.key)
		res += graphvizEdges(node.right)
	}
	return res
}

//nolint:unused,deadcode
//noinspection GoUnusedFunction
func graphviz[Key, Value any](
	node *node[Key, Value],
) string {
	return "digraph {\nnode [style=filled]\n" + graphvizNodes(node) + graphvizEdges(node) + "}\n"
}
