package monsoon

import (
	"fmt"
	"strings"
	"testing"
)

func TestRsaEncryptWithPublic(t *testing.T) {
	_, e := RsaEncryptFileWithPublic("c:/ReadMe.md", "c:/1.data")
	if e != nil{
		fmt.Println(e)
	}

	_, e1 := RsaDecryptFileWithPrivte("c:/1.data","c:/ReadMe1.md")
	if e1 != nil{
		fmt.Println(e1)
	}
}

func TestRsaDecryptWithPrivte(t *testing.T) {
	rsaInfo := "rsainfo"
	rsaPath := "D:\\test\\rsafile"
	fileMd5 := func(fileName string) {
		if(IsFile(fileName)){
			sha,_ := Sha256File(fileName)

			rsaFile := rsaPath+"\\"+sha
			dstFile := strings.ReplaceAll(strings.ReplaceAll(fileName,"C:\\workspace\\","D:\\test\\dst\\"),"\\","/")
			rest := fmt.Sprintf("%s -> %s -> %s\n",fileName,rsaFile,dstFile)
			WriteString(rsaPath+"\\"+rsaInfo,rest)
			fmt.Print(rest)
			go func(){
				_, e := RsaEncryptFileWithPublic(fileName,rsaFile)
				if e != nil{
					fmt.Println(e)
				}
				_, e1 := RsaDecryptFileWithPrivte(rsaFile,dstFile)
				if e1 != nil{
					fmt.Println(e1)
				}
			}()

		}

	}
	WalkDirFilesHandler("C:\\workspace\\",fileMd5)
}
