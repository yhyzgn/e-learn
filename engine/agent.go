// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2021-08-31 14:33
// version: 1.0.0
// desc   :

package engine

import (
	"go.zoe.im/surferua"
)

// NextUserAgent ...
func NextUserAgent() string {
	return surferua.New().String()
}
