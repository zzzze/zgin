package zgin

import (
	"fmt"
	"strings"
)

type node struct {
	pattern  string
	part     string
	children []*node
	isWild   bool
}

func (n *node) exactMatchChild(part string) *node {
	for _, child := range n.children {
    if child.part == part {
			return child
		}
	}
	return nil
}

func (n *node) fuzzyMatchChild(part string) *node {
	for _, child := range n.children {
    if child.isWild {
			return child
		}
	}
	return nil
}

func getWildNode(nodes []*node) *node {
  for _, node := range nodes {
    if node.isWild {
      return node
    }
  }
  return nil
}

func isWild(part string) bool {
  return part[0] == ':' || part[0] == '*'
}

func (n *node) insert(pattern string, parts []string, height int) {
  if len(parts) == height {
    n.pattern = pattern
    return
  }
	part := parts[height]
	child := n.exactMatchChild(part)
	if child == nil {
    wildChild := getWildNode(n.children)
    if wildChild != nil && isWild(part) {
      panic(fmt.Sprintf("%q in %q conflict with %q", part, pattern, wildChild.part))
    }
		child = &node{part: part, isWild: isWild(part)}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

func (n *node) search(parts []string, height int) *node {
  if len(parts) == height || strings.HasPrefix(n.part, "*") {
    if n.pattern == "" {
      return nil
    }
    return n
  }
  part := parts[height]
  child := n.exactMatchChild(part)
  if child == nil {
    child = n.fuzzyMatchChild(part)
  }
  if child != nil {
    if result := child.search(parts, height+1); result != nil {
      return result
    }
  }
  return nil
}
