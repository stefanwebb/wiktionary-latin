package main

import (
	"fmt"
	"log"
	"os"
	"wiktionary"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)

	langs, _ := wiktionary.CountLevel2Headings("../../enwiktionary-latest-pages-articles.xml")
	fmt.Println(langs)
}
