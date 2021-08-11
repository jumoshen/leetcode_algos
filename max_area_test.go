package algos

import (
	"testing"
)

func TestMaxArea(t *testing.T) {
	type test struct {
		input []int
		want  int
	}

	heights := map[string]test{
		"height1": {[]int{1, 2, 3}, 2},
		"height2": {[]int{7, 4, 3, 5, 6}, 24},
	}

	for key, height := range heights {

		if res := MaxArea(height.input); res != height.want {
			t.Errorf("name:%s,except:%d,got:%d", key, height.want, res)
		}
	}
}
