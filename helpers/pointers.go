package helpers

func Ptr[T any](val T) *T {
	return &val
}

func ArrToPtrArr[T any](arr []T) []*T {
	prtArr := make([]*T, len(arr))
	for i := range arr {
		prtArr[i] = Ptr(arr[i])
	}
	return prtArr
}
