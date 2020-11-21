package main

import (
	"bytes"
	"errors"
	"github.com/gookit/color"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
	"golang.org/x/net/html"
	"io"
	"io/ioutil"
	"log"
	"opslabGo/monsoon"
	"strings"
)

func color_markdown_print(content string) {
	// 利用 bluemonday 解析
	unsafe := blackfriday.Run([]byte(content))
	htmlContent := string(bluemonday.UGCPolicy().SanitizeBytes(unsafe))
	//fmt.Println(htmlContent)
	doc, err := html.Parse(strings.NewReader(htmlContent))
	if err != nil {
		log.Panic(err)
	}
	bn, err := Body(doc)
	if err != nil {
		return
	}

	h1 := color.HEXStyle("FFFFFF", "FF0000")
	h2 := color.HEXStyle("FFFFFF", "FF0000")
	h3 := color.HEXStyle("FFFFFF", "FF0000")
	h4 := color.HEXStyle("FFFFFF", "FF0000")
	h5 := color.HEXStyle("FFFFFF", "FF0000")
	h6 := color.HEXStyle("FFFFFF", "FF0000")
	Code := color.Success
	Comment := color.Gray

	for cc := bn.FirstChild; cc != nil; cc = cc.NextSibling {
		text := &bytes.Buffer{}
		collectText(cc, text)
		textContent := text.String()
		switch cc.Data {
		case "h1":
			h1.Println(textContent)
		case "h2":
			h2.Println(textContent)
		case "h3":
			h3.Println(textContent)
		case "h4":
			h4.Println(textContent)
		case "h5":
			h5.Println(textContent)
		case "h6":
			h6.Println(textContent)
		case "p":
			p := renderNode(cc)
			pLen := len(p)
			p = p[3 : pLen-4]
			if strings.HasPrefix(p, "\n<code>") {
				codeLine := strings.Split(p, "\n")
				codeLen := len(codeLine)
				for i := 1; i < codeLen-1; i++ {
					line := codeLine[i]
					if strings.HasPrefix(line, "#") {
						color.Gray.Println()
						//Comment.Println(line)
					} else {
						color.Green.Println(line)
					}
				}
			} else {
				Comment.Println(textContent)
			}
		default:
			Code.Println(textContent)
		}
	}

}

func renderNode(n *html.Node) string {
	var buf bytes.Buffer
	w := io.Writer(&buf)
	html.Render(w, n)
	return buf.String()
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

func main() {

	//confile := flag.String("conf", "", "the configuration file")
	//flag.Parse()
	//if *confile == "" {
	//	fmt.Println("Please specify the configuration file")
	//	return
	//}

	cmdPath := "/data/workspace/useful-command"
	files, _, _ := monsoon.WalkDirFiles(cmdPath, "md")
	for _, file := range files {

		fileContent, err := ioutil.ReadFile(file)
		if err != nil {
			color.Red.Println("ReadFileError")
			return
		}
		color_markdown_print(string(fileContent))
	}

}