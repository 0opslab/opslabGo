package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gookit/color"
	"io/ioutil"
	"opslabGo/monsoon"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"runtime"
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

var App = HelperConfig{}

//创建并加载配置文件
func createLoadConfFile(v interface{}) error{
	sysType := runtime.GOOS
	jsonConfigFile := ""
	if sysType == "linux" {
		jsonConfigFile = "/etc/cmdhelp.conf"
	}
	if sysType == "windows" {
		jsonConfigFile = "C:/windows/cmdhelp.conf"
	}
	if !monsoon.FileIsExist(jsonConfigFile) {
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
		monsoon.WriteString(jsonConfigFile, content)
		err := json.Unmarshal([]byte("content"), &v)
		return err
	} else {
		content, err := ioutil.ReadFile(jsonConfigFile)
		if err == nil {
			err = json.Unmarshal(content, &v)
		}
		return err
	}

}

func main() {
	err := createLoadConfFile(&App)
	if err != nil{
		color.Red.Println("ReadConfigFileError")
		return
	}

	fmt.Println(App)

	//fileName := "C:\\workspace\\useful-command\\doc\\git.md"
	//fileContent, err := ioutil.ReadFile(fileName)
	//if err != nil {
	//	color.Red.Println("ReadFileError")
	//	return
	//}
	//color_print(string(fileContent))


	cmdPath := "/data/workspace/useful-command"
	for i, v := range os.Args {
		if i > 0 {
			files, _, _ := monsoon.WalkDirFiles(cmdPath, "md")
			for _, file := range files {
				if strings.ToLower(TrimFile(file)) == strings.ToLower(v) {
					fileContent, err := ioutil.ReadFile(file)
					if err != nil {
						color.Red.Println("ReadFileError")
						return
					}
					color_print(string(fileContent))
				}
			}
		}
	}

}

/**
 * @func 命令行彩色打印
 */
func color_print(content string) {
	h1 := color.HEXStyle(App.STYLE.H1_FG, App.STYLE.H1_BG)
	h2 := color.HEXStyle(App.STYLE.H2_FG, App.STYLE.H2_BG)
	h3 := color.HEXStyle(App.STYLE.H3_FG, App.STYLE.H3_BG)
	h4 := color.HEXStyle(App.STYLE.H4_FG, App.STYLE.H4_BG)
	h5 := color.HEXStyle(App.STYLE.H5_FG, App.STYLE.H5_BG)
	h6 := color.HEXStyle(App.STYLE.H6_FG, App.STYLE.H6_BG)
	Code := color.Success
	Comment := color.Gray
	content = strings.ReplaceAll(content, "\r\n", "\n")
	content = strings.ReplaceAll(content, "\r", "\n")
	lines := strings.Split(content, "\n")
	var buffer bytes.Buffer
	adds := false

	var s5 []string
	for _, line := range lines {
		if strings.HasPrefix(line, "```") && !adds {
			adds = true
		}

		if adds {
			buffer.WriteString(line + "\n")
			if line == "```" {
				s5 = append(s5, buffer.String())
				buffer.Reset()
				adds = false
			}
		} else {
			s5 = append(s5, line)
		}

	}

	for _, value := range s5 {
		if strings.HasPrefix(value, "######") {
			value = value[6:]
			h6.Printf("%-120s", trimStr(value))
			Comment.Println("")
		} else if strings.HasPrefix(value, "#####") {
			value = value[5:]
			h5.Printf("%-120s", trimStr(value))
			Comment.Println("")
		} else if strings.HasPrefix(value, "####") {
			value = value[4:]
			h4.Printf("%-120s", trimStr(value))
			Comment.Println("")
		} else if strings.HasPrefix(value, "###") {
			value = value[3:]
			h3.Printf("%-120s", trimStr(value))
			Comment.Println("")
		} else if strings.HasPrefix(value, "##") {
			value = value[2:]
			h2.Printf("%-120s", trimStr(value))
			Comment.Println("")
		} else if strings.HasPrefix(value, "#") {
			value = value[1:]
			h1.Printf("%-120s", trimStr(value))
			Comment.Println("")
		} else if strings.HasPrefix(value, "```") {
			ll := strings.Split(value, "\n")
			for _, l := range ll {
				if strings.HasPrefix(l, "```") {
					continue
				}
				if strings.HasPrefix(l, "#") || strings.HasPrefix(l, "/") {
					Comment.Println(delComment(trimStr(l)))
				} else if strings.HasPrefix(l, "^") {
					Code.Println(trimStr(l[1:]))
				} else {
					Code.Println(trimStr(l))
				}
			}
		} else {
			Comment.Println(delComment(trimStr(value)))
		}
	}
}


//@func 利用正则表达式压缩字符串，去除空格或制表符匹配一个或多个空白符的正则表达式
func trimStr(strs string) string {
	return strings.Trim(strings.Trim(strings.Trim(strs, " "), "\n"), "\r")
}


//@func 删除行收的助手符号# 和 /
func delComment(strs string) string {
	re3, _ := regexp.Compile("(^#{1,})|(^/{1,})")
	rep := re3.ReplaceAllString(strs, "");
	return trimStr(rep)
}


//@func 获取文件名
func TrimFile(fullFilename string) string {
	//获取文件名带后缀
	filenameWithSuffix := filepath.Base(fullFilename)
	//fmt.Println("filenameWithSuffix =", filenameWithSuffix)
	//获取文件后缀
	fileSuffix := path.Ext(filenameWithSuffix)
	//fmt.Println("fileSuffix =", fileSuffix)

	//获取文件名
	filenameOnly := strings.TrimSuffix(filenameWithSuffix, fileSuffix)
	return filenameOnly
}
