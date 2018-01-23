package demo

import (
	"os"
	"fmt"
)

func file_demo(){
	//文件路径
	userFile := "/tmp/test.txt"
	//根据路径创建File的内存地址
	fout,err := os.Create(userFile)
	//延迟关闭资源
	defer fout.Close()
	if err != nil{
		fmt.Println(userFile,err)
		return
	}
	//循环写入数据到文件
	for i:=0;i<10;i++{
		//写入字符串
		fout.WriteString("Hello world!\r\n")
		//强转成byte slice后再写入
		fout.Write([]byte("abcd!\r\n"))
	}

	//打开文件,返回File的内存地址
	fin,err := os.Open(userFile)
	//延迟关闭资源
	defer fin.Close()
	if err != nil{
		fmt.Println(userFile,err)
		return
	}
	//创建一个初始容量为1024的slice,作为缓冲容器
	buf := make([]byte,1024)
	for{
		//循环读取文件数据到缓冲容器中,返回读取到的个数
		n,_ := fin.Read(buf)

		if 0==n{
			break //如果读到个数为0,则读取完毕,跳出循环
		}
		//从缓冲slice中写出数据,从slice下标0到n,通过os.Stdout写出到控制台
		os.Stdout.Write(buf[:n])
	}
}
