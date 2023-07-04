package tools

type Int64Sort []int64

func (a Int64Sort) Len() int           { return len(a) }
func (a Int64Sort) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Int64Sort) Less(i, j int) bool { return a[i] < a[j] }
