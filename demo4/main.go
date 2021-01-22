package main

import (
	"fmt"
	"sync"
)

func main() {
	result := DistributeCompute([]int{2, 5, 6, 9})
	fmt.Println(result)
}

func DistributeCompute(params []int) []int {
	paramsCount := len(params)
	type resStruct struct {
		idx int
		res int
	}
	resChan := make(chan resStruct, paramsCount) //？这里通道的容量能否小于 paramsCount？为什么

	var wg sync.WaitGroup
	wg.Add(paramsCount)

	for i, num := range params {
		//wg.Add(1) //能否用这种写法替代上面的 wg.Add(paramsCount)
		go func(i, num int, reschan chan<- resStruct) { //这里携程执行的函数，是否可以去掉第三个参数？第三个参数的意义是什么
			defer wg.Done()
			resChan <- resStruct{i, num * num}
		}(i, num, resChan)
	}

	wg.Wait()      //这里使用 wg 是否是必须的，不用会怎样？
	close(resChan) //如果没有这行会怎么样

	result := make([]int, paramsCount, paramsCount) //注意这里长度要赋值
	for res := range resChan {
		result[res.idx] = res.res
	}

	return result
}
