package demo

import (
	"os"
	"bufio"
	"fmt"
	"io"
	"strings"
)

func writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

func load_file(filename string, list *[]string) {
	fi, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		var line = string(a)
		str := strings.Replace(line, "'", "\"", -1)
		str = strings.Replace(line, "\n", "", -1)
		*list = append(*list, str)
	}

	fmt.Println("init list. len", len(*list))
}


//func main(){
//	prose_list := make([]string,0)
//	load_file("c:/data.json",&prose_list)
//
//
//	writeLines(prose_list,"c:/data1.txt")
//
//}