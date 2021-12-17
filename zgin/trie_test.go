package zgin

import (
	"log"
	"strings"
	"testing"
)

func assertSuccess(root *node, search []string, expectedPattern string) {
  n := root.search(search, 0)
  if n == nil {
    log.Fatalf("can't find %q", strings.Join(search, "/"))
  }
  if n.pattern != expectedPattern {
    log.Fatalf("expect %q, got %q", expectedPattern, n.pattern)
  }
}

func TestInsertAndSearch(t *testing.T) {
  root := &node{}
  root.insert("/p/:name", []string{"p", ":name"}, 0)
  root.insert("/p/zhangsan", []string{"p", "zhangsan"}, 0)
  root.insert("/p/zhangsan/b", []string{"p", "zhangsan", "b"}, 0)
  root.insert("/p/:name/b", []string{"p", ":name", "b"}, 0)
  root.insert("/p/hello/:age", []string{"p", "hello", ":age"}, 0)
  assertSuccess(root, []string{"p", "zhangsan"}, "/p/zhangsan")
  assertSuccess(root, []string{"p", "lisi"}, "/p/:name")
  assertSuccess(root, []string{"p", "zhangsan", "b"}, "/p/zhangsan/b")
  assertSuccess(root, []string{"p", "wangwu", "b"}, "/p/:name/b")
  assertSuccess(root, []string{"p", "hello", "10"}, "/p/hello/:age")
}
