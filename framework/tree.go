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

func (t *TreeNode) Search(pathname string) func(rw http.ResponseWriter, r *http.Request) {
	node := t
	params := strings.Split(pathname, "/")

	for _, param := range params {
		child := node.findChild(param)
		if child == nil {
			return nil
		}
		node = child
	}
	return node.handler
}
