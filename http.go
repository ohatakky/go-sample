package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	// URLの生成
	base, _ := url.Parse("http://example.com")
	reference, _ := url.Parse("/test?a=1")
	endpoint := base.ResolveReference(reference).String()
	fmt.Println(endpoint)

	req, _ := http.NewRequest("Get", endpoint, nil)
	// req.Header.Add()
	q := req.URL.Query()
	fmt.Println(q)

	var client *http.Client = &http.Client{}
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
