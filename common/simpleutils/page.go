package simpleutils

func CalcTotalPage(pageSize, total int32) int32 {
	return (total + pageSize - 1) / pageSize
}
