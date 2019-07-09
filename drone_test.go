package dronetest

import (
	"fmt"
	"testing"
)

func Test_Run(t *testing.T) {
	for i := 0; i < 100; i++ {
		c := i + 1
		fmt.Println(c)
	}
}
