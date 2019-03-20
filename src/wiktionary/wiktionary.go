package wiktionary

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"regexp"
	"time"
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

// CountLevel2Headings counts number of unique second level headings in Wiktionary XML
func CountLevel2Headings(filename string) (map[string]int, error) {
	// Counts of language tags
	var countLanguages = map[string]int{}
	countPages := 0

	// Open XML file and start decoder
	xmlFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, nil
	}
	defer xmlFile.Close()
	decoder := xml.NewDecoder(xmlFile)

	// This regular expression looks for second level headings, e.g. "==English=="
	r := regexp.MustCompile(`(?m)^={2}[^=]+?={2}`)

	fmt.Println("Counting second level headings in Wiktionary XML")
	start := time.Now()
	for {
		// Read tokens from the XML document in a stream.
		t, _ := decoder.Token()
		if err == io.EOF {
			break
		}

		switch se := t.(type) {
		case xml.StartElement:
			//fmt.Println(se.Name.Local)
			//fmt.Println(t)

			//if se.Name.Local == "entry" {
			if se.Name.Local == "page" {
				t, _ = decoder.Token()
				//for t.(type) != xml.StartElement

				var p page
				// decode a whole chunk of following XML into the
				// variable p which is a Page (se above)
				decoder.DecodeElement(&p, &se)
				if p.Ns != 0 {
					continue
				}
				//print(p.Revision[0].Text)

				loc := r.FindAllStringIndex(p.Revision[0].Text, -1)
				//print(loc)
				for _, idx := range loc {
					//substr := p.Revision[0].Text[idx[0]:idx[1]]
					//fmt.Println(substr)
					countLanguages[p.Revision[0].Text[(idx[0]+2):(idx[1]-2)]]++
				}

				countPages++

				if countPages%100000 == 0 {
					fmt.Printf("%d pages...\n", countPages)
				}
			}

		default:
		}

	}

	elapsed := time.Since(start).Minutes()
	fmt.Printf("Took %f minutes\n", elapsed)

	return countLanguages, nil
}

// CountPages counts number of pages in NS 0, i.e. word pages
func CountPages(filename string) (int, error) {
	// Counts of language tags
	var countPages = 0

	// Open XML file and start decoder
	xmlFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return 0, nil
	}
	defer xmlFile.Close()
	decoder := xml.NewDecoder(xmlFile)

	fmt.Println("Counting pages in Wiktionary XML")
	start := time.Now()
	for {
		// Read tokens from the XML document in a stream.
		t, err := decoder.Token()
		if err == io.EOF {
			break
		}

		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local == "page" {
				t, _ = decoder.Token()
				var p page
				// decode a whole chunk of following XML into the
				// variable p which is a Page (se above)
				decoder.DecodeElement(&p, &se)
				if p.Ns != 0 {
					continue
				}

				countPages++

				if countPages%100000 == 0 {
					fmt.Printf("%d pages...\n", countPages)
				}
			}

		default:
		}

	}

	elapsed := time.Since(start).Minutes()
	fmt.Printf("Took %f minutes\n", elapsed)

	return countPages, nil
}
