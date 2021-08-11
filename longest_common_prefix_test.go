package algos

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	type test struct {
		input []string
		want  string
	}

	tests := map[string]test{
		"strs-1": {
			input: []string{"abcdefg", "abcde", "abcrf"},
			want:  "abc",
		},
		"strs-2": {
			input: []string{"good", "god", "gift"},
			want:  "g",
		},
	}

	for key, test := range tests {
		res := LongestCommonPrefix(test.input)

		if res != test.want {
			t.Errorf("name:%s, except:%s, got:%s", key, test.want, res)
		}
	}
}
