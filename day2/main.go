package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Submarine struct {
	position int
	depth    int
}

func (s *Submarine) exec(c []string) error {
	if len(c) != 2 {
		return errors.New("invalid command")
	}

	commandType := c[0]
	commandValue, err := strconv.Atoi(c[1])
	if err != nil {
		return err
	}
	switch commandType {
	case "forward":
		s.position += commandValue
		return nil
	case "down":
		s.depth += commandValue
		return nil
	case "up":
		s.depth -= commandValue
		return nil
	}

	return errors.New("invalid command")
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

	var submarine Submarine
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		command := strings.Fields(scanner.Text())
		if err = submarine.exec(command); err != nil {
			panic(err)
		}
	}
	check(scanner.Err())

	fmt.Printf("Answer: %d\n", submarine.depth*submarine.position)
}
