package utils

/**
判断两个数组中的数据是否一致
*/
func CompareArrayAllValue(arr1 []uint8, arr2 []uint8) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}
