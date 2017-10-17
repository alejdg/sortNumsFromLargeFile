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
	for len(h) > n {
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
	h := []int{0}
	cc := make(chan []int)

	for ; p > 0; p-- {
		go func() {
			hn := []int{0}
			for cn := range c {
				hn = sortAndSize(meta(cn, hn), n)
			}
			cc <- hn
			wg.Done()
		}()
	}
	go func(){
		wg.Wait()
		for cv := range cc {
			h = append(h, cv...)
	  }
	}()
  h = sortAndSize(h, n)
	return <-cc
}

func main() {
	start := time.Now()

	var n, p int
	var fp string
	c := make(chan int, 10000000)

	switch a := len(os.Args); a {
	case 1:
		fp = "/Workspace/large_file.txt"
		n = 10
		p = 2
	case 2:
		fp = os.Args[1]
		n = 10
		p = 2
	case 3:
		fp = os.Args[1]
		n, _ = strconv.Atoi(os.Args[2])
		p = 2
	default:
		fp = os.Args[1]
		n, _ = strconv.Atoi(os.Args[2])
		p, _ = strconv.Atoi(os.Args[3])
	}

	go putOnQueue(c, fp, p)

	h := topN(c, n, p)

	elapsed := time.Since(start)
	fmt.Printf("Result: %v\n", h)
	fmt.Printf("Executed in %v with %d workers.\n", elapsed, p)
}

