package util



import (
	"time"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"fmt"
	"errors"
	"github.com/google/logger"
	"strings"
	"opslabGo/gosearch/conf"
)


/**
 * 获取GOqueryDOC
 */
func GoQueryDoc(url string,timeout time.Duration) (doc *goquery.Document){
	chan_doc := make(chan *goquery.Document, 1)

	go func(url string) {
		client := &http.Client{}
		html, err := GetWithClient(url, client, nil, nil)
		if err == nil{
			doc, _ := goquery.NewDocumentFromReader(strings.NewReader(html))
			chan_doc <- doc
		}else{
			logger.Error("request error",url)
		}

	}(url)

	select {
	case res := <-chan_doc:
		return res
	case <-time.After(timeout):

		doc, _ := goquery.NewDocumentFromReader(strings.NewReader(App.HTML_EMPATY))
		return doc
	}
}

func GoQueryDocWithHeaders(url string, Headers map[string]string,timeout time.Duration) (doc *goquery.Document) {
	chan_doc := make(chan *goquery.Document, 1)

	go func(url string, Headers map[string]string) {
		html, err := GetHtmlWithHeaders(url, Headers)
		if err == nil{
			doc, _ := goquery.NewDocumentFromReader(strings.NewReader(html))
			chan_doc <- doc
		}else{
			logger.Error("request error",url)
		}
	}(url,Headers)

	select {
	case res := <-chan_doc:
		return res
	case <-time.After(timeout):
		doc, _ := goquery.NewDocumentFromReader(strings.NewReader(App.HTML_EMPATY))
		return doc
	}
}

func GoQueryDocWithCookies(url string, cookies map[string]string,timeout time.Duration) (doc *goquery.Document) {
	chan_doc := make(chan *goquery.Document, 1)

	go func(url string,cookies map[string]string) {
		html, err := GetHtmlWithCookies(url, cookies)
		if err == nil{
			doc, _ := goquery.NewDocumentFromReader(strings.NewReader(html))
			chan_doc <- doc
		}else{
			logger.Error("request error",url)
		}
	}(url,cookies)

	select {
	case res := <-chan_doc:
		return res
	case <-time.After(timeout):
		doc, _ := goquery.NewDocumentFromReader(strings.NewReader(App.HTML_EMPATY))
		return doc
	}
}

func GoQueryDocCustom(url string, headers map[string]string, cookies map[string]string,timeout time.Duration) (doc *goquery.Document) {
	chan_doc := make(chan *goquery.Document, 1)

	go func(url string, headers map[string]string, cookies map[string]string) {
		html, err := GetHtmlCustom(url, headers, cookies)
		if err == nil{
			doc, _ := goquery.NewDocumentFromReader(strings.NewReader(html))
			chan_doc <- doc
		}else{
			logger.Error("request error",url)
		}
	}(url,headers,cookies)

	select {
	case res := <-chan_doc:
		return res
	case <-time.After(timeout):
		doc, _ := goquery.NewDocumentFromReader(strings.NewReader(App.HTML_EMPATY))
		return doc
	}
}


/**
 获取html
 使用原生的goquery方法请求并返回文档类型
 */
func GoQueryRequestDoc(url string, refer string, timeout time.Duration) (doc *goquery.Document, err error) {
	chan_doc := make(chan *goquery.Document, 1)

	go func(url string, refer string) {
		headers := &http.Header{}
		headers.Set("User-Agent", App.HTTPHEADER_USERAGENT)
		headers.Set("Accept", App.HTTPHEADER_ACCEPT)
		headers.Set("X-Requested-With", App.HTTPHEADER_REQUESTEDWITH)
		headers.Set("Accept-Encoding", App.HTTPHEADER_ACCEPTENCODING)
		headers.Set("Accept-Language", App.HTTPHEADER_ACCEPTLANGUAGE)
		headers.Set("Referer", refer)

		docs, err := goquery.NewDocument(url)
		if err != nil {
			logger.Warning("request Doc Error ", url, err)
		} else {
			chan_doc <- docs
		}

	}(url, refer)

	select {
	case res := <-chan_doc:
		return res, nil
	case <-time.After(time.Second * 1):
		fmt.Println("request doc timeout ", url)
		return nil, errors.New("request-doc-timeout")
	}
}


/**
 获取html
 使用原生的goquery方法请求并返回文档类型
 */
func GoQueryRequestDocNoError(url string, refer string, timeout time.Duration) (doc *goquery.Document) {
	chan_doc := make(chan *goquery.Document, 1)

	go func(url string, refer string) {
		headers := &http.Header{}
		headers.Set("User-Agent", App.HTTPHEADER_USERAGENT)
		headers.Set("Accept", App.HTTPHEADER_ACCEPT)
		headers.Set("X-Requested-With", App.HTTPHEADER_REQUESTEDWITH)
		headers.Set("Accept-Encoding", App.HTTPHEADER_ACCEPTENCODING)
		headers.Set("Accept-Language", App.HTTPHEADER_ACCEPTLANGUAGE)
		headers.Set("Referer", refer)

		docs, err := goquery.NewDocument(url)
		if err != nil {
			logger.Warning("request Doc Error ", url, err)
		} else {
			chan_doc <- docs
		}

	}(url, refer)


	select {
	case res := <-chan_doc:

		if(App.IS_DEBUG){
			html, _ := res.Html()
			fmt.Println("Request doc:", html)
		}
		return res
	case <-time.After(time.Second * 1):
		logger.Info("request doc timeout ", url)
		doc, _ := goquery.NewDocumentFromReader(strings.NewReader(App.HTML_EMPATY))
		return doc
	}
}