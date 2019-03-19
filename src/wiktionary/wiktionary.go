package wiktionary

import (
	"encoding/xml"
	"fmt"
	"os"
	"regexp"
)

// CountLevel2Headings counts number of unique second level headings in Wiktionary XML
func CountLevel2Headings(filename string) {
	// Counts of language tags
	var countLanguages = map[string]int{}

	// Open XML file and start decoder
	xmlFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer xmlFile.Close()
	decoder := xml.NewDecoder(xmlFile)

	// This regular expression looks for second level headings, e.g. "==English=="
	r := regexp.MustCompile(`(?m)^={2}[^=]+?={2}`)

	// TODO: Replace 1000 loops with an EOF cond
	for i := 0; i < 1000; i++ {
		// Read tokens from the XML document in a stream.
		t, _ := decoder.Token()

		//if t == nil {
		//	break
		//}

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
					substr := p.Revision[0].Text[idx[0]:idx[1]]
					fmt.Println(substr)
					countLanguages[p.Revision[0].Text[(idx[0]+2):(idx[1]-2)]]
				}
				//fmt.Print(e.Content[loc[0]:loc[1]])

			}

		default:
		}

		//break
	}
}
