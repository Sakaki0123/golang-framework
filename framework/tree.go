package framework

import (
	"strings"
)

type TreeNode struct {
	children []*TreeNode
	handler  func(ctx *MyContext)
	param    string
	parent   *TreeNode
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
				parent:   node,
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

func (this *TreeNode) Search(pathname string) *TreeNode {
	params := strings.Split(pathname, "/")

	return dfs(this, params)
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

func (this *TreeNode) ParseParams(pathname string) map[string]string {
	node := this
	pathname = strings.TrimSuffix(pathname, "/")
	paramArr := strings.Split(pathname, "/")

	paramDicts := make(map[string]string)
	for i:=len(paramArr)-1;i>0;i-- {
		if isGeneral(node.param) {
			paramDicts[node.param] = paramArr[i]
		}
		node = node.parent
	}
	return paramDicts
}
