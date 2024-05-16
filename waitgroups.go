package main

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

var wg sync.WaitGroup
var mut sync.Mutex

var primes []int

func checkPrimeW(val int) bool {
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

func PrimeW(st, en int) {

	defer wg.Done()

	for i := st; i <= en; i++ {
		ans := checkPrimeW(i)
		if ans {

			mut.Lock()

			primes = append(primes, i)

			mut.Unlock()

		}
	}
}

func WaitGroup() {

	start := time.Now()

	wg.Add(4)

	go PrimeW(2, 250)
	go PrimeW(251, 500)
	go PrimeW(501, 750)
	go PrimeW(751, 1000)

	wg.Wait()
	end := time.Since(start)

	sort.Slice(primes, func(i, j int) bool {
		return primes[i] < primes[j]
	})

	fmt.Println(primes)

	fmt.Println("Total number of prime is ", len(primes))

	fmt.Println("Total time taken is ", end)

}
