package service

/**
 * @author: tangye
 * @Date: 2019/9/17 15:48
 * @Description:
 */

func StringSliceEqual(a, b []string) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
