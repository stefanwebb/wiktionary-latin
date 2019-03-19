package main

import (
	"encoding/xml"
	"wiktionary"
)

// Represents a <data> element
type page struct {
	XMLName  xml.Name   `xml:"page"`
	Title    string     `xml:"title"`
	Ns       int        `xml:"ns"`
	Revision []revision `xml:"revision"`
}

// Represents an <entry> element
type revision struct {
	Timestamp string `xml:"timestamp"`
	Text      string `xml:"text"`
}

func main() {
	langs = wiktionary.CountLevel2Headings("enwiktionary-latest-pages-articles.xml")
	print(langs)
}
