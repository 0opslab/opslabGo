package demo

import (
	"bufio"
	"io"
	"os"
)

func Bufio_Dmoe() {
	file_name := "/local/workspace/opslabGo/data/tmp/go_file.txt"
	file_out_name := "/local/workspace/opslabGo/data/tmp/go_file_buf_out.txt"
	fi, err := os.Open(file_name)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	//创建一个读取缓冲流
	r := bufio.NewReader(fi)

	fo, err := os.Create(file_out_name)
	if err != nil {
		panic(err)
	}
	//创建输出缓冲流
	w := bufio.NewWriter(fo)

	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}

		if n2, err := w.Write(buf[:n]); err != nil {
			panic(err)
		} else if n2 != n {
			panic("error in writing")
		}
	}
	if err = w.Flush(); err != nil {
		panic(err)
	}

}