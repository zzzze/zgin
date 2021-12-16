package zgin

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct {
  // origin objects
	Writer     http.ResponseWriter
	Req        *http.Request
  // request info
	Path       string
	Method     string
  // response info
	StatusCode int
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
  return &Context{
    Writer: w,
    Req: req,
    Path: req.URL.Path,
    Method: req.Method,
  }
}

func (c *Context) PostForm(key string) string {
  return c.Req.FormValue(key)
}

func (c *Context) Query(key string) string {
  return c.Req.URL.Query().Get(key)
}

func (c *Context) Status(code int) {
  c.StatusCode = code
  c.Writer.WriteHeader(code)
}

func (c *Context) SetHeader(key string, value string) {
  c.Writer.Header().Set(key, value)
}

func (c *Context) String(code int, format string, values ...interface{}) {
  c.Status(code)
  c.SetHeader(http.CanonicalHeaderKey("Content-Type"), "text/plain")
  fmt.Fprintf(c.Writer, format, values...)
}

func (c *Context) JSON(code int, obj interface{}) {
  c.Status(code)
  c.SetHeader(http.CanonicalHeaderKey("Content-Type"), "application/json")
  encoder := json.NewEncoder(c.Writer)
  if err := encoder.Encode(obj); err != nil {
    http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
  }
}

func (c *Context) Data(code int, data []byte) {
  c.Status(code)
  c.Writer.Write(data)
}

func (c *Context) HTML(code int, html string) {
  c.Status(code)
  c.SetHeader(http.CanonicalHeaderKey("Content-Type"), "text/html")
  c.Writer.Write([]byte(html))
}
