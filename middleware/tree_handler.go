package middleware

import (
	"go-web/common"
	"net/http"
	"strings"
)

type HandlerBasedOnTree struct {
	root *Node
}

type Node struct {
	path     string
	children []*Node
	// 如果这是叶子结点，匹配上就可以使用此方法
	handler HandlerFunc
}

// ServeHTTP 查找路由过程
func (h *HandlerBasedOnTree) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	handler, found := h.findRouter(request.URL.Path)
	if !found {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte("page is not found"))
		return
	}
	handler(common.NewContext(writer, request))
}

func (h *HandlerBasedOnTree) findRouter(path string) (HandlerFunc, bool) {
	paths := strings.Split(strings.Trim(path, "/"), "/")
	cur := h.root
	for _, path := range paths {
		matchChild, found := h.findMatchChild(cur, path)
		if !found {
			return nil, false
		}
		cur = matchChild
	}
	if cur.handler == nil {
		// 执行到这里是这种情况
		// 正确路由：/user/login
		// 访问的路由：/user
		return nil, false
	}
	return cur.handler, true
}

// Route 匹配路由过程
func (h *HandlerBasedOnTree) Route(method string, pattern string, handleFunc HandlerFunc) {
	// 去掉前后的"/"
	pattern = strings.Trim(pattern, "/")
	// 拿到url的各节点
	// eg：/user/login
	// [user,login]
	paths := strings.Split(pattern, "/")

	cur := h.root
	for index, path := range paths {
		if mathChild, ok := h.findMatchChild(cur, path); ok {
			cur = mathChild
		} else {
			// 若未找到这个节点，则创建路径之后的所有节点
			h.createSubTree(cur, paths[index:], handleFunc)
			return
		}
	}
}

func (h *HandlerBasedOnTree) createSubTree(root *Node, paths []string, handlerFunc HandlerFunc) {
	cur := root
	for _, path := range paths {
		nn := NewNode(path)
		cur.children = append(cur.children, nn)
		cur = nn
	}
	cur.handler = handlerFunc
}

func NewNode(path string) *Node {
	return &Node{
		path:     path,
		children: make([]*Node, 0, 2),
	}
}

func (h *HandlerBasedOnTree) findMatchChild(root *Node, path string) (*Node, bool) {
	for _, child := range root.children {
		if child.path == path {
			return child, true
		}
	}
	return nil, false
}
