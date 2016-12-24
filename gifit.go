package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	url "net/url"
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
	ID       string `json: "id"`
	Slug     string `json: "slug"`
	EmbedURL string `json: "embed_url"`
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

const GiphyPublicAPIKey = "dc6zaTOxFJmzC"

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

	// var useMarkdown bool
	// var useEmbeddable bool
	// var copyToClipboard bool
	var useStillImage bool

	var out string

	flag.BoolVar(&useStillImage, "i", false, "use still image instead of gif")
	// flag.BoolVar(&useMarkdown, "m", true, "format as markdown | DEFAULT")
	// flag.BoolVar(&useEmbeddable, "e", false, "use embeddable | incompatible with markdown")
	// flag.BoolVar(&copyToClipboard, "p", false, "copy to clipboard")

	flag.Parse()

	// if useEmbeddable {
	// 	useMarkdown = false
	// 	fmt.Println("Using embeddable.")
	// }
	// if useMarkdown {
	// 	fmt.Println("Using markdowny.")
	// }

	// $@
	nonflagArgs := flag.Args()
	if len(nonflagArgs) == 0 {
		fmt.Println("Useage: $ gifit hello kitty")
		fmt.Println("-i : use a still image")
		// fmt.Println("-m : format as markdown  | cannot use with -e | DEFAULT")
		// fmt.Println("-e : format as embed url | cannot use with -m")
		// fmt.Println("-p : copy to clipboard (pbcopy)")
		return
	}

	queryArgsString := strings.Join(nonflagArgs, " ")
	// fmt.Printf("qargs = %s\n", queryArgsString)
	encodedQuery := url.QueryEscape(queryArgsString)

	res := GiphyQueryResponse{}
	// http://api.giphy.com/v1/gifs/search?q=hello+kitty&limit=1&api_key=dc6zaTOxFJmzC
	getJson("http://api.giphy.com/v1/gifs/search?q="+encodedQuery+"&limit=1&api_key="+GiphyPublicAPIKey, &res)

	if len(res.Data) == 0 {
		fmt.Println("Got 0 results for the GIF search. Shiiii. Try again?")
		return
	}

	gifSource := res.Data[0].Images.Downsized.Url
	stillImageSource := res.Data[0].Images.Downsized_Still.Url
	// embeddableURL := res.Data[0].EmbedURL

	if !useStillImage {
		out = formatMarkdownImageMarkup(encodedQuery, gifSource)
	} else {
		out = formatMarkdownImageMarkup(encodedQuery, stillImageSource)
		// fmt.Println(out)
	}

	// if copyToClipboard {
	// 	c := exec.Command("pbcopy", out)
	// 	err := c.Run()
	// 	if err != nil {
	// 		fmt.Print(err)
	// 		return
	// 	}
	// 	fmt.Printf("(Copied to clipboard.) %s", out)
	// } else {
	fmt.Printf("%s", out)
	// }

}
