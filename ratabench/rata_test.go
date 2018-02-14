package ratabench

import (
	"net/http"
	"testing"

	"github.com/tedsuo/rata"
)

func BenchmarkMux(b *testing.B) {
	routes := rata.Routes{{
		Name:   "test",
		Path:   "/v1/v1",
		Method: rata.GET,
	}}
	handlers := rata.Handlers{
		"test": http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
	}
	router, err := rata.NewRouter(routes, handlers)
	if err != nil {
		b.Fatal(err)
	}

	request, _ := http.NewRequest("GET", "/v1/v1", nil)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		router.ServeHTTP(nil, request)
	}
}

func BenchmarkVar(b *testing.B) {
	routes := rata.Routes{{
		Name:   "test",
		Path:   "/v1/:v1",
		Method: rata.GET,
	}}
	handlers := rata.Handlers{
		"test": http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
	}
	router, err := rata.NewRouter(routes, handlers)
	if err != nil {
		b.Fatal(err)
	}

	request, _ := http.NewRequest("GET", "/v1/anything", nil)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		router.ServeHTTP(nil, request)
	}
}

func BenchmarkMultiVar(b *testing.B) {
	routes := rata.Routes{{
		Name:   "test",
		Path:   "/component/:v1/sub_component/:v2",
		Method: rata.GET,
	}}
	handlers := rata.Handlers{
		"test": http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
	}
	router, err := rata.NewRouter(routes, handlers)
	if err != nil {
		b.Fatal(err)
	}

	request, _ := http.NewRequest("GET", "/component/something/sub_component/anything", nil)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		router.ServeHTTP(nil, request)
	}
}
