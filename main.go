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
	"reflect"
	"errors"
	"time"
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

type Config struct {
	ControllerOn	*bool `json:"controller_on"`	
	StartHour	*int  `json:"start_hour"`
	StartMinute	*int  `json:"start_minute"`
	WarmingRate	*int  `json:"warming_rate"`
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

/* Gets the state of the lights in a room and returns them in a Group struct */
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

/* Reads the specified config file and returns a Config variable with the 
controller's configuration
*/
func readConfig( configFileName string , config *Config) error {

	configFile, err := ioutil.ReadFile( "./" + configFileName )
	if err != nil {
		return err
	}
	if err = json.Unmarshal(configFile, config); err != nil {
		return err
	}
	
	//reflection to get members of struct Config to check that none are nil
	v := reflect.ValueOf(config)

	// v is a reflect.Value type, we know it is a pointer (because config is a pointer) so we dereference with Elem()
	for i := 0; i < v.Elem().NumField(); i++ {
		// Get struct with Elem. get field in it 
		field := v.Elem().Field(i)

		// field is of type reflect.Ptr, so we use the library IsNil() function to check that it is not nil	
		if field.IsNil() {
			return errors.New(  "error in config, likely due to incorrect JSON")
		}
	}

	return nil
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

	var config Config
	if err := readConfig( "config", &config) ; err!= nil {
		fmt.Println( err )
		return
	}
	if *(config.ControllerOn) == false {
		return
	}

	// need to check if we are after the time to start running
	t := time.Now()
	hour, minute, _ := t.Clock()
	// This current if means that the controller will always consider midnight as before the time to change the lights
	if !( ( hour > *(config.StartHour) ) || ( hour == *(config.StartHour) && minute >= *(config.StartMinute) ) ) {
		//Not time to change the lights
		fmt.Println("Not time to change the lights")
		return	
	}
	
	var roomNumber int = 1 /* TODO get the room number dynamically */

	group := getState(keyURL, roomNumber) 
	actualCt := group.Action.Ct

	if actualCt < MAX_CT {
		group.Action.Ct += *(config.WarmingRate)
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
