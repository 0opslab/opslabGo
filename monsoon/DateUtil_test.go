package monsoon

import (
	"fmt"
	"testing"
)

func TestNewGoTime(t *testing.T) {
	date := NewGoTime()
	fmt.Println(date.Now())
}