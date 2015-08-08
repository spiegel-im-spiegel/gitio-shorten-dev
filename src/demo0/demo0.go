package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
)

func main() {
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)
	writer.WriteField("url", "https://github.com/spiegel-im-spiegel")
	writer.Close()

	resp, err := http.Post("http://git.io", "multipart/form-data; boundary="+writer.Boundary(), &buffer)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(string(body))
}

