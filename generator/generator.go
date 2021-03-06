package generator

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Generator struct
type Generator interface {
	// Generate a random phrase
	Generate() string
}

// generator implements Generator interface
type generator struct {
	Dictionary1 string
	Dictionary2 string

	Dictionary1Lenght int
	Dictionary2Lenght int
}

//New creates a new instance of generator
func New() Generator {
	g := new(generator)
	g.Dictionary1 = "generator/dictionaries/phrases.txt"
	g.Dictionary2 = "generator/dictionaries/auxiliaries.txt"

	g.Dictionary1Lenght = linesInFile(g.Dictionary1)
	g.Dictionary2Lenght = linesInFile(g.Dictionary2)

	return g
}

// Generate a random phrase
func (g *generator) Generate() string {
	phrase := getLine(g.Dictionary1, rand.Int()%g.Dictionary1Lenght)

	var result string
	var buffer string

	for i := 0; i < len(phrase); i++ {
		buffer = string([]rune(phrase)[i])

		if buffer == "-" {
			if string([]rune(phrase)[i+1]) == "-" {
				buffer = getLine(g.Dictionary2, rand.Int()%g.Dictionary2Lenght)
				i++
			}
		}
		result += buffer
	}

	return result
}

// linesInFile counts the lines in the file
func linesInFile(fileName string) int {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)
	var count int = 0

	for scanner.Scan() {
		scanner.Text()
		count++
	}

	defer f.Close()
	return count
}

// getLine get the line in the n position
func getLine(fileName string, n int) string {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)
	var line string

	for i := 0; i < n; i++ {
		line = scanner.Text()
	}

	defer f.Close()
	return line
}
