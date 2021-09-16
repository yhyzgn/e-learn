// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2019-11-13 上午11:17
// version: 1.0.0
// desc   :

package util

import (
	"fmt"
	"testing"
)

func TestAddURLQuery(t *testing.T) {
	url := "/test"

	url = AddURLQuery(url, "id", "34")
	fmt.Println(url)

	url = AddURLQuery(url, "name", "Jason")
	fmt.Println(url)

}
