package tools

import (
	"fmt"
	"testing"
)

func TestFormatCmd(t *testing.T) {
	cmd := "dis vers,];dis users"
	fmt.Println(FormatCmd(cmd))
}
