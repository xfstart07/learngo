// Author: Xu Fei
// Date: 2018/9/27
package main

import (
	"time"
	"math/rand"
	"fmt"
	"math/big"

	crand "crypto/rand"
	"sync"
)

func genRand(wp *sync.WaitGroup, num int) {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(time.Now().UnixNano())
	//for i := 0; i < 100; i++ {
	//	x := rand.Intn(1000000)
	//	fmt.Printf("Rand 第 %d 个 %d\n", num, x)
	//}
	wp.Done()
}

func genCRand(wp *sync.WaitGroup, num int) {
	for i := 0; i < 100; i++ {
		x, _ := crand.Int(crand.Reader, big.NewInt(1000000))
		fmt.Printf("CRand 第 %d 个 %d\n", num, x)
	}
	wp.Done()
}

func main() {

	wp :=  &sync.WaitGroup{}
	for i:= 0; i<100; i++ {
		wp.Add(1)
		go genRand(wp, i)
	}

	//for i:= 0; i<100; i++ {
	//	wp.Add(1)
	//	go genCRand(wp, i)
	//}

	wp.Wait()
}
