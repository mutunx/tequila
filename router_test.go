package tequila

import (
	"fmt"
	"reflect"
	"testing"
)

func newTestRouter() *router {
	r := newRouter()

	r.addRoute("GET", "/hello", nil)
	r.addRoute("GET", "/hello/:lang/say", nil)
	r.addRoute("GET", "/hello/file/*fileType", nil)
	return r
}

func TestParsePath(t *testing.T) {
	ok := reflect.DeepEqual(parsePath("hello/world"), []string{"hello", "world"})
	ok = ok && reflect.DeepEqual(parsePath("hello/:lang/say"), []string{"hello", ":lang", "say"})
	ok = ok && reflect.DeepEqual(parsePath("hello/*filePath/say"), []string{"hello", "*filePath"})
	ok = ok && reflect.DeepEqual(parsePath("hello/*"), []string{"hello", "*"})
	if !ok {
		t.Fatal("test parsePath fail")
	}
}

func TestGetRoute(t *testing.T) {
	r := newTestRouter()
	n, ps := r.getRoute("GET", "/hello/go/say")

	if n == nil {
		t.Fatal("nil shouldn't be returned")
	}

	if n.Path != "/hello/:lang/say" {
		t.Fatal("should match /hello/:lang/say")
	}

	if ps["lang"] != "go" {
		t.Fatal("name should be equal to 'go'")
	}

	fmt.Printf("matched path: %s, params['name']: %s\n", n.Path, ps["name"])

}
