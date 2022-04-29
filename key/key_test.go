package key

import (
	"fmt"
	"testing"
)

func TestShaConvert(t *testing.T) {
	a := ShaConvert("ramin")

	fmt.Printf("%x\n", a)
	if false {
		t.Errorf("error")
	}
}
