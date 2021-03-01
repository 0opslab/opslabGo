package monsoon

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

//import (
//	"strconv"
//	"syscall"
//)
//
//
////windows获取系统盘符
//func GetLogicalDrives() []string {
//	kernel32 := syscall.MustLoadDLL("kernel32.dll")
//	GetLogicalDrives := kernel32.MustFindProc("GetLogicalDrives")
//	n, _, _ := GetLogicalDrives.Call()
//	s := strconv.FormatInt(int64(n), 2)
//	var drives_all = []string{"A:", "B:", "C:", "D:", "E:", "F:", "G:", "H:", "I:",
//		"J:", "K:", "L:", "M:", "N:", "O:", "P：", "Q：", "R：", "S：", "T：", "U：",
//		"V：", "W：", "X：", "Y：", "Z："}
//	temp := drives_all[0:len(s)]
//	var d []string
//	for i, v := range s {
//		if v == 49 {
//			l := len(s) - i - 1
//			d = append(d, temp[l])
//		}
//	}
//	var drives []string
//	for i, v := range d {
//		drives = append(drives[i:], append([]string{v}, drives[:i]...)...)
//	}
//	return drives
//}

/**
	获取程序当前所在目录
 */
func GetCurrentDirectory() string {
	//返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	//将\替换成/
	return strings.Replace(dir, "\\", "/", -1)
}
