package zgin

import (
	"log"
	"net/http"
	"strings"
)

type router struct {
  roots map[string]*node
  handlers map[string]HandlerFunc
}

func newRouter() *router {
  return &router{
    roots: make(map[string]*node),
    handlers: make(map[string]HandlerFunc),
  }
}

func parsePattern(pattern string) []string {
  vs := strings.Split(pattern, "/")
  parts := make([]string, 0)
  for _, item := range vs {
    if item != "" {
      parts = append(parts, item)
      if item[0] == '*' {  // only one "*" is allowed
        break
      }
    }
  }
  return parts
}

func (r *router) addRoute(method, pattern string, handler HandlerFunc) {
  parts := parsePattern(pattern)
  if _, ok := r.roots[method]; !ok {
    r.roots[method] = &node{}
  }
  r.roots[method].insert(pattern, parts, 0)
  log.Printf("Route %4s - %s", method, pattern)
  key := method + "-" + pattern
  r.handlers[key] = handler
}

func (r *router) getRoute(method, path string) (*node, map[string]string) {
  root, ok := r.roots[method]
  if !ok {
    return nil, nil
  }
  searchParts := parsePattern(path)
  n := root.search(searchParts, 0)
  if n == nil {
    return nil, nil
  }
  params := make(map[string]string)
  parts := parsePattern(n.pattern)
  for index, part := range parts {
    if part[0] == ':' {
      params[part[1:]] = searchParts[index]
    }
    if part[0] == '*' && len(part) > 1 {
      params[part[1:]] = strings.Join(searchParts[index:], "/")
    }
  }
  return n, params
}

func (r *router) handle(c *Context) {
  n, params := r.getRoute(c.Method, c.Path)
  if n != nil {
    c.Params = params
    key := c.Method + "-" + n.pattern
    r.handlers[key](c)
  } else {
    c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
  }
}
