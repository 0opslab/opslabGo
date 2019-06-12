package main

import (
	"bytes"
	"io"
	"os"

)

func main(){
	proverbs := new(bytes.Buffer)
	proverbs.WriteString("Channels orchestrate mutexes serialize\n")
	proverbs.WriteString("Cgo is not Go\n")
	proverbs.WriteString("Errors are values\n")
	proverbs.WriteString("Don't panic\n")

	piper, pipew := io.Pipe()

	// 将 proverbs 写入 pipew 这一端
	go func() {
		defer pipew.Close()
		io.Copy(pipew, proverbs)
	}()

	// 从另一端 piper 中读取数据并拷贝到标准输出
	io.Copy(os.Stdout, piper)
	piper.Close()

}
