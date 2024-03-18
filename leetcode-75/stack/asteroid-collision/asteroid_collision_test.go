package removestarsfromstring

import (
	"testing"
)

func TestRemoveString(t *testing.T) {
	var results = map[string]string {
		"leet**cod*e": "lecoe",
		"erase*****": "",
	}

	for k, wanted := range results {
		if removeStars(k) != wanted {
			t.Errorf("wanted %s, got %s\n", wanted, removeStars(k))
		}
	}
}
