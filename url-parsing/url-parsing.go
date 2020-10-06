package main

import (
	"fmt"
	"net/url"
)

func main() {

	host := `https://www.learnku.com/docs/gobyexample/2020/url-analysis/6306?user=1&age=2`

	urls , e := url.Parse(host)
	if e != nil {
		panic(e)
	}

	fmt.Println(urls.Scheme)
	fmt.Println(urls.Host)
	fmt.Println(urls.Path)
	fmt.Println(urls.RawQuery)
	fmt.Println(url.ParseQuery(urls.RawQuery))

}
