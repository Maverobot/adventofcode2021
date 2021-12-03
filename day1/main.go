package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	MeasurementsLen = 3
)

type Measurement struct {
	values []int
}

func (m *Measurement) append(value int) {
	m.values = append(m.values, value)
	if len(m.values) > MeasurementsLen {
		m.values = m.values[1:]
	}
}

func (m *Measurement) isComplete() bool {
	return len(m.values) == MeasurementsLen
}

func (m *Measurement) sum() int {
	if !m.isComplete() {
		panic(errors.New("sum cannot be computed for incomplete measurements"))
	}

	result := 0
	for _, value := range m.values {
		result += value
	}
	return result
}

func check(e error) {
	if e != nil {
		fmt.Println(e)
		os.Exit(2)
	}
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

	scanner := bufio.NewScanner(file)
	var first = true
	var startCheck = false
	var window1 Measurement
	var window2 Measurement
	var count = 0
	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		check(err)
		if first {
			first = false
			window1.append(value)
			continue
		}
		window2.append(value)

		if window1.isComplete() && window2.isComplete() {
			startCheck = true
			if window2.sum() > window1.sum() {
				count++
			}
		} else if startCheck {
			break
		}
		window1.append(value)
	}
	check(scanner.Err())

	fmt.Printf("Answer: %d\n", count)
}
