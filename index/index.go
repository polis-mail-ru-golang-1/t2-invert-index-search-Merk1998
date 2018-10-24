package index

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func InitMap(args []string) (map[string]map[string]int, error) {
	mapOfMap := make(map[string]map[string]int)
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
	return mapOfMap, nil
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
			if s == "" {
				continue
			}
			wordSlice = append(wordSlice, s)
		}
	}
	return wordSlice
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