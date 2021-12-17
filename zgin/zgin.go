package zgin

import (
	"net/http"
	"strings"
)

type HandlerFunc func(c *Context)

type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc
	engine      *Engine
}

type Engine struct {
	*RouterGroup
	router *router
	groups []*RouterGroup
}

func New() *Engine {
  engine := &Engine{
		router: newRouter(),
	}
  engine.RouterGroup = &RouterGroup{
    engine: engine,
  }
  engine.groups = []*RouterGroup{engine.RouterGroup}
  return engine
}

func (group *RouterGroup) Group(prefix string) *RouterGroup {
  newGroup := &RouterGroup{
    prefix: group.prefix + prefix,
    engine: group.engine,
  }
  group.engine.groups = append(group.engine.groups, newGroup)
  return newGroup
}

func (group *RouterGroup) addRoute(method, pattern string, handler HandlerFunc) {
	group.engine.router.addRoute(method, group.prefix + pattern, handler)
}

func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute(http.MethodGet, pattern, handler)
}

func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute(http.MethodPost, pattern, handler)
}

func (group *RouterGroup) Use(middlewares ...HandlerFunc) {
  group.middlewares = append(group.middlewares, middlewares...)
}

func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
  middlewares := make([]HandlerFunc, 0)
  for _, g := range engine.groups {
    if strings.HasPrefix(req.URL.Path, g.prefix) {
      middlewares = append(middlewares, g.middlewares...)
    }
  }
	c := newContext(w, req)
  c.handlers = middlewares
	engine.router.handle(c)
}
