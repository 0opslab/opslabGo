package monsoon

import (
	"fmt"
	"testing"
)

func TestGetLogicalDrives_Windows(t *testing.T) {
	drivers  := GetLogicalDrives()
	fmt.Println(drivers)
}
