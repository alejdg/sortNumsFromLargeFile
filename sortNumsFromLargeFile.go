package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"time"
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

func main() {
	start := time.Now()
	// n := os.Args[1:]
	n := 10
	h := []int{0}
	// fp := os.Args[2:]
	// f, err := os.Open(fp)

	// f, err := os.Open("/home/agomes/Workspace/1m.txt")
	f, err := os.Open("/home/agomes/Workspace/temp.txt")
	// f, err := os.Open("/home/agomes/Workspace/large_file.txt")
  check(err)

  // Scan the file line by line to avoid putting the whole file in memory
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		cn, err := strconv.Atoi(scanner.Text())
		check(err)
		h = sortAndSize(meta(cn, h), n)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	elapsed := time.Since(start)
	fmt.Printf("Result: %v\n", h)
	fmt.Printf("Executed in %v\n", elapsed)

}

