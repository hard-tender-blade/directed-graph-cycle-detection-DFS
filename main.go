package main

// # directed graph cycle detection using DFS (depth first search)
func main() {
	graph := graph16NoCycle()

	visited := make(map[string]bool)
	recStack := make(map[string]bool)

	for node := range graph {
		if isCyclic(graph, node, visited, recStack) {
			println("Graph contains cycle ❌")
			return
		}
	}

	println("Graph does not contain cycle ✅")
}

func isCyclic(graph map[string][]string, node string, visited map[string]bool, recStack map[string]bool) bool {
	if recStack[node] {
		return true
	}

	if visited[node] {
		return false
	}

	visited[node] = true
	recStack[node] = true

	for _, neighbor := range graph[node] {
		if isCyclic(graph, neighbor, visited, recStack) {
			return true
		}
	}

	recStack[node] = false

	return false
}
