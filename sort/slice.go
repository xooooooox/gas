package sort

// ReverseSlice reverse slice
func ReverseSlice(s []interface{}) []interface{} {
	length := len(s)
	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
