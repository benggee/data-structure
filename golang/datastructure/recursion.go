package datastructure

type recursion struct {
}

func Recursion() *recursion {
	return &recursion{}
}

func (r *recursion) Sum(arr []int) int {
	return r.sum(arr, 0)
}

func (r *recursion) sum(arr []int, level int) int {
	if (level == len(arr)) {
		return 0
	}
	return arr[level] + r.sum(arr, level+1)
}

