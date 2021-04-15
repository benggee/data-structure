package tool


func VerifyOrder(arr []int64) bool {
	aLen := len(arr)
	// 验证是否正确
	for i := 0; i < aLen-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}