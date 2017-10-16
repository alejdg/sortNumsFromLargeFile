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

func populateQueue(c chan int, fp string) {

	f, _ := os.Open(fp)

	// Scan the file line by line to avoid putting the whole file in memory
	s := bufio.NewScanner(f)
	i:= 0
	for s.Scan() {
		cn, _ := strconv.Atoi(s.Text())
		putOnQueue(cn, c)
		i++
	}
	close(c)
	fmt.Printf("On queue: %v\n", i)
	fmt.Printf("Fila completa\n")
	if err := s.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func putOnQueue(cn int, c chan int) {
	c<- cn
}

func readFromQueue(c chan int) int {
	n := <- c
	return n
}

// func topN(c chan int, n int, h []int) []int {
func topN(c chan int, n, p int) (/*<-chan */[]int) {
	// j := 0
	h := []int{0}
	// h := make([]int, 0)
	fmt.Printf("TopN: %v\n", n)
	cc := make(chan []int)
	var mutex = &sync.Mutex{}
	for j := 0; p > 0; p-- {
		go func() {
			i := 0
			for range c {
				mutex.Lock()
				h = sortAndSize(meta(readFromQueue(c), h), n)
				mutex.Unlock()
				i++
			}
			fmt.Printf("Interactions[%d]: %d\n", j, i)
			j++
			fmt.Printf("Result: %v\n", h)
			cc <- h
			time.Sleep(5*time.Second)


		}()
	}
	for range c{
			fmt.Println("Ok true")
			time.Sleep(5*time.Second)
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

	var n int
	var fp string
	c := make(chan int)
	p := 2

	switch a := len(os.Args); a {
	case 1:
		fp = "/home/alejdg/Workspace/half.txt"
		n = 10
	case 2:
		fp = os.Args[1]
		n = 10
	default:
		fp = os.Args[1]
		n, _ = strconv.Atoi(os.Args[2])
	}
	// h := []int{0}
	// f, err := os.Open(fp)
	// check(err)

	// go populateQueue(c, f)
	// for j:=0; p > 0; p-- {
		go populateQueue(c, fp)
		// fmt.Printf("J:%d\n", j)
		// j++
	// }
	fmt.Println("teste")

	// h, i := topN(c, n, p)
	h := topN(c, n, p)


	elapsed := time.Since(start)
	fmt.Printf("Result: %v\n", h)
	fmt.Printf("Executed in %v\n", elapsed)
	// fmt.Printf("Iteractions: %v\n", i)

}

