package funcs

import (
	"os"
	"fmt"
	"bufio"
	"io"
	"strings"
)

func Read_File_Line(filename string)(lines []string,err error) {

	fi, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return nil,err
	}
	defer fi.Close()

	lines = make([]string,0,10)
	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		var line = string(a)
		str := strings.Replace(line, "'", "\"", -1)
		str = strings.Replace(line, "\n", "", -1)
		lines = append(lines, str)
	}
	return lines,nil
}
