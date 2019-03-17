package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"regexp"
)

/*type entry struct {
	XMLName   xml.Name `xml:"entry"`
	Timestamp string   `xml:"timestamp,attr"`
	Title     string   `xml:"title,attr"`
	Content   string   `xml:",chardata"`
}*/

/*type entry struct {
	XMLName   xml.Name `xml:"entry"`
	Timestamp string   `xml:"timestamp,attr"`
	Title     string   `xml:"title,attr"`
	Content   string   `xml:",chardata"`
}*/

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

/*type ns struct {
	XMLName xml.Name `xml:"ns"`
	Content string   `xml:",chardata"`
}

type title struct {
	XMLName xml.Name `xml:"title"`
	Content string   `xml:",chardata"`
}

type text struct {
	XMLName xml.Name `xml:"text"`
	Content string   `xml:",chardata"`
}

type timestamp struct {
	XMLName xml.Name `xml:"timestamp"`
	Content string   `xml:",chardata"`
}*/

//var filter, _ = regexp.Compile("^file:.*|^talk:.*|^special:.*|^wikipedia:.*|^wiktionary:.*|^user:.*|^user_talk:.*")

func main() {
	//xmlFile, err := os.Open("wiktionary-latin.xml")
	xmlFile, err := os.Open("enwiktionary-latest-pages-articles.xml")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer xmlFile.Close()

	decoder := xml.NewDecoder(xmlFile)
	r := regexp.MustCompile(`(?m)^={2}[^=]+?={2}`)
	//idx := r.FindAllStringIndex("==abc==", -1)
	//print(idx)

	//total := 0
	//var inElement string
	//for {
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
				}
				//fmt.Print(e.Content[loc[0]:loc[1]])

			}

		default:
		}

		//break
	}
}
