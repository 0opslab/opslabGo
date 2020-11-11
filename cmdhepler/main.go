package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gookit/color"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
	"golang.org/x/net/html"
	"io"
	"io/ioutil"
	"log"
	"strings"
)

type HelperConfig struct {
	//监听地址和端口
	ADDR string `json:'ADDR'`
	//文件写入路径
	PATH string `json:'PATH'`
	//同步的地址
	STYLE struct {
		H1_FG string `json STYLE.h1_fg`
		H1_BG string `json STYLE.h1_bg`
		H2_FG string `json STYLE.h2_fg`
		H2_BG string `json STYLE.h2_bg`
		H3_FG string `json STYLE.h3_fg`
		H3_BG string `json STYLE.h3_bg`
		H4_FG string `json STYLE.h4_fg`
		H4_BG string `json STYLE.h4_bg`
		H5_FG string `json STYLE.h5_fg`
		H5_BG string `json STYLE.h5_bg`
		H6_FG string `json STYLE.h6_fg`
		H6_BG string `json STYLE.h6_bg`
	} `json:'STYLE'`
}

var app = HelperConfig{}

func main() {

	//confile := flag.String("conf", "", "the configuration file")
	//flag.Parse()
	//if *confile == "" {
	//	fmt.Println("Please specify the configuration file")
	//	return
	//}
	jsonConfigFile := "C:\\workspace\\opslabGo\\cmdhepler\\conf\\server1.conf"
	content, fileErr := ioutil.ReadFile(jsonConfigFile)
	if fileErr != nil {
		color.Red.Println("ReadConfigFileError")
		return
	}
	err := json.Unmarshal(content, &app)
	if err != nil {
		color.Red.Println("ParseConfigFileError")
		return
	}

	fileName := "C:\\workspace\\useful-documents\\db\\dba-command.md"
	fileContent := loadFileContent(fileName)
	color_print(fileContent)
}
func loadFileContent(file string) string {
	content, fileErr := ioutil.ReadFile(file)
	if fileErr != nil {
		panic("Read Config Error")
	}
	return string(content)
}

func color_print(content string) {
	// 利用 bluemonday 解析
	unsafe := blackfriday.Run([]byte(content))
	htmlContent := string(bluemonday.UGCPolicy().SanitizeBytes(unsafe))
	fmt.Println(htmlContent)
	doc, err := html.Parse(strings.NewReader(htmlContent))
	if err != nil {
		log.Panic(err)
	}
	bn, err := Body(doc)
	if err != nil {
		return
	}
	//fmt.Println("@@@@",bn)
	//fmt.Println("@@@@")
	//body := renderNode(bn)
	//fmt.Println(body)
	h1 := color.HEXStyle(app.STYLE.H1_FG, app.STYLE.H1_BG)
	h2 := color.HEXStyle(app.STYLE.H2_FG, app.STYLE.H2_BG)
	h3 := color.HEXStyle(app.STYLE.H3_FG, app.STYLE.H3_BG)
	h4 := color.HEXStyle(app.STYLE.H4_FG, app.STYLE.H4_BG)
	h5 := color.HEXStyle(app.STYLE.H5_FG, app.STYLE.H5_BG)
	h6 := color.HEXStyle(app.STYLE.H6_FG, app.STYLE.H6_BG)

	for cc := bn.FirstChild; cc != nil; cc = cc.NextSibling {
		text := &bytes.Buffer{}
		collectText(cc, text)

		switch cc.Data {
		case "h1":
			h1.Println(text)
		case "h2":
			h2.Println(text)
		case "h3":
			h3.Println(text)
		case "h4":
			h4.Println(text)
		case "h5":
			h5.Println(text)
		case "h6":
			h6.Println(text)
		case "p":
			p := renderNode(cc)
			pLen := len(p)
			p = p[3:pLen-4]
			if strings.HasPrefix(p, "\n<code>") {
				codeLine := strings.Split(p, "\n")
				codeLen := len(codeLine)
				for i := 1; i < codeLen -1; i++ {
					line := codeLine[i]
					if strings.HasPrefix(line, "#") {
						color.Gray.Println(line)
					} else {
						color.Green.Println(line)
					}
				}
			} else {
				color.Gray.Println(p)
			}

			//text.WriteString(cc.FirstChild.Data)
			//fmt.Println(text)
			//case "code":
			//	fmt.Println("@@@@")
			//	color.Green.Println(text)
		}
	}

	//lines := strings.Split(content, "\n")
	//for i := 0; i < len(lines); i++ {
	//	line := lines[i]
	//	i := strings.HasPrefix(line, "#")
	//	if (i) {
	//		re3, _ := regexp.Compile("^#{1,}\\s+");
	//		line := re3.ReplaceAllString(line, "");
	//		h1.Println(line)
	//	} else {
	//		color.Gray.Println(line)
	//	}
	//
	//}

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
