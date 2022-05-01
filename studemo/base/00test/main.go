package main

import (
	"fmt"
	"strings"
)

func main() {
	strs := "https://m.baidu.com/s?word=%!E(MISSING)7%!B(MISSING)A%!A(MISSING)2%!E(MISSING)8%!A(MISSING)D%!A(MISSING)6%!E(MISSING)7%!B(MISSING)%!B(MISSING)8%!E(MISSING)5%B3%!E(MISSING)6%!E(MISSING)%!A(MISSING)8%!E(MISSING)8%!D(MISSING)%!"
	fmt.Println(!strings.Contains(strs,"(MISSING)"))
}
