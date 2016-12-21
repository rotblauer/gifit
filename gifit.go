package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	url "net/url"
	"os"
	"os/exec"
	"strings"
	"time"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

type GiphyQueryResponse struct {
	Data       []Gif         `json: "data"`
	Meta       ResMeta       `json: "meta"`
	Pagination ResPagination `json: "pagination"`
}
type Gif struct {
	ID   string `json: "id"`
	Slug string `json: "slug"`
	// there's a lot more... TODO?
}
type ResPagination struct {
	Total_count int `json: "total_count"`
	Count       int `json: "count"`
	Offset      int `json: "offset"`
}
type ResMeta struct {
	Status      int    `json: "status"`
	Msg         string `json: "msg"`
	Response_id string `json: "response_id"`
}

const GiphyPublicAPIKey = "dc6zaTOxFJmzC"

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func formatMarkdownImageMarkup(altText, source string) string {
	return "![" + altText + "](" + source + ")"
}

func main() {
	// $@
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Useage: $ gifit hello kitty")
		return
	}
	commit_message := strings.Join(args, " ")
	encoded_query := url.QueryEscape(commit_message)

	res := GiphyQueryResponse{}
	getJson("http://api.giphy.com/v1/gifs/search?q="+encoded_query+"&limit=1&api_key="+GiphyPublicAPIKey, &res)

	if len(res.Data) == 0 {
		fmt.Println("Got 0 results for the GIF search. Shiiii. Try again?")
		return
	}

	gif_id := res.Data[0].ID

	// TODO: handle different kinds of file url provided in Gif json object (?) so we don't have to hardcode url?
	markdownImageSource := "http://i.giphy.com/" + gif_id + ".gif"

	ga := exec.Command("git", "add", "-A")
	ga_out, ga_err := ga.Output()
	if ga_err != nil {
		fmt.Println(ga_err)
		return
	}
	fmt.Println(ga_out)

	gc := exec.Command("git", "commit", "-m", "'"+formatMarkdownImageMarkup(commit_message, markdownImageSource)+"'")
	gc_out, gc_err := gc.Output()
	if gc_err != nil {
		fmt.Println(gc_err)
		return
	}
	fmt.Println(gc_out)

}
