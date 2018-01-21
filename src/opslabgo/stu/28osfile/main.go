package main

import "fmt"
import (
	"os"
	"bufio"
	"io/ioutil"
	"io"
	"path/filepath"
)

func main()  {
	file_name :="/local/workspace/opslabGo/data/tmp/go_file.txt"
	err := write(file_name)
	if err != nil{
		fmt.Println(file_name,err)
	}

	file_name1 :="/local/workspace/opslabGo/data/tmp/go_file1.txt"
	err0 := write(file_name1)
	if err0 != nil{
		fmt.Println(file_name1,err0)
	}

	file_dst := "/local/workspace/opslabGo/data/tmp/go_file_dst.txt"
	rst,errcopy := copyFile(file_name,file_dst)
	if errcopy != nil{
		fmt.Println("copy error:",errcopy)
	}
	fmt.Println("copy rst:",rst)

	err1 := readByByte(file_name)
	if err != nil{
		fmt.Println(file_name,err1)
	}

	err2 := readByLine(file_name)
	if err != nil{
		fmt.Println(file_name,err2)
	}

	str,err3 := readAll(file_name)
	if err3 != nil{
		fmt.Println(file_name," ReadAll error:",err3)
	}
	fmt.Println("file content:",str)

	//遍历目录
	file_list("/local/workspace/opslabGo/data/")
}

/**
 写文件
 */
func write(file_name string) error{
	fout,err := os.Create(file_name)
	defer  fout.Close()

	if err != nil{
		return err
	}

	for i:=0;i<10 ;i++  {
		fout.WriteString("0opslab.com\r\n")
		fout.Write([]byte("hello man"))
	}
	return nil
}

func writeString(file_name string){
	file,err := os.OpenFile(file_name,os.O_CREATE|os.O_WRONLY,0755)
	if err != nil{
		fmt.Println("error",err)
		return
	}
	defer file.Close()

	fileWriter := bufio.NewWriter(file)
	for i:=0;i<10 ;i++  {
		fileWriter.WriteString("good")
	}
	fileWriter.Flush()
}
/**
 读文件
 */
func readAll(file_name string)(string,error){
	file,err := os.OpenFile(file_name,os.O_CREATE|os.O_RDWR,0666)
	if nil == err{
		defer file.Close()

		buf,err := ioutil.ReadAll(file)
		if nil == err{
			return string(buf),nil
		}
		return "",err
	}else{
		return "",err
	}
}

func readByByte(file_name string) error{
	fin,err := os.Open(file_name)
	defer fin.Close()

	if err != nil{
		return err
	}

	buf := make([]byte,1024)
	for{
		n,_ := fin.Read(buf)
		if 0 == n {
			break
		}
		fmt.Println(buf[:n])
		fmt.Println(string(buf[:n]))
	}
	return nil
}
func readByLine(file_name string) error{
	fin,err := os.Open(file_name)
	defer fin.Close()

	if err == nil{
		buff := bufio.NewReader(fin)
		for{
			//以\r\n为结束符读入一行
			//此处如果使用ReadLine()方法更加完美
			line,err := buff.ReadString('\n')
			if err != nil{
				return err
			}
			fmt.Println("read line =>",line)
		}
	}else{
		return err
	}

}

func copyFile(src_file_name string,dst_file_name string)(writeen int64,err error){
	src,err := os.Open(src_file_name)
	if err != nil{
		return -1,err
	}
	defer src.Close()

	dst,err := os.OpenFile(dst_file_name,os.O_CREATE|os.O_WRONLY,0644)
	if err != nil{
		return -1,err
	}
	defer dst.Close()

	return io.Copy(dst,src)
}

/**
 遍历文件夹
 */
func file_list(path string){
	err := filepath.Walk(path,func(path string,f os.FileInfo,err error) error{
		if(f ==nil){return err}
		if f.IsDir(){return nil}
		println(path)
		return nil
	})
	if err != nil{
		fmt.Println("filepath.Walt() return %v\n",err)
	}
}