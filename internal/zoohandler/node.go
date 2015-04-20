package zoohandler

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/samuel/go-zookeeper/zk"
)

type RequestType byte

type Node struct {
	zh *ZooHandler

	name     string
	path     string
	fullpath string
}

func (node Node) Name() string {
	return node.name
}

func (node Node) Path() string {
	return node.path
}

func (node Node) FullPath() string {
	return node.fullpath
}

func (node Node) NumberName() (int64, bool) {
	i, err := strconv.ParseInt(node.name, 10, 63)
	if err != nil {
		return 0, false
	}
	return i, true
}

func (node Node) Root() bool {
	if node.fullpath == "/" {
		return true
	}
	return false
}

func (node Node) Parent() (Node, error) {
	if node.Root() {
		return Node{}, fmt.Errorf("Root node has no parent")
	}
	return node.Child(node.path)
}

func (node Node) Child(name string) (Node, error) {
	var fullpath string
	if name == "" || strings.ContainsRune(name, '/') {
		return Node{}, fmt.Errorf("Invalid child name (%s)", name)
	}

	if node.Root() {
		fullpath = "/" + name
	} else {
		fullpath = node.fullpath + "/" + name
	}
	child, err := node.zh.NewNode(fullpath)
	if err != nil {
		return Node{}, err
	}
	return child, nil
}

func (node *Node) Sequence() (int64, bool) {
	idx := strings.LastIndex(node.name, "-")
	if idx == -1 {
		return 0, false
	}
	seqstr := node.name[idx+1:]
	i, err := strconv.ParseInt(seqstr, 10, 63)
	if err != nil {
		return 0, false
	}
	return i, true
}

func (zh *ZooHandler) NewNode(fullpath string) (Node, error) {
	var name, path string
	if fullpath != "/" {
		idx := strings.LastIndex(fullpath, "/")
		if fullpath == "" || idx == -1 || fullpath[0] != '/' {
			return Node{}, fmt.Errorf("Path (%s) must start with '/'", fullpath)
		}
		name = fullpath[idx+1:]
		path = fullpath[:idx]
	}
	node := Node{
		zh:       zh,
		name:     name,
		path:     path,
		fullpath: fullpath,
	}
	return node, nil
}

func (nr NodeResult) Error() string {
	return fmt.Sprintf("Error result for %s operation on %s: %v", nr.Op, nr.Node.fullpath, nr.Err)
}

func (node Node) Get(setWatch bool) NodeResult {
	var data []byte
	var stat *zk.Stat
	var watch <-chan zk.Event
	var err error

	if setWatch {
		data, stat, watch, err = node.zh.conn.GetW(node.fullpath)
	} else {
		data, stat, err = node.zh.conn.Get(node.fullpath)
	}
	if err != nil {
		return NodeResult{
			Node: node,
			Op:   "get",
			Err:  err,
		}
	}

	nr := NodeResult{
		Node:  node,
		Op:    "get",
		Stat:  stat,
		Watch: watch,
		Data:  data,
	}
	return nr
}

func (node Node) GetChildren(setWatch bool) NodeResult {
	var children []string
	var stat *zk.Stat
	var watch <-chan zk.Event
	var err error

	if setWatch {
		children, stat, watch, err = node.zh.conn.ChildrenW(node.fullpath)
	} else {
		children, stat, err = node.zh.conn.Children(node.fullpath)
	}
	if err != nil {
		return NodeResult{
			Node: node,
			Op:   "children",
			Err:  err,
		}
	}

	nr := NodeResult{
		Node:     node,
		Op:       "children",
		Stat:     stat,
		Watch:    watch,
		Children: make([]Node, len(children)),
		ChildMap: make(map[string]Node, len(children)),
	}
	for i, name := range children {
		childnode, err := node.Child(name)
		if err != nil {
			return NodeResult{
				Node: node,
				Op:   "children",
				Err:  err,
			}
		}
		nr.Children[i] = childnode
		nr.ChildMap[name] = childnode
	}
	return nr
}

type NodeResult struct {
	Node
	Op       string
	Err      error
	Stat     *zk.Stat
	Watch    <-chan zk.Event
	Data     []byte
	Children []Node
	ChildMap map[string]Node
}

func (nr NodeResult) Json(obj interface{}) error {
	return json.Unmarshal(nr.Data, obj)
}
