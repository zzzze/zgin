package zgin

import (
	"fmt"
	"reflect"
	"testing"
)

func newTestRoute() *router {
	r := newRouter()
	r.addRoute("GET", "/", nil)
	r.addRoute("GET", "/hello/:name", nil)
	r.addRoute("GET", "/hello/b/c", nil)
	r.addRoute("GET", "/hi/:name", nil)
	r.addRoute("GET", "/assets/:filepath", nil)
	return r
}

func TestParsePattern(t *testing.T) {
	ok := reflect.DeepEqual(parsePattern("/p/:name"), []string{"p", ":name"})
	ok = ok && reflect.DeepEqual(parsePattern("/p/*"), []string{"p", "*"})
	ok = ok && reflect.DeepEqual(parsePattern("/p/*name/*"), []string{"p", "*name"})
	if !ok {
		t.Fatal("test parsePattern failed")
	}
}

func TestGetRoute(t *testing.T) {
	r := newTestRoute()
	n, ps := r.getRoute("GET", "/hello/zhangsan")
	if n == nil {
		t.Fatal("nil shouldn't be returned")
	}
	if n.pattern != "/hello/:name" {
		t.Fatal("shoud match /hello/:name")
	}
	if ps["name"] != "zhangsan" {
		t.Fatal("name should be equal to 'zhangsan'")
	}
	fmt.Printf("matched path %s, params['name']: %s\n", n.pattern, ps["name"])
}
