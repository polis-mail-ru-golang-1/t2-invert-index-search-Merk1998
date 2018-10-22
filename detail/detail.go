package detail

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func InitMap(mapOfMap map[string]map[string]int) ([]string, error) {
	args := os.Args[1:]
	for _, inputFile := range (args) {
		strSlice, err := ReadString(inputFile)
		if err != nil {
			return nil, err
		}

		wordsSlice := ReadWords(strSlice)
		for _, i := range wordsSlice {
			_, ok := mapOfMap[i]
			if !ok {
				mapOfMap[i] = map[string]int{}
			}
			mapOfMap[i][inputFile]++
		}
	}
	return args, nil
}

func ReadString (inputFile string) ([]string, error) {
	fileName, err := os.Open(inputFile)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := fileName.Close(); err != nil {
			panic(err)
		}
	} ()

	var strBuf []string

	scanner := bufio.NewScanner(fileName)
	for scanner.Scan() {
		strBuf = append(strBuf, scanner.Text())
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return strBuf, nil
}

func ReadWords (strBufSlice []string) ([]string) {
	var wordSlice []string

	for _, s := range strBufSlice {
		strBuf := strings.Split(s, " ")

		for _, s = range(strBuf) {
			wordSlice = append(wordSlice, s)
		}
	}
	return wordSlice
}

func ReadFromStdin() []string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var str []string
	str = append(str, scanner.Text())
	return ReadWords(str)
}

func PrintResult(count map[string]int) {
	for len(count) > 0 {
		var max int
		var tmp string
		for i, val := range count {
			if val > max {
				tmp = i
				max = val
			}
		}
		fmt.Printf("- %s; совпадений - %d\n", tmp, max)
		delete(count, tmp)
	}
}

func PrintMap(mapOfMap map[string]map[string]int) {
	for str, val := range mapOfMap {
		fmt.Printf("%12s %+v\n", str, val)
	}
}