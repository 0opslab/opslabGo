package monsoon

import (
	"fmt"
	"strings"
	"testing"
)

func TestRandom_String(t *testing.T) {
	rand := NewRandom()
	fmt.Println(rand.String(8,Numeric))
	fmt.Println(rand.String(32,Alphanumeric))


	dateUtil  := NewGoTime()
	igGen := NewRandomGen(strings.ReplaceAll(dateUtil.GetYmd(),"-",""),1000001)
	for i:=0;i< 1000 ;i++  {
		fmt.Println(igGen.Get())
	}


	for i:=0;i<1000;i++{
		fmt.Println(TimeUUID())
	}
}