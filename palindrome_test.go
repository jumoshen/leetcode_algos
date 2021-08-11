package algos

import "testing"

func TestIsPalindrome(t *testing.T) {
	type test struct {
		input int
		want  bool
	}

	numbers := map[string]test{
		"is": {1324231, true},
		"not":{12345, false},
	}

	for k, v := range numbers{
		res := IsPalindrome(v.input)

		if res != v.want{
			t.Errorf("name%s, except:%v, got:%v", k, v.want, res)
		}
	}
}
