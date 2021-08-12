package algos

import "testing"

func TestRemoveElement(t *testing.T) {
	type test struct {
		input []int
		element int
		want  int
	}

	tests := map[string]test{
		"test1": {input: []int{1, 2, 2, 3, 4, 5, 5, 6}, element: 2, want: 6},
		"test2": {input: []int{2, 2}, element: 2, want: 0},
	}

	for k, test := range tests {
		res := RemoveElement(test.input, test.element)

		if res != test.want {
			t.Errorf("name:%s,except:%d, got:%d", k, test.want, res)
		}
	}
}
