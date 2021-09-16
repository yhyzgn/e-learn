// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2021-08-31 14:34
// version: 1.0.0
// desc   :

package engine

import (
	"testing"
)

func TestNextIP(t *testing.T) {
	for i := 0; i < 20; i++ {
		t.Log(NextIP())
	}
}
