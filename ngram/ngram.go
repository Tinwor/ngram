package ngram

import (
	"bufio"
	"os"
	"strings"
)

type Ngram struct {
}

func NewNgram() *Ngram {
	return &Ngram{}
}

type CustomSplit func(string) []string

func (n Ngram) GetNGramFromArray(minNgramLen, maxNgramLen int, words []string) map[string]int {
	dictionary := make(map[string]int)
	var array []string
	for _, word := range words {
		array = append(array, word)
		if len(array) >= maxNgramLen {
			for i := 0; i < maxNgramLen-minNgramLen+1; i++ {
				key := strings.Join(array[:minNgramLen+i], " ")
				dictionary[key]++
			}
			array = array[1:]
		}
	}

	return dictionary
}

func (n Ngram) GetNGramFromFile(filePath string, minNgramLen, maxNgramLen int, fn CustomSplit) (map[string]int, error) {
	dictionary := make(map[string]int)
	var array []string
	file, err := os.Open(filePath)
	if err != nil {
		defer file.Close()
		return nil, err

	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words := fn(scanner.Text())
		for _, word := range words {
			array = append(array, word)
			if len(array) >= maxNgramLen {
				for i := 0; i < maxNgramLen-minNgramLen+1; i++ {
					key := strings.Join(array[:minNgramLen+i], " ")
					dictionary[key]++
				}
				array = array[1:]
			}
		}
	}
	return dictionary, nil
}
