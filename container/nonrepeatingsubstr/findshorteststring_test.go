package nonrepeatingsubstr

import (
	"testing"
)

func TestSubStr(t *testing.T)  {
	tests := []struct{
		longStr string
		ansInt int
	}{
		// Normal cases
		{"abcabcbb", 3},
		{"pwwkew", 3},

		// Edge cases
		{"", 0},
		{"b", 1},
		{"bbbbbbbbb", 1},
		{"abcabcabcd", 4},

		// Chinese support
		{"这里是慕课网", 6},
		{"一二三二一", 3},
		{"黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花", 8},
	}

	for _,tt := range tests {
		if actual := lengthOfNoRepeatingSubstring(tt.longStr); actual != tt.ansInt{
			t.Errorf("got %d for input %s; "+
				"expected %d",
				actual, tt.longStr, tt.ansInt)
		}
	}
}

func BenchmarkSubStr(b *testing.B)  {
	longStr := "黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"
	ansInt := 8
	for i :=0; i < b.N; i++ {
		if actual := lengthOfNoRepeatingSubstring(longStr); actual != ansInt{
			b.Errorf("got %d for input %s; "+
				"expected %d",
				actual, longStr, ansInt)
		}
	}
}
