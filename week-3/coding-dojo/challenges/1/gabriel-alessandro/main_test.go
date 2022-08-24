package main

import (
	"fmt"
	"testing"
)

func TestFifo(t *testing.T) {

	cases := map[string]struct {
		output func()
	}{
		"test": {
			function: func() { fmt.Println("1") },
		},
	}

	for _, cs := range cases {
		fifoImpl := fifoImpl()
		data := fifoImpl.Pop()
		if &data != &cs.output {
			t.Error()
		}
	}

}
