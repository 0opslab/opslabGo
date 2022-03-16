package main

import (
	"encoding/json"
	"fmt"
)

type HelperConfig struct {
	//监听地址和端口
	ADDR string `json:"ADDR"`
	//文件写入路径
	PATH string `json:"PATH"`
	//同步的地址
	STYLE StyleConfig `json:"STYLE"`
}

type StyleConfig struct {
	H1_FG string `json:"h1_fg"`
	H1_BG string `json:"h1_bg"`
	H2_FG string `json:"h2_fg"`
	H2_BG string `json:"h2_bg"`
	H3_FG string `json:"h3_fg"`
	H3_BG string `json:"h3_bg"`
	H4_FG string `json:"h4_fg"`
	H4_BG string `json:"h4_bg"`
	H5_FG string `json:"h5_fg"`
	H5_BG string `json:"h5_bg"`
	H6_FG string `json:"h6_fg"`
	H6_BG string `json:"h6_bg"`
}

var App = HelperConfig{}

func main() {
	content := `{
		"addr":"0.0.0.0:9090",
	"path":"c:/var/upload/wwww/",
	"fileNameLength":"11",
	"STYLE":{
		"h1_fg":"FFFFFF",
		"h1_bg":"FF0000",
		"h2_fg":"FFFFFF",
		"h2_bg":"FF0000",
		"h3_fg":"FFFFFF",
		"h3_bg":"FF0000",
		"h4_fg":"FFFFFF",
		"h4_bg":"FF0000",
		"h5_fg":"FFFFFF",
		"h5_bg":"FF0000",
		"h6_fg":"FFFFFF",
		"h6_bg":"FF0000"
	}
	}`

	err := json.Unmarshal([]byte(content), &App)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(App.ADDR)
	fmt.Println(App.PATH)
}
