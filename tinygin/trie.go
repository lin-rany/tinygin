package tinygin

import (
	"log"
	"strings"
)

type node struct {
	// whole path
	path string
	// split path
	part string
	// childnode
	child []*node
	// iswildcard
	iswild bool
}

func Parseparts(path string) []string {
	parts := strings.Split(path, "/")
	ans := make([]string, 0)
	for _, item := range parts {
		if item != "" {
			ans = append(ans, item)
			if item[0] == '*' {
				break
			}
		}
	}
	log.Printf("path:%v parts:%v ans:%v", path, parts, ans)
	return ans
}

func (n *node) Insert(path string, parts []string, height int) {
	if len(parts) == height {
		n.path = path
		return
	}
	part := parts[height]
	ch := n.MatchNext(part)
	if ch == nil {
		ch = &node{part: part, iswild: part[0] == ':' || part[0] == '*'}
		n.child = append(n.child, ch)
	}
	ch.Insert(path, parts, height+1)
}

func (n *node) MatchNext(part string) *node {
	for _, ch := range n.child {
		if ch.iswild || ch.part == part {
			return ch
		}
	}
	return nil
}
func (n *node) Search(path string, parts []string, height int) *node {
	if height == len(parts) || strings.HasPrefix(n.part, "*") {
		if n.path == "" {
			return nil
		}
		return n
	}
	part := parts[height]
	chs := n.SearchNext(part)
	for _, ch := range chs {
		if res := ch.Search(path, parts, height+1); res != nil {
			return res
		}
	}
	return nil
}
func (n *node) SearchNext(part string) []*node {
	ans := make([]*node, 0)
	for _, ch := range n.child {
		if ch.iswild || ch.part == part {
			ans = append(ans, ch)
		}
	}
	return ans
}
