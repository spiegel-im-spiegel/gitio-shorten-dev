package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func main() {
	//arguments
	flag.Parse()
	argsStr := flag.Args()
	if len(argsStr) < 1 {
		fmt.Fprintln(os.Stderr, os.ErrInvalid)
		return
	}
	urlStr := argsStr[0]

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
