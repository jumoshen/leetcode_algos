package algos

import (
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	type test struct {
		input []int
		want  int
	}

	tests := map[string]test{
		"test1": {input: []int{1, 2, 2, 3, 4, 5, 5, 6}, want: 6},
		"test2": {input: []int{2, 2}, want: 1},
	}

	for k, test := range tests {
		res := RemoveDuplicates(test.input)

		if res != test.want {
			t.Errorf("name:%s,except:%d, got:%d", k, test.want, res)
		}
	}
}
