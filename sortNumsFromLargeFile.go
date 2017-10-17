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

// Read the numbers on file and put them on a queue
func putOnQueue(c chan int, fp string, p int) {
	defer close(c)
	f, _ := os.Open(fp)
	s := bufio.NewScanner(f)
	for s.Scan() {
		cn, _ := strconv.Atoi(s.Text())
		c <- cn
	}
	if err := s.Err(); err != nil {
		fmt.Println(os.Stderr, "reading standard input:", err)
	}
}

func topN(c chan int, n, p int) ([]int) {
	var wg sync.WaitGroup
	wg.Add(p)
	var mutex = &sync.Mutex{}
	h := []int{0}
	cc := make(chan []int)

	for ; p > 0; p-- {
		go func() {
			for cn := range c {
				mutex.Lock()
				h = sortAndSize(meta(cn, h), n)
				mutex.Unlock()
			}
			cc <- h
			wg.Done()
		}()
	}
	return <-cc
}

func main() {
	start := time.Now()

	var n, p int
	var fp string
	c := make(chan int, 1000000)

	// 
	switch a := len(os.Args); a {
	case 1:
		fp = "/Workspace/large_file.txt"
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

	go putOnQueue(c, fp, p)

	h := topN(c, n, p)

	elapsed := time.Since(start)
	fmt.Printf("Result: %v\n", h)
	fmt.Printf("Executed in %v\n", elapsed)
}

