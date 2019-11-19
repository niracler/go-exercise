package infra

import (
	"math/rand"
	"time"
)

//二倍均值算法
func DoubleAverage(count int64, amount int64) int64 {
	const min = 1

	if count == 1 {
		return amount
	}

	max := amount - min*count //最大可用金额
	avg := max / count        //最大可用平均值
	avg2 := 2*avg + min

	rand.Seed(time.Now().UnixNano())
	x := rand.Int63n(avg2) + min

	return x
}

//基础资源上下文结构体
type StarterContext map[string]interface{}

type BaseStarter struct {
}

//func main() {
//	count, amount := int64(100), int64(100)
//	remain := amount
//	sum := int64(0)
//
//	for i := int64(0); i < count; i++ {
//		x := DoubleAverage(count-i, remain)
//		remain -= x
//		sum += x
//		fmt.Println(i+1, "=", float64(x)/float64(100))
//	}
//	fmt.Println("总和", float64(sum)/float64(100))
//}
