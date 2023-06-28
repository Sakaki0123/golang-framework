package framework

import (
	"net/http"
	"strings"
)

type TreeNode struct {
	children []*TreeNode
	handler  func(rw http.ResponseWriter, r *http.Request)
	param    string
}

func Contructor() TreeNode {
	return TreeNode{
		param:    "",
		children: []*TreeNode{},
	}
}

func (t *TreeNode) Insert(pathname string, handler func(rw http.ResponseWriter, r *http.Request)) {
	node := t
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

func (t *TreeNode) findChild(param string) *TreeNode {
	for _, child := range t.children {
		if child.param == param {
			return child
		}
	}
	return nil
}

// 再起的にpathを読むようにする
func (t *TreeNode) Search(pathname string) func(rw http.ResponseWriter, r *http.Request) {
	node := t
	params := strings.Split(pathname, "/")

	result := dfs(node, params)
	if result == nil {
		return nil
	}

	return node.handler
}

func dfs(node *TreeNode, params []string) *TreeNode {
	currentParams := params[0]
	isLastParam := len(params) == 1
	for _, child := range node.children {
		if isLastParam {
			if isGeneral(child.param) {
				return child
			}
			if child.param == currentParams {
				return child
			}
		}

		if !isGeneral(child.param) && child.param != currentParams {
			continue
		}
		result := dfs(child, params[1:])

		if result != nil {
			return result
		}
	}

	return nil
}
