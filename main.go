package main

import "fmt"
import "net/http"
import "encoding/json"

type request struct {
	Name string
}

type response struct {
	Message string
}

type greeter struct {
	Default string
}

func NewGreeter(greeting string) *greeter {
	g := new(greeter)
	g.Default = greeting
	return g
}

func (this *greeter) ServeHTTP(respWriter http.ResponseWriter, point *http.Request) {
	req := new(request)
	decoder := json.NewDecoder(point.Body)

	if err := decoder.Decode(req); err != nil {
		panic(err)
	}

	resp := new(response)
	resp.Message = this.Default

	res2B, _ := json.Marshal(resp)

	fmt.Fprintf(respWriter, string(res2B))
}

func main() {
	g := NewGreeter("bye")
	http.Handle("/event1", g)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}

}
