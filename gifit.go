package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	url "net/url"
	"strconv"
	"strings"
	"time"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

type GiphyQueryResponse struct {
	Data       []Gif         `json: "data"`
	Meta       ResMeta       `json: "meta"`
	Pagination ResPagination `json: "pagination"`
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
type Gif struct {
	ID        string `json: "id"`
	Slug      string `json: "slug"`
	Embed_URL string `json: "embed_url"`
	// there's a lot more... TODO?
	Images ImagesObj `json: "images"`
}
type ImagesObj struct {
	Downsized       ImagesTypeObj `json: "downsized"`
	Downsized_Still ImagesTypeObj `json: "downsized_still"`
	// &c
}
type ImagesTypeObj struct {
	Url string `json: "url"`
	// &c
}

const GiphyPublicAPIKey string = "dc6zaTOxFJmzC"
const NumberQueryResults int = 20

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	j := json.NewDecoder(r.Body).Decode(target)

	return j
}

func formatMarkdownImageMarkup(altText, source string) string {
	return "![" + altText + "](" + source + ")"
}

func main() {

	var useEmbeddable bool
	var useStillImage bool

	var out string

	flag.BoolVar(&useStillImage, "s", false, "use still image instead of gif") // "s" for "still"
	flag.BoolVar(&useEmbeddable, "e", false, "use embeddable | incompatible with still image")

	flag.Parse()

	// embed haz priority
	if useEmbeddable {
		useStillImage = false
	}
	if useStillImage {
		useEmbeddable = false
	}

	// $@
	nonflagArgs := flag.Args()
	if len(nonflagArgs) == 0 {
		fmt.Println("Useage: $ gifit hello kitty")
		fmt.Println("   -i : use a still image")
		fmt.Println("   -e : format as embed url; CANNOT use with -i")
		return
	}

	queryArgsString := strings.Join(nonflagArgs, " ")
	encodedQuery := url.QueryEscape(queryArgsString)

	res := GiphyQueryResponse{}
	s := strconv.Itoa(NumberQueryResults)
	getJson("http://api.giphy.com/v1/gifs/search?q="+encodedQuery+"&limit="+s+"&api_key="+GiphyPublicAPIKey, &res)
	// http://api.giphy.com/v1/gifs/search?q=hello+kitty&limit=1&api_key=dc6zaTOxFJmzC

	if len(res.Data) == 0 {
		fmt.Printf("Got %d results for the GIF search. Shiiii. Try again?", len(res.Data))
		return
	}

	unixxx := time.Now().Unix()
	rand.Seed(unixxx)
	r := rand.Intn(len(res.Data) - 1) // [0, n]

	gifSource := res.Data[r].Images.Downsized.Url
	stillImageSource := res.Data[r].Images.Downsized_Still.Url
	embeddableURL := res.Data[r].Embed_URL

	out = formatMarkdownImageMarkup(encodedQuery, gifSource) // default with overrides below in case of option flags
	if useStillImage {
		out = formatMarkdownImageMarkup(encodedQuery, stillImageSource)
	}
	if useEmbeddable {
		out = embeddableURL
	}

	fmt.Printf("%s", out)

}
