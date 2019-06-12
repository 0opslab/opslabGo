package demo

import (
	"io/ioutil"
	"fmt"
)

func IOUtil_Demo() {
	file_name := "/local/workspace/opslabGo/data/tmp/go_file.txt"
	file_iout_name := "/local/workspace/opslabGo/data/tmp/go_file_ioutil_out.txt"

	//文件读取
	b, err := ioutil.ReadFile(file_name)
	if err != nil {
		fmt.Print(err)
	}
	str := string(b)
	fmt.Println("file-content:", str)

	//文件写入
	err = ioutil.WriteFile(file_iout_name, []byte(str), 0644)
	if err != nil {
		panic(err)
	}
}
