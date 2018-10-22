package main

import (
	"fmt"
	detail "github.com/polis-mail-ru-golang-1/t2-invert-index-search-Merk1998/detail"
)

func main() {
	mapOfMap := make(map[string]map[string]int)

	_, err := detail.InitMap(mapOfMap)
	if err != nil {
		fmt.Println(err)
		return
	}

	detail.PrintMap(mapOfMap)

	str := detail.ReadFromStdin()
	c := make(map[string]int)
	for i, item := range mapOfMap {
		for fileName, count := range item {
			for _, word := range str {
				if i == word {
					c[fileName] += count
				}
			}
		}
	}

	detail.PrintResult(c)
}