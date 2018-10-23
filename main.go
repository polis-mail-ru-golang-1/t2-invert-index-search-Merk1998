package main

import (
	"bufio"
	"fmt"
	"os"
	index "github.com/polis-mail-ru-golang-1/t2-invert-index-search-Merk1998/index"
	//"./index"
)

func main() {
	args := os.Args[1:]
	invertIndexMap, err := index.InitMap(args)
	if err != nil {
		fmt.Println(err)
		return
	}

	index.PrintMap(invertIndexMap)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var str []string
	str = append(str, scanner.Text())
	str1 := index.ReadWords(str)
	c := make(map[string]int)
	for i, item := range invertIndexMap {
		for fileName, count := range item {
			for _, word := range str1 {
				if i == word {
					c[fileName] += count
				}
			}
		}
	}
	index.PrintResult(c)
}