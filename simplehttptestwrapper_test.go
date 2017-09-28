package simplehttptestwrapper

import (
	"fmt"
	"net/http"
	"testing"
)

func Test1(t *testing.T) {
	a := func() http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Hello, you hit foo!")
		}
	}

	b, err := HTTPTestHandler(a(), "GET", "/foo", HTTPTestRequestConfig{})
	fmt.Println("hi", b, err)

}
