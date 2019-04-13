package wiktionary

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strings"
	"time"
)

// Represents a <data> element
type lemma struct {
	XMLName   xml.Name `xml:"lemma"`
	Title     string   `xml:"title,attr"`
	Timestamp string   `xml:"timestamp,attr"`
	Contents  string   `xml:",chardata"`
}

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
		t, err := decoder.Token()
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

// FindLevel2HeadingsTypos counts number of unique second level headings in Wiktionary XML
func FindLevel2HeadingsTypos(filename string) error {
	// Counts of language tags
	countPages := 0

	// Open XML file and start decoder
	xmlFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer xmlFile.Close()
	decoder := xml.NewDecoder(xmlFile)

	// This regular expression looks for second level headings, e.g. "==English=="
	r := regexp.MustCompile(`(?m)^={2}[^=]+?={2}`)

	fmt.Println("Counting second level headings in Wiktionary XML")
	start := time.Now()
	for {
		// Read tokens from the XML document in a stream.
		t, err := decoder.Token()
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
				// variable p which is a Page (see above)
				decoder.DecodeElement(&p, &se)
				if p.Ns != 0 {
					continue
				}
				//print(p.Revision[0].Text)

				loc := r.FindAllStringIndex(p.Revision[0].Text, -1)
				//print(loc)
				for _, idx := range loc {

					lang := p.Revision[0].Text[(idx[0] + 2):(idx[1] - 2)]
					switch lang {
					case "your mom lollol":
						fmt.Printf("%s => %s\n", p.Title, lang)
					case "West Frisian, Dutch, English and German":
						fmt.Printf("%s => %s\n", p.Title, lang)
					case "Mecayapan Nahautl":
						fmt.Printf("%s => %s\n", p.Title, lang)
					default:
						if lang[0] == ' ' {
							fmt.Printf("%s => %s\n", p.Title, lang)
						}
					}
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

	return nil
}

// ExtractLatin extracts the Latin entries of Wiktionary
func ExtractLatin(inputFilename string, outputFilename string) error {
	// Counts of language tags
	countPages := 0

	// Open XML file and start decoder
	inFile, err := os.Open(inputFilename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer inFile.Close()
	decoder := xml.NewDecoder(inFile)

	// Open XML file and start encoder
	outFile, err := os.Create(outputFilename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer outFile.Close()
	encoder := xml.NewEncoder(outFile)

	// This regular expression looks for second level headings, e.g. "==English=="
	r := regexp.MustCompile(`(?m)^={2}[^=]+?={2}`)

	fmt.Println("Extracting Latin entries from Wiktionary XML")
	start := time.Now()
	for {
		// Read tokens from the XML document in a stream.
		t, err := decoder.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error decoding token:", err)
			break
		}

		switch se := t.(type) {
		case xml.StartElement:
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
				for i, idx := range loc {

					lang := p.Revision[0].Text[(idx[0] + 2):(idx[1] - 2)]
					if lang == "Latin" {
						var contents string
						if i < len(loc)-1 {
							contents = strings.TrimSpace(p.Revision[0].Text[idx[1]:loc[i+1][0]])
						} else {
							contents = strings.TrimSpace(p.Revision[0].Text[idx[1]:])
						}
						word := &lemma{Title: p.Title, Timestamp: p.Revision[0].Timestamp, Contents: contents}
						encoder.Encode(word)
					}
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

	return nil
}

// SortLatinHeadwords processes the extracted Latin XML into a sorted text file of headwords
func SortLatinHeadwords(inputFilename string, outputFilename string) error {
	// Counts of language tags
	countPages := 0
	headwords := make([]string, 0, 10000)

	// Open XML file and start decoder
	inFile, err := os.Open(inputFilename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer inFile.Close()
	decoder := xml.NewDecoder(inFile)

	// Open text file for output
	outFile, err := os.Create(outputFilename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer outFile.Close()

	fmt.Println("Reading Latin headwords")
	start := time.Now()
	for {
		// Read tokens from the XML document in a stream.
		t, err := decoder.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error decoding token:", err)
			break
		}

		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local == "lemma" {
				t, _ = decoder.Token()
				//for t.(type) != xml.StartElement

				var l lemma
				// decode a whole chunk of following XML into the
				// variable p which is a Page (se above)
				decoder.DecodeElement(&l, &se)

				headwords = append(headwords, strings.ToLower(l.Title))

				if countPages%10000 == 0 {
					fmt.Printf("%d pages...\n", countPages)
				}

				countPages++
			}

		default:
		}

	}

	fmt.Println("Sorting Latin headwords")
	sort.Strings(headwords)

	fmt.Println("Saving to text file")
	isValidWord := regexp.MustCompile(`^[a-zA-Z-.0-9 ]+$`).MatchString
	w := bufio.NewWriter(outFile)
	lastString := ""
	for _, s := range headwords {
		if s == lastString {
			continue
		}

		if isValidWord(s) {
			w.WriteString(s)
			w.WriteString("\n")
		} else {
			fmt.Println(s)
		}
		lastString = s
	}
	w.Flush()

	elapsed := time.Since(start).Minutes()
	fmt.Printf("Took %f minutes\n", elapsed)

	return nil
}
