package zgin

import (
	"net/http"
)

type HandlerFunc func(c *Context)

type Engine struct {
  router *router
}

func New() *Engine {
  return &Engine{
    router: newRouter(),
  }
}

func (engine *Engine) addRoute(method, pattern string, handler HandlerFunc) {
  engine.router.addRoute(method, pattern, handler)
}

func (engine *Engine) GET(pattern string, handler HandlerFunc) {
  engine.addRoute(http.MethodGet, pattern, handler)
}

func (engine *Engine) POST(pattern string, handler HandlerFunc) {
  engine.addRoute(http.MethodPost, pattern, handler)
}

func (engine *Engine) Run(addr string) error {
  return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
  c := newContext(w, req)
  engine.router.handler(c)
}
