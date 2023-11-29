package JOmap

import (
	"testing"
)

func TestSA_JoinKey(t *testing.T) {
	sa := SA{
		"name": "rama",
		"age":  "20",
	}
	sa2 := SA{
		"rama": "name",
		"20":   "age",
	}

	sa.Merge(sa2)
	t.Log(sa)
	t.Log(sa.JoinKey("", "asd"))
}
