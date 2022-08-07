package tinygin

import (
	"fmt"
	"reflect"
	"testing"
)

func newTestRouter() *Router {
	r := NewRouter()
	r.AddRoute("GET", "/", nil)
	r.AddRoute("GET", "/hello/:name", nil)
	r.AddRoute("GET", "/hello/b/c", nil)
	r.AddRoute("GET", "/hi/:name", nil)
	r.AddRoute("GET", "/assets/*filepath", nil)
	return r
}
func Test_split(t *testing.T) {
	ok := reflect.DeepEqual(Parseparts("/id/:name"), []string{"id", ":name"})
	ok = ok && reflect.DeepEqual(Parseparts("/x/y/z/*"), []string{"x", "y", "z", "*"})
	ok = ok && reflect.DeepEqual(Parseparts("/*/x/y"), []string{"*"})
	ok = ok && reflect.DeepEqual(Parseparts("/*x/y"), []string{"*x"})
	if !ok {
		t.Fatalf("test fail")
	}
}

func TestGetRoute(t *testing.T) {
	r := newTestRouter()
	n, ps := r.GetRoute("GET", "/assets/geektutu")

	if n == nil {
		t.Fatal("nil shouldn't be returned")
	}

	if n.path != "/assets/*filepath" {
		t.Fatal("should match /assets/*filepath")
	}

	if ps["filepath"] != "geektutu" {
		t.Fatal("geektutu should be equal to 'geektutu'")
	}

	fmt.Printf("matched path: %s, params['name']: %s\n", n.path, ps["name"])

}
