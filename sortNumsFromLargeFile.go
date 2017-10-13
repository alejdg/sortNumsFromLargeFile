package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)
	
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func meta(n int, h []int) []int {
	if n > h[0] {
		h = append(h, n)
	} else if c := contains(h, n); n > h[len(h)-1] && !c {
			h = append(h, n)
	}
	return h
}

func sortAndSize (h []int, n int) []int {
	// Sort the array
	sort.Sort(sort.Reverse(sort.IntSlice(h)))
	// And remove the last if needed
	if len(h) > n {
		h = h[:len(h)-1]
	}
	return h
}

func contains(s []int, e int) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

func main() {
	// n := os.Args[1:]
	n := 10
	h := []int{0}
	c := 0
	// fp := os.Args[2:]
	// f, err := os.Open(fp)

	f, err := os.Open("/home/alejdg/Workspace/ruby/half.txt")
    check(err)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		cn, err := strconv.Atoi(scanner.Text())
		h = meta(cn, h)
		h = sortAndSize(h, n)
		check(err)
		c++
		fmt.Println(c)
		// fmt.Println(h)
		// fmt.Println(scanner.Text()) // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	fmt.Println(h)

}

