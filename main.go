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

func main() {
	http.HandleFunc("/event1", func(respWriter http.ResponseWriter, point *http.Request) {
		req := new(request)
		decoder := json.NewDecoder(point.Body)
		err := decoder.Decode(req)
		if err != nil {
			panic(err)
		}

		resp := new(response)
		resp.Message = "Bye"
		res2B, _ := json.Marshal(resp)

		fmt.Fprintf(respWriter, string(res2B))
	})

	http.ListenAndServe(":8080", nil)
}
