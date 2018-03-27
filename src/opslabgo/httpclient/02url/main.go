package main

import (
	"net/url"
	"log"
	"fmt"
)

func main() {

	u, err := url.Parse("http://bing.com/search?q=go-url")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u)

	u.Scheme = "https"
	u.Host = "google.com"
	q := u.Query()
	q.Set("q", "golang")
	q.Set("s", "编程")
	u.RawQuery = q.Encode()
	// https://google.com/search?q=golang&s=%E7%BC%96%E7%A8%8B
	fmt.Println(u)

	fmt.Println(u.Path)
	fmt.Println(u.RawPath)
	fmt.Println(u.String())

	u1, err := url.Parse("../../..//search?q=dotnet")
	if err != nil {
		log.Fatal(err)
	}
	base, err := url.Parse("http://example.com/directory/")
	if err != nil {
		log.Fatal(err)
	}
	//http://example.com/search?q=dotnet
	fmt.Println(base.ResolveReference(u1))

	v := url.Values{}
	v.Set("name", "Ava")
	v.Add("friend", "Jess")
	v.Add("friend", "Sarah")
	v.Add("friend", "Zoe")
	// v.Encode() == "name=Ava&friend=Jess&friend=Sarah&friend=Zoe"

	//Ava
	fmt.Println(v.Get("name"))
	//Jess
	fmt.Println(v.Get("friend"))
	//[Jess Sarah Zoe]
	fmt.Println(v["friend"])
}
