package main

import (
	"bytes"
	"errors"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"log"
	"strings"
)

func main() {
	htmlStr := `
<h2>7z
</h2>

<p>A file archiver with highest compression ratio

Args:
a       add
d       delete
e       extract
l       list
t       test
u       update
x       extract with full paths

</p>

<h3>Example:
</h3>

<p>7z a -t7z -m0=lzma -mx=9 -mfb=64 -md=32m -ms=on archive.7z dir1&lt;/&gt;

-t7z        7z archive
-m0=lzma    lzma method
-mx=9       level of compression = 9 (ultra)
-mfb=64     number of fast bytes for lzma = 64
-md=32m     dictionary size = 32 Mb
-ms=on      solid archive = on

7z exit codes:
0       normal (no errors or warnings)
1       warning (non-fatal errors)
2       fatal error
7       bad cli arguments
8       not enough memory for operation
255     process was interrupted
</p>

`

	//doc, err := html.Parse(strings.NewReader(htmlStr))
	//if err != nil {
	//	log.Panic(err)
	//}
	//bn, err := Body(doc)
	//if err != nil {
	//	return
	//}
	doc, err := html.Parse(strings.NewReader(htmlStr))
	if err != nil {
		log.Panic(err)
	}
	bn, err := Body(doc)
	if err != nil {
		return
	}
	//body := renderNode(bn)
	//fmt.Println(body)
	for cc := bn.FirstChild; cc != nil; cc = cc.NextSibling {
		//node := bn.FirstChild
		text := &bytes.Buffer{}
		collectText(cc,text)
		fmt.Println(cc.Data,text)

	}


	//parseHtml(doc)
	//fmt.Println(doc.)

	//for c := doc.FirstChild; c != nil; c = c.NextSibling {
	//	node := c.FirstChild
	//	fmt.Println(c.Type, c.Data, node)
	//}

	//z := html.NewTokenizer(strings.NewReader(htmlStr))
	//depth := 0
	//for {
	//	tt := z.Next()
	//	switch tt {
	//	case html.ErrorToken:
	//		z.Err()
	//		break
	//	case html.TextToken:
	//		if depth > 0 {
	//			// emitBytes should copy the []byte it receives,
	//			// if it doesn't process it immediately.
	//			//emitBytes(z.Text())
	//			print(z.Text())
	//		}
	//	case html.StartTagToken, html.EndTagToken:
	//		tn, _ := z.TagName()
	//		if len(tn) == 1 && tn[0] == 'a' {
	//			if tt == html.StartTagToken {
	//				depth++
	//			} else {
	//				depth--
	//			}
	//		}
	//	}
	//}


}
func collectText(n *html.Node, buf *bytes.Buffer) {
	if n.Type == html.TextNode {
		buf.WriteString(n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		collectText(c, buf)
	}
}
func Body(doc *html.Node) (*html.Node, error) {
	var body *html.Node
	var crawler func(*html.Node)
	crawler = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "body" {
			body = node
			return
		}
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			crawler(child)
		}
	}
	crawler(doc)
	if body != nil {
		return body, nil
	}
	return nil, errors.New("Missing <body> in the node tree")
}

func renderNode(n *html.Node) string {
	var buf bytes.Buffer
	w := io.Writer(&buf)
	html.Render(w, n)
	return buf.String()
}

func parseHtml(n *html.Node) {
	log.Printf("node type:%d  ", n.Type)
	switch n.Type {
	case html.ErrorNode:
		log.Printf("ErrorNode(%p):%+v", n, n)
	case html.TextNode:
		log.Printf("TextNode(%p):%+v", n, n)
	case html.DocumentNode:
		log.Printf("DocumentNode(%p):%+v", n, n)
	case html.ElementNode:
		log.Printf("ElementNode(%p):%+v", n, n)
	case html.CommentNode:
		log.Printf("CommentNode(%p):%+v", n, n)
	case html.DoctypeNode:
		log.Printf("DoctypeNode(%p):%+v", n, n)
	}

	if n.Type == html.ElementNode && n.Data == "a" {
		if n.FirstChild != nil && n.FirstChild.Type == html.TextNode {
			log.Printf("href:%s, text:%s \n", n.Attr[0].Val, n.FirstChild.Data)
		} else {
			log.Printf("href:%s \n", n.Attr[0].Val)
		}

	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		parseHtml(c)
	}
}
