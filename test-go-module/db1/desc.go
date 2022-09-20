package db1

func Addanddesc(a int ,b int) int {
	r := Add(a,b)
	return r - b
}