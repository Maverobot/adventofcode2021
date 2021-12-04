package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
		os.Exit(2)
	}
}

func hasBit(n int, pos int) bool {
	val := n & (1 << pos)
	return (val > 0)
}

type Report struct {
	ones []int
	size int
}

func (r *Report) addRow(s string) {
	if len(r.ones) == 0 {
		r.ones = make([]int, len(s))
	}
	if len(s) != len(r.ones) {
		panic(errors.New("inconsistent lenght of the row"))
	}
	value, err := strconv.ParseInt(s, 2, 64)
	check(err)
	for i := 0; i < len(s); i++ {
		if hasBit(int(value), i) {
			r.ones[i]++
		}
	}
	r.size++
}

func (r *Report) getGamma() int {
	var gamma int
	for i := 0; i < len(r.ones); i++ {
		if r.ones[i] > r.size/2 {
			gamma |= (1 << i)
		}
	}
	return gamma
}

func (r *Report) getEpsilon() int {
	epsilon := ^r.getGamma()
	// Only keep the lower len(r.ones) bits. Clear the rest bits.
	var mask int
	for i := 0; i < len(r.ones); i++ {
		mask |= (1 << i)
	}
	return epsilon & mask
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s + data-file\n", os.Args[0])
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var report Report
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		report.addRow(row)
	}
	check(scanner.Err())

	fmt.Printf("Answer: %d\n", report.getGamma()*report.getEpsilon())
}
