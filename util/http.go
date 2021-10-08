// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2019-11-04 下午12:32
// version: 1.0.0
// desc   :

package util

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
)

// RecycleRequestBody 复用 request.Body
//
// 获取到本来的 request.Body
// 再把获取到的设置回去
func RecycleRequestBody(req *http.Request) []byte {
	if req != nil && req.Body != nil {
		bs, _ := ioutil.ReadAll(req.Body)
		req.Body = ioutil.NopCloser(bytes.NewBuffer(bs))
		return bs
	}
	return nil
}

// SetRequestHeader 设置请求头
//
// 指定大小写
func SetRequestHeader(req *http.Request, key, value string) {
	req.Header.Set(key, value)
}

// SetResponseHeader 设置响应头
//
// 指定大小写
func SetResponseHeader(res *http.Response, key, value string) {
	res.Header.Set(key, value)
}

// SetResponseWriterHeader 设置响应头
//
// 指定大小写
func SetResponseWriterHeader(res http.ResponseWriter, key, value string) {
	res.Header().Set(key, value)
}

// AddURLQuery 向 url 中添加 query 参数
//
// 添加 url 参数
func AddURLQuery(url, key, value string) string {
	var sb strings.Builder
	sb.WriteString(url)

	if strings.Contains(url, "?") {
		// 如果不以 ? 结尾，也不以 & 结尾，就加上 & 连接符
		if !strings.HasSuffix(url, "?") && !strings.HasSuffix(url, "&") {
			sb.WriteString("&")
		}
	} else {
		sb.WriteString("?")
	}
	sb.WriteString(key)
	sb.WriteString("=")
	sb.WriteString(value)
	return sb.String()
}

// IsCorsRequest 是否是跨域请求
func IsCorsRequest(r *http.Request) bool {
	return r.Header.Get("Origin") != ""
}

// ShouldAbortRequest 是否应该终止跨域请求
func ShouldAbortRequest(r *http.Request) bool {
	// Access-Control-Request-Method 出现于 Options 预检请求头中
	return IsCorsRequest(r) && r.Method == http.MethodOptions && r.Header.Get("Access-Control-Request-Method") != ""
}
