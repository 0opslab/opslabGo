package stringsdemo

import (
	"fmt"
	"strings"
)

//给 fmt.Println 一个短名字的别名
var p = fmt.Println


func demo() {
	fmt.Println(strings.Count("five", ""))
	p("Contains:  ", strings.Contains("test", "es"))
	p("Count:     ", strings.Count("test", "t"))
	p("HasPrefix: ", strings.HasPrefix("test", "te"))
	p("HasSuffix: ", strings.HasSuffix("test", "st"))
	p("Index:     ", strings.Index("test", "e"))
	p("Join:      ", strings.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", strings.Repeat("a", 5))
	p("Replace:   ", strings.Replace("foo", "o", "0", -1))
	p("Replace:   ", strings.Replace("foo", "o", "0", 1))
	p("Split:     ", strings.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", strings.ToLower("TEST"))
	p("ToUpper:   ", strings.ToUpper("test"))
	/*
		p := point{1, 2}
		fmt.Printf("%v\n", p)

		fmt.Printf("%+v\n", p)

		fmt.Printf("%#v\n", p)

		fmt.Printf("%T\n", p)

		fmt.Printf("%t\n", true)

		fmt.Printf("%d\n", 123)

		fmt.Printf("%b\n", 14)

		fmt.Printf("%c\n", 33)

		fmt.Printf("%x\n", 456)

		fmt.Printf("%f\n", 78.9)

		fmt.Printf("%e\n", 123400000.0)
		fmt.Printf("%E\n", 123400000.0)

		fmt.Printf("%s\n", "\"string\"")

		fmt.Printf("%q\n", "\"string\"")

		fmt.Printf("%x\n", "hex this")

		fmt.Printf("%p\n", &p)

		fmt.Printf("|%6d|%6d|\n", 12, 345)

		fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)

		fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

		fmt.Printf("|%6s|%6s|\n", "foo", "b")

		fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

		s := fmt.Sprintf("a %s", "string")
		fmt.Println(s)

		fmt.Fprintf(os.Stderr, "an %s\n", "error")
	*/
}
