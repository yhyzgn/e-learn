// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2021-06-06 14:19
// version: 1.0.0
// desc   :

package util

// SliceUnion 并集
func SliceUnion(slice1, slice2 []string) []string {
	flag := make(map[string]int)
	for _, v := range slice1 {
		flag[v]++
	}
	for _, v := range slice2 {
		if flag[v] == 0 {
			slice1 = append(slice1, v)
		}
	}
	return slice1
}

// SliceIntersect 求交集
func SliceIntersect(slice1, slice2 []string) []string {
	flag := make(map[string]int)
	for _, v := range slice1 {
		flag[v]++
	}
	res := make([]string, 0)
	for _, v := range slice2 {
		if flag[v] > 0 {
			res = append(res, v)
		}
	}
	return res
}

// SliceDifference 求差集
func SliceDifference(slice1, slice2 []string) []string {
	flag := make(map[string]int)
	// 先求交集
	inter := SliceIntersect(slice1, slice2)
	for _, v := range inter {
		flag[v]++
	}

	// 不在交集里的，才是差异
	res := make([]string, 0)
	for _, v := range slice2 {
		if flag[v] == 0 {
			res = append(res, v)
		}
	}
	return res
}
