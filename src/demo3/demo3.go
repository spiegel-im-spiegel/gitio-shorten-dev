package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/spiegel-im-spiegel/gitioapi"
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
	shortUrl, err := gitioapi.Encode(&gitioapi.Param{Url: urlStr})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Print(shortUrl)
}
