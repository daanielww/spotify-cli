package main

import (
	"encoding/json"
	"fmt"
	"github.com/TylerBrock/colorjson"
	"github.com/daanielww/spotify-cli/client/config"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	c, err := config.GetConfiguration()
	if err != nil {
		log.Fatalf("couldn't get config: %v", err)
	}

	fmt.Println("making request to: ", c.Url)
	resp, err := http.Get(c.Url)
	if err != nil {
		log.Fatalf("error performing request: %v", err)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("error performing ioutil.ReadAll: %v", err)
	}

	var obj map[string]interface{}
	json.Unmarshal(responseBody, &obj)
	fmt.Println(obj)

	f := colorjson.NewFormatter()
	f.Indent = 4

	// Marshall the Colorized JSON
	s, _ := f.Marshal(obj)
	fmt.Println(string(s))
}