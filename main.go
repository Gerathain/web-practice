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
	"strconv"
)

const MAX_CT int = 454
const MIN_CT int = 153

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

func putRequest(keyURL string, data io.Reader, roomNumber int, putLocation string) {
	client := &http.Client{}

	URL := keyURL + "/groups/" + strconv.Itoa(roomNumber) + "/" + putLocation

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

func getState(url string, roomNumber int) Group {
	group := Group{}

	fullURL := url + "/groups/" + strconv.Itoa(roomNumber)
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

/* TODO need to change the function to get the URL of the bridge properly.
This should not be too hard as it is documented by Philips */
func getBridgeUrl() string {
	return "http://192.168.0.17/api/"
}

func getApiKey( fileName string ) string {
	k, _ := ioutil.ReadFile( "./" + fileName )
	apiKey := strings.TrimSuffix(string(k), "\n")

	return apiKey
}

func main() {
	/* This will be useful if I want the script to return something to the
	   user via HTTP
	
	g := NewGreeter("bye")
	http.Handle("/event1", g)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
	*/

	keyURL := getBridgeUrl() + getApiKey( "apiKey" )

	var roomNumber int = 1 /* TODO get the room number dynamically */

	group := getState(keyURL, roomNumber) 
	actualCt := group.Action.Ct

	if actualCt < MAX_CT {
		/* for the minute, increasing ct by 10 per cycle seems good. and means the script doesnt have to run so frequently,
		might change to 5 or so */
		group.Action.Ct += 10
	}

	res, _ := json.Marshal(group.Action)

	if group.GroupState.All_on {
		/* PUT the group with the new ct
		need a to change func (or create new one) as it currently is hard coded to post to the wrong level
		To get around this, I am creating a second JSON object at a different level, is quite wasteful*/
		putRequest(keyURL, bytes.NewReader(res), roomNumber, "action")
	} 
	/*else lights are off so do nothing*/
}
