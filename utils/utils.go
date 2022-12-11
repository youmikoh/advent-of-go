package utils

import (
	"bufio"
	"math"
	"os"
	"strconv"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func LoadFile(fileName string) *os.File {
	file, err := os.Open(fileName)
	Check(err)
	return file
}

func StreamFile[T any](fileName string, lineParser func(string) T) <-chan T {
	stream := make(chan T)

	go func() {
		defer close(stream)

		file := LoadFile(fileName)

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			stream <- lineParser(line)
		}

		if err := scanner.Err(); err != nil {
			Check(err)
		}
	}()

	return stream
}

func StreamInput[T any](lineParser func(string) T) <-chan T {
	return StreamFile("input.txt", lineParser)
}

func StreamTest[T any](lineParser func(string) T) <-chan T {
	return StreamFile("test.txt", lineParser)
}

func Base2To10(binary []int) int {
	var sum float64
	p := len(binary) - 1
	for i, b := range binary {
		sum += float64(b) * math.Pow(2, float64(p-i))
	}
	return int(sum)
}

func Strings2Ints(strings []string) []int {
	var ints []int
	for _, s := range strings {
		i, err := strconv.Atoi(s)
		Check(err)
		ints = append(ints, i)
	}
	return ints
}

func Abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}

func Sign(x int) int {
	if x < 0 {
		return -1
	}
	return 1
}

func Reverse[T any](slice []T) []T {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}
