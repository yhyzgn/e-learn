// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2021-08-31 14:34
// version: 1.0.0
// desc   :

package engine

import (
	"math/rand"
	"testing"
)

func TestUserAgent(t *testing.T) {
	for i := 0; i < 20; i++ {
		t.Log(NextUserAgent())
	}

	t.Log(rand.Intn(10) + 1)
	t.Log(rand.Intn(10) + 1)
	t.Log(rand.Intn(10) + 1)
	t.Log(rand.Intn(10) + 1)
}
