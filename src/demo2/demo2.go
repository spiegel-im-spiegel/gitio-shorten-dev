package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func main() {
	//arguments
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, os.ErrInvalid)
		return
	}
	urlStr := os.Args[1]

	//shortening url
	resp, err := http.PostForm("http://git.io", url.Values{"url": {urlStr}})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Fprintln(os.Stderr, resp.Header.Get("Status"))
	if string(body) != urlStr {
		fmt.Fprintln(os.Stderr, string(body))
	}
	fmt.Fprint(os.Stdout, resp.Header.Get("Location"))
}
