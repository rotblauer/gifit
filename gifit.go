package main

import (
	"fmt"
	"net/http"
	"time"
)

import "encoding/json"

type Data struct {
	ID string
}

var myClient = &http.Client{Timeout: 10 * time.Second}

const GiphyPublicAPIKey = "dc6zaTOxFJmzC"

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func main() {
	// var d interface{}
	// var d interface{}
	var d []Data
	getJson("http://api.giphy.com/v1/gifs/search?q=funny+cat&api_key=dc6zaTOxFJmzC", &d)
	fmt.Println(d[0])
	// m := d.(map[string]interface{})
	// for k, v := range m {
	// 	switch vv := v.(type) {
	// 	case string:
	// 		fmt.Println(k, "is string", vv)
	// 	case int:
	// 		fmt.Println(k, "is int", vv)
	// 	case []interface{}:
	// 		fmt.Println(k, "is an array:")
	// 		for i, u := range vv {
	// 			fmt.Println(i, u)
	// 		}
	// 	default:
	// 		fmt.Println(k, "is of a type I don't know how to handle")
	// 	}
	// }

	// fmt.Println(m.data[0].id)

}
