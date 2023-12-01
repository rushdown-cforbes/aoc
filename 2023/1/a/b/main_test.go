package main

import "testing"

func TestCheckNum(t *testing.T) {
	cases := []struct {
		input  []byte
		output string
	}{
		{[]byte("one"), "1"},
		{[]byte("oone"), "1"},
		{[]byte("ontwoe"), "2"},
	}
	for caseNum, c := range cases {
		output := ""
		for i := range c.input {
			num, valid := checkNumber(c.input, i)
			if valid {
				output = num
			}
		}
		if output != c.output {
			t.Errorf("output number not correct case %d: %s %s", caseNum, output, c.output)
		}
	}
}
