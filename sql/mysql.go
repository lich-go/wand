package sql

// Limit 计算分页偏移量
func Limit(Current int, Num int) (int, int) {
	if Current < 1 {
		Current = 1
	}
	if Num < 1 {
		Num = 10
	}
	offset := (Current - 1) * Num
	return Num, offset
}
