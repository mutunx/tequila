package tequila

import "strings"

type node struct {
	Path     string  //访问路径
	Part     string  //路径块
	children []*node //子节点
	isWild   bool
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

// 搜索用 匹配多子节点 因为在搜索时可以遇到匹配到动态路由和相同地址块的
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, n := range n.children {
		if part == n.Part || n.isWild {
			nodes = append(nodes, n)
		}
	}
	return nodes
}

func (n *node) insert(path string, parts []string, height int) {
	// 最末节点添加全路径 遇到*开头则直接返回
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
			Part:   part,
			isWild: strings.HasPrefix(part, ":") || strings.HasPrefix(part, "*"), // 判断是否时动态路由进行赋值
		}
		n.children = append(n.children, child)
	}

	child.insert(path, parts, height+1)
}

/**
输入:地址块和高度
返回:节点||nil
查找地址块是否有匹配的节点,有则返回节点没有返回nil
*/
func (n *node) search(parts []string, height int) *node {
	// 递归方法 结束标志为
	// 1.地址块长度等于长度:表示已经到达最后位置
	// 2.n的地址不为空:在插入成功时再最后一个节点添加了请求地址其他节点的请求地址为空
	// 3.n.part部分为*开头 则返回全部
	if len(parts) == height || strings.HasPrefix(n.Part, "*") {
		if n.Path == "" {
			return nil
		}
		return n
	}

	// 获取下一次处理的数据
	part := parts[height]
	// 获取所有匹配的子节点 相同节点和动态节点
	children := n.matchChildren(part)
	// 逐个节点树开始找
	for _, child := range children {
		// 递归往下找 遇到 有子节点,子节点符合建结束标志 则返回节点否则返回nil
		r := child.search(parts, height+1)
		if r != nil {
			return r
		}
	}

	return nil
}
