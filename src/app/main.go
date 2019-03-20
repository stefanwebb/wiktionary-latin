package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"wiktionary"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)

	path := "C:/Users/Stefan Webb/OneDrive - OnTheHub - The University of Oxford/Programming/Go/latin-wiktionary"
	filename := "enwiktionary-latest-pages-articles.xml"
	//pages, _ := wiktionary.CountPages(filepath.Join(path, filename))
	//fmt.Printf("\n%d pages in total\n", pages)

	langs, _ := wiktionary.CountLevel2Headings(filepath.Join(path, filename))
	//fmt.Println(langs)

	// Sort keys, i.e. language headers, and save counts to file in sorted order
	keys := make([]string, 0, len(langs))
	for k := range langs {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	f, err := os.Create("wiktionary-lang-counts.txt")
	check(err)
	defer f.Close()

	w := bufio.NewWriter(f)
	for _, k := range keys {
		w.WriteString(fmt.Sprintf("%s, %d\n", k, langs[k]))
		//fmt.Println(k, langs[k])
	}
	w.Flush()
}
