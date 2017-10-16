package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"time"
	// "sync"
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

func readFromQueue(c chan int, n int) int {
	n = <- c
	return n
}

func topN(c chan int, n int, h []int) []int {
	// j := 0
	// h := []int{0}
		fmt.Printf("TopN: %v\n", n)
	for range c {
	// for i:= range c {
		// fmt.Printf("Partial[%v]: %v\n", i, h)
		h = sortAndSize(meta(readFromQueue(c, n), h), n)
		// fmt.Printf("Partial[%v]: %v\n", i, h)
	}
	return h
}


func main() {
	start := time.Now()

	var n int
	var fp string
	c := make(chan int, 10000)

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
	h := []int{0}
	// f, err := os.Open(fp)
	// check(err)

	// go populateQueue(c, f)
	go populateQueue(c, fp)

	h = topN(c, n, h)

	elapsed := time.Since(start)
	// fmt.Printf("Queue: \n")
	// for i := range c {
	// 	fmt.Println(i)
	// }
	fmt.Printf("Result: %v\n", h)
	fmt.Printf("Executed in %v\n", elapsed)

}

