package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"time"
	"sync"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

// Put a value in the top list if it deserves to be in it
func meta(n int, h []int) []int {
	if n > h[0] {
		h = append(h, n)
	} else if c := contains(h, n); n > h[len(h)-1] && !c {
			h = append(h, n)
	}
	return h
}

// Sort the top list and cut out the surplus
func sortAndSize (h []int, n int) []int {
	sort.Sort(sort.Reverse(sort.IntSlice(h)))
	// Remove the last if needed
	if len(h) > n {
		h = h[:len(h)-1]
	}
	return h
}

// Verify if a number is in a slice
func contains(s []int, e int) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

func populateQueue(c chan int, fp string, p int) {

	f, _ := os.Open(fp)
	var wg sync.WaitGroup
	wg.Add(p)

	for ; p > 0; p-- {
		go func() {
			defer wg.Done()
			i:= 0
			s := bufio.NewScanner(f)
			for s.Scan() {
				cn, _ := strconv.Atoi(s.Text())
				c <- cn
				if i == 0 {
					fmt.Println(s.Text())
				}
				i++
			}
			if err := s.Err(); err != nil {
				fmt.Println(os.Stderr, "reading standard input:", err)
			}
			fmt.Printf("On queue: %v\n", i)
			fmt.Printf("Fila completa\n")
		}()
	}
	go func() {
		wg.Wait()
		close(c)
	}()
}

func putOnQueue(cn int, c chan int) {
	c<- cn
}

// func topN(c chan int, n int, h []int) []int {
func topN(c chan int, n, p int) (/*<-chan */[]int) {
	h := []int{0}
	cc := make(chan []int)
	var mutex = &sync.Mutex{}
	for ; p > 0; p-- {
		go func() {
			i := 0
			for cn := range c {
				mutex.Lock()
				h = sortAndSize(meta(cn, h), n)
				mutex.Unlock()
				i++
			}
			fmt.Printf("Interactions: %d\n", i)
			// fmt.Printf("Partial Result: %v\n", h)
			cc <- h

		}()
	}
	return <-cc
}

// func fanIn(c1,c2 chan int) []int {
// 	// cc := make(chan []int)
// 	// for p > 0; p-- {
// 	// 	h, i := <-topN(c, n)
// 	// 	fmt.Printf("H%d: %v - interactions:%v\n", i, h, i)
// 	// }
// 	h1, i1 := topN(c, n)
// 	// h1, i1 := topN(c, n)
// 	// h2, i2 := topN(c, n)
// 	// fmt.Printf("H1: %v - interactions:%v\n", h1, i1)
// 	// fmt.Printf("H2: %v - interactions:%v\n", h2, i2)

// 	for n := range h2 {
// 		h1 = meta(n, h1)
// 	}
// 	return h1
// } 


func main() {
	start := time.Now()

	var n, p int
	var fp string
	// c := make(chan int)
	c := make(chan int, 100000)

	switch a := len(os.Args); a {
	case 1:
		fp = "/home/alejdg/Workspace/half.txt"
		p = 1
		n = 10
	case 2:
		fp = os.Args[1]
		p = 1
		n = 10
	case 3:
		fp = os.Args[1]
		p, _ = strconv.Atoi(os.Args[2])
		n = 10
	default:
		fp = os.Args[1]
		p, _ = strconv.Atoi(os.Args[2])
		n, _ = strconv.Atoi(os.Args[3])
	}
	// h := []int{0}
	// f, err := os.Open(fp)
	// check(err)

	go populateQueue(c, fp, p)

	// h, i := topN(c, n, p)
	h := topN(c, n, p)


	elapsed := time.Since(start)
	fmt.Printf("Result: %v\n", h)
	fmt.Printf("Executed in %v\n", elapsed)
	// fmt.Printf("Iteractions: %v\n", i)

}

