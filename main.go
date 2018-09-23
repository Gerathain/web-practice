package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type State struct {
	On bool `json:"on,omitempty"`
	Ct int  `json:"ct,omitempty"`
}

type GroupState struct {
	All_on bool `json:"all_on"`
	Any_on bool `json:"any_on"`
}

type Action struct {
	On        bool   `json:"on"`
	Bri       int    `json:"bri"`
	Ct        int    `json:"ct"`
	Alert     string `json:"alert"`
	Colormode string `json:"colormode"`
}

type Group struct {
	Name       string     `json:"name"`
	Lights     []string   `json:"lights"`
	GroupState GroupState `json:"state"`
	Recycle    bool       `json:"recycle"`
	Class      string     `json:"class"`
	Action     Action     `json:"action"`
}

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

/*
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
*/

func putRequest(keyURL string, data io.Reader) {
	client := &http.Client{}

	URL := keyURL + "/groups/1/action"

	req, err := http.NewRequest(http.MethodPut, URL, data)
	if err != nil {
		log.Fatal(err)
	}

	var resp *http.Response
	resp, err = client.Do(req)

	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func getState(url string) Group {
	group := Group{}

	fullURL := url + "/groups/1/"
	response, err := http.Get(fullURL)
	if err != nil {
		fmt.Println(err)
	} else {
		//var dat map[string]interface{}

		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
		} else {
			if err := json.Unmarshal(contents, &group); err != nil {
				fmt.Println(err)
			} else {
				return group
			}
		}
	}

	return group
}

func main() {
	/*g := NewGreeter("bye")
	http.Handle("/event1", g)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
	*/

	k, _ := ioutil.ReadFile("./apiKey")
	apiKey := strings.TrimSuffix(string(k), "\n")
	keyURL := "http://192.168.1.4/api/" + string(apiKey)
	fmt.Println(keyURL)

	s := new(State)
	s.On = true
	res, _ := json.Marshal(s)

	s2 := new(State)
	s2.Ct = 200
	res2, _ := json.Marshal(s2)

	group := getState(keyURL)
	on := group.GroupState.All_on
	ct := group.Action.Ct
	if ct < MAX_CT { /*TODO define max ct*/
		ct++ /*TODO decide on rate at which lights get warmer */
	}
	res3, _ := json.Marshall(group)
	fmt.Printf("On: %t\n", on)
	if on {
		/* put the group with the new ct
		need a to change func (or create new one) as it currently is hard coded to post to the wrong level*/
		putRequest(keyURL, bytes.NewReader(res2))
	} else {
		putRequest(keyURL, bytes.NewReader(res))
	}
}
