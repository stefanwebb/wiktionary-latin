package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"wiktionary"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)

	//langs, _ := wiktionary.CountLevel2Headings("../../enwiktionary-latest-pages-articles.xml")
	//fmt.Println(langs)

	path := "C:/Users/Stefan Webb/OneDrive - OnTheHub - The University of Oxford/Programming/Go/latin-wiktionary"
	filename := "enwiktionary-latest-pages-articles.xml"
	pages, _ := wiktionary.CountPages(filepath.Join(path, filename))
	fmt.Printf("\n%d pages in total\n", pages)
}
