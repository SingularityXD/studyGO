package main

import (
	"fmt"
	"math/rand"
)

func RandInt64(min, max int64) int64 {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Int63n(max-min) + min
}
func main() {
	balance := [...]string{"Space Adventures", "SpaceX", "Virgin Galactic"}
	count := 4
	fmt.Printf("%-10v     %-10v     %-10v     %-10v\n", "太空航行公司", "飞行天数", "飞行类型", "价格(百万美元)")
	for count >= 0 {
		co := balance[rand.Intn(len(balance))]
		const distance = 62100000
		speedRand := RandInt64(16, 31)
		speed := speedRand * 60 * 60 * 24
		days := distance / speed
		flyTypeRand := RandInt64(1, 3)
		flyType := "单程"
		if flyTypeRand > 1 {
			flyType = "往返"
		}
		singleWayPrice := speedRand*100 + 2000
		fmt.Printf("%-10v     %-10v     %-10v     %-10v\n", co, days, flyType, singleWayPrice)
		count--
	}
}
