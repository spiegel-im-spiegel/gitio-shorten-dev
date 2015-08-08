package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {
	resp, err := http.PostForm("http://git.io", url.Values{"url": {"https://github.com/spiegel-im-spiegel"}})
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("  Status: ", resp.Header.Get("Status"))
	log.Println("Location: ", resp.Header.Get("Location"))
	log.Println("    Body: ", string(body))
}
