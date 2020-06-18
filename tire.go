package tequila

type node struct {
	Path     string  //访问路径
	Part     string  //路径块
	children []*node //子节点
}

// 匹配part是否一致
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if part == child.Part {
			return child
		}
	}
	return nil
}

func (n *node) insert(path string, parts []string, height int) {
	// 最末节点添加全路径
	if len(parts) == height {
		n.Path = path
		return
	}

	// 获取下一次处理的数据
	part := parts[height]
	// 判断是否为空 为空则创建
	child := n.matchChild(part)

	if child == nil {
		child = &node{
			Part: part,
		}
		n.children = append(n.children, child)
	}

	n.insert(path, parts, height+1)
}

func (n *node) search(parts []string, height int) *node {
	// 最末节点添加全路径
	if len(parts) == height && n.Path != "" {
		return n
	}

	// 获取下一次处理的数据
	part := parts[height]
	// 判断是否为空 为空则返回nil
	child := n.matchChild(part)

	if child == nil {
		return nil
	}

	n.search(parts, height+1)
	return n
}
