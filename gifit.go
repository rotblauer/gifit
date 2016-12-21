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
	ID          string `json: "id"`
	Slug        string `json: "slug"`
	BitlyGifUrl string `json: "bitly_gif_url"`
	// there's a lot more... TODO?
	Images ImagesObj `json: "images"`
}
type ImagesObj struct {
	Downsized ImagesTypeObj `json: "downsized"`
	// &c
}
type ImagesTypeObj struct {
	Url string `json: "url"`
	// &c
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

	useMarkdown := true
	useEmbeddable := false
	copyToClipboard := false

	var out string

	// flag.BoolVar(&useMarkdown, "m", false, "format as markdown")
	// flag.BoolVar(&useEmbeddable, "e", false, "use embeddable")
	// flag.BoolVar(&copyToClipboard, "p", false, "copy to clipboard")

	// flag.Parse()

	// if !useEmbeddable {
	// 	useMarkdown = true
	// 	fmt.Println("Using embed.")
	// } else {
	// 	fmt.Println("Using markdown.")
	// }

	// $@
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Useage: $ gifit hello kitty")
		// fmt.Println("-m : format as markdown  | cannot use with -e | DEFAULT")
		// fmt.Println("-e : format as embed url | cannot use with -m")
		// fmt.Println("-p : copy to clipboard (pbcopy)")
		return
	}
	// Remove all args that are flags.
	var queryArgs []string
	for _, arg := range args {
		if !strings.HasPrefix(arg, "-") {
			queryArgs = append(queryArgs, arg)
		}
	}
	queryArgsString := strings.Join(queryArgs, " ")
	encodedQuery := url.QueryEscape(queryArgsString)

	res := GiphyQueryResponse{}
	getJson("http://api.giphy.com/v1/gifs/search?q="+encodedQuery+"&limit=1&api_key="+GiphyPublicAPIKey, &res)

	if len(res.Data) == 0 {
		fmt.Println("Got 0 results for the GIF search. Shiiii. Try again?")
		return
	}

	markdownImageSource := res.Data[0].Images.Downsized.Url
	embeddableUrl := res.Data[0].BitlyGifUrl

	if useMarkdown {
		out = formatMarkdownImageMarkup(queryArgsString, markdownImageSource)
	}
	if useEmbeddable {
		out = embeddableUrl
	}
	if copyToClipboard {
		c := exec.Command("pbcopy", out)
		err := c.Run()
		if err != nil {
			fmt.Print(err)
			return
		}
		fmt.Printf("(Copied to clipboard.) %s", out)
	} else {
		fmt.Printf("%s", out)
	}

	// ga := exec.Command("git", "add", "-A")
	// ga_out, ga_err := ga.Output()
	// if ga_err != nil {
	// 	fmt.Println(ga_err)
	// 	return
	// }
	// fmt.Println(string(ga_out))

	// gc := exec.Command("git", "commit", "-m", "'"+formatMarkdownImageMarkup(commit_message, markdownImageSource)+"'")
	// gc_out, gc_err := gc.Output()
	// if gc_err != nil {
	// 	fmt.Println(gc_err)
	// 	return
	// }
	// fmt.Println(string(gc_out))

}
