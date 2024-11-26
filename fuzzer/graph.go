package fuzzer

type NodeType int

const (
	Rule NodeType = iota
	MultiRule
)

type Node struct {
	Name  string
	edges []*Edge
	typee NodeType
	price int
}

type Edge struct {
	from  *Node
	to    *Node
	price int
}

func ConstructGraph(lang languageRules) []*Node {
	graph := make(map[string]*Node)
	terminalNodes := make([]*Node, 0)

	for key := range lang {
		graph[key] = &Node{Name: key, typee: Rule, price: 1000000}
	}

	for key, value := range lang {
		node := graph[key]
		for _, expr := range value {
			rules := split_rule(expr.output)

			if len(rules) == 0 {
				node.price = 0
				terminalNodes = append(terminalNodes, node)
				continue
			}

			// TODO: If the rule is a multi rule, we should create a new node
			// connect this node to the new node, and then connect the new node to all the rules
			// in the multi rule

			for _, token := range rules {
				node.edges = append(node.edges, &Edge{from: node, to: graph[token], price: 100000})
			}
		}
	}

	if len(terminalNodes) == 0 {
		panic("No start nodes found. Grammar has no terminals")
	}

	return terminalNodes
}

func split_rule(rule string) []string {
	output := make([]string, 0)

	// find all <> pairs, where no < or > is inside
	start := 0
	for i := 0; i < len(rule); i++ {
		if rule[i] == '<' {
			if i > 0 && rule[i-1] != '<' {
				start = i
			}
		} else if rule[i] == '>' {
			if i > 0 && rule[i-1] != '>' {
				output = append(output, rule[start:i+1])
			}
		}
	}

	return output
}

/*
func BFS(start *Node) []int {
	visited := make(map[int]bool)
	queue := list.New()
	result := []int{}

	queue.PushBack(start)
	visited[start.id] = true

	for queue.Len() > 0 {
		element := queue.Front()
		queue.Remove(element)
		node := element.Value.(*Node)
		result = append(result, node.id)

		for _, edge := range node.Edges {
			if !visited[edge.To.id] {
				queue.PushBack(edge.To)
				visited[edge.To.id] = true
			}
		}
	}

	return result
} */
