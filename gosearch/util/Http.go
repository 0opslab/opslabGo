package util

import (
	"net/http"
	"net/url"
	"errors"
	"io/ioutil"
	"opslabGo/gosearch/conf"
	"fmt"
)

/**
 * url转码
 */
func UrlEncode(uri string)(rsult string){
	u,_ := url.Parse(uri)
	q := u.Query()
	u.RawQuery = q.Encode()
	return u.String()
}

/**
 * 完成最底层的Get请求
 * GetWithClient returns the HTML returned by the url using a provided HTTP client
 */
func GetWithClient(url string, client *http.Client, Headers map[string]string, Cookies map[string]string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", errors.New("Request error " + url)
	}
	DEFUALT_HEADERS := map[string]string{
		"User-Agent": App.HTTPHEADER_USERAGENT,
		"Accept": App.HTTPHEADER_ACCEPT,
		"X-Requested-With":App.HTTPHEADER_REQUESTEDWITH,
		"Accept-Encoding":App.HTTPHEADER_ACCEPTENCODING,
		"Accept-Language":App.HTTPHEADER_ACCEPTLANGUAGE}
	for key, value := range DEFUALT_HEADERS {
		req.Header.Set(key, value)
	}
	for hName, hValue := range Headers {
		req.Header.Set(hName, hValue)
	}
	for cName, cValue := range Cookies {
		req.AddCookie(&http.Cookie{
			Name:  cName,
			Value: cValue,
		})
	}
	resp, err := client.Do(req)
	if err == nil  && resp.StatusCode == 200 {
		defer resp.Body.Close()
		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		if(App.IS_DEBUG){
			fmt.Println("request url :",url)
			fmt.Println(string(bytes))
		}
		return string(bytes), nil

	}else{
		return "", errors.New("GET request to " + url)
	}


}

/**
 * 获取指定url的html
 * 使用默认的http client获取html
 */
func GetHtml(url string) (string, error) {
	client := &http.Client{}
	return GetWithClient(url, client, nil, nil)
}

func GetHtmlWithHeaders(url string, Headers map[string]string) (string, error) {
	client := &http.Client{}

	return GetWithClient(url, client, Headers, nil)
}

func GetHtmlWithCookies(url string, cookies map[string]string) (string, error) {
	client := &http.Client{}
	return GetWithClient(url, client, nil, cookies)
}

func GetHtmlCustom(url string, headers map[string]string, cookies map[string]string) (string, error) {
	client := &http.Client{}
	return GetWithClient(url, client, headers, cookies)
}










