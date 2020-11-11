package main

import (
	"fmt"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
	"golang.org/x/net/html"
	"io/ioutil"
	"log"
	"strings"
)

func markdownFile2html(fileName string) string{
	content, fileErr := ioutil.ReadFile(fileName)
	if fileErr != nil {
		panic("ReadConfigFileError")
	}
	unsafe := blackfriday.Run([]byte(content))
	htmlContent := string(bluemonday.UGCPolicy().SanitizeBytes(unsafe))

	return string(htmlContent)
}

func parseHtml(n *html.Node) map[string]string{
	set := map[string]string{}
	//switch n.Type {
	//case html.ErrorNode:
	//
	//	log.Printf(n.Data)
	//case html.TextNode:
	//	log.Printf(n.Data)
	//case html.DocumentNode:
	//	log.Printf(n.Data)
	//case html.ElementNode:
	//	log.Printf(n.Data)
	//case html.CommentNode:
	//	log.Printf(n.Data)
	//case html.DoctypeNode:
	//	log.Printf(n.Data)
	//}
	if n.Type == html.ElementNode{
		var tt = n.Data
		if strings.HasSuffix(tt,">") && strings.HasPrefix(tt,"<"){
			set[n.Data] = n.Data
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				set1 := parseHtml(c)
				for k ,v := range set1 { set[k] = v }
			}
		}
	}
	return set

}

func getHtmlTag(htmlContent string) map[string]string{
	doc, err := html.Parse(strings.NewReader(htmlContent))
	if err != nil {
		log.Panic(err)
	}
	set := parseHtml(doc)
	return set
}



func main() {
	markdownHtml  :=  markdownFile2html("C:\\workspace\\useful-documents\\db\\dba-command.md")
	fmt.Println(markdownHtml)
	set := getHtmlTag(markdownHtml)
	for key := range  set{
		fmt.Println(key)
	}
}
