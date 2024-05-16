package main

import (
	"fmt"
	"time"
)

type data struct {
	value    int
	funcType int
}

func checkPrime(val int) bool {
	c := 0
	for i := 1; i <= val; i++ {
		if val%i == 0 {
			c++
		}
	}
	if c > 2 {
		return false
	} else {
		return true
	}
}

func Prime(st, en, idx int, ch chan data) {

	for i := st; i <= en; i++ {
		ans := checkPrime(i)
		if ans {
			da := data{
				value:    i,
				funcType: idx,
			}

			ch <- da
		}
	}
	close(ch)

}

func Channel() {

	averagePrimeValue, cnt := 0, 0

	pri := map[int]int{}
	ch1 := make(chan data)
	ch2 := make(chan data)
	ch3 := make(chan data)
	ch4 := make(chan data)

	start := time.Now()

	//ch := make(chan int)
	go Prime(2, 1000, 1, ch1)
	go Prime(2, 1000, 2, ch2)
	go Prime(2, 1000, 3, ch3)
	go Prime(2, 1000, 4, ch4)

	var v1, v2, v3, v4 bool
	for {

		select {
		case val, ok := <-ch1:
			if !ok {
				v1 = true
				break
			}

			_, ok = pri[val.value]
			if !ok {
				pri[val.value] = val.funcType
				averagePrimeValue += val.value
				cnt++
			}
		case val, ok := <-ch2:
			if !ok {
				v2 = true
				break
			}

			_, ok = pri[val.value]
			if !ok {
				pri[val.value] = val.funcType
				averagePrimeValue += val.value
				cnt++
			}
		case val, ok := <-ch3:
			if !ok {
				v3 = true
				break
			}

			_, ok = pri[val.value]
			if !ok {
				pri[val.value] = val.funcType
				averagePrimeValue += val.value
				cnt++
			}
		case val, ok := <-ch4:
			if !ok {
				v4 = true
				break
			}

			_, ok = pri[val.value]
			if !ok {
				pri[val.value] = val.funcType
				averagePrimeValue += val.value
				cnt++
			}

		}
		if v1 && v2 && v3 && v4 {
			break
		}

	}
	end := time.Since(start)

	fmt.Println("Total time taken is ", end)

	fmt.Println("Total prime number is ", cnt)
	fmt.Println("Sum of prime number is ", averagePrimeValue)

	var avg float64

	avg = float64(averagePrimeValue) / float64(cnt)

	fmt.Println("Average sum is ", avg)

}
