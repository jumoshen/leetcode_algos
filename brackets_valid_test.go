package algos

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	type test struct {
		input string
		want  bool
	}

	tests := map[string]test{
		"test1": {input: "{}()[]", want: true},
		"test2": {input: "{()}[]", want: true},
	}

	for k, test := range tests {
		res := IsValid(test.input)

		if res != test.want {
			t.Errorf("name:%s,except:%t, got:%t", k, test.want, res)
		}
	}
}
