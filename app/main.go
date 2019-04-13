package main

import (
	"fmt"
	"go/build"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/stefanwebb/wiktionary-latin/wiktionary"
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

	datapath := path.Join(build.Default.GOPATH, "/data")
	//inFilename := "enwiktionary-latest-pages-articles.xml"
	outFilename := "enwiktionary-latin.xml"

	//wiktionary.ExtractLatin(filepath.Join(path, inFilename), filepath.Join(path, outFilename))

	headwordsFilename := "latin-headwords.txt"
	wiktionary.SortLatinHeadwords(filepath.Join(datapath, outFilename), filepath.Join(datapath, headwordsFilename))

	//pages, _ := wiktionary.CountPages(filepath.Join(path, filename))
	//fmt.Printf("\n%d pages in total\n", pages)

	//_ = wiktionary.FindLevel2HeadingsTypos(filepath.Join(path, filename))

	// Sort keys, i.e. language headers, and save counts to file in sorted order
	/*langs, _ := wiktionary.CountLevel2Headings(filepath.Join(path, filename))
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
	w.Flush()*/
}
