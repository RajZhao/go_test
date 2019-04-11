package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		hasHttp := strings.HasPrefix(url,"http://")
		fmt.Println("has Http", hasHttp)
		if !hasHttp{
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprint(os.Stderr, "fetch:%v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
