package models

type FrequentNum struct {
	Num   int
	Count int
}

func NewFrequentNum(num int, count int) FrequentNum {
	return FrequentNum{
		Num:   num,
		Count: count,
	}
}
