package framework

import (
	"strings"
)

type TreeNode struct {
	children []*TreeNode
	handler  func(ctx *MyContext)
	param    string
}

func Contructor() TreeNode {
	return TreeNode{
		param:    "",
		children: []*TreeNode{},
	}
}

func isGeneral(param string) bool {
	return strings.HasPrefix(param, ":")
}

func (this *TreeNode) Insert(pathname string, handler func(ctx *MyContext)) {
	node := this

	params := strings.Split(pathname, "/")

	for _, param := range params {
		child := node.findChild(param)

		if child == nil {
			child = &TreeNode{
				param:    param,
				children: []*TreeNode{},
			}
			node.children = append(node.children, child)
		}
		node = child
	}

	node.handler = handler
}

func (this *TreeNode) findChild(param string) *TreeNode {
	for _, child := range this.children {
		if child.param == param {
			return child
		}
	}
	return nil
}

func (this *TreeNode) Search(pathname string) func(ctx *MyContext) {
	params := strings.Split(pathname, "/")

	result := dfs(this, params)

	if result == nil {
		return nil
	}

	return result.handler
}

func dfs(node *TreeNode, params []string) *TreeNode {
	currentParam := params[0]
	isLastParam := len(params) == 1

	for _, child := range node.children {

		if isLastParam {
			if isGeneral(child.param) {
				return child
			}

			if child.param == currentParam {
				return child
			}

			continue
		}

		if !isGeneral(child.param) && child.param != currentParam {
			continue
		}

		result := dfs(child, params[1:])

		if result != nil {
			return result
		}
	}

	return nil
}
