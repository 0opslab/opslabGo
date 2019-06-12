package monsoon

import (
	"fmt"
	"os"
	"path"
	"testing"
)

func TestWriteString(t *testing.T) {
	WriteString("C:/Users/Administrator/Desktop/test1/test2/test.txt", "1111111111/n2222222222//33333")
	strs := "如果某个东西长得像鸭子，像鸭子一样游泳，像鸭子一样嘎嘎叫，那它就可以被看成是一只鸭子。"
	WriteBytes("C:/Users/Administrator/Desktop/test1.txt", []byte(strs))
}

func TestReadByte(t *testing.T) {
	file := "C:/Windows/System32/drivers/etc/hosts"
	bytes, err := ReadByte(file,8)
	if err != nil {
		fmt.Println("ReadError:" + err.Error())
	}
	fmt.Println("==============>")
	fmt.Println(bytes,"====>",bytesToHexString(bytes),"===>",string(bytes))
}


func TestReadFile(t *testing.T) {
	file := "C:/Windows/System32/drivers/etc/hosts"
	bytes, err := ReadFile(file)
	if err != nil {
		fmt.Println("ReadError:" + err.Error())
	}
	fmt.Print(string(bytes))
}

func TestFileInfo(t *testing.T) {
	file := "C:/Windows/System32/drivers/etc/hosts"
	if FileIsExist(file) {
		fileSize, _ := FileSize(file)
		fileMd5, _ := Md5File(file)
		fileSH1, _ := Sha1File(file)
		fileSH2, _ := Sha256File(file)
		fileSH5, _ := Sha512File(file)

		fmt.Printf("FileInfo:fileSize=%v fileMd5=%s fileSH1=%s fileSH2=%s fileSH5=%s/n", fileSize, fileMd5, fileSH1, fileSH2, fileSH5)
	}

}

func TestWalkDirFiles(t *testing.T) {
	path := "c:/workspace/"
	files, dirs, _ := WalkDirFiles(path, "java")
	for _, file := range files {
		fmt.Println(file)
	}
	for _, dir := range dirs {
		fmt.Println(dir)
	}
}

func TestWalkDirFilesFilter(t *testing.T) {
	files, _, _ := WalkDirFilesFilter("c:/workspace/opslabJava", func(filename string) bool {
		fi, e := os.Stat(filename)
		if e != nil {
			return false
		}

		if (fi.IsDir()) {
			return false
		} else {
			if (path.Ext(filename) == ".java") {
				return true
			}
		}
		return false
	})
	for _, file := range files {
		fmt.Println(file)
	}
}

// 遍历指定文件夹并计算其中文件的MD5进行输出
func TestWalkDirFilesHandler(t *testing.T) {
	fileMd5 := func(fileName string) {
		if(IsFile(fileName)){
			md5,_ := Md5File(fileName)
			fmt.Println(fileName,"==========>",md5)
		}

	}
	WalkDirFilesHandler("c:/workspace/opslabJava",fileMd5)
}

func TestCopy(t *testing.T) {
	if res,err := CopyFile("C:/Windows/System32/drivers/etc/hosts","C:/Users/Administrator/Desktop/hosts"); err != nil{
		fmt.Println("CopyFile Error:",err)
	}else{
		fmt.Println("CopyFile status:",res)
	}

	if _, err := CopyDir("C:/workspace/doc/", "C:/Users/Administrator/Desktop/doc/"); err != nil {
		fmt.Println(err)
	}
}


func TestZip(t *testing.T){
	if err := ZipCompress("C:/workspace/doc","C:/Users/Administrator/Desktop/11.zip"); err != nil{
		fmt.Println("压缩文件错误:",err)
	}
	if err := ZipUnCompress("C:/Users/Administrator/Desktop/11.zip","C:/Users/Administrator/Desktop/docs"); err != nil{
		fmt.Println("压缩文件错误:",err)
	}

}