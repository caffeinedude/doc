package main

import (
	"fmt"
	"math/rand"
	"time"
)

//origin原数组 count 随机取出的个数，最终返回一个count容量的目标数组

func MicsSlice(origin []int64, count int) []int64 {
	tmpOrigin := make([]int64, len(origin))
	copy(tmpOrigin, origin)
	//一定要seed
	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(tmpOrigin), func(i int, j int) {
		tmpOrigin[i], tmpOrigin[j] = tmpOrigin[j], tmpOrigin[i]
	})

	result := make([]int64, 0, count)
	for index, value := range tmpOrigin {
		if index == count {
			break
		}
		result = append(result, value)
	}
	return result
}
func main() {

	dest := MicsSlice([]int64{1, 2, 3, 4, 5, 6, 7}, 2)
	fmt.Println("result....", dest)
}
